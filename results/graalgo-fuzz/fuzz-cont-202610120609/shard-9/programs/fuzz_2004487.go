package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint16; F1 int32; F2 int8 }
type GenS10 struct { F0 int16; F1 bool; F2 bool }
func genFn0(p0 string, p1 string) string {
    if (float64(float32((-2.25 / (0.1 + 0.1)))) < 1.5) {
        p0 = "\n\t"
        for fi4 := 0; fi4 < 20; fi4++ {
            p1 = p0
            for fi5 := 0; fi5 < 5; fi5++ {
                p1 = (p0 + ((p0 + "hello") + p0))
                switch (fi5 + -329) {
                case 32767:
                    p0 = p0
                case 4294967295:
                    var v6 uint64 = 4294967295
                    _ = v6
                case -1:
                    p1 = (p0 + p1)
                }
                _ = fi5
            }
            v7 := []int{(fi4 | fi4)}
            _ = fi4
            _ = v7
        }
        var v8 float64 = (float64((1.5 * 3.14159)) * float64(3.14159))
        _ = v8
    } else {
        var v9 bool = (21 != 0)
        p1 = "foo bar"
        _ = v9
    }
    var v10 float64 = (((-2.25 - (-2.25 * float64(0.0))) * -2.25) * 1e10)
    var v11 bool = ((true && (!(((1.5 * ((float64((1.5 * 0.1)) / 1e10) * ((v10 - v10) + v10))) == v10) && false))) || true)
    var v12 bool = (28 == 1)
    if (950 <= 886) {
        for fi13 := 0; fi13 < 9; fi13++ {
            if false {
                switch (^(fi13)) {
                case 112:
                    v11 = false
                case -211:
                    for fi14 := 0; fi14 < 11; fi14++ {
                        p1 = ((((p1 + p0) + ("" + (p1 + (("hello" + ("" + p0)) + p0)))) + p1) + (p1 + ((((p1 + p1) + (p0 + p0)) + p1) + p0)))
                        v15 := []int{(fi14 / (((fi14 << ((78) & 63))) | 1)), (^(fi13)), fi14}
                        fmt.Println(p1, v12)
                        _ = fi14
                        _ = v15
                    }
                default:
                    for fi16 := 0; fi16 < 18; fi16++ {
                        v11 = (1 >= -630)
                        if ((42 == 0) && true) {
                            if (-1 <= -688) {
                                var v17 int = (fi16 % ((fi13) | 1))
                                v17 -= fi16
                                fmt.Println(v12)
                                _ = v17
                            }
                        }
                        v18 := []int{(fi16 % (((fi13 & (fi13 * fi16))) | 1)), fi13, fi16}
                        _ = fi16
                        _ = v18
                    }
                }
            } else {
                for fi19 := 0; fi19 < 9; fi19++ {
                    fmt.Println(v11, v11)
                    if v12 {
                        switch 460 {
                        case -9223372036854775808:
                            var v20 int32 = -988
                            _ = v20
                        case 9223372036854775807:
                            v21 := GenS1{F0: 100}
                            _ = v21
                        case -361:
                            var v22 string = (p1 + p0)
                            _ = v22
                        default:
                            p0 = p0
                        }
                        p0 = "foo bar"
                        v11 = (43 <= 21)
                    } else {
                        var v23 string = (p0 + p0)
                        v24 := []int{(fi13 >> ((67) & 63)), (fi13 ^ fi19)}
                        if (128 == 255) {
                            v11 = (-1 == 9223372036854775807)
                            v10 = 3.14159
                        }
                        _ = v23
                        _ = v24
                    }
                    _ = fi19
                }
                if (!v11) {
                    p0 = p0
                    v10 = 0.0
                    v25 := GenS2{F0: 2, F1: (!v11), F2: 255}
                    _ = v25
                } else {
                    v26 := []int{}
                    var v27 string = "\n\t"
                    _ = v26
                    _ = v27
                }
                v28 := []int{fi13, fi13}
                _ = v28
            }
            v12 = (float32((v10 * v10)) >= 1e10)
            _ = fi13
        }
        v10 -= ((0.1 + (float64((1e10 - (0.1 + 1.5))) + 0.0)) * v10)
        if ((!v11) || v12) {
            for fi29 := 0; fi29 < 3; fi29++ {
                v10 += float64((3.14159 / ((0.1 - 3.14159) + (((1.5 / 3.14159) / (3.14159 - 0.0)) + 1.5))))
                fmt.Println(p1)
                _ = fi29
            }
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return ("hello" + "")
}

func main() {
    var v1 uint8 = 255
    var v2 uint16 = 1
    v3 := []uint32{2, 4294967295, 1, 2147483647}
    var v30 uint8 = v1
    for fi31 := 0; fi31 < 10; fi31++ {
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 uint8 = (v1 << ((4294967295) & 63))
            _ = fi32
            _ = v33
        }
        var v34 int = (fi31 & (fi31 << ((0) & 63)))
        _ = fi31
        _ = v34
    }
    fmt.Println(v2, v30, v1)
    v30 -= v1
    var v35 int = -989
    var v36 string = "foo bar"
    for fi37 := 0; fi37 < 16; fi37++ {
        var v38 uint16 = 1
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi39 := 0; fi39 < 20; fi39++ {
            v40 := GenS3{F0: 0}
            v41 := GenS4{F0: 42, F1: (v30 * v30)}
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = fi37
        _ = v38
    }
    v42 := GenS2{F0: v2, F1: (true && true), F2: v1}
    for fi43 := 0; fi43 < 14; fi43++ {
        v44 := GenS2{F0: 0, F1: false, F2: (^(v30))}
        if len(v3) > 0 {
            v3[0] = 111
        }
        _ = fi43
        _ = v44
    }
    fmt.Println(v1)
    switch v36 {
    case "hello":
        switch v35 {
        case 1:
            switch 32767 {
            case -9223372036854775808:
                if (!false) {
                    fmt.Println(v2, v35, v30)
                    var v45 int32 = -281
                    v35 -= v35
                    _ = v45
                }
            case 9223372036854775807:
                var v46 uint32 = 2
                _ = v46
            default:
                if false {
                    _ = len(v3)
                    for fi47 := 0; fi47 < 4; fi47++ {
                        _ = len(v3)
                        var v48 uint16 = (v2 << ((42) & 63))
                        v48 -= (v48 << ((255) & 63))
                        _ = fi47
                        _ = v48
                    }
                    if (!(660 == 0)) {
                        v49 := GenS1{F0: -1}
                        _ = v49
                    }
                }
            }
        case -1:
            var v50 float64 = 1e10
            _ = v50
        default:
            v30 = (v1 >> ((255) & 63))
        }
    case "a":
        v3 = append(v3, 278)
    }
    if false {
        if (1 <= 2) {
            switch ("hello" + "a") {
            case "a":
                v51 := GenS5{F0: -2147483648, F1: 42}
                _ = v51
            case "\n\t":
                switch v35 {
                case 42:
                    if false {
                        for fi52 := 0; fi52 < 6; fi52++ {
                            if len(v3) > 0 {
                                v3[0] = 1
                            }
                            _ = fi52
                        }
                        if false {
                            fmt.Println(v36, v2)
                            v3 = append(v3, 0)
                            switch v36 {
                            case "foo bar":
                                switch v36 {
                                case "a":
                                    v35 -= (v35 / ((-1) | 1))
                                case "hello":
                                    v3 = append(v3, 485)
                                }
                            case "a":
                                for fi53 := 0; fi53 < 13; fi53++ {
                                    if ((float64(3.14159) + 0.0) >= -2.25) {
                                        if (!true) {
                                            v54 := GenS4{F0: 1, F1: v1}
                                            var v55 string = ""
                                            v56 := GenS4{F0: 9223372036854775807, F1: v1}
                                            _ = v54
                                            _ = v55
                                            _ = v56
                                        }
                                        switch (v36 + "a") {
                                        case "hello":
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v57 int64 = -391
                                                _ = v57
                                            case "\n\t":
                                                if false {
                                                    switch v35 {
                                                    case -586:
                                                        var v58 int32 = 890
                                                        _ = v58
                                                    case 2147483647:
                                                        v3 = append(v3, 70)
                                                    }
                                                    fmt.Println(fi53)
                                                }
                                            case "hello":
                                                if false {
                                                    switch v36 {
                                                    case "\n\t":
                                                        v3 = append(v3, 246)
                                                    case "foo bar":
                                                        v30 = v1
                                                    }
                                                    var v59 int8 = 127
                                                    _ = v59
                                                }
                                            default:
                                                _ = len(v3)
                                            }
                                        case "a":
                                            var v60 float64 = 3.14159
                                            _ = v60
                                        case "":
                                            v35 -= v35
                                        default:
                                            fmt.Println(v35, v35, fi53)
                                        }
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    v1 += (^(v30))
                                    _ = fi53
                                }
                            }
                        } else {
                            if (v2 != v2) {
                                v30 += 185
                            }
                            v2 = v2
                        }
                        for fi61 := 0; fi61 < 15; fi61++ {
                            var v62 int64 = 1
                            switch (v35 + (fi61 % ((v35) | 1))) {
                            case 127:
                                var v63 int8 = -128
                                _ = v63
                            }
                            _ = fi61
                            _ = v62
                        }
                    }
                case 1:
                    var v64 float32 = 0.1
                    _ = v64
                case 624:
                    v65 := GenS6{F0: v36, F1: (^(v2))}
                    _ = v65
                }
            case "foo bar":
                var v66 uint64 = 49
                _ = v66
            }
            if false {
                if len(v3) > 0 {
                    v3[0] = 65535
                }
                _ = len(v3)
                if (v2 <= uint16(v1)) {
                    v67 := GenS4{F0: -1, F1: v1}
                    var v68 uint = 768
                    _ = v67
                    _ = v68
                }
            }
        }
    } else {
        v36 = "hello"
        v36 = v36
        if (0 > 42) {
            for fi69 := 0; fi69 < 2; fi69++ {
                v30 = (v30 & (v30 / ((v1) | 1)))
                _ = fi69
            }
            for fi70 := 0; fi70 < 16; fi70++ {
                switch (fi70 | fi70) {
                case 9223372036854775807:
                    v35 = fi70
                default:
                    switch v36 {
                    case "":
                        for fi71 := 0; fi71 < 17; fi71++ {
                            if (1 == 91) {
                                v30 = (v30 | 1)
                            } else {
                                var v72 float64 = float64(-2.25)
                                v36 = v36
                                _ = v72
                            }
                            if (46 != 127) {
                                switch (fi71 + v35) {
                                case 32768:
                                    var v73 uint8 = (^(v1))
                                    if false {
                                        var v74 uint64 = 18446744073709551615
                                        fmt.Println(v30)
                                        _ = v74
                                    } else {
                                        switch (fi71 & (fi70 ^ (fi71 / ((9223372036854775807) | 1)))) {
                                        case -938:
                                            fmt.Println(v1)
                                        case 558:
                                            if true {
                                                var v75 uint8 = (v30 - v73)
                                                _ = v75
                                            } else {
                                                switch v36 {
                                                case "hello":
                                                    v76 := GenS1{F0: -108}
                                                    _ = v76
                                                case "a":
                                                    fmt.Println(v30, v2, v30)
                                                case "\n\t":
                                                    v77 := GenS2{F0: v2, F1: true, F2: v30}
                                                    _ = v77
                                                default:
                                                    switch v36 {
                                                    case "foo bar":
                                                        fmt.Println(v30, v2)
                                                    case "a":
                                                        v73 += uint8(v2)
                                                    }
                                                }
                                            }
                                        default:
                                            if (!(227 > uint(v2))) {
                                                var v78 int64 = 1
                                                switch (v36 + "a") {
                                                case "a":
                                                    v3 = append(v3, 0)
                                                }
                                                _ = v78
                                            }
                                        }
                                        switch v36 {
                                        case "":
                                            if (false && (!true)) {
                                                v79 := GenS4{F0: -283, F1: (v73 << ((42) & 63))}
                                                if ((2 > 4294967295) || false) {
                                                    fmt.Println(v2, v73, v36)
                                                    var v80 float32 = 3.14159
                                                    v80 = -2.25
                                                    _ = v80
                                                } else {
                                                    var v81 float64 = 1.5
                                                    _ = v81
                                                }
                                                var v82 uint8 = 55
                                                _ = v79
                                                _ = v82
                                            }
                                        case "hello":
                                            if true {
                                                var v83 string = "foo bar"
                                                var v84 uint16 = (v2 - v2)
                                                _ = v83
                                                _ = v84
                                            }
                                        }
                                    }
                                    v1 = (v73 / ((225) | 1))
                                    _ = v73
                                }
                                switch fi70 {
                                case -772:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v85 := GenS7{F0: -539}
                                    fmt.Println(fi71, v36)
                                    _ = v85
                                default:
                                    var v86 float64 = (-2.25 * 0.0)
                                    _ = v86
                                }
                                fmt.Println(v1)
                            }
                            v2 = (v2 + (v2 + (v2 % (((v2 << ((526) & 63))) | 1))))
                            _ = fi71
                        }
                    case "\n\t":
                        v30 += v30
                    default:
                        v30 = (v30 % ((v30) | 1))
                    }
                }
                _ = fi70
            }
        } else {
            v3 = append(v3, 2)
            if true {
                v35 = 42
                var v87 uint16 = (v2 % ((v2) | 1))
                for fi88 := 0; fi88 < 8; fi88++ {
                    for fi89 := 0; fi89 < 19; fi89++ {
                        var v90 int16 = 40
                        var v91 bool = (0 < 657)
                        _ = fi89
                        _ = v90
                        _ = v91
                    }
                    var v92 string = ("foo bar" + "hello")
                    if true {
                        if false {
                            v3 = append(v3, 65535)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            switch ((v36 + "\n\t") + "a") {
                            case "hello":
                                v30 = (v30 >> ((281) & 63))
                            case "":
                                v3 = append(v3, 856)
                            }
                            if true {
                                fmt.Println(fi88, v87, fi88)
                                for fi93 := 0; fi93 < 8; fi93++ {
                                    var v94 int = (fi88 & v35)
                                    var v95 bool = false
                                    v3 = append(v3, 1)
                                    _ = fi93
                                    _ = v94
                                    _ = v95
                                }
                            } else {
                                var v96 int8 = 1
                                _ = v96
                            }
                            if true {
                                if true {
                                    var v97 int8 = -53
                                    v98 := GenS8{F0: uint64(v30)}
                                    if ((v97 / ((v97) | 1)) == (v97 / ((v97) | 1))) {
                                        if false {
                                            for fi99 := 0; fi99 < 7; fi99++ {
                                                if true {
                                                    if ((v87 < 65535) || (839 < uint32(v1))) {
                                                        v97 = v97
                                                        if (!(1.5 >= 1.5)) {
                                                            v1 -= (v1 / (((v1 | (v1 & (^(v1))))) | 1))
                                                        }
                                                        v100 := GenS6{F0: v36, F1: (v87 & v87)}
                                                        _ = v100
                                                    }
                                                }
                                                if ((true || ((1.5 / 1.0) <= 1e10)) && true) {
                                                    v92 = ("hello" + (v36 + "\n\t"))
                                                    switch -9223372036854775808 {
                                                    case 2147483647:
                                                        v101 := GenS1{F0: (v97 ^ (-(v97)))}
                                                        _ = v101
                                                    case 128:
                                                        var v102 uint = 476
                                                        _ = v102
                                                    }
                                                    v36 = "\n\t"
                                                }
                                                v1 = (v30 & (^(v1)))
                                                if (!true) {
                                                    var v103 int = fi88
                                                    _ = v103
                                                } else {
                                                    if (4294967295 >= 0) {
                                                        switch v35 {
                                                        case -228:
                                                            fmt.Println(v35, v30)
                                                        case -436:
                                                            var v104 string = "foo bar"
                                                            _ = v104
                                                        default:
                                                            if false {
                                                                var v105 uint32 = 133
                                                                _ = v105
                                                            }
                                                        }
                                                    }
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v106 uint8 = v1
                                                _ = fi99
                                                _ = v106
                                            }
                                            v107 := GenS9{F0: (v2 ^ (v2 << ((65535) & 63))), F1: 42, F2: v97}
                                            _ = v107
                                        }
                                        var v108 string = v92
                                        switch int(v30) {
                                        case -323:
                                            var v109 uint8 = v30
                                            _ = v109
                                        case 4294967295:
                                            var v110 int32 = -565
                                            _ = v110
                                        case -1:
                                            fmt.Println(v97)
                                        }
                                        _ = v108
                                    } else {
                                        v1 = v30
                                    }
                                    _ = v97
                                    _ = v98
                                } else {
                                    if false {
                                        v2 = (v87 % (((v87 << ((497) & 63))) | 1))
                                        v87 = v2
                                    }
                                }
                                switch v36 {
                                case "\n\t":
                                    for fi111 := 0; fi111 < 13; fi111++ {
                                        if false {
                                            var v112 int32 = 1
                                            if len(v3) > 0 {
                                                v3[0] = 2
                                            }
                                            _ = v112
                                        }
                                        fmt.Println(v35)
                                        if len(v3) > 0 {
                                            v3[0] = 1
                                        }
                                        _ = fi111
                                    }
                                default:
                                    if (-1 >= 1000) {
                                        var v113 float64 = 0.1
                                        _ = v113
                                    } else {
                                        v1 += 77
                                        for fi114 := 0; fi114 < 16; fi114++ {
                                            v115 := GenS10{F0: -32768, F1: true, F2: ((!(1 > -128)) || (float32(1.5) >= float32(float64(float32((((float64(3.14159) - 0.1) / (((1.5 * 1.5) - float64(0.0)) * (float64(0.1) - 1.5))) + 0.0))))))}
                                            _ = fi114
                                            _ = v115
                                        }
                                        if false {
                                            v87 += v2
                                        } else {
                                            v116 := GenS6{F0: (("foo bar" + v36) + v36), F1: 1}
                                            _ = v116
                                        }
                                    }
                                }
                            }
                        }
                        var v117 int8 = 55
                        var v118 bool = false
                        _ = v117
                        _ = v118
                    }
                    _ = fi88
                    _ = v92
                }
                _ = v87
            } else {
                v119 := GenS8{F0: 589}
                switch (v36 + (((v36 + "foo bar") + "hello") + (v36 + v36))) {
                case "\n\t":
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                case "a":
                    if ((v1 >> ((42) & 63)) >= 42) {
                        var v120 int64 = 0
                        _ = v120
                    }
                default:
                    var v121 float32 = (3.14159 + (0.0 / 3.14159))
                    _ = v121
                }
                _ = v119
            }
            fmt.Println(v2, v2, v1)
        }
    }
    var v122 int16 = 70
    var v123 uint16 = (v2 & (v2 + v2))
    fmt.Println("v1:", v1)
    fmt.Println("v122:", v122)
    fmt.Println("v123:", v123)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v122
    _ = v123
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v42
}
