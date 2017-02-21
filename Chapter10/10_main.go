package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"runtime"
)

func (w *PreffixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
        out := make(chan Request)
        go func() {
            for msg := range in {
                s, ok := msg.Data.(string)
             
                if !ok {
                    msg.handler(nil)
                    continue
                }
             
                msg.Data = strings.ToUpper(s)
             
                out <- msg
            }
            close(out)
        }()
        return out
    }
