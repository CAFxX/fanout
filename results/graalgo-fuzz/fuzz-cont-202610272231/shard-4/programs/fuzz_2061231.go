package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int; F2 int }
type GenS3 struct { F0 uint64; F1 bool; F2 int8 }
type GenS4 struct { F0 int8; F1 uint16; F2 bool }
func genFn0(p0 int32) int {
    v4 := []int{1, 1}
    if len(v4) > 0 {
        v4[0] = 0
    }
    p0 += (p0 % (((p0 - (p0 - p0))) | 1))
    var v5 int8 = -1
    v6 := GenS1{F0: -69, F1: 0, F2: 42}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -739
}

func main() {
    var v1 int8 = 1
    var v2 uint8 = 1
    var v3 int = 0
    var v7 float64 = 0.0
    v8 := []int{768, (v3 % ((42) | 1)), (v3 / ((genFn0(-220)) | 1)), -248}
    if false {
        v2 = v2
        fmt.Println(v3, v3, v7)
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    } else {
        _ = len(v8)
    }
    var v9 uint64 = 976
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := GenS2{F0: v9, F1: (fi10 % ((int(v2)) | 1)), F2: v3}
            _ = fi11
            _ = v12
        }
        switch (v3 * v3) {
        case -9223372036854775808:
            v13 := GenS2{F0: (v9 & (v9 % ((v9) | 1))), F1: genFn0(0), F2: (v3 - (-(v3)))}
            _ = v13
        case -1:
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 string = "\n\t"
                v3 = (fi14 - genFn0(-1))
                _ = fi14
                _ = v15
            }
        default:
            v9 -= v9
        }
        if false {
            v1 += (v1 + (-(v1)))
            v9 = 1
        }
        _ = fi10
    }
    if len(v8) > 0 {
        v8[0] = v3
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 uint8 = (v2 & v2)
        var v18 int16 = 1
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS3{F0: v9, F1: false, F2: (v1 | (v1 | 59))}
    if false {
        v20 := GenS4{F0: v1, F1: 32767, F2: false}
        var v21 int64 = 1
        _ = v20
        _ = v21
    }
    switch (v3 & (v3 << ((703) & 63))) {
    case 752:
        v8 = append(v8, 2147483647)
    default:
        var v22 uint16 = 962
        _ = v22
    }
    fmt.Println(v9, v9, v3)
    switch genFn0(42) {
    case 32767:
        for fi23 := 0; fi23 < 3; fi23++ {
            var v24 float64 = v7
            v8 = append(v8, (fi23 * genFn0(1)))
            _ = fi23
            _ = v24
        }
    case 0:
        fmt.Println(v3)
    }
    var v25 string = ""
    v3 -= v3
    var v26 int16 = 0
    fmt.Println(v3)
    v1 = -123
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
