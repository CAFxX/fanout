package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int; F1 int; F2 uint }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint64) bool {
    var v4 float64 = 3.14159
    v5 := GenS1{F0: 2, F1: 1}
    var v6 uint32 = 42
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return (!true)
}
func genFn1(p0 int16, p1 uint64) int8 {
    var v7 int64 = -9223372036854775808
    v7 += v7
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 250
    var v3 bool = true
    v8 := GenS1{F0: 65535, F1: -128}
    fmt.Println(v3, v3)
    fmt.Println(v3)
    if (v2 != (v2 & (v2 ^ (v2 | v2)))) {
        fmt.Println(v2, v2)
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := v1[415]
            var v11 uint64 = 4294967295
            if (-97 != -110) {
                v2 += 0
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 bool = (1000 != 0)
    if (!(-871 <= 42)) {
        v13 := []int{-441, 128, 127, -9223372036854775808}
        for fi14 := 0; fi14 < 1; fi14++ {
            if ("hello" == "\n\t") {
                v15 := GenS1{F0: 558, F1: 99}
                v16 := GenS1{F0: 327, F1: -1}
                fmt.Println(v2, v12, v2)
                _ = v15
                _ = v16
            }
            _ = fi14
        }
        var v17 string = "a"
        _ = v13
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        switch ("hello" + "hello") {
        case "\n\t":
            _ = len(v1)
        case "":
            switch "hello" {
            case "\n\t":
                v19 := GenS2{F0: fi18, F1: fi18, F2: 65535}
                _ = v19
            }
        case "foo bar":
            if (1e10 < (((0.1 / 0.1) * 0.1) * 0.1)) {
                fmt.Println(fi18, fi18)
                if (float64((0.0 + 1e10)) < 1e10) {
                    v12 = v12
                    fmt.Println(v2)
                }
            }
        default:
            for fi20 := 0; fi20 < 2; fi20++ {
                switch 1 {
                case 128:
                    v21 := GenS1{F0: 105, F1: genFn1(1000, 42)}
                    _ = v21
                case 1:
                    var v22 int32 = 675
                    _ = v22
                case -1:
                    v2 += v2
                }
                v3 = v3
                _ = fi20
            }
        }
        _ = fi18
    }
    v23 := GenS2{F0: 32767, F1: -1, F2: 255}
    var v24 uint8 = v2
    var v25 bool = (!false)
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v24)
    case "a":
        switch 768 {
        case -2147483648:
            if true {
                fmt.Println(v12, v25)
                switch 32767 {
                case -1:
                    if v12 {
                        v26 := GenS3{F0: -2147483648}
                        _ = v26
                    } else {
                        fmt.Println(v25)
                    }
                default:
                    var v27 uint64 = 386
                    _ = v27
                }
                var v28 int32 = 42
                _ = v28
            } else {
                fmt.Println(v12, v24, v12)
            }
        case -9223372036854775808:
            for fi29 := 0; fi29 < 12; fi29++ {
                v30 := GenS1{F0: 255, F1: genFn1(158, 1)}
                fmt.Println(v25, v25, v24)
                var v31 int16 = -32768
                _ = fi29
                _ = v30
                _ = v31
            }
        }
    default:
        var v32 int16 = 718
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v8
}
