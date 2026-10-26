package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int16 }
type GenS2 struct { F0 bool }
func genFn0(p0 int64, p1 string, p2 int64) uint {
    var v5 int64 = p2
    v6 := GenS1{F0: 287, F1: (v5 ^ p2), F2: -32768}
    fmt.Println(p0, p2, p2)
    var v7 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 598
}

func main() {
    var v1 float64 = -2.25
    var v2 uint32 = 42
    v3 := make(map[string]string)
    v4 := []uint32{65535, v2, v2}
    var v8 uint16 = 120
    for fi9 := 0; fi9 < 18; fi9++ {
        v2 += v2
        _ = len(v3)
        _ = fi9
    }
    switch 128 {
    case 32767:
        fmt.Println(v2)
    case -468:
        var v10 int8 = -1
        _ = v10
    case 0:
        v11 := GenS2{F0: (32767 <= 1000)}
        _ = v11
    default:
        for fi12 := 0; fi12 < 7; fi12++ {
            switch (fi12 / ((127) | 1)) {
            case 2147483647:
                fmt.Println(v1)
            }
            var v13 uint16 = v8
            for fi14 := 0; fi14 < 3; fi14++ {
                fmt.Println(fi14, v2)
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := v3[("\n\t" + "a")]
        _ = fi15
        _ = v16
    }
    var v17 bool = false
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 uint64 = 99
        v20 := GenS2{F0: v17}
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println(v8)
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 float32 = -2.25
        if v17 {
            v8 = (^(v8))
        } else {
            v2 = 61
        }
        if (!(v1 >= float64(v22))) {
            switch ("a" + "") {
            case "\n\t":
                v22 += (v22 / (v22 * (v22 + float32(v1))))
            case "foo bar":
                for fi23 := 0; fi23 < 8; fi23++ {
                    v24 := GenS1{F0: 255, F1: int64(fi21), F2: 798}
                    var v25 int32 = -83
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            default:
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 float64 = (v1 / v1)
                    var v28 int64 = 554
                    fmt.Println(v8, v28)
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            var v29 int16 = 32767
            for fi30 := 0; fi30 < 10; fi30++ {
                v31 := GenS2{F0: (v1 == v1)}
                fmt.Println(v1, fi30, v2)
                _ = fi30
                _ = v31
            }
            _ = v29
        } else {
            if len(v4) > 0 {
                v4[0] = (v2 << ((65535) & 63))
            }
            var v32 uint = 65535
            _ = v32
        }
        _ = fi21
        _ = v22
    }
    var v33 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v8
}
