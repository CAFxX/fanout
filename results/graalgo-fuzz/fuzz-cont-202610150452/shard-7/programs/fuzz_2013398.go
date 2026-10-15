package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint32; F2 int8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int64, p1 string, p2 uint) int32 {
    v5 := []int{-1}
    p1 = ("a" + p1)
    if (false && false) {
        var v6 int32 = 42
        _ = v6
    } else {
        p0 = (p0 / (((^(p0))) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1(p0 int16) int {
    if (!false) {
        if true {
            var v7 uint8 = 42
            var v8 float64 = (0.0 - (1.5 / (float64((3.14159 + -2.25)) / 3.14159)))
            _ = v7
            _ = v8
        }
        var v9 bool = true
        if (!(16 > 3)) {
            p0 = (^(p0))
        } else {
            var v10 uint8 = 142
            _ = v10
        }
        _ = v9
    } else {
        if true {
            p0 -= (-(p0))
        } else {
            var v11 float64 = -2.25
            v12 := []int{}
            switch -2147483648 {
            case 127:
                var v13 int8 = -1
                _ = v13
            case 42:
                var v14 float32 = ((3.14159 * ((-2.25 / -2.25) - 0.0)) * 0.1)
                _ = v14
            }
            _ = v11
            _ = v12
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            fmt.Println(fi15, fi15)
            _ = fi15
        }
        fmt.Println(p0)
    }
    v16 := GenS1{F0: p0}
    p0 += (p0 >> ((65535) & 63))
    _ = p0
    _ = v16
    return 4294967295
}

func main() {
    v1 := []uint32{2, 2, 0, 65535}
    var v2 int16 = 420
    v3 := make(map[string]uint32)
    v4 := []string{"hello", "hello", ""}
    switch "a" {
    case "\n\t":
        if true {
            v17 := GenS2{F0: 42, F1: 299}
            var v18 float32 = (3.14159 + 3.14159)
            _ = v17
            _ = v18
        } else {
            for fi19 := 0; fi19 < 15; fi19++ {
                v20 := GenS2{F0: 243, F1: 0}
                _ = fi19
                _ = v20
            }
            var v21 uint16 = 1
            var v22 int = -317
            _ = v21
            _ = v22
        }
    default:
        v2 += v2
    }
    v2 = (v2 + (^(v2)))
    var v23 int32 = 974
    if (100 <= -55) {
        delete(v3, "a")
        fmt.Println(v2, v2)
    } else {
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 int = fi24
            var v26 int8 = 42
            if true {
                v2 -= v2
                fmt.Println(v23, v26)
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        var v27 int32 = v23
        if (0.1 >= 0.1) {
            v28 := GenS3{F0: v27, F1: 42, F2: -82}
            v2 += (v2 ^ (v2 ^ v2))
            var v29 uint32 = 4294967295
            _ = v28
            _ = v29
        }
        _ = v27
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        switch ("a" + "hello") {
        case "\n\t":
            v2 -= v2
        }
        var v31 float32 = (0.1 * ((((1.5 / 1.5) / 1.0) / 1.5) + 1e10))
        var v32 float64 = ((1.5 + 0.0) + (-2.25 - -2.25))
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 float64 = 1.5
    var v34 int = 42
    v35 := v3[("hello" + "")]
    v36 := GenS4{F0: 2147483647}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
