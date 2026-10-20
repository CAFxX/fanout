package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 int64; F1 int16; F2 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 string; F1 int16 }
type GenS10 struct { F0 int64; F1 uint16 }
type GenS11 struct { F0 uint; F1 bool }
type GenS12 struct { F0 uint; F1 uint64; F2 string }
type GenS13 struct { F0 int8; F1 int64 }
type GenS14 struct { F0 int; F1 int64; F2 uint }
type GenS15 struct { F0 bool; F1 string; F2 int8 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 int32; F1 int; F2 uint8 }
type GenS18 struct { F0 int; F1 int64 }
type GenS19 struct { F0 int8; F1 uint64 }
type GenS20 struct { F0 bool; F1 bool; F2 int16 }
type GenS21 struct { F0 int16; F1 int32; F2 int8 }
type GenS22 struct { F0 int64; F1 uint64 }
type GenS23 struct { F0 bool }
type GenS24 struct { F0 int16 }
type GenS25 struct { F0 int32; F1 int32; F2 uint64 }
type GenS26 struct { F0 uint }
type GenS27 struct { F0 int8; F1 bool }
type GenS28 struct { F0 int16; F1 int16 }
type GenS29 struct { F0 bool; F1 uint32 }
type GenS30 struct { F0 string; F1 int64; F2 uint16 }
func genFn0() uint16 {
    var v7 float32 = 1e10
    switch -1 {
    case 9223372036854775807:
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 string = "a"
            switch fi8 {
            case 32767:
                if (!true) {
                    v7 += 1.5
                }
            }
            _ = fi8
            _ = v9
        }
    default:
        var v10 uint16 = 65535
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        if ((42 <= 255) || (1 < 42)) {
            var v12 string = "hello"
            _ = v12
        } else {
            v13 := []int{(fi11 & fi11), (fi11 & fi11), (fi11 / (((-(fi11))) | 1))}
            switch fi11 {
            case -910:
                var v14 string = "foo bar"
                _ = v14
            }
            _ = v13
        }
        if (fi11 == 32767) {
            for fi15 := 0; fi15 < 16; fi15++ {
                fmt.Println(fi11, fi11)
                _ = fi15
            }
            var v16 uint64 = 42
            _ = v16
        } else {
            v17 := GenS1{F0: "", F1: 127}
            var v18 int = fi11
            v19 := GenS2{F0: false}
            _ = v17
            _ = v18
            _ = v19
        }
        var v20 uint64 = 2
        _ = fi11
        _ = v20
    }
    v21 := []int{2147483647, 4294967295, -9223372036854775808, 128}
    for fi22 := 0; fi22 < 18; fi22++ {
        v7 = (((v7 * (((((((v7 + v7) / v7) * v7) - (v7 / v7)) + v7) - (float32((((0.0 / ((0.0 + 0.1) - 1.5)) + -2.25) + (((-2.25 * 1.5) + 1.5) * float64(0.0)))) - ((float32(-2.25) / (0.0 + -2.25)) + v7))) - float32(1.5))) + (-2.25 - (v7 * v7))) * 3.14159)
        if false {
            v23 := GenS2{F0: false}
            fmt.Println(v7, v7)
            if true {
                fmt.Println(fi22)
                v7 += 0.0
                if (-474 < 0) {
                    v7 -= ((float32(0.0) / v7) - v7)
                    var v24 float32 = (0.0 * -2.25)
                    var v25 uint = 255
                    _ = v24
                    _ = v25
                } else {
                    switch "hello" {
                    case "":
                        if (0 == 14) {
                            var v26 uint64 = 18446744073709551615
                            v26 += (v26 & (^(v26)))
                            for fi27 := 0; fi27 < 3; fi27++ {
                                fmt.Println(v7)
                                _ = fi27
                            }
                            _ = v26
                        } else {
                            fmt.Println(v7)
                        }
                    case "hello":
                        var v28 float64 = 1e10
                        _ = v28
                    case "a":
                        if false {
                            if (0 != 2) {
                                for fi29 := 0; fi29 < 8; fi29++ {
                                    if (2 >= 42) {
                                        var v30 float32 = (v7 + v7)
                                        if false {
                                            var v31 int64 = -9223372036854775808
                                            _ = v31
                                        }
                                        _ = v30
                                    }
                                    if ((!false) && (!(((!(!false)) || true) || (1 > -1)))) {
                                        if (-1 < 0) {
                                            var v32 bool = true
                                            if (759 > -187) {
                                                var v33 uint = 4294967295
                                                v21 = append(v21, fi29)
                                                _ = v33
                                            } else {
                                                fmt.Println(fi22)
                                                switch "\n\t" {
                                                case "hello":
                                                    if false {
                                                        switch -2147483648 {
                                                        case 4294967295:
                                                            v32 = v32
                                                        case 128:
                                                            if len(v21) > 0 {
                                                                v21[0] = fi22
                                                            }
                                                        }
                                                    }
                                                default:
                                                    v7 -= ((1.5 + v7) / ((v7 - 1.5) - (float32((-2.25 - 3.14159)) - float32(0.1))))
                                                }
                                            }
                                            _ = v32
                                        }
                                        var v34 uint32 = 593
                                        var v35 float32 = (float32(1.5) + (3.14159 + v7))
                                        _ = v34
                                        _ = v35
                                    }
                                    var v36 float64 = 3.14159
                                    _ = fi29
                                    _ = v36
                                }
                                v37 := GenS2{F0: false}
                                fmt.Println(v7, fi22)
                                _ = v37
                            } else {
                                var v38 float32 = v7
                                _ = v38
                            }
                            if ((1e10 * -2.25) > -2.25) {
                                v39 := GenS3{F0: 13}
                                _ = v39
                            }
                        } else {
                            _ = len(v21)
                            var v40 int64 = 739
                            if (-1 > 42) {
                                if (2 <= 0) {
                                    var v41 uint32 = 1
                                    v7 += (v7 - (v7 + v7))
                                    _ = v41
                                }
                            } else {
                                var v42 uint8 = 127
                                var v43 float32 = 1.5
                                var v44 float64 = 0.1
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                            _ = v40
                        }
                    }
                }
            } else {
                for fi45 := 0; fi45 < 14; fi45++ {
                    switch "" {
                    case "\n\t":
                        if (false || true) {
                            v46 := GenS4{F0: (-(fi45)), F1: 0}
                            _ = v46
                        } else {
                            if ((fi22 - fi45) >= (fi22 << ((1) & 63))) {
                                switch 492 {
                                case 42:
                                    if false {
                                        v7 = v7
                                    } else {
                                        v47 := GenS5{F0: 1, F1: -559, F2: fi45}
                                        _ = len(v21)
                                        var v48 string = ("\n\t" + "")
                                        _ = v47
                                        _ = v48
                                    }
                                case -26:
                                    var v49 uint8 = 2
                                    _ = v49
                                case 32768:
                                    var v50 uint64 = 288
                                    _ = v50
                                default:
                                    if true {
                                        v7 = v7
                                        var v51 int32 = -1
                                        _ = v51
                                    }
                                }
                            }
                        }
                    case "a":
                        v52 := GenS3{F0: 595}
                        _ = v52
                    }
                    var v53 int32 = 0
                    if (false && false) {
                        var v54 uint64 = 274
                        v54 += v54
                        _ = v54
                    } else {
                        var v55 uint64 = 70
                        v55 -= v55
                        if (436 > 158) {
                            if len(v21) > 0 {
                                v21[0] = 4294967295
                            }
                            v53 = 42
                            if (v55 >= (^(v55))) {
                                fmt.Println(fi45)
                                if false {
                                    v56 := GenS2{F0: true}
                                    var v57 int64 = 1000
                                    fmt.Println(v53)
                                    _ = v56
                                    _ = v57
                                }
                                if (65535 != 450) {
                                    if (!((fi22 != (fi45 + (fi22 | (fi22 % ((int(v55)) | 1))))) && false)) {
                                        if false {
                                            _ = len(v21)
                                        }
                                        switch fi22 {
                                        case 0:
                                            fmt.Println(v53, v7)
                                        case 2147483647:
                                            fmt.Println(v7, v55)
                                        }
                                        v53 -= (v53 << ((424) & 63))
                                        if len(v21) > 0 {
                                            v21[0] = (fi45 * fi22)
                                        }
                                        v7 = (float32(1.5) + 0.0)
                                    } else {
                                        v53 = -2147483648
                                    }
                                    v7 -= (v7 + v7)
                                }
                            } else {
                                switch (("hello" + (("" + "") + ("hello" + ("\n\t" + (("\n\t" + "\n\t") + "foo bar"))))) + ("\n\t" + (("hello" + "") + ("" + "foo bar")))) {
                                case "hello":
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                case "\n\t":
                                    var v58 bool = (!false)
                                    _ = v58
                                case "":
                                    v21 = append(v21, (fi22 + -9223372036854775808))
                                default:
                                    v55 -= 1
                                }
                                _ = len(v21)
                            }
                        } else {
                            if true {
                                if len(v21) > 0 {
                                    v21[0] = fi45
                                }
                            } else {
                                if len(v21) > 0 {
                                    v21[0] = fi45
                                }
                                v21 = append(v21, 4294967295)
                                if (795 > 255) {
                                    var v59 int8 = 42
                                    v7 += ((v7 / -2.25) * float32((0.0 + 0.0)))
                                    fmt.Println(v7, v53)
                                    _ = v59
                                } else {
                                    v7 = v7
                                    fmt.Println(fi45, fi45)
                                }
                                v60 := GenS6{F0: (fi22 % ((-1) | 1))}
                                v55 = (^(v55))
                                _ = v60
                            }
                        }
                        _ = v55
                    }
                    var v61 uint8 = 0
                    v62 := GenS1{F0: ("\n\t" + "a"), F1: (v61 * v61)}
                    v63 := GenS7{F0: 574, F1: true}
                    if false {
                        v64 := GenS8{F0: 1}
                        var v65 uint = 2
                        switch "a" {
                        case "hello":
                            v53 += v53
                            var v66 uint8 = 127
                            switch (fi22 << ((18446744073709551615) & 63)) {
                            case 128:
                                if len(v21) > 0 {
                                    v21[0] = int(v61)
                                }
                            }
                            switch "hello" {
                            case "a":
                                if false {
                                    if (false && true) {
                                        var v67 bool = (v65 != (v65 & (v65 % ((1) | 1))))
                                        v7 += v7
                                        var v68 int8 = 127
                                        _ = len(v21)
                                        _ = v67
                                        _ = v68
                                    }
                                    if false {
                                        v7 = v7
                                        v65 = (v65 - 2)
                                        switch ("a" + "\n\t") {
                                        case "\n\t":
                                            v7 = float32(float64((v7 * float32(0.0))))
                                        case "":
                                            switch fi22 {
                                            case 2147483647:
                                                v69 := GenS9{F0: "", F1: -32768}
                                                _ = v69
                                            case 1:
                                                v70 := GenS10{F0: 514, F1: 2}
                                                _ = v70
                                            case -275:
                                                v66 = (v61 % ((1) | 1))
                                            default:
                                                v66 -= v61
                                            }
                                        case "a":
                                            v65 += v65
                                        }
                                    }
                                } else {
                                    var v71 uint = v65
                                    _ = v71
                                }
                            case "":
                                if true {
                                    if false {
                                        if false {
                                            v7 = v7
                                            v72 := GenS9{F0: "a", F1: -32768}
                                            _ = v72
                                        } else {
                                            v73 := GenS11{F0: (v65 >> (((v65 >> ((v65) & 63))) & 63)), F1: true}
                                            _ = v73
                                        }
                                        var v74 float64 = 3.14159
                                        var v75 int32 = (v53 << ((v65) & 63))
                                        var v76 int64 = 42
                                        v77 := GenS12{F0: 418, F1: 2, F2: "\n\t"}
                                        var v78 int16 = 1000
                                        v21 = append(v21, (fi22 & -2147483648))
                                        _ = v74
                                        _ = v75
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                    } else {
                                        switch ("foo bar" + "") {
                                        case "a":
                                            v65 += (v65 & v65)
                                        }
                                        var v79 uint64 = 4294967295
                                        _ = v79
                                    }
                                    var v80 int32 = int32(fi45)
                                    v81 := GenS12{F0: (v65 | v65), F1: 18446744073709551615, F2: ("hello" + "a")}
                                    _ = v80
                                    _ = v81
                                } else {
                                    var v82 uint32 = 363
                                    v82 += (v82 / ((v82) | 1))
                                    _ = v82
                                }
                            }
                            v83 := GenS13{F0: -1, F1: -1}
                            var v84 uint32 = 4294967295
                            _ = v66
                            _ = v83
                            _ = v84
                        case "a":
                            v85 := GenS14{F0: (fi22 + int(v53)), F1: 0, F2: v65}
                            _ = v85
                        default:
                            _ = len(v21)
                        }
                        v86 := GenS15{F0: true, F1: ("\n\t" + ""), F2: -14}
                        _ = v64
                        _ = v65
                        _ = v86
                    } else {
                        if true {
                            var v87 uint = 4294967295
                            _ = v87
                        }
                        var v88 uint16 = 2
                        switch (fi22 | fi45) {
                        case 128:
                            var v89 int64 = 42
                            _ = v89
                        }
                        _ = v88
                    }
                    switch "" {
                    case "":
                        v61 = 128
                        fmt.Println(v53, v7)
                        var v90 int8 = -1
                        _ = v90
                    }
                    switch -28 {
                    case -436:
                        v91 := GenS3{F0: 4294967295}
                        _ = v91
                    case 18:
                        if true {
                            if (2 <= 42) {
                                v92 := GenS16{F0: -121}
                                v93 := GenS17{F0: 593, F1: 1, F2: v61}
                                var v94 float64 = (-2.25 - 0.1)
                                _ = v92
                                _ = v93
                                _ = v94
                            }
                        }
                    }
                    _ = fi45
                    _ = v53
                    _ = v61
                    _ = v62
                    _ = v63
                }
                if len(v21) > 0 {
                    v21[0] = (^(fi22))
                }
            }
            _ = v23
        }
        _ = fi22
    }
    _ = v21
    _ = v7
    return 65535
}
func genFn1(p0 int16, p1 uint32, p2 uint64) int16 {
    var v95 string = "\n\t"
    if false {
        if true {
            switch 128 {
            case -40:
                v95 = v95
            case -9223372036854775808:
                fmt.Println(p2, v95)
            case -2147483648:
                if ((!true) && true) {
                    fmt.Println(p0)
                    var v96 uint = 99
                    _ = v96
                } else {
                    var v97 uint8 = 1
                    if (!((v97 - (v97 << ((0) & 63))) > 0)) {
                        fmt.Println(p1, v97)
                    } else {
                        if (false || (20 < 100)) {
                            v98 := []int{-1, 0, 9223372036854775807}
                            _ = v98
                        } else {
                            for fi99 := 0; fi99 < 8; fi99++ {
                                v100 := []int{fi99, int(v97)}
                                _ = fi99
                                _ = v100
                            }
                        }
                    }
                    if (true || true) {
                        var v101 uint16 = 0
                        var v102 uint32 = (p1 + (p1 % (((p1 << ((1) & 63))) | 1)))
                        _ = v101
                        _ = v102
                    } else {
                        var v103 float64 = ((((-2.25 + 1e10) + (3.14159 / 1.0)) - 3.14159) * 0.0)
                        _ = v103
                    }
                    _ = v97
                }
            }
            v104 := []int{128, 2147483647}
            v105 := GenS18{F0: -2147483648, F1: 42}
            _ = v104
            _ = v105
        }
        for fi106 := 0; fi106 < 9; fi106++ {
            p0 += p0
            v107 := GenS19{F0: -1, F1: p2}
            var v108 string = ("\n\t" + v95)
            _ = fi106
            _ = v107
            _ = v108
        }
    } else {
        v95 = v95
        for fi109 := 0; fi109 < 7; fi109++ {
            for fi110 := 0; fi110 < 20; fi110++ {
                v95 = "foo bar"
                _ = fi110
            }
            if (genFn0() > 65535) {
                switch v95 {
                case "":
                    if (1 <= uint8(p2)) {
                        v111 := GenS20{F0: false, F1: (1 == 201), F2: p0}
                        var v112 uint64 = 451
                        _ = v111
                        _ = v112
                    }
                }
                v113 := GenS8{F0: 42}
                _ = v113
            }
            switch (fi109 % ((fi109) | 1)) {
            case 9223372036854775807:
                v114 := GenS21{F0: (^(p0)), F1: 705, F2: 1}
                _ = v114
            case -1:
                if (42 >= 9223372036854775807) {
                    switch ("hello" + v95) {
                    case "":
                        var v115 int16 = int16(p1)
                        _ = v115
                    case "foo bar":
                        fmt.Println(v95, p0, p1)
                    case "a":
                        for fi116 := 0; fi116 < 4; fi116++ {
                            var v117 uint64 = (p2 / ((602) | 1))
                            switch ("a" + v95) {
                            case "\n\t":
                                p2 += p2
                            case "hello":
                                p1 = p1
                            }
                            _ = fi116
                            _ = v117
                        }
                    default:
                        var v118 uint8 = 128
                        _ = v118
                    }
                } else {
                    fmt.Println(v95, fi109)
                    v119 := GenS22{F0: -9223372036854775808, F1: p2}
                    _ = v119
                }
            case -9223372036854775808:
                var v120 uint = 4294967295
                _ = v120
            }
            _ = fi109
        }
    }
    p2 = (p2 & 0)
    fmt.Println(p2, p2, v95)
    _ = p0
    _ = p1
    _ = p2
    _ = v95
    return (p0 ^ p0)
}

