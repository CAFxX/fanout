package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 bool; F2 int32 }
type GenS3 struct { F0 int; F1 uint16; F2 int64 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS8 struct { F0 int64; F1 int8 }
type GenS9 struct { F0 int16; F1 int16 }
type GenS10 struct { F0 int64; F1 int8 }
type GenS11 struct { F0 int16; F1 int32; F2 int32 }
type GenS12 struct { F0 uint32; F1 int64; F2 int32 }

func main() {
    var v1 uint64 = 1
    var v2 uint64 = (v1 + (^(v1)))
    var v3 string = ("hello" + "a")
    var v4 float64 = -2.25
    var v5 float64 = float64((float32((3.14159 * v4)) * float32(float64(3.14159))))
    v6 := make(map[int]uint32)
    v5 = v4
    var v7 uint8 = 0
    if false {
        if ((923 > 2) || (-1 < 2147483647)) {
            fmt.Println(v7, v7, v4)
            var v8 int16 = 756
            v9 := []int{-2147483648, 128, -303, -1}
            _ = v8
            _ = v9
        } else {
            if (42 > 512) {
                if false {
                    if (0 == 2) {
                        var v10 uint32 = 42
                        if (true && (v4 >= v4)) {
                            var v11 int16 = 42
                            var v12 int16 = v11
                            v13 := GenS1{F0: 623}
                            _ = v11
                            _ = v12
                            _ = v13
                        } else {
                            fmt.Println(v4, v7, v7)
                            var v14 float64 = ((1.5 + 0.1) * v4)
                            _ = v14
                        }
                        var v15 uint32 = v10
                        _ = v10
                        _ = v15
                    }
                    v7 = v7
                    v16 := GenS2{F0: v7, F1: ((v7 % ((42) | 1)) > (v7 << ((42) & 63))), F2: 2147483647}
                    _ = v16
                } else {
                    v17 := GenS2{F0: (^(v7)), F1: ((!true) && true), F2: 556}
                    _ = v17
                }
            } else {
                v18 := []int{9223372036854775807, -2147483648, -2147483648, -9223372036854775808}
                v2 = (^(v2))
                if false {
                    for fi19 := 0; fi19 < 1; fi19++ {
                        switch 32767 {
                        case 42:
                            v3 = "hello"
                        case -434:
                            v20 := GenS1{F0: 65535}
                            _ = v20
                        case -9223372036854775808:
                            if (65535 <= 22) {
                                var v21 bool = true
                                v22 := v6[4294967295]
                                var v23 bool = (!v21)
                                _ = v21
                                _ = v22
                                _ = v23
                            } else {
                                if (!false) {
                                    v24 := GenS3{F0: fi19, F1: 0, F2: -199}
                                    if false {
                                        _ = len(v6)
                                    } else {
                                        v4 -= (v4 + (((v4 - v4) / 1e10) - (v5 / 3.14159)))
                                        switch (fi19 - (^(fi19))) {
                                        case -2147483648:
                                            var v25 bool = (0 <= 515)
                                            _ = v25
                                        default:
                                            switch v3 {
                                            case "":
                                                switch fi19 {
                                                case -611:
                                                    if false {
                                                        v26 := GenS3{F0: fi19, F1: 42, F2: -1}
                                                        fmt.Println(v4)
                                                        _ = v26
                                                    }
                                                case -568:
                                                    if false {
                                                        if (100 != 0) {
                                                            var v27 int64 = 224
                                                            fmt.Println(v7)
                                                            _ = v27
                                                        } else {
                                                            v2 = 295
                                                            v28 := v6[fi19]
                                                            _ = v28
                                                        }
                                                        fmt.Println(v7, v4, v2)
                                                        v3 = v3
                                                    }
                                                default:
                                                    var v29 float32 = 0.1
                                                    _ = v29
                                                }
                                            case "\n\t":
                                                fmt.Println(v3)
                                            case "a":
                                                var v30 uint32 = 196
                                                _ = v30
                                            default:
                                                v18 = append(v18, 9223372036854775807)
                                            }
                                        }
                                        for fi31 := 0; fi31 < 1; fi31++ {
                                            fmt.Println(v4, v7)
                                            v18 = append(v18, 42)
                                            if false {
                                                var v32 float64 = (-2.25 - 3.14159)
                                                v33 := GenS2{F0: (v7 % ((0) | 1)), F1: true, F2: 0}
                                                _ = v32
                                                _ = v33
                                            } else {
                                                var v34 float32 = 3.14159
                                                switch v3 {
                                                case "\n\t":
                                                    var v35 uint8 = v7
                                                    _ = v35
                                                case "":
                                                    if true {
                                                        v34 = 1.5
                                                        if (true && true) {
                                                            v2 = (v1 / ((0) | 1))
                                                        } else {
                                                            var v36 int8 = -60
                                                            var v37 bool = ((^(v7)) != (^(v7)))
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        var v38 uint8 = v7
                                                        _ = v38
                                                    }
                                                case "a":
                                                    v7 = (^(v7))
                                                default:
                                                    if (-922 <= 9223372036854775807) {
                                                        v1 += v2
                                                        if (-200 != 52) {
                                                            switch "foo bar" {
                                                            case "":
                                                                if false {
                                                                    v2 += v1
                                                                } else {
                                                                    var v39 bool = (255 <= uint16(v2))
                                                                    _ = v39
                                                                }
                                                            }
                                                            var v40 uint8 = (v7 * (^(v7)))
                                                            switch fi31 {
                                                            case 2147483647:
                                                                var v41 int8 = 100
                                                                _ = v41
                                                            case 9223372036854775807:
                                                                v42 := GenS3{F0: (fi19 ^ (fi31 ^ 32768)), F1: 32767, F2: 1}
                                                                _ = v42
                                                            default:
                                                                v43 := GenS4{F0: ((("a" + v3) + v3) + "foo bar"), F1: 4294967295}
                                                                fmt.Println(v4)
                                                                var v44 int = (fi19 ^ (-(fi31)))
                                                                _ = v43
                                                                _ = v44
                                                            }
                                                            _ = v40
                                                        }
                                                        if (false || (1 < 100)) {
                                                            fmt.Println(v3)
                                                            if false {
                                                                var v45 float32 = ((v34 - ((((v34 + (((v34 * 1.5) - 0.0) + v34)) + (v34 - v34)) / ((v34 + float32(1e10)) + ((v34 + v34) * float32((-2.25 * 1.5))))) * float32(v4))) / ((3.14159 * 1e10) + v34))
                                                                if (!((v7 != v7) || (2 >= 600))) {
                                                                    switch v3 {
                                                                    case "":
                                                                        var v46 int64 = 0
                                                                        _ = v46
                                                                    default:
                                                                        fmt.Println(v34, v7, fi19)
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        v4 -= (v5 / 1.0)
                                                                        v47 := v6[(fi19 & fi19)]
                                                                        if ((((0.1 - ((((v4 * float64(v34)) + 1e10) * (float64(0.1) - (v4 * (v4 * (0.1 * 3.14159))))) / v4)) < v4) || true) || false) {
                                                                            if false {
                                                                                v48 := GenS2{F0: v7, F1: (v2 >= 4294967295), F2: 2147483647}
                                                                                var v49 int16 = int16(v1)
                                                                                v4 += (((1e10 + ((0.0 * v5) * v5)) / v5) / (v4 / (v4 - v5)))
                                                                                switch 518 {
                                                                                case 1:
                                                                                    if true {
                                                                                        fmt.Println(fi31, fi31)
                                                                                    } else {
                                                                                        v50 := GenS4{F0: ("foo bar" + v3), F1: 42}
                                                                                        _ = v50
                                                                                    }
                                                                                default:
                                                                                    v51 := GenS5{F0: -853, F1: (v7 << ((18446744073709551615) & 63)), F2: (v47 | v47)}
                                                                                    _ = v51
                                                                                }
                                                                                _ = v48
                                                                                _ = v49
                                                                            } else {
                                                                                var v52 int32 = 441
                                                                                v45 = ((v45 * v34) + -2.25)
                                                                                if true {
                                                                                    fmt.Println(v52, v2)
                                                                                    var v53 bool = false
                                                                                    _ = v53
                                                                                }
                                                                                if true {
                                                                                    fmt.Println(v5, v47, v52)
                                                                                    var v54 string = v3
                                                                                    v1 = 2
                                                                                    _ = v54
                                                                                } else {
                                                                                    v55 := GenS2{F0: v7, F1: (!true), F2: v52}
                                                                                    _ = v55
                                                                                }
                                                                                v56 := GenS1{F0: 65535}
                                                                                var v57 uint = 18446744073709551615
                                                                                _ = v52
                                                                                _ = v56
                                                                                _ = v57
                                                                            }
                                                                            fmt.Println(fi19)
                                                                            switch (fi19 + fi31) {
                                                                            case -243:
                                                                                if (!(int16(v7) <= -214)) {
                                                                                    if len(v18) >= 2 {
                                                                                        v18 = v18[1:len(v18)]
                                                                                    }
                                                                                    v1 -= (^(v2))
                                                                                }
                                                                            case -263:
                                                                                v2 += (v2 % ((v1) | 1))
                                                                            case -9223372036854775808:
                                                                                if ((v47 << ((2) & 63)) == v47) {
                                                                                    fmt.Println(v5, v2)
                                                                                }
                                                                            default:
                                                                                _ = len(v6)
                                                                            }
                                                                        } else {
                                                                            var v58 string = (("" + ("\n\t" + (v3 + (((v3 + ("" + v3)) + v3) + v3)))) + v3)
                                                                            _ = len(v6)
                                                                            var v59 uint = 65535
                                                                            _ = v58
                                                                            _ = v59
                                                                        }
                                                                        v60 := GenS5{F0: 0, F1: (v7 % (((v7 / ((v7) | 1))) | 1)), F2: v47}
                                                                        _ = v47
                                                                        _ = v60
                                                                    } else {
                                                                        v61 := GenS1{F0: 255}
                                                                        v62 := GenS6{F0: uint32(v2), F1: -690}
                                                                        _ = v61
                                                                        _ = v62
                                                                    }
                                                                    if true {
                                                                        v5 -= v5
                                                                        var v63 uint16 = 255
                                                                        _ = v63
                                                                    }
                                                                }
                                                                v4 = v5
                                                                var v64 int64 = 383
                                                                _ = v45
                                                                _ = v64
                                                            } else {
                                                                v34 -= float32(v5)
                                                            }
                                                            var v65 bool = true
                                                            _ = v65
                                                        } else {
                                                            var v66 uint32 = 0
                                                            var v67 uint64 = v2
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                    } else {
                                                        var v68 int8 = 100
                                                        v69 := GenS1{F0: 380}
                                                        v3 = v3
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                }
                                                if len(v18) >= 2 {
                                                    v18 = v18[1:len(v18)]
                                                }
                                                _ = v34
                                            }
                                            _ = fi31
                                        }
                                    }
                                    v7 += 1
                                    _ = v24
                                } else {
                                    if ((-25 >= 529) && ((v2 | (^(v1))) < v2)) {
                                        if (0 == -940) {
                                            v70 := GenS4{F0: (((((v3 + ((v3 + "hello") + (((("\n\t" + "hello") + v3) + ("hello" + ("foo bar" + "foo bar"))) + (v3 + v3)))) + v3) + v3) + "foo bar") + ((v3 + "a") + ((v3 + v3) + (v3 + v3)))), F1: 679}
                                            for fi71 := 0; fi71 < 18; fi71++ {
                                                switch fi71 {
                                                case 128:
                                                    if (!(uint16(v2) >= 652)) {
                                                        delete(v6, 128)
                                                        v7 += (^(v7))
                                                        v4 = 1.5
                                                    }
                                                case 1:
                                                    if false {
                                                        var v72 uint64 = (v1 & (v1 >> ((255) & 63)))
                                                        _ = v72
                                                    }
                                                case 9223372036854775807:
                                                    if ((^(v1)) >= 18446744073709551615) {
                                                        var v73 uint64 = v1
                                                        v3 = ("\n\t" + (("foo bar" + ((v3 + "") + v3)) + "\n\t"))
                                                        _ = v73
                                                    }
                                                }
                                                v1 = v2
                                                _ = fi71
                                            }
                                            var v74 uint = 18446744073709551615
                                            _ = v70
                                            _ = v74
                                        } else {
                                            fmt.Println(v4)
                                        }
                                        var v75 int32 = -1
                                        var v76 bool = (false || (370 < 0))
                                        _ = v75
                                        _ = v76
                                    }
                                    v4 -= (float64(-2.25) * v4)
                                    fmt.Println(fi19)
                                }
                                var v77 uint8 = (v7 - v7)
                                switch "hello" {
                                case "foo bar":
                                    v5 -= ((float64((float32(float64((1e10 * 1.5))) / 3.14159)) - (3.14159 / (float64((float32(v4) + ((3.14159 / -2.25) - 0.0))) * ((v5 - v5) * (float64(0.0) - 0.0))))) + v4)
                                default:
                                    for fi78 := 0; fi78 < 4; fi78++ {
                                        v4 += (1e10 + -2.25)
                                        var v79 uint32 = 1
                                        _ = fi78
                                        _ = v79
                                    }
                                }
                                _ = v77
                            }
                        default:
                            fmt.Println(v7, v5, v4)
                        }
                        if false {
                            v80 := GenS5{F0: -2147483648, F1: v7, F2: 65535}
                            var v81 uint = 0
                            _ = v80
                            _ = v81
                        }
                        _ = fi19
                    }
                    var v82 uint64 = (v1 % ((v1) | 1))
                    v82 = (v82 - v2)
                    _ = v82
                } else {
                    switch ((("" + v3) + v3) + ("foo bar" + "foo bar")) {
                    case "\n\t":
                        v83 := v6[1]
                        _ = v83
                    case "a":
                        v84 := GenS7{F0: 0, F1: 4294967295, F2: v7}
                        _ = v84
                    case "":
                        var v85 uint16 = 754
                        _ = v85
                    default:
                        v86 := GenS4{F0: v3, F1: 2}
                        _ = v86
                    }
                    if true {
                        var v87 uint = 2
                        v88 := GenS8{F0: int64(v87), F1: -1}
                        for fi89 := 0; fi89 < 2; fi89++ {
                            var v90 uint = v87
                            _ = fi89
                            _ = v90
                        }
                        _ = v87
                        _ = v88
                    }
                    var v91 uint8 = 4
                    _ = v91
                }
                _ = v18
            }
        }
    } else {
        var v92 float64 = ((((1.5 / v5) + -2.25) / v4) + (-2.25 - 1.5))
        switch "a" {
        case "foo bar":
            var v93 int32 = -896
            _ = v93
        case "hello":
            v92 = v92
        }
        _ = v92
    }
    if (v3 <= ("" + v3)) {
        v94 := []int{0, -437}
        _ = v94
    } else {
        var v95 uint16 = 65535
        for fi96 := 0; fi96 < 14; fi96++ {
            for fi97 := 0; fi97 < 20; fi97++ {
                v5 = (v5 * -2.25)
                var v98 uint32 = 201
                var v99 uint32 = (v98 ^ 1)
                _ = fi97
                _ = v98
                _ = v99
            }
            _ = fi96
        }
        if true {
            v95 += v95
            var v100 int16 = 0
            if (3.14159 > (0.1 + (1e10 * 0.1))) {
                v4 = v4
                fmt.Println(v2, v1)
            } else {
                v101 := GenS9{F0: (v100 << ((1) & 63)), F1: (v100 + 32767)}
                var v102 int32 = -965
                v103 := GenS10{F0: -963, F1: 42}
                _ = v101
                _ = v102
                _ = v103
            }
            _ = v100
        } else {
            for fi104 := 0; fi104 < 5; fi104++ {
                v105 := GenS2{F0: v7, F1: (!true), F2: 317}
                _ = fi104
                _ = v105
            }
            v3 = v3
            switch "foo bar" {
            case "hello":
                v106 := GenS6{F0: 4294967295, F1: 42}
                _ = v106
            case "\n\t":
                var v107 int16 = 1
                _ = v107
            default:
                v108 := []int{128, 4294967295, -9223372036854775808, -670}
                _ = v108
            }
        }
        _ = v95
    }
    fmt.Println(v3)
    var v109 int64 = 1
    for fi110 := 0; fi110 < 7; fi110++ {
        for fi111 := 0; fi111 < 8; fi111++ {
            var v112 uint8 = v7
            var v113 int64 = (v109 >> ((737) & 63))
            v114 := v6[fi111]
            _ = fi111
            _ = v112
            _ = v113
            _ = v114
        }
        v115 := []int{}
        if (((v5 / float64(-2.25)) > 0.1) && (!true)) {
            for fi116 := 0; fi116 < 4; fi116++ {
                var v117 uint16 = 42
                v1 += v1
                v118 := v6[-1]
                _ = fi116
                _ = v117
                _ = v118
            }
            if len(v115) >= 2 {
                v115 = v115[1:len(v115)]
            }
            if (-8 < 1000) {
                var v119 int64 = v109
                fmt.Println(v5, v7)
                v120 := GenS11{F0: 0, F1: 0, F2: -805}
                _ = v119
                _ = v120
            } else {
                switch v3 {
                case "\n\t":
                    v121 := GenS2{F0: v7, F1: ("\n\t" > "a"), F2: 42}
                    _ = v121
                }
                if (-279 > -408) {
                    v122 := GenS3{F0: 2147483647, F1: 42, F2: v109}
                    v3 = "foo bar"
                    v123 := GenS8{F0: (v109 % ((v109) | 1)), F1: 5}
                    _ = v122
                    _ = v123
                } else {
                    if (0 == v1) {
                        for fi124 := 0; fi124 < 8; fi124++ {
                            var v125 uint32 = 247
                            _ = fi124
                            _ = v125
                        }
                        var v126 uint64 = 0
                        var v127 int16 = 581
                        _ = v126
                        _ = v127
                    } else {
                        switch -2147483648 {
                        case 0:
                            var v128 uint16 = 521
                            _ = v128
                        }
                        fmt.Println(v5)
                    }
                    fmt.Println(v1, v2, v5)
                    for fi129 := 0; fi129 < 12; fi129++ {
                        v5 -= (v4 - (((v4 + v4) + (0.0 / ((1e10 - v5) * v5))) - v5))
                        if true {
                            v5 = (v4 * -2.25)
                            var v130 int32 = 371
                            if len(v115) > 0 {
                                v115[0] = fi110
                            }
                            var v131 int16 = -1
                            _ = v130
                            _ = v131
                        }
                        v2 = (v2 & v2)
                        var v132 float64 = v5
                        v133 := GenS12{F0: 2147483647, F1: (v109 & (v109 % (((v109 * 0)) | 1))), F2: 0}
                        _ = fi129
                        _ = v132
                        _ = v133
                    }
                }
            }
        }
        _ = fi110
        _ = v115
    }
    switch 32768 {
    case -9223372036854775808:
        v5 = 3.14159
    default:
        v109 -= -1
    }
    var v134 int = -841
    v134 = (-(v134))
    v7 = (v7 - 101)
    v135 := []int{}
    v136 := GenS5{F0: 0, F1: 2, F2: uint32(v7)}
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v134:", v134)
    fmt.Println("slicelen v135:", len(v135), v135)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v109
    _ = v134
    _ = v135
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
