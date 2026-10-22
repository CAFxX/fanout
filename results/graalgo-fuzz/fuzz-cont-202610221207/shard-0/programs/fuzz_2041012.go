package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int32; F1 bool; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 uint8; F1 int64 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS10 struct { F0 int16; F1 uint64; F2 int16 }
type GenS11 struct { F0 int8; F1 string; F2 int8 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int; F1 int; F2 uint32 }
type GenS15 struct { F0 uint32; F1 int16 }
type GenS16 struct { F0 uint16; F1 int8 }
type GenS17 struct { F0 uint32; F1 uint64 }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 int32 }
func genFn0() string {
    var v4 int = 127
    if (!(false && false)) {
        fmt.Println(v4)
        if false {
            var v5 int8 = 42
            v6 := []int{(^(v4))}
            v6 = append(v6, v4)
            _ = v5
            _ = v6
        } else {
            if (float64(0.0) <= (-2.25 + 1e10)) {
                for fi7 := 0; fi7 < 16; fi7++ {
                    v4 = -143
                    var v8 int16 = 42
                    for fi9 := 0; fi9 < 8; fi9++ {
                        v4 -= v4
                        v4 -= fi7
                        fmt.Println(v8)
                        v4 += 32767
                        if (-1 != -133) {
                            if true {
                                v8 = (v8 << ((42) & 63))
                            } else {
                                if (1 > 2) {
                                    v8 += (-(v8))
                                    var v10 bool = (!(127 >= 0))
                                    _ = v10
                                }
                            }
                            var v11 uint8 = 1
                            fmt.Println(v11)
                            var v12 float32 = (-2.25 / 1.0)
                            v13 := GenS1{F0: 4294967295, F1: v8}
                            if false {
                                switch -9223372036854775808 {
                                case -853:
                                    v14 := []int{9223372036854775807, fi9, (v4 % ((v4) | 1)), -1}
                                    _ = v14
                                }
                                if (0 != 76) {
                                    v15 := GenS1{F0: 2, F1: 666}
                                    v8 -= 42
                                    v16 := []int{fi9, fi7, 0, (fi7 / ((462) | 1))}
                                    _ = v15
                                    _ = v16
                                }
                                if true {
                                    v8 -= (v8 / (((v8 * v8)) | 1))
                                }
                            }
                            var v17 int64 = -9223372036854775808
                            _ = v11
                            _ = v12
                            _ = v13
                            _ = v17
                        }
                        v18 := GenS1{F0: 18446744073709551615, F1: v8}
                        v19 := []int{1, 32767}
                        fmt.Println(fi7, v4)
                        v8 += 116
                        _ = fi9
                        _ = v18
                        _ = v19
                    }
                    _ = fi7
                    _ = v8
                }
            } else {
                v4 = v4
                var v20 int64 = 1
                var v21 uint16 = 147
                _ = v20
                _ = v21
            }
            var v22 int32 = 748
            fmt.Println(v4, v4)
            _ = v22
        }
    }
    if (!(2 >= 127)) {
        v23 := []int{-1, (^(v4)), (v4 >> ((4294967295) & 63)), v4}
        var v24 uint16 = 1
        v23 = append(v23, v4)
        _ = v23
        _ = v24
    }
    var v25 uint64 = 18446744073709551615
    _ = v25
    _ = v4
    return ""
}
func genFn1(p0 uint64, p1 string) uint16 {
    v26 := GenS1{F0: 908, F1: -1}
    var v27 int16 = -501
    switch p1 {
    case "hello":
        v28 := []int{114, 42, 187}
        _ = v28
    case "\n\t":
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 uint64 = (p0 | 2)
            fmt.Println(p0, v30, v27)
            var v31 int64 = 0
            _ = fi29
            _ = v30
            _ = v31
        }
    default:
        if (!(p1 == (p1 + (p1 + "\n\t")))) {
            var v32 string = p1
            _ = v32
        }
    }
    p0 = 775
    p1 = "\n\t"
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return 273
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (0 > 65535)
    var v3 uint8 = 0
    if (0 > 42) {
        var v33 float64 = 3.14159
        for fi34 := 0; fi34 < 7; fi34++ {
            v2 = false
            v35 := GenS2{F0: int32(fi34), F1: (0 != 819), F2: -128}
            switch "a" {
            case "":
                var v36 uint32 = 65535
                _ = v36
            case "hello":
                v37 := GenS3{F0: 12}
                _ = v37
            case "\n\t":
                _ = len(v1)
            default:
                v3 = v3
            }
            _ = fi34
            _ = v35
        }
        _ = v33
    }
    if (v2 && true) {
        v3 = (^(v3))
        var v38 string = "a"
        var v39 int32 = 42
        _ = v38
        _ = v39
    } else {
        var v40 int64 = -311
        var v41 bool = (!v2)
        var v42 int8 = -109
        _ = v40
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 6; fi43++ {
        v44 := v1[2147483647]
        _ = fi43
        _ = v44
    }
    _ = len(v1)
    v45 := []int{108, -907}
    var v46 int64 = -909
    for fi47 := 0; fi47 < 7; fi47++ {
        if (fi47 > fi47) {
            v48 := GenS3{F0: 172}
            _ = v48
        }
        v46 = (^(v46))
        _ = fi47
    }
    _ = len(v45)
    if v2 {
        if (1 >= 533) {
            v2 = v2
            fmt.Println(v3)
        }
        v3 += v3
    } else {
        if v2 {
            v2 = (!v2)
            v3 = v3
            var v49 float64 = (float64((0.0 * -2.25)) - (-2.25 - 0.0))
            _ = v49
        }
    }
    v50 := GenS3{F0: 42}
    if (0.1 >= (1.5 - -2.25)) {
        v46 = -687
        v46 -= (v46 >> ((127) & 63))
        if (65535 >= 65535) {
            var v51 int8 = 15
            v2 = true
            _ = v51
        } else {
            var v52 bool = (v3 == 255)
            _ = v52
        }
    } else {
        v45 = append(v45, 42)
        if ((false || false) || (v2 || (v2 || (-9223372036854775808 < 32767)))) {
            if true {
                v46 += v46
                v1[176] = 9
                switch -1 {
                case 128:
                    v2 = (-717 != 32767)
                default:
                    v46 = (v46 ^ v46)
                }
            }
            var v53 uint32 = 42
            if len(v45) > 0 {
                v45[0] = 32767
            }
            _ = v53
        } else {
            for fi54 := 0; fi54 < 7; fi54++ {
                v55 := GenS4{F0: (("a" + ("hello" + "\n\t")) + "foo bar"), F1: (("a" + ("\n\t" + (("\n\t" + "a") + ""))) + (("a" + "a") + ("hello" + ("a" + ""))))}
                if true {
                    v56 := GenS5{F0: 128, F1: (^(v46))}
                    for fi57 := 0; fi57 < 5; fi57++ {
                        var v58 uint8 = 2
                        _ = fi57
                        _ = v58
                    }
                    _ = v56
                } else {
                    for fi59 := 0; fi59 < 8; fi59++ {
                        if v2 {
                            var v60 int = 306
                            _ = v60
                        } else {
                            v61 := GenS4{F0: ("\n\t" + "hello"), F1: ("a" + "")}
                            _ = len(v45)
                            _ = v61
                        }
                        if len(v45) >= 2 {
                            v45 = v45[1:len(v45)]
                        }
                        v3 = v3
                        _ = fi59
                    }
                    if true {
                        var v62 int16 = 831
                        for fi63 := 0; fi63 < 20; fi63++ {
                            var v64 int16 = (v62 | v62)
                            var v65 uint32 = 4294967295
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                        if false {
                            var v66 int = (fi54 << ((65535) & 63))
                            switch "hello" {
                            case "\n\t":
                                var v67 uint32 = 125
                                _ = v67
                            }
                            _ = v66
                        } else {
                            if true {
                                fmt.Println(fi54, v46, v62)
                                if v2 {
                                    if true {
                                        var v68 int64 = (v46 + v46)
                                        _ = v68
                                    }
                                    v2 = v2
                                    v45 = append(v45, fi54)
                                } else {
                                    for fi69 := 0; fi69 < 13; fi69++ {
                                        v45 = append(v45, 2147483647)
                                        var v70 uint64 = 706
                                        _ = fi69
                                        _ = v70
                                    }
                                }
                            } else {
                                fmt.Println(fi54, v46, v46)
                                if v2 {
                                    for fi71 := 0; fi71 < 11; fi71++ {
                                        v3 = (v3 | (v3 + v3))
                                        if (!false) {
                                            v1[65535] = 624
                                            if v2 {
                                                var v72 uint16 = 716
                                                fmt.Println(fi54, v72)
                                                if (v2 || v2) {
                                                    v73 := GenS6{F0: 2, F1: 567}
                                                    _ = v73
                                                } else {
                                                    v72 = v72
                                                    var v74 uint32 = 396
                                                    var v75 int64 = -524
                                                    _ = v74
                                                    _ = v75
                                                }
                                                switch "foo bar" {
                                                case "":
                                                    fmt.Println(v62, fi54)
                                                case "\n\t":
                                                    var v76 uint16 = (v72 - genFn1(1, ""))
                                                    _ = v76
                                                case "a":
                                                    if true {
                                                        if (genFn1(2, "a") >= (^(v72))) {
                                                            var v77 float32 = 3.14159
                                                            v78 := GenS7{F0: 2}
                                                            var v79 uint = 0
                                                            switch fi71 {
                                                            case 536:
                                                                v46 = 42
                                                            }
                                                            _ = v77
                                                            _ = v78
                                                            _ = v79
                                                        } else {
                                                            switch "\n\t" {
                                                            case "":
                                                                var v80 bool = true
                                                                _ = v80
                                                            case "foo bar":
                                                                v3 -= (v3 | v3)
                                                            case "a":
                                                                v72 += 65535
                                                            }
                                                            v81 := GenS8{F0: (fi71 * (fi71 >> ((0) & 63)))}
                                                            _ = v81
                                                        }
                                                    } else {
                                                        var v82 uint64 = 656
                                                        _ = v82
                                                    }
                                                default:
                                                    switch fi71 {
                                                    case 32768:
                                                        if v2 {
                                                            v83 := GenS3{F0: v72}
                                                            var v84 float64 = ((1.5 / (1e10 / 3.14159)) - 3.14159)
                                                            _ = v83
                                                            _ = v84
                                                        } else {
                                                            var v85 float32 = ((3.14159 - 1e10) / 0.1)
                                                            _ = v85
                                                        }
                                                        var v86 int8 = 67
                                                        v45 = append(v45, fi71)
                                                        _ = v86
                                                    case 1:
                                                        if (!(v2 || v2)) {
                                                            _ = len(v1)
                                                            var v87 float32 = 3.14159
                                                            fmt.Println(v62, v2, v72)
                                                            _ = v87
                                                        } else {
                                                            var v88 uint8 = v3
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v89 := GenS9{F0: 684, F1: 128, F2: (v72 - v72)}
                                                                _ = v89
                                                            }
                                                            var v90 int8 = -46
                                                            _ = v88
                                                            _ = v90
                                                        }
                                                    case -2147483648:
                                                        v91 := GenS10{F0: (v62 & 401), F1: 133, F2: (v62 & 32767)}
                                                        _ = v91
                                                    default:
                                                        v62 += (-(v62))
                                                    }
                                                }
                                                v92 := GenS10{F0: (v62 + v62), F1: 4294967295, F2: v62}
                                                v93 := GenS9{F0: 695, F1: v3, F2: (v72 >> ((413) & 63))}
                                                var v94 uint64 = 437
                                                v2 = (genFn1(v94, "hello") < (v72 * (v72 >> ((42) & 63))))
                                                _ = v72
                                                _ = v92
                                                _ = v93
                                                _ = v94
                                            } else {
                                                v95 := GenS4{F0: ("a" + "foo bar"), F1: (((("hello" + "a") + "hello") + "") + "a")}
                                                var v96 uint8 = v3
                                                var v97 int16 = (^(v62))
                                                fmt.Println(v96)
                                                _ = v95
                                                _ = v96
                                                _ = v97
                                            }
                                            if (!v2) {
                                                fmt.Println(fi54)
                                            } else {
                                                v98 := GenS1{F0: 18446744073709551615, F1: v62}
                                                var v99 bool = (!(v2 || (255 != 754)))
                                                _ = v98
                                                _ = v99
                                            }
                                        } else {
                                            v3 = 0
                                        }
                                        var v100 float64 = (1e10 + (0.0 + 0.1))
                                        v45 = append(v45, (-(fi71)))
                                        if true {
                                            v2 = v2
                                            v3 = (v3 % ((v3) | 1))
                                        } else {
                                            fmt.Println(v3, v3)
                                            var v101 int64 = (v46 ^ v46)
                                            v102 := GenS11{F0: 0, F1: "foo bar", F2: 42}
                                            _ = v101
                                            _ = v102
                                        }
                                        fmt.Println(fi54, v62, v3)
                                        switch ("\n\t" + "a") {
                                        case "":
                                            switch "a" {
                                            case "\n\t":
                                                if v2 {
                                                    if ("foo bar" == "a") {
                                                        _ = len(v45)
                                                        fmt.Println(v100, fi54)
                                                        var v103 int16 = (v62 | v62)
                                                        _ = v103
                                                    } else {
                                                        if v2 {
                                                            switch "hello" {
                                                            case "foo bar":
                                                                var v104 float32 = 3.14159
                                                                _ = v104
                                                            case "":
                                                                v2 = true
                                                            case "a":
                                                                var v105 uint64 = 42
                                                                _ = v105
                                                            default:
                                                                var v106 int = (fi71 * 1)
                                                                _ = v106
                                                            }
                                                            if len(v45) >= 2 {
                                                                v45 = v45[1:len(v45)]
                                                            }
                                                        } else {
                                                            if true {
                                                                v1[93] = 490
                                                            } else {
                                                                if (620 <= 373) {
                                                                    var v107 uint = 123
                                                                    var v108 uint32 = 65535
                                                                    _ = len(v45)
                                                                    v109 := GenS12{F0: int8(v107)}
                                                                    v110 := GenS12{F0: -36}
                                                                    _ = v107
                                                                    _ = v108
                                                                    _ = v109
                                                                    _ = v110
                                                                } else {
                                                                    v100 -= -2.25
                                                                    var v111 int = -391
                                                                    _ = v111
                                                                }
                                                                switch (-(fi71)) {
                                                                case 18:
                                                                    switch "a" {
                                                                    case "hello":
                                                                        if v2 {
                                                                            fmt.Println(v3, fi54)
                                                                        } else {
                                                                            fmt.Println(v3, fi54)
                                                                            v2 = (-32768 != v62)
                                                                        }
                                                                    case "":
                                                                        _ = len(v45)
                                                                    default:
                                                                        v112 := v1[1]
                                                                        _ = v112
                                                                    }
                                                                case 9223372036854775807:
                                                                    var v113 float64 = v100
                                                                    _ = v113
                                                                default:
                                                                    switch 32768 {
                                                                    case 32768:
                                                                        v114 := v1[1]
                                                                        _ = v114
                                                                    case 0:
                                                                        if ((!(!((873 != 42) || v2))) && (true || (false && ((v2 && (-1 < (v46 * -1))) || (!(!v2)))))) {
                                                                            var v115 int = fi54
                                                                            var v116 string = "hello"
                                                                            _ = v115
                                                                            _ = v116
                                                                        }
                                                                    default:
                                                                        if v2 {
                                                                            _ = len(v1)
                                                                            v117 := GenS13{F0: "foo bar"}
                                                                            _ = v117
                                                                        } else {
                                                                            v118 := v1[285]
                                                                            var v119 string = "foo bar"
                                                                            v120 := GenS14{F0: 42, F1: 573, F2: v118}
                                                                            if ((-(fi71)) != (fi71 * fi71)) {
                                                                                if len(v45) >= 2 {
                                                                                    v45 = v45[1:len(v45)]
                                                                                }
                                                                                var v121 bool = (v46 >= (-(v46)))
                                                                                var v122 float64 = v100
                                                                                var v123 uint16 = 65535
                                                                                _ = v121
                                                                                _ = v122
                                                                                _ = v123
                                                                            }
                                                                            if (v62 >= v62) {
                                                                                var v124 int8 = int8(v62)
                                                                                v125 := GenS13{F0: ""}
                                                                                var v126 uint16 = 0
                                                                                v127 := v1[uint32(v126)]
                                                                                v119 = "\n\t"
                                                                                v3 -= (^(v3))
                                                                                var v128 bool = true
                                                                                _ = v124
                                                                                _ = v125
                                                                                _ = v126
                                                                                _ = v127
                                                                                _ = v128
                                                                            } else {
                                                                                var v129 int64 = v46
                                                                                switch ((v119 + ((v119 + v119) + "a")) + "hello") {
                                                                                case "a":
                                                                                    _ = len(v45)
                                                                                case "\n\t":
                                                                                    if v2 {
                                                                                        var v130 uint64 = 118
                                                                                        v45 = append(v45, fi54)
                                                                                        _ = v130
                                                                                    }
                                                                                    if v2 {
                                                                                        v131 := GenS5{F0: (^(v3)), F1: (^(v129))}
                                                                                        switch v119 {
                                                                                        case "a":
                                                                                            var v132 uint64 = 580
                                                                                            _ = v132
                                                                                        case "\n\t":
                                                                                            switch v119 {
                                                                                            case "a":
                                                                                                var v133 uint8 = v3
                                                                                                _ = v133
                                                                                            }
                                                                                            v3 = (v3 | 2)
                                                                                            v62 += v62
                                                                                            var v134 uint16 = 0
                                                                                            _ = v134
                                                                                        }
                                                                                        var v135 uint16 = 134
                                                                                        v129 -= (v46 | (v129 % (((v46 & (v46 | (-(v46))))) | 1)))
                                                                                        if v2 {
                                                                                            if false {
                                                                                                v136 := GenS15{F0: uint32(v3), F1: (^(v62))}
                                                                                                _ = v136
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v118, v3)
                                                                                        if len(v45) > 0 {
                                                                                            v45[0] = (fi71 % ((186) | 1))
                                                                                        }
                                                                                        v119 = ""
                                                                                        _ = v131
                                                                                        _ = v135
                                                                                    } else {
                                                                                        var v137 uint8 = (^(v3))
                                                                                        v46 -= 195
                                                                                        _ = v137
                                                                                    }
                                                                                    if (42 <= genFn1(860, v119)) {
                                                                                        if ((!(-72 < 66)) || (!(4294967295 == 2))) {
                                                                                            switch "foo bar" {
                                                                                            case "\n\t":
                                                                                                v62 = 484
                                                                                            }
                                                                                            if (42 <= v118) {
                                                                                                var v138 string = ((((v119 + v119) + ("" + v119)) + v119) + (v119 + v119))
                                                                                                _ = v138
                                                                                            }
                                                                                        } else {
                                                                                            var v139 uint16 = 32767
                                                                                            _ = v139
                                                                                        }
                                                                                        var v140 uint16 = 1
                                                                                        var v141 uint8 = v3
                                                                                        v142 := GenS12{F0: 127}
                                                                                        _ = len(v45)
                                                                                        _ = v140
                                                                                        _ = v141
                                                                                        _ = v142
                                                                                    } else {
                                                                                        v2 = (964 != -1)
                                                                                        var v143 int16 = v62
                                                                                        _ = v143
                                                                                    }
                                                                                case "foo bar":
                                                                                    v144 := GenS16{F0: 42, F1: 18}
                                                                                    _ = v144
                                                                                default:
                                                                                    fmt.Println(fi54, v100, v62)
                                                                                }
                                                                                var v145 bool = (v2 || true)
                                                                                var v146 bool = ((fi71 - 128) > fi71)
                                                                                v145 = (v145 && ((-126 <= 100) || true))
                                                                                v146 = true
                                                                                _ = v129
                                                                                _ = v145
                                                                                _ = v146
                                                                            }
                                                                            _ = v118
                                                                            _ = v119
                                                                            _ = v120
                                                                        }
                                                                    }
                                                                    var v147 string = "foo bar"
                                                                    _ = v147
                                                                }
                                                            }
                                                            if ((fi71 % ((fi54) | 1)) != fi54) {
                                                                var v148 uint16 = 255
                                                                if false {
                                                                    v2 = false
                                                                } else {
                                                                    if (1 <= 2147483647) {
                                                                        fmt.Println(v148, v2)
                                                                    } else {
                                                                        fmt.Println(v100, v148)
                                                                        var v149 uint32 = 65535
                                                                        _ = v149
                                                                    }
                                                                    switch 9223372036854775807 {
                                                                    case 2147483647:
                                                                        var v150 uint32 = 87
                                                                        _ = v150
                                                                    case 0:
                                                                        v151 := GenS17{F0: uint32(v3), F1: 42}
                                                                        _ = v151
                                                                    case 42:
                                                                        v100 -= (v100 - v100)
                                                                    default:
                                                                        if v2 {
                                                                            var v152 uint = 252
                                                                            _ = v152
                                                                        }
                                                                    }
                                                                }
                                                                _ = v148
                                                            } else {
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    v153 := GenS7{F0: 18446744073709551615}
                                                                    _ = v153
                                                                default:
                                                                    fmt.Println(v62, fi71, fi54)
                                                                }
                                                                v46 += v46
                                                            }
                                                            delete(v1, 1)
                                                        }
                                                        v1[1] = 4294967295
                                                        if v2 {
                                                            v100 += (1.5 / (v100 / float64(float32(v100))))
                                                        }
                                                    }
                                                    var v154 int16 = (v62 ^ (v62 - v62))
                                                    var v155 int64 = (v46 >> ((2) & 63))
                                                    _ = v154
                                                    _ = v155
                                                } else {
                                                    var v156 uint64 = 18446744073709551615
                                                    var v157 uint32 = 4294967295
                                                    _ = v156
                                                    _ = v157
                                                }
                                            default:
                                                _ = len(v1)
                                            }
                                        }
                                        var v158 uint32 = 42
                                        var v159 uint16 = 32767
                                        _ = fi71
                                        _ = v100
                                        _ = v158
                                        _ = v159
                                    }
                                    v160 := GenS18{F0: genFn1(0, "foo bar")}
                                    for fi161 := 0; fi161 < 4; fi161++ {
                                        v62 = v62
                                        switch "\n\t" {
                                        case "\n\t":
                                            v162 := GenS4{F0: "a", F1: "\n\t"}
                                            _ = v162
                                        case "foo bar":
                                            v163 := GenS18{F0: 65535}
                                            _ = v163
                                        default:
                                            switch (fi54 * 9223372036854775807) {
                                            case 42:
                                                switch fi161 {
                                                case 127:
                                                    if ((v3 | v3) > v3) {
                                                        v2 = v2
                                                        if false {
                                                            delete(v1, uint32(v3))
                                                        }
                                                    }
                                                    var v164 float64 = (((-2.25 * float64(1.5)) + 1e10) * 0.1)
                                                    _ = v164
                                                case 9223372036854775807:
                                                    var v165 uint64 = 512
                                                    _ = v165
                                                case 128:
                                                    v62 -= v62
                                                }
                                            case 32767:
                                                v166 := GenS19{F0: 2147483647}
                                                _ = v166
                                            case 0:
                                                v62 += (v62 - int16(v46))
                                            }
                                        }
                                        if v2 {
                                            var v167 int = fi161
                                            v168 := GenS8{F0: 82}
                                            _ = v167
                                            _ = v168
                                        } else {
                                            v1[4294967295] = 170
                                            fmt.Println(v3, fi54, v46)
                                            var v169 bool = v2
                                            _ = v169
                                        }
                                        _ = fi161
                                    }
                                    _ = v160
                                }
                            }
                            fmt.Println(v3, fi54, v2)
                            v170 := GenS4{F0: ("hello" + ""), F1: ("hello" + "")}
                            _ = v170
                        }
                        _ = v62
                    }
                }
                for fi171 := 0; fi171 < 3; fi171++ {
                    v172 := GenS19{F0: 2147483647}
                    v173 := GenS9{F0: genFn1(2, "foo bar"), F1: v3, F2: 175}
                    fmt.Println(v3)
                    _ = fi171
                    _ = v172
                    _ = v173
                }
                _ = fi54
                _ = v55
            }
            var v174 float32 = (float32(1.5) + (((float32(0.1) + float32(3.14159)) / -2.25) * 0.0))
            _ = v174
        }
    }
    v3 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v50
}
