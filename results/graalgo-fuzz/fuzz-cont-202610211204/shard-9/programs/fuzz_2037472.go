package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint16; F1 int64; F2 int16 }

func main() {
    v1 := []uint32{2147483647}
    var v2 bool = false
    var v3 float32 = 0.0
    var v4 float32 = (float32(0.1) - (v3 - v3))
    v5 := make(map[string]int)
    v5["a"] = 0
    var v6 float32 = float32((1.5 - -2.25))
    v7 := GenS1{F0: 42, F1: ("" + ("a" + "hello"))}
    v8 := GenS2{F0: 42, F1: 868, F2: 32767}
    v3 = v4
    var v9 int64 = -160
    v9 += v9
    v10 := GenS1{F0: 2, F1: (((("" + ("" + "")) + ("a" + "\n\t")) + "\n\t") + "")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
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
