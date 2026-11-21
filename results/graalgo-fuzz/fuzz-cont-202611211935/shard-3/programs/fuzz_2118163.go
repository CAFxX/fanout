package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 string; F2 bool }
func genFn0() uint16 {
    v7 := []int{128, 42, 9223372036854775807}
    var v8 float32 = 0.1
    _ = v7
    _ = v8
    return 313
}
func genFn1(p0 uint, p1 int64) int32 {
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 bool = true
        for fi11 := 0; fi11 < 18; fi11++ {
            p0 -= (p0 % (((p0 * p0)) | 1))
            var v12 uint32 = 690
            fmt.Println(v10, v10)
            _ = fi11
            _ = v12
        }
        v10 = false
        _ = fi9
        _ = v10
    }
    var v13 uint16 = 2
    var v14 uint64 = 0
    v13 += (v13 % (((v13 ^ 258)) | 1))
    p0 += (p0 ^ 792)
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = ("" + (("foo bar" + "\n\t") + "\n\t"))
    v3 := []int{}
    v4 := make(map[string]string)
    var v5 uint32 = 0
    v6 := []int{9223372036854775807, -1}
    v15 := GenS1{F0: genFn0(), F1: (v5 * v5)}
    for fi16 := 0; fi16 < 2; fi16++ {
        v17 := GenS2{F0: -1}
        v5 += (^(v5))
        _ = fi16
        _ = v17
    }
    var v18 uint = 0
    switch -9223372036854775808 {
    case -1:
        var v19 uint8 = 42
        _ = v19
    default:
        v5 = 42
    }
    var v20 int32 = 0
    for fi21 := 0; fi21 < 13; fi21++ {
        if (!true) {
            switch (fi21 << (((v18 << ((v18) & 63))) & 63)) {
            case 42:
                var v22 int32 = v20
                _ = v22
            case -582:
                fmt.Println(v5, v5)
            }
            v18 = (v18 * (v18 - (v18 / ((364) | 1))))
        } else {
            var v23 uint8 = 178
            _ = v23
        }
        var v24 int16 = 32767
        _ = fi21
        _ = v24
    }
    var v25 uint32 = (v5 + (^(v5)))
    v25 = (v5 - v25)
    v26 := GenS3{F0: (809 == 306), F1: (v2 + v2), F2: (3.14159 == 3.14159)}
    var v27 int64 = 1
    if (0 > 42) {
        v20 = (v20 + 0)
    } else {
        _ = len(v3)
        if true {
            var v28 float32 = 0.0
            var v29 int32 = genFn1(v18, -14)
            _ = v28
            _ = v29
        } else {
            v27 += (v27 + v27)
        }
    }
    v20 = (v20 & (v20 / (((v20 | (-(v20)))) | 1)))
    if false {
        var v30 string = v2
        _ = v30
    }
    fmt.Println(v20)
    var v31 int16 = 1
    if len(v6) > 0 {
        v6[0] = 0
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
