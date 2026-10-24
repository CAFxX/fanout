package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS5 struct { F0 int; F1 uint; F2 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int64; F1 string; F2 string }
type GenS8 struct { F0 int64; F1 string; F2 int64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint8; F1 int }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint16; F1 int64 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int16; F1 string; F2 uint64 }
func genFn0(p0 string, p1 int16) uint8 {
    v6 := GenS1{F0: 176}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v6
    return 2
}
func genFn1(p0 int) int16 {
    if true {
        switch "hello" {
        case "\n\t":
            switch ("hello" + "") {
            case "foo bar":
                var v7 int = (p0 / ((9223372036854775807) | 1))
                _ = v7
            case "a":
                p0 = -2147483648
            }
        case "a":
            var v8 uint32 = 0
            _ = v8
        case "hello":
            var v9 uint = 2
            _ = v9
        }
        var v10 string = "a"
        if ((false && false) || true) {
            if true {
                switch p0 {
                case 517:
                    if false {
                        var v11 uint32 = 2147483648
                        v12 := []int{}
                        _ = len(v12)
                        _ = v11
                        _ = v12
                    }
                default:
                    if (3.14159 < 1e10) {
                        var v13 uint8 = 218
                        switch p0 {
                        case 32768:
                            v14 := GenS2{F0: (0 <= 18446744073709551615), F1: "\n\t"}
                            _ = v14
                        case -2147483648:
                            for fi15 := 0; fi15 < 6; fi15++ {
                                for fi16 := 0; fi16 < 3; fi16++ {
                                    switch -9223372036854775808 {
                                    case -9223372036854775808:
                                        v17 := GenS1{F0: (^(v13))}
                                        _ = v17
                                    case -2147483648:
                                        var v18 float64 = 3.14159
                                        _ = v18
                                    default:
                                        var v19 uint32 = 424
                                        _ = v19
                                    }
                                    v20 := GenS2{F0: true, F1: v10}
                                    p0 -= (^(fi16))
                                    _ = fi16
                                    _ = v20
                                }
                                _ = fi15
                            }
                        case -801:
                            v21 := []int{}
                            _ = v21
                        default:
                            if (!false) {
                                for fi22 := 0; fi22 < 12; fi22++ {
                                    for fi23 := 0; fi23 < 17; fi23++ {
                                        v10 = ("a" + v10)
                                        v24 := []int{128, (fi23 ^ (fi22 + 2147483647))}
                                        _ = fi23
                                        _ = v24
                                    }
                                    v25 := []int{(fi22 >> ((945) & 63)), (-(fi22)), (p0 >> ((4294967295) & 63))}
                                    _ = fi22
                                    _ = v25
                                }
                            } else {
                                v26 := []int{499, (p0 ^ 0), 128}
                                var v27 int8 = 60
                                _ = v26
                                _ = v27
                            }
                        }
                        v28 := GenS3{F0: 2147483647}
                        _ = v13
                        _ = v28
                    } else {
                        p0 -= p0
                    }
                }
                for fi29 := 0; fi29 < 13; fi29++ {
                    switch v10 {
                    case "":
                        v30 := GenS1{F0: 1}
                        _ = v30
                    }
                    fmt.Println(p0)
                    _ = fi29
                }
                v10 = v10
            } else {
                v31 := []int{(p0 % ((-9223372036854775808) | 1)), p0}
                _ = v31
            }
        }
        _ = v10
    }
    fmt.Println(p0)
    _ = p0
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 string = ""
    var v3 string = v2
    v4 := []int{782, 127, 104, 128}
    v5 := []int{4294967295, 128, -212}
    if (42 != -194) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if len(v5) > 0 {
            v5[0] = 32768
        }
    } else {
        fmt.Println(v2, v2)
    }
    if len(v4) > 0 {
        v4[0] = 127
    }
    var v32 uint16 = 206
    if (0.1 != float64((0.0 - (((((-2.25 - float32(1.5)) - (0.0 + (0.0 - -2.25))) + (float32(1.5) - 1.5)) * (1e10 + float32(0.1))) - ((-2.25 - (3.14159 + 1e10)) * 1e10))))) {
        v2 = ((v2 + v3) + "")
        for fi33 := 0; fi33 < 20; fi33++ {
            switch fi33 {
            case 128:
                v2 = ("hello" + ("" + "\n\t"))
            }
            v34 := GenS3{F0: 42}
            _ = fi33
            _ = v34
        }
    } else {
        switch v3 {
        case "a":
            if (-462 <= 0) {
                var v35 int16 = 32767
                switch 32768 {
                case 4294967295:
                    v36 := GenS3{F0: -2147483648}
                    _ = v36
                default:
                    var v37 uint = 0
                    _ = v37
                }
                _ = v35
            }
        default:
            var v38 int = 4294967295
            _ = v38
        }
        fmt.Println(v3, v32, v32)
    }
    v32 = v32
    v32 = v32
    switch ("a" + "hello") {
    case "hello":
        var v39 uint32 = 2147483648
        _ = v39
    case "a":
        if (65535 == 2147483648) {
            switch v2 {
            case "hello":
                v2 = v2
            case "foo bar":
                if (-86 <= 849) {
                    var v40 int16 = 32767
                    fmt.Println(v2, v2)
                    v41 := GenS2{F0: true, F1: "hello"}
                    _ = v40
                    _ = v41
                }
            }
        } else {
            if ((18446744073709551615 == 4294967295) && false) {
                var v42 uint = 528
                v43 := GenS2{F0: (1 < 4294967295), F1: "hello"}
                v2 = v3
                _ = v42
                _ = v43
            }
            var v44 int8 = 42
            _ = v44
        }
    case "foo bar":
        if true {
            var v45 uint = 65535
            switch 4294967295 {
            case 32768:
                if ((605 != -990) && false) {
                    for fi46 := 0; fi46 < 4; fi46++ {
                        v3 = ""
                        if true {
                            if false {
                                var v47 uint64 = 0
                                var v48 uint8 = uint8(v45)
                                _ = v47
                                _ = v48
                            }
                            if (!(true && ((v32 - (v32 ^ 404)) >= (v32 | 2)))) {
                                v49 := GenS2{F0: (true && (!true)), F1: "foo bar"}
                                v5 = append(v5, (fi46 % (((fi46 ^ (fi46 % (((-(fi46))) | 1)))) | 1)))
                                _ = v49
                            }
                        } else {
                            for fi50 := 0; fi50 < 2; fi50++ {
                                if true {
                                    var v51 uint8 = 2
                                    v3 = "a"
                                    _ = v51
                                }
                                if (!true) {
                                    if (!false) {
                                        fmt.Println(fi46, fi50)
                                    }
                                    if (2147483647 < -2147483648) {
                                        _ = len(v1)
                                        v32 -= v32
                                        if false {
                                            v32 += v32
                                            v3 = "\n\t"
                                            switch 2147483647 {
                                            case 32767:
                                                if false {
                                                    switch ("" + "\n\t") {
                                                    case "a":
                                                        var v52 uint64 = 2
                                                        _ = v52
                                                    case "\n\t":
                                                        if (2147483647 <= 259) {
                                                            v45 -= (v45 % ((v45) | 1))
                                                            var v53 bool = (((1.5 - (-2.25 * 0.1)) < (0.1 + 3.14159)) && (255 < 244))
                                                            _ = v53
                                                        } else {
                                                            if (!true) {
                                                                var v54 int = (fi50 * fi50)
                                                                var v55 bool = (((v3 + ("hello" + (v2 + "foo bar"))) == ((v3 + v2) + v2)) && ((1.5 >= 1.5) || ((false && true) || false)))
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                            var v56 int32 = 0
                                                            _ = v56
                                                        }
                                                    }
                                                    var v57 int8 = -128
                                                    _ = v57
                                                }
                                            case 4294967295:
                                                _ = len(v4)
                                            }
                                        } else {
                                            fmt.Println(v3)
                                            var v58 uint64 = 635
                                            var v59 int16 = 32767
                                            switch ("a" + "a") {
                                            case "a":
                                                if false {
                                                    v60 := GenS1{F0: 42}
                                                    var v61 int = -490
                                                    var v62 int = v61
                                                    _ = v60
                                                    _ = v61
                                                    _ = v62
                                                }
                                            case "hello":
                                                v45 = v45
                                            }
                                            _ = v58
                                            _ = v59
                                        }
                                    }
                                    if false {
                                        var v63 float64 = (float64((0.1 / (1e10 / -2.25))) * float64(1e10))
                                        _ = v63
                                    }
                                }
                                v32 = (v32 >> ((v45) & 63))
                                _ = fi50
                            }
                            for fi64 := 0; fi64 < 3; fi64++ {
                                var v65 int16 = 1000
                                _ = fi64
                                _ = v65
                            }
                        }
                        switch fi46 {
                        case 42:
                            v45 = (v45 + v45)
                        case 377:
                            var v66 float32 = 0.0
                            _ = v66
                        case 127:
                            if ((v2 + (((v3 + "") + (v3 + (("foo bar" + v2) + v3))) + v2)) <= "a") {
                                if (618 == 2147483647) {
                                    if (1.5 != (((0.0 - (0.0 * (-2.25 / -2.25))) / 1.0) / ((3.14159 / 1.5) * -2.25))) {
                                        if (v32 <= (v32 | v32)) {
                                            for fi67 := 0; fi67 < 8; fi67++ {
                                                var v68 uint16 = 224
                                                v69 := GenS3{F0: 42}
                                                _ = fi67
                                                _ = v68
                                                _ = v69
                                            }
                                            var v70 int = fi46
                                            _ = v70
                                        }
                                        if (true || true) {
                                            var v71 uint = (v45 + (v45 / ((v45) | 1)))
                                            v2 = v3
                                            _ = v71
                                        } else {
                                            switch v2 {
                                            case "":
                                                v72 := GenS4{F0: 42, F1: v32, F2: 0}
                                                _ = v72
                                            case "a":
                                                fmt.Println(v32, fi46, fi46)
                                            }
                                        }
                                    } else {
                                        var v73 int64 = 1
                                        _ = v73
                                    }
                                    v74 := GenS1{F0: 42}
                                    var v75 uint32 = 2147483647
                                    _ = v74
                                    _ = v75
                                }
                            }
                        default:
                            fmt.Println(v32, fi46, fi46)
                        }
                        _ = fi46
                    }
                    var v76 float32 = 0.0
                    var v77 int64 = 1
                    _ = v76
                    _ = v77
                } else {
                    for fi78 := 0; fi78 < 4; fi78++ {
                        var v79 uint8 = 255
                        v80 := GenS5{F0: (fi78 >> (((v45 - (v45 / ((4294967295) | 1)))) & 63)), F1: 65535, F2: 72}
                        _ = fi78
                        _ = v79
                        _ = v80
                    }
                    if true {
                        if ((v32 >> ((v45) & 63)) != (v32 + 65535)) {
                            var v81 uint16 = v32
                            for fi82 := 0; fi82 < 12; fi82++ {
                                v45 -= (v45 << (((v45 * (v45 / ((v45) | 1)))) & 63))
                                for fi83 := 0; fi83 < 17; fi83++ {
                                    v1[4294967295] = (fi83 - 4294967295)
                                    v84 := GenS6{F0: 42}
                                    _ = fi83
                                    _ = v84
                                }
                                _ = fi82
                            }
                            _ = v81
                        }
                        if false {
                            switch -2147483648 {
                            case 128:
                                if false {
                                    v85 := GenS4{F0: 0, F1: v32, F2: -1}
                                    var v86 int32 = -1
                                    _ = v85
                                    _ = v86
                                } else {
                                    v32 -= v32
                                }
                            case 746:
                                var v87 float32 = (1.5 + 3.14159)
                                _ = v87
                            }
                            for fi88 := 0; fi88 < 6; fi88++ {
                                if len(v4) > 0 {
                                    v4[0] = fi88
                                }
                                _ = fi88
                            }
                            for fi89 := 0; fi89 < 15; fi89++ {
                                v90 := GenS2{F0: false, F1: v2}
                                for fi91 := 0; fi91 < 17; fi91++ {
                                    v2 = ("\n\t" + v3)
                                    _ = fi91
                                }
                                _ = fi89
                                _ = v90
                            }
                        }
                        var v92 uint8 = uint8(v32)
                        _ = v92
                    }
                    fmt.Println(v45)
                }
            }
            _ = v45
        } else {
            if true {
                if ((!(false || ((!true) || (!(!false))))) && false) {
                    var v93 int64 = -1
                    _ = v93
                } else {
                    var v94 string = v2
                    v3 = (v2 + v2)
                    var v95 bool = false
                    _ = v94
                    _ = v95
                }
                var v96 int16 = genFn1(9223372036854775807)
                var v97 bool = (!true)
                _ = v96
                _ = v97
            } else {
                v32 -= v32
            }
            switch 128 {
            case 0:
                switch "a" {
                case "foo bar":
                    v32 = 1
                case "hello":
                    switch 127 {
                    case 1:
                        var v98 uint64 = 2
                        _ = v98
                    default:
                        if false {
                            if len(v5) > 0 {
                                v5[0] = -119
                            }
                        }
                    }
                case "\n\t":
                    switch -626 {
                    case 838:
                        v99 := GenS7{F0: -9223372036854775808, F1: v2, F2: v3}
                        _ = v99
                    case 9223372036854775807:
                        v2 = (v3 + "\n\t")
                    case 0:
                        _ = len(v1)
                    default:
                        if false {
                            for fi100 := 0; fi100 < 15; fi100++ {
                                if (1 != 0) {
                                    if true {
                                        fmt.Println(fi100, v2, v32)
                                        fmt.Println(v2, v2)
                                        if (32767 != genFn1(-2147483648)) {
                                            if true {
                                                v2 = "a"
                                                var v101 float64 = (0.0 + float64(((0.1 - 1e10) * 0.0)))
                                                v32 = (v32 | (v32 | 866))
                                                _ = v101
                                            }
                                            for fi102 := 0; fi102 < 9; fi102++ {
                                                v3 = "a"
                                                v32 = v32
                                                if (!((fi100 / (((fi102 << ((18446744073709551615) & 63))) | 1)) < (fi102 - (fi102 << ((18446744073709551615) & 63))))) {
                                                    switch fi100 {
                                                    case 615:
                                                        v103 := GenS8{F0: 0, F1: v3, F2: 743}
                                                        _ = v103
                                                    case -595:
                                                        fmt.Println(fi100)
                                                    case 42:
                                                        v32 = 58
                                                        if true {
                                                            switch ("hello" + v3) {
                                                            case "foo bar":
                                                                switch fi102 {
                                                                case 691:
                                                                    switch v2 {
                                                                    case "foo bar":
                                                                        fmt.Println(v2, v2)
                                                                    case "\n\t":
                                                                        v2 = v2
                                                                    }
                                                                case 32768:
                                                                    switch v2 {
                                                                    case "\n\t":
                                                                        var v104 uint = 1
                                                                        _ = v104
                                                                    case "foo bar":
                                                                        v105 := v1[fi102]
                                                                        _ = v105
                                                                    default:
                                                                        if true {
                                                                            v32 += (v32 * 103)
                                                                            switch (fi100 - fi100) {
                                                                            case 127:
                                                                                if (0.0 >= -2.25) {
                                                                                    if ((!(-9223372036854775808 == fi102)) || ((v3 + v3) <= ((v3 + (("hello" + "foo bar") + ("a" + "\n\t"))) + ""))) {
                                                                                        var v106 uint = 18446744073709551615
                                                                                        _ = v106
                                                                                    } else {
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = (-(fi102))
                                                                                        }
                                                                                        if false {
                                                                                            switch fi102 {
                                                                                            case 32768:
                                                                                                if false {
                                                                                                    if false {
                                                                                                        var v107 string = "foo bar"
                                                                                                        if (337 > 633) {
                                                                                                            v32 = (v32 ^ 2)
                                                                                                            if false {
                                                                                                                if false {
                                                                                                                    var v108 uint8 = 128
                                                                                                                    _ = v108
                                                                                                                } else {
                                                                                                                    if ((1e10 - (1.5 * 0.1)) >= (((0.0 + 3.14159) / ((float64(1.5) / 1.0) / (((0.1 * 1.5) - -2.25) / float64((0.1 / 0.1))))) + 1e10)) {
                                                                                                                        _ = len(v1)
                                                                                                                        if len(v4) > 0 {
                                                                                                                            v4[0] = fi100
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        v32 = 255
                                                                                                                        var v109 uint16 = v32
                                                                                                                        _ = v109
                                                                                                                    }
                                                                                                                    v110 := GenS9{F0: 251}
                                                                                                                    _ = v110
                                                                                                                }
                                                                                                                switch fi102 {
                                                                                                                case 2147483647:
                                                                                                                    fmt.Println(fi100, v3, v2)
                                                                                                                case 544:
                                                                                                                    var v111 uint8 = 233
                                                                                                                    _ = v111
                                                                                                                case -2147483648:
                                                                                                                    delete(v1, -326)
                                                                                                                }
                                                                                                                var v112 uint = 0
                                                                                                                _ = v112
                                                                                                            } else {
                                                                                                                if false {
                                                                                                                    v2 = (("foo bar" + "\n\t") + v107)
                                                                                                                    switch fi100 {
                                                                                                                    case 431:
                                                                                                                        v32 += (v32 * (v32 >> ((18446744073709551615) & 63)))
                                                                                                                    case -9223372036854775808:
                                                                                                                        if (v3 > v2) {
                                                                                                                            var v113 int32 = -1
                                                                                                                            _ = len(v1)
                                                                                                                            fmt.Println(v113, v107)
                                                                                                                            _ = v113
                                                                                                                        } else {
                                                                                                                            if (true && (!((1.5 >= 1e10) || ((float64(1.5) / 0.1) == 3.14159)))) {
                                                                                                                                switch "foo bar" {
                                                                                                                                case "\n\t":
                                                                                                                                    fmt.Println(v107, v2, v2)
                                                                                                                                case "":
                                                                                                                                    v107 = ((v2 + "hello") + "foo bar")
                                                                                                                                case "a":
                                                                                                                                    var v114 string = v3
                                                                                                                                    _ = v114
                                                                                                                                default:
                                                                                                                                    var v115 int32 = 2147483647
                                                                                                                                    _ = v115
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                    case 128:
                                                                                                                        var v116 int16 = 0
                                                                                                                        _ = v116
                                                                                                                    }
                                                                                                                    v117 := GenS1{F0: 2}
                                                                                                                    _ = v117
                                                                                                                } else {
                                                                                                                    if (0 > 1) {
                                                                                                                        v3 = (v107 + v3)
                                                                                                                        if true {
                                                                                                                            var v118 int = 32767
                                                                                                                            v119 := GenS10{F0: 127, F1: (^(fi102))}
                                                                                                                            _ = v118
                                                                                                                            _ = v119
                                                                                                                        } else {
                                                                                                                            var v120 float32 = 0.0
                                                                                                                            _ = v120
                                                                                                                        }
                                                                                                                        v5 = append(v5, (fi102 ^ -2147483648))
                                                                                                                    } else {
                                                                                                                        v121 := GenS4{F0: 0, F1: 998, F2: -1}
                                                                                                                        var v122 bool = false
                                                                                                                        var v123 int16 = -820
                                                                                                                        _ = v121
                                                                                                                        _ = v122
                                                                                                                        _ = v123
                                                                                                                    }
                                                                                                                    v124 := GenS1{F0: 2}
                                                                                                                    var v125 int32 = 1
                                                                                                                    _ = v124
                                                                                                                    _ = v125
                                                                                                                }
                                                                                                                v2 = "a"
                                                                                                                v3 = (v2 + v2)
                                                                                                            }
                                                                                                        }
                                                                                                        v3 = "hello"
                                                                                                        _ = v107
                                                                                                    }
                                                                                                    var v126 float32 = -2.25
                                                                                                    _ = v126
                                                                                                }
                                                                                            case -9223372036854775808:
                                                                                                var v127 bool = true
                                                                                                _ = v127
                                                                                            }
                                                                                            v128 := GenS11{F0: int64(fi102)}
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = -207
                                                                                            }
                                                                                            _ = v128
                                                                                        } else {
                                                                                            var v129 uint32 = 2147483648
                                                                                            fmt.Println(v3, fi100)
                                                                                            _ = v129
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    v130 := GenS3{F0: 1}
                                                                                    fmt.Println(v2)
                                                                                    _ = v130
                                                                                }
                                                                            case 1:
                                                                                switch fi100 {
                                                                                case 813:
                                                                                    fmt.Println(v3, fi102)
                                                                                case -242:
                                                                                    v131 := GenS2{F0: (0.0 == -2.25), F1: (v2 + v3)}
                                                                                    _ = v131
                                                                                case 2147483647:
                                                                                    v132 := GenS7{F0: 57, F1: (v3 + v3), F2: "a"}
                                                                                    _ = v132
                                                                                }
                                                                            }
                                                                            var v133 int16 = 118
                                                                            _ = v133
                                                                        } else {
                                                                            v134 := GenS12{F0: (v32 % ((v32) | 1)), F1: 0}
                                                                            switch 650 {
                                                                            case -736:
                                                                                v5 = append(v5, fi102)
                                                                            default:
                                                                                v135 := v1[127]
                                                                                _ = v135
                                                                            }
                                                                            v5 = append(v5, fi102)
                                                                            var v136 uint16 = v32
                                                                            var v137 float64 = float64(1.5)
                                                                            _ = v134
                                                                            _ = v136
                                                                            _ = v137
                                                                        }
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                v32 += 1
                                                            default:
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            }
                                                        } else {
                                                            switch (-(fi100)) {
                                                            case 32767:
                                                                switch (-(fi100)) {
                                                                case 4294967295:
                                                                    var v138 uint32 = 1
                                                                    _ = v138
                                                                case -816:
                                                                    if (42 < -242) {
                                                                        switch ("a" + (v3 + v3)) {
                                                                        case "foo bar":
                                                                            v3 = (("foo bar" + "foo bar") + "\n\t")
                                                                        case "":
                                                                            v139 := GenS8{F0: -762, F1: v3, F2: 42}
                                                                            _ = v139
                                                                        case "hello":
                                                                            v140 := GenS13{F0: 629}
                                                                            _ = v140
                                                                        default:
                                                                            v141 := GenS14{F0: 1000}
                                                                            _ = v141
                                                                        }
                                                                    } else {
                                                                        switch fi102 {
                                                                        case -1:
                                                                            switch fi100 {
                                                                            case 790:
                                                                                if ((-128 == 42) && (!(!false))) {
                                                                                    if false {
                                                                                        v142 := GenS1{F0: 120}
                                                                                        _ = v142
                                                                                    }
                                                                                    var v143 int8 = 106
                                                                                    _ = v143
                                                                                } else {
                                                                                    var v144 int64 = 1
                                                                                    _ = v144
                                                                                }
                                                                            case 1:
                                                                                fmt.Println(fi102)
                                                                            default:
                                                                                v1[fi100] = 4294967295
                                                                            }
                                                                        }
                                                                    }
                                                                case 128:
                                                                    var v145 uint8 = 208
                                                                    _ = v145
                                                                }
                                                            case 127:
                                                                v146 := GenS10{F0: 127, F1: 2147483647}
                                                                _ = v146
                                                            }
                                                            v147 := GenS4{F0: 42, F1: 32767, F2: -1}
                                                            switch v3 {
                                                            case "hello":
                                                                var v148 int32 = 0
                                                                _ = v148
                                                            default:
                                                                v149 := GenS11{F0: -668}
                                                                _ = v149
                                                            }
                                                            _ = v147
                                                        }
                                                    default:
                                                        if len(v4) > 0 {
                                                            v4[0] = 32767
                                                        }
                                                    }
                                                    var v150 uint = 0
                                                    _ = v150
                                                } else {
                                                    var v151 float64 = 1.5
                                                    v2 = "a"
                                                    var v152 bool = false
                                                    _ = v151
                                                    _ = v152
                                                }
                                                _ = fi102
                                            }
                                        }
                                    } else {
                                        var v153 int16 = -441
                                        if len(v4) > 0 {
                                            v4[0] = fi100
                                        }
                                        _ = v153
                                    }
                                } else {
                                    for fi154 := 0; fi154 < 7; fi154++ {
                                        var v155 int16 = 1
                                        _ = fi154
                                        _ = v155
                                    }
                                }
                                _ = fi100
                            }
                        }
                    }
                }
            case 9223372036854775807:
                v2 = v2
            case 4294967295:
                for fi156 := 0; fi156 < 6; fi156++ {
                    if true {
                        var v157 int = fi156
                        if (984 < 97) {
                            for fi158 := 0; fi158 < 18; fi158++ {
                                switch "foo bar" {
                                case "":
                                    v159 := GenS6{F0: 2}
                                    _ = v159
                                default:
                                    if false {
                                        var v160 uint = 1
                                        v1[fi158] = -2147483648
                                        delete(v1, fi156)
                                        _ = v160
                                    } else {
                                        var v161 int16 = 1000
                                        v32 = 65535
                                        _ = v161
                                    }
                                }
                                _ = fi158
                            }
                            var v162 uint = 42
                            var v163 int = v157
                            _ = v162
                            _ = v163
                        } else {
                            var v164 bool = (887 != 818)
                            _ = v164
                        }
                        var v165 bool = false
                        _ = v157
                        _ = v165
                    } else {
                        v1[-430] = 922
                        var v166 int = (fi156 & int(v32))
                        _ = v166
                    }
                    _ = fi156
                }
            }
        }
    }
    v2 = "hello"
    delete(v1, 0)
    for fi167 := 0; fi167 < 17; fi167++ {
        var v168 uint32 = 613
        _ = fi167
        _ = v168
    }
    v32 -= 32767
    if (!true) {
        v169 := GenS2{F0: (!(65535 > uint32(v32))), F1: v3}
        _ = len(v4)
        var v170 uint = 2
        _ = v169
        _ = v170
    } else {
        delete(v1, -2147483648)
    }
    delete(v1, 308)
    v32 -= v32
    v171 := GenS15{F0: 0, F1: "a", F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v171
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
