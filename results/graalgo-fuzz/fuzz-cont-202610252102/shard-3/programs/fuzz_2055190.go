package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int16, p1 uint16, p2 bool) uint {
    v5 := []int{2147483647}
    var v6 uint16 = (p1 / ((2) | 1))
    var v7 bool = p2
    v8 := GenS1{F0: 65535, F1: 100}
    var v9 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 65535
}
func genFn1(p0 uint, p1 string, p2 int32) uint32 {
    p0 = 23
    var v10 int32 = (p2 + p2)
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 string = p1
        switch "hello" {
        case "\n\t":
            if (!(9223372036854775807 >= -620)) {
                var v13 int64 = 597
                _ = v13
            } else {
                var v14 string = ((v12 + (v12 + ("" + p1))) + (p1 + v12))
                p2 += p2
                p0 += (p0 >> ((1) & 63))
                _ = v14
            }
        case "a":
            for fi15 := 0; fi15 < 10; fi15++ {
                v16 := GenS2{F0: -32768}
                if ((32767 != 0) && (v10 <= -1)) {
                    v12 = ("\n\t" + v12)
                    v17 := GenS2{F0: -32768}
                    fmt.Println(v12)
                    _ = v17
                }
                _ = fi15
                _ = v16
            }
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 497
}

func main() {
    var v1 float64 = float64((float32(0.1) + 0.0))
    v2 := make(map[int]uint32)
    var v3 uint16 = 1
    var v4 int64 = 840
    v18 := []int{-2147483648, 32768}
    _ = len(v18)
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    switch -9223372036854775808 {
    case 561:
        v3 = (v3 + v3)
    default:
        fmt.Println(v1)
    }
    var v19 uint32 = 42
    v3 += v3
    if true {
        var v20 bool = true
        var v21 uint64 = 1
        _ = v20
        _ = v21
    } else {
        v2[-1] = v19
    }
    v4 -= v4
    v3 += (v3 - 255)
    if false {
        v22 := GenS1{F0: 1, F1: 1}
        if len(v18) > 0 {
            v18[0] = 773
        }
        switch 1 {
        case -381:
            if false {
                var v23 int16 = 32767
                _ = v23
            }
        case 128:
            var v24 int32 = 2147483647
            _ = v24
        case 32767:
            var v25 uint8 = 144
            _ = v25
        default:
            if false {
                var v26 uint8 = 2
                var v27 uint = 876
                _ = v26
                _ = v27
            }
        }
        _ = v22
    } else {
        v28 := GenS1{F0: genFn0(-78, (^(v3)), false), F1: 1}
        _ = v28
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        v30 := v2[(fi29 * (fi29 + -354))]
        _ = fi29
        _ = v30
    }
    var v31 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
