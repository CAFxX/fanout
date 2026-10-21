package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint; F1 uint8; F2 int }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 int32; F1 int64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint64 }
func genFn0() uint64 {
    v5 := []int{42}
    var v6 int = 4294967295
    _ = v5
    _ = v6
    return 598
}

func main() {
    var v1 uint8 = 255
    var v2 int64 = -9223372036854775808
    v3 := []uint32{2}
    var v4 bool = true
    var v7 float64 = 0.1
    switch 127 {
    case -1:
        var v8 int = 128
        _ = v8
    case -9223372036854775808:
        var v9 float64 = 3.14159
        _ = v9
    default:
        v2 = 1
    }
    var v10 uint32 = 580
    v4 = (1 != -1)
    _ = len(v3)
    var v11 int64 = 9223372036854775807
    switch "" {
    case "a":
        v10 -= v10
    default:
        if v4 {
            if (!false) {
                v4 = (128 > 383)
                var v12 int16 = -1
                for fi13 := 0; fi13 < 4; fi13++ {
                    if false {
                        v11 += (^(v2))
                    } else {
                        for fi14 := 0; fi14 < 14; fi14++ {
                            fmt.Println(v12, v12, fi14)
                            var v15 int16 = (v12 / ((v12) | 1))
                            _ = fi14
                            _ = v15
                        }
                        fmt.Println(v12, v1)
                        var v16 int32 = -275
                        _ = v16
                    }
                    v3 = append(v3, (v10 % ((v10) | 1)))
                    _ = fi13
                }
                _ = v12
            } else {
                if ((true || ((v10 << ((2) & 63)) >= 691)) && (86 <= 100)) {
                    if v4 {
                        v7 -= (v7 * (((0.1 + 1.5) * (1e10 + 0.0)) + (1.5 + (0.0 + (((((v7 - v7) - v7) / (v7 - (float64(3.14159) * float64(-2.25)))) - float64((1e10 + 1.5))) - 0.1)))))
                    }
                    fmt.Println(v1)
                } else {
                    fmt.Println(v10, v11)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
        }
    }
    var v17 uint32 = 65535
    var v18 int16 = 0
    _ = len(v3)
    var v19 float64 = (v7 / float64(0.1))
    for fi20 := 0; fi20 < 12; fi20++ {
        switch "a" {
        case "":
            if (!(v1 != v1)) {
                v10 += 0
                v21 := GenS1{F0: (^(v2))}
                var v22 int16 = v18
                _ = v21
                _ = v22
            } else {
                switch "hello" {
                case "":
                    v10 -= v10
                case "hello":
                    v19 -= (v7 * (((v7 / v7) + ((v7 / (v19 * float64(((1.5 + 3.14159) * (0.1 / 1e10))))) / 1.0)) - (float64((0.0 + (0.0 + float32((1.5 * v19))))) - 3.14159)))
                case "foo bar":
                    if (-1 > 0) {
                        if (v4 && v4) {
                            v23 := GenS1{F0: v2}
                            v19 += float64(((3.14159 * 1e10) - float32((v7 + (v19 - (v7 / (float64(1.5) - v7)))))))
                            _ = v23
                        } else {
                            v24 := GenS2{F0: 1, F1: 451, F2: -201}
                            _ = v24
                        }
                        v25 := GenS2{F0: 255, F1: v17, F2: 2147483647}
                        if v4 {
                            for fi26 := 0; fi26 < 15; fi26++ {
                                var v27 int = (fi26 + 4294967295)
                                _ = fi26
                                _ = v27
                            }
                        }
                        _ = v25
                    } else {
                        v28 := GenS3{F0: 500, F1: uint8(v10), F2: fi20}
                        _ = v28
                    }
                default:
                    if (32767 < 42) {
                        for fi29 := 0; fi29 < 9; fi29++ {
                            if v4 {
                                var v30 int64 = (v2 ^ v2)
                                _ = v30
                            }
                            _ = fi29
                        }
                        for fi31 := 0; fi31 < 13; fi31++ {
                            var v32 int64 = (^(v2))
                            switch (fi31 + -457) {
                            case 128:
                                v33 := GenS4{F0: 18446744073709551615, F1: 755}
                                _ = v33
                            }
                            v4 = ((!v4) || (v2 > (-(v11))))
                            _ = fi31
                            _ = v32
                        }
                        for fi34 := 0; fi34 < 12; fi34++ {
                            fmt.Println(v2, v18, v19)
                            fmt.Println(v18, fi34, v11)
                            if len(v3) > 0 {
                                v3[0] = 609
                            }
                            _ = fi34
                        }
                    }
                }
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(fi20, v7)
                case "foo bar":
                    for fi35 := 0; fi35 < 10; fi35++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v19 = float64(1e10)
                        v36 := GenS1{F0: (^(v11))}
                        _ = fi35
                        _ = v36
                    }
                case "a":
                    switch (("hello" + "\n\t") + "foo bar") {
                    case "hello":
                        var v37 float64 = (v19 + 1e10)
                        _ = v37
                    case "a":
                        v38 := GenS2{F0: 255, F1: v17, F2: 64}
                        _ = v38
                    case "foo bar":
                        var v39 uint64 = 805
                        _ = v39
                    }
                default:
                    for fi40 := 0; fi40 < 5; fi40++ {
                        fmt.Println(v7, v18, fi20)
                        v4 = v4
                        v41 := GenS5{F0: 42, F1: genFn0()}
                        _ = fi40
                        _ = v41
                    }
                }
            }
        case "a":
            for fi42 := 0; fi42 < 9; fi42++ {
                if true {
                    if true {
                        if len(v3) > 0 {
                            v3[0] = v17
                        }
                        var v43 float32 = 1.5
                        v10 = (v17 + 279)
                        _ = v43
                    } else {
                        switch fi42 {
                        case 128:
                            if ((float64(0.0) / (float64(((float32(v19) - 1e10) / 0.1)) * 1e10)) >= v7) {
                                v19 = -2.25
                            } else {
                                var v44 string = "\n\t"
                                v1 = v1
                                v7 += float64(3.14159)
                                _ = v44
                            }
                        }
                        v19 += float64(1e10)
                    }
                    v45 := GenS1{F0: 1}
                    if (v4 || (((v4 || (fi42 == (fi42 + int(v10)))) && ((0 <= int64(v1)) && ((((^(fi20)) >= -2147483648) && v4) || v4))) || ((310 <= 0) || v4))) {
                        switch (fi20 >> ((65535) & 63)) {
                        case -9223372036854775808:
                            v2 = (v2 * v2)
                        default:
                            var v46 int32 = 1
                            _ = v46
                        }
                    }
                    _ = v45
                }
                v47 := GenS6{F0: -849, F1: v2}
                _ = fi42
                _ = v47
            }
        case "\n\t":
            switch (fi20 & (fi20 * fi20)) {
            case 1:
                v48 := GenS6{F0: -2147483648, F1: 0}
                _ = v48
            case 2147483647:
                v49 := GenS7{F0: v10}
                _ = v49
            default:
                var v50 int8 = 0
                _ = v50
            }
        }
        for fi51 := 0; fi51 < 17; fi51++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v52 := GenS8{F0: (("" + ("a" + ("a" + "hello"))) + ("\n\t" + "foo bar"))}
            var v53 float32 = 3.14159
            v3 = append(v3, v17)
            _ = fi51
            _ = v52
            _ = v53
        }
        _ = fi20
    }
    var v54 int8 = 127
    switch (((("a" + ("foo bar" + "hello")) + "") + "\n\t") + "a") {
    case "hello":
        for fi55 := 0; fi55 < 6; fi55++ {
            if false {
                v2 = 42
                for fi56 := 0; fi56 < 13; fi56++ {
                    _ = len(v3)
                    fmt.Println(v19)
                    if (!v4) {
                        var v57 uint = 255
                        v10 -= v17
                        var v58 uint32 = v17
                        _ = v57
                        _ = v58
                    } else {
                        v1 = 0
                        if ((^(v17)) <= (v10 << ((uint(v17)) & 63))) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v59 float32 = 0.0
                            var v60 uint = 18446744073709551615
                            _ = v59
                            _ = v60
                        } else {
                            v10 -= v10
                            v61 := GenS8{F0: "\n\t"}
                            v19 = ((v19 - float64(3.14159)) * v7)
                            _ = v61
                        }
                        if v4 {
                            fmt.Println(fi56, v7, v17)
                            v3 = append(v3, 1)
                        } else {
                            var v62 int = (fi56 << ((255) & 63))
                            _ = len(v3)
                            if false {
                                if len(v3) > 0 {
                                    v3[0] = v10
                                }
                            }
                            if false {
                                var v63 uint16 = 42
                                _ = len(v3)
                                var v64 uint = 42
                                var v65 float64 = ((v7 - 0.0) / v19)
                                v66 := GenS1{F0: -682}
                                _ = v63
                                _ = v64
                                _ = v65
                                _ = v66
                            }
                            _ = v62
                        }
                        var v67 float64 = ((0.1 / v7) * (0.1 - (1.5 / v7)))
                        if len(v3) > 0 {
                            v3[0] = (^(v10))
                        }
                        _ = v67
                    }
                    var v68 float64 = float64(1e10)
                    var v69 float32 = (0.0 * (1e10 * (0.1 / 0.1)))
                    _ = fi56
                    _ = v68
                    _ = v69
                }
                if (686 < 255) {
                    var v70 int32 = 42
                    var v71 float64 = (1.5 / 1.0)
                    _ = v70
                    _ = v71
                }
            }
            for fi72 := 0; fi72 < 18; fi72++ {
                switch "\n\t" {
                case "":
                    var v73 bool = v4
                    _ = v73
                case "a":
                    switch (fi55 * fi55) {
                    case 504:
                        if v4 {
                            var v74 uint8 = v1
                            var v75 bool = (v74 != 184)
                            v10 -= 2147483648
                            _ = v74
                            _ = v75
                        } else {
                            v2 += (-(v2))
                            v76 := GenS8{F0: ""}
                            if v4 {
                                v3 = append(v3, (v10 >> ((255) & 63)))
                                v4 = (4294967295 > uint64(v17))
                                fmt.Println(fi72)
                                var v77 int = (fi72 & 741)
                                _ = v77
                            }
                            var v78 int8 = v54
                            var v79 uint64 = 23
                            _ = v76
                            _ = v78
                            _ = v79
                        }
                    case 4294967295:
                        v80 := GenS4{F0: uint64(v1), F1: 0}
                        _ = v80
                    case 195:
                        v81 := GenS4{F0: uint64(v17), F1: 32767}
                        _ = v81
                    default:
                        v11 = 688
                    }
                default:
                    if len(v3) > 0 {
                        v3[0] = (v10 | v10)
                    }
                }
                var v82 bool = (true && v4)
                _ = fi72
                _ = v82
            }
            _ = fi55
        }
    case "\n\t":
        v83 := GenS9{F0: 2}
        _ = v83
    case "":
        var v84 int8 = (v54 % ((int8(v11)) | 1))
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v7
}
