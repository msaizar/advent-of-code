package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
  "math"
)

func get_fuel_for_mass(i int) int {
  var fuel = int(math.Floor(float64(i)/3) - 2)
  if fuel > 0 {
    return fuel + get_fuel_for_mass(fuel)
  } else {
    return 0
  }
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
      total = total + i
    }
	}
  fmt.Println(total)
}
