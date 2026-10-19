package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 bool; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool }
func genFn0() uint {
    var v4 int16 = 1000
    var v5 bool = (false && (false && true))
    v6 := GenS1{F0: false}
    v7 := GenS1{F0: (!true)}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 float64 = (3.14159 + 1.5)
    v2 := []uint32{754, 344, 2147483647, 2147483648}
    var v3 uint16 = 380
    v1 += (float64(3.14159) + float64(1e10))
    var v8 uint32 = 2147483648
    var v9 bool = (1.5 >= (3.14159 + ((1.5 * -2.25) - ((v1 + float64(0.1)) / 1e10))))
    var v10 int32 = -1
    v1 -= (v1 * 1.5)
    v10 = v10
    v10 = 42
    v11 := GenS2{F0: 249, F1: v10}
    v9 = v9
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 uint16 = v3
        var v14 uint = 42
        _ = fi12
        _ = v13
        _ = v14
    }
    if v9 {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if v9 {
            v3 -= uint16(v8)
            var v15 uint16 = v3
            _ = v15
        } else {
            v16 := GenS3{F0: true, F1: 1000, F2: 574}
            v8 -= (v8 - 65535)
            var v17 bool = false
            _ = v16
            _ = v17
        }
    } else {
        var v18 uint = 0
        _ = v18
    }
    var v19 int64 = -739
    if true {
        for fi20 := 0; fi20 < 18; fi20++ {
            switch fi20 {
            case 423:
                v1 += 1e10
            case 710:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 0:
                fmt.Println(fi20)
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v8 -= (v8 & 65535)
            _ = fi20
        }
        for fi21 := 0; fi21 < 19; fi21++ {
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 float64 = v1
                v9 = (77 > 92)
                _ = fi22
                _ = v23
            }
            if (1e10 > (v1 - v1)) {
                v3 += v3
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                var v24 uint = 4294967295
                var v25 uint64 = 0
                var v26 int64 = (-(v19))
                _ = v24
                _ = v25
                _ = v26
            }
            _ = fi21
        }
        v27 := GenS4{F0: v8}
        _ = v27
    } else {
        v28 := GenS5{F0: false}
        var v29 bool = ((100 != -38) || v9)
        fmt.Println(v8)
        _ = v28
        _ = v29
    }
    var v30 float64 = float64(3.14159)
    v31 := GenS1{F0: v9}
    var v32 int16 = -1
    v32 = v32
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v8
    _ = v9
}
