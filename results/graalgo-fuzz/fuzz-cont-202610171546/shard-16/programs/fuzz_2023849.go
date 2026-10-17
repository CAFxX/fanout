package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 int; F1 bool; F2 int32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string; F1 int8; F2 uint16 }
type GenS10 struct { F0 uint64 }
func genFn0(p0 int, p1 uint, p2 bool) uint {
    v4 := GenS1{F0: -111, F1: ((32767 < 42) || p2)}
    var v5 uint = p1
    if p2 {
        var v6 int8 = -128
        _ = v6
    } else {
        v7 := []int{}
        if len(v7) > 0 {
            v7[0] = (p0 >> (((v5 ^ v5)) & 63))
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return v5
}
func genFn1() bool {
    v8 := GenS2{F0: 1, F1: 1, F2: 65535}
    var v9 bool = (false || false)
    fmt.Println(v9)
    _ = v8
    _ = v9
    return v9
}

func main() {
    var v1 int = 4294967295
    var v2 uint8 = 128
    var v3 uint = 65535
    if true {
        v2 += (v2 << (((v3 ^ (v3 / ((0) | 1)))) & 63))
    } else {
        v2 -= (v2 + v2)
        v10 := []int{128, -66}
        v10 = append(v10, (-(v1)))
        _ = v10
    }
    var v11 float64 = ((1.5 * ((0.1 / (3.14159 + 1.5)) / 1.5)) * (-2.25 * (-2.25 - 1.5)))
    fmt.Println(v1, v3)
    var v12 int16 = 1000
    switch v1 {
    case 1:
        switch v1 {
        case 42:
            var v13 uint16 = 356
            _ = v13
        case 127:
            if ((v1 < v1) || (-1 <= -2147483648)) {
                for fi14 := 0; fi14 < 15; fi14++ {
                    v11 = ((v11 * 3.14159) + (-2.25 + v11))
                    _ = fi14
                }
                v15 := GenS2{F0: 116, F1: 4294967295, F2: 65535}
                _ = v15
            }
        }
    }
    v11 += v11
    for fi16 := 0; fi16 < 4; fi16++ {
        for fi17 := 0; fi17 < 1; fi17++ {
            v1 += (fi16 << (((v3 >> (((v3 << (((v3 | (v3 ^ (v3 % (((v3 + (v3 | v3))) | 1))))) & 63))) & 63))) & 63))
            switch 483 {
            case 4294967295:
                v2 = v2
            default:
                fmt.Println(v11, v2)
            }
            _ = fi17
        }
        v3 = v3
        v18 := GenS3{F0: 42, F1: (^(v3))}
        _ = fi16
        _ = v18
    }
    v19 := []int{v1, v1}
    v3 -= (^(v3))
    if false {
        v19 = append(v19, v1)
        var v20 int64 = -1
        switch (v1 % (((^(v1))) | 1)) {
        case -2147483648:
            switch v1 {
            case 211:
                var v21 float32 = (float32(v11) * 3.14159)
                _ = v21
            case 127:
                switch ("a" + "\n\t") {
                case "a":
                    var v22 uint = v3
                    _ = v22
                case "hello":
                    switch "hello" {
                    case "":
                        switch "a" {
                        case "\n\t":
                            for fi23 := 0; fi23 < 15; fi23++ {
                                var v24 int32 = 0
                                switch 1 {
                                case 42:
                                    if false {
                                        var v25 string = "a"
                                        v26 := GenS4{F0: 0}
                                        _ = v25
                                        _ = v26
                                    } else {
                                        for fi27 := 0; fi27 < 20; fi27++ {
                                            v28 := GenS5{F0: v3, F1: v20}
                                            _ = fi27
                                            _ = v28
                                        }
                                    }
                                case 4294967295:
                                    for fi29 := 0; fi29 < 16; fi29++ {
                                        switch int(v24) {
                                        case 4294967295:
                                            var v30 int16 = (v12 % ((v12) | 1))
                                            _ = v30
                                        default:
                                            v31 := GenS4{F0: 691}
                                            _ = v31
                                        }
                                        _ = fi29
                                    }
                                case 32767:
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                default:
                                    _ = len(v19)
                                }
                                var v32 float64 = v11
                                _ = fi23
                                _ = v24
                                _ = v32
                            }
                        case "":
                            v19 = append(v19, (^(v1)))
                        case "a":
                            for fi33 := 0; fi33 < 15; fi33++ {
                                switch "a" {
                                case "a":
                                    v34 := GenS6{F0: 42, F1: v3}
                                    _ = v34
                                case "hello":
                                    switch "\n\t" {
                                    case "a":
                                        var v35 int8 = 42
                                        _ = v35
                                    case "":
                                        _ = len(v19)
                                    default:
                                        switch "hello" {
                                        case "hello":
                                            v19 = append(v19, v1)
                                        case "a":
                                            var v36 int = fi33
                                            _ = v36
                                        case "foo bar":
                                            switch (("" + "\n\t") + ("\n\t" + "a")) {
                                            case "a":
                                                switch v1 {
                                                case 9223372036854775807:
                                                    fmt.Println(v12)
                                                case 2147483647:
                                                    fmt.Println(fi33, v12)
                                                case 4294967295:
                                                    fmt.Println(v12)
                                                default:
                                                    v37 := GenS1{F0: 100, F1: false}
                                                    _ = v37
                                                }
                                            case "hello":
                                                for fi38 := 0; fi38 < 11; fi38++ {
                                                    switch v1 {
                                                    case -821:
                                                        if len(v19) > 0 {
                                                            v19[0] = 32767
                                                        }
                                                    case 744:
                                                        switch (fi38 ^ v1) {
                                                        case 0:
                                                            if (!true) {
                                                                fmt.Println(v11, v11, v11)
                                                            }
                                                        default:
                                                            v11 = 0.0
                                                        }
                                                    case -1:
                                                        if len(v19) > 0 {
                                                            v19[0] = fi33
                                                        }
                                                    }
                                                    if len(v19) >= 2 {
                                                        v19 = v19[1:len(v19)]
                                                    }
                                                    fmt.Println(fi38, v1)
                                                    fmt.Println(v3, fi33, v3)
                                                    _ = fi38
                                                }
                                            case "foo bar":
                                                switch "a" {
                                                case "a":
                                                    v1 -= int(v12)
                                                case "\n\t":
                                                    for fi39 := 0; fi39 < 15; fi39++ {
                                                        fmt.Println(v2, fi39)
                                                        _ = fi39
                                                    }
                                                default:
                                                    v20 -= (v20 * v20)
                                                }
                                            }
                                        }
                                    }
                                }
                                _ = fi33
                            }
                        }
                    default:
                        v40 := GenS2{F0: 0, F1: 1, F2: 255}
                        _ = v40
                    }
                case "foo bar":
                    var v41 int = v1
                    _ = v41
                default:
                    var v42 int = (v1 - v1)
                    _ = v42
                }
            case -1:
                for fi43 := 0; fi43 < 17; fi43++ {
                    for fi44 := 0; fi44 < 8; fi44++ {
                        v12 = -32768
                        _ = fi44
                    }
                    for fi45 := 0; fi45 < 14; fi45++ {
                        if true {
                            var v46 float64 = (1.5 - ((float64((0.0 + (3.14159 * 0.0))) * v11) - (0.1 / ((v11 + v11) + v11))))
                            v47 := GenS7{F0: -32, F1: false, F2: -2147483648}
                            _ = v46
                            _ = v47
                        }
                        _ = fi45
                    }
                    if false {
                        fmt.Println(v2, v2, v1)
                    }
                    _ = fi43
                }
            default:
                var v48 uint32 = 606
                _ = v48
            }
        case -9223372036854775808:
            var v49 int = 570
            _ = v49
        case 32768:
            fmt.Println(v11, v2, v20)
        default:
            var v50 float32 = 0.0
            _ = v50
        }
        _ = v20
    } else {
        var v51 uint8 = (v2 * (v2 % ((v2) | 1)))
        switch "foo bar" {
        case "\n\t":
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
        case "foo bar":
            var v52 string = ("hello" + (("a" + ("" + "foo bar")) + "a"))
            _ = v52
        case "a":
            v1 = (^(v1))
        default:
            if true {
                v2 -= v51
                if ((((-2.25 + 0.0) + 0.0) / (3.14159 - ((3.14159 / 3.14159) - (1.5 * 0.1)))) >= 0.0) {
                    v11 -= ((v11 + -2.25) / 0.1)
                    v3 = (^(v3))
                }
            } else {
                var v53 float32 = 1.5
                var v54 float64 = v11
                if false {
                    v19 = append(v19, v1)
                }
                _ = v53
                _ = v54
            }
        }
        for fi55 := 0; fi55 < 14; fi55++ {
            var v56 uint8 = 1
            v57 := GenS8{F0: uint(v56)}
            v58 := GenS9{F0: "\n\t", F1: 46, F2: 614}
            _ = fi55
            _ = v56
            _ = v57
            _ = v58
        }
        _ = v51
    }
    for fi59 := 0; fi59 < 14; fi59++ {
        var v60 int64 = 1
        v19 = append(v19, (fi59 & fi59))
        _ = fi59
        _ = v60
    }
    for fi61 := 0; fi61 < 13; fi61++ {
        fmt.Println(v1, v1)
        v3 = (v3 + v3)
        _ = fi61
    }
    v1 += (v1 + v1)
    var v62 uint = 0
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    switch "foo bar" {
    case "":
        var v63 uint32 = 2147483647
        _ = v63
    case "a":
        for fi64 := 0; fi64 < 11; fi64++ {
            for fi65 := 0; fi65 < 14; fi65++ {
                v2 = 2
                _ = fi65
            }
            _ = fi64
        }
    default:
        for fi66 := 0; fi66 < 14; fi66++ {
            if false {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                v67 := GenS10{F0: 4294967295}
                switch fi66 {
                case -475:
                    if len(v19) > 0 {
                        v19[0] = fi66
                    }
                default:
                    v2 += v2
                }
                _ = v67
            } else {
                v62 = genFn0(v1, 42, false)
            }
            _ = fi66
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v62
}
