package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS5 struct { F0 int16; F1 bool }
func genFn0() int64 {
    switch 127 {
    case 32768:
        var v6 int = 0
        _ = v6
    case 128:
        if false {
            var v7 int = 32768
            v7 -= 9223372036854775807
            v7 = v7
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS1{F0: 42, F1: 2147483647, F2: 763}
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 255
    _ = v10
    return -1
}

func main() {
    var v1 bool = false
    var v2 uint8 = 2
    v3 := []uint32{2147483647, 69, 2147483647}
    v4 := make(map[int]int)
    var v5 uint64 = 4294967295
    switch -483 {
    case 32767:
        v11 := GenS1{F0: -69, F1: 144, F2: 467}
        _ = v11
    default:
        v12 := v4[128]
        _ = v12
    }
    if ((("foo bar" + ("" + "\n\t")) != ((("foo bar" + "hello") + "a") + "hello")) || (float64(1.5) != 3.14159)) {
        v13 := GenS2{F0: "", F1: 4294967295}
        fmt.Println(v2, v5, v2)
        if (false || (42 < 138)) {
            var v14 uint32 = 1
            _ = v14
        } else {
            v5 += v5
            var v15 float32 = 0.0
            v15 -= v15
            _ = v15
        }
        _ = v13
    } else {
        for fi16 := 0; fi16 < 18; fi16++ {
            v4[(fi16 - 2147483647)] = fi16
            fmt.Println(v1, v5)
            _ = len(v4)
            _ = fi16
        }
    }
    v2 = v2
    if true {
        v4[834] = 424
        fmt.Println(v1, v5)
        v5 += v5
    } else {
        fmt.Println(v5)
    }
    if v1 {
        var v17 float64 = float64(((1e10 * (((((1e10 + 0.1) * -2.25) / 1.0) / (-2.25 - 1.5)) / 1.5)) * (-2.25 * 0.0)))
        var v18 uint = 460
        if v1 {
            var v19 uint8 = (v2 - v2)
            _ = len(v4)
            for fi20 := 0; fi20 < 1; fi20++ {
                v1 = true
                var v21 int8 = -128
                for fi22 := 0; fi22 < 15; fi22++ {
                    var v23 uint32 = 630
                    _ = fi22
                    _ = v23
                }
                _ = fi20
                _ = v21
            }
            _ = v19
        }
        _ = v17
        _ = v18
    }
    switch 4294967295 {
    case 920:
        if (!v1) {
            fmt.Println(v2, v5, v1)
            var v24 uint64 = v5
            var v25 int = -1
            _ = v24
            _ = v25
        }
    case -831:
        var v26 float32 = 3.14159
        _ = v26
    case -2147483648:
        var v27 float32 = 0.1
        _ = v27
    }
    var v28 uint = 1
    switch "foo bar" {
    case "hello":
        v2 += 255
    case "":
        v28 = (v28 | (v28 ^ 42))
    }
    switch 42 {
    case 4294967295:
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 float64 = 1e10
            v30 -= v30
            if v1 {
                v31 := GenS2{F0: "foo bar", F1: 0}
                switch fi29 {
                case 127:
                    fmt.Println(v2)
                case -1:
                    v32 := GenS3{F0: v2, F1: 950, F2: 100}
                    _ = v32
                case 923:
                    var v33 uint32 = 2147483648
                    _ = v33
                }
                for fi34 := 0; fi34 < 11; fi34++ {
                    v2 += v2
                    _ = fi34
                }
                _ = v31
            } else {
                if (-887 <= 0) {
                    v35 := GenS4{F0: 1, F1: 1, F2: v5}
                    fmt.Println(v1, v1)
                    v30 += float64(3.14159)
                    _ = v35
                }
                var v36 uint8 = 2
                _ = v36
            }
            _ = fi29
            _ = v30
        }
    case 32768:
        var v37 uint16 = 0
        _ = v37
    case 127:
        v38 := GenS5{F0: -32768, F1: v1}
        _ = v38
    }
    delete(v4, -272)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
