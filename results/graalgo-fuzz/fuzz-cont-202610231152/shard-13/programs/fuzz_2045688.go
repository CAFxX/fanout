package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int64 }
type GenS2 struct { F0 int32; F1 int32; F2 int16 }
type GenS3 struct { F0 uint8; F1 int8; F2 int32 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 int32; F1 int; F2 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 string; F1 uint16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 bool; F1 int32 }
type GenS12 struct { F0 uint64; F1 int8; F2 int64 }
type GenS13 struct { F0 uint; F1 int64; F2 int32 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 uint8; F1 uint; F2 string }
type GenS17 struct { F0 uint }
func genFn0(p0 uint32) uint32 {
    var v7 int64 = 0
    if (128 == 127) {
        v8 := []int{0, 897, 2147483647}
        for fi9 := 0; fi9 < 15; fi9++ {
            p0 = p0
            v7 -= v7
            _ = fi9
        }
        _ = v8
    } else {
        for fi10 := 0; fi10 < 8; fi10++ {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v7)
            case "a":
                p0 += (p0 + (p0 + p0))
            case "hello":
                v11 := []int{42, fi10, -2147483648, 239}
                _ = v11
            default:
                v12 := GenS1{F0: 9223372036854775807, F1: 6, F2: (v7 / ((v7) | 1))}
                _ = v12
            }
            switch "a" {
            case "foo bar":
                var v13 int = (fi10 | -211)
                _ = v13
            case "\n\t":
                v14 := []int{fi10, fi10, fi10}
                _ = v14
            default:
                p0 = p0
            }
            p0 += p0
            _ = fi10
        }
        if true {
            p0 -= p0
            fmt.Println(p0, v7)
            p0 -= 712
        }
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := []int{fi15, (-(fi15)), (fi15 ^ 4294967295), -783}
            _ = len(v16)
            _ = fi15
            _ = v16
        }
    }
    v17 := GenS2{F0: 2147483647, F1: -1, F2: -626}
    _ = p0
    _ = v17
    _ = v7
    return 716
}