func main() {
    v1 := []int{}
    var v2 bool = true
    var v3 int32 = 2147483647
    v4 := []string{"\n\t", (("hello" + (("" + "") + "a")) + "\n\t"), ("" + "\n\t")}
    var v5 int8 = 100
    var v6 uint16 = 65535
    v6 += v6
    for fi121 := 0; fi121 < 18; fi121++ {
        if true {
            switch "a" {
            case "\n\t":
                v5 -= v5
            case "a":
                fmt.Println(v5, v2)
            case "foo bar":
                var v122 int32 = v3
                _ = v122
            }
            if (-439 < (v3 / ((v3) | 1))) {
                v123 := GenS18{F0: fi121, F1: 0}
                if ((("\n\t" == "hello") || (!v2)) || false) {
                    switch (("foo bar" + "\n\t") + ("a" + (("\n\t" + "hello") + ("foo bar" + ((("foo bar" + (("hello" + "\n\t") + ((("a" + (("a" + "") + "hello")) + "") + ("foo bar" + "a")))) + ("\n\t" + "\n\t")) + "\n\t"))))) {
                    case "a":
                        v124 := GenS23{F0: false}
                        _ = v124
                    }
                    v6 = (v6 ^ (v6 - genFn0()))
                    fmt.Println(v2)
                } else {
                    var v125 bool = v2
                    switch "\n\t" {
                    case "foo bar":
                        if ((-1 <= 870) && ((!true) || (!(v3 <= -339)))) {
                            for fi126 := 0; fi126 < 10; fi126++ {
                                v125 = v125
                                v1 = append(v1, (fi126 << ((829) & 63)))
                                _ = fi126
                            }
                            v127 := GenS24{F0: 1000}
                            switch "" {
                            case "\n\t":
                                var v128 float32 = float32(-2.25)
                                _ = v128
                            case "foo bar":
                                for fi129 := 0; fi129 < 10; fi129++ {
                                    v5 += v5
                                    var v130 float64 = 3.14159
                                    _ = fi129
                                    _ = v130
                                }
                            }
                            _ = v127
                        } else {
                            switch (fi121 * (^(fi121))) {
                            case 32768:
                                var v131 int64 = -657
                                _ = v131
                            case 9223372036854775807:
                                if ((0.0 * float64(1e10)) < ((3.14159 * 3.14159) - float64(1e10))) {
                                    var v132 uint = 932
                                    v133 := GenS25{F0: v3, F1: -234, F2: 0}
                                    v3 += v3
                                    _ = v132
                                    _ = v133
                                } else {
                                    for fi134 := 0; fi134 < 14; fi134++ {
                                        if false {
                                            v135 := GenS20{F0: (!(!(0 == 2))), F1: v2, F2: -527}
                                            fmt.Println(v6)
                                            _ = v135
                                        }
                                        v2 = false
                                        _ = fi134
                                    }
                                    for fi136 := 0; fi136 < 20; fi136++ {
                                        v4 = append(v4, (("a" + "\n\t") + "a"))
                                        _ = fi136
                                    }
                                }
                            default:
                                _ = len(v1)
                            }
                            var v137 bool = (!v125)
                            if v2 {
                                var v138 bool = v2
                                if v137 {
                                    fmt.Println(v137)
                                }
                                _ = v138
                            }
                            _ = v137
                        }
                    case "hello":
                        v3 = v3
                    case "a":
                        v139 := GenS23{F0: (v2 || true)}
                        _ = v139
                    }
                    _ = v125
                }
                _ = v123
            }
        }
        _ = fi121
    }
    v140 := GenS19{F0: 127, F1: 866}
    v6 = 454
    var v141 uint = 65535
    fmt.Println(v2, v3, v6)
    switch "\n\t" {
    case "":
        var v142 float64 = (-2.25 / 1e10)
        _ = v142
    case "a":
        var v143 float64 = 1.5
        _ = v143
    case "foo bar":
        if v2 {
            fmt.Println(v141)
        } else {
            v144 := GenS26{F0: v141}
            var v145 float64 = ((1.5 - ((3.14159 + 0.1) + 1e10)) + -2.25)
            _ = v144
            _ = v145
        }
    default:
        _ = len(v1)
    }
    switch 187 {
    case -1:
        v1 = append(v1, 950)
    case 9223372036854775807:
        for fi146 := 0; fi146 < 9; fi146++ {
            if v2 {
                var v147 int64 = 42
                v1 = append(v1, (fi146 >> ((0) & 63)))
                v148 := GenS24{F0: -32768}
                _ = v147
                _ = v148
            } else {
                var v149 float32 = (-2.25 - float32(1e10))
                v150 := GenS27{F0: v5, F1: v2}
                var v151 float32 = 1e10
                _ = v149
                _ = v150
                _ = v151
            }
            _ = len(v4)
            _ = fi146
        }
    case 51:
        v6 += v6
    default:
        v152 := GenS28{F0: 36, F1: 1000}
        _ = v152
    }
    v5 = v5
    if (!v2) {
        if (949 >= 0) {
            for fi153 := 0; fi153 < 20; fi153++ {
                switch "" {
                case "a":
                    v154 := GenS29{F0: true, F1: 621}
                    _ = v154
                default:
                    v155 := GenS30{F0: ("hello" + "hello"), F1: -9223372036854775808, F2: (v6 & v6)}
                    _ = v155
                }
                var v156 int64 = -9223372036854775808
                v156 = 9223372036854775807
                _ = fi153
                _ = v156
            }
            v141 += 1
        } else {
            switch -1 {
            case 744:
                if ((1e10 > 3.14159) || ((211 == (v141 - (^(v141)))) || v2)) {
                    var v157 string = ("hello" + "\n\t")
                    _ = v157
                }
            default:
                for fi158 := 0; fi158 < 11; fi158++ {
                    _ = len(v4)
                    _ = fi158
                }
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v141:", v141)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v140
    _ = v141
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
