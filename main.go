package main

import (
	//"fmt"
	"fmt"
	gui "github.com/gen2brain/raylib-go/raygui"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Timer struct {
	Hours   int
	Minutes int
}

func main() {
	rl.InitWindow(800, 800, "SleepSave")
	var timer Timer
	for !rl.WindowShouldClose() {
		var clock = time.Now()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawText(clock.Format("03:04 am"), 250, 100, 60, rl.Blue)
		rl.EndDrawing()
	}
}