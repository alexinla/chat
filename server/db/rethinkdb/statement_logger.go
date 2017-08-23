package rethinkdb

import (
    "os"
    "log"
    "fmt"
    "path/filepath"
    "time"
    rdb "gopkg.in/gorethink/gorethink.v2"
)

var LOG_FILENAME = "statements_log.txt"

func logStatement(tag string, st rdb.Term) rdb.Term {
    
    path, _ := os.Executable()
    text := fmt.Sprintf("%v [RethinkDB Statement] %v: %v\n\n", time.Now().Format("2006-01-02 15:04:05"), tag, st)
    f, err := os.OpenFile(filepath.Dir(path) + "/" + LOG_FILENAME, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0600)
    if err != nil {
        log.Println(err)
    }
    defer f.Close()
    f.WriteString(text)
    
    return st
}