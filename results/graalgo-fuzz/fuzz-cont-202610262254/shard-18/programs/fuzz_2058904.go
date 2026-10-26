package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 uint16; F1 uint; F2 int32 }
type GenS8 struct { F0 uint32; F1 uint16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint32; F1 uint32; F2 uint }
type GenS11 struct { F0 int; F1 int64; F2 int }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint8 }
func genFn0(p0 int, p1 int64) uint8 {
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int64 = (p1 >> ((913) & 63))
        for fi9 := 0; fi9 < 19; fi9++ {
            p0 = 1
            var v10 string = "\n\t"
            if true {
                if false {
                    if (!true) {
                        v11 := []int{fi9, p0, fi9, 1}
                        fmt.Println(p1)
                        _ = v11
                    }
                } else {
                    v12 := GenS1{F0: 255, F1: "foo bar"}
                    _ = v12
                }
                var v13 int16 = 254
                var v14 uint16 = 264
                _ = v13
                _ = v14
            }
            _ = fi9
            _ = v10
        }
        v8 = (p1 & p1)
        _ = fi7
        _ = v8
    }
    p1 -= (p1 - p1)
    _ = p0
    _ = p1
    return 202
}

func main() {
    var v1 int = 1
    var v2 int32 = 598
    v3 := []uint32{0}
    var v4 bool = true
    var v5 int64 = -9223372036854775808
    var v6 uint8 = 104
    if true {
        if (v4 && (!(65535 <= 48))) {
            var v15 int64 = v5
            _ = v15
        } else {
            v4 = v4
            v6 -= 2
        }
    } else {
        if true {
            v1 -= 9223372036854775807
            switch int(v5) {
            case 32768:
                if len(v3) > 0 {
                    v3[0] = 65535
                }
            case 4294967295:
                var v16 int32 = v2
                _ = v16
            case 42:
                v17 := GenS2{F0: v5}
                _ = v17
            }
        } else {
            for fi18 := 0; fi18 < 20; fi18++ {
                v19 := GenS3{F0: 0, F1: 65535, F2: v6}
                var v20 int32 = 2147483647
                for fi21 := 0; fi21 < 11; fi21++ {
                    var v22 uint32 = 0
                    switch int(v22) {
                    case -9223372036854775808:
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(fi18, fi21)
                            v3 = append(v3, uint32(v6))
                        case "foo bar":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case 128:
                        if (!(!v4)) {
                            fmt.Println(v2, fi18, v1)
                            var v23 int = (fi18 + (fi18 << ((2) & 63)))
                            var v24 float64 = ((1.5 - 0.0) - 1.5)
                            _ = v23
                            _ = v24
                        } else {
                            v2 -= (v20 >> ((65535) & 63))
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                            v5 -= 1
                        }
                    case 4294967295:
                        _ = len(v3)
                    default:
                        var v25 bool = (-1 < 42)
                        _ = v25
                    }
                    _ = fi21
                    _ = v22
                }
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    }
    v26 := GenS3{F0: 32767, F1: 0, F2: v6}
    switch v1 {
    case 9223372036854775807:
        v27 := GenS2{F0: v5}
        _ = v27
    default:
        v1 += v1
    }
    if true {
        if true {
            switch v1 {
            case 2147483647:
                for fi28 := 0; fi28 < 5; fi28++ {
                    switch -2147483648 {
                    case 127:
                        for fi29 := 0; fi29 < 7; fi29++ {
                            if false {
                                if v4 {
                                    v5 += (v5 * (v5 % ((v5) | 1)))
                                    var v30 int32 = v2
                                    _ = v30
                                } else {
                                    var v31 uint8 = v6
                                    fmt.Println(fi29, v31)
                                    switch ("" + "\n\t") {
                                    case "hello":
                                        fmt.Println(v6, v31)
                                    case "a":
                                        var v32 int = fi28
                                        _ = v32
                                    default:
                                        v33 := GenS3{F0: 1, F1: 42, F2: (^(v6))}
                                        _ = v33
                                    }
                                    _ = v31
                                }
                                v34 := GenS4{F0: 2147483647}
                                if len(v3) > 0 {
                                    v3[0] = 471
                                }
                                _ = v34
                            } else {
                                var v35 int16 = -1
                                _ = v35
                            }
                            _ = fi29
                        }
                    case 42:
                        v6 += v6
                    case 2147483647:
                        for fi36 := 0; fi36 < 11; fi36++ {
                            var v37 int = fi36
                            _ = fi36
                            _ = v37
                        }
                    default:
                        if v4 {
                            v5 = (v5 & (v5 % (((-(v5))) | 1)))
                            var v38 uint = 613
                            _ = v38
                        }
                    }
                    _ = fi28
                }
            case 32767:
                if (v5 != -1) {
                    for fi39 := 0; fi39 < 2; fi39++ {
                        _ = len(v3)
                        v5 = -787
                        _ = fi39
                    }
                    if v4 {
                        switch "\n\t" {
                        case "foo bar":
                            var v40 int8 = 42
                            _ = v40
                        }
                        v41 := GenS5{F0: v5}
                        _ = v41
                    } else {
                        for fi42 := 0; fi42 < 13; fi42++ {
                            switch 864 {
                            case -211:
                                var v43 int16 = 0
                                _ = v43
                            case 378:
                                if (1 == 920) {
                                    if (3.14159 <= float32(-2.25)) {
                                        if v4 {
                                            v44 := GenS1{F0: 65535, F1: "\n\t"}
                                            v45 := GenS3{F0: 1000, F1: 4294967295, F2: (v6 % ((v6) | 1))}
                                            _ = v44
                                            _ = v45
                                        }
                                        var v46 float64 = (float64(0.1) / 0.1)
                                        if (float32(v46) == 3.14159) {
                                            if ("a" < ((("foo bar" + "\n\t") + "hello") + ("a" + ""))) {
                                                for fi47 := 0; fi47 < 11; fi47++ {
                                                    if len(v3) > 0 {
                                                        v3[0] = 0
                                                    }
                                                    _ = fi47
                                                }
                                                switch ("\n\t" + ("hello" + "\n\t")) {
                                                case "\n\t":
                                                    var v48 uint32 = 193
                                                    _ = v48
                                                case "a":
                                                    if true {
                                                        for fi49 := 0; fi49 < 18; fi49++ {
                                                            var v50 string = ""
                                                            var v51 uint8 = (v6 & v6)
                                                            _ = fi49
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    } else {
                                                        v3 = append(v3, 2147483647)
                                                        if (v46 == (v46 / float64(1e10))) {
                                                            var v52 uint64 = 189
                                                            for fi53 := 0; fi53 < 19; fi53++ {
                                                                v54 := GenS2{F0: (v5 & v5)}
                                                                switch ((("foo bar" + "a") + ("foo bar" + ("a" + "\n\t"))) + (("hello" + "") + ("" + ("hello" + "hello")))) {
                                                                case "":
                                                                    _ = len(v3)
                                                                    var v55 uint8 = v6
                                                                    switch ("foo bar" + (("\n\t" + ("\n\t" + "foo bar")) + "hello")) {
                                                                    case "a":
                                                                        v56 := GenS1{F0: 2, F1: (("foo bar" + "\n\t") + ("a" + "foo bar"))}
                                                                        _ = v56
                                                                    }
                                                                    _ = v55
                                                                case "foo bar":
                                                                    fmt.Println(v6)
                                                                case "a":
                                                                    v57 := GenS6{F0: 0, F1: 1000}
                                                                    _ = v57
                                                                default:
                                                                    if (!(1.5 > ((((float64(1.5) - v46) * v46) / (1.5 * v46)) * (1e10 * (v46 * ((0.0 * v46) * 0.0)))))) {
                                                                        v58 := GenS6{F0: 65535, F1: -1}
                                                                        _ = v58
                                                                    }
                                                                    _ = len(v3)
                                                                }
                                                                _ = fi53
                                                                _ = v54
                                                            }
                                                            v59 := GenS7{F0: 42, F1: 65535, F2: 2147483647}
                                                            _ = v52
                                                            _ = v59
                                                        } else {
                                                            var v60 uint32 = 2147483648
                                                            v60 = 803
                                                            for fi61 := 0; fi61 < 11; fi61++ {
                                                                v4 = (1 > 1000)
                                                                v5 -= int64(v6)
                                                                fmt.Println(v5)
                                                                v62 := GenS8{F0: v60, F1: 461}
                                                                var v63 int32 = (v2 * v2)
                                                                _ = fi61
                                                                _ = v62
                                                                _ = v63
                                                            }
                                                            _ = v60
                                                        }
                                                        switch 2147483647 {
                                                        case 0:
                                                            v46 = (((v46 / (v46 * (v46 / (1.5 * v46)))) - 0.0) * 3.14159)
                                                        case 9223372036854775807:
                                                            var v64 int8 = 100
                                                            _ = v64
                                                        case 32767:
                                                            var v65 bool = (!v4)
                                                            _ = v65
                                                        }
                                                    }
                                                default:
                                                    v5 = v5
                                                }
                                            } else {
                                                for fi66 := 0; fi66 < 2; fi66++ {
                                                    v67 := GenS4{F0: (v2 ^ 144)}
                                                    var v68 uint16 = 2
                                                    _ = fi66
                                                    _ = v67
                                                    _ = v68
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                for fi69 := 0; fi69 < 7; fi69++ {
                                                    if v4 {
                                                        var v70 float64 = (float64((float32(1.5) + ((-2.25 * (float32((float64(0.0) / float64((((-2.25 + 1e10) * 1e10) / (((1.5 / 0.1) - 3.14159) * 1.5))))) - 1e10)) + 0.1))) - v46)
                                                        if len(v3) > 0 {
                                                            v3[0] = 378
                                                        }
                                                        v6 = (v6 & (^(v6)))
                                                        _ = v70
                                                    } else {
                                                        v71 := GenS9{F0: v6}
                                                        var v72 int64 = v5
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                    var v73 int = (^(fi42))
                                                    _ = fi69
                                                    _ = v73
                                                }
                                            }
                                            v1 -= fi42
                                        }
                                        _ = v46
                                    }
                                }
                            case 4294967295:
                                v4 = (true && (0 >= 0))
                            }
                            v4 = (4294967295 != 746)
                            for fi74 := 0; fi74 < 15; fi74++ {
                                var v75 uint8 = (v6 / (((v6 | (^(v6)))) | 1))
                                v76 := GenS9{F0: (v6 ^ v75)}
                                _ = fi74
                                _ = v75
                                _ = v76
                            }
                            _ = fi42
                        }
                        var v77 uint32 = 908
                        _ = v77
                    }
                    for fi78 := 0; fi78 < 13; fi78++ {
                        if true {
                            var v79 bool = ((v5 <= (v5 | (v5 - v5))) || (((0.1 / 3.14159) * ((0.0 - float32(3.14159)) * 1e10)) == 1e10))
                            fmt.Println(v2, v5, v1)
                            _ = v79
                        } else {
                            if v4 {
                                if ((true && v4) || (1.5 > 0.0)) {
                                    var v80 int16 = 1
                                    var v81 int32 = (^(v2))
                                    _ = v80
                                    _ = v81
                                }
                                v2 = v2
                            }
                        }
                        _ = fi78
                    }
                } else {
                    v4 = (uint(v6) >= 4294967295)
                    if (42 == 39) {
                        fmt.Println(v1, v6)
                        for fi82 := 0; fi82 < 15; fi82++ {
                            for fi83 := 0; fi83 < 19; fi83++ {
                                var v84 int16 = int16(v6)
                                if false {
                                    if v4 {
                                        if (!true) {
                                            v2 = v2
                                        } else {
                                            v85 := GenS10{F0: 2147483647, F1: 2147483648, F2: 559}
                                            _ = v85
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = 411
                                        }
                                        fmt.Println(fi83, fi83, v2)
                                    }
                                    if (v4 || v4) {
                                        var v86 int8 = -128
                                        v5 += (-(v5))
                                        v5 -= (v5 - (v5 | -9223372036854775808))
                                        _ = v86
                                    }
                                    v2 -= v2
                                } else {
                                    if ((v84 | v84) <= (v84 & (v84 >> ((42) & 63)))) {
                                        v87 := GenS5{F0: v5}
                                        _ = v87
                                    }
                                }
                                v88 := GenS9{F0: 127}
                                _ = fi83
                                _ = v84
                                _ = v88
                            }
                            v5 = (-(v5))
                            var v89 int32 = (v2 | 0)
                            _ = fi82
                            _ = v89
                        }
                        fmt.Println(v4, v6, v2)
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    var v90 bool = ((0.0 * (((0.1 * -2.25) - (1.5 + 0.1)) - (0.1 - 1e10))) >= -2.25)
                    _ = v90
                }
            default:
                for fi91 := 0; fi91 < 4; fi91++ {
                    var v92 uint8 = v6
                    _ = fi91
                    _ = v92
                }
            }
        } else {
            if (!true) {
                fmt.Println(v4)
                _ = len(v3)
                if false {
                    fmt.Println(v5, v1, v1)
                }
            } else {
                var v93 uint = 4294967295
                _ = v93
            }
        }
        if (true && (!v4)) {
            for fi94 := 0; fi94 < 10; fi94++ {
                var v95 int8 = 17
                fmt.Println(v2)
                _ = fi94
                _ = v95
            }
            fmt.Println(v5)
        }
    } else {
        var v96 int64 = (v5 & v5)
        if (!((v6 % ((v6) | 1)) > v6)) {
            v97 := GenS11{F0: v1, F1: -1, F2: (v1 ^ 2147483647)}
            v4 = (1 < 42)
            _ = v97
        } else {
            v98 := GenS3{F0: 1, F1: 721, F2: (^(v6))}
            v4 = (((((v1 ^ v1) != (-(v1))) && (!(!v4))) || (0.1 >= float64(1e10))) || ((!((true && (float32(3.14159) < -2.25)) || v4)) || v4))
            _ = v98
        }
        var v99 uint8 = genFn0((v1 * (v1 << ((1) & 63))), (v5 / (((v5 / (((v5 >> ((541) & 63))) | 1))) | 1)))
        _ = v96
        _ = v99
    }
    for fi100 := 0; fi100 < 6; fi100++ {
        v5 = 923
        v3 = append(v3, 919)
        _ = fi100
    }
    for fi101 := 0; fi101 < 5; fi101++ {
        v4 = (65535 <= 42)
        _ = fi101
    }
    var v102 uint32 = 149
    for fi103 := 0; fi103 < 19; fi103++ {
        v4 = (1 > -41)
        _ = fi103
    }
    for fi104 := 0; fi104 < 19; fi104++ {
        var v105 float32 = (1.5 * 0.0)
        switch "" {
        case "a":
            if (v2 != (v2 | (^(v2)))) {
                v1 = fi104
                if true {
                    var v106 uint32 = (v102 % (((v102 * v102)) | 1))
                    _ = v106
                }
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        case "foo bar":
            var v107 uint16 = 65535
            _ = v107
        case "\n\t":
            v108 := GenS12{F0: v2}
            _ = v108
        }
        v6 = (v6 & (^(v6)))
        _ = fi104
        _ = v105
    }
    if false {
        _ = len(v3)
    }
    switch (^(v1)) {
    case 0:
        var v109 uint = 0
        _ = v109
    case -872:
        var v110 bool = v4
        _ = v110
    case 4294967295:
        var v111 int16 = 32767
        _ = v111
    default:
        if len(v3) > 0 {
            v3[0] = v102
        }
    }
    var v112 int64 = (v5 * v5)
    v113 := GenS5{F0: int64(v2)}
    if ((v4 && true) && (v5 != v5)) {
        var v114 uint8 = 0
        var v115 float64 = 1.5
        _ = v114
        _ = v115
    } else {
        for fi116 := 0; fi116 < 10; fi116++ {
            for fi117 := 0; fi117 < 2; fi117++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch 528 {
                case -9223372036854775808:
                    var v118 float64 = -2.25
                    _ = v118
                case -970:
                    if false {
                        var v119 uint8 = 2
                        fmt.Println(fi116)
                        var v120 float32 = (1.5 - float32((float64(3.14159) * 3.14159)))
                        switch (("\n\t" + "") + "a") {
                        case "foo bar":
                            v4 = (52 != 2)
                        case "a":
                            if (215 == 837) {
                                if v4 {
                                    var v121 int = (fi116 ^ (v1 / ((-130) | 1)))
                                    if (v4 || v4) {
                                        var v122 float64 = (1e10 * 0.0)
                                        v102 -= (v102 + v102)
                                        _ = v122
                                    }
                                    var v123 int16 = 42
                                    _ = v121
                                    _ = v123
                                }
                                var v124 uint8 = (v6 << ((1) & 63))
                                v102 -= (v102 ^ (^(v102)))
                                _ = v124
                            }
                        case "\n\t":
                            v6 += v119
                        default:
                            v112 = (v5 - v5)
                        }
                        _ = v119
                        _ = v120
                    }
                case 532:
                    v3 = append(v3, 4294967295)
                default:
                    fmt.Println(v2, v102)
                }
                switch 4294967295 {
                case -406:
                    v112 -= (^(v112))
                case -671:
                    v1 += fi116
                case 143:
                    if v4 {
                        v1 -= -611
                    }
                    switch ("hello" + ("foo bar" + "a")) {
                    case "a":
                        v3 = append(v3, v102)
                    case "hello":
                        v125 := GenS11{F0: -2147483648, F1: v5, F2: v1}
                        _ = v125
                    }
                default:
                    var v126 string = "hello"
                    _ = v126
                }
                _ = fi117
            }
            _ = fi116
        }
    }
    v127 := GenS13{F0: v6}
    var v128 uint8 = (^(v6))
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v112:", v112)
    fmt.Println("v128:", v128)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v102
    _ = v112
    _ = v113
    _ = v127
    _ = v128
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
