package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int8 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint16; F1 int16; F2 int }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int64; F1 int32; F2 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 string; F1 bool; F2 int }
type GenS9 struct { F0 string; F1 uint16 }
type GenS10 struct { F0 uint32; F1 string }
type GenS11 struct { F0 int; F1 string; F2 string }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 int16 = 1000
    var v4 uint = 31
    var v5 float32 = 0.0
    var v6 uint64 = 42
    v3 = -184
    v7 := GenS1{F0: -32768, F1: -1, F2: -64}
    switch "\n\t" {
    case "":
        if (!(65535 != 0)) {
            v8 := v1[94]
            if (0.0 < 3.14159) {
                v3 = v3
                if (0 > v4) {
                    if (55 == 128) {
                        if false {
                            for fi9 := 0; fi9 < 5; fi9++ {
                                var v10 uint32 = 65535
                                v11 := []int{}
                                _ = fi9
                                _ = v10
                                _ = v11
                            }
                            if true {
                                for fi12 := 0; fi12 < 16; fi12++ {
                                    var v13 uint64 = (^(v6))
                                    v5 += v5
                                    for fi14 := 0; fi14 < 7; fi14++ {
                                        v15 := GenS1{F0: (v3 | v3), F1: 435, F2: 42}
                                        _ = fi14
                                        _ = v15
                                    }
                                    _ = fi12
                                    _ = v13
                                }
                                v16 := GenS1{F0: -591, F1: 2147483647, F2: 127}
                                _ = v16
                            } else {
                                v17 := []int{v8, (v8 ^ (v8 - (v8 & 1)))}
                                var v18 uint64 = 0
                                _ = v17
                                _ = v18
                            }
                            v3 = (v3 % (((v3 - 42)) | 1))
                        } else {
                            v8 -= v8
                            if true {
                                if (!true) {
                                    if (false && true) {
                                        for fi19 := 0; fi19 < 19; fi19++ {
                                            v8 = fi19
                                            _ = fi19
                                        }
                                    } else {
                                        v6 = v6
                                        v8 -= (v8 ^ (^(v8)))
                                        if ((32767 >= 42) || ((2 != 152) || false)) {
                                            v2[1] = "\n\t"
                                            if true {
                                                for fi20 := 0; fi20 < 20; fi20++ {
                                                    v4 -= 4294967295
                                                    if (!false) {
                                                        var v21 float32 = (v5 + 0.1)
                                                        _ = v21
                                                    }
                                                    _ = fi20
                                                }
                                                if false {
                                                    if false {
                                                        v22 := GenS1{F0: v3, F1: 42, F2: -33}
                                                        _ = v22
                                                    } else {
                                                        var v23 uint8 = 128
                                                        for fi24 := 0; fi24 < 18; fi24++ {
                                                            var v25 int = fi24
                                                            _ = fi24
                                                            _ = v25
                                                        }
                                                        _ = v23
                                                    }
                                                    for fi26 := 0; fi26 < 6; fi26++ {
                                                        v27 := GenS1{F0: v3, F1: -190, F2: -1}
                                                        _ = fi26
                                                        _ = v27
                                                    }
                                                }
                                            } else {
                                                delete(v1, 417)
                                            }
                                        }
                                    }
                                    if true {
                                        for fi28 := 0; fi28 < 15; fi28++ {
                                            v5 = v5
                                            var v29 uint32 = 878
                                            _ = fi28
                                            _ = v29
                                        }
                                    } else {
                                        v6 = (^(v6))
                                        var v30 bool = false
                                        var v31 string = "hello"
                                        _ = v30
                                        _ = v31
                                    }
                                    v8 = 9223372036854775807
                                }
                                var v32 bool = (((50 == 1) || (uint16(v4) < 0)) || (42 == -9223372036854775808))
                                _ = v32
                            } else {
                                var v33 float32 = v5
                                switch v8 {
                                case 128:
                                    fmt.Println(v33, v33, v3)
                                case 0:
                                    v34 := GenS2{F0: "hello", F1: 1}
                                    _ = v34
                                case 32767:
                                    var v35 uint8 = 0
                                    _ = v35
                                default:
                                    v36 := []int{(-(v8)), int(v3)}
                                    _ = v36
                                }
                                _ = v33
                            }
                            for fi37 := 0; fi37 < 1; fi37++ {
                                v5 -= v5
                                v5 = v5
                                _ = fi37
                            }
                        }
                        v38 := v1[1]
                        _ = v38
                    }
                } else {
                    if false {
                        v6 -= (v6 | (v6 >> (((v4 ^ 42)) & 63)))
                    } else {
                        fmt.Println(v5, v4, v5)
                        v39 := GenS2{F0: ((("foo bar" + "") + "") + ""), F1: 42}
                        _ = v39
                    }
                    if (2147483647 <= 168) {
                        if true {
                            v40 := GenS1{F0: v3, F1: 2147483647, F2: 0}
                            _ = v40
                        }
                        var v41 string = (("a" + ("\n\t" + "foo bar")) + ("a" + "hello"))
                        _ = v41
                    } else {
                        var v42 int8 = -128
                        _ = v42
                    }
                    if true {
                        v43 := GenS3{F0: uint16(v4), F1: v3, F2: 0}
                        v8 -= 42
                        for fi44 := 0; fi44 < 14; fi44++ {
                            var v45 int8 = -1
                            _ = fi44
                            _ = v45
                        }
                        _ = v43
                    }
                }
                v6 += (^(v6))
            }
            _ = v8
        } else {
            v46 := []int{111, -1, 4294967295, 615}
            if false {
                for fi47 := 0; fi47 < 12; fi47++ {
                    var v48 string = ("hello" + "foo bar")
                    _ = fi47
                    _ = v48
                }
            } else {
                var v49 uint64 = v6
                var v50 uint8 = 42
                switch -600 {
                case 0:
                    if false {
                        var v51 int8 = 127
                        delete(v2, 2147483648)
                        delete(v2, uint32(v50))
                        _ = v51
                    }
                case -1:
                    var v52 uint32 = 2147483648
                    _ = v52
                case 1:
                    delete(v1, 2)
                default:
                    v53 := v2[65535]
                    _ = v53
                }
                _ = v49
                _ = v50
            }
            var v54 int64 = 118
            _ = v46
            _ = v54
        }
    }
    for fi55 := 0; fi55 < 5; fi55++ {
        fmt.Println(v3, v4, v5)
        var v56 float64 = 3.14159
        _ = fi55
        _ = v56
    }
    var v57 int8 = 0
    fmt.Println(v5, v57)
    switch 2147483647 {
    case 0:
        fmt.Println(v6)
    case 647:
        v5 = (float32(1.5) / v5)
    case 127:
        for fi58 := 0; fi58 < 17; fi58++ {
            for fi59 := 0; fi59 < 7; fi59++ {
                v5 = 0.0
                _ = fi59
            }
            v60 := GenS4{F0: 2147483648, F1: 18446744073709551615}
            _ = fi58
            _ = v60
        }
    }
    v61 := []int{-745, 0, 32767, 461}
    var v62 bool = false
    var v63 int = -9223372036854775808
    var v64 uint16 = 41
    if ((^(v64)) >= 367) {
        v64 = v64
    }
    switch -1 {
    case -995:
        if (v62 && ((!false) || (1 < 61))) {
            v3 = (v3 * (-(v3)))
            v3 = 0
            for fi65 := 0; fi65 < 18; fi65++ {
                if (true && ("\n\t" == (("\n\t" + (("\n\t" + "") + "\n\t")) + "a"))) {
                    if (!v62) {
                        v6 -= (v6 << ((1) & 63))
                        var v66 uint8 = 139
                        _ = v66
                    } else {
                        if ("\n\t" >= ((("" + ("\n\t" + "foo bar")) + "\n\t") + "")) {
                            v1[2] = fi65
                            var v67 float64 = (0.1 / (3.14159 / 3.14159))
                            _ = len(v61)
                            _ = v67
                        }
                        for fi68 := 0; fi68 < 15; fi68++ {
                            v3 -= v3
                            if false {
                                var v69 float64 = (((3.14159 + (float64(v5) / ((1.5 - 0.1) / float64(v5)))) * float64((-2.25 - (v5 * float32(-2.25))))) / (1.5 + float64(v5)))
                                _ = v69
                            } else {
                                if (0 != -1) {
                                    v6 = (v6 | (v6 / (((v6 | (v6 | 42))) | 1)))
                                } else {
                                    var v70 uint = v4
                                    _ = v70
                                }
                                v62 = (v57 < int8(v4))
                                v71 := GenS5{F0: 42, F1: 653, F2: v64}
                                _ = v71
                            }
                            switch fi65 {
                            case 9223372036854775807:
                                if ((-(v57)) != -1) {
                                    switch (("foo bar" + "foo bar") + "a") {
                                    case "a":
                                        v6 = (v6 / ((2) | 1))
                                    case "":
                                        v72 := GenS6{F0: fi68}
                                        _ = v72
                                    default:
                                        v73 := v1[631]
                                        _ = v73
                                    }
                                }
                            case 42:
                                if true {
                                    var v74 uint64 = (v6 >> (((^(v4))) & 63))
                                    var v75 int16 = v3
                                    _ = v74
                                    _ = v75
                                }
                            }
                            _ = fi68
                        }
                    }
                    var v76 int32 = -108
                    switch (v63 / ((fi65) | 1)) {
                    case -164:
                        v77 := GenS1{F0: v3, F1: v76, F2: (v57 * (v57 / ((127) | 1)))}
                        _ = v77
                    case 128:
                        var v78 float32 = ((((v5 - ((v5 - float32(1.5)) + (-2.25 / (v5 + 0.1)))) / float32(3.14159)) - (1.5 / 1.0)) / 3.14159)
                        _ = v78
                    default:
                        v79 := GenS3{F0: v64, F1: 0, F2: int(v4)}
                        _ = v79
                    }
                    _ = v76
                }
                _ = fi65
            }
        }
    case 1:
        v57 -= (-(v57))
    default:
        _ = len(v2)
    }
    v80 := GenS7{F0: 0, F1: v62}
    for fi81 := 0; fi81 < 6; fi81++ {
        for fi82 := 0; fi82 < 9; fi82++ {
            if (227 >= -933) {
                v2[374] = ("a" + "")
                switch "" {
                case "\n\t":
                    if (v62 || (-9223372036854775808 < 651)) {
                        if v62 {
                            v4 = v4
                            switch 0 {
                            case 4294967295:
                                v4 += v4
                                var v83 int = -9223372036854775808
                                v84 := GenS6{F0: fi81}
                                v85 := GenS8{F0: "foo bar", F1: v62, F2: 854}
                                fmt.Println(fi81, v64, v6)
                                _ = v83
                                _ = v84
                                _ = v85
                            }
                            v86 := GenS6{F0: fi81}
                            _ = v86
                        } else {
                            switch "" {
                            case "hello":
                                _ = len(v1)
                            default:
                                var v87 int32 = 487
                                _ = v87
                            }
                            v64 = (v64 | (v64 & v64))
                            var v88 float32 = (-2.25 * v5)
                            _ = v88
                        }
                    } else {
                        var v89 float32 = float32(((1.5 * 3.14159) - 1.5))
                        if v62 {
                            v90 := GenS4{F0: 2147483648, F1: (^(v6))}
                            _ = v90
                        } else {
                            var v91 uint = (v4 - v4)
                            if false {
                                v92 := GenS9{F0: "foo bar", F1: v64}
                                _ = v92
                            }
                            _ = v91
                        }
                        var v93 float64 = (1e10 / (float64((float32(-2.25) / v89)) / 0.1))
                        _ = v89
                        _ = v93
                    }
                    _ = len(v61)
                    var v94 uint32 = 42
                    _ = v94
                case "hello":
                    if (v6 <= (v6 - v6)) {
                        switch "\n\t" {
                        case "":
                            if false {
                                fmt.Println(v63)
                                if true {
                                    if true {
                                        v6 = (v6 + 2)
                                    }
                                }
                                var v95 uint16 = (v64 ^ (^(v64)))
                                _ = v95
                            }
                        default:
                            switch "" {
                            case "hello":
                                if ((v3 % ((-330) | 1)) == v3) {
                                    v64 = (v64 * v64)
                                    if false {
                                        if len(v61) > 0 {
                                            v61[0] = (fi81 << ((4294967295) & 63))
                                        }
                                        v5 = (v5 / 1.0)
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v96 int64 = -1
                                            _ = v96
                                        case "a":
                                            var v97 uint32 = 65535
                                            _ = v97
                                        case "":
                                            v61 = append(v61, v63)
                                        default:
                                            v98 := GenS7{F0: v64, F1: true}
                                            _ = v98
                                        }
                                    }
                                }
                            case "":
                                v99 := GenS10{F0: 90, F1: "hello"}
                                if v62 {
                                    if (119 != 0) {
                                        v57 += 0
                                    } else {
                                        if ("" >= "foo bar") {
                                            var v100 float32 = 0.1
                                            _ = v100
                                        } else {
                                            var v101 int8 = v57
                                            _ = v101
                                        }
                                    }
                                    switch fi81 {
                                    case -2147483648:
                                        var v102 uint8 = 1
                                        _ = v102
                                    case 2147483647:
                                        delete(v2, 2147483648)
                                        v62 = ((^(v64)) <= (^(v64)))
                                        v103 := v1[2147483648]
                                        _ = v103
                                    case 651:
                                        if len(v61) > 0 {
                                            v61[0] = fi81
                                        }
                                    default:
                                        if false {
                                            fmt.Println(v64, v3, v62)
                                            v6 = v6
                                            var v104 bool = (1 >= -2147483648)
                                            v105 := GenS11{F0: v63, F1: "foo bar", F2: "\n\t"}
                                            _ = v104
                                            _ = v105
                                        }
                                    }
                                    v106 := GenS3{F0: (v64 % ((1) | 1)), F1: v3, F2: (v63 & 9223372036854775807)}
                                    _ = v106
                                }
                                switch (v63 ^ (fi82 << ((v4) & 63))) {
                                case -1:
                                    v3 -= 42
                                case 769:
                                    var v107 uint64 = v6
                                    _ = v107
                                }
                                var v108 bool = false
                                _ = v108
                                _ = v99
                            case "foo bar":
                                v5 = float32(1.5)
                            default:
                                v57 = (v57 % (((v57 << ((255) & 63))) | 1))
                            }
                        }
                    }
                case "":
                    var v109 int16 = v3
                    _ = v109
                default:
                    v64 -= v64
                }
            } else {
                delete(v2, 495)
                v6 = v6
            }
            _ = fi82
        }
        switch ("foo bar" + ("a" + "\n\t")) {
        case "foo bar":
            if len(v61) > 0 {
                v61[0] = v63
            }
        case "a":
            v6 -= 42
        case "\n\t":
            fmt.Println(v4)
        }
        _ = fi81
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v7
    _ = v80
}
