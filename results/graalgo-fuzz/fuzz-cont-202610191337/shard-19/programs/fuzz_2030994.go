package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int32; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 uint16; F2 string }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 uint64; F1 int32; F2 int32 }
type GenS8 struct { F0 int; F1 string }
type GenS9 struct { F0 string }
type GenS10 struct { F0 bool; F1 int32; F2 int64 }
type GenS11 struct { F0 string; F1 int64; F2 uint64 }
func genFn0(p0 bool, p1 uint8, p2 string) int {
    p2 = p2
    for fi6 := 0; fi6 < 6; fi6++ {
        p0 = ("a" == p2)
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}
func genFn1(p0 uint16, p1 int) bool {
    var v7 int16 = 998
    switch int(p0) {
    case -1:
        switch "foo bar" {
        case "":
            if (-2147483648 == -2147483648) {
                for fi8 := 0; fi8 < 19; fi8++ {
                    p0 -= p0
                    _ = fi8
                }
            }
        case "a":
            p0 = p0
        default:
            switch ("" + "a") {
            case "\n\t":
                for fi9 := 0; fi9 < 2; fi9++ {
                    fmt.Println(p1, v7, p1)
                    for fi10 := 0; fi10 < 9; fi10++ {
                        switch genFn0((false && (false || false)), 8, "a") {
                        case 42:
                            v11 := []int{(fi9 / (((p1 - fi10)) | 1))}
                            _ = v11
                        }
                        p0 = (p0 ^ p0)
                        switch (p1 % ((fi9) | 1)) {
                        case 636:
                            v12 := []int{fi10, 4294967295}
                            _ = v12
                        case 128:
                            v7 = v7
                        }
                        _ = fi10
                    }
                    _ = fi9
                }
            case "a":
                p0 += p0
            }
        }
    case -2147483648:
        v7 -= 0
    default:
        fmt.Println(v7, p0)
    }
    _ = p0
    _ = p1
    _ = v7
    return false
}

func main() {
    var v1 int16 = -32768
    v2 := []string{(("hello" + (("a" + (("a" + "foo bar") + "a")) + ("foo bar" + ("foo bar" + (("foo bar" + (("\n\t" + (("a" + (("a" + ("foo bar" + "hello")) + ("a" + ("" + "")))) + "hello")) + "")) + ""))))) + ("\n\t" + ((("a" + "") + "hello") + "a"))), "\n\t", "hello", "foo bar"}
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]string)
    var v5 float64 = ((((0.0 + 1e10) - 1.5) * 0.1) + 0.1)
    var v13 string = ""
    if (!false) {
        v14 := GenS1{F0: 722}
        if (!false) {
            for fi15 := 0; fi15 < 4; fi15++ {
                fmt.Println(fi15, fi15, v1)
                fmt.Println(v13, v5)
                _ = fi15
            }
            v16 := GenS1{F0: 825}
            switch v13 {
            case "\n\t":
                switch 1 {
                case 1:
                    for fi17 := 0; fi17 < 2; fi17++ {
                        v18 := GenS1{F0: 18446744073709551615}
                        switch v13 {
                        case "foo bar":
                            var v19 float64 = (v5 * 0.0)
                            _ = v19
                        }
                        _ = fi17
                        _ = v18
                    }
                }
            case "hello":
                switch genFn0(false, 128, ("" + "")) {
                case 32768:
                    v2 = append(v2, ("foo bar" + v13))
                case 544:
                    v2 = append(v2, (v13 + (v13 + "a")))
                case 127:
                    fmt.Println(v13, v1, v5)
                }
            case "":
                if (481 != -2147483648) {
                    v20 := GenS2{F0: true, F1: -1, F2: 1}
                    var v21 float32 = (((-2.25 * 3.14159) / float32(-2.25)) * 0.1)
                    var v22 uint64 = 2
                    _ = v20
                    _ = v21
                    _ = v22
                }
            }
            _ = v16
        } else {
            v23 := GenS2{F0: (true && false), F1: 480, F2: 0}
            v24 := GenS3{F0: 0}
            _ = v23
            _ = v24
        }
        _ = v14
    } else {
        for fi25 := 0; fi25 < 6; fi25++ {
            var v26 float32 = -2.25
            var v27 uint64 = 271
            for fi28 := 0; fi28 < 3; fi28++ {
                if len(v2) > 0 {
                    v2[0] = (v13 + (v13 + "\n\t"))
                }
                switch genFn0((!true), 42, ((((v13 + v13) + "hello") + "hello") + "a")) {
                case 691:
                    v29 := GenS3{F0: 0}
                    _ = v29
                }
                _ = fi28
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        if false {
            switch "" {
            case "hello":
                v1 = v1
            }
            v13 = v13
            for fi30 := 0; fi30 < 5; fi30++ {
                for fi31 := 0; fi31 < 4; fi31++ {
                    v5 -= ((v5 / (v5 - (1e10 - v5))) + v5)
                    v2 = append(v2, "foo bar")
                    if (int64(v1) < 0) {
                        var v32 uint = 606
                        v32 -= (v32 / ((v32) | 1))
                        v3[2147483648] = (fi30 % (((-(fi31))) | 1))
                        v32 = (v32 - v32)
                        _ = v32
                    } else {
                        if false {
                            v33 := GenS4{F0: (fi31 << ((18446744073709551615) & 63))}
                            v2 = append(v2, (v13 + "foo bar"))
                            _ = v33
                        }
                        fmt.Println(v5)
                        var v34 int64 = -375
                        _ = v34
                    }
                    _ = fi31
                }
                fmt.Println(v5, fi30, fi30)
                if (!true) {
                    var v35 int16 = (v1 + (v1 - (v1 % ((v1) | 1))))
                    v3[4294967295] = fi30
                    for fi36 := 0; fi36 < 3; fi36++ {
                        if true {
                            v37 := GenS5{F0: 1, F1: 2, F2: ("" + (v13 + "foo bar"))}
                            var v38 bool = (true && true)
                            if v38 {
                                v39 := GenS5{F0: 485, F1: 65535, F2: (v13 + "hello")}
                                _ = v39
                            } else {
                                if (-128 > int8(fi30)) {
                                    var v40 uint8 = 0
                                    fmt.Println(v13, fi36, v1)
                                    v41 := GenS2{F0: (!(18446744073709551615 != 0)), F1: -2147483648, F2: 373}
                                    _ = v40
                                    _ = v41
                                }
                                var v42 float32 = 0.0
                                v38 = (!v38)
                                _ = v42
                            }
                            _ = v37
                            _ = v38
                        } else {
                            var v43 int16 = -32768
                            if len(v2) > 0 {
                                v2[0] = v13
                            }
                            v5 = (v5 * 0.0)
                            _ = v43
                        }
                        _ = fi36
                    }
                    _ = v35
                } else {
                    _ = len(v3)
                    if false {
                        v13 = v13
                    }
                    v1 -= (v1 + (v1 * (^(v1))))
                }
                _ = fi30
            }
        } else {
            switch "hello" {
            case "foo bar":
                v1 += v1
            default:
                v2 = append(v2, ("\n\t" + "a"))
            }
            delete(v4, 1)
            var v44 bool = true
            _ = v44
        }
    }
    v1 = (-(v1))
    delete(v4, 65535)
    v1 -= (v1 - 0)
    fmt.Println(v1)
    if true {
        v5 = v5
        v5 = v5
    } else {
        _ = len(v2)
        if true {
            v45 := v3[896]
            fmt.Println(v13)
            _ = v45
        }
    }
    switch v13 {
    case "hello":
        _ = len(v3)
    case "a":
        fmt.Println(v13)
    default:
        v3[2147483647] = 42
    }
    var v46 int16 = 1000
    if (!false) {
        var v47 float32 = 0.1
        var v48 string = (("" + (((v13 + v13) + v13) + v13)) + v13)
        for fi49 := 0; fi49 < 8; fi49++ {
            fmt.Println(v13, v5, v13)
            if (!false) {
                for fi50 := 0; fi50 < 20; fi50++ {
                    var v51 int16 = v46
                    var v52 float32 = float32(((1e10 / v5) - -2.25))
                    if false {
                        fmt.Println(v52)
                        v53 := v3[2147483647]
                        fmt.Println(v52)
                        if (!true) {
                            var v54 int16 = -32768
                            var v55 bool = (1 > 4294967295)
                            var v56 uint = 65535
                            _ = v54
                            _ = v55
                            _ = v56
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v57 := GenS5{F0: 662, F1: 999, F2: v48}
                            _ = v57
                        }
                        _ = v53
                    } else {
                        var v58 int64 = 0
                        var v59 uint64 = 18446744073709551615
                        fmt.Println(v51, v5)
                        _ = v58
                        _ = v59
                    }
                    _ = fi50
                    _ = v51
                    _ = v52
                }
                v60 := GenS3{F0: 1}
                _ = v60
            }
            for fi61 := 0; fi61 < 12; fi61++ {
                var v62 uint8 = 0
                var v63 uint16 = 886
                _ = fi61
                _ = v62
                _ = v63
            }
            _ = fi49
        }
        _ = v47
        _ = v48
    } else {
        switch (v13 + v13) {
        case "":
            v4[453] = v13
        case "hello":
            for fi64 := 0; fi64 < 17; fi64++ {
                if (!(!(!(0 == 1)))) {
                    var v65 uint16 = 42
                    fmt.Println(v5)
                    _ = v65
                }
                _ = fi64
            }
        default:
            fmt.Println(v5)
        }
    }
    v66 := GenS1{F0: 108}
    if (255 <= 42) {
        for fi67 := 0; fi67 < 3; fi67++ {
            var v68 int8 = -1
            v3[935] = 42
            switch fi67 {
            case -224:
                var v69 bool = true
                _ = v69
            case -2147483648:
                if true {
                    for fi70 := 0; fi70 < 8; fi70++ {
                        var v71 uint16 = 2
                        if ((-(fi67)) <= -498) {
                            v1 += (v1 ^ 1000)
                            var v72 bool = ((false && (1 == 42)) || true)
                            _ = v72
                        } else {
                            if false {
                                if false {
                                    if (65535 <= 2) {
                                        v73 := GenS4{F0: (fi70 | (fi67 & fi67))}
                                        var v74 string = (("\n\t" + "hello") + "")
                                        _ = v73
                                        _ = v74
                                    } else {
                                        if (true || (70 < 897)) {
                                            switch (^(fi70)) {
                                            case 81:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            case 4294967295:
                                                v68 -= (v68 % (((v68 / (((v68 >> ((65535) & 63))) | 1))) | 1))
                                            default:
                                                switch v13 {
                                                case "":
                                                    if true {
                                                        v13 = "\n\t"
                                                        switch 1 {
                                                        case -2147483648:
                                                            var v75 uint8 = 73
                                                            _ = v75
                                                        case 32768:
                                                            v76 := GenS6{F0: (728 == 0), F1: (^(v71))}
                                                            _ = v76
                                                        case 9223372036854775807:
                                                            fmt.Println(fi70)
                                                        }
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = "hello"
                                                        }
                                                        switch fi70 {
                                                        case 127:
                                                            var v77 int16 = v46
                                                            _ = v77
                                                        case -1:
                                                            v71 = (v71 & (v71 % ((v71) | 1)))
                                                        case 9:
                                                            if (false && true) {
                                                                if true {
                                                                    fmt.Println(fi70, fi67)
                                                                    v1 = v46
                                                                    v5 = v5
                                                                }
                                                                v5 = v5
                                                            }
                                                        }
                                                    }
                                                case "a":
                                                    if (((v1 << ((18446744073709551615) & 63)) > (-(v46))) && ((v5 * float64(0.1)) <= 0.1)) {
                                                        if (false && true) {
                                                            var v78 uint8 = 177
                                                            var v79 int8 = 0
                                                            if true {
                                                                var v80 int = (-(fi67))
                                                                _ = v80
                                                            } else {
                                                                if (true && true) {
                                                                    if ((v78 + (v78 * v78)) == 1) {
                                                                        if (-1 > (^(v46))) {
                                                                            v4[1] = "foo bar"
                                                                            switch v13 {
                                                                            case "hello":
                                                                                v68 = v79
                                                                            }
                                                                            var v81 uint16 = 1
                                                                            var v82 string = v13
                                                                            fmt.Println(v13, v82, v13)
                                                                            _ = v81
                                                                            _ = v82
                                                                        } else {
                                                                            v2 = append(v2, "\n\t")
                                                                            var v83 uint16 = (^(v71))
                                                                            _ = v83
                                                                        }
                                                                        switch v13 {
                                                                        case "\n\t":
                                                                            switch (v13 + v13) {
                                                                            case "hello":
                                                                                v84 := GenS7{F0: 203, F1: 42, F2: 457}
                                                                                v85 := GenS5{F0: 575, F1: (v71 >> ((4294967295) & 63)), F2: (v13 + "foo bar")}
                                                                                _ = v84
                                                                                _ = v85
                                                                            default:
                                                                                v13 = ("hello" + v13)
                                                                            }
                                                                        case "foo bar":
                                                                            v78 = v78
                                                                        case "hello":
                                                                            v79 = -128
                                                                        }
                                                                    } else {
                                                                        var v86 int32 = 42
                                                                        if false {
                                                                            v87 := GenS6{F0: (!true), F1: 1}
                                                                            v86 += 42
                                                                            _ = v87
                                                                        }
                                                                        _ = v86
                                                                    }
                                                                    var v88 uint8 = (v78 | (v78 & 42))
                                                                    v79 = (v68 - (v68 << ((275) & 63)))
                                                                    switch fi70 {
                                                                    case -2147483648:
                                                                        v89 := GenS8{F0: (fi70 + 1), F1: ("a" + "a")}
                                                                        _ = v89
                                                                    case 4294967295:
                                                                        fmt.Println(v79, v5)
                                                                    }
                                                                    if true {
                                                                        var v90 float32 = (0.0 * (((((3.14159 - ((3.14159 + 0.1) + 1.5)) - 1.5) / ((3.14159 + 0.1) - -2.25)) - 3.14159) + 0.1))
                                                                        v1 += 1
                                                                        _ = v90
                                                                    }
                                                                    _ = v88
                                                                }
                                                                if false {
                                                                    if true {
                                                                        v1 = v1
                                                                    } else {
                                                                        if false {
                                                                            v68 -= v68
                                                                        }
                                                                        fmt.Println(v68)
                                                                        if (true && false) {
                                                                            v13 = "hello"
                                                                            delete(v3, 1)
                                                                            v91 := GenS9{F0: v13}
                                                                            _ = v91
                                                                        }
                                                                    }
                                                                    switch genFn0((true && false), (v78 ^ v78), v13) {
                                                                    case -272:
                                                                        v5 -= 0.0
                                                                    case 127:
                                                                        v46 = 1000
                                                                    }
                                                                } else {
                                                                    var v92 int32 = -1
                                                                    if ((v1 >> ((730) & 63)) >= int16(v71)) {
                                                                        var v93 uint = 0
                                                                        _ = v93
                                                                    }
                                                                    fmt.Println(v78)
                                                                    _ = v92
                                                                }
                                                            }
                                                            _ = v78
                                                            _ = v79
                                                        }
                                                        var v94 int = (fi70 | genFn0(true, 63, v13))
                                                        _ = v94
                                                    } else {
                                                        v5 += 0.0
                                                    }
                                                }
                                            }
                                        }
                                        v13 = v13
                                        _ = len(v2)
                                    }
                                    var v95 float32 = 3.14159
                                    _ = len(v2)
                                    _ = v95
                                } else {
                                    v1 = 42
                                    v71 -= v71
                                    _ = len(v2)
                                }
                                fmt.Println(v68, v46)
                                v5 = 0.1
                                var v96 int16 = v1
                                if ((^(v71)) > v71) {
                                    var v97 string = (v13 + (v13 + ((v13 + v13) + ("\n\t" + "hello"))))
                                    switch v97 {
                                    case "":
                                        v5 += 1e10
                                    default:
                                        var v98 uint16 = (^(v71))
                                        _ = v98
                                    }
                                    if false {
                                        v1 = (-(v96))
                                    }
                                    _ = v97
                                }
                                _ = v96
                            }
                            v71 += 255
                            if false {
                                if false {
                                    var v99 int32 = int32(v68)
                                    var v100 bool = false
                                    _ = v100
                                    _ = v99
                                } else {
                                    fmt.Println(v68, v46, v5)
                                }
                                v46 -= (v1 * 42)
                                fmt.Println(v46)
                                if true {
                                    if len(v2) > 0 {
                                        v2[0] = "foo bar"
                                    }
                                    v2 = append(v2, v13)
                                    if (!(!false)) {
                                        if false {
                                            if true {
                                                v4[42] = v13
                                                var v101 int = (^(fi70))
                                                _ = v101
                                            } else {
                                                v2 = append(v2, (v13 + (v13 + v13)))
                                                fmt.Println(v1, v5)
                                                v102 := GenS4{F0: (fi70 | 32768)}
                                                _ = v102
                                            }
                                            switch genFn0((!false), 1, v13) {
                                            case -126:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            case 996:
                                                if (((!((v68 * (v68 >> ((65535) & 63))) >= v68)) || false) && false) {
                                                    if true {
                                                        fmt.Println(v46, v68)
                                                        switch v13 {
                                                        case "a":
                                                            var v103 int32 = -2147483648
                                                            _ = v103
                                                        case "hello":
                                                            v13 = v13
                                                        }
                                                        v13 = v13
                                                    }
                                                    fmt.Println(v13, v68, v13)
                                                }
                                            default:
                                                if (!(42 <= -1)) {
                                                    if ((("foo bar" + "hello") + (v13 + "foo bar")) <= v13) {
                                                        var v104 uint8 = 128
                                                        v104 += 53
                                                        switch (fi70 ^ (fi67 | 9223372036854775807)) {
                                                        case -232:
                                                            v105 := GenS9{F0: v13}
                                                            if false {
                                                                v106 := GenS4{F0: genFn0(true, (v104 / ((v104) | 1)), (((v13 + v13) + "hello") + v13))}
                                                                var v107 bool = (float32(v5) >= (1e10 * float32(-2.25)))
                                                                _ = v106
                                                                _ = v107
                                                            }
                                                            _ = v105
                                                        case -991:
                                                            var v108 uint32 = 65535
                                                            _ = v108
                                                        case 9223372036854775807:
                                                            v68 = v68
                                                        }
                                                        _ = v104
                                                    }
                                                }
                                            }
                                            var v109 bool = false
                                            _ = v109
                                        } else {
                                            var v110 int = fi70
                                            _ = v110
                                        }
                                    }
                                } else {
                                    fmt.Println(v5, v68)
                                    v68 -= (v68 & v68)
                                    v3[962] = (fi70 - (fi70 / ((genFn0(false, 172, v13)) | 1)))
                                }
                                if false {
                                    v5 += v5
                                    v46 = int16(v68)
                                }
                                v2 = append(v2, ("\n\t" + (("a" + "\n\t") + (v13 + (v13 + (v13 + (v13 + v13)))))))
                                v111 := v4[65535]
                                _ = v111
                            }
                            v71 -= 1
                        }
                        delete(v4, 65535)
                        v68 = (v68 & v68)
                        _ = fi70
                        _ = v71
                    }
                    fmt.Println(fi67)
                    fmt.Println(v1, fi67)
                } else {
                    var v112 string = ((v13 + (("foo bar" + (v13 + v13)) + "foo bar")) + "")
                    if (false || (!(!(!true)))) {
                        for fi113 := 0; fi113 < 6; fi113++ {
                            if (false && false) {
                                switch v112 {
                                case "hello":
                                    v112 = (((v112 + v112) + (v13 + (v112 + v112))) + ("foo bar" + "\n\t"))
                                case "":
                                    var v114 bool = (!(false && true))
                                    _ = v114
                                default:
                                    v115 := GenS10{F0: (128 <= 128), F1: -1, F2: 0}
                                    _ = v115
                                }
                                fmt.Println(fi67, v1, v1)
                            }
                            switch (fi113 << ((255) & 63)) {
                            case 128:
                                if (int16(fi67) == v1) {
                                    v1 += v1
                                } else {
                                    _ = len(v2)
                                    fmt.Println(v1, v68, v112)
                                }
                            case 32768:
                                v116 := GenS11{F0: (("hello" + "foo bar") + v112), F1: 1, F2: 2}
                                _ = v116
                            case 32767:
                                if false {
                                    fmt.Println(fi67)
                                    var v117 uint = 0
                                    _ = v117
                                }
                            }
                            _ = fi113
                        }
                    }
                    switch (fi67 ^ (^(fi67))) {
                    case 32768:
                        v1 -= -923
                    case 0:
                        for fi118 := 0; fi118 < 3; fi118++ {
                            v112 = v13
                            v46 = v46
                            fmt.Println(fi118, v112)
                            _ = fi118
                        }
                    default:
                        fmt.Println(v112, v46)
                    }
                    _ = v112
                }
            case -501:
                v4[42] = (v13 + v13)
            }
            _ = fi67
            _ = v68
        }
        var v119 uint = 65535
        _ = v119
    }
    v46 = -32768
    fmt.Println(v46, v46, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v66
}
