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

  for noun:=0; noun<=99; noun++ {
    for verb:=0; verb<=99; verb++ {
      var integer_array = make([]int, 0, len(lines))

      for _, l := range lines {
        n, err := strconv.Atoi(l)
        if err != nil {
          fmt.Println(err)
          return
        }
        integer_array = append(integer_array, n)
      }

      var result = 0
      integer_array[1] = noun
      integer_array[2] = verb
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
      if integer_array[0] == 19690720 {
        fmt.Println("Noun: ", noun, "Verb: ", verb)
        fmt.Println("Solution: ", 100 * noun + verb)
      }
    }
  }
}
