package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 string; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint8; F1 int64 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int32; F1 int8 }
type GenS11 struct { F0 int8; F1 int8 }
type GenS12 struct { F0 string; F1 int16 }
type GenS13 struct { F0 int8; F1 uint; F2 bool }
type GenS14 struct { F0 int; F1 int32 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 int64; F1 uint16 }
type GenS17 struct { F0 int8; F1 int; F2 int }
type GenS18 struct { F0 string; F1 int }
type GenS19 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS20 struct { F0 int; F1 int8 }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 int; F1 uint8 }

func main() {
    var v1 uint16 = 1
    var v2 float32 = (((-2.25 + 3.14159) - -2.25) - 0.0)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    var v5 int16 = 824
    var v6 bool = (255 <= 1)
    var v7 bool = ((v1 / ((255) | 1)) == v1)
    v5 = (v5 * v5)
    for fi8 := 0; fi8 < 10; fi8++ {
        if v7 {
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 bool = true
                v2 = (float32(1.5) + (1.5 - v2))
                var v11 uint8 = 0
                _ = fi9
                _ = v10
                _ = v11
            }
            v12 := GenS1{F0: 0, F1: (v1 << ((1) & 63))}
            v13 := GenS2{F0: (("a" + "hello") + "hello"), F1: -1, F2: 249}
            _ = v12
            _ = v13
        } else {
            if (!true) {
                v5 = 42
                v14 := []int{(fi8 & (fi8 ^ (fi8 / ((fi8) | 1))))}
                if v6 {
                    var v15 int32 = -300
                    _ = v15
                }
                _ = v14
            } else {
                v16 := GenS3{F0: 376}
                for fi17 := 0; fi17 < 12; fi17++ {
                    if v6 {
                        var v18 int64 = -9223372036854775808
                        switch ((("foo bar" + ("" + (("a" + (("foo bar" + "\n\t") + "\n\t")) + "a"))) + (("" + ("a" + ("\n\t" + "\n\t"))) + "hello")) + ("hello" + "foo bar")) {
                        case "\n\t":
                            v18 -= -536
                        }
                        v19 := GenS4{F0: 2147483648, F1: v6, F2: 882}
                        v20 := GenS5{F0: -14}
                        switch ((("hello" + "\n\t") + "") + "a") {
                        case "":
                            var v21 int8 = 42
                            _ = v21
                        case "a":
                            fmt.Println(v2, v6, v1)
                        case "hello":
                            var v22 int64 = -1
                            _ = v22
                        default:
                            switch "" {
                            case "":
                                switch (fi8 & 0) {
                                case -516:
                                    v23 := GenS1{F0: 42, F1: (^(v1))}
                                    _ = v23
                                case 1:
                                    v24 := GenS5{F0: -1}
                                    _ = v24
                                case -2147483648:
                                    if (v7 && (v7 && (v6 && true))) {
                                        v25 := GenS4{F0: 65535, F1: (42 > 91), F2: 2}
                                        v26 := GenS6{F0: false, F1: 102}
                                        v2 -= (v2 / (0.0 - float32(1e10)))
                                        _ = v25
                                        _ = v26
                                    }
                                default:
                                    v7 = true
                                }
                            }
                        }
                        v2 += (v2 + v2)
                        var v27 int32 = 0
                        _ = v18
                        _ = v19
                        _ = v20
                        _ = v27
                    } else {
                        if true {
                            var v28 int8 = 0
                            _ = v28
                        } else {
                            v29 := GenS4{F0: 0, F1: (!false), F2: 201}
                            if v6 {
                                if ((v1 - 1) < v1) {
                                    v5 -= v5
                                }
                                v30 := GenS7{F0: 2147483647}
                                _ = v30
                            }
                            v5 = v5
                            v6 = false
                            _ = v29
                        }
                        var v31 uint32 = 0
                        _ = v31
                    }
                    v1 += v1
                    _ = fi17
                }
                v7 = ((-9223372036854775808 == -1) || (v6 && true))
                _ = v16
            }
            v7 = (127 < 42)
            var v32 int16 = (v5 >> ((810) & 63))
            _ = v32
        }
        _ = fi8
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        switch ("\n\t" + "a") {
        case "\n\t":
            if true {
                for fi34 := 0; fi34 < 17; fi34++ {
                    var v35 int64 = 1
                    _ = fi34
                    _ = v35
                }
                var v36 uint64 = 117
                v37 := GenS1{F0: 42, F1: (v1 << ((42) & 63))}
                _ = v36
                _ = v37
            }
        case "":
            if false {
                for fi38 := 0; fi38 < 10; fi38++ {
                    var v39 uint8 = 0
                    v5 = (v5 >> ((18446744073709551615) & 63))
                    if true {
                        v7 = (9223372036854775807 <= 848)
                    } else {
                        var v40 int16 = (v5 + (v5 / (((v5 % ((32767) | 1))) | 1)))
                        fmt.Println(fi38, v1, v7)
                        v41 := []int{1, fi33}
                        _ = v40
                        _ = v41
                    }
                    _ = fi38
                    _ = v39
                }
            } else {
                v5 = -1
            }
        default:
            v4[2147483648] = fi33
        }
        var v42 int32 = -543
        if true {
            v42 = v42
        }
        _ = fi33
        _ = v42
    }
    switch "hello" {
    case "":
        var v43 int = 239
        _ = v43
    default:
        var v44 int16 = (v5 % ((1000) | 1))
        _ = v44
    }
    delete(v4, 4294967295)
    switch "" {
    case "foo bar":
        var v45 uint32 = 1
        _ = v45
    default:
        if v7 {
            v1 -= 255
            for fi46 := 0; fi46 < 10; fi46++ {
                v47 := GenS7{F0: 42}
                _ = fi46
                _ = v47
            }
            v5 = -349
        } else {
            if (("a" + "foo bar") == ("foo bar" + (("hello" + ("\n\t" + ("\n\t" + "\n\t"))) + ""))) {
                v7 = v7
                var v48 int8 = 0
                _ = v48
            } else {
                v49 := v4[232]
                fmt.Println(v7, v7, v1)
                switch 42 {
                case -773:
                    v1 = 876
                default:
                    var v50 int = int(v5)
                    _ = v50
                }
                _ = v49
            }
            fmt.Println(v1, v5)
            v51 := GenS8{F0: 42, F1: -949}
            _ = v51
        }
    }
    for fi52 := 0; fi52 < 3; fi52++ {
        v4[2147483648] = 42
        var v53 uint = 1
        switch fi52 {
        case -9223372036854775808:
            v7 = (-154 < 42)
        case -2147483648:
            switch "a" {
            case "a":
                v54 := v3[32767]
                _ = v54
            case "\n\t":
                var v55 bool = (!(v5 >= v5))
                _ = v55
            default:
                v3[-849] = v6
            }
        case 9223372036854775807:
            if v6 {
                fmt.Println(v1)
            }
        }
        _ = fi52
        _ = v53
    }
    v56 := GenS1{F0: 0, F1: v1}
    var v57 uint16 = (v1 | v1)
    switch "hello" {
    case "hello":
        var v58 int32 = 0
        _ = v58
    }
    for fi59 := 0; fi59 < 1; fi59++ {
        switch "a" {
        case "hello":
            v60 := GenS2{F0: ("" + "hello"), F1: int64(fi59), F2: 174}
            _ = v60
        case "":
            var v61 uint = 435
            _ = v61
        case "\n\t":
            for fi62 := 0; fi62 < 8; fi62++ {
                switch "a" {
                case "hello":
                    switch ("hello" + "foo bar") {
                    case "a":
                        var v63 uint = 2
                        _ = v63
                    case "foo bar":
                        var v64 uint = 65535
                        v65 := GenS4{F0: 2147483647, F1: ((1 < -128) && (!(42 == 1))), F2: 2}
                        v66 := GenS5{F0: 100}
                        _ = v64
                        _ = v65
                        _ = v66
                    }
                    if ((v5 - (v5 - 1000)) > v5) {
                        fmt.Println(fi62, v5, v1)
                    } else {
                        _ = len(v4)
                        var v67 float64 = float64(v2)
                        switch "hello" {
                        case "a":
                            fmt.Println(fi59, v2)
                        default:
                            switch "" {
                            case "\n\t":
                                var v68 float32 = (v2 - 0.1)
                                _ = v68
                            case "":
                                fmt.Println(fi59, v67, v7)
                            }
                            if false {
                                var v69 int8 = 42
                                v69 = (^(v69))
                                v70 := v3[(fi59 % ((fi62) | 1))]
                                _ = v69
                                _ = v70
                            } else {
                                v71 := GenS9{F0: 1}
                                _ = v71
                            }
                        }
                        _ = v67
                    }
                    if v7 {
                        switch "\n\t" {
                        case "a":
                            if v6 {
                                v5 = 1
                                if ((!v6) && v6) {
                                    var v72 uint64 = 616
                                    v73 := GenS8{F0: 2, F1: 789}
                                    v5 = v5
                                    var v74 float64 = 0.0
                                    v1 -= (^(v57))
                                    fmt.Println(v72, fi59)
                                    _ = v72
                                    _ = v73
                                    _ = v74
                                } else {
                                    fmt.Println(v7, fi59, fi59)
                                }
                                if true {
                                    fmt.Println(v57)
                                    if v7 {
                                        v75 := GenS10{F0: 1, F1: 127}
                                        fmt.Println(fi59, fi59)
                                        v2 = (v2 / 1.0)
                                        _ = v75
                                    }
                                } else {
                                    var v76 uint16 = (v1 ^ (^(v57)))
                                    if false {
                                        v5 = v5
                                        if false {
                                            var v77 int8 = 106
                                            _ = v77
                                        }
                                    }
                                    _ = v76
                                }
                                var v78 uint32 = 42
                                _ = v78
                            }
                        case "\n\t":
                            v5 = v5
                        }
                    } else {
                        v79 := GenS11{F0: 100, F1: 0}
                        v80 := GenS12{F0: "\n\t", F1: (v5 * (v5 << ((255) & 63)))}
                        v6 = (1 >= int64(v1))
                        if v7 {
                            v1 += (v57 + 159)
                            v7 = v6
                        } else {
                            v81 := GenS13{F0: 1, F1: 255, F2: false}
                            var v82 uint = 1
                            var v83 float32 = v2
                            _ = v81
                            _ = v82
                            _ = v83
                        }
                        _ = v79
                        _ = v80
                    }
                default:
                    if (v7 && (-704 != -159)) {
                        var v84 string = "a"
                        _ = v84
                    } else {
                        var v85 string = ("foo bar" + "hello")
                        _ = v85
                    }
                }
                v86 := []int{}
                v3[(fi62 + (fi62 ^ 255))] = (!(-1 <= -1))
                _ = fi62
                _ = v86
            }
        default:
            if (1 > 214) {
                if (!v6) {
                    v7 = (!v6)
                    switch (fi59 | 32768) {
                    case -825:
                        v87 := GenS8{F0: 255, F1: 1}
                        _ = v87
                    case 32767:
                        if (9223372036854775807 > 0) {
                            var v88 int64 = -16
                            for fi89 := 0; fi89 < 16; fi89++ {
                                switch (("foo bar" + (("hello" + "\n\t") + "a")) + ("a" + "hello")) {
                                case "":
                                    v7 = (-1 <= (v5 << ((42) & 63)))
                                case "hello":
                                    switch ((((("foo bar" + ("a" + ("" + ("hello" + "hello")))) + (((("a" + ("hello" + "\n\t")) + (("" + "hello") + "foo bar")) + "") + "\n\t")) + ((("a" + ("\n\t" + ("a" + "\n\t"))) + "") + "a")) + ("hello" + "\n\t")) + "a") {
                                    case "foo bar":
                                        var v90 float32 = v2
                                        _ = v90
                                    default:
                                        switch 128 {
                                        case 127:
                                            if v6 {
                                                fmt.Println(v6)
                                                if (v7 && (14 == 1)) {
                                                    v57 += 1
                                                    v2 += 0.1
                                                    v91 := GenS1{F0: 128, F1: v1}
                                                    v92 := GenS1{F0: 42, F1: (v1 >> ((uint(v57)) & 63))}
                                                    if ((v88 % ((v88) | 1)) == v88) {
                                                        if v7 {
                                                            var v93 uint = 710
                                                            if v6 {
                                                                v94 := []int{fi89, 0}
                                                                if ((v5 >= v5) && ((v5 | 835) == (v5 | v5))) {
                                                                    var v95 int64 = int64(v93)
                                                                    _ = v95
                                                                } else {
                                                                    var v96 int8 = -128
                                                                    _ = v96
                                                                }
                                                                if v6 {
                                                                    _ = len(v94)
                                                                    var v97 int = fi89
                                                                    var v98 float32 = ((v2 + (((v2 - float32(1e10)) + 0.1) * (1.5 / (-2.25 + v2)))) * 1e10)
                                                                    _ = v97
                                                                    _ = v98
                                                                }
                                                                _ = v94
                                                            } else {
                                                                v99 := v4[75]
                                                                if v6 {
                                                                    v5 = (v5 - (v5 & -407))
                                                                    switch (fi59 - (fi59 / (((fi89 << ((v93) & 63))) | 1))) {
                                                                    case 42:
                                                                        fmt.Println(v6, v2)
                                                                    default:
                                                                        if true {
                                                                            v1 -= v57
                                                                            switch "a" {
                                                                            case "a":
                                                                                var v100 uint32 = 510
                                                                                _ = v100
                                                                            }
                                                                        }
                                                                    }
                                                                    v101 := GenS11{F0: 0, F1: int8(v93)}
                                                                    fmt.Println(v88, v5, v2)
                                                                    _ = v101
                                                                }
                                                                fmt.Println(v6)
                                                                _ = v99
                                                            }
                                                            if v6 {
                                                                var v102 uint16 = (v1 - v1)
                                                                v6 = v7
                                                                v103 := GenS14{F0: fi89, F1: -2147483648}
                                                                _ = v102
                                                                _ = v103
                                                            } else {
                                                                v104 := GenS5{F0: 42}
                                                                _ = v104
                                                            }
                                                            if (18446744073709551615 > 18446744073709551615) {
                                                                fmt.Println(v7, v2, v5)
                                                            } else {
                                                                fmt.Println(v5, v57)
                                                            }
                                                            fmt.Println(v57, v1, v6)
                                                            if false {
                                                                v105 := []int{fi59}
                                                                if true {
                                                                    if len(v105) >= 2 {
                                                                        v105 = v105[1:len(v105)]
                                                                    }
                                                                }
                                                                _ = v105
                                                            } else {
                                                                v6 = (127 > 127)
                                                                v3[fi59] = (1 >= v88)
                                                            }
                                                            _ = v93
                                                        }
                                                        var v106 uint8 = 95
                                                        _ = v106
                                                    }
                                                    _ = v91
                                                    _ = v92
                                                } else {
                                                    switch (fi59 | fi59) {
                                                    case -9223372036854775808:
                                                        var v107 uint = 402
                                                        _ = v107
                                                    }
                                                    v108 := v4[42]
                                                    _ = v108
                                                }
                                                v57 = v1
                                                v5 += v5
                                            }
                                        }
                                    }
                                case "foo bar":
                                    v1 = (^(v1))
                                    v109 := GenS15{F0: -2147483648}
                                    _ = v109
                                }
                                _ = fi89
                            }
                            v110 := []int{fi59}
                            _ = v110
                            _ = v88
                        }
                    case -9223372036854775808:
                        for fi111 := 0; fi111 < 20; fi111++ {
                            v3[(^(fi59))] = v6
                            _ = fi111
                        }
                    }
                    var v112 string = "\n\t"
                    _ = v112
                } else {
                    fmt.Println(v5, fi59, v5)
                }
                var v113 uint8 = 194
                _ = v113
            }
        }
        for fi114 := 0; fi114 < 12; fi114++ {
            if (!(!(!(true || v7)))) {
                if false {
                    fmt.Println(v2, v2)
                }
                v5 = 1000
                if (-9223372036854775808 > 9223372036854775807) {
                    v115 := GenS16{F0: 42, F1: (v1 % ((v1) | 1))}
                    switch (-(fi114)) {
                    case -2147483648:
                        switch "a" {
                        case "":
                            var v116 uint16 = v57
                            _ = v116
                        case "foo bar":
                            var v117 uint64 = 17
                            _ = v117
                        }
                    case 0:
                        v3[(fi114 >> ((65535) & 63))] = false
                    case -9223372036854775808:
                        var v118 float64 = 0.1
                        v3[(fi114 >> ((65535) & 63))] = (!true)
                        _ = v118
                    default:
                        if (!v7) {
                            fmt.Println(v7, v7, v6)
                            delete(v4, 185)
                            var v119 uint16 = (v57 % ((243) | 1))
                            _ = v119
                        }
                    }
                    _ = v115
                } else {
                    _ = len(v4)
                    fmt.Println(fi114)
                    fmt.Println(fi59, v5, fi114)
                }
            } else {
                switch ("" + ("hello" + "a")) {
                case "a":
                    v57 += (v57 & (v1 * (^(v57))))
                case "foo bar":
                    switch fi114 {
                    case 1:
                        v6 = (((0 > 42) || true) && v7)
                    case -2147483648:
                        v3[(^(fi114))] = (104 > 100)
                    case 2147483647:
                        v120 := GenS1{F0: 255, F1: (^(v1))}
                        _ = v120
                    default:
                        v2 = v2
                    }
                default:
                    if (v6 && v6) {
                        if (v6 && v7) {
                            v121 := GenS13{F0: -128, F1: 18446744073709551615, F2: ((9223372036854775807 == -9223372036854775808) || (!((float32(-2.25) + 1e10) != v2)))}
                            switch (fi59 << ((2) & 63)) {
                            case 986:
                                var v122 float32 = 0.0
                                _ = v122
                            }
                            v123 := GenS5{F0: 42}
                            _ = v121
                            _ = v123
                        } else {
                            switch -9223372036854775808 {
                            case 128:
                                v124 := GenS17{F0: 53, F1: (fi59 ^ 4294967295), F2: fi59}
                                _ = v124
                            case 1:
                                v57 += v1
                            case 4294967295:
                                v125 := GenS16{F0: 285, F1: v1}
                                _ = v125
                            }
                            v126 := GenS8{F0: 127, F1: -414}
                            switch (fi59 * 2147483647) {
                            case -539:
                                if (!v7) {
                                    var v127 int32 = -592
                                    _ = v127
                                } else {
                                    if (0.0 >= float32(0.1)) {
                                        delete(v3, (fi114 >> ((4294967295) & 63)))
                                        if (128 < 1) {
                                            fmt.Println(fi59)
                                            var v128 uint = 4294967295
                                            v129 := GenS18{F0: "\n\t", F1: (fi59 + (fi59 / ((42) | 1)))}
                                            _ = v128
                                            _ = v129
                                        }
                                        if false {
                                            var v130 string = ("hello" + "")
                                            v131 := []int{fi59, 997, (-(fi114))}
                                            _ = v130
                                            _ = v131
                                        }
                                        if (false || (-2147483648 > 2147483647)) {
                                            v132 := []int{fi59}
                                            delete(v3, 331)
                                            _ = v132
                                        }
                                    }
                                }
                            case 127:
                                v1 = v1
                            }
                            delete(v4, 2147483647)
                            _ = v126
                        }
                    }
                    v57 += v57
                }
                if (!(1 <= 948)) {
                    v133 := GenS19{F0: 0, F1: 479, F2: 175}
                    v7 = v6
                    switch fi59 {
                    case -9223372036854775808:
                        v57 = 2
                    case 128:
                        var v134 int32 = 1
                        _ = v134
                    }
                    _ = v133
                } else {
                    v135 := GenS20{F0: fi114, F1: -128}
                    switch "foo bar" {
                    case "foo bar":
                        if (0.0 >= (float32((0.0 + 1.5)) - 1.5)) {
                            v1 = v1
                            var v136 string = "hello"
                            v1 = 2
                            _ = v136
                        } else {
                            v6 = (-1 <= 0)
                            v7 = true
                            var v137 bool = v7
                            if (float64(v2) > 1e10) {
                                var v138 int16 = v5
                                if (v137 && (!v7)) {
                                    fmt.Println(fi114, v138)
                                    v139 := []int{(fi114 | 0), fi59}
                                    _ = len(v139)
                                    _ = v139
                                } else {
                                    fmt.Println(v57, v5, v137)
                                    var v140 int16 = (v138 * (v138 + v138))
                                    var v141 uint = 255
                                    _ = v140
                                    _ = v141
                                }
                                v5 += (v5 * v5)
                                if v6 {
                                    delete(v3, fi59)
                                    if (42 != 1) {
                                        var v142 uint64 = 2
                                        switch (fi114 ^ (^(fi59))) {
                                        case 1:
                                            v6 = (1 != 42)
                                        }
                                        _ = v142
                                    }
                                    var v143 int32 = 42
                                    fmt.Println(v57, v2)
                                    var v144 uint = 18446744073709551615
                                    _ = v143
                                    _ = v144
                                } else {
                                    switch fi59 {
                                    case -9223372036854775808:
                                        var v145 float64 = 0.0
                                        _ = v145
                                    default:
                                        var v146 int8 = 100
                                        _ = v146
                                    }
                                    v138 += -493
                                }
                                _ = v138
                            }
                            _ = v137
                        }
                    case "hello":
                        var v147 uint8 = 128
                        _ = v147
                    case "a":
                        v148 := GenS21{F0: -1}
                        _ = v148
                    default:
                        v1 = (v57 - v57)
                        var v149 int32 = 499
                        _ = v149
                    }
                    if (((v57 + (^(v57))) >= (v57 ^ v57)) || (4294967295 <= 2147483648)) {
                        if (-1 == 25) {
                            v150 := GenS8{F0: 2, F1: -365}
                            if true {
                                if false {
                                    switch "\n\t" {
                                    case "foo bar":
                                        v151 := []int{}
                                        _ = v151
                                    case "hello":
                                        var v152 uint = 2
                                        if ((^(v5)) == (v5 + v5)) {
                                            var v153 int16 = v5
                                            var v154 int8 = -98
                                            if (-1 <= 9223372036854775807) {
                                                v2 += float32(float64(((float32(3.14159) * 0.0) + -2.25)))
                                                var v155 string = (("" + "\n\t") + "")
                                                var v156 float64 = ((float64(v2) * 1.5) - (float64((v2 - ((0.1 * (((1e10 / (v2 * 0.0)) * float32((float64(v2) * 0.0))) * (1.5 * 3.14159))) * ((v2 / v2) / 0.1)))) / 1.0))
                                                _ = v155
                                                _ = v156
                                            }
                                            _ = v153
                                            _ = v154
                                        }
                                        v152 = (v152 >> ((0) & 63))
                                        _ = v152
                                    case "":
                                        v7 = v7
                                    default:
                                        v157 := []int{(fi59 << ((974) & 63)), (fi114 * (fi59 + fi59)), fi59}
                                        _ = v157
                                    }
                                    var v158 uint32 = 42
                                    _ = v158
                                } else {
                                    var v159 float64 = 0.1
                                    switch (fi59 << ((42) & 63)) {
                                    case 128:
                                        switch fi114 {
                                        case -9223372036854775808:
                                            var v160 uint = 65535
                                            _ = v160
                                        }
                                    case 9223372036854775807:
                                        v161 := GenS7{F0: int32(v5)}
                                        _ = v161
                                    }
                                    _ = v159
                                }
                            }
                            _ = v150
                        }
                        v7 = v7
                        switch "" {
                        case "foo bar":
                            fmt.Println(v1)
                        }
                    }
                    _ = v135
                }
            }
            _ = fi114
        }
        _ = fi59
    }
    var v162 uint = 1
    if true {
        v5 += (^(v5))
        v57 = (v57 | v1)
    }
    v2 = ((v2 / v2) / float32(1e10))
    v163 := GenS22{F0: -2147483648, F1: uint8(v162)}
    var v164 uint16 = v57
    fmt.Println("v1:", v1)
    fmt.Println("v162:", v162)
    fmt.Println("v164:", v164)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v162
    _ = v163
    _ = v164
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v6
    _ = v7
}
