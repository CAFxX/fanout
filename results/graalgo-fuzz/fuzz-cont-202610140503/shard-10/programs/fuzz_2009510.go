package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int8; F2 string }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int16; F1 string; F2 int }
type GenS5 struct { F0 int32; F1 int8; F2 int8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 uint32 }
type GenS8 struct { F0 uint8; F1 uint }
type GenS9 struct { F0 int8; F1 int64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 string; F1 int64; F2 uint64 }
type GenS12 struct { F0 string; F1 int16 }
type GenS13 struct { F0 uint64; F1 int8; F2 int64 }
type GenS14 struct { F0 uint64; F1 uint8 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 uint32; F1 int16; F2 int16 }
type GenS17 struct { F0 int }
func genFn0(p0 int64, p1 uint, p2 int) int32 {
    var v6 int32 = int32(p0)
    for fi7 := 0; fi7 < 8; fi7++ {
        switch (p2 * fi7) {
        case 42:
            var v8 uint16 = 255
            _ = v8
        case 569:
            fmt.Println(p2, fi7)
        case 127:
            v9 := GenS1{F0: fi7}
            _ = v9
        default:
            p0 = (p0 - p0)
        }
        if (false && false) {
            v10 := []int{}
            _ = v10
        } else {
            if false {
                var v11 uint = p1
                fmt.Println(fi7)
                for fi12 := 0; fi12 < 17; fi12++ {
                    if (18446744073709551615 <= 625) {
                        if (false || ((32767 != 535) && (2 != 73))) {
                            v6 += (v6 >> ((395) & 63))
                            if true {
                                v13 := []int{p2}
                                _ = v13
                            }
                            fmt.Println(p0, p1)
                        } else {
                            p1 = (p1 / ((v11) | 1))
                        }
                        var v14 uint = v11
                        v15 := GenS2{F0: 0, F1: 127, F2: "a"}
                        _ = v14
                        _ = v15
                    } else {
                        var v16 uint16 = 544
                        _ = v16
                    }
                    p1 = v11
                    _ = fi12
                }
                _ = v11
            }
            v6 = (^(v6))
            if ((!(!((-9223372036854775808 <= 32768) || true))) || (42 <= 199)) {
                p0 -= 0
                for fi17 := 0; fi17 < 7; fi17++ {
                    v18 := GenS1{F0: 32768}
                    p1 += p1
                    v6 += v6
                    _ = fi17
                    _ = v18
                }
            } else {
                switch (p2 & p2) {
                case 128:
                    var v19 float32 = 1e10
                    _ = v19
                case -918:
                    v20 := GenS1{F0: 0}
                    _ = v20
                case 162:
                    fmt.Println(v6, p1)
                }
                if true {
                    v21 := GenS3{F0: 108, F1: true}
                    _ = v21
                }
            }
        }
        _ = fi7
    }
    v22 := []int{p2}
    if (-1 < -32768) {
        if (0 == 38) {
            p1 += p1
            v23 := GenS1{F0: p2}
            var v24 uint64 = 0
            _ = v23
            _ = v24
        } else {
            for fi25 := 0; fi25 < 4; fi25++ {
                p1 = (p1 - 0)
                _ = fi25
            }
        }
    } else {
        switch "a" {
        case "foo bar":
            p1 = (p1 / ((2) | 1))
        case "a":
            v26 := GenS1{F0: p2}
            _ = v26
        default:
            switch p2 {
            case 0:
                for fi27 := 0; fi27 < 7; fi27++ {
                    if (-2.25 < 1.5) {
                        var v28 uint32 = 2147483648
                        _ = v28
                    } else {
                        if (true || (941 == -461)) {
                            switch "hello" {
                            case "a":
                                p2 = (fi27 & (fi27 | fi27))
                            case "\n\t":
                                var v29 float32 = 0.0
                                _ = v29
                            case "hello":
                                switch ("foo bar" + (("foo bar" + "") + (("hello" + "foo bar") + "hello"))) {
                                case "foo bar":
                                    _ = len(v22)
                                case "a":
                                    for fi30 := 0; fi30 < 15; fi30++ {
                                        var v31 uint32 = 42
                                        if true {
                                            v6 = v6
                                            v32 := GenS4{F0: 32767, F1: "hello", F2: p2}
                                            _ = v32
                                        }
                                        _ = fi30
                                        _ = v31
                                    }
                                case "hello":
                                    v33 := GenS1{F0: (fi27 | (fi27 + (-(fi27))))}
                                    _ = v33
                                default:
                                    if false {
                                        for fi34 := 0; fi34 < 13; fi34++ {
                                            p2 -= (p2 & (p2 % ((32768) | 1)))
                                            if (905 != 32767) {
                                                v6 += (v6 * v6)
                                                if len(v22) > 0 {
                                                    v22[0] = -507
                                                }
                                                p1 += 0
                                                if true {
                                                    var v35 int = p2
                                                    _ = v35
                                                }
                                            } else {
                                                if (128 == 232) {
                                                    p2 += p2
                                                } else {
                                                    var v36 uint = 2
                                                    v37 := GenS2{F0: 768, F1: 1, F2: "hello"}
                                                    fmt.Println(p0)
                                                    _ = v36
                                                    _ = v37
                                                }
                                                var v38 float32 = 0.0
                                                _ = v38
                                            }
                                            _ = fi34
                                        }
                                        v6 -= 1
                                    }
                                }
                            default:
                                for fi39 := 0; fi39 < 7; fi39++ {
                                    var v40 int64 = p0
                                    var v41 uint64 = 1
                                    _ = fi39
                                    _ = v40
                                    _ = v41
                                }
                            }
                        } else {
                            for fi42 := 0; fi42 < 5; fi42++ {
                                p0 -= (p0 / ((p0) | 1))
                                if false {
                                    p1 = p1
                                    v6 = v6
                                } else {
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(fi42, p1, fi27)
                                    default:
                                        fmt.Println(v6, p2)
                                    }
                                    var v43 float64 = ((((-2.25 * 3.14159) - 1e10) + 1.5) / (float64(1.5) * 3.14159))
                                    _ = v43
                                }
                                _ = fi42
                            }
                            if (1 >= 42) {
                                switch ("a" + ("\n\t" + "\n\t")) {
                                case "\n\t":
                                    if true {
                                        if true {
                                            if (uint32(p1) <= 2) {
                                                switch "" {
                                                case "":
                                                    if (545 <= 42) {
                                                        v6 -= v6
                                                    } else {
                                                        var v44 int = p2
                                                        p1 = p1
                                                        p1 += p1
                                                        _ = v44
                                                    }
                                                case "hello":
                                                    _ = len(v22)
                                                default:
                                                    switch "foo bar" {
                                                    case "":
                                                        if false {
                                                            v45 := GenS5{F0: 42, F1: 127, F2: 0}
                                                            var v46 int64 = (p0 + p0)
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                    case "\n\t":
                                                        v47 := GenS3{F0: (p2 | (^(fi27))), F1: (true || true)}
                                                        _ = v47
                                                    default:
                                                        v6 = -1
                                                    }
                                                }
                                                for fi48 := 0; fi48 < 1; fi48++ {
                                                    v49 := GenS6{F0: p1}
                                                    _ = fi48
                                                    _ = v49
                                                }
                                                if len(v22) > 0 {
                                                    v22[0] = fi27
                                                }
                                            }
                                        } else {
                                            p1 -= (^(p1))
                                            _ = len(v22)
                                            p1 = (p1 + p1)
                                        }
                                        p0 -= p0
                                        var v50 uint16 = 255
                                        _ = v50
                                    } else {
                                        switch "foo bar" {
                                        case "a":
                                            var v51 int8 = 127
                                            _ = v51
                                        default:
                                            for fi52 := 0; fi52 < 3; fi52++ {
                                                if (true || false) {
                                                    if len(v22) > 0 {
                                                        v22[0] = fi27
                                                    }
                                                    switch p2 {
                                                    case 1:
                                                        _ = len(v22)
                                                    case 32768:
                                                        if (!false) {
                                                            p0 += 42
                                                            _ = len(v22)
                                                            p0 = p0
                                                        }
                                                    }
                                                } else {
                                                    var v53 float64 = (float64(0.1) * 3.14159)
                                                    switch fi52 {
                                                    case 2147483647:
                                                        v54 := GenS7{F0: 32767, F1: 820}
                                                        _ = v54
                                                    }
                                                    p0 += 1
                                                    _ = v53
                                                }
                                                _ = fi52
                                            }
                                        }
                                        fmt.Println(fi27)
                                    }
                                case "":
                                    v55 := GenS4{F0: 32767, F1: "hello", F2: (fi27 / (((^(fi27))) | 1))}
                                    _ = v55
                                default:
                                    p0 = 1
                                }
                                var v56 uint = 437
                                if true {
                                    v6 = v6
                                    fmt.Println(fi27, p1, p0)
                                } else {
                                    switch -2147483648 {
                                    case -2147483648:
                                        switch ("a" + "foo bar") {
                                        case "\n\t":
                                            switch ("" + "a") {
                                            case "foo bar":
                                                p2 -= 455
                                            case "hello":
                                                fmt.Println(fi27)
                                            }
                                        }
                                    default:
                                        var v57 string = "\n\t"
                                        _ = v57
                                    }
                                }
                                _ = v56
                            } else {
                                v58 := GenS6{F0: 501}
                                for fi59 := 0; fi59 < 16; fi59++ {
                                    var v60 float32 = (((((-2.25 / 3.14159) * (-2.25 - 1e10)) * 3.14159) / -2.25) + 1e10)
                                    var v61 uint16 = 42
                                    v62 := GenS8{F0: 96, F1: p1}
                                    _ = fi59
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                }
                                v6 += 746
                                _ = v58
                            }
                        }
                        v63 := GenS2{F0: 42, F1: -123, F2: "a"}
                        if true {
                            var v64 string = ("foo bar" + "hello")
                            if (true && true) {
                                for fi65 := 0; fi65 < 18; fi65++ {
                                    var v66 int16 = 741
                                    _ = fi65
                                    _ = v66
                                }
                            }
                            _ = v64
                        }
                        _ = v63
                    }
                    fmt.Println(p2)
                    _ = fi27
                }
            case 127:
                switch (("a" + "") + "") {
                case "foo bar":
                    v6 += v6
                default:
                    var v67 float32 = ((0.1 / -2.25) * 0.1)
                    _ = v67
                }
            case 4294967295:
                var v68 float64 = float64(0.1)
                _ = v68
            default:
                p1 += (p1 + 32)
            }
        }
        if false {
            var v69 int32 = 795
            _ = v69
        }
        fmt.Println(p2, p0)
    }
    for fi70 := 0; fi70 < 9; fi70++ {
        var v71 uint16 = 446
        v72 := GenS9{F0: 119, F1: p0}
        v71 = v71
        _ = fi70
        _ = v71
        _ = v72
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v6
    return v6
}
func genFn1() uint8 {
    v73 := GenS10{F0: -97}
    var v74 int32 = 42
    if false {
        if (20 >= 41) {
            fmt.Println(v74)
            v75 := []int{32767, 0, 42}
            _ = v75
        } else {
            if (0 <= -407) {
                if true {
                    for fi76 := 0; fi76 < 15; fi76++ {
                        var v77 uint8 = 255
                        v77 += (^(v77))
                        if (!false) {
                            if ((0.1 + 0.0) < -2.25) {
                                fmt.Println(v74)
                                v78 := GenS1{F0: fi76}
                                _ = v78
                            } else {
                                if false {
                                    for fi79 := 0; fi79 < 10; fi79++ {
                                        switch (fi76 * fi76) {
                                        case 525:
                                            v74 = (v74 ^ (v74 * 42))
                                        case 128:
                                            v80 := GenS1{F0: fi76}
                                            _ = v80
                                        case 244:
                                            if ((!((0.1 - 3.14159) < 1e10)) || true) {
                                                v74 += (v74 << ((866) & 63))
                                                var v81 uint8 = v77
                                                v77 -= 128
                                                _ = v81
                                            }
                                        default:
                                            v74 += (v74 & v74)
                                        }
                                        switch (fi76 / ((fi79) | 1)) {
                                        case 32768:
                                            switch ("hello" + "hello") {
                                            case "hello":
                                                fmt.Println(v74, fi76)
                                            case "foo bar":
                                                if true {
                                                    if (0 >= 18446744073709551615) {
                                                        if false {
                                                            v82 := []int{95, fi79, fi79}
                                                            if len(v82) > 0 {
                                                                v82[0] = fi79
                                                            }
                                                            _ = v82
                                                        }
                                                    } else {
                                                        v77 = (v77 ^ (v77 | 255))
                                                        v74 = (v74 + v74)
                                                        var v83 uint = 18446744073709551615
                                                        _ = v83
                                                    }
                                                    var v84 int8 = -128
                                                    v85 := []int{fi79}
                                                    _ = v84
                                                    _ = v85
                                                } else {
                                                    switch fi79 {
                                                    case 2147483647:
                                                        v86 := GenS10{F0: 0}
                                                        v77 = v77
                                                        v77 = v77
                                                        var v87 float32 = (1e10 / 3.14159)
                                                        _ = v86
                                                        _ = v87
                                                    default:
                                                        v74 = v74
                                                        v77 -= v77
                                                    }
                                                    var v88 int64 = int64(fi76)
                                                    switch (-(fi79)) {
                                                    case 179:
                                                        v88 -= (v88 << ((42) & 63))
                                                    }
                                                    v89 := GenS11{F0: "", F1: -9223372036854775808, F2: 1}
                                                    var v90 uint = 1
                                                    v91 := GenS1{F0: (fi79 & fi79)}
                                                    if true {
                                                        v92 := []int{}
                                                        _ = v92
                                                    }
                                                    var v93 int8 = -128
                                                    switch (fi76 ^ (fi76 & (-(fi79)))) {
                                                    case -92:
                                                        if (false && false) {
                                                            switch "hello" {
                                                            case "hello":
                                                                var v94 uint32 = 281
                                                                _ = v94
                                                            case "\n\t":
                                                                fmt.Println(v90)
                                                            case "a":
                                                                fmt.Println(v90, v93, v74)
                                                            }
                                                            v95 := []int{}
                                                            if (v77 == v77) {
                                                                var v96 bool = ((v90 % ((v90) | 1)) == 1)
                                                                if true {
                                                                    var v97 int8 = v93
                                                                    _ = v97
                                                                }
                                                                fmt.Println(v96, fi79)
                                                                if len(v95) > 0 {
                                                                    v95[0] = fi76
                                                                }
                                                                var v98 bool = (v96 && true)
                                                                var v99 int = fi76
                                                                _ = v96
                                                                _ = v98
                                                                _ = v99
                                                            } else {
                                                                var v100 string = "a"
                                                                if len(v95) >= 2 {
                                                                    v95 = v95[1:len(v95)]
                                                                }
                                                                _ = v100
                                                            }
                                                            v101 := GenS1{F0: fi76}
                                                            v88 = v88
                                                            _ = v101
                                                            _ = v95
                                                        } else {
                                                            var v102 uint64 = 18446744073709551615
                                                            _ = v102
                                                        }
                                                    case -304:
                                                        if (v93 >= -1) {
                                                            v103 := GenS3{F0: fi79, F1: false}
                                                            v77 -= (v77 + 127)
                                                            v77 = v77
                                                            _ = v103
                                                        }
                                                    }
                                                    var v104 float64 = 1e10
                                                    v105 := []int{449}
                                                    _ = v104
                                                    _ = v105
                                                    _ = v88
                                                    _ = v89
                                                    _ = v90
                                                    _ = v91
                                                    _ = v93
                                                }
                                            }
                                        default:
                                            v106 := []int{fi79, fi79}
                                            _ = v106
                                        }
                                        v77 += v77
                                        _ = fi79
                                    }
                                } else {
                                    v77 = (v77 + 199)
                                    for fi107 := 0; fi107 < 19; fi107++ {
                                        v108 := GenS9{F0: 1, F1: 42}
                                        v77 -= v77
                                        _ = fi107
                                        _ = v108
                                    }
                                }
                                var v109 int64 = 1
                                var v110 int16 = -1
                                _ = v109
                                _ = v110
                            }
                            for fi111 := 0; fi111 < 15; fi111++ {
                                var v112 string = ""
                                if (!(true || true)) {
                                    var v113 int16 = 1
                                    _ = v113
                                }
                                _ = fi111
                                _ = v112
                            }
                        } else {
                            var v114 float32 = 0.1
                            switch fi76 {
                            case -1:
                                if false {
                                    fmt.Println(v114, v74)
                                }
                            case 2147483647:
                                fmt.Println(v77, v114)
                            default:
                                v115 := GenS8{F0: (v77 | (v77 >> ((42) & 63))), F1: 18446744073709551615}
                                _ = v115
                            }
                            for fi116 := 0; fi116 < 12; fi116++ {
                                var v117 float64 = (1e10 * ((1e10 - -2.25) / (float64(0.1) * -2.25)))
                                v77 -= 255
                                fmt.Println(v77, v114)
                                _ = fi116
                                _ = v117
                            }
                            _ = v114
                        }
                        _ = fi76
                        _ = v77
                    }
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        var v118 int8 = 15
                        _ = v118
                    default:
                        switch "a" {
                        case "hello":
                            var v119 string = "foo bar"
                            _ = v119
                        }
                    }
                }
                var v120 int64 = 705
                _ = v120
            }
            var v121 bool = false
            if v121 {
                v74 -= (v74 & v74)
                var v122 int16 = 1
                for fi123 := 0; fi123 < 20; fi123++ {
                    v124 := GenS12{F0: "hello", F1: 32767}
                    fmt.Println(v121)
                    if v121 {
                        v125 := GenS5{F0: -1, F1: -128, F2: 42}
                        _ = v125
                    } else {
                        var v126 int8 = -103
                        var v127 bool = (2147483647 < 2147483648)
                        var v128 uint = 841
                        _ = v126
                        _ = v127
                        _ = v128
                    }
                    _ = fi123
                    _ = v124
                }
                _ = v122
            } else {
                switch ((("hello" + ((("" + "foo bar") + "a") + "")) + ("a" + "")) + ((((("\n\t" + ("" + ("hello" + "hello"))) + "a") + "foo bar") + ("" + "a")) + "hello")) {
                case "hello":
                    v121 = v121
                default:
                    var v129 int = 9223372036854775807
                    _ = v129
                }
            }
            _ = v121
        }
    }
    _ = v73
    _ = v74
    return 0
}

