package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint32 = 823
    var v2 int32 = -2147483648
    v3 := make(map[int]string)
    if (v2 == -1) {
        delete(v3, 127)
    } else {
        var v4 float64 = float64(float32(float64(1.5)))
        _ = v4
    }
    fmt.Println(v2)
    v5 := GenS1{F0: 220}
    var v6 string = "foo bar"
    v7 := GenS2{F0: 930, F1: 1}
    v8 := GenS3{F0: 0}
    v1 = (v1 / ((v1) | 1))
    var v9 float32 = (1e10 / 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
