package main

import (
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

var files []string

// visit - searches provided path for valid fortunes to store for use later
//
// filters out "offensive" and unnecessary files (*.dat and symlinks)
func visit(path string, f os.FileInfo, err error)error  {
  
  if err != nil{
    log.Fatal(err)
  }
  // excluded "offensive" fortunes
  if strings.Contains(path, "/off/"){
    return nil
  }

  // exclued *.dat files
  if strings.Contains(path, ".dat"){
    return nil
  }

  if strings.Contains(path, ".u8"){
    return nil
  }

  // excluded any sub dirs
  if f.IsDir(){
    return nil
  }

  files = append(files, path)
  return nil

}


// randomInt returns in >= min, < max
func randomInt(min int, max int)int  {

  rtn := min + rand.Intn(max-min)

  return rtn
  
}


// main - main entry point
func main()  {

  // var declarations
  var quotes []string

  // some command line options...
  debugPtr := flag.Bool("d", false, "enable verbose logging")

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

  // walk the root dir of the fortunes
  err = filepath.Walk(root, visit)
  if err != nil {
    log.Panic(err)
  }

  // log the file contents
  if *debugPtr {
    log.Println("Found", len(files), "files:")
    for _, v := range files {
      vSlc := strings.Split(v, "/")
      log.Println("  >", vSlc[len(vSlc)-1])
    }
  }

  // get a random index of a file
  rndIndex := randomInt(0, len(files))
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

