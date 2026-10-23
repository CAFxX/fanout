package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int8; F2 uint }
type GenS5 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS6 struct { F0 bool; F1 uint8; F2 int8 }
func genFn0(p0 int64) bool {
    var v6 uint32 = 4294967295
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            v6 += 2147483647
            v6 = v6
            v6 = v6
            _ = fi8
        }
        for fi9 := 0; fi9 < 4; fi9++ {
            switch fi7 {
            case -1:
                if (!false) {
                    v10 := GenS1{F0: 353}
                    if (fi9 < fi7) {
                        var v11 float64 = (1e10 / (3.14159 - (0.0 - 0.0)))
                        var v12 int = fi7
                        _ = v11
                        _ = v12
                    } else {
                        fmt.Println(p0, fi7, fi7)
                        v6 = v6
                        v13 := []int{(fi7 ^ fi9), (-(fi7)), fi7}
                        _ = v13
                    }
                    v14 := GenS1{F0: 4294967295}
                    _ = v10
                    _ = v14
                } else {
                    switch (fi9 + (^(fi9))) {
                    case -1:
                        if (2147483648 > 42) {
                            var v15 int32 = 500
                            v15 = v15
                            var v16 bool = true
                            _ = v15
                            _ = v16
                        } else {
                            fmt.Println(fi9, p0, p0)
                            p0 -= 0
                        }
                    default:
                        v17 := GenS2{F0: 0, F1: 0, F2: 255}
                        _ = v17
                    }
                    var v18 uint8 = 1
                    _ = v18
                }
            case 42:
                var v19 int = fi7
                _ = v19
            case 0:
                p0 = p0
            default:
                if (!false) {
                    v20 := []int{(fi9 - (fi9 ^ (fi7 >> ((408) & 63)))), fi9, fi7, (fi9 & (fi9 + fi9))}
                    p0 = -22
                    _ = v20
                } else {
                    p0 = 9223372036854775807
                    v21 := GenS2{F0: 42, F1: -1, F2: 862}
                    var v22 string = "foo bar"
                    _ = v21
                    _ = v22
                }
            }
            v6 += v6
            _ = fi9
        }
        _ = fi7
    }
    _ = p0
    _ = v6
    return (!false)
}

func main() {
    var v1 bool = false
    var v2 uint16 = 1
    var v3 int32 = -620
    var v4 int64 = 42
    var v5 uint8 = 255
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 int8 = 0
        switch 1 {
        case 128:
            v25 := GenS3{F0: (((-2.25 * float64((1.5 + ((((-2.25 + -2.25) - 3.14159) + ((-2.25 * -2.25) + float32((0.1 / 1.0)))) - 1.5)))) * (1e10 + ((1e10 + 0.1) * 0.0))) < 1.5)}
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    if v1 {
        var v26 uint8 = v5
        var v27 bool = (!true)
        v5 = 2
        _ = v26
        _ = v27
    } else {
        v28 := GenS4{F0: (v5 - v5), F1: 42, F2: 4294967295}
        _ = v28
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        switch "hello" {
        case "":
            switch (("" + "a") + "a") {
            case "hello":
                fmt.Println(v2)
            }
        }
        v30 := GenS5{F0: v4, F1: 334, F2: 0}
        _ = fi29
        _ = v30
    }
    v2 = v2
    var v31 int32 = v3
    var v32 int16 = -424
    v32 -= (v32 - (v32 + v32))
    v33 := []int{-1, 1, -64}
    fmt.Println(v2, v4)
    v2 -= v2
    v34 := GenS6{F0: v1, F1: 1, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
