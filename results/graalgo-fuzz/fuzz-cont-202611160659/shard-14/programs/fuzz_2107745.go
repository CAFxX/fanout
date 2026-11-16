package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int64; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
func genFn0(p0 int32) int8 {
    v6 := GenS1{F0: -641}
    fmt.Println(p0)
    p0 = p0
    for fi7 := 0; fi7 < 1; fi7++ {
        switch fi7 {
        case -405:
            var v8 uint = 42
            _ = v8
        }
        _ = fi7
    }
    _ = p0
    _ = v6
    return 38
}
func genFn1(p0 uint32, p1 string, p2 int8) uint32 {
    if (-428 < 0) {
        p1 = p1
        var v9 uint16 = 1
        v10 := []int{-844, 2147483647}
        _ = v10
        _ = v9
    } else {
        var v11 int16 = 1
        switch "" {
        case "":
            v12 := GenS1{F0: v11}
            _ = v12
        case "foo bar":
            v11 -= (v11 & -32768)
        }
        _ = v11
    }
    v13 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return p0
}

func main() {
    var v1 float64 = ((float64(((1.5 + ((0.1 * (float32(-2.25) + -2.25)) / 0.1)) + -2.25)) / 1.5) - -2.25)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    var v4 int16 = -1
    var v5 int16 = -598
    var v14 int8 = 100
    fmt.Println(v14, v1, v4)
    v15 := GenS2{F0: -1, F1: 1, F2: 0}
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 int16 = (v5 | (v5 ^ (^(v4))))
        if false {
            if false {
                v14 -= 87
                var v18 int8 = genFn0(838)
                _ = v18
            }
            fmt.Println(v1, v5)
            v1 += v1
        } else {
            v19 := v3[127]
            var v20 bool = (255 != 65535)
            _ = v19
            _ = v20
        }
        var v21 float64 = float64(1.5)
        _ = fi16
        _ = v17
        _ = v21
    }
    v4 = (v4 << ((780) & 63))
    switch 42 {
    case 0:
        if (618 != 18446744073709551615) {
            v22 := GenS3{F0: 4294967295}
            fmt.Println(v4, v5)
            var v23 int16 = 1
            _ = v22
            _ = v23
        }
    }
    switch 128 {
    case -873:
        var v24 int16 = v5
        _ = v24
    case 42:
        fmt.Println(v1)
    }
    v5 += (v4 | -654)
    for fi25 := 0; fi25 < 7; fi25++ {
        for fi26 := 0; fi26 < 12; fi26++ {
            v5 = (v5 << ((18446744073709551615) & 63))
            v1 += v1
            _ = fi26
        }
        v27 := GenS2{F0: -708, F1: 0, F2: -1}
        v28 := GenS4{F0: (!(v4 < v4))}
        _ = fi25
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
