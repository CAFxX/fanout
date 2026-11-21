package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
func genFn0(p0 int, p1 int16, p2 int32) int16 {
    var v7 int8 = 42
    var v8 float32 = (((3.14159 * float32(3.14159)) * ((0.0 + 1.5) * 1e10)) / float32((-2.25 / (1e10 + (float64(0.1) * 3.14159)))))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    var v3 uint64 = 42
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    var v6 uint32 = 0
    var v9 int = -1
    if true {
        v3 = 0
    } else {
        var v10 int8 = 1
        fmt.Println(v9, v3)
        _ = v10
    }
    v11 := GenS1{F0: 32767, F1: 1}
    v12 := GenS1{F0: 27, F1: 1}
    v3 += (v3 >> ((2) & 63))
    var v13 uint8 = uint8(v6)
    var v14 uint8 = 1
    v13 = (v14 | v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
