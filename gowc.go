package main

import (
  "fmt"
  "os"
  "bufio"
  "flag"
)

func main() {
      filename := flag.String("f", "", "The file path to count lines")

      flag.Parse()

    if len(*filename) == 0 {
        fmt.Println("you must pass -f=filepath")
        return
    }

    fmt.Println("Counting line for file", *filename, "...")

    f, err := os.Open(*filename)
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

    fmt.Println(*filename, "has", lineCounter, "lines")

}
