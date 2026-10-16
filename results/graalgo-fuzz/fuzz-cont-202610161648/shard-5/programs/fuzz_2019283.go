package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0(p0 uint) uint16 {
    var v6 int64 = 42
    if (-1 < 849) {
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS1{F0: true}
            _ = fi7
            _ = v8
        }
        var v9 uint32 = 42
        p0 += (^(p0))
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 bool = true
        v12 := []int{}
        p0 += (p0 - 984)
        _ = fi10
        _ = v11
        _ = v12
    }
    p0 = (p0 ^ p0)
    _ = p0
    _ = v6
    return 359
}
func genFn1(p0 uint32, p1 uint, p2 int) int16 {
    p2 -= p2
    var v13 int = p2
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := []int{2147483647, v13}
        var v16 int = (fi14 << ((p1) & 63))
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -452
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 1
    v3 := []int{9223372036854775807, 0}
    var v4 uint16 = 32767
    var v5 int8 = -128
    v4 = (v4 * (v4 / (((^(v4))) | 1)))
    var v17 uint32 = 459
    v4 += v4
    if true {
        fmt.Println(v5)
        if true {
            delete(v1, 2)
        } else {
            v5 = v5
            for fi18 := 0; fi18 < 15; fi18++ {
                _ = len(v3)
                var v19 uint16 = (v4 % ((v4) | 1))
                v20 := GenS2{F0: 1, F1: v5}
                _ = fi18
                _ = v19
                _ = v20
            }
            var v21 uint8 = 127
            _ = v21
        }
    } else {
        var v22 float32 = (0.0 * 3.14159)
        fmt.Println(v2, v2)
        _ = v22
    }
    v17 = 65535
    var v23 uint64 = 0
    for fi24 := 0; fi24 < 20; fi24++ {
        v2 += v2
        var v25 uint16 = genFn0(4294967295)
        _ = fi24
        _ = v25
    }
    v17 -= uint32(v2)
    var v26 int8 = v5
    var v27 int16 = 696
    if false {
        var v28 int16 = (v27 ^ v27)
        _ = v28
    } else {
        if (true || ((v26 << ((65535) & 63)) < v26)) {
            v29 := GenS2{F0: (-(v26)), F1: 56}
            _ = v29
        }
        v17 += v17
    }
    var v30 uint16 = v4
    var v31 int8 = (v5 | (-(v26)))
    v3 = append(v3, 32767)
    if len(v3) > 0 {
        v3[0] = -549
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
