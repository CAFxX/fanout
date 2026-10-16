package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int16; F2 string }
type GenS3 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint32; F1 uint64 }
type GenS9 struct { F0 int16; F1 uint32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int64; F1 uint16 }
type GenS12 struct { F0 int32; F1 bool; F2 string }
type GenS13 struct { F0 int32; F1 uint8; F2 bool }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int64; F1 uint32 }
type GenS17 struct { F0 int; F1 uint16; F2 uint }
type GenS18 struct { F0 int8 }
func genFn0(p0 int64) int {
    v6 := GenS1{F0: -129}
    p0 = p0
    var v7 uint64 = 2
    if ((!false) && true) {
        var v8 int32 = 0
        v7 -= (v7 ^ v7)
        _ = v8
    }
    _ = p0
    _ = v6
    _ = v7
    return 408
}
func genFn1(p0 string) int8 {
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS2{F0: 139, F1: 1000, F2: p0}
        if true {
            var v11 int = -9223372036854775808
            for fi12 := 0; fi12 < 11; fi12++ {
                v13 := []int{371, 490}
                v14 := GenS3{F0: 229, F1: int16(fi9), F2: 18446744073709551615}
                v15 := GenS4{F0: -9223372036854775808}
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            p0 = p0
            _ = v11
        }
        v16 := GenS5{F0: 551}
        _ = fi9
        _ = v10
        _ = v16
    }
    if true {
        var v17 int32 = 0
        for fi18 := 0; fi18 < 3; fi18++ {
            v19 := []int{fi18, fi18, fi18}
            if ("" <= p0) {
                v20 := GenS5{F0: 42}
                switch p0 {
                case "":
                    fmt.Println(p0)
                case "hello":
                    var v21 bool = (18446744073709551615 >= 42)
                    _ = v21
                case "a":
                    if (true || (4294967295 != 42)) {
                        v17 += -233
                    }
                default:
                    fmt.Println(p0, fi18, v17)
                }
                _ = v20
            }
            v19 = append(v19, fi18)
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    p0 = ""
    for fi22 := 0; fi22 < 4; fi22++ {
        v23 := GenS4{F0: -1}
        v24 := []int{}
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := GenS6{F0: 2147483647, F1: 42, F2: 1}
    _ = p0
    _ = v25
    return -99
}

func main() {
    var v1 int = 42
    v2 := make(map[uint32]int)
    var v3 string = ("" + "")
    v4 := make(map[string]int)
    var v5 int8 = 127
    if false {
        v26 := GenS1{F0: 0}
        if true {
            v27 := GenS6{F0: 1, F1: 4294967295, F2: 703}
            v5 += (v5 - v5)
            if false {
                fmt.Println(v3)
                var v28 string = ("foo bar" + (v3 + ""))
                var v29 float32 = 1e10
                _ = v28
                _ = v29
            }
            _ = v27
        } else {
            var v30 uint8 = 128
            _ = v30
        }
        v31 := GenS5{F0: 42}
        _ = v26
        _ = v31
    }
    var v32 uint = 1
    fmt.Println(v3, v5)
    var v33 uint64 = 415
    v3 = v3
    switch v1 {
    case 9223372036854775807:
        if (1000 == int16(v5)) {
            v1 = -379
            var v34 string = v3
            _ = v34
        }
    }
    var v35 int64 = -1
    if false {
        v36 := GenS4{F0: v35}
        v37 := GenS6{F0: 42, F1: (v32 % ((775) | 1)), F2: 0}
        _ = v36
        _ = v37
    } else {
        var v38 int8 = (v5 * (v5 + v5))
        _ = v38
    }
    for fi39 := 0; fi39 < 15; fi39++ {
        if (false || ((((v32 << ((v32) & 63)) > (^(v32))) || false) || (2 <= uint16(v33)))) {
            switch "\n\t" {
            case "":
                for fi40 := 0; fi40 < 20; fi40++ {
                    v41 := GenS7{F0: (v33 | v33)}
                    _ = fi40
                    _ = v41
                }
            case "\n\t":
                var v42 int = (fi39 ^ genFn0(int64(v32)))
                _ = v42
            default:
                v1 += v1
            }
            switch (v1 & genFn0(1)) {
            case 2147483647:
                var v43 bool = true
                _ = v43
            case -9223372036854775808:
                var v44 uint16 = 42
                _ = v44
            case 127:
                var v45 int8 = v5
                _ = v45
            default:
                if (!(127 == 42)) {
                    v46 := v2[2147483647]
                    _ = v46
                } else {
                    v32 -= v32
                    if (((uint16(v32) != 32767) || (!false)) || true) {
                        var v47 string = v3
                        _ = v47
                    } else {
                        switch (v3 + v3) {
                        case "a":
                            v5 -= 36
                        }
                    }
                }
            }
            v48 := GenS6{F0: 4294967295, F1: 4294967295, F2: 929}
            _ = v48
        } else {
            var v49 uint32 = 42
            switch fi39 {
            case 2147483647:
                v5 += v5
            case 127:
                if ((^(v32)) == v32) {
                    var v50 int16 = 252
                    if true {
                        if false {
                            for fi51 := 0; fi51 < 3; fi51++ {
                                if (!true) {
                                    var v52 float64 = float64(0.0)
                                    _ = v52
                                } else {
                                    if false {
                                        fmt.Println(v33, v33, v1)
                                        v53 := []int{fi39, genFn0(-24)}
                                        v49 = 1
                                        _ = len(v2)
                                        _ = v53
                                    } else {
                                        var v54 int16 = v50
                                        _ = v54
                                    }
                                    var v55 int = genFn0(v35)
                                    if false {
                                        v56 := GenS8{F0: (v49 - 0), F1: 42}
                                        var v57 bool = false
                                        _ = v56
                                        _ = v57
                                    }
                                    v49 -= (v49 << (((v32 | 59)) & 63))
                                    fmt.Println(fi51, fi51, fi51)
                                    v1 = int(v33)
                                    v58 := GenS2{F0: 1, F1: (v50 | v50), F2: v3}
                                    var v59 uint32 = 2147483648
                                    _ = v55
                                    _ = v58
                                    _ = v59
                                }
                                v3 = v3
                                _ = fi51
                            }
                        } else {
                            v60 := GenS8{F0: 42, F1: 356}
                            _ = v60
                        }
                        for fi61 := 0; fi61 < 2; fi61++ {
                            v35 += v35
                            _ = fi61
                        }
                        v62 := GenS9{F0: (v50 / (((-(v50))) | 1)), F1: v49}
                        _ = v62
                    }
                    fmt.Println(v1)
                    _ = v50
                } else {
                    var v63 int64 = int64(v5)
                    v64 := v4[v3]
                    v33 -= (v33 * (v33 * v33))
                    _ = v63
                    _ = v64
                }
            case 895:
                switch v3 {
                case "foo bar":
                    v65 := GenS9{F0: 0, F1: (v49 >> ((v32) & 63))}
                    _ = v65
                case "hello":
                    var v66 int32 = 619
                    _ = v66
                case "a":
                    v4["\n\t"] = -1
                default:
                    v3 = v3
                }
            default:
                switch ("foo bar" + "") {
                case "":
                    v32 = v32
                case "foo bar":
                    var v67 int32 = -194
                    _ = v67
                case "a":
                    v32 -= (^(v32))
                default:
                    for fi68 := 0; fi68 < 4; fi68++ {
                        var v69 float32 = 1e10
                        if false {
                            if (uint16(v33) > 32767) {
                                if ((v35 / (((v35 * -9223372036854775808)) | 1)) <= v35) {
                                    v1 = genFn0(0)
                                    var v70 uint32 = 4294967295
                                    _ = v70
                                } else {
                                    switch genFn0(1) {
                                    case 0:
                                        v2[v49] = fi68
                                    case 4294967295:
                                        var v71 string = v3
                                        _ = v71
                                    case 1:
                                        v72 := GenS10{F0: false}
                                        _ = v72
                                    default:
                                        var v73 float64 = float64(v69)
                                        _ = v73
                                    }
                                }
                                v32 -= (v32 % (((v32 ^ 1)) | 1))
                            } else {
                                v2[v49] = (v1 * (^(fi39)))
                                v74 := []int{(fi39 ^ 32768), fi68}
                                if true {
                                    if true {
                                        var v75 float32 = v69
                                        _ = v75
                                    } else {
                                        switch v3 {
                                        case "":
                                            v5 += v5
                                        case "a":
                                            var v76 uint32 = 99
                                            _ = v76
                                        case "hello":
                                            _ = len(v2)
                                        default:
                                            var v77 int64 = 1
                                            _ = v77
                                        }
                                    }
                                    v78 := v4[(v3 + "\n\t")]
                                    if true {
                                        v74 = append(v74, int(v35))
                                        v79 := GenS10{F0: (-1 < -784)}
                                        if (!(!true)) {
                                            fmt.Println(v33)
                                            var v80 float64 = (3.14159 / 1.5)
                                            _ = v80
                                        } else {
                                            v81 := GenS5{F0: v32}
                                            if (1 == 555) {
                                                if ((!false) || false) {
                                                    if false {
                                                        v82 := GenS11{F0: 0, F1: 2}
                                                        var v83 uint8 = 2
                                                        v69 -= 1.5
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                }
                                                var v84 uint32 = v49
                                                _ = v84
                                            } else {
                                                v85 := GenS12{F0: -826, F1: ((1000 == -1) || ((v32 & (v32 & v32)) <= 270)), F2: v3}
                                                _ = len(v4)
                                                _ = v85
                                            }
                                            var v86 uint32 = uint32(v33)
                                            _ = v81
                                            _ = v86
                                        }
                                        _ = v79
                                    }
                                    fmt.Println(v78, v32)
                                    switch v3 {
                                    case "a":
                                        v87 := GenS7{F0: v33}
                                        _ = v87
                                    case "\n\t":
                                        var v88 uint16 = 653
                                        _ = v88
                                    default:
                                        v33 += v33
                                    }
                                    switch "a" {
                                    case "hello":
                                        var v89 uint = 933
                                        _ = v89
                                    }
                                    _ = v78
                                } else {
                                    v5 = 100
                                    v90 := GenS6{F0: v49, F1: (v32 << ((v32) & 63)), F2: (v49 / (((v49 + v49)) | 1))}
                                    _ = v90
                                }
                                _ = v74
                            }
                            var v91 uint32 = v49
                            v49 += (v91 ^ (v49 >> ((v32) & 63)))
                            _ = v91
                        }
                        _ = fi68
                        _ = v69
                    }
                }
            }
            v4[""] = -9223372036854775808
            _ = v49
        }
        _ = fi39
    }
    v92 := GenS13{F0: 586, F1: 127, F2: false}
    switch ("\n\t" + v3) {
    case "hello":
        v33 = 1
    case "a":
        for fi93 := 0; fi93 < 8; fi93++ {
            fmt.Println(v5)
            _ = fi93
        }
    case "\n\t":
        _ = len(v4)
    default:
        if (v32 == (v32 - v32)) {
            v94 := v4[v3]
            for fi95 := 0; fi95 < 5; fi95++ {
                var v96 float32 = 1.5
                switch (v94 & v1) {
                case -9223372036854775808:
                    v97 := []int{(fi95 >> ((v32) & 63)), -2147483648}
                    _ = v97
                case -1:
                    v98 := GenS14{F0: 958}
                    _ = v98
                case 676:
                    for fi99 := 0; fi99 < 20; fi99++ {
                        v100 := []int{(v94 * fi95)}
                        switch v3 {
                        case "":
                            if (!(1 <= 511)) {
                                if (true && ((-2147483648 <= 42) && (-2.25 == 1.5))) {
                                    v35 = (v35 * v35)
                                }
                                switch v3 {
                                case "\n\t":
                                    switch ((v3 + ((("foo bar" + ((v3 + v3) + v3)) + v3) + ((v3 + "a") + v3))) + (v3 + v3)) {
                                    case "hello":
                                        v35 -= 9223372036854775807
                                    case "":
                                        fmt.Println(v32, fi95, fi95)
                                    }
                                    switch v1 {
                                    case 0:
                                        v32 -= (v32 >> ((65535) & 63))
                                    case 32768:
                                        if len(v100) >= 2 {
                                            v100 = v100[1:len(v100)]
                                        }
                                    }
                                    fmt.Println(fi99, fi95)
                                    v96 -= (v96 * 1e10)
                                case "hello":
                                    switch ((v3 + "") + "hello") {
                                    case "":
                                        var v101 int = genFn0(v35)
                                        _ = v101
                                    case "foo bar":
                                        fmt.Println(fi99, v1)
                                    case "\n\t":
                                        delete(v2, 2147483648)
                                    }
                                default:
                                    switch v3 {
                                    case "foo bar":
                                        var v102 uint32 = 42
                                        var v103 int = 32767
                                        _ = v102
                                        _ = v103
                                    case "a":
                                        var v104 float32 = v96
                                        _ = v104
                                    case "":
                                        switch (("foo bar" + (v3 + v3)) + "foo bar") {
                                        case "a":
                                            fmt.Println(fi99, v3)
                                        case "foo bar":
                                            if (false && (191 != 0)) {
                                                if true {
                                                    var v105 int16 = 0
                                                    var v106 int16 = v105
                                                    _ = v105
                                                    _ = v106
                                                }
                                                switch "" {
                                                case "hello":
                                                    v107 := v4[v3]
                                                    _ = v107
                                                case "\n\t":
                                                    v1 += (v1 << (((v32 + v32)) & 63))
                                                case "a":
                                                    fmt.Println(v3)
                                                }
                                                v108 := v2[42]
                                                fmt.Println(v33)
                                                _ = v108
                                            }
                                        case "":
                                            var v109 uint32 = 2147483648
                                            _ = v109
                                        default:
                                            fmt.Println(v1)
                                        }
                                    default:
                                        var v110 uint16 = 599
                                        _ = v110
                                    }
                                }
                            } else {
                                var v111 int8 = (v5 ^ v5)
                                _ = v111
                            }
                        default:
                            v96 -= float32(1.5)
                        }
                        if false {
                            v1 -= genFn0((v35 << ((v32) & 63)))
                        } else {
                            v32 -= v32
                            if false {
                                var v112 uint16 = 1
                                var v113 uint32 = 177
                                v114 := GenS13{F0: 0, F1: 0, F2: (v32 != (v32 - (v32 * v32)))}
                                _ = v112
                                _ = v113
                                _ = v114
                            }
                        }
                        fmt.Println(v1, fi95)
                        _ = fi99
                        _ = v100
                    }
                default:
                    for fi115 := 0; fi115 < 4; fi115++ {
                        var v116 int8 = (v5 * 42)
                        if false {
                            v117 := GenS12{F0: -2147483648, F1: (true && ((!true) || false)), F2: v3}
                            switch ("a" + v3) {
                            case "a":
                                if ((42 < 0) || (-1 != 1000)) {
                                    switch 32767 {
                                    case 42:
                                        fmt.Println(v3, v1)
                                    case 32767:
                                        var v118 uint32 = 2
                                        _ = v118
                                    }
                                }
                            }
                            if false {
                                var v119 int16 = 1
                                _ = v119
                            }
                            _ = v117
                        }
                        v32 -= (v32 + 4294967295)
                        _ = fi115
                        _ = v116
                    }
                }
                for fi120 := 0; fi120 < 2; fi120++ {
                    _ = len(v4)
                    v121 := GenS1{F0: -374}
                    var v122 bool = false
                    var v123 int64 = -846
                    fmt.Println(v3, v1)
                    _ = fi120
                    _ = v121
                    _ = v122
                    _ = v123
                }
                _ = fi95
                _ = v96
            }
            _ = v94
        }
    }
    for fi124 := 0; fi124 < 3; fi124++ {
        fmt.Println(v35)
        fmt.Println(v3, v5)
        for fi125 := 0; fi125 < 18; fi125++ {
            v1 = 732
            _ = fi125
        }
        _ = fi124
    }
    for fi126 := 0; fi126 < 5; fi126++ {
        v127 := GenS15{F0: 42}
        switch v3 {
        case "":
            fmt.Println(v33)
        case "hello":
            var v128 uint8 = 1
            _ = v128
        case "a":
            for fi129 := 0; fi129 < 17; fi129++ {
                if false {
                    var v130 uint32 = 511
                    v3 = v3
                    v131 := []int{v1}
                    var v132 int64 = -475
                    _ = v130
                    _ = v131
                    _ = v132
                } else {
                    v5 = (-(v5))
                }
                _ = fi129
            }
        default:
            v5 = -1
        }
        v133 := GenS11{F0: 42, F1: 767}
        _ = fi126
        _ = v127
        _ = v133
    }
    v134 := v4[v3]
    for fi135 := 0; fi135 < 12; fi135++ {
        v33 = (v33 % ((0) | 1))
        switch v3 {
        case "":
            if false {
                switch "hello" {
                case "foo bar":
                    var v136 int16 = 32767
                    _ = v136
                default:
                    if (false && true) {
                        v137 := []int{(v1 % ((0) | 1)), (fi135 * (^(v1)))}
                        if (true || (!false)) {
                            v5 -= (v5 + (v5 ^ v5))
                        } else {
                            var v138 bool = (!(337 == -1))
                            _ = len(v4)
                            _ = v138
                        }
                        fmt.Println(v32, v35, v134)
                        _ = v137
                    } else {
                        switch (v134 % ((v134) | 1)) {
                        case 2147483647:
                            v139 := GenS16{F0: -320, F1: 222}
                            _ = v139
                        default:
                            if (true || (!(42 < 42))) {
                                delete(v2, 2147483648)
                                fmt.Println(fi135, v1)
                            }
                        }
                    }
                }
                var v140 uint16 = 362
                v141 := GenS17{F0: 42, F1: (v140 / ((v140) | 1)), F2: 18446744073709551615}
                _ = v140
                _ = v141
            } else {
                var v142 uint32 = 42
                v143 := []int{}
                _ = v142
                _ = v143
            }
        }
        var v144 uint64 = (v33 / ((v33) | 1))
        _ = fi135
        _ = v144
    }
    for fi145 := 0; fi145 < 20; fi145++ {
        switch v3 {
        case "":
            if (255 < 559) {
                var v146 int64 = -120
                for fi147 := 0; fi147 < 12; fi147++ {
                    var v148 int16 = -422
                    v33 -= v33
                    _ = fi147
                    _ = v148
                }
                var v149 int32 = 0
                _ = v146
                _ = v149
            }
        }
        _ = fi145
    }
    v150 := []int{-2147483648, 9223372036854775807, (v134 * v1), v134}
    v151 := GenS18{F0: (v5 / ((-1) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v134:", v134)
    fmt.Println("slicelen v150:", len(v150), v150)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v134
    _ = v150
    _ = v151
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v4
    _ = v5
    _ = v92
}
