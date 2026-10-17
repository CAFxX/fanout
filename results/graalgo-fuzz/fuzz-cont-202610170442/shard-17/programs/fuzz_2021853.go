package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 bool; F2 int32 }
type GenS4 struct { F0 uint64; F1 uint16; F2 bool }
type GenS5 struct { F0 uint8; F1 int8; F2 int32 }
type GenS6 struct { F0 uint32; F1 int; F2 int }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int32; F1 int32; F2 uint8 }
type GenS10 struct { F0 int8; F1 int32 }
type GenS11 struct { F0 int; F1 string; F2 uint8 }
type GenS12 struct { F0 uint8; F1 uint8 }
type GenS13 struct { F0 int16 }
func genFn0(p0 int, p1 uint32, p2 int8) uint8 {
    var v7 uint64 = 993
    v8 := GenS1{F0: false}
    for fi9 := 0; fi9 < 18; fi9++ {
        if (!(false && false)) {
            v7 += v7
            var v10 uint32 = p1
            var v11 float64 = 3.14159
            _ = v10
            _ = v11
        } else {
            var v12 bool = false
            var v13 float64 = float64((0.1 - ((3.14159 / 0.1) / (((((1e10 * ((1.5 * float32((1e10 * (1.5 / (1e10 + 0.0))))) + 1e10)) - 0.0) / 3.14159) / 0.1) - ((1.5 / 1e10) + float32(0.0))))))
            var v14 int32 = 0
            _ = v12
            _ = v13
            _ = v14
        }
        v15 := GenS2{F0: 1}
        _ = fi9
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}
func genFn1() uint8 {
    if false {
        var v16 int = 32767
        switch ("" + "a") {
        case "\n\t":
            v16 += (v16 << ((42) & 63))
        case "hello":
            var v17 int16 = -1
            _ = v17
        default:
            switch ("a" + "\n\t") {
            case "\n\t":
                var v18 float64 = -2.25
                _ = v18
            case "foo bar":
                var v19 int32 = 0
                _ = v19
            }
        }
        _ = v16
    } else {
        if true {
            var v20 int8 = 0
            switch "" {
            case "hello":
                var v21 float64 = 1.5
                _ = v21
            case "\n\t":
                fmt.Println(v20)
            default:
                v20 -= v20
            }
            _ = v20
        } else {
            var v22 int = -582
            for fi23 := 0; fi23 < 3; fi23++ {
                var v24 int16 = 0
                fmt.Println(v22)
                for fi25 := 0; fi25 < 1; fi25++ {
                    switch (fi23 / (((fi23 >> ((319) & 63))) | 1)) {
                    case 9223372036854775807:
                        var v26 int8 = 127
                        _ = v26
                    case 950:
                        var v27 int8 = 42
                        _ = v27
                    default:
                        var v28 uint16 = 32767
                        _ = v28
                    }
                    switch fi25 {
                    case -2147483648:
                        v29 := GenS3{F0: 2, F1: true, F2: -231}
                        _ = v29
                    default:
                        var v30 uint8 = 0
                        _ = v30
                    }
                    _ = fi25
                }
                _ = fi23
                _ = v24
            }
            v31 := GenS1{F0: (0 <= 836)}
            _ = v22
            _ = v31
        }
        var v32 uint64 = 744
        _ = v32
    }
    var v33 string = ""
    v34 := GenS4{F0: 740, F1: 2, F2: (0 > 42)}
    var v35 int = 537
    if (42 >= 4294967295) {
        v36 := []int{v35, -9223372036854775808}
        var v37 int64 = -1
        if true {
            v38 := GenS2{F0: 32767}
            _ = v38
        } else {
            if false {
                v39 := GenS5{F0: 128, F1: -6, F2: 2147483647}
                fmt.Println(v33, v33)
                v35 = 731
                _ = v39
            }
            switch (v35 % (((v35 % ((v35) | 1))) | 1)) {
            case 32767:
                v40 := GenS3{F0: 18446744073709551615, F1: (!(((true && (0.1 > 0.1)) && false) || false)), F2: 328}
                _ = v40
            case 2147483647:
                for fi41 := 0; fi41 < 10; fi41++ {
                    if len(v36) >= 2 {
                        v36 = v36[1:len(v36)]
                    }
                    fmt.Println(v33, v33, v33)
                    _ = fi41
                }
            }
        }
        _ = v36
        _ = v37
    } else {
        v42 := []int{v35, (v35 + 2147483647), (v35 - -892), v35}
        v35 = v35
        var v43 int16 = 0
        _ = v42
        _ = v43
    }
    _ = v33
    _ = v34
    _ = v35
    return 128
}

func main() {
    var v1 uint = 1
    var v2 float64 = 1e10
    v3 := make(map[uint32]int)
    var v4 string = "a"
    v5 := make(map[int]bool)
    var v6 string = v4
    fmt.Println(v4)
    if false {
        switch 109 {
        case 32768:
            for fi44 := 0; fi44 < 9; fi44++ {
                delete(v5, (fi44 - fi44))
                v1 = v1
                _ = fi44
            }
        case 128:
            var v45 int = int(v1)
            _ = v45
        case -761:
            fmt.Println(v6, v6)
        }
        v2 += v2
    } else {
        switch 1 {
        case 42:
            var v46 float32 = (3.14159 + -2.25)
            _ = v46
        case 910:
            v1 = (v1 << (((v1 / ((v1) | 1))) & 63))
        case 32768:
            v2 = (0.1 * 3.14159)
        default:
            v6 = "hello"
        }
    }
    var v47 uint = v1
    switch "foo bar" {
    case "\n\t":
        v47 = 42
    case "":
        v47 = v47
    case "hello":
        for fi48 := 0; fi48 < 11; fi48++ {
            v4 = "foo bar"
            if true {
                for fi49 := 0; fi49 < 15; fi49++ {
                    v50 := GenS2{F0: 42}
                    v47 = (v47 + (v47 >> ((926) & 63)))
                    if true {
                        v51 := v3[0]
                        _ = v51
                    } else {
                        fmt.Println(v4, fi48)
                    }
                    _ = fi49
                    _ = v50
                }
            } else {
                if (42 <= 42) {
                    for fi52 := 0; fi52 < 5; fi52++ {
                        if (((v47 / ((710) | 1)) < v1) && (!(((v1 < (v47 ^ (^(v1)))) || true) && false))) {
                            v47 = (v47 / (((v1 & (v1 << (((v47 + v47)) & 63)))) | 1))
                            v47 -= v47
                            var v53 bool = true
                            if (4294967295 <= 0) {
                                switch (-(fi48)) {
                                case 661:
                                    var v54 float64 = v2
                                    _ = v54
                                case 2147483647:
                                    v55 := v3[942]
                                    _ = v55
                                case 757:
                                    var v56 bool = (127 == 100)
                                    _ = v56
                                }
                                v1 = v47
                            } else {
                                v1 += (v1 | v47)
                                if v53 {
                                    var v57 int16 = int16(v47)
                                    var v58 float32 = 3.14159
                                    var v59 int16 = (v57 << ((2) & 63))
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                                fmt.Println(fi48)
                            }
                            v60 := v3[4294967295]
                            _ = v53
                            _ = v60
                        } else {
                            var v61 bool = true
                            if ((!(!(42 != -9223372036854775808))) && ((fi52 * (fi52 / ((fi48) | 1))) <= 42)) {
                                if v61 {
                                    fmt.Println(fi48)
                                    if false {
                                        var v62 bool = v61
                                        switch 9223372036854775807 {
                                        case 471:
                                            v63 := GenS1{F0: v62}
                                            _ = v63
                                        case -604:
                                            if true {
                                                switch (fi48 & (fi52 - (fi48 + 0))) {
                                                case 32768:
                                                    var v64 int64 = 432
                                                    _ = v64
                                                }
                                                v65 := GenS6{F0: 4294967295, F1: (fi52 % ((fi52) | 1)), F2: fi52}
                                                _ = v65
                                            }
                                        default:
                                            if (780 > 207) {
                                                _ = len(v3)
                                            } else {
                                                v66 := GenS1{F0: (4294967295 > 732)}
                                                if false {
                                                    switch fi52 {
                                                    case 1:
                                                        var v67 int16 = 1000
                                                        _ = v67
                                                    case 0:
                                                        var v68 uint = (v47 ^ 42)
                                                        fmt.Println(v47)
                                                        v69 := GenS3{F0: v68, F1: false, F2: 681}
                                                        _ = v68
                                                        _ = v69
                                                    default:
                                                        v70 := GenS4{F0: 18446744073709551615, F1: 2, F2: true}
                                                        _ = v70
                                                    }
                                                    var v71 int8 = 32
                                                    fmt.Println(fi48)
                                                    v72 := GenS7{F0: ((v71 / ((v71) | 1)) <= v71), F1: v4}
                                                    if false {
                                                        if true {
                                                            v73 := GenS4{F0: 4294967295, F1: 32767, F2: v61}
                                                            v61 = v61
                                                            _ = v73
                                                        } else {
                                                            switch "hello" {
                                                            case "foo bar":
                                                                var v74 bool = (0 > 42)
                                                                _ = v74
                                                            case "\n\t":
                                                                v4 = v6
                                                            case "hello":
                                                                v62 = (0 <= -2147483648)
                                                            }
                                                        }
                                                        if (1 == 194) {
                                                            var v75 int8 = (v71 + v71)
                                                            var v76 int8 = (v71 << (((^(v1))) & 63))
                                                            v2 = 1e10
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                    }
                                                    _ = v71
                                                    _ = v72
                                                } else {
                                                    if (true || true) {
                                                        var v77 uint32 = 42
                                                        if true {
                                                            fmt.Println(v61, v62, v62)
                                                            v77 += (v77 - v77)
                                                        } else {
                                                            v78 := GenS3{F0: (v1 & (v1 * 42)), F1: ((v47 ^ (v47 - v1)) != v47), F2: -813}
                                                            delete(v5, fi48)
                                                            switch (v6 + v4) {
                                                            case "foo bar":
                                                                switch v4 {
                                                                case "hello":
                                                                    v77 = 65535
                                                                case "foo bar":
                                                                    var v79 int = fi48
                                                                    _ = v79
                                                                }
                                                            case "hello":
                                                                v77 = (v77 % ((v77) | 1))
                                                            }
                                                            v80 := GenS8{F0: 1}
                                                            if v62 {
                                                                v1 -= (v47 ^ v1)
                                                                fmt.Println(v77, v6, v61)
                                                            } else {
                                                                var v81 bool = v62
                                                                fmt.Println(v2, v1, fi52)
                                                                if false {
                                                                    v82 := []int{}
                                                                    v62 = true
                                                                    _ = v82
                                                                } else {
                                                                    var v83 uint32 = v77
                                                                    var v84 uint = v47
                                                                    v85 := GenS9{F0: -2147483648, F1: -1, F2: 163}
                                                                    _ = v83
                                                                    _ = v84
                                                                    _ = v85
                                                                }
                                                                var v86 uint16 = 65535
                                                                if v62 {
                                                                    if (((float32(float64(0.1)) + (float32((v2 * (v2 / (v2 - 0.1)))) + 1e10)) * (0.0 - 1.5)) != ((1.5 - 3.14159) + 3.14159)) {
                                                                        v62 = (v47 >= v47)
                                                                        v87 := GenS1{F0: (42 > 554)}
                                                                        switch (fi48 ^ (fi48 << (((v47 ^ 4294967295)) & 63))) {
                                                                        case 2147483647:
                                                                            switch (fi48 << (((^(v1))) & 63)) {
                                                                            case 32767:
                                                                                switch (^(fi48)) {
                                                                                case 42:
                                                                                    v88 := GenS6{F0: v77, F1: -249, F2: -1}
                                                                                    _ = v88
                                                                                case -35:
                                                                                    var v89 bool = v62
                                                                                    _ = v89
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v87
                                                                    } else {
                                                                        var v90 uint = v1
                                                                        _ = v90
                                                                    }
                                                                    if (v47 > 2) {
                                                                        v62 = v81
                                                                    } else {
                                                                        v91 := []int{(^(fi48))}
                                                                        _ = v91
                                                                    }
                                                                    if (2 != uint8(v86)) {
                                                                        var v92 uint64 = 777
                                                                        _ = v92
                                                                    } else {
                                                                        v77 = 0
                                                                        v93 := GenS10{F0: 0, F1: -640}
                                                                        _ = v93
                                                                    }
                                                                } else {
                                                                    var v94 float32 = 0.0
                                                                    v2 = 1.5
                                                                    _ = v94
                                                                }
                                                                if v62 {
                                                                    delete(v3, uint32(v1))
                                                                    v95 := GenS11{F0: fi52, F1: (v6 + ((("\n\t" + v4) + "\n\t") + v6)), F2: 1}
                                                                    var v96 uint64 = 18446744073709551615
                                                                    _ = v95
                                                                    _ = v96
                                                                }
                                                                _ = v81
                                                                _ = v86
                                                            }
                                                            _ = v78
                                                            _ = v80
                                                        }
                                                        v4 = v6
                                                        _ = v77
                                                    } else {
                                                        v97 := GenS4{F0: 1, F1: 370, F2: (651 <= 1)}
                                                        v98 := []int{2147483647}
                                                        _ = v97
                                                        _ = v98
                                                    }
                                                    v4 = v4
                                                    v6 = (v4 + "")
                                                }
                                                _ = v66
                                            }
                                        }
                                        var v99 string = v6
                                        _ = v62
                                        _ = v99
                                    } else {
                                        fmt.Println(fi48)
                                        v100 := GenS12{F0: 0, F1: 127}
                                        _ = v100
                                    }
                                } else {
                                    v2 = v2
                                }
                            } else {
                                if false {
                                    var v101 int8 = -30
                                    _ = v101
                                }
                                switch "a" {
                                case "\n\t":
                                    var v102 uint64 = 4294967295
                                    _ = v102
                                default:
                                    var v103 uint = 18446744073709551615
                                    _ = v103
                                }
                            }
                            var v104 int32 = 288
                            _ = v104
                            _ = v61
                        }
                        _ = fi52
                    }
                    if (65535 == 0) {
                        for fi105 := 0; fi105 < 17; fi105++ {
                            var v106 int16 = -32768
                            v1 += v47
                            if false {
                                var v107 uint8 = 127
                                _ = v107
                            } else {
                                v4 = (v4 + "foo bar")
                            }
                            _ = fi105
                            _ = v106
                        }
                        if false {
                            var v108 uint8 = 32
                            _ = v108
                        }
                        if true {
                            fmt.Println(v4)
                            v4 = v4
                        }
                    } else {
                        v109 := []int{}
                        _ = v109
                    }
                    v2 = v2
                } else {
                    v110 := GenS13{F0: 1}
                    _ = v110
                }
            }
            for fi111 := 0; fi111 < 11; fi111++ {
                if (0 >= 255) {
                    if false {
                        fmt.Println(v1)
                        var v112 float32 = -2.25
                        _ = v112
                    } else {
                        var v113 string = (v6 + (v6 + (v4 + "a")))
                        v1 -= (v47 + v47)
                        if true {
                            v114 := []int{fi111, (fi48 >> ((v1) & 63)), (fi48 & 32768)}
                            delete(v5, 1)
                            _ = v114
                        }
                        switch -2147483648 {
                        case -254:
                            v113 = ("a" + (((v113 + v113) + v113) + v113))
                        case -2147483648:
                            v47 += 42
                        case 420:
                            var v115 uint64 = 701
                            _ = v115
                        default:
                            var v116 uint16 = 255
                            _ = v116
                        }
                        v1 = 1
                        v4 = (v4 + v113)
                        _ = v113
                    }
                }
                var v117 int8 = -128
                _ = fi111
                _ = v117
            }
            _ = fi48
        }
    default:
        for fi118 := 0; fi118 < 10; fi118++ {
            v119 := []int{}
            switch v6 {
            case "":
                v4 = v4
            case "hello":
                fmt.Println(v1)
            case "foo bar":
                v1 += v47
            }
            for fi120 := 0; fi120 < 9; fi120++ {
                var v121 bool = (2 <= 65535)
                v4 = "a"
                var v122 int32 = 0
                _ = fi120
                _ = v121
                _ = v122
            }
            _ = fi118
            _ = v119
        }
    }
    switch v6 {
    case "a":
        if false {
            if (42 > -1) {
                v123 := []int{}
                _ = v123
            } else {
                switch "\n\t" {
                case "foo bar":
                    var v124 float64 = ((((((float64(3.14159) / v2) * (-2.25 * (0.1 / float64(1.5)))) * (v2 / (v2 + (v2 + 0.0)))) / ((v2 - -2.25) - (v2 + 3.14159))) * v2) + float64(float32(v2)))
                    _ = v124
                case "hello":
                    v125 := v5[128]
                    _ = v125
                case "\n\t":
                    v126 := []int{-309, 2147483647, 32768}
                    _ = v126
                }
                var v127 float32 = (-2.25 * (-2.25 * 1e10))
                var v128 int = 4294967295
                _ = v127
                _ = v128
            }
            fmt.Println(v47, v4, v47)
        }
    case "hello":
        v2 = v2
    case "foo bar":
        v129 := v5[1]
        _ = v129
    }
    if (3.14159 == 1.5) {
        if true {
            for fi130 := 0; fi130 < 11; fi130++ {
                fmt.Println(v4)
                v1 = v47
                _ = fi130
            }
        }
    } else {
        v131 := GenS3{F0: 196, F1: true, F2: 1}
        var v132 string = v4
        fmt.Println(v6, v47)
        _ = v131
        _ = v132
    }
    v1 = (v47 & v47)
    v133 := GenS9{F0: -545, F1: -111, F2: 221}
    var v134 int8 = 94
    var v135 int64 = 42
    var v136 int16 = 42
    v2 = ((v2 * float64(((0.0 * 1e10) * ((-2.25 + float32(float64((3.14159 / -2.25)))) + 1.5)))) / (v2 + v2))
    var v137 int8 = v134
    fmt.Println("v1:", v1)
    fmt.Println("v134:", v134)
    fmt.Println("v135:", v135)
    fmt.Println("v136:", v136)
    fmt.Println("v137:", v137)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v133
    _ = v134
    _ = v135
    _ = v136
    _ = v137
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
