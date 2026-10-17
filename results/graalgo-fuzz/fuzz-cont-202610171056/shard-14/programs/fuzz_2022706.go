package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 int; F1 int; F2 int32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint64; F1 int8; F2 int32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint16; F1 uint64; F2 uint }
type GenS12 struct { F0 bool; F1 int32; F2 uint32 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS18 struct { F0 bool; F1 int8; F2 int32 }
type GenS19 struct { F0 string; F1 int16 }
type GenS20 struct { F0 uint64 }

func main() {
    v1 := []uint32{2147483647}
    v2 := []uint32{36, 65535, 65535}
    v3 := []string{"\n\t", "a", "\n\t"}
    var v4 int64 = 9223372036854775807
    if (v4 == 9223372036854775807) {
        var v5 uint = 18446744073709551615
        switch ((("\n\t" + "hello") + "a") + "\n\t") {
        case "foo bar":
            if (true || (255 <= 127)) {
                var v6 uint32 = 4294967295
                v6 = 1
                _ = v6
            } else {
                for fi7 := 0; fi7 < 19; fi7++ {
                    v8 := GenS1{F0: 968}
                    if false {
                        switch 128 {
                        case -854:
                            for fi9 := 0; fi9 < 20; fi9++ {
                                switch ("" + "") {
                                case "foo bar":
                                    var v10 int64 = -758
                                    _ = v10
                                case "\n\t":
                                    v5 += 350
                                case "":
                                    fmt.Println(fi9)
                                default:
                                    switch ("\n\t" + "hello") {
                                    case "hello":
                                        v11 := GenS2{F0: ("hello" + "\n\t"), F1: (^(v5)), F2: 0}
                                        _ = v11
                                    case "":
                                        switch "foo bar" {
                                        case "foo bar":
                                            v12 := GenS1{F0: 377}
                                            _ = v12
                                        case "a":
                                            var v13 int16 = 1
                                            _ = v13
                                        case "\n\t":
                                            var v14 int = (fi7 * 414)
                                            _ = v14
                                        }
                                    default:
                                        var v15 int = fi7
                                        _ = v15
                                    }
                                }
                                _ = fi9
                            }
                        case 943:
                            if (false && false) {
                                _ = len(v1)
                            }
                        }
                        v5 -= (v5 & 2)
                    } else {
                        switch 2147483647 {
                        case 32768:
                            _ = len(v3)
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        switch ("hello" + (("\n\t" + ("\n\t" + ("\n\t" + (("\n\t" + "hello") + ("a" + (("hello" + ("" + "")) + ((("a" + ("hello" + "foo bar")) + "a") + "a"))))))) + (("hello" + "foo bar") + "hello"))) {
                        case "\n\t":
                            if false {
                                if (true && false) {
                                    fmt.Println(v5)
                                }
                                if true {
                                    fmt.Println(fi7, fi7, v5)
                                } else {
                                    switch "a" {
                                    case "":
                                        v5 -= (v5 >> (((^(v5))) & 63))
                                    }
                                }
                                fmt.Println(v5, v4)
                            } else {
                                var v16 int = fi7
                                _ = v16
                            }
                        case "":
                            v2 = append(v2, 65535)
                        }
                    }
                    _ = fi7
                    _ = v8
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                for fi17 := 0; fi17 < 8; fi17++ {
                    v5 = (^(v5))
                    _ = fi17
                }
            }
        case "\n\t":
            var v18 int = 42
            _ = v18
        default:
            v4 += (v4 << (((v5 << (((v5 % ((18446744073709551615) | 1))) & 63))) & 63))
        }
        _ = v5
    }
    var v19 int16 = 1000
    for fi20 := 0; fi20 < 10; fi20++ {
        if ((float32(float64(-2.25)) - 0.1) < 1.5) {
            var v21 uint8 = 81
            _ = v21
        }
        if true {
            if false {
                v4 = v4
                var v22 int8 = 2
                var v23 float32 = 0.0
                _ = v22
                _ = v23
            } else {
                v19 -= v19
                if true {
                    if (fi20 != -394) {
                        if (((true && false) || (!(!(("hello" + "") < "hello")))) && true) {
                            _ = len(v3)
                            v24 := GenS2{F0: "", F1: 138, F2: 1}
                            _ = v24
                        }
                        var v25 uint16 = 1
                        _ = v25
                    }
                    fmt.Println(v4)
                    fmt.Println(fi20, v4)
                }
                v26 := GenS3{F0: 4294967295}
                _ = v26
            }
            for fi27 := 0; fi27 < 16; fi27++ {
                if (fi20 >= (fi20 & -23)) {
                    v4 = (-(v4))
                    switch ("a" + "foo bar") {
                    case "":
                        v28 := GenS2{F0: "foo bar", F1: 18446744073709551615, F2: 42}
                        _ = v28
                    }
                }
                if false {
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (100 <= -108) {
                        v4 -= v4
                        switch ("\n\t" + "hello") {
                        case "a":
                            v19 += 806
                        case "foo bar":
                            v19 = (v19 / (((v19 << ((600) & 63))) | 1))
                            if (0 == 707) {
                                switch (fi27 << ((65535) & 63)) {
                                case 128:
                                    var v29 uint32 = 65535
                                    _ = v29
                                }
                                v30 := GenS4{F0: 294, F1: 0}
                                v19 -= v19
                                _ = v30
                            }
                        case "hello":
                            var v31 float64 = 1e10
                            v19 = 0
                            _ = v31
                        }
                    }
                } else {
                    if false {
                        var v32 int8 = 42
                        var v33 float64 = 0.0
                        if false {
                            fmt.Println(fi27)
                            v32 += (v32 << ((255) & 63))
                        } else {
                            if false {
                                v19 = v19
                                if (-511 < 548) {
                                    var v34 uint64 = 18446744073709551615
                                    v32 = int8(v4)
                                    v35 := GenS5{F0: fi27}
                                    _ = v34
                                    _ = v35
                                } else {
                                    var v36 int16 = v19
                                    _ = v36
                                }
                            } else {
                                if (!false) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if (fi27 < fi27) {
                                        var v37 int = fi20
                                        _ = v37
                                    } else {
                                        var v38 int8 = -122
                                        v3 = append(v3, ("foo bar" + ("" + ("hello" + ("\n\t" + "hello")))))
                                        if false {
                                            v1 = append(v1, 4294967295)
                                            if (210 < 127) {
                                                v33 += v33
                                                if ((fi20 + fi27) >= 2147483647) {
                                                    if (104 <= 1) {
                                                        if false {
                                                            var v39 float32 = 0.1
                                                            if ((255 < 65535) || (-337 >= (v19 >> ((866) & 63)))) {
                                                                switch "" {
                                                                case "foo bar":
                                                                    var v40 int = fi27
                                                                    _ = v40
                                                                default:
                                                                    v41 := GenS6{F0: 0, F1: -619}
                                                                    _ = v41
                                                                }
                                                                var v42 bool = ((!false) && (-2147483648 == int32(v19)))
                                                                v32 = (v32 % ((110) | 1))
                                                                _ = v42
                                                            }
                                                            if (!((-(v4)) > (v4 & v4))) {
                                                                switch (fi20 >> ((2) & 63)) {
                                                                case 971:
                                                                    var v43 int64 = v4
                                                                    _ = v43
                                                                case 127:
                                                                    v38 -= v32
                                                                case 32767:
                                                                    _ = len(v2)
                                                                default:
                                                                    if (42 < -1) {
                                                                        var v44 uint8 = 127
                                                                        v38 = v38
                                                                        if false {
                                                                            if (true && (1.5 >= v39)) {
                                                                                fmt.Println(v32)
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                v32 += -1
                                                                            } else {
                                                                                if (42 > -500) {
                                                                                    fmt.Println(v32, v19)
                                                                                    v45 := GenS7{F0: (fi20 - (-(fi27))), F1: (fi27 - 32767), F2: 1}
                                                                                    _ = v45
                                                                                }
                                                                                v46 := GenS4{F0: 0, F1: 0}
                                                                                if (false && false) {
                                                                                    switch (-(fi20)) {
                                                                                    case -295:
                                                                                        v32 = int8(fi27)
                                                                                    case 32768:
                                                                                        if false {
                                                                                            var v47 int64 = v4
                                                                                            v33 += v33
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                            _ = v47
                                                                                        } else {
                                                                                            var v48 int8 = v38
                                                                                            var v49 uint32 = 1
                                                                                            if ((v32 - v48) != -1) {
                                                                                                if false {
                                                                                                    fmt.Println(v39)
                                                                                                    v1 = append(v1, v49)
                                                                                                }
                                                                                                fmt.Println(v19, v32, v44)
                                                                                                if (1.5 < float32(((float64(v39) * -2.25) + (v33 + v33)))) {
                                                                                                    var v50 int64 = 1
                                                                                                    var v51 string = "\n\t"
                                                                                                    v52 := GenS2{F0: (v51 + "foo bar"), F1: 255, F2: 109}
                                                                                                    _ = v50
                                                                                                    _ = v51
                                                                                                    _ = v52
                                                                                                }
                                                                                            }
                                                                                            var v53 uint8 = v44
                                                                                            switch "\n\t" {
                                                                                            case "\n\t":
                                                                                                v54 := GenS8{F0: 146}
                                                                                                _ = v54
                                                                                            case "foo bar":
                                                                                                v55 := GenS9{F0: 4294967295, F1: (-(v32)), F2: 42}
                                                                                                _ = v55
                                                                                            case "hello":
                                                                                                v56 := GenS8{F0: 124}
                                                                                                v49 += (v49 & (v49 - 4294967295))
                                                                                                if ((!(true || ((65535 != 0) || (v4 >= (v4 << ((71) & 63)))))) || (true && true)) {
                                                                                                    var v57 int64 = (v4 ^ int64(v19))
                                                                                                    _ = v57
                                                                                                }
                                                                                                _ = v56
                                                                                            default:
                                                                                                if (((1e10 * v33) + ((0.0 + 1e10) + 3.14159)) < 1.5) {
                                                                                                    var v58 int = 794
                                                                                                    v59 := GenS1{F0: 2}
                                                                                                    _ = v58
                                                                                                    _ = v59
                                                                                                }
                                                                                            }
                                                                                            var v60 string = "hello"
                                                                                            _ = v48
                                                                                            _ = v49
                                                                                            _ = v53
                                                                                            _ = v60
                                                                                        }
                                                                                    case 1:
                                                                                        v61 := GenS7{F0: (-(fi27)), F1: fi27, F2: 1}
                                                                                        _ = v61
                                                                                    default:
                                                                                        if (1 != 0) {
                                                                                            v62 := GenS10{F0: 1}
                                                                                            v1 = append(v1, 779)
                                                                                            _ = v62
                                                                                        }
                                                                                    }
                                                                                    v38 -= (^(v38))
                                                                                    var v63 int64 = 155
                                                                                    _ = v63
                                                                                } else {
                                                                                    v32 = (-(v32))
                                                                                }
                                                                                _ = v46
                                                                            }
                                                                            if (376 > 691) {
                                                                                var v64 uint8 = (v44 * 255)
                                                                                _ = v64
                                                                            }
                                                                        }
                                                                        _ = v44
                                                                    } else {
                                                                        if true {
                                                                            var v65 uint64 = 18446744073709551615
                                                                            _ = v65
                                                                        }
                                                                        var v66 uint32 = 479
                                                                        v66 += (v66 * 2)
                                                                        _ = v66
                                                                    }
                                                                }
                                                                var v67 uint32 = 2147483647
                                                                fmt.Println(fi27)
                                                                _ = v67
                                                            } else {
                                                                v68 := GenS10{F0: 1}
                                                                var v69 float64 = v33
                                                                var v70 uint32 = 65535
                                                                _ = v68
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                            if (-9223372036854775808 <= fi20) {
                                                                fmt.Println(fi20)
                                                                if len(v2) > 0 {
                                                                    v2[0] = 4294967295
                                                                }
                                                            }
                                                            switch ("\n\t" + "a") {
                                                            case "a":
                                                                fmt.Println(fi20, v38, v33)
                                                            default:
                                                                v39 += (v39 / (float32(((v33 * float64((float32(v33) / ((v39 / (1.5 * (-2.25 / (0.1 * ((((0.0 / v39) * v39) / float32(((0.0 + v33) * float64(3.14159)))) + (-2.25 * (float32(v33) * -2.25))))))) * 1e10)))) - 1.5)) / (v39 * 3.14159)))
                                                            }
                                                            _ = v39
                                                        }
                                                    }
                                                    switch fi27 {
                                                    case -499:
                                                        if len(v2) > 0 {
                                                            v2[0] = 65535
                                                        }
                                                    default:
                                                        var v71 int16 = (-(v19))
                                                        _ = v71
                                                    }
                                                    v72 := GenS11{F0: 2, F1: 2, F2: 2}
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v72
                                                }
                                            } else {
                                                v73 := GenS11{F0: 500, F1: 1, F2: 637}
                                                _ = v73
                                            }
                                        }
                                        v74 := GenS1{F0: 255}
                                        _ = v38
                                        _ = v74
                                    }
                                }
                                var v75 int = (fi20 - 181)
                                v4 = int64(fi20)
                                _ = v75
                            }
                            switch 127 {
                            case -2147483648:
                                if false {
                                    if true {
                                        v33 += (v33 + 0.0)
                                    } else {
                                        var v76 int16 = v19
                                        var v77 int8 = (v32 - (v32 ^ (v32 + (v32 * -88))))
                                        v78 := GenS12{F0: false, F1: int32(fi27), F2: 2}
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                    }
                                } else {
                                    var v79 uint = 2
                                    v33 += 0.0
                                    var v80 int16 = (v19 % ((1) | 1))
                                    _ = v79
                                    _ = v80
                                }
                            }
                        }
                        _ = v32
                        _ = v33
                    }
                    fmt.Println(fi27, v19)
                    switch 9223372036854775807 {
                    case 729:
                        v4 = 42
                    case -184:
                        if true {
                            if false {
                                _ = len(v1)
                                v2 = append(v2, 0)
                                v81 := GenS13{F0: 0}
                                _ = v81
                            }
                            if false {
                                v82 := GenS7{F0: 32767, F1: fi20, F2: 821}
                                _ = v82
                            }
                            if (true && true) {
                                if true {
                                    if true {
                                        var v83 float64 = (0.1 * 0.1)
                                        v84 := GenS7{F0: fi20, F1: 2147483647, F2: 492}
                                        _ = v83
                                        _ = v84
                                    } else {
                                        v4 += v4
                                    }
                                    v85 := GenS14{F0: 445}
                                    _ = v85
                                }
                                v19 = (v19 | (v19 << ((0) & 63)))
                                if false {
                                    switch "" {
                                    case "a":
                                        var v86 int64 = v4
                                        _ = v86
                                    default:
                                        var v87 uint = 255
                                        _ = v87
                                    }
                                } else {
                                    fmt.Println(fi20, fi20)
                                    if (!(v4 > v4)) {
                                        if true {
                                            var v88 bool = false
                                            if true {
                                                v19 += (v19 << ((2) & 63))
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if true {
                                                    v88 = false
                                                }
                                            }
                                            _ = v88
                                        }
                                        var v89 uint32 = 516
                                        if true {
                                            v89 -= (v89 ^ (v89 ^ (v89 + 42)))
                                        } else {
                                            if false {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v19 -= v19
                                                    _ = len(v2)
                                                    var v90 float32 = 3.14159
                                                    _ = v90
                                                case "":
                                                    v4 = (-(v4))
                                                default:
                                                    var v91 float32 = -2.25
                                                    _ = v91
                                                }
                                                if false {
                                                    if len(v3) > 0 {
                                                        v3[0] = ((("hello" + "\n\t") + "\n\t") + "")
                                                    }
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v89 += v89
                                                    case "a":
                                                        var v92 float32 = -2.25
                                                        _ = v92
                                                    case "\n\t":
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            if ((0 == 95) || false) {
                                                                v4 -= 0
                                                                var v93 bool = false
                                                                _ = v93
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    var v94 uint8 = 2
                                                    var v95 uint64 = 823
                                                    _ = v94
                                                    _ = v95
                                                }
                                            }
                                        }
                                        switch (((("" + "hello") + (((("" + ("" + "a")) + "a") + "foo bar") + "foo bar")) + ("hello" + (("hello" + ("a" + "hello")) + "a"))) + ("foo bar" + "\n\t")) {
                                        case "":
                                            fmt.Println(v89, v4)
                                        }
                                        var v96 string = "foo bar"
                                        _ = v89
                                        _ = v96
                                    } else {
                                        var v97 uint8 = 177
                                        v4 = v4
                                        v97 -= (v97 / ((v97) | 1))
                                        _ = v97
                                    }
                                }
                            }
                        }
                    case 128:
                        switch ("foo bar" + ("foo bar" + ("foo bar" + ("" + ("hello" + "a"))))) {
                        case "a":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "":
                            var v98 int32 = 396
                            _ = v98
                        }
                    default:
                        v4 += v4
                    }
                    var v99 int8 = 37
                    _ = v99
                }
                switch ("hello" + (("hello" + ("a" + "a")) + "")) {
                case "":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v100 float64 = 0.1
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v101 := GenS1{F0: 697}
                    _ = v100
                    _ = v101
                case "hello":
                    v102 := GenS6{F0: 100, F1: 1}
                    _ = v102
                case "\n\t":
                    var v103 string = "hello"
                    _ = v103
                }
                _ = fi27
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            for fi104 := 0; fi104 < 6; fi104++ {
                v105 := GenS10{F0: -2147483648}
                v19 -= (v19 | -800)
                _ = fi104
                _ = v105
            }
            v19 = (v19 / ((v19) | 1))
            var v106 float64 = ((-2.25 / (-2.25 + 0.0)) + (1.5 - (-2.25 + 0.1)))
            _ = v106
        }
        _ = fi20
    }
    for fi107 := 0; fi107 < 5; fi107++ {
        for fi108 := 0; fi108 < 7; fi108++ {
            v19 -= (v19 % (((v19 % ((v19) | 1))) | 1))
            fmt.Println(fi107, fi107)
            _ = fi108
        }
        _ = fi107
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v109 := GenS15{F0: 2}
    var v110 uint64 = 782
    switch "" {
    case "foo bar":
        if (v4 >= 94) {
            for fi111 := 0; fi111 < 7; fi111++ {
                v112 := GenS14{F0: 513}
                for fi113 := 0; fi113 < 14; fi113++ {
                    var v114 uint32 = 1
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = len(v3)
                    _ = fi113
                    _ = v114
                }
                for fi115 := 0; fi115 < 9; fi115++ {
                    if (!(-2147483648 >= -2147483648)) {
                        v4 = (v4 / ((v4) | 1))
                        fmt.Println(fi115, fi115)
                    } else {
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        v2 = append(v2, 1)
                    }
                    v1 = append(v1, 42)
                    if (!false) {
                        var v116 int64 = (v4 + v4)
                        v116 += v4
                        _ = v116
                    }
                    var v117 uint32 = 65535
                    var v118 uint32 = (v117 >> ((18446744073709551615) & 63))
                    _ = fi115
                    _ = v117
                    _ = v118
                }
                _ = fi111
                _ = v112
            }
            for fi119 := 0; fi119 < 18; fi119++ {
                v2 = append(v2, 873)
                _ = fi119
            }
            for fi120 := 0; fi120 < 17; fi120++ {
                switch ("a" + "") {
                case "a":
                    if ((v4 % (((v4 * v4)) | 1)) < v4) {
                        var v121 uint32 = 744
                        var v122 int = fi120
                        if (255 > 243) {
                            for fi123 := 0; fi123 < 18; fi123++ {
                                v4 += v4
                                if true {
                                    switch "" {
                                    case "hello":
                                        fmt.Println(v19, v121)
                                        var v124 int8 = -74
                                        switch fi120 {
                                        case 9223372036854775807:
                                            if (1.5 != ((float64(0.1) + 1.5) / 3.14159)) {
                                                if ("" <= "") {
                                                    if false {
                                                        if (42 > 2) {
                                                            v4 = -450
                                                        }
                                                        var v125 bool = ((v19 / (((^(v19))) | 1)) < v19)
                                                        switch (fi123 / (((v122 ^ (-(fi123)))) | 1)) {
                                                        case 175:
                                                            if ((v122 * fi120) > fi120) {
                                                                v122 -= 42
                                                                var v126 int16 = 42
                                                                _ = v126
                                                            } else {
                                                                var v127 uint32 = 4294967295
                                                                v121 += (v127 / ((v127) | 1))
                                                                var v128 uint64 = v110
                                                                _ = v127
                                                                _ = v128
                                                            }
                                                        case 9223372036854775807:
                                                            switch "a" {
                                                            case "a":
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    if true {
                                                                        var v129 int16 = v19
                                                                        _ = v129
                                                                    } else {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = "a"
                                                                        }
                                                                        _ = len(v3)
                                                                        v130 := GenS6{F0: v124, F1: 42}
                                                                        _ = v130
                                                                    }
                                                                default:
                                                                    if (255 <= 933) {
                                                                        v131 := GenS11{F0: 9, F1: 2, F2: 42}
                                                                        v132 := GenS1{F0: 862}
                                                                        _ = v131
                                                                        _ = v132
                                                                    } else {
                                                                        var v133 int8 = v124
                                                                        if v125 {
                                                                            var v134 int8 = (-(v133))
                                                                            var v135 uint8 = 1
                                                                            if (v121 > v121) {
                                                                                _ = len(v1)
                                                                            }
                                                                            _ = v134
                                                                            _ = v135
                                                                        } else {
                                                                            v122 -= (^(fi120))
                                                                            var v136 uint = 42
                                                                            v137 := GenS8{F0: 1}
                                                                            var v138 uint = v136
                                                                            _ = v136
                                                                            _ = v137
                                                                            _ = v138
                                                                        }
                                                                        v122 += fi123
                                                                        _ = v133
                                                                    }
                                                                }
                                                            }
                                                        case 32768:
                                                            if v125 {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        }
                                                        var v139 int = v122
                                                        _ = v125
                                                        _ = v139
                                                    }
                                                    switch int(v110) {
                                                    case 128:
                                                        v122 = v122
                                                    default:
                                                        v4 -= 9223372036854775807
                                                    }
                                                }
                                            } else {
                                                if false {
                                                    fmt.Println(v19, v4, v4)
                                                    v4 -= int64(v19)
                                                    v140 := GenS16{F0: 1}
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = v140
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                        case 42:
                                            v122 = fi120
                                        case 1:
                                            v2 = append(v2, v121)
                                        }
                                        _ = v124
                                    case "":
                                        _ = len(v2)
                                    case "foo bar":
                                        if false {
                                            switch ("" + ("\n\t" + "foo bar")) {
                                            case "\n\t":
                                                var v141 uint8 = 0
                                                _ = v141
                                            }
                                        } else {
                                            var v142 uint = 65535
                                            if true {
                                                v4 = 0
                                                var v143 string = ""
                                                if false {
                                                    var v144 int8 = 100
                                                    v110 += (v110 * v110)
                                                    switch fi123 {
                                                    case -9223372036854775808:
                                                        switch v143 {
                                                        case "foo bar":
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        default:
                                                            if (false || true) {
                                                                v145 := GenS15{F0: v121}
                                                                fmt.Println(v4, v143)
                                                                _ = v145
                                                            } else {
                                                                if len(v1) > 0 {
                                                                    v1[0] = 2
                                                                }
                                                            }
                                                        }
                                                    case -1:
                                                        if (!false) {
                                                            var v146 string = "foo bar"
                                                            var v147 uint32 = v121
                                                            _ = v146
                                                            _ = v147
                                                        } else {
                                                            switch v143 {
                                                            case "a":
                                                                v148 := GenS4{F0: 65535, F1: 1}
                                                                _ = v148
                                                            case "":
                                                                var v149 uint8 = 2
                                                                _ = v149
                                                            }
                                                        }
                                                    case -2147483648:
                                                        var v150 uint64 = 18446744073709551615
                                                        _ = v150
                                                    }
                                                    _ = v144
                                                }
                                                if false {
                                                    fmt.Println(v19, v121, v143)
                                                }
                                                v19 = (v19 >> (((v142 << ((v142) & 63))) & 63))
                                                if true {
                                                    var v151 int = 32767
                                                    v152 := GenS17{F0: 0, F1: 934, F2: (v121 & v121)}
                                                    v122 += fi123
                                                    var v153 uint8 = 0
                                                    _ = v151
                                                    _ = v152
                                                    _ = v153
                                                } else {
                                                    v154 := GenS18{F0: false, F1: -109, F2: 980}
                                                    switch v122 {
                                                    case 127:
                                                        switch "a" {
                                                        case "":
                                                            if false {
                                                                var v155 int32 = 42
                                                                v156 := GenS11{F0: 432, F1: v110, F2: (v142 << ((42) & 63))}
                                                                _ = v155
                                                                _ = v156
                                                            }
                                                        case "foo bar":
                                                            v110 = (v110 & (v110 - (v110 * v110)))
                                                        case "\n\t":
                                                            v157 := GenS4{F0: uint16(v110), F1: 283}
                                                            _ = v157
                                                        default:
                                                            v2 = append(v2, 2147483648)
                                                        }
                                                    case -9223372036854775808:
                                                        var v158 int = (-(fi120))
                                                        _ = v158
                                                    default:
                                                        var v159 uint64 = v110
                                                        _ = v159
                                                    }
                                                    _ = v154
                                                }
                                                _ = v143
                                            }
                                            _ = v142
                                        }
                                    default:
                                        fmt.Println(v19)
                                    }
                                    var v160 int32 = 342
                                    fmt.Println(v4, v160, fi123)
                                    _ = v160
                                } else {
                                    var v161 bool = false
                                    _ = v161
                                }
                                _ = fi123
                            }
                            if (((v110 >> ((42) & 63)) < (v110 - (v110 - (v110 / (((v110 * v110)) | 1))))) && (-2.25 < ((1e10 * 0.0) / (0.0 - (-2.25 - 0.1))))) {
                                var v162 uint8 = 6
                                _ = v162
                            }
                            var v163 int16 = (v19 & v19)
                            _ = v163
                        }
                        _ = v121
                        _ = v122
                    }
                }
                _ = fi120
            }
        }
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "":
        for fi164 := 0; fi164 < 3; fi164++ {
            v165 := GenS19{F0: "\n\t", F1: (v19 >> ((uint(v110)) & 63))}
            v19 -= (v19 / ((v19) | 1))
            _ = fi164
            _ = v165
        }
    }
    var v166 float32 = 0.0
    if true {
        v110 -= (v110 * v110)
        if false {
            if len(v1) > 0 {
                v1[0] = 2
            }
            fmt.Println(v110)
            v19 = v19
        } else {
            var v167 string = "a"
            if true {
                if true {
                    if false {
                        if len(v1) > 0 {
                            v1[0] = 121
                        }
                        if (4 != 1) {
                            v168 := GenS20{F0: (v110 ^ (v110 + v110))}
                            _ = len(v1)
                            _ = v168
                        } else {
                            v167 = v167
                        }
                        if (false && ((v167 + "") != (v167 + "\n\t"))) {
                            if (-1 >= -120) {
                                if len(v3) > 0 {
                                    v3[0] = "foo bar"
                                }
                            }
                            var v169 uint = 2
                            _ = v169
                        }
                    } else {
                        v4 = -1
                        for fi170 := 0; fi170 < 6; fi170++ {
                            switch -211 {
                            case 42:
                                v166 = (v166 / v166)
                            case -2147483648:
                                if (-128 == 83) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            }
                            _ = fi170
                        }
                    }
                }
            }
            switch v167 {
            case "a":
                v19 -= 1
            case "foo bar":
                v110 = v110
            }
            _ = v167
        }
        v171 := GenS14{F0: 315}
        _ = v171
    }
    var v172 uint64 = (v110 | v110)
    for fi173 := 0; fi173 < 12; fi173++ {
        switch "foo bar" {
        case "a":
            fmt.Println(v172)
        case "hello":
            var v174 int = fi173
            _ = v174
        default:
            v172 -= v172
        }
        switch "foo bar" {
        case "\n\t":
            v19 = (v19 % (((v19 + -32768)) | 1))
        case "hello":
            _ = len(v1)
        }
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        _ = fi173
    }
    for fi175 := 0; fi175 < 3; fi175++ {
        v166 = (0.0 + -2.25)
        _ = fi175
    }
    v172 -= (v172 * (v172 >> ((255) & 63)))
    var v176 float32 = 0.0
    v176 -= v166
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v110:", v110)
    fmt.Println("v166:", v166)
    fmt.Println("v172:", v172)
    fmt.Println("v176:", v176)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v109
    _ = v110
    _ = v166
    _ = v172
    _ = v176
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
