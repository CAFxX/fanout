package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 uint16; F2 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int; F1 int; F2 uint32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 string; F1 bool }
type GenS11 struct { F0 int32; F1 int }
type GenS12 struct { F0 uint; F1 uint32 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 bool; F1 int32 }
type GenS15 struct { F0 bool }
func genFn0(p0 uint32, p1 uint64, p2 string) int16 {
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 float64 = ((0.0 + float64(1.5)) / 3.14159)
        p2 = p2
        switch (fi7 + fi7) {
        case 4294967295:
            v9 := []int{(fi7 * fi7)}
            _ = v9
        case -20:
            var v10 uint = 1
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    var v11 uint32 = 65535
    var v12 int32 = 0
    fmt.Println(p0, p1, v11)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[int]string)
    v3 := []uint32{2147483647, 2, 2147483647}
    v4 := []string{("a" + "\n\t"), ""}
    var v5 int = -493
    var v6 string = ("foo bar" + ("hello" + "hello"))
    switch (v5 >> ((622) & 63)) {
    case 2147483647:
        _ = len(v2)
    case 1:
        for fi13 := 0; fi13 < 8; fi13++ {
            for fi14 := 0; fi14 < 1; fi14++ {
                if (2 <= 65535) {
                    v6 = (v6 + ((v6 + v6) + v6))
                    v2[v5] = v6
                }
                fmt.Println(v5)
                _ = fi14
            }
            _ = fi13
        }
    }
    fmt.Println(v5)
    v6 = v6
    _ = len(v2)
    if true {
        delete(v2, v5)
        switch v5 {
        case 9223372036854775807:
            v1 -= 2
        case -1:
            for fi15 := 0; fi15 < 2; fi15++ {
                switch "foo bar" {
                case "hello":
                    if len(v3) > 0 {
                        v3[0] = 239
                    }
                case "foo bar":
                    v5 = (fi15 << ((255) & 63))
                case "a":
                    if ((v6 == v6) || false) {
                        for fi16 := 0; fi16 < 20; fi16++ {
                            _ = len(v4)
                            switch (fi15 / (((fi15 / (((fi16 >> ((132) & 63))) | 1))) | 1)) {
                            case 32768:
                                if false {
                                    var v17 int16 = 447
                                    fmt.Println(v5)
                                    var v18 bool = (!(!false))
                                    _ = v17
                                    _ = v18
                                } else {
                                    var v19 int32 = 42
                                    v20 := GenS1{F0: -128, F1: 0, F2: 885}
                                    var v21 int = fi16
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                                v4 = append(v4, ((v6 + v6) + "hello"))
                                v22 := GenS2{F0: v6}
                                _ = v22
                            default:
                                fmt.Println(fi15, fi15)
                                switch (fi15 & fi16) {
                                case -470:
                                    v6 = v6
                                case -1:
                                    v5 += 42
                                case 42:
                                    switch "hello" {
                                    case "":
                                        v6 = "hello"
                                    }
                                default:
                                    if (127 < v1) {
                                        v23 := GenS1{F0: 109, F1: 255, F2: 0}
                                        switch (v6 + "foo bar") {
                                        case "a":
                                            v24 := GenS1{F0: 1, F1: 550, F2: 0}
                                            _ = v24
                                        case "foo bar":
                                            v25 := GenS1{F0: -1, F1: 1, F2: 646}
                                            var v26 uint16 = 93
                                            fmt.Println(v5, v1, v26)
                                            _ = v25
                                            _ = v26
                                        case "":
                                            v2[(v5 / ((fi15) | 1))] = "foo bar"
                                            v5 -= -561
                                        }
                                        v5 = (fi15 % ((-9223372036854775808) | 1))
                                        _ = v23
                                    } else {
                                        fmt.Println(v5)
                                        var v27 int32 = 1
                                        delete(v2, 9223372036854775807)
                                        _ = v27
                                    }
                                }
                                v5 = (fi15 * 128)
                                if true {
                                    v5 = v5
                                    v28 := GenS1{F0: -1, F1: 2, F2: 57}
                                    _ = v28
                                }
                                if false {
                                    if true {
                                        v29 := GenS2{F0: "foo bar"}
                                        _ = v29
                                    }
                                    switch "" {
                                    case "a":
                                        v1 -= 1
                                    case "foo bar":
                                        v30 := v2[2147483647]
                                        _ = v30
                                    default:
                                        v1 += (v1 + 127)
                                    }
                                    switch 127 {
                                    case 9223372036854775807:
                                        switch v6 {
                                        case "\n\t":
                                            v31 := GenS2{F0: ((("" + v6) + v6) + "")}
                                            _ = v31
                                        default:
                                            switch "foo bar" {
                                            case "hello":
                                                v1 = 2
                                            }
                                            v32 := GenS3{F0: 0, F1: 2, F2: 1}
                                            if true {
                                                fmt.Println(v1)
                                            }
                                            _ = v32
                                        }
                                    case -157:
                                        v33 := GenS4{F0: 9223372036854775807}
                                        _ = v33
                                    }
                                    v5 = (fi16 % ((fi16) | 1))
                                } else {
                                    v5 -= v5
                                }
                            }
                            v5 = 32768
                            if len(v4) > 0 {
                                v4[0] = "hello"
                            }
                            var v34 int32 = 42
                            var v35 int64 = 42
                            v34 = -2147483648
                            delete(v2, -2147483648)
                            fmt.Println(fi15)
                            v5 -= (fi15 ^ (fi15 - 32767))
                            var v36 float32 = (1.5 * 3.14159)
                            _ = fi16
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        for fi37 := 0; fi37 < 4; fi37++ {
                            var v38 int32 = 158
                            v39 := GenS5{F0: 843}
                            v5 = v5
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    } else {
                        if (1 > -1) {
                            if len(v3) > 0 {
                                v3[0] = 4294967295
                            }
                        } else {
                            for fi40 := 0; fi40 < 14; fi40++ {
                                var v41 int64 = 1
                                v42 := GenS6{F0: fi15, F1: fi15, F2: 0}
                                var v43 bool = true
                                if v43 {
                                    v43 = true
                                    switch -238 {
                                    case 711:
                                        var v44 int32 = 1
                                        _ = v44
                                    }
                                    if (!(((fi40 ^ v5) <= (fi40 + v5)) && v43)) {
                                        if (!v43) {
                                            if (-1 >= int32(fi40)) {
                                                switch "" {
                                                case "":
                                                    _ = len(v3)
                                                case "foo bar":
                                                    var v45 int16 = 1000
                                                    _ = v45
                                                case "a":
                                                    if v43 {
                                                        if ((((1.5 + 1.5) + 0.0) - (1e10 - 1.5)) >= 0.0) {
                                                            var v46 uint = 65535
                                                            fmt.Println(fi15, v1)
                                                            delete(v2, v5)
                                                            _ = v46
                                                        }
                                                    } else {
                                                        v6 = "a"
                                                        if len(v3) > 0 {
                                                            v3[0] = 810
                                                        }
                                                        v6 = "\n\t"
                                                        v47 := GenS7{F0: 1}
                                                        v1 -= v1
                                                        _ = v47
                                                    }
                                                default:
                                                    var v48 int = (-(fi40))
                                                    _ = v48
                                                }
                                                if (!true) {
                                                    v49 := GenS8{F0: (((v5 | (v5 & 70)) != -9223372036854775808) && true)}
                                                    _ = v49
                                                }
                                                v50 := GenS9{F0: 42}
                                                var v51 int8 = 1
                                                v43 = (v1 <= 2)
                                                v52 := GenS2{F0: v6}
                                                var v53 uint16 = 138
                                                _ = v50
                                                _ = v51
                                                _ = v52
                                                _ = v53
                                            }
                                            var v54 int32 = 1
                                            var v55 uint32 = 4294967295
                                            _ = len(v3)
                                            fmt.Println(v55, fi15, fi15)
                                            v56 := GenS10{F0: v6, F1: v43}
                                            _ = v54
                                            _ = v55
                                            _ = v56
                                        }
                                        switch "hello" {
                                        case "":
                                            fmt.Println(v6)
                                        case "a":
                                            var v57 uint32 = 2147483648
                                            _ = v57
                                        default:
                                            if (v43 || v43) {
                                                fmt.Println(v6)
                                                if (2 != 987) {
                                                    switch fi40 {
                                                    case 816:
                                                        v58 := GenS11{F0: 42, F1: (fi15 << ((71) & 63))}
                                                        _ = v58
                                                    case 42:
                                                        fmt.Println(v43, v43, fi15)
                                                    }
                                                    fmt.Println(fi15, v41)
                                                    var v59 uint64 = 0
                                                    _ = v59
                                                }
                                            }
                                        }
                                        var v60 string = "foo bar"
                                        fmt.Println(v1, fi40, v5)
                                        v61 := GenS8{F0: (0 == 65535)}
                                        v60 = "\n\t"
                                        v62 := GenS1{F0: 100, F1: 1, F2: 920}
                                        _ = v60
                                        _ = v61
                                        _ = v62
                                    }
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = 4294967295
                                    }
                                }
                                v3 = append(v3, 796)
                                _ = fi40
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                        }
                    }
                }
                switch fi15 {
                case -641:
                    for fi63 := 0; fi63 < 15; fi63++ {
                        var v64 int64 = 9223372036854775807
                        var v65 uint32 = 2
                        _ = fi63
                        _ = v64
                        _ = v65
                    }
                default:
                    for fi66 := 0; fi66 < 17; fi66++ {
                        switch v6 {
                        case "\n\t":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case "hello":
                            if true {
                                if len(v3) > 0 {
                                    v3[0] = 1
                                }
                                var v67 float32 = 1e10
                                fmt.Println(v5, fi66)
                                _ = len(v2)
                                _ = v67
                            }
                        case "foo bar":
                            if true {
                                fmt.Println(v6, fi15)
                                v68 := GenS4{F0: 42}
                                if (!(165 < 42)) {
                                    switch (v5 * (-(fi15))) {
                                    case 2147483647:
                                        v69 := GenS4{F0: -1}
                                        _ = v69
                                    case 128:
                                        var v70 int16 = 0
                                        _ = v70
                                    }
                                    var v71 int64 = 0
                                    fmt.Println(fi15)
                                    _ = v71
                                } else {
                                    switch ((v6 + "a") + ("foo bar" + (v6 + v6))) {
                                    case "a":
                                        v72 := GenS8{F0: (!false)}
                                        _ = v72
                                    case "foo bar":
                                        v73 := v2[(fi15 * -2147483648)]
                                        _ = v73
                                    }
                                    var v74 int8 = 1
                                    _ = len(v2)
                                    _ = v74
                                }
                                _ = v68
                            } else {
                                v75 := GenS12{F0: 0, F1: 65535}
                                switch fi66 {
                                case -2147483648:
                                    v1 += (^(v1))
                                case 42:
                                    v76 := GenS11{F0: 662, F1: fi66}
                                    _ = v76
                                case 1:
                                    v6 = v6
                                }
                                if false {
                                    switch ("\n\t" + (("\n\t" + v6) + v6)) {
                                    case "hello":
                                        v1 += (v1 | (v1 * v1))
                                    case "a":
                                        switch v6 {
                                        case "":
                                            v6 = "hello"
                                        case "hello":
                                            fmt.Println(v6, v1, fi66)
                                        default:
                                            fmt.Println(fi66, v1)
                                        }
                                    case "":
                                        v77 := GenS13{F0: 1}
                                        _ = v77
                                    }
                                    var v78 int32 = 1
                                    var v79 uint32 = 421
                                    v80 := GenS2{F0: (v6 + v6)}
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                }
                                fmt.Println(v6, v1)
                                fmt.Println(v6)
                                _ = v75
                            }
                        default:
                            var v81 uint64 = 188
                            _ = v81
                        }
                        _ = fi66
                    }
                }
                switch ((v6 + v6) + ("hello" + ("hello" + ((v6 + "foo bar") + v6)))) {
                case "a":
                    if true {
                        var v82 int32 = 1
                        _ = v82
                    }
                case "":
                    var v83 string = ((v6 + v6) + v6)
                    _ = v83
                case "hello":
                    var v84 uint8 = 0
                    _ = v84
                }
                _ = fi15
            }
        case 35:
            v6 = ("a" + v6)
        }
        v1 -= v1
    }
    v85 := GenS13{F0: 255}
    fmt.Println(v6, v5)
    v86 := GenS11{F0: -1, F1: (v5 / ((-964) | 1))}
    if false {
        if (-745 >= -2147483648) {
            for fi87 := 0; fi87 < 2; fi87++ {
                switch 1 {
                case 1:
                    var v88 uint16 = 447
                    _ = v88
                case 127:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    switch (v6 + v6) {
                    case "hello":
                        if true {
                            if (-104 >= 127) {
                                v89 := GenS3{F0: 31, F1: 2, F2: 100}
                                v5 += v5
                                for fi90 := 0; fi90 < 17; fi90++ {
                                    v91 := GenS4{F0: 42}
                                    if ((^(v5)) >= fi87) {
                                        fmt.Println(v1)
                                    } else {
                                        if true {
                                            if (!(int16(fi87) != 910)) {
                                                v92 := GenS14{F0: false, F1: -2147483648}
                                                _ = v92
                                            }
                                        }
                                        v1 = (v1 * (v1 % ((v1) | 1)))
                                    }
                                    v93 := v2[fi87]
                                    if false {
                                        v5 -= fi90
                                        var v94 float32 = (float32(float64((0.1 - (3.14159 + (3.14159 - (-2.25 - 0.0)))))) * 1.5)
                                        var v95 uint32 = 130
                                        _ = v94
                                        _ = v95
                                    } else {
                                        v96 := GenS5{F0: 919}
                                        _ = v96
                                    }
                                    var v97 uint16 = 32767
                                    _ = fi90
                                    _ = v91
                                    _ = v93
                                    _ = v97
                                }
                                _ = v89
                            } else {
                                for fi98 := 0; fi98 < 17; fi98++ {
                                    v99 := GenS9{F0: 865}
                                    _ = fi98
                                    _ = v99
                                }
                                fmt.Println(fi87, v6, v6)
                            }
                        }
                    case "a":
                        fmt.Println(v1)
                    }
                }
                v100 := GenS1{F0: 127, F1: 255, F2: 962}
                _ = fi87
                _ = v100
            }
            v101 := GenS7{F0: v5}
            v102 := GenS5{F0: 42}
            _ = v101
            _ = v102
        }
        var v103 float32 = (0.0 * (((float32((0.1 * (0.0 - (-2.25 / -2.25)))) / (1.5 * float32((0.1 * 1.5)))) - (3.14159 - 1.5)) * -2.25))
        _ = v103
    } else {
        v5 = (v5 + v5)
        var v104 uint16 = 255
        var v105 uint32 = 2147483647
        _ = v104
        _ = v105
    }
    var v106 int = (^(v5))
    for fi107 := 0; fi107 < 9; fi107++ {
        var v108 uint8 = 42
        switch v6 {
        case "a":
            for fi109 := 0; fi109 < 3; fi109++ {
                v4 = append(v4, v6)
                var v110 int8 = -1
                v110 -= 0
                _ = fi109
                _ = v110
            }
        case "foo bar":
            switch 9223372036854775807 {
            case 128:
                v5 = (v5 | (v5 ^ v5))
            default:
                switch (v6 + v6) {
                case "\n\t":
                    if true {
                        if (float32(float64(float32(((-2.25 - -2.25) + -2.25)))) == ((0.1 * 3.14159) * -2.25)) {
                            switch (v106 - v5) {
                            case -573:
                                for fi111 := 0; fi111 < 19; fi111++ {
                                    var v112 uint8 = 255
                                    _ = len(v4)
                                    _ = fi111
                                    _ = v112
                                }
                            case 2147483647:
                                v108 = v1
                            case -180:
                                v1 = v1
                            default:
                                var v113 uint32 = 804
                                _ = v113
                            }
                        }
                        var v114 uint32 = 1
                        v115 := GenS12{F0: 922, F1: v114}
                        _ = v114
                        _ = v115
                    } else {
                        switch fi107 {
                        case 674:
                            _ = len(v3)
                        case 127:
                            if (true || false) {
                                fmt.Println(v106, v5, v5)
                            } else {
                                for fi116 := 0; fi116 < 11; fi116++ {
                                    if len(v4) > 0 {
                                        v4[0] = (v6 + (("hello" + v6) + ((("foo bar" + (v6 + "\n\t")) + v6) + "\n\t")))
                                    }
                                    switch v6 {
                                    case "a":
                                        if (255 <= 65535) {
                                            v106 -= 47
                                        } else {
                                            if (1 <= 0) {
                                                var v117 uint8 = 128
                                                switch (v5 ^ (fi107 / ((v106) | 1))) {
                                                case 4294967295:
                                                    v6 = v6
                                                }
                                                v118 := GenS2{F0: v6}
                                                _ = v117
                                                _ = v118
                                            } else {
                                                var v119 uint64 = 854
                                                v119 += v119
                                                if false {
                                                    switch 4294967295 {
                                                    case 127:
                                                        fmt.Println(v1, v108)
                                                    case 4294967295:
                                                        v120 := v2[(fi116 - (v5 / ((88) | 1)))]
                                                        _ = v120
                                                    }
                                                    v2[v5] = v6
                                                } else {
                                                    v1 += (v108 % ((v1) | 1))
                                                    v106 -= int(v108)
                                                }
                                                _ = v119
                                            }
                                            if false {
                                                var v121 bool = true
                                                v122 := GenS8{F0: (1 == uint64(v1))}
                                                _ = v121
                                                _ = v122
                                            } else {
                                                _ = len(v3)
                                            }
                                        }
                                    case "\n\t":
                                        switch "foo bar" {
                                        case "":
                                            switch v5 {
                                            case -1:
                                                if (42 == 2) {
                                                    delete(v2, 44)
                                                }
                                            case 128:
                                                v1 += 255
                                            case 9223372036854775807:
                                                switch (v6 + "\n\t") {
                                                case "hello":
                                                    var v123 int64 = -939
                                                    _ = v123
                                                }
                                            default:
                                                _ = len(v2)
                                            }
                                        default:
                                            v106 += 0
                                        }
                                    default:
                                        v124 := GenS15{F0: ((!true) && ((v106 * (fi107 / (((fi116 * v5)) | 1))) <= v5))}
                                        _ = v124
                                    }
                                    _ = fi116
                                }
                            }
                        }
                        var v125 string = "a"
                        for fi126 := 0; fi126 < 8; fi126++ {
                            if false {
                                if false {
                                    fmt.Println(v106, fi107, v1)
                                    var v127 int = fi107
                                    v128 := v2[(fi107 | v127)]
                                    var v129 bool = true
                                    switch "hello" {
                                    case "a":
                                        v128 = v6
                                    case "":
                                        if false {
                                            v2[fi107] = "\n\t"
                                            v5 = fi126
                                        }
                                    case "hello":
                                        var v130 uint64 = 4294967295
                                        _ = v130
                                    }
                                    v131 := GenS1{F0: 91, F1: 773, F2: 111}
                                    _ = v127
                                    _ = v128
                                    _ = v129
                                    _ = v131
                                }
                                var v132 uint16 = 2
                                var v133 float32 = 0.0
                                _ = v132
                                _ = v133
                            }
                            var v134 uint16 = 32767
                            var v135 int64 = -9223372036854775808
                            _ = fi126
                            _ = v134
                            _ = v135
                        }
                        _ = v125
                    }
                case "":
                    v136 := v2[(fi107 >> ((0) & 63))]
                    _ = v136
                case "hello":
                    delete(v2, (^(v5)))
                }
            }
        }
        _ = fi107
        _ = v108
    }
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v85
    _ = v86
}
