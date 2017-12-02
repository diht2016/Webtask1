package main

//import "fmt"
import "unicode"

/*
func main() {
  input := []int{0, 3, 2, 5}
  result := RemoveEven(input)
  fmt.Println(result) // Должно напечататься [3 5]
  
  gen := PowerGenerator(2)
  fmt.Println(gen()) // Должно напечатать 2
  fmt.Println(gen()) // Должно напечатать 4
  fmt.Println(gen()) // Должно напечатать 8
  
  fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
  // Должно напечатать 2
}
*/

func RemoveEven(s []int) (a []int) {
  for _, n := range s {
    if n % 2 == 1 {
      a = append(a, n)
    }
  }
  return
}

func PowerGenerator(n int) (func() int) {
  k := 1
  return func() (int) {
    k *= n
    return k
  }
}

func DifferentWordsCount(s string) (count int) {
  next := -1
  word := ""
  dict := make(map[string]bool)

  add := func() {
    if len(word) > 0 {
      if (!dict[word]) {
        count++
        dict[word] = true
      }
      word = ""
    }
  }
  
  for pos, c := range s {
    if unicode.IsLetter(c) {
      if (pos != next) {
        add()
      }
      word = word + string(unicode.ToLower(c))
      next = pos + 1
    }
  }
  
  add()

  return
}
