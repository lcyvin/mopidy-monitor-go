# polybar-mopidy-volume-monitor
## A simple go program that monitors volume change events in mopidy/MPD for use in polybar

### Usage
In your polybar config, add:
```
[module/mopidy-volume]
type = custom/script
exec = "killall mopidy-volume &> /dev/null; /path/to/mopidy-volume-executable"
exec-if = pgrep -x mopidy; or mpd, setting this ensures the module doesn't show up if mopidy/mpd isn't running
tail = true
label = (%output%%); %output% is the token for the stdout values, the second percent and parens are just formatting
scroll-up = /path/to/mopidy-volume-control --step 2 inc; step arg is optional, default is 2
scroll-down = /path/to/mopidy-volume-control --step 2 dec
```
remember to remove the comments (in polybar config, comments are denoted with `;`, note that the semicolon in the 
exec-if line is part of a sh command and not a polybar comment)

### Installation
1. Clone this repository
2. `cd` into the cloned repo and run `make install`
3. Ensure `~/.local/bin` is on your `$PATH`
4. That's it, you're done, test things out and update your polybar config to your liking.
