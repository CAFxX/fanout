package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int8; F2 int32 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 int8; F1 int; F2 uint8 }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int64; F1 int64 }
type GenS12 struct { F0 uint; F1 uint64 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int; F1 int32; F2 string }
type GenS17 struct { F0 uint; F1 string }
type GenS18 struct { F0 uint64 }
type GenS19 struct { F0 uint16; F1 uint; F2 int8 }
type GenS20 struct { F0 int32; F1 uint32; F2 string }
type GenS21 struct { F0 int }
type GenS22 struct { F0 int; F1 uint8 }
func genFn0(p0 uint16, p1 string) int {
    v5 := []int{-422}
    var v6 int16 = 32767
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1() uint64 {
    v7 := GenS1{F0: ((("foo bar" + "a") + "a") + ("hello" + "")), F1: 1}
    if true {
        var v8 string = ""
        v9 := []int{9223372036854775807, genFn0(0, (v8 + v8))}
        _ = v8
        _ = v9
    }
    v10 := []int{9223372036854775807, 615, genFn0(626, ("hello" + "foo bar"))}
    var v11 float64 = 3.14159
    _ = v10
    _ = v11
    _ = v7
    return 727
}

func main() {
    var v1 uint32 = 65535
    var v2 uint8 = 234
    var v3 int32 = 1
    v4 := []uint32{v1, 4294967295, v1}
    switch "\n\t" {
    case "a":
        if (-125 >= -14) {
            if (!true) {
                var v12 uint16 = 224
                fmt.Println(v2, v12)
                _ = v12
            } else {
                var v13 uint = 1
                var v14 int16 = 0
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 16; fi15++ {
                v1 -= v1
                var v16 uint64 = 986
                _ = fi15
                _ = v16
            }
            for fi17 := 0; fi17 < 1; fi17++ {
                fmt.Println(v2, v3)
                var v18 int64 = -1
                switch genFn0(1, "") {
                case 9223372036854775807:
                    v19 := GenS2{F0: 1}
                    _ = v19
                default:
                    if (!(-131 < (-(fi17)))) {
                        v1 = 1
                        var v20 int16 = 1
                        _ = v20
                    } else {
                        switch "" {
                        case "foo bar":
                            switch -2147483648 {
                            case -193:
                                for fi21 := 0; fi21 < 9; fi21++ {
                                    fmt.Println(v3, v1)
                                    _ = fi21
                                }
                            case -840:
                                fmt.Println(fi17, v1)
                            case 32767:
                                if false {
                                    var v22 uint64 = 4294967295
                                    v23 := GenS2{F0: (v1 << ((0) & 63))}
                                    if false {
                                        var v24 int8 = -1
                                        _ = v24
                                    } else {
                                        v22 = (v22 - 174)
                                        v1 = (v1 * (v1 << ((87) & 63)))
                                        v22 -= v22
                                    }
                                    _ = v22
                                    _ = v23
                                } else {
                                    if true {
                                        for fi25 := 0; fi25 < 17; fi25++ {
                                            if true {
                                                v26 := GenS1{F0: "hello", F1: v1}
                                                var v27 float64 = (3.14159 - 3.14159)
                                                switch ("a" + ("foo bar" + "")) {
                                                case "":
                                                    if ((fi17 | (fi25 / ((genFn0(0, ((("hello" + "\n\t") + "\n\t") + "a"))) | 1))) > (fi25 / (((fi25 ^ genFn0(334, ""))) | 1))) {
                                                        var v28 uint64 = genFn1()
                                                        switch "" {
                                                        case "hello":
                                                            v4 = append(v4, v1)
                                                        case "foo bar":
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v29 := GenS3{F0: 127}
                                                                _ = v29
                                                            case "\n\t":
                                                                v30 := GenS1{F0: ("a" + "\n\t"), F1: 2}
                                                                _ = v30
                                                            case "a":
                                                                if false {
                                                                    v31 := GenS2{F0: (v1 / (((^(v1))) | 1))}
                                                                    if (true || (float32(v27) != float32(0.0))) {
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            switch fi25 {
                                                                            case 42:
                                                                                if ((fi25 - (fi17 % (((fi17 ^ fi17)) | 1))) != genFn0(731, "foo bar")) {
                                                                                    var v32 uint32 = v1
                                                                                    _ = v32
                                                                                } else {
                                                                                    var v33 uint64 = v28
                                                                                    v18 += (v18 % ((0) | 1))
                                                                                    _ = v33
                                                                                }
                                                                            case 0:
                                                                                fmt.Println(fi25, v18, fi25)
                                                                            case 128:
                                                                                if (v27 != (v27 + ((0.1 - v27) * v27))) {
                                                                                    v34 := GenS4{F0: ((v2 / ((v2) | 1)) != (v2 + v2)), F1: -128, F2: (v3 % ((-1) | 1))}
                                                                                    v18 -= (v18 + -9223372036854775808)
                                                                                    _ = v34
                                                                                }
                                                                                if (!("" > ("foo bar" + ""))) {
                                                                                    fmt.Println(v28)
                                                                                    v35 := GenS5{F0: (v18 ^ v18), F1: 493}
                                                                                    switch fi25 {
                                                                                    case 0:
                                                                                        if false {
                                                                                            var v36 int8 = 35
                                                                                            _ = v36
                                                                                        }
                                                                                    default:
                                                                                        var v37 uint64 = v28
                                                                                        _ = v37
                                                                                    }
                                                                                    _ = v35
                                                                                } else {
                                                                                    var v38 int32 = 915
                                                                                    v4 = append(v4, v1)
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = v1
                                                                                    }
                                                                                    _ = v38
                                                                                }
                                                                            default:
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = uint32(v28)
                                                                                }
                                                                            }
                                                                        case "":
                                                                            v1 -= 686
                                                                            v4 = append(v4, v1)
                                                                            v3 += v3
                                                                        case "a":
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            v28 = 644
                                                                            switch ("foo bar" + "") {
                                                                            case "a":
                                                                                switch genFn0(255, "hello") {
                                                                                case -670:
                                                                                    switch ("foo bar" + (("hello" + ("a" + "hello")) + "")) {
                                                                                    case "hello":
                                                                                        v27 -= v27
                                                                                        v27 -= v27
                                                                                    }
                                                                                case 4294967295:
                                                                                    fmt.Println(v18)
                                                                                case -9223372036854775808:
                                                                                    fmt.Println(v3)
                                                                                }
                                                                            case "foo bar":
                                                                                var v39 int = fi25
                                                                                _ = v39
                                                                            default:
                                                                                v40 := GenS4{F0: (255 != uint(v2)), F1: -128, F2: int32(v18)}
                                                                                _ = v40
                                                                            }
                                                                            v27 = ((-2.25 + (v27 + 0.1)) / (v27 + 0.1))
                                                                            fmt.Println(v2, v3)
                                                                        default:
                                                                            switch fi25 {
                                                                            case 0:
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            case 32767:
                                                                                var v41 float32 = 1.5
                                                                                _ = v41
                                                                            case -209:
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            default:
                                                                                fmt.Println(v1, fi25, v27)
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            v18 = 42
                                                                            v42 := GenS6{F0: (32767 >= 42), F1: ""}
                                                                            var v43 int16 = 0
                                                                            _ = v42
                                                                            _ = v43
                                                                        } else {
                                                                            var v44 int16 = 1
                                                                            fmt.Println(v18)
                                                                            if false {
                                                                                fmt.Println(fi25)
                                                                                v3 += v3
                                                                            }
                                                                            _ = v44
                                                                        }
                                                                        if (646 > 1) {
                                                                            if true {
                                                                                v45 := GenS7{F0: 52, F1: (fi17 * fi25), F2: (v2 ^ uint8(v28))}
                                                                                _ = v45
                                                                            }
                                                                            v46 := GenS3{F0: 42}
                                                                            _ = v46
                                                                        }
                                                                        v47 := GenS6{F0: (50 == 100), F1: "\n\t"}
                                                                        var v48 int16 = int16(v2)
                                                                        _ = v47
                                                                        _ = v48
                                                                    }
                                                                    if len(v4) > 0 {
                                                                        v4[0] = (v1 << ((1) & 63))
                                                                    }
                                                                    _ = v31
                                                                }
                                                            }
                                                        default:
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        _ = v28
                                                    } else {
                                                        if false {
                                                            v27 -= ((-2.25 + v27) - (v27 + (v27 + v27)))
                                                        }
                                                    }
                                                    v49 := GenS8{F0: (genFn0(1, "hello") != fi17), F1: (fi25 ^ fi17)}
                                                    _ = v49
                                                default:
                                                    v1 += (v1 >> ((42) & 63))
                                                }
                                                _ = v26
                                                _ = v27
                                            }
                                            v50 := GenS8{F0: true, F1: genFn0(255, "\n\t")}
                                            var v51 uint32 = v1
                                            _ = fi25
                                            _ = v50
                                            _ = v51
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        fmt.Println(v18, v1)
                                    }
                                }
                            default:
                                fmt.Println(v18, v2, fi17)
                            }
                        case "\n\t":
                            for fi52 := 0; fi52 < 13; fi52++ {
                                var v53 int64 = (v18 % (((v18 + (v18 ^ 9223372036854775807))) | 1))
                                _ = fi52
                                _ = v53
                            }
                        }
                        for fi54 := 0; fi54 < 14; fi54++ {
                            var v55 int32 = (v3 * 600)
                            var v56 int16 = 0
                            _ = fi54
                            _ = v55
                            _ = v56
                        }
                    }
                }
                _ = fi17
                _ = v18
            }
        } else {
            fmt.Println(v3)
            _ = len(v4)
            switch 2147483647 {
            case 497:
                var v57 int16 = -678
                _ = v57
            default:
                var v58 float32 = (0.0 + (float32(1.5) + 0.1))
                _ = v58
            }
        }
    default:
        v59 := GenS3{F0: 1}
        _ = v59
    }
    switch 127 {
    case 9223372036854775807:
        var v60 bool = (false || (!true))
        _ = v60
    }
    if ((0 != 456) && (!(!true))) {
        for fi61 := 0; fi61 < 17; fi61++ {
            for fi62 := 0; fi62 < 1; fi62++ {
                var v63 float32 = 1.5
                v64 := GenS4{F0: true, F1: -54, F2: -2147483648}
                _ = fi62
                _ = v63
                _ = v64
            }
            switch "hello" {
            case "hello":
                _ = len(v4)
            }
            _ = fi61
        }
        _ = len(v4)
        fmt.Println(v3, v3, v3)
    } else {
        v65 := GenS2{F0: v1}
        _ = v65
    }
    fmt.Println(v1, v2, v3)
    v66 := GenS9{F0: 32767, F1: -673, F2: 879}
    for fi67 := 0; fi67 < 14; fi67++ {
        var v68 float32 = 0.0
        var v69 int64 = -192
        _ = fi67
        _ = v68
        _ = v69
    }
    if len(v4) > 0 {
        v4[0] = (^(v1))
    }
    v70 := GenS7{F0: -128, F1: 437, F2: 1}
    fmt.Println(v2, v3, v2)
    var v71 uint8 = (v2 - v2)
    for fi72 := 0; fi72 < 14; fi72++ {
        var v73 int8 = 74
        var v74 string = (("" + "") + "hello")
        v3 += v3
        _ = fi72
        _ = v73
        _ = v74
    }
    var v75 float32 = float32((0.1 * 1.5))
    v1 += v1
    var v76 float32 = v75
    v4 = append(v4, v1)
    switch ("" + ("" + "a")) {
    case "":
        if ((v71 | v2) >= (^(v71))) {
            v71 -= 2
            v77 := GenS10{F0: -819}
            var v78 string = "\n\t"
            _ = v77
            _ = v78
        } else {
            v3 -= v3
            if (42 > -510) {
                for fi79 := 0; fi79 < 14; fi79++ {
                    if len(v4) > 0 {
                        v4[0] = v1
                    }
                    if true {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    v80 := GenS2{F0: v1}
                    _ = fi79
                    _ = v80
                }
            } else {
                v81 := GenS11{F0: 42, F1: 0}
                _ = v81
            }
        }
    case "hello":
        v82 := GenS7{F0: 100, F1: 42, F2: v2}
        _ = v82
    case "foo bar":
        switch 9223372036854775807 {
        case 127:
            v75 = (((v76 / (v75 - v75)) / v75) - v76)
        case 925:
            v83 := GenS4{F0: (828 != -9223372036854775808), F1: -1, F2: 2147483647}
            _ = v83
        default:
            v76 += 0.1
        }
    default:
        if ((0 >= -1) && true) {
            var v84 bool = (127 <= 0)
            if ((v76 * ((v76 - v75) - 0.0)) != v76) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v85 int16 = int16(v71)
                _ = v85
            } else {
                fmt.Println(v71, v1, v76)
                for fi86 := 0; fi86 < 1; fi86++ {
                    v4 = append(v4, (v1 << ((4294967295) & 63)))
                    _ = fi86
                }
                fmt.Println(v3)
            }
            _ = v84
        }
    }
    for fi87 := 0; fi87 < 5; fi87++ {
        var v88 int8 = -54
        switch "a" {
        case "\n\t":
            if false {
                v2 += 127
            } else {
                switch ("" + ("" + "")) {
                case "hello":
                    switch (("" + ("\n\t" + "hello")) + "a") {
                    case "hello":
                        fmt.Println(v75, v2, v75)
                    case "foo bar":
                        for fi89 := 0; fi89 < 1; fi89++ {
                            v75 -= v75
                            v76 -= (((1.5 + (1.5 - v75)) + 1e10) / (1.5 / float32((3.14159 - 1.5))))
                            if true {
                                if false {
                                    if ("\n\t" < "") {
                                        if true {
                                            _ = len(v4)
                                            var v90 string = ("a" + "\n\t")
                                            _ = v90
                                        }
                                        v88 = (-(v88))
                                        var v91 int16 = 1000
                                        _ = v91
                                    }
                                } else {
                                    var v92 int32 = v3
                                    v93 := GenS12{F0: 65535, F1: 42}
                                    _ = v92
                                    _ = v93
                                }
                                var v94 int8 = v88
                                v95 := GenS13{F0: 1}
                                _ = v94
                                _ = v95
                            } else {
                                var v96 uint16 = 32767
                                v2 = (v71 + (^(v71)))
                                _ = v96
                            }
                            if ((^(v2)) != v71) {
                                if (fi89 > fi89) {
                                    if ((v71 | (^(v71))) >= (v71 | 255)) {
                                        v97 := GenS7{F0: (v88 ^ -55), F1: (fi89 / ((genFn0(1, "a")) | 1)), F2: 255}
                                        fmt.Println(v88, v76)
                                        if false {
                                            v98 := GenS14{F0: -470}
                                            fmt.Println(fi89)
                                            _ = v98
                                        }
                                        if (v1 < v1) {
                                            v2 = v71
                                        }
                                        v99 := GenS6{F0: true, F1: "\n\t"}
                                        var v100 uint = 65535
                                        _ = v100
                                        _ = v97
                                        _ = v99
                                    } else {
                                        var v101 float32 = 1e10
                                        v102 := GenS12{F0: 0, F1: uint64(v1)}
                                        _ = v101
                                        _ = v102
                                    }
                                } else {
                                    if true {
                                        var v103 int32 = (v3 % ((int32(fi87)) | 1))
                                        if true {
                                            v75 = v75
                                            v3 += v3
                                            if true {
                                                var v104 float64 = 1e10
                                                var v105 uint64 = 0
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v104
                                                _ = v105
                                            } else {
                                                var v106 string = ("hello" + "foo bar")
                                                var v107 float64 = float64(v76)
                                                if (4294967295 != genFn1()) {
                                                    v108 := GenS4{F0: true, F1: (v88 / ((v88) | 1)), F2: (v3 * -188)}
                                                    fmt.Println(v76, v71)
                                                    v109 := GenS12{F0: 0, F1: genFn1()}
                                                    v76 = (v76 / 3.14159)
                                                    _ = v108
                                                    _ = v109
                                                }
                                                v110 := GenS11{F0: -695, F1: -1}
                                                _ = v106
                                                _ = v107
                                                _ = v110
                                            }
                                            var v111 int32 = v3
                                            var v112 string = "hello"
                                            _ = v111
                                            _ = v112
                                        } else {
                                            v71 -= 42
                                            switch 4294967295 {
                                            case 32768:
                                                v113 := GenS2{F0: 42}
                                                _ = v113
                                            default:
                                                var v114 float64 = 0.0
                                                _ = v114
                                            }
                                            fmt.Println(v1, v76)
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v115 float64 = 1.5
                                            _ = v115
                                        case "a":
                                            v116 := GenS15{F0: v71}
                                            _ = v116
                                        default:
                                            if false {
                                                var v117 string = ("\n\t" + "hello")
                                                _ = v117
                                            } else {
                                                var v118 int8 = (v88 | v88)
                                                switch "" {
                                                case "foo bar":
                                                    v119 := GenS16{F0: (fi87 >> ((1) & 63)), F1: -327, F2: ("\n\t" + "\n\t")}
                                                    _ = v119
                                                case "a":
                                                    if len(v4) > 0 {
                                                        v4[0] = v1
                                                    }
                                                }
                                                v120 := GenS17{F0: 4294967295, F1: ("\n\t" + "foo bar")}
                                                _ = v118
                                                _ = v120
                                            }
                                        }
                                        _ = v103
                                    } else {
                                        if false {
                                            var v121 float32 = v76
                                            _ = v121
                                        } else {
                                            fmt.Println(v2, v1)
                                            var v122 int16 = 32767
                                            _ = v122
                                        }
                                        var v123 int32 = v3
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if true {
                                            v124 := GenS5{F0: int64(fi87), F1: -747}
                                            _ = v124
                                        } else {
                                            switch "hello" {
                                            case "\n\t":
                                                var v125 uint16 = 665
                                                _ = v125
                                            }
                                            _ = len(v4)
                                        }
                                        v126 := GenS17{F0: 292, F1: ((((("foo bar" + ("foo bar" + ("a" + ""))) + "a") + "") + "hello") + "\n\t")}
                                        if (((1.5 / (((1e10 * 1e10) / 3.14159) + 1.5)) * float64((v75 + v75))) > (0.1 + float64(v75))) {
                                            v127 := GenS4{F0: ((!true) && (!false)), F1: -84, F2: 905}
                                            _ = v127
                                        }
                                        _ = v123
                                        _ = v126
                                    }
                                    _ = len(v4)
                                }
                                v88 += (v88 % (((v88 / ((-105) | 1))) | 1))
                            } else {
                                if (((true || false) || true) && ((v71 * (^(v71))) >= (v2 << ((1) & 63)))) {
                                    var v128 float32 = 3.14159
                                    v1 += (v1 | (^(v1)))
                                    switch ("" + "hello") {
                                    case "a":
                                        if true {
                                            v129 := GenS18{F0: 18446744073709551615}
                                            var v130 string = "foo bar"
                                            v88 = (v88 & v88)
                                            _ = v129
                                            _ = v130
                                        }
                                    default:
                                        _ = len(v4)
                                    }
                                    _ = v128
                                } else {
                                    if ((v88 << ((0) & 63)) < v88) {
                                        fmt.Println(v76, v2, v1)
                                        if false {
                                            if len(v4) > 0 {
                                                v4[0] = 602
                                            }
                                            v2 = (v2 << ((4294967295) & 63))
                                        }
                                        if (false || false) {
                                            if true {
                                                if len(v4) > 0 {
                                                    v4[0] = (^(v1))
                                                }
                                                switch (fi87 | fi87) {
                                                case 32767:
                                                    fmt.Println(v76, v71)
                                                    var v131 int64 = -9223372036854775808
                                                    _ = v131
                                                case 927:
                                                    if (!((fi89 + (fi89 + fi89)) == genFn0(499, ("hello" + ("foo bar" + (((("a" + "\n\t") + "\n\t") + (("\n\t" + "foo bar") + ("" + "a"))) + "a")))))) {
                                                        if (v1 == (v1 % (((v1 % (((^(v1))) | 1))) | 1))) {
                                                            var v132 bool = false
                                                            var v133 int8 = (v88 / (((v88 >> ((uint(v1)) & 63))) | 1))
                                                            var v134 uint32 = 4294967295
                                                            _ = v132
                                                            _ = v133
                                                            _ = v134
                                                        } else {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            var v135 int64 = -182
                                                            fmt.Println(v135, v1, v135)
                                                            _ = v135
                                                        }
                                                        v136 := GenS18{F0: 989}
                                                        _ = v136
                                                    } else {
                                                        fmt.Println(v2)
                                                    }
                                                    var v137 float64 = 0.0
                                                    _ = v137
                                                }
                                            }
                                            v71 = v71
                                            var v138 uint8 = v2
                                            _ = v138
                                        }
                                    } else {
                                        fmt.Println(fi89, v3)
                                    }
                                }
                                v75 = v75
                                _ = len(v4)
                                var v139 string = (("a" + ("foo bar" + ("\n\t" + ("\n\t" + ("foo bar" + ("foo bar" + "")))))) + "foo bar")
                                v140 := GenS17{F0: 1, F1: ("foo bar" + "\n\t")}
                                _ = v139
                                _ = v140
                            }
                            _ = fi89
                        }
                    }
                case "a":
                    v76 = 0.0
                case "":
                    fmt.Println(v88)
                default:
                    switch "\n\t" {
                    case "hello":
                        for fi141 := 0; fi141 < 8; fi141++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if false {
                                if (v88 >= (v88 & v88)) {
                                    if true {
                                        v76 = float32(-2.25)
                                        var v142 int32 = v3
                                        fmt.Println(v1)
                                        if true {
                                            var v143 int64 = -96
                                            _ = v143
                                        }
                                        _ = len(v4)
                                        _ = v142
                                    } else {
                                        if ((fi87 / ((42) | 1)) <= 42) {
                                            switch fi87 {
                                            case -2147483648:
                                                var v144 uint32 = v1
                                                _ = v144
                                            case 9223372036854775807:
                                                fmt.Println(v1)
                                                v145 := GenS1{F0: "", F1: (v1 / ((v1) | 1))}
                                                _ = v145
                                            default:
                                                v88 = 127
                                            }
                                            var v146 int64 = 255
                                            _ = v146
                                        } else {
                                            if (1 < 32767) {
                                                if true {
                                                    fmt.Println(v76, v71, v76)
                                                    v4 = append(v4, v1)
                                                }
                                                var v147 int32 = v3
                                                if (2147483647 >= int32(fi87)) {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v148 := GenS19{F0: 0, F1: uint(v1), F2: (v88 & int8(v1))}
                                                    _ = v148
                                                }
                                                var v149 string = "\n\t"
                                                _ = v147
                                                _ = v149
                                            }
                                            v76 += v76
                                        }
                                        v75 += v75
                                        if ((827 > 442) || (float32((1e10 / 3.14159)) <= v75)) {
                                            var v150 uint = 1
                                            _ = v150
                                        }
                                    }
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    var v151 uint8 = v71
                                    _ = v151
                                case "hello":
                                    if (!false) {
                                        if true {
                                            if false {
                                                var v152 float32 = (v75 - ((float32(1.5) / 1.0) + (1.5 - v75)))
                                                if true {
                                                    switch "" {
                                                    case "hello":
                                                        if len(v4) > 0 {
                                                            v4[0] = 2147483647
                                                        }
                                                    }
                                                }
                                                switch (-(fi87)) {
                                                case 854:
                                                    v153 := GenS11{F0: -1, F1: 1}
                                                    _ = v153
                                                case 32767:
                                                    if (42 >= 1) {
                                                        var v154 int8 = (v88 / (((-(v88))) | 1))
                                                        fmt.Println(v152, fi87)
                                                        v155 := GenS14{F0: 0}
                                                        switch -856 {
                                                        case 2147483647:
                                                            if true {
                                                                switch (fi141 ^ (^(fi87))) {
                                                                case -496:
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                case -198:
                                                                    if (!true) {
                                                                        var v156 int64 = -1
                                                                        v2 += 42
                                                                        var v157 uint = 202
                                                                        _ = v156
                                                                        _ = v157
                                                                    } else {
                                                                        switch fi141 {
                                                                        case 0:
                                                                            _ = len(v4)
                                                                            v158 := GenS1{F0: "hello", F1: v1}
                                                                            _ = v158
                                                                        default:
                                                                            v159 := GenS20{F0: v3, F1: 1, F2: "a"}
                                                                            _ = v159
                                                                        }
                                                                        v4 = append(v4, (v1 | 2))
                                                                    }
                                                                case 128:
                                                                    var v160 float64 = 0.0
                                                                    _ = v160
                                                                }
                                                            } else {
                                                                v1 -= 2147483648
                                                                var v161 float32 = v152
                                                                var v162 uint8 = 255
                                                                _ = v161
                                                                _ = v162
                                                            }
                                                        case 4294967295:
                                                            if len(v4) > 0 {
                                                                v4[0] = (^(v1))
                                                            }
                                                        }
                                                        _ = v154
                                                        _ = v155
                                                    } else {
                                                        if len(v4) > 0 {
                                                            v4[0] = (v1 << ((42) & 63))
                                                        }
                                                    }
                                                default:
                                                    v163 := GenS21{F0: (fi87 + fi141)}
                                                    _ = v163
                                                }
                                                _ = v152
                                            } else {
                                                v71 = 228
                                                var v164 uint32 = v1
                                                _ = v164
                                            }
                                            fmt.Println(v2, v71)
                                            v76 += (v75 + ((3.14159 - 0.0) / -2.25))
                                        }
                                        switch 32768 {
                                        case 4294967295:
                                            if ((!true) && (true || true)) {
                                                var v165 uint16 = 42
                                                _ = v165
                                            } else {
                                                if ((0.1 + (1.5 / 1.0)) == (-2.25 / (0.0 + ((1.5 + 0.1) + 3.14159)))) {
                                                    var v166 int32 = (v3 ^ v3)
                                                    fmt.Println(v2, v76, v3)
                                                    _ = v166
                                                }
                                                if false {
                                                    fmt.Println(fi87)
                                                }
                                            }
                                        case 0:
                                            var v167 uint8 = 145
                                            _ = v167
                                        case -945:
                                            var v168 uint16 = 2
                                            _ = v168
                                        }
                                        v3 = v3
                                    }
                                case "a":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                var v169 float64 = 0.0
                                _ = v169
                            }
                            _ = fi141
                        }
                    case "\n\t":
                        v75 = (v75 / 0.1)
                    case "foo bar":
                        v76 = float32(0.0)
                    }
                }
            }
        case "foo bar":
            switch ("hello" + ("foo bar" + "hello")) {
            case "hello":
                fmt.Println(v2)
            }
        case "a":
            var v170 float32 = (0.0 + (float32(1.5) / v75))
            _ = v170
        default:
            if (1e10 > ((float64(v76) * 1.5) + 0.0)) {
                v3 = -1
            } else {
                if true {
                    for fi171 := 0; fi171 < 7; fi171++ {
                        v76 -= ((v76 + 1e10) - ((3.14159 * v76) - (-2.25 / v75)))
                        _ = fi171
                    }
                    switch "foo bar" {
                    case "":
                        v172 := GenS16{F0: 127, F1: 42, F2: "hello"}
                        _ = v172
                    case "a":
                        v76 += v76
                    }
                } else {
                    var v173 int = genFn0(45, "a")
                    fmt.Println(v3, v173, v1)
                    _ = v173
                }
                if len(v4) > 0 {
                    v4[0] = v1
                }
                switch ("hello" + "hello") {
                case "foo bar":
                    var v174 bool = (!true)
                    _ = v174
                case "\n\t":
                    v175 := GenS22{F0: (fi87 % (((fi87 % ((int(v3)) | 1))) | 1)), F1: v2}
                    _ = v175
                default:
                    var v176 uint32 = (v1 * (^(v1)))
                    _ = v176
                }
            }
        }
        var v177 uint = 186
        _ = fi87
        _ = v177
        _ = v88
    }
    if (2 <= 2) {
        var v178 int16 = 42
        fmt.Println(v2)
        v76 = v75
        _ = v178
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v71:", v71)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v66
    _ = v70
    _ = v71
    _ = v75
    _ = v76
}
