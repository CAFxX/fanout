package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint16; F1 string; F2 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int8; F1 uint }
type GenS10 struct { F0 uint }
func genFn0(p0 uint, p1 uint, p2 uint32) uint64 {
    if (false || false) {
        p2 -= 11
        var v7 int8 = 0
        _ = v7
    }
    var v8 int = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}
func genFn1() int32 {
    v9 := GenS1{F0: 18446744073709551615, F1: 42, F2: 127}
    var v10 uint16 = 1
    if true {
        v10 = v10
        v11 := []int{2147483647, 847}
        _ = v11
    } else {
        switch "" {
        case "foo bar":
            if (float32(1.5) < (0.0 / 1.5)) {
                v12 := []int{9223372036854775807, -9223372036854775808, 32767, -686}
                var v13 float64 = 1.5
                v13 = v13
                _ = v12
                _ = v13
            } else {
                v10 = v10
            }
        case "":
            var v14 uint32 = 2
            _ = v14
        case "\n\t":
            v15 := []int{50, 1, 2147483647}
            _ = v15
        }
        v16 := []int{}
        var v17 string = "hello"
        _ = v16
        _ = v17
    }
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = ("\n\t" + "a")
    var v3 int8 = -30
    v4 := make(map[string]bool)
    v5 := []int{795, 42, 2147483647, 127}
    var v6 bool = (141 >= 1)
    for fi18 := 0; fi18 < 8; fi18++ {
        if true {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi19 := 0; fi19 < 19; fi19++ {
                var v20 int = 1
                _ = fi19
                _ = v20
            }
            v3 -= (v3 & (v3 & v3))
        }
        _ = fi18
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        if true {
            switch (^(fi21)) {
            case -180:
                var v22 int = fi21
                _ = v22
            default:
                if (!v6) {
                    v6 = v6
                } else {
                    var v23 float64 = 3.14159
                    v3 = v3
                    delete(v4, (v2 + ((v2 + "foo bar") + v2)))
                    _ = v23
                }
            }
            var v24 uint = 0
            v24 -= v24
            _ = v24
        }
        for fi25 := 0; fi25 < 3; fi25++ {
            v2 = v2
            fmt.Println(v3, v6, v6)
            v26 := GenS1{F0: 18446744073709551615, F1: (fi25 - (fi25 % (((fi21 ^ fi21)) | 1))), F2: 119}
            _ = fi25
            _ = v26
        }
        var v27 int32 = 614
        _ = fi21
        _ = v27
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 int = fi28
        var v30 uint32 = 0
        _ = fi28
        _ = v29
        _ = v30
    }
    v2 = "a"
    var v31 int8 = v3
    for fi32 := 0; fi32 < 15; fi32++ {
        for fi33 := 0; fi33 < 6; fi33++ {
            var v34 uint32 = 2147483648
            if (v6 && v6) {
                if (v6 && (!v6)) {
                    fmt.Println(v34, v2)
                    switch "foo bar" {
                    case "a":
                        v1[v34] = v34
                        var v35 uint8 = 255
                        var v36 int = (fi33 | -876)
                        switch (v2 + v2) {
                        case "\n\t":
                            v34 = v34
                        case "a":
                            v1[2147483647] = v34
                        case "hello":
                            if v6 {
                                var v37 uint16 = uint16(v35)
                                _ = v37
                            } else {
                                if (v35 == v35) {
                                    var v38 uint16 = 0
                                    var v39 int32 = 0
                                    if v6 {
                                        var v40 int32 = int32(v38)
                                        _ = v40
                                    } else {
                                        if true {
                                            v3 = v31
                                        } else {
                                            var v41 bool = v6
                                            var v42 float64 = 0.1
                                            switch v2 {
                                            case "":
                                                v43 := GenS1{F0: 4294967295, F1: (fi33 << ((4294967295) & 63)), F2: v35}
                                                switch (v2 + "a") {
                                                case "":
                                                    var v44 int64 = 1
                                                    _ = v44
                                                }
                                                v6 = ((v3 & int8(v36)) != (v3 & v3))
                                                _ = v43
                                            case "\n\t":
                                                var v45 bool = false
                                                _ = v45
                                            }
                                            _ = v41
                                            _ = v42
                                        }
                                        var v46 int16 = -32768
                                        _ = v46
                                    }
                                    _ = v38
                                    _ = v39
                                }
                                var v47 uint16 = 32767
                                _ = v47
                            }
                            var v48 int64 = 432
                            v49 := GenS2{F0: 42, F1: v31}
                            _ = v48
                            _ = v49
                        }
                        _ = v35
                        _ = v36
                    case "\n\t":
                        fmt.Println(v31, v34)
                    }
                    v5 = append(v5, 127)
                } else {
                    switch (v2 + "foo bar") {
                    case "foo bar":
                        var v50 int16 = 262
                        _ = v50
                    }
                }
                var v51 uint32 = v34
                v31 += -128
                _ = v51
            }
            switch v2 {
            case "a":
                fmt.Println(v6)
                switch fi32 {
                case 32768:
                    var v52 uint = 255
                    _ = v52
                default:
                    var v53 int32 = -1
                    _ = v53
                }
            default:
                switch fi32 {
                case 4294967295:
                    fmt.Println(v6, fi32)
                case -1:
                    if (!v6) {
                        var v54 uint16 = 1
                        if (-2.25 < (float32(3.14159) + (3.14159 / ((float32(0.0) - (-2.25 + 1e10)) - 3.14159)))) {
                            if (764 < uint64(v34)) {
                                var v55 uint16 = 2
                                switch ("" + (v2 + (v2 + ("\n\t" + "\n\t")))) {
                                case "a":
                                    var v56 uint64 = 2
                                    _ = v56
                                case "hello":
                                    if v6 {
                                        if ((v3 * v31) <= v31) {
                                            var v57 int8 = v3
                                            var v58 uint = 0
                                            switch ((v2 + ((("\n\t" + v2) + v2) + "hello")) + (("a" + "\n\t") + v2)) {
                                            case "a":
                                                v58 = (v58 << (((v58 + (v58 * 2))) & 63))
                                            }
                                            _ = v57
                                            _ = v58
                                        } else {
                                            if false {
                                                fmt.Println(v2)
                                                v3 = (v3 & -118)
                                            }
                                            switch (fi33 << ((331) & 63)) {
                                            case -2147483648:
                                                switch v2 {
                                                case "a":
                                                    var v59 int8 = 0
                                                    _ = v59
                                                }
                                            case -817:
                                                var v60 uint = 765
                                                _ = v60
                                            case 128:
                                                fmt.Println(fi33)
                                            }
                                            v61 := GenS1{F0: 4294967295, F1: fi32, F2: 127}
                                            _ = v61
                                        }
                                        var v62 float64 = 0.0
                                        v3 = (v3 % ((-4) | 1))
                                        _ = v62
                                    } else {
                                        var v63 float32 = 1e10
                                        fmt.Println(v55, v55, v6)
                                        _ = v63
                                    }
                                case "":
                                    v5 = append(v5, fi33)
                                default:
                                    v64 := GenS2{F0: 42, F1: 1}
                                    _ = v64
                                }
                                v31 = v3
                                _ = v55
                            } else {
                                _ = len(v5)
                                fmt.Println(fi32, v31, v6)
                            }
                            v54 = v54
                            if (0.1 <= (1.5 + 1.5)) {
                                fmt.Println(v3, fi33)
                                v65 := GenS2{F0: -32768, F1: 100}
                                _ = v65
                            } else {
                                var v66 int64 = -902
                                fmt.Println(v31)
                                if v6 {
                                    var v67 bool = v6
                                    if (v54 == v54) {
                                        var v68 uint64 = 848
                                        if len(v5) > 0 {
                                            v5[0] = 762
                                        }
                                        v69 := v1[(v34 + v34)]
                                        _ = v68
                                        _ = v69
                                    }
                                    _ = v67
                                } else {
                                    var v70 int = (fi32 ^ fi33)
                                    switch (fi32 & fi32) {
                                    case 4294967295:
                                        v71 := GenS2{F0: 42, F1: (v3 >> ((892) & 63))}
                                        _ = v71
                                    }
                                    var v72 int16 = -32768
                                    _ = v70
                                    _ = v72
                                }
                                _ = v66
                            }
                        }
                        _ = v54
                    } else {
                        v31 += v31
                        if false {
                            fmt.Println(v31)
                            switch v2 {
                            case "hello":
                                if v6 {
                                    var v73 bool = (!v6)
                                    v5 = append(v5, fi32)
                                    v34 -= v34
                                    _ = v73
                                } else {
                                    _ = len(v5)
                                    fmt.Println(v31, v2, fi33)
                                }
                            case "a":
                                if (!v6) {
                                    fmt.Println(v31, v6)
                                    if (-1 > 9223372036854775807) {
                                        fmt.Println(v31)
                                    } else {
                                        var v74 int32 = -2147483648
                                        _ = v74
                                    }
                                }
                            case "\n\t":
                                v31 = v3
                            }
                            if (403 != v34) {
                                switch (-(fi33)) {
                                case 32:
                                    var v75 uint32 = v34
                                    _ = v75
                                case -780:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                case 42:
                                    var v76 int16 = -1
                                    _ = v76
                                default:
                                    v77 := GenS2{F0: 32767, F1: -1}
                                    _ = v77
                                }
                                switch ((v2 + (v2 + ((v2 + v2) + "\n\t"))) + v2) {
                                case "":
                                    v3 = (v3 * (-(v3)))
                                case "a":
                                    switch "" {
                                    case "foo bar":
                                        delete(v4, (v2 + v2))
                                    case "a":
                                        if (v6 && false) {
                                            v31 -= (v31 & v3)
                                            var v78 float64 = ((((-2.25 * (-2.25 + 0.0)) + (1e10 * (1e10 + 3.14159))) + -2.25) * (0.0 * 3.14159))
                                            if false {
                                                v6 = false
                                                v34 -= (v34 + (v34 ^ (v34 / ((v34) | 1))))
                                                _ = len(v1)
                                            }
                                            _ = v78
                                        }
                                    default:
                                        v3 -= 127
                                    }
                                case "foo bar":
                                    v31 -= 0
                                default:
                                    var v79 int64 = 9223372036854775807
                                    _ = v79
                                }
                                v80 := GenS1{F0: 18446744073709551615, F1: (fi33 * -1), F2: 54}
                                _ = v80
                            }
                        } else {
                            if true {
                                v81 := GenS3{F0: 1, F1: v2, F2: genFn1()}
                                switch fi32 {
                                case -175:
                                    fmt.Println(fi32, v6, v2)
                                }
                                _ = v81
                            }
                            switch "hello" {
                            case "\n\t":
                                _ = len(v5)
                            default:
                                fmt.Println(fi32, v2, fi32)
                            }
                        }
                    }
                }
            }
            _ = fi33
            _ = v34
        }
        if false {
            v82 := GenS1{F0: 799, F1: (^(fi32)), F2: 2}
            v83 := GenS4{F0: 42}
            var v84 uint64 = 0
            _ = v82
            _ = v83
            _ = v84
        }
        switch "" {
        case "hello":
            for fi85 := 0; fi85 < 2; fi85++ {
                if v6 {
                    v2 = ""
                    v6 = true
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v86 uint8 = 255
                    _ = v86
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                v87 := GenS5{F0: ((-(v3)) != (v31 << ((0) & 63)))}
                _ = fi85
                _ = v87
            }
        case "foo bar":
            if false {
                v6 = true
                var v88 float32 = 0.1
                for fi89 := 0; fi89 < 15; fi89++ {
                    var v90 uint64 = 4294967295
                    if (2 >= 128) {
                        if len(v5) > 0 {
                            v5[0] = (-(fi32))
                        }
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi89
                    _ = v90
                }
                _ = v88
            } else {
                var v91 uint32 = 871
                v31 = v3
                _ = v91
            }
        }
        _ = fi32
    }
    for fi92 := 0; fi92 < 9; fi92++ {
        var v93 bool = true
        _ = fi92
        _ = v93
    }
    fmt.Println(v31)
    v94 := GenS6{F0: 4294967295, F1: 0}
    v95 := GenS1{F0: 1, F1: -511, F2: 1}
    for fi96 := 0; fi96 < 6; fi96++ {
        if (v6 || v6) {
            var v97 float32 = (0.1 * 3.14159)
            _ = v97
        }
        _ = len(v4)
        _ = fi96
    }
    if (7 == v3) {
        for fi98 := 0; fi98 < 8; fi98++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi98
        }
        v99 := GenS7{F0: (v3 / ((v31) | 1)), F1: 950}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v99
    } else {
        var v100 string = ((v2 + v2) + (v2 + (v2 + v2)))
        v100 = v100
        _ = v100
    }
    for fi101 := 0; fi101 < 19; fi101++ {
        for fi102 := 0; fi102 < 7; fi102++ {
            v103 := GenS5{F0: (2147483647 > -684)}
            v2 = "a"
            if (!(!v6)) {
                switch (fi102 ^ int(v31)) {
                case 615:
                    if v6 {
                        v6 = ((0.0 >= 1e10) || v6)
                    }
                case 32767:
                    v2 = v2
                case -1:
                    switch (fi101 & (fi101 % (((fi102 - 9223372036854775807)) | 1))) {
                    case -2147483648:
                        if len(v5) > 0 {
                            v5[0] = int(v31)
                        }
                    case 4294967295:
                        v2 = ("foo bar" + v2)
                        if (1 < 593) {
                            var v104 uint64 = 42
                            _ = v104
                        } else {
                            v6 = (v6 && (!false))
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if (!v6) {
                                var v105 int = -9223372036854775808
                                _ = v105
                            } else {
                                if v6 {
                                    var v106 string = v2
                                    v107 := GenS5{F0: (((-1 <= 1) || (v6 && v6)) || v6)}
                                    fmt.Println(v2, v3)
                                    _ = v106
                                    _ = v107
                                } else {
                                    var v108 uint64 = 795
                                    v3 += (v3 << ((255) & 63))
                                    v109 := GenS8{F0: (v108 * v108)}
                                    _ = v108
                                    _ = v109
                                }
                            }
                        }
                    default:
                        v3 -= v31
                    }
                }
                switch fi102 {
                case -2147483648:
                    v6 = false
                case 0:
                    var v110 string = v2
                    _ = v110
                default:
                    var v111 string = ("a" + v2)
                    _ = v111
                }
                v6 = (!(v6 || v6))
            }
            _ = fi102
            _ = v103
        }
        for fi112 := 0; fi112 < 13; fi112++ {
            v3 -= 0
            _ = fi112
        }
        _ = fi101
    }
    for fi113 := 0; fi113 < 18; fi113++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi113
    }
    var v114 uint8 = 166
    v2 = v2
    for fi115 := 0; fi115 < 13; fi115++ {
        if v6 {
            switch fi115 {
            case 4294967295:
                for fi116 := 0; fi116 < 4; fi116++ {
                    var v117 int8 = (-(v3))
                    _ = fi116
                    _ = v117
                }
            case 1:
                fmt.Println(v31, v3)
            }
            v118 := GenS9{F0: (v3 - -128), F1: 4294967295}
            _ = v118
        } else {
            v2 = (("\n\t" + v2) + (v2 + v2))
            for fi119 := 0; fi119 < 19; fi119++ {
                v31 -= int8(v114)
                v120 := GenS10{F0: 65535}
                var v121 string = v2
                _ = fi119
                _ = v120
                _ = v121
            }
        }
        var v122 uint64 = 0
        _ = fi115
        _ = v122
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v114:", v114)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v114
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v94
    _ = v95
}
