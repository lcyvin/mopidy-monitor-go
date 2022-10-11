# polybar-mopidy-volume-monitor
## A simple go program that monitors volume change events in mopidy/MPD for use in polybar

![screenshot of i3wm with polybar showing mopidy's internal volume](1665515329-screenshot.png)

### Usage
In your polybar config, add:
```
[module/mopidy-volume]
type = custom/script
exec = "killall mopidy-volume-server &> /dev/null; /path/to/mopidy-volume-executable"
exec-if = pgrep -x mopidy; or mpd, setting this ensures the module doesn't show up if mopidy/mpd isn't running
tail = true
label = (%output%%); %output% is the token for the stdout values, the second percent and parens are just formatting
scroll-up = /path/to/mopidy-volume-control volume -i --step 2; step arg is optional, default is 2
scroll-down = /path/to/mopidy-volume-control volume -d --step 2
```
remember to remove the comments (in polybar config, comments are denoted with `;`, note that the semicolon in the 
exec-if line is part of a sh command and not a polybar comment)

If you want to add seek forward/backward buttons (eg, for skipping thru a podcast), add the following:
```
[module/mopidy-seek-backward]
type = custom/script
exec = echo " "
exec-if = ps aux | grep mopidy-volume-server | grep -v grep
format = <label>
label = ◀
click-left = /path/to/mopidy-volume-control seek -b -t 30

[module/mopidy-seek-forward]
type = custom/script
exec = echo " "
exec-if = ps aux | grep mopidy-volume-server | grep -v grep
format = <label>
label = ▶
click-left = /path/to//mopidy-volume-control seek -f -t 30
```

### Installation
1. Clone this repository
2. `cd` into the cloned repo and run `make install`
3. Ensure `~/.local/bin` is on your `$PATH`
4. That's it, you're done, test things out and update your polybar config to your liking.
