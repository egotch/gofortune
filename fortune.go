package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func main()  {

  // define the command to run
  fortuneCommand := exec.Command("fortune", "-f")
  pipe, err := fortuneCommand.StderrPipe()
  if err != nil {
    panic(err)
  }

  // kick off the command and grab the stderr output
  fortuneCommand.Start()
  outputStream := bufio.NewScanner(pipe)
  outputStream.Scan()
  fmt.Println(outputStream.Text())

  // get the path from the returned line
  line := outputStream.Text()
  path := line[strings.Index(line, "/"):]
  fmt.Println(string(path))
}

