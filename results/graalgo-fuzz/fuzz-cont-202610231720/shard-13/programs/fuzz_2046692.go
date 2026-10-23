package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 uint; F1 bool; F2 int16 }
type GenS7 struct { F0 uint8; F1 string; F2 uint }
type GenS8 struct { F0 string; F1 int8 }
type GenS9 struct { F0 int16; F1 uint32 }
type GenS10 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS11 struct { F0 uint8; F1 uint32 }
func genFn0() int32 {
    var v6 int = -223
    v6 = v6
    v7 := []int{}
    var v8 int32 = 2147483647
    var v9 uint16 = 0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (v8 / (((v8 ^ (v8 << ((uint(v9)) & 63)))) | 1))
}
func genFn1(p0 int16, p1 int8) int16 {
    var v10 int16 = p0
    var v11 int16 = (p0 - 818)
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS1{F0: 1}
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 uint64 = 471
            var v16 int32 = 104
            fmt.Println(fi12)
            v11 += (v11 | p0)
            p1 = -27
            v17 := GenS2{F0: (v15 - v15)}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        v18 := []int{(fi12 >> ((18446744073709551615) & 63)), 652}
        _ = fi12
        _ = v13
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return p0
}

func main() {
    var v1 int = 32767
    v2 := []int{v1}
    v3 := make(map[int]int)
    var v4 string = ((("a" + ("hello" + "")) + "a") + ("hello" + ((("a" + "a") + "hello") + "\n\t")))
    var v5 int8 = 56
    v2 = append(v2, -9223372036854775808)
    v1 = v1
    v4 = v4
    for fi19 := 0; fi19 < 7; fi19++ {
        v2 = append(v2, -580)
        for fi20 := 0; fi20 < 6; fi20++ {
            v4 = "hello"
            _ = fi20
        }
        var v21 uint64 = 4294967295
        _ = fi19
        _ = v21
    }
    if true {
        v1 = (^(v1))
    } else {
        switch ((v4 + "hello") + ((v4 + v4) + (v4 + v4))) {
        case "foo bar":
            fmt.Println(v4)
        }
        if (v1 >= -791) {
            if (false || (v1 > v1)) {
                for fi22 := 0; fi22 < 12; fi22++ {
                    var v23 float64 = (-2.25 + 0.1)
                    _ = fi22
                    _ = v23
                }
                var v24 int = 1
                if ((-(v24)) >= (v24 % ((v1) | 1))) {
                    fmt.Println(v1, v24, v4)
                }
                _ = v24
            } else {
                fmt.Println(v4, v5, v4)
                fmt.Println(v5, v4)
            }
            var v25 uint8 = 148
            switch v1 {
            case 1:
                v1 += (v1 * v1)
            case -1:
                if (0.0 > 0.0) {
                    var v26 int8 = (v5 + v5)
                    _ = v26
                } else {
                    v1 += v1
                    var v27 int16 = 32767
                    _ = v27
                }
            case 32767:
                fmt.Println(v1, v4, v5)
            }
            _ = v25
        } else {
            for fi28 := 0; fi28 < 4; fi28++ {
                fmt.Println(v4, fi28)
                var v29 uint32 = 451
                v1 += v1
                _ = fi28
                _ = v29
            }
            var v30 int64 = -9223372036854775808
            _ = v30
        }
        v4 = (v4 + "hello")
    }
    v5 = v5
    var v31 float64 = 1e10
    v32 := GenS1{F0: 1}
    for fi33 := 0; fi33 < 11; fi33++ {
        fmt.Println(v1, v4, v4)
        v34 := GenS3{F0: 65535, F1: v5}
        fmt.Println(v31, v31)
        _ = fi33
        _ = v34
    }
    if (false || false) {
        if false {
            switch v1 {
            case 9223372036854775807:
                for fi35 := 0; fi35 < 20; fi35++ {
                    v5 += v5
                    if (42 < fi35) {
                        if (true && ((!(-1 > -1)) || ((!true) && true))) {
                            var v36 bool = ((v5 & (v5 - (v5 >> ((255) & 63)))) > (v5 << ((4294967295) & 63)))
                            var v37 int = 1
                            v5 += 0
                            _ = v36
                            _ = v37
                        } else {
                            if false {
                                fmt.Println(v4)
                            }
                        }
                        if false {
                            if true {
                                var v38 uint8 = 102
                                fmt.Println(v4, v4, v31)
                                switch "hello" {
                                case "":
                                    if (2147483647 < 630) {
                                        for fi39 := 0; fi39 < 6; fi39++ {
                                            v40 := GenS4{F0: -32768, F1: 1}
                                            _ = fi39
                                            _ = v40
                                        }
                                    } else {
                                        var v41 float32 = 1e10
                                        _ = v41
                                    }
                                case "\n\t":
                                    var v42 uint32 = 647
                                    _ = v42
                                }
                                _ = v38
                            }
                            v5 = v5
                        } else {
                            for fi43 := 0; fi43 < 14; fi43++ {
                                v44 := GenS5{F0: 2147483647, F1: (v4 + v4)}
                                v45 := GenS2{F0: 451}
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                        }
                        switch v4 {
                        case "hello":
                            if ((v1 == v1) && true) {
                                for fi46 := 0; fi46 < 18; fi46++ {
                                    v1 = (fi35 ^ v1)
                                    _ = fi46
                                }
                                var v47 uint = 18446744073709551615
                                v47 = v47
                                _ = v47
                            } else {
                                v48 := GenS6{F0: 263, F1: true, F2: -306}
                                _ = v48
                            }
                        }
                    } else {
                        v49 := GenS7{F0: 127, F1: "\n\t", F2: 65535}
                        var v50 int8 = v5
                        var v51 float64 = 1e10
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                    if (v4 > "") {
                        if (v5 == 42) {
                            v31 -= 1.5
                            var v52 uint64 = 18446744073709551615
                            switch "\n\t" {
                            case "foo bar":
                                var v53 int = (fi35 + (v1 & (v1 << ((2) & 63))))
                                _ = v53
                            default:
                                var v54 int32 = 1
                                _ = v54
                            }
                            _ = v52
                        }
                        v55 := GenS4{F0: -1, F1: -9223372036854775808}
                        _ = v55
                    }
                    _ = fi35
                }
            default:
                v56 := GenS1{F0: 128}
                _ = v56
            }
        } else {
            var v57 uint = 1
            switch (v1 * (v1 >> (((^(v57))) & 63))) {
            case 1:
                if false {
                    switch v1 {
                    case -1:
                        for fi58 := 0; fi58 < 7; fi58++ {
                            switch 162 {
                            case 127:
                                fmt.Println(v57, v31, v31)
                            case 128:
                                v4 = v4
                            }
                            _ = fi58
                        }
                    case 127:
                        switch 32767 {
                        case 42:
                            switch (((v4 + ("\n\t" + v4)) + v4) + v4) {
                            case "":
                                var v59 uint32 = 2
                                _ = v59
                            case "hello":
                                var v60 int16 = -287
                                _ = v60
                            default:
                                for fi61 := 0; fi61 < 12; fi61++ {
                                    var v62 float32 = (1.5 * (1e10 - -2.25))
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi61
                                    _ = v62
                                }
                            }
                        }
                    default:
                        var v63 int32 = -596
                        _ = v63
                    }
                    if false {
                        var v64 uint = v57
                        delete(v3, (v1 * (v1 + v1)))
                        switch v4 {
                        case "hello":
                            v65 := GenS2{F0: 0}
                            _ = v65
                        case "":
                            v4 = v4
                        case "a":
                            switch v1 {
                            case -40:
                                for fi66 := 0; fi66 < 10; fi66++ {
                                    v5 = 127
                                    var v67 int16 = 1
                                    _ = fi66
                                    _ = v67
                                }
                            case -390:
                                if (true || false) {
                                    for fi68 := 0; fi68 < 18; fi68++ {
                                        v1 = -537
                                        _ = fi68
                                    }
                                }
                            case -376:
                                _ = len(v3)
                            default:
                                for fi69 := 0; fi69 < 3; fi69++ {
                                    for fi70 := 0; fi70 < 10; fi70++ {
                                        v71 := v3[fi70]
                                        if (1 == -1) {
                                            v3[v71] = (-(v71))
                                            v72 := v3[v71]
                                            v64 += (v57 * (^(v57)))
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = v72
                                        } else {
                                            var v73 float64 = ((v31 + ((v31 + v31) - 0.0)) / 1e10)
                                            _ = v73
                                        }
                                        if true {
                                            if true {
                                                var v74 uint = (v57 | (v64 - 1))
                                                _ = v74
                                            }
                                            var v75 float32 = -2.25
                                            v31 -= 1e10
                                            v57 = v57
                                            _ = v75
                                        } else {
                                            if true {
                                                var v76 uint8 = 190
                                                if true {
                                                    _ = len(v2)
                                                    if (!false) {
                                                        if ((-786 <= 1000) || false) {
                                                            var v77 bool = true
                                                            fmt.Println(fi70, v31)
                                                            v78 := GenS6{F0: v64, F1: v77, F2: 0}
                                                            _ = v77
                                                            _ = v78
                                                        } else {
                                                            var v79 uint32 = 4294967295
                                                            fmt.Println(v79, v31)
                                                            fmt.Println(v79)
                                                            _ = v79
                                                        }
                                                        if (true || true) {
                                                            delete(v3, (v71 % (((v71 | (-(fi69)))) | 1)))
                                                        } else {
                                                            v71 += (fi70 << (((v64 ^ 1)) & 63))
                                                            var v80 int32 = -2147483648
                                                            _ = v80
                                                        }
                                                    } else {
                                                        if (((false || (-376 <= -2147483648)) || false) || true) {
                                                            v1 += fi69
                                                        }
                                                        var v81 uint32 = 571
                                                        _ = v81
                                                    }
                                                }
                                                var v82 int64 = 9223372036854775807
                                                switch v4 {
                                                case "a":
                                                    fmt.Println(v5, v76, v57)
                                                case "\n\t":
                                                    fmt.Println(fi70)
                                                default:
                                                    v31 += v31
                                                }
                                                _ = v76
                                                _ = v82
                                            } else {
                                                var v83 bool = ((0 >= -136) && (int64(v5) != -754))
                                                v84 := GenS7{F0: 255, F1: "a", F2: v57}
                                                _ = v83
                                                _ = v84
                                            }
                                            var v85 int64 = 1
                                            v86 := GenS8{F0: "hello", F1: 0}
                                            _ = v85
                                            _ = v86
                                        }
                                        fmt.Println(v1)
                                        var v87 uint32 = 592
                                        var v88 int32 = -806
                                        v31 -= ((v31 - v31) * 3.14159)
                                        _ = len(v3)
                                        _ = fi70
                                        _ = v71
                                        _ = v87
                                        _ = v88
                                    }
                                    for fi89 := 0; fi89 < 11; fi89++ {
                                        v4 = ""
                                        var v90 int64 = 872
                                        v91 := GenS6{F0: (v64 - 4294967295), F1: false, F2: -907}
                                        _ = fi89
                                        _ = v90
                                        _ = v91
                                    }
                                    v3[(^(fi69))] = (fi69 & (v1 >> ((v64) & 63)))
                                    _ = fi69
                                }
                            }
                        default:
                            switch v1 {
                            case 70:
                                var v92 int16 = 0
                                _ = v92
                            case 2147483647:
                                for fi93 := 0; fi93 < 15; fi93++ {
                                    if (42 < 2147483647) {
                                        var v94 int16 = 32767
                                        for fi95 := 0; fi95 < 16; fi95++ {
                                            fmt.Println(v57)
                                            v96 := GenS2{F0: 0}
                                            _ = fi95
                                            _ = v96
                                        }
                                        _ = v94
                                    } else {
                                        v97 := GenS9{F0: -625, F1: 411}
                                        v1 -= v1
                                        v98 := GenS9{F0: 0, F1: 65535}
                                        _ = v97
                                        _ = v98
                                    }
                                    _ = fi93
                                }
                            default:
                                v99 := v3[(v1 ^ (v1 << ((4294967295) & 63)))]
                                _ = v99
                            }
                        }
                        _ = v64
                    } else {
                        var v100 int16 = genFn1(797, v5)
                        v57 += (v57 | v57)
                        v101 := v3[v1]
                        _ = v100
                        _ = v101
                    }
                    for fi102 := 0; fi102 < 17; fi102++ {
                        for fi103 := 0; fi103 < 18; fi103++ {
                            var v104 int16 = 1000
                            var v105 uint32 = 42
                            _ = fi103
                            _ = v104
                            _ = v105
                        }
                        var v106 uint = 18446744073709551615
                        _ = fi102
                        _ = v106
                    }
                }
            case 2147483647:
                var v107 int = (v1 / ((int(v57)) | 1))
                _ = v107
            case 4294967295:
                v108 := GenS7{F0: 200, F1: v4, F2: (^(v57))}
                _ = v108
            default:
                if false {
                    switch 42 {
                    case 32767:
                        var v109 uint = v57
                        _ = v109
                    case 154:
                        var v110 uint64 = 42
                        _ = v110
                    case 2147483647:
                        v57 = (^(v57))
                    default:
                        for fi111 := 0; fi111 < 7; fi111++ {
                            fmt.Println(v1, v4)
                            fmt.Println(v31)
                            v57 -= 2
                            _ = fi111
                        }
                    }
                    for fi112 := 0; fi112 < 3; fi112++ {
                        v113 := GenS6{F0: v57, F1: ((9223372036854775807 <= (v1 % (((v1 >> ((v57) & 63))) | 1))) && true), F2: 196}
                        v114 := GenS6{F0: (v57 / ((255) | 1)), F1: true, F2: 1}
                        v115 := v3[fi112]
                        _ = fi112
                        _ = v113
                        _ = v114
                        _ = v115
                    }
                } else {
                    if (32767 >= 0) {
                        switch v4 {
                        case "hello":
                            v116 := GenS10{F0: 70, F1: ((-(v5)) <= v5), F2: 0}
                            _ = v116
                        default:
                            fmt.Println(v1)
                        }
                        v117 := GenS9{F0: -32768, F1: 576}
                        _ = v117
                    } else {
                        var v118 uint8 = 2
                        v119 := GenS11{F0: (^(v118)), F1: 680}
                        fmt.Println(v57, v118)
                        _ = v118
                        _ = v119
                    }
                    v120 := GenS7{F0: 127, F1: ((v4 + "hello") + (((("" + v4) + v4) + "hello") + ((v4 + v4) + (v4 + v4)))), F2: v57}
                    _ = v120
                }
            }
            v121 := v3[(v1 >> ((v57) & 63))]
            _ = v121
            _ = v57
        }
        if (1 >= 42) {
            v122 := GenS10{F0: 1, F1: (18446744073709551615 == 18446744073709551615), F2: 42}
            v123 := GenS1{F0: 2}
            _ = v122
            _ = v123
        }
        var v124 int32 = -2147483648
        _ = v124
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
