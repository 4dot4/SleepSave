package main

import (
	"fmt"
	"os/exec"
	"time"

	gui "github.com/gen2brain/raylib-go/raygui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 800, "SleepSave")
	var timer time.Duration
	var setTimer = true
	var endtime time.Time
	var x, y int32 = 0, 0

	// var FrameCounter uint
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyA) {
			x -= 5
			fmt.Println("x: ", x)
		}
		if rl.IsKeyDown(rl.KeyD) {
			x += 5
			fmt.Println("x: ", x)
		}
		if rl.IsKeyDown(rl.KeyW) {
			y -= 5
			fmt.Println("y: ", y)
		}
		if rl.IsKeyDown(rl.KeyS) {
			y += 5
			fmt.Println("y: ", y)
		}
		var clock = time.Now()
		endtime = clock.Add(timer)
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		if setTimer {
			if gui.Button(rl.Rectangle{350, 580, 100, 50}, "set timer") {
				setTimer = false
				fmt.Println(setTimer, "dentro do butao")
				go func(duration *time.Duration) {
					for *duration != 0 {
						time.Sleep(time.Second)
						*duration -= time.Second
					}

					if err := exec.Command("cmd", "/C", "shutdown", "/s").Run(); err != nil {
						fmt.Println("Failed to initiate shutdown:", err)
					}
				}(&timer)
			}
		}
		if gui.LabelButton(rl.Rectangle{200, 320, 20, 20}, "+ 1 hour") {
			timer += time.Hour
		}
		if gui.LabelButton(rl.Rectangle{510, 320, 20, 20}, "+ 30 minutes") {
			timer += 30 * time.Minute
		}

		if gui.LabelButton(rl.Rectangle{200, 520, 20, 20}, "- 1 hour") {
			if timer-time.Hour >= 0 {

				timer -= time.Hour
			}
		}
		if gui.LabelButton(rl.Rectangle{510, 520, 20, 20}, "- 30 minutes") {
			if timer-30*time.Minute >= 0 {

				timer -= 30 * time.Minute
			}
		}
		rl.DrawText("it's "+clock.Format(time.Kitchen), 250, 100, 60, rl.Blue)
		rl.DrawText("End Time "+endtime.Format(time.Kitchen), 165, 710, 60, rl.Red)

		format := fmt.Sprintf("%v", timer)
		rl.DrawText(format, 300, 400, 60, rl.Green)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}