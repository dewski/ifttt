# ifttt

 [![](https://godoc.org/github.com/dewski/ifttt?status.svg)](http://godoc.org/github.com/dewski/ifttt)

```go
package main

import (
  "fmt"
  "github.com/dewski/ifttt"
)

func main() {
  // Automatically reads IFTTT_MAKER_KEY if set
  ifttt.SetKey("1234")
  values := ifttt.Values{
    Value1: "1600 Pennsylvania Ave NW, Washington, DC 20500",
    Value2: 1,
    Value3: 45,
  }
  err := ifttt.Deliver("trip-finish", values)
  if err != nil {
    panic(err)
  }

  fmt.Println("Success!")
}
```
