package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int; F2 uint }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int64; F1 uint; F2 uint64 }
type GenS13 struct { F0 uint8; F1 uint16 }
type GenS14 struct { F0 uint; F1 string }
type GenS15 struct { F0 int }
func genFn0(p0 int32, p1 int8) int16 {
    p0 += (p0 >> ((298) & 63))
    v7 := GenS1{F0: 32767, F1: p0}
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 int16) int16 {
    if (84 == 100) {
        v8 := GenS1{F0: 32768, F1: 517}
        for fi9 := 0; fi9 < 7; fi9++ {
            if true {
                for fi10 := 0; fi10 < 13; fi10++ {
                    var v11 uint = 4294967295
                    var v12 uint64 = 0
                    v13 := []int{fi9, (fi10 ^ int(p0))}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                if false {
                    for fi14 := 0; fi14 < 4; fi14++ {
                        p0 = (p0 | p0)
                        var v15 uint = 2
                        _ = fi14
                        _ = v15
                    }
                } else {
                    v16 := GenS1{F0: fi9, F1: 42}
                    _ = v16
                }
                for fi17 := 0; fi17 < 16; fi17++ {
                    p0 = p0
                    _ = fi17
                }
            } else {
                if false {
                    fmt.Println(p0, fi9)
                    v18 := GenS2{F0: 42}
                    var v19 bool = (true || true)
                    _ = v18
                    _ = v19
                } else {
                    v20 := GenS1{F0: (fi9 ^ (fi9 + (^(fi9)))), F1: 42}
                    _ = v20
                }
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 uint64 = 0
                    p0 -= p0
                    v23 := GenS1{F0: fi9, F1: 42}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                if (!(false || (!false))) {
                    v24 := GenS2{F0: 100}
                    p0 -= (p0 & (p0 - (p0 + 0)))
                    _ = v24
                } else {
                    p0 -= (p0 / ((0) | 1))
                    if false {
                        var v25 int = 1
                        v25 -= (v25 - 651)
                        var v26 int32 = 0
                        _ = v25
                        _ = v26
                    } else {
                        if (true || (32767 < int16(fi9))) {
                            if (0 > 42) {
                                for fi27 := 0; fi27 < 15; fi27++ {
                                    p0 = (p0 & (p0 << ((2) & 63)))
                                    var v28 float32 = 1e10
                                    fmt.Println(p0, v28)
                                    _ = fi27
                                    _ = v28
                                }
                                v29 := []int{}
                                for fi30 := 0; fi30 < 9; fi30++ {
                                    var v31 uint8 = 172
                                    if ((184 > 670) && (!true)) {
                                        switch -1 {
                                        case 32768:
                                            var v32 int8 = int8(v31)
                                            var v33 string = ""
                                            v31 += v31
                                            switch v33 {
                                            case "hello":
                                                var v34 uint64 = 2
                                                if (false && true) {
                                                    v35 := GenS2{F0: 0}
                                                    _ = v35
                                                }
                                                _ = v34
                                            default:
                                                fmt.Println(fi9, v32)
                                            }
                                            _ = v32
                                            _ = v33
                                        case 0:
                                            fmt.Println(fi30)
                                            v31 += (v31 - (v31 ^ v31))
                                            switch 32768 {
                                            case 529:
                                                fmt.Println(p0, fi30, fi9)
                                            case 9223372036854775807:
                                                switch 9223372036854775807 {
                                                case 178:
                                                    var v36 uint8 = 42
                                                    _ = v36
                                                default:
                                                    switch (fi9 >> ((625) & 63)) {
                                                    case -655:
                                                        var v37 int = (-(fi9))
                                                        _ = v37
                                                    case 9223372036854775807:
                                                        var v38 int16 = (p0 >> ((18446744073709551615) & 63))
                                                        _ = v38
                                                    case -2147483648:
                                                        if (!(false || false)) {
                                                            switch ("\n\t" + "hello") {
                                                            case "":
                                                                if true {
                                                                    v39 := GenS2{F0: -128}
                                                                    v31 += v31
                                                                    var v40 uint16 = 298
                                                                    v40 -= (^(v40))
                                                                    _ = len(v29)
                                                                    _ = v39
                                                                    _ = v40
                                                                }
                                                            case "foo bar":
                                                                if false {
                                                                    fmt.Println(fi30, fi30, p0)
                                                                    v31 += v31
                                                                    v41 := GenS3{F0: 0, F1: 32767, F2: 1}
                                                                    if (440 == -1) {
                                                                        _ = len(v29)
                                                                        p0 = p0
                                                                        fmt.Println(p0)
                                                                    } else {
                                                                        v31 = v31
                                                                        p0 = (p0 & (p0 % ((p0) | 1)))
                                                                        p0 -= (p0 ^ int16(fi30))
                                                                    }
                                                                    switch fi30 {
                                                                    case -38:
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            if true {
                                                                                v42 := GenS4{F0: 769, F1: 0}
                                                                                fmt.Println(p0)
                                                                                _ = v42
                                                                            } else {
                                                                                v31 = (v31 * (v31 + v31))
                                                                                if false {
                                                                                    v29 = append(v29, 0)
                                                                                    v29 = append(v29, 127)
                                                                                    p0 = -32768
                                                                                    fmt.Println(fi30, fi9, fi30)
                                                                                    if len(v29) >= 2 {
                                                                                        v29 = v29[1:len(v29)]
                                                                                    }
                                                                                    var v43 uint64 = 2
                                                                                    _ = v43
                                                                                } else {
                                                                                    fmt.Println(fi30)
                                                                                }
                                                                                if ((true && (!(-2.25 == -2.25))) && ((p0 ^ 1) <= p0)) {
                                                                                    switch fi30 {
                                                                                    case 792:
                                                                                        v31 -= (v31 + 127)
                                                                                    case 0:
                                                                                        v29 = append(v29, fi30)
                                                                                    default:
                                                                                        var v44 uint16 = 610
                                                                                        _ = v44
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    var v45 int = fi30
                                                                                    _ = v45
                                                                                } else {
                                                                                    if len(v29) >= 2 {
                                                                                        v29 = v29[1:len(v29)]
                                                                                    }
                                                                                }
                                                                                v46 := GenS5{F0: p0}
                                                                                p0 = genFn0(0, 42)
                                                                                p0 = 1000
                                                                                _ = v46
                                                                            }
                                                                            p0 = p0
                                                                        case "":
                                                                            fmt.Println(v31)
                                                                            v47 := GenS1{F0: fi30, F1: int32(v31)}
                                                                            _ = v47
                                                                        default:
                                                                            var v48 uint = 0
                                                                            _ = v48
                                                                        }
                                                                    case 9223372036854775807:
                                                                        if len(v29) >= 2 {
                                                                            v29 = v29[1:len(v29)]
                                                                        }
                                                                    }
                                                                    _ = v41
                                                                }
                                                            case "a":
                                                                p0 = 32767
                                                            default:
                                                                var v49 int32 = 42
                                                                _ = v49
                                                            }
                                                            v50 := GenS6{F0: ("a" + "\n\t")}
                                                            _ = v50
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    _ = fi30
                                    _ = v31
                                }
                                _ = v29
                            } else {
                                if (0.1 > 1.5) {
                                    p0 = p0
                                    switch "foo bar" {
                                    case "\n\t":
                                        if (!true) {
                                            v51 := []int{}
                                            _ = v51
                                        }
                                    }
                                }
                                v52 := GenS7{F0: fi9, F1: 65535}
                                _ = v52
                            }
                            var v53 uint16 = 32767
                            _ = v53
                        }
                        if false {
                            var v54 int16 = p0
                            fmt.Println(p0)
                            _ = v54
                        } else {
                            p0 = (p0 - -1)
                            if (!true) {
                                p0 = -1
                            }
                        }
                        switch (fi9 - (^(fi9))) {
                        case 4294967295:
                            p0 += p0
                        }
                    }
                    if (0 < 42) {
                        v55 := []int{fi9, 127, 4294967295, 2147483647}
                        _ = v55
                    }
                }
            }
            switch "a" {
            case "hello":
                if false {
                    p0 -= (p0 / ((32767) | 1))
                }
            }
            _ = fi9
        }
        _ = v8
    }
    switch "hello" {
    case "foo bar":
        for fi56 := 0; fi56 < 6; fi56++ {
            fmt.Println(p0, fi56)
            p0 = p0
            _ = fi56
        }
    default:
        for fi57 := 0; fi57 < 2; fi57++ {
            for fi58 := 0; fi58 < 16; fi58++ {
                switch 951 {
                case -1:
                    v59 := GenS1{F0: -101, F1: 2147483647}
                    _ = v59
                }
                _ = fi58
            }
            _ = fi57
        }
    }
    for fi60 := 0; fi60 < 5; fi60++ {
        for fi61 := 0; fi61 < 13; fi61++ {
            fmt.Println(p0)
            _ = fi61
        }
        for fi62 := 0; fi62 < 13; fi62++ {
            if false {
                var v63 bool = ((-128 != int8(fi60)) && true)
                var v64 string = ""
                _ = v63
                _ = v64
            }
            _ = fi62
        }
        v65 := []int{fi60, 1, fi60, 32767}
        _ = fi60
        _ = v65
    }
    v66 := []int{}
    _ = p0
    _ = v66
    return 0
}

func main() {
    var v1 int = 390
    var v2 uint8 = 127
    var v3 int16 = 0
    var v4 bool = (61 > 42)
    v5 := []uint32{748, 4294967295}
    var v6 bool = (42 >= 307)
    v6 = (48 > 42)
    v67 := GenS4{F0: uint64(v2), F1: 42}
    if v4 {
        v68 := GenS5{F0: v3}
        _ = len(v5)
        if (false || v4) {
            var v69 uint32 = 986
            v3 += (v3 / (((v3 | -32768)) | 1))
            v1 += v1
            _ = v69
        }
        _ = v68
    }
    var v70 uint64 = 42
    v6 = ((((0.0 - -2.25) + float64(1.5)) - (0.1 + (-2.25 / (1.5 / float64(1e10))))) > 1.5)
    if ((false && (false || false)) || v6) {
        for fi71 := 0; fi71 < 4; fi71++ {
            for fi72 := 0; fi72 < 17; fi72++ {
                _ = len(v5)
                _ = fi72
            }
            v6 = (0.1 != 0.1)
            _ = fi71
        }
        if (-50 != 42) {
            if (127 > 100) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            v3 = (v3 + genFn0(1, 100))
            v4 = ((v3 / ((genFn1(1)) | 1)) > v3)
        } else {
            v5 = append(v5, 65535)
        }
        switch ("\n\t" + "hello") {
        case "":
            for fi73 := 0; fi73 < 19; fi73++ {
                if v6 {
                    for fi74 := 0; fi74 < 10; fi74++ {
                        var v75 int64 = 781
                        switch fi74 {
                        case 400:
                            v70 = (v70 % ((v70) | 1))
                        case 2147483647:
                            fmt.Println(fi74, v2)
                        }
                        var v76 uint8 = (v2 / ((v2) | 1))
                        _ = fi74
                        _ = v75
                        _ = v76
                    }
                    if (!(0 < 903)) {
                        v70 -= (v70 / (((^(v70))) | 1))
                        for fi77 := 0; fi77 < 5; fi77++ {
                            var v78 int64 = 429
                            if v6 {
                                fmt.Println(v4, fi77, v2)
                                fmt.Println(v2)
                                v5 = append(v5, 0)
                            } else {
                                v79 := GenS8{F0: 494}
                                v1 += fi77
                                _ = v79
                            }
                            _ = fi77
                            _ = v78
                        }
                        var v80 int64 = -9223372036854775808
                        _ = v80
                    }
                    v81 := GenS2{F0: 1}
                    _ = v81
                }
                for fi82 := 0; fi82 < 8; fi82++ {
                    _ = len(v5)
                    if false {
                        var v83 string = "a"
                        _ = v83
                    } else {
                        fmt.Println(v3, fi73)
                        var v84 uint32 = 65535
                        if v6 {
                            var v85 float64 = 0.1
                            fmt.Println(v3, v1, fi82)
                            v84 += v84
                            _ = v85
                        } else {
                            var v86 int64 = 1
                            v6 = (0.1 == -2.25)
                            var v87 float64 = -2.25
                            v3 += 42
                            v88 := GenS4{F0: v70, F1: 1}
                            if false {
                                v70 = (v70 + v70)
                                _ = len(v5)
                            }
                            switch "hello" {
                            case "foo bar":
                                v89 := GenS1{F0: (^(fi73)), F1: 2147483647}
                                _ = v89
                            default:
                                v90 := GenS5{F0: (v3 * (v3 + int16(v86)))}
                                _ = v90
                            }
                            var v91 uint16 = 1
                            v87 -= v87
                            var v92 string = "\n\t"
                            v70 += (v70 / (((v70 - (v70 ^ uint64(v84)))) | 1))
                            _ = v86
                            _ = v87
                            _ = v88
                            _ = v91
                            _ = v92
                        }
                        _ = v84
                    }
                    var v93 int16 = (v3 * (v3 + (-(v3))))
                    _ = fi82
                    _ = v93
                }
                _ = fi73
            }
        case "\n\t":
            if false {
                var v94 int16 = int16(v1)
                if v4 {
                    v3 = genFn1(v3)
                }
                _ = v94
            }
        case "hello":
            if v6 {
                fmt.Println(v70, v70, v1)
                var v95 uint16 = 0
                v4 = v4
                _ = v95
            } else {
                v3 = genFn1(393)
            }
        default:
            for fi96 := 0; fi96 < 11; fi96++ {
                var v97 string = ("" + "hello")
                switch (^(fi96)) {
                case -925:
                    v5 = append(v5, 2)
                }
                _ = fi96
                _ = v97
            }
        }
    } else {
        if (-128 >= 1) {
            for fi98 := 0; fi98 < 17; fi98++ {
                if (!("\n\t" > (("foo bar" + "a") + ("hello" + ("a" + "a"))))) {
                    if false {
                        v70 -= 871
                    } else {
                        if false {
                            v2 += 39
                        } else {
                            var v99 uint16 = 889
                            _ = v99
                        }
                        for fi100 := 0; fi100 < 6; fi100++ {
                            var v101 int32 = -2147483648
                            v102 := GenS4{F0: v70, F1: (v70 % ((42) | 1))}
                            v101 = (v101 & (-(v101)))
                            switch "\n\t" {
                            case "hello":
                                if v4 {
                                    if ((-2.25 - 0.0) > ((((1.5 / (-2.25 - float32(1.5))) + 0.1) - 1.5) + (float32((1.5 - 0.1)) * 0.1))) {
                                        var v103 int = v1
                                        v104 := GenS3{F0: 42, F1: (fi98 - fi100), F2: uint(v2)}
                                        _ = v103
                                        _ = v104
                                    } else {
                                        if (!((1.5 > (1e10 + ((1.5 / float64((1e10 - (((1e10 * -2.25) * 0.0) - 1.5)))) / 1.5))) || (885 <= 65535))) {
                                            if v6 {
                                                if (0.0 == 0.0) {
                                                    var v105 float64 = ((1e10 + float64(1.5)) + -2.25)
                                                    v4 = true
                                                    v6 = false
                                                    fmt.Println(v3, v70, v70)
                                                    _ = v105
                                                } else {
                                                    if ((-(v101)) > v101) {
                                                        switch 89 {
                                                        case -1:
                                                            if v6 {
                                                                if v6 {
                                                                    if (!v4) {
                                                                        if v6 {
                                                                            v4 = false
                                                                            var v106 int = (fi100 | fi100)
                                                                            _ = v106
                                                                        } else {
                                                                            v4 = (((false && (v3 > genFn1(v3))) || true) || ((^(v2)) <= (v2 | 255)))
                                                                            var v107 uint = 18446744073709551615
                                                                            if v6 {
                                                                                switch ("hello" + "hello") {
                                                                                case "hello":
                                                                                    v108 := GenS3{F0: 0, F1: -2147483648, F2: (v107 / ((v107) | 1))}
                                                                                    _ = v108
                                                                                case "\n\t":
                                                                                    var v109 int32 = (v101 >> ((v107) & 63))
                                                                                    _ = v109
                                                                                case "a":
                                                                                    if v4 {
                                                                                        if v4 {
                                                                                            fmt.Println(v1, v4)
                                                                                            if (0 <= 21) {
                                                                                                fmt.Println(v70, v1)
                                                                                            } else {
                                                                                                v110 := GenS7{F0: 128, F1: 42}
                                                                                                var v111 int = -2147483648
                                                                                                _ = v110
                                                                                                _ = v111
                                                                                            }
                                                                                        } else {
                                                                                            var v112 uint = v107
                                                                                            var v113 int64 = -901
                                                                                            _ = v112
                                                                                            _ = v113
                                                                                        }
                                                                                        v114 := GenS9{F0: 1, F1: 127}
                                                                                        _ = v114
                                                                                    }
                                                                                default:
                                                                                    var v115 uint = (v107 << (((v107 + v107)) & 63))
                                                                                    if v4 {
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                        v4 = v4
                                                                                        v101 = (-(v101))
                                                                                        fmt.Println(v115, v6, v1)
                                                                                        var v116 int32 = int32(v1)
                                                                                        _ = v116
                                                                                    } else {
                                                                                        fmt.Println(v4, v115, v4)
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = 4294967295
                                                                                        }
                                                                                        if v4 {
                                                                                            var v117 int8 = 100
                                                                                            if v6 {
                                                                                                var v118 int16 = -1
                                                                                                _ = v118
                                                                                            }
                                                                                            v4 = ((v2 | (v2 << (((v115 & v115)) & 63))) <= 31)
                                                                                            _ = v117
                                                                                        }
                                                                                    }
                                                                                    _ = v115
                                                                                }
                                                                            } else {
                                                                                v119 := GenS10{F0: (0 <= 2147483647)}
                                                                                _ = v119
                                                                            }
                                                                            v120 := GenS5{F0: v3}
                                                                            v107 += 255
                                                                            _ = v107
                                                                            _ = v120
                                                                        }
                                                                        if len(v5) > 0 {
                                                                            v5[0] = 866
                                                                        }
                                                                        v121 := GenS6{F0: "\n\t"}
                                                                        if len(v5) > 0 {
                                                                            v5[0] = 580
                                                                        }
                                                                        _ = v121
                                                                    } else {
                                                                        v6 = (!false)
                                                                    }
                                                                    var v122 uint64 = (v70 << ((4294967295) & 63))
                                                                    v1 = (v1 & -841)
                                                                    _ = v122
                                                                }
                                                                var v123 int8 = 42
                                                                v1 = fi100
                                                                _ = v123
                                                            } else {
                                                                switch 569 {
                                                                case 9223372036854775807:
                                                                    if true {
                                                                        if len(v5) > 0 {
                                                                            v5[0] = 2147483648
                                                                        }
                                                                    }
                                                                }
                                                                if ((1.5 / 0.1) >= 0.1) {
                                                                    var v124 bool = (3.14159 > float64(0.0))
                                                                    _ = v124
                                                                } else {
                                                                    fmt.Println(v70)
                                                                    v125 := GenS11{F0: v3}
                                                                    _ = v125
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        var v126 uint32 = 2147483647
                                                        _ = v126
                                                    }
                                                    var v127 uint8 = v2
                                                    if true {
                                                        _ = len(v5)
                                                    }
                                                    _ = v127
                                                }
                                                var v128 float32 = float32((1e10 + 0.0))
                                                _ = v128
                                            } else {
                                                v70 += v70
                                                switch "foo bar" {
                                                case "":
                                                    if (((-2.25 < 1.5) || false) && false) {
                                                        fmt.Println(v70)
                                                        switch 973 {
                                                        case -237:
                                                            var v129 float32 = 1.5
                                                            _ = v129
                                                        case -423:
                                                            if len(v5) > 0 {
                                                                v5[0] = 2147483647
                                                            }
                                                        case 561:
                                                            var v130 string = "a"
                                                            _ = v130
                                                        default:
                                                            if v6 {
                                                                if ("hello" == "\n\t") {
                                                                    if (2 >= 2) {
                                                                        if (true && (false && v4)) {
                                                                            fmt.Println(v70, v1, v101)
                                                                            fmt.Println(v6)
                                                                            v3 += v3
                                                                        }
                                                                        v6 = v6
                                                                        v70 += 2
                                                                    } else {
                                                                        if ((!((^(v70)) > v70)) || (float64((0.1 / 1e10)) <= 1.5)) {
                                                                            switch ("hello" + "hello") {
                                                                            case "foo bar":
                                                                                var v131 float64 = 1e10
                                                                                _ = v131
                                                                            case "":
                                                                                var v132 uint = 377
                                                                                _ = v132
                                                                            default:
                                                                                if (1 < -128) {
                                                                                    v101 += (v101 >> ((2) & 63))
                                                                                    v4 = v4
                                                                                    v3 += 1
                                                                                }
                                                                            }
                                                                        }
                                                                        fmt.Println(v3, v6)
                                                                        var v133 uint8 = v2
                                                                        _ = v133
                                                                    }
                                                                    v134 := GenS12{F0: 0, F1: 18446744073709551615, F2: uint64(v2)}
                                                                    v135 := GenS13{F0: (v2 % ((v2) | 1)), F1: 0}
                                                                    _ = v134
                                                                    _ = v135
                                                                }
                                                            }
                                                        }
                                                    }
                                                case "a":
                                                    if (true || ((3.14159 * 0.1) >= 3.14159)) {
                                                        if false {
                                                            var v136 int32 = -1
                                                            _ = v136
                                                        } else {
                                                            fmt.Println(v4)
                                                            v137 := GenS14{F0: 30, F1: ""}
                                                            _ = v137
                                                        }
                                                        if (!("a" == ("hello" + "a"))) {
                                                            v6 = v4
                                                        } else {
                                                            v138 := GenS15{F0: fi100}
                                                            var v139 int64 = 1
                                                            _ = v138
                                                            _ = v139
                                                        }
                                                    } else {
                                                        v101 -= v101
                                                        var v140 int32 = (v101 * v101)
                                                        fmt.Println(fi98)
                                                        _ = v140
                                                    }
                                                default:
                                                    v2 += v2
                                                }
                                                var v141 int64 = -9223372036854775808
                                                v141 = (v141 - v141)
                                                fmt.Println(fi98, v1, v141)
                                                if (v6 || false) {
                                                    var v142 int64 = v141
                                                    _ = v142
                                                }
                                                _ = v141
                                            }
                                            var v143 int = 9223372036854775807
                                            v144 := GenS11{F0: 32767}
                                            _ = v143
                                            _ = v144
                                        }
                                        var v145 uint8 = v2
                                        var v146 uint32 = 880
                                        _ = v145
                                        _ = v146
                                    }
                                    var v147 int64 = 42
                                    if true {
                                        v5 = append(v5, 0)
                                    } else {
                                        v148 := GenS3{F0: 4294967295, F1: (v1 + (v1 | fi100)), F2: 65535}
                                        _ = v148
                                    }
                                    _ = v147
                                }
                            }
                            if false {
                                _ = len(v5)
                                var v149 uint16 = 65535
                                v6 = false
                                v1 -= fi100
                                _ = v149
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if len(v5) > 0 {
                                v5[0] = 488
                            }
                            _ = fi100
                            _ = v101
                            _ = v102
                        }
                        var v150 string = "foo bar"
                        _ = v150
                    }
                    switch "a" {
                    case "foo bar":
                        for fi151 := 0; fi151 < 4; fi151++ {
                            var v152 bool = v4
                            v4 = (false || (-877 > -9223372036854775808))
                            _ = fi151
                            _ = v152
                        }
                    case "a":
                        if false {
                            switch ("foo bar" + ("hello" + "foo bar")) {
                            case "foo bar":
                                v153 := GenS7{F0: (v1 + int(v3)), F1: 65535}
                                _ = v153
                            case "":
                                v5 = append(v5, 2)
                            default:
                                if len(v5) > 0 {
                                    v5[0] = 2
                                }
                            }
                        }
                    }
                } else {
                    var v154 int16 = 711
                    var v155 bool = ((("" + "hello") > "") && ((!((("hello" + "") >= "foo bar") || true)) || (831 == 127)))
                    var v156 uint = 1
                    _ = v154
                    _ = v155
                    _ = v156
                }
                for fi157 := 0; fi157 < 8; fi157++ {
                    v3 = -1
                    _ = fi157
                }
                switch v1 {
                case -2147483648:
                    if len(v5) > 0 {
                        v5[0] = 4294967295
                    }
                case 2147483647:
                    v1 -= v1
                case -9223372036854775808:
                    for fi158 := 0; fi158 < 16; fi158++ {
                        if (127 > -59) {
                            if len(v5) > 0 {
                                v5[0] = 0
                            }
                        }
                        var v159 int32 = 2147483647
                        v2 -= v2
                        _ = fi158
                        _ = v159
                    }
                default:
                    if true {
                        v2 = v2
                    }
                }
                _ = fi98
            }
            var v160 uint32 = 1
            _ = v160
        } else {
            switch "\n\t" {
            case "a":
                var v161 uint64 = v70
                _ = v161
            default:
                v6 = (true || (42 <= 97))
            }
        }
        v4 = ("a" < ("foo bar" + "hello"))
        for fi162 := 0; fi162 < 4; fi162++ {
            var v163 int32 = 1
            _ = fi162
            _ = v163
        }
    }
    v4 = (-2147483648 != -1)
    var v164 uint64 = 4294967295
    if v4 {
        for fi165 := 0; fi165 < 3; fi165++ {
            if len(v5) > 0 {
                v5[0] = 451
            }
            v166 := GenS5{F0: (v3 % ((v3) | 1))}
            _ = fi165
            _ = v166
        }
    } else {
        switch (v1 & v1) {
        case 1:
            v167 := GenS2{F0: -1}
            _ = v167
        }
        fmt.Println(v3, v2)
    }
    var v168 float32 = ((0.1 / -2.25) - 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v164:", v164)
    fmt.Println("v168:", v168)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v164
    _ = v168
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v70
}
