package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint16; F2 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint }
func genFn0(p0 int, p1 int, p2 uint16) int8 {
    v7 := GenS1{F0: 645}
    var v8 int8 = 42
    var v9 float32 = (3.14159 / (-2.25 / -2.25))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return v8
}

func main() {
    v1 := []uint32{1, 4294967295, 4294967295}
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    v5 := []uint32{4294967295, 42, 467, 2147483648}
    var v6 int16 = -719
    switch ("" + "a") {
    case "hello":
        var v10 bool = false
        _ = v10
    }
    v11 := GenS1{F0: -9223372036854775808}
    v12 := GenS2{F0: ("foo bar" + "foo bar"), F1: 42, F2: 18446744073709551615}
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            switch ("foo bar" + "foo bar") {
            case "a":
                if true {
                    v4["foo bar"] = 2147483647
                } else {
                    _ = len(v1)
                    delete(v4, "foo bar")
                    if false {
                        v6 = 0
                    }
                }
            case "foo bar":
                if (128 <= 42) {
                    v15 := v4[("hello" + "hello")]
                    _ = v15
                }
            default:
                var v16 string = "foo bar"
                _ = v16
            }
            var v17 float32 = (1e10 * 0.0)
            if true {
                v6 += 1
                fmt.Println(v6, v17)
                v18 := GenS3{F0: 619}
                _ = v18
            } else {
                var v19 bool = true
                var v20 uint8 = 128
                if (587 >= 255) {
                    v5 = append(v5, 82)
                    switch (((("hello" + (("hello" + "hello") + ("foo bar" + "a"))) + ("\n\t" + ("\n\t" + "hello"))) + ("foo bar" + "a")) + "foo bar") {
                    case "hello":
                        var v21 uint8 = 208
                        if (!(!(0 == 562))) {
                            switch (fi13 & 32768) {
                            case -75:
                                v6 -= (^(v6))
                            case -355:
                                switch (fi14 ^ (fi13 << ((2) & 63))) {
                                case -9223372036854775808:
                                    var v22 int64 = 42
                                    _ = v22
                                case 127:
                                    var v23 uint8 = v20
                                    _ = v23
                                }
                            }
                            var v24 uint32 = 2
                            _ = v24
                        } else {
                            var v25 uint = 42
                            v20 -= v20
                            var v26 int8 = 42
                            _ = v25
                            _ = v26
                        }
                        _ = v21
                    default:
                        var v27 int32 = 2147483647
                        _ = v27
                    }
                    v28 := GenS2{F0: "hello", F1: 255, F2: 255}
                    _ = v28
                }
                _ = v19
                _ = v20
            }
            _ = fi14
            _ = v17
        }
        _ = fi13
    }
    switch ("\n\t" + "hello") {
    case "\n\t":
        var v29 float64 = ((0.0 + (3.14159 - float64(float32(-2.25)))) - (0.1 + 0.1))
        _ = v29
    case "":
        var v30 uint16 = 2
        _ = v30
    case "foo bar":
        switch "hello" {
        case "foo bar":
            v6 = -688
        case "hello":
            v31 := GenS4{F0: v6, F1: 100, F2: 128}
            _ = v31
        default:
            switch -1 {
            case -146:
                var v32 int = 2147483647
                _ = v32
            case 32768:
                v6 -= (-(v6))
            default:
                v33 := GenS1{F0: -9223372036854775808}
                _ = v33
            }
        }
    default:
        var v34 int32 = 199
        _ = v34
    }
    v6 = (v6 | (v6 / ((0) | 1)))
    v3[("\n\t" + "")] = (("hello" + "") + (("hello" + "foo bar") + ("foo bar" + ("" + ("hello" + "foo bar")))))
    v6 = 42
    if len(v5) > 0 {
        v5[0] = 2147483648
    }
    v35 := GenS5{F0: 255}
    if (322 <= -9223372036854775808) {
        v6 += v6
        v6 -= v6
        var v36 int = 32767
        _ = v36
    }
    v6 -= (v6 & -32768)
    v37 := GenS4{F0: v6, F1: 3, F2: 72}
    switch ("" + "a") {
    case "":
        v6 -= v6
    case "foo bar":
        var v38 uint64 = 473
        _ = v38
    default:
        delete(v2, 2)
    }
    fmt.Println(v6)
    v39 := GenS5{F0: 148}
    for fi40 := 0; fi40 < 12; fi40++ {
        var v41 uint32 = 992
        v41 = 0
        var v42 uint8 = 2
        _ = fi40
        _ = v41
        _ = v42
    }
    v43 := GenS2{F0: ("a" + "hello"), F1: 42, F2: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v39
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
