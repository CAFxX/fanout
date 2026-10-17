package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 int; F2 string }
type GenS6 struct { F0 uint; F1 uint; F2 int64 }
type GenS7 struct { F0 string; F1 int64; F2 uint8 }
type GenS8 struct { F0 int32; F1 uint }
type GenS9 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS10 struct { F0 int16; F1 int64 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint8; F1 bool; F2 int64 }
type GenS13 struct { F0 int64; F1 int }
type GenS14 struct { F0 bool }
func genFn0(p0 int64, p1 string, p2 uint) bool {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint = 255
        v8 := GenS1{F0: (p1 + "foo bar"), F1: 11}
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(p0, p0)
    var v9 int8 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return false
}
func genFn1(p0 uint64, p1 uint32) uint32 {
    var v10 uint = 18446744073709551615
    v10 = 2
    var v11 uint8 = 209
    if false {
        v12 := GenS1{F0: ("foo bar" + "\n\t"), F1: v11}
        _ = v12
    } else {
        p1 = p1
        if (-416 >= 32767) {
            v13 := []int{}
            if true {
                v14 := GenS1{F0: "", F1: 41}
                var v15 bool = true
                _ = v14
                _ = v15
            } else {
                fmt.Println(p1)
                switch int(p0) {
                case -9223372036854775808:
                    fmt.Println(v10, v11)
                case 2147483647:
                    var v16 uint16 = 0
                    _ = v16
                default:
                    v17 := GenS2{F0: ("" + ("a" + "hello")), F1: 42}
                    _ = v17
                }
                for fi18 := 0; fi18 < 11; fi18++ {
                    var v19 int32 = 2147483647
                    v10 = (v10 / (((v10 - uint(v11))) | 1))
                    _ = fi18
                    _ = v19
                }
            }
            _ = v13
        } else {
            v20 := []int{1, 797, 0}
            _ = v20
        }
    }
    v21 := []int{42, -1, -1}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v21
    return (p1 & p1)
}

