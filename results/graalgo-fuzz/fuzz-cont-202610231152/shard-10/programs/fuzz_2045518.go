package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 int; F1 uint; F2 uint16 }
type GenS8 struct { F0 int16; F1 string }
type GenS9 struct { F0 bool; F1 uint64 }
type GenS10 struct { F0 uint; F1 string; F2 int16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int64; F1 bool; F2 bool }
type GenS14 struct { F0 int; F1 string; F2 uint32 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 string; F1 string }
type GenS17 struct { F0 string }
func genFn0(p0 bool) int {
    fmt.Println(p0)
    var v5 float32 = 3.14159
    if true {
        v6 := []int{}
        v5 = 0.0
        v7 := GenS1{F0: 65535, F1: 1}
        _ = v6
        _ = v7
    } else {
        v8 := []int{-996, 4294967295}
        for fi9 := 0; fi9 < 9; fi9++ {
            fmt.Println(fi9)
            if (2 <= 810) {
                switch "\n\t" {
                case "":
                    p0 = ((v5 - v5) == v5)
                default:
                    for fi10 := 0; fi10 < 19; fi10++ {
                        v11 := GenS1{F0: 209, F1: fi9}
                        _ = fi10
                        _ = v11
                    }
                }
            } else {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                if ((p0 && p0) || (p0 && p0)) {
                    fmt.Println(v5, fi9)
                    v12 := GenS2{F0: 1, F1: 546}
                    var v13 uint32 = 2147483647
                    _ = v12
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        p0 = true
                        _ = fi14
                    }
                    if (!p0) {
                        for fi15 := 0; fi15 < 13; fi15++ {
                            if p0 {
                                if p0 {
                                    v8 = append(v8, (fi9 ^ (fi15 * (fi15 & (fi9 - fi9)))))
                                    var v16 string = "hello"
                                    fmt.Println(v5, v5, p0)
                                    _ = v16
                                }
                            } else {
                                if false {
                                    var v17 int8 = -97
                                    _ = v17
                                }
                            }
                            var v18 uint8 = 152
                            var v19 int32 = 251
                            _ = fi15
                            _ = v18
                            _ = v19
                        }
                        switch (fi9 << ((703) & 63)) {
                        case 906:
                            fmt.Println(p0)
                        case 127:
                            switch (("" + "a") + "a") {
                            case "a":
                                p0 = true
                            }
                        case 107:
                            _ = len(v8)
                        }
                    }
                }
                for fi20 := 0; fi20 < 6; fi20++ {
                    fmt.Println(fi9)
                    v5 += float32(1e10)
                    _ = fi20
                }
            }
            _ = len(v8)
            _ = fi9
        }
        fmt.Println(v5, v5)
        _ = v8
    }
    v21 := GenS3{F0: -425, F1: 2}
    _ = p0
    _ = v21
    _ = v5
    return 127
}

