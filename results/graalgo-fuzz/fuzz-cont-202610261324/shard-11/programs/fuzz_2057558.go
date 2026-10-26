package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint16; F1 int8; F2 int16 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 int8; F1 int16; F2 int16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool; F1 int16; F2 uint32 }
type GenS9 struct { F0 int }

func main() {
    var v1 uint8 = 1
    v2 := []uint32{}
    v3 := make(map[uint32]bool)
    var v4 int32 = 340
    switch "\n\t" {
    case "":
        var v5 int64 = 1
        _ = v5
    case "foo bar":
        v6 := GenS1{F0: 970, F1: (v4 + v4)}
        _ = v6
    case "hello":
        var v7 uint64 = 42
        _ = v7
    }
    if false {
        for fi8 := 0; fi8 < 4; fi8++ {
            switch ("a" + "foo bar") {
            case "\n\t":
                switch fi8 {
                case 4294967295:
                    v1 = (v1 ^ v1)
                default:
                    v3[4294967295] = ((false || (false && true)) || (true && false))
                }
            case "foo bar":
                switch (("foo bar" + ("a" + ("hello" + ("hello" + "")))) + "") {
                case "\n\t":
                    var v9 int16 = -32768
                    _ = v9
                case "":
                    for fi10 := 0; fi10 < 9; fi10++ {
                        fmt.Println(v4, fi8, fi8)
                        switch fi10 {
                        case 0:
                            v1 -= 2
                        case 32767:
                            var v11 float32 = 1.5
                            _ = v11
                        }
                        _ = fi10
                    }
                default:
                    fmt.Println(v1, fi8)
                }
            }
            fmt.Println(fi8)
            _ = fi8
        }
        if (false && (false && (42 > 423))) {
            v4 += (v4 % ((v4) | 1))
        }
        if false {
            v3[4294967295] = false
            if (100 >= -1) {
                _ = len(v2)
            } else {
                switch 433 {
                case -947:
                    for fi12 := 0; fi12 < 9; fi12++ {
                        v1 = 23
                        for fi13 := 0; fi13 < 9; fi13++ {
                            var v14 int = 85
                            v15 := GenS1{F0: 0, F1: (v4 << ((0) & 63))}
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        var v16 uint8 = 149
                        _ = fi12
                        _ = v16
                    }
                case 4294967295:
                    var v17 float64 = 3.14159
                    _ = v17
                default:
                    v18 := v3[65535]
                    _ = v18
                }
                v1 -= v1
            }
        } else {
            _ = len(v3)
            var v19 uint16 = 751
            if (v4 >= v4) {
                _ = len(v3)
                if len(v2) > 0 {
                    v2[0] = 828
                }
                v3[2] = true
            }
            _ = v19
        }
    } else {
        for fi20 := 0; fi20 < 2; fi20++ {
            _ = len(v3)
            _ = fi20
        }
    }
    v21 := GenS2{F0: 1, F1: -1, F2: -1}
    v4 += 2147483647
    var v22 int64 = 42
    fmt.Println(v4)
    v22 = int64(v1)
    v23 := GenS2{F0: 2, F1: 47, F2: 0}
    var v24 uint64 = 954
    var v25 uint32 = 609
    for fi26 := 0; fi26 < 8; fi26++ {
        fmt.Println(v4, v24, fi26)
        switch "a" {
        case "":
            switch (-(fi26)) {
            case 9223372036854775807:
                v27 := GenS1{F0: 1, F1: v4}
                _ = v27
            case 42:
                var v28 uint16 = 119
                _ = v28
            default:
                switch 128 {
                case 183:
                    var v29 int16 = 0
                    _ = v29
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
        case "\n\t":
            fmt.Println(v24, v4, v22)
        case "hello":
            if false {
                var v30 uint32 = 717
                _ = v30
            } else {
                v31 := GenS3{F0: -173, F1: (v1 | 0), F2: (v24 % (((v24 >> ((65535) & 63))) | 1))}
                for fi32 := 0; fi32 < 13; fi32++ {
                    switch fi26 {
                    case 32767:
                        fmt.Println(v1, v24)
                        fmt.Println(fi32, v1)
                        var v33 int32 = (v4 / ((-1) | 1))
                        _ = v33
                    case -1:
                        var v34 int8 = 1
                        _ = v34
                    case -76:
                        var v35 int8 = 127
                        _ = v35
                    }
                    v36 := GenS3{F0: 1000, F1: v1, F2: 0}
                    v2 = append(v2, (v25 / (((v25 << ((0) & 63))) | 1)))
                    _ = fi32
                    _ = v36
                }
                _ = v31
            }
        default:
            switch (^(fi26)) {
            case 104:
                var v37 int16 = 32767
                _ = v37
            case -9223372036854775808:
                if (fi26 == (fi26 * fi26)) {
                    v38 := GenS2{F0: 202, F1: 100, F2: int16(v1)}
                    switch "\n\t" {
                    case "\n\t":
                        v1 += (v1 * (v1 ^ (v1 << ((303) & 63))))
                    case "a":
                        v39 := GenS3{F0: 489, F1: (^(v1)), F2: 4294967295}
                        _ = v39
                    default:
                        if ((v4 / ((v4) | 1)) < v4) {
                            v40 := GenS1{F0: 841, F1: v4}
                            v41 := GenS2{F0: 0, F1: 42, F2: 651}
                            _ = v40
                            _ = v41
                        } else {
                            v42 := GenS3{F0: int16(v25), F1: v1, F2: v24}
                            _ = v42
                        }
                    }
                    _ = v38
                } else {
                    fmt.Println(fi26)
                    var v43 int16 = 32767
                    fmt.Println(v43)
                    _ = v43
                }
            case 1:
                for fi44 := 0; fi44 < 15; fi44++ {
                    v22 += (v22 + 0)
                    v25 += v25
                    _ = fi44
                }
            default:
                var v45 int64 = v22
                _ = v45
            }
        }
        _ = fi26
    }
    v46 := GenS2{F0: 0, F1: 42, F2: 249}
    v47 := GenS4{F0: v24, F1: 1}
    if false {
        v22 = (v22 - v22)
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v4)
    case "foo bar":
        v22 = v22
    default:
        if false {
            switch -9223372036854775808 {
            case -2147483648:
                var v48 int16 = -32768
                _ = v48
            case -905:
                switch -925 {
                case 42:
                    if (!(!false)) {
                        if (v1 >= v1) {
                            v49 := GenS5{F0: "foo bar", F1: 2}
                            _ = v49
                        } else {
                            fmt.Println(v4, v4, v24)
                            var v50 uint64 = (v24 & 611)
                            var v51 float32 = 1.5
                            _ = v50
                            _ = v51
                        }
                    } else {
                        var v52 string = ""
                        v22 = v22
                        _ = v52
                    }
                case 0:
                    v25 -= (v25 / (((v25 - v25)) | 1))
                default:
                    switch (((((((("foo bar" + ("" + "foo bar")) + ("hello" + "a")) + ("hello" + "foo bar")) + ("hello" + "")) + "hello") + ("\n\t" + "a")) + "hello") + "hello") {
                    case "hello":
                        v53 := GenS2{F0: 825, F1: -1, F2: -32768}
                        _ = v53
                    }
                }
            case 4294967295:
                switch 32767 {
                case -1:
                    v54 := GenS2{F0: 32767, F1: 100, F2: 834}
                    _ = v54
                case -2147483648:
                    for fi55 := 0; fi55 < 7; fi55++ {
                        if true {
                            var v56 int64 = v22
                            for fi57 := 0; fi57 < 11; fi57++ {
                                switch ("a" + "") {
                                case "\n\t":
                                    var v58 int64 = (^(v22))
                                    _ = v58
                                case "":
                                    v4 += (v4 % (((v4 | 785)) | 1))
                                case "foo bar":
                                    var v59 uint16 = 288
                                    if (127 != v1) {
                                        var v60 uint16 = (v59 ^ v59)
                                        fmt.Println(v24)
                                        fmt.Println(fi55, v59)
                                        _ = v60
                                    } else {
                                        if (!(63 >= -55)) {
                                            if (973 >= (v59 - (^(v59)))) {
                                                var v61 int8 = 0
                                                var v62 int32 = v4
                                                if ((!true) || (0.1 > 3.14159)) {
                                                    if true {
                                                        if true {
                                                            v25 = (v25 / ((v25) | 1))
                                                            v63 := v3[(v25 | v25)]
                                                            _ = v63
                                                        }
                                                        switch (^(fi57)) {
                                                        case 865:
                                                            _ = len(v3)
                                                        default:
                                                            fmt.Println(v25, v56, fi57)
                                                        }
                                                    }
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v1 = 178
                                                _ = v61
                                                _ = v62
                                            }
                                            fmt.Println(v56, v22)
                                            _ = len(v3)
                                            switch "\n\t" {
                                            case "hello":
                                                if (0 != 42) {
                                                    v24 -= 4294967295
                                                }
                                            case "foo bar":
                                                var v64 float64 = (0.1 * 1.5)
                                                v22 = v22
                                                _ = v64
                                            case "\n\t":
                                                delete(v3, (v25 * (^(v25))))
                                            }
                                        } else {
                                            fmt.Println(v1, v56, v59)
                                            var v65 string = ""
                                            fmt.Println(fi55)
                                            _ = v65
                                        }
                                        v2 = append(v2, 2147483648)
                                    }
                                    _ = v59
                                }
                                switch (fi57 << ((0) & 63)) {
                                case 9223372036854775807:
                                    switch (fi57 % ((fi55) | 1)) {
                                    case 466:
                                        if len(v2) > 0 {
                                            v2[0] = 1
                                        }
                                    case 32767:
                                        var v66 float64 = 0.0
                                        _ = v66
                                    }
                                }
                                v1 = 0
                                _ = fi57
                            }
                            if false {
                                switch "hello" {
                                case "a":
                                    if false {
                                        v67 := GenS1{F0: 18446744073709551615, F1: 0}
                                        _ = v67
                                    } else {
                                        var v68 int = fi55
                                        v69 := GenS6{F0: 0, F1: 874, F2: -745}
                                        if false {
                                            if true {
                                                var v70 uint32 = v25
                                                for fi71 := 0; fi71 < 3; fi71++ {
                                                    var v72 int64 = v22
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v4 += v4
                                                    _ = fi71
                                                    _ = v72
                                                }
                                                switch (v68 & fi55) {
                                                case 42:
                                                    v73 := GenS7{F0: v24}
                                                    _ = v73
                                                }
                                                _ = v70
                                            } else {
                                                v56 -= v56
                                                if true {
                                                    var v74 float32 = -2.25
                                                    v56 -= 107
                                                    v3[2] = (!(1 > 1))
                                                    _ = v74
                                                }
                                            }
                                        }
                                        _ = v68
                                        _ = v69
                                    }
                                }
                            }
                            _ = v56
                        } else {
                            v75 := GenS4{F0: (v24 % ((v24) | 1)), F1: (v22 << ((42) & 63))}
                            switch int(v22) {
                            case 566:
                                if ((fi55 % ((fi55) | 1)) != fi55) {
                                    v22 += (v22 ^ v22)
                                    v22 -= -1
                                    var v76 uint64 = 4294967295
                                    _ = v76
                                } else {
                                    var v77 int8 = 42
                                    for fi78 := 0; fi78 < 17; fi78++ {
                                        v79 := GenS3{F0: -32768, F1: (v1 ^ v1), F2: 1}
                                        var v80 float64 = (3.14159 - ((float64(1e10) * 3.14159) + (0.1 - -2.25)))
                                        _ = fi78
                                        _ = v79
                                        _ = v80
                                    }
                                    fmt.Println(v22)
                                    _ = v77
                                }
                            default:
                                fmt.Println(v25, v22)
                            }
                            _ = v75
                        }
                        v4 = (v4 ^ (v4 % ((-1) | 1)))
                        _ = fi55
                    }
                default:
                    v81 := GenS8{F0: true, F1: -629, F2: (v25 % (((v25 % ((65535) | 1))) | 1))}
                    _ = v81
                }
            default:
                var v82 int32 = v4
                _ = v82
            }
            v83 := GenS5{F0: "foo bar", F1: 42}
            for fi84 := 0; fi84 < 3; fi84++ {
                if false {
                    for fi85 := 0; fi85 < 5; fi85++ {
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v86 int32 = (v4 ^ v4)
                            fmt.Println(v86, fi84)
                            var v87 int16 = 42
                            var v88 uint = 353
                            _ = v86
                            _ = v87
                            _ = v88
                        }
                        _ = fi85
                    }
                }
                _ = fi84
            }
            _ = v83
        } else {
            var v89 int32 = (v4 | v4)
            v90 := GenS9{F0: -2147483648}
            _ = v89
            _ = v90
        }
    }
    v22 += (^(v22))
    var v91 uint = 1
    v92 := v3[v25]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v91
    _ = v92
}
