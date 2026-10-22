package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 string; F1 string; F2 int64 }
type GenS10 struct { F0 int; F1 string }
type GenS11 struct { F0 int64; F1 int8; F2 int8 }
type GenS12 struct { F0 uint; F1 int16; F2 uint8 }
type GenS13 struct { F0 int16; F1 uint8; F2 int32 }
type GenS14 struct { F0 uint16; F1 int32; F2 string }
type GenS15 struct { F0 int32; F1 uint8 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 string }
func genFn0(p0 uint) uint32 {
    p0 = 264
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: fi7}
        var v9 int16 = 69
        var v10 int32 = 1
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: fi11}
        _ = fi11
        _ = v12
    }
    if (true && (true || (0 >= 435))) {
        fmt.Println(p0)
    }
    _ = p0
    return 109
}
func genFn1() uint {
    var v13 string = "a"
    var v14 uint32 = 0
    switch v13 {
    case "hello":
        fmt.Println(v14, v13)
    case "foo bar":
        for fi15 := 0; fi15 < 14; fi15++ {
            v16 := []int{(fi15 * 42), fi15, fi15, fi15}
            v17 := GenS1{F0: 128}
            switch (fi15 - -1) {
            case 128:
                switch "" {
                case "foo bar":
                    v18 := GenS1{F0: fi15}
                    _ = v18
                default:
                    _ = len(v16)
                }
            case 42:
                v14 = 242
            case 4294967295:
                fmt.Println(v13, fi15, fi15)
            }
            _ = fi15
            _ = v16
            _ = v17
        }
    default:
        if true {
            v19 := []int{-9223372036854775808, 742}
            switch v13 {
            case "foo bar":
                var v20 uint64 = 2
                _ = v20
            case "\n\t":
                var v21 float64 = ((1.5 * 1.5) * (0.0 / -2.25))
                _ = v21
            }
            v14 = v14
            _ = v19
        } else {
            for fi22 := 0; fi22 < 3; fi22++ {
                fmt.Println(v13, v14)
                var v23 uint16 = 986
                _ = fi22
                _ = v23
            }
        }
    }
    switch ("\n\t" + ("\n\t" + v13)) {
    case "":
        v24 := GenS1{F0: 773}
        _ = v24
    case "\n\t":
        v25 := []int{-2147483648, 9223372036854775807, -2147483648, 9223372036854775807}
        _ = v25
    default:
        var v26 uint8 = 128
        _ = v26
    }
    var v27 float32 = (0.1 - (-2.25 * (3.14159 / -2.25)))
    _ = v13
    _ = v14
    _ = v27
    return 65535
}

