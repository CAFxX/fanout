package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint64; F1 int; F2 uint64 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS6 struct { F0 uint16; F1 bool; F2 int16 }
type GenS7 struct { F0 uint32; F1 uint16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int16; F1 uint; F2 int8 }
type GenS11 struct { F0 int32; F1 uint16 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int64; F1 int8; F2 int8 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint32; F1 uint; F2 int32 }
type GenS17 struct { F0 uint }
func genFn0(p0 int, p1 uint64) uint8 {
    v4 := []int{p0, -636}
    var v5 int8 = int8(p1)
    var v6 int16 = 42
    p0 = p0
    switch "\n\t" {
    case "a":
        if true {
            p1 = p1
            v6 = v6
            v7 := GenS1{F0: int32(p1), F1: 35}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 uint64 = p1
                _ = fi8
                _ = v9
            }
            switch ("a" + "hello") {
            case "":
                v5 = v5
            case "\n\t":
                p1 = (p1 + 213)
            default:
                p1 = (p1 | p1)
            }
            p0 = (p0 << ((42) & 63))
        }
    case "":
        switch "foo bar" {
        case "\n\t":
            var v10 uint8 = 64
            _ = v10
        }
    case "foo bar":
        fmt.Println(p1, v6, p1)
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 127
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 0.0
    var v3 int = -1
    if false {
        for fi11 := 0; fi11 < 4; fi11++ {
            if (149 != 109) {
                if false {
                    for fi12 := 0; fi12 < 13; fi12++ {
                        var v13 int32 = 1
                        _ = fi12
                        _ = v13
                    }
                    v3 += 127
                }
                switch fi11 {
                case -399:
                    var v14 uint64 = 1
                    _ = v14
                default:
                    if true {
                        var v15 string = "hello"
                        for fi16 := 0; fi16 < 2; fi16++ {
                            switch (fi16 & (^(fi11))) {
                            case 32767:
                                v17 := v1[("a" + ((v15 + "") + v15))]
                                _ = v17
                            }
                            v2 -= ((v2 * ((v2 / (1.5 * (0.1 / v2))) - (v2 * ((1e10 / float64(float32(1.5))) + 0.1)))) - v2)
                            _ = fi16
                        }
                        _ = v15
                    } else {
                        v18 := GenS2{F0: 9223372036854775807}
                        _ = v18
                    }
                }
                var v19 int8 = 1
                _ = v19
            } else {
                v20 := GenS2{F0: -883}
                v3 = -873
                _ = v20
            }
            for fi21 := 0; fi21 < 7; fi21++ {
                if false {
                    v22 := GenS3{F0: "", F1: 466, F2: 2}
                    v23 := []int{fi11, v3, (-(fi21))}
                    var v24 uint32 = 4294967295
                    _ = v22
                    _ = v23
                    _ = v24
                }
                fmt.Println(v2)
                if (74 >= 1) {
                    v25 := GenS1{F0: -747, F1: 1}
                    _ = v25
                } else {
                    if false {
                        v26 := GenS4{F0: 2, F1: (v3 / ((140) | 1)), F2: 18446744073709551615}
                        _ = v26
                    }
                    var v27 uint16 = 42
                    _ = v27
                }
                _ = fi21
            }
            for fi28 := 0; fi28 < 20; fi28++ {
                if false {
                    var v29 uint64 = 341
                    _ = v29
                } else {
                    if false {
                        if ((138 == 128) && true) {
                            if true {
                                v2 += 0.1
                            } else {
                                v2 -= (v2 / v2)
                                fmt.Println(fi11, v2)
                                if false {
                                    var v30 uint = 42
                                    var v31 int16 = 0
                                    _ = v30
                                    _ = v31
                                } else {
                                    v3 -= (fi28 + fi28)
                                    switch ("foo bar" + "\n\t") {
                                    case "a":
                                        v32 := GenS5{F0: 2, F1: 127, F2: 317}
                                        _ = v32
                                    case "foo bar":
                                        v33 := GenS6{F0: 84, F1: false, F2: 42}
                                        _ = v33
                                    default:
                                        fmt.Println(v2)
                                    }
                                }
                            }
                            v34 := GenS5{F0: 127, F1: genFn0(-707, 1), F2: 4294967295}
                            v35 := GenS4{F0: 543, F1: fi28, F2: 18446744073709551615}
                            _ = v34
                            _ = v35
                        }
                        v36 := GenS4{F0: 42, F1: 127, F2: 1}
                        var v37 int32 = -1
                        _ = v36
                        _ = v37
                    }
                }
                delete(v1, (("a" + "\n\t") + "a"))
                switch v3 {
                case -996:
                    v2 += (((v2 / 0.1) + (v2 / 1e10)) - v2)
                case 296:
                    v38 := []int{}
                    _ = v38
                case 1:
                    v39 := GenS2{F0: 1}
                    _ = v39
                }
                v3 -= (-(fi28))
                _ = fi28
            }
            _ = fi11
        }
        fmt.Println(v2)
    }
    var v40 uint64 = 2
    if (0 >= 1) {
        v41 := GenS3{F0: "", F1: 814, F2: 4294967295}
        if false {
            for fi42 := 0; fi42 < 15; fi42++ {
                v40 = (v40 % (((v40 + 0)) | 1))
                v43 := GenS7{F0: 768, F1: 42}
                _ = fi42
                _ = v43
            }
            if (535 >= 0) {
                var v44 uint8 = 42
                fmt.Println(v44)
                _ = v44
            }
            for fi45 := 0; fi45 < 14; fi45++ {
                fmt.Println(v2, v40, v2)
                _ = fi45
            }
        } else {
            v40 += v40
            v46 := GenS7{F0: 1, F1: 2}
            _ = v46
        }
        for fi47 := 0; fi47 < 20; fi47++ {
            var v48 int = (fi47 * (v3 | (v3 ^ fi47)))
            delete(v1, ((((("foo bar" + ((("" + "a") + "\n\t") + "foo bar")) + "hello") + "") + "foo bar") + "hello"))
            for fi49 := 0; fi49 < 17; fi49++ {
                var v50 float64 = v2
                _ = fi49
                _ = v50
            }
            _ = fi47
            _ = v48
        }
        _ = v41
    }
    switch "foo bar" {
    case "foo bar":
        if (!(235 < 4294967295)) {
            v51 := GenS8{F0: 967}
            if false {
                switch "hello" {
                case "hello":
                    _ = len(v1)
                default:
                    var v52 int32 = -233
                    _ = v52
                }
                var v53 uint64 = v40
                _ = v53
            } else {
                v54 := GenS5{F0: 223, F1: 0, F2: 345}
                _ = v54
            }
            v3 += (v3 / (((v3 + v3)) | 1))
            _ = v51
        }
    case "hello":
        if true {
            switch "a" {
            case "":
                var v55 int16 = 42
                _ = v55
            }
            v56 := GenS9{F0: -1}
            _ = v56
        } else {
            fmt.Println(v2, v2, v40)
        }
    case "a":
        if (!("" != "foo bar")) {
            delete(v1, ("hello" + "hello"))
            v57 := GenS9{F0: 22}
            _ = v57
        } else {
            switch v3 {
            case 0:
                switch v3 {
                case 0:
                    fmt.Println(v40, v40)
                case 9223372036854775807:
                    v3 = (v3 & 128)
                case 42:
                    var v58 float32 = (0.0 / (1.5 / (((0.1 + (1e10 - 1.5)) / float32(0.1)) / ((float32(v2) * 3.14159) - 3.14159))))
                    _ = v58
                default:
                    v59 := GenS10{F0: 1000, F1: 104, F2: -63}
                    _ = v59
                }
            case -162:
                switch (v3 & v3) {
                case -9223372036854775808:
                    v40 = (v40 - 2)
                case -2147483648:
                    if (!false) {
                        var v60 float64 = v2
                        _ = v60
                    }
                }
            default:
                v61 := []int{32767, (v3 | (-(v3))), v3}
                _ = v61
            }
            for fi62 := 0; fi62 < 18; fi62++ {
                v63 := v1["foo bar"]
                _ = fi62
                _ = v63
            }
        }
    default:
        v64 := []int{v3}
        _ = v64
    }
    switch 4294967295 {
    case -9223372036854775808:
        var v65 uint16 = 32767
        _ = v65
    case 127:
        for fi66 := 0; fi66 < 18; fi66++ {
            for fi67 := 0; fi67 < 8; fi67++ {
                v40 += 4294967295
                _ = fi67
            }
            fmt.Println(v40, v3)
            _ = fi66
        }
    case 42:
        switch ("hello" + (("foo bar" + "hello") + ((("foo bar" + "") + "hello") + (("a" + "a") + "foo bar")))) {
        case "":
            var v68 int = v3
            _ = v68
        case "a":
            fmt.Println(v3, v2, v3)
        }
    }
    for fi69 := 0; fi69 < 1; fi69++ {
        v2 = v2
        var v70 int = fi69
        _ = fi69
        _ = v70
    }
    v40 += v40
    for fi71 := 0; fi71 < 12; fi71++ {
        switch fi71 {
        case -868:
            if true {
                v40 = 18446744073709551615
            } else {
                for fi72 := 0; fi72 < 15; fi72++ {
                    v3 -= (fi72 ^ (^(fi72)))
                    v3 = fi71
                    v73 := GenS11{F0: 42, F1: 91}
                    _ = fi72
                    _ = v73
                }
                v74 := GenS12{F0: (v3 & -1)}
                if (757 > 0) {
                    v75 := GenS13{F0: 42, F1: -6, F2: 100}
                    if true {
                        var v76 int32 = 42
                        for fi77 := 0; fi77 < 15; fi77++ {
                            switch "\n\t" {
                            case "hello":
                                switch (v3 << ((637) & 63)) {
                                case 32768:
                                    v3 -= v3
                                case 0:
                                    var v78 bool = true
                                    _ = v78
                                default:
                                    fmt.Println(fi71)
                                }
                            case "foo bar":
                                v79 := GenS7{F0: 958, F1: 2}
                                _ = v79
                            case "a":
                                v80 := GenS1{F0: v76, F1: v40}
                                _ = v80
                            default:
                                v81 := GenS4{F0: v40, F1: fi71, F2: 18446744073709551615}
                                _ = v81
                            }
                            var v82 string = "hello"
                            v3 = fi71
                            v2 = 1.5
                            v2 -= v2
                            if true {
                                v76 = v76
                                v76 = v76
                                v83 := []int{(fi71 / ((fi71) | 1)), (-(fi71))}
                                v84 := GenS3{F0: ("\n\t" + v82), F1: 2, F2: 0}
                                fmt.Println(fi77, v2, v3)
                                _ = v83
                                _ = v84
                            }
                            if true {
                                v3 -= v3
                                if false {
                                    switch 1 {
                                    case 913:
                                        if (-698 > -409) {
                                            delete(v1, ((v82 + (("a" + ((v82 + "foo bar") + "hello")) + v82)) + ("hello" + v82)))
                                            v40 -= (v40 - (v40 << ((42) & 63)))
                                            var v85 uint64 = (^(v40))
                                            _ = v85
                                        } else {
                                            v3 += -1
                                            var v86 float64 = v2
                                            v87 := GenS5{F0: 1, F1: 2, F2: 582}
                                            _ = v86
                                            _ = v87
                                        }
                                    case -2147483648:
                                        v3 -= (^(fi77))
                                    case 807:
                                        _ = len(v1)
                                    default:
                                        if true {
                                            if ((1 >= 1) && false) {
                                                if true {
                                                    var v88 uint16 = 2
                                                    _ = v88
                                                }
                                                v89 := GenS6{F0: 42, F1: true, F2: -32768}
                                                v76 = (v76 - (v76 / ((-667) | 1)))
                                                _ = v89
                                            }
                                            var v90 float64 = (float64((1e10 * (((-2.25 / 1e10) * float32(v2)) + ((3.14159 * 0.1) * (1.5 / (0.1 + 1.5)))))) + v2)
                                            var v91 float32 = float32(-2.25)
                                            v76 -= v76
                                            var v92 float32 = ((v91 / v91) - 0.1)
                                            _ = v90
                                            _ = v91
                                            _ = v92
                                        }
                                    }
                                    v2 += float64(0.0)
                                } else {
                                    var v93 float64 = 1e10
                                    if true {
                                        var v94 string = (v82 + (v82 + ""))
                                        if (324 != 2147483647) {
                                            switch ((v82 + v82) + ((("" + v94) + ("a" + "")) + v82)) {
                                            case "hello":
                                                var v95 uint8 = 42
                                                v93 += v2
                                                fmt.Println(v95)
                                                _ = v95
                                            case "a":
                                                switch v82 {
                                                case "":
                                                    var v96 int = (fi71 - fi77)
                                                    _ = v96
                                                default:
                                                    switch -2147483648 {
                                                    case -2147483648:
                                                        var v97 int64 = 0
                                                        _ = v97
                                                    case 810:
                                                        if false {
                                                            v98 := []int{fi71, v3}
                                                            v2 = (v2 - float64(((-2.25 * (((1e10 + float32(3.14159)) - (((1e10 * 1.5) - (1.5 / -2.25)) + 0.0)) / 3.14159)) / 1e10)))
                                                            delete(v1, (v94 + v82))
                                                            var v99 string = v82
                                                            _ = v98
                                                            _ = v99
                                                        } else {
                                                            if true {
                                                                if false {
                                                                    v40 = 1
                                                                    var v100 int64 = 42
                                                                    if (255 == uint8(v40)) {
                                                                        var v101 int32 = 0
                                                                        v102 := GenS10{F0: 881, F1: 255, F2: -1}
                                                                        v94 = (((v82 + v94) + ("\n\t" + v94)) + v94)
                                                                        _ = v101
                                                                        _ = v102
                                                                    }
                                                                    _ = v100
                                                                }
                                                                var v103 float32 = (3.14159 - 0.1)
                                                                var v104 int64 = -626
                                                                _ = v103
                                                                _ = v104
                                                            } else {
                                                                v105 := v1[((v82 + v94) + v82)]
                                                                v3 += (^(v3))
                                                                var v106 uint8 = 125
                                                                _ = v105
                                                                _ = v106
                                                            }
                                                            var v107 int16 = 0
                                                            var v108 int64 = 9223372036854775807
                                                            if true {
                                                                if ((v107 % (((v107 | 481)) | 1)) != 42) {
                                                                    switch "hello" {
                                                                    case "a":
                                                                        v109 := GenS2{F0: v108}
                                                                        _ = v109
                                                                    default:
                                                                        v2 = v93
                                                                    }
                                                                    if ((v93 / (((v93 - (v93 - -2.25)) + (v2 / (v2 + float64(0.1)))) / v93)) >= float64(-2.25)) {
                                                                        var v110 float64 = v2
                                                                        _ = v110
                                                                    } else {
                                                                        switch 42 {
                                                                        case -171:
                                                                            switch 9223372036854775807 {
                                                                            case 9223372036854775807:
                                                                                switch ("a" + v94) {
                                                                                case "foo bar":
                                                                                    v111 := GenS6{F0: 32767, F1: true, F2: v107}
                                                                                    _ = v111
                                                                                case "":
                                                                                    v112 := GenS2{F0: (v108 & (-(v108)))}
                                                                                    var v113 int16 = v107
                                                                                    v114 := GenS12{F0: fi77}
                                                                                    _ = v112
                                                                                    _ = v113
                                                                                    _ = v114
                                                                                default:
                                                                                    var v115 uint = 0
                                                                                    _ = v115
                                                                                }
                                                                            case 1:
                                                                                fmt.Println(v108, v82, v82)
                                                                            case 0:
                                                                                v116 := GenS14{F0: (v3 & -958)}
                                                                                _ = v116
                                                                            }
                                                                        case -284:
                                                                            v3 = (-(fi71))
                                                                        default:
                                                                            v117 := GenS7{F0: 802, F1: 65535}
                                                                            _ = v117
                                                                        }
                                                                        v94 = ("" + ("a" + (v82 + (("hello" + (((("foo bar" + "hello") + "hello") + ((v94 + "foo bar") + (v82 + "foo bar"))) + v94)) + v94))))
                                                                    }
                                                                    v94 = (v82 + v82)
                                                                    if true {
                                                                        if true {
                                                                            var v118 float64 = (3.14159 * v93)
                                                                            _ = v118
                                                                        } else {
                                                                            switch fi77 {
                                                                            case -281:
                                                                                v2 = 0.0
                                                                            case -1:
                                                                                _ = len(v1)
                                                                            }
                                                                        }
                                                                        v119 := GenS15{F0: v108}
                                                                        v120 := []int{}
                                                                        _ = v119
                                                                        _ = v120
                                                                    } else {
                                                                        v121 := GenS3{F0: ("foo bar" + (v82 + v94)), F1: 510, F2: 2}
                                                                        var v122 uint = 79
                                                                        _ = v121
                                                                        _ = v122
                                                                    }
                                                                    var v123 int = (-(v3))
                                                                    _ = v123
                                                                } else {
                                                                    v124 := GenS5{F0: 163, F1: 1, F2: 42}
                                                                    var v125 string = (v82 + v94)
                                                                    _ = v124
                                                                    _ = v125
                                                                }
                                                            } else {
                                                                var v126 uint32 = 2
                                                                _ = v126
                                                            }
                                                            _ = v107
                                                            _ = v108
                                                        }
                                                        delete(v1, ((("a" + v82) + v82) + (v82 + (v82 + "foo bar"))))
                                                    case 0:
                                                        fmt.Println(v93, fi77, v82)
                                                    default:
                                                        var v127 int8 = -1
                                                        _ = v127
                                                    }
                                                }
                                            default:
                                                var v128 float64 = v93
                                                _ = v128
                                            }
                                            var v129 int16 = 0
                                            _ = v129
                                        } else {
                                            v130 := []int{(v3 >> ((1) & 63)), (-(fi71)), fi77}
                                            _ = v130
                                        }
                                        _ = v94
                                    } else {
                                        if false {
                                            v1[v82] = (fi77 << ((504) & 63))
                                            v131 := []int{311}
                                            _ = v131
                                        } else {
                                            delete(v1, "a")
                                            var v132 float32 = 1e10
                                            var v133 int16 = -1
                                            _ = v132
                                            _ = v133
                                        }
                                        switch 214 {
                                        case 9223372036854775807:
                                            switch v82 {
                                            case "a":
                                                var v134 int32 = -924
                                                _ = v134
                                            default:
                                                fmt.Println(v40)
                                            }
                                        }
                                        v135 := GenS15{F0: 291}
                                        _ = v135
                                    }
                                    _ = v93
                                }
                            } else {
                                var v136 bool = true
                                var v137 float32 = 3.14159
                                _ = v136
                                _ = v137
                            }
                            v138 := []int{(fi77 & (v3 * -1))}
                            _ = fi77
                            _ = v138
                            _ = v82
                        }
                        v76 -= (v76 + (-(v76)))
                        _ = v76
                    } else {
                        var v139 uint32 = 87
                        if true {
                            if false {
                                fmt.Println(v40, v40)
                                if false {
                                    for fi140 := 0; fi140 < 11; fi140++ {
                                        if ("hello" != "hello") {
                                            switch 32767 {
                                            case -9223372036854775808:
                                                v141 := GenS16{F0: v139, F1: 1, F2: 0}
                                                _ = v141
                                            default:
                                                if (!(v139 > v139)) {
                                                    v40 -= v40
                                                    v1[""] = (v3 | (v3 << ((uint(v139)) & 63)))
                                                }
                                                var v142 float64 = v2
                                                v143 := []int{0, fi71, fi140, fi140}
                                                var v144 uint32 = 2147483647
                                                var v145 uint = 4294967295
                                                v146 := v1["hello"]
                                                switch "foo bar" {
                                                case "a":
                                                    switch (fi71 * (-(fi140))) {
                                                    case 1:
                                                        v147 := GenS17{F0: 18446744073709551615}
                                                        _ = v147
                                                    case 4294967295:
                                                        _ = len(v1)
                                                    case -2147483648:
                                                        v2 -= ((v2 + v142) - (v2 - v2))
                                                    default:
                                                        if ("foo bar" > "\n\t") {
                                                            v148 := GenS12{F0: -9223372036854775808}
                                                            var v149 int64 = 9223372036854775807
                                                            var v150 int64 = (v149 << ((v145) & 63))
                                                            _ = v148
                                                            _ = v149
                                                            _ = v150
                                                        }
                                                    }
                                                case "hello":
                                                    switch "foo bar" {
                                                    case "":
                                                        v1["\n\t"] = fi71
                                                    }
                                                case "":
                                                    if (v40 <= uint64(v145)) {
                                                        if true {
                                                            v151 := GenS2{F0: 0}
                                                            var v152 int = v146
                                                            _ = v151
                                                            _ = v152
                                                        } else {
                                                            v145 = (^(v145))
                                                            if len(v143) >= 2 {
                                                                v143 = v143[1:len(v143)]
                                                            }
                                                            var v153 uint32 = (v144 >> (((v145 + uint(v40))) & 63))
                                                            _ = v153
                                                        }
                                                    }
                                                default:
                                                    v139 += 4294967295
                                                }
                                                _ = v142
                                                _ = v143
                                                _ = v144
                                                _ = v145
                                                _ = v146
                                            }
                                            v40 += (v40 & v40)
                                            v154 := GenS9{F0: -1}
                                            _ = v154
                                        }
                                        var v155 string = ("hello" + "hello")
                                        _ = fi140
                                        _ = v155
                                    }
                                    var v156 int64 = -1
                                    var v157 float64 = v2
                                    _ = v156
                                    _ = v157
                                }
                                for fi158 := 0; fi158 < 16; fi158++ {
                                    v2 = v2
                                    v139 = (^(v139))
                                    _ = fi158
                                }
                            } else {
                                for fi159 := 0; fi159 < 15; fi159++ {
                                    switch "" {
                                    case "":
                                        v139 -= v139
                                    case "hello":
                                        v2 = v2
                                    }
                                    _ = fi159
                                }
                                var v160 float32 = (((-2.25 * 3.14159) + 3.14159) - 0.0)
                                _ = v160
                            }
                        } else {
                            for fi161 := 0; fi161 < 11; fi161++ {
                                if (480 == 441) {
                                    var v162 int8 = 100
                                    v163 := GenS6{F0: 65535, F1: ((("hello" + "foo bar") + "\n\t") != ("" + ("a" + ("" + "a")))), F2: 42}
                                    _ = v162
                                    _ = v163
                                } else {
                                    _ = len(v1)
                                    fmt.Println(fi71, v2)
                                }
                                var v164 float64 = -2.25
                                _ = fi161
                                _ = v164
                            }
                            var v165 uint16 = 727
                            switch -9223372036854775808 {
                            case 127:
                                var v166 int16 = 42
                                _ = v166
                            case 323:
                                if true {
                                    v1["\n\t"] = (v3 % ((-2147483648) | 1))
                                    v167 := []int{fi71, fi71}
                                    if false {
                                        v3 = fi71
                                        var v168 uint16 = (v165 / ((v165) | 1))
                                        _ = v168
                                    } else {
                                        var v169 float64 = ((float64((0.0 + 3.14159)) / 3.14159) / ((v2 + v2) + v2))
                                        switch (-(v3)) {
                                        case 1:
                                            v170 := GenS7{F0: (v139 / ((2147483647) | 1)), F1: 1}
                                            _ = v170
                                        }
                                        _ = v169
                                    }
                                    _ = v167
                                } else {
                                    v171 := []int{fi71, 127, -532}
                                    for fi172 := 0; fi172 < 5; fi172++ {
                                        var v173 uint32 = (v139 + (v139 & (v139 & v139)))
                                        if false {
                                            fmt.Println(v40)
                                            _ = len(v1)
                                            switch ("\n\t" + "a") {
                                            case "":
                                                v165 = v165
                                            case "hello":
                                                _ = len(v171)
                                            case "foo bar":
                                                delete(v1, ("a" + "\n\t"))
                                            default:
                                                v3 -= (-(fi172))
                                            }
                                        } else {
                                            v3 = 4294967295
                                        }
                                        _ = fi172
                                        _ = v173
                                    }
                                    _ = v171
                                }
                            case 22:
                                for fi174 := 0; fi174 < 15; fi174++ {
                                    var v175 uint8 = 0
                                    _ = fi174
                                    _ = v175
                                }
                            }
                            _ = v165
                        }
                        _ = v139
                    }
                    v176 := []int{(v3 & 0), v3, (-(fi71)), -731}
                    _ = v176
                    _ = v75
                } else {
                    switch "\n\t" {
                    case "":
                        v177 := []int{fi71, v3, v3, 0}
                        _ = v177
                    case "foo bar":
                        for fi178 := 0; fi178 < 14; fi178++ {
                            if false {
                                v2 = float64(3.14159)
                                var v179 uint8 = 128
                                if false {
                                    v180 := []int{fi71, fi71}
                                    v3 = fi71
                                    switch (-(fi71)) {
                                    case 127:
                                        var v181 float32 = (0.1 + (1.5 * (0.1 - -2.25)))
                                        _ = v181
                                    case -1:
                                        v3 = fi71
                                    }
                                    _ = v180
                                } else {
                                    fmt.Println(fi178, v40, fi71)
                                    v40 = (v40 >> ((2) & 63))
                                    v1[""] = fi71
                                }
                                _ = v179
                            }
                            var v182 string = "hello"
                            var v183 string = v182
                            _ = fi178
                            _ = v182
                            _ = v183
                        }
                    case "hello":
                        v2 = float64((1e10 + (float32((float64((-2.25 * float32((v2 * (((v2 - 1.5) - (v2 / v2)) / (v2 + (((v2 * v2) + v2) / (0.1 + -2.25)))))))) + 3.14159)) * 1e10)))
                    default:
                        var v184 int64 = -118
                        _ = v184
                    }
                    fmt.Println(v40, fi71, fi71)
                }
                _ = v74
            }
        case 1:
            v185 := []int{-1, v3}
            _ = v185
        case 0:
            v2 -= -2.25
        default:
            v186 := GenS6{F0: 32767, F1: false, F2: 32767}
            _ = v186
        }
        var v187 int16 = 0
        _ = fi71
        _ = v187
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
}
