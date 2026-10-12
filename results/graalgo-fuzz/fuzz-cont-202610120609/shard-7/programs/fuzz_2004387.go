package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint8; F1 string; F2 uint8 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 uint32; F1 int8 }
type GenS8 struct { F0 bool; F1 string }
type GenS9 struct { F0 bool; F1 int64 }
type GenS10 struct { F0 int; F1 bool }
func genFn0() int32 {
    var v7 bool = true
    if v7 {
        v7 = (!v7)
    } else {
        var v8 uint64 = 0
        var v9 bool = v7
        _ = v8
        _ = v9
    }
    _ = v7
    return -1
}
func genFn1() uint {
    switch -842 {
    case -1:
        var v10 int16 = -32768
        _ = v10
    case 128:
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := []int{}
            var v13 bool = (2147483647 <= 1)
            _ = fi11
            _ = v12
            _ = v13
        }
    case -761:
        v14 := GenS1{F0: -84, F1: 0}
        _ = v14
    default:
        var v15 int8 = 0
        _ = v15
    }
    if (true || ((-56 != 42) && (-2.25 < (0.1 - 3.14159)))) {
        v16 := []int{-2147483648, -999}
        if ((!(!false)) && (42 < 18446744073709551615)) {
            var v17 int32 = 259
            _ = v17
        } else {
            v18 := GenS1{F0: -1, F1: 905}
            var v19 uint32 = 953
            var v20 uint32 = v19
            _ = v18
            _ = v19
            _ = v20
        }
        var v21 int = -2147483648
        _ = v16
        _ = v21
    } else {
        var v22 int32 = 42
        v22 -= (v22 & genFn0())
        _ = v22
    }
    v23 := GenS1{F0: 42, F1: 895}
    _ = v23
    return 858
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]string)
    v3 := make(map[uint32]string)
    var v4 int16 = -908
    var v5 uint8 = 42
    v6 := make(map[int]uint32)
    switch (((("" + "") + "foo bar") + ("\n\t" + "")) + "foo bar") {
    case "\n\t":
        for fi24 := 0; fi24 < 9; fi24++ {
            fmt.Println(fi24)
            fmt.Println(v4, fi24)
            _ = fi24
        }
    case "foo bar":
        var v25 string = "foo bar"
        _ = v25
    default:
        if true {
            for fi26 := 0; fi26 < 19; fi26++ {
                fmt.Println(fi26, fi26)
                for fi27 := 0; fi27 < 10; fi27++ {
                    var v28 float64 = ((1.5 + (-2.25 + float64(1.5))) * 1.5)
                    v1 += v1
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
            v4 = (-(v4))
            _ = len(v3)
        } else {
            for fi29 := 0; fi29 < 3; fi29++ {
                if (42 <= -27) {
                    switch ("a" + ("\n\t" + "hello")) {
                    case "foo bar":
                        if (2 != 2147483648) {
                            var v30 float32 = 0.1
                            if (v5 == (^(v5))) {
                                var v31 uint16 = 0
                                switch "\n\t" {
                                case "hello":
                                    var v32 uint32 = 942
                                    _ = v32
                                default:
                                    _ = len(v6)
                                }
                                _ = v31
                            }
                            v1 += 2
                            _ = v30
                        } else {
                            switch (-(fi29)) {
                            case -622:
                                if false {
                                    for fi33 := 0; fi33 < 20; fi33++ {
                                        fmt.Println(v4, v1)
                                        fmt.Println(fi29)
                                        _ = fi33
                                    }
                                    v4 -= v4
                                } else {
                                    var v34 int = 2147483647
                                    var v35 float32 = (-2.25 * 3.14159)
                                    switch (v34 | (v34 | fi29)) {
                                    case -482:
                                        v36 := []int{-3, (-(fi29))}
                                        _ = v36
                                    case 720:
                                        v5 = (^(v5))
                                    default:
                                        fmt.Println(v34, v5)
                                    }
                                    _ = v34
                                    _ = v35
                                }
                            case 1:
                                if (65535 < 255) {
                                    if ((fi29 & fi29) <= fi29) {
                                        fmt.Println(v1)
                                    } else {
                                        v37 := GenS1{F0: 42, F1: 705}
                                        v4 -= (v4 % ((v4) | 1))
                                        fmt.Println(v1, v4)
                                        _ = v37
                                    }
                                }
                            case 127:
                                v5 += (v5 * 255)
                            default:
                                v1 += (v1 / ((950) | 1))
                            }
                            var v38 uint8 = (v5 & 210)
                            switch "\n\t" {
                            case "a":
                                var v39 int64 = 0
                                _ = v39
                            default:
                                for fi40 := 0; fi40 < 14; fi40++ {
                                    fmt.Println(v38, v4)
                                    var v41 uint32 = 2147483648
                                    var v42 uint8 = (v5 >> ((genFn1()) & 63))
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                }
                            }
                            _ = v38
                        }
                    case "\n\t":
                        var v43 uint8 = 255
                        _ = v43
                    default:
                        v4 = v4
                    }
                } else {
                    var v44 uint = 2
                    var v45 uint32 = 728
                    _ = v44
                    _ = v45
                }
                v1 = (v1 ^ (v1 + (v1 >> ((2) & 63))))
                fmt.Println(v1, v5, v5)
                _ = fi29
            }
            for fi46 := 0; fi46 < 5; fi46++ {
                v47 := GenS2{F0: 42, F1: "\n\t", F2: 128}
                for fi48 := 0; fi48 < 3; fi48++ {
                    var v49 uint = 884
                    var v50 int8 = 0
                    fmt.Println(v49, v5, v5)
                    v51 := GenS1{F0: -24, F1: (^(v49))}
                    _ = fi48
                    _ = v49
                    _ = v50
                    _ = v51
                }
                var v52 uint8 = 213
                _ = fi46
                _ = v47
                _ = v52
            }
            fmt.Println(v1)
        }
    }
    delete(v6, -981)
    var v53 float64 = (0.0 + -2.25)
    v5 += (v5 + (v5 % (((^(v5))) | 1)))
    fmt.Println(v1, v53)
    v5 += v5
    fmt.Println(v1, v53, v5)
    for fi54 := 0; fi54 < 7; fi54++ {
        var v55 uint = 406
        switch 19 {
        case 9223372036854775807:
            var v56 int32 = -598
            _ = v56
        case 832:
            switch ("hello" + "hello") {
            case "foo bar":
                v57 := GenS2{F0: v5, F1: "a", F2: v5}
                _ = v57
            case "hello":
                v4 += int16(v55)
            case "\n\t":
                if false {
                    v53 += 1e10
                    var v58 float64 = (v53 * v53)
                    _ = v58
                }
            default:
                v55 = (^(v55))
            }
        case 32768:
            if true {
                v59 := GenS1{F0: 81, F1: genFn1()}
                if false {
                    if true {
                        var v60 uint = v55
                        var v61 uint32 = 2147483648
                        _ = v60
                        _ = v61
                    }
                } else {
                    for fi62 := 0; fi62 < 8; fi62++ {
                        var v63 uint32 = 2147483647
                        fmt.Println(v5, fi62)
                        _ = fi62
                        _ = v63
                    }
                    for fi64 := 0; fi64 < 18; fi64++ {
                        v65 := GenS3{F0: (v55 % ((v55) | 1)), F1: 32}
                        v5 += 85
                        var v66 uint64 = (^(v1))
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                    var v67 string = "foo bar"
                    _ = v67
                }
                _ = v59
            } else {
                v53 = v53
            }
        default:
            v5 = v5
        }
        _ = fi54
        _ = v55
    }
    v53 = v53
    if true {
        var v68 float64 = (float64(-2.25) + ((v53 - (0.1 + (0.1 / v53))) * (v53 + 1.5)))
        v69 := GenS1{F0: -128, F1: 18446744073709551615}
        v53 = v68
        _ = v68
        _ = v69
    }
    fmt.Println(v53, v5)
    for fi70 := 0; fi70 < 8; fi70++ {
        v71 := []int{(fi70 % ((fi70) | 1)), -9223372036854775808, fi70}
        _ = fi70
        _ = v71
    }
    if true {
        v72 := GenS2{F0: 2, F1: ("foo bar" + ""), F2: 2}
        var v73 float64 = v53
        _ = v72
        _ = v73
    }
    var v74 string = ("hello" + "\n\t")
    var v75 int = 423
    var v76 uint = 644
    if true {
        if (false && false) {
            if false {
                if false {
                    v77 := v6[-2147483648]
                    switch v75 {
                    case 127:
                        if false {
                            var v78 int16 = v4
                            _ = v78
                        }
                    case 583:
                        if ((v5 + v5) >= 255) {
                            v4 += v4
                            for fi79 := 0; fi79 < 15; fi79++ {
                                if false {
                                    v1 += (^(v1))
                                    switch -1 {
                                    case 2147483647:
                                        if false {
                                            if (!(v77 > (^(v77)))) {
                                                switch -411 {
                                                case -116:
                                                    if false {
                                                        if ((-9223372036854775808 <= 42) || (0 >= 0)) {
                                                            var v80 uint = v76
                                                            if (true || true) {
                                                                var v81 float32 = -2.25
                                                                for fi82 := 0; fi82 < 9; fi82++ {
                                                                    var v83 int32 = 1
                                                                    delete(v2, fi82)
                                                                    if true {
                                                                        if true {
                                                                            var v84 uint32 = uint32(v80)
                                                                            switch 42 {
                                                                            case -1:
                                                                                var v85 uint8 = 255
                                                                                _ = v85
                                                                            default:
                                                                                if (!(uint64(v5) >= 281)) {
                                                                                    fmt.Println(v5, v76)
                                                                                    v86 := GenS3{F0: (v76 + (v76 & v80)), F1: 43}
                                                                                    if false {
                                                                                        fmt.Println(v4, v84)
                                                                                    } else {
                                                                                        v87 := GenS2{F0: (v5 ^ 255), F1: v74, F2: (v5 / ((v5) | 1))}
                                                                                        delete(v6, (fi82 % (((fi82 & v75)) | 1)))
                                                                                        v88 := []int{-869, -817}
                                                                                        _ = v87
                                                                                        _ = v88
                                                                                    }
                                                                                    _ = v86
                                                                                }
                                                                            }
                                                                            switch ((((("a" + "\n\t") + "") + "") + v74) + ("" + v74)) {
                                                                            case "foo bar":
                                                                                v53 = 1e10
                                                                                var v89 int16 = (v4 / (((^(v4))) | 1))
                                                                                _ = v89
                                                                            case "\n\t":
                                                                                v77 -= uint32(v5)
                                                                            }
                                                                            _ = v84
                                                                        }
                                                                        if (((v53 - 1.5) >= (3.14159 + v53)) && (!(-2147483648 <= (^(v83))))) {
                                                                            switch fi79 {
                                                                            case 408:
                                                                                v75 -= fi79
                                                                            case 9223372036854775807:
                                                                                switch v74 {
                                                                                case "\n\t":
                                                                                    v90 := GenS2{F0: v5, F1: "foo bar", F2: (v5 + 255)}
                                                                                    _ = v90
                                                                                case "":
                                                                                    v91 := v3[(v77 & 176)]
                                                                                    _ = v91
                                                                                case "hello":
                                                                                    v81 += 1e10
                                                                                default:
                                                                                    var v92 string = ((v74 + ((v74 + ((v74 + "foo bar") + v74)) + (v74 + (v74 + v74)))) + v74)
                                                                                    _ = v92
                                                                                }
                                                                            case -2147483648:
                                                                                if (v77 >= (^(v77))) {
                                                                                    if ((v77 & 2) == v77) {
                                                                                        if true {
                                                                                            v53 = 3.14159
                                                                                            var v93 uint8 = uint8(v77)
                                                                                            _ = v93
                                                                                        } else {
                                                                                            v80 = (v80 ^ (v80 ^ (v76 | 42)))
                                                                                        }
                                                                                        v80 += genFn1()
                                                                                        var v94 uint32 = 42
                                                                                        _ = v94
                                                                                    }
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v95 := []int{(v75 ^ v75)}
                                                                            if len(v95) > 0 {
                                                                                v95[0] = (-(v75))
                                                                            }
                                                                            var v96 uint16 = uint16(v80)
                                                                            _ = v95
                                                                            _ = v96
                                                                        }
                                                                    }
                                                                    _ = fi82
                                                                    _ = v83
                                                                }
                                                                var v97 int = v75
                                                                _ = v81
                                                                _ = v97
                                                            } else {
                                                                if (true || true) {
                                                                    var v98 int = fi79
                                                                    _ = v98
                                                                }
                                                                v99 := GenS1{F0: -76, F1: (v76 + (v76 / ((v76) | 1)))}
                                                                if true {
                                                                    v100 := GenS1{F0: -1, F1: 255}
                                                                    switch fi79 {
                                                                    case 42:
                                                                        fmt.Println(v77)
                                                                    case 128:
                                                                        if ((((v53 / (0.0 + v53)) / 3.14159) + v53) < v53) {
                                                                            for fi101 := 0; fi101 < 7; fi101++ {
                                                                                fmt.Println(v5)
                                                                                _ = fi101
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v100
                                                                }
                                                                _ = v99
                                                            }
                                                            for fi102 := 0; fi102 < 12; fi102++ {
                                                                fmt.Println(fi79)
                                                                v103 := GenS1{F0: 1, F1: 18446744073709551615}
                                                                _ = fi102
                                                                _ = v103
                                                            }
                                                            _ = v80
                                                        }
                                                        for fi104 := 0; fi104 < 20; fi104++ {
                                                            v5 = 42
                                                            switch ((v74 + (v74 + v74)) + "hello") {
                                                            case "a":
                                                                v53 -= (float64(1.5) - v53)
                                                            default:
                                                                if true {
                                                                    v4 += 890
                                                                    if (781 > 42) {
                                                                        v105 := GenS4{F0: (v76 >> ((255) & 63)), F1: -9223372036854775808}
                                                                        var v106 int = (^(fi79))
                                                                        _ = v105
                                                                        _ = v106
                                                                    } else {
                                                                        if true {
                                                                            fmt.Println(v4)
                                                                            v5 = v5
                                                                        }
                                                                    }
                                                                    _ = len(v6)
                                                                } else {
                                                                    if (((!(genFn1() == (v76 | (v76 * (v76 << (((^(v76))) & 63)))))) && true) || true) {
                                                                        if (0 > 3) {
                                                                            var v107 int = fi104
                                                                            v74 = v74
                                                                            v108 := GenS4{F0: (^(v76)), F1: 1}
                                                                            v109 := v3[697]
                                                                            _ = v107
                                                                            _ = v108
                                                                            _ = v109
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            v110 := []int{4294967295}
                                                                            var v111 uint = genFn1()
                                                                            fmt.Println(v77, v76, v77)
                                                                            _ = v110
                                                                            _ = v111
                                                                        }
                                                                        _ = len(v6)
                                                                        v77 = (v77 << ((v76) & 63))
                                                                        var v112 float32 = -2.25
                                                                        var v113 uint32 = v77
                                                                        _ = v112
                                                                        _ = v113
                                                                    }
                                                                    v77 -= (v77 - 2147483648)
                                                                    fmt.Println(fi104, v53, v76)
                                                                }
                                                            }
                                                            _ = fi104
                                                        }
                                                    }
                                                default:
                                                    var v114 uint8 = v5
                                                    _ = v114
                                                }
                                                v3[(v77 >> (((v76 % (((v76 % (((^(v76))) | 1))) | 1))) & 63))] = v74
                                                var v115 string = ("" + (((v74 + v74) + "a") + "foo bar"))
                                                _ = v115
                                            }
                                        }
                                    case 9223372036854775807:
                                        for fi116 := 0; fi116 < 9; fi116++ {
                                            v117 := GenS5{F0: v77, F1: 0, F2: -128}
                                            if ((^(v5)) == v5) {
                                                switch ("a" + v74) {
                                                case "":
                                                    v75 = (-(fi116))
                                                }
                                                var v118 int8 = 42
                                                _ = v118
                                            } else {
                                                var v119 uint = (v76 << ((42) & 63))
                                                _ = v119
                                            }
                                            switch v74 {
                                            case "foo bar":
                                                fmt.Println(v74, fi79)
                                            default:
                                                v120 := GenS3{F0: (^(v76)), F1: 100}
                                                _ = v120
                                            }
                                            switch v75 {
                                            case 128:
                                                if false {
                                                    fmt.Println(fi79, v76)
                                                    v121 := []int{fi116, v75, fi116, (fi79 ^ (-(fi79)))}
                                                    var v122 uint32 = v77
                                                    if false {
                                                        v53 += -2.25
                                                        if (!((42 <= -1) && (-1 <= 87))) {
                                                            if (!(-771 >= -9223372036854775808)) {
                                                                v1 = 0
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    var v123 int32 = int32(v5)
                                                                    _ = v123
                                                                }
                                                            } else {
                                                                v2[fi116] = "a"
                                                            }
                                                        }
                                                    }
                                                    _ = v121
                                                    _ = v122
                                                }
                                            case -9223372036854775808:
                                                fmt.Println(v76, fi79, v4)
                                            case -444:
                                                var v124 bool = (2147483647 >= 42)
                                                _ = v124
                                            default:
                                                var v125 int = fi116
                                                _ = v125
                                            }
                                            _ = fi116
                                            _ = v117
                                        }
                                    case 1:
                                        v126 := []int{fi79, 127}
                                        _ = v126
                                    default:
                                        v127 := GenS6{F0: 1, F1: v1}
                                        _ = v127
                                    }
                                }
                                _ = fi79
                            }
                            for fi128 := 0; fi128 < 9; fi128++ {
                                switch "\n\t" {
                                case "hello":
                                    v77 -= 416
                                case "a":
                                    v4 = (-(v4))
                                default:
                                    fmt.Println(fi128)
                                }
                                for fi129 := 0; fi129 < 4; fi129++ {
                                    var v130 int64 = 1
                                    v131 := []int{(fi129 % (((v75 ^ fi128)) | 1)), fi129}
                                    var v132 int16 = (v4 ^ v4)
                                    _ = fi129
                                    _ = v130
                                    _ = v131
                                    _ = v132
                                }
                                _ = fi128
                            }
                        }
                    case 2147483647:
                        var v133 int = (v75 << ((v76) & 63))
                        _ = v133
                    }
                    for fi134 := 0; fi134 < 5; fi134++ {
                        for fi135 := 0; fi135 < 10; fi135++ {
                            var v136 int64 = -9223372036854775808
                            _ = fi135
                            _ = v136
                        }
                        switch (v74 + v74) {
                        case "":
                            var v137 uint64 = v1
                            _ = v137
                        case "hello":
                            v53 = (v53 / ((float64((((0.0 - -2.25) * 0.0) + (0.1 / 0.1))) / v53) - (float64(0.1) / (0.1 / ((v53 * v53) * v53)))))
                        case "a":
                            v75 = (^(v75))
                        }
                        _ = fi134
                    }
                    _ = v77
                }
                var v138 bool = true
                _ = v138
            } else {
                var v139 int32 = 2147483647
                v140 := GenS7{F0: uint32(v5), F1: 42}
                v141 := GenS8{F0: true, F1: "a"}
                _ = v139
                _ = v140
                _ = v141
            }
        }
        if false {
            v142 := v2[0]
            v143 := GenS9{F0: (-99 == 100), F1: 42}
            _ = v142
            _ = v143
        }
        for fi144 := 0; fi144 < 12; fi144++ {
            switch 9223372036854775807 {
            case 127:
                for fi145 := 0; fi145 < 15; fi145++ {
                    v53 += 0.1
                    v76 = v76
                    _ = fi145
                }
            case 188:
                fmt.Println(fi144, v4, v53)
            case 32768:
                switch -959 {
                case 9223372036854775807:
                    v146 := GenS3{F0: (v76 * v76), F1: -11}
                    _ = v146
                case 128:
                    if ((((^(v4)) > (v4 - v4)) || false) && true) {
                        if true {
                            v1 = 18446744073709551615
                            v147 := []int{}
                            for fi148 := 0; fi148 < 14; fi148++ {
                                fmt.Println(v53, v74, v74)
                                switch "" {
                                case "foo bar":
                                    var v149 uint = (v76 % ((genFn1()) | 1))
                                    _ = v149
                                case "a":
                                    v2[(v75 | 32767)] = (v74 + v74)
                                }
                                v5 -= v5
                                _ = fi148
                            }
                            _ = v147
                        }
                    }
                case 32768:
                    var v150 float32 = (0.0 / 1.0)
                    _ = v150
                default:
                    v151 := GenS10{F0: (fi144 & (v75 / ((v75) | 1))), F1: (1 >= 42)}
                    _ = v151
                }
            }
            for fi152 := 0; fi152 < 6; fi152++ {
                var v153 float64 = 0.1
                v1 = (^(v1))
                _ = fi152
                _ = v153
            }
            _ = fi144
        }
    }
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v74
    _ = v75
    _ = v76
}
