# polybar-mopidy-volume-monitor
## A simple go program that monitors volume change events in mopidy/MPD for use in polybar. Can also be used to skip forward/back in podcasts and songs

<p align="center">
  <img max-height="500px" src="https://github.com/lcyvin/polybar-mopidy-volume-monitor/blob/main/example.png?raw=true" alt="i3wm with polybar showing a music player and the corresponding polybar output with application-specific volume and skip forward/backward buttons"/>
</p>

### Motivation
Polybar's built-in MPD module handles the bulk of needs for mopidy/mpd-related controls, however, it lacks a few features like static time skipping (polybar's implementation is percentage based) and volume level control. Because mopidy tends to skew a bit louder than the rest of the programs on my computer, I like to manage its volume independently. Additionally, as I use mopidy for podcasts as well as music, being able to skip forward in set *n*-second intervals allows for a much more pleasant listening experience. This project aims to solve those specific gaps, not replace the existing mpd module. 

### Usage
You can find an example polybar config [here](polybar-config.ini), if you want to just copy it straight into
your config file. Don't forget to add the modules to your bar config, eg:
`modules-right = [...] mpd mopidy-volume mopidy-seek-backward mopidy-seek-forward [...]`

If the module gives an error or does not render, ensure that the server and client were installed into your user's ~/.local/bin folder.

In your polybar config, add:
```ini
[module/mopidy-volume]
type = custom/script
exec = "killall mopidy-volume-server &> /dev/null; /path/to/mopidy-volume-executable"
exec-if = pgrep -x mopidy; or mpd, setting this ensures the module doesn't show up if mopidy/mpd isn't running
tail = true
label = (%output%%); results in (XX%) output
scroll-up = /path/to/mopidy-volume-control volume -i --step 2; step arg is optional, default is 2
scroll-down = /path/to/mopidy-volume-control volume -d --step 2
```
remember to remove the comments (in polybar config, comments are denoted with `;`, note that the semicolon in the 
exec-if line is part of a sh command and not a polybar comment)

If you want to add seek forward/backward buttons (eg, for skipping thru a podcast), refer to the following example:
```ini
; using individual modules as opposed to one module with a huge format string to keep things clean
; exec is set to echo " " so that the icon will show up, otherwise polybar will see no output
; and not render the module
[module/mopidy-seek-backward]
type = custom/script
exec = echo " "
exec-if = ps aux | grep mopidy-volume-server | grep -v grep
format = <label>
label = ◀
click-left = /path/to/mopidy-volume-control seek -b -t 30
click-right = /path/to/mopidy-volume-control seek -b -t 10

[module/mopidy-seek-forward]
type = custom/script
exec = echo " "
exec-if = ps aux | grep mopidy-volume-server | grep -v grep
format = <label>
label = ▶
click-left = /path/to/mopidy-volume-control seek -f -t 30
click-right = /path/to/mopidy-volume-control seek -f -t 10
```

### Installation
1. Clone this repository
2. `cd` into the cloned repo and run `make install` as your user, not root
3. Ensure `~/.local/bin` is on your `$PATH`
4. That's it, you're done, test things out and update your polybar config to your liking.
