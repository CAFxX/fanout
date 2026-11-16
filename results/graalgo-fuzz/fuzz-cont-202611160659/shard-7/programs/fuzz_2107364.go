package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 int32; F2 int16 }
type GenS6 struct { F0 uint8; F1 string; F2 uint8 }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 int32; F1 uint64 }
type GenS9 struct { F0 int32; F1 uint }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int32; F1 uint32 }
type GenS12 struct { F0 string; F1 uint16; F2 uint64 }
type GenS13 struct { F0 int8 }
func genFn0(p0 string, p1 string, p2 int64) string {
    var v5 int64 = p2
    var v6 float32 = 1.5
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p1
}
func genFn1() string {
    if true {
        var v7 float32 = (0.1 / 1.0)
        v7 += v7
        _ = v7
    } else {
        var v8 uint64 = 0
        var v9 bool = false
        if (42 < -1) {
            v10 := []int{42, -1}
            v9 = (42 <= 2)
            if (v9 && (0 >= 42)) {
                v9 = (-1 <= 0)
                for fi11 := 0; fi11 < 4; fi11++ {
                    if (2 >= 42) {
                        if v9 {
                            var v12 uint32 = 1
                            switch fi11 {
                            case 32768:
                                v10 = append(v10, fi11)
                            }
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 5; fi13++ {
                            v14 := GenS1{F0: -603, F1: 32767, F2: 0}
                            v15 := GenS1{F0: 0, F1: 1, F2: 42}
                            var v16 int = fi11
                            _ = fi13
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                    } else {
                        for fi17 := 0; fi17 < 12; fi17++ {
                            var v18 uint16 = 1
                            _ = fi17
                            _ = v18
                        }
                    }
                    var v19 int8 = -9
                    _ = fi11
                    _ = v19
                }
                var v20 uint64 = 4294967295
                _ = v20
            } else {
                for fi21 := 0; fi21 < 3; fi21++ {
                    for fi22 := 0; fi22 < 4; fi22++ {
                        if (32 > 54) {
                            if v9 {
                                v23 := GenS1{F0: -699, F1: 42, F2: 1}
                                _ = v23
                            }
                            if (v9 && v9) {
                                var v24 float64 = (1e10 * ((0.1 - (float64((0.0 / (1.5 + 1e10))) * (1.5 / -2.25))) - 1.5))
                                v25 := GenS1{F0: 9223372036854775807, F1: 65535, F2: 127}
                                _ = v24
                                _ = v25
                            } else {
                                var v26 bool = (1000 != -251)
                                v9 = true
                                if v26 {
                                    v8 += (v8 + 0)
                                    if v26 {
                                        if true {
                                            if true {
                                                v8 -= v8
                                                switch -1 {
                                                case 79:
                                                    v10 = append(v10, (fi21 * fi21))
                                                case -1:
                                                    var v27 float32 = float32((0.0 - float64(0.0)))
                                                    _ = v27
                                                }
                                            } else {
                                                if (false || (41 == 0)) {
                                                    fmt.Println(v26, fi21, v26)
                                                } else {
                                                    var v28 uint = 349
                                                    _ = v28
                                                }
                                                v8 = (^(v8))
                                            }
                                            var v29 uint8 = 129
                                            v29 = (v29 % (((v29 | 65)) | 1))
                                            _ = v29
                                        } else {
                                            var v30 float64 = 3.14159
                                            var v31 float32 = 1.5
                                            _ = v30
                                            _ = v31
                                        }
                                        if v9 {
                                            v32 := GenS2{F0: 42}
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                            _ = v32
                                        }
                                        v33 := GenS3{F0: 42}
                                        _ = v33
                                    }
                                    fmt.Println(fi21)
                                } else {
                                    fmt.Println(fi21)
                                }
                                _ = v26
                            }
                            fmt.Println(fi22, fi21)
                        }
                        fmt.Println(fi21, v8, fi22)
                        _ = fi22
                    }
                    if (!false) {
                        var v34 float32 = 1.5
                        v35 := GenS3{F0: -700}
                        if v9 {
                            v8 += 4294967295
                        }
                        _ = v34
                        _ = v35
                    }
                    _ = fi21
                }
                for fi36 := 0; fi36 < 9; fi36++ {
                    var v37 uint = 65535
                    for fi38 := 0; fi38 < 4; fi38++ {
                        fmt.Println(fi38, v37)
                        var v39 uint16 = 0
                        var v40 int16 = int16(v8)
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    var v41 float32 = (((1.5 / 1.0) / (0.1 / 1.0)) + ((1e10 / float32((-2.25 - (((0.0 * float64((0.1 - 1.5))) - 1.5) + 1e10)))) + 3.14159))
                    _ = fi36
                    _ = v37
                    _ = v41
                }
            }
            _ = v10
        } else {
            if v9 {
                var v42 uint32 = 630
                v43 := []int{127, 4294967295}
                _ = v42
                _ = v43
            }
            v9 = true
            if (!v9) {
                if false {
                    for fi44 := 0; fi44 < 10; fi44++ {
                        v45 := GenS4{F0: 116}
                        v46 := []int{(fi44 << ((1) & 63))}
                        var v47 float64 = 0.0
                        _ = fi44
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    var v48 uint64 = (v8 | (v8 + (v8 | v8)))
                    _ = v48
                }
                switch 42 {
                case -165:
                    for fi49 := 0; fi49 < 13; fi49++ {
                        switch fi49 {
                        case 32767:
                            var v50 int64 = -487
                            _ = v50
                        case 42:
                            for fi51 := 0; fi51 < 11; fi51++ {
                                v52 := GenS1{F0: 42, F1: 1, F2: -85}
                                var v53 int64 = 42
                                var v54 uint64 = v8
                                _ = fi51
                                _ = v52
                                _ = v53
                                _ = v54
                            }
                        case 4294967295:
                            v8 -= v8
                        default:
                            v55 := GenS2{F0: 18446744073709551615}
                            _ = v55
                        }
                        switch "hello" {
                        case "a":
                            for fi56 := 0; fi56 < 1; fi56++ {
                                switch fi49 {
                                case -897:
                                    if v9 {
                                        var v57 uint = 877
                                        var v58 int16 = 1000
                                        v58 -= v58
                                        _ = v57
                                        _ = v58
                                    } else {
                                        var v59 uint32 = 2147483648
                                        _ = v59
                                    }
                                case -2147483648:
                                    v9 = (!(-693 > -1))
                                }
                                var v60 bool = ((^(v8)) <= v8)
                                v61 := []int{4294967295, (fi56 | fi49), -63}
                                v62 := GenS1{F0: -908, F1: 32767, F2: 42}
                                _ = fi56
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                        }
                        var v63 uint16 = 65535
                        _ = fi49
                        _ = v63
                    }
                default:
                    if (float32((((0.1 * (3.14159 / 0.1)) / (0.1 + 1e10)) + 0.0)) == 1.5) {
                        v64 := GenS5{F0: 2147483647, F1: 0, F2: -1}
                        var v65 int16 = 32767
                        _ = v64
                        _ = v65
                    } else {
                        var v66 uint32 = 2147483648
                        _ = v66
                    }
                }
            }
        }
        _ = v8
        _ = v9
    }
    for fi67 := 0; fi67 < 5; fi67++ {
        if true {
            var v68 uint = 763
            for fi69 := 0; fi69 < 19; fi69++ {
                v70 := []int{fi69, -2147483648}
                v71 := GenS2{F0: 614}
                if len(v70) >= 2 {
                    v70 = v70[1:len(v70)]
                }
                _ = fi69
                _ = v70
                _ = v71
            }
            v68 -= v68
            _ = v68
        }
        switch (fi67 % ((fi67) | 1)) {
        case -9223372036854775808:
            v72 := []int{fi67}
            _ = v72
        }
        _ = fi67
    }
    v73 := GenS6{F0: 98, F1: ("foo bar" + ((("a" + "") + "hello") + ("" + ("hello" + ((("hello" + "hello") + "\n\t") + ("hello" + "hello")))))), F2: 255}
    _ = v73
    return ("foo bar" + "")
}

func main() {
    var v1 int16 = -851
    var v2 float64 = (1.5 - 1.5)
    var v3 int = 776
    var v4 int = v3
    v74 := GenS7{F0: v1, F1: -1}
    if false {
        if true {
            v75 := []int{-742, v3, (v3 - 127)}
            var v76 int8 = 1
            _ = v75
            _ = v76
        } else {
            switch "a" {
            case "foo bar":
                var v77 int32 = 0
                _ = v77
            case "\n\t":
                if false {
                    var v78 uint16 = 161
                    for fi79 := 0; fi79 < 1; fi79++ {
                        for fi80 := 0; fi80 < 8; fi80++ {
                            v81 := []int{(fi80 / ((fi79) | 1)), (v4 | (^(v3))), (fi80 + (fi80 & 176))}
                            if ((true || true) || true) {
                                if true {
                                    v82 := GenS2{F0: 65535}
                                    _ = v82
                                }
                                v83 := GenS5{F0: -1, F1: 889, F2: (v1 * 1000)}
                                v4 += fi80
                                if (!(100 > 127)) {
                                    var v84 int32 = -2147483648
                                    if (0 != 18446744073709551615) {
                                        fmt.Println(v84, v4)
                                        v85 := GenS8{F0: (-(v84)), F1: 18446744073709551615}
                                        if (-128 != 42) {
                                            v84 = v84
                                        } else {
                                            var v86 int = (fi80 / ((fi80) | 1))
                                            var v87 uint = 932
                                            v87 = v87
                                            v88 := GenS9{F0: (v84 + v84), F1: 298}
                                            _ = v86
                                            _ = v87
                                            _ = v88
                                        }
                                        if (((fi80 | (v4 % ((v4) | 1))) < 2147483647) && ((true || true) || (0 >= -74))) {
                                            if true {
                                                switch (fi80 - (v3 - -1)) {
                                                case 32768:
                                                    v84 = v84
                                                }
                                            } else {
                                                var v89 int8 = 100
                                                v89 -= -60
                                                _ = v89
                                            }
                                            v78 -= 42
                                        } else {
                                            v1 -= v1
                                        }
                                        if len(v81) > 0 {
                                            v81[0] = v3
                                        }
                                        _ = v85
                                    } else {
                                        if (42 <= 1) {
                                            var v90 uint8 = 2
                                            _ = v90
                                        }
                                        if len(v81) >= 2 {
                                            v81 = v81[1:len(v81)]
                                        }
                                    }
                                    if len(v81) > 0 {
                                        v81[0] = (fi80 % ((9223372036854775807) | 1))
                                    }
                                    var v91 int32 = (v84 ^ v84)
                                    _ = v84
                                    _ = v91
                                }
                                _ = v83
                            } else {
                                v92 := GenS10{F0: 18446744073709551615}
                                _ = v92
                            }
                            _ = fi80
                            _ = v81
                        }
                        var v93 uint16 = v78
                        _ = fi79
                        _ = v93
                    }
                    v94 := GenS6{F0: 200, F1: ("foo bar" + ("hello" + "a")), F2: 42}
                    _ = v78
                    _ = v94
                } else {
                    v3 += -134
                    fmt.Println(v2, v3, v1)
                    for fi95 := 0; fi95 < 17; fi95++ {
                        fmt.Println(v3, fi95, fi95)
                        switch "\n\t" {
                        case "\n\t":
                            v4 = 127
                        case "":
                            v96 := GenS1{F0: 1, F1: 0, F2: -128}
                            _ = v96
                        default:
                            v3 = -9223372036854775808
                        }
                        if true {
                            var v97 int32 = 42
                            v98 := GenS1{F0: 663, F1: 816, F2: -128}
                            v3 -= fi95
                            _ = v97
                            _ = v98
                        } else {
                            for fi99 := 0; fi99 < 15; fi99++ {
                                fmt.Println(fi99, fi99, fi99)
                                _ = fi99
                            }
                            if false {
                                v4 = v4
                            }
                            for fi100 := 0; fi100 < 12; fi100++ {
                                v4 = (fi95 & 4294967295)
                                v2 += 0.0
                                _ = fi100
                            }
                        }
                        _ = fi95
                    }
                }
            }
            v2 = (v2 - 3.14159)
            if true {
                for fi101 := 0; fi101 < 18; fi101++ {
                    var v102 uint16 = 0
                    fmt.Println(v4)
                    _ = fi101
                    _ = v102
                }
                v4 -= 926
                v103 := GenS11{F0: 217, F1: 1}
                _ = v103
            }
        }
        if ((255 == 1) || (!false)) {
            for fi104 := 0; fi104 < 2; fi104++ {
                v105 := []int{v3}
                if true {
                    var v106 int32 = 2147483647
                    if (true && false) {
                        var v107 string = "hello"
                        v108 := GenS11{F0: (v106 ^ 2147483647), F1: 42}
                        _ = v107
                        _ = v108
                    }
                    if ((0.0 <= (1e10 - -2.25)) || false) {
                        var v109 uint = 499
                        for fi110 := 0; fi110 < 14; fi110++ {
                            var v111 uint8 = 129
                            _ = fi110
                            _ = v111
                        }
                        switch "hello" {
                        case "":
                            fmt.Println(v4, v4, v109)
                        case "\n\t":
                            var v112 float64 = ((3.14159 / float64((float32((((v2 * float64(1.5)) * v2) - ((0.0 * v2) * (float64(-2.25) * (v2 / v2))))) * float32(v2)))) + v2)
                            _ = v112
                        case "a":
                            var v113 int = (^(fi104))
                            _ = v113
                        }
                        _ = v109
                    }
                    _ = v106
                } else {
                    for fi114 := 0; fi114 < 3; fi114++ {
                        if (!(!((v2 <= v2) || true))) {
                            v2 = (1e10 - v2)
                        }
                        if true {
                            fmt.Println(v3, v3, v1)
                        } else {
                            var v115 uint32 = 42
                            if len(v105) > 0 {
                                v105[0] = (-(v3))
                            }
                            _ = v115
                        }
                        switch v3 {
                        case 9223372036854775807:
                            var v116 int8 = 0
                            _ = v116
                        case -9223372036854775808:
                            var v117 uint64 = 42
                            _ = v117
                        case 32767:
                            var v118 int16 = v1
                            _ = v118
                        default:
                            v119 := GenS9{F0: 634, F1: 42}
                            _ = v119
                        }
                        _ = fi114
                    }
                }
                _ = fi104
                _ = v105
            }
            v120 := []int{v3, v3}
            fmt.Println(v2)
            _ = v120
        }
    }
    v4 = v3
    v2 = 1.5
    v121 := GenS12{F0: "\n\t", F1: 255, F2: 887}
    switch v3 {
    case -197:
        v122 := GenS13{F0: 1}
        _ = v122
    case 4294967295:
        if ("" != "a") {
            v2 = (3.14159 - (1.5 + float64((((3.14159 * 1.5) + (0.0 - -2.25)) / float32(v2)))))
        }
    default:
        var v123 uint = 1
        _ = v123
    }
    for fi124 := 0; fi124 < 8; fi124++ {
        var v125 int64 = 995
        _ = fi124
        _ = v125
    }
    v2 -= float64(0.1)
    var v126 float32 = (3.14159 / (1e10 - 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v126:", v126)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v121
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v74
}
