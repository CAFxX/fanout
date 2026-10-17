package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int64 }
func genFn0(p0 uint64, p1 bool) int16 {
    fmt.Println(p0, p1)
    v5 := GenS1{F0: 0, F1: 0, F2: -925}
    _ = p0
    _ = p1
    _ = v5
    return -963
}

func main() {
    var v1 string = ""
    var v2 uint32 = 0
    var v3 int = 854
    v4 := make(map[uint32]uint32)
    v2 = 847
    v6 := GenS1{F0: 24, F1: (v2 ^ (v2 / ((v2) | 1))), F2: -9223372036854775808}
    v2 += v2
    var v7 uint32 = (v2 % ((42) | 1))
    var v8 uint32 = (v2 >> ((1) & 63))
    var v9 int16 = -1
    v2 -= 0
    if ((v9 * (v9 - v9)) <= (v9 / ((v9) | 1))) {
        fmt.Println(v1, v8)
    } else {
        fmt.Println(v7, v9)
    }
    fmt.Println(v1, v2)
    v4[v8] = v2
    var v10 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
