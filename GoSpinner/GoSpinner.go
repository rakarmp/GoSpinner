package gospinner

import (
	"fmt"
	"time"
)

// spinner struct
type LoadingSpinner struct {
	frames map[string][]string
}

// new instance in LoadingSpinner
func NewLoadingSpinner() *LoadingSpinner {
	frames := make(map[string][]string)

	// You can add new frames in here
	frames["default"] = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	frames["arrow"] = []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}
	frames["circle"] = []string{"◐", "◓", "◑", "◒"}
	frames["star"] = []string{"✶", "✸", "✹", "✺", "✹", "✷"}

	return &LoadingSpinner{frames}
}

/*
Start loading spinner
Select variant handler
*/

func (ls *LoadingSpinner) Start(variant string, duration time.Duration) {
	frames, ok := ls.frames[variant]
	if !ok {
		frames = ls.frames["default"]
	}

	ticker := time.NewTicker(duration)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				done <- true
				return
			default:
				for _, frame := range frames {
					fmt.Printf("\033[2K\r%s Loading...", frame)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}()

	<-done
	ticker.Stop()
	fmt.Print("\033[2K\r")
}
