package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 uint64; F2 int }
type GenS3 struct { F0 uint64; F1 int32; F2 int16 }
type GenS4 struct { F0 string; F1 int8; F2 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 int16; F1 int }
type GenS9 struct { F0 string; F1 int16; F2 uint32 }
type GenS10 struct { F0 bool; F1 bool }
type GenS11 struct { F0 string; F1 bool }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 string; F1 int8 }
type GenS14 struct { F0 int32; F1 int64 }
type GenS15 struct { F0 uint; F1 int32; F2 uint }
func genFn0(p0 uint64) bool {
    for fi4 := 0; fi4 < 20; fi4++ {
        if true {
            var v5 int8 = 127
            v6 := GenS1{F0: 1}
            v7 := GenS2{F0: -466, F1: (p0 | 2), F2: fi4}
            _ = v5
            _ = v6
            _ = v7
        }
        p0 -= 1
        var v8 int32 = -1
        _ = fi4
        _ = v8
    }
    switch "foo bar" {
    case "foo bar":
        v9 := []int{127, -321, 237, -2147483648}
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        switch (fi10 * 1) {
        case 42:
            if (!true) {
                p0 = p0
            } else {
                fmt.Println(fi10, fi10)
                var v11 uint16 = 217
                _ = v11
            }
        case 2147483647:
            fmt.Println(p0, fi10)
        case -2147483648:
            p0 += p0
        }
        _ = fi10
    }
    _ = p0
    return true
}

