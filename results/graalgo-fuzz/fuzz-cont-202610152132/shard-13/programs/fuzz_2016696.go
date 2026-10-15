package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int64; F2 int8 }
type GenS3 struct { F0 uint64; F1 string }
func genFn0() int {
    v7 := GenS1{F0: 1}
    var v8 int = 2147483647
    if ((false && (!(((0 == -1) || (float64(1e10) < 0.0)) || false))) && (!(true && true))) {
        v8 = v8
        v9 := GenS1{F0: 706}
        _ = v9
    } else {
        var v10 uint8 = 128
        _ = v10
    }
    v11 := []int{-1, v8, 2147483647}
    _ = v11
    _ = v7
    _ = v8
    return v8
}
func genFn1(p0 int32) uint32 {
    p0 += (p0 / (((p0 / ((p0) | 1))) | 1))
    p0 -= (p0 % ((p0) | 1))
    _ = p0
    return 0
}

func main() {
    v1 := []int{-244, 590, -749, -185}
    v2 := []int{32767}
    var v3 uint32 = 65535
    v4 := make(map[string]string)
    v5 := make(map[int]uint32)
    var v6 uint32 = v3
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            v3 += v6
            var v14 float64 = 1.5
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    switch 128 {
    case 32767:
        fmt.Println(v3, v6)
    default:
        if (!(65535 < 1)) {
            var v15 float64 = -2.25
            if false {
                v16 := GenS1{F0: 32767}
                switch "foo bar" {
                case "hello":
                    switch "foo bar" {
                    case "hello":
                        var v17 uint = 0
                        _ = v17
                    case "":
                        _ = len(v1)
                    default:
                        var v18 uint64 = 18446744073709551615
                        _ = v18
                    }
                case "a":
                    var v19 uint32 = v3
                    _ = v19
                case "":
                    _ = len(v4)
                }
                var v20 float64 = v15
                _ = v16
                _ = v20
            } else {
                var v21 int64 = 197
                _ = v21
            }
            if false {
                var v22 int16 = -278
                _ = len(v4)
                v3 = (v6 ^ v3)
                _ = v22
            }
            _ = v15
        }
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        for fi24 := 0; fi24 < 9; fi24++ {
            if false {
                fmt.Println(v6, fi24)
                v6 = 1
                if (-727 < -2147483648) {
                    var v25 bool = true
                    v3 = v6
                    _ = v25
                }
            }
            v26 := GenS1{F0: -233}
            v3 -= genFn1(2147483647)
            _ = fi24
            _ = v26
        }
        var v27 int32 = int32(fi23)
        _ = fi23
        _ = v27
    }
    v28 := GenS1{F0: -1}
    v3 = (v3 - 2)
    for fi29 := 0; fi29 < 9; fi29++ {
        v30 := GenS1{F0: -1}
        var v31 int = (fi29 - genFn0())
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 string = ("a" + ("a" + "\n\t"))
    var v33 float32 = float32((((1e10 * (3.14159 - 1e10)) / (1.5 * 1.5)) / 1e10))
    var v34 int8 = -128
    if true {
        delete(v5, -2147483648)
        if (v6 >= genFn1(-835)) {
            fmt.Println(v3, v3, v34)
        } else {
            _ = len(v5)
            for fi35 := 0; fi35 < 7; fi35++ {
                v33 = 0.0
                _ = fi35
            }
        }
    }
    v36 := GenS2{F0: 42, F1: -9223372036854775808, F2: 100}
    for fi37 := 0; fi37 < 12; fi37++ {
        v6 += 781
        var v38 int16 = 1
        fmt.Println(v38, v3)
        _ = fi37
        _ = v38
    }
    if true {
        var v39 int = 2147483647
        v6 += (^(v3))
        _ = v39
    }
    v3 = 946
    v34 = v34
    v40 := GenS3{F0: 758, F1: ((v32 + v32) + v32)}
    var v41 float64 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
