package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint }
func genFn0() bool {
    switch 1 {
    case 2147483647:
        var v6 int16 = 707
        _ = v6
    case -1:
        var v7 uint8 = 0
        _ = v7
    case -767:
        v8 := []int{127, 4294967295}
        _ = v8
    }
    v9 := []int{128, 214, 42}
    _ = v9
    return (-9223372036854775808 <= 0)
}
func genFn1(p0 int16, p1 int16, p2 uint) int64 {
    if false {
        p2 = (^(p2))
        var v10 int64 = 0
        v11 := GenS1{F0: -2147483648}
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 1; fi12++ {
            p2 = (p2 / (((p2 | p2)) | 1))
            fmt.Println(fi12, p2)
            _ = fi12
        }
        p2 = 2
    }
    var v13 bool = (!(42 == uint8(p2)))
    switch 127 {
    case 32768:
        fmt.Println(p2)
    case 204:
        v14 := []int{}
        _ = v14
    case 128:
        for fi15 := 0; fi15 < 13; fi15++ {
            for fi16 := 0; fi16 < 6; fi16++ {
                v17 := GenS1{F0: 42}
                _ = fi16
                _ = v17
            }
            var v18 float32 = 0.0
            v18 = 0.1
            _ = fi15
            _ = v18
        }
    default:
        var v19 float32 = 0.1
        _ = v19
    }
    v20 := GenS1{F0: -138}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v20
    return 0
}

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    var v4 int8 = 42
    var v5 float64 = ((0.0 - ((3.14159 * float64((3.14159 * 0.1))) / 1e10)) * 3.14159)
    v21 := GenS2{F0: "foo bar", F1: v1}
    switch "" {
    case "":
        v5 = v5
    default:
        v22 := []int{}
        _ = v22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint = 1
        var v25 int = (fi23 | (^(fi23)))
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 bool = (!false)
    var v27 float64 = v5
    v2[(v1 - v1)] = (-1 == 1)
    v28 := []int{(v1 | (v1 << ((0) & 63)))}
    var v29 int = v1
    v4 -= (v4 + v4)
    _ = len(v3)
    v5 += (float64(1e10) / v5)
    switch 680 {
    case 1:
        if v26 {
            var v30 bool = false
            _ = v30
        }
    case 32767:
        for fi31 := 0; fi31 < 4; fi31++ {
            v28 = append(v28, (v29 | 180))
            _ = fi31
        }
    case 127:
        var v32 float32 = -2.25
        _ = v32
    }
    v1 = 127
    fmt.Println(v26, v27)
    v33 := GenS3{F0: 0}
    v1 -= 32768
    var v34 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
