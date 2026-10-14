package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 bool; F1 uint64; F2 uint }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int; F1 int32 }
type GenS13 struct { F0 uint; F1 int16 }
type GenS14 struct { F0 int16; F1 uint8; F2 int }
func genFn0(p0 uint, p1 int32) uint16 {
    v4 := []int{9223372036854775807, 32768, 829, 9223372036854775807}
    if false {
        v4 = append(v4, 280)
        p1 = p1
    } else {
        switch ("hello" + ("foo bar" + ((("a" + "a") + (("" + "hello") + "foo bar")) + "a"))) {
        case "hello":
            fmt.Println(p1, p1)
        case "":
            var v5 int8 = -1
            _ = v5
        }
        if (true || true) {
            v6 := GenS1{F0: uint8(p0)}
            _ = v6
        } else {
            _ = len(v4)
        }
        var v7 float32 = (3.14159 / (-2.25 * ((1e10 / 1.5) - (3.14159 + 0.0))))
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return 255
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{("foo bar" + (("" + ("hello" + "foo bar")) + "a"))}
    v3 := make(map[int]string)
    if false {
        for fi8 := 0; fi8 < 20; fi8++ {
            if false {
                var v9 int = -1
                var v10 uint32 = 4294967295
                v10 += v10
                _ = v10
                _ = v9
            } else {
                v1[(fi8 >> ((1) & 63))] = 42
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v11 uint32 = 997
                _ = v11
            }
            _ = fi8
        }
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 uint = 42
            _ = fi12
            _ = v13
        }
        if (false && (-662 <= 1000)) {
            v14 := GenS1{F0: 147}
            _ = v14
        } else {
            v15 := GenS1{F0: 255}
            var v16 int = 9223372036854775807
            _ = v15
            _ = v16
        }
    }
    if (0 != 42) {
        delete(v3, 127)
        var v17 int32 = -697
        var v18 uint32 = 2147483647
        _ = v17
        _ = v18
    }
    var v19 int = 685
    v19 -= 32767
    switch (v19 | (v19 * -436)) {
    case 9223372036854775807:
        v19 = (v19 - (v19 >> ((18446744073709551615) & 63)))
    }
    if (genFn0(18446744073709551615, 2147483647) == 32767) {
        v19 -= v19
        var v20 uint = 0
        if (1.5 != float32(1.5)) {
            for fi21 := 0; fi21 < 10; fi21++ {
                if (565 > 42) {
                    for fi22 := 0; fi22 < 19; fi22++ {
                        switch 42 {
                        case -1:
                            fmt.Println(v20)
                        }
                        var v23 string = "a"
                        switch 32767 {
                        case 128:
                            if false {
                                if (!(!false)) {
                                    delete(v1, (fi21 >> ((0) & 63)))
                                    var v24 uint64 = 1
                                    _ = v24
                                } else {
                                    v25 := GenS2{F0: 1}
                                    _ = v25
                                }
                                var v26 uint8 = 90
                                v27 := GenS2{F0: 2147483647}
                                _ = v26
                                _ = v27
                            } else {
                                var v28 bool = true
                                _ = v28
                            }
                        case 0:
                            fmt.Println(v19)
                        default:
                            if (true && false) {
                                v23 = v23
                                v23 = "foo bar"
                                switch ("foo bar" + v23) {
                                case "foo bar":
                                    switch 686 {
                                    case 2147483647:
                                        switch v23 {
                                        case "hello":
                                            if len(v2) > 0 {
                                                v2[0] = "a"
                                            }
                                        case "":
                                            v3[-964] = ((((v23 + v23) + "") + "hello") + "a")
                                        default:
                                            v29 := GenS3{F0: fi22}
                                            _ = v29
                                        }
                                    case -2147483648:
                                        fmt.Println(v19, fi22, v20)
                                    case 128:
                                        v19 -= (v19 * 543)
                                    default:
                                        if true {
                                            v23 = (("" + v23) + "\n\t")
                                            var v30 uint = (v20 - (v20 + 65535))
                                            v31 := GenS4{F0: (v20 ^ (v20 >> (((v20 << (((^(v30))) & 63))) & 63)))}
                                            v19 -= -16
                                            _ = v30
                                            _ = v31
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = (("foo bar" + v23) + v23)
                                            }
                                            _ = len(v2)
                                            if true {
                                                v32 := v3[fi22]
                                                fmt.Println(v23)
                                                _ = v32
                                            }
                                        }
                                        var v33 int8 = int8(v19)
                                        v34 := v1[fi22]
                                        v35 := GenS5{F0: 255}
                                        fmt.Println(v20)
                                        var v36 uint32 = (v34 & (v34 % (((v34 & 383)) | 1)))
                                        fmt.Println(fi21, fi22, v36)
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                }
                            }
                        }
                        _ = fi22
                        _ = v23
                    }
                } else {
                    var v37 string = "foo bar"
                    switch (fi21 >> ((v20) & 63)) {
                    case 0:
                        if false {
                            if false {
                                for fi38 := 0; fi38 < 9; fi38++ {
                                    var v39 uint32 = 727
                                    fmt.Println(v39)
                                    v40 := v3[439]
                                    switch (^(fi38)) {
                                    case 32768:
                                        v41 := v1[(fi21 - fi38)]
                                        _ = v41
                                    case 0:
                                        switch (v37 + v37) {
                                        case "\n\t":
                                            delete(v3, v19)
                                        }
                                    case 32767:
                                        fmt.Println(v20)
                                    default:
                                        if false {
                                            v37 = (v40 + (("a" + v40) + ""))
                                        }
                                    }
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                                var v42 uint32 = 2147483647
                                _ = v42
                            }
                            if (v37 > (v37 + (("" + v37) + (v37 + v37)))) {
                                switch ((v37 + "foo bar") + v37) {
                                case "":
                                    for fi43 := 0; fi43 < 19; fi43++ {
                                        var v44 string = v37
                                        v19 = fi21
                                        v45 := GenS4{F0: 239}
                                        _ = fi43
                                        _ = v44
                                        _ = v45
                                    }
                                case "hello":
                                    v19 = fi21
                                case "foo bar":
                                    var v46 bool = true
                                    _ = v46
                                default:
                                    v47 := v3[(fi21 % ((fi21) | 1))]
                                    _ = v47
                                }
                                v48 := v1[(v19 ^ v19)]
                                fmt.Println(v48, fi21, v19)
                                _ = v48
                            }
                        }
                    default:
                        switch (^(fi21)) {
                        case -843:
                            if (v19 <= (v19 / (((^(fi21))) | 1))) {
                                v49 := GenS2{F0: 4294967295}
                                _ = v49
                            } else {
                                var v50 uint8 = 203
                                v51 := GenS2{F0: uint32(v20)}
                                var v52 int64 = 42
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                        case 4294967295:
                            if (3.14159 >= (1.5 - 3.14159)) {
                                v37 = v37
                                v19 = v19
                                for fi53 := 0; fi53 < 11; fi53++ {
                                    v19 -= fi53
                                    _ = fi53
                                }
                            } else {
                                fmt.Println(v19, v37, v20)
                                v20 += (v20 << ((v20) & 63))
                                v54 := GenS4{F0: v20}
                                _ = v54
                            }
                        default:
                            for fi55 := 0; fi55 < 20; fi55++ {
                                var v56 int = (-(fi55))
                                _ = fi55
                                _ = v56
                            }
                        }
                    }
                    switch "" {
                    case "\n\t":
                        fmt.Println(v19, fi21)
                    case "hello":
                        var v57 int = (v19 << (((v20 + 42)) & 63))
                        _ = v57
                    default:
                        for fi58 := 0; fi58 < 14; fi58++ {
                            delete(v3, (fi58 << ((0) & 63)))
                            _ = fi58
                        }
                    }
                    _ = v37
                }
                for fi59 := 0; fi59 < 5; fi59++ {
                    v60 := GenS5{F0: 42}
                    switch v19 {
                    case 42:
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                    case 127:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    fmt.Println(v20)
                    _ = fi59
                    _ = v60
                }
                if (70 != 42) {
                    var v61 float64 = float64(float32((3.14159 + -2.25)))
                    _ = v61
                }
                _ = fi21
            }
            if true {
                var v62 float32 = (((((1e10 * ((float32(0.0) / 3.14159) / 0.1)) * 0.1) + 1.5) - 1e10) * (0.0 - (-2.25 / 1.0)))
                fmt.Println(v19, v19)
                _ = v62
            }
        } else {
            if (1 == v19) {
                v63 := GenS2{F0: 2}
                _ = v63
            }
            var v64 uint32 = 2147483647
            for fi65 := 0; fi65 < 13; fi65++ {
                switch (fi65 >> ((256) & 63)) {
                case 127:
                    v64 -= 683
                case 2147483647:
                    switch (fi65 & -121) {
                    case 32768:
                        var v66 uint = (v20 ^ v20)
                        _ = v66
                    case 9223372036854775807:
                        for fi67 := 0; fi67 < 10; fi67++ {
                            v20 += (v20 + (v20 % (((v20 >> ((v20) & 63))) | 1)))
                            _ = fi67
                        }
                    case 42:
                        if (!false) {
                            switch (^(fi65)) {
                            case 2147483647:
                                v64 += 2
                            case -739:
                                v19 = (fi65 % (((fi65 << ((2) & 63))) | 1))
                            case -215:
                                if (!true) {
                                    switch "hello" {
                                    case "a":
                                        v20 += v20
                                    default:
                                        var v68 int64 = -670
                                        _ = v68
                                    }
                                    for fi69 := 0; fi69 < 1; fi69++ {
                                        var v70 float64 = 0.0
                                        _ = fi69
                                        _ = v70
                                    }
                                }
                            }
                            if true {
                                var v71 uint8 = 255
                                _ = v71
                            } else {
                                for fi72 := 0; fi72 < 2; fi72++ {
                                    v20 += v20
                                    _ = fi72
                                }
                            }
                            v20 = (v20 | v20)
                        } else {
                            for fi73 := 0; fi73 < 15; fi73++ {
                                v20 = (v20 * 18446744073709551615)
                                v74 := v3[v19]
                                _ = fi73
                                _ = v74
                            }
                            var v75 uint8 = 171
                            v76 := GenS2{F0: uint32(v75)}
                            _ = v75
                            _ = v76
                        }
                    }
                case 128:
                    v64 = 65535
                }
                _ = fi65
            }
            _ = v64
        }
        _ = v20
    } else {
        v77 := v1[(v19 >> ((65535) & 63))]
        _ = v77
    }
    v19 -= (v19 / ((4294967295) | 1))
    if true {
        var v78 uint32 = 4294967295
        var v79 string = ("" + "\n\t")
        _ = v78
        _ = v79
    } else {
        var v80 int32 = -1
        v19 = -835
        v19 = 1
        _ = v80
    }
    if true {
        v81 := GenS6{F0: "\n\t", F1: 2}
        if ((3.14159 * -2.25) > 0.1) {
            for fi82 := 0; fi82 < 14; fi82++ {
                v19 += (-(v19))
                _ = fi82
            }
        } else {
            _ = len(v3)
            for fi83 := 0; fi83 < 20; fi83++ {
                v84 := GenS6{F0: "", F1: 0}
                fmt.Println(v19)
                _ = fi83
                _ = v84
            }
        }
        _ = v81
    }
    v19 += v19
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v85 int32 = -1
    if true {
        fmt.Println(v19, v19)
    }
    if ((v19 % ((245) | 1)) >= (v19 & 32767)) {
        var v86 uint32 = 763
        v19 = v19
        var v87 int16 = 1000
        _ = v86
        _ = v87
    }
    if (1 >= 0) {
        var v88 uint16 = 32767
        _ = v88
    } else {
        fmt.Println(v85)
    }
    for fi89 := 0; fi89 < 12; fi89++ {
        var v90 uint16 = 123
        v85 += (-(v85))
        _ = fi89
        _ = v90
    }
    if ((("foo bar" + "hello") + "a") > "hello") {
        v2 = append(v2, "\n\t")
        if (1 >= 42) {
            for fi91 := 0; fi91 < 7; fi91++ {
                for fi92 := 0; fi92 < 16; fi92++ {
                    v85 += v85
                    _ = fi92
                }
                v85 = (v85 / (((v85 - v85)) | 1))
                _ = fi91
            }
            switch "a" {
            case "\n\t":
                for fi93 := 0; fi93 < 6; fi93++ {
                    for fi94 := 0; fi94 < 2; fi94++ {
                        var v95 float64 = (-2.25 + float64((((((3.14159 / (0.0 - 0.1)) - 0.0) + 3.14159) / 3.14159) * float32(float64(-2.25)))))
                        switch fi93 {
                        case 32767:
                            v85 -= 116
                        case 2147483647:
                            if false {
                                v96 := GenS1{F0: 0}
                                _ = v96
                            } else {
                                fmt.Println(fi94, fi93, fi93)
                            }
                        default:
                            var v97 uint32 = 2
                            _ = v97
                        }
                        v95 += 3.14159
                        _ = fi94
                        _ = v95
                    }
                    _ = fi93
                }
            case "foo bar":
                switch 127 {
                case 1:
                    if false {
                        var v98 bool = true
                        if (v19 != v19) {
                            v3[v19] = ("" + ("" + "\n\t"))
                            fmt.Println(v19, v85, v19)
                        } else {
                            v85 = -2147483648
                        }
                        _ = v98
                    } else {
                        switch v19 {
                        case -617:
                            for fi99 := 0; fi99 < 18; fi99++ {
                                var v100 int16 = 1
                                switch (v19 & (-(v19))) {
                                case 0:
                                    var v101 uint32 = 42
                                    _ = v101
                                case 373:
                                    fmt.Println(v19)
                                default:
                                    var v102 uint = 1
                                    _ = v102
                                }
                                _ = fi99
                                _ = v100
                            }
                        case -2147483648:
                            for fi103 := 0; fi103 < 15; fi103++ {
                                v19 = fi103
                                var v104 string = "foo bar"
                                _ = fi103
                                _ = v104
                            }
                        case 32767:
                            var v105 int = v19
                            _ = v105
                        default:
                            var v106 uint8 = 171
                            _ = v106
                        }
                        v19 = (-(v19))
                    }
                }
            }
            v107 := GenS7{F0: 1}
            _ = v107
        }
        switch "\n\t" {
        case "hello":
            for fi108 := 0; fi108 < 14; fi108++ {
                v19 = (fi108 | (^(fi108)))
                _ = fi108
            }
        }
    } else {
        if true {
            fmt.Println(v19, v19)
            if ("a" < "") {
                fmt.Println(v19, v19)
                if (42 <= 4294967295) {
                    v109 := GenS8{F0: "\n\t"}
                    delete(v3, v19)
                    _ = v109
                }
                v110 := GenS9{F0: (0 >= 100), F1: 201, F2: 65535}
                _ = v110
            } else {
                for fi111 := 0; fi111 < 14; fi111++ {
                    for fi112 := 0; fi112 < 20; fi112++ {
                        delete(v3, 42)
                        if (0 > 2) {
                            var v113 uint8 = 151
                            switch "\n\t" {
                            case "foo bar":
                                var v114 int8 = -1
                                _ = v114
                            default:
                                v115 := GenS2{F0: 191}
                                _ = v115
                            }
                            _ = v113
                        }
                        switch 2147483647 {
                        case 32767:
                            switch 1 {
                            case 1:
                                var v116 uint64 = 692
                                _ = v116
                            }
                        case 4294967295:
                            v85 += v85
                        case 868:
                            v19 = v19
                        default:
                            v117 := GenS5{F0: 65535}
                            _ = v117
                        }
                        _ = fi112
                    }
                    fmt.Println(v85)
                    if false {
                        switch (v19 + -9223372036854775808) {
                        case -2147483648:
                            var v118 int = v19
                            _ = v118
                        default:
                            for fi119 := 0; fi119 < 1; fi119++ {
                                var v120 uint8 = 127
                                v121 := GenS1{F0: v120}
                                v122 := GenS3{F0: 722}
                                v120 -= v120
                                _ = fi119
                                _ = v120
                                _ = v121
                                _ = v122
                            }
                        }
                    } else {
                        fmt.Println(fi111, fi111, v19)
                        var v123 uint = 2
                        switch (v19 ^ (-(v19))) {
                        case 128:
                            var v124 uint32 = 113
                            _ = v124
                        default:
                            for fi125 := 0; fi125 < 3; fi125++ {
                                fmt.Println(v19, fi125)
                                v19 = 4294967295
                                _ = len(v1)
                                _ = fi125
                            }
                        }
                        _ = v123
                    }
                    _ = fi111
                }
            }
            var v126 uint16 = genFn0(4294967295, (v85 | v85))
            _ = v126
        }
        if false {
            var v127 int8 = -128
            _ = v127
        } else {
            for fi128 := 0; fi128 < 4; fi128++ {
                v19 = fi128
                v129 := GenS10{F0: -9223372036854775808}
                v1[(fi128 + fi128)] = 42
                _ = fi128
                _ = v129
            }
            switch "a" {
            case "a":
                fmt.Println(v85)
            case "hello":
                if false {
                    v85 = 610
                } else {
                    for fi130 := 0; fi130 < 8; fi130++ {
                        if (4294967295 > 853) {
                            fmt.Println(fi130, fi130)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v131 uint = 42
                            _ = v131
                        }
                        v85 -= (v85 - v85)
                        var v132 uint8 = 59
                        _ = fi130
                        _ = v132
                    }
                    for fi133 := 0; fi133 < 20; fi133++ {
                        fmt.Println(fi133, fi133)
                        _ = fi133
                    }
                }
            }
            for fi134 := 0; fi134 < 14; fi134++ {
                v85 = v85
                if (!(!false)) {
                    var v135 int32 = (v85 << ((42) & 63))
                    for fi136 := 0; fi136 < 3; fi136++ {
                        fmt.Println(fi134)
                        _ = fi136
                    }
                    _ = v135
                }
                v137 := GenS4{F0: 4294967295}
                _ = fi134
                _ = v137
            }
        }
        switch (v19 - v19) {
        case 179:
            var v138 uint = 0
            _ = v138
        case -1:
            if false {
                for fi139 := 0; fi139 < 14; fi139++ {
                    if true {
                        if (float32(0.1) == 0.0) {
                            if true {
                                v140 := GenS11{F0: -79}
                                _ = v140
                            } else {
                                v141 := GenS3{F0: 188}
                                if true {
                                    for fi142 := 0; fi142 < 6; fi142++ {
                                        var v143 bool = (604 != v85)
                                        if (1 > -88) {
                                            fmt.Println(fi142, v143, fi139)
                                            _ = len(v1)
                                            v144 := v3[(^(fi139))]
                                            _ = v144
                                        } else {
                                            if false {
                                                var v145 string = "hello"
                                                v146 := GenS3{F0: (fi142 << ((4294967295) & 63))}
                                                _ = v145
                                                _ = v146
                                            }
                                            if (!(!(42 != 1))) {
                                                if (((((("" + "a") + "hello") + "hello") + "a") < "") || ((2 > 2) || true)) {
                                                    v143 = v143
                                                }
                                            }
                                        }
                                        var v147 uint8 = 246
                                        _ = fi142
                                        _ = v143
                                        _ = v147
                                    }
                                    var v148 int16 = -586
                                    if (18446744073709551615 < 876) {
                                        v19 -= 32767
                                        v19 = (^(v19))
                                        switch 2147483647 {
                                        case -2147483648:
                                            var v149 float64 = 1.5
                                            _ = v149
                                        case -1:
                                            _ = len(v2)
                                        case 128:
                                            v148 = (v148 / ((v148) | 1))
                                        }
                                    }
                                    _ = v148
                                } else {
                                    var v150 int16 = 32767
                                    var v151 uint32 = 1
                                    _ = v150
                                    _ = v151
                                }
                                fmt.Println(v19, fi139)
                                _ = v141
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v19 += (v19 ^ (^(fi139)))
                        }
                        var v152 float32 = 0.0
                        _ = v152
                    }
                    var v153 string = ("a" + "\n\t")
                    fmt.Println(v153, v153)
                    _ = fi139
                    _ = v153
                }
                switch v19 {
                case -1:
                    if false {
                        v19 -= (v19 ^ 42)
                        v154 := GenS10{F0: -1}
                        var v155 uint32 = 464
                        _ = v154
                        _ = v155
                    }
                case 42:
                    for fi156 := 0; fi156 < 3; fi156++ {
                        var v157 int64 = -9223372036854775808
                        if false {
                            if true {
                                v157 = v157
                                v158 := GenS7{F0: v157}
                                switch (v19 >> ((4294967295) & 63)) {
                                case -2147483648:
                                    v19 += (fi156 << ((42) & 63))
                                case 92:
                                    for fi159 := 0; fi159 < 16; fi159++ {
                                        v85 = (v85 ^ (-(v85)))
                                        _ = fi159
                                    }
                                case 32767:
                                    var v160 float32 = 0.1
                                    _ = v160
                                default:
                                    for fi161 := 0; fi161 < 10; fi161++ {
                                        var v162 float32 = float32(1.5)
                                        v162 = ((v162 + (v162 + v162)) - ((3.14159 + (-2.25 / v162)) * float32(1.5)))
                                        _ = fi161
                                        _ = v162
                                    }
                                }
                                _ = v158
                            } else {
                                v163 := GenS1{F0: 42}
                                var v164 uint32 = 808
                                if (false && (((v164 >> ((42) & 63)) != 319) && true)) {
                                    switch (-(fi156)) {
                                    case -1:
                                        var v165 float32 = 0.1
                                        _ = v165
                                    case -711:
                                        fmt.Println(fi156)
                                    default:
                                        v166 := GenS3{F0: 9223372036854775807}
                                        _ = v166
                                    }
                                } else {
                                    for fi167 := 0; fi167 < 18; fi167++ {
                                        var v168 uint64 = 18446744073709551615
                                        v169 := GenS10{F0: -310}
                                        v170 := GenS12{F0: (-(v19)), F1: (v85 / ((42) | 1))}
                                        _ = fi167
                                        _ = v168
                                        _ = v169
                                        _ = v170
                                    }
                                    for fi171 := 0; fi171 < 3; fi171++ {
                                        fmt.Println(v19, v157)
                                        v85 = v85
                                        fmt.Println(v164, v19, fi171)
                                        _ = fi171
                                    }
                                    v164 = (v164 % ((557) | 1))
                                }
                                _ = v163
                                _ = v164
                            }
                        } else {
                            if true {
                                switch "foo bar" {
                                case "\n\t":
                                    var v172 uint32 = 944
                                    _ = v172
                                default:
                                    _ = len(v3)
                                }
                                var v173 bool = ((!false) || false)
                                _ = v173
                            } else {
                                for fi174 := 0; fi174 < 15; fi174++ {
                                    var v175 uint = 42
                                    _ = fi174
                                    _ = v175
                                }
                                v176 := GenS10{F0: (v157 * v157)}
                                if false {
                                    v85 = v85
                                    if false {
                                        for fi177 := 0; fi177 < 2; fi177++ {
                                            if false {
                                                if false {
                                                    fmt.Println(fi177, fi156, v19)
                                                    var v178 int32 = v85
                                                    _ = v178
                                                } else {
                                                    v19 += 544
                                                    if true {
                                                        v85 -= int32(v157)
                                                        fmt.Println(fi177, fi177, v85)
                                                    } else {
                                                        switch "a" {
                                                        case "a":
                                                            fmt.Println(v19, v19)
                                                        }
                                                        var v179 int = fi156
                                                        v179 = (v19 * (v19 >> ((448) & 63)))
                                                        _ = v179
                                                    }
                                                }
                                                if (!false) {
                                                    var v180 bool = (2 <= 140)
                                                    _ = v180
                                                }
                                                v85 = v85
                                            } else {
                                                if (false || ((false || ((v85 != v85) && false)) && ("\n\t" <= ((("" + "hello") + "\n\t") + "")))) {
                                                    v181 := GenS1{F0: 1}
                                                    _ = v181
                                                }
                                                var v182 int32 = v85
                                                _ = v182
                                            }
                                            v183 := GenS13{F0: 255, F1: 42}
                                            var v184 bool = false
                                            v85 -= v85
                                            fmt.Println(v85)
                                            _ = fi177
                                            _ = v183
                                            _ = v184
                                        }
                                        v19 += 556
                                    }
                                    switch v19 {
                                    case 4294967295:
                                        var v185 bool = (false && false)
                                        _ = v185
                                    }
                                }
                                _ = v176
                            }
                        }
                        _ = fi156
                        _ = v157
                    }
                }
            } else {
                v186 := GenS14{F0: 42, F1: 133, F2: 0}
                v187 := GenS11{F0: -1}
                _ = v186
                _ = v187
            }
        default:
            v85 = (-(v85))
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v85:", v85)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v85
}
