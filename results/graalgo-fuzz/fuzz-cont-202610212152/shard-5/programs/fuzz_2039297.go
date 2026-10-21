package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int16; F1 int16; F2 int32 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 int; F1 int8; F2 uint8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint64; F1 string }
type GenS11 struct { F0 int8; F1 int64 }
type GenS12 struct { F0 bool; F1 int }
type GenS13 struct { F0 uint32 }
func genFn0(p0 int8, p1 int16) int8 {
    fmt.Println(p1)
    if (!(42 > 127)) {
        var v7 int = 2147483647
        v7 -= v7
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := GenS1{F0: 16}
            v11 := []int{fi8, fi9, (fi8 << ((110) & 63)), 2147483647}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    if (18446744073709551615 < 18446744073709551615) {
        var v12 int8 = 42
        _ = v12
    }
    _ = p0
    _ = p1
    return (p0 << ((0) & 63))
}
func genFn1() int64 {
    if true {
        var v13 int32 = 0
        var v14 string = ((("" + "") + "hello") + ("hello" + "\n\t"))
        _ = v13
        _ = v14
    } else {
        var v15 uint8 = 42
        switch -2147483648 {
        case -949:
            v15 += 127
        default:
            switch ("hello" + (("hello" + "\n\t") + "foo bar")) {
            case "hello":
                switch ("hello" + ("" + "")) {
                case "":
                    v15 -= v15
                case "foo bar":
                    fmt.Println(v15)
                default:
                    var v16 int16 = 1000
                    _ = v16
                }
            default:
                v15 = v15
            }
        }
        _ = v15
    }
    v17 := []int{}
    var v18 int32 = -1
    v18 += v18
    for fi19 := 0; fi19 < 16; fi19++ {
        v17 = append(v17, (fi19 - fi19))
        _ = fi19
    }
    _ = v17
    _ = v18
    return 9223372036854775807
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = -201
    v3 := make(map[uint32]bool)
    var v4 float32 = -2.25
    var v5 float64 = ((float64(1e10) * (1e10 - 0.1)) + (1.5 / -2.25))
    var v6 bool = (true || true)
    fmt.Println(v4, v1, v5)
    switch ("\n\t" + "hello") {
    case "a":
        for fi20 := 0; fi20 < 9; fi20++ {
            v3[4294967295] = v6
            fmt.Println(v5, v1, fi20)
            _ = fi20
        }
    }
    if v6 {
        var v21 int16 = -1
        var v22 uint8 = 234
        _ = v21
        _ = v22
    } else {
        v23 := []int{553, 32767, 421}
        _ = v23
    }
    _ = len(v3)
    v24 := GenS1{F0: 717}
    var v25 float64 = v5
    fmt.Println(v2, v2)
    if (v2 < 625) {
        for fi26 := 0; fi26 < 19; fi26++ {
            if (((v4 != 1.5) || true) && v6) {
                for fi27 := 0; fi27 < 18; fi27++ {
                    v2 = v2
                    var v28 uint64 = 166
                    _ = fi27
                    _ = v28
                }
                if false {
                    v29 := v3[2]
                    _ = v29
                }
                switch "" {
                case "":
                    delete(v3, 548)
                case "foo bar":
                    for fi30 := 0; fi30 < 17; fi30++ {
                        v2 += v2
                        v31 := GenS1{F0: 2}
                        if v6 {
                            if (0 != 1) {
                                var v32 string = "a"
                                v33 := GenS1{F0: 65535}
                                _ = v32
                                _ = v33
                            } else {
                                fmt.Println(fi30, v2, fi26)
                                switch (fi30 * fi26) {
                                case 2147483647:
                                    if ((("hello" + "foo bar") + "foo bar") > "foo bar") {
                                        switch "a" {
                                        case "\n\t":
                                            v34 := v3[4294967295]
                                            _ = v34
                                        case "foo bar":
                                            if (!(float32(v5) < float32((v25 + -2.25)))) {
                                                v25 = v5
                                            }
                                        case "":
                                            v25 -= 1.5
                                        }
                                        if (0.1 == (v4 * 0.1)) {
                                            var v35 uint16 = 1
                                            _ = v35
                                        }
                                        switch 127 {
                                        case 127:
                                            v36 := GenS2{F0: 18446744073709551615, F1: 4294967295}
                                            _ = v36
                                        case 9223372036854775807:
                                            switch (fi26 | 128) {
                                            case 1:
                                                fmt.Println(fi26, v1)
                                            default:
                                                var v37 int32 = 1
                                                _ = v37
                                            }
                                        case 0:
                                            v6 = v6
                                        default:
                                            if false {
                                                if (0 >= 108) {
                                                    var v38 int64 = v1
                                                    _ = v38
                                                }
                                            }
                                        }
                                    }
                                case -324:
                                    var v39 float64 = (3.14159 * float64(0.0))
                                    _ = v39
                                case -272:
                                    v40 := GenS2{F0: 255, F1: 755}
                                    _ = v40
                                default:
                                    var v41 float64 = (v5 + v5)
                                    _ = v41
                                }
                                v42 := GenS2{F0: 280, F1: 580}
                                _ = v42
                            }
                            var v43 int16 = (v2 | v2)
                            var v44 bool = (0 != 0)
                            _ = v43
                            _ = v44
                        }
                        _ = fi30
                        _ = v31
                    }
                case "\n\t":
                    var v45 int32 = 1
                    _ = v45
                default:
                    v46 := []int{1, 32767, (-(fi26)), fi26}
                    _ = v46
                }
            }
            _ = fi26
        }
        var v47 int = 32768
        for fi48 := 0; fi48 < 18; fi48++ {
            v49 := GenS3{F0: -2147483648, F1: 115}
            v50 := GenS4{F0: (v2 >> ((18446744073709551615) & 63)), F1: v2, F2: 42}
            _ = fi48
            _ = v49
            _ = v50
        }
        _ = v47
    }
    var v51 uint8 = 127
    if v6 {
        switch ("foo bar" + ("\n\t" + "a")) {
        case "hello":
            v52 := GenS4{F0: v2, F1: v2, F2: 42}
            _ = v52
        case "":
            if v6 {
                _ = len(v3)
                v51 = 255
            } else {
                var v53 float32 = float32((1e10 * (v25 - v25)))
                fmt.Println(v25, v4)
                v54 := GenS5{F0: 0, F1: ("" + ("" + ("" + "\n\t")))}
                _ = v53
                _ = v54
            }
        }
        var v55 bool = (!(((v1 / ((42) | 1)) >= (-(v1))) || v6))
        _ = v55
    } else {
        fmt.Println(v5, v2)
        for fi56 := 0; fi56 < 1; fi56++ {
            if v6 {
                var v57 int = int(v51)
                v6 = v6
                for fi58 := 0; fi58 < 8; fi58++ {
                    v59 := GenS6{F0: 856}
                    if (65535 >= 4294967295) {
                        switch int(v2) {
                        case 9223372036854775807:
                            v51 -= v51
                        case 844:
                            v60 := []int{(fi58 << ((0) & 63)), fi56, -9223372036854775808, (fi58 | 0)}
                            _ = v60
                        }
                    } else {
                        if (!v6) {
                            fmt.Println(v4)
                            v61 := GenS5{F0: 834, F1: ""}
                            fmt.Println(v6)
                            _ = v61
                        }
                        switch 9223372036854775807 {
                        case -9223372036854775808:
                            var v62 bool = v6
                            _ = v62
                        default:
                            v4 += (v4 + v4)
                        }
                        switch (fi56 & (-(v57))) {
                        case 42:
                            v4 = float32((((v25 + 0.1) + ((v25 / -2.25) - v25)) * v5))
                        }
                    }
                    v1 += (v1 / (((v1 * v1)) | 1))
                    _ = fi58
                    _ = v59
                }
                _ = v57
            } else {
                var v63 int32 = -794
                for fi64 := 0; fi64 < 17; fi64++ {
                    v65 := v3[4294967295]
                    var v66 uint32 = 1
                    var v67 int64 = v1
                    _ = fi64
                    _ = v65
                    _ = v66
                    _ = v67
                }
                for fi68 := 0; fi68 < 19; fi68++ {
                    if v6 {
                        var v69 float64 = v5
                        _ = v69
                    }
                    var v70 float64 = v25
                    v71 := GenS5{F0: 42, F1: "\n\t"}
                    _ = fi68
                    _ = v70
                    _ = v71
                }
                _ = v63
            }
            _ = fi56
        }
    }
    for fi72 := 0; fi72 < 12; fi72++ {
        v1 -= (v1 ^ v1)
        _ = fi72
    }
    fmt.Println(v6)
    if v6 {
        if true {
            for fi73 := 0; fi73 < 2; fi73++ {
                var v74 uint8 = 38
                _ = fi73
                _ = v74
            }
            if v6 {
                v75 := []int{128, 1, 554, 2147483647}
                _ = v75
            }
            var v76 uint8 = (v51 % ((255) | 1))
            _ = v76
        } else {
            var v77 float32 = v4
            fmt.Println(v25, v77, v77)
            v4 -= v4
            _ = v77
        }
        var v78 uint8 = (v51 << ((uint(v51)) & 63))
        _ = v78
    } else {
        var v79 int32 = int32(v2)
        _ = v79
    }
    if (true || v6) {
        if ((v51 ^ 128) >= (v51 - v51)) {
            v51 -= 127
            switch 128 {
            case 127:
                if false {
                    for fi80 := 0; fi80 < 2; fi80++ {
                        switch fi80 {
                        case -2147483648:
                            for fi81 := 0; fi81 < 9; fi81++ {
                                var v82 string = ("foo bar" + (("hello" + "foo bar") + ((("foo bar" + ("a" + "hello")) + "\n\t") + "\n\t")))
                                var v83 int64 = (v1 | (v1 + v1))
                                _ = fi81
                                _ = v82
                                _ = v83
                            }
                        case -1:
                            if false {
                                for fi84 := 0; fi84 < 20; fi84++ {
                                    var v85 bool = (!false)
                                    _ = fi84
                                    _ = v85
                                }
                                if v6 {
                                    var v86 uint16 = 32767
                                    _ = v86
                                } else {
                                    v87 := GenS1{F0: 386}
                                    var v88 float64 = v5
                                    _ = v87
                                    _ = v88
                                }
                                var v89 int32 = int32(v1)
                                _ = v89
                            } else {
                                switch ((("a" + ("" + "hello")) + "foo bar") + "a") {
                                case "\n\t":
                                    v90 := GenS7{F0: v2, F1: (v51 ^ v51)}
                                    _ = v90
                                }
                                for fi91 := 0; fi91 < 9; fi91++ {
                                    fmt.Println(v5, v1)
                                    v92 := []int{0, 0, (fi91 & 128), -884}
                                    _ = fi91
                                    _ = v92
                                }
                            }
                        case 2147483647:
                            var v93 int32 = 1
                            _ = v93
                        }
                        switch ("foo bar" + "foo bar") {
                        case "\n\t":
                            v25 -= v5
                        case "":
                            if ((fi80 + (fi80 << ((42) & 63))) != 497) {
                                var v94 bool = ((((3.14159 * 0.1) - ((float64((float32(float64(v4)) * 3.14159)) - v5) / ((v5 * 0.1) + (v25 + (3.14159 / v5))))) <= 0.0) || true)
                                v51 += v51
                                _ = v94
                            }
                        case "hello":
                            v95 := GenS8{F0: fi80, F1: genFn0(1, (v2 / ((42) | 1))), F2: (v51 << ((79) & 63))}
                            _ = v95
                        }
                        _ = fi80
                    }
                    var v96 bool = ((v1 << ((2) & 63)) != genFn1())
                    _ = v96
                }
            default:
                var v97 int8 = 10
                _ = v97
            }
        }
        switch 128 {
        case -593:
            for fi98 := 0; fi98 < 17; fi98++ {
                switch ("a" + "hello") {
                case "hello":
                    _ = len(v3)
                }
                v6 = (1 == 2)
                if (false || (620 > 342)) {
                    v1 += (v1 - (v1 | v1))
                }
                _ = fi98
            }
        }
    }
    for fi99 := 0; fi99 < 7; fi99++ {
        switch ("\n\t" + "foo bar") {
        case "a":
            var v100 uint16 = 255
            _ = v100
        case "foo bar":
            v1 -= v1
        case "":
            v101 := []int{(fi99 + 128), fi99, 1, fi99}
            _ = v101
        }
        _ = fi99
    }
    for fi102 := 0; fi102 < 6; fi102++ {
        for fi103 := 0; fi103 < 13; fi103++ {
            v104 := GenS3{F0: -1, F1: 647}
            v105 := GenS4{F0: 32767, F1: (v2 % ((0) | 1)), F2: 379}
            _ = fi103
            _ = v104
            _ = v105
        }
        _ = len(v3)
        delete(v3, 0)
        _ = fi102
    }
    if (((!(541 <= 42)) || (!v6)) && (!v6)) {
        v106 := GenS9{F0: 4294967295}
        fmt.Println(v6, v6)
        _ = v106
    } else {
        if false {
            fmt.Println(v25)
            v107 := GenS4{F0: v2, F1: v2, F2: 385}
            v108 := GenS10{F0: 42, F1: "\n\t"}
            _ = v107
            _ = v108
        } else {
            v51 = (^(v51))
            if v6 {
                switch "a" {
                case "hello":
                    v25 = -2.25
                }
                for fi109 := 0; fi109 < 11; fi109++ {
                    v110 := []int{(-(fi109)), (fi109 & fi109)}
                    _ = fi109
                    _ = v110
                }
                v111 := GenS6{F0: -1}
                _ = v111
            }
            if v6 {
                switch -439 {
                case -9223372036854775808:
                    if (!(v6 || (!v6))) {
                        v112 := GenS1{F0: 1}
                        v6 = false
                        _ = v112
                    }
                }
                switch "a" {
                case "a":
                    var v113 uint = 42
                    _ = v113
                case "hello":
                    v114 := v3[727]
                    _ = v114
                case "":
                    switch ("\n\t" + "hello") {
                    case "a":
                        fmt.Println(v1, v4, v1)
                    case "\n\t":
                        switch 127 {
                        case 42:
                            v3[65535] = ((306 != 131) && (!v6))
                        case -9223372036854775808:
                            fmt.Println(v5)
                        default:
                            if v6 {
                                v115 := []int{0}
                                var v116 uint32 = 299
                                _ = v115
                                _ = v116
                            } else {
                                if (18446744073709551615 <= 4294967295) {
                                    switch ("" + ("a" + (("foo bar" + "foo bar") + ("a" + "a")))) {
                                    case "a":
                                        for fi117 := 0; fi117 < 4; fi117++ {
                                            switch int(v2) {
                                            case 32767:
                                                var v118 string = "hello"
                                                _ = v118
                                            case 128:
                                                fmt.Println(v5, v25)
                                            case 32768:
                                                if v6 {
                                                    if (-270 >= fi117) {
                                                        switch (fi117 & fi117) {
                                                        case 127:
                                                            var v119 int64 = (v1 - v1)
                                                            _ = v119
                                                        case 377:
                                                            v120 := GenS9{F0: 4294967295}
                                                            _ = v120
                                                        case -711:
                                                            if v6 {
                                                                v3[0] = false
                                                            }
                                                        }
                                                        var v121 uint64 = 1
                                                        _ = v121
                                                    }
                                                    fmt.Println(fi117)
                                                    delete(v3, 641)
                                                }
                                            }
                                            if v6 {
                                                for fi122 := 0; fi122 < 7; fi122++ {
                                                    if (!(v25 == v5)) {
                                                        var v123 uint8 = (v51 % (((^(v51))) | 1))
                                                        var v124 int = 605
                                                        _ = v123
                                                        _ = v124
                                                    }
                                                    v6 = true
                                                    if false {
                                                        v2 += 32767
                                                        delete(v3, 42)
                                                        v125 := GenS4{F0: v2, F1: v2, F2: int32(v1)}
                                                        _ = v125
                                                    } else {
                                                        v6 = true
                                                        v126 := GenS10{F0: uint64(v51), F1: ((("\n\t" + "\n\t") + (("foo bar" + ("a" + "\n\t")) + "\n\t")) + ("hello" + (((("a" + "a") + "") + "foo bar") + "foo bar")))}
                                                        if ((v51 > v51) && (32767 == 2)) {
                                                            v1 = v1
                                                            if (!true) {
                                                                if true {
                                                                    switch 128 {
                                                                    case 0:
                                                                        var v127 int = (fi117 >> ((787) & 63))
                                                                        _ = v127
                                                                    }
                                                                    var v128 uint64 = 286
                                                                    v129 := v3[2147483648]
                                                                    var v130 uint16 = 2
                                                                    if v129 {
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            if (v6 || (!v6)) {
                                                                                v6 = v6
                                                                                v131 := []int{}
                                                                                _ = v131
                                                                            }
                                                                        case "a":
                                                                            switch "a" {
                                                                            case "hello":
                                                                                v132 := []int{fi122, (fi122 % ((fi117) | 1)), (fi122 * (fi117 | fi117))}
                                                                                _ = v132
                                                                            case "":
                                                                                var v133 uint64 = (^(v128))
                                                                                _ = v133
                                                                            case "a":
                                                                                var v134 int64 = (-(v1))
                                                                                _ = v134
                                                                            default:
                                                                                v4 -= -2.25
                                                                            }
                                                                        case "":
                                                                            if (v5 > ((v5 + 0.0) / 3.14159)) {
                                                                                var v135 uint16 = (v130 >> ((42) & 63))
                                                                                _ = v135
                                                                            }
                                                                        }
                                                                        v136 := []int{-829, -248}
                                                                        switch (fi117 / ((fi117) | 1)) {
                                                                        case -1:
                                                                            if true {
                                                                                v130 = v130
                                                                                switch (("hello" + "foo bar") + ("a" + "a")) {
                                                                                case "foo bar":
                                                                                    var v137 int32 = 0
                                                                                    _ = v137
                                                                                case "":
                                                                                    fmt.Println(v128, v129)
                                                                                default:
                                                                                    if v6 {
                                                                                        var v138 uint16 = (v130 - (v130 + (v130 & v130)))
                                                                                        var v139 int8 = -20
                                                                                        _ = v138
                                                                                        _ = v139
                                                                                    } else {
                                                                                        v136 = append(v136, fi117)
                                                                                        var v140 int16 = v2
                                                                                        _ = v140
                                                                                    }
                                                                                }
                                                                                v141 := GenS3{F0: 0, F1: 879}
                                                                                _ = v141
                                                                            } else {
                                                                                v142 := GenS11{F0: 44, F1: v1}
                                                                                _ = v142
                                                                            }
                                                                        }
                                                                        _ = v136
                                                                    }
                                                                    _ = v128
                                                                    _ = v129
                                                                    _ = v130
                                                                } else {
                                                                    _ = len(v3)
                                                                    v143 := GenS4{F0: 42, F1: 1, F2: 2147483647}
                                                                    var v144 int16 = v2
                                                                    _ = v143
                                                                    _ = v144
                                                                }
                                                            } else {
                                                                var v145 uint16 = 2
                                                                v3[2147483648] = (v2 == 0)
                                                                if v6 {
                                                                    v146 := GenS2{F0: 18446744073709551615, F1: 1}
                                                                    fmt.Println(v6, v5)
                                                                    _ = v146
                                                                }
                                                                _ = v145
                                                            }
                                                            fmt.Println(v6, v5)
                                                            v147 := GenS12{F0: true, F1: (-(fi122))}
                                                            var v148 bool = v6
                                                            _ = v147
                                                            _ = v148
                                                        }
                                                        _ = v126
                                                    }
                                                    fmt.Println(v5, v51, fi117)
                                                    _ = fi122
                                                }
                                                switch fi117 {
                                                case 42:
                                                    delete(v3, uint32(v51))
                                                case -2147483648:
                                                    var v149 float32 = 3.14159
                                                    _ = v149
                                                case 0:
                                                    for fi150 := 0; fi150 < 7; fi150++ {
                                                        v151 := GenS8{F0: fi150, F1: 42, F2: (^(v51))}
                                                        var v152 int = int(v51)
                                                        if (int32(v2) > 2147483647) {
                                                            v153 := []int{fi150, fi150, fi150, (fi150 & fi150)}
                                                            _ = v153
                                                        }
                                                        _ = fi150
                                                        _ = v151
                                                        _ = v152
                                                    }
                                                }
                                            } else {
                                                v154 := []int{-9223372036854775808}
                                                var v155 string = "foo bar"
                                                var v156 int = (fi117 ^ fi117)
                                                _ = v154
                                                _ = v155
                                                _ = v156
                                            }
                                            _ = fi117
                                        }
                                    }
                                    v157 := v3[0]
                                    fmt.Println(v25)
                                    _ = v157
                                } else {
                                    for fi158 := 0; fi158 < 8; fi158++ {
                                        for fi159 := 0; fi159 < 18; fi159++ {
                                            var v160 int16 = 1000
                                            _ = fi159
                                            _ = v160
                                        }
                                        v2 += (v2 | (v2 - (^(v2))))
                                        _ = fi158
                                    }
                                    if true {
                                        v161 := []int{786, 32767}
                                        if v6 {
                                            if v6 {
                                                var v162 uint = 18446744073709551615
                                                var v163 uint64 = 399
                                                _ = v162
                                                _ = v163
                                            } else {
                                                switch "" {
                                                case "\n\t":
                                                    v51 += v51
                                                default:
                                                    v2 -= v2
                                                }
                                                var v164 bool = (v2 > v2)
                                                _ = v164
                                            }
                                            v165 := GenS6{F0: 42}
                                            _ = v165
                                        } else {
                                            v166 := GenS13{F0: 798}
                                            _ = v166
                                        }
                                        fmt.Println(v51)
                                        _ = v161
                                    }
                                }
                                v167 := GenS9{F0: 709}
                                _ = v167
                            }
                        }
                    case "":
                        var v168 int32 = 555
                        _ = v168
                    }
                default:
                    var v169 uint64 = 870
                    _ = v169
                }
                var v170 int16 = (^(v2))
                _ = v170
            } else {
                if true {
                    v5 += v25
                }
            }
        }
        v171 := GenS3{F0: 875, F1: 592}
        v172 := []int{64}
        _ = v171
        _ = v172
    }
    fmt.Println(v4, v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
