package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

type IDoer interface {
	DoSomething()
}

type DoerFunc func()

func (f DoerFunc) DoSomething() {
	fmt.Printf("%s\n", "DoSomething")
	f()
}

func Try(d IDoer) {
	d.DoSomething()
}

func TryFunc(d DoerFunc) {
	Try(d)
}

func main() {
	testSuperAdapter()
}

func pipe2() {
	file, _ := os.OpenFile("/tmp/pipe", os.O_CREATE, 0666)
	file.Close()

	file, _ = os.OpenFile("/tmp/pipe", os.O_WRONLY, 0666)
	defer file.Close()

	pipeReader, pipeWriter := io.Pipe()
	defer pipeWriter.Close()
	defer pipeReader.Close()

	counter := Counter{
		Writer: pipeWriter,
	}

	tee := io.TeeReader(pipeReader, file)
	go func() {
		io.Copy(os.Stdout, tee)
	}()

	counter.Count(5)
}

func movePointers() {
	ar := make([]int, 3)
	ar[0] = 0
	ar[1] = 1
	ar[2] = 2

	first := &ar[0]
	second := *first + 1
	fmt.Printf("%s,%s,%s\n", reflect.TypeOf(ar), reflect.TypeOf(first), reflect.TypeOf(second))
}

type MyObject struct{}
type ObjectInherits struct {
	MyObject
}

func compositionIsNotInheritance() {
	composite(&MyObject{})
}

func composite(m *MyObject) {
	println("hello")
}

func pipe1() {

	file1, _ := os.OpenFile("/tmp/pipe", os.O_RDWR, 0666)
	defer file1.Close()

	destinationBytes := make([]byte, 7)
	file1.Read(destinationBytes)

	fmt.Printf("%s\n", string(destinationBytes))

	pr, pw := io.Pipe()
	defer pw.Close()
	defer pr.Close()

	counter := Counter{
		Writer: pw,
	}

	go func(pr io.Reader) {
		tee := io.TeeReader(pr, file1)
		io.Copy(os.Stdout, tee)
	}(pr)

	counter.Count(5)
}

// --------------------------------------------------

type Counter struct {
	Writer io.Writer
}

func (f *Counter) Counter2(n uint64) uint64 {
	if n == 0 {
		println(strconv.Itoa(0))
		return 0
	}

	cur := n
	println(strconv.FormatUint(cur, 10))
	return f.Count(n - 1)
}

func (f *Counter) Count(n uint64) uint64 {
	if n == 0 {
		f.Writer.Write([]byte(strconv.Itoa(0) + "\n"))
		return 0
	}

	cur := n
	f.Writer.Write([]byte(strconv.FormatUint(cur, 10) + "\n"))
	return f.Count(n - 1)
}

type MockReader struct{}

func (m *MockReader) Read(p []byte) (n int, err error) {
	p = []byte("mario")
	return len(p), nil
}

func testSuperAdapter() {
	s := SuperAdapter{&MockReader{}, os.Stdout}
	s.test()
}

type SuperAdapter struct {
	Reader io.Reader
	Writer io.Writer
}

func (s *SuperAdapter) test() {
	var b []byte
	s.Reader.Read(b)
	s.Writer.Write(b)
}
