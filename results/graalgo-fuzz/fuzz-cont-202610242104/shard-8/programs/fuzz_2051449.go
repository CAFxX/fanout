package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint; F2 int16 }
type GenS4 struct { F0 int32; F1 int; F2 int64 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 bool; F1 uint8; F2 bool }
type GenS7 struct { F0 uint64; F1 int8 }
type GenS8 struct { F0 string; F1 uint }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int32; F1 int; F2 uint64 }
type GenS11 struct { F0 uint; F1 int16 }
type GenS12 struct { F0 int; F1 uint }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 string; F1 bool; F2 int16 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS18 struct { F0 string }
type GenS19 struct { F0 uint32; F1 uint32 }
type GenS20 struct { F0 int16; F1 bool }
type GenS21 struct { F0 string; F1 int16; F2 uint16 }
type GenS22 struct { F0 uint16 }
type GenS23 struct { F0 int64 }
type GenS24 struct { F0 bool }
type GenS25 struct { F0 uint32; F1 int16; F2 int32 }
type GenS26 struct { F0 int; F1 int; F2 int }
type GenS27 struct { F0 uint16 }
type GenS28 struct { F0 string; F1 int8; F2 uint16 }
type GenS29 struct { F0 int64 }
type GenS30 struct { F0 string; F1 uint }
type GenS31 struct { F0 string; F1 uint }
type GenS32 struct { F0 uint; F1 int32 }
type GenS33 struct { F0 string; F1 uint16; F2 uint8 }
type GenS34 struct { F0 uint16 }
type GenS35 struct { F0 uint16; F1 uint; F2 int16 }
type GenS36 struct { F0 int32; F1 int16; F2 uint8 }
type GenS37 struct { F0 int; F1 uint64 }
type GenS38 struct { F0 uint8 }
func genFn0(p0 string) bool {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 3; fi4++ {
        switch 42 {
        case -455:
            p0 = "foo bar"
        case -549:
            for fi5 := 0; fi5 < 2; fi5++ {
                v6 := GenS1{F0: 980}
                v7 := []int{}
                var v8 uint32 = 4294967295
                _ = fi5
                _ = v6
                _ = v7
                _ = v8
            }
        default:
            if ((-84 <= 100) && (1 != 0)) {
                var v9 string = (p0 + p0)
                v9 = "a"
                var v10 bool = (42 == 0)
                _ = v10
                _ = v9
            } else {
                var v11 int64 = -1
                var v12 float32 = 0.1
                _ = v11
                _ = v12
            }
        }
        if false {
            var v13 uint = 42
            _ = v13
        } else {
            fmt.Println(fi4, p0)
        }
        _ = fi4
    }
    switch (p0 + (p0 + p0)) {
    case "\n\t":
        p0 = "a"
    case "foo bar":
        p0 = "\n\t"
    case "":
        var v14 bool = true
        _ = v14
    }
    if (0 != 42) {
        if false {
            if false {
                fmt.Println(p0)
                v15 := []int{-431, 805, 9223372036854775807, -440}
                _ = v15
            }
            v16 := []int{-1, 32767, 0, 2147483647}
            fmt.Println(p0)
            _ = v16
        } else {
            var v17 int32 = 1
            if true {
                var v18 int = 32767
                for fi19 := 0; fi19 < 18; fi19++ {
                    var v20 uint = 2
                    switch "a" {
                    case "hello":
                        if (!false) {
                            switch v18 {
                            case -161:
                                if (-9223372036854775808 >= -1) {
                                    switch p0 {
                                    case "foo bar":
                                        var v21 int = v18
                                        _ = v21
                                    case "\n\t":
                                        v18 += (v18 - (v18 % ((fi19) | 1)))
                                    default:
                                        switch fi19 {
                                        case 9223372036854775807:
                                            var v22 int = (v18 % (((fi19 << ((v20) & 63))) | 1))
                                            _ = v22
                                        default:
                                            var v23 float64 = float64(float32(0.1))
                                            _ = v23
                                        }
                                    }
                                    fmt.Println(fi19, v17)
                                    for fi24 := 0; fi24 < 11; fi24++ {
                                        v25 := []int{60}
                                        var v26 uint = 166
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                    }
                                } else {
                                    v17 += (v17 | 2147483647)
                                }
                            case -246:
                                if true {
                                    p0 = "\n\t"
                                    for fi27 := 0; fi27 < 7; fi27++ {
                                        if (42 != -9223372036854775808) {
                                            if (!true) {
                                                var v28 uint32 = 65535
                                                v17 += (v17 + 479)
                                                if (1 > int64(v20)) {
                                                    v29 := []int{fi27, (fi19 - fi19)}
                                                    _ = v29
                                                }
                                                if true {
                                                    v28 = v28
                                                }
                                                v30 := []int{(fi19 & v18), fi19, fi19, (v18 | fi19)}
                                                if len(v30) >= 2 {
                                                    v30 = v30[1:len(v30)]
                                                }
                                                _ = v28
                                                _ = v30
                                            }
                                        } else {
                                            fmt.Println(p0, fi19, v17)
                                        }
                                        _ = fi27
                                    }
                                    v20 = v20
                                }
                            case -166:
                                switch (p0 + ((("a" + p0) + p0) + p0)) {
                                case "\n\t":
                                    if false {
                                        if true {
                                            for fi31 := 0; fi31 < 11; fi31++ {
                                                if (true && (44 != 0)) {
                                                    v32 := []int{512, (-(fi19)), -9223372036854775808, (fi31 >> (((v20 + 18446744073709551615)) & 63))}
                                                    _ = v32
                                                } else {
                                                    v20 += 255
                                                }
                                                v20 = v20
                                                _ = fi31
                                            }
                                            v20 = 255
                                            p0 = "\n\t"
                                        } else {
                                            fmt.Println(fi19, v18)
                                        }
                                        p0 = "foo bar"
                                    } else {
                                        if true {
                                            v33 := GenS1{F0: 0}
                                            p0 = (p0 + p0)
                                            _ = v33
                                        } else {
                                            if true {
                                                v34 := []int{}
                                                var v35 uint16 = 854
                                                fmt.Println(v35, v35)
                                                _ = v34
                                                _ = v35
                                            } else {
                                                v36 := []int{}
                                                v37 := GenS1{F0: 444}
                                                v20 = (v20 & (v20 * v20))
                                                _ = v36
                                                _ = v37
                                            }
                                            v38 := []int{}
                                            if (true || ((!true) || false)) {
                                                switch (v18 - v18) {
                                                case 1:
                                                    if (100 != 127) {
                                                        v39 := GenS2{F0: v20}
                                                        if (303 == 890) {
                                                            switch fi19 {
                                                            case -2147483648:
                                                                fmt.Println(p0, fi19, p0)
                                                            default:
                                                                fmt.Println(p0, fi19)
                                                            }
                                                            var v40 float32 = (1.5 - (((-2.25 * float32(0.1)) + (1.5 + 0.1)) / 0.1))
                                                            _ = v40
                                                        }
                                                        _ = v39
                                                    } else {
                                                        switch v18 {
                                                        case 9223372036854775807:
                                                            p0 = ("hello" + "\n\t")
                                                        case -134:
                                                            p0 = p0
                                                        default:
                                                            fmt.Println(v17)
                                                        }
                                                        p0 = p0
                                                    }
                                                }
                                            } else {
                                                for fi41 := 0; fi41 < 11; fi41++ {
                                                    if (true && true) {
                                                        var v42 float32 = 3.14159
                                                        if len(v38) > 0 {
                                                            v38[0] = v18
                                                        }
                                                        if (v18 > (v18 | v18)) {
                                                            v38 = append(v38, (fi19 ^ (fi41 | 625)))
                                                        } else {
                                                            p0 = "a"
                                                        }
                                                        _ = v42
                                                    }
                                                    switch (v18 | v18) {
                                                    case 32767:
                                                        p0 = ((("\n\t" + (p0 + "")) + "\n\t") + "foo bar")
                                                    case 0:
                                                        var v43 int = v18
                                                        _ = v43
                                                    default:
                                                        var v44 int8 = -128
                                                        _ = v44
                                                    }
                                                    _ = fi41
                                                }
                                                p0 = (p0 + "\n\t")
                                            }
                                            _ = v38
                                        }
                                        var v45 int16 = -853
                                        var v46 int = (v18 >> ((v20) & 63))
                                        _ = v45
                                        _ = v46
                                    }
                                case "foo bar":
                                    var v47 float32 = 3.14159
                                    _ = v47
                                default:
                                    v48 := []int{(fi19 - (-(fi19))), (fi19 % (((v18 / ((fi19) | 1))) | 1)), fi19}
                                    _ = v48
                                }
                            default:
                                if false {
                                    var v49 uint32 = 586
                                    _ = v49
                                }
                            }
                            v50 := []int{}
                            _ = v50
                        }
                    default:
                        p0 = "a"
                    }
                    _ = fi19
                    _ = v20
                }
                _ = v18
            } else {
                var v51 uint8 = 34
                var v52 int32 = v17
                var v53 uint32 = 0
                _ = v51
                _ = v52
                _ = v53
            }
            for fi54 := 0; fi54 < 18; fi54++ {
                p0 = "a"
                _ = fi54
            }
            _ = v17
        }
    } else {
        p0 = p0
    }
    var v55 float64 = float64(3.14159)
    _ = p0
    _ = v55
    return false
}
func genFn1() bool {
    var v56 float32 = -2.25
    for fi57 := 0; fi57 < 14; fi57++ {
        for fi58 := 0; fi58 < 18; fi58++ {
            if (true && false) {
                if true {
                    if (1 != 4294967295) {
                        switch ("hello" + ("hello" + ((("" + (((("\n\t" + "hello") + "a") + "\n\t") + "")) + ("foo bar" + "")) + (("" + "hello") + "hello")))) {
                        case "a":
                            var v59 uint16 = 42
                            _ = v59
                        case "hello":
                            if true {
                                v60 := GenS1{F0: 755}
                                fmt.Println(fi57)
                                fmt.Println(fi57, fi57, fi58)
                                if (1 < -405) {
                                    v61 := GenS2{F0: 125}
                                    v56 = v56
                                    v62 := []int{1, (fi57 & fi58)}
                                    _ = v61
                                    _ = v62
                                }
                                _ = v60
                            }
                        case "":
                            switch "" {
                            case "hello":
                                v63 := []int{(fi57 * fi57), fi58, fi57, 4294967295}
                                _ = v63
                            case "":
                                v56 -= v56
                            case "foo bar":
                                fmt.Println(v56, v56)
                            default:
                                v56 -= v56
                            }
                        }
                        v64 := GenS3{F0: (false || false), F1: 65535, F2: 1000}
                        v56 -= float32((0.1 + (0.1 * (-2.25 * 1e10))))
                        _ = v64
                    } else {
                        switch fi57 {
                        case 9223372036854775807:
                            v56 = v56
                        case 4294967295:
                            var v65 bool = (2 > 0)
                            _ = v65
                        case 106:
                            fmt.Println(fi58, fi58)
                        }
                        if true {
                            switch fi57 {
                            case 686:
                                v56 += v56
                            case 4294967295:
                                v66 := GenS4{F0: 2147483647, F1: (-(fi57)), F2: -515}
                                _ = v66
                            case 6:
                                fmt.Println(fi58, fi58)
                            }
                        } else {
                            var v67 string = ("\n\t" + "\n\t")
                            v68 := []int{}
                            _ = v67
                            _ = v68
                        }
                        v69 := []int{}
                        _ = v69
                    }
                    v56 += 3.14159
                    var v70 int32 = -1
                    _ = v70
                }
                v56 = v56
                fmt.Println(v56, v56, fi58)
            }
            v56 = (v56 - v56)
            _ = fi58
        }
        _ = fi57
    }
    v56 += v56
    var v71 int64 = 904
    _ = v56
    _ = v71
    return (true && false)
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"\n\t", "a"}
    var v3 float64 = (0.1 / float64(((0.0 / 0.1) - -2.25)))
    if false {
        if (true && (false || true)) {
            switch "\n\t" {
            case "\n\t":
                if (32767 >= 1) {
                    v3 = v3
                } else {
                    v72 := GenS5{F0: 44, F1: 127}
                    var v73 bool = true
                    _ = v72
                    _ = v73
                }
            case "hello":
                if true {
                    _ = len(v1)
                    fmt.Println(v3)
                }
            default:
                v74 := GenS3{F0: true, F1: 65535, F2: 32767}
                _ = v74
            }
            fmt.Println(v3)
        } else {
            v75 := GenS6{F0: (1.5 < (0.0 / 3.14159)), F1: 252, F2: (!(141 != -1))}
            _ = v75
        }
        var v76 float64 = 0.1
        v3 += v76
        _ = v76
    } else {
        fmt.Println(v3)
        v3 -= 0.0
        if ((1 != -9223372036854775808) && true) {
            var v77 int8 = -99
            v3 = v3
            _ = v77
        }
    }
    _ = len(v2)
    for fi78 := 0; fi78 < 18; fi78++ {
        var v79 int64 = 1
        _ = fi78
        _ = v79
    }
    var v80 string = "\n\t"
    fmt.Println(v3, v3)
    v3 += v3
    var v81 float32 = (((float32(v3) - 0.0) / 3.14159) / float32(v3))
    fmt.Println(v81)
    v80 = (("" + v80) + "foo bar")
    if false {
        if true {
            v81 += v81
        }
        var v82 float64 = v3
        fmt.Println(v82, v80, v80)
        _ = v82
    }
    switch 100 {
    case 32768:
        for fi83 := 0; fi83 < 1; fi83++ {
            v80 = v80
            var v84 uint = 552
            for fi85 := 0; fi85 < 18; fi85++ {
                v86 := GenS7{F0: 0, F1: 0}
                switch fi85 {
                case 914:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = fi85
                _ = v86
            }
            _ = fi83
            _ = v84
        }
    case -56:
        v80 = ((v80 + ((v80 + "a") + "hello")) + (v80 + "foo bar"))
    case -613:
        v81 = v81
    default:
        var v87 uint64 = 18446744073709551615
        _ = v87
    }
    v3 = (v3 / 0.1)
    switch 127 {
    case 2147483647:
        if true {
            if (916 >= 614) {
                for fi88 := 0; fi88 < 10; fi88++ {
                    v2 = append(v2, ((((v80 + v80) + v80) + "\n\t") + "foo bar"))
                    switch fi88 {
                    case 32767:
                        for fi89 := 0; fi89 < 15; fi89++ {
                            var v90 bool = (!true)
                            v91 := v1[615]
                            _ = fi89
                            _ = v90
                            _ = v91
                        }
                    case -9223372036854775808:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case -360:
                        v92 := GenS4{F0: -682, F1: fi88, F2: 0}
                        _ = v92
                    }
                    _ = fi88
                }
            }
            switch ((((v80 + (v80 + ((((v80 + v80) + v80) + v80) + v80))) + v80) + ("hello" + ("foo bar" + (("foo bar" + "foo bar") + "\n\t")))) + v80) {
            case "foo bar":
                var v93 uint32 = 981
                _ = v93
            case "hello":
                var v94 bool = false
                _ = v94
            case "\n\t":
                v95 := GenS8{F0: v80, F1: 4294967295}
                _ = v95
            }
            for fi96 := 0; fi96 < 7; fi96++ {
                if false {
                    if (1 < 0) {
                        switch (fi96 - -1) {
                        case 0:
                            var v97 uint8 = 8
                            _ = v97
                        }
                        for fi98 := 0; fi98 < 10; fi98++ {
                            fmt.Println(v81)
                            _ = fi98
                        }
                        for fi99 := 0; fi99 < 2; fi99++ {
                            if false {
                                v80 = (v80 + v80)
                                fmt.Println(fi99, v80, fi96)
                            } else {
                                if (!(!false)) {
                                    fmt.Println(v80)
                                    if true {
                                        if ((!false) || false) {
                                            v3 = 3.14159
                                            switch v80 {
                                            case "\n\t":
                                                switch (fi96 - fi96) {
                                                case 128:
                                                    v100 := GenS5{F0: 127, F1: 81}
                                                    _ = v100
                                                case 42:
                                                    v81 += (v81 * v81)
                                                case 1:
                                                    if ((fi99 / ((127) | 1)) < -2147483648) {
                                                        v81 = 0.1
                                                        v81 = 1.5
                                                        var v101 float64 = float64(v81)
                                                        fmt.Println(fi96)
                                                        fmt.Println(fi99)
                                                        var v102 int16 = -908
                                                        if (v81 > ((v81 + v81) - 3.14159)) {
                                                            fmt.Println(v101)
                                                            var v103 float32 = v81
                                                            var v104 int32 = 2147483647
                                                            _ = v103
                                                            _ = v104
                                                        } else {
                                                            v101 += (v101 * float64((v81 * float32(1e10))))
                                                            _ = len(v2)
                                                        }
                                                        _ = v101
                                                        _ = v102
                                                    }
                                                }
                                            case "a":
                                                if true {
                                                    var v105 string = v80
                                                    v80 = "foo bar"
                                                    var v106 float64 = ((v3 * float64(1.5)) + (v3 * v3))
                                                    _ = v105
                                                    _ = v106
                                                } else {
                                                    v80 = v80
                                                }
                                            default:
                                                if false {
                                                    v107 := GenS9{F0: 42}
                                                    switch fi96 {
                                                    case -1:
                                                        v80 = "foo bar"
                                                    case 9223372036854775807:
                                                        v2 = append(v2, ((v80 + (v80 + "hello")) + v80))
                                                    }
                                                    v108 := v1[230]
                                                    _ = v107
                                                    _ = v108
                                                }
                                                if (65535 >= 0) {
                                                    if true {
                                                        v81 = v81
                                                        var v109 int16 = -923
                                                        v1[4294967295] = ((fi99 ^ fi99) < 32767)
                                                        _ = v109
                                                    }
                                                    switch v80 {
                                                    case "\n\t":
                                                        var v110 string = v80
                                                        _ = v110
                                                    case "a":
                                                        _ = len(v1)
                                                    default:
                                                        switch fi96 {
                                                        case 1:
                                                            if false {
                                                                switch -856 {
                                                                case -9223372036854775808:
                                                                    v111 := GenS9{F0: 65535}
                                                                    v81 -= 3.14159
                                                                    var v112 uint32 = 0
                                                                    v80 = "a"
                                                                    var v113 uint8 = 95
                                                                    _ = v111
                                                                    _ = v112
                                                                    _ = v113
                                                                case 496:
                                                                    var v114 uint64 = 0
                                                                    _ = v114
                                                                case -67:
                                                                    v3 = 1e10
                                                                default:
                                                                    if true {
                                                                        var v115 int64 = 1
                                                                        _ = v115
                                                                    }
                                                                }
                                                                if (18446744073709551615 < 4294967295) {
                                                                    switch (((v80 + ("hello" + "a")) + "a") + v80) {
                                                                    case "a":
                                                                        v3 = v3
                                                                    case "hello":
                                                                        if true {
                                                                            var v116 float32 = (v81 - (v81 - v81))
                                                                            v117 := GenS1{F0: 65535}
                                                                            _ = v116
                                                                            _ = v117
                                                                        } else {
                                                                            if false {
                                                                                switch v80 {
                                                                                case "hello":
                                                                                    var v118 bool = false
                                                                                    _ = v118
                                                                                case "\n\t":
                                                                                    var v119 float64 = (1e10 / v3)
                                                                                    _ = v119
                                                                                case "":
                                                                                    if (32767 <= -207) {
                                                                                        var v120 uint32 = 450
                                                                                        v120 = v120
                                                                                        delete(v1, (v120 % ((v120) | 1)))
                                                                                        _ = v120
                                                                                    } else {
                                                                                        fmt.Println(v80, v80, fi99)
                                                                                        fmt.Println(fi99, fi99, v81)
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v2 = append(v2, "hello")
                                                                                fmt.Println(v80, v81)
                                                                                fmt.Println(fi99, v3)
                                                                            }
                                                                            if (false && false) {
                                                                                v80 = v80
                                                                                v81 = v81
                                                                            } else {
                                                                                v81 = 1.5
                                                                                if (!(42 != 42)) {
                                                                                    switch "foo bar" {
                                                                                    case "a":
                                                                                        var v121 uint = 255
                                                                                        _ = v121
                                                                                    case "hello":
                                                                                        v80 = (v80 + ("foo bar" + v80))
                                                                                    }
                                                                                    if (32767 == 0) {
                                                                                        if (!(0 > 1)) {
                                                                                            var v122 uint16 = 0
                                                                                            _ = v122
                                                                                        } else {
                                                                                            if (2 <= 32767) {
                                                                                                var v123 int32 = 0
                                                                                                _ = v123
                                                                                            } else {
                                                                                                fmt.Println(fi99)
                                                                                            }
                                                                                            var v124 float32 = v81
                                                                                            _ = v124
                                                                                        }
                                                                                        var v125 int64 = 0
                                                                                        v125 -= v125
                                                                                        _ = v125
                                                                                    }
                                                                                } else {
                                                                                    switch (^(fi96)) {
                                                                                    case 647:
                                                                                        v2 = append(v2, v80)
                                                                                    case -9223372036854775808:
                                                                                        v80 = v80
                                                                                    }
                                                                                    var v126 string = "foo bar"
                                                                                    v127 := GenS5{F0: 42, F1: 127}
                                                                                    var v128 uint = 677
                                                                                    _ = v126
                                                                                    _ = v127
                                                                                    _ = v128
                                                                                }
                                                                            }
                                                                        }
                                                                    case "\n\t":
                                                                        if true {
                                                                            switch v80 {
                                                                            case "a":
                                                                                v81 = v81
                                                                            case "":
                                                                                switch (fi96 / (((fi99 / ((284) | 1))) | 1)) {
                                                                                case 32767:
                                                                                    v81 -= (v81 / v81)
                                                                                case 128:
                                                                                    if (-962 == (fi96 / ((127) | 1))) {
                                                                                        if (!((v3 == 1.5) && true)) {
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                        } else {
                                                                                            switch v80 {
                                                                                            case "a":
                                                                                                var v129 bool = false
                                                                                                _ = v129
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v81, v3)
                                                                                        var v130 string = ("a" + (v80 + v80))
                                                                                        fmt.Println(fi96, v130)
                                                                                        _ = v130
                                                                                    }
                                                                                case 1:
                                                                                    fmt.Println(fi96, v81)
                                                                                default:
                                                                                    v81 += (v81 - (0.0 - (v81 / 0.1)))
                                                                                }
                                                                            default:
                                                                                if true {
                                                                                    fmt.Println(v3, v3)
                                                                                    _ = len(v1)
                                                                                    _ = len(v1)
                                                                                } else {
                                                                                    v131 := v1[1]
                                                                                    v3 = v3
                                                                                    _ = v131
                                                                                }
                                                                            }
                                                                            if true {
                                                                                fmt.Println(fi99, v3)
                                                                                if false {
                                                                                    var v132 int32 = 0
                                                                                    _ = v132
                                                                                } else {
                                                                                    fmt.Println(fi96, fi96, fi99)
                                                                                    var v133 string = v80
                                                                                    var v134 int16 = -78
                                                                                    var v135 uint = 2
                                                                                    var v136 int64 = 9223372036854775807
                                                                                    if (v3 >= ((v3 / (-2.25 - v3)) * v3)) {
                                                                                        v134 = 1
                                                                                        v136 = (v136 & -1)
                                                                                        if (970 != 1) {
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = "foo bar"
                                                                                            }
                                                                                        } else {
                                                                                            v3 = v3
                                                                                            var v137 bool = (0 > 2147483648)
                                                                                            v138 := GenS10{F0: 1, F1: 32767, F2: 2}
                                                                                            _ = v137
                                                                                            _ = v138
                                                                                        }
                                                                                    } else {
                                                                                        v3 -= (v3 + 1e10)
                                                                                        fmt.Println(v135, fi96)
                                                                                    }
                                                                                    if (!(((1e10 * (-2.25 + v81)) * 3.14159) > v81)) {
                                                                                        var v139 string = ((v80 + "foo bar") + "hello")
                                                                                        v140 := GenS11{F0: (^(v135)), F1: -1}
                                                                                        v80 = v133
                                                                                        _ = v139
                                                                                        _ = v140
                                                                                    } else {
                                                                                        if ((false && false) || (0.1 > v81)) {
                                                                                            _ = len(v2)
                                                                                        } else {
                                                                                            fmt.Println(fi96)
                                                                                            if false {
                                                                                                v141 := GenS12{F0: fi96, F1: 4294967295}
                                                                                                var v142 float64 = v3
                                                                                                _ = v141
                                                                                                _ = v142
                                                                                            } else {
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = "\n\t"
                                                                                                }
                                                                                                delete(v1, 65535)
                                                                                                v143 := GenS11{F0: v135, F1: (v134 / ((v134) | 1))}
                                                                                                var v144 bool = false
                                                                                                v136 += v136
                                                                                                v145 := GenS13{F0: 0}
                                                                                                switch "" {
                                                                                                case "\n\t":
                                                                                                    v136 = (v136 << (((v135 % ((895) | 1))) & 63))
                                                                                                case "foo bar":
                                                                                                    switch (((v80 + v80) + (v80 + "foo bar")) + ((((v133 + (((v80 + (((v80 + "hello") + "a") + v133)) + "\n\t") + (v80 + (v80 + ((v80 + v80) + (("hello" + "foo bar") + (v80 + v80))))))) + "a") + ("hello" + "a")) + (v80 + (v80 + "a")))) {
                                                                                                    case "\n\t":
                                                                                                        switch "foo bar" {
                                                                                                        case "\n\t":
                                                                                                            fmt.Println(v80, fi96, v3)
                                                                                                        case "a":
                                                                                                            var v146 float64 = 0.1
                                                                                                            _ = v146
                                                                                                        default:
                                                                                                            switch 4294967295 {
                                                                                                            case -2147483648:
                                                                                                                if true {
                                                                                                                    if v144 {
                                                                                                                        var v147 int32 = int32(v136)
                                                                                                                        _ = v147
                                                                                                                    } else {
                                                                                                                        var v148 int8 = -51
                                                                                                                        _ = v148
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        var v149 int64 = (v136 + (^(v136)))
                                                                                                                        _ = v149
                                                                                                                    } else {
                                                                                                                        fmt.Println(v80, v136, v81)
                                                                                                                    }
                                                                                                                    v134 -= (v134 % (((v134 << (((v135 - v135)) & 63))) | 1))
                                                                                                                    v134 = (^(v134))
                                                                                                                    v80 = (v80 + v80)
                                                                                                                    var v150 uint64 = 42
                                                                                                                    _ = v150
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    case "":
                                                                                                        if true {
                                                                                                            switch (fi99 | fi99) {
                                                                                                            case 42:
                                                                                                                v134 -= v134
                                                                                                            case -2147483648:
                                                                                                                var v151 bool = false
                                                                                                                _ = v151
                                                                                                            }
                                                                                                            fmt.Println(v144)
                                                                                                        }
                                                                                                        switch (fi96 << (((v135 + (v135 & (^(v135))))) & 63)) {
                                                                                                        case 2147483647:
                                                                                                            if (42 == -1) {
                                                                                                                var v152 int32 = 42
                                                                                                                v133 = ""
                                                                                                                if (128 != 153) {
                                                                                                                    delete(v1, 136)
                                                                                                                    fmt.Println(v134, v136)
                                                                                                                    v134 = -522
                                                                                                                } else {
                                                                                                                    if v144 {
                                                                                                                        if true {
                                                                                                                            if v144 {
                                                                                                                                var v153 bool = (true || v144)
                                                                                                                                if v153 {
                                                                                                                                    v144 = v153
                                                                                                                                    if false {
                                                                                                                                        if ((!v144) || ((v134 + -32768) == 1000)) {
                                                                                                                                            v2 = append(v2, (("\n\t" + ("foo bar" + v133)) + "hello"))
                                                                                                                                            v135 -= v135
                                                                                                                                        } else {
                                                                                                                                            switch (fi96 % ((fi96) | 1)) {
                                                                                                                                            case 1:
                                                                                                                                                switch v133 {
                                                                                                                                                case "\n\t":
                                                                                                                                                    v133 = v133
                                                                                                                                                }
                                                                                                                                                if false {
                                                                                                                                                    v154 := GenS8{F0: (v133 + v80), F1: v135}
                                                                                                                                                    _ = v154
                                                                                                                                                } else {
                                                                                                                                                    var v155 int32 = v152
                                                                                                                                                    var v156 int = (fi99 / (((fi96 * fi96)) | 1))
                                                                                                                                                    var v157 uint16 = 59
                                                                                                                                                    switch ((v80 + "\n\t") + "hello") {
                                                                                                                                                    case "\n\t":
                                                                                                                                                        if (v153 || (((1.5 + v3) + (1e10 / (v3 - (v3 + v3)))) > v3)) {
                                                                                                                                                            if (v144 || (!(v134 < v134))) {
                                                                                                                                                                v136 = (^(v136))
                                                                                                                                                                v2 = append(v2, v133)
                                                                                                                                                            }
                                                                                                                                                            if ((v135 >> (((v135 | v135)) & 63)) > v135) {
                                                                                                                                                                if v153 {
                                                                                                                                                                    v158 := GenS9{F0: 586}
                                                                                                                                                                    var v159 uint64 = uint64(v157)
                                                                                                                                                                    _ = v158
                                                                                                                                                                    _ = v159
                                                                                                                                                                }
                                                                                                                                                                if true {
                                                                                                                                                                    if v153 {
                                                                                                                                                                        v160 := GenS9{F0: 65535}
                                                                                                                                                                        var v161 uint64 = 4294967295
                                                                                                                                                                        if (65535 < (^(v157))) {
                                                                                                                                                                            v157 = v157
                                                                                                                                                                        } else {
                                                                                                                                                                            var v162 int64 = (v136 << ((323) & 63))
                                                                                                                                                                            _ = v162
                                                                                                                                                                        }
                                                                                                                                                                        _ = v160
                                                                                                                                                                        _ = v161
                                                                                                                                                                    } else {
                                                                                                                                                                        var v163 uint16 = v157
                                                                                                                                                                        var v164 int16 = (v134 & 775)
                                                                                                                                                                        _ = v163
                                                                                                                                                                        _ = v164
                                                                                                                                                                    }
                                                                                                                                                                } else {
                                                                                                                                                                    v152 = (v155 << (((v135 << ((0) & 63))) & 63))
                                                                                                                                                                }
                                                                                                                                                                if ((v144 && (-2.25 != ((v3 + (((-2.25 + (v3 + ((0.1 / (3.14159 * v3)) + v3))) * v3) + v3)) / -2.25))) && false) {
                                                                                                                                                                    v165 := GenS5{F0: -128, F1: 42}
                                                                                                                                                                    _ = v165
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    _ = v155
                                                                                                                                                    _ = v156
                                                                                                                                                    _ = v157
                                                                                                                                                }
                                                                                                                                                var v166 int = fi99
                                                                                                                                                v167 := GenS14{F0: 32767}
                                                                                                                                                _ = v166
                                                                                                                                                _ = v167
                                                                                                                                            case 4294967295:
                                                                                                                                                v168 := GenS10{F0: 1, F1: -55, F2: 42}
                                                                                                                                                _ = v168
                                                                                                                                            case 42:
                                                                                                                                                if (9223372036854775807 == v136) {
                                                                                                                                                    v169 := GenS10{F0: (v152 | v152), F1: (fi99 + (fi96 % ((int(v152)) | 1))), F2: 2}
                                                                                                                                                    v133 = v80
                                                                                                                                                    _ = v169
                                                                                                                                                } else {
                                                                                                                                                    fmt.Println(v3, v3, v80)
                                                                                                                                                    fmt.Println(v81, fi99, v133)
                                                                                                                                                    fmt.Println(v153, v135, v133)
                                                                                                                                                }
                                                                                                                                                v170 := GenS13{F0: v134}
                                                                                                                                                if (false || (v3 > (3.14159 - (v3 - v3)))) {
                                                                                                                                                    v171 := GenS9{F0: 494}
                                                                                                                                                    var v172 string = v80
                                                                                                                                                    _ = v171
                                                                                                                                                    _ = v172
                                                                                                                                                } else {
                                                                                                                                                    switch fi96 {
                                                                                                                                                    case 1:
                                                                                                                                                        var v173 uint16 = 242
                                                                                                                                                        _ = v173
                                                                                                                                                    }
                                                                                                                                                    var v174 uint16 = 1
                                                                                                                                                    _ = v174
                                                                                                                                                }
                                                                                                                                                _ = v170
                                                                                                                                            default:
                                                                                                                                                var v175 uint16 = 937
                                                                                                                                                _ = v175
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        if len(v2) > 0 {
                                                                                                                                            v2[0] = (v133 + v133)
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    v2 = append(v2, "hello")
                                                                                                                                }
                                                                                                                                _ = v153
                                                                                                                            } else {
                                                                                                                                fmt.Println(v144, fi96, fi96)
                                                                                                                            }
                                                                                                                            var v176 uint32 = 553
                                                                                                                            var v177 int16 = 991
                                                                                                                            _ = v176
                                                                                                                            _ = v177
                                                                                                                        } else {
                                                                                                                            v144 = v144
                                                                                                                        }
                                                                                                                        var v178 string = ((("hello" + (((v80 + (("hello" + "a") + ("foo bar" + ((v80 + "a") + v80)))) + (v133 + "hello")) + (((v133 + v80) + ((v133 + v80) + (("a" + v133) + v80))) + (((v80 + (((v80 + v80) + v80) + ((v80 + "") + "\n\t"))) + v80) + (v133 + (v80 + "hello")))))) + v133) + "\n\t")
                                                                                                                        var v179 uint64 = 18446744073709551615
                                                                                                                        _ = v178
                                                                                                                        _ = v179
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v152
                                                                                                            }
                                                                                                        default:
                                                                                                            v144 = v144
                                                                                                        }
                                                                                                        switch fi99 {
                                                                                                        case -2147483648:
                                                                                                            v135 = 42
                                                                                                        case 775:
                                                                                                            var v180 float64 = ((v3 * v3) / 3.14159)
                                                                                                            _ = v180
                                                                                                        default:
                                                                                                            switch fi96 {
                                                                                                            case 42:
                                                                                                                if v144 {
                                                                                                                    _ = len(v1)
                                                                                                                    v181 := v1[2147483647]
                                                                                                                    v81 = v81
                                                                                                                    v80 = v80
                                                                                                                    switch 9223372036854775807 {
                                                                                                                    case 564:
                                                                                                                        v136 -= (^(v136))
                                                                                                                    case 128:
                                                                                                                        if v144 {
                                                                                                                            if (("hello" != ("\n\t" + v133)) && v181) {
                                                                                                                                v182 := GenS15{F0: v80, F1: v144, F2: v134}
                                                                                                                                v136 = (v136 / (((^(v136))) | 1))
                                                                                                                                if (false || true) {
                                                                                                                                    if v181 {
                                                                                                                                        var v183 uint16 = 0
                                                                                                                                        v184 := GenS10{F0: 0, F1: (fi99 << (((^(v135))) & 63)), F2: 18446744073709551615}
                                                                                                                                        fmt.Println(fi96, fi96, v181)
                                                                                                                                        _ = v183
                                                                                                                                        _ = v184
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                _ = v182
                                                                                                                            }
                                                                                                                            v185 := GenS6{F0: (!(!(true && (!v144)))), F1: 255, F2: ((42 < fi96) && false)}
                                                                                                                            _ = v185
                                                                                                                        } else {
                                                                                                                            if (false || (!((v134 == v134) && false))) {
                                                                                                                                v135 = (v135 / ((v135) | 1))
                                                                                                                                v135 += (v135 & v135)
                                                                                                                                v181 = (!((((0.0 + (((float32((v3 + float64(v81))) - -2.25) + (0.1 - v81)) + ((v81 + v81) - (1e10 * 1e10)))) * v81) - (v81 - ((float32((0.0 - float64(v81))) / 0.1) + v81))) != -2.25))
                                                                                                                            }
                                                                                                                            if v181 {
                                                                                                                                v186 := GenS9{F0: 4294967295}
                                                                                                                                v144 = (v144 || true)
                                                                                                                                switch "a" {
                                                                                                                                case "a":
                                                                                                                                    fmt.Println(v135, fi99, fi99)
                                                                                                                                    if (65535 != 32767) {
                                                                                                                                        v81 = (0.1 + 1.5)
                                                                                                                                        switch (("foo bar" + v133) + ("hello" + "a")) {
                                                                                                                                        case "hello":
                                                                                                                                            if ((v136 & v136) == 42) {
                                                                                                                                                if v181 {
                                                                                                                                                    v144 = v144
                                                                                                                                                    v81 -= v81
                                                                                                                                                    var v187 int8 = 100
                                                                                                                                                    _ = v187
                                                                                                                                                } else {
                                                                                                                                                    var v188 int8 = 127
                                                                                                                                                    v134 = -34
                                                                                                                                                    _ = v188
                                                                                                                                                }
                                                                                                                                                _ = len(v2)
                                                                                                                                                if true {
                                                                                                                                                    v189 := GenS2{F0: 785}
                                                                                                                                                    var v190 int16 = 42
                                                                                                                                                    v191 := GenS10{F0: 42, F1: 193, F2: 2}
                                                                                                                                                    _ = v189
                                                                                                                                                    _ = v190
                                                                                                                                                    _ = v191
                                                                                                                                                } else {
                                                                                                                                                    v133 = "\n\t"
                                                                                                                                                    if len(v2) >= 2 {
                                                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                                if v144 {
                                                                                                                                                    switch (fi96 ^ 1) {
                                                                                                                                                    case 1:
                                                                                                                                                        v192 := GenS16{F0: 1}
                                                                                                                                                        if (4 == 127) {
                                                                                                                                                            if ("" >= v80) {
                                                                                                                                                                _ = len(v2)
                                                                                                                                                                v144 = v181
                                                                                                                                                            } else {
                                                                                                                                                                if (v144 && true) {
                                                                                                                                                                    var v193 uint32 = 42
                                                                                                                                                                    switch "" {
                                                                                                                                                                    case "":
                                                                                                                                                                        var v194 int8 = 126
                                                                                                                                                                        _ = v194
                                                                                                                                                                    case "foo bar":
                                                                                                                                                                        v195 := GenS13{F0: int16(v193)}
                                                                                                                                                                        switch (fi96 >> (((v135 << ((1) & 63))) & 63)) {
                                                                                                                                                                        case 42:
                                                                                                                                                                            delete(v1, (^(v193)))
                                                                                                                                                                        default:
                                                                                                                                                                            v196 := GenS9{F0: (v193 * (v193 % (((v193 + v193)) | 1)))}
                                                                                                                                                                            _ = v196
                                                                                                                                                                        }
                                                                                                                                                                        v197 := GenS8{F0: v133, F1: v135}
                                                                                                                                                                        _ = v195
                                                                                                                                                                        _ = v197
                                                                                                                                                                    }
                                                                                                                                                                    if v144 {
                                                                                                                                                                        fmt.Println(v133)
                                                                                                                                                                        var v198 int8 = 0
                                                                                                                                                                        _ = v198
                                                                                                                                                                    }
                                                                                                                                                                    _ = v193
                                                                                                                                                                } else {
                                                                                                                                                                    var v199 uint32 = 612
                                                                                                                                                                    v200 := GenS16{F0: (fi99 - int(v199))}
                                                                                                                                                                    var v201 bool = v181
                                                                                                                                                                    if ((v134 > v134) && (((-2.25 >= (0.0 * v3)) || ((v144 && v201) && true)) && (0 >= 164))) {
                                                                                                                                                                        v181 = v201
                                                                                                                                                                    } else {
                                                                                                                                                                        v3 -= ((v3 - (float64(v81) - (v3 / v3))) / v3)
                                                                                                                                                                        v133 = ""
                                                                                                                                                                        if len(v2) >= 2 {
                                                                                                                                                                            v2 = v2[1:len(v2)]
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                    _ = v199
                                                                                                                                                                    _ = v200
                                                                                                                                                                    _ = v201
                                                                                                                                                                }
                                                                                                                                                                switch (fi96 | -9223372036854775808) {
                                                                                                                                                                case 817:
                                                                                                                                                                    fmt.Println(v80, v80)
                                                                                                                                                                case 29:
                                                                                                                                                                    if ((!(459 > 73)) || (uint16(v135) != 65535)) {
                                                                                                                                                                        var v202 bool = false
                                                                                                                                                                        v136 = v136
                                                                                                                                                                        v203 := GenS3{F0: (!v181), F1: v135, F2: v134}
                                                                                                                                                                        var v204 int = -432
                                                                                                                                                                        v3 -= v3
                                                                                                                                                                        _ = v202
                                                                                                                                                                        _ = v203
                                                                                                                                                                        _ = v204
                                                                                                                                                                    } else {
                                                                                                                                                                        v144 = v181
                                                                                                                                                                        v1[1] = true
                                                                                                                                                                        var v205 uint64 = 42
                                                                                                                                                                        v134 += -1
                                                                                                                                                                        v206 := GenS17{F0: (v136 - (v136 << ((2) & 63))), F1: 67, F2: 32767}
                                                                                                                                                                        fmt.Println(v3)
                                                                                                                                                                        _ = len(v2)
                                                                                                                                                                        _ = v205
                                                                                                                                                                        _ = v206
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                                if v181 {
                                                                                                                                                                    v1[2147483647] = ((!v144) && v144)
                                                                                                                                                                    v135 = (v135 ^ 65535)
                                                                                                                                                                    var v207 int64 = v136
                                                                                                                                                                    _ = v207
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                            v208 := v1[63]
                                                                                                                                                            if true {
                                                                                                                                                                var v209 bool = false
                                                                                                                                                                _ = v209
                                                                                                                                                            } else {
                                                                                                                                                                if v208 {
                                                                                                                                                                    v144 = (!(v3 == 3.14159))
                                                                                                                                                                    switch v80 {
                                                                                                                                                                    case "foo bar":
                                                                                                                                                                        if (4294967295 != 65535) {
                                                                                                                                                                            if ((!(1 != (v134 % (((v134 >> ((v135) & 63))) | 1)))) || false) {
                                                                                                                                                                                if (!((v136 - 42) == v136)) {
                                                                                                                                                                                    v208 = v181
                                                                                                                                                                                    switch (fi96 >> (((v135 - v135)) & 63)) {
                                                                                                                                                                                    case -1:
                                                                                                                                                                                        var v210 int32 = -1
                                                                                                                                                                                        _ = v210
                                                                                                                                                                                    case 32767:
                                                                                                                                                                                        var v211 int = fi99
                                                                                                                                                                                        _ = v211
                                                                                                                                                                                    default:
                                                                                                                                                                                        var v212 bool = v181
                                                                                                                                                                                        _ = v212
                                                                                                                                                                                    }
                                                                                                                                                                                }
                                                                                                                                                                                v181 = true
                                                                                                                                                                            }
                                                                                                                                                                            var v213 int32 = 2147483647
                                                                                                                                                                            var v214 float64 = (v3 + 0.0)
                                                                                                                                                                            _ = v213
                                                                                                                                                                            _ = v214
                                                                                                                                                                        } else {
                                                                                                                                                                            var v215 string = ("foo bar" + (v133 + v80))
                                                                                                                                                                            var v216 string = "hello"
                                                                                                                                                                            v217 := GenS18{F0: ((v80 + "") + (v215 + v133))}
                                                                                                                                                                            _ = v215
                                                                                                                                                                            _ = v216
                                                                                                                                                                            _ = v217
                                                                                                                                                                        }
                                                                                                                                                                    default:
                                                                                                                                                                        v218 := GenS19{F0: 2147483648, F1: 220}
                                                                                                                                                                        _ = v218
                                                                                                                                                                    }
                                                                                                                                                                    fmt.Println(v208)
                                                                                                                                                                }
                                                                                                                                                                if (!false) {
                                                                                                                                                                    switch ("a" + "a") {
                                                                                                                                                                    case "":
                                                                                                                                                                        v133 = v80
                                                                                                                                                                    case "a":
                                                                                                                                                                        var v219 int32 = 0
                                                                                                                                                                        _ = v219
                                                                                                                                                                    }
                                                                                                                                                                } else {
                                                                                                                                                                    fmt.Println(v135, fi96, v208)
                                                                                                                                                                    var v220 uint64 = 4294967295
                                                                                                                                                                    var v221 string = "\n\t"
                                                                                                                                                                    var v222 int64 = v136
                                                                                                                                                                    var v223 uint = (v135 * (v135 / ((v135) | 1)))
                                                                                                                                                                    var v224 bool = v181
                                                                                                                                                                    if v181 {
                                                                                                                                                                        fmt.Println(v223, v220)
                                                                                                                                                                        if len(v2) > 0 {
                                                                                                                                                                            v2[0] = v133
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                    _ = v220
                                                                                                                                                                    _ = v221
                                                                                                                                                                    _ = v222
                                                                                                                                                                    _ = v223
                                                                                                                                                                    _ = v224
                                                                                                                                                                }
                                                                                                                                                                if (42 < 1) {
                                                                                                                                                                    _ = len(v1)
                                                                                                                                                                    v144 = v208
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                            var v225 int64 = v136
                                                                                                                                                            var v226 int16 = (v134 / ((v134) | 1))
                                                                                                                                                            var v227 string = v133
                                                                                                                                                            _ = v208
                                                                                                                                                            _ = v225
                                                                                                                                                            _ = v226
                                                                                                                                                            _ = v227
                                                                                                                                                        } else {
                                                                                                                                                            var v228 int32 = 1
                                                                                                                                                            _ = v228
                                                                                                                                                        }
                                                                                                                                                        _ = len(v2)
                                                                                                                                                        if (false && v144) {
                                                                                                                                                            switch fi96 {
                                                                                                                                                            case -1:
                                                                                                                                                                v229 := GenS11{F0: v135, F1: (v134 & -1)}
                                                                                                                                                                _ = v229
                                                                                                                                                            case 1:
                                                                                                                                                                fmt.Println(v136, fi96)
                                                                                                                                                            case -2147483648:
                                                                                                                                                                if false {
                                                                                                                                                                    fmt.Println(v136, v81, v80)
                                                                                                                                                                }
                                                                                                                                                            default:
                                                                                                                                                                v181 = true
                                                                                                                                                            }
                                                                                                                                                            v230 := GenS20{F0: -950, F1: v144}
                                                                                                                                                            _ = v230
                                                                                                                                                        } else {
                                                                                                                                                            if ((v136 >= (v136 & (v136 % ((v136) | 1)))) || ((((935 <= 2) || (false || v144)) || (v134 >= v134)) || false)) {
                                                                                                                                                                _ = len(v1)
                                                                                                                                                                if (2 == 42) {
                                                                                                                                                                    v133 = v80
                                                                                                                                                                }
                                                                                                                                                            } else {
                                                                                                                                                                var v231 uint32 = 2147483647
                                                                                                                                                                _ = v231
                                                                                                                                                            }
                                                                                                                                                            var v232 uint16 = 0
                                                                                                                                                            v233 := GenS18{F0: "\n\t"}
                                                                                                                                                            _ = v232
                                                                                                                                                            _ = v233
                                                                                                                                                        }
                                                                                                                                                        fmt.Println(v136, v136)
                                                                                                                                                        v134 = (v134 + v134)
                                                                                                                                                        _ = v192
                                                                                                                                                    }
                                                                                                                                                    v136 -= (-(v136))
                                                                                                                                                    if v144 {
                                                                                                                                                        v1[48] = true
                                                                                                                                                        switch (fi99 / ((fi96) | 1)) {
                                                                                                                                                        case 32768:
                                                                                                                                                            v234 := GenS21{F0: v80, F1: v134, F2: 2}
                                                                                                                                                            _ = v234
                                                                                                                                                        case 319:
                                                                                                                                                            var v235 uint16 = 255
                                                                                                                                                            _ = v235
                                                                                                                                                        }
                                                                                                                                                        switch v80 {
                                                                                                                                                        case "a":
                                                                                                                                                            if v181 {
                                                                                                                                                                var v236 uint = 2
                                                                                                                                                                v133 = v80
                                                                                                                                                                _ = v236
                                                                                                                                                            }
                                                                                                                                                        case "hello":
                                                                                                                                                            fmt.Println(v133, v80, v133)
                                                                                                                                                            v133 = "a"
                                                                                                                                                        }
                                                                                                                                                        v133 = "hello"
                                                                                                                                                    } else {
                                                                                                                                                        v134 -= (^(v134))
                                                                                                                                                    }
                                                                                                                                                    v237 := GenS18{F0: (v133 + "")}
                                                                                                                                                    switch (fi99 | (fi99 % (((fi99 << ((65535) & 63))) | 1))) {
                                                                                                                                                    case 32768:
                                                                                                                                                        v2 = append(v2, ((v80 + ("a" + (("foo bar" + ("hello" + v133)) + "hello"))) + v80))
                                                                                                                                                    case 42:
                                                                                                                                                        if v181 {
                                                                                                                                                            v3 = (v3 * (float64((v81 + (0.1 / v81))) + v3))
                                                                                                                                                        } else {
                                                                                                                                                            v80 = "foo bar"
                                                                                                                                                        }
                                                                                                                                                    case 920:
                                                                                                                                                        if (!(false || false)) {
                                                                                                                                                            switch ((v80 + (v133 + v133)) + v133) {
                                                                                                                                                            case "hello":
                                                                                                                                                                _ = len(v1)
                                                                                                                                                            case "":
                                                                                                                                                                switch v133 {
                                                                                                                                                                case "a":
                                                                                                                                                                    fmt.Println(fi96, v81, v181)
                                                                                                                                                                case "":
                                                                                                                                                                    v144 = v181
                                                                                                                                                                default:
                                                                                                                                                                    var v238 int32 = 2147483647
                                                                                                                                                                    _ = v238
                                                                                                                                                                }
                                                                                                                                                            case "a":
                                                                                                                                                                switch fi99 {
                                                                                                                                                                case 127:
                                                                                                                                                                    v144 = (!((v135 * v135) > (v135 * (v135 | (^(v135))))))
                                                                                                                                                                default:
                                                                                                                                                                    var v239 uint32 = 65535
                                                                                                                                                                    _ = v239
                                                                                                                                                                }
                                                                                                                                                            default:
                                                                                                                                                                v181 = (!v144)
                                                                                                                                                            }
                                                                                                                                                            if (!(!((0.0 + v3) != 0.0))) {
                                                                                                                                                                var v240 bool = (((v133 + v80) + ("hello" + "\n\t")) >= "hello")
                                                                                                                                                                _ = v240
                                                                                                                                                            } else {
                                                                                                                                                                var v241 int8 = 100
                                                                                                                                                                if len(v2) > 0 {
                                                                                                                                                                    v2[0] = v133
                                                                                                                                                                }
                                                                                                                                                                _ = v241
                                                                                                                                                            }
                                                                                                                                                            if ((v135 % (((v135 & 4294967295)) | 1)) == (v135 + v135)) {
                                                                                                                                                                var v242 float64 = 0.1
                                                                                                                                                                v243 := GenS11{F0: (v135 >> ((204) & 63)), F1: (v134 ^ -331)}
                                                                                                                                                                var v244 float32 = v81
                                                                                                                                                                _ = v242
                                                                                                                                                                _ = v243
                                                                                                                                                                _ = v244
                                                                                                                                                            } else {
                                                                                                                                                                delete(v1, 4294967295)
                                                                                                                                                                v245 := GenS2{F0: v135}
                                                                                                                                                                v246 := v1[2147483647]
                                                                                                                                                                if false {
                                                                                                                                                                    v81 += v81
                                                                                                                                                                } else {
                                                                                                                                                                    v181 = true
                                                                                                                                                                    fmt.Println(fi99, v81, v181)
                                                                                                                                                                }
                                                                                                                                                                v246 = (v181 && (!(98 != 127)))
                                                                                                                                                                _ = v245
                                                                                                                                                                _ = v246
                                                                                                                                                            }
                                                                                                                                                        } else {
                                                                                                                                                            v247 := GenS22{F0: 32767}
                                                                                                                                                            v248 := v1[42]
                                                                                                                                                            if (1e10 != float32(-2.25)) {
                                                                                                                                                                _ = len(v1)
                                                                                                                                                                v144 = v248
                                                                                                                                                            } else {
                                                                                                                                                                var v249 uint16 = 42
                                                                                                                                                                var v250 uint = 18446744073709551615
                                                                                                                                                                _ = v249
                                                                                                                                                                _ = v250
                                                                                                                                                            }
                                                                                                                                                            _ = v247
                                                                                                                                                            _ = v248
                                                                                                                                                        }
                                                                                                                                                        if false {
                                                                                                                                                            v181 = true
                                                                                                                                                        } else {
                                                                                                                                                            if len(v2) >= 2 {
                                                                                                                                                                v2 = v2[1:len(v2)]
                                                                                                                                                            }
                                                                                                                                                            v134 -= int16(fi99)
                                                                                                                                                        }
                                                                                                                                                        v251 := GenS23{F0: (-(v136))}
                                                                                                                                                        _ = v251
                                                                                                                                                    default:
                                                                                                                                                        if len(v2) >= 2 {
                                                                                                                                                            v2 = v2[1:len(v2)]
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    var v252 float32 = (v81 / v81)
                                                                                                                                                    _ = v237
                                                                                                                                                    _ = v252
                                                                                                                                                } else {
                                                                                                                                                    if (!v181) {
                                                                                                                                                        var v253 string = v80
                                                                                                                                                        _ = v253
                                                                                                                                                    }
                                                                                                                                                    if ((float32(-2.25) + v81) >= 1.5) {
                                                                                                                                                        if true {
                                                                                                                                                            var v254 float32 = v81
                                                                                                                                                            v3 += ((0.1 * (-2.25 - v3)) / v3)
                                                                                                                                                            if len(v2) >= 2 {
                                                                                                                                                                v2 = v2[1:len(v2)]
                                                                                                                                                            }
                                                                                                                                                            _ = v254
                                                                                                                                                        }
                                                                                                                                                        v136 = v136
                                                                                                                                                        var v255 int = (fi99 - (fi96 * -2147483648))
                                                                                                                                                        _ = v255
                                                                                                                                                    }
                                                                                                                                                    var v256 bool = (127 <= 1)
                                                                                                                                                    _ = v256
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        var v257 int = fi96
                                                                                                                                        switch 4294967295 {
                                                                                                                                        case 367:
                                                                                                                                            v81 = -2.25
                                                                                                                                        case 2147483647:
                                                                                                                                            v258 := GenS4{F0: -992, F1: (fi99 ^ -2147483648), F2: v136}
                                                                                                                                            _ = v258
                                                                                                                                        default:
                                                                                                                                            v144 = (v144 || v144)
                                                                                                                                        }
                                                                                                                                        v259 := GenS24{F0: (!(v3 != v3))}
                                                                                                                                        _ = v257
                                                                                                                                        _ = v259
                                                                                                                                    } else {
                                                                                                                                        if true {
                                                                                                                                            fmt.Println(v181)
                                                                                                                                            v133 = ((("" + (((v80 + (v133 + v133)) + v133) + ("foo bar" + v133))) + (v80 + v133)) + "")
                                                                                                                                            if (0 == 42) {
                                                                                                                                                v181 = (((v136 ^ v136) != (v136 / (((-(v136))) | 1))) || v181)
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    v144 = (!((!((1 == 2) && (!false))) || (false && false)))
                                                                                                                                case "foo bar":
                                                                                                                                    v260 := GenS5{F0: -59, F1: 106}
                                                                                                                                    _ = v260
                                                                                                                                default:
                                                                                                                                    v80 = v80
                                                                                                                                }
                                                                                                                                _ = v186
                                                                                                                            } else {
                                                                                                                                var v261 uint64 = 1
                                                                                                                                if (-54 > -91) {
                                                                                                                                    var v262 uint16 = 46
                                                                                                                                    _ = v262
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v135)
                                                                                                                                    if ((fi96 + fi99) != fi96) {
                                                                                                                                        fmt.Println(v181, v144, v261)
                                                                                                                                        if v144 {
                                                                                                                                            if len(v2) >= 2 {
                                                                                                                                                v2 = v2[1:len(v2)]
                                                                                                                                            }
                                                                                                                                            fmt.Println(v134, v81)
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        var v263 float64 = v3
                                                                                                                                        var v264 int = fi96
                                                                                                                                        _ = v263
                                                                                                                                        _ = v264
                                                                                                                                    }
                                                                                                                                    _ = len(v2)
                                                                                                                                }
                                                                                                                                var v265 bool = v144
                                                                                                                                _ = v261
                                                                                                                                _ = v265
                                                                                                                            }
                                                                                                                            var v266 int8 = 122
                                                                                                                            _ = v266
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v181
                                                                                                                }
                                                                                                                if false {
                                                                                                                    var v267 string = v133
                                                                                                                    if false {
                                                                                                                        if ((127 < 1) && (!(int(v135) != (^(fi96))))) {
                                                                                                                            var v268 uint8 = 42
                                                                                                                            if v144 {
                                                                                                                                switch ((((v80 + "foo bar") + "hello") + v133) + "hello") {
                                                                                                                                case "foo bar":
                                                                                                                                    v1[4294967295] = (!v144)
                                                                                                                                case "\n\t":
                                                                                                                                    if ("a" != "hello") {
                                                                                                                                        fmt.Println(v144, v134, v81)
                                                                                                                                        var v269 uint = v135
                                                                                                                                        v268 += 1
                                                                                                                                        v267 = "hello"
                                                                                                                                        _ = v269
                                                                                                                                    } else {
                                                                                                                                        var v270 bool = (!false)
                                                                                                                                        v81 = 3.14159
                                                                                                                                        _ = v270
                                                                                                                                    }
                                                                                                                                case "":
                                                                                                                                    v134 = (v134 << ((v135) & 63))
                                                                                                                                }
                                                                                                                                var v271 string = "hello"
                                                                                                                                _ = v271
                                                                                                                            }
                                                                                                                            _ = v268
                                                                                                                        }
                                                                                                                        fmt.Println(v81)
                                                                                                                    } else {
                                                                                                                        var v272 uint64 = 0
                                                                                                                        _ = v272
                                                                                                                    }
                                                                                                                    _ = v267
                                                                                                                } else {
                                                                                                                    var v273 uint8 = 0
                                                                                                                    switch -9223372036854775808 {
                                                                                                                    case 32767:
                                                                                                                        v133 = "foo bar"
                                                                                                                    }
                                                                                                                    _ = v273
                                                                                                                }
                                                                                                            }
                                                                                                            v274 := v1[1]
                                                                                                            _ = v274
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        v3 += v3
                                                                                                    default:
                                                                                                        var v275 int32 = 0
                                                                                                        _ = v275
                                                                                                    }
                                                                                                }
                                                                                                v276 := GenS21{F0: "", F1: v134, F2: 2}
                                                                                                if true {
                                                                                                    v277 := GenS11{F0: 65535, F1: v134}
                                                                                                    if (v144 && v144) {
                                                                                                        switch ("hello" + ("hello" + v133)) {
                                                                                                        case "hello":
                                                                                                            var v278 string = (v80 + "\n\t")
                                                                                                            _ = v278
                                                                                                        }
                                                                                                        if true {
                                                                                                            if v144 {
                                                                                                                v144 = v144
                                                                                                            } else {
                                                                                                                var v279 float32 = 1.5
                                                                                                                switch fi99 {
                                                                                                                case -192:
                                                                                                                    var v280 float32 = 1.5
                                                                                                                    _ = v280
                                                                                                                }
                                                                                                                v281 := v1[0]
                                                                                                                v144 = (((v281 && (!v281)) && false) || true)
                                                                                                                v282 := GenS25{F0: 348, F1: (v134 << (((v135 << (((v135 / ((v135) | 1))) & 63))) & 63)), F2: 2147483647}
                                                                                                                _ = v279
                                                                                                                _ = v281
                                                                                                                _ = v282
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v283 uint32 = 0
                                                                                                            _ = v283
                                                                                                        }
                                                                                                        v284 := v1[2147483647]
                                                                                                        _ = v284
                                                                                                    } else {
                                                                                                        fmt.Println(v81, fi99, v3)
                                                                                                        var v285 float64 = 0.1
                                                                                                        var v286 float64 = (1e10 - (0.0 / ((3.14159 / (((0.1 + (v3 - float64(v81))) - (v3 / (v285 / v285))) + -2.25)) + v3)))
                                                                                                        _ = v285
                                                                                                        _ = v286
                                                                                                    }
                                                                                                    var v287 int = fi99
                                                                                                    _ = v277
                                                                                                    _ = v287
                                                                                                } else {
                                                                                                    v288 := GenS16{F0: fi99}
                                                                                                    var v289 int = 127
                                                                                                    _ = v288
                                                                                                    _ = v289
                                                                                                }
                                                                                                _ = v143
                                                                                                _ = v144
                                                                                                _ = v145
                                                                                                _ = v276
                                                                                            }
                                                                                            fmt.Println(v81, fi96)
                                                                                            switch (fi99 % (((fi99 * (^(fi99)))) | 1)) {
                                                                                            case 9223372036854775807:
                                                                                                _ = len(v2)
                                                                                            case 32767:
                                                                                                var v290 float64 = float64(((((float32(v3) + 0.0) - float32(1e10)) / 3.14159) + ((v81 / v81) * (((1e10 - (1.5 - v81)) / 1.0) - 1.5))))
                                                                                                _ = v290
                                                                                            case -2147483648:
                                                                                                if true {
                                                                                                    var v291 bool = false
                                                                                                    _ = v291
                                                                                                } else {
                                                                                                    if (457 <= 0) {
                                                                                                        _ = len(v1)
                                                                                                        v135 -= v135
                                                                                                        v134 = (v134 | v134)
                                                                                                    }
                                                                                                    if (true && (-406 < v136)) {
                                                                                                        v292 := GenS4{F0: 42, F1: fi99, F2: v136}
                                                                                                        if true {
                                                                                                            v3 -= 3.14159
                                                                                                            var v293 uint8 = 255
                                                                                                            v3 = float64(3.14159)
                                                                                                            _ = v293
                                                                                                        }
                                                                                                        _ = v292
                                                                                                    } else {
                                                                                                        var v294 uint16 = 861
                                                                                                        var v295 string = v133
                                                                                                        v135 = 2
                                                                                                        _ = v294
                                                                                                        _ = v295
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                fmt.Println(v80, v136)
                                                                                            }
                                                                                        }
                                                                                        if (!true) {
                                                                                            v134 = (v134 ^ v134)
                                                                                            if (-32768 <= (v134 >> (((v135 >> ((v135) & 63))) & 63))) {
                                                                                                v80 = v133
                                                                                                var v296 string = ""
                                                                                                var v297 uint = (v135 >> ((v135) & 63))
                                                                                                if true {
                                                                                                    if false {
                                                                                                        fmt.Println(v297)
                                                                                                        fmt.Println(fi96, v296)
                                                                                                        v298 := GenS1{F0: 65535}
                                                                                                        _ = v298
                                                                                                    } else {
                                                                                                        if (((184 <= 18446744073709551615) || (v136 < (^(v136)))) && true) {
                                                                                                            switch -1 {
                                                                                                            case -2147483648:
                                                                                                                var v299 string = "foo bar"
                                                                                                                _ = v299
                                                                                                            case 835:
                                                                                                                v300 := GenS16{F0: (-(fi99))}
                                                                                                                _ = v300
                                                                                                            default:
                                                                                                                if (v135 <= 18446744073709551615) {
                                                                                                                    if len(v2) > 0 {
                                                                                                                        v2[0] = "hello"
                                                                                                                    }
                                                                                                                    if (2147483647 > 0) {
                                                                                                                        v301 := GenS26{F0: (fi96 % (((fi96 | 128)) | 1)), F1: 613, F2: fi96}
                                                                                                                        _ = v301
                                                                                                                    } else {
                                                                                                                        var v302 uint32 = 42
                                                                                                                        _ = v302
                                                                                                                    }
                                                                                                                    switch fi96 {
                                                                                                                    case -305:
                                                                                                                        var v303 int64 = (v136 << ((v135) & 63))
                                                                                                                        _ = v303
                                                                                                                    default:
                                                                                                                        if (true || true) {
                                                                                                                            var v304 int8 = 0
                                                                                                                            switch fi96 {
                                                                                                                            case 4294967295:
                                                                                                                                var v305 uint = v135
                                                                                                                                _ = v305
                                                                                                                            case -784:
                                                                                                                                switch (v296 + ("a" + v296)) {
                                                                                                                                case "a":
                                                                                                                                    switch (fi99 >> ((v297) & 63)) {
                                                                                                                                    case 994:
                                                                                                                                        v80 = "foo bar"
                                                                                                                                    case -178:
                                                                                                                                        if ((v3 - v3) != v3) {
                                                                                                                                            v306 := GenS21{F0: ("a" + "a"), F1: 42, F2: 1}
                                                                                                                                            if (!true) {
                                                                                                                                                v136 -= v136
                                                                                                                                                if true {
                                                                                                                                                    v307 := GenS17{F0: (v136 ^ (v136 | v136)), F1: 2147483648, F2: 144}
                                                                                                                                                    var v308 float32 = float32(1e10)
                                                                                                                                                    _ = v307
                                                                                                                                                    _ = v308
                                                                                                                                                } else {
                                                                                                                                                    v309 := v1[2]
                                                                                                                                                    fmt.Println(v297, v81, v134)
                                                                                                                                                    fmt.Println(fi99, v304, fi99)
                                                                                                                                                    _ = v309
                                                                                                                                                }
                                                                                                                                            } else {
                                                                                                                                                v310 := GenS27{F0: 2}
                                                                                                                                                fmt.Println(v133, v3, v135)
                                                                                                                                                _ = v310
                                                                                                                                            }
                                                                                                                                            _ = v306
                                                                                                                                        } else {
                                                                                                                                            if (!false) {
                                                                                                                                                var v311 int = (fi96 ^ -9223372036854775808)
                                                                                                                                                _ = v311
                                                                                                                                            }
                                                                                                                                            if ((v304 <= -128) && false) {
                                                                                                                                                v136 = v136
                                                                                                                                                if (v136 != 42) {
                                                                                                                                                    var v312 float32 = v81
                                                                                                                                                    v313 := GenS1{F0: 0}
                                                                                                                                                    switch v80 {
                                                                                                                                                    case "":
                                                                                                                                                        var v314 uint16 = 255
                                                                                                                                                        v2 = append(v2, "a")
                                                                                                                                                        if (v297 != v297) {
                                                                                                                                                            var v315 float64 = 1.5
                                                                                                                                                            if (42 > 2147483648) {
                                                                                                                                                                v316 := GenS28{F0: v133, F1: (^(v304)), F2: v314}
                                                                                                                                                                _ = v316
                                                                                                                                                            } else {
                                                                                                                                                                _ = len(v1)
                                                                                                                                                            }
                                                                                                                                                            var v317 string = v133
                                                                                                                                                            fmt.Println(v317)
                                                                                                                                                            _ = len(v1)
                                                                                                                                                            _ = v315
                                                                                                                                                            _ = v317
                                                                                                                                                        }
                                                                                                                                                        _ = v314
                                                                                                                                                    case "foo bar":
                                                                                                                                                        v3 += v3
                                                                                                                                                    case "\n\t":
                                                                                                                                                        v3 += (((float64(v312) - (v3 - (0.0 - 3.14159))) - (v3 / (float64(((v81 * 0.0) / 0.1)) + v3))) - (0.0 + v3))
                                                                                                                                                    default:
                                                                                                                                                        v2 = append(v2, (v80 + v80))
                                                                                                                                                    }
                                                                                                                                                    _ = v312
                                                                                                                                                    _ = v313
                                                                                                                                                }
                                                                                                                                                v135 += (^(v297))
                                                                                                                                            } else {
                                                                                                                                                v1[2147483647] = (255 >= 28)
                                                                                                                                                v296 = v80
                                                                                                                                                v80 = "hello"
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                case "hello":
                                                                                                                                    var v318 uint8 = 1
                                                                                                                                    _ = v318
                                                                                                                                default:
                                                                                                                                    v304 -= (v304 - (v304 % (((v304 / ((v304) | 1))) | 1)))
                                                                                                                                }
                                                                                                                            case -1:
                                                                                                                                var v319 int16 = -1
                                                                                                                                _ = v319
                                                                                                                            default:
                                                                                                                                var v320 uint16 = 2
                                                                                                                                _ = v320
                                                                                                                            }
                                                                                                                            var v321 uint8 = 42
                                                                                                                            if true {
                                                                                                                                _ = len(v1)
                                                                                                                                if ((false || false) && false) {
                                                                                                                                    _ = len(v1)
                                                                                                                                    fmt.Println(fi99)
                                                                                                                                } else {
                                                                                                                                    fmt.Println(fi99, v136)
                                                                                                                                }
                                                                                                                                switch (-(fi99)) {
                                                                                                                                case 105:
                                                                                                                                    var v322 string = ""
                                                                                                                                    _ = v322
                                                                                                                                case -1:
                                                                                                                                    var v323 int64 = 136
                                                                                                                                    v297 -= (v135 >> (((v135 - (v135 | (^(v297))))) & 63))
                                                                                                                                    var v324 uint64 = 42
                                                                                                                                    _ = v323
                                                                                                                                    _ = v324
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                if (v296 == (v133 + (("a" + "\n\t") + v133))) {
                                                                                                                                    var v325 int64 = (v136 << ((v297) & 63))
                                                                                                                                    if false {
                                                                                                                                        var v326 int16 = v134
                                                                                                                                        _ = len(v2)
                                                                                                                                        _ = v326
                                                                                                                                    }
                                                                                                                                    _ = v325
                                                                                                                                }
                                                                                                                                fmt.Println(fi96)
                                                                                                                            }
                                                                                                                            _ = v304
                                                                                                                            _ = v321
                                                                                                                        } else {
                                                                                                                            fmt.Println(fi96, v3, v133)
                                                                                                                            fmt.Println(v134, v297)
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            var v327 int64 = v136
                                                                                                            if false {
                                                                                                                var v328 int = (fi99 & -2147483648)
                                                                                                                if len(v2) > 0 {
                                                                                                                    v2[0] = "\n\t"
                                                                                                                }
                                                                                                                _ = v328
                                                                                                            }
                                                                                                            _ = v327
                                                                                                        }
                                                                                                    }
                                                                                                    var v329 string = ("a" + ("hello" + v133))
                                                                                                    _ = v329
                                                                                                } else {
                                                                                                    var v330 int32 = -1
                                                                                                    v296 = (v296 + "a")
                                                                                                    delete(v1, uint32(v297))
                                                                                                    _ = v330
                                                                                                }
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                if false {
                                                                                                    var v331 uint = v135
                                                                                                    v133 = ((((((v80 + v80) + (v80 + v80)) + ((v133 + "") + "hello")) + ("" + "")) + (((v80 + v80) + "") + v296)) + v296)
                                                                                                    v332 := GenS29{F0: -685}
                                                                                                    _ = v331
                                                                                                    _ = v332
                                                                                                } else {
                                                                                                    v333 := GenS7{F0: 2, F1: -1}
                                                                                                    _ = v333
                                                                                                }
                                                                                                _ = v296
                                                                                                _ = v297
                                                                                            }
                                                                                        } else {
                                                                                            v133 = v133
                                                                                        }
                                                                                        v135 = (v135 % ((65535) | 1))
                                                                                    }
                                                                                    if true {
                                                                                        v136 += 234
                                                                                    }
                                                                                    var v334 uint32 = 705
                                                                                    _ = v133
                                                                                    _ = v134
                                                                                    _ = v135
                                                                                    _ = v136
                                                                                    _ = v334
                                                                                }
                                                                                v3 -= 0.1
                                                                            } else {
                                                                                var v335 uint32 = 65535
                                                                                if (!(false && true)) {
                                                                                    var v336 uint32 = v335
                                                                                    v336 += (v335 + v336)
                                                                                    v337 := GenS22{F0: 65535}
                                                                                    _ = v336
                                                                                    _ = v337
                                                                                }
                                                                                var v338 int32 = -1
                                                                                var v339 uint64 = 911
                                                                                v340 := GenS30{F0: (v80 + ("\n\t" + "\n\t")), F1: 949}
                                                                                _ = v335
                                                                                _ = v338
                                                                                _ = v339
                                                                                _ = v340
                                                                            }
                                                                            if false {
                                                                                v1[157] = (v3 > 0.1)
                                                                                var v341 float64 = v3
                                                                                _ = v341
                                                                            } else {
                                                                                var v342 int16 = 1
                                                                                v3 += v3
                                                                                _ = v342
                                                                            }
                                                                            v3 -= (1e10 + v3)
                                                                        } else {
                                                                            v343 := GenS31{F0: (v80 + ""), F1: 255}
                                                                            _ = v343
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            var v344 string = v80
                                                                            v345 := GenS24{F0: false}
                                                                            var v346 uint64 = 42
                                                                            _ = v344
                                                                            _ = v345
                                                                            _ = v346
                                                                        }
                                                                    }
                                                                    var v347 uint16 = 642
                                                                    _ = v347
                                                                } else {
                                                                    switch v80 {
                                                                    case "":
                                                                        var v348 uint8 = 66
                                                                        switch v80 {
                                                                        case "foo bar":
                                                                            if false {
                                                                                var v349 int16 = 1
                                                                                _ = v349
                                                                            } else {
                                                                                v350 := GenS18{F0: v80}
                                                                                var v351 float64 = (((3.14159 + ((float64(v81) + ((v3 - 1.5) * ((v3 - v3) - v3))) - 0.0)) - v3) * ((v3 - v3) + ((((float64(float32((v3 + 0.0))) / (1.5 - 0.1)) + -2.25) - float64(1.5)) - v3)))
                                                                                _ = v350
                                                                                _ = v351
                                                                            }
                                                                        default:
                                                                            v81 = v81
                                                                        }
                                                                        var v352 int8 = 127
                                                                        _ = len(v1)
                                                                        switch 1 {
                                                                        case 2147483647:
                                                                            v80 = v80
                                                                        case 472:
                                                                            var v353 string = (v80 + (("hello" + v80) + "foo bar"))
                                                                            _ = v353
                                                                        }
                                                                        v348 = v348
                                                                        _ = v348
                                                                        _ = v352
                                                                    case "\n\t":
                                                                        v80 = v80
                                                                    case "foo bar":
                                                                        delete(v1, 1)
                                                                    default:
                                                                        v80 = "\n\t"
                                                                    }
                                                                    if (!(!(1 >= -2147483648))) {
                                                                        fmt.Println(v3, v3, fi96)
                                                                        v81 = 0.0
                                                                        var v354 uint64 = 696
                                                                        _ = v354
                                                                    }
                                                                    v355 := GenS32{F0: 18446744073709551615, F1: 686}
                                                                    _ = v355
                                                                }
                                                            } else {
                                                                var v356 int16 = 1
                                                                if true {
                                                                    v2 = append(v2, v80)
                                                                    v3 = 1e10
                                                                    switch 127 {
                                                                    case -127:
                                                                        v3 += 0.0
                                                                        v3 = 1e10
                                                                        v357 := GenS17{F0: 0, F1: 383, F2: 42}
                                                                        switch "hello" {
                                                                        case "":
                                                                            var v358 float64 = (0.0 + ((float64(v81) / (((float64((float32(v3) + (v81 * v81))) + v3) - -2.25) / (((v3 + float64(-2.25)) - 0.0) - 1e10))) - (1e10 + ((v3 / v3) / (float64((1e10 * (((0.0 * 0.1) / float32(float64(v81))) - v81))) - ((v3 - float64(v81)) + v3))))))
                                                                            _ = v358
                                                                        case "a":
                                                                            v81 = 1e10
                                                                        }
                                                                        _ = v357
                                                                    case 779:
                                                                        var v359 int16 = -374
                                                                        _ = v359
                                                                    case 9223372036854775807:
                                                                        if true {
                                                                            var v360 uint = 888
                                                                            if false {
                                                                                var v361 uint32 = 2147483647
                                                                                v362 := GenS22{F0: 65535}
                                                                                _ = v361
                                                                                _ = v362
                                                                            }
                                                                            v363 := GenS33{F0: "\n\t", F1: 0, F2: 42}
                                                                            _ = v360
                                                                            _ = v363
                                                                        }
                                                                        v364 := GenS32{F0: 355, F1: 627}
                                                                        _ = v364
                                                                    default:
                                                                        fmt.Println(v80, v81, v3)
                                                                    }
                                                                } else {
                                                                    switch v80 {
                                                                    case "\n\t":
                                                                        var v365 int64 = 42
                                                                        _ = v365
                                                                    default:
                                                                        v366 := GenS29{F0: -9223372036854775808}
                                                                        var v367 int8 = 100
                                                                        v80 = ("\n\t" + (v80 + v80))
                                                                        switch v80 {
                                                                        case "foo bar":
                                                                            var v368 uint = 65535
                                                                            switch fi96 {
                                                                            case -9223372036854775808:
                                                                                var v369 int32 = 42
                                                                                if false {
                                                                                    var v370 uint64 = 907
                                                                                    v369 = v369
                                                                                    _ = v370
                                                                                }
                                                                                if false {
                                                                                    switch (-(fi99)) {
                                                                                    case -2147483648:
                                                                                        delete(v1, 42)
                                                                                    }
                                                                                    v80 = (v80 + v80)
                                                                                    switch (v80 + v80) {
                                                                                    case "":
                                                                                        v367 = v367
                                                                                    case "\n\t":
                                                                                        if (0 < 128) {
                                                                                            fmt.Println(v368, fi99)
                                                                                            var v371 uint64 = 4294967295
                                                                                            var v372 uint16 = 684
                                                                                            _ = v371
                                                                                            _ = v372
                                                                                        }
                                                                                    case "a":
                                                                                        v356 = (v356 >> (((v368 + (v368 % ((42) | 1)))) & 63))
                                                                                    }
                                                                                }
                                                                                _ = v369
                                                                            }
                                                                            fmt.Println(v80)
                                                                            _ = v368
                                                                        }
                                                                        _ = v366
                                                                        _ = v367
                                                                    }
                                                                    fmt.Println(v356)
                                                                    fmt.Println(v81, v81, fi96)
                                                                    delete(v1, 42)
                                                                }
                                                                if true {
                                                                    v3 = 1e10
                                                                } else {
                                                                    v3 += v3
                                                                    if len(v2) > 0 {
                                                                        v2[0] = v80
                                                                    }
                                                                    v373 := GenS4{F0: 42, F1: fi99, F2: 42}
                                                                    _ = v373
                                                                }
                                                                _ = v356
                                                            }
                                                        case -2147483648:
                                                            v81 = 3.14159
                                                        default:
                                                            v374 := GenS34{F0: 0}
                                                            _ = v374
                                                        }
                                                    }
                                                    _ = len(v1)
                                                } else {
                                                    if false {
                                                        v375 := GenS35{F0: 358, F1: 255, F2: -32768}
                                                        var v376 int = 2147483647
                                                        v377 := GenS36{F0: int32(v376), F1: 42, F2: 0}
                                                        _ = v375
                                                        _ = v376
                                                        _ = v377
                                                    }
                                                    v3 += v3
                                                }
                                            }
                                            v80 = v80
                                        }
                                        var v378 int8 = 95
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v379 int64 = 1
                                            _ = v379
                                        case "hello":
                                            if (false && true) {
                                                fmt.Println(fi96, v3, v3)
                                                _ = len(v2)
                                            } else {
                                                var v380 int64 = -9223372036854775808
                                                switch 0 {
                                                case 32768:
                                                    _ = len(v1)
                                                case 2147483647:
                                                    var v381 uint8 = 127
                                                    _ = v381
                                                }
                                                var v382 float64 = v3
                                                _ = v380
                                                _ = v382
                                            }
                                        case "":
                                            v80 = v80
                                        default:
                                            v3 += v3
                                        }
                                        v1[65535] = (false && true)
                                        if (-556 != -9223372036854775808) {
                                            if false {
                                                var v383 string = v80
                                                _ = v383
                                            }
                                        } else {
                                            v378 -= v378
                                        }
                                        _ = v378
                                    }
                                    switch (fi96 ^ fi99) {
                                    case -2147483648:
                                        var v384 int64 = 0
                                        _ = v384
                                    case 2147483647:
                                        switch (("\n\t" + (v80 + (((v80 + v80) + v80) + v80))) + (v80 + ("\n\t" + ("hello" + (v80 + v80))))) {
                                        case "foo bar":
                                            fmt.Println(fi99)
                                        case "\n\t":
                                            if (42 <= 2) {
                                                v80 = v80
                                                if true {
                                                    var v385 int32 = -664
                                                    _ = v385
                                                } else {
                                                    switch (fi99 - fi99) {
                                                    case -9223372036854775808:
                                                        switch 32767 {
                                                        case -2147483648:
                                                            fmt.Println(v3, v81)
                                                        case -110:
                                                            var v386 int = (fi96 & fi99)
                                                            _ = v386
                                                        }
                                                    case 869:
                                                        var v387 uint32 = 0
                                                        _ = v387
                                                    case 4294967295:
                                                        var v388 uint8 = 42
                                                        _ = v388
                                                    }
                                                }
                                                v389 := GenS22{F0: 31}
                                                _ = v389
                                            } else {
                                                var v390 bool = false
                                                v81 -= (1.5 * 3.14159)
                                                switch (("\n\t" + (v80 + "\n\t")) + v80) {
                                                case "\n\t":
                                                    fmt.Println(fi96, v80)
                                                    v391 := GenS2{F0: 42}
                                                    switch fi96 {
                                                    case 222:
                                                        if (18446744073709551615 != 18446744073709551615) {
                                                            v392 := v1[32]
                                                            switch "" {
                                                            case "foo bar":
                                                                v392 = v390
                                                            case "hello":
                                                                fmt.Println(v81)
                                                            default:
                                                                v393 := GenS20{F0: 901, F1: v390}
                                                                _ = v393
                                                            }
                                                            _ = v392
                                                        }
                                                    case -1:
                                                        if v390 {
                                                            if (871 == 100) {
                                                                if true {
                                                                    var v394 uint16 = 0
                                                                    switch (v80 + v80) {
                                                                    case "a":
                                                                        v394 += (v394 << ((0) & 63))
                                                                    case "\n\t":
                                                                        v395 := GenS9{F0: 2}
                                                                        _ = v395
                                                                    default:
                                                                        switch (fi96 - 1) {
                                                                        case 0:
                                                                            var v396 bool = (508 != 916)
                                                                            _ = v396
                                                                        default:
                                                                            v397 := GenS13{F0: -337}
                                                                            _ = v397
                                                                        }
                                                                    }
                                                                    _ = v394
                                                                } else {
                                                                    var v398 uint32 = 457
                                                                    _ = v398
                                                                }
                                                            }
                                                            if (v390 && false) {
                                                                fmt.Println(v80, fi99, v3)
                                                                switch v80 {
                                                                case "foo bar":
                                                                    fmt.Println(v390, v3)
                                                                case "hello":
                                                                    var v399 int16 = 42
                                                                    _ = v399
                                                                case "":
                                                                    if v390 {
                                                                        var v400 int = fi99
                                                                        _ = v400
                                                                    }
                                                                default:
                                                                    v3 += (0.1 + v3)
                                                                }
                                                                if ((-9223372036854775808 > -506) || false) {
                                                                    var v401 float64 = v3
                                                                    var v402 uint64 = 39
                                                                    _ = v401
                                                                    _ = v402
                                                                } else {
                                                                    v403 := GenS37{F0: fi99, F1: 4294967295}
                                                                    v390 = v390
                                                                    _ = v403
                                                                }
                                                            }
                                                        } else {
                                                            v81 = 1.5
                                                            if false {
                                                                var v404 bool = (v390 || v390)
                                                                _ = v404
                                                            } else {
                                                                var v405 uint16 = 0
                                                                fmt.Println(v390, fi96, v405)
                                                                v3 = -2.25
                                                                _ = v405
                                                            }
                                                        }
                                                    default:
                                                        var v406 int16 = 0
                                                        _ = v406
                                                    }
                                                    _ = v391
                                                }
                                                _ = v390
                                            }
                                        }
                                    }
                                }
                                v407 := GenS30{F0: (("" + "a") + v80), F1: 445}
                                _ = v407
                            }
                            v2 = append(v2, v80)
                            v408 := GenS38{F0: 2}
                            _ = fi99
                            _ = v408
                        }
                    }
                    switch (v80 + (v80 + "hello")) {
                    case "":
                        fmt.Println(v81)
                    case "\n\t":
                        v80 = (v80 + v80)
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                } else {
                    v1[0] = (true || true)
                    v409 := GenS25{F0: 1, F1: 643, F2: 2147483647}
                    _ = v409
                }
                if false {
                    if (0 == 1000) {
                        var v410 uint64 = 42
                        if false {
                            var v411 int = fi96
                            _ = v411
                        } else {
                            fmt.Println(v80, v3)
                        }
                        var v412 uint16 = 42
                        _ = v410
                        _ = v412
                    }
                    v413 := GenS37{F0: (fi96 + (fi96 & 2147483647)), F1: 2}
                    var v414 uint32 = 4294967295
                    _ = v413
                    _ = v414
                }
                _ = fi96
            }
        } else {
            v415 := GenS31{F0: v80, F1: 1}
            _ = v415
        }
    case 200:
        _ = len(v2)
    case -9223372036854775808:
        fmt.Println(v3, v80, v81)
    }
    var v416 uint = 1
    switch "\n\t" {
    case "foo bar":
        v81 += v81
    case "\n\t":
        var v417 int8 = -97
        _ = v417
    }
    for fi418 := 0; fi418 < 19; fi418++ {
        v416 = 4294967295
        _ = fi418
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v416:", v416)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v416
    _ = v80
    _ = v81
}
