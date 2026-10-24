package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 bool }
type GenS2 struct { F0 string; F1 int64; F2 int32 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS6 struct { F0 bool; F1 string; F2 int64 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS9 struct { F0 int64; F1 uint8 }
type GenS10 struct { F0 int16; F1 int16; F2 int64 }
type GenS11 struct { F0 int8; F1 uint; F2 uint64 }
type GenS12 struct { F0 int; F1 uint8; F2 uint }
type GenS13 struct { F0 int64; F1 int8 }
func genFn0() int32 {
    switch 32767 {
    case -18:
        var v7 float32 = -2.25
        _ = v7
    case 128:
        if (0.0 <= 3.14159) {
            var v8 int8 = 100
            _ = v8
        } else {
            var v9 bool = (18446744073709551615 == 18446744073709551615)
            var v10 string = (("hello" + "") + "hello")
            var v11 uint32 = 340
            _ = v10
            _ = v11
            _ = v9
        }
    case -370:
        var v12 int16 = -865
        _ = v12
    default:
        if false {
            v13 := []int{128}
            for fi14 := 0; fi14 < 5; fi14++ {
                v15 := GenS1{F0: 827, F1: (((3.14159 + 0.0) + float64((-2.25 * float32(3.14159)))) < 3.14159), F2: false}
                fmt.Println(fi14)
                _ = fi14
                _ = v15
            }
            _ = v13
        } else {
            if false {
                var v16 uint = 731
                v17 := GenS2{F0: "hello", F1: -9223372036854775808, F2: 0}
                _ = v16
                _ = v17
            } else {
                var v18 uint = 42
                _ = v18
            }
            v19 := GenS1{F0: 2, F1: false, F2: (897 < -384)}
            switch 32768 {
            case -752:
                var v20 uint8 = 118
                _ = v20
            case 160:
                var v21 float64 = 1.5
                _ = v21
            default:
                for fi22 := 0; fi22 < 20; fi22++ {
                    v23 := GenS1{F0: 684, F1: true, F2: (1000 >= 1)}
                    _ = fi22
                    _ = v23
                }
            }
            _ = v19
        }
    }
    switch -2147483648 {
    case 752:
        var v24 uint = 450
        _ = v24
    case 128:
        v25 := GenS3{F0: 1, F1: 0, F2: 0}
        _ = v25
    default:
        var v26 bool = ((1e10 - (-2.25 + 0.0)) == 3.14159)
        _ = v26
    }
    var v27 string = "\n\t"
    var v28 uint16 = 1
    if (false && false) {
        var v29 int = -2147483648
        var v30 string = v27
        _ = v29
        _ = v30
    } else {
        for fi31 := 0; fi31 < 9; fi31++ {
            for fi32 := 0; fi32 < 12; fi32++ {
                var v33 int8 = 42
                v28 = v28
                switch v27 {
                case "foo bar":
                    var v34 uint32 = 4294967295
                    _ = v34
                case "":
                    if (42 != 1000) {
                        var v35 float64 = 3.14159
                        v36 := []int{fi31, fi32, (-(fi31)), (fi32 << ((4294967295) & 63))}
                        v35 += 0.0
                        _ = v35
                        _ = v36
                    } else {
                        v27 = v27
                        fmt.Println(v33)
                    }
                case "hello":
                    v37 := []int{(fi32 & 1), 128, fi32}
                    _ = v37
                default:
                    v38 := []int{-88, -1, fi32}
                    if len(v38) >= 2 {
                        v38 = v38[1:len(v38)]
                    }
                    switch v27 {
                    case "":
                        v39 := GenS1{F0: (^(v28)), F1: false, F2: (288 < 65535)}
                        _ = v39
                    case "\n\t":
                        var v40 int32 = 1
                        _ = v40
                    }
                    _ = v38
                }
                _ = fi32
                _ = v33
            }
            var v41 uint16 = v28
            v41 = v41
            _ = fi31
            _ = v41
        }
        v27 = "\n\t"
        v42 := GenS1{F0: (v28 / ((v28) | 1)), F1: false, F2: (0 <= 0)}
        _ = v42
    }
    _ = v27
    _ = v28
    return 1
}
func genFn1() uint {
    var v43 int16 = -1
    var v44 int64 = -464
    _ = v43
    _ = v44
    return 529
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 42
    var v3 int = 201
    var v4 uint = 2
    var v5 string = ("foo bar" + ("hello" + ("" + (("foo bar" + ("hello" + "\n\t")) + ("hello" + "")))))
    var v6 uint = v4
    for fi45 := 0; fi45 < 12; fi45++ {
        v4 += (^(v6))
        _ = fi45
    }
    v4 += v4
    v3 = -348
    switch v3 {
    case 204:
        if false {
            fmt.Println(v6, v2)
            if true {
                for fi46 := 0; fi46 < 1; fi46++ {
                    for fi47 := 0; fi47 < 17; fi47++ {
                        v3 = (fi46 & (v3 % (((v3 / (((fi46 / ((617) | 1))) | 1))) | 1)))
                        _ = fi47
                    }
                    var v48 uint8 = 255
                    v49 := []int{}
                    _ = fi46
                    _ = v48
                    _ = v49
                }
            } else {
                switch "hello" {
                case "":
                    if (!(-9223372036854775808 > -573)) {
                        if ((-452 >= -1) && true) {
                            v2 -= 1
                            if true {
                                v6 = genFn1()
                                v5 = ""
                                switch "hello" {
                                case "a":
                                    switch v3 {
                                    case 546:
                                        var v50 uint = (v4 * (v4 / ((42) | 1)))
                                        _ = v50
                                    case 127:
                                        for fi51 := 0; fi51 < 1; fi51++ {
                                            var v52 int64 = int64(v4)
                                            _ = fi51
                                            _ = v52
                                        }
                                    case -2147483648:
                                        delete(v1, ("a" + v5))
                                    default:
                                        var v53 bool = false
                                        _ = v53
                                    }
                                default:
                                    v54 := []int{}
                                    _ = v54
                                }
                            }
                            v2 += v2
                        }
                        if ((v3 - (v3 << ((v6) & 63))) >= v3) {
                            v5 = ""
                            var v55 float32 = 1e10
                            v56 := GenS4{F0: 75}
                            _ = v55
                            _ = v56
                        }
                    } else {
                        v57 := []int{v3, int(v2), -1, (v3 >> ((v4) & 63))}
                        switch (^(v3)) {
                        case -9223372036854775808:
                            v3 = (v3 ^ v3)
                        case -26:
                            if (true || false) {
                                var v58 int8 = 0
                                v59 := GenS2{F0: v5, F1: 939, F2: -1}
                                for fi60 := 0; fi60 < 7; fi60++ {
                                    for fi61 := 0; fi61 < 6; fi61++ {
                                        var v62 float64 = (-2.25 + (3.14159 / float64(1e10)))
                                        _ = fi61
                                        _ = v62
                                    }
                                    var v63 uint64 = 1
                                    _ = fi60
                                    _ = v63
                                }
                                _ = v58
                                _ = v59
                            } else {
                                var v64 int32 = -280
                                _ = v64
                            }
                        case -2147483648:
                            for fi65 := 0; fi65 < 9; fi65++ {
                                for fi66 := 0; fi66 < 16; fi66++ {
                                    if false {
                                        var v67 uint16 = 0
                                        switch 659 {
                                        case -2147483648:
                                            v68 := v1[v5]
                                            _ = v68
                                        case 346:
                                            v2 -= v2
                                        default:
                                            if false {
                                                if true {
                                                    v67 += 174
                                                    fmt.Println(v4, v4, v6)
                                                } else {
                                                    if true {
                                                        switch -54 {
                                                        case -429:
                                                            fmt.Println(v3, v6, v4)
                                                        case 379:
                                                            v69 := GenS3{F0: 18446744073709551615, F1: 18446744073709551615, F2: 127}
                                                            _ = v69
                                                        case 2147483647:
                                                            v70 := GenS5{F0: 2147483647, F1: 2, F2: (v2 % (((v2 ^ 1)) | 1))}
                                                            _ = v70
                                                        default:
                                                            v4 = 18446744073709551615
                                                        }
                                                    }
                                                }
                                                var v71 int32 = 0
                                                _ = v71
                                            } else {
                                                if false {
                                                    fmt.Println(v5)
                                                    fmt.Println(v3, fi65)
                                                } else {
                                                    if false {
                                                        fmt.Println(fi65, v67)
                                                    } else {
                                                        v72 := GenS6{F0: false, F1: v5, F2: 9223372036854775807}
                                                        var v73 uint32 = uint32(v4)
                                                        switch fi65 {
                                                        case 9223372036854775807:
                                                            if false {
                                                                v74 := GenS7{F0: 42, F1: -41}
                                                                if ((!(!false)) || false) {
                                                                    var v75 uint32 = (v73 % ((v73) | 1))
                                                                    delete(v1, ((("" + ((v5 + "foo bar") + ("" + v5))) + v5) + v5))
                                                                    _ = v75
                                                                }
                                                                _ = v74
                                                            }
                                                        case 0:
                                                            fmt.Println(v6)
                                                        case -479:
                                                            switch 962 {
                                                            case 653:
                                                                switch v5 {
                                                                case "":
                                                                    if len(v57) > 0 {
                                                                        v57[0] = (fi66 << (((v6 & (v6 >> (((v4 * v6)) & 63)))) & 63))
                                                                    }
                                                                default:
                                                                    v67 += v67
                                                                }
                                                            case 32768:
                                                                if false {
                                                                    if true {
                                                                        v73 = uint32(v6)
                                                                        var v76 string = (v5 + "")
                                                                        _ = v76
                                                                    }
                                                                } else {
                                                                    var v77 uint = genFn1()
                                                                    _ = len(v1)
                                                                    _ = v77
                                                                }
                                                            case 0:
                                                                var v78 int64 = 9223372036854775807
                                                                _ = v78
                                                            }
                                                        }
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                    var v79 float64 = (3.14159 - -2.25)
                                                    v1[(v5 + v5)] = v3
                                                    _ = v79
                                                }
                                            }
                                        }
                                        _ = v67
                                    } else {
                                        var v80 int = v3
                                        if (465 != 255) {
                                            switch v5 {
                                            case "hello":
                                                v4 += v4
                                            case "foo bar":
                                                switch 2147483647 {
                                                case 1:
                                                    var v81 int8 = 100
                                                    _ = v81
                                                case 0:
                                                    var v82 uint64 = 2
                                                    _ = v82
                                                case 32768:
                                                    var v83 float64 = 3.14159
                                                    _ = v83
                                                }
                                            }
                                            fmt.Println(v6, v6, fi65)
                                            v2 -= -32768
                                        }
                                        if false {
                                            var v84 int = fi65
                                            v84 += 4294967295
                                            _ = v84
                                        } else {
                                            v85 := GenS8{F0: 9223372036854775807, F1: 2, F2: 0}
                                            var v86 uint16 = 65535
                                            _ = v85
                                            _ = v86
                                        }
                                        _ = v80
                                    }
                                    if true {
                                        v4 += genFn1()
                                        v87 := GenS2{F0: (("" + (v5 + v5)) + "hello"), F1: 42, F2: 2147483647}
                                        _ = len(v1)
                                        _ = v87
                                    }
                                    v88 := GenS2{F0: v5, F1: 42, F2: 1}
                                    _ = fi66
                                    _ = v88
                                }
                                _ = fi65
                            }
                        default:
                            v57 = append(v57, v3)
                        }
                        _ = len(v57)
                        _ = v57
                    }
                case "hello":
                    v2 = v2
                }
                for fi89 := 0; fi89 < 17; fi89++ {
                    v6 += (v4 / (((v6 >> (((v6 * 18446744073709551615)) & 63))) | 1))
                    _ = fi89
                }
                fmt.Println(v2)
            }
        }
    default:
        var v90 float64 = 1e10
        _ = v90
    }
    for fi91 := 0; fi91 < 1; fi91++ {
        v4 = v6
        v3 = v3
        if true {
            for fi92 := 0; fi92 < 16; fi92++ {
                switch (("foo bar" + "foo bar") + "a") {
                case "\n\t":
                    v93 := GenS6{F0: true, F1: v5, F2: 0}
                    _ = v93
                case "":
                    if (0 == -367) {
                        switch v3 {
                        case 32767:
                            fmt.Println(v2, v3, v4)
                        }
                        if true {
                            if (!true) {
                                if (669 <= 623) {
                                    var v94 uint32 = 542
                                    var v95 float32 = -2.25
                                    _ = v94
                                    _ = v95
                                } else {
                                    v4 -= (v6 % ((v4) | 1))
                                    v96 := []int{fi91, int(v6), v3}
                                    _ = v96
                                }
                                fmt.Println(v5)
                                var v97 int16 = v2
                                _ = v97
                            }
                            v98 := GenS9{F0: 9223372036854775807, F1: 0}
                            switch "\n\t" {
                            case "a":
                                fmt.Println(fi92, v3)
                            case "":
                                var v99 int16 = v2
                                _ = v99
                            default:
                                v6 -= (^(v4))
                            }
                            _ = v98
                        } else {
                            var v100 float32 = 0.0
                            var v101 uint64 = 2
                            var v102 uint16 = 989
                            _ = v100
                            _ = v101
                            _ = v102
                        }
                        v2 += v2
                    } else {
                        if false {
                            if (3.14159 >= (3.14159 * (0.1 - 0.1))) {
                                v103 := []int{(fi92 % (((v3 + v3)) | 1)), fi92, fi91}
                                v3 = v3
                                if (true || false) {
                                    v104 := GenS8{F0: 495, F1: 0, F2: 527}
                                    _ = v104
                                }
                                _ = v103
                            } else {
                                if ((v4 * 42) == (v6 + (^(v6)))) {
                                    if true {
                                        fmt.Println(v5, fi92, v6)
                                        v6 = v6
                                    } else {
                                        var v105 float64 = (0.1 / (float64(float32(-2.25)) / float64(((((0.0 - (((0.1 * -2.25) * float32(1.5)) + 3.14159)) / 0.1) / (1e10 / 3.14159)) / 1.0))))
                                        switch (^(fi92)) {
                                        case -1:
                                            var v106 string = ("a" + "a")
                                            _ = v106
                                        case 32768:
                                            if false {
                                                v107 := GenS10{F0: v2, F1: v2, F2: -1}
                                                if false {
                                                    v105 = float64(1e10)
                                                    var v108 float32 = (float32((0.1 - v105)) + 0.1)
                                                    _ = v108
                                                } else {
                                                    var v109 uint8 = 117
                                                    fmt.Println(v105)
                                                    _ = v109
                                                }
                                                v110 := GenS11{F0: 11, F1: (^(v4)), F2: 0}
                                                v105 -= -2.25
                                                _ = len(v1)
                                                _ = v107
                                                _ = v110
                                            } else {
                                                var v111 int16 = -1
                                                _ = v111
                                            }
                                        case 127:
                                            var v112 uint8 = 42
                                            _ = v112
                                        }
                                        _ = v105
                                    }
                                    var v113 uint8 = 220
                                    _ = v113
                                }
                                var v114 int64 = 42
                                switch "a" {
                                case "a":
                                    var v115 int64 = v114
                                    _ = v115
                                default:
                                    fmt.Println(v114, fi91)
                                }
                                _ = v114
                            }
                            v116 := GenS2{F0: (v5 + v5), F1: -1, F2: -2147483648}
                            var v117 uint64 = 2
                            _ = v116
                            _ = v117
                        } else {
                            v118 := GenS7{F0: 660, F1: 115}
                            var v119 int32 = -1
                            _ = v118
                            _ = v119
                        }
                    }
                default:
                    v6 = (^(v4))
                }
                var v120 string = v5
                _ = fi92
                _ = v120
            }
            v121 := GenS5{F0: 4294967295, F1: 374, F2: v2}
            _ = v121
        } else {
            var v122 bool = ((-9223372036854775808 > 1) && (1 >= 400))
            _ = v122
        }
        _ = fi91
    }
    switch ((("a" + "foo bar") + (v5 + (((("hello" + (v5 + "foo bar")) + (v5 + v5)) + "") + (("hello" + (v5 + v5)) + v5)))) + "\n\t") {
    case "":
        var v123 int32 = 1
        _ = v123
    default:
        var v124 bool = true
        _ = v124
    }
    if false {
        if false {
            var v125 bool = false
            _ = v125
        } else {
            var v126 uint32 = 42
            switch "" {
            case "hello":
                v6 = 42
            default:
                v2 = 807
            }
            _ = v126
        }
    } else {
        if false {
            v127 := GenS10{F0: (v2 >> (((v6 >> ((v4) & 63))) & 63)), F1: (v2 + (v2 - (v2 + (v2 + (v2 - (-(v2))))))), F2: -685}
            switch ((v5 + "") + v5) {
            case "foo bar":
                var v128 string = "hello"
                _ = v128
            case "a":
                fmt.Println(v6, v4, v2)
            case "":
                switch "hello" {
                case "a":
                    var v129 int16 = v2
                    _ = v129
                case "hello":
                    var v130 string = v5
                    _ = v130
                }
            }
            _ = v127
        } else {
            var v131 float32 = (((3.14159 * (1e10 + (0.1 / ((-2.25 - 3.14159) - 1e10)))) / ((1e10 - -2.25) + 0.1)) * (((((0.1 - (0.0 + 1.5)) / 1.0) * ((0.0 - 0.1) + 1.5)) / 3.14159) / 0.1))
            var v132 float32 = v131
            v5 = v5
            _ = v131
            _ = v132
        }
        if true {
            var v133 float32 = -2.25
            for fi134 := 0; fi134 < 10; fi134++ {
                v3 = v3
                var v135 uint64 = 4294967295
                v2 = v2
                _ = fi134
                _ = v135
            }
            var v136 int16 = (v2 ^ v2)
            _ = v133
            _ = v136
        } else {
            var v137 int64 = int64(v2)
            v138 := []int{v3, v3, v3, v3}
            if (((v5 + v5) > "\n\t") || true) {
                v5 = "foo bar"
                fmt.Println(v5)
            } else {
                var v139 bool = (!(false && ("\n\t" >= v5)))
                v140 := v1[""]
                var v141 int16 = v2
                _ = v139
                _ = v140
                _ = v141
            }
            _ = v137
            _ = v138
        }
        v142 := GenS10{F0: v2, F1: (v2 % ((v2) | 1)), F2: 784}
        _ = v142
    }
    v143 := v1[v5]
    for fi144 := 0; fi144 < 5; fi144++ {
        v2 = (v2 & v2)
        for fi145 := 0; fi145 < 6; fi145++ {
            v146 := GenS5{F0: 2, F1: 826, F2: (^(v2))}
            v147 := GenS9{F0: -1, F1: 2}
            var v148 float32 = float32(0.1)
            _ = len(v1)
            _ = fi145
            _ = v146
            _ = v147
            _ = v148
        }
        _ = fi144
    }
    switch (v3 + (-(v3))) {
    case 2147483647:
        v149 := GenS12{F0: v143, F1: uint8(v6), F2: (v6 * (v6 / ((v6) | 1)))}
        _ = v149
    default:
        if false {
            v150 := GenS13{F0: 0, F1: -120}
            _ = v150
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v143:", v143)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v143
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
