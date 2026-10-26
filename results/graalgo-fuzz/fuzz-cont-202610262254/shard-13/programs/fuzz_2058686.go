package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
func genFn0(p0 int32, p1 int16, p2 uint) uint8 {
    var v6 int16 = 1000
    if (!true) {
        p0 = p0
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 224
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]bool)
    var v3 float32 = (-2.25 * float32(0.0))
    v4 := []uint32{v1, v1, 1}
    var v5 float32 = float32((1e10 - 1e10))
    var v7 uint8 = 0
    v3 = v5
    var v8 uint8 = v7
    var v9 uint16 = 103
    v10 := GenS1{F0: v1, F1: 428}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 uint8 = (v8 << ((1) & 63))
            var v13 int8 = 1
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 int32 = 1
        _ = v14
    }
    if false {
        for fi15 := 0; fi15 < 8; fi15++ {
            v1 += v1
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 uint64 = 42
                v7 += (^(v8))
                var v18 int8 = 0
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v1, fi16)
                _ = fi16
                _ = v17
                _ = v18
            }
            for fi19 := 0; fi19 < 9; fi19++ {
                var v20 string = "foo bar"
                v2[(v1 << ((186) & 63))] = true
                v4 = append(v4, 65535)
                _ = fi19
                _ = v20
            }
            _ = fi15
        }
        for fi21 := 0; fi21 < 19; fi21++ {
            if len(v4) > 0 {
                v4[0] = v1
            }
            _ = fi21
        }
        v1 += 65535
    }
    v4 = append(v4, 877)
    var v22 int32 = 1
    for fi23 := 0; fi23 < 5; fi23++ {
        if len(v4) > 0 {
            v4[0] = (v1 - (v1 & (^(v1))))
        }
        _ = fi23
    }
    switch 1 {
    case 0:
        v24 := GenS1{F0: 2147483647, F1: 0}
        _ = v24
    case -957:
        v3 = (float32(((1e10 / 1.0) * (-2.25 - ((((0.1 / 1.0) * (0.1 * 1e10)) - 1.5) * 0.1)))) - 1e10)
    default:
        v1 -= (v1 - 2)
    }
    if true {
        fmt.Println(v9, v7, v3)
        var v25 float64 = 0.1
        var v26 int16 = 0
        _ = v25
        _ = v26
    }
    v7 += v8
    var v27 int64 = 1
    for fi28 := 0; fi28 < 6; fi28++ {
        var v29 int64 = (v27 | (v27 % ((653) | 1)))
        v4 = append(v4, v1)
        v2[(v1 * (v1 >> ((2) & 63)))] = false
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
