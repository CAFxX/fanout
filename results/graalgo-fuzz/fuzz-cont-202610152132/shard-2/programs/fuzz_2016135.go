package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint64 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 bool }
func genFn0() uint64 {
    v5 := []int{42, 1, -327, 9223372036854775807}
    var v6 uint64 = 1
    _ = v5
    _ = v6
    return 654
}
func genFn1() int32 {
    if (!("hello" != "a")) {
        if (-128 == 0) {
            for fi7 := 0; fi7 < 16; fi7++ {
                v8 := []int{32767, (fi7 * 128)}
                if false {
                    var v9 int = fi7
                    v9 = v9
                    _ = v9
                }
                _ = fi7
                _ = v8
            }
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := []int{596, -2147483648, fi10}
            v11 = append(v11, (fi10 - fi10))
            if (42 == 128) {
                var v12 int64 = -1
                v12 = v12
                _ = v12
            }
            _ = fi10
            _ = v11
        }
        if false {
            switch 32767 {
            case 0:
                var v13 uint32 = 195
                _ = v13
            case 2147483647:
                if false {
                    for fi14 := 0; fi14 < 13; fi14++ {
                        v15 := GenS1{F0: -850, F1: -615, F2: genFn0()}
                        if false {
                            if true {
                                if false {
                                    var v16 int32 = 1
                                    _ = v16
                                } else {
                                    var v17 uint32 = 2147483647
                                    var v18 int = (fi14 - 92)
                                    for fi19 := 0; fi19 < 4; fi19++ {
                                        var v20 bool = (595 == 1)
                                        _ = fi19
                                        _ = v20
                                    }
                                    _ = v17
                                    _ = v18
                                }
                                for fi21 := 0; fi21 < 10; fi21++ {
                                    switch "hello" {
                                    case "foo bar":
                                        var v22 float32 = 0.0
                                        _ = v22
                                    case "hello":
                                        var v23 float32 = -2.25
                                        _ = v23
                                    case "a":
                                        v24 := []int{(fi14 - fi21), 96, fi21}
                                        _ = v24
                                    default:
                                        v25 := []int{}
                                        _ = v25
                                    }
                                    _ = fi21
                                }
                            }
                        }
                        fmt.Println(fi14)
                        _ = fi14
                        _ = v15
                    }
                    var v26 uint = 102
                    var v27 int = -810
                    _ = v26
                    _ = v27
                } else {
                    v28 := []int{}
                    _ = v28
                }
            case -530:
                var v29 uint = 255
                _ = v29
            default:
                for fi30 := 0; fi30 < 4; fi30++ {
                    v31 := []int{(fi30 << ((18446744073709551615) & 63)), -828, (fi30 ^ -9223372036854775808), (fi30 - -9223372036854775808)}
                    _ = fi30
                    _ = v31
                }
            }
            var v32 uint16 = 2
            v32 -= (^(v32))
            _ = v32
        }
    } else {
        var v33 bool = false
        v33 = (127 == 9223372036854775807)
        _ = v33
    }
    var v34 int = 32767
    switch "foo bar" {
    case "\n\t":
        v35 := []int{(v34 ^ (v34 % (((v34 * v34)) | 1))), v34}
        _ = v35
    case "hello":
        var v36 float32 = (-2.25 * (((float32((float64(1e10) / 1.5)) * -2.25) - ((1e10 - (0.0 / (float32(1e10) + float32(0.0)))) / float32(-2.25))) + ((1.5 / ((0.0 - 1e10) / (float32(0.0) - ((float32(0.0) * (((1e10 + (((3.14159 + 1e10) + 3.14159) / 3.14159)) + (0.1 / 1.5)) * (0.0 - (1e10 + (((0.1 / float32(-2.25)) - (float32(0.1) / (0.1 / 1.5))) * float32(0.1)))))) - ((0.1 / (1e10 - -2.25)) + 0.1))))) - 1.5)))
        _ = v36
    case "foo bar":
        for fi37 := 0; fi37 < 5; fi37++ {
            if (127 > -128) {
                v34 = fi37
            } else {
                var v38 uint8 = 1
                for fi39 := 0; fi39 < 20; fi39++ {
                    var v40 float64 = 0.1
                    _ = fi39
                    _ = v40
                }
                var v41 float64 = -2.25
                _ = v38
                _ = v41
            }
            _ = fi37
        }
    }
    v42 := []int{v34, v34}
    if true {
        v34 = (v34 << ((2) & 63))
        var v43 int8 = -1
        _ = v43
    } else {
        v44 := GenS1{F0: 2147483647, F1: v34, F2: 4294967295}
        switch ("\n\t" + ("" + ("\n\t" + "\n\t"))) {
        case "foo bar":
            for fi45 := 0; fi45 < 11; fi45++ {
                if len(v42) >= 2 {
                    v42 = v42[1:len(v42)]
                }
                _ = len(v42)
                _ = fi45
            }
        case "\n\t":
            var v46 uint64 = 669
            _ = v46
        }
        if (-1 > 2147483647) {
            var v47 uint = 0
            fmt.Println(v47, v47)
            var v48 bool = ((!true) && false)
            _ = v47
            _ = v48
        } else {
            var v49 uint8 = 128
            var v50 uint8 = v49
            var v51 uint8 = v49
            _ = v49
            _ = v50
            _ = v51
        }
        _ = v44
    }
    _ = v34
    _ = v42
    return 2147483647
}

