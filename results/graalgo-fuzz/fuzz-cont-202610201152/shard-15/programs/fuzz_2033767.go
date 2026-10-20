package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint8; F2 string }

func main() {
    var v1 int = -781
    v2 := make(map[uint32]string)
    var v3 int = 1
    var v4 uint64 = 0
    var v5 string = ("hello" + "\n\t")
    var v6 uint = 940
    v7 := GenS1{F0: (v6 / ((2) | 1))}
    var v8 int8 = -128
    var v9 int8 = (v8 << ((v6) & 63))
    v8 = v8
    v10 := GenS2{F0: (v5 + "\n\t"), F1: 198, F2: "hello"}
    v8 = 18
    var v11 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
