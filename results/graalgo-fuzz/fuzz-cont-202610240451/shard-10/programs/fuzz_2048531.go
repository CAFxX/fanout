package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint16; F1 int32; F2 string }
type GenS9 struct { F0 int16; F1 uint32; F2 uint }
type GenS10 struct { F0 bool; F1 uint8 }
type GenS11 struct { F0 int16; F1 int32; F2 int }
type GenS12 struct { F0 int32; F1 int16; F2 uint8 }
type GenS13 struct { F0 string; F1 uint64; F2 uint8 }
type GenS14 struct { F0 int32 }
func genFn0() uint8 {
    var v6 float32 = 3.14159
    switch 2147483647 {
    case 128:
        v6 += -2.25
    case 2147483647:
        v7 := []int{42, 342, -484}
        _ = v7
    default:
        var v8 uint16 = 255
        _ = v8
    }
    if (!true) {
        if false {
            var v9 int32 = 1
            if false {
                fmt.Println(v9)
                v10 := []int{795}
                var v11 int32 = v9
                _ = v10
                _ = v11
            }
            _ = v9
        } else {
            fmt.Println(v6)
        }
        var v12 int = 4294967295
        _ = v12
    } else {
        v13 := []int{-619}
        switch "hello" {
        case "\n\t":
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 uint = 255
                v16 := GenS1{F0: 42}
                for fi17 := 0; fi17 < 2; fi17++ {
                    v15 += (v15 << ((v15) & 63))
                    var v18 uint16 = 32767
                    v18 = 414
                    var v19 int8 = 42
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        case "a":
            for fi20 := 0; fi20 < 1; fi20++ {
                v6 += v6
                _ = fi20
            }
        case "hello":
            var v21 int = 0
            _ = v21
        default:
            fmt.Println(v6)
        }
        _ = v13
    }
    v22 := GenS1{F0: 293}
    _ = v22
    _ = v6
    return 168
}

func main() {
    var v1 uint = 255
    var v2 int64 = 0
    var v3 uint64 = 18446744073709551615
    v4 := make(map[uint32]string)
    var v5 int64 = -65
    var v23 uint8 = 211
    for fi24 := 0; fi24 < 18; fi24++ {
        var v25 uint = (^(v1))
        var v26 uint64 = v3
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 uint = (v1 / ((18446744073709551615) | 1))
    var v28 uint8 = (v23 | v23)
    for fi29 := 0; fi29 < 7; fi29++ {
        if true {
            v30 := v4[65535]
            _ = v30
        } else {
            fmt.Println(v2)
        }
        fmt.Println(fi29, v27, v28)
        var v31 int8 = 53
        _ = fi29
        _ = v31
    }
    if true {
        v1 -= v1
        fmt.Println(v5, v28, v1)
        fmt.Println(v27)
    } else {
        fmt.Println(v28, v28, v5)
        fmt.Println(v1)
    }
    var v32 bool = true
    v23 = (v23 << ((v27) & 63))
    fmt.Println(v5, v2)
    for fi33 := 0; fi33 < 15; fi33++ {
        var v34 int16 = 159
        _ = fi33
        _ = v34
    }
    if true {
        if (false && (((((((1.5 / 1e10) - -2.25) != (1.5 - (1e10 - 1.5))) && (v32 || v32)) && ((!v32) || (!((v5 * (v2 & v5)) > v2)))) || ("foo bar" == "foo bar")) || true)) {
            if (!((v27 >> ((1) & 63)) <= v27)) {
                var v35 uint16 = 0
                if true {
                    var v36 float32 = -2.25
                    v32 = (((0 > 0) && (true || v32)) && v32)
                    v2 -= (v5 * 1)
                    _ = v36
                } else {
                    for fi37 := 0; fi37 < 8; fi37++ {
                        v38 := []int{(fi37 & (fi37 | fi37)), -2147483648, 127, -1}
                        _ = fi37
                        _ = v38
                    }
                }
                _ = v35
            } else {
                v39 := GenS2{F0: 2}
                var v40 int8 = -128
                switch 128 {
                case 32768:
                    var v41 int = 128
                    _ = v41
                case 32767:
                    v42 := []int{128, 42, 4294967295, 1}
                    _ = v42
                case -771:
                    for fi43 := 0; fi43 < 11; fi43++ {
                        v44 := GenS1{F0: 218}
                        var v45 uint64 = (v3 >> (((v1 % ((v27) | 1))) & 63))
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                }
                _ = v39
                _ = v40
            }
            v4[2] = "\n\t"
        }
        var v46 bool = false
        _ = v46
    } else {
        if (v27 != (v27 & v27)) {
            var v47 int64 = int64(v1)
            v48 := GenS1{F0: (v3 - v3)}
            v28 = 0
            _ = v47
            _ = v48
        }
    }
    if ((^(v27)) <= (v1 % ((v27) | 1))) {
        if (433 == 42) {
            v49 := GenS3{F0: (v23 / ((v28) | 1)), F1: 2147483647, F2: 1}
            fmt.Println(v28)
            _ = v49
        }
        if false {
            v50 := []int{-9223372036854775808, 854, 32768, 9223372036854775807}
            if (!((float32(0.0) >= -2.25) || (0 <= 2147483647))) {
                fmt.Println(v28, v2, v28)
            }
            v51 := GenS3{F0: genFn0(), F1: 846, F2: 1}
            _ = v50
            _ = v51
        } else {
            fmt.Println(v1, v5)
        }
    }
    if v32 {
        switch ("hello" + "") {
        case "\n\t":
            switch 4294967295 {
            case 32768:
                v3 = (v3 + v3)
            case -2147483648:
                v52 := GenS4{F0: 1}
                _ = v52
            case -680:
                delete(v4, 65535)
            default:
                if false {
                    if (!(!v32)) {
                        switch 2147483647 {
                        case 127:
                            v53 := GenS5{F0: v23, F1: (v23 / (((v23 ^ genFn0())) | 1))}
                            _ = v53
                        case 575:
                            for fi54 := 0; fi54 < 6; fi54++ {
                                v55 := GenS3{F0: (v23 & uint8(v1)), F1: -1, F2: 255}
                                var v56 uint8 = v23
                                _ = fi54
                                _ = v55
                                _ = v56
                            }
                        default:
                            if false {
                                switch 2147483647 {
                                case 1:
                                    if (v32 || v32) {
                                        v57 := GenS6{F0: v5}
                                        v58 := []int{42, 2147483647, 32767}
                                        _ = v57
                                        _ = v58
                                    } else {
                                        v59 := []int{-473}
                                        if ((v3 ^ 850) > v3) {
                                            v1 = (v27 | v1)
                                            for fi60 := 0; fi60 < 6; fi60++ {
                                                if v32 {
                                                    v61 := GenS2{F0: 42}
                                                    _ = v61
                                                } else {
                                                    if true {
                                                        var v62 uint = 0
                                                        fmt.Println(v62, v62)
                                                        v3 = 42
                                                        _ = v62
                                                    }
                                                }
                                                _ = fi60
                                            }
                                        }
                                        for fi63 := 0; fi63 < 11; fi63++ {
                                            if v32 {
                                                fmt.Println(v32)
                                                var v64 int = -705
                                                _ = v64
                                            } else {
                                                if true {
                                                    v5 = 251
                                                } else {
                                                    _ = len(v4)
                                                    switch "hello" {
                                                    case "a":
                                                        for fi65 := 0; fi65 < 20; fi65++ {
                                                            v66 := GenS2{F0: 0}
                                                            v67 := v4[628]
                                                            _ = fi65
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                    case "foo bar":
                                                        if v32 {
                                                            switch (^(fi63)) {
                                                            case 1:
                                                                var v68 int = fi63
                                                                _ = v68
                                                            default:
                                                                for fi69 := 0; fi69 < 6; fi69++ {
                                                                    if (true || true) {
                                                                        var v70 string = "hello"
                                                                        v71 := GenS6{F0: (v5 ^ 832)}
                                                                        v1 = 18446744073709551615
                                                                        _ = v70
                                                                        _ = v71
                                                                    }
                                                                    if v32 {
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            var v72 float32 = 1.5
                                                                            v73 := GenS7{F0: int64(v1)}
                                                                            _ = v72
                                                                            _ = v73
                                                                        case "foo bar":
                                                                            var v74 uint64 = 0
                                                                            _ = v74
                                                                        }
                                                                        fmt.Println(fi63, v23)
                                                                    } else {
                                                                        if (0 != (fi69 * fi63)) {
                                                                            v28 = (v23 / ((v28) | 1))
                                                                            v75 := GenS3{F0: genFn0(), F1: 745, F2: 2}
                                                                            v76 := GenS6{F0: (v5 % ((14) | 1))}
                                                                            v3 -= 4294967295
                                                                            _ = len(v4)
                                                                            fmt.Println(v32)
                                                                            _ = v75
                                                                            _ = v76
                                                                        }
                                                                    }
                                                                    fmt.Println(fi63)
                                                                    _ = fi69
                                                                }
                                                            }
                                                        } else {
                                                            var v77 bool = false
                                                            if (-2.25 <= 3.14159) {
                                                                fmt.Println(v32, v5)
                                                                v78 := GenS7{F0: int64(fi63)}
                                                                _ = v78
                                                            }
                                                            _ = v77
                                                        }
                                                    default:
                                                        var v79 uint32 = 4294967295
                                                        _ = v79
                                                    }
                                                }
                                                switch (("" + ("\n\t" + "\n\t")) + ("" + "foo bar")) {
                                                case "foo bar":
                                                    fmt.Println(v3, v3, v5)
                                                case "a":
                                                    var v80 int = fi63
                                                    _ = v80
                                                case "hello":
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        if true {
                                                            var v81 string = "foo bar"
                                                            _ = v81
                                                        } else {
                                                            if (((0.1 + (float32(float64((3.14159 * ((float32((0.0 * (0.0 + 1e10))) / ((1e10 - -2.25) + float32(3.14159))) + ((3.14159 + (1.5 - -2.25)) + ((-2.25 + float32((-2.25 * -2.25))) * (((float32(float64(0.1)) / 1.0) - 3.14159) / 1.0))))))) / ((((1.5 * -2.25) / (-2.25 / 0.1)) / 3.14159) - 1.5))) + 1e10) >= 1.5) {
                                                                fmt.Println(v32)
                                                                v2 = -9223372036854775808
                                                                if v32 {
                                                                    switch ("hello" + ("\n\t" + "hello")) {
                                                                    case "a":
                                                                        if (true || false) {
                                                                            if true {
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    for fi82 := 0; fi82 < 16; fi82++ {
                                                                                        var v83 uint16 = 0
                                                                                        _ = fi82
                                                                                        _ = v83
                                                                                    }
                                                                                case "\n\t":
                                                                                    fmt.Println(v2)
                                                                                case "a":
                                                                                    switch fi63 {
                                                                                    case 42:
                                                                                        var v84 int32 = -682
                                                                                        _ = v84
                                                                                    case 127:
                                                                                        var v85 int = (-(fi63))
                                                                                        _ = v85
                                                                                    case 128:
                                                                                        if (v32 && (v32 || (4294967295 <= 2147483648))) {
                                                                                            v5 += v5
                                                                                        } else {
                                                                                            var v86 uint = (v27 | 0)
                                                                                            _ = v86
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    switch (fi63 * 0) {
                                                                                    case -9223372036854775808:
                                                                                        v23 = uint8(v1)
                                                                                    case 847:
                                                                                        _ = len(v4)
                                                                                    }
                                                                                }
                                                                                var v87 bool = (42 > 863)
                                                                                _ = v87
                                                                            }
                                                                            for fi88 := 0; fi88 < 17; fi88++ {
                                                                                v89 := GenS3{F0: 68, F1: 1, F2: 55}
                                                                                _ = fi88
                                                                                _ = v89
                                                                            }
                                                                            v90 := GenS4{F0: int32(fi63)}
                                                                            _ = v90
                                                                        } else {
                                                                            var v91 uint8 = genFn0()
                                                                            v23 = 127
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                v92 := GenS4{F0: 641}
                                                                                _ = v92
                                                                            case "a":
                                                                                if v32 {
                                                                                    if v32 {
                                                                                        var v93 uint16 = 255
                                                                                        if len(v59) >= 2 {
                                                                                            v59 = v59[1:len(v59)]
                                                                                        }
                                                                                        _ = v93
                                                                                    } else {
                                                                                        var v94 int8 = 1
                                                                                        fmt.Println(v27, v27, v28)
                                                                                        _ = v94
                                                                                    }
                                                                                    if v32 {
                                                                                        if v32 {
                                                                                            v3 = v3
                                                                                        }
                                                                                        var v95 uint64 = (v3 + (^(v3)))
                                                                                        for fi96 := 0; fi96 < 2; fi96++ {
                                                                                            fmt.Println(fi63)
                                                                                            if (2 <= 46) {
                                                                                                v97 := GenS7{F0: v2}
                                                                                                if (!v32) {
                                                                                                    var v98 uint32 = 22
                                                                                                    _ = v98
                                                                                                } else {
                                                                                                    switch (fi63 * fi63) {
                                                                                                    case -287:
                                                                                                        if v32 {
                                                                                                            fmt.Println(v32)
                                                                                                        }
                                                                                                        if (980 == (v1 * (^(v1)))) {
                                                                                                            if v32 {
                                                                                                                v99 := GenS8{F0: uint16(v27), F1: 2147483647, F2: "foo bar"}
                                                                                                                _ = v99
                                                                                                            } else {
                                                                                                                var v100 uint32 = uint32(v3)
                                                                                                                if false {
                                                                                                                    if true {
                                                                                                                        if true {
                                                                                                                            if len(v59) >= 2 {
                                                                                                                                v59 = v59[1:len(v59)]
                                                                                                                            }
                                                                                                                            var v101 uint64 = v95
                                                                                                                            if (v27 <= v1) {
                                                                                                                                fmt.Println(v91, v100, v101)
                                                                                                                                v91 = (^(v23))
                                                                                                                            } else {
                                                                                                                                v1 = v27
                                                                                                                                v5 += (v5 + v5)
                                                                                                                            }
                                                                                                                            switch "" {
                                                                                                                            case "a":
                                                                                                                                var v102 int8 = 8
                                                                                                                                _ = v102
                                                                                                                            case "":
                                                                                                                                v101 = v95
                                                                                                                            case "\n\t":
                                                                                                                                fmt.Println(v32, v32)
                                                                                                                            default:
                                                                                                                                _ = len(v59)
                                                                                                                            }
                                                                                                                            if v32 {
                                                                                                                                v23 = genFn0()
                                                                                                                                var v103 int32 = -1
                                                                                                                                _ = v103
                                                                                                                            } else {
                                                                                                                                if v32 {
                                                                                                                                    if len(v59) >= 2 {
                                                                                                                                        v59 = v59[1:len(v59)]
                                                                                                                                    }
                                                                                                                                    if (true || (56 >= genFn0())) {
                                                                                                                                        v104 := GenS9{F0: 1, F1: (v100 - 2147483647), F2: (v27 / (((v27 + (v1 | v27))) | 1))}
                                                                                                                                        v4[v100] = "foo bar"
                                                                                                                                        var v105 string = ("\n\t" + ("foo bar" + ("\n\t" + (("\n\t" + ("hello" + ("\n\t" + "\n\t"))) + ""))))
                                                                                                                                        _ = v104
                                                                                                                                        _ = v105
                                                                                                                                    } else {
                                                                                                                                        v106 := GenS6{F0: 9223372036854775807}
                                                                                                                                        fmt.Println(v23)
                                                                                                                                        v32 = ((^(v100)) >= (v100 / ((v100) | 1)))
                                                                                                                                        _ = v106
                                                                                                                                    }
                                                                                                                                    var v107 uint = 4294967295
                                                                                                                                    _ = v107
                                                                                                                                }
                                                                                                                                var v108 bool = ((v5 * -907) >= (v2 & 42))
                                                                                                                                var v109 int16 = 1
                                                                                                                                _ = v108
                                                                                                                                _ = v109
                                                                                                                            }
                                                                                                                            fmt.Println(fi96)
                                                                                                                            _ = v101
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        var v110 int = 128
                                                                                                                        _ = v110
                                                                                                                    }
                                                                                                                    switch (^(fi96)) {
                                                                                                                    case -1:
                                                                                                                        v28 += (^(v28))
                                                                                                                        if true {
                                                                                                                            v91 += genFn0()
                                                                                                                            v4[(v100 | 65535)] = ""
                                                                                                                            if len(v59) > 0 {
                                                                                                                                v59[0] = fi63
                                                                                                                            }
                                                                                                                        }
                                                                                                                    case 1:
                                                                                                                        if len(v59) >= 2 {
                                                                                                                            v59 = v59[1:len(v59)]
                                                                                                                        }
                                                                                                                    case 32768:
                                                                                                                        v4[(^(v100))] = ("a" + "a")
                                                                                                                    }
                                                                                                                    v59 = append(v59, fi96)
                                                                                                                }
                                                                                                                var v111 uint64 = v3
                                                                                                                _ = v100
                                                                                                                _ = v111
                                                                                                            }
                                                                                                            var v112 bool = (!v32)
                                                                                                            if len(v59) >= 2 {
                                                                                                                v59 = v59[1:len(v59)]
                                                                                                            }
                                                                                                            _ = v112
                                                                                                        }
                                                                                                        v4[2147483648] = ("a" + "\n\t")
                                                                                                    case -266:
                                                                                                        fmt.Println(v2, v91)
                                                                                                    default:
                                                                                                        var v113 int8 = -67
                                                                                                        _ = v113
                                                                                                    }
                                                                                                    if ((!false) || v32) {
                                                                                                        var v114 int32 = -1
                                                                                                        _ = v114
                                                                                                    }
                                                                                                }
                                                                                                _ = v97
                                                                                            }
                                                                                            _ = fi96
                                                                                        }
                                                                                        _ = v95
                                                                                    }
                                                                                } else {
                                                                                    if v32 {
                                                                                        if len(v59) >= 2 {
                                                                                            v59 = v59[1:len(v59)]
                                                                                        }
                                                                                        var v115 int = fi63
                                                                                        _ = v115
                                                                                    }
                                                                                    v116 := GenS5{F0: uint8(v1), F1: 195}
                                                                                    v32 = v32
                                                                                    _ = v116
                                                                                }
                                                                            default:
                                                                                for fi117 := 0; fi117 < 19; fi117++ {
                                                                                    fmt.Println(fi63, v2, v2)
                                                                                    _ = fi117
                                                                                }
                                                                            }
                                                                            _ = v91
                                                                        }
                                                                    default:
                                                                        fmt.Println(v5, v28, v5)
                                                                    }
                                                                    switch "a" {
                                                                    case "hello":
                                                                        fmt.Println(v2, fi63, v1)
                                                                    case "a":
                                                                        _ = len(v4)
                                                                    case "foo bar":
                                                                        v32 = (0 == 1000)
                                                                    default:
                                                                        if true {
                                                                            for fi118 := 0; fi118 < 5; fi118++ {
                                                                                _ = len(v59)
                                                                                var v119 uint64 = (v3 % ((4294967295) | 1))
                                                                                _ = fi118
                                                                                _ = v119
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    v120 := GenS7{F0: int64(v3)}
                                                                    v121 := GenS9{F0: 0, F1: 0, F2: 1}
                                                                    if (!false) {
                                                                        if true {
                                                                            v2 = v2
                                                                        }
                                                                        var v122 bool = (-2.25 <= -2.25)
                                                                        _ = v122
                                                                    }
                                                                    _ = v120
                                                                    _ = v121
                                                                }
                                                            }
                                                        }
                                                    case "":
                                                        if (-2.25 > 3.14159) {
                                                            switch (fi63 & (fi63 % ((fi63) | 1))) {
                                                            case 42:
                                                                v2 = (v5 | 42)
                                                            case -768:
                                                                for fi123 := 0; fi123 < 10; fi123++ {
                                                                    v4[956] = (("a" + "a") + "a")
                                                                    var v124 bool = false
                                                                    switch fi63 {
                                                                    case 1:
                                                                        v125 := GenS10{F0: false, F1: (v28 / ((genFn0()) | 1))}
                                                                        _ = v125
                                                                    }
                                                                    _ = fi123
                                                                    _ = v124
                                                                }
                                                            default:
                                                                for fi126 := 0; fi126 < 6; fi126++ {
                                                                    fmt.Println(v5)
                                                                    fmt.Println(v28)
                                                                    _ = fi126
                                                                }
                                                            }
                                                        } else {
                                                            fmt.Println(v27, v5)
                                                        }
                                                    case "a":
                                                        v127 := GenS11{F0: 32767, F1: 1, F2: (-(fi63))}
                                                        _ = v127
                                                    }
                                                default:
                                                    _ = len(v59)
                                                }
                                                v32 = (v27 >= (v27 % ((1) | 1)))
                                            }
                                            _ = fi63
                                        }
                                        _ = v59
                                    }
                                case 32768:
                                    fmt.Println(v3, v23)
                                case -9223372036854775808:
                                    v2 -= -487
                                }
                                var v128 uint16 = 445
                                _ = v128
                            } else {
                                if false {
                                    v129 := GenS12{F0: -429, F1: -32768, F2: uint8(v1)}
                                    _ = v129
                                } else {
                                    v5 += -1
                                }
                                var v130 uint8 = genFn0()
                                _ = v130
                            }
                        }
                        if (!(1 != 1)) {
                            switch -137 {
                            case 9223372036854775807:
                                switch "foo bar" {
                                case "foo bar":
                                    for fi131 := 0; fi131 < 4; fi131++ {
                                        for fi132 := 0; fi132 < 17; fi132++ {
                                            if v32 {
                                                if (2 > 2147483648) {
                                                    v133 := GenS5{F0: (v28 - (^(v23))), F1: 255}
                                                    _ = v133
                                                }
                                                fmt.Println(v2, v28)
                                                if v32 {
                                                    var v134 int64 = v2
                                                    v134 = 0
                                                    fmt.Println(v3, v28)
                                                    _ = v134
                                                } else {
                                                    var v135 string = (("" + "\n\t") + "\n\t")
                                                    v136 := GenS13{F0: (v135 + v135), F1: v3, F2: v28}
                                                    if (false || ("foo bar" > "foo bar")) {
                                                        var v137 float32 = 1.5
                                                        switch fi132 {
                                                        case 851:
                                                            v23 = v23
                                                        case 0:
                                                            var v138 uint32 = 485
                                                            _ = v138
                                                        }
                                                        _ = v137
                                                    }
                                                    _ = v135
                                                    _ = v136
                                                }
                                            } else {
                                                v139 := GenS6{F0: v5}
                                                _ = v139
                                            }
                                            var v140 uint8 = uint8(v1)
                                            _ = fi132
                                            _ = v140
                                        }
                                        var v141 uint = 4294967295
                                        v32 = true
                                        _ = fi131
                                        _ = v141
                                    }
                                case "":
                                    var v142 string = ("foo bar" + ("a" + "\n\t"))
                                    _ = v142
                                }
                            case -637:
                                v5 -= (v5 % (((v5 | v5)) | 1))
                            }
                            if v32 {
                                v3 += v3
                            }
                            for fi143 := 0; fi143 < 9; fi143++ {
                                var v144 uint8 = uint8(v3)
                                var v145 float64 = 1e10
                                var v146 uint = (v1 ^ v27)
                                _ = fi143
                                _ = v144
                                _ = v145
                                _ = v146
                            }
                        } else {
                            if true {
                                v2 -= (v2 % ((v5) | 1))
                            } else {
                                var v147 int16 = 0
                                if (!(v32 && (("a" > "hello") || v32))) {
                                    v148 := GenS13{F0: "foo bar", F1: v3, F2: v28}
                                    _ = v148
                                } else {
                                    v149 := GenS8{F0: 42, F1: -829, F2: (("" + "hello") + "\n\t")}
                                    if v32 {
                                        var v150 float64 = (float64(3.14159) * (-2.25 + 1.5))
                                        _ = v150
                                    }
                                    _ = v149
                                }
                                for fi151 := 0; fi151 < 19; fi151++ {
                                    if v32 {
                                        v152 := GenS2{F0: 42}
                                        _ = v152
                                    }
                                    v4[2147483647] = "\n\t"
                                    v5 += (v5 & v5)
                                    _ = fi151
                                }
                                _ = v147
                            }
                            fmt.Println(v3, v3)
                        }
                    }
                    var v153 int16 = 182
                    _ = v153
                }
            }
        case "hello":
            v3 -= (v3 * (v3 & (v3 - 1)))
        }
    } else {
        v27 = (v27 ^ v1)
        var v154 int32 = -1
        _ = v154
    }
    if false {
        switch -2147483648 {
        case -2147483648:
            if (0.0 < 3.14159) {
                switch ("hello" + "") {
                case "a":
                    fmt.Println(v27, v3)
                case "\n\t":
                    v155 := GenS3{F0: 128, F1: 1, F2: 0}
                    _ = v155
                }
            } else {
                v156 := []int{128, 2147483647, 9223372036854775807, int(v3)}
                delete(v4, 4294967295)
                _ = v156
            }
        case 579:
            v157 := GenS2{F0: 1}
            _ = v157
        }
        switch 0 {
        case 127:
            v32 = v32
        }
    } else {
        v23 -= (v23 << (((v1 << ((v1) & 63))) & 63))
        switch 127 {
        case 32767:
            v27 -= v27
        case 9223372036854775807:
            if (v32 || true) {
                _ = len(v4)
                switch 128 {
                case -2147483648:
                    v1 -= uint(v23)
                case 739:
                    var v158 float64 = 0.0
                    _ = v158
                default:
                    v159 := GenS4{F0: 554}
                    _ = v159
                }
                v160 := GenS14{F0: -59}
                _ = v160
            } else {
                if (3.14159 != 1e10) {
                    var v161 int = int(v28)
                    _ = v161
                }
                for fi162 := 0; fi162 < 15; fi162++ {
                    switch ("" + "hello") {
                    case "hello":
                        switch "a" {
                        case "\n\t":
                            v163 := GenS5{F0: genFn0(), F1: genFn0()}
                            _ = v163
                        case "a":
                            v23 = v28
                        case "hello":
                            var v164 uint64 = v3
                            _ = v164
                        }
                    default:
                        for fi165 := 0; fi165 < 19; fi165++ {
                            var v166 int8 = -1
                            v167 := []int{358, (fi162 - -9223372036854775808), fi165, 2147483647}
                            v1 += (v1 * 255)
                            _ = fi165
                            _ = v166
                            _ = v167
                        }
                    }
                    _ = fi162
                }
                v168 := GenS8{F0: 32767, F1: -2147483648, F2: ("\n\t" + "a")}
                _ = v168
            }
        default:
            for fi169 := 0; fi169 < 20; fi169++ {
                var v170 uint64 = (v3 | (^(v3)))
                _ = fi169
                _ = v170
            }
        }
    }
    var v171 uint16 = 186
    fmt.Println("v1:", v1)
    fmt.Println("v171:", v171)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v171
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
