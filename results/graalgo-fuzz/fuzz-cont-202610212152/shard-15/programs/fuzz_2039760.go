package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 uint16; F2 bool }
type GenS5 struct { F0 int16; F1 int16; F2 uint }
type GenS6 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS7 struct { F0 int16; F1 int32; F2 uint16 }
type GenS8 struct { F0 int32; F1 int64 }
type GenS9 struct { F0 uint; F1 int32 }
type GenS10 struct { F0 uint8; F1 int64; F2 int }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint; F1 int32 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int; F1 int16 }
type GenS15 struct { F0 uint32; F1 int8; F2 int8 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 int16; F1 uint32; F2 int }
type GenS18 struct { F0 int32; F1 int64; F2 uint32 }
func genFn0() int16 {
    var v5 uint = 1
    switch -928 {
    case -1:
        v5 = 18446744073709551615
    default:
        v5 += (v5 / ((v5) | 1))
    }
    v6 := GenS1{F0: 255, F1: 42, F2: 0}
    v7 := GenS1{F0: 2, F1: 851, F2: (^(v5))}
    v5 = (v5 * v5)
    _ = v5
    _ = v6
    _ = v7
    return 1000
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{2147483648}
    v3 := make(map[string]string)
    v4 := []string{"\n\t", "hello"}
    var v8 uint = 4294967295
    if false {
        var v9 uint64 = 2
        _ = len(v4)
        switch -9223372036854775808 {
        case 42:
            if true {
                v9 = (v9 / ((18446744073709551615) | 1))
                fmt.Println(v8, v9)
                for fi10 := 0; fi10 < 18; fi10++ {
                    var v11 uint64 = (v9 & 0)
                    _ = fi10
                    _ = v11
                }
            } else {
                for fi12 := 0; fi12 < 18; fi12++ {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        v8 -= 18446744073709551615
                        v4 = append(v4, ("\n\t" + ((("foo bar" + "foo bar") + ("\n\t" + (("foo bar" + "a") + "\n\t"))) + "hello")))
                        _ = fi13
                    }
                    _ = fi12
                }
            }
        case 324:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 32767:
            var v14 float32 = (((0.0 - (((0.1 + (1.5 + float32(0.0))) / -2.25) * 3.14159)) * float32(1e10)) * ((1.5 - 1.5) + 3.14159))
            _ = v14
        default:
            for fi15 := 0; fi15 < 14; fi15++ {
                switch fi15 {
                case 42:
                    v8 = 2
                case -980:
                    for fi16 := 0; fi16 < 6; fi16++ {
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        _ = fi16
                    }
                }
                v17 := GenS1{F0: 32767, F1: v9, F2: (v8 - (v8 << ((v8) & 63)))}
                v18 := GenS1{F0: 478, F1: 2, F2: v8}
                _ = fi15
                _ = v17
                _ = v18
            }
        }
        _ = v9
    }
    var v19 uint8 = 42
    _ = len(v2)
    var v20 int16 = -757
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 int8 = -128
        _ = fi21
        _ = v22
    }
    v20 = (v20 % ((genFn0()) | 1))
    v20 -= v20
    if (32767 >= 302) {
        var v23 float32 = -2.25
        _ = v23
    } else {
        v8 = v8
    }
    switch -469 {
    case 128:
        switch "hello" {
        case "hello":
            switch ("hello" + ((("a" + "") + "foo bar") + ("foo bar" + "a"))) {
            case "hello":
                for fi24 := 0; fi24 < 6; fi24++ {
                    for fi25 := 0; fi25 < 20; fi25++ {
                        v26 := GenS1{F0: uint16(v8), F1: 4294967295, F2: v8}
                        var v27 float32 = float32((-2.25 - ((0.1 - float64(-2.25)) + ((3.14159 / float64(1.5)) - 0.1))))
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    _ = fi24
                }
            default:
                v8 -= v8
            }
        case "":
            switch "hello" {
            case "a":
                v19 += v19
            }
        case "\n\t":
            switch "hello" {
            case "hello":
                v19 = 0
            case "a":
                if ((true || false) || true) {
                    v28 := GenS1{F0: 37, F1: 42, F2: v8}
                    var v29 int = 1
                    v30 := GenS1{F0: 1, F1: 0, F2: v8}
                    _ = v28
                    _ = v29
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 12; fi31++ {
                        v32 := GenS1{F0: 633, F1: 469, F2: 46}
                        for fi33 := 0; fi33 < 13; fi33++ {
                            v8 -= v8
                            v8 = v8
                            _ = fi33
                        }
                        _ = fi31
                        _ = v32
                    }
                }
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
    }
    v19 -= (v19 << ((42) & 63))
    v4 = append(v4, "a")
    v4 = append(v4, "")
    if (v19 != 128) {
        v8 = v8
    } else {
        v19 += (v19 * v19)
    }
    if (((1e10 * ((1.5 + 1e10) * 1e10)) * 1.5) <= 0.1) {
        switch "\n\t" {
        case "\n\t":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        if true {
            for fi34 := 0; fi34 < 12; fi34++ {
                v35 := GenS2{F0: 2147483647, F1: 2, F2: v19}
                for fi36 := 0; fi36 < 13; fi36++ {
                    if false {
                        v37 := GenS1{F0: 1, F1: 1, F2: (v8 | v8)}
                        _ = v37
                    } else {
                        switch fi36 {
                        case 42:
                            var v38 bool = (((!(((!(!(true || (!(42 == uint64(v8)))))) && false) && (!false))) || (!(true && (100 == 100)))) || (true && (42 != 0)))
                            v19 += (v19 << (((v8 ^ 1)) & 63))
                            if len(v2) > 0 {
                                v2[0] = 2147483647
                            }
                            _ = v38
                        case 127:
                            var v39 int64 = 804
                            _ = v39
                        }
                    }
                    var v40 int = (^(fi34))
                    if (v8 < 297) {
                        var v41 uint = v8
                        v40 += v40
                        var v42 int = v40
                        v43 := GenS3{F0: "foo bar"}
                        if ("a" <= ("hello" + ("" + ("hello" + "a")))) {
                            v44 := GenS4{F0: int8(v20), F1: 0, F2: true}
                            var v45 float64 = float64((0.1 + 0.0))
                            var v46 int8 = 111
                            var v47 string = (("hello" + "") + "hello")
                            _ = v44
                            _ = v45
                            _ = v46
                            _ = v47
                        } else {
                            fmt.Println(v41, fi36, v42)
                            v19 = 255
                        }
                        _ = v41
                        _ = v42
                        _ = v43
                    } else {
                        var v48 string = "\n\t"
                        _ = v48
                    }
                    if ((!false) && (-420 < -2147483648)) {
                        var v49 uint32 = 0
                        _ = len(v2)
                        fmt.Println(fi36, v49, v19)
                        switch (-(fi36)) {
                        case 127:
                            if ((^(v19)) != v19) {
                                var v50 int64 = -585
                                _ = v50
                            }
                        default:
                            fmt.Println(v40)
                        }
                        _ = len(v3)
                        v20 = v20
                        _ = v49
                    }
                    _ = fi36
                    _ = v40
                }
                v19 += 0
                _ = fi34
                _ = v35
            }
            var v51 uint8 = (^(v19))
            _ = v51
        } else {
            fmt.Println(v8, v8, v8)
        }
        if (((!(!false)) || false) || false) {
            v8 = v8
            var v52 int32 = 1
            v53 := GenS4{F0: 0, F1: uint16(v8), F2: (18446744073709551615 != 0)}
            _ = v52
            _ = v53
        } else {
            if false {
                for fi54 := 0; fi54 < 14; fi54++ {
                    v55 := v3[("a" + "")]
                    _ = fi54
                    _ = v55
                }
                if ((((v8 * v8) == v8) || false) && (2147483647 > 2147483647)) {
                    delete(v1, ("" + "\n\t"))
                } else {
                    switch "hello" {
                    case "foo bar":
                        for fi56 := 0; fi56 < 18; fi56++ {
                            for fi57 := 0; fi57 < 6; fi57++ {
                                if false {
                                    _ = len(v4)
                                    v58 := GenS5{F0: v20, F1: -32768, F2: v8}
                                    v1["a"] = 606
                                    _ = v58
                                }
                                _ = fi57
                            }
                            _ = fi56
                        }
                    default:
                        if true {
                            v8 = (v8 / ((2) | 1))
                            if (int8(v19) < -128) {
                                v59 := GenS4{F0: -1, F1: 0, F2: (42 >= 18446744073709551615)}
                                fmt.Println(v20)
                                v60 := GenS6{F0: ((v19 ^ v19) < (^(v19))), F1: 278, F2: 127}
                                _ = v59
                                _ = v60
                            }
                        } else {
                            var v61 string = "foo bar"
                            v62 := GenS3{F0: (v61 + "a")}
                            if (!(false && true)) {
                                for fi63 := 0; fi63 < 18; fi63++ {
                                    switch -2147483648 {
                                    case 939:
                                        v64 := GenS6{F0: false, F1: 42, F2: (^(v19))}
                                        _ = v64
                                    case 4294967295:
                                        for fi65 := 0; fi65 < 3; fi65++ {
                                            if (true && false) {
                                                var v66 uint64 = 42
                                                v19 = 2
                                                v67 := GenS2{F0: 4294967295, F1: (^(v66)), F2: 2}
                                                _ = v66
                                                _ = v67
                                            } else {
                                                v8 -= 42
                                                fmt.Println(v20)
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v68 := GenS7{F0: 799, F1: 42, F2: 0}
                                                _ = len(v3)
                                                var v69 int64 = 42
                                                _ = v68
                                                _ = v69
                                            }
                                            _ = fi65
                                        }
                                    case 934:
                                        for fi70 := 0; fi70 < 8; fi70++ {
                                            if true {
                                                if false {
                                                    switch v61 {
                                                    case "hello":
                                                        v61 = ((v61 + (((((v61 + (("hello" + "foo bar") + (v61 + v61))) + v61) + "\n\t") + (v61 + "a")) + (v61 + v61))) + ("" + "a"))
                                                    }
                                                }
                                            } else {
                                                if true {
                                                    var v71 uint = v8
                                                    v72 := GenS8{F0: 2147483647, F1: 1}
                                                    _ = v71
                                                    _ = v72
                                                }
                                            }
                                            v19 = (v19 ^ v19)
                                            v20 = v20
                                            _ = fi70
                                        }
                                    default:
                                        fmt.Println(v20, fi63)
                                    }
                                    switch (fi63 / ((9223372036854775807) | 1)) {
                                    case 310:
                                        if (4294967295 < 1) {
                                            switch ("\n\t" + (("a" + (v61 + v61)) + ("" + ""))) {
                                            case "foo bar":
                                                if (!(!true)) {
                                                    var v73 int = -1
                                                    v1[(v61 + (v61 + (v61 + ("\n\t" + v61))))] = 384
                                                    _ = v73
                                                } else {
                                                    var v74 string = "hello"
                                                    var v75 string = ("foo bar" + (v61 + "\n\t"))
                                                    _ = v74
                                                    _ = v75
                                                }
                                            case "\n\t":
                                                var v76 uint32 = 2147483647
                                                _ = v76
                                            }
                                            v77 := GenS9{F0: (v8 / (((^(v8))) | 1)), F1: 1}
                                            if true {
                                                for fi78 := 0; fi78 < 19; fi78++ {
                                                    v79 := GenS10{F0: 127, F1: 42, F2: (-(fi78))}
                                                    v20 = v20
                                                    switch "hello" {
                                                    case "":
                                                        if (fi63 >= 48) {
                                                            if true {
                                                                if ((908 > 1) && (1 >= v8)) {
                                                                    switch v61 {
                                                                    case "\n\t":
                                                                        if false {
                                                                            _ = len(v1)
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                switch v61 {
                                                                                case "hello":
                                                                                    v80 := GenS11{F0: -55}
                                                                                    _ = v80
                                                                                case "a":
                                                                                    v61 = v61
                                                                                case "foo bar":
                                                                                    v8 = (^(v8))
                                                                                }
                                                                            }
                                                                            v19 -= (^(v19))
                                                                        }
                                                                    case "":
                                                                        v19 = (v19 >> ((731) & 63))
                                                                    case "foo bar":
                                                                        var v81 int64 = -9223372036854775808
                                                                        _ = v81
                                                                    }
                                                                }
                                                                v19 = 151
                                                                switch ((("a" + v61) + v61) + "") {
                                                                case "a":
                                                                    _ = len(v4)
                                                                case "\n\t":
                                                                    switch ((v61 + "foo bar") + v61) {
                                                                    case "a":
                                                                        var v82 uint = v8
                                                                        _ = v82
                                                                    }
                                                                case "foo bar":
                                                                    if false {
                                                                        if ((437 != 0) || true) {
                                                                            v61 = v61
                                                                        } else {
                                                                            var v83 string = v61
                                                                            _ = v83
                                                                        }
                                                                    }
                                                                default:
                                                                    if false {
                                                                        v19 = v19
                                                                        var v84 uint = (v8 | (v8 / ((v8) | 1)))
                                                                        _ = v84
                                                                    } else {
                                                                        _ = len(v2)
                                                                    }
                                                                }
                                                                v8 -= uint(v19)
                                                                fmt.Println(v20, v19, v61)
                                                            }
                                                            var v85 int64 = 386
                                                            var v86 float64 = 1.5
                                                            v8 = 33
                                                            if ((401 >= 493) && (!true)) {
                                                                v87 := GenS12{F0: v8, F1: -334}
                                                                _ = v87
                                                            } else {
                                                                var v88 int64 = (v85 / ((v85) | 1))
                                                                var v89 uint32 = 0
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                            v90 := GenS13{F0: 938}
                                                            var v91 bool = (((1 <= 421) || false) || ((v8 << ((v8) & 63)) > v8))
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            _ = v85
                                                            _ = v86
                                                            _ = v90
                                                            _ = v91
                                                        } else {
                                                            v92 := GenS7{F0: (v20 ^ v20), F1: -1, F2: 1}
                                                            _ = v92
                                                        }
                                                    }
                                                    switch "a" {
                                                    case "foo bar":
                                                        if false {
                                                            if true {
                                                                v93 := GenS14{F0: int(v20), F1: (v20 & (v20 - -211))}
                                                                v94 := GenS11{F0: 32767}
                                                                v8 = v8
                                                                _ = v93
                                                                _ = v94
                                                            } else {
                                                                v61 = (v61 + v61)
                                                                switch -2147483648 {
                                                                case 706:
                                                                    var v95 int64 = 258
                                                                    _ = v95
                                                                case 127:
                                                                    if false {
                                                                        v8 -= (v8 % ((v8) | 1))
                                                                        if (!(false || true)) {
                                                                            var v96 float32 = (0.0 * 1e10)
                                                                            if (false || (true || (!(v96 < v96)))) {
                                                                                v3[v61] = (("\n\t" + (v61 + "\n\t")) + "hello")
                                                                                if false {
                                                                                    v8 = 18446744073709551615
                                                                                    var v97 uint64 = 1
                                                                                    _ = v97
                                                                                }
                                                                            } else {
                                                                                var v98 string = v61
                                                                                v20 = 0
                                                                                _ = v98
                                                                            }
                                                                            var v99 int = 318
                                                                            if (true || ((!(((1e10 / 1.5) * v96) > 3.14159)) || false)) {
                                                                                v8 = 721
                                                                                v61 = "foo bar"
                                                                            } else {
                                                                                _ = len(v1)
                                                                            }
                                                                            _ = v96
                                                                            _ = v99
                                                                        } else {
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                if (42 == 18446744073709551615) {
                                                                                    v4 = append(v4, ((v61 + v61) + v61))
                                                                                    v20 = genFn0()
                                                                                    var v100 uint = (v8 ^ (v8 & (v8 - v8)))
                                                                                    v19 -= (v19 % ((2) | 1))
                                                                                    v8 -= (v100 / ((v100) | 1))
                                                                                    if (fi78 <= (fi63 ^ fi63)) {
                                                                                        var v101 int = fi63
                                                                                        var v102 uint32 = 333
                                                                                        _ = v101
                                                                                        _ = v102
                                                                                    } else {
                                                                                        v3[("" + ((("" + v61) + v61) + v61))] = v61
                                                                                    }
                                                                                    _ = v100
                                                                                } else {
                                                                                    if false {
                                                                                        v8 = (v8 / (((^(v8))) | 1))
                                                                                        v19 += v19
                                                                                        v19 += v19
                                                                                    } else {
                                                                                        v103 := GenS15{F0: 206, F1: 100, F2: 1}
                                                                                        _ = v103
                                                                                    }
                                                                                    v19 = 128
                                                                                    v104 := GenS6{F0: false, F1: 1, F2: v19}
                                                                                    _ = v104
                                                                                }
                                                                            case "a":
                                                                                v19 = v19
                                                                            case "":
                                                                                v61 = v61
                                                                            }
                                                                            if (65535 < 42) {
                                                                                if true {
                                                                                    v8 = v8
                                                                                    v105 := GenS2{F0: 42, F1: 18446744073709551615, F2: v19}
                                                                                    fmt.Println(v61, v61)
                                                                                    _ = v105
                                                                                }
                                                                            } else {
                                                                                v19 = (v19 + (v19 / ((196) | 1)))
                                                                                if (!(-2.25 != 0.0)) {
                                                                                    var v106 int8 = 50
                                                                                    fmt.Println(fi78, fi63)
                                                                                    if ((v20 ^ 32767) == (v20 << ((v8) & 63))) {
                                                                                        var v107 int64 = -9223372036854775808
                                                                                        v108 := v1[(v61 + v61)]
                                                                                        _ = v107
                                                                                        _ = v108
                                                                                    } else {
                                                                                        switch (v61 + (("hello" + "\n\t") + ("" + "\n\t"))) {
                                                                                        case "hello":
                                                                                            v8 = (v8 / ((42) | 1))
                                                                                        }
                                                                                        v8 += 42
                                                                                        if true {
                                                                                            if ((0.0 + 0.0) <= (1.5 / 1.0)) {
                                                                                                v1[(((v61 + (((v61 + "") + ("a" + v61)) + v61)) + (("\n\t" + v61) + v61)) + "hello")] = 545
                                                                                                if (((v19 | (v19 >> ((uint(v19)) & 63))) < (^(v19))) || false) {
                                                                                                    if (615 < 33) {
                                                                                                        v8 -= v8
                                                                                                        var v109 int = (fi78 * -9223372036854775808)
                                                                                                        _ = v109
                                                                                                    } else {
                                                                                                        fmt.Println(fi78, v61, v8)
                                                                                                        if ((v20 | (v20 + (^(v20)))) <= 32767) {
                                                                                                            fmt.Println(fi63)
                                                                                                            if (42 == (v8 / ((0) | 1))) {
                                                                                                                v110 := GenS16{F0: 771}
                                                                                                                fmt.Println(fi63)
                                                                                                                _ = v110
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v20)
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(v106)
                                                                                                v20 += -586
                                                                                            }
                                                                                            fmt.Println(v20, v8, fi78)
                                                                                            var v111 uint64 = 0
                                                                                            _ = v111
                                                                                        }
                                                                                        if true {
                                                                                            if (32767 == uint16(v8)) {
                                                                                                if len(v4) > 0 {
                                                                                                    v4[0] = "a"
                                                                                                }
                                                                                                if (("foo bar" > v61) && (!false)) {
                                                                                                    v112 := GenS16{F0: 18446744073709551615}
                                                                                                    if false {
                                                                                                        v113 := GenS7{F0: (v20 * 32767), F1: 2147483647, F2: 689}
                                                                                                        var v114 int32 = 164
                                                                                                        if false {
                                                                                                            v114 = v114
                                                                                                            switch -1 {
                                                                                                            case 693:
                                                                                                                v20 = -22
                                                                                                            }
                                                                                                            fmt.Println(v19, v106)
                                                                                                        }
                                                                                                        if len(v4) > 0 {
                                                                                                            v4[0] = (((v61 + ("" + v61)) + v61) + v61)
                                                                                                        }
                                                                                                        _ = v113
                                                                                                        _ = v114
                                                                                                    }
                                                                                                    _ = v112
                                                                                                }
                                                                                                var v115 uint = v8
                                                                                                _ = v115
                                                                                            }
                                                                                            switch fi63 {
                                                                                            case -1:
                                                                                                var v116 int8 = -1
                                                                                                _ = v116
                                                                                            }
                                                                                            var v117 string = "\n\t"
                                                                                            var v118 uint64 = 42
                                                                                            fmt.Println(v19, v20)
                                                                                            _ = len(v3)
                                                                                            _ = v117
                                                                                            _ = v118
                                                                                        }
                                                                                        v61 = v61
                                                                                        if true {
                                                                                            v20 += (v20 / ((genFn0()) | 1))
                                                                                            var v119 uint8 = (v19 ^ 1)
                                                                                            _ = v119
                                                                                        }
                                                                                    }
                                                                                    _ = v106
                                                                                }
                                                                            }
                                                                            var v120 int = 128
                                                                            var v121 uint16 = uint16(v8)
                                                                            switch v61 {
                                                                            case "hello":
                                                                                switch (v61 + (((v61 + "hello") + v61) + "")) {
                                                                                case "":
                                                                                    if (((float32((0.0 * (0.1 - (3.14159 - ((-2.25 + 1.5) - 1.5))))) * (0.0 / 1e10)) * 3.14159) == float32(((0.1 * (((float64(float32(1.5)) - 0.1) * 3.14159) + 3.14159)) * 0.1))) {
                                                                                        var v122 uint32 = 967
                                                                                        v123 := GenS14{F0: 42, F1: (v20 / ((v20) | 1))}
                                                                                        _ = v122
                                                                                        _ = v123
                                                                                    }
                                                                                case "\n\t":
                                                                                    var v124 float64 = -2.25
                                                                                    _ = v124
                                                                                default:
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                }
                                                                            case "a":
                                                                                fmt.Println(v19, v8, v61)
                                                                            case "foo bar":
                                                                                var v125 uint32 = 1
                                                                                _ = v125
                                                                            }
                                                                            fmt.Println(v61)
                                                                            _ = v120
                                                                            _ = v121
                                                                        }
                                                                        var v126 int8 = -1
                                                                        if ((v8 + 292) != 255) {
                                                                            if (false || ((true && false) && ((v19 & 42) > (v19 / ((128) | 1))))) {
                                                                                var v127 uint8 = (v19 * 255)
                                                                                var v128 uint16 = 1
                                                                                _ = v127
                                                                                _ = v128
                                                                            } else {
                                                                                switch (fi78 | (fi63 - fi63)) {
                                                                                case 0:
                                                                                    v8 = uint(v19)
                                                                                default:
                                                                                    var v129 uint64 = 1
                                                                                    _ = v129
                                                                                }
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                var v130 uint = (v8 + (v8 | v8))
                                                                                fmt.Println(fi63)
                                                                                _ = v130
                                                                            }
                                                                            v8 = 1
                                                                            if true {
                                                                                var v131 uint32 = 1
                                                                                var v132 uint32 = 2
                                                                                _ = v131
                                                                                _ = v132
                                                                            }
                                                                        }
                                                                        v133 := GenS17{F0: v20, F1: 4294967295, F2: (fi63 << ((v8) & 63))}
                                                                        _ = v126
                                                                        _ = v133
                                                                    } else {
                                                                        var v134 float64 = (0.1 + -2.25)
                                                                        delete(v1, ("a" + v61))
                                                                        v135 := GenS16{F0: 0}
                                                                        _ = v134
                                                                        _ = v135
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                    _ = fi78
                                                    _ = v79
                                                }
                                                var v136 int16 = (v20 / ((v20) | 1))
                                                _ = v136
                                            }
                                            _ = v77
                                        } else {
                                            if ((3.14159 + 1e10) != 3.14159) {
                                                v61 = v61
                                            }
                                            var v137 bool = (!(-856 >= 42))
                                            for fi138 := 0; fi138 < 12; fi138++ {
                                                var v139 float32 = 1e10
                                                var v140 int16 = 1000
                                                _ = fi138
                                                _ = v139
                                                _ = v140
                                            }
                                            _ = v137
                                        }
                                    case -9223372036854775808:
                                        v141 := GenS10{F0: (v19 % ((127) | 1)), F1: -9223372036854775808, F2: fi63}
                                        _ = v141
                                    default:
                                        for fi142 := 0; fi142 < 8; fi142++ {
                                            v143 := GenS18{F0: 1, F1: -1, F2: 4294967295}
                                            var v144 float32 = (3.14159 - (3.14159 + -2.25))
                                            v8 = (v8 ^ v8)
                                            _ = fi142
                                            _ = v143
                                            _ = v144
                                        }
                                    }
                                    _ = fi63
                                }
                                var v145 float64 = 3.14159
                                if (!(((!false) && false) || (!(!((v19 << (((v8 / ((4294967295) | 1))) & 63)) >= 2))))) {
                                    for fi146 := 0; fi146 < 20; fi146++ {
                                        if false {
                                            switch v61 {
                                            case "foo bar":
                                                v147 := GenS6{F0: ((fi146 | fi146) <= (fi146 >> ((v8) & 63))), F1: 32767, F2: 1}
                                                _ = v147
                                            case "hello":
                                                fmt.Println(v8)
                                            case "":
                                                fmt.Println(v61)
                                            }
                                            var v148 string = (v61 + "")
                                            _ = v148
                                        } else {
                                            for fi149 := 0; fi149 < 2; fi149++ {
                                                v20 -= (v20 & (-(v20)))
                                                fmt.Println(v61, v8)
                                                switch ((v61 + "hello") + v61) {
                                                case "a":
                                                    _ = len(v1)
                                                default:
                                                    var v150 float32 = (1.5 - (((0.0 / 0.1) - -2.25) * (0.1 / 1.5)))
                                                    _ = v150
                                                }
                                                _ = fi149
                                            }
                                            v151 := v1[v61]
                                            _ = v151
                                        }
                                        v8 = (v8 ^ (v8 + v8))
                                        _ = fi146
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = 1
                                    }
                                    v1[(v61 + (v61 + v61))] = 260
                                } else {
                                    var v152 int8 = 1
                                    var v153 bool = ((!true) && true)
                                    _ = v152
                                    _ = v153
                                }
                                _ = v145
                            }
                            _ = v61
                            _ = v62
                        }
                    }
                }
                if len(v4) > 0 {
                    v4[0] = ("foo bar" + "")
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
}
