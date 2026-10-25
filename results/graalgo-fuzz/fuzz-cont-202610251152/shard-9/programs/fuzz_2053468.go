package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 uint8 }
func genFn0() uint16 {
    v5 := []int{}
    v6 := GenS1{F0: 255, F1: 4294967295}
    var v7 int8 = 100
    for fi8 := 0; fi8 < 8; fi8++ {
        _ = len(v5)
        _ = fi8
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = v5
    _ = v6
    _ = v7
    return 32767
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[int]string)
    v3 := make(map[string]uint32)
    v4 := []int{}
    v1 = (^(v1))
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int8 = 1
        v11 := GenS2{F0: v1, F1: -128}
        if (false && true) {
            if (((0.0 * 1.5) - (1.5 * float64(((float32(float64(0.0)) * float32((0.0 / 3.14159))) + ((((1.5 + 0.1) - float32(float64(((3.14159 - (float32((3.14159 + (1e10 / 1.0))) * (3.14159 + 3.14159))) / 0.1)))) + ((((1e10 + (-2.25 * (((((1.5 * (1.5 * 1.5)) * ((1.5 * 0.1) - 0.0)) + 0.0) + 0.1) - 3.14159))) * 0.0) - float32((0.1 + (((((0.1 * -2.25) * float64(0.1)) - 3.14159) + 0.0) - 1e10)))) / 1.0)) - float32(0.0)))))) != 0.1) {
                if true {
                    v12 := GenS1{F0: 65535, F1: 886}
                    for fi13 := 0; fi13 < 11; fi13++ {
                        var v14 int16 = -32768
                        switch "hello" {
                        case "a":
                            v1 -= 1
                            v15 := GenS2{F0: (v1 / ((v1) | 1)), F1: (v10 % (((v10 >> ((4294967295) & 63))) | 1))}
                            _ = v15
                        default:
                            if (4294967295 != 945) {
                                var v16 float64 = 1e10
                                v17 := GenS3{F0: (-(v14))}
                                var v18 float32 = (0.0 - 1e10)
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                        }
                        var v19 float64 = (0.1 / 3.14159)
                        _ = fi13
                        _ = v14
                        _ = v19
                    }
                    v10 = (^(v10))
                    _ = v12
                } else {
                    var v20 uint32 = 65535
                    for fi21 := 0; fi21 < 9; fi21++ {
                        v10 -= v10
                        _ = fi21
                    }
                    if false {
                        if true {
                            if (!(65535 < 4294967295)) {
                                fmt.Println(v1, v20, v1)
                                for fi22 := 0; fi22 < 10; fi22++ {
                                    delete(v2, (fi22 % ((fi9) | 1)))
                                    switch (fi9 << ((65535) & 63)) {
                                    case 4294967295:
                                        var v23 uint64 = 959
                                        _ = v23
                                    case 9223372036854775807:
                                        v1 += (v20 * (v1 + (^(v20))))
                                    case -9223372036854775808:
                                        v20 = v1
                                    default:
                                        var v24 int16 = 32767
                                        _ = v24
                                    }
                                    var v25 uint = 590
                                    _ = fi22
                                    _ = v25
                                }
                                var v26 int64 = 1
                                _ = v26
                            } else {
                                v1 -= (v20 - (v20 - 457))
                                if (-741 == 0) {
                                    delete(v2, (^(fi9)))
                                    if (0 < 32767) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                }
                            }
                            _ = len(v3)
                            if ((v1 % ((4294967295) | 1)) < 470) {
                                v27 := GenS3{F0: 578}
                                _ = v27
                            }
                        } else {
                            if (1 == 166) {
                                var v28 int64 = 9223372036854775807
                                v29 := GenS3{F0: -32768}
                                _ = v28
                                _ = v29
                            }
                            switch (fi9 & (fi9 | fi9)) {
                            case 128:
                                switch (((((("a" + "foo bar") + "\n\t") + "foo bar") + "hello") + "") + "\n\t") {
                                case "foo bar":
                                    var v30 int8 = (v10 + (v10 / ((v10) | 1)))
                                    _ = v30
                                case "a":
                                    for fi31 := 0; fi31 < 13; fi31++ {
                                        v20 = 959
                                        _ = fi31
                                    }
                                case "\n\t":
                                    for fi32 := 0; fi32 < 17; fi32++ {
                                        var v33 int32 = -700
                                        fmt.Println(v33)
                                        var v34 float32 = (1e10 * (0.1 + (-2.25 + 0.0)))
                                        v35 := GenS2{F0: v20, F1: -128}
                                        _ = fi32
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                }
                            default:
                                var v36 int8 = -128
                                _ = v36
                            }
                        }
                        if (0 == 1) {
                            var v37 int16 = -246
                            var v38 uint = 65535
                            _ = v37
                            _ = v38
                        }
                        v10 -= -128
                    } else {
                        var v39 string = "a"
                        switch "\n\t" {
                        case "a":
                            for fi40 := 0; fi40 < 10; fi40++ {
                                var v41 int = int(v10)
                                v42 := GenS1{F0: genFn0(), F1: 18446744073709551615}
                                var v43 float32 = (float32(float64(((float32(float64(1e10)) - 0.1) / float32(1e10)))) - 1e10)
                                _ = fi40
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                        default:
                            for fi44 := 0; fi44 < 17; fi44++ {
                                v45 := GenS2{F0: v20, F1: v10}
                                _ = fi44
                                _ = v45
                            }
                        }
                        for fi46 := 0; fi46 < 11; fi46++ {
                            v20 = (v1 - v1)
                            v47 := v3[v39]
                            _ = fi46
                            _ = v47
                        }
                        _ = v39
                    }
                    _ = v20
                }
            } else {
                if false {
                    var v48 uint = 255
                    _ = v48
                } else {
                    for fi49 := 0; fi49 < 13; fi49++ {
                        v1 += (^(v1))
                        _ = fi49
                    }
                }
                if true {
                    var v50 bool = false
                    for fi51 := 0; fi51 < 14; fi51++ {
                        _ = len(v4)
                        if v50 {
                            v52 := GenS4{F0: v10}
                            _ = v52
                        } else {
                            if len(v4) > 0 {
                                v4[0] = (fi9 | (fi51 / (((^(fi51))) | 1)))
                            }
                            var v53 int8 = v10
                            _ = v53
                        }
                        fmt.Println(v10, v1)
                        _ = fi51
                    }
                    _ = v50
                } else {
                    if false {
                        var v54 uint8 = 1
                        for fi55 := 0; fi55 < 15; fi55++ {
                            if ((!true) && (0 < 9223372036854775807)) {
                                v4 = append(v4, fi55)
                                _ = len(v3)
                                v54 = (v54 ^ v54)
                            }
                            _ = fi55
                        }
                        _ = v54
                    } else {
                        v1 += 4294967295
                        v56 := v2[36]
                        _ = v56
                    }
                }
            }
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v1)
    for fi57 := 0; fi57 < 2; fi57++ {
        v58 := GenS5{F0: 18}
        var v59 uint8 = 90
        _ = fi57
        _ = v58
        _ = v59
    }
    var v60 int16 = 32767
    for fi61 := 0; fi61 < 15; fi61++ {
        for fi62 := 0; fi62 < 1; fi62++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v63 := GenS6{F0: -11, F1: 869}
            _ = fi62
            _ = v63
        }
        _ = fi61
    }
    switch -1 {
    case -608:
        if true {
            switch "hello" {
            case "":
                v2[32767] = ""
            default:
                v1 = (v1 + v1)
            }
            if (42 < 1) {
                var v64 string = "a"
                fmt.Println(v1, v60, v64)
                _ = v64
            } else {
                switch "a" {
                case "":
                    v65 := GenS7{F0: 1}
                    _ = v65
                case "hello":
                    v66 := GenS8{F0: -1, F1: 1}
                    _ = v66
                }
                for fi67 := 0; fi67 < 16; fi67++ {
                    switch "a" {
                    case "a":
                        var v68 uint = 42
                        _ = v68
                    case "hello":
                        if true {
                            v1 += 2
                            if len(v4) > 0 {
                                v4[0] = (fi67 ^ -9223372036854775808)
                            }
                        }
                    }
                    switch ("a" + ("" + "hello")) {
                    case "a":
                        delete(v3, "\n\t")
                    case "":
                        v60 = v60
                    case "hello":
                        var v69 int64 = 0
                        _ = v69
                    default:
                        var v70 int = -661
                        _ = v70
                    }
                    var v71 int16 = v60
                    _ = fi67
                    _ = v71
                }
            }
            switch "hello" {
            case "":
                if (2147483647 <= 0) {
                    v1 = (v1 >> ((351) & 63))
                    if (359 != 276) {
                        v60 -= (v60 ^ v60)
                        v1 = (^(v1))
                    } else {
                        v72 := GenS1{F0: 115, F1: 363}
                        var v73 string = ""
                        fmt.Println(v60, v73)
                        _ = v72
                        _ = v73
                    }
                    _ = len(v4)
                } else {
                    if (!(!(0 <= -1))) {
                        if false {
                            v74 := GenS9{F0: -934}
                            _ = v74
                        }
                        v1 = v1
                        var v75 bool = (!false)
                        _ = v75
                    } else {
                        if (true || (!(true || (1 < 42)))) {
                            fmt.Println(v60, v1)
                        }
                        v76 := GenS5{F0: 153}
                        v60 -= v60
                        _ = v76
                    }
                    fmt.Println(v60, v1)
                }
            default:
                if len(v4) > 0 {
                    v4[0] = 9223372036854775807
                }
            }
        } else {
            if (0.0 < 3.14159) {
                v77 := GenS10{F0: -738, F1: 255}
                v1 += v1
                for fi78 := 0; fi78 < 19; fi78++ {
                    v60 = (v60 << ((751) & 63))
                    if (true && true) {
                        if (((!false) || false) && (true && (!true))) {
                            fmt.Println(v60, fi78, v1)
                            for fi79 := 0; fi79 < 11; fi79++ {
                                _ = len(v4)
                                _ = fi79
                            }
                            var v80 bool = (772 >= 42)
                            _ = v80
                        }
                        fmt.Println(v1, fi78, fi78)
                        fmt.Println(fi78, fi78, v60)
                    }
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(v60, v1)
                    case "hello":
                        v60 = -840
                    }
                    _ = fi78
                }
                _ = v77
            } else {
                fmt.Println(v1)
            }
            if len(v4) > 0 {
                v4[0] = 42
            }
        }
    default:
        if true {
            switch 127 {
            case 792:
                var v81 int16 = 111
                _ = v81
            case -251:
                var v82 uint32 = (^(v1))
                _ = v82
            default:
                v60 -= (v60 ^ (v60 / ((v60) | 1)))
            }
        }
    }
    v60 = v60
    v60 = v60
    if ((float32(float64((-2.25 + (float32(0.0) + 1.5)))) - 1e10) < 0.1) {
        v60 += (-(v60))
    } else {
        if (435 != 32767) {
            v83 := GenS10{F0: 0, F1: 127}
            _ = v83
        }
    }
    if (!true) {
        var v84 float64 = (0.1 * 1.5)
        if (1e10 <= (((-2.25 * (1e10 / float32(v84))) - (0.0 + 0.1)) / 0.1)) {
            if true {
                delete(v3, "a")
                switch 9223372036854775807 {
                case 2147483647:
                    if false {
                        fmt.Println(v84, v60)
                        var v85 int = -73
                        v84 = (v84 * v84)
                        _ = v85
                    }
                case 933:
                    for fi86 := 0; fi86 < 6; fi86++ {
                        for fi87 := 0; fi87 < 11; fi87++ {
                            if true {
                                v1 -= (^(v1))
                                v1 += v1
                            } else {
                                fmt.Println(fi87, v1, fi87)
                                switch fi87 {
                                case -9223372036854775808:
                                    delete(v3, ("hello" + (("\n\t" + "\n\t") + "hello")))
                                case -532:
                                    var v88 uint8 = 2
                                    _ = v88
                                default:
                                    var v89 uint32 = v1
                                    _ = v89
                                }
                            }
                            _ = fi87
                        }
                        _ = fi86
                    }
                default:
                    v90 := GenS1{F0: genFn0(), F1: 2}
                    _ = v90
                }
            }
            var v91 float32 = float32(v84)
            _ = v91
        }
        _ = v84
    } else {
        for fi92 := 0; fi92 < 12; fi92++ {
            if (!false) {
                var v93 string = ""
                _ = v93
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v2[fi92] = ("a" + ("hello" + "a"))
            }
            _ = fi92
        }
        v60 = -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v60
}
