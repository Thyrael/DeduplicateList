package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main () {
  fmt.Println("Hello, world.")

  path := os.Args[1] // Get file path
  fmt.Println("File to De-Duplicate : " +path)

  file, err := os.Open(path)
  if err != nil {
    return
  }
  defer file.Close()

  lines := []string{}
  duplicateLines := []string{}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
      line := scanner.Text()
      if len(lines) == 0 {
        lines = append(lines, line)
      } else {
          findDuplicateLine := false
          for i := range lines {
            if strings.Compare(lines[i], line) == 0 {
              fmt.Println("Duplicate line" + line)
              duplicateLines = append(duplicateLines, line)
              findDuplicateLine = true
              break
            }
          }
          if !findDuplicateLine {
            lines = append(lines, line)
          }
      }
  }


  f, err := os.Create("duplicate.txt")
  datawriter := bufio.NewWriter(f)
  for i := range duplicateLines {
    datawriter.WriteString(duplicateLines[i] + "\r\n")
  }

  datawriter.Flush()
  file.Close()
  //fmt.Println(len(lines))

}
