package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 uint32; F2 string }

func main() {
    var v1 uint16 = 885
    var v2 float32 = -2.25
    var v3 uint = 426
    var v4 float64 = (((((-2.25 / -2.25) - (0.1 * 3.14159)) + -2.25) + (((1e10 + 1.5) * 0.0) - -2.25)) + 3.14159)
    v3 = 255
    v5 := []int{}
    fmt.Println(v3)
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v6 string = "hello"
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            if (float32(v4) > ((float32(((v4 * (1.5 / v4)) / 0.1)) - v2) * 1.5)) {
                fmt.Println(fi7, v3)
            } else {
                v4 -= (v4 * v4)
                fmt.Println(v3, fi7, v1)
            }
            _ = fi8
        }
        _ = fi7
    }
    if (1000 >= 32767) {
        v9 := GenS1{F0: true}
        _ = v9
    } else {
        v2 += (((v2 + v2) - (v2 + v2)) - ((((0.0 * float32(3.14159)) + (0.1 + (1.5 * float32((1e10 - v4))))) - 1e10) + v2))
        var v10 string = ("a" + ("a" + "a"))
        _ = v10
    }
    if (2 > 2) {
        v1 = (v1 % ((v1) | 1))
        if (false || (370 > 4294967295)) {
            fmt.Println(v2)
            switch "foo bar" {
            case "\n\t":
                v4 += float64(((v2 / ((float32(3.14159) * (v2 * 0.1)) + v2)) + (v2 / 1.5)))
            case "hello":
                v11 := GenS2{F0: (v1 + 32767), F1: 127}
                _ = v11
            }
            switch "a" {
            case "\n\t":
                if false {
                    switch "" {
                    case "\n\t":
                        var v12 uint32 = 2
                        _ = v12
                    case "a":
                        if (!false) {
                            for fi13 := 0; fi13 < 4; fi13++ {
                                var v14 uint8 = 46
                                switch fi13 {
                                case 2147483647:
                                    switch 259 {
                                    case 4294967295:
                                        var v15 uint8 = v14
                                        _ = v15
                                    case 42:
                                        var v16 float32 = (float32(v4) * float32(-2.25))
                                        _ = v16
                                    case -961:
                                        for fi17 := 0; fi17 < 16; fi17++ {
                                            var v18 int32 = 9
                                            _ = fi17
                                            _ = v18
                                        }
                                    }
                                case -2147483648:
                                    switch "\n\t" {
                                    case "hello":
                                        if ((v1 << (((v3 << ((1) & 63))) & 63)) >= (v1 % ((v1) | 1))) {
                                            var v19 int8 = 0
                                            v20 := GenS2{F0: v1, F1: (v19 << ((uint(v1)) & 63))}
                                            _ = len(v5)
                                            _ = v19
                                            _ = v20
                                        }
                                    }
                                case -762:
                                    v5 = append(v5, (fi13 * fi13))
                                default:
                                    v14 = (^(v14))
                                }
                                _ = fi13
                                _ = v14
                            }
                            v21 := GenS1{F0: true}
                            _ = v21
                        }
                    case "":
                        if true {
                            v4 -= float64(v2)
                        } else {
                            var v22 uint8 = 2
                            if false {
                                v23 := GenS3{F0: 0}
                                var v24 int = 127
                                v24 -= 0
                                _ = v23
                                _ = v24
                            } else {
                                switch "" {
                                case "hello":
                                    switch -9223372036854775808 {
                                    case -112:
                                        v25 := GenS3{F0: 2}
                                        _ = v25
                                    }
                                case "\n\t":
                                    var v26 float32 = v2
                                    _ = v26
                                default:
                                    var v27 bool = (839 >= (^(v1)))
                                    _ = v27
                                }
                                _ = len(v5)
                                var v28 uint32 = 1
                                _ = v28
                            }
                            var v29 uint = 255
                            _ = v22
                            _ = v29
                        }
                    default:
                        v30 := GenS1{F0: (676 != 891)}
                        _ = v30
                    }
                    v31 := GenS4{F0: int16(v3), F1: 0}
                    _ = v31
                }
            case "hello":
                v32 := GenS5{F0: (false && true)}
                _ = v32
            case "":
                for fi33 := 0; fi33 < 16; fi33++ {
                    v2 = v2
                    _ = fi33
                }
            default:
                v3 = (v3 * v3)
            }
        }
    }
    var v34 string = "a"
    v4 += 3.14159
    var v35 int64 = 1
    for fi36 := 0; fi36 < 4; fi36++ {
        v4 += -2.25
        v37 := GenS6{F0: 416, F1: 195, F2: v34}
        _ = fi36
        _ = v37
    }
    switch ((v34 + "") + "hello") {
    case "a":
        var v38 float32 = v2
        _ = v38
    case "":
        var v39 uint32 = 429
        _ = v39
    default:
        var v40 string = v34
        _ = v40
    }
    if true {
        v1 = (v1 & (v1 << ((v3) & 63)))
        var v41 bool = (!false)
        v1 = v1
        _ = v41
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v3 += 552
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
