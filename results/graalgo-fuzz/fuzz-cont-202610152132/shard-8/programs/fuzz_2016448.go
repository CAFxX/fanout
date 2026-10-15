package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int64; F2 string }

func main() {
    var v1 int8 = 100
    var v2 int64 = 0
    var v3 uint16 = 255
    var v4 int32 = -237
    v5 := []uint32{}
    v6 := make(map[string]string)
    switch "foo bar" {
    case "":
        var v7 int64 = (v2 << ((42) & 63))
        _ = v7
    case "hello":
        for fi8 := 0; fi8 < 1; fi8++ {
            fmt.Println(fi8, v1, v2)
            _ = fi8
        }
    }
    fmt.Println(v1)
    switch "" {
    case "a":
        switch "\n\t" {
        case "\n\t":
            v5 = append(v5, 42)
        case "":
            if true {
                v3 = v3
                if true {
                    var v9 float32 = 0.0
                    v2 = 42
                    v1 += (^(v1))
                    _ = v9
                } else {
                    v6["a"] = "a"
                    var v10 bool = false
                    v4 += int32(v3)
                    _ = v10
                }
            }
        case "foo bar":
            switch ("a" + ((("hello" + ("foo bar" + ("a" + "\n\t"))) + "hello") + "a")) {
            case "a":
                _ = len(v6)
            default:
                var v11 uint16 = v3
                _ = v11
            }
        default:
            for fi12 := 0; fi12 < 12; fi12++ {
                switch -2147483648 {
                case 373:
                    _ = len(v6)
                case -686:
                    v13 := GenS1{F0: 0, F1: -877}
                    _ = v13
                case 42:
                    for fi14 := 0; fi14 < 7; fi14++ {
                        var v15 int = int(v4)
                        _ = fi14
                        _ = v15
                    }
                default:
                    v3 = (^(v3))
                }
                v16 := GenS2{F0: 4294967295}
                _ = fi12
                _ = v16
            }
        }
    case "\n\t":
        fmt.Println(v3, v3)
    case "":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        var v17 int = -2147483648
        _ = v17
    }
    v5 = append(v5, 163)
    if ("" < "") {
        switch "" {
        case "\n\t":
            if (v1 > 100) {
                v18 := GenS3{F0: (v4 / (((v4 ^ (^(v4)))) | 1)), F1: v2, F2: (("hello" + ("\n\t" + ("hello" + "hello"))) + (("" + "foo bar") + "hello"))}
                var v19 int = 202
                fmt.Println(v19, v4, v1)
                _ = v18
                _ = v19
            }
        default:
            v3 = (v3 % ((v3) | 1))
        }
    } else {
        _ = len(v5)
    }
    fmt.Println(v4, v3)
    v1 = (v1 >> ((875) & 63))
    fmt.Println(v4, v3)
    fmt.Println(v3)
    switch -309 {
    case 0:
        for fi20 := 0; fi20 < 17; fi20++ {
            for fi21 := 0; fi21 < 18; fi21++ {
                var v22 string = (("hello" + "") + (((("" + (("\n\t" + (("hello" + ("foo bar" + "a")) + ("a" + ""))) + "foo bar")) + "\n\t") + "\n\t") + ("" + "\n\t")))
                _ = fi21
                _ = v22
            }
            var v23 int64 = int64(fi20)
            _ = fi20
            _ = v23
        }
    case 831:
        v2 -= v2
    default:
        v2 -= (v2 * 42)
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 int64 = (-(v2))
        _ = fi24
        _ = v25
    }
    v3 = v3
    if (((1.5 + ((1.5 * 3.14159) - 0.1)) / 1.0) <= -2.25) {
        switch -2147483648 {
        case -1:
            var v26 uint8 = 23
            _ = v26
        case 2147483647:
            var v27 uint64 = 0
            _ = v27
        case 9223372036854775807:
            var v28 int64 = v2
            _ = v28
        }
    } else {
        v1 += 100
        v29 := GenS3{F0: (v4 & -823), F1: (v2 - v2), F2: ((("foo bar" + "foo bar") + "\n\t") + "\n\t")}
        v1 -= (v1 | (v1 | (v1 / ((v1) | 1))))
        _ = v29
    }
    if (32767 >= int16(v2)) {
        if (227 > 1) {
            v4 += 0
        } else {
            if false {
                v3 = v3
                if false {
                    v3 += v3
                    for fi30 := 0; fi30 < 1; fi30++ {
                        _ = len(v6)
                        fmt.Println(fi30)
                        fmt.Println(fi30)
                        _ = fi30
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                v31 := GenS3{F0: v4, F1: 1, F2: "hello"}
                _ = v31
            }
        }
    }
    v1 = (v1 << ((2) & 63))
    if false {
        var v32 int32 = 0
        v33 := GenS2{F0: 39}
        _ = v32
        _ = v33
    }
    if (-988 != v4) {
        v2 += (v2 << ((2) & 63))
        switch "hello" {
        case "":
            v3 -= v3
        case "\n\t":
            var v34 uint = 63
            _ = v34
        case "foo bar":
            fmt.Println(v2)
        }
        v4 = 1
    } else {
        v35 := GenS1{F0: 119, F1: -1}
        v1 = v1
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
