package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int8; F2 int }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 int; F1 int16; F2 uint16 }
type GenS8 struct { F0 int32; F1 uint16; F2 int }
type GenS9 struct { F0 bool; F1 string; F2 uint64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int32; F1 uint; F2 int16 }
type GenS12 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint; F1 uint64; F2 int64 }
type GenS15 struct { F0 int }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 2
    var v3 int = -691
    if false {
        fmt.Println(v3)
        if (false && (("foo bar" + "\n\t") > "hello")) {
            if false {
                fmt.Println(v3)
                v4 := GenS1{F0: -1, F1: 42}
                v2 = v2
                _ = v4
            }
            var v5 uint16 = 42
            _ = v5
        } else {
            v6 := GenS2{F0: -38}
            var v7 int8 = 127
            switch "a" {
            case "hello":
                for fi8 := 0; fi8 < 1; fi8++ {
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v3, v2)
                    case "hello":
                        v9 := GenS1{F0: -204, F1: -1}
                        _ = v9
                    }
                    v10 := []int{32768, v3, (v3 - -2147483648)}
                    v11 := GenS1{F0: -712, F1: 1}
                    _ = fi8
                    _ = v10
                    _ = v11
                }
            }
            _ = v6
            _ = v7
        }
    } else {
        if true {
            fmt.Println(v3, v3)
        } else {
            switch 0 {
            case 579:
                var v12 uint8 = 42
                _ = v12
            case 296:
                v2 = (v2 * v2)
            default:
                v13 := []int{-32, (v3 ^ (v3 >> ((v2) & 63))), v3}
                _ = v13
            }
            if ((true && (false && true)) || (("hello" + "a") <= ((("\n\t" + "hello") + "foo bar") + ("foo bar" + ("\n\t" + ""))))) {
                if false {
                    if true {
                        switch "" {
                        case "\n\t":
                            var v14 uint8 = 127
                            _ = v14
                        case "foo bar":
                            v15 := []int{v3, v3}
                            _ = v15
                        }
                        fmt.Println(v2, v2)
                        v1[v3] = (v3 - 9223372036854775807)
                    } else {
                        v3 -= v3
                    }
                    v3 = (v3 - v3)
                    v2 = (v2 ^ v2)
                } else {
                    switch v3 {
                    case 0:
                        v3 += 9223372036854775807
                    default:
                        v16 := v1[v3]
                        _ = v16
                    }
                    if (-2147483648 < (v3 & 9223372036854775807)) {
                        delete(v1, (v3 >> ((65535) & 63)))
                        if (-81 < -1) {
                            v2 = 2
                        } else {
                            for fi17 := 0; fi17 < 20; fi17++ {
                                var v18 int32 = 2147483647
                                _ = fi17
                                _ = v18
                            }
                            switch v3 {
                            case -2147483648:
                                var v19 int64 = -1
                                _ = v19
                            case -9223372036854775808:
                                for fi20 := 0; fi20 < 16; fi20++ {
                                    if true {
                                        switch "hello" {
                                        case "a":
                                            switch -740 {
                                            case -165:
                                                if (724 < 4294967295) {
                                                    if false {
                                                        if (((((((("foo bar" + ("" + "")) + "foo bar") + "foo bar") + "a") + "\n\t") + "foo bar") + "hello") != "hello") {
                                                            for fi21 := 0; fi21 < 8; fi21++ {
                                                                var v22 uint32 = 1
                                                                v23 := []int{9223372036854775807, (v3 - fi20), fi20, fi20}
                                                                var v24 int64 = -1
                                                                _ = fi21
                                                                _ = v22
                                                                _ = v23
                                                                _ = v24
                                                            }
                                                        }
                                                        if true {
                                                            switch "" {
                                                            case "":
                                                                _ = len(v1)
                                                            default:
                                                                v25 := GenS2{F0: -761}
                                                                _ = v25
                                                            }
                                                            var v26 float32 = 1.5
                                                            _ = v26
                                                        } else {
                                                            for fi27 := 0; fi27 < 17; fi27++ {
                                                                var v28 int = v3
                                                                var v29 float64 = 1.5
                                                                _ = fi27
                                                                _ = v28
                                                                _ = v29
                                                            }
                                                            switch (("" + ("a" + (("\n\t" + "foo bar") + ("a" + ("a" + "\n\t"))))) + ("" + "\n\t")) {
                                                            case "hello":
                                                                if true {
                                                                    var v30 int16 = 42
                                                                    v2 -= (v2 ^ (v2 >> ((4294967295) & 63)))
                                                                    var v31 uint8 = 255
                                                                    _ = v30
                                                                    _ = v31
                                                                } else {
                                                                    var v32 uint64 = 2
                                                                    _ = v32
                                                                }
                                                            case "a":
                                                                v3 = fi20
                                                            }
                                                            for fi33 := 0; fi33 < 12; fi33++ {
                                                                v2 = v2
                                                                if (1 > 398) {
                                                                    var v34 string = "foo bar"
                                                                    var v35 string = (("foo bar" + v34) + "")
                                                                    _ = v34
                                                                    _ = v35
                                                                }
                                                                var v36 int8 = 116
                                                                _ = fi33
                                                                _ = v36
                                                            }
                                                        }
                                                    } else {
                                                        v37 := GenS1{F0: 42, F1: 32767}
                                                        var v38 string = (("a" + "\n\t") + ("" + ((("hello" + "foo bar") + ("foo bar" + ("\n\t" + ("a" + "")))) + "foo bar")))
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                    v2 += 592
                                                } else {
                                                    var v39 int64 = 9223372036854775807
                                                    var v40 bool = (1e10 > (((1.5 / 0.1) / 1.5) + (float64(((((3.14159 + 1e10) * 1e10) * 1.5) / 1.0)) / 1e10)))
                                                    _ = v39
                                                    _ = v40
                                                }
                                            }
                                        case "\n\t":
                                            switch ("foo bar" + "") {
                                            case "":
                                                var v41 uint64 = 363
                                                _ = v41
                                            case "a":
                                                switch 128 {
                                                case 1:
                                                    for fi42 := 0; fi42 < 2; fi42++ {
                                                        if (15 < 1) {
                                                            fmt.Println(fi42, fi42, v2)
                                                        }
                                                        _ = fi42
                                                    }
                                                }
                                            default:
                                                v43 := GenS3{F0: 12}
                                                _ = v43
                                            }
                                        case "":
                                            var v44 uint16 = 2
                                            _ = v44
                                        default:
                                            v45 := []int{v3}
                                            _ = v45
                                        }
                                        switch (^(v3)) {
                                        case 128:
                                            v1[v3] = fi20
                                        }
                                    }
                                    _ = fi20
                                }
                            }
                            _ = len(v1)
                        }
                    }
                    if (!(false || true)) {
                        for fi46 := 0; fi46 < 14; fi46++ {
                            var v47 uint64 = 770
                            v2 = v2
                            _ = fi46
                            _ = v47
                        }
                    } else {
                        fmt.Println(v2, v2)
                    }
                }
            }
        }
        var v48 uint8 = 127
        var v49 uint8 = 0
        _ = v48
        _ = v49
    }
    for fi50 := 0; fi50 < 5; fi50++ {
        v3 += -2147483648
        if (false && ((("hello" + "\n\t") + (((("\n\t" + ("" + "foo bar")) + "hello") + ("" + "a")) + "foo bar")) == "")) {
            if true {
                var v51 int64 = -606
                var v52 int = fi50
                _ = v51
                _ = v52
            } else {
                v3 -= v3
            }
            fmt.Println(v3, fi50, v3)
            for fi53 := 0; fi53 < 8; fi53++ {
                v54 := GenS4{F0: 2147483647, F1: -20, F2: fi50}
                v3 = -2147483648
                _ = fi53
                _ = v54
            }
        } else {
            var v55 int16 = -241
            _ = v55
        }
        if ((float64(0.1) / ((1.5 / 1.0) / (float64(3.14159) - (0.1 * 1.5)))) != (((float64(((float32(1e10) + 0.0) - 3.14159)) * (((0.0 * -2.25) - float64((-2.25 * float32(1e10)))) * 1e10)) * float64(3.14159)) * (3.14159 / (0.0 - (float64(float32(-2.25)) / (1.5 / (3.14159 + 0.0))))))) {
            v2 = 2
        }
        _ = fi50
    }
    var v56 int = 144
    if true {
        var v57 uint8 = 141
        v58 := v1[v56]
        var v59 uint8 = 127
        _ = v57
        _ = v58
        _ = v59
    }
    fmt.Println(v3)
    switch "hello" {
    case "":
        if true {
            if (42 == 2) {
                var v60 uint8 = 128
                v61 := v1[(v3 ^ -549)]
                _ = v60
                _ = v61
            }
            v62 := GenS2{F0: 1000}
            if ((0.0 - 0.1) > float64(0.1)) {
                v63 := GenS5{F0: "\n\t", F1: ("hello" + "")}
                _ = v63
            } else {
                v64 := v1[(v3 ^ v56)]
                _ = v64
            }
            _ = v62
        } else {
            v65 := GenS6{F0: 0, F1: (!(65535 != 255))}
            _ = len(v1)
            _ = v65
        }
    case "a":
        for fi66 := 0; fi66 < 16; fi66++ {
            v56 = -2147483648
            for fi67 := 0; fi67 < 9; fi67++ {
                var v68 int8 = 0
                var v69 string = (("hello" + "\n\t") + ("hello" + "\n\t"))
                _ = fi67
                _ = v68
                _ = v69
            }
            for fi70 := 0; fi70 < 13; fi70++ {
                v71 := GenS1{F0: 0, F1: 1}
                if (true && ((!(("a" < "hello") && (-1 > -2147483648))) || false)) {
                    v72 := []int{v3}
                    if (2 < 2) {
                        var v73 int32 = -1
                        _ = v73
                    }
                    _ = v72
                } else {
                    switch "" {
                    case "":
                        if (0.1 < (3.14159 * 1.5)) {
                            var v74 uint32 = 2
                            _ = v74
                        }
                    case "\n\t":
                        v75 := GenS1{F0: 57, F1: 32767}
                        _ = v75
                    case "hello":
                        if (0 > 2147483648) {
                            v76 := []int{v3, (v56 | (fi70 | (v3 & fi70))), -272, v3}
                            _ = v76
                        } else {
                            var v77 uint8 = 128
                            _ = v77
                        }
                    default:
                        var v78 uint = v2
                        _ = v78
                    }
                    var v79 int16 = 42
                    fmt.Println(v2, v56)
                    _ = v79
                }
                switch (fi66 ^ v3) {
                case 2147483647:
                    v3 = 1
                case -9223372036854775808:
                    switch "hello" {
                    case "a":
                        var v80 float64 = 0.1
                        _ = v80
                    default:
                        var v81 int32 = -2147483648
                        _ = v81
                    }
                case -910:
                    if (239 >= 2147483648) {
                        var v82 uint = v2
                        v3 += fi66
                        fmt.Println(v82, fi66, v3)
                        _ = v82
                    }
                default:
                    var v83 int8 = -59
                    _ = v83
                }
                _ = fi70
                _ = v71
            }
            _ = fi66
        }
    case "\n\t":
        v56 += (v56 & v3)
    }
    v84 := GenS2{F0: 422}
    switch 1 {
    case -384:
        fmt.Println(v3, v2, v3)
    default:
        v2 -= v2
    }
    switch "\n\t" {
    case "foo bar":
        for fi85 := 0; fi85 < 3; fi85++ {
            for fi86 := 0; fi86 < 11; fi86++ {
                fmt.Println(v2, v3, v56)
                v87 := GenS5{F0: ("a" + "a"), F1: "\n\t"}
                switch "foo bar" {
                case "hello":
                    if ((!true) || true) {
                        var v88 float32 = float32(1e10)
                        switch "\n\t" {
                        case "hello":
                            var v89 uint8 = 255
                            _ = v89
                        case "a":
                            v90 := v1[(fi85 | fi86)]
                            _ = v90
                        default:
                            fmt.Println(v3, v56)
                        }
                        v91 := GenS3{F0: 1}
                        _ = v88
                        _ = v91
                    } else {
                        v92 := GenS7{F0: (fi85 + fi85), F1: 423, F2: 42}
                        if true {
                            v3 += (fi85 | -1)
                        }
                        v93 := GenS2{F0: 0}
                        _ = v92
                        _ = v93
                    }
                case "":
                    if ((!(1 <= -901)) || ((0 < -9223372036854775808) && false)) {
                        fmt.Println(v2)
                        var v94 int = (fi85 << (((v2 & v2)) & 63))
                        v95 := GenS6{F0: 1, F1: false}
                        _ = v94
                        _ = v95
                    } else {
                        switch ("hello" + "hello") {
                        case "\n\t":
                            var v96 string = (("a" + "hello") + "")
                            _ = v96
                        case "hello":
                            v97 := []int{}
                            _ = v97
                        case "foo bar":
                            v56 -= (v3 / ((2147483647) | 1))
                        }
                    }
                }
                _ = fi86
                _ = v87
            }
            if true {
                fmt.Println(v56, v2, fi85)
                switch (fi85 * 42) {
                case -2147483648:
                    if true {
                        var v98 int16 = 42
                        fmt.Println(v3, v56, v2)
                        _ = v98
                    } else {
                        for fi99 := 0; fi99 < 20; fi99++ {
                            if false {
                                switch 265 {
                                case 0:
                                    v56 = (fi99 >> (((v2 | v2)) & 63))
                                case 649:
                                    var v100 int64 = -1
                                    v101 := GenS8{F0: 42, F1: 334, F2: fi85}
                                    fmt.Println(v2, v2)
                                    _ = v100
                                    _ = v101
                                case 9223372036854775807:
                                    if ((fi99 ^ fi99) >= (fi99 * v56)) {
                                        var v102 int = (v3 / (((v3 + 128)) | 1))
                                        if (("\n\t" + "\n\t") <= "a") {
                                            var v103 int16 = 42
                                            if true {
                                                var v104 bool = (1 >= 0)
                                                var v105 uint32 = 650
                                                _ = v104
                                                _ = v105
                                            }
                                            switch (fi99 / ((v56) | 1)) {
                                            case 42:
                                                v56 += (-(v56))
                                            case 32767:
                                                v3 = fi85
                                            case 127:
                                                var v106 uint8 = 186
                                                _ = v106
                                            }
                                            v56 -= v56
                                            v107 := []int{32768}
                                            if false {
                                                switch "hello" {
                                                case "\n\t":
                                                    var v108 int64 = 233
                                                    _ = v108
                                                default:
                                                    var v109 uint64 = 42
                                                    _ = v109
                                                }
                                                v102 = v102
                                                if (!(0 >= -37)) {
                                                    v3 = v56
                                                    switch -2147483648 {
                                                    case 32767:
                                                        var v110 int16 = 585
                                                        _ = v110
                                                    case 444:
                                                        var v111 int32 = 1
                                                        _ = v111
                                                    case 166:
                                                        var v112 string = ((("" + ("a" + "hello")) + "foo bar") + ("\n\t" + ("" + "")))
                                                        _ = v112
                                                    default:
                                                        var v113 uint16 = 555
                                                        _ = v113
                                                    }
                                                }
                                            } else {
                                                if false {
                                                    if ((v103 & v103) > (v103 % ((v103) | 1))) {
                                                        if (510 <= 2147483647) {
                                                            if (0 >= -19) {
                                                                var v114 int = (fi85 | int(v2))
                                                                v115 := GenS7{F0: (fi99 / (((v102 & v114)) | 1)), F1: (v103 ^ v103), F2: 439}
                                                                v116 := GenS2{F0: v103}
                                                                _ = v114
                                                                _ = v115
                                                                _ = v116
                                                            }
                                                            switch "a" {
                                                            case "":
                                                                v103 = (-(v103))
                                                            default:
                                                                if len(v107) > 0 {
                                                                    v107[0] = -1
                                                                }
                                                            }
                                                            var v117 uint8 = 42
                                                            _ = v117
                                                        } else {
                                                            v103 += v103
                                                            if true {
                                                                if (("\n\t" + "foo bar") != ("a" + ("hello" + ((("" + "a") + "a") + "hello")))) {
                                                                    if true {
                                                                        v118 := GenS9{F0: false, F1: "a", F2: 2}
                                                                        if false {
                                                                            var v119 uint16 = 667
                                                                            v56 = fi85
                                                                            _ = v119
                                                                        } else {
                                                                            v120 := GenS10{F0: 65535}
                                                                            var v121 int8 = 0
                                                                            v102 -= (v56 - (v56 + fi85))
                                                                            _ = v120
                                                                            _ = v121
                                                                        }
                                                                        v56 -= (v3 << ((1) & 63))
                                                                        _ = v118
                                                                    }
                                                                    fmt.Println(v3)
                                                                    var v122 float64 = (1.5 - 3.14159)
                                                                    _ = v122
                                                                }
                                                                fmt.Println(fi85, v103, v2)
                                                            }
                                                            var v123 uint64 = 641
                                                            _ = v123
                                                        }
                                                        switch (v102 + -128) {
                                                        case 0:
                                                            var v124 float32 = (0.1 - 3.14159)
                                                            _ = v124
                                                        case 32768:
                                                            v125 := GenS3{F0: 241}
                                                            _ = v125
                                                        case 32767:
                                                            v126 := v1[-141]
                                                            _ = v126
                                                        }
                                                    }
                                                    switch "a" {
                                                    case "foo bar":
                                                        v2 = (v2 + (^(v2)))
                                                    case "a":
                                                        var v127 uint32 = 42
                                                        _ = v127
                                                    case "":
                                                        v103 = (v103 / (((-(v103))) | 1))
                                                        if ((v2 >> ((v2) & 63)) == 0) {
                                                            var v128 int = 2147483647
                                                            v107 = append(v107, (-(v128)))
                                                            var v129 string = ("a" + "\n\t")
                                                            if (false || true) {
                                                                var v130 int32 = 122
                                                                if (false || false) {
                                                                    if true {
                                                                        var v131 int16 = -673
                                                                        if false {
                                                                            var v132 uint8 = 34
                                                                            var v133 int8 = 42
                                                                            var v134 int = fi85
                                                                            _ = v132
                                                                            _ = v133
                                                                            _ = v134
                                                                        }
                                                                        _ = v131
                                                                    }
                                                                    fmt.Println(v2)
                                                                } else {
                                                                    if (fi99 != (^(v3))) {
                                                                        _ = len(v1)
                                                                    }
                                                                }
                                                                _ = v130
                                                            }
                                                            v2 += 18446744073709551615
                                                            _ = v128
                                                            _ = v129
                                                        } else {
                                                            v107 = append(v107, (v102 + fi99))
                                                            v135 := GenS7{F0: fi85, F1: v103, F2: 0}
                                                            _ = v135
                                                        }
                                                        v103 = v103
                                                        var v136 int64 = 0
                                                        var v137 uint = v2
                                                        _ = v136
                                                        _ = v137
                                                    default:
                                                        v103 += (v103 | 1)
                                                    }
                                                    var v138 int16 = -934
                                                    _ = v138
                                                } else {
                                                    var v139 int8 = 15
                                                    v2 += (^(v2))
                                                    if true {
                                                        v1[181] = fi85
                                                    }
                                                    _ = v139
                                                }
                                                var v140 int8 = 100
                                                _ = v140
                                            }
                                            _ = v103
                                            _ = v107
                                        } else {
                                            var v141 int8 = 1
                                            switch -232 {
                                            case 32767:
                                                v142 := GenS4{F0: 0, F1: v141, F2: fi99}
                                                _ = v142
                                            default:
                                                v143 := []int{fi99, fi85}
                                                _ = v143
                                            }
                                            _ = v141
                                        }
                                        _ = len(v1)
                                        v144 := []int{(fi99 >> (((v2 & v2)) & 63)), 4294967295}
                                        _ = v102
                                        _ = v144
                                    } else {
                                        switch "a" {
                                        case "\n\t":
                                            if (!(!true)) {
                                                _ = len(v1)
                                                v3 = (v56 - fi85)
                                                v3 = 32767
                                            } else {
                                                switch v56 {
                                                case 128:
                                                    var v145 int16 = 216
                                                    _ = v145
                                                }
                                            }
                                        case "a":
                                            v56 = (-(v3))
                                        }
                                        v3 += -1
                                        v2 = v2
                                    }
                                }
                            } else {
                                v3 = (v56 >> ((v2) & 63))
                                var v146 uint = (v2 >> ((v2) & 63))
                                var v147 int32 = 176
                                fmt.Println(v2)
                                v148 := GenS11{F0: v147, F1: v146, F2: 0}
                                _ = v146
                                _ = v147
                                _ = v148
                            }
                            _ = fi99
                        }
                    }
                case 128:
                    v2 = (v2 - (v2 % (((^(v2))) | 1)))
                }
                v149 := GenS9{F0: false, F1: (("foo bar" + (("" + ("\n\t" + (("a" + ((("hello" + (("a" + "\n\t") + ("foo bar" + "foo bar"))) + ("hello" + "\n\t")) + "a")) + "hello"))) + "foo bar")) + "\n\t"), F2: 4294967295}
                _ = v149
            } else {
                v150 := GenS11{F0: -2147483648, F1: (v2 >> ((v2) & 63)), F2: int16(v2)}
                if false {
                    switch "\n\t" {
                    case "foo bar":
                        v3 -= (v3 >> (((v2 - 0)) & 63))
                    default:
                        if true {
                            for fi151 := 0; fi151 < 17; fi151++ {
                                switch (^(v3)) {
                                case 4294967295:
                                    v56 = 0
                                case 1:
                                    v56 -= (v3 << ((v2) & 63))
                                }
                                _ = fi151
                            }
                        }
                    }
                } else {
                    v56 = (v56 + (fi85 - (fi85 / ((v3) | 1))))
                }
                v56 -= fi85
                _ = v150
            }
            if (!(-2.25 <= (float32(1.5) + 1e10))) {
                switch 32767 {
                case 127:
                    if ((v2 | (v2 & (v2 % (((^(v2))) | 1)))) < (v2 << (((^(v2))) & 63))) {
                        var v152 int16 = 1
                        _ = v152
                    } else {
                        v56 = (v56 - v56)
                        var v153 float32 = (1.5 + 3.14159)
                        _ = v153
                    }
                case 9223372036854775807:
                    fmt.Println(v3, v2, fi85)
                case 32767:
                    v56 = v3
                default:
                    if false {
                        v1[(v3 * 127)] = v3
                    }
                }
                for fi154 := 0; fi154 < 8; fi154++ {
                    if (!(588 != 1)) {
                        fmt.Println(v2)
                        var v155 int16 = -32768
                        if true {
                            var v156 int32 = 188
                            _ = v156
                        } else {
                            if false {
                                if (false || (42 <= 9223372036854775807)) {
                                    v157 := []int{fi85}
                                    _ = v157
                                }
                                v56 -= 32767
                                v155 = (v155 * (v155 & (-(v155))))
                                var v158 bool = ((true || (1e10 < 1e10)) || true)
                                _ = v158
                            } else {
                                var v159 bool = false
                                if false {
                                    v160 := GenS12{F0: 65535, F1: -9223372036854775808, F2: 65535}
                                    v161 := GenS13{F0: "\n\t"}
                                    delete(v1, 2147483647)
                                    switch (fi154 & fi85) {
                                    case -1:
                                        if v159 {
                                            var v162 uint8 = 128
                                            var v163 int64 = -242
                                            _ = len(v1)
                                            v164 := GenS1{F0: v155, F1: 1000}
                                            _ = v162
                                            _ = v163
                                            _ = v164
                                        } else {
                                            var v165 int8 = 100
                                            fmt.Println(v3, v3, v165)
                                            switch fi85 {
                                            case 32767:
                                                v166 := []int{int(v155), 2147483647}
                                                _ = v166
                                            case 2147483647:
                                                fmt.Println(fi154, fi85, v3)
                                            }
                                            switch "hello" {
                                            case "\n\t":
                                                v159 = v159
                                            case "hello":
                                                switch "hello" {
                                                case "":
                                                    var v167 uint32 = 2147483648
                                                    _ = v167
                                                case "\n\t":
                                                    v168 := GenS11{F0: -2147483648, F1: v2, F2: v155}
                                                    _ = v168
                                                case "hello":
                                                    if true {
                                                        if v159 {
                                                            if v159 {
                                                                fmt.Println(v159, v159, v3)
                                                                v169 := GenS14{F0: v2, F1: 313, F2: 416}
                                                                v2 = (^(v2))
                                                                _ = v169
                                                            } else {
                                                                var v170 uint8 = 255
                                                                switch (fi154 / ((fi154) | 1)) {
                                                                case -274:
                                                                    if (!(!v159)) {
                                                                        v170 = 225
                                                                        v2 -= (v2 + v2)
                                                                        if ((446 >= 997) || v159) {
                                                                            if (true || (-146 <= v3)) {
                                                                                var v171 uint64 = 424
                                                                                _ = v171
                                                                            } else {
                                                                                v172 := []int{}
                                                                                var v173 bool = (-9223372036854775808 <= 9223372036854775807)
                                                                                _ = v172
                                                                                _ = v173
                                                                            }
                                                                            var v174 uint = v2
                                                                            _ = v174
                                                                        } else {
                                                                            var v175 uint = (^(v2))
                                                                            v56 += 32767
                                                                            _ = v175
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v2, v159, v155)
                                                                        if (float64(0.0) <= 0.1) {
                                                                            fmt.Println(v155, fi154, v155)
                                                                            v56 = (^(fi154))
                                                                        }
                                                                        v165 = v165
                                                                    }
                                                                }
                                                                _ = v170
                                                            }
                                                            if false {
                                                                if (!v159) {
                                                                    if (0 == 398) {
                                                                        if v159 {
                                                                            var v176 int64 = -9223372036854775808
                                                                            _ = v176
                                                                        }
                                                                        var v177 float64 = -2.25
                                                                        v2 = v2
                                                                        _ = v177
                                                                    } else {
                                                                        v178 := []int{128}
                                                                        _ = v178
                                                                    }
                                                                    var v179 string = (((("" + "\n\t") + "a") + "hello") + "a")
                                                                    _ = v179
                                                                }
                                                                var v180 int16 = (v155 << (((v2 << ((v2) & 63))) & 63))
                                                                var v181 int32 = 1
                                                                v3 = (fi154 << (((v2 | v2)) & 63))
                                                                _ = v180
                                                                _ = v181
                                                            }
                                                            switch (v3 | (fi85 << (((v2 << (((v2 << ((294) & 63))) & 63))) & 63))) {
                                                            case -594:
                                                                v182 := []int{int(v165), -703, (v56 >> (((v2 | v2)) & 63))}
                                                                _ = v182
                                                            }
                                                        } else {
                                                            switch fi85 {
                                                            case 9223372036854775807:
                                                                var v183 int32 = -1
                                                                _ = v183
                                                            case 2147483647:
                                                                var v184 uint16 = 1
                                                                _ = v184
                                                            }
                                                            fmt.Println(v159, v2)
                                                            v185 := GenS15{F0: (-(fi85))}
                                                            _ = v185
                                                        }
                                                    }
                                                    fmt.Println(v165)
                                                default:
                                                    fmt.Println(v155, v3)
                                                }
                                                v186 := []int{42, (v56 * 0)}
                                                if ((float32((-2.25 - -2.25)) / float32((((float64(0.1) - 1.5) * float64(-2.25)) / 1.5))) != 3.14159) {
                                                    v187 := GenS13{F0: "a"}
                                                    _ = v187
                                                }
                                                _ = v186
                                            default:
                                                if v159 {
                                                    fmt.Println(v3, v56)
                                                }
                                                _ = len(v1)
                                                fmt.Println(v165, fi85)
                                            }
                                            _ = v165
                                        }
                                    default:
                                        var v188 float32 = 0.0
                                        _ = v188
                                    }
                                    delete(v1, (^(v3)))
                                    _ = v160
                                    _ = v161
                                } else {
                                    var v189 float64 = 1.5
                                    v56 = (v3 & 731)
                                    _ = v189
                                }
                                if v159 {
                                    if (!((!(42 > 1)) && ((fi154 % ((int(v2)) | 1)) > (v56 & 23)))) {
                                        fmt.Println(v159, v155, v56)
                                    } else {
                                        delete(v1, (fi85 + v3))
                                    }
                                    var v190 uint32 = 1
                                    _ = v190
                                }
                                _ = v159
                            }
                            if false {
                                fmt.Println(fi154, fi85, fi85)
                                var v191 bool = false
                                _ = v191
                            }
                        }
                        _ = v155
                    }
                    var v192 uint32 = 2
                    _ = fi154
                    _ = v192
                }
                if ((12 > 255) && ((0 != 420) || false)) {
                    v193 := GenS11{F0: 1, F1: (v2 * (v2 % (((v2 % ((v2) | 1))) | 1))), F2: 0}
                    v2 -= (v2 & (v2 + (v2 % (((v2 & 255)) | 1))))
                    _ = v193
                } else {
                    switch (("hello" + "\n\t") + "a") {
                    case "":
                        v194 := v1[316]
                        _ = v194
                    default:
                        v195 := []int{42, v3, v3}
                        _ = v195
                    }
                    var v196 float32 = 0.1
                    v197 := []int{}
                    _ = v196
                    _ = v197
                }
            } else {
                v198 := GenS4{F0: -2147483648, F1: 127, F2: (-(v3))}
                _ = v198
            }
            _ = fi85
        }
    case "hello":
        for fi199 := 0; fi199 < 6; fi199++ {
            var v200 string = "\n\t"
            for fi201 := 0; fi201 < 14; fi201++ {
                v2 = v2
                var v202 uint8 = 255
                _ = fi201
                _ = v202
            }
            _ = fi199
            _ = v200
        }
    default:
        v203 := []int{}
        _ = v203
    }
    var v204 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v204:", v204)
    fmt.Println("v3:", v3)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v204
    _ = v3
    _ = v56
    _ = v84
}
