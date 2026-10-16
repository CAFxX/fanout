package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint; F1 uint64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS12 struct { F0 int16; F1 int8 }
type GenS13 struct { F0 uint64; F1 int8; F2 int32 }
type GenS14 struct { F0 int64; F1 int32 }
type GenS15 struct { F0 int8; F1 bool }
type GenS16 struct { F0 string; F1 uint16; F2 int16 }
type GenS17 struct { F0 uint64; F1 int16 }
type GenS18 struct { F0 int; F1 int32 }
func genFn0() uint {
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 int32 = -1
        v6 := []int{fi4, 9223372036854775807}
        switch (fi4 / (((fi4 ^ (fi4 >> ((1) & 63)))) | 1)) {
        case 2147483647:
            v5 = (v5 & v5)
        case 4294967295:
            for fi7 := 0; fi7 < 17; fi7++ {
                if false {
                    v6 = append(v6, (fi7 << ((968) & 63)))
                }
                if (!true) {
                    v5 = (v5 / (((v5 % ((v5) | 1))) | 1))
                    v5 = 1
                    if false {
                        var v8 uint64 = 53
                        v5 += 0
                        var v9 float32 = 0.1
                        _ = v8
                        _ = v9
                    } else {
                        v10 := GenS1{F0: -1, F1: 91, F2: 232}
                        if (!(169 <= 0)) {
                            if (!(!(3.14159 != (0.1 - float64(0.1))))) {
                                switch "" {
                                case "foo bar":
                                    var v11 int32 = 2147483647
                                    _ = v11
                                default:
                                    var v12 float64 = (1.5 * 0.0)
                                    _ = v12
                                }
                                v5 = -1
                            }
                        }
                        v5 = v5
                        if len(v6) > 0 {
                            v6[0] = -1
                        }
                        _ = v10
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                var v13 uint8 = 127
                _ = fi7
                _ = v13
            }
        default:
            v5 = v5
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    var v14 int8 = -128
    _ = v14
    return 65535
}
func genFn1(p0 int, p1 uint32, p2 string) int64 {
    p2 = ("a" + ((("a" + ("" + p2)) + (p2 + p2)) + ("foo bar" + p2)))
    switch 128 {
    case 4294967295:
        v15 := GenS1{F0: 0, F1: 255, F2: genFn0()}
        _ = v15
    case 128:
        if false {
            p0 = (p0 / ((p0) | 1))
            var v16 uint32 = 968
            var v17 float64 = 3.14159
            _ = v16
            _ = v17
        } else {
            v18 := GenS2{F0: 1000}
            v19 := []int{-1}
            v20 := GenS2{F0: -1}
            _ = v18
            _ = v19
            _ = v20
        }
    case 809:
        var v21 int32 = -1
        _ = v21
    }
    var v22 uint = 65535
    v23 := GenS3{F0: 740}
    for fi24 := 0; fi24 < 5; fi24++ {
        switch p2 {
        case "hello":
            switch p2 {
            case "a":
                if false {
                    var v25 float32 = 3.14159
                    v22 -= (v22 - v22)
                    _ = v25
                }
            case "\n\t":
                var v26 uint = v22
                _ = v26
            }
        default:
            v27 := GenS1{F0: 9223372036854775807, F1: 2, F2: (v22 / (((^(v22))) | 1))}
            _ = v27
        }
        p1 -= p1
        fmt.Println(v22)
        _ = fi24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    return -82
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 734
    var v3 float64 = float64(-2.25)
    v2 = 506
    v2 -= v2
    var v28 int16 = 1000
    fmt.Println(v28)
    for fi29 := 0; fi29 < 18; fi29++ {
        for fi30 := 0; fi30 < 18; fi30++ {
            var v31 uint64 = 0
            if (true || (true || (144 < 18446744073709551615))) {
                var v32 int = 4294967295
                if true {
                    var v33 int = fi29
                    _ = v33
                }
                _ = v32
            }
            if (true || true) {
                v34 := GenS4{F0: 99, F1: 0}
                _ = v34
            } else {
                if ((true && true) || (!true)) {
                    if (fi30 != int(v28)) {
                        var v35 int64 = 9223372036854775807
                        _ = v35
                    } else {
                        var v36 float32 = 0.0
                        v37 := []int{0}
                        v2 -= v2
                        if (0 != 2) {
                            v37 = append(v37, 32767)
                        } else {
                            v3 -= ((((float64(((v36 * v36) - (v36 - float32(v3)))) / v3) / ((float64(v36) * (1e10 * (float64(1e10) - v3))) - v3)) * float64(1.5)) / v3)
                        }
                        _ = len(v1)
                        v38 := GenS5{F0: "\n\t"}
                        switch ("foo bar" + ("hello" + "foo bar")) {
                        case "foo bar":
                            if false {
                                switch "" {
                                case "\n\t":
                                    v39 := GenS6{F0: v31}
                                    _ = v39
                                case "foo bar":
                                    v36 += 3.14159
                                case "hello":
                                    var v40 int32 = -2147483648
                                    _ = v40
                                default:
                                    if true {
                                        var v41 int32 = 2147483647
                                        switch "\n\t" {
                                        case "hello":
                                            if false {
                                                v36 = 0.0
                                                if (821 <= v31) {
                                                    if false {
                                                        var v42 bool = (!(v36 >= 0.1))
                                                        var v43 int = fi30
                                                        _ = v42
                                                        _ = v43
                                                    } else {
                                                        v2 = v2
                                                        if true {
                                                            if (302 >= 0) {
                                                                v28 += (v28 - 1000)
                                                            } else {
                                                                fmt.Println(v41, fi30)
                                                                v41 = (v41 & v41)
                                                                _ = len(v37)
                                                            }
                                                            if true {
                                                                v28 = v28
                                                                var v44 float32 = v36
                                                                if ("foo bar" < ("foo bar" + ("\n\t" + (("\n\t" + (((("foo bar" + "foo bar") + "a") + ("hello" + "")) + "\n\t")) + "")))) {
                                                                    v36 = -2.25
                                                                } else {
                                                                    v3 -= v3
                                                                    v37 = append(v37, fi29)
                                                                    if ((^(v2)) >= v2) {
                                                                        var v45 uint32 = v2
                                                                        var v46 int16 = v28
                                                                        _ = v45
                                                                        _ = v46
                                                                    } else {
                                                                        v47 := GenS7{F0: 0}
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            fmt.Println(v28, v2)
                                                                        case "\n\t":
                                                                            if false {
                                                                                var v48 string = "hello"
                                                                                var v49 int64 = 255
                                                                                if false {
                                                                                    v50 := GenS7{F0: 891}
                                                                                    delete(v1, (v2 >> ((genFn0()) & 63)))
                                                                                    _ = v50
                                                                                }
                                                                                _ = v48
                                                                                _ = v49
                                                                            }
                                                                        case "a":
                                                                            var v51 bool = (42 >= 65535)
                                                                            var v52 uint = genFn0()
                                                                            var v53 uint8 = 1
                                                                            _ = v51
                                                                            _ = v52
                                                                            _ = v53
                                                                        }
                                                                        switch ("a" + ("\n\t" + "hello")) {
                                                                        case "":
                                                                            v54 := v1[(^(v2))]
                                                                            _ = v54
                                                                        }
                                                                        _ = v47
                                                                    }
                                                                }
                                                                _ = v44
                                                            } else {
                                                                if (!true) {
                                                                    v2 = 65535
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        v1[2] = true
                                                                        if false {
                                                                            var v55 int32 = v41
                                                                            switch -1 {
                                                                            case 128:
                                                                                switch "foo bar" {
                                                                                case "a":
                                                                                    v2 = v2
                                                                                default:
                                                                                    v56 := v1[(v2 & v2)]
                                                                                    _ = v56
                                                                                }
                                                                            case 42:
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    v41 += 2147483647
                                                                                case "a":
                                                                                    if len(v37) > 0 {
                                                                                        v37[0] = fi30
                                                                                    }
                                                                                    v57 := GenS7{F0: 42}
                                                                                    if true {
                                                                                        fmt.Println(v28, v36)
                                                                                    } else {
                                                                                        var v58 int16 = 1
                                                                                        _ = len(v37)
                                                                                        switch "foo bar" {
                                                                                        case "hello":
                                                                                            fmt.Println(v31, fi30, v28)
                                                                                        case "":
                                                                                            var v59 uint64 = 1
                                                                                            _ = v59
                                                                                        case "a":
                                                                                            switch "foo bar" {
                                                                                            case "":
                                                                                                var v60 float64 = v3
                                                                                                _ = v60
                                                                                            case "a":
                                                                                                if (false || (v36 <= v36)) {
                                                                                                    v61 := GenS6{F0: (v31 << ((1) & 63))}
                                                                                                    fmt.Println(v2)
                                                                                                    switch fi30 {
                                                                                                    case 128:
                                                                                                        v62 := GenS8{F0: 1, F1: v31}
                                                                                                        _ = v62
                                                                                                    case -122:
                                                                                                        var v63 float32 = 0.0
                                                                                                        _ = v63
                                                                                                    }
                                                                                                    _ = v61
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            if false {
                                                                                                if ((v58 > -1) && (v3 >= v3)) {
                                                                                                    v3 = 1.5
                                                                                                }
                                                                                                if (true && true) {
                                                                                                    var v64 uint16 = 32767
                                                                                                    var v65 uint8 = 127
                                                                                                    var v66 int8 = -49
                                                                                                    _ = v64
                                                                                                    _ = v65
                                                                                                    _ = v66
                                                                                                } else {
                                                                                                    v67 := v1[(v2 | (v2 | 2147483647))]
                                                                                                    _ = v67
                                                                                                }
                                                                                                var v68 float64 = (float64(1.5) / (3.14159 - (v3 - 0.1)))
                                                                                                v1[(v2 + v2)] = true
                                                                                                if false {
                                                                                                    var v69 uint64 = v31
                                                                                                    v68 -= -2.25
                                                                                                    switch "\n\t" {
                                                                                                    case "foo bar":
                                                                                                        v68 -= 0.0
                                                                                                    }
                                                                                                    _ = v69
                                                                                                }
                                                                                                fmt.Println(fi29)
                                                                                                v58 += (v58 + v28)
                                                                                                _ = v68
                                                                                            }
                                                                                        }
                                                                                        var v70 string = ""
                                                                                        v71 := GenS9{F0: (^(v31))}
                                                                                        var v72 int8 = -77
                                                                                        _ = v58
                                                                                        _ = v70
                                                                                        _ = v71
                                                                                        _ = v72
                                                                                    }
                                                                                    _ = v57
                                                                                }
                                                                            case 9223372036854775807:
                                                                                v3 = float64(0.1)
                                                                            default:
                                                                                var v73 int32 = (v55 + v41)
                                                                                _ = v73
                                                                            }
                                                                            v74 := GenS6{F0: (v31 / (((v31 << ((65535) & 63))) | 1))}
                                                                            _ = v55
                                                                            _ = v74
                                                                        } else {
                                                                            if (v41 <= 42) {
                                                                                var v75 int8 = -128
                                                                                _ = v75
                                                                            } else {
                                                                                switch -2147483648 {
                                                                                case 823:
                                                                                    v76 := GenS5{F0: ("a" + (("hello" + ("foo bar" + ("a" + "\n\t"))) + "\n\t"))}
                                                                                    _ = v76
                                                                                default:
                                                                                    v41 = v41
                                                                                }
                                                                                var v77 int16 = (v28 % (((v28 % ((v28) | 1))) | 1))
                                                                                if true {
                                                                                    var v78 uint = 2
                                                                                    _ = v78
                                                                                }
                                                                                _ = v77
                                                                            }
                                                                            var v79 string = "\n\t"
                                                                            _ = v79
                                                                        }
                                                                    case "a":
                                                                        v36 += v36
                                                                    default:
                                                                        var v80 uint = genFn0()
                                                                        _ = v80
                                                                    }
                                                                    if (false && true) {
                                                                        _ = len(v37)
                                                                        v2 += v2
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            fmt.Println(v3, v31)
                                                                        }
                                                                    } else {
                                                                        if (true && false) {
                                                                            v81 := GenS6{F0: v31}
                                                                            _ = v81
                                                                        } else {
                                                                            v82 := GenS10{F0: (^(v28))}
                                                                            if true {
                                                                                if (!false) {
                                                                                    var v83 uint = uint(v31)
                                                                                    v1[802] = (((-21 <= 67) || true) && (v28 <= 741))
                                                                                    if true {
                                                                                        var v84 uint = (v83 >> ((2) & 63))
                                                                                        if len(v37) > 0 {
                                                                                            v37[0] = fi30
                                                                                        }
                                                                                        if ((fi30 ^ (fi29 + 669)) < (fi30 << ((v83) & 63))) {
                                                                                            switch "a" {
                                                                                            case "\n\t":
                                                                                                if true {
                                                                                                    switch (fi30 / ((fi29) | 1)) {
                                                                                                    case -9223372036854775808:
                                                                                                        v85 := GenS4{F0: 128, F1: (v28 & (-(v28)))}
                                                                                                        _ = v85
                                                                                                    default:
                                                                                                        delete(v1, 2)
                                                                                                    }
                                                                                                    var v86 uint = (v83 % (((v83 % ((574) | 1))) | 1))
                                                                                                    fmt.Println(fi30)
                                                                                                    _ = v86
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            var v87 int = fi29
                                                                                            v88 := GenS7{F0: 32767}
                                                                                            v3 -= (v3 + v3)
                                                                                            var v89 int16 = v28
                                                                                            v90 := GenS5{F0: "hello"}
                                                                                            fmt.Println(v84, v2)
                                                                                            var v91 int16 = v89
                                                                                            v31 = 900
                                                                                            _ = v87
                                                                                            _ = v88
                                                                                            _ = v89
                                                                                            _ = v90
                                                                                            _ = v91
                                                                                        }
                                                                                        _ = v84
                                                                                    } else {
                                                                                        switch "hello" {
                                                                                        case "hello":
                                                                                            var v92 int64 = 263
                                                                                            _ = v92
                                                                                        case "foo bar":
                                                                                            if false {
                                                                                                var v93 int8 = 1
                                                                                                if true {
                                                                                                    var v94 int = (fi30 % (((fi30 * (fi29 * fi30))) | 1))
                                                                                                    _ = v94
                                                                                                } else {
                                                                                                    v28 = 0
                                                                                                    var v95 int = (-(fi30))
                                                                                                    if len(v37) >= 2 {
                                                                                                        v37 = v37[1:len(v37)]
                                                                                                    }
                                                                                                    _ = v95
                                                                                                }
                                                                                                _ = v93
                                                                                            }
                                                                                        case "\n\t":
                                                                                            fmt.Println(v36)
                                                                                        }
                                                                                        var v96 uint64 = (^(v31))
                                                                                        var v97 bool = ((float64(v36) - float64((v36 * float32(((-2.25 - 1.5) + 1e10))))) == -2.25)
                                                                                        _ = v96
                                                                                        _ = v97
                                                                                    }
                                                                                    _ = v83
                                                                                }
                                                                                fmt.Println(fi30, fi30, v41)
                                                                            }
                                                                            _ = v82
                                                                        }
                                                                        v98 := GenS11{F0: 0, F1: 513, F2: 1}
                                                                        var v99 int32 = (v41 ^ 1)
                                                                        _ = v98
                                                                        _ = v99
                                                                    }
                                                                    v100 := GenS10{F0: v28}
                                                                    var v101 bool = (61 < 128)
                                                                    v102 := v1[4294967295]
                                                                    if len(v37) >= 2 {
                                                                        v37 = v37[1:len(v37)]
                                                                    }
                                                                    _ = v100
                                                                    _ = v101
                                                                    _ = v102
                                                                }
                                                            }
                                                        } else {
                                                            var v103 int32 = v41
                                                            var v104 int = (fi29 + -918)
                                                            if ((((v103 - -2147483648) >= v41) || ((v31 >> ((255) & 63)) >= v31)) && (v2 < (v2 % ((v2) | 1)))) {
                                                                var v105 uint16 = 0
                                                                v31 = uint64(v105)
                                                                v37 = append(v37, (fi30 & v104))
                                                                v103 -= (v103 << ((18446744073709551615) & 63))
                                                                v106 := GenS12{F0: (v28 % (((-(v28))) | 1)), F1: int8(v28)}
                                                                _ = v105
                                                                _ = v106
                                                            } else {
                                                                if true {
                                                                    var v107 bool = true
                                                                    if v107 {
                                                                        delete(v1, uint32(v31))
                                                                        if (!v107) {
                                                                            fmt.Println(v104, v103)
                                                                            var v108 float64 = v3
                                                                            v41 += 0
                                                                            _ = v108
                                                                        } else {
                                                                            if len(v37) >= 2 {
                                                                                v37 = v37[1:len(v37)]
                                                                            }
                                                                            v36 += v36
                                                                            fmt.Println(v41, v28, fi29)
                                                                            v109 := GenS1{F0: genFn1((fi29 * 9223372036854775807), (v2 << ((270) & 63)), "\n\t"), F1: 127, F2: 65535}
                                                                            var v110 uint16 = 631
                                                                            _ = v109
                                                                            _ = v110
                                                                        }
                                                                    } else {
                                                                        v111 := GenS13{F0: (v31 / ((1) | 1)), F1: int8(v2), F2: (v41 << ((699) & 63))}
                                                                        _ = v111
                                                                    }
                                                                    v112 := GenS13{F0: v31, F1: -25, F2: v103}
                                                                    v113 := GenS14{F0: genFn1(fi29, v2, ("hello" + ("" + ("foo bar" + ("" + "\n\t"))))), F1: (v41 & v41)}
                                                                    if v107 {
                                                                        if v107 {
                                                                            v41 = -2147483648
                                                                            v1[(v2 & (v2 >> ((255) & 63)))] = (-9223372036854775808 < 9223372036854775807)
                                                                        }
                                                                        v107 = (!(((109 > 1) && v107) || (1 < (fi29 & (fi30 | -9223372036854775808)))))
                                                                    } else {
                                                                        fmt.Println(v28, v104)
                                                                        var v114 uint64 = v31
                                                                        _ = v114
                                                                    }
                                                                    var v115 float64 = v3
                                                                    if len(v37) > 0 {
                                                                        v37[0] = (fi30 - fi30)
                                                                    }
                                                                    v116 := GenS15{F0: 0, F1: false}
                                                                    _ = v107
                                                                    _ = v112
                                                                    _ = v113
                                                                    _ = v115
                                                                    _ = v116
                                                                } else {
                                                                    switch 127 {
                                                                    case -1:
                                                                        v103 = 2147483647
                                                                    }
                                                                    if true {
                                                                        if len(v37) > 0 {
                                                                            v37[0] = (fi29 + 127)
                                                                        }
                                                                        if true {
                                                                            v1[v2] = false
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            _ = v103
                                                            _ = v104
                                                        }
                                                        if ((3.14159 < float32(3.14159)) && (-9223372036854775808 <= 42)) {
                                                            var v117 int16 = (v28 + 42)
                                                            var v118 uint64 = v31
                                                            _ = v117
                                                            _ = v118
                                                        } else {
                                                            fmt.Println(fi30, v31, fi30)
                                                            v119 := GenS16{F0: ("\n\t" + ("a" + "foo bar")), F1: 255, F2: (v28 / ((-1) | 1))}
                                                            _ = v119
                                                        }
                                                    }
                                                    switch 83 {
                                                    case -99:
                                                        v120 := GenS12{F0: (v28 >> ((4294967295) & 63)), F1: -84}
                                                        _ = v120
                                                    case 1:
                                                        var v121 int16 = v28
                                                        _ = v121
                                                    default:
                                                        v122 := GenS17{F0: 42, F1: (^(v28))}
                                                        _ = v122
                                                    }
                                                    if len(v37) >= 2 {
                                                        v37 = v37[1:len(v37)]
                                                    }
                                                } else {
                                                    v2 = (v2 / ((v2) | 1))
                                                    v3 = ((v3 * (v3 / v3)) / -2.25)
                                                    if true {
                                                        switch ("" + ("hello" + "\n\t")) {
                                                        case "a":
                                                            if ((v2 == 990) || true) {
                                                                var v123 int32 = (v41 ^ 0)
                                                                _ = v123
                                                            }
                                                        case "hello":
                                                            v3 = v3
                                                        case "":
                                                            fmt.Println(v41)
                                                        default:
                                                            if len(v37) >= 2 {
                                                                v37 = v37[1:len(v37)]
                                                            }
                                                        }
                                                    } else {
                                                        switch (fi29 / ((fi29) | 1)) {
                                                        case 128:
                                                            v124 := GenS10{F0: v28}
                                                            _ = v124
                                                        case 4294967295:
                                                            fmt.Println(v28, v3)
                                                        default:
                                                            if len(v37) >= 2 {
                                                                v37 = v37[1:len(v37)]
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                        case "foo bar":
                                            v125 := v1[2147483647]
                                            _ = v125
                                        case "a":
                                            var v126 int16 = (v28 / ((-1) | 1))
                                            _ = v126
                                        default:
                                            v3 -= v3
                                        }
                                        _ = v41
                                    }
                                }
                                v36 = ((float32(v3) + v36) / 1.0)
                            } else {
                                fmt.Println(fi30)
                                var v127 float64 = v3
                                _ = v127
                            }
                        case "hello":
                            delete(v1, (v2 % ((v2) | 1)))
                        default:
                            if len(v37) >= 2 {
                                v37 = v37[1:len(v37)]
                            }
                        }
                        v3 = v3
                        _ = len(v37)
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                }
                var v128 bool = false
                if false {
                    var v129 int8 = -112
                    fmt.Println(v129)
                    switch "a" {
                    case "a":
                        v129 -= v129
                    case "":
                        var v130 bool = (!(((1.5 + 0.0) + -2.25) <= -2.25))
                        _ = v130
                    }
                    _ = v129
                }
                _ = v128
            }
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    var v131 int = -772
    var v132 float32 = ((0.1 * 0.1) - ((1.5 * ((float32(v3) / (1.5 - 0.1)) / float32((float64((1.5 + 1e10)) * v3)))) / float32((v3 / v3))))
    v1[(v2 - (v2 << ((45) & 63)))] = false
    for fi133 := 0; fi133 < 12; fi133++ {
        v134 := GenS14{F0: 712, F1: 42}
        _ = fi133
        _ = v134
    }
    var v135 int = -102
    v136 := GenS18{F0: (v131 - 2147483647), F1: -510}
    if (!false) {
        v135 += v135
        for fi137 := 0; fi137 < 18; fi137++ {
            var v138 uint16 = 32767
            for fi139 := 0; fi139 < 17; fi139++ {
                var v140 float64 = 1e10
                v140 = v3
                _ = fi139
                _ = v140
            }
            _ = fi137
            _ = v138
        }
        for fi141 := 0; fi141 < 3; fi141++ {
            switch v131 {
            case -320:
                var v142 int32 = -871
                _ = v142
            case -1:
                var v143 float64 = (v3 + float64(float32((0.1 / v3))))
                _ = v143
            case -706:
                for fi144 := 0; fi144 < 13; fi144++ {
                    switch fi144 {
                    case 32767:
                        var v145 bool = false
                        _ = v145
                    }
                    v3 = (float64((float32(((0.0 - v3) - float64(v132))) - v132)) * (v3 + v3))
                    v146 := GenS14{F0: 271, F1: 1}
                    _ = fi144
                    _ = v146
                }
            default:
                var v147 uint64 = 0
                _ = v147
            }
            fmt.Println(v132)
            _ = fi141
        }
    }
    var v148 string = ""
    var v149 uint = 255
    for fi150 := 0; fi150 < 10; fi150++ {
        var v151 bool = (!(((!false) && true) || true))
        switch ((v148 + v148) + "foo bar") {
        case "\n\t":
            v148 = "foo bar"
        default:
            v152 := GenS15{F0: 42, F1: v151}
            _ = v152
        }
        for fi153 := 0; fi153 < 18; fi153++ {
            var v154 uint = (v149 - 4294967295)
            _ = fi153
            _ = v154
        }
        _ = fi150
        _ = v151
    }
    var v155 float32 = 3.14159
    var v156 int32 = -210
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v131:", v131)
    fmt.Println("v132:", v132)
    fmt.Println("v135:", v135)
    fmt.Println("v148:", v148)
    fmt.Println("v149:", v149)
    fmt.Println("v155:", v155)
    fmt.Println("v156:", v156)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v131
    _ = v132
    _ = v135
    _ = v136
    _ = v148
    _ = v149
    _ = v155
    _ = v156
    _ = v2
    _ = v28
    _ = v3
}
