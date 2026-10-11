package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 uint32; F2 string }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 uint16; F1 int16; F2 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint; F1 int64 }
type GenS11 struct { F0 uint8; F1 int }
type GenS12 struct { F0 int; F1 int32 }
type GenS13 struct { F0 string; F1 int64; F2 uint16 }
func genFn0(p0 uint8) uint64 {
    switch "\n\t" {
    case "hello":
        v7 := []int{}
        _ = v7
    }
    v8 := GenS1{F0: 565}
    v9 := []int{42, 128}
    fmt.Println(p0)
    p0 -= (p0 - (p0 % ((p0) | 1)))
    _ = p0
    _ = v8
    _ = v9
    return 478
}
func genFn1(p0 int64) int32 {
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := []int{fi10}
        _ = fi10
        _ = v11
    }
    var v12 uint64 = 42
    fmt.Println(v12, v12)
    v13 := []int{-309}
    var v14 int8 = -128
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return -2147483648
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 bool = (711 > -1)
    v4 := []string{"foo bar", "\n\t", ("hello" + ("\n\t" + "a")), ""}
    v5 := make(map[string]bool)
    var v6 int8 = -1
    v15 := GenS1{F0: 65535}
    if ("foo bar" != ((("a" + "hello") + "") + ((((((("a" + "a") + "") + "") + "hello") + "hello") + (("hello" + "") + (((("" + "\n\t") + ("hello" + "")) + "\n\t") + "a"))) + ""))) {
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 bool = v3
            v18 := GenS2{F0: 0}
            v3 = v17
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 uint32 = 481
        _ = v19
    } else {
        if v3 {
            for fi20 := 0; fi20 < 14; fi20++ {
                if (128 == 0) {
                    switch "" {
                    case "":
                        if (866 >= 42) {
                            v3 = v3
                        }
                    case "a":
                        fmt.Println(v6, v3, fi20)
                    }
                }
                _ = fi20
            }
            v21 := GenS1{F0: 32767}
            v22 := GenS2{F0: 2}
            _ = v21
            _ = v22
        } else {
            if (!v3) {
                fmt.Println(v6)
                for fi23 := 0; fi23 < 12; fi23++ {
                    var v24 uint8 = 0
                    var v25 int = (-(fi23))
                    if (65535 == 625) {
                        switch v25 {
                        case 128:
                            if (!false) {
                                v3 = (80 > 1)
                            } else {
                                v25 -= v25
                                if false {
                                    if v3 {
                                        if true {
                                            switch "foo bar" {
                                            case "\n\t":
                                                if (!(!v3)) {
                                                    fmt.Println(v3, fi23, v24)
                                                    if v3 {
                                                        v25 -= 9223372036854775807
                                                    }
                                                }
                                            case "a":
                                                v25 += 127
                                            default:
                                                v25 -= (v25 & 128)
                                            }
                                            var v26 float32 = (0.1 - 0.0)
                                            if false {
                                                fmt.Println(v24, fi23)
                                                for fi27 := 0; fi27 < 1; fi27++ {
                                                    v28 := GenS2{F0: 0}
                                                    if true {
                                                        delete(v5, "\n\t")
                                                    }
                                                    var v29 int64 = 0
                                                    _ = fi27
                                                    _ = v28
                                                    _ = v29
                                                }
                                                for fi30 := 0; fi30 < 2; fi30++ {
                                                    v31 := GenS2{F0: 636}
                                                    switch "a" {
                                                    case "\n\t":
                                                        var v32 float64 = (1.5 * 1e10)
                                                        _ = v32
                                                    case "foo bar":
                                                        v6 = (v6 ^ 1)
                                                    }
                                                    v33 := GenS3{F0: 316, F1: -32768}
                                                    _ = fi30
                                                    _ = v31
                                                    _ = v33
                                                }
                                            } else {
                                                var v34 int32 = 42
                                                _ = v34
                                            }
                                            _ = v26
                                        }
                                        v4 = append(v4, "\n\t")
                                    }
                                    for fi35 := 0; fi35 < 17; fi35++ {
                                        v24 = v24
                                        v36 := GenS4{F0: 4294967295}
                                        _ = fi35
                                        _ = v36
                                    }
                                } else {
                                    v37 := GenS1{F0: 0}
                                    switch (fi23 & v25) {
                                    case 1:
                                        v25 += fi23
                                    case -1:
                                        switch (("a" + "hello") + "a") {
                                        case "foo bar":
                                            switch "" {
                                            case "hello":
                                                fmt.Println(v25)
                                            case "foo bar":
                                                v6 = v6
                                            case "\n\t":
                                                fmt.Println(fi23, v3, v25)
                                            }
                                        case "\n\t":
                                            var v38 int = fi23
                                            _ = v38
                                        default:
                                            v4 = append(v4, ("hello" + "hello"))
                                        }
                                    default:
                                        v25 = (v25 - (v25 | (fi23 % (((v25 ^ v25)) | 1))))
                                    }
                                    _ = v37
                                }
                            }
                        }
                        fmt.Println(v6)
                    }
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            } else {
                var v39 float32 = 0.0
                for fi40 := 0; fi40 < 8; fi40++ {
                    if ((fi40 - fi40) <= 9223372036854775807) {
                        var v41 uint32 = 35
                        _ = v41
                    }
                    if (v3 && true) {
                        var v42 uint32 = 4294967295
                        _ = v42
                    } else {
                        fmt.Println(fi40, v39, v39)
                        for fi43 := 0; fi43 < 19; fi43++ {
                            v44 := v5[""]
                            switch ("hello" + ((("" + "foo bar") + "") + "hello")) {
                            case "hello":
                                v45 := GenS2{F0: 2147483647}
                                _ = v45
                            case "foo bar":
                                v3 = v3
                            case "a":
                                v44 = (32767 <= -1)
                            default:
                                v46 := GenS5{F0: 9223372036854775807, F1: 2, F2: "hello"}
                                v3 = v44
                                _ = v46
                            }
                            _ = fi43
                            _ = v44
                        }
                    }
                    _ = fi40
                }
                v3 = true
                _ = v39
            }
            v47 := GenS4{F0: 206}
            _ = v47
        }
        var v48 uint32 = 2147483647
        _ = v48
    }
    switch -724 {
    case -90:
        if v3 {
            var v49 uint64 = genFn0(1)
            switch 9223372036854775807 {
            case -1:
                v3 = (0.1 < 0.0)
            case 4294967295:
                v49 += 0
            default:
                if v3 {
                    _ = len(v4)
                    switch (("foo bar" + "\n\t") + "foo bar") {
                    case "":
                        v50 := v2[-1]
                        _ = v50
                    default:
                        var v51 int = 128
                        _ = v51
                    }
                }
            }
            _ = v49
        } else {
            if (((!(42 == 0)) || true) && (v3 && v3)) {
                for fi52 := 0; fi52 < 3; fi52++ {
                    var v53 bool = (v3 && v3)
                    v53 = ((v3 || v3) && v53)
                    v5["a"] = false
                    _ = fi52
                    _ = v53
                }
                v3 = (!(0.1 <= 3.14159))
            } else {
                fmt.Println(v3, v3)
                if (!v3) {
                    var v54 int32 = 1
                    v55 := GenS4{F0: 1}
                    if false {
                        switch "" {
                        case "a":
                            if false {
                                if v3 {
                                    v56 := v5[("a" + "a")]
                                    _ = v56
                                } else {
                                    if (!true) {
                                        if (-9223372036854775808 >= 1) {
                                            fmt.Println(v6, v6)
                                            var v57 int64 = 0
                                            _ = v57
                                        } else {
                                            v3 = ((!(!(v3 || (1 == 2)))) && false)
                                        }
                                        switch -1 {
                                        case 79:
                                            var v58 bool = ((v54 | (v54 % (((v54 * (v54 & v54))) | 1))) <= v54)
                                            _ = v58
                                        case -2147483648:
                                            if true {
                                                for fi59 := 0; fi59 < 10; fi59++ {
                                                    var v60 int = fi59
                                                    if (v3 && (992 != -750)) {
                                                        if (v3 || v3) {
                                                            v4 = append(v4, "hello")
                                                        }
                                                        fmt.Println(v60, v6)
                                                        var v61 float32 = 1e10
                                                        _ = v61
                                                    }
                                                    _ = fi59
                                                    _ = v60
                                                }
                                            }
                                        case -960:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                    }
                                }
                                for fi62 := 0; fi62 < 9; fi62++ {
                                    for fi63 := 0; fi63 < 17; fi63++ {
                                        switch "\n\t" {
                                        case "a":
                                            if (1 >= 145) {
                                                if v3 {
                                                    switch (^(fi63)) {
                                                    case 127:
                                                        v64 := GenS6{F0: 2147483647, F1: 255}
                                                        v54 = 1
                                                        var v65 uint64 = 18446744073709551615
                                                        _ = v64
                                                        _ = v65
                                                    default:
                                                        fmt.Println(v3)
                                                    }
                                                    var v66 float64 = (1.5 + 3.14159)
                                                    var v67 uint16 = 255
                                                    v6 -= (v6 * v6)
                                                    _ = v66
                                                    _ = v67
                                                } else {
                                                    var v68 uint = 65535
                                                    _ = v68
                                                }
                                            }
                                        }
                                        var v69 bool = ("\n\t" <= ("a" + ("foo bar" + "a")))
                                        if (v6 <= (v6 | int8(v54))) {
                                            v69 = v69
                                            if len(v4) > 0 {
                                                v4[0] = ("a" + "\n\t")
                                            }
                                        }
                                        _ = fi63
                                        _ = v69
                                    }
                                    _ = fi62
                                }
                            }
                        case "\n\t":
                            var v70 uint64 = 2
                            _ = v70
                        case "hello":
                            var v71 int64 = 0
                            _ = v71
                        }
                        fmt.Println(v6, v54, v6)
                        for fi72 := 0; fi72 < 14; fi72++ {
                            switch -55 {
                            case 32768:
                                var v73 int = (^(fi72))
                                _ = v73
                            case 42:
                                fmt.Println(v3)
                            default:
                                switch 127 {
                                case 32767:
                                    switch -1 {
                                    case -2147483648:
                                        var v74 int8 = v6
                                        _ = v74
                                    }
                                case -1:
                                    var v75 int32 = -1
                                    _ = v75
                                }
                            }
                            _ = fi72
                        }
                    } else {
                        for fi76 := 0; fi76 < 9; fi76++ {
                            for fi77 := 0; fi77 < 19; fi77++ {
                                v6 = v6
                                _ = fi77
                            }
                            fmt.Println(fi76, v54, v54)
                            var v78 int32 = v54
                            _ = fi76
                            _ = v78
                        }
                    }
                    _ = v54
                    _ = v55
                } else {
                    v79 := GenS7{F0: 2147483647, F1: 828}
                    _ = v79
                }
            }
        }
    default:
        var v80 uint64 = 831
        _ = v80
    }
    v81 := GenS8{F0: 42, F1: 32767, F2: -9223372036854775808}
    var v82 int8 = v6
    switch 2147483647 {
    case 424:
        delete(v5, ("" + ""))
    }
    if (974 < 65535) {
        for fi83 := 0; fi83 < 3; fi83++ {
            var v84 uint16 = 476
            v85 := GenS3{F0: 1, F1: 1000}
            _ = fi83
            _ = v84
            _ = v85
        }
        v3 = v3
        if true {
            for fi86 := 0; fi86 < 12; fi86++ {
                if v3 {
                    if ((v6 >> ((255) & 63)) <= v6) {
                        if ((v3 && true) || (!(2 < 4294967295))) {
                            v3 = (718 != 1000)
                            if v3 {
                                for fi87 := 0; fi87 < 20; fi87++ {
                                    if v3 {
                                        var v88 string = ("\n\t" + "")
                                        _ = v88
                                    }
                                    _ = fi87
                                }
                                for fi89 := 0; fi89 < 4; fi89++ {
                                    v4 = append(v4, (("foo bar" + "") + "a"))
                                    var v90 uint = 255
                                    _ = fi89
                                    _ = v90
                                }
                            }
                            v91 := GenS6{F0: -2147483648, F1: 0}
                            _ = v91
                        }
                        var v92 int32 = 734
                        v93 := GenS5{F0: 42, F1: 2, F2: "foo bar"}
                        _ = v92
                        _ = v93
                    } else {
                        v94 := GenS7{F0: fi86, F1: 169}
                        v95 := GenS9{F0: ""}
                        v96 := GenS5{F0: -257, F1: 65535, F2: "hello"}
                        _ = v94
                        _ = v95
                        _ = v96
                    }
                }
                _ = fi86
            }
            v3 = (!v3)
        } else {
            if len(v4) > 0 {
                v4[0] = ("" + "a")
            }
            v82 += 1
            var v97 int64 = -1
            _ = v97
        }
    }
    for fi98 := 0; fi98 < 4; fi98++ {
        switch fi98 {
        case 9223372036854775807:
            fmt.Println(fi98)
        case 4294967295:
            var v99 int64 = 966
            _ = v99
        }
        switch ((("foo bar" + "\n\t") + "") + ("a" + "a")) {
        case "":
            if true {
                for fi100 := 0; fi100 < 15; fi100++ {
                    if (!v3) {
                        v101 := GenS10{F0: 255, F1: 9223372036854775807}
                        _ = v101
                    }
                    _ = fi100
                }
                if v3 {
                    delete(v2, (fi98 ^ (fi98 << ((803) & 63))))
                    if (false || v3) {
                        for fi102 := 0; fi102 < 19; fi102++ {
                            var v103 uint16 = 196
                            switch "" {
                            case "":
                                var v104 uint8 = 144
                                _ = v104
                            }
                            if (879 != 1) {
                                var v105 uint8 = 253
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v105
                            } else {
                                _ = len(v4)
                                if (!(v3 && (!(!false)))) {
                                    v106 := GenS4{F0: 2147483648}
                                    v103 += v103
                                    _ = v106
                                }
                                switch fi102 {
                                case -515:
                                    if v3 {
                                        switch 0 {
                                        case 32768:
                                            fmt.Println(v6, v82)
                                        case 779:
                                            v82 = (-(v82))
                                        case -2147483648:
                                            var v107 float64 = 3.14159
                                            _ = v107
                                        }
                                        if true {
                                            var v108 int32 = 2147483647
                                            _ = v108
                                        }
                                    }
                                case 32767:
                                    v82 += v82
                                default:
                                    var v109 int32 = 1
                                    _ = v109
                                }
                            }
                            _ = fi102
                            _ = v103
                        }
                        switch "foo bar" {
                        case "hello":
                            var v110 bool = (255 != 358)
                            _ = v110
                        default:
                            for fi111 := 0; fi111 < 11; fi111++ {
                                if v3 {
                                    v6 = v82
                                }
                                _ = fi111
                            }
                        }
                    }
                } else {
                    if (v3 && ((v82 * v82) < (v6 >> ((1) & 63)))) {
                        for fi112 := 0; fi112 < 5; fi112++ {
                            switch (fi112 * (-(fi112))) {
                            case 127:
                                if (v82 == 42) {
                                    var v113 int16 = 1
                                    var v114 float32 = (((3.14159 / -2.25) + 0.1) - -2.25)
                                    _ = v113
                                    _ = v114
                                } else {
                                    v3 = (1 <= 32767)
                                    var v115 int32 = 2147483647
                                    _ = v115
                                }
                            default:
                                if (!(v3 || (-2147483648 >= genFn1(0)))) {
                                    v116 := GenS11{F0: 2, F1: fi98}
                                    _ = v116
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = ""
                                    }
                                    v6 += v82
                                }
                                fmt.Println(v3)
                            }
                            var v117 float32 = (3.14159 - 1.5)
                            fmt.Println(v117, v3)
                            if (370 >= 42) {
                                v118 := GenS12{F0: (^(fi112)), F1: -2147483648}
                                _ = len(v5)
                                switch (("" + "foo bar") + "\n\t") {
                                case "hello":
                                    var v119 uint = 444
                                    _ = v119
                                }
                                _ = v118
                            }
                            _ = fi112
                            _ = v117
                        }
                        if ((885 <= 0) || false) {
                            var v120 uint8 = 42
                            fmt.Println(v6, v82, v120)
                            var v121 uint = 248
                            _ = v120
                            _ = v121
                        }
                        for fi122 := 0; fi122 < 14; fi122++ {
                            if true {
                                if (1 <= v6) {
                                    if v3 {
                                        fmt.Println(fi98, fi122)
                                        v6 = (v82 & v6)
                                        v6 -= (v6 + -98)
                                    }
                                } else {
                                    var v123 uint64 = 42
                                    _ = v123
                                }
                                switch ("\n\t" + (((("\n\t" + "") + ((("a" + "") + "\n\t") + "a")) + "\n\t") + "")) {
                                case "\n\t":
                                    v2[fi122] = (v3 || ((0.0 != 1e10) || v3))
                                    var v124 int = 42
                                    v125 := v5["\n\t"]
                                    _ = v124
                                    _ = v125
                                case "hello":
                                    if v3 {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v126 int = fi122
                                        _ = v126
                                    } else {
                                        switch fi98 {
                                        case 32767:
                                            if (56 < 2) {
                                                var v127 float32 = 3.14159
                                                v128 := GenS6{F0: genFn1(-1), F1: 1}
                                                v6 = 100
                                                _ = v127
                                                _ = v128
                                            } else {
                                                var v129 uint64 = genFn0(2)
                                                _ = v129
                                            }
                                        case -116:
                                            fmt.Println(v82, v82)
                                        case 127:
                                            if len(v4) > 0 {
                                                v4[0] = "foo bar"
                                            }
                                        }
                                        v3 = true
                                        _ = len(v4)
                                    }
                                }
                                v1[(fi122 | 32768)] = 639
                            }
                            _ = fi122
                        }
                    }
                }
            }
        case "\n\t":
            if v3 {
                v130 := v2[666]
                _ = v130
            } else {
                v82 -= 42
            }
        case "hello":
            switch (fi98 & (^(fi98))) {
            case 128:
                v4 = append(v4, "")
            case 9223372036854775807:
                _ = len(v2)
            case -2147483648:
                var v131 uint16 = 32767
                _ = v131
            default:
                fmt.Println(v3, v6, v82)
            }
        default:
            v132 := GenS13{F0: "hello", F1: 0, F2: 683}
            _ = v132
        }
        _ = fi98
    }
    fmt.Println(v82)
    if (42 >= 4294967295) {
        fmt.Println(v3)
        if (!(v3 && ((v3 || true) || v3))) {
            var v133 float64 = 3.14159
            if false {
                var v134 uint64 = 1
                _ = v134
            }
            _ = v133
        } else {
            if v3 {
                var v135 uint32 = 2147483648
                var v136 string = ""
                v82 = 0
                _ = v135
                _ = v136
            } else {
                v82 = v6
                fmt.Println(v3, v3, v6)
                v137 := GenS8{F0: 738, F1: -483, F2: 0}
                _ = v137
            }
            var v138 uint64 = 4294967295
            _ = v138
        }
    }
    v6 = -128
    fmt.Println(v3, v6)
    var v139 int64 = 1
    var v140 int = -71
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v139:", v139)
    fmt.Println("v140:", v140)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v139
    _ = v140
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v82
}
