package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int; F1 uint; F2 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 uint; F2 int }
type GenS5 struct { F0 bool; F1 bool; F2 uint }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 int8; F1 int }
type GenS8 struct { F0 uint16; F1 int }
type GenS9 struct { F0 string; F1 int32; F2 int8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS12 struct { F0 uint }
func genFn0(p0 int16, p1 int32, p2 string) int16 {
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 bool = (p1 > 627)
            p2 = p2
            _ = fi7
            _ = v8
        }
        p0 += p0
        _ = fi6
    }
    fmt.Println(p1, p0, p2)
    v9 := GenS1{F0: (p1 / ((p1) | 1)), F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (p0 - 478)
}
func genFn1(p0 uint16, p1 int16, p2 int32) int8 {
    var v10 uint = 129
    for fi11 := 0; fi11 < 12; fi11++ {
        p2 += (p2 - (p2 - (p2 % ((-2147483648) | 1))))
        switch fi11 {
        case 2147483647:
            fmt.Println(fi11, v10)
        case 32767:
            v12 := []int{fi11}
            _ = v12
        case 697:
            v13 := []int{-9223372036854775808, fi11}
            _ = v13
        }
        v10 -= v10
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 0
    var v3 uint32 = 1
    v4 := make(map[uint32]int)
    var v5 int = 19
    if (42 <= 730) {
        v2 -= v2
        v3 -= v3
        v3 = (v3 ^ (v3 >> ((857) & 63)))
    } else {
        v2 = 1
        v3 += v3
        fmt.Println(v2, v5, v3)
    }
    v2 = -468
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 string = ""
            v3 -= (^(v3))
            _ = fi15
            _ = v16
        }
        switch fi14 {
        case -2147483648:
            v17 := GenS1{F0: 8, F1: 0}
            _ = v17
        case 349:
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 string = "a"
                v20 := GenS2{F0: 42, F1: 479, F2: (v3 >> ((4294967295) & 63))}
                if ((v19 + v19) == v19) {
                    fmt.Println(v3, v19)
                    v3 -= (v3 ^ 390)
                    var v21 int16 = -625
                    switch "a" {
                    case "":
                        switch (^(fi14)) {
                        case 1:
                            if false {
                                v2 += v2
                                var v22 uint16 = 249
                                v23 := GenS1{F0: 2147483647, F1: 840}
                                _ = v22
                                _ = v23
                            } else {
                                v21 = (v21 & (v21 % ((v21) | 1)))
                                v5 = (fi18 >> ((698) & 63))
                                var v24 string = v19
                                _ = v24
                            }
                        case 127:
                            if false {
                                fmt.Println(fi18, v19)
                            }
                        case -526:
                            _ = len(v4)
                        }
                    case "foo bar":
                        switch (v19 + v19) {
                        case "\n\t":
                            v25 := GenS3{F0: -2147483648}
                            _ = v25
                        case "hello":
                            if (-2147483648 > -985) {
                                var v26 uint16 = 32767
                                _ = v26
                            }
                        }
                    }
                    v27 := GenS4{F0: 308, F1: 42, F2: int(v3)}
                    var v28 int16 = (v21 - 32767)
                    var v29 uint8 = 128
                    if (18446744073709551615 != 28) {
                        v29 -= (v29 * (^(v29)))
                    }
                    fmt.Println(v19, v28, v5)
                    _ = v21
                    _ = v27
                    _ = v28
                    _ = v29
                }
                _ = fi18
                _ = v19
                _ = v20
            }
        case -795:
            for fi30 := 0; fi30 < 8; fi30++ {
                if false {
                    if true {
                        v31 := GenS3{F0: 941}
                        switch "" {
                        case "a":
                            _ = len(v4)
                        case "\n\t":
                            fmt.Println(v2, v5, v3)
                        case "hello":
                            v3 += 2
                        default:
                            v32 := []int{(fi30 | (fi14 * int(v2))), 0, v5}
                            _ = v32
                        }
                        _ = v31
                    } else {
                        v33 := []int{}
                        var v34 uint64 = 654
                        v35 := GenS1{F0: int32(v3), F1: 0}
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                    v36 := GenS5{F0: true, F1: (0.1 == 3.14159), F2: 42}
                    fmt.Println(v3, v3)
                    v37 := v4[(^(v3))]
                    _ = v36
                    _ = v37
                }
                v38 := []int{v5, fi14, (fi30 % (((fi30 & fi14)) | 1)), fi30}
                v3 -= 2
                if false {
                    if true {
                        if false {
                            switch "foo bar" {
                            case "\n\t":
                                var v39 int8 = genFn1(255, 1000, 1)
                                _ = v39
                            case "a":
                                if (!true) {
                                    v40 := GenS6{F0: false, F1: 0}
                                    fmt.Println(v2, v5)
                                    if true {
                                        v41 := GenS7{F0: 100, F1: fi14}
                                        v3 -= v3
                                        _ = v41
                                    } else {
                                        v42 := GenS7{F0: -92, F1: (fi30 & -2147483648)}
                                        _ = v42
                                    }
                                    _ = v40
                                } else {
                                    var v43 uint8 = 229
                                    v5 = fi30
                                    _ = v43
                                }
                            case "foo bar":
                                switch fi14 {
                                case 128:
                                    v44 := v4[(v3 / ((v3) | 1))]
                                    _ = v44
                                }
                            }
                            var v45 int32 = 0
                            if false {
                                v2 += v2
                                fmt.Println(v45)
                                var v46 int8 = 45
                                _ = v46
                            } else {
                                v45 -= (^(v45))
                                switch "\n\t" {
                                case "a":
                                    v47 := GenS3{F0: (v45 | v45)}
                                    _ = v47
                                }
                                switch (fi30 - (v5 - -403)) {
                                case 4294967295:
                                    v48 := GenS6{F0: false, F1: 255}
                                    _ = v48
                                }
                            }
                            _ = v45
                        } else {
                            if true {
                                v49 := GenS4{F0: (^(v5)), F1: 42, F2: (fi30 & (-(v5)))}
                                fmt.Println(v3)
                                if (100 == genFn1(0, 32767, 1)) {
                                    if (false || false) {
                                        switch 9223372036854775807 {
                                        case 9223372036854775807:
                                            if true {
                                                var v50 int64 = v2
                                                v51 := GenS6{F0: (v2 == v2), F1: 4294967295}
                                                var v52 uint = 65535
                                                _ = v50
                                                _ = v51
                                                _ = v52
                                            } else {
                                                var v53 uint64 = 204
                                                if true {
                                                    var v54 uint64 = 42
                                                    var v55 int64 = v2
                                                    _ = v54
                                                    _ = v55
                                                }
                                                var v56 string = ""
                                                _ = v53
                                                _ = v56
                                            }
                                        case -1:
                                            v2 -= -1
                                        case -430:
                                            if true {
                                                _ = len(v4)
                                                v38 = append(v38, 604)
                                                var v57 int64 = (v2 ^ v2)
                                                _ = v57
                                            }
                                        default:
                                            if true {
                                                if len(v38) >= 2 {
                                                    v38 = v38[1:len(v38)]
                                                }
                                                if ((-18 == -128) || ((-128 <= 0) && true)) {
                                                    switch "a" {
                                                    case "a":
                                                        v3 = v3
                                                    default:
                                                        var v58 int8 = -128
                                                        v5 = fi30
                                                        _ = v58
                                                    }
                                                    if (0 != 257) {
                                                        _ = len(v38)
                                                    } else {
                                                        switch 802 {
                                                        case 310:
                                                            if false {
                                                                switch fi30 {
                                                                case 32767:
                                                                    if false {
                                                                        var v59 int16 = genFn0(genFn0(789, -571, "\n\t"), -143, "foo bar")
                                                                        if false {
                                                                            var v60 int = (v5 + int(v2))
                                                                            if (2147483647 > (fi30 >> ((42) & 63))) {
                                                                                fmt.Println(fi30, v5, v59)
                                                                                if true {
                                                                                    fmt.Println(fi30)
                                                                                }
                                                                                var v61 uint = 1
                                                                                _ = v61
                                                                            }
                                                                            fmt.Println(fi14, v2, v59)
                                                                            v62 := GenS2{F0: fi30, F1: 357, F2: (^(v3))}
                                                                            _ = len(v4)
                                                                            v63 := GenS5{F0: ((v5 % (((v60 | -113)) | 1)) != (fi14 / (((fi14 << ((0) & 63))) | 1))), F1: true, F2: 65535}
                                                                            _ = v60
                                                                            _ = v62
                                                                            _ = v63
                                                                        } else {
                                                                            v64 := GenS4{F0: (fi30 + fi14), F1: 255, F2: (^(fi14))}
                                                                            _ = v64
                                                                        }
                                                                        fmt.Println(v5, fi30)
                                                                        if (!(1.5 == 0.1)) {
                                                                            v59 = (v59 / ((-798) | 1))
                                                                        } else {
                                                                            var v65 bool = false
                                                                            _ = v65
                                                                        }
                                                                        _ = v59
                                                                    } else {
                                                                        v2 += 0
                                                                    }
                                                                case -1:
                                                                    fmt.Println(v5, v3, v2)
                                                                case 1:
                                                                    switch (v5 << ((0) & 63)) {
                                                                    case 4294967295:
                                                                        v38 = append(v38, 71)
                                                                    case 524:
                                                                        var v66 bool = true
                                                                        _ = v66
                                                                    }
                                                                    var v67 float32 = (0.0 + 1e10)
                                                                    _ = v67
                                                                }
                                                                if (9223372036854775807 != 0) {
                                                                    v3 = (v3 | v3)
                                                                    var v68 int16 = 134
                                                                    var v69 int32 = 1
                                                                    _ = v68
                                                                    _ = v69
                                                                } else {
                                                                    v70 := GenS3{F0: 42}
                                                                    v4[(^(v3))] = (fi14 | fi30)
                                                                    if true {
                                                                        var v71 uint32 = v3
                                                                        v2 = v2
                                                                        v3 = (v3 << ((1) & 63))
                                                                        _ = v71
                                                                    }
                                                                    _ = v70
                                                                }
                                                                _ = len(v1)
                                                            } else {
                                                                if true {
                                                                    switch (fi14 + v5) {
                                                                    case 127:
                                                                        var v72 int64 = (v2 << ((65535) & 63))
                                                                        _ = v72
                                                                    }
                                                                    v3 += v3
                                                                    v2 -= (v2 << ((255) & 63))
                                                                }
                                                            }
                                                        case 382:
                                                            v5 -= (v5 >> ((2) & 63))
                                                        case -1:
                                                            var v73 uint16 = 0
                                                            _ = v73
                                                        }
                                                    }
                                                } else {
                                                    var v74 int64 = (-(v2))
                                                    _ = v74
                                                }
                                                fmt.Println(v5)
                                                if (127 != -45) {
                                                    var v75 uint16 = 0
                                                    _ = v75
                                                } else {
                                                    if len(v38) >= 2 {
                                                        v38 = v38[1:len(v38)]
                                                    }
                                                    var v76 int = v5
                                                    _ = v76
                                                }
                                                switch (fi30 << ((2) & 63)) {
                                                case 42:
                                                    v77 := GenS8{F0: 32767, F1: -1}
                                                    _ = v77
                                                case 1:
                                                    var v78 uint8 = 128
                                                    _ = v78
                                                case 32767:
                                                    switch v5 {
                                                    case 127:
                                                        fmt.Println(v2, v2, v2)
                                                    case 32767:
                                                        v79 := v4[v3]
                                                        _ = v79
                                                    }
                                                }
                                            }
                                        }
                                        if false {
                                            switch v5 {
                                            case 2147483647:
                                                switch (v5 & (v5 / ((1) | 1))) {
                                                case 626:
                                                    switch (fi30 + (v5 >> ((186) & 63))) {
                                                    case -159:
                                                        fmt.Println(fi14, v2, fi14)
                                                    }
                                                }
                                            case 127:
                                                if true {
                                                    if (!((!(42 >= 32767)) || (15 >= -128))) {
                                                        if false {
                                                            var v80 int = int(v3)
                                                            _ = v80
                                                        } else {
                                                            v3 = (v3 ^ v3)
                                                            v81 := GenS8{F0: 255, F1: (fi14 + fi30)}
                                                            if false {
                                                                v2 = -93
                                                                var v82 int8 = -128
                                                                _ = v82
                                                            } else {
                                                                _ = len(v4)
                                                                v83 := GenS2{F0: 735, F1: 4294967295, F2: v3}
                                                                if (((genFn0(523, -1, ("\n\t" + "\n\t")) < 32767) || false) || true) {
                                                                    var v84 int32 = -1
                                                                    _ = v84
                                                                } else {
                                                                    v85 := v4[v3]
                                                                    v5 = (v85 << ((984) & 63))
                                                                    _ = v85
                                                                }
                                                                _ = v83
                                                            }
                                                            v86 := GenS1{F0: 42, F1: -1}
                                                            _ = v81
                                                            _ = v86
                                                        }
                                                    } else {
                                                        v87 := GenS1{F0: 93, F1: 549}
                                                        if (v5 <= fi14) {
                                                            var v88 uint = 673
                                                            v5 = 9223372036854775807
                                                            var v89 uint32 = (v3 % ((v3) | 1))
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                        _ = v87
                                                    }
                                                    switch v5 {
                                                    case 1:
                                                        fmt.Println(fi14, v3, v2)
                                                        if (-1 <= -1) {
                                                            var v90 uint = 619
                                                            var v91 uint8 = 42
                                                            _ = v90
                                                            _ = v91
                                                        }
                                                    case -920:
                                                        var v92 int16 = -743
                                                        _ = v92
                                                    case 128:
                                                        delete(v1, (v3 ^ (^(v3))))
                                                    }
                                                }
                                                _ = len(v38)
                                            default:
                                                v3 = v3
                                                delete(v4, v3)
                                            }
                                        } else {
                                            switch ("a" + ("a" + "foo bar")) {
                                            case "\n\t":
                                                fmt.Println(fi14)
                                                var v93 uint = 599
                                                if true {
                                                    v94 := v1[(v3 / ((v3) | 1))]
                                                    var v95 float32 = (0.0 * 0.1)
                                                    fmt.Println(v93, fi14, fi14)
                                                    _ = v94
                                                    _ = v95
                                                }
                                                _ = v93
                                            case "":
                                                var v96 int16 = 0
                                                _ = v96
                                            case "foo bar":
                                                if (0 >= 803) {
                                                    _ = len(v4)
                                                }
                                            default:
                                                var v97 int16 = 32767
                                                _ = v97
                                            }
                                        }
                                        var v98 float64 = float64(3.14159)
                                        _ = v98
                                    } else {
                                        v2 = (^(v2))
                                    }
                                    v3 = 2147483647
                                    _ = len(v1)
                                }
                                _ = v49
                            }
                            v99 := GenS9{F0: "\n\t", F1: 2147483647, F2: 12}
                            _ = v99
                        }
                        v3 = 42
                        var v100 bool = (38 > 363)
                        _ = v100
                    } else {
                        v2 += v2
                        if false {
                            var v101 int = (^(fi14))
                            var v102 int64 = (v2 + v2)
                            switch (("\n\t" + "\n\t") + "foo bar") {
                            case "foo bar":
                                if true {
                                    if false {
                                        fmt.Println(v5)
                                        var v103 uint32 = (v3 % ((v3) | 1))
                                        if (127 < 6) {
                                            var v104 uint64 = 72
                                            _ = v104
                                        }
                                        _ = v103
                                    }
                                    var v105 bool = (1e10 > (((0.0 * 0.1) / (1.5 - float64(-2.25))) - float64((0.1 + 1e10))))
                                    fmt.Println(v105, v5)
                                    if v105 {
                                        v38 = append(v38, (v5 - fi30))
                                        var v106 int64 = v2
                                        if v105 {
                                            delete(v1, v3)
                                            v107 := GenS6{F0: v105, F1: 18446744073709551615}
                                            var v108 float64 = 3.14159
                                            _ = v107
                                            _ = v108
                                        }
                                        _ = len(v38)
                                        v2 = (v2 >> ((18446744073709551615) & 63))
                                        _ = v106
                                    }
                                    fmt.Println(v102, fi30, fi14)
                                    _ = v105
                                }
                            case "a":
                                v102 = (v102 + (v2 >> ((278) & 63)))
                            case "hello":
                                v109 := GenS9{F0: ("a" + "foo bar"), F1: 19, F2: 100}
                                _ = v109
                            }
                            _ = v101
                            _ = v102
                        } else {
                            v110 := GenS8{F0: 478, F1: (v5 << ((2) & 63))}
                            v111 := GenS10{F0: 2}
                            _ = v110
                            _ = v111
                        }
                        v2 += 71
                    }
                    var v112 uint16 = 65535
                    _ = v112
                }
                v113 := GenS4{F0: fi30, F1: 2, F2: (fi14 - (v5 >> ((4294967295) & 63)))}
                _ = fi30
                _ = v113
                _ = v38
            }
        default:
            var v114 float32 = ((((-2.25 / 0.1) - (float32(1.5) - (-2.25 - 3.14159))) + 1e10) - 0.1)
            _ = v114
        }
        _ = fi14
    }
    if false {
        var v115 string = "foo bar"
        fmt.Println(v5)
        fmt.Println(v3)
        _ = v115
    }
    if true {
        if (v3 < (^(v3))) {
            v3 = v3
            var v116 uint16 = 2
            if (v5 >= v5) {
                var v117 uint8 = 249
                fmt.Println(v3)
                _ = v117
            } else {
                v118 := v4[2147483647]
                for fi119 := 0; fi119 < 6; fi119++ {
                    if false {
                        for fi120 := 0; fi120 < 14; fi120++ {
                            var v121 int8 = 0
                            v122 := GenS9{F0: ((("foo bar" + (("" + "foo bar") + "a")) + "\n\t") + "a"), F1: -1, F2: (v121 - (v121 + (v121 & v121)))}
                            v123 := []int{-533, (v5 ^ (v5 - -9223372036854775808)), (v5 - 0), v5}
                            _ = fi120
                            _ = v121
                            _ = v122
                            _ = v123
                        }
                        v124 := GenS11{F0: (false && (fi119 > v118)), F1: 0, F2: v3}
                        _ = v124
                    } else {
                        if true {
                            var v125 uint32 = v3
                            _ = v125
                        }
                    }
                    if true {
                        var v126 float64 = ((3.14159 / 3.14159) * -2.25)
                        v3 += v3
                        var v127 int16 = -1
                        _ = v126
                        _ = v127
                    } else {
                        v2 = (-(v2))
                        switch "" {
                        case "foo bar":
                            if (2 == 90) {
                                v4[2] = (fi119 ^ 128)
                            } else {
                                v128 := GenS9{F0: "hello", F1: 1, F2: 0}
                                fmt.Println(v2, fi119, v118)
                                v129 := GenS4{F0: (v5 / ((32767) | 1)), F1: uint(v116), F2: fi119}
                                _ = v128
                                _ = v129
                            }
                        case "":
                            fmt.Println(v5)
                        default:
                            switch (v118 << ((65535) & 63)) {
                            case 127:
                                if false {
                                    for fi130 := 0; fi130 < 15; fi130++ {
                                        v3 = v3
                                        var v131 bool = (true && ((-2.25 / 3.14159) != 3.14159))
                                        if (100 > 76) {
                                            v116 = (v116 % (((v116 | v116)) | 1))
                                        } else {
                                            fmt.Println(fi119, fi130)
                                            v4[v3] = 127
                                            v2 = v2
                                        }
                                        _ = fi130
                                        _ = v131
                                    }
                                    v116 = 32767
                                    fmt.Println(fi119, v5)
                                }
                            default:
                                v5 += fi119
                            }
                        }
                    }
                    _ = fi119
                }
                _ = v118
            }
            _ = v116
        } else {
            if ((v5 > (v5 << ((0) & 63))) || false) {
                v3 += 484
                v3 = v3
                v132 := GenS1{F0: 42, F1: -1}
                _ = v132
            }
            v2 += v2
            if (!false) {
                v133 := GenS5{F0: false, F1: false, F2: 65535}
                for fi134 := 0; fi134 < 15; fi134++ {
                    v3 = (v3 / ((187) | 1))
                    if false {
                        v135 := []int{2147483647}
                        _ = v135
                    }
                    _ = fi134
                }
                _ = v133
            }
        }
    }
    v136 := []int{v5}
    v137 := GenS12{F0: 0}
    if false {
        var v138 uint = 4294967295
        _ = v138
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v136:", len(v136), v136)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v136
    _ = v137
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
