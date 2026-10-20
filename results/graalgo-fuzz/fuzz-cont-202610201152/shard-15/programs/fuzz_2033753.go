package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint; F2 int32 }
type GenS5 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 uint16; F1 int; F2 uint8 }
type GenS8 struct { F0 uint; F1 uint8; F2 int }
func genFn0() uint32 {
    var v5 int16 = 1
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 uint32 = 1
        _ = fi6
        _ = v7
    }
    v5 = 42
    v8 := []int{86, -440, 127}
    _ = v5
    _ = v8
    return 696
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "foo bar"
    var v3 int = 0
    var v4 uint64 = 1
    var v9 uint32 = 1
    v10 := GenS1{F0: -128, F1: genFn0()}
    switch (v2 + v2) {
    case "foo bar":
        _ = len(v1)
    case "hello":
        v9 += genFn0()
    }
    fmt.Println(v4, v4)
    var v11 float64 = (-2.25 + -2.25)
    if false {
        fmt.Println(v3)
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            switch "a" {
            case "a":
                switch v2 {
                case "a":
                    switch fi12 {
                    case 2147483647:
                        if false {
                            v4 = (v4 * 18446744073709551615)
                        }
                        v9 = v9
                        _ = len(v1)
                    default:
                        if false {
                            if (818 > 0) {
                                var v14 uint = 65535
                                v3 += v3
                                v11 += (v11 + v11)
                                _ = v14
                            } else {
                                v15 := GenS2{F0: -9223372036854775808, F1: 1}
                                v16 := []int{fi12, fi12, 32767}
                                _ = v15
                                _ = v16
                            }
                        }
                        v17 := GenS1{F0: -72, F1: (v9 / ((v9) | 1))}
                        if (598 == 255) {
                            if false {
                                v3 = (fi12 | 329)
                                v9 = 2
                                var v18 uint16 = 65535
                                _ = v18
                            }
                            switch "\n\t" {
                            case "":
                                v19 := []int{}
                                _ = v19
                            case "hello":
                                if true {
                                    if (!true) {
                                        v9 = (v9 ^ v9)
                                        v3 += (^(fi12))
                                    } else {
                                        var v20 int8 = 127
                                        switch (v3 % (((fi13 + (fi12 ^ fi12))) | 1)) {
                                        case 749:
                                            v9 = v9
                                        default:
                                            if false {
                                                v21 := GenS1{F0: v20, F1: v9}
                                                fmt.Println(v2, v11)
                                                _ = v21
                                            } else {
                                                var v22 float64 = v11
                                                _ = v22
                                            }
                                            v23 := []int{v3, v3}
                                            fmt.Println(fi12, v11, fi13)
                                            _ = v23
                                        }
                                        if true {
                                            fmt.Println(fi12, v20, v9)
                                        }
                                        if (0 >= 1) {
                                            if (0 == 18446744073709551615) {
                                                var v24 uint32 = v9
                                                if (true && false) {
                                                    fmt.Println(v2, v4)
                                                } else {
                                                    v25 := GenS2{F0: 0, F1: 9}
                                                    v2 = v2
                                                    v9 = 876
                                                    var v26 string = ("\n\t" + v2)
                                                    v26 = "a"
                                                    fmt.Println(v24)
                                                    if (1 > v4) {
                                                        v27 := GenS3{F0: 160}
                                                        _ = v27
                                                    }
                                                    if false {
                                                        v28 := GenS3{F0: 0}
                                                        var v29 int32 = int32(v20)
                                                        fmt.Println(fi13)
                                                        _ = v28
                                                        _ = v29
                                                    } else {
                                                        v3 += fi12
                                                        fmt.Println(v26)
                                                        v24 = 42
                                                    }
                                                    fmt.Println(fi12, v4, v4)
                                                    _ = v25
                                                    _ = v26
                                                }
                                                _ = v24
                                            }
                                            v20 = (v20 & v20)
                                        }
                                        switch (v3 & (fi13 + fi12)) {
                                        case 127:
                                            var v30 int = -2147483648
                                            _ = v30
                                        case 32767:
                                            v31 := GenS2{F0: 1, F1: 1000}
                                            _ = v31
                                        case 32768:
                                            v32 := GenS2{F0: -9223372036854775808, F1: 1}
                                            _ = v32
                                        default:
                                            var v33 float32 = (0.0 / 1.5)
                                            _ = v33
                                        }
                                        if true {
                                            v34 := GenS4{F0: 1, F1: 65535, F2: int32(v20)}
                                            switch v2 {
                                            case "a":
                                                if true {
                                                    v35 := []int{v3, fi13, fi13, (fi13 * fi13)}
                                                    if false {
                                                        var v36 int = (^(v3))
                                                        var v37 uint64 = (v4 ^ (v4 / (((v4 | v4)) | 1)))
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                    v38 := GenS5{F0: v4, F1: 0, F2: 42}
                                                    var v39 int64 = 42
                                                    v40 := GenS1{F0: (v20 ^ 100), F1: 4294967295}
                                                    _ = v35
                                                    _ = v38
                                                    _ = v39
                                                    _ = v40
                                                } else {
                                                    v3 -= fi12
                                                    var v41 uint8 = 2
                                                    _ = v41
                                                }
                                            case "":
                                                fmt.Println(fi12, v4, fi13)
                                            default:
                                                var v42 bool = ((32767 != uint16(v9)) && (2 < 2))
                                                _ = v42
                                            }
                                            if true {
                                                var v43 uint64 = v4
                                                if false {
                                                    var v44 float32 = -2.25
                                                    var v45 int32 = -2147483648
                                                    _ = v44
                                                    _ = v45
                                                } else {
                                                    switch "" {
                                                    case "\n\t":
                                                        if false {
                                                            _ = len(v1)
                                                            delete(v1, v9)
                                                            v20 -= v20
                                                        } else {
                                                            switch 101 {
                                                            case 32767:
                                                                v4 += 4294967295
                                                            case 9223372036854775807:
                                                                v2 = v2
                                                            case 4294967295:
                                                                if (!(28 == int16(v9))) {
                                                                    v4 += 0
                                                                } else {
                                                                    v43 = v4
                                                                    v46 := GenS5{F0: (v43 + v43), F1: 0, F2: 0}
                                                                    v4 = (v4 % ((18446744073709551615) | 1))
                                                                    _ = v46
                                                                }
                                                            }
                                                            fmt.Println(v2, v2)
                                                        }
                                                    default:
                                                        v43 = v43
                                                    }
                                                }
                                                switch fi12 {
                                                case -2147483648:
                                                    v43 += (^(v4))
                                                default:
                                                    var v47 int = v3
                                                    switch v3 {
                                                    case 128:
                                                        v47 += fi13
                                                    default:
                                                        if true {
                                                            switch (("a" + (((v2 + (v2 + (((v2 + ("\n\t" + (v2 + (v2 + "")))) + ((((("" + (v2 + ((("\n\t" + (("" + v2) + v2)) + v2) + v2))) + v2) + v2) + v2) + v2)) + "a"))) + v2) + (v2 + "\n\t"))) + "hello") {
                                                            case "":
                                                                if true {
                                                                    var v48 int = (fi12 - -2147483648)
                                                                    _ = v48
                                                                }
                                                            }
                                                        }
                                                    }
                                                    if true {
                                                        v49 := GenS6{F0: 42, F1: 136}
                                                        var v50 int16 = 1
                                                        if (false && true) {
                                                            delete(v1, 169)
                                                        } else {
                                                            if true {
                                                                if (4294967295 != 893) {
                                                                    fmt.Println(v2)
                                                                    if (((v2 + (("hello" + v2) + ("hello" + v2))) >= "foo bar") && (18446744073709551615 > 1)) {
                                                                        v1[genFn0()] = "a"
                                                                    } else {
                                                                        var v51 bool = false
                                                                        _ = v51
                                                                    }
                                                                } else {
                                                                    v47 = (-(fi12))
                                                                    var v52 float32 = (1.5 / -2.25)
                                                                    _ = v52
                                                                }
                                                                var v53 int16 = (-(v50))
                                                                v54 := []int{}
                                                                _ = v53
                                                                _ = v54
                                                            } else {
                                                                v55 := []int{42, fi12, (^(fi12)), v47}
                                                                _ = v55
                                                            }
                                                        }
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                    if (uint(v9) >= 65535) {
                                                        v3 += fi12
                                                        switch v2 {
                                                        case "hello":
                                                            var v56 int8 = 100
                                                            _ = v56
                                                        case "":
                                                            var v57 int8 = (v20 ^ -46)
                                                            _ = v57
                                                        case "foo bar":
                                                            v1[(v9 % ((v9) | 1))] = (v2 + ((v2 + ("" + ("\n\t" + ("" + "hello")))) + v2))
                                                        }
                                                        switch v3 {
                                                        case 0:
                                                            delete(v1, (v9 / (((v9 ^ 65535)) | 1)))
                                                        }
                                                    }
                                                    if true {
                                                        v58 := []int{(-(v3)), fi12}
                                                        var v59 float64 = 1e10
                                                        if false {
                                                            var v60 uint32 = (v9 & (v9 + 77))
                                                            v4 += v4
                                                            _ = v60
                                                        } else {
                                                            if ((^(v9)) != (v9 * 21)) {
                                                                var v61 uint = 835
                                                                _ = v61
                                                            }
                                                            v62 := GenS7{F0: 2, F1: -33, F2: 1}
                                                            if (true || true) {
                                                                fmt.Println(v47, v11)
                                                                v58 = append(v58, -337)
                                                                fmt.Println(v20, v9, v2)
                                                            } else {
                                                                var v63 int32 = -1
                                                                _ = v63
                                                            }
                                                            _ = v62
                                                        }
                                                        _ = v58
                                                        _ = v59
                                                    } else {
                                                        v2 = v2
                                                    }
                                                    _ = v47
                                                }
                                                _ = v43
                                            } else {
                                                if false {
                                                    v11 -= float64(((3.14159 + 0.1) + 0.1))
                                                    switch v3 {
                                                    case 2147483647:
                                                        v64 := []int{}
                                                        _ = v64
                                                    case -225:
                                                        if false {
                                                            if true {
                                                                switch v2 {
                                                                case "":
                                                                    var v65 float64 = (v11 - 0.0)
                                                                    _ = v65
                                                                case "foo bar":
                                                                    var v66 int = fi12
                                                                    _ = v66
                                                                case "a":
                                                                    if (2 == 2147483648) {
                                                                        var v67 string = (v2 + "a")
                                                                        _ = len(v1)
                                                                        v20 -= 92
                                                                        _ = v67
                                                                    } else {
                                                                        if true {
                                                                            v1[v9] = ""
                                                                        } else {
                                                                            v68 := GenS4{F0: 42, F1: 2, F2: 0}
                                                                            var v69 uint8 = 128
                                                                            _ = v68
                                                                            _ = v69
                                                                        }
                                                                    }
                                                                default:
                                                                    if false {
                                                                        v9 += v9
                                                                    }
                                                                }
                                                                switch v2 {
                                                                case "":
                                                                    switch ((v2 + v2) + (v2 + (v2 + v2))) {
                                                                    case "\n\t":
                                                                        v70 := []int{(fi13 ^ -1), (^(fi13)), (fi12 * (fi12 * (-(fi12))))}
                                                                        _ = v70
                                                                    }
                                                                case "a":
                                                                    v71 := GenS7{F0: 0, F1: (^(fi12)), F2: 255}
                                                                    _ = v71
                                                                case "hello":
                                                                    var v72 string = v2
                                                                    _ = v72
                                                                }
                                                                v20 -= v20
                                                            }
                                                            var v73 int8 = (v20 + 0)
                                                            var v74 uint32 = genFn0()
                                                            _ = v73
                                                            _ = v74
                                                        }
                                                    case 733:
                                                        v11 -= -2.25
                                                    }
                                                }
                                                switch v2 {
                                                case "":
                                                    if (0.0 > 0.0) {
                                                        if true {
                                                            switch 625 {
                                                            case 32767:
                                                                v75 := GenS4{F0: 0, F1: 42, F2: -86}
                                                                _ = v75
                                                            case -127:
                                                                var v76 int = (fi12 ^ int(v9))
                                                                _ = v76
                                                            }
                                                            var v77 int64 = 9223372036854775807
                                                            fmt.Println(v9, v11)
                                                            v78 := []int{fi13}
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                        v20 = 84
                                                        v79 := GenS5{F0: v4, F1: 210, F2: uint8(v4)}
                                                        v80 := GenS2{F0: 9223372036854775807, F1: 1000}
                                                        _ = v79
                                                        _ = v80
                                                    } else {
                                                        var v81 int64 = -1
                                                        if (!true) {
                                                            var v82 int32 = 531
                                                            if (0.0 < v11) {
                                                                if false {
                                                                    var v83 uint32 = 4294967295
                                                                    v3 = fi13
                                                                    if true {
                                                                        if true {
                                                                            v3 = (v3 >> ((255) & 63))
                                                                            v84 := GenS8{F0: 2, F1: 2, F2: -1}
                                                                            _ = v84
                                                                        } else {
                                                                            v85 := v1[(v9 & 2)]
                                                                            _ = v85
                                                                        }
                                                                    }
                                                                    _ = v83
                                                                }
                                                                switch "a" {
                                                                case "\n\t":
                                                                    v20 = (v20 | 127)
                                                                }
                                                            } else {
                                                                var v86 uint8 = 14
                                                                v87 := []int{}
                                                                _ = v86
                                                                _ = v87
                                                            }
                                                            _ = v82
                                                        }
                                                        v88 := GenS4{F0: 0, F1: 2, F2: 1}
                                                        _ = v81
                                                        _ = v88
                                                    }
                                                case "\n\t":
                                                    switch v2 {
                                                    case "a":
                                                        v9 = v9
                                                    case "hello":
                                                        if true {
                                                            v89 := GenS3{F0: 42}
                                                            _ = v89
                                                        } else {
                                                            if (-1 != 2147483647) {
                                                                switch fi12 {
                                                                case 1:
                                                                    var v90 uint = 1
                                                                    _ = v90
                                                                case 128:
                                                                    v91 := GenS1{F0: v20, F1: (v9 / ((v9) | 1))}
                                                                    var v92 uint = uint(v9)
                                                                    _ = v91
                                                                    _ = v92
                                                                case 127:
                                                                    var v93 int8 = (v20 | (-(v20)))
                                                                    _ = v93
                                                                default:
                                                                    v94 := GenS4{F0: 1, F1: 2, F2: -196}
                                                                    _ = v94
                                                                }
                                                            } else {
                                                                if (-2147483648 != 2147483647) {
                                                                    _ = len(v1)
                                                                    v11 -= v11
                                                                    if false {
                                                                        if (((((0.1 * float32(float64((3.14159 + (3.14159 / 1e10))))) - -2.25) * -2.25) / 1e10) >= float32(((v11 + v11) + ((v11 + 1.5) - 0.0)))) {
                                                                            if (65535 <= 2) {
                                                                                v95 := v1[(^(v9))]
                                                                                v1[genFn0()] = ("a" + "\n\t")
                                                                                var v96 float64 = v11
                                                                                _ = v95
                                                                                _ = v96
                                                                            } else {
                                                                                fmt.Println(v2)
                                                                                var v97 int64 = -1
                                                                                _ = v97
                                                                            }
                                                                        }
                                                                        var v98 float32 = -2.25
                                                                        _ = v98
                                                                    } else {
                                                                        v99 := []int{fi12}
                                                                        var v100 bool = (!false)
                                                                        v11 = float64(float32((((0.0 / (v11 - ((v11 / ((float64(((1.5 - 1.5) - float32(1e10))) + ((v11 - ((1e10 / v11) + (((v11 - float64(0.1)) - ((((1.5 * 1.5) / (v11 / v11)) * float64(1e10)) * v11)) - (3.14159 / (float64((0.0 * 0.1)) + v11))))) - (1.5 - float64(0.1)))) / v11)) - 1.5))) - v11) * v11)))
                                                                        _ = v100
                                                                        _ = v99
                                                                    }
                                                                }
                                                                if true {
                                                                    var v101 string = ""
                                                                    _ = v101
                                                                }
                                                                fmt.Println(v20, v2, v3)
                                                                if false {
                                                                    if true {
                                                                        var v102 int64 = -9223372036854775808
                                                                        v103 := GenS4{F0: 321, F1: 255, F2: 0}
                                                                        switch ("a" + v2) {
                                                                        case "foo bar":
                                                                            v20 += v20
                                                                        default:
                                                                            v4 = 784
                                                                        }
                                                                        v104 := []int{(v3 >> ((2) & 63))}
                                                                        _ = v102
                                                                        _ = v103
                                                                        _ = v104
                                                                    } else {
                                                                        fmt.Println(v20)
                                                                        var v105 uint = 4294967295
                                                                        _ = v105
                                                                    }
                                                                    v2 = ("\n\t" + v2)
                                                                } else {
                                                                    var v106 uint = 42
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        v107 := GenS2{F0: -1, F1: -108}
                                                                        _ = v107
                                                                    case "hello":
                                                                        v108 := []int{(fi13 << ((v106) & 63))}
                                                                        _ = v108
                                                                    case "foo bar":
                                                                        switch v3 {
                                                                        case -453:
                                                                            v109 := GenS4{F0: 0, F1: v106, F2: 904}
                                                                            _ = v109
                                                                        default:
                                                                            if (!(v2 <= "")) {
                                                                                if (false || (-470 < 42)) {
                                                                                    var v110 string = v2
                                                                                    _ = v110
                                                                                }
                                                                                var v111 float32 = float32(v11)
                                                                                fmt.Println(v9, fi13, v11)
                                                                                _ = v111
                                                                            } else {
                                                                                switch -2147483648 {
                                                                                case 490:
                                                                                    var v112 string = v2
                                                                                    _ = v112
                                                                                }
                                                                                var v113 int = v3
                                                                                fmt.Println(v9, fi13)
                                                                                _ = v113
                                                                            }
                                                                        }
                                                                    default:
                                                                        v4 = (v4 & (v4 | (^(v4))))
                                                                    }
                                                                    _ = v106
                                                                }
                                                                if false {
                                                                    var v114 int = (-(fi12))
                                                                    v3 += 32767
                                                                    _ = v114
                                                                } else {
                                                                    v115 := GenS8{F0: 87, F1: 219, F2: 1}
                                                                    fmt.Println(v2, v9)
                                                                    _ = v115
                                                                }
                                                            }
                                                            v116 := GenS1{F0: v20, F1: 392}
                                                            _ = v116
                                                        }
                                                    case "foo bar":
                                                        if (0 > -1) {
                                                            v2 = (v2 + (v2 + v2))
                                                            delete(v1, v9)
                                                        } else {
                                                            fmt.Println(v20, v9)
                                                            v20 = -128
                                                            v117 := []int{}
                                                            v9 += v9
                                                            var v118 bool = true
                                                            _ = v117
                                                            _ = v118
                                                        }
                                                    default:
                                                        fmt.Println(v2, v20, v9)
                                                    }
                                                }
                                                v119 := GenS4{F0: 49, F1: uint(v4), F2: 0}
                                                _ = v119
                                            }
                                            _ = v34
                                        }
                                        _ = v20
                                    }
                                    var v120 bool = true
                                    var v121 int32 = -10
                                    _ = v120
                                    _ = v121
                                } else {
                                    if (0 != 336) {
                                        v2 = "hello"
                                    } else {
                                        v9 -= v9
                                        var v122 uint = 176
                                        var v123 float64 = v11
                                        _ = v122
                                        _ = v123
                                    }
                                    if false {
                                        switch -1 {
                                        case 32768:
                                            v124 := GenS3{F0: 1}
                                            _ = v124
                                        case 1:
                                            var v125 float32 = ((1.5 / 1.5) * 0.1)
                                            _ = v125
                                        case -2147483648:
                                            v4 = 18446744073709551615
                                        }
                                        v126 := []int{(fi13 ^ (fi12 % (((fi12 - int(v4))) | 1))), (v3 * (^(fi12))), fi12}
                                        var v127 int = (fi13 << ((2) & 63))
                                        v128 := GenS2{F0: 830, F1: 32767}
                                        v129 := GenS2{F0: 9223372036854775807, F1: 0}
                                        _ = v126
                                        _ = v127
                                        _ = v128
                                        _ = v129
                                    }
                                    v130 := []int{v3, 815}
                                    _ = v130
                                }
                            }
                        }
                        _ = v17
                    }
                }
            case "foo bar":
                v131 := []int{v3}
                _ = v131
            case "hello":
                var v132 uint8 = 255
                _ = v132
            }
            _ = len(v1)
            var v133 int32 = -520
            _ = fi13
            _ = v133
        }
        var v134 int8 = -1
        var v135 int64 = -1
        _ = fi12
        _ = v134
        _ = v135
    }
    if (genFn0() <= 65535) {
        for fi136 := 0; fi136 < 13; fi136++ {
            switch (v3 << ((18446744073709551615) & 63)) {
            case 128:
                switch "hello" {
                case "":
                    v9 = 4294967295
                case "a":
                    v4 = v4
                case "\n\t":
                    for fi137 := 0; fi137 < 4; fi137++ {
                        fmt.Println(fi136, fi136, fi136)
                        _ = fi137
                    }
                }
            case 853:
                if false {
                    var v138 string = (v2 + ("" + ((v2 + ("" + "a")) + v2)))
                    v139 := GenS1{F0: -128, F1: (v9 ^ (v9 >> ((256) & 63)))}
                    _ = v138
                    _ = v139
                }
            case 0:
                v140 := GenS2{F0: -1, F1: -32768}
                _ = v140
            }
            _ = fi136
        }
        var v141 int64 = int64(v3)
        _ = v141
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
