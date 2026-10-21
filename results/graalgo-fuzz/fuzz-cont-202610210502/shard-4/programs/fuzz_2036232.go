package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 int16; F1 uint8; F2 bool }
type GenS6 struct { F0 int16; F1 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int8; F1 uint }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS13 struct { F0 string; F1 bool }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int32; F1 int16; F2 int }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 int; F1 int16; F2 int32 }
func genFn0() bool {
    switch "hello" {
    case "foo bar":
        var v6 int8 = 1
        _ = v6
    default:
        if false {
            var v7 int16 = 562
            _ = v7
        }
    }
    v8 := []int{1}
    var v9 int32 = 0
    _ = v8
    _ = v9
    return true
}
func genFn1(p0 int32) int32 {
    switch "" {
    case "":
        p0 = -1
    case "foo bar":
        p0 -= -993
    case "hello":
        var v10 uint8 = 2
        _ = v10
    default:
        p0 += (^(p0))
    }
    fmt.Println(p0)
    v11 := GenS1{F0: 0}
    fmt.Println(p0)
    if false {
        for fi12 := 0; fi12 < 3; fi12++ {
            p0 = int32(fi12)
            v13 := []int{(fi12 * int(p0))}
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: -15}
        v15 := []int{0}
        _ = v14
        _ = v15
    } else {
        v16 := GenS1{F0: 118}
        v17 := GenS1{F0: -1}
        _ = v16
        _ = v17
    }
    _ = p0
    _ = v11
    return -2147483648
}

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    var v5 float64 = float64(((0.1 / 3.14159) + 0.0))
    v18 := GenS2{F0: -42, F1: 2147483647}
    for fi19 := 0; fi19 < 7; fi19++ {
        if false {
            v20 := v2[fi19]
            var v21 float32 = (3.14159 - 0.1)
            v20 += fi19
            _ = v20
            _ = v21
        }
        var v22 int = -799
        _ = fi19
        _ = v22
    }
    if ((1000 > -1) && ((32768 > 128) && true)) {
        var v23 float32 = (0.0 - (1e10 + (-2.25 - -2.25)))
        _ = v23
    }
    v5 = 1.5
    v5 += v5
    v5 = float64(-2.25)
    switch -9223372036854775808 {
    case 366:
        for fi24 := 0; fi24 < 1; fi24++ {
            switch fi24 {
            case -9223372036854775808:
                var v25 int32 = -1
                _ = v25
            case 0:
                var v26 uint8 = 176
                _ = v26
            }
            _ = fi24
        }
    case -9223372036854775808:
        var v27 int16 = -32768
        _ = v27
    case 0:
        var v28 uint64 = 18446744073709551615
        _ = v28
    default:
        fmt.Println(v5)
    }
    v5 = -2.25
    v5 = v5
    v29 := GenS1{F0: 127}
    switch ("\n\t" + ("hello" + "\n\t")) {
    case "hello":
        v2[32768] = 127
    case "foo bar":
        var v30 int = 0
        _ = v30
    case "\n\t":
        var v31 uint = 65535
        _ = v31
    default:
        if true {
            for fi32 := 0; fi32 < 10; fi32++ {
                switch 9223372036854775807 {
                case 1:
                    if false {
                        if (!false) {
                            v33 := GenS3{F0: -1, F1: 0}
                            for fi34 := 0; fi34 < 6; fi34++ {
                                var v35 int16 = 32767
                                v5 += 3.14159
                                v5 -= 0.1
                                var v36 int64 = 467
                                v37 := GenS2{F0: 28, F1: 2147483647}
                                var v38 uint8 = 109
                                _ = fi34
                                _ = v35
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                            _ = v33
                        }
                    } else {
                        for fi39 := 0; fi39 < 7; fi39++ {
                            switch fi32 {
                            case -9223372036854775808:
                                v40 := GenS1{F0: 127}
                                _ = v40
                            case 200:
                                v5 = 3.14159
                            case 4294967295:
                                if (42 >= 9223372036854775807) {
                                    v5 -= (v5 / 1e10)
                                    if true {
                                        switch (^(fi39)) {
                                        case 17:
                                            v5 -= 1e10
                                        case 807:
                                            var v41 int8 = 14
                                            _ = v41
                                        }
                                        if (fi32 <= fi39) {
                                            var v42 uint16 = 769
                                            v43 := GenS3{F0: -1, F1: 4294967295}
                                            fmt.Println(fi39, v5, v42)
                                            _ = v42
                                            _ = v43
                                        } else {
                                            if (!((float32(float64(0.1)) != (0.0 * 1.5)) || true)) {
                                                fmt.Println(fi39, v5)
                                                var v44 float64 = (v5 * (float64(0.0) / ((0.1 + 1e10) / (0.0 - v5))))
                                                v45 := GenS3{F0: 0, F1: 42}
                                                switch "a" {
                                                case "hello":
                                                    switch "a" {
                                                    case "foo bar":
                                                        if true {
                                                            if (genFn1(1) < genFn1(-186)) {
                                                                v5 -= 1e10
                                                            }
                                                            v5 -= ((0.1 - v44) - 1.5)
                                                            if (4294967295 != 42) {
                                                                v5 = v44
                                                                v44 = 0.1
                                                            }
                                                        }
                                                    case "a":
                                                        v44 -= (v44 + -2.25)
                                                    case "":
                                                        if (!(0 != 4294967295)) {
                                                            _ = len(v2)
                                                            switch 32767 {
                                                            case 0:
                                                                if true {
                                                                    var v46 bool = (true && (18446744073709551615 > 1))
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 765
                                                                    }
                                                                    fmt.Println(v44, v5, v44)
                                                                    v47 := GenS2{F0: -1, F1: 0}
                                                                    if true {
                                                                        var v48 int = fi39
                                                                        v48 -= 4294967295
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        _ = v48
                                                                    }
                                                                    _ = v46
                                                                    _ = v47
                                                                }
                                                            case 982:
                                                                fmt.Println(fi32)
                                                                switch "a" {
                                                                case "\n\t":
                                                                    fmt.Println(v44, fi39, fi39)
                                                                case "foo bar":
                                                                    var v49 uint32 = 753
                                                                    _ = v49
                                                                case "hello":
                                                                    v5 += v5
                                                                    v50 := GenS2{F0: 127, F1: -1}
                                                                    _ = v50
                                                                }
                                                            case 4294967295:
                                                                if (!(!true)) {
                                                                    var v51 int32 = genFn1(genFn1(898))
                                                                    v52 := GenS4{F0: fi39, F1: 1}
                                                                    _ = v51
                                                                    _ = v52
                                                                } else {
                                                                    var v53 int8 = 127
                                                                    if (2 >= 0) {
                                                                        switch (fi39 % (((-(fi39))) | 1)) {
                                                                        case 32767:
                                                                            v44 += float64(1e10)
                                                                        default:
                                                                            v5 = v5
                                                                        }
                                                                    } else {
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        v5 = (1.5 - 1.5)
                                                                        _ = len(v3)
                                                                    }
                                                                    _ = v53
                                                                }
                                                            }
                                                            var v54 float32 = 0.1
                                                            _ = v54
                                                        } else {
                                                            v44 -= v44
                                                        }
                                                    }
                                                default:
                                                    switch -516 {
                                                    case -1:
                                                        fmt.Println(v5, v44)
                                                    case 0:
                                                        switch ("" + "hello") {
                                                        case "\n\t":
                                                            v55 := GenS5{F0: -636, F1: 0, F2: (2 > 1)}
                                                            _ = v55
                                                        case "a":
                                                            if (true || ("" == "\n\t")) {
                                                                _ = len(v3)
                                                                var v56 int = 1
                                                                _ = v56
                                                            }
                                                        default:
                                                            var v57 int32 = 42
                                                            _ = v57
                                                        }
                                                    default:
                                                        var v58 string = ""
                                                        _ = v58
                                                    }
                                                }
                                                _ = v44
                                                _ = v45
                                            } else {
                                                if (1 <= -128) {
                                                    var v59 float32 = 0.0
                                                    _ = v59
                                                } else {
                                                    if true {
                                                        v5 += 0.0
                                                        v5 += v5
                                                        v5 -= v5
                                                        var v60 uint = 597
                                                        _ = v60
                                                    } else {
                                                        v61 := v4["hello"]
                                                        _ = v61
                                                    }
                                                }
                                                v62 := GenS3{F0: 150, F1: 817}
                                                v63 := v3[(fi32 | fi32)]
                                                _ = v62
                                                _ = v63
                                            }
                                        }
                                        v64 := GenS5{F0: -732, F1: 42, F2: true}
                                        _ = v64
                                    }
                                } else {
                                    v65 := GenS1{F0: 0}
                                    var v66 int = (fi32 % (((fi32 + 4294967295)) | 1))
                                    fmt.Println(fi32, v5)
                                    _ = v65
                                    _ = v66
                                }
                            }
                            var v67 uint64 = 18446744073709551615
                            delete(v4, ("\n\t" + "hello"))
                            v68 := GenS6{F0: 42, F1: true}
                            _ = fi39
                            _ = v67
                            _ = v68
                        }
                    }
                case -9223372036854775808:
                    switch ("a" + "a") {
                    case "":
                        v69 := GenS4{F0: fi32, F1: 0}
                        _ = v69
                    case "hello":
                        var v70 int8 = 0
                        _ = v70
                    case "foo bar":
                        v5 = v5
                    default:
                        v1 = append(v1, 2147483647)
                    }
                case 128:
                    switch (fi32 % ((fi32) | 1)) {
                    case 4294967295:
                        switch "hello" {
                        case "a":
                            for fi71 := 0; fi71 < 3; fi71++ {
                                var v72 float64 = float64((float32(0.1) / (((float32(v5) * 0.0) - (1.5 / 1e10)) / ((((0.1 / 1.5) + 0.0) / 0.1) * ((((((1e10 - 1e10) - 3.14159) - 0.0) + 1.5) / 3.14159) / (1e10 - (1.5 * (3.14159 / (0.1 / -2.25)))))))))
                                var v73 int8 = 42
                                _ = fi71
                                _ = v72
                                _ = v73
                            }
                        default:
                            for fi74 := 0; fi74 < 1; fi74++ {
                                _ = len(v2)
                                v75 := GenS5{F0: -1, F1: 0, F2: true}
                                _ = fi74
                                _ = v75
                            }
                        }
                    case 572:
                        fmt.Println(v5)
                    case -2147483648:
                        fmt.Println(fi32, fi32)
                    default:
                        if false {
                            fmt.Println(v5, v5)
                        } else {
                            if (127 < 255) {
                                if (true && (true && (false && (0 == 9223372036854775807)))) {
                                    v5 += (v5 * v5)
                                    var v76 float32 = 0.0
                                    v5 = ((1.5 / 1.5) - v5)
                                    _ = v76
                                } else {
                                    if true {
                                        delete(v4, "foo bar")
                                        if (true || true) {
                                            _ = len(v1)
                                            if (!("\n\t" <= ("a" + "a"))) {
                                                var v77 int64 = 42
                                                var v78 string = (("foo bar" + "a") + "foo bar")
                                                _ = v77
                                                _ = v78
                                            } else {
                                                v79 := GenS7{F0: -128}
                                                var v80 string = (("a" + "a") + "a")
                                                if (!(!((!true) && true))) {
                                                    v4[(v80 + ("\n\t" + v80))] = -919
                                                    var v81 uint = 691
                                                    fmt.Println(fi32, v5, v80)
                                                    _ = v81
                                                }
                                                _ = v79
                                                _ = v80
                                            }
                                            var v82 uint32 = 929
                                            _ = v82
                                        } else {
                                            var v83 string = "\n\t"
                                            if false {
                                                for fi84 := 0; fi84 < 3; fi84++ {
                                                    v83 = (v83 + v83)
                                                    _ = fi84
                                                }
                                                for fi85 := 0; fi85 < 9; fi85++ {
                                                    v5 += v5
                                                    _ = fi85
                                                }
                                            } else {
                                                var v86 int16 = 42
                                                var v87 bool = (-606 != 0)
                                                _ = v86
                                                _ = v87
                                            }
                                            if (!true) {
                                                var v88 float32 = -2.25
                                                for fi89 := 0; fi89 < 4; fi89++ {
                                                    v90 := GenS1{F0: -128}
                                                    switch 2147483647 {
                                                    case 4294967295:
                                                        v5 += v5
                                                    case 328:
                                                        var v91 uint = 0
                                                        v91 += (v91 / (((^(v91))) | 1))
                                                        switch v83 {
                                                        case "\n\t":
                                                            switch (fi89 | (^(fi32))) {
                                                            case 127:
                                                                if false {
                                                                    switch "hello" {
                                                                    case "":
                                                                        v92 := GenS5{F0: -1, F1: 2, F2: false}
                                                                        _ = v92
                                                                    case "a":
                                                                        switch fi89 {
                                                                        case 4294967295:
                                                                            var v93 uint = (v91 / ((v91) | 1))
                                                                            _ = v93
                                                                        case -243:
                                                                            fmt.Println(fi32, v91)
                                                                        case 42:
                                                                            var v94 bool = true
                                                                            _ = v94
                                                                        default:
                                                                            var v95 int = -9223372036854775808
                                                                            _ = v95
                                                                        }
                                                                    }
                                                                } else {
                                                                    _ = len(v1)
                                                                }
                                                            }
                                                        case "":
                                                            var v96 uint8 = 128
                                                            if true {
                                                                v97 := GenS8{F0: 42, F1: 255}
                                                                if true {
                                                                    v83 = "hello"
                                                                } else {
                                                                    switch fi89 {
                                                                    case -173:
                                                                        v5 += v5
                                                                    case 32768:
                                                                        _ = len(v2)
                                                                    }
                                                                    v98 := GenS9{F0: fi32}
                                                                    v91 += v91
                                                                    _ = len(v1)
                                                                    if false {
                                                                        v1 = append(v1, 160)
                                                                    }
                                                                    _ = v98
                                                                }
                                                                v96 = v96
                                                                _ = v97
                                                            }
                                                            v88 -= v88
                                                            _ = v96
                                                        default:
                                                            var v99 string = v83
                                                            _ = v99
                                                        }
                                                        _ = v91
                                                    case 633:
                                                        switch (-(fi32)) {
                                                        case 596:
                                                            var v100 int32 = genFn1(0)
                                                            _ = v100
                                                        default:
                                                            var v101 float32 = ((-2.25 / (((float32(float64(v88)) * v88) * (0.1 * (((v88 - (v88 + v88)) * v88) / (v88 * (float32(0.1) / v88))))) * (v88 / v88))) / float32(float64(1.5)))
                                                            _ = v101
                                                        }
                                                    default:
                                                        var v102 int16 = -1
                                                        _ = v102
                                                    }
                                                    _ = fi89
                                                    _ = v90
                                                }
                                                _ = v88
                                            } else {
                                                v5 -= v5
                                                v5 += v5
                                            }
                                            _ = v83
                                        }
                                        v103 := GenS10{F0: 4294967295, F1: -1, F2: 127}
                                        _ = v103
                                    } else {
                                        switch 928 {
                                        case -1:
                                            switch ("" + "foo bar") {
                                            case "":
                                                fmt.Println(v5)
                                            case "foo bar":
                                                var v104 string = "hello"
                                                _ = v104
                                            }
                                        case 42:
                                            delete(v2, (^(fi32)))
                                        case -523:
                                            v105 := GenS4{F0: (fi32 >> ((0) & 63)), F1: 9223372036854775807}
                                            _ = v105
                                        }
                                        v106 := GenS11{F0: 32767}
                                        _ = v106
                                    }
                                    _ = len(v1)
                                }
                            } else {
                                switch fi32 {
                                case 1:
                                    if true {
                                        if ((false || false) || false) {
                                            var v107 int = fi32
                                            _ = v107
                                        }
                                        if false {
                                            if false {
                                                if true {
                                                    var v108 int16 = 0
                                                    for fi109 := 0; fi109 < 9; fi109++ {
                                                        var v110 bool = (-1 < 9223372036854775807)
                                                        _ = fi109
                                                        _ = v110
                                                    }
                                                    _ = v108
                                                }
                                                var v111 uint64 = 1
                                                _ = v111
                                            }
                                        } else {
                                            v5 += v5
                                            v5 = 0.1
                                        }
                                        if (42 > 745) {
                                            var v112 float32 = -2.25
                                            _ = v112
                                        }
                                    } else {
                                        v5 += (((float64(1e10) + ((1e10 / float64(1.5)) + 1.5)) + v5) - (0.0 * v5))
                                        var v113 bool = (2147483647 != 0)
                                        var v114 int64 = 707
                                        _ = v113
                                        _ = v114
                                    }
                                case 0:
                                    v5 = v5
                                case -89:
                                    if true {
                                        v1 = append(v1, 42)
                                        switch (fi32 & (fi32 * fi32)) {
                                        case 9223372036854775807:
                                            _ = len(v4)
                                        case 127:
                                            v5 -= v5
                                        }
                                        var v115 uint32 = 42
                                        _ = v115
                                    }
                                }
                                for fi116 := 0; fi116 < 20; fi116++ {
                                    v117 := GenS3{F0: -399, F1: 1}
                                    delete(v4, ("a" + ("\n\t" + (("\n\t" + "") + ("\n\t" + "a")))))
                                    if true {
                                        v5 = v5
                                        switch fi116 {
                                        case 32768:
                                            if (32767 != 32767) {
                                                if false {
                                                    var v118 int32 = genFn1(0)
                                                    _ = v118
                                                }
                                                v119 := GenS12{F0: 9223372036854775807, F1: 65535, F2: 42}
                                                _ = v119
                                            }
                                        case -263:
                                            v5 = v5
                                        }
                                        v5 += (v5 * 3.14159)
                                        v1 = append(v1, 463)
                                        if true {
                                            var v120 bool = (false || (481 > 4294967295))
                                            switch -983 {
                                            case -9223372036854775808:
                                                switch "\n\t" {
                                                case "":
                                                    var v121 int32 = -335
                                                    _ = v121
                                                }
                                            case 32767:
                                                switch -929 {
                                                case 32767:
                                                    v122 := GenS13{F0: ("hello" + "\n\t"), F1: v120}
                                                    _ = v122
                                                case 9223372036854775807:
                                                    var v123 bool = (v120 || (false && (!v120)))
                                                    _ = v123
                                                }
                                            case -370:
                                                if (443 != 649) {
                                                    if v120 {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v4["hello"] = (fi32 >> ((1) & 63))
                                                        default:
                                                            var v124 bool = false
                                                            _ = v124
                                                        }
                                                    }
                                                    v125 := GenS14{F0: 1}
                                                    v120 = false
                                                    _ = v125
                                                }
                                            default:
                                                var v126 uint8 = 1
                                                _ = v126
                                            }
                                            v127 := GenS4{F0: (fi32 % ((42) | 1)), F1: -619}
                                            _ = v120
                                            _ = v127
                                        } else {
                                            if true {
                                                v5 += (v5 - v5)
                                                if (false || true) {
                                                    var v128 int = (-(fi116))
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v5 -= -2.25
                                                    case "foo bar":
                                                        v129 := GenS7{F0: 42}
                                                        _ = v129
                                                    default:
                                                        if true {
                                                            v128 = 32768
                                                            var v130 uint64 = 718
                                                            v128 = 0
                                                            v5 = ((((float64((1e10 * 3.14159)) + v5) * v5) * 1.5) + v5)
                                                            var v131 uint64 = v130
                                                            if (false || false) {
                                                                var v132 float32 = 1e10
                                                                if (-1 != -128) {
                                                                    var v133 float32 = v132
                                                                    var v134 int = fi116
                                                                    if false {
                                                                        fmt.Println(v131, v5, v132)
                                                                        v135 := GenS15{F0: 1, F1: 169, F2: 0}
                                                                        if true {
                                                                            v136 := GenS6{F0: 551, F1: ((v5 * -2.25) != v5)}
                                                                            switch (v128 | fi32) {
                                                                            case 32767:
                                                                                var v137 int64 = 581
                                                                                if (1000 <= 99) {
                                                                                    v138 := GenS6{F0: -1, F1: ((v131 == (v130 / ((v131) | 1))) && (-117 > 127))}
                                                                                    var v139 int8 = -19
                                                                                    _ = v138
                                                                                    _ = v139
                                                                                } else {
                                                                                    if (!true) {
                                                                                        delete(v4, "")
                                                                                        var v140 uint16 = 100
                                                                                        delete(v3, (v134 & int(v131)))
                                                                                        var v141 float32 = (v133 / (0.1 * v132))
                                                                                        switch "foo bar" {
                                                                                        case "":
                                                                                            v142 := GenS1{F0: 127}
                                                                                            _ = v142
                                                                                        case "foo bar":
                                                                                            fmt.Println(v134)
                                                                                        case "\n\t":
                                                                                            v133 += 1e10
                                                                                        }
                                                                                        _ = v140
                                                                                        _ = v141
                                                                                    } else {
                                                                                        _ = len(v4)
                                                                                        v143 := GenS16{F0: 9223372036854775807}
                                                                                        v5 += (v5 + (((((v5 * (v5 / v5)) * (0.0 / 1.5)) * v5) - 0.1) / (1.5 / (v5 / ((0.0 * -2.25) * float64(v132))))))
                                                                                        if false {
                                                                                            v130 -= (v130 / ((4294967295) | 1))
                                                                                            var v144 uint8 = 255
                                                                                            v130 = v130
                                                                                            _ = v144
                                                                                        }
                                                                                        _ = v143
                                                                                    }
                                                                                    switch ("\n\t" + ("" + "\n\t")) {
                                                                                    case "hello":
                                                                                        v134 -= 42
                                                                                    case "":
                                                                                        switch fi116 {
                                                                                        case 844:
                                                                                            _ = len(v4)
                                                                                        default:
                                                                                            v145 := GenS11{F0: 231}
                                                                                            _ = v145
                                                                                        }
                                                                                    case "foo bar":
                                                                                        v146 := GenS17{F0: v131}
                                                                                        _ = v146
                                                                                    default:
                                                                                        v147 := v3[(fi116 / ((0) | 1))]
                                                                                        v131 += 1
                                                                                        fmt.Println(fi116, v137, fi116)
                                                                                        var v148 uint64 = (v130 % (((v131 >> ((18446744073709551615) & 63))) | 1))
                                                                                        _ = v147
                                                                                        _ = v148
                                                                                    }
                                                                                    v3[-2147483648] = (0.1 >= v5)
                                                                                }
                                                                                _ = v137
                                                                            }
                                                                            _ = v136
                                                                        }
                                                                        _ = v135
                                                                    }
                                                                    _ = v133
                                                                    _ = v134
                                                                } else {
                                                                    v5 = (1.5 - v5)
                                                                }
                                                                _ = v132
                                                            }
                                                            switch v128 {
                                                            case -1:
                                                                var v149 string = ""
                                                                _ = v149
                                                            case 720:
                                                                if false {
                                                                    if true {
                                                                        v5 += (((((((float64((-2.25 / 3.14159)) / float64((0.0 - 1e10))) / v5) / (((v5 + -2.25) * (v5 * 3.14159)) * (v5 + (1.5 * v5)))) / v5) / 1e10) - v5) / float64(float32(-2.25)))
                                                                        var v150 int = fi116
                                                                        _ = v150
                                                                    } else {
                                                                        _ = len(v4)
                                                                    }
                                                                    _ = len(v1)
                                                                } else {
                                                                    var v151 float64 = v5
                                                                    switch 2147483647 {
                                                                    case 849:
                                                                        var v152 uint16 = 185
                                                                        _ = v152
                                                                    case 32767:
                                                                        v131 = (v130 / ((v131) | 1))
                                                                    }
                                                                    v5 -= v5
                                                                    _ = v151
                                                                }
                                                            }
                                                            var v153 float64 = v5
                                                            v131 = (^(v131))
                                                            if len(v1) > 0 {
                                                                v1[0] = 2147483647
                                                            }
                                                            _ = v130
                                                            _ = v131
                                                            _ = v153
                                                        } else {
                                                            v128 -= 42
                                                            if false {
                                                                v154 := GenS12{F0: 1, F1: 2, F2: 42}
                                                                _ = v154
                                                            }
                                                        }
                                                    }
                                                    _ = v128
                                                } else {
                                                    if false {
                                                        v5 += 0.0
                                                        switch (-(fi116)) {
                                                        case 868:
                                                            if true {
                                                                switch (fi32 * (fi116 | (fi32 ^ fi116))) {
                                                                case 736:
                                                                    _ = len(v1)
                                                                default:
                                                                    if false {
                                                                        v5 += v5
                                                                        _ = len(v1)
                                                                    }
                                                                }
                                                                v5 -= v5
                                                            }
                                                        case 1:
                                                            v1 = append(v1, 540)
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = 452
                                                        }
                                                    } else {
                                                        if (255 <= 312) {
                                                            var v155 uint8 = 0
                                                            if (!false) {
                                                                v155 = v155
                                                                if true {
                                                                    v5 -= 3.14159
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                } else {
                                                                    v156 := GenS10{F0: 4294967295, F1: -1, F2: (v155 & (v155 ^ 1))}
                                                                    v5 += (1e10 + v5)
                                                                    _ = v156
                                                                }
                                                            }
                                                            var v157 float64 = 1e10
                                                            _ = v155
                                                            _ = v157
                                                        }
                                                    }
                                                    fmt.Println(fi116, fi32, v5)
                                                    fmt.Println(fi32, v5)
                                                    fmt.Println(fi32)
                                                }
                                            }
                                        }
                                    } else {
                                        var v158 int8 = 42
                                        if len(v1) > 0 {
                                            v1[0] = 576
                                        }
                                        switch fi32 {
                                        case 4294967295:
                                            if len(v1) > 0 {
                                                v1[0] = 1
                                            }
                                            v5 = v5
                                            v159 := GenS18{F0: fi32, F1: 42, F2: -387}
                                            _ = v159
                                        case 128:
                                            v5 = v5
                                        }
                                        _ = v158
                                    }
                                    _ = fi116
                                    _ = v117
                                }
                            }
                        }
                    }
                }
                v5 = float64((1e10 * (float32((((float64(0.0) * v5) / ((v5 + float64(0.0)) * 0.0)) + 0.0)) + ((0.0 + ((float32(float64(((1.5 * 1.5) + 0.0))) * (1e10 - 1.5)) * -2.25)) * (1e10 * (float32(float64(3.14159)) - 1e10))))))
                switch -173 {
                case 554:
                    for fi160 := 0; fi160 < 16; fi160++ {
                        if len(v1) > 0 {
                            v1[0] = 72
                        }
                        v5 -= v5
                        fmt.Println(fi160, fi32)
                        v5 -= v5
                        _ = fi160
                    }
                case 42:
                    var v161 uint64 = 0
                    _ = v161
                case -9223372036854775808:
                    if false {
                        for fi162 := 0; fi162 < 10; fi162++ {
                            v1 = append(v1, 562)
                            v5 = (0.0 * v5)
                            var v163 float64 = (0.1 - v5)
                            _ = fi162
                            _ = v163
                        }
                        if (2147483648 != 418) {
                            switch fi32 {
                            case 1:
                                fmt.Println(fi32)
                            case 127:
                                var v164 int = (fi32 + -1)
                                _ = v164
                            case 0:
                                v5 = v5
                            default:
                                v5 = v5
                            }
                        }
                        var v165 int = fi32
                        _ = v165
                    }
                }
                _ = fi32
            }
            var v166 int8 = 42
            _ = v166
        } else {
            v5 = 1e10
            var v167 int = 1
            v167 += v167
            _ = v167
        }
    }
    var v168 int = 4294967295
    var v169 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v168:", v168)
    fmt.Println("v169:", v169)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v168
    _ = v169
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
