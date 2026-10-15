package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int16 }
type GenS2 struct { F0 uint8; F1 int; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 uint32; F1 int64; F2 string }

func main() {
    var v1 uint16 = 65535
    var v2 uint64 = 664
    var v3 string = "a"
    for fi4 := 0; fi4 < 13; fi4++ {
        v2 += 790
        if false {
            v1 = 1
        } else {
            v2 += v2
            switch fi4 {
            case -480:
                for fi5 := 0; fi5 < 7; fi5++ {
                    v3 = v3
                    v2 += v2
                    _ = fi5
                }
            }
            if false {
                v6 := GenS1{F0: 2, F1: 1, F2: 943}
                if true {
                    var v7 uint = 896
                    var v8 int = fi4
                    _ = v7
                    _ = v8
                } else {
                    v2 += v2
                }
                v9 := GenS1{F0: (v2 * v2), F1: -1, F2: 1}
                _ = v6
                _ = v9
            } else {
                v3 = (v3 + (v3 + "hello"))
                var v10 int = fi4
                for fi11 := 0; fi11 < 11; fi11++ {
                    fmt.Println(v1, v10, v10)
                    var v12 uint8 = 127
                    _ = fi11
                    _ = v12
                }
                _ = v10
            }
        }
        _ = fi4
    }
    if (v3 != v3) {
        fmt.Println(v3)
        v13 := []int{}
        v2 += v2
        _ = v13
    } else {
        v14 := GenS1{F0: 50, F1: 1, F2: -3}
        _ = v14
    }
    v1 = uint16(v2)
    var v15 uint8 = 255
    var v16 uint8 = (v15 * v15)
    v16 = (v15 ^ 155)
    fmt.Println(v16, v2)
    if (!true) {
        if (1 < 2147483648) {
            if false {
                v2 -= v2
                fmt.Println(v15)
            }
            if false {
                if ((-2.25 * float64((((0.1 / float32(-2.25)) * (float32(float64(1.5)) - (0.1 / -2.25))) - 3.14159))) < (1.5 - 0.1)) {
                    v17 := GenS1{F0: v2, F1: 1, F2: -1}
                    switch v3 {
                    case "a":
                        var v18 uint8 = v16
                        _ = v18
                    default:
                        v1 = v1
                    }
                    v19 := []int{}
                    _ = v17
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 10; fi20++ {
                        v21 := []int{(^(fi20)), (fi20 ^ 128), (fi20 % (((fi20 / (((fi20 * -65)) | 1))) | 1)), fi20}
                        _ = fi20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 19; fi22++ {
                        v23 := GenS2{F0: 1, F1: (^(fi22)), F2: (v2 & v2)}
                        var v24 bool = true
                        for fi25 := 0; fi25 < 4; fi25++ {
                            var v26 int8 = -15
                            v27 := []int{(-(fi22)), fi25}
                            _ = len(v27)
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    fmt.Println(v2, v16, v15)
                }
                switch "" {
                case "hello":
                    var v28 uint8 = 127
                    _ = v28
                case "foo bar":
                    v2 = v2
                case "":
                    if (!true) {
                        switch v3 {
                        case "\n\t":
                            v16 = (^(v15))
                        case "":
                            var v29 int32 = -2147483648
                            _ = v29
                        case "a":
                            v15 = uint8(v1)
                        default:
                            v30 := GenS3{F0: false}
                            _ = v30
                        }
                        if false {
                            for fi31 := 0; fi31 < 16; fi31++ {
                                switch fi31 {
                                case -9223372036854775808:
                                    var v32 int64 = 477
                                    _ = v32
                                case -2147483648:
                                    if (!true) {
                                        fmt.Println(v3, v16)
                                        if true {
                                            if (!false) {
                                                fmt.Println(v16, v16, v1)
                                            } else {
                                                var v33 int64 = 9223372036854775807
                                                v3 = v3
                                                _ = v33
                                            }
                                            v15 += (v16 << ((255) & 63))
                                            fmt.Println(v3, v16, v3)
                                        }
                                        for fi34 := 0; fi34 < 9; fi34++ {
                                            if (65535 <= v1) {
                                                v35 := GenS4{F0: (0.0 > (1.5 / (3.14159 + (0.0 * (3.14159 - float32((((0.1 * ((1e10 * 3.14159) + 1.5)) + 0.0) * float64(((1e10 + 3.14159) / 0.1))))))))), F1: 2}
                                                fmt.Println(v16, v3)
                                                v36 := GenS5{F0: (65535 <= v1), F1: (v1 * v1)}
                                                _ = v35
                                                _ = v36
                                            } else {
                                                fmt.Println(v1, fi31)
                                            }
                                            _ = fi34
                                        }
                                    } else {
                                        v37 := GenS5{F0: true, F1: (^(v1))}
                                        switch (fi31 / ((4294967295) | 1)) {
                                        case 9223372036854775807:
                                            var v38 float32 = (float32(1e10) / -2.25)
                                            _ = v38
                                        default:
                                            fmt.Println(v2, fi31)
                                        }
                                        v15 = (v15 | 155)
                                        _ = v37
                                    }
                                default:
                                    var v39 uint = 18446744073709551615
                                    _ = v39
                                }
                                v40 := []int{fi31, -205}
                                _ = fi31
                                _ = v40
                            }
                            var v41 float32 = ((1.5 - (-2.25 / float32(-2.25))) / 3.14159)
                            v42 := []int{32767}
                            _ = v41
                            _ = v42
                        } else {
                            v43 := []int{-445, 2147483647, 32768}
                            var v44 bool = true
                            v45 := GenS4{F0: v44, F1: 255}
                            _ = v43
                            _ = v44
                            _ = v45
                        }
                        v3 = "hello"
                    }
                default:
                    v46 := []int{0, 32768, -2147483648, 42}
                    _ = v46
                }
            } else {
                for fi47 := 0; fi47 < 13; fi47++ {
                    if false {
                        for fi48 := 0; fi48 < 13; fi48++ {
                            switch v3 {
                            case "":
                                switch ("foo bar" + ("a" + "a")) {
                                case "a":
                                    v16 -= (v15 + (v16 % (((^(v16))) | 1)))
                                case "foo bar":
                                    v49 := []int{(-(fi48)), (fi48 * fi48), 9223372036854775807}
                                    _ = v49
                                default:
                                    switch (-(fi47)) {
                                    case 128:
                                        var v50 int16 = 841
                                        _ = v50
                                    }
                                }
                            case "foo bar":
                                if false {
                                    v16 += (v15 ^ 1)
                                    fmt.Println(fi47)
                                    fmt.Println(v15, v3)
                                    v51 := GenS5{F0: true, F1: 0}
                                    var v52 int64 = -314
                                    if false {
                                        if (0 <= uint32(v1)) {
                                            v53 := []int{}
                                            _ = v53
                                        } else {
                                            fmt.Println(v16)
                                            v54 := []int{24, (fi47 * fi48)}
                                            fmt.Println(fi48)
                                            _ = v54
                                        }
                                    } else {
                                        v55 := GenS6{F0: 255}
                                        var v56 uint8 = 0
                                        v57 := []int{fi48}
                                        _ = v55
                                        _ = v56
                                        _ = v57
                                    }
                                    _ = v51
                                    _ = v52
                                } else {
                                    v58 := GenS1{F0: 4294967295, F1: 0, F2: -451}
                                    switch (v3 + ((v3 + "\n\t") + v3)) {
                                    case "":
                                        if false {
                                            var v59 int64 = -1
                                            v1 += v1
                                            fmt.Println(v59)
                                            _ = v59
                                        } else {
                                            v2 = (v2 ^ v2)
                                            v60 := GenS5{F0: false, F1: (v1 / ((32767) | 1))}
                                            v1 = (^(v1))
                                            _ = v60
                                        }
                                    default:
                                        v61 := GenS5{F0: true, F1: 42}
                                        _ = v61
                                    }
                                    _ = v58
                                }
                            default:
                                v62 := []int{fi47}
                                _ = v62
                            }
                            v2 -= 0
                            var v63 int32 = -2147483648
                            _ = fi48
                            _ = v63
                        }
                        v2 -= (v2 << ((1) & 63))
                        var v64 float64 = (1.5 - (1.5 / (0.0 + (1.5 - 1e10))))
                        _ = v64
                    }
                    _ = fi47
                }
                if (388 == 2147483647) {
                    if ((-9223372036854775808 < 128) && (((-2.25 - 3.14159) - float64((1e10 + float32(0.1)))) == float64(float32(0.0)))) {
                        var v65 uint64 = v2
                        _ = v65
                    } else {
                        var v66 int = -246
                        _ = v66
                    }
                    v67 := []int{}
                    _ = v67
                } else {
                    v68 := GenS7{F0: -1, F1: 159}
                    _ = v68
                }
            }
        }
    } else {
        if true {
            v69 := []int{-683, 32768, 32768}
            fmt.Println(v2, v1)
            if (!((^(v1)) == (v1 & uint16(v15)))) {
                if len(v69) >= 2 {
                    v69 = v69[1:len(v69)]
                }
                if true {
                    for fi70 := 0; fi70 < 16; fi70++ {
                        v15 -= v15
                        if ((0 > v15) || false) {
                            v1 += v1
                            switch 303 {
                            case 294:
                                if false {
                                    switch ("\n\t" + v3) {
                                    case "foo bar":
                                        if false {
                                            if true {
                                                v3 = v3
                                                v15 = (v16 & 192)
                                                v71 := GenS7{F0: int8(v2), F1: (v1 >> ((1) & 63))}
                                                _ = v71
                                            } else {
                                                v72 := GenS6{F0: 2}
                                                _ = v72
                                            }
                                            fmt.Println(v3)
                                            v16 -= uint8(v1)
                                        }
                                    default:
                                        switch 128 {
                                        case 32767:
                                            v16 = 42
                                        case 143:
                                            var v73 bool = true
                                            _ = v73
                                        default:
                                            v16 -= v16
                                        }
                                    }
                                }
                            case 709:
                                _ = len(v69)
                            }
                            for fi74 := 0; fi74 < 15; fi74++ {
                                var v75 int = (fi74 % (((fi74 / ((fi74) | 1))) | 1))
                                switch (v75 / (((v75 | v75)) | 1)) {
                                case 4294967295:
                                    v1 -= v1
                                case 42:
                                    var v76 float32 = 3.14159
                                    _ = v76
                                }
                                if (-2147483648 < 204) {
                                    var v77 uint16 = 42
                                    var v78 uint64 = (v2 & (v2 / ((42) | 1)))
                                    fmt.Println(v75, fi74, v75)
                                    v75 += fi74
                                    if len(v69) > 0 {
                                        v69[0] = (v75 >> ((352) & 63))
                                    }
                                    if len(v69) > 0 {
                                        v69[0] = (-(fi74))
                                    }
                                    _ = v77
                                    _ = v78
                                } else {
                                    if (1 <= 65535) {
                                        if true {
                                            if true {
                                                v3 = v3
                                                v3 = ((v3 + ("hello" + ("" + (v3 + v3)))) + v3)
                                                v75 -= (fi70 + 128)
                                            }
                                        } else {
                                            v15 -= (v16 + v15)
                                            var v79 uint = 552
                                            _ = v79
                                        }
                                    }
                                    v80 := GenS4{F0: (65535 <= 42), F1: 255}
                                    switch v3 {
                                    case "foo bar":
                                        fmt.Println(v2, v75)
                                        v15 -= v16
                                        if (-9223372036854775808 >= -791) {
                                            v15 += 1
                                        } else {
                                            if false {
                                                v81 := GenS4{F0: (0.0 != 0.1), F1: 2}
                                                switch v3 {
                                                case "":
                                                    switch fi74 {
                                                    case 128:
                                                        switch v3 {
                                                        case "foo bar":
                                                            switch "" {
                                                            case "\n\t":
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    if (v1 > (v1 % (((v1 / ((42) | 1))) | 1))) {
                                                                        v3 = (v3 + (v3 + ((("\n\t" + v3) + (("\n\t" + (("hello" + v3) + (v3 + (v3 + v3)))) + v3)) + ("a" + v3))))
                                                                        if len(v69) >= 2 {
                                                                            v69 = v69[1:len(v69)]
                                                                        }
                                                                    } else {
                                                                        v69 = append(v69, fi70)
                                                                        v75 = v75
                                                                    }
                                                                }
                                                            default:
                                                                if false {
                                                                    var v82 uint64 = (v2 << ((18446744073709551615) & 63))
                                                                    fmt.Println(fi70, v2)
                                                                    switch ("hello" + (v3 + "")) {
                                                                    case "":
                                                                        var v83 int64 = -1
                                                                        _ = v83
                                                                    case "foo bar":
                                                                        var v84 int = fi74
                                                                        switch (v84 << ((4294967295) & 63)) {
                                                                        case -941:
                                                                            if ((v16 >> ((2) & 63)) <= (^(v16))) {
                                                                                fmt.Println(v3)
                                                                            } else {
                                                                                var v85 int32 = 874
                                                                                fmt.Println(v82, v15)
                                                                                _ = v85
                                                                            }
                                                                        case 128:
                                                                            var v86 bool = (false && true)
                                                                            _ = v86
                                                                        }
                                                                        _ = v84
                                                                    case "\n\t":
                                                                        var v87 float64 = ((1.5 - (1.5 * float64(((0.0 * (0.0 * ((1.5 * -2.25) - 1e10))) / 0.1)))) / 3.14159)
                                                                        _ = v87
                                                                    }
                                                                    _ = v82
                                                                } else {
                                                                    v88 := GenS4{F0: (!true), F1: 2}
                                                                    if len(v69) >= 2 {
                                                                        v69 = v69[1:len(v69)]
                                                                    }
                                                                    v1 += (v1 + v1)
                                                                    fmt.Println(v16)
                                                                    _ = v88
                                                                }
                                                            }
                                                        default:
                                                            _ = len(v69)
                                                        }
                                                    default:
                                                        var v89 uint64 = (v2 & 344)
                                                        _ = v89
                                                    }
                                                }
                                                var v90 int64 = -1
                                                _ = v81
                                                _ = v90
                                            } else {
                                                if true {
                                                    if len(v69) >= 2 {
                                                        v69 = v69[1:len(v69)]
                                                    }
                                                }
                                                fmt.Println(fi74, v75)
                                                if (((true || true) || (false && true)) && true) {
                                                    if true {
                                                        var v91 uint16 = (v1 + v1)
                                                        _ = v91
                                                    }
                                                    var v92 int = (-(fi70))
                                                    v93 := GenS3{F0: false}
                                                    var v94 uint32 = 2147483647
                                                    switch (v75 / (((fi74 | v75)) | 1)) {
                                                    case 32768:
                                                        v15 -= 71
                                                    }
                                                    _ = v92
                                                    _ = v93
                                                    _ = v94
                                                }
                                                v95 := GenS6{F0: 237}
                                                _ = v95
                                            }
                                            v2 = 698
                                            fmt.Println(v15, v75, v1)
                                            var v96 int8 = 100
                                            _ = v96
                                        }
                                        if len(v69) >= 2 {
                                            v69 = v69[1:len(v69)]
                                        }
                                        fmt.Println(fi70, v1, v1)
                                    case "a":
                                        var v97 int = 32768
                                        _ = v97
                                    case "hello":
                                        v3 = v3
                                    default:
                                        v69 = append(v69, 2147483647)
                                    }
                                    var v98 int64 = int64(fi74)
                                    _ = v80
                                    _ = v98
                                }
                                _ = fi74
                                _ = v75
                            }
                        }
                        if true {
                            for fi99 := 0; fi99 < 14; fi99++ {
                                if true {
                                    fmt.Println(v15, v16)
                                    switch v3 {
                                    case "a":
                                        fmt.Println(fi70)
                                    }
                                    v16 += (v16 / ((v16) | 1))
                                    v100 := GenS8{F0: 2147483647, F1: 1, F2: "foo bar"}
                                    _ = v100
                                } else {
                                    var v101 float64 = (float64(float32(0.0)) * 3.14159)
                                    v69 = append(v69, fi70)
                                    _ = v101
                                }
                                fmt.Println(v3)
                                switch "" {
                                case "foo bar":
                                    v102 := GenS7{F0: 127, F1: 94}
                                    _ = v102
                                default:
                                    fmt.Println(v3, v1, v3)
                                }
                                if (!((!false) || true)) {
                                    fmt.Println(fi70, v16)
                                    var v103 int16 = 32767
                                    _ = v103
                                } else {
                                    v1 = v1
                                    fmt.Println(v15)
                                }
                                _ = fi99
                            }
                        } else {
                            var v104 int8 = 1
                            v105 := GenS5{F0: false, F1: (v1 | v1)}
                            v106 := GenS3{F0: true}
                            _ = v104
                            _ = v105
                            _ = v106
                        }
                        _ = fi70
                    }
                    for fi107 := 0; fi107 < 5; fi107++ {
                        for fi108 := 0; fi108 < 14; fi108++ {
                            var v109 int64 = -9223372036854775808
                            v16 = v16
                            var v110 uint16 = v1
                            _ = fi108
                            _ = v109
                            _ = v110
                        }
                        _ = fi107
                    }
                    for fi111 := 0; fi111 < 5; fi111++ {
                        v3 = ""
                        v1 += (v1 >> ((1) & 63))
                        _ = fi111
                    }
                }
            } else {
                var v112 uint64 = (v2 % ((v2) | 1))
                _ = v112
            }
            _ = v69
        }
        v113 := []int{1}
        v3 = "hello"
        _ = v113
    }
    for fi114 := 0; fi114 < 12; fi114++ {
        v115 := GenS6{F0: (v15 ^ v16)}
        v116 := GenS1{F0: v2, F1: -2147483648, F2: 1}
        var v117 float32 = -2.25
        _ = fi114
        _ = v115
        _ = v116
        _ = v117
    }
    v118 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v118:", len(v118), v118)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v118
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
