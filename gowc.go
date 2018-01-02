package main

import (
  "fmt"
  "os"
  "bufio"
  "flag"
)

func main() {
  fileNamePtr := flag.String("f", "", "")

  flag.Parse()

  if len(*fileNamePtr) == 0 {
    fmt.Println("you must pass -f=filepath")
    return
  }

  fmt.Println("Counting line for file", *fileNamePtr)

  f, err := os.Open(*fileNamePtr)
  if err != nil {
    fmt.Println(err)
    return
  }

  defer f.Close()

  bufferedReader := bufio.NewReader(f)

  lineCounter := 0
  for {
    // 10 -> is new line (\n)
    _, err := bufferedReader.ReadString(10)
    if err != nil {
      break
    }

    lineCounter++
  }

  fmt.Println("file has", lineCounter, "lines")

}
