package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 int64; F2 uint8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int16; F1 uint16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int32; F1 uint; F2 uint64 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint16; F1 int16 }
type GenS14 struct { F0 string; F1 int64 }
type GenS15 struct { F0 string; F1 int32; F2 int16 }
func genFn0(p0 uint, p1 uint64) int {
    p0 = (p0 >> ((1) & 63))
    var v4 float64 = 0.0
    switch 2147483647 {
    case 1:
        for fi5 := 0; fi5 < 10; fi5++ {
            v6 := GenS1{F0: 9223372036854775807}
            _ = fi5
            _ = v6
        }
    default:
        if (-1 >= 980) {
            v7 := []int{32768, 32767}
            var v8 int = -9223372036854775808
            v8 += (v8 + (v8 | (v8 | (v8 | v8))))
            _ = v7
            _ = v8
        } else {
            var v9 int8 = 1
            if true {
                var v10 uint16 = 32767
                v11 := []int{1, 9223372036854775807, 127}
                _ = v10
                _ = v11
            } else {
                switch int(p1) {
                case 662:
                    for fi12 := 0; fi12 < 4; fi12++ {
                        switch (fi12 >> (((p0 >> ((p0) & 63))) & 63)) {
                        case 469:
                            v9 -= (v9 % ((107) | 1))
                        }
                        v4 += (0.1 + 0.1)
                        _ = fi12
                    }
                case -785:
                    v9 = v9
                default:
                    var v13 bool = false
                    _ = v13
                }
                if (v4 == (1.5 + ((v4 - (v4 / v4)) * 1e10))) {
                    v14 := GenS1{F0: 748}
                    v15 := GenS2{F0: 32767}
                    _ = v14
                    _ = v15
                }
                v16 := GenS3{F0: 2147483647}
                _ = v16
            }
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    var v1 int32 = 42
    var v2 bool = (0 > 255)
    var v3 uint32 = 253
    var v17 float64 = 0.1
    var v18 uint32 = v3
    fmt.Println(v3, v2, v2)
    if (1 <= 1) {
        v17 += v17
        v18 += v3
    } else {
        if (42 <= 195) {
            v19 := []int{}
            _ = v19
        }
        fmt.Println(v3, v3)
        v20 := GenS4{F0: uint16(v3), F1: 328}
        _ = v20
    }
    v21 := []int{-337}
    v3 += v18
    v18 = 545
    fmt.Println(v3, v1, v17)
    if v2 {
        fmt.Println(v1, v3, v18)
        var v22 float64 = ((float64(1e10) + v17) + (float64(1e10) / (3.14159 - v17)))
        _ = v22
    } else {
        fmt.Println(v18)
        for fi23 := 0; fi23 < 8; fi23++ {
            for fi24 := 0; fi24 < 7; fi24++ {
                v25 := GenS5{F0: 95}
                _ = fi24
                _ = v25
            }
            fmt.Println(v1)
            _ = fi23
        }
        for fi26 := 0; fi26 < 8; fi26++ {
            v18 += 2
            var v27 bool = false
            _ = fi26
            _ = v27
        }
    }
    fmt.Println(v1, v18, v3)
    for fi28 := 0; fi28 < 7; fi28++ {
        if (0.0 != v17) {
            var v29 uint = 0
            _ = v29
        }
        var v30 int32 = 853
        if true {
            if v2 {
                var v31 uint16 = 0
                _ = v31
            } else {
                v32 := GenS5{F0: uint8(v18)}
                var v33 float64 = 3.14159
                for fi34 := 0; fi34 < 7; fi34++ {
                    _ = len(v21)
                    _ = fi34
                }
                _ = v32
                _ = v33
            }
        }
        _ = fi28
        _ = v30
    }
    var v35 float64 = 0.1
    if v2 {
        for fi36 := 0; fi36 < 12; fi36++ {
            fmt.Println(v35, v2)
            switch genFn0(2, 2) {
            case 42:
                for fi37 := 0; fi37 < 16; fi37++ {
                    v38 := GenS5{F0: 128}
                    _ = fi37
                    _ = v38
                }
            case 555:
                if (true && false) {
                    var v39 bool = (127 <= -1)
                    switch genFn0(673, 2) {
                    case 807:
                        if v39 {
                            v40 := GenS6{F0: 0}
                            for fi41 := 0; fi41 < 10; fi41++ {
                                var v42 string = ((("\n\t" + ("" + "foo bar")) + "foo bar") + "\n\t")
                                _ = fi41
                                _ = v42
                            }
                            if (fi36 >= 127) {
                                fmt.Println(v18, v18)
                            }
                            _ = v40
                        }
                    case 2147483647:
                        if len(v21) > 0 {
                            v21[0] = (fi36 | 32768)
                        }
                    }
                    fmt.Println(fi36)
                    _ = v39
                }
            case 32768:
                if v2 {
                    var v43 uint32 = v3
                    _ = v43
                } else {
                    fmt.Println(v1)
                    if ((1 < 217) && v2) {
                        if (-2.25 != 0.0) {
                            var v44 int16 = 42
                            if (!false) {
                                if (128 == 1) {
                                    var v45 float64 = ((3.14159 * float64((1e10 / ((0.1 * 0.1) * 0.1)))) * v35)
                                    _ = v45
                                }
                            } else {
                                var v46 uint16 = 2
                                if false {
                                    v47 := GenS3{F0: genFn0(255, 4294967295)}
                                    for fi48 := 0; fi48 < 19; fi48++ {
                                        v21 = append(v21, (fi36 * (fi36 ^ (fi48 - genFn0(18446744073709551615, 374)))))
                                        v18 += v3
                                        _ = fi48
                                    }
                                    for fi49 := 0; fi49 < 19; fi49++ {
                                        v18 = v3
                                        v50 := GenS7{F0: ((1e10 - 1e10) < (-2.25 * (0.1 * ((-2.25 * 0.1) * -2.25)))), F1: 0, F2: 1}
                                        _ = fi49
                                        _ = v50
                                    }
                                    _ = v47
                                }
                                _ = v46
                            }
                            switch fi36 {
                            case 9223372036854775807:
                                if false {
                                    switch ("a" + "a") {
                                    case "":
                                        fmt.Println(v17, v2, v2)
                                    }
                                    fmt.Println(v17, v1, fi36)
                                }
                            case -2147483648:
                                v21 = append(v21, (fi36 & (fi36 / (((fi36 / (((fi36 % ((921) | 1))) | 1))) | 1))))
                            case -1:
                                v17 += (float64(0.1) + 3.14159)
                            }
                            _ = v44
                        }
                    } else {
                        var v51 uint8 = 47
                        _ = v51
                    }
                    v35 = (v17 + (v35 * (v17 + (((v35 / v17) - 0.0) / v35))))
                }
            default:
                if (42 <= v1) {
                    switch fi36 {
                    case 128:
                        v17 = 0.0
                    case -2147483648:
                        var v52 int64 = -1
                        _ = v52
                    }
                } else {
                    if (v2 && (0 >= (^(v3)))) {
                        for fi53 := 0; fi53 < 2; fi53++ {
                            if false {
                                if true {
                                    var v54 int32 = -1
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    _ = v54
                                }
                            }
                            if (2 <= 801) {
                                var v55 uint8 = 128
                                _ = v55
                            }
                            v56 := GenS3{F0: (fi53 - -395)}
                            var v57 bool = v2
                            var v58 string = "a"
                            _ = fi53
                            _ = v56
                            _ = v57
                            _ = v58
                        }
                        if (506 != 255) {
                            v3 = 2147483648
                            var v59 uint = 4294967295
                            switch "hello" {
                            case "":
                                fmt.Println(fi36, v17, v35)
                            case "\n\t":
                                v60 := GenS8{F0: 238}
                                _ = v60
                            case "a":
                                v59 -= v59
                            }
                            _ = v59
                        }
                        v1 = 42
                    } else {
                        if ((1.5 * (1e10 - (-2.25 * -2.25))) >= 1.5) {
                            for fi61 := 0; fi61 < 4; fi61++ {
                                var v62 float32 = (3.14159 + 1e10)
                                if (!((v17 + 3.14159) > (((0.1 * 0.0) - v35) + v35))) {
                                    var v63 int8 = 0
                                    var v64 bool = (42 < 127)
                                    _ = v63
                                    _ = v64
                                } else {
                                    switch fi61 {
                                    case -1:
                                        switch "foo bar" {
                                        case "a":
                                            switch fi36 {
                                            case 127:
                                                switch (-(fi36)) {
                                                case 127:
                                                    var v65 bool = v2
                                                    _ = v65
                                                case -368:
                                                    fmt.Println(v3, v62, v35)
                                                    if false {
                                                        fmt.Println(v17, v17, fi61)
                                                        if (!(660 == 0)) {
                                                            var v66 int32 = v1
                                                            var v67 bool = true
                                                            v2 = false
                                                            _ = v66
                                                            _ = v67
                                                        } else {
                                                            if len(v21) >= 2 {
                                                                v21 = v21[1:len(v21)]
                                                            }
                                                            v68 := GenS9{F0: 42, F1: 255}
                                                            switch (^(fi61)) {
                                                            case 127:
                                                                var v69 uint32 = v18
                                                                _ = v69
                                                            case 2147483647:
                                                                v17 = (1e10 + (0.1 * ((1e10 * float64((v62 / ((v62 * ((v62 * (((v62 / v62) + 0.0) * v62)) * float32((v17 + 3.14159)))) - v62)))) / 1e10)))
                                                            case 20:
                                                                if (!(-1 != 0)) {
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        v1 = -2147483648
                                                                        var v70 float32 = float32(v17)
                                                                        if true {
                                                                            var v71 int64 = 9223372036854775807
                                                                            var v72 int16 = -1
                                                                            var v73 uint16 = 65535
                                                                            _ = v71
                                                                            _ = v72
                                                                            _ = v73
                                                                        }
                                                                        _ = v70
                                                                    case "":
                                                                        var v74 uint64 = 941
                                                                        _ = v74
                                                                    }
                                                                } else {
                                                                    if len(v21) >= 2 {
                                                                        v21 = v21[1:len(v21)]
                                                                    }
                                                                    fmt.Println(v62)
                                                                }
                                                            default:
                                                                var v75 int8 = 127
                                                                _ = v75
                                                            }
                                                            _ = v68
                                                        }
                                                        switch (fi61 << ((1) & 63)) {
                                                        case 4294967295:
                                                            v76 := GenS5{F0: 255}
                                                            _ = v76
                                                        case -9223372036854775808:
                                                            var v77 uint8 = 128
                                                            _ = v77
                                                        }
                                                    } else {
                                                        v18 -= v3
                                                        v78 := GenS10{F0: 176}
                                                        _ = v78
                                                    }
                                                    _ = len(v21)
                                                    if true {
                                                        if v2 {
                                                            var v79 string = "a"
                                                            var v80 float64 = float64(0.0)
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                    } else {
                                                        v2 = false
                                                        fmt.Println(v2)
                                                        if len(v21) >= 2 {
                                                            v21 = v21[1:len(v21)]
                                                        }
                                                    }
                                                case 42:
                                                    var v81 float64 = 0.0
                                                    _ = v81
                                                }
                                            case 868:
                                                var v82 uint64 = 652
                                                _ = v82
                                            case 1:
                                                if (!((fi36 - fi61) < (fi36 - (fi36 >> ((0) & 63))))) {
                                                    var v83 int8 = 1
                                                    if len(v21) > 0 {
                                                        v21[0] = (fi61 >> ((uint(v18)) & 63))
                                                    }
                                                    v84 := GenS10{F0: 197}
                                                    _ = v83
                                                    _ = v84
                                                } else {
                                                    fmt.Println(v18)
                                                    if v2 {
                                                        if len(v21) > 0 {
                                                            v21[0] = (fi36 + -76)
                                                        }
                                                        if len(v21) >= 2 {
                                                            v21 = v21[1:len(v21)]
                                                        }
                                                        if (0 != 302) {
                                                            v85 := GenS1{F0: (fi61 % (((-(fi36))) | 1))}
                                                            v86 := GenS11{F0: v1, F1: 775, F2: 2}
                                                            _ = v85
                                                            _ = v86
                                                        } else {
                                                            if true {
                                                                v3 += v18
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    _ = len(v21)
                                                                case "a":
                                                                    if (!v2) {
                                                                        var v87 int32 = 1
                                                                        _ = v87
                                                                    } else {
                                                                        if len(v21) >= 2 {
                                                                            v21 = v21[1:len(v21)]
                                                                        }
                                                                        var v88 string = "foo bar"
                                                                        v21 = append(v21, (fi36 ^ (^(fi36))))
                                                                        _ = v88
                                                                    }
                                                                default:
                                                                    if false {
                                                                        if v2 {
                                                                            if false {
                                                                                v21 = append(v21, 2147483647)
                                                                            } else {
                                                                                fmt.Println(v62, v18, v17)
                                                                                switch "a" {
                                                                                case "":
                                                                                    var v89 bool = (-889 != -58)
                                                                                    _ = v89
                                                                                case "hello":
                                                                                    if false {
                                                                                        if v2 {
                                                                                            v90 := GenS4{F0: 299, F1: 0}
                                                                                            switch (-(fi61)) {
                                                                                            case 32768:
                                                                                                if (!(0 > -501)) {
                                                                                                    if len(v21) > 0 {
                                                                                                        v21[0] = (fi36 - int(v1))
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v2 = v2
                                                                                            var v91 uint = 682
                                                                                            v1 += v1
                                                                                            _ = v90
                                                                                            _ = v91
                                                                                        } else {
                                                                                            v62 = (0.1 / (v62 - ((((((v62 / float32(v17)) / (v62 + 3.14159)) / v62) + ((float32((1.5 - v17)) - 3.14159) + float32(float64((3.14159 - v62))))) + 0.0) + (v62 - 3.14159))))
                                                                                        }
                                                                                        v92 := GenS12{F0: (-131 == 479)}
                                                                                        var v93 uint32 = v18
                                                                                        _ = v92
                                                                                        _ = v93
                                                                                    }
                                                                                default:
                                                                                    v17 = v17
                                                                                }
                                                                                fmt.Println(v1)
                                                                                if v2 {
                                                                                    var v94 uint32 = 2147483648
                                                                                    v1 = (v1 & v1)
                                                                                    _ = v94
                                                                                } else {
                                                                                    _ = len(v21)
                                                                                    if len(v21) > 0 {
                                                                                        v21[0] = -371
                                                                                    }
                                                                                    v18 += (^(v18))
                                                                                }
                                                                            }
                                                                        }
                                                                        switch ("foo bar" + "\n\t") {
                                                                        case "a":
                                                                            v17 = -2.25
                                                                        case "\n\t":
                                                                            if len(v21) > 0 {
                                                                                v21[0] = (fi61 + (-(fi36)))
                                                                            }
                                                                        default:
                                                                            var v95 uint8 = 128
                                                                            _ = v95
                                                                        }
                                                                    }
                                                                }
                                                                if len(v21) >= 2 {
                                                                    v21 = v21[1:len(v21)]
                                                                }
                                                            } else {
                                                                switch (fi36 & -9223372036854775808) {
                                                                case 56:
                                                                    if len(v21) > 0 {
                                                                        v21[0] = genFn0(18446744073709551615, 42)
                                                                    }
                                                                case 1:
                                                                    var v96 uint64 = 2
                                                                    _ = v96
                                                                default:
                                                                    fmt.Println(v2, v1, v2)
                                                                }
                                                                switch "" {
                                                                case "\n\t":
                                                                    fmt.Println(v1, v3, v62)
                                                                case "foo bar":
                                                                    v97 := GenS9{F0: 1000, F1: 255}
                                                                    _ = v97
                                                                default:
                                                                    var v98 int8 = 0
                                                                    _ = v98
                                                                }
                                                                var v99 int = 32768
                                                                _ = v99
                                                            }
                                                            var v100 bool = true
                                                            fmt.Println(fi61, fi61)
                                                            _ = v100
                                                        }
                                                    } else {
                                                        fmt.Println(v62, fi61)
                                                    }
                                                    var v101 string = "a"
                                                    v35 = v35
                                                    _ = v101
                                                }
                                            }
                                        case "":
                                            _ = len(v21)
                                        default:
                                            if (!(!(!v2))) {
                                                _ = len(v21)
                                                var v102 int16 = 0
                                                v103 := GenS5{F0: 12}
                                                _ = v102
                                                _ = v103
                                            }
                                        }
                                    default:
                                        if v2 {
                                            var v104 bool = (127 >= 42)
                                            _ = len(v21)
                                            v105 := GenS13{F0: 448, F1: 32767}
                                            _ = v104
                                            _ = v105
                                        } else {
                                            v106 := GenS14{F0: ("" + "foo bar"), F1: -1}
                                            var v107 int16 = 1000
                                            _ = v106
                                            _ = v107
                                        }
                                    }
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                }
                                var v108 int64 = -898
                                v109 := GenS2{F0: -766}
                                if (!v2) {
                                    fmt.Println(fi61, fi36)
                                    if false {
                                        if ((v108 & (v108 - (v108 >> ((1) & 63)))) != v108) {
                                            v110 := GenS15{F0: (((("" + ("\n\t" + "")) + ("hello" + "hello")) + "\n\t") + (("" + "a") + ("a" + ("" + "")))), F1: v1, F2: 658}
                                            if (1 <= 42) {
                                                var v111 int8 = 113
                                                if v2 {
                                                    fmt.Println(v35, v3, v35)
                                                }
                                                var v112 string = ("hello" + "")
                                                _ = v111
                                                _ = v112
                                            }
                                            _ = v110
                                        }
                                        if len(v21) >= 2 {
                                            v21 = v21[1:len(v21)]
                                        }
                                    }
                                }
                                _ = fi61
                                _ = v108
                                _ = v109
                                _ = v62
                            }
                            fmt.Println(v3, v18)
                            v1 = (v1 * -81)
                        }
                    }
                    fmt.Println(v18)
                    if len(v21) > 0 {
                        v21[0] = (fi36 - 4294967295)
                    }
                }
            }
            _ = fi36
        }
    } else {
        for fi113 := 0; fi113 < 7; fi113++ {
            fmt.Println(v2, v35, v3)
            for fi114 := 0; fi114 < 7; fi114++ {
                var v115 int16 = 710
                var v116 int32 = v1
                fmt.Println(v1, v17)
                _ = fi114
                _ = v115
                _ = v116
            }
            _ = fi113
        }
        v35 = float64(0.0)
        v17 = v17
    }
    if (42 != -387) {
        var v117 uint16 = 255
        var v118 uint16 = 550
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
        _ = v117
        _ = v118
    }
    _ = len(v21)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v35
}
