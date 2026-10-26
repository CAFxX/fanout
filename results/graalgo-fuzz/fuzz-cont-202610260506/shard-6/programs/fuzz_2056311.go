package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int8 }
func genFn0() int16 {
    switch -1 {
    case 32768:
        for fi4 := 0; fi4 < 11; fi4++ {
            if false {
                for fi5 := 0; fi5 < 18; fi5++ {
                    var v6 uint16 = 1
                    fmt.Println(v6)
                    v7 := GenS1{F0: (-2.25 <= 0.1), F1: 1, F2: -66}
                    _ = fi5
                    _ = v6
                    _ = v7
                }
            } else {
                var v8 uint16 = 42
                if true {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        v10 := []int{(^(fi4)), fi4, 4294967295, (fi9 | fi4)}
                        switch (fi9 % (((fi4 * (fi4 & fi9))) | 1)) {
                        case -450:
                            var v11 int64 = -1
                            _ = v11
                        case -29:
                            v8 += (^(v8))
                        default:
                            if true {
                                var v12 uint32 = 42
                                switch ("\n\t" + ("hello" + ("\n\t" + "\n\t"))) {
                                case "foo bar":
                                    var v13 float32 = (3.14159 - 0.0)
                                    _ = v13
                                default:
                                    v8 -= (v8 & (^(v8)))
                                }
                                _ = v12
                            }
                        }
                        if ((!(v8 == 42)) && false) {
                            switch (fi9 / ((128) | 1)) {
                            case 732:
                                v8 = (v8 - v8)
                            case 32768:
                                v8 = 42
                            }
                            switch (("a" + (("foo bar" + "a") + "")) + ("" + (("\n\t" + "hello") + ("" + "foo bar")))) {
                            case "a":
                                if (4294967295 == 18446744073709551615) {
                                    v8 += v8
                                }
                            case "\n\t":
                                var v14 int16 = -32768
                                _ = v14
                            }
                            v8 = (v8 + (v8 / ((v8) | 1)))
                            v8 = v8
                        } else {
                            if (!true) {
                                switch ((("hello" + "foo bar") + ("foo bar" + "")) + "hello") {
                                case "hello":
                                    var v15 float64 = 3.14159
                                    _ = v15
                                case "\n\t":
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                }
                                v8 += (v8 / ((v8) | 1))
                            }
                            fmt.Println(fi9, fi4)
                        }
                        _ = fi9
                        _ = v10
                    }
                    if true {
                        v8 -= v8
                        v8 = v8
                    }
                }
                _ = v8
            }
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 uint32 = 1
                var v18 int8 = -1
                v19 := []int{-174, 539, (fi16 >> ((uint(v17)) & 63))}
                switch "a" {
                case "a":
                    var v20 uint64 = 2
                    _ = v20
                case "foo bar":
                    v18 += (v18 ^ 0)
                case "\n\t":
                    switch ("" + "a") {
                    case "foo bar":
                        var v21 int64 = 42
                        _ = v21
                    case "":
                        v17 = 777
                    default:
                        var v22 int16 = -341
                        _ = v22
                    }
                }
                var v23 float64 = (1.5 * 0.0)
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
                _ = v23
            }
            _ = fi4
        }
    default:
        var v24 uint = 427
        _ = v24
    }
    switch ("\n\t" + "a") {
    case "":
        var v25 uint32 = 2147483647
        _ = v25
    case "\n\t":
        var v26 uint8 = 88
        _ = v26
    case "hello":
        v27 := []int{729, 32768, 128}
        _ = v27
    default:
        if (9223372036854775807 != -1) {
            var v28 int16 = 1
            var v29 float64 = (3.14159 / float64(3.14159))
            var v30 int64 = 9223372036854775807
            _ = v28
            _ = v29
            _ = v30
        }
    }
    return -1
}

func main() {
    v1 := []int{489, -1, 2147483647, 0}
    v2 := make(map[uint32]string)
    v3 := []int{128}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v31 int = 32768
    v31 = (^(v31))
    var v32 int64 = 241
    for fi33 := 0; fi33 < 4; fi33++ {
        var v34 uint16 = 65535
        for fi35 := 0; fi35 < 15; fi35++ {
            v31 += v31
            _ = fi35
        }
        fmt.Println(v34, v34, fi33)
        _ = fi33
        _ = v34
    }
    for fi36 := 0; fi36 < 17; fi36++ {
        fmt.Println(v32, v31)
        for fi37 := 0; fi37 < 1; fi37++ {
            v32 -= v32
            _ = fi37
        }
        _ = fi36
    }
    var v38 float32 = float32((0.0 + 1e10))
    fmt.Println(v38, v31, v31)
    var v39 int16 = 1000
    var v40 int16 = -32768
    v38 = v38
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v39
    _ = v40
}
