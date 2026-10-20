package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint }
type GenS2 struct { F0 int32; F1 uint32; F2 int64 }
type GenS3 struct { F0 bool; F1 uint16; F2 int8 }
func genFn0(p0 int) int16 {
    if (1e10 < 3.14159) {
        var v4 float64 = 0.1
        switch "\n\t" {
        case "a":
            v5 := []int{p0, 591}
            _ = v5
        case "hello":
            var v6 uint64 = 391
            _ = v6
        default:
            v4 = (((0.1 / 1.5) - (((float64((1.5 / -2.25)) / -2.25) + v4) / (((v4 / (v4 + v4)) * v4) + (0.1 * (1.5 * (float64(1e10) - v4)))))) - 1.5)
        }
        _ = v4
    } else {
        v7 := []int{(p0 >> ((633) & 63)), (p0 + (p0 * (-(p0)))), (p0 % ((p0) | 1)), -1}
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -32
    v3 := make(map[uint32]string)
    var v8 int16 = 42
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 int = (fi9 >> ((18446744073709551615) & 63))
        for fi11 := 0; fi11 < 1; fi11++ {
            fmt.Println(fi11, fi9, v2)
            fmt.Println(fi11, v2, v2)
            if (42 == 0) {
                var v12 int = fi9
                _ = v12
            }
            _ = fi11
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            fmt.Println(v8)
            v14 := []int{(fi13 << ((255) & 63)), 1, fi13, 490}
            fmt.Println(v10)
            _ = fi13
            _ = v14
        }
        _ = fi9
        _ = v10
    }
    var v15 float32 = ((1e10 / (((float32(0.0) / 1.5) * (((3.14159 * float32((1.5 / 1.5))) / (float32(0.0) + -2.25)) * 1.5)) - (((-2.25 - float32(0.1)) - float32(0.1)) + ((3.14159 - -2.25) / 0.1)))) - 1e10)
    if false {
        v16 := GenS1{F0: 1, F1: 690, F2: 2}
        var v17 bool = true
        v17 = true
        _ = v16
        _ = v17
    } else {
        var v18 uint64 = 2
        _ = v18
    }
    fmt.Println(v8)
    var v19 string = ("hello" + "")
    if ((1e10 - 1e10) >= 0.1) {
        var v20 int16 = (-(v8))
        fmt.Println(v19, v15)
        v20 += (v8 % (((v8 << ((894) & 63))) | 1))
        _ = v20
    }
    v21 := []int{}
    v22 := GenS2{F0: 42, F1: 65535, F2: 1}
    fmt.Println(v15, v8)
    v19 = "foo bar"
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    if ((628 == 127) && false) {
        v8 += 1000
        v8 += v8
        v2 = -57
    } else {
        v23 := GenS1{F0: 794, F1: 98, F2: 1}
        _ = v23
    }
    v15 -= (((v15 / v15) - v15) / v15)
    var v24 int16 = -32768
    for fi25 := 0; fi25 < 5; fi25++ {
        v19 = v19
        v2 += (v2 & (v2 >> ((2) & 63)))
        for fi26 := 0; fi26 < 20; fi26++ {
            v2 -= -128
            v27 := GenS2{F0: 1, F1: 1, F2: -934}
            switch ("" + v19) {
            case "\n\t":
                delete(v3, 2)
            default:
                v15 += v15
            }
            v2 = -128
            v8 -= (v8 * v8)
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    v28 := GenS3{F0: (float32((float64((0.1 * v15)) / (0.0 - ((((1e10 + 3.14159) + float64(float32(0.0))) / (1.5 + 1.5)) - float64((0.1 * v15)))))) != v15), F1: 1, F2: (v2 + 98)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v28
    _ = v3
    _ = v8
}
