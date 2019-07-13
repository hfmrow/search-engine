// miscFunc.go

/*
*	©2019 H.F.M. MIT license
 */

package main

import (
	"fmt"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	g "github.com/hfmrow/searchEngine/genLib"
)

// Handling arguments from command line.
func cmdLineParser(args []string) {
	if len(args) > 1 && g.ScanFile(args[1]).IsExists {
		mainOptions.LastDirectory = g.ScanFile(args[1]).Path
		return
	}
	name := g.SplitFilepath(g.SplitFilepath("").ExecFullName).Base
	fmt.Printf("%s %s %s %s\n%s\n%s\n\nUsage: %s \"PathToScan\"\n",
		mainOptions.AboutOptions.AppName,
		mainOptions.AboutOptions.AppVers,
		mainOptions.AboutOptions.YearCreat,
		mainOptions.AboutOptions.AppCreats,
		mainOptions.AboutOptions.LicenseShort,
		mainOptions.AboutOptions.Description,
		name)
}

// Measuring lapse (may be multiples) between operations
type Bench struct {
	lapse   []time.Time
	label   []string
	Results []string
}

func (b *Bench) Lapse(label ...string) {
	b.lapse = append(b.lapse, time.Now())
	if len(label) == 0 {
		label = append(label, fmt.Sprintf("%d", len(b.lapse)))
	}
	b.label = append(b.label, label...)
}

func (b *Bench) ResetAndStart(label ...string) {
	b.lapse = b.lapse[:0]
	b.label = b.label[:0]
	b.Results = b.Results[:0]
	b.Lapse(label...)
}
func (b *Bench) Stop() {
	b.Lapse("")
	lapseCount := len(b.lapse) - 1
	var calculateLapse = func(count int, start, stop time.Time) {
		diff := stop.Sub(start).Nanoseconds()
		min := (diff / 1000000000) / 60
		sec := diff/1000000000 - (min * 60)
		ms := (diff / 1000000) - (sec * 1000)
		b.Results = append(b.Results, fmt.Sprintf("%vm%vs%vms", min, sec, ms))
	}
	b.Results = b.Results[:0]
	if lapseCount > 1 {
		for idx := 0; idx < len(b.lapse)-1; idx++ {
			calculateLapse(idx, b.lapse[idx], b.lapse[idx+1])
		}
	}
	calculateLapse(lapseCount, b.lapse[0], b.lapse[lapseCount])
	b.label = b.label[:0]
	b.lapse = b.lapse[:0]
}

// Check: Display error messages in HR version with onClickJump enabled in
// my favourite Golang IDE editor. Return true if error exist.
func Check(err error, message ...string) (state bool) {
	remInside := regexp.MustCompile(`[\s\p{Zs}]{2,}`) //	to match 2 or more whitespace symbols inside a string
	var msgs string
	if err != nil {
		state = true
		if len(message) != 0 { // Make string with messages if exists
			for _, mess := range message {
				msgs += `[` + mess + `]`
			}
		}
		pc, file, line, ok := runtime.Caller(1) //	(pc uintptr, file string, line int, ok bool)
		if ok == false {                        // Remove "== false" if needed
			fName := runtime.FuncForPC(pc).Name()
			fmt.Printf("[%s][%s][File: %s][Func: %s][Line: %d]\n", msgs, err.Error(), file, fName, line)
		} else {
			stack := strings.Split(fmt.Sprintf("%s", debug.Stack()), "\n")
			for idx := 5; idx < len(stack)-1; idx = idx + 2 {
				//	To match 2 or more whitespace leading/ending/inside a string (include \t, \n)
				mess1 := strings.Join(strings.Fields(stack[idx]), " ")
				mess2 := strings.TrimSpace(remInside.ReplaceAllString(stack[idx+1], " "))
				fmt.Printf("%s[%s][%s]\n", msgs, err.Error(), strings.Join([]string{mess1, mess2}, "]["))
			}
		}
	}
	return state
}
