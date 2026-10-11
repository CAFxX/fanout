package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint8; F1 bool; F2 uint }
type GenS4 struct { F0 int; F1 uint8; F2 int16 }
func genFn0(p0 string, p1 uint16) int16 {
    fmt.Println(p1, p0)
    if (0 == 1) {
        if false {
            p1 = 65535
        }
        var v4 bool = false
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 uint32 = 917
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := []int{127, fi5, 32768, fi7}
                var v9 uint8 = 127
                if v4 {
                    v6 = 2147483647
                } else {
                    v10 := GenS1{F0: int32(fi7)}
                    var v11 float32 = (1e10 / (1e10 * 1.5))
                    if len(v8) > 0 {
                        v8[0] = fi5
                    }
                    _ = v10
                    _ = v11
                }
                _ = fi7
                _ = v8
                _ = v9
            }
            v12 := GenS2{F0: v4, F1: 127}
            _ = fi5
            _ = v12
            _ = v6
        }
        _ = v4
    }
    var v13 int8 = 1
    if false {
        v14 := []int{4294967295, 815}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v13
    return 1000
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float64 = -2.25
    var v3 uint64 = 0
    var v15 int16 = genFn0("a", 0)
    var v16 int8 = 42
    var v17 string = "a"
    switch 469 {
    case -209:
        v18 := []int{4294967295, -2147483648, 128, 1}
        _ = v18
    case 696:
        fmt.Println(v15, v16)
    default:
        var v19 string = v17
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        switch fi20 {
        case 128:
            var v21 uint64 = (v3 & 2)
            _ = v21
        case 328:
            for fi22 := 0; fi22 < 19; fi22++ {
                v23 := GenS3{F0: 127, F1: (v15 == int16(fi20)), F2: 816}
                var v24 int16 = (v15 << ((42) & 63))
                var v25 uint16 = 0
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
        default:
            for fi26 := 0; fi26 < 9; fi26++ {
                v1 = v1
                v27 := GenS4{F0: fi20, F1: 42, F2: (v15 | 32767)}
                _ = fi26
                _ = v27
            }
        }
        v28 := GenS1{F0: 2147483647}
        var v29 uint64 = 18446744073709551615
        _ = fi20
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        v31 := []int{fi30}
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 bool = true
        fmt.Println(v1, v2)
        _ = fi32
        _ = v33
    }
    switch -207 {
    case 395:
        for fi34 := 0; fi34 < 15; fi34++ {
            v17 = (v17 + v17)
            if false {
                fmt.Println(fi34, v15)
                v15 -= (-(v15))
            }
            _ = fi34
        }
    case 32767:
        var v35 uint8 = 65
        _ = v35
    case 0:
        fmt.Println(v16)
    default:
        var v36 float64 = 0.0
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
