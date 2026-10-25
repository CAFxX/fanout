package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 bool; F2 int64 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 uint64; F2 int }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint16; F1 int16; F2 int }
type GenS12 struct { F0 int64; F1 uint64; F2 int64 }
type GenS13 struct { F0 uint; F1 int64 }

func main() {
    var v1 int8 = -1
    v2 := []uint32{2147483647, 2147483647, 313, 2}
    v3 := []string{"a", "hello", "a", "a"}
    fmt.Println(v1)
    if (!true) {
        _ = len(v3)
        for fi4 := 0; fi4 < 6; fi4++ {
            var v5 int = fi4
            _ = fi4
            _ = v5
        }
    } else {
        var v6 float32 = 1e10
        v6 += (float32((-2.25 + 1e10)) + v6)
        _ = v6
    }
    if false {
        v1 = v1
        if false {
            if len(v3) > 0 {
                v3[0] = ("foo bar" + "foo bar")
            }
        } else {
            var v7 float64 = -2.25
            _ = v7
        }
        v3 = append(v3, "a")
    } else {
        if len(v2) > 0 {
            v2[0] = 0
        }
        for fi8 := 0; fi8 < 3; fi8++ {
            if (-14 < 732) {
                if (v1 != 75) {
                    var v9 int64 = 669
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 7; fi10++ {
                        fmt.Println(fi8)
                        _ = fi10
                    }
                    for fi11 := 0; fi11 < 3; fi11++ {
                        var v12 int = fi8
                        _ = fi11
                        _ = v12
                    }
                    v1 -= 1
                }
                for fi13 := 0; fi13 < 20; fi13++ {
                    v1 = (-(v1))
                    var v14 int16 = 65
                    _ = fi13
                    _ = v14
                }
                if false {
                    fmt.Println(v1, fi8)
                    fmt.Println(fi8, fi8)
                    v15 := GenS1{F0: "\n\t"}
                    _ = v15
                }
            }
            v16 := GenS1{F0: "foo bar"}
            switch 32767 {
            case 1:
                var v17 int8 = (-(v1))
                _ = v17
            case 270:
                v3 = append(v3, "hello")
            default:
                v1 += (v1 % ((v1) | 1))
            }
            _ = fi8
            _ = v16
        }
        v3 = append(v3, "\n\t")
    }
    switch (("foo bar" + "foo bar") + "\n\t") {
    case "hello":
        v18 := GenS1{F0: ((("\n\t" + "a") + ("foo bar" + (("hello" + ("hello" + (("a" + "\n\t") + "foo bar"))) + ("hello" + ("foo bar" + "hello"))))) + "")}
        _ = v18
    case "\n\t":
        var v19 uint = 255
        _ = v19
    default:
        v3 = append(v3, "\n\t")
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 int16 = 428
        var v22 float32 = 0.1
        v1 -= (-(v1))
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint64 = 18446744073709551615
    fmt.Println(v23, v1)
    if true {
        var v24 int32 = 1
        if true {
            v23 -= v23
            for fi25 := 0; fi25 < 14; fi25++ {
                if false {
                    _ = len(v2)
                }
                if true {
                    v24 -= v24
                } else {
                    fmt.Println(v1, v1, fi25)
                    if len(v3) > 0 {
                        v3[0] = (("a" + "foo bar") + ("foo bar" + ("" + "foo bar")))
                    }
                    for fi26 := 0; fi26 < 8; fi26++ {
                        v23 -= v23
                        _ = fi26
                    }
                }
                _ = fi25
            }
            switch -9223372036854775808 {
            case 32767:
                for fi27 := 0; fi27 < 1; fi27++ {
                    var v28 int8 = -128
                    _ = fi27
                    _ = v28
                }
            case 1:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 0:
                v29 := GenS1{F0: "a"}
                _ = v29
            default:
                var v30 float64 = 0.0
                _ = v30
            }
        }
        _ = v24
    } else {
        var v31 string = "foo bar"
        var v32 int8 = (^(v1))
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        v23 -= v23
        if len(v3) > 0 {
            v3[0] = "a"
        }
        _ = fi33
    }
    var v34 uint16 = 0
    for fi35 := 0; fi35 < 11; fi35++ {
        for fi36 := 0; fi36 < 9; fi36++ {
            if false {
                if (int32(v34) < -1) {
                    switch ("foo bar" + (("foo bar" + "foo bar") + "")) {
                    case "a":
                        switch fi35 {
                        case 0:
                            v37 := GenS1{F0: "foo bar"}
                            _ = v37
                        case 128:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case -1:
                            if false {
                                if ((v34 << ((0) & 63)) <= (v34 | v34)) {
                                    var v38 uint32 = 833
                                    if true {
                                        v39 := GenS1{F0: "\n\t"}
                                        var v40 float64 = 3.14159
                                        _ = v39
                                        _ = v40
                                    } else {
                                        if true {
                                            if (true || (1 < 1)) {
                                                switch (fi35 >> ((42) & 63)) {
                                                case 349:
                                                    var v41 string = (("foo bar" + "hello") + "foo bar")
                                                    _ = v41
                                                case 32768:
                                                    var v42 int16 = 42
                                                    _ = v42
                                                case 42:
                                                    var v43 uint32 = v38
                                                    _ = v43
                                                default:
                                                    if false {
                                                        var v44 float32 = ((((0.1 - ((0.1 * 1e10) / 3.14159)) * (1e10 - 0.1)) - 0.1) / (1.5 / -2.25))
                                                        _ = v44
                                                    }
                                                }
                                                var v45 uint64 = (v23 ^ v23)
                                                v45 = v23
                                                _ = v45
                                            } else {
                                                var v46 int = -9223372036854775808
                                                var v47 bool = true
                                                var v48 uint = uint(v34)
                                                _ = v46
                                                _ = v47
                                                _ = v48
                                            }
                                            var v49 uint32 = (^(v38))
                                            var v50 int8 = 0
                                            if true {
                                                _ = len(v2)
                                            } else {
                                                _ = len(v3)
                                                v49 += v49
                                                v50 -= (^(v50))
                                                v23 -= uint64(v49)
                                            }
                                            _ = v49
                                            _ = v50
                                        } else {
                                            var v51 int = (^(fi35))
                                            if true {
                                                if false {
                                                    fmt.Println(fi36, v51, v38)
                                                    switch 9223372036854775807 {
                                                    case 1:
                                                        var v52 int = -1
                                                        _ = v52
                                                    case -9223372036854775808:
                                                        fmt.Println(fi35, v51, v38)
                                                    case 289:
                                                        fmt.Println(v51, v23, v1)
                                                        if false {
                                                            switch fi36 {
                                                            case 9223372036854775807:
                                                                var v53 uint64 = (^(v23))
                                                                _ = v53
                                                            default:
                                                                var v54 int64 = 1
                                                                _ = v54
                                                            }
                                                            if true {
                                                                if (v23 > v23) {
                                                                    v34 = (v34 << ((965) & 63))
                                                                    if true {
                                                                        if (1.5 >= 0.1) {
                                                                            _ = len(v2)
                                                                        } else {
                                                                            v55 := GenS1{F0: "a"}
                                                                            if false {
                                                                                var v56 uint64 = 44
                                                                                _ = v56
                                                                            }
                                                                            switch -401 {
                                                                            case 32768:
                                                                                v57 := GenS1{F0: (("a" + (("" + "") + "a")) + "a")}
                                                                                _ = v57
                                                                            case 9223372036854775807:
                                                                                v58 := GenS2{F0: "", F1: false, F2: -1}
                                                                                _ = v58
                                                                            }
                                                                            _ = v55
                                                                        }
                                                                        var v59 uint16 = v34
                                                                        if true {
                                                                            var v60 float32 = (((-2.25 / -2.25) * 3.14159) / 1.0)
                                                                            _ = v60
                                                                        }
                                                                        var v61 int8 = (v1 | (^(v1)))
                                                                        _ = v59
                                                                        _ = v61
                                                                    }
                                                                    v3 = append(v3, "foo bar")
                                                                }
                                                                v62 := GenS3{F0: (("hello" + (("a" + "") + "hello")) + "foo bar"), F1: 831}
                                                                if ((-2147483648 == 0) && false) {
                                                                    fmt.Println(fi36, v34)
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        if true {
                                                                            var v63 float64 = 0.1
                                                                            v38 = v38
                                                                            _ = v63
                                                                        }
                                                                    case "a":
                                                                        _ = len(v3)
                                                                    default:
                                                                        v34 += v34
                                                                    }
                                                                    v38 -= uint32(v23)
                                                                }
                                                                _ = v62
                                                            }
                                                            v64 := GenS4{F0: 715, F1: (v34 & 2), F2: -1}
                                                            _ = v64
                                                        }
                                                        var v65 int16 = -32768
                                                        _ = v65
                                                    }
                                                }
                                                switch "a" {
                                                case "":
                                                    v66 := GenS3{F0: "hello", F1: -234}
                                                    _ = v66
                                                case "hello":
                                                    v38 = (v38 ^ v38)
                                                case "\n\t":
                                                    var v67 uint8 = 76
                                                    switch "\n\t" {
                                                    case "":
                                                        if false {
                                                            fmt.Println(v38, v1, fi36)
                                                            if false {
                                                                fmt.Println(fi35, v1)
                                                                if (float32((float64(3.14159) * (3.14159 - 3.14159))) >= float32(((float64(float32(0.0)) / float64((-2.25 / (float32((-2.25 / ((((0.0 + 0.1) + 0.1) - 0.0) / ((1e10 / float64(((((1.5 - (1e10 * 0.1)) * float32(float64(1e10))) * -2.25) + ((-2.25 / (3.14159 - 0.1)) * (1.5 / 0.1))))) + -2.25)))) / 1e10)))) / -2.25))) {
                                                                    switch ((("a" + "hello") + "\n\t") + "a") {
                                                                    case "a":
                                                                        var v68 float64 = float64(((-2.25 / float32(((0.1 * 0.1) / (1.5 / 3.14159)))) - 3.14159))
                                                                        _ = v68
                                                                    case "foo bar":
                                                                        v23 = v23
                                                                    }
                                                                } else {
                                                                    v69 := GenS3{F0: ((("\n\t" + "a") + "a") + ("" + "\n\t")), F1: 42}
                                                                    if len(v2) > 0 {
                                                                        v2[0] = uint32(v34)
                                                                    }
                                                                    var v70 uint8 = v67
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                            } else {
                                                                var v71 bool = true
                                                                if (!(true || (623 == 1))) {
                                                                    switch (v51 << ((42) & 63)) {
                                                                    case 4294967295:
                                                                        var v72 int16 = 1
                                                                        _ = v72
                                                                    case 1:
                                                                        var v73 string = "\n\t"
                                                                        _ = v73
                                                                    case 250:
                                                                        var v74 uint16 = v34
                                                                        _ = v74
                                                                    }
                                                                    var v75 int32 = 1
                                                                    _ = v75
                                                                } else {
                                                                    switch "hello" {
                                                                    case "":
                                                                        fmt.Println(v71, v71)
                                                                    case "a":
                                                                        if ((v51 ^ fi35) >= 933) {
                                                                            fmt.Println(v51, v38)
                                                                        }
                                                                    case "\n\t":
                                                                        if len(v2) > 0 {
                                                                            v2[0] = 0
                                                                        }
                                                                    }
                                                                }
                                                                _ = v71
                                                            }
                                                        } else {
                                                            v23 = (v23 + uint64(v67))
                                                            var v76 bool = (0 != -1)
                                                            var v77 uint16 = (v34 / ((1) | 1))
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                    case "foo bar":
                                                        if len(v3) > 0 {
                                                            v3[0] = ""
                                                        }
                                                    default:
                                                        var v78 uint16 = (v34 - 42)
                                                        _ = v78
                                                    }
                                                    _ = v67
                                                }
                                            } else {
                                                v23 = 42
                                                var v79 string = "\n\t"
                                                _ = v79
                                            }
                                            _ = v51
                                        }
                                    }
                                    v3 = append(v3, ("a" + ("\n\t" + "foo bar")))
                                    _ = v38
                                } else {
                                    v1 = int8(v23)
                                    v23 = 1
                                    var v80 uint64 = (^(v23))
                                    _ = v80
                                }
                                if (680 != -2147483648) {
                                    v81 := GenS5{F0: 9223372036854775807}
                                    v82 := GenS6{F0: 583, F1: v23, F2: (fi36 + -1)}
                                    switch "\n\t" {
                                    case "hello":
                                        if true {
                                            v3 = append(v3, "\n\t")
                                            if len(v3) > 0 {
                                                v3[0] = (("hello" + "") + "hello")
                                            }
                                            var v83 string = (("" + ((("a" + ("hello" + "\n\t")) + ((("foo bar" + "a") + ("" + ("" + "hello"))) + "a")) + "a")) + "a")
                                            _ = v83
                                        } else {
                                            var v84 float64 = ((0.1 / (1.5 + -2.25)) - 0.0)
                                            v85 := GenS7{F0: (v34 / ((878) | 1))}
                                            _ = v84
                                            _ = v85
                                        }
                                    case "a":
                                        if (779 < 2147483647) {
                                            v1 += v1
                                            if true {
                                                switch fi36 {
                                                case -336:
                                                    if len(v2) > 0 {
                                                        v2[0] = 635
                                                    }
                                                case 128:
                                                    v86 := GenS8{F0: 0, F1: (v1 >> ((1) & 63))}
                                                    _ = v86
                                                case 644:
                                                    v23 = v23
                                                default:
                                                    var v87 bool = (-9223372036854775808 > (fi36 ^ -2147483648))
                                                    _ = v87
                                                }
                                                switch fi36 {
                                                case 42:
                                                    var v88 uint = 65535
                                                    _ = v88
                                                default:
                                                    fmt.Println(v34, v1, v34)
                                                }
                                            } else {
                                                _ = len(v3)
                                            }
                                        } else {
                                            var v89 int64 = 607
                                            _ = v89
                                        }
                                    default:
                                        fmt.Println(v23)
                                        var v90 float64 = 3.14159
                                        _ = v90
                                    }
                                    _ = v81
                                    _ = v82
                                } else {
                                    var v91 int = fi35
                                    v34 = (v34 | v34)
                                    _ = v91
                                }
                            }
                        }
                    default:
                        var v92 float64 = 0.1
                        _ = v92
                    }
                    switch (fi35 * fi36) {
                    case 42:
                        v93 := GenS6{F0: 2, F1: (v23 + 18446744073709551615), F2: fi36}
                        _ = v93
                    case -843:
                        var v94 uint64 = (v23 | (v23 - v23))
                        _ = v94
                    case 2147483647:
                        var v95 bool = (v23 > 1)
                        _ = v95
                    }
                } else {
                    if true {
                        switch fi35 {
                        case 1:
                            if true {
                                if (1 <= 1) {
                                    v96 := GenS2{F0: "\n\t", F1: false, F2: -9223372036854775808}
                                    if true {
                                        if (false && false) {
                                            if len(v2) > 0 {
                                                v2[0] = 976
                                            }
                                        } else {
                                            var v97 uint64 = v23
                                            switch fi35 {
                                            case -1:
                                                var v98 uint16 = (^(v34))
                                                _ = v98
                                            case -9223372036854775808:
                                                var v99 int16 = 337
                                                v100 := GenS9{F0: v99}
                                                if ((v99 / ((v99) | 1)) <= (v99 & v99)) {
                                                    v101 := GenS1{F0: ("foo bar" + "")}
                                                    var v102 int32 = 1
                                                    v102 = v102
                                                    _ = v101
                                                    _ = v102
                                                } else {
                                                    var v103 bool = (1 > -378)
                                                    if false {
                                                        v104 := GenS10{F0: int8(fi36)}
                                                        v99 += 42
                                                        v105 := GenS11{F0: v34, F1: v99, F2: (fi36 * (fi35 ^ fi36))}
                                                        _ = v104
                                                        _ = v105
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = "\n\t"
                                                    }
                                                    _ = v103
                                                }
                                                var v106 int8 = v1
                                                _ = v100
                                                _ = v106
                                                _ = v99
                                            }
                                            _ = v97
                                        }
                                        if true {
                                            if true {
                                                var v107 int64 = 9223372036854775807
                                                _ = v107
                                            } else {
                                                v108 := GenS12{F0: -9223372036854775808, F1: v23, F2: -9223372036854775808}
                                                if false {
                                                    if (65535 == 544) {
                                                        var v109 int16 = -491
                                                        if true {
                                                            if (((-2.25 + 3.14159) / (0.1 * 3.14159)) <= (float32(1e10) - 0.0)) {
                                                                v109 = 24
                                                            }
                                                            var v110 int32 = 42
                                                            var v111 int = 9223372036854775807
                                                            _ = v110
                                                            _ = v111
                                                        }
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v1 += (-(v1))
                                                        _ = v109
                                                    } else {
                                                        var v112 uint8 = 0
                                                        v113 := GenS13{F0: 65535, F1: 9223372036854775807}
                                                        _ = v112
                                                        _ = v113
                                                    }
                                                    var v114 bool = true
                                                    var v115 uint64 = v23
                                                    _ = v114
                                                    _ = v115
                                                } else {
                                                    var v116 int32 = 0
                                                    _ = v116
                                                }
                                                _ = v108
                                            }
                                            if (0 != (v34 << ((1) & 63))) {
                                                if len(v2) > 0 {
                                                    v2[0] = 2147483647
                                                }
                                                switch (fi36 ^ fi36) {
                                                case -9223372036854775808:
                                                    v34 += v34
                                                    if (false || false) {
                                                        var v117 uint32 = 1
                                                        var v118 uint32 = v117
                                                        _ = v117
                                                        _ = v118
                                                    }
                                                    if true {
                                                        var v119 int64 = 299
                                                        if true {
                                                            var v120 uint16 = 42
                                                            _ = len(v3)
                                                            _ = v120
                                                        } else {
                                                            v34 = (^(v34))
                                                            if (42 != v34) {
                                                                if (3.14159 > 1.5) {
                                                                    v119 -= v119
                                                                }
                                                            } else {
                                                                v1 = v1
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v1 -= (v1 << ((1) & 63))
                                                                v23 = (v23 % ((42) | 1))
                                                            }
                                                        }
                                                        _ = v119
                                                    }
                                                case 4294967295:
                                                    v23 += (v23 / ((v23) | 1))
                                                default:
                                                    switch 127 {
                                                    case 32768:
                                                        v23 = 18446744073709551615
                                                    case 2147483647:
                                                        var v121 bool = (9 > 245)
                                                        _ = v121
                                                    case 128:
                                                        v122 := GenS7{F0: v34}
                                                        _ = v122
                                                    default:
                                                        fmt.Println(v1)
                                                    }
                                                }
                                            }
                                        } else {
                                            v34 = uint16(v23)
                                        }
                                        fmt.Println(v1, v23)
                                    } else {
                                        var v123 uint8 = 2
                                        v124 := GenS2{F0: "foo bar", F1: (1000 >= 32767), F2: -805}
                                        _ = v123
                                        _ = v124
                                    }
                                    var v125 uint = 42
                                    _ = v125
                                    _ = v96
                                } else {
                                    fmt.Println(fi35, v23)
                                    fmt.Println(fi35, v1, v1)
                                }
                            } else {
                                _ = len(v3)
                                fmt.Println(v1)
                            }
                        case -474:
                            if len(v2) > 0 {
                                v2[0] = 42
                            }
                        case 670:
                            var v126 int8 = v1
                            _ = v126
                        }
                        v23 -= 18446744073709551615
                        var v127 float64 = 1.5
                        v2 = append(v2, 1)
                        _ = v127
                    } else {
                        fmt.Println(v34)
                    }
                    _ = len(v2)
                }
            } else {
                if ((v34 % (((v34 % ((1) | 1))) | 1)) != 453) {
                    v3 = append(v3, "\n\t")
                    fmt.Println(fi35, v23, fi35)
                } else {
                    var v128 int32 = 797
                    if (1.5 < 1.5) {
                        var v129 bool = false
                        var v130 uint64 = 816
                        switch (("hello" + ("" + "")) + ("a" + "foo bar")) {
                        case "a":
                            v23 += (v23 - v23)
                        case "hello":
                            var v131 int64 = 0
                            _ = v131
                        default:
                            if v129 {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        _ = v129
                        _ = v130
                    } else {
                        switch (fi35 << ((675) & 63)) {
                        case 1:
                            var v132 uint64 = v23
                            _ = v132
                        case 814:
                            v128 += (v128 + (v128 + -404))
                        }
                    }
                    _ = v128
                }
            }
            v23 = (v23 | (v23 ^ (v23 / ((v23) | 1))))
            _ = fi36
        }
        v34 = v34
        _ = fi35
    }
    if true {
        for fi133 := 0; fi133 < 2; fi133++ {
            var v134 int = fi133
            for fi135 := 0; fi135 < 20; fi135++ {
                if (2147483647 < 872) {
                    switch -492 {
                    case 32767:
                        fmt.Println(fi133)
                    case 128:
                        if false {
                            v134 = (fi135 / ((fi135) | 1))
                            fmt.Println(fi133, v134)
                            if true {
                                v3 = append(v3, (("hello" + "\n\t") + "hello"))
                                var v136 int16 = 0
                                v137 := GenS11{F0: v34, F1: (v136 / (((v136 | v136)) | 1)), F2: 298}
                                _ = v136
                                _ = v137
                            } else {
                                if (255 != 127) {
                                    var v138 int = (-(fi133))
                                    _ = v138
                                }
                                var v139 uint32 = 341
                                _ = v139
                            }
                        } else {
                            var v140 uint16 = 2
                            _ = v140
                        }
                    case 2147483647:
                        if (42 == 0) {
                            if len(v3) > 0 {
                                v3[0] = ("a" + ("foo bar" + "hello"))
                            }
                            var v141 int8 = v1
                            var v142 uint8 = 33
                            _ = v141
                            _ = v142
                        }
                    default:
                        fmt.Println(fi133, v23, v134)
                    }
                    v143 := GenS5{F0: -972}
                    _ = v143
                } else {
                    fmt.Println(v1)
                    var v144 int16 = 1000
                    fmt.Println(v144, v34, fi135)
                    _ = v144
                }
                fmt.Println(v134)
                _ = fi135
            }
            if len(v3) > 0 {
                v3[0] = ("hello" + (("hello" + "foo bar") + "foo bar"))
            }
            _ = fi133
            _ = v134
        }
        for fi145 := 0; fi145 < 6; fi145++ {
            if (((v23 - v23) == v23) && true) {
                fmt.Println(v23, v1)
            } else {
                if (1 > 1) {
                    switch "" {
                    case "foo bar":
                        switch fi145 {
                        case -1:
                            var v146 int32 = 2147483647
                            _ = v146
                        case 127:
                            var v147 uint8 = 196
                            _ = v147
                        case 9223372036854775807:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case "":
                        var v148 bool = ("foo bar" <= "foo bar")
                        _ = v148
                    case "a":
                        for fi149 := 0; fi149 < 13; fi149++ {
                            fmt.Println(v34)
                            v150 := GenS9{F0: -662}
                            var v151 float32 = (3.14159 / (((1e10 + 1.5) - (-2.25 * ((-2.25 * ((0.0 * ((3.14159 - (3.14159 - (-2.25 * 1.5))) - 1e10)) * (-2.25 - (1.5 / 1.5)))) + (1.5 * 1.5)))) * ((0.0 - 1e10) - (((((1.5 - 1e10) / ((1.5 * 3.14159) - (1.5 + -2.25))) + 0.1) - float32(0.0)) + (((float32(1.5) - 1.5) * 1e10) * (float32(0.0) + 1.5))))))
                            _ = fi149
                            _ = v150
                            _ = v151
                        }
                    default:
                        v23 = (v23 / (((v23 ^ 696)) | 1))
                    }
                }
                for fi152 := 0; fi152 < 18; fi152++ {
                    v23 -= v23
                    _ = fi152
                }
                var v153 float64 = -2.25
                _ = v153
            }
            _ = fi145
        }
        var v154 float32 = 1.5
        _ = v154
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v155 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v155:", v155)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v155
    _ = v2
    _ = v23
    _ = v3
    _ = v34
}
