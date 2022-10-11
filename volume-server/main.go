package main

import (
  "fmt"
  "log"
  "strconv"
  "net"
  "time"
  "net/http"
  "net/rpc"
  "github.com/fhs/gompd/v2/mpd"
)

var (
  client *mpd.Client
)

type Result struct {
  Err error
  Msgs []string
}

type Step struct {
  Value int
}

type Volume struct {
  //Volume provides volume control functions to client processes
  value int
}

func (v *Volume) get() error {
  vals,err := client.Status()
  if err != nil {
    return err
  }

  i,err := strconv.Atoi(vals["volume"])
  if err != nil {
    return err
  }

  v.value = i
  return nil
}

//Increment increases volume by step amount
func (v *Volume) Increment(s *Step, res *Result) error {
  err := v.get()
  if err != nil {
    return err
  }

  vol := v.value + s.Value
  if vol > 100 {
    vol = 100
  }
  err = client.SetVolume(vol)
  if err != nil {
    return err
  }

  v.get()

  return nil
}

//Decerement decreases volume by step amount
func (v *Volume) Decrement(s *Step, res *Result) error {
  err := v.get()
  if err != nil {
  }

  vol := v.value - s.Value
  if vol < 0 {
    vol = 0
  }
  err = client.SetVolume(vol)
  v.get()
  if err != nil {

  }

  return nil
}

func evtHandler(evt string, client *mpd.Client, resChan chan<- Result) {
  res := Result{}
  vals, err := client.Status()
  if err != nil {
    res.Err = err
  } else {
    res.Msgs = append(res.Msgs, vals["volume"])
  }

  resChan <- res
}

func init() {
  c,err := mpd.Dial("tcp", ":6600")
  if err != nil {
    panic(err)
  }

  client = c
}

func main() {
  vol := new(Volume)
  rpc.Register(vol)
  rpc.HandleHTTP()
  l, e := net.Listen("tcp", ":9834")
  if e != nil {
    log.Fatal(e)
  }

  go http.Serve(l, nil)

  //we need to print a value to stdout for this to show up in polybar
  //on launch
  vals,err := client.Status()
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(vals["volume"])
  }

  //keepalive ticker
  ticker := time.NewTicker(20*time.Second)

  handlerResult := make(chan Result) 

  watcher, err := mpd.NewWatcher("tcp", ":6600", "", "mixer")
  if err != nil {
    log.Fatal(err)
  }
  
  defer ticker.Stop()
  defer watcher.Close()
  defer close(handlerResult)

  for {
    select {
    case evt := <- watcher.Event:
      go evtHandler(evt, client ,handlerResult)
    case watchErr := <- watcher.Error:
      fmt.Println(watchErr)
    case hr := <- handlerResult:
      if hr.Err != nil {
        log.Fatal(err)
      }
      for _,m := range hr.Msgs {
        fmt.Println(m)
      }
    case <- ticker.C:
    err := client.Ping()
    if err != nil {
      log.Fatal(err)
      }
    }
  }
}
