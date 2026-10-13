package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 uint; F2 int8 }
type GenS6 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint32; F1 int8; F2 int }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint64; F1 int8 }
type GenS13 struct { F0 int8 }
func genFn0(p0 int8) uint8 {
    fmt.Println(p0)
    switch 265 {
    case -2147483648:
        for fi7 := 0; fi7 < 19; fi7++ {
            fmt.Println(p0, fi7)
            if (32767 <= fi7) {
                v8 := []int{(fi7 % (((fi7 & -973)) | 1)), (fi7 | 1), 9223372036854775807, fi7}
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                var v9 int = (fi7 % (((^(fi7))) | 1))
                _ = v8
                _ = v9
            } else {
                fmt.Println(p0, p0)
            }
            p0 += (p0 | (p0 >> ((65535) & 63)))
            _ = fi7
        }
    case -1:
        p0 += p0
    case 1:
        var v10 int32 = -1
        _ = v10
    }
    v11 := GenS1{F0: 1, F1: -9223372036854775808, F2: 1000}
    var v12 uint64 = 0
    _ = p0
    _ = v11
    _ = v12
    return 87
}

func main() {
    v1 := []int{267, 128, 2147483647}
    var v2 int32 = 42
    var v3 uint64 = 726
    var v4 int = 2147483647
    var v5 int16 = -1
    var v6 uint8 = 0
    var v13 int32 = (-(v2))
    if true {
        if true {
            v5 += v5
            v14 := GenS2{F0: (-2.25 <= 1.5)}
            for fi15 := 0; fi15 < 3; fi15++ {
                v16 := GenS3{F0: "a"}
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 float32 = 0.1
            var v19 uint = 200
            _ = fi17
            _ = v18
            _ = v19
        }
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            v4 -= fi20
            if (!(true || ((((v3 | 682) >= (v3 * (v3 % ((57) | 1)))) || ((v6 ^ genFn0(-46)) < v6)) && (!true)))) {
                v13 = 42
                if (695 > 0) {
                    var v21 float64 = 1.5
                    if (!true) {
                        for fi22 := 0; fi22 < 11; fi22++ {
                            if false {
                                v23 := GenS4{F0: fi20}
                                _ = v23
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi22
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = -847
                    }
                    _ = v21
                } else {
                    var v24 bool = false
                    for fi25 := 0; fi25 < 13; fi25++ {
                        var v26 int16 = v5
                        _ = fi25
                        _ = v26
                    }
                    _ = v24
                }
                switch (v4 * (v4 | fi20)) {
                case 42:
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v1 = append(v1, v4)
                        _ = fi27
                    }
                }
            }
            fmt.Println(v3, v4, v6)
            _ = fi20
        }
        fmt.Println(v6, v2)
        v2 = v13
    }
    switch "foo bar" {
    case "foo bar":
        switch (v4 - (v4 >> ((4294967295) & 63))) {
        case -354:
            if false {
                var v28 string = "hello"
                for fi29 := 0; fi29 < 11; fi29++ {
                    var v30 bool = false
                    switch "" {
                    case "":
                        v31 := GenS5{F0: 18446744073709551615, F1: 903, F2: 100}
                        _ = v31
                    case "hello":
                        v6 = v6
                    case "\n\t":
                        for fi32 := 0; fi32 < 11; fi32++ {
                            v30 = (!(((v3 / ((18446744073709551615) | 1)) != (v3 % (((^(v3))) | 1))) && v30))
                            var v33 uint32 = 684
                            _ = fi32
                            _ = v33
                        }
                    }
                    if ((2 > 2147483647) || v30) {
                        if (0.1 < -2.25) {
                            for fi34 := 0; fi34 < 12; fi34++ {
                                v13 = v13
                                _ = fi34
                            }
                            v35 := GenS6{F0: 404, F1: 42, F2: v2}
                            _ = v35
                        }
                        var v36 int16 = 799
                        _ = v36
                    }
                    _ = fi29
                    _ = v30
                }
                _ = v28
            } else {
                fmt.Println(v13, v6)
                v37 := GenS7{F0: (v13 << ((0) & 63)), F1: 255}
                if (false && (2147483648 >= 314)) {
                    var v38 int = v4
                    if false {
                        v39 := GenS2{F0: ("foo bar" >= ((("a" + ("" + "\n\t")) + "") + ("" + ("" + "foo bar"))))}
                        switch "hello" {
                        case "a":
                            if (v5 != v5) {
                                for fi40 := 0; fi40 < 11; fi40++ {
                                    var v41 string = ("" + "a")
                                    var v42 float32 = 1e10
                                    switch (v4 | (v4 << ((65535) & 63))) {
                                    case 128:
                                        v43 := GenS8{F0: 665}
                                        _ = v43
                                    case 0:
                                        if (true && (32767 == 317)) {
                                            if false {
                                                v1 = append(v1, fi40)
                                                var v44 string = v41
                                                _ = v44
                                            } else {
                                                switch v41 {
                                                case "":
                                                    v42 += 0.0
                                                }
                                                v1 = append(v1, (v38 - -938))
                                                switch v41 {
                                                case "a":
                                                    var v45 int8 = 42
                                                    _ = v45
                                                }
                                            }
                                        }
                                    default:
                                        for fi46 := 0; fi46 < 1; fi46++ {
                                            v38 = v4
                                            fmt.Println(v13, fi40, v41)
                                            _ = fi46
                                        }
                                    }
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                }
                                var v47 uint = 962
                                _ = v47
                            } else {
                                v13 = 1
                                v48 := GenS5{F0: 18446744073709551615, F1: uint(v3), F2: -128}
                                for fi49 := 0; fi49 < 14; fi49++ {
                                    for fi50 := 0; fi50 < 17; fi50++ {
                                        var v51 bool = (!false)
                                        switch ((("" + "foo bar") + "hello") + ("" + "foo bar")) {
                                        case "foo bar":
                                            if ((v6 | v6) != 1) {
                                                fmt.Println(v38, v38)
                                                v3 = 0
                                                var v52 uint64 = v3
                                                _ = v52
                                            }
                                        case "a":
                                            var v53 float64 = 1.5
                                            v5 = v5
                                            v1 = append(v1, 4294967295)
                                            _ = v53
                                        default:
                                            var v54 int16 = v5
                                            var v55 float64 = 0.0
                                            _ = v54
                                            _ = v55
                                        }
                                        v13 -= v13
                                        _ = fi50
                                        _ = v51
                                    }
                                    if (!(1 > -474)) {
                                        if true {
                                            v2 = -661
                                            fmt.Println(v13, v2, v6)
                                            var v56 uint16 = 900
                                            _ = v56
                                        } else {
                                            switch 42 {
                                            case 4294967295:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case 9:
                                                v57 := GenS9{F0: 2147483647, F1: -1, F2: (fi49 * 9223372036854775807)}
                                                _ = v57
                                            case 32767:
                                                v3 += (v3 ^ 18446744073709551615)
                                            default:
                                                if false {
                                                    switch -432 {
                                                    case 309:
                                                        v2 = 952
                                                    case 42:
                                                        _ = len(v1)
                                                    case -696:
                                                        fmt.Println(v2, v13, v2)
                                                    }
                                                    switch ("" + "a") {
                                                    case "":
                                                        for fi58 := 0; fi58 < 13; fi58++ {
                                                            v59 := GenS5{F0: 1, F1: uint(v6), F2: 1}
                                                            if true {
                                                                fmt.Println(v6, v4)
                                                                fmt.Println(v13, v4)
                                                                fmt.Println(v4, v13, v5)
                                                                var v60 uint32 = 0
                                                                _ = v60
                                                            } else {
                                                                v6 = v6
                                                            }
                                                            _ = fi58
                                                            _ = v59
                                                        }
                                                    case "\n\t":
                                                        if (42 <= 127) {
                                                            for fi61 := 0; fi61 < 16; fi61++ {
                                                                var v62 float64 = (-2.25 / 0.1)
                                                                _ = fi61
                                                                _ = v62
                                                            }
                                                            if false {
                                                                for fi63 := 0; fi63 < 14; fi63++ {
                                                                    switch (^(fi49)) {
                                                                    case -2147483648:
                                                                        if (0.0 < 0.0) {
                                                                            var v64 uint8 = genFn0(80)
                                                                            switch ("hello" + "foo bar") {
                                                                            case "":
                                                                                if (2 != 0) {
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v64
                                                                        } else {
                                                                            var v65 string = "foo bar"
                                                                            var v66 uint16 = 42
                                                                            _ = v65
                                                                            _ = v66
                                                                        }
                                                                    case 146:
                                                                        switch fi63 {
                                                                        case -1:
                                                                            var v67 int32 = 2147483647
                                                                            _ = v67
                                                                        }
                                                                    case 2147483647:
                                                                        fmt.Println(v38)
                                                                    }
                                                                    v4 = (fi63 / ((245) | 1))
                                                                    _ = fi63
                                                                }
                                                            }
                                                            v5 -= 259
                                                        }
                                                    case "hello":
                                                        var v68 int8 = -128
                                                        _ = v68
                                                    }
                                                    var v69 uint8 = v6
                                                    _ = v69
                                                }
                                            }
                                            if (65535 <= 0) {
                                                v70 := GenS9{F0: 488, F1: -116, F2: 127}
                                                var v71 int8 = 127
                                                v13 += (v2 + v2)
                                                _ = v70
                                                _ = v71
                                            } else {
                                                var v72 uint8 = 67
                                                _ = v72
                                            }
                                            v1 = append(v1, v38)
                                        }
                                        v73 := GenS2{F0: false}
                                        _ = v73
                                    } else {
                                        var v74 bool = true
                                        fmt.Println(v74, v13)
                                        _ = v74
                                    }
                                    _ = fi49
                                }
                                _ = v48
                            }
                        }
                        for fi75 := 0; fi75 < 3; fi75++ {
                            for fi76 := 0; fi76 < 1; fi76++ {
                                v77 := GenS10{F0: 0}
                                switch "a" {
                                case "":
                                    var v78 float64 = 3.14159
                                    _ = v78
                                case "foo bar":
                                    fmt.Println(v13)
                                }
                                _ = fi76
                                _ = v77
                            }
                            for fi79 := 0; fi79 < 2; fi79++ {
                                v3 = v3
                                if ((!(32767 <= 223)) || false) {
                                    if (float32(0.0) != 1.5) {
                                        v80 := GenS10{F0: 42}
                                        var v81 uint64 = (^(v3))
                                        _ = v80
                                        _ = v81
                                    }
                                } else {
                                    if false {
                                        var v82 uint64 = (v3 >> ((0) & 63))
                                        _ = len(v1)
                                        v83 := GenS9{F0: 784, F1: -57, F2: fi75}
                                        var v84 uint64 = (v3 - (v82 * 0))
                                        var v85 float64 = float64((float32(-2.25) + 1.5))
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                        _ = v85
                                    }
                                    v38 += (fi79 * (fi75 | v38))
                                    switch "" {
                                    case "foo bar":
                                        var v86 float32 = float32(1e10)
                                        _ = v86
                                    case "":
                                        var v87 int = -2147483648
                                        _ = v87
                                    default:
                                        switch fi75 {
                                        case -9223372036854775808:
                                            v38 = fi75
                                        }
                                    }
                                    var v88 int32 = (v2 >> ((971) & 63))
                                    _ = v88
                                }
                                if true {
                                    var v89 string = ""
                                    v90 := GenS7{F0: v2, F1: 32767}
                                    v13 -= 42
                                    _ = v89
                                    _ = v90
                                } else {
                                    v91 := GenS11{F0: (v5 / ((32767) | 1))}
                                    var v92 uint8 = v6
                                    _ = v91
                                    _ = v92
                                }
                                _ = fi79
                            }
                            _ = fi75
                        }
                        _ = v39
                    }
                    for fi93 := 0; fi93 < 3; fi93++ {
                        var v94 int64 = 115
                        v13 = (-(v2))
                        v95 := GenS7{F0: v13, F1: 2}
                        _ = fi93
                        _ = v94
                        _ = v95
                    }
                    _ = v38
                } else {
                    for fi96 := 0; fi96 < 13; fi96++ {
                        v6 -= 127
                        var v97 bool = ((false || false) && (false && (4294967295 == 374)))
                        switch "foo bar" {
                        case "\n\t":
                            var v98 string = "foo bar"
                            _ = v98
                        default:
                            for fi99 := 0; fi99 < 10; fi99++ {
                                v100 := GenS5{F0: 2, F1: 1, F2: -128}
                                var v101 int8 = 0
                                v6 += v6
                                switch "foo bar" {
                                case "":
                                    switch "a" {
                                    case "hello":
                                        if true {
                                            v102 := GenS12{F0: v3, F1: v101}
                                            _ = v102
                                        }
                                    case "\n\t":
                                        fmt.Println(v101, v3, fi99)
                                    case "foo bar":
                                        if (!false) {
                                            switch ("\n\t" + (("a" + "a") + "\n\t")) {
                                            case "a":
                                                v103 := GenS13{F0: (v101 * (v101 / (((v101 - -128)) | 1)))}
                                                _ = v103
                                            case "\n\t":
                                                if len(v1) > 0 {
                                                    v1[0] = (fi99 / ((fi99) | 1))
                                                }
                                            }
                                            if (!(0 <= 0)) {
                                                fmt.Println(v6, v6)
                                                v2 += -680
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            } else {
                                                v13 -= v13
                                                var v104 string = "hello"
                                                v101 -= v101
                                                var v105 int64 = -231
                                                v13 = (^(v13))
                                                _ = v104
                                                _ = v105
                                            }
                                            var v106 uint = 42
                                            _ = v106
                                        } else {
                                            if ((v2 / ((v13) | 1)) != 42) {
                                                switch fi96 {
                                                case -1:
                                                    var v107 int32 = v13
                                                    _ = v107
                                                }
                                                v13 = (^(v13))
                                                var v108 string = (("a" + "hello") + "\n\t")
                                                _ = v108
                                            }
                                        }
                                    }
                                case "foo bar":
                                    v101 = (v101 - int8(v3))
                                case "\n\t":
                                    v109 := GenS1{F0: 271, F1: -539, F2: v5}
                                    _ = v109
                                default:
                                    v110 := GenS6{F0: 260, F1: 1, F2: -2147483648}
                                    _ = v110
                                }
                                v6 = (v6 % ((2) | 1))
                                _ = fi99
                                _ = v100
                                _ = v101
                            }
                        }
                        _ = fi96
                        _ = v97
                    }
                }
                _ = v37
            }
        case -2147483648:
            var v111 float64 = 0.0
            _ = v111
        default:
            var v112 uint16 = 65535
            _ = v112
        }
    default:
        var v113 uint8 = genFn0(101)
        _ = v113
    }
    if len(v1) > 0 {
        v1[0] = (v4 - (v4 & v4))
    }
    var v114 int32 = v13
    var v115 int16 = v5
    v5 += 691
    v116 := GenS9{F0: 2, F1: -128, F2: (v4 & v4)}
    for fi117 := 0; fi117 < 14; fi117++ {
        var v118 int32 = v2
        var v119 bool = true
        _ = fi117
        _ = v118
        _ = v119
    }
    v115 -= v5
    v120 := GenS4{F0: v4}
    for fi121 := 0; fi121 < 12; fi121++ {
        _ = len(v1)
        if len(v1) > 0 {
            v1[0] = (v4 << ((486) & 63))
        }
        _ = fi121
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v114:", v114)
    fmt.Println("v115:", v115)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v114
    _ = v115
    _ = v116
    _ = v120
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
