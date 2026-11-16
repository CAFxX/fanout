package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint8; F1 int8; F2 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int; F1 string; F2 int64 }
type GenS8 struct { F0 uint64; F1 int16; F2 int32 }
type GenS9 struct { F0 uint8; F1 int; F2 uint8 }
type GenS10 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS11 struct { F0 int32; F1 uint64 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint64 }
func genFn0(p0 int16, p1 bool) int16 {
    p0 = p0
    var v7 int8 = 101
    _ = p0
    _ = p1
    _ = v7
    return -32768
}
func genFn1(p0 int32, p1 int) uint8 {
    for fi8 := 0; fi8 < 10; fi8++ {
        fmt.Println(p1, p0, p1)
        var v9 int16 = int16(p1)
        var v10 uint = 18446744073709551615
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        fmt.Println(p1)
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS1{F0: 255, F1: (p0 % ((p0) | 1))}
            var v14 int32 = 0
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    switch (^(p1)) {
    case 127:
        p0 = p0
    case 4294967295:
        switch "hello" {
        case "\n\t":
            v15 := GenS2{F0: 255}
            _ = v15
        case "foo bar":
            if true {
                fmt.Println(p0, p1)
                var v16 int8 = 1
                var v17 uint = 4294967295
                _ = v16
                _ = v17
            }
        default:
            switch (((("" + "a") + "\n\t") + (("a" + "") + ("foo bar" + "foo bar"))) + "\n\t") {
            case "a":
                p0 -= 0
            default:
                v18 := []int{(p1 + 128), -222, p1, (p1 >> ((18446744073709551615) & 63))}
                _ = v18
            }
        }
    case 2147483647:
        p0 = 334
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{570, 2}
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    v5 := []string{}
    v6 := []string{"a", "a", "\n\t"}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((false || (-384 < 4294967295)) || true) {
        var v19 uint32 = 0
        if true {
            v19 = v19
            v19 += 4294967295
            if (true || false) {
                var v20 string = ("hello" + ("foo bar" + ("\n\t" + "\n\t")))
                _ = v20
            } else {
                var v21 int64 = -9223372036854775808
                _ = v21
            }
        }
        if false {
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 int32 = 101
                if false {
                    if false {
                        fmt.Println(fi22, v19)
                        var v24 int32 = 1
                        _ = v24
                    }
                    v23 += -312
                    for fi25 := 0; fi25 < 20; fi25++ {
                        if len(v5) > 0 {
                            v5[0] = "foo bar"
                        }
                        _ = fi25
                    }
                }
                for fi26 := 0; fi26 < 9; fi26++ {
                    var v27 bool = (false && false)
                    var v28 uint16 = 351
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = fi22
                _ = v23
            }
        }
        _ = v19
    }
    if (-9223372036854775808 >= 42) {
        var v29 float64 = ((1.5 - 0.1) / 1.5)
        v30 := GenS1{F0: 1, F1: -1}
        if (100 > -117) {
            v29 = ((3.14159 + v29) - ((v29 * float64(0.0)) / float64(float32(v29))))
            var v31 uint32 = 2147483648
            _ = v31
        } else {
            v32 := GenS2{F0: 1}
            _ = v32
        }
        _ = v29
        _ = v30
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v33 uint64 = 29
    v3[-1] = ((921 <= uint16(v33)) && false)
    if true {
        var v34 int16 = -32768
        _ = v34
    } else {
        for fi35 := 0; fi35 < 11; fi35++ {
            switch ((((("a" + "\n\t") + "") + ("" + ("" + ("a" + ("" + ""))))) + "a") + ("a" + "foo bar")) {
            case "a":
                if (!(!(-2.25 > float32(0.0)))) {
                    v1["a"] = 918
                    switch fi35 {
                    case -506:
                        for fi36 := 0; fi36 < 6; fi36++ {
                            if false {
                                var v37 int = (fi35 | (fi35 + (fi35 + (fi36 % ((4294967295) | 1)))))
                                fmt.Println(fi36)
                                _ = v37
                            }
                            v38 := GenS3{F0: (v33 >> ((408) & 63)), F1: 1}
                            _ = fi36
                            _ = v38
                        }
                    case 9223372036854775807:
                        v33 = (v33 - 4294967295)
                    case -9223372036854775808:
                        if len(v6) > 0 {
                            v6[0] = "\n\t"
                        }
                    default:
                        fmt.Println(v33, v33)
                    }
                } else {
                    v39 := GenS4{F0: 0, F1: 100, F2: 354}
                    _ = v39
                }
            default:
                var v40 uint8 = 127
                _ = v40
            }
            _ = len(v3)
            _ = fi35
        }
    }
    for fi41 := 0; fi41 < 5; fi41++ {
        if true {
            if false {
                v33 = (v33 + 35)
                switch 32768 {
                case -2147483648:
                    if false {
                        v33 = (v33 % ((v33) | 1))
                    } else {
                        for fi42 := 0; fi42 < 7; fi42++ {
                            var v43 float32 = 0.1
                            if true {
                                if false {
                                    v33 = v33
                                    if false {
                                        v44 := GenS2{F0: 240}
                                        v45 := GenS1{F0: genFn1(0, fi41), F1: 2147483647}
                                        v2 = append(v2, 2147483648)
                                        _ = v44
                                        _ = v45
                                    } else {
                                        var v46 int8 = 116
                                        v47 := GenS5{F0: 1}
                                        delete(v1, "a")
                                        _ = v46
                                        _ = v47
                                    }
                                    v43 -= v43
                                }
                                var v48 bool = false
                                _ = len(v1)
                                _ = v48
                            } else {
                                if (((3.14159 - (((1e10 / 0.1) * 1e10) + 1e10)) + -2.25) > (-2.25 - (-2.25 + 1.5))) {
                                    var v49 int64 = 42
                                    var v50 int = (fi42 % ((-9223372036854775808) | 1))
                                    switch fi42 {
                                    case 128:
                                        var v51 uint8 = genFn1(42, fi41)
                                        _ = v51
                                    default:
                                        v52 := GenS6{F0: fi42}
                                        _ = v52
                                    }
                                    v33 = (v33 << ((1) & 63))
                                    var v53 float32 = -2.25
                                    if (false && false) {
                                        delete(v1, "")
                                        v43 = -2.25
                                    } else {
                                        v54 := GenS5{F0: -1}
                                        if false {
                                            switch ("hello" + "hello") {
                                            case "\n\t":
                                                v49 = v49
                                            default:
                                                if true {
                                                    var v55 uint = 2
                                                    if (v55 == 42) {
                                                        v56 := GenS7{F0: 42, F1: "a", F2: 42}
                                                        switch "a" {
                                                        case "\n\t":
                                                            _ = len(v3)
                                                        case "a":
                                                            switch (("a" + "a") + "") {
                                                            case "":
                                                                v53 -= v53
                                                            default:
                                                                fmt.Println(v55, fi41)
                                                            }
                                                        }
                                                        v5 = append(v5, "hello")
                                                        v50 += (fi41 << ((v55) & 63))
                                                        delete(v4, 2147483647)
                                                        _ = v56
                                                    }
                                                    if false {
                                                        if true {
                                                            if (!(0 < 255)) {
                                                                v33 -= 136
                                                                var v57 float64 = 0.1
                                                                _ = v57
                                                            }
                                                            v58 := GenS7{F0: fi41, F1: ("a" + "\n\t"), F2: (v49 >> (((v55 - 206)) & 63))}
                                                            _ = v58
                                                        }
                                                        v43 = v53
                                                    } else {
                                                        v53 = (1e10 / (1.5 / v43))
                                                        if ((0.1 >= 3.14159) && (v43 < v43)) {
                                                            var v59 uint32 = 4294967295
                                                            var v60 int64 = (^(v49))
                                                            _ = len(v6)
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                        if (412 >= 42) {
                                                            v55 = 647
                                                            fmt.Println(v50)
                                                            v53 = float32(float64((v53 * (3.14159 * 1e10))))
                                                            fmt.Println(v53)
                                                            v55 = v55
                                                        } else {
                                                            v43 += float32(0.0)
                                                            var v61 string = "a"
                                                            v62 := GenS8{F0: 1, F1: 650, F2: 42}
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                    }
                                                    _ = v55
                                                }
                                            }
                                            if true {
                                                v63 := GenS5{F0: 569}
                                                v33 = (v33 - 966)
                                                v43 = 0.1
                                                _ = v63
                                            } else {
                                                if (1e10 == float64(v53)) {
                                                    v64 := GenS9{F0: 221, F1: fi42, F2: 255}
                                                    v33 += 423
                                                    fmt.Println(v33, fi42)
                                                    v53 -= (0.0 + (float32(-2.25) - v53))
                                                    v65 := GenS5{F0: 42}
                                                    _ = v64
                                                    _ = v65
                                                }
                                                switch ("hello" + (("\n\t" + (("foo bar" + (((((("a" + "\n\t") + ("a" + "foo bar")) + "\n\t") + (("a" + "foo bar") + "foo bar")) + "\n\t") + ("hello" + ""))) + "a")) + ("a" + "foo bar"))) {
                                                case "hello":
                                                    switch (fi41 - (v50 % (((fi42 - (fi42 & fi42))) | 1))) {
                                                    case 966:
                                                        v66 := GenS10{F0: v33, F1: (!(-876 < -1)), F2: 903}
                                                        _ = v66
                                                    case 9223372036854775807:
                                                        if true {
                                                            fmt.Println(v33, v50)
                                                            switch v50 {
                                                            case 9223372036854775807:
                                                                if (v49 > 0) {
                                                                    if false {
                                                                        _ = len(v5)
                                                                        delete(v4, 0)
                                                                    } else {
                                                                        var v67 bool = false
                                                                        _ = v67
                                                                    }
                                                                    if true {
                                                                        switch (-(fi42)) {
                                                                        case -181:
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                var v68 int16 = 42
                                                                                fmt.Println(v49, fi41, fi41)
                                                                                var v69 string = ("" + "a")
                                                                                _ = v68
                                                                                _ = v69
                                                                            default:
                                                                                switch (("\n\t" + "a") + ("hello" + "hello")) {
                                                                                case "":
                                                                                    var v70 int = v50
                                                                                    _ = v70
                                                                                case "a":
                                                                                    var v71 uint = 750
                                                                                    _ = v71
                                                                                case "\n\t":
                                                                                    v3[(v50 >> ((1) & 63))] = true
                                                                                default:
                                                                                    var v72 int8 = 76
                                                                                    _ = v72
                                                                                }
                                                                            }
                                                                        case -1:
                                                                            v53 -= (((v43 + (float32(0.1) + v43)) - 0.0) + 1e10)
                                                                        }
                                                                        if false {
                                                                            var v73 uint8 = 2
                                                                            v74 := v4[v50]
                                                                            fmt.Println(fi41)
                                                                            _ = v73
                                                                            _ = v74
                                                                        }
                                                                        v53 += v43
                                                                        fmt.Println(v43)
                                                                        _ = len(v4)
                                                                    } else {
                                                                        switch fi41 {
                                                                        case 42:
                                                                            if (9223372036854775807 > (v49 * -1)) {
                                                                                v75 := GenS9{F0: 128, F1: (fi41 + (fi41 >> ((1) & 63))), F2: 0}
                                                                                _ = v75
                                                                            } else {
                                                                                var v76 int = -125
                                                                                _ = len(v3)
                                                                                _ = v76
                                                                            }
                                                                        case 4294967295:
                                                                            if true {
                                                                                if (v50 <= fi41) {
                                                                                    switch (((("hello" + "") + (("" + "") + "foo bar")) + ("\n\t" + "a")) + ("a" + "foo bar")) {
                                                                                    case "a":
                                                                                        v50 = 480
                                                                                    default:
                                                                                        v53 += 1e10
                                                                                    }
                                                                                    var v77 float64 = 1.5
                                                                                    _ = v77
                                                                                }
                                                                                v78 := v3[v50]
                                                                                _ = len(v3)
                                                                                v43 = v53
                                                                                var v79 bool = true
                                                                                var v80 int = (v50 % ((fi42) | 1))
                                                                                _ = v78
                                                                                _ = v79
                                                                                _ = v80
                                                                            } else {
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = "\n\t"
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v81 int8 = 42
                                                                            _ = v81
                                                                        }
                                                                        if false {
                                                                            var v82 uint64 = 2
                                                                            var v83 uint16 = 65535
                                                                            _ = v82
                                                                            _ = v83
                                                                        }
                                                                        var v84 int32 = -263
                                                                        if (1 > int32(fi42)) {
                                                                            v53 += v43
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                if ("foo bar" == "hello") {
                                                                                    var v85 int64 = (v49 * (v49 % (((v49 >> ((42) & 63))) | 1)))
                                                                                    _ = v85
                                                                                } else {
                                                                                    var v86 string = (((("foo bar" + "hello") + ("" + "hello")) + "") + ("a" + "a"))
                                                                                    _ = v86
                                                                                }
                                                                                switch 127 {
                                                                                case 771:
                                                                                    var v87 string = "a"
                                                                                    _ = v87
                                                                                case -2147483648:
                                                                                    var v88 float64 = -2.25
                                                                                    _ = v88
                                                                                case 32767:
                                                                                    var v89 uint32 = 2147483648
                                                                                    _ = v89
                                                                                default:
                                                                                    var v90 uint8 = 222
                                                                                    var v91 int = fi42
                                                                                    v92 := GenS6{F0: (fi41 + 32767)}
                                                                                    _ = v90
                                                                                    _ = v91
                                                                                    _ = v92
                                                                                }
                                                                                var v93 string = ("" + "hello")
                                                                                fmt.Println(v84)
                                                                                switch int(v49) {
                                                                                case -2147483648:
                                                                                    if (true && true) {
                                                                                        switch 1 {
                                                                                        case -816:
                                                                                            v93 = v93
                                                                                        case -2147483648:
                                                                                            var v94 int8 = -128
                                                                                            _ = v94
                                                                                        }
                                                                                        v53 += 0.0
                                                                                    }
                                                                                }
                                                                                switch 0 {
                                                                                case 32767:
                                                                                    switch ((("a" + "") + v93) + ((("" + (v93 + (v93 + "\n\t"))) + v93) + "\n\t")) {
                                                                                    case "hello":
                                                                                        var v95 bool = false
                                                                                        _ = v95
                                                                                    default:
                                                                                        delete(v1, "\n\t")
                                                                                    }
                                                                                }
                                                                                v43 -= (((((float32((-2.25 - 1.5)) / ((3.14159 - (float32(-2.25) - 1e10)) + (1.5 * v53))) * 1e10) - float32((1e10 * 0.0))) + v43) + ((-2.25 - (float32(((((0.1 + 3.14159) + float64(-2.25)) + -2.25) * 1e10)) - 0.0)) * 1e10))
                                                                                _ = v93
                                                                            default:
                                                                                var v96 uint = 255
                                                                                _ = v96
                                                                            }
                                                                        }
                                                                        _ = v84
                                                                    }
                                                                    v97 := GenS3{F0: 4294967295, F1: 19}
                                                                    _ = v97
                                                                }
                                                            case -2147483648:
                                                                var v98 int16 = -32768
                                                                _ = v98
                                                            default:
                                                                fmt.Println(v50, v50)
                                                            }
                                                        } else {
                                                            fmt.Println(v49)
                                                            v50 += 1
                                                            var v99 float32 = v43
                                                            if false {
                                                                _ = len(v3)
                                                                if false {
                                                                    v100 := GenS7{F0: v50, F1: ("\n\t" + "a"), F2: 363}
                                                                    _ = v100
                                                                } else {
                                                                    var v101 int16 = 32767
                                                                    _ = v101
                                                                }
                                                                if false {
                                                                    if (32767 >= 42) {
                                                                        if (((v99 * (v99 - v53)) == 0.1) || false) {
                                                                            if (true || (v99 >= (1e10 * 0.1))) {
                                                                                switch "\n\t" {
                                                                                case "foo bar":
                                                                                    fmt.Println(fi42)
                                                                                case "a":
                                                                                    var v102 int64 = v49
                                                                                    _ = v102
                                                                                case "\n\t":
                                                                                    v53 = 3.14159
                                                                                default:
                                                                                    v5 = append(v5, "hello")
                                                                                }
                                                                                switch (^(v50)) {
                                                                                case -9223372036854775808:
                                                                                    fmt.Println(fi41, v33, fi42)
                                                                                }
                                                                            } else {
                                                                                if (!true) {
                                                                                    var v103 int8 = 42
                                                                                    _ = v103
                                                                                } else {
                                                                                    _ = len(v1)
                                                                                    var v104 uint = 194
                                                                                    switch ("a" + ("\n\t" + "foo bar")) {
                                                                                    case "foo bar":
                                                                                        var v105 bool = (1 <= genFn1(1, fi42))
                                                                                        _ = v105
                                                                                    default:
                                                                                        v106 := GenS11{F0: -1, F1: 75}
                                                                                        _ = v106
                                                                                    }
                                                                                    v104 = v104
                                                                                    _ = v104
                                                                                }
                                                                                v53 += (float32(0.0) + 1e10)
                                                                                v43 = v53
                                                                            }
                                                                        } else {
                                                                            switch (("a" + ("foo bar" + "")) + "hello") {
                                                                            case "\n\t":
                                                                                v49 += (v49 % ((844) | 1))
                                                                            }
                                                                        }
                                                                        v107 := GenS12{F0: (v49 | 522)}
                                                                        _ = v107
                                                                    } else {
                                                                        fmt.Println(v43)
                                                                        v108 := GenS5{F0: 0}
                                                                        _ = v108
                                                                    }
                                                                }
                                                            }
                                                            _ = v99
                                                        }
                                                    case 127:
                                                        v109 := GenS13{F0: (v33 ^ (v33 << ((18446744073709551615) & 63)))}
                                                        _ = v109
                                                    }
                                                default:
                                                    v50 = -1
                                                }
                                            }
                                            var v110 int64 = (v49 / (((v49 * (v49 & v49))) | 1))
                                            fmt.Println(v53, fi42, v50)
                                            _ = v110
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = 2
                                            }
                                            v33 = (v33 | (v33 & v33))
                                            switch v50 {
                                            case -466:
                                                var v111 int32 = 1
                                                _ = v111
                                            case -9223372036854775808:
                                                switch (("" + ("a" + "a")) + "foo bar") {
                                                case "":
                                                    if (!false) {
                                                        fmt.Println(fi42, v49, v33)
                                                        v33 = (v33 * v33)
                                                        var v112 string = "\n\t"
                                                        _ = v112
                                                    } else {
                                                        if (1.5 <= 1e10) {
                                                            v113 := GenS2{F0: 65535}
                                                            v1["hello"] = 4294967295
                                                            fmt.Println(v49, fi42, v33)
                                                            _ = v113
                                                        }
                                                    }
                                                default:
                                                    var v114 int8 = 0
                                                    _ = v114
                                                }
                                            case -227:
                                                v115 := GenS6{F0: 128}
                                                _ = v115
                                            default:
                                                v53 -= (v53 / (0.0 / v43))
                                            }
                                        }
                                        _ = v54
                                    }
                                    var v116 uint = 673
                                    _ = v116
                                    _ = v49
                                    _ = v50
                                    _ = v53
                                }
                            }
                            switch "hello" {
                            case "":
                                var v117 float64 = 3.14159
                                _ = v117
                            case "\n\t":
                                v118 := GenS6{F0: (fi41 >> ((4294967295) & 63))}
                                _ = v118
                            case "foo bar":
                                v43 = v43
                            }
                            fmt.Println(v43, fi42)
                            var v119 int32 = -2147483648
                            v119 = 42
                            _ = fi42
                            _ = v119
                            _ = v43
                        }
                    }
                case 127:
                    var v120 int16 = 0
                    _ = v120
                default:
                    for fi121 := 0; fi121 < 3; fi121++ {
                        v33 -= (v33 & 1)
                        _ = fi121
                    }
                }
                v33 -= (^(v33))
            }
        } else {
            if ((-2.25 >= 0.0) || true) {
                v33 -= v33
            }
        }
        fmt.Println(v33, fi41)
        for fi122 := 0; fi122 < 13; fi122++ {
            v33 += v33
            if (1 >= 1) {
                var v123 bool = false
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v123
            }
            _ = fi122
        }
        _ = fi41
    }
    fmt.Println(v33)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
