package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int8 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int; F1 int8 }
func genFn0(p0 uint64) int8 {
    if (1 >= 508) {
        var v7 uint32 = 807
        _ = v7
    } else {
        p0 = p0
        if false {
            if (p0 > 1) {
                var v8 int16 = 42
                v9 := GenS1{F0: ("" > "hello"), F1: "hello", F2: -46}
                if (int8(v8) >= -71) {
                    v8 += v8
                } else {
                    var v10 int16 = (v8 & (v8 | v8))
                    switch "a" {
                    case "":
                        var v11 uint32 = 2147483647
                        _ = v11
                    case "foo bar":
                        fmt.Println(v10, v8, v10)
                    case "\n\t":
                        v12 := []int{42, -1, -9223372036854775808}
                        _ = v12
                    }
                    switch "\n\t" {
                    case "a":
                        v13 := []int{}
                        _ = v13
                    }
                    _ = v10
                }
                _ = v8
                _ = v9
            } else {
                for fi14 := 0; fi14 < 19; fi14++ {
                    fmt.Println(p0, p0)
                    _ = fi14
                }
                var v15 uint16 = 771
                _ = v15
            }
            v16 := GenS2{F0: 664, F1: -1, F2: 255}
            _ = v16
        } else {
            var v17 int8 = 1
            var v18 float32 = float32((3.14159 - ((0.1 - 0.0) / float64((((1e10 + 3.14159) + 1.5) / float32(-2.25))))))
            _ = v17
            _ = v18
        }
    }
    switch 32767 {
    case 32767:
        switch 32768 {
        case 32768:
            var v19 int32 = -630
            _ = v19
        }
    case -336:
        v20 := []int{-643, -963, -2147483648, 108}
        _ = v20
    case 1:
        p0 += p0
    }
    _ = p0
    return 42
}

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[uint32]bool)
    var v3 bool = ((0.0 * 0.0) <= (3.14159 - -2.25))
    var v4 uint16 = 0
    v5 := make(map[int]uint32)
    var v6 string = "hello"
    var v21 float64 = ((0.0 * 0.0) - 1e10)
    if (32768 == -1) {
        if (-128 > 5) {
            if v3 {
                var v22 int16 = int16(v4)
                switch v6 {
                case "foo bar":
                    for fi23 := 0; fi23 < 17; fi23++ {
                        v22 += (^(v22))
                        _ = fi23
                    }
                case "hello":
                    v21 = ((v21 - float64(float32(3.14159))) / 1.5)
                case "\n\t":
                    fmt.Println(v3)
                default:
                    v24 := GenS2{F0: (v4 << ((255) & 63)), F1: -16, F2: v4}
                    _ = v24
                }
                _ = v22
            }
        }
        var v25 float64 = 3.14159
        var v26 int32 = -2147483648
        _ = v25
        _ = v26
    }
    var v27 int32 = 0
    v28 := GenS1{F0: (!((v3 && v3) && (1 != -3))), F1: "a", F2: -1}
    v6 = v6
    v5[-185] = 2147483648
    if v3 {
        if (genFn0(1) == 42) {
            fmt.Println(v4, v4)
            for fi29 := 0; fi29 < 5; fi29++ {
                var v30 int32 = int32(v4)
                v5[(fi29 + (fi29 - fi29))] = 2147483647
                fmt.Println(v4)
                _ = fi29
                _ = v30
            }
            var v31 int32 = (v27 & (-(v27)))
            _ = v31
        }
        fmt.Println(v21)
        var v32 int64 = 9223372036854775807
        _ = v32
    }
    v27 += v27
    var v33 uint32 = 2147483648
    if v3 {
        v3 = (v6 != v6)
        var v34 uint = 586
        _ = v34
    } else {
        v35 := GenS3{F0: 0, F1: 0}
        for fi36 := 0; fi36 < 13; fi36++ {
            for fi37 := 0; fi37 < 13; fi37++ {
                v38 := GenS1{F0: false, F1: "\n\t", F2: 100}
                var v39 float64 = v21
                var v40 uint = 349
                v39 -= v21
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
            }
            var v41 float32 = float32((0.0 * 0.0))
            _ = fi36
            _ = v41
        }
        var v42 uint32 = (v33 / (((v33 % ((v33) | 1))) | 1))
        _ = v35
        _ = v42
    }
    var v43 uint8 = 228
    for fi44 := 0; fi44 < 9; fi44++ {
        v33 = 1
        _ = fi44
    }
    fmt.Println(v3, v6, v3)
    v45 := GenS4{F0: 32767, F1: 18}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v43
    _ = v45
    _ = v5
    _ = v6
}
