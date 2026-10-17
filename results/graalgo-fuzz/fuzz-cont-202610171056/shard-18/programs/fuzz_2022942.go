package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 string; F2 string }
func genFn0(p0 string) uint8 {
    if true {
        var v6 int32 = -94
        if true {
            v7 := []int{-870}
            _ = v7
        }
        var v8 int32 = v6
        _ = v6
        _ = v8
    } else {
        switch ("a" + "hello") {
        case "":
            var v9 uint32 = 65535
            _ = v9
        default:
            v10 := GenS1{F0: 2}
            _ = v10
        }
    }
    v11 := GenS1{F0: 4294967295}
    switch p0 {
    case "a":
        var v12 string = ((((p0 + (p0 + ((p0 + "a") + (((p0 + "foo bar") + p0) + (p0 + (p0 + p0)))))) + "foo bar") + p0) + (p0 + (p0 + "a")))
        _ = v12
    case "":
        v13 := GenS1{F0: 266}
        _ = v13
    default:
        v14 := []int{-9223372036854775808, -1}
        _ = v14
    }
    var v15 int16 = 477
    var v16 int64 = -759
    _ = p0
    _ = v11
    _ = v15
    _ = v16
    return 127
}
func genFn1(p0 uint8, p1 int64) uint64 {
    var v17 uint32 = 488
    v18 := []int{-169, -911, 596}
    v17 += 65535
    v17 = v17
    p1 = 265
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 0
}

func main() {
    var v1 bool = true
    var v2 uint64 = 722
    var v3 uint = 42
    var v4 int8 = 127
    var v5 int8 = -1
    fmt.Println(v4, v1, v2)
    v19 := []int{32768, 0, 3}
    v1 = true
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 int32 = 1
        _ = fi20
        _ = v21
    }
    switch "" {
    case "\n\t":
        if v1 {
            fmt.Println(v5)
            v1 = false
            var v22 int64 = 42
            _ = v22
        } else {
            if (!(-1 > 909)) {
                v23 := GenS1{F0: v2}
                fmt.Println(v2, v3, v1)
                var v24 uint = (v3 * (v3 + (v3 * (v3 + v3))))
                _ = v23
                _ = v24
            }
            var v25 int64 = 0
            v25 -= v25
            _ = v25
        }
    case "a":
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
    case "hello":
        v26 := GenS1{F0: uint64(v3)}
        _ = v26
    }
    v27 := GenS2{F0: (!false), F1: "hello", F2: "foo bar"}
    var v28 int8 = (v5 + (v4 << (((v3 % (((v3 - 4294967295)) | 1))) & 63)))
    var v29 string = "hello"
    fmt.Println(v29, v4)
    if (float32(-2.25) > (1e10 / 1e10)) {
        v1 = v1
        fmt.Println(v3, v29)
        var v30 float64 = 1e10
        _ = v30
    } else {
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        for fi31 := 0; fi31 < 3; fi31++ {
            v2 += v2
            v3 += 42
            if true {
                v2 = 2
                if (2 <= 959) {
                    fmt.Println(v3)
                }
            } else {
                if (v1 || v1) {
                    for fi32 := 0; fi32 < 17; fi32++ {
                        _ = len(v19)
                        v29 = ""
                        _ = fi32
                    }
                    v28 = (v28 % ((42) | 1))
                }
                v1 = v1
                var v33 int32 = 0
                _ = v33
            }
            _ = fi31
        }
        var v34 bool = true
        _ = v34
    }
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
