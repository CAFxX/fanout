package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 uint64 = 514
    v4 := make(map[string]string)
    var v5 int = -437
    v6 := make(map[string]int)
    _ = len(v4)
    v5 -= 0
    v3 += v3
    var v7 int = (^(v5))
    v8 := GenS1{F0: true, F1: (v5 % (((^(v7))) | 1)), F2: 1}
    var v9 float32 = 0.0
    v3 = v3
    switch (("hello" + "") + "hello") {
    case "a":
        var v10 float32 = v9
        _ = v10
    case "\n\t":
        var v11 uint16 = 32767
        _ = v11
    case "hello":
        v3 -= (^(v3))
    default:
        v12 := GenS1{F0: true, F1: v7, F2: 583}
        _ = v12
    }
    switch v7 {
    case 42:
        var v13 float64 = (1.5 / 1.5)
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
