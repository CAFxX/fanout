package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int; F2 int64 }
func genFn0(p0 uint, p1 uint32) int8 {
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 uint8 = 175
            fmt.Println(fi7, fi7)
            switch (("\n\t" + "") + "\n\t") {
            case "a":
                v10 := []int{fi7}
                _ = v10
            case "foo bar":
                v9 -= v9
            case "":
                var v11 int = fi7
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        for fi12 := 0; fi12 < 16; fi12++ {
            fmt.Println(fi7, p1)
            v13 := []int{}
            p0 = (p0 - (p0 >> (((p0 >> ((p0) & 63))) & 63)))
            _ = fi12
            _ = v13
        }
        _ = fi7
    }
    p1 = (p1 % ((p1) | 1))
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "hello"
    var v3 float64 = (1e10 - ((0.0 + (float64((1.5 / (1.5 + 1e10))) - 0.0)) + 0.1))
    var v4 uint16 = 608
    v5 := make(map[string]uint32)
    v6 := []string{}
    if true {
        v14 := GenS1{F0: -1, F1: 0}
        v2 = "a"
        _ = v14
    }
    var v15 int64 = -1
    var v16 int32 = 2147483647
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := GenS2{F0: fi17}
        v4 += 255
        var v19 float32 = float32(v3)
        _ = fi17
        _ = v18
        _ = v19
    }
    if len(v6) > 0 {
        v6[0] = (v2 + v2)
    }
    var v20 string = v2
    v15 = 194
    if ((uint32(v4) == 1) || false) {
        var v21 string = (v20 + v2)
        v16 = v16
        _ = v21
    } else {
        switch v2 {
        case "\n\t":
            if (504 < 42) {
                v22 := GenS1{F0: (-(v16)), F1: 97}
                _ = v22
            }
        case "":
            switch ((v20 + v20) + v2) {
            case "\n\t":
                var v23 int = 127
                _ = v23
            default:
                var v24 float64 = ((float64(float32((-2.25 / v3))) + v3) / v3)
                _ = v24
            }
        case "hello":
            if (65535 < 4294967295) {
                var v25 float64 = (1.5 + -2.25)
                _ = len(v5)
                var v26 int = 0
                _ = v25
                _ = v26
            } else {
                var v27 int = 0
                _ = v27
            }
        default:
            v28 := GenS3{F0: (v16 - v16), F1: -571, F2: v15}
            _ = v28
        }
        delete(v5, v20)
    }
    if true {
        v29 := GenS1{F0: 42, F1: 2}
        _ = v29
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        var v31 uint16 = (v4 - (^(v4)))
        v4 = v31
        _ = fi30
        _ = v31
    }
    v4 += 65535
    var v32 uint8 = 0
    for fi33 := 0; fi33 < 8; fi33++ {
        fmt.Println(v32, v32)
        v2 = "a"
        _ = fi33
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        var v35 uint64 = 4294967295
        _ = fi34
        _ = v35
    }
    fmt.Println(v2)
    if (4294967295 < 42) {
        var v36 uint = 4294967295
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
