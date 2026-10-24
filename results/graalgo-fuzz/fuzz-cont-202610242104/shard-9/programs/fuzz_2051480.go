package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int8; F2 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS9 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS10 struct { F0 int; F1 int16; F2 uint8 }
type GenS11 struct { F0 int32; F1 int16 }
type GenS12 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS13 struct { F0 int32; F1 uint32 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 bool; F1 uint64 }
type GenS16 struct { F0 uint64; F1 int64; F2 string }
type GenS17 struct { F0 uint32; F1 bool }
type GenS18 struct { F0 uint16; F1 string }
type GenS19 struct { F0 int }
type GenS20 struct { F0 bool; F1 bool }
type GenS21 struct { F0 string; F1 string; F2 uint64 }
type GenS22 struct { F0 uint64; F1 int8 }
type GenS23 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS24 struct { F0 uint16; F1 int32 }
func genFn0() int64 {
    var v6 uint = 2
    v7 := GenS1{F0: false}
    v8 := GenS2{F0: ""}
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 int32 = -1
        switch -1 {
        case 32767:
            var v11 float32 = (1.5 * (0.0 + 3.14159))
            _ = v11
        default:
            switch (^(fi9)) {
            case -412:
                var v12 int16 = 32767
                _ = v12
            case -357:
                v13 := []int{(fi9 | (fi9 & fi9)), -9223372036854775808, fi9, fi9}
                _ = v13
            default:
                v14 := GenS3{F0: 1, F1: 45, F2: true}
                _ = v14
            }
        }
        _ = fi9
        _ = v10
    }
    _ = v6
    _ = v7
    _ = v8
    return 164
}
func genFn1(p0 int16) bool {
    var v15 uint32 = 778
    var v16 float32 = 3.14159
    _ = p0
    _ = v15
    _ = v16
    return true
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{2}
    var v3 bool = (("\n\t" + ("\n\t" + ("foo bar" + "foo bar"))) < ("hello" + "hello"))
    var v4 int16 = 42
    v5 := make(map[uint32]uint32)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = false
    if (3.14159 < ((1e10 / 1e10) + (0.0 + float32((-2.25 - (0.1 + 3.14159)))))) {
        fmt.Println(v3, v3)
    } else {
        var v17 int = 128
        switch (v17 >> ((255) & 63)) {
        case 4294967295:
            _ = len(v2)
        default:
            v18 := GenS4{F0: (^(v4))}
            _ = v18
        }
        if (!false) {
            v17 -= v17
            switch (("a" + ((("foo bar" + "a") + ((("a" + "\n\t") + "foo bar") + "foo bar")) + "hello")) + "") {
            case "":
                switch v17 {
                case 2147483647:
                    var v19 uint8 = 1
                    _ = v19
                case -728:
                    var v20 int = -2147483648
                    _ = v20
                }
            case "a":
                switch 42 {
                case -312:
                    for fi21 := 0; fi21 < 9; fi21++ {
                        if (v3 && (false || (!(683 > 18446744073709551615)))) {
                            switch 128 {
                            case -1:
                                switch "" {
                                case "a":
                                    for fi22 := 0; fi22 < 14; fi22++ {
                                        delete(v1, 42)
                                        _ = fi22
                                    }
                                case "hello":
                                    v4 += (v4 ^ (^(v4)))
                                case "\n\t":
                                    v3 = true
                                }
                            }
                            v17 = (v17 << ((2) & 63))
                            fmt.Println(fi21)
                        } else {
                            for fi23 := 0; fi23 < 13; fi23++ {
                                if ((true && (v3 && v3)) && (1 < 1)) {
                                    v4 = (v4 & v4)
                                } else {
                                    fmt.Println(v4)
                                    if v3 {
                                        v4 = (-(v4))
                                    }
                                    var v24 int = (fi21 % ((v17) | 1))
                                    _ = v24
                                }
                                v17 -= 2147483647
                                _ = fi23
                            }
                            switch fi21 {
                            case -2147483648:
                                v4 += (v4 + (v4 + 1000))
                            case -9223372036854775808:
                                for fi25 := 0; fi25 < 8; fi25++ {
                                    switch "" {
                                    case "foo bar":
                                        switch -9223372036854775808 {
                                        case 9223372036854775807:
                                            if (v3 || (v3 && (2147483648 == 644))) {
                                                v4 = v4
                                            }
                                        case 4294967295:
                                            fmt.Println(v3, v4, v3)
                                        default:
                                            var v26 int16 = -628
                                            _ = v26
                                        }
                                    case "hello":
                                        if v3 {
                                            v17 = (v17 * (fi25 | (-(fi21))))
                                            if (v4 > v4) {
                                                var v27 float64 = -2.25
                                                v4 = v4
                                                if v3 {
                                                    switch "" {
                                                    case "a":
                                                        v3 = ((-20 > -1) && ((-2.25 / v27) >= 1.5))
                                                    }
                                                    var v28 float32 = 0.0
                                                    _ = v28
                                                } else {
                                                    v17 -= (v17 >> ((3) & 63))
                                                    var v29 uint32 = 536
                                                    _ = len(v1)
                                                    v1[(^(v29))] = true
                                                    _ = v29
                                                }
                                                _ = v27
                                            } else {
                                                v4 = v4
                                                v4 -= (v4 >> ((65535) & 63))
                                                _ = len(v1)
                                                fmt.Println(v4, v17, fi25)
                                            }
                                            v30 := GenS1{F0: v3}
                                            _ = v30
                                        }
                                    default:
                                        var v31 int16 = v4
                                        v4 = v31
                                        fmt.Println(v31)
                                        _ = v31
                                    }
                                    var v32 uint8 = 2
                                    fmt.Println(v3, v17, v17)
                                    fmt.Println(v32, v32)
                                    _ = fi25
                                    _ = v32
                                }
                            }
                            var v33 uint64 = 1
                            _ = v33
                        }
                        v34 := v1[65535]
                        for fi35 := 0; fi35 < 6; fi35++ {
                            switch "a" {
                            case "hello":
                                if (255 <= 18446744073709551615) {
                                    v36 := GenS1{F0: v34}
                                    fmt.Println(fi21)
                                    var v37 uint32 = 42
                                    fmt.Println(v37, v17, v34)
                                    switch (v17 << ((42) & 63)) {
                                    case -2147483648:
                                        v34 = v3
                                    default:
                                        switch (^(v17)) {
                                        case 127:
                                            fmt.Println(v17, v37, v17)
                                        case -848:
                                            v37 = (v37 / ((v37) | 1))
                                        }
                                    }
                                    if (!(v17 > v17)) {
                                        _ = len(v1)
                                    }
                                    v37 -= (v37 ^ 4294967295)
                                    _ = v36
                                    _ = v37
                                } else {
                                    v34 = (true || v34)
                                }
                            case "\n\t":
                                fmt.Println(v3, v4, fi35)
                            }
                            _ = fi35
                        }
                        _ = fi21
                        _ = v34
                    }
                case -1:
                    var v38 int32 = 1
                    _ = v38
                default:
                    if len(v2) > 0 {
                        v2[0] = 155
                    }
                }
            case "foo bar":
                switch v17 {
                case -1:
                    var v39 uint8 = 255
                    _ = v39
                case 1:
                    if false {
                        fmt.Println(v3, v17)
                        if v3 {
                            if (42 > 2147483648) {
                                if v3 {
                                    fmt.Println(v17, v4, v4)
                                    v4 = v4
                                    fmt.Println(v4, v4)
                                }
                            } else {
                                if (v17 <= 9223372036854775807) {
                                    if v3 {
                                        switch 1 {
                                        case 444:
                                            if len(v2) > 0 {
                                                v2[0] = 0
                                            }
                                        case 2147483647:
                                            switch "" {
                                            case "foo bar":
                                                if true {
                                                    if (true && (734 <= 4294967295)) {
                                                        v5[4294967295] = 1
                                                    } else {
                                                        v40 := GenS5{F0: (v17 >> ((2) & 63))}
                                                        _ = v40
                                                    }
                                                    if v3 {
                                                        if v3 {
                                                            for fi41 := 0; fi41 < 20; fi41++ {
                                                                _ = len(v5)
                                                                var v42 float64 = 3.14159
                                                                _ = fi41
                                                                _ = v42
                                                            }
                                                            v3 = (255 != 65535)
                                                        } else {
                                                            v4 -= (v4 | v4)
                                                            v4 -= v4
                                                            v4 -= 32767
                                                        }
                                                        for fi43 := 0; fi43 < 3; fi43++ {
                                                            fmt.Println(v4, fi43, v3)
                                                            _ = fi43
                                                        }
                                                        if v3 {
                                                            var v44 int8 = -128
                                                            _ = v44
                                                        } else {
                                                            var v45 int8 = 42
                                                            for fi46 := 0; fi46 < 14; fi46++ {
                                                                switch (fi46 % ((fi46) | 1)) {
                                                                case -416:
                                                                    if (!v3) {
                                                                        var v47 uint64 = 711
                                                                        if (42 > genFn0()) {
                                                                            if ((!(113 != 2)) && (577 == 2)) {
                                                                                v47 = v47
                                                                                var v48 int16 = (-(v4))
                                                                                _ = v48
                                                                            } else {
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            }
                                                                        } else {
                                                                            for fi49 := 0; fi49 < 3; fi49++ {
                                                                                v17 = (^(fi46))
                                                                                var v50 uint16 = 0
                                                                                var v51 int16 = (v4 ^ v4)
                                                                                _ = fi49
                                                                                _ = v50
                                                                                _ = v51
                                                                            }
                                                                            var v52 int64 = -535
                                                                            var v53 uint32 = 1
                                                                            _ = v52
                                                                            _ = v53
                                                                        }
                                                                        var v54 uint = 4294967295
                                                                        _ = v47
                                                                        _ = v54
                                                                    }
                                                                case 4294967295:
                                                                    v55 := GenS1{F0: (v17 != (fi46 ^ 0))}
                                                                    _ = v55
                                                                default:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 2147483647
                                                                    }
                                                                }
                                                                var v56 uint64 = 1
                                                                _ = fi46
                                                                _ = v56
                                                            }
                                                            v17 -= v17
                                                            _ = v45
                                                        }
                                                    } else {
                                                        v57 := GenS3{F0: (v4 ^ v4), F1: 42, F2: v3}
                                                        _ = v57
                                                    }
                                                    if (("a" + "\n\t") <= (("" + "") + "hello")) {
                                                        var v58 uint8 = 127
                                                        _ = v58
                                                    } else {
                                                        v17 = v17
                                                    }
                                                }
                                            case "a":
                                                v4 += 33
                                            case "hello":
                                                v4 -= v4
                                            }
                                        case 128:
                                            if (!v3) {
                                                v2 = append(v2, 2147483648)
                                                var v59 uint16 = 829
                                                _ = v59
                                            }
                                        default:
                                            var v60 int = v17
                                            _ = v60
                                        }
                                        v61 := v5[0]
                                        fmt.Println(v4, v17, v17)
                                        _ = v61
                                    } else {
                                        var v62 bool = v3
                                        switch int(v4) {
                                        case -444:
                                            v4 -= (v4 - v4)
                                        case -447:
                                            for fi63 := 0; fi63 < 13; fi63++ {
                                                v64 := GenS6{F0: 2}
                                                v5[1] = 4294967295
                                                if v3 {
                                                    v3 = v62
                                                    var v65 string = "a"
                                                    _ = v65
                                                } else {
                                                    if (true || (v17 > (fi63 ^ 547))) {
                                                        switch (^(v17)) {
                                                        case 0:
                                                            var v66 int32 = 42
                                                            _ = v66
                                                        case 32768:
                                                            v67 := GenS6{F0: 2}
                                                            _ = v67
                                                        default:
                                                            for fi68 := 0; fi68 < 14; fi68++ {
                                                                v2 = append(v2, 2147483648)
                                                                v69 := GenS1{F0: (false && ((!(false && ((!v62) && (v62 || v62)))) && (!(65535 > 909))))}
                                                                if v62 {
                                                                    switch (fi63 - (fi68 - (-(fi63)))) {
                                                                    case 42:
                                                                        v70 := GenS5{F0: int(v4)}
                                                                        _ = v70
                                                                    default:
                                                                        v3 = (v62 && false)
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        var v71 int16 = v4
                                                                        switch fi63 {
                                                                        case 9223372036854775807:
                                                                            var v72 uint = 992
                                                                            _ = v72
                                                                        }
                                                                        v73 := GenS7{F0: (100 <= -103)}
                                                                        _ = v71
                                                                        _ = v73
                                                                    } else {
                                                                        if ((v62 || ((v4 <= -32768) || v3)) && (!v62)) {
                                                                            var v74 int64 = -9223372036854775808
                                                                            switch ("" + "") {
                                                                            case "\n\t":
                                                                                _ = len(v2)
                                                                            }
                                                                            switch fi63 {
                                                                            case 32767:
                                                                                var v75 int16 = v4
                                                                                _ = v75
                                                                            case 0:
                                                                                v76 := GenS8{F0: 104, F1: 2, F2: v4}
                                                                                _ = v76
                                                                            }
                                                                            _ = v74
                                                                        } else {
                                                                            v3 = ("foo bar" != ("a" + "a"))
                                                                        }
                                                                        v77 := v5[77]
                                                                        _ = v77
                                                                    }
                                                                    if v62 {
                                                                        v17 = 0
                                                                        v78 := GenS9{F0: 42, F1: 42, F2: 2}
                                                                        v62 = (0 < 679)
                                                                        _ = v78
                                                                    }
                                                                }
                                                                _ = fi68
                                                                _ = v69
                                                            }
                                                        }
                                                        if (4294967295 >= 4294967295) {
                                                            var v79 uint16 = 295
                                                            _ = len(v1)
                                                            if v3 {
                                                                v80 := GenS4{F0: (v4 * v4)}
                                                                _ = v80
                                                            }
                                                            _ = v79
                                                        } else {
                                                            for fi81 := 0; fi81 < 6; fi81++ {
                                                                v82 := GenS10{F0: (^(fi81)), F1: 32767, F2: 127}
                                                                var v83 int32 = 2147483647
                                                                _ = len(v5)
                                                                _ = fi81
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                        }
                                                        v3 = (65535 >= 118)
                                                    }
                                                    for fi84 := 0; fi84 < 9; fi84++ {
                                                        v3 = true
                                                        switch ("\n\t" + "hello") {
                                                        case "a":
                                                            if (-2147483648 >= 1) {
                                                                v62 = v3
                                                            }
                                                        case "":
                                                            if (!((!(861 > 1)) || true)) {
                                                                switch fi84 {
                                                                case -367:
                                                                    var v85 uint8 = 1
                                                                    _ = v85
                                                                case 32767:
                                                                    if (!(292 != 432)) {
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            v86 := GenS8{F0: 255, F1: 15, F2: 0}
                                                                            _ = v86
                                                                        case "":
                                                                            if (!false) {
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = 65535
                                                                                }
                                                                                v1[2147483648] = ((float32((-2.25 / 1e10)) <= (0.0 + 0.1)) || ((fi84 | (-(fi84))) >= (v17 >> ((2) & 63))))
                                                                                v3 = true
                                                                            }
                                                                        default:
                                                                            v87 := GenS7{F0: (((true || v62) && v3) && v3)}
                                                                            _ = v87
                                                                        }
                                                                        var v88 float64 = 1e10
                                                                        v4 -= 42
                                                                        _ = v88
                                                                    }
                                                                default:
                                                                    v62 = v62
                                                                    if true {
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            var v89 float64 = 3.14159
                                                                            _ = v89
                                                                        case "foo bar":
                                                                            if true {
                                                                                if v62 {
                                                                                    _ = len(v5)
                                                                                    var v90 uint64 = 42
                                                                                    var v91 float32 = 1.5
                                                                                    v2 = append(v2, 95)
                                                                                    _ = v90
                                                                                    _ = v91
                                                                                } else {
                                                                                    var v92 string = "\n\t"
                                                                                    _ = v92
                                                                                }
                                                                                switch -9223372036854775808 {
                                                                                case -9223372036854775808:
                                                                                    v4 -= v4
                                                                                }
                                                                                v3 = ((true || (42 < 127)) && v62)
                                                                            } else {
                                                                                if true {
                                                                                    var v93 int = (v17 * fi63)
                                                                                    if v3 {
                                                                                        if v3 {
                                                                                            v17 = (v17 / ((v93) | 1))
                                                                                            fmt.Println(v62)
                                                                                            v94 := GenS11{F0: 1, F1: v4}
                                                                                            _ = v94
                                                                                        } else {
                                                                                            var v95 int = v93
                                                                                            _ = v95
                                                                                        }
                                                                                    } else {
                                                                                        delete(v5, 2)
                                                                                    }
                                                                                    v96 := GenS2{F0: "foo bar"}
                                                                                    if false {
                                                                                        if true {
                                                                                            fmt.Println(v17, fi84, fi84)
                                                                                        } else {
                                                                                            v62 = true
                                                                                            if true {
                                                                                                switch (fi63 - (^(fi84))) {
                                                                                                case -9223372036854775808:
                                                                                                    if false {
                                                                                                        v97 := GenS12{F0: 4294967295, F1: 42, F2: 545}
                                                                                                        _ = v97
                                                                                                    }
                                                                                                    var v98 uint16 = 253
                                                                                                    if ((false && v3) || (v3 && v3)) {
                                                                                                        v17 = v17
                                                                                                    }
                                                                                                    switch "foo bar" {
                                                                                                    case "":
                                                                                                        var v99 uint64 = 18446744073709551615
                                                                                                        _ = v99
                                                                                                    case "foo bar":
                                                                                                        var v100 float32 = 1.5
                                                                                                        _ = v100
                                                                                                    case "a":
                                                                                                        fmt.Println(v17, v62)
                                                                                                    default:
                                                                                                        switch fi84 {
                                                                                                        case 524:
                                                                                                            if (!(!(!((!true) || (0 <= 0))))) {
                                                                                                                v98 -= (v98 + (v98 << ((2) & 63)))
                                                                                                            }
                                                                                                        case -9223372036854775808:
                                                                                                            if v3 {
                                                                                                                if len(v2) > 0 {
                                                                                                                    v2[0] = 182
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v101 uint64 = 0
                                                                                                                var v102 uint32 = 336
                                                                                                                _ = v101
                                                                                                                _ = v102
                                                                                                            }
                                                                                                        case 4294967295:
                                                                                                            var v103 float32 = 3.14159
                                                                                                            _ = v103
                                                                                                        }
                                                                                                    }
                                                                                                    v104 := v1[65535]
                                                                                                    _ = v104
                                                                                                    _ = v98
                                                                                                default:
                                                                                                    v3 = v3
                                                                                                    if len(v2) >= 2 {
                                                                                                        v2 = v2[1:len(v2)]
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                v62 = (("" <= "a") || v62)
                                                                                                v3 = (!false)
                                                                                                if true {
                                                                                                    v105 := GenS13{F0: 42, F1: 65535}
                                                                                                    var v106 uint8 = 255
                                                                                                    _ = v105
                                                                                                    _ = v106
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v93 = 392
                                                                                    } else {
                                                                                        v4 = 570
                                                                                        v93 = (fi63 ^ (fi63 - 42))
                                                                                    }
                                                                                    var v107 int64 = 9223372036854775807
                                                                                    fmt.Println(v62, fi63)
                                                                                    v4 = (v4 - v4)
                                                                                    _ = v107
                                                                                    _ = v93
                                                                                    _ = v96
                                                                                } else {
                                                                                    v108 := GenS14{F0: 138}
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = 842
                                                                                    }
                                                                                    v109 := GenS15{F0: (!v62), F1: 419}
                                                                                    _ = v108
                                                                                    _ = v109
                                                                                }
                                                                                var v110 int32 = -2147483648
                                                                                _ = v110
                                                                            }
                                                                        case "":
                                                                            if (169 < 101) {
                                                                                var v111 int64 = -1
                                                                                switch fi63 {
                                                                                case 2147483647:
                                                                                    var v112 uint16 = 0
                                                                                    switch fi63 {
                                                                                    case 32768:
                                                                                        v1[0] = v62
                                                                                    default:
                                                                                        var v113 float64 = 3.14159
                                                                                        _ = v113
                                                                                    }
                                                                                    _ = v112
                                                                                case -2147483648:
                                                                                    _ = len(v1)
                                                                                case 379:
                                                                                    var v114 float64 = 1e10
                                                                                    _ = v114
                                                                                }
                                                                                _ = v111
                                                                            } else {
                                                                                if (103 <= 18446744073709551615) {
                                                                                    switch fi84 {
                                                                                    case 42:
                                                                                        if v62 {
                                                                                            _ = len(v2)
                                                                                            var v115 bool = v3
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            _ = v115
                                                                                        }
                                                                                    default:
                                                                                        v116 := GenS12{F0: 864, F1: 40, F2: 542}
                                                                                        _ = v116
                                                                                    }
                                                                                }
                                                                            }
                                                                        default:
                                                                            switch fi84 {
                                                                            case -1:
                                                                                var v117 int = fi84
                                                                                _ = v117
                                                                            default:
                                                                                v62 = (255 == 260)
                                                                            }
                                                                        }
                                                                        v118 := GenS16{F0: 1, F1: 564, F2: ""}
                                                                        v119 := GenS17{F0: 630, F1: (959 == -1)}
                                                                        if v62 {
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                v62 = (true && v3)
                                                                            }
                                                                        } else {
                                                                            var v120 int16 = 42
                                                                            if (!false) {
                                                                                v17 += fi63
                                                                            }
                                                                            if false {
                                                                                v4 = 1
                                                                            }
                                                                            _ = v120
                                                                        }
                                                                        _ = v118
                                                                        _ = v119
                                                                    } else {
                                                                        v121 := GenS12{F0: 18446744073709551615, F1: 668, F2: 954}
                                                                        var v122 uint16 = 422
                                                                        if (1e10 <= 1e10) {
                                                                            v123 := GenS2{F0: "\n\t"}
                                                                            _ = v123
                                                                        }
                                                                        if len(v2) > 0 {
                                                                            v2[0] = 951
                                                                        }
                                                                        _ = v121
                                                                        _ = v122
                                                                    }
                                                                    v3 = (v3 || (4294967295 >= 42))
                                                                }
                                                                fmt.Println(v3, v62)
                                                                if (598 < 732) {
                                                                    var v124 int16 = -1
                                                                    _ = v124
                                                                }
                                                            }
                                                        case "foo bar":
                                                            fmt.Println(fi84, fi63, v62)
                                                        }
                                                        _ = fi84
                                                    }
                                                    for fi125 := 0; fi125 < 12; fi125++ {
                                                        var v126 bool = v3
                                                        v127 := GenS18{F0: 32767, F1: ""}
                                                        _ = fi125
                                                        _ = v126
                                                        _ = v127
                                                    }
                                                }
                                                _ = fi63
                                                _ = v64
                                            }
                                        }
                                        v17 += (v17 | (v17 % (((v17 / ((v17) | 1))) | 1)))
                                        _ = v62
                                    }
                                    v5[2] = 2147483648
                                    v4 += 195
                                } else {
                                    if (320 < -2147483648) {
                                        v3 = (3.14159 <= float64(3.14159))
                                        for fi128 := 0; fi128 < 16; fi128++ {
                                            fmt.Println(v4, v4)
                                            if (169 <= 1) {
                                                _ = len(v2)
                                                if ("foo bar" == "hello") {
                                                    v129 := GenS7{F0: v3}
                                                    _ = v129
                                                }
                                                if ((!((!v3) && v3)) || v3) {
                                                    var v130 uint = 957
                                                    v5[2] = 2
                                                    _ = v130
                                                } else {
                                                    v4 = 32767
                                                }
                                            }
                                            v131 := GenS18{F0: 344, F1: "\n\t"}
                                            _ = fi128
                                            _ = v131
                                        }
                                    } else {
                                        v132 := GenS9{F0: 2147483647, F1: 17, F2: 0}
                                        switch (v17 ^ (v17 + 127)) {
                                        case -2147483648:
                                            _ = len(v5)
                                        }
                                        _ = v132
                                    }
                                    switch v17 {
                                    case -2147483648:
                                        var v133 int16 = v4
                                        _ = v133
                                    }
                                }
                                v3 = v3
                                var v134 float64 = (1.5 * 0.0)
                                _ = v134
                            }
                            if v3 {
                                v17 -= (v17 + (v17 + (v17 << ((453) & 63))))
                            }
                            v135 := GenS4{F0: v4}
                            _ = v135
                        } else {
                            if (1 == (v4 & v4)) {
                                switch v17 {
                                case 42:
                                    fmt.Println(v3, v17, v3)
                                case -805:
                                    switch (v17 + 0) {
                                    case -783:
                                        if (!(v3 || v3)) {
                                            v136 := GenS19{F0: 32768}
                                            _ = v136
                                        }
                                    }
                                }
                                fmt.Println(v17, v17, v17)
                            } else {
                                var v137 int8 = -128
                                if (65535 != 1) {
                                    switch "\n\t" {
                                    case "a":
                                        var v138 int32 = 0
                                        _ = v138
                                    case "hello":
                                        if false {
                                            fmt.Println(v137, v3)
                                            switch ("foo bar" + "foo bar") {
                                            case "":
                                                var v139 int16 = int16(v17)
                                                _ = v139
                                            case "\n\t":
                                                if (1 != 4294967295) {
                                                    for fi140 := 0; fi140 < 4; fi140++ {
                                                        for fi141 := 0; fi141 < 14; fi141++ {
                                                            switch "a" {
                                                            case "a":
                                                                v142 := GenS20{F0: (0 > 0), F1: true}
                                                                _ = v142
                                                            case "foo bar":
                                                                v2 = append(v2, 0)
                                                                v143 := GenS8{F0: 226, F1: 0, F2: (v4 % (((v4 & (-(v4)))) | 1))}
                                                                _ = v143
                                                            case "hello":
                                                                if (42 < v4) {
                                                                    _ = len(v5)
                                                                    v144 := GenS5{F0: -2147483648}
                                                                    var v145 uint16 = 2
                                                                    if true {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    } else {
                                                                        var v146 int64 = -115
                                                                        v17 += (fi141 * (fi141 % ((v17) | 1)))
                                                                        var v147 int = (fi140 << ((255) & 63))
                                                                        var v148 uint32 = 2
                                                                        _ = v146
                                                                        _ = v147
                                                                        _ = v148
                                                                    }
                                                                    _ = v144
                                                                    _ = v145
                                                                } else {
                                                                    v4 = -18
                                                                    v149 := GenS9{F0: 2147483648, F1: v137, F2: 1}
                                                                    if v3 {
                                                                        switch (((("hello" + "foo bar") + "\n\t") + "a") + "hello") {
                                                                        case "a":
                                                                            if false {
                                                                                switch 353 {
                                                                                case -1:
                                                                                    v3 = (v3 && v3)
                                                                                case -61:
                                                                                    fmt.Println(fi141)
                                                                                case 42:
                                                                                    var v150 float32 = (0.1 + 3.14159)
                                                                                    _ = v150
                                                                                default:
                                                                                    var v151 bool = (int8(v4) <= -1)
                                                                                    if (-557 <= 955) {
                                                                                        v17 += (fi141 % (((^(fi140))) | 1))
                                                                                    } else {
                                                                                        if (42 == 42) {
                                                                                            v152 := GenS1{F0: false}
                                                                                            v153 := GenS2{F0: "a"}
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            _ = v152
                                                                                            _ = v153
                                                                                        } else {
                                                                                            v3 = (!(!(v3 && true)))
                                                                                        }
                                                                                    }
                                                                                    _ = v151
                                                                                }
                                                                                v154 := GenS7{F0: false}
                                                                                _ = v154
                                                                            } else {
                                                                                var v155 uint = 42
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                _ = v155
                                                                            }
                                                                        case "\n\t":
                                                                            var v156 int = (fi140 | fi140)
                                                                            _ = v156
                                                                        case "":
                                                                            v157 := GenS21{F0: "a", F1: "", F2: 18446744073709551615}
                                                                            _ = v157
                                                                        }
                                                                        var v158 uint64 = 1
                                                                        _ = v158
                                                                    }
                                                                    _ = v149
                                                                }
                                                            default:
                                                                v4 -= (v4 - -32768)
                                                            }
                                                            _ = fi141
                                                        }
                                                        for fi159 := 0; fi159 < 16; fi159++ {
                                                            if (2147483648 <= 42) {
                                                                v160 := v1[2147483648]
                                                                switch "hello" {
                                                                case "hello":
                                                                    var v161 uint64 = 1
                                                                    _ = v161
                                                                case "foo bar":
                                                                    v162 := GenS22{F0: 191, F1: -128}
                                                                    _ = v162
                                                                case "a":
                                                                    v163 := GenS1{F0: (fi159 < fi140)}
                                                                    _ = v163
                                                                }
                                                                fmt.Println(fi140)
                                                                _ = v160
                                                            } else {
                                                                if false {
                                                                    v164 := GenS21{F0: "hello", F1: ("a" + ""), F2: 1}
                                                                    var v165 int8 = -46
                                                                    var v166 bool = v3
                                                                    fmt.Println(v17, v3)
                                                                    var v167 string = "a"
                                                                    _ = v164
                                                                    _ = v165
                                                                    _ = v166
                                                                    _ = v167
                                                                }
                                                                var v168 int32 = 42
                                                                if (42 < 544) {
                                                                    v169 := GenS7{F0: v3}
                                                                    var v170 float32 = 1.5
                                                                    _ = v169
                                                                    _ = v170
                                                                } else {
                                                                    var v171 string = "foo bar"
                                                                    v172 := GenS8{F0: 145, F1: 127, F2: v4}
                                                                    _ = v171
                                                                    _ = v172
                                                                }
                                                                fmt.Println(v3, v168)
                                                                _ = v168
                                                            }
                                                            var v173 uint8 = 2
                                                            _ = len(v1)
                                                            _ = fi159
                                                            _ = v173
                                                        }
                                                        _ = fi140
                                                    }
                                                }
                                            case "foo bar":
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case "a":
                                                    v3 = (!((v137 - v137) <= v137))
                                                default:
                                                    var v174 int8 = 0
                                                    _ = v174
                                                }
                                            default:
                                                v17 = (v17 ^ (v17 / (((v17 >> ((65535) & 63))) | 1)))
                                            }
                                        } else {
                                            v4 -= (v4 | v4)
                                            for fi175 := 0; fi175 < 16; fi175++ {
                                                var v176 float32 = ((0.0 * (0.1 / 1.5)) * 0.1)
                                                switch ("" + "foo bar") {
                                                case "a":
                                                    if (181 > 0) {
                                                        var v177 uint = 18446744073709551615
                                                        if v3 {
                                                            if (v176 >= -2.25) {
                                                                fmt.Println(fi175, v17)
                                                            }
                                                        }
                                                        var v178 bool = v3
                                                        _ = v177
                                                        _ = v178
                                                    }
                                                case "":
                                                    var v179 int64 = -909
                                                    _ = v179
                                                }
                                                if ((-1 > 0) && (v3 || (v3 || v3))) {
                                                    if true {
                                                        v3 = (1 <= v4)
                                                        _ = len(v1)
                                                        var v180 int16 = (v4 * (^(v4)))
                                                        _ = v180
                                                    } else {
                                                        var v181 uint8 = 255
                                                        _ = v181
                                                    }
                                                }
                                                _ = fi175
                                                _ = v176
                                            }
                                        }
                                    case "foo bar":
                                        var v182 float64 = (float64(0.0) / float64((-2.25 - (0.0 * ((-2.25 / 1.5) - -2.25)))))
                                        _ = v182
                                    }
                                    for fi183 := 0; fi183 < 4; fi183++ {
                                        v3 = (false && ("a" > ""))
                                        for fi184 := 0; fi184 < 2; fi184++ {
                                            v137 += v137
                                            var v185 float32 = (-2.25 - (1e10 + 1.5))
                                            _ = fi184
                                            _ = v185
                                        }
                                        _ = fi183
                                    }
                                    if v3 {
                                        if v3 {
                                            var v186 int32 = -1
                                            var v187 float64 = 0.0
                                            var v188 int32 = -547
                                            _ = v186
                                            _ = v187
                                            _ = v188
                                        }
                                        var v189 string = "foo bar"
                                        _ = v189
                                    }
                                }
                                v190 := GenS23{F0: 189, F1: 0, F2: -816}
                                _ = v137
                                _ = v190
                            }
                            if v3 {
                                var v191 bool = v3
                                _ = v191
                            }
                            var v192 uint = 2
                            _ = v192
                        }
                        var v193 float64 = 1e10
                        _ = v193
                    }
                default:
                    v194 := GenS17{F0: 1, F1: v3}
                    _ = v194
                }
            default:
                v2 = append(v2, 143)
            }
            v4 = (v4 ^ v4)
        }
        _ = v17
    }
    var v195 int = 0
    switch "" {
    case "\n\t":
        v196 := GenS3{F0: (-(v4)), F1: 79, F2: v3}
        _ = v196
    case "hello":
        v4 += v4
    case "foo bar":
        var v197 int64 = -627
        _ = v197
    default:
        var v198 float64 = (1.5 + 1.5)
        _ = v198
    }
    v5[2147483648] = 530
    if true {
        v3 = false
    }
    var v199 float32 = float32(3.14159)
    v199 += ((float32(0.1) + float32(-2.25)) - (-2.25 / ((((0.1 / float32(0.1)) - (v199 / (1e10 + 0.1))) * ((v199 / (v199 + v199)) * 1e10)) + v199)))
    for fi200 := 0; fi200 < 17; fi200++ {
        v195 = v195
        v195 = (v195 - (fi200 & (fi200 ^ 1)))
        v201 := GenS21{F0: (("" + "") + (("a" + ((("" + "hello") + "\n\t") + ("a" + "\n\t"))) + "foo bar")), F1: "a", F2: 191}
        _ = fi200
        _ = v201
    }
    if false {
        _ = len(v2)
    } else {
        fmt.Println(v195, v4)
    }
    if false {
        var v202 int16 = (v4 + v4)
        v4 += v4
        var v203 float64 = ((3.14159 - 3.14159) + 3.14159)
        _ = v202
        _ = v203
    } else {
        switch ((("a" + (("\n\t" + "a") + "foo bar")) + "\n\t") + "a") {
        case "foo bar":
            for fi204 := 0; fi204 < 16; fi204++ {
                v199 = 0.1
                var v205 uint32 = 42
                for fi206 := 0; fi206 < 10; fi206++ {
                    var v207 int8 = -128
                    if ((!(v3 && v3)) && (32767 == 32767)) {
                        if ((-9223372036854775808 != -9223372036854775808) && v3) {
                            v195 = -1
                            v207 -= 122
                            v4 -= v4
                        }
                        fmt.Println(fi204, v3)
                        v207 -= (v207 - v207)
                    }
                    _ = fi206
                    _ = v207
                }
                _ = fi204
                _ = v205
            }
        case "a":
            v208 := GenS24{F0: 32767, F1: -2147483648}
            _ = v208
        }
        var v209 int = 32767
        _ = v209
    }
    var v210 int64 = 9223372036854775807
    v4 += -867
    if (42 != 18446744073709551615) {
        v210 -= v210
    }
    v195 = v195
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v195:", v195)
    fmt.Println("v199:", v199)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v210:", v210)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v195
    _ = v199
    _ = v2
    _ = v210
    _ = v3
    _ = v4
    _ = v5
}