func main() {
    var v1 int64 = 929
    var v2 int16 = -1
    var v3 int8 = 1
    var v12 float64 = 1.5
    switch "a" {
    case "hello":
        v13 := []int{917}
        _ = v13
    case "foo bar":
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := GenS2{F0: -681, F1: 1, F2: fi14}
            v3 = (v3 / (((v3 & (v3 - (^(v3))))) | 1))
            _ = fi14
            _ = v15
        }
    default:
        fmt.Println(v12, v12, v1)
    }
    v16 := []int{2147483647, -2147483648}
    var v17 string = ""
    if true {
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        fmt.Println(v1, v2, v12)
    }
    if (!true) {
        if ((!false) || ((0.1 * (float64((0.0 / 3.14159)) + 3.14159)) >= 0.0)) {
            v18 := GenS3{F0: 18446744073709551615, F1: 2147483647, F2: (^(v2))}
            v19 := GenS1{F0: 9223372036854775807}
            switch v17 {
            case "foo bar":
                v2 = (v2 % (((v2 >> ((1) & 63))) | 1))
            case "":
                for fi20 := 0; fi20 < 18; fi20++ {
                    if false {
                        for fi21 := 0; fi21 < 17; fi21++ {
                            v1 = (^(v1))
                            _ = fi21
                        }
                        for fi22 := 0; fi22 < 6; fi22++ {
                            if (v12 <= float64(1e10)) {
                                switch (fi22 / (((fi20 ^ (fi20 >> ((0) & 63)))) | 1)) {
                                case 1:
                                    v17 = v17
                                case 9223372036854775807:
                                    switch (fi22 * 32767) {
                                    case -2147483648:
                                        v1 += v1
                                    case 128:
                                        v23 := GenS3{F0: 1, F1: 2147483647, F2: v2}
                                        _ = v23
                                    case -961:
                                        var v24 uint8 = 42
                                        _ = v24
                                    }
                                }
                                var v25 uint16 = 1
                                v25 += 255
                                _ = v25
                            }
                            if (false || false) {
                                switch (^(fi22)) {
                                case -184:
                                    switch fi22 {
                                    case 128:
                                        var v26 uint64 = 895
                                        _ = v26
                                    case 127:
                                        var v27 uint16 = 42
                                        _ = v27
                                    case -2147483648:
                                        v28 := GenS1{F0: (v1 - (v1 / (((-(v1))) | 1)))}
                                        _ = v28
                                    default:
                                        switch (v17 + (("hello" + v17) + "\n\t")) {
                                        case "a":
                                            fmt.Println(v2, v12)
                                        case "":
                                            _ = len(v16)
                                        case "\n\t":
                                            if len(v16) > 0 {
                                                v16[0] = fi22
                                            }
                                        default:
                                            v12 -= float64(float32((0.1 - v12)))
                                        }
                                    }
                                case 42:
                                    if (false && (v17 > ("hello" + v17))) {
                                        if (((float64(1.5) < ((v12 - 1.5) * 3.14159)) && (true && true)) || (0 == -250)) {
                                            if len(v16) >= 2 {
                                                v16 = v16[1:len(v16)]
                                            }
                                            var v29 float64 = float64(0.0)
                                            _ = len(v16)
                                            _ = v29
                                        } else {
                                            v30 := GenS4{F0: v17, F1: int8(v1), F2: (false && false)}
                                            var v31 int32 = 1
                                            _ = v30
                                            _ = v31
                                        }
                                    } else {
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                    }
                                }
                            } else {
                                fmt.Println(v1, v17)
                            }
                            if false {
                                v32 := GenS1{F0: v1}
                                var v33 float64 = (v12 - ((v12 + ((v12 / ((v12 - ((v12 * v12) - ((((v12 - v12) + v12) - -2.25) - v12))) - 3.14159)) / 1.5)) - (v12 * -2.25)))
                                fmt.Println(v17, fi22, fi20)
                                _ = v32
                                _ = v33
                            } else {
                                var v34 float64 = float64(float32(1e10))
                                if false {
                                    switch (fi22 | -1) {
                                    case -899:
                                        v35 := GenS5{F0: 2}
                                        _ = v35
                                    case 128:
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                        switch (v17 + v17) {
                                        case "\n\t":
                                            var v36 int = fi20
                                            _ = v36
                                        case "a":
                                            if ((fi22 * fi22) <= (fi20 >> ((328) & 63))) {
                                                var v37 int16 = v2
                                                _ = v37
                                            }
                                        }
                                        var v38 int32 = 42
                                        _ = v38
                                    case 525:
                                        var v39 int32 = 0
                                        _ = v39
                                    default:
                                        v1 += 42
                                    }
                                    fmt.Println(v2, v1, v17)
                                } else {
                                    if (2147483648 < 65535) {
                                        v40 := GenS1{F0: v1}
                                        v41 := GenS6{F0: 32767}
                                        var v42 uint32 = 42
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    }
                                }
                                _ = v34
                            }
                            _ = fi22
                        }
                        fmt.Println(v12)
                    } else {
                        var v43 float64 = v12
                        _ = v43
                    }
                    _ = fi20
                }
            case "hello":
                for fi44 := 0; fi44 < 2; fi44++ {
                    if false {
                        v1 = (v1 - 1)
                    }
                    for fi45 := 0; fi45 < 7; fi45++ {
                        if false {
                            switch fi45 {
                            case 127:
                                var v46 uint16 = 2
                                _ = v46
                            case 921:
                                var v47 uint32 = 2147483647
                                _ = v47
                            }
                            if len(v16) > 0 {
                                v16[0] = fi44
                            }
                            switch (fi45 | fi45) {
                            case 1:
                                var v48 float32 = (1e10 + -2.25)
                                _ = v48
                            }
                        } else {
                            var v49 int32 = 42
                            if false {
                                var v50 uint64 = 388
                                v3 -= int8(fi44)
                                if (18446744073709551615 != 42) {
                                    v51 := GenS7{F0: 720, F1: (v1 % (((v1 - -430)) | 1))}
                                    var v52 float64 = v12
                                    if ((false && true) && false) {
                                        var v53 int64 = v1
                                        v17 = v17
                                        _ = v53
                                    } else {
                                        v50 += (v50 + (v50 + (v50 << ((0) & 63))))
                                        fmt.Println(v12, v17, v52)
                                        var v54 float32 = -2.25
                                        var v55 int8 = (-(v3))
                                        switch fi45 {
                                        case 42:
                                            v56 := GenS5{F0: 1}
                                            var v57 uint = 65535
                                            switch fi44 {
                                            case 0:
                                                var v58 bool = true
                                                _ = v58
                                            }
                                            _ = v56
                                            _ = v57
                                        case 32768:
                                            v59 := GenS6{F0: 287}
                                            _ = v59
                                        case -1:
                                            switch 127 {
                                            case 206:
                                                v60 := GenS3{F0: (v50 % ((v50) | 1)), F1: v49, F2: v2}
                                                _ = v60
                                            default:
                                                v3 = v55
                                            }
                                        }
                                        _ = v54
                                        _ = v55
                                    }
                                    _ = v51
                                    _ = v52
                                }
                                _ = v50
                            } else {
                                fmt.Println(fi44, v1)
                            }
                            fmt.Println(fi45, v3)
                            _ = v49
                        }
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        _ = fi45
                    }
                    _ = fi44
                }
            }
            _ = v18
            _ = v19
        }
        v61 := GenS3{F0: 1, F1: 2147483647, F2: int16(v1)}
        _ = len(v16)
        _ = v61
    }
    v2 = -707
    var v62 uint16 = 42
    v62 += 42
    if false {
        var v63 uint16 = v62
        v3 = v3
        if (!(!false)) {
            if (!((v63 < v63) && false)) {
                if false {
                    v64 := GenS2{F0: 42, F1: 407, F2: -9223372036854775808}
                    _ = v64
                }
            } else {
                if true {
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    switch 127 {
                    case 127:
                        v65 := GenS1{F0: v1}
                        _ = v65
                    case 32767:
                        for fi66 := 0; fi66 < 19; fi66++ {
                            var v67 int32 = 1
                            var v68 int32 = v67
                            _ = fi66
                            _ = v67
                            _ = v68
                        }
                    case -946:
                        for fi69 := 0; fi69 < 16; fi69++ {
                            switch (fi69 + (^(fi69))) {
                            case -180:
                                if ((v17 + "hello") >= v17) {
                                    switch 127 {
                                    case -185:
                                        if false {
                                            if (!(v1 >= (-(v1)))) {
                                                v12 = v12
                                            }
                                            fmt.Println(fi69, v12, fi69)
                                        } else {
                                            switch (fi69 << ((0) & 63)) {
                                            case 466:
                                                v70 := GenS8{F0: 798, F1: 4294967295}
                                                _ = v70
                                            case -262:
                                                fmt.Println(v2)
                                            case 14:
                                                v63 = (v63 * v63)
                                            default:
                                                v3 = v3
                                            }
                                            fmt.Println(v63)
                                        }
                                    case 127:
                                        v3 = v3
                                    }
                                } else {
                                    v71 := GenS2{F0: 2147483647, F1: 1, F2: -170}
                                    _ = v71
                                }
                            default:
                                switch fi69 {
                                case 968:
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                case 127:
                                    if (!(false || false)) {
                                        switch (v17 + ((v17 + v17) + "foo bar")) {
                                        case "a":
                                            v72 := GenS1{F0: (^(v1))}
                                            _ = v72
                                        case "":
                                            fmt.Println(v3, fi69)
                                        case "\n\t":
                                            var v73 int32 = -1
                                            _ = v73
                                        default:
                                            var v74 string = v17
                                            _ = v74
                                        }
                                    }
                                }
                            }
                            v75 := GenS9{F0: v17, F1: v2, F2: 138}
                            v17 = v17
                            _ = fi69
                            _ = v75
                        }
                    }
                } else {
                    var v76 uint = 18446744073709551615
                    v17 = (((v17 + "foo bar") + (("hello" + v17) + ("foo bar" + (v17 + (v17 + v17))))) + v17)
                    _ = v76
                }
                for fi77 := 0; fi77 < 6; fi77++ {
                    if true {
                        if ((float64(0.1) * float64(-2.25)) != ((0.1 * 1e10) * v12)) {
                            if false {
                                switch (("" + (v17 + v17)) + ("" + (v17 + v17))) {
                                case "a":
                                    v2 -= (v2 + v2)
                                }
                            }
                            fmt.Println(v63, v62, v12)
                        } else {
                            fmt.Println(v3, v17, fi77)
                            v3 -= 1
                            var v78 bool = (true || (false || ((0 <= 1) && false)))
                            _ = v78
                        }
                        var v79 int16 = -667
                        if (v63 > v63) {
                            var v80 float64 = float64((0.0 / float32((v12 * (((0.0 - 1.5) / (v12 - float64((0.0 * (0.0 + 0.1))))) * v12)))))
                            _ = v80
                        } else {
                            var v81 int16 = v2
                            _ = v81
                        }
                        _ = v79
                    }
                    var v82 uint64 = 459
                    var v83 uint32 = 2
                    _ = fi77
                    _ = v82
                    _ = v83
                }
                v84 := GenS9{F0: v17, F1: (v2 * (^(v2))), F2: 0}
                _ = v84
            }
        } else {
            for fi85 := 0; fi85 < 12; fi85++ {
                for fi86 := 0; fi86 < 6; fi86++ {
                    v87 := GenS9{F0: "", F1: v2, F2: 1}
                    var v88 uint8 = 1
                    _ = fi86
                    _ = v87
                    _ = v88
                }
                _ = fi85
            }
        }
        _ = v63
    }
    switch ("" + v17) {
    case "foo bar":
        switch 136 {
        case -354:
            if ((v2 % ((6) | 1)) >= v2) {
                v2 = int16(v1)
            }
        default:
            var v89 uint = 42
            _ = v89
        }
    case "hello":
        for fi90 := 0; fi90 < 1; fi90++ {
            v91 := GenS10{F0: true, F1: false}
            _ = fi90
            _ = v91
        }
    }
    switch "foo bar" {
    case "":
        v3 = v3
    default:
        for fi92 := 0; fi92 < 6; fi92++ {
            fmt.Println(v62)
            if true {
                if (127 > 42) {
                    fmt.Println(v17, v2)
                    var v93 int16 = v2
                    _ = v93
                }
                var v94 uint32 = 4294967295
                _ = v94
            } else {
                v3 += v3
                v95 := GenS11{F0: v17, F1: true}
                v12 += v12
                _ = v95
            }
            _ = fi92
        }
    }
    v17 = "hello"
    if false {
        var v96 uint8 = 179
        for fi97 := 0; fi97 < 11; fi97++ {
            for fi98 := 0; fi98 < 15; fi98++ {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                var v99 int8 = v3
                _ = fi98
                _ = v99
            }
            fmt.Println(v1, v12, v2)
            _ = fi97
        }
        _ = v96
    } else {
        var v100 float64 = (0.1 / 1.0)
        v1 += (v1 / ((v1) | 1))
        for fi101 := 0; fi101 < 16; fi101++ {
            v102 := GenS6{F0: v62}
            switch fi101 {
            case 127:
                switch 127 {
                case -582:
                    var v103 string = v17
                    _ = v103
                case 4294967295:
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                case 127:
                    if (true && ((fi101 < 1) && (!(!false)))) {
                        if (v3 > (v3 << ((0) & 63))) {
                            for fi104 := 0; fi104 < 14; fi104++ {
                                var v105 float32 = ((1.5 / 1.0) * ((0.0 - 1e10) - 1.5))
                                _ = fi104
                                _ = v105
                            }
                            for fi106 := 0; fi106 < 5; fi106++ {
                                switch "foo bar" {
                                case "\n\t":
                                    var v107 int64 = v1
                                    _ = v107
                                default:
                                    v100 = v100
                                }
                                if (42 == 4294967295) {
                                    switch ("\n\t" + v17) {
                                    case "hello":
                                        var v108 int32 = 42
                                        _ = v108
                                    case "foo bar":
                                        v109 := GenS4{F0: v17, F1: v3, F2: true}
                                        _ = v109
                                    default:
                                        var v110 int16 = (v2 | 1)
                                        _ = v110
                                    }
                                    v17 = ("\n\t" + v17)
                                } else {
                                    if (1 >= 4294967295) {
                                        v111 := GenS12{F0: 2}
                                        var v112 float64 = v100
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                        _ = v111
                                        _ = v112
                                    } else {
                                        v113 := GenS5{F0: 0}
                                        _ = v113
                                    }
                                    v114 := GenS13{F0: v17, F1: -125}
                                    _ = v114
                                }
                                var v115 uint16 = (v62 / (((v62 >> ((18446744073709551615) & 63))) | 1))
                                _ = fi106
                                _ = v115
                            }
                            v116 := GenS14{F0: -2147483648, F1: (v1 << ((255) & 63))}
                            _ = v116
                        } else {
                            var v117 int64 = 1
                            if ((v17 + "") < v17) {
                                var v118 float32 = -2.25
                                v16 = append(v16, 0)
                                _ = v118
                            } else {
                                if false {
                                    v100 -= v100
                                    var v119 int32 = 0
                                    switch (^(fi101)) {
                                    case -2147483648:
                                        if (true && ((fi101 < (fi101 >> ((18446744073709551615) & 63))) && false)) {
                                            var v120 uint = 4294967295
                                            switch "foo bar" {
                                            case "\n\t":
                                                if false {
                                                    if (false && true) {
                                                        v1 += v117
                                                    }
                                                    v121 := GenS10{F0: ((0 > v3) || false), F1: false}
                                                    var v122 int16 = v2
                                                    _ = v121
                                                    _ = v122
                                                }
                                            case "foo bar":
                                                if len(v16) >= 2 {
                                                    v16 = v16[1:len(v16)]
                                                }
                                            case "a":
                                                for fi123 := 0; fi123 < 19; fi123++ {
                                                    if len(v16) >= 2 {
                                                        v16 = v16[1:len(v16)]
                                                    }
                                                    _ = len(v16)
                                                    _ = fi123
                                                }
                                            }
                                            _ = v120
                                        }
                                    }
                                    _ = v119
                                }
                                _ = len(v16)
                                var v124 uint8 = 255
                                _ = v124
                            }
                            _ = v117
                        }
                    } else {
                        if (0 > -2147483648) {
                            if true {
                                if (4294967295 < 1) {
                                    _ = len(v16)
                                } else {
                                    if ((v1 & v1) > (v1 - v1)) {
                                        var v125 string = "foo bar"
                                        _ = v125
                                    } else {
                                        var v126 int16 = v2
                                        var v127 int64 = (v1 & 835)
                                        var v128 int8 = (-(v3))
                                        _ = v126
                                        _ = v127
                                        _ = v128
                                    }
                                    var v129 int32 = 1
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                    _ = v129
                                }
                                var v130 int64 = 42
                                _ = v130
                            }
                            var v131 uint16 = 691
                            v2 += (v2 | (v2 << ((18446744073709551615) & 63)))
                            _ = v131
                        } else {
                            var v132 uint16 = 65535
                            for fi133 := 0; fi133 < 3; fi133++ {
                                if true {
                                    if len(v16) > 0 {
                                        v16[0] = 0
                                    }
                                    var v134 bool = (!true)
                                    _ = v134
                                }
                                v100 = 3.14159
                                _ = fi133
                            }
                            if false {
                                v135 := GenS15{F0: 77, F1: 42, F2: 305}
                                _ = v135
                            } else {
                                switch (fi101 >> ((42) & 63)) {
                                case 128:
                                    v16 = append(v16, fi101)
                                }
                                v132 = v132
                            }
                            _ = v132
                        }
                        _ = len(v16)
                        var v136 uint8 = 127
                        _ = v136
                    }
                }
            case 4294967295:
                var v137 int8 = v3
                _ = v137
            }
            _ = fi101
            _ = v102
        }
        _ = v100
    }
    if ((!(2147483647 <= 2147483647)) && (int32(v62) > -1)) {
        fmt.Println(v12)
        v12 = (-2.25 + v12)
    }
    for fi138 := 0; fi138 < 16; fi138++ {
        v16 = append(v16, (fi138 * 396))
        _ = fi138
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v62
}
