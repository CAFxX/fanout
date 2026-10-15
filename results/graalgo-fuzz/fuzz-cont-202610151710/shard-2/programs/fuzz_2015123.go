package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 int32; F1 uint }
type GenS7 struct { F0 uint16; F1 uint32 }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 string; F1 int8; F2 int8 }
type GenS11 struct { F0 string; F1 uint64 }
func genFn0(p0 int, p1 int, p2 int32) int8 {
    var v6 int8 = -1
    p1 = (^(p0))
    for fi7 := 0; fi7 < 4; fi7++ {
        if false {
            if ((true || (!true)) && ((!(1 < 18446744073709551615)) && false)) {
                var v8 uint64 = 1
                _ = v8
            }
            var v9 int64 = 42
            _ = v9
        }
        _ = fi7
    }
    if false {
        fmt.Println(v6, p2)
        v10 := []int{(-(p1)), (p1 + (p1 & p0)), -1}
        var v11 uint64 = 4294967295
        _ = v10
        _ = v11
    }
    v12 := []int{42}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    return -128
}

func main() {
    v1 := []uint32{0, 65535, 466, 2147483648}
    v2 := []uint32{}
    v3 := []uint32{715}
    var v4 string = (("" + "foo bar") + "foo bar")
    var v5 int32 = 42
    var v13 uint16 = 818
    if (false || (-83 != genFn0(32768, -1, v5))) {
        for fi14 := 0; fi14 < 9; fi14++ {
            switch fi14 {
            case 0:
                switch int(v5) {
                case -544:
                    for fi15 := 0; fi15 < 13; fi15++ {
                        var v16 int16 = -461
                        var v17 int64 = -1
                        switch 2147483647 {
                        case -63:
                            var v18 bool = true
                            _ = v18
                        case 127:
                            switch (v4 + v4) {
                            case "hello":
                                if true {
                                    v4 = (v4 + v4)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            case "\n\t":
                                v5 += v5
                            }
                        }
                        var v19 float32 = (((1.5 + (-2.25 + 1e10)) * -2.25) - 0.1)
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v19
                    }
                case -1:
                    v13 += (v13 * (^(v13)))
                case -2147483648:
                    v5 = (v5 % ((v5) | 1))
                default:
                    v13 = (^(v13))
                }
            case 32768:
                for fi20 := 0; fi20 < 13; fi20++ {
                    if (int32(fi14) < (^(v5))) {
                        v13 -= v13
                        var v21 int32 = v5
                        fmt.Println(v13, v5)
                        _ = v21
                    } else {
                        if true {
                            v2 = append(v2, 2)
                            if false {
                                v5 = v5
                                v5 = (v5 / (((v5 & v5)) | 1))
                            }
                        } else {
                            var v22 uint16 = 0
                            v5 -= v5
                            v13 = v22
                            _ = v22
                        }
                        var v23 uint = 255
                        var v24 int = 684
                        _ = v23
                        _ = v24
                    }
                    v4 = v4
                    if false {
                        v5 = v5
                    } else {
                        var v25 float64 = (3.14159 + 0.0)
                        _ = v25
                    }
                    switch (fi14 + fi14) {
                    case 9223372036854775807:
                        var v26 int = fi20
                        _ = v26
                    }
                    _ = fi20
                }
            }
            _ = fi14
        }
        var v27 string = v4
        if (true || (42 < 4294967295)) {
            for fi28 := 0; fi28 < 19; fi28++ {
                v27 = (v27 + "a")
                v5 += -1
                var v29 int = (fi28 - fi28)
                _ = fi28
                _ = v29
            }
        } else {
            if false {
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
            }
            for fi30 := 0; fi30 < 14; fi30++ {
                switch v4 {
                case "hello":
                    v13 -= 0
                case "\n\t":
                    if false {
                        v27 = ("" + (v4 + v27))
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                        v2 = append(v2, 1)
                        var v31 int32 = v5
                        _ = v31
                    }
                case "a":
                    fmt.Println(fi30)
                }
                _ = fi30
            }
            v4 = "foo bar"
        }
        _ = v27
    }
    var v32 int64 = 25
    for fi33 := 0; fi33 < 9; fi33++ {
        v4 = "a"
        _ = len(v1)
        fmt.Println(v13)
        _ = fi33
    }
    v34 := GenS1{F0: -74, F1: 65535}
    fmt.Println(v5, v4, v32)
    v35 := GenS1{F0: -11, F1: 400}
    v13 -= 255
    fmt.Println(v13)
    v2 = append(v2, 2147483648)
    for fi36 := 0; fi36 < 3; fi36++ {
        v32 += v32
        v37 := GenS1{F0: -21, F1: 916}
        var v38 int8 = 0
        _ = fi36
        _ = v37
        _ = v38
    }
    switch v4 {
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 65535
        }
    default:
        if true {
            v5 -= 0
            if ((v32 ^ v32) > -9223372036854775808) {
                if false {
                    v5 = (v5 | -671)
                    if true {
                        v32 = (^(v32))
                        fmt.Println(v32, v32, v13)
                    } else {
                        if (uint32(v13) > 475) {
                            if (2147483647 <= (v5 ^ 42)) {
                                if (-2.25 >= 0.1) {
                                    switch v4 {
                                    case "\n\t":
                                        switch "hello" {
                                        case "hello":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case "":
                                            var v39 bool = (!(!(-9223372036854775808 != 770)))
                                            _ = v39
                                        case "foo bar":
                                            v32 = 0
                                        default:
                                            v13 += (v13 >> ((4294967295) & 63))
                                        }
                                    }
                                    var v40 int16 = 42
                                    if len(v2) > 0 {
                                        v2[0] = 2
                                    }
                                    _ = v40
                                }
                            } else {
                                v41 := GenS1{F0: 1, F1: 323}
                                for fi42 := 0; fi42 < 10; fi42++ {
                                    var v43 uint = 4294967295
                                    switch v4 {
                                    case "":
                                        var v44 int64 = v32
                                        _ = v44
                                    case "foo bar":
                                        switch fi42 {
                                        case -309:
                                            v13 -= v13
                                        case -1:
                                            var v45 bool = false
                                            _ = v45
                                        case 32767:
                                            var v46 int32 = -2147483648
                                            _ = v46
                                        }
                                    default:
                                        var v47 int64 = v32
                                        _ = v47
                                    }
                                    v32 = v32
                                    _ = fi42
                                    _ = v43
                                }
                                _ = v41
                            }
                            v5 -= -399
                        } else {
                            v13 = (v13 - 415)
                        }
                        v4 = v4
                        var v48 float32 = float32(1e10)
                        _ = v48
                    }
                    for fi49 := 0; fi49 < 7; fi49++ {
                        if (1 != 2) {
                            for fi50 := 0; fi50 < 16; fi50++ {
                                var v51 int16 = 42
                                v52 := GenS2{F0: 128}
                                if false {
                                    var v53 float64 = (1e10 - (1e10 / ((-2.25 / 1e10) - -2.25)))
                                    if (true || true) {
                                        _ = len(v3)
                                        v54 := GenS3{F0: (!((v51 / ((v51) | 1)) != int16(v5)))}
                                        var v55 int32 = v5
                                        _ = v54
                                        _ = v55
                                    }
                                    v53 -= float64(1.5)
                                    _ = v53
                                } else {
                                    v5 = (v5 >> ((1) & 63))
                                    if true {
                                        fmt.Println(v32, v32)
                                    } else {
                                        fmt.Println(v13, v13)
                                    }
                                    if true {
                                        switch (v4 + "hello") {
                                        case "hello":
                                            v56 := GenS2{F0: 128}
                                            _ = v56
                                        }
                                        var v57 uint8 = 0
                                        _ = v57
                                    }
                                    v58 := GenS4{F0: 42, F1: int16(v5), F2: 143}
                                    _ = v58
                                }
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                            v59 := GenS1{F0: -1, F1: 65535}
                            for fi60 := 0; fi60 < 14; fi60++ {
                                fmt.Println(v5, fi49, fi49)
                                _ = fi60
                            }
                            _ = v59
                        } else {
                            for fi61 := 0; fi61 < 1; fi61++ {
                                fmt.Println(v32)
                                v4 = (v4 + v4)
                                _ = fi61
                            }
                            switch fi49 {
                            case -165:
                                if (!true) {
                                    switch ("hello" + "\n\t") {
                                    case "foo bar":
                                        var v62 int = fi49
                                        _ = v62
                                    case "a":
                                        fmt.Println(v13, fi49)
                                    case "\n\t":
                                        if true {
                                            var v63 string = (v4 + v4)
                                            var v64 string = (v63 + v4)
                                            _ = len(v3)
                                            _ = v63
                                            _ = v64
                                        }
                                    }
                                    if false {
                                        if true {
                                            var v65 int8 = -1
                                            if (((v13 - (v13 << ((18446744073709551615) & 63))) <= v13) && (fi49 == (fi49 | (fi49 | fi49)))) {
                                                fmt.Println(fi49)
                                                if true {
                                                    v66 := GenS5{F0: false, F1: 18446744073709551615}
                                                    var v67 float32 = 1e10
                                                    switch v4 {
                                                    case "foo bar":
                                                        if false {
                                                            if false {
                                                                for fi68 := 0; fi68 < 3; fi68++ {
                                                                    v5 -= 2147483647
                                                                    _ = fi68
                                                                }
                                                                for fi69 := 0; fi69 < 17; fi69++ {
                                                                    v65 -= genFn0((^(fi69)), -617, v5)
                                                                    var v70 float64 = 3.14159
                                                                    fmt.Println(v32, v13, v70)
                                                                    _ = fi69
                                                                    _ = v70
                                                                }
                                                                if (true && true) {
                                                                    for fi71 := 0; fi71 < 9; fi71++ {
                                                                        var v72 uint16 = 489
                                                                        v73 := GenS1{F0: v65, F1: 42}
                                                                        _ = fi71
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                    switch 32767 {
                                                                    case 2147483647:
                                                                        v2 = append(v2, 42)
                                                                    case 4294967295:
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                    default:
                                                                        if (!false) {
                                                                            var v74 uint8 = 2
                                                                            _ = v74
                                                                        }
                                                                    }
                                                                    v67 = 0.0
                                                                }
                                                            }
                                                            if ((180 != 18446744073709551615) && (false || true)) {
                                                                v75 := GenS6{F0: 1, F1: 255}
                                                                switch v4 {
                                                                case "":
                                                                    for fi76 := 0; fi76 < 18; fi76++ {
                                                                        var v77 int = fi76
                                                                        _ = fi76
                                                                        _ = v77
                                                                    }
                                                                case "foo bar":
                                                                    fmt.Println(v13)
                                                                default:
                                                                    for fi78 := 0; fi78 < 15; fi78++ {
                                                                        v65 -= (v65 % ((v65) | 1))
                                                                        var v79 float32 = 1e10
                                                                        _ = fi78
                                                                        _ = v79
                                                                    }
                                                                }
                                                                v80 := GenS7{F0: v13, F1: 4294967295}
                                                                _ = v75
                                                                _ = v80
                                                            }
                                                            for fi81 := 0; fi81 < 18; fi81++ {
                                                                fmt.Println(fi49)
                                                                _ = fi81
                                                            }
                                                        }
                                                    default:
                                                        v82 := GenS8{F0: ((32767 == v13) && false), F1: 9223372036854775807}
                                                        _ = v82
                                                    }
                                                    _ = v66
                                                    _ = v67
                                                }
                                                v65 += 1
                                            }
                                            _ = v65
                                        }
                                        if (false || false) {
                                            var v83 bool = false
                                            v32 = v32
                                            _ = v83
                                        }
                                    }
                                }
                            default:
                                v13 += 255
                            }
                            var v84 int8 = -110
                            _ = v84
                        }
                        _ = fi49
                    }
                }
            }
            for fi85 := 0; fi85 < 14; fi85++ {
                for fi86 := 0; fi86 < 16; fi86++ {
                    var v87 float64 = (3.14159 - (1e10 + -2.25))
                    _ = fi86
                    _ = v87
                }
                fmt.Println(fi85)
                _ = fi85
            }
        } else {
            if (1000 < -1) {
                var v88 int16 = 42
                v89 := GenS9{F0: 1}
                if ((1 < 1) || (!(true && (-79 < 1)))) {
                    fmt.Println(v13, v13, v13)
                    for fi90 := 0; fi90 < 17; fi90++ {
                        _ = len(v3)
                        var v91 uint = 18446744073709551615
                        _ = fi90
                        _ = v91
                    }
                } else {
                    var v92 float32 = (float32(1e10) - (1e10 / 1.5))
                    var v93 uint8 = 42
                    var v94 uint64 = 2
                    _ = v92
                    _ = v93
                    _ = v94
                }
                _ = v88
                _ = v89
            } else {
                for fi95 := 0; fi95 < 17; fi95++ {
                    var v96 int64 = 0
                    _ = fi95
                    _ = v96
                }
                if false {
                    for fi97 := 0; fi97 < 19; fi97++ {
                        var v98 uint32 = 260
                        for fi99 := 0; fi99 < 6; fi99++ {
                            var v100 float32 = ((1e10 / 0.1) * (0.0 * -2.25))
                            v98 += v98
                            v101 := GenS6{F0: (v5 ^ v5), F1: 18446744073709551615}
                            _ = fi99
                            _ = v100
                            _ = v101
                        }
                        v5 -= -955
                        _ = fi97
                        _ = v98
                    }
                    var v102 int64 = -387
                    v13 += 42
                    _ = v102
                }
            }
            if false {
                for fi103 := 0; fi103 < 12; fi103++ {
                    v5 -= v5
                    if true {
                        v4 = "foo bar"
                    }
                    fmt.Println(fi103, v5)
                    _ = fi103
                }
                switch "a" {
                case "foo bar":
                    var v104 uint64 = 1
                    _ = v104
                case "hello":
                    var v105 int = 0
                    _ = v105
                }
                var v106 int = 32767
                _ = v106
            }
            switch (v4 + "\n\t") {
            case "":
                for fi107 := 0; fi107 < 6; fi107++ {
                    _ = len(v2)
                    _ = fi107
                }
            case "a":
                if (true && false) {
                    var v108 uint32 = 42
                    var v109 bool = true
                    var v110 int16 = 32767
                    _ = v108
                    _ = v109
                    _ = v110
                } else {
                    if (!false) {
                        if false {
                            fmt.Println(v5)
                            v3 = append(v3, 361)
                            var v111 uint32 = 2
                            _ = v111
                        } else {
                            if false {
                                v13 = (v13 ^ v13)
                            }
                            fmt.Println(v4, v32, v5)
                        }
                        fmt.Println(v32)
                        for fi112 := 0; fi112 < 19; fi112++ {
                            switch v4 {
                            case "":
                                v113 := GenS9{F0: 4294967295}
                                _ = v113
                            }
                            for fi114 := 0; fi114 < 13; fi114++ {
                                var v115 int = (fi112 | (fi114 & fi114))
                                if true {
                                    fmt.Println(v115, v5)
                                    v5 = v5
                                }
                                v115 += (fi112 - (^(fi112)))
                                _ = fi114
                                _ = v115
                            }
                            for fi116 := 0; fi116 < 13; fi116++ {
                                v32 += v32
                                _ = fi116
                            }
                            _ = fi112
                        }
                    } else {
                        v3 = append(v3, 2147483648)
                        if (!(0 >= v13)) {
                            v117 := GenS10{F0: v4, F1: -63, F2: 42}
                            _ = v117
                        }
                        if (813 >= 2) {
                            v118 := GenS11{F0: v4, F1: 2}
                            _ = v118
                        }
                    }
                    fmt.Println(v5, v5)
                }
            default:
                if true {
                    v119 := GenS10{F0: "hello", F1: -1, F2: 0}
                    v5 = v5
                    _ = v119
                } else {
                    for fi120 := 0; fi120 < 16; fi120++ {
                        if (v5 > 1) {
                            v13 -= (v13 << ((2) & 63))
                            for fi121 := 0; fi121 < 2; fi121++ {
                                var v122 string = v4
                                var v123 int16 = -32768
                                fmt.Println(v123, v13, v123)
                                _ = fi121
                                _ = v122
                                _ = v123
                            }
                        } else {
                            v4 = (v4 + ("foo bar" + ""))
                            v4 = (v4 + (("foo bar" + ("foo bar" + "a")) + ("\n\t" + v4)))
                            var v124 int16 = -966
                            _ = v124
                        }
                        if true {
                            v32 = (v32 ^ v32)
                        } else {
                            switch v4 {
                            case "\n\t":
                                var v125 uint32 = 162
                                _ = v125
                            case "hello":
                                var v126 int8 = 42
                                _ = v126
                            case "":
                                _ = len(v2)
                            default:
                                v127 := GenS9{F0: 42}
                                _ = v127
                            }
                        }
                        _ = fi120
                    }
                    v13 = (v13 - v13)
                }
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
