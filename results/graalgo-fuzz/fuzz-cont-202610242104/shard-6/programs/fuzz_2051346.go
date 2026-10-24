package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int8 }
type GenS2 struct { F0 int64; F1 int8; F2 bool }
type GenS3 struct { F0 uint; F1 uint; F2 int16 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 int32; F1 uint64; F2 int8 }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 uint32; F1 int8; F2 int }
func genFn0(p0 int16) bool {
    var v4 uint8 = 84
    if (!true) {
        switch "" {
        case "hello":
            fmt.Println(v4, p0)
        }
        var v5 int8 = 0
        _ = v5
    }
    v4 -= v4
    var v6 uint16 = 255
    _ = p0
    _ = v4
    _ = v6
    return true
}
func genFn1(p0 int16) uint {
    switch 981 {
    case 343:
        v7 := GenS1{F0: -9, F1: -2147483648, F2: 127}
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: 43, F1: -9223372036854775808, F2: 1}
        p0 = p0
        _ = fi8
        _ = v9
    }
    _ = p0
    return 42
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[uint32]uint32)
    v3 := []int{}
    if true {
        v1 = 2
        v10 := GenS2{F0: -914, F1: 127, F2: true}
        _ = v10
    } else {
        var v11 float32 = ((0.1 * 0.1) - 1.5)
        _ = v11
    }
    v12 := GenS3{F0: 65535, F1: 4294967295, F2: 1000}
    v13 := GenS4{F0: 2147483648, F1: 42}
    switch -1 {
    case -227:
        v1 -= (v1 >> ((18446744073709551615) & 63))
    case 32768:
        if true {
            var v14 uint = 4294967295
            if false {
                var v15 int16 = -1
                for fi16 := 0; fi16 < 10; fi16++ {
                    for fi17 := 0; fi17 < 16; fi17++ {
                        if ((v15 | 42) == v15) {
                            var v18 float64 = ((0.1 / 0.1) - 0.1)
                            _ = v18
                        }
                        var v19 uint8 = 2
                        _ = fi17
                        _ = v19
                    }
                    _ = fi16
                }
                _ = v15
            }
            for fi20 := 0; fi20 < 6; fi20++ {
                _ = len(v3)
                v21 := GenS5{F0: 637, F1: true}
                switch fi20 {
                case 571:
                    v22 := GenS6{F0: 2147483647, F1: 4294967295, F2: 100}
                    _ = v22
                case 859:
                    v14 += uint(v1)
                case 302:
                    for fi23 := 0; fi23 < 1; fi23++ {
                        v24 := GenS7{F0: "foo bar", F1: 362}
                        fmt.Println(v1, v1)
                        _ = fi23
                        _ = v24
                    }
                default:
                    v25 := GenS5{F0: 2147483647, F1: ((1.5 == -2.25) && (82 < 255))}
                    _ = v25
                }
                _ = fi20
                _ = v21
            }
            _ = v14
        }
    default:
        v3 = append(v3, 32768)
    }
    fmt.Println(v1)
    switch (("" + ("a" + ("foo bar" + "hello"))) + "\n\t") {
    case "hello":
        if len(v3) > 0 {
            v3[0] = -926
        }
    case "":
        _ = len(v2)
    case "\n\t":
        var v26 int32 = -2147483648
        _ = v26
    }
    if false {
        for fi27 := 0; fi27 < 12; fi27++ {
            v2[42] = 65535
            _ = fi27
        }
        v28 := GenS8{F0: 42, F1: 0, F2: 755}
        if (255 <= 42) {
            for fi29 := 0; fi29 < 6; fi29++ {
                v1 = v1
                _ = fi29
            }
            if false {
                var v30 uint16 = 42
                fmt.Println(v1)
                var v31 float32 = (float32((1e10 - 0.0)) / 1.0)
                _ = v30
                _ = v31
            }
            var v32 string = ("foo bar" + "foo bar")
            _ = v32
        }
        _ = v28
    }
    var v33 uint16 = (^(v1))
    v1 += (v33 ^ v1)
    fmt.Println(v33)
    fmt.Println(v33)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v33
}
