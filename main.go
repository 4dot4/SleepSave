package main

import (
	"fmt"
	"time"

	gui "github.com/gen2brain/raylib-go/raygui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 800, "SleepSave")
	var timer time.Duration
	for !rl.WindowShouldClose() {
		var clock = time.Now()
		var endtime = clock.Add(timer)
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		if gui.LabelButton(rl.Rectangle{300, 400, 20, 20}, "+ 1 hour") {
			timer += time.Hour
		}
		if gui.LabelButton(rl.Rectangle{300, 600, 20, 20}, "- 1 hour") {
			timer -= time.Hour
		}
		if gui.LabelButton(rl.Rectangle{400, 400, 20, 20}, "+ 30 minutes") {
			timer += 30 * time.Minute
		}
		if gui.LabelButton(rl.Rectangle{400, 600, 20, 20}, "- 30 minutes") {
			timer -= 30 * time.Minute
		}
		rl.DrawText("it's "+clock.Format(time.Kitchen), 250, 100, 60, rl.Blue)
		rl.DrawText("End Time "+endtime.Format(time.Kitchen), 100, 700, 60, rl.Red)

		format := fmt.Sprintf("%v", timer)
		rl.DrawText(format, 300, 500, 60, rl.Green)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}