package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 bool; F1 int8; F2 uint64 }
type GenS8 struct { F0 int16; F1 uint16; F2 int32 }
type GenS9 struct { F0 string }
func genFn0(p0 int16, p1 string, p2 bool) int {
    p1 = (("foo bar" + p1) + p1)
    p2 = (128 > 128)
    if (0 <= 0) {
        for fi6 := 0; fi6 < 5; fi6++ {
            v7 := GenS1{F0: p0, F1: 0, F2: 1}
            for fi8 := 0; fi8 < 10; fi8++ {
                var v9 uint8 = 255
                var v10 int = fi6
                _ = fi8
                _ = v10
                _ = v9
            }
            switch p1 {
            case "foo bar":
                if p2 {
                    if (-2.25 >= 0.1) {
                        p1 = "hello"
                        if p2 {
                            switch (fi6 / ((42) | 1)) {
                            case 32768:
                                if (42 == 4294967295) {
                                    v11 := GenS1{F0: (p0 * (p0 + p0)), F1: 1, F2: 0}
                                    _ = v11
                                } else {
                                    v12 := GenS1{F0: (p0 % ((1000) | 1)), F1: 920, F2: 0}
                                    _ = v12
                                }
                            case -9223372036854775808:
                                var v13 uint8 = 1
                                _ = v13
                            }
                            var v14 int8 = 127
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 9; fi15++ {
                                switch 2147483647 {
                                case 0:
                                    switch fi6 {
                                    case 32768:
                                        fmt.Println(fi6, fi15)
                                    case 310:
                                        v16 := GenS2{F0: fi15}
                                        _ = v16
                                    case -577:
                                        p2 = ((^(fi6)) > (fi15 / ((787) | 1)))
                                    }
                                    v17 := []int{(fi15 / ((-2147483648) | 1)), (fi6 >> ((2) & 63)), (fi15 - fi15), 127}
                                    _ = v17
                                }
                                v18 := GenS3{F0: 128}
                                _ = fi15
                                _ = v18
                            }
                            var v19 string = p1
                            if (0.1 <= (-2.25 - 1.5)) {
                                if (!(65535 < 510)) {
                                    p1 = ("\n\t" + v19)
                                    p2 = (!(!p2))
                                } else {
                                    v20 := GenS1{F0: -462, F1: 2147483647, F2: 0}
                                    _ = v20
                                }
                                if (1e10 < float64(-2.25)) {
                                    if (p2 || (((fi6 - (fi6 % (((fi6 * (fi6 << ((255) & 63)))) | 1))) <= 0) || true)) {
                                        var v21 int8 = -30
                                        if p2 {
                                            p0 += (^(p0))
                                            for fi22 := 0; fi22 < 16; fi22++ {
                                                p0 += (p0 * (p0 - 1))
                                                var v23 int16 = p0
                                                _ = fi22
                                                _ = v23
                                            }
                                            v19 = "hello"
                                        }
                                        v24 := []int{fi6}
                                        _ = v21
                                        _ = v24
                                    }
                                } else {
                                    if (65535 > 2) {
                                        v25 := []int{(fi6 & -1), fi6}
                                        p0 = (p0 | p0)
                                        _ = v25
                                    } else {
                                        switch "\n\t" {
                                        case "":
                                            fmt.Println(p2)
                                        case "hello":
                                            p1 = (v19 + "")
                                        }
                                        var v26 uint32 = 2147483647
                                        v27 := []int{fi6, 1, (^(fi6))}
                                        _ = v26
                                        _ = v27
                                    }
                                }
                                v19 = (v19 + p1)
                            } else {
                                v28 := []int{}
                                _ = v28
                            }
                            _ = v19
                        }
                        v29 := []int{}
                        _ = v29
                    }
                    var v30 float32 = (0.1 * ((float32(1.5) / 1.0) * float32(0.0)))
                    _ = v30
                }
            case "a":
                v31 := []int{fi6}
                _ = v31
            case "":
                switch 0 {
                case 32768:
                    for fi32 := 0; fi32 < 15; fi32++ {
                        if (65535 != 550) {
                            p2 = (!p2)
                            var v33 int16 = -32768
                            _ = v33
                        }
                        _ = fi32
                    }
                case 0:
                    switch "a" {
                    case "a":
                        v34 := GenS4{F0: -603}
                        _ = v34
                    }
                }
            }
            _ = fi6
            _ = v7
        }
        v35 := []int{-2147483648, 0}
        var v36 string = p1
        _ = v35
        _ = v36
    }
    if (904 == 2) {
        var v37 uint32 = 591
        v37 = (v37 % ((65535) | 1))
        if p2 {
            var v38 int32 = 0
            _ = v38
        }
        _ = v37
    }
    var v39 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v39
    return 482
}

