package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
  "math"
)

func get_fuel_for_mass(i int) int {
  return int(math.Floor(float64(i)/3) - 2)
}

func main() {
  var filename = "input.txt"
  var total = 0
  content, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Println(err)
  }
  lines := strings.Split(string(content), "\n")
  for _, value := range lines {
    i, err := strconv.Atoi(value)
    if err != nil {
      i = 0
    } else {
      i = get_fuel_for_mass(i)
      for i > 0 {
        total = total + i
        i = get_fuel_for_mass(i)
      }
    }
	}
  fmt.Println(total)
}
