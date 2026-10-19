package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int64; F2 bool }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int32; F1 uint64; F2 int }
type GenS5 struct { F0 int64; F1 string; F2 uint16 }
type GenS6 struct { F0 int16; F1 int64; F2 int32 }
type GenS7 struct { F0 string; F1 string; F2 int64 }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 int32; F1 uint32; F2 int32 }
type GenS10 struct { F0 int16; F1 string }
type GenS11 struct { F0 bool; F1 uint32 }
type GenS12 struct { F0 uint32; F1 string; F2 int }
type GenS13 struct { F0 int8; F1 int64 }
func genFn0(p0 int64, p1 string, p2 uint16) string {
    if (!((true && ((1.5 / 1e10) != (3.14159 / 0.1))) && true)) {
        p2 += p2
        v4 := []int{128, 1, 4294967295}
        var v5 float64 = ((((1.5 + 0.1) / 1.0) / 3.14159) * -2.25)
        _ = v4
        _ = v5
    } else {
        var v6 bool = (false && (true && true))
        var v7 int16 = 32767
        v7 = v7
        _ = v6
        _ = v7
    }
    fmt.Println(p1, p0, p0)
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    return ((p1 + p1) + p1)
}
func genFn1(p0 int16, p1 int64) string {
    fmt.Println(p1, p0)
    p0 = (p0 | p0)
    var v8 uint32 = 2
    v9 := []int{4294967295, 9223372036854775807, 32767, 215}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return ("foo bar" + ("hello" + "hello"))
}

