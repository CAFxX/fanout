package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 string }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS4 struct { F0 string; F1 uint8; F2 uint }
type GenS5 struct { F0 string; F1 uint32; F2 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint32; F1 int8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int16; F1 uint8; F2 string }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int16 }
func genFn0(p0 int, p1 uint8) uint16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            p1 -= (p1 + 0)
            v7 := []int{(fi5 - (^(fi5))), fi6, 299}
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        switch (p0 ^ p0) {
        case 32767:
            switch "\n\t" {
            case "":
                p1 = p1
            case "foo bar":
                switch "foo bar" {
                case "a":
                    p0 -= p0
                case "":
                    switch p0 {
                    case -748:
                        v8 := GenS1{F0: (!(42 > 0)), F1: 2147483647, F2: "hello"}
                        _ = v8
                    case 32767:
                        v9 := GenS2{F0: 83, F1: 32767}
                        _ = v9
                    default:
                        p0 -= (p0 ^ (p0 | -441))
                    }
                }
            case "hello":
                for fi10 := 0; fi10 < 7; fi10++ {
                    switch (^(fi10)) {
                    case 2147483647:
                        var v11 float64 = 1.5
                        _ = v11
                    case 9223372036854775807:
                        fmt.Println(fi10)
                    case 987:
                        v12 := GenS3{F0: (p1 << ((78) & 63)), F1: 1, F2: 4294967295}
                        _ = v12
                    }
                    _ = fi10
                }
            default:
                switch "a" {
                case "a":
                    v13 := []int{}
                    _ = v13
                case "":
                    var v14 bool = true
                    _ = v14
                default:
                    switch (p0 >> ((18446744073709551615) & 63)) {
                    case 4294967295:
                        p0 -= -88
                    case -2147483648:
                        if false {
                            fmt.Println(p1)
                            var v15 float32 = (-2.25 + 0.1)
                            switch p0 {
                            case 32767:
                                if (-1 > -2147483648) {
                                    p0 += p0
                                }
                            case 75:
                                if true {
                                    var v16 uint64 = 4294967295
                                    switch ("\n\t" + "a") {
                                    case "hello":
                                        var v17 uint8 = p1
                                        _ = v17
                                    case "a":
                                        if ((p0 | -9223372036854775808) != -2147483648) {
                                            v18 := GenS1{F0: (!false), F1: 42, F2: "a"}
                                            switch ("hello" + "a") {
                                            case "foo bar":
                                                for fi19 := 0; fi19 < 13; fi19++ {
                                                    v20 := GenS4{F0: ("" + "foo bar"), F1: (p1 >> ((1) & 63)), F2: 65535}
                                                    _ = fi19
                                                    _ = v20
                                                }
                                            case "hello":
                                                v21 := []int{32768}
                                                _ = v21
                                            default:
                                                v16 += 1
                                            }
                                            var v22 uint8 = 255
                                            _ = v18
                                            _ = v22
                                        }
                                    default:
                                        for fi23 := 0; fi23 < 19; fi23++ {
                                            p1 = p1
                                            _ = fi23
                                        }
                                    }
                                    p0 = (p0 + (p0 | p0))
                                    _ = v16
                                } else {
                                    if true {
                                        p1 += (^(p1))
                                        fmt.Println(v15, p1, p0)
                                        p1 = p1
                                    }
                                    if (663 == 834) {
                                        p0 = p0
                                    }
                                    var v24 uint32 = 656
                                    _ = v24
                                }
                            case 42:
                                if true {
                                    if ((!(0 < 32767)) && (-60 < 36)) {
                                        p1 += p1
                                        v25 := []int{}
                                        _ = v25
                                    } else {
                                        v15 += (3.14159 + 0.0)
                                    }
                                } else {
                                    p1 = (p1 | 128)
                                }
                            }
                            _ = v15
                        } else {
                            if false {
                                p0 -= (^(p0))
                                v26 := []int{}
                                _ = v26
                            }
                        }
                    }
                }
            }
        }
    case -9223372036854775808:
        for fi27 := 0; fi27 < 5; fi27++ {
            for fi28 := 0; fi28 < 3; fi28++ {
                v29 := []int{(fi28 + 128), (fi28 & p0), p0}
                _ = fi28
                _ = v29
            }
            v30 := []int{(-(p0)), (p0 & p0), (fi27 / ((127) | 1))}
            var v31 int8 = 100
            _ = fi27
            _ = v30
            _ = v31
        }
    case 0:
        v32 := GenS1{F0: false, F1: 4294967295, F2: "a"}
        _ = v32
    default:
        p0 += (p0 % ((37) | 1))
    }
    v33 := GenS2{F0: 65535, F1: 1}
    var v34 int16 = 0
    _ = p0
    _ = p1
    _ = v33
    _ = v34
    return uint16(p1)
}