func main() {
    v1 := []int{1, 4294967295}
    var v2 int = -1
    var v3 int8 = 42
    var v4 uint = 4294967295
    v5 := []uint32{0}
    var v6 int = v2
    var v18 bool = false
    var v19 uint = (v4 | 467)
    v19 -= 38
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 float64 = 0.0
        v18 = (1 > 1)
        for fi22 := 0; fi22 < 14; fi22++ {
            fmt.Println(v21, v21)
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    if v18 {
        if (!(v3 != (^(v3)))) {
            if (genFn0(2) > 22) {
                fmt.Println(v2, v18, v19)
            } else {
                v4 += v19
                for fi23 := 0; fi23 < 19; fi23++ {
                    v5 = append(v5, 607)
                    _ = fi23
                }
            }
            if v18 {
                v24 := GenS3{F0: 42, F1: v3, F2: 0}
                var v25 int32 = 773
                for fi26 := 0; fi26 < 16; fi26++ {
                    if len(v5) > 0 {
                        v5[0] = 563
                    }
                    v4 = v19
                    if (!v18) {
                        switch fi26 {
                        case -9223372036854775808:
                            var v27 float64 = -2.25
                            _ = v27
                        case 680:
                            fmt.Println(v3, v4)
                        case 128:
                            var v28 bool = (1 <= 32767)
                            _ = v28
                        default:
                            if v18 {
                                v2 = fi26
                            } else {
                                switch 4294967295 {
                                case -164:
                                    v25 -= v25
                                }
                            }
                        }
                        switch (fi26 / ((v2) | 1)) {
                        case -633:
                            fmt.Println(v3)
                        case -9223372036854775808:
                            for fi29 := 0; fi29 < 5; fi29++ {
                                v30 := GenS1{F0: -612, F1: -32768, F2: 1}
                                v31 := GenS4{F0: 1, F1: 220}
                                switch "a" {
                                case "hello":
                                    v32 := GenS1{F0: -1, F1: -1, F2: -1}
                                    _ = v32
                                case "\n\t":
                                    var v33 int16 = 42
                                    _ = v33
                                case "":
                                    var v34 int8 = (v3 / (((^(v3))) | 1))
                                    _ = v34
                                }
                                var v35 uint64 = 4294967295
                                _ = fi29
                                _ = v30
                                _ = v31
                                _ = v35
                            }
                        }
                        var v36 int8 = (v3 & v3)
                        _ = v36
                    }
                    _ = fi26
                }
                _ = v24
                _ = v25
            } else {
                switch ("\n\t" + "hello") {
                case "":
                    switch "hello" {
                    case "hello":
                        var v37 int32 = -1
                        _ = v37
                    default:
                        if (((float64((1.5 + 0.1)) + 1e10) + (-2.25 * (-2.25 + 1.5))) < (((1.5 + 0.1) * 3.14159) * (1e10 / 1.0))) {
                            v38 := GenS5{F0: -2147483648, F1: v2, F2: v6}
                            var v39 string = "foo bar"
                            _ = v38
                            _ = v39
                        } else {
                            switch "foo bar" {
                            case "a":
                                if (4294967295 >= 759) {
                                    v40 := GenS6{F0: false}
                                    _ = v40
                                } else {
                                    var v41 string = "\n\t"
                                    v1 = append(v1, -9223372036854775808)
                                    _ = v41
                                }
                            case "hello":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case "\n\t":
                                for fi42 := 0; fi42 < 8; fi42++ {
                                    if v18 {
                                        var v43 int64 = -1
                                        var v44 int64 = v43
                                        _ = v43
                                        _ = v44
                                    } else {
                                        if (v18 && (!(!(139 > -9223372036854775808)))) {
                                            v6 -= v6
                                            switch "" {
                                            case "hello":
                                                switch 258 {
                                                case 32768:
                                                    fmt.Println(v2, v4, fi42)
                                                case -1:
                                                    if (-9223372036854775808 == v6) {
                                                        if ((!false) && v18) {
                                                            if ((0 == -290) || (!false)) {
                                                                var v45 int16 = 1
                                                                var v46 uint64 = 223
                                                                _ = v45
                                                                _ = v46
                                                            } else {
                                                                v47 := GenS1{F0: int64(v19), F1: -982, F2: 42}
                                                                _ = v47
                                                            }
                                                            fmt.Println(v4, v6)
                                                            switch "a" {
                                                            case "hello":
                                                                var v48 int8 = 1
                                                                _ = v48
                                                            case "":
                                                                v19 -= (v19 >> (((v4 & v4)) & 63))
                                                            case "\n\t":
                                                                v4 -= (^(v19))
                                                            }
                                                        }
                                                        switch "a" {
                                                        case "":
                                                            if v18 {
                                                                var v49 int8 = 1
                                                                _ = v49
                                                            }
                                                        case "\n\t":
                                                            if (!v18) {
                                                                if (0.0 <= ((1.5 + 1e10) * float32(1.5))) {
                                                                    v6 += fi42
                                                                    v50 := GenS7{F0: (!false)}
                                                                    v51 := GenS1{F0: 150, F1: 32767, F2: 0}
                                                                    _ = v50
                                                                    _ = v51
                                                                }
                                                                fmt.Println(fi42, fi42, v2)
                                                                var v52 int16 = 29
                                                                _ = v52
                                                            } else {
                                                                for fi53 := 0; fi53 < 6; fi53++ {
                                                                    if v18 {
                                                                        fmt.Println(v6)
                                                                        switch ("" + ("hello" + "foo bar")) {
                                                                        case "":
                                                                            if (-811 != 1) {
                                                                                v18 = (308 != 2)
                                                                                var v54 int64 = -1
                                                                                _ = v54
                                                                            } else {
                                                                                v55 := GenS4{F0: 502, F1: genFn0(0)}
                                                                                _ = v55
                                                                            }
                                                                            if ((225 >= 127) && false) {
                                                                                var v56 bool = (v18 || (v2 <= (v6 | v6)))
                                                                                v57 := GenS8{F0: 1}
                                                                                _ = v56
                                                                                _ = v57
                                                                            } else {
                                                                                var v58 float32 = (0.0 / 0.1)
                                                                                v59 := GenS8{F0: 255}
                                                                                v18 = false
                                                                                v60 := GenS2{F0: 0, F1: 1, F2: -261}
                                                                                _ = v58
                                                                                _ = v59
                                                                                _ = v60
                                                                            }
                                                                        case "\n\t":
                                                                            switch 2147483647 {
                                                                            case 805:
                                                                                v2 -= fi42
                                                                            case -551:
                                                                                var v61 int16 = 42
                                                                                _ = v61
                                                                            default:
                                                                                if ("foo bar" == "a") {
                                                                                    if v18 {
                                                                                        v2 -= (fi42 | (^(fi42)))
                                                                                        v1 = append(v1, (fi53 | v6))
                                                                                        if (!true) {
                                                                                            v62 := GenS5{F0: -1, F1: (v6 * (fi42 + 32768)), F2: -550}
                                                                                            if (("a" + "hello") < ("hello" + "\n\t")) {
                                                                                                fmt.Println(fi42)
                                                                                                v3 = -1
                                                                                                v63 := GenS3{F0: 255, F1: (^(v3)), F2: -1}
                                                                                                _ = v63
                                                                                            } else {
                                                                                                var v64 float32 = (0.0 - 0.1)
                                                                                                var v65 uint = v4
                                                                                                var v66 uint = 18446744073709551615
                                                                                                var v67 bool = true
                                                                                                v5 = append(v5, 0)
                                                                                                _ = v64
                                                                                                _ = v65
                                                                                                _ = v66
                                                                                                _ = v67
                                                                                            }
                                                                                            var v68 uint64 = 2
                                                                                            _ = v62
                                                                                            _ = v68
                                                                                        }
                                                                                    }
                                                                                    if (v18 || (2 >= 229)) {
                                                                                        v6 = (v6 - (^(fi42)))
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            v4 = 0
                                                                                        }
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                    }
                                                                                    var v69 uint32 = 1
                                                                                    _ = v69
                                                                                } else {
                                                                                    v70 := GenS4{F0: 794, F1: genFn0(4294967295)}
                                                                                    switch "a" {
                                                                                    case "hello":
                                                                                        v18 = true
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = (fi42 / ((v6) | 1))
                                                                                        }
                                                                                        switch fi53 {
                                                                                        case 42:
                                                                                            fmt.Println(v18, v3)
                                                                                        case 321:
                                                                                            v2 = v2
                                                                                        case -699:
                                                                                            _ = len(v5)
                                                                                        }
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = 630
                                                                                        }
                                                                                    case "":
                                                                                        var v71 int16 = 42
                                                                                        _ = v71
                                                                                    }
                                                                                    v72 := GenS5{F0: -1, F1: fi53, F2: -1}
                                                                                    _ = v70
                                                                                    _ = v72
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v73 int = (^(fi42))
                                                                            _ = v73
                                                                        }
                                                                        var v74 int16 = -1
                                                                        _ = v74
                                                                    } else {
                                                                        v2 = fi42
                                                                        if v18 {
                                                                            var v75 int32 = 0
                                                                            _ = v75
                                                                        } else {
                                                                            if (!(1 != 797)) {
                                                                                if (("hello" <= "a") || (v18 || (1 != -729))) {
                                                                                    if v18 {
                                                                                        v76 := GenS2{F0: 0, F1: -368, F2: 1000}
                                                                                        _ = v76
                                                                                    } else {
                                                                                        _ = len(v5)
                                                                                    }
                                                                                    v77 := GenS1{F0: -236, F1: -987, F2: -690}
                                                                                    _ = v77
                                                                                } else {
                                                                                    v78 := GenS9{F0: ("\n\t" + ""), F1: 32767}
                                                                                    v6 += (fi42 >> ((v4) & 63))
                                                                                    v3 = (v3 >> ((v4) & 63))
                                                                                    _ = v78
                                                                                }
                                                                            }
                                                                        }
                                                                        if ((((!(uint32(v19) >= 2147483647)) && true) || v18) && (36 >= 42)) {
                                                                            var v79 int16 = -1
                                                                            fmt.Println(fi53, v4, v3)
                                                                            v80 := GenS10{F0: -466}
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                    }
                                                                    _ = len(v1)
                                                                    var v81 uint = (v19 * (v19 * (^(v4))))
                                                                    _ = fi53
                                                                    _ = v81
                                                                }
                                                                var v82 int = (v2 & (v6 & int(v3)))
                                                                _ = v82
                                                            }
                                                        case "a":
                                                            v83 := GenS10{F0: 42}
                                                            _ = v83
                                                        }
                                                    }
                                                default:
                                                    if v18 {
                                                        v18 = true
                                                        if false {
                                                            var v84 int64 = -331
                                                            var v85 uint32 = genFn0(695)
                                                            _ = v84
                                                            _ = v85
                                                        } else {
                                                            var v86 string = (("a" + "hello") + ("hello" + ("foo bar" + "hello")))
                                                            _ = v86
                                                        }
                                                        var v87 float32 = 3.14159
                                                        _ = v87
                                                    } else {
                                                        if v18 {
                                                            var v88 float32 = 1.5
                                                            for fi89 := 0; fi89 < 5; fi89++ {
                                                                v4 = 701
                                                                if v18 {
                                                                    var v90 int = fi89
                                                                    _ = len(v5)
                                                                    _ = v90
                                                                } else {
                                                                    switch fi42 {
                                                                    case -2147483648:
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            v91 := GenS4{F0: 65535, F1: 1}
                                                                            _ = v91
                                                                        case "hello":
                                                                            v18 = (v18 && (!(-110 != v3)))
                                                                        default:
                                                                            fmt.Println(v3, v2, fi89)
                                                                        }
                                                                    default:
                                                                        var v92 uint64 = 195
                                                                        _ = v92
                                                                    }
                                                                    v93 := GenS10{F0: -9223372036854775808}
                                                                    _ = v93
                                                                }
                                                                _ = fi89
                                                            }
                                                            _ = v88
                                                        }
                                                        v3 = v3
                                                        for fi94 := 0; fi94 < 16; fi94++ {
                                                            var v95 uint = 1
                                                            if v18 {
                                                                v3 = (v3 / (((v3 - 1)) | 1))
                                                                var v96 uint8 = 147
                                                                _ = v96
                                                            } else {
                                                                v3 -= -62
                                                                if len(v1) > 0 {
                                                                    v1[0] = -2147483648
                                                                }
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            }
                                                            if v18 {
                                                                v97 := GenS11{F0: ((42 == 961) || (!(230 == -9223372036854775808))), F1: -2147483648}
                                                                _ = v97
                                                            }
                                                            _ = fi94
                                                            _ = v95
                                                        }
                                                    }
                                                }
                                            case "a":
                                                v98 := GenS12{F0: 0, F1: (v3 & (v3 / ((v3) | 1))), F2: -1}
                                                _ = v98
                                            case "\n\t":
                                                fmt.Println(fi42, v2)
                                            }
                                            for fi99 := 0; fi99 < 8; fi99++ {
                                                fmt.Println(v6, v19)
                                                if ((!v18) && false) {
                                                    switch 217 {
                                                    case 32767:
                                                        v19 -= v19
                                                        var v100 uint16 = 32767
                                                        _ = v100
                                                    }
                                                    switch fi99 {
                                                    case 159:
                                                        switch "" {
                                                        case "":
                                                            switch "a" {
                                                            case "a":
                                                                if (254 == 127) {
                                                                    switch 734 {
                                                                    case 128:
                                                                        v101 := GenS13{F0: v19, F1: -273, F2: -340}
                                                                        _ = v101
                                                                    case -9223372036854775808:
                                                                        var v102 uint64 = 42
                                                                        _ = v102
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        v6 = (v2 / (((v2 | fi99)) | 1))
                                                                    default:
                                                                        switch 491 {
                                                                        case -125:
                                                                            var v103 uint8 = 1
                                                                            _ = v103
                                                                        case 32768:
                                                                            if true {
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                                var v104 uint32 = 738
                                                                                _ = v104
                                                                            }
                                                                        case 454:
                                                                            var v105 uint16 = 775
                                                                            if v18 {
                                                                                fmt.Println(v3)
                                                                                v5 = append(v5, 0)
                                                                            }
                                                                            var v106 string = "hello"
                                                                            _ = v105
                                                                            _ = v106
                                                                        default:
                                                                            fmt.Println(fi42, v18)
                                                                        }
                                                                    }
                                                                }
                                                            case "hello":
                                                                v3 = -128
                                                            case "foo bar":
                                                                switch ("\n\t" + ("\n\t" + (("a" + "") + "a"))) {
                                                                case "foo bar":
                                                                    v6 -= 9223372036854775807
                                                                case "\n\t":
                                                                    fmt.Println(fi99, fi99, v4)
                                                                case "a":
                                                                    var v107 int64 = -842
                                                                    _ = v107
                                                                }
                                                            default:
                                                                v108 := GenS8{F0: 2}
                                                                _ = v108
                                                            }
                                                        default:
                                                            v6 += (fi99 ^ 2147483647)
                                                        }
                                                    case 169:
                                                        var v109 int8 = v3
                                                        _ = v109
                                                    }
                                                    var v110 string = ("hello" + "a")
                                                    _ = v110
                                                }
                                                switch 100 {
                                                case 2147483647:
                                                    switch (fi42 << (((v19 / ((v19) | 1))) & 63)) {
                                                    case 32767:
                                                        switch "hello" {
                                                        case "foo bar":
                                                            v2 = -749
                                                        }
                                                    case 426:
                                                        var v111 float32 = (1.5 * 3.14159)
                                                        v6 -= (v2 - (v6 >> (((v4 | v19)) & 63)))
                                                        if false {
                                                            v4 = (v4 | (^(v19)))
                                                            var v112 bool = v18
                                                            _ = v112
                                                        } else {
                                                            var v113 uint32 = 2
                                                            _ = v113
                                                        }
                                                        v3 = (-(v3))
                                                        var v114 float64 = 1.5
                                                        var v115 int = v6
                                                        _ = v111
                                                        _ = v114
                                                        _ = v115
                                                    case 0:
                                                        v2 += v6
                                                    }
                                                case 9223372036854775807:
                                                    var v116 uint64 = 4294967295
                                                    _ = v116
                                                default:
                                                    var v117 int64 = -657
                                                    _ = v117
                                                }
                                                _ = fi99
                                            }
                                        }
                                        for fi118 := 0; fi118 < 2; fi118++ {
                                            _ = len(v5)
                                            v18 = true
                                            var v119 uint32 = genFn0(42)
                                            _ = fi118
                                            _ = v119
                                        }
                                    }
                                    _ = fi42
                                }
                            }
                            if (2147483647 > 42) {
                                if false {
                                    _ = len(v1)
                                } else {
                                    v2 = v6
                                    if v18 {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    } else {
                                        if (!false) {
                                            var v120 int64 = -32
                                            fmt.Println(v19)
                                            _ = v120
                                        } else {
                                            v19 = v4
                                            fmt.Println(v2, v3, v2)
                                            v121 := GenS5{F0: 1, F1: v2, F2: (v6 >> ((18446744073709551615) & 63))}
                                            _ = v121
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                }
                                v18 = ((1 > -1) || (v18 && (!v18)))
                            }
                            v18 = true
                        }
                    }
                case "foo bar":
                    _ = len(v5)
                case "hello":
                    var v122 uint64 = 0
                    _ = v122
                }
            }
            v3 = (v3 << (((^(v19))) & 63))
        } else {
            v123 := GenS14{F0: 42}
            _ = v123
        }
        var v124 int32 = -2147483648
        _ = len(v5)
        _ = v124
    } else {
        v18 = false
        if ((v4 + (v19 << ((845) & 63))) > v19) {
            if (-1 > 42) {
                switch "a" {
                case "\n\t":
                    if ("foo bar" != "foo bar") {
                        var v125 uint32 = 2147483648
                        v126 := GenS5{F0: 0, F1: (v2 / ((1) | 1)), F2: (v6 & (v6 + 42))}
                        var v127 int8 = (^(v3))
                        _ = v125
                        _ = v126
                        _ = v127
                    } else {
                        for fi128 := 0; fi128 < 12; fi128++ {
                            v129 := GenS15{F0: 399}
                            _ = fi128
                            _ = v129
                        }
                    }
                case "":
                    v130 := GenS16{F0: 42, F1: 18446744073709551615, F2: "a"}
                    _ = v130
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            } else {
                switch v6 {
                case -1:
                    var v131 float32 = 1e10
                    _ = v131
                case 1:
                    var v132 float64 = 0.0
                    _ = v132
                case -100:
                    var v133 int64 = int64(v3)
                    _ = v133
                default:
                    var v134 int16 = 0
                    _ = v134
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            switch v2 {
            case 120:
                for fi135 := 0; fi135 < 10; fi135++ {
                    switch "" {
                    case "foo bar":
                        fmt.Println(v3)
                    case "a":
                        v2 = (fi135 & (fi135 - -9223372036854775808))
                    case "":
                        var v136 int32 = 796
                        _ = v136
                    default:
                        for fi137 := 0; fi137 < 13; fi137++ {
                            var v138 uint16 = 32767
                            _ = fi137
                            _ = v138
                        }
                    }
                    _ = fi135
                }
            case 128:
                if (!false) {
                    var v139 bool = true
                    var v140 int16 = 1000
                    _ = v139
                    _ = v140
                }
            case 312:
                v1 = append(v1, v6)
            default:
                v6 = (v2 + -272)
            }
        } else {
            var v141 int64 = 9223372036854775807
            switch 547 {
            case 128:
                switch int(v3) {
                case -1:
                    for fi142 := 0; fi142 < 6; fi142++ {
                        fmt.Println(v3, v3)
                        _ = fi142
                    }
                case 32767:
                    if len(v5) > 0 {
                        v5[0] = 42
                    }
                default:
                    v143 := GenS5{F0: 1, F1: v2, F2: 1}
                    _ = v143
                }
            case -9223372036854775808:
                if (0 <= 255) {
                    switch (v6 + v6) {
                    case -190:
                        var v144 int32 = -165
                        _ = v144
                    default:
                        var v145 uint16 = 1
                        _ = v145
                    }
                    v146 := GenS1{F0: v141, F1: 42, F2: -556}
                    _ = v146
                } else {
                    v2 += (v2 % (((v6 % ((v6) | 1))) | 1))
                    fmt.Println(v4)
                }
            default:
                v18 = ("hello" < "")
            }
            v147 := GenS1{F0: 0, F1: int16(v141), F2: (-(v141))}
            _ = v141
            _ = v147
        }
        v148 := GenS10{F0: 1}
        _ = v148
    }
    for fi149 := 0; fi149 < 20; fi149++ {
        for fi150 := 0; fi150 < 7; fi150++ {
            v151 := GenS5{F0: 0, F1: (fi149 - fi149), F2: -1}
            fmt.Println(v3, fi149, fi150)
            v152 := GenS7{F0: (2 != 0)}
            var v153 uint16 = 1
            _ = fi150
            _ = v151
            _ = v152
            _ = v153
        }
        v18 = (("" + "a") != "")
        var v154 uint8 = 128
        _ = fi149
        _ = v154
    }
    switch (-(v2)) {
    case -1:
        v5 = append(v5, 1)
    case -9223372036854775808:
        var v155 uint32 = 2
        _ = v155
    case 718:
        var v156 float32 = (0.0 + 0.0)
        _ = v156
    }
    switch (-(v6)) {
    case -680:
        if true {
            fmt.Println(v2, v6, v18)
        } else {
            if (!(-85 <= -1)) {
                if v18 {
                    var v157 float64 = (((1.5 - 0.0) * -2.25) / (-2.25 + ((((-2.25 - (float64((0.0 / 3.14159)) * (3.14159 - (0.0 * -2.25)))) - 0.0) * -2.25) - ((((1e10 - 3.14159) * float64((0.1 * 0.1))) / ((0.1 + float64(1.5)) - 0.1)) * (0.0 + 1.5)))))
                    fmt.Println(v3)
                    _ = v157
                }
                for fi158 := 0; fi158 < 2; fi158++ {
                    if v18 {
                        fmt.Println(fi158)
                    } else {
                        v159 := GenS17{F0: (v4 ^ v4)}
                        var v160 uint64 = 2
                        _ = v159
                        _ = v160
                    }
                    v4 += (v4 & v4)
                    _ = fi158
                }
                var v161 uint64 = 1
                _ = v161
            }
            fmt.Println(v19, v3)
        }
    }
    v162 := GenS6{F0: ("hello" <= ("foo bar" + (("hello" + ("foo bar" + "foo bar")) + "\n\t")))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v162
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
