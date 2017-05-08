package main

import (
  "encoding/json"
  "fmt"
  "log"
  "io/ioutil"
  "os"
  "strings"
)

// var pages func = getPages()
// for _, p := range pages {
//     fmt.Println(p.toString())
// }
// fmt.Println(toJson(pages))

// m, err := getMessage(ws)
// if err != nil {
//   log.Fatal(err)
// }

type Page struct {
    keyword string `json:"keyword"`
    pun   string `json:"pun"`
}

// func (p Page) toString() string {
//     return toJson(p)
// }

// func toJson(p interface{}) string {
//     bytes, err := json.Marshal(p)
//     if err != nil {
//         fmt.Println(err.Error())
//         os.Exit(1)
//     }

//     return string(bytes)
// }


// main function here sets up the websocket and runs the main loop for getting messages from the slackgroup

func main() {
  if len(os.Args) != 2 {
    fmt.Fprintf(os.Stderr, "usage: mybot slack-bot-token\n")
    os.Exit(1)
  }

  // start a websocket-based Real Time API session
  ws, id := slackConnect(os.Args[1])
  fmt.Println("mybot ready, ^C exits")

  for {
    // read each incoming message
    m, err := getMessage(ws)
    if err != nil {
      log.Fatal(err)
    }
    // see if we're mentioned
    if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+id+">") {
      // if so try to parse if
      parts := strings.Fields(m.Text)
      if len(parts) >= 1 {
        // looks good, get the quote and reply with the result
        go func(m Message) {
          m.Text = "<sassy>Eh, NO!</sassy>"

          //function call at m.Text = \theFunction()\ should be triggering the response by calling a method that returns the response through postMessage(ws,m)

          postMessage(ws, m)
        }(m)
        // NOTE: the Message object is copied, this is intentional
      } else {
        // huh?
        m.Text = fmt.Sprintf("sorry, that does not compute\n")
        postMessage(ws, m)
      }
    }
  }

}

  func getPages() []Page {
    raw, err := ioutil.ReadFile("./puns.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    var c []Page
    json.Unmarshal(raw, &c)
    return c
  }


/*
mybot - Illustrative Slack bot in Go
Copyright (c) 2015 RapidLoop
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
