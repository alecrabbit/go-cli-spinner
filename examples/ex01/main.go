package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/alecrabbit/go-cli-spinner"
	"github.com/alecrabbit/go-cli-spinner/color"
)

const dots = "...................."

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	messages := map[int]string{
		0:  "Starting",
		3:  "Initializing",
		6:  "Gathering data",
		9:  "Processing",
		16: "Processing",
		25: "Processing",
		44: "Processing",
		60: "Processing",
		79: "Processing",
		82: "Still processing",
		90: "Be patient",
		95: "Almost there",
	}

	s, err := spinner.New(
		// Override default refresh interval, each CharSet has it's own recommended refresh interval
		spinner.Interval(120),
		// Override default color level support, default: TColor256
		spinner.ColorLevel(color.TColor256),
		spinner.ProgressFormat("[%4s]"),     // [  7%]
		spinner.MessageFormat("(%s)"),       // (message)
		spinner.Format("-%s -"),       // -⠏ -
		spinner.Prefix("\x1b[38;5;161mprefix\x1b[0m"),
		spinner.FinalMessage("\x1b[38;5;34mDone!\x1b[0m\n"),
		spinner.Reverse(),
		// spinner.HideCursor(false),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Print(dots)

	// Start spinner
	s.Start()
	for i := 0; i <= 100; i++ {
		// Doing some work 1
		time.Sleep(duration())
		if m, ok := messages[i]; ok {
			// Printing execution message
			s.Erase()       // optional if you're absolutely sure that your messages are longer
			fmt.Println(m)  //
			fmt.Print(dots) // string to show that spinner can be used in inline mode
			s.Current()     // Write current frame to output(optional - for smooth animation)
		}
		// Simulating spinner message
		if i > 10 && i < 20 || i > 40 && i < 60 { // Sometimes there are no messages
			s.Message("")
		} else {
			s.Message(fmt.Sprintf("Status message at %s", time.Now().Format("15:04:05")))
		}
		// Simulating spinner progress
		if i > 50 && i < 70 { // between 50% and 70%
			s.Progress(0) // hide progress indicator
		} else {
			s.Progress(float32(i) / float32(100)) // float32 0..1
		}
		// Doing some work 2
		time.Sleep(duration())
	}
	// Stop spinner
	s.Stop()
	time.Sleep(100 * time.Millisecond)
}

func duration() time.Duration {
	return (50 + time.Duration(rand.Intn(50))) * time.Millisecond
}
