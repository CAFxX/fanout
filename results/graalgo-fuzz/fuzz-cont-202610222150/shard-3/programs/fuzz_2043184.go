package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32; F1 int; F2 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 string; F1 bool }
func genFn0(p0 uint) int32 {
    var v6 uint64 = 533
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(fi7)
        var v8 uint16 = 1
        _ = fi7
        _ = v8
    }
    var v9 uint16 = 1
    _ = p0
    _ = v6
    _ = v9
    return -464
}
func genFn1(p0 bool, p1 bool, p2 uint) int8 {
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := []int{}
        if ((3.14159 <= ((0.0 * 1.5) - 1.5)) || p1) {
            v12 := GenS1{F0: 23, F1: 32767}
            p0 = p1
            v13 := GenS1{F0: genFn0((p2 % ((p2) | 1))), F1: 2}
            _ = v12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            if false {
                var v16 int64 = 42
                switch "hello" {
                case "a":
                    v17 := []int{}
                    _ = v17
                case "foo bar":
                    fmt.Println(p0, p1)
                case "\n\t":
                    if p1 {
                        v16 += -912
                        p1 = false
                    } else {
                        var v18 uint64 = 1
                        _ = v18
                    }
                default:
                    var v19 int = (fi14 / (((-(fi14))) | 1))
                    _ = v19
                }
                p1 = p1
                _ = v16
            } else {
                var v20 int32 = -1
                switch ("foo bar" + "a") {
                case "a":
                    p0 = (!p0)
                case "\n\t":
                    var v21 uint64 = 4294967295
                    _ = v21
                default:
                    switch "\n\t" {
                    case "\n\t":
                        switch "" {
                        case "":
                            p1 = p1
                        }
                    case "a":
                        var v22 uint = (p2 << (((p2 * 18446744073709551615)) & 63))
                        _ = v22
                    default:
                        v20 += v20
                    }
                }
                var v23 uint32 = 1
                _ = v20
                _ = v23
            }
            p2 = p2
            _ = fi15
        }
        if (3.14159 == 0.1) {
            p2 = 42
            var v24 int = fi14
            _ = v24
        }
        v25 := []int{fi14, (fi14 ^ (fi14 & (fi14 / ((fi14) | 1)))), (^(fi14)), -728}
        _ = fi14
        _ = v25
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        for fi27 := 0; fi27 < 2; fi27++ {
            var v28 uint16 = 255
            p0 = (255 != 128)
            _ = fi27
            _ = v28
        }
        p2 -= p2
        _ = fi26
    }
    var v29 string = "\n\t"
    var v30 int8 = -109
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    _ = v30
    return (v30 + v30)
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 42
    v3 := make(map[string]int)
    var v4 float32 = 3.14159
    v5 := []uint32{2147483648, 431, 670}
    fmt.Println(v4, v4)
    v2 = v2
    for fi31 := 0; fi31 < 2; fi31++ {
        if false {
            var v32 uint8 = 127
            _ = v32
        }
        delete(v3, ("hello" + "\n\t"))
        var v33 float32 = v4
        _ = fi31
        _ = v33
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        v2 = v2
        v35 := GenS2{F0: 0}
        fmt.Println(v2)
        _ = fi34
        _ = v35
    }
    switch 9223372036854775807 {
    case -9223372036854775808:
        v1["hello"] = "\n\t"
    case -146:
        switch 874 {
        case 127:
            if false {
                v36 := v3[""]
                var v37 int16 = 1000
                if (65535 > (^(v2))) {
                    switch v36 {
                    case 32767:
                        var v38 int16 = v37
                        _ = v38
                    case 42:
                        var v39 int = 2147483647
                        _ = v39
                    }
                } else {
                    if true {
                        switch (v36 % (((^(v36))) | 1)) {
                        case 922:
                            switch (((("a" + "\n\t") + "foo bar") + "\n\t") + "\n\t") {
                            case "hello":
                                if ((v4 + v4) <= (float32(3.14159) / v4)) {
                                    v40 := GenS2{F0: 1}
                                    fmt.Println(v2, v4, v36)
                                    v41 := GenS2{F0: 1}
                                    _ = v40
                                    _ = v41
                                } else {
                                    v5 = append(v5, 4294967295)
                                }
                            case "a":
                                v4 = v4
                            case "foo bar":
                                for fi42 := 0; fi42 < 8; fi42++ {
                                    _ = len(v1)
                                    v2 -= 4294967295
                                    if (true || false) {
                                        fmt.Println(fi42, v37)
                                        switch "" {
                                        case "foo bar":
                                            v43 := GenS3{F0: v37, F1: v37}
                                            _ = v43
                                        case "a":
                                            var v44 uint8 = 127
                                            _ = v44
                                        case "hello":
                                            switch fi42 {
                                            case 173:
                                                if false {
                                                    v37 = -32768
                                                    v1[("\n\t" + ("" + ("\n\t" + (("" + "hello") + "\n\t"))))] = "hello"
                                                } else {
                                                    v45 := GenS4{F0: 550}
                                                    _ = v45
                                                }
                                            case 1:
                                                _ = len(v5)
                                            case 2147483647:
                                                var v46 int8 = -52
                                                _ = v46
                                            }
                                        }
                                        v4 = v4
                                    }
                                    _ = fi42
                                }
                            }
                        case 783:
                            for fi47 := 0; fi47 < 10; fi47++ {
                                if false {
                                    switch fi47 {
                                    case 32767:
                                        var v48 uint = 637
                                        _ = v48
                                    case 9223372036854775807:
                                        v49 := GenS3{F0: -32768, F1: (v37 + (v37 + 80))}
                                        _ = v49
                                    }
                                } else {
                                    if (-1 != -109) {
                                        v50 := GenS4{F0: 32767}
                                        v36 = 128
                                        _ = v50
                                    } else {
                                        if (42 == 202) {
                                            var v51 uint8 = 255
                                            _ = v51
                                        }
                                        var v52 int64 = 9223372036854775807
                                        _ = v52
                                    }
                                    if true {
                                        if (!(195 > 128)) {
                                            switch ("\n\t" + "hello") {
                                            case "hello":
                                                v4 = (v4 - (v4 + v4))
                                            case "":
                                                v36 -= v36
                                            }
                                            if true {
                                                v2 -= v2
                                                if true {
                                                    var v53 int32 = 1
                                                    switch fi47 {
                                                    case -9223372036854775808:
                                                        v37 = 1000
                                                    case 128:
                                                        for fi54 := 0; fi54 < 17; fi54++ {
                                                            var v55 bool = (((0 == 2) && (!(!false))) && false)
                                                            _ = fi54
                                                            _ = v55
                                                        }
                                                    case 4294967295:
                                                        fmt.Println(v36, v2, v2)
                                                    default:
                                                        var v56 float64 = float64(0.0)
                                                        _ = v56
                                                    }
                                                    for fi57 := 0; fi57 < 19; fi57++ {
                                                        var v58 string = ("" + ("hello" + ""))
                                                        _ = fi57
                                                        _ = v58
                                                    }
                                                    _ = v53
                                                } else {
                                                    fmt.Println(v37)
                                                }
                                            }
                                        } else {
                                            v59 := GenS2{F0: -9223372036854775808}
                                            var v60 uint16 = 2
                                            fmt.Println(v4, v4, v60)
                                            _ = v59
                                            _ = v60
                                        }
                                        v4 = 0.1
                                        var v61 int32 = genFn0(65535)
                                        _ = v61
                                    } else {
                                        if false {
                                            v62 := GenS5{F0: 0, F1: -9223372036854775808, F2: 798}
                                            v3["foo bar"] = (-(v36))
                                            var v63 uint16 = 255
                                            _ = v62
                                            _ = v63
                                        } else {
                                            v37 = int16(v36)
                                            v64 := v1[("foo bar" + (("foo bar" + "") + "a"))]
                                            var v65 int = -1
                                            _ = v64
                                            _ = v65
                                        }
                                    }
                                }
                                v4 += v4
                                _ = fi47
                            }
                        case 42:
                            _ = len(v3)
                        default:
                            v2 = (v2 >> ((v2) & 63))
                        }
                        v66 := v1["hello"]
                        _ = v66
                    } else {
                        if (!true) {
                            _ = len(v3)
                            if (false && (!(42 < 127))) {
                                var v67 string = ""
                                _ = v67
                            }
                        }
                        switch "hello" {
                        case "foo bar":
                            for fi68 := 0; fi68 < 15; fi68++ {
                                if false {
                                    if (!true) {
                                        v37 -= 277
                                        fmt.Println(v2)
                                    }
                                    fmt.Println(fi68, v36)
                                    _ = len(v1)
                                }
                                var v69 string = "hello"
                                _ = fi68
                                _ = v69
                            }
                        case "a":
                            var v70 uint8 = 1
                            _ = v70
                        default:
                            var v71 int16 = (v37 * 0)
                            _ = v71
                        }
                    }
                }
                _ = v36
                _ = v37
            }
        default:
            if (!false) {
                for fi72 := 0; fi72 < 20; fi72++ {
                    v73 := GenS3{F0: 32767, F1: 1000}
                    fmt.Println(v4, v4)
                    if (!(((3.14159 / v4) / v4) < v4)) {
                        fmt.Println(v2, v4)
                        var v74 int16 = 42
                        switch fi72 {
                        case 0:
                            for fi75 := 0; fi75 < 2; fi75++ {
                                var v76 uint32 = 524
                                v76 += v76
                                v77 := GenS4{F0: uint16(v76)}
                                var v78 int16 = v74
                                _ = fi75
                                _ = v76
                                _ = v77
                                _ = v78
                            }
                        default:
                            switch 204 {
                            case 280:
                                v4 = (v4 * float32(1.5))
                            case 0:
                                v79 := GenS1{F0: -859, F1: 2}
                                _ = v79
                            default:
                                for fi80 := 0; fi80 < 6; fi80++ {
                                    if true {
                                        switch fi72 {
                                        case 4294967295:
                                            v74 = v74
                                        default:
                                            var v81 int32 = 2147483647
                                            _ = v81
                                        }
                                        v2 = 658
                                    } else {
                                        v4 = (v4 - float32((0.1 / 1e10)))
                                        v74 -= (v74 | (v74 % ((v74) | 1)))
                                    }
                                    var v82 float64 = float64(3.14159)
                                    var v83 int16 = -453
                                    _ = fi80
                                    _ = v82
                                    _ = v83
                                }
                            }
                        }
                        _ = v74
                    }
                    _ = fi72
                    _ = v73
                }
            }
        }
    case 32768:
        fmt.Println(v2, v4)
    }
    switch "hello" {
    case "hello":
        switch "foo bar" {
        case "hello":
            v2 = v2
        case "":
            if (103 >= 1) {
                fmt.Println(v2, v4)
            }
        default:
            v84 := GenS2{F0: -9223372036854775808}
            _ = v84
        }
    case "a":
        if (true && true) {
            if true {
                var v85 float64 = float64((1e10 * v4))
                _ = v85
            }
            if true {
                if (0 < 2) {
                    if (-374 < 1) {
                        delete(v3, "hello")
                        if ((v2 << ((v2) & 63)) == (^(v2))) {
                            var v86 float32 = float32(float64(1e10))
                            v4 = v4
                            _ = v86
                        } else {
                            var v87 int = 650
                            v87 = v87
                            if (-1 <= 32767) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                for fi88 := 0; fi88 < 18; fi88++ {
                                    v89 := GenS6{F0: false}
                                    if true {
                                        for fi90 := 0; fi90 < 17; fi90++ {
                                            fmt.Println(v87)
                                            var v91 uint = (v2 + v2)
                                            _ = fi90
                                            _ = v91
                                        }
                                        v92 := GenS6{F0: true}
                                        for fi93 := 0; fi93 < 20; fi93++ {
                                            var v94 uint8 = 181
                                            v3["foo bar"] = 127
                                            fmt.Println(fi93)
                                            _ = fi93
                                            _ = v94
                                        }
                                        _ = v92
                                    } else {
                                        v87 += 127
                                        v95 := GenS4{F0: 255}
                                        _ = v95
                                    }
                                    _ = fi88
                                    _ = v89
                                }
                                v2 = (v2 + (^(v2)))
                            }
                            _ = v87
                        }
                    } else {
                        var v96 int16 = 32767
                        switch "a" {
                        case "":
                            if true {
                                var v97 uint64 = 526
                                switch "foo bar" {
                                case "":
                                    v4 += v4
                                }
                                if len(v5) > 0 {
                                    v5[0] = 65535
                                }
                                _ = v97
                            } else {
                                switch "" {
                                case "foo bar":
                                    v96 += (v96 >> (((v2 >> ((4294967295) & 63))) & 63))
                                default:
                                    var v98 int32 = 175
                                    _ = v98
                                }
                                var v99 string = ("\n\t" + "hello")
                                if false {
                                    var v100 float32 = float32(0.1)
                                    v1[v99] = ("a" + v99)
                                    _ = v100
                                }
                                _ = v99
                            }
                        default:
                            for fi101 := 0; fi101 < 17; fi101++ {
                                if len(v5) > 0 {
                                    v5[0] = 1
                                }
                                v96 += (v96 - v96)
                                _ = fi101
                            }
                        }
                        switch -900 {
                        case 907:
                            if (((174 <= 127) && false) && (42 >= 382)) {
                                if (65535 < 2) {
                                    for fi102 := 0; fi102 < 7; fi102++ {
                                        v5 = append(v5, 172)
                                        var v103 int = fi102
                                        v4 -= v4
                                        _ = fi102
                                        _ = v103
                                    }
                                    switch (("a" + "foo bar") + ("\n\t" + "a")) {
                                    case "foo bar":
                                        for fi104 := 0; fi104 < 8; fi104++ {
                                            if (false || (-128 > 127)) {
                                                v105 := GenS7{F0: genFn0((v2 << ((725) & 63)))}
                                                _ = v105
                                            }
                                            v96 += v96
                                            _ = fi104
                                        }
                                    case "\n\t":
                                        if len(v5) > 0 {
                                            v5[0] = 0
                                        }
                                    default:
                                        fmt.Println(v4, v2)
                                    }
                                    v106 := GenS5{F0: -2147483648, F1: -990, F2: 303}
                                    _ = v106
                                }
                            } else {
                                var v107 string = ((("foo bar" + "\n\t") + ("foo bar" + "foo bar")) + "\n\t")
                                _ = v107
                            }
                        case -503:
                            v5 = append(v5, 74)
                        case 306:
                            var v108 uint8 = 220
                            _ = v108
                        default:
                            if (false && ("\n\t" <= "hello")) {
                                if (!false) {
                                    v4 = float32(3.14159)
                                    if false {
                                        switch ("\n\t" + "hello") {
                                        case "\n\t":
                                            v4 = v4
                                        case "hello":
                                            if true {
                                                switch "\n\t" {
                                                case "a":
                                                    switch 244 {
                                                    case -9223372036854775808:
                                                        v2 -= 287
                                                    case -734:
                                                        _ = len(v5)
                                                    case 127:
                                                        if len(v5) > 0 {
                                                            v5[0] = 129
                                                        }
                                                    }
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                        default:
                                            fmt.Println(v4)
                                        }
                                        if (0 < 78) {
                                            v109 := v3["\n\t"]
                                            for fi110 := 0; fi110 < 5; fi110++ {
                                                switch v109 {
                                                case 2147483647:
                                                    v111 := v3["\n\t"]
                                                    _ = v111
                                                default:
                                                    var v112 int32 = 1
                                                    _ = v112
                                                }
                                                switch v109 {
                                                case 42:
                                                    v109 += v109
                                                case 2147483647:
                                                    for fi113 := 0; fi113 < 9; fi113++ {
                                                        var v114 string = ("\n\t" + "\n\t")
                                                        _ = fi113
                                                        _ = v114
                                                    }
                                                case -466:
                                                    var v115 bool = (true || true)
                                                    _ = v115
                                                }
                                                var v116 float32 = 3.14159
                                                _ = fi110
                                                _ = v116
                                            }
                                            _ = v109
                                        }
                                        var v117 int8 = 1
                                        _ = v117
                                    } else {
                                        var v118 string = "a"
                                        for fi119 := 0; fi119 < 2; fi119++ {
                                            v1[v118] = v118
                                            _ = fi119
                                        }
                                        _ = v118
                                    }
                                    fmt.Println(v96)
                                }
                            } else {
                                for fi120 := 0; fi120 < 18; fi120++ {
                                    v3["\n\t"] = fi120
                                    v2 = (v2 >> ((v2) & 63))
                                    _ = len(v5)
                                    _ = fi120
                                }
                                for fi121 := 0; fi121 < 11; fi121++ {
                                    for fi122 := 0; fi122 < 1; fi122++ {
                                        v96 -= (v96 - int16(fi122))
                                        _ = fi122
                                    }
                                    _ = fi121
                                }
                                v1["a"] = ""
                            }
                        }
                        _ = v96
                    }
                    var v123 uint16 = 843
                    var v124 int = 2147483647
                    _ = v123
                    _ = v124
                } else {
                    v2 += v2
                }
                for fi125 := 0; fi125 < 19; fi125++ {
                    fmt.Println(fi125, v4, v4)
                    if (!(!(1 >= 2))) {
                        v126 := GenS5{F0: 882, F1: fi125, F2: 32767}
                        v127 := GenS1{F0: 1, F1: 0}
                        v4 = v4
                        _ = v126
                        _ = v127
                    } else {
                        switch "" {
                        case "":
                            v4 -= float32((3.14159 * 1.5))
                        case "a":
                            fmt.Println(v4, v2, v4)
                        case "foo bar":
                            var v128 int8 = 96
                            _ = v128
                        }
                        fmt.Println(v2)
                        var v129 int32 = 0
                        _ = v129
                    }
                    _ = fi125
                }
            }
            var v130 uint64 = 42
            _ = v130
        } else {
            var v131 uint8 = 217
            var v132 bool = true
            _ = v131
            _ = v132
        }
    case "\n\t":
        if true {
            fmt.Println(v4, v4)
        } else {
            fmt.Println(v2, v2)
            var v133 uint = v2
            _ = v133
        }
    }
    fmt.Println(v2, v4)
    if (false || (65535 != 0)) {
        v2 = v2
        if ((false || false) && false) {
            switch -9223372036854775808 {
            case -473:
                v134 := GenS2{F0: 520}
                _ = v134
            case 128:
                fmt.Println(v4, v2)
            case 4294967295:
                if (true && false) {
                    switch 4294967295 {
                    case 32767:
                        if true {
                            v2 -= v2
                        } else {
                            v2 = (^(v2))
                            v4 += (0.1 - ((0.0 * 1.5) * 0.1))
                        }
                    default:
                        v135 := v1["foo bar"]
                        _ = v135
                    }
                    for fi136 := 0; fi136 < 4; fi136++ {
                        v137 := GenS2{F0: 42}
                        _ = fi136
                        _ = v137
                    }
                }
            default:
                v138 := v1[""]
                _ = v138
            }
        }
    } else {
        v139 := GenS2{F0: -9223372036854775808}
        _ = v139
    }
    if ((-323 <= 42) && true) {
        var v140 int = -56
        v140 = v140
        _ = v140
    } else {
        switch ("" + "a") {
        case "":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "\n\t":
            for fi141 := 0; fi141 < 2; fi141++ {
                if (!true) {
                    v142 := GenS5{F0: 2147483647, F1: 127, F2: 34}
                    _ = v142
                } else {
                    v143 := GenS8{F0: -1, F1: 0}
                    v2 = (^(v2))
                    for fi144 := 0; fi144 < 6; fi144++ {
                        v145 := v3["hello"]
                        _ = fi144
                        _ = v145
                    }
                    _ = v143
                }
                switch fi141 {
                case 32767:
                    v4 -= (float32(-2.25) * v4)
                case 4294967295:
                    var v146 uint16 = 2
                    _ = v146
                case -503:
                    var v147 float32 = v4
                    _ = v147
                }
                _ = fi141
            }
        }
    }
    if false {
        v4 = (float32(-2.25) * (3.14159 - v4))
    }
    v5 = append(v5, 1)
    for fi148 := 0; fi148 < 18; fi148++ {
        v5 = append(v5, 210)
        switch (fi148 % ((fi148) | 1)) {
        case 32767:
            var v149 int8 = 66
            _ = v149
        case -14:
            v150 := GenS9{F0: "foo bar", F1: false}
            _ = v150
        default:
            var v151 int32 = -519
            _ = v151
        }
        _ = fi148
    }
    v4 = 1e10
    var v152 int = -9223372036854775808
    v2 += v2
    v5 = append(v5, 4294967295)
    var v153 int32 = 2147483647
    fmt.Println(v153, v4, v153)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v152:", v152)
    fmt.Println("v153:", v153)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v152
    _ = v153
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
