package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 uint16; F1 int8; F2 uint8 }
func genFn0() string {
    var v5 int8 = -123
    switch ("a" + "\n\t") {
    case "\n\t":
        v6 := []int{-9223372036854775808, 9223372036854775807}
        _ = v6
    case "a":
        v7 := []int{-2147483648, 32768, 128, 32768}
        _ = v7
    }
    switch -300 {
    case 127:
        v5 -= (v5 / (((v5 << ((184) & 63))) | 1))
    case 329:
        if true {
            v8 := GenS1{F0: 900}
            switch "" {
            case "":
                var v9 uint16 = 2
                _ = v9
            default:
                v10 := GenS2{F0: true, F1: 2, F2: 1}
                _ = v10
            }
            _ = v8
        }
    }
    var v11 uint64 = 2
    _ = v11
    _ = v5
    return (("\n\t" + "a") + "")
}
func genFn1(p0 uint8, p1 uint32) string {
    v12 := GenS3{F0: -14, F1: 706}
    if (128 != 958) {
        var v13 int32 = -683
        var v14 uint8 = (p0 + (p0 << ((1) & 63)))
        var v15 int8 = 127
        _ = v13
        _ = v14
        _ = v15
    } else {
        p0 -= 0
        fmt.Println(p0, p1)
        if (((p1 ^ (p1 % (((p1 << ((2) & 63))) | 1))) != (p1 << ((437) & 63))) && true) {
            p0 += (^(p0))
            var v16 uint16 = 42
            fmt.Println(p1, v16)
            _ = v16
        } else {
            var v17 int8 = 42
            fmt.Println(p0)
            v18 := []int{-816, 147, 9223372036854775807}
            _ = v17
            _ = v18
        }
    }
    p0 = p0
    if true {
        p1 = (p1 / ((p1) | 1))
        p1 += p1
    }
    _ = p0
    _ = p1
    _ = v12
    return "a"
}

func main() {
    var v1 float64 = (3.14159 - 0.1)
    var v2 float32 = float32(v1)
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    v2 += v2
    _ = len(v3)
    for fi19 := 0; fi19 < 3; fi19++ {
        for fi20 := 0; fi20 < 19; fi20++ {
            switch fi20 {
            case -9223372036854775808:
                var v21 int16 = -1
                _ = v21
            case 11:
                var v22 string = "foo bar"
                _ = v22
            case -293:
                v2 -= v2
            default:
                v23 := GenS3{F0: 127, F1: 2}
                _ = v23
            }
            v24 := GenS2{F0: (!true), F1: 42, F2: 32767}
            _ = fi20
            _ = v24
        }
        v25 := GenS4{F0: (-(fi19)), F1: 32768}
        v26 := GenS2{F0: false, F1: 0, F2: 255}
        _ = fi19
        _ = v25
        _ = v26
    }
    if false {
        delete(v3, 9223372036854775807)
        v27 := GenS4{F0: 42, F1: 9223372036854775807}
        _ = v27
    }
    v28 := []int{}
    _ = len(v28)
    var v29 uint8 = 255
    v2 = v2
    for fi30 := 0; fi30 < 13; fi30++ {
        v29 = 255
        v31 := GenS5{F0: 255, F1: 42, F2: v29}
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        v1 = 0.0
        var v33 uint64 = 42
        switch (("\n\t" + "") + "\n\t") {
        case "\n\t":
            v29 = (v29 | 1)
        default:
            var v34 int32 = -2147483648
            _ = v34
        }
        _ = fi32
        _ = v33
    }
    var v35 int64 = -827
    for fi36 := 0; fi36 < 8; fi36++ {
        v37 := v4[0]
        _ = len(v3)
        _ = fi36
        _ = v37
    }
    switch -9223372036854775808 {
    case 4294967295:
        if ("" <= "\n\t") {
            v28 = append(v28, 32768)
        }
    case 32767:
        var v38 uint8 = (v29 * (v29 ^ v29))
        _ = v38
    default:
        if (true || true) {
            var v39 uint64 = 42
            _ = v39
        } else {
            fmt.Println(v35)
            v40 := GenS3{F0: 48, F1: 408}
            var v41 float32 = 0.1
            _ = v40
            _ = v41
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v4
}
