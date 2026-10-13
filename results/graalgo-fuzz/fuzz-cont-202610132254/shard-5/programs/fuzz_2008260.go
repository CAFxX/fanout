package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int }
func genFn0(p0 int32, p1 uint) int {
    p0 += (-(p0))
    if (255 > 255) {
        var v6 uint16 = 1
        p0 = p0
        p1 += p1
        _ = v6
    }
    var v7 float32 = (1.5 - (-2.25 + ((0.0 / 0.1) + 0.0)))
    v8 := GenS1{F0: 2}
    for fi9 := 0; fi9 < 7; fi9++ {
        v7 -= v7
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint64, p1 string) int8 {
    var v10 int64 = 42
    p0 += p0
    if (0.1 <= float32(1.5)) {
        fmt.Println(p0, v10)
    } else {
        v11 := GenS1{F0: 2147483647}
        v12 := GenS1{F0: uint32(p0)}
        if false {
            v13 := []int{32768}
            switch 42 {
            case 4294967295:
                var v14 int64 = (v10 % (((v10 | (v10 | (v10 % (((v10 | (v10 | v10))) | 1))))) | 1))
                _ = v14
            case 32767:
                p1 = p1
            case 2147483647:
                v15 := GenS2{F0: 2147483647, F1: -2147483648}
                _ = v15
            default:
                fmt.Println(p1, v10)
            }
            _ = v13
        } else {
            v16 := []int{int(v10), 1}
            _ = v16
        }
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return -58
}

func main() {
    var v1 float32 = -2.25
    var v2 int16 = 42
    v3 := make(map[uint32]int)
    v4 := []int{-382, -772}
    var v5 int = -157
    fmt.Println(v1)
    v2 = (v2 * (v2 - -32768))
    v5 += genFn0(42, 1)
    for fi17 := 0; fi17 < 2; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            if true {
                v5 += (fi17 - genFn0(-462, 292))
                _ = len(v4)
            }
            switch (^(fi18)) {
            case 748:
                v1 = 0.0
            case 42:
                var v19 int32 = int32(fi17)
                _ = v19
            case 128:
                v20 := GenS2{F0: 0, F1: (-(fi18))}
                _ = v20
            }
            _ = fi18
        }
        v21 := GenS2{F0: 874, F1: (fi17 | (-(v5)))}
        _ = fi17
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 uint16 = 255
        if (false && true) {
            var v24 int64 = -1
            _ = v24
        } else {
            for fi25 := 0; fi25 < 5; fi25++ {
                var v26 int64 = -1
                var v27 uint8 = 0
                v5 = 1
                _ = fi25
                _ = v26
                _ = v27
            }
            for fi28 := 0; fi28 < 5; fi28++ {
                v1 = v1
                v23 = v23
                _ = fi28
            }
        }
        _ = fi22
        _ = v23
    }
    var v29 uint8 = 1
    for fi30 := 0; fi30 < 7; fi30++ {
        var v31 int64 = -1
        _ = fi30
        _ = v31
    }
    v1 += (1e10 + ((v1 / v1) - (float32(0.1) - (((v1 * (float32(1.5) - ((v1 + ((v1 * (1e10 / v1)) + (v1 + v1))) - (-2.25 * (v1 / (v1 - 0.0)))))) / v1) - v1))))
    v32 := GenS1{F0: 315}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
