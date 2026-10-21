package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16; F1 int8; F2 int16 }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 int16; F1 int8; F2 uint16 }
type GenS9 struct { F0 uint16; F1 int; F2 uint32 }
type GenS10 struct { F0 int8; F1 uint32 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int64; F1 uint16; F2 string }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 string }

func main() {
    var v1 uint16 = 0
    var v2 uint64 = 864
    v3 := make(map[int]string)
    var v4 uint8 = 2
    v5 := []int{70, -368}
    for fi6 := 0; fi6 < 14; fi6++ {
        _ = len(v5)
        v7 := GenS1{F0: "hello", F1: 1, F2: (v1 < 42)}
        var v8 int = (fi6 ^ fi6)
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 string = "hello"
    v1 = v1
    var v10 uint32 = 2147483647
    v2 -= (v2 ^ v2)
    v5 = append(v5, 32767)
    var v11 uint8 = (v4 ^ v4)
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 float32 = 3.14159
        if false {
            if ((^(v4)) < v4) {
                if (!true) {
                    var v14 uint = 255
                    if (true || (127 >= 100)) {
                        v14 = v14
                        delete(v3, (fi12 - fi12))
                        if ((false && true) && (v9 == ((((("" + ("\n\t" + v9)) + v9) + ("foo bar" + (v9 + ""))) + (("foo bar" + "hello") + "\n\t")) + (("foo bar" + "\n\t") + (v9 + v9))))) {
                            var v15 uint32 = (v10 + (v10 % ((42) | 1)))
                            v4 = (^(v11))
                            _ = v15
                        } else {
                            for fi16 := 0; fi16 < 18; fi16++ {
                                var v17 int8 = -1
                                if true {
                                    var v18 int = -456
                                    fmt.Println(v13, v9)
                                    if true {
                                        v9 = v9
                                        switch (fi12 / ((32767) | 1)) {
                                        case 32767:
                                            var v19 int8 = (v17 ^ v17)
                                            _ = v19
                                        case 582:
                                            v20 := GenS1{F0: v9, F1: v17, F2: (42 < -243)}
                                            _ = v20
                                        default:
                                            fmt.Println(v17)
                                            v21 := GenS2{F0: (v10 ^ (v10 * 0))}
                                            _ = v21
                                        }
                                    }
                                    _ = v18
                                } else {
                                    v3[fi12] = v9
                                    if (!false) {
                                        v10 = 373
                                        var v22 float64 = 1e10
                                        v14 = 0
                                        v4 = v4
                                        v10 = v10
                                        v23 := GenS3{F0: (v10 * v10), F1: 42}
                                        v24 := GenS3{F0: 176, F1: 901}
                                        v17 += v17
                                        var v25 bool = false
                                        v26 := GenS4{F0: v9}
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    }
                                }
                                var v27 int16 = 0
                                _ = fi16
                                _ = v17
                                _ = v27
                            }
                            var v28 int32 = -1
                            for fi29 := 0; fi29 < 13; fi29++ {
                                v11 = 128
                                if len(v5) > 0 {
                                    v5[0] = fi29
                                }
                                v14 += (v14 & v14)
                                _ = fi29
                            }
                            _ = v28
                        }
                    } else {
                        v13 = 0.1
                        for fi30 := 0; fi30 < 12; fi30++ {
                            if false {
                                fmt.Println(v13)
                                if len(v5) > 0 {
                                    v5[0] = -1
                                }
                            } else {
                                fmt.Println(v2, v9, v10)
                                if true {
                                    switch v9 {
                                    case "hello":
                                        if (v14 < (^(v14))) {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            switch v9 {
                                            case "":
                                                var v31 uint16 = v1
                                                _ = v31
                                            case "a":
                                                if true {
                                                    v4 = (^(v11))
                                                } else {
                                                    v32 := GenS5{F0: ((v11 << ((v14) & 63)) <= 0)}
                                                    var v33 int64 = -499
                                                    var v34 int = (fi30 | (-(fi30)))
                                                    v35 := GenS2{F0: v10}
                                                    _ = v32
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                }
                                            }
                                            v4 += v11
                                            if true {
                                                v36 := GenS6{F0: (v1 | (^(v1))), F1: 42, F2: 1}
                                                _ = len(v3)
                                                v37 := GenS5{F0: true}
                                                _ = v36
                                                _ = v37
                                            } else {
                                                var v38 int16 = int16(v2)
                                                v39 := GenS7{F0: 1, F1: 2}
                                                _ = v38
                                                _ = v39
                                            }
                                            v4 = v11
                                            if true {
                                                var v40 bool = (-2147483648 > 2147483647)
                                                v9 = ""
                                                _ = v40
                                            }
                                            _ = len(v5)
                                        } else {
                                            var v41 uint16 = 42
                                            _ = v41
                                        }
                                    case "foo bar":
                                        v13 -= v13
                                    case "":
                                        v42 := GenS1{F0: ("foo bar" + "foo bar"), F1: -128, F2: (uint64(v11) >= v2)}
                                        _ = v42
                                    default:
                                        if true {
                                            v43 := GenS8{F0: 0, F1: 42, F2: (^(v1))}
                                            fmt.Println(fi30, v4, v13)
                                            _ = v43
                                        }
                                    }
                                    if ((v4 >> (((v14 >> (((v14 ^ (^(v14)))) & 63))) & 63)) == (v11 % (((^(v11))) | 1))) {
                                        fmt.Println(v2)
                                        switch (-(fi30)) {
                                        case 1:
                                            if false {
                                                if true {
                                                    v44 := GenS9{F0: (v1 & 32767), F1: fi30, F2: (v10 % ((2147483647) | 1))}
                                                    v14 = (v14 - 18446744073709551615)
                                                    var v45 int = fi12
                                                    _ = v44
                                                    _ = v45
                                                }
                                                v2 -= v2
                                                switch 9223372036854775807 {
                                                case 609:
                                                    v46 := v3[4294967295]
                                                    _ = v46
                                                }
                                                switch ("foo bar" + "hello") {
                                                case "hello":
                                                    v1 = v1
                                                default:
                                                    if false {
                                                        var v47 int8 = 127
                                                        v48 := GenS8{F0: -32768, F1: v47, F2: v1}
                                                        _ = v47
                                                        _ = v48
                                                    } else {
                                                        _ = len(v5)
                                                        v10 += uint32(v4)
                                                        v49 := v3[fi30]
                                                        _ = v49
                                                    }
                                                }
                                                fmt.Println(v9, v10, v10)
                                            }
                                        case 968:
                                            fmt.Println(v10, v13)
                                        default:
                                            switch -9223372036854775808 {
                                            case -2147483648:
                                                v50 := GenS10{F0: 15, F1: v10}
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                if false {
                                                    var v51 uint = (v14 * v14)
                                                    if false {
                                                        fmt.Println(v1, v14)
                                                        if (true && (1.5 <= 0.1)) {
                                                            var v52 int8 = 42
                                                            var v53 float32 = ((v13 + v13) / v13)
                                                            if false {
                                                                var v54 uint8 = (v4 >> (((v14 - 218)) & 63))
                                                                if (42 < -32768) {
                                                                    switch (v9 + "hello") {
                                                                    case "\n\t":
                                                                        v4 = (v4 >> ((v14) & 63))
                                                                    case "hello":
                                                                        var v55 int32 = 2147483647
                                                                        _ = v55
                                                                    case "foo bar":
                                                                        var v56 string = v9
                                                                        _ = v56
                                                                    default:
                                                                        if true {
                                                                            switch (^(fi12)) {
                                                                            case 1:
                                                                                v57 := GenS11{F0: (v2 % (((v2 << (((v51 ^ 40)) & 63))) | 1))}
                                                                                _ = v57
                                                                            default:
                                                                                v13 -= (1e10 / 1e10)
                                                                            }
                                                                            fmt.Println(v51)
                                                                            fmt.Println(v51, v2, v51)
                                                                        } else {
                                                                            v58 := GenS3{F0: (^(v10)), F1: 76}
                                                                            _ = v58
                                                                        }
                                                                    }
                                                                } else {
                                                                    v59 := GenS3{F0: v10, F1: 0}
                                                                    if false {
                                                                        v1 += 32767
                                                                        fmt.Println(v52, v11)
                                                                        var v60 string = v9
                                                                        _ = v60
                                                                    }
                                                                    _ = v59
                                                                }
                                                                _ = v54
                                                            } else {
                                                                if (42 < 0) {
                                                                    v53 -= 0.0
                                                                    if (true || ((v2 > v2) && false)) {
                                                                        v5 = append(v5, 164)
                                                                        var v61 int = 331
                                                                        _ = v61
                                                                    } else {
                                                                        v62 := GenS12{F0: -9223372036854775808, F1: uint16(v4), F2: v9}
                                                                        fmt.Println(v10)
                                                                        var v63 int = (fi12 / (((-(fi30))) | 1))
                                                                        fmt.Println(v14)
                                                                        var v64 uint16 = 390
                                                                        _ = v62
                                                                        _ = v63
                                                                        _ = v64
                                                                    }
                                                                    var v65 bool = (((!true) || true) && true)
                                                                    _ = v65
                                                                } else {
                                                                    if false {
                                                                        var v66 uint = (v51 << ((uint(v11)) & 63))
                                                                        _ = v66
                                                                    } else {
                                                                        var v67 string = v9
                                                                        v51 -= v14
                                                                        v9 = "hello"
                                                                        _ = v67
                                                                    }
                                                                }
                                                                if false {
                                                                    v53 = ((v53 - (1e10 + v53)) + (0.0 * v53))
                                                                    v51 -= 2
                                                                    var v68 float64 = 1.5
                                                                    _ = v68
                                                                }
                                                            }
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                        v5 = append(v5, (fi30 % ((fi30) | 1)))
                                                    }
                                                    _ = v51
                                                } else {
                                                    fmt.Println(v10)
                                                    if ((1e10 + ((3.14159 + ((((1e10 - (((1.5 / -2.25) * (1e10 - (1e10 + (-2.25 + 3.14159)))) / 1.5)) - (1.5 + 0.1)) * (0.1 * 1e10)) + -2.25)) * (-2.25 / 1.0))) > -2.25) {
                                                        if false {
                                                            if (false && true) {
                                                                var v69 bool = false
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    v9 = ("hello" + "a")
                                                                case "a":
                                                                    v69 = false
                                                                default:
                                                                    var v70 uint64 = 42
                                                                    _ = v70
                                                                }
                                                                switch (v9 + "") {
                                                                case "":
                                                                    var v71 uint64 = (v2 * v2)
                                                                    _ = v71
                                                                case "hello":
                                                                    if v69 {
                                                                        v69 = (100 == 18)
                                                                    } else {
                                                                        v72 := GenS13{F0: v11}
                                                                        _ = v72
                                                                    }
                                                                default:
                                                                    v9 = "hello"
                                                                }
                                                                _ = v69
                                                            }
                                                        }
                                                    } else {
                                                        var v73 bool = (false || ((true || (v9 > v9)) || false))
                                                        v74 := GenS14{F0: "hello"}
                                                        fmt.Println(v13)
                                                        v10 += (v10 | v10)
                                                        _ = v73
                                                        _ = v74
                                                    }
                                                    var v75 float64 = 3.14159
                                                    _ = v75
                                                }
                                                if (int32(v10) > 805) {
                                                    var v76 int64 = 9223372036854775807
                                                    _ = v76
                                                }
                                                _ = v50
                                            case -983:
                                                fmt.Println(v10)
                                            }
                                        }
                                        var v77 uint16 = 343
                                        var v78 float64 = 3.14159
                                        if true {
                                            if true {
                                                switch (-(fi30)) {
                                                case 32767:
                                                    if len(v5) > 0 {
                                                        v5[0] = (fi30 ^ fi12)
                                                    }
                                                case -585:
                                                    v11 = v4
                                                case 32768:
                                                    v13 += v13
                                                default:
                                                    v78 = (v78 + 0.1)
                                                }
                                                var v79 float64 = float64(float32(-2.25))
                                                switch fi30 {
                                                case 127:
                                                    v80 := GenS14{F0: ((v9 + "\n\t") + v9)}
                                                    _ = v80
                                                }
                                                _ = v79
                                            }
                                            switch fi30 {
                                            case 2147483647:
                                                var v81 bool = ((v14 - 2) <= v14)
                                                v11 -= uint8(v14)
                                                switch (fi30 + fi12) {
                                                case 420:
                                                    fmt.Println(v11, v13)
                                                    v13 += float32(v78)
                                                    if v81 {
                                                        if v81 {
                                                            var v82 bool = v81
                                                            _ = v82
                                                        } else {
                                                            v13 = ((v13 - 3.14159) - float32(v78))
                                                            var v83 uint32 = v10
                                                            if ((fi12 & (fi30 ^ (^(fi12)))) == fi30) {
                                                                fmt.Println(v2, v10)
                                                            }
                                                            _ = v83
                                                        }
                                                        var v84 int = (-(fi12))
                                                        _ = v84
                                                    }
                                                default:
                                                    v10 = 2147483648
                                                }
                                                _ = v81
                                            case 42:
                                                v85 := GenS3{F0: v10, F1: v14}
                                                _ = v85
                                            case 4294967295:
                                                v11 = (v11 * v11)
                                            default:
                                                v13 = ((v13 * (3.14159 * (v13 * float32(((v78 - v78) * v78))))) + 1.5)
                                            }
                                        }
                                        v3[(fi12 >> ((v14) & 63))] = "foo bar"
                                        _ = len(v5)
                                        _ = v77
                                        _ = v78
                                    }
                                    _ = len(v3)
                                } else {
                                    var v86 string = ((((("hello" + (v9 + "")) + v9) + v9) + ("a" + v9)) + "")
                                    switch ((v86 + (v86 + v86)) + ("hello" + (v86 + ((v86 + v9) + "foo bar")))) {
                                    case "\n\t":
                                        v11 += (^(v4))
                                    }
                                    v9 = v86
                                    _ = v86
                                }
                            }
                            switch ((v9 + "foo bar") + ("\n\t" + v9)) {
                            case "hello":
                                var v87 int8 = 127
                                _ = v87
                            case "":
                                if true {
                                    switch v9 {
                                    case "hello":
                                        v9 = "a"
                                    default:
                                        v88 := GenS8{F0: -66, F1: 127, F2: (v1 << ((v14) & 63))}
                                        _ = v88
                                    }
                                    switch v9 {
                                    case "foo bar":
                                        fmt.Println(v13, fi12)
                                    case "":
                                        if false {
                                            var v89 float32 = (0.1 * float32(3.14159))
                                            var v90 float64 = (1.5 * 3.14159)
                                            v9 = v9
                                            _ = v89
                                            _ = v90
                                        }
                                    case "a":
                                        var v91 float32 = ((1e10 / float32(-2.25)) - 0.0)
                                        v13 -= float32((3.14159 * 0.0))
                                        var v92 int32 = 653
                                        _ = v91
                                        _ = v92
                                    default:
                                        v4 = 42
                                    }
                                    fmt.Println(v4, v14)
                                    if true {
                                        v4 = (v11 % (((^(v11))) | 1))
                                    } else {
                                        fmt.Println(v9, fi12, v4)
                                    }
                                } else {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v93 string = ((v9 + "foo bar") + "hello")
                                    v1 += uint16(v4)
                                    _ = v93
                                }
                            case "foo bar":
                                switch fi30 {
                                case 127:
                                    v94 := GenS15{F0: -1}
                                    _ = v94
                                case -1:
                                    if false {
                                        fmt.Println(v9, v13)
                                        if (!true) {
                                            v10 = (^(v10))
                                        } else {
                                            var v95 uint64 = v2
                                            v96 := GenS13{F0: 0}
                                            v13 = 3.14159
                                            _ = v95
                                            _ = v96
                                        }
                                        _ = len(v3)
                                        v2 = uint64(v1)
                                        v10 = v10
                                        switch fi30 {
                                        case -2147483648:
                                            v11 -= 59
                                        default:
                                            v5 = append(v5, fi30)
                                        }
                                    } else {
                                        v97 := GenS11{F0: 1}
                                        _ = v97
                                    }
                                default:
                                    if len(v5) > 0 {
                                        v5[0] = fi30
                                    }
                                    fmt.Println(v2, v4, v9)
                                    var v98 float64 = (1e10 - (float64(0.0) + (0.1 - -2.25)))
                                    _ = v98
                                }
                            }
                            _ = fi30
                        }
                        v2 = 1
                    }
                    var v99 int = (^(fi12))
                    _ = v14
                    _ = v99
                }
                v3[fi12] = v9
                var v100 uint = 65535
                _ = v100
            } else {
                var v101 uint64 = v2
                fmt.Println(v4)
                var v102 uint16 = (v1 + (v1 & (v1 - v1)))
                _ = v101
                _ = v102
            }
            if ((v1 ^ v1) <= v1) {
                fmt.Println(v1, v4)
                var v103 uint32 = 820
                _ = v103
            } else {
                for fi104 := 0; fi104 < 8; fi104++ {
                    var v105 uint8 = uint8(v2)
                    v106 := v3[1]
                    v107 := GenS12{F0: 0, F1: v1, F2: v9}
                    _ = fi104
                    _ = v105
                    _ = v106
                    _ = v107
                }
            }
        } else {
            delete(v3, (fi12 * fi12))
            v9 = (v9 + "")
        }
        if (!(57 != 127)) {
            var v108 uint64 = 2
            v10 -= v10
            _ = v108
        } else {
            v109 := GenS8{F0: 1, F1: 42, F2: (v1 & v1)}
            _ = v109
        }
        _ = fi12
        _ = v13
    }
    v5 = append(v5, -909)
    var v110 uint8 = (v11 / ((uint8(v2)) | 1))
    var v111 int = 1
    v112 := GenS11{F0: v2}
    v4 += 1
    for fi113 := 0; fi113 < 11; fi113++ {
        var v114 uint8 = 1
        for fi115 := 0; fi115 < 6; fi115++ {
            if false {
                var v116 string = "hello"
                if true {
                    if false {
                        var v117 int16 = 42
                        var v118 float64 = ((((-2.25 + 1.5) - 1.5) + 1.5) / 1.5)
                        v119 := GenS6{F0: uint16(v2), F1: 42, F2: v117}
                        _ = v117
                        _ = v118
                        _ = v119
                    }
                }
                _ = v116
            } else {
                v114 += (v110 / ((127) | 1))
            }
            if false {
                var v120 float64 = 0.1
                _ = v120
            }
            _ = fi115
        }
        _ = fi113
        _ = v114
    }
    v5 = append(v5, (v111 & (v111 | (v111 % ((v111) | 1)))))
    v111 -= v111
    switch "a" {
    case "\n\t":
        v121 := GenS16{F0: ((v9 + ("a" + v9)) + (v9 + v9))}
        _ = v121
    case "a":
        v11 = (v11 << ((18446744073709551615) & 63))
    }
    v110 = 163
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v110:", v110)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v110
    _ = v111
    _ = v112
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
