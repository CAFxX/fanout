package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]string)
    var v3 uint8 = 127
    var v4 uint16 = 32767
    v4 -= v4
    var v5 string = ("a" + "foo bar")
    var v6 uint64 = 199
    var v7 float64 = 1e10
    var v8 float32 = -2.25
    v9 := GenS1{F0: true, F1: int(v3)}
    v7 -= v7
    v5 = (v5 + v5)
    v10 := GenS2{F0: (v4 >> ((996) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
