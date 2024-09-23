package main

import (
  // internal
  "github.com/egotch/gofortune/helpers"

  // standard
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)



// main - main entry point
func main()  {

  // var declarations
  var quotes []string

  // some command line options...
  debugPtr := flag.Bool("d", false, "enable verbose logging")
  frtnPtr := flag.Bool("f", false, "pull from FORTUNES section")
  rdlPtr := flag.Bool("r", false, "pull lines from the riddles section")
  litPtr := flag.Bool("l", false, "pull lines from LITERATURE section")


  flag.Parse()


  // define the command to run
  fortuneCommand := exec.Command("fortune", "-f")
  pipe, err := fortuneCommand.StderrPipe()
  if err != nil {
    log.Panic(err)
  }

  // kick off the command and grab the stderr output
  fortuneCommand.Start()
  outputStream := bufio.NewScanner(pipe)
  outputStream.Scan()
  if *debugPtr {
    log.Println(outputStream.Text())
  }

  // get the path from the returned line
  line := outputStream.Text()
  root := line[strings.Index(line, "/"):]

  files, err := helpers.Visit(root, *frtnPtr, *litPtr, *rdlPtr)

  // log the file contents
  if *debugPtr {
    log.Println("Found", len(files), "files:")
    for _, v := range files {
      vSlc := strings.Split(v, "/")
      log.Println("  >", vSlc[len(vSlc)-1])
    }
  }

  // get a random index of a file
  rndIndex := helpers.RandomInt(0, len(files))
  rndFile := files[rndIndex]
  if *debugPtr{
    log.Println("fetched random file:", rndFile)
  }

  // fetch a random line from the random file!
  file, err := os.Open(rndFile)
  if err != nil{
    panic(err)
  }

  defer file.Close()

  lines, err := io.ReadAll(file)
  if err != nil{
    panic(err)
  }

  //quotes = strings.Split(string(lines), "%")
  tmp_lines := strings.Split(string(lines), "%")
  quotes = append(quotes, tmp_lines...)
 
  // get a random quote and print it out
  rndQuote := quotes[randomInt(0, len(quotes))]
  fmt.Println(rndQuote)


}

