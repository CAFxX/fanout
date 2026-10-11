package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint; F2 bool }
type GenS3 struct { F0 uint }
func genFn0(p0 string, p1 bool) uint64 {
    var v5 bool = (0 != 255)
    p0 = p0
    if ((v5 || false) || (128 < 194)) {
        if (((float64((0.0 + -2.25)) * 3.14159) != float64(-2.25)) || false) {
            var v6 int16 = -1
            fmt.Println(p1)
            _ = v6
        }
        fmt.Println(p0)
    }
    v5 = v5
    var v7 int8 = 100
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return 1
}
func genFn1(p0 uint32) int32 {
    switch -1 {
    case 32768:
        v8 := GenS1{F0: 1}
        _ = v8
    }
    p0 -= 1
    if false {
        v9 := GenS2{F0: false, F1: 4294967295, F2: true}
        _ = v9
    } else {
        if (0 > 1) {
            if false {
                for fi10 := 0; fi10 < 20; fi10++ {
                    var v11 uint = 0
                    p0 -= 1
                    _ = fi10
                    _ = v11
                }
                p0 -= p0
                p0 += (^(p0))
            }
            for fi12 := 0; fi12 < 8; fi12++ {
                v13 := []int{(^(fi12)), fi12, fi12}
                _ = fi12
                _ = v13
            }
            v14 := []int{9223372036854775807, 1, -9223372036854775808, 32768}
            _ = v14
        }
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 int8 = -128
            _ = fi15
            _ = v16
        }
    }
    _ = p0
    return 187
}

func main() {
    var v1 int = 128
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    var v4 uint16 = 634
    for fi17 := 0; fi17 < 12; fi17++ {
        for fi18 := 0; fi18 < 8; fi18++ {
            v19 := GenS3{F0: 255}
            var v20 int = fi18
            v1 -= v20
            _ = fi18
            _ = v19
            _ = v20
        }
        v1 = (fi17 * (fi17 ^ fi17))
        _ = fi17
    }
    v1 = (v1 * -586)
    v1 += (v1 * v1)
    v4 -= (v4 >> ((0) & 63))
    var v21 int8 = -128
    v22 := []int{}
    fmt.Println(v4, v21, v4)
    v1 += v1
    var v23 uint16 = (v4 - v4)
    v21 = 127
    delete(v2, (-(v1)))
    var v24 uint64 = 42
    var v25 uint64 = (^(v24))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
