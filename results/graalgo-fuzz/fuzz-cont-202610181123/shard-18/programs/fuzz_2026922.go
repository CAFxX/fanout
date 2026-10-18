package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS2 struct { F0 string; F1 int16; F2 uint8 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16; F1 string }
type GenS8 struct { F0 uint8; F1 int8 }
type GenS9 struct { F0 bool; F1 uint }
type GenS10 struct { F0 uint32; F1 bool }
type GenS11 struct { F0 uint64; F1 int16 }
type GenS12 struct { F0 int32; F1 string; F2 int8 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint; F1 int16 }
type GenS15 struct { F0 uint32; F1 string }
func genFn0() uint64 {
    var v5 float64 = (1.5 + (1e10 - (1e10 - 0.1)))
    v5 += (v5 / ((1e10 - 3.14159) - v5))
    v5 -= 3.14159
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := []int{fi6, fi6, (fi6 | fi6), (^(fi6))}
        _ = fi6
        _ = v7
    }
    _ = v5
    return 4294967295
}
func genFn1(p0 int) uint64 {
    var v8 uint8 = 1
    for fi9 := 0; fi9 < 14; fi9++ {
        switch "" {
        case "hello":
            v10 := GenS1{F0: 1, F1: 127, F2: 4294967295}
            _ = v10
        case "a":
            v11 := GenS1{F0: 900, F1: 1, F2: 0}
            _ = v11
        case "\n\t":
            v12 := []int{p0, (fi9 % ((fi9) | 1)), fi9, p0}
            _ = v12
        }
        _ = fi9
    }
    p0 = -692
    var v13 int8 = 1
    p0 = (p0 % (((p0 - (-(p0)))) | 1))
    _ = p0
    _ = v13
    _ = v8
    return 63
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]uint32)
    var v3 bool = false
    var v4 int8 = -1
    v3 = (36 != 42)
    v14 := GenS1{F0: uint32(v1), F1: v4, F2: 65535}
    v15 := GenS1{F0: 43, F1: v4, F2: 703}
    v4 = 42
    if (0.1 < 3.14159) {
        switch 712 {
        case 9223372036854775807:
            fmt.Println(v3, v3)
        case 0:
            if true {
                v3 = (42 <= 0)
                var v16 int = -46
                v17 := GenS1{F0: 599, F1: (v4 & -27), F2: 2147483647}
                _ = v16
                _ = v17
            } else {
                var v18 float32 = (0.1 - (1.5 - 3.14159))
                v4 = (v4 >> ((0) & 63))
                _ = v18
            }
        default:
            v4 += v4
        }
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 float64 = 0.1
            fmt.Println(v1, v3, fi19)
            var v21 string = ""
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        if (!(0.0 < float32((float64(3.14159) + ((3.14159 * float64(1.5)) / 3.14159))))) {
            for fi22 := 0; fi22 < 13; fi22++ {
                var v23 uint16 = 152
                _ = fi22
                _ = v23
            }
            switch "a" {
            case "foo bar":
                if ((-2.25 * 1.5) <= (0.0 + ((((1.5 - 0.0) / -2.25) - (-2.25 * 1e10)) * float64((-2.25 / (1e10 / (1e10 / -2.25))))))) {
                    for fi24 := 0; fi24 < 1; fi24++ {
                        switch ("foo bar" + "a") {
                        case "":
                            fmt.Println(v3)
                        }
                        fmt.Println(fi24, v4)
                        for fi25 := 0; fi25 < 7; fi25++ {
                            var v26 int32 = 42
                            _ = fi25
                            _ = v26
                        }
                        _ = fi24
                    }
                    var v27 string = ""
                    _ = v27
                } else {
                    if (false && false) {
                        var v28 int32 = -2147483648
                        if v3 {
                            switch "foo bar" {
                            case "hello":
                                v29 := GenS1{F0: 615, F1: (v4 | (-(v4))), F2: 2147483647}
                                _ = v29
                            default:
                                v30 := GenS1{F0: 2, F1: (v4 % (((v4 % ((v4) | 1))) | 1)), F2: 42}
                                _ = v30
                            }
                            if true {
                                v31 := GenS2{F0: "", F1: 1000, F2: 39}
                                fmt.Println(v4, v3)
                                _ = v31
                            } else {
                                if ((v28 >= (v28 << ((294) & 63))) && false) {
                                    for fi32 := 0; fi32 < 16; fi32++ {
                                        delete(v2, 4294967295)
                                        var v33 bool = (!false)
                                        _ = fi32
                                        _ = v33
                                    }
                                    fmt.Println(v28, v3, v1)
                                } else {
                                    if v3 {
                                        v28 -= v28
                                    }
                                    var v34 int32 = 879
                                    _ = v34
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    _ = len(v2)
                                }
                            }
                        } else {
                            switch (("" + "\n\t") + "a") {
                            case "":
                                var v35 uint = 940
                                _ = v35
                            case "\n\t":
                                v36 := GenS3{F0: 1, F1: 992}
                                _ = v36
                            case "a":
                                fmt.Println(v1, v1, v28)
                            }
                        }
                        for fi37 := 0; fi37 < 5; fi37++ {
                            var v38 int64 = -1
                            for fi39 := 0; fi39 < 19; fi39++ {
                                var v40 int32 = int32(v1)
                                var v41 bool = v3
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                            _ = fi37
                            _ = v38
                        }
                        _ = v28
                    } else {
                        var v42 int64 = -9223372036854775808
                        var v43 uint32 = 250
                        if v3 {
                            var v44 int8 = v4
                            var v45 string = ""
                            var v46 int16 = -1
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        _ = v42
                        _ = v43
                    }
                    var v47 int8 = (v4 | v4)
                    for fi48 := 0; fi48 < 17; fi48++ {
                        for fi49 := 0; fi49 < 16; fi49++ {
                            switch "foo bar" {
                            case "":
                                v50 := []int{(-(fi49))}
                                _ = v50
                            case "hello":
                                v51 := GenS3{F0: 623, F1: -1}
                                switch (fi48 >> ((18446744073709551615) & 63)) {
                                case 2147483647:
                                    var v52 string = "foo bar"
                                    _ = v52
                                case -508:
                                    v47 = v47
                                }
                                if (true && true) {
                                    v1 = genFn0()
                                    v53 := []int{fi48, (fi49 + 209), 918}
                                    _ = v53
                                }
                                _ = v51
                            case "\n\t":
                                var v54 uint8 = 127
                                _ = v54
                            }
                            fmt.Println(v1)
                            _ = fi49
                        }
                        _ = fi48
                    }
                    _ = v47
                }
            case "\n\t":
                switch "foo bar" {
                case "foo bar":
                    v4 += v4
                }
            default:
                for fi55 := 0; fi55 < 9; fi55++ {
                    v56 := GenS1{F0: 591, F1: (v4 << ((42) & 63)), F2: 2}
                    v3 = v3
                    for fi57 := 0; fi57 < 11; fi57++ {
                        switch (fi55 & fi57) {
                        case -516:
                            v58 := []int{(fi57 | fi55)}
                            var v59 int32 = -2147483648
                            v1 = (v1 ^ (v1 << ((856) & 63)))
                            _ = v58
                            _ = v59
                        }
                        v60 := []int{fi57, 1, fi57, 2147483647}
                        switch 681 {
                        case 32768:
                            if ((false && v3) && v3) {
                                var v61 uint32 = 65535
                                _ = v61
                            }
                        case -1:
                            v2[547] = 66
                        }
                        if false {
                            var v62 int = 32767
                            fmt.Println(v4, v62)
                            var v63 int = (fi55 % ((42) | 1))
                            fmt.Println(v3)
                            _ = v62
                            _ = v63
                        } else {
                            v64 := GenS4{F0: 1, F1: ("foo bar" + "a")}
                            var v65 uint8 = 127
                            if (v3 || v3) {
                                if (((v65 % (((v65 - (v65 | v65))) | 1)) < 121) && v3) {
                                    if len(v60) >= 2 {
                                        v60 = v60[1:len(v60)]
                                    }
                                    v65 = (v65 | (v65 % (((v65 / ((v65) | 1))) | 1)))
                                    if false {
                                        v66 := GenS4{F0: 42, F1: "a"}
                                        fmt.Println(v4)
                                        _ = v66
                                    }
                                }
                            }
                            _ = v64
                            _ = v65
                        }
                        fmt.Println(v4, fi57)
                        _ = fi57
                        _ = v60
                    }
                    _ = fi55
                    _ = v56
                }
            }
        } else {
            if (((v3 || v3) && true) || (v3 && v3)) {
                for fi67 := 0; fi67 < 7; fi67++ {
                    v4 = v4
                    v1 = v1
                    var v68 bool = true
                    _ = fi67
                    _ = v68
                }
            } else {
                switch (("hello" + ("a" + (("hello" + ("a" + (("foo bar" + "") + ("\n\t" + "a")))) + ""))) + "foo bar") {
                case "foo bar":
                    v69 := GenS4{F0: -9223372036854775808, F1: ""}
                    _ = v69
                case "hello":
                    fmt.Println(v4)
                default:
                    v3 = true
                }
                for fi70 := 0; fi70 < 6; fi70++ {
                    fmt.Println(v4, fi70)
                    v71 := []int{9223372036854775807}
                    v72 := GenS5{F0: v1, F1: 65535}
                    _ = fi70
                    _ = v71
                    _ = v72
                }
                _ = len(v2)
            }
            v2[765] = 42
        }
        switch 32768 {
        case 2147483647:
            var v73 int = 42
            _ = v73
        case -2147483648:
            var v74 uint8 = 139
            _ = v74
        }
        var v75 uint = 255
        _ = v75
    }
    fmt.Println(v4)
    if v3 {
        var v76 float32 = 3.14159
        fmt.Println(v76, v4, v4)
        v3 = (v1 >= v1)
        _ = v76
    } else {
        v1 = (v1 | v1)
        v77 := GenS6{F0: 65535}
        v1 = v1
        _ = v77
    }
    for fi78 := 0; fi78 < 2; fi78++ {
        for fi79 := 0; fi79 < 7; fi79++ {
            fmt.Println(v3, v3, fi79)
            fmt.Println(v1)
            fmt.Println(fi79, v4, fi78)
            _ = fi79
        }
        var v80 bool = false
        _ = fi78
        _ = v80
    }
    if (false || (!(v3 && ((2 >= 4294967295) || (!(!v3)))))) {
        switch "a" {
        case "a":
            fmt.Println(v3, v4, v3)
        case "":
            var v81 uint8 = 42
            _ = v81
        case "hello":
            fmt.Println(v3, v4)
        default:
            v82 := GenS5{F0: (v1 % ((v1) | 1)), F1: 42}
            _ = v82
        }
        if v3 {
            for fi83 := 0; fi83 < 3; fi83++ {
                var v84 bool = true
                _ = fi83
                _ = v84
            }
            var v85 int = 42
            var v86 bool = (!v3)
            _ = v85
            _ = v86
        }
    } else {
        var v87 int = 0
        _ = v87
    }
    for fi88 := 0; fi88 < 20; fi88++ {
        var v89 bool = (176 >= 42)
        for fi90 := 0; fi90 < 10; fi90++ {
            if (!v3) {
                fmt.Println(v1)
                v91 := []int{(fi88 & int(v4))}
                switch (fi90 ^ fi90) {
                case 4294967295:
                    v89 = (1 >= 620)
                }
                _ = v91
            } else {
                if v3 {
                    var v92 int16 = 1000
                    _ = v92
                }
                var v93 float32 = 1.5
                v3 = (false && (v4 > (v4 << ((4294967295) & 63))))
                _ = v93
            }
            if v3 {
                if (4294967295 != 65535) {
                    if true {
                        var v94 float32 = (1.5 + 0.1)
                        v95 := GenS7{F0: 0, F1: "\n\t"}
                        _ = v94
                        _ = v95
                    }
                    switch fi88 {
                    case 42:
                        if (0.0 == float64(0.1)) {
                            if v89 {
                                _ = len(v2)
                            } else {
                                v1 -= genFn1((fi90 ^ -961))
                                fmt.Println(v89)
                                fmt.Println(fi88, v89, v1)
                                fmt.Println(fi88, v3)
                                v3 = ((!(4294967295 <= 1)) || ("foo bar" != "hello"))
                                v89 = v89
                                fmt.Println(v4, v89)
                            }
                            var v96 uint = 18446744073709551615
                            _ = v96
                        }
                    default:
                        v97 := v2[fi88]
                        _ = v97
                    }
                    v4 -= (v4 - (v4 - 115))
                    v1 = genFn0()
                    var v98 float64 = -2.25
                    _ = v98
                } else {
                    v2[32767] = 65535
                    v99 := GenS8{F0: 212, F1: 100}
                    _ = v99
                }
                v1 += (v1 >> ((707) & 63))
                _ = len(v2)
            }
            _ = fi90
        }
        if false {
            v1 = 0
            for fi100 := 0; fi100 < 6; fi100++ {
                v89 = v89
                v101 := GenS9{F0: v89, F1: 65535}
                var v102 int = fi100
                _ = fi100
                _ = v101
                _ = v102
            }
        } else {
            switch 32767 {
            case 960:
                for fi103 := 0; fi103 < 6; fi103++ {
                    v104 := GenS1{F0: 2147483648, F1: v4, F2: 1}
                    v105 := GenS10{F0: 491, F1: true}
                    _ = len(v2)
                    if ((127 >= 0) && false) {
                        var v106 int64 = -18
                        fmt.Println(v3, fi103)
                        delete(v2, 128)
                        v89 = v3
                        _ = v106
                    } else {
                        var v107 uint16 = 1
                        v3 = (!v89)
                        switch (("\n\t" + "") + "a") {
                        case "":
                            switch fi103 {
                            case 9223372036854775807:
                                fmt.Println(v1)
                            case 127:
                                var v108 uint8 = 42
                                _ = v108
                            case 4294967295:
                                if v3 {
                                    var v109 uint8 = 1
                                    fmt.Println(fi103)
                                    _ = v109
                                } else {
                                    v110 := GenS11{F0: (^(v1)), F1: -1}
                                    fmt.Println(fi103, v4)
                                    _ = v110
                                }
                            default:
                                if v89 {
                                    var v111 int64 = 0
                                    switch (fi103 ^ fi88) {
                                    case 0:
                                        v112 := GenS12{F0: 524, F1: ("a" + "\n\t"), F2: 0}
                                        _ = v112
                                    case -617:
                                        var v113 uint16 = v107
                                        var v114 float64 = ((3.14159 * (3.14159 * 1e10)) * (1.5 / 1.0))
                                        v115 := []int{230}
                                        _ = v113
                                        _ = v114
                                        _ = v115
                                    }
                                    _ = v111
                                } else {
                                    var v116 uint64 = v1
                                    if false {
                                        if (v3 && v3) {
                                            var v117 uint64 = (v116 ^ 2)
                                            _ = v117
                                        }
                                        v118 := []int{int(v107), (fi88 - fi103)}
                                        fmt.Println(v89)
                                        _ = v118
                                    } else {
                                        v119 := []int{fi103, (fi103 - fi88), (fi103 >> ((42) & 63)), (fi103 - (fi103 % (((^(fi103))) | 1)))}
                                        v120 := GenS13{F0: 42}
                                        v1 = 685
                                        _ = v119
                                        _ = v120
                                    }
                                    if v3 {
                                        switch "\n\t" {
                                        case "\n\t":
                                            v121 := GenS14{F0: 42, F1: 1000}
                                            _ = v121
                                        default:
                                            var v122 int64 = 9223372036854775807
                                            _ = v122
                                        }
                                        var v123 uint8 = 0
                                        var v124 uint16 = v107
                                        _ = v123
                                        _ = v124
                                    }
                                    v125 := GenS1{F0: 181, F1: (^(v4)), F2: 42}
                                    if true {
                                        switch int(v107) {
                                        case -340:
                                            var v126 int64 = 880
                                            _ = v126
                                        case 4294967295:
                                            fmt.Println(v89, v107)
                                        case 32767:
                                            fmt.Println(fi88)
                                        }
                                        v116 = v116
                                        var v127 string = "foo bar"
                                        _ = v127
                                    }
                                    v107 += (v107 / (((v107 << ((724) & 63))) | 1))
                                    fmt.Println(v4)
                                    _ = v116
                                    _ = v125
                                }
                            }
                        case "foo bar":
                            if false {
                                var v128 string = (("hello" + "hello") + "foo bar")
                                switch fi103 {
                                case -31:
                                    v107 = (v107 << ((255) & 63))
                                default:
                                    v129 := GenS1{F0: 345, F1: (v4 >> ((1) & 63)), F2: 4294967295}
                                    _ = v129
                                }
                                v2[fi88] = 0
                                _ = v128
                            }
                        case "hello":
                            var v130 int = 128
                            _ = v130
                        }
                        _ = len(v2)
                        _ = v107
                    }
                    _ = fi103
                    _ = v104
                    _ = v105
                }
            case 4294967295:
                v4 = 0
            default:
                v3 = ((v1 << ((1) & 63)) == v1)
            }
            v3 = (42 != 255)
            v3 = v3
        }
        _ = fi88
        _ = v89
    }
    switch "hello" {
    case "hello":
        v131 := []int{}
        _ = v131
    case "a":
        var v132 bool = (false && (128 != 255))
        _ = v132
    case "foo bar":
        v133 := GenS6{F0: 2147483648}
        _ = v133
    }
    var v134 int64 = 9223372036854775807
    v135 := []int{32768, -2147483648, 127, 9223372036854775807}
    for fi136 := 0; fi136 < 3; fi136++ {
        v3 = false
        if v3 {
            if ((!(!v3)) && (65535 < 425)) {
                var v137 uint64 = genFn0()
                if (1 > 0) {
                    if v3 {
                        v3 = v3
                    }
                    for fi138 := 0; fi138 < 4; fi138++ {
                        _ = len(v2)
                        _ = fi138
                    }
                }
                switch (^(fi136)) {
                case 4294967295:
                    var v139 int = 32767
                    _ = v139
                }
                _ = v137
            }
        } else {
            _ = len(v2)
            for fi140 := 0; fi140 < 7; fi140++ {
                if (((-919 > 42) || false) || v3) {
                    var v141 uint32 = 2147483648
                    _ = len(v135)
                    v134 -= (-(v134))
                    _ = v141
                } else {
                    _ = len(v2)
                    var v142 uint16 = 255
                    v143 := GenS10{F0: 348, F1: ((^(v142)) != v142)}
                    _ = v142
                    _ = v143
                }
                v134 = (v134 | (-(v134)))
                _ = fi140
            }
        }
        switch fi136 {
        case 4294967295:
            v144 := GenS10{F0: 4294967295, F1: v3}
            _ = v144
        case 1:
            _ = len(v135)
        case 98:
            var v145 uint64 = (v1 ^ (v1 % ((genFn1((fi136 << ((255) & 63)))) | 1)))
            _ = v145
        }
        _ = fi136
    }
    if true {
        switch 4294967295 {
        case -686:
            v146 := v2[-88]
            _ = v146
        }
    } else {
        switch "\n\t" {
        case "foo bar":
            switch 1 {
            case -9223372036854775808:
                fmt.Println(v1, v134)
            default:
                v2[42] = 744
            }
        case "":
            v3 = (v3 || false)
        case "hello":
            _ = len(v2)
        default:
            for fi147 := 0; fi147 < 15; fi147++ {
                var v148 uint8 = 255
                switch "a" {
                case "":
                    v149 := GenS15{F0: uint32(v148), F1: "a"}
                    _ = v149
                default:
                    switch int(v148) {
                    case 0:
                        fmt.Println(v4, v3)
                    case 280:
                        var v150 int = (-(fi147))
                        _ = v150
                    case 127:
                        var v151 uint64 = (v1 >> ((65535) & 63))
                        _ = v151
                    }
                }
                _ = fi147
                _ = v148
            }
        }
        v135 = append(v135, 128)
        fmt.Println(v3, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v134:", v134)
    fmt.Println("slicelen v135:", len(v135), v135)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v134
    _ = v135
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
