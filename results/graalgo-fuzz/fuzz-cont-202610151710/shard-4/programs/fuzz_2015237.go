package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint16; F1 uint; F2 bool }
type GenS5 struct { F0 uint16; F1 int; F2 uint64 }
type GenS6 struct { F0 bool; F1 int16; F2 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int8; F1 int16; F2 uint8 }
type GenS12 struct { F0 bool; F1 int8; F2 bool }
func genFn0(p0 bool, p1 int8, p2 int32) int16 {
    if (p1 == 42) {
        for fi4 := 0; fi4 < 19; fi4++ {
            var v5 uint64 = 29
            v6 := GenS1{F0: (v5 / ((v5) | 1)), F1: 899}
            fmt.Println(p1)
            _ = fi4
            _ = v5
            _ = v6
        }
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := []int{fi7, fi7, 128, (-(fi7))}
            v8 = append(v8, (fi7 & 2147483647))
            p0 = p0
            _ = fi7
            _ = v8
        }
    }
    if p0 {
        fmt.Println(p1)
    } else {
        switch "foo bar" {
        case "":
            p2 -= (-(p2))
        case "hello":
            switch "hello" {
            case "\n\t":
                if (0.1 != 0.0) {
                    switch ("hello" + "") {
                    case "\n\t":
                        v9 := []int{-2147483648, 2147483647, -746, -9223372036854775808}
                        _ = v9
                    default:
                        switch (("a" + "hello") + "\n\t") {
                        case "\n\t":
                            p0 = ((p1 % ((p1) | 1)) == (p1 | (p1 & (p1 % ((p1) | 1)))))
                        default:
                            p2 = (p2 ^ p2)
                        }
                    }
                    p0 = (2 >= 42)
                }
            default:
                var v10 uint = 1
                _ = v10
            }
        }
        if false {
            var v11 int16 = 1
            v12 := GenS2{F0: 127, F1: 0}
            p1 = p1
            _ = v11
            _ = v12
        } else {
            var v13 string = ((("\n\t" + "") + "foo bar") + "")
            _ = v13
        }
        if true {
            v14 := GenS3{F0: 42, F1: 18446744073709551615}
            v15 := []int{-166, 4294967295, -9223372036854775808}
            _ = v14
            _ = v15
        }
    }
    var v16 uint8 = 128
    if p0 {
        v17 := []int{128, 0, 32767}
        for fi18 := 0; fi18 < 1; fi18++ {
            v16 = 236
            if (377 == 675) {
                var v19 string = (("a" + "a") + ("" + "foo bar"))
                v20 := GenS3{F0: 9223372036854775807, F1: 1}
                for fi21 := 0; fi21 < 7; fi21++ {
                    if true {
                        if p0 {
                            p0 = ((p0 && false) || p0)
                            fmt.Println(p0)
                            fmt.Println(fi21, p1, p2)
                        }
                        var v22 uint64 = 0
                        fmt.Println(v19, fi18, p1)
                        p0 = ((!p0) && (!(!p0)))
                        v23 := GenS1{F0: (^(v22)), F1: 2}
                        _ = v22
                        _ = v23
                    }
                    var v24 int16 = 525
                    _ = fi21
                    _ = v24
                }
                _ = v19
                _ = v20
            }
            _ = fi18
        }
        _ = v17
    } else {
        var v25 int = -2147483648
        p1 = -128
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return -32768
}
func genFn1(p0 uint64) int64 {
    if true {
        if true {
            v26 := []int{42, 0, 1}
            for fi27 := 0; fi27 < 20; fi27++ {
                if false {
                    p0 = (p0 % ((p0) | 1))
                } else {
                    var v28 uint16 = 208
                    _ = v28
                }
                _ = fi27
            }
            _ = v26
        } else {
            fmt.Println(p0)
        }
        v29 := GenS1{F0: p0, F1: 523}
        var v30 bool = true
        _ = v29
        _ = v30
    }
    v31 := GenS1{F0: p0, F1: 42}
    var v32 uint32 = 203
    _ = p0
    _ = v31
    _ = v32
    return 9223372036854775807
}

