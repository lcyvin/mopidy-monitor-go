package main

import (
  "fmt"
	"log"
	"net/rpc"
	"time"

	"github.com/jxskiss/mcli"
)

type rpcObj interface{}

type SeekTime struct {
  Amount time.Duration
}

type Step struct {
  Value int
}

func runCmd(cmd string, args any, output any) error {
  client, err := rpc.DialHTTP("tcp", ":9834")
  if err != nil {
    return err
  }

  err = client.Call(cmd, args, output)
  if err != nil {
    return err
  }

  return nil
} 

func volumeCmd() {
  var args struct {
    Step int `cli:"-s, --step, amount to change volume by (default 2)"`
    Inc bool `cli:"-i, --increment, increment volume by Step amount"`
    Dec bool `cli:"-d, --decrement, decrement volume by Step amount"`
  }
  mcli.Parse(&args)

  if args.Step == 0 {
    args.Step = 2
  }

  if args.Inc && args.Dec {
    log.Fatal("Either inc OR dec should be set, not both")
  } else if (args.Inc || args.Dec) && args.Step < 0 {
    log.Fatal("Either use absolute value for step and pass inc/dec flag, or remove inc/dec flag")
  }
  
  var cmd string
  if args.Dec {
    cmd = "Volume.Decrement"
  } else if args.Step < 0 {
    cmd = "Volume.Decrement"
  } else {
    cmd = "Volume.Increment"
  }

  s := &Step{Value: args.Step}

  err := runCmd(cmd, s, nil)
  if err != nil {
    log.Fatal(err)
  }
}

func seekCmd() {
  var args struct {
    Interval int `cli:"-t, --time, amount to seek forward/backward by in seconds (negative for backward seeking, default 10)"`
    Forward bool `cli:"-f, --forward, set interval to positive (forward) value"`
    Backward bool `cli:"-b, --backward, set interval to negative (backward) value"`
  }
  mcli.Parse(&args)

  if args.Interval == 0 {
    args.Interval = 10
  }

  if args.Forward && args.Backward {
    log.Fatal("Either forward OR backward should be set, not both")
  } else if (args.Forward || args.Backward) && args.Interval < 0 {
    log.Fatal("Either use absolute value for interval and pass forward/backward flag, or remove forward/backward flag")
  }

  var i int
  cmd := "Seek.Jump"
  if args.Backward {
    i = args.Interval*-1
    fmt.Println(i)
  } else {
    i = args.Interval
  }

  st := &SeekTime{Amount: time.Duration(i)*time.Second}
  err := runCmd(cmd, st, nil)
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  mcli.Add("volume", volumeCmd, "control the volume of the music server")
  mcli.Add("seek", seekCmd, "jump forward/backward in the playing track")
  mcli.Run()
}
