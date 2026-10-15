package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 int; F2 uint }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint; F2 int16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 uint8; F2 int }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int64; F1 uint32 }
type GenS11 struct { F0 uint8; F1 uint16 }
type GenS12 struct { F0 uint16; F1 uint64 }
type GenS13 struct { F0 uint; F1 string }
func genFn0(p0 int, p1 int) uint32 {
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: 4294967295}
        p1 += 32767
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint = 65535
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    p0 -= (p1 >> ((42) & 63))
    v11 := []int{(p0 % (((-(p1))) | 1)), (p1 - p1), 1}
    _ = p0
    _ = p1
    _ = v11
    return 2147483648
}
func genFn1(p0 uint8, p1 uint8) uint8 {
    for fi12 := 0; fi12 < 3; fi12++ {
        if true {
            switch ("" + "\n\t") {
            case "a":
                v13 := []int{fi12, fi12, 1, fi12}
                _ = v13
            case "\n\t":
                v14 := GenS1{F0: 0}
                _ = v14
            case "hello":
                fmt.Println(fi12, p0, p0)
            }
            fmt.Println(fi12, p0, p1)
            for fi15 := 0; fi15 < 7; fi15++ {
                if false {
                    fmt.Println(fi15)
                    var v16 int64 = 42
                    _ = v16
                }
                var v17 string = ((("hello" + "a") + "foo bar") + "a")
                var v18 float64 = (-2.25 * 1e10)
                _ = fi15
                _ = v17
                _ = v18
            }
        } else {
            var v19 string = ("foo bar" + "")
            p1 += (p0 & (^(p0)))
            v20 := GenS2{F0: p1, F1: 42}
            _ = v19
            _ = v20
        }
        _ = fi12
    }
    if (((!false) && (!false)) || (0 > 31)) {
        var v21 string = (("" + ("a" + "a")) + ("" + ("hello" + "")))
        if false {
            fmt.Println(p0)
        } else {
            var v22 float32 = float32(-2.25)
            switch "\n\t" {
            case "foo bar":
                v23 := GenS3{F0: 270}
                _ = v23
            }
            _ = v22
        }
        _ = v21
    } else {
        if false {
            var v24 int16 = -32768
            var v25 bool = (false || false)
            if (1 > -9223372036854775808) {
                var v26 int = -850
                for fi27 := 0; fi27 < 10; fi27++ {
                    if (("a" + "hello") == "a") {
                        if (v25 && (v25 && v25)) {
                            for fi28 := 0; fi28 < 4; fi28++ {
                                p1 = (p0 * 42)
                                switch "foo bar" {
                                case "":
                                    v29 := GenS2{F0: 127, F1: -9223372036854775808}
                                    _ = v29
                                default:
                                    v25 = ((1 < 0) || (!true))
                                }
                                if (v25 || (v25 && (1 <= 0))) {
                                    if v25 {
                                        var v30 float32 = (((1e10 / (1e10 / 3.14159)) + float32(1.5)) - 0.0)
                                        switch "\n\t" {
                                        case "a":
                                            p0 -= (p1 / (((p0 / (((p1 - p0)) | 1))) | 1))
                                        case "foo bar":
                                            var v31 uint16 = 65535
                                            _ = v31
                                        case "":
                                            v32 := GenS2{F0: p0, F1: 9223372036854775807}
                                            var v33 int = int(v24)
                                            var v34 uint64 = 0
                                            _ = v32
                                            _ = v33
                                            _ = v34
                                        }
                                        var v35 uint = 791
                                        _ = v30
                                        _ = v35
                                    }
                                }
                                _ = fi28
                            }
                        } else {
                            var v36 uint32 = 2147483647
                            if v25 {
                                if (!((v36 << ((65535) & 63)) < v36)) {
                                    var v37 uint = 18446744073709551615
                                    fmt.Println(v37)
                                    _ = v37
                                }
                                switch "a" {
                                case "":
                                    if false {
                                        v38 := []int{-629, fi27}
                                        _ = v38
                                    } else {
                                        if v25 {
                                            v36 -= v36
                                            var v39 int16 = v24
                                            switch "hello" {
                                            case "\n\t":
                                                for fi40 := 0; fi40 < 18; fi40++ {
                                                    v26 -= v26
                                                    v41 := []int{}
                                                    _ = fi40
                                                    _ = v41
                                                }
                                            case "hello":
                                                switch 32767 {
                                                case 9223372036854775807:
                                                    v42 := []int{}
                                                    _ = v42
                                                case 127:
                                                    var v43 int64 = -780
                                                    _ = v43
                                                }
                                            case "a":
                                                v44 := GenS3{F0: 4294967295}
                                                _ = v44
                                            }
                                            _ = v39
                                        } else {
                                            v24 -= (v24 + 1)
                                            v45 := []int{(fi27 >> ((855) & 63))}
                                            fmt.Println(p1, v25, p1)
                                            _ = v45
                                        }
                                        var v46 int64 = -9223372036854775808
                                        v47 := GenS1{F0: 179}
                                        _ = v46
                                        _ = v47
                                    }
                                }
                            }
                            _ = v36
                        }
                        var v48 int32 = 0
                        _ = v48
                    }
                    p1 = p1
                    _ = fi27
                }
                v49 := GenS1{F0: 664}
                _ = v26
                _ = v49
            }
            _ = v24
            _ = v25
        }
        var v50 uint64 = 688
        _ = v50
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 int64 = 0
    var v2 string = "foo bar"
    v3 := make(map[uint32]uint32)
    var v4 float32 = (0.1 - (1.5 - 1e10))
    v5 := []string{(("a" + "hello") + "a"), v2, v2, (v2 + v2)}
    var v6 float32 = v4
    v51 := GenS4{F0: 2, F1: 32767, F2: 24}
    var v52 int32 = 42
    if false {
        v4 -= (v4 / v6)
    } else {
        _ = len(v3)
        v5 = append(v5, "\n\t")
    }
    var v53 uint8 = genFn1(1, 128)
    for fi54 := 0; fi54 < 6; fi54++ {
        var v55 int = fi54
        _ = fi54
        _ = v55
    }
    for fi56 := 0; fi56 < 19; fi56++ {
        if (fi56 > (fi56 ^ (fi56 & (fi56 * -321)))) {
            if false {
                v52 = v52
            } else {
                if ((float64(v6) / 1.5) != (1e10 / 1e10)) {
                    _ = len(v5)
                    v57 := GenS5{F0: 128, F1: v53}
                    var v58 uint16 = 1
                    _ = v57
                    _ = v58
                } else {
                    switch ("" + v2) {
                    case "":
                        v53 += (v53 + v53)
                    case "a":
                        v2 = ""
                    case "\n\t":
                        for fi59 := 0; fi59 < 4; fi59++ {
                            v60 := GenS6{F0: false, F1: 255, F2: 0}
                            _ = fi59
                            _ = v60
                        }
                    default:
                        v2 = ((("" + v2) + (("foo bar" + ("a" + v2)) + "hello")) + ("\n\t" + (v2 + v2)))
                    }
                    var v61 uint16 = 577
                    var v62 int16 = 137
                    _ = v61
                    _ = v62
                }
            }
        } else {
            v2 = ""
            var v63 uint8 = 128
            for fi64 := 0; fi64 < 15; fi64++ {
                v65 := GenS2{F0: 22, F1: (v1 / (((v1 + (v1 + (v1 + int64(fi64))))) | 1))}
                if true {
                    if ((1.5 == v4) && false) {
                        switch ("hello" + ("" + v2)) {
                        case "":
                            switch v2 {
                            case "foo bar":
                                v66 := GenS3{F0: 234}
                                _ = v66
                            case "a":
                                switch v2 {
                                case "":
                                    v53 = v63
                                    fmt.Println(fi56, v6)
                                default:
                                    v63 -= v53
                                }
                            }
                        default:
                            v67 := GenS1{F0: 4294967295}
                            _ = v67
                        }
                    } else {
                        if false {
                            if ("hello" != v2) {
                                if false {
                                    if false {
                                        var v68 float32 = (v4 - v4)
                                        v4 += (((((3.14159 - (float32((float64(v6) * 1e10)) * ((0.1 / v68) - -2.25))) + v4) + v6) + (((v4 + 0.1) * ((v6 / v4) * v68)) + v4)) / v6)
                                        v3[2147483647] = 721
                                        _ = v68
                                    } else {
                                        v69 := GenS7{F0: -1}
                                        v70 := GenS8{F0: 42, F1: (v63 ^ (v53 >> ((255) & 63))), F2: (fi64 % (((fi64 << ((65535) & 63))) | 1))}
                                        switch 32768 {
                                        case 0:
                                            if true {
                                                v4 = v4
                                                v1 += 1
                                            }
                                        default:
                                            v71 := GenS5{F0: v53, F1: (^(v53))}
                                            _ = v71
                                        }
                                        _ = v69
                                        _ = v70
                                    }
                                    var v72 uint64 = 0
                                    _ = v72
                                }
                            } else {
                                var v73 uint16 = 0
                                var v74 int64 = -1
                                _ = v73
                                _ = v74
                            }
                            v2 = v2
                            var v75 string = (v2 + "\n\t")
                            _ = v75
                        }
                        v76 := GenS9{F0: 2}
                        if (!true) {
                            v63 += 49
                        }
                        _ = v76
                    }
                } else {
                    v5 = append(v5, (((((v2 + "a") + v2) + v2) + "a") + v2))
                    switch ("foo bar" + ("\n\t" + "a")) {
                    case "":
                        switch 128 {
                        case -550:
                            v63 = (v63 >> ((2) & 63))
                        default:
                            fmt.Println(v4)
                        }
                    case "foo bar":
                        if false {
                            switch (fi56 ^ (fi64 << ((4294967295) & 63))) {
                            case 9223372036854775807:
                                v77 := GenS4{F0: 0, F1: 880, F2: 885}
                                _ = v77
                            default:
                                switch 32768 {
                                case 127:
                                    if true {
                                        if false {
                                            var v78 uint16 = 2
                                            switch fi64 {
                                            case -9223372036854775808:
                                                v79 := GenS2{F0: 2, F1: v1}
                                                _ = v79
                                            case 42:
                                                var v80 uint16 = 32767
                                                _ = v80
                                            case 128:
                                                v81 := GenS10{F0: -9223372036854775808, F1: 65535}
                                                _ = v81
                                            }
                                            _ = v78
                                        }
                                        v82 := GenS11{F0: 0, F1: 255}
                                        _ = v82
                                    }
                                case 2147483647:
                                    var v83 uint64 = 0
                                    _ = v83
                                case -984:
                                    var v84 float64 = ((1.5 - 3.14159) * 3.14159)
                                    _ = v84
                                }
                            }
                            v85 := GenS11{F0: (^(v53)), F1: 0}
                            _ = v85
                        } else {
                            v86 := GenS3{F0: 262}
                            fmt.Println(v2, v52, v4)
                            _ = v86
                        }
                    }
                }
                _ = fi64
                _ = v65
            }
            _ = v63
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi56
    }
    v87 := GenS2{F0: (^(v53)), F1: 0}
    if (true && (1 > 729)) {
        var v88 float64 = (-2.25 + (3.14159 - -2.25))
        if false {
            for fi89 := 0; fi89 < 15; fi89++ {
                var v90 uint8 = v53
                _ = fi89
                _ = v90
            }
            v91 := GenS12{F0: 2, F1: 411}
            var v92 string = (v2 + "")
            _ = v91
            _ = v92
        } else {
            v4 += 3.14159
        }
        _ = v88
    } else {
        switch "hello" {
        case "hello":
            if ((1.5 - -2.25) < 3.14159) {
                for fi93 := 0; fi93 < 4; fi93++ {
                    v3[570] = 431
                    v94 := GenS4{F0: 0, F1: (^(fi93)), F2: 42}
                    _ = fi93
                    _ = v94
                }
                v95 := GenS13{F0: 42, F1: "a"}
                v53 -= v53
                _ = v95
            } else {
                switch 4294967295 {
                case -9223372036854775808:
                    v6 = float32(1e10)
                case -1:
                    fmt.Println(v52)
                }
                var v96 int64 = (v1 * int64(v53))
                _ = v96
            }
        case "\n\t":
            if (1 >= -159) {
                if (((-1 != -1) || false) || true) {
                    if true {
                        for fi97 := 0; fi97 < 17; fi97++ {
                            var v98 uint8 = (v53 % (((v53 % (((v53 / ((1) | 1))) | 1))) | 1))
                            v1 = v1
                            _ = fi97
                            _ = v98
                        }
                        var v99 int32 = (v52 / ((42) | 1))
                        v5 = append(v5, v2)
                        _ = v99
                    } else {
                        v100 := v3[245]
                        var v101 int16 = 0
                        _ = v100
                        _ = v101
                    }
                    delete(v3, 500)
                    v102 := GenS5{F0: v53, F1: (^(v53))}
                    _ = v102
                } else {
                    switch 1 {
                    case -902:
                        switch 1 {
                        case 830:
                            v53 -= (v53 + (v53 << ((725) & 63)))
                        case 4294967295:
                            for fi103 := 0; fi103 < 8; fi103++ {
                                for fi104 := 0; fi104 < 19; fi104++ {
                                    var v105 uint16 = 132
                                    _ = fi104
                                    _ = v105
                                }
                                v106 := GenS13{F0: 855, F1: "\n\t"}
                                for fi107 := 0; fi107 < 15; fi107++ {
                                    v1 = 9223372036854775807
                                    _ = fi107
                                }
                                _ = fi103
                                _ = v106
                            }
                        case 127:
                            fmt.Println(v52)
                        }
                    case 32768:
                        fmt.Println(v52, v6)
                    case -9223372036854775808:
                        for fi108 := 0; fi108 < 19; fi108++ {
                            for fi109 := 0; fi109 < 7; fi109++ {
                                if (!(float64(v6) <= float64(float32(0.0)))) {
                                    if ((genFn1(255, v53) < (^(v53))) && true) {
                                        var v110 uint32 = 2147483648
                                        var v111 uint16 = 2
                                        _ = v110
                                        _ = v111
                                    } else {
                                        var v112 uint16 = 1
                                        _ = v112
                                    }
                                    var v113 int16 = 453
                                    _ = v113
                                }
                                _ = fi109
                            }
                            _ = fi108
                        }
                    }
                    fmt.Println(v53, v52, v53)
                }
            } else {
                fmt.Println(v6, v1, v2)
            }
        case "a":
            switch v2 {
            case "hello":
                v6 = 3.14159
            case "\n\t":
                switch "" {
                case "a":
                    fmt.Println(v53, v52)
                }
            default:
                for fi114 := 0; fi114 < 19; fi114++ {
                    var v115 uint64 = 2
                    var v116 float64 = 1.5
                    _ = fi114
                    _ = v115
                    _ = v116
                }
            }
        default:
            var v117 int = -714
            _ = v117
        }
    }
    switch "\n\t" {
    case "":
        v6 += v6
    default:
        if true {
            v53 -= (^(v53))
        } else {
            for fi118 := 0; fi118 < 16; fi118++ {
                v3[2147483647] = 65535
                switch "foo bar" {
                case "":
                    var v119 int64 = v1
                    _ = v119
                case "hello":
                    switch fi118 {
                    case 127:
                        v6 += v6
                    }
                case "foo bar":
                    fmt.Println(v6, v6, v52)
                }
                if ((1e10 + v4) == v4) {
                    fmt.Println(v6)
                } else {
                    v120 := GenS1{F0: 4294967295}
                    _ = v120
                }
                _ = fi118
            }
        }
    }
    v2 = v2
    var v121 int16 = 0
    switch -280 {
    case -2147483648:
        v53 = (v53 ^ v53)
    case -9223372036854775808:
        v1 = v1
    default:
        v6 = ((0.1 * v4) / 1.0)
    }
    if (!false) {
        switch "foo bar" {
        case "\n\t":
            switch (v2 + "\n\t") {
            case "":
                v1 = v1
            }
        case "":
            for fi122 := 0; fi122 < 14; fi122++ {
                for fi123 := 0; fi123 < 8; fi123++ {
                    var v124 int = (fi123 - fi123)
                    _ = fi123
                    _ = v124
                }
                _ = fi122
            }
        case "hello":
            var v125 float64 = float64(((((v4 / (float32((1.5 / 1e10)) - (v4 * v6))) * -2.25) + v6) + ((v4 - (0.0 + 0.1)) + (v4 * 3.14159))))
            _ = v125
        }
        if true {
            delete(v3, 65535)
            var v126 int16 = 1000
            v6 = v6
            _ = v126
        } else {
            for fi127 := 0; fi127 < 11; fi127++ {
                v128 := GenS11{F0: (v53 * v53), F1: 1}
                for fi129 := 0; fi129 < 8; fi129++ {
                    var v130 int8 = 13
                    _ = fi129
                    _ = v130
                }
                _ = fi127
                _ = v128
            }
        }
    } else {
        for fi131 := 0; fi131 < 2; fi131++ {
            v6 += v6
            v132 := v3[795]
            _ = fi131
            _ = v132
        }
    }
    for fi133 := 0; fi133 < 3; fi133++ {
        var v134 uint16 = 150
        _ = fi133
        _ = v134
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v3[2] = 1
    fmt.Println("v1:", v1)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v6
    _ = v87
}
