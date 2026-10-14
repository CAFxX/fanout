package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint }
type GenS3 struct { F0 bool; F1 uint16; F2 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS6 struct { F0 uint; F1 uint16; F2 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint16; F1 uint16 }
type GenS10 struct { F0 int16; F1 string }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int64; F1 int8; F2 int64 }
type GenS13 struct { F0 int; F1 int64; F2 int16 }
type GenS14 struct { F0 uint16 }
func genFn0(p0 int64) uint {
    v5 := GenS1{F0: 42}
    p0 += p0
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint = 437
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    p0 = (^(p0))
    _ = p0
    _ = v5
    return 272
}
func genFn1(p0 int64, p1 int32) uint16 {
    v8 := []int{608, 0, 444}
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := GenS2{F0: p1, F1: 246, F2: 2}
        var v11 uint16 = 65535
        if false {
            var v12 uint8 = 1
            for fi13 := 0; fi13 < 15; fi13++ {
                if false {
                    if len(v8) > 0 {
                        v8[0] = (fi9 % ((127) | 1))
                    }
                    v14 := GenS1{F0: (p1 % ((int32(v11)) | 1))}
                    v12 = v12
                    _ = v14
                }
                _ = fi13
            }
            _ = len(v8)
            _ = v12
        } else {
            p0 += p0
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    if ("hello" <= "a") {
        var v15 uint64 = 42
        var v16 uint16 = 540
        switch -947 {
        case -446:
            v8 = append(v8, 128)
        case 780:
            _ = len(v8)
        }
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v8
    return 42
}

func main() {
    var v1 int8 = -128
    var v2 int64 = -1
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    var v17 uint16 = 42
    for fi18 := 0; fi18 < 5; fi18++ {
        v17 = (v17 | v17)
        var v19 int16 = 42
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 int64 = (v2 & 42)
            if (v1 < (v1 / ((42) | 1))) {
                var v22 int32 = -1
                _ = v22
            } else {
                _ = len(v3)
                var v23 uint64 = 4294967295
                _ = v23
            }
            v21 -= -751
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v17)
    fmt.Println(v1, v1)
    fmt.Println(v2)
    for fi24 := 0; fi24 < 18; fi24++ {
        var v25 string = "\n\t"
        switch v25 {
        case "a":
            fmt.Println(v17, v17, fi24)
        case "\n\t":
            v1 -= (v1 ^ 69)
        }
        var v26 uint32 = 939
        _ = fi24
        _ = v25
        _ = v26
    }
    v17 = (v17 % (((v17 * 170)) | 1))
    var v27 uint8 = 2
    v28 := GenS2{F0: -2147483648, F1: 42, F2: 42}
    v29 := GenS1{F0: -907}
    v30 := []int{-9223372036854775808, 956, 184}
    if false {
        v27 = v27
        switch -1 {
        case 128:
            v31 := GenS3{F0: false, F1: (v17 << ((4294967295) & 63)), F2: 42}
            _ = v31
        case 429:
            switch 9223372036854775807 {
            case 128:
                v2 -= -9223372036854775808
            default:
                var v32 float64 = (-2.25 + (1.5 / 3.14159))
                _ = v32
            }
        case -908:
            var v33 float32 = 1e10
            _ = v33
        default:
            if true {
                v17 = (^(v17))
                for fi34 := 0; fi34 < 14; fi34++ {
                    switch "\n\t" {
                    case "":
                        if false {
                            var v35 bool = (!(-2.25 == float64(1.5)))
                            _ = v35
                        }
                    }
                    v27 += (v27 << ((255) & 63))
                    _ = fi34
                }
                for fi36 := 0; fi36 < 3; fi36++ {
                    var v37 float64 = 1e10
                    v1 += (v1 >> ((genFn0(v2)) & 63))
                    _ = fi36
                    _ = v37
                }
            }
        }
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        for fi39 := 0; fi39 < 17; fi39++ {
            if ((v2 - (v2 | (v2 & 1))) <= (v2 - v2)) {
                var v40 int64 = v2
                var v41 float64 = 3.14159
                if (false && (v1 < (v1 - -128))) {
                    if ((false || true) || true) {
                        if (438 > v17) {
                            v42 := GenS1{F0: -2147483648}
                            switch ((("foo bar" + "\n\t") + "\n\t") + "\n\t") {
                            case "a":
                                v43 := GenS2{F0: 598, F1: 0, F2: 1}
                                _ = v43
                            case "foo bar":
                                if false {
                                    if (false || (false && false)) {
                                        v44 := GenS1{F0: 42}
                                        v45 := GenS1{F0: 898}
                                        if (!(!(float32(v41) > 3.14159))) {
                                            v46 := GenS4{F0: uint32(v27)}
                                            if len(v30) > 0 {
                                                v30[0] = fi38
                                            }
                                            var v47 uint16 = (v17 & genFn1(-9223372036854775808, 1))
                                            v48 := GenS3{F0: false, F1: 32767, F2: 0}
                                            v49 := GenS5{F0: genFn1((v2 ^ (v40 + v2)), 2147483647), F1: 242, F2: 1}
                                            v40 = v2
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                        } else {
                                            if true {
                                                var v50 int16 = -1
                                                fmt.Println(v17, v2, v17)
                                                v40 = v40
                                                _ = v50
                                            }
                                            v51 := GenS4{F0: 0}
                                            _ = v51
                                        }
                                        _ = v44
                                        _ = v45
                                    } else {
                                        fmt.Println(v1)
                                        v52 := GenS4{F0: 4294967295}
                                        _ = v52
                                    }
                                    v53 := GenS2{F0: 78, F1: 168, F2: 65535}
                                    _ = v53
                                } else {
                                    var v54 bool = false
                                    _ = v54
                                }
                                switch -493 {
                                case 32768:
                                    var v55 int8 = (v1 - (v1 | (v1 * (v1 % (((v1 >> ((538) & 63))) | 1)))))
                                    _ = v55
                                default:
                                    switch (fi38 & (fi39 | -809)) {
                                    case -686:
                                        v2 = v40
                                    case 32767:
                                        v41 -= v41
                                    case -9223372036854775808:
                                        var v56 int32 = 2147483647
                                        _ = v56
                                    }
                                }
                                var v57 float32 = (-2.25 - -2.25)
                                switch "" {
                                case "":
                                    v58 := GenS4{F0: 0}
                                    _ = v58
                                default:
                                    if (v17 <= genFn1((^(v40)), 1)) {
                                        fmt.Println(v27)
                                        if (257 < (v2 % ((-761) | 1))) {
                                            var v59 int = fi38
                                            switch "a" {
                                            case "\n\t":
                                                if false {
                                                    if false {
                                                        var v60 int16 = 42
                                                        var v61 uint8 = v27
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                    switch ("\n\t" + "hello") {
                                                    case "foo bar":
                                                        v41 = v41
                                                    default:
                                                        v62 := GenS1{F0: 1}
                                                        _ = v62
                                                    }
                                                } else {
                                                    v41 -= v41
                                                    switch ("hello" + "a") {
                                                    case "a":
                                                        if true {
                                                            v40 -= (v2 << ((42) & 63))
                                                        } else {
                                                            var v63 uint32 = 227
                                                            v59 += fi38
                                                            v59 += int(v1)
                                                            v64 := v3[(v63 >> ((11) & 63))]
                                                            v65 := GenS2{F0: 549, F1: 2, F2: 0}
                                                            v2 += v2
                                                            v66 := GenS6{F0: 981, F1: v17, F2: v1}
                                                            var v67 uint16 = 65535
                                                            _ = v63
                                                            _ = v64
                                                            _ = v65
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        if (((-2.25 * v57) / v57) >= (0.1 * v57)) {
                                                            _ = len(v30)
                                                            v30 = append(v30, (fi39 * (v59 % ((-625) | 1))))
                                                        } else {
                                                            v59 += v59
                                                        }
                                                        v1 -= v1
                                                        fmt.Println(v40)
                                                        v57 = float32(v41)
                                                    case "":
                                                        v4["foo bar"] = v59
                                                    default:
                                                        var v68 uint16 = 1
                                                        _ = v68
                                                    }
                                                    var v69 uint32 = 42
                                                    _ = v69
                                                }
                                            case "":
                                                var v70 int64 = 608
                                                _ = v70
                                            default:
                                                v71 := GenS5{F0: genFn1((v2 | -1), 42), F1: 877, F2: 880}
                                                _ = v71
                                            }
                                            fmt.Println(fi38)
                                            _ = v59
                                        } else {
                                            if (true || true) {
                                                var v72 int16 = -1
                                                fmt.Println(v41, v40)
                                                var v73 uint8 = (v27 * 127)
                                                _ = v72
                                                _ = v73
                                            } else {
                                                var v74 uint16 = 1
                                                _ = v74
                                            }
                                            v2 = v2
                                            v75 := GenS2{F0: -1, F1: 498, F2: 558}
                                            _ = v75
                                        }
                                        if ((32767 == int16(v40)) || false) {
                                            v17 -= genFn1(v40, 985)
                                            switch (("foo bar" + ("hello" + ("" + "foo bar"))) + "") {
                                            case "hello":
                                                fmt.Println(v41)
                                            case "\n\t":
                                                if true {
                                                    var v76 float64 = v41
                                                    v1 -= 57
                                                    _ = v76
                                                }
                                            case "a":
                                                v77 := GenS6{F0: genFn0(v40), F1: v17, F2: v1}
                                                _ = v77
                                            }
                                        } else {
                                            v41 = (1e10 - v41)
                                            v2 = 1
                                            v17 = v17
                                        }
                                        v78 := GenS4{F0: 2147483647}
                                        fmt.Println(v41, v2)
                                        _ = v78
                                    } else {
                                        switch fi39 {
                                        case 244:
                                            fmt.Println(fi39, v2, fi39)
                                        }
                                    }
                                }
                                var v79 int8 = (v1 & (v1 << ((290) & 63)))
                                v2 = (-(v40))
                                switch 0 {
                                case 42:
                                    var v80 int64 = v2
                                    _ = v80
                                case 32768:
                                    switch (fi39 | (-(fi39))) {
                                    case -1:
                                        var v81 int8 = (v79 % ((100) | 1))
                                        _ = v81
                                    default:
                                        if true {
                                            fmt.Println(v40, v27)
                                            delete(v4, "a")
                                        } else {
                                            v40 -= 406
                                        }
                                        var v82 uint64 = 42
                                        delete(v4, "\n\t")
                                        v17 = genFn1(v40, 356)
                                        var v83 uint16 = 622
                                        _ = v82
                                        _ = v83
                                    }
                                case -758:
                                    delete(v4, "")
                                }
                                _ = v57
                                _ = v79
                            }
                            v84 := GenS4{F0: 2}
                            _ = v42
                            _ = v84
                        }
                        var v85 uint64 = 178
                        v86 := GenS4{F0: 0}
                        _ = v85
                        _ = v86
                    }
                    v87 := GenS7{F0: (v1 * v1)}
                    switch (^(fi39)) {
                    case 127:
                        switch ("" + ("\n\t" + "a")) {
                        case "a":
                            if false {
                                v88 := GenS8{F0: 52}
                                var v89 uint32 = 42
                                if false {
                                    v30 = append(v30, fi38)
                                    v2 -= -1
                                } else {
                                    v90 := GenS5{F0: 255, F1: 42, F2: 18446744073709551615}
                                    fmt.Println(v27, v17)
                                    _ = v90
                                }
                                _ = v88
                                _ = v89
                            }
                        case "hello":
                            v91 := GenS9{F0: 2, F1: (v17 << ((800) & 63))}
                            var v92 int = 665
                            fmt.Println(v92)
                            _ = v91
                            _ = v92
                        case "":
                            v93 := GenS4{F0: 4294967295}
                            _ = v93
                        default:
                            v40 += (v40 * (v40 / (((v40 - (v2 * v2))) | 1)))
                            v30 = append(v30, 127)
                            fmt.Println(v17, fi38)
                            var v94 int32 = 0
                            var v95 uint8 = (^(v27))
                            v96 := GenS10{F0: 42, F1: ("hello" + ("\n\t" + ""))}
                            var v97 uint8 = 127
                            v2 = v40
                            _ = v94
                            _ = v95
                            _ = v96
                            _ = v97
                        }
                    case -800:
                        if (4294967295 >= 255) {
                            if ((42 > -2147483648) && false) {
                                v98 := GenS11{F0: 9223372036854775807}
                                v41 = v41
                                var v99 float32 = 1e10
                                _ = v98
                                _ = v99
                            } else {
                                if (!(true && (42 > 18446744073709551615))) {
                                    if (false && (1 == (^(v27)))) {
                                        var v100 int8 = v1
                                        fmt.Println(v40, v17)
                                        fmt.Println(v40, v1)
                                        v2 += -1
                                        _ = v100
                                    }
                                    if ((^(fi39)) <= fi38) {
                                        v17 = v17
                                        if false {
                                            v101 := GenS12{F0: -1, F1: (v1 ^ v1), F2: v2}
                                            _ = v101
                                        } else {
                                            var v102 int16 = 248
                                            _ = v102
                                        }
                                    } else {
                                        if false {
                                            switch fi39 {
                                            case -2147483648:
                                                v103 := GenS8{F0: 2147483647}
                                                _ = v103
                                            case -172:
                                                var v104 float64 = v41
                                                _ = v104
                                            case 9223372036854775807:
                                                if false {
                                                    v30 = append(v30, int(v1))
                                                    v105 := GenS11{F0: 9223372036854775807}
                                                    var v106 uint16 = (v17 | (v17 >> ((203) & 63)))
                                                    _ = v105
                                                    _ = v106
                                                }
                                            }
                                            v40 = (v40 >> ((812) & 63))
                                        } else {
                                            switch 32767 {
                                            case 128:
                                                _ = len(v4)
                                            case -9223372036854775808:
                                                if len(v30) > 0 {
                                                    v30[0] = -774
                                                }
                                            }
                                            var v107 uint8 = (v27 | (v27 ^ 205))
                                            var v108 uint64 = 1
                                            v109 := GenS3{F0: true, F1: v17, F2: 2147483647}
                                            _ = v107
                                            _ = v108
                                            _ = v109
                                        }
                                        v41 += float64(0.0)
                                        v110 := GenS13{F0: -291, F1: -1, F2: -54}
                                        _ = v110
                                    }
                                }
                                if false {
                                    var v111 float64 = v41
                                    v111 = 0.0
                                    _ = v111
                                } else {
                                    if len(v30) > 0 {
                                        v30[0] = fi39
                                    }
                                    v41 = v41
                                }
                            }
                        }
                    case 714:
                        switch fi39 {
                        case 128:
                            if (1.5 == ((float64(1.5) + v41) - v41)) {
                                var v112 float32 = 1.5
                                _ = v112
                            } else {
                                var v113 uint32 = 4294967295
                                var v114 uint64 = 2
                                var v115 uint64 = (^(v114))
                                _ = v113
                                _ = v114
                                _ = v115
                            }
                        case -507:
                            if (true || true) {
                                fmt.Println(v1, v41, v40)
                            }
                        case -451:
                            v116 := GenS14{F0: v17}
                            _ = v116
                        default:
                            v2 = v2
                        }
                    default:
                        v40 = v2
                    }
                    v3[2147483647] = ((v1 / ((38) | 1)) < (v1 * (-(v1))))
                    v27 = v27
                    _ = v87
                }
                _ = v40
                _ = v41
            }
            _ = fi39
        }
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
