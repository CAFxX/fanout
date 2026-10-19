package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int8; F2 int }
type GenS4 struct { F0 string; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64; F1 int64 }
type GenS9 struct { F0 uint64; F1 string }
type GenS10 struct { F0 uint64; F1 int64 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS14 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS15 struct { F0 bool; F1 int8; F2 int }
type GenS16 struct { F0 bool }
func genFn0(p0 uint16, p1 int, p2 int) int64 {
    if (1000 <= 42) {
        v5 := []int{(p2 ^ (p2 + int(p0))), -1, (p1 << ((0) & 63)), (^(p1))}
        for fi6 := 0; fi6 < 3; fi6++ {
            if (false && (!(0 < -1))) {
                for fi7 := 0; fi7 < 14; fi7++ {
                    p2 = fi7
                    var v8 int64 = -9223372036854775808
                    switch (p2 >> ((628) & 63)) {
                    case 4294967295:
                        v9 := GenS1{F0: true}
                        _ = v9
                    default:
                        v10 := GenS2{F0: 42}
                        _ = v10
                    }
                    _ = fi7
                    _ = v8
                }
                v11 := GenS3{F0: -32768, F1: 0, F2: p2}
                if false {
                    if true {
                        for fi12 := 0; fi12 < 16; fi12++ {
                            var v13 int8 = 1
                            var v14 float64 = -2.25
                            v15 := GenS4{F0: "a", F1: 1, F2: v13}
                            _ = fi12
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                    }
                    var v16 uint32 = 0
                    _ = v16
                }
                _ = v11
            }
            var v17 float32 = 1e10
            if true {
                v18 := GenS5{F0: 854, F1: (p0 & (p0 - p0))}
                v17 -= 3.14159
                _ = v18
            }
            _ = fi6
            _ = v17
        }
        _ = v5
    }
    switch "hello" {
    case "hello":
        if false {
            v19 := []int{p1, -9223372036854775808, p2}
            switch p2 {
            case 127:
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
            default:
                var v20 uint8 = 1
                _ = v20
            }
            var v21 uint16 = p0
            _ = v19
            _ = v21
        }
    }
    p1 = 32767
    p0 -= (p0 % ((255) | 1))
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 uint8) uint8 {
    if ((!true) || ((1 > 255) || false)) {
        p0 = (p0 << ((976) & 63))
        if (!false) {
            var v22 uint8 = 127
            switch "a" {
            case "a":
                var v23 int8 = 35
                _ = v23
            case "hello":
                v24 := []int{}
                _ = v24
            case "":
                p0 -= (^(v22))
            }
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 int8 = 100
                _ = fi25
                _ = v26
            }
            _ = v22
        } else {
            switch ("" + ("foo bar" + (("\n\t" + "hello") + (("hello" + ("" + ("foo bar" + "hello"))) + ("\n\t" + ("hello" + "a")))))) {
            case "hello":
                for fi27 := 0; fi27 < 6; fi27++ {
                    v28 := GenS6{F0: genFn0(244, (fi27 - -9223372036854775808), fi27), F1: 405}
                    _ = fi27
                    _ = v28
                }
            case "foo bar":
                v29 := []int{-1, -1, 501, 127}
                _ = v29
            case "\n\t":
                v30 := []int{-9223372036854775808, 32767}
                _ = v30
            default:
                var v31 int = 4294967295
                _ = v31
            }
        }
        p0 -= (p0 >> ((1) & 63))
    } else {
        var v32 uint16 = 255
        p0 = p0
        if true {
            var v33 string = "a"
            var v34 bool = (false || (1 == 2))
            _ = v33
            _ = v34
        } else {
            fmt.Println(v32)
        }
        _ = v32
    }
    for fi35 := 0; fi35 < 15; fi35++ {
        for fi36 := 0; fi36 < 16; fi36++ {
            if (true || false) {
                if (false || (467 <= 0)) {
                    fmt.Println(fi35, fi35)
                    v37 := []int{0, 42}
                    var v38 uint8 = 2
                    _ = v37
                    _ = v38
                } else {
                    p0 += 127
                }
                var v39 uint8 = 156
                _ = v39
            }
            v40 := GenS1{F0: ((-32768 <= -407) && true)}
            var v41 int32 = 1
            _ = fi36
            _ = v40
            _ = v41
        }
        if true {
            p0 += p0
            v42 := []int{(fi35 ^ (fi35 * -1))}
            switch "a" {
            case "a":
                if false {
                    if (genFn0(65535, fi35, (^(fi35))) < genFn0(255, fi35, fi35)) {
                        p0 += p0
                        for fi43 := 0; fi43 < 20; fi43++ {
                            v44 := GenS7{F0: 1}
                            v45 := GenS4{F0: "foo bar", F1: 42, F2: 127}
                            p0 = (p0 + (^(p0)))
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                        var v46 int64 = genFn0(1, (fi35 + 127), 42)
                        _ = v46
                    } else {
                        switch "" {
                        case "foo bar":
                            var v47 uint32 = 4294967295
                            _ = v47
                        case "a":
                            var v48 uint16 = 521
                            _ = v48
                        case "":
                            for fi49 := 0; fi49 < 10; fi49++ {
                                v50 := GenS6{F0: -9223372036854775808, F1: 1}
                                switch "hello" {
                                case "a":
                                    if true {
                                        v51 := GenS3{F0: -32768, F1: -128, F2: (fi35 % (((fi35 * 32767)) | 1))}
                                        var v52 uint16 = 408
                                        switch "hello" {
                                        case "foo bar":
                                            v53 := GenS8{F0: 2, F1: 1}
                                            _ = v53
                                        }
                                        _ = v51
                                        _ = v52
                                    } else {
                                        if (true || true) {
                                            switch fi35 {
                                            case 42:
                                                var v54 uint32 = 2147483647
                                                _ = v54
                                            case 43:
                                                fmt.Println(fi35, fi35)
                                            default:
                                                v55 := GenS2{F0: 139}
                                                _ = v55
                                            }
                                        } else {
                                            fmt.Println(fi49)
                                            _ = len(v42)
                                            if true {
                                                v56 := GenS7{F0: 42}
                                                if true {
                                                    switch ((("\n\t" + ("hello" + "\n\t")) + "a") + "a") {
                                                    case "":
                                                        if false {
                                                            if (2 < 4294967295) {
                                                                var v57 uint8 = p0
                                                                _ = v57
                                                            } else {
                                                                if len(v42) >= 2 {
                                                                    v42 = v42[1:len(v42)]
                                                                }
                                                            }
                                                            var v58 int8 = 13
                                                            _ = v58
                                                        } else {
                                                            switch "" {
                                                            case "hello":
                                                                var v59 string = ("hello" + "a")
                                                                _ = v59
                                                            }
                                                            switch fi35 {
                                                            case -1:
                                                                switch ("a" + (("a" + "") + "hello")) {
                                                                case "":
                                                                    var v60 int = fi49
                                                                    _ = v60
                                                                default:
                                                                    fmt.Println(p0, fi49, p0)
                                                                }
                                                            case 32767:
                                                                p0 = (p0 - (^(p0)))
                                                            default:
                                                                if (true && true) {
                                                                    var v61 uint = 18446744073709551615
                                                                    fmt.Println(fi35)
                                                                    if len(v42) >= 2 {
                                                                        v42 = v42[1:len(v42)]
                                                                    }
                                                                    var v62 bool = (v61 == (v61 % ((v61) | 1)))
                                                                    var v63 uint = v61
                                                                    _ = v61
                                                                    _ = v62
                                                                    _ = v63
                                                                } else {
                                                                    if true {
                                                                        v64 := GenS3{F0: -32768, F1: 127, F2: (fi49 ^ fi49)}
                                                                        _ = v64
                                                                    }
                                                                    _ = len(v42)
                                                                    var v65 bool = true
                                                                    var v66 bool = v65
                                                                    _ = v65
                                                                    _ = v66
                                                                }
                                                            }
                                                        }
                                                        var v67 uint64 = 0
                                                        var v68 int32 = -1
                                                        _ = v67
                                                        _ = v68
                                                    case "hello":
                                                        var v69 int32 = int32(fi49)
                                                        _ = v69
                                                    }
                                                    _ = len(v42)
                                                } else {
                                                    switch (fi49 - 2147483647) {
                                                    case 42:
                                                        if (((((0.1 + (float64(1.5) - 0.1)) / 1.0) >= -2.25) || (!false)) && (1000 != 274)) {
                                                            switch ("a" + "") {
                                                            case "a":
                                                                if (!true) {
                                                                    p0 -= p0
                                                                }
                                                            default:
                                                                var v70 uint8 = 65
                                                                var v71 int32 = -1
                                                                _ = v70
                                                                _ = v71
                                                            }
                                                            switch fi35 {
                                                            case 412:
                                                                var v72 float64 = 0.1
                                                                _ = v72
                                                            case 42:
                                                                var v73 float32 = 0.0
                                                                _ = v73
                                                            case 4294967295:
                                                                var v74 uint8 = (p0 >> ((65535) & 63))
                                                                if (-32768 > 32767) {
                                                                    switch fi49 {
                                                                    case -9223372036854775808:
                                                                        v74 = v74
                                                                    }
                                                                    fmt.Println(fi49)
                                                                    var v75 int = (^(fi35))
                                                                    _ = v75
                                                                } else {
                                                                    var v76 string = ("a" + "foo bar")
                                                                    _ = v76
                                                                }
                                                                _ = v74
                                                            default:
                                                                if (537 > 1) {
                                                                    p0 -= (^(p0))
                                                                    p0 -= p0
                                                                    var v77 int64 = 9223372036854775807
                                                                    _ = v77
                                                                } else {
                                                                    v78 := GenS6{F0: 911, F1: 319}
                                                                    var v79 float64 = 0.0
                                                                    _ = v78
                                                                    _ = v79
                                                                }
                                                            }
                                                            v80 := GenS9{F0: 4294967295, F1: "hello"}
                                                            _ = v80
                                                        } else {
                                                            fmt.Println(fi49)
                                                        }
                                                    case 128:
                                                        fmt.Println(p0)
                                                    default:
                                                        p0 -= (^(p0))
                                                    }
                                                }
                                                _ = v56
                                            } else {
                                                p0 = p0
                                            }
                                        }
                                    }
                                }
                                _ = fi49
                                _ = v50
                            }
                        }
                    }
                    fmt.Println(p0, p0)
                    v81 := GenS10{F0: 42, F1: -325}
                    _ = v81
                }
            default:
                p0 -= (p0 ^ p0)
            }
            _ = v42
        }
        var v82 float32 = 1.5
        _ = fi35
        _ = v82
    }
    for fi83 := 0; fi83 < 17; fi83++ {
        if true {
            fmt.Println(p0, fi83)
            p0 = p0
            fmt.Println(p0, fi83)
        }
        v84 := GenS11{F0: 184}
        if false {
            v85 := []int{(fi83 << ((190) & 63)), fi83, fi83, fi83}
            if true {
                var v86 uint = 4294967295
                for fi87 := 0; fi87 < 7; fi87++ {
                    var v88 uint16 = 348
                    _ = fi87
                    _ = v88
                }
                switch "hello" {
                case "":
                    for fi89 := 0; fi89 < 14; fi89++ {
                        v86 = v86
                        _ = fi89
                    }
                case "a":
                    if (float32(1e10) > (0.0 / 1e10)) {
                        if false {
                            var v90 bool = false
                            for fi91 := 0; fi91 < 4; fi91++ {
                                v92 := GenS12{F0: 0}
                                _ = fi91
                                _ = v92
                            }
                            var v93 uint16 = 65535
                            _ = v90
                            _ = v93
                        } else {
                            if ((!((!true) && ((false || (2147483647 < -877)) || false))) && true) {
                                var v94 string = "\n\t"
                                if true {
                                    for fi95 := 0; fi95 < 11; fi95++ {
                                        v86 = (^(v86))
                                        _ = fi95
                                    }
                                    if true {
                                        var v96 int64 = 9223372036854775807
                                        if len(v85) > 0 {
                                            v85[0] = 2147483647
                                        }
                                        if len(v85) >= 2 {
                                            v85 = v85[1:len(v85)]
                                        }
                                        _ = v96
                                    } else {
                                        v86 = (v86 << (((v86 ^ 65535)) & 63))
                                        fmt.Println(p0, p0, p0)
                                    }
                                } else {
                                    if len(v85) > 0 {
                                        v85[0] = (^(fi83))
                                    }
                                }
                                _ = v94
                            } else {
                                if (!(42 < -128)) {
                                    var v97 uint8 = (p0 >> (((v86 & v86)) & 63))
                                    fmt.Println(p0, fi83, p0)
                                    if true {
                                        if (true && true) {
                                            fmt.Println(v86)
                                            var v98 uint8 = (^(v97))
                                            switch (("" + "") + "a") {
                                            case "\n\t":
                                                v99 := GenS13{F0: 0, F1: 1, F2: 678}
                                                _ = v99
                                            default:
                                                fmt.Println(v86, v86)
                                            }
                                            _ = v98
                                        } else {
                                            v100 := GenS3{F0: 0, F1: -119, F2: fi83}
                                            _ = v100
                                        }
                                        p0 += (v97 - 42)
                                    } else {
                                        switch "a" {
                                        case "hello":
                                            if true {
                                                v101 := GenS5{F0: (v86 << (((v86 + (v86 & v86))) & 63)), F1: 278}
                                                v102 := GenS14{F0: 127, F1: 42, F2: 255}
                                                fmt.Println(p0)
                                                _ = v101
                                                _ = v102
                                            } else {
                                                v97 -= p0
                                            }
                                        }
                                        var v103 uint64 = 112
                                        var v104 string = "hello"
                                        _ = v103
                                        _ = v104
                                    }
                                    _ = v97
                                } else {
                                    fmt.Println(v86, v86)
                                }
                            }
                            var v105 float32 = 0.1
                            if (((v105 / 3.14159) <= v105) && true) {
                                fmt.Println(p0, p0)
                            }
                            _ = v105
                        }
                        _ = len(v85)
                    } else {
                        if ("a" != "") {
                            switch (fi83 + -204) {
                            case -9223372036854775808:
                                for fi106 := 0; fi106 < 5; fi106++ {
                                    switch "" {
                                    case "":
                                        p0 -= 2
                                    case "hello":
                                        if (-1 < 1) {
                                            fmt.Println(fi83, fi106)
                                            var v107 float32 = -2.25
                                            _ = v107
                                        } else {
                                            v86 += 65535
                                            var v108 int = 32768
                                            switch -961 {
                                            case 42:
                                                var v109 uint16 = 32767
                                                _ = v109
                                            case 9223372036854775807:
                                                switch "hello" {
                                                case "a":
                                                    fmt.Println(fi106, v108, fi106)
                                                default:
                                                    switch (v108 & (v108 % ((v108) | 1))) {
                                                    case -9223372036854775808:
                                                        p0 -= 136
                                                    case 32767:
                                                        v86 -= (v86 * v86)
                                                    case 1:
                                                        v110 := GenS9{F0: 2, F1: ("\n\t" + ("" + ("\n\t" + "hello")))}
                                                        _ = v110
                                                    }
                                                }
                                            default:
                                                var v111 float32 = (0.1 + 1.5)
                                                _ = v111
                                            }
                                            _ = v108
                                        }
                                    }
                                    var v112 int64 = 0
                                    _ = fi106
                                    _ = v112
                                }
                            case 582:
                                v86 -= 4294967295
                            case 496:
                                if false {
                                    v113 := GenS15{F0: (69 == 71), F1: -27, F2: (fi83 + fi83)}
                                    _ = v113
                                } else {
                                    for fi114 := 0; fi114 < 17; fi114++ {
                                        if len(v85) >= 2 {
                                            v85 = v85[1:len(v85)]
                                        }
                                        if false {
                                            p0 = (p0 << ((v86) & 63))
                                            switch "" {
                                            case "foo bar":
                                                switch fi114 {
                                                case -2147483648:
                                                    v115 := GenS16{F0: false}
                                                    _ = v115
                                                case 42:
                                                    fmt.Println(fi114, v86)
                                                case 1:
                                                    if len(v85) > 0 {
                                                        v85[0] = (-(fi83))
                                                    }
                                                default:
                                                    switch 2147483647 {
                                                    case 4294967295:
                                                        switch "" {
                                                        case "":
                                                            v85 = append(v85, 2147483647)
                                                        default:
                                                            var v116 float32 = -2.25
                                                            _ = v116
                                                        }
                                                    }
                                                }
                                            case "a":
                                                p0 = p0
                                            default:
                                                switch 404 {
                                                case 128:
                                                    if (false || false) {
                                                        p0 = (p0 >> (((v86 * v86)) & 63))
                                                        var v117 string = ""
                                                        var v118 int16 = 42
                                                        _ = v117
                                                        _ = v118
                                                    } else {
                                                        p0 -= p0
                                                    }
                                                default:
                                                    var v119 string = "\n\t"
                                                    _ = v119
                                                }
                                            }
                                            var v120 uint8 = p0
                                            v120 -= 0
                                            _ = v120
                                        } else {
                                            v86 -= 255
                                        }
                                        _ = fi114
                                    }
                                    if false {
                                        p0 = 127
                                        var v121 int16 = 1
                                        _ = v121
                                    }
                                    v86 -= (v86 / (((v86 - (v86 + v86))) | 1))
                                }
                            default:
                                if true {
                                    p0 = (p0 * (p0 * p0))
                                } else {
                                    v85 = append(v85, fi83)
                                    for fi122 := 0; fi122 < 11; fi122++ {
                                        v86 = (v86 % (((v86 & 0)) | 1))
                                        switch (fi122 | fi122) {
                                        case 9223372036854775807:
                                            var v123 int64 = 1
                                            _ = v123
                                        default:
                                            p0 = p0
                                        }
                                        var v124 string = "\n\t"
                                        var v125 int64 = 1
                                        _ = fi122
                                        _ = v124
                                        _ = v125
                                    }
                                }
                            }
                            var v126 uint32 = 65535
                            for fi127 := 0; fi127 < 17; fi127++ {
                                fmt.Println(v86, v86, p0)
                                var v128 uint32 = v126
                                v129 := GenS3{F0: 0, F1: 1, F2: fi127}
                                _ = fi127
                                _ = v128
                                _ = v129
                            }
                            _ = v126
                        }
                        var v130 int8 = 127
                        p0 -= (p0 / (((p0 % ((181) | 1))) | 1))
                        _ = v130
                    }
                }
                _ = v86
            }
            _ = v85
        }
        _ = fi83
        _ = v84
    }
    v131 := []int{}
    p0 = p0
    _ = p0
    _ = v131
    return (p0 % (((^(p0))) | 1))
}

func main() {
    var v1 int32 = -2147483648
    var v2 int8 = 127
    v3 := make(map[int]uint32)
    var v4 uint16 = 1
    v1 -= 797
    switch 0 {
    case 1:
        _ = len(v3)
    case -542:
        v2 = (v2 / (((v2 ^ -78)) | 1))
    case 42:
        v4 -= 2
    }
    if true {
        v2 = 0
        fmt.Println(v2)
        var v132 int64 = 42
        _ = v132
    } else {
        v4 -= v4
    }
    v2 = v2
    fmt.Println(v1, v4, v4)
    v4 = (^(v4))
    v133 := []int{}
    for fi134 := 0; fi134 < 3; fi134++ {
        if (!true) {
            var v135 string = "foo bar"
            for fi136 := 0; fi136 < 2; fi136++ {
                var v137 int32 = 2147483647
                if len(v133) >= 2 {
                    v133 = v133[1:len(v133)]
                }
                _ = fi136
                _ = v137
            }
            _ = v135
        }
        v138 := GenS15{F0: ("foo bar" <= ("hello" + "")), F1: v2, F2: 2147483647}
        fmt.Println(fi134, fi134, v4)
        _ = fi134
        _ = v138
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v133:", len(v133), v133)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v133
    _ = v2
    _ = v3
    _ = v4
}