func main() {
    var v1 int32 = 0
    var v2 uint64 = 18446744073709551615
    var v3 int64 = 42
    v33 := GenS4{F0: 1, F1: 4294967295, F2: (!true)}
    var v34 int64 = (v3 | -1)
    var v35 float32 = 1e10
    fmt.Println(v1)
    switch "foo bar" {
    case "a":
        var v36 uint32 = 42
        _ = v36
    case "foo bar":
        fmt.Println(v34, v35, v2)
    case "\n\t":
        v35 += (v35 - 1e10)
    default:
        if ((730 > 42) || ("foo bar" != "hello")) {
            for fi37 := 0; fi37 < 17; fi37++ {
                if true {
                    if (2 != 0) {
                        v38 := []int{1, (fi37 & (fi37 / ((fi37) | 1))), fi37}
                        _ = v38
                    }
                } else {
                    if ((1e10 - (0.0 - v35)) == (v35 / 1e10)) {
                        var v39 uint8 = 42
                        if true {
                            v40 := []int{}
                            fmt.Println(v2, v35, v2)
                            v40 = append(v40, 876)
                            _ = v40
                        } else {
                            v1 -= (v1 | 0)
                            fmt.Println(v1)
                            if false {
                                var v41 int16 = -32768
                                var v42 uint32 = 0
                                _ = v41
                                _ = v42
                            }
                        }
                        var v43 uint16 = 2
                        _ = v39
                        _ = v43
                    } else {
                        switch (fi37 << ((434) & 63)) {
                        case -968:
                            v44 := GenS3{F0: (v3 + v34), F1: v2}
                            _ = v44
                        case -660:
                            switch fi37 {
                            case 32767:
                                fmt.Println(v3, v2)
                            }
                        default:
                            v45 := GenS1{F0: (v2 & v2), F1: 1}
                            _ = v45
                        }
                        if false {
                            var v46 uint32 = 586
                            for fi47 := 0; fi47 < 3; fi47++ {
                                v48 := GenS4{F0: 1, F1: 42, F2: (v1 != v1)}
                                switch fi47 {
                                case 127:
                                    var v49 int = fi47
                                    _ = v49
                                }
                                v50 := GenS3{F0: v34, F1: v2}
                                _ = fi47
                                _ = v48
                                _ = v50
                            }
                            if (11 != -128) {
                                for fi51 := 0; fi51 < 18; fi51++ {
                                    fmt.Println(v46, fi37, v35)
                                    _ = fi51
                                }
                                var v52 int16 = genFn0((255 >= 1), -128, v1)
                                _ = v52
                            }
                            _ = v46
                        } else {
                            v53 := GenS1{F0: v2, F1: 65535}
                            v1 = v1
                            var v54 string = "a"
                            _ = v53
                            _ = v54
                        }
                        for fi55 := 0; fi55 < 15; fi55++ {
                            var v56 int = fi37
                            v35 = 1.5
                            var v57 float32 = 1.5
                            v58 := []int{4294967295, fi55, (fi37 + v56)}
                            switch "" {
                            case "":
                                v56 = int(v2)
                            case "hello":
                                if ((!("hello" <= (("" + "hello") + ("" + ("foo bar" + "hello"))))) || true) {
                                    v57 -= v35
                                    v58 = append(v58, (v56 << ((0) & 63)))
                                    var v59 uint8 = 90
                                    _ = v59
                                } else {
                                    if ("a" != "a") {
                                        fmt.Println(fi55, v3)
                                        v1 = v1
                                        v35 = -2.25
                                    } else {
                                        v34 = (v34 % ((1) | 1))
                                        v57 += (v35 * v57)
                                    }
                                }
                            case "\n\t":
                                if ((0.1 * float64(v35)) != 1.5) {
                                    v34 -= v34
                                    var v60 bool = (true || (("a" + "hello") < ("" + (("a" + (("foo bar" + "") + ("a" + "hello"))) + ("a" + (("hello" + "foo bar") + "foo bar"))))))
                                    fmt.Println(v57, v57)
                                    var v61 uint16 = 255
                                    _ = v60
                                    _ = v61
                                } else {
                                    if len(v58) >= 2 {
                                        v58 = v58[1:len(v58)]
                                    }
                                    var v62 int32 = (v1 + v1)
                                    if true {
                                        fmt.Println(v1)
                                        if (false && false) {
                                            var v63 bool = true
                                            if (0 < 42) {
                                                var v64 uint8 = 255
                                                var v65 int = fi37
                                                if len(v58) >= 2 {
                                                    v58 = v58[1:len(v58)]
                                                }
                                                _ = v64
                                                _ = v65
                                            } else {
                                                v58 = append(v58, fi37)
                                                if len(v58) > 0 {
                                                    v58[0] = fi37
                                                }
                                                fmt.Println(fi37, v62)
                                            }
                                            _ = v63
                                        }
                                        v66 := GenS2{F0: 1, F1: 42}
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v67 float32 = v57
                                            _ = v67
                                        case "":
                                            v68 := GenS4{F0: 65535, F1: 417, F2: (!false)}
                                            if true {
                                                fmt.Println(v2)
                                            }
                                            var v69 float64 = 1e10
                                            v58 = append(v58, (v56 << ((255) & 63)))
                                            _ = v68
                                            _ = v69
                                        default:
                                            fmt.Println(fi37, v1)
                                        }
                                        if true {
                                            v70 := GenS5{F0: 65535, F1: fi55, F2: v2}
                                            if false {
                                                switch "a" {
                                                case "a":
                                                    v3 -= (v34 * v3)
                                                case "foo bar":
                                                    v34 += (v3 / (((v3 ^ 42)) | 1))
                                                    var v71 int32 = (v62 & (^(v1)))
                                                    if (594 < 1) {
                                                        v35 = v35
                                                        if len(v58) >= 2 {
                                                            v58 = v58[1:len(v58)]
                                                        }
                                                        var v72 uint16 = 42
                                                        _ = v72
                                                    }
                                                    _ = v71
                                                case "":
                                                    var v73 uint16 = 327
                                                    _ = v73
                                                }
                                            }
                                            _ = v70
                                        }
                                        _ = v66
                                    }
                                    _ = v62
                                }
                            default:
                                if false {
                                    _ = len(v58)
                                } else {
                                    fmt.Println(v2, v2)
                                    fmt.Println(fi55)
                                    if (("" != ("hello" + (("" + (("foo bar" + "a") + "foo bar")) + "foo bar"))) || (!(true || false))) {
                                        switch -472 {
                                        case 9223372036854775807:
                                            fmt.Println(v2, v2, v56)
                                        }
                                        v74 := GenS4{F0: 356, F1: 0, F2: true}
                                        _ = v74
                                    }
                                    switch "" {
                                    case "foo bar":
                                        var v75 float32 = v57
                                        _ = v75
                                    case "a":
                                        var v76 uint8 = 65
                                        _ = v76
                                    default:
                                        if len(v58) > 0 {
                                            v58[0] = (-(fi55))
                                        }
                                    }
                                    if (!(v56 > v56)) {
                                        if true {
                                            if (!(1 <= 68)) {
                                                switch (-(v56)) {
                                                case 1:
                                                    var v77 bool = (!((18446744073709551615 == 1) || (false && false)))
                                                    _ = v77
                                                default:
                                                    var v78 int16 = -32768
                                                    _ = v78
                                                }
                                                if true {
                                                    if false {
                                                        v79 := GenS2{F0: 127, F1: (fi55 - (fi55 / ((v56) | 1)))}
                                                        var v80 float32 = (float32(-2.25) - (v57 + (3.14159 * 3.14159)))
                                                        _ = v79
                                                        _ = v80
                                                    } else {
                                                        fmt.Println(fi37)
                                                        var v81 int = 283
                                                        switch ("" + "a") {
                                                        case "hello":
                                                            switch (^(v81)) {
                                                            case 2147483647:
                                                                if (!((-32768 <= 1) && ((v56 + v56) >= -830))) {
                                                                    v82 := GenS6{F0: (!true), F1: -32768, F2: 990}
                                                                    if true {
                                                                        v56 += v81
                                                                    }
                                                                    if len(v58) >= 2 {
                                                                        v58 = v58[1:len(v58)]
                                                                    }
                                                                    _ = v82
                                                                }
                                                            }
                                                        default:
                                                            v1 = -2147483648
                                                        }
                                                        v35 = (1.5 - (v57 * ((v35 + v35) + (v35 + v57))))
                                                        fmt.Println(v56, v57)
                                                        switch "hello" {
                                                        case "":
                                                            v3 -= (v34 % ((v34) | 1))
                                                        default:
                                                            v83 := GenS7{F0: 389}
                                                            _ = v83
                                                        }
                                                        switch fi37 {
                                                        case -967:
                                                            _ = len(v58)
                                                        case 127:
                                                            var v84 bool = false
                                                            _ = v84
                                                        case -9223372036854775808:
                                                            v57 = v35
                                                        }
                                                        _ = v81
                                                    }
                                                }
                                            } else {
                                                var v85 float64 = 0.0
                                                _ = v85
                                            }
                                            v58 = append(v58, (fi37 * (^(v56))))
                                            switch int(v1) {
                                            case 128:
                                                if len(v58) >= 2 {
                                                    v58 = v58[1:len(v58)]
                                                }
                                            case -1:
                                                v58 = append(v58, fi37)
                                            case 42:
                                                v86 := GenS8{F0: fi37, F1: 37}
                                                _ = v86
                                            default:
                                                if ((false && true) && (255 >= 42)) {
                                                    v87 := GenS9{F0: 415}
                                                    v56 = -484
                                                    v56 = (fi55 >> ((65535) & 63))
                                                    _ = v87
                                                } else {
                                                    if len(v58) > 0 {
                                                        v58[0] = (fi55 >> ((4294967295) & 63))
                                                    }
                                                    var v88 float32 = v57
                                                    _ = v88
                                                }
                                            }
                                        } else {
                                            if false {
                                                v34 += v3
                                                if len(v58) > 0 {
                                                    v58[0] = fi37
                                                }
                                                var v89 uint64 = (^(v2))
                                                _ = v89
                                            } else {
                                                var v90 uint16 = 255
                                                var v91 int = 2147483647
                                                fmt.Println(v91, v1)
                                                if false {
                                                    switch 32767 {
                                                    case 128:
                                                        v92 := GenS3{F0: v34, F1: 470}
                                                        _ = v92
                                                    case 131:
                                                        v3 = genFn1(v2)
                                                    default:
                                                        _ = len(v58)
                                                    }
                                                } else {
                                                    v91 -= 32768
                                                    v3 -= 195
                                                }
                                                if (0 >= 100) {
                                                    v34 -= genFn1(1)
                                                } else {
                                                    var v93 int64 = (^(v34))
                                                    v94 := GenS2{F0: 2, F1: 42}
                                                    _ = v93
                                                    _ = v94
                                                }
                                                _ = v90
                                                _ = v91
                                            }
                                            if (true && true) {
                                                var v95 string = ("\n\t" + ("\n\t" + "a"))
                                                if (true || false) {
                                                    v56 = int(v2)
                                                }
                                                v2 += 2
                                                _ = v95
                                            }
                                        }
                                    } else {
                                        var v96 string = "\n\t"
                                        var v97 int32 = v1
                                        _ = v96
                                        _ = v97
                                    }
                                }
                            }
                            _ = fi55
                            _ = v56
                            _ = v57
                            _ = v58
                        }
                    }
                }
                var v98 float32 = float32(3.14159)
                v99 := GenS2{F0: 255, F1: fi37}
                _ = fi37
                _ = v98
                _ = v99
            }
        } else {
            v100 := []int{32767, -78}
            var v101 uint8 = 128
            switch "\n\t" {
            case "":
                var v102 int16 = 1
                _ = v102
            case "a":
                if len(v100) >= 2 {
                    v100 = v100[1:len(v100)]
                }
            }
            _ = v100
            _ = v101
        }
    }
    var v103 uint = 65535
    fmt.Println(v103, v34)
    fmt.Println(v1, v3, v2)
    for fi104 := 0; fi104 < 16; fi104++ {
        var v105 bool = (140 <= 69)
        var v106 uint16 = 42
        _ = fi104
        _ = v105
        _ = v106
    }
    switch "\n\t" {
    case "\n\t":
        for fi107 := 0; fi107 < 10; fi107++ {
            v108 := []int{0, fi107, fi107}
            v109 := GenS10{F0: 18446744073709551615}
            _ = fi107
            _ = v108
            _ = v109
        }
    case "a":
        switch 243 {
        case 32768:
            switch "foo bar" {
            case "foo bar":
                switch 127 {
                case -9223372036854775808:
                    var v110 uint32 = 1
                    _ = v110
                case 2147483647:
                    v3 += (v34 * v3)
                }
            default:
                v111 := []int{}
                _ = v111
            }
        case 9223372036854775807:
            v1 = -2147483648
        case -2147483648:
            var v112 string = ("\n\t" + "")
            _ = v112
        }
    case "foo bar":
        fmt.Println(v2, v35)
    }
    if false {
        if false {
            switch "a" {
            case "":
                var v113 uint8 = 128
                _ = v113
            case "\n\t":
                var v114 uint8 = 0
                _ = v114
            case "foo bar":
                for fi115 := 0; fi115 < 2; fi115++ {
                    fmt.Println(v1, v34)
                    _ = fi115
                }
            }
        }
        var v116 uint8 = 18
        var v117 int8 = -64
        _ = v116
        _ = v117
    } else {
        fmt.Println(v3)
        if (4294967295 < 283) {
            v118 := GenS11{F0: 0, F1: 32767, F2: 2}
            _ = v118
        } else {
            v3 = genFn1(v2)
            fmt.Println(v103, v3)
        }
        switch "" {
        case "":
            for fi119 := 0; fi119 < 9; fi119++ {
                var v120 int16 = genFn0(true, 85, v1)
                v103 -= v103
                _ = fi119
                _ = v120
            }
        case "foo bar":
            v121 := []int{662}
            _ = v121
        case "\n\t":
            for fi122 := 0; fi122 < 11; fi122++ {
                if true {
                    v123 := []int{32767, (-(fi122)), fi122, (fi122 + fi122)}
                    var v124 uint = 2
                    _ = v123
                    _ = v124
                } else {
                    v125 := GenS12{F0: (false && (!true)), F1: 33, F2: (42 > 629)}
                    _ = v125
                }
                _ = fi122
            }
        default:
            for fi126 := 0; fi126 < 2; fi126++ {
                v127 := GenS6{F0: (1 < 482), F1: -32768, F2: 32767}
                _ = fi126
                _ = v127
            }
        }
    }
    v128 := []int{-960, 32767}
    if (!true) {
        if ((v2 | v2) < (v2 * 2)) {
            if len(v128) > 0 {
                v128[0] = -770
            }
        } else {
            var v129 uint16 = 1
            _ = v129
        }
        v2 += 2
        v35 += v35
    }
    if len(v128) >= 2 {
        v128 = v128[1:len(v128)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("slicelen v128:", len(v128), v128)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v103
    _ = v128
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
