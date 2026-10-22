package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS2 struct { F0 string; F1 int8; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint8; F1 int64; F2 int64 }
type GenS6 struct { F0 int64; F1 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64; F1 uint }
type GenS9 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS10 struct { F0 int16; F1 int }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int32; F1 string }
type GenS13 struct { F0 uint8; F1 int8 }
type GenS14 struct { F0 int64; F1 uint }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 int16; F1 uint64 }
type GenS17 struct { F0 uint8; F1 int8 }
type GenS18 struct { F0 bool; F1 int }
type GenS19 struct { F0 uint; F1 bool; F2 int64 }
type GenS20 struct { F0 uint; F1 int16 }
type GenS21 struct { F0 string; F1 uint }
func genFn0() int32 {
    switch 9223372036854775807 {
    case 32767:
        v6 := GenS1{F0: 634, F1: 54, F2: 255}
        _ = v6
    default:
        if false {
            var v7 uint8 = 42
            v8 := GenS2{F0: ("foo bar" + "a"), F1: -111, F2: 1}
            _ = v7
            _ = v8
        }
    }
    v9 := []int{4294967295, 1, 128}
    if (1 > -2147483648) {
        for fi10 := 0; fi10 < 2; fi10++ {
            if (100 >= 0) {
                var v11 int8 = 126
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                _ = v11
            } else {
                var v12 uint64 = 698
                var v13 int16 = -1
                switch "" {
                case "hello":
                    var v14 int64 = 42
                    _ = v14
                case "":
                    if true {
                        v15 := GenS1{F0: 4294967295, F1: -55, F2: 42}
                        var v16 uint16 = 21
                        var v17 float64 = 0.0
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                case "\n\t":
                    v18 := GenS1{F0: 65535, F1: -22, F2: 91}
                    _ = v18
                default:
                    for fi19 := 0; fi19 < 1; fi19++ {
                        v20 := GenS3{F0: 487}
                        var v21 float64 = -2.25
                        v22 := GenS1{F0: 82, F1: -13, F2: 25}
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                }
                _ = v12
                _ = v13
            }
            _ = fi10
        }
        v23 := GenS1{F0: 2147483648, F1: -128, F2: 127}
        if (1e10 > 0.1) {
            switch "" {
            case "hello":
                if len(v9) > 0 {
                    v9[0] = 32768
                }
            case "a":
                var v24 string = ((((("foo bar" + ("hello" + "foo bar")) + "foo bar") + ("\n\t" + "")) + "") + "\n\t")
                _ = v24
            case "\n\t":
                var v25 uint32 = 4294967295
                _ = v25
            default:
                var v26 int = 128
                _ = v26
            }
            for fi27 := 0; fi27 < 3; fi27++ {
                v9 = append(v9, fi27)
                for fi28 := 0; fi28 < 14; fi28++ {
                    v9 = append(v9, (-(fi28)))
                    if (42 <= 163) {
                        var v29 uint = 770
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        switch (^(fi28)) {
                        case -686:
                            var v30 bool = (false && ((v29 ^ v29) > v29))
                            _ = v30
                        case 127:
                            v31 := GenS4{F0: 42, F1: 2147483647, F2: 1}
                            _ = v31
                        case 541:
                            fmt.Println(v29, fi27)
                        }
                        if false {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            v29 -= 18446744073709551615
                            v29 += v29
                        } else {
                            fmt.Println(v29, v29, fi27)
                            v29 += v29
                        }
                        v9 = append(v9, 32767)
                        switch fi27 {
                        case 128:
                            if (false && (1e10 != (3.14159 * 1.5))) {
                                v9 = append(v9, (fi28 & fi27))
                                v29 += v29
                            } else {
                                var v32 int64 = 42
                                _ = v32
                            }
                        case 127:
                            if true {
                                switch 0 {
                                case 300:
                                    v33 := GenS1{F0: 4294967295, F1: -128, F2: 127}
                                    _ = v33
                                case -1:
                                    v29 = 1
                                case -216:
                                    var v34 uint32 = 2
                                    _ = v34
                                }
                                if false {
                                    v29 -= 428
                                    if len(v9) > 0 {
                                        v9[0] = -1
                                    }
                                }
                            } else {
                                fmt.Println(v29)
                                switch "" {
                                case "\n\t":
                                    var v35 uint8 = 255
                                    _ = v35
                                }
                                v29 = (v29 + v29)
                            }
                            if ((916 < 1) || ((42 < 4294967295) && true)) {
                                v36 := GenS4{F0: 782, F1: 1, F2: 513}
                                v37 := GenS4{F0: 9223372036854775807, F1: 0, F2: 42}
                                var v38 int8 = 1
                                _ = v36
                                _ = v37
                                _ = v38
                            } else {
                                v39 := GenS5{F0: 128, F1: 9223372036854775807, F2: -9223372036854775808}
                                if true {
                                    fmt.Println(v29)
                                    var v40 float32 = 0.0
                                    var v41 int16 = 0
                                    _ = v40
                                    _ = v41
                                } else {
                                    v29 -= (v29 - v29)
                                    v29 = v29
                                }
                                _ = v39
                            }
                            var v42 uint16 = 1
                            _ = v42
                        case 975:
                            var v43 string = "a"
                            _ = v43
                        default:
                            switch "" {
                            case "foo bar":
                                v29 = (v29 + v29)
                            case "\n\t":
                                var v44 float64 = -2.25
                                _ = v44
                            case "hello":
                                fmt.Println(fi27, fi27)
                            default:
                                if ((-(fi28)) >= fi27) {
                                    v45 := GenS6{F0: -1, F1: v29}
                                    fmt.Println(fi27)
                                    var v46 int8 = 1
                                    _ = v45
                                    _ = v46
                                } else {
                                    v29 = 65535
                                    var v47 uint = v29
                                    v48 := GenS7{F0: 495}
                                    _ = v47
                                    _ = v48
                                }
                            }
                        }
                        _ = v29
                    }
                    v49 := GenS8{F0: 2, F1: 18446744073709551615}
                    var v50 string = "a"
                    fmt.Println(v50, fi28)
                    var v51 int64 = 42
                    if (-57 < -41) {
                        _ = len(v9)
                        var v52 bool = false
                        v50 = (v50 + v50)
                        _ = v52
                    } else {
                        fmt.Println(fi28, fi27)
                        v53 := GenS9{F0: 255, F1: 127, F2: 42}
                        var v54 int8 = 42
                        var v55 uint64 = 2
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                    if true {
                        if len(v9) > 0 {
                            v9[0] = 127
                        }
                        var v56 float64 = (0.0 - (1e10 - (3.14159 + -2.25)))
                        v51 = (v51 & (v51 & v51))
                        _ = v56
                    }
                    _ = fi28
                    _ = v49
                    _ = v50
                    _ = v51
                }
                _ = fi27
            }
        } else {
            if (!(65535 != 0)) {
                v57 := GenS7{F0: 2}
                if (!(65535 <= 65535)) {
                    var v58 int8 = -1
                    fmt.Println(v58)
                    var v59 uint = 2
                    _ = v58
                    _ = v59
                } else {
                    for fi60 := 0; fi60 < 17; fi60++ {
                        var v61 uint8 = 127
                        v62 := GenS4{F0: 0, F1: -878, F2: 915}
                        var v63 int32 = 0
                        _ = fi60
                        _ = v61
                        _ = v62
                        _ = v63
                    }
                    v64 := GenS8{F0: 2, F1: 0}
                    _ = v64
                }
                _ = v57
            } else {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            }
            if (!(true || (!false))) {
                var v65 int8 = -70
                switch ("\n\t" + ("\n\t" + "foo bar")) {
                case "":
                    v66 := GenS10{F0: -32768, F1: -1}
                    _ = v66
                }
                _ = v65
            } else {
                var v67 string = "hello"
                _ = v67
            }
            var v68 int32 = 1
            _ = v68
        }
        _ = v23
    }
    if (0 == 97) {
        var v69 float64 = (1.5 / 3.14159)
        if true {
            fmt.Println(v69)
            for fi70 := 0; fi70 < 12; fi70++ {
                _ = len(v9)
                _ = fi70
            }
            fmt.Println(v69)
        }
        for fi71 := 0; fi71 < 15; fi71++ {
            _ = len(v9)
            _ = fi71
        }
        _ = v69
    }
    _ = v9
    return 1
}
func genFn1(p0 uint, p1 uint) int32 {
    if true {
        for fi72 := 0; fi72 < 16; fi72++ {
            p1 = 18446744073709551615
            switch ("hello" + ("" + "foo bar")) {
            case "\n\t":
                v73 := GenS11{F0: -425}
                _ = v73
            default:
                for fi74 := 0; fi74 < 8; fi74++ {
                    v75 := GenS12{F0: 42, F1: ""}
                    _ = fi74
                    _ = v75
                }
            }
            v76 := []int{}
            _ = fi72
            _ = v76
        }
    } else {
        p1 = p1
        switch "a" {
        case "foo bar":
            switch 922 {
            case 9223372036854775807:
                v77 := []int{128, 42, 128}
                _ = v77
            case 42:
                fmt.Println(p0, p0)
            }
        case "a":
            var v78 int = -72
            _ = v78
        case "\n\t":
            for fi79 := 0; fi79 < 19; fi79++ {
                v80 := GenS11{F0: 9223372036854775807}
                p0 -= 4294967295
                v81 := GenS7{F0: 1}
                _ = fi79
                _ = v80
                _ = v81
            }
        }
        v82 := []int{-2147483648, 4294967295, -471}
        _ = v82
    }
    switch 42 {
    case 32767:
        switch "hello" {
        case "\n\t":
            if (758 == 417) {
                var v83 int8 = 127
                _ = v83
            } else {
                for fi84 := 0; fi84 < 16; fi84++ {
                    fmt.Println(p1, p1, p1)
                    for fi85 := 0; fi85 < 18; fi85++ {
                        if (351 > -9223372036854775808) {
                            var v86 float32 = ((-2.25 / 3.14159) / 1.0)
                            _ = v86
                        }
                        _ = fi85
                    }
                    v87 := GenS13{F0: 127, F1: 89}
                    _ = fi84
                    _ = v87
                }
                for fi88 := 0; fi88 < 2; fi88++ {
                    var v89 int32 = 574
                    fmt.Println(p1)
                    _ = fi88
                    _ = v89
                }
            }
        case "a":
            fmt.Println(p0)
        }
    }
    p0 -= (p1 & p0)
    var v90 float32 = 0.1
    _ = p0
    _ = p1
    _ = v90
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{65535}
    var v3 int16 = 42
    v4 := make(map[string]int)
    var v5 int16 = v3
    fmt.Println(v3, v3)
    v3 -= (v5 ^ (v3 ^ 32767))
    if false {
        for fi91 := 0; fi91 < 16; fi91++ {
            var v92 bool = (-78 == 42)
            var v93 bool = (343 < 0)
            _ = fi91
            _ = v92
            _ = v93
        }
    }
    v94 := GenS14{F0: -9223372036854775808, F1: 42}
    v2 = append(v2, 65535)
    var v95 bool = (!false)
    for fi96 := 0; fi96 < 5; fi96++ {
        if (-2.25 != -2.25) {
            for fi97 := 0; fi97 < 17; fi97++ {
                v98 := v4[""]
                v99 := GenS1{F0: 65535, F1: -128, F2: 255}
                _ = fi97
                _ = v98
                _ = v99
            }
            v3 = (v5 % (((v5 * 0)) | 1))
            if (143 >= 65535) {
                for fi100 := 0; fi100 < 7; fi100++ {
                    delete(v4, ("hello" + ("foo bar" + ("hello" + ("foo bar" + "hello")))))
                    if (!true) {
                        if (547 <= 1) {
                            v5 = v5
                            var v101 float64 = 3.14159
                            _ = v101
                        } else {
                            if (v95 && true) {
                                if v95 {
                                    v5 -= (^(v5))
                                    if false {
                                        switch ("foo bar" + (("a" + (("hello" + "foo bar") + "foo bar")) + "foo bar")) {
                                        case "":
                                            var v102 uint64 = 750
                                            _ = v102
                                        case "foo bar":
                                            _ = len(v4)
                                        case "a":
                                            v5 = -852
                                        default:
                                            switch ("\n\t" + "\n\t") {
                                            case "a":
                                                var v103 float32 = 3.14159
                                                _ = v103
                                            case "\n\t":
                                                if v95 {
                                                    v104 := GenS15{F0: -9223372036854775808}
                                                    _ = v104
                                                } else {
                                                    delete(v1, "")
                                                }
                                            }
                                        }
                                        var v105 float32 = 0.1
                                        var v106 bool = v95
                                        var v107 int = 514
                                        v3 += int16(fi100)
                                        var v108 bool = (4294967295 == 2)
                                        _ = v105
                                        _ = v106
                                        _ = v107
                                        _ = v108
                                    } else {
                                        if (v95 && v95) {
                                            switch "foo bar" {
                                            case "a":
                                                _ = len(v1)
                                            case "foo bar":
                                                _ = len(v4)
                                            }
                                            v5 = -1
                                            if v95 {
                                                if v95 {
                                                    switch (fi96 + 42) {
                                                    case 739:
                                                        v3 -= -509
                                                    case 32767:
                                                        if (v3 <= 123) {
                                                            var v109 int32 = 2147483647
                                                            if v95 {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v5 = 668
                                                            } else {
                                                                fmt.Println(v95, fi96)
                                                            }
                                                            v110 := GenS16{F0: (-(v3)), F1: 387}
                                                            _ = v109
                                                            _ = v110
                                                        } else {
                                                            v111 := GenS3{F0: genFn1(381, 42)}
                                                            var v112 string = "hello"
                                                            if (!(-1 == 9223372036854775807)) {
                                                                v5 += v3
                                                                v113 := GenS7{F0: 46}
                                                                _ = v113
                                                            }
                                                            _ = v111
                                                            _ = v112
                                                        }
                                                    default:
                                                        v2 = append(v2, 528)
                                                        v114 := GenS15{F0: 1}
                                                        _ = v114
                                                    }
                                                    v5 -= (v5 ^ -375)
                                                    fmt.Println(v5, v95)
                                                } else {
                                                    v95 = (47 != -80)
                                                    var v115 float32 = (1.5 * 1e10)
                                                    v116 := GenS8{F0: 18446744073709551615, F1: 454}
                                                    v117 := GenS17{F0: 128, F1: 42}
                                                    _ = v115
                                                    _ = v116
                                                    _ = v117
                                                }
                                                fmt.Println(fi96, v5, v3)
                                                fmt.Println(fi96, fi100, fi100)
                                            }
                                            if (-9223372036854775808 <= -1) {
                                                if (0 != (-(v3))) {
                                                    v95 = false
                                                    var v118 uint8 = 0
                                                    v119 := GenS13{F0: v118, F1: -57}
                                                    var v120 uint16 = 207
                                                    v2 = append(v2, 636)
                                                    _ = v118
                                                    _ = v119
                                                    _ = v120
                                                }
                                                if false {
                                                    if (0 < 183) {
                                                        v121 := GenS16{F0: (v5 & 730), F1: 18446744073709551615}
                                                        var v122 bool = (v95 || v95)
                                                        fmt.Println(v5, v3, v3)
                                                        v3 += (v3 | 1)
                                                        _ = v121
                                                        _ = v122
                                                    } else {
                                                        v123 := GenS18{F0: true, F1: (fi96 * 727)}
                                                        _ = v123
                                                    }
                                                } else {
                                                    if true {
                                                        var v124 bool = (2147483648 != 4294967295)
                                                        v5 += 32767
                                                        _ = v124
                                                    }
                                                    v125 := GenS9{F0: 0, F1: 255, F2: 42}
                                                    switch "a" {
                                                    case "\n\t":
                                                        v3 = (v3 * (v5 % ((v5) | 1)))
                                                    case "":
                                                        fmt.Println(fi96, v3)
                                                    case "a":
                                                        var v126 float64 = 1e10
                                                        _ = v126
                                                    default:
                                                        v95 = (v5 >= 42)
                                                    }
                                                    v5 = 94
                                                    _ = v125
                                                }
                                            } else {
                                                v127 := GenS6{F0: 1, F1: 42}
                                                v1[("hello" + "a")] = (fi96 & fi96)
                                                if (v95 && (v3 == (v3 >> ((2) & 63)))) {
                                                    switch fi100 {
                                                    case 31:
                                                        if (!false) {
                                                            v3 += v3
                                                            var v128 uint8 = 2
                                                            switch (("hello" + "") + "foo bar") {
                                                            case "foo bar":
                                                                if len(v2) > 0 {
                                                                    v2[0] = 240
                                                                }
                                                            case "\n\t":
                                                                switch 9223372036854775807 {
                                                                case -75:
                                                                    var v129 uint32 = 1
                                                                    _ = v129
                                                                case 0:
                                                                    if (v95 || (1 == 4294967295)) {
                                                                        v3 += (v3 << ((4294967295) & 63))
                                                                    } else {
                                                                        v128 = (v128 ^ v128)
                                                                    }
                                                                case 2147483647:
                                                                    var v130 string = ("" + "a")
                                                                    _ = v130
                                                                default:
                                                                    if (!((v128 ^ v128) < v128)) {
                                                                        var v131 int = fi100
                                                                        _ = v131
                                                                    }
                                                                }
                                                            default:
                                                                if v95 {
                                                                    var v132 float64 = -2.25
                                                                    var v133 bool = (4294967295 <= 4294967295)
                                                                    if true {
                                                                        v134 := GenS9{F0: 65535, F1: (v128 >> ((2) & 63)), F2: -128}
                                                                        var v135 uint16 = 1
                                                                        _ = v134
                                                                        _ = v135
                                                                    }
                                                                    _ = v132
                                                                    _ = v133
                                                                } else {
                                                                    fmt.Println(v95)
                                                                    fmt.Println(v95, fi100, v5)
                                                                }
                                                            }
                                                            _ = v128
                                                        } else {
                                                            v136 := GenS19{F0: 0, F1: v95, F2: -9223372036854775808}
                                                            if len(v2) > 0 {
                                                                v2[0] = 192
                                                            }
                                                            if v95 {
                                                                if len(v2) > 0 {
                                                                    v2[0] = 739
                                                                }
                                                                switch "foo bar" {
                                                                case "":
                                                                    v5 -= v5
                                                                    var v137 int8 = 0
                                                                    _ = v137
                                                                case "hello":
                                                                    v3 = 42
                                                                default:
                                                                    switch ((("a" + "\n\t") + "") + "") {
                                                                    case "a":
                                                                        var v138 int32 = -2147483648
                                                                        _ = v138
                                                                    default:
                                                                        delete(v4, "\n\t")
                                                                        fmt.Println(fi100)
                                                                        var v139 int8 = 100
                                                                        _ = v139
                                                                    }
                                                                }
                                                            }
                                                            var v140 uint64 = 18446744073709551615
                                                            _ = v136
                                                            _ = v140
                                                        }
                                                    }
                                                    var v141 int = fi100
                                                    _ = v141
                                                } else {
                                                    var v142 int64 = -1
                                                    var v143 float32 = 0.0
                                                    _ = len(v4)
                                                    _ = v142
                                                    _ = v143
                                                }
                                                _ = v127
                                            }
                                            v3 = (v5 / (((v5 + v5)) | 1))
                                            v144 := GenS20{F0: 65535, F1: v3}
                                            _ = v144
                                        }
                                        delete(v4, "")
                                    }
                                    var v145 int16 = v5
                                    _ = v145
                                }
                            } else {
                                v95 = (2147483647 != 42)
                                var v146 float64 = 0.0
                                switch "" {
                                case "hello":
                                    var v147 string = "a"
                                    _ = v147
                                }
                                if false {
                                    v5 = 1000
                                    if v95 {
                                        if v95 {
                                            if (2 == 564) {
                                                v95 = v95
                                                v3 += (-(v5))
                                            }
                                        } else {
                                            var v148 uint64 = 860
                                            delete(v1, "a")
                                            _ = v148
                                        }
                                        var v149 int16 = -1
                                        switch "" {
                                        case "\n\t":
                                            v150 := GenS21{F0: "a", F1: 18446744073709551615}
                                            _ = v150
                                        case "foo bar":
                                            var v151 bool = false
                                            _ = v151
                                        case "":
                                            v5 += (v5 + (v5 & v5))
                                        }
                                        v149 -= (v5 ^ (v149 - -1))
                                        _ = v149
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = 309
                                        }
                                        var v152 uint64 = 18446744073709551615
                                        v3 = (v3 - 32767)
                                        _ = v152
                                    }
                                }
                                _ = v146
                            }
                        }
                    }
                    var v153 float32 = 0.0
                    v153 = v153
                    _ = fi100
                    _ = v153
                }
            }
        }
        _ = fi96
    }
    if v95 {
        delete(v1, "\n\t")
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v95 = (2 == 391)
    }
    v154 := GenS21{F0: "", F1: 18446744073709551615}
    v3 += 247
    var v155 int = 131
    var v156 uint16 = 355
    switch 32767 {
    case -1:
        if len(v2) > 0 {
            v2[0] = 860
        }
    case 1:
        switch 0 {
        case 127:
            _ = len(v1)
        case 0:
            switch "a" {
            case "":
                for fi157 := 0; fi157 < 11; fi157++ {
                    if len(v2) > 0 {
                        v2[0] = 961
                    }
                    _ = fi157
                }
            }
        default:
            for fi158 := 0; fi158 < 4; fi158++ {
                var v159 float32 = -2.25
                for fi160 := 0; fi160 < 9; fi160++ {
                    v159 = 3.14159
                    _ = fi160
                }
                var v161 int = -34
                _ = fi158
                _ = v159
                _ = v161
            }
        }
    default:
        v156 = v156
    }
    for fi162 := 0; fi162 < 14; fi162++ {
        if v95 {
            for fi163 := 0; fi163 < 4; fi163++ {
                if (42 <= 676) {
                    var v164 int8 = -1
                    _ = v164
                } else {
                    v156 -= v156
                    v156 += v156
                    var v165 uint32 = 666
                    _ = v165
                }
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
                if true {
                    var v166 bool = v95
                    _ = v166
                }
                _ = fi163
            }
            for fi167 := 0; fi167 < 18; fi167++ {
                var v168 float64 = -2.25
                var v169 uint = 0
                _ = fi167
                _ = v168
                _ = v169
            }
        } else {
            var v170 uint32 = 2
            var v171 uint8 = 47
            _ = v170
            _ = v171
        }
        var v172 bool = v95
        if true {
            v172 = (!v172)
        }
        _ = fi162
        _ = v172
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v155:", v155)
    fmt.Println("v156:", v156)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v154
    _ = v155
    _ = v156
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v94
    _ = v95
}
