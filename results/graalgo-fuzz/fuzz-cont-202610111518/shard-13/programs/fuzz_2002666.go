package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint32; F1 int16; F2 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint16; F1 int64 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 bool; F1 int; F2 int }
type GenS11 struct { F0 string; F1 int32; F2 uint }
type GenS12 struct { F0 int }
type GenS13 struct { F0 string; F1 int16 }
type GenS14 struct { F0 uint64; F1 int16 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 int16; F1 bool }
type GenS17 struct { F0 uint; F1 string; F2 int }
type GenS18 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 uint; F1 uint64 }
type GenS21 struct { F0 uint8; F1 int64; F2 int32 }
type GenS22 struct { F0 uint64 }
func genFn0(p0 uint64) uint {
    fmt.Println(p0)
    var v7 float32 = 1e10
    for fi8 := 0; fi8 < 8; fi8++ {
        p0 -= 4294967295
        _ = fi8
    }
    _ = p0
    _ = v7
    return 480
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{371, -306, 790, 2147483647}
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    v5 := []string{(("a" + (((("foo bar" + "hello") + (("" + "") + "a")) + "foo bar") + "\n\t")) + "\n\t"), "foo bar", "a"}
    var v6 int64 = v1
    if ((!(1e10 <= (0.0 / (((0.1 / 1.5) / (0.0 - 1.5)) + 3.14159)))) && true) {
        var v9 uint = 255
        switch "\n\t" {
        case "\n\t":
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := GenS1{F0: 1, F1: -1}
                fmt.Println(v1, v6, v1)
                var v12 int16 = 32767
                _ = fi10
                _ = v11
                _ = v12
            }
        default:
            v6 -= -9223372036854775808
        }
        _ = v9
    }
    var v13 int16 = -32768
    fmt.Println(v1, v6, v1)
    v1 -= v1
    v14 := GenS1{F0: (-(v13)), F1: 833}
    if false {
        v15 := GenS2{F0: 1, F1: (!true), F2: 42}
        for fi16 := 0; fi16 < 5; fi16++ {
            if ((((1e10 + 1.5) + -2.25) - 0.0) > 0.0) {
                v1 -= v1
            } else {
                for fi17 := 0; fi17 < 15; fi17++ {
                    v18 := GenS1{F0: (v13 + v13), F1: 893}
                    var v19 int32 = -2147483648
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            v6 = v1
            for fi20 := 0; fi20 < 4; fi20++ {
                _ = len(v3)
                if true {
                    v21 := GenS2{F0: 0, F1: ((2 != 603) || (-730 > 0)), F2: 2}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v6 += 42
                    var v22 uint16 = 1
                    _ = v21
                    _ = v22
                }
                _ = fi20
            }
            _ = fi16
        }
        switch (("\n\t" + "a") + "") {
        case "foo bar":
            v1 = v1
        case "a":
            v23 := GenS2{F0: 1, F1: true, F2: 4294967295}
            _ = v23
        }
        _ = v15
    }
    v24 := GenS1{F0: (v13 % (((v13 | v13)) | 1)), F1: 2147483647}
    v25 := GenS3{F0: -883, F1: 42}
    var v26 int8 = -5
    switch 0 {
    case 32767:
        v26 += (-(v26))
    case -2147483648:
        for fi27 := 0; fi27 < 9; fi27++ {
            switch (-(fi27)) {
            case 32768:
                if true {
                    v13 += 935
                    v6 = 1
                    switch (-(fi27)) {
                    case 384:
                        if false {
                            v28 := GenS4{F0: 2147483648, F1: (-(v13)), F2: 0}
                            v1 = v6
                            v26 += (v26 & v26)
                            _ = v28
                        } else {
                            switch (-(fi27)) {
                            case 42:
                                var v29 int64 = v1
                                _ = v29
                            case 0:
                                v30 := GenS5{F0: "hello"}
                                _ = v30
                            default:
                                v26 += (-(v26))
                            }
                            var v31 float32 = 1.5
                            _ = v31
                        }
                    default:
                        switch (fi27 << ((287) & 63)) {
                        case 128:
                            var v32 bool = (0.0 == (1.5 - 0.0))
                            _ = v32
                        default:
                            v26 += (v26 | 0)
                        }
                    }
                } else {
                    _ = len(v2)
                    var v33 uint16 = 32767
                    v6 -= (v1 | -1)
                    _ = v33
                }
            default:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = fi27
        }
    case 128:
        v13 = -875
    default:
        if (867 <= 4294967295) {
            var v34 uint64 = 1
            v13 = 32767
            var v35 float64 = 0.0
            _ = v34
            _ = v35
        } else {
            var v36 uint = 255
            _ = v36
        }
    }
    if false {
        v37 := GenS6{F0: -557, F1: -2147483648}
        var v38 uint32 = 381
        for fi39 := 0; fi39 < 3; fi39++ {
            v13 = v13
            v5 = append(v5, "a")
            if (-513 <= v13) {
                if false {
                    v5 = append(v5, "")
                    if true {
                        if true {
                            var v40 bool = (-539 == int32(v1))
                            _ = v40
                        }
                        var v41 float32 = 1e10
                        switch "hello" {
                        case "foo bar":
                            v42 := GenS6{F0: -1, F1: (fi39 ^ fi39)}
                            _ = v42
                        default:
                            if false {
                                var v43 float32 = 1e10
                                for fi44 := 0; fi44 < 10; fi44++ {
                                    var v45 uint16 = 42
                                    _ = fi44
                                    _ = v45
                                }
                                _ = len(v4)
                                _ = v43
                            } else {
                                if false {
                                    var v46 string = "\n\t"
                                    v47 := GenS2{F0: 0, F1: true, F2: 591}
                                    _ = v46
                                    _ = v47
                                } else {
                                    if len(v5) > 0 {
                                        v5[0] = ""
                                    }
                                    var v48 int16 = (v13 / ((v13) | 1))
                                    v13 = (v48 << ((18446744073709551615) & 63))
                                    _ = v48
                                }
                            }
                        }
                        _ = v41
                    } else {
                        var v49 uint8 = 1
                        switch (fi39 >> ((398) & 63)) {
                        case 1:
                            v6 = v6
                        case 2147483647:
                            var v50 uint = 62
                            _ = v50
                        case 345:
                            if true {
                                if (((3.14159 - (0.0 / 1.0)) * 1.5) > -2.25) {
                                    v1 = 42
                                    var v51 float32 = 0.0
                                    _ = v51
                                } else {
                                    var v52 bool = ((float64(-2.25) + (float64((3.14159 / (-2.25 + -2.25))) - -2.25)) != float64((0.1 * (1.5 * ((1.5 + 0.1) + (1.5 / float32(3.14159)))))))
                                    var v53 string = ("a" + "")
                                    _ = v52
                                    _ = v53
                                }
                                switch fi39 {
                                case 9223372036854775807:
                                    if true {
                                        var v54 uint64 = 42
                                        v55 := GenS2{F0: uint64(v49), F1: (v49 >= (v49 >> ((783) & 63))), F2: 2}
                                        _ = v54
                                        _ = v55
                                    } else {
                                        v38 = (v38 + (v38 / (((v38 + (^(v38)))) | 1)))
                                    }
                                case 359:
                                    if true {
                                        for fi56 := 0; fi56 < 13; fi56++ {
                                            switch ("hello" + "") {
                                            case "\n\t":
                                                v38 = (v38 * v38)
                                                v57 := GenS2{F0: 2, F1: false, F2: 255}
                                                _ = v57
                                            case "hello":
                                                if false {
                                                    switch "hello" {
                                                    case "hello":
                                                        switch "\n\t" {
                                                        case "hello":
                                                            var v58 uint64 = 1
                                                            _ = v58
                                                        default:
                                                            v38 -= v38
                                                        }
                                                    case "a":
                                                        v59 := GenS5{F0: ("" + ("foo bar" + "a"))}
                                                        _ = v59
                                                    case "\n\t":
                                                        fmt.Println(v6, v49)
                                                    default:
                                                        v38 = (v38 % ((v38) | 1))
                                                    }
                                                    if (false || false) {
                                                        switch (^(fi39)) {
                                                        case 42:
                                                            v60 := GenS2{F0: 1, F1: false, F2: 42}
                                                            _ = v60
                                                        default:
                                                            v26 += v26
                                                        }
                                                        v61 := GenS3{F0: -679, F1: 1}
                                                        var v62 uint64 = uint64(v38)
                                                        _ = v61
                                                        _ = v62
                                                    } else {
                                                        if (((false || (!(((((0.1 / 1.0) + (1.5 * 1e10)) / 1.0) / ((1e10 - -2.25) - 0.1)) >= (0.0 * 3.14159)))) || false) || (true && true)) {
                                                            var v63 uint8 = (v49 ^ v49)
                                                            _ = v63
                                                        }
                                                    }
                                                    var v64 float64 = 3.14159
                                                    _ = v64
                                                } else {
                                                    v65 := GenS5{F0: "a"}
                                                    if (42 == 4294967295) {
                                                        var v66 float32 = ((0.1 + 1e10) - 3.14159)
                                                        v26 -= -63
                                                        _ = v66
                                                    } else {
                                                        var v67 uint = 587
                                                        _ = v67
                                                    }
                                                    v68 := GenS7{F0: "foo bar", F1: true}
                                                    _ = v65
                                                    _ = v68
                                                }
                                            }
                                            _ = fi56
                                        }
                                    } else {
                                        if false {
                                            var v69 string = "\n\t"
                                            var v70 string = v69
                                            _ = v69
                                            _ = v70
                                        } else {
                                            if ((!((fi39 << ((genFn0(18446744073709551615)) & 63)) >= (fi39 * fi39))) && false) {
                                                v13 += v13
                                                v2 = append(v2, fi39)
                                            }
                                            v13 = v13
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    }
                                default:
                                    switch "" {
                                    case "a":
                                        var v71 int8 = v26
                                        _ = v71
                                    case "foo bar":
                                        var v72 string = "foo bar"
                                        _ = v72
                                    }
                                }
                                switch (("hello" + "a") + (("\n\t" + "foo bar") + ("" + (("" + "foo bar") + ("a" + (((("" + "hello") + "") + "") + ((("\n\t" + "hello") + "foo bar") + "\n\t"))))))) {
                                case "a":
                                    for fi73 := 0; fi73 < 7; fi73++ {
                                        if ((v26 + 0) > -128) {
                                            v49 += (v49 >> ((840) & 63))
                                            if true {
                                                v1 -= v6
                                            }
                                            var v74 uint16 = 0
                                            _ = v74
                                        }
                                        v26 = int8(v13)
                                        if (1 > (^(v38))) {
                                            v26 += (v26 ^ -128)
                                            var v75 float64 = (-2.25 / ((-2.25 * (1e10 + 0.0)) + float64(1.5)))
                                            v76 := v4["\n\t"]
                                            v1 += v1
                                            if (841 != 18446744073709551615) {
                                                var v77 string = ""
                                                fmt.Println(fi39, fi73, fi39)
                                                var v78 int16 = 1000
                                                _ = v77
                                                _ = v78
                                            } else {
                                                if false {
                                                    v79 := GenS8{F0: 0, F1: 226}
                                                    fmt.Println(v13, fi73, v1)
                                                    _ = v79
                                                } else {
                                                    if (65535 > v76) {
                                                        var v80 float32 = ((float32(v75) + 0.0) * 1e10)
                                                        v13 += (v13 | 32767)
                                                        _ = v80
                                                    } else {
                                                        if false {
                                                            v1 += (v1 & v6)
                                                            switch 2147483647 {
                                                            case 4294967295:
                                                                fmt.Println(fi73, v75)
                                                            case 2147483647:
                                                                if ((-(fi39)) != fi39) {
                                                                    v13 = (v13 & v13)
                                                                    fmt.Println(v75)
                                                                } else {
                                                                    var v81 int32 = -283
                                                                    if true {
                                                                        if false {
                                                                            var v82 uint = 1
                                                                            fmt.Println(fi73, v76)
                                                                            v75 = (v75 * ((3.14159 - (float64(0.0) + v75)) + v75))
                                                                            var v83 uint32 = (v76 ^ (v76 >> (((^(v82))) & 63)))
                                                                            _ = v82
                                                                            _ = v83
                                                                        }
                                                                        var v84 uint64 = 1
                                                                        var v85 float64 = v75
                                                                        _ = v84
                                                                        _ = v85
                                                                    } else {
                                                                        switch fi73 {
                                                                        case 4294967295:
                                                                            v13 = (-(v13))
                                                                        case 9223372036854775807:
                                                                            if ((0 > 42) && (true && (v81 < (v81 * -830)))) {
                                                                                if (2147483648 == 381) {
                                                                                    fmt.Println(fi73)
                                                                                    var v86 bool = (!(false && (false && false)))
                                                                                    _ = v86
                                                                                } else {
                                                                                    if (0.1 == v75) {
                                                                                        if (v75 >= 1e10) {
                                                                                            v87 := GenS4{F0: v38, F1: (v13 * (^(v13))), F2: v81}
                                                                                            if true {
                                                                                                if false {
                                                                                                    fmt.Println(v6, v6)
                                                                                                }
                                                                                                switch ("foo bar" + "\n\t") {
                                                                                                case "hello":
                                                                                                    fmt.Println(v6, v13)
                                                                                                case "\n\t":
                                                                                                    v88 := GenS8{F0: 42, F1: v1}
                                                                                                    _ = v88
                                                                                                default:
                                                                                                    v26 += v26
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(v76, v76, v1)
                                                                                                var v89 int16 = int16(v1)
                                                                                                v4[("" + "\n\t")] = v76
                                                                                                _ = v89
                                                                                            }
                                                                                            _ = v87
                                                                                        }
                                                                                        var v90 int8 = 127
                                                                                        var v91 int16 = (v13 ^ -1)
                                                                                        _ = v90
                                                                                        _ = v91
                                                                                    } else {
                                                                                        fmt.Println(v1, v38, v6)
                                                                                        var v92 float32 = 1e10
                                                                                        _ = v92
                                                                                    }
                                                                                    fmt.Println(v1, v76, v38)
                                                                                    v38 -= v76
                                                                                }
                                                                            }
                                                                        case 128:
                                                                            if (1 > 18446744073709551615) {
                                                                                if (!((-2.25 <= 1.5) || false)) {
                                                                                    v13 += v13
                                                                                    v5 = append(v5, ("foo bar" + "hello"))
                                                                                }
                                                                                var v93 float32 = ((1e10 - 1.5) - (-2.25 + float32(v75)))
                                                                                if false {
                                                                                    _ = len(v4)
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = "\n\t"
                                                                                    }
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = ""
                                                                                    }
                                                                                } else {
                                                                                    if true {
                                                                                        v38 -= 4294967295
                                                                                        fmt.Println(v26)
                                                                                    }
                                                                                    fmt.Println(v13, v38, v49)
                                                                                }
                                                                                _ = v93
                                                                            } else {
                                                                                var v94 float64 = v75
                                                                                v95 := GenS9{F0: 32767}
                                                                                v96 := v4["foo bar"]
                                                                                _ = v94
                                                                                _ = v95
                                                                                _ = v96
                                                                            }
                                                                        }
                                                                    }
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    _ = v81
                                                                }
                                                            case -522:
                                                                fmt.Println(v1, v75)
                                                            }
                                                        } else {
                                                            delete(v4, "a")
                                                            fmt.Println(v49, fi73)
                                                        }
                                                        var v97 int64 = (v1 >> ((0) & 63))
                                                        v13 = v13
                                                        _ = v97
                                                    }
                                                    if true {
                                                        _ = len(v3)
                                                        if true {
                                                            switch fi39 {
                                                            case 494:
                                                                switch fi73 {
                                                                case 1:
                                                                    v98 := GenS3{F0: 0, F1: 42}
                                                                    _ = v98
                                                                case -280:
                                                                    switch ("" + ("hello" + "foo bar")) {
                                                                    case "":
                                                                        v49 += (^(v49))
                                                                    case "\n\t":
                                                                        if true {
                                                                            var v99 string = "a"
                                                                            var v100 int32 = -1
                                                                            v101 := GenS10{F0: (106 <= v49), F1: (-(fi39)), F2: (^(fi39))}
                                                                            var v102 string = v99
                                                                            _ = v100
                                                                            _ = v101
                                                                            _ = v102
                                                                            _ = v99
                                                                        } else {
                                                                            switch (("" + "\n\t") + ("hello" + ("\n\t" + "\n\t"))) {
                                                                            case "\n\t":
                                                                                fmt.Println(v6, v76)
                                                                            case "":
                                                                                v75 = v75
                                                                            default:
                                                                                var v103 uint64 = 922
                                                                                _ = v103
                                                                            }
                                                                            if (uint8(v76) > v49) {
                                                                                v104 := GenS11{F0: (("a" + "") + "a"), F1: 457, F2: 0}
                                                                                fmt.Println(fi73, v49, v13)
                                                                                var v105 float32 = 0.0
                                                                                v106 := GenS12{F0: (fi73 << ((1) & 63))}
                                                                                _ = v104
                                                                                _ = v105
                                                                                _ = v106
                                                                            } else {
                                                                                v6 = (v6 / ((0) | 1))
                                                                                v75 += v75
                                                                            }
                                                                        }
                                                                    case "foo bar":
                                                                        switch 128 {
                                                                        case 2147483647:
                                                                            v5 = append(v5, ("hello" + ((("foo bar" + ("a" + "hello")) + ("hello" + "a")) + (("" + (("a" + (("a" + "foo bar") + (("\n\t" + ("foo bar" + ("foo bar" + "\n\t"))) + (("\n\t" + ("foo bar" + (("foo bar" + "foo bar") + "a"))) + "\n\t")))) + "foo bar")) + "foo bar"))))
                                                                        case 1:
                                                                            v1 = v6
                                                                        case 9223372036854775807:
                                                                            v6 -= (^(v1))
                                                                        }
                                                                    }
                                                                case 127:
                                                                    if (genFn0(928) != 42) {
                                                                        switch ("" + "") {
                                                                        case "foo bar":
                                                                            switch fi73 {
                                                                            case -267:
                                                                                var v107 int8 = (v26 + v26)
                                                                                _ = v107
                                                                            case 32767:
                                                                                switch (fi73 ^ (-(fi39))) {
                                                                                case 9223372036854775807:
                                                                                    fmt.Println(v75, v76)
                                                                                case -707:
                                                                                    v26 -= 44
                                                                                case 127:
                                                                                    fmt.Println(v76, fi73, v26)
                                                                                default:
                                                                                    v13 = 42
                                                                                    if false {
                                                                                        v108 := GenS12{F0: (^(fi73))}
                                                                                        if (1 > v26) {
                                                                                            v49 -= v49
                                                                                        } else {
                                                                                            var v109 uint32 = (^(v76))
                                                                                            v38 = v76
                                                                                            var v110 uint16 = 136
                                                                                            _ = v109
                                                                                            _ = v110
                                                                                        }
                                                                                        _ = v108
                                                                                    }
                                                                                    if true {
                                                                                        v13 -= 1
                                                                                        v1 = 42
                                                                                        if (true || true) {
                                                                                            v2 = append(v2, int(v26))
                                                                                            switch fi73 {
                                                                                            case 383:
                                                                                                if (65535 < (v38 | 261)) {
                                                                                                    v13 += v13
                                                                                                    _ = len(v3)
                                                                                                    var v111 uint8 = (v49 | (v49 << ((4294967295) & 63)))
                                                                                                    _ = v111
                                                                                                } else {
                                                                                                    v112 := GenS13{F0: ("hello" + ("foo bar" + "foo bar")), F1: (^(v13))}
                                                                                                    fmt.Println(v75)
                                                                                                    _ = v112
                                                                                                }
                                                                                            default:
                                                                                                switch int(v49) {
                                                                                                case 4294967295:
                                                                                                    if (uint16(v38) <= 473) {
                                                                                                        v38 += (v76 & (v76 >> ((255) & 63)))
                                                                                                        if (!true) {
                                                                                                            v26 = v26
                                                                                                        }
                                                                                                        v13 = int16(v6)
                                                                                                    }
                                                                                                case -441:
                                                                                                    v113 := GenS14{F0: 0, F1: v13}
                                                                                                    _ = v113
                                                                                                }
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    _ = len(v3)
                                                                                                default:
                                                                                                    v26 += (v26 & (v26 << ((0) & 63)))
                                                                                                }
                                                                                                switch fi39 {
                                                                                                case -9223372036854775808:
                                                                                                    var v114 bool = false
                                                                                                    _ = v114
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(fi73, v76)
                                                                                        } else {
                                                                                            switch ("hello" + "") {
                                                                                            case "foo bar":
                                                                                                var v115 int64 = v1
                                                                                                _ = v115
                                                                                            default:
                                                                                                switch "foo bar" {
                                                                                                case "":
                                                                                                    v5 = append(v5, "hello")
                                                                                                    var v116 float32 = (0.0 * (-2.25 + (1.5 - float32(((float64(1.5) / 1e10) / v75)))))
                                                                                                    _ = v116
                                                                                                case "foo bar":
                                                                                                    delete(v3, uint32(v49))
                                                                                                    v49 += 1
                                                                                                case "\n\t":
                                                                                                    if true {
                                                                                                        v49 = v49
                                                                                                    }
                                                                                                default:
                                                                                                    v26 = v26
                                                                                                }
                                                                                            }
                                                                                            v26 = -70
                                                                                        }
                                                                                    } else {
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = ("hello" + ("foo bar" + "foo bar"))
                                                                                        }
                                                                                        v117 := GenS15{F0: "\n\t"}
                                                                                        if false {
                                                                                            switch (("" + "foo bar") + ("" + "\n\t")) {
                                                                                            case "a":
                                                                                                if (v26 <= (v26 | v26)) {
                                                                                                    v118 := GenS13{F0: "\n\t", F1: 0}
                                                                                                    _ = v118
                                                                                                }
                                                                                            case "":
                                                                                                v6 = v6
                                                                                            case "hello":
                                                                                                fmt.Println(v38)
                                                                                            }
                                                                                            if true {
                                                                                                switch fi39 {
                                                                                                case -1:
                                                                                                    v119 := GenS16{F0: v13, F1: (v13 <= v13)}
                                                                                                    _ = v119
                                                                                                default:
                                                                                                    v5 = append(v5, ("a" + "\n\t"))
                                                                                                }
                                                                                            } else {
                                                                                                if false {
                                                                                                    var v120 float64 = float64(0.1)
                                                                                                    _ = v120
                                                                                                } else {
                                                                                                    var v121 int = fi39
                                                                                                    v121 = v121
                                                                                                    if false {
                                                                                                        v122 := GenS13{F0: "foo bar", F1: v13}
                                                                                                        var v123 uint8 = v49
                                                                                                        if (42 == 65535) {
                                                                                                            var v124 int = 0
                                                                                                            fmt.Println(v76)
                                                                                                            _ = v124
                                                                                                        }
                                                                                                        _ = v122
                                                                                                        _ = v123
                                                                                                    } else {
                                                                                                        v125 := GenS11{F0: "", F1: 648, F2: 0}
                                                                                                        _ = v125
                                                                                                    }
                                                                                                    var v126 int8 = int8(v76)
                                                                                                    _ = v121
                                                                                                    _ = v126
                                                                                                }
                                                                                                v6 += 42
                                                                                                v75 -= v75
                                                                                                v127 := GenS14{F0: 2, F1: int16(fi73)}
                                                                                                v49 += (v49 & (v49 / (((^(v49))) | 1)))
                                                                                                _ = v127
                                                                                            }
                                                                                            if false {
                                                                                                v4["hello"] = 665
                                                                                            }
                                                                                        }
                                                                                        if ((v13 / ((-13) | 1)) == v13) {
                                                                                            fmt.Println(v13, v13, v76)
                                                                                        } else {
                                                                                            var v128 uint8 = (v49 + (v49 ^ (v49 & v49)))
                                                                                            switch (fi39 ^ fi39) {
                                                                                            case -1:
                                                                                                var v129 string = "a"
                                                                                                _ = v129
                                                                                            }
                                                                                            v130 := GenS4{F0: 787, F1: 1, F2: 22}
                                                                                            v75 = 1e10
                                                                                            fmt.Println(v76)
                                                                                            _ = v128
                                                                                            _ = v130
                                                                                        }
                                                                                        var v131 float64 = 1e10
                                                                                        _ = v117
                                                                                        _ = v131
                                                                                    }
                                                                                    fmt.Println(v75, v49)
                                                                                }
                                                                                var v132 float32 = 0.1
                                                                                var v133 int8 = (^(v26))
                                                                                _ = v132
                                                                                _ = v133
                                                                            default:
                                                                                var v134 int16 = (v13 / (((v13 % ((v13) | 1))) | 1))
                                                                                _ = v134
                                                                            }
                                                                        default:
                                                                            var v135 uint64 = 1
                                                                            _ = v135
                                                                        }
                                                                        if ((^(v49)) != 112) {
                                                                            var v136 bool = (!(!true))
                                                                            v3[(v76 / (((^(v76))) | 1))] = (fi39 - (fi39 ^ fi39))
                                                                            v75 -= ((((v75 * (0.0 / 1.0)) + (v75 * 1e10)) - (v75 / v75)) / float64(float32(v75)))
                                                                            _ = v136
                                                                        }
                                                                        if ((fi39 ^ int(v13)) <= (fi73 - fi73)) {
                                                                            switch fi39 {
                                                                            case 2147483647:
                                                                                v76 = v38
                                                                            case 691:
                                                                                var v137 float64 = v75
                                                                                delete(v4, "a")
                                                                                v49 = v49
                                                                                _ = v137
                                                                            case 4294967295:
                                                                                _ = len(v5)
                                                                            }
                                                                            if false {
                                                                                var v138 float32 = 1.5
                                                                                v76 = 645
                                                                                _ = v138
                                                                            }
                                                                        }
                                                                        fmt.Println(v49, v38)
                                                                        v49 = (v49 + (^(v49)))
                                                                    } else {
                                                                        v13 = v13
                                                                        var v139 float32 = 3.14159
                                                                        v49 = (v49 << ((42) & 63))
                                                                        _ = v139
                                                                    }
                                                                }
                                                            case 128:
                                                                v13 = v13
                                                            default:
                                                                if (!(-1 >= -1)) {
                                                                    if false {
                                                                        if (256 != 53) {
                                                                            if true {
                                                                                v38 += (v38 & (v38 << ((712) & 63)))
                                                                            }
                                                                            delete(v4, "foo bar")
                                                                        }
                                                                        var v140 string = "\n\t"
                                                                        switch 1 {
                                                                        case -284:
                                                                            if (v13 < 1) {
                                                                                fmt.Println(v6, v49, v75)
                                                                                if false {
                                                                                    switch (v140 + ("foo bar" + (v140 + v140))) {
                                                                                    case "hello":
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                    case "a":
                                                                                        v3[(^(v38))] = fi73
                                                                                    case "\n\t":
                                                                                        fmt.Println(v76)
                                                                                        var v141 int32 = 0
                                                                                        var v142 uint16 = 1
                                                                                        v143 := v4["\n\t"]
                                                                                        _ = v141
                                                                                        _ = v142
                                                                                        _ = v143
                                                                                    default:
                                                                                        var v144 float64 = v75
                                                                                        _ = v144
                                                                                    }
                                                                                } else {
                                                                                    switch fi73 {
                                                                                    case 0:
                                                                                        var v145 int64 = (v6 / (((-(v6))) | 1))
                                                                                        _ = v145
                                                                                    case 1:
                                                                                        _ = len(v2)
                                                                                    default:
                                                                                        var v146 uint32 = v76
                                                                                        _ = v146
                                                                                    }
                                                                                    if false {
                                                                                        switch v140 {
                                                                                        case "hello":
                                                                                            v26 = 38
                                                                                            if (!(!(3.14159 > 1e10))) {
                                                                                                switch -1 {
                                                                                                case 1:
                                                                                                    var v147 uint16 = 255
                                                                                                    _ = v147
                                                                                                default:
                                                                                                    if (968 < 2) {
                                                                                                        v140 = (v140 + v140)
                                                                                                        v148 := GenS17{F0: 2, F1: v140, F2: fi73}
                                                                                                        _ = v148
                                                                                                    } else {
                                                                                                        fmt.Println(v76, v6)
                                                                                                        delete(v4, (v140 + ((v140 + ("foo bar" + v140)) + v140)))
                                                                                                    }
                                                                                                }
                                                                                                _ = len(v5)
                                                                                            } else {
                                                                                                v13 = v13
                                                                                                if (((0.1 <= 0.0) && true) && true) {
                                                                                                    if len(v5) > 0 {
                                                                                                        v5[0] = v140
                                                                                                    }
                                                                                                } else {
                                                                                                    delete(v4, "\n\t")
                                                                                                    var v149 int16 = v13
                                                                                                    _ = len(v5)
                                                                                                    v38 += uint32(v49)
                                                                                                    _ = v149
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v2 = append(v2, fi73)
                                                                                    } else {
                                                                                        v49 = v49
                                                                                        var v150 uint32 = v76
                                                                                        var v151 uint8 = v49
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = 2147483647
                                                                                        }
                                                                                        _ = v150
                                                                                        _ = v151
                                                                                    }
                                                                                    if true {
                                                                                        v152 := v3[v76]
                                                                                        var v153 uint = 255
                                                                                        _ = v152
                                                                                        _ = v153
                                                                                    }
                                                                                    v154 := GenS2{F0: 4294967295, F1: false, F2: 18446744073709551615}
                                                                                    _ = v154
                                                                                }
                                                                            } else {
                                                                                switch (^(fi39)) {
                                                                                case 4294967295:
                                                                                    var v155 uint8 = 1
                                                                                    _ = v155
                                                                                default:
                                                                                    switch ("foo bar" + v140) {
                                                                                    case "hello":
                                                                                        if (((v140 + "hello") < "foo bar") && false) {
                                                                                            fmt.Println(v75)
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v38 += (v76 << ((uint(v49)) & 63))
                                                                                    default:
                                                                                        _ = len(v5)
                                                                                    }
                                                                                }
                                                                            }
                                                                        case 2147483647:
                                                                            v26 = v26
                                                                        case 32768:
                                                                            v156 := GenS8{F0: 32767, F1: (v1 & (v6 - v1))}
                                                                            _ = v156
                                                                        default:
                                                                            if true {
                                                                                var v157 int64 = v6
                                                                                if (false && ((v6 / (((v6 ^ (v157 + v1))) | 1)) >= 0)) {
                                                                                    var v158 uint64 = 86
                                                                                    v38 = (v38 | (^(v76)))
                                                                                    if true {
                                                                                        v159 := GenS18{F0: -194, F1: 1, F2: v49}
                                                                                        _ = v159
                                                                                    } else {
                                                                                        var v160 uint64 = (v158 << ((65535) & 63))
                                                                                        _ = len(v3)
                                                                                        fmt.Println(v38)
                                                                                        _ = v160
                                                                                    }
                                                                                    _ = v158
                                                                                } else {
                                                                                    v76 -= v38
                                                                                    if ((((v26 / ((v26) | 1)) != v26) || (true || false)) && false) {
                                                                                        fmt.Println(v26, v49, v157)
                                                                                        fmt.Println(v75, v26)
                                                                                        var v161 uint64 = 2
                                                                                        var v162 float32 = 1.5
                                                                                        if true {
                                                                                            fmt.Println(v140, fi73)
                                                                                            var v163 int8 = (v26 % ((100) | 1))
                                                                                            if (814 > 0) {
                                                                                                var v164 bool = false
                                                                                                v157 -= (v1 >> ((genFn0(v161)) & 63))
                                                                                                var v165 int64 = v6
                                                                                                _ = v164
                                                                                                _ = v165
                                                                                            } else {
                                                                                                v13 = (v13 * v13)
                                                                                                v166 := GenS2{F0: v161, F1: ((fi73 * fi73) == fi73), F2: 58}
                                                                                                _ = v166
                                                                                            }
                                                                                            switch fi73 {
                                                                                            case 474:
                                                                                                switch (v140 + "") {
                                                                                                case "a":
                                                                                                    _ = len(v3)
                                                                                                case "hello":
                                                                                                    delete(v3, (v38 / ((v76) | 1)))
                                                                                                }
                                                                                            case -9223372036854775808:
                                                                                                if false {
                                                                                                    v76 = v76
                                                                                                    fmt.Println(v49, v75, v49)
                                                                                                } else {
                                                                                                    v167 := GenS6{F0: 1, F1: fi39}
                                                                                                    if true {
                                                                                                        if (42 <= 4294967295) {
                                                                                                            var v168 int32 = -1
                                                                                                            var v169 uint16 = 2
                                                                                                            var v170 bool = true
                                                                                                            _ = v168
                                                                                                            _ = v169
                                                                                                            _ = v170
                                                                                                        } else {
                                                                                                            v13 += -383
                                                                                                            if false {
                                                                                                                v171 := GenS19{F0: 615}
                                                                                                                fmt.Println(v38, v49)
                                                                                                                v157 = (v157 % (((v157 << ((1) & 63))) | 1))
                                                                                                                switch "hello" {
                                                                                                                case "foo bar":
                                                                                                                    v172 := GenS13{F0: (v140 + v140), F1: v13}
                                                                                                                    _ = v172
                                                                                                                case "a":
                                                                                                                    var v173 string = (v140 + v140)
                                                                                                                    _ = v173
                                                                                                                case "hello":
                                                                                                                    var v174 int = fi73
                                                                                                                    _ = v174
                                                                                                                default:
                                                                                                                    var v175 uint = 65535
                                                                                                                    _ = v175
                                                                                                                }
                                                                                                                _ = v171
                                                                                                            } else {
                                                                                                                var v176 uint8 = 255
                                                                                                                v75 = ((-2.25 + v75) / ((v75 + v75) + v75))
                                                                                                                v162 = (-2.25 / ((0.0 / v162) * v162))
                                                                                                                _ = v176
                                                                                                            }
                                                                                                            fmt.Println(v26, fi39)
                                                                                                        }
                                                                                                        v157 = v6
                                                                                                        if (((true && false) && true) && false) {
                                                                                                            fmt.Println(v6)
                                                                                                            var v177 int = -289
                                                                                                            switch v140 {
                                                                                                            case "hello":
                                                                                                                v5 = append(v5, v140)
                                                                                                            case "":
                                                                                                                var v178 uint8 = (v49 ^ 127)
                                                                                                                _ = v178
                                                                                                            case "a":
                                                                                                                var v179 int = fi39
                                                                                                                _ = v179
                                                                                                            default:
                                                                                                                v180 := GenS20{F0: 2, F1: uint64(v76)}
                                                                                                                _ = v180
                                                                                                            }
                                                                                                            _ = v177
                                                                                                        }
                                                                                                    }
                                                                                                    v5 = append(v5, v140)
                                                                                                    _ = v167
                                                                                                }
                                                                                            default:
                                                                                                var v181 int8 = (v163 << ((18446744073709551615) & 63))
                                                                                                _ = v181
                                                                                            }
                                                                                            v26 += (v26 * 0)
                                                                                            _ = v163
                                                                                        } else {
                                                                                            var v182 uint32 = 2147483648
                                                                                            _ = v182
                                                                                        }
                                                                                        _ = v161
                                                                                        _ = v162
                                                                                    }
                                                                                }
                                                                                var v183 int32 = -2147483648
                                                                                _ = v157
                                                                                _ = v183
                                                                            }
                                                                        }
                                                                        if (!(fi39 == fi73)) {
                                                                            if len(v5) > 0 {
                                                                                v5[0] = v140
                                                                            }
                                                                        }
                                                                        _ = v140
                                                                    } else {
                                                                        var v184 uint32 = (v76 & 4294967295)
                                                                        v38 -= v184
                                                                        var v185 int8 = v26
                                                                        switch ("" + "a") {
                                                                        case "a":
                                                                            v6 = v6
                                                                            v38 = (^(v76))
                                                                        default:
                                                                            var v186 uint = 23
                                                                            _ = v186
                                                                        }
                                                                        _ = v184
                                                                        _ = v185
                                                                    }
                                                                    v76 = v38
                                                                    v187 := v3[(v38 * v38)]
                                                                    v188 := GenS13{F0: "", F1: v13}
                                                                    _ = v187
                                                                    _ = v188
                                                                } else {
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    v189 := GenS6{F0: -2147483648, F1: fi73}
                                                                    v190 := GenS12{F0: fi39}
                                                                    _ = v189
                                                                    _ = v190
                                                                }
                                                            }
                                                        } else {
                                                            if false {
                                                                if (!(true && false)) {
                                                                    fmt.Println(fi73, fi39)
                                                                } else {
                                                                    fmt.Println(v6)
                                                                    v1 = (v6 | 9223372036854775807)
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        if true {
                                                                            fmt.Println(fi73)
                                                                            v76 += v38
                                                                            if ((v26 | 100) > v26) {
                                                                                var v191 int = (fi73 / (((fi73 * (fi73 & fi39))) | 1))
                                                                                if true {
                                                                                    fmt.Println(fi39, v13)
                                                                                    var v192 int = (fi73 << ((0) & 63))
                                                                                    if (false || (((((v1 >= (v6 - (v1 - v6))) && (v26 == 58)) || (float64((-2.25 - 1e10)) > (-2.25 + (((v75 - v75) + 1.5) + v75)))) && false) || false)) {
                                                                                        var v193 uint64 = uint64(v76)
                                                                                        _ = v193
                                                                                    }
                                                                                    _ = v192
                                                                                } else {
                                                                                    v194 := GenS15{F0: ""}
                                                                                    _ = v194
                                                                                }
                                                                                _ = v191
                                                                            } else {
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = fi73
                                                                                }
                                                                            }
                                                                            if false {
                                                                                var v195 uint8 = 223
                                                                                v196 := v4["a"]
                                                                                var v197 int8 = v26
                                                                                _ = v195
                                                                                _ = v196
                                                                                _ = v197
                                                                            } else {
                                                                                v198 := GenS6{F0: 0, F1: (^(fi39))}
                                                                                var v199 uint8 = v49
                                                                                _ = len(v3)
                                                                                switch (fi73 & fi73) {
                                                                                case 32768:
                                                                                    var v200 float64 = v75
                                                                                    _ = v200
                                                                                }
                                                                                if false {
                                                                                    v6 += int64(v13)
                                                                                    var v201 string = "\n\t"
                                                                                    fmt.Println(v199, v75, v75)
                                                                                    _ = v201
                                                                                }
                                                                                _ = v198
                                                                                _ = v199
                                                                            }
                                                                        }
                                                                    case "":
                                                                        if true {
                                                                            v38 = 2
                                                                        }
                                                                    default:
                                                                        if (false || ((-(v26)) <= -72)) {
                                                                            switch -1 {
                                                                            case -1:
                                                                                v202 := GenS17{F0: 1, F1: "\n\t", F2: (fi73 + (fi73 + fi73))}
                                                                                _ = v202
                                                                            case 4294967295:
                                                                                v76 = (v38 / ((65535) | 1))
                                                                            default:
                                                                                if true {
                                                                                    fmt.Println(v38, v38)
                                                                                    delete(v4, "")
                                                                                }
                                                                            }
                                                                            v26 = (v26 * (^(v26)))
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                var v203 string = "hello"
                                                                _ = v203
                                                            }
                                                            fmt.Println(v6, fi39)
                                                            v38 += v76
                                                            var v204 int8 = v26
                                                            var v205 uint8 = 128
                                                            _ = v204
                                                            _ = v205
                                                        }
                                                        v38 -= (^(v76))
                                                    } else {
                                                        v2 = append(v2, 127)
                                                        var v206 uint = 215
                                                        _ = v206
                                                    }
                                                }
                                            }
                                            _ = v75
                                            _ = v76
                                        }
                                        _ = fi73
                                    }
                                case "foo bar":
                                    for fi207 := 0; fi207 < 18; fi207++ {
                                        v208 := GenS6{F0: 1, F1: (fi207 / (((^(fi39))) | 1))}
                                        v6 += (v6 * v6)
                                        _ = fi207
                                        _ = v208
                                    }
                                }
                            } else {
                                v209 := GenS16{F0: (-(v13)), F1: true}
                                v210 := GenS21{F0: v49, F1: v1, F2: 2147483647}
                                _ = v209
                                _ = v210
                            }
                        }
                        _ = len(v4)
                        _ = v49
                    }
                } else {
                    v211 := GenS20{F0: 823, F1: 281}
                    _ = v211
                }
            } else {
                var v212 uint = 683
                if (v26 != 0) {
                    for fi213 := 0; fi213 < 12; fi213++ {
                        var v214 float32 = ((1.5 * -2.25) + 0.1)
                        if true {
                            delete(v4, ("a" + "a"))
                        } else {
                            fmt.Println(fi213, v13, v38)
                        }
                        _ = fi213
                        _ = v214
                    }
                    v6 = (v6 ^ (v1 ^ (^(v6))))
                    v215 := GenS13{F0: "", F1: v13}
                    _ = v215
                }
                _ = v212
            }
            _ = fi39
        }
        _ = v37
        _ = v38
    } else {
        _ = len(v5)
        v216 := GenS22{F0: 1}
        for fi217 := 0; fi217 < 20; fi217++ {
            v1 -= (v6 % ((v1) | 1))
            var v218 uint = genFn0(6)
            if (v13 > (v13 | v13)) {
                v13 -= -1
                v219 := GenS7{F0: "a", F1: false}
                v13 -= (-(v13))
                _ = v219
            } else {
                v13 = int16(v6)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = fi217
            _ = v218
        }
        _ = v216
    }
    var v220 int32 = 2147483647
    if (0.0 == (-2.25 / (((1.5 / (1e10 * float32(-2.25))) / 1.5) * -2.25))) {
        for fi221 := 0; fi221 < 19; fi221++ {
            for fi222 := 0; fi222 < 9; fi222++ {
                v2 = append(v2, (fi222 - int(v13)))
                _ = fi222
            }
            _ = fi221
        }
    } else {
        var v223 int32 = (v220 - 2147483647)
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
        _ = v223
    }
    v4["a"] = 0
    for fi224 := 0; fi224 < 6; fi224++ {
        var v225 uint16 = 42
        var v226 string = (("a" + "a") + "a")
        _ = fi224
        _ = v225
        _ = v226
    }
    switch -355 {
    case 32768:
        var v227 float32 = (0.0 + -2.25)
        _ = v227
    case -2147483648:
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
    }
    var v228 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v220:", v220)
    fmt.Println("v228:", v228)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v220
    _ = v228
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
