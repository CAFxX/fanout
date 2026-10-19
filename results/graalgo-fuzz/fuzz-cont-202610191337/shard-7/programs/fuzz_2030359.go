package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int32; F2 bool }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 int32; F1 int8; F2 string }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8; F1 string; F2 int8 }
type GenS10 struct { F0 string; F1 uint }
type GenS11 struct { F0 uint32; F1 int16; F2 int64 }
type GenS12 struct { F0 uint64; F1 int64 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int32; F1 uint16 }
type GenS15 struct { F0 int; F1 uint; F2 bool }
func genFn0() uint32 {
    var v6 uint16 = 42
    v6 -= (v6 * v6)
    if false {
        var v7 int16 = 42
        fmt.Println(v6, v6)
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        switch 55 {
        case -9223372036854775808:
            v6 -= v6
        case -693:
            fmt.Println(fi8)
        case -961:
            v9 := GenS1{F0: (v6 | v6)}
            _ = v9
        default:
            switch ((("a" + "a") + "\n\t") + "") {
            case "":
                v6 -= (v6 & (v6 - (v6 | 34)))
            case "foo bar":
                fmt.Println(v6)
            }
        }
        v10 := []int{-9223372036854775808}
        var v11 bool = false
        _ = fi8
        _ = v10
        _ = v11
    }
    var v12 int = 32768
    _ = v12
    _ = v6
    return 2147483648
}
func genFn1(p0 string) int32 {
    v13 := GenS2{F0: (p0 + p0), F1: -596, F2: (257 >= -423)}
    p0 = "hello"
    for fi14 := 0; fi14 < 2; fi14++ {
        fmt.Println(fi14)
        v15 := GenS1{F0: 342}
        var v16 uint = 255
        _ = fi14
        _ = v15
        _ = v16
    }
    if (false || true) {
        p0 = ("a" + p0)
    } else {
        p0 = ""
    }
    _ = p0
    _ = v13
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 0.0
    var v3 uint64 = 4294967295
    v4 := make(map[int]uint32)
    v5 := []int{-9223372036854775808, -959, 2147483647}
    delete(v4, -408)
    var v17 float32 = (3.14159 + (1e10 + -2.25))
    for fi18 := 0; fi18 < 6; fi18++ {
        if (-23 == 0) {
            for fi19 := 0; fi19 < 3; fi19++ {
                v20 := GenS3{F0: -843, F1: "\n\t"}
                _ = fi19
                _ = v20
            }
            var v21 uint8 = 255
            if (42 > fi18) {
                _ = len(v1)
            } else {
                var v22 float32 = ((0.0 / (float32(3.14159) + 0.1)) - (v17 / v17))
                _ = v22
            }
            _ = v21
        } else {
            v3 += (v3 ^ (v3 / (((v3 / ((v3) | 1))) | 1)))
            v23 := GenS4{F0: (true && false), F1: 641}
            _ = v23
        }
        switch "a" {
        case "a":
            var v24 string = "hello"
            _ = v24
        }
        _ = fi18
    }
    var v25 int32 = 790
    if true {
        v3 = 4294967295
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 uint32 = 1
        switch "hello" {
        case "foo bar":
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                var v28 uint32 = v27
                _ = v28
            default:
                if (false && ("a" == "hello")) {
                    v3 = 2
                    for fi29 := 0; fi29 < 9; fi29++ {
                        fmt.Println(fi26, v2, v17)
                        fmt.Println(v27, v3)
                        v30 := GenS3{F0: 1, F1: (("" + "\n\t") + ("\n\t" + ("\n\t" + "a")))}
                        _ = fi29
                        _ = v30
                    }
                    v5 = append(v5, (fi26 + fi26))
                } else {
                    v31 := GenS5{F0: 0}
                    var v32 bool = ((^(v3)) != v3)
                    for fi33 := 0; fi33 < 17; fi33++ {
                        v3 = (v3 << ((440) & 63))
                        if v32 {
                            switch fi33 {
                            case 9223372036854775807:
                                _ = len(v5)
                            case -1:
                                if v32 {
                                    v34 := GenS5{F0: 100}
                                    v17 += (0.0 - v17)
                                    v35 := GenS1{F0: 1}
                                    _ = v34
                                    _ = v35
                                } else {
                                    v2 = float64((v17 * 3.14159))
                                    fmt.Println(v2, v17, fi33)
                                }
                            }
                            v36 := GenS1{F0: 255}
                            _ = v36
                        }
                        var v37 uint8 = 127
                        _ = fi33
                        _ = v37
                    }
                    _ = v31
                    _ = v32
                }
            }
        default:
            v38 := GenS6{F0: false, F1: (fi26 * fi26)}
            _ = v38
        }
        v25 -= (-(v25))
        _ = fi26
        _ = v27
    }
    v39 := GenS4{F0: (4294967295 > v3), F1: 18446744073709551615}
    var v40 int32 = (v25 | 2147483647)
    v17 = -2.25
    v25 = (v25 % (((v40 + v25)) | 1))
    v41 := GenS3{F0: genFn1("\n\t"), F1: ""}
    v2 -= v2
    for fi42 := 0; fi42 < 9; fi42++ {
        for fi43 := 0; fi43 < 5; fi43++ {
            if true {
                if ((69 > 0) || false) {
                    if true {
                        var v44 int32 = (-(v40))
                        _ = v44
                    } else {
                        switch ("hello" + "\n\t") {
                        case "":
                            fmt.Println(fi42, v25)
                        case "a":
                            var v45 int = fi42
                            _ = v45
                        case "foo bar":
                            _ = len(v4)
                        }
                    }
                }
                var v46 uint64 = v3
                _ = v46
            }
            fmt.Println(v40)
            if true {
                switch "hello" {
                case "a":
                    var v47 int8 = -80
                    _ = v47
                }
                if (2 != 620) {
                    var v48 int16 = -221
                    _ = v48
                }
                _ = len(v1)
                v49 := GenS7{F0: v40, F1: -128, F2: (("" + (("hello" + "hello") + "foo bar")) + "hello")}
                _ = v49
            } else {
                v50 := GenS8{F0: (v40 + genFn1(("foo bar" + ((("foo bar" + "\n\t") + "\n\t") + ("foo bar" + ("foo bar" + "hello"))))))}
                if (!false) {
                    v51 := GenS9{F0: 241, F1: "\n\t", F2: 127}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v51
                } else {
                    var v52 float32 = 1.5
                    var v53 bool = false
                    _ = len(v1)
                    _ = v52
                    _ = v53
                }
                var v54 bool = (!(!(343 == 414)))
                _ = v50
                _ = v54
            }
            fmt.Println(v2, fi43)
            _ = fi43
        }
        _ = fi42
    }
    _ = len(v5)
    fmt.Println(v2, v17, v25)
    if false {
        var v55 string = "hello"
        var v56 bool = false
        _ = v55
        _ = v56
    } else {
        if false {
            if (-108 <= v40) {
                v40 = 2147483647
                var v57 int64 = 9223372036854775807
                _ = v57
            }
        }
    }
    var v58 bool = true
    switch 32768 {
    case 524:
        switch (("" + ("" + "hello")) + "foo bar") {
        case "a":
            if (240 != 4294967295) {
                _ = len(v1)
                var v59 int = 127
                _ = v59
            }
        case "\n\t":
            if (0 >= 42) {
                if ((!((((("foo bar" + "a") + "") + "hello") + "") >= "foo bar")) && (-128 <= 60)) {
                    for fi60 := 0; fi60 < 3; fi60++ {
                        fmt.Println(v3, v3)
                        switch "hello" {
                        case "a":
                            fmt.Println(v3)
                        }
                        _ = fi60
                    }
                    if (0 > 926) {
                        fmt.Println(v3)
                        if v58 {
                            v61 := GenS5{F0: -1}
                            switch -9223372036854775808 {
                            case 32767:
                                var v62 int32 = (v40 << ((25) & 63))
                                _ = v62
                            case 1:
                                for fi63 := 0; fi63 < 3; fi63++ {
                                    var v64 float32 = 0.1
                                    v65 := GenS8{F0: (-(v40))}
                                    _ = fi63
                                    _ = v64
                                    _ = v65
                                }
                            case 32768:
                                switch -502 {
                                case 32767:
                                    v25 += (^(v25))
                                case -116:
                                    var v66 uint8 = 2
                                    _ = v66
                                }
                            default:
                                if v58 {
                                    if (v58 && v58) {
                                        for fi67 := 0; fi67 < 1; fi67++ {
                                            v3 += v3
                                            _ = fi67
                                        }
                                        v17 = (1e10 + v17)
                                        if (!v58) {
                                            _ = len(v1)
                                            switch 279 {
                                            case 1:
                                                if (!(18446744073709551615 <= 255)) {
                                                    var v68 float32 = v17
                                                    _ = v68
                                                } else {
                                                    var v69 float32 = -2.25
                                                    var v70 bool = (0 != 1)
                                                    delete(v4, -1)
                                                    _ = v69
                                                    _ = v70
                                                }
                                            default:
                                                fmt.Println(v17)
                                            }
                                            switch (("" + "a") + "a") {
                                            case "hello":
                                                if v58 {
                                                    var v71 int8 = 109
                                                    for fi72 := 0; fi72 < 2; fi72++ {
                                                        delete(v4, fi72)
                                                        _ = fi72
                                                    }
                                                    v5 = append(v5, -1)
                                                    _ = v71
                                                }
                                            case "":
                                                if true {
                                                    if (!(!(false && ((255 >= 721) || ((v2 / 3.14159) > 0.0))))) {
                                                        var v73 uint32 = 42
                                                        _ = v73
                                                    }
                                                } else {
                                                    var v74 float64 = ((v2 - v2) * v2)
                                                    for fi75 := 0; fi75 < 9; fi75++ {
                                                        switch "a" {
                                                        case "hello":
                                                            for fi76 := 0; fi76 < 20; fi76++ {
                                                                if (!true) {
                                                                    var v77 uint8 = 42
                                                                    _ = v77
                                                                } else {
                                                                    v3 += 18446744073709551615
                                                                    v78 := GenS3{F0: (^(v40)), F1: "foo bar"}
                                                                    fmt.Println(v58, v2, fi76)
                                                                    _ = v78
                                                                }
                                                                _ = fi76
                                                            }
                                                        }
                                                        _ = fi75
                                                    }
                                                    v40 -= genFn1("\n\t")
                                                    _ = v74
                                                }
                                            default:
                                                v17 = (v17 * 1.5)
                                            }
                                        } else {
                                            var v79 uint16 = 2
                                            v17 = v17
                                            _ = v79
                                        }
                                    } else {
                                        if false {
                                            for fi80 := 0; fi80 < 14; fi80++ {
                                                if false {
                                                    if (true || (!(219 > 32767))) {
                                                        if (!v58) {
                                                            for fi81 := 0; fi81 < 5; fi81++ {
                                                                if v58 {
                                                                    v5 = append(v5, (fi81 + fi81))
                                                                    fmt.Println(fi81, v17)
                                                                    delete(v4, fi80)
                                                                } else {
                                                                    fmt.Println(v3, v40)
                                                                }
                                                                fmt.Println(v17)
                                                                _ = fi81
                                                            }
                                                        } else {
                                                            switch (fi80 >> ((18446744073709551615) & 63)) {
                                                            case 2147483647:
                                                                if ((2 >= 809) || true) {
                                                                    if v58 {
                                                                        var v82 int = (fi80 % ((1) | 1))
                                                                        _ = v82
                                                                    }
                                                                    for fi83 := 0; fi83 < 1; fi83++ {
                                                                        fmt.Println(v25, v3)
                                                                        _ = len(v4)
                                                                        v40 = (v40 ^ 689)
                                                                        _ = fi83
                                                                    }
                                                                    v84 := GenS1{F0: 753}
                                                                    _ = v84
                                                                } else {
                                                                    var v85 uint8 = 0
                                                                    for fi86 := 0; fi86 < 2; fi86++ {
                                                                        switch "hello" {
                                                                        case "":
                                                                            var v87 uint16 = 255
                                                                            var v88 bool = ((genFn0() == 0) && v58)
                                                                            _ = v87
                                                                            _ = v88
                                                                        case "a":
                                                                            var v89 uint32 = 684
                                                                            _ = v89
                                                                        }
                                                                        v3 += 0
                                                                        _ = fi86
                                                                    }
                                                                    v58 = v58
                                                                    _ = v85
                                                                }
                                                            case 218:
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            case 803:
                                                                var v90 int16 = 1
                                                                _ = v90
                                                            }
                                                            var v91 uint32 = genFn0()
                                                            _ = v91
                                                        }
                                                    }
                                                    for fi92 := 0; fi92 < 18; fi92++ {
                                                        var v93 uint16 = 32767
                                                        if v58 {
                                                            v58 = v58
                                                            v3 = 591
                                                        } else {
                                                            var v94 string = "foo bar"
                                                            var v95 int = 32767
                                                            _ = v94
                                                            _ = v95
                                                        }
                                                        if (v93 <= 65535) {
                                                            v96 := GenS2{F0: ((("a" + ("foo bar" + "hello")) + "\n\t") + ""), F1: (v25 + -2147483648), F2: ((!((v58 || (fi80 != (fi80 / ((fi80) | 1)))) || v58)) && (!false))}
                                                            _ = v96
                                                        } else {
                                                            v97 := GenS10{F0: (((("hello" + "a") + "") + "\n\t") + ("" + ("\n\t" + "hello"))), F1: 255}
                                                            switch "a" {
                                                            case "foo bar":
                                                                v5 = append(v5, fi80)
                                                            }
                                                            v98 := GenS8{F0: 2147483647}
                                                            _ = v97
                                                            _ = v98
                                                        }
                                                        _ = fi92
                                                        _ = v93
                                                    }
                                                } else {
                                                    v5 = append(v5, (fi80 << ((959) & 63)))
                                                    if true {
                                                        for fi99 := 0; fi99 < 16; fi99++ {
                                                            switch "a" {
                                                            case "hello":
                                                                var v100 int64 = 42
                                                                _ = v100
                                                            }
                                                            v17 = 1.5
                                                            switch "foo bar" {
                                                            case "a":
                                                                v58 = true
                                                            default:
                                                                v101 := GenS11{F0: 489, F1: 383, F2: 42}
                                                                _ = v101
                                                            }
                                                            _ = fi99
                                                        }
                                                        for fi102 := 0; fi102 < 17; fi102++ {
                                                            v58 = ((v58 || v58) || true)
                                                            switch (fi102 / (((fi80 & (^(fi80)))) | 1)) {
                                                            case -940:
                                                                var v103 uint = 1
                                                                _ = v103
                                                            default:
                                                                var v104 bool = v58
                                                                _ = v104
                                                            }
                                                            _ = fi102
                                                        }
                                                    }
                                                }
                                                var v105 uint8 = 161
                                                v105 -= v105
                                                _ = fi80
                                                _ = v105
                                            }
                                            var v106 int64 = 9223372036854775807
                                            v107 := GenS12{F0: v3, F1: (v106 + (v106 | v106))}
                                            _ = v106
                                            _ = v107
                                        }
                                        if ((0 >= 2147483647) && (-1 < -1)) {
                                            var v108 string = "\n\t"
                                            _ = len(v1)
                                            v109 := GenS13{F0: 32767}
                                            _ = v108
                                            _ = v109
                                        }
                                        fmt.Println(v58, v3, v40)
                                    }
                                }
                            }
                            _ = v61
                        }
                        if (true || (846 != 9223372036854775807)) {
                            var v110 uint64 = v3
                            if v58 {
                                if len(v5) > 0 {
                                    v5[0] = 32767
                                }
                                v111 := GenS10{F0: ("foo bar" + ("foo bar" + "hello")), F1: 392}
                                v112 := GenS14{F0: v25, F1: 42}
                                _ = v111
                                _ = v112
                            }
                            _ = v110
                        }
                    }
                } else {
                    var v113 uint8 = 128
                    v17 = float32((v2 * v2))
                    _ = v113
                }
                for fi114 := 0; fi114 < 1; fi114++ {
                    if v58 {
                        var v115 float64 = float64(v17)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v115
                    }
                    _ = fi114
                }
            } else {
                fmt.Println(v17, v40)
            }
        case "hello":
            switch "foo bar" {
            case "a":
                v116 := GenS2{F0: "foo bar", F1: (v25 & v40), F2: (-571 >= 0)}
                _ = v116
            case "":
                for fi117 := 0; fi117 < 20; fi117++ {
                    if false {
                        v4[(fi117 - fi117)] = 1
                        v2 = v2
                        _ = len(v5)
                    }
                    v118 := GenS15{F0: (^(fi117)), F1: 998, F2: (684 > 65535)}
                    _ = fi117
                    _ = v118
                }
            case "foo bar":
                v2 = v2
            default:
                var v119 bool = ((v40 & v40) >= v25)
                _ = v119
            }
        }
    default:
        for fi120 := 0; fi120 < 17; fi120++ {
            switch 42 {
            case 127:
                v58 = false
            default:
                v121 := GenS12{F0: v3, F1: 9223372036854775807}
                _ = v121
            }
            v122 := GenS9{F0: 255, F1: "a", F2: 100}
            _ = fi120
            _ = v122
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v58
}
