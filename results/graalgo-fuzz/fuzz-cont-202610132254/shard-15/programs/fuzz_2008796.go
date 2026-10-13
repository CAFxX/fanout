package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint; F1 uint64; F2 string }
type GenS3 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS4 struct { F0 int64; F1 int; F2 bool }
type GenS5 struct { F0 string; F1 int; F2 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16; F1 uint32 }
type GenS8 struct { F0 int32; F1 int; F2 bool }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint64; F1 int16; F2 uint }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 bool; F1 int16 }
type GenS14 struct { F0 int8; F1 uint; F2 int8 }
type GenS15 struct { F0 uint64; F1 int64 }
type GenS16 struct { F0 string; F1 int16; F2 uint32 }
func genFn0(p0 uint32, p1 uint8, p2 uint16) int16 {
    v4 := []int{}
    var v5 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1000
}
func genFn1(p0 uint64, p1 int8, p2 int16) uint32 {
    if true {
        switch 127 {
        case 1:
            switch -600 {
            case 2147483647:
                v6 := GenS1{F0: 65535, F1: p1}
                _ = v6
            default:
                if (false || true) {
                    if ((true && ("foo bar" > "hello")) && (!false)) {
                        p1 = p1
                        if false {
                            v7 := []int{467}
                            _ = v7
                        }
                        v8 := GenS1{F0: 796, F1: p1}
                        _ = v8
                    }
                    v9 := GenS2{F0: 234, F1: p0, F2: "a"}
                    _ = v9
                } else {
                    if false {
                        for fi10 := 0; fi10 < 3; fi10++ {
                            if true {
                                fmt.Println(p1, fi10, p1)
                            }
                            _ = fi10
                        }
                        v11 := []int{9223372036854775807, 0, 127}
                        var v12 int8 = (p1 & p1)
                        _ = v11
                        _ = v12
                    } else {
                        fmt.Println(p2, p2, p1)
                    }
                    p2 = (p2 >> ((255) & 63))
                    var v13 string = "hello"
                    _ = v13
                }
            }
        case -1:
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 int8 = (^(p1))
                fmt.Println(v15, fi14)
                p1 += (-(v15))
                _ = fi14
                _ = v15
            }
        default:
            switch 42 {
            case 2147483647:
                v16 := GenS3{F0: (p0 << ((2) & 63)), F1: p2, F2: 420}
                _ = v16
            default:
                for fi17 := 0; fi17 < 20; fi17++ {
                    switch ("hello" + "") {
                    case "foo bar":
                        p1 -= -1
                    case "a":
                        var v18 uint8 = 46
                        _ = v18
                    default:
                        v19 := GenS1{F0: 2, F1: p1}
                        _ = v19
                    }
                    _ = fi17
                }
            }
        }
    } else {
        p0 = (p0 % ((p0) | 1))
        if (!true) {
            v20 := GenS3{F0: p0, F1: 0, F2: 1}
            _ = v20
        }
    }
    v21 := GenS1{F0: 1, F1: p1}
    v22 := GenS3{F0: (p0 >> ((65535) & 63)), F1: p2, F2: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return 2147483648
}

