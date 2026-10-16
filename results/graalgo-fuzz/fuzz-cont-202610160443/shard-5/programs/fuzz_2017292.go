package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int; F1 string; F2 int16 }
type GenS3 struct { F0 bool; F1 string; F2 uint64 }
type GenS4 struct { F0 string; F1 bool; F2 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int64; F1 uint16 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 bool; F1 uint64; F2 uint32 }
func genFn0(p0 uint64, p1 uint16, p2 uint16) uint32 {
    v7 := []int{-1, 32768, -2147483648, -724}
    fmt.Println(p1)
    for fi8 := 0; fi8 < 11; fi8++ {
        p0 = (^(p0))
        v9 := GenS1{F0: 0, F1: 2}
        var v10 uint8 = 2
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        p1 = p1
        _ = len(v7)
        _ = fi11
    }
    p0 = (p0 + p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 34
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = false
    var v3 bool = v2
    var v4 uint64 = 4294967295
    var v5 int = 4294967295
    var v6 uint = 2
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            if v2 {
                switch fi13 {
                case 128:
                    var v14 int32 = -1
                    _ = v14
                case -417:
                    switch "" {
                    case "hello":
                        if (0.0 <= -2.25) {
                            switch -983 {
                            case -980:
                                v6 += (^(v6))
                            case 32768:
                                var v15 int16 = 32767
                                _ = v15
                            }
                        } else {
                            var v16 int64 = 622
                            var v17 uint8 = 0
                            var v18 int32 = int32(v5)
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                    case "\n\t":
                        var v19 float32 = 0.1
                        _ = v19
                    case "a":
                        if (v2 && v2) {
                            fmt.Println(fi13, fi13)
                            v20 := []int{fi12, (^(fi13)), 42, fi13}
                            v5 -= fi12
                            _ = v20
                        } else {
                            var v21 uint32 = 2147483647
                            v4 -= (v4 % ((v4) | 1))
                            v21 = v21
                            switch fi13 {
                            case 32767:
                                v22 := GenS2{F0: (-(fi13)), F1: "a", F2: -372}
                                _ = v22
                            case 32768:
                                var v23 int = 127
                                _ = v23
                            case 1:
                                var v24 bool = (212 == 0)
                                _ = v24
                            }
                            var v25 float32 = 1e10
                            _ = v21
                            _ = v25
                        }
                    }
                case 9223372036854775807:
                    v26 := []int{(v5 << ((v6) & 63)), v5}
                    _ = v26
                default:
                    var v27 int = 127
                    _ = v27
                }
                v28 := v1[fi13]
                v29 := []int{}
                switch fi12 {
                case -1:
                    v28 = (v28 + v28)
                case 4294967295:
                    var v30 uint16 = 65535
                    _ = v30
                }
                fmt.Println(v28, v2)
                _ = v28
                _ = v29
            }
            v31 := GenS3{F0: v3, F1: "a", F2: (v4 - 18446744073709551615)}
            var v32 int8 = 2
            _ = fi13
            _ = v31
            _ = v32
        }
        if v2 {
            switch "hello" {
            case "":
                var v33 float32 = 1e10
                _ = v33
            case "foo bar":
                switch (fi12 / (((v5 << ((v6) & 63))) | 1)) {
                case 784:
                    var v34 uint32 = 212
                    _ = v34
                }
            }
            switch "hello" {
            case "\n\t":
                v2 = (3.14159 < 1.5)
            case "":
                if false {
                    if (!v2) {
                        var v35 bool = v2
                        _ = v35
                    }
                    fmt.Println(v6, fi12)
                    v36 := []int{(^(v5)), (fi12 - (fi12 << (((v6 + v6)) & 63)))}
                    _ = v36
                }
            default:
                for fi37 := 0; fi37 < 18; fi37++ {
                    switch "\n\t" {
                    case "foo bar":
                        v5 += fi37
                    case "":
                        switch ("foo bar" + "a") {
                        case "foo bar":
                            v2 = (!(!false))
                        case "\n\t":
                            v38 := GenS3{F0: (0 <= 42), F1: "foo bar", F2: (v4 | v4)}
                            _ = v38
                        case "hello":
                            v39 := GenS4{F0: "a", F1: v2, F2: 0}
                            _ = v39
                        default:
                            if (v2 || (!false)) {
                                v3 = ((v6 >> (((v6 << ((42) & 63))) & 63)) != v6)
                            } else {
                                v5 -= (v5 ^ 9223372036854775807)
                                v2 = (!((false && v3) && (!("a" >= ("a" + "hello")))))
                            }
                        }
                    case "a":
                        if v2 {
                            var v40 int64 = 0
                            v41 := GenS5{F0: 0}
                            fmt.Println(v2, v5)
                            _ = v40
                            _ = v41
                        } else {
                            if (755 <= 1) {
                                v42 := GenS5{F0: 0}
                                v43 := []int{}
                                fmt.Println(v6)
                                _ = v42
                                _ = v43
                            }
                            switch "\n\t" {
                            case "":
                                var v44 float64 = 3.14159
                                _ = v44
                            }
                        }
                    }
                    _ = fi37
                }
            }
            v2 = v3
        } else {
            switch "" {
            case "hello":
                v45 := GenS2{F0: (fi12 % ((fi12) | 1)), F1: "\n\t", F2: -684}
                _ = v45
            case "":
                switch (^(v5)) {
                case 32768:
                    v46 := []int{(v5 << ((1) & 63)), v5}
                    _ = v46
                case -1:
                    var v47 uint32 = genFn0(v4, 32767, 926)
                    _ = v47
                }
            default:
                for fi48 := 0; fi48 < 11; fi48++ {
                    if v3 {
                        var v49 uint16 = 2
                        v49 += 2
                        _ = v49
                    } else {
                        v2 = (!v3)
                        v50 := []int{2147483647}
                        if (!v3) {
                            fmt.Println(fi48, v5, fi48)
                        } else {
                            v5 = v5
                            v2 = v3
                        }
                        _ = v50
                    }
                    if (int8(fi12) >= int8(v5)) {
                        if (!true) {
                            delete(v1, (fi12 - 99))
                            v51 := v1[fi12]
                            _ = v51
                        } else {
                            v52 := []int{}
                            var v53 int16 = 1
                            v54 := GenS6{F0: v6}
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                        var v55 uint8 = 128
                        fmt.Println(v5, v3, v5)
                        var v56 int16 = 32767
                        v3 = v2
                        _ = v55
                        _ = v56
                    }
                    v57 := GenS1{F0: 2147483647, F1: 127}
                    _ = fi48
                    _ = v57
                }
            }
        }
        var v58 float32 = (0.1 / -2.25)
        _ = fi12
        _ = v58
    }
    _ = len(v1)
    fmt.Println(v2)
    if (42 > 2) {
        fmt.Println(v3, v5, v4)
        for fi59 := 0; fi59 < 8; fi59++ {
            v60 := []int{311}
            _ = fi59
            _ = v60
        }
        if (127 > (v5 + int(v6))) {
            v61 := []int{v5, 693, v5, (v5 - v5)}
            v6 += 721
            v62 := GenS2{F0: v5, F1: ("" + "\n\t"), F2: 139}
            _ = v61
            _ = v62
        } else {
            v5 += v5
        }
    } else {
        v63 := v1[v5]
        v64 := GenS6{F0: v6}
        v6 += v6
        _ = v63
        _ = v64
    }
    v3 = (v2 && (v2 && (!(false && v2))))
    switch ("hello" + "\n\t") {
    case "\n\t":
        var v65 uint = v6
        _ = v65
    case "a":
        for fi66 := 0; fi66 < 11; fi66++ {
            if v3 {
                var v67 int16 = int16(fi66)
                var v68 int64 = -9223372036854775808
                _ = v67
                _ = v68
            }
            if true {
                if (-914 < 42) {
                    for fi69 := 0; fi69 < 8; fi69++ {
                        switch 32768 {
                        case 0:
                            if ((!(1 <= 42)) && false) {
                                var v70 int32 = 1
                                if v2 {
                                    v71 := []int{208}
                                    v72 := GenS3{F0: (!false), F1: ("foo bar" + "\n\t"), F2: v4}
                                    _ = v71
                                    _ = v72
                                }
                                if v3 {
                                    var v73 string = "hello"
                                    switch 4294967295 {
                                    case 32768:
                                        var v74 int = fi66
                                        _ = v74
                                    case -559:
                                        var v75 int8 = 127
                                        _ = v75
                                    case 4294967295:
                                        switch 32767 {
                                        case -9223372036854775808:
                                            var v76 int32 = (v70 | (^(v70)))
                                            _ = v76
                                        default:
                                            delete(v1, (^(fi69)))
                                        }
                                    }
                                    _ = v73
                                }
                                _ = v70
                            } else {
                                var v77 string = ("\n\t" + "foo bar")
                                var v78 int = -9223372036854775808
                                v79 := []int{}
                                v6 = v6
                                fmt.Println(v4)
                                _ = v77
                                _ = v78
                                _ = v79
                            }
                            v4 = (^(v4))
                            switch "\n\t" {
                            case "":
                                v4 = v4
                            case "a":
                                var v80 float64 = 1e10
                                _ = v80
                            }
                            if v2 {
                                switch "" {
                                case "foo bar":
                                    v81 := GenS6{F0: v6}
                                    _ = v81
                                case "\n\t":
                                    var v82 int8 = 100
                                    _ = v82
                                default:
                                    v2 = v3
                                    delete(v1, -9223372036854775808)
                                    if false {
                                        var v83 int = v5
                                        switch (v5 * fi69) {
                                        case 994:
                                            v1[(-(fi69))] = "a"
                                        default:
                                            switch (fi69 >> ((1) & 63)) {
                                            case 885:
                                                v3 = true
                                            }
                                        }
                                        _ = v83
                                    }
                                }
                                if ((fi69 << (((^(v6))) & 63)) != v5) {
                                    v5 = fi66
                                    if (1 < -1) {
                                        if ((1000 <= 1000) && v3) {
                                            v84 := GenS7{F0: (!false), F1: 2147483647}
                                            _ = v84
                                        }
                                    } else {
                                        v4 -= v4
                                    }
                                } else {
                                    v85 := []int{(^(v5)), (fi66 << (((v6 ^ (v6 - 0))) & 63))}
                                    v4 = (v4 % ((0) | 1))
                                    _ = v85
                                }
                                v5 = (^(v5))
                            } else {
                                switch "hello" {
                                case "\n\t":
                                    var v86 uint64 = 4294967295
                                    _ = v86
                                case "a":
                                    v4 -= 920
                                default:
                                    delete(v1, (v5 / ((int(v6)) | 1)))
                                }
                                fmt.Println(v2, v2)
                                if (v2 && true) {
                                    var v87 int = (fi66 & fi66)
                                    _ = v87
                                } else {
                                    v88 := GenS8{F0: 960}
                                    _ = v88
                                }
                            }
                            switch "\n\t" {
                            case "\n\t":
                                if (!(42 <= 255)) {
                                    if v3 {
                                        fmt.Println(fi69, v5)
                                    } else {
                                        v2 = (!true)
                                        v4 = (v4 / (((v4 & (v4 ^ (v4 & v4)))) | 1))
                                        var v89 int8 = -88
                                        _ = v89
                                    }
                                    if v3 {
                                        v4 = v4
                                        v3 = v3
                                    }
                                    fmt.Println(v6, v2, fi69)
                                } else {
                                    v3 = v3
                                    v90 := GenS8{F0: -596}
                                    _ = v90
                                }
                            }
                        case -9223372036854775808:
                            if true {
                                fmt.Println(fi69, v6)
                                v5 -= (v5 * fi69)
                                fmt.Println(fi66)
                                v91 := GenS4{F0: "foo bar", F1: ((42 == 127) && (v3 || ((!v2) && v3))), F2: -698}
                                _ = v91
                            }
                        case 4294967295:
                            switch fi66 {
                            case -292:
                                v92 := []int{-176}
                                _ = v92
                            default:
                                v2 = v2
                            }
                        default:
                            var v93 uint32 = 0
                            _ = v93
                        }
                        var v94 int16 = 1
                        _ = fi69
                        _ = v94
                    }
                    v6 -= 809
                } else {
                    var v95 int64 = 1
                    fmt.Println(v5, v5)
                    switch "" {
                    case "foo bar":
                        v96 := GenS9{F0: v4}
                        _ = v96
                    }
                    _ = v95
                }
            }
            var v97 uint64 = v4
            _ = fi66
            _ = v97
        }
    default:
        var v98 int64 = 9223372036854775807
        _ = v98
    }
    v99 := GenS10{F0: 1}
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v5, v4)
    case "hello":
        if (128 > 227) {
            fmt.Println(v6, v5, v6)
            var v100 uint = 65535
            fmt.Println(v3, v4)
            _ = v100
        } else {
            v101 := GenS7{F0: v2, F1: 65535}
            for fi102 := 0; fi102 < 14; fi102++ {
                if (190 != 2) {
                    v3 = (3.14159 < 0.0)
                    var v103 string = "foo bar"
                    _ = v103
                }
                _ = fi102
            }
            _ = v101
        }
    default:
        for fi104 := 0; fi104 < 19; fi104++ {
            if (!v2) {
                switch "\n\t" {
                case "a":
                    if true {
                        v105 := GenS4{F0: "foo bar", F1: (((((-2.25 - 0.1) / 1.5) + 0.0) + (float64(-2.25) - 0.0)) != 0.0), F2: -1}
                        _ = v105
                    }
                case "\n\t":
                    v5 += (fi104 * 42)
                case "hello":
                    if v3 {
                        var v106 int8 = -83
                        switch "hello" {
                        case "hello":
                            v107 := []int{(fi104 % ((fi104) | 1)), 0}
                            _ = v107
                        default:
                            fmt.Println(v6)
                        }
                        v108 := GenS11{F0: 42}
                        _ = v106
                        _ = v108
                    } else {
                        for fi109 := 0; fi109 < 3; fi109++ {
                            v5 = v5
                            switch "foo bar" {
                            case "\n\t":
                                v4 -= (v4 >> ((4294967295) & 63))
                                v110 := GenS8{F0: -965}
                                _ = v110
                            case "a":
                                fmt.Println(fi109, v4, v5)
                            default:
                                fmt.Println(v6, v2)
                            }
                            if true {
                                switch "a" {
                                case "foo bar":
                                    var v111 uint32 = 65535
                                    _ = v111
                                case "":
                                    if (156 <= 42) {
                                        fmt.Println(v2, v2, v3)
                                    } else {
                                        switch (^(v5)) {
                                        case 4294967295:
                                            switch "hello" {
                                            case "foo bar":
                                                switch ("hello" + (((("foo bar" + ("\n\t" + ("" + "\n\t"))) + "hello") + "a") + "\n\t")) {
                                                case "foo bar":
                                                    if v3 {
                                                        var v112 uint = 1
                                                        v1[(fi109 & 128)] = ("a" + (("hello" + "") + ""))
                                                        var v113 uint32 = 2147483648
                                                        _ = v112
                                                        _ = v113
                                                    }
                                                case "a":
                                                    if ((("foo bar" + "hello") + "\n\t") == (((("a" + "") + ("hello" + "a")) + "a") + "\n\t")) {
                                                        fmt.Println(v4, v6, fi109)
                                                        v5 -= 42
                                                    }
                                                    v114 := GenS12{F0: 1, F1: 443}
                                                    _ = v114
                                                case "\n\t":
                                                    var v115 float64 = 1e10
                                                    _ = v115
                                                default:
                                                    fmt.Println(v5)
                                                }
                                            case "a":
                                                if false {
                                                    fmt.Println(fi109, v3)
                                                }
                                            case "":
                                                var v116 int32 = -2147483648
                                                _ = v116
                                            }
                                        case 606:
                                            if v3 {
                                                var v117 int64 = -9223372036854775808
                                                v5 -= (fi104 & v5)
                                                fmt.Println(v6, v3, v2)
                                                _ = v117
                                            }
                                        default:
                                            var v118 int32 = 2147483647
                                            _ = v118
                                        }
                                    }
                                default:
                                    fmt.Println(v3, fi109)
                                }
                                switch (fi109 >> ((v6) & 63)) {
                                case 128:
                                    if (v6 < v6) {
                                        v119 := v1[(fi104 >> ((uint(v4)) & 63))]
                                        var v120 uint16 = 0
                                        if v3 {
                                            var v121 int32 = 849
                                            _ = v121
                                        }
                                        switch ((v119 + (v119 + v119)) + "foo bar") {
                                        case "\n\t":
                                            var v122 uint64 = v4
                                            _ = v122
                                        case "hello":
                                            fmt.Println(v2, v2)
                                        case "a":
                                            if ((fi104 / (((^(v5))) | 1)) == fi109) {
                                                if (v120 != (^(v120))) {
                                                    fmt.Println(v5, v4)
                                                    var v123 string = v119
                                                    v124 := GenS13{F0: v6}
                                                    _ = v123
                                                    _ = v124
                                                }
                                                v5 += v5
                                                var v125 int16 = 0
                                                v2 = v3
                                                if (((1 != -1) || (255 < 32767)) && v3) {
                                                    if v3 {
                                                        v125 = v125
                                                        switch "a" {
                                                        case "foo bar":
                                                            var v126 float32 = ((float32(float64(0.1)) * 3.14159) - (-2.25 + (float32(0.0) - float32((1e10 * 0.0)))))
                                                            _ = v126
                                                        }
                                                        if (1000 <= -32768) {
                                                            fmt.Println(v6)
                                                        }
                                                    } else {
                                                        var v127 bool = (42 != 1)
                                                        if (19 == 2) {
                                                            v128 := GenS8{F0: v125}
                                                            _ = v128
                                                        }
                                                        fmt.Println(v127, v127, v120)
                                                        _ = v127
                                                    }
                                                }
                                                var v129 bool = false
                                                _ = v125
                                                _ = v129
                                            } else {
                                                if (-2.25 == ((1.5 / 1.5) * 3.14159)) {
                                                    v6 += 65535
                                                    v130 := GenS11{F0: 0}
                                                    _ = v130
                                                }
                                                v131 := []int{fi104}
                                                _ = v131
                                            }
                                        default:
                                            var v132 uint64 = uint64(v6)
                                            v133 := GenS9{F0: v4}
                                            v134 := GenS11{F0: 0}
                                            _ = v132
                                            _ = v133
                                            _ = v134
                                        }
                                        var v135 int8 = 123
                                        _ = v119
                                        _ = v120
                                        _ = v135
                                    } else {
                                        v136 := GenS10{F0: uint16(v6)}
                                        _ = v136
                                    }
                                case 127:
                                    switch (fi104 << ((255) & 63)) {
                                    case 128:
                                        v3 = ("hello" <= "")
                                    case 4294967295:
                                        switch (("a" + "foo bar") + "") {
                                        case "foo bar":
                                            if v2 {
                                                var v137 float32 = 0.1
                                                _ = v137
                                            }
                                        }
                                    }
                                default:
                                    delete(v1, v5)
                                }
                                fmt.Println(v6, v6, v3)
                            }
                            _ = fi109
                        }
                        v3 = v3
                        fmt.Println(v3)
                    }
                }
                v138 := GenS2{F0: v5, F1: (("\n\t" + "") + "foo bar"), F2: -557}
                fmt.Println(v5, v5, v5)
                _ = v138
            }
            v139 := []int{fi104}
            v3 = false
            _ = fi104
            _ = v139
        }
    }
    v5 = (v5 % ((v5) | 1))
    var v140 string = "\n\t"
    fmt.Println(v140, v2)
    fmt.Println(v140, v3, v2)
    var v141 float64 = (float64(0.1) - 1.5)
    switch (-(v5)) {
    case -174:
        v142 := GenS14{F0: (9223372036854775807 == -1), F1: (v4 + (v4 % (((v4 ^ v4)) | 1))), F2: 42}
        _ = v142
    }
    var v143 int8 = 0
    v144 := []int{v5, 165, v5, v5}
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v140:", v140)
    fmt.Println("v141:", v141)
    fmt.Println("v143:", v143)
    fmt.Println("slicelen v144:", len(v144), v144)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v140
    _ = v141
    _ = v143
    _ = v144
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
