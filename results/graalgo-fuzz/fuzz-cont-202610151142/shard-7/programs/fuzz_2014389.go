package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 uint8; F2 string }
type GenS4 struct { F0 int64; F1 int8; F2 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS7 struct { F0 uint; F1 int64; F2 string }
type GenS8 struct { F0 int16; F1 uint8; F2 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool; F1 int32 }
type GenS11 struct { F0 uint; F1 int32 }
type GenS12 struct { F0 int32; F1 uint8; F2 bool }
type GenS13 struct { F0 uint }
func genFn0(p0 int16, p1 uint32, p2 int) int16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := []int{p2, (p2 ^ (fi6 / ((560) | 1))), (fi6 ^ (p2 / ((-1) | 1)))}
            p2 += (p2 + (fi5 & fi6))
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 19; fi8++ {
            p2 -= 413
            _ = fi8
        }
        if (!(1e10 != -2.25)) {
            p0 += -772
            v9 := []int{}
            p2 -= p2
            _ = v9
        } else {
            p1 = (p1 * (p1 >> ((2) & 63)))
            var v10 uint32 = 1
            fmt.Println(p1)
            _ = v10
        }
        _ = fi5
    }
    v11 := []int{}
    for fi12 := 0; fi12 < 5; fi12++ {
        p1 -= (p1 * p1)
        for fi13 := 0; fi13 < 14; fi13++ {
            if (-128 <= -118) {
                if (!false) {
                    p1 -= 834
                    var v14 int64 = 42
                    p1 += (p1 - (p1 >> ((18446744073709551615) & 63)))
                    switch "" {
                    case "hello":
                        var v15 uint = 832
                        if false {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            v16 := GenS1{F0: false}
                            p2 = p2
                            _ = v16
                        } else {
                            v14 = -9223372036854775808
                            var v17 float32 = (float32(0.0) + -2.25)
                            _ = v17
                        }
                        if true {
                            var v18 uint32 = p1
                            fmt.Println(p2, v14, v18)
                            _ = v18
                        }
                        _ = v15
                    case "foo bar":
                        v19 := GenS1{F0: true}
                        _ = v19
                    }
                    if (!(!false)) {
                        p1 += (p1 | 979)
                        _ = len(v11)
                        p1 += 497
                        switch (fi12 * 128) {
                        case 42:
                            p0 = (^(p0))
                        case -9223372036854775808:
                            var v20 int64 = v14
                            _ = v20
                        case 0:
                            if true {
                                v14 = -9223372036854775808
                            } else {
                                switch -1 {
                                case 1:
                                    var v21 int32 = -1
                                    _ = v21
                                }
                                v22 := GenS1{F0: (0 <= (-(p0)))}
                                var v23 uint16 = 0
                                _ = v22
                                _ = v23
                            }
                        default:
                            switch fi12 {
                            case 32767:
                                v24 := GenS1{F0: (true || (true || false))}
                                _ = v24
                            default:
                                if (p0 >= 525) {
                                    if (v14 >= (v14 * (-(v14)))) {
                                        v25 := GenS2{F0: "foo bar"}
                                        _ = v25
                                    }
                                } else {
                                    var v26 int = (fi13 >> ((1) & 63))
                                    v27 := GenS2{F0: ((("foo bar" + "\n\t") + ("\n\t" + "hello")) + "foo bar")}
                                    v28 := GenS3{F0: 1, F1: 128, F2: "a"}
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            }
                        }
                        v29 := GenS4{F0: v14, F1: 67, F2: v14}
                        _ = v29
                    }
                    fmt.Println(p2, v14, p1)
                    p1 += 65535
                    _ = v14
                }
                switch (fi13 ^ 0) {
                case -779:
                    var v30 int16 = (p0 * (p0 - int16(p1)))
                    _ = v30
                }
            }
            _ = fi13
        }
        p0 -= (p0 / ((p0) | 1))
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (p0 | (p0 >> ((890) & 63)))
}
func genFn1() uint64 {
    var v31 bool = (true && true)
    var v32 uint16 = 1
    v33 := GenS2{F0: "a"}
    if v31 {
        fmt.Println(v32)
        v31 = false
    }
    _ = v31
    _ = v32
    _ = v33
    return 4294967295
}

func main() {
    var v1 bool = false
    var v2 bool = v1
    v3 := make(map[string]string)
    v4 := []string{"", "\n\t"}
    v1 = true
    if (42 == -1) {
        _ = len(v3)
        var v34 uint = 5
        for fi35 := 0; fi35 < 19; fi35++ {
            var v36 float64 = 0.1
            var v37 int8 = -82
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = v34
    } else {
        for fi38 := 0; fi38 < 20; fi38++ {
            v39 := GenS4{F0: 9223372036854775807, F1: 1, F2: int64(fi38)}
            if v2 {
                var v40 string = ("" + ("foo bar" + "foo bar"))
                var v41 float64 = 1e10
                _ = v40
                _ = v41
            }
            switch (^(fi38)) {
            case 964:
                switch "" {
                case "foo bar":
                    v1 = ((924 != 397) && v1)
                case "\n\t":
                    v1 = false
                default:
                    if v2 {
                        v1 = (v1 || (1 == -9223372036854775808))
                        fmt.Println(fi38)
                    }
                }
            }
            _ = fi38
            _ = v39
        }
        v2 = (v1 && (v2 || (100 < 39)))
    }
    for fi42 := 0; fi42 < 8; fi42++ {
        var v43 uint64 = genFn1()
        _ = fi42
        _ = v43
    }
    v2 = ((42 == -10) || false)
    for fi44 := 0; fi44 < 12; fi44++ {
        var v45 bool = (((v2 && (false && ("\n\t" == (("" + (((("foo bar" + ("" + "hello")) + ("hello" + "foo bar")) + "hello") + "foo bar")) + "foo bar")))) && v1) && true)
        v46 := GenS1{F0: ("a" > "\n\t")}
        _ = fi44
        _ = v45
        _ = v46
    }
    if v2 {
        for fi47 := 0; fi47 < 16; fi47++ {
            var v48 uint16 = 476
            v49 := GenS1{F0: (!(42 == 410))}
            _ = fi47
            _ = v48
            _ = v49
        }
    }
    if v2 {
        if (!v1) {
            v2 = v1
        }
        v50 := GenS1{F0: v2}
        _ = v50
    } else {
        var v51 uint64 = 2
        var v52 int8 = 77
        _ = v51
        _ = v52
    }
    var v53 int8 = 42
    for fi54 := 0; fi54 < 18; fi54++ {
        v1 = (-246 <= fi54)
        v1 = (v2 && v1)
        if false {
            delete(v3, "a")
            for fi55 := 0; fi55 < 11; fi55++ {
                if true {
                    v2 = v1
                    if v2 {
                        v56 := v3[""]
                        v1 = (v2 && (v1 && (!v1)))
                        v56 = v56
                        _ = v56
                    }
                }
                _ = fi55
            }
            if v1 {
                if v1 {
                    var v57 int8 = -1
                    _ = v57
                }
                _ = len(v3)
                v58 := GenS5{F0: "a"}
                _ = v58
            }
        } else {
            for fi59 := 0; fi59 < 16; fi59++ {
                var v60 int64 = 0
                var v61 string = ""
                _ = fi59
                _ = v60
                _ = v61
            }
            v2 = v1
        }
        _ = fi54
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        if true {
            _ = len(v4)
        } else {
            switch 42 {
            case 127:
                fmt.Println(v1, v1)
            case -513:
                v53 = v53
            case -981:
                v62 := GenS4{F0: -9223372036854775808, F1: (-(v53)), F2: 42}
                _ = v62
            default:
                var v63 int32 = -2147483648
                _ = v63
            }
            fmt.Println(v53, v1)
            if ((0.0 - -2.25) <= (0.1 + 0.1)) {
                v64 := GenS1{F0: (((float32(1.5) + -2.25) / 1.5) < float32((0.0 - 1.5)))}
                v4 = append(v4, ("" + "a"))
                switch 996 {
                case -256:
                    switch 4294967295 {
                    case -2147483648:
                        var v65 float64 = 3.14159
                        _ = v65
                    default:
                        if v1 {
                            for fi66 := 0; fi66 < 19; fi66++ {
                                switch (fi66 | (fi66 % ((fi66) | 1))) {
                                case 32768:
                                    switch "\n\t" {
                                    case "":
                                        v67 := GenS4{F0: -9223372036854775808, F1: (v53 / ((v53) | 1)), F2: -401}
                                        _ = v67
                                    case "a":
                                        var v68 int = (fi66 - 128)
                                        _ = v68
                                    }
                                default:
                                    if (0 == 2) {
                                        v2 = v1
                                    }
                                }
                                v69 := GenS5{F0: "a"}
                                _ = fi66
                                _ = v69
                            }
                        }
                    }
                case 96:
                    var v70 uint8 = 0
                    _ = v70
                case 127:
                    v1 = (1.5 != (1.5 / 1.0))
                }
                _ = v64
            } else {
                for fi71 := 0; fi71 < 17; fi71++ {
                    if (1 != 18446744073709551615) {
                        fmt.Println(v53, v2)
                    }
                    v3[(((("foo bar" + ("" + ("\n\t" + "foo bar"))) + ("hello" + "\n\t")) + "hello") + "\n\t")] = "a"
                    fmt.Println(v1, v1)
                    _ = fi71
                }
            }
        }
        if v2 {
            v2 = v1
        } else {
            fmt.Println(v53, v1, v2)
            if len(v4) > 0 {
                v4[0] = "a"
            }
            for fi72 := 0; fi72 < 6; fi72++ {
                if true {
                    for fi73 := 0; fi73 < 16; fi73++ {
                        v2 = true
                        fmt.Println(v2, fi73, v1)
                        _ = fi73
                    }
                    v2 = false
                    if ((0 == 988) || (-276 != 1)) {
                        switch (fi72 - 128) {
                        case 842:
                            if len(v4) > 0 {
                                v4[0] = "\n\t"
                            }
                        case 128:
                            switch fi72 {
                            case 0:
                                v53 -= v53
                            default:
                                var v74 uint8 = 128
                                _ = v74
                            }
                        default:
                            var v75 uint8 = 2
                            _ = v75
                        }
                        if (-2147483648 >= 210) {
                            if v2 {
                                var v76 uint32 = 0
                                _ = v76
                            }
                        }
                        v77 := GenS5{F0: "hello"}
                        _ = v77
                    }
                } else {
                    for fi78 := 0; fi78 < 17; fi78++ {
                        _ = len(v3)
                        if (!(v53 < 38)) {
                            var v79 uint = 186
                            v4 = append(v4, "foo bar")
                            _ = v79
                        }
                        _ = len(v3)
                        _ = fi78
                    }
                    for fi80 := 0; fi80 < 2; fi80++ {
                        var v81 int64 = -1
                        v53 -= (v53 - int8(fi80))
                        switch "" {
                        case "":
                            var v82 uint = 36
                            _ = v82
                        case "foo bar":
                            switch ("a" + "hello") {
                            case "":
                                var v83 int8 = (v53 | (-(v53)))
                                _ = v83
                            case "foo bar":
                                var v84 float32 = -2.25
                                _ = v84
                            }
                        }
                        var v85 uint32 = 657
                        switch (fi72 - (fi72 & fi80)) {
                        case 1:
                            var v86 int64 = (-(v81))
                            _ = v86
                        }
                        v1 = ((-(v53)) > 127)
                        v85 += (v85 & (v85 >> ((255) & 63)))
                        if v2 {
                            var v87 int = 127
                            _ = v87
                        }
                        v88 := GenS6{F0: 2, F1: 65535, F2: -128}
                        fmt.Println(v53, v53, v1)
                        var v89 uint = 340
                        _ = fi80
                        _ = v81
                        _ = v85
                        _ = v88
                        _ = v89
                    }
                    for fi90 := 0; fi90 < 10; fi90++ {
                        if true {
                            var v91 int8 = v53
                            v92 := v3["a"]
                            fmt.Println(v53, v91, v2)
                            if v1 {
                                if (32767 != 2) {
                                    v93 := v3["hello"]
                                    v2 = (0.1 > 3.14159)
                                    v94 := GenS7{F0: 0, F1: -217, F2: v93}
                                    v91 = v91
                                    _ = v93
                                    _ = v94
                                } else {
                                    fmt.Println(v92)
                                    var v95 int16 = -1
                                    _ = v95
                                }
                            } else {
                                if (509 <= -313) {
                                    switch (fi90 ^ 1) {
                                    case -9223372036854775808:
                                        var v96 int16 = -613
                                        _ = v96
                                    case -575:
                                        var v97 float32 = (0.0 + float32(0.0))
                                        _ = v97
                                    case -56:
                                        var v98 uint = 18446744073709551615
                                        _ = v98
                                    default:
                                        v91 = v53
                                    }
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                if (1 >= 180) {
                                    fmt.Println(v92)
                                    v99 := GenS6{F0: 316, F1: 255, F2: (v91 - (^(v91)))}
                                    _ = v99
                                }
                                switch v92 {
                                case "":
                                    if v2 {
                                        var v100 float64 = float64(1e10)
                                        var v101 int = fi90
                                        _ = v100
                                        _ = v101
                                    } else {
                                        v102 := GenS1{F0: (26 < 42)}
                                        _ = v102
                                    }
                                case "hello":
                                    fmt.Println(v91)
                                case "\n\t":
                                    v1 = (v2 && v1)
                                }
                            }
                            _ = v91
                            _ = v92
                        }
                        _ = fi90
                    }
                }
                if v1 {
                    var v103 int8 = v53
                    _ = v103
                } else {
                    v3[("hello" + "\n\t")] = ("foo bar" + ("" + ("foo bar" + "foo bar")))
                    if len(v4) > 0 {
                        v4[0] = ("hello" + "hello")
                    }
                }
                _ = fi72
            }
        }
        switch -659 {
        case 242:
            switch 357 {
            case 42:
                if ((int64(v53) != 1) && (26 > 171)) {
                    v4 = append(v4, ("hello" + "hello"))
                    var v104 uint = 66
                    _ = v104
                } else {
                    v2 = v1
                    var v105 int32 = -914
                    _ = v105
                }
            case 32767:
                switch 860 {
                case 42:
                    v106 := GenS8{F0: 1, F1: 0, F2: (v53 >> ((1) & 63))}
                    _ = v106
                case 28:
                    for fi107 := 0; fi107 < 8; fi107++ {
                        _ = len(v3)
                        for fi108 := 0; fi108 < 8; fi108++ {
                            switch fi108 {
                            case 0:
                                if (!v2) {
                                    if v1 {
                                        v109 := GenS1{F0: (42 != 1)}
                                        var v110 int32 = 0
                                        switch ("\n\t" + "foo bar") {
                                        case "\n\t":
                                            _ = len(v4)
                                            var v111 int32 = v110
                                            _ = v111
                                        case "hello":
                                            v112 := v3[((("hello" + ("" + ("\n\t" + "foo bar"))) + "\n\t") + "foo bar")]
                                            _ = v112
                                        case "":
                                            var v113 float64 = ((-2.25 + ((1.5 + 1.5) - float64(float32(1.5)))) + 1.5)
                                            _ = v113
                                        default:
                                            fmt.Println(v2, fi108)
                                            switch (fi107 >> ((1) & 63)) {
                                            case 9223372036854775807:
                                                var v114 uint16 = 2
                                                _ = v114
                                            case -2147483648:
                                                var v115 uint8 = 0
                                                v116 := GenS9{F0: 91}
                                                _ = v115
                                                _ = v116
                                            case -990:
                                                var v117 int64 = 9223372036854775807
                                                _ = v117
                                            default:
                                                v4 = append(v4, "\n\t")
                                            }
                                            switch ("\n\t" + "hello") {
                                            case "":
                                                v1 = (385 >= 4294967295)
                                            }
                                        }
                                        _ = v109
                                        _ = v110
                                    }
                                }
                            default:
                                var v118 uint8 = 2
                                _ = v118
                            }
                            switch -340 {
                            case 42:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case 588:
                                v1 = (3.14159 >= 1e10)
                                var v119 bool = (!v1)
                                _ = v119
                            }
                            _ = fi108
                        }
                        if (!v2) {
                            delete(v3, "\n\t")
                            for fi120 := 0; fi120 < 12; fi120++ {
                                var v121 uint8 = 255
                                var v122 uint32 = 1
                                var v123 float64 = ((-2.25 + 3.14159) - 0.0)
                                _ = fi120
                                _ = v121
                                _ = v122
                                _ = v123
                            }
                            for fi124 := 0; fi124 < 6; fi124++ {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v125 := GenS10{F0: (264 > 716), F1: 407}
                                _ = fi124
                                _ = v125
                            }
                        } else {
                            for fi126 := 0; fi126 < 20; fi126++ {
                                v53 -= v53
                                _ = fi126
                            }
                            for fi127 := 0; fi127 < 17; fi127++ {
                                fmt.Println(fi127, v2, fi107)
                                fmt.Println(v53, fi107, v2)
                                var v128 uint8 = 1
                                var v129 uint32 = 146
                                if v2 {
                                    v128 -= 127
                                    var v130 string = ("hello" + "hello")
                                    _ = v130
                                }
                                var v131 bool = ((v53 * v53) < (v53 * (v53 / (((v53 * (-(v53)))) | 1))))
                                _ = fi127
                                _ = v128
                                _ = v129
                                _ = v131
                            }
                            var v132 uint16 = 65535
                            _ = v132
                        }
                        _ = fi107
                    }
                case 127:
                    v133 := v3["foo bar"]
                    _ = v133
                default:
                    v53 = (^(v53))
                }
            case 4294967295:
                for fi134 := 0; fi134 < 18; fi134++ {
                    for fi135 := 0; fi135 < 16; fi135++ {
                        if (!((18446744073709551615 == 607) || (2 < 0))) {
                            switch "" {
                            case "a":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case "\n\t":
                                v2 = v1
                                if (!(!(!(593 == 0)))) {
                                    var v136 uint16 = 2
                                    _ = v136
                                }
                                var v137 uint64 = 18446744073709551615
                                _ = v137
                            case "hello":
                                if v2 {
                                    v2 = (0 <= 260)
                                }
                            }
                        } else {
                            v138 := GenS1{F0: ((0.0 * 0.1) <= (1e10 * 1.5))}
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v2 = v2
                            _ = v138
                        }
                        _ = len(v3)
                        if (!false) {
                            if len(v4) > 0 {
                                v4[0] = "hello"
                            }
                            fmt.Println(v53, fi134, fi135)
                            if v1 {
                                fmt.Println(v2, v53)
                                switch fi134 {
                                case 42:
                                    v2 = ((0 > 0) && v2)
                                }
                                if len(v4) > 0 {
                                    v4[0] = "hello"
                                }
                            }
                        } else {
                            var v139 uint64 = 2
                            switch "a" {
                            case "hello":
                                v1 = v2
                            }
                            switch "a" {
                            case "":
                                v2 = v1
                                switch (fi134 * fi135) {
                                case 0:
                                    v3[""] = ("" + (("foo bar" + "a") + "hello"))
                                }
                                fmt.Println(v53, v139, v2)
                            }
                            _ = v139
                        }
                        _ = fi135
                    }
                    _ = fi134
                }
            default:
                v1 = (118 > 0)
            }
        default:
            v2 = (!(!((0 >= 2) && v1)))
        }
    }
    var v140 int32 = 42
    v140 += v140
    v141 := GenS11{F0: 888, F1: v140}
    switch "a" {
    case "foo bar":
        if v2 {
            for fi142 := 0; fi142 < 5; fi142++ {
                switch "\n\t" {
                case "foo bar":
                    v140 = int32(fi142)
                case "":
                    switch -440 {
                    case 127:
                        switch "foo bar" {
                        case "hello":
                            var v143 float32 = 0.1
                            _ = v143
                        }
                    }
                case "hello":
                    for fi144 := 0; fi144 < 10; fi144++ {
                        fmt.Println(v1)
                        _ = fi144
                    }
                default:
                    for fi145 := 0; fi145 < 14; fi145++ {
                        switch (^(fi145)) {
                        case 2147483647:
                            if false {
                                fmt.Println(fi145, fi142)
                                var v146 uint16 = 0
                                var v147 uint = 18446744073709551615
                                _ = v146
                                _ = v147
                            }
                        case 1:
                            v140 = (v140 & 1)
                        case -2147483648:
                            var v148 float32 = ((-2.25 * (-2.25 - (((0.0 * float32((3.14159 / -2.25))) - -2.25) / 1.5))) - (1e10 + 3.14159))
                            _ = v148
                        default:
                            var v149 float32 = ((3.14159 + -2.25) - ((float32(1.5) * (-2.25 - -2.25)) + 0.1))
                            _ = v149
                        }
                        if ((0 != 2) || (((-2.25 - 1e10) - -2.25) != (float64(-2.25) + -2.25))) {
                            v150 := GenS3{F0: -1, F1: 128, F2: (("" + "hello") + ("" + (("a" + ("foo bar" + ("" + ""))) + "foo bar")))}
                            _ = v150
                        }
                        _ = fi145
                    }
                }
                delete(v3, (("" + "hello") + "foo bar"))
                _ = fi142
            }
            v151 := GenS12{F0: -156, F1: 2, F2: (true || v1)}
            v140 = 2147483647
            _ = v151
        }
    case "":
        var v152 float64 = 0.0
        _ = v152
    case "a":
        if true {
            v153 := v3["hello"]
            _ = v153
        }
    default:
        switch -699 {
        case -2147483648:
            if (!true) {
                fmt.Println(v1, v2, v140)
                if v2 {
                    switch -9223372036854775808 {
                    case 127:
                        v3[""] = ""
                    case 0:
                        v2 = v1
                    default:
                        if v2 {
                            var v154 int8 = (v53 & (^(v53)))
                            var v155 bool = false
                            fmt.Println(v154, v155)
                            _ = v154
                            _ = v155
                        } else {
                            v156 := GenS11{F0: 65535, F1: (v140 + v140)}
                            _ = v156
                        }
                    }
                } else {
                    var v157 string = "hello"
                    v158 := GenS9{F0: 811}
                    _ = v157
                    _ = v158
                }
                v2 = (814 > 42)
            } else {
                if v2 {
                    for fi159 := 0; fi159 < 7; fi159++ {
                        if false {
                            if v1 {
                                v2 = (42 != 128)
                                var v160 uint32 = 2147483647
                                v2 = (2 != 115)
                                _ = v160
                            }
                            v161 := GenS12{F0: -848, F1: 64, F2: false}
                            _ = v161
                        }
                        _ = fi159
                    }
                    v2 = (-727 >= -1)
                } else {
                    v140 += v140
                    switch "\n\t" {
                    case "":
                        v53 -= v53
                    }
                    var v162 float64 = 3.14159
                    _ = v162
                }
                var v163 uint8 = 0
                var v164 int16 = 598
                _ = v163
                _ = v164
            }
        case 127:
            for fi165 := 0; fi165 < 17; fi165++ {
                v166 := GenS13{F0: 956}
                _ = fi165
                _ = v166
            }
        case 2147483647:
            v1 = (genFn0(32767, 621, -9223372036854775808) != -558)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v140:", v140)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v140
    _ = v141
    _ = v2
    _ = v3
    _ = v4
    _ = v53
}
