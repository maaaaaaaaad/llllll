package main

type gogo interface{
  one(n int) int
}

func one(n int) int {
  return n + 100
}

func main() {
  m := 1
}