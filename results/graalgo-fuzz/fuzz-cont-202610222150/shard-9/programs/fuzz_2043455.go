package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int; F1 int }

func main() {
    var v1 float64 = ((1e10 + float64((3.14159 - (((1e10 * (1.5 / 1e10)) - ((((0.1 + (1e10 - -2.25)) / (0.0 - 3.14159)) - 0.0) + (-2.25 / 1.0))) * 0.0)))) * 0.1)
    v2 := make(map[string]int)
    var v3 uint16 = 0
    var v4 int16 = -47
    var v5 int = -1
    switch "" {
    case "\n\t":
        fmt.Println(v3, v3, v5)
    case "foo bar":
        v4 -= 0
    }
    v5 = (v5 + -9223372036854775808)
    fmt.Println(v3)
    fmt.Println(v5, v1)
    if (!(v4 > v4)) {
        var v6 int = v5
        v1 -= 3.14159
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := GenS1{F0: 2147483647, F1: 738, F2: 850}
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        switch ("foo bar" + "") {
        case "foo bar":
            v5 = 2147483647
        default:
            if (!false) {
                if ("" > "foo bar") {
                    var v9 bool = (false || true)
                    var v10 float64 = 1e10
                    v3 = (v3 * (v3 % (((v3 ^ 42)) | 1)))
                    _ = v10
                    _ = v9
                }
                if (!false) {
                    for fi11 := 0; fi11 < 2; fi11++ {
                        var v12 uint32 = 4294967295
                        _ = fi11
                        _ = v12
                    }
                    if false {
                        fmt.Println(v3, v5, v5)
                        v5 += (v5 % ((v5) | 1))
                        if (true && ((true || false) && false)) {
                            v3 -= v3
                            fmt.Println(v5, v4, v3)
                            fmt.Println(v4, v4)
                        } else {
                            if (65535 > 2) {
                                var v13 string = ("foo bar" + "foo bar")
                                switch "a" {
                                case "\n\t":
                                    v1 += v1
                                case "hello":
                                    var v14 int16 = v4
                                    _ = v14
                                case "a":
                                    if false {
                                        _ = len(v2)
                                    } else {
                                        var v15 float64 = 1.5
                                        _ = v15
                                    }
                                }
                                _ = v13
                            }
                            var v16 int32 = -203
                            _ = v16
                        }
                    }
                } else {
                    for fi17 := 0; fi17 < 3; fi17++ {
                        v3 += v3
                        _ = len(v2)
                        v18 := GenS1{F0: -1, F1: 779, F2: 1}
                        _ = fi17
                        _ = v18
                    }
                    switch v5 {
                    case 692:
                        v1 = 0.0
                    case 435:
                        var v19 uint = 2
                        _ = v19
                    default:
                        var v20 uint8 = 92
                        _ = v20
                    }
                    fmt.Println(v1, v4, v3)
                }
            }
        }
    }
    v21 := v2[("a" + "a")]
    v22 := GenS1{F0: -1, F1: -2147483648, F2: 2}
    for fi23 := 0; fi23 < 9; fi23++ {
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 uint64 = 1
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    var v26 float32 = (1e10 / -2.25)
    v27 := GenS2{F0: v3, F1: 130, F2: v3}
    v4 -= (v4 % (((-(v4))) | 1))
    v4 -= (v4 ^ (v4 % (((v4 ^ v4)) | 1)))
    v28 := GenS3{F0: v21, F1: (^(v21))}
    for fi29 := 0; fi29 < 14; fi29++ {
        var v30 string = "foo bar"
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
