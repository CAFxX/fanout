package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS3 struct { F0 int8; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 uint8; F1 uint }
type GenS8 struct { F0 int32; F1 uint64; F2 int16 }
type GenS9 struct { F0 int16; F1 uint32; F2 int64 }
type GenS10 struct { F0 int32; F1 int; F2 int8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int8; F1 uint64 }
type GenS13 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS14 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int64; F1 bool }
func genFn0(p0 uint32, p1 uint8, p2 uint32) uint32 {
    var v7 int16 = -32768
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            p2 -= p2
            _ = fi9
        }
        fmt.Println(p1)
        p2 = 42
        _ = fi8
    }
    if ((p1 == p1) || false) {
        if (709 != 1) {
            switch 1 {
            case 32768:
                var v10 uint = 2
                _ = v10
            case -544:
                switch "\n\t" {
                case "\n\t":
                    var v11 int32 = -1
                    _ = v11
                case "foo bar":
                    v12 := []int{546, 1, 32768, 0}
                    _ = v12
                default:
                    var v13 int8 = -128
                    _ = v13
                }
            }
            if true {
                p0 -= 695
                v14 := GenS1{F0: v7, F1: 128, F2: 894}
                _ = v14
            } else {
                fmt.Println(v7)
                if (59 >= (p1 - (p1 ^ (^(p1))))) {
                    var v15 uint32 = p0
                    var v16 uint64 = 1
                    _ = v15
                    _ = v16
                }
            }
            if ("\n\t" < "\n\t") {
                var v17 uint32 = p2
                _ = v17
            } else {
                var v18 uint32 = p0
                p2 += (p2 | 602)
                _ = v18
            }
        } else {
            var v19 int32 = 217
            fmt.Println(p1)
            if false {
                var v20 int16 = 42
                v21 := []int{-1, 942, 9223372036854775807}
                _ = v20
                _ = v21
            } else {
                v22 := GenS2{F0: uint32(p1), F1: p0, F2: -1}
                var v23 int = -2147483648
                v24 := []int{v23}
                _ = v22
                _ = v23
                _ = v24
            }
            _ = v19
        }
        var v25 uint32 = 729
        _ = v25
    } else {
        fmt.Println(p0, p1, p2)
        var v26 float32 = 0.1
        v27 := []int{32768, 529}
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        v29 := []int{(fi28 ^ (^(fi28))), (fi28 % ((9223372036854775807) | 1)), 128}
        for fi30 := 0; fi30 < 10; fi30++ {
            v7 = (-(v7))
            if false {
                fmt.Println(p0, v7, p2)
                p2 -= (p0 << ((1) & 63))
                p2 += 65535
            } else {
                var v31 string = ""
                _ = v31
            }
            _ = fi30
        }
        _ = fi28
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 float64 = 1.5
    var v4 uint16 = 42
    var v5 int8 = -1
    v6 := make(map[string]int)
    v4 = (^(v4))
    for fi32 := 0; fi32 < 15; fi32++ {
        _ = len(v2)
        v4 = 2
        if (int8(v4) != (v5 - v5)) {
            switch (fi32 & (fi32 * fi32)) {
            case 4294967295:
                if (false && false) {
                    v33 := []int{19, 387}
                    v34 := GenS1{F0: 0, F1: 128, F2: 0}
                    fmt.Println(v3)
                    _ = v33
                    _ = v34
                } else {
                    if (4294967295 <= genFn0(2, 128, 771)) {
                        var v35 int32 = int32(v5)
                        for fi36 := 0; fi36 < 19; fi36++ {
                            v37 := GenS1{F0: 1000, F1: 1, F2: v4}
                            v5 = -1
                            _ = fi36
                            _ = v37
                        }
                        _ = v35
                    } else {
                        v5 = (v5 * v5)
                    }
                    v3 = (((float64(1e10) * (-2.25 / (v3 - (v3 - (-2.25 / ((float64((3.14159 - 1e10)) - v3) - float64(0.0))))))) - -2.25) - (v3 + (v3 - (0.1 - (((3.14159 * (((((v3 - v3) / v3) / ((0.1 + 0.1) / 3.14159)) * ((v3 - float64(((float32((v3 * (0.0 * ((v3 * v3) / 1.5)))) * 3.14159) + (0.0 / (3.14159 / (3.14159 + (0.0 / 3.14159))))))) + v3)) + ((float64((1e10 - float32((v3 * 1.5)))) + (v3 * float64(-2.25))) + -2.25))) - v3) / 0.1)))))
                }
            default:
                switch (fi32 % (((fi32 & (fi32 + (fi32 / ((fi32) | 1))))) | 1)) {
                case -9223372036854775808:
                    var v38 float32 = -2.25
                    _ = v38
                default:
                    v5 = (v5 & v5)
                }
            }
            switch "foo bar" {
            case "a":
                if (v4 != (v4 - 897)) {
                    v39 := GenS1{F0: -32768, F1: 128, F2: v4}
                    _ = v39
                }
            case "foo bar":
                if true {
                    if true {
                        switch -2147483648 {
                        case -252:
                            for fi40 := 0; fi40 < 5; fi40++ {
                                switch ("" + "foo bar") {
                                case "":
                                    var v41 uint32 = 42
                                    _ = v41
                                default:
                                    var v42 uint64 = 0
                                    v43 := GenS3{F0: 55, F1: 0, F2: 2147483647}
                                    _ = v42
                                    _ = v43
                                }
                                if ((v5 & -128) <= v5) {
                                    v4 = v4
                                    v5 = (-(v5))
                                    var v44 int16 = -1
                                    _ = v44
                                } else {
                                    v3 = v3
                                    if (true || (((v3 / float64(0.1)) * v3) <= v3)) {
                                        v45 := GenS4{F0: 65535, F1: 1, F2: 18446744073709551615}
                                        _ = v45
                                    } else {
                                        var v46 uint32 = 279
                                        if (!(!(65535 >= 0))) {
                                            if (1000 != 32767) {
                                                v47 := GenS4{F0: v46, F1: 2147483647, F2: 388}
                                                delete(v6, "\n\t")
                                                var v48 uint32 = (^(v46))
                                                _ = v47
                                                _ = v48
                                            } else {
                                                v3 -= v3
                                                var v49 uint64 = 674
                                                v2[65535] = "\n\t"
                                                _ = v49
                                            }
                                            fmt.Println(v46, fi40, v4)
                                            if false {
                                                var v50 int8 = v5
                                                v51 := GenS1{F0: int16(v46), F1: 0, F2: v4}
                                                _ = v50
                                                _ = v51
                                            }
                                            var v52 float32 = (0.0 - 1.5)
                                            var v53 uint = 42
                                            if (v53 == 319) {
                                                v54 := GenS5{F0: (v5 + v5)}
                                                var v55 float64 = (((v3 * v3) + float64(1.5)) + v3)
                                                _ = v54
                                                _ = v55
                                            }
                                            _ = v52
                                            _ = v53
                                        } else {
                                            switch "a" {
                                            case "a":
                                                if ((^(fi32)) >= (fi40 << ((2) & 63))) {
                                                    var v56 float64 = float64(float32(v3))
                                                    v3 = v3
                                                    _ = v56
                                                } else {
                                                    switch (fi40 / ((fi32) | 1)) {
                                                    case -2147483648:
                                                        v57 := GenS1{F0: 0, F1: 255, F2: 65535}
                                                        _ = v57
                                                    case 38:
                                                        v58 := GenS4{F0: 956, F1: -865, F2: 0}
                                                        if true {
                                                            fmt.Println(v46, v46, v3)
                                                            var v59 bool = (true && false)
                                                            _ = v59
                                                        } else {
                                                            switch 32768 {
                                                            case 2147483647:
                                                                _ = len(v1)
                                                            }
                                                            fmt.Println(v4, v46)
                                                            switch (fi32 / (((fi40 ^ (fi32 - fi32))) | 1)) {
                                                            case 1:
                                                                if (2 == 4294967295) {
                                                                    if true {
                                                                        var v60 int16 = 1000
                                                                        var v61 uint = 255
                                                                        var v62 int16 = v60
                                                                        _ = v60
                                                                        _ = v61
                                                                        _ = v62
                                                                    } else {
                                                                        if false {
                                                                            switch fi40 {
                                                                            case 4294967295:
                                                                                switch -979 {
                                                                                case 128:
                                                                                    var v63 uint64 = 42
                                                                                    _ = v63
                                                                                case 537:
                                                                                    fmt.Println(v46, fi40)
                                                                                default:
                                                                                    fmt.Println(fi40, fi40, fi40)
                                                                                    if (-407 == 0) {
                                                                                        if (0 < 0) {
                                                                                            v4 += 829
                                                                                            fmt.Println(v5, v3, fi40)
                                                                                            v64 := GenS4{F0: 65535, F1: -1, F2: 42}
                                                                                            _ = v64
                                                                                        } else {
                                                                                            v65 := GenS6{F0: 37, F1: 0}
                                                                                            _ = v65
                                                                                        }
                                                                                    } else {
                                                                                        if (-1 < 42) {
                                                                                            if false {
                                                                                                var v66 uint32 = genFn0((v46 & (^(v46))), 2, (v46 % ((v46) | 1)))
                                                                                                v67 := []int{fi32, (fi32 << ((255) & 63)), fi32}
                                                                                                var v68 int16 = 803
                                                                                                _ = v66
                                                                                                _ = v67
                                                                                                _ = v68
                                                                                            }
                                                                                        } else {
                                                                                            var v69 uint16 = 0
                                                                                            _ = v69
                                                                                        }
                                                                                        if false {
                                                                                            if true {
                                                                                                v46 += v46
                                                                                                var v70 uint8 = 128
                                                                                                _ = v70
                                                                                            }
                                                                                            var v71 float64 = (0.0 / v3)
                                                                                            if (0 == 573) {
                                                                                                var v72 int8 = -112
                                                                                                _ = v72
                                                                                            } else {
                                                                                                if false {
                                                                                                    v46 += 2147483648
                                                                                                    var v73 uint64 = 42
                                                                                                    var v74 uint16 = v4
                                                                                                    _ = v73
                                                                                                    _ = v74
                                                                                                }
                                                                                            }
                                                                                            _ = v71
                                                                                        }
                                                                                        v46 = v46
                                                                                    }
                                                                                    if (0 != 0) {
                                                                                        v4 = (^(v4))
                                                                                    }
                                                                                    if false {
                                                                                        var v75 bool = (v4 <= v4)
                                                                                        if (v4 <= 65535) {
                                                                                            if true {
                                                                                                if v75 {
                                                                                                    v46 += v46
                                                                                                }
                                                                                                var v76 uint32 = (^(v46))
                                                                                                v76 -= genFn0((v76 / ((v46) | 1)), 1, (v76 / (((v76 / ((v46) | 1))) | 1)))
                                                                                                _ = v76
                                                                                            } else {
                                                                                                v77 := v2[(v46 - (v46 | 455))]
                                                                                                var v78 int64 = -9223372036854775808
                                                                                                fmt.Println(fi40, fi32, fi32)
                                                                                                _ = v77
                                                                                                _ = v78
                                                                                            }
                                                                                            fmt.Println(v5)
                                                                                        } else {
                                                                                            if ((!("foo bar" <= ("hello" + "foo bar"))) && v75) {
                                                                                                v79 := GenS5{F0: v5}
                                                                                                var v80 int16 = -222
                                                                                                _ = v79
                                                                                                _ = v80
                                                                                            }
                                                                                            if v75 {
                                                                                                v75 = (-9223372036854775808 >= -1)
                                                                                                var v81 uint32 = (v46 - v46)
                                                                                                _ = v81
                                                                                            } else {
                                                                                                v2[(v46 * (v46 | v46))] = "a"
                                                                                                v4 += v4
                                                                                            }
                                                                                            v3 += v3
                                                                                        }
                                                                                        if false {
                                                                                            var v82 int32 = -2147483648
                                                                                            var v83 uint32 = 0
                                                                                            _ = v82
                                                                                            _ = v83
                                                                                        } else {
                                                                                            var v84 uint = 42
                                                                                            fmt.Println(fi40, fi40, fi32)
                                                                                            var v85 uint32 = (v46 - 65535)
                                                                                            _ = v84
                                                                                            _ = v85
                                                                                        }
                                                                                        _ = v75
                                                                                    }
                                                                                    v5 = 127
                                                                                    var v86 uint16 = v4
                                                                                    v87 := []int{(fi32 / ((-322) | 1)), (fi40 << ((607) & 63)), (fi32 % ((42) | 1)), 32768}
                                                                                    v5 = 100
                                                                                    v88 := v6["a"]
                                                                                    _ = v86
                                                                                    _ = v87
                                                                                    _ = v88
                                                                                }
                                                                            case -282:
                                                                                switch fi32 {
                                                                                case 42:
                                                                                    v89 := GenS3{F0: int8(v46), F1: 4294967295, F2: genFn0(v46, 144, v46)}
                                                                                    _ = v89
                                                                                case -9223372036854775808:
                                                                                    if (!(-733 < (fi40 - (^(fi40))))) {
                                                                                        v90 := GenS7{F0: 97, F1: 868}
                                                                                        fmt.Println(fi40, v46, v46)
                                                                                        v4 = v4
                                                                                        _ = v90
                                                                                    }
                                                                                    v91 := GenS5{F0: -1}
                                                                                    if (uint64(v46) == 42) {
                                                                                        fmt.Println(fi40, fi32)
                                                                                    }
                                                                                    _ = v91
                                                                                case 0:
                                                                                    var v92 bool = (((-2.25 * 0.0) / 0.1) != 0.1)
                                                                                    _ = v92
                                                                                }
                                                                            default:
                                                                                switch (fi32 % ((int(v4)) | 1)) {
                                                                                case 2147483647:
                                                                                    v93 := []int{(fi40 * fi32), fi40, (fi32 * (fi40 * (fi32 >> ((42) & 63)))), (fi32 % ((fi40) | 1))}
                                                                                    _ = v93
                                                                                case -9223372036854775808:
                                                                                    var v94 int64 = 0
                                                                                    _ = v94
                                                                                default:
                                                                                    var v95 uint32 = genFn0(v46, 7, v46)
                                                                                    _ = v95
                                                                                }
                                                                                var v96 uint32 = genFn0(v46, 2, genFn0(v46, 1, (v46 | (v46 >> ((896) & 63)))))
                                                                                v5 = v5
                                                                                _ = v96
                                                                            }
                                                                            v97 := GenS4{F0: 2, F1: -2147483648, F2: 2}
                                                                            _ = v97
                                                                        }
                                                                        fmt.Println(v3, v46, v3)
                                                                    }
                                                                    v4 = (^(v4))
                                                                    v3 = ((v3 * v3) - float64((0.1 * 0.0)))
                                                                }
                                                            default:
                                                                fmt.Println(v4, v5, v5)
                                                            }
                                                            v98 := GenS8{F0: 2147483647, F1: 0, F2: 1000}
                                                            _ = v98
                                                        }
                                                        if true {
                                                            if (false && false) {
                                                                fmt.Println(v4)
                                                                v99 := GenS7{F0: 14, F1: 42}
                                                                fmt.Println(fi32, v4)
                                                                _ = v99
                                                            }
                                                        }
                                                        _ = v58
                                                    case 2147483647:
                                                        var v100 string = "\n\t"
                                                        _ = v100
                                                    }
                                                    switch (-(fi40)) {
                                                    case -2147483648:
                                                        v101 := []int{(-(fi40)), -9223372036854775808}
                                                        _ = v101
                                                    case 0:
                                                        if (1 < 110) {
                                                            if true {
                                                                v4 = (v4 & (v4 << ((65535) & 63)))
                                                                v102 := GenS2{F0: v46, F1: 758, F2: 100}
                                                                _ = v102
                                                            } else {
                                                                var v103 bool = ((4294967295 >= 4294967295) && false)
                                                                var v104 bool = v103
                                                                v2[v46] = "foo bar"
                                                                _ = v103
                                                                _ = v104
                                                            }
                                                            v1[(v46 - genFn0(genFn0((v46 << ((777) & 63)), 139, (v46 + 887)), 127, (v46 - (v46 * 194))))] = fi32
                                                            v105 := GenS9{F0: 1000, F1: (^(v46)), F2: -1}
                                                            _ = v105
                                                        }
                                                    case 486:
                                                        var v106 int64 = 9223372036854775807
                                                        _ = v106
                                                    default:
                                                        v3 += 1e10
                                                    }
                                                }
                                            }
                                            if ((!true) && true) {
                                                v46 += v46
                                                switch (fi32 & fi40) {
                                                case 32768:
                                                    v5 = v5
                                                case -9223372036854775808:
                                                    var v107 uint = 0
                                                    _ = v107
                                                default:
                                                    var v108 uint16 = v4
                                                    _ = v108
                                                }
                                                fmt.Println(v46, fi32, v3)
                                            }
                                        }
                                        v109 := []int{fi40, (^(fi40))}
                                        _ = v109
                                        _ = v46
                                    }
                                }
                                _ = fi40
                            }
                        case -721:
                            var v110 int16 = 42
                            _ = v110
                        case -2147483648:
                            v3 = 0.1
                        }
                        switch (("\n\t" + "\n\t") + "\n\t") {
                        case "\n\t":
                            var v111 int32 = 42
                            _ = v111
                        case "":
                            v3 += -2.25
                        }
                    } else {
                        fmt.Println(v3, v4, v3)
                        v112 := []int{-156, (fi32 * fi32), (fi32 >> ((968) & 63))}
                        _ = v112
                    }
                    if false {
                        if (false && (-9223372036854775808 == 1)) {
                            v4 += v4
                            if false {
                                delete(v2, 302)
                                var v113 string = ("foo bar" + "\n\t")
                                if true {
                                    v114 := GenS10{F0: 42, F1: fi32, F2: v5}
                                    var v115 int16 = -1
                                    for fi116 := 0; fi116 < 6; fi116++ {
                                        v117 := GenS2{F0: 0, F1: genFn0(65535, 127, 2), F2: (v5 | v5)}
                                        v4 += v4
                                        _ = fi116
                                        _ = v117
                                    }
                                    _ = v114
                                    _ = v115
                                } else {
                                    var v118 int64 = 42
                                    v119 := []int{fi32, (fi32 * fi32), fi32, (fi32 | 42)}
                                    _ = v118
                                    _ = v119
                                }
                                _ = v113
                            }
                        }
                        v120 := []int{fi32, (^(fi32)), fi32}
                        fmt.Println(v5)
                        _ = v120
                    } else {
                        v121 := GenS5{F0: v5}
                        var v122 uint8 = 42
                        _ = v121
                        _ = v122
                    }
                } else {
                    fmt.Println(fi32, v3, v4)
                }
            default:
                if (float32(v3) != float32(((v3 - v3) - v3))) {
                    _ = len(v1)
                    for fi123 := 0; fi123 < 17; fi123++ {
                        _ = len(v1)
                        switch "foo bar" {
                        case "a":
                            v124 := GenS5{F0: (^(v5))}
                            _ = v124
                        case "\n\t":
                            v125 := []int{}
                            _ = v125
                        default:
                            v126 := GenS11{F0: (true && false)}
                            _ = v126
                        }
                        if (!(genFn0(2147483647, uint8(v4), 2147483647) == 65535)) {
                            v127 := v2[1]
                            var v128 int32 = 2147483647
                            v3 += 0.0
                            var v129 int32 = (v128 + (v128 ^ v128))
                            v4 -= (v4 << ((4294967295) & 63))
                            v129 = (v128 % ((2147483647) | 1))
                            _ = v127
                            _ = v128
                            _ = v129
                        }
                        _ = fi123
                    }
                    var v130 bool = false
                    _ = v130
                } else {
                    v131 := GenS1{F0: -926, F1: 255, F2: 65535}
                    v5 = v5
                    _ = v131
                }
            }
            v132 := []int{fi32, fi32}
            _ = v132
        }
        _ = fi32
    }
    for fi133 := 0; fi133 < 11; fi133++ {
        delete(v6, "\n\t")
        var v134 uint32 = genFn0(1, 128, 341)
        fmt.Println(fi133)
        _ = fi133
        _ = v134
    }
    fmt.Println(v4, v4, v5)
    v135 := []int{9223372036854775807}
    var v136 int32 = 2147483647
    v137 := GenS12{F0: (v5 / ((127) | 1)), F1: 4294967295}
    if true {
        _ = len(v135)
        var v138 uint16 = v4
        if false {
            if (!true) {
                v139 := GenS13{F0: v136, F1: 65535, F2: 255}
                var v140 string = ""
                _ = v139
                _ = v140
            }
            fmt.Println(v5, v3)
        } else {
            var v141 uint8 = 42
            if ((v141 << ((900) & 63)) <= (v141 | v141)) {
                var v142 uint64 = 0
                var v143 uint32 = 2147483648
                if (true || false) {
                    fmt.Println(v4, v4)
                }
                _ = v142
                _ = v143
            }
            if true {
                fmt.Println(v138, v141, v141)
                v144 := GenS11{F0: true}
                var v145 uint = 65535
                _ = v144
                _ = v145
            } else {
                var v146 float64 = v3
                _ = v146
            }
            _ = v141
        }
        _ = v138
    }
    v147 := GenS14{F0: 236, F1: 0, F2: v4}
    v148 := GenS15{F0: (-(v136))}
    v149 := GenS5{F0: (v5 - (^(v5)))}
    if (!false) {
        if (42 > v5) {
            if (1 != 65535) {
                v150 := GenS16{F0: 1, F1: false}
                v3 = v3
                _ = v150
            }
            for fi151 := 0; fi151 < 2; fi151++ {
                _ = len(v135)
                _ = fi151
            }
            v152 := GenS14{F0: 255, F1: 18446744073709551615, F2: 0}
            _ = v152
        }
    }
    var v153 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v135:", len(v135), v135)
    fmt.Println("v136:", v136)
    fmt.Println("v153:", v153)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v135
    _ = v136
    _ = v137
    _ = v147
    _ = v148
    _ = v149
    _ = v153
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
