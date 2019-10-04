package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alecrabbit/go-cli-spinner"
	"github.com/alecrabbit/go-cli-spinner/color"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	messages := []string{
		"Starting",
		"Initializing",
		"Gathering data",
		"Checking data",
		"Checking weather",
		"Processing",
		"Processing",
		"Processing",
		"Processing",
		"Processing",
		"Processing",
		"Processing",
		"Processing",
		"Almost there",
		"Be patient",
	}

	s, _ := spinner.New(
		spinner.Interval(250),
		spinner.ColorLevel(color.TColor256),
		spinner.ProgressFormat("[%4s]"), // [  7%]
		spinner.MessageFormat("(%s)"),   // (message)
	)
	// s.Colors(spinner.Set{Set256: spinner.ColorSets[spinner.C256Rainbow]})
	s.FinalMessage = "Done!\n"
	// s.HideCursor = false
	s.Reversed = true
	// s.prefix = " " // spinner prefix

	// Start spinner
	s.Start()
	// for _, m := range messages {
	l := len(messages)
	for i, m := range messages {
		// Doing some work 1
		time.Sleep(duration())
		// Printing execution message
		{
			s.Erase() // optional if you're absolutely sure that your messages are longer
			fmt.Println(m)
			fmt.Print("..................................") // string to show that spinner can be used in inline mode
			s.Current()                                     // Write current frame to output(optional - for smooth amination)
		}
		// Simulating spinner message
		if rand.Intn(16) > 7 {
			s.Message("") // Sometimes there are no messages
		} else {
			s.Message(fmt.Sprintf("Message at %s", time.Now().Format("15:04:05")))
		}
		// Doing some work 2
		time.Sleep(duration())
		// Simulating spinner progress
		s.Progress(float32(i) / float32(l)) // float32 0..1
	}
	time.Sleep(1 * time.Second)
	// Stop spinner
	s.Stop()
	time.Sleep(1 * time.Second)
}

func duration() time.Duration {
	return (200 + time.Duration(rand.Intn(1600))) * time.Millisecond
}
