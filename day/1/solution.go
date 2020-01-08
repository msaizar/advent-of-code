package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
  "math"
)

func main() {
  var filename = "input.txt"
  var total = 0
  var j = 0
  content, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Println(err)
  }
  lines := strings.Split(string(content), "\n")
  for _, value := range lines {
    i, err := strconv.Atoi(value)
    if err != nil {
      j = 0
    } else {
      j = int(math.Floor(float64(i)/3) - 2)
    }
    total = total + j
	}
  fmt.Println(total)
}
