package main

import (
	"bufio"
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
  log.Println(outputStream.Text())

  // get the path from the returned line
  line := outputStream.Text()
  root := line[strings.Index(line, "/"):]

  // walk the root dir of the fortunes
  err = filepath.Walk(root, visit)
  if err != nil {
    log.Panic(err)
  }

  // log the file contents
  log.Println("Found", len(files), "files:")
  for _, v := range files {
    vSlc := strings.Split(v, "/")
    log.Println("  >", vSlc[len(vSlc)-1])
  }

  // get a random index of a file
  rndIndex := randomInt(0, len(files))
  rndFile := files[rndIndex]
  log.Println("fetched random file:", rndFile)
}

