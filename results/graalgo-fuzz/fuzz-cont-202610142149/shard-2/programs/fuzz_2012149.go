package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint }

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[uint32]int)
    v3 := make(map[string]bool)
    v4 := GenS1{F0: false, F1: "hello", F2: 79}
    switch 1 {
    case 293:
        var v5 float32 = (1e10 - 1.5)
        _ = v5
    case -743:
        var v6 uint16 = 0
        _ = v6
    }
    if (9223372036854775807 >= 1) {
        var v7 uint8 = 120
        _ = v7
    } else {
        v8 := GenS1{F0: (!(false && (!false))), F1: "a", F2: 42}
        _ = v8
    }
    var v9 uint8 = 93
    var v10 int64 = -9223372036854775808
    var v11 float64 = (0.0 + (0.0 + (-2.25 * float64(3.14159))))
    v11 = 0.1
    fmt.Println(v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
