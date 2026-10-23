package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 int32; F2 uint8 }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16; F1 int64; F2 int16 }
type GenS11 struct { F0 int64; F1 uint32 }
type GenS12 struct { F0 bool; F1 bool; F2 int16 }
type GenS13 struct { F0 uint8; F1 string }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int64; F1 uint16; F2 int8 }
type GenS16 struct { F0 int8; F1 bool }
type GenS17 struct { F0 int16; F1 bool; F2 int64 }
type GenS18 struct { F0 int8; F1 int }
type GenS19 struct { F0 string }
type GenS20 struct { F0 int8; F1 uint16; F2 int }
type GenS21 struct { F0 int32; F1 uint8; F2 uint }
type GenS22 struct { F0 uint; F1 int8; F2 uint16 }
func genFn0() uint64 {
    switch "a" {
    case "hello":
        var v4 uint64 = 1
        _ = v4
    case "\n\t":
        v5 := GenS1{F0: ""}
        _ = v5
    }
    var v6 uint = 4294967295
    var v7 bool = false
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int8, p1 uint16) uint64 {
    var v8 int16 = 42
    v9 := []int{0, 0}
    p1 += (p1 / (((p1 + (^(p1)))) | 1))
    if ((v8 / ((v8) | 1)) != 582) {
        if true {
            for fi10 := 0; fi10 < 3; fi10++ {
                v8 -= v8
                var v11 uint8 = 181
                var v12 bool = false
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    } else {
        var v13 bool = false
        v8 = (^(v8))
        if true {
            var v14 uint16 = p1
            fmt.Println(v14, p0)
            p0 = p0
            _ = v14
        } else {
            var v15 int16 = v8
            v15 += 32767
            if (0 != 625) {
                v13 = v13
            }
            _ = v15
        }
        _ = v13
    }
    v9 = append(v9, 1)
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return uint64(p1)
}

func main() {
    var v1 string = "foo bar"
    var v2 bool = true
    var v3 string = (v1 + ((v1 + "\n\t") + ("hello" + v1)))
    var v16 uint8 = 42
    switch 42 {
    case -789:
        v17 := GenS2{F0: 42}
        _ = v17
    }
    switch -2147483648 {
    case 4294967295:
        if (!(!(!true))) {
            for fi18 := 0; fi18 < 11; fi18++ {
                for fi19 := 0; fi19 < 1; fi19++ {
                    fmt.Println(v1, v2)
                    var v20 int32 = -228
                    if (v1 <= v3) {
                        v1 = v3
                        v1 = v1
                    }
                    _ = fi19
                    _ = v20
                }
                fmt.Println(fi18, v2, v2)
                var v21 uint16 = 65535
                _ = fi18
                _ = v21
            }
            var v22 uint32 = 2147483648
            _ = v22
        } else {
            v23 := []int{-786, 581, 258, -2147483648}
            v24 := GenS1{F0: v3}
            _ = v23
            _ = v24
        }
    case -9223372036854775808:
        if false {
            var v25 int32 = -1
            if (0 <= -1) {
                v26 := GenS2{F0: 102}
                if (2 != 42) {
                    fmt.Println(v2)
                } else {
                    fmt.Println(v25, v25, v16)
                    fmt.Println(v25)
                    if v2 {
                        v3 = "\n\t"
                    } else {
                        var v27 int16 = -32768
                        if v2 {
                            var v28 uint8 = (v16 % ((v16) | 1))
                            _ = v28
                        }
                        _ = v27
                    }
                }
                for fi29 := 0; fi29 < 10; fi29++ {
                    switch 608 {
                    case -1:
                        v3 = (v1 + v1)
                    case -2147483648:
                        for fi30 := 0; fi30 < 3; fi30++ {
                            v31 := GenS3{F0: 56, F1: -1}
                            v32 := []int{(fi29 / ((-249) | 1))}
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                    default:
                        v3 = ((v1 + (v3 + (("\n\t" + "a") + v1))) + (v3 + v3))
                    }
                    v33 := []int{(fi29 * fi29)}
                    _ = fi29
                    _ = v33
                }
                _ = v26
            }
            _ = v25
        } else {
            if v2 {
                var v34 uint8 = v16
                _ = v34
            } else {
                v3 = v1
            }
            if v2 {
                v35 := GenS4{F0: 71}
                _ = v35
            } else {
                var v36 uint8 = 38
                v37 := []int{}
                _ = v36
                _ = v37
            }
            v16 = v16
        }
    default:
        fmt.Println(v1)
    }
    fmt.Println(v1, v2, v2)
    for fi38 := 0; fi38 < 17; fi38++ {
        var v39 int = fi38
        for fi40 := 0; fi40 < 11; fi40++ {
            switch ((("\n\t" + ("foo bar" + "a")) + v1) + "a") {
            case "a":
                v41 := GenS2{F0: 100}
                _ = v41
            default:
                v3 = (v1 + (v1 + ""))
                v42 := GenS4{F0: 255}
                if (0 == 18446744073709551615) {
                    if false {
                        if true {
                            switch "" {
                            case "":
                                v3 = (v1 + ((v3 + "hello") + "hello"))
                            case "foo bar":
                                if (1 > 84) {
                                    switch v3 {
                                    case "":
                                        if v2 {
                                            var v43 int32 = 42
                                            _ = v43
                                        }
                                    case "foo bar":
                                        if true {
                                            v44 := GenS3{F0: v16, F1: 0}
                                            switch fi38 {
                                            case 127:
                                                v16 += 255
                                            case -1:
                                                v45 := GenS5{F0: (v16 - v16)}
                                                _ = v45
                                            }
                                            v1 = v3
                                            _ = v44
                                        } else {
                                            var v46 string = v3
                                            fmt.Println(v2, v3, v39)
                                            _ = v46
                                        }
                                    default:
                                        var v47 uint32 = 2
                                        _ = v47
                                    }
                                    v48 := GenS3{F0: (^(v16)), F1: 845}
                                    _ = v48
                                } else {
                                    var v49 uint64 = 1
                                    _ = v49
                                }
                            case "a":
                                switch v1 {
                                case "":
                                    var v50 int8 = 100
                                    _ = v50
                                case "a":
                                    var v51 uint8 = v16
                                    _ = v51
                                case "hello":
                                    v39 += 2147483647
                                }
                            default:
                                fmt.Println(v3)
                            }
                            switch (fi40 & v39) {
                            case 9223372036854775807:
                                var v52 uint64 = 309
                                _ = v52
                            case 4294967295:
                                if v2 {
                                    v53 := []int{}
                                    if (-32768 == -1) {
                                        v54 := GenS6{F0: (92 <= -92), F1: -763, F2: v16}
                                        var v55 uint16 = 32767
                                        _ = v54
                                        _ = v55
                                    } else {
                                        v16 += (v16 - v16)
                                    }
                                    _ = v53
                                }
                            case 127:
                                if v2 {
                                    v3 = (v3 + (v1 + v1))
                                } else {
                                    fmt.Println(v39, v2, v16)
                                }
                                v39 -= (v39 ^ 9223372036854775807)
                            default:
                                v56 := []int{v39, int(v16)}
                                _ = v56
                            }
                        } else {
                            switch v3 {
                            case "foo bar":
                                v1 = v3
                            }
                            v2 = (0 == 42)
                        }
                        fmt.Println(fi38, fi40)
                        if v2 {
                            var v57 int16 = 32767
                            v57 = v57
                            _ = v57
                        } else {
                            if (!(((-2147483648 >= -240) && v2) || ((((2147483647 >= 4294967295) || true) && (v2 || v2)) && v2))) {
                                var v58 int16 = 325
                                if true {
                                    var v59 uint8 = v16
                                    fmt.Println(v3, fi38, v3)
                                    var v60 uint = 65535
                                    _ = v59
                                    _ = v60
                                } else {
                                    var v61 float32 = (-2.25 - (0.0 - (float32(1.5) * (1e10 + (float32(0.1) - 0.0)))))
                                    _ = v61
                                }
                                if v2 {
                                    var v62 uint = 161
                                    var v63 bool = true
                                    _ = v62
                                    _ = v63
                                } else {
                                    v1 = v3
                                    switch (fi38 % (((fi38 / (((fi38 | (^(v39)))) | 1))) | 1)) {
                                    case 250:
                                        if (false || (671 >= 42)) {
                                            if true {
                                                v3 = v1
                                                if true {
                                                    v64 := GenS5{F0: v16}
                                                    v65 := GenS7{F0: (v3 + v3), F1: 0}
                                                    _ = v64
                                                    _ = v65
                                                } else {
                                                    v66 := GenS3{F0: 127, F1: (v58 | v58)}
                                                    switch ("\n\t" + (("\n\t" + (v1 + (((v1 + v3) + (((v1 + ("\n\t" + "hello")) + v1) + v3)) + "foo bar"))) + "")) {
                                                    case "":
                                                        var v67 float32 = ((float32(3.14159) * (1.5 / 3.14159)) - 0.1)
                                                        _ = v67
                                                    case "\n\t":
                                                        v2 = v2
                                                    case "a":
                                                        v68 := []int{127}
                                                        _ = v68
                                                    default:
                                                        var v69 uint = 1
                                                        _ = v69
                                                    }
                                                    _ = v66
                                                }
                                            } else {
                                                v3 = "a"
                                            }
                                            v70 := []int{}
                                            _ = v70
                                        }
                                    case -150:
                                        v16 += 1
                                    case -113:
                                        fmt.Println(v1, fi38)
                                    }
                                    var v71 int = v39
                                    if v2 {
                                        v72 := []int{(v71 % (((v71 >> ((255) & 63))) | 1)), v71, fi38}
                                        v39 += v71
                                        var v73 uint = 454
                                        _ = v72
                                        _ = v73
                                    } else {
                                        switch (v1 + ((((v1 + ("a" + v3)) + "") + v3) + "\n\t")) {
                                        case "\n\t":
                                            switch (fi40 / ((v71) | 1)) {
                                            case 9223372036854775807:
                                                fmt.Println(v2)
                                                v71 -= -517
                                            default:
                                                v2 = v2
                                            }
                                        default:
                                            if (v3 < v3) {
                                                var v74 int = v39
                                                if (0 == 917) {
                                                    if v2 {
                                                        v75 := []int{v71, -214, (-(fi40))}
                                                        _ = v75
                                                    }
                                                }
                                                var v76 int16 = v58
                                                _ = v74
                                                _ = v76
                                            } else {
                                                var v77 uint64 = 42
                                                _ = v77
                                            }
                                        }
                                    }
                                    _ = v71
                                }
                                _ = v58
                            } else {
                                switch v1 {
                                case "\n\t":
                                    var v78 int = v39
                                    _ = v78
                                case "":
                                    switch (fi40 + (v39 - fi38)) {
                                    case -271:
                                        v79 := []int{fi40}
                                        _ = v79
                                    default:
                                        if (1e10 != float64((1e10 / ((0.0 + 0.1) - 1.5)))) {
                                            if v2 {
                                                if v2 {
                                                    var v80 uint64 = 555
                                                    v81 := []int{(v39 - (v39 << ((4294967295) & 63)))}
                                                    v81 = append(v81, fi40)
                                                    _ = v80
                                                    _ = v81
                                                } else {
                                                    var v82 uint32 = 252
                                                    _ = v82
                                                }
                                            }
                                            if false {
                                                v16 -= v16
                                            } else {
                                                var v83 uint32 = 2147483647
                                                v84 := []int{fi40, (fi38 % ((fi38) | 1)), (fi40 ^ fi40)}
                                                fmt.Println(fi38, v16, v83)
                                                _ = v83
                                                _ = v84
                                            }
                                        } else {
                                            if v2 {
                                                fmt.Println(v39)
                                                v85 := []int{(fi40 & 127), (fi38 >> ((1) & 63))}
                                                v1 = v1
                                                _ = v85
                                            } else {
                                                v86 := GenS3{F0: 127, F1: 1}
                                                _ = v86
                                            }
                                            v39 = 4294967295
                                            v1 = ""
                                            var v87 bool = v2
                                            if (("" != ("hello" + (v1 + ""))) && (!(18446744073709551615 <= 0))) {
                                                if v2 {
                                                    fmt.Println(v1, v2)
                                                    var v88 float32 = 0.1
                                                    v87 = (!true)
                                                    _ = v88
                                                } else {
                                                    var v89 string = v3
                                                    fmt.Println(v87)
                                                    switch fi38 {
                                                    case 1:
                                                        var v90 int32 = 0
                                                        _ = v90
                                                    }
                                                    _ = v89
                                                }
                                                var v91 int16 = 42
                                                _ = v91
                                            } else {
                                                switch "hello" {
                                                case "a":
                                                    v1 = v1
                                                case "":
                                                    v39 = -9223372036854775808
                                                case "\n\t":
                                                    v87 = (!(-1 >= -1))
                                                }
                                                switch "hello" {
                                                case "foo bar":
                                                    v92 := GenS8{F0: 9223372036854775807}
                                                    _ = v92
                                                case "":
                                                    var v93 bool = (1 != 0)
                                                    _ = v93
                                                case "hello":
                                                    v94 := []int{}
                                                    _ = v94
                                                }
                                                if (!v2) {
                                                    v87 = ((fi40 + (v39 ^ (fi40 % ((fi40) | 1)))) == (fi38 + v39))
                                                }
                                            }
                                            _ = v87
                                        }
                                    }
                                default:
                                    switch fi40 {
                                    case 1:
                                        v16 = 0
                                    case 202:
                                        if v2 {
                                            v95 := GenS9{F0: 914}
                                            var v96 int16 = -32768
                                            _ = v95
                                            _ = v96
                                        } else {
                                            var v97 float32 = (float32(3.14159) / float32(1.5))
                                            _ = v97
                                        }
                                    case -9223372036854775808:
                                        v2 = (0 >= 42)
                                    default:
                                        switch "a" {
                                        case "\n\t":
                                            v39 -= fi38
                                        case "foo bar":
                                            v98 := []int{-117, 127, (fi40 * (fi40 ^ 32767))}
                                            _ = v98
                                        case "a":
                                            v99 := GenS3{F0: v16, F1: 32767}
                                            _ = v99
                                        }
                                    }
                                }
                                v100 := []int{(fi38 % ((fi38) | 1)), 32768, (v39 ^ fi38)}
                                fmt.Println(v39, v16)
                                _ = v100
                            }
                        }
                    } else {
                        v2 = (0 > 42)
                    }
                    v101 := []int{fi40}
                    v102 := GenS2{F0: -128}
                    _ = v101
                    _ = v102
                } else {
                    if (1 >= 336) {
                        var v103 uint64 = 18446744073709551615
                        var v104 float64 = 0.0
                        if ((-2.25 * 0.0) >= 3.14159) {
                            var v105 uint64 = v103
                            var v106 string = v1
                            if ((!((^(v103)) >= v103)) && (((v104 * v104) <= v104) || (687 < 1))) {
                                var v107 bool = v2
                                _ = v107
                            }
                            var v108 uint64 = (v103 + v105)
                            _ = v105
                            _ = v106
                            _ = v108
                        } else {
                            var v109 float64 = v104
                            v104 -= (v109 * 1.5)
                            if (v2 && ((-468 != 1000) || false)) {
                                fmt.Println(v3, v103)
                                var v110 int8 = 1
                                _ = v110
                            }
                            switch "\n\t" {
                            case "a":
                                if v2 {
                                    if false {
                                        if v2 {
                                            if v2 {
                                                v109 += ((float64(1e10) * ((0.0 * (0.0 / v109)) + v104)) * ((v109 + (((v104 * 3.14159) * float64(1.5)) / -2.25)) - float64(3.14159)))
                                            }
                                            v1 = (((v1 + "\n\t") + "foo bar") + v3)
                                            if v2 {
                                                v111 := GenS10{F0: 618, F1: 0, F2: -73}
                                                _ = v111
                                            }
                                        } else {
                                            if false {
                                                v112 := []int{}
                                                var v113 uint8 = (v16 & 0)
                                                _ = v112
                                                _ = v113
                                            }
                                            v114 := []int{fi38, v39, 4294967295}
                                            _ = v114
                                        }
                                    } else {
                                        var v115 int64 = 313
                                        v16 -= 1
                                        switch "" {
                                        case "\n\t":
                                            v116 := GenS8{F0: (v115 / ((9223372036854775807) | 1))}
                                            _ = v116
                                        }
                                        _ = v115
                                    }
                                } else {
                                    if (v16 == 1) {
                                        var v117 int64 = 275
                                        if (float32(-2.25) > ((float32(((0.1 * (v104 + 3.14159)) + 0.0)) - (0.0 * ((3.14159 / 1.5) - 3.14159))) / 0.1)) {
                                            var v118 uint16 = 1
                                            v16 = (v16 * v16)
                                            switch 1 {
                                            case -105:
                                                var v119 int16 = -471
                                                _ = v119
                                            }
                                            _ = v118
                                        }
                                        switch "foo bar" {
                                        case "hello":
                                            v39 += v39
                                        case "\n\t":
                                            v109 -= v104
                                        default:
                                            v2 = v2
                                        }
                                        _ = v117
                                    } else {
                                        var v120 float32 = 3.14159
                                        if v2 {
                                            v121 := []int{(-(fi38)), (fi40 | 0), (-(fi40)), (^(fi38))}
                                            v109 -= v104
                                            if v2 {
                                                switch (fi40 - 32767) {
                                                case 128:
                                                    if v2 {
                                                        v39 += 128
                                                        switch fi40 {
                                                        case 9223372036854775807:
                                                            if len(v121) > 0 {
                                                                v121[0] = (-(v39))
                                                            }
                                                        case 4294967295:
                                                            v39 += (-(v39))
                                                        default:
                                                            fmt.Println(v109)
                                                        }
                                                    }
                                                case 968:
                                                    v39 = (v39 | v39)
                                                case -9223372036854775808:
                                                    if v2 {
                                                        var v122 uint32 = 4294967295
                                                        v120 += (1.5 * (float32(-2.25) + 0.1))
                                                        if v2 {
                                                            if ((v103 >> ((4294967295) & 63)) == (v103 ^ (v103 - (v103 % ((v103) | 1))))) {
                                                                if (true || v2) {
                                                                    v123 := GenS9{F0: v103}
                                                                    var v124 float32 = v120
                                                                    if true {
                                                                        var v125 float64 = v104
                                                                        v3 = "\n\t"
                                                                        _ = v125
                                                                    } else {
                                                                        fmt.Println(v3)
                                                                        switch (fi40 & (v39 << ((4294967295) & 63))) {
                                                                        case 0:
                                                                            if len(v121) > 0 {
                                                                                v121[0] = fi38
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v123
                                                                    _ = v124
                                                                } else {
                                                                    if false {
                                                                        v16 = (v16 & (v16 / ((235) | 1)))
                                                                        var v126 int16 = 494
                                                                        _ = len(v121)
                                                                        _ = v126
                                                                    } else {
                                                                        v120 -= v120
                                                                    }
                                                                }
                                                                if (true && (v122 >= 2147483647)) {
                                                                    var v127 int16 = -32768
                                                                    var v128 int = 967
                                                                    v109 -= v109
                                                                    v1 = v3
                                                                    switch 127 {
                                                                    case 2147483647:
                                                                        var v129 uint16 = 325
                                                                        _ = v129
                                                                    case 229:
                                                                        if len(v121) > 0 {
                                                                            v121[0] = 4294967295
                                                                        }
                                                                    case 127:
                                                                        v16 -= (v16 | v16)
                                                                    default:
                                                                        if (v16 <= 255) {
                                                                            v130 := GenS11{F0: -632, F1: 42}
                                                                            if v2 {
                                                                                v131 := GenS4{F0: 255}
                                                                                _ = v131
                                                                            }
                                                                            _ = v130
                                                                        }
                                                                    }
                                                                    _ = v127
                                                                    _ = v128
                                                                } else {
                                                                    var v132 int = fi38
                                                                    _ = v132
                                                                }
                                                            } else {
                                                                v133 := GenS10{F0: 0, F1: -9223372036854775808, F2: 1000}
                                                                fmt.Println(v2, v103, v104)
                                                                _ = v133
                                                            }
                                                            var v134 uint8 = v16
                                                            v3 = v1
                                                            _ = v134
                                                        } else {
                                                            fmt.Println(v104, v16, fi38)
                                                            if v2 {
                                                                v135 := GenS8{F0: 654}
                                                                if len(v121) >= 2 {
                                                                    v121 = v121[1:len(v121)]
                                                                }
                                                                var v136 int64 = 195
                                                                _ = v135
                                                                _ = v136
                                                            }
                                                            var v137 uint8 = v16
                                                            _ = v137
                                                        }
                                                        v120 = 3.14159
                                                        if ((fi38 >> ((786) & 63)) >= v39) {
                                                            v138 := GenS11{F0: 481, F1: v122}
                                                            v2 = ((255 >= 4294967295) || ((true && ((v103 + v103) <= (^(v103)))) || v2))
                                                            v139 := GenS5{F0: 207}
                                                            _ = v138
                                                            _ = v139
                                                        }
                                                        _ = v122
                                                    }
                                                default:
                                                    if (v2 || (v2 || false)) {
                                                        v3 = v3
                                                        fmt.Println(fi40, v104, v1)
                                                    }
                                                }
                                                var v140 int = (fi40 | (-(fi38)))
                                                var v141 uint16 = 2
                                                _ = v140
                                                _ = v141
                                            }
                                            _ = v121
                                        } else {
                                            fmt.Println(v120, v109, v103)
                                        }
                                        var v142 uint64 = v103
                                        _ = v120
                                        _ = v142
                                    }
                                    switch (-(v39)) {
                                    case 32768:
                                        fmt.Println(v1, v104, v104)
                                    }
                                    var v143 string = ""
                                    v104 += v104
                                    v2 = false
                                    v3 = v143
                                    _ = v143
                                }
                                fmt.Println(v39, fi38)
                            case "foo bar":
                                v144 := []int{fi40, 42, v39, 127}
                                _ = v144
                            default:
                                v145 := []int{(fi40 >> ((183) & 63))}
                                _ = v145
                            }
                            var v146 int8 = 127
                            v109 = v104
                            if v2 {
                                var v147 int32 = 2147483647
                                v16 = 226
                                v148 := []int{(v39 >> ((0) & 63)), (v39 << ((603) & 63)), -903, (fi40 / (((fi40 | (fi40 * (fi38 >> ((0) & 63))))) | 1))}
                                _ = v147
                                _ = v148
                            }
                            switch (fi40 / ((32768) | 1)) {
                            case -1:
                                if (42 != -1) {
                                    if v2 {
                                        v149 := []int{fi40, (v39 << ((42) & 63))}
                                        v3 = (v3 + v1)
                                        if ((false || true) || (("foo bar" <= (v3 + v1)) && (int32(v103) <= 2147483647))) {
                                            if len(v149) >= 2 {
                                                v149 = v149[1:len(v149)]
                                            }
                                            v3 = (v1 + v3)
                                            switch (^(v39)) {
                                            case -9223372036854775808:
                                                var v150 int16 = 363
                                                _ = v150
                                            case 2147483647:
                                                switch "a" {
                                                case "\n\t":
                                                    var v151 uint8 = 2
                                                    _ = v151
                                                }
                                            }
                                            switch v3 {
                                            case "foo bar":
                                                var v152 float64 = 3.14159
                                                _ = v152
                                            }
                                        }
                                        switch -713 {
                                        case 9223372036854775807:
                                            v146 -= (v146 | 42)
                                        case 32768:
                                            var v153 float64 = v109
                                            _ = v153
                                        case 128:
                                            fmt.Println(v3, v109)
                                        }
                                        _ = v149
                                    } else {
                                        v154 := GenS12{F0: (true && (!v2)), F1: (0.0 <= (float64((1.5 + (1e10 / (0.0 * float32((v104 - v104)))))) * v104)), F2: 1}
                                        _ = v154
                                    }
                                    var v155 float32 = float32(float64(-2.25))
                                    _ = v155
                                }
                            default:
                                if v2 {
                                    var v156 float64 = 1.5
                                    v39 += fi40
                                    _ = v156
                                }
                            }
                            v157 := GenS13{F0: 1, F1: ((v3 + v1) + "")}
                            _ = v109
                            _ = v146
                            _ = v157
                        }
                        _ = v103
                        _ = v104
                    } else {
                        if (((1.5 - (0.0 / 1e10)) - -2.25) > float64(-2.25)) {
                            if v2 {
                                v158 := GenS14{F0: true}
                                _ = v158
                            }
                            v159 := []int{(fi38 * -288)}
                            switch v1 {
                            case "foo bar":
                                v160 := GenS15{F0: 260, F1: 65535, F2: -128}
                                _ = v160
                            case "":
                                if true {
                                    var v161 uint64 = 4294967295
                                    switch v1 {
                                    case "a":
                                        v1 = v3
                                    case "":
                                        var v162 uint = 18446744073709551615
                                        _ = v162
                                    case "foo bar":
                                        var v163 int64 = -9223372036854775808
                                        _ = v163
                                    }
                                    v164 := GenS16{F0: 1, F1: true}
                                    _ = v161
                                    _ = v164
                                } else {
                                    switch ("\n\t" + (v1 + (v1 + v3))) {
                                    case "\n\t":
                                        var v165 string = v1
                                        _ = v165
                                    case "foo bar":
                                        if len(v159) >= 2 {
                                            v159 = v159[1:len(v159)]
                                        }
                                    case "":
                                        v39 -= fi38
                                    default:
                                        v166 := GenS16{F0: -23, F1: (1 > -1)}
                                        _ = v166
                                    }
                                    var v167 uint64 = genFn0()
                                    fmt.Println(v2, v167)
                                    _ = v167
                                }
                            case "hello":
                                v168 := GenS17{F0: 0, F1: v2, F2: 0}
                                _ = v168
                            }
                            _ = v159
                        } else {
                            var v169 int32 = -931
                            switch 255 {
                            case 127:
                                var v170 int32 = int32(v16)
                                _ = v170
                            case 32767:
                                v171 := []int{int(v169), (fi40 >> ((2) & 63))}
                                _ = v171
                            case -400:
                                v172 := GenS18{F0: -115, F1: (fi38 >> ((42) & 63))}
                                _ = v172
                            default:
                                switch v1 {
                                case "foo bar":
                                    var v173 uint64 = 18446744073709551615
                                    _ = v173
                                case "a":
                                    v3 = v3
                                case "":
                                    var v174 int16 = 42
                                    _ = v174
                                default:
                                    if v2 {
                                        v1 = v3
                                        v175 := GenS13{F0: v16, F1: v3}
                                        _ = v175
                                    } else {
                                        if true {
                                            if true {
                                                if (42 == v16) {
                                                    v1 = (((v1 + "foo bar") + v3) + v1)
                                                    v39 = fi38
                                                } else {
                                                    var v176 uint64 = 309
                                                    var v177 int32 = v169
                                                    fmt.Println(fi40, v16, v3)
                                                    _ = v176
                                                    _ = v177
                                                }
                                                var v178 uint8 = (^(v16))
                                                v1 = v1
                                                _ = v178
                                            } else {
                                                var v179 string = ("foo bar" + v1)
                                                _ = v179
                                            }
                                            v180 := []int{(-(fi38)), fi40, 2147483647, 275}
                                            var v181 uint16 = 589
                                            _ = v180
                                            _ = v181
                                        } else {
                                            if v2 {
                                                if v2 {
                                                    switch (v39 << ((77) & 63)) {
                                                    case -920:
                                                        v3 = v3
                                                    case 32768:
                                                        v182 := GenS8{F0: 1}
                                                        _ = v182
                                                    default:
                                                        fmt.Println(v169, fi40, v169)
                                                    }
                                                    if false {
                                                        fmt.Println(v169, v3, v2)
                                                        switch v3 {
                                                        case "hello":
                                                            v39 += 127
                                                        case "":
                                                            fmt.Println(v169, v3)
                                                            v183 := GenS14{F0: (2 == 42)}
                                                            var v184 uint16 = 297
                                                            switch fi40 {
                                                            case 32767:
                                                                if (((!true) && v2) || (3.14159 == (3.14159 / float64(1e10)))) {
                                                                    var v185 string = v3
                                                                    fmt.Println(v185, v184)
                                                                    _ = v185
                                                                } else {
                                                                    fmt.Println(v2)
                                                                }
                                                            case -531:
                                                                var v186 int32 = (v169 ^ -811)
                                                                _ = v186
                                                            default:
                                                                if v2 {
                                                                    switch fi40 {
                                                                    case -9223372036854775808:
                                                                        v187 := GenS10{F0: -1, F1: 884, F2: 1000}
                                                                        v1 = "foo bar"
                                                                        switch (v39 / (((fi40 >> ((4294967295) & 63))) | 1)) {
                                                                        case 32767:
                                                                            v188 := []int{(v39 / (((v39 >> ((42) & 63))) | 1)), (v39 + (-(fi40)))}
                                                                            _ = v188
                                                                        case -1:
                                                                            fmt.Println(fi38)
                                                                        default:
                                                                            switch v39 {
                                                                            case 127:
                                                                                switch (("a" + ("\n\t" + (((("foo bar" + ((v3 + v3) + ("hello" + ""))) + v3) + "a") + v1))) + "\n\t") {
                                                                                case "foo bar":
                                                                                    var v189 int = (v39 % ((fi40) | 1))
                                                                                    _ = v189
                                                                                case "hello":
                                                                                    v16 += (v16 / (((v16 + 127)) | 1))
                                                                                }
                                                                            default:
                                                                                v16 -= v16
                                                                            }
                                                                            v16 = (v16 % ((v16) | 1))
                                                                            v16 -= (v16 % ((v16) | 1))
                                                                        }
                                                                        v190 := []int{4294967295}
                                                                        var v191 string = v1
                                                                        _ = v187
                                                                        _ = v190
                                                                        _ = v191
                                                                    }
                                                                    fmt.Println(v2, v39, v16)
                                                                    if (((((3.14159 / 1.5) - 3.14159) - (-2.25 - -2.25)) / (-2.25 + 0.0)) != ((0.1 * 0.0) + 1e10)) {
                                                                        v192 := []int{}
                                                                        _ = v192
                                                                    } else {
                                                                        fmt.Println(v16)
                                                                        v193 := []int{fi40, fi40, (fi40 << ((628) & 63))}
                                                                        _ = v193
                                                                    }
                                                                } else {
                                                                    v3 = "a"
                                                                    v194 := []int{-2147483648, 32767}
                                                                    _ = v194
                                                                }
                                                                v195 := []int{(fi40 / ((v39) | 1)), fi40, fi40, 128}
                                                                _ = v195
                                                            }
                                                            fmt.Println(v1, v39)
                                                            switch "" {
                                                            case "foo bar":
                                                                if (v3 >= (("foo bar" + "") + v1)) {
                                                                    if (867 != 2147483647) {
                                                                        v169 += v169
                                                                        var v196 int = (fi40 >> ((4294967295) & 63))
                                                                        _ = v196
                                                                    } else {
                                                                        v197 := GenS2{F0: 0}
                                                                        var v198 int8 = 127
                                                                        _ = v197
                                                                        _ = v198
                                                                    }
                                                                    var v199 int = fi38
                                                                    _ = v199
                                                                }
                                                            case "hello":
                                                                v200 := GenS18{F0: -1, F1: (fi38 << ((65535) & 63))}
                                                                _ = v200
                                                            case "\n\t":
                                                                var v201 float32 = -2.25
                                                                _ = v201
                                                            }
                                                            fmt.Println(v16)
                                                            v3 = ("foo bar" + v1)
                                                            var v202 float64 = 1e10
                                                            v2 = v2
                                                            fmt.Println(v1, v39, v1)
                                                            _ = v183
                                                            _ = v184
                                                            _ = v202
                                                        case "foo bar":
                                                            v3 = (v3 + ((v1 + v3) + "hello"))
                                                        }
                                                    }
                                                }
                                            } else {
                                                v203 := []int{(^(fi40))}
                                                _ = v203
                                            }
                                        }
                                    }
                                }
                            }
                            v169 = -1
                            _ = v169
                        }
                    }
                }
                _ = v42
            }
            _ = fi40
        }
        _ = fi38
        _ = v39
    }
    for fi204 := 0; fi204 < 17; fi204++ {
        for fi205 := 0; fi205 < 1; fi205++ {
            switch fi205 {
            case 0:
                v1 = (v3 + v1)
            default:
                var v206 uint32 = 2147483648
                _ = v206
            }
            fmt.Println(fi204, fi204)
            v207 := GenS19{F0: ""}
            switch v1 {
            case "a":
                v208 := GenS20{F0: 1, F1: 461, F2: (^(fi204))}
                _ = v208
            case "foo bar":
                var v209 float64 = 1.5
                _ = v209
            case "hello":
                if (float64((0.0 / 1.0)) == 1e10) {
                    v210 := GenS21{F0: -2147483648, F1: (v16 >> ((893) & 63)), F2: 65535}
                    v1 = v3
                    _ = v210
                }
            }
            var v211 float32 = float32(1.5)
            _ = fi205
            _ = v207
            _ = v211
        }
        _ = fi204
    }
    v212 := GenS22{F0: 4294967295, F1: -122, F2: 576}
    v213 := []int{573}
    _ = len(v213)
    if len(v213) > 0 {
        v213[0] = 32767
    }
    v214 := GenS8{F0: 0}
    for fi215 := 0; fi215 < 1; fi215++ {
        v3 = "a"
        v3 = (v3 + v1)
        _ = fi215
    }
    v3 = v1
    v3 = (v3 + v1)
    v3 = ("foo bar" + v1)
    v216 := GenS18{F0: 23, F1: -706}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v213:", len(v213), v213)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v212
    _ = v213
    _ = v214
    _ = v216
    _ = v3
}
