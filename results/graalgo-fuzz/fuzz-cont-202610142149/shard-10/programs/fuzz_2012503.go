package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int }
type GenS2 struct { F0 string; F1 int16; F2 uint }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]bool)
    var v3 int16 = 1000
    var v4 float64 = ((0.1 - 1.5) - float64((0.0 + -2.25)))
    v5 := []uint32{2147483648, 790, 0}
    if ((true && (!(!true))) || false) {
        switch "" {
        case "\n\t":
            fmt.Println(v3, v4, v1)
        case "":
            for fi6 := 0; fi6 < 14; fi6++ {
                switch "foo bar" {
                case "foo bar":
                    v3 = v3
                case "hello":
                    v7 := GenS1{F0: "", F1: 18446744073709551615, F2: (^(fi6))}
                    _ = v7
                default:
                    switch fi6 {
                    case -555:
                        if (18446744073709551615 > 817) {
                            fmt.Println(v3)
                            var v8 uint16 = 528
                            for fi9 := 0; fi9 < 7; fi9++ {
                                var v10 int8 = 107
                                v10 += -52
                                fmt.Println(v3)
                                _ = fi9
                                _ = v10
                            }
                            _ = v8
                        }
                    }
                }
                var v11 string = "\n\t"
                _ = fi6
                _ = v11
            }
        default:
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v12 := GenS2{F0: "foo bar", F1: v3, F2: uint(v1)}
                _ = v12
            }
        }
    } else {
        v4 = ((0.1 * v4) - v4)
        v3 += v3
        v13 := GenS3{F0: 443}
        _ = v13
    }
    if true {
        switch "foo bar" {
        case "\n\t":
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 uint32 = 1
                _ = fi14
                _ = v15
            }
        }
        if (128 != 584) {
            var v16 uint = 143
            _ = v16
        } else {
            fmt.Println(v3, v4)
            var v17 uint = 1
            var v18 bool = false
            _ = v17
            _ = v18
        }
    } else {
        switch 9223372036854775807 {
        case 42:
            v1 -= (v1 ^ 127)
        case -2147483648:
            var v19 bool = true
            _ = v19
        default:
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
        }
        v20 := v2["foo bar"]
        _ = v20
    }
    v1 -= 2
    v21 := GenS1{F0: "a", F1: 23, F2: 2147483647}
    fmt.Println(v1)
    var v22 uint = 0
    if (141 == 2147483647) {
        v2["hello"] = true
        switch 865 {
        case 32768:
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 int32 = int32(fi23)
                switch "hello" {
                case "hello":
                    var v25 int = 32767
                    _ = v25
                }
                _ = fi23
                _ = v24
            }
        default:
            var v26 uint16 = 316
            _ = v26
        }
        var v27 uint32 = 2
        _ = v27
    }
    v28 := GenS2{F0: "hello", F1: (v3 << ((18446744073709551615) & 63)), F2: 255}
    for fi29 := 0; fi29 < 14; fi29++ {
        v22 = (v22 % ((2) | 1))
        if ((fi29 << (((v22 ^ (v22 | 943))) & 63)) < fi29) {
            delete(v2, "hello")
            v30 := GenS3{F0: (v22 * (^(v22)))}
            _ = v30
        }
        _ = fi29
    }
    v3 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
