package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/olandr/wharf/pkg"
)

func flags() *regexp.Regexp {
	var pattern string
	flag.StringVar(&pattern, "pattern", "", "regex pattern to search for")
	flag.Parse()
	if pattern != "" {
		return regexp.MustCompile(pattern)
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pattern := flags()
	rorc := pkg.NewRorColuriser()
	for scanner.Scan() {
		fmt.Println(rorc.ApplyStyle(scanner.Text(), pattern))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
