package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int8; F1 int64 }
func genFn0(p0 uint64) uint8 {
    v7 := []int{}
    var v8 uint16 = 998
    _ = p0
    _ = v7
    _ = v8
    return 216
}
func genFn1(p0 uint16, p1 uint) int32 {
    p0 = (p0 & p0)
    switch "a" {
    case "\n\t":
        var v9 uint8 = 128
        _ = v9
    }
    if true {
        v10 := GenS1{F0: 197, F1: uint(p0), F2: 128}
        _ = v10
    }
    var v11 int32 = 569
    _ = p0
    _ = p1
    _ = v11
    return v11
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]int)
    v3 := []int{0, 496, -934, 0}
    var v4 uint16 = 42
    var v5 uint64 = 2
    var v6 uint = 42
    var v12 float64 = 3.14159
    if (true && (255 >= (^(v6)))) {
        v13 := GenS1{F0: 136, F1: (v6 + (v6 >> (((v6 | (v6 >> ((42) & 63)))) & 63))), F2: 127}
        _ = v13
    }
    if false {
        v5 -= (v5 * v5)
        fmt.Println(v12, v4, v4)
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 float64 = (v12 - v12)
            _ = fi14
            _ = v15
        }
    } else {
        var v16 int64 = 1
        v17 := GenS1{F0: genFn0(0), F1: 1, F2: -1}
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 107, F1: (v6 | (v6 << ((18446744073709551615) & 63))), F2: -2147483648}
    var v19 string = "a"
    _ = len(v3)
    v4 -= 964
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := GenS1{F0: 42, F1: v6, F2: 2147483647}
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        switch v19 {
        case "foo bar":
            v23 := GenS1{F0: 44, F1: (v6 ^ 0), F2: (fi22 | 127)}
            _ = v23
        default:
            var v24 int16 = 498
            _ = v24
        }
        v5 = 48
        _ = fi22
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        switch fi25 {
        case 4294967295:
            v26 := GenS2{F0: false}
            _ = v26
        case 42:
            v6 -= 4294967295
        case -2147483648:
            v12 -= v12
        default:
            v27 := v2[(v19 + v19)]
            _ = v27
        }
        v12 = v12
        for fi28 := 0; fi28 < 18; fi28++ {
            var v29 uint32 = 841
            _ = fi28
            _ = v29
        }
        _ = fi25
    }
    if (((("a" + v19) + v19) + v19) > v19) {
        switch -644 {
        case 127:
            if (false && ("a" >= v19)) {
                fmt.Println(v5, v5)
            } else {
                v19 = v19
                for fi30 := 0; fi30 < 10; fi30++ {
                    var v31 uint8 = 32
                    switch v19 {
                    case "a":
                        v12 += v12
                    case "\n\t":
                        v32 := GenS3{F0: 128, F1: -456}
                        _ = v32
                    default:
                        var v33 bool = ((-(fi30)) == fi30)
                        _ = v33
                    }
                    _ = fi30
                    _ = v31
                }
            }
        case -9223372036854775808:
            fmt.Println(v4, v5)
        case -2147483648:
            var v34 float64 = ((0.1 / (v12 - 3.14159)) - (0.1 + ((1.5 + (0.0 + ((v12 - (((v12 - v12) / (1.5 / (1.5 + ((float64(float32(v12)) / v12) / 3.14159)))) * v12)) * (1e10 * v12)))) * v12)))
            _ = v34
        }
        if true {
            for fi35 := 0; fi35 < 2; fi35++ {
                v36 := GenS3{F0: 143, F1: -772}
                v5 += (v5 - (v5 << (((^(v6))) & 63)))
                _ = fi35
                _ = v36
            }
            var v37 uint8 = genFn0(v5)
            switch -2147483648 {
            case 2147483647:
                v38 := GenS4{F0: 100, F1: 9223372036854775807}
                _ = v38
            }
            _ = v37
        }
    }
    var v39 float32 = 0.1
    v19 = v19
    v19 = v19
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
