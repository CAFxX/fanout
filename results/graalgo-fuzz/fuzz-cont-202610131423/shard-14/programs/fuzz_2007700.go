package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int64, p1 bool, p2 int) uint {
    v5 := GenS1{F0: 4294967295, F1: 1, F2: 255}
    p1 = (((p2 & (p2 * 1)) <= (p2 - 128)) || (!p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = float32((((((0.0 - ((float64(0.0) / 1.0) + 1e10)) + float64((1e10 + -2.25))) / ((1e10 - 1.5) / -2.25)) + 1e10) * 1.5))
    v3 := make(map[string]bool)
    v4 := make(map[string]string)
    if (!true) {
        v2 = v2
        v6 := GenS2{F0: -274}
        var v7 int32 = 2147483647
        _ = v6
        _ = v7
    } else {
        v2 = (float32(float64(v2)) * v2)
    }
    v8 := GenS1{F0: 985, F1: 158, F2: 0}
    var v9 int16 = -32768
    var v10 bool = ("foo bar" > "a")
    var v11 float64 = (-2.25 / 1e10)
    v12 := GenS2{F0: 42}
    var v13 uint16 = 65535
    v11 = float64(v2)
    fmt.Println(v2)
    fmt.Println(v13, v11, v10)
    v14 := GenS2{F0: 42}
    v9 -= (v9 / ((v9) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
