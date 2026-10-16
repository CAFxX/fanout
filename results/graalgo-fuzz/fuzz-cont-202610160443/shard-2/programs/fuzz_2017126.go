package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint32 }

func main() {
    var v1 int16 = 42
    v2 := make(map[int]string)
    var v3 float64 = (3.14159 - (3.14159 / ((3.14159 - ((3.14159 - (0.1 / 0.1)) + -2.25)) / -2.25)))
    var v4 int = 42
    v1 -= (v1 % ((v1) | 1))
    v4 = v4
    v5 := GenS1{F0: 43, F1: 628}
    v6 := GenS2{F0: ((!false) && ((0.1 != 3.14159) && (1 != 42))), F1: 776}
    fmt.Println(v1)
    v3 -= 0.0
    v3 += v3
    fmt.Println(v3, v3, v3)
    v4 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
