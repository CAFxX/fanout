package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 uint16; F1 int32; F2 int64 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 uint; F1 string; F2 string }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 uint8; F1 bool }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 string }
func genFn0(p0 uint) int32 {
    var v5 float64 = -2.25
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 int8 = -117
        switch ("foo bar" + "") {
        case "hello":
            v8 := []int{}
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    switch "foo bar" {
    case "a":
        fmt.Println(p0, p0)
    default:
        switch ("foo bar" + "hello") {
        case "a":
            var v9 float32 = 1.5
            _ = v9
        }
    }
    var v10 int = -25
    _ = p0
    _ = v10
    _ = v5
    return 129
}
func genFn1(p0 uint32, p1 uint) int16 {
    p0 -= 980
    v11 := GenS1{F0: "a", F1: 0, F2: 1}
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(p0, p1)
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v11
    return 0
}

func main() {
    var v1 bool = false
    var v2 bool = (!true)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    var v13 string = "foo bar"
    for fi14 := 0; fi14 < 13; fi14++ {
        v13 = ("a" + v13)
        v15 := []int{(fi14 ^ 0), (fi14 - (fi14 << ((1) & 63))), fi14}
        if (!(((v13 + "\n\t") + ((v13 + v13) + v13)) == (v13 + v13))) {
            switch 2147483647 {
            case 129:
                v4["hello"] = 0
            case -968:
                for fi16 := 0; fi16 < 18; fi16++ {
                    if (100 > -77) {
                        v13 = v13
                        switch (^(fi16)) {
                        case 9223372036854775807:
                            switch (("" + v13) + (v13 + v13)) {
                            case "a":
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                            case "foo bar":
                                fmt.Println(fi16)
                            case "hello":
                                switch ((v13 + v13) + (v13 + v13)) {
                                case "a":
                                    var v17 float64 = 0.1
                                    _ = v17
                                case "foo bar":
                                    v13 = v13
                                case "\n\t":
                                    var v18 int = fi16
                                    _ = v18
                                default:
                                    switch fi14 {
                                    case 32768:
                                        v1 = (v1 && v2)
                                    }
                                    if false {
                                        switch "a" {
                                        case "":
                                            var v19 int64 = 382
                                            _ = v19
                                        case "foo bar":
                                            var v20 bool = v2
                                            _ = v20
                                        default:
                                            if (459 < 150) {
                                                v13 = (("\n\t" + "\n\t") + (v13 + "foo bar"))
                                                var v21 uint8 = 2
                                                v2 = v1
                                                _ = v21
                                            }
                                        }
                                        var v22 uint8 = 136
                                        _ = v22
                                    } else {
                                        if len(v15) > 0 {
                                            v15[0] = fi16
                                        }
                                        v23 := GenS1{F0: v13, F1: 0, F2: 1}
                                        _ = len(v15)
                                        _ = v23
                                    }
                                    if v1 {
                                        switch fi16 {
                                        case 592:
                                            if v1 {
                                                fmt.Println(fi14, v1, v13)
                                                _ = len(v15)
                                                v24 := GenS1{F0: v13, F1: 32767, F2: 255}
                                                _ = v24
                                            }
                                        case 42:
                                            if v1 {
                                                switch 9223372036854775807 {
                                                case 705:
                                                    if len(v15) > 0 {
                                                        v15[0] = fi14
                                                    }
                                                case 0:
                                                    v25 := GenS1{F0: (v13 + (("\n\t" + "\n\t") + "a")), F1: 2, F2: 85}
                                                    _ = v25
                                                default:
                                                    v26 := GenS2{F0: genFn1(614, 42), F1: 242}
                                                    _ = v26
                                                }
                                                fmt.Println(fi16, fi16)
                                            }
                                        default:
                                            v27 := GenS2{F0: 0, F1: 2}
                                            _ = v27
                                        }
                                    } else {
                                        fmt.Println(v1, fi14, v13)
                                        v28 := GenS3{F0: 42, F1: 42}
                                        _ = v28
                                    }
                                    v29 := GenS4{F0: 237, F1: 32767}
                                    var v30 float64 = (((0.0 - 1e10) - 0.0) + 1e10)
                                    var v31 uint64 = 1
                                    v2 = v2
                                    if v1 {
                                        fmt.Println(v2, v1)
                                        if true {
                                            if (18446744073709551615 <= 888) {
                                                switch -940 {
                                                case 4294967295:
                                                    if v2 {
                                                        fmt.Println(v2, v1)
                                                        v32 := GenS2{F0: 1, F1: 0}
                                                        v31 = (^(v31))
                                                        var v33 uint16 = 2
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                case 32768:
                                                    if ((42 > 9223372036854775807) || (119 <= 128)) {
                                                        var v34 bool = (!((1 <= 326) || (!(!v1))))
                                                        if (!(1 < 233)) {
                                                            v34 = (!v1)
                                                            var v35 int64 = 9223372036854775807
                                                            var v36 int32 = 42
                                                            _ = v35
                                                            _ = v36
                                                        } else {
                                                            if (1 > -1) {
                                                                if (252 == (v31 - 604)) {
                                                                    v37 := GenS5{F0: 584, F1: -242, F2: 9223372036854775807}
                                                                    var v38 float32 = (0.0 + 1.5)
                                                                    fmt.Println(v30, v38, v31)
                                                                    _ = v37
                                                                    _ = v38
                                                                }
                                                                var v39 float32 = float32(float64(float32(((((1.5 / float64(1e10)) / 3.14159) + (((3.14159 * (v30 / v30)) * v30) * float64(float32((v30 - 1e10))))) + v30))))
                                                                _ = v39
                                                            } else {
                                                                var v40 uint64 = v31
                                                                _ = v40
                                                            }
                                                        }
                                                        fmt.Println(fi14, v34, v34)
                                                        _ = v34
                                                    }
                                                }
                                                v31 -= v31
                                                var v41 uint8 = 127
                                                _ = v41
                                            } else {
                                                v42 := GenS6{F0: -884, F1: v2}
                                                if len(v15) > 0 {
                                                    v15[0] = (fi14 | -9223372036854775808)
                                                }
                                                _ = v42
                                            }
                                        } else {
                                            v13 = ("\n\t" + v13)
                                            v1 = v1
                                            v1 = false
                                        }
                                    } else {
                                        fmt.Println(v1, v13, fi16)
                                        v1 = true
                                        var v43 int = 463
                                        _ = v43
                                    }
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                            default:
                                var v44 uint32 = 0
                                if (!false) {
                                    var v45 int = 127
                                    var v46 bool = true
                                    v15 = append(v15, fi14)
                                    fmt.Println(fi14, v1, v1)
                                    v13 = v13
                                    _ = v45
                                    _ = v46
                                }
                                switch (v13 + (v13 + "")) {
                                case "foo bar":
                                    v47 := GenS7{F0: 255, F1: v13, F2: ""}
                                    _ = v47
                                case "":
                                    var v48 int = (fi16 >> ((65535) & 63))
                                    if v1 {
                                        v44 -= (v44 | (v44 + (v44 * (v44 * v44))))
                                        fmt.Println(v44, v48)
                                        _ = len(v15)
                                    }
                                    _ = v48
                                case "\n\t":
                                    var v49 uint32 = 0
                                    _ = v49
                                default:
                                    var v50 int64 = 0
                                    _ = v50
                                }
                                _ = v44
                            }
                        default:
                            if true {
                                v13 = v13
                                v51 := v4[v13]
                                fmt.Println(v2, fi14, v1)
                                _ = v51
                            }
                        }
                        switch 1 {
                        case 32768:
                            if ((("hello" + v13) < "") || (0 <= 65535)) {
                                switch 813 {
                                case 0:
                                    v1 = (128 > 38)
                                default:
                                    v2 = true
                                }
                                v2 = v1
                            } else {
                                v52 := GenS5{F0: 42, F1: 172, F2: -472}
                                v53 := GenS4{F0: 42, F1: (fi14 | fi14)}
                                v54 := GenS3{F0: 0, F1: 0}
                                _ = v52
                                _ = v53
                                _ = v54
                            }
                        }
                    }
                    v13 = ((v13 + v13) + v13)
                    _ = fi16
                }
            case 467:
                v2 = (4294967295 <= 837)
            }
        } else {
            fmt.Println(fi14)
            for fi55 := 0; fi55 < 1; fi55++ {
                v2 = (-128 == -9)
                if len(v15) > 0 {
                    v15[0] = (fi55 | 4294967295)
                }
                v2 = (!v1)
                v2 = true
                fmt.Println(v13)
                fmt.Println(fi14, fi55)
                _ = fi55
            }
        }
        _ = fi14
        _ = v15
    }
    switch "hello" {
    case "foo bar":
        var v56 uint32 = 248
        _ = v56
    case "hello":
        var v57 float64 = 0.0
        _ = v57
    default:
        fmt.Println(v2)
    }
    for fi58 := 0; fi58 < 14; fi58++ {
        switch ((v13 + v13) + v13) {
        case "\n\t":
            v2 = v1
        default:
            for fi59 := 0; fi59 < 19; fi59++ {
                if (2147483648 > 0) {
                    if (((1.5 + 3.14159) / (((1e10 + 0.0) * (float64(3.14159) / 1.0)) + -2.25)) < 0.1) {
                        v1 = (-5 != 127)
                    } else {
                        v60 := GenS5{F0: 430, F1: -2147483648, F2: 0}
                        v1 = v1
                        _ = v60
                    }
                    if ((!true) || (false || v2)) {
                        v2 = (((0.1 / ((0.0 * 1e10) + 0.1)) + (0.1 / 3.14159)) < 1e10)
                        if v1 {
                            v2 = ((!true) || (1 >= 32767))
                            v61 := v3[2]
                            switch (v13 + "foo bar") {
                            case "hello":
                                v13 = ((v13 + (v13 + "hello")) + ("\n\t" + v13))
                            }
                            _ = v61
                        }
                        var v62 uint8 = 35
                        if (!v1) {
                            v63 := GenS5{F0: 32767, F1: 1, F2: 690}
                            fmt.Println(v1)
                            var v64 int16 = 42
                            _ = v63
                            _ = v64
                        }
                        _ = v62
                    }
                    var v65 int32 = 2147483647
                    _ = v65
                } else {
                    delete(v4, v13)
                    v66 := []int{(fi59 % ((-912) | 1)), fi58}
                    _ = v66
                }
                v67 := GenS8{F0: v1, F1: fi58}
                v68 := v4[v13]
                _ = fi59
                _ = v67
                _ = v68
            }
        }
        v1 = ((v13 + ((v13 + v13) + "a")) == "\n\t")
        for fi69 := 0; fi69 < 19; fi69++ {
            switch (-(fi69)) {
            case 202:
                v70 := GenS1{F0: "\n\t", F1: 779, F2: 0}
                _ = v70
            case -1:
                v2 = (!v2)
            default:
                switch (fi69 * fi58) {
                case 127:
                    v4["a"] = 1
                case 485:
                    switch "a" {
                    case "":
                        delete(v4, ("hello" + (v13 + ("hello" + ("a" + (v13 + v13))))))
                    case "foo bar":
                        var v71 string = "foo bar"
                        _ = v71
                    default:
                        v1 = (-1 <= 1)
                    }
                case -9223372036854775808:
                    if (!(178 <= 1)) {
                        if v2 {
                            if (740 > 42) {
                                if v1 {
                                    if v1 {
                                        fmt.Println(v1, v2, fi69)
                                        if v2 {
                                            v13 = v13
                                            v72 := []int{4294967295, (fi58 % ((fi58) | 1)), 32767, 42}
                                            _ = v72
                                        }
                                        var v73 float32 = ((((1e10 * 1e10) * float32(float64(3.14159))) - -2.25) - 1e10)
                                        _ = v73
                                    } else {
                                        if (((3.14159 / float32((float64((1.5 + 1e10)) * (((1.5 / 1.0) * (1.5 - 0.0)) - float64(1.5))))) == (1.5 / ((1.5 / 3.14159) + float32(0.1)))) && ((1.5 / 3.14159) <= 1e10)) {
                                            switch fi69 {
                                            case 1:
                                                v74 := []int{128, (fi69 * fi69), (fi69 << ((0) & 63)), (fi58 * (fi58 >> ((517) & 63)))}
                                                _ = v74
                                            case -209:
                                                v2 = v1
                                            case 219:
                                                v75 := []int{}
                                                _ = v75
                                            }
                                            v13 = v13
                                            v76 := []int{fi58}
                                            _ = v76
                                        } else {
                                            v13 = v13
                                        }
                                        v77 := GenS3{F0: 758, F1: 152}
                                        fmt.Println(v2, v1, fi69)
                                        _ = v77
                                    }
                                }
                                fmt.Println(v2, fi69, fi58)
                                var v78 uint8 = 42
                                _ = v78
                            } else {
                                v2 = (4294967295 > 177)
                                v79 := GenS2{F0: 0, F1: 42}
                                _ = v79
                            }
                            v80 := GenS3{F0: -2147483648, F1: -9223372036854775808}
                            _ = v80
                        } else {
                            v81 := GenS7{F0: 0, F1: v13, F2: v13}
                            _ = v81
                        }
                        v82 := GenS1{F0: "hello", F1: 247, F2: 2}
                        v13 = (((v13 + ("hello" + (v13 + ((((v13 + v13) + v13) + v13) + v13)))) + "a") + (("" + v13) + v13))
                        v1 = (42 != 292)
                        v2 = v2
                        _ = v82
                    }
                default:
                    switch 128 {
                    case 2147483647:
                        v83 := []int{}
                        switch (fi69 & (fi69 & 9223372036854775807)) {
                        case 0:
                            switch v13 {
                            case "\n\t":
                                var v84 float32 = -2.25
                                _ = v84
                            }
                        case -2147483648:
                            if v1 {
                                var v85 float64 = -2.25
                                var v86 uint16 = 333
                                if (v2 && false) {
                                    var v87 string = v13
                                    var v88 int = fi58
                                    v89 := v4[((v13 + v87) + v87)]
                                    _ = v87
                                    _ = v88
                                    _ = v89
                                } else {
                                    var v90 uint32 = 1
                                    switch fi69 {
                                    case -80:
                                        v85 = (1e10 * (1.5 - ((3.14159 + 0.0) - v85)))
                                    case 437:
                                        v91 := GenS9{F0: 42, F1: true}
                                        _ = v91
                                    }
                                    var v92 uint = 65535
                                    if len(v83) > 0 {
                                        v83[0] = (^(fi69))
                                    }
                                    if ((v13 + "\n\t") <= v13) {
                                        v93 := GenS8{F0: v2, F1: (fi58 >> ((v92) & 63))}
                                        v2 = false
                                        v90 -= 42
                                        _ = v93
                                    }
                                    _ = v90
                                    _ = v92
                                }
                                _ = v85
                                _ = v86
                            }
                        default:
                            var v94 uint32 = 2147483647
                            _ = v94
                        }
                        _ = v83
                    case 42:
                        v95 := GenS10{F0: 2}
                        _ = v95
                    }
                }
            }
            if true {
                if true {
                    v3[914] = 182
                    fmt.Println(v2)
                }
            }
            if false {
                if (18446744073709551615 < 2) {
                    fmt.Println(v13, v1, fi69)
                } else {
                    var v96 int = 32767
                    var v97 int32 = 2147483647
                    v98 := []int{(v96 % ((fi58) | 1))}
                    var v99 string = (v13 + v13)
                    v100 := GenS10{F0: 65535}
                    _ = v100
                    _ = v96
                    _ = v97
                    _ = v98
                    _ = v99
                }
                fmt.Println(fi69, v2)
                fmt.Println(fi58, v2)
            }
            _ = fi69
        }
        _ = fi58
    }
    if v2 {
        v1 = (1 > -32768)
        if v2 {
            if (0 == 42) {
                fmt.Println(v13, v2, v2)
                for fi101 := 0; fi101 < 9; fi101++ {
                    if (127 <= 0) {
                        v13 = ("a" + (v13 + v13))
                        var v102 string = "a"
                        _ = v102
                    }
                    v2 = ((32768 == fi101) || v2)
                    var v103 uint64 = 280
                    _ = fi101
                    _ = v103
                }
            } else {
                v2 = true
                fmt.Println(v2)
            }
            var v104 int16 = 32767
            _ = v104
        } else {
            for fi105 := 0; fi105 < 8; fi105++ {
                fmt.Println(fi105, fi105)
                fmt.Println(v1, v2)
                _ = fi105
            }
            if (!(!true)) {
                if (float64(-2.25) <= (-2.25 / 1.5)) {
                    fmt.Println(v13, v13)
                } else {
                    var v106 string = v13
                    for fi107 := 0; fi107 < 11; fi107++ {
                        var v108 int32 = 2147483647
                        switch v106 {
                        case "":
                            fmt.Println(v2, fi107)
                        default:
                            switch (fi107 ^ fi107) {
                            case 2147483647:
                                for fi109 := 0; fi109 < 6; fi109++ {
                                    v106 = ("foo bar" + "hello")
                                    var v110 uint8 = 255
                                    fmt.Println(v1, v2, fi109)
                                    var v111 int32 = genFn0(1)
                                    var v112 float32 = 1e10
                                    var v113 int32 = (v111 & v111)
                                    var v114 int64 = -853
                                    _ = fi109
                                    _ = v110
                                    _ = v111
                                    _ = v112
                                    _ = v113
                                    _ = v114
                                }
                            }
                        }
                        _ = fi107
                        _ = v108
                    }
                    fmt.Println(v2, v2, v2)
                    _ = v106
                }
                v1 = (-1 >= 42)
            }
            v115 := GenS1{F0: v13, F1: 65535, F2: 219}
            _ = v115
        }
    }
    switch "a" {
    case "foo bar":
        v3[2] = 2147483647
    case "\n\t":
        var v116 int32 = -2147483648
        _ = v116
    case "a":
        v117 := GenS4{F0: 65535, F1: 0}
        _ = v117
    }
    for fi118 := 0; fi118 < 9; fi118++ {
        var v119 uint = 4294967295
        if (v1 && (-1 == genFn0(42))) {
            v119 -= (^(v119))
            if (939 < 2147483648) {
                for fi120 := 0; fi120 < 18; fi120++ {
                    fmt.Println(v13)
                    if false {
                        v121 := GenS11{F0: 2}
                        _ = v121
                    }
                    _ = fi120
                }
            } else {
                for fi122 := 0; fi122 < 9; fi122++ {
                    var v123 bool = v1
                    v124 := []int{fi118, fi122}
                    _ = fi122
                    _ = v123
                    _ = v124
                }
            }
        } else {
            switch ((v13 + ("a" + v13)) + (("foo bar" + (v13 + "")) + v13)) {
            case "hello":
                v125 := GenS5{F0: 65535, F1: 1, F2: 9223372036854775807}
                _ = v125
            case "a":
                var v126 uint8 = 127
                _ = v126
            case "\n\t":
                fmt.Println(v13)
            }
        }
        _ = fi118
        _ = v119
    }
    var v127 int32 = -2147483648
    var v128 int64 = -9223372036854775808
    var v129 int64 = v128
    switch v13 {
    case "":
        if ((-(v127)) == v127) {
            var v130 uint64 = 2
            v1 = (false || (!false))
            v127 = (v127 & (v127 | (-(v127))))
            _ = v130
        }
    case "a":
        v131 := v3[42]
        _ = v131
    }
    v3[2147483647] = 496
    if ((42 <= 100) || (v129 > (v129 ^ (v128 >> ((18446744073709551615) & 63))))) {
        v127 -= v127
    }
    for fi132 := 0; fi132 < 7; fi132++ {
        switch "" {
        case "a":
            for fi133 := 0; fi133 < 19; fi133++ {
                switch fi132 {
                case -1:
                    v2 = (!v1)
                    if (v1 || ((!(18446744073709551615 >= 0)) && (((-2.25 * ((-2.25 * ((float32(3.14159) - 3.14159) / 3.14159)) * -2.25)) / 3.14159) <= (0.1 * 0.0)))) {
                        v134 := GenS12{F0: 42}
                        switch "" {
                        case "\n\t":
                            var v135 float32 = (float32(1.5) * 1.5)
                            _ = v135
                        }
                        _ = v134
                    }
                    delete(v3, 219)
                case 9223372036854775807:
                    if v2 {
                        var v136 uint = 18446744073709551615
                        v137 := GenS3{F0: (v127 + v127), F1: 42}
                        fmt.Println(v13)
                        _ = v136
                        _ = v137
                    }
                case 128:
                    if (615 >= 525) {
                        v138 := GenS13{F0: v13}
                        v128 += 9223372036854775807
                        v139 := GenS5{F0: 680, F1: genFn0(0), F2: v129}
                        _ = v138
                        _ = v139
                    }
                default:
                    v1 = v1
                }
                v128 = 216
                v140 := GenS5{F0: 302, F1: (v127 << ((60) & 63)), F2: v128}
                _ = fi133
                _ = v140
            }
        case "\n\t":
            v141 := []int{}
            _ = v141
        default:
            var v142 float32 = ((1e10 * (0.0 + 1e10)) * float32(float64(1e10)))
            _ = v142
        }
        v143 := GenS5{F0: 1, F1: (v127 ^ v127), F2: (v129 << ((255) & 63))}
        _ = fi132
        _ = v143
    }
    var v144 uint8 = 1
    v145 := v3[143]
    var v146 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v127:", v127)
    fmt.Println("v128:", v128)
    fmt.Println("v129:", v129)
    fmt.Println("v13:", v13)
    fmt.Println("v144:", v144)
    fmt.Println("v145:", v145)
    fmt.Println("v146:", v146)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v127
    _ = v128
    _ = v129
    _ = v13
    _ = v144
    _ = v145
    _ = v146
    _ = v2
    _ = v3
    _ = v4
}
