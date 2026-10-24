package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int16 }
func genFn0(p0 int) int16 {
    fmt.Println(p0)
    var v4 uint32 = 85
    var v5 float32 = 1.5
    v6 := GenS1{F0: uint16(v4)}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 uint32 = 65535
    var v2 int32 = -971
    v3 := make(map[uint32]int)
    fmt.Println(v1, v2)
    if (510 <= -343) {
        var v7 uint = 1
        fmt.Println(v7)
        _ = v7
    }
    v2 = (^(v2))
    var v8 float64 = float64(0.0)
    var v9 uint = 583
    fmt.Println(v8, v9)
    for fi10 := 0; fi10 < 3; fi10++ {
        v2 = v2
        v9 = v9
        _ = fi10
    }
    var v11 uint32 = (v1 | (v1 << ((v9) & 63)))
    var v12 uint64 = 634
    if true {
        v13 := GenS2{F0: 254, F1: 1000}
        _ = len(v3)
        if (false || (true && (((^(v9)) == 4294967295) || false))) {
            v1 -= v11
            var v14 int64 = -444
            _ = v14
        }
        _ = v13
    } else {
        v12 += (v12 / ((v12) | 1))
        v11 = 943
        v1 = (v1 >> ((v9) & 63))
    }
    var v15 int = -9223372036854775808
    v16 := []int{(v15 - v15), v15, v15, (v15 + v15)}
    v12 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
