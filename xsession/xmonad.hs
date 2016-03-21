import XMonad
import XMonad.Config.Desktop
import XMonad.Util.EZConfig

import XMonad.Hooks.ManageDocks
import XMonad.Layout.Fullscreen
import XMonad.Layout.NoBorders

import XMonad.Hooks.EwmhDesktops (ewmh)
import System.Taffybar.Hooks.PagerHints (pagerHints)

-- all this stuff is cargo-culted to the max
-- I'm taking a friend's advice:
-- "just treat Haskell as a domain-specific language for window management"

lowerVolume    = "<XF86AudioLowerVolume>"
lowerVolumeCMD = "amixer set Master 3%-"
raiseVolume    = "<XF86AudioRaiseVolume>"
raiseVolumeCMD = "amixer set Master 3%+ unmute"
muteVolume     = "<XF86AudioMute>"
muteVolumeCMD  = "amixer set Master toggle"

myKeys = [ ("<Print>"    , spawn "gnome-screenshot -i" ) -- screenshot
         , (raiseVolume  , spawn raiseVolumeCMD ) -- raise volume
         , (lowerVolume  , spawn lowerVolumeCMD ) -- lower volume
         , (muteVolume   , spawn muteVolumeCMD ) -- mute volume
         ]

myLayout = avoidStruts (
    Tall 1 (3/100) (1/2) |||
    Mirror (Tall 1 (3/100) (1/2))) |||
    noBorders (fullscreenFull Full)

myConfig = desktopConfig
    {
        terminal = "urxvt",
        layoutHook = myLayout,
        modMask = mod4Mask
    }
    `additionalKeysP` myKeys

main = do
    xmonad $ ewmh $ pagerHints $ myConfig
