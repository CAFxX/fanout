package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 int32; F1 uint }
type GenS7 struct { F0 uint64; F1 int; F2 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint64; F1 int; F2 int8 }
type GenS13 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS14 struct { F0 uint8; F1 int8 }
type GenS15 struct { F0 string; F1 int32 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 bool; F1 int; F2 int }
type GenS18 struct { F0 string; F1 int }
type GenS19 struct { F0 int32; F1 int64 }
type GenS20 struct { F0 uint16; F1 int8; F2 int8 }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 int16; F1 uint16 }
type GenS23 struct { F0 uint8; F1 uint }
func genFn0(p0 int32, p1 int8, p2 int8) int64 {
    v5 := []int{127, 32768, 0, 42}
    if (!true) {
        if false {
            if (42 > 0) {
                v6 := GenS1{F0: 1}
                p2 += 127
                var v7 uint8 = 187
                _ = v6
                _ = v7
            }
            switch "\n\t" {
            case "":
                for fi8 := 0; fi8 < 12; fi8++ {
                    _ = len(v5)
                    _ = fi8
                }
            default:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            var v9 float64 = float64(-2.25)
            _ = v9
        }
        var v10 uint16 = 352
        _ = v10
    } else {
        if false {
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 float32 = (3.14159 - 1e10)
                _ = len(v5)
                if true {
                    v12 = float32(1.5)
                    fmt.Println(p2, p0, p0)
                } else {
                    p0 = p0
                    var v13 float32 = 0.0
                    for fi14 := 0; fi14 < 6; fi14++ {
                        var v15 int32 = 42
                        _ = fi14
                        _ = v15
                    }
                    _ = v13
                }
                _ = fi11
                _ = v12
            }
            switch "\n\t" {
            case "hello":
                var v16 uint16 = 32767
                _ = v16
            case "foo bar":
                switch 0 {
                case -169:
                    var v17 uint = 2
                    _ = v17
                case 1:
                    var v18 float32 = float32(3.14159)
                    _ = v18
                case 32768:
                    p1 = 117
                default:
                    switch 42 {
                    case 2147483647:
                        for fi19 := 0; fi19 < 5; fi19++ {
                            if false {
                                if false {
                                    p1 = (p1 ^ (p2 * (p2 >> ((65535) & 63))))
                                    if (p0 == 0) {
                                        fmt.Println(p0, p0)
                                        fmt.Println(p2)
                                    } else {
                                        fmt.Println(p2)
                                        var v20 uint64 = 540
                                        _ = v20
                                    }
                                    v21 := GenS1{F0: 32767}
                                    _ = v21
                                } else {
                                    var v22 uint = 1
                                    v23 := GenS2{F0: 2, F1: 43, F2: true}
                                    _ = v22
                                    _ = v23
                                }
                                switch ("\n\t" + (("" + "foo bar") + "\n\t")) {
                                case "\n\t":
                                    p1 = -128
                                default:
                                    var v24 bool = (2 > 146)
                                    _ = v24
                                }
                            } else {
                                for fi25 := 0; fi25 < 7; fi25++ {
                                    if true {
                                        fmt.Println(p1, p1, p2)
                                    }
                                    var v26 string = "a"
                                    _ = fi25
                                    _ = v26
                                }
                            }
                            _ = fi19
                        }
                    case 128:
                        if len(v5) > 0 {
                            v5[0] = 1
                        }
                    case -263:
                        p1 = p2
                    default:
                        v27 := GenS3{F0: (true && (false && true))}
                        _ = v27
                    }
                }
            }
        }
        if false {
            v28 := GenS1{F0: -816}
            _ = v28
        }
        p0 += -2147483648
    }
    fmt.Println(p1, p1, p1)
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -249
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[int]int)
    var v3 int = v1
    v4 := make(map[string]int)
    v29 := []int{}
    if false {
        var v30 uint8 = 2
        var v31 string = "foo bar"
        _ = v30
        _ = v31
    } else {
        for fi32 := 0; fi32 < 20; fi32++ {
            v33 := GenS3{F0: true}
            v34 := GenS3{F0: (!false)}
            v35 := GenS2{F0: 117, F1: 0, F2: (65535 == 65535)}
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
        v36 := GenS2{F0: 1, F1: 100, F2: true}
        if (false && true) {
            if ((v3 * (-(v1))) > v1) {
                var v37 bool = (2 < 1)
                for fi38 := 0; fi38 < 10; fi38++ {
                    v39 := v2[-107]
                    if v37 {
                        switch "a" {
                        case "hello":
                            v37 = (600 < 0)
                        case "\n\t":
                            for fi40 := 0; fi40 < 5; fi40++ {
                                switch (v39 % ((fi40) | 1)) {
                                case -1:
                                    v3 += v3
                                    switch (v39 ^ 9223372036854775807) {
                                    case 4294967295:
                                        v41 := GenS3{F0: (99 <= 32767)}
                                        _ = v41
                                    case 9223372036854775807:
                                        var v42 int64 = -270
                                        _ = v42
                                    case 0:
                                        switch fi38 {
                                        case -956:
                                            fmt.Println(v39, v39, v3)
                                        case 128:
                                            switch "\n\t" {
                                            case "a":
                                                switch ("hello" + "\n\t") {
                                                case "foo bar":
                                                    var v43 bool = (-812 > 0)
                                                    _ = v43
                                                default:
                                                    switch (("a" + "\n\t") + "\n\t") {
                                                    case "a":
                                                        v37 = v37
                                                    }
                                                }
                                            case "\n\t":
                                                var v44 float64 = 1e10
                                                _ = v44
                                            case "hello":
                                                fmt.Println(fi40, v37, v39)
                                            default:
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v45 float32 = 0.1
                                                    _ = v45
                                                default:
                                                    v37 = true
                                                }
                                            }
                                        case 199:
                                            var v46 string = "\n\t"
                                            _ = v46
                                        }
                                    }
                                    v4["\n\t"] = -232
                                    var v47 int32 = int32(v39)
                                    _ = v47
                                }
                                var v48 uint = 65535
                                _ = fi40
                                _ = v48
                            }
                        case "":
                            v39 = (fi38 % ((32768) | 1))
                        default:
                            fmt.Println(v1)
                        }
                        if (float64(1.5) <= float64(3.14159)) {
                            v37 = false
                            if v37 {
                                v49 := v2[(v3 >> ((65535) & 63))]
                                v50 := GenS4{F0: 42, F1: (!v37), F2: 2147483647}
                                _ = v49
                                _ = v50
                            }
                            for fi51 := 0; fi51 < 15; fi51++ {
                                v52 := GenS5{F0: 1, F1: 42}
                                v53 := GenS3{F0: v37}
                                v1 -= -2147483648
                                _ = fi51
                                _ = v52
                                _ = v53
                            }
                        } else {
                            v54 := GenS6{F0: 1, F1: 779}
                            for fi55 := 0; fi55 < 5; fi55++ {
                                v1 -= fi38
                                v56 := GenS4{F0: 65535, F1: v37, F2: 4294967295}
                                _ = fi55
                                _ = v56
                            }
                            for fi57 := 0; fi57 < 6; fi57++ {
                                var v58 float64 = 0.0
                                switch -9223372036854775808 {
                                case 4294967295:
                                    v58 = 0.1
                                case 983:
                                    if ((!true) || v37) {
                                        v37 = false
                                        switch 127 {
                                        case -232:
                                            delete(v2, v3)
                                        case 4294967295:
                                            var v59 bool = v37
                                            _ = v59
                                        case -1:
                                            var v60 int32 = 181
                                            var v61 uint32 = 313
                                            if (v37 && (fi57 <= -943)) {
                                                var v62 int64 = 1
                                                _ = v62
                                            }
                                            _ = v60
                                            _ = v61
                                        }
                                        if v37 {
                                            if (!v37) {
                                                var v63 int8 = 42
                                                v29 = append(v29, v3)
                                                v3 += v1
                                                _ = v63
                                            }
                                            v58 += v58
                                        } else {
                                            if len(v29) >= 2 {
                                                v29 = v29[1:len(v29)]
                                            }
                                            var v64 uint64 = 1
                                            var v65 uint = 527
                                            _ = v64
                                            _ = v65
                                        }
                                    }
                                default:
                                    switch ("" + ("hello" + "a")) {
                                    case "foo bar":
                                        var v66 float64 = float64(1.5)
                                        _ = v66
                                    }
                                }
                                _ = fi57
                                _ = v58
                            }
                            _ = v54
                        }
                    } else {
                        v67 := GenS7{F0: 18446744073709551615, F1: fi38, F2: 0}
                        if (65535 != 2) {
                            for fi68 := 0; fi68 < 8; fi68++ {
                                var v69 int32 = -636
                                v69 = (v69 & v69)
                                var v70 int8 = 127
                                _ = fi68
                                _ = v69
                                _ = v70
                            }
                            fmt.Println(v3)
                        } else {
                            if len(v29) > 0 {
                                v29[0] = fi38
                            }
                            switch 127 {
                            case 0:
                                v71 := GenS3{F0: v37}
                                _ = v71
                            }
                        }
                        fmt.Println(v39, v37)
                        _ = v67
                    }
                    if (1 < 1) {
                        v37 = (!(false && (1.5 <= 0.1)))
                        v37 = v37
                        v1 += (v39 - fi38)
                    } else {
                        var v72 uint16 = 42
                        fmt.Println(v39, v72, v1)
                        v73 := GenS5{F0: 42, F1: 657}
                        _ = v72
                        _ = v73
                    }
                    _ = fi38
                    _ = v39
                }
                _ = v37
            }
            v1 -= 127
        } else {
            var v74 float64 = (1e10 * 1e10)
            for fi75 := 0; fi75 < 19; fi75++ {
                v1 -= 200
                for fi76 := 0; fi76 < 18; fi76++ {
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                    var v77 float64 = 0.0
                    v78 := GenS8{F0: 253}
                    var v79 uint = 255
                    _ = fi76
                    _ = v77
                    _ = v78
                    _ = v79
                }
                var v80 int32 = -2147483648
                _ = fi75
                _ = v80
            }
            _ = v74
        }
        _ = v36
    }
    if true {
        if ((("\n\t" + "foo bar") + "foo bar") < "foo bar") {
            v2[v1] = 667
            v81 := GenS9{F0: 65535, F1: 601, F2: 0}
            _ = v81
        } else {
            if (!(1 <= 42)) {
                var v82 string = (("hello" + ("hello" + ((("foo bar" + "\n\t") + "a") + "\n\t"))) + (("" + "foo bar") + ("\n\t" + "a")))
                for fi83 := 0; fi83 < 9; fi83++ {
                    var v84 bool = true
                    v1 = -657
                    fmt.Println(v84, v1, v1)
                    _ = fi83
                    _ = v84
                }
                v4["hello"] = v1
                _ = v82
            }
            fmt.Println(v1, v1)
        }
        var v85 int64 = -9223372036854775808
        v85 = v85
        _ = v85
    } else {
        if true {
            v86 := GenS7{F0: 0, F1: (v3 | v3), F2: 1000}
            fmt.Println(v1)
            if (!true) {
                if false {
                    var v87 int = (-(v1))
                    _ = v87
                }
                switch v1 {
                case -321:
                    if ("foo bar" > ("\n\t" + "a")) {
                        for fi88 := 0; fi88 < 9; fi88++ {
                            for fi89 := 0; fi89 < 12; fi89++ {
                                var v90 int = 0
                                if false {
                                    v29 = append(v29, (v3 | (fi88 - 1)))
                                    if (true || false) {
                                        if len(v29) > 0 {
                                            v29[0] = v3
                                        }
                                    } else {
                                        fmt.Println(v1, fi89, v1)
                                        v90 -= fi88
                                    }
                                    switch "\n\t" {
                                    case "foo bar":
                                        v3 = (fi89 * 617)
                                    case "hello":
                                        if false {
                                            v29 = append(v29, v3)
                                            switch v3 {
                                            case 0:
                                                v1 -= (fi88 << ((409) & 63))
                                            case 9223372036854775807:
                                                v90 = 9223372036854775807
                                            }
                                            v1 += v90
                                        }
                                    default:
                                        v91 := GenS10{F0: true}
                                        _ = v91
                                    }
                                } else {
                                    v92 := GenS11{F0: 14}
                                    v90 = v3
                                    if (true && true) {
                                        var v93 int64 = -1
                                        var v94 int16 = 1000
                                        if (128 != 255) {
                                            switch "\n\t" {
                                            case "":
                                                v1 += (v1 - v1)
                                                v90 = (fi88 * 127)
                                                if (!(false || (1e10 >= 1.5))) {
                                                    v2[v1] = fi88
                                                }
                                                switch (v90 - (v90 / ((v3) | 1))) {
                                                case -43:
                                                    v95 := GenS12{F0: 481, F1: 32768, F2: 1}
                                                    _ = v95
                                                case 1:
                                                    v90 += 0
                                                case 0:
                                                    if (fi89 <= 4294967295) {
                                                        v3 = 1
                                                        var v96 int32 = 712
                                                        _ = v96
                                                    } else {
                                                        var v97 uint8 = 0
                                                        _ = v97
                                                    }
                                                default:
                                                    var v98 uint16 = 65535
                                                    _ = v98
                                                }
                                                v93 += v93
                                                v94 += (v94 & (-(v94)))
                                            case "foo bar":
                                                var v99 float64 = 1.5
                                                _ = v99
                                            case "hello":
                                                var v100 uint16 = 255
                                                _ = v100
                                            }
                                            v90 -= (^(v90))
                                            switch (fi89 % ((fi89) | 1)) {
                                            case 0:
                                                switch (fi88 / ((fi88) | 1)) {
                                                case 32768:
                                                    v101 := GenS11{F0: 0}
                                                    _ = v101
                                                default:
                                                    switch "" {
                                                    case "":
                                                        var v102 uint32 = 65535
                                                        _ = v102
                                                    case "foo bar":
                                                        v103 := GenS5{F0: 127, F1: 42}
                                                        _ = v103
                                                    default:
                                                        v1 -= v3
                                                    }
                                                }
                                            case 1:
                                                var v104 bool = true
                                                _ = v104
                                            case -480:
                                                v1 -= v3
                                            }
                                        }
                                        _ = v93
                                        _ = v94
                                    }
                                    if (!(false && ((((127 > 42) || (-1 != 42)) && ((!(42 == 2)) && true)) || false))) {
                                        v105 := GenS13{F0: 9223372036854775807, F1: 179, F2: 643}
                                        v106 := GenS14{F0: 2, F1: 0}
                                        _ = v105
                                        _ = v106
                                    } else {
                                        if (genFn0(42, 55, 42) < 0) {
                                            var v107 int64 = 10
                                            v108 := GenS15{F0: "\n\t", F1: -2147483648}
                                            fmt.Println(v90, v1)
                                            v90 += fi89
                                            _ = v107
                                            _ = v108
                                        }
                                        v90 += v1
                                    }
                                    _ = v92
                                }
                                _ = fi89
                                _ = v90
                            }
                            v3 = (-(v1))
                            fmt.Println(v1, fi88)
                            _ = fi88
                        }
                        var v109 uint64 = 18446744073709551615
                        _ = v109
                    }
                case 42:
                    switch (v3 >> ((255) & 63)) {
                    case -669:
                        v1 = (v1 % ((-557) | 1))
                    default:
                        var v110 float64 = (-2.25 - 3.14159)
                        _ = v110
                    }
                case -882:
                    v1 = v1
                }
            } else {
                v111 := v4["hello"]
                _ = v111
            }
            _ = v86
        } else {
            fmt.Println(v3)
            var v112 uint8 = 127
            v113 := GenS8{F0: (v112 ^ v112)}
            _ = v112
            _ = v113
        }
        var v114 int64 = -123
        var v115 int = v3
        _ = v114
        _ = v115
    }
    var v116 uint8 = 243
    delete(v4, ("foo bar" + "a"))
    v117 := GenS11{F0: int8(v116)}
    var v118 float64 = (3.14159 * (3.14159 + -2.25))
    v2[(^(v1))] = 42
    _ = len(v2)
    var v119 bool = (v116 < (v116 << ((18446744073709551615) & 63)))
    fmt.Println(v3, v116)
    if ((v1 % (((v3 / ((v3) | 1))) | 1)) < (v3 + (v3 % ((v3) | 1)))) {
        if false {
            for fi120 := 0; fi120 < 19; fi120++ {
                var v121 int = 127
                for fi122 := 0; fi122 < 3; fi122++ {
                    if v119 {
                        if false {
                            v121 = (fi120 | 9223372036854775807)
                            v123 := GenS16{F0: 787}
                            _ = v123
                        } else {
                            var v124 float32 = 0.1
                            _ = v124
                        }
                        switch fi120 {
                        case -9223372036854775808:
                            v125 := v4[(("\n\t" + "a") + "")]
                            _ = v125
                        case -403:
                            if true {
                                if true {
                                    if (-1 >= 100) {
                                        var v126 int = v121
                                        v126 -= fi120
                                        _ = v126
                                    }
                                    _ = len(v4)
                                } else {
                                    var v127 string = ("a" + ("hello" + ("" + "foo bar")))
                                    _ = v127
                                }
                                v128 := GenS17{F0: v119, F1: -518, F2: (v1 % ((fi122) | 1))}
                                _ = v128
                            } else {
                                var v129 string = "\n\t"
                                fmt.Println(v3, fi122)
                                if v119 {
                                    switch fi122 {
                                    case 9223372036854775807:
                                        var v130 uint16 = 2
                                        _ = v130
                                    case -992:
                                        switch "" {
                                        case "":
                                            if v119 {
                                                v1 -= (v121 % ((-9223372036854775808) | 1))
                                            } else {
                                                v119 = false
                                                switch ("a" + "\n\t") {
                                                case "":
                                                    fmt.Println(v119, fi120)
                                                default:
                                                    fmt.Println(v129)
                                                }
                                            }
                                            v121 = (^(v3))
                                        case "hello":
                                            v116 = v116
                                        }
                                    case 42:
                                        if (v118 <= v118) {
                                            delete(v4, v129)
                                            if true {
                                                fmt.Println(v116, v116)
                                            } else {
                                                v118 -= 1.5
                                            }
                                        }
                                    default:
                                        v131 := GenS1{F0: -171}
                                        _ = v131
                                    }
                                    if true {
                                        v121 = 42
                                    } else {
                                        fmt.Println(v3, v116, v121)
                                    }
                                    v119 = v119
                                    v2[(^(v121))] = v3
                                    v119 = v119
                                    v29 = append(v29, (^(fi122)))
                                } else {
                                    v119 = ((v116 % ((5) | 1)) <= 2)
                                    switch fi120 {
                                    case 127:
                                        if len(v29) > 0 {
                                            v29[0] = v3
                                        }
                                    case 939:
                                        v132 := GenS16{F0: 42}
                                        _ = v132
                                    case 32768:
                                        v2[(v121 << ((4294967295) & 63))] = fi122
                                    default:
                                        switch (v129 + v129) {
                                        case "\n\t":
                                            if (!(v129 < v129)) {
                                                v119 = (!v119)
                                                v121 = 42
                                                var v133 int16 = 42
                                                switch (-(fi120)) {
                                                case 4294967295:
                                                    if (true || (v116 <= 255)) {
                                                        v119 = v119
                                                    } else {
                                                        switch ("hello" + "foo bar") {
                                                        case "":
                                                            var v134 uint32 = 49
                                                            _ = v134
                                                        case "foo bar":
                                                            if v119 {
                                                                if true {
                                                                    v116 = 0
                                                                } else {
                                                                    fmt.Println(fi120)
                                                                }
                                                                if v119 {
                                                                    var v135 uint64 = 426
                                                                    v3 = 42
                                                                    var v136 float32 = (0.0 * -2.25)
                                                                    _ = v135
                                                                    _ = v136
                                                                } else {
                                                                    switch ((((v129 + "") + (v129 + (((v129 + "foo bar") + v129) + v129))) + v129) + (("hello" + "\n\t") + v129)) {
                                                                    case "hello":
                                                                        var v137 uint64 = 2
                                                                        _ = v137
                                                                    }
                                                                    if false {
                                                                        switch fi120 {
                                                                        case -61:
                                                                            v118 -= float64(0.0)
                                                                        case 42:
                                                                            v119 = v119
                                                                        case -758:
                                                                            if ((-9223372036854775808 <= 9223372036854775807) && (1 >= 65535)) {
                                                                                var v138 float64 = (v118 * (v118 * 0.0))
                                                                                fmt.Println(v116)
                                                                                _ = v138
                                                                            }
                                                                        }
                                                                        _ = len(v29)
                                                                        v139 := v4[(v129 + (v129 + "foo bar"))]
                                                                        _ = v139
                                                                    } else {
                                                                        if v119 {
                                                                            v118 -= float64((0.1 / float32(v118)))
                                                                            switch v129 {
                                                                            case "\n\t":
                                                                                v140 := GenS18{F0: "", F1: 1}
                                                                                _ = v140
                                                                            case "a":
                                                                                var v141 int8 = -1
                                                                                _ = v141
                                                                            default:
                                                                                var v142 int32 = 1
                                                                                _ = v142
                                                                            }
                                                                            var v143 int32 = -144
                                                                            _ = v143
                                                                        } else {
                                                                            var v144 int = (v3 / ((v1) | 1))
                                                                            _ = v144
                                                                        }
                                                                        v145 := GenS19{F0: -1, F1: 942}
                                                                        fmt.Println(v133)
                                                                        var v146 int64 = 638
                                                                        var v147 uint16 = 32767
                                                                        var v148 float32 = 1e10
                                                                        var v149 uint32 = 42
                                                                        v148 += float32(((v118 - float64(0.1)) * v118))
                                                                        _ = v145
                                                                        _ = v146
                                                                        _ = v147
                                                                        _ = v148
                                                                        _ = v149
                                                                    }
                                                                }
                                                            }
                                                        case "\n\t":
                                                            v133 += v133
                                                        }
                                                        if ("\n\t" < ("a" + v129)) {
                                                            if (false || v119) {
                                                                if ((765 < 467) || false) {
                                                                    var v150 float64 = v118
                                                                    if (v119 && v119) {
                                                                        v3 += v3
                                                                        if (!v119) {
                                                                            v150 = v118
                                                                        }
                                                                        v3 = (^(v121))
                                                                        var v151 uint64 = 565
                                                                        _ = v151
                                                                    } else {
                                                                        var v152 int32 = -2147483648
                                                                        if (true && v119) {
                                                                            v133 = v133
                                                                            v29 = append(v29, (fi120 ^ v121))
                                                                        }
                                                                        _ = v152
                                                                    }
                                                                    _ = v150
                                                                }
                                                            }
                                                            switch v3 {
                                                            case -650:
                                                                v118 += 0.1
                                                            case 118:
                                                                fmt.Println(fi120, v3, v1)
                                                            case 9223372036854775807:
                                                                if (622 == -2147483648) {
                                                                    fmt.Println(v118)
                                                                }
                                                            default:
                                                                v153 := GenS9{F0: 2147483647, F1: 959, F2: 2}
                                                                _ = v153
                                                            }
                                                            switch 2147483647 {
                                                            case 0:
                                                                fmt.Println(v121, fi120)
                                                            }
                                                        }
                                                        var v154 bool = v119
                                                        v118 = 3.14159
                                                        v155 := GenS20{F0: 255, F1: -90, F2: 42}
                                                        _ = v154
                                                        _ = v155
                                                    }
                                                case 800:
                                                    if false {
                                                        if len(v29) >= 2 {
                                                            v29 = v29[1:len(v29)]
                                                        }
                                                    }
                                                case 0:
                                                    if v119 {
                                                        v3 = fi120
                                                    } else {
                                                        var v156 uint64 = 2
                                                        _ = v156
                                                    }
                                                default:
                                                    fmt.Println(v1, v3, v116)
                                                }
                                                fmt.Println(fi120)
                                                fmt.Println(v119)
                                                _ = v133
                                            }
                                        case "":
                                            if v119 {
                                                switch v121 {
                                                case -2147483648:
                                                    v119 = (((1e10 + 0.0) + float32(v118)) != 0.0)
                                                case 9223372036854775807:
                                                    v157 := GenS21{F0: (1 >= 2147483647)}
                                                    _ = v157
                                                case 32767:
                                                    fmt.Println(v1, v121)
                                                default:
                                                    v1 = v121
                                                }
                                                if v119 {
                                                    if v119 {
                                                        v121 -= v3
                                                        switch "" {
                                                        case "\n\t":
                                                            v158 := GenS22{F0: int16(v1), F1: 1}
                                                            _ = v158
                                                        case "hello":
                                                            v129 = ("foo bar" + "\n\t")
                                                        case "a":
                                                            v29 = append(v29, (-(v1)))
                                                        default:
                                                            var v159 int64 = 1
                                                            _ = v159
                                                        }
                                                    }
                                                    v160 := GenS17{F0: (!false), F1: v121, F2: -13}
                                                    v161 := GenS23{F0: (v116 << ((2) & 63)), F1: 0}
                                                    _ = v160
                                                    _ = v161
                                                }
                                                var v162 uint32 = 407
                                                _ = v162
                                            }
                                        case "a":
                                            var v163 int32 = -1
                                            _ = v163
                                        default:
                                            var v164 int = (v121 / ((32768) | 1))
                                            _ = v164
                                        }
                                    }
                                    switch v129 {
                                    case "hello":
                                        if ((float32(v118) - ((0.0 / 0.1) + 0.0)) > 0.0) {
                                            if len(v29) > 0 {
                                                v29[0] = (v1 - (v1 << ((589) & 63)))
                                            }
                                            v119 = true
                                        } else {
                                            _ = len(v29)
                                            v165 := GenS13{F0: 0, F1: 255, F2: 42}
                                            _ = v165
                                        }
                                    case "foo bar":
                                        v119 = (42 >= 0)
                                    case "\n\t":
                                        v166 := GenS15{F0: v129, F1: -1}
                                        _ = v166
                                    default:
                                        v118 += float64(3.14159)
                                    }
                                }
                                _ = v129
                            }
                        case 9223372036854775807:
                            delete(v4, "")
                        default:
                            _ = len(v4)
                        }
                    }
                    var v167 int64 = 42
                    var v168 uint16 = 2
                    _ = fi122
                    _ = v167
                    _ = v168
                }
                v3 -= (v121 >> ((2) & 63))
                _ = fi120
                _ = v121
            }
            v169 := GenS20{F0: 1, F1: 23, F2: 42}
            fmt.Println(v119, v119)
            _ = v169
        }
    } else {
        switch -2147483648 {
        case -9223372036854775808:
            fmt.Println(v1, v116)
        case 2147483647:
            var v170 int32 = -1
            _ = v170
        }
    }
    for fi171 := 0; fi171 < 13; fi171++ {
        v118 = (v118 - 3.14159)
        for fi172 := 0; fi172 < 3; fi172++ {
            if (!(true || (v119 && (2 == 4294967295)))) {
                fmt.Println(v1, fi172)
                v1 -= fi171
            } else {
                v29 = append(v29, fi172)
                if (false || v119) {
                    v173 := GenS18{F0: "hello", F1: -959}
                    if (669 != 865) {
                        _ = len(v4)
                        switch "a" {
                        case "a":
                            var v174 int32 = -313
                            _ = v174
                        case "\n\t":
                            var v175 int64 = -660
                            _ = v175
                        }
                        if v119 {
                            switch 32768 {
                            case 42:
                                v118 += (0.0 - 3.14159)
                            }
                        } else {
                            switch 4294967295 {
                            case 2147483647:
                                var v176 int = (v3 * v3)
                                _ = v176
                            case -2147483648:
                                delete(v4, (("a" + ("foo bar" + "hello")) + ("hello" + "")))
                            default:
                                var v177 string = "foo bar"
                                _ = v177
                            }
                            v178 := v4[((((((("a" + ("hello" + (((("" + ("a" + "a")) + "hello") + ("foo bar" + ("a" + "hello"))) + ("" + "\n\t")))) + ("\n\t" + "a")) + ((("" + "\n\t") + ("a" + ((("a" + "") + "\n\t") + ""))) + ("foo bar" + (("\n\t" + "foo bar") + "")))) + ("" + "foo bar")) + "a") + "a") + "hello")]
                            if true {
                                var v179 uint = uint(v116)
                                v180 := GenS12{F0: 643, F1: (-(fi172)), F2: 49}
                                switch v1 {
                                case -902:
                                    if v119 {
                                        var v181 int = v3
                                        fmt.Println(v116)
                                        _ = v181
                                    } else {
                                        var v182 uint16 = 42
                                        _ = v182
                                    }
                                case 127:
                                    v1 = (v178 ^ 127)
                                }
                                var v183 uint16 = 346
                                _ = v179
                                _ = v180
                                _ = v183
                            } else {
                                if (42 == 783) {
                                    fmt.Println(fi172, v118)
                                    var v184 uint = 50
                                    fmt.Println(v178, fi171)
                                    _ = v184
                                } else {
                                    v118 = ((((v118 * -2.25) + v118) * (v118 / (v118 * (float64(-2.25) - 0.0)))) + (3.14159 - 1e10))
                                }
                                var v185 bool = (151 > 0)
                                _ = v185
                            }
                            _ = v178
                        }
                    } else {
                        var v186 uint16 = 255
                        var v187 float64 = (1e10 - (v118 + 1.5))
                        v119 = v119
                        _ = v186
                        _ = v187
                    }
                    _ = v173
                } else {
                    var v188 int32 = int32(fi172)
                    delete(v4, "hello")
                    _ = v188
                }
            }
            if v119 {
                if v119 {
                    fmt.Println(v3)
                    if true {
                        var v189 int = 2147483647
                        var v190 bool = (1 <= 227)
                        v3 = 128
                        _ = v189
                        _ = v190
                    }
                    if (42 >= 2) {
                        var v191 float32 = (float32(0.1) * 1.5)
                        _ = v191
                    } else {
                        var v192 uint32 = 65535
                        _ = v192
                    }
                }
                fmt.Println(v3, fi172)
            }
            fmt.Println(v1, fi172)
            v193 := GenS9{F0: 0, F1: 1, F2: 255}
            _ = fi172
            _ = v193
        }
        _ = fi171
    }
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v116
    _ = v117
    _ = v118
    _ = v119
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
