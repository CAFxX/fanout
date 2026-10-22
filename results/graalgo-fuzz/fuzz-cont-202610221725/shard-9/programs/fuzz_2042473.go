package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 uint; F2 int16 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 bool; F1 int16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint32; F1 int32 }
type GenS12 struct { F0 uint16; F1 int32 }
type GenS13 struct { F0 int8; F1 bool; F2 int8 }
type GenS14 struct { F0 uint8; F1 uint64; F2 int }
type GenS15 struct { F0 string; F1 int64 }
type GenS16 struct { F0 string; F1 int16; F2 uint32 }
type GenS17 struct { F0 uint8; F1 string; F2 int32 }
type GenS18 struct { F0 uint16; F1 int8; F2 uint }
type GenS19 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS20 struct { F0 uint8; F1 int32; F2 int32 }
type GenS21 struct { F0 int; F1 int64; F2 string }
func genFn0() uint16 {
    v5 := []int{9223372036854775807, 2147483647}
    v6 := GenS1{F0: 255}
    var v7 int64 = 9223372036854775807
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int) uint8 {
    p0 = (p0 ^ (p0 ^ p0))
    v8 := GenS1{F0: 0}
    fmt.Println(p0)
    var v9 uint64 = 4294967295
    _ = p0
    _ = v8
    _ = v9
    return 94
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{42}
    var v3 int64 = -773
    var v4 int16 = -131
    switch "" {
    case "\n\t":
        for fi10 := 0; fi10 < 2; fi10++ {
            v11 := GenS2{F0: 0, F1: 0, F2: false}
            _ = fi10
            _ = v11
        }
    case "":
        if true {
            v3 = (v3 + (v3 >> ((42) & 63)))
            if false {
                if true {
                    v12 := GenS1{F0: 683}
                    if false {
                        for fi13 := 0; fi13 < 12; fi13++ {
                            var v14 uint8 = 242
                            v14 = (v14 | 2)
                            if (625 == 386) {
                                var v15 float64 = (((3.14159 - ((-2.25 * (0.0 * (3.14159 / 1.0))) + (1.5 + 1e10))) * (((((((0.0 - -2.25) + 1e10) - 1.5) * 0.1) * -2.25) - 1.5) * float64(1e10))) / float64((3.14159 * ((1e10 * (((1e10 + 0.1) - ((0.1 - 0.0) / (0.0 - (3.14159 + 0.0)))) * ((0.0 - 0.1) / ((3.14159 + 0.0) * -2.25)))) + 0.1))))
                                fmt.Println(fi13)
                                _ = v15
                            }
                            _ = fi13
                            _ = v14
                        }
                        if true {
                            v16 := GenS1{F0: 2}
                            fmt.Println(v3, v4)
                            v17 := GenS1{F0: 369}
                            _ = v16
                            _ = v17
                        } else {
                            var v18 int32 = -455
                            v18 = (-(v18))
                            for fi19 := 0; fi19 < 15; fi19++ {
                                if false {
                                    if true {
                                        v20 := GenS1{F0: 42}
                                        _ = v20
                                    } else {
                                        v21 := GenS3{F0: v4}
                                        fmt.Println(v4, fi19, v4)
                                        for fi22 := 0; fi22 < 16; fi22++ {
                                            if true {
                                                if (true && false) {
                                                    v18 = v18
                                                } else {
                                                    if true {
                                                        v3 = v3
                                                        var v23 uint16 = genFn0()
                                                        v18 -= v18
                                                        _ = v23
                                                    } else {
                                                        _ = len(v1)
                                                        v24 := v1[4294967295]
                                                        _ = v24
                                                    }
                                                }
                                            } else {
                                                v25 := GenS3{F0: 718}
                                                switch 128 {
                                                case -9223372036854775808:
                                                    var v26 float64 = float64(1.5)
                                                    _ = v26
                                                case -2147483648:
                                                    v27 := GenS2{F0: 1, F1: 218, F2: true}
                                                    _ = v27
                                                }
                                                v28 := v1[2147483647]
                                                v3 = v3
                                                _ = v25
                                                _ = v28
                                            }
                                            if (false && (true || (fi22 >= (fi19 & fi22)))) {
                                                v3 = -9223372036854775808
                                                var v29 int16 = 1000
                                                var v30 uint8 = 179
                                                _ = v29
                                                _ = v30
                                            } else {
                                                v31 := GenS2{F0: 0, F1: 65535, F2: true}
                                                if (("foo bar" + (("a" + (("hello" + "hello") + "hello")) + (("\n\t" + (("" + ("hello" + "a")) + "a")) + "a"))) >= "hello") {
                                                    fmt.Println(v18)
                                                    var v32 uint16 = 0
                                                    if ((998 <= 2) && (fi19 > (fi22 + fi22))) {
                                                        v33 := GenS2{F0: 2147483647, F1: 423, F2: (true || false)}
                                                        _ = v33
                                                    } else {
                                                        fmt.Println(v3, v32)
                                                        var v34 uint8 = genFn1((fi22 >> ((4294967295) & 63)))
                                                        _ = v34
                                                    }
                                                    _ = v32
                                                } else {
                                                    var v35 uint = 596
                                                    _ = v35
                                                }
                                                _ = v31
                                            }
                                            if ((v3 / ((v3) | 1)) <= v3) {
                                                v3 -= 0
                                                if false {
                                                    if true {
                                                        switch fi22 {
                                                        case 709:
                                                            if true {
                                                                switch "a" {
                                                                case "a":
                                                                    var v36 int32 = v18
                                                                    _ = len(v2)
                                                                    var v37 uint64 = 593
                                                                    _ = v36
                                                                    _ = v37
                                                                case "":
                                                                    v3 = (v3 - (v3 & -732))
                                                                case "hello":
                                                                    v4 += int16(fi19)
                                                                }
                                                                v38 := GenS2{F0: 42, F1: 236, F2: (3.14159 < -2.25)}
                                                                if false {
                                                                    var v39 uint8 = genFn1(590)
                                                                    _ = v39
                                                                }
                                                                _ = v38
                                                            } else {
                                                                v40 := GenS2{F0: 65535, F1: 481, F2: (!false)}
                                                                fmt.Println(v18)
                                                                v41 := GenS4{F0: 1, F1: 0, F2: v4}
                                                                if (0 < 741) {
                                                                    v42 := v1[42]
                                                                    var v43 uint = 42
                                                                    var v44 bool = false
                                                                    _ = v42
                                                                    _ = v43
                                                                    _ = v44
                                                                } else {
                                                                    v45 := GenS5{F0: "foo bar", F1: genFn1(fi19)}
                                                                    _ = v45
                                                                }
                                                                _ = len(v2)
                                                                _ = v40
                                                                _ = v41
                                                            }
                                                        case -168:
                                                            v3 = 42
                                                        }
                                                        if (true && (1.5 < (0.1 + float32(0.1)))) {
                                                            if len(v2) > 0 {
                                                                v2[0] = 369
                                                            }
                                                            var v46 int64 = -1
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v4 -= 1000
                                                            _ = v46
                                                        } else {
                                                            if ((!((fi22 ^ (^(fi22))) <= fi19)) || (int64(fi19) == -669)) {
                                                                if (false && false) {
                                                                    var v47 float64 = float64(1.5)
                                                                    v48 := GenS6{F0: 0, F1: ("foo bar" + ("a" + ""))}
                                                                    _ = v47
                                                                    _ = v48
                                                                } else {
                                                                    var v49 int64 = v3
                                                                    v18 += (v18 << ((1) & 63))
                                                                    _ = v49
                                                                }
                                                                _ = len(v2)
                                                                var v50 uint32 = 65535
                                                                _ = v50
                                                            }
                                                            v3 = v3
                                                        }
                                                    }
                                                    v1[82] = (^(fi19))
                                                } else {
                                                    fmt.Println(fi22, fi19, v3)
                                                    switch (fi19 & fi22) {
                                                    case 32767:
                                                        v51 := v1[65535]
                                                        _ = v51
                                                    case 128:
                                                        switch 32768 {
                                                        case 128:
                                                            v52 := GenS7{F0: 2, F1: (v3 / (((-(v3))) | 1))}
                                                            v4 = (v4 / ((v4) | 1))
                                                            _ = v52
                                                        case -1:
                                                            var v53 uint64 = 0
                                                            _ = v53
                                                        default:
                                                            fmt.Println(v3)
                                                        }
                                                    default:
                                                        switch fi19 {
                                                        case -538:
                                                            v4 = 354
                                                        default:
                                                            v3 = (v3 >> ((690) & 63))
                                                        }
                                                    }
                                                    v4 = (-(v4))
                                                    var v54 int32 = v18
                                                    var v55 float32 = 1e10
                                                    _ = v54
                                                    _ = v55
                                                }
                                                var v56 int32 = -2147483648
                                                _ = v56
                                            } else {
                                                var v57 float64 = float64((1.5 + ((float32(3.14159) - 1e10) - 3.14159)))
                                                v58 := GenS8{F0: false, F1: (v4 % ((v4) | 1))}
                                                v59 := GenS9{F0: -1}
                                                _ = v57
                                                _ = v58
                                                _ = v59
                                            }
                                            _ = fi22
                                        }
                                        _ = v21
                                    }
                                    switch fi19 {
                                    case -2147483648:
                                        if true {
                                            if false {
                                                v1[42] = 32768
                                                fmt.Println(fi19, v3)
                                            } else {
                                                switch (fi19 & (fi19 >> ((111) & 63))) {
                                                case 127:
                                                    switch "\n\t" {
                                                    case "":
                                                        var v60 uint = 65535
                                                        _ = v60
                                                    case "foo bar":
                                                        v3 = int64(fi19)
                                                    default:
                                                        v4 = (v4 % ((v4) | 1))
                                                    }
                                                }
                                                for fi61 := 0; fi61 < 19; fi61++ {
                                                    var v62 int = fi61
                                                    if false {
                                                        v2 = append(v2, 42)
                                                        if true {
                                                            v63 := v1[65535]
                                                            switch fi19 {
                                                            case -836:
                                                                if true {
                                                                    switch (-(fi19)) {
                                                                    case 42:
                                                                        switch ((("\n\t" + ("\n\t" + "\n\t")) + "\n\t") + "") {
                                                                        case "":
                                                                            var v64 int16 = v4
                                                                            _ = v64
                                                                        default:
                                                                            v65 := GenS10{F0: 4294967295}
                                                                            _ = v65
                                                                        }
                                                                    case 2147483647:
                                                                        fmt.Println(v63, v4, fi61)
                                                                        v3 = v3
                                                                        v3 -= v3
                                                                    case 9223372036854775807:
                                                                        v18 += v18
                                                                    }
                                                                    v63 = v62
                                                                    v62 -= (fi61 % (((fi61 << ((2) & 63))) | 1))
                                                                } else {
                                                                    v66 := GenS9{F0: -1}
                                                                    _ = v66
                                                                }
                                                                fmt.Println(v18, fi19, v3)
                                                                var v67 int8 = 1
                                                                _ = v67
                                                            case -1:
                                                                v18 = v18
                                                            default:
                                                                var v68 int16 = (v4 % ((v4) | 1))
                                                                var v69 int64 = v3
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                            _ = v63
                                                        } else {
                                                            var v70 float32 = 0.1
                                                            v71 := GenS11{F0: 969, F1: 2147483647}
                                                            if false {
                                                                var v72 int16 = (v4 | 173)
                                                                switch (("foo bar" + "") + ((("foo bar" + ("a" + "")) + (("\n\t" + "") + (("" + "\n\t") + "\n\t"))) + "foo bar")) {
                                                                case "":
                                                                    _ = len(v2)
                                                                case "a":
                                                                    switch (fi19 ^ (^(v62))) {
                                                                    case 245:
                                                                        var v73 float32 = (1e10 * ((v70 - v70) - v70))
                                                                        _ = v73
                                                                    case -914:
                                                                        v4 = (v4 + (v4 | 42))
                                                                    case 198:
                                                                        if false {
                                                                            delete(v1, 0)
                                                                            if len(v2) > 0 {
                                                                                v2[0] = 505
                                                                            }
                                                                            var v74 int32 = 0
                                                                            _ = v74
                                                                        } else {
                                                                            var v75 uint32 = 4294967295
                                                                            var v76 int16 = v4
                                                                            _ = v75
                                                                            _ = v76
                                                                        }
                                                                    }
                                                                }
                                                                _ = v72
                                                            }
                                                            if true {
                                                                if true {
                                                                    fmt.Println(v70, v3, v18)
                                                                } else {
                                                                    delete(v1, 2)
                                                                    v77 := GenS7{F0: 18446744073709551615, F1: 42}
                                                                    v62 -= 2147483647
                                                                    _ = v77
                                                                }
                                                                switch fi61 {
                                                                case 2147483647:
                                                                    if false {
                                                                        if false {
                                                                            v18 -= 42
                                                                            fmt.Println(fi61, v3, v70)
                                                                            var v78 uint8 = genFn1((fi19 << ((1) & 63)))
                                                                            _ = v78
                                                                        }
                                                                        var v79 int16 = v4
                                                                        _ = v79
                                                                    } else {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = 339
                                                                        }
                                                                        if true {
                                                                            fmt.Println(v18, v70)
                                                                            var v80 int16 = v4
                                                                            _ = v80
                                                                        }
                                                                    }
                                                                }
                                                                if true {
                                                                    var v81 int8 = 42
                                                                    fmt.Println(fi19, fi61, v18)
                                                                    var v82 uint64 = 503
                                                                    v83 := GenS12{F0: 755, F1: v18}
                                                                    _ = v81
                                                                    _ = v82
                                                                    _ = v83
                                                                }
                                                            } else {
                                                                delete(v1, 746)
                                                                var v84 int8 = 42
                                                                fmt.Println(fi61)
                                                                _ = v84
                                                            }
                                                            _ = len(v1)
                                                            fmt.Println(v4, v4, v70)
                                                            var v85 int64 = v3
                                                            if true {
                                                                if (3 != 228) {
                                                                    v62 = v62
                                                                }
                                                                v86 := GenS13{F0: 109, F1: true, F2: -128}
                                                                if false {
                                                                    v87 := GenS14{F0: 2, F1: 18446744073709551615, F2: int(v18)}
                                                                    v88 := v1[4294967295]
                                                                    _ = v87
                                                                    _ = v88
                                                                } else {
                                                                    v18 -= -771
                                                                }
                                                                _ = v86
                                                            } else {
                                                                if false {
                                                                    fmt.Println(v70, fi61, fi61)
                                                                    v89 := v1[65535]
                                                                    var v90 int = (fi61 ^ (-(fi19)))
                                                                    _ = v89
                                                                    _ = v90
                                                                } else {
                                                                    var v91 uint8 = 127
                                                                    v92 := GenS4{F0: (v91 - 1), F1: 65535, F2: 32767}
                                                                    _ = v91
                                                                    _ = v92
                                                                }
                                                                var v93 uint = 18446744073709551615
                                                                _ = v93
                                                            }
                                                            var v94 float64 = 1.5
                                                            _ = v70
                                                            _ = v71
                                                            _ = v85
                                                            _ = v94
                                                        }
                                                    }
                                                    _ = fi61
                                                    _ = v62
                                                }
                                            }
                                            var v95 uint64 = 4294967295
                                            _ = v95
                                        } else {
                                            fmt.Println(v3)
                                        }
                                    default:
                                        var v96 uint = 65535
                                        _ = v96
                                    }
                                    if false {
                                        for fi97 := 0; fi97 < 16; fi97++ {
                                            v98 := GenS15{F0: "\n\t", F1: v3}
                                            _ = fi97
                                            _ = v98
                                        }
                                        var v99 uint64 = 42
                                        _ = v99
                                    } else {
                                        fmt.Println(v4, fi19, v3)
                                    }
                                }
                                _ = fi19
                            }
                            _ = v18
                        }
                        for fi100 := 0; fi100 < 13; fi100++ {
                            v2 = append(v2, 2147483647)
                            _ = fi100
                        }
                    }
                    v101 := GenS6{F0: 1, F1: (("" + "a") + "a")}
                    _ = v101
                    _ = v12
                }
                v3 = (v3 - (v3 << ((65535) & 63)))
            }
            switch 9223372036854775807 {
            case 4294967295:
                v102 := GenS9{F0: 1}
                _ = v102
            case -9223372036854775808:
                var v103 uint16 = 294
                _ = v103
            default:
                fmt.Println(v4)
            }
        } else {
            switch "\n\t" {
            case "foo bar":
                for fi104 := 0; fi104 < 15; fi104++ {
                    v4 = v4
                    switch "foo bar" {
                    case "\n\t":
                        v105 := v1[4294967295]
                        _ = v105
                    case "a":
                        v106 := GenS16{F0: "\n\t", F1: (-(v4)), F2: 4294967295}
                        _ = v106
                    }
                    if (false || (-2147483648 >= -217)) {
                        if (0 > int8(fi104)) {
                            v107 := GenS17{F0: 189, F1: "a", F2: 72}
                            _ = v107
                        }
                        v2 = append(v2, 2)
                    }
                    _ = fi104
                }
            case "a":
                for fi108 := 0; fi108 < 12; fi108++ {
                    var v109 string = "foo bar"
                    for fi110 := 0; fi110 < 20; fi110++ {
                        v3 = 9223372036854775807
                        if (127 > 128) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        _ = fi110
                    }
                    _ = fi108
                    _ = v109
                }
            case "\n\t":
                if false {
                    if true {
                        var v111 string = "foo bar"
                        for fi112 := 0; fi112 < 7; fi112++ {
                            var v113 uint8 = genFn1(631)
                            _ = fi112
                            _ = v113
                        }
                        _ = v111
                    }
                    for fi114 := 0; fi114 < 11; fi114++ {
                        var v115 int64 = (v3 - (v3 - (v3 * 42)))
                        var v116 int8 = 127
                        _ = fi114
                        _ = v115
                        _ = v116
                    }
                } else {
                    var v117 uint = 137
                    v3 += (v3 / ((v3) | 1))
                    var v118 string = "foo bar"
                    _ = v117
                    _ = v118
                }
            default:
                v4 = (v4 >> ((1) & 63))
            }
        }
    default:
        v4 += (v4 | (^(v4)))
    }
    if (true && (885 == 248)) {
        var v119 bool = (!(3.14159 > -2.25))
        v120 := v1[0]
        _ = v119
        _ = v120
    }
    v121 := GenS18{F0: 255, F1: -1, F2: 503}
    switch -423 {
    case -9223372036854775808:
        if ("foo bar" != "a") {
            v122 := GenS7{F0: 1, F1: -1}
            _ = v122
        }
    case 2147483647:
        if (!true) {
            v2 = append(v2, 42)
            v4 = 1000
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case 4294967295:
        if (!true) {
            for fi123 := 0; fi123 < 19; fi123++ {
                v124 := GenS4{F0: 42, F1: 65535, F2: v4}
                for fi125 := 0; fi125 < 8; fi125++ {
                    v126 := GenS19{F0: 831, F1: 2, F2: 2}
                    var v127 int = fi123
                    v128 := GenS12{F0: 739, F1: -1}
                    _ = fi125
                    _ = v126
                    _ = v127
                    _ = v128
                }
                _ = fi123
                _ = v124
            }
        } else {
            fmt.Println(v3)
            if (1 < -1) {
                for fi129 := 0; fi129 < 12; fi129++ {
                    if ((fi129 << ((486) & 63)) > (fi129 | (fi129 >> ((831) & 63)))) {
                        for fi130 := 0; fi130 < 17; fi130++ {
                            var v131 float64 = (0.0 / 1e10)
                            _ = fi130
                            _ = v131
                        }
                        v132 := GenS20{F0: 1, F1: 0, F2: -2147483648}
                        if true {
                            v133 := GenS4{F0: 179, F1: 378, F2: (-(v4))}
                            switch 4294967295 {
                            case 42:
                                if (!true) {
                                    v134 := GenS12{F0: 680, F1: -1}
                                    _ = v134
                                }
                            }
                            switch ("foo bar" + "a") {
                            case "a":
                                var v135 float32 = 1e10
                                _ = v135
                            case "\n\t":
                                for fi136 := 0; fi136 < 17; fi136++ {
                                    var v137 uint = 2
                                    var v138 int32 = 1
                                    _ = fi136
                                    _ = v137
                                    _ = v138
                                }
                            default:
                                var v139 bool = ((v4 & 32767) >= int16(v3))
                                _ = v139
                            }
                            _ = v133
                        } else {
                            switch fi129 {
                            case 9223372036854775807:
                                _ = len(v1)
                            case 4294967295:
                                v1[363] = -453
                            default:
                                for fi140 := 0; fi140 < 7; fi140++ {
                                    if len(v2) > 0 {
                                        v2[0] = 65535
                                    }
                                    v4 -= (v4 + 0)
                                    v141 := GenS12{F0: 1, F1: -364}
                                    _ = fi140
                                    _ = v141
                                }
                            }
                        }
                        _ = v132
                    }
                    _ = fi129
                }
                var v142 float32 = 1e10
                _ = v142
            } else {
                if (!(!(154 > 2147483647))) {
                    v3 = v3
                } else {
                    _ = len(v2)
                }
                var v143 float32 = (0.0 * -2.25)
                var v144 int32 = -1
                _ = v143
                _ = v144
            }
            switch 480 {
            case 2147483647:
                for fi145 := 0; fi145 < 4; fi145++ {
                    v146 := GenS21{F0: (fi145 / ((int(v3)) | 1)), F1: v3, F2: ("\n\t" + ("" + "hello"))}
                    _ = fi145
                    _ = v146
                }
            case 1:
                var v147 uint16 = 2
                _ = v147
            case 127:
                switch 127 {
                case -1:
                    for fi148 := 0; fi148 < 15; fi148++ {
                        delete(v1, 65535)
                        _ = fi148
                    }
                case -2147483648:
                    v4 = v4
                case 32767:
                    var v149 int32 = -779
                    _ = v149
                }
            }
        }
    default:
        v150 := GenS7{F0: 542, F1: v3}
        _ = v150
    }
    if (true && true) {
        for fi151 := 0; fi151 < 11; fi151++ {
            v152 := GenS21{F0: fi151, F1: v3, F2: ("" + "\n\t")}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi151
            _ = v152
        }
        for fi153 := 0; fi153 < 9; fi153++ {
            var v154 int8 = 121
            _ = fi153
            _ = v154
        }
        fmt.Println(v4)
    } else {
        if (genFn0() == 2) {
            var v155 float64 = 0.1
            v4 -= -367
            _ = v155
        }
        var v156 float32 = 1e10
        _ = v156
    }
    v4 += 1000
    for fi157 := 0; fi157 < 9; fi157++ {
        var v158 uint = 42
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if false {
            v158 += 0
        }
        _ = fi157
        _ = v158
    }
    if true {
        var v159 string = "foo bar"
        for fi160 := 0; fi160 < 18; fi160++ {
            _ = len(v2)
            _ = fi160
        }
        _ = v159
    }
    for fi161 := 0; fi161 < 12; fi161++ {
        switch "\n\t" {
        case "hello":
            delete(v1, 2)
        case "":
            var v162 float64 = 0.1
            _ = v162
        default:
            switch "a" {
            case "hello":
                fmt.Println(fi161, v3, v3)
            case "foo bar":
                v4 -= 907
            default:
                v4 = v4
            }
        }
        var v163 int16 = (v4 - v4)
        for fi164 := 0; fi164 < 17; fi164++ {
            switch ("hello" + "foo bar") {
            case "hello":
                switch "a" {
                case "hello":
                    switch (((("a" + "hello") + "hello") + "") + "hello") {
                    case "hello":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "a":
                        var v165 int = (^(fi164))
                        _ = v165
                    case "\n\t":
                        fmt.Println(fi164, fi161)
                    }
                default:
                    fmt.Println(v3, fi164)
                }
            case "foo bar":
                var v166 float64 = float64(float32(float64(3.14159)))
                _ = v166
            default:
                var v167 float32 = 1e10
                var v168 float32 = v167
                var v169 int32 = 2147483647
                _ = v167
                _ = v168
                _ = v169
            }
            v4 += v4
            _ = fi164
        }
        _ = fi161
        _ = v163
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v121
    _ = v2
    _ = v3
    _ = v4
}
