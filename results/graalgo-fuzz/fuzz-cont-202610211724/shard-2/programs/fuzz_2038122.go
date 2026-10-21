package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int; F2 uint16 }
type GenS4 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 int16; F1 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 int64; F1 bool }
type GenS11 struct { F0 uint64 }
func genFn0(p0 int16, p1 int8, p2 uint16) uint {
    p0 = 0
    p0 = p0
    p0 += (p0 % ((-584) | 1))
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 uint, p1 uint16, p2 int8) int16 {
    var v7 float32 = (3.14159 * 3.14159)
    v8 := []int{9223372036854775807}
    var v9 float64 = (1e10 - 0.1)
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS1{F0: ("a" == "\n\t"), F1: 51, F2: (p2 | (p2 * p2))}
        _ = len(v8)
        var v12 float64 = v9
        _ = fi10
        _ = v11
        _ = v12
    }
    p0 = (p0 ^ (^(p0)))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 643
}

func main() {
    var v1 uint = 886
    var v2 int8 = -89
    v3 := make(map[uint32]bool)
    var v4 float64 = (1.5 / 3.14159)
    var v5 uint8 = 1
    v6 := make(map[string]bool)
    switch "\n\t" {
    case "a":
        switch 42 {
        case 128:
            v5 += v5
        }
    case "":
        if true {
            v13 := GenS1{F0: (32767 <= 418), F1: v5, F2: int8(v1)}
            fmt.Println(v1, v2, v4)
            _ = v13
        } else {
            var v14 bool = (false || (2147483647 >= -2147483648))
            _ = v14
        }
    case "foo bar":
        v15 := GenS1{F0: ((v2 ^ (v2 * v2)) >= (v2 & 49)), F1: (v5 + v5), F2: 127}
        _ = v15
    }
    if true {
        if false {
            v5 -= (^(v5))
            if (!true) {
                if true {
                    var v16 string = "\n\t"
                    _ = v16
                } else {
                    switch 128 {
                    case 64:
                        var v17 int16 = genFn1((v1 % ((0) | 1)), 0, (v2 % (((v2 - (-(v2)))) | 1)))
                        _ = v17
                    case -185:
                        v6[""] = false
                    case 32768:
                        switch -9223372036854775808 {
                        case 128:
                            switch ("a" + (("foo bar" + "hello") + ("hello" + "\n\t"))) {
                            case "\n\t":
                                fmt.Println(v4)
                            default:
                                if true {
                                    var v18 uint16 = 463
                                    for fi19 := 0; fi19 < 16; fi19++ {
                                        v20 := v3[0]
                                        var v21 uint16 = v18
                                        v22 := GenS2{F0: -2147483648}
                                        _ = fi19
                                        _ = v20
                                        _ = v21
                                        _ = v22
                                    }
                                    v23 := []int{32768, 127, 127, -9223372036854775808}
                                    _ = v18
                                    _ = v23
                                } else {
                                    v5 = 42
                                    var v24 int8 = -67
                                    v25 := GenS2{F0: -1}
                                    _ = v24
                                    _ = v25
                                }
                            }
                        default:
                            switch "hello" {
                            case "foo bar":
                                v1 += (v1 / (((v1 ^ (v1 & (v1 - (v1 + (^(v1))))))) | 1))
                            case "":
                                v26 := GenS3{F0: "", F1: 2147483647, F2: 621}
                                _ = v26
                            case "hello":
                                var v27 string = ("" + "a")
                                _ = v27
                            }
                        }
                    default:
                        v28 := GenS1{F0: (false && (4294967295 >= 65535)), F1: v5, F2: (v2 >> ((v1) & 63))}
                        _ = v28
                    }
                    _ = len(v6)
                    var v29 float32 = ((1.5 - (-2.25 - 0.0)) + 0.1)
                    _ = v29
                }
                v4 -= 3.14159
            }
        }
    }
    v30 := GenS4{F0: 18446744073709551615, F1: 1, F2: (^(v5))}
    fmt.Println(v4, v1)
    if (((0.1 / (3.14159 + 0.0)) >= 1e10) || (!(((-(v2)) >= v2) || (!true)))) {
        v2 += (v2 << (((v1 / ((105) | 1))) & 63))
        var v31 int16 = 470
        _ = v31
    } else {
        var v32 float32 = float32(v4)
        switch ("\n\t" + ("a" + "foo bar")) {
        case "":
            v33 := []int{-781, -324, 225, -862}
            _ = v33
        case "\n\t":
            if false {
                if (-1 < v2) {
                    v34 := GenS5{F0: 9223372036854775807, F1: 284}
                    _ = v34
                } else {
                    fmt.Println(v2, v32, v32)
                }
                v35 := GenS6{F0: -466, F1: genFn0(genFn1(v1, 255, v2), v2, 2)}
                v36 := []int{int(v2), 9223372036854775807}
                _ = v35
                _ = v36
            }
        }
        v37 := GenS7{F0: 233, F1: 42}
        _ = v32
        _ = v37
    }
    fmt.Println(v5, v4, v5)
    if (false || false) {
        for fi38 := 0; fi38 < 6; fi38++ {
            switch "foo bar" {
            case "a":
                if (1e10 < v4) {
                    for fi39 := 0; fi39 < 11; fi39++ {
                        v4 += -2.25
                        v4 -= (v4 * (float64(0.1) - v4))
                        _ = fi39
                    }
                    v40 := []int{fi38, fi38, 283}
                    v41 := GenS8{F0: (^(v1))}
                    _ = v40
                    _ = v41
                }
            case "":
                var v42 uint8 = (^(v5))
                _ = v42
            default:
                v5 -= 127
            }
            _ = fi38
        }
        var v43 uint64 = 2
        _ = v43
    } else {
        var v44 float32 = float32(v4)
        _ = v44
    }
    var v45 float32 = 0.0
    if (!true) {
        if true {
            var v46 int64 = -9223372036854775808
            _ = v46
        }
        var v47 int16 = 36
        _ = v47
    } else {
        if (false && (v4 == (v4 / float64(3.14159)))) {
            v1 += 255
            if (642 >= 32767) {
                v5 += v5
            }
        }
        for fi48 := 0; fi48 < 18; fi48++ {
            v49 := GenS3{F0: ("a" + "hello"), F1: (fi48 * (fi48 << (((^(v1))) & 63))), F2: 640}
            v50 := []int{fi48, fi48}
            v45 = 3.14159
            _ = fi48
            _ = v49
            _ = v50
        }
        for fi51 := 0; fi51 < 18; fi51++ {
            if true {
                switch ("foo bar" + "\n\t") {
                case "hello":
                    v1 += (^(v1))
                default:
                    v6["\n\t"] = true
                }
                v52 := []int{fi51, -2147483648, fi51}
                _ = v52
            } else {
                v53 := v3[65535]
                if (-235 < 0) {
                    if (18446744073709551615 < 1) {
                        if true {
                            v54 := GenS9{F0: 0, F1: 831}
                            var v55 uint16 = 1
                            if false {
                                if (42 > 18446744073709551615) {
                                    var v56 float64 = v4
                                    _ = v56
                                }
                            }
                            _ = v54
                            _ = v55
                        }
                    }
                    for fi57 := 0; fi57 < 5; fi57++ {
                        var v58 uint64 = 907
                        v1 -= v1
                        var v59 uint64 = 0
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                    switch "a" {
                    case "foo bar":
                        if false {
                            v60 := GenS10{F0: -1, F1: (65535 == 235)}
                            v61 := []int{}
                            _ = v60
                            _ = v61
                        }
                    default:
                        fmt.Println(v45, fi51, v45)
                    }
                } else {
                    v62 := []int{fi51, fi51, fi51}
                    _ = v62
                }
                _ = v53
            }
            _ = fi51
        }
    }
    var v63 string = ("" + "a")
    var v64 float64 = v4
    for fi65 := 0; fi65 < 8; fi65++ {
        v2 += -128
        var v66 int64 = 1
        _ = fi65
        _ = v66
    }
    if true {
        switch v63 {
        case "\n\t":
            v1 = v1
        default:
            v63 = v63
        }
        v67 := GenS4{F0: 402, F1: -2147483648, F2: (v5 | (v5 + v5))}
        var v68 uint32 = 2
        _ = v67
        _ = v68
    } else {
        for fi69 := 0; fi69 < 14; fi69++ {
            for fi70 := 0; fi70 < 10; fi70++ {
                switch "a" {
                case "":
                    v5 += v5
                }
                if (!true) {
                    var v71 uint16 = 0
                    var v72 uint = (v1 * (v1 - (v1 / (((v1 >> ((2) & 63))) | 1))))
                    v73 := GenS10{F0: 730, F1: false}
                    _ = v71
                    _ = v72
                    _ = v73
                }
                _ = fi70
            }
            _ = fi69
        }
    }
    v45 = (((1.5 / v45) / 1e10) + v45)
    v63 = v63
    v45 = (0.1 / (1.5 * v45))
    switch 9223372036854775807 {
    case -9223372036854775808:
        fmt.Println(v5, v63)
    case 128:
        v63 = ("\n\t" + ((((v63 + (v63 + ((((v63 + "foo bar") + "a") + v63) + "foo bar"))) + (((v63 + (v63 + v63)) + v63) + "a")) + "\n\t") + ("a" + "hello")))
    case -2147483648:
        v74 := GenS11{F0: 1}
        _ = v74
    default:
        var v75 uint = genFn0(0, v2, 42)
        _ = v75
    }
    switch ("" + v63) {
    case "foo bar":
        var v76 uint = 821
        _ = v76
    case "hello":
        v45 = v45
    default:
        v64 = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v63
    _ = v64
}
