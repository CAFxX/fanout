package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int8; F1 int64; F2 uint8 }
type GenS3 struct { F0 int64; F1 int64; F2 int16 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 string; F1 bool; F2 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64; F1 int8; F2 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int; F1 int8 }
type GenS10 struct { F0 int64; F1 uint; F2 int16 }
type GenS11 struct { F0 uint; F1 int64 }
type GenS12 struct { F0 bool; F1 uint64 }
type GenS13 struct { F0 uint; F1 uint8; F2 int }
type GenS14 struct { F0 int16; F1 uint32; F2 bool }
type GenS15 struct { F0 int16; F1 uint8; F2 bool }
type GenS16 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS17 struct { F0 bool; F1 string; F2 int16 }
type GenS18 struct { F0 int; F1 uint32 }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 int; F1 uint }
type GenS21 struct { F0 uint8; F1 int16 }
type GenS22 struct { F0 uint16; F1 string }
type GenS23 struct { F0 uint32 }
type GenS24 struct { F0 int16 }
func genFn0(p0 int8, p1 uint64) uint32 {
    var v6 float64 = 1e10
    var v7 int16 = -32768
    var v8 uint64 = p1
    v9 := GenS1{F0: (!false), F1: -1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return uint32(v8)
}

func main() {
    var v1 uint = 65535
    v2 := make(map[uint32]int)
    var v3 uint64 = 928
    var v4 uint8 = 127
    v5 := []int{32768}
    for fi10 := 0; fi10 < 17; fi10++ {
        v4 -= uint8(v3)
        if false {
            if true {
                switch (fi10 & fi10) {
                case -9223372036854775808:
                    var v11 float64 = (-2.25 / -2.25)
                    _ = v11
                case -2147483648:
                    for fi12 := 0; fi12 < 11; fi12++ {
                        _ = len(v5)
                        _ = fi12
                    }
                case 1:
                    var v13 int8 = -1
                    _ = v13
                }
            } else {
                switch "\n\t" {
                case "\n\t":
                    v4 = uint8(v1)
                case "foo bar":
                    var v14 uint8 = (^(v4))
                    _ = v14
                }
                fmt.Println(v3, fi10, v4)
            }
        } else {
            v3 -= (v3 * v3)
        }
        _ = fi10
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        v16 := GenS2{F0: -1, F1: 9223372036854775807, F2: uint8(v3)}
        v17 := GenS2{F0: -55, F1: int64(v4), F2: uint8(v3)}
        if false {
            switch "foo bar" {
            case "hello":
                for fi18 := 0; fi18 < 19; fi18++ {
                    if (("a" >= "a") || ((506 < -1) && (!true))) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v19 uint64 = uint64(v1)
                        fmt.Println(fi15, v4)
                        _ = v19
                    } else {
                        if ((v1 % (((v1 & (v1 - (^(v1))))) | 1)) >= v1) {
                            _ = len(v2)
                        }
                        var v20 uint = (v1 - v1)
                        if true {
                            v4 += v4
                        }
                        _ = v20
                    }
                    _ = fi18
                }
            }
            for fi21 := 0; fi21 < 15; fi21++ {
                v5 = append(v5, int(v1))
                _ = fi21
            }
            for fi22 := 0; fi22 < 6; fi22++ {
                v23 := GenS3{F0: 0, F1: -9223372036854775808, F2: int16(v1)}
                switch ((("foo bar" + "hello") + "\n\t") + ("foo bar" + (((("hello" + (("a" + "hello") + "a")) + "foo bar") + "") + ""))) {
                case "\n\t":
                    if false {
                        v24 := GenS1{F0: true, F1: -390}
                        switch fi22 {
                        case -87:
                            if (3.14159 <= (0.1 + -2.25)) {
                                switch (fi15 - 42) {
                                case 1:
                                    var v25 int = fi15
                                    _ = v25
                                case 42:
                                    var v26 uint8 = v4
                                    _ = v26
                                case -79:
                                    v5 = append(v5, -546)
                                default:
                                    v27 := GenS4{F0: 0, F1: "foo bar"}
                                    _ = v27
                                }
                                v4 -= (^(v4))
                            } else {
                                if true {
                                    switch (^(fi22)) {
                                    case 2147483647:
                                        v28 := v2[2147483648]
                                        _ = v28
                                    case -2147483648:
                                        switch fi22 {
                                        case 4294967295:
                                            v3 -= v3
                                        case 712:
                                            if (0.0 == 1.5) {
                                                switch (-(fi15)) {
                                                case 107:
                                                    v4 = 127
                                                case 42:
                                                    v3 -= 490
                                                case 9223372036854775807:
                                                    if false {
                                                        if false {
                                                            v4 -= 234
                                                        }
                                                        fmt.Println(v1)
                                                    } else {
                                                        v5 = append(v5, fi22)
                                                    }
                                                }
                                                var v29 float32 = 3.14159
                                                _ = v29
                                            }
                                        case -9223372036854775808:
                                            v1 = (v1 * (v1 - (v1 ^ v1)))
                                        }
                                    }
                                } else {
                                    if (v3 <= (v3 >> (((^(v1))) & 63))) {
                                        v30 := GenS3{F0: 0, F1: 708, F2: 32767}
                                        _ = v30
                                    } else {
                                        fmt.Println(fi22)
                                    }
                                    if (float64(0.0) > (1e10 + 0.1)) {
                                        if ((genFn0(-128, (v3 / ((0) | 1))) != 65535) && false) {
                                            delete(v2, 0)
                                            v31 := GenS1{F0: false, F1: 0}
                                            v32 := GenS2{F0: -1, F1: 0, F2: (v4 % ((v4) | 1))}
                                            _ = v31
                                            _ = v32
                                        } else {
                                            v33 := GenS5{F0: "hello", F1: (2147483648 <= 712), F2: ("foo bar" <= "foo bar")}
                                            var v34 int8 = 42
                                            _ = v33
                                            _ = v34
                                        }
                                    }
                                }
                            }
                        case 0:
                            v35 := GenS2{F0: 31, F1: 1, F2: 128}
                            _ = v35
                        default:
                            v5 = append(v5, (fi15 + -9223372036854775808))
                        }
                        _ = v24
                    } else {
                        v2[uint32(v3)] = (fi22 ^ -2147483648)
                        switch ("foo bar" + "foo bar") {
                        case "hello":
                            switch "foo bar" {
                            case "\n\t":
                                v36 := GenS2{F0: 1, F1: 9223372036854775807, F2: v4}
                                _ = v36
                            case "":
                                switch (("a" + "") + "foo bar") {
                                case "foo bar":
                                    fmt.Println(v4)
                                    if true {
                                        v1 = (v1 << ((852) & 63))
                                    }
                                    _ = len(v2)
                                case "hello":
                                    delete(v2, 2147483648)
                                }
                            default:
                                v37 := v2[434]
                                fmt.Println(v1, v37, v37)
                                v4 = (^(v4))
                                switch int(v1) {
                                case 0:
                                    if (true && (((((1.5 + 1.5) * (1.5 / 3.14159)) > (1e10 - -2.25)) && (true && false)) || false)) {
                                        v37 = fi22
                                    } else {
                                        if ("" >= "hello") {
                                            if (1 < 42) {
                                                if true {
                                                    v1 = (v1 & 18446744073709551615)
                                                } else {
                                                    v38 := v2[genFn0(-1, (v3 & v3))]
                                                    if len(v5) > 0 {
                                                        v5[0] = -1
                                                    }
                                                    var v39 int64 = -222
                                                    _ = v38
                                                    _ = v39
                                                }
                                            } else {
                                                v40 := GenS6{F0: "\n\t"}
                                                if true {
                                                    v4 += (v4 >> (((^(v1))) & 63))
                                                }
                                                if false {
                                                    v1 += 147
                                                }
                                                switch (fi22 << (((v1 / ((uint(v3)) | 1))) & 63)) {
                                                case 127:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                case 32767:
                                                    v41 := GenS7{F0: -775, F1: 42, F2: -433}
                                                    _ = v41
                                                default:
                                                    v42 := GenS8{F0: -2147483648}
                                                    _ = v42
                                                }
                                                var v43 int8 = 52
                                                var v44 int8 = v43
                                                _ = v40
                                                _ = v43
                                                _ = v44
                                            }
                                            v45 := GenS9{F0: -101, F1: 30}
                                            var v46 uint64 = 8
                                            _ = v45
                                            _ = v46
                                        } else {
                                            v1 = v1
                                            fmt.Println(fi15)
                                        }
                                    }
                                case -9223372036854775808:
                                    v47 := GenS4{F0: 0, F1: "foo bar"}
                                    _ = v47
                                case -607:
                                    v4 -= 220
                                }
                                _ = v37
                            }
                        }
                    }
                case "a":
                    v4 = v4
                default:
                    switch (-(fi22)) {
                    case 9223372036854775807:
                        v4 += 255
                    default:
                        if (-599 >= 42) {
                            var v48 uint8 = v4
                            _ = v48
                        }
                    }
                }
                switch ("\n\t" + ("hello" + ("foo bar" + ("a" + "\n\t")))) {
                case "a":
                    if (false || true) {
                        switch (fi15 - (fi15 * (fi22 & fi22))) {
                        case -656:
                            var v49 int8 = 67
                            _ = v49
                        case -975:
                            v50 := GenS7{F0: 683, F1: 123, F2: 1}
                            _ = v50
                        default:
                            v5 = append(v5, 127)
                        }
                        if false {
                            if true {
                                switch 2147483647 {
                                case 127:
                                    _ = len(v5)
                                case 957:
                                    if true {
                                        fmt.Println(v4)
                                        fmt.Println(v4, fi15)
                                        v51 := GenS10{F0: 42, F1: (v1 / ((v1) | 1)), F2: 1}
                                        _ = v51
                                    } else {
                                        fmt.Println(v3, fi22, v1)
                                        delete(v2, 742)
                                        var v52 int64 = -9223372036854775808
                                        _ = v52
                                    }
                                case -522:
                                    v53 := GenS11{F0: (v1 >> (((v1 & v1)) & 63)), F1: 1}
                                    _ = v53
                                }
                                var v54 int64 = -1
                                v3 = (v3 >> ((18446744073709551615) & 63))
                                _ = v54
                            }
                            v55 := GenS12{F0: (1 == (^(fi15))), F1: (v3 + v3)}
                            v3 = v3
                            _ = v55
                        } else {
                            switch ("\n\t" + "") {
                            case "\n\t":
                                v3 = v3
                            case "foo bar":
                                if (-190 != -2147483648) {
                                    _ = len(v5)
                                    if (776 > 32767) {
                                        if true {
                                            var v56 uint = (^(v1))
                                            _ = v56
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                } else {
                                    v57 := v2[65535]
                                    v3 -= (v3 % (((v3 << (((v1 * 1)) & 63))) | 1))
                                    if true {
                                        v58 := GenS7{F0: -25, F1: -1, F2: 988}
                                        var v59 uint32 = 0
                                        _ = v58
                                        _ = v59
                                    } else {
                                        if (v4 != (v4 >> (((v1 + (v1 >> ((v1) & 63)))) & 63))) {
                                            if ((!true) || (!true)) {
                                                var v60 int = fi22
                                                v5 = append(v5, v57)
                                                _ = v60
                                            } else {
                                                v61 := GenS11{F0: (v1 / ((527) | 1)), F1: 1}
                                                fmt.Println(v4)
                                                _ = v61
                                            }
                                            var v62 float32 = (1e10 / -2.25)
                                            v4 += (^(v4))
                                            switch ("\n\t" + "hello") {
                                            case "foo bar":
                                                v4 += (v4 / ((81) | 1))
                                            case "\n\t":
                                                switch v57 {
                                                case -2147483648:
                                                    if false {
                                                        v62 -= -2.25
                                                        fmt.Println(fi15, fi22, v62)
                                                    } else {
                                                        v4 += 0
                                                        if false {
                                                            v63 := GenS7{F0: 1, F1: int8(v3), F2: 444}
                                                            v64 := GenS2{F0: -128, F1: 9223372036854775807, F2: 128}
                                                            _ = v63
                                                            _ = v64
                                                        } else {
                                                            v65 := GenS5{F0: "\n\t", F1: true, F2: (fi15 == fi15)}
                                                            switch ((("" + "foo bar") + "") + ("" + "")) {
                                                            case "hello":
                                                                v66 := GenS13{F0: (^(v1)), F1: v4, F2: fi15}
                                                                _ = v66
                                                            case "foo bar":
                                                                v57 -= fi22
                                                            case "":
                                                                var v67 uint8 = v4
                                                                _ = v67
                                                            }
                                                            v2[806] = (^(v57))
                                                            _ = v65
                                                        }
                                                        _ = len(v2)
                                                    }
                                                case -737:
                                                    var v68 string = ("a" + ("\n\t" + "\n\t"))
                                                    _ = v68
                                                default:
                                                    v5 = append(v5, fi22)
                                                }
                                            case "hello":
                                                var v69 float64 = float64(float32(0.0))
                                                _ = v69
                                            }
                                            _ = v62
                                        } else {
                                            fmt.Println(v4)
                                            fmt.Println(v3)
                                        }
                                        if (!((0.1 + float32(3.14159)) != (1.5 - 0.0))) {
                                            fmt.Println(v3, fi22)
                                            if (127 >= 100) {
                                                fmt.Println(v3, fi22)
                                                if false {
                                                    v57 += fi15
                                                    switch (fi15 % (((^(fi15))) | 1)) {
                                                    case 737:
                                                        v1 = (v1 / ((v1) | 1))
                                                        v1 -= 2
                                                        var v70 uint = v1
                                                        v4 = (v4 >> ((4294967295) & 63))
                                                        var v71 int64 = 696
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                    switch -2147483648 {
                                                    case 127:
                                                        delete(v2, 42)
                                                    case 307:
                                                        var v72 int64 = -1
                                                        _ = v72
                                                    default:
                                                        if ((((1.5 - 1.5) * ((3.14159 + 0.0) - 1e10)) * float32(1e10)) >= 0.0) {
                                                            switch (-(v57)) {
                                                            case 2147483647:
                                                                v4 += v4
                                                            case 32767:
                                                                if (("foo bar" <= ("" + "hello")) || true) {
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        var v73 float64 = 3.14159
                                                                        _ = v73
                                                                    } else {
                                                                        var v74 string = ("foo bar" + "foo bar")
                                                                        var v75 float32 = 3.14159
                                                                        if (!false) {
                                                                            v57 = fi22
                                                                            if false {
                                                                                v4 = (v4 ^ v4)
                                                                                delete(v2, 0)
                                                                                v76 := GenS11{F0: v1, F1: -424}
                                                                                _ = v76
                                                                            } else {
                                                                                v77 := GenS14{F0: -1, F1: 4294967295, F2: false}
                                                                                if (v74 >= "") {
                                                                                    var v78 string = "a"
                                                                                    fmt.Println(v3)
                                                                                    _ = v78
                                                                                }
                                                                                var v79 float64 = 1.5
                                                                                _ = v77
                                                                                _ = v79
                                                                            }
                                                                            v1 = v1
                                                                        } else {
                                                                            if true {
                                                                                v74 = v74
                                                                                v57 = v57
                                                                            } else {
                                                                                fmt.Println(v4)
                                                                            }
                                                                        }
                                                                        _ = v74
                                                                        _ = v75
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        var v80 int16 = 32767
                                                                        _ = v80
                                                                    case "":
                                                                        var v81 uint64 = v3
                                                                        _ = v81
                                                                    default:
                                                                        var v82 int64 = -9223372036854775808
                                                                        _ = v82
                                                                    }
                                                                }
                                                            default:
                                                                fmt.Println(v4)
                                                            }
                                                        }
                                                        if true {
                                                            var v83 bool = (!true)
                                                            var v84 uint8 = v4
                                                            if len(v5) > 0 {
                                                                v5[0] = v57
                                                            }
                                                            var v85 uint16 = 32767
                                                            var v86 int8 = -47
                                                            v87 := GenS15{F0: -763, F1: (^(v84)), F2: true}
                                                            switch fi15 {
                                                            case 42:
                                                                v1 -= (v1 >> ((18446744073709551615) & 63))
                                                            }
                                                            _ = v83
                                                            _ = v84
                                                            _ = v85
                                                            _ = v86
                                                            _ = v87
                                                        } else {
                                                            _ = len(v2)
                                                            var v88 int64 = 9223372036854775807
                                                            var v89 float64 = float64((-2.25 + float32(0.0)))
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                        v90 := GenS6{F0: "foo bar"}
                                                        _ = v90
                                                    }
                                                    v1 = (v1 ^ (^(v1)))
                                                    fmt.Println(v4, v57, v4)
                                                    v91 := GenS10{F0: 1, F1: v1, F2: int16(v1)}
                                                    if ((v1 + (^(v1))) <= v1) {
                                                        v2[2] = -540
                                                        var v92 float64 = float64(3.14159)
                                                        v93 := GenS5{F0: "hello", F1: false, F2: (884 != 2147483647)}
                                                        _ = v92
                                                        _ = v93
                                                    } else {
                                                        var v94 int64 = 1
                                                        _ = v94
                                                    }
                                                    _ = v91
                                                }
                                                fmt.Println(v3, v1, v4)
                                            }
                                            switch "foo bar" {
                                            case "hello":
                                                if true {
                                                    switch ("" + "hello") {
                                                    case "a":
                                                        var v95 float64 = (-2.25 - -2.25)
                                                        _ = v95
                                                    case "\n\t":
                                                        v96 := v2[genFn0(42, v3)]
                                                        v1 -= v1
                                                        _ = v96
                                                    case "hello":
                                                        v5 = append(v5, fi15)
                                                    }
                                                    v3 = 42
                                                }
                                            case "\n\t":
                                                if true {
                                                    if false {
                                                        if true {
                                                            fmt.Println(fi15, fi15, v3)
                                                            if false {
                                                                if true {
                                                                    fmt.Println(v4, v4, v57)
                                                                    v2[991] = (fi15 >> ((v1) & 63))
                                                                    if (("" + "hello") > (("" + "hello") + "")) {
                                                                        v57 += (fi22 >> ((v1) & 63))
                                                                        var v97 bool = (118 != 108)
                                                                        _ = v97
                                                                    } else {
                                                                        v3 -= 489
                                                                    }
                                                                }
                                                                var v98 float32 = (0.0 * 0.1)
                                                                _ = v98
                                                            } else {
                                                                if false {
                                                                    if (-1 < -32768) {
                                                                        var v99 float32 = 0.0
                                                                        fmt.Println(fi22, v4, v99)
                                                                        _ = v99
                                                                    }
                                                                } else {
                                                                    v3 = v3
                                                                    if (!false) {
                                                                        fmt.Println(fi22, v4, fi15)
                                                                        v57 = (fi15 * (v57 << (((v1 | 0)) & 63)))
                                                                    } else {
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            v57 = 32767
                                                                        case "":
                                                                            switch (("hello" + "hello") + (("\n\t" + "foo bar") + "a")) {
                                                                            case "":
                                                                                if true {
                                                                                    fmt.Println(v4, v1)
                                                                                    v3 -= 730
                                                                                }
                                                                            case "a":
                                                                                var v100 int64 = 1
                                                                                _ = v100
                                                                            }
                                                                        default:
                                                                            v3 = (v3 << (((^(v1))) & 63))
                                                                        }
                                                                        var v101 uint64 = (v3 ^ 42)
                                                                        if false {
                                                                            fmt.Println(v1, fi15, v57)
                                                                            var v102 uint32 = 0
                                                                            var v103 uint16 = 377
                                                                            v101 = uint64(v4)
                                                                            if true {
                                                                                if (genFn0(100, (v101 / ((v101) | 1))) >= 42) {
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        if true {
                                                                                            if false {
                                                                                                v104 := GenS9{F0: -1, F1: 1}
                                                                                                v105 := GenS15{F0: 32767, F1: (v4 | v4), F2: (((v1 + (v1 - v1)) != (v1 + (v1 - (v1 << ((uint(v103)) & 63))))) && false)}
                                                                                                _ = v104
                                                                                                _ = v105
                                                                                            } else {
                                                                                                v3 = v3
                                                                                            }
                                                                                            v106 := GenS6{F0: "foo bar"}
                                                                                            v107 := GenS8{F0: -2147483648}
                                                                                            _ = v106
                                                                                            _ = v107
                                                                                        } else {
                                                                                            if len(v5) > 0 {
                                                                                                v5[0] = (^(fi15))
                                                                                            }
                                                                                        }
                                                                                    case "a":
                                                                                        v108 := v2[genFn0(41, 259)]
                                                                                        var v109 float64 = ((float64(3.14159) + float64(0.1)) + 0.1)
                                                                                        if (true || (-9223372036854775808 == -9223372036854775808)) {
                                                                                            var v110 float64 = 1.5
                                                                                            var v111 uint64 = (v101 << (((v1 - 819)) & 63))
                                                                                            v4 = (v4 + (v4 >> ((v1) & 63)))
                                                                                            _ = v110
                                                                                            _ = v111
                                                                                        }
                                                                                        _ = v108
                                                                                        _ = v109
                                                                                    default:
                                                                                        v112 := GenS15{F0: -611, F1: v4, F2: (-521 < -814)}
                                                                                        v113 := GenS3{F0: -1, F1: 0, F2: 32767}
                                                                                        if ((fi22 & fi15) > fi22) {
                                                                                            v114 := GenS16{F0: 32767, F1: (v103 + (v103 ^ v103)), F2: v3}
                                                                                            _ = v114
                                                                                        } else {
                                                                                            v5 = append(v5, (fi22 * (fi15 + (fi15 << (((v1 << (((v1 ^ v1)) & 63))) & 63)))))
                                                                                            if ((v103 & (^(v103))) != 32767) {
                                                                                                switch (^(fi22)) {
                                                                                                case 128:
                                                                                                    var v115 string = ""
                                                                                                    _ = v115
                                                                                                case -1:
                                                                                                    switch "a" {
                                                                                                    case "":
                                                                                                        if len(v5) > 0 {
                                                                                                            v5[0] = fi22
                                                                                                        }
                                                                                                        v1 = (^(v1))
                                                                                                        fmt.Println(v102, v3, v57)
                                                                                                    case "a":
                                                                                                        if false {
                                                                                                            _ = len(v2)
                                                                                                            var v116 int8 = -128
                                                                                                            _ = v116
                                                                                                        } else {
                                                                                                            var v117 int8 = -24
                                                                                                            var v118 uint32 = v102
                                                                                                            _ = v117
                                                                                                            _ = v118
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        v102 += genFn0(-115, v3)
                                                                                                    }
                                                                                                    var v119 int16 = 145
                                                                                                    _ = v119
                                                                                                }
                                                                                                v102 -= v102
                                                                                            }
                                                                                            switch ("" + "a") {
                                                                                            case "foo bar":
                                                                                                if false {
                                                                                                    v103 += (v103 << ((uint(v4)) & 63))
                                                                                                    fmt.Println(v1, v103)
                                                                                                } else {
                                                                                                    if (1.5 == 1e10) {
                                                                                                        _ = len(v5)
                                                                                                        v103 += 32767
                                                                                                    } else {
                                                                                                        if len(v5) >= 2 {
                                                                                                            v5 = v5[1:len(v5)]
                                                                                                        }
                                                                                                        var v120 int = fi15
                                                                                                        _ = v120
                                                                                                    }
                                                                                                    switch fi15 {
                                                                                                    case 54:
                                                                                                        var v121 float64 = (0.1 * (1e10 / 1.0))
                                                                                                        _ = v121
                                                                                                    case 1:
                                                                                                        _ = len(v2)
                                                                                                    default:
                                                                                                        var v122 uint16 = (^(v103))
                                                                                                        _ = v122
                                                                                                    }
                                                                                                    if false {
                                                                                                        v123 := GenS14{F0: 400, F1: (^(v102)), F2: ("foo bar" != "a")}
                                                                                                        if false {
                                                                                                            var v124 float32 = 1e10
                                                                                                            var v125 uint16 = 255
                                                                                                            v102 = genFn0(int8(v57), v3)
                                                                                                            if ((v101 << ((v1) & 63)) == 1) {
                                                                                                                v4 = v4
                                                                                                            } else {
                                                                                                                if true {
                                                                                                                    v125 -= 39
                                                                                                                }
                                                                                                                fmt.Println(v57)
                                                                                                                fmt.Println(v1)
                                                                                                                if true {
                                                                                                                    v126 := v2[(v102 << (((^(v1))) & 63))]
                                                                                                                    switch ("a" + "\n\t") {
                                                                                                                    case "":
                                                                                                                        var v127 uint32 = 65535
                                                                                                                        _ = v127
                                                                                                                    default:
                                                                                                                        switch "hello" {
                                                                                                                        case "a":
                                                                                                                            v103 += 32767
                                                                                                                        case "":
                                                                                                                            switch "a" {
                                                                                                                            case "\n\t":
                                                                                                                                _ = len(v2)
                                                                                                                            }
                                                                                                                        case "foo bar":
                                                                                                                            switch ("hello" + ("a" + ("\n\t" + (("hello" + "") + "\n\t")))) {
                                                                                                                            case "":
                                                                                                                                v5 = append(v5, -11)
                                                                                                                            case "a":
                                                                                                                                v1 = v1
                                                                                                                            case "\n\t":
                                                                                                                                fmt.Println(v3, v102)
                                                                                                                            default:
                                                                                                                                if len(v5) > 0 {
                                                                                                                                    v5[0] = (fi15 / ((fi22) | 1))
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v126
                                                                                                                } else {
                                                                                                                    v101 = v101
                                                                                                                    v128 := GenS12{F0: (0 <= 0), F1: (v101 + v3)}
                                                                                                                    _ = v128
                                                                                                                }
                                                                                                            }
                                                                                                            fmt.Println(v102, fi22)
                                                                                                            _ = v124
                                                                                                            _ = v125
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                _ = len(v5)
                                                                                                                v129 := v2[genFn0(42, 961)]
                                                                                                                fmt.Println(v4)
                                                                                                                _ = v129
                                                                                                            }
                                                                                                            fmt.Println(v101, v1)
                                                                                                            v130 := GenS7{F0: 0, F1: -99, F2: -1}
                                                                                                            var v131 uint = (v1 / (((v1 & 2)) | 1))
                                                                                                            _ = v130
                                                                                                            _ = v131
                                                                                                        }
                                                                                                        _ = v123
                                                                                                    } else {
                                                                                                        _ = len(v5)
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                if false {
                                                                                                    fmt.Println(v4, v102)
                                                                                                } else {
                                                                                                    v132 := GenS1{F0: true, F1: 178}
                                                                                                    _ = len(v5)
                                                                                                    _ = v132
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v3, v103)
                                                                                        v133 := GenS8{F0: -502}
                                                                                        _ = v112
                                                                                        _ = v113
                                                                                        _ = v133
                                                                                    }
                                                                                    switch "hello" {
                                                                                    case "a":
                                                                                        if true {
                                                                                            if false {
                                                                                                fmt.Println(v1)
                                                                                            } else {
                                                                                                v134 := GenS16{F0: -1, F1: 42, F2: v101}
                                                                                                _ = v134
                                                                                            }
                                                                                            if len(v5) >= 2 {
                                                                                                v5 = v5[1:len(v5)]
                                                                                            }
                                                                                            v135 := GenS17{F0: (-932 <= -2147483648), F1: "\n\t", F2: -32768}
                                                                                            _ = v135
                                                                                        }
                                                                                    default:
                                                                                        if (1 >= int64(v4)) {
                                                                                            fmt.Println(v103, v101)
                                                                                        }
                                                                                        v136 := GenS18{F0: v57, F1: uint32(v101)}
                                                                                        var v137 int8 = int8(v4)
                                                                                        _ = v136
                                                                                        _ = v137
                                                                                    }
                                                                                    if (!false) {
                                                                                        v57 = fi15
                                                                                        if (false || (1.5 >= 1e10)) {
                                                                                            v101 -= 0
                                                                                        } else {
                                                                                            v57 += (v57 * v57)
                                                                                            v138 := GenS5{F0: ("foo bar" + "a"), F1: (false && true), F2: false}
                                                                                            _ = v138
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v139 := GenS18{F0: (-(v57)), F1: (^(v102))}
                                                                                _ = v139
                                                                            }
                                                                            _ = v102
                                                                            _ = v103
                                                                        } else {
                                                                            if true {
                                                                                switch 1 {
                                                                                case 127:
                                                                                    var v140 int32 = 0
                                                                                    _ = v140
                                                                                case 4294967295:
                                                                                    var v141 bool = (4294967295 == uint32(v4))
                                                                                    _ = v141
                                                                                case -398:
                                                                                    switch "a" {
                                                                                    case "\n\t":
                                                                                        var v142 bool = (false || ((^(v1)) != v1))
                                                                                        _ = v142
                                                                                    case "":
                                                                                        var v143 bool = false
                                                                                        _ = v143
                                                                                    }
                                                                                default:
                                                                                    switch "foo bar" {
                                                                                    case "hello":
                                                                                        var v144 uint64 = v101
                                                                                        _ = v144
                                                                                    case "\n\t":
                                                                                        var v145 bool = true
                                                                                        _ = v145
                                                                                    case "":
                                                                                        if false {
                                                                                            v101 = 996
                                                                                            var v146 int = v57
                                                                                            _ = v146
                                                                                        }
                                                                                    }
                                                                                }
                                                                                switch "hello" {
                                                                                case "\n\t":
                                                                                    v5 = append(v5, 0)
                                                                                case "a":
                                                                                    v3 += (^(v101))
                                                                                case "":
                                                                                    if (111 == 0) {
                                                                                        switch "\n\t" {
                                                                                        case "hello":
                                                                                            var v147 bool = false
                                                                                            _ = v147
                                                                                        case "a":
                                                                                            v148 := GenS19{F0: 1}
                                                                                            fmt.Println(v1)
                                                                                            _ = v148
                                                                                        case "\n\t":
                                                                                            var v149 int32 = 1
                                                                                            _ = v149
                                                                                        }
                                                                                        v3 = (v3 >> (((v1 * (v1 << ((v1) & 63)))) & 63))
                                                                                        v150 := GenS20{F0: fi22, F1: (v1 + (v1 * (v1 + 42)))}
                                                                                        _ = v150
                                                                                    } else {
                                                                                        var v151 uint16 = 504
                                                                                        var v152 float64 = 1e10
                                                                                        _ = v151
                                                                                        _ = v152
                                                                                    }
                                                                                default:
                                                                                    v3 = v101
                                                                                }
                                                                            }
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                if true {
                                                                                    v3 -= 0
                                                                                } else {
                                                                                    v4 = (v4 & 255)
                                                                                    if false {
                                                                                        switch fi15 {
                                                                                        case 128:
                                                                                            v101 -= v3
                                                                                        default:
                                                                                            var v153 float64 = ((float64(0.0) - 3.14159) - -2.25)
                                                                                            _ = v153
                                                                                        }
                                                                                        fmt.Println(v101, v4, v1)
                                                                                        v154 := GenS9{F0: (fi22 + v57), F1: -1}
                                                                                        _ = v154
                                                                                    } else {
                                                                                        v3 = (v3 - 18446744073709551615)
                                                                                        var v155 float64 = 3.14159
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                        _ = v155
                                                                                    }
                                                                                    v1 -= (^(v1))
                                                                                }
                                                                                v156 := GenS15{F0: -763, F1: (^(v4)), F2: ((v1 << ((v1) & 63)) == v1)}
                                                                                var v157 float32 = 1.5
                                                                                _ = v156
                                                                                _ = v157
                                                                            case "hello":
                                                                                v3 -= 42
                                                                            }
                                                                        }
                                                                        _ = v101
                                                                    }
                                                                    var v158 float64 = (float64(0.0) * 3.14159)
                                                                    var v159 float64 = v158
                                                                    switch ("a" + "\n\t") {
                                                                    case "foo bar":
                                                                        if true {
                                                                            fmt.Println(v159)
                                                                        }
                                                                    case "hello":
                                                                        v160 := GenS7{F0: 9223372036854775807, F1: 89, F2: 1}
                                                                        _ = v160
                                                                    }
                                                                    _ = v158
                                                                    _ = v159
                                                                }
                                                                fmt.Println(v1, v57)
                                                                if true {
                                                                    v161 := GenS10{F0: -9223372036854775808, F1: (v1 + (v1 + v1)), F2: 32767}
                                                                    switch "hello" {
                                                                    case "":
                                                                        var v162 bool = (!true)
                                                                        _ = v162
                                                                    case "a":
                                                                        v4 = v4
                                                                    case "hello":
                                                                        v5 = append(v5, (fi15 / ((128) | 1)))
                                                                    }
                                                                    _ = v161
                                                                } else {
                                                                    fmt.Println(fi15, v4, v1)
                                                                    var v163 int8 = 1
                                                                    _ = v163
                                                                }
                                                            }
                                                            fmt.Println(v4, v4)
                                                            switch "a" {
                                                            case "foo bar":
                                                                var v164 float64 = (1e10 - -2.25)
                                                                _ = v164
                                                            case "":
                                                                v2[42] = 128
                                                            case "a":
                                                                v1 = 255
                                                            }
                                                            if ((!(!(43 == 42))) || false) {
                                                                var v165 uint16 = 1
                                                                switch -819 {
                                                                case 127:
                                                                    v2[0] = -9223372036854775808
                                                                case 2147483647:
                                                                    var v166 uint64 = v3
                                                                    _ = v166
                                                                }
                                                                _ = v165
                                                            }
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        } else {
                                                            v3 = v3
                                                            v167 := GenS15{F0: int16(v4), F1: 165, F2: (65535 == 702)}
                                                            _ = v167
                                                        }
                                                    }
                                                } else {
                                                    v168 := GenS21{F0: (v4 / ((v4) | 1)), F1: 1}
                                                    _ = v168
                                                }
                                            case "foo bar":
                                                var v169 float32 = (3.14159 + 1e10)
                                                _ = v169
                                            default:
                                                fmt.Println(v1, v57)
                                            }
                                        }
                                        if ((-2.25 - (float32(1e10) + (0.0 / 1.0))) <= ((0.0 + 3.14159) * 0.1)) {
                                            fmt.Println(v4, fi15)
                                        } else {
                                            if true {
                                                _ = len(v5)
                                                v4 = v4
                                            }
                                        }
                                        switch "a" {
                                        case "foo bar":
                                            if true {
                                                v170 := GenS18{F0: (^(v57)), F1: uint32(v4)}
                                                _ = v170
                                            } else {
                                                _ = len(v2)
                                                var v171 float64 = 3.14159
                                                var v172 int8 = -1
                                                _ = v171
                                                _ = v172
                                            }
                                            fmt.Println(fi22, v57)
                                            v3 += v3
                                        case "":
                                            fmt.Println(v3)
                                        case "hello":
                                            v173 := GenS22{F0: 693, F1: "a"}
                                            _ = v173
                                        default:
                                            fmt.Println(v4)
                                            var v174 uint8 = (v4 - 40)
                                            if (100 < 100) {
                                                var v175 int8 = 123
                                                _ = v175
                                            } else {
                                                v176 := GenS7{F0: -9223372036854775808, F1: 127, F2: -32768}
                                                _ = v176
                                            }
                                            v4 = v4
                                            _ = v174
                                        }
                                    }
                                    _ = v57
                                }
                            case "hello":
                                var v177 int64 = 42
                                _ = v177
                            }
                        }
                        var v178 uint32 = 2147483648
                        _ = v178
                    } else {
                        v179 := GenS23{F0: 156}
                        var v180 string = "\n\t"
                        _ = v179
                        _ = v180
                    }
                case "\n\t":
                    fmt.Println(v1)
                }
                _ = fi22
                _ = v23
            }
        } else {
            if len(v5) > 0 {
                v5[0] = 32767
            }
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    if false {
        v4 += (v4 / (((^(v4))) | 1))
        for fi181 := 0; fi181 < 15; fi181++ {
            if true {
                delete(v2, 42)
                v2[479] = fi181
                for fi182 := 0; fi182 < 1; fi182++ {
                    var v183 uint8 = v4
                    switch (fi182 << ((255) & 63)) {
                    case -2147483648:
                        v184 := GenS13{F0: 65535, F1: (v183 & (v4 * (^(v183)))), F2: 128}
                        _ = v184
                    }
                    _ = fi182
                    _ = v183
                }
            } else {
                var v185 uint8 = v4
                _ = v185
            }
            _ = fi181
        }
        v186 := GenS6{F0: "\n\t"}
        _ = v186
    }
    v2[2] = 214
    if (!true) {
        if true {
            v4 += v4
            v1 = v1
            var v187 float32 = (3.14159 - 0.1)
            _ = v187
        }
    }
    v2[genFn0(-103, v3)] = -567
    for fi188 := 0; fi188 < 7; fi188++ {
        for fi189 := 0; fi189 < 1; fi189++ {
            _ = len(v5)
            v190 := GenS12{F0: false, F1: 2}
            v4 += (^(v4))
            var v191 int8 = -12
            v192 := GenS24{F0: -286}
            if (("" + (("foo bar" + "") + "a")) < "") {
                var v193 float64 = ((1e10 / ((0.1 * float64(3.14159)) * 1.5)) * (-2.25 - 1e10))
                _ = v193
            }
            switch ("foo bar" + ("\n\t" + "a")) {
            case "\n\t":
                v1 = v1
            case "a":
                fmt.Println(v4)
            case "foo bar":
                v1 = 1
            }
            _ = fi189
            _ = v190
            _ = v191
            _ = v192
        }
        fmt.Println(v4)
        _ = fi188
    }
    var v194 int = -1
    if (3.14159 > 1.5) {
        fmt.Println(v194)
        if (false || (211 < 1)) {
            v3 = 4294967295
            v5 = append(v5, 662)
        }
        v3 -= 1
    } else {
        switch (v194 | 127) {
        case 459:
            for fi195 := 0; fi195 < 9; fi195++ {
                fmt.Println(v194, fi195, v3)
                _ = fi195
            }
        case 1:
            var v196 int16 = 1000
            _ = v196
        }
        for fi197 := 0; fi197 < 6; fi197++ {
            v3 = v3
            var v198 int8 = 42
            _ = fi197
            _ = v198
        }
    }
    var v199 int = -66
    for fi200 := 0; fi200 < 5; fi200++ {
        var v201 int32 = -2147483648
        delete(v2, 65535)
        v202 := GenS4{F0: -377, F1: "a"}
        _ = fi200
        _ = v201
        _ = v202
    }
    fmt.Println("v1:", v1)
    fmt.Println("v194:", v194)
    fmt.Println("v199:", v199)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v194
    _ = v199
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
