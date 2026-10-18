package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 int16; F2 string }
type GenS6 struct { F0 string; F1 int32; F2 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int; F1 uint; F2 int8 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint; F1 uint; F2 int32 }
type GenS11 struct { F0 string; F1 int8; F2 int32 }
func genFn0(p0 uint, p1 bool) string {
    fmt.Println(p1, p1)
    p0 -= 315
    v6 := []int{127}
    p1 = false
    v7 := GenS1{F0: 42, F1: ("a" + (("\n\t" + (("\n\t" + "") + ("a" + ""))) + "\n\t"))}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ""
}
func genFn1(p0 int8, p1 uint, p2 int32) uint {
    var v8 uint = 255
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: -470, F1: "\n\t"}
        var v11 int32 = p2
        fmt.Println(p0, p1, v11)
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := []int{414}
    var v13 uint8 = 213
    if len(v12) > 0 {
        v12[0] = -1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v8
    return (p1 & p1)
}

func main() {
    var v1 float64 = (0.0 + (((1e10 + 1.5) + 0.1) * 1.5))
    var v2 bool = false
    var v3 int8 = 100
    v4 := []uint32{42, 0}
    var v5 uint = 18446744073709551615
    if (0 < 9223372036854775807) {
        if false {
            switch "a" {
            case "foo bar":
                v3 = 127
            case "":
                var v14 float64 = v1
                _ = v14
            case "a":
                var v15 float64 = (1.5 / v1)
                _ = v15
            }
        }
        fmt.Println(v2, v1)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v16 := GenS1{F0: 1000, F1: ("a" + "\n\t")}
    if (2 < 42) {
        var v17 int32 = 0
        v1 -= (v1 / 3.14159)
        _ = v17
    }
    var v18 uint8 = 42
    if (270 > 0) {
        switch "hello" {
        case "foo bar":
            v3 = v3
        case "hello":
            for fi19 := 0; fi19 < 4; fi19++ {
                if v2 {
                    v1 = ((-2.25 * v1) + v1)
                } else {
                    v3 -= (v3 + v3)
                }
                var v20 int16 = 420
                if true {
                    v1 -= v1
                }
                _ = fi19
                _ = v20
            }
        case "\n\t":
            var v21 uint8 = 2
            _ = v21
        default:
            if false {
                v4 = append(v4, 42)
            } else {
                if v2 {
                    v2 = (v2 && (-128 < 42))
                    var v22 uint16 = 255
                    var v23 uint16 = (^(v22))
                    _ = v22
                    _ = v23
                } else {
                    v24 := GenS2{F0: 0, F1: v2}
                    _ = v24
                }
                if (42 <= 255) {
                    var v25 uint32 = 42
                    fmt.Println(v18, v1, v25)
                    if v2 {
                        var v26 float32 = 1.5
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 16; fi27++ {
                            v1 = 0.0
                            v28 := GenS2{F0: 2, F1: v2}
                            for fi29 := 0; fi29 < 12; fi29++ {
                                v25 = v25
                                v30 := GenS2{F0: 32767, F1: v2}
                                _ = fi29
                                _ = v30
                            }
                            _ = fi27
                            _ = v28
                        }
                        fmt.Println(v5, v18, v25)
                    }
                    _ = v25
                } else {
                    for fi31 := 0; fi31 < 15; fi31++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v2 = ((v18 >> ((0) & 63)) > uint8(v5))
                        _ = fi31
                    }
                }
                v18 += v18
            }
        }
        if false {
            var v32 string = "hello"
            switch ("\n\t" + ((v32 + v32) + (("" + (v32 + ((v32 + v32) + ((v32 + (v32 + v32)) + v32)))) + ""))) {
            case "a":
                for fi33 := 0; fi33 < 13; fi33++ {
                    v34 := GenS1{F0: 1, F1: v32}
                    if false {
                        v35 := GenS3{F0: v2}
                        _ = v35
                    } else {
                        v36 := GenS4{F0: (v18 % ((uint8(v5)) | 1))}
                        if (v2 && (0.0 >= (float32((((3.14159 * v1) + 1e10) - v1)) / 0.1))) {
                            v4 = append(v4, 744)
                            v1 = ((v1 + float64(0.1)) / (v1 - v1))
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v36
                    }
                    if false {
                        v37 := GenS1{F0: 32767, F1: "foo bar"}
                        _ = v37
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            var v38 uint64 = 4294967295
                            _ = v38
                        case "a":
                            if true {
                                for fi39 := 0; fi39 < 7; fi39++ {
                                    if (42 > 42) {
                                        var v40 string = v32
                                        v41 := GenS3{F0: false}
                                        _ = v40
                                        _ = v41
                                    } else {
                                        if v2 {
                                            v42 := GenS3{F0: true}
                                            _ = v42
                                        } else {
                                            if (true && false) {
                                                var v43 int16 = 1
                                                v44 := GenS5{F0: v43, F1: (v43 & v43), F2: ("" + "hello")}
                                                _ = v43
                                                _ = v44
                                            } else {
                                                if (32767 <= 1000) {
                                                    switch v32 {
                                                    case "hello":
                                                        v45 := GenS4{F0: v18}
                                                        _ = v45
                                                    case "foo bar":
                                                        if (-1 != -1) {
                                                            v46 := GenS4{F0: v18}
                                                            _ = v46
                                                        } else {
                                                            v3 = (-(v3))
                                                        }
                                                        if false {
                                                            v4 = append(v4, 2147483648)
                                                            var v47 uint = (v5 + 65535)
                                                            _ = v47
                                                        } else {
                                                            v48 := GenS2{F0: 1, F1: true}
                                                            v18 += (v18 * (v18 ^ (v18 - 1)))
                                                            v3 = (v3 / ((-128) | 1))
                                                            v32 = v32
                                                            var v49 float32 = (-2.25 / (float32(1.5) + ((1.5 / 1e10) + (-2.25 - -2.25))))
                                                            var v50 int64 = 1
                                                            _ = v48
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                    default:
                                                        if v2 {
                                                            v1 = 0.0
                                                            switch fi33 {
                                                            case 128:
                                                                var v51 int = (fi33 * (fi33 >> (((v5 + v5)) & 63)))
                                                                _ = v51
                                                            case -210:
                                                                switch (fi39 ^ fi33) {
                                                                case 127:
                                                                    v18 = v18
                                                                }
                                                            default:
                                                                if v2 {
                                                                    var v52 string = v32
                                                                    var v53 uint32 = 667
                                                                    var v54 int8 = (v3 % ((-128) | 1))
                                                                    v54 = (^(v54))
                                                                    fmt.Println(v3, v1)
                                                                    _ = v52
                                                                    _ = v53
                                                                    _ = v54
                                                                } else {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    v55 := GenS6{F0: "a", F1: 1, F2: 574}
                                                                    fmt.Println(fi33, v18, v18)
                                                                    v56 := GenS6{F0: v32, F1: 1, F2: 396}
                                                                    var v57 float32 = -2.25
                                                                    _ = v55
                                                                    _ = v56
                                                                    _ = v57
                                                                }
                                                            }
                                                        } else {
                                                            var v58 uint = (^(v5))
                                                            if (!("\n\t" <= (v32 + v32))) {
                                                                v59 := GenS1{F0: 0, F1: (("\n\t" + v32) + v32)}
                                                                _ = v59
                                                            } else {
                                                                v60 := GenS3{F0: (true || ((3.14159 > 1e10) || v2))}
                                                                v61 := GenS1{F0: 0, F1: v32}
                                                                if (-107 != 1) {
                                                                    var v62 float64 = v1
                                                                    _ = len(v4)
                                                                    if true {
                                                                        fmt.Println(v1, fi39, v2)
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        if true {
                                                                            v63 := GenS2{F0: 761, F1: (v2 || ((38 == 255) || false))}
                                                                            var v64 uint32 = 42
                                                                            var v65 int8 = 100
                                                                            _ = v63
                                                                            _ = v64
                                                                            _ = v65
                                                                        } else {
                                                                            switch (-(fi39)) {
                                                                            case -56:
                                                                                v2 = (true && ((1 > 689) && v2))
                                                                            }
                                                                            if true {
                                                                                v66 := GenS2{F0: 32767, F1: (!v2)}
                                                                                var v67 uint64 = 0
                                                                                var v68 int64 = 42
                                                                                _ = v66
                                                                                _ = v67
                                                                                _ = v68
                                                                            }
                                                                        }
                                                                        var v69 int32 = -2147483648
                                                                        switch ((v32 + v32) + (v32 + (v32 + v32))) {
                                                                        case "\n\t":
                                                                            var v70 int8 = (v3 - 42)
                                                                            _ = v70
                                                                        case "hello":
                                                                            if len(v4) > 0 {
                                                                                v4[0] = 212
                                                                            }
                                                                        }
                                                                        v62 += 1.5
                                                                        fmt.Println(fi39)
                                                                        _ = v69
                                                                    }
                                                                    _ = v62
                                                                } else {
                                                                    fmt.Println(v1)
                                                                    v3 = v3
                                                                    v4 = append(v4, 2)
                                                                }
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                            _ = len(v4)
                                                            _ = v58
                                                        }
                                                    }
                                                    var v71 float32 = ((0.1 - (1.5 * ((float32(0.0) * (-2.25 / (1.5 * (1.5 + 0.0)))) + 0.1))) * 0.0)
                                                    if (2 <= 229) {
                                                        v72 := GenS5{F0: -1, F1: 1, F2: "foo bar"}
                                                        _ = len(v4)
                                                        switch "foo bar" {
                                                        case "":
                                                            var v73 float64 = v1
                                                            _ = v73
                                                        default:
                                                            v74 := GenS4{F0: v18}
                                                            if v2 {
                                                                if v2 {
                                                                    v2 = v2
                                                                    var v75 int8 = v3
                                                                    var v76 int64 = 1
                                                                    _ = v75
                                                                    _ = v76
                                                                } else {
                                                                    switch (v32 + "foo bar") {
                                                                    case "a":
                                                                        v5 = v5
                                                                    case "foo bar":
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    }
                                                                    var v77 float32 = v71
                                                                    _ = v77
                                                                }
                                                                if (181 != (v18 * 199)) {
                                                                    var v78 uint64 = 18446744073709551615
                                                                    v3 = (v3 | v3)
                                                                    _ = v78
                                                                }
                                                                switch (^(fi39)) {
                                                                case -837:
                                                                    v2 = (v71 < 0.1)
                                                                    var v79 bool = (558 >= fi39)
                                                                    fmt.Println(v1, v71, v18)
                                                                    _ = v79
                                                                case 2147483647:
                                                                    var v80 string = "foo bar"
                                                                    _ = v80
                                                                }
                                                            }
                                                            _ = v74
                                                        }
                                                        _ = v72
                                                    }
                                                    v3 = v3
                                                    _ = v71
                                                } else {
                                                    fmt.Println(fi39)
                                                    var v81 int = 497
                                                    fmt.Println(v3, fi33, fi39)
                                                    _ = v81
                                                }
                                                var v82 uint16 = 42
                                                switch fi33 {
                                                case 0:
                                                    v4 = append(v4, 311)
                                                case -2147483648:
                                                    var v83 uint = (^(v5))
                                                    fmt.Println(v2, v83, fi33)
                                                    switch v32 {
                                                    case "foo bar":
                                                        v82 -= 0
                                                    case "\n\t":
                                                        v18 = v18
                                                    default:
                                                        fmt.Println(fi33)
                                                    }
                                                    _ = v83
                                                case 711:
                                                    if v2 {
                                                        switch 942 {
                                                        case 394:
                                                            v32 = v32
                                                        case 32768:
                                                            var v84 int32 = 0
                                                            _ = v84
                                                        case 767:
                                                            v85 := GenS2{F0: (v82 / (((v82 ^ 32767)) | 1)), F1: (!v2)}
                                                            _ = v85
                                                        }
                                                        v4 = append(v4, 2147483648)
                                                        if true {
                                                            if v2 {
                                                                if v2 {
                                                                    if true {
                                                                        v5 -= v5
                                                                    } else {
                                                                        if len(v4) > 0 {
                                                                            v4[0] = 359
                                                                        }
                                                                        if true {
                                                                            var v86 float32 = 1e10
                                                                            _ = v86
                                                                        } else {
                                                                            v87 := GenS2{F0: 2, F1: (!true)}
                                                                            _ = v87
                                                                        }
                                                                        var v88 uint8 = 255
                                                                        _ = v88
                                                                    }
                                                                    if ((int32(v5) < 2147483647) && ((fi39 < 55) || (int32(v82) != 877))) {
                                                                        v82 -= (^(v82))
                                                                        if false {
                                                                            var v89 uint64 = 608
                                                                            if (!true) {
                                                                                switch (fi39 ^ (fi33 % ((fi39) | 1))) {
                                                                                case -9223372036854775808:
                                                                                    v82 = v82
                                                                                case 128:
                                                                                    var v90 float32 = (1.5 - 0.1)
                                                                                    _ = v90
                                                                                default:
                                                                                    fmt.Println(v32, v82, v18)
                                                                                }
                                                                                fmt.Println(v32)
                                                                                switch fi33 {
                                                                                case 127:
                                                                                    v89 = v89
                                                                                default:
                                                                                    var v91 int64 = -9223372036854775808
                                                                                    _ = v91
                                                                                }
                                                                                v5 -= v5
                                                                            }
                                                                            var v92 uint16 = (v82 << (((v5 % (((^(v5))) | 1))) & 63))
                                                                            _ = v89
                                                                            _ = v92
                                                                        } else {
                                                                            if v2 {
                                                                                var v93 int16 = -32768
                                                                                if (990 != 0) {
                                                                                    v94 := GenS7{F0: "hello"}
                                                                                    _ = v94
                                                                                } else {
                                                                                    var v95 uint16 = 255
                                                                                    v32 = ("" + v32)
                                                                                    _ = v95
                                                                                }
                                                                                _ = v93
                                                                            } else {
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                                fmt.Println(v3, v2)
                                                                            }
                                                                            fmt.Println(v2, v82)
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v96 uint16 = v82
                                                                    _ = v96
                                                                }
                                                                fmt.Println(fi33, v2)
                                                            } else {
                                                                var v97 int32 = int32(fi33)
                                                                var v98 float64 = v1
                                                                _ = v97
                                                                _ = v98
                                                            }
                                                        } else {
                                                            _ = len(v4)
                                                            switch 127 {
                                                            case 2147483647:
                                                                var v99 float32 = float32(-2.25)
                                                                _ = v99
                                                            case 0:
                                                                v5 += v5
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v82
                                            }
                                            v5 = genFn1(v3, v5, -594)
                                            _ = len(v4)
                                            v100 := GenS4{F0: v18}
                                            if (v2 || (v1 == (0.0 * v1))) {
                                                fmt.Println(v5, v5)
                                                v101 := GenS6{F0: v32, F1: int32(v3), F2: 42}
                                                v5 += v5
                                                switch (fi39 >> ((v5) & 63)) {
                                                case 582:
                                                    v5 += (v5 & (^(v5)))
                                                case 624:
                                                    v32 = v32
                                                }
                                                v18 = v18
                                                _ = v101
                                            } else {
                                                v32 = v32
                                                fmt.Println(v3, v1, v18)
                                                if (float64(1e10) >= (((-2.25 - 0.1) + (1e10 - 0.1)) + v1)) {
                                                    v18 -= v18
                                                } else {
                                                    var v102 uint32 = 96
                                                    fmt.Println(v102)
                                                    var v103 bool = (4294967295 <= 1)
                                                    _ = v102
                                                    _ = v103
                                                }
                                            }
                                            _ = v100
                                        }
                                        if v2 {
                                            if v2 {
                                                switch (-(fi33)) {
                                                case 66:
                                                    v4 = append(v4, 1)
                                                case 127:
                                                    v104 := GenS4{F0: v18}
                                                    _ = v104
                                                case 32768:
                                                    v5 = genFn1((-(v3)), 65535, -460)
                                                }
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v105 float32 = 0.1
                                            _ = v105
                                        }
                                        v3 += (v3 - (v3 ^ v3))
                                        var v106 float32 = (3.14159 / 3.14159)
                                        _ = v106
                                    }
                                    var v107 uint = v5
                                    _ = fi39
                                    _ = v107
                                }
                            }
                        default:
                            v108 := GenS1{F0: int16(v5), F1: ((v32 + ("foo bar" + "hello")) + "foo bar")}
                            _ = v108
                        }
                        v109 := GenS7{F0: ""}
                        v110 := GenS4{F0: v18}
                        _ = v109
                        _ = v110
                    }
                    _ = fi33
                    _ = v34
                }
            }
            _ = v32
        } else {
            v18 += 255
        }
        if v2 {
            switch -2147483648 {
            case 9223372036854775807:
                v4 = append(v4, 2147483648)
            }
            for fi111 := 0; fi111 < 11; fi111++ {
                var v112 string = ("" + "")
                v5 = genFn1((^(v3)), (v5 / ((genFn1((v3 / ((-128) | 1)), 1, 2147483647)) | 1)), 176)
                v5 -= v5
                _ = fi111
                _ = v112
            }
            if false {
                v113 := GenS8{F0: 128, F1: 1, F2: 43}
                if false {
                    v5 -= genFn1((-(v3)), 255, 1)
                } else {
                    if false {
                        v5 += v5
                        var v114 uint32 = 219
                        for fi115 := 0; fi115 < 7; fi115++ {
                            for fi116 := 0; fi116 < 1; fi116++ {
                                if len(v4) > 0 {
                                    v4[0] = (^(v114))
                                }
                                v5 -= (v5 - (v5 | genFn1((^(v3)), 2, 0)))
                                _ = fi116
                            }
                            _ = fi115
                        }
                        _ = v114
                    } else {
                        for fi117 := 0; fi117 < 15; fi117++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            for fi118 := 0; fi118 < 20; fi118++ {
                                v5 = genFn1((v3 << ((1) & 63)), (v5 << ((v5) & 63)), 42)
                                var v119 uint32 = 4294967295
                                _ = fi118
                                _ = v119
                            }
                            if (9223372036854775807 > 0) {
                                fmt.Println(v1)
                                v1 = v1
                            } else {
                                switch fi117 {
                                case 32767:
                                    v1 -= v1
                                case 42:
                                    var v120 int8 = v3
                                    _ = v120
                                default:
                                    fmt.Println(v1, v5, fi117)
                                }
                                var v121 string = "foo bar"
                                var v122 int16 = 42
                                _ = v121
                                _ = v122
                            }
                            _ = fi117
                        }
                        v2 = (9223372036854775807 > 1)
                        for fi123 := 0; fi123 < 20; fi123++ {
                            v18 = (^(v18))
                            _ = fi123
                        }
                    }
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(v1, v2, v5)
                    case "hello":
                        v124 := GenS7{F0: ("hello" + "foo bar")}
                        _ = v124
                    default:
                        var v125 uint64 = 447
                        _ = v125
                    }
                }
                fmt.Println(v1, v5)
                _ = v113
            }
        } else {
            fmt.Println(v1, v3, v2)
        }
    }
    v126 := GenS5{F0: 1, F1: 43, F2: ((("foo bar" + ("a" + "a")) + ("hello" + "a")) + "hello")}
    for fi127 := 0; fi127 < 4; fi127++ {
        if (0.1 == ((float32(1.5) - 3.14159) - 0.0)) {
            v3 = v3
            var v128 int64 = -1
            v18 = v18
            _ = v128
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v129 int = (fi127 + fi127)
        _ = fi127
        _ = v129
    }
    for fi130 := 0; fi130 < 20; fi130++ {
        var v131 int64 = -1
        if v2 {
            v2 = v2
        }
        _ = fi130
        _ = v131
    }
    switch 42 {
    case 0:
        fmt.Println(v3, v2)
    }
    switch (((((("foo bar" + "") + "hello") + (("" + "a") + "")) + "a") + ("" + (("foo bar" + "") + (("foo bar" + "hello") + ("" + "a"))))) + "\n\t") {
    case "a":
        fmt.Println(v3, v18, v5)
    case "foo bar":
        v5 += (v5 - (v5 >> ((4294967295) & 63)))
    }
    var v132 uint8 = (v18 * (v18 + v18))
    fmt.Println(v5, v132)
    if ((v132 & (v132 << ((v5) & 63))) > v18) {
        for fi133 := 0; fi133 < 9; fi133++ {
            for fi134 := 0; fi134 < 1; fi134++ {
                var v135 float64 = float64(1.5)
                if true {
                    if (!(4294967295 == 445)) {
                        v18 = v18
                        v3 += 1
                        if ((("foo bar" + ("\n\t" + (("hello" + "\n\t") + "foo bar"))) + (("hello" + "hello") + "a")) >= ((("a" + "") + "a") + ("" + ""))) {
                            switch "a" {
                            case "\n\t":
                                var v136 int = fi133
                                v18 -= (v18 & 1)
                                _ = v136
                            }
                            if false {
                                switch 127 {
                                case 994:
                                    if (!true) {
                                        switch fi134 {
                                        case 1:
                                            switch ("hello" + "\n\t") {
                                            case "a":
                                                var v137 int8 = (v3 + 7)
                                                _ = v137
                                            default:
                                                v3 -= v3
                                            }
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = 760
                                        }
                                        if ("" == "a") {
                                            if (false && v2) {
                                                if true {
                                                    v138 := GenS2{F0: 255, F1: ((fi133 % ((fi134) | 1)) <= (^(fi134)))}
                                                    if ((!v2) && (v5 == v5)) {
                                                        v132 = v132
                                                        var v139 string = "foo bar"
                                                        if (!true) {
                                                            var v140 float32 = 1e10
                                                            v140 = v140
                                                            _ = v140
                                                        }
                                                        _ = v139
                                                    } else {
                                                        v1 = (v135 - ((v135 + float64(0.0)) - v135))
                                                        fmt.Println(v2, v1)
                                                        var v141 uint = 255
                                                        _ = v141
                                                    }
                                                    var v142 int32 = 439
                                                    _ = v138
                                                    _ = v142
                                                } else {
                                                    if v2 {
                                                        if v2 {
                                                            var v143 bool = v2
                                                            _ = v143
                                                        }
                                                    }
                                                    fmt.Println(v135, v3, v5)
                                                }
                                                var v144 bool = true
                                                _ = v144
                                            } else {
                                                var v145 uint32 = 1
                                                v146 := GenS9{F0: -830}
                                                v1 += float64(1e10)
                                                _ = v145
                                                _ = v146
                                            }
                                        } else {
                                            var v147 uint64 = 0
                                            var v148 uint8 = v18
                                            _ = v147
                                            _ = v148
                                        }
                                    }
                                }
                            }
                            var v149 uint = 255
                            v150 := GenS3{F0: v2}
                            var v151 int = (^(fi134))
                            _ = v149
                            _ = v150
                            _ = v151
                        }
                    }
                    switch ("" + "") {
                    case "":
                        fmt.Println(v2)
                    case "\n\t":
                        v152 := GenS5{F0: 32767, F1: 42, F2: "foo bar"}
                        _ = v152
                    }
                } else {
                    v153 := GenS4{F0: 128}
                    fmt.Println(v3, v135)
                    fmt.Println(v132)
                    v18 -= v18
                    _ = v153
                }
                if (!(true || v2)) {
                    v132 += v132
                    fmt.Println(fi133, v132, v5)
                } else {
                    var v154 string = ""
                    if len(v4) > 0 {
                        v4[0] = 0
                    }
                    _ = v154
                }
                _ = fi134
                _ = v135
            }
            v155 := GenS10{F0: genFn1((v3 / ((31) | 1)), genFn1(42, v5, -1), 1), F1: v5, F2: 384}
            _ = fi133
            _ = v155
        }
    } else {
        var v156 uint64 = 0
        _ = v156
    }
    var v157 int32 = -16
    v5 += 379
    for fi158 := 0; fi158 < 14; fi158++ {
        var v159 int64 = 9223372036854775807
        v157 = (v157 + v157)
        _ = fi158
        _ = v159
    }
    for fi160 := 0; fi160 < 15; fi160++ {
        if (!v2) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v161 := GenS11{F0: ("\n\t" + "hello"), F1: 42, F2: (-(v157))}
            _ = v161
        }
        _ = fi160
    }
    fmt.Println("v1:", v1)
    fmt.Println("v132:", v132)
    fmt.Println("v157:", v157)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v126
    _ = v132
    _ = v157
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
