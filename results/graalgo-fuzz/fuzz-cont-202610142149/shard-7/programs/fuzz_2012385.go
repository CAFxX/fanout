package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }
func genFn0(p0 int, p1 uint) uint16 {
    var v6 int64 = 1
    p0 -= p0
    _ = p0
    _ = p1
    _ = v6
    return 255
}
func genFn1(p0 uint16, p1 uint32, p2 int) uint32 {
    var v7 int32 = -51
    v8 := []int{p2}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        p1 = p1
        _ = fi9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        p2 += fi10
        p1 = (p1 % (((^(p1))) | 1))
        p0 -= p0
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p1 % (((^(p1))) | 1))
}

func main() {
    var v1 uint = 0
    var v2 string = ""
    v3 := []uint32{42, 644}
    v4 := []uint32{58}
    v5 := make(map[string]string)
    if ((v2 > ((v2 + ("\n\t" + v2)) + "foo bar")) && (821 != 255)) {
        v2 = ("" + (("hello" + (v2 + "a")) + v2))
        var v11 int8 = -72
        _ = v11
    } else {
        for fi12 := 0; fi12 < 15; fi12++ {
            switch v2 {
            case "foo bar":
                var v13 float64 = -2.25
                _ = v13
            default:
                v2 = "\n\t"
            }
            switch fi12 {
            case -9223372036854775808:
                v2 = v2
            case 0:
                v1 = (v1 | (v1 + (v1 * 65535)))
            default:
                fmt.Println(fi12, fi12, v1)
            }
            _ = fi12
        }
    }
    var v14 uint = (^(v1))
    switch "" {
    case "":
        v15 := GenS1{F0: 2147483648, F1: false, F2: v2}
        _ = v15
    case "a":
        switch v2 {
        case "\n\t":
            if (!false) {
                fmt.Println(v14, v1)
            } else {
                var v16 float64 = (3.14159 * 0.1)
                _ = len(v4)
                switch "hello" {
                case "a":
                    v16 += ((v16 - float64((((1e10 * float32(float64((0.0 - ((1.5 * 0.1) - (-2.25 - ((-2.25 + 1.5) - -2.25))))))) - float32((v16 - v16))) / 1.5))) / v16)
                case "":
                    v2 = "foo bar"
                case "hello":
                    if (!(-9223372036854775808 <= 536)) {
                        var v17 int64 = 42
                        _ = v17
                    } else {
                        fmt.Println(v14, v14, v14)
                        switch v2 {
                        case "foo bar":
                            v14 -= v1
                        default:
                            var v18 uint32 = 989
                            _ = v18
                        }
                    }
                default:
                    for fi19 := 0; fi19 < 6; fi19++ {
                        if (((((float32(1.5) / (((0.0 + ((0.0 + ((1e10 * 3.14159) / -2.25)) / float32(1e10))) - 1.5) * 0.1)) + 1.5) + 0.0) > 0.1) || true) {
                            _ = len(v4)
                            var v20 bool = false
                            var v21 int32 = 1
                            _ = v20
                            _ = v21
                        } else {
                            if false {
                                switch 786 {
                                case -209:
                                    v22 := GenS1{F0: 0, F1: true, F2: "a"}
                                    _ = v22
                                case 2147483647:
                                    if (!((127 > 0) || true)) {
                                        v1 = v1
                                    }
                                default:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                if (1 < int16(v14)) {
                                    var v23 uint8 = 1
                                    _ = v23
                                }
                                for fi24 := 0; fi24 < 16; fi24++ {
                                    fmt.Println(v16, v14)
                                    v16 += v16
                                    _ = fi24
                                }
                            }
                            _ = len(v5)
                        }
                        v25 := GenS1{F0: 1, F1: false, F2: (("hello" + v2) + (v2 + v2))}
                        for fi26 := 0; fi26 < 6; fi26++ {
                            if false {
                                v14 -= v1
                                var v27 int32 = 2147483647
                                _ = v27
                            } else {
                                v28 := GenS1{F0: 313, F1: (!true), F2: ((v2 + "a") + v2)}
                                var v29 uint = (v14 * (^(v14)))
                                _ = v28
                                _ = v29
                            }
                            _ = fi26
                        }
                        _ = fi19
                        _ = v25
                    }
                }
                _ = v16
            }
        case "":
            var v30 float64 = (3.14159 + (3.14159 - (-2.25 - 0.1)))
            _ = v30
        }
    case "hello":
        v31 := GenS2{F0: 42}
        _ = v31
    }
    v14 += 0
    fmt.Println(v1, v2, v1)
    fmt.Println(v1, v14, v2)
    for fi32 := 0; fi32 < 19; fi32++ {
        v33 := GenS1{F0: 2, F1: (18446744073709551615 <= 2), F2: v2}
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        v2 = ""
        v2 = (v2 + ((v2 + v2) + v2))
        _ = fi34
    }
    if (!false) {
        var v35 float64 = (((0.0 - 0.0) + 3.14159) / 1.0)
        _ = v35
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        var v37 bool = ("a" != "foo bar")
        var v38 int16 = 127
        v37 = (2 <= 255)
        _ = fi36
        _ = v37
        _ = v38
    }
    v14 -= (v14 ^ v1)
    var v39 int32 = -682
    v40 := GenS3{F0: genFn1(genFn0(-2147483648, 255), genFn1(255, 65535, 0), -132)}
    v41 := GenS4{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