func main() {
    v1 := []string{}
    var v2 int8 = 42
    var v3 uint32 = 65535
    var v4 int32 = -2147483648
    v5 := []uint32{v3}
    var v6 float64 = 1.5
    v28 := GenS1{F0: 1}
    v6 = v6
    if (!true) {
        for fi29 := 0; fi29 < 13; fi29++ {
            switch fi29 {
            case -450:
                switch fi29 {
                case -1:
                    for fi30 := 0; fi30 < 17; fi30++ {
                        if (false && false) {
                            v2 -= (v2 ^ (-(v2)))
                        }
                        v2 = (-(v2))
                        var v31 float32 = (-2.25 * (3.14159 / (float32(v6) - 3.14159)))
                        _ = fi30
                        _ = v31
                    }
                case 127:
                    if (!((^(v3)) < (v3 & 2147483648))) {
                        for fi32 := 0; fi32 < 2; fi32++ {
                            _ = len(v1)
                            v33 := GenS2{F0: 1}
                            var v34 uint32 = 927
                            switch (fi32 / (((fi29 % (((^(fi32))) | 1))) | 1)) {
                            case -9223372036854775808:
                                v34 = v34
                            case -1:
                                if (!true) {
                                    if false {
                                        switch "a" {
                                        case "":
                                            fmt.Println(v6)
                                        case "a":
                                            switch fi29 {
                                            case 127:
                                                v35 := GenS2{F0: 42}
                                                var v36 int64 = -485
                                                _ = v35
                                                _ = v36
                                            default:
                                                v3 -= (v34 << ((71) & 63))
                                                switch fi29 {
                                                case -752:
                                                    switch -520 {
                                                    case 4294967295:
                                                        var v37 float32 = 3.14159
                                                        _ = v37
                                                    }
                                                }
                                                _ = len(v5)
                                                if false {
                                                    if (331 < 0) {
                                                        var v38 int8 = 100
                                                        _ = v38
                                                    } else {
                                                        v3 = (^(v3))
                                                        fmt.Println(v34)
                                                        v39 := GenS2{F0: 42}
                                                        _ = v39
                                                    }
                                                    var v40 uint = uint(v3)
                                                    _ = v40
                                                }
                                            }
                                        case "hello":
                                            if true {
                                                v2 += -39
                                                if false {
                                                    fmt.Println(fi29)
                                                    switch "" {
                                                    case "hello":
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    default:
                                                        if (false && true) {
                                                            switch -781 {
                                                            case 9223372036854775807:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            case 840:
                                                                if (!false) {
                                                                    v41 := GenS3{F0: 1, F1: (v34 | v34)}
                                                                    v6 += v6
                                                                    v3 = v3
                                                                    _ = v41
                                                                } else {
                                                                    v6 = float64(1e10)
                                                                    v4 -= v4
                                                                    switch (fi29 << ((genFn1()) & 63)) {
                                                                    case 80:
                                                                        v34 += v34
                                                                    }
                                                                }
                                                            case -1:
                                                                if true {
                                                                    if (42 == v4) {
                                                                        var v42 int16 = -744
                                                                        _ = v42
                                                                    } else {
                                                                        if false {
                                                                            v6 += v6
                                                                            fmt.Println(fi29, fi32, v4)
                                                                        } else {
                                                                            if (18446744073709551615 == 2) {
                                                                                v34 = (v3 ^ 1)
                                                                                fmt.Println(v34)
                                                                                fmt.Println(v34, v4, v2)
                                                                            }
                                                                            v1 = append(v1, ("hello" + ((("a" + ("" + "foo bar")) + "foo bar") + ("a" + ""))))
                                                                            var v43 int8 = -128
                                                                            var v44 int8 = 100
                                                                            switch 2147483647 {
                                                                            case -445:
                                                                                v45 := GenS2{F0: 65535}
                                                                                _ = v45
                                                                            case 128:
                                                                                var v46 uint8 = 0
                                                                                _ = v46
                                                                            case 0:
                                                                                if (true || true) {
                                                                                    v47 := GenS2{F0: 65535}
                                                                                    _ = len(v1)
                                                                                    _ = v47
                                                                                } else {
                                                                                    if true {
                                                                                        var v48 int16 = 42
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        v48 -= (v48 & 32767)
                                                                                        if false {
                                                                                            v49 := GenS4{F0: (v4 | v4), F1: (v4 << ((1) & 63))}
                                                                                            _ = v49
                                                                                        } else {
                                                                                            fmt.Println(v4)
                                                                                        }
                                                                                        _ = v48
                                                                                    } else {
                                                                                        var v50 int16 = -401
                                                                                        fmt.Println(v2, v3)
                                                                                        _ = v50
                                                                                    }
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                    var v51 bool = true
                                                                                    _ = v51
                                                                                }
                                                                            }
                                                                            fmt.Println(v2, v3, v6)
                                                                            _ = v43
                                                                            _ = v44
                                                                        }
                                                                        var v52 int64 = -9223372036854775808
                                                                        var v53 uint16 = 65535
                                                                        _ = v52
                                                                        _ = v53
                                                                    }
                                                                    switch ("foo bar" + "hello") {
                                                                    case "hello":
                                                                        var v54 uint8 = 128
                                                                        _ = v54
                                                                    case "foo bar":
                                                                        v55 := GenS3{F0: 444, F1: v3}
                                                                        _ = v55
                                                                    case "a":
                                                                        var v56 int16 = 939
                                                                        _ = v56
                                                                    default:
                                                                        fmt.Println(v34, fi29)
                                                                    }
                                                                    v57 := GenS2{F0: 42}
                                                                    _ = v57
                                                                }
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            }
                                                            v1 = append(v1, "")
                                                        } else {
                                                            var v58 float32 = 1e10
                                                            _ = v58
                                                        }
                                                    }
                                                    if false {
                                                        v59 := GenS4{F0: (^(v4)), F1: 2147483647}
                                                        _ = v59
                                                    } else {
                                                        v60 := GenS5{F0: 2, F1: -32768, F2: 255}
                                                        _ = v60
                                                    }
                                                }
                                                v34 = (v3 + v34)
                                            }
                                        default:
                                            if (-9223372036854775808 > 9223372036854775807) {
                                                v61 := GenS6{F0: 1, F1: 255}
                                                _ = v61
                                            }
                                            var v62 uint8 = 205
                                            _ = v62
                                        }
                                        if false {
                                            var v63 float32 = -2.25
                                            _ = v63
                                        } else {
                                            var v64 int8 = v2
                                            _ = v64
                                        }
                                        if true {
                                            v65 := GenS7{F0: 252}
                                            var v66 float64 = v6
                                            _ = v65
                                            _ = v66
                                        } else {
                                            var v67 int8 = 100
                                            var v68 int8 = (v2 % (((v67 / ((-99) | 1))) | 1))
                                            switch 248 {
                                            case 32767:
                                                var v69 int = 32767
                                                _ = v69
                                            case 32768:
                                                if (1 != 582) {
                                                    v5 = append(v5, (^(v34)))
                                                    var v70 uint64 = 0
                                                    _ = v70
                                                } else {
                                                    fmt.Println(v6)
                                                    var v71 int64 = 0
                                                    fmt.Println(v3, v71, v67)
                                                    _ = v71
                                                }
                                            case -2147483648:
                                                var v72 uint64 = 2
                                                _ = v72
                                            default:
                                                var v73 int16 = -1
                                                _ = v73
                                            }
                                            _ = v67
                                            _ = v68
                                        }
                                    } else {
                                        v4 = v4
                                        if false {
                                            switch 1 {
                                            case 42:
                                                v74 := GenS8{F0: fi32}
                                                _ = v74
                                            case 128:
                                                var v75 uint64 = 479
                                                _ = v75
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if true {
                                                var v76 int32 = (v4 + v4)
                                                var v77 float64 = v6
                                                if true {
                                                    if (!(v34 == 304)) {
                                                        if len(v5) > 0 {
                                                            v5[0] = v3
                                                        }
                                                        if (genFn0(2) >= (v34 * v3)) {
                                                            v78 := GenS6{F0: 65535, F1: 42}
                                                            _ = v78
                                                        } else {
                                                            var v79 uint = genFn1()
                                                            v3 -= (v34 % (((^(v3))) | 1))
                                                            _ = v79
                                                        }
                                                        if len(v5) > 0 {
                                                            v5[0] = genFn0(255)
                                                        }
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = ("a" + "\n\t")
                                                    }
                                                    if false {
                                                        if false {
                                                            var v80 uint32 = 387
                                                            v1 = append(v1, "\n\t")
                                                            _ = v80
                                                        } else {
                                                            v4 += v4
                                                            var v81 uint64 = 1
                                                            _ = v81
                                                        }
                                                        if (true && true) {
                                                            v3 += (v34 | v3)
                                                        }
                                                        var v82 uint32 = (^(v34))
                                                        if len(v5) > 0 {
                                                            v5[0] = (^(v3))
                                                        }
                                                        switch (("hello" + (("" + "hello") + ("hello" + ("\n\t" + "")))) + "hello") {
                                                        case "":
                                                            v4 = (v76 % (((v4 - -926)) | 1))
                                                        case "\n\t":
                                                            if ((v3 + 307) != (v82 | v3)) {
                                                                v83 := GenS3{F0: 255, F1: (v34 >> ((4294967295) & 63))}
                                                                v34 = v3
                                                                _ = v83
                                                            }
                                                        default:
                                                            _ = len(v5)
                                                        }
                                                        if len(v5) > 0 {
                                                            v5[0] = (v34 - 2147483647)
                                                        }
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        _ = v82
                                                    } else {
                                                        v84 := GenS1{F0: 0}
                                                        if (v2 == v2) {
                                                            var v85 int32 = v4
                                                            _ = v85
                                                        } else {
                                                            var v86 uint = 4294967295
                                                            var v87 int8 = v2
                                                            _ = v86
                                                            _ = v87
                                                        }
                                                        _ = v84
                                                    }
                                                } else {
                                                    fmt.Println(v3, v34, v77)
                                                }
                                                _ = v76
                                                _ = v77
                                            }
                                        } else {
                                            if true {
                                                switch fi32 {
                                                case 127:
                                                    fmt.Println(v6, v4)
                                                case 42:
                                                    fmt.Println(v4, v2, v3)
                                                case 9223372036854775807:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                default:
                                                    fmt.Println(fi29, fi29, v3)
                                                }
                                                v88 := GenS3{F0: 357, F1: (v3 ^ v34)}
                                                var v89 uint32 = genFn0(255)
                                                _ = v88
                                                _ = v89
                                            } else {
                                                switch (fi32 & 32768) {
                                                case 4294967295:
                                                    var v90 uint = 1
                                                    _ = v90
                                                case -929:
                                                    if (!(true || true)) {
                                                        v2 += (v2 | v2)
                                                        var v91 bool = (3.14159 > 3.14159)
                                                        _ = v91
                                                    }
                                                case 101:
                                                    _ = len(v5)
                                                default:
                                                    var v92 uint8 = 1
                                                    if (0 < 596) {
                                                        switch -1 {
                                                        case 477:
                                                            var v93 uint = 676
                                                            _ = v93
                                                        case 128:
                                                            var v94 uint = 42
                                                            _ = v94
                                                        default:
                                                            var v95 bool = (true || false)
                                                            _ = v95
                                                        }
                                                        v6 += (float64(((3.14159 - 1e10) * -2.25)) + v6)
                                                    }
                                                    v6 = v6
                                                    _ = v92
                                                }
                                            }
                                        }
                                        v2 = (-(v2))
                                    }
                                }
                            case 0:
                                var v96 uint32 = v3
                                _ = v96
                            default:
                                _ = len(v5)
                            }
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        if ((-(fi29)) > -653) {
                            v97 := GenS4{F0: -614, F1: (v4 / ((0) | 1))}
                            var v98 int32 = (v4 + v4)
                            _ = v97
                            _ = v98
                        } else {
                            var v99 float32 = -2.25
                            for fi100 := 0; fi100 < 11; fi100++ {
                                fmt.Println(v4, v99, v2)
                                v4 = v4
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = fi100
                            }
                            _ = v99
                        }
                        if true {
                            for fi101 := 0; fi101 < 10; fi101++ {
                                v6 -= v6
                                v4 += (v4 * (v4 >> ((139) & 63)))
                                _ = fi101
                            }
                        }
                    } else {
                        if true {
                            var v102 bool = ((!(0 < 42)) && true)
                            _ = v102
                        } else {
                            if true {
                                for fi103 := 0; fi103 < 9; fi103++ {
                                    if (1000 > 42) {
                                        if false {
                                            v104 := GenS4{F0: 0, F1: v4}
                                            _ = v104
                                        } else {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        v105 := GenS9{F0: "\n\t", F1: "\n\t", F2: int64(v3)}
                                        _ = v105
                                    }
                                    if (true && (-9223372036854775808 < 9223372036854775807)) {
                                        var v106 uint = 1
                                        _ = v106
                                    } else {
                                        if (32767 >= fi29) {
                                            _ = len(v5)
                                            if ("a" > (("hello" + "\n\t") + "foo bar")) {
                                                var v107 uint32 = v3
                                                var v108 int = (fi103 + fi103)
                                                _ = v107
                                                _ = v108
                                            } else {
                                                switch ("\n\t" + "") {
                                                case "":
                                                    if true {
                                                        v5 = append(v5, 4294967295)
                                                    }
                                                default:
                                                    v2 = v2
                                                }
                                            }
                                        } else {
                                            v109 := GenS10{F0: (fi103 + 0), F1: "\n\t"}
                                            if len(v5) > 0 {
                                                v5[0] = (^(v3))
                                            }
                                            v110 := GenS6{F0: 1, F1: 1}
                                            _ = v109
                                            _ = v110
                                        }
                                        fmt.Println(v2, fi103, fi103)
                                        v111 := GenS7{F0: 0}
                                        var v112 uint8 = 206
                                        _ = v111
                                        _ = v112
                                    }
                                    v4 = 0
                                    _ = fi103
                                }
                            } else {
                                v113 := GenS11{F0: 0, F1: (v2 ^ v2), F2: -51}
                                _ = v113
                            }
                        }
                        if false {
                            v4 = (^(v4))
                        }
                        if false {
                            var v114 uint16 = 42
                            _ = v114
                        }
                    }
                case -648:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            case 1:
                for fi115 := 0; fi115 < 16; fi115++ {
                    fmt.Println(v4)
                    _ = fi115
                }
            default:
                v6 -= v6
            }
            var v116 string = "\n\t"
            switch v116 {
            case "":
                fmt.Println(v6)
            case "hello":
                v6 = (0.1 / (v6 - float64((float32(v6) / -2.25))))
            case "\n\t":
                var v117 bool = true
                _ = v117
            default:
                if false {
                    if false {
                        var v118 uint32 = 4294967295
                        _ = v118
                    } else {
                        var v119 float64 = v6
                        _ = v119
                    }
                    if false {
                        for fi120 := 0; fi120 < 1; fi120++ {
                            v4 += (v4 / ((v4) | 1))
                            if (!(0 <= 114)) {
                                v1 = append(v1, v116)
                                v3 -= 0
                                fmt.Println(fi120, v4, v116)
                            } else {
                                fmt.Println(v6)
                            }
                            v116 = v116
                            _ = fi120
                        }
                        var v121 int32 = (v4 >> ((4294967295) & 63))
                        var v122 uint = 18446744073709551615
                        _ = v121
                        _ = v122
                    } else {
                        for fi123 := 0; fi123 < 12; fi123++ {
                            switch int(v4) {
                            case -9223372036854775808:
                                _ = len(v1)
                            case -659:
                                var v124 uint64 = 2
                                _ = v124
                            default:
                                fmt.Println(v4, fi123, v4)
                            }
                            _ = fi123
                        }
                        for fi125 := 0; fi125 < 9; fi125++ {
                            fmt.Println(fi125, v2)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi125
                        }
                    }
                    fmt.Println(fi29)
                } else {
                    if (4294967295 >= 462) {
                        v2 += 1
                        var v126 int32 = (v4 ^ v4)
                        v3 = 1
                        _ = v126
                    } else {
                        v127 := GenS12{F0: 2, F1: -843, F2: 1}
                        if false {
                            for fi128 := 0; fi128 < 17; fi128++ {
                                v3 = v3
                                _ = len(v5)
                                var v129 uint32 = 2
                                if (1 != 303) {
                                    if true {
                                        fmt.Println(v2, fi128)
                                        var v130 float64 = v6
                                        switch "" {
                                        case "hello":
                                            v3 -= (v3 | (v129 << ((genFn1()) & 63)))
                                        default:
                                            v131 := GenS13{F0: 42, F1: 255, F2: (^(v4))}
                                            _ = v131
                                        }
                                        _ = v130
                                    } else {
                                        v116 = v116
                                        var v132 string = v116
                                        _ = v132
                                    }
                                    switch fi128 {
                                    case 2147483647:
                                        v3 = 252
                                    }
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = fi128
                                _ = v129
                            }
                            _ = len(v5)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            v4 = v4
                            v2 -= v2
                            if len(v5) > 0 {
                                v5[0] = (^(v3))
                            }
                        }
                        _ = v127
                    }
                }
            }
            _ = fi29
            _ = v116
        }
    } else {
        fmt.Println(v2, v3, v6)
        v133 := GenS7{F0: 1}
        _ = v133
    }
    for fi134 := 0; fi134 < 11; fi134++ {
        v135 := GenS14{F0: 801, F1: -652, F2: ""}
        _ = fi134
        _ = v135
    }
    var v136 uint64 = 0
    for fi137 := 0; fi137 < 19; fi137++ {
        var v138 int8 = (v2 * 127)
        if (-321 >= fi137) {
            var v139 string = ""
            _ = v139
        } else {
            for fi140 := 0; fi140 < 17; fi140++ {
                if (v136 > (^(v136))) {
                    var v141 float32 = 1e10
                    var v142 float64 = -2.25
                    if false {
                        var v143 uint16 = 2
                        var v144 uint16 = 2
                        _ = v143
                        _ = v144
                    } else {
                        if (0 > 1) {
                            var v145 bool = (!(!true))
                            _ = v145
                        } else {
                            v142 -= v6
                            v136 = (v136 << ((876) & 63))
                            v3 += (^(v3))
                        }
                        v146 := GenS5{F0: (v136 - 4294967295), F1: int16(fi137), F2: 633}
                        v147 := GenS6{F0: 991, F1: 2}
                        _ = v146
                        _ = v147
                    }
                    v148 := GenS15{F0: v4, F1: 223}
                    _ = v141
                    _ = v142
                    _ = v148
                } else {
                    v149 := GenS16{F0: 0}
                    _ = v149
                }
                var v150 float32 = (3.14159 * 3.14159)
                var v151 uint = 0
                _ = fi140
                _ = v150
                _ = v151
            }
            var v152 uint8 = 127
            _ = v152
        }
        _ = fi137
        _ = v138
    }
    var v153 uint8 = 42
    v136 = v136
    switch 1 {
    case -2147483648:
        for fi154 := 0; fi154 < 18; fi154++ {
            var v155 uint32 = (v3 >> ((1) & 63))
            _ = fi154
            _ = v155
        }
    case 800:
        v4 -= -2147483648
    case 32768:
        fmt.Println(v136, v4, v153)
    }
    var v156 int32 = (v4 / ((int32(v3)) | 1))
    v157 := GenS17{F0: "a"}
    switch 32767 {
    case -41:
        v158 := GenS10{F0: 32767, F1: ""}
        _ = v158
    case 32768:
        fmt.Println(v156, v4)
    default:
        if true {
            fmt.Println(v6, v6, v153)
            if len(v1) > 0 {
                v1[0] = ""
            }
            switch "hello" {
            case "foo bar":
                v156 -= v156
            }
        }
    }
    v6 += (v6 * v6)
    var v159 uint32 = (v3 << ((18446744073709551615) & 63))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v136:", v136)
    fmt.Println("v153:", v153)
    fmt.Println("v156:", v156)
    fmt.Println("v159:", v159)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v136
    _ = v153
    _ = v156
    _ = v157
    _ = v159
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
