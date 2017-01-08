package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
)

var (
	Version   string
	Revision  string
	ByteCount int64
	Filename  string
)

// https://golang.org/doc/effective_go.html#constants
const (
	_        = iota // ignore first value by assigning to blank identifier
	KB int64 = 1 << (10 * iota)
	MB
	GB
)

func printVersion() {
	fmt.Fprintln(os.Stdout, "Version:", Version)
	fmt.Fprintln(os.Stdout, "Revision:", Revision)
}

func byteCount(b, k, m, g int64) (int64, string) {
	if b > 0 {
		return int64(b), fmt.Sprintf("%v", b)
	}
	if k > 0 {
		return k * KB, fmt.Sprintf("%vKiB", k)
	}
	if m > 0 {
		return m * MB, fmt.Sprintf("%vMiB", m)
	}
	if g > 0 {
		return g * GB, fmt.Sprintf("%vGiB", m)
	}
	return 1 * MB, "1MiB"
}

func filename(args []string, byteString string) string {
	if len(args) == 0 {
		return byteString + ".dat"
	}
	return args[0]
}

func init() {
	var b, k, m, g int64
	var version bool
	flag.Int64Var(&b, "b", 0, "Byte")
	flag.Int64Var(&k, "k", 0, "KiB (kibibyte)")
	flag.Int64Var(&m, "m", 0, "MiB (mebibyte)")
	flag.Int64Var(&g, "g", 0, "GiB (gibibyte)")
	flag.BoolVar(&version, "v", false, "Print version.")
	flag.Parse()

	if version {
		printVersion()
		os.Exit(ExitCodeOK)
	}

	var byteString string
	ByteCount, byteString = byteCount(b, k, m, g)
	Filename = filename(flag.Args(), byteString)
}

func main() {
	fmt.Fprintln(os.Stdout, "filename:", Filename)
	fmt.Fprintln(os.Stdout, "byte    :", ByteCount)

	file, err := os.Create(Filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(ExitCodeError)
	}
	defer file.Close()

	mw := NewMonitorWriter(os.Stdout, ByteCount)
	w := io.MultiWriter(file, mw)

	if _, err := io.CopyN(w, rand.Reader, ByteCount); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(ExitCodeError)
	}

	fmt.Fprintf(os.Stdout, "\ndone\n")
	os.Exit(ExitCodeOK)
}
