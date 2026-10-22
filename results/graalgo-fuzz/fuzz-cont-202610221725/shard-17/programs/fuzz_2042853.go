package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 int8; F1 int64; F2 string }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int32; F1 int16; F2 uint64 }
type GenS10 struct { F0 int64; F1 uint16; F2 bool }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint; F1 bool; F2 string }
type GenS13 struct { F0 uint8; F1 string; F2 int32 }
type GenS14 struct { F0 int8; F1 int8; F2 uint16 }
func genFn0(p0 int32) string {
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    if false {
        var v8 string = "a"
        var v9 int16 = -545
        switch ("\n\t" + ((v8 + "") + (("a" + v8) + v8))) {
        case "\n\t":
            v9 = -991
        case "hello":
            var v10 string = v8
            _ = v10
        case "foo bar":
            for fi11 := 0; fi11 < 10; fi11++ {
                p0 -= (p0 << ((4294967295) & 63))
                fmt.Println(v9, v8)
                var v12 uint = 42
                _ = fi11
                _ = v12
            }
        default:
            v9 += (v9 & v9)
        }
        _ = v8
        _ = v9
    } else {
        if (true || false) {
            fmt.Println(p0)
        } else {
            switch 9223372036854775807 {
            case 548:
                var v13 uint32 = 2
                _ = v13
            }
        }
        for fi14 := 0; fi14 < 7; fi14++ {
            fmt.Println(p0)
            fmt.Println(fi14)
            fmt.Println(p0, p0)
            _ = fi14
        }
        p0 -= (p0 << ((10) & 63))
    }
    p0 = p0
    p0 -= (p0 + -2147483648)
    _ = p0
    return "\n\t"
}
func genFn1(p0 bool) int {
    v15 := []int{-509}
    for fi16 := 0; fi16 < 14; fi16++ {
        v17 := GenS1{F0: 65535}
        for fi18 := 0; fi18 < 14; fi18++ {
            if p0 {
                fmt.Println(fi18)
                if p0 {
                    if true {
                        v19 := GenS2{F0: 42, F1: 32767}
                        _ = v19
                    } else {
                        p0 = (!p0)
                        fmt.Println(p0, fi16, fi18)
                    }
                    if (1.5 > -2.25) {
                        v15 = append(v15, fi18)
                    } else {
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                    }
                }
            } else {
                p0 = false
                var v20 uint = 2
                _ = v20
            }
            _ = fi18
        }
        if (4294967295 < 0) {
            if p0 {
                if (true && (!p0)) {
                    for fi21 := 0; fi21 < 13; fi21++ {
                        if p0 {
                            p0 = false
                            switch fi16 {
                            case -888:
                                if (537 == 708) {
                                    if p0 {
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                        v15 = append(v15, -2147483648)
                                    } else {
                                        v22 := GenS3{F0: -2147483648}
                                        var v23 float64 = -2.25
                                        if (1 <= 42) {
                                            v24 := GenS2{F0: 127, F1: 65535}
                                            var v25 int8 = 100
                                            _ = v24
                                            _ = v25
                                        }
                                        switch (fi16 - fi21) {
                                        case -999:
                                            switch "foo bar" {
                                            case "foo bar":
                                                if p0 {
                                                    fmt.Println(v23)
                                                    switch (("\n\t" + ("foo bar" + "\n\t")) + "") {
                                                    case "a":
                                                        p0 = (-81 == -9223372036854775808)
                                                    case "hello":
                                                        switch (-(fi21)) {
                                                        case -97:
                                                            fmt.Println(p0, fi16, v23)
                                                        case 32767:
                                                            if p0 {
                                                                if len(v15) >= 2 {
                                                                    v15 = v15[1:len(v15)]
                                                                }
                                                                if len(v15) > 0 {
                                                                    v15[0] = 0
                                                                }
                                                            }
                                                        }
                                                        _ = len(v15)
                                                        if (true && true) {
                                                            fmt.Println(p0, v23, fi21)
                                                            var v26 int32 = 1
                                                            if (0 < 0) {
                                                                if p0 {
                                                                    v23 -= float64(((1e10 * 0.0) * (1.5 - ((0.0 + 0.1) / (((1.5 * 1e10) - (1.5 - (0.0 + float32((-2.25 / v23))))) - float32((1e10 * 1.5)))))))
                                                                } else {
                                                                    v27 := GenS4{F0: (v26 | -2147483648), F1: 554}
                                                                    var v28 uint = 2
                                                                    _ = v27
                                                                    _ = v28
                                                                }
                                                            } else {
                                                                p0 = (!((p0 || ((255 == 22) && p0)) || p0))
                                                                var v29 float64 = 0.1
                                                                v15 = append(v15, fi16)
                                                                _ = v29
                                                            }
                                                            _ = v26
                                                        }
                                                    }
                                                    v30 := GenS4{F0: -1, F1: 9223372036854775807}
                                                    _ = v30
                                                } else {
                                                    if len(v15) > 0 {
                                                        v15[0] = fi16
                                                    }
                                                    var v31 string = ""
                                                    v32 := GenS2{F0: -15, F1: 65535}
                                                    p0 = p0
                                                    if (0 <= 1) {
                                                        v23 -= v23
                                                        fmt.Println(v31)
                                                    } else {
                                                        var v33 float32 = 1e10
                                                        var v34 float64 = -2.25
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                    _ = v31
                                                    _ = v32
                                                }
                                            case "hello":
                                                if (22 <= 42) {
                                                    if p0 {
                                                        var v35 uint = 487
                                                        switch fi16 {
                                                        case 32768:
                                                            v36 := GenS5{F0: 0, F1: 439}
                                                            _ = v36
                                                        }
                                                        var v37 uint32 = 65535
                                                        _ = v35
                                                        _ = v37
                                                    } else {
                                                        v38 := GenS3{F0: 2147483647}
                                                        switch (("hello" + (("hello" + "") + ("foo bar" + "\n\t"))) + "a") {
                                                        case "foo bar":
                                                            p0 = (!(false && (!p0)))
                                                        case "a":
                                                            v39 := GenS2{F0: 10, F1: 782}
                                                            _ = v39
                                                        }
                                                        _ = v38
                                                    }
                                                    p0 = false
                                                }
                                            case "\n\t":
                                                v40 := GenS6{F0: 21, F1: -531, F2: ""}
                                                _ = v40
                                            }
                                        case 160:
                                            p0 = (252 > 669)
                                        }
                                        v23 -= v23
                                        _ = v22
                                        _ = v23
                                    }
                                    p0 = (!p0)
                                    var v41 uint64 = 2
                                    _ = v41
                                }
                            }
                            switch fi16 {
                            case -1:
                                var v42 uint32 = 2
                                _ = v42
                            case 4294967295:
                                fmt.Println(fi21, fi21, fi21)
                            default:
                                if p0 {
                                    fmt.Println(p0, fi16, p0)
                                    if (fi16 >= 430) {
                                        if (!(p0 || p0)) {
                                            switch "foo bar" {
                                            case "\n\t":
                                                if ((-2.25 <= 0.1) && (255 > 42)) {
                                                    p0 = (false && p0)
                                                    if len(v15) > 0 {
                                                        v15[0] = (fi21 - (fi21 * fi21))
                                                    }
                                                    switch fi21 {
                                                    case 2147483647:
                                                        var v43 uint = 4294967295
                                                        _ = v43
                                                    }
                                                    var v44 int8 = -114
                                                    p0 = (-9223372036854775808 != 42)
                                                    v45 := GenS4{F0: 356, F1: 0}
                                                    var v46 int16 = int16(fi16)
                                                    fmt.Println(fi16)
                                                    fmt.Println(v46, v46, fi21)
                                                    _ = v44
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    p0 = p0
                                                    var v47 int64 = 9223372036854775807
                                                    switch ("a" + ("" + ((("foo bar" + "") + ("foo bar" + (("" + "\n\t") + ""))) + "hello"))) {
                                                    case "hello":
                                                        v48 := GenS1{F0: 2}
                                                        _ = v48
                                                    case "\n\t":
                                                        var v49 uint16 = 226
                                                        _ = v49
                                                    default:
                                                        if false {
                                                            fmt.Println(fi21, fi16, v47)
                                                            v47 += (v47 | (-(v47)))
                                                            var v50 int32 = 42
                                                            _ = v50
                                                        } else {
                                                            v51 := GenS6{F0: 1, F1: 0, F2: "\n\t"}
                                                            _ = v51
                                                        }
                                                    }
                                                    _ = v47
                                                }
                                            case "hello":
                                                var v52 string = "\n\t"
                                                _ = v52
                                            case "":
                                                p0 = p0
                                                var v53 int64 = -9223372036854775808
                                                if (4294967295 >= 724) {
                                                    var v54 uint8 = 255
                                                    if true {
                                                        fmt.Println(p0)
                                                        fmt.Println(v54, v54, fi16)
                                                        p0 = (1e10 == -2.25)
                                                    } else {
                                                        v54 = v54
                                                        v54 += 42
                                                        v53 += (v53 ^ v53)
                                                    }
                                                    _ = v54
                                                } else {
                                                    fmt.Println(v53)
                                                    var v55 int32 = -2147483648
                                                    _ = v55
                                                }
                                                v56 := GenS7{F0: 137, F1: (fi21 * 32767)}
                                                _ = v53
                                                _ = v56
                                            }
                                            v57 := GenS5{F0: 42, F1: 42}
                                            v58 := GenS7{F0: 42, F1: (fi21 + fi21)}
                                            _ = v57
                                            _ = v58
                                        } else {
                                            v59 := GenS8{F0: 65535}
                                            v60 := GenS9{F0: -435, F1: 333, F2: 2}
                                            _ = v59
                                            _ = v60
                                        }
                                        switch fi21 {
                                        case 0:
                                            if (0 != 1) {
                                                if len(v15) >= 2 {
                                                    v15 = v15[1:len(v15)]
                                                }
                                                if p0 {
                                                    if len(v15) >= 2 {
                                                        v15 = v15[1:len(v15)]
                                                    }
                                                    var v61 int = fi16
                                                    _ = v61
                                                } else {
                                                    var v62 int = (fi16 ^ -646)
                                                    _ = v62
                                                }
                                                p0 = (1 >= 26)
                                            }
                                        case 624:
                                            switch (-(fi21)) {
                                            case 0:
                                                var v63 int32 = 664
                                                _ = v63
                                            default:
                                                p0 = ((0 <= -2147483648) && (p0 && false))
                                            }
                                        case 32767:
                                            var v64 int = (fi21 & 726)
                                            _ = v64
                                        }
                                        var v65 float32 = (((1e10 + (float32(0.0) / 1.0)) / (3.14159 - -2.25)) / float32(3.14159))
                                        _ = v65
                                    }
                                    if (p0 || ((1e10 <= (1.5 / 1.5)) && false)) {
                                        switch fi16 {
                                        case -1:
                                            p0 = p0
                                        case 32768:
                                            switch (("a" + "foo bar") + "hello") {
                                            case "\n\t":
                                                var v66 string = ""
                                                _ = v66
                                            case "":
                                                if (666 == -1) {
                                                    fmt.Println(fi21, p0, fi16)
                                                    if len(v15) > 0 {
                                                        v15[0] = fi16
                                                    }
                                                }
                                            }
                                        default:
                                            var v67 int = fi21
                                            _ = v67
                                        }
                                        switch "hello" {
                                        case "\n\t":
                                            v68 := GenS1{F0: 65535}
                                            _ = v68
                                        case "a":
                                            switch (fi16 / (((fi21 >> ((664) & 63))) | 1)) {
                                            case -9223372036854775808:
                                                p0 = p0
                                            }
                                        case "foo bar":
                                            fmt.Println(fi21)
                                        default:
                                            if len(v15) > 0 {
                                                v15[0] = fi21
                                            }
                                        }
                                        var v69 int16 = -981
                                        _ = v69
                                    }
                                    if true {
                                        p0 = (!((!false) && ("" > "hello")))
                                        p0 = p0
                                    } else {
                                        var v70 bool = p0
                                        if (1 == 219) {
                                            p0 = (-2147483648 < 1)
                                        } else {
                                            if (((("a" + "\n\t") + "hello") + "foo bar") >= "hello") {
                                                p0 = p0
                                                if len(v15) > 0 {
                                                    v15[0] = -804
                                                }
                                                v70 = true
                                            } else {
                                                if len(v15) > 0 {
                                                    v15[0] = 4294967295
                                                }
                                                p0 = false
                                                if v70 {
                                                    var v71 int8 = 127
                                                    p0 = ((18446744073709551615 <= 42) && p0)
                                                    _ = v71
                                                }
                                            }
                                            var v72 int = (fi16 ^ fi21)
                                            var v73 int8 = 1
                                            var v74 int64 = 1
                                            _ = v72
                                            _ = v73
                                            _ = v74
                                        }
                                        _ = v70
                                    }
                                    _ = len(v15)
                                }
                            }
                        } else {
                            v15 = append(v15, (-(fi16)))
                            var v75 int = (fi16 | fi16)
                            _ = len(v15)
                            _ = v75
                        }
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        if (p0 && p0) {
                            p0 = (((3.14159 * ((float64(1e10) - 0.0) * (1.5 * (0.1 + -2.25)))) * ((((1e10 * (1e10 + 1.5)) - 1e10) - 0.0) * -2.25)) == 3.14159)
                            p0 = p0
                        }
                        var v76 float32 = 3.14159
                        p0 = p0
                        if (false && true) {
                            p0 = p0
                            p0 = p0
                            _ = len(v15)
                        }
                        var v77 string = "\n\t"
                        _ = fi21
                        _ = v76
                        _ = v77
                    }
                    v78 := GenS10{F0: 0, F1: 1, F2: p0}
                    _ = v78
                } else {
                    v79 := GenS2{F0: 0, F1: 0}
                    _ = v79
                }
                v15 = append(v15, fi16)
            }
            if true {
                var v80 uint32 = 65535
                _ = v80
            }
        }
        _ = fi16
        _ = v17
    }
    p0 = p0
    _ = p0
    _ = v15
    return -2147483648
}

func main() {
    var v1 float32 = (3.14159 - 1.5)
    v2 := make(map[string]int)
    var v3 uint32 = 529
    var v4 uint = 240
    v5 := []uint32{592}
    var v6 int = 32768
    var v81 uint32 = 1
    var v82 uint32 = 2147483648
    fmt.Println(v81, v82, v6)
    var v83 float32 = 1e10
    var v84 uint32 = v3
    var v85 uint8 = 67
    for fi86 := 0; fi86 < 12; fi86++ {
        v81 = (v84 >> (((^(v4))) & 63))
        if (v1 > v1) {
            var v87 int64 = 9223372036854775807
            _ = v87
        } else {
            for fi88 := 0; fi88 < 19; fi88++ {
                switch ((("hello" + (("foo bar" + "hello") + ((("" + "foo bar") + "\n\t") + "foo bar"))) + "foo bar") + "foo bar") {
                case "":
                    fmt.Println(v85, fi86)
                default:
                    v84 -= 42
                }
                switch (v6 ^ genFn1((!false))) {
                case 2147483647:
                    delete(v2, (("hello" + "") + ("" + "hello")))
                case 128:
                    switch ("foo bar" + "a") {
                    case "a":
                        var v89 float32 = v1
                        _ = v89
                    default:
                        switch genFn1((2147483647 != int32(v81))) {
                        case -1:
                            _ = len(v2)
                        }
                    }
                    if (-1 > 100) {
                        var v90 float32 = float32(((1e10 * (0.1 + 1.5)) + 3.14159))
                        switch ("" + "\n\t") {
                        case "a":
                            switch ("a" + ("foo bar" + "foo bar")) {
                            case "\n\t":
                                v91 := GenS1{F0: 886}
                                _ = v91
                            case "hello":
                                if (!(!true)) {
                                    switch "" {
                                    case "a":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    case "\n\t":
                                        var v92 uint32 = (v84 % ((2147483647) | 1))
                                        _ = v92
                                    case "hello":
                                        switch (fi88 >> ((v4) & 63)) {
                                        case 1:
                                            if (int64(v4) != -642) {
                                                v93 := GenS11{F0: "\n\t"}
                                                _ = v93
                                            }
                                        }
                                    default:
                                        fmt.Println(v83)
                                    }
                                }
                            case "":
                                v3 = (^(v3))
                            default:
                                fmt.Println(v3, fi86, v81)
                                var v94 float64 = 0.1
                                if (false || true) {
                                    var v95 uint8 = 79
                                    if true {
                                        if (3.14159 <= 1e10) {
                                            switch genFn1((!false)) {
                                            case -1:
                                                v96 := GenS8{F0: (v4 & v4)}
                                                _ = v96
                                            case 4294967295:
                                                v82 += 163
                                            case 1:
                                                v83 = v90
                                            }
                                        }
                                        v85 = (^(v95))
                                        var v97 uint = (^(v4))
                                        var v98 uint8 = v85
                                        var v99 uint = v4
                                        _ = v97
                                        _ = v98
                                        _ = v99
                                    } else {
                                        v6 += genFn1(false)
                                        switch fi86 {
                                        case 32768:
                                            v100 := GenS8{F0: v4}
                                            _ = v100
                                        case 128:
                                            delete(v2, "hello")
                                        }
                                    }
                                    if (v94 != v94) {
                                        _ = len(v2)
                                        var v101 uint = 18446744073709551615
                                        _ = v101
                                    } else {
                                        _ = len(v2)
                                        fmt.Println(v3)
                                        v102 := GenS12{F0: 2, F1: (689 >= -3), F2: ""}
                                        _ = v102
                                    }
                                    _ = v95
                                } else {
                                    switch "a" {
                                    case "hello":
                                        fmt.Println(v6, v82)
                                    default:
                                        if (593 != (^(v4))) {
                                            if (-192 > 742) {
                                                if (-96 != 1) {
                                                    fmt.Println(v94)
                                                    if (!false) {
                                                        switch "a" {
                                                        case "foo bar":
                                                            v103 := GenS13{F0: (v85 << (((v4 << ((v4) & 63))) & 63)), F1: "foo bar", F2: -2147483648}
                                                            _ = v103
                                                        case "a":
                                                            var v104 int32 = 435
                                                            _ = v104
                                                        default:
                                                            var v105 bool = (363 > 0)
                                                            _ = v105
                                                        }
                                                    }
                                                }
                                                var v106 float64 = 0.1
                                                _ = v106
                                            } else {
                                                fmt.Println(fi86, v85)
                                                var v107 int32 = -587
                                                v108 := GenS10{F0: 42, F1: 2, F2: (int16(v84) <= 0)}
                                                _ = v107
                                                _ = v108
                                            }
                                            var v109 int64 = -1
                                            var v110 uint = (v4 << ((1) & 63))
                                            _ = v109
                                            _ = v110
                                        } else {
                                            if false {
                                                v111 := GenS4{F0: -26, F1: 9223372036854775807}
                                                _ = v111
                                            }
                                            switch 32768 {
                                            case -934:
                                                v6 -= genFn1(true)
                                            case -411:
                                                v84 = (^(v3))
                                            case -946:
                                                var v112 uint64 = uint64(v81)
                                                _ = v112
                                            default:
                                                if len(v5) > 0 {
                                                    v5[0] = 1
                                                }
                                            }
                                            switch (fi88 % ((v6) | 1)) {
                                            case -1:
                                                delete(v2, "\n\t")
                                            case -914:
                                                var v113 int32 = 42
                                                _ = v113
                                            case 607:
                                                v114 := GenS14{F0: int8(v85), F1: -101, F2: 763}
                                                _ = v114
                                            default:
                                                v4 = (v4 & (^(v4)))
                                            }
                                        }
                                    }
                                }
                                _ = v94
                            }
                        case "foo bar":
                            var v115 int32 = -2147483648
                            _ = v115
                        case "":
                            fmt.Println(v81, v81, v3)
                            v81 += v82
                        default:
                            var v116 int16 = 395
                            _ = v116
                        }
                        _ = v90
                    } else {
                        v3 = 2147483647
                    }
                }
                _ = fi88
            }
        }
        v4 += uint(v3)
        _ = fi86
    }
    var v117 int16 = 1000
    for fi118 := 0; fi118 < 15; fi118++ {
        switch "hello" {
        case "":
            v6 = fi118
        case "\n\t":
            var v119 uint = v4
            _ = v119
        default:
            var v120 float32 = 3.14159
            _ = v120
        }
        var v121 uint64 = 42
        v2[("" + ("foo bar" + "\n\t"))] = (fi118 << ((4294967295) & 63))
        _ = fi118
        _ = v121
    }
    if true {
        v117 = (v117 | (v117 % (((v117 - -32768)) | 1)))
        if (!(true && ((v85 % ((57) | 1)) <= v85))) {
            if (0 >= 2147483647) {
                v122 := v2[("\n\t" + "foo bar")]
                var v123 uint16 = 1
                _ = v122
                _ = v123
            } else {
                fmt.Println(v81)
                v124 := v2["\n\t"]
                v124 -= (v6 - 127)
                _ = v124
            }
            var v125 uint = (v4 % (((v4 * 374)) | 1))
            v82 += v84
            _ = v125
        } else {
            if true {
                v82 = (v3 / (((^(v84))) | 1))
            }
        }
        fmt.Println(v84, v6)
    } else {
        v83 -= ((float32((1.5 + float64(float32(0.0)))) - (v83 / v83)) - -2.25)
        delete(v2, "hello")
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v82
    _ = v83
    _ = v84
    _ = v85
}
