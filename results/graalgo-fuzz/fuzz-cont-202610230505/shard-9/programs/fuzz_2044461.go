package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 string }

func main() {
    var v1 int16 = 42
    v2 := []int{63, 32768}
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    var v5 float32 = 1.5
    var v6 uint32 = 42
    _ = len(v3)
    v6 = v6
    v5 = float32(1e10)
    v7 := GenS1{F0: 127, F1: 40, F2: "\n\t"}
    fmt.Println(v6)
    v1 -= 42
    v8 := GenS1{F0: 198, F1: 0, F2: "hello"}
    fmt.Println(v1)
    _ = len(v2)
    if ((v5 - ((v5 * 0.0) - v5)) > v5) {
        fmt.Println(v1)
    } else {
        for fi9 := 0; fi9 < 11; fi9++ {
            v5 = v5
            _ = fi9
        }
        v1 -= v1
    }
    var v10 int64 = -254
    v11 := GenS1{F0: 255, F1: 100, F2: "hello"}
    fmt.Println(v6, v1, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
}
