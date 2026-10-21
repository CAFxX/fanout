package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint32; F1 int; F2 int64 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8; F1 int; F2 string }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int; F1 uint64; F2 string }
type GenS11 struct { F0 string; F1 int16; F2 int64 }
type GenS12 struct { F0 int64; F1 int8 }
type GenS13 struct { F0 uint16 }
func genFn0(p0 uint64, p1 int16, p2 int64) uint8 {
    switch "foo bar" {
    case "a":
        if false {
            p1 = (p1 | (p1 >> ((255) & 63)))
            var v7 int64 = (p2 - p2)
            _ = v7
        } else {
            if false {
                v8 := GenS1{F0: 1}
                _ = v8
            }
            if (("hello" < "foo bar") || (!(!true))) {
                var v9 bool = (!(float64(float32((3.14159 / (0.1 + float64(1e10))))) != float64(-2.25)))
                for fi10 := 0; fi10 < 5; fi10++ {
                    v11 := GenS1{F0: 255}
                    _ = fi10
                    _ = v11
                }
                var v12 bool = v9
                _ = v12
                _ = v9
            }
            p2 = (p2 ^ p2)
        }
    case "foo bar":
        v13 := GenS1{F0: 42}
        _ = v13
    }
    if (0 <= -1) {
        v14 := GenS1{F0: 127}
        _ = v14
    }
    if false {
        var v15 bool = ("" != ((("a" + ("foo bar" + (("hello" + "a") + "\n\t"))) + "") + ("a" + "a")))
        if v15 {
            v15 = (283 != 2)
        }
        v15 = v15
        _ = v15
    } else {
        switch 32767 {
        case 9223372036854775807:
            if ("" > "a") {
                v16 := []int{32768, -2147483648, 127, 127}
                switch 42 {
                case -895:
                    var v17 uint = 209
                    _ = v17
                case -2147483648:
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                case -9223372036854775808:
                    p1 = (p1 / ((-1) | 1))
                default:
                    var v18 int8 = 127
                    _ = v18
                }
                if (429 <= 622) {
                    var v19 uint32 = 204
                    _ = v19
                } else {
                    var v20 string = ((("" + "foo bar") + "hello") + ("hello" + "a"))
                    p0 = 862
                    v21 := GenS2{F0: 531, F1: 4294967295, F2: 728}
                    _ = v20
                    _ = v21
                }
                _ = v16
            } else {
                fmt.Println(p2, p1, p2)
                if (false || (231 > 65535)) {
                    var v22 uint8 = 255
                    var v23 int8 = 127
                    p1 -= int16(p2)
                    _ = v22
                    _ = v23
                }
                switch 2147483647 {
                case 4294967295:
                    var v24 uint8 = 2
                    _ = v24
                }
            }
        default:
            if false {
                v25 := []int{32768, -1}
                _ = v25
            } else {
                for fi26 := 0; fi26 < 13; fi26++ {
                    var v27 uint = 1
                    p1 -= (p1 % ((983) | 1))
                    _ = fi26
                    _ = v27
                }
            }
        }
    }
    if (127 == 1) {
        switch "hello" {
        case "hello":
            fmt.Println(p1, p0)
        }
    } else {
        for fi28 := 0; fi28 < 10; fi28++ {
            v29 := []int{fi28, (fi28 << ((185) & 63))}
            var v30 int = fi28
            if ((!(!false)) || false) {
                v31 := GenS1{F0: 37}
                _ = v31
            } else {
                var v32 int32 = -1
                _ = v32
            }
            _ = fi28
            _ = v29
            _ = v30
        }
        fmt.Println(p2, p0)
        for fi33 := 0; fi33 < 19; fi33++ {
            var v34 uint64 = (p0 - 0)
            _ = fi33
            _ = v34
        }
    }
    var v35 int8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v35
    return 0
}

