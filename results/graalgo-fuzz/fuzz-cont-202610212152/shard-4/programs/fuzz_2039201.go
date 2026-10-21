package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 bool; F1 bool; F2 int }
func genFn0(p0 uint8, p1 int) uint {
    fmt.Println(p1)
    if (2 > 65535) {
        var v7 uint16 = 255
        v8 := GenS1{F0: -476}
        p0 = (p0 ^ 42)
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := []int{}
        p1 = (-(p1))
        _ = fi9
        _ = v10
    }
    var v11 int8 = 0
    _ = p0
    _ = p1
    _ = v11
    return 4294967295
}
func genFn1(p0 int64, p1 uint32) uint8 {
    v12 := []int{-9223372036854775808, 33}
    v13 := GenS2{F0: 4294967295, F1: 1}
    switch 42 {
    case -719:
        fmt.Println(p0, p1)
    case 2147483647:
        p0 -= -9223372036854775808
    case 911:
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    var v14 float32 = float32((0.1 - 0.1))
    var v15 int8 = 72
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 0
}

func main() {
    var v1 float64 = (1.5 * -2.25)
    var v2 uint16 = 42
    var v3 int = 170
    var v4 float64 = v1
    var v5 int32 = 621
    var v6 int32 = (v5 | v5)
    if (844 != 42) {
        v1 -= v4
        var v16 bool = ((311 < 1000) || false)
        fmt.Println(v3)
        _ = v16
    }
    fmt.Println(v3)
    switch 0 {
    case 32768:
        for fi17 := 0; fi17 < 6; fi17++ {
            v18 := []int{4294967295, v3, fi17, fi17}
            _ = fi17
            _ = v18
        }
    case 32767:
        v6 = (v5 + (v6 * (-(v5))))
    default:
        v19 := GenS3{F0: ((3.14159 + (((-2.25 / -2.25) / (-2.25 / ((0.1 * 0.1) + (3.14159 + (((0.1 / (-2.25 - ((0.1 + 1e10) / 1e10))) - 0.1) + 3.14159))))) / 1.5)) < (float32(float64((3.14159 - 1.5))) * ((1.5 * (-2.25 + (1e10 / 1e10))) - (1e10 / 1e10)))), F1: false, F2: (v3 % ((v3) | 1))}
        _ = v19
    }
    v20 := []int{v3}
    if true {
        _ = len(v20)
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 int16 = 613
        if len(v20) > 0 {
            v20[0] = (fi21 ^ fi21)
        }
        _ = fi21
        _ = v22
    }
    fmt.Println(v3, v1)
    v4 = v1
    v2 = (v2 ^ v2)
    v3 += (-(v3))
    switch "" {
    case "hello":
        v23 := GenS1{F0: v6}
        _ = v23
    case "":
        if ((!true) || (!true)) {
            fmt.Println(v5)
            if true {
                v20 = append(v20, v3)
            }
            v3 = (v3 << ((42) & 63))
        }
    case "\n\t":
        if false {
            var v24 string = "\n\t"
            if true {
                _ = len(v20)
            }
            var v25 float32 = (0.1 + 0.0)
            _ = v24
            _ = v25
        }
    default:
        if (-804 >= -1) {
            v20 = append(v20, v3)
        }
    }
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    var v26 uint32 = 2
    var v27 uint = genFn0(genFn1(0, 65535), v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
