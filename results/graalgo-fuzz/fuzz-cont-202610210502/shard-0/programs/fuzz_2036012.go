package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int64 }
type GenS2 struct { F0 bool; F1 uint8; F2 bool }
type GenS3 struct { F0 int8; F1 string; F2 uint64 }
type GenS4 struct { F0 bool; F1 int64; F2 uint64 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS8 struct { F0 int16; F1 uint8; F2 int }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int8; F1 bool }
type GenS11 struct { F0 uint32; F1 bool }
type GenS12 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS13 struct { F0 uint; F1 bool; F2 int }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 uint8; F1 uint32; F2 uint8 }
func genFn0(p0 uint32) uint {
    var v7 uint8 = 42
    var v8 bool = true
    fmt.Println(v8)
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 0
    v3 := []uint32{4294967295, 975}
    v4 := []string{"foo bar"}
    v5 := make(map[string]string)
    v6 := []uint32{285, 0, 65535, 2}
    for fi9 := 0; fi9 < 12; fi9++ {
        switch fi9 {
        case 42:
            v10 := v1[2]
            _ = v10
        }
        _ = fi9
    }
    var v11 string = "foo bar"
    var v12 uint32 = 0
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 17; fi14++ {
            fmt.Println(v11, v12)
            v15 := GenS1{F0: 303, F1: -32768, F2: -542}
            var v16 bool = false
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    if false {
        var v17 uint = v2
        if true {
            var v18 string = v11
            var v19 int16 = 32767
            for fi20 := 0; fi20 < 19; fi20++ {
                delete(v5, v18)
                v12 = (v12 & uint32(v2))
                v2 += (v17 >> ((genFn0(v12)) & 63))
                _ = fi20
            }
            _ = v18
            _ = v19
        }
        for fi21 := 0; fi21 < 16; fi21++ {
            for fi22 := 0; fi22 < 20; fi22++ {
                var v23 string = v11
                var v24 uint64 = 42
                _ = fi22
                _ = v23
                _ = v24
            }
            if (68 <= -1) {
                for fi25 := 0; fi25 < 13; fi25++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    if (127 >= 0) {
                        v17 += 4294967295
                        fmt.Println(v17)
                        v26 := GenS1{F0: 975, F1: 1, F2: -875}
                        v12 = 2
                        _ = v26
                    }
                    _ = fi25
                }
                if false {
                    if false {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                } else {
                    fmt.Println(v2, v2, v2)
                    var v27 string = (((v11 + v11) + v11) + (v11 + (v11 + v11)))
                    _ = v27
                }
            } else {
                v28 := v1[(v12 | (v12 << ((v17) & 63)))]
                v11 = v11
                _ = v28
            }
            _ = fi21
        }
        _ = v17
    } else {
        v29 := GenS2{F0: true, F1: 42, F2: true}
        if true {
            v2 = v2
            var v30 uint64 = 828
            if false {
                switch ("a" + v11) {
                case "foo bar":
                    var v31 float32 = 1e10
                    _ = v31
                case "\n\t":
                    switch 4294967295 {
                    case 42:
                        if (232 < 2) {
                            for fi32 := 0; fi32 < 11; fi32++ {
                                if (9223372036854775807 != -2147483648) {
                                    if (!(!true)) {
                                        if ((0.1 > 3.14159) && (((v12 > uint32(v30)) && ((v2 >> ((v2) & 63)) >= (v2 + v2))) || true)) {
                                            if (v12 != 1) {
                                                switch v11 {
                                                case "a":
                                                    var v33 float64 = ((3.14159 / 1.5) + (0.0 + ((3.14159 + 0.1) + float64((-2.25 - float32((1e10 + 0.1)))))))
                                                    _ = v33
                                                case "":
                                                    v11 = ("a" + v11)
                                                case "foo bar":
                                                    v30 -= (^(v30))
                                                }
                                                var v34 float32 = 0.0
                                                _ = v34
                                            } else {
                                                switch (fi32 / ((127) | 1)) {
                                                case 533:
                                                    switch 32768 {
                                                    case -2147483648:
                                                        fmt.Println(fi32)
                                                    default:
                                                        switch (-(fi32)) {
                                                        case 727:
                                                            v35 := GenS3{F0: 33, F1: v11, F2: (^(v30))}
                                                            _ = v35
                                                        case 32768:
                                                            v36 := GenS4{F0: true, F1: -9223372036854775808, F2: v30}
                                                            _ = v36
                                                        case 9223372036854775807:
                                                            var v37 bool = false
                                                            _ = v37
                                                        }
                                                    }
                                                default:
                                                    v12 = (^(v12))
                                                }
                                            }
                                            if ((-9223372036854775808 >= -9223372036854775808) && (int64(v30) <= 42)) {
                                                for fi38 := 0; fi38 < 9; fi38++ {
                                                    fmt.Println(fi32)
                                                    switch (fi38 % ((0) | 1)) {
                                                    case 32768:
                                                        fmt.Println(fi32, v12, v2)
                                                    case 128:
                                                        if false {
                                                            switch "\n\t" {
                                                            case "a":
                                                                if true {
                                                                    var v39 int = (fi32 ^ (fi32 ^ -9223372036854775808))
                                                                    var v40 uint16 = uint16(v30)
                                                                    v11 = (("\n\t" + v11) + v11)
                                                                    _ = v39
                                                                    _ = v40
                                                                } else {
                                                                    var v41 bool = (!(false && true))
                                                                    _ = v41
                                                                }
                                                            case "foo bar":
                                                                v2 = (v2 * v2)
                                                            case "hello":
                                                                if false {
                                                                    var v42 uint = v2
                                                                    if (!true) {
                                                                        v43 := GenS1{F0: (v30 & (v30 ^ (v30 + (v30 % ((v30) | 1))))), F1: -1, F2: 0}
                                                                        v42 += genFn0(v12)
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        _ = v43
                                                                    } else {
                                                                        var v44 uint64 = v30
                                                                        _ = v44
                                                                    }
                                                                    fmt.Println(v2)
                                                                    _ = v42
                                                                } else {
                                                                    v2 += v2
                                                                    v45 := v5[v11]
                                                                    _ = v45
                                                                }
                                                            default:
                                                                switch fi32 {
                                                                case 128:
                                                                    if false {
                                                                        if false {
                                                                            var v46 float32 = float32(3.14159)
                                                                            _ = v46
                                                                        }
                                                                        switch (("foo bar" + v11) + (v11 + v11)) {
                                                                        case "\n\t":
                                                                            v47 := v1[4294967295]
                                                                            _ = v47
                                                                        case "foo bar":
                                                                            var v48 uint = v2
                                                                            _ = v48
                                                                        case "a":
                                                                            v30 += (v30 >> ((v2) & 63))
                                                                        default:
                                                                            var v49 bool = (true || (true && (0 > genFn0(65535))))
                                                                            _ = v49
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            v30 = v30
                                                                        }
                                                                        delete(v5, v11)
                                                                        v11 = v11
                                                                    }
                                                                }
                                                            }
                                                            var v50 uint32 = 2
                                                            if (((0.1 + -2.25) - float32(float64(0.1))) < ((3.14159 + (((3.14159 * 0.1) / 1.0) * (3.14159 * float32(-2.25)))) / -2.25)) {
                                                                var v51 uint64 = v30
                                                                _ = v51
                                                            } else {
                                                                fmt.Println(v11, fi38)
                                                            }
                                                            _ = v50
                                                        }
                                                    }
                                                    switch fi38 {
                                                    case -721:
                                                        v12 -= v12
                                                    case -1:
                                                        if false {
                                                            v11 = "\n\t"
                                                            v11 = ((("a" + v11) + "") + (v11 + v11))
                                                            v1[v12] = (fi32 ^ (fi32 - (fi38 + fi32)))
                                                            if (!true) {
                                                                var v52 float64 = 0.0
                                                                v53 := GenS3{F0: 127, F1: v11, F2: v30}
                                                                _ = v52
                                                                _ = v53
                                                            } else {
                                                                fmt.Println(fi32)
                                                                if len(v4) > 0 {
                                                                    v4[0] = v11
                                                                }
                                                                var v54 int8 = 42
                                                                _ = v54
                                                            }
                                                        } else {
                                                            v55 := GenS3{F0: -58, F1: ("" + "hello"), F2: v30}
                                                            v56 := GenS5{F0: 42, F1: 128}
                                                            fmt.Println(v2, fi32)
                                                            v57 := GenS6{F0: -953}
                                                            _ = v55
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                    }
                                                    _ = fi38
                                                }
                                                var v58 bool = (((4294967295 != fi32) || (!(!(!((true && false) || (v11 != "\n\t")))))) && (false || true))
                                                var v59 int32 = -1
                                                _ = v58
                                                _ = v59
                                            } else {
                                                if false {
                                                    v60 := GenS5{F0: 127, F1: 255}
                                                    switch "hello" {
                                                    case "foo bar":
                                                        if len(v6) > 0 {
                                                            v6[0] = (v12 | v12)
                                                        }
                                                    default:
                                                        v12 = v12
                                                    }
                                                    _ = v60
                                                } else {
                                                    fmt.Println(v2, v30, v11)
                                                    if true {
                                                        if (0.1 != 0.1) {
                                                            if false {
                                                                for fi61 := 0; fi61 < 14; fi61++ {
                                                                    switch (("\n\t" + v11) + ("" + v11)) {
                                                                    case "hello":
                                                                        fmt.Println(v12, v12)
                                                                    case "foo bar":
                                                                        v62 := GenS7{F0: 50, F1: uint16(v30), F2: 255}
                                                                        _ = v62
                                                                    }
                                                                    switch v11 {
                                                                    case "a":
                                                                        fmt.Println(fi32, v11)
                                                                    case "\n\t":
                                                                        delete(v1, 56)
                                                                    case "hello":
                                                                        var v63 int16 = -1
                                                                        _ = v63
                                                                    }
                                                                    _ = fi61
                                                                }
                                                            }
                                                            v64 := GenS2{F0: (!(true || false)), F1: uint8(v2), F2: true}
                                                            _ = len(v6)
                                                            _ = v64
                                                        }
                                                        var v65 uint8 = 0
                                                        var v66 float32 = 0.1
                                                        _ = v65
                                                        _ = v66
                                                    } else {
                                                        var v67 uint8 = 188
                                                        v68 := GenS5{F0: -59, F1: v67}
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                }
                                            }
                                            for fi69 := 0; fi69 < 20; fi69++ {
                                                var v70 int32 = 42
                                                v12 = (^(v12))
                                                _ = fi69
                                                _ = v70
                                            }
                                        } else {
                                            fmt.Println(v12, v11, fi32)
                                            v71 := GenS2{F0: true, F1: 255, F2: (-367 >= -9223372036854775808)}
                                            _ = v71
                                        }
                                        var v72 uint64 = v30
                                        _ = v72
                                    } else {
                                        v11 = (v11 + "\n\t")
                                        v73 := v5["hello"]
                                        if true {
                                            var v74 uint32 = v12
                                            _ = v74
                                        } else {
                                            _ = len(v3)
                                            for fi75 := 0; fi75 < 6; fi75++ {
                                                v2 += v2
                                                _ = fi75
                                            }
                                        }
                                        _ = v73
                                    }
                                    for fi76 := 0; fi76 < 13; fi76++ {
                                        var v77 bool = false
                                        _ = fi76
                                        _ = v77
                                    }
                                } else {
                                    for fi78 := 0; fi78 < 9; fi78++ {
                                        v79 := GenS8{F0: 42, F1: 209, F2: fi78}
                                        var v80 uint = (v2 >> ((255) & 63))
                                        _ = fi78
                                        _ = v79
                                        _ = v80
                                    }
                                    var v81 int8 = -48
                                    switch (fi32 | (fi32 - (fi32 << ((genFn0((^(v12)))) & 63)))) {
                                    case 261:
                                        v30 += 1
                                    case -2147483648:
                                        for fi82 := 0; fi82 < 5; fi82++ {
                                            if true {
                                                var v83 int64 = -9223372036854775808
                                                fmt.Println(v30, v83)
                                                fmt.Println(fi32, v12, v30)
                                                _ = v83
                                            }
                                            v84 := v1[(v12 * 4294967295)]
                                            _ = fi82
                                            _ = v84
                                        }
                                    case 42:
                                        delete(v5, "\n\t")
                                    }
                                    _ = v81
                                }
                                if (42 <= 100) {
                                    if false {
                                        v30 -= (v30 >> ((genFn0((v12 + v12))) & 63))
                                        for fi85 := 0; fi85 < 20; fi85++ {
                                            v30 = (v30 >> ((1) & 63))
                                            if true {
                                                v86 := GenS9{F0: -1}
                                                _ = v86
                                            }
                                            v30 = 1
                                            fmt.Println(fi32, v2, v11)
                                            _ = fi85
                                        }
                                    } else {
                                        var v87 uint16 = 823
                                        fmt.Println(v87, fi32)
                                        _ = len(v1)
                                        _ = v87
                                    }
                                }
                                if (false && false) {
                                    v2 += (v2 << ((v2) & 63))
                                    fmt.Println(v11, v11, v30)
                                    v11 = "a"
                                }
                                _ = fi32
                            }
                        } else {
                            if (1.5 >= 3.14159) {
                                if true {
                                    for fi88 := 0; fi88 < 12; fi88++ {
                                        _ = len(v1)
                                        _ = fi88
                                    }
                                    v89 := GenS7{F0: (^(v30)), F1: 1, F2: 255}
                                    _ = v89
                                } else {
                                    if false {
                                        if (-9223372036854775808 > -1) {
                                            var v90 int64 = 0
                                            v90 += (v90 / ((v90) | 1))
                                            _ = v90
                                        }
                                        var v91 string = "foo bar"
                                        _ = v91
                                    } else {
                                        switch v11 {
                                        case "\n\t":
                                            switch -9223372036854775808 {
                                            case 2147483647:
                                                var v92 uint = (v2 + genFn0((v12 + 65535)))
                                                _ = v92
                                            default:
                                                var v93 bool = (false && (true || (42 < 548)))
                                                _ = v93
                                            }
                                        case "hello":
                                            var v94 int32 = -1
                                            _ = v94
                                        case "foo bar":
                                            fmt.Println(v2, v30)
                                        }
                                        var v95 float64 = (1.5 + 1.5)
                                        _ = v95
                                    }
                                    if (127 != 128) {
                                        if false {
                                            v96 := GenS1{F0: v30, F1: -673, F2: 739}
                                            _ = v96
                                        }
                                        v97 := GenS6{F0: 1}
                                        _ = v97
                                    } else {
                                        var v98 uint16 = uint16(v12)
                                        var v99 string = v11
                                        _ = v98
                                        _ = v99
                                    }
                                    v100 := GenS10{F0: 0, F1: true}
                                    _ = v100
                                }
                            } else {
                                v101 := GenS11{F0: v12, F1: (true && (v30 < 4294967295))}
                                delete(v5, (v11 + v11))
                                _ = v101
                            }
                            for fi102 := 0; fi102 < 5; fi102++ {
                                switch v11 {
                                case "a":
                                    var v103 int16 = 32767
                                    _ = v103
                                case "":
                                    delete(v1, 282)
                                case "\n\t":
                                    _ = len(v3)
                                default:
                                    v104 := GenS2{F0: (474 == 0), F1: 0, F2: (v30 == v30)}
                                    _ = v104
                                }
                                switch "hello" {
                                case "\n\t":
                                    v105 := GenS12{F0: (v30 >> ((1) & 63)), F1: ((^(v30)) < (^(v30))), F2: 65535}
                                    _ = v105
                                default:
                                    if true {
                                        v12 -= (^(v12))
                                        switch v11 {
                                        case "hello":
                                            v12 -= v12
                                        case "":
                                            v12 = 0
                                        case "\n\t":
                                            for fi106 := 0; fi106 < 17; fi106++ {
                                                switch v11 {
                                                case "foo bar":
                                                    var v107 int16 = 398
                                                    _ = v107
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = "a"
                                                }
                                                var v108 uint64 = v30
                                                _ = fi106
                                                _ = v108
                                            }
                                        default:
                                            v11 = "foo bar"
                                        }
                                    } else {
                                        v5[((v11 + "\n\t") + ((v11 + "a") + "a"))] = v11
                                    }
                                }
                                _ = fi102
                            }
                        }
                    case -276:
                        switch v11 {
                        case "\n\t":
                            v109 := GenS13{F0: v2, F1: (!false), F2: 127}
                            _ = v109
                        case "hello":
                            if (false && false) {
                                v110 := GenS5{F0: 107, F1: 42}
                                if (v30 != v30) {
                                    fmt.Println(v30, v2)
                                    for fi111 := 0; fi111 < 5; fi111++ {
                                        _ = len(v1)
                                        var v112 int = fi111
                                        v11 = (((v11 + v11) + "\n\t") + "foo bar")
                                        _ = fi111
                                        _ = v112
                                    }
                                    v2 -= (v2 % (((v2 << ((v2) & 63))) | 1))
                                } else {
                                    for fi113 := 0; fi113 < 18; fi113++ {
                                        v2 += genFn0(v12)
                                        _ = fi113
                                    }
                                }
                                fmt.Println(v2, v30, v30)
                                _ = v110
                            } else {
                                if (-2.25 <= 0.1) {
                                    var v114 float32 = 1.5
                                    v2 = v2
                                    var v115 int32 = -2147483648
                                    _ = v114
                                    _ = v115
                                } else {
                                    var v116 int32 = 0
                                    fmt.Println(v30, v2)
                                    if true {
                                        v117 := GenS3{F0: 42, F1: v11, F2: 42}
                                        v11 = v11
                                        v116 -= 2147483647
                                        _ = v117
                                    } else {
                                        var v118 float64 = 0.0
                                        if false {
                                            for fi119 := 0; fi119 < 13; fi119++ {
                                                var v120 float64 = float64(1e10)
                                                var v121 uint16 = 447
                                                var v122 int16 = -1
                                                _ = fi119
                                                _ = v120
                                                _ = v121
                                                _ = v122
                                            }
                                        } else {
                                            var v123 uint16 = 898
                                            if false {
                                                v11 = v11
                                                for fi124 := 0; fi124 < 7; fi124++ {
                                                    v125 := GenS1{F0: (^(v30)), F1: -1, F2: 0}
                                                    fmt.Println(v118, v116, v2)
                                                    switch -9223372036854775808 {
                                                    case 1:
                                                        var v126 float32 = float32((v118 * v118))
                                                        _ = v126
                                                    case -2147483648:
                                                        v123 -= (v123 >> ((uint(v123)) & 63))
                                                    }
                                                    _ = fi124
                                                    _ = v125
                                                }
                                            }
                                            _ = v123
                                        }
                                        var v127 int16 = 1000
                                        _ = v118
                                        _ = v127
                                    }
                                    _ = v116
                                }
                                v12 += (^(v12))
                            }
                        default:
                            _ = len(v5)
                        }
                    case 789:
                        for fi128 := 0; fi128 < 18; fi128++ {
                            switch "" {
                            case "hello":
                                if true {
                                    var v129 int8 = 100
                                    v11 = v11
                                    _ = v129
                                }
                            case "a":
                                v130 := v1[v12]
                                _ = v130
                            default:
                                var v131 bool = true
                                _ = v131
                            }
                            var v132 uint16 = 1
                            _ = fi128
                            _ = v132
                        }
                    default:
                        v133 := GenS10{F0: 100, F1: (65535 != 0)}
                        _ = v133
                    }
                case "hello":
                    fmt.Println(v12, v12)
                }
                var v134 float64 = 1e10
                _ = v134
            } else {
                if true {
                    for fi135 := 0; fi135 < 10; fi135++ {
                        v2 = genFn0(v12)
                        _ = fi135
                    }
                } else {
                    for fi136 := 0; fi136 < 7; fi136++ {
                        v11 = ("\n\t" + (v11 + "foo bar"))
                        var v137 uint64 = (v30 | v30)
                        if len(v3) > 0 {
                            v3[0] = (v12 * v12)
                        }
                        _ = fi136
                        _ = v137
                    }
                    if (!true) {
                        if false {
                            if (!true) {
                                if (42 <= 475) {
                                    for fi138 := 0; fi138 < 9; fi138++ {
                                        v139 := GenS3{F0: 127, F1: "", F2: v30}
                                        _ = fi138
                                        _ = v139
                                    }
                                    if false {
                                        for fi140 := 0; fi140 < 7; fi140++ {
                                            for fi141 := 0; fi141 < 19; fi141++ {
                                                if (0 >= 0) {
                                                    var v142 int16 = 0
                                                    v12 += (v12 ^ 2147483648)
                                                    _ = v142
                                                }
                                                switch v11 {
                                                case "":
                                                    fmt.Println(fi141)
                                                case "foo bar":
                                                    switch (fi141 / (((^(fi140))) | 1)) {
                                                    case 241:
                                                        v143 := GenS6{F0: -1}
                                                        if true {
                                                            _ = len(v3)
                                                            if false {
                                                                v11 = v11
                                                            }
                                                        }
                                                        _ = v143
                                                    case 4294967295:
                                                        var v144 uint32 = v12
                                                        _ = v144
                                                    case 534:
                                                        switch fi141 {
                                                        case 9223372036854775807:
                                                            var v145 int32 = 976
                                                            _ = v145
                                                        case -1:
                                                            fmt.Println(v30, fi140)
                                                        }
                                                        var v146 uint16 = 95
                                                        _ = v146
                                                    default:
                                                        v6 = append(v6, (^(v12)))
                                                    }
                                                }
                                                if true {
                                                    v147 := GenS14{F0: 1}
                                                    _ = v147
                                                }
                                                v148 := v5[(v11 + (((("foo bar" + ("\n\t" + v11)) + v11) + ("hello" + ("\n\t" + "hello"))) + ((("a" + "foo bar") + (v11 + v11)) + v11)))]
                                                v11 = v148
                                                _ = fi141
                                                _ = v148
                                            }
                                            fmt.Println(v30, fi140)
                                            _ = fi140
                                        }
                                        switch "hello" {
                                        case "\n\t":
                                            v6 = append(v6, v12)
                                        case "hello":
                                            v30 = (v30 - (v30 << ((173) & 63)))
                                        case "":
                                            v2 = v2
                                        }
                                    }
                                    var v149 uint = (^(v2))
                                    _ = v149
                                }
                                for fi150 := 0; fi150 < 11; fi150++ {
                                    var v151 int32 = -891
                                    _ = fi150
                                    _ = v151
                                }
                            } else {
                                v152 := v5[v11]
                                var v153 int32 = 2147483647
                                v1[v12] = 0
                                _ = v152
                                _ = v153
                            }
                            v2 = (v2 ^ v2)
                            if true {
                                v12 = v12
                                switch ((("hello" + v11) + ((v11 + ("foo bar" + "")) + "")) + "a") {
                                case "a":
                                    v154 := GenS15{F0: 128}
                                    _ = v154
                                case "":
                                    v11 = ("hello" + v11)
                                default:
                                    fmt.Println(v30)
                                }
                            } else {
                                var v155 uint8 = 0
                                if (0 == 1) {
                                    var v156 int8 = 42
                                    v157 := GenS15{F0: v155}
                                    _ = v156
                                    _ = v157
                                }
                                _ = v155
                            }
                        }
                    } else {
                        var v158 int8 = 127
                        if true {
                            fmt.Println(v2, v2)
                            v2 = genFn0(v12)
                            fmt.Println(v11, v11)
                        } else {
                            v30 = (v30 / ((v30) | 1))
                            if true {
                                if true {
                                    v159 := GenS11{F0: (v12 - (v12 / ((0) | 1))), F1: false}
                                    fmt.Println(v30, v158, v2)
                                    _ = v159
                                } else {
                                    v11 = v11
                                }
                                v4 = append(v4, ("foo bar" + ((v11 + v11) + ("hello" + "a"))))
                                for fi160 := 0; fi160 < 17; fi160++ {
                                    for fi161 := 0; fi161 < 16; fi161++ {
                                        if ((0 <= 0) || true) {
                                            v162 := GenS8{F0: int16(v12), F1: 128, F2: fi160}
                                            var v163 uint = v2
                                            v164 := v1[v12]
                                            _ = v162
                                            _ = v163
                                            _ = v164
                                        } else {
                                            switch "" {
                                            case "hello":
                                                fmt.Println(v2, fi160, v11)
                                            default:
                                                if ((^(v2)) == v2) {
                                                    var v165 int16 = -32768
                                                    var v166 int16 = (^(v165))
                                                    _ = v165
                                                    _ = v166
                                                }
                                            }
                                        }
                                        var v167 float32 = (float32((1e10 - 1e10)) + -2.25)
                                        if (v30 <= 4294967295) {
                                            if false {
                                                var v168 uint32 = v12
                                                switch fi161 {
                                                case 127:
                                                    var v169 int32 = 42
                                                    _ = v169
                                                case 969:
                                                    if true {
                                                        v2 = v2
                                                        v1[(^(v168))] = (fi160 & 0)
                                                        if (((v158 / ((0) | 1)) != (v158 << (((v2 * 18446744073709551615)) & 63))) || (false && (!true))) {
                                                            fmt.Println(v168, fi160)
                                                        } else {
                                                            var v170 string = v11
                                                            v30 += (^(v30))
                                                            v168 += v168
                                                            _ = v170
                                                        }
                                                    }
                                                default:
                                                    v171 := v1[2147483647]
                                                    _ = v171
                                                }
                                                _ = v168
                                            } else {
                                                var v172 uint8 = uint8(v12)
                                                if len(v4) > 0 {
                                                    v4[0] = v11
                                                }
                                                v167 = v167
                                                _ = v172
                                            }
                                            var v173 int8 = int8(v12)
                                            v30 = v30
                                            _ = v173
                                        } else {
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            var v174 int = int(v30)
                                            v1[v12] = 4294967295
                                            _ = v174
                                        }
                                        _ = fi161
                                        _ = v167
                                    }
                                    _ = fi160
                                }
                            } else {
                                var v175 uint16 = 871
                                var v176 uint32 = v12
                                _ = v175
                                _ = v176
                            }
                        }
                        v177 := v5[((("a" + ((v11 + v11) + (v11 + "\n\t"))) + v11) + v11)]
                        _ = v158
                        _ = v177
                    }
                }
                var v178 uint64 = (v30 | v30)
                _ = v178
            }
            _ = v30
        }
        switch v11 {
        case "\n\t":
            var v179 float64 = 0.0
            _ = v179
        default:
            v11 = (v11 + v11)
        }
        _ = v29
    }
    fmt.Println(v2)
    var v180 int64 = -1
    v181 := GenS16{F0: 0, F1: v12, F2: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v180:", v180)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v180
    _ = v181
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
