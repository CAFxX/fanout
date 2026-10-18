package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint16; F1 string; F2 int16 }
type GenS3 struct { F0 int64; F1 int }

func main() {
    v1 := []string{("\n\t" + ((((((("a" + "hello") + ("" + ("hello" + ("foo bar" + "a")))) + "") + "foo bar") + "\n\t") + "foo bar") + "")), ("" + ("\n\t" + "foo bar"))}
    v2 := make(map[uint32]bool)
    var v3 string = "a"
    v4 := []string{(v3 + (v3 + (v3 + v3))), v3}
    if true {
        v5 := GenS1{F0: (306 <= 958), F1: 128}
        _ = v5
    }
    switch 2147483647 {
    case 127:
        v3 = v3
    }
    v3 = (("hello" + "") + "\n\t")
    v3 = ""
    v6 := GenS1{F0: false, F1: 211}
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS2{F0: 65535, F1: "\n\t", F2: 1}
        for fi9 := 0; fi9 < 20; fi9++ {
            if false {
                v10 := GenS1{F0: (1000 > 42), F1: fi9}
                switch (fi9 | fi9) {
                case -9223372036854775808:
                    var v11 int8 = 42
                    _ = v11
                case 2147483647:
                    fmt.Println(v3, fi7, v3)
                default:
                    if (true || (0 <= -1)) {
                        fmt.Println(fi7, fi9)
                        v3 = ("\n\t" + ("" + v3))
                    }
                    if (253 >= 255) {
                        if false {
                            v3 = v3
                            var v12 int32 = 0
                            fmt.Println(fi7, v3, fi7)
                            _ = v12
                        } else {
                            if true {
                                fmt.Println(fi9, fi9, fi9)
                                switch "foo bar" {
                                case "hello":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                v3 = (v3 + v3)
                            }
                            v3 = v3
                            var v13 uint16 = 2
                            _ = v13
                        }
                        v3 = ("foo bar" + v3)
                        v3 = v3
                    } else {
                        fmt.Println(fi7)
                        var v14 uint64 = 2
                        v15 := GenS3{F0: -1, F1: (fi9 / (((fi9 ^ (fi7 + fi7))) | 1))}
                        _ = v14
                        _ = v15
                    }
                }
                _ = v10
            } else {
                v3 = (("a" + "") + (v3 + ""))
                v2[2] = (float32(float64((1.5 / 1e10))) <= 3.14159)
                switch (v3 + v3) {
                case "a":
                    var v16 float32 = ((1e10 - ((-2.25 / (0.0 - (((((0.1 / 1.0) + 0.1) / 1e10) * 0.1) - -2.25))) + 0.0)) + 0.1)
                    _ = v16
                case "hello":
                    var v17 bool = (18446744073709551615 <= 42)
                    _ = v17
                case "":
                    switch (fi7 / ((fi7) | 1)) {
                    case -360:
                        var v18 uint8 = 127
                        _ = v18
                    default:
                        fmt.Println(fi7, fi9, fi7)
                    }
                }
                v3 = (v3 + "hello")
            }
            v19 := GenS3{F0: -1, F1: (-(fi9))}
            switch "foo bar" {
            case "a":
                switch (fi9 * -2147483648) {
                case -861:
                    var v20 uint8 = 255
                    _ = v20
                case 32768:
                    var v21 uint64 = 4294967295
                    _ = v21
                case 583:
                    _ = len(v4)
                }
            case "\n\t":
                var v22 int = (^(fi9))
                _ = v22
            case "foo bar":
                fmt.Println(fi9, fi7)
            }
            _ = fi9
            _ = v19
        }
        _ = fi7
        _ = v8
    }
    switch 9223372036854775807 {
    case -2147483648:
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 int64 = -1
            _ = fi23
            _ = v24
        }
    case 127:
        var v25 bool = (false && true)
        _ = v25
    default:
        var v26 int = -2147483648
        _ = v26
    }
    v3 = "a"
    fmt.Println(v3)
    var v27 uint16 = 578
    var v28 float32 = (float32(0.0) - 3.14159)
    fmt.Println(v3, v27, v3)
    v27 = 42
    if ((("" + v3) + ((v3 + v3) + v3)) != ("foo bar" + "a")) {
        fmt.Println(v27, v3)
        if false {
            var v29 uint8 = 42
            fmt.Println(v27, v3)
            _ = len(v1)
            _ = v29
        }
    }
    fmt.Println(v28)
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v6
}
