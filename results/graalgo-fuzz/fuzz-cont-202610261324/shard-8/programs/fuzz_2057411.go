package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int16; F1 int64; F2 int64 }
type GenS6 struct { F0 int8; F1 int64; F2 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32; F1 uint; F2 int64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint; F1 bool; F2 uint64 }
type GenS12 struct { F0 int32; F1 uint }
type GenS13 struct { F0 uint; F1 uint64 }
type GenS14 struct { F0 uint64; F1 int8 }

func main() {
    var v1 int8 = 127
    var v2 uint64 = 1
    var v3 int32 = 0
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]string)
    if false {
        v1 += (v1 & (v1 ^ (v1 & (v1 & -128))))
        v6 := []int{-1, 770, 128}
        _ = v6
    }
    if true {
        v7 := GenS1{F0: 9223372036854775807, F1: 32767}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int64 = -480
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 uint32 = 1
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: -1, F1: 1}
        _ = v12
    }
    var v13 int32 = -1
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := GenS2{F0: 1, F1: 42, F2: 41}
        _ = fi14
        _ = v15
    }
    _ = len(v5)
    if true {
        fmt.Println(v3)
        var v16 float32 = (float32((0.1 - (0.1 + (1.5 + -2.25)))) + float32(-2.25))
        v17 := GenS1{F0: 42, F1: 3}
        _ = v16
        _ = v17
    } else {
        if false {
            if true {
                if (0.0 > (float64((0.0 / 1.5)) * 1e10)) {
                    for fi18 := 0; fi18 < 7; fi18++ {
                        var v19 uint64 = (v2 >> ((255) & 63))
                        var v20 int = (fi18 << ((2) & 63))
                        var v21 int16 = 1000
                        _ = fi18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    var v22 uint64 = v2
                    v23 := GenS1{F0: 42, F1: 1}
                    _ = v22
                    _ = v23
                }
            } else {
                for fi24 := 0; fi24 < 18; fi24++ {
                    v2 = v2
                    _ = len(v4)
                    _ = fi24
                }
            }
            if false {
                var v25 float32 = 1.5
                if (0 > (v13 * 342)) {
                    for fi26 := 0; fi26 < 3; fi26++ {
                        for fi27 := 0; fi27 < 10; fi27++ {
                            switch "a" {
                            case "foo bar":
                                var v28 int32 = v13
                                _ = v28
                            }
                            v29 := []int{(fi27 & fi27), fi26}
                            var v30 uint64 = v2
                            _ = fi27
                            _ = v29
                            _ = v30
                        }
                        v1 = (v1 / ((-99) | 1))
                        _ = fi26
                    }
                    v31 := GenS3{F0: -9223372036854775808}
                    var v32 int16 = 1
                    _ = v31
                    _ = v32
                } else {
                    for fi33 := 0; fi33 < 5; fi33++ {
                        if true {
                            v2 = (v2 | (v2 / ((84) | 1)))
                            delete(v4, 2147483647)
                            var v34 string = "hello"
                            _ = v34
                        }
                        if false {
                            for fi35 := 0; fi35 < 1; fi35++ {
                                if ((!false) && true) {
                                    fmt.Println(v2, v13)
                                } else {
                                    v2 = v2
                                }
                                fmt.Println(v2)
                                var v36 int8 = (v1 % ((v1) | 1))
                                v25 += v25
                                fmt.Println(v13, v36, v2)
                                _ = fi35
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 10; fi37++ {
                                v38 := []int{-2147483648}
                                if false {
                                    if (32767 > 1000) {
                                        if false {
                                            _ = len(v5)
                                        } else {
                                            switch "a" {
                                            case "":
                                                if true {
                                                    v39 := GenS2{F0: (^(v2)), F1: -706, F2: 255}
                                                    if false {
                                                        var v40 uint64 = v2
                                                        var v41 int = fi33
                                                        if (223 > 42) {
                                                            var v42 int16 = 1
                                                            v43 := GenS4{F0: 2, F1: 42}
                                                            _ = v42
                                                            _ = v43
                                                        } else {
                                                            var v44 uint16 = 558
                                                            var v45 int16 = 42
                                                            v41 = -234
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                        if ((2147483648 == 961) && true) {
                                                            var v46 bool = (2 < 4294967295)
                                                            var v47 uint = 255
                                                            switch "\n\t" {
                                                            case "":
                                                                var v48 int = 9223372036854775807
                                                                _ = v48
                                                            default:
                                                                switch int(v1) {
                                                                case -190:
                                                                    v49 := GenS2{F0: (v2 << ((v47) & 63)), F1: -1, F2: 729}
                                                                    _ = v49
                                                                case -9223372036854775808:
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            var v50 float64 = (((float64((v25 - v25)) / 3.14159) / float64((1e10 * float32(float64(float32(1e10)))))) / 1.0)
                                                                            _ = v50
                                                                        case "":
                                                                            v41 = (v41 | (v41 | fi33))
                                                                        default:
                                                                            v41 -= v41
                                                                        }
                                                                    case "hello":
                                                                        fmt.Println(v1, fi37)
                                                                    default:
                                                                        v2 -= (^(v2))
                                                                    }
                                                                case 42:
                                                                    switch fi33 {
                                                                    case 4294967295:
                                                                        v2 += v2
                                                                    default:
                                                                        fmt.Println(v46, v13, v1)
                                                                    }
                                                                }
                                                            }
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                        _ = v40
                                                        _ = v41
                                                    } else {
                                                        switch (("\n\t" + "") + "a") {
                                                        case "":
                                                            if true {
                                                                var v51 uint = 9
                                                                var v52 uint = (v51 ^ (^(v51)))
                                                                v4[814] = 65535
                                                                if len(v38) > 0 {
                                                                    v38[0] = fi33
                                                                }
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                        }
                                                        v53 := GenS4{F0: 612, F1: 42}
                                                        v54 := v4[71]
                                                        v55 := GenS5{F0: 0, F1: 1, F2: -9223372036854775808}
                                                        v56 := GenS6{F0: v1, F1: -356, F2: v2}
                                                        _ = v53
                                                        _ = v54
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                    _ = v39
                                                }
                                            case "foo bar":
                                                var v57 int16 = -32768
                                                _ = v57
                                            case "a":
                                                v5[2] = ""
                                            default:
                                                var v58 int8 = -15
                                                _ = v58
                                            }
                                        }
                                    }
                                    _ = len(v38)
                                    v1 = 0
                                }
                                fmt.Println(fi37, v25)
                                _ = fi37
                                _ = v38
                            }
                        } else {
                            fmt.Println(fi33, fi33)
                            if false {
                                if (v13 > 0) {
                                    if (1.5 >= 0.0) {
                                        fmt.Println(v13, fi33, v13)
                                    } else {
                                        var v59 uint = 1
                                        _ = v59
                                    }
                                }
                                v60 := []int{fi33, 4294967295, 32768, fi33}
                                _ = v60
                            } else {
                                fmt.Println(v25)
                                v61 := GenS4{F0: 65535, F1: 143}
                                _ = v61
                            }
                            v62 := GenS1{F0: 993, F1: 1000}
                            _ = v62
                        }
                        _ = fi33
                    }
                }
                _ = v25
            }
            for fi63 := 0; fi63 < 6; fi63++ {
                v3 = (v3 + (-(v3)))
                _ = fi63
            }
        } else {
            for fi64 := 0; fi64 < 17; fi64++ {
                _ = len(v4)
                _ = fi64
            }
            v65 := GenS5{F0: -889, F1: -1, F2: -9223372036854775808}
            for fi66 := 0; fi66 < 12; fi66++ {
                if true {
                    v2 += v2
                    fmt.Println(v2, v2)
                    if (v3 <= 1) {
                        switch -645 {
                        case -9223372036854775808:
                            fmt.Println(v3, v3)
                        case 613:
                            v67 := GenS7{F0: 127}
                            _ = v67
                        case 127:
                            v68 := GenS7{F0: 42}
                            _ = v68
                        }
                    } else {
                        var v69 uint = 161
                        _ = v69
                    }
                }
                if (3.14159 > 0.0) {
                    switch (fi66 % ((fi66) | 1)) {
                    case 2147483647:
                        var v70 int32 = 42
                        _ = v70
                    case -2147483648:
                        v13 -= (v3 << ((0) & 63))
                    }
                    switch "a" {
                    case "":
                        switch fi66 {
                        case 620:
                            if true {
                                for fi71 := 0; fi71 < 12; fi71++ {
                                    v1 -= (v1 << ((255) & 63))
                                    var v72 bool = (1 <= 127)
                                    var v73 int = fi66
                                    _ = fi71
                                    _ = v72
                                    _ = v73
                                }
                                if (127 == 194) {
                                    var v74 uint64 = v2
                                    fmt.Println(v1, v2)
                                    _ = v74
                                }
                            } else {
                                var v75 float32 = 0.1
                                if true {
                                    var v76 int = (fi66 | (fi66 << ((669) & 63)))
                                    _ = v76
                                } else {
                                    var v77 float64 = 3.14159
                                    if (2 != 1) {
                                        var v78 uint8 = 128
                                        _ = v78
                                    }
                                    _ = v77
                                }
                                var v79 int32 = 109
                                _ = v75
                                _ = v79
                            }
                        case -2147483648:
                            delete(v4, 0)
                        case 128:
                            var v80 bool = true
                            _ = v80
                        default:
                            if true {
                                var v81 int8 = -128
                                v82 := GenS3{F0: 42}
                                _ = v81
                                _ = v82
                            } else {
                                var v83 float64 = 1e10
                                var v84 int = int(v2)
                                _ = v83
                                _ = v84
                            }
                        }
                    case "foo bar":
                        v4[2] = 42
                    case "a":
                        switch (("hello" + "hello") + "") {
                        case "hello":
                            fmt.Println(v1)
                        case "\n\t":
                            v85 := v5[65535]
                            _ = v85
                        default:
                            for fi86 := 0; fi86 < 20; fi86++ {
                                var v87 uint = 42
                                switch (fi66 << ((18446744073709551615) & 63)) {
                                case 88:
                                    var v88 uint = 0
                                    _ = v88
                                case 2147483647:
                                    var v89 uint16 = 0
                                    _ = v89
                                }
                                v90 := GenS8{F0: 2, F1: 2, F2: -967}
                                if true {
                                    v13 -= v13
                                    if (434 >= -686) {
                                        if true {
                                            var v91 int64 = -512
                                            v92 := GenS8{F0: 2147483648, F1: (v87 / (((v87 << (((v87 ^ v87)) & 63))) | 1)), F2: (v91 * (v91 + (v91 * v91)))}
                                            v93 := GenS2{F0: v2, F1: (v91 / ((v91) | 1)), F2: v87}
                                            var v94 string = ("a" + "a")
                                            _ = v91
                                            _ = v92
                                            _ = v93
                                            _ = v94
                                        } else {
                                            v95 := GenS7{F0: 151}
                                            _ = v95
                                        }
                                        v96 := GenS9{F0: -1}
                                        _ = v96
                                    } else {
                                        var v97 int32 = v13
                                        if true {
                                            switch (fi86 + (fi66 ^ fi86)) {
                                            case 0:
                                                if (!true) {
                                                    fmt.Println(v97, v87)
                                                    switch 4294967295 {
                                                    case 42:
                                                        var v98 uint32 = 0
                                                        if false {
                                                            var v99 float32 = -2.25
                                                            if true {
                                                                var v100 uint32 = v98
                                                                _ = v100
                                                            } else {
                                                                var v101 uint32 = v98
                                                                _ = v101
                                                            }
                                                            var v102 float32 = v99
                                                            if (!true) {
                                                                var v103 uint64 = (v2 % ((v2) | 1))
                                                                v4[(^(v98))] = (^(v98))
                                                                _ = v103
                                                            }
                                                            v5[v98] = "\n\t"
                                                            _ = v102
                                                            _ = v99
                                                        } else {
                                                            switch "hello" {
                                                            case "":
                                                                var v104 int32 = 0
                                                                _ = v104
                                                            case "\n\t":
                                                                v105 := GenS10{F0: 0}
                                                                _ = v105
                                                            case "hello":
                                                                v106 := GenS11{F0: v87, F1: (2147483647 == (v98 | v98)), F2: (^(v2))}
                                                                _ = v106
                                                            }
                                                            fmt.Println(v3, v1, v2)
                                                            v107 := GenS1{F0: -1, F1: 0}
                                                            _ = v107
                                                        }
                                                        _ = v98
                                                    }
                                                    switch fi86 {
                                                    case 128:
                                                        switch (("foo bar" + ("\n\t" + ("a" + "foo bar"))) + "") {
                                                        case "foo bar":
                                                            _ = len(v5)
                                                        }
                                                    }
                                                }
                                            case -800:
                                                if (false && (("a" >= ("a" + "\n\t")) && ((fi66 >= 35) || true))) {
                                                    fmt.Println(v87, v97, v87)
                                                    v13 += (v97 & int32(v1))
                                                    v108 := GenS12{F0: (v13 % (((v13 + (v3 * (^(v13))))) | 1)), F1: v87}
                                                    _ = v108
                                                }
                                            case -378:
                                                v2 = (v2 / (((v2 / ((v2) | 1))) | 1))
                                            default:
                                                var v109 uint16 = uint16(v2)
                                                _ = v109
                                            }
                                            if (v87 <= 2) {
                                                v13 = (-(v97))
                                                v110 := GenS13{F0: (v87 & v87), F1: v2}
                                                if (false || false) {
                                                    if true {
                                                        var v111 bool = false
                                                        if false {
                                                            v5[934] = "hello"
                                                        }
                                                        fmt.Println(v2, v13)
                                                        _ = v111
                                                    } else {
                                                        var v112 bool = (33 < 128)
                                                        v4[839] = 4294967295
                                                        _ = v112
                                                    }
                                                    v113 := GenS4{F0: (v87 & 864), F1: 2}
                                                    _ = v113
                                                } else {
                                                    if (v1 != v1) {
                                                        if (255 != 2) {
                                                            v2 -= (v2 % (((v2 >> ((4294967295) & 63))) | 1))
                                                            v114 := GenS13{F0: v87, F1: (v2 ^ (v2 % (((v2 >> ((v87) & 63))) | 1)))}
                                                            _ = v114
                                                        }
                                                        var v115 float64 = 1e10
                                                        v116 := GenS11{F0: v87, F1: true, F2: 1}
                                                        _ = v115
                                                        _ = v116
                                                    } else {
                                                        v117 := GenS8{F0: uint32(v2), F1: v87, F2: -623}
                                                        v13 -= (v13 | int32(v87))
                                                        v13 = -1
                                                        _ = v117
                                                    }
                                                    fmt.Println(v97, v13, v1)
                                                    v118 := []int{(fi66 >> ((18446744073709551615) & 63)), (-(fi86))}
                                                    _ = v118
                                                }
                                                _ = v110
                                            } else {
                                                v119 := []int{fi86, fi66}
                                                _ = v119
                                            }
                                            fmt.Println(v3)
                                        }
                                        _ = v97
                                    }
                                } else {
                                    v3 = v13
                                    switch ("foo bar" + "a") {
                                    case "\n\t":
                                        var v120 int64 = 516
                                        var v121 float64 = 0.1
                                        _ = v120
                                        _ = v121
                                    case "hello":
                                        v13 += 1
                                    }
                                }
                                var v122 float64 = ((0.0 / float64(0.1)) * 1e10)
                                _ = fi86
                                _ = v122
                                _ = v87
                                _ = v90
                            }
                        }
                    default:
                        for fi123 := 0; fi123 < 12; fi123++ {
                            v1 -= -69
                            _ = fi123
                        }
                    }
                    v124 := GenS14{F0: (v2 - v2), F1: 127}
                    _ = v124
                }
                _ = fi66
            }
            _ = v65
        }
    }
    var v125 int = -1
    fmt.Println(v125, v13)
    var v126 int64 = 0
    v2 = v2
    var v127 int16 = -1
    v2 -= v2
    v128 := v5[19]
    fmt.Println("v1:", v1)
    fmt.Println("v125:", v125)
    fmt.Println("v126:", v126)
    fmt.Println("v127:", v127)
    fmt.Println("v128:", v128)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v125
    _ = v126
    _ = v127
    _ = v128
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
