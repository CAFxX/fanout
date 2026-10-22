package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint; F1 uint16; F2 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 int8; F1 uint8; F2 int }
type GenS9 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int32; F1 uint8 }
type GenS12 struct { F0 uint64; F1 int8 }
type GenS13 struct { F0 uint; F1 int32 }
type GenS14 struct { F0 uint64; F1 int16; F2 bool }

func main() {
    var v1 uint64 = 2
    var v2 bool = (32767 >= -900)
    v3 := make(map[int]uint32)
    var v4 uint16 = 147
    var v5 string = "foo bar"
    switch -9223372036854775808 {
    case -294:
        var v6 bool = (194 == 215)
        _ = v6
    }
    var v7 int16 = 1000
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 bool = false
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 uint8 = 2
            switch v5 {
            case "":
                v12 := GenS1{F0: v11, F1: 261}
                _ = v12
            }
            if true {
                v2 = v9
                var v13 uint64 = v1
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        switch v5 {
        case "\n\t":
            if (false || (-2147483648 >= 2147483647)) {
                for fi14 := 0; fi14 < 2; fi14++ {
                    if (!true) {
                        if (42 <= 127) {
                            var v15 int16 = (v7 | v7)
                            _ = v15
                        } else {
                            if ((v2 && true) || (v2 && false)) {
                                if v2 {
                                    if (((1 < 1) || true) && v2) {
                                        v3[fi8] = 927
                                        var v16 int64 = -471
                                        _ = v16
                                    }
                                    switch "foo bar" {
                                    case "":
                                        _ = len(v3)
                                    case "hello":
                                        v17 := GenS1{F0: 128, F1: 991}
                                        _ = v17
                                    default:
                                        switch -988 {
                                        case 9223372036854775807:
                                            var v18 uint64 = (v1 ^ v1)
                                            if v2 {
                                                v19 := []int{fi14, (fi14 & (fi8 / (((fi8 / (((fi8 * (fi8 << ((2) & 63)))) | 1))) | 1))), fi14}
                                                if false {
                                                    var v20 uint = 915
                                                    v21 := GenS2{F0: -1}
                                                    _ = v20
                                                    _ = v21
                                                } else {
                                                    v5 = v5
                                                    var v22 int = fi8
                                                    _ = v22
                                                }
                                                switch "" {
                                                case "a":
                                                    if v9 {
                                                        if false {
                                                            if v2 {
                                                                var v23 bool = false
                                                                _ = v23
                                                            } else {
                                                                if (!v9) {
                                                                    v9 = v9
                                                                    switch v5 {
                                                                    case "a":
                                                                        v24 := v3[fi14]
                                                                        _ = v24
                                                                    case "\n\t":
                                                                        var v25 int16 = v7
                                                                        _ = v25
                                                                    case "hello":
                                                                        var v26 uint64 = v18
                                                                        _ = v26
                                                                    }
                                                                    v27 := GenS2{F0: 548}
                                                                    _ = v27
                                                                } else {
                                                                    v18 = 651
                                                                }
                                                                if (133 > 127) {
                                                                    _ = len(v19)
                                                                    if len(v19) >= 2 {
                                                                        v19 = v19[1:len(v19)]
                                                                    }
                                                                } else {
                                                                    v28 := GenS3{F0: 2, F1: -199}
                                                                    v19 = append(v19, fi8)
                                                                    _ = len(v3)
                                                                    _ = v28
                                                                }
                                                                if ((-(fi8)) > (fi14 / (((fi14 & fi8)) | 1))) {
                                                                    if (-45 >= 1) {
                                                                        var v29 int32 = 0
                                                                        _ = v29
                                                                    }
                                                                }
                                                            }
                                                            switch v5 {
                                                            case "":
                                                                switch (v5 + ((v5 + v5) + ("foo bar" + ("\n\t" + v5)))) {
                                                                case "foo bar":
                                                                    if len(v19) >= 2 {
                                                                        v19 = v19[1:len(v19)]
                                                                    }
                                                                default:
                                                                    fmt.Println(fi8)
                                                                }
                                                            case "\n\t":
                                                                if (-9223372036854775808 != -630) {
                                                                    _ = len(v19)
                                                                    v5 = v5
                                                                }
                                                            default:
                                                                v30 := GenS3{F0: 128, F1: int32(fi8)}
                                                                _ = v30
                                                            }
                                                        } else {
                                                            var v31 string = "foo bar"
                                                            v32 := GenS4{F0: 419, F1: 32767, F2: 1}
                                                            _ = v31
                                                            _ = v32
                                                        }
                                                        switch v5 {
                                                        case "a":
                                                            var v33 uint32 = 2147483648
                                                            _ = v33
                                                        case "foo bar":
                                                            var v34 bool = (v9 && (!false))
                                                            _ = v34
                                                        case "":
                                                            switch fi14 {
                                                            case 32768:
                                                                var v35 uint8 = 1
                                                                _ = v35
                                                            case -1:
                                                                v9 = v2
                                                            case 1:
                                                                if (fi14 != -612) {
                                                                    v7 = (v7 + v7)
                                                                    fmt.Println(fi14)
                                                                } else {
                                                                    switch fi8 {
                                                                    case 127:
                                                                        if (float32(3.14159) <= (3.14159 - 0.1)) {
                                                                            var v36 float32 = 3.14159
                                                                            switch v5 {
                                                                            case "foo bar":
                                                                                if v9 {
                                                                                    switch (fi8 | fi8) {
                                                                                    case 32767:
                                                                                        v1 -= 42
                                                                                    default:
                                                                                        v5 = "hello"
                                                                                        fmt.Println(v5, v1, v36)
                                                                                    }
                                                                                } else {
                                                                                    var v37 int64 = 1
                                                                                    _ = v37
                                                                                }
                                                                            default:
                                                                                v4 -= (v4 | v4)
                                                                            }
                                                                            _ = v36
                                                                        } else {
                                                                            v4 -= 42
                                                                            var v38 uint32 = 2
                                                                            if v2 {
                                                                                switch v5 {
                                                                                case "\n\t":
                                                                                    var v39 int = fi8
                                                                                    _ = v39
                                                                                case "":
                                                                                    var v40 uint16 = v4
                                                                                    _ = v40
                                                                                }
                                                                                v5 = "foo bar"
                                                                                v41 := GenS2{F0: -470}
                                                                                _ = v41
                                                                            }
                                                                            _ = len(v3)
                                                                            _ = v38
                                                                        }
                                                                    case 4294967295:
                                                                        var v42 uint8 = 127
                                                                        _ = v42
                                                                    case 984:
                                                                        switch v5 {
                                                                        case "foo bar":
                                                                            v43 := GenS1{F0: 2, F1: 4294967295}
                                                                            _ = v43
                                                                        case "a":
                                                                            var v44 uint64 = v18
                                                                            var v45 int32 = -1
                                                                            var v46 int8 = int8(v1)
                                                                            _ = v44
                                                                            _ = v45
                                                                            _ = v46
                                                                        case "":
                                                                            v9 = false
                                                                        default:
                                                                            v47 := GenS4{F0: 147, F1: (v4 ^ (v4 ^ (v4 >> ((uint(v1)) & 63)))), F2: 42}
                                                                            if v9 {
                                                                                var v48 uint64 = (v1 - (v18 / ((v1) | 1)))
                                                                                v3[fi8] = 65535
                                                                                _ = v48
                                                                            }
                                                                            v5 = (v5 + v5)
                                                                            var v49 float32 = ((float32(float64(3.14159)) / ((1.5 / 1.0) / 3.14159)) + float32(0.1))
                                                                            v50 := v3[(fi8 % (((-(fi14))) | 1))]
                                                                            _ = v47
                                                                            _ = v49
                                                                            _ = v50
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            v51 := GenS2{F0: 0}
                                                                            _ = v51
                                                                        }
                                                                    }
                                                                    var v52 string = v5
                                                                    _ = v52
                                                                }
                                                            }
                                                        default:
                                                            var v53 string = v5
                                                            _ = v53
                                                        }
                                                    }
                                                    switch (("a" + v5) + ((("foo bar" + (v5 + v5)) + ((("foo bar" + (((v5 + (v5 + v5)) + (v5 + v5)) + "a")) + (v5 + "\n\t")) + v5)) + v5)) {
                                                    case "":
                                                        if v2 {
                                                            if len(v19) >= 2 {
                                                                v19 = v19[1:len(v19)]
                                                            }
                                                            if len(v19) >= 2 {
                                                                v19 = v19[1:len(v19)]
                                                            }
                                                            if false {
                                                                v1 -= 176
                                                            } else {
                                                                v19 = append(v19, (fi14 | fi8))
                                                            }
                                                        } else {
                                                            v19 = append(v19, (fi14 << ((65535) & 63)))
                                                            v5 = "a"
                                                            var v54 int64 = 1
                                                            v2 = (("\n\t" + "\n\t") >= v5)
                                                            var v55 int16 = v7
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                    case "\n\t":
                                                        v56 := GenS4{F0: 18446744073709551615, F1: (v4 + 32767), F2: 91}
                                                        _ = v56
                                                    case "hello":
                                                        switch v5 {
                                                        case "a":
                                                            if len(v19) >= 2 {
                                                                v19 = v19[1:len(v19)]
                                                            }
                                                        case "":
                                                            var v57 int64 = int64(v7)
                                                            _ = v57
                                                        case "hello":
                                                            var v58 float64 = (1e10 - (-2.25 * -2.25))
                                                            _ = v58
                                                        default:
                                                            v59 := GenS3{F0: 43, F1: 945}
                                                            fmt.Println(v4, fi14, v4)
                                                            if v9 {
                                                                v60 := GenS5{F0: 127}
                                                                v5 = v5
                                                                if v2 {
                                                                    v18 -= 2
                                                                    v61 := GenS6{F0: (v7 ^ int16(v4)), F1: 2}
                                                                    v5 = "a"
                                                                    _ = v61
                                                                }
                                                                if true {
                                                                    v62 := GenS7{F0: fi8, F1: 48}
                                                                    var v63 uint = uint(v1)
                                                                    _ = v62
                                                                    _ = v63
                                                                }
                                                                var v64 string = (("a" + "hello") + v5)
                                                                _ = len(v19)
                                                                _ = v60
                                                                _ = v64
                                                            } else {
                                                                if len(v19) >= 2 {
                                                                    v19 = v19[1:len(v19)]
                                                                }
                                                                if len(v19) > 0 {
                                                                    v19[0] = fi14
                                                                }
                                                            }
                                                            _ = v59
                                                        }
                                                    default:
                                                        v18 += v1
                                                    }
                                                    fmt.Println(v1)
                                                    v2 = ((3.14159 <= 1.5) && v2)
                                                    var v65 uint64 = (v1 >> ((42) & 63))
                                                    _ = v65
                                                }
                                                _ = v19
                                            } else {
                                                if (v2 && (!(!v2))) {
                                                    v9 = (797 != 4294967295)
                                                    switch v5 {
                                                    case "a":
                                                        v5 = (v5 + "hello")
                                                    }
                                                    switch "hello" {
                                                    case "":
                                                        v5 = ((v5 + "foo bar") + v5)
                                                    case "foo bar":
                                                        var v66 uint32 = 6
                                                        var v67 float64 = 0.0
                                                        v68 := GenS8{F0: 0, F1: 255, F2: (^(fi14))}
                                                        _ = v66
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                } else {
                                                    var v69 int8 = -115
                                                    switch (((v5 + (("" + "hello") + v5)) + (("a" + v5) + v5)) + ((v5 + "") + (("" + ("" + "")) + (v5 + (v5 + "foo bar"))))) {
                                                    case "hello":
                                                        if v2 {
                                                            var v70 int8 = v69
                                                            if (3.14159 != 0.0) {
                                                                var v71 int16 = int16(v4)
                                                                _ = v71
                                                            }
                                                            _ = v70
                                                        } else {
                                                            if true {
                                                                v72 := GenS1{F0: 1, F1: uint(v4)}
                                                                _ = v72
                                                            }
                                                            switch (-(fi14)) {
                                                            case -1:
                                                                var v73 int8 = 1
                                                                _ = v73
                                                            }
                                                            v74 := []int{fi8, int(v7), fi14, 127}
                                                            _ = v74
                                                        }
                                                    case "\n\t":
                                                        v1 += v1
                                                    }
                                                    v9 = false
                                                    fmt.Println(v18, v1, v5)
                                                    fmt.Println(v4)
                                                    v75 := GenS9{F0: 2, F1: -666, F2: 87}
                                                    _ = v69
                                                    _ = v75
                                                }
                                                v76 := GenS10{F0: 18446744073709551615}
                                                var v77 float32 = 0.0
                                                _ = v76
                                                _ = v77
                                            }
                                            _ = v18
                                        case 384:
                                            if ((v2 || (false || (!v2))) && v9) {
                                                var v78 bool = (933 < 110)
                                                _ = v78
                                            } else {
                                                v3[(fi8 + (fi8 / (((fi14 ^ (fi14 >> ((18446744073709551615) & 63)))) | 1)))] = 65535
                                                v7 = v7
                                            }
                                        default:
                                            v7 -= (v7 ^ v7)
                                        }
                                    }
                                }
                                _ = len(v3)
                                var v79 int64 = 16
                                _ = v79
                            }
                        }
                        v2 = v2
                        if ((true && (v9 || (1 < 127))) && (-103 >= -116)) {
                            if v9 {
                                v9 = false
                                if true {
                                    switch (fi8 >> ((4294967295) & 63)) {
                                    case 32768:
                                        var v80 uint8 = 255
                                        _ = v80
                                    case -9223372036854775808:
                                        fmt.Println(fi14)
                                    case 127:
                                        var v81 int8 = -14
                                        _ = v81
                                    default:
                                        if (true && v9) {
                                            fmt.Println(v2, v7)
                                            v7 = (^(v7))
                                        } else {
                                            var v82 bool = v9
                                            if ((!(((float32(((-2.25 / 1.5) * (1e10 - float64(-2.25)))) != (-2.25 - 1.5)) || v2) || v82)) && true) {
                                                switch v5 {
                                                case "hello":
                                                    var v83 float32 = 1.5
                                                    _ = v83
                                                case "a":
                                                    v1 = (^(v1))
                                                default:
                                                    v84 := GenS11{F0: -231, F1: 82}
                                                    _ = v84
                                                }
                                                switch fi8 {
                                                case -1:
                                                    v4 = (^(v4))
                                                case 439:
                                                    if true {
                                                        var v85 int8 = 121
                                                        _ = v85
                                                    }
                                                }
                                            } else {
                                                var v86 uint = 82
                                                _ = v86
                                            }
                                            _ = v82
                                        }
                                    }
                                    var v87 uint16 = (^(v4))
                                    _ = v87
                                } else {
                                    v88 := []int{}
                                    v5 = v5
                                    var v89 bool = v9
                                    _ = v88
                                    _ = v89
                                }
                            } else {
                                var v90 int32 = 188
                                v91 := []int{}
                                v92 := GenS7{F0: (fi8 * (fi14 - fi8)), F1: 243}
                                if v9 {
                                    var v93 uint8 = 255
                                    var v94 string = (("a" + v5) + (v5 + "a"))
                                    _ = v93
                                    _ = v94
                                } else {
                                    var v95 float32 = -2.25
                                    if len(v91) >= 2 {
                                        v91 = v91[1:len(v91)]
                                    }
                                    var v96 uint8 = 127
                                    _ = v95
                                    _ = v96
                                }
                                v7 -= (v7 >> ((54) & 63))
                                _ = v90
                                _ = v91
                                _ = v92
                            }
                            v97 := GenS7{F0: (fi14 - -635), F1: 42}
                            var v98 int8 = -128
                            _ = v97
                            _ = v98
                        } else {
                            var v99 uint32 = 21
                            if ((v5 + v5) < v5) {
                                v5 = "foo bar"
                                if v2 {
                                    delete(v3, fi14)
                                } else {
                                    var v100 string = ((((("hello" + v5) + "") + v5) + (v5 + v5)) + v5)
                                    _ = v100
                                }
                            }
                            var v101 float64 = (((0.1 * 3.14159) - (0.1 / -2.25)) + float64(float32((1.5 + (1.5 - float64(1e10))))))
                            _ = v101
                            _ = v99
                        }
                        v102 := GenS12{F0: v1, F1: 48}
                        v5 = (v5 + v5)
                        _ = v102
                    } else {
                        fmt.Println(v2)
                    }
                    _ = fi14
                }
            } else {
                fmt.Println(fi8, v5)
                if (int32(fi8) >= 468) {
                    for fi103 := 0; fi103 < 14; fi103++ {
                        v104 := []int{fi103, 127}
                        v104 = append(v104, (fi8 + fi103))
                        _ = fi103
                        _ = v104
                    }
                    v105 := GenS12{F0: 18446744073709551615, F1: -128}
                    _ = v105
                } else {
                    if ((v1 | (v1 * v1)) > uint64(v4)) {
                        v106 := GenS3{F0: 47, F1: -1}
                        _ = v106
                    } else {
                        v4 += (v4 & (^(v4)))
                        var v107 uint32 = 0
                        _ = v107
                    }
                    fmt.Println(fi8, v9)
                }
                v108 := v3[(fi8 / (((fi8 ^ 4294967295)) | 1))]
                _ = v108
            }
        case "":
            for fi109 := 0; fi109 < 13; fi109++ {
                v9 = (42 <= 1)
                var v110 string = (("\n\t" + "a") + ((v5 + (v5 + "")) + (v5 + v5)))
                _ = fi109
                _ = v110
            }
        case "hello":
            if v9 {
                v111 := GenS3{F0: 128, F1: -1}
                fmt.Println(v9)
                _ = v111
            }
        }
        _ = fi8
        _ = v9
    }
    var v112 uint64 = v1
    var v113 uint64 = 767
    for fi114 := 0; fi114 < 1; fi114++ {
        for fi115 := 0; fi115 < 11; fi115++ {
            var v116 int64 = 192
            _ = fi115
            _ = v116
        }
        v117 := []int{fi114, (-(fi114)), (^(fi114)), (fi114 | (fi114 << ((0) & 63)))}
        _ = fi114
        _ = v117
    }
    var v118 int8 = 124
    v119 := GenS13{F0: 255, F1: -1}
    for fi120 := 0; fi120 < 11; fi120++ {
        var v121 int = fi120
        if true {
            var v122 int16 = 899
            v123 := GenS14{F0: (v112 * (v112 % ((v113) | 1))), F1: v122, F2: (!((!true) || ("a" <= v5)))}
            var v124 uint16 = v4
            _ = v122
            _ = v123
            _ = v124
        }
        _ = fi120
        _ = v121
    }
    fmt.Println(v7, v5)
    for fi125 := 0; fi125 < 17; fi125++ {
        var v126 bool = v2
        _ = fi125
        _ = v126
    }
    var v127 bool = false
    fmt.Println(v2, v7, v127)
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v113:", v113)
    fmt.Println("v118:", v118)
    fmt.Println("v127:", v127)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v112
    _ = v113
    _ = v118
    _ = v119
    _ = v127
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
