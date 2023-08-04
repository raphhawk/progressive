package progressive

import (
	"fmt"
	"log"
	"time"
)

// ColorMap contains set of colors for the
// progress bar visualization
var ColorMap = map[string]string{
	"Reset":  "\033[0m",
	"Red":    "\033[31m",
	"Green":  "\033[32m",
	"Yellow": "\033[33m",
	"Blue":   "\033[34m",
	"Purple": "\033[35m",
	"Cyan":   "\033[36m",
	"Gray":   "\033[37m",
	"White":  "\033[97m",
}

// ValidProgress: an interface for
// validation of progress data comming from external src.
type ValidProgress interface {
	CloseProgress()
	GetProgress() (int, string, string)
	PassProgress(percent int, status string)
}

// ProgressChannel: a structure that holds information about
// respective progress bar
type ProgressChannel struct {
	Progress        chan int
	ProgressPercent int
	ProgressStatus  string
	ProgressName    string
}

// GetProgress: gets progress, name, status of the process
func (p *ProgressChannel) GetProgress() (int, string, string) {
	return <-p.Progress, p.ProgressName, p.ProgressStatus + "               "
}

// PassProgress: pass pre calculated progress and status of the process
// to the ProgressChannel
func (p *ProgressChannel) PassProgress(percent int, status string) {
	if percent > 100 || percent < 0 {
		log.Fatalf("Invalid progress: Exceeding bounds")
	} else if percent < p.ProgressPercent {
		log.Fatalf("Invalid progress: Decreasing Progress")
	}
	p.ProgressPercent = percent
	if len(status) > 0 {
		p.ProgressStatus = status
	}
	p.Progress <- percent
	time.Sleep(1 * time.Second)
}

// CloseProgress: closes a ProgressChannel explicitly
func (p *ProgressChannel) CloseProgress() { close(p.Progress) }

// ProgressBar: Displays a concurrent progress bar on
// the cli/terminal
func ProgressBar(
	progressLength int,
	updateLength ValidProgress,
	fillColor string,
	barColor string,
) {
	filler, toFill := "█", "░"
	for {
		progress1, name, status := updateLength.GetProgress()
		progress := (progress1 * progressLength) / 100
		fmt.Print(" \r")
		for j := 0; j < progress; j++ {
			fmt.Print(fillColor + filler)
		}
		for k := 0; k < (progressLength - progress); k++ {
			fmt.Print(barColor + toFill)
		}
		fmt.Printf(" %v%% %v %v", (progress * 100 / progressLength), name, status)
		fmt.Print("\r ")
		if progress1 == 100 {
			break
		}
	}
	fmt.Println()
}
