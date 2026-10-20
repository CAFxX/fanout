package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 bool; F2 uint32 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 uint8; F1 int16 }
type GenS9 struct { F0 uint32; F1 uint64 }
type GenS10 struct { F0 int8; F1 int64; F2 string }
func genFn0() uint {
    var v6 uint16 = 678
    var v7 uint64 = 441
    _ = v6
    _ = v7
    return 221
}

func main() {
    var v1 float32 = 0.0
    var v2 float64 = 1e10
    v3 := make(map[string]uint32)
    var v4 float64 = v2
    var v5 int64 = 1
    var v8 int64 = v5
    if true {
        switch 128 {
        case 128:
            var v9 int64 = v5
            _ = v9
        case 42:
            v5 = 9223372036854775807
        case 539:
            v10 := GenS1{F0: 1, F1: 18446744073709551615}
            _ = v10
        }
        switch (("" + "hello") + ((((("hello" + "a") + "") + "\n\t") + "\n\t") + "\n\t")) {
        case "\n\t":
            v11 := GenS1{F0: 349, F1: 4294967295}
            _ = v11
        case "a":
            if true {
                v12 := GenS1{F0: 833, F1: 2}
                delete(v3, "\n\t")
                _ = v12
            } else {
                if (42 < 0) {
                    if true {
                        v4 = 0.1
                        delete(v3, ("" + ("hello" + "\n\t")))
                        v13 := GenS1{F0: 0, F1: 42}
                        _ = v13
                    } else {
                        var v14 uint64 = 2
                        if ((-(v5)) < -1) {
                            v5 = v5
                        } else {
                            v15 := GenS2{F0: 127}
                            v4 = (v4 + (v2 / ((v2 + (v2 / (float64((0.0 * 3.14159)) - v2))) - v4)))
                            _ = v15
                        }
                        _ = v14
                    }
                    fmt.Println(v8, v4)
                    var v16 bool = false
                    _ = v16
                }
                var v17 int16 = -32768
                fmt.Println(v17, v5)
                _ = v17
            }
        case "":
            v2 += v4
        default:
            if (246 <= 781) {
                var v18 uint32 = 0
                _ = v18
            } else {
                if (18446744073709551615 <= 1) {
                    v19 := GenS2{F0: 9223372036854775807}
                    _ = v19
                }
            }
        }
    }
    v20 := GenS1{F0: 42, F1: 0}
    v2 = v4
    fmt.Println(v8, v1, v8)
    var v21 uint16 = 1
    v22 := GenS2{F0: int(v21)}
    for fi23 := 0; fi23 < 14; fi23++ {
        v24 := GenS1{F0: 42, F1: 42}
        var v25 string = "hello"
        switch v25 {
        case "a":
            var v26 uint8 = 0
            _ = v26
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    var v27 bool = (0 < 2)
    fmt.Println(v8, v21)
    if true {
        var v28 uint16 = (v21 % ((v21) | 1))
        v29 := GenS2{F0: 32767}
        _ = v28
        _ = v29
    } else {
        v30 := GenS1{F0: 1, F1: 540}
        _ = v30
    }
    if v27 {
        if ((v27 && (!((-1 == 0) && (1000 >= int16(v21))))) || (!true)) {
            for fi31 := 0; fi31 < 3; fi31++ {
                delete(v3, ("foo bar" + "hello"))
                v32 := GenS2{F0: (fi31 % (((^(fi31))) | 1))}
                if (42 <= 1) {
                    var v33 float64 = (v2 - 1e10)
                    for fi34 := 0; fi34 < 8; fi34++ {
                        delete(v3, "a")
                        var v35 uint = 630
                        _ = fi34
                        _ = v35
                    }
                    _ = v33
                } else {
                    switch fi31 {
                    case 363:
                        var v36 float64 = 3.14159
                        _ = v36
                    case 310:
                        if v27 {
                            if true {
                                switch "hello" {
                                case "hello":
                                    fmt.Println(v8, v27, v27)
                                case "\n\t":
                                    var v37 string = "a"
                                    _ = v37
                                default:
                                    var v38 uint16 = (v21 >> ((genFn0()) & 63))
                                    _ = v38
                                }
                                switch -177 {
                                case 42:
                                    v1 = v1
                                default:
                                    for fi39 := 0; fi39 < 16; fi39++ {
                                        if true {
                                            v4 = v2
                                        } else {
                                            if ((103 != -21) || (!true)) {
                                                switch (fi39 / ((2147483647) | 1)) {
                                                case -863:
                                                    if v27 {
                                                        fmt.Println(v4, v27, v4)
                                                        v40 := GenS3{F0: v8, F1: v27, F2: 2}
                                                        v27 = (!v27)
                                                        _ = v40
                                                    } else {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            fmt.Println(v27, v4, fi39)
                                                        case "\n\t":
                                                            var v41 float64 = ((v2 / ((v4 + v4) * (v4 - 3.14159))) / v4)
                                                            _ = v41
                                                        case "hello":
                                                            v5 = (v8 * v5)
                                                        default:
                                                            v2 -= v4
                                                        }
                                                        fmt.Println(fi31, v8, v27)
                                                        if false {
                                                            switch (((("a" + "foo bar") + ("" + "\n\t")) + ((((("foo bar" + "foo bar") + "\n\t") + "a") + "") + "a")) + "\n\t") {
                                                            case "\n\t":
                                                                v1 -= (0.1 * v1)
                                                            case "hello":
                                                                if false {
                                                                    switch 32768 {
                                                                    case -2147483648:
                                                                        v42 := []int{(fi31 - 714)}
                                                                        _ = v42
                                                                    }
                                                                    if v27 {
                                                                        switch ("" + "a") {
                                                                        case "":
                                                                            var v43 float64 = v4
                                                                            _ = v43
                                                                        case "a":
                                                                            v44 := GenS4{F0: 42, F1: 66}
                                                                            _ = v44
                                                                        case "hello":
                                                                            var v45 float32 = v1
                                                                            _ = v45
                                                                        default:
                                                                            var v46 uint32 = 2147483648
                                                                            _ = v46
                                                                        }
                                                                        if ((!true) || v27) {
                                                                            if v27 {
                                                                                _ = len(v3)
                                                                            } else {
                                                                                switch (("a" + (("\n\t" + "a") + ("hello" + ("" + ("hello" + "\n\t"))))) + "") {
                                                                                case "foo bar":
                                                                                    switch (("\n\t" + (("foo bar" + "foo bar") + "")) + "hello") {
                                                                                    case "a":
                                                                                        v1 = (float32(v4) - v1)
                                                                                    case "foo bar":
                                                                                        delete(v3, "hello")
                                                                                    }
                                                                                case "\n\t":
                                                                                    v47 := []int{-857, (-(fi39)), fi31}
                                                                                    _ = v47
                                                                                }
                                                                                _ = len(v3)
                                                                            }
                                                                            var v48 int64 = v5
                                                                            _ = v48
                                                                        } else {
                                                                            switch fi31 {
                                                                            case 42:
                                                                                v49 := []int{0, -9223372036854775808}
                                                                                _ = v49
                                                                            case 0:
                                                                                var v50 uint32 = 42
                                                                                _ = v50
                                                                            }
                                                                        }
                                                                        if true {
                                                                            v3["foo bar"] = 42
                                                                        }
                                                                    }
                                                                    delete(v3, "")
                                                                } else {
                                                                    if v27 {
                                                                        fmt.Println(fi39, v8)
                                                                        v51 := GenS5{F0: fi39}
                                                                        _ = len(v3)
                                                                        _ = v51
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v21, v2)
                                                            }
                                                            v27 = v27
                                                        }
                                                    }
                                                case 32768:
                                                    var v52 uint64 = 440
                                                    _ = v52
                                                case 979:
                                                    fmt.Println(v8, fi39)
                                                }
                                                switch ((("a" + "a") + (("" + "a") + "")) + "hello") {
                                                case "foo bar":
                                                    switch fi31 {
                                                    case 0:
                                                        v2 = 1e10
                                                        v53 := v3[("\n\t" + "\n\t")]
                                                        if true {
                                                            var v54 int16 = 0
                                                            _ = v54
                                                        } else {
                                                            var v55 uint16 = 811
                                                            _ = v55
                                                        }
                                                        _ = v53
                                                    case 2147483647:
                                                        var v56 bool = v27
                                                        _ = v56
                                                    case -9223372036854775808:
                                                        switch 1 {
                                                        case -9223372036854775808:
                                                            var v57 int8 = 1
                                                            _ = v57
                                                        }
                                                    }
                                                case "hello":
                                                    var v58 int = -177
                                                    v59 := GenS5{F0: (-(fi31))}
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v60 float32 = (1.5 * v1)
                                                        _ = v60
                                                    case "a":
                                                        v21 = v21
                                                    case "hello":
                                                        v8 = (v5 & -1)
                                                    }
                                                    _ = v58
                                                    _ = v59
                                                case "\n\t":
                                                    if v27 {
                                                        if v27 {
                                                            v61 := []int{(fi39 % ((fi39) | 1)), fi31, (fi39 & int(v21)), (fi31 / ((32767) | 1))}
                                                            _ = v61
                                                        } else {
                                                            var v62 float32 = v1
                                                            _ = v62
                                                        }
                                                        v63 := v3[("\n\t" + "foo bar")]
                                                        _ = v63
                                                    } else {
                                                        switch fi39 {
                                                        case 32768:
                                                            var v64 bool = (!(2 != 189))
                                                            _ = v64
                                                        case -345:
                                                            v8 -= v5
                                                        case -424:
                                                            var v65 uint8 = 2
                                                            _ = v65
                                                        default:
                                                            switch fi39 {
                                                            case 2147483647:
                                                                if v27 {
                                                                    v3["hello"] = 1
                                                                    v66 := GenS1{F0: 18446744073709551615, F1: 783}
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        if (-7 == 92) {
                                                                            v67 := GenS6{F0: false}
                                                                            _ = v67
                                                                        }
                                                                    }
                                                                    _ = v66
                                                                }
                                                            case 1:
                                                                v1 -= v1
                                                            case -566:
                                                                delete(v3, "\n\t")
                                                            default:
                                                                v3[("a" + "a")] = 4294967295
                                                            }
                                                        }
                                                        v68 := GenS2{F0: (fi31 + fi31)}
                                                        switch "hello" {
                                                        case "hello":
                                                            v69 := GenS7{F0: (("hello" + "a") + "\n\t"), F1: 65535}
                                                            _ = v69
                                                        case "\n\t":
                                                            v70 := GenS1{F0: uint64(v21), F1: uint64(v21)}
                                                            _ = v70
                                                        case "foo bar":
                                                            fmt.Println(v1, fi31)
                                                        }
                                                        _ = v68
                                                    }
                                                default:
                                                    if false {
                                                        delete(v3, "")
                                                        v71 := v3["a"]
                                                        _ = v71
                                                    } else {
                                                        if ("a" <= "") {
                                                            if true {
                                                                v3["hello"] = 282
                                                            }
                                                            var v72 uint8 = 2
                                                            v73 := []int{}
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                        var v74 uint16 = v21
                                                        v2 -= v4
                                                        var v75 float64 = v4
                                                        var v76 uint8 = 0
                                                        _ = v74
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                }
                                            }
                                            var v77 string = ("hello" + "")
                                            if v27 {
                                                switch "foo bar" {
                                                case "a":
                                                    delete(v3, "a")
                                                case "\n\t":
                                                    if (-1 <= -1) {
                                                        if v27 {
                                                            v8 = -1
                                                            var v78 uint64 = 4294967295
                                                            var v79 uint32 = 605
                                                            _ = v78
                                                            _ = v79
                                                        }
                                                        switch fi39 {
                                                        case -1:
                                                            v8 -= v5
                                                        case 32768:
                                                            if (-1 <= int8(v5)) {
                                                                if v27 {
                                                                    v80 := []int{(fi31 >> ((genFn0()) & 63)), (-(fi31)), (fi39 << ((42) & 63))}
                                                                    _ = v80
                                                                }
                                                            }
                                                        case 0:
                                                            if v27 {
                                                                v2 -= 1.5
                                                                v77 = "\n\t"
                                                            } else {
                                                                if ((v21 < 904) && v27) {
                                                                    fmt.Println(v1)
                                                                } else {
                                                                    fmt.Println(v1)
                                                                    v77 = v77
                                                                    var v81 bool = (1 == 398)
                                                                    _ = v81
                                                                }
                                                                switch (fi39 & (fi39 % ((fi39) | 1))) {
                                                                case 128:
                                                                    var v82 bool = ((((v1 * 3.14159) - 3.14159) * (v1 + v1)) >= (float32(((float64(float32(v4)) / (((-2.25 * 0.1) + v2) / (((((1.5 + v2) - ((((float64(((v1 - ((((0.1 + 0.1) * v1) * 0.0) / v1)) * 3.14159)) / v2) * 1e10) * 0.1) * float64((1e10 + v1)))) / (v4 - 0.0)) + 1.5) * -2.25))) - 1.5)) + float32(0.0)))
                                                                    _ = v82
                                                                case 9223372036854775807:
                                                                    v8 = (^(v8))
                                                                case -156:
                                                                    var v83 int = (fi39 / ((32767) | 1))
                                                                    _ = v83
                                                                }
                                                            }
                                                        }
                                                    }
                                                case "hello":
                                                    fmt.Println(v27, v21, v8)
                                                default:
                                                    v84 := GenS7{F0: (v77 + v77), F1: (v21 | 42)}
                                                    _ = v84
                                                }
                                                switch (fi31 + (fi31 | (fi31 - (^(fi31))))) {
                                                case -368:
                                                    v85 := GenS1{F0: 4294967295, F1: 4294967295}
                                                    _ = v85
                                                case 4294967295:
                                                    if false {
                                                        if v27 {
                                                            v86 := v3[(("a" + v77) + "a")]
                                                            _ = v86
                                                        }
                                                        fmt.Println(v4)
                                                        fmt.Println(v1, v21)
                                                    } else {
                                                        v8 -= v8
                                                        v87 := []int{fi39, (fi31 ^ 2147483647), (fi39 / ((fi39) | 1)), fi31}
                                                        v2 = v4
                                                        _ = v87
                                                    }
                                                default:
                                                    v1 -= 1.5
                                                }
                                            }
                                            _ = v77
                                        }
                                        fmt.Println(v2, fi39, v8)
                                        var v88 uint = genFn0()
                                        _ = fi39
                                        _ = v88
                                    }
                                }
                                var v89 int32 = 2147483647
                                _ = v89
                            }
                        } else {
                            v4 -= v4
                            var v90 string = ""
                            v91 := []int{-9223372036854775808, (fi31 % ((-783) | 1)), (fi31 | (^(fi31))), (fi31 * fi31)}
                            _ = v90
                            _ = v91
                        }
                    default:
                        v4 = v4
                    }
                    v92 := []int{(-(fi31)), fi31}
                    _ = v92
                }
                _ = fi31
                _ = v32
            }
        }
        var v93 uint64 = 1
        if (v1 == (v1 * ((v1 - 0.0) + v1))) {
            v94 := v3["foo bar"]
            v95 := GenS7{F0: "foo bar", F1: (v21 * v21)}
            _ = v94
            _ = v95
        }
        _ = v93
    }
    switch 2147483647 {
    case 42:
        switch -59 {
        case 139:
            var v96 int32 = -1
            _ = v96
        case -906:
            var v97 bool = false
            _ = v97
        }
    case 2147483647:
        v98 := []int{-1, -974}
        _ = v98
    }
    if ("\n\t" == "\n\t") {
        switch "" {
        case "":
            if false {
                if (!((106 != 0) || (584 > -1))) {
                    v27 = v27
                    v99 := []int{242}
                    _ = v99
                } else {
                    if v27 {
                        for fi100 := 0; fi100 < 13; fi100++ {
                            v101 := []int{-2147483648, fi100, (^(fi100)), fi100}
                            for fi102 := 0; fi102 < 4; fi102++ {
                                v101 = append(v101, fi102)
                                v8 -= v5
                                var v103 int16 = 613
                                _ = fi102
                                _ = v103
                            }
                            var v104 int64 = v8
                            _ = fi100
                            _ = v101
                            _ = v104
                        }
                        v21 += 2
                        v27 = true
                    }
                    var v105 float32 = (v1 + 1e10)
                    _ = v105
                }
            }
        default:
            switch 865 {
            case -679:
                var v106 uint64 = 2
                _ = v106
            case 956:
                for fi107 := 0; fi107 < 9; fi107++ {
                    v108 := v3["hello"]
                    _ = fi107
                    _ = v108
                }
            case 1:
                var v109 bool = ((638 <= 705) || false)
                _ = v109
            default:
                v3["foo bar"] = uint32(v21)
            }
        }
        if v27 {
            var v110 bool = false
            v8 -= (^(v8))
            _ = v110
        } else {
            for fi111 := 0; fi111 < 4; fi111++ {
                v2 = v4
                _ = fi111
            }
            v112 := GenS8{F0: 2, F1: -32768}
            switch "foo bar" {
            case "foo bar":
                switch 0 {
                case 127:
                    var v113 int16 = 0
                    _ = v113
                case 42:
                    if false {
                        v21 -= (^(v21))
                        var v114 uint64 = 921
                        _ = v114
                    }
                default:
                    switch "\n\t" {
                    case "hello":
                        if false {
                            var v115 int = 127
                            _ = v115
                        } else {
                            var v116 float32 = float32(3.14159)
                            switch 742 {
                            case 128:
                                v117 := GenS4{F0: 31, F1: 0}
                                _ = v117
                            default:
                                v27 = false
                            }
                            _ = v116
                        }
                    }
                }
            case "":
                var v118 uint = 42
                _ = v118
            case "hello":
                if v27 {
                    for fi119 := 0; fi119 < 8; fi119++ {
                        v1 -= v1
                        _ = fi119
                    }
                    switch ((("a" + (("\n\t" + "hello") + "\n\t")) + "a") + "hello") {
                    case "\n\t":
                        switch "\n\t" {
                        case "":
                            switch -865 {
                            case 4294967295:
                                v120 := GenS4{F0: 127, F1: 42}
                                _ = v120
                            case 42:
                                switch "foo bar" {
                                case "\n\t":
                                    v4 = 3.14159
                                case "a":
                                    var v121 int = -214
                                    _ = v121
                                default:
                                    v8 += v8
                                }
                            default:
                                v122 := GenS3{F0: v8, F1: (v27 && v27), F2: 42}
                                _ = v122
                            }
                        default:
                            v4 = ((v4 * v4) / 1e10)
                        }
                    case "foo bar":
                        v123 := []int{}
                        _ = v123
                    case "a":
                        switch "hello" {
                        case "hello":
                            v124 := GenS6{F0: true}
                            _ = v124
                        case "\n\t":
                            for fi125 := 0; fi125 < 12; fi125++ {
                                v21 = 2
                                v27 = v27
                                for fi126 := 0; fi126 < 11; fi126++ {
                                    var v127 int = fi126
                                    _ = fi126
                                    _ = v127
                                }
                                _ = fi125
                            }
                        case "":
                            for fi128 := 0; fi128 < 15; fi128++ {
                                fmt.Println(v4, fi128)
                                var v129 bool = ((!(1 >= -1)) || v27)
                                _ = fi128
                                _ = v129
                            }
                        default:
                            var v130 int = -2147483648
                            _ = v130
                        }
                    default:
                        var v131 float32 = (v1 / 1e10)
                        _ = v131
                    }
                }
            }
            _ = v112
        }
        v5 -= -853
    } else {
        v1 = (v1 * (v1 + float32(1.5)))
        if (!(("hello" + "foo bar") >= "\n\t")) {
            var v132 uint32 = 576
            _ = v132
        } else {
            if (631 >= 0) {
                for fi133 := 0; fi133 < 10; fi133++ {
                    if v27 {
                        var v134 bool = v27
                        if (2147483647 <= 118) {
                            for fi135 := 0; fi135 < 8; fi135++ {
                                v8 += (-(v8))
                                v136 := []int{(fi133 % ((0) | 1)), fi135, int(v8), fi133}
                                _ = fi135
                                _ = v136
                            }
                            for fi137 := 0; fi137 < 11; fi137++ {
                                switch "hello" {
                                case "a":
                                    if (-1 <= 24) {
                                        var v138 int32 = -2147483648
                                        v139 := GenS4{F0: -90, F1: 128}
                                        _ = v138
                                        _ = v139
                                    }
                                default:
                                    var v140 uint64 = 18446744073709551615
                                    _ = v140
                                }
                                var v141 uint = 406
                                _ = fi137
                                _ = v141
                            }
                        } else {
                            v21 = v21
                            for fi142 := 0; fi142 < 19; fi142++ {
                                var v143 float32 = (v1 * v1)
                                fmt.Println(v1, fi142)
                                if (-32768 < -284) {
                                    var v144 int16 = -1
                                    v145 := GenS9{F0: 2147483648, F1: 376}
                                    fmt.Println(v143, v5)
                                    fmt.Println(v2, fi142)
                                    _ = v144
                                    _ = v145
                                } else {
                                    if (127 >= 100) {
                                        if v27 {
                                            var v146 int8 = 42
                                            v147 := GenS7{F0: "a", F1: (^(v21))}
                                            var v148 int64 = v8
                                            _ = v146
                                            _ = v147
                                            _ = v148
                                        }
                                        var v149 int32 = -95
                                        _ = v149
                                    }
                                    var v150 int8 = -1
                                    var v151 uint32 = 1
                                    _ = v150
                                    _ = v151
                                }
                                var v152 uint64 = 850
                                _ = fi142
                                _ = v143
                                _ = v152
                            }
                        }
                        _ = v134
                    }
                    for fi153 := 0; fi153 < 5; fi153++ {
                        var v154 bool = (128 < 1)
                        _ = fi153
                        _ = v154
                    }
                    _ = fi133
                }
            } else {
                var v155 int64 = (v5 >> ((0) & 63))
                switch "\n\t" {
                case "\n\t":
                    var v156 int32 = 31
                    _ = v156
                case "foo bar":
                    v157 := []int{}
                    _ = v157
                }
                v158 := GenS10{F0: -1, F1: (-(v8)), F2: "hello"}
                _ = v155
                _ = v158
            }
        }
        v5 -= (v5 >> ((0) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
