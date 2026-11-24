package main

import (
	"ani-track/src/tui"
	"log"
)

func main() {

	app := tui.SetupUI()
	if err := app.Run(); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
