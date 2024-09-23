package helpers

import (
  "log"
  "strings"
  "os"
	"path/filepath"
  "math/rand"
)


// visit - searches provided path for valid fortunes to store for use later
//
// includes flags for including TYPES of read backs
//
// ability to filter on fortunes, literature, and/or riddles
// and any combination there in
func Visit(root string, fortunes bool, literature bool, riddles bool)([]string, error)  {
  
  var files []string

  // if all flags are false, return anything
  if ! fortunes && ! literature && ! riddles {
    fortunes = true
    literature = true
    riddles = true
  }

  err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {

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

    // check if we care about fortunes
    if strings.Contains(path, "fortune") && fortunes {
      files = append(files, path)
    }
    // check if we care about riddles
    if strings.Contains(path, "riddle") && riddles {
      files = append(files, path)
    }
    // check if we care about literature
    if strings.Contains(path, "literature") && literature {
      files = append(files, path)
    }
    return nil
  })

  return files, err
}


// randomInt returns in >= min, < max
func RandomInt(min int, max int)int  {

  rtn := min + rand.Intn(max-min)

  return rtn
  
}
