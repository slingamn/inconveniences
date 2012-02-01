## END provided .bashrc, begin customizations

# prompt with full directory name
PS1="[\u@\h \w]\$ "

alias gfa='git fetch --all --prune'
alias gpp='git pull --prune'
alias gsp='git stash show -p'
alias gsd='git stash drop'
# like in Windows 95 when you could watch the little piece of paper get crumpled
# and soar lazily into the recycle bin:
alias stashdrop='git stash && git stash drop'

# "git branch diff"
function gitbd {
	REMOTE=`git remote | head -n 1`
	git diff $REMOTE/master...HEAD $@
}

# hackish analogue to Alt-F2 "run" in gnome shell
function run {
	$@ & &> /dev/null
}

PATH=$PATH:~/Scripts
