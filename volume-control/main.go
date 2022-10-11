package main

import (
  "fmt"
  "log"
  "net/rpc"
  "github.com/jxskiss/mcli"
)

type Step struct {
  Value int
}

func main() {
  var serverArg string
  var cmdargs struct {
    Step int `cli:"-s, --step, amount to change volume by" default:2`
    Arg string `cli:"#R, text, the function to run -- can be [Inc]rement or [Dec]rement"`
  } 

  mcli.Parse(&cmdargs)
  switch cmdargs.Arg {
  case "Inc", "inc", "Increment", "increment":
    serverArg = "Volume.Increment"
  case "Dec", "Decrement", "dec", "decrement":
    serverArg = "Volume.Decrement"
  default:
    log.Fatalf("Unknown argument type: %s", cmdargs.Arg)
  }

  // handle default value for step
  if cmdargs.Step == 0 {
    cmdargs.Step = 2
  }

  client,err := rpc.DialHTTP("tcp", ":9834")
  if err != nil {
    log.Fatal(err)
  }

  args := &Step{Value:cmdargs.Step}
  err = client.Call(serverArg, args, nil)
  if err != nil {
    fmt.Println(err)
  }
}