func main() {
    v1 := []int{}
    var v2 int32 = 2147483647
    v3 := make(map[uint32]string)
    v10 := GenS1{F0: ("\n\t" + ("hello" + "foo bar"))}
    var v11 int8 = 127
    v11 -= v11
    fmt.Println(v11, v11)
    if false {
        v12 := GenS1{F0: ""}
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 float64 = 0.1
            switch "a" {
            case "foo bar":
                for fi15 := 0; fi15 < 18; fi15++ {
                    var v16 int = fi15
                    _ = fi15
                    _ = v16
                }
            case "\n\t":
                fmt.Println(v14, v14)
            default:
                _ = len(v1)
            }
            var v17 int16 = 32767
            _ = fi13
            _ = v14
            _ = v17
        }
        _ = v12
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 uint = 4294967295
        for fi20 := 0; fi20 < 15; fi20++ {
            if (!(((-2.25 >= 1.5) || (((0.1 * 1.5) - 0.1) == (3.14159 / ((float32(((0.1 * 1.5) * float64(-2.25))) + ((1.5 - (1.5 + (0.0 - (float32((float64((0.0 * -2.25)) * 0.1)) - 0.1)))) * float32(float64(3.14159)))) / -2.25)))) && true)) {
                var v21 uint16 = 1
                _ = v21
            }
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    _ = len(v1)
    if false {
        switch "\n\t" {
        case "foo bar":
            v11 = (v11 << ((818) & 63))
        case "hello":
            for fi22 := 0; fi22 < 16; fi22++ {
                delete(v3, 738)
                if (1 != 9223372036854775807) {
                    fmt.Println(v11)
                    for fi23 := 0; fi23 < 14; fi23++ {
                        v11 += -33
                        _ = fi23
                    }
                    v24 := GenS1{F0: ("hello" + "foo bar")}
                    _ = v24
                }
                _ = fi22
            }
        case "a":
            v11 = 1
        }
        for fi25 := 0; fi25 < 2; fi25++ {
            fmt.Println(v2, v2, v11)
            _ = fi25
        }
    }
    v2 += v2
    var v26 uint32 = 96
    if (true || true) {
        for fi27 := 0; fi27 < 6; fi27++ {
            var v28 int = (fi27 * 128)
            switch "a" {
            case "foo bar":
                for fi29 := 0; fi29 < 18; fi29++ {
                    switch "foo bar" {
                    case "foo bar":
                        if (!(!(!true))) {
                            if false {
                                var v30 int8 = v11
                                _ = v30
                            }
                            var v31 uint = 853
                            _ = v31
                        }
                    default:
                        v28 = v28
                        switch v28 {
                        case 4294967295:
                            delete(v3, (^(v26)))
                        case -9223372036854775808:
                            var v32 int = (fi29 / ((v28) | 1))
                            _ = v32
                        case 127:
                            v11 = (v11 - 1)
                        }
                    }
                    _ = fi29
                }
            case "a":
                fmt.Println(v11, fi27, v28)
            case "":
                for fi33 := 0; fi33 < 17; fi33++ {
                    var v34 uint = 688
                    var v35 int64 = 42
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            default:
                if false {
                    switch fi27 {
                    case 0:
                        if false {
                            _ = len(v3)
                            v36 := GenS2{F0: 1, F1: int64(v2), F2: (42 < 189)}
                            _ = v36
                        }
                    }
                    v37 := GenS2{F0: 1, F1: 42, F2: false}
                    _ = v37
                }
            }
            v2 += v2
            _ = fi27
            _ = v28
        }
        if (false && true) {
            var v38 int64 = -1
            var v39 uint16 = 42
            _ = v38
            _ = v39
        } else {
            var v40 string = ""
            v3[2] = v40
            _ = v40
        }
    } else {
        fmt.Println(v11, v26)
        v26 = 0
    }
    if false {
        if (109 <= 128) {
            fmt.Println(v26, v2)
            fmt.Println(v26, v2)
        }
        v2 += 42
    } else {
        var v41 string = ("" + "\n\t")
        _ = v41
    }
    switch 1 {
    case 1:
        var v42 uint16 = 42
        _ = v42
    default:
        v11 -= 0
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        if false {
            v44 := GenS3{F0: 55, F1: 2}
            v45 := GenS3{F0: 107, F1: 65535}
            v46 := GenS2{F0: 2, F1: 1, F2: false}
            _ = v44
            _ = v45
            _ = v46
        } else {
            var v47 float32 = ((1.5 + float32(1e10)) - 0.0)
            _ = v47
        }
        _ = fi43
    }
    if (1000 >= -1) {
        for fi48 := 0; fi48 < 4; fi48++ {
            switch (-(fi48)) {
            case 128:
                switch ("foo bar" + ("" + ("\n\t" + ""))) {
                case "\n\t":
                    var v49 int32 = v2
                    _ = v49
                case "hello":
                    var v50 int = 32768
                    _ = v50
                }
            case -353:
                v51 := GenS4{F0: v2, F1: uint64(v26), F2: fi48}
                _ = v51
            case 4294967295:
                v2 = v2
            default:
                for fi52 := 0; fi52 < 19; fi52++ {
                    var v53 float64 = -2.25
                    _ = fi52
                    _ = v53
                }
            }
            _ = fi48
        }
        if (1 <= 0) {
            var v54 uint64 = 1
            v3[v26] = "hello"
            _ = v54
        } else {
            v55 := v3[v26]
            if ((!false) && (true || true)) {
                v11 = v11
                var v56 uint = 65535
                switch 1 {
                case 610:
                    v57 := GenS2{F0: 32767, F1: -9223372036854775808, F2: true}
                    _ = v57
                default:
                    if (((1e10 / 1e10) - (1.5 - 1.5)) <= 1.5) {
                        if false {
                            switch 32767 {
                            case 459:
                                v58 := GenS5{F0: -1, F1: (((v55 + v55) + ("hello" + v55)) + v55), F2: 735}
                                _ = v58
                            case 128:
                                switch 32767 {
                                case 9223372036854775807:
                                    v11 = v11
                                case 163:
                                    var v59 int64 = -738
                                    _ = v59
                                }
                            case 0:
                                var v60 int8 = -127
                                _ = v60
                            default:
                                v61 := GenS3{F0: 1, F1: v56}
                                _ = v61
                            }
                        } else {
                            if ((v26 ^ 768) <= v26) {
                                v62 := GenS1{F0: "hello"}
                                v63 := GenS5{F0: 1, F1: (v55 + "\n\t"), F2: 1}
                                if (true || ((v56 / ((216) | 1)) != v56)) {
                                    fmt.Println(v26)
                                    v26 -= (v26 << ((2) & 63))
                                    v64 := GenS5{F0: 1, F1: v55, F2: 0}
                                    _ = v64
                                }
                                _ = v62
                                _ = v63
                            } else {
                                v55 = (v55 + (v55 + v55))
                                var v65 float32 = (1e10 * 1.5)
                                _ = v65
                            }
                        }
                        switch 1 {
                        case -2147483648:
                            delete(v3, 42)
                        case 121:
                            fmt.Println(v2)
                        case -9223372036854775808:
                            for fi66 := 0; fi66 < 3; fi66++ {
                                var v67 int64 = 1
                                for fi68 := 0; fi68 < 13; fi68++ {
                                    fmt.Println(v55)
                                    _ = fi68
                                }
                                fmt.Println(fi66, v11)
                                _ = fi66
                                _ = v67
                            }
                        }
                        if (true && false) {
                            v2 = (v2 - 1)
                            var v69 uint = 388
                            _ = v69
                        } else {
                            fmt.Println(v11)
                            for fi70 := 0; fi70 < 16; fi70++ {
                                var v71 int64 = 1
                                _ = fi70
                                _ = v71
                            }
                        }
                    } else {
                        switch v55 {
                        case "hello":
                            v56 -= v56
                        case "":
                            fmt.Println(v56, v56)
                        case "\n\t":
                            if (!false) {
                                switch v55 {
                                case "foo bar":
                                    switch -1 {
                                    case -2147483648:
                                        delete(v3, (v26 >> ((1) & 63)))
                                    case 601:
                                        fmt.Println(v55)
                                    default:
                                        if false {
                                            v55 = v55
                                            var v72 int = 42
                                            var v73 uint = 1
                                            _ = v72
                                            _ = v73
                                        } else {
                                            fmt.Println(v55, v11, v26)
                                            v74 := GenS6{F0: 1000, F1: -1, F2: v2}
                                            v11 = v11
                                            _ = v74
                                        }
                                    }
                                }
                                for fi75 := 0; fi75 < 9; fi75++ {
                                    fmt.Println(v56)
                                    if (!(2 >= 78)) {
                                        v2 = v2
                                    }
                                    var v76 int16 = 1000
                                    _ = fi75
                                    _ = v76
                                }
                            } else {
                                var v77 int8 = v11
                                v26 = v26
                                switch 128 {
                                case 32767:
                                    var v78 int8 = v11
                                    _ = v78
                                case 32768:
                                    v3[9] = "hello"
                                case 4294967295:
                                    for fi79 := 0; fi79 < 15; fi79++ {
                                        switch "" {
                                        case "foo bar":
                                            _ = len(v3)
                                        }
                                        v56 = v56
                                        v2 -= (v2 >> (((^(v56))) & 63))
                                        _ = fi79
                                    }
                                }
                                _ = v77
                            }
                        }
                        switch -1 {
                        case 128:
                            if (127 != 43) {
                                switch ((v55 + v55) + v55) {
                                case "hello":
                                    v56 = (v56 ^ v56)
                                case "foo bar":
                                    for fi80 := 0; fi80 < 18; fi80++ {
                                        switch "foo bar" {
                                        case "a":
                                            for fi81 := 0; fi81 < 5; fi81++ {
                                                switch (fi80 << ((2) & 63)) {
                                                case 4294967295:
                                                    switch fi80 {
                                                    case 127:
                                                        v11 += (v11 & (v11 - -56))
                                                    case 128:
                                                        fmt.Println(fi81, fi80)
                                                    }
                                                }
                                                v82 := GenS7{F0: v55, F1: v55, F2: 9223372036854775807}
                                                v83 := GenS8{F0: v26, F1: v26}
                                                switch 32768 {
                                                case 1:
                                                    v11 -= v11
                                                case 42:
                                                    _ = len(v3)
                                                }
                                                _ = fi81
                                                _ = v82
                                                _ = v83
                                            }
                                        }
                                        _ = fi80
                                    }
                                default:
                                    v55 = (v55 + v55)
                                }
                                v84 := GenS9{F0: 0, F1: v26, F2: -2147483648}
                                _ = v84
                            } else {
                                fmt.Println(v11, v26, v55)
                                var v85 int32 = (v2 & (v2 + 991))
                                v11 += (v11 & (v11 * 100))
                                _ = v85
                            }
                        default:
                            fmt.Println(v55)
                        }
                        v56 = (v56 % ((4294967295) | 1))
                    }
                }
                _ = v56
            }
            _ = v55
        }
        for fi86 := 0; fi86 < 2; fi86++ {
            switch (-(fi86)) {
            case 127:
                switch ("" + "hello") {
                case "a":
                    delete(v3, (v26 / (((^(v26))) | 1)))
                case "\n\t":
                    _ = len(v1)
                default:
                    if true {
                        if (-353 == 9223372036854775807) {
                            var v87 uint8 = 64
                            var v88 float32 = 1.5
                            _ = v87
                            _ = v88
                        } else {
                            _ = len(v1)
                        }
                        switch fi86 {
                        case 42:
                            var v89 int16 = -632
                            _ = v89
                        case 4294967295:
                            v11 += -128
                        case -2147483648:
                            switch "a" {
                            case "\n\t":
                                var v90 string = ((("\n\t" + ("hello" + "")) + ((("a" + (("hello" + "a") + (("a" + "foo bar") + "a"))) + "a") + "\n\t")) + "\n\t")
                                _ = v90
                            case "foo bar":
                                v3[0] = (("foo bar" + ((("" + "foo bar") + "hello") + ("hello" + "\n\t"))) + ("\n\t" + "a"))
                            case "":
                                if true {
                                    for fi91 := 0; fi91 < 6; fi91++ {
                                        if true {
                                            var v92 int32 = 2147483647
                                            _ = v92
                                        }
                                        _ = fi91
                                    }
                                    var v93 string = (("foo bar" + "a") + "hello")
                                    _ = v93
                                }
                            }
                        default:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v11 = v11
                    } else {
                        for fi94 := 0; fi94 < 14; fi94++ {
                            v2 = v2
                            var v95 uint32 = (^(v26))
                            fmt.Println(v26)
                            _ = fi94
                            _ = v95
                        }
                        for fi96 := 0; fi96 < 10; fi96++ {
                            if true {
                                if false {
                                    var v97 float64 = (-2.25 / float64(((1e10 - -2.25) - 3.14159)))
                                    v98 := GenS1{F0: ("foo bar" + (("hello" + "") + ("hello" + ("foo bar" + "\n\t"))))}
                                    var v99 int8 = 0
                                    var v100 int8 = 0
                                    _ = v100
                                    _ = v97
                                    _ = v98
                                    _ = v99
                                }
                                v101 := GenS4{F0: 42, F1: 1, F2: 128}
                                _ = v101
                            }
                            switch (-(fi86)) {
                            case 0:
                                var v102 uint8 = 0
                                _ = v102
                            default:
                                var v103 float64 = 0.0
                                _ = v103
                            }
                            _ = fi96
                        }
                        var v104 int8 = int8(v26)
                        _ = v104
                    }
                }
            default:
                switch "hello" {
                case "":
                    switch (fi86 % (((fi86 + (-(fi86)))) | 1)) {
                    case 32768:
                        if true {
                            fmt.Println(v2)
                            fmt.Println(v2, v2, fi86)
                            var v105 int = (fi86 * -2147483648)
                            _ = v105
                        }
                    default:
                        v106 := GenS10{F0: int16(fi86), F1: "\n\t"}
                        _ = v106
                    }
                case "\n\t":
                    v11 -= v11
                }
            }
            _ = fi86
        }
    } else {
        if false {
            v11 = v11
            fmt.Println(v11, v11)
            switch "" {
            case "a":
                if true {
                    v2 += 973
                    var v107 int8 = -128
                    switch 2147483647 {
                    case 42:
                        _ = len(v3)
                    case 2147483647:
                        var v108 int8 = v107
                        _ = v108
                    case 150:
                        var v109 uint64 = 2
                        _ = v109
                    default:
                        if false {
                            _ = len(v1)
                            switch ("a" + ("hello" + "foo bar")) {
                            case "hello":
                                var v110 int64 = -1
                                _ = v110
                            default:
                                for fi111 := 0; fi111 < 18; fi111++ {
                                    v26 = v26
                                    _ = fi111
                                }
                            }
                            switch 606 {
                            case 2147483647:
                                var v112 int16 = -1
                                _ = v112
                            case -914:
                                fmt.Println(v11)
                            case -1:
                                v113 := GenS11{F0: false, F1: (v26 & 4294967295)}
                                _ = v113
                            }
                        }
                    }
                    _ = v107
                } else {
                    v114 := GenS12{F0: 42, F1: "", F2: 2147483647}
                    _ = v114
                }
            default:
                v115 := GenS13{F0: (-(v11)), F1: 892}
                _ = v115
            }
        }
        v116 := v3[v26]
        v117 := GenS3{F0: 2, F1: 18446744073709551615}
        _ = v116
        _ = v117
    }
    switch 128 {
    case -1:
        v1 = append(v1, 42)
    case 128:
        var v118 float32 = ((0.1 * 0.1) - 1e10)
        _ = v118
    case 4294967295:
        v1 = append(v1, 771)
    default:
        v119 := GenS9{F0: v2, F1: v26, F2: (v2 ^ -2147483648)}
        _ = v119
    }
    v2 -= int32(v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v3
}
