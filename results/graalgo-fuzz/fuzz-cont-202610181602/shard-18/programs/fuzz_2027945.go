package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 int8; F1 uint16; F2 uint }
func genFn0() uint8 {
    var v6 int16 = -32768
    switch -9223372036854775808 {
    case 4294967295:
        var v7 uint32 = 0
        _ = v7
    default:
        var v8 uint16 = 0
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS1{F0: -1, F1: 69}
        _ = fi9
        _ = v10
    }
    _ = v6
    return 128
}

func main() {
    v1 := []string{("a" + "a")}
    var v2 float32 = (3.14159 - (float32(1e10) / 1.0))
    v3 := []int{-888, -679, -475}
    v4 := make(map[int]int)
    v5 := make(map[int]int)
    for fi11 := 0; fi11 < 11; fi11++ {
        if (((v2 + float32(1.5)) + 1e10) == (0.1 / -2.25)) {
            var v12 int32 = 2147483647
            _ = v12
        } else {
            v2 -= v2
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi11
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        if (!false) {
            v14 := GenS1{F0: 9223372036854775807, F1: 89}
            _ = v14
        }
        v15 := GenS2{F0: 0}
        fmt.Println(fi13)
        _ = fi13
        _ = v15
    }
    v16 := GenS3{F0: 255, F1: -32768, F2: 32767}
    var v17 uint64 = 0
    for fi18 := 0; fi18 < 6; fi18++ {
        switch (fi18 << ((255) & 63)) {
        case 1:
            var v19 float64 = (((float64(v2) - 0.0) * ((0.1 + 3.14159) / (0.0 + (1.5 + float64(v2))))) * 0.1)
            _ = v19
        }
        v20 := GenS2{F0: 32767}
        _ = fi18
        _ = v20
    }
    v4[127] = 128
    switch (("foo bar" + "") + "foo bar") {
    case "foo bar":
        v2 = float32(-2.25)
    case "hello":
        switch (("a" + "\n\t") + "hello") {
        case "":
            switch "\n\t" {
            case "":
                v17 = (v17 % ((v17) | 1))
            case "\n\t":
                var v21 uint32 = 4294967295
                _ = v21
            case "hello":
                for fi22 := 0; fi22 < 7; fi22++ {
                    for fi23 := 0; fi23 < 14; fi23++ {
                        fmt.Println(v2, v17, fi22)
                        v2 -= 1e10
                        _ = fi23
                    }
                    _ = fi22
                }
            default:
                v2 += 3.14159
            }
        default:
            v24 := GenS4{F0: 4294967295, F1: 4294967295}
            _ = v24
        }
    case "a":
        v2 = float32((3.14159 / 3.14159))
    default:
        v25 := v5[42]
        _ = v25
    }
    if ((v17 ^ (v17 & (v17 - (v17 | v17)))) == (^(v17))) {
        v17 += (v17 & v17)
    }
    var v26 uint = 255
    var v27 bool = (1 <= 6)
    switch ("\n\t" + "foo bar") {
    case "hello":
        var v28 int = int(v17)
        _ = v28
    case "foo bar":
        v2 -= v2
    }
    v2 -= v2
    v29 := GenS3{F0: 42, F1: int16(v17), F2: 831}
    for fi30 := 0; fi30 < 19; fi30++ {
        fmt.Println(v26, v27)
        var v31 uint32 = 65535
        _ = fi30
        _ = v31
    }
    switch "foo bar" {
    case "":
        v27 = (!v27)
    case "\n\t":
        _ = len(v3)
    }
    if true {
        _ = len(v1)
        v26 = (^(v26))
        v32 := GenS5{F0: -18, F1: 42, F2: uint(v17)}
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
