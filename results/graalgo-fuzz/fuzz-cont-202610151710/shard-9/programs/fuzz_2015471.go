package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 bool; F2 bool }
type GenS5 struct { F0 int; F1 uint64; F2 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS8 struct { F0 int32; F1 uint32 }
type GenS9 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS10 struct { F0 string; F1 bool; F2 int8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int; F1 int32; F2 uint16 }
type GenS13 struct { F0 uint32; F1 int32 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 uint8; F1 uint; F2 bool }
type GenS18 struct { F0 uint16; F1 uint64 }
type GenS19 struct { F0 string }
type GenS20 struct { F0 int; F1 uint; F2 int16 }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 uint8 }
func genFn0(p0 uint, p1 bool, p2 uint) int64 {
    var v4 float32 = (3.14159 * -2.25)
    var v5 bool = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -2147483648
    v3 := make(map[uint32]uint32)
    var v6 int64 = -171
    v7 := []int{2147483647, -1, 42}
    if true {
        for fi8 := 0; fi8 < 4; fi8++ {
            _ = len(v3)
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := GenS1{F0: "", F1: 776}
                var v11 int16 = -32768
                _ = fi9
                _ = v10
                _ = v11
            }
            fmt.Println(v2, fi8)
            _ = fi8
        }
        v12 := GenS1{F0: ("" + "hello"), F1: 409}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := v3[967]
            for fi15 := 0; fi15 < 4; fi15++ {
                if ((v2 * v2) < 812) {
                    var v16 int = fi13
                    switch "a" {
                    case "foo bar":
                        if true {
                            switch 42 {
                            case 128:
                                fmt.Println(v2, v14, v2)
                                switch "" {
                                case "":
                                    v17 := GenS1{F0: (("hello" + "\n\t") + ("a" + "")), F1: 412}
                                    _ = v17
                                case "\n\t":
                                    v18 := GenS1{F0: ("\n\t" + ("\n\t" + "")), F1: 1}
                                    _ = v18
                                case "hello":
                                    fmt.Println(v6)
                                }
                                v19 := GenS1{F0: "foo bar", F1: 306}
                                v16 += fi13
                                _ = v19
                            case 1:
                                switch fi15 {
                                case 42:
                                    var v20 string = "\n\t"
                                    _ = v20
                                case -1:
                                    var v21 bool = true
                                    _ = v21
                                case 589:
                                    if ((!(1.5 < -2.25)) && true) {
                                        v14 = (v14 + v14)
                                    } else {
                                        v22 := GenS2{F0: (v14 / (((v14 + v14)) | 1))}
                                        v7 = append(v7, fi15)
                                        if true {
                                            v6 += v6
                                        } else {
                                            if false {
                                                v23 := GenS2{F0: (v14 << ((1) & 63))}
                                                v6 -= v6
                                                if len(v7) > 0 {
                                                    v7[0] = (fi13 - (fi15 / ((fi15) | 1)))
                                                }
                                                var v24 uint = 490
                                                v6 = genFn0((^(v24)), (!(true || false)), 255)
                                                v25 := GenS3{F0: 214}
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                            }
                                            v7 = append(v7, fi15)
                                        }
                                        _ = v22
                                    }
                                default:
                                    switch fi15 {
                                    case -9223372036854775808:
                                        fmt.Println(v16)
                                    case 42:
                                        fmt.Println(v16, fi15, v14)
                                    case 32768:
                                        fmt.Println(fi13, v6)
                                    }
                                }
                            default:
                                var v26 int64 = -380
                                _ = v26
                            }
                        }
                    case "a":
                        if (v14 < 42) {
                            if false {
                                var v27 string = ((("foo bar" + "foo bar") + "a") + "hello")
                                fmt.Println(v27)
                                var v28 uint32 = 1
                                _ = v27
                                _ = v28
                            }
                            switch -982 {
                            case -9223372036854775808:
                                v29 := v1[v16]
                                _ = v29
                            }
                        } else {
                            if true {
                                var v30 uint = 2
                                v14 = v14
                                v31 := GenS4{F0: v2, F1: (!((v16 & (fi15 | fi13)) > 32767)), F2: false}
                                _ = v30
                                _ = v31
                            }
                            fmt.Println(v6)
                            var v32 uint64 = 1
                            delete(v1, 1)
                            _ = v32
                        }
                    default:
                        var v33 uint = 1
                        _ = v33
                    }
                    _ = v16
                }
                var v34 uint8 = 145
                fmt.Println(fi15)
                _ = fi15
                _ = v34
            }
            _ = fi13
            _ = v14
        }
        _ = len(v3)
        v35 := GenS5{F0: 42, F1: 2, F2: 713}
        _ = v35
    }
    for fi36 := 0; fi36 < 15; fi36++ {
        fmt.Println(v2, fi36)
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi36
    }
    var v37 int64 = v6
    var v38 uint = 2
    for fi39 := 0; fi39 < 17; fi39++ {
        if true {
            fmt.Println(v38, v37)
            for fi40 := 0; fi40 < 2; fi40++ {
                var v41 string = ""
                var v42 bool = (!(522 != 858))
                _ = fi40
                _ = v41
                _ = v42
            }
        } else {
            if false {
                switch "hello" {
                case "foo bar":
                    var v43 string = "a"
                    _ = v43
                case "hello":
                    v2 += v2
                case "":
                    switch (fi39 + (fi39 * (fi39 << (((v38 >> ((255) & 63))) & 63)))) {
                    case 128:
                        if true {
                            delete(v3, 475)
                            var v44 string = ""
                            _ = v44
                        }
                    default:
                        if true {
                            switch "" {
                            case "foo bar":
                                for fi45 := 0; fi45 < 9; fi45++ {
                                    var v46 float32 = 0.1
                                    _ = len(v1)
                                    v46 = (-2.25 / v46)
                                    _ = fi45
                                    _ = v46
                                }
                            case "hello":
                                if (false && (!(true || (!(v37 > -1))))) {
                                    var v47 bool = (42 != 954)
                                    _ = v47
                                } else {
                                    if true {
                                        if false {
                                            var v48 float64 = 0.1
                                            for fi49 := 0; fi49 < 9; fi49++ {
                                                v2 = (^(v2))
                                                v38 = v38
                                                _ = fi49
                                            }
                                            _ = v48
                                        }
                                    } else {
                                        fmt.Println(v38, v37)
                                        var v50 string = (((("foo bar" + "foo bar") + ("hello" + "")) + "foo bar") + ((("foo bar" + "hello") + "foo bar") + "\n\t"))
                                        _ = v50
                                    }
                                    v51 := GenS6{F0: 101}
                                    var v52 uint16 = 65535
                                    _ = v51
                                    _ = v52
                                }
                            default:
                                for fi53 := 0; fi53 < 7; fi53++ {
                                    v37 -= 0
                                    switch "foo bar" {
                                    case "hello":
                                        v54 := GenS2{F0: 1}
                                        _ = v54
                                    default:
                                        switch ("" + ("\n\t" + ("\n\t" + (((("" + (("foo bar" + "hello") + "a")) + ("\n\t" + (("hello" + ("foo bar" + "")) + "a"))) + "") + "\n\t")))) {
                                        case "a":
                                            if len(v7) > 0 {
                                                v7[0] = (fi53 ^ 128)
                                            }
                                        }
                                    }
                                    if true {
                                        fmt.Println(fi53)
                                        var v55 bool = (0 > 128)
                                        if false {
                                            if ((404 < 931) && (v55 && (!(!v55)))) {
                                                v37 += v37
                                                var v56 int16 = int16(v6)
                                                _ = v56
                                            } else {
                                                v2 = (v2 % ((v2) | 1))
                                            }
                                            if len(v7) > 0 {
                                                v7[0] = (fi53 << (((^(v38))) & 63))
                                            }
                                            var v57 uint = 590
                                            _ = v57
                                        }
                                        _ = v55
                                    } else {
                                        switch "hello" {
                                        case "\n\t":
                                            var v58 int32 = (v2 >> (((v38 * v38)) & 63))
                                            _ = v58
                                        }
                                        v59 := GenS1{F0: "\n\t", F1: v38}
                                        v37 = v6
                                        _ = v59
                                    }
                                    var v60 uint32 = 286
                                    v61 := v3[v60]
                                    if (true || ((0 < 951) || (0.1 > float32(1.5)))) {
                                        var v62 uint32 = (v61 ^ (v61 + 2))
                                        if true {
                                            var v63 bool = true
                                            if v63 {
                                                v64 := GenS7{F0: 42, F1: v62, F2: (v37 & (v6 + genFn0((^(v38)), v63, 457)))}
                                                var v65 float64 = float64((float32(-2.25) - (-2.25 - 1.5)))
                                                _ = v64
                                                _ = v65
                                            } else {
                                                var v66 uint8 = 127
                                                if v63 {
                                                    if (v63 && true) {
                                                        _ = len(v7)
                                                        if ((^(fi39)) == (fi39 << ((915) & 63))) {
                                                            v67 := GenS8{F0: 42, F1: v61}
                                                            if ((42 != (v38 + v38)) || false) {
                                                                v66 -= (v66 >> ((v38) & 63))
                                                                v62 = (^(v60))
                                                                switch ((("" + "foo bar") + "foo bar") + "\n\t") {
                                                                case "a":
                                                                    _ = len(v1)
                                                                case "":
                                                                    v60 = (v60 ^ (v60 * v60))
                                                                default:
                                                                    if (-32768 == int16(v37)) {
                                                                        fmt.Println(fi53, v61)
                                                                    } else {
                                                                        var v68 uint16 = 0
                                                                        switch ("foo bar" + "foo bar") {
                                                                        case "foo bar":
                                                                            v69 := GenS2{F0: (v60 * (v61 + (v62 >> (((^(v38))) & 63))))}
                                                                            _ = v69
                                                                        case "\n\t":
                                                                            var v70 int8 = 1
                                                                            _ = v70
                                                                        default:
                                                                            v71 := GenS4{F0: (v2 % (((v2 << ((v38) & 63))) | 1)), F1: v63, F2: (!v63)}
                                                                            _ = v71
                                                                        }
                                                                        if (1 <= 0) {
                                                                            var v72 uint8 = 255
                                                                            v1[127] = ((("" + "foo bar") + ("a" + "")) <= "foo bar")
                                                                            v62 = (^(v62))
                                                                            v2 += v2
                                                                            if (true && (-42 > -6)) {
                                                                                switch ("" + "\n\t") {
                                                                                case "a":
                                                                                    _ = len(v7)
                                                                                    fmt.Println(v37, v37, v2)
                                                                                    var v73 int16 = -1
                                                                                    _ = v73
                                                                                case "":
                                                                                    fmt.Println(v63)
                                                                                }
                                                                            } else {
                                                                                var v74 int16 = 32767
                                                                                var v75 float64 = 1.5
                                                                                if ((fi39 * fi39) > (fi39 & fi39)) {
                                                                                    var v76 float32 = 1e10
                                                                                    if len(v7) > 0 {
                                                                                        v7[0] = (fi39 ^ (fi39 % (((^(fi39))) | 1)))
                                                                                    }
                                                                                    if v63 {
                                                                                        var v77 uint32 = v60
                                                                                        _ = v77
                                                                                    } else {
                                                                                        fmt.Println(v72)
                                                                                        var v78 uint = (v38 * uint(v68))
                                                                                        _ = v78
                                                                                    }
                                                                                    v79 := GenS5{F0: fi39, F1: 440, F2: 2}
                                                                                    _ = v76
                                                                                    _ = v79
                                                                                } else {
                                                                                    if v63 {
                                                                                        v63 = (v37 == genFn0(v38, (-48 <= -128), (v38 - v38)))
                                                                                    } else {
                                                                                        v7 = append(v7, fi39)
                                                                                        v80 := GenS9{F0: -1, F1: 2, F2: v68}
                                                                                        v62 += 425
                                                                                        _ = v80
                                                                                    }
                                                                                    v7 = append(v7, int(v38))
                                                                                }
                                                                                if v63 {
                                                                                    v6 -= v6
                                                                                    v38 -= 18446744073709551615
                                                                                    _ = len(v7)
                                                                                } else {
                                                                                    var v81 int = fi39
                                                                                    if true {
                                                                                        v60 += 1
                                                                                        if true {
                                                                                            v37 -= v6
                                                                                        }
                                                                                        fmt.Println(v62, v37)
                                                                                        var v82 int32 = int32(v68)
                                                                                        _ = v82
                                                                                    }
                                                                                    _ = v81
                                                                                }
                                                                                _ = v74
                                                                                _ = v75
                                                                            }
                                                                            var v83 int64 = v37
                                                                            switch int(v68) {
                                                                            case 4294967295:
                                                                                var v84 int16 = 32767
                                                                                if (126 == 42) {
                                                                                    var v85 int64 = genFn0((v38 & (v38 ^ 1)), (128 == v72), (^(v38)))
                                                                                    _ = v85
                                                                                } else {
                                                                                    fmt.Println(fi39, v83)
                                                                                }
                                                                                _ = v84
                                                                            }
                                                                            _ = v72
                                                                            _ = v83
                                                                        }
                                                                        _ = v68
                                                                    }
                                                                }
                                                                v63 = true
                                                                v2 = (v2 % (((v2 & v2)) | 1))
                                                            } else {
                                                                var v86 int64 = (v6 ^ v6)
                                                                fmt.Println(fi39, v60, fi53)
                                                                _ = v86
                                                            }
                                                            if (v63 || v63) {
                                                                v87 := GenS1{F0: "\n\t", F1: v38}
                                                                if true {
                                                                    var v88 uint16 = 65535
                                                                    if v63 {
                                                                        v89 := GenS6{F0: 127}
                                                                        if v63 {
                                                                            v90 := GenS8{F0: v2, F1: (v61 | 1)}
                                                                            v1[(fi39 * (-(fi39)))] = (("a" + "foo bar") != (("" + "hello") + "hello"))
                                                                            v91 := GenS10{F0: "\n\t", F1: v63, F2: 127}
                                                                            v92 := GenS3{F0: (v66 << (((v38 % (((v38 ^ 228)) | 1))) & 63))}
                                                                            v93 := GenS2{F0: (^(v62))}
                                                                            _ = v90
                                                                            _ = v91
                                                                            _ = v92
                                                                            _ = v93
                                                                        } else {
                                                                            if (float32(1e10) > 0.1) {
                                                                                var v94 float64 = 1.5
                                                                                v95 := GenS7{F0: (v88 - 32767), F1: (^(v62)), F2: v6}
                                                                                v96 := GenS11{F0: 7}
                                                                                _ = v94
                                                                                _ = v95
                                                                                _ = v96
                                                                            }
                                                                            var v97 int = (fi39 / ((fi53) | 1))
                                                                            v98 := GenS12{F0: fi53, F1: (v2 * v2), F2: 1}
                                                                            _ = v97
                                                                            _ = v98
                                                                        }
                                                                        var v99 int = fi39
                                                                        var v100 int8 = -1
                                                                        _ = v100
                                                                        _ = v89
                                                                        _ = v99
                                                                    } else {
                                                                        if v63 {
                                                                            var v101 uint64 = 4294967295
                                                                            v3[(v61 >> ((v38) & 63))] = v62
                                                                            _ = v101
                                                                        } else {
                                                                            if ((v37 * (v6 ^ v6)) != genFn0(v38, v63, (^(v38)))) {
                                                                                var v102 int32 = -626
                                                                                _ = v102
                                                                            }
                                                                            if v63 {
                                                                                var v103 uint16 = v88
                                                                                v6 = v6
                                                                                _ = v103
                                                                            }
                                                                        }
                                                                        v3[v61] = (v61 * (v60 << ((v38) & 63)))
                                                                        switch (fi39 * (fi53 << ((v38) & 63))) {
                                                                        case -2147483648:
                                                                            var v104 uint = 42
                                                                            if (v60 >= (v62 ^ v62)) {
                                                                                var v105 string = ((("foo bar" + "a") + "") + ("a" + ("foo bar" + "foo bar")))
                                                                                v106 := GenS12{F0: fi53, F1: v2, F2: (v88 / ((1) | 1))}
                                                                                _ = v105
                                                                                _ = v106
                                                                            } else {
                                                                                var v107 int16 = 0
                                                                                if false {
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        var v108 int8 = 42
                                                                                        _ = v108
                                                                                    case "":
                                                                                        var v109 uint64 = 186
                                                                                        _ = v109
                                                                                    case "foo bar":
                                                                                        var v110 string = "a"
                                                                                        _ = v110
                                                                                    default:
                                                                                        v2 += (v2 & int32(fi39))
                                                                                        _ = len(v7)
                                                                                        var v111 float32 = (3.14159 / 1e10)
                                                                                        v112 := GenS13{F0: v60, F1: (^(v2))}
                                                                                        v113 := GenS11{F0: -64}
                                                                                        _ = v111
                                                                                        _ = v112
                                                                                        _ = v113
                                                                                    }
                                                                                    v61 += (v61 << ((uint(v61)) & 63))
                                                                                } else {
                                                                                    var v114 int16 = v107
                                                                                    v104 -= uint(v88)
                                                                                    _ = v114
                                                                                }
                                                                                _ = v107
                                                                            }
                                                                            v61 = v62
                                                                            _ = v104
                                                                        case -9223372036854775808:
                                                                            if (-2.25 > 0.0) {
                                                                                v63 = ((v62 / (((v62 % (((v60 * (v62 / ((v60) | 1)))) | 1))) | 1)) >= (v60 / (((v60 / ((2) | 1))) | 1)))
                                                                            }
                                                                        }
                                                                        var v115 int32 = v2
                                                                        if false {
                                                                            if (v66 != 255) {
                                                                                v61 -= (v62 % ((v62) | 1))
                                                                                if false {
                                                                                    v2 += v115
                                                                                } else {
                                                                                    v116 := GenS2{F0: (v62 - (^(v60)))}
                                                                                    _ = v116
                                                                                }
                                                                                v117 := GenS14{F0: (fi39 & 346)}
                                                                                _ = v117
                                                                            }
                                                                            v63 = v63
                                                                            fmt.Println(v60, v63, v2)
                                                                            delete(v3, v62)
                                                                            fmt.Println(v62)
                                                                            v63 = ((("" + "") + "\n\t") < ("" + "\n\t"))
                                                                        } else {
                                                                            v118 := GenS2{F0: v60}
                                                                            var v119 uint32 = v60
                                                                            _ = v118
                                                                            _ = v119
                                                                        }
                                                                        _ = v115
                                                                    }
                                                                    v120 := GenS15{F0: ((("hello" + "a") + "hello") + "\n\t")}
                                                                    _ = v120
                                                                    _ = v88
                                                                } else {
                                                                    var v121 uint64 = 394
                                                                    _ = v121
                                                                }
                                                                _ = v87
                                                            } else {
                                                                _ = len(v7)
                                                                if false {
                                                                    v62 -= 2
                                                                    var v122 uint8 = (^(v66))
                                                                    v123 := GenS16{F0: uint8(v62)}
                                                                    var v124 float32 = 0.1
                                                                    _ = v122
                                                                    _ = v123
                                                                    _ = v124
                                                                } else {
                                                                    if len(v7) >= 2 {
                                                                        v7 = v7[1:len(v7)]
                                                                    }
                                                                    if (0 >= 42) {
                                                                        switch 542 {
                                                                        case -687:
                                                                            v125 := GenS3{F0: (^(v66))}
                                                                            _ = v125
                                                                        case -881:
                                                                            var v126 int64 = v6
                                                                            _ = v126
                                                                        case 799:
                                                                            fmt.Println(v37, v6)
                                                                        }
                                                                        v62 = v61
                                                                    } else {
                                                                        var v127 float32 = 1.5
                                                                        fmt.Println(fi39)
                                                                        _ = v127
                                                                    }
                                                                    v38 += v38
                                                                    if true {
                                                                        v128 := GenS17{F0: 12, F1: (v38 + v38), F2: v63}
                                                                        v129 := v1[fi53]
                                                                        _ = v128
                                                                        _ = v129
                                                                    }
                                                                }
                                                            }
                                                            _ = v67
                                                        } else {
                                                            switch (fi53 - fi39) {
                                                            case 128:
                                                                var v130 int = -2147483648
                                                                _ = v130
                                                            case 723:
                                                                var v131 float64 = 3.14159
                                                                _ = v131
                                                            case 4294967295:
                                                                _ = len(v3)
                                                            default:
                                                                var v132 bool = (276 <= 1)
                                                                if v63 {
                                                                    _ = len(v3)
                                                                    v38 -= (v38 + (^(v38)))
                                                                } else {
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        v60 = (v60 >> ((v38) & 63))
                                                                    default:
                                                                        if v132 {
                                                                            fmt.Println(fi53, v38, v66)
                                                                            switch 32768 {
                                                                            case -720:
                                                                                var v133 int = (fi53 << ((4294967295) & 63))
                                                                                _ = v133
                                                                            case 32768:
                                                                                v134 := GenS13{F0: (^(v62)), F1: 836}
                                                                                _ = v134
                                                                            default:
                                                                                v61 = (v61 - (v61 >> ((v38) & 63)))
                                                                            }
                                                                            var v135 string = "foo bar"
                                                                            _ = v135
                                                                        }
                                                                    }
                                                                    fmt.Println(v38)
                                                                }
                                                                _ = v132
                                                            }
                                                            fmt.Println(v60, v2)
                                                            var v136 string = "hello"
                                                            _ = v136
                                                        }
                                                    } else {
                                                        fmt.Println(v66, fi39)
                                                        switch "a" {
                                                        case "foo bar":
                                                            var v137 uint64 = 1
                                                            _ = v137
                                                        case "a":
                                                            if false {
                                                                switch ("" + "foo bar") {
                                                                case "hello":
                                                                    v138 := GenS6{F0: 42}
                                                                    _ = v138
                                                                case "":
                                                                    fmt.Println(v6)
                                                                default:
                                                                    var v139 uint = (^(v38))
                                                                    _ = v139
                                                                }
                                                            }
                                                        default:
                                                            if (!(v63 && false)) {
                                                                v140 := GenS12{F0: fi53, F1: v2, F2: 32767}
                                                                _ = v140
                                                            }
                                                        }
                                                        var v141 uint32 = (^(v61))
                                                        _ = v141
                                                    }
                                                    v3[950] = 106
                                                    if (true || v63) {
                                                        v62 = v61
                                                        switch "hello" {
                                                        case "":
                                                            v37 -= 0
                                                        }
                                                        v142 := GenS6{F0: -1}
                                                        _ = v142
                                                    } else {
                                                        fmt.Println(v63, v37, v62)
                                                        var v143 uint64 = 2
                                                        switch (fi53 * (fi53 | (fi53 - fi39))) {
                                                        case 767:
                                                            fmt.Println(fi39, v6, v38)
                                                        }
                                                        _ = v143
                                                    }
                                                    switch fi39 {
                                                    case 9223372036854775807:
                                                        _ = len(v7)
                                                    case 127:
                                                        switch "a" {
                                                        case "a":
                                                            v7 = append(v7, (fi39 | -385))
                                                        case "foo bar":
                                                            v144 := GenS6{F0: 1}
                                                            _ = v144
                                                        case "hello":
                                                            if v63 {
                                                                _ = len(v7)
                                                            }
                                                        }
                                                    case 42:
                                                        v63 = false
                                                    }
                                                    var v145 uint32 = v60
                                                    v6 = (v6 ^ 9223372036854775807)
                                                    v146 := GenS2{F0: v62}
                                                    _ = v145
                                                    _ = v146
                                                } else {
                                                    fmt.Println(fi39)
                                                }
                                                v38 = v38
                                                _ = v66
                                            }
                                            _ = v63
                                        } else {
                                            if (v38 != v38) {
                                                fmt.Println(v61, v61)
                                                if false {
                                                    v2 = v2
                                                    if (-2.25 != -2.25) {
                                                        v147 := GenS16{F0: 128}
                                                        _ = len(v1)
                                                        _ = v147
                                                    }
                                                    fmt.Println(v37)
                                                }
                                                v1[fi39] = false
                                            } else {
                                                var v148 uint = v38
                                                _ = v148
                                            }
                                            if false {
                                                fmt.Println(v61, v60, v38)
                                                v6 -= (v37 ^ v37)
                                            } else {
                                                v37 = genFn0(v38, (!false), v38)
                                                var v149 int16 = 0
                                                _ = v149
                                            }
                                        }
                                        var v150 uint16 = 32767
                                        var v151 float64 = 3.14159
                                        v62 -= 65535
                                        _ = v150
                                        _ = v151
                                        _ = v62
                                    } else {
                                        if true {
                                            var v152 uint32 = (v60 * 934)
                                            fmt.Println(v38)
                                            var v153 int64 = int64(v2)
                                            _ = v152
                                            _ = v153
                                        }
                                    }
                                    _ = fi53
                                    _ = v60
                                    _ = v61
                                }
                            }
                            switch "" {
                            case "\n\t":
                                var v154 uint8 = 1
                                _ = v154
                            case "foo bar":
                                if true {
                                    for fi155 := 0; fi155 < 5; fi155++ {
                                        fmt.Println(v38, fi155)
                                        v6 = v6
                                        v38 = v38
                                        _ = fi155
                                    }
                                } else {
                                    switch ("\n\t" + "") {
                                    case "\n\t":
                                        for fi156 := 0; fi156 < 12; fi156++ {
                                            v3[312] = 389
                                            _ = fi156
                                        }
                                    case "foo bar":
                                        v6 -= v6
                                    case "":
                                        if ((v2 >> (((^(v38))) & 63)) < v2) {
                                            v6 = (v6 & v37)
                                            fmt.Println(fi39)
                                        }
                                    default:
                                        fmt.Println(fi39, v37, v6)
                                    }
                                }
                            case "":
                                var v157 uint = (v38 & 2)
                                _ = v157
                            default:
                                if true {
                                    fmt.Println(v6, v37, v6)
                                    var v158 int64 = 0
                                    _ = v158
                                } else {
                                    switch fi39 {
                                    case 127:
                                        if false {
                                            if true {
                                                v3[2] = 761
                                                if true {
                                                    fmt.Println(v38, fi39, v38)
                                                } else {
                                                    fmt.Println(fi39, v2, v38)
                                                }
                                                for fi159 := 0; fi159 < 2; fi159++ {
                                                    v6 = -9223372036854775808
                                                    v37 = (v6 % (((v37 * (-(v37)))) | 1))
                                                    v38 = 65535
                                                    _ = fi159
                                                }
                                            }
                                        }
                                    case 9223372036854775807:
                                        for fi160 := 0; fi160 < 12; fi160++ {
                                            var v161 uint64 = 649
                                            v161 = (v161 << ((v38) & 63))
                                            _ = fi160
                                            _ = v161
                                        }
                                    }
                                    v38 = (v38 | 18446744073709551615)
                                }
                            }
                            if ((v37 / ((1) | 1)) == -1) {
                                var v162 string = "a"
                                _ = v162
                            } else {
                                var v163 uint64 = 18446744073709551615
                                _ = v163
                            }
                        } else {
                            v164 := GenS18{F0: 772, F1: 18446744073709551615}
                            if true {
                                var v165 uint16 = 65535
                                var v166 float64 = -2.25
                                v38 += 781
                                _ = v165
                                _ = v166
                            } else {
                                var v167 int8 = -128
                                _ = len(v1)
                                _ = v167
                            }
                            var v168 uint = 640
                            _ = v164
                            _ = v168
                        }
                    }
                }
                fmt.Println(v37, v2, v38)
                v6 = 0
            } else {
                v169 := GenS3{F0: 0}
                var v170 int8 = 13
                switch (-(fi39)) {
                case 0:
                    for fi171 := 0; fi171 < 7; fi171++ {
                        switch (^(fi39)) {
                        case 1:
                            var v172 string = "foo bar"
                            v170 += -128
                            var v173 string = ("foo bar" + v172)
                            _ = v172
                            _ = v173
                        case 32768:
                            var v174 uint16 = 0
                            _ = v174
                        default:
                            v6 -= v37
                        }
                        v6 -= (v6 * (v6 % ((genFn0(18446744073709551615, (32767 < 810), v38)) | 1)))
                        _ = fi171
                    }
                case 1:
                    for fi175 := 0; fi175 < 10; fi175++ {
                        v176 := GenS19{F0: ""}
                        if ((fi39 & fi175) < (fi175 ^ 32768)) {
                            v37 = (v37 ^ (v37 << (((^(v38))) & 63)))
                            v177 := GenS16{F0: 0}
                            var v178 float64 = (float64((0.0 - 1.5)) + 0.1)
                            _ = v177
                            _ = v178
                        } else {
                            v38 -= (v38 ^ (^(v38)))
                            var v179 int16 = -1
                            _ = v179
                        }
                        switch (fi39 & 695) {
                        case 127:
                            if ((((0.1 - ((0.1 + 1e10) * -2.25)) / float32((-2.25 - 1e10))) + (0.1 - 1e10)) <= 1.5) {
                                var v180 float64 = -2.25
                                fmt.Println(v170, v6)
                                _ = v180
                            }
                            v181 := GenS11{F0: v170}
                            v2 -= (v2 << (((^(v38))) & 63))
                            _ = v181
                        case -935:
                            v2 = (v2 / ((v2) | 1))
                        case 128:
                            v38 += v38
                        default:
                            var v182 int = (fi39 << (((v38 + v38)) & 63))
                            _ = v182
                        }
                        var v183 bool = false
                        _ = fi175
                        _ = v176
                        _ = v183
                    }
                default:
                    v184 := GenS20{F0: fi39, F1: 968, F2: 1}
                    _ = v184
                }
                _ = v169
                _ = v170
            }
        }
        if (0.0 == 0.0) {
            var v185 float64 = 3.14159
            var v186 uint16 = 1
            v187 := GenS21{F0: -1}
            _ = v185
            _ = v186
            _ = v187
        } else {
            var v188 int8 = 0
            var v189 uint = (v38 << ((2) & 63))
            var v190 uint64 = 41
            _ = v188
            _ = v189
            _ = v190
        }
        _ = fi39
    }
    v191 := GenS22{F0: 207}
    var v192 uint = v38
    var v193 bool = true
    var v194 string = "hello"
    v2 = (^(v2))
    for fi195 := 0; fi195 < 13; fi195++ {
        if false {
            switch (fi195 ^ -2147483648) {
            case 0:
                for fi196 := 0; fi196 < 16; fi196++ {
                    v192 -= 42
                    v194 = v194
                    _ = fi196
                }
            case -9223372036854775808:
                for fi197 := 0; fi197 < 9; fi197++ {
                    v194 = v194
                    var v198 uint8 = 0
                    if false {
                        _ = len(v3)
                        switch int(v38) {
                        case -2147483648:
                            fmt.Println(v37, v6)
                        }
                        if (!(!(230 == 1))) {
                            fmt.Println(v193, v193, v194)
                            v38 -= v38
                            if ((-2.25 - -2.25) == ((((-2.25 - 0.0) * (float32(0.0) + -2.25)) / ((1.5 - 0.0) + 1.5)) - 0.1)) {
                                v37 = v37
                                fmt.Println(v37)
                                var v199 bool = true
                                _ = v199
                            }
                        }
                    }
                    _ = fi197
                    _ = v198
                }
            case 128:
                v200 := GenS5{F0: (fi195 << ((v38) & 63)), F1: uint64(v192), F2: 42}
                _ = v200
            default:
                for fi201 := 0; fi201 < 14; fi201++ {
                    _ = len(v7)
                    _ = fi201
                }
            }
        }
        v192 = v192
        v202 := GenS17{F0: 128, F1: 0, F2: v193}
        _ = fi195
        _ = v202
    }
    if v193 {
        v1[9223372036854775807] = false
        var v203 uint64 = 1
        _ = v203
    }
    v192 = 566
    v7 = append(v7, -352)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v192:", v192)
    fmt.Println("v193:", v193)
    fmt.Println("v194:", v194)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v191
    _ = v192
    _ = v193
    _ = v194
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v6
    _ = v7
}
