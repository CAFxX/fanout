package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint32 }
type GenS2 struct { F0 int; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 int8; F2 uint64 }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 int16; F1 uint16; F2 int32 }
type GenS10 struct { F0 int64; F1 uint }
type GenS11 struct { F0 int; F1 int16 }
type GenS12 struct { F0 string; F1 uint64; F2 bool }
type GenS13 struct { F0 int16; F1 int16; F2 uint }
type GenS14 struct { F0 int; F1 int16; F2 string }
func genFn0(p0 bool) uint32 {
    v6 := GenS1{F0: 0, F1: 219, F2: 2147483647}
    var v7 uint64 = 253
    _ = p0
    _ = v6
    _ = v7
    return 65535
}
func genFn1() uint64 {
    var v8 int16 = -108
    v8 -= v8
    v8 += (-(v8))
    v8 = 42
    var v9 uint16 = 606
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 int64 = 888
    var v2 int64 = (-(v1))
    v3 := make(map[uint32]string)
    var v4 float64 = 1.5
    var v5 int32 = 2147483647
    if (0.1 == ((1.5 / (1.5 * (3.14159 + -2.25))) / -2.25)) {
        v10 := GenS1{F0: -52, F1: 42, F2: 2147483648}
        if (-2.25 > 0.0) {
            if true {
                var v11 float32 = 0.1
                v12 := GenS1{F0: -128, F1: 18446744073709551615, F2: 919}
                _ = len(v3)
                _ = v11
                _ = v12
            }
            var v13 uint64 = 818
            var v14 int16 = 32767
            _ = v13
            _ = v14
        }
        v1 += v1
        _ = v10
    } else {
        fmt.Println(v5, v5, v1)
        v15 := []int{0, -9223372036854775808, 32768}
        for fi16 := 0; fi16 < 7; fi16++ {
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 bool = (!((v5 | v5) == -198))
                var v19 int8 = 21
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi16
        }
        _ = v15
    }
    if (!false) {
        for fi20 := 0; fi20 < 20; fi20++ {
            fmt.Println(v2, v2)
            v5 -= -661
            switch "\n\t" {
            case "\n\t":
                delete(v3, 2)
            case "a":
                v21 := v3[2147483647]
                _ = v21
            case "":
                for fi22 := 0; fi22 < 14; fi22++ {
                    if (true || (!false)) {
                        if (!(true && false)) {
                            v23 := []int{(-(fi22)), (fi22 + 648), 933, fi22}
                            v1 -= (v1 << ((0) & 63))
                            fmt.Println(fi20)
                            v4 -= v4
                            _ = v23
                        } else {
                            v24 := GenS2{F0: (fi20 ^ fi22), F1: 118, F2: 105}
                            v25 := GenS3{F0: 324, F1: genFn1()}
                            var v26 int8 = 13
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        v27 := v3[990]
                        v4 += v4
                        _ = v27
                    }
                    switch "foo bar" {
                    case "a":
                        if false {
                            var v28 uint8 = 1
                            fmt.Println(v28, v28)
                            v5 = v5
                            switch 42 {
                            case 32768:
                                if ((v5 << ((100) & 63)) != v5) {
                                    switch (fi22 >> ((42) & 63)) {
                                    case 4294967295:
                                        v29 := GenS4{F0: false}
                                        _ = v29
                                    case -9223372036854775808:
                                        v30 := []int{(-(fi20)), (fi22 - -2147483648), (fi20 + (fi20 ^ fi22))}
                                        _ = v30
                                    }
                                    v31 := GenS4{F0: ("\n\t" <= "a")}
                                    v28 = v28
                                    _ = v31
                                } else {
                                    v3[980] = ""
                                    v1 = v1
                                    v32 := GenS5{F0: ("hello" + ""), F1: -127, F2: 42}
                                    _ = v32
                                }
                            case 2147483647:
                                v33 := []int{}
                                _ = v33
                            case 32767:
                                v34 := []int{42, fi22}
                                _ = v34
                            default:
                                if (128 < (v28 & (v28 & (^(v28))))) {
                                    if true {
                                        var v35 uint = 469
                                        v36 := []int{(fi20 - 32768), fi22}
                                        _ = len(v3)
                                        _ = v35
                                        _ = v36
                                    }
                                    v1 = v2
                                    v4 = (0.0 + 0.1)
                                } else {
                                    fmt.Println(fi22, fi20, fi22)
                                    switch "foo bar" {
                                    case "foo bar":
                                        if false {
                                            v5 = (v5 << ((236) & 63))
                                            fmt.Println(fi20)
                                        } else {
                                            v3[65535] = ("foo bar" + "foo bar")
                                            v37 := []int{(fi22 + fi20), (fi20 | (fi20 ^ 4294967295)), (fi20 / ((fi20) | 1)), -942}
                                            _ = v37
                                        }
                                    case "hello":
                                        var v38 int = fi22
                                        _ = v38
                                    default:
                                        var v39 bool = false
                                        _ = v39
                                    }
                                    delete(v3, 4294967295)
                                }
                            }
                            _ = v28
                        }
                    case "\n\t":
                        fmt.Println(v1, v1)
                    case "foo bar":
                        if (true || (162 < 255)) {
                            var v40 uint8 = 30
                            v40 = (v40 + (v40 << ((1) & 63)))
                            _ = v40
                        } else {
                            v2 = (^(v2))
                            var v41 string = "foo bar"
                            _ = v41
                        }
                    default:
                        v5 += (v5 - 1)
                    }
                    var v42 bool = false
                    _ = fi22
                    _ = v42
                }
            }
            _ = fi20
        }
        v43 := GenS6{F0: -58, F1: (v1 * (v1 & (v2 / ((v2) | 1))))}
        _ = v43
    }
    switch 32767 {
    case 1:
        switch "hello" {
        case "foo bar":
            for fi44 := 0; fi44 < 6; fi44++ {
                if true {
                    v45 := []int{(fi44 + 0), fi44, fi44, 128}
                    _ = v45
                } else {
                    for fi46 := 0; fi46 < 1; fi46++ {
                        v47 := GenS2{F0: fi44, F1: 2, F2: 94}
                        var v48 uint64 = 968
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                    v5 -= (v5 << ((42) & 63))
                }
                _ = fi44
            }
        }
    default:
        switch "" {
        case "foo bar":
            fmt.Println(v5, v4, v2)
        case "hello":
            for fi49 := 0; fi49 < 1; fi49++ {
                for fi50 := 0; fi50 < 3; fi50++ {
                    var v51 uint64 = 18446744073709551615
                    switch "a" {
                    case "":
                        var v52 uint8 = 128
                        _ = v52
                    case "hello":
                        switch "" {
                        case "hello":
                            v1 -= 388
                        default:
                            v2 = int64(v5)
                        }
                    }
                    _ = fi50
                    _ = v51
                }
                if (42 >= -444) {
                    switch "\n\t" {
                    case "":
                        v5 = -771
                    case "a":
                        fmt.Println(fi49)
                    }
                    v53 := GenS6{F0: -128, F1: (v1 / ((1) | 1))}
                    _ = v53
                } else {
                    for fi54 := 0; fi54 < 15; fi54++ {
                        var v55 string = ""
                        if false {
                            var v56 string = (v55 + (v55 + "foo bar"))
                            v5 -= (v5 & (v5 - (v5 & v5)))
                            var v57 int16 = 356
                            _ = v56
                            _ = v57
                        } else {
                            v58 := []int{0, -9223372036854775808, fi49}
                            v58 = append(v58, -2147483648)
                            _ = v58
                        }
                        if true {
                            switch fi54 {
                            case -641:
                                delete(v3, 1)
                                v59 := GenS3{F0: 1, F1: 42}
                                _ = v59
                            case -9223372036854775808:
                                if false {
                                    v60 := GenS7{F0: 1, F1: 1}
                                    var v61 bool = false
                                    if v61 {
                                        _ = len(v3)
                                    } else {
                                        v62 := []int{fi54}
                                        v1 += 1
                                        _ = v62
                                    }
                                    _ = v60
                                    _ = v61
                                }
                            case -2147483648:
                                v63 := GenS3{F0: genFn1(), F1: 548}
                                _ = v63
                            default:
                                if true {
                                    fmt.Println(v1, v1)
                                    v64 := GenS8{F0: 42, F1: int8(v2)}
                                    _ = len(v3)
                                    _ = v64
                                }
                            }
                            var v65 int16 = int16(v5)
                            var v66 int8 = 127
                            var v67 int16 = (v65 | v65)
                            v68 := GenS9{F0: v65, F1: 42, F2: v5}
                            _ = len(v3)
                            _ = v65
                            _ = v66
                            _ = v67
                            _ = v68
                        } else {
                            if (!(42 > 227)) {
                                if (1 != -32768) {
                                    var v69 int16 = 32767
                                    var v70 string = (v55 + (("a" + v55) + "\n\t"))
                                    var v71 int32 = -109
                                    _ = v69
                                    _ = v70
                                    _ = v71
                                }
                            }
                            v72 := []int{}
                            fmt.Println(fi54, v2, fi54)
                            var v73 uint16 = 153
                            v55 = v55
                            v74 := GenS10{F0: 42, F1: 18446744073709551615}
                            _ = v72
                            _ = v73
                            _ = v74
                        }
                        fmt.Println(v1, fi54)
                        _ = fi54
                        _ = v55
                    }
                    v5 = v5
                }
                switch -68 {
                case 1:
                    if (255 == 323) {
                        if false {
                            var v75 int = (^(fi49))
                            if ((!("hello" >= ((("a" + "") + "hello") + "foo bar"))) && ((127 > 2) || true)) {
                                fmt.Println(v75, v5)
                            }
                            var v76 uint64 = 18446744073709551615
                            _ = v75
                            _ = v76
                        } else {
                            fmt.Println(v4, v5, v4)
                            if false {
                                for fi77 := 0; fi77 < 4; fi77++ {
                                    v1 = (v1 - v2)
                                    switch (("" + "foo bar") + ("a" + "")) {
                                    case "":
                                        v78 := []int{572, (fi49 / ((fi77) | 1)), (-(fi49))}
                                        _ = v78
                                    }
                                    _ = fi77
                                }
                                v2 = v1
                                v79 := GenS8{F0: -32768, F1: -128}
                                _ = v79
                            }
                            var v80 int16 = -935
                            _ = v80
                        }
                    }
                case 32768:
                    fmt.Println(v4, v5)
                case -2147483648:
                    v2 += v2
                default:
                    var v81 int16 = -489
                    _ = v81
                }
                _ = fi49
            }
        case "\n\t":
            v4 = v4
        default:
            if ((3.14159 - (1e10 + ((3.14159 + ((1.5 / 3.14159) - float32(((float64(0.0) * -2.25) - (1e10 / v4))))) / 0.1))) > 0.0) {
                var v82 float64 = (v4 * v4)
                _ = v82
            }
        }
    }
    switch 2147483647 {
    case 32767:
        var v83 uint = 255
        _ = v83
    }
    var v84 int32 = -1
    v85 := []int{42, 758, 127}
    v5 = (v84 << ((927) & 63))
    if true {
        fmt.Println(v4)
        v86 := v3[2147483648]
        var v87 int16 = 427
        _ = v86
        _ = v87
    } else {
        if true {
            fmt.Println(v1, v5, v5)
            var v88 uint = 4294967295
            for fi89 := 0; fi89 < 13; fi89++ {
                if (fi89 != fi89) {
                    switch "foo bar" {
                    case "":
                        v5 += v5
                    case "a":
                        if (false || (true || (false && (!true)))) {
                            for fi90 := 0; fi90 < 4; fi90++ {
                                var v91 int = (fi89 ^ (fi89 + fi90))
                                switch (^(fi89)) {
                                case 32768:
                                    v1 = -1
                                case -1:
                                    v1 = v1
                                case 9223372036854775807:
                                    v88 = 2
                                default:
                                    v92 := GenS7{F0: 4294967295, F1: genFn0((3.14159 >= 0.1))}
                                    _ = v92
                                }
                                _ = fi90
                                _ = v91
                            }
                            var v93 int32 = (v84 - (v5 & 0))
                            _ = v93
                        }
                    default:
                        var v94 float64 = v4
                        _ = v94
                    }
                    v84 = (v84 | (-(v84)))
                }
                if (true || (false || false)) {
                    var v95 uint64 = 0
                    _ = v95
                }
                _ = fi89
            }
            _ = v88
        }
    }
    switch "" {
    case "\n\t":
        v96 := GenS7{F0: 42, F1: 42}
        _ = v96
    case "":
        if len(v85) >= 2 {
            v85 = v85[1:len(v85)]
        }
    case "hello":
        if true {
            if (65535 <= 476) {
                for fi97 := 0; fi97 < 6; fi97++ {
                    for fi98 := 0; fi98 < 15; fi98++ {
                        fmt.Println(v2, v4)
                        switch "foo bar" {
                        case "":
                            v5 = 78
                        case "\n\t":
                            switch fi98 {
                            case 2147483647:
                                v99 := GenS10{F0: 9223372036854775807, F1: 65535}
                                _ = v99
                            case 9223372036854775807:
                                if false {
                                    v84 -= -1
                                    if (!false) {
                                        var v100 int = (fi98 << ((54) & 63))
                                        switch fi97 {
                                        case 127:
                                            var v101 uint32 = 209
                                            _ = v101
                                        case 32767:
                                            v1 -= v2
                                        case 702:
                                            v85 = append(v85, 880)
                                        }
                                        _ = v100
                                    }
                                    if (((((1.5 + (1.5 - -2.25)) * float32(-2.25)) * 0.0) == 1.5) && (!true)) {
                                        v4 += 0.0
                                        switch "foo bar" {
                                        case "\n\t":
                                            switch "foo bar" {
                                            case "":
                                                v84 = -2147483648
                                                switch -1 {
                                                case 9223372036854775807:
                                                    v102 := GenS8{F0: 876, F1: -128}
                                                    v3[2] = ""
                                                    var v103 uint32 = 2
                                                    _ = v102
                                                    _ = v103
                                                case 464:
                                                    switch (("foo bar" + "\n\t") + "hello") {
                                                    case "\n\t":
                                                        if (2 < 4294967295) {
                                                            var v104 int32 = (v5 * v84)
                                                            var v105 uint64 = 2
                                                            var v106 bool = false
                                                            _ = v104
                                                            _ = v105
                                                            _ = v106
                                                        } else {
                                                            var v107 float64 = ((1.5 + v4) + 1e10)
                                                            v85 = append(v85, int(v84))
                                                            if len(v85) >= 2 {
                                                                v85 = v85[1:len(v85)]
                                                            }
                                                            _ = v107
                                                        }
                                                    case "":
                                                        v108 := GenS11{F0: (fi98 >> ((786) & 63)), F1: 32767}
                                                        _ = v108
                                                    case "a":
                                                        if len(v85) > 0 {
                                                            v85[0] = (fi97 ^ (fi97 / ((127) | 1)))
                                                        }
                                                    default:
                                                        v4 = ((((v4 * v4) / (v4 - v4)) - float64(3.14159)) + v4)
                                                    }
                                                case 208:
                                                    fmt.Println(v84)
                                                }
                                            case "a":
                                                switch (fi97 / ((fi97) | 1)) {
                                                case 373:
                                                    var v109 uint8 = 255
                                                    _ = v109
                                                default:
                                                    v110 := GenS8{F0: 0, F1: 28}
                                                    _ = len(v85)
                                                    var v111 int32 = -1
                                                    _ = v110
                                                    _ = v111
                                                }
                                            case "\n\t":
                                                v5 -= v5
                                            default:
                                                var v112 bool = false
                                                _ = v112
                                            }
                                        }
                                    }
                                } else {
                                    v1 += -1
                                    var v113 uint = 0
                                    _ = v113
                                }
                            }
                        default:
                            v114 := GenS12{F0: "hello", F1: 2, F2: (0 == (v84 ^ v5))}
                            v2 += v1
                            fmt.Println(v1)
                            _ = v114
                        }
                        if (((v4 * 0.0) / float64(-2.25)) < (((v4 - 3.14159) * 1.5) / 0.1)) {
                            var v115 int16 = -797
                            switch "" {
                            case "foo bar":
                                var v116 int64 = v2
                                _ = v116
                            case "hello":
                                v1 = (v1 * v1)
                            case "\n\t":
                                var v117 int64 = v2
                                _ = v117
                            default:
                                var v118 float64 = (float64(1.5) - 3.14159)
                                _ = v118
                            }
                            _ = v115
                        } else {
                            if true {
                                fmt.Println(fi97)
                                switch "hello" {
                                case "foo bar":
                                    var v119 uint8 = 2
                                    _ = v119
                                case "hello":
                                    v84 = (-(v84))
                                case "a":
                                    if true {
                                        var v120 float64 = 3.14159
                                        var v121 float64 = (v4 + ((float64((3.14159 * (1.5 + (float32(1e10) * (((((float32(v120) - float32(v4)) - 0.1) / 0.1) * float32((1.5 - 3.14159))) - ((((1e10 * -2.25) - 0.0) * 3.14159) * -2.25)))))) / v4) + 1.5))
                                        v120 = v120
                                        _ = v120
                                        _ = v121
                                    } else {
                                        _ = len(v85)
                                        v122 := GenS13{F0: 32767, F1: 1, F2: 460}
                                        v3[42] = "a"
                                        _ = v122
                                    }
                                default:
                                    if (!false) {
                                        v123 := GenS14{F0: (fi97 >> ((18446744073709551615) & 63)), F1: 42, F2: (("" + ("" + (("a" + "a") + ("foo bar" + ("" + "a"))))) + "foo bar")}
                                        _ = v123
                                    }
                                }
                                if (!false) {
                                    v124 := GenS5{F0: "\n\t", F1: -128, F2: 2}
                                    _ = v124
                                }
                            }
                            _ = len(v3)
                            var v125 uint8 = 0
                            _ = v125
                        }
                        _ = fi98
                    }
                    v4 += (float64(1e10) + v4)
                    if len(v85) > 0 {
                        v85[0] = (fi97 ^ (fi97 - (^(fi97))))
                    }
                    _ = fi97
                }
                var v126 int = -2147483648
                _ = v126
            } else {
                var v127 int16 = 42
                var v128 int16 = 0
                _ = v127
                _ = v128
            }
            switch 32767 {
            case 9223372036854775807:
                switch 1 {
                case 102:
                    var v129 int32 = v5
                    _ = v129
                case 9223372036854775807:
                    var v130 uint8 = 68
                    _ = v130
                }
            }
        }
    }
    v131 := GenS8{F0: -1, F1: 90}
    switch 9223372036854775807 {
    case 0:
        var v132 uint = 670
        _ = v132
    case -1:
        _ = len(v85)
    default:
        v5 = (v84 % (((v84 | (v5 * v5))) | 1))
    }
    for fi133 := 0; fi133 < 12; fi133++ {
        v134 := GenS7{F0: 403, F1: 2147483647}
        var v135 string = "foo bar"
        var v136 int64 = v1
        _ = fi133
        _ = v134
        _ = v135
        _ = v136
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v84:", v84)
    fmt.Println("slicelen v85:", len(v85), v85)
    _ = v1
    _ = v131
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v84
    _ = v85
}
