package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int16 }
func genFn0(p0 uint, p1 bool, p2 uint32) uint32 {
    for fi7 := 0; fi7 < 1; fi7++ {
        p2 = p2
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 float64 = 1e10
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(p2, p1, p1)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    v1 := []int{32768}
    v2 := make(map[int]string)
    var v3 uint8 = 255
    v4 := make(map[int]uint32)
    var v5 int = 0
    var v6 uint64 = 1
    var v11 uint8 = (v3 % ((1) | 1))
    if (true || true) {
        var v12 int64 = 0
        if (!(42 >= 42)) {
            v12 -= (v12 / (((v12 >> ((941) & 63))) | 1))
        } else {
            var v13 uint64 = v6
            _ = v13
        }
        _ = v12
    }
    fmt.Println(v3)
    for fi14 := 0; fi14 < 14; fi14++ {
        v3 += (v3 + v3)
        _ = fi14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS1{F0: 42, F1: 1, F2: -32768}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi15
        _ = v16
    }
    var v17 uint16 = 386
    if false {
        if (1 > 991) {
            v17 += 255
            v1 = append(v1, v5)
        }
    }
    var v18 uint32 = 561
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
