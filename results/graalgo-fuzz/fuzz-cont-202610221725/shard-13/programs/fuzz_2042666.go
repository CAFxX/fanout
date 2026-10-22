package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint; F2 int8 }
type GenS5 struct { F0 string; F1 int; F2 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 string; F1 int8; F2 int }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint; F1 uint }
type GenS12 struct { F0 int8; F1 uint8 }
type GenS13 struct { F0 uint32; F1 int16 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int16; F1 uint }
type GenS16 struct { F0 uint; F1 string }
type GenS17 struct { F0 string; F1 uint; F2 uint32 }
type GenS18 struct { F0 bool; F1 int16 }
type GenS19 struct { F0 int; F1 int32 }
type GenS20 struct { F0 uint8 }
type GenS21 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS22 struct { F0 uint64; F1 uint64; F2 uint }
type GenS23 struct { F0 int8 }
type GenS24 struct { F0 uint8; F1 uint32 }
type GenS25 struct { F0 int32 }
type GenS26 struct { F0 int16; F1 int8; F2 int64 }
type GenS27 struct { F0 int8; F1 uint8 }
type GenS28 struct { F0 string }
func genFn0(p0 int32, p1 int16, p2 uint64) int {
    v5 := GenS1{F0: 42}
    switch ("foo bar" + ("\n\t" + "\n\t")) {
    case "":
        if false {
            v6 := []int{}
            p0 = (p0 * p0)
            _ = v6
        } else {
            v7 := []int{127, 2147483647, 2147483647, int(p1)}
            if true {
                for fi8 := 0; fi8 < 1; fi8++ {
                    fmt.Println(p1)
                    var v9 uint = 615
                    v10 := GenS2{F0: 0, F1: 977}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                var v11 uint64 = p2
                _ = v11
            }
            var v12 string = ("" + "\n\t")
            _ = v12
            _ = v7
        }
    }
    v13 := []int{}
    if false {
        var v14 uint64 = (p2 + (p2 << ((609) & 63)))
        v14 = 42
        if false {
            var v15 int16 = -1
            v14 += (^(v14))
            v16 := GenS2{F0: 1, F1: 4294967295}
            _ = v15
            _ = v16
        } else {
            if (127 >= 255) {
                var v17 float64 = 0.1
                _ = v17
            }
        }
        _ = v14
    }
    v18 := GenS2{F0: 1, F1: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v18
    _ = v5
    return 809
}
func genFn1(p0 int64, p1 uint64, p2 int) uint8 {
    fmt.Println(p1, p2, p1)
    fmt.Println(p2, p2)
    for fi19 := 0; fi19 < 14; fi19++ {
        for fi20 := 0; fi20 < 14; fi20++ {
            if ((p1 | (p1 - p1)) != p1) {
                v21 := GenS3{F0: -602}
                fmt.Println(fi19, fi19)
                _ = v21
            } else {
                var v22 int = -2147483648
                v23 := []int{(fi19 % ((-581) | 1)), (fi20 | genFn0(-2147483648, 1, 18446744073709551615))}
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        _ = fi19
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    v1 := make(map[int]int)
    var v2 int = 128
    var v3 int32 = 2147483647
    var v4 uint8 = 1
    if (2 >= 241) {
        var v24 int = v2
        v25 := GenS3{F0: v3}
        _ = v24
        _ = v25
    }
    v26 := []int{(v2 - (v2 | 4294967295)), 42, genFn0(-1, 1000, 18446744073709551615), genFn0(v3, 1000, 2)}
    var v27 int32 = -2147483648
    var v28 int16 = 1000
    for fi29 := 0; fi29 < 18; fi29++ {
        if (true || (!false)) {
            v28 -= -683
            var v30 uint64 = 1
            fmt.Println(v28)
            _ = v30
        }
        if (!(0.0 <= 0.0)) {
            v3 = v3
            _ = len(v1)
            for fi31 := 0; fi31 < 1; fi31++ {
                fmt.Println(v3, fi31)
                if true {
                    fmt.Println(v3, v4)
                    switch "\n\t" {
                    case "foo bar":
                        if len(v26) > 0 {
                            v26[0] = fi29
                        }
                    case "\n\t":
                        v32 := GenS4{F0: false, F1: 255, F2: 100}
                        _ = v32
                    case "a":
                        v33 := GenS5{F0: "", F1: genFn0(int32(v28), 1000, 403), F2: 0}
                        _ = v33
                    default:
                        switch v2 {
                        case 32768:
                            var v34 int16 = (v28 / ((v28) | 1))
                            _ = v34
                        case 1:
                            var v35 float64 = 1e10
                            _ = v35
                        case -9223372036854775808:
                            var v36 int8 = -128
                            _ = v36
                        }
                    }
                    var v37 float64 = (1.5 * 3.14159)
                    _ = v37
                }
                _ = fi31
            }
        }
        _ = fi29
    }
    if false {
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
    } else {
        if (0.1 < (float64((-2.25 + (float32(3.14159) - 1e10))) / 0.1)) {
            fmt.Println(v4)
            for fi38 := 0; fi38 < 7; fi38++ {
                v4 += v4
                if true {
                    var v39 uint = 2
                    if ((!true) && ((1e10 > 0.1) && (!(!true)))) {
                        if false {
                            switch "" {
                            case "foo bar":
                                if ((!true) || false) {
                                    if true {
                                        v40 := GenS6{F0: (!(!true))}
                                        if (1 != 42) {
                                            v4 -= (v4 ^ v4)
                                            if false {
                                                v41 := GenS7{F0: 546}
                                                for fi42 := 0; fi42 < 15; fi42++ {
                                                    fmt.Println(v27, v39, v27)
                                                    v3 -= int32(fi42)
                                                    _ = fi42
                                                }
                                                for fi43 := 0; fi43 < 15; fi43++ {
                                                    if (!true) {
                                                        var v44 uint64 = 4294967295
                                                        v45 := GenS7{F0: v28}
                                                        _ = len(v1)
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                    if (255 != uint16(v39)) {
                                                        var v46 float64 = 0.1
                                                        switch v2 {
                                                        case 0:
                                                            v2 -= genFn0(v3, v28, 896)
                                                        case 42:
                                                            var v47 bool = true
                                                            _ = v47
                                                        case -9223372036854775808:
                                                            v48 := GenS6{F0: (v4 <= v4)}
                                                            _ = v48
                                                        default:
                                                            var v49 string = "foo bar"
                                                            _ = v49
                                                        }
                                                        if false {
                                                            fmt.Println(v3, v28)
                                                            if false {
                                                                if false {
                                                                    if true {
                                                                        v27 -= (v3 & v27)
                                                                        var v50 int8 = 0
                                                                        v50 = -56
                                                                        if (((((("" + "foo bar") + ("hello" + "\n\t")) + "a") + ("hello" + "a")) + ("hello" + "\n\t")) > "hello") {
                                                                            _ = len(v1)
                                                                            v51 := GenS1{F0: 1}
                                                                            _ = v51
                                                                        }
                                                                        v2 = -227
                                                                        var v52 string = "\n\t"
                                                                        _ = v50
                                                                        _ = v52
                                                                    } else {
                                                                        v39 -= 859
                                                                    }
                                                                    v53 := GenS8{F0: 18446744073709551615}
                                                                    _ = v53
                                                                }
                                                                fmt.Println(v39, v3)
                                                                v26 = append(v26, (^(v2)))
                                                            }
                                                            var v54 float64 = (((float64(3.14159) * 0.0) / (v46 + (v46 + float64(0.0)))) / v46)
                                                            _ = v54
                                                        }
                                                        _ = v46
                                                    } else {
                                                        if false {
                                                            if (((v27 - int32(v4)) < v3) || false) {
                                                                fmt.Println(v3, v27)
                                                                if false {
                                                                    fmt.Println(v28, fi43)
                                                                    _ = len(v26)
                                                                } else {
                                                                    v39 = (v39 & (v39 ^ (v39 >> ((v39) & 63))))
                                                                    if (914 == 65535) {
                                                                        var v55 uint = v39
                                                                        _ = v55
                                                                    }
                                                                    var v56 int64 = -9223372036854775808
                                                                    _ = v56
                                                                }
                                                                var v57 string = ((("\n\t" + ("foo bar" + "foo bar")) + "foo bar") + "\n\t")
                                                                _ = v57
                                                            } else {
                                                                switch 32768 {
                                                                case 720:
                                                                    var v58 bool = (!false)
                                                                    _ = v58
                                                                case -75:
                                                                    var v59 string = "\n\t"
                                                                    _ = v59
                                                                case 9223372036854775807:
                                                                    delete(v1, 42)
                                                                }
                                                                switch (^(fi43)) {
                                                                case 127:
                                                                    switch (("hello" + "hello") + "hello") {
                                                                    case "":
                                                                        var v60 float64 = 1.5
                                                                        _ = v60
                                                                    case "hello":
                                                                        v2 -= (v2 >> ((v39) & 63))
                                                                    default:
                                                                        v61 := GenS9{F0: ("" + "a"), F1: -128, F2: 304}
                                                                        _ = v61
                                                                    }
                                                                }
                                                                switch "hello" {
                                                                case "a":
                                                                    if (255 < (v39 * 0)) {
                                                                        v62 := GenS9{F0: "hello", F1: 1, F2: (fi43 ^ -150)}
                                                                        _ = v62
                                                                    } else {
                                                                        fmt.Println(fi38, v3, v28)
                                                                        v39 -= uint(v4)
                                                                        var v63 int16 = -631
                                                                        if true {
                                                                            v28 = v28
                                                                            var v64 int16 = v63
                                                                            var v65 uint = v39
                                                                            _ = v64
                                                                            _ = v65
                                                                        } else {
                                                                            v2 = (fi43 & (v2 & (fi38 / (((^(fi43))) | 1))))
                                                                            v39 -= (^(v39))
                                                                            v63 = v63
                                                                        }
                                                                        _ = v63
                                                                    }
                                                                case "hello":
                                                                    if len(v26) > 0 {
                                                                        v26[0] = 42
                                                                    }
                                                                case "foo bar":
                                                                    var v66 int16 = v28
                                                                    _ = v66
                                                                default:
                                                                    v67 := GenS7{F0: v28}
                                                                    _ = v67
                                                                }
                                                                v2 = (^(fi38))
                                                                v28 = v28
                                                            }
                                                            var v68 float64 = float64(3.14159)
                                                            _ = v68
                                                        }
                                                    }
                                                    if (1 < 9223372036854775807) {
                                                        var v69 uint64 = 931
                                                        if true {
                                                            var v70 int = (fi43 + (^(fi43)))
                                                            fmt.Println(v28)
                                                            _ = v70
                                                        }
                                                        if ("\n\t" != "") {
                                                            v71 := GenS1{F0: 715}
                                                            _ = v71
                                                        }
                                                        _ = v69
                                                    }
                                                    _ = fi43
                                                }
                                                _ = v41
                                            } else {
                                                for fi72 := 0; fi72 < 16; fi72++ {
                                                    v4 -= (v4 & v4)
                                                    var v73 int32 = (-(v3))
                                                    v3 -= (v73 & v3)
                                                    v74 := GenS3{F0: (v27 >> ((v39) & 63))}
                                                    _ = fi72
                                                    _ = v73
                                                    _ = v74
                                                }
                                            }
                                        } else {
                                            if ((v39 ^ (v39 << ((12) & 63))) == (v39 >> (((v39 | (v39 | v39))) & 63))) {
                                                v2 = (-(v2))
                                                v75 := v1[(v2 / ((fi38) | 1))]
                                                _ = v75
                                            } else {
                                                v76 := GenS10{F0: (v3 | -1)}
                                                switch fi38 {
                                                case 383:
                                                    if (-128 >= -1) {
                                                        v27 = 639
                                                        v3 = (v27 - (v3 | (v3 / (((v27 << ((42) & 63))) | 1))))
                                                        fmt.Println(v4, v2, fi38)
                                                    }
                                                case -733:
                                                    if (false && ((v28 << ((v39) & 63)) >= -32768)) {
                                                        var v77 uint = v39
                                                        _ = v77
                                                    }
                                                case 2147483647:
                                                    var v78 uint64 = 18446744073709551615
                                                    _ = v78
                                                default:
                                                    var v79 uint16 = 42
                                                    _ = v79
                                                }
                                                _ = v76
                                            }
                                            for fi80 := 0; fi80 < 8; fi80++ {
                                                v81 := GenS5{F0: "a", F1: (fi80 + v2), F2: 2}
                                                v82 := GenS10{F0: v3}
                                                v3 -= v3
                                                v83 := GenS6{F0: false}
                                                if len(v26) >= 2 {
                                                    v26 = v26[1:len(v26)]
                                                }
                                                v84 := GenS5{F0: "a", F1: -9223372036854775808, F2: 338}
                                                var v85 int32 = (v3 >> ((uint(v4)) & 63))
                                                _ = fi80
                                                _ = v81
                                                _ = v82
                                                _ = v83
                                                _ = v84
                                                _ = v85
                                            }
                                            v86 := GenS11{F0: 0, F1: (^(v39))}
                                            _ = v86
                                        }
                                        _ = v40
                                    } else {
                                        switch "\n\t" {
                                        case "\n\t":
                                            switch (-(fi38)) {
                                            case -2147483648:
                                                _ = len(v26)
                                            case 697:
                                                fmt.Println(v39)
                                            case 276:
                                                if (false || false) {
                                                    delete(v1, 526)
                                                }
                                            }
                                        case "a":
                                            if true {
                                                if (((v2 & (fi38 * (fi38 | int(v3)))) != (fi38 | (-(fi38)))) && (v39 > (v39 | v39))) {
                                                    v87 := GenS12{F0: -22, F1: (v4 * 42)}
                                                    var v88 int64 = 585
                                                    _ = v87
                                                    _ = v88
                                                }
                                                if true {
                                                    v89 := GenS2{F0: (v4 / ((v4) | 1)), F1: 906}
                                                    v39 = (v39 << (((v39 ^ v39)) & 63))
                                                    var v90 int64 = -776
                                                    _ = v89
                                                    _ = v90
                                                } else {
                                                    v2 = genFn0(-294, -32768, 1)
                                                    fmt.Println(v2, fi38, v2)
                                                    var v91 uint64 = uint64(v4)
                                                    _ = v91
                                                }
                                                v4 = 42
                                            } else {
                                                if true {
                                                    var v92 float32 = ((float32(0.1) / 0.1) / float32((float64((1e10 - -2.25)) * 3.14159)))
                                                    _ = v92
                                                }
                                            }
                                        case "foo bar":
                                            var v93 uint8 = v4
                                            _ = v93
                                        default:
                                            fmt.Println(v2, fi38, v27)
                                        }
                                        if false {
                                            v94 := v1[genFn0((-(v3)), v28, 4294967295)]
                                            _ = v94
                                        } else {
                                            if ((("" + "a") + "foo bar") >= ("\n\t" + "")) {
                                                if (((^(v3)) < v3) || true) {
                                                    v39 = v39
                                                    var v95 float64 = (0.0 + float64(-2.25))
                                                    var v96 uint = v39
                                                    _ = v95
                                                    _ = v96
                                                }
                                            } else {
                                                var v97 bool = false
                                                var v98 float32 = (1e10 * (3.14159 / ((float32(3.14159) * 0.1) * (1.5 + 3.14159))))
                                                _ = len(v26)
                                                _ = v97
                                                _ = v98
                                            }
                                            if true {
                                                var v99 uint8 = (v4 / (((v4 + v4)) | 1))
                                                _ = v99
                                            }
                                            v100 := GenS5{F0: "hello", F1: 1, F2: 2147483647}
                                            _ = v100
                                        }
                                    }
                                    if len(v26) > 0 {
                                        v26[0] = v2
                                    }
                                }
                            case "":
                                v101 := GenS3{F0: (v3 / ((v27) | 1))}
                                _ = v101
                            case "\n\t":
                                switch "foo bar" {
                                case "a":
                                    var v102 string = "a"
                                    _ = v102
                                default:
                                    v27 += v27
                                }
                            default:
                                var v103 uint = (v39 / (((v39 - 4294967295)) | 1))
                                _ = v103
                            }
                        } else {
                            fmt.Println(v28, v2)
                        }
                        if (-1 != -6) {
                            var v104 string = "\n\t"
                            var v105 bool = false
                            if ((!(v39 <= (v39 - 255))) || v105) {
                                var v106 bool = true
                                var v107 uint64 = 18446744073709551615
                                var v108 int8 = 47
                                _ = v106
                                _ = v107
                                _ = v108
                            }
                            _ = v104
                            _ = v105
                        } else {
                            if (1 > v39) {
                                if true {
                                    v109 := GenS13{F0: 0, F1: v28}
                                    _ = v109
                                }
                            }
                            for fi110 := 0; fi110 < 7; fi110++ {
                                var v111 bool = false
                                if (((1.5 - 3.14159) / 1.0) > 0.1) {
                                    switch "a" {
                                    case "foo bar":
                                        if true {
                                            v3 = (v27 ^ v27)
                                            v111 = (0.1 <= (-2.25 + float64((1.5 - 3.14159))))
                                        }
                                    case "hello":
                                        v112 := GenS9{F0: "", F1: 127, F2: 608}
                                        var v113 uint8 = v4
                                        _ = v112
                                        _ = v113
                                    }
                                    if (0.0 >= 0.1) {
                                        var v114 int8 = -49
                                        var v115 uint8 = v4
                                        switch (v2 << ((2) & 63)) {
                                        case 32768:
                                            v116 := GenS1{F0: uint16(v39)}
                                            _ = v116
                                        case -1:
                                            var v117 uint = (v39 | uint(v4))
                                            switch fi110 {
                                            case -2147483648:
                                                fmt.Println(v27, v117)
                                            case -9223372036854775808:
                                                var v118 bool = (255 >= 32767)
                                                _ = v118
                                            case -255:
                                                if v111 {
                                                    v115 = (^(v115))
                                                    delete(v1, (v2 % ((0) | 1)))
                                                }
                                            }
                                            v119 := GenS14{F0: "foo bar"}
                                            fmt.Println(v114, v3)
                                            _ = v117
                                            _ = v119
                                        case 2147483647:
                                            v28 += 0
                                        }
                                        _ = v114
                                        _ = v115
                                    } else {
                                        v28 = (v28 | v28)
                                        v120 := GenS15{F0: v28, F1: v39}
                                        v121 := GenS7{F0: v28}
                                        _ = v120
                                        _ = v121
                                    }
                                    v111 = false
                                    if (false && (2 != 0)) {
                                        var v122 int32 = (-(v27))
                                        switch ("a" + "a") {
                                        case "":
                                            switch fi110 {
                                            case 109:
                                                fmt.Println(v28, v4)
                                            case 32767:
                                                v111 = (42 != (v122 - v122))
                                            case -9223372036854775808:
                                                var v123 uint = v39
                                                _ = v123
                                            default:
                                                v111 = (((v4 / ((v4) | 1)) >= v4) && v111)
                                            }
                                        default:
                                            if (-128 == -1) {
                                                var v124 uint32 = 1
                                                v39 += (v39 / (((v39 | v39)) | 1))
                                                var v125 uint8 = (^(v4))
                                                _ = v124
                                                _ = v125
                                            }
                                            v39 = v39
                                            if v111 {
                                                switch (-(fi38)) {
                                                case 32768:
                                                    if v111 {
                                                        var v126 int64 = 624
                                                        switch (-(fi38)) {
                                                        case 32767:
                                                            fmt.Println(v3)
                                                        case 4294967295:
                                                            switch fi38 {
                                                            case 436:
                                                                switch (-(v2)) {
                                                                case 32768:
                                                                    var v127 int64 = 190
                                                                    _ = v127
                                                                case 32767:
                                                                    if len(v26) > 0 {
                                                                        v26[0] = -2147483648
                                                                    }
                                                                default:
                                                                    v27 = (v122 >> (((v39 << (((v39 << ((v39) & 63))) & 63))) & 63))
                                                                }
                                                            case 1:
                                                                if len(v26) > 0 {
                                                                    v26[0] = (-(fi110))
                                                                }
                                                            case -2147483648:
                                                                v111 = v111
                                                                v128 := GenS16{F0: v39, F1: ("\n\t" + "\n\t")}
                                                                _ = v128
                                                            }
                                                        case -505:
                                                            var v129 uint8 = (v4 | (^(v4)))
                                                            switch (fi110 + (fi110 ^ genFn0((v122 >> ((v39) & 63)), (v28 + (v28 ^ v28)), 18446744073709551615))) {
                                                            case 1:
                                                                fmt.Println(v3, v126, v122)
                                                            default:
                                                                if v111 {
                                                                    fmt.Println(v4, fi110)
                                                                    if (("a" + (("a" + "foo bar") + "\n\t")) == "hello") {
                                                                        v27 = (v122 / (((v122 | (v27 ^ v3))) | 1))
                                                                        v28 -= (v28 | v28)
                                                                        fmt.Println(v111, v27, v39)
                                                                    }
                                                                    var v130 int64 = (v126 / (((^(v126))) | 1))
                                                                    var v131 uint16 = 27
                                                                    _ = v130
                                                                    _ = v131
                                                                } else {
                                                                    var v132 int = fi110
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        var v133 uint = v39
                                                                        _ = v133
                                                                    case "a":
                                                                        switch (^(fi38)) {
                                                                        case 32768:
                                                                            fmt.Println(v3, v4)
                                                                        case -9223372036854775808:
                                                                            _ = len(v26)
                                                                        case 127:
                                                                            v26 = append(v26, genFn0(v3, 1000, 11))
                                                                            var v134 int16 = v28
                                                                            if (v111 && v111) {
                                                                                v129 = 168
                                                                            }
                                                                            _ = v134
                                                                        }
                                                                    }
                                                                    _ = v132
                                                                }
                                                            }
                                                            _ = v129
                                                        }
                                                        v27 -= 798
                                                        _ = v126
                                                    } else {
                                                        if (!(((1.5 - float32(float64(float32(1.5)))) <= 0.0) || (!(2147483648 == 42)))) {
                                                            var v135 float32 = 1e10
                                                            if v111 {
                                                                var v136 int16 = v28
                                                                var v137 int16 = v28
                                                                _ = v136
                                                                _ = v137
                                                            }
                                                            _ = v135
                                                        } else {
                                                            fmt.Println(v2)
                                                            var v138 uint8 = genFn1(1, 778, (^(v2)))
                                                            v39 -= 42
                                                            _ = v138
                                                        }
                                                    }
                                                    var v139 uint32 = 4294967295
                                                    _ = v139
                                                default:
                                                    _ = len(v1)
                                                }
                                                fmt.Println(fi110, v2, fi38)
                                                if (v39 == (v39 ^ (^(v39)))) {
                                                    v111 = ((-2.25 * 0.1) <= (3.14159 / 1e10))
                                                    v122 += (v27 >> ((v39) & 63))
                                                }
                                                v27 += (v122 % ((v122) | 1))
                                            }
                                        }
                                        if len(v26) > 0 {
                                            v26[0] = -1
                                        }
                                        _ = v122
                                    } else {
                                        v3 += v27
                                    }
                                }
                                fmt.Println(v27)
                                var v140 int8 = 127
                                switch -9223372036854775808 {
                                case 640:
                                    if ((v27 >> (((v39 - 1)) & 63)) != (-(v3))) {
                                        fmt.Println(v4, fi38)
                                        switch (("foo bar" + "a") + ("foo bar" + ("\n\t" + ""))) {
                                        case "\n\t":
                                            var v141 bool = v111
                                            _ = v141
                                        }
                                    } else {
                                        v3 = 286
                                        if false {
                                            v142 := v1[(^(v2))]
                                            _ = v142
                                        }
                                        v143 := GenS15{F0: v28, F1: v39}
                                        if (!v111) {
                                            var v144 int64 = -1
                                            if (false && ((v4 > 42) || true)) {
                                                switch "a" {
                                                case "hello":
                                                    fmt.Println(v4)
                                                    fmt.Println(v4, v111, v111)
                                                    switch genFn0((v3 | v3), 1, 42) {
                                                    case 42:
                                                        _ = len(v26)
                                                    }
                                                case "\n\t":
                                                    var v145 uint16 = 42
                                                    _ = v145
                                                case "":
                                                    switch "hello" {
                                                    case "\n\t":
                                                        delete(v1, fi110)
                                                    case "foo bar":
                                                        fmt.Println(v4, v144)
                                                    case "hello":
                                                        v146 := v1[v2]
                                                        _ = v146
                                                    }
                                                }
                                                fmt.Println(v28)
                                            } else {
                                                v140 = (v140 + (^(v140)))
                                            }
                                            _ = v144
                                        } else {
                                            v111 = true
                                            v147 := GenS17{F0: "\n\t", F1: 18446744073709551615, F2: 2147483647}
                                            _ = v147
                                        }
                                        v3 = (^(v3))
                                        if v111 {
                                            _ = len(v26)
                                        }
                                        v39 += (v39 >> ((v39) & 63))
                                        if false {
                                            fmt.Println(v140)
                                            var v148 uint16 = 110
                                            v149 := GenS10{F0: 1}
                                            v150 := GenS18{F0: false, F1: (v28 & (v28 % ((0) | 1)))}
                                            _ = v148
                                            _ = v149
                                            _ = v150
                                        }
                                        fmt.Println(v140, fi38)
                                        switch "" {
                                        case "hello":
                                            fmt.Println(v3)
                                        case "\n\t":
                                            v1[(^(fi38))] = 128
                                        }
                                        _ = v143
                                    }
                                default:
                                    v3 = 8
                                }
                                switch "\n\t" {
                                case "hello":
                                    v111 = false
                                case "":
                                    _ = len(v1)
                                }
                                _ = fi110
                                _ = v111
                                _ = v140
                            }
                            v1[fi38] = 238
                        }
                        v1[genFn0(v3, 1000, 144)] = (fi38 / (((fi38 % ((v2) | 1))) | 1))
                    } else {
                        var v151 int8 = 97
                        v28 -= (v28 ^ int16(fi38))
                        _ = v151
                    }
                    for fi152 := 0; fi152 < 1; fi152++ {
                        var v153 float64 = 3.14159
                        _ = fi152
                        _ = v153
                    }
                    _ = v39
                }
                _ = fi38
            }
        } else {
            fmt.Println(v3, v28)
            for fi154 := 0; fi154 < 15; fi154++ {
                var v155 int32 = (^(v27))
                v2 += (^(v2))
                _ = fi154
                _ = v155
            }
            v2 += v2
        }
        v156 := GenS12{F0: 0, F1: v4}
        v4 = v4
        _ = v156
    }
    for fi157 := 0; fi157 < 5; fi157++ {
        if true {
            v158 := GenS3{F0: v3}
            _ = v158
        } else {
            _ = len(v1)
            switch "\n\t" {
            case "a":
                if true {
                    v28 += 767
                    if false {
                        v159 := GenS19{F0: fi157, F1: 1}
                        _ = v159
                    }
                }
            case "foo bar":
                switch (v2 - (fi157 << ((470) & 63))) {
                case 128:
                    var v160 int16 = (v28 % ((int16(v2)) | 1))
                    _ = v160
                case -1:
                    switch v2 {
                    case 75:
                        if true {
                            var v161 uint64 = 115
                            if len(v26) >= 2 {
                                v26 = v26[1:len(v26)]
                            }
                            fmt.Println(v28, v161)
                            _ = v161
                        }
                    case 128:
                        if len(v26) > 0 {
                            v26[0] = 128
                        }
                    case -1:
                        var v162 int64 = 0
                        _ = v162
                    }
                }
            case "":
                v163 := GenS20{F0: 95}
                _ = v163
            default:
                if true {
                    fmt.Println(v4, v2)
                    _ = len(v26)
                    v164 := v1[v2]
                    _ = v164
                } else {
                    if false {
                        for fi165 := 0; fi165 < 12; fi165++ {
                            fmt.Println(v28)
                            _ = fi165
                        }
                        if (3.14159 < (-2.25 * 3.14159)) {
                            if (18446744073709551615 > 4294967295) {
                                if (!false) {
                                    var v166 uint32 = 65535
                                    _ = v166
                                } else {
                                    fmt.Println(v4, v4, v3)
                                }
                            } else {
                                fmt.Println(v27, v27, v28)
                            }
                        } else {
                            if true {
                                var v167 int = (v2 | genFn0(v3, 42, 42))
                                for fi168 := 0; fi168 < 18; fi168++ {
                                    switch (fi157 - (v167 & genFn0(42, v28, 0))) {
                                    case -716:
                                        var v169 int16 = (v28 % ((42) | 1))
                                        _ = v169
                                    case -2147483648:
                                        if true {
                                            v3 += (^(v27))
                                            v170 := GenS11{F0: 0, F1: 2}
                                            if len(v26) >= 2 {
                                                v26 = v26[1:len(v26)]
                                            }
                                            _ = v170
                                        }
                                    default:
                                        v1[genFn0(42, v28, 550)] = (v2 | genFn0((v27 % (((v3 % ((v3) | 1))) | 1)), (-(v28)), 18446744073709551615))
                                    }
                                    var v171 string = (("foo bar" + "hello") + ("a" + ""))
                                    _ = fi168
                                    _ = v171
                                }
                                _ = v167
                            }
                        }
                    } else {
                        var v172 bool = true
                        _ = v172
                    }
                    var v173 string = (("" + "foo bar") + "hello")
                    for fi174 := 0; fi174 < 13; fi174++ {
                        v175 := GenS7{F0: -1}
                        fmt.Println(v27, v3, v2)
                        switch v173 {
                        case "a":
                            v27 = 1
                        case "\n\t":
                            if true {
                                v27 -= (v27 >> ((42) & 63))
                                fmt.Println(v28)
                                var v176 bool = (236 != 42)
                                _ = v176
                            }
                        default:
                            v27 = 42
                        }
                        _ = fi174
                        _ = v175
                    }
                    _ = v173
                }
            }
        }
        if false {
            v4 = genFn1(-1, 379, (v2 >> ((767) & 63)))
        } else {
            switch -9223372036854775808 {
            case 32768:
                if ((65535 == 65535) && ((!(42 <= 18)) || false)) {
                    var v177 int16 = 32767
                    if (!true) {
                        var v178 bool = (!true)
                        _ = v178
                    }
                    var v179 uint8 = (v4 | (v4 & genFn1(-1, 762, fi157)))
                    _ = v177
                    _ = v179
                } else {
                    if false {
                        for fi180 := 0; fi180 < 1; fi180++ {
                            v27 = 0
                            switch (v2 | genFn0((v27 << ((255) & 63)), v28, 42)) {
                            case 32767:
                                if (0.0 >= ((1.5 + (0.1 + (1.5 / (0.1 * (0.0 - 1.5))))) / 1.0)) {
                                    var v181 uint = 42
                                    if ((18446744073709551615 == (v181 | (v181 ^ v181))) && (42 > 2)) {
                                        switch "\n\t" {
                                        case "":
                                            v181 = (v181 % ((v181) | 1))
                                            v2 = fi180
                                            v1[(-(fi180))] = (-(v2))
                                        }
                                        v181 += (v181 | v181)
                                        fmt.Println(fi180)
                                    } else {
                                        v181 -= 224
                                        v27 = v3
                                        v182 := GenS14{F0: "a"}
                                        _ = v182
                                    }
                                    v3 += (v27 * (v27 | (v27 / ((42) | 1))))
                                    var v183 float32 = (1.5 / 1.5)
                                    _ = len(v26)
                                    fmt.Println(fi157, v3, v3)
                                    if len(v26) > 0 {
                                        v26[0] = v2
                                    }
                                    _ = v181
                                    _ = v183
                                }
                            case 822:
                                switch fi157 {
                                case 32768:
                                    if ((!("\n\t" == "")) || (30 > 127)) {
                                        if false {
                                            switch (-(fi180)) {
                                            case -669:
                                                if true {
                                                    if false {
                                                        var v184 bool = ((!(true || (1.5 >= (0.1 - -2.25)))) && false)
                                                        _ = v184
                                                    } else {
                                                        var v185 float32 = ((-2.25 + 1e10) * (0.1 + -2.25))
                                                        _ = v185
                                                    }
                                                    if len(v26) >= 2 {
                                                        v26 = v26[1:len(v26)]
                                                    }
                                                } else {
                                                    v186 := GenS21{F0: 127, F1: 128, F2: 652}
                                                    var v187 float32 = 1.5
                                                    var v188 bool = false
                                                    _ = v186
                                                    _ = v187
                                                    _ = v188
                                                }
                                            case 0:
                                                v189 := GenS22{F0: 18446744073709551615, F1: 977, F2: 31}
                                                _ = v189
                                            case 32768:
                                                v27 = 2147483647
                                            }
                                        } else {
                                            var v190 uint8 = v4
                                            var v191 int16 = int16(v190)
                                            switch -2147483648 {
                                            case 4294967295:
                                                v2 = 957
                                            case 42:
                                                v192 := GenS23{F0: -128}
                                                _ = v192
                                            }
                                            _ = v190
                                            _ = v191
                                        }
                                        var v193 uint32 = 679
                                        if true {
                                            v27 = v27
                                            switch (-(v2)) {
                                            case 0:
                                                v4 = v4
                                            case -641:
                                                if (true && (!(v193 <= v193))) {
                                                    v193 += (v193 >> ((18446744073709551615) & 63))
                                                    fmt.Println(fi157)
                                                    var v194 int32 = (-(v27))
                                                    var v195 int16 = v28
                                                    var v196 bool = (uint64(v4) < 1)
                                                    _ = v194
                                                    _ = v195
                                                    _ = v196
                                                }
                                            case -857:
                                                fmt.Println(v4)
                                                switch "\n\t" {
                                                case "foo bar":
                                                    fmt.Println(v28)
                                                }
                                                v197 := GenS15{F0: (-(v28)), F1: 4294967295}
                                                _ = v197
                                            }
                                            v27 = v3
                                        }
                                        _ = v193
                                    } else {
                                        var v198 uint16 = 1
                                        _ = v198
                                    }
                                default:
                                    v27 = (-(v27))
                                }
                            default:
                                v199 := GenS24{F0: v4, F1: 1}
                                _ = v199
                            }
                            var v200 float32 = 0.0
                            var v201 uint32 = 4294967295
                            _ = fi180
                            _ = v200
                            _ = v201
                        }
                        var v202 int8 = -128
                        _ = v202
                    } else {
                        switch 79 {
                        case 1:
                            var v203 uint = 65535
                            _ = v203
                        }
                    }
                }
            }
        }
        _ = fi157
    }
    if true {
        v27 += (v27 >> ((928) & 63))
    }
    v4 += v4
    if true {
        v204 := GenS22{F0: 280, F1: 34, F2: 0}
        var v205 int8 = 0
        var v206 int64 = 0
        _ = v204
        _ = v205
        _ = v206
    } else {
        v26 = append(v26, v2)
        var v207 uint64 = 0
        _ = len(v1)
        _ = v207
    }
    v208 := GenS25{F0: v27}
    if ((((float32(1e10) - ((0.0 + 3.14159) + 0.1)) == -2.25) || true) && (0.0 != -2.25)) {
        if false {
            v4 -= (v4 >> ((18446744073709551615) & 63))
            if (420 > 42) {
                if true {
                    v1[v2] = v2
                }
            }
        } else {
            for fi209 := 0; fi209 < 18; fi209++ {
                v210 := GenS20{F0: 1}
                _ = fi209
                _ = v210
            }
            for fi211 := 0; fi211 < 4; fi211++ {
                v212 := GenS12{F0: 42, F1: 96}
                _ = len(v26)
                v4 += (v4 ^ (v4 << ((42) & 63)))
                _ = fi211
                _ = v212
            }
            for fi213 := 0; fi213 < 6; fi213++ {
                if ((v4 == (v4 ^ (v4 - genFn1(0, 506, genFn0((-(v27)), v28, 218))))) || false) {
                    v3 = v3
                    v28 = v28
                    for fi214 := 0; fi214 < 15; fi214++ {
                        v2 = fi213
                        v215 := GenS26{F0: (-(v28)), F1: -1, F2: 42}
                        fmt.Println(v28, v2)
                        _ = fi214
                        _ = v215
                    }
                }
                for fi216 := 0; fi216 < 16; fi216++ {
                    var v217 float32 = 3.14159
                    v27 -= v27
                    if (!(-108 > -127)) {
                        if true {
                            var v218 uint64 = 4294967295
                            _ = v218
                        } else {
                            switch ("\n\t" + "foo bar") {
                            case "":
                                var v219 float64 = 0.1
                                _ = v219
                            }
                        }
                        switch fi213 {
                        case 4294967295:
                            fmt.Println(v28)
                        }
                        if (!true) {
                            v27 += (v27 * v3)
                            v217 = v217
                            v1[(fi216 - (fi216 % ((fi216) | 1)))] = int(v4)
                        } else {
                            v220 := GenS10{F0: (v27 - (v27 * -310))}
                            _ = len(v26)
                            _ = v220
                        }
                        var v221 uint8 = v4
                        _ = v221
                    } else {
                        if false {
                            var v222 uint16 = 0
                            v223 := GenS27{F0: -1, F1: 127}
                            _ = v222
                            _ = v223
                        }
                        v224 := GenS14{F0: "foo bar"}
                        _ = v224
                    }
                    _ = fi216
                    _ = v217
                }
                delete(v1, (fi213 << ((296) & 63)))
                _ = fi213
            }
        }
        if false {
            for fi225 := 0; fi225 < 15; fi225++ {
                for fi226 := 0; fi226 < 4; fi226++ {
                    v28 += v28
                    _ = fi226
                }
                _ = fi225
            }
        } else {
            if false {
                v227 := GenS28{F0: "a"}
                if (v28 == v28) {
                    v4 -= 0
                    if true {
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        if (2147483647 != 2) {
                            var v228 int32 = (v3 * (v27 - v27))
                            _ = v228
                        } else {
                            v229 := GenS7{F0: v28}
                            _ = len(v26)
                            switch (^(v2)) {
                            case 127:
                                fmt.Println(v4)
                            default:
                                v26 = append(v26, 1)
                            }
                            _ = v229
                        }
                    }
                    v3 = v3
                }
                _ = v227
            } else {
                v26 = append(v26, 4294967295)
                switch ("foo bar" + (("a" + "") + ("foo bar" + ((("\n\t" + "") + "") + ("\n\t" + (("\n\t" + "") + ("a" + ""))))))) {
                case "foo bar":
                    v230 := GenS3{F0: (v3 % ((v3) | 1))}
                    _ = v230
                default:
                    var v231 int8 = -38
                    _ = v231
                }
                if (!(!false)) {
                    v232 := GenS22{F0: uint64(v4), F1: 0, F2: 18446744073709551615}
                    _ = v232
                }
            }
            v28 = (v28 / ((v28) | 1))
        }
        if (65535 == 32767) {
            var v233 bool = (true && false)
            v28 = (v28 >> ((4294967295) & 63))
            fmt.Println(v3, v28)
            _ = v233
        } else {
            fmt.Println(v3, v3, v27)
        }
    }
    v26 = append(v26, (^(v2)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v208
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
