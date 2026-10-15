package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int8; F2 int }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 int; F2 uint32 }
type GenS9 struct { F0 int8; F1 string; F2 uint64 }
type GenS10 struct { F0 int64; F1 string }
type GenS11 struct { F0 uint16; F1 uint; F2 int16 }
type GenS12 struct { F0 uint8; F1 int64 }
type GenS13 struct { F0 int }
func genFn0(p0 uint) uint32 {
    v7 := GenS1{F0: 1}
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := []int{(fi8 & int(p0))}
        p0 = (p0 >> ((p0) & 63))
        _ = fi8
        _ = v9
    }
    v10 := []int{127}
    _ = p0
    _ = v10
    _ = v7
    return 989
}

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    var v4 int64 = 9223372036854775807
    var v5 uint32 = 130
    v6 := []int{}
    if true {
        v4 -= 42
        v1 += (v1 | v4)
    } else {
        if ((!false) && true) {
            v11 := GenS2{F0: 2, F1: -128, F2: 0}
            for fi12 := 0; fi12 < 5; fi12++ {
                for fi13 := 0; fi13 < 19; fi13++ {
                    v6 = append(v6, (fi13 & fi13))
                    _ = fi13
                }
                for fi14 := 0; fi14 < 6; fi14++ {
                    var v15 uint8 = 149
                    fmt.Println(v15)
                    v1 -= -672
                    _ = fi14
                    _ = v15
                }
                var v16 uint = 255
                _ = fi12
                _ = v16
            }
            _ = v11
        } else {
            v17 := GenS1{F0: 127}
            switch ("" + (("\n\t" + "\n\t") + "\n\t")) {
            case "a":
                for fi18 := 0; fi18 < 4; fi18++ {
                    v19 := GenS2{F0: 255, F1: 1, F2: fi18}
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    for fi20 := 0; fi20 < 11; fi20++ {
                        var v21 int32 = -2147483648
                        _ = fi20
                        _ = v21
                    }
                    _ = fi18
                    _ = v19
                }
            case "":
                _ = len(v6)
            case "hello":
                var v22 uint = 320
                _ = v22
            }
            if true {
                delete(v2, 2147483647)
                v5 -= 4294967295
            } else {
                if len(v6) > 0 {
                    v6[0] = 0
                }
                v23 := GenS3{F0: ("foo bar" + ("\n\t" + ("hello" + ((("hello" + (("hello" + "\n\t") + "hello")) + ("" + ("foo bar" + "\n\t"))) + ((("\n\t" + "") + "\n\t") + ("a" + ("a" + "\n\t"))))))), F1: (42 < 9223372036854775807)}
                v24 := v2[(v5 & genFn0(98))]
                _ = v23
                _ = v24
            }
            _ = v17
        }
    }
    switch 770 {
    case -2147483648:
        for fi25 := 0; fi25 < 5; fi25++ {
            v26 := GenS4{F0: ("hello" + "hello"), F1: 128}
            fmt.Println(v1, v1, v5)
            v4 = v4
            _ = fi25
            _ = v26
        }
    case 128:
        v27 := GenS3{F0: "hello", F1: true}
        _ = v27
    }
    var v28 int8 = -100
    var v29 uint = 1
    v29 -= (v29 * v29)
    for fi30 := 0; fi30 < 11; fi30++ {
        var v31 int64 = v4
        fmt.Println(v29, v1, v5)
        for fi32 := 0; fi32 < 4; fi32++ {
            var v33 uint16 = 42
            fmt.Println(fi30, v33)
            _ = fi32
            _ = v33
        }
        _ = fi30
        _ = v31
    }
    for fi34 := 0; fi34 < 3; fi34++ {
        v28 = (v28 - (-(v28)))
        var v35 uint64 = uint64(v29)
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 16; fi36++ {
        var v37 int16 = -32768
        var v38 int64 = (v1 | v4)
        if false {
            _ = len(v6)
            var v39 int8 = v28
            _ = v39
        } else {
            if true {
                fmt.Println(v37)
                for fi40 := 0; fi40 < 18; fi40++ {
                    if ((^(v29)) < (v29 ^ v29)) {
                        switch (fi40 / (((^(fi40))) | 1)) {
                        case 32768:
                            if true {
                                if (!((1 <= 660) && true)) {
                                    v41 := v3[v5]
                                    v42 := GenS2{F0: 1, F1: (v28 + (v28 << ((v29) & 63))), F2: 2147483647}
                                    _ = v41
                                    _ = v42
                                } else {
                                    if ((v5 | genFn0((v29 << ((v29) & 63)))) == (v5 / (((v5 + (v5 - (v5 % ((v5) | 1))))) | 1))) {
                                        v5 = 42
                                    }
                                    var v43 uint16 = 2
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = v43
                                }
                            }
                            var v44 uint8 = 4
                            _ = v44
                        }
                        v37 += v37
                        var v45 int16 = (-(v37))
                        _ = v45
                    }
                    _ = fi40
                }
                switch ("a" + ("hello" + "")) {
                case "foo bar":
                    fmt.Println(fi36, v37)
                case "\n\t":
                    if false {
                        if (0 <= uint16(v29)) {
                            var v46 uint = 127
                            for fi47 := 0; fi47 < 1; fi47++ {
                                v37 = 32767
                                v46 = v46
                                if (true && ((false || false) && false)) {
                                    fmt.Println(v28, fi47, fi47)
                                }
                                _ = fi47
                            }
                            _ = v46
                        } else {
                            var v48 int64 = (v4 ^ 0)
                            _ = v48
                        }
                    }
                case "hello":
                    switch (fi36 - fi36) {
                    case 42:
                        if false {
                            for fi49 := 0; fi49 < 12; fi49++ {
                                var v50 uint32 = v5
                                fmt.Println(v50, fi49)
                                fmt.Println(v4, v4, v29)
                                switch (fi49 - fi49) {
                                case 128:
                                    if (v29 < v29) {
                                        fmt.Println(v5)
                                        v51 := GenS5{F0: (fi49 / (((fi49 / ((823) | 1))) | 1))}
                                        var v52 int64 = -1
                                        _ = v51
                                        _ = v52
                                    }
                                default:
                                    if true {
                                        switch fi36 {
                                        case -2147483648:
                                            switch -1 {
                                            case -731:
                                                v53 := GenS6{F0: 18446744073709551615, F1: 1}
                                                _ = v53
                                            default:
                                                fmt.Println(fi49, v38, v5)
                                            }
                                        default:
                                            v29 = (v29 ^ 290)
                                        }
                                        if (!true) {
                                            switch "a" {
                                            case "foo bar":
                                                if (v29 == 533) {
                                                    v54 := v2[(v5 / ((v50) | 1))]
                                                    fmt.Println(v1, v38)
                                                    if true {
                                                        if (!false) {
                                                            v29 = (v29 / ((v29) | 1))
                                                            switch v54 {
                                                            case "foo bar":
                                                                v4 -= (v1 << (((v29 & 427)) & 63))
                                                            }
                                                        } else {
                                                            v55 := GenS5{F0: (fi49 % ((fi36) | 1))}
                                                            fmt.Println(v4)
                                                            var v56 int16 = (v37 / ((v37) | 1))
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                        if true {
                                                            if false {
                                                                v5 = (v50 | (v50 & (v5 + (^(v5)))))
                                                                _ = len(v6)
                                                                var v57 int8 = v28
                                                                v58 := GenS1{F0: (v57 - (v28 / ((61) | 1)))}
                                                                var v59 int8 = v57
                                                                _ = v57
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                            fmt.Println(v54, v28)
                                                        } else {
                                                            fmt.Println(v4, v37)
                                                            var v60 float32 = 0.0
                                                            _ = v60
                                                        }
                                                        _ = len(v2)
                                                        switch v54 {
                                                        case "hello":
                                                            _ = len(v6)
                                                        case "\n\t":
                                                            v38 = (v1 + v4)
                                                            v61 := GenS3{F0: "foo bar", F1: ((!(1 < 42)) || (!true))}
                                                            delete(v2, 4294967295)
                                                            _ = v61
                                                        default:
                                                            v37 = v37
                                                        }
                                                        fmt.Println(v5, v37, v54)
                                                        v38 = (v4 << ((2) & 63))
                                                    } else {
                                                        v29 = v29
                                                    }
                                                    _ = v54
                                                } else {
                                                    switch "a" {
                                                    case "foo bar":
                                                        fmt.Println(v28)
                                                    }
                                                    fmt.Println(v28, v28)
                                                    switch (fi49 / ((-114) | 1)) {
                                                    case 128:
                                                        var v62 bool = true
                                                        _ = v62
                                                    default:
                                                        if false {
                                                            v1 -= v38
                                                        }
                                                    }
                                                }
                                            case "":
                                                var v63 int = fi49
                                                _ = v63
                                            }
                                            fmt.Println(v5, v28, v29)
                                            switch (("a" + "") + "a") {
                                            case "a":
                                                if (!((genFn0(v29) == v50) && (2147483647 <= (^(v50))))) {
                                                    var v64 string = "\n\t"
                                                    _ = v64
                                                }
                                                v65 := GenS7{F0: ("" + ("\n\t" + ("\n\t" + (("foo bar" + "\n\t") + ("" + ("" + ((("" + "a") + "foo bar") + ("\n\t" + "\n\t"))))))))}
                                                var v66 int64 = -1
                                                _ = v65
                                                _ = v66
                                            case "":
                                                switch (fi49 / ((fi49) | 1)) {
                                                case -2147483648:
                                                    switch (fi36 ^ (fi36 % ((fi36) | 1))) {
                                                    case 4294967295:
                                                        delete(v2, v50)
                                                    case 32768:
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    case 685:
                                                        switch "" {
                                                        case "hello":
                                                            fmt.Println(v37)
                                                        case "a":
                                                            fmt.Println(v28, fi36, v29)
                                                        default:
                                                            v29 = (v29 & (^(v29)))
                                                        }
                                                    }
                                                default:
                                                    var v67 string = (((((("foo bar" + "") + "hello") + "") + "") + (("\n\t" + "\n\t") + "hello")) + "hello")
                                                    _ = v67
                                                }
                                            case "\n\t":
                                                v28 += (^(v28))
                                            default:
                                                switch fi36 {
                                                case 1:
                                                    var v68 int16 = (^(v37))
                                                    _ = v68
                                                }
                                                if ((^(v37)) <= 1) {
                                                    v69 := GenS8{F0: (!true), F1: 710, F2: genFn0(1)}
                                                    v70 := GenS2{F0: 501, F1: v28, F2: (-(fi36))}
                                                    _ = len(v6)
                                                    _ = v69
                                                    _ = v70
                                                } else {
                                                    fmt.Println(fi49, v29)
                                                    v71 := GenS3{F0: ("a" + "\n\t"), F1: ((-2.25 + (((1e10 * 0.0) - (0.1 + 1.5)) / (((-2.25 / 1.5) - (0.0 + (0.1 * 0.1))) + 1e10))) == ((((((float64((0.1 + 3.14159)) / -2.25) / 1.5) - ((0.1 + (-2.25 + 1.5)) - 1e10)) + 0.1) - -2.25) * 3.14159))}
                                                    delete(v3, 2147483647)
                                                    _ = v71
                                                }
                                            }
                                        } else {
                                            var v72 float32 = 1.5
                                            v3[4294967295] = 2147483647
                                            _ = v72
                                        }
                                    } else {
                                        fmt.Println(fi36)
                                    }
                                    fmt.Println(v28)
                                }
                                v50 += (v5 / ((673) | 1))
                                v5 += v5
                                _ = fi49
                                _ = v50
                            }
                        } else {
                            switch 2147483647 {
                            case 0:
                                var v73 uint16 = 42
                                _ = v73
                            }
                            for fi74 := 0; fi74 < 13; fi74++ {
                                v4 = (v4 % ((v4) | 1))
                                _ = fi74
                            }
                        }
                    }
                }
            }
            _ = len(v6)
        }
        _ = fi36
        _ = v37
        _ = v38
    }
    if false {
        for fi75 := 0; fi75 < 3; fi75++ {
            v76 := GenS9{F0: (v28 & (v28 + -128)), F1: "a", F2: 1}
            var v77 string = (("\n\t" + "\n\t") + "a")
            _ = fi75
            _ = v76
            _ = v77
        }
        v78 := GenS10{F0: (^(v1)), F1: "hello"}
        _ = v78
    } else {
        switch ("" + "\n\t") {
        case "\n\t":
            if (!(0 > (v4 >> ((42) & 63)))) {
                v29 -= 65535
                v6 = append(v6, 42)
            }
        }
        switch "hello" {
        case "hello":
            if (!false) {
                v29 = (v29 + (v29 >> ((v29) & 63)))
                var v79 uint16 = 1
                _ = v79
            } else {
                if (-371 <= 0) {
                    for fi80 := 0; fi80 < 10; fi80++ {
                        v81 := GenS4{F0: ("foo bar" + "hello"), F1: 70}
                        _ = fi80
                        _ = v81
                    }
                    if false {
                        var v82 string = "foo bar"
                        fmt.Println(v82)
                        _ = v82
                    }
                } else {
                    var v83 uint64 = 754
                    for fi84 := 0; fi84 < 18; fi84++ {
                        for fi85 := 0; fi85 < 19; fi85++ {
                            var v86 uint16 = 65535
                            switch fi84 {
                            case -9223372036854775808:
                                if (!(false || (false || (!true)))) {
                                    v87 := GenS8{F0: (!true), F1: 304, F2: (v5 ^ 2147483648)}
                                    switch "a" {
                                    case "\n\t":
                                        fmt.Println(v83, v1, fi84)
                                    case "hello":
                                        v88 := GenS4{F0: "", F1: 255}
                                        _ = v88
                                    default:
                                        if true {
                                            fmt.Println(v29)
                                            var v89 uint = uint(v5)
                                            _ = len(v2)
                                            _ = v89
                                        }
                                    }
                                    _ = v87
                                }
                            case -152:
                                v86 += 255
                            default:
                                if (!true) {
                                    _ = len(v6)
                                    v6 = append(v6, -2147483648)
                                    v5 -= uint32(v83)
                                    _ = len(v2)
                                    var v90 int64 = 0
                                    v1 = v90
                                    v91 := v3[2]
                                    fmt.Println(v91)
                                    _ = v90
                                    _ = v91
                                }
                                var v92 int32 = -243
                                if len(v6) > 0 {
                                    v6[0] = fi84
                                }
                                var v93 uint64 = 1
                                _ = v92
                                _ = v93
                            }
                            _ = fi85
                            _ = v86
                        }
                        var v94 uint64 = uint64(v29)
                        _ = fi84
                        _ = v94
                    }
                    for fi95 := 0; fi95 < 10; fi95++ {
                        for fi96 := 0; fi96 < 9; fi96++ {
                            var v97 uint64 = (v83 ^ (v83 - (v83 >> ((v29) & 63))))
                            _ = fi96
                            _ = v97
                        }
                        var v98 uint = uint(v83)
                        _ = fi95
                        _ = v98
                    }
                    _ = v83
                }
                v99 := GenS11{F0: 0, F1: (v29 | v29), F2: -32768}
                v100 := GenS3{F0: (("foo bar" + "a") + ("foo bar" + ("" + ""))), F1: (genFn0((^(v29))) >= 1)}
                _ = v100
                _ = v99
            }
        case "\n\t":
            if ((3.14159 != 1.5) || true) {
                v1 -= 1
                v28 -= 86
                v28 = v28
            } else {
                fmt.Println(v28, v29, v28)
            }
        case "a":
            v101 := GenS12{F0: 0, F1: (v1 & 0)}
            _ = v101
        }
        var v102 int32 = -366
        _ = v102
    }
    v103 := v2[v5]
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = len(v6)
    var v104 int16 = 858
    delete(v2, (v5 >> ((255) & 63)))
    v29 = v29
    if true {
        delete(v2, 324)
    }
    var v105 float64 = 1e10
    if true {
        v106 := GenS6{F0: 343, F1: 255}
        _ = len(v2)
        _ = v106
    } else {
        for fi107 := 0; fi107 < 12; fi107++ {
            if false {
                for fi108 := 0; fi108 < 20; fi108++ {
                    if true {
                        v4 += v4
                    } else {
                        var v109 string = "a"
                        switch fi107 {
                        case 805:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        default:
                            v2[genFn0((^(v29)))] = v109
                        }
                        fmt.Println(v5, v109)
                        _ = v109
                    }
                    switch 32768 {
                    case 32768:
                        v5 -= (v5 * 387)
                    case 921:
                        var v110 int32 = 2147483647
                        _ = v110
                    case 128:
                        if ((1 > 2) || false) {
                            if (!(334 == 4294967295)) {
                                v103 = "foo bar"
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if false {
                                    var v111 uint = (v29 << ((65535) & 63))
                                    v29 = v29
                                    _ = v111
                                }
                            }
                            v112 := GenS11{F0: 64, F1: (v29 + (v29 & (^(v29)))), F2: (v104 - v104)}
                            _ = v112
                        }
                    default:
                        if (2147483647 < -1) {
                            if true {
                                v2[(v5 | (v5 / ((v5) | 1)))] = ("foo bar" + v103)
                            }
                            v28 = v28
                            v113 := v2[(v5 - (^(v5)))]
                            _ = v113
                        }
                    }
                    v105 = (0.1 - v105)
                    _ = fi108
                }
                var v114 int32 = int32(v104)
                v115 := GenS13{F0: fi107}
                _ = v114
                _ = v115
            } else {
                v116 := GenS6{F0: 62, F1: 492}
                _ = v116
            }
            var v117 int16 = v104
            switch (v103 + v103) {
            case "":
                v4 = v4
            case "hello":
                var v118 float32 = (float32(((v105 / v105) * ((((v105 * v105) / 1e10) * 1e10) + v105))) / float32(((v105 - 1e10) * (v105 - 1e10))))
                _ = v118
            case "\n\t":
                if false {
                    if true {
                        fmt.Println(v29, v4)
                        switch ((v103 + v103) + (v103 + (v103 + "foo bar"))) {
                        case "hello":
                            if true {
                                fmt.Println(v117, fi107)
                                v119 := GenS3{F0: v103, F1: (false && false)}
                                var v120 uint16 = 2
                                _ = v119
                                _ = v120
                            }
                        case "a":
                            var v121 float32 = 0.0
                            _ = v121
                        default:
                            var v122 int16 = (-(v104))
                            _ = v122
                        }
                        v117 -= -648
                    }
                    fmt.Println(v4)
                    for fi123 := 0; fi123 < 20; fi123++ {
                        var v124 bool = ((3.14159 / (v105 - 1e10)) <= ((float64(float32(1.5)) + (1e10 - 1.5)) - v105))
                        _ = fi123
                        _ = v124
                    }
                } else {
                    delete(v3, genFn0(v29))
                    v5 = 2147483647
                    v125 := GenS11{F0: 255, F1: (v29 - 820), F2: v104}
                    _ = v125
                }
            }
            _ = fi107
            _ = v117
        }
        var v126 uint32 = 42
        _ = v126
    }
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v103
    _ = v104
    _ = v105
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
