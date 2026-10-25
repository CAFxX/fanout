package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool }
func genFn0() int {
    if ((881 == 0) || false) {
        var v4 uint32 = 2147483648
        _ = v4
    } else {
        v5 := GenS1{F0: -1, F1: ((-1 <= 0) && (857 != 0)), F2: 4294967295}
        if true {
            v6 := GenS2{F0: 151, F1: -26}
            _ = v6
        } else {
            var v7 int8 = -1
            fmt.Println(v7)
            v7 += v7
            _ = v7
        }
        var v8 uint64 = 1
        _ = v5
        _ = v8
    }
    v9 := GenS3{F0: 1}
    _ = v9
    return 128
}
func genFn1() int64 {
    var v10 float64 = 0.0
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := GenS4{F0: 437, F1: 42}
        _ = fi11
        _ = v12
    }
    v10 = (((0.0 + 1.5) / 1.5) - v10)
    if true {
        v13 := []int{128}
        if len(v13) > 0 {
            v13[0] = 1
        }
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        _ = v13
    }
    _ = v10
    return 42
}

func main() {
    var v1 int16 = 1000
    v2 := []uint32{2, 2147483647, 520, 2147483648}
    v3 := make(map[uint32]uint32)
    var v14 uint8 = 76
    var v15 int16 = (v1 / (((v1 ^ (^(v1)))) | 1))
    v16 := GenS3{F0: 65535}
    var v17 float32 = -2.25
    var v18 int8 = 1
    for fi19 := 0; fi19 < 19; fi19++ {
        if (!(!(347 != (fi19 | (^(fi19)))))) {
            _ = len(v2)
        }
        fmt.Println(v1, v14)
        v20 := GenS5{F0: "hello"}
        _ = fi19
        _ = v20
    }
    switch 127 {
    case -1:
        var v21 bool = ((false || (1 > -78)) || false)
        _ = v21
    default:
        for fi22 := 0; fi22 < 16; fi22++ {
            fmt.Println(fi22, v1, v14)
            v2 = append(v2, 65535)
            _ = fi22
        }
    }
    if (1e10 <= 0.0) {
        fmt.Println(v17, v15)
    }
    var v23 uint32 = uint32(v14)
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 float32 = (0.0 / (((v17 + 1.5) - (v17 / v17)) + v17))
        v18 = v18
        if false {
            v26 := GenS6{F0: false}
            if ((v14 + 166) >= (v14 | 191)) {
                if (-9223372036854775808 > -1) {
                    var v27 float64 = ((float64(v17) + (0.0 / 1.0)) / 3.14159)
                    _ = v27
                }
            }
            fmt.Println(v18)
            _ = v26
        }
        _ = fi24
        _ = v25
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        v23 = 2147483648
        var v29 uint16 = 42
        var v30 uint32 = 2
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
}
