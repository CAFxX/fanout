package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int16 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint32; F1 int32; F2 int16 }
type GenS6 struct { F0 int16; F1 int8; F2 int64 }
type GenS7 struct { F0 uint16; F1 uint64; F2 bool }
type GenS8 struct { F0 uint; F1 uint16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS12 struct { F0 uint16; F1 string; F2 string }
func genFn0(p0 uint, p1 int, p2 uint32) int64 {
    for fi4 := 0; fi4 < 2; fi4++ {
        switch "hello" {
        case "\n\t":
            if true {
                v5 := GenS1{F0: 24, F1: 18446744073709551615, F2: 42}
                p2 -= p2
                switch "" {
                case "a":
                    var v6 uint32 = p2
                    _ = v6
                case "":
                    if false {
                        p2 -= 2147483648
                        v7 := GenS1{F0: 1, F1: (p0 + (p0 % ((p0) | 1))), F2: 1}
                        _ = v7
                    } else {
                        p0 = 0
                        p0 += p0
                        v8 := []int{}
                        _ = v8
                    }
                case "\n\t":
                    var v9 uint = (p0 / ((p0) | 1))
                    _ = v9
                default:
                    if (false && (42 >= (p2 % ((p2) | 1)))) {
                        if false {
                            if false {
                                v10 := []int{(fi4 | -9223372036854775808), -318, fi4, fi4}
                                fmt.Println(p2, p2, p2)
                                _ = v10
                            } else {
                                p2 += (p2 / (((p2 | p2)) | 1))
                                var v11 float64 = -2.25
                                _ = v11
                            }
                            v12 := []int{446, (p1 | (-(fi4)))}
                            _ = v12
                        }
                        fmt.Println(p2, fi4, p0)
                        p2 += (p2 / (((p2 << ((42) & 63))) | 1))
                    }
                }
                _ = v5
            }
        case "foo bar":
            if true {
                v13 := GenS1{F0: -128, F1: 4294967295, F2: 0}
                switch fi4 {
                case 0:
                    var v14 float64 = ((1e10 - 0.0) / 0.1)
                    _ = v14
                case 1:
                    var v15 string = ("" + "\n\t")
                    _ = v15
                }
                v16 := []int{(^(fi4)), p1, 9223372036854775807, fi4}
                _ = v13
                _ = v16
            }
        }
        v17 := []int{(p1 + 299), (p1 >> ((162) & 63)), (fi4 >> (((^(p0))) & 63)), p1}
        var v18 uint = p0
        _ = fi4
        _ = v17
        _ = v18
    }
    var v19 int64 = 422
    var v20 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return v19
}
func genFn1(p0 int16, p1 uint16) uint64 {
    if (!(65535 != 4294967295)) {
        fmt.Println(p0, p0)
    } else {
        var v21 int8 = 0
        var v22 uint = 65535
        _ = v21
        _ = v22
    }
    fmt.Println(p0)
    for fi23 := 0; fi23 < 1; fi23++ {
        for fi24 := 0; fi24 < 17; fi24++ {
            v25 := GenS1{F0: 100, F1: 42, F2: (p0 % ((p0) | 1))}
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    _ = p0
    _ = p1
    return 612
}

func main() {
    var v1 bool = ((false && (0.1 < 0.0)) || true)
    v2 := []string{"foo bar", ""}
    v3 := make(map[int]bool)
    v1 = (1 != 65535)
    v1 = ((3.14159 + (0.1 * 0.1)) != (1e10 * (1.5 * (float32((0.1 - 0.1)) - float32(1.5)))))
    if v1 {
        fmt.Println(v1)
        v26 := v3[4294967295]
        _ = v26
    } else {
        if v1 {
            var v27 int32 = 0
            var v28 int64 = -110
            v29 := GenS1{F0: -128, F1: 1, F2: 707}
            _ = v27
            _ = v28
            _ = v29
        }
    }
    v1 = (false || v1)
    switch ("hello" + "") {
    case "foo bar":
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v1)
        }
    case "a":
        delete(v3, -734)
    default:
        v1 = false
    }
    v30 := GenS2{F0: ((626 > 2147483647) && (v1 && v1)), F1: 1}
    switch "hello" {
    case "":
        for fi31 := 0; fi31 < 10; fi31++ {
            var v32 int = fi31
            _ = fi31
            _ = v32
        }
    case "hello":
        var v33 int8 = 100
        _ = v33
    case "\n\t":
        fmt.Println(v1)
    }
    switch "hello" {
    case "a":
        switch "foo bar" {
        case "a":
            if ((!(255 >= 255)) || v1) {
                var v34 string = "foo bar"
                var v35 int8 = 0
                _ = v34
                _ = v35
            }
        case "":
            v2 = append(v2, (("\n\t" + "a") + ""))
        case "foo bar":
            v1 = (42 == 127)
        default:
            switch (("" + ("" + "\n\t")) + "hello") {
            case "hello":
                _ = len(v2)
            case "a":
                v1 = v1
            case "\n\t":
                switch -1 {
                case 32768:
                    v1 = (false || v1)
                case -881:
                    v2 = append(v2, "\n\t")
                case 634:
                    _ = len(v3)
                }
            }
        }
    case "\n\t":
        if v1 {
            var v36 int = -882
            _ = v36
        } else {
            if true {
                if (2 == 127) {
                    for fi37 := 0; fi37 < 1; fi37++ {
                        fmt.Println(v1, fi37)
                        var v38 string = "\n\t"
                        _ = fi37
                        _ = v38
                    }
                    if (((!((!(v1 || v1)) && v1)) || ((42 != 771) && v1)) && ((1.5 / ((0.1 * (1e10 - float64(((((1e10 * (0.1 / float32(1e10))) - -2.25) / 1.0) - 0.0)))) * (3.14159 / (0.1 / 1e10)))) > 1.5)) {
                        var v39 int64 = -9223372036854775808
                        _ = v39
                    }
                } else {
                    for fi40 := 0; fi40 < 20; fi40++ {
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                        switch (-(fi40)) {
                        case 32768:
                            v1 = (-2147483648 >= 476)
                        default:
                            v1 = v1
                        }
                        _ = fi40
                    }
                    fmt.Println(v1)
                }
                if false {
                    var v41 uint32 = 2
                    for fi42 := 0; fi42 < 3; fi42++ {
                        v43 := GenS3{F0: v41, F1: 595}
                        _ = fi42
                        _ = v43
                    }
                    if true {
                        switch "" {
                        case "":
                            var v44 int64 = -9223372036854775808
                            _ = v44
                        case "\n\t":
                            if v1 {
                                v45 := v3[1]
                                _ = v45
                            } else {
                                v46 := GenS4{F0: 271, F1: genFn0(183, 0, 2147483648)}
                                var v47 int8 = 1
                                switch 42 {
                                case 32767:
                                    if len(v2) > 0 {
                                        v2[0] = "hello"
                                    }
                                case 0:
                                    fmt.Println(v41, v41)
                                case -673:
                                    if v1 {
                                        var v48 float32 = 0.1
                                        v49 := GenS4{F0: (v41 | (v41 & 228)), F1: 729}
                                        var v50 uint32 = (v41 - v41)
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                    }
                                }
                                _ = v46
                                _ = v47
                            }
                        }
                        for fi51 := 0; fi51 < 4; fi51++ {
                            fmt.Println(fi51, fi51)
                            switch "a" {
                            case "foo bar":
                                var v52 float64 = 1.5
                                _ = v52
                            case "":
                                var v53 bool = (v1 && v1)
                                _ = v53
                            default:
                                v3[fi51] = (!true)
                            }
                            _ = fi51
                        }
                    }
                    _ = v41
                }
            }
            switch 42 {
            case -1:
                v1 = true
            case 42:
                if (!v1) {
                    v1 = (0 == 1)
                    var v54 int64 = -9223372036854775808
                    v1 = (42 <= 32767)
                    _ = v54
                }
            case 32767:
                if (1 > 9223372036854775807) {
                    var v55 uint = 64
                    _ = v55
                }
            }
            v1 = (32767 <= 259)
        }
    case "hello":
        for fi56 := 0; fi56 < 13; fi56++ {
            if (42 >= -2147483648) {
                switch fi56 {
                case 1:
                    v57 := v3[(fi56 + fi56)]
                    _ = v57
                }
            }
            _ = fi56
        }
    default:
        _ = len(v2)
    }
    if v1 {
        var v58 float32 = (3.14159 / (float32(float64(3.14159)) - (1.5 / ((1e10 / (-2.25 / 1.5)) - (-2.25 * 1e10)))))
        fmt.Println(v58)
        var v59 int = 9223372036854775807
        _ = v58
        _ = v59
    } else {
        switch 4294967295 {
        case -162:
            switch 4294967295 {
            case 535:
                fmt.Println(v1)
            default:
                for fi60 := 0; fi60 < 15; fi60++ {
                    v1 = (631 <= 4294967295)
                    v61 := GenS2{F0: v1, F1: 42}
                    var v62 uint16 = 2
                    _ = fi60
                    _ = v61
                    _ = v62
                }
            }
        default:
            if v1 {
                switch 988 {
                case 4294967295:
                    fmt.Println(v1)
                default:
                    v2 = append(v2, (("hello" + "hello") + "hello"))
                }
            } else {
                var v63 int64 = -984
                _ = v63
            }
        }
    }
    for fi64 := 0; fi64 < 7; fi64++ {
        for fi65 := 0; fi65 < 2; fi65++ {
            var v66 int32 = -1
            _ = fi65
            _ = v66
        }
        _ = fi64
    }
    if true {
        if v1 {
            fmt.Println(v1)
            fmt.Println(v1)
        }
    } else {
        var v67 int16 = -315
        var v68 int = 127
        if (v1 && (!((-2.25 / 0.1) == (-2.25 / 1.0)))) {
            if (2147483648 < 2147483648) {
                v67 += 649
            }
        }
        _ = v67
        _ = v68
    }
    for fi69 := 0; fi69 < 12; fi69++ {
        fmt.Println(fi69)
        switch fi69 {
        case 42:
            fmt.Println(fi69, fi69)
        case 2147483647:
            if (!(!(("foo bar" + "") >= ("\n\t" + ("a" + "hello"))))) {
                for fi70 := 0; fi70 < 12; fi70++ {
                    var v71 uint32 = 193
                    if v1 {
                        var v72 uint16 = 32767
                        if (-226 == 608) {
                            if v1 {
                                var v73 int = fi70
                                var v74 float32 = float32(1.5)
                                if v1 {
                                    var v75 bool = false
                                    v72 = 13
                                    v74 = (v74 * (0.1 / 1e10))
                                    if ((genFn1(-955, (v72 ^ v72)) < genFn1(856, (v72 + v72))) && ((^(v71)) <= (^(v71)))) {
                                        if true {
                                            v75 = v1
                                            if (!v1) {
                                                v71 -= (v71 - v71)
                                                _ = len(v2)
                                                var v76 int16 = 1
                                                _ = v76
                                            }
                                        }
                                        switch "" {
                                        case "hello":
                                            v1 = ((1.5 + 1.5) >= (0.0 / 0.1))
                                        }
                                        var v77 int64 = -1
                                        if false {
                                            if v75 {
                                                if (!((0.0 + (1.5 / 3.14159)) != 1e10)) {
                                                    var v78 float32 = (v74 + -2.25)
                                                    if (v72 <= (v72 % (((v72 << ((78) & 63))) | 1))) {
                                                        if (4294967295 <= 2) {
                                                            v3[(v73 / (((v73 / (((fi70 << ((357) & 63))) | 1))) | 1))] = v1
                                                        } else {
                                                            var v79 bool = ((!v1) || v75)
                                                            v80 := GenS3{F0: v71, F1: 42}
                                                            if (!(0 >= uint64(v71))) {
                                                                v74 += (v74 + 3.14159)
                                                            }
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                    }
                                                    fmt.Println(fi70, fi69)
                                                    _ = v78
                                                }
                                                fmt.Println(fi69, v77)
                                            }
                                        }
                                        fmt.Println(v1)
                                        _ = v77
                                    }
                                    _ = v75
                                } else {
                                    v72 += v72
                                }
                                _ = v73
                                _ = v74
                            }
                            if (!(!("\n\t" < "a"))) {
                                if v1 {
                                    if (true || v1) {
                                        if false {
                                            if v1 {
                                                if v1 {
                                                    v71 -= (v71 * (v71 >> ((2) & 63)))
                                                    v81 := GenS4{F0: (v71 + v71), F1: -149}
                                                    var v82 uint8 = 0
                                                    _ = v81
                                                    _ = v82
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    switch 2147483647 {
                                                    case 4294967295:
                                                        var v83 int32 = int32(v72)
                                                        _ = v83
                                                    default:
                                                        var v84 bool = (((v72 - v72) < v72) && false)
                                                        v1 = v1
                                                        _ = v84
                                                    }
                                                    var v85 int64 = 782
                                                    _ = v85
                                                }
                                                _ = len(v2)
                                            } else {
                                                switch ("\n\t" + ("\n\t" + ("" + ("foo bar" + "")))) {
                                                case "":
                                                    if len(v2) > 0 {
                                                        v2[0] = (("\n\t" + ("\n\t" + ("foo bar" + ""))) + "a")
                                                    }
                                                case "\n\t":
                                                    v86 := GenS5{F0: v71, F1: 0, F2: 1}
                                                    _ = v86
                                                case "hello":
                                                    if true {
                                                        v1 = false
                                                        fmt.Println(v71)
                                                        v1 = true
                                                    } else {
                                                        v87 := v3[fi69]
                                                        var v88 bool = v1
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                default:
                                                    v89 := GenS6{F0: 42, F1: 86, F2: genFn0(2, (fi69 ^ 32768), (v71 + 516))}
                                                    _ = v89
                                                }
                                            }
                                            v71 += (v71 << ((1) & 63))
                                            v1 = (781 >= uint64(v72))
                                            switch ("hello" + "hello") {
                                            case "\n\t":
                                                v90 := GenS7{F0: (v72 - (v72 ^ (^(v72)))), F1: 2, F2: (32767 < 32767)}
                                                _ = v90
                                            case "hello":
                                                fmt.Println(v1, v71, v71)
                                            }
                                        } else {
                                            if (v1 && (v1 && (!v1))) {
                                                if len(v2) > 0 {
                                                    v2[0] = ("a" + "a")
                                                }
                                                switch ("hello" + "hello") {
                                                case "":
                                                    var v91 int64 = 203
                                                    _ = v91
                                                case "foo bar":
                                                    v71 += (v71 & v71)
                                                case "hello":
                                                    v71 += (v71 / ((65535) | 1))
                                                default:
                                                    v92 := GenS3{F0: (v71 >> ((194) & 63)), F1: -997}
                                                    _ = v92
                                                }
                                            }
                                            _ = len(v3)
                                            switch "a" {
                                            case "hello":
                                                v72 += v72
                                            default:
                                                var v93 string = (("\n\t" + "\n\t") + (((("\n\t" + "a") + "hello") + "foo bar") + "\n\t"))
                                                _ = v93
                                            }
                                            fmt.Println(v72, v1, v71)
                                        }
                                        var v94 int32 = -2147483648
                                        _ = v94
                                    }
                                    v1 = v1
                                } else {
                                    v3[(^(fi69))] = (!v1)
                                    delete(v3, (^(fi69)))
                                    v2 = append(v2, ("" + "foo bar"))
                                    if (!v1) {
                                        var v95 uint16 = 32767
                                        _ = v95
                                    }
                                }
                            } else {
                                switch "" {
                                case "foo bar":
                                    v71 = (v71 >> ((42) & 63))
                                case "\n\t":
                                    v71 += 760
                                }
                                fmt.Println(fi69)
                                v96 := GenS2{F0: false, F1: 1}
                                _ = v96
                            }
                            _ = len(v3)
                        }
                        switch "foo bar" {
                        case "\n\t":
                            v97 := GenS8{F0: 255, F1: v72}
                            _ = v97
                        case "a":
                            if (255 >= 162) {
                                if (-2.25 == -2.25) {
                                    fmt.Println(fi69, v1)
                                    var v98 float32 = (((1e10 * 1.5) / -2.25) * 1.5)
                                    var v99 int32 = 1
                                    _ = v98
                                    _ = v99
                                } else {
                                    v100 := v3[(fi69 & fi69)]
                                    switch "hello" {
                                    case "a":
                                        if (1 >= 255) {
                                            v101 := GenS4{F0: (v71 * v71), F1: -1}
                                            if (true && v1) {
                                                v100 = v1
                                                var v102 int64 = 9223372036854775807
                                                var v103 uint = 456
                                                _ = v102
                                                _ = v103
                                            }
                                            v1 = ("a" > "foo bar")
                                            _ = v101
                                        }
                                    case "":
                                        v104 := GenS7{F0: (v72 - (^(v72))), F1: 0, F2: v1}
                                        _ = v104
                                    }
                                    _ = v100
                                }
                            } else {
                                switch "\n\t" {
                                case "hello":
                                    v105 := GenS3{F0: (v71 * v71), F1: -589}
                                    _ = v105
                                }
                            }
                            fmt.Println(v72, fi69, v72)
                            var v106 int8 = -1
                            var v107 float64 = (0.0 - 3.14159)
                            _ = v106
                            _ = v107
                        default:
                            var v108 int64 = -9223372036854775808
                            v109 := GenS9{F0: "foo bar"}
                            v110 := GenS10{F0: 42}
                            _ = v108
                            _ = v109
                            _ = v110
                        }
                        if v1 {
                            v72 -= 42
                            v111 := GenS11{F0: 0, F1: (v72 + v72), F2: 236}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v111
                        }
                        v112 := GenS9{F0: "a"}
                        v113 := GenS7{F0: v72, F1: 1, F2: (("hello" < ("foo bar" + "")) && v1)}
                        var v114 int32 = 0
                        _ = v112
                        _ = v113
                        _ = v114
                        _ = v72
                    } else {
                        if v1 {
                            v115 := v3[fi69]
                            v116 := v3[(-(fi70))]
                            _ = v115
                            _ = v116
                        } else {
                            var v117 int16 = 1
                            _ = v117
                        }
                        v71 += v71
                    }
                    _ = fi70
                    _ = v71
                }
                var v118 int = (fi69 / (((fi69 * 1)) | 1))
                var v119 int16 = 1000
                _ = v118
                _ = v119
            } else {
                for fi120 := 0; fi120 < 12; fi120++ {
                    v121 := GenS12{F0: 255, F1: (("hello" + "") + ("" + "hello")), F2: ("a" + "foo bar")}
                    _ = fi120
                    _ = v121
                }
            }
        case 931:
            fmt.Println(fi69, v1)
        }
        for fi122 := 0; fi122 < 15; fi122++ {
            var v123 int = fi122
            v1 = false
            _ = fi122
            _ = v123
        }
        _ = fi69
    }
    v124 := v3[-911]
    if v124 {
        v124 = (0 > 1)
        fmt.Println(v1)
        v125 := GenS9{F0: "foo bar"}
        _ = v125
    }
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v124
    _ = v2
    _ = v3
    _ = v30
}
