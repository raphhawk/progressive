package progressive

import (
	"fmt"
	"log"
	"time"
)

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

/*Example

ch := make(chan int)

var p ValidProgress = &ProgressChannel{ch, 0, "", "Process 1"}
var pt ValidProgress = &ProgressChannel{ch, 0, "", "Process 2"}

go progressBar(25, p, colorMap["Green"], colorMap["Gray"])
go progressBar(25, pt, colorMap["Purple"], colorMap["Gray"])

p.passProgress(20, "Initializing Containers")
p.passProgress(50, "Creating Kube Clusters")

pt.passProgress(28, "Initializing Services")

p.passProgress(75, "Deploying Application")
pt.passProgress(34, "Monitorings Pod Health")
p.passProgress(100, "Operation Complete")

pt.passProgress(56, "Service #1 Deployed")
pt.passProgress(82, "Service #2 Deployed")
pt.passProgress(100, "Operation Complete")

pt.closeProgress()
*/

type ValidProgress interface {
	closeProgress()
	getProgress() (int, string, string)
	passProgress(percent int, status string)
}

type ProgressChannel struct {
	progress        chan int
	progressPercent int
	progressStatus  string
	progressName    string
}

func (p *ProgressChannel) getProgress() (int, string, string) {
	return <-p.progress, p.progressName, p.progressStatus + "               "
}
func (p *ProgressChannel) passProgress(percent int, status string) {
	if percent > 100 || percent < 0 {
		log.Fatalf("Invalid progress: Exceeding bounds")
	} else if percent < p.progressPercent {
		log.Fatalf("Invalid progress: Decreasing Progress")
	}
	p.progressPercent = percent
	if len(status) > 0 {
		p.progressStatus = status
	}
	p.progress <- percent
	time.Sleep(1 * time.Second)
}

func (p *ProgressChannel) closeProgress() { close(p.progress) }

func progressBar(
	progressLength int,
	updateLength ValidProgress,
	fillColor string,
	barColor string,
) {
	filler, toFill := "█", "░"
	for {
		progress1, name, status := updateLength.getProgress()
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