func main() {
    var v1 float32 = 0.0
    var v2 int64 = -14
    var v3 uint32 = 2147483648
    var v4 uint64 = 18446744073709551615
    v5 := make(map[string]bool)
    v5[("a" + "hello")] = (1000 >= 42)
    v130 := []int{32768, 4294967295, 128, -2147483648}
    if (((^(v4)) >= (v4 >> ((612) & 63))) || true) {
        v131 := GenS13{F0: v4, F1: 127, F2: v2}
        _ = v131
    }
    v2 = v2
    switch 9223372036854775807 {
    case 795:
        v3 = (^(v3))
    case 638:
        var v132 string = ""
        _ = v132
    default:
        for fi133 := 0; fi133 < 16; fi133++ {
            var v134 uint32 = (^(v3))
            if false {
                switch ("foo bar" + "a") {
                case "hello":
                    v4 += (v4 & uint64(v3))
                default:
                    var v135 int64 = (-(v2))
                    _ = v135
                }
                _ = len(v5)
                _ = len(v130)
            }
            _ = fi133
            _ = v134
        }
    }
    for fi136 := 0; fi136 < 12; fi136++ {
        delete(v5, "\n\t")
        _ = fi136
    }
    switch -1 {
    case 4294967295:
        v1 = v1
    }
    v2 = -1
    v137 := v5[(("hello" + "a") + ("\n\t" + (("\n\t" + "") + ("hello" + "a"))))]
    v137 = (0.1 == (3.14159 - -2.25))
    v137 = (854 == 65535)
    switch 793 {
    case 127:
        for fi138 := 0; fi138 < 9; fi138++ {
            v139 := GenS14{F0: v4, F1: 2}
            switch "\n\t" {
            case "a":
                for fi140 := 0; fi140 < 3; fi140++ {
                    fmt.Println(fi138)
                    var v141 int64 = int64(v3)
                    if v137 {
                        if len(v130) >= 2 {
                            v130 = v130[1:len(v130)]
                        }
                        switch int(v2) {
                        case -9223372036854775808:
                            _ = len(v130)
                        case 0:
                            v130 = append(v130, -2147483648)
                        case 9223372036854775807:
                            if v137 {
                                var v142 int32 = -1
                                v137 = false
                                _ = v142
                            } else {
                                if ((v4 << ((65535) & 63)) == 1) {
                                    var v143 uint = 269
                                    if (-423 == -1) {
                                        var v144 bool = v137
                                        v145 := GenS1{F0: int(v143)}
                                        _ = v144
                                        _ = v145
                                    }
                                    v1 = (v1 + 0.0)
                                    var v146 uint8 = 0
                                    _ = v143
                                    _ = v146
                                } else {
                                    var v147 float64 = 1.5
                                    var v148 int16 = 861
                                    _ = v147
                                    _ = v148
                                }
                            }
                        default:
                            v149 := GenS15{F0: (v4 & v4)}
                            _ = v149
                        }
                        v137 = (1 > 127)
                    }
                    _ = fi140
                    _ = v141
                }
            case "\n\t":
                v150 := GenS16{F0: 2, F1: 1, F2: 1000}
                _ = v150
            default:
                var v151 int64 = (v2 * (v2 | v2))
                _ = v151
            }
            _ = fi138
            _ = v139
        }
    default:
        var v152 int64 = (-(v2))
        _ = v152
    }
    v1 = v1
    var v153 string = "\n\t"
    v154 := GenS17{F0: -23}
    if len(v130) > 0 {
        v130[0] = 32768
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v130:", len(v130), v130)
    fmt.Println("v137:", v137)
    fmt.Println("v153:", v153)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v130
    _ = v137
    _ = v153
    _ = v154
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
