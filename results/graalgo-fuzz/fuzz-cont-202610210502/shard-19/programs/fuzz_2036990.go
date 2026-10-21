package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 string; F1 string; F2 int32 }
type GenS9 struct { F0 uint8; F1 int }
type GenS10 struct { F0 int64; F1 uint16; F2 int64 }
type GenS11 struct { F0 int8; F1 bool }
func genFn0(p0 int, p1 int64, p2 int8) uint8 {
    v7 := GenS1{F0: 0, F1: 18446744073709551615, F2: 109}
    for fi8 := 0; fi8 < 3; fi8++ {
        p1 = (p1 >> ((1) & 63))
        _ = fi8
    }
    v9 := GenS1{F0: 2147483647, F1: 415, F2: 42}
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS2{F0: 4294967295, F1: (1 < 2)}
        v12 := []int{(p0 ^ p0), 9223372036854775807}
        _ = fi10
        _ = v11
        _ = v12
    }
    if false {
        p2 += 92
    } else {
        v13 := []int{}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return 55
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float32 = (-2.25 * 3.14159)
    v3 := make(map[int]string)
    var v4 bool = false
    var v5 float64 = 0.0
    var v6 int16 = int16(v1)
    switch 1 {
    case -2147483648:
        for fi14 := 0; fi14 < 17; fi14++ {
            v15 := GenS1{F0: 0, F1: 1, F2: genFn0((^(fi14)), (v1 + v1), 1)}
            var v16 float64 = (v5 - v5)
            _ = fi14
            _ = v15
            _ = v16
        }
    case 42:
        if v4 {
            if v4 {
                for fi17 := 0; fi17 < 17; fi17++ {
                    fmt.Println(v5, v4)
                    if v4 {
                        var v18 uint64 = 530
                        v2 += v2
                        var v19 int32 = 2147483647
                        _ = v18
                        _ = v19
                    } else {
                        v20 := GenS1{F0: 0, F1: 0, F2: 1}
                        var v21 uint16 = 607
                        _ = v20
                        _ = v21
                    }
                    fmt.Println(v2)
                    _ = fi17
                }
            } else {
                fmt.Println(v4)
                v1 = v1
                for fi22 := 0; fi22 < 18; fi22++ {
                    if false {
                        if true {
                            v2 = v2
                            var v23 int16 = v6
                            var v24 bool = v4
                            _ = v23
                            _ = v24
                        }
                        switch (fi22 + int(v1)) {
                        case -9223372036854775808:
                            v25 := []int{int(v1), 4294967295, (-(fi22))}
                            _ = v25
                        case 2147483647:
                            v4 = (116 < 17)
                        default:
                            v26 := []int{(fi22 + fi22), (fi22 * fi22), -2147483648, 956}
                            _ = v26
                        }
                        v5 = v5
                    } else {
                        fmt.Println(v5, v1, v6)
                        for fi27 := 0; fi27 < 14; fi27++ {
                            v1 += (v1 - (-(v1)))
                            if v4 {
                                v4 = v4
                            }
                            v6 = v6
                            _ = fi27
                        }
                        fmt.Println(v1, v4, v5)
                    }
                    var v28 uint64 = 4294967295
                    _ = fi22
                    _ = v28
                }
            }
        }
    case 128:
        var v29 bool = (18446744073709551615 != 18446744073709551615)
        _ = v29
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        for fi31 := 0; fi31 < 14; fi31++ {
            var v32 float64 = v5
            v32 = (((v32 * v32) - (v32 / float64(-2.25))) * (float64(v2) / v32))
            if (!false) {
                var v33 string = "hello"
                _ = v33
            }
            _ = fi31
            _ = v32
        }
        var v34 string = "hello"
        _ = fi30
        _ = v34
    }
    v2 = ((v2 * 3.14159) * (float32(1.5) - v2))
    if false {
        for fi35 := 0; fi35 < 20; fi35++ {
            for fi36 := 0; fi36 < 7; fi36++ {
                switch "a" {
                case "a":
                    switch (fi35 ^ fi35) {
                    case 32767:
                        switch "\n\t" {
                        case "foo bar":
                            switch ("" + "\n\t") {
                            case "a":
                                switch fi35 {
                                case 0:
                                    v2 -= 1e10
                                case -700:
                                    v1 += v1
                                default:
                                    v37 := []int{}
                                    _ = v37
                                }
                            case "foo bar":
                                switch "" {
                                case "a":
                                    var v38 int8 = 100
                                    _ = v38
                                case "\n\t":
                                    v4 = v4
                                default:
                                    v3[127] = ""
                                }
                            default:
                                _ = len(v3)
                            }
                        case "":
                            v39 := GenS3{F0: 42}
                            _ = v39
                        case "a":
                            v1 += (v1 * 1)
                        default:
                            v1 -= 1
                        }
                    default:
                        v40 := []int{fi35, -998, 469, (fi35 / ((fi36) | 1))}
                        _ = v40
                    }
                case "":
                    switch (fi36 >> ((434) & 63)) {
                    case 1:
                        var v41 int32 = 88
                        _ = v41
                    case -2147483648:
                        v5 += (float64((1.5 + v2)) + v5)
                        var v42 int64 = (-(v1))
                        if false {
                            if ((127 != 68) && (-384 >= 837)) {
                                switch ((("\n\t" + "a") + "a") + "hello") {
                                case "\n\t":
                                    var v43 uint8 = 1
                                    _ = v43
                                case "foo bar":
                                    switch (fi36 - (fi36 << ((1) & 63))) {
                                    case 4294967295:
                                        v6 = (v6 & v6)
                                    case -2147483648:
                                        v4 = (-1 != v42)
                                    case -51:
                                        if v4 {
                                            fmt.Println(fi36)
                                        } else {
                                            switch (fi36 | (fi35 / (((fi35 % (((fi35 | fi36)) | 1))) | 1))) {
                                            case 9223372036854775807:
                                                var v44 uint16 = 32767
                                                _ = v44
                                            case 2147483647:
                                                v45 := []int{441, 32767}
                                                _ = v45
                                            case 127:
                                                v46 := []int{(^(fi36)), fi36}
                                                _ = v46
                                            default:
                                                var v47 uint = 33
                                                _ = v47
                                            }
                                            switch 9223372036854775807 {
                                            case 32768:
                                                v3[(^(fi35))] = (("hello" + ("hello" + "\n\t")) + "\n\t")
                                            case 1:
                                                var v48 bool = true
                                                _ = v48
                                            case 276:
                                                fmt.Println(v6, v6)
                                            default:
                                                switch -102 {
                                                case 419:
                                                    if (!v4) {
                                                        v49 := []int{(fi36 - fi35), 2147483647}
                                                        fmt.Println(fi35)
                                                        _ = v49
                                                    }
                                                case -854:
                                                    fmt.Println(fi35)
                                                case 472:
                                                    v42 = 0
                                                }
                                            }
                                        }
                                    }
                                case "":
                                    v42 -= (v1 & v1)
                                default:
                                    v42 = (v1 >> ((0) & 63))
                                }
                                if (!(v6 == 42)) {
                                    v50 := []int{}
                                    _ = len(v3)
                                    _ = v50
                                }
                                v51 := GenS2{F0: 65535, F1: v4}
                                fmt.Println(v1)
                                v52 := v3[fi35]
                                v42 = v1
                                _ = v51
                                _ = v52
                            } else {
                                v53 := GenS2{F0: 464, F1: v4}
                                _ = v53
                            }
                            v6 += (v6 - int16(fi36))
                            var v54 string = "foo bar"
                            _ = v54
                        }
                        _ = len(v3)
                        _ = v42
                    default:
                        if (v6 <= -1) {
                            fmt.Println(v1)
                            delete(v3, (fi36 / (((fi35 << ((65535) & 63))) | 1)))
                            var v55 float64 = 1.5
                            _ = v55
                        }
                    }
                    v56 := GenS2{F0: 0, F1: (!(!v4))}
                    _ = v56
                default:
                    v5 -= 1.5
                }
                _ = fi36
            }
            _ = fi35
        }
        var v57 uint16 = 42
        for fi58 := 0; fi58 < 12; fi58++ {
            v59 := []int{fi58, (fi58 ^ (-(fi58)))}
            if v4 {
                v6 += (v6 - 0)
                for fi60 := 0; fi60 < 5; fi60++ {
                    var v61 uint64 = 450
                    _ = fi60
                    _ = v61
                }
            } else {
                var v62 int = 128
                _ = v62
            }
            if (4294967295 < 788) {
                var v63 uint64 = 0
                _ = v63
            }
            _ = fi58
            _ = v59
        }
        _ = v57
    } else {
        v64 := []int{42, 0}
        if v4 {
            switch "hello" {
            case "":
                var v65 bool = true
                _ = v65
            case "hello":
                var v66 uint16 = 255
                _ = v66
            default:
                for fi67 := 0; fi67 < 20; fi67++ {
                    for fi68 := 0; fi68 < 10; fi68++ {
                        if (v6 > (v6 - v6)) {
                            fmt.Println(v5, v5, v2)
                        }
                        v64 = append(v64, 32768)
                        _ = fi68
                    }
                    _ = fi67
                }
            }
        } else {
            if v4 {
                switch 0 {
                case 42:
                    if v4 {
                        var v69 uint64 = 4294967295
                        _ = v69
                    } else {
                        var v70 int64 = v1
                        _ = v70
                    }
                case 32767:
                    for fi71 := 0; fi71 < 13; fi71++ {
                        if len(v64) >= 2 {
                            v64 = v64[1:len(v64)]
                        }
                        if true {
                            if v4 {
                                fmt.Println(v6)
                            }
                            if false {
                                if v4 {
                                    v4 = v4
                                    var v72 int32 = 78
                                    _ = v72
                                }
                                var v73 int32 = -2147483648
                                v6 += v6
                                _ = v73
                            }
                            var v74 string = ("" + "a")
                            _ = v74
                        }
                        v5 -= (v5 * v5)
                        _ = fi71
                    }
                default:
                    _ = len(v64)
                }
                for fi75 := 0; fi75 < 1; fi75++ {
                    var v76 int8 = 17
                    _ = fi75
                    _ = v76
                }
            }
            v3[-216] = "foo bar"
            v77 := GenS4{F0: 918, F1: "\n\t"}
            _ = v77
        }
        var v78 int = 32768
        _ = v64
        _ = v78
    }
    fmt.Println(v1)
    switch 2147483647 {
    case 551:
        if ((-2.25 * ((3.14159 / (((v2 / (-2.25 - (v2 / v2))) / (v2 - v2)) / ((((v2 * float32(v5)) + v2) + v2) / (v2 - (v2 * ((v2 / v2) / (1e10 / (0.1 * (1e10 * (v2 / (-2.25 + v2))))))))))) / ((1.5 - (v2 - 1e10)) + 3.14159))) > (v2 + (-2.25 - v2))) {
            var v79 int32 = -592
            _ = v79
        } else {
            if v4 {
                for fi80 := 0; fi80 < 19; fi80++ {
                    for fi81 := 0; fi81 < 5; fi81++ {
                        v82 := GenS3{F0: 0}
                        v4 = v4
                        v83 := v3[fi80]
                        _ = fi81
                        _ = v82
                        _ = v83
                    }
                    if (53 < 0) {
                        for fi84 := 0; fi84 < 14; fi84++ {
                            if v4 {
                                if (true || ((v6 >> ((4294967295) & 63)) != -441)) {
                                    if (v4 && true) {
                                        v1 -= (v1 - v1)
                                        v85 := GenS2{F0: 1, F1: false}
                                        _ = v85
                                    } else {
                                        v86 := GenS3{F0: 450}
                                        var v87 int32 = 2147483647
                                        _ = v86
                                        _ = v87
                                    }
                                }
                                var v88 uint = 42
                                fmt.Println(v1, v4)
                                _ = v88
                            }
                            v89 := []int{fi84, fi84, fi80, fi80}
                            _ = fi84
                            _ = v89
                        }
                        v6 = v6
                    } else {
                        var v90 string = ""
                        var v91 float32 = (1e10 + 3.14159)
                        if (127 >= -108) {
                            if (v4 && v4) {
                                if (!v4) {
                                    v92 := GenS4{F0: 4294967295, F1: v90}
                                    var v93 uint64 = 2
                                    v6 = -32768
                                    _ = v92
                                    _ = v93
                                } else {
                                    var v94 float32 = 0.1
                                    v95 := GenS5{F0: fi80, F1: 2}
                                    var v96 int = 217
                                    _ = v94
                                    _ = v95
                                    _ = v96
                                }
                                v91 = v91
                            } else {
                                for fi97 := 0; fi97 < 1; fi97++ {
                                    v5 = v5
                                    v1 += -104
                                    v98 := GenS5{F0: (fi80 ^ (^(fi97))), F1: 0}
                                    _ = fi97
                                    _ = v98
                                }
                            }
                            if ((!(260 <= 1)) || (v6 >= v6)) {
                                v99 := GenS6{F0: int16(fi80)}
                                for fi100 := 0; fi100 < 4; fi100++ {
                                    v101 := []int{(fi100 / ((127) | 1))}
                                    v4 = v4
                                    _ = fi100
                                    _ = v101
                                }
                                _ = v99
                            }
                            switch 2147483647 {
                            case -9223372036854775808:
                                v91 = v2
                            case 42:
                                if v4 {
                                    switch v90 {
                                    case "hello":
                                        v102 := GenS1{F0: 42, F1: 1, F2: 128}
                                        _ = v102
                                    case "a":
                                        v103 := []int{fi80, (fi80 + (^(fi80))), fi80, fi80}
                                        _ = v103
                                    case "":
                                        v104 := GenS3{F0: 42}
                                        _ = v104
                                    }
                                    switch v90 {
                                    case "hello":
                                        var v105 uint16 = 255
                                        _ = v105
                                    case "\n\t":
                                        var v106 float64 = 1e10
                                        _ = v106
                                    default:
                                        if true {
                                            switch (fi80 ^ fi80) {
                                            case 2147483647:
                                                switch ("\n\t" + ("a" + "")) {
                                                case "hello":
                                                    fmt.Println(v2)
                                                case "":
                                                    v2 -= v2
                                                case "\n\t":
                                                    v91 = v91
                                                default:
                                                    v4 = v4
                                                }
                                            case 932:
                                                v5 = (float64((v91 + v2)) - v5)
                                            default:
                                                v107 := []int{}
                                                _ = v107
                                            }
                                        }
                                    }
                                    var v108 uint8 = 1
                                    _ = v108
                                } else {
                                    for fi109 := 0; fi109 < 6; fi109++ {
                                        v110 := GenS5{F0: fi80, F1: 4294967295}
                                        _ = fi109
                                        _ = v110
                                    }
                                    if false {
                                        if false {
                                            if v4 {
                                                if (172 > 26) {
                                                    if true {
                                                        v111 := GenS7{F0: 2, F1: 2}
                                                        var v112 uint8 = 2
                                                        v5 += ((v5 + (v5 + float64((1e10 + v91)))) / 0.1)
                                                        _ = v111
                                                        _ = v112
                                                    } else {
                                                        if v4 {
                                                            var v113 int64 = v1
                                                            if (((v90 == (v90 + v90)) && v4) && ((("hello" + ("foo bar" + (v90 + ("" + v90)))) + v90) <= "\n\t")) {
                                                                v113 = -9223372036854775808
                                                            }
                                                            _ = v113
                                                        } else {
                                                            v5 += v5
                                                            v91 = (3.14159 * (v91 * v91))
                                                        }
                                                        v91 = v2
                                                        v114 := v3[(-(fi80))]
                                                        _ = v114
                                                    }
                                                }
                                            }
                                        } else {
                                            switch ("" + "foo bar") {
                                            case "foo bar":
                                                var v115 uint16 = 32767
                                                _ = v115
                                            case "\n\t":
                                                var v116 int = (^(fi80))
                                                _ = v116
                                            case "a":
                                                switch (-(fi80)) {
                                                case -1:
                                                    v117 := []int{fi80, (fi80 + fi80)}
                                                    _ = v117
                                                case 4294967295:
                                                    for fi118 := 0; fi118 < 19; fi118++ {
                                                        fmt.Println(fi80)
                                                        var v119 bool = true
                                                        var v120 string = "hello"
                                                        v121 := []int{fi80, fi118, 0}
                                                        v5 = v5
                                                        _ = fi118
                                                        _ = v119
                                                        _ = v120
                                                        _ = v121
                                                    }
                                                case -2147483648:
                                                    fmt.Println(v90)
                                                default:
                                                    if v4 {
                                                        for fi122 := 0; fi122 < 7; fi122++ {
                                                            v123 := GenS8{F0: "\n\t", F1: v90, F2: 0}
                                                            v124 := GenS9{F0: 202, F1: -430}
                                                            fmt.Println(v91)
                                                            _ = fi122
                                                            _ = v123
                                                            _ = v124
                                                        }
                                                        v2 = (v91 / (v91 * 0.0))
                                                        var v125 uint16 = 32767
                                                        _ = v125
                                                    }
                                                }
                                            default:
                                                var v126 uint16 = 79
                                                _ = v126
                                            }
                                            var v127 int64 = 42
                                            if (!((-2147483648 <= 0) || true)) {
                                                if (!true) {
                                                    v2 = (v2 / (v91 + (v91 - (((3.14159 + (v91 * (v2 * (1e10 / v91)))) - v91) - (v2 / (v2 - (float32(v5) / 3.14159)))))))
                                                    v128 := GenS4{F0: 42, F1: "a"}
                                                    v127 += -9223372036854775808
                                                    _ = v128
                                                } else {
                                                    var v129 string = v90
                                                    var v130 uint64 = 2
                                                    fmt.Println(fi80, v127)
                                                    _ = v129
                                                    _ = v130
                                                }
                                            }
                                            _ = v127
                                        }
                                        v6 += (-(v6))
                                    }
                                    for fi131 := 0; fi131 < 6; fi131++ {
                                        fmt.Println(v2)
                                        v6 -= (v6 / ((-1) | 1))
                                        v3[32768] = v90
                                        _ = fi131
                                    }
                                }
                            default:
                                var v132 int16 = 32767
                                _ = v132
                            }
                        } else {
                            v133 := []int{fi80}
                            if (!((fi80 & (fi80 ^ (^(fi80)))) == (-(fi80)))) {
                                if ((42 != 815) && true) {
                                    var v134 uint64 = 42
                                    if ((-534 < -878) && ((v90 + "") == v90)) {
                                        if ((!true) && (v4 || (1000 >= int16(fi80)))) {
                                            var v135 bool = (v4 || v4)
                                            _ = v135
                                        } else {
                                            _ = len(v3)
                                            fmt.Println(v5, fi80)
                                            v90 = "hello"
                                        }
                                        if len(v133) >= 2 {
                                            v133 = v133[1:len(v133)]
                                        }
                                        var v136 float32 = -2.25
                                        _ = v136
                                    }
                                    _ = v134
                                }
                                if (255 > 153) {
                                    var v137 uint32 = 2147483648
                                    v137 = (v137 / (((v137 + (v137 + 0))) | 1))
                                    fmt.Println(v137, v91)
                                    _ = v137
                                } else {
                                    v1 = v1
                                    for fi138 := 0; fi138 < 8; fi138++ {
                                        v139 := GenS2{F0: 42, F1: true}
                                        _ = fi138
                                        _ = v139
                                    }
                                    v140 := GenS10{F0: (v1 ^ 367), F1: 255, F2: (v1 >> ((65535) & 63))}
                                    _ = v140
                                }
                                v141 := v3[-2147483648]
                                _ = v141
                            } else {
                                _ = len(v133)
                                for fi142 := 0; fi142 < 7; fi142++ {
                                    if ((v1 & v1) != v1) {
                                        fmt.Println(v2, v90)
                                        fmt.Println(fi142, v2, v90)
                                        v133 = append(v133, fi142)
                                        var v143 string = v90
                                        _ = v143
                                    } else {
                                        var v144 uint64 = 4294967295
                                        _ = v144
                                    }
                                    if true {
                                        v91 = 1.5
                                        var v145 uint8 = 2
                                        v133 = append(v133, (fi80 ^ fi80))
                                        v4 = v4
                                        _ = v145
                                    }
                                    if len(v133) > 0 {
                                        v133[0] = -419
                                    }
                                    _ = len(v133)
                                    _ = fi142
                                }
                                switch 42 {
                                case -768:
                                    if len(v133) > 0 {
                                        v133[0] = -2147483648
                                    }
                                case -1:
                                    if true {
                                        fmt.Println(v5)
                                        v146 := GenS6{F0: (v6 + 829)}
                                        _ = v146
                                    } else {
                                        v147 := GenS7{F0: 18446744073709551615, F1: 0}
                                        for fi148 := 0; fi148 < 17; fi148++ {
                                            v4 = ((false && (91 != 1)) || v4)
                                            v133 = append(v133, fi80)
                                            var v149 int8 = 42
                                            _ = fi148
                                            _ = v149
                                        }
                                        v90 = v90
                                        _ = v147
                                    }
                                case 32767:
                                    fmt.Println(v2, v90)
                                default:
                                    var v150 int64 = (^(v1))
                                    _ = v150
                                }
                            }
                            v4 = true
                            _ = v133
                        }
                        _ = v90
                        _ = v91
                    }
                    if true {
                        v151 := v3[-2147483648]
                        _ = v151
                    } else {
                        var v152 int = int(v1)
                        var v153 float32 = 0.0
                        _ = v152
                        _ = v153
                    }
                    _ = fi80
                }
            }
            var v154 uint = 1
            _ = v154
        }
    }
    var v155 int = -497
    var v156 uint8 = genFn0(v155, 381, -1)
    for fi157 := 0; fi157 < 19; fi157++ {
        switch fi157 {
        case 960:
            fmt.Println(v5, v1, v1)
        default:
            var v158 uint = 4294967295
            _ = v158
        }
        if ((v155 ^ fi157) < fi157) {
            v159 := GenS9{F0: v156, F1: 128}
            _ = v159
        } else {
            for fi160 := 0; fi160 < 11; fi160++ {
                v161 := []int{fi157, -278, 127}
                _ = fi160
                _ = v161
            }
            switch "hello" {
            case "a":
                fmt.Println(v4, v4)
            }
        }
        _ = fi157
    }
    v6 -= 0
    switch ("foo bar" + ((("hello" + "") + "\n\t") + "foo bar")) {
    case "\n\t":
        if v4 {
            switch v155 {
            case 32767:
                if true {
                    for fi162 := 0; fi162 < 9; fi162++ {
                        fmt.Println(v1, v6, v156)
                        for fi163 := 0; fi163 < 4; fi163++ {
                            var v164 string = "foo bar"
                            if (746 > 4294967295) {
                                var v165 uint64 = uint64(v156)
                                v166 := GenS9{F0: v156, F1: fi163}
                                var v167 uint16 = 779
                                _ = v165
                                _ = v166
                                _ = v167
                            }
                            _ = fi163
                            _ = v164
                        }
                        _ = fi162
                    }
                } else {
                    v168 := GenS7{F0: 230, F1: 65535}
                    _ = v168
                }
            case 2147483647:
                var v169 bool = v4
                _ = v169
            case 128:
                fmt.Println(v1, v1, v156)
            default:
                var v170 uint16 = 32767
                _ = v170
            }
        } else {
            v155 -= v155
            _ = len(v3)
            for fi171 := 0; fi171 < 6; fi171++ {
                if true {
                    var v172 int64 = (^(v1))
                    if v4 {
                        var v173 float64 = float64(1e10)
                        if ((v1 << ((2) & 63)) > (^(v1))) {
                            v5 += 0.1
                        } else {
                            if v4 {
                                var v174 int8 = -102
                                var v175 int32 = -2147483648
                                _ = v174
                                _ = v175
                            } else {
                                fmt.Println(v156, v5, v156)
                                switch (("hello" + "foo bar") + "\n\t") {
                                case "foo bar":
                                    if v4 {
                                        for fi176 := 0; fi176 < 7; fi176++ {
                                            v5 -= 1.5
                                            _ = fi176
                                        }
                                        var v177 uint8 = v156
                                        _ = v177
                                    } else {
                                        for fi178 := 0; fi178 < 6; fi178++ {
                                            var v179 uint16 = 255
                                            _ = fi178
                                            _ = v179
                                        }
                                        switch "foo bar" {
                                        case "hello":
                                            var v180 int16 = (v6 | 1)
                                            _ = v180
                                        case "":
                                            if true {
                                                switch "hello" {
                                                case "foo bar":
                                                    v1 += 1
                                                case "hello":
                                                    if v4 {
                                                        switch fi171 {
                                                        case -1:
                                                            var v181 float64 = v173
                                                            _ = v181
                                                        case 0:
                                                            if (18446744073709551615 != 0) {
                                                                var v182 uint = 65535
                                                                _ = v182
                                                            } else {
                                                                v1 = (^(v1))
                                                                var v183 uint = 847
                                                                _ = v183
                                                            }
                                                        default:
                                                            v184 := GenS11{F0: 101, F1: v4}
                                                            _ = v184
                                                        }
                                                        v4 = false
                                                        v5 += v5
                                                    }
                                                }
                                            }
                                        default:
                                            var v185 uint16 = uint16(v156)
                                            _ = v185
                                        }
                                    }
                                case "hello":
                                    if (!(-43 == int8(v6))) {
                                        v2 += -2.25
                                    }
                                }
                            }
                        }
                        _ = v173
                    }
                    _ = v172
                }
                for fi186 := 0; fi186 < 8; fi186++ {
                    var v187 int8 = -60
                    _ = fi186
                    _ = v187
                }
                _ = fi171
            }
        }
    case "foo bar":
        var v188 int16 = v6
        _ = v188
    case "":
        fmt.Println(v2)
    }
    v189 := GenS3{F0: 4294967295}
    v190 := GenS10{F0: v1, F1: 32767, F2: 42}
    if ((false || (v4 || v4)) && (9223372036854775807 != (v1 % ((v1) | 1)))) {
        v4 = v4
        v1 = (v1 * (v1 * -9223372036854775808))
        switch ("hello" + "hello") {
        case "a":
            v6 = v6
        case "hello":
            v191 := GenS5{F0: (v155 >> ((976) & 63)), F1: 42}
            _ = v191
        case "\n\t":
            v192 := GenS5{F0: -384, F1: 571}
            _ = v192
        default:
            var v193 int = 0
            _ = v193
        }
    } else {
        v2 += v2
        var v194 uint = 0
        _ = v194
    }
    v195 := GenS5{F0: int(v1), F1: 933}
    if false {
        switch v155 {
        case 935:
            var v196 int32 = 0
            _ = v196
        case 32767:
            if v4 {
                var v197 uint = 65535
                v155 -= v155
                v198 := GenS3{F0: uint64(v156)}
                _ = v197
                _ = v198
            }
        default:
            var v199 int8 = 127
            _ = v199
        }
    }
    v2 += (((1.5 * 3.14159) + ((v2 - ((-2.25 - 1e10) + v2)) / v2)) + v2)
    var v200 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v155:", v155)
    fmt.Println("v156:", v156)
    fmt.Println("v2:", v2)
    fmt.Println("v200:", v200)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v155
    _ = v156
    _ = v189
    _ = v190
    _ = v195
    _ = v2
    _ = v200
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
