package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS3 struct { F0 int16; F1 int8; F2 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 int8; F1 int32; F2 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 bool; F1 int64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int64; F1 int16; F2 uint8 }
type GenS12 struct { F0 uint8; F1 int; F2 uint32 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 string; F1 string; F2 int8 }
type GenS15 struct { F0 string; F1 bool }
type GenS16 struct { F0 uint; F1 int64; F2 uint32 }
type GenS17 struct { F0 string }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 int; F1 int32 }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 bool; F1 int32 }
type GenS23 struct { F0 string; F1 int64; F2 string }
type GenS24 struct { F0 string; F1 int16 }
type GenS25 struct { F0 uint16 }
func genFn0(p0 uint8) int16 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 4; fi5++ {
        p0 = (^(p0))
        v6 := GenS1{F0: 65535, F1: (p0 | p0), F2: 0}
        var v7 uint8 = (p0 ^ p0)
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := []int{fi9, 9223372036854775807, 127, (fi8 % (((^(fi9))) | 1))}
            _ = fi9
            _ = v10
        }
        fmt.Println(fi8)
        _ = fi8
    }
    _ = p0
    return 42
}
func genFn1(p0 string, p1 int64, p2 uint64) int16 {
    if false {
        var v11 int = -335
        fmt.Println(p1, v11)
        _ = v11
    }
    v12 := GenS2{F0: p2, F1: 1, F2: 138}
    v13 := []int{}
    for fi14 := 0; fi14 < 7; fi14++ {
        if false {
            switch (((p0 + (p0 + ((p0 + "") + p0))) + ((p0 + p0) + p0)) + p0) {
            case "foo bar":
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            }
            switch (fi14 & fi14) {
            case 275:
                for fi15 := 0; fi15 < 5; fi15++ {
                    var v16 float64 = 0.0
                    _ = fi15
                    _ = v16
                }
            case 32767:
                v17 := GenS2{F0: 4294967295, F1: 42, F2: 42}
                _ = v17
            case 199:
                p1 -= (-(p1))
            }
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        }
        for fi18 := 0; fi18 < 18; fi18++ {
            if false {
                fmt.Println(fi14, p0, p2)
                if len(v13) > 0 {
                    v13[0] = fi18
                }
                if (float32(0.1) != 0.0) {
                    var v19 uint64 = (p2 + p2)
                    _ = v19
                }
            }
            _ = fi18
        }
        _ = fi14
    }
    if len(v13) > 0 {
        v13[0] = 42
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 1
}

func main() {
    var v1 uint32 = 2
    var v2 float64 = (-2.25 / 3.14159)
    var v3 uint64 = 18446744073709551615
    var v4 uint8 = 67
    var v20 float64 = (v2 + float64(1e10))
    fmt.Println(v4, v4)
    v21 := []int{}
    for fi22 := 0; fi22 < 18; fi22++ {
        if true {
            v23 := GenS2{F0: (v3 ^ (v3 - v3)), F1: -2147483648, F2: v4}
            if false {
                if (false || (!false)) {
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    v2 = v20
                }
                v24 := GenS3{F0: genFn0((v4 & v4)), F1: -1, F2: 100}
                for fi25 := 0; fi25 < 9; fi25++ {
                    var v26 bool = false
                    v21 = append(v21, 128)
                    _ = fi25
                    _ = v26
                }
                _ = v24
            } else {
                var v27 string = "hello"
                fmt.Println(v2)
                _ = v27
            }
            v28 := GenS3{F0: -618, F1: 80, F2: 107}
            _ = v23
            _ = v28
        } else {
            v29 := GenS4{F0: fi22}
            _ = v29
        }
        _ = fi22
    }
    if false {
        v20 -= 0.1
        v30 := GenS4{F0: 128}
        v20 = (((v2 + v2) / (v2 / 1.0)) - ((v2 / v2) / (v20 - v2)))
        _ = v30
    } else {
        switch 0 {
        case 1:
            var v31 string = (("foo bar" + "hello") + "\n\t")
            _ = v31
        case 936:
            var v32 string = "\n\t"
            _ = v32
        default:
            var v33 uint16 = 824
            _ = v33
        }
        v3 = 934
    }
    var v34 int8 = -1
    var v35 float64 = v20
    var v36 float64 = v2
    if (false && (9223372036854775807 < -1)) {
        v1 = v1
    } else {
        if true {
            v21 = append(v21, -9223372036854775808)
        }
        v21 = append(v21, 9223372036854775807)
        switch 4294967295 {
        case 42:
            if true {
                for fi37 := 0; fi37 < 19; fi37++ {
                    for fi38 := 0; fi38 < 8; fi38++ {
                        var v39 uint = 65535
                        v40 := GenS5{F0: v1, F1: 9223372036854775807}
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    if (!(!true)) {
                        for fi41 := 0; fi41 < 15; fi41++ {
                            if true {
                                v42 := GenS6{F0: v34, F1: -1, F2: ("foo bar" + "foo bar")}
                                _ = v42
                            } else {
                                if (false && ("" <= (("foo bar" + ("" + "\n\t")) + "a"))) {
                                    if ("a" < ("\n\t" + "foo bar")) {
                                        v43 := GenS4{F0: fi41}
                                        var v44 int32 = 2147483647
                                        _ = v43
                                        _ = v44
                                    }
                                }
                                v45 := GenS7{F0: 255}
                                switch "hello" {
                                case "foo bar":
                                    if true {
                                        var v46 int16 = 42
                                        v47 := GenS2{F0: uint64(v1), F1: -1, F2: v4}
                                        v48 := GenS8{F0: (fi41 >> ((410) & 63)), F1: -1}
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    }
                                case "\n\t":
                                    var v49 int16 = 0
                                    _ = v49
                                }
                                _ = v45
                            }
                            _ = fi41
                        }
                        if len(v21) > 0 {
                            v21[0] = fi37
                        }
                    }
                    _ = fi37
                }
                v50 := GenS7{F0: uint16(v1)}
                _ = v50
            } else {
                v34 += v34
                if (!(false && true)) {
                    v20 -= ((((((v36 * v35) - v36) + (((-2.25 / v20) / v35) + (v20 / (v36 / v2)))) / 3.14159) + v20) - v2)
                    if false {
                        for fi51 := 0; fi51 < 8; fi51++ {
                            v52 := GenS9{F0: false, F1: 1}
                            var v53 bool = (32767 < -202)
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                    }
                } else {
                    var v54 uint64 = v3
                    if (("\n\t" + "a") >= (((("a" + "\n\t") + "") + "") + ("" + "a"))) {
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                    } else {
                        v55 := GenS10{F0: 9223372036854775807}
                        switch 42 {
                        case -2147483648:
                            switch "a" {
                            case "":
                                switch 127 {
                                case 0:
                                    v35 += 0.1
                                case 127:
                                    var v56 int = 1
                                    _ = v56
                                case 362:
                                    for fi57 := 0; fi57 < 4; fi57++ {
                                        switch fi57 {
                                        case 128:
                                            var v58 int32 = -179
                                            _ = v58
                                        case 9223372036854775807:
                                            _ = len(v21)
                                        case 4294967295:
                                            v59 := GenS10{F0: 42}
                                            _ = v59
                                        }
                                        if (1000 == -259) {
                                            for fi60 := 0; fi60 < 5; fi60++ {
                                                fmt.Println(v54)
                                                var v61 int64 = 42
                                                fmt.Println(v61)
                                                _ = fi60
                                                _ = v61
                                            }
                                            if true {
                                                var v62 uint64 = 613
                                                switch (fi57 << ((255) & 63)) {
                                                case 1:
                                                    if true {
                                                        for fi63 := 0; fi63 < 16; fi63++ {
                                                            if len(v21) > 0 {
                                                                v21[0] = fi57
                                                            }
                                                            _ = fi63
                                                        }
                                                        for fi64 := 0; fi64 < 7; fi64++ {
                                                            fmt.Println(v62, v54)
                                                            var v65 string = "a"
                                                            _ = fi64
                                                            _ = v65
                                                        }
                                                        v54 = v54
                                                    } else {
                                                        for fi66 := 0; fi66 < 14; fi66++ {
                                                            _ = len(v21)
                                                            _ = fi66
                                                        }
                                                    }
                                                case -718:
                                                    v54 += (v62 | v62)
                                                case 2147483647:
                                                    switch (fi57 >> ((2) & 63)) {
                                                    case -275:
                                                        if true {
                                                            for fi67 := 0; fi67 < 2; fi67++ {
                                                                if (true && true) {
                                                                    v1 -= v1
                                                                }
                                                                switch "a" {
                                                                case "\n\t":
                                                                    var v68 uint = 256
                                                                    _ = v68
                                                                case "":
                                                                    switch (fi67 ^ 288) {
                                                                    case 4294967295:
                                                                        fmt.Println(v3, v62)
                                                                    }
                                                                case "foo bar":
                                                                    v69 := GenS11{F0: 765, F1: -478, F2: (v4 - (v4 | 42))}
                                                                    _ = v69
                                                                }
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    var v70 uint = 90
                                                                    _ = v70
                                                                case "":
                                                                    fmt.Println(v34)
                                                                default:
                                                                    v54 -= v54
                                                                }
                                                                _ = fi67
                                                            }
                                                            var v71 uint = 672
                                                            v72 := GenS12{F0: (v4 / ((183) | 1)), F1: 32768, F2: v1}
                                                            _ = v71
                                                            _ = v72
                                                        } else {
                                                            if true {
                                                                if true {
                                                                    var v73 float32 = (float32((float64(1e10) - v2)) - (float32(1e10) * ((float32((((0.0 - 0.1) / ((-2.25 + v20) * ((((float64((0.0 - 0.0)) * (v2 + v2)) - (-2.25 + v20)) + v2) + v20))) * (v2 - v20))) + 3.14159) - float32(((v2 / (v36 * 1.5)) / (v2 / (((v20 + v35) * v35) - v20)))))))
                                                                    var v74 int16 = 32767
                                                                    _ = v73
                                                                    _ = v74
                                                                } else {
                                                                    v2 += (v36 / 1e10)
                                                                    var v75 int16 = -32768
                                                                    var v76 uint16 = 42
                                                                    _ = v75
                                                                    _ = v76
                                                                }
                                                                var v77 int16 = 1
                                                                v78 := GenS13{F0: (v62 - (v62 - v3))}
                                                                _ = v77
                                                                _ = v78
                                                            } else {
                                                                v4 = 128
                                                                var v79 uint = 65535
                                                                fmt.Println(v54, v20)
                                                                _ = v79
                                                            }
                                                            var v80 uint32 = (v1 / ((4294967295) | 1))
                                                            var v81 bool = false
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                    case -2147483648:
                                                        v3 = 1
                                                    case 119:
                                                        var v82 int8 = v34
                                                        _ = v82
                                                    default:
                                                        switch (fi57 >> ((18446744073709551615) & 63)) {
                                                        case 0:
                                                            if true {
                                                                if true {
                                                                    var v83 uint8 = 0
                                                                    if false {
                                                                        v84 := GenS13{F0: 866}
                                                                        fmt.Println(v35, v4)
                                                                        _ = v84
                                                                    }
                                                                    _ = v83
                                                                }
                                                                for fi85 := 0; fi85 < 11; fi85++ {
                                                                    v86 := GenS3{F0: 212, F1: v34, F2: (v34 / ((v34) | 1))}
                                                                    fmt.Println(v54, v4)
                                                                    _ = fi85
                                                                    _ = v86
                                                                }
                                                                switch (fi57 ^ 9223372036854775807) {
                                                                case 32767:
                                                                    var v87 int64 = -9223372036854775808
                                                                    _ = v87
                                                                case 128:
                                                                    fmt.Println(v3)
                                                                case 1:
                                                                    if (683 != int32(v4)) {
                                                                        v1 -= 4294967295
                                                                    } else {
                                                                        for fi88 := 0; fi88 < 14; fi88++ {
                                                                            if len(v21) > 0 {
                                                                                v21[0] = -177
                                                                            }
                                                                            switch fi88 {
                                                                            case 4294967295:
                                                                                switch -544 {
                                                                                case 1:
                                                                                    fmt.Println(v2)
                                                                                default:
                                                                                    fmt.Println(v1)
                                                                                }
                                                                            case 42:
                                                                                v89 := GenS14{F0: "hello", F1: "hello", F2: (v34 ^ v34)}
                                                                                v3 -= (v62 / (((v62 ^ v3)) | 1))
                                                                                v90 := GenS4{F0: fi57}
                                                                                _ = v89
                                                                                _ = v90
                                                                            case 1:
                                                                                var v91 bool = (-1 >= -1)
                                                                                _ = v91
                                                                            default:
                                                                                fmt.Println(v3, v4)
                                                                            }
                                                                            v92 := GenS3{F0: 104, F1: v34, F2: v34}
                                                                            var v93 float32 = 0.0
                                                                            fmt.Println(v1)
                                                                            _ = fi88
                                                                            _ = v92
                                                                            _ = v93
                                                                        }
                                                                        v3 = 2
                                                                    }
                                                                default:
                                                                    fmt.Println(v20)
                                                                }
                                                            } else {
                                                                v20 = v20
                                                                switch fi57 {
                                                                case 4294967295:
                                                                    v21 = append(v21, (fi57 * fi57))
                                                                }
                                                                v94 := GenS12{F0: (v4 + 87), F1: (fi57 - (fi57 << ((4294967295) & 63))), F2: (^(v1))}
                                                                _ = v94
                                                            }
                                                        case -2147483648:
                                                            v34 += (v34 & -1)
                                                        default:
                                                            var v95 bool = false
                                                            _ = v95
                                                        }
                                                    }
                                                default:
                                                    var v96 int8 = (v34 << ((uint(v62)) & 63))
                                                    _ = v96
                                                }
                                                var v97 float64 = float64(3.14159)
                                                _ = v62
                                                _ = v97
                                            }
                                            var v98 bool = (v1 < 2147483647)
                                            _ = v98
                                        }
                                        if false {
                                            switch "\n\t" {
                                            case "a":
                                                if ((v4 + (^(v4))) < v4) {
                                                    if len(v21) >= 2 {
                                                        v21 = v21[1:len(v21)]
                                                    }
                                                    switch "hello" {
                                                    case "":
                                                        var v99 uint = 42
                                                        _ = v99
                                                    }
                                                    if (1 > -614) {
                                                        v100 := GenS1{F0: (v1 - (v1 << ((18446744073709551615) & 63))), F1: v4, F2: -2147483648}
                                                        var v101 string = ""
                                                        _ = v100
                                                        _ = v101
                                                    }
                                                }
                                            }
                                            if false {
                                                v102 := GenS15{F0: "", F1: (!(true || true))}
                                                if ((542 != v1) && ((fi57 << ((439) & 63)) == fi57)) {
                                                    for fi103 := 0; fi103 < 14; fi103++ {
                                                        v104 := GenS16{F0: 18446744073709551615, F1: 164, F2: (v1 | v1)}
                                                        var v105 int = (fi103 >> ((255) & 63))
                                                        _ = fi103
                                                        _ = v104
                                                        _ = v105
                                                    }
                                                    if false {
                                                        _ = len(v21)
                                                        v2 -= ((v35 / float64(1e10)) - v20)
                                                    } else {
                                                        if len(v21) > 0 {
                                                            v21[0] = (^(fi57))
                                                        }
                                                        var v106 bool = (!false)
                                                        if true {
                                                            if (-613 < 1000) {
                                                                if len(v21) > 0 {
                                                                    v21[0] = 306
                                                                }
                                                                fmt.Println(v1)
                                                            } else {
                                                                var v107 int16 = -32768
                                                                for fi108 := 0; fi108 < 9; fi108++ {
                                                                    if (v106 && (((0.1 - ((float32((0.1 * v2)) + 0.1) + 1e10)) - 1.5) != (-2.25 / 1.5))) {
                                                                        v34 = (v34 * (v34 | (v34 * v34)))
                                                                        var v109 string = "a"
                                                                        _ = v109
                                                                    }
                                                                    v21 = append(v21, (fi57 >> ((4294967295) & 63)))
                                                                    _ = fi108
                                                                }
                                                                v34 += v34
                                                                _ = v107
                                                            }
                                                            fmt.Println(v54, fi57)
                                                            var v110 int32 = 42
                                                            _ = v110
                                                        }
                                                        _ = v106
                                                    }
                                                } else {
                                                    if false {
                                                        if len(v21) >= 2 {
                                                            v21 = v21[1:len(v21)]
                                                        }
                                                        switch fi57 {
                                                        case 128:
                                                            for fi111 := 0; fi111 < 4; fi111++ {
                                                                if true {
                                                                    if len(v21) > 0 {
                                                                        v21[0] = (fi57 << ((42) & 63))
                                                                    }
                                                                    v35 -= 0.0
                                                                    v112 := GenS8{F0: fi57, F1: 1}
                                                                    _ = v112
                                                                } else {
                                                                    if (!false) {
                                                                        fmt.Println(v36, v2, v34)
                                                                        v20 = v2
                                                                    }
                                                                    var v113 uint32 = uint32(v54)
                                                                    if false {
                                                                        var v114 int = 32768
                                                                        var v115 float64 = (v2 - float64((0.1 - 0.1)))
                                                                        _ = v114
                                                                        _ = v115
                                                                    }
                                                                    if false {
                                                                        switch (fi57 - fi111) {
                                                                        case 32767:
                                                                            v116 := GenS17{F0: "foo bar"}
                                                                            _ = v116
                                                                        case -921:
                                                                            if (453 <= 774) {
                                                                                if (false && ((^(v3)) < (v54 >> ((0) & 63)))) {
                                                                                    if false {
                                                                                        var v117 bool = false
                                                                                        _ = v117
                                                                                    } else {
                                                                                        if (v113 == 4294967295) {
                                                                                            fmt.Println(v2, v54)
                                                                                        } else {
                                                                                            if (!(65535 > uint16(v4))) {
                                                                                                fmt.Println(v4)
                                                                                                v34 -= (v34 | v34)
                                                                                                var v118 int16 = 393
                                                                                                fmt.Println(v34, fi57)
                                                                                                _ = v118
                                                                                            }
                                                                                            v34 = (v34 + v34)
                                                                                            var v119 float64 = 1e10
                                                                                            _ = v119
                                                                                        }
                                                                                        if len(v21) >= 2 {
                                                                                            v21 = v21[1:len(v21)]
                                                                                        }
                                                                                    }
                                                                                    v36 += v20
                                                                                    v3 -= 333
                                                                                }
                                                                                fmt.Println(v4)
                                                                            }
                                                                        default:
                                                                            switch fi111 {
                                                                            case -9223372036854775808:
                                                                                if false {
                                                                                    if true {
                                                                                        switch fi57 {
                                                                                        case 769:
                                                                                            var v120 float64 = v2
                                                                                            _ = v120
                                                                                        case 4294967295:
                                                                                            var v121 int32 = -1
                                                                                            _ = v121
                                                                                        }
                                                                                    } else {
                                                                                        var v122 uint64 = (v54 / (((v54 ^ (v3 + (v54 % (((v3 << ((0) & 63))) | 1))))) | 1))
                                                                                        _ = v122
                                                                                    }
                                                                                }
                                                                            case 4294967295:
                                                                                v2 = v2
                                                                            default:
                                                                                v3 += (v3 / ((v3) | 1))
                                                                            }
                                                                        }
                                                                        v123 := GenS18{F0: 37}
                                                                        _ = v123
                                                                    }
                                                                    v113 = (v113 % ((v1) | 1))
                                                                    _ = v113
                                                                }
                                                                if (fi57 == fi111) {
                                                                    v124 := GenS15{F0: "", F1: true}
                                                                    fmt.Println(fi57, v36)
                                                                    v2 = 1e10
                                                                    var v125 int16 = int16(v34)
                                                                    _ = v124
                                                                    _ = v125
                                                                } else {
                                                                    if len(v21) > 0 {
                                                                        v21[0] = fi57
                                                                    }
                                                                }
                                                                v34 = v34
                                                                _ = fi111
                                                            }
                                                        default:
                                                            fmt.Println(v2)
                                                        }
                                                    } else {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v20 -= 3.14159
                                                        case "hello":
                                                            _ = len(v21)
                                                        }
                                                        v126 := GenS15{F0: ("" + "hello"), F1: (!(!true))}
                                                        for fi127 := 0; fi127 < 4; fi127++ {
                                                            fmt.Println(v54, v54, v35)
                                                            _ = fi127
                                                        }
                                                        _ = v126
                                                    }
                                                    v34 += (v34 % ((v34) | 1))
                                                }
                                                v4 = (v4 | (^(v4)))
                                                _ = v102
                                            } else {
                                                switch "a" {
                                                case "a":
                                                    if false {
                                                        var v128 float32 = (-2.25 * 3.14159)
                                                        if len(v21) > 0 {
                                                            v21[0] = -2147483648
                                                        }
                                                        if ((true && true) && (0 != 255)) {
                                                            if true {
                                                                _ = len(v21)
                                                            }
                                                            fmt.Println(v4, v3, v128)
                                                            if false {
                                                                var v129 uint16 = 65535
                                                                var v130 uint = 469
                                                                if true {
                                                                    if (false && (!((v129 == 65535) || true))) {
                                                                        if true {
                                                                            v131 := GenS12{F0: v4, F1: (fi57 % ((fi57) | 1)), F2: v1}
                                                                            _ = v131
                                                                        } else {
                                                                            v132 := GenS15{F0: "a", F1: (v34 >= (v34 + (v34 * int8(v4))))}
                                                                            var v133 float32 = 0.0
                                                                            _ = v132
                                                                            _ = v133
                                                                        }
                                                                        if true {
                                                                            if true {
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    var v134 int64 = 42
                                                                                    _ = v134
                                                                                case "hello":
                                                                                    v135 := GenS5{F0: 4294967295, F1: 9223372036854775807}
                                                                                    _ = v135
                                                                                default:
                                                                                    for fi136 := 0; fi136 < 8; fi136++ {
                                                                                        var v137 uint16 = uint16(v1)
                                                                                        v138 := GenS18{F0: -1}
                                                                                        _ = fi136
                                                                                        _ = v137
                                                                                        _ = v138
                                                                                    }
                                                                                }
                                                                                fmt.Println(v35)
                                                                                var v139 uint16 = (v129 | v129)
                                                                                _ = v139
                                                                            }
                                                                        } else {
                                                                            var v140 float32 = (((v128 * v128) / v128) + (3.14159 - (3.14159 - (0.0 * v128))))
                                                                            _ = len(v21)
                                                                            fmt.Println(v35, v20)
                                                                            _ = v140
                                                                        }
                                                                    }
                                                                }
                                                                _ = v129
                                                                _ = v130
                                                            } else {
                                                                if len(v21) > 0 {
                                                                    v21[0] = -1
                                                                }
                                                                fmt.Println(v36, v1, v34)
                                                                var v141 int16 = 895
                                                                _ = v141
                                                            }
                                                        } else {
                                                            v4 = (v4 >> ((2) & 63))
                                                        }
                                                        _ = v128
                                                    } else {
                                                        v142 := GenS19{F0: int16(v1)}
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            for fi143 := 0; fi143 < 16; fi143++ {
                                                                var v144 int = (fi57 * (-(fi57)))
                                                                fmt.Println(v36, v34)
                                                                _ = len(v21)
                                                                _ = fi143
                                                                _ = v144
                                                            }
                                                        case "\n\t":
                                                            fmt.Println(v3, v2)
                                                        case "":
                                                            v1 = (v1 + (^(v1)))
                                                        }
                                                        var v145 float64 = 0.0
                                                        _ = v142
                                                        _ = v145
                                                    }
                                                default:
                                                    v21 = append(v21, (fi57 % ((fi57) | 1)))
                                                }
                                            }
                                        } else {
                                            v146 := GenS3{F0: 0, F1: (v34 * (v34 % (((v34 ^ (v34 | 0))) | 1))), F2: v34}
                                            _ = v146
                                        }
                                        _ = fi57
                                    }
                                default:
                                    var v147 float32 = 0.1
                                    _ = v147
                                }
                            case "hello":
                                v36 += (((-2.25 * v36) * 0.0) - v2)
                            case "a":
                                v4 = 2
                            default:
                                for fi148 := 0; fi148 < 12; fi148++ {
                                    if true {
                                        _ = len(v21)
                                        if (true || (!false)) {
                                            v149 := GenS20{F0: 458, F1: 1}
                                            var v150 float32 = (1.5 - float32(v35))
                                            if (((((false || true) && (-849 != 0)) && true) && ((1 == 1) && ((v36 - v2) != v20))) && true) {
                                                var v151 int = -9223372036854775808
                                                if true {
                                                    v20 += v35
                                                    switch (fi148 | 32768) {
                                                    case 2147483647:
                                                        v3 = v3
                                                    case -9223372036854775808:
                                                        v2 += float64(float32(3.14159))
                                                    default:
                                                        v3 -= 0
                                                    }
                                                } else {
                                                    var v152 int16 = -1
                                                    _ = v152
                                                }
                                                _ = v151
                                            }
                                            _ = v149
                                            _ = v150
                                        }
                                        v153 := GenS17{F0: "a"}
                                        _ = v153
                                    }
                                    if (false || true) {
                                        for fi154 := 0; fi154 < 4; fi154++ {
                                            if false {
                                                v3 -= v54
                                                fmt.Println(fi148, v4, v34)
                                                v155 := GenS21{F0: 75}
                                                _ = v155
                                            } else {
                                                var v156 uint = 911
                                                v2 = (-2.25 - v2)
                                                v1 += v1
                                                _ = v156
                                            }
                                            if false {
                                                var v157 int32 = -1
                                                v158 := GenS21{F0: 65535}
                                                fmt.Println(v2)
                                                _ = v157
                                                _ = v158
                                            } else {
                                                if (!(2147483647 > 42)) {
                                                    fmt.Println(fi148)
                                                    if true {
                                                        var v159 bool = false
                                                        v160 := GenS8{F0: fi148, F1: -1}
                                                        _ = v159
                                                        _ = v160
                                                    } else {
                                                        var v161 string = "foo bar"
                                                        _ = v161
                                                    }
                                                    fmt.Println(fi148, v4, fi148)
                                                }
                                            }
                                            _ = fi154
                                        }
                                    } else {
                                        if (!(!(v1 != 319))) {
                                            if false {
                                                v34 += (v34 - (v34 >> ((255) & 63)))
                                                v20 -= float64(0.1)
                                                var v162 int8 = (v34 ^ v34)
                                                _ = v162
                                            }
                                            var v163 uint32 = v1
                                            for fi164 := 0; fi164 < 5; fi164++ {
                                                v54 += (v3 | (v54 ^ 53))
                                                var v165 uint32 = v1
                                                v1 += v165
                                                _ = fi164
                                                _ = v165
                                            }
                                            _ = v163
                                        } else {
                                            v3 = v54
                                            for fi166 := 0; fi166 < 4; fi166++ {
                                                switch (^(fi148)) {
                                                case 1:
                                                    if false {
                                                        v167 := GenS21{F0: 42}
                                                        if false {
                                                            fmt.Println(v3, fi148)
                                                            switch "" {
                                                            case "\n\t":
                                                                _ = len(v21)
                                                            case "hello":
                                                                v36 += v20
                                                            case "":
                                                                var v168 uint = 909
                                                                _ = v168
                                                            default:
                                                                var v169 uint64 = (v3 % ((v3) | 1))
                                                                var v170 int8 = (v34 % ((v34) | 1))
                                                                if false {
                                                                    fmt.Println(v36, v35, v20)
                                                                    var v171 uint32 = 2
                                                                    _ = v171
                                                                } else {
                                                                    var v172 bool = true
                                                                    v173 := GenS22{F0: true, F1: -1}
                                                                    if false {
                                                                        if len(v21) >= 2 {
                                                                            v21 = v21[1:len(v21)]
                                                                        }
                                                                    }
                                                                    _ = v172
                                                                    _ = v173
                                                                }
                                                                _ = v169
                                                                _ = v170
                                                            }
                                                        } else {
                                                            var v174 int16 = 0
                                                            v175 := GenS23{F0: "a", F1: -1, F2: "foo bar"}
                                                            _ = v174
                                                            _ = v175
                                                        }
                                                        _ = v167
                                                    }
                                                }
                                                v2 += 3.14159
                                                if true {
                                                    v21 = append(v21, fi166)
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v176 := GenS24{F0: "", F1: 405}
                                                        _ = v176
                                                    }
                                                }
                                                _ = fi166
                                            }
                                            v20 = (v36 + (-2.25 - ((-2.25 * float64(float32(float64(0.1)))) - (1.5 + v20))))
                                        }
                                    }
                                    if true {
                                        switch "foo bar" {
                                        case "\n\t":
                                            fmt.Println(v36, v54)
                                        case "hello":
                                            var v177 uint64 = v54
                                            _ = v177
                                        default:
                                            v34 = -72
                                        }
                                        if len(v21) >= 2 {
                                            v21 = v21[1:len(v21)]
                                        }
                                    }
                                    _ = fi148
                                }
                            }
                        default:
                            v178 := GenS19{F0: 1}
                            _ = v178
                        }
                        if true {
                            fmt.Println(v34, v36)
                            v2 += v36
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                if len(v21) > 0 {
                                    v21[0] = -9223372036854775808
                                }
                            default:
                                var v179 string = "\n\t"
                                _ = v179
                            }
                            switch "hello" {
                            case "hello":
                                if (v1 == (v1 << ((255) & 63))) {
                                    for fi180 := 0; fi180 < 5; fi180++ {
                                        v2 += v20
                                        _ = len(v21)
                                        _ = fi180
                                    }
                                }
                            }
                        }
                        _ = v55
                    }
                    v181 := GenS25{F0: 32767}
                    _ = v181
                    _ = v54
                }
            }
        }
    }
    var v182 int = -2147483648
    v36 += ((1e10 * v20) * v36)
    var v183 string = ("hello" + "")
    var v184 uint = uint(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v182:", v182)
    fmt.Println("v183:", v183)
    fmt.Println("v184:", v184)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v182
    _ = v183
    _ = v184
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
