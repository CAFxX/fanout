package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint; F1 uint; F2 int8 }
type GenS5 struct { F0 int; F1 int64; F2 bool }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int8; F1 string; F2 int64 }
type GenS11 struct { F0 uint; F1 int32; F2 int }
type GenS12 struct { F0 int32; F1 int32 }
type GenS13 struct { F0 int32; F1 uint32 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 bool }
func genFn0(p0 int64, p1 int) uint8 {
    var v4 int64 = p0
    if false {
        for fi5 := 0; fi5 < 15; fi5++ {
            switch p1 {
            case -1:
                v4 += v4
            case -96:
                var v6 bool = true
                _ = v6
            default:
                if (!true) {
                    fmt.Println(fi5, p0)
                    p0 -= (^(v4))
                    p0 -= (p0 * (p0 << ((1) & 63)))
                } else {
                    v4 = (-(v4))
                    v7 := GenS1{F0: 551, F1: (v4 % ((p0) | 1)), F2: (-(v4))}
                    _ = v7
                }
            }
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 int16 = -699
                _ = fi8
                _ = v9
            }
            var v10 uint16 = 1
            _ = fi5
            _ = v10
        }
    } else {
        if false {
            var v11 float64 = float64(3.14159)
            var v12 int64 = v4
            _ = v11
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    var v1 bool = (1e10 < 3.14159)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    var v13 bool = (1 != 1000)
    if (!v13) {
        if (!(true || ((!true) || true))) {
            v13 = v13
            v1 = v13
            fmt.Println(v1, v1)
        } else {
            var v14 int32 = 0
            v15 := []int{}
            if ((!((65535 != 0) || false)) && (1.5 <= 1.5)) {
                if (42 > -32768) {
                    v16 := GenS2{F0: 0}
                    if (-9223372036854775808 < 0) {
                        if (float64((0.0 / ((1e10 - (3.14159 / -2.25)) * (1.5 - (3.14159 * 3.14159))))) == float64(1e10)) {
                            switch 127 {
                            case 32768:
                                v17 := GenS2{F0: 1}
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 20; fi18++ {
                                if v13 {
                                    v3[42] = fi18
                                    var v19 uint16 = 953
                                    switch "\n\t" {
                                    case "a":
                                        v13 = ((v19 / ((65535) | 1)) == v19)
                                    case "foo bar":
                                        v13 = (1.5 == 0.1)
                                    }
                                    _ = v19
                                }
                                _ = fi18
                            }
                        }
                        v20 := GenS1{F0: 1, F1: 590, F2: -9223372036854775808}
                        var v21 int32 = v14
                        _ = v20
                        _ = v21
                    }
                    if (516 == 778) {
                        var v22 int16 = 1
                        _ = v22
                    }
                    _ = v16
                } else {
                    if ((!(v1 || (126 == 2))) || (v13 && true)) {
                        switch ("hello" + "foo bar") {
                        case "foo bar":
                            var v23 int = -1
                            _ = v23
                        default:
                            var v24 int8 = 127
                            _ = v24
                        }
                    } else {
                        var v25 int = 128
                        var v26 int64 = 42
                        v27 := GenS3{F0: 18446744073709551615, F1: 42}
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    if v13 {
                        var v28 float64 = float64(-2.25)
                        _ = v28
                    }
                    _ = len(v15)
                }
                var v29 string = "a"
                _ = v29
            }
            _ = v14
            _ = v15
        }
        v30 := []int{}
        if v1 {
            for fi31 := 0; fi31 < 4; fi31++ {
                var v32 uint64 = 0
                var v33 uint8 = 114
                _ = fi31
                _ = v32
                _ = v33
            }
        } else {
            var v34 int32 = 42
            var v35 string = "\n\t"
            _ = v34
            _ = v35
        }
        _ = v30
    }
    if true {
        v13 = (v1 || true)
        if (1 != 0) {
            for fi36 := 0; fi36 < 14; fi36++ {
                for fi37 := 0; fi37 < 8; fi37++ {
                    var v38 string = "hello"
                    _ = fi37
                    _ = v38
                }
                _ = fi36
            }
            fmt.Println(v13, v13)
            switch "" {
            case "":
                v1 = (-2.25 >= 3.14159)
            }
        } else {
            switch "hello" {
            case "hello":
                if true {
                    for fi39 := 0; fi39 < 4; fi39++ {
                        if (0.1 < (-2.25 + (1e10 * (3.14159 * float64(1.5))))) {
                            v40 := GenS3{F0: 145, F1: 4294967295}
                            if (!false) {
                                v41 := []int{fi39, 32767, (fi39 % ((42) | 1)), (fi39 & fi39)}
                                v13 = false
                                _ = v41
                            } else {
                                var v42 float64 = ((((1.5 / 1e10) / 1.0) / 3.14159) - (1e10 + -2.25))
                                var v43 bool = (!((fi39 + fi39) < (fi39 ^ fi39)))
                                _ = v42
                                _ = v43
                            }
                            _ = v40
                        } else {
                            for fi44 := 0; fi44 < 19; fi44++ {
                                var v45 int = (fi44 & (fi44 | (fi39 << ((2) & 63))))
                                _ = fi44
                                _ = v45
                            }
                            v1 = v1
                            v13 = v1
                        }
                        for fi46 := 0; fi46 < 6; fi46++ {
                            v1 = ((1e10 / float32(0.1)) > 0.0)
                            _ = fi46
                        }
                        _ = fi39
                    }
                    if v13 {
                        var v47 float32 = 0.1
                        for fi48 := 0; fi48 < 17; fi48++ {
                            var v49 float64 = (1e10 + 0.1)
                            var v50 int16 = 32767
                            v51 := GenS4{F0: 93, F1: 488, F2: -46}
                            _ = fi48
                            _ = v49
                            _ = v50
                            _ = v51
                        }
                        _ = v47
                    }
                    switch -1 {
                    case -1:
                        switch 4294967295 {
                        case 587:
                            v1 = true
                        case -2147483648:
                            fmt.Println(v13, v1)
                        }
                    default:
                        switch -1 {
                        case -1:
                            fmt.Println(v13, v13)
                        case 2147483647:
                            if (889 > -463) {
                                switch "a" {
                                case "foo bar":
                                    if (65535 < 727) {
                                        switch 32768 {
                                        case 32768:
                                            var v52 uint = 65535
                                            _ = v52
                                        case 910:
                                            v53 := GenS5{F0: 0, F1: 9223372036854775807, F2: v13}
                                            _ = v53
                                        }
                                        switch ("hello" + ("foo bar" + "a")) {
                                        case "foo bar":
                                            if true {
                                                v13 = false
                                                var v54 float32 = 3.14159
                                                _ = v54
                                            }
                                        case "a":
                                            fmt.Println(v1)
                                        case "\n\t":
                                            _ = len(v2)
                                        default:
                                            fmt.Println(v1, v1)
                                        }
                                        if true {
                                            switch "foo bar" {
                                            case "\n\t":
                                                fmt.Println(v1, v1)
                                            case "a":
                                                v55 := GenS1{F0: 4294967295, F1: -9223372036854775808, F2: 9223372036854775807}
                                                _ = v55
                                            case "":
                                                for fi56 := 0; fi56 < 16; fi56++ {
                                                    var v57 uint32 = 1
                                                    _ = fi56
                                                    _ = v57
                                                }
                                            default:
                                                v13 = ((!v1) || (495 != 0))
                                            }
                                        } else {
                                            switch -617 {
                                            case -9223372036854775808:
                                                var v58 bool = (-1 > 43)
                                                _ = v58
                                            case 42:
                                                var v59 uint8 = genFn0(9223372036854775807, 42)
                                                _ = v59
                                            }
                                            v1 = true
                                            var v60 uint64 = 18446744073709551615
                                            _ = v60
                                        }
                                    } else {
                                        var v61 uint8 = 0
                                        if true {
                                            switch ("a" + "foo bar") {
                                            case "foo bar":
                                                for fi62 := 0; fi62 < 6; fi62++ {
                                                    for fi63 := 0; fi63 < 20; fi63++ {
                                                        v64 := []int{(fi63 - 32768), fi63}
                                                        v1 = (!v1)
                                                        var v65 int = -2147483648
                                                        if (10 > 42) {
                                                            v66 := GenS2{F0: 255}
                                                            v2[2] = ((("\n\t" + "") + "\n\t") + "foo bar")
                                                            _ = v66
                                                        } else {
                                                            var v67 uint16 = 65535
                                                            _ = v67
                                                        }
                                                        delete(v3, 0)
                                                        _ = fi63
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    _ = fi62
                                                }
                                            }
                                            fmt.Println(v61, v13)
                                        }
                                        _ = v61
                                    }
                                }
                                v68 := []int{-518, -9223372036854775808, 32768}
                                v1 = ((255 == 194) || ((v1 || v13) && (!v1)))
                                _ = v68
                            }
                        }
                    }
                } else {
                    v1 = (9223372036854775807 != 277)
                    var v69 uint = 0
                    _ = v69
                }
            }
            for fi70 := 0; fi70 < 2; fi70++ {
                var v71 float32 = 0.0
                fmt.Println(v1, v1)
                v72 := GenS4{F0: 65535, F1: 4294967295, F2: 0}
                _ = fi70
                _ = v71
                _ = v72
            }
            switch ("\n\t" + "") {
            case "foo bar":
                for fi73 := 0; fi73 < 3; fi73++ {
                    v74 := GenS6{F0: 42, F1: 1}
                    var v75 uint32 = 1
                    _ = fi73
                    _ = v74
                    _ = v75
                }
            case "\n\t":
                v1 = (!((!((("a" + "\n\t") < "hello") || v1)) || ((((float32(((0.0 - 1e10) + 0.1)) * 1.5) * 0.0) >= 3.14159) || (((("" + "") + "a") + ("hello" + (("\n\t" + "foo bar") + ("hello" + (("hello" + "a") + "foo bar"))))) < "\n\t"))))
            default:
                var v76 uint8 = genFn0(-1, 9223372036854775807)
                _ = v76
            }
        }
        v77 := GenS6{F0: -128, F1: 32767}
        _ = v77
    }
    for fi78 := 0; fi78 < 16; fi78++ {
        for fi79 := 0; fi79 < 4; fi79++ {
            fmt.Println(v1, fi78, v13)
            _ = fi79
        }
        switch "\n\t" {
        case "\n\t":
            v80 := v2[2147483648]
            _ = v80
        case "foo bar":
            var v81 float32 = 1.5
            _ = v81
        default:
            for fi82 := 0; fi82 < 5; fi82++ {
                var v83 uint64 = 18446744073709551615
                v13 = true
                var v84 float64 = 0.0
                fmt.Println(v13, fi78, v1)
                v85 := GenS4{F0: 18446744073709551615, F1: 18446744073709551615, F2: 42}
                delete(v3, 852)
                var v86 int32 = -1
                v1 = v13
                _ = fi82
                _ = v83
                _ = v84
                _ = v85
                _ = v86
            }
        }
        _ = fi78
    }
    v13 = true
    for fi87 := 0; fi87 < 3; fi87++ {
        _ = len(v3)
        var v88 float32 = 0.1
        _ = fi87
        _ = v88
    }
    _ = len(v2)
    var v89 int8 = 42
    for fi90 := 0; fi90 < 6; fi90++ {
        switch (("foo bar" + "foo bar") + "hello") {
        case "foo bar":
            if (42 <= 0) {
                var v91 int64 = 721
                if false {
                    for fi92 := 0; fi92 < 2; fi92++ {
                        switch fi92 {
                        case -2147483648:
                            v93 := GenS7{F0: (genFn0(v91, 1) != 42)}
                            _ = v93
                        case 42:
                            fmt.Println(fi90)
                        default:
                            v94 := []int{}
                            _ = v94
                        }
                        switch "a" {
                        case "a":
                            v91 = -9223372036854775808
                        case "hello":
                            switch fi90 {
                            case -9223372036854775808:
                                if (3.14159 >= (1.5 * 1.5)) {
                                    switch ("\n\t" + "hello") {
                                    case "":
                                        switch "foo bar" {
                                        case "\n\t":
                                            if ((!(!(1 > (v91 ^ (^(v91)))))) && v13) {
                                                v1 = (2 > 0)
                                            } else {
                                                fmt.Println(fi92, v1)
                                                v89 -= v89
                                                fmt.Println(fi92, v91)
                                                if (4294967295 < 2) {
                                                    if v13 {
                                                        v13 = v1
                                                        if v13 {
                                                            v91 -= (v91 ^ v91)
                                                            v13 = (!(true || false))
                                                        }
                                                    } else {
                                                        if v1 {
                                                            v13 = v1
                                                        } else {
                                                            v95 := []int{1}
                                                            var v96 uint = 65535
                                                            _ = v95
                                                            _ = v96
                                                        }
                                                    }
                                                    var v97 uint8 = 162
                                                    _ = v97
                                                }
                                                v98 := GenS8{F0: 2}
                                                _ = v98
                                            }
                                        default:
                                            switch (-(fi90)) {
                                            case -553:
                                                var v99 float64 = (-2.25 * (0.0 + float64((3.14159 / 0.1))))
                                                v100 := GenS2{F0: 272}
                                                _ = v100
                                                _ = v99
                                            default:
                                                var v101 float32 = 1.5
                                                _ = v101
                                            }
                                        }
                                        var v102 bool = v13
                                        _ = v102
                                    case "hello":
                                        v103 := GenS7{F0: false}
                                        _ = v103
                                    }
                                    fmt.Println(v91, v89)
                                    var v104 float64 = -2.25
                                    _ = v104
                                }
                            }
                        case "foo bar":
                            var v105 string = ""
                            _ = v105
                        }
                        _ = fi92
                    }
                    if v1 {
                        var v106 float64 = ((0.1 - 1.5) / 1.0)
                        fmt.Println(fi90)
                        _ = v106
                    }
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        var v107 string = "\n\t"
                        _ = v107
                    }
                }
                v108 := []int{(fi90 & fi90), -9223372036854775808, 9223372036854775807, fi90}
                _ = v108
                _ = v91
            }
        case "":
            var v109 bool = v1
            _ = v109
        default:
            v110 := GenS2{F0: 0}
            _ = v110
        }
        for fi111 := 0; fi111 < 19; fi111++ {
            switch "a" {
            case "foo bar":
                v112 := GenS9{F0: 281}
                fmt.Println(v13, fi111, v89)
                _ = v112
            case "a":
                if (42 != 1) {
                    if true {
                        fmt.Println(fi111)
                    } else {
                        if false {
                            v1 = v1
                        } else {
                            fmt.Println(v13, v89, v13)
                            var v113 uint = 18446744073709551615
                            _ = v113
                        }
                    }
                    v13 = v1
                    v114 := []int{}
                    _ = v114
                } else {
                    if v1 {
                        var v115 int8 = 42
                        _ = v115
                    } else {
                        v89 += (v89 & 74)
                        switch fi111 {
                        case 4294967295:
                            var v116 int32 = 42
                            _ = v116
                        default:
                            if ((v13 || v1) && true) {
                                switch ("" + "\n\t") {
                                case "foo bar":
                                    v89 -= v89
                                case "\n\t":
                                    v117 := GenS10{F0: v89, F1: (("\n\t" + "") + ("a" + ("\n\t" + "\n\t"))), F2: 1}
                                    _ = v117
                                default:
                                    if (!v1) {
                                        switch ("a" + "") {
                                        case "foo bar":
                                            v89 -= (v89 ^ (^(v89)))
                                        case "hello":
                                            var v118 uint32 = 352
                                            _ = v118
                                        default:
                                            v1 = v13
                                        }
                                        _ = len(v3)
                                    }
                                }
                                if v1 {
                                    var v119 int8 = v89
                                    _ = v119
                                } else {
                                    switch (fi111 + 4294967295) {
                                    case 597:
                                        fmt.Println(fi111, v13)
                                    case 1:
                                        var v120 float32 = (-2.25 - 1e10)
                                        _ = v120
                                    case 0:
                                        v121 := GenS1{F0: 2, F1: 0, F2: 9223372036854775807}
                                        _ = v121
                                    default:
                                        var v122 int = (fi111 >> ((222) & 63))
                                        _ = v122
                                    }
                                    var v123 float32 = 0.0
                                    if (-357 > -1) {
                                        switch (fi90 * -1) {
                                        case 503:
                                            v124 := GenS8{F0: 159}
                                            _ = v124
                                        case 4294967295:
                                            var v125 float32 = (((((float32(float64(v123)) - v123) * 3.14159) + (v123 * 0.1)) / v123) / (v123 / v123))
                                            _ = v125
                                        case -416:
                                            v126 := GenS11{F0: 867, F1: 138, F2: fi111}
                                            _ = v126
                                        default:
                                            v89 += 127
                                            var v127 bool = v13
                                            _ = v127
                                        }
                                        if (111 <= fi111) {
                                            v123 -= float32((((0.0 - 1.5) + 0.1) - (((-2.25 + -2.25) * ((0.0 / 0.1) * (1.5 + 1.5))) * 1e10)))
                                            if (v13 && true) {
                                                var v128 int32 = 2147483647
                                                _ = v128
                                            } else {
                                                if ("a" != ("\n\t" + ("\n\t" + ""))) {
                                                    if v13 {
                                                        fmt.Println(fi111, v89, v1)
                                                        fmt.Println(fi90, v89, v89)
                                                    } else {
                                                        v129 := []int{fi111}
                                                        fmt.Println(fi111)
                                                        _ = v129
                                                    }
                                                    v13 = ((1 == 42) && false)
                                                    var v130 string = "\n\t"
                                                    _ = v130
                                                }
                                                if v1 {
                                                    v2[4294967295] = ("a" + ("hello" + ("hello" + (("" + "\n\t") + "a"))))
                                                    v1 = false
                                                    if ((("foo bar" + "hello") > ("hello" + "")) && (223 < 65535)) {
                                                        v131 := []int{0}
                                                        _ = v131
                                                    } else {
                                                        v13 = ((65535 == 2147483648) || v1)
                                                        var v132 int = (fi90 / (((^(fi90))) | 1))
                                                        _ = v132
                                                    }
                                                } else {
                                                    if (v1 || ((("foo bar" + "foo bar") + "a") <= "")) {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            switch fi111 {
                                                            case -1:
                                                                v133 := []int{fi90}
                                                                _ = v133
                                                            case 128:
                                                                v134 := GenS3{F0: 18446744073709551615, F1: 675}
                                                                _ = v134
                                                            default:
                                                                if (!v1) {
                                                                    v3[42] = fi90
                                                                    if v1 {
                                                                        v1 = false
                                                                        fmt.Println(v1)
                                                                    }
                                                                    v1 = v13
                                                                }
                                                            }
                                                        case "":
                                                            v1 = v13
                                                        case "foo bar":
                                                            switch fi111 {
                                                            case -9223372036854775808:
                                                                v123 -= 0.1
                                                            default:
                                                                v123 = v123
                                                            }
                                                        default:
                                                            v89 += (v89 - (-(v89)))
                                                        }
                                                        var v135 uint32 = 2147483647
                                                        var v136 uint8 = 255
                                                        _ = v135
                                                        _ = v136
                                                    }
                                                    fmt.Println(fi90)
                                                }
                                                fmt.Println(v1, v123, v123)
                                            }
                                            if false {
                                                v13 = true
                                                v137 := GenS12{F0: 406, F1: 0}
                                                _ = v137
                                            }
                                        } else {
                                            switch "a" {
                                            case "":
                                                var v138 uint16 = 42
                                                _ = v138
                                            case "a":
                                                switch (fi111 << ((65535) & 63)) {
                                                case 0:
                                                    v139 := []int{fi111, (^(fi90)), 42, fi90}
                                                    _ = v139
                                                case -9223372036854775808:
                                                    var v140 uint = 38
                                                    _ = v140
                                                case 950:
                                                    var v141 uint32 = 42
                                                    _ = v141
                                                default:
                                                    if (128 == (fi111 & fi111)) {
                                                        v89 = v89
                                                    } else {
                                                        v142 := GenS12{F0: 81, F1: 1}
                                                        v13 = (643 >= 0)
                                                        var v143 float64 = -2.25
                                                        var v144 uint = 2
                                                        switch (fi111 >> (((v144 * v144)) & 63)) {
                                                        case -775:
                                                            var v145 float64 = 0.0
                                                            _ = v145
                                                        }
                                                        if v1 {
                                                            if ((v89 - -1) > -128) {
                                                                v143 += v143
                                                                fmt.Println(v144, v123, v13)
                                                            }
                                                            var v146 int64 = -942
                                                            v147 := []int{4294967295, (^(fi90))}
                                                            delete(v3, 0)
                                                            _ = v146
                                                            _ = v147
                                                        } else {
                                                            var v148 int = -1
                                                            var v149 uint8 = 241
                                                            v150 := []int{v148}
                                                            _ = v148
                                                            _ = v149
                                                            _ = v150
                                                        }
                                                        var v151 uint32 = 632
                                                        if v1 {
                                                            var v152 int8 = 42
                                                            v123 -= (v123 - v123)
                                                            _ = v152
                                                        }
                                                        fmt.Println(v13, fi90)
                                                        v153 := []int{fi90, fi90, (fi111 | (fi90 * fi111))}
                                                        var v154 int8 = (v89 - v89)
                                                        _ = v142
                                                        _ = v143
                                                        _ = v144
                                                        _ = v151
                                                        _ = v153
                                                        _ = v154
                                                    }
                                                }
                                            }
                                            var v155 string = "a"
                                            _ = v155
                                        }
                                        var v156 uint8 = 2
                                        _ = v156
                                    }
                                    v1 = (2 <= 42)
                                    _ = v123
                                }
                                v1 = v13
                                if v1 {
                                    fmt.Println(v1, v89, fi111)
                                    fmt.Println(v89, fi90, v1)
                                    var v157 uint64 = 1
                                    _ = v157
                                } else {
                                    v1 = v1
                                    switch ("foo bar" + "hello") {
                                    case "\n\t":
                                        v1 = ("" < (("hello" + ("\n\t" + (("a" + "\n\t") + "\n\t"))) + "a"))
                                    case "foo bar":
                                        var v158 uint64 = 0
                                        _ = v158
                                    default:
                                        if (0 != 1) {
                                            switch fi111 {
                                            case -261:
                                                v13 = v1
                                            case 0:
                                                switch -917 {
                                                case -697:
                                                    v1 = v13
                                                default:
                                                    v89 = v89
                                                }
                                            case 153:
                                                fmt.Println(v89, v1)
                                            }
                                        }
                                    }
                                    v159 := GenS7{F0: (fi111 > fi111)}
                                    _ = v159
                                }
                                delete(v2, 2147483647)
                            } else {
                                v160 := GenS13{F0: 0, F1: 4294967295}
                                var v161 float32 = float32(1e10)
                                _ = v160
                                _ = v161
                            }
                        }
                    }
                }
            case "hello":
                var v162 int64 = -806
                _ = v162
            }
            if (0 < 18446744073709551615) {
                var v163 uint16 = 65535
                if (-2147483648 < -1) {
                    v1 = false
                    fmt.Println(fi90, fi111, v13)
                    var v164 int8 = v89
                    if true {
                        v165 := GenS10{F0: (v164 + v89), F1: (("" + "foo bar") + "a"), F2: 1}
                        _ = v165
                    }
                    var v166 uint64 = 1
                    _ = v164
                    _ = v166
                } else {
                    v167 := []int{(-(fi90)), (fi90 + fi111)}
                    _ = v167
                }
                _ = v163
            } else {
                v89 += (-(v89))
                v13 = (32767 == fi90)
                var v168 uint = 42
                if v1 {
                    fmt.Println(fi90, fi90, v168)
                    var v169 int16 = 0
                    v170 := []int{}
                    _ = v169
                    _ = v170
                }
                _ = v168
            }
            v1 = (!(-2.25 <= -2.25))
            _ = fi111
        }
        var v171 uint = 787
        _ = fi90
        _ = v171
    }
    v172 := GenS14{F0: 986}
    var v173 float64 = 1.5
    if v13 {
        v1 = true
        var v174 int32 = 399
        _ = v174
    }
    v173 = v173
    var v175 uint16 = 42
    for fi176 := 0; fi176 < 10; fi176++ {
        v177 := GenS15{F0: v1}
        var v178 uint = 1
        _ = fi176
        _ = v177
        _ = v178
    }
    var v179 uint8 = 127
    fmt.Println(v175, v89, v13)
    v180 := GenS4{F0: 42, F1: 42, F2: -121}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v173:", v173)
    fmt.Println("v175:", v175)
    fmt.Println("v179:", v179)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v89:", v89)
    _ = v1
    _ = v13
    _ = v172
    _ = v173
    _ = v175
    _ = v179
    _ = v180
    _ = v2
    _ = v3
    _ = v89
}