func main() {
    var v1 uint8 = 255
    v2 := []uint32{}
    var v3 int16 = 1000
    v4 := make(map[uint32]uint32)
    for fi22 := 0; fi22 < 3; fi22++ {
        v3 = (v3 ^ (v3 >> ((886) & 63)))
        fmt.Println(v1)
        _ = fi22
    }
    v23 := GenS3{F0: 9223372036854775807, F1: 65535}
    var v24 string = "a"
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (genFn0(true) >= 4294967295) {
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 int32 = 2147483647
            v27 := GenS4{F0: 65535}
            _ = fi25
            _ = v26
            _ = v27
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if false {
        var v28 float32 = 0.0
        v29 := GenS5{F0: 32768}
        _ = v28
        _ = v29
    } else {
        if true {
            v3 += 42
            if false {
                v30 := v4[799]
                _ = v30
            } else {
                _ = len(v4)
                v31 := GenS6{F0: 255, F1: 2}
                _ = v31
            }
        } else {
            v4[42] = 712
            v24 = v24
            if (-9223372036854775808 == -1) {
                v2 = append(v2, 16)
            } else {
                for fi32 := 0; fi32 < 17; fi32++ {
                    for fi33 := 0; fi33 < 2; fi33++ {
                        v34 := GenS7{F0: fi33, F1: 65535, F2: 255}
                        _ = fi33
                        _ = v34
                    }
                    v35 := GenS8{F0: v3, F1: v24}
                    v1 = 127
                    _ = fi32
                    _ = v35
                }
                v1 -= (v1 * v1)
                if false {
                    switch "a" {
                    case "":
                        v1 += 127
                    case "foo bar":
                        if (1e10 >= 0.0) {
                            v24 = v24
                            fmt.Println(v24, v1, v3)
                            v24 = v24
                        }
                    default:
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                    }
                    for fi36 := 0; fi36 < 16; fi36++ {
                        v37 := GenS8{F0: (v3 / (((v3 - v3)) | 1)), F1: (((("hello" + "\n\t") + v24) + v24) + v24)}
                        for fi38 := 0; fi38 < 20; fi38++ {
                            switch "foo bar" {
                            case "hello":
                                v39 := GenS9{F0: true, F1: 4294967295}
                                _ = v39
                            case "a":
                                v3 = 32767
                            case "foo bar":
                                if true {
                                    var v40 string = v24
                                    v41 := GenS3{F0: (fi38 / (((fi36 << ((18446744073709551615) & 63))) | 1)), F1: 255}
                                    v24 = "a"
                                    _ = v40
                                    _ = v41
                                }
                            default:
                                switch fi38 {
                                case -839:
                                    v3 = (v3 % ((1) | 1))
                                default:
                                    switch v24 {
                                    case "a":
                                        var v42 int64 = int64(v1)
                                        _ = v42
                                    case "foo bar":
                                        fmt.Println(v24, v24, v3)
                                    }
                                }
                            }
                            if (!(1e10 < (0.0 / 1.0))) {
                                v43 := GenS10{F0: 2, F1: "hello", F2: (^(v3))}
                                if true {
                                    if true {
                                        delete(v4, 2147483648)
                                        fmt.Println(v24, v24, v1)
                                        v44 := GenS11{F0: 4294967295}
                                        _ = v44
                                    } else {
                                        v2 = append(v2, 0)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    var v45 string = ""
                                    _ = v45
                                }
                                _ = v43
                            }
                            var v46 int32 = 2147483647
                            if true {
                                if true {
                                    v47 := GenS6{F0: 2, F1: 814}
                                    _ = v47
                                }
                            }
                            v1 = v1
                            _ = fi38
                            _ = v46
                        }
                        _ = fi36
                        _ = v37
                    }
                } else {
                    v1 = (^(v1))
                    v1 -= (v1 * v1)
                    switch (v24 + "\n\t") {
                    case "hello":
                        v48 := v4[360]
                        _ = v48
                    case "":
                        for fi49 := 0; fi49 < 12; fi49++ {
                            if true {
                                var v50 float32 = (((-2.25 * 0.1) * ((3.14159 - (0.0 * 1.5)) / 1.0)) * float32(1e10))
                                fmt.Println(fi49, v1)
                                _ = v50
                            } else {
                                for fi51 := 0; fi51 < 12; fi51++ {
                                    v24 = "a"
                                    _ = fi51
                                }
                                var v52 uint32 = 65535
                                if true {
                                    var v53 uint64 = 93
                                    var v54 uint16 = 65535
                                    _ = v53
                                    _ = v54
                                }
                                _ = v52
                            }
                            for fi55 := 0; fi55 < 19; fi55++ {
                                fmt.Println(v1, fi55)
                                _ = fi55
                            }
                            _ = fi49
                        }
                    case "a":
                        for fi56 := 0; fi56 < 17; fi56++ {
                            var v57 uint32 = 868
                            _ = fi56
                            _ = v57
                        }
                    }
                }
            }
        }
        var v58 uint = 65535
        if (float32(1e10) <= 1e10) {
            if ((^(v58)) >= 65535) {
                var v59 int64 = -9223372036854775808
                for fi60 := 0; fi60 < 18; fi60++ {
                    switch (fi60 / ((382) | 1)) {
                    case 128:
                        switch (fi60 * (fi60 / ((int(v3)) | 1))) {
                        case 4294967295:
                            switch (fi60 << (((^(v58))) & 63)) {
                            case 32768:
                                v61 := GenS12{F0: (-(fi60))}
                                _ = v61
                            default:
                                var v62 uint64 = 205
                                _ = v62
                            }
                        case 2147483647:
                            v1 = (v1 % ((v1) | 1))
                        default:
                            for fi63 := 0; fi63 < 7; fi63++ {
                                fmt.Println(v3)
                                fmt.Println(fi63, v1, fi63)
                                _ = fi63
                            }
                        }
                    case 32767:
                        if false {
                            switch fi60 {
                            case -218:
                                if true {
                                    fmt.Println(v58, v24, v3)
                                } else {
                                    if (!true) {
                                        var v64 uint32 = 1
                                        _ = v64
                                    }
                                    if false {
                                        delete(v4, 42)
                                        switch v24 {
                                        case "":
                                            fmt.Println(v24, v59)
                                        default:
                                            var v65 int64 = (-(v59))
                                            _ = v65
                                        }
                                    } else {
                                        for fi66 := 0; fi66 < 11; fi66++ {
                                            delete(v4, 2147483647)
                                            _ = fi66
                                        }
                                        for fi67 := 0; fi67 < 4; fi67++ {
                                            v24 = ("" + (v24 + "a"))
                                            _ = fi67
                                        }
                                    }
                                    for fi68 := 0; fi68 < 16; fi68++ {
                                        v69 := GenS13{F0: v59, F1: (v1 < 17), F2: (2 == v1)}
                                        _ = fi68
                                        _ = v69
                                    }
                                }
                            default:
                                switch v24 {
                                case "":
                                    for fi70 := 0; fi70 < 17; fi70++ {
                                        switch fi60 {
                                        case 32767:
                                            v71 := GenS14{F0: (-(fi70)), F1: ((v24 + "a") + "foo bar"), F2: 0}
                                            _ = v71
                                        case 2147483647:
                                            fmt.Println(fi60, v1, v3)
                                        default:
                                            var v72 uint8 = v1
                                            _ = v72
                                        }
                                        v58 = 308
                                        _ = fi70
                                    }
                                case "\n\t":
                                    if ((fi60 >> ((v58) & 63)) == (^(fi60))) {
                                        for fi73 := 0; fi73 < 12; fi73++ {
                                            v4[2147483647] = 880
                                            if (false && true) {
                                                var v74 bool = (true || false)
                                                if (v59 <= v59) {
                                                    v3 += (v3 & (^(v3)))
                                                }
                                                _ = v74
                                            } else {
                                                var v75 float32 = 0.0
                                                switch v24 {
                                                case "\n\t":
                                                    v76 := GenS15{F0: v24}
                                                    _ = v76
                                                default:
                                                    v77 := GenS12{F0: fi60}
                                                    _ = v77
                                                }
                                                switch int(v59) {
                                                case -159:
                                                    var v78 uint32 = 255
                                                    _ = v78
                                                case 32768:
                                                    switch ("\n\t" + (("foo bar" + (v24 + "\n\t")) + "hello")) {
                                                    case "":
                                                        _ = len(v2)
                                                        if ((!false) || false) {
                                                            if len(v2) > 0 {
                                                                v2[0] = 119
                                                            }
                                                            var v79 int16 = v3
                                                            _ = v79
                                                        }
                                                        var v80 uint16 = 65535
                                                        _ = v80
                                                    case "hello":
                                                        if (int64(v3) <= v59) {
                                                            v1 -= v1
                                                            fmt.Println(v3)
                                                        }
                                                    case "a":
                                                        v81 := GenS6{F0: 871, F1: 772}
                                                        _ = v81
                                                    default:
                                                        fmt.Println(v58)
                                                    }
                                                    switch fi60 {
                                                    case 127:
                                                        fmt.Println(v1, v59)
                                                    case 42:
                                                        var v82 bool = true
                                                        _ = v82
                                                    }
                                                case -2147483648:
                                                    if (v1 != v1) {
                                                        if false {
                                                            switch ((v24 + "") + v24) {
                                                            case "hello":
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            case "":
                                                                var v83 int64 = (-(v59))
                                                                _ = v83
                                                            case "a":
                                                                v2 = append(v2, 0)
                                                            }
                                                            v84 := GenS16{F0: (v24 + v24), F1: v24}
                                                            v58 = v58
                                                            _ = v84
                                                        }
                                                    } else {
                                                        v75 -= v75
                                                    }
                                                default:
                                                    if len(v2) > 0 {
                                                        v2[0] = 4294967295
                                                    }
                                                }
                                                _ = v75
                                            }
                                            _ = fi73
                                        }
                                        var v85 int8 = 127
                                        _ = v85
                                    }
                                }
                            }
                            var v86 string = ""
                            v2 = append(v2, uint32(v58))
                            _ = v86
                        }
                    case 4294967295:
                        for fi87 := 0; fi87 < 13; fi87++ {
                            var v88 uint64 = 2
                            _ = fi87
                            _ = v88
                        }
                    default:
                        v1 = v1
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi89 := 0; fi89 < 3; fi89++ {
                        if true {
                            v3 += v3
                            v58 = (v58 | (v58 & (v58 * 37)))
                            var v90 bool = (!(4294967295 > fi60))
                            _ = v90
                        }
                        v91 := GenS4{F0: 2147483648}
                        _ = fi89
                        _ = v91
                    }
                    _ = fi60
                }
                _ = v59
            } else {
                var v92 string = ""
                fmt.Println(v24, v58, v1)
                if ((false || (601 != 438)) && true) {
                    switch -307 {
                    case 0:
                        var v93 bool = true
                        _ = v93
                    case 127:
                        var v94 float32 = ((1e10 / float32(-2.25)) + 0.1)
                        _ = v94
                    case 42:
                        var v95 uint8 = uint8(v58)
                        _ = v95
                    }
                    var v96 float32 = 0.1
                    _ = v96
                }
                _ = v92
            }
        }
        _ = v58
    }
    var v97 int32 = -2147483648
    v98 := GenS2{F0: -740, F1: v97}
    var v99 int8 = 42
    for fi100 := 0; fi100 < 5; fi100++ {
        v97 = (v97 - (v97 << ((1) & 63)))
        if false {
            var v101 uint32 = 801
            _ = v101
        } else {
            v99 += v99
            var v102 string = ("\n\t" + (("foo bar" + (v24 + v24)) + v24))
            _ = v102
        }
        for fi103 := 0; fi103 < 12; fi103++ {
            v104 := GenS17{F0: v24}
            _ = fi103
            _ = v104
        }
        _ = fi100
    }
    var v105 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v97:", v97)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v105
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v97
    _ = v98
    _ = v99
}
