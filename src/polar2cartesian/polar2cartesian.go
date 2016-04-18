package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "runtime"
)

type polar struct {
    radius    float64
    angle     float64
}

type cartesian struct {
    x    float64
    y    float64
}

var prompt = "Enter a radius and an angle (in degress), e.g., 12.5 90," + "or %s to quit."

func init() {
    if runtime.GOOS == "windows" {
        prompt = fmt.Sprintf(prompt, "Ctrl + Z, Enter")
    } else {
        prompt = fmt.Sprintf(prompt, "Ctrl + D")
    }
}

func main() {
    questions := make(chan polar)
    defer close(questions)
    answers := createSolver(questions)
    defer close(answers)
    interact(questions, answers)
}

func createSolver(questions chan polar) chan cartesian {
    answers := make(chan cartesian)
    go func() {
        for {
            polarCoord := <-questions
            angle := polarCoord.angle * math.Pi / 180.0
            x := polarCoord.radius * math.Cos(angle)
            y := polarCoord.radius * math.Sin(angle)
            answers <- cartesian{x, y}
        }
    }()
    return answers
}

const result = "Polar radius = %.02f angle = %.02f ->Cartesian x = %.02f y = %.02f\n"
func interact(questions chan polar, answers chan cartesian) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println(prompt)
    for {
        fmt.Println("Radius and angle: ")
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        fmt.Println("here")
        var radius, angle float64
        fmt.Println(line)
        if _, err := fmt.Sscanf(string(line), "%f %f", &radius, &angle); err != nil {
            fmt.Println(os.Stderr, "invalid input", err)
            continue
        }
        questions <-polar{radius, angle}
        coord := <-answers
        fmt.Printf(result, radius, angle, coord.x, coord.y)
    }
    fmt.Println()
}
