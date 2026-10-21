package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 bool; F2 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS6 struct { F0 uint; F1 int64; F2 int }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int; F1 uint32; F2 uint8 }
type GenS9 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS10 struct { F0 uint8; F1 int64; F2 bool }
type GenS11 struct { F0 int8; F1 int16; F2 int32 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint16; F1 uint; F2 int32 }
func genFn0(p0 uint16, p1 uint64, p2 uint8) int16 {
    v6 := []int{32768, -118, 19}
    p2 = 61
    var v7 float64 = 0.0
    for fi8 := 0; fi8 < 12; fi8++ {
        p1 += p1
        _ = fi8
    }
    v9 := GenS1{F0: 1, F1: 42, F2: 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v9
    return 42
}
func genFn1(p0 int16, p1 uint8) int64 {
    p1 = (p1 % ((p1) | 1))
    p0 += (p0 * p0)
    if (!true) {
        for fi10 := 0; fi10 < 19; fi10++ {
            for fi11 := 0; fi11 < 11; fi11++ {
                fmt.Println(fi11, p0, p0)
                p1 += 127
                if (false && (934 > 431)) {
                    var v12 float32 = 3.14159
                    v13 := GenS1{F0: (^(p0)), F1: 42, F2: -939}
                    switch "" {
                    case "a":
                        fmt.Println(p1, p0)
                    case "foo bar":
                        fmt.Println(p1)
                        switch fi10 {
                        case -288:
                            p0 -= (-(p0))
                        }
                        var v14 int32 = -1
                        var v15 string = "hello"
                        _ = v14
                        _ = v15
                    case "\n\t":
                        fmt.Println(fi10, v12)
                    }
                    p0 += p0
                    var v16 int16 = (p0 | (p0 * p0))
                    v17 := GenS2{F0: 65535}
                    _ = v12
                    _ = v13
                    _ = v16
                    _ = v17
                }
                p0 = 80
                _ = fi11
            }
            _ = fi10
        }
        var v18 float32 = (0.1 + 1e10)
        _ = v18
    } else {
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    return -9223372036854775808
}

func main() {
    var v1 int8 = 126
    var v2 int16 = 262
    var v3 uint8 = 255
    var v4 int = 4294967295
    var v5 uint64 = 4294967295
    var v19 string = "hello"
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 int16 = (v2 / ((v2) | 1))
        v19 = v19
        var v22 uint64 = uint64(v3)
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint16 = 156
    switch v4 {
    case 128:
        v24 := []int{(v4 << ((785) & 63))}
        _ = v24
    case -2147483648:
        v25 := []int{int(v3), (v4 << ((0) & 63)), v4}
        _ = v25
    default:
        v1 = v1
    }
    v26 := []int{int(v3), -9223372036854775808}
    if true {
        if (genFn0(662, (v5 * v5), 201) <= v2) {
            if false {
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                for fi27 := 0; fi27 < 1; fi27++ {
                    for fi28 := 0; fi28 < 9; fi28++ {
                        var v29 int = v4
                        v19 = (v19 + (((v19 + v19) + v19) + (v19 + (v19 + v19))))
                        v1 += (v1 & v1)
                        _ = fi28
                        _ = v29
                    }
                    var v30 uint = 65535
                    _ = fi27
                    _ = v30
                }
            } else {
                if (!false) {
                    if (v3 != v3) {
                        for fi31 := 0; fi31 < 9; fi31++ {
                            var v32 int16 = 1000
                            _ = fi31
                            _ = v32
                        }
                        if true {
                            for fi33 := 0; fi33 < 15; fi33++ {
                                switch (v19 + v19) {
                                case "foo bar":
                                    if len(v26) >= 2 {
                                        v26 = v26[1:len(v26)]
                                    }
                                case "\n\t":
                                    v34 := GenS3{F0: ((v1 > v1) || false), F1: false, F2: 0}
                                    _ = v34
                                }
                                fmt.Println(v5, fi33)
                                var v35 uint32 = uint32(v23)
                                _ = fi33
                                _ = v35
                            }
                        } else {
                            if (1 != 1) {
                                for fi36 := 0; fi36 < 12; fi36++ {
                                    fmt.Println(v23, v1, v2)
                                    v37 := GenS2{F0: 65535}
                                    _ = fi36
                                    _ = v37
                                }
                                v38 := GenS3{F0: false, F1: (v5 != (v5 + (v5 ^ v5))), F2: genFn1(v2, (v3 | (v3 | (^(v3)))))}
                                _ = v38
                            } else {
                                v4 = (v4 ^ (v4 | v4))
                                v2 -= genFn0(0, v5, 97)
                                for fi39 := 0; fi39 < 16; fi39++ {
                                    fmt.Println(v4, v19, fi39)
                                    _ = fi39
                                }
                            }
                            for fi40 := 0; fi40 < 13; fi40++ {
                                switch v19 {
                                case "hello":
                                    if (((float32((-2.25 - -2.25)) + 1.5) >= 0.0) && ((v1 ^ 42) <= (v1 >> ((uint(v5)) & 63)))) {
                                        v5 -= 1
                                    }
                                default:
                                    v41 := GenS2{F0: 42}
                                    _ = v41
                                }
                                _ = fi40
                            }
                            switch (v4 & (v4 + (v4 * (v4 >> ((327) & 63))))) {
                            case -124:
                                if (int32(v5) > 496) {
                                    var v42 uint16 = v23
                                    var v43 uint = 255
                                    _ = v42
                                    _ = v43
                                } else {
                                    var v44 uint8 = v3
                                    _ = v44
                                }
                            }
                        }
                        var v45 uint64 = uint64(v3)
                        _ = v45
                    }
                    for fi46 := 0; fi46 < 19; fi46++ {
                        var v47 uint64 = v5
                        if (1 == 421) {
                            for fi48 := 0; fi48 < 17; fi48++ {
                                fmt.Println(fi46, fi46, v3)
                                _ = fi48
                            }
                        }
                        _ = fi46
                        _ = v47
                    }
                    switch v4 {
                    case 4294967295:
                        if false {
                            fmt.Println(v5, v19, v3)
                            for fi49 := 0; fi49 < 5; fi49++ {
                                var v50 uint32 = 971
                                _ = fi49
                                _ = v50
                            }
                        }
                    default:
                        var v51 int8 = v1
                        _ = v51
                    }
                }
            }
            if false {
                if true {
                    var v52 int16 = (v2 ^ 32767)
                    _ = v52
                }
                v53 := GenS2{F0: 392}
                _ = v53
            }
        }
        v54 := GenS1{F0: (v2 + v2), F1: (v1 & (v1 & int8(v3))), F2: 1}
        _ = v54
    } else {
        v55 := GenS4{F0: v23}
        var v56 float64 = float64(float32(float64(1e10)))
        _ = v55
        _ = v56
    }
    if true {
        for fi57 := 0; fi57 < 17; fi57++ {
            for fi58 := 0; fi58 < 15; fi58++ {
                var v59 int8 = 42
                if ((v3 >= v3) && true) {
                    var v60 int16 = genFn0((v23 - 2), (v5 - (v5 << ((728) & 63))), 127)
                    v1 -= v59
                    v61 := GenS4{F0: uint16(v3)}
                    _ = v60
                    _ = v61
                } else {
                    fmt.Println(v4)
                }
                _ = fi58
                _ = v59
            }
            for fi62 := 0; fi62 < 6; fi62++ {
                if ((float64(3.14159) + float64(1e10)) >= 1.5) {
                    fmt.Println(v5, fi57, v19)
                    v3 = 127
                }
                v1 = v1
                _ = fi62
            }
            var v63 float32 = (3.14159 / float32(1e10))
            _ = fi57
            _ = v63
        }
        v4 += (v4 >> ((uint(v5)) & 63))
    } else {
        v64 := GenS2{F0: uint32(v5)}
        var v65 uint64 = 1
        v1 += v1
        _ = v64
        _ = v65
    }
    if (v23 > v23) {
        _ = len(v26)
        if false {
            fmt.Println(v2, v19)
            v66 := GenS5{F0: 0, F1: (v3 % (((v3 | 117)) | 1)), F2: -1}
            v67 := GenS6{F0: 478, F1: genFn1((v2 / ((genFn0(2, 2, v3)) | 1)), v3), F2: (v4 & (v4 ^ (v4 - v4)))}
            _ = v66
            _ = v67
        } else {
            v68 := GenS7{F0: 42}
            switch (v4 + v4) {
            case 0:
                fmt.Println(v3)
            case 916:
                var v69 int8 = (v1 + (v1 % (((v1 & 100)) | 1)))
                _ = v69
            case 32768:
                v70 := GenS6{F0: 2, F1: int64(v23), F2: v4}
                _ = v70
            }
            if true {
                var v71 uint64 = (v5 | (^(v5)))
                for fi72 := 0; fi72 < 6; fi72++ {
                    var v73 bool = true
                    _ = fi72
                    _ = v73
                }
                _ = v71
            }
            _ = v68
        }
        switch (-(v4)) {
        case 32768:
            v26 = append(v26, (v4 + (v4 & v4)))
        case 0:
            switch v4 {
            case -167:
                if false {
                    if false {
                        if true {
                            switch v19 {
                            case "hello":
                                v74 := GenS8{F0: (^(v4)), F1: 2, F2: 0}
                                _ = v74
                            case "":
                                v4 = v4
                            case "foo bar":
                                for fi75 := 0; fi75 < 15; fi75++ {
                                    if len(v26) >= 2 {
                                        v26 = v26[1:len(v26)]
                                    }
                                    _ = fi75
                                }
                            default:
                                v4 -= v4
                            }
                            v76 := GenS9{F0: int16(v4), F1: 42, F2: 42}
                            _ = v76
                        } else {
                            v23 = (v23 ^ v23)
                        }
                        var v77 int = (v4 + v4)
                        for fi78 := 0; fi78 < 3; fi78++ {
                            v79 := GenS10{F0: uint8(v23), F1: genFn1((v2 / ((284) | 1)), v3), F2: false}
                            _ = fi78
                            _ = v79
                        }
                        _ = v77
                    }
                    if len(v26) > 0 {
                        v26[0] = (^(v4))
                    }
                } else {
                    v4 -= v4
                }
            case 831:
                _ = len(v26)
            case 127:
                if true {
                    if len(v26) > 0 {
                        v26[0] = (v4 << ((823) & 63))
                    }
                    v4 = v4
                    v5 += 0
                } else {
                    var v80 bool = (-2147483648 < 42)
                    v81 := GenS11{F0: (v1 & (v1 % (((-(v1))) | 1))), F1: genFn0(42, v5, v3), F2: 1}
                    _ = v80
                    _ = v81
                }
            default:
                for fi82 := 0; fi82 < 11; fi82++ {
                    var v83 float32 = 1e10
                    v3 = v3
                    _ = fi82
                    _ = v83
                }
            }
        case -2147483648:
            switch (v4 & v4) {
            case -1:
                switch (v19 + "") {
                case "foo bar":
                    for fi84 := 0; fi84 < 16; fi84++ {
                        v2 += v2
                        _ = fi84
                    }
                case "a":
                    for fi85 := 0; fi85 < 8; fi85++ {
                        fmt.Println(v23)
                        _ = fi85
                    }
                default:
                    var v86 uint = 65535
                    _ = v86
                }
            }
        }
    } else {
        switch v4 {
        case 128:
            for fi87 := 0; fi87 < 7; fi87++ {
                for fi88 := 0; fi88 < 17; fi88++ {
                    switch fi87 {
                    case 127:
                        if true {
                            v4 -= fi87
                        } else {
                            fmt.Println(v1, v3, v23)
                            v23 -= 65535
                            fmt.Println(v4, v4)
                        }
                    case 194:
                        switch v19 {
                        case "foo bar":
                            switch v19 {
                            case "hello":
                                var v89 int32 = 1
                                _ = v89
                            case "":
                                v2 -= genFn0(uint16(v5), (v5 & (v5 / ((v5) | 1))), (v3 - 0))
                                var v90 string = "\n\t"
                                _ = v90
                            default:
                                if (396 != 255) {
                                    var v91 string = ""
                                    if true {
                                        v2 = (-(v2))
                                        v2 -= v2
                                    }
                                    _ = v91
                                }
                            }
                        default:
                            v92 := GenS7{F0: 250}
                            fmt.Println(v19, v23, fi88)
                            var v93 uint64 = (v5 << ((4294967295) & 63))
                            _ = v92
                            _ = v93
                        }
                    case -606:
                        if true {
                            switch v4 {
                            case 127:
                                var v94 float32 = 1e10
                                _ = v94
                            default:
                                v19 = ((v19 + v19) + v19)
                            }
                            v95 := GenS2{F0: uint32(v3)}
                            v26 = append(v26, (fi87 + v4))
                            _ = v95
                        } else {
                            if false {
                                var v96 int64 = -9223372036854775808
                                _ = v96
                            } else {
                                switch (v4 << ((67) & 63)) {
                                case -218:
                                    v97 := GenS1{F0: genFn0(802, v5, v3), F1: (^(v1)), F2: 9223372036854775807}
                                    _ = v97
                                case 468:
                                    if (v5 < (v5 ^ (v5 % (((v5 ^ (v5 % ((814) | 1)))) | 1)))) {
                                        switch ("hello" + v19) {
                                        case "foo bar":
                                            var v98 uint64 = (^(v5))
                                            _ = v98
                                        default:
                                            v4 = fi87
                                        }
                                        v4 -= (v4 >> ((0) & 63))
                                        var v99 float32 = 0.1
                                        _ = v99
                                    } else {
                                        var v100 uint32 = 2147483647
                                        var v101 uint = 255
                                        _ = v100
                                        _ = v101
                                    }
                                case 1:
                                    fmt.Println(v1, v4)
                                }
                            }
                        }
                    default:
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                    }
                    _ = fi88
                }
                var v102 uint64 = 4294967295
                v26 = append(v26, v4)
                _ = fi87
                _ = v102
            }
        default:
            _ = len(v26)
        }
        for fi103 := 0; fi103 < 10; fi103++ {
            if ((float32((3.14159 + (float64((3.14159 - ((1.5 - float32(-2.25)) + (-2.25 * 0.0)))) + 0.1))) * (3.14159 * -2.25)) >= 3.14159) {
                if len(v26) > 0 {
                    v26[0] = (v4 | (^(v4)))
                }
                _ = len(v26)
            } else {
                var v104 bool = false
                var v105 int32 = 1
                v106 := GenS12{F0: v23}
                _ = v104
                _ = v105
                _ = v106
            }
            v1 = v1
            _ = fi103
        }
    }
    if false {
        fmt.Println(v1, v1, v5)
        v107 := GenS4{F0: v23}
        var v108 uint16 = 255
        _ = v107
        _ = v108
    } else {
        v1 = v1
        var v109 int64 = -1
        _ = v109
    }
    var v110 uint8 = (v3 - (v3 - (v3 % ((v3) | 1))))
    if true {
        for fi111 := 0; fi111 < 11; fi111++ {
            switch -1 {
            case -467:
                v3 += uint8(v5)
            case -283:
                for fi112 := 0; fi112 < 1; fi112++ {
                    if true {
                        v3 = v3
                    }
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                    v113 := GenS3{F0: true, F1: ((v1 >> ((934) & 63)) <= v1), F2: int64(v2)}
                    _ = fi112
                    _ = v113
                }
            case -2147483648:
                var v114 int = fi111
                _ = v114
            }
            switch v19 {
            case "a":
                fmt.Println(v19, v2, fi111)
            case "hello":
                if true {
                    for fi115 := 0; fi115 < 15; fi115++ {
                        var v116 int32 = -1
                        if true {
                            fmt.Println(v1, v19, v3)
                            v110 = v110
                            if true {
                                var v117 uint64 = (v5 - v5)
                                _ = len(v26)
                                fmt.Println(v1, v116)
                                _ = v117
                            } else {
                                v3 = (v3 >> ((612) & 63))
                            }
                        } else {
                            switch v19 {
                            case "\n\t":
                                if len(v26) > 0 {
                                    v26[0] = fi111
                                }
                            }
                        }
                        var v118 uint32 = 4294967295
                        var v119 int64 = -9223372036854775808
                        _ = fi115
                        _ = v116
                        _ = v118
                        _ = v119
                    }
                    fmt.Println(v19, v110)
                    v120 := GenS13{F0: v23, F1: 18446744073709551615, F2: 0}
                    _ = v120
                }
            }
            fmt.Println(v110)
            _ = fi111
        }
    } else {
        var v121 float64 = ((-2.25 / ((1.5 - (3.14159 + float64(1.5))) / 1.5)) + 0.1)
        v19 = "a"
        _ = v121
    }
    var v122 float64 = 1e10
    var v123 uint64 = v5
    var v124 uint64 = v5
    v3 = v110
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("v122:", v122)
    fmt.Println("v123:", v123)
    fmt.Println("v124:", v124)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v110
    _ = v122
    _ = v123
    _ = v124
    _ = v19
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
