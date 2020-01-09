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
  content, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Println(err)
  }
  lines := strings.Split(strings.TrimSuffix(string(content), "\n"), ",")
  var integer_array = make([]int, 0, len(lines))

  for _, l := range lines {
    if len(l) == 0 { continue }
    n, err := strconv.Atoi(l)
    if err != nil {
      fmt.Println(err)
      return
    }
    integer_array = append(integer_array, n)
  }

  var result = 0
  integer_array[1] = 12
  integer_array[2] = 2
  for position, value := range integer_array {
    if math.Mod(float64(position), 4) == 0 {
      if value == 1 {
        result = integer_array[integer_array[position+1]] + integer_array[integer_array[position+2]]
        integer_array[integer_array[position+3]] = result
      } else if value == 2 {
        result = integer_array[integer_array[position+1]] * integer_array[integer_array[position+2]]
        integer_array[integer_array[position+3]] = result
      } else if value == 99 {
        break
      } else {
        fmt.Println("Unknown OPCODE", value, "at position", position)
      }
    }
  }
  fmt.Println("Position 0: ", integer_array[0])
}
