package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 string }
type GenS2 struct { F0 int8; F1 string; F2 int8 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint8, p1 int64) uint64 {
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 bool = false
        var v8 uint64 = 4294967295
        if true {
            for fi9 := 0; fi9 < 5; fi9++ {
                v10 := GenS1{F0: 128, F1: (v8 ^ (^(v8))), F2: ("a" + "foo bar")}
                if (1e10 > 0.1) {
                    v11 := GenS2{F0: 52, F1: ("\n\t" + ""), F2: 42}
                    v12 := []int{-1, fi9}
                    _ = v11
                    _ = v12
                } else {
                    fmt.Println(v8, p0)
                    v8 += (v8 << ((629) & 63))
                }
                var v13 uint32 = 686
                _ = fi9
                _ = v10
                _ = v13
            }
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v14 float32 = ((3.14159 + (0.1 * 0.0)) / (((0.1 / 1.0) + ((1.5 - (0.1 + (0.0 * 1e10))) / 1e10)) / 1e10))
    v15 := []int{}
    v16 := GenS3{F0: 1, F1: (p1 * p1), F2: 939}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return 1
}
func genFn1(p0 int32, p1 uint16) int {
    fmt.Println(p0, p1)
    if true {
        v17 := GenS4{F0: "foo bar", F1: 1}
        _ = v17
    }
    fmt.Println(p0, p0)
    if true {
        switch 0 {
        case -849:
            p1 = 255
        case -752:
            v18 := GenS2{F0: 100, F1: "foo bar", F2: -128}
            _ = v18
        }
        var v19 float32 = 1.5
        _ = v19
    } else {
        var v20 uint32 = 42
        var v21 bool = true
        switch "\n\t" {
        case "hello":
            v21 = v21
        }
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    var v1 float64 = (((((((0.1 * 0.0) * ((-2.25 / 1e10) + 1e10)) - (-2.25 + -2.25)) * 1.5) * float64(1.5)) - 3.14159) - 0.0)
    var v2 float32 = (1e10 * 0.0)
    v3 := []uint32{692, 2147483647}
    v4 := []int{9223372036854775807}
    v5 := []int{128, 1, -2147483648}
    fmt.Println(v1, v1)
    if (false && (1 != 65535)) {
        v1 += 1e10
        for fi22 := 0; fi22 < 19; fi22++ {
            if (-32768 != -943) {
                switch 9223372036854775807 {
                case 614:
                    _ = len(v5)
                case 9223372036854775807:
                    v2 = float32((0.1 - (((0.1 + 0.0) + v1) + 0.0)))
                default:
                    switch fi22 {
                    case 0:
                        v1 += (3.14159 * v1)
                    case 128:
                        v2 += (0.0 / ((v2 + v2) * v2))
                    case -2147483648:
                        var v23 int64 = -9223372036854775808
                        _ = v23
                    }
                }
                v24 := GenS2{F0: 0, F1: ("" + "hello"), F2: -128}
                _ = v24
            }
            if (false && (949 < 42)) {
                var v25 bool = false
                v25 = (false && true)
                v25 = v25
                _ = v25
            } else {
                var v26 uint32 = 154
                _ = v26
            }
            _ = fi22
        }
        var v27 float32 = v2
        _ = v27
    }
    if (1000 == 32767) {
        var v28 uint16 = 65535
        _ = v28
    }
    v29 := GenS5{F0: 0}
    if false {
        var v30 float64 = 1e10
        _ = v30
    } else {
        var v31 uint = 0
        _ = v31
    }
    v2 = ((3.14159 * 3.14159) + 0.1)
    var v32 float32 = 0.0
    v2 = v2
    v1 += v1
    var v33 int8 = 62
    if (true || true) {
        fmt.Println(v32)
    }
    v34 := GenS2{F0: (v33 - v33), F1: "", F2: 32}
    for fi35 := 0; fi35 < 15; fi35++ {
        for fi36 := 0; fi36 < 14; fi36++ {
            v32 = float32(v1)
            fmt.Println(v32, v1, v33)
            switch "hello" {
            case "hello":
                if true {
                    if false {
                        var v37 int = fi36
                        v32 += v2
                        fmt.Println(v2)
                        fmt.Println(v37, v2, v37)
                        _ = v37
                    }
                    var v38 int = int(v33)
                    _ = v38
                } else {
                    fmt.Println(v2)
                    v32 -= 1.5
                    fmt.Println(v32, v2, v32)
                    if true {
                        v39 := GenS5{F0: 127}
                        _ = v39
                    }
                }
            case "foo bar":
                if (fi35 <= (fi35 / (((^(fi35))) | 1))) {
                    fmt.Println(fi36, fi35)
                }
            }
            _ = fi36
        }
        _ = fi35
    }
    if (18446744073709551615 > 166) {
        var v40 string = (("\n\t" + ("a" + "hello")) + ("\n\t" + (("\n\t" + "\n\t") + "\n\t")))
        _ = v40
    }
    var v41 uint32 = 946
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v5
}
