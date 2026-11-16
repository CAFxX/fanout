package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 bool }
type GenS2 struct { F0 int32; F1 int; F2 int16 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 uint64; F2 int16 }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 int64; F1 uint32; F2 string }
type GenS9 struct { F0 int32; F1 uint32; F2 int32 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint64; F1 string; F2 int }
type GenS12 struct { F0 int64; F1 uint8; F2 string }
type GenS13 struct { F0 int64 }
func genFn0(p0 bool) int32 {
    p0 = p0
    fmt.Println(p0)
    for fi7 := 0; fi7 < 15; fi7++ {
        p0 = p0
        switch (fi7 | fi7) {
        case -2147483648:
            fmt.Println(fi7, p0)
        case 1:
            for fi8 := 0; fi8 < 18; fi8++ {
                switch "a" {
                case "hello":
                    var v9 bool = (213 > 0)
                    _ = v9
                }
                _ = fi8
            }
        case 32767:
            v10 := GenS1{F0: fi7, F1: 2, F2: true}
            _ = v10
        default:
            p0 = false
        }
        _ = fi7
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int16 = -32768
        v13 := []int{(fi11 + (fi11 << ((0) & 63))), 2147483647, (fi11 - 32767)}
        v13 = append(v13, 802)
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    return -1
}

func main() {
    v1 := []string{"\n\t", ("hello" + ("foo bar" + ("a" + "a"))), "\n\t"}
    v2 := make(map[int]bool)
    var v3 uint32 = 65535
    var v4 uint8 = 128
    v5 := []uint32{v3, (^(v3)), v3, v3}
    var v6 int8 = 6
    v3 = v3
    v2[94] = (751 != 0)
    if false {
        if true {
            fmt.Println(v6, v4, v4)
            if ((3.14159 / 0.1) < ((1e10 + (float64(0.0) + 0.0)) / 1e10)) {
                fmt.Println(v6)
            }
            v3 += (v3 >> ((42) & 63))
        } else {
            if (v3 != (v3 & v3)) {
                v14 := GenS1{F0: 127, F1: 88, F2: (1.5 > 0.0)}
                var v15 float32 = 1.5
                switch "\n\t" {
                case "\n\t":
                    if ((!(966 > 42)) && false) {
                        for fi16 := 0; fi16 < 19; fi16++ {
                            v15 -= float32(0.1)
                            _ = fi16
                        }
                        var v17 int16 = -32768
                        _ = v17
                    } else {
                        if false {
                            _ = len(v1)
                            v15 -= float32(1e10)
                            var v18 float64 = (-2.25 + (-2.25 - (1e10 - float64(0.1))))
                            _ = v18
                        }
                        v19 := GenS1{F0: 0, F1: v4, F2: false}
                        _ = v19
                    }
                case "":
                    var v20 bool = (42 <= 0)
                    _ = v20
                case "a":
                    v3 = (v3 + (^(v3)))
                }
                _ = v14
                _ = v15
            }
            var v21 uint8 = 255
            v22 := GenS2{F0: int32(v4), F1: 0, F2: -1}
            _ = v21
            _ = v22
        }
        v4 += (v4 & (v4 / ((v4) | 1)))
    } else {
        if (455 == 148) {
            var v23 float32 = (float32((0.1 - 3.14159)) * 0.0)
            for fi24 := 0; fi24 < 1; fi24++ {
                fmt.Println(v4, v6, v4)
                var v25 int16 = 1000
                if ((!true) || false) {
                    v25 = -1
                    var v26 int = (fi24 % (((fi24 ^ (-(fi24)))) | 1))
                    for fi27 := 0; fi27 < 15; fi27++ {
                        switch (v26 << ((365) & 63)) {
                        case -9223372036854775808:
                            var v28 uint32 = v3
                            _ = v28
                        case 0:
                            if (false || (uint(v4) <= 255)) {
                                if false {
                                    if true {
                                        if len(v5) > 0 {
                                            v5[0] = (v3 - v3)
                                        }
                                        switch "hello" {
                                        case "a":
                                            v23 += 1e10
                                        case "foo bar":
                                            if true {
                                                if (!true) {
                                                    switch ("\n\t" + ("" + "a")) {
                                                    case "foo bar":
                                                        v23 -= (0.1 + (v23 / (v23 / (((float32(1.5) - v23) - 3.14159) / (((float32(3.14159) + 0.1) * ((((float32(3.14159) + (v23 / v23)) * ((v23 * (0.0 + v23)) * (v23 + (((float32(1e10) * (v23 - float32(0.0))) / v23) / (float32(0.0) + float32((3.14159 * (1.5 + 1e10)))))))) * (v23 / -2.25)) + 3.14159)) + v23)))))
                                                    case "hello":
                                                        v29 := GenS3{F0: (v4 & 127), F1: "\n\t"}
                                                        var v30 int64 = -1
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                    if false {
                                                        switch "\n\t" {
                                                        case "a":
                                                            v6 += -84
                                                            var v31 int32 = -1
                                                            _ = v31
                                                        case "foo bar":
                                                            fmt.Println(v25)
                                                        case "\n\t":
                                                            v1 = append(v1, "\n\t")
                                                        default:
                                                            v1 = append(v1, ("hello" + ""))
                                                        }
                                                        v23 = (v23 / v23)
                                                        v6 -= 42
                                                    }
                                                }
                                                fmt.Println(v4, v25, v3)
                                            }
                                            var v32 int16 = -32768
                                            var v33 int = fi27
                                            _ = v32
                                            _ = v33
                                        default:
                                            if false {
                                                v6 = -1
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v34 := GenS2{F0: 2147483647, F1: (v26 % ((v26) | 1)), F2: 32767}
                                                    _ = v34
                                                default:
                                                    fmt.Println(v3, fi24, fi24)
                                                }
                                            } else {
                                                fmt.Println(v6, v4)
                                                var v35 float32 = (float32(3.14159) / v23)
                                                _ = v35
                                            }
                                        }
                                    } else {
                                        v36 := GenS4{F0: -1, F1: (522 <= 256)}
                                        v25 = -856
                                        fmt.Println(v26, fi24, v4)
                                        var v37 int8 = 100
                                        _ = v36
                                        _ = v37
                                    }
                                    v38 := GenS5{F0: 32767}
                                    _ = v38
                                } else {
                                    var v39 int32 = -1
                                    _ = v39
                                }
                                v26 -= v26
                                switch 4294967295 {
                                case 9223372036854775807:
                                    v2[(-(fi24))] = (!false)
                                case 32767:
                                    var v40 int8 = (v6 >> ((2) & 63))
                                    _ = v40
                                case 128:
                                    var v41 float64 = (1e10 + float64((v23 + v23)))
                                    if true {
                                        fmt.Println(v3, v25)
                                    } else {
                                        v42 := GenS5{F0: 990}
                                        _ = v42
                                    }
                                    var v43 uint = 1
                                    _ = v41
                                    _ = v43
                                default:
                                    v44 := GenS5{F0: 65535}
                                    _ = v44
                                }
                            } else {
                                v25 += (v25 | -392)
                            }
                            fmt.Println(v25, v6, v3)
                        case -1:
                            if (413 >= -792) {
                                v6 = v6
                                var v45 string = "\n\t"
                                var v46 uint32 = 4294967295
                                _ = v45
                                _ = v46
                            } else {
                                v25 = 0
                                v4 = (v4 & 174)
                                var v47 uint16 = 32767
                                _ = v47
                            }
                        default:
                            v48 := GenS2{F0: -1, F1: (^(fi24)), F2: (v25 % ((12) | 1))}
                            _ = v48
                        }
                        _ = fi27
                    }
                    _ = v26
                } else {
                    var v49 bool = (-2147483648 == -2147483648)
                    if v49 {
                        v3 = (v3 | (v3 | v3))
                        fmt.Println(v49, v4)
                    } else {
                        for fi50 := 0; fi50 < 5; fi50++ {
                            fmt.Println(v4)
                            _ = fi50
                        }
                        var v51 string = "hello"
                        _ = v51
                    }
                    var v52 int64 = 1
                    _ = v49
                    _ = v52
                }
                _ = fi24
                _ = v25
            }
            for fi53 := 0; fi53 < 11; fi53++ {
                v4 = 52
                _ = fi53
            }
            _ = v23
        }
        switch -9223372036854775808 {
        case -2147483648:
            v54 := GenS1{F0: 823, F1: 0, F2: (!true)}
            _ = v54
        default:
            v55 := GenS1{F0: 4294967295, F1: (^(v4)), F2: true}
            _ = v55
        }
        for fi56 := 0; fi56 < 17; fi56++ {
            v57 := GenS3{F0: v4, F1: ("hello" + ("" + ""))}
            _ = fi56
            _ = v57
        }
    }
    var v58 uint16 = 32767
    if (!(!(false && true))) {
        switch 1 {
        case -2147483648:
            v58 = 0
        case -633:
            v59 := GenS1{F0: 32768, F1: (v4 % (((v4 % ((v4) | 1))) | 1)), F2: (!(!false))}
            _ = v59
        default:
            for fi60 := 0; fi60 < 4; fi60++ {
                v58 = (v58 >> ((uint(v58)) & 63))
                if (1000 <= 596) {
                    fmt.Println(v6, v4, fi60)
                    var v61 uint8 = 190
                    if (true && false) {
                        var v62 uint32 = v3
                        _ = len(v2)
                        v4 -= 0
                        _ = v62
                    }
                    _ = v61
                }
                v63 := GenS6{F0: "\n\t", F1: 4294967295, F2: 1}
                _ = fi60
                _ = v63
            }
        }
        switch "hello" {
        case "":
            if false {
                var v64 int = 0
                v65 := GenS7{F0: 4294967295, F1: (v64 % (((^(v64))) | 1))}
                _ = v64
                _ = v65
            }
        case "foo bar":
            switch "a" {
            case "":
                v66 := v2[32767]
                _ = v66
            case "foo bar":
                if true {
                    v6 -= 100
                }
            }
        }
        if true {
            _ = len(v5)
        } else {
            for fi67 := 0; fi67 < 4; fi67++ {
                v6 -= -1
                _ = fi67
            }
            var v68 uint16 = (v58 >> ((718) & 63))
            _ = v68
        }
    }
    v4 = (v4 ^ v4)
    if true {
        if ((-(v6)) > v6) {
            var v69 int8 = -1
            switch "foo bar" {
            case "a":
                v70 := GenS8{F0: -116, F1: 42, F2: ("\n\t" + "a")}
                _ = v70
            }
            _ = v69
        } else {
            v4 += 0
            if true {
                if false {
                    for fi71 := 0; fi71 < 2; fi71++ {
                        fmt.Println(fi71, v6, v58)
                        _ = fi71
                    }
                    var v72 string = "\n\t"
                    v73 := GenS9{F0: 690, F1: 1, F2: -728}
                    _ = v72
                    _ = v73
                } else {
                    var v74 int16 = -32768
                    fmt.Println(v3, v6)
                    _ = v74
                }
            } else {
                for fi75 := 0; fi75 < 1; fi75++ {
                    var v76 uint32 = (v3 | v3)
                    var v77 int32 = 42
                    if false {
                        v77 = 350
                    } else {
                        v2[(fi75 << ((65535) & 63))] = false
                    }
                    _ = fi75
                    _ = v76
                    _ = v77
                }
                v6 = v6
                v78 := GenS2{F0: 1, F1: -410, F2: 0}
                _ = v78
            }
        }
        for fi79 := 0; fi79 < 4; fi79++ {
            var v80 float64 = 3.14159
            v81 := GenS3{F0: v4, F1: ("foo bar" + "a")}
            v82 := GenS5{F0: (v58 ^ v58)}
            _ = fi79
            _ = v80
            _ = v81
            _ = v82
        }
        switch 4294967295 {
        case 127:
            for fi83 := 0; fi83 < 12; fi83++ {
                if ((false && true) && (0 > 0)) {
                    v3 += 2147483647
                    v84 := GenS1{F0: fi83, F1: v4, F2: (v58 <= v58)}
                    switch 127 {
                    case -440:
                        switch "hello" {
                        case "\n\t":
                            switch "a" {
                            case "\n\t":
                                var v85 int32 = -27
                                _ = v85
                            case "hello":
                                v86 := GenS10{F0: 51}
                                _ = v86
                            case "":
                                v4 = uint8(v58)
                            }
                        case "hello":
                            v3 += (v3 / (((v3 % (((v3 ^ (v3 * v3))) | 1))) | 1))
                        case "":
                            var v87 bool = (-171 <= 0)
                            _ = v87
                        }
                    case -698:
                        if ("hello" != "\n\t") {
                            if true {
                                for fi88 := 0; fi88 < 16; fi88++ {
                                    if (v3 != (v3 | 65535)) {
                                        v6 = v6
                                    }
                                    _ = len(v2)
                                    var v89 uint64 = 387
                                    _ = fi88
                                    _ = v89
                                }
                                var v90 string = "a"
                                v90 = v90
                                _ = v90
                            } else {
                                v91 := GenS2{F0: -561, F1: fi83, F2: 42}
                                _ = v91
                            }
                            v6 = (v6 + v6)
                        } else {
                            var v92 uint8 = (v4 % (((v4 | (v4 / ((v4) | 1)))) | 1))
                            v4 = 2
                            if ((v6 <= v6) && (v6 == v6)) {
                                switch "" {
                                case "hello":
                                    v1 = append(v1, "foo bar")
                                }
                            }
                            _ = v92
                        }
                    case 32768:
                        var v93 int16 = 42
                        _ = v93
                    default:
                        switch (^(fi83)) {
                        case 32767:
                            var v94 int8 = (v6 % ((v6) | 1))
                            _ = v94
                        case -1:
                            fmt.Println(v58, v3)
                        default:
                            if true {
                                v95 := v2[(fi83 >> ((121) & 63))]
                                v95 = ((v6 & int8(v3)) > v6)
                                _ = v95
                            }
                        }
                    }
                    _ = v84
                } else {
                    _ = len(v2)
                    if (fi83 <= (fi83 | 128)) {
                        v58 -= v58
                        v58 = v58
                    } else {
                        for fi96 := 0; fi96 < 3; fi96++ {
                            switch "" {
                            case "":
                                v2[(fi83 * (fi96 / (((-(fi83))) | 1)))] = true
                            case "a":
                                v97 := v2[fi83]
                                _ = v97
                            }
                            _ = fi96
                        }
                    }
                }
                _ = fi83
            }
        }
    }
    fmt.Println(v58)
    if (!(42 == 18446744073709551615)) {
        v98 := GenS11{F0: 18446744073709551615, F1: ("a" + ((("hello" + "\n\t") + (("\n\t" + "foo bar") + "hello")) + ((((("foo bar" + "foo bar") + "a") + "\n\t") + ("a" + "")) + "hello"))), F2: 320}
        if true {
            v58 = (v58 << ((42) & 63))
            v1 = append(v1, "foo bar")
        }
        if true {
            switch "foo bar" {
            case "foo bar":
                v3 -= (^(v3))
            default:
                for fi99 := 0; fi99 < 9; fi99++ {
                    var v100 string = ("hello" + "")
                    v101 := GenS1{F0: fi99, F1: (v4 % (((^(v4))) | 1)), F2: true}
                    _ = fi99
                    _ = v100
                    _ = v101
                }
            }
            v102 := GenS12{F0: 1, F1: v4, F2: ("a" + "\n\t")}
            for fi103 := 0; fi103 < 1; fi103++ {
                switch ("hello" + (("hello" + "") + "")) {
                case "a":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "":
                    switch "foo bar" {
                    case "a":
                        v104 := GenS4{F0: 9223372036854775807, F1: true}
                        _ = v104
                    default:
                        var v105 float64 = 3.14159
                        _ = v105
                    }
                default:
                    fmt.Println(v3)
                }
                if false {
                    _ = len(v2)
                }
                if true {
                    if true {
                        switch "" {
                        case "foo bar":
                            switch 127 {
                            case -269:
                                v106 := GenS9{F0: -943, F1: v3, F2: 0}
                                _ = v106
                            }
                        default:
                            v3 += v3
                        }
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    v107 := GenS13{F0: -211}
                    if (18446744073709551615 != 18446744073709551615) {
                        v5 = append(v5, v3)
                        var v108 float32 = 0.1
                        var v109 uint32 = 65535
                        _ = v108
                        _ = v109
                    } else {
                        var v110 float32 = -2.25
                        if false {
                            v110 -= v110
                        }
                        _ = v110
                    }
                    _ = v107
                } else {
                    v3 = (^(v3))
                }
                _ = fi103
            }
            _ = v102
        }
        _ = v98
    } else {
        for fi111 := 0; fi111 < 18; fi111++ {
            v6 += v6
            var v112 bool = false
            _ = fi111
            _ = v112
        }
    }
    if true {
        v58 -= (v58 & v58)
        var v113 uint16 = (v58 * v58)
        _ = v113
    } else {
        v58 = 1
        switch -2147483648 {
        case 4294967295:
            _ = len(v1)
        case 128:
            var v114 bool = (3.14159 == -2.25)
            _ = v114
        case 127:
            if (!((v3 % ((0) | 1)) >= v3)) {
                v115 := GenS7{F0: 585, F1: 9223372036854775807}
                var v116 int32 = 0
                v117 := GenS5{F0: 255}
                _ = v115
                _ = v116
                _ = v117
            }
        }
    }
    switch -9223372036854775808 {
    case 431:
        v2[0] = false
    default:
        switch int(v4) {
        case 899:
            v118 := v2[255]
            _ = v118
        }
    }
    if true {
        v119 := GenS4{F0: -861, F1: (true && true)}
        if (262 <= 602) {
            var v120 int8 = 36
            v121 := GenS3{F0: v4, F1: ""}
            _ = v120
            _ = v121
        } else {
            var v122 uint64 = 0
            _ = v122
        }
        _ = v119
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v123 uint = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v123:", v123)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v123
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
}
