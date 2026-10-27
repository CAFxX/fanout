package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint; F1 string; F2 int16 }
type GenS5 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 bool; F1 string }
type GenS9 struct { F0 int64; F1 uint32; F2 uint }
type GenS10 struct { F0 bool; F1 int8; F2 bool }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int64; F1 bool; F2 bool }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int32; F1 int32; F2 int }
type GenS15 struct { F0 int16; F1 uint64 }
type GenS16 struct { F0 int8; F1 int64; F2 int16 }
func genFn0() uint8 {
    var v6 int = 42
    v6 += (v6 + v6)
    _ = v6
    return 0
}
func genFn1(p0 int64, p1 string) uint {
    if false {
        v7 := []int{-198, -9223372036854775808, 32767, 32767}
        switch -938 {
        case -603:
            if false {
                for fi8 := 0; fi8 < 17; fi8++ {
                    p0 += -9223372036854775808
                    p1 = p1
                    p0 += (p0 / (((-(p0))) | 1))
                    _ = fi8
                }
            } else {
                if true {
                    for fi9 := 0; fi9 < 14; fi9++ {
                        var v10 float64 = 0.0
                        for fi11 := 0; fi11 < 16; fi11++ {
                            if false {
                                var v12 int64 = p0
                                if (0 >= 255) {
                                    switch 4294967295 {
                                    case 32768:
                                        var v13 int64 = v12
                                        _ = v13
                                    case -790:
                                        v12 -= (p0 % (((v12 / (((v12 ^ (p0 << ((0) & 63)))) | 1))) | 1))
                                    case 42:
                                        var v14 float32 = -2.25
                                        _ = v14
                                    default:
                                        var v15 float64 = (float64(float32((v10 / v10))) / 1.0)
                                        _ = v15
                                    }
                                }
                                _ = len(v7)
                                v10 += 1.5
                                _ = v12
                            }
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = fi11
                        }
                        var v16 uint = 255
                        _ = fi9
                        _ = v10
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 4; fi17++ {
                        p1 = p1
                        v7 = append(v7, fi17)
                        _ = fi17
                    }
                } else {
                    var v18 string = p1
                    _ = v18
                }
                var v19 uint = 0
                if (!(false && (1.5 >= 1e10))) {
                    p1 = (("hello" + "hello") + p1)
                }
                _ = v19
            }
        default:
            p0 = p0
        }
        p1 = p1
        _ = v7
    } else {
        p0 -= p0
        var v20 float64 = 0.1
        _ = v20
    }
    v21 := GenS1{F0: 46, F1: 32767}
    p1 = p1
    _ = p0
    _ = p1
    _ = v21
    return 784
}

func main() {
    var v1 uint64 = 352
    var v2 int8 = -97
    var v3 uint = 744
    var v4 int64 = 42
    var v5 bool = false
    var v22 float64 = ((-2.25 - float64(-2.25)) * ((-2.25 * 0.1) - (1e10 + 0.1)))
    v23 := GenS1{F0: uint8(v1), F1: 1000}
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 uint64 = 42
        v1 += (v1 / (((^(v25))) | 1))
        for fi26 := 0; fi26 < 12; fi26++ {
            if (v5 || ((("a" + "a") + "") < "\n\t")) {
                var v27 bool = (v5 && ((fi26 + 0) >= int(v4)))
                v28 := []int{}
                var v29 uint32 = 0
                _ = v27
                _ = v28
                _ = v29
            } else {
                var v30 uint = 1
                v4 = 0
                v5 = v5
                _ = v30
            }
            if v5 {
                if (v3 >= v3) {
                    v31 := GenS2{F0: fi26, F1: (v2 - (-(v2)))}
                    switch "hello" {
                    case "":
                        fmt.Println(v5)
                    case "foo bar":
                        fmt.Println(v5, v4, fi24)
                        v32 := GenS1{F0: 239, F1: -1}
                        _ = v32
                    }
                    fmt.Println(fi24, v2, v5)
                    _ = v31
                } else {
                    fmt.Println(fi24, fi24, v3)
                }
            } else {
                switch "" {
                case "a":
                    if ((375 <= 1000) && v5) {
                        v33 := []int{fi26}
                        v22 = v22
                        _ = v33
                    }
                case "foo bar":
                    v34 := []int{(fi24 + 4294967295)}
                    _ = v34
                case "hello":
                    v25 = (v1 - v1)
                }
            }
            _ = fi26
        }
        _ = fi24
        _ = v25
    }
    v35 := GenS2{F0: -60, F1: 1}
    for fi36 := 0; fi36 < 14; fi36++ {
        var v37 uint16 = 0
        v1 = v1
        _ = fi36
        _ = v37
    }
    v38 := []int{59, -108}
    for fi39 := 0; fi39 < 19; fi39++ {
        v40 := GenS3{F0: "hello", F1: 42}
        if len(v38) >= 2 {
            v38 = v38[1:len(v38)]
        }
        var v41 uint16 = 255
        _ = fi39
        _ = v40
        _ = v41
    }
    var v42 uint32 = 36
    fmt.Println(v1)
    fmt.Println(v5)
    switch "hello" {
    case "hello":
        var v43 uint8 = 128
        _ = v43
    case "":
        v5 = (332 <= 0)
    default:
        if (2 <= genFn0()) {
            if (v4 >= 302) {
                for fi44 := 0; fi44 < 1; fi44++ {
                    switch "foo bar" {
                    case "\n\t":
                        var v45 float64 = (v22 + v22)
                        _ = v45
                    case "":
                        v46 := GenS4{F0: 889, F1: ("\n\t" + ("\n\t" + "")), F2: 957}
                        _ = v46
                    default:
                        v38 = append(v38, (fi44 / ((fi44) | 1)))
                    }
                    if v5 {
                        if false {
                            v47 := GenS5{F0: 42, F1: 255, F2: v1}
                            _ = v47
                        } else {
                            fmt.Println(v2, v42, v3)
                            if (v5 || ((v42 / ((v42) | 1)) >= 2147483647)) {
                                var v48 uint32 = uint32(v1)
                                if false {
                                    var v49 uint16 = uint16(v1)
                                    for fi50 := 0; fi50 < 11; fi50++ {
                                        _ = len(v38)
                                        v1 = (v1 << ((v3) & 63))
                                        _ = len(v38)
                                        _ = fi50
                                    }
                                    var v51 int32 = 0
                                    _ = v49
                                    _ = v51
                                }
                                if (!(!false)) {
                                    if len(v38) > 0 {
                                        v38[0] = (fi44 << (((v3 - (v3 & v3))) & 63))
                                    }
                                }
                                _ = v48
                            } else {
                                v4 -= (-(v4))
                                switch "\n\t" {
                                case "a":
                                    var v52 uint16 = 933
                                    _ = v52
                                default:
                                    var v53 int8 = v2
                                    _ = v53
                                }
                            }
                            fmt.Println(v4, v5)
                        }
                        switch 127 {
                        case -708:
                            if len(v38) > 0 {
                                v38[0] = (fi44 + (fi44 & fi44))
                            }
                        case 42:
                            fmt.Println(v4)
                        default:
                            var v54 int8 = v2
                            _ = v54
                        }
                    } else {
                        v55 := GenS6{F0: 1, F1: fi44}
                        _ = v55
                    }
                    if (!(!((42 > -132) || false))) {
                        var v56 uint = v3
                        switch fi44 {
                        case 667:
                            if v5 {
                                _ = len(v38)
                                v22 -= (float64((-2.25 - 3.14159)) * (v22 / (((v22 - (float64(1.5) * v22)) - 0.0) / float64((((1e10 - float32(3.14159)) - 3.14159) * -2.25)))))
                                v4 = 566
                            }
                        case 32768:
                            if v5 {
                                v5 = (v4 >= 42)
                                v38 = append(v38, (^(fi44)))
                            } else {
                                v4 -= 0
                                if v5 {
                                    if len(v38) > 0 {
                                        v38[0] = (fi44 | fi44)
                                    }
                                    v57 := GenS7{F0: "\n\t", F1: "a"}
                                    if v5 {
                                        var v58 uint = genFn1(-215, "a")
                                        v38 = append(v38, (^(fi44)))
                                        for fi59 := 0; fi59 < 2; fi59++ {
                                            if true {
                                                var v60 bool = ((v1 - 4294967295) < v1)
                                                v61 := GenS2{F0: fi44, F1: v2}
                                                _ = v60
                                                _ = v61
                                            }
                                            _ = fi59
                                        }
                                        _ = v58
                                    } else {
                                        v56 += 529
                                        for fi62 := 0; fi62 < 8; fi62++ {
                                            v3 -= 4294967295
                                            v63 := GenS8{F0: (-2147483648 > fi44), F1: ("a" + "foo bar")}
                                            _ = fi62
                                            _ = v63
                                        }
                                    }
                                    _ = v57
                                } else {
                                    if len(v38) > 0 {
                                        v38[0] = fi44
                                    }
                                }
                                for fi64 := 0; fi64 < 5; fi64++ {
                                    switch fi44 {
                                    case -955:
                                        v65 := GenS9{F0: v4, F1: (v42 - v42), F2: 293}
                                        _ = v65
                                    }
                                    var v66 uint64 = 4294967295
                                    v22 = (v22 + 0.1)
                                    _ = fi64
                                    _ = v66
                                }
                            }
                        }
                        _ = v56
                    } else {
                        if ((-1 >= 0) && (v5 || ((^(v3)) >= (v3 + (v3 - uint(v1)))))) {
                            switch fi44 {
                            case -245:
                                if false {
                                    v67 := GenS10{F0: (!(!v5)), F1: (v2 / (((v2 + (^(v2)))) | 1)), F2: (v22 <= v22)}
                                    var v68 int32 = -1
                                    if len(v38) >= 2 {
                                        v38 = v38[1:len(v38)]
                                    }
                                    _ = v67
                                    _ = v68
                                } else {
                                    v5 = ((^(v3)) >= 0)
                                }
                            case 32768:
                                var v69 int64 = (v4 | (-(v4)))
                                _ = v69
                            case 42:
                                fmt.Println(v3, v22, v4)
                            }
                            for fi70 := 0; fi70 < 11; fi70++ {
                                _ = len(v38)
                                switch -9223372036854775808 {
                                case -9223372036854775808:
                                    v71 := GenS11{F0: fi44}
                                    _ = v71
                                case 109:
                                    if len(v38) > 0 {
                                        v38[0] = (fi44 >> (((v3 & (^(v3)))) & 63))
                                    }
                                case -2147483648:
                                    v5 = v5
                                }
                                var v72 uint8 = 128
                                _ = fi70
                                _ = v72
                            }
                        }
                    }
                    _ = fi44
                }
                switch ("hello" + "\n\t") {
                case "a":
                    v22 = (float64(0.1) / ((v22 / float64((-2.25 * 1e10))) / (3.14159 - v22)))
                case "foo bar":
                    v2 += v2
                default:
                    for fi73 := 0; fi73 < 9; fi73++ {
                        v74 := GenS4{F0: v3, F1: "foo bar", F2: 552}
                        if v5 {
                            var v75 uint32 = (v42 % ((2147483647) | 1))
                            _ = v75
                        } else {
                            switch (-(fi73)) {
                            case 80:
                                fmt.Println(v4, v3, fi73)
                            case 32767:
                                v38 = append(v38, fi73)
                            case 543:
                                for fi76 := 0; fi76 < 1; fi76++ {
                                    if (v5 && ((v42 + 1) == v42)) {
                                        v42 = (v42 + 2147483647)
                                        v38 = append(v38, (fi76 * (fi76 << ((v3) & 63))))
                                    } else {
                                        v5 = ((!v5) && (!(201 >= 42)))
                                        v3 -= (^(v3))
                                        v5 = (-704 != -1)
                                    }
                                    v77 := GenS12{F0: v4, F1: true, F2: ((3.14159 < (((0.1 - 1.5) + (0.1 + -2.25)) - 3.14159)) && (!v5))}
                                    v78 := GenS2{F0: fi73, F1: 34}
                                    _ = fi76
                                    _ = v77
                                    _ = v78
                                }
                            default:
                                if (-2.25 == 0.1) {
                                    var v79 float64 = v22
                                    v80 := GenS13{F0: 1}
                                    _ = v79
                                    _ = v80
                                } else {
                                    if (-9223372036854775808 > (fi73 - fi73)) {
                                        var v81 uint8 = 234
                                        var v82 float64 = v22
                                        for fi83 := 0; fi83 < 20; fi83++ {
                                            v3 = v3
                                            var v84 int16 = 0
                                            _ = len(v38)
                                            v5 = v5
                                            v2 -= int8(v84)
                                            switch 9223372036854775807 {
                                            case 128:
                                                v22 = v22
                                                fmt.Println(v3, fi73)
                                            case 2147483647:
                                                fmt.Println(fi83, v82, v4)
                                            default:
                                                var v85 uint64 = v1
                                                _ = v85
                                            }
                                            _ = fi83
                                            _ = v84
                                        }
                                        _ = v81
                                        _ = v82
                                    } else {
                                        var v86 int8 = v2
                                        _ = v86
                                    }
                                    var v87 uint8 = 195
                                    for fi88 := 0; fi88 < 10; fi88++ {
                                        fmt.Println(v3, v1, fi73)
                                        _ = len(v38)
                                        var v89 bool = false
                                        _ = fi88
                                        _ = v89
                                    }
                                    _ = v87
                                }
                            }
                            for fi90 := 0; fi90 < 17; fi90++ {
                                if (0 == 127) {
                                    var v91 int8 = 0
                                    switch "a" {
                                    case "hello":
                                        if v5 {
                                            switch "foo bar" {
                                            case "a":
                                                var v92 int32 = 42
                                                _ = v92
                                            case "":
                                                if (true || ((v42 ^ (v42 + v42)) >= v42)) {
                                                    var v93 float32 = -2.25
                                                    v93 = ((float32(v22) * 1e10) + (0.1 * 0.1))
                                                    v94 := GenS14{F0: -1, F1: -187, F2: 42}
                                                    _ = v93
                                                    _ = v94
                                                } else {
                                                    var v95 uint64 = (v1 ^ (v1 | v1))
                                                    _ = v95
                                                }
                                            case "\n\t":
                                                v4 -= (v4 >> ((v3) & 63))
                                            }
                                            var v96 uint = (v3 * 381)
                                            _ = v96
                                        }
                                        if len(v38) > 0 {
                                            v38[0] = -1
                                        }
                                    case "a":
                                        switch "foo bar" {
                                        case "\n\t":
                                            if (127 != (v91 & v91)) {
                                                switch (("" + "\n\t") + ("" + "\n\t")) {
                                                case "hello":
                                                    v42 = (v42 % ((v42) | 1))
                                                default:
                                                    if (false || (!v5)) {
                                                        if (!(2147483647 > 42)) {
                                                            v4 = v4
                                                        }
                                                    }
                                                }
                                                if len(v38) >= 2 {
                                                    v38 = v38[1:len(v38)]
                                                }
                                                var v97 int = (^(fi73))
                                                _ = v97
                                            }
                                        case "a":
                                            var v98 int32 = 42
                                            _ = v98
                                        case "hello":
                                            var v99 int8 = (v91 + v2)
                                            _ = v99
                                        }
                                    default:
                                        var v100 uint = v3
                                        _ = v100
                                    }
                                    v42 = v42
                                    _ = v91
                                } else {
                                    v3 += 18446744073709551615
                                    if v5 {
                                        switch "hello" {
                                        case "":
                                            var v101 uint32 = (v42 >> ((v3) & 63))
                                            _ = v101
                                        case "hello":
                                            v22 = float64(((1.5 - (1e10 * 0.1)) - 0.1))
                                        default:
                                            if v5 {
                                                v5 = v5
                                                var v102 int8 = (v2 << (((^(v3))) & 63))
                                                var v103 uint = v3
                                                _ = v102
                                                _ = v103
                                            }
                                        }
                                    }
                                    v104 := GenS15{F0: -241, F1: 4294967295}
                                    _ = v104
                                }
                                if (-128 < (v2 / (((v2 | (^(v2)))) | 1))) {
                                    v4 += 1
                                    if true {
                                        v4 = (v4 + (v4 ^ (^(v4))))
                                        var v105 int32 = 1
                                        v22 = 0.1
                                        _ = v105
                                    } else {
                                        v106 := GenS3{F0: ("hello" + "hello"), F1: uint16(v3)}
                                        if (v5 || (!v5)) {
                                            var v107 uint16 = 0
                                            v108 := GenS16{F0: (v2 * (v2 << ((v3) & 63))), F1: (^(v4)), F2: 1}
                                            var v109 uint = genFn1((v4 & (v4 ^ v4)), "\n\t")
                                            _ = v107
                                            _ = v108
                                            _ = v109
                                        }
                                        if (v1 > (v1 >> ((uint(v1)) & 63))) {
                                            var v110 int16 = -616
                                            _ = v110
                                        } else {
                                            if v5 {
                                                var v111 uint8 = 0
                                                if len(v38) >= 2 {
                                                    v38 = v38[1:len(v38)]
                                                }
                                                if v5 {
                                                    _ = len(v38)
                                                    v112 := GenS9{F0: v4, F1: (v42 * (^(v42))), F2: 0}
                                                    _ = v112
                                                } else {
                                                    v38 = append(v38, (^(fi73)))
                                                    var v113 int = fi90
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v114 uint64 = (^(v1))
                                                        _ = v114
                                                    case "a":
                                                        v2 += v2
                                                    }
                                                    v115 := GenS9{F0: v4, F1: v42, F2: uint(v111)}
                                                    v111 -= v111
                                                    v1 -= 0
                                                    _ = v113
                                                    _ = v115
                                                }
                                                if len(v38) >= 2 {
                                                    v38 = v38[1:len(v38)]
                                                }
                                                switch fi73 {
                                                case -826:
                                                    var v116 float32 = 3.14159
                                                    _ = v116
                                                case 2147483647:
                                                    var v117 uint16 = 42
                                                    _ = v117
                                                }
                                                var v118 int32 = -2147483648
                                                var v119 uint32 = (v42 & (^(v42)))
                                                fmt.Println(v1, v111)
                                                _ = v111
                                                _ = v118
                                                _ = v119
                                            }
                                            if false {
                                                v120 := GenS12{F0: v4, F1: v5, F2: (!true)}
                                                _ = v120
                                            }
                                            fmt.Println(v42, v3, fi90)
                                        }
                                        _ = v106
                                    }
                                } else {
                                    _ = len(v38)
                                    v42 = v42
                                }
                                v121 := GenS16{F0: (-(v2)), F1: v4, F2: -795}
                                _ = fi90
                                _ = v121
                            }
                            if v5 {
                                fmt.Println(v3, v1, v42)
                            }
                        }
                        if (v5 && (2 < v42)) {
                            v38 = append(v38, fi73)
                        }
                        _ = fi73
                        _ = v74
                    }
                }
            } else {
                if (-2147483648 != 604) {
                    if v5 {
                        v122 := GenS8{F0: (18446744073709551615 <= v1), F1: ("hello" + "foo bar")}
                        _ = v122
                    }
                    for fi123 := 0; fi123 < 6; fi123++ {
                        v1 = v1
                        v5 = true
                        v1 = 18446744073709551615
                        _ = fi123
                    }
                    _ = len(v38)
                } else {
                    var v124 string = ((("foo bar" + "foo bar") + ("\n\t" + "\n\t")) + ("a" + "foo bar"))
                    if (!true) {
                        switch "foo bar" {
                        case "a":
                            if false {
                                for fi125 := 0; fi125 < 11; fi125++ {
                                    if (!(v5 && (v42 <= (v42 % (((^(v42))) | 1))))) {
                                        var v126 int64 = 69
                                        _ = v126
                                    }
                                    v124 = v124
                                    var v127 uint16 = uint16(v42)
                                    _ = fi125
                                    _ = v127
                                }
                            }
                        default:
                            v3 = v3
                        }
                        if len(v38) > 0 {
                            v38[0] = 32767
                        }
                        var v128 float64 = v22
                        _ = v128
                    }
                    _ = v124
                }
            }
            for fi129 := 0; fi129 < 13; fi129++ {
                v42 += 224
                v3 = v3
                if len(v38) >= 2 {
                    v38 = v38[1:len(v38)]
                }
                _ = fi129
            }
        } else {
            if (false || v5) {
                v3 = genFn1((v4 | -1), "")
                v22 = v22
            }
        }
    }
    v130 := GenS13{F0: 409}
    _ = len(v38)
    v131 := GenS7{F0: "", F1: ("" + "foo bar")}
    for fi132 := 0; fi132 < 16; fi132++ {
        v133 := GenS1{F0: 2, F1: -362}
        if true {
            for fi134 := 0; fi134 < 1; fi134++ {
                if false {
                    var v135 int32 = -2147483648
                    _ = v135
                } else {
                    var v136 float64 = float64(3.14159)
                    v1 = uint64(v3)
                    var v137 uint8 = uint8(v1)
                    _ = v136
                    _ = v137
                }
                v2 = (v2 % (((v2 % ((127) | 1))) | 1))
                if len(v38) >= 2 {
                    v38 = v38[1:len(v38)]
                }
                _ = fi134
            }
        }
        if v5 {
            for fi138 := 0; fi138 < 20; fi138++ {
                if (0 >= 98) {
                    _ = len(v38)
                    _ = len(v38)
                    fmt.Println(fi138, v42, v4)
                }
                v3 = (^(v3))
                v139 := GenS16{F0: v2, F1: v4, F2: 960}
                _ = fi138
                _ = v139
            }
        }
        _ = fi132
        _ = v133
    }
    var v140 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v140:", v140)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v130
    _ = v131
    _ = v140
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v42
    _ = v5
}
