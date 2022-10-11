#!/bin/sh
echo "Please enter an installation directory, or hit enter for default [$HOME/.local/bin]"
echo "-----------------------------------------------------------------------------------"
read INSTALLDIR
if [ $INSTALLDIR == "" ] 
then
  INSTALLDIR="$HOME/.local/bin"
fi

if [ "${INSTALLDIR:0:1}" != "/" ] && [ "${INSTALLDIR:0:1}" != "~" ] 
then
  INSTALLDIR="$PWD/$INSTALLDIR"
fi

eval INSTALLDIR="$INSTALLDIR"

if [ ! -d "$INSTALLDIR" ] 
then
  mkdir -p $INSTALLDIR
fi

wget -O "$INSTALLDIR/mopidy-control-go" https://github.com/lcyvin/mopidy-monitor-go/releases/download/0.5/mopidy-control-go
wget -O "$INSTALLDIR/mopidy-monitor-go" https://github.com/lcyvin/mopidy-monitor-go/releases/download/0.5/mopidy-monitor-go
chmod +x "$INSTALLDIR/mopidy-monitor-go"
chmod +x "$INSTALLDIR/mopidy-control-go"

echo "Installation complete. Ensure that $INSTALLDIR is in your PATH to use the program without its full path"
