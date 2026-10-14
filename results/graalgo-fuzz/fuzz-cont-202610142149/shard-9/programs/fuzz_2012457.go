package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint16; F1 bool; F2 int32 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int8; F1 bool; F2 uint64 }
func genFn0(p0 int32) int {
    v7 := []int{-1, -2147483648, 1, 4294967295}
    if (2 < 2) {
        var v8 int32 = p0
        var v9 int64 = 673
        v10 := GenS1{F0: 364, F1: 1}
        _ = v10
        _ = v8
        _ = v9
    }
    p0 += (p0 * (p0 / ((p0) | 1)))
    if false {
        var v11 uint64 = 4294967295
        var v12 int64 = 599
        if ((v12 / ((v12) | 1)) > v12) {
            if len(v7) > 0 {
                v7[0] = -573
            }
            fmt.Println(v11)
            v13 := GenS1{F0: (v11 | v11), F1: -1}
            _ = v13
        }
        _ = v11
        _ = v12
    }
    v14 := GenS1{F0: 591, F1: 127}
    _ = p0
    _ = v14
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 bool = (false && (!("a" < "\n\t")))
    var v2 int8 = 0
    var v3 uint64 = 749
    var v4 int16 = -711
    var v5 uint = 42
    v6 := make(map[uint32]int)
    var v15 float32 = (((3.14159 + float32(float64(1e10))) + 3.14159) / 1e10)
    v6[4294967295] = 127
    v15 += ((1.5 + v15) + v15)
    switch 128 {
    case 4294967295:
        for fi16 := 0; fi16 < 19; fi16++ {
            v17 := GenS2{F0: 2, F1: 482, F2: 100}
            _ = fi16
            _ = v17
        }
    case -579:
        v18 := GenS3{F0: 1, F1: v3}
        _ = v18
    case -1:
        fmt.Println(v15)
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS4{F0: 962, F1: (219 >= 42), F2: 291}
        _ = fi19
        _ = v20
    }
    v2 -= 0
    var v21 float32 = (v15 / 1e10)
    v22 := []int{2147483647, genFn0(683), 32767, -313}
    v23 := GenS2{F0: 255, F1: (^(v5)), F2: v2}
    if ((!v1) && ((v2 != 100) || (v1 && (!v1)))) {
        var v24 uint = 0
        _ = v24
    } else {
        v25 := GenS2{F0: 2, F1: 1, F2: -1}
        if (1 >= 255) {
            v26 := GenS5{F0: -843, F1: -32768}
            fmt.Println(v3)
            _ = v26
        } else {
            if false {
                if v1 {
                    v27 := GenS6{F0: v2, F1: v1, F2: v3}
                    _ = v27
                } else {
                    v21 = 3.14159
                }
                fmt.Println(v15)
            }
            fmt.Println(v1, v15, v4)
            var v28 bool = false
            _ = v28
        }
        switch "foo bar" {
        case "":
            delete(v6, 343)
        case "foo bar":
            var v29 int64 = -1
            _ = v29
        case "\n\t":
            var v30 float64 = float64((-2.25 / (1.5 - 0.0)))
            _ = v30
        }
        _ = v25
    }
    v22 = append(v22, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