func main() {
    var v1 uint16 = 941
    var v2 uint16 = (v1 & v1)
    v3 := []uint32{2147483647, 611, 2147483648}
    var v4 uint32 = 4294967295
    v35 := GenS5{F0: "hello", F1: v4, F2: 1}
    v4 = 2147483647
    var v36 float64 = 3.14159
    if false {
        var v37 int = 2147483647
        var v38 bool = (uint16(v4) >= v1)
        _ = v37
        _ = v38
    } else {
        if (76 >= 65) {
            if (false || false) {
                v4 -= 2147483648
                v36 += v36
                var v39 uint32 = v4
                _ = v39
            } else {
                if (!(42 >= 2147483647)) {
                    v1 = genFn0(204, 233)
                }
                for fi40 := 0; fi40 < 3; fi40++ {
                    fmt.Println(v4, v1, fi40)
                    v41 := GenS4{F0: ("\n\t" + "a"), F1: 128, F2: 42}
                    v42 := GenS6{F0: 255}
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            }
            switch "\n\t" {
            case "a":
                var v43 string = "foo bar"
                _ = v43
            default:
                fmt.Println(v36, v4, v4)
            }
        }
    }
    for fi44 := 0; fi44 < 1; fi44++ {
        if len(v3) > 0 {
            v3[0] = (^(v4))
        }
        for fi45 := 0; fi45 < 6; fi45++ {
            var v46 int = (fi45 >> ((4294967295) & 63))
            v47 := GenS1{F0: false, F1: 0, F2: "\n\t"}
            v2 = (v1 & (v1 << ((18446744073709551615) & 63)))
            var v48 string = "hello"
            v36 = (v36 + ((v36 / 0.1) - float64(((3.14159 + 1e10) * 0.0))))
            if true {
                v49 := GenS4{F0: "a", F1: 1, F2: 0}
                var v50 uint = 184
                _ = v49
                _ = v50
            } else {
                v4 = v4
            }
            switch (v48 + ((v48 + ((v48 + (v48 + (v48 + "hello"))) + ("foo bar" + v48))) + ((v48 + ("hello" + (v48 + ("foo bar" + v48)))) + (v48 + "\n\t")))) {
            case "foo bar":
                v46 = 4294967295
            }
            _ = fi45
            _ = v46
            _ = v47
            _ = v48
        }
        if (1e10 == (float32(v36) / 1.0)) {
            v36 -= v36
            fmt.Println(v2, v4, v36)
            for fi51 := 0; fi51 < 6; fi51++ {
                var v52 int = (fi51 << ((92) & 63))
                v2 = genFn0((v52 ^ 32768), 2)
                _ = fi51
                _ = v52
            }
        } else {
            v3 = append(v3, 2)
            if false {
                if false {
                    _ = len(v3)
                    if true {
                        if true {
                            var v53 float32 = 0.1
                            switch "" {
                            case "foo bar":
                                switch fi44 {
                                case 9223372036854775807:
                                    for fi54 := 0; fi54 < 13; fi54++ {
                                        switch fi44 {
                                        case 1:
                                            fmt.Println(v36, v36, fi54)
                                        case 693:
                                            if (490 != v4) {
                                                if len(v3) > 0 {
                                                    v3[0] = v4
                                                }
                                                if true {
                                                    v55 := GenS1{F0: false, F1: 42, F2: "foo bar"}
                                                    _ = v55
                                                }
                                                var v56 int8 = 42
                                                switch ("\n\t" + "a") {
                                                case "a":
                                                    switch (fi44 / ((fi54) | 1)) {
                                                    case 9223372036854775807:
                                                        var v57 uint32 = (v4 - (v4 % (((v4 << ((700) & 63))) | 1)))
                                                        _ = v57
                                                    default:
                                                        v4 = 0
                                                    }
                                                default:
                                                    fmt.Println(v53)
                                                }
                                                fmt.Println(fi54, v2)
                                                _ = v56
                                            }
                                        case -426:
                                            v3 = append(v3, 2)
                                        default:
                                            v36 = (v36 + ((1.5 - v36) / (float64((v53 - 3.14159)) - 1e10)))
                                        }
                                        v4 = v4
                                        var v58 int = fi44
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v59 int8 = 41
                                        fmt.Println(v2, v2, v4)
                                        _ = fi54
                                        _ = v58
                                        _ = v59
                                    }
                                }
                            case "hello":
                                var v60 uint32 = (v4 * 42)
                                _ = v60
                            case "":
                                v61 := GenS3{F0: 0, F1: 255, F2: 1}
                                _ = v61
                            default:
                                var v62 uint16 = v2
                                _ = v62
                            }
                            fmt.Println(fi44, v53, v2)
                            _ = v53
                        } else {
                            for fi63 := 0; fi63 < 13; fi63++ {
                                fmt.Println(v1)
                                if false {
                                    v64 := GenS4{F0: "\n\t", F1: uint8(v2), F2: 358}
                                    v1 = v1
                                    v65 := GenS7{F0: 18446744073709551615}
                                    _ = v64
                                    _ = v65
                                }
                                _ = fi63
                            }
                            if false {
                                for fi66 := 0; fi66 < 7; fi66++ {
                                    if (-1 >= -115) {
                                        v4 -= v4
                                        var v67 float64 = v36
                                        if false {
                                            v36 = 1.5
                                            v68 := GenS2{F0: 2147483647, F1: (v1 << ((881) & 63))}
                                            var v69 uint64 = 759
                                            _ = v68
                                            _ = v69
                                        }
                                        v36 += (3.14159 + v67)
                                        var v70 uint = 1
                                        _ = v67
                                        _ = v70
                                    } else {
                                        v2 += (v2 ^ 65535)
                                        v36 += (0.1 / (float64(0.1) + (0.0 + v36)))
                                    }
                                    switch fi44 {
                                    case 764:
                                        v3 = append(v3, (^(v4)))
                                    case 619:
                                        v36 -= (-2.25 + float64(((((0.1 + 1e10) + 1e10) + 0.0) - 0.0)))
                                    default:
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(fi66, v4)
                                        default:
                                            var v71 int = fi44
                                            _ = v71
                                        }
                                        switch 0 {
                                        case 32767:
                                            v36 = v36
                                        case 4294967295:
                                            var v72 int32 = -1
                                            _ = v72
                                        case 2147483647:
                                            v73 := GenS8{F0: (!(0.0 >= 1e10))}
                                            _ = v73
                                        default:
                                            v74 := GenS8{F0: false}
                                            _ = v74
                                        }
                                    }
                                    if (true && false) {
                                        fmt.Println(fi44)
                                        v3 = append(v3, 2147483647)
                                    } else {
                                        var v75 bool = true
                                        switch ("\n\t" + "hello") {
                                        case "hello":
                                            if false {
                                                v76 := GenS9{F0: 27, F1: -128}
                                                _ = v76
                                            } else {
                                                if false {
                                                    v3 = append(v3, uint32(v1))
                                                    var v77 int16 = 65
                                                    var v78 uint16 = v2
                                                    _ = v77
                                                    _ = v78
                                                }
                                                v79 := GenS8{F0: (((v36 * (v36 * float64(0.0))) - (v36 * (((v36 - (v36 * 1.5)) * float64(0.0)) + (v36 - (v36 + -2.25))))) >= ((0.0 - (float64(float32(1.5)) * v36)) * -2.25))}
                                                _ = v79
                                            }
                                        }
                                        if (v36 >= ((1.5 / ((0.0 / v36) / v36)) * -2.25)) {
                                            v80 := GenS3{F0: 128, F1: 127, F2: 4294967295}
                                            switch -318 {
                                            case 9223372036854775807:
                                                switch (("" + "") + ("a" + "\n\t")) {
                                                case "":
                                                    fmt.Println(v36, v4, v1)
                                                default:
                                                    _ = len(v3)
                                                }
                                            default:
                                                v75 = true
                                            }
                                            _ = v80
                                        } else {
                                            v36 -= v36
                                            var v81 uint32 = v4
                                            var v82 int16 = 32767
                                            _ = v81
                                            _ = v82
                                        }
                                        v36 -= float64(0.0)
                                        _ = v75
                                    }
                                    _ = fi66
                                }
                                v4 = (v4 & v4)
                                v4 += v4
                            } else {
                                var v83 int64 = -71
                                switch fi44 {
                                case 42:
                                    fmt.Println(v1, v83, v1)
                                }
                                _ = v83
                            }
                        }
                        var v84 string = ("a" + "hello")
                        v1 += 132
                        _ = v84
                    } else {
                        for fi85 := 0; fi85 < 18; fi85++ {
                            if true {
                                v3 = append(v3, (v4 << ((425) & 63)))
                                var v86 int32 = -1
                                _ = v86
                            } else {
                                switch -9223372036854775808 {
                                case -184:
                                    v4 += 1
                                case 1:
                                    var v87 int8 = 0
                                    _ = v87
                                }
                            }
                            _ = fi85
                        }
                        if ("hello" > "") {
                            v88 := GenS3{F0: 127, F1: 128, F2: 42}
                            var v89 uint32 = 330
                            _ = v88
                            _ = v89
                        } else {
                            for fi90 := 0; fi90 < 1; fi90++ {
                                var v91 int16 = 148
                                _ = fi90
                                _ = v91
                            }
                            switch (("a" + "") + ("\n\t" + ("foo bar" + "a"))) {
                            case "":
                                var v92 int16 = 1000
                                _ = v92
                            case "hello":
                                var v93 bool = false
                                _ = v93
                            default:
                                v94 := GenS1{F0: true, F1: v4, F2: ("a" + ("foo bar" + "a"))}
                                _ = v94
                            }
                        }
                    }
                }
                if len(v3) > 0 {
                    v3[0] = (^(v4))
                }
            } else {
                for fi95 := 0; fi95 < 14; fi95++ {
                    v96 := GenS9{F0: v4, F1: int8(fi95)}
                    switch ("hello" + "hello") {
                    case "foo bar":
                        var v97 string = "foo bar"
                        _ = v97
                    default:
                        if (!false) {
                            if true {
                                var v98 uint16 = 2
                                var v99 uint32 = v4
                                _ = v98
                                _ = v99
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v100 float32 = 0.0
                            _ = v100
                        } else {
                            v4 = 513
                            if len(v3) > 0 {
                                v3[0] = (v4 % ((v4) | 1))
                            }
                            v36 = float64(3.14159)
                        }
                    }
                    _ = fi95
                    _ = v96
                }
            }
        }
        _ = fi44
    }
    switch 528 {
    case 1:
        var v101 int64 = -959
        _ = v101
    case 2147483647:
        if (true || true) {
            v102 := GenS5{F0: ("\n\t" + "foo bar"), F1: (v4 - uint32(v1)), F2: 1}
            _ = v102
        } else {
            _ = len(v3)
        }
    case -9223372036854775808:
        var v103 float32 = 3.14159
        _ = v103
    default:
        if (("hello" <= ("" + "")) && false) {
            for fi104 := 0; fi104 < 11; fi104++ {
                var v105 uint = 812
                _ = fi104
                _ = v105
            }
            for fi106 := 0; fi106 < 15; fi106++ {
                switch "hello" {
                case "":
                    for fi107 := 0; fi107 < 11; fi107++ {
                        switch "" {
                        case "\n\t":
                            _ = len(v3)
                            v108 := GenS10{F0: 0}
                            fmt.Println(v2, v2, v36)
                            _ = v108
                        case "":
                            if false {
                                fmt.Println(v4, v2)
                                var v109 int32 = 784
                                v110 := GenS10{F0: 42}
                                _ = v109
                                _ = v110
                            }
                        default:
                            var v111 uint32 = (v4 / ((1) | 1))
                            _ = v111
                        }
                        _ = fi107
                    }
                case "foo bar":
                    _ = len(v3)
                case "hello":
                    var v112 uint16 = (v1 << ((893) & 63))
                    _ = v112
                default:
                    var v113 uint64 = 0
                    _ = v113
                }
                v36 -= (v36 * v36)
                var v114 int64 = 9223372036854775807
                _ = fi106
                _ = v114
            }
        } else {
            var v115 float32 = -2.25
            _ = len(v3)
            var v116 uint32 = v4
            _ = v115
            _ = v116
        }
    }
    switch 32767 {
    case 635:
        if (false || (!false)) {
            if true {
                if false {
                    v1 = v1
                } else {
                    if (-43 >= 122) {
                        var v117 uint64 = 42
                        v118 := GenS11{F0: 1000, F1: 224, F2: "hello"}
                        _ = v117
                        _ = v118
                    } else {
                        var v119 int64 = 842
                        if true {
                            v2 = 32767
                        } else {
                            for fi120 := 0; fi120 < 14; fi120++ {
                                v4 = 4294967295
                                _ = fi120
                            }
                        }
                        _ = v119
                    }
                    for fi121 := 0; fi121 < 9; fi121++ {
                        fmt.Println(v2)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi121
                    }
                }
                if true {
                    switch ("\n\t" + "") {
                    case "":
                        var v122 uint32 = 42
                        _ = v122
                    }
                    v2 = 68
                } else {
                    v4 -= v4
                    switch -151 {
                    case 42:
                        fmt.Println(v2, v1)
                    default:
                        var v123 uint64 = 47
                        _ = v123
                    }
                }
                fmt.Println(v4)
            } else {
                v124 := GenS11{F0: -32768, F1: 255, F2: "a"}
                _ = v124
            }
        } else {
            v2 = v2
            v125 := GenS1{F0: ((true && false) || false), F1: v4, F2: "hello"}
            _ = v125
        }
    default:
        if false {
            switch "foo bar" {
            case "":
                for fi126 := 0; fi126 < 3; fi126++ {
                    fmt.Println(v2, v2)
                    for fi127 := 0; fi127 < 14; fi127++ {
                        if false {
                            if false {
                                _ = len(v3)
                                if true {
                                    v2 = 255
                                    var v128 int = fi127
                                    var v129 int64 = 357
                                    _ = v128
                                    _ = v129
                                } else {
                                    var v130 int16 = 1000
                                    v36 -= (-2.25 - float64(((1.5 - float32(v36)) * 3.14159)))
                                    _ = v130
                                }
                            }
                            v131 := GenS1{F0: true, F1: v4, F2: "foo bar"}
                            switch fi126 {
                            case 2147483647:
                                switch ("foo bar" + "\n\t") {
                                case "":
                                    var v132 int16 = -32768
                                    _ = v132
                                case "hello":
                                    if true {
                                        var v133 uint = 18446744073709551615
                                        fmt.Println(v2, v2)
                                        _ = v133
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = v4
                                        }
                                        v36 = ((float64((3.14159 * (3.14159 - 3.14159))) / float64(3.14159)) * v36)
                                    }
                                case "\n\t":
                                    var v134 uint32 = v4
                                    _ = v134
                                default:
                                    if true {
                                        var v135 int8 = -128
                                        var v136 uint64 = 4294967295
                                        _ = v135
                                        _ = v136
                                    } else {
                                        v2 = uint16(v4)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v137 int16 = 873
                                        _ = v137
                                    }
                                }
                            default:
                                fmt.Println(fi126)
                            }
                            _ = v131
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if true {
                                if (0 < 127) {
                                    if (-1 == -176) {
                                        v4 -= 675
                                    } else {
                                        _ = len(v3)
                                        v138 := GenS9{F0: (^(v4)), F1: 1}
                                        _ = v138
                                    }
                                    var v139 int8 = 1
                                    v1 = uint16(v4)
                                    if false {
                                        var v140 bool = true
                                        var v141 int8 = (^(v139))
                                        var v142 float64 = 0.1
                                        var v143 string = ("" + "")
                                        switch v143 {
                                        case "\n\t":
                                            var v144 int64 = 0
                                            _ = v144
                                        default:
                                            switch v143 {
                                            case "foo bar":
                                                var v145 uint16 = (v1 + (^(v1)))
                                                _ = v145
                                            case "a":
                                                v143 = v143
                                                fmt.Println(v140, v143, fi127)
                                            }
                                        }
                                        var v146 float32 = (0.0 + ((((1e10 / 1e10) * 1e10) - float32(float64(-2.25))) - ((-2.25 * (float32(0.0) / 1.5)) / (0.1 - -2.25))))
                                        _ = v140
                                        _ = v141
                                        _ = v142
                                        _ = v143
                                        _ = v146
                                    }
                                    var v147 float32 = 0.1
                                    _ = v139
                                    _ = v147
                                }
                            }
                            switch (fi127 - fi126) {
                            case -2147483648:
                                switch (fi126 | fi126) {
                                case 32768:
                                    switch (((("hello" + "foo bar") + "hello") + (("" + ("hello" + ("" + "foo bar"))) + ("foo bar" + "foo bar"))) + ("" + "a")) {
                                    case "a":
                                        v36 = (v36 - ((1.5 / v36) + 3.14159))
                                    case "\n\t":
                                        if true {
                                            v148 := GenS11{F0: 32767, F1: 255, F2: "a"}
                                            v3 = append(v3, (v4 / ((v4) | 1)))
                                            switch (fi126 & (fi127 / (((fi126 / ((fi126) | 1))) | 1))) {
                                            case 9223372036854775807:
                                                if true {
                                                    switch fi127 {
                                                    case -1:
                                                        if false {
                                                            v1 = (v2 * (v1 | (v2 | v2)))
                                                            fmt.Println(fi127, fi127)
                                                            switch (fi126 - 9223372036854775807) {
                                                            case -9223372036854775808:
                                                                if (false || ((fi127 - fi127) < fi127)) {
                                                                    v1 += v1
                                                                    v149 := GenS11{F0: 32767, F1: 2, F2: ("hello" + "hello")}
                                                                    _ = v149
                                                                }
                                                            case 2147483647:
                                                                v36 = ((-2.25 + v36) / (0.0 + 1e10))
                                                            case 9223372036854775807:
                                                                v1 = genFn0(fi127, 80)
                                                            }
                                                            v150 := GenS2{F0: (v4 >> ((0) & 63)), F1: v2}
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = len(v3)
                                                            _ = v150
                                                        } else {
                                                            var v151 int32 = 2147483647
                                                            v4 -= v4
                                                            _ = v151
                                                        }
                                                    case 127:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    default:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                    v2 -= (v1 << ((65535) & 63))
                                                    var v152 uint = 629
                                                    v36 += (float64(0.1) * v36)
                                                    _ = v152
                                                } else {
                                                    switch (fi126 >> ((2) & 63)) {
                                                    case 128:
                                                        v2 += (v2 >> ((255) & 63))
                                                    default:
                                                        v153 := GenS3{F0: 255, F1: 128, F2: 71}
                                                        _ = v153
                                                    }
                                                    v154 := GenS3{F0: 124, F1: 0, F2: 18446744073709551615}
                                                    _ = v154
                                                }
                                            default:
                                                v4 += v4
                                            }
                                            _ = v148
                                        } else {
                                            fmt.Println(fi127, v1)
                                        }
                                    case "foo bar":
                                        if (!true) {
                                            fmt.Println(fi126, fi127)
                                            v1 += 1
                                            switch "" {
                                            case "":
                                                var v155 uint8 = 33
                                                _ = v155
                                            case "a":
                                                v156 := GenS12{F0: "hello"}
                                                _ = v156
                                            }
                                        }
                                    }
                                }
                            case -9223372036854775808:
                                v157 := GenS6{F0: 81}
                                _ = v157
                            case 919:
                                var v158 int64 = -701
                                _ = v158
                            }
                        }
                        _ = fi127
                    }
                    _ = fi126
                }
            }
            switch ("" + ("\n\t" + "")) {
            case "\n\t":
                if true {
                    var v159 int16 = -659
                    if false {
                        if false {
                            var v160 uint16 = v2
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v160
                        } else {
                            var v161 int32 = int32(v4)
                            v162 := GenS13{F0: 1}
                            _ = v161
                            _ = v162
                        }
                        var v163 uint64 = 927
                        _ = v163
                    }
                    if (0 >= 18446744073709551615) {
                        if true {
                            v36 += v36
                        } else {
                            for fi164 := 0; fi164 < 2; fi164++ {
                                var v165 float64 = -2.25
                                if (-2147483648 > -561) {
                                    var v166 uint32 = (v4 * uint32(v2))
                                    _ = v166
                                } else {
                                    v36 = ((((v165 * (v36 - 1e10)) / ((v36 / 1.0) / v36)) / v165) + v165)
                                    v167 := GenS7{F0: 1}
                                    _ = v167
                                }
                                _ = fi164
                                _ = v165
                            }
                        }
                        v4 = v4
                    } else {
                        for fi168 := 0; fi168 < 3; fi168++ {
                            var v169 uint32 = 447
                            v36 = v36
                            _ = fi168
                            _ = v169
                        }
                        switch 32768 {
                        case 127:
                            v1 = (v1 >> ((255) & 63))
                        case 128:
                            v170 := GenS8{F0: (!(float64(1e10) == 0.1))}
                            _ = v170
                        case -154:
                            if (!false) {
                                switch 467 {
                                case 318:
                                    if len(v3) > 0 {
                                        v3[0] = 0
                                    }
                                case 2147483647:
                                    for fi171 := 0; fi171 < 5; fi171++ {
                                        for fi172 := 0; fi172 < 12; fi172++ {
                                            fmt.Println(v159)
                                            v3 = append(v3, (v4 + v4))
                                            _ = fi172
                                        }
                                        for fi173 := 0; fi173 < 4; fi173++ {
                                            if false {
                                                var v174 bool = (((v36 / v36) - v36) == v36)
                                                _ = v174
                                            }
                                            var v175 uint16 = v2
                                            var v176 bool = true
                                            _ = fi173
                                            _ = v175
                                            _ = v176
                                        }
                                        _ = fi171
                                    }
                                case 128:
                                    fmt.Println(v36, v4, v1)
                                default:
                                    var v177 int16 = 1000
                                    _ = v177
                                }
                                var v178 uint64 = 42
                                fmt.Println(v1, v159)
                                _ = v178
                            }
                        default:
                            fmt.Println(v36, v36)
                        }
                    }
                    _ = v159
                }
            }
        }
    }
    if ((1000 != 1000) && false) {
        v2 = (v2 / ((genFn0(-582, 1)) | 1))
        var v179 uint = 18446744073709551615
        _ = v179
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    switch ("" + "\n\t") {
    case "hello":
        var v180 int32 = 42
        _ = v180
    default:
        v1 = (v2 / (((v1 / ((255) | 1))) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