func main() {
    var v1 uint32 = 65535
    var v2 uint8 = 42
    var v3 int8 = 127
    var v4 uint16 = 32767
    var v5 bool = false
    v2 = 42
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 int8 = v3
        v24 := []int{fi22, (fi22 ^ fi22)}
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := GenS2{F0: "a", F1: 1}
    v26 := []int{}
    switch 127 {
    case 128:
        v27 := GenS3{F0: 42}
        _ = v27
    case 32768:
        v4 = 255
    case -425:
        fmt.Println(v3, v2, v2)
    default:
        var v28 float32 = 0.1
        _ = v28
    }
    var v29 float32 = (float32((((1e10 - 1e10) * -2.25) - 3.14159)) / 1.5)
    v30 := GenS1{F0: "", F1: 8}
    v29 = (v29 - ((0.0 / (v29 * v29)) + v29))
    v5 = v5
    if v5 {
        if v5 {
            for fi31 := 0; fi31 < 6; fi31++ {
                fmt.Println(v1, v5)
                _ = fi31
            }
            for fi32 := 0; fi32 < 14; fi32++ {
                for fi33 := 0; fi33 < 4; fi33++ {
                    var v34 int32 = 69
                    _ = fi33
                    _ = v34
                }
                _ = fi32
            }
        } else {
            var v35 int16 = 1
            _ = v35
        }
    }
    v2 = v2
    var v36 string = ""
    if v5 {
        for fi37 := 0; fi37 < 5; fi37++ {
            switch "hello" {
            case "a":
                for fi38 := 0; fi38 < 18; fi38++ {
                    v29 += v29
                    if v5 {
                        if len(v26) > 0 {
                            v26[0] = (fi37 / ((4294967295) | 1))
                        }
                        var v39 uint8 = 127
                        _ = v39
                    } else {
                        fmt.Println(v29, v2, fi37)
                        v4 += v4
                        if v5 {
                            if (fi37 > 32767) {
                                v36 = ((("" + v36) + (v36 + v36)) + "hello")
                                fmt.Println(v4, v1, fi38)
                                _ = len(v26)
                            } else {
                                var v40 int32 = -2147483648
                                _ = v40
                            }
                        } else {
                            if (v5 && (!false)) {
                                var v41 string = (v36 + v36)
                                var v42 int8 = (-(v3))
                                v29 = ((0.1 - v29) + v29)
                                _ = v41
                                _ = v42
                            }
                            var v43 uint64 = 1
                            v44 := GenS4{F0: "foo bar"}
                            _ = v43
                            _ = v44
                        }
                    }
                    switch -953 {
                    case 128:
                        if len(v26) > 0 {
                            v26[0] = fi37
                        }
                    case 1:
                        if len(v26) > 0 {
                            v26[0] = fi38
                        }
                    default:
                        var v45 int32 = 660
                        _ = v45
                    }
                    _ = fi38
                }
            case "hello":
                for fi46 := 0; fi46 < 19; fi46++ {
                    v47 := GenS4{F0: (v36 + v36)}
                    fmt.Println(v36)
                    var v48 uint32 = (^(v1))
                    _ = fi46
                    _ = v47
                    _ = v48
                }
            case "":
                switch fi37 {
                case 1:
                    var v49 bool = false
                    _ = v49
                case 42:
                    var v50 uint8 = v2
                    _ = v50
                default:
                    fmt.Println(v1)
                }
            }
            v1 -= (v1 * v1)
            _ = fi37
        }
        var v51 int8 = -1
        for fi52 := 0; fi52 < 4; fi52++ {
            var v53 bool = v5
            v54 := GenS2{F0: "", F1: 42}
            v2 += v2
            _ = fi52
            _ = v53
            _ = v54
        }
        _ = v51
    } else {
        var v55 uint64 = 0
        v56 := GenS5{F0: -346, F1: 793, F2: v36}
        _ = v55
        _ = v56
    }
    for fi57 := 0; fi57 < 13; fi57++ {
        v29 = (v29 + v29)
        v4 += (^(v4))
        _ = fi57
    }
    switch (v36 + "") {
    case "hello":
        v2 = v2
    case "foo bar":
        _ = len(v26)
    default:
        for fi58 := 0; fi58 < 7; fi58++ {
            v29 += -2.25
            v59 := GenS2{F0: v36, F1: 42}
            _ = fi58
            _ = v59
        }
    }
    v60 := GenS6{F0: 4294967295, F1: 18446744073709551615, F2: -1}
    for fi61 := 0; fi61 < 1; fi61++ {
        for fi62 := 0; fi62 < 13; fi62++ {
            v63 := GenS7{F0: (v36 + v36), F1: -755, F2: (v2 / ((v2) | 1))}
            v64 := GenS8{F0: -2147483648, F1: 0}
            _ = fi62
            _ = v63
            _ = v64
        }
        for fi65 := 0; fi65 < 9; fi65++ {
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
            _ = fi65
        }
        _ = fi61
    }
    for fi66 := 0; fi66 < 8; fi66++ {
        for fi67 := 0; fi67 < 19; fi67++ {
            if v5 {
                v2 = 42
                var v68 bool = v5
                var v69 uint64 = 1
                switch ((v36 + v36) + ("" + ((v36 + ((v36 + "hello") + (((v36 + (v36 + "hello")) + v36) + ("hello" + ("hello" + v36))))) + v36))) {
                case "a":
                    if len(v26) > 0 {
                        v26[0] = (-(fi66))
                    }
                default:
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                }
                _ = v68
                _ = v69
            }
            switch (v36 + v36) {
            case "":
                if len(v26) > 0 {
                    v26[0] = (fi67 >> ((4294967295) & 63))
                }
            case "a":
                if v5 {
                    var v70 string = "a"
                    if len(v26) > 0 {
                        v26[0] = 32768
                    }
                    var v71 uint16 = (v4 ^ v4)
                    _ = v70
                    _ = v71
                } else {
                    if v5 {
                        var v72 int64 = 1
                        if (v5 || true) {
                            v73 := GenS9{F0: (^(v1)), F1: (v72 & (v72 % (((v72 << ((18446744073709551615) & 63))) | 1))), F2: v1}
                            if true {
                                v74 := GenS8{F0: -2147483648, F1: 4294967295}
                                var v75 uint32 = (v1 % ((genFn1(42, 2147483647)) | 1))
                                _ = v74
                                _ = v75
                            } else {
                                var v76 uint32 = (^(v1))
                                if v5 {
                                    switch (fi67 << ((2) & 63)) {
                                    case -645:
                                        if v5 {
                                            v77 := GenS9{F0: v76, F1: 1, F2: v1}
                                            if v5 {
                                                var v78 int16 = 1
                                                v79 := GenS4{F0: v36}
                                                if v5 {
                                                    var v80 uint32 = 1
                                                    v78 = int16(v72)
                                                    var v81 uint8 = (v2 - 255)
                                                    _ = v80
                                                    _ = v81
                                                }
                                                _ = v78
                                                _ = v79
                                            }
                                            var v82 uint64 = 1
                                            _ = v77
                                            _ = v82
                                        }
                                    case -602:
                                        v4 = v4
                                    }
                                    if len(v26) >= 2 {
                                        v26 = v26[1:len(v26)]
                                    }
                                    v3 = (v3 | v3)
                                }
                                _ = v76
                            }
                            if (((v1 + genFn1(485, v1)) > v1) && (1000 > -32768)) {
                                var v83 bool = false
                                if (1 > -32768) {
                                    var v84 uint = 694
                                    var v85 uint8 = (^(v2))
                                    _ = v84
                                    _ = v85
                                } else {
                                    switch (v36 + v36) {
                                    case "":
                                        if false {
                                            v72 += v72
                                        }
                                    case "hello":
                                        v3 = (v3 * v3)
                                        v4 = v4
                                        _ = len(v26)
                                    case "\n\t":
                                        if (((0 < 42) || v5) || (1.5 >= -2.25)) {
                                            fmt.Println(v2, fi66, fi66)
                                            v26 = append(v26, fi66)
                                        }
                                    }
                                }
                                v86 := GenS10{F0: -340, F1: v72}
                                _ = v83
                                _ = v86
                            }
                            _ = v73
                        } else {
                            var v87 float64 = -2.25
                            _ = v87
                        }
                        _ = v72
                    }
                    switch fi66 {
                    case 127:
                        v3 = v3
                        v36 = (v36 + v36)
                        var v88 uint64 = 4294967295
                        _ = v88
                    case 1:
                        v89 := GenS11{F0: (!(!((838 >= (fi67 & fi66)) || (v5 && (42 > 42)))))}
                        _ = v89
                    case -9223372036854775808:
                        if true {
                            fmt.Println(v1, v29, v1)
                            v90 := GenS1{F0: "a", F1: v2}
                            _ = v90
                        } else {
                            v5 = v5
                            var v91 string = "\n\t"
                            if len(v26) > 0 {
                                v26[0] = fi66
                            }
                            switch fi67 {
                            case 1:
                                fmt.Println(fi66, v91)
                            default:
                                v36 = v91
                            }
                            if (!v5) {
                                v92 := GenS12{F0: 186, F1: v5, F2: 9223372036854775807}
                                switch ("" + (((v36 + v91) + ((v91 + "") + ("" + ""))) + ((("hello" + ("foo bar" + (v36 + "foo bar"))) + v91) + v91))) {
                                case "\n\t":
                                    v1 -= v1
                                case "hello":
                                    switch int(v4) {
                                    case 127:
                                        v1 -= v1
                                        v1 = 800
                                        v26 = append(v26, fi67)
                                        fmt.Println(v29, v91)
                                        v93 := GenS13{F0: 1, F1: (fi66 ^ (fi67 / ((fi66) | 1)))}
                                        _ = v93
                                    case 2147483647:
                                        if len(v26) > 0 {
                                            v26[0] = fi67
                                        }
                                    case 32767:
                                        fmt.Println(v5, fi67)
                                    default:
                                        _ = len(v26)
                                    }
                                case "a":
                                    switch (v36 + "foo bar") {
                                    case "a":
                                        if len(v26) > 0 {
                                            v26[0] = (fi67 - fi67)
                                        }
                                    case "":
                                        switch (v91 + (((((("foo bar" + (((("\n\t" + v91) + "a") + v36) + "")) + v91) + v91) + "hello") + v36) + (v36 + "a"))) {
                                        case "\n\t":
                                            var v94 string = v91
                                            _ = v94
                                        case "hello":
                                            fmt.Println(v2, v3)
                                        case "a":
                                            if (!false) {
                                                var v95 uint32 = (v1 & genFn1(2, (v1 * (v1 | (v1 * v1)))))
                                                var v96 float64 = 0.0
                                                _ = v95
                                                _ = v96
                                            }
                                        }
                                    case "\n\t":
                                        switch ("a" + v36) {
                                        case "foo bar":
                                            if (((^(fi66)) >= fi67) && (v5 && true)) {
                                                var v97 int32 = 2147483647
                                                var v98 string = v91
                                                if (18446744073709551615 >= 1) {
                                                    var v99 string = (v98 + v91)
                                                    var v100 uint = 916
                                                    v99 = (((v99 + ("\n\t" + (((v91 + (v99 + v91)) + v98) + v99))) + "hello") + "hello")
                                                    _ = v100
                                                    _ = v99
                                                } else {
                                                    v101 := GenS7{F0: "\n\t", F1: 9223372036854775807, F2: 42}
                                                    switch (fi67 % ((int(v1)) | 1)) {
                                                    case 4294967295:
                                                        var v102 int64 = -9223372036854775808
                                                        _ = v102
                                                    }
                                                    fmt.Println(fi66, v1)
                                                    _ = v101
                                                }
                                                _ = v97
                                                _ = v98
                                            }
                                        }
                                    default:
                                        var v103 int8 = v3
                                        _ = v103
                                    }
                                default:
                                    if (!v5) {
                                        v1 = 65535
                                        var v104 float64 = 3.14159
                                        v105 := GenS5{F0: 42, F1: -991, F2: "hello"}
                                        _ = v104
                                        _ = v105
                                    } else {
                                        if len(v26) > 0 {
                                            v26[0] = (fi66 % ((fi66) | 1))
                                        }
                                        if v5 {
                                            switch (fi66 | (fi67 - (fi67 * -9223372036854775808))) {
                                            case 127:
                                                if (int64(v4) == -1) {
                                                    fmt.Println(v36)
                                                    var v106 bool = (!((1.5 * ((1.5 * v29) / v29)) <= ((float32(0.0) - -2.25) * 1e10)))
                                                    _ = v106
                                                }
                                            case 903:
                                                if (-32768 >= 1000) {
                                                    fmt.Println(fi67, v29)
                                                }
                                            case 32767:
                                                v91 = v36
                                            default:
                                                fmt.Println(fi67)
                                            }
                                            if len(v26) >= 2 {
                                                v26 = v26[1:len(v26)]
                                            }
                                            v36 = (v91 + ("hello" + v36))
                                        }
                                    }
                                }
                                _ = len(v26)
                                _ = v92
                            }
                            _ = v91
                        }
                    default:
                        v107 := GenS9{F0: genFn1(311, (^(v1))), F1: 687, F2: uint32(v2)}
                        _ = v107
                    }
                    switch (fi66 & (fi66 >> ((1) & 63))) {
                    case -2147483648:
                        var v108 uint64 = 42
                        if len(v26) > 0 {
                            v26[0] = fi66
                        }
                        _ = v108
                    case -606:
                        v1 += v1
                        fmt.Println(v1)
                    case 128:
                        var v109 uint16 = (v4 / ((v4) | 1))
                        _ = v109
                    }
                }
                var v110 int32 = 42
                var v111 int32 = v110
                var v112 uint64 = 521
                if false {
                    if v5 {
                        v112 -= uint64(v1)
                        var v113 string = (v36 + v36)
                        switch int(v4) {
                        case -825:
                            v112 = v112
                        default:
                            fmt.Println(v2, v4)
                        }
                        _ = v113
                    } else {
                        v2 = v2
                    }
                    if (!false) {
                        v26 = append(v26, 128)
                    } else {
                        var v114 int8 = v3
                        _ = v114
                    }
                }
                _ = v110
                _ = v111
                _ = v112
            default:
                if (("foo bar" + v36) < (v36 + "hello")) {
                    v36 = v36
                    v4 += v4
                } else {
                    fmt.Println(v29)
                    if v5 {
                        v115 := GenS4{F0: v36}
                        _ = v115
                    } else {
                        _ = len(v26)
                        v26 = append(v26, fi66)
                        var v116 uint64 = 609
                        var v117 uint8 = (v2 ^ v2)
                        v118 := GenS3{F0: 1000}
                        _ = v116
                        _ = v117
                        _ = v118
                    }
                }
            }
            _ = fi67
        }
        if (v5 || false) {
            v119 := GenS14{F0: (v5 && (42 != 1))}
            _ = v119
        }
        if len(v26) > 0 {
            v26[0] = fi66
        }
        _ = fi66
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
    _ = v60
}