func main() {
    var v1 int32 = 0
    var v2 int64 = 0
    v3 := []int{127, 446, 9223372036854775807}
    v4 := []int{}
    v5 := make(map[int]int)
    var v6 int = -221
    if false {
        switch (("" + "\n\t") + "a") {
        case "\n\t":
            var v36 float32 = (float32(0.0) * (1.5 - 3.14159))
            _ = v36
        default:
            if (!((!true) && true)) {
                for fi37 := 0; fi37 < 5; fi37++ {
                    v1 += (v1 + 42)
                    _ = fi37
                }
                fmt.Println(v1, v6)
            }
        }
    } else {
        if false {
            v38 := GenS2{F0: 18446744073709551615, F1: 556, F2: 232}
            v39 := GenS3{F0: 2, F1: v2}
            if (42 <= 0) {
                var v40 float32 = (((1.5 * (1.5 - 0.1)) + (3.14159 - ((3.14159 * float32((0.1 - (((3.14159 / 1e10) / 0.1) + 0.0)))) / float32((((3.14159 * (((((0.1 * 0.0) * ((0.1 * (((0.0 - 0.1) / -2.25) - (-2.25 * (((0.1 * ((-2.25 / -2.25) - -2.25)) / 1e10) - 3.14159)))) / 3.14159)) * 1e10) - 0.0) - (1.5 + -2.25))) / (3.14159 + 1.5)) / 1.5))))) / 3.14159)
                if false {
                    delete(v5, v6)
                    if (!false) {
                        for fi41 := 0; fi41 < 5; fi41++ {
                            for fi42 := 0; fi42 < 17; fi42++ {
                                var v43 int64 = v2
                                if true {
                                    v5[(v6 >> ((255) & 63))] = v6
                                }
                                _ = fi42
                                _ = v43
                            }
                            if (1 >= 1000) {
                                fmt.Println(v6)
                                fmt.Println(v40, fi41)
                            }
                            _ = fi41
                        }
                    } else {
                        _ = len(v3)
                        fmt.Println(v2, v40, v6)
                        if false {
                            v1 += (v1 / ((1) | 1))
                            for fi44 := 0; fi44 < 18; fi44++ {
                                for fi45 := 0; fi45 < 19; fi45++ {
                                    var v46 bool = (!false)
                                    v6 += -742
                                    _ = fi45
                                    _ = v46
                                }
                                v47 := GenS3{F0: 0, F1: (^(v2))}
                                v48 := GenS4{F0: 2, F1: -1, F2: (v2 >> ((18446744073709551615) & 63))}
                                _ = fi44
                                _ = v47
                                _ = v48
                            }
                            v6 = (v6 ^ (^(v6)))
                        } else {
                            if false {
                                if len(v4) > 0 {
                                    v4[0] = v6
                                }
                                fmt.Println(v2, v6)
                            } else {
                                v49 := GenS4{F0: 2, F1: 0, F2: v2}
                                v5[-1] = 9223372036854775807
                                _ = v49
                            }
                        }
                    }
                } else {
                    v50 := GenS2{F0: 42, F1: 769, F2: 4294967295}
                    _ = v50
                }
                _ = v40
            } else {
                switch v6 {
                case -656:
                    for fi51 := 0; fi51 < 10; fi51++ {
                        v52 := GenS3{F0: 0, F1: v2}
                        _ = fi51
                        _ = v52
                    }
                case 782:
                    if true {
                        switch (v6 - v6) {
                        case 108:
                            var v53 uint8 = 2
                            _ = v53
                        default:
                            var v54 float64 = (0.1 * 3.14159)
                            _ = v54
                        }
                        var v55 bool = (!true)
                        for fi56 := 0; fi56 < 10; fi56++ {
                            v55 = true
                            if (!true) {
                                switch (v6 - (fi56 % (((v6 & 32768)) | 1))) {
                                case 127:
                                    switch "" {
                                    case "foo bar":
                                        if v55 {
                                            v2 += 1
                                            v3 = append(v3, (v6 - (fi56 * fi56)))
                                            fmt.Println(fi56, fi56, fi56)
                                        } else {
                                            v57 := GenS5{F0: fi56, F1: 4294967295}
                                            v6 = fi56
                                            var v58 uint32 = 459
                                            _ = v57
                                            _ = v58
                                        }
                                    default:
                                        v6 -= (v6 / ((fi56) | 1))
                                    }
                                case 42:
                                    if v55 {
                                        var v59 string = ("foo bar" + "\n\t")
                                        var v60 string = (v59 + v59)
                                        v1 -= (v1 + (v1 & v1))
                                        _ = v59
                                        _ = v60
                                    } else {
                                        var v61 bool = (1 > 175)
                                        _ = v61
                                    }
                                case 32767:
                                    if (192 != 2) {
                                        if true {
                                            v62 := GenS5{F0: (v6 / (((fi56 % (((-(fi56))) | 1))) | 1)), F1: 944}
                                            var v63 bool = v55
                                            var v64 uint32 = 860
                                            _ = v62
                                            _ = v63
                                            _ = v64
                                        } else {
                                            if ((!("\n\t" > "")) || ((false && v55) || v55)) {
                                                fmt.Println(v55, v2, v55)
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v1 += (v1 ^ v1)
                                        }
                                        v65 := GenS5{F0: v6, F1: 65535}
                                        switch "" {
                                        case "\n\t":
                                            v2 = (v2 << ((2) & 63))
                                        case "a":
                                            v2 = v2
                                        case "":
                                            if len(v3) > 0 {
                                                v3[0] = 748
                                            }
                                        default:
                                            for fi66 := 0; fi66 < 9; fi66++ {
                                                if false {
                                                    v67 := GenS6{F0: 52}
                                                    v68 := GenS2{F0: 255, F1: 1, F2: 2147483648}
                                                    v69 := GenS6{F0: 1}
                                                    _ = v67
                                                    _ = v68
                                                    _ = v69
                                                }
                                                v55 = ((0.1 < (1.5 + (float32(1.5) * 1e10))) || true)
                                                v2 += 9223372036854775807
                                                fmt.Println(fi66, v2)
                                                v6 = 128
                                                v6 = (fi66 * (-(v6)))
                                                _ = fi66
                                            }
                                        }
                                        _ = v65
                                    } else {
                                        v6 -= (v6 + (^(v6)))
                                        switch fi56 {
                                        case 174:
                                            fmt.Println(v2, fi56, v2)
                                        case -1:
                                            v6 = 509
                                        case 2147483647:
                                            var v70 uint = 4294967295
                                            _ = v70
                                        }
                                    }
                                default:
                                    v71 := GenS5{F0: (v6 << ((2) & 63)), F1: 18446744073709551615}
                                    _ = v71
                                }
                                var v72 int16 = -669
                                _ = v72
                            } else {
                                switch 42 {
                                case -708:
                                    _ = len(v3)
                                case 252:
                                    v73 := GenS7{F0: 127, F1: (-(fi56)), F2: "foo bar"}
                                    _ = v73
                                }
                                v74 := GenS5{F0: 32767, F1: 404}
                                v75 := GenS8{F0: 1}
                                _ = v74
                                _ = v75
                            }
                            switch 32767 {
                            case 433:
                                v76 := GenS3{F0: 1, F1: (v2 & -608)}
                                _ = v76
                            case 2147483647:
                                for fi77 := 0; fi77 < 19; fi77++ {
                                    v1 += 42
                                    v6 += (-(fi56))
                                    switch fi56 {
                                    case 4294967295:
                                        var v78 uint64 = 1
                                        _ = v78
                                    case 42:
                                        v55 = v55
                                    default:
                                        v79 := GenS4{F0: 847, F1: fi56, F2: (v2 + (v2 << ((618) & 63)))}
                                        _ = v79
                                    }
                                    _ = fi77
                                }
                            case 127:
                                var v80 int16 = -683
                                _ = v80
                            default:
                                var v81 int = (fi56 ^ (fi56 >> ((4294967295) & 63)))
                                _ = v81
                            }
                            _ = fi56
                        }
                        _ = v55
                    } else {
                        if (true || (false || ((64 == 42) || ((0 < 1) && true)))) {
                            v6 -= int(v1)
                        } else {
                            var v82 float32 = ((3.14159 / (((-2.25 - 0.0) - float32(0.1)) * (0.1 - 1.5))) - 1.5)
                            _ = v82
                        }
                        switch "foo bar" {
                        case "foo bar":
                            var v83 uint8 = 255
                            _ = v83
                        case "":
                            v6 = (v6 * v6)
                        }
                        for fi84 := 0; fi84 < 4; fi84++ {
                            switch "\n\t" {
                            case "foo bar":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                var v85 uint16 = 916
                                _ = v85
                            }
                            for fi86 := 0; fi86 < 3; fi86++ {
                                switch ("a" + "") {
                                case "foo bar":
                                    v1 -= (v1 * v1)
                                case "":
                                    if false {
                                        fmt.Println(v2, fi86, v6)
                                        var v87 uint = 735
                                        var v88 int8 = 0
                                        _ = v87
                                        _ = v88
                                    }
                                }
                                _ = fi86
                            }
                            _ = fi84
                        }
                    }
                }
            }
            _ = v38
            _ = v39
        } else {
            fmt.Println(v2, v2, v1)
            if false {
                switch (v6 ^ 42) {
                case 1:
                    var v89 string = "\n\t"
                    _ = v89
                case 0:
                    var v90 int16 = 982
                    _ = v90
                default:
                    if (v6 > v6) {
                        if ("foo bar" > "a") {
                            if (-32768 <= 0) {
                                for fi91 := 0; fi91 < 2; fi91++ {
                                    var v92 uint32 = 290
                                    for fi93 := 0; fi93 < 2; fi93++ {
                                        v92 = (^(v92))
                                        switch (v6 % (((^(v6))) | 1)) {
                                        case -9223372036854775808:
                                            v94 := GenS3{F0: 0, F1: v2}
                                            _ = v94
                                        case 703:
                                            fmt.Println(fi91, v2, fi93)
                                        case 9223372036854775807:
                                            v92 -= (v92 % ((v92) | 1))
                                        default:
                                            switch "a" {
                                            case "hello":
                                                if (v1 < (v1 >> ((4294967295) & 63))) {
                                                    var v95 int8 = 1
                                                    _ = len(v3)
                                                    v3 = append(v3, (fi93 ^ (fi91 | (fi91 / ((579) | 1)))))
                                                    _ = v95
                                                } else {
                                                    var v96 int8 = 1
                                                    var v97 uint8 = genFn0(0, 1000, (v2 >> ((640) & 63)))
                                                    _ = v96
                                                    _ = v97
                                                }
                                            case "foo bar":
                                                fmt.Println(v92, v2)
                                            default:
                                                v2 = (v2 / ((42) | 1))
                                            }
                                        }
                                        var v98 uint32 = (v92 / ((762) | 1))
                                        _ = fi93
                                        _ = v98
                                    }
                                    _ = fi91
                                    _ = v92
                                }
                            }
                        }
                        for fi99 := 0; fi99 < 2; fi99++ {
                            switch fi99 {
                            case 152:
                                fmt.Println(v1, v6)
                            case -1:
                                v100 := GenS5{F0: -2147483648, F1: 65535}
                                _ = v100
                            case -531:
                                v101 := GenS1{F0: 1}
                                _ = v101
                            default:
                                var v102 string = ("foo bar" + "foo bar")
                                _ = v102
                            }
                            v1 = -120
                            _ = fi99
                        }
                        switch (v6 >> ((32) & 63)) {
                        case 32767:
                            if true {
                                v2 += v2
                                v103 := GenS8{F0: 2}
                                v1 = v1
                                _ = v103
                            }
                        }
                    } else {
                        for fi104 := 0; fi104 < 9; fi104++ {
                            for fi105 := 0; fi105 < 19; fi105++ {
                                v1 = v1
                                if false {
                                    v2 = v2
                                    v3 = append(v3, fi104)
                                    var v106 int16 = -32768
                                    _ = v106
                                } else {
                                    if true {
                                        var v107 bool = (0.1 > (1e10 - (-2.25 / (3.14159 * 3.14159))))
                                        v108 := GenS7{F0: 42, F1: (fi105 + (fi104 % (((fi104 - (v6 << ((18446744073709551615) & 63)))) | 1))), F2: (("" + ("\n\t" + "a")) + ((("a" + ("\n\t" + "foo bar")) + ("foo bar" + "hello")) + "a"))}
                                        var v109 uint = 42
                                        _ = v107
                                        _ = v108
                                        _ = v109
                                    }
                                    switch ("foo bar" + "foo bar") {
                                    case "":
                                        switch (fi105 >> ((255) & 63)) {
                                        case 2147483647:
                                            var v110 float64 = ((-2.25 + -2.25) - (float64(-2.25) - 0.0))
                                            _ = v110
                                        }
                                    }
                                }
                                var v111 int16 = 42
                                _ = fi105
                                _ = v111
                            }
                            _ = fi104
                        }
                    }
                }
                for fi112 := 0; fi112 < 10; fi112++ {
                    v1 = (v1 | (v1 % ((v1) | 1)))
                    for fi113 := 0; fi113 < 8; fi113++ {
                        v2 = -9223372036854775808
                        var v114 float32 = 1e10
                        var v115 int8 = 0
                        _ = fi113
                        _ = v114
                        _ = v115
                    }
                    var v116 uint64 = 1
                    _ = fi112
                    _ = v116
                }
                switch "foo bar" {
                case "hello":
                    var v117 uint8 = genFn0(403, 1000, 9223372036854775807)
                    _ = v117
                case "a":
                    switch v6 {
                    case -471:
                        if true {
                            var v118 bool = true
                            var v119 float64 = ((1.5 + 0.1) - (((3.14159 - 1.5) / 3.14159) / 3.14159))
                            if (false || true) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v120 int16 = 1000
                                _ = v120
                            } else {
                                if v118 {
                                    fmt.Println(v119, v2)
                                    for fi121 := 0; fi121 < 1; fi121++ {
                                        fmt.Println(v118, v118)
                                        v122 := GenS9{F0: -1}
                                        _ = fi121
                                        _ = v122
                                    }
                                    var v123 uint32 = 4294967295
                                    _ = v123
                                } else {
                                    v6 = (^(v6))
                                    v119 += 1.5
                                    switch 1 {
                                    case -1:
                                        var v124 uint16 = 515
                                        _ = v124
                                    case -58:
                                        v125 := GenS1{F0: genFn0(42, int16(v6), (v2 - 9223372036854775807))}
                                        _ = v125
                                    }
                                }
                                fmt.Println(v1)
                                if (42 >= -59) {
                                    v4 = append(v4, (v6 >> ((0) & 63)))
                                    if v118 {
                                        if v118 {
                                            for fi126 := 0; fi126 < 11; fi126++ {
                                                if (100 < 83) {
                                                    v119 = 1.5
                                                }
                                                v127 := GenS10{F0: (v6 * (v6 & 9223372036854775807)), F1: 4294967295, F2: "foo bar"}
                                                _ = fi126
                                                _ = v127
                                            }
                                            var v128 uint64 = 623
                                            _ = v128
                                        } else {
                                            var v129 uint16 = 0
                                            _ = v129
                                        }
                                        fmt.Println(v118)
                                    } else {
                                        v118 = false
                                    }
                                }
                            }
                            _ = v118
                            _ = v119
                        }
                    default:
                        v2 = 42
                    }
                case "foo bar":
                    if ((v1 ^ v1) <= (v1 << ((0) & 63))) {
                        _ = len(v5)
                        var v130 string = "\n\t"
                        for fi131 := 0; fi131 < 15; fi131++ {
                            for fi132 := 0; fi132 < 2; fi132++ {
                                v133 := GenS6{F0: 0}
                                switch (v6 - (fi131 + 998)) {
                                case -2147483648:
                                    v134 := GenS11{F0: "hello", F1: 1000, F2: (v2 >> ((255) & 63))}
                                    _ = v134
                                case 2147483647:
                                    var v135 int32 = (v1 >> ((4294967295) & 63))
                                    _ = v135
                                case 127:
                                    if ((v130 + v130) < (v130 + ((v130 + v130) + (v130 + v130)))) {
                                        v130 = v130
                                        var v136 uint64 = 10
                                        _ = v136
                                    } else {
                                        var v137 uint = 1
                                        _ = v137
                                    }
                                default:
                                    v138 := GenS5{F0: v6, F1: 42}
                                    _ = v138
                                }
                                var v139 int32 = v1
                                _ = fi132
                                _ = v133
                                _ = v139
                            }
                            _ = fi131
                        }
                        _ = v130
                    } else {
                        _ = len(v3)
                        var v140 int8 = 100
                        _ = v140
                    }
                default:
                    var v141 int16 = -32768
                    _ = v141
                }
            }
        }
        if false {
            v142 := GenS11{F0: (("a" + ("\n\t" + ((("\n\t" + "") + "foo bar") + "a"))) + "a"), F1: 0, F2: (v2 / ((v2) | 1))}
            var v143 int = (v6 ^ -795)
            _ = v142
            _ = v143
        } else {
            fmt.Println(v2)
        }
        var v144 int32 = v1
        _ = v144
    }
    v2 -= int64(v1)
    var v145 uint8 = 42
    if false {
        v6 += 494
        v145 -= 2
    }
    if (!(42 <= 65535)) {
        var v146 int32 = (v1 * -2147483648)
        v2 -= (v2 ^ 42)
        _ = v146
    }
    switch 676 {
    case 909:
        v2 = (-(v2))
    default:
        if (true && (false && (!(!(!true))))) {
            for fi147 := 0; fi147 < 10; fi147++ {
                var v148 float32 = 0.0
                v2 += v2
                _ = fi147
                _ = v148
            }
            if (((65535 < 0) || (0 != 207)) && (true && (true || false))) {
                fmt.Println(v1, v145)
                v149 := GenS5{F0: (v6 & (v6 >> ((2) & 63))), F1: 237}
                for fi150 := 0; fi150 < 19; fi150++ {
                    var v151 uint32 = 2147483648
                    fmt.Println(v2, v6)
                    _ = fi150
                    _ = v151
                }
                _ = v149
            }
        }
    }
    var v152 string = ("foo bar" + "foo bar")
    v145 -= genFn0(42, 42, 412)
    fmt.Println(v1)
    var v153 uint64 = 18446744073709551615
    v154 := GenS12{F0: v2, F1: -76}
    var v155 uint64 = v153
    fmt.Println(v153, v145)
    switch (v152 + ((v152 + (("\n\t" + "a") + v152)) + ("" + v152))) {
    case "":
        fmt.Println(v2, v152)
    }
    for fi156 := 0; fi156 < 20; fi156++ {
        var v157 int64 = (v2 ^ v2)
        v3 = append(v3, 42)
        _ = fi156
        _ = v157
    }
    v158 := GenS13{F0: 65535}
    var v159 int16 = 1000
    switch (-(v6)) {
    case 9223372036854775807:
        switch "\n\t" {
        case "\n\t":
            v153 = (^(v155))
        case "hello":
            var v160 uint16 = 455
            _ = v160
        case "foo bar":
            if len(v4) > 0 {
                v4[0] = (^(v6))
            }
        default:
            _ = len(v5)
        }
    case 32768:
        for fi161 := 0; fi161 < 3; fi161++ {
            var v162 int16 = 32767
            for fi163 := 0; fi163 < 2; fi163++ {
                v152 = v152
                v6 = (-(v6))
                _ = fi163
            }
            _ = fi161
            _ = v162
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v145:", v145)
    fmt.Println("v152:", v152)
    fmt.Println("v153:", v153)
    fmt.Println("v155:", v155)
    fmt.Println("v159:", v159)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v145
    _ = v152
    _ = v153
    _ = v154
    _ = v155
    _ = v158
    _ = v159
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