func main() {
    v1 := []string{("foo bar" + "a")}
    var v2 int = -634
    var v3 int16 = 32767
    v23 := GenS1{F0: 65535, F1: -1}
    if (1 <= 847) {
        var v24 int32 = 0
        if (1 >= 1) {
            var v25 string = ""
            switch "hello" {
            case "":
                if (false || (407 < 1)) {
                    fmt.Println(v2, v3, v3)
                    if (18446744073709551615 < 574) {
                        var v26 uint16 = 65535
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi27 := 0; fi27 < 15; fi27++ {
                            v28 := GenS1{F0: 32767, F1: -34}
                            v25 = v25
                            _ = fi27
                            _ = v28
                        }
                        _ = v26
                    } else {
                        v3 = v3
                        if true {
                            switch v2 {
                            case 128:
                                var v29 int8 = -128
                                _ = v29
                            case 1:
                                switch "\n\t" {
                                case "foo bar":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            case 983:
                                v24 = v24
                            }
                        } else {
                            for fi30 := 0; fi30 < 17; fi30++ {
                                for fi31 := 0; fi31 < 18; fi31++ {
                                    var v32 int16 = (v3 / (((^(v3))) | 1))
                                    v32 = genFn0(65535, 200, 0)
                                    v33 := GenS1{F0: 42, F1: 35}
                                    _ = len(v1)
                                    _ = fi31
                                    _ = v32
                                    _ = v33
                                }
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                                v24 -= (v24 + (-(v24)))
                                _ = fi30
                            }
                        }
                    }
                } else {
                    if false {
                        _ = len(v1)
                        switch (^(v2)) {
                        case -35:
                            if (!true) {
                                v24 = v24
                            }
                        case -1:
                            v34 := GenS2{F0: 65535, F1: 42, F2: "hello"}
                            _ = v34
                        case 1:
                            var v35 bool = true
                            _ = v35
                        default:
                            if true {
                                var v36 float64 = (3.14159 / (0.1 / 1.0))
                                var v37 int64 = 9223372036854775807
                                _ = v36
                                _ = v37
                            } else {
                                if (!false) {
                                    v3 = v3
                                    if (-1 >= -1) {
                                        for fi38 := 0; fi38 < 11; fi38++ {
                                            v24 += v24
                                            var v39 float64 = (0.0 - 0.1)
                                            if len(v1) > 0 {
                                                v1[0] = v25
                                            }
                                            _ = fi38
                                            _ = v39
                                        }
                                    } else {
                                        v25 = ("" + v25)
                                        v40 := GenS3{F0: 0, F1: genFn0(315, 25, 42), F2: 442}
                                        _ = v40
                                    }
                                    var v41 int8 = int8(v2)
                                    _ = v41
                                } else {
                                    if true {
                                        _ = len(v1)
                                        v2 = (v2 + v2)
                                    }
                                }
                                if (!true) {
                                    v3 += (v3 * (v3 / ((v3) | 1)))
                                    fmt.Println(v24, v3)
                                } else {
                                    fmt.Println(v25, v24)
                                }
                                for fi42 := 0; fi42 < 15; fi42++ {
                                    if ((!((!((65535 > 2147483648) && false)) || false)) || false) {
                                        switch -1 {
                                        case -9223372036854775808:
                                            fmt.Println(fi42, v24)
                                        case 32767:
                                            v25 = (("a" + v25) + "")
                                        case 2147483647:
                                            var v43 int16 = (^(v3))
                                            _ = v43
                                        default:
                                            var v44 uint64 = 367
                                            _ = v44
                                        }
                                    }
                                    v25 = (("" + (v25 + v25)) + v25)
                                    _ = fi42
                                }
                            }
                        }
                    } else {
                        v1 = append(v1, v25)
                    }
                    var v45 uint16 = 32767
                    if ((611 < uint32(v45)) && ((v45 >> ((1) & 63)) > 411)) {
                        v1 = append(v1, ((v25 + ((((((v25 + v25) + v25) + (("foo bar" + v25) + "")) + v25) + (v25 + "hello")) + "foo bar")) + v25))
                        fmt.Println(v2, v45)
                        switch "hello" {
                        case "":
                            fmt.Println(v2, v45, v25)
                        case "hello":
                            for fi46 := 0; fi46 < 6; fi46++ {
                                v47 := GenS4{F0: int64(v24), F1: (-(v2)), F2: false}
                                v48 := GenS1{F0: 0, F1: 108}
                                _ = fi46
                                _ = v47
                                _ = v48
                            }
                        case "foo bar":
                            if (((v25 + "foo bar") + ((v25 + "foo bar") + "a")) >= "") {
                                v2 -= v2
                            } else {
                                switch ("" + ((v25 + v25) + v25)) {
                                case "a":
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                case "":
                                    v1 = append(v1, "foo bar")
                                case "hello":
                                    var v49 int = v2
                                    _ = v49
                                default:
                                    if (v3 > genFn0(genFn1(891, 72, (v3 ^ (v3 % (((v3 | (v3 >> ((18446744073709551615) & 63)))) | 1)))), uint8(v45), (v45 << ((65535) & 63)))) {
                                        if (65 < 1) {
                                            for fi50 := 0; fi50 < 20; fi50++ {
                                                var v51 int8 = 127
                                                for fi52 := 0; fi52 < 20; fi52++ {
                                                    if len(v1) > 0 {
                                                        v1[0] = v25
                                                    }
                                                    var v53 uint = 2
                                                    _ = fi52
                                                    _ = v53
                                                }
                                                _ = fi50
                                                _ = v51
                                            }
                                            var v54 float32 = 0.0
                                            if false {
                                                v55 := GenS5{F0: (("a" + v25) + v25), F1: -1, F2: 158}
                                                _ = v55
                                            } else {
                                                var v56 uint8 = 255
                                                var v57 int = v2
                                                if ((v56 >= (^(v56))) || ((float32((0.1 / 1.5)) + float32((float64((v54 + v54)) + 1e10))) <= (((v54 + 1e10) * (3.14159 * (v54 - v54))) / float32((-2.25 * float64(-2.25)))))) {
                                                    v58 := GenS6{F0: (v56 & (v56 << ((4294967295) & 63)))}
                                                    v1 = append(v1, v25)
                                                    var v59 int8 = 1
                                                    _ = v58
                                                    _ = v59
                                                }
                                                _ = v56
                                                _ = v57
                                            }
                                            _ = v54
                                        } else {
                                            v60 := GenS7{F0: 1, F1: 125}
                                            var v61 int = v2
                                            _ = v60
                                            _ = v61
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            v25 = v25
                                        case "hello":
                                            for fi62 := 0; fi62 < 6; fi62++ {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v63 := GenS1{F0: v45, F1: 0}
                                                _ = fi62
                                                _ = v63
                                            }
                                        default:
                                            v24 = v24
                                        }
                                    }
                                }
                                var v64 uint = 394
                                if len(v1) > 0 {
                                    v1[0] = (v25 + v25)
                                }
                                _ = v64
                            }
                        default:
                            for fi65 := 0; fi65 < 5; fi65++ {
                                for fi66 := 0; fi66 < 3; fi66++ {
                                    if false {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        fmt.Println(v2, fi65)
                                        var v67 int = v2
                                        _ = v67
                                    } else {
                                        v68 := GenS8{F0: v24, F1: (v2 & v2), F2: true}
                                        var v69 string = ""
                                        _ = v68
                                        _ = v69
                                    }
                                    _ = fi66
                                }
                                _ = fi65
                            }
                        }
                    } else {
                        for fi70 := 0; fi70 < 18; fi70++ {
                            v71 := GenS9{F0: v2}
                            v45 = 65535
                            var v72 int8 = 42
                            _ = fi70
                            _ = v71
                            _ = v72
                        }
                        switch ("" + "") {
                        case "a":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "foo bar":
                            for fi73 := 0; fi73 < 8; fi73++ {
                                _ = len(v1)
                                if (v25 <= (v25 + ("\n\t" + ((("\n\t" + ("\n\t" + (v25 + ""))) + v25) + (v25 + v25))))) {
                                    v24 = (v24 * (v24 + v24))
                                }
                                _ = fi73
                            }
                        case "\n\t":
                            for fi74 := 0; fi74 < 17; fi74++ {
                                var v75 uint64 = 18446744073709551615
                                if false {
                                    fmt.Println(v45, v25)
                                }
                                if (genFn1(v75, 1, genFn0(1, 127, v45)) < 161) {
                                    for fi76 := 0; fi76 < 4; fi76++ {
                                        if false {
                                            v77 := GenS10{F0: -1}
                                            v24 = -2147483648
                                            var v78 uint8 = 127
                                            _ = v77
                                            _ = v78
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = v25
                                        }
                                        if (42 <= 255) {
                                            v79 := GenS11{F0: 42, F1: int16(v75), F2: 65535}
                                            var v80 float32 = 3.14159
                                            _ = v79
                                            _ = v80
                                        } else {
                                            v75 -= 1
                                        }
                                        _ = fi76
                                    }
                                    fmt.Println(v45, v25, v75)
                                    fmt.Println(v24, v24, v24)
                                }
                                _ = fi74
                                _ = v75
                            }
                        default:
                            if (1 > 1) {
                                v2 -= (-(v2))
                                v3 = (v3 >> ((18446744073709551615) & 63))
                                for fi81 := 0; fi81 < 3; fi81++ {
                                    for fi82 := 0; fi82 < 8; fi82++ {
                                        if ((1 > 220) || ((-1 < 100) || (false || false))) {
                                            v83 := GenS8{F0: v24, F1: fi82, F2: true}
                                            var v84 int32 = -2147483648
                                            switch fi81 {
                                            case 2147483647:
                                                _ = len(v1)
                                            default:
                                                if false {
                                                    v84 = (v24 + 42)
                                                }
                                            }
                                            if true {
                                                v3 -= -976
                                            }
                                            _ = v83
                                            _ = v84
                                        }
                                        _ = fi82
                                    }
                                    v85 := GenS12{F0: 1}
                                    _ = fi81
                                    _ = v85
                                }
                            }
                        }
                    }
                    _ = v45
                }
            case "hello":
                var v86 uint = 18446744073709551615
                _ = v86
            }
            var v87 float64 = (((1.5 + (float64(float32(float64(-2.25))) + float64((1e10 - 0.1)))) * ((((0.0 / (((float64((((((float32(0.1) * (float32(((3.14159 + 1.5) * (float64(float32(((0.1 + 0.0) + -2.25))) * 1.5))) / 1e10)) + 1e10) * (3.14159 / float32(1e10))) - (1.5 + ((1e10 * (1e10 + 0.1)) - float32(-2.25)))) - 3.14159)) / ((((-2.25 - 0.0) + -2.25) - 3.14159) * (3.14159 + ((3.14159 / (0.1 - 1e10)) + (float64((0.0 - ((3.14159 * 0.0) / 3.14159))) + float64(0.1)))))) - 0.0) / 1.0)) + 3.14159) * ((3.14159 + (((1e10 * -2.25) - float64(-2.25)) + 1.5)) + 0.0)) - (0.1 - 1e10))) / 1.0)
            _ = v25
            _ = v87
        }
        for fi88 := 0; fi88 < 12; fi88++ {
            fmt.Println(fi88, fi88)
            _ = fi88
        }
        _ = v24
    }
    if (153 >= -9223372036854775808) {
        for fi89 := 0; fi89 < 9; fi89++ {
            switch -83 {
            case -2147483648:
                var v90 int8 = 127
                _ = v90
            case 9223372036854775807:
                fmt.Println(v3, fi89, v2)
            default:
                v91 := GenS13{F0: false, F1: 32767}
                _ = v91
            }
            switch "foo bar" {
            case "":
                v3 -= genFn0(65535, 42, 255)
            case "foo bar":
                if true {
                    v92 := GenS6{F0: 127}
                    switch "\n\t" {
                    case "hello":
                        for fi93 := 0; fi93 < 7; fi93++ {
                            v94 := GenS14{F0: -1, F1: 997, F2: -128}
                            var v95 float64 = (3.14159 * 0.0)
                            v95 += ((v95 * v95) - (1.5 / float64(((float32(v95) * 1.5) * 0.0))))
                            v96 := GenS1{F0: 65535, F1: 0}
                            _ = fi93
                            _ = v94
                            _ = v95
                            _ = v96
                        }
                    default:
                        for fi97 := 0; fi97 < 13; fi97++ {
                            if false {
                                if len(v1) > 0 {
                                    v1[0] = "\n\t"
                                }
                                v98 := GenS15{F0: 0, F1: -1}
                                _ = v98
                            } else {
                                var v99 uint8 = 193
                                fmt.Println(fi97)
                                _ = v99
                            }
                            v2 = (fi89 * fi97)
                            switch v2 {
                            case -9223372036854775808:
                                var v100 uint16 = 1
                                _ = v100
                            case -470:
                                if ((v3 > (v3 % ((v3) | 1))) || (float32(0.0) < (float32(float64(1.5)) / 1.0))) {
                                    var v101 float32 = (1.5 - 0.0)
                                    if true {
                                        var v102 float64 = 1e10
                                        _ = v102
                                    }
                                    v101 = 1.5
                                    _ = v101
                                } else {
                                    v2 = fi89
                                    var v103 float32 = (((3.14159 - -2.25) / 1.5) - 1e10)
                                    _ = v103
                                }
                            }
                            _ = fi97
                        }
                    }
                    if true {
                        switch fi89 {
                        case 283:
                            if false {
                                var v104 float32 = 1e10
                                if len(v1) > 0 {
                                    v1[0] = ("a" + "hello")
                                }
                                _ = v104
                            }
                        default:
                            if (!(-9223372036854775808 <= -1)) {
                                v1 = append(v1, "hello")
                                v105 := GenS16{F0: "hello", F1: genFn0(477, 1, 0), F2: 65535}
                                _ = v105
                            }
                        }
                    }
                    _ = v92
                } else {
                    v106 := GenS15{F0: 4294967295, F1: 9223372036854775807}
                    v3 = 0
                    _ = v106
                }
            default:
                var v107 bool = true
                _ = v107
            }
            _ = fi89
        }
    }
    v3 -= (v3 / ((v3) | 1))
    var v108 uint = 2
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    for fi109 := 0; fi109 < 16; fi109++ {
        var v110 string = "\n\t"
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi109
        _ = v110
    }
    var v111 uint16 = 65535
    v112 := GenS10{F0: -9223372036854775808}
    var v113 int16 = -173
    for fi114 := 0; fi114 < 8; fi114++ {
        if true {
            v113 += 42
        }
        _ = fi114
    }
    for fi115 := 0; fi115 < 8; fi115++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v113, v2, v108)
        _ = fi115
    }
    for fi116 := 0; fi116 < 18; fi116++ {
        v111 += v111
        if true {
            v2 = (fi116 % ((128) | 1))
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
        }
        var v117 int64 = 9223372036854775807
        _ = fi116
        _ = v117
    }
    var v118 uint64 = 4294967295
    v113 -= (v3 | v3)
    var v119 string = ""
    if (v119 > (v119 + v119)) {
        var v120 int8 = int8(v111)
        _ = v120
    }
    var v121 uint16 = v111
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v108:", v108)
    fmt.Println("v111:", v111)
    fmt.Println("v113:", v113)
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v108
    _ = v111
    _ = v112
    _ = v113
    _ = v118
    _ = v119
    _ = v121
    _ = v2
    _ = v23
    _ = v3
}
