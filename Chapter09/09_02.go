package barrier

import (
    "bytes"
    "io"
    "os"
    "strings"
    "testing"
)

func captureBarrierOutput(endpoints ...string) string {
    reader, writer, _ := os.Pipe()

    os.Stdout = writer

    out := make(chan string)
    go func() {
      var buf bytes.Buffer
      io.Copy(&buf, reader)
      out <- buf.String()
    }()

    barrier(endpoints...)

    writer.Close()
    temp := <-out

    return temp
}