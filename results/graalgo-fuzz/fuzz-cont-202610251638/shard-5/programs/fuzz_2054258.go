package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 bool; F1 string; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint; F1 string; F2 int }
type GenS7 struct { F0 string; F1 int16; F2 int16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint32; F1 int }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int8; F1 int; F2 uint32 }
type GenS12 struct { F0 int8; F1 int32; F2 uint8 }
type GenS13 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS14 struct { F0 uint16; F1 string; F2 string }
func genFn0() uint {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint32 = 2147483648
        v9 := []int{fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint = 66
            _ = fi11
            _ = v12
        }
        switch "a" {
        case "a":
            switch "" {
            case "\n\t":
                var v13 int64 = 9223372036854775807
                _ = v13
            }
        case "\n\t":
            var v14 uint32 = 429
            _ = v14
        default:
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 int32 = -678
                if (0 <= 55) {
                    v17 := GenS1{F0: -32768, F1: -1}
                    _ = v17
                } else {
                    v16 = 1
                    var v18 int16 = 525
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
        }
        _ = fi10
    }
    if true {
        v19 := GenS1{F0: 293, F1: 495}
        for fi20 := 0; fi20 < 5; fi20++ {
            v21 := []int{(fi20 >> ((42) & 63))}
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    return 2
}

func main() {
    var v1 uint64 = 0
    var v2 uint16 = 314
    v3 := make(map[int]uint32)
    var v4 uint8 = 255
    var v5 float64 = (1e10 * 1.5)
    var v6 bool = (2147483647 != 1)
    v22 := GenS1{F0: 32767, F1: 2147483647}
    v23 := GenS1{F0: 1, F1: 128}
    if (!(!(1.5 <= float64(0.1)))) {
        if v6 {
            v24 := []int{922, 32768}
            var v25 uint64 = (v1 ^ (v1 | (v1 << ((65535) & 63))))
            _ = v24
            _ = v25
        } else {
            for fi26 := 0; fi26 < 15; fi26++ {
                v6 = v6
                for fi27 := 0; fi27 < 20; fi27++ {
                    var v28 int = (fi26 / (((fi27 - fi27)) | 1))
                    _ = fi27
                    _ = v28
                }
                for fi29 := 0; fi29 < 19; fi29++ {
                    v1 = (v1 >> ((genFn0()) & 63))
                    var v30 float64 = 3.14159
                    if (v5 != (((v30 * v30) / (1.5 + v5)) - v5)) {
                        var v31 string = "a"
                        _ = v31
                    } else {
                        if v6 {
                            if true {
                                delete(v3, (fi26 & fi29))
                            } else {
                                var v32 int8 = -1
                                v2 -= (^(v2))
                                if v6 {
                                    if (!(4294967295 != 1)) {
                                        fmt.Println(v30, v6, v6)
                                        fmt.Println(v4, v1, v32)
                                    } else {
                                        v33 := []int{4294967295, (fi29 | 42), -696, (fi26 / ((fi29) | 1))}
                                        _ = v33
                                    }
                                } else {
                                    v4 += (v4 - 0)
                                    if (255 > 255) {
                                        v1 -= v1
                                    }
                                }
                                _ = v32
                            }
                            if (32767 >= 834) {
                                v1 = v1
                                switch fi29 {
                                case -2147483648:
                                    v2 += (v2 << ((65535) & 63))
                                }
                            }
                        }
                        var v34 string = "foo bar"
                        var v35 uint64 = v1
                        _ = v34
                        _ = v35
                    }
                    _ = fi29
                    _ = v30
                }
                _ = fi26
            }
            for fi36 := 0; fi36 < 14; fi36++ {
                v3[(-(fi36))] = 2
                _ = fi36
            }
        }
        if false {
            v37 := GenS2{F0: v6, F1: (("" + "") + "foo bar"), F2: 255}
            _ = v37
        } else {
            if (-184 == -116) {
                if (0.1 == -2.25) {
                    v38 := GenS1{F0: -1, F1: 626}
                    _ = v38
                } else {
                    fmt.Println(v5, v2)
                    v4 = v4
                }
                v3[0] = 195
                if false {
                    if v6 {
                        switch "hello" {
                        case "a":
                            switch "a" {
                            case "\n\t":
                                v39 := GenS2{F0: (!v6), F1: ("foo bar" + ""), F2: 942}
                                _ = v39
                            }
                        case "\n\t":
                            for fi40 := 0; fi40 < 1; fi40++ {
                                v3[fi40] = 4294967295
                                _ = fi40
                            }
                        case "foo bar":
                            fmt.Println(v1, v4, v5)
                        default:
                            for fi41 := 0; fi41 < 5; fi41++ {
                                delete(v3, (fi41 << ((499) & 63)))
                                v4 += (v4 * v4)
                                _ = fi41
                            }
                        }
                        _ = len(v3)
                    }
                }
            } else {
                var v42 bool = false
                v3[32767] = 29
                _ = v42
            }
            for fi43 := 0; fi43 < 2; fi43++ {
                for fi44 := 0; fi44 < 11; fi44++ {
                    if v6 {
                        v45 := []int{}
                        if (v6 && v6) {
                            v4 = v4
                            switch ("\n\t" + "a") {
                            case "a":
                                var v46 uint64 = (v1 | (v1 + (v1 >> ((642) & 63))))
                                _ = v46
                            case "hello":
                                var v47 uint64 = v1
                                _ = v47
                            case "foo bar":
                                fmt.Println(v1, v1)
                            }
                        } else {
                            var v48 int16 = 449
                            v48 = (v48 | (^(v48)))
                            if false {
                                var v49 uint64 = (v1 + (^(v1)))
                                _ = v49
                            } else {
                                if v6 {
                                    fmt.Println(v48, fi43, v6)
                                    v50 := GenS1{F0: 492, F1: (^(fi43))}
                                    var v51 int64 = 1
                                    _ = v50
                                    _ = v51
                                } else {
                                    delete(v3, fi43)
                                    var v52 uint = 42
                                    _ = v52
                                }
                                v53 := GenS2{F0: v6, F1: ("a" + ("hello" + (("foo bar" + (("hello" + ("hello" + "foo bar")) + "")) + ("\n\t" + (("" + "hello") + "hello"))))), F2: 781}
                                _ = v53
                            }
                            _ = v48
                        }
                        v6 = v6
                        v3[(fi43 | (fi44 - (fi44 % ((fi44) | 1))))] = 2147483647
                        if v6 {
                            var v54 float64 = float64(float32(v5))
                            switch 1 {
                            case -2147483648:
                                var v55 int16 = -32768
                                _ = v55
                            default:
                                v6 = (42 >= -1)
                            }
                            v56 := GenS3{F0: 235}
                            _ = v54
                            _ = v56
                        }
                        var v57 uint16 = (v2 | (v2 / (((v2 << ((238) & 63))) | 1)))
                        if (!false) {
                            var v58 bool = (!(v2 < v2))
                            v59 := GenS3{F0: uint(v2)}
                            _ = v58
                            _ = v59
                        } else {
                            var v60 int32 = 0
                            if v6 {
                                if (!v6) {
                                    v61 := GenS4{F0: (^(fi43)), F1: 0}
                                    v62 := GenS3{F0: 2}
                                    v5 += v5
                                    fmt.Println(fi43, v5)
                                    v4 = 127
                                    _ = v61
                                    _ = v62
                                }
                                fmt.Println(v57, v6, v60)
                            } else {
                                v3[fi44] = 0
                            }
                            _ = v60
                        }
                        _ = v45
                        _ = v57
                    }
                    _ = fi44
                }
                _ = fi43
            }
        }
        for fi63 := 0; fi63 < 10; fi63++ {
            v64 := GenS4{F0: fi63, F1: 0}
            for fi65 := 0; fi65 < 9; fi65++ {
                fmt.Println(v1, fi63, v1)
                _ = len(v3)
                var v66 bool = ((!v6) && ((!(405 >= 0)) && false))
                _ = fi65
                _ = v66
            }
            if v6 {
                v1 -= v1
                v67 := GenS5{F0: fi63}
                _ = v67
            } else {
                for fi68 := 0; fi68 < 20; fi68++ {
                    v69 := []int{}
                    var v70 int8 = 42
                    var v71 bool = v6
                    _ = fi68
                    _ = v69
                    _ = v70
                    _ = v71
                }
            }
            _ = fi63
            _ = v64
        }
    }
    switch "a" {
    case "\n\t":
        for fi72 := 0; fi72 < 1; fi72++ {
            var v73 int16 = -689
            _ = fi72
            _ = v73
        }
    case "a":
        for fi74 := 0; fi74 < 7; fi74++ {
            v2 += 255
            if (!v6) {
                v2 = 627
                switch (fi74 & (fi74 << ((uint(v4)) & 63))) {
                case 1:
                    var v75 uint32 = 68
                    _ = v75
                case -55:
                    v76 := GenS3{F0: genFn0()}
                    _ = v76
                case 128:
                    v77 := []int{fi74, -513, 1}
                    _ = v77
                default:
                    v5 = v5
                }
                if (((1e10 / 1.0) - v5) == (((0.1 / v5) - 3.14159) + 0.1)) {
                    for fi78 := 0; fi78 < 6; fi78++ {
                        var v79 int = fi78
                        v2 = (v2 - 65535)
                        var v80 bool = v6
                        _ = fi78
                        _ = v79
                        _ = v80
                    }
                    fmt.Println(v1, v6, fi74)
                    switch fi74 {
                    case -1:
                        var v81 int64 = 42
                        _ = v81
                    case 698:
                        if (v1 > v1) {
                            v82 := GenS1{F0: 172, F1: 42}
                            for fi83 := 0; fi83 < 16; fi83++ {
                                v84 := GenS1{F0: 42, F1: fi74}
                                var v85 int8 = -128
                                _ = fi83
                                _ = v84
                                _ = v85
                            }
                            _ = v82
                        }
                    default:
                        v86 := []int{fi74, fi74}
                        _ = v86
                    }
                }
            } else {
                if (-14 >= 0) {
                    v87 := v3[4294967295]
                    _ = v87
                } else {
                    var v88 uint64 = (v1 | v1)
                    v89 := GenS5{F0: 4294967295}
                    var v90 uint = 0
                    _ = v88
                    _ = v89
                    _ = v90
                }
                v4 = (v4 * 128)
                switch ("hello" + ("a" + "foo bar")) {
                case "\n\t":
                    v91 := []int{}
                    _ = v91
                case "hello":
                    if (false && (fi74 == (fi74 ^ (fi74 + (fi74 ^ fi74))))) {
                        switch (fi74 & fi74) {
                        case -545:
                            v92 := v3[(fi74 ^ (fi74 >> ((255) & 63)))]
                            _ = v92
                        case 1:
                            var v93 float64 = v5
                            _ = v93
                        default:
                            if v6 {
                                switch 2147483647 {
                                case 32767:
                                    for fi94 := 0; fi94 < 4; fi94++ {
                                        v95 := GenS6{F0: 255, F1: ((((("a" + ("" + "hello")) + "hello") + "hello") + "\n\t") + ""), F2: (fi74 / ((-917) | 1))}
                                        v96 := GenS7{F0: "\n\t", F1: 32767, F2: 0}
                                        if v6 {
                                            switch (fi94 & (fi94 / ((fi94) | 1))) {
                                            case -40:
                                                v97 := []int{(fi74 * -2147483648), int(v4), 128}
                                                _ = v97
                                            }
                                        }
                                        fmt.Println(v1, v1)
                                        if (-212 == 0) {
                                            if (float64(-2.25) > v5) {
                                                var v98 int8 = -1
                                                v2 = (^(v2))
                                                var v99 uint64 = v1
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v100 := GenS5{F0: fi94}
                                                    _ = v100
                                                case "":
                                                    delete(v3, (fi94 & (fi74 - (fi94 | fi94))))
                                                case "hello":
                                                    if (0.1 >= 1.5) {
                                                        v6 = (v6 && (v5 == (0.1 / v5)))
                                                        v101 := GenS8{F0: uint(v4)}
                                                        var v102 uint = 255
                                                        _ = v101
                                                        _ = v102
                                                    }
                                                default:
                                                    if ((v4 - (v4 + 2)) <= (v4 ^ (^(v4)))) {
                                                        switch 4294967295 {
                                                        case 108:
                                                            if v6 {
                                                                if v6 {
                                                                    v103 := GenS9{F0: 1, F1: fi74}
                                                                    _ = v103
                                                                }
                                                                v104 := v3[fi94]
                                                                if (!false) {
                                                                    v6 = false
                                                                    var v105 string = ("" + "hello")
                                                                    _ = v105
                                                                }
                                                                _ = v104
                                                            }
                                                        default:
                                                            var v106 int8 = v98
                                                            v107 := GenS10{F0: 1000}
                                                            fmt.Println(v5, v6, fi94)
                                                            _ = v106
                                                            _ = v107
                                                        }
                                                    } else {
                                                        v4 += v4
                                                    }
                                                }
                                                var v108 uint = 0
                                                switch (^(fi94)) {
                                                case 0:
                                                    fmt.Println(v6, v5)
                                                case 9223372036854775807:
                                                    v5 = v5
                                                case -9223372036854775808:
                                                    if ((!(!(!(!((float64(3.14159) + v5) == (v5 + v5)))))) && (344 <= 4294967295)) {
                                                        var v109 uint8 = v4
                                                        var v110 int16 = 881
                                                        var v111 int64 = -1
                                                        if ((v110 / (((-(v110))) | 1)) < 1) {
                                                            fmt.Println(fi94)
                                                            var v112 int32 = 1
                                                            _ = v112
                                                        } else {
                                                            switch (-(fi94)) {
                                                            case 9223372036854775807:
                                                                fmt.Println(v2, v108, v1)
                                                            case -9223372036854775808:
                                                                switch (-(fi74)) {
                                                                case 9223372036854775807:
                                                                    var v113 int64 = v111
                                                                    _ = v113
                                                                case 32768:
                                                                    var v114 uint32 = 2147483647
                                                                    _ = v114
                                                                case 0:
                                                                    switch fi94 {
                                                                    case 127:
                                                                        v109 = 42
                                                                    case 2147483647:
                                                                        v115 := []int{}
                                                                        _ = v115
                                                                    case -2147483648:
                                                                        fmt.Println(v2, v5)
                                                                    }
                                                                default:
                                                                    v116 := GenS9{F0: 0, F1: fi94}
                                                                    _ = v116
                                                                }
                                                            case -2147483648:
                                                                var v117 int16 = int16(v4)
                                                                _ = v117
                                                            default:
                                                                if true {
                                                                    switch (fi74 * fi94) {
                                                                    case 2147483647:
                                                                        v118 := []int{-332, fi94}
                                                                        _ = v118
                                                                    case 127:
                                                                        switch int(v2) {
                                                                        case -2147483648:
                                                                            if (!v6) {
                                                                                v119 := []int{(fi74 & fi74), fi74, 127, (-(fi94))}
                                                                                _ = v119
                                                                            }
                                                                        }
                                                                    case 1:
                                                                        v120 := GenS6{F0: (v108 & uint(v4)), F1: "", F2: 42}
                                                                        _ = v120
                                                                    }
                                                                    switch ("a" + "foo bar") {
                                                                    case "foo bar":
                                                                        v121 := GenS11{F0: int8(fi94), F1: int(v4), F2: 2}
                                                                        _ = v121
                                                                    default:
                                                                        if (("hello" < ((((((("a" + ("a" + ("hello" + ("foo bar" + "")))) + "\n\t") + "hello") + (("\n\t" + "a") + ("foo bar" + ("hello" + "a")))) + "") + "\n\t") + "foo bar")) && true) {
                                                                            var v122 int64 = v111
                                                                            v123 := []int{-651, (fi94 / (((fi94 / ((167) | 1))) | 1)), fi94, (fi74 | (fi74 / ((int(v122)) | 1)))}
                                                                            if (!v6) {
                                                                                fmt.Println(v4, v99, v2)
                                                                                if v6 {
                                                                                    if v6 {
                                                                                        if false {
                                                                                            var v124 float64 = float64((1e10 / 0.1))
                                                                                            if len(v123) >= 2 {
                                                                                                v123 = v123[1:len(v123)]
                                                                                            }
                                                                                            _ = v124
                                                                                        } else {
                                                                                            switch "" {
                                                                                            case "hello":
                                                                                                v2 = (v2 / ((v2) | 1))
                                                                                            default:
                                                                                                var v125 int64 = 909
                                                                                                _ = v125
                                                                                            }
                                                                                            fmt.Println(v2)
                                                                                            fmt.Println(v99, fi74, fi94)
                                                                                        }
                                                                                        v126 := v3[fi74]
                                                                                        var v127 bool = (42 < v1)
                                                                                        _ = v126
                                                                                        _ = v127
                                                                                    } else {
                                                                                        if v6 {
                                                                                            var v128 string = ("" + "foo bar")
                                                                                            delete(v3, fi94)
                                                                                            var v129 int = 782
                                                                                            _ = v128
                                                                                            _ = v129
                                                                                        } else {
                                                                                            var v130 int32 = -607
                                                                                            _ = v130
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    var v131 int64 = (v111 / (((v122 << ((v108) & 63))) | 1))
                                                                                    fmt.Println(v99)
                                                                                    _ = len(v123)
                                                                                    _ = v131
                                                                                }
                                                                                v132 := GenS12{F0: v98, F1: -780, F2: 128}
                                                                                _ = v132
                                                                            } else {
                                                                                fmt.Println(v122, v122, v2)
                                                                                var v133 int32 = 231
                                                                                _ = v133
                                                                            }
                                                                            _ = len(v3)
                                                                            _ = v122
                                                                            _ = v123
                                                                        }
                                                                    }
                                                                    if (v110 < (^(v110))) {
                                                                        var v134 uint = (v108 << (((v108 * 0)) & 63))
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            v135 := v3[849]
                                                                            _ = v135
                                                                        }
                                                                        if false {
                                                                            var v136 int64 = 1
                                                                            v134 = genFn0()
                                                                            if (fi74 >= -34) {
                                                                                v137 := GenS4{F0: fi94, F1: (^(v111))}
                                                                                fmt.Println(v2)
                                                                                if ("a" > ("hello" + "a")) {
                                                                                    v138 := GenS13{F0: (v110 | v110), F1: (v2 | v2), F2: v2}
                                                                                    if v6 {
                                                                                        _ = len(v3)
                                                                                        var v139 int8 = v98
                                                                                        v140 := GenS6{F0: (^(v108)), F1: ("hello" + ("hello" + (("hello" + "a") + ("foo bar" + "")))), F2: (fi94 ^ 4294967295)}
                                                                                        var v141 int64 = 0
                                                                                        v6 = (!(-1 >= 0))
                                                                                        fmt.Println(v136, fi74, v136)
                                                                                        _ = v139
                                                                                        _ = v140
                                                                                        _ = v141
                                                                                    }
                                                                                    var v142 bool = (!v6)
                                                                                    _ = v138
                                                                                    _ = v142
                                                                                }
                                                                                _ = v137
                                                                            }
                                                                            v1 -= (v1 & (^(v99)))
                                                                            _ = v136
                                                                        } else {
                                                                            v134 += v108
                                                                            v143 := v3[(fi74 + int(v109))]
                                                                            v144 := GenS14{F0: v2, F1: "a", F2: ("\n\t" + ("a" + "\n\t"))}
                                                                            _ = v143
                                                                            _ = v144
                                                                        }
                                                                        _ = v134
                                                                    }
                                                                } else {
                                                                    v109 += (v109 & (v4 * (v4 ^ (v109 - v4))))
                                                                }
                                                            }
                                                            var v145 uint64 = (v1 + 572)
                                                            v146 := GenS4{F0: fi94, F1: (-(v111))}
                                                            fmt.Println(fi94)
                                                            v109 -= v109
                                                            v108 = (v108 * v108)
                                                            if v6 {
                                                                v1 = (^(v145))
                                                            }
                                                            _ = v145
                                                            _ = v146
                                                        }
                                                        _ = v109
                                                        _ = v110
                                                        _ = v111
                                                    }
                                                }
                                                v147 := v3[(fi94 - (fi74 | fi74))]
                                                _ = v108
                                                _ = v147
                                                _ = v98
                                                _ = v99
                                            } else {
                                                var v148 uint = 255
                                                _ = v148
                                            }
                                            v6 = true
                                            var v149 string = "foo bar"
                                            if true {
                                                v150 := GenS12{F0: 127, F1: -1, F2: v4}
                                                _ = v150
                                            }
                                            fmt.Println(v2, v1, v6)
                                            _ = v149
                                        }
                                        var v151 string = ("foo bar" + "a")
                                        _ = fi94
                                        _ = v151
                                        _ = v95
                                        _ = v96
                                    }
                                case -1:
                                    for fi152 := 0; fi152 < 19; fi152++ {
                                        var v153 int8 = -1
                                        v6 = v6
                                        v153 = 0
                                        _ = fi152
                                        _ = v153
                                    }
                                }
                                v154 := []int{(-(fi74)), int(v1), 103, fi74}
                                _ = len(v3)
                                _ = v154
                            }
                        }
                        if v6 {
                            for fi155 := 0; fi155 < 10; fi155++ {
                                v1 = v1
                                _ = fi155
                            }
                        }
                    }
                default:
                    var v156 uint = 4294967295
                    _ = v156
                }
            }
            _ = fi74
        }
    default:
        var v157 uint64 = v1
        _ = v157
    }
    for fi158 := 0; fi158 < 8; fi158++ {
        v1 = (v1 | (^(v1)))
        if (1 <= -203) {
            var v159 int = (-(fi158))
            _ = v159
        }
        _ = fi158
    }
    for fi160 := 0; fi160 < 12; fi160++ {
        var v161 float32 = (((1e10 / 1.0) - ((((float32((v5 + (float64(1e10) * ((v5 * (v5 * v5)) + v5)))) + (0.1 + (1e10 * 0.0))) - (float32(v5) + (1.5 * ((0.1 * (1e10 + 3.14159)) * 3.14159)))) + float32(3.14159)) + 0.0)) * -2.25)
        _ = fi160
        _ = v161
    }
    delete(v3, 1)
    v162 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v162:", len(v162), v162)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v162
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
