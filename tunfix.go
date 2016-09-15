package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//
// problem: if you resume a screen session, your ssh
// tunnel parameters have changed.
//
// solution:
// a) record the new parameters in ~/.ssh/last.tun
//    on each new login. One time setup, see below.
// b) once inside the resumed screen session,
//    run tunfix
// c) source the ~/.ssh/.tuno file that has
//    just been written, to update your shell.
//
// One time setup:
//
// add to .bash_profile this line:
// env | grep SSH > $HOME/.ssh/last.tun
//
// license: MIT
//
func main() {
	home := os.Getenv("HOME")
	if home == "" {
		panic("could not read $HOME from env")
	}
	f, err := ioutil.ReadFile(home + "/.ssh/last.tun")
	panicOn(err)
	ofn := home + "/.ssh/.tuno"
	o, err := os.OpenFile(ofn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0774)
	panicOn(err)

	lines := strings.Split(string(f), "\n")
	for _, line := range lines {
		splt := strings.Split(line, "=")
		if len(splt) == 2 {
			fmt.Fprintf(o, "export %s=\"%s\"\n", splt[0], splt[1])
		}
	}
	err = o.Close()
	panicOn(err)
	fmt.Printf("cat %s && source %s # <- do this\n", ofn, ofn)
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}
