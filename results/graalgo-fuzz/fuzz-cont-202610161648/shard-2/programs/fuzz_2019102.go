package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int }
type GenS2 struct { F0 int16; F1 uint; F2 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 4294967295
    v3 := make(map[uint32]uint32)
    v4 := GenS1{F0: 127, F1: -1, F2: 9223372036854775807}
    var v5 float64 = (0.1 * 0.0)
    v5 = v5
    if false {
        v6 := GenS2{F0: 42, F1: 781, F2: -9223372036854775808}
        v5 = (v5 + v5)
        v5 += ((1e10 - v5) / (v5 - ((v5 * v5) * (((1e10 + float64(((0.1 / -2.25) / (3.14159 / 3.14159)))) / float64((0.1 / 1e10))) * 1.5))))
        _ = v6
    }
    v7 := []int{-9223372036854775808}
    fmt.Println(v2)
    if (((((!true) && true) || (!(!false))) || true) || (!(((!(1 != 1)) || true) || ((-126 == 46) || false)))) {
        v5 += (v5 + v5)
        v2 -= (v2 >> ((42) & 63))
    }
    var v8 bool = false
    var v9 int32 = -74
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