func main() {
    v1 := []string{(("hello" + "foo bar") + (("\n\t" + "hello") + "")), "", "hello"}
    var v2 int32 = 2147483647
    var v3 int32 = v2
    v4 := make(map[int]bool)
    v5 := make(map[int]string)
    for fi40 := 0; fi40 < 4; fi40++ {
        var v41 float64 = (0.0 + 1.5)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v42 bool = true
        _ = fi40
        _ = v41
        _ = v42
    }
    if true {
        if true {
            var v43 uint64 = 546
            var v44 int64 = 890
            _ = v43
            _ = v44
        }
        if false {
            switch 32768 {
            case 32767:
                v2 = (v3 % (((-(v2))) | 1))
            case -9223372036854775808:
                var v45 uint64 = 2
                _ = v45
            default:
                v46 := GenS2{F0: genFn0(32767, "hello", true)}
                _ = v46
            }
            v3 -= (v3 << ((18446744073709551615) & 63))
            for fi47 := 0; fi47 < 14; fi47++ {
                for fi48 := 0; fi48 < 16; fi48++ {
                    switch "hello" {
                    case "":
                        var v49 uint64 = 809
                        v3 = v2
                        var v50 uint64 = (v49 | v49)
                        v1 = append(v1, "hello")
                        if true {
                            var v51 uint32 = 674
                            v3 = (v2 << ((32) & 63))
                            _ = v51
                        } else {
                            var v52 uint16 = 964
                            v53 := GenS5{F0: "foo bar"}
                            var v54 int32 = v2
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                        _ = v49
                        _ = v50
                    default:
                        if (42 <= -452) {
                            v2 = 306
                            var v55 bool = false
                            _ = v55
                        }
                    }
                    _ = fi48
                }
                for fi56 := 0; fi56 < 9; fi56++ {
                    v57 := GenS6{F0: 127, F1: 530}
                    switch "a" {
                    case "a":
                        var v58 uint16 = 65535
                        _ = v58
                    case "":
                        var v59 float64 = 0.0
                        _ = v59
                    default:
                        if (float32(3.14159) == 1.5) {
                            switch (fi56 - fi47) {
                            case 32767:
                                v2 = 1
                                if len(v1) > 0 {
                                    v1[0] = "a"
                                }
                                switch genFn0(32767, "", (!((fi47 | fi47) <= genFn0(-695, "hello", (255 <= 1))))) {
                                case 4294967295:
                                    if (true && (!false)) {
                                        v60 := GenS7{F0: ((42 <= -248) || false), F1: -54, F2: 4294967295}
                                        _ = v60
                                    } else {
                                        var v61 float64 = ((float64(((((0.0 / 1.0) - 3.14159) - (-2.25 + (float32(0.1) * (3.14159 * 1e10)))) / 1.5)) / ((3.14159 / 1e10) / (1.5 / 0.1))) / -2.25)
                                        _ = v61
                                    }
                                case -144:
                                    if (true && (!(255 >= 55))) {
                                        if (true || true) {
                                            if ((!((^(v3)) == v3)) && (!(!(!((((float64(3.14159) / (-2.25 * 1.5)) + -2.25) / (1e10 + float64(0.0))) >= 3.14159))))) {
                                                fmt.Println(fi56, fi47, fi47)
                                                v62 := GenS3{F0: 0}
                                                fmt.Println(fi56, fi47)
                                                _ = v62
                                            } else {
                                                var v63 uint = 42
                                                v63 = (v63 * v63)
                                                switch fi56 {
                                                case 42:
                                                    switch ("" + "") {
                                                    case "\n\t":
                                                        if (-128 == int8(v2)) {
                                                            var v64 int16 = 1
                                                            var v65 int32 = -2147483648
                                                            var v66 bool = ((fi47 + (^(fi56))) != (fi56 << ((1) & 63)))
                                                            _ = v64
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                    case "a":
                                                        fmt.Println(fi47, fi56, v63)
                                                    }
                                                case 451:
                                                    fmt.Println(v2, v63)
                                                case 128:
                                                    v3 = v2
                                                }
                                                _ = v63
                                            }
                                        }
                                    }
                                case 315:
                                    switch ("\n\t" + "\n\t") {
                                    case "hello":
                                        if ((((!false) || ("" != "hello")) || true) || (128 <= fi56)) {
                                            switch 2147483647 {
                                            case 0:
                                                if (true && false) {
                                                    _ = len(v1)
                                                    v3 = (v2 >> ((1) & 63))
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            case 127:
                                                _ = len(v1)
                                                v67 := GenS7{F0: (false || ((0 != 114) || ((!(!false)) && (351 != 1)))), F1: 1, F2: 2}
                                                fmt.Println(fi47)
                                                _ = v67
                                            case -9223372036854775808:
                                                v68 := GenS8{F0: 0, F1: 65535, F2: (v3 % (((v3 - (^(v2)))) | 1))}
                                                _ = v68
                                            }
                                        }
                                    case "a":
                                        v2 += 1
                                    default:
                                        var v69 float32 = 0.1
                                        _ = v69
                                    }
                                }
                            case -9223372036854775808:
                                fmt.Println(v3)
                            default:
                                v2 += (v3 | -412)
                            }
                            if false {
                                if (true && false) {
                                    delete(v5, (fi47 ^ (fi47 >> ((4294967295) & 63))))
                                    if true {
                                        v3 -= v3
                                    } else {
                                        var v70 uint = 0
                                        delete(v5, (^(fi47)))
                                        _ = v70
                                    }
                                    v71 := GenS8{F0: -32768, F1: 138, F2: v3}
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v72 uint32 = 2147483647
                                        v2 = v2
                                        _ = v72
                                    case "a":
                                        if ((3.14159 / float32(1e10)) <= (-2.25 / 1e10)) {
                                            v3 += (-(v2))
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if false {
                                                v2 = -2147483648
                                                fmt.Println(fi56, fi56, fi56)
                                            }
                                        }
                                    default:
                                        var v73 int = genFn0(1000, "foo bar", false)
                                        _ = v73
                                    }
                                    _ = v71
                                }
                                var v74 int16 = 1
                                _ = v74
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if len(v1) > 0 {
                                    v1[0] = ("\n\t" + "\n\t")
                                }
                                v75 := GenS7{F0: true, F1: int8(v2), F2: 24}
                                _ = v75
                            }
                            v2 += -1
                        } else {
                            var v76 int8 = 42
                            v4[fi47] = (!false)
                            fmt.Println(v76, fi47, fi47)
                            _ = v76
                        }
                        v2 = (v2 >> ((960) & 63))
                    }
                    v77 := v5[fi56]
                    _ = fi56
                    _ = v57
                    _ = v77
                }
                v3 += (v2 + v2)
                _ = fi47
            }
        }
    }
    if (2 < 255) {
        v2 -= -168
        if true {
            v2 = v3
            if false {
                var v78 uint32 = 2
                _ = v78
            }
        }
    } else {
        for fi79 := 0; fi79 < 5; fi79++ {
            switch "a" {
            case "foo bar":
                var v80 float32 = 1e10
                _ = v80
            }
            _ = fi79
        }
        v2 = (v2 + v3)
        fmt.Println(v2, v3)
    }
    switch genFn0(-508, "", (!false)) {
    case 868:
        fmt.Println(v3)
    }
    v2 = (v3 / ((v3) | 1))
    if (!(-32768 == -1)) {
        v81 := GenS4{F0: 1}
        v82 := GenS8{F0: 77, F1: 741, F2: v3}
        _ = v81
        _ = v82
    } else {
        v3 -= (v3 ^ (v2 | -2147483648))
        v83 := GenS2{F0: genFn0(1000, (("foo bar" + "\n\t") + "foo bar"), ((!true) && (!((!true) || (!false)))))}
        _ = v83
    }
    for fi84 := 0; fi84 < 10; fi84++ {
        var v85 uint8 = 42
        for fi86 := 0; fi86 < 9; fi86++ {
            var v87 int32 = int32(v85)
            var v88 uint32 = 320
            v88 = v88
            _ = fi86
            _ = v87
            _ = v88
        }
        var v89 uint16 = 788
        _ = fi84
        _ = v85
        _ = v89
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    if (!(("" != "a") && (-756 <= -32768))) {
        v90 := GenS3{F0: 194}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v90
    }
    if true {
        var v91 int64 = 354
        v92 := GenS6{F0: -128, F1: 219}
        _ = v91
        _ = v92
    }
    v3 -= 9
    v2 -= v3
    if (!true) {
        var v93 uint8 = 128
        v3 -= v3
        switch "foo bar" {
        case "a":
            fmt.Println(v3)
        default:
            switch -664 {
            case 9223372036854775807:
                v2 += v3
            case 498:
                for fi94 := 0; fi94 < 16; fi94++ {
                    v93 = (^(v93))
                    switch ("a" + "foo bar") {
                    case "\n\t":
                        switch "a" {
                        case "\n\t":
                            if false {
                                v5[genFn0(1000, "\n\t", true)] = "foo bar"
                                var v95 uint32 = 4294967295
                                v96 := GenS1{F0: -32768, F1: (-(v3)), F2: 20}
                                _ = v95
                                _ = v96
                            }
                        case "":
                            v2 = (v3 % ((282) | 1))
                        case "hello":
                            if false {
                                v3 -= (v2 << ((1) & 63))
                                fmt.Println(v93)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        }
                    case "foo bar":
                        var v97 uint8 = v93
                        _ = v97
                    default:
                        var v98 uint8 = v93
                        _ = v98
                    }
                    _ = fi94
                }
            case -9223372036854775808:
                switch "\n\t" {
                case "a":
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                case "foo bar":
                    v2 += v2
                case "":
                    var v99 int = genFn0(0, ("foo bar" + ""), ((!(2 >= 18446744073709551615)) || (42 < 255)))
                    _ = v99
                default:
                    fmt.Println(v3, v2, v3)
                }
            }
        }
        _ = v93
    }
    if true {
        v100 := GenS8{F0: -32768, F1: 0, F2: (v2 + 132)}
        switch "\n\t" {
        case "a":
            for fi101 := 0; fi101 < 15; fi101++ {
                var v102 int32 = v3
                fmt.Println(v2)
                var v103 float32 = (1.5 / 1.0)
                _ = fi101
                _ = v102
                _ = v103
            }
        case "hello":
            v104 := GenS8{F0: -1, F1: 32767, F2: -620}
            _ = v104
        }
        var v105 int16 = 1000
        _ = v100
        _ = v105
    } else {
        for fi106 := 0; fi106 < 11; fi106++ {
            delete(v5, fi106)
            _ = fi106
        }
        v107 := GenS9{F0: "a"}
        _ = v107
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
