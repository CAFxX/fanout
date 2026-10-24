package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8; F1 int64 }
type GenS8 struct { F0 int8; F1 int; F2 uint16 }
func genFn0(p0 uint, p1 int32) int64 {
    fmt.Println(p0)
    if (!false) {
        fmt.Println(p1, p1)
        var v6 float32 = 3.14159
        var v7 int8 = 127
        _ = v6
        _ = v7
    }
    v8 := []int{-9223372036854775808}
    var v9 int64 = -9223372036854775808
    var v10 float64 = ((0.1 / 3.14159) + (1.5 * -2.25))
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return -449
}
func genFn1() int16 {
    switch 0 {
    case -756:
        v11 := GenS1{F0: 2}
        _ = v11
    }
    var v12 float32 = float32(((1e10 - (float64(1.5) * 1.5)) + 3.14159))
    v12 += (v12 * (0.1 / (0.0 - 0.1)))
    fmt.Println(v12)
    v13 := []int{0}
    _ = v12
    _ = v13
    return 0
}

func main() {
    v1 := []int{-793, -2147483648, 0, -1}
    var v2 uint8 = 255
    var v3 bool = (!((!false) && (2 <= (^(v2)))))
    v4 := []int{47, -1, 1}
    var v5 uint16 = 42
    _ = len(v4)
    var v14 int8 = 100
    v1 = append(v1, -1)
    switch "foo bar" {
    case "a":
        if v3 {
            if v3 {
                v15 := GenS1{F0: 0}
                var v16 uint32 = 65535
                for fi17 := 0; fi17 < 17; fi17++ {
                    if ((int16(fi17) >= 0) && (!true)) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = len(v4)
                    }
                    var v18 uint64 = 2
                    _ = fi17
                    _ = v18
                }
                _ = v15
                _ = v16
            }
            for fi19 := 0; fi19 < 11; fi19++ {
                if (v3 && (!(4294967295 >= 4294967295))) {
                    for fi20 := 0; fi20 < 8; fi20++ {
                        v14 = v14
                        if false {
                            var v21 int32 = -1
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v21
                        }
                        _ = fi20
                    }
                    var v22 uint16 = v5
                    switch "" {
                    case "hello":
                        v3 = true
                    case "":
                        v1 = append(v1, fi19)
                    case "foo bar":
                        v23 := GenS1{F0: 1}
                        _ = v23
                    }
                    _ = v22
                } else {
                    switch (fi19 ^ fi19) {
                    case -888:
                        for fi24 := 0; fi24 < 17; fi24++ {
                            v25 := GenS2{F0: (857 == 323)}
                            _ = fi24
                            _ = v25
                        }
                    default:
                        var v26 string = (((("hello" + "hello") + "foo bar") + "") + ("a" + "\n\t"))
                        _ = v26
                    }
                }
                _ = fi19
            }
            for fi27 := 0; fi27 < 7; fi27++ {
                switch (fi27 << ((2) & 63)) {
                case -679:
                    var v28 int8 = 1
                    _ = v28
                case 127:
                    switch (("a" + ("hello" + "foo bar")) + "\n\t") {
                    case "":
                        var v29 int = (fi27 & 157)
                        _ = v29
                    default:
                        v30 := GenS1{F0: 2}
                        _ = v30
                    }
                default:
                    if len(v4) > 0 {
                        v4[0] = (fi27 / (((fi27 | int(v5))) | 1))
                    }
                }
                v31 := GenS1{F0: 0}
                switch (("\n\t" + "a") + "foo bar") {
                case "foo bar":
                    for fi32 := 0; fi32 < 4; fi32++ {
                        v5 -= 32767
                        var v33 uint64 = 42
                        var v34 string = "\n\t"
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                case "":
                    v35 := GenS1{F0: 2}
                    _ = v35
                case "a":
                    if (1 != 1) {
                        for fi36 := 0; fi36 < 16; fi36++ {
                            var v37 uint = 286
                            _ = fi36
                            _ = v37
                        }
                        if (!(!v3)) {
                            switch fi27 {
                            case 4294967295:
                                var v38 bool = v3
                                _ = v38
                            case -753:
                                var v39 string = "foo bar"
                                _ = v39
                            default:
                                for fi40 := 0; fi40 < 8; fi40++ {
                                    v2 -= (v2 >> ((255) & 63))
                                    _ = fi40
                                }
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                default:
                    if false {
                        var v41 uint32 = 4294967295
                        if len(v1) > 0 {
                            v1[0] = fi27
                        }
                        for fi42 := 0; fi42 < 9; fi42++ {
                            v1 = append(v1, fi27)
                            if (v3 && ((fi42 - 128) == fi27)) {
                                v1 = append(v1, fi27)
                                var v43 int = fi27
                                v3 = v3
                                _ = v43
                            }
                            _ = fi42
                        }
                        _ = v41
                    }
                }
                _ = fi27
                _ = v31
            }
        } else {
            if ((0.1 != float64((1e10 / 1.0))) && true) {
                v44 := GenS2{F0: (uint16(v2) >= 65535)}
                _ = v44
            } else {
                if (0.0 == (3.14159 / 1e10)) {
                    if true {
                        if len(v1) > 0 {
                            v1[0] = 355
                        }
                        var v45 uint = 65535
                        _ = v45
                    }
                    fmt.Println(v3)
                }
                for fi46 := 0; fi46 < 1; fi46++ {
                    var v47 float64 = float64(-2.25)
                    fmt.Println(v5)
                    fmt.Println(v3, v47)
                    _ = fi46
                    _ = v47
                }
                v48 := GenS1{F0: 68}
                _ = v48
            }
            _ = len(v4)
            v49 := GenS2{F0: (!(127 <= 22))}
            _ = v49
        }
    default:
        if v3 {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v5 += v5
            v1 = append(v1, -352)
        }
    }
    fmt.Println(v5, v3, v14)
    var v50 float32 = (1.5 * 0.0)
    var v51 int8 = 42
    if ((^(v2)) > uint8(v5)) {
        fmt.Println(v14, v51)
    }
    switch 646 {
    case 32767:
        var v52 uint64 = 923
        _ = v52
    case -9223372036854775808:
        for fi53 := 0; fi53 < 15; fi53++ {
            var v54 uint64 = 2
            v55 := GenS1{F0: v54}
            _ = fi53
            _ = v54
            _ = v55
        }
    default:
        v56 := GenS2{F0: v3}
        _ = v56
    }
    if ("hello" != "foo bar") {
        v57 := GenS2{F0: (!true)}
        _ = v57
    } else {
        var v58 float64 = (1.5 + 3.14159)
        if (127 <= 8) {
            if v3 {
                v59 := GenS3{F0: 0}
                var v60 uint16 = v5
                if v3 {
                    switch 199 {
                    case 32768:
                        switch "foo bar" {
                        case "":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "\n\t":
                            fmt.Println(v60, v14, v50)
                        default:
                            switch ("foo bar" + (("foo bar" + "\n\t") + "\n\t")) {
                            case "":
                                switch 32768 {
                                case 32767:
                                    var v61 float32 = ((-2.25 - v50) - (float32((v58 / v58)) + v50))
                                    _ = v61
                                default:
                                    if ((v3 || (!false)) && false) {
                                        if ((v50 + (v50 / 1e10)) >= v50) {
                                            var v62 int16 = 0
                                            v63 := GenS4{F0: v2, F1: false}
                                            _ = v62
                                            _ = v63
                                        } else {
                                            if false {
                                                for fi64 := 0; fi64 < 14; fi64++ {
                                                    if false {
                                                        var v65 uint = 381
                                                        switch ("\n\t" + "") {
                                                        case "foo bar":
                                                            v58 -= ((v58 * v58) + float64(3.14159))
                                                        }
                                                        _ = v65
                                                    }
                                                    if v3 {
                                                        v66 := GenS4{F0: v2, F1: ((-(fi64)) != -1)}
                                                        _ = v66
                                                    }
                                                    v14 -= (v14 - (v14 % (((v14 | 100)) | 1)))
                                                    _ = fi64
                                                }
                                                if ((-857 >= -1) && v3) {
                                                    v60 -= (^(v5))
                                                    switch 2147483647 {
                                                    case 32767:
                                                        for fi67 := 0; fi67 < 20; fi67++ {
                                                            v68 := GenS2{F0: (237 != 1)}
                                                            v69 := GenS4{F0: (^(v2)), F1: v3}
                                                            switch "hello" {
                                                            case "hello":
                                                                var v70 uint32 = uint32(v2)
                                                                _ = v70
                                                            case "\n\t":
                                                                var v71 int16 = 886
                                                                _ = v71
                                                            default:
                                                                var v72 uint32 = 2147483648
                                                                _ = v72
                                                            }
                                                            _ = fi67
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                    }
                                                } else {
                                                    for fi73 := 0; fi73 < 7; fi73++ {
                                                        var v74 int64 = -1
                                                        var v75 int16 = 0
                                                        switch fi73 {
                                                        case 42:
                                                            var v76 float64 = (3.14159 + v58)
                                                            _ = v76
                                                        case 927:
                                                            v77 := GenS5{F0: 0, F1: 42}
                                                            _ = v77
                                                        }
                                                        _ = fi73
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                    switch ("" + "foo bar") {
                                                    case "a":
                                                        for fi78 := 0; fi78 < 18; fi78++ {
                                                            if len(v1) > 0 {
                                                                v1[0] = 32767
                                                            }
                                                            var v79 uint = 2
                                                            _ = fi78
                                                            _ = v79
                                                        }
                                                    case "":
                                                        v60 = v5
                                                    }
                                                }
                                                v14 += v51
                                            } else {
                                                var v80 uint32 = 667
                                                switch -9223372036854775808 {
                                                case -9223372036854775808:
                                                    switch -9223372036854775808 {
                                                    case 42:
                                                        fmt.Println(v51, v50)
                                                    }
                                                }
                                                for fi81 := 0; fi81 < 8; fi81++ {
                                                    v60 -= (v5 - uint16(v80))
                                                    var v82 uint8 = (v2 % ((v2) | 1))
                                                    switch "" {
                                                    case "a":
                                                        for fi83 := 0; fi83 < 13; fi83++ {
                                                            v60 = v60
                                                            if v3 {
                                                                fmt.Println(v50, v2, v50)
                                                                var v84 int32 = -2147483648
                                                                var v85 uint32 = 2
                                                                _ = v84
                                                                _ = v85
                                                            } else {
                                                                var v86 string = "hello"
                                                                v50 += 1.5
                                                                fmt.Println(v60, v86)
                                                                _ = v86
                                                            }
                                                            var v87 uint8 = 127
                                                            if v3 {
                                                                v60 = 740
                                                                var v88 int32 = -2147483648
                                                                _ = v88
                                                            } else {
                                                                var v89 uint32 = v80
                                                                v90 := GenS2{F0: false}
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                            switch "\n\t" {
                                                            case "a":
                                                                switch (("foo bar" + "a") + "") {
                                                                case "":
                                                                    if true {
                                                                        v82 -= 1
                                                                        v87 += (v82 - (v82 / ((1) | 1)))
                                                                    } else {
                                                                        v91 := GenS2{F0: false}
                                                                        var v92 string = "a"
                                                                        _ = v91
                                                                        _ = v92
                                                                    }
                                                                case "\n\t":
                                                                    v93 := GenS4{F0: 127, F1: true}
                                                                    _ = v93
                                                                case "foo bar":
                                                                    if len(v4) > 0 {
                                                                        v4[0] = (-(fi81))
                                                                    }
                                                                default:
                                                                    if v3 {
                                                                        var v94 uint16 = v60
                                                                        v1 = append(v1, fi83)
                                                                        var v95 uint = 42
                                                                        v1 = append(v1, fi83)
                                                                        if v3 {
                                                                            fmt.Println(v95, v80, v94)
                                                                            v1 = append(v1, fi83)
                                                                        }
                                                                        _ = v94
                                                                        _ = v95
                                                                    }
                                                                }
                                                                var v96 int8 = (v51 << ((42) & 63))
                                                                var v97 int = 127
                                                                _ = v96
                                                                _ = v97
                                                            case "hello":
                                                                v14 -= (v14 | 42)
                                                            case "\n\t":
                                                                var v98 int64 = -1
                                                                _ = v98
                                                            default:
                                                                switch (((("hello" + ("hello" + "a")) + "") + ("hello" + ("" + "\n\t"))) + "a") {
                                                                case "\n\t":
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        var v99 uint8 = 49
                                                                        _ = v99
                                                                    default:
                                                                        v100 := GenS6{F0: fi83}
                                                                        _ = v100
                                                                    }
                                                                case "":
                                                                    var v101 uint32 = (^(v80))
                                                                    _ = v101
                                                                case "hello":
                                                                    var v102 uint32 = v80
                                                                    _ = v102
                                                                default:
                                                                    if len(v4) > 0 {
                                                                        v4[0] = -2147483648
                                                                    }
                                                                }
                                                            }
                                                            _ = fi83
                                                            _ = v87
                                                        }
                                                    case "":
                                                        var v103 uint64 = 763
                                                        _ = v103
                                                    case "\n\t":
                                                        switch 42 {
                                                        case 1000:
                                                            v3 = (v3 && (1 > 340))
                                                        case 127:
                                                            v104 := GenS7{F0: (v2 - 42), F1: 42}
                                                            _ = v104
                                                        case -695:
                                                            var v105 bool = v3
                                                            _ = v105
                                                        }
                                                    default:
                                                        v106 := GenS2{F0: (338 >= -1)}
                                                        _ = v106
                                                    }
                                                    _ = fi81
                                                    _ = v82
                                                }
                                                _ = v80
                                            }
                                            if false {
                                                v4 = append(v4, 0)
                                            } else {
                                                v60 = 255
                                            }
                                        }
                                        if v3 {
                                            switch 127 {
                                            case 42:
                                                var v107 int8 = (v51 % (((v51 | 127)) | 1))
                                                _ = v107
                                            default:
                                                for fi108 := 0; fi108 < 6; fi108++ {
                                                    var v109 float32 = (((float32(float64(float32(v58))) * 1.5) * v50) - v50)
                                                    v1 = append(v1, fi108)
                                                    if v3 {
                                                        switch "hello" {
                                                        case "a":
                                                            v2 = 0
                                                        case "hello":
                                                            var v110 float32 = 0.1
                                                            _ = v110
                                                        }
                                                        v111 := GenS7{F0: (v2 & v2), F1: 9223372036854775807}
                                                        _ = v111
                                                    } else {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        v112 := GenS3{F0: 9223372036854775807}
                                                        var v113 float32 = float32((0.0 + 0.0))
                                                        _ = v112
                                                        _ = v113
                                                    }
                                                    _ = fi108
                                                    _ = v109
                                                }
                                            }
                                            var v114 float64 = v58
                                            var v115 int16 = 193
                                            _ = v114
                                            _ = v115
                                        } else {
                                            for fi116 := 0; fi116 < 1; fi116++ {
                                                if true {
                                                    v2 = v2
                                                    if (int32(v51) > -2147483648) {
                                                        for fi117 := 0; fi117 < 18; fi117++ {
                                                            if true {
                                                                v60 = v60
                                                            }
                                                            _ = fi117
                                                        }
                                                    } else {
                                                        if ((^(v5)) >= (v60 & (v60 >> ((0) & 63)))) {
                                                            if (int8(v5) >= (v14 & (v14 + -128))) {
                                                                var v118 int32 = 42
                                                                _ = v118
                                                            }
                                                            switch "hello" {
                                                            case "a":
                                                                _ = len(v4)
                                                            case "hello":
                                                                v119 := GenS6{F0: -1}
                                                                _ = v119
                                                            case "foo bar":
                                                                var v120 int64 = 0
                                                                _ = v120
                                                            }
                                                        }
                                                        var v121 float64 = 0.0
                                                        for fi122 := 0; fi122 < 20; fi122++ {
                                                            v123 := GenS8{F0: v14, F1: (-(fi122)), F2: (v60 ^ 32767)}
                                                            var v124 float32 = float32(v58)
                                                            _ = fi122
                                                            _ = v123
                                                            _ = v124
                                                        }
                                                        _ = v121
                                                    }
                                                    v125 := GenS3{F0: -9223372036854775808}
                                                    _ = v125
                                                }
                                                if ((2 != 1) || (fi116 > fi116)) {
                                                    v126 := GenS3{F0: 0}
                                                    var v127 float32 = 1e10
                                                    _ = v126
                                                    _ = v127
                                                }
                                                _ = fi116
                                            }
                                            if v3 {
                                                var v128 uint = 18446744073709551615
                                                var v129 float64 = v58
                                                _ = v128
                                                _ = v129
                                            } else {
                                                v2 -= (v2 >> ((65535) & 63))
                                                var v130 float64 = ((float64(0.0) * 0.0) + 1e10)
                                                _ = v130
                                            }
                                            if true {
                                                var v131 int8 = (v51 >> ((18446744073709551615) & 63))
                                                var v132 float32 = v50
                                                _ = v131
                                                _ = v132
                                            }
                                        }
                                    }
                                }
                            case "a":
                                var v133 int16 = 0
                                _ = v133
                            }
                        }
                    case 127:
                        for fi134 := 0; fi134 < 4; fi134++ {
                            switch fi134 {
                            case -2147483648:
                                for fi135 := 0; fi135 < 18; fi135++ {
                                    fmt.Println(fi135, v2, v58)
                                    _ = fi135
                                }
                            default:
                                for fi136 := 0; fi136 < 19; fi136++ {
                                    switch fi136 {
                                    case 434:
                                        var v137 uint16 = v5
                                        _ = v137
                                    case 32767:
                                        var v138 uint = 4294967295
                                        _ = v138
                                    case 42:
                                        var v139 string = ("hello" + "a")
                                        _ = v139
                                    default:
                                        fmt.Println(v3, v60, v51)
                                    }
                                    fmt.Println(v14)
                                    _ = fi136
                                }
                            }
                            v3 = (!(((2147483647 <= -2147483648) || (18446744073709551615 <= 1)) && (-9223372036854775808 >= (fi134 | fi134))))
                            v51 -= v14
                            _ = fi134
                        }
                    default:
                        var v140 int8 = v14
                        _ = v140
                    }
                    if (32767 != (v5 * v60)) {
                        if (255 < 42) {
                            v58 = v58
                            if len(v1) > 0 {
                                v1[0] = 4294967295
                            }
                        }
                        v1 = append(v1, 2147483647)
                    } else {
                        var v141 int32 = 2147483647
                        v141 += v141
                        v58 = (0.0 + 0.0)
                        _ = v141
                    }
                    v60 += 0
                } else {
                    v4 = append(v4, 1)
                }
                _ = v59
                _ = v60
            }
        }
        v5 = v5
        _ = v58
    }
    if len(v1) > 0 {
        v1[0] = 32768
    }
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
}
