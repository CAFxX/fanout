package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int; F2 int32 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint8, p1 int32) int64 {
    if (-1 != -9223372036854775808) {
        if (("\n\t" != ("a" + "hello")) || true) {
            switch -9223372036854775808 {
            case 9223372036854775807:
                switch 127 {
                case -320:
                    v6 := GenS1{F0: true, F1: 42, F2: -1}
                    _ = v6
                case -1:
                    v7 := GenS2{F0: -3, F1: 12}
                    _ = v7
                }
            }
        } else {
            var v8 uint = 349
            v8 = v8
            _ = v8
        }
        p0 = p0
        fmt.Println(p0)
    }
    var v9 float32 = 0.0
    switch -750 {
    case 4294967295:
        v10 := []int{127, -324, 4294967295, -784}
        _ = v10
    case 42:
        fmt.Println(p0, p0, v9)
    default:
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 float64 = float64(v9)
            _ = fi11
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = v9
    return -9223372036854775808
}
func genFn1() int64 {
    var v13 uint16 = 32767
    switch ("" + "a") {
    case "hello":
        if false {
            v14 := []int{-413, 9223372036854775807}
            var v15 uint32 = 162
            switch -182 {
            case 1:
                var v16 int64 = 42
                _ = v16
            default:
                v17 := GenS2{F0: -268, F1: -645}
                _ = v17
            }
            _ = v14
            _ = v15
        } else {
            v13 += v13
            fmt.Println(v13)
        }
    case "foo bar":
        v18 := []int{}
        _ = v18
    default:
        if false {
            v19 := []int{32767, 2147483647}
            if (-40 <= 127) {
                v20 := GenS1{F0: (1000 <= -1), F1: 4294967295, F2: -218}
                _ = v20
            }
            v21 := GenS1{F0: false, F1: 65535, F2: 996}
            _ = v19
            _ = v21
        } else {
            switch "\n\t" {
            case "a":
                if (v13 > 42) {
                    v22 := []int{9223372036854775807, -687, int(v13), -349}
                    var v23 bool = true
                    _ = v22
                    _ = v23
                } else {
                    v24 := []int{}
                    v13 -= (v13 ^ (v13 * 2))
                    var v25 float64 = (3.14159 * ((3.14159 + 0.0) / 1.0))
                    _ = v24
                    _ = v25
                }
            case "\n\t":
                var v26 uint8 = 94
                _ = v26
            case "":
                var v27 float64 = -2.25
                _ = v27
            default:
                v13 -= v13
            }
        }
    }
    var v28 string = "foo bar"
    _ = v13
    _ = v28
    return 1
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]string)
    v3 := []uint32{4294967295, 65535, 2}
    v4 := make(map[int]uint32)
    var v5 uint32 = 42
    switch 32768 {
    case 9223372036854775807:
        var v29 uint8 = 240
        _ = v29
    case 4294967295:
        if (1.5 != (float32(0.1) * 1e10)) {
            var v30 int16 = -1
            v4[-1] = (v5 << ((2) & 63))
            for fi31 := 0; fi31 < 13; fi31++ {
                var v32 float32 = 3.14159
                _ = fi31
                _ = v32
            }
            _ = v30
        } else {
            switch -1 {
            case 2147483647:
                v1 = v1
            case 38:
                switch "a" {
                case "":
                    v33 := GenS3{F0: (v5 - v5)}
                    _ = v33
                case "a":
                    v34 := GenS4{F0: 42, F1: -9223372036854775808, F2: 2147483647}
                    _ = v34
                default:
                    v5 = (v5 * (v5 | v5))
                }
            case -9223372036854775808:
                if true {
                    v1 += (v1 ^ v1)
                    v35 := GenS2{F0: 0, F1: 42}
                    _ = v35
                }
            }
            v36 := GenS5{F0: v5}
            _ = v36
        }
    case 69:
        var v37 uint64 = uint64(v5)
        _ = v37
    default:
        v38 := GenS3{F0: 904}
        _ = v38
    }
    var v39 int64 = -9223372036854775808
    var v40 uint32 = (^(v5))
    var v41 uint = 835
    fmt.Println(v40, v40)
    v40 -= (v40 / ((uint32(v41)) | 1))
    var v42 uint16 = 42
    v43 := v2["hello"]
    var v44 int8 = -105
    v4[0] = (^(v5))
    v45 := GenS1{F0: false, F1: (v41 - 255), F2: 127}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
