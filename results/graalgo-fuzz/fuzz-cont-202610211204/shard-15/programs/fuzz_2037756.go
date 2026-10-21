package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint; F2 int }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 uint16; F2 uint }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS11 struct { F0 int64; F1 int64 }
type GenS12 struct { F0 int64; F1 uint64 }
type GenS13 struct { F0 uint; F1 uint8 }
type GenS14 struct { F0 uint8 }

func main() {
    var v1 int16 = 264
    var v2 int32 = -271
    v3 := make(map[int]int)
    var v4 uint16 = 42
    fmt.Println(v2, v4, v2)
    fmt.Println(v1)
    var v5 string = (("foo bar" + "") + ("" + "hello"))
    var v6 uint8 = 0
    switch (v5 + "hello") {
    case "":
        if (!false) {
            var v7 bool = (0.0 >= 3.14159)
            for fi8 := 0; fi8 < 15; fi8++ {
                switch ("foo bar" + ((((v5 + (("a" + v5) + "")) + ((v5 + v5) + "hello")) + v5) + v5)) {
                case "foo bar":
                    switch (v5 + "hello") {
                    case "hello":
                        v9 := GenS1{F0: v4, F1: v4}
                        _ = v9
                    case "\n\t":
                        fmt.Println(v4, v4, v6)
                    default:
                        if (!(9223372036854775807 >= -1)) {
                            v6 += v6
                        }
                    }
                }
                v2 = v2
                for fi10 := 0; fi10 < 7; fi10++ {
                    switch fi8 {
                    case -354:
                        if v7 {
                            v11 := GenS2{F0: v2, F1: 1, F2: (fi10 * fi10)}
                            v12 := v3[(fi8 - -183)]
                            _ = v11
                            _ = v12
                        } else {
                            v1 -= v1
                        }
                    case 299:
                        if false {
                            v13 := GenS3{F0: -89, F1: -128}
                            _ = v13
                        }
                    default:
                        switch 503 {
                        case 4294967295:
                            var v14 int8 = -81
                            _ = v14
                        }
                        var v15 int16 = (v1 | v1)
                        _ = v15
                    }
                    _ = fi10
                }
                _ = fi8
            }
            if (v7 && true) {
                v16 := GenS2{F0: v2, F1: 18446744073709551615, F2: 4294967295}
                fmt.Println(v5, v2)
                _ = v16
            }
            _ = v7
        } else {
            v3[128] = 127
        }
    case "hello":
        v17 := []int{-2147483648, 0, 42}
        _ = v17
    case "\n\t":
        for fi18 := 0; fi18 < 1; fi18++ {
            fmt.Println(v6, v2)
            for fi19 := 0; fi19 < 1; fi19++ {
                v20 := v3[-2147483648]
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    default:
        if (false && true) {
            switch -1 {
            case -799:
                v1 = (v1 % ((v1) | 1))
            case 9223372036854775807:
                v21 := []int{42, 42, 128}
                _ = v21
            case -478:
                v22 := GenS2{F0: (-(v2)), F1: 18446744073709551615, F2: 4294967295}
                _ = v22
            default:
                v23 := GenS4{F0: 128, F1: v6}
                _ = v23
            }
            if false {
                var v24 bool = ((2 > 0) || false)
                _ = v24
            }
            var v25 float32 = 0.0
            _ = v25
        } else {
            var v26 uint32 = 4294967295
            switch int(v26) {
            case 389:
                for fi27 := 0; fi27 < 13; fi27++ {
                    v6 = v6
                    var v28 int16 = (v1 >> ((0) & 63))
                    _ = fi27
                    _ = v28
                }
            default:
                switch 259 {
                case -2147483648:
                    switch 2147483647 {
                    case 32768:
                        for fi29 := 0; fi29 < 20; fi29++ {
                            v2 -= (v2 - (v2 - (v2 & 0)))
                            switch (fi29 * (fi29 << ((0) & 63))) {
                            case 9223372036854775807:
                                v30 := GenS5{F0: -9223372036854775808}
                                _ = v30
                            case 1:
                                switch v5 {
                                case "":
                                    for fi31 := 0; fi31 < 10; fi31++ {
                                        switch fi29 {
                                        case -607:
                                            switch "a" {
                                            case "foo bar":
                                                v32 := GenS3{F0: 0, F1: 100}
                                                _ = v32
                                            case "hello":
                                                v33 := []int{fi29}
                                                _ = v33
                                            case "\n\t":
                                                v34 := GenS3{F0: 0, F1: 42}
                                                _ = v34
                                            default:
                                                v5 = v5
                                            }
                                        default:
                                            switch "\n\t" {
                                            case "foo bar":
                                                v35 := GenS3{F0: 0, F1: 127}
                                                _ = v35
                                            case "a":
                                                if (255 < v6) {
                                                    v36 := []int{}
                                                    _ = v36
                                                } else {
                                                    var v37 float32 = 0.1
                                                    _ = v37
                                                }
                                            case "":
                                                delete(v3, (fi31 * (fi29 + 1)))
                                            }
                                        }
                                        _ = len(v3)
                                        switch 32768 {
                                        case -540:
                                            v1 = -1
                                        }
                                        _ = fi31
                                    }
                                case "hello":
                                    fmt.Println(v26, v5, fi29)
                                case "a":
                                    if false {
                                        switch (fi29 & (fi29 * (fi29 ^ fi29))) {
                                        case 2147483647:
                                            v38 := []int{}
                                            _ = v38
                                        case 228:
                                            v4 = (v4 | v4)
                                        }
                                        var v39 float64 = float64(0.1)
                                        for fi40 := 0; fi40 < 4; fi40++ {
                                            v41 := GenS3{F0: -36, F1: 127}
                                            if (!false) {
                                                if false {
                                                    var v42 uint16 = v4
                                                    switch (("" + ("foo bar" + "a")) + "") {
                                                    case "a":
                                                        v4 -= 1
                                                    case "\n\t":
                                                        if true {
                                                            if false {
                                                                var v43 int16 = (-(v1))
                                                                var v44 int64 = 40
                                                                _ = v43
                                                                _ = v44
                                                            } else {
                                                                v45 := GenS3{F0: 42, F1: -122}
                                                                v2 = 2147483647
                                                                v2 -= 352
                                                                _ = v45
                                                            }
                                                            v46 := GenS2{F0: v2, F1: 393, F2: int(v26)}
                                                            _ = v46
                                                        }
                                                    case "foo bar":
                                                        switch fi29 {
                                                        case -1:
                                                            switch ((v5 + "hello") + (v5 + (v5 + ((("" + v5) + v5) + ("hello" + "hello"))))) {
                                                            case "a":
                                                                var v47 float32 = 1.5
                                                                _ = v47
                                                            case "":
                                                                v4 = v42
                                                            }
                                                        }
                                                    default:
                                                        switch v5 {
                                                        case "":
                                                            v5 = v5
                                                        }
                                                    }
                                                    _ = v42
                                                }
                                            } else {
                                                var v48 float64 = v39
                                                delete(v3, fi40)
                                                if ((v26 / ((848) | 1)) > 4294967295) {
                                                    v49 := GenS1{F0: (v4 << ((4294967295) & 63)), F1: (v4 * 65535)}
                                                    v50 := []int{9223372036854775807, 530, 127, (fi29 / ((0) | 1))}
                                                    _ = v49
                                                    _ = v50
                                                }
                                                _ = v48
                                            }
                                            var v51 uint16 = v4
                                            _ = fi40
                                            _ = v41
                                            _ = v51
                                        }
                                        _ = v39
                                    } else {
                                        v3[(fi29 >> ((18446744073709551615) & 63))] = -482
                                        fmt.Println(v5, v6, v6)
                                    }
                                }
                            case 42:
                                v2 += (v2 & v2)
                            default:
                                v52 := GenS6{F0: 39, F1: (v4 * v4), F2: 2}
                                _ = v52
                            }
                            if (-2.25 <= (((1.5 - (0.0 - float64((0.0 + 1.5)))) / -2.25) * 3.14159)) {
                                v3[4294967295] = fi29
                                switch 2147483647 {
                                case -1:
                                    var v53 string = v5
                                    _ = v53
                                case 32767:
                                    var v54 int32 = (-(v2))
                                    _ = v54
                                case 42:
                                    if false {
                                        switch (("\n\t" + v5) + "a") {
                                        case "\n\t":
                                            switch v5 {
                                            case "\n\t":
                                                v55 := []int{128}
                                                _ = v55
                                            case "foo bar":
                                                for fi56 := 0; fi56 < 18; fi56++ {
                                                    if true {
                                                        var v57 float32 = -2.25
                                                        switch 4294967295 {
                                                        case 127:
                                                            v26 = 0
                                                        default:
                                                            v4 = v4
                                                        }
                                                        v58 := GenS5{F0: 42}
                                                        _ = v57
                                                        _ = v58
                                                    } else {
                                                        var v59 bool = true
                                                        _ = v59
                                                    }
                                                    switch (v5 + "foo bar") {
                                                    case "\n\t":
                                                        v6 += uint8(v4)
                                                    case "":
                                                        if (uint64(v6) > 842) {
                                                            if (118 != 94) {
                                                                if ((v26 << ((277) & 63)) > uint32(v4)) {
                                                                    switch -2147483648 {
                                                                    case 4294967295:
                                                                        switch (fi29 - (fi56 & fi56)) {
                                                                        case 127:
                                                                            v26 = v26
                                                                        }
                                                                    }
                                                                    if ((!(42 != uint64(v4))) || false) {
                                                                        var v60 float64 = float64(-2.25)
                                                                        if (fi56 >= fi56) {
                                                                            v61 := GenS2{F0: (^(v2)), F1: 914, F2: 1}
                                                                            v3[9223372036854775807] = fi56
                                                                            _ = v61
                                                                        }
                                                                        _ = v60
                                                                    }
                                                                    if true {
                                                                        var v62 uint = 255
                                                                        v6 = v6
                                                                        switch (fi29 - (fi29 / (((fi29 >> ((v62) & 63))) | 1))) {
                                                                        case 128:
                                                                            v63 := GenS1{F0: 32767, F1: v4}
                                                                            _ = v63
                                                                        case 562:
                                                                            _ = len(v3)
                                                                        default:
                                                                            v64 := GenS3{F0: -34, F1: -40}
                                                                            _ = v64
                                                                        }
                                                                        _ = v62
                                                                    }
                                                                }
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    _ = len(v3)
                                                                case "\n\t":
                                                                    v65 := GenS1{F0: 32767, F1: (v4 * 1)}
                                                                    _ = v65
                                                                }
                                                            } else {
                                                                var v66 uint32 = 340
                                                                v67 := GenS4{F0: (fi29 << ((255) & 63)), F1: v6}
                                                                _ = v66
                                                                _ = v67
                                                            }
                                                        }
                                                    default:
                                                        var v68 uint64 = 1
                                                        _ = v68
                                                    }
                                                    _ = fi56
                                                }
                                            }
                                        case "hello":
                                            fmt.Println(v5, v1, v1)
                                        }
                                        fmt.Println(v26, v5)
                                        for fi69 := 0; fi69 < 12; fi69++ {
                                            switch (fi29 << ((255) & 63)) {
                                            case -2147483648:
                                                switch "hello" {
                                                case "":
                                                    v70 := GenS5{F0: -957}
                                                    _ = v70
                                                case "foo bar":
                                                    var v71 float32 = 0.0
                                                    _ = v71
                                                default:
                                                    v4 += 32767
                                                }
                                            }
                                            fmt.Println(v2, fi69, v5)
                                            _ = fi69
                                        }
                                    } else {
                                        fmt.Println(v1, v6, v1)
                                    }
                                }
                            } else {
                                var v72 int16 = v1
                                for fi73 := 0; fi73 < 15; fi73++ {
                                    var v74 uint8 = (v6 ^ (v6 >> ((2) & 63)))
                                    _ = fi73
                                    _ = v74
                                }
                                v75 := []int{fi29}
                                _ = v72
                                _ = v75
                            }
                            _ = fi29
                        }
                    case 127:
                        var v76 string = v5
                        _ = v76
                    default:
                        if (false || false) {
                            for fi77 := 0; fi77 < 3; fi77++ {
                                var v78 uint64 = 2
                                _ = fi77
                                _ = v78
                            }
                            switch 47 {
                            case 9223372036854775807:
                                for fi79 := 0; fi79 < 3; fi79++ {
                                    v6 += v6
                                    _ = fi79
                                }
                            case 1:
                                v6 = (v6 - 1)
                            default:
                                var v80 uint = uint(v6)
                                _ = v80
                            }
                            switch (("hello" + v5) + v5) {
                            case "\n\t":
                                if (!(!true)) {
                                    for fi81 := 0; fi81 < 5; fi81++ {
                                        var v82 uint = 4294967295
                                        _ = fi81
                                        _ = v82
                                    }
                                }
                            case "hello":
                                v83 := GenS7{F0: ("\n\t" + v5)}
                                _ = v83
                            case "":
                                for fi84 := 0; fi84 < 5; fi84++ {
                                    for fi85 := 0; fi85 < 17; fi85++ {
                                        v26 = (v26 & (v26 * v26))
                                        v86 := []int{}
                                        _ = fi85
                                        _ = v86
                                    }
                                    switch 4294967295 {
                                    case -69:
                                        switch ("a" + "a") {
                                        case "a":
                                            v1 = 32767
                                        }
                                    case 983:
                                        var v87 float64 = 0.1
                                        _ = v87
                                    default:
                                        if false {
                                            for fi88 := 0; fi88 < 10; fi88++ {
                                                if ((!(int32(v4) < int32(v6))) && false) {
                                                    if false {
                                                        v89 := GenS8{F0: uint32(v4), F1: (^(v6)), F2: -32768}
                                                        delete(v3, (fi88 % (((fi88 + int(v4))) | 1)))
                                                        _ = v89
                                                    } else {
                                                        if true {
                                                            if (!true) {
                                                                var v90 uint = 0
                                                                _ = v90
                                                            } else {
                                                                fmt.Println(v5, v2)
                                                                fmt.Println(fi88)
                                                                v91 := GenS1{F0: (^(v4)), F1: (v4 % (((^(v4))) | 1))}
                                                                var v92 string = (("" + v5) + ("hello" + ""))
                                                                switch (("\n\t" + "") + "hello") {
                                                                case "\n\t":
                                                                    var v93 uint = 18446744073709551615
                                                                    _ = v93
                                                                }
                                                                fmt.Println(v1, v1)
                                                                v6 = v6
                                                                _ = v91
                                                                _ = v92
                                                            }
                                                            v4 = v4
                                                            switch (("hello" + v5) + v5) {
                                                            case "a":
                                                                if true {
                                                                    v6 = 127
                                                                }
                                                            }
                                                        } else {
                                                            var v94 uint32 = (^(v26))
                                                            _ = v94
                                                        }
                                                        switch -598 {
                                                        case -527:
                                                            var v95 uint16 = (^(v4))
                                                            _ = v95
                                                        }
                                                    }
                                                    switch (fi88 / ((-9223372036854775808) | 1)) {
                                                    case 128:
                                                        var v96 int16 = (v1 & 1)
                                                        _ = v96
                                                    default:
                                                        v97 := GenS9{F0: (v26 % (((^(v26))) | 1)), F1: 0}
                                                        _ = v97
                                                    }
                                                    v98 := GenS10{F0: 42, F1: 812, F2: (v26 * (v26 / (((v26 / (((v26 * (v26 + v26))) | 1))) | 1)))}
                                                    _ = v98
                                                }
                                                v6 = (v6 + v6)
                                                _ = fi88
                                            }
                                            for fi99 := 0; fi99 < 12; fi99++ {
                                                switch "" {
                                                case "hello":
                                                    fmt.Println(v2, fi84)
                                                case "":
                                                    fmt.Println(v6, fi84, v4)
                                                case "\n\t":
                                                    var v100 uint32 = (v26 + (v26 >> ((2) & 63)))
                                                    _ = v100
                                                }
                                                v101 := GenS3{F0: 1, F1: -33}
                                                delete(v3, fi99)
                                                v5 = v5
                                                var v102 uint16 = (v4 * (v4 / (((^(v4))) | 1)))
                                                _ = fi99
                                                _ = v101
                                                _ = v102
                                            }
                                        } else {
                                            for fi103 := 0; fi103 < 17; fi103++ {
                                                v6 = v6
                                                v104 := []int{}
                                                _ = fi103
                                                _ = v104
                                            }
                                            v105 := GenS1{F0: v4, F1: v4}
                                            for fi106 := 0; fi106 < 19; fi106++ {
                                                fmt.Println(v5, v6)
                                                var v107 int8 = 70
                                                v108 := GenS1{F0: v4, F1: 1}
                                                switch ((v5 + v5) + ("\n\t" + v5)) {
                                                case "a":
                                                    var v109 int8 = int8(v2)
                                                    _ = v109
                                                case "foo bar":
                                                    var v110 int64 = 1
                                                    _ = v110
                                                case "":
                                                    v111 := GenS4{F0: fi84, F1: v6}
                                                    _ = v111
                                                }
                                                var v112 uint = 65535
                                                _ = fi106
                                                _ = v107
                                                _ = v108
                                                _ = v112
                                            }
                                            _ = v105
                                        }
                                    }
                                    v6 = v6
                                    _ = fi84
                                }
                            }
                        } else {
                            if (65535 > 935) {
                                switch "\n\t" {
                                case "\n\t":
                                    if true {
                                        fmt.Println(v2, v5, v1)
                                        if (526 == 0) {
                                            var v113 int64 = 0
                                            var v114 int64 = v113
                                            v26 = 65535
                                            _ = v113
                                            _ = v114
                                        }
                                        for fi115 := 0; fi115 < 11; fi115++ {
                                            for fi116 := 0; fi116 < 5; fi116++ {
                                                fmt.Println(v26, fi115)
                                                v6 = (v6 + v6)
                                                v6 += uint8(v4)
                                                v117 := GenS7{F0: ""}
                                                if false {
                                                    var v118 string = ""
                                                    v119 := []int{}
                                                    _ = v118
                                                    _ = v119
                                                } else {
                                                    var v120 int64 = 1
                                                    fmt.Println(v2, v26, v1)
                                                    v121 := []int{(-(fi115)), (fi115 << ((1) & 63))}
                                                    _ = v120
                                                    _ = v121
                                                }
                                                if false {
                                                    var v122 uint8 = (v6 << ((450) & 63))
                                                    fmt.Println(v26, fi115)
                                                    v6 = (^(v6))
                                                    _ = v122
                                                } else {
                                                    if (false || false) {
                                                        var v123 float64 = -2.25
                                                        switch (fi115 ^ fi116) {
                                                        case -865:
                                                            v1 = v1
                                                        case -1:
                                                            v124 := v3[-852]
                                                            _ = v124
                                                        case -2147483648:
                                                            fmt.Println(v4, v123)
                                                        }
                                                        _ = v123
                                                    }
                                                    switch -753 {
                                                    case 9223372036854775807:
                                                        var v125 uint32 = (v26 + 0)
                                                        _ = v125
                                                    case 32767:
                                                        v126 := []int{}
                                                        _ = v126
                                                    case 1:
                                                        fmt.Println(v2, v2, v5)
                                                    }
                                                    v127 := []int{-2147483648, 4294967295, 1}
                                                    var v128 float32 = (3.14159 - 0.0)
                                                    _ = v127
                                                    _ = v128
                                                }
                                                _ = fi116
                                                _ = v117
                                            }
                                            v129 := GenS11{F0: 1, F1: 9223372036854775807}
                                            for fi130 := 0; fi130 < 4; fi130++ {
                                                var v131 uint = 18446744073709551615
                                                _ = fi130
                                                _ = v131
                                            }
                                            _ = fi115
                                            _ = v129
                                        }
                                    } else {
                                        _ = len(v3)
                                        if true {
                                            var v132 int = 32767
                                            v26 = (^(v26))
                                            _ = v132
                                        } else {
                                            switch "\n\t" {
                                            case "hello":
                                                v133 := []int{-1, 9223372036854775807}
                                                _ = v133
                                            default:
                                                var v134 uint64 = 0
                                                _ = v134
                                            }
                                        }
                                        for fi135 := 0; fi135 < 3; fi135++ {
                                            fmt.Println(v1)
                                            var v136 string = "a"
                                            _ = fi135
                                            _ = v136
                                        }
                                    }
                                }
                                if (false || false) {
                                    v137 := GenS3{F0: 86, F1: -1}
                                    _ = v137
                                } else {
                                    for fi138 := 0; fi138 < 1; fi138++ {
                                        v139 := GenS3{F0: 0, F1: 111}
                                        var v140 uint32 = 2
                                        var v141 uint16 = v4
                                        _ = fi138
                                        _ = v139
                                        _ = v140
                                        _ = v141
                                    }
                                    v1 += (-(v1))
                                    v1 -= -575
                                }
                            } else {
                                fmt.Println(v6, v6, v5)
                                if (!((v26 % (((v26 << ((uint(v4)) & 63))) | 1)) >= 315)) {
                                    for fi142 := 0; fi142 < 20; fi142++ {
                                        var v143 int16 = 1000
                                        _ = fi142
                                        _ = v143
                                    }
                                    v144 := GenS12{F0: 1, F1: 42}
                                    _ = v144
                                } else {
                                    v6 += (v6 - v6)
                                    var v145 bool = (((v2 - v2) <= v2) || true)
                                    _ = v145
                                }
                            }
                            fmt.Println(v5, v6, v1)
                        }
                    }
                default:
                    v5 = ""
                }
            }
            _ = v26
        }
    }
    switch -9223372036854775808 {
    case 32767:
        v146 := GenS12{F0: 1, F1: 662}
        _ = v146
    case 32768:
        v6 = (v6 ^ (v6 + v6))
    default:
        switch -507 {
        case 42:
            for fi147 := 0; fi147 < 16; fi147++ {
                if (1e10 > (1.5 + 1.5)) {
                    v148 := GenS10{F0: 2, F1: 1, F2: 2}
                    var v149 float32 = (-2.25 * -2.25)
                    v150 := GenS13{F0: 937, F1: v6}
                    _ = v148
                    _ = v149
                    _ = v150
                }
                for fi151 := 0; fi151 < 4; fi151++ {
                    var v152 uint = 298
                    fmt.Println(v5, v2)
                    v153 := []int{-1, -9223372036854775808, fi147}
                    _ = fi151
                    _ = v152
                    _ = v153
                }
                switch ("foo bar" + "") {
                case "a":
                    var v154 uint8 = v6
                    _ = v154
                case "":
                    v5 = v5
                }
                _ = fi147
            }
        default:
            v155 := GenS7{F0: v5}
            _ = v155
        }
    }
    var v156 int = 0
    delete(v3, v156)
    switch (v156 | (v156 ^ (v156 / ((v156) | 1)))) {
    case -572:
        v4 -= (v4 % ((v4) | 1))
    case -661:
        v157 := GenS7{F0: ""}
        _ = v157
    }
    v158 := GenS14{F0: (v6 ^ v6)}
    if false {
        v159 := GenS7{F0: ((v5 + v5) + v5)}
        _ = v159
    } else {
        v160 := []int{}
        v5 = v5
        _ = v160
    }
    var v161 uint = 65535
    v162 := []int{(v156 & 909), 2147483647, (v156 - v156), v156}
    fmt.Println("v1:", v1)
    fmt.Println("v156:", v156)
    fmt.Println("v161:", v161)
    fmt.Println("slicelen v162:", len(v162), v162)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v156
    _ = v158
    _ = v161
    _ = v162
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
