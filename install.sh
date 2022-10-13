#!/bin/sh -i

case $- in
*i*)
  echo "Please enter an installation directory, or hit enter for default [$HOME/.local/bin]"
  echo "-----------------------------------------------------------------------------------"
  read -r INSTALLDIR # interactive shell
  if [ "$INSTALLDIR" = "" ]
  then
   eval INSTALLDIR="$HOME/.local/bin"
  fi
  eval INSTALLDIR="$INSTALLDIR"
;;
*)
  INSTALLDIR="$HOME/.local/bin" # non-interactive shell
;;
esac

if [ ! -d "$INSTALLDIR" ]
then
  mkdir -p "$INSTALLDIR"
fi

wget -O "$INSTALLDIR/mopidy-control-go" https://github.com/lcyvin/mopidy-monitor-go/releases/download/0.5.1/mopidy-control-go
wget -O "$INSTALLDIR/mopidy-monitor-go" https://github.com/lcyvin/mopidy-monitor-go/releases/download/0.5.1/mopidy-monitor-go
chmod +x "$INSTALLDIR/mopidy-monitor-go"
chmod +x "$INSTALLDIR/mopidy-control-go"

echo "Installation complete. Ensure that $INSTALLDIR is in your PATH to use the program without its full path"
