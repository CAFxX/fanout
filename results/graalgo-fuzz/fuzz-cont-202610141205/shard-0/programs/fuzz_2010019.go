package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 int16; F1 uint16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int32; F1 int32; F2 int16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint32; F1 int }
type GenS12 struct { F0 uint32; F1 uint; F2 int16 }
type GenS13 struct { F0 int32; F1 int8; F2 uint64 }
type GenS14 struct { F0 string; F1 uint }
type GenS15 struct { F0 uint; F1 uint64; F2 bool }
type GenS16 struct { F0 uint64; F1 uint }
func genFn0() uint64 {
    var v4 uint32 = 65535
    v4 -= (v4 | (v4 + 229))
    v5 := []int{-37, 32768, -2147483648, -934}
    v4 = 813
    var v6 bool = (false || true)
    _ = v4
    _ = v5
    _ = v6
    return 223
}
func genFn1(p0 uint8, p1 uint16, p2 uint16) uint16 {
    var v7 bool = (4294967295 >= 255)
    if v7 {
        var v8 int64 = 1
        switch "\n\t" {
        case "":
            if (3.14159 >= 3.14159) {
                p1 = p1
                v9 := []int{42, 32767, 1, 211}
                fmt.Println(p1, v8, p0)
                _ = v9
            } else {
                v10 := []int{1, 32767, -278}
                for fi11 := 0; fi11 < 16; fi11++ {
                    var v12 float64 = (((1.5 / 3.14159) - 3.14159) * float64(3.14159))
                    for fi13 := 0; fi13 < 14; fi13++ {
                        var v14 uint32 = 1
                        _ = len(v10)
                        v15 := GenS1{F0: -128, F1: p0, F2: int(v8)}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    var v16 float32 = float32(3.14159)
                    _ = fi11
                    _ = v12
                    _ = v16
                }
                v17 := GenS1{F0: 0, F1: 2, F2: 127}
                _ = v10
                _ = v17
            }
        default:
            p1 = p1
        }
        var v18 uint8 = 255
        _ = v18
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p2
}

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]bool)
    var v3 int = -9223372036854775808
    var v19 int = v3
    var v20 int = (-(v19))
    switch v19 {
    case 0:
        for fi21 := 0; fi21 < 17; fi21++ {
            v22 := []int{9223372036854775807, (fi21 & 32768), v19}
            switch 1 {
            case -36:
                switch 839 {
                case 127:
                    v3 = (^(v20))
                case -9223372036854775808:
                    v19 = v3
                default:
                    var v23 int = (fi21 * 0)
                    _ = v23
                }
            }
            _ = fi21
            _ = v22
        }
    case 676:
        for fi24 := 0; fi24 < 8; fi24++ {
            switch (-(v20)) {
            case 896:
                v25 := GenS2{F0: 1, F1: 367}
                _ = v25
            case -103:
                if true {
                    for fi26 := 0; fi26 < 19; fi26++ {
                        v27 := []int{v19, (fi26 / ((127) | 1))}
                        if true {
                            v28 := GenS3{F0: v1}
                            if true {
                                v1 -= (v1 - v1)
                                v29 := GenS1{F0: 42, F1: 1, F2: (fi24 - -942)}
                                _ = v29
                            } else {
                                fmt.Println(v19)
                            }
                            _ = v28
                        } else {
                            v3 = (fi26 & fi26)
                            v19 -= fi26
                            v30 := GenS1{F0: 1, F1: 69, F2: v3}
                            v1 = 624
                            _ = v30
                        }
                        v3 += -139
                        v3 = (fi24 + v20)
                        _ = fi26
                        _ = v27
                    }
                    switch ("" + ("" + "hello")) {
                    case "hello":
                        var v31 int = v19
                        _ = v31
                    case "":
                        var v32 string = "a"
                        _ = v32
                    case "a":
                        switch v19 {
                        case 128:
                            v3 = (-(fi24))
                        case 2147483647:
                            if true {
                                v33 := GenS4{F0: 119}
                                fmt.Println(v1)
                                _ = v33
                            }
                        case 599:
                            v34 := []int{(v3 << ((18446744073709551615) & 63))}
                            _ = v34
                        }
                    default:
                        v35 := []int{fi24, v20}
                        _ = v35
                    }
                } else {
                    v36 := []int{}
                    _ = v36
                }
            default:
                fmt.Println(v1, v19, v20)
            }
            v3 = v20
            switch ("foo bar" + (("a" + "hello") + (("hello" + (("foo bar" + ("foo bar" + "")) + "foo bar")) + "\n\t"))) {
            case "":
                v37 := GenS1{F0: 0, F1: 0, F2: -2147483648}
                _ = v37
            default:
                if (false && false) {
                    for fi38 := 0; fi38 < 20; fi38++ {
                        v20 -= (v20 | v19)
                        var v39 int = (v3 % ((32768) | 1))
                        if (false && true) {
                            _ = len(v2)
                            fmt.Println(v39, fi38, v3)
                        }
                        _ = fi38
                        _ = v39
                    }
                    fmt.Println(fi24, fi24, v1)
                    if true {
                        switch "a" {
                        case "hello":
                            v40 := GenS5{F0: 55}
                            _ = v40
                        }
                        v41 := GenS3{F0: (v1 >> ((195) & 63))}
                        _ = v41
                    }
                }
            }
            _ = fi24
        }
    case 642:
        v42 := GenS5{F0: 1}
        _ = v42
    default:
        switch ("hello" + "") {
        case "hello":
            var v43 float32 = -2.25
            _ = v43
        default:
            var v44 int32 = 2147483647
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 3; fi45++ {
        _ = len(v2)
        switch ("\n\t" + ("" + ("foo bar" + ""))) {
        case "\n\t":
            if ("" == "") {
                v19 -= (fi45 * (fi45 - v3))
                for fi46 := 0; fi46 < 6; fi46++ {
                    v3 -= (fi46 + (fi46 >> ((1) & 63)))
                    var v47 uint64 = 2
                    v1 += (v1 % (((v1 & 18446744073709551615)) | 1))
                    _ = fi46
                    _ = v47
                }
            } else {
                var v48 bool = false
                delete(v2, 65535)
                for fi49 := 0; fi49 < 12; fi49++ {
                    fmt.Println(v48, v48)
                    switch "hello" {
                    case "":
                        var v50 uint16 = 0
                        _ = v50
                    default:
                        switch ("hello" + "a") {
                        case "\n\t":
                            v51 := v2[2]
                            _ = v51
                        default:
                            v52 := GenS6{F0: -29, F1: fi45}
                            _ = v52
                        }
                    }
                    _ = fi49
                }
                _ = v48
            }
        case "foo bar":
            if false {
                var v53 int32 = -1
                var v54 string = ""
                for fi55 := 0; fi55 < 6; fi55++ {
                    var v56 int32 = v53
                    delete(v2, 212)
                    _ = fi55
                    _ = v56
                }
                _ = v53
                _ = v54
            } else {
                var v57 uint64 = 852
                v20 += (-(v3))
                switch (v20 / (((-(v20))) | 1)) {
                case 127:
                    v58 := GenS1{F0: 0, F1: 42, F2: v19}
                    _ = v58
                }
                _ = v57
            }
        case "":
            v59 := []int{v3, -2147483648, v20}
            _ = v59
        default:
            var v60 int = 4294967295
            _ = v60
        }
        v61 := []int{}
        _ = fi45
        _ = v61
    }
    fmt.Println(v1, v1)
    for fi62 := 0; fi62 < 3; fi62++ {
        if true {
            v63 := GenS7{F0: 1, F1: genFn1(113, 2, 0)}
            fmt.Println(fi62)
            _ = v63
        } else {
            v1 -= (v1 / ((v1) | 1))
        }
        var v64 int = (fi62 * (v20 | 1))
        var v65 int = (v19 + 634)
        _ = fi62
        _ = v64
        _ = v65
    }
    switch (^(v20)) {
    case 32767:
        fmt.Println(v19, v3, v1)
    case 9223372036854775807:
        for fi66 := 0; fi66 < 2; fi66++ {
            if (((float32(0.0) - 0.1) - 0.0) < float32(((1e10 + -2.25) / 0.1))) {
                v3 -= 1
                for fi67 := 0; fi67 < 16; fi67++ {
                    v1 += (v1 % ((713) | 1))
                    switch "" {
                    case "\n\t":
                        v68 := GenS8{F0: -9223372036854775808}
                        _ = v68
                    }
                    _ = fi67
                }
                var v69 uint16 = 887
                _ = v69
            } else {
                for fi70 := 0; fi70 < 14; fi70++ {
                    fmt.Println(v19)
                    switch ("\n\t" + (("hello" + "hello") + (("" + "a") + "hello"))) {
                    case "foo bar":
                        var v71 uint32 = 42
                        _ = v71
                    case "\n\t":
                        var v72 uint8 = uint8(v1)
                        if true {
                            v1 = 0
                            if (v20 >= (fi70 ^ -9223372036854775808)) {
                                var v73 int8 = 42
                                var v74 bool = true
                                v74 = v74
                                _ = v73
                                _ = v74
                            }
                            var v75 uint16 = 255
                            var v76 int64 = 942
                            _ = v75
                            _ = v76
                        } else {
                            v77 := []int{v19, (fi66 * fi70), 9223372036854775807, (v20 & -9223372036854775808)}
                            v1 = v1
                            _ = v77
                        }
                        _ = v72
                    case "hello":
                        switch ("" + "a") {
                        case "hello":
                            if true {
                                fmt.Println(fi70)
                                if (true || false) {
                                    v78 := []int{(v20 | v20), (v3 & 214)}
                                    _ = v78
                                }
                                v79 := GenS9{F0: 1, F1: -521, F2: 42}
                                _ = v79
                            } else {
                                var v80 int = v3
                                if (1.5 == 0.1) {
                                    switch "" {
                                    case "foo bar":
                                        v3 = fi66
                                    case "hello":
                                        if ((v80 | fi70) > (v80 | v80)) {
                                            var v81 float32 = 1e10
                                            v82 := GenS10{F0: 1}
                                            _ = v81
                                            _ = v82
                                        }
                                        var v83 float64 = 0.0
                                        var v84 int8 = -1
                                        _ = v83
                                        _ = v84
                                    case "\n\t":
                                        switch (fi70 - (fi66 & (v20 | v3))) {
                                        case -29:
                                            v3 += (v80 ^ (fi66 * (v20 & v19)))
                                        case -150:
                                            var v85 int16 = 1
                                            _ = v85
                                        }
                                    default:
                                        v86 := GenS11{F0: 859, F1: fi70}
                                        _ = v86
                                    }
                                    if false {
                                        var v87 int8 = 42
                                        switch (v20 ^ fi66) {
                                        case 128:
                                            var v88 uint16 = 255
                                            _ = v88
                                        case 2147483647:
                                            if false {
                                                v19 -= (v3 & (v19 * v3))
                                                v89 := GenS4{F0: 42}
                                                v90 := GenS2{F0: 9223372036854775807, F1: v1}
                                                _ = v89
                                                _ = v90
                                            } else {
                                                if false {
                                                    v91 := GenS12{F0: 42, F1: (^(v1)), F2: 0}
                                                    if true {
                                                        switch fi66 {
                                                        case 948:
                                                            v92 := GenS11{F0: 2147483647, F1: -886}
                                                            _ = v92
                                                        default:
                                                            var v93 int16 = 1000
                                                            _ = v93
                                                        }
                                                        v1 += v1
                                                        if (!true) {
                                                            v94 := GenS2{F0: 0, F1: (v1 ^ v1)}
                                                            var v95 int32 = 0
                                                            v96 := GenS13{F0: v95, F1: 42, F2: 868}
                                                            _ = v94
                                                            _ = v95
                                                            _ = v96
                                                        } else {
                                                            fmt.Println(v80, v1, v3)
                                                            v87 -= (v87 - v87)
                                                            var v97 bool = false
                                                            _ = v97
                                                        }
                                                    }
                                                    _ = v91
                                                } else {
                                                    v98 := []int{-673, fi66, (v20 / (((fi70 | v19)) | 1))}
                                                    v3 = v80
                                                    _ = v98
                                                }
                                                v87 -= 100
                                                var v99 int = fi66
                                                _ = v99
                                            }
                                        case 32767:
                                            var v100 int = 127
                                            _ = v100
                                        }
                                        v101 := GenS2{F0: 9223372036854775807, F1: (v1 & 18446744073709551615)}
                                        _ = v101
                                        _ = v87
                                    } else {
                                        if true {
                                            v102 := []int{fi70}
                                            v3 += 128
                                            switch "a" {
                                            case "":
                                                var v103 int64 = 42
                                                _ = v103
                                            }
                                            _ = v102
                                        } else {
                                            var v104 string = ""
                                            delete(v2, 4294967295)
                                            v1 -= v1
                                            _ = v104
                                        }
                                        if false {
                                            var v105 string = "foo bar"
                                            var v106 int8 = -99
                                            var v107 int16 = 0
                                            _ = v105
                                            _ = v106
                                            _ = v107
                                        }
                                        if true {
                                            fmt.Println(v1, fi70)
                                            var v108 int64 = -561
                                            switch v3 {
                                            case -2147483648:
                                                switch (v20 + int(v108)) {
                                                case 42:
                                                    var v109 int64 = (-(v108))
                                                    _ = v109
                                                default:
                                                    v80 -= (v80 % (((v3 - (v80 % ((-1) | 1)))) | 1))
                                                }
                                            default:
                                                switch v19 {
                                                case 2147483647:
                                                    if true {
                                                        v1 = v1
                                                        if false {
                                                            v110 := []int{v80, (v20 + (v20 ^ fi66))}
                                                            _ = v110
                                                        }
                                                        if false {
                                                            fmt.Println(v108, fi66, v20)
                                                            var v111 float32 = (float32(1e10) - 0.1)
                                                            _ = v111
                                                        }
                                                        v112 := GenS5{F0: 100}
                                                        v19 = 664
                                                        _ = v112
                                                    } else {
                                                        v113 := GenS14{F0: "\n\t", F1: v1}
                                                        var v114 string = "hello"
                                                        _ = v113
                                                        _ = v114
                                                    }
                                                case 42:
                                                    v115 := GenS15{F0: v1, F1: 1, F2: false}
                                                    _ = v115
                                                case 9223372036854775807:
                                                    v116 := GenS9{F0: -140, F1: 2147483647, F2: -1}
                                                    _ = v116
                                                default:
                                                    switch fi70 {
                                                    case 4294967295:
                                                        v20 += v3
                                                    case -2147483648:
                                                        var v117 int32 = 549
                                                        _ = v117
                                                    case 9223372036854775807:
                                                        v1 = (v1 / (((^(v1))) | 1))
                                                    }
                                                }
                                            }
                                            var v118 uint = v1
                                            v19 = 4294967295
                                            v3 -= (v20 * fi70)
                                            _ = v108
                                            _ = v118
                                        } else {
                                            switch ("hello" + (("hello" + "") + "hello")) {
                                            case "":
                                                v119 := GenS1{F0: 0, F1: 255, F2: (v19 % (((fi70 | 86)) | 1))}
                                                _ = v119
                                            case "hello":
                                                var v120 int64 = int64(v80)
                                                _ = v120
                                            case "\n\t":
                                                var v121 int16 = 61
                                                v122 := GenS10{F0: -71}
                                                switch (^(v20)) {
                                                case 32767:
                                                    v123 := []int{}
                                                    _ = v123
                                                case 334:
                                                    var v124 float32 = -2.25
                                                    _ = v124
                                                }
                                                fmt.Println(fi66, v1)
                                                switch fi70 {
                                                case -945:
                                                    var v125 bool = true
                                                    _ = v125
                                                case 741:
                                                    v19 = v20
                                                case 575:
                                                    var v126 int8 = -57
                                                    _ = v126
                                                }
                                                if (-9223372036854775808 == 1) {
                                                    v20 -= v19
                                                    var v127 float64 = -2.25
                                                    v19 = v80
                                                    fmt.Println(v3)
                                                    v128 := []int{v19, 2147483647, 4294967295}
                                                    _ = v127
                                                    _ = v128
                                                } else {
                                                    var v129 bool = false
                                                    _ = v129
                                                }
                                                _ = v121
                                                _ = v122
                                            default:
                                                switch ("" + ((("" + "\n\t") + (((("foo bar" + ("\n\t" + ("" + ("foo bar" + "a")))) + "foo bar") + (("hello" + ("hello" + "hello")) + ("a" + "a"))) + "\n\t")) + "a")) {
                                                case "\n\t":
                                                    switch -348 {
                                                    case -189:
                                                        switch ("\n\t" + "\n\t") {
                                                        case "\n\t":
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v3 = (fi66 & v80)
                                                            default:
                                                                v1 += v1
                                                            }
                                                        case "foo bar":
                                                            v130 := []int{v19, v80, v20, v80}
                                                            _ = v130
                                                        default:
                                                            v19 = 32768
                                                        }
                                                    case 32767:
                                                        var v131 uint = v1
                                                        var v132 float32 = 0.1
                                                        fmt.Println(v1)
                                                        _ = v131
                                                        _ = v132
                                                    }
                                                case "foo bar":
                                                    if (1 > -302) {
                                                        var v133 int64 = -9223372036854775808
                                                        v133 += (v133 + (v133 | 9223372036854775807))
                                                        switch (v19 << (((v1 | (^(v1)))) & 63)) {
                                                        case 32767:
                                                            v20 -= (fi70 << (((v1 | v1)) & 63))
                                                        }
                                                        _ = v133
                                                    } else {
                                                        var v134 uint32 = 65535
                                                        _ = v134
                                                    }
                                                }
                                            }
                                            v135 := GenS12{F0: 593, F1: 255, F2: 1000}
                                            _ = v135
                                        }
                                    }
                                    switch "a" {
                                    case "":
                                        v3 -= (v20 * v3)
                                    }
                                    v136 := v2[562]
                                    v2[2] = false
                                    _ = v136
                                } else {
                                    v20 += v20
                                    v19 = -2147483648
                                    v137 := GenS3{F0: (^(v1))}
                                    _ = v137
                                }
                                _ = v80
                            }
                            v138 := []int{fi70, (fi66 << ((1) & 63))}
                            var v139 int = v19
                            _ = v138
                            _ = v139
                        }
                    default:
                        fmt.Println(fi66, v3, v1)
                    }
                    _ = fi70
                }
            }
            v3 = (v20 | v20)
            var v140 int8 = -128
            _ = fi66
            _ = v140
        }
    }
    fmt.Println(v19, v20)
    for fi141 := 0; fi141 < 18; fi141++ {
        v142 := []int{v19, (fi141 & fi141), v20}
        _ = len(v142)
        _ = fi141
        _ = v142
    }
    for fi143 := 0; fi143 < 16; fi143++ {
        v20 -= (v3 + (v3 * -9223372036854775808))
        var v144 string = "hello"
        v2[210] = true
        _ = fi143
        _ = v144
    }
    var v145 uint = (^(v1))
    v146 := GenS16{F0: 1, F1: v145}
    var v147 int8 = 77
    var v148 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v145:", v145)
    fmt.Println("v147:", v147)
    fmt.Println("v148:", v148)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v145
    _ = v146
    _ = v147
    _ = v148
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
