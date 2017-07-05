package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/go-gl/gl/v2.1/gl" // OR: github.com/go-gl/gl/v4.1-core/gl
	"github.com/go-gl/glfw/v3.2/glfw"

	"./object"
	//"./play"
	"math/rand"
)

const (
	vertexShaderSource = `
		#version 330
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
		#version 330
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(1, 1, 1, 1.0);
		}
	` + "\x00"

	fps = 60
)

func main() {
	runtime.LockOSThread()
	memory_stats := &runtime.MemStats{}

	window := initGlfw()
	defer glfw.Terminate()
	program := initOpenGL()

	game_board := object.Board{Program: program, Window: window}
	game_board.MakeCells()
	game_board.MakeEnemyTanks(1)
	game_board.MakePlayerTanks(1)
	//game_board.DestroyTank(0)

	game_board.PlayerTanks[0].RotateRight()
	game_board.PlayerTanks[0].Fire()
	game_board.PlayerTanks[0].Fire()
	game_board.PlayerTanks[0].Fire()
	game_board.PlayerTanks[0].Fire()
	game_board.PlayerTanks[0].Fire()
	game_board.PlayerTanks[0].Fire()

	game_board.EnemyTanks[0].MoveToPosition(50, 2)

	for !window.ShouldClose() {
		t := time.Now()

		for i := 0; i < len(game_board.EnemyTanks); i++ {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			if r1.Float32() > 0.45 {
				game_board.EnemyTanks[i].RotateRight()
			} else {
				game_board.EnemyTanks[i].RotateLeft()
			}
			if r1.Float32() > 0.7 {
				//game_board.EnemyTanks[i].Fire()
			}

			//game_board.EnemyTanks[i].MoveForward()
			//tank.MoveToPosition(50,50)
			//log.Printf("Enemy 0 bullet count %d", len(game_board.EnemyTanks[0].Bullet))
		}

		runtime.ReadMemStats(memory_stats)
		//fmt.Println("Time;Allocated;Total Allocated; System Memory;Num Gc;Heap Allocated;Heap System;Heap Objects;Heap Released;\n")
		//fmt.Printf("%s;%d;%d;%d;%d;%d;%d;%d;%d;\n", time.Now(), memory_stats.Alloc, memory_stats.TotalAlloc, memory_stats.Sys, memory_stats.NumGC, memory_stats.HeapAlloc, memory_stats.HeapSys, memory_stats.HeapObjects, memory_stats.HeapReleased)
		game_board.Draw()

		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(object.WindowWidth, object.WindowHeight, "Just for fun", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

// initOpenGL initializes OpenGL and returns an initialized program.
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		logging := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logging))

		return 0, fmt.Errorf("failed to compile %v: %v", source, logging)
	}

	return shader, nil
}