func main() {
    var v1 uint16 = 65535
    var v2 int32 = -89
    var v3 float64 = -2.25
    v4 := []string{""}
    v52 := GenS2{F0: "a", F1: 21}
    for fi53 := 0; fi53 < 3; fi53++ {
        v54 := GenS1{F0: v2, F1: (fi53 + fi53), F2: 4294967295}
        _ = fi53
        _ = v54
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v55 := GenS1{F0: v2, F1: 1, F2: 0}
    for fi56 := 0; fi56 < 12; fi56++ {
        var v57 int64 = 42
        if (!true) {
            v57 += 1
            for fi58 := 0; fi58 < 6; fi58++ {
                fmt.Println(v1, fi56)
                if (!true) {
                    var v59 int = fi56
                    if true {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    fmt.Println(fi56)
                    _ = v59
                } else {
                    v3 = v3
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                    _ = len(v4)
                }
                switch "hello" {
                case "foo bar":
                    if false {
                        var v60 string = "hello"
                        var v61 int64 = v57
                        fmt.Println(fi58, v2, v60)
                        if false {
                            v57 = (v57 - v61)
                            switch "a" {
                            case "a":
                                v1 -= v1
                            case "":
                                if len(v4) > 0 {
                                    v4[0] = "a"
                                }
                                if true {
                                    var v62 uint8 = 28
                                    v2 = (v2 >> ((864) & 63))
                                    var v63 int16 = -299
                                    _ = v62
                                    _ = v63
                                } else {
                                    var v64 float32 = 1e10
                                    if false {
                                        if (42 >= 2147483647) {
                                            v65 := GenS1{F0: (v2 % ((2147483647) | 1)), F1: fi56, F2: 967}
                                            var v66 float64 = v3
                                            _ = v65
                                            _ = v66
                                        } else {
                                            var v67 float32 = (float32(3.14159) * (((v64 + (v64 * ((((float32(v3) + v64) * (v64 * v64)) * v64) * v64))) / (float32(float64(float32((v3 - 1e10)))) * v64)) + float32(-2.25)))
                                            v67 = v64
                                            var v68 float64 = (v3 / v3)
                                            _ = v67
                                            _ = v68
                                        }
                                    } else {
                                        var v69 int16 = int16(v61)
                                        var v70 uint8 = 27
                                        switch 128 {
                                        case 127:
                                            if len(v4) > 0 {
                                                v4[0] = v60
                                            }
                                        case 157:
                                            var v71 int = (fi56 << ((42) & 63))
                                            _ = v71
                                        case 1:
                                            var v72 int16 = v69
                                            _ = v72
                                        default:
                                            v2 -= (v2 >> ((1) & 63))
                                        }
                                        _ = v69
                                        _ = v70
                                    }
                                    var v73 int64 = (-(v61))
                                    _ = v64
                                    _ = v73
                                }
                            case "foo bar":
                                fmt.Println(v1, v1, v60)
                                v2 -= v2
                                var v74 uint16 = (v1 & v1)
                                _ = v74
                            default:
                                if true {
                                    fmt.Println(v2, v2, fi58)
                                    var v75 float64 = v3
                                    var v76 string = (v60 + "\n\t")
                                    _ = v75
                                    _ = v76
                                } else {
                                    if false {
                                        fmt.Println(v2, v3, v2)
                                        v4 = append(v4, v60)
                                    }
                                }
                            }
                            var v77 bool = true
                            _ = v77
                        }
                        var v78 float32 = ((0.1 + float32(((v3 / 1e10) - (0.1 / (0.0 + v3))))) / -2.25)
                        _ = v60
                        _ = v61
                        _ = v78
                    }
                default:
                    var v79 int8 = -128
                    _ = v79
                }
                _ = fi58
            }
            v3 = 0.1
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi56
        _ = v57
    }
    for fi80 := 0; fi80 < 2; fi80++ {
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        _ = fi80
    }
    switch -2147483648 {
    case 128:
        if false {
            v1 -= v1
            switch "\n\t" {
            case "":
                var v81 int16 = -20
                _ = v81
            case "hello":
                fmt.Println(v3)
            case "foo bar":
                v82 := GenS1{F0: v2, F1: 713, F2: 18446744073709551615}
                _ = v82
            }
            for fi83 := 0; fi83 < 8; fi83++ {
                _ = len(v4)
                if true {
                    if true {
                        v84 := GenS3{F0: 376, F1: 1}
                        if (!false) {
                            var v85 int32 = (^(v2))
                            var v86 float64 = v3
                            _ = v85
                            _ = v86
                        } else {
                            v1 += 42
                        }
                        _ = v84
                    }
                    v2 = (-(v2))
                    var v87 uint32 = 415
                    _ = v87
                } else {
                    v3 += float64((((1.5 * (float32(v3) * float32(((-2.25 + 0.0) - -2.25)))) + -2.25) - ((-2.25 * (-2.25 / 1.0)) - (1.5 + float32(v3)))))
                }
                _ = fi83
            }
        } else {
            v2 -= -512
            var v88 int16 = -871
            _ = v88
        }
    }
    var v89 string = "a"
    v90 := GenS2{F0: (v89 + "foo bar"), F1: 127}
    v2 = -883
    v89 = (("foo bar" + "\n\t") + v89)
    var v91 int32 = (v2 + (v2 / ((42) | 1)))
    for fi92 := 0; fi92 < 7; fi92++ {
        v93 := GenS4{F0: v2, F1: "foo bar"}
        if false {
            v94 := GenS2{F0: "\n\t", F1: 255}
            _ = v94
        }
        v95 := GenS3{F0: 9223372036854775807, F1: 4294967295}
        _ = fi92
        _ = v93
        _ = v95
    }
    if (genFn0() <= 2) {
        for fi96 := 0; fi96 < 18; fi96++ {
            switch fi96 {
            case 4294967295:
                v97 := GenS3{F0: 9223372036854775807, F1: 42}
                _ = v97
            case -1:
                v89 = v89
            }
            _ = fi96
        }
        _ = len(v4)
        fmt.Println(v1)
    }
    v98 := GenS5{F0: 0, F1: 1}
    v99 := GenS2{F0: (v89 + v89), F1: 0}
    if (((3.14159 * 0.1) * 0.1) != (0.0 * float32(0.1))) {
        fmt.Println(v91, v89, v3)
        v100 := GenS5{F0: 0, F1: 0}
        for fi101 := 0; fi101 < 6; fi101++ {
            var v102 int = fi101
            v103 := GenS2{F0: "hello", F1: 255}
            if (v3 >= v3) {
                var v104 int64 = 671
                var v105 int8 = -128
                var v106 uint64 = 4294967295
                _ = v104
                _ = v105
                _ = v106
            }
            _ = fi101
            _ = v102
            _ = v103
        }
        _ = v100
    } else {
        v89 = v89
        if (!((v1 >> ((1) & 63)) > 255)) {
            v107 := GenS4{F0: genFn1(), F1: ("foo bar" + ((v89 + v89) + v89))}
            v4 = append(v4, v89)
            if false {
                fmt.Println(v2, v2, v3)
                var v108 uint16 = v1
                switch 0 {
                case -2147483648:
                    for fi109 := 0; fi109 < 13; fi109++ {
                        v110 := GenS6{F0: fi109, F1: 15}
                        v111 := GenS7{F0: false}
                        if ((float32((v3 + v3)) / 1.0) > (1e10 * (1.5 / 1.0))) {
                            switch v89 {
                            case "\n\t":
                                if false {
                                    v91 -= v91
                                    var v112 float32 = -2.25
                                    v113 := GenS3{F0: 9223372036854775807, F1: 18446744073709551615}
                                    _ = v112
                                    _ = v113
                                }
                            default:
                                if ((!(!((0 >= 42) && false))) && true) {
                                    for fi114 := 0; fi114 < 1; fi114++ {
                                        v115 := GenS3{F0: -9223372036854775808, F1: 65535}
                                        if (1 != 42) {
                                            if true {
                                                fmt.Println(v91)
                                                _ = len(v4)
                                            } else {
                                                v4 = append(v4, v89)
                                                if true {
                                                    v116 := GenS3{F0: 1, F1: 255}
                                                    var v117 int8 = -1
                                                    _ = v116
                                                    _ = v117
                                                } else {
                                                    var v118 uint64 = 0
                                                    v119 := GenS2{F0: v89, F1: 10}
                                                    _ = v118
                                                    _ = v119
                                                }
                                                v91 += (v91 * (v2 / ((genFn1()) | 1)))
                                            }
                                            fmt.Println(v2, v91)
                                        } else {
                                            v1 = 639
                                        }
                                        var v120 uint8 = 0
                                        v2 = v2
                                        if (!true) {
                                            v108 = uint16(v120)
                                            var v121 string = ("foo bar" + v89)
                                            _ = v121
                                        } else {
                                            if len(v4) > 0 {
                                                v4[0] = v89
                                            }
                                        }
                                        var v122 uint8 = (v120 + (^(v120)))
                                        _ = fi114
                                        _ = v115
                                        _ = v120
                                        _ = v122
                                    }
                                    v91 -= v2
                                } else {
                                    var v123 uint8 = 1
                                    if ((1 <= 0) || false) {
                                        var v124 int8 = 42
                                        switch -1 {
                                        case 127:
                                            var v125 int = (-(fi109))
                                            _ = v125
                                        case -466:
                                            for fi126 := 0; fi126 < 1; fi126++ {
                                                if true {
                                                    v127 := GenS6{F0: int(v108), F1: v123}
                                                    v3 += (1e10 * v3)
                                                    _ = v127
                                                }
                                                _ = fi126
                                            }
                                        case 128:
                                            var v128 uint64 = 2
                                            _ = v128
                                        }
                                        var v129 float64 = (-2.25 / v3)
                                        _ = v124
                                        _ = v129
                                    } else {
                                        switch fi109 {
                                        case 32767:
                                            v3 += v3
                                        case 128:
                                            v130 := GenS6{F0: (fi109 % ((127) | 1)), F1: 127}
                                            _ = v130
                                        }
                                    }
                                    _ = v123
                                }
                            }
                            if (-1 < 0) {
                                v4 = append(v4, "\n\t")
                                var v131 uint64 = 2
                                v3 += v3
                                _ = v131
                            } else {
                                if true {
                                    if (fi109 < (fi109 | fi109)) {
                                        var v132 string = ((((v89 + v89) + "hello") + "\n\t") + "a")
                                        var v133 uint32 = 42
                                        v133 = v133
                                        _ = v132
                                        _ = v133
                                    } else {
                                        var v134 uint32 = 244
                                        _ = v134
                                    }
                                    switch "" {
                                    case "a":
                                        for fi135 := 0; fi135 < 2; fi135++ {
                                            v1 += v1
                                            v108 = (v1 * v108)
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = fi135
                                        }
                                    case "\n\t":
                                        _ = len(v4)
                                    case "foo bar":
                                        var v136 uint = 187
                                        _ = v136
                                    }
                                    v137 := GenS7{F0: false}
                                    _ = v137
                                }
                            }
                        }
                        _ = fi109
                        _ = v110
                        _ = v111
                    }
                case -274:
                    var v138 uint = 1
                    _ = v138
                case 9223372036854775807:
                    var v139 uint = 832
                    _ = v139
                }
                _ = v108
            }
            _ = v107
        } else {
            var v140 int64 = 858
            if (0.0 == 3.14159) {
                v4 = append(v4, v89)
            }
            _ = v140
        }
        fmt.Println(v3)
    }
    if (!true) {
        fmt.Println(v1)
    } else {
        for fi141 := 0; fi141 < 11; fi141++ {
            var v142 uint8 = 42
            _ = fi141
            _ = v142
        }
        for fi143 := 0; fi143 < 16; fi143++ {
            v91 = (v2 >> ((1) & 63))
            _ = fi143
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v89:", v89)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v55
    _ = v89
    _ = v90
    _ = v91
    _ = v98
    _ = v99
}
