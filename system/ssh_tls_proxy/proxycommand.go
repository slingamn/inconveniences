package main

import (
	"crypto/tls"
	"io"
	"log"
	"os"
	"sync"
)

// This is a ProxyCommand for ssh that wraps the ssh protocol in TLS. Follow these steps:
// 1. Compile this file with `go build proxycommand.go`
// 2. Install it in your home directory on the client machine,
// e.g., `mv proxycommand ~/bin/proxycommand_ssh_in_tls`
// 3. Add the following block to ~/.ssh/config:
// Host my-host.domain
//     ProxyCommand ~/bin/proxycommand_ssh_in_tls %h:443
// 4. If you don't have have a valid TLS certificate, add `insecure` to the command
// (this is probably OK because you can rely on the known_hosts file to authenticate
// the actual sshd):
// Host my-host.domain
//     ProxyCommand ~/bin/proxycommand_ssh_in_tls %h:443 insecure

// connects two io.ReadWriteCloser; reads from the first are written to the second,
// and vice versa
type Socat struct {
	c1 io.ReadWriteCloser
	c2 io.ReadWriteCloser

	done      chan error
	closeOnce sync.Once
	closeErr  error
}

func NewSocat(c1, c2 io.ReadWriteCloser) *Socat {
	c := &Socat{
		c1:   c1,
		c2:   c2,
		done: make(chan error, 2),
	}
	go c.funnel(c1, c2)
	go c.funnel(c2, c1)
	return c
}

func (t *Socat) funnel(d1, d2 io.ReadWriteCloser) {
	_, err := io.Copy(d1, d2)
	t.done <- err
}

func (t *Socat) Wait() (err error) {
	err = <-t.done
	t.Close()
	return err
}

func (t *Socat) Close() (err error) {
	t.closeOnce.Do(func() {
		t.closeErr = t.realClose()
	})
	return t.closeErr
}

func (t *Socat) realClose() (err error) {
	e1 := t.c1.Close()
	e2 := t.c2.Close()
	if e1 != nil {
		return e1
	}
	return e2
}

// StdioRWC wraps os.Stdin and os.Stdout to expose an io.ReadWriteCloser.
type StdioRWC struct{}

func (s StdioRWC) Read(buf []byte) (n int, err error) {
	return os.Stdin.Read(buf)
}

func (s StdioRWC) Write(buf []byte) (n int, err error) {
	return os.Stdout.Write(buf)
}

func (s StdioRWC) Close() (err error) {
	return nil
}

func main_(args []string) (code int) {
	if len(args) < 1 {
		log.Printf("insufficient arguments\n")
		return 1
	}
	addr := args[0]
	var conf *tls.Config
	if len(args) > 1 && args[1] == "insecure" {
		conf = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	conn, err := tls.Dial("tcp", addr, conf)
	if err != nil {
		log.Printf("error dialing remote server %s: %v\n", addr, err)
		return 1
	}
	socat := NewSocat(conn, StdioRWC{})
	socat.Wait()
	return 0
}

func main() {
	os.Exit(main_(os.Args[1:]))
}
