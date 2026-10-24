package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint32; F1 int; F2 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32; F1 uint16; F2 int64 }
func genFn0() uint64 {
    var v6 uint = 68
    switch -359 {
    case 1:
        v7 := []int{-9223372036854775808}
        _ = v7
    }
    switch 770 {
    case 2147483647:
        if (1 <= 1) {
            for fi8 := 0; fi8 < 13; fi8++ {
                switch (fi8 / (((fi8 / ((fi8) | 1))) | 1)) {
                case -2147483648:
                    v6 -= 1
                case 4294967295:
                    v6 = (v6 - v6)
                default:
                    for fi9 := 0; fi9 < 7; fi9++ {
                        fmt.Println(v6, fi8, fi8)
                        v10 := []int{(fi9 >> ((4294967295) & 63)), -2147483648}
                        var v11 uint = (v6 & v6)
                        var v12 uint16 = 32767
                        _ = fi9
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                }
                v6 = (^(v6))
                v13 := GenS1{F0: 407, F1: uint8(v6), F2: 2147483647}
                _ = fi8
                _ = v13
            }
        }
    }
    fmt.Println(v6)
    if false {
        fmt.Println(v6)
        v14 := GenS2{F0: 0}
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 int8 = -125
            _ = fi15
            _ = v16
        }
        _ = v14
    } else {
        v6 -= 18446744073709551615
        v6 = v6
    }
    _ = v6
    return 2
}
func genFn1(p0 uint8) int32 {
    fmt.Println(p0)
    v17 := []int{2147483647, 128}
    if len(v17) > 0 {
        v17[0] = 2147483647
    }
    _ = p0
    _ = v17
    return 42
}

func main() {
    var v1 int32 = 0
    var v2 uint16 = 255
    var v3 int16 = 509
    v4 := []uint32{42, 182, 2147483647}
    v5 := make(map[int]bool)
    v18 := GenS3{F0: (v3 | (v3 - 1)), F1: "\n\t"}
    fmt.Println(v3, v1)
    switch "a" {
    case "foo bar":
        var v19 string = "a"
        _ = v19
    case "":
        switch (("a" + "a") + "foo bar") {
        case "":
            v3 = v3
        case "\n\t":
            fmt.Println(v2)
        default:
            if false {
                var v20 float64 = (((1.5 + 0.0) * ((((0.1 - float64(((3.14159 + 1.5) / 1e10))) * 0.1) - 3.14159) + 1e10)) / (1.5 / 0.1))
                _ = v20
            } else {
                var v21 int64 = 42
                var v22 uint8 = 149
                _ = len(v5)
                _ = v21
                _ = v22
            }
        }
    default:
        fmt.Println(v1)
    }
    v4 = append(v4, 620)
    v23 := GenS4{F0: 4294967295, F1: 2147483647, F2: ("foo bar" + "")}
    v3 -= (v3 / ((537) | 1))
    v2 = 0
    _ = len(v5)
    switch 2147483647 {
    case 975:
        if true {
            var v24 string = ("" + ("" + ("\n\t" + "a")))
            _ = v24
        }
    default:
        if (false || true) {
            v3 = v3
            var v25 int64 = 0
            _ = v25
        } else {
            switch ("" + (("hello" + "hello") + "foo bar")) {
            case "\n\t":
                var v26 int64 = 42
                _ = v26
            }
            switch ("a" + "a") {
            case "hello":
                v27 := GenS5{F0: -128}
                _ = v27
            case "":
                var v28 uint8 = 255
                _ = v28
            default:
                v3 -= 515
            }
        }
    }
    var v29 int16 = (v3 | v3)
    fmt.Println(v1)
    v30 := GenS6{F0: genFn1(0), F1: 476, F2: -578}
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
