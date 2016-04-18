package main
import (
    "fmt"
)

func main() {
    a := "5.0 30.5"
    var v1, v2 float64
    fmt.Sscanf(a, "%f %f", &v1, &v2)
    fmt.Println(v1, v2)
}
