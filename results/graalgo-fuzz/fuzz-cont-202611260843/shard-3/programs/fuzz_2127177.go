package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint; F1 int64; F2 uint64 }
type GenS6 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint; F1 uint16; F2 int16 }
type GenS9 struct { F0 uint32; F1 uint64 }
type GenS10 struct { F0 uint8; F1 int64; F2 int }
type GenS11 struct { F0 int16; F1 uint; F2 uint }
type GenS12 struct { F0 uint16; F1 uint16 }
type GenS13 struct { F0 uint16; F1 int }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint8; F1 int32 }
type GenS16 struct { F0 uint; F1 uint32; F2 int }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 int; F1 uint; F2 bool }
type GenS19 struct { F0 uint32; F1 int; F2 bool }
type GenS20 struct { F0 bool; F1 int8 }
func genFn0(p0 int16) uint16 {
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 uint16 = 1
        switch "a" {
        case "\n\t":
            if false {
                switch "hello" {
                case "\n\t":
                    var v6 bool = false
                    _ = v6
                case "foo bar":
                    var v7 uint32 = 116
                    _ = v7
                case "hello":
                    var v8 int16 = p0
                    _ = v8
                default:
                    for fi9 := 0; fi9 < 2; fi9++ {
                        if false {
                            fmt.Println(fi4, fi4)
                            if true {
                                v5 += (v5 % (((^(v5))) | 1))
                            }
                        } else {
                            var v10 int16 = (p0 + -1)
                            var v11 uint64 = 18446744073709551615
                            if false {
                                v12 := []int{fi9, (fi4 - fi4), (fi4 / (((^(fi4))) | 1)), (fi9 * (fi4 >> ((18446744073709551615) & 63)))}
                                var v13 bool = (true && false)
                                if true {
                                    var v14 uint64 = 2
                                    if true {
                                        if (1.5 >= 0.1) {
                                            switch "foo bar" {
                                            case "\n\t":
                                                v15 := GenS1{F0: 42, F1: uint8(v14)}
                                                _ = v15
                                            case "hello":
                                                var v16 int32 = -2147483648
                                                _ = v16
                                            }
                                        } else {
                                            var v17 int16 = v10
                                            _ = v17
                                        }
                                        v5 += (v5 + 255)
                                        if ((18446744073709551615 > (v14 % ((553) | 1))) || (false && false)) {
                                            v12 = append(v12, (fi4 & int(v11)))
                                            v18 := GenS1{F0: 100, F1: 127}
                                            _ = v18
                                        }
                                    }
                                    fmt.Println(v10)
                                    _ = v14
                                } else {
                                    v19 := GenS1{F0: -1, F1: 32}
                                    var v20 int32 = 992
                                    fmt.Println(p0)
                                    var v21 int16 = p0
                                    var v22 int64 = -9223372036854775808
                                    if (!true) {
                                        var v23 int32 = v20
                                        v10 += (v10 % ((32767) | 1))
                                        v24 := GenS1{F0: 1, F1: 79}
                                        _ = v23
                                        _ = v24
                                    }
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                }
                                _ = v12
                                _ = v13
                            }
                            _ = v10
                            _ = v11
                        }
                        fmt.Println(fi4, fi9)
                        v5 = (v5 << ((65535) & 63))
                        _ = fi9
                    }
                }
            }
        }
        for fi25 := 0; fi25 < 11; fi25++ {
            p0 += 278
            switch "hello" {
            case "a":
                switch fi4 {
                case -980:
                    switch "a" {
                    case "a":
                        switch fi25 {
                        case 1:
                            v26 := GenS1{F0: -1, F1: 186}
                            _ = v26
                        default:
                            var v27 int = fi4
                            _ = v27
                        }
                    case "foo bar":
                        v28 := GenS1{F0: 107, F1: 49}
                        _ = v28
                    }
                case 4294967295:
                    if (!(0 >= 4294967295)) {
                        v29 := GenS2{F0: 359}
                        if true {
                            var v30 uint8 = 142
                            p0 = (p0 * p0)
                            if (-1 > (^(p0))) {
                                fmt.Println(v30)
                                var v31 int = (fi25 % ((fi4) | 1))
                                var v32 bool = (!(!true))
                                v30 = 0
                                _ = v31
                                _ = v32
                            } else {
                                v33 := GenS3{F0: (v5 % ((v5) | 1))}
                                v5 = v5
                                var v34 int8 = 127
                                var v35 float32 = 0.0
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                            _ = v30
                        }
                        _ = v29
                    } else {
                        if (!true) {
                            var v36 uint = 144
                            v5 -= v5
                            _ = v36
                        }
                    }
                case -774:
                    v5 = (v5 - 495)
                }
            case "":
                p0 = p0
            default:
                fmt.Println(p0)
            }
            p0 = (p0 % ((p0) | 1))
            _ = fi25
        }
        _ = fi4
        _ = v5
    }
    var v37 int16 = p0
    _ = p0
    _ = v37
    return 255
}
func genFn1() uint32 {
    switch (("a" + (("" + ((("a" + "a") + ("foo bar" + "\n\t")) + (("a" + "hello") + (((("\n\t" + ("\n\t" + (("" + "") + "foo bar"))) + "") + "") + "foo bar")))) + "")) + ("\n\t" + ("\n\t" + "foo bar"))) {
    case "\n\t":
        var v38 uint64 = 4294967295
        _ = v38
    }
    var v39 float64 = (0.0 - ((float64(1e10) / (((0.0 * 1.5) + ((1e10 + 0.0) / 1.0)) - 1.5)) + 0.1))
    for fi40 := 0; fi40 < 9; fi40++ {
        if true {
            if false {
                var v41 bool = false
                _ = v41
            } else {
                var v42 int = 1
                var v43 uint32 = 42
                switch v42 {
                case 1:
                    for fi44 := 0; fi44 < 13; fi44++ {
                        switch (-(v42)) {
                        case -2147483648:
                            v45 := []int{(fi44 % ((fi44) | 1))}
                            _ = v45
                        case 42:
                            fmt.Println(v39, v39, v42)
                        }
                        var v46 float32 = -2.25
                        v47 := GenS4{F0: 42, F1: 2147483647, F2: 0}
                        _ = fi44
                        _ = v46
                        _ = v47
                    }
                default:
                    var v48 bool = false
                    _ = v48
                }
                _ = v42
                _ = v43
            }
            fmt.Println(v39)
            v49 := GenS5{F0: 42, F1: 1, F2: 4294967295}
            _ = v49
        }
        _ = fi40
    }
    for fi50 := 0; fi50 < 8; fi50++ {
        switch fi50 {
        case 32768:
            for fi51 := 0; fi51 < 19; fi51++ {
                v39 += v39
                var v52 uint8 = 132
                var v53 uint = 802
                switch (fi51 % ((-2147483648) | 1)) {
                case 9223372036854775807:
                    switch "a" {
                    case "\n\t":
                        var v54 uint64 = 284
                        _ = v54
                    case "foo bar":
                        var v55 int8 = 88
                        _ = v55
                    case "":
                        v39 = (1.5 * 0.0)
                    }
                case -9223372036854775808:
                    var v56 bool = (-31 >= -68)
                    _ = v56
                default:
                    if true {
                        v39 += ((0.0 * (v39 * 1.5)) - (v39 * v39))
                        v57 := []int{32768, (fi50 << (((v53 << ((204) & 63))) & 63))}
                        var v58 float32 = (-2.25 - (0.0 / 3.14159))
                        v59 := GenS5{F0: 18446744073709551615, F1: 477, F2: 481}
                        if false {
                            if len(v57) >= 2 {
                                v57 = v57[1:len(v57)]
                            }
                            if len(v57) > 0 {
                                v57[0] = fi50
                            }
                        }
                        _ = v57
                        _ = v58
                        _ = v59
                    } else {
                        switch fi50 {
                        case 0:
                            v60 := []int{}
                            _ = v60
                        case 126:
                            switch fi50 {
                            case 1:
                                if (!(!false)) {
                                    if (!(genFn0(-32768) != genFn0(-754))) {
                                        if false {
                                            v61 := GenS6{F0: 2, F1: v52, F2: 42}
                                            v62 := []int{(fi50 % ((fi51) | 1))}
                                            switch "a" {
                                            case "a":
                                                if (("foo bar" + ("hello" + "foo bar")) > ("" + ("" + "\n\t"))) {
                                                    if true {
                                                        v52 = (^(v52))
                                                    }
                                                }
                                            case "\n\t":
                                                var v63 uint16 = 32767
                                                _ = v63
                                            case "hello":
                                                var v64 int = fi51
                                                _ = v64
                                            }
                                            var v65 int8 = -128
                                            _ = v61
                                            _ = v62
                                            _ = v65
                                        }
                                        if ("foo bar" < "\n\t") {
                                            if false {
                                                var v66 int32 = -2147483648
                                                var v67 uint64 = 42
                                                if ((v52 % ((v52) | 1)) == 1) {
                                                    v68 := []int{}
                                                    _ = len(v68)
                                                    v69 := GenS1{F0: 0, F1: (v52 * 128)}
                                                    _ = v68
                                                    _ = v69
                                                }
                                                _ = v66
                                                _ = v67
                                            } else {
                                                v70 := []int{}
                                                var v71 uint16 = 486
                                                v71 = 2
                                                _ = v70
                                                _ = v71
                                            }
                                        } else {
                                            v72 := []int{-2147483648}
                                            fmt.Println(fi51, v52)
                                            _ = v72
                                        }
                                        v73 := GenS7{F0: 1000}
                                        _ = v73
                                    }
                                    switch (fi50 % ((fi50) | 1)) {
                                    case 42:
                                        if (1 != 2147483648) {
                                            if false {
                                                v53 += (^(v53))
                                            } else {
                                                var v74 bool = ((!false) || (!(false && (0 != 255))))
                                                v75 := GenS8{F0: 0, F1: uint16(v53), F2: 1}
                                                _ = v74
                                                _ = v75
                                            }
                                        }
                                    case -1:
                                        v53 = 65535
                                    case 127:
                                        v52 = v52
                                    default:
                                        var v76 float64 = float64(1e10)
                                        _ = v76
                                    }
                                    var v77 int64 = -1
                                    _ = v77
                                }
                            case 406:
                                var v78 float64 = 1e10
                                _ = v78
                            case 0:
                                switch 9223372036854775807 {
                                case 454:
                                    var v79 uint64 = 4294967295
                                    _ = v79
                                default:
                                    v80 := GenS9{F0: 42, F1: 4294967295}
                                    _ = v80
                                }
                            default:
                                fmt.Println(fi51)
                            }
                        case -2147483648:
                            v52 -= 185
                        default:
                            fmt.Println(v52, v53, fi51)
                        }
                    }
                    v81 := GenS10{F0: v52, F1: 787, F2: fi51}
                    _ = v81
                }
                _ = fi51
                _ = v52
                _ = v53
            }
        case 1:
            for fi82 := 0; fi82 < 2; fi82++ {
                var v83 uint16 = 667
                _ = fi82
                _ = v83
            }
        case 2147483647:
            var v84 uint8 = 1
            _ = v84
        }
        v39 = 0.0
        v39 += float64(-2.25)
        _ = fi50
    }
    _ = v39
    return 65535
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 42
    var v3 int8 = -128
    switch 9223372036854775807 {
    case -2147483648:
        fmt.Println(v3, v2)
    case 0:
        v2 += v2
    }
    if (false && (v3 == -1)) {
        v3 = (v3 >> ((v2) & 63))
        var v85 float64 = (((((0.1 - 1.5) - 0.1) - ((3.14159 + (0.1 * ((1.5 - (((-2.25 - -2.25) / (0.1 + 0.0)) - ((0.0 / 3.14159) + 0.0))) + 1.5))) + float64(0.1))) - ((1e10 / float64(((1.5 - float32(1e10)) / (((float32(float64(1.5)) - 1.5) - 0.1) * 1.5)))) - 0.0)) / -2.25)
        _ = v85
    } else {
        v3 = 100
        var v86 string = ""
        _ = v86
    }
    var v87 int64 = -1
    for fi88 := 0; fi88 < 20; fi88++ {
        var v89 uint64 = 570
        for fi90 := 0; fi90 < 9; fi90++ {
            v91 := []int{}
            _ = fi90
            _ = v91
        }
        _ = fi88
        _ = v89
    }
    v3 += v3
    if ("foo bar" < "foo bar") {
        v92 := GenS4{F0: -1, F1: -972, F2: 1}
        switch 128 {
        case 42:
            delete(v1, 252)
        case 32768:
            v93 := []int{9223372036854775807, -9223372036854775808, 4294967295, 0}
            _ = v93
        default:
            for fi94 := 0; fi94 < 11; fi94++ {
                for fi95 := 0; fi95 < 20; fi95++ {
                    var v96 bool = true
                    var v97 string = ("" + ("\n\t" + ""))
                    _ = fi95
                    _ = v96
                    _ = v97
                }
                v98 := []int{42}
                for fi99 := 0; fi99 < 12; fi99++ {
                    v100 := GenS11{F0: 442, F1: (v2 * 0), F2: (v2 & (v2 / (((v2 | (^(v2)))) | 1)))}
                    v87 = v87
                    switch (fi94 % (((fi94 | -981)) | 1)) {
                    case 1:
                        if ((2 == 42) && false) {
                            var v101 uint64 = 779
                            var v102 int32 = 1
                            _ = v101
                            _ = v102
                        } else {
                            switch -215 {
                            case 2147483647:
                                switch "foo bar" {
                                case "\n\t":
                                    if (fi94 != (fi94 << ((257) & 63))) {
                                        var v103 float32 = 1.5
                                        v87 = 0
                                        if true {
                                            if (18446744073709551615 == (^(v2))) {
                                                v2 = 255
                                                delete(v1, fi99)
                                                v104 := GenS12{F0: 921, F1: 0}
                                                _ = v104
                                            }
                                            fmt.Println(v87)
                                        } else {
                                            v105 := GenS6{F0: 0, F1: 0, F2: v3}
                                            if false {
                                                if false {
                                                    var v106 string = ("foo bar" + "foo bar")
                                                    _ = v106
                                                }
                                                v3 = v3
                                            } else {
                                                v107 := GenS13{F0: 0, F1: fi94}
                                                v3 += v3
                                                _ = v107
                                            }
                                            var v108 uint64 = 2
                                            if true {
                                                _ = len(v98)
                                                _ = len(v98)
                                            } else {
                                                v109 := GenS4{F0: 1, F1: -628, F2: (^(v108))}
                                                _ = v109
                                            }
                                            var v110 uint64 = (^(v108))
                                            _ = v105
                                            _ = v108
                                            _ = v110
                                        }
                                        _ = v103
                                    }
                                case "a":
                                    if (((1e10 / (1.5 * 1.5)) - ((float64(0.0) + 1.5) * 1.5)) != 1e10) {
                                        _ = len(v98)
                                        _ = len(v98)
                                        v111 := GenS14{F0: ((1e10 / 0.1) <= ((float64(1e10) - 0.0) + 1.5))}
                                        _ = v111
                                    } else {
                                        v112 := GenS15{F0: 128, F1: 2147483647}
                                        v2 += (v2 & v2)
                                        v113 := GenS16{F0: v2, F1: 1, F2: (fi94 % ((0) | 1))}
                                        _ = v112
                                        _ = v113
                                    }
                                default:
                                    var v114 string = "\n\t"
                                    v115 := GenS11{F0: 358, F1: (v2 & (v2 % (((^(v2))) | 1))), F2: 25}
                                    v114 = (v114 + ("" + v114))
                                    _ = v114
                                    _ = v115
                                }
                            case 959:
                                v2 += v2
                            case 128:
                                switch fi99 {
                                case 127:
                                    var v116 uint8 = 127
                                    _ = v116
                                }
                            }
                            v117 := v1[32768]
                            v118 := GenS17{F0: 4294967295}
                            _ = v117
                            _ = v118
                        }
                    case 32767:
                        var v119 uint64 = 459
                        _ = v119
                    }
                    _ = fi99
                    _ = v100
                }
                _ = fi94
                _ = v98
            }
        }
        _ = v92
    } else {
        v87 -= (v87 + v87)
        for fi120 := 0; fi120 < 9; fi120++ {
            switch "foo bar" {
            case "foo bar":
                for fi121 := 0; fi121 < 10; fi121++ {
                    var v122 uint16 = 2
                    _ = len(v1)
                    v3 -= v3
                    _ = fi121
                    _ = v122
                }
            case "":
                fmt.Println(v3)
            case "a":
                v123 := v1[(fi120 - 2147483647)]
                _ = v123
            default:
                var v124 uint64 = 2
                _ = v124
            }
            switch ("a" + "") {
            case "hello":
                if (!true) {
                    if false {
                        for fi125 := 0; fi125 < 19; fi125++ {
                            switch fi120 {
                            case 14:
                                v87 -= 42
                            case 128:
                                v87 -= (v87 - (v87 * (v87 >> (((v2 << ((255) & 63))) & 63))))
                                fmt.Println(v2, v3)
                                switch "foo bar" {
                                case "foo bar":
                                    v126 := GenS18{F0: (fi125 / ((int(v87)) | 1)), F1: 1, F2: true}
                                    var v127 int64 = (^(v87))
                                    _ = v126
                                    _ = v127
                                }
                                var v128 int64 = -455
                                _ = v128
                            case 42:
                                switch (fi120 * fi125) {
                                case 167:
                                    delete(v1, (fi120 | -9223372036854775808))
                                    v129 := []int{int(v2), fi120, fi125}
                                    v130 := GenS15{F0: 211, F1: -2147483648}
                                    _ = v129
                                    _ = v130
                                default:
                                    var v131 uint32 = 0
                                    _ = v131
                                }
                            }
                            _ = fi125
                        }
                        var v132 uint64 = 2
                        _ = v132
                    }
                    var v133 int32 = -2147483648
                    for fi134 := 0; fi134 < 11; fi134++ {
                        if (128 == fi120) {
                            v3 = v3
                            v87 = -1
                            fmt.Println(fi120)
                        }
                        switch fi120 {
                        case -698:
                            var v135 int = (^(fi120))
                            _ = v135
                        case -1:
                            v3 -= -110
                        case -871:
                            var v136 int = (fi120 | (fi120 / ((fi134) | 1)))
                            _ = v136
                        }
                        switch "" {
                        case "a":
                            if true {
                                delete(v1, fi134)
                                var v137 int32 = -85
                                fmt.Println(v137, v3)
                                _ = v137
                            }
                        case "foo bar":
                            if false {
                                if ((0.1 >= 1.5) && (((4294967295 <= 0) && true) || ((((float32((1e10 + 0.0)) * 3.14159) / 0.1) / 1e10) >= 1.5))) {
                                    var v138 int = (fi134 % ((fi120) | 1))
                                    if (1 > v2) {
                                        v139 := []int{(^(v138))}
                                        _ = v139
                                    }
                                    _ = v138
                                }
                            }
                        case "":
                            v140 := []int{(fi120 % (((fi120 / ((fi120) | 1))) | 1)), fi134, (fi120 / ((-2147483648) | 1))}
                            _ = v140
                        }
                        _ = fi134
                    }
                    _ = v133
                } else {
                    for fi141 := 0; fi141 < 16; fi141++ {
                        var v142 string = "a"
                        _ = fi141
                        _ = v142
                    }
                }
            default:
                _ = len(v1)
            }
            v2 += (v2 << ((v2) & 63))
            _ = fi120
        }
    }
    v143 := GenS11{F0: -404, F1: (v2 / ((v2) | 1)), F2: (^(v2))}
    var v144 uint8 = 0
    v2 -= 42
    if true {
        if true {
            v3 += 100
            if false {
                var v145 int = -454
                switch (-(v145)) {
                case -2147483648:
                    if true {
                        v144 = v144
                        var v146 uint64 = 4294967295
                        _ = v146
                    }
                case 1:
                    fmt.Println(v2)
                default:
                    for fi147 := 0; fi147 < 18; fi147++ {
                        for fi148 := 0; fi148 < 7; fi148++ {
                            _ = len(v1)
                            var v149 bool = (1.5 != 0.1)
                            switch 127 {
                            case 9223372036854775807:
                                if false {
                                    v150 := GenS10{F0: (v144 + v144), F1: 387, F2: (fi148 ^ 9223372036854775807)}
                                    _ = v150
                                } else {
                                    fmt.Println(v87, v144)
                                    v145 += (fi147 * fi147)
                                }
                            }
                            _ = fi148
                            _ = v149
                        }
                        fmt.Println(v144, v2, fi147)
                        var v151 int32 = -1
                        _ = fi147
                        _ = v151
                    }
                }
                _ = v145
            } else {
                for fi152 := 0; fi152 < 5; fi152++ {
                    fmt.Println(v2, v144)
                    switch "" {
                    case "foo bar":
                        v153 := []int{}
                        _ = v153
                    default:
                        var v154 int64 = 0
                        _ = v154
                    }
                    _ = fi152
                }
            }
            v155 := GenS6{F0: 65535, F1: v144, F2: v3}
            _ = v155
        }
        var v156 int = 32768
        _ = v156
    }
    var v157 uint = v2
    v158 := GenS19{F0: 99, F1: 0, F2: (!(true || (false || true)))}
    if (3.14159 != (((float64(-2.25) * ((1.5 + (3.14159 - -2.25)) + 0.0)) * 0.1) - 0.0)) {
        v159 := []int{9223372036854775807}
        var v160 int8 = (v3 * -58)
        _ = v159
        _ = v160
    }
    v144 += v144
    v161 := GenS3{F0: 255}
    v162 := GenS20{F0: true, F1: int8(v144)}
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v144:", v144)
    fmt.Println("v157:", v157)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v143
    _ = v144
    _ = v157
    _ = v158
    _ = v161
    _ = v162
    _ = v2
    _ = v3
    _ = v87
}
