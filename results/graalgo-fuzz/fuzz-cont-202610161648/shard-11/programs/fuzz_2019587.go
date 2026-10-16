package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int8; F1 bool; F2 bool }
type GenS4 struct { F0 string; F1 uint32; F2 int64 }
func genFn0(p0 int16, p1 int16) uint16 {
    p1 += p1
    p0 -= (^(p1))
    p0 += p0
    _ = p0
    _ = p1
    return 880
}
func genFn1(p0 uint32) int32 {
    for fi7 := 0; fi7 < 20; fi7++ {
        p0 -= p0
        v8 := GenS1{F0: 727, F1: 122}
        _ = fi7
        _ = v8
    }
    var v9 uint8 = uint8(p0)
    fmt.Println(v9)
    v9 += (v9 - (v9 >> ((4294967295) & 63)))
    var v10 uint16 = 333
    _ = p0
    _ = v10
    _ = v9
    return 2147483647
}

func main() {
    var v1 float64 = ((1.5 * 1e10) / 1e10)
    var v2 uint16 = 283
    var v3 int64 = -83
    var v4 int = -195
    var v5 float64 = (1e10 / -2.25)
    var v6 int = v4
    v11 := GenS2{F0: 1, F1: 4294967295}
    v12 := []int{(v6 + -656), 32767, (v6 & (v4 % (((v6 | 976)) | 1))), (-(v4))}
    var v13 int = (v6 | -1)
    switch "a" {
    case "\n\t":
        v1 -= v5
    case "foo bar":
        if true {
            if false {
                _ = len(v12)
            } else {
                fmt.Println(v5)
                var v14 int16 = 76
                _ = v14
            }
        } else {
            if (!true) {
                for fi15 := 0; fi15 < 15; fi15++ {
                    var v16 int16 = 42
                    var v17 int8 = 100
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                v6 -= (v13 + (^(v6)))
            }
            if false {
                if len(v12) > 0 {
                    v12[0] = -2147483648
                }
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            }
        }
    default:
        var v18 int8 = -83
        _ = v18
    }
    v6 = 32768
    if false {
        v5 = v5
        v6 = (v13 % (((v4 + (v13 << ((255) & 63)))) | 1))
        var v19 float64 = 0.0
        _ = v19
    } else {
        switch v4 {
        case -2147483648:
            var v20 uint8 = 2
            _ = v20
        case 4294967295:
            var v21 float32 = 0.0
            _ = v21
        case 7:
            if len(v12) > 0 {
                v12[0] = (v13 - (v6 & -9223372036854775808))
            }
        default:
            if (true && true) {
                switch "foo bar" {
                case "a":
                    v3 += (v3 % ((930) | 1))
                default:
                    for fi22 := 0; fi22 < 13; fi22++ {
                        if (true && true) {
                            for fi23 := 0; fi23 < 7; fi23++ {
                                switch "" {
                                case "foo bar":
                                    switch 127 {
                                    case 4294967295:
                                        fmt.Println(v3, v5)
                                    }
                                case "\n\t":
                                    switch "hello" {
                                    case "a":
                                        var v24 float64 = v1
                                        _ = v24
                                    case "\n\t":
                                        var v25 int16 = 309
                                        _ = v25
                                    case "hello":
                                        v2 = (v2 << ((18446744073709551615) & 63))
                                    }
                                case "a":
                                    fmt.Println(v13)
                                }
                                _ = len(v12)
                                if false {
                                    v26 := GenS1{F0: uint(v2), F1: 100}
                                    var v27 float32 = 1.5
                                    _ = v26
                                    _ = v27
                                } else {
                                    fmt.Println(v6)
                                }
                                _ = fi23
                            }
                            if ("a" > "foo bar") {
                                if true {
                                    for fi28 := 0; fi28 < 14; fi28++ {
                                        if (!(((1e10 * (1.5 / (1.5 / (1.5 + 1e10)))) / 1.0) != (3.14159 / (3.14159 / 1.0)))) {
                                            _ = len(v12)
                                            var v29 uint8 = 177
                                            v5 = ((float64(0.1) * v1) - v5)
                                            _ = v29
                                        }
                                        var v30 uint = 18446744073709551615
                                        _ = fi28
                                        _ = v30
                                    }
                                }
                            }
                            var v31 float64 = (((((v5 * ((float64(1e10) * float64(float32(v5))) / v1)) - v5) * (v5 / ((v1 - v1) / float64(1.5)))) - float64(1.5)) * (v1 - ((((0.1 + 0.0) - (((v1 + v5) / (v5 / ((v5 / v1) * 1.5))) * (v5 - (v5 * float64(((1e10 + (0.0 / (-2.25 - float32((v1 / 0.1))))) - 0.1)))))) * v1) + (float64(1e10) + v5))))
                            _ = v31
                        }
                        v12 = append(v12, v4)
                        v2 = v2
                        _ = fi22
                    }
                }
                v6 = (-(v13))
                fmt.Println(v3)
            }
        }
        if false {
            var v32 string = (("\n\t" + "foo bar") + "hello")
            if len(v12) > 0 {
                v12[0] = v4
            }
            v2 = (v2 * genFn0(0, -544))
            _ = v32
        } else {
            v33 := GenS3{F0: -128, F1: false, F2: false}
            v6 = (^(v13))
            _ = v33
        }
        v4 += 9223372036854775807
    }
    if len(v12) > 0 {
        v12[0] = -724
    }
    switch 42 {
    case -9223372036854775808:
        if true {
            var v34 string = "foo bar"
            v13 += (v6 & (v4 % ((v6) | 1)))
            _ = v34
        }
    case 9223372036854775807:
        var v35 int = v6
        _ = v35
    case 128:
        v13 += 32768
    default:
        if len(v12) > 0 {
            v12[0] = (v13 & 4294967295)
        }
    }
    if true {
        for fi36 := 0; fi36 < 13; fi36++ {
            switch fi36 {
            case 32768:
                v37 := GenS3{F0: -73, F1: (!true), F2: (2 < 2)}
                _ = v37
            case 4294967295:
                for fi38 := 0; fi38 < 19; fi38++ {
                    var v39 uint = 18446744073709551615
                    _ = fi38
                    _ = v39
                }
            case 215:
                v5 = 0.0
            }
            v4 = (^(v6))
            v6 -= (v6 - 2147483647)
            _ = fi36
        }
        for fi40 := 0; fi40 < 14; fi40++ {
            if (!(!false)) {
                switch "" {
                case "":
                    switch "hello" {
                    case "a":
                        v3 += (v3 + (^(v3)))
                    default:
                        v41 := GenS2{F0: 676, F1: (fi40 & v4)}
                        _ = v41
                    }
                case "\n\t":
                    v42 := GenS3{F0: -48, F1: (v4 > 32767), F2: (255 != 0)}
                    _ = v42
                case "foo bar":
                    v5 = float64(float32(1e10))
                }
            }
            v43 := GenS4{F0: "\n\t", F1: 4294967295, F2: (v3 / ((42) | 1))}
            _ = fi40
            _ = v43
        }
    }
    var v44 uint32 = 2147483648
    v5 -= v1
    var v45 int8 = -1
    v4 -= v6
    for fi46 := 0; fi46 < 1; fi46++ {
        var v47 int16 = 0
        _ = fi46
        _ = v47
    }
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
