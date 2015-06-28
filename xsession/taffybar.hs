import System.Taffybar

import System.Taffybar.Systray
import System.Taffybar.TaffyPager
import System.Taffybar.SimpleClock
import System.Taffybar.FreedesktopNotifications

--import System.Information.Memory
--import System.Information.CPU

main = do
  let clock = textClockNew Nothing "<span fgcolor='orange'>%a %b %_d %l:%M %p</span>" 1
      pager = taffyPagerNew defaultPagerConfig
      note = notifyAreaNew defaultNotificationConfig
      tray = systrayNew
  defaultTaffybar defaultTaffybarConfig { startWidgets = [ pager ]
                                        , endWidgets = [ tray, clock, note ]
                                        }
