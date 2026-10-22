package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int16 }
type GenS2 struct { F0 uint16; F1 int16; F2 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint32; F1 int }
type GenS10 struct { F0 string; F1 uint16; F2 int8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int16; F1 uint64 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int16; F1 int64 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 int32; F1 bool; F2 int16 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 uint8 }
type GenS19 struct { F0 int; F1 uint64 }
type GenS20 struct { F0 uint8 }
type GenS21 struct { F0 uint16; F1 string; F2 uint8 }
type GenS22 struct { F0 int32; F1 bool }
type GenS23 struct { F0 int }
type GenS24 struct { F0 uint64 }
type GenS25 struct { F0 int32 }
type GenS26 struct { F0 string }
type GenS27 struct { F0 int; F1 uint64 }
type GenS28 struct { F0 uint64; F1 string; F2 int64 }
type GenS29 struct { F0 int64; F1 uint64 }
type GenS30 struct { F0 int64; F1 int64; F2 uint }
type GenS31 struct { F0 bool; F1 uint64; F2 uint }
type GenS32 struct { F0 int32 }
type GenS33 struct { F0 int64; F1 int16 }
func genFn0(p0 string, p1 uint32) int8 {
    v6 := GenS1{F0: true, F1: ("a" + ""), F2: 32767}
    for fi7 := 0; fi7 < 3; fi7++ {
        p1 = 65535
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int = -256
            p0 = ("a" + "\n\t")
            fmt.Println(v9, fi8, p0)
            _ = fi8
            _ = v9
        }
        v10 := []int{fi7, (fi7 - 416), -9223372036854775808}
        _ = fi7
        _ = v10
    }
    fmt.Println(p1)
    switch ((("" + p0) + ("" + p0)) + p0) {
    case "a":
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 uint32 = 1
            v13 := GenS2{F0: 65535, F1: 1000, F2: "a"}
            p1 -= (p1 | (v12 >> ((579) & 63)))
            _ = fi11
            _ = v12
            _ = v13
        }
    case "foo bar":
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 int16 = -1
            p0 = "a"
            _ = fi14
            _ = v15
        }
    case "":
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    _ = v6
    return 77
}

func main() {
    var v1 uint64 = 921
    var v2 int8 = 93
    var v3 float32 = 1.5
    var v4 uint8 = 128
    v5 := make(map[int]bool)
    var v16 bool = (750 == -760)
    if ((238 >= -9223372036854775808) && (v2 < (v2 + (v2 * 100)))) {
        v17 := GenS2{F0: 65535, F1: 0, F2: "a"}
        _ = v17
    } else {
        switch -697 {
        case 32767:
            v18 := GenS1{F0: true, F1: ("hello" + "a"), F2: -1}
            _ = v18
        default:
            delete(v5, 572)
        }
        v19 := []int{0, 32768, -2147483648}
        _ = v19
    }
    v1 -= (v1 - v1)
    var v20 uint16 = 255
    var v21 uint64 = 4294967295
    v22 := GenS3{F0: int(v4)}
    v23 := GenS1{F0: (1e10 <= 1e10), F1: "a", F2: 32767}
    if false {
        v24 := GenS1{F0: v16, F1: "hello", F2: 32767}
        _ = v24
    } else {
        v20 -= v20
        v25 := v5[1]
        switch -1 {
        case 810:
            v26 := GenS4{F0: 945, F1: "foo bar"}
            _ = v26
        case 150:
            v27 := GenS5{F0: 32767}
            _ = v27
        }
        _ = v25
    }
    switch -9223372036854775808 {
    case 4294967295:
        if ((v1 * (v1 | v21)) <= v1) {
            v5[1] = ((^(v4)) <= v4)
            fmt.Println(v21, v16)
        }
    }
    var v28 float64 = float64((1.5 * v3))
    v29 := GenS1{F0: (v4 < v4), F1: "hello", F2: -1}
    for fi30 := 0; fi30 < 7; fi30++ {
        var v31 int64 = 42
        _ = fi30
        _ = v31
    }
    v32 := []int{-286, -9223372036854775808, -104}
    if ((v21 / (((v21 << ((874) & 63))) | 1)) <= (v21 + (^(v21)))) {
        fmt.Println(v21, v4)
        if true {
            for fi33 := 0; fi33 < 3; fi33++ {
                for fi34 := 0; fi34 < 2; fi34++ {
                    if (!(53 < 1)) {
                        switch fi33 {
                        case 9223372036854775807:
                            v35 := GenS5{F0: 821}
                            _ = v35
                        case 127:
                            v36 := GenS6{F0: genFn0(("" + ("foo bar" + "a")), 42)}
                            _ = v36
                        case -81:
                            if false {
                                fmt.Println(v16)
                                v37 := GenS7{F0: 388, F1: 794}
                                switch (fi34 % ((fi33) | 1)) {
                                case 668:
                                    var v38 float32 = 0.1
                                    _ = v38
                                case 890:
                                    v28 -= ((-2.25 / v28) * v28)
                                default:
                                    if v16 {
                                        if len(v32) >= 2 {
                                            v32 = v32[1:len(v32)]
                                        }
                                    } else {
                                        v32 = append(v32, (fi33 & (fi34 + (fi33 * -138))))
                                        var v39 int32 = 42
                                        switch "hello" {
                                        case "foo bar":
                                            v40 := GenS5{F0: (fi34 & (fi34 + (fi33 | (^(fi33)))))}
                                            _ = v40
                                        case "a":
                                            fmt.Println(v4)
                                        default:
                                            v41 := v5[-431]
                                            _ = v41
                                        }
                                        if v16 {
                                            delete(v5, fi34)
                                            v28 = 0.1
                                            v42 := v5[1]
                                            _ = v42
                                        } else {
                                            fmt.Println(v20)
                                            v43 := GenS1{F0: false, F1: "foo bar", F2: -915}
                                            _ = v43
                                        }
                                        fmt.Println(v1, v2)
                                        _ = v39
                                    }
                                    if false {
                                        v16 = (v16 && v16)
                                        v44 := GenS8{F0: ("" + ("\n\t" + ("" + ("a" + ("a" + (("hello" + "foo bar") + ("foo bar" + "hello")))))))}
                                        if v16 {
                                            var v45 float64 = 0.0
                                            _ = v45
                                        } else {
                                            fmt.Println(fi33, v20, v4)
                                        }
                                        _ = v44
                                    } else {
                                        if false {
                                            v16 = v16
                                            delete(v5, fi33)
                                            v46 := GenS7{F0: v20, F1: 1}
                                            _ = v46
                                        }
                                    }
                                }
                                var v47 int = fi33
                                _ = v37
                                _ = v47
                            }
                            if v16 {
                                if (65535 == 2) {
                                    if (!false) {
                                        switch fi33 {
                                        case -1:
                                            switch ((("" + "hello") + "\n\t") + (("a" + "hello") + "hello")) {
                                            case "":
                                                var v48 uint64 = v1
                                                _ = v48
                                            case "\n\t":
                                                v28 = float64(-2.25)
                                            case "hello":
                                                if true {
                                                    v21 = v1
                                                    v32 = append(v32, -398)
                                                    switch ("" + "a") {
                                                    case "\n\t":
                                                        fmt.Println(fi34)
                                                    }
                                                    if v16 {
                                                        fmt.Println(v4, fi34, v21)
                                                        v16 = (v16 && true)
                                                    } else {
                                                        var v49 int64 = -1
                                                        v3 -= ((3.14159 / float32(v28)) - v3)
                                                        v4 -= 1
                                                        v50 := GenS5{F0: fi33}
                                                        v2 = (v2 + 1)
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                    fmt.Println(v4, v28)
                                                    if v16 {
                                                        if v16 {
                                                            var v51 int16 = 32767
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                v1 -= v21
                                                            case "":
                                                                fmt.Println(v3)
                                                            case "hello":
                                                                v32 = append(v32, (fi34 + fi33))
                                                                v28 -= (v28 / (0.1 + 1.5))
                                                                fmt.Println(fi33, v3, v28)
                                                                v2 -= (v2 / (((v2 << ((4294967295) & 63))) | 1))
                                                            }
                                                            v20 = v20
                                                            _ = v51
                                                        } else {
                                                            var v52 bool = v16
                                                            _ = v52
                                                        }
                                                    }
                                                    v20 = 65535
                                                }
                                                v28 = v28
                                                v28 += v28
                                            default:
                                                v53 := GenS8{F0: "hello"}
                                                _ = v53
                                            }
                                        case -2147483648:
                                            switch (-(fi33)) {
                                            case 1:
                                                v54 := GenS9{F0: 610, F1: fi34}
                                                _ = v54
                                            case 32767:
                                                fmt.Println(v2, v3, v1)
                                            case 2147483647:
                                                var v55 uint16 = (v20 | (^(v20)))
                                                _ = v55
                                            }
                                        default:
                                            v56 := v5[(fi33 << ((0) & 63))]
                                            _ = v56
                                        }
                                        switch (fi34 << ((4294967295) & 63)) {
                                        case 1:
                                            v2 += -128
                                        case 42:
                                            switch "hello" {
                                            case "hello":
                                                v16 = (((fi33 & 32767) <= (fi33 << ((42) & 63))) && v16)
                                            }
                                        }
                                    } else {
                                        if (!v16) {
                                            _ = len(v5)
                                            var v57 bool = (2147483647 < 32767)
                                            v58 := GenS9{F0: 316, F1: fi34}
                                            v59 := GenS9{F0: 2, F1: (fi33 & fi33)}
                                            _ = v57
                                            _ = v58
                                            _ = v59
                                        }
                                        v60 := v5[-86]
                                        _ = v60
                                    }
                                    fmt.Println(v20, v4)
                                } else {
                                    if (v1 > 628) {
                                        _ = len(v5)
                                        switch "a" {
                                        case "":
                                            switch fi33 {
                                            case 4294967295:
                                                v28 = (1.5 * (-2.25 + (v28 + (0.1 / v28))))
                                            }
                                        default:
                                            var v61 float64 = v28
                                            _ = v61
                                        }
                                        fmt.Println(v3, v4)
                                    } else {
                                        if v16 {
                                            var v62 bool = v16
                                            fmt.Println(v1, v21)
                                            var v63 int = (fi33 | fi34)
                                            _ = v62
                                            _ = v63
                                        }
                                        var v64 float32 = (1e10 * (-2.25 * v3))
                                        _ = v64
                                    }
                                    v20 = (v20 ^ 255)
                                    if v16 {
                                        v21 = 178
                                        v65 := GenS6{F0: 1}
                                        _ = v65
                                    } else {
                                        fmt.Println(v16, fi33)
                                    }
                                    switch (fi33 >> ((255) & 63)) {
                                    case -1:
                                        v66 := GenS10{F0: "hello", F1: v20, F2: (^(v2))}
                                        _ = v66
                                    }
                                }
                            }
                        }
                    } else {
                        if ((^(v21)) == 4294967295) {
                            var v67 int16 = 1
                            v32 = append(v32, fi34)
                            switch (fi33 & 1) {
                            case 576:
                                v68 := GenS9{F0: 2147483647, F1: (fi34 / ((fi34) | 1))}
                                _ = v68
                            case -4:
                                if ((32767 == v20) || v16) {
                                    var v69 bool = v16
                                    if true {
                                        v70 := v5[(^(fi33))]
                                        v71 := GenS1{F0: (v4 != v4), F1: "a", F2: int16(v1)}
                                        v72 := GenS11{F0: (v4 - v4)}
                                        var v73 int16 = v67
                                        fmt.Println(v28, v20)
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                        _ = v73
                                    } else {
                                        v69 = (629 == (^(v1)))
                                    }
                                    _ = v69
                                } else {
                                    fmt.Println(v20, v67, v16)
                                    if v16 {
                                        v32 = append(v32, -2147483648)
                                    }
                                }
                            }
                            _ = v67
                        }
                        v74 := GenS3{F0: fi33}
                        if v16 {
                            if true {
                                v5[fi33] = (4294967295 != 792)
                            }
                            switch (-(fi34)) {
                            case 4294967295:
                                var v75 uint16 = v20
                                var v76 uint64 = 4294967295
                                _ = v75
                                _ = v76
                            }
                            v16 = v16
                        } else {
                            switch "foo bar" {
                            case "hello":
                                var v77 int64 = 243
                                _ = v77
                            case "a":
                                if (!v16) {
                                    v78 := GenS11{F0: (v4 / ((1) | 1))}
                                    v28 = float64(((float32((((v28 + (0.1 / 3.14159)) / v28) - v28)) + 0.1) + (v3 / v3)))
                                    fmt.Println(v20, v28)
                                    _ = v78
                                } else {
                                    if v16 {
                                        v3 = v3
                                        if len(v32) >= 2 {
                                            v32 = v32[1:len(v32)]
                                        }
                                        switch "" {
                                        case "foo bar":
                                            v79 := GenS3{F0: 32767}
                                            _ = v79
                                        case "a":
                                            if len(v32) >= 2 {
                                                v32 = v32[1:len(v32)]
                                            }
                                        case "\n\t":
                                            var v80 uint = 403
                                            _ = v80
                                        }
                                    } else {
                                        switch fi33 {
                                        case 32767:
                                            _ = len(v32)
                                        default:
                                            if len(v32) > 0 {
                                                v32[0] = -2147483648
                                            }
                                            v3 -= ((float32(-2.25) + v3) / v3)
                                        }
                                    }
                                    if (!v16) {
                                        if v16 {
                                            if (!v16) {
                                                if (723 < 0) {
                                                    v5[fi33] = (734 != 32767)
                                                } else {
                                                    var v81 string = "hello"
                                                    if v16 {
                                                        v82 := GenS11{F0: v4}
                                                        _ = v82
                                                    }
                                                    if ((v3 / (((((v3 * v3) * (3.14159 * 0.1)) / 3.14159) / ((v3 / 3.14159) / v3)) * v3)) > ((((0.0 - ((0.1 / (float32((float64(((v3 * v3) - (v3 - ((v3 * (v3 - v3)) - v3)))) + 0.1)) + v3)) / 1.5)) * 1e10) * (float32(3.14159) + (((v3 * 0.0) - (v3 + ((-2.25 * v3) + v3))) * v3))) - v3)) {
                                                        if len(v32) > 0 {
                                                            v32[0] = fi33
                                                        }
                                                    }
                                                    _ = v81
                                                }
                                                v16 = v16
                                                v83 := GenS4{F0: 542, F1: ("hello" + ("\n\t" + (("foo bar" + "a") + "foo bar")))}
                                                _ = v83
                                            } else {
                                                v84 := GenS12{F0: 691, F1: (v1 & (v21 | v1))}
                                                _ = v84
                                            }
                                            v3 -= 1e10
                                            switch fi34 {
                                            case 32767:
                                                if ((!(v3 >= (v3 - (v3 / v3)))) || v16) {
                                                    v20 -= v20
                                                } else {
                                                    if v16 {
                                                        v32 = append(v32, (fi34 & 694))
                                                        v85 := GenS12{F0: 42, F1: v1}
                                                        if (!v16) {
                                                            v86 := GenS5{F0: fi33}
                                                            var v87 uint8 = (^(v4))
                                                            var v88 int8 = v2
                                                            _ = v86
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                        _ = v85
                                                    } else {
                                                        v20 = (^(v20))
                                                    }
                                                    if v16 {
                                                        v1 = (^(v21))
                                                    } else {
                                                        fmt.Println(fi34, v4, v20)
                                                        v2 += genFn0(((((("hello" + "foo bar") + "foo bar") + ("" + "hello")) + ("" + ("a" + "hello"))) + ""), 312)
                                                        v28 = 1.5
                                                    }
                                                    v3 += (v3 / (v3 - v3))
                                                }
                                            case 4294967295:
                                                v4 -= v4
                                            }
                                            v20 += (^(v20))
                                            v28 = (((v28 / -2.25) * v28) / ((0.0 - (v28 / v28)) * v28))
                                        } else {
                                            delete(v5, (fi33 << ((1) & 63)))
                                            switch fi33 {
                                            case 432:
                                                switch (fi34 % (((^(fi34))) | 1)) {
                                                case 4294967295:
                                                    var v89 int16 = -1
                                                    _ = v89
                                                case -909:
                                                    switch ("hello" + "hello") {
                                                    case "\n\t":
                                                        if (!(3.14159 < v28)) {
                                                            v20 += 32767
                                                            var v90 uint16 = v20
                                                            fmt.Println(v16, fi33, v1)
                                                            _ = v90
                                                        } else {
                                                            fmt.Println(fi34)
                                                            fmt.Println(v21, v2)
                                                            var v91 bool = false
                                                            _ = v91
                                                        }
                                                        var v92 float32 = -2.25
                                                        v28 += float64(((0.0 - v92) * v3))
                                                        _ = v92
                                                    default:
                                                        var v93 int = (fi34 / (((fi33 + fi34)) | 1))
                                                        _ = v93
                                                    }
                                                case 9223372036854775807:
                                                    v4 -= (v4 >> ((4294967295) & 63))
                                                default:
                                                    switch "" {
                                                    case "foo bar":
                                                        switch int(v2) {
                                                        case 32767:
                                                            fmt.Println(fi33, v21, v16)
                                                        case 597:
                                                            if v16 {
                                                                var v94 uint32 = 1
                                                                _ = v94
                                                            }
                                                            var v95 uint64 = (v1 ^ uint64(v4))
                                                            fmt.Println(v1, v95)
                                                            _ = v95
                                                        default:
                                                            v96 := GenS11{F0: v4}
                                                            var v97 bool = v16
                                                            _ = len(v5)
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                    default:
                                                        _ = len(v5)
                                                    }
                                                    var v98 int8 = 100
                                                    v99 := GenS13{F0: 65535}
                                                    _ = v98
                                                    _ = v99
                                                }
                                            default:
                                                switch 128 {
                                                case 32768:
                                                    v4 = 0
                                                case 32767:
                                                    var v100 int16 = -1
                                                    _ = v100
                                                case 4294967295:
                                                    v101 := GenS14{F0: -1, F1: 9223372036854775807}
                                                    _ = v101
                                                }
                                            }
                                        }
                                        switch "" {
                                        case "\n\t":
                                            if (false && true) {
                                                v102 := GenS2{F0: (v20 << ((633) & 63)), F1: 0, F2: ""}
                                                if (!(0 >= -9223372036854775808)) {
                                                    var v103 uint64 = v1
                                                    _ = v103
                                                }
                                                if v16 {
                                                    switch fi34 {
                                                    case 4294967295:
                                                        _ = len(v5)
                                                    case 32768:
                                                        switch ("\n\t" + "hello") {
                                                        case "hello":
                                                            fmt.Println(v3, v28, v28)
                                                        }
                                                    case 585:
                                                        v20 = v20
                                                    default:
                                                        if ("foo bar" != "hello") {
                                                            var v104 uint8 = (v4 - 2)
                                                            _ = v104
                                                        } else {
                                                            v105 := GenS2{F0: v20, F1: 1000, F2: ("hello" + "foo bar")}
                                                            v16 = (v16 || (true && (-9223372036854775808 != 1)))
                                                            _ = v105
                                                        }
                                                    }
                                                    switch (fi34 ^ fi33) {
                                                    case 1:
                                                        if v16 {
                                                            switch "hello" {
                                                            case "hello":
                                                                v106 := v5[fi34]
                                                                _ = v106
                                                            case "foo bar":
                                                                if (-803 > -1) {
                                                                    delete(v5, (-(fi34)))
                                                                }
                                                            }
                                                        }
                                                    case 42:
                                                        delete(v5, (fi33 | -606))
                                                    }
                                                    v107 := GenS7{F0: 7, F1: v21}
                                                    _ = v107
                                                }
                                                _ = v102
                                            } else {
                                                v108 := GenS8{F0: "foo bar"}
                                                if (!v16) {
                                                    var v109 string = ("hello" + (("foo bar" + "a") + "foo bar"))
                                                    var v110 int32 = 2147483647
                                                    _ = v109
                                                    _ = v110
                                                } else {
                                                    if false {
                                                        if false {
                                                            switch int(v4) {
                                                            case 2147483647:
                                                                var v111 float64 = -2.25
                                                                v112 := GenS15{F0: 65535}
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    fmt.Println(v3, v111)
                                                                default:
                                                                    if v16 {
                                                                        if len(v32) >= 2 {
                                                                            v32 = v32[1:len(v32)]
                                                                        }
                                                                    }
                                                                }
                                                                v32 = append(v32, fi33)
                                                                _ = v111
                                                                _ = v112
                                                            case 9223372036854775807:
                                                                v113 := GenS10{F0: "a", F1: 294, F2: -70}
                                                                _ = v113
                                                            }
                                                            v2 += (-(v2))
                                                        } else {
                                                            if false {
                                                                v1 -= uint64(v20)
                                                                switch "a" {
                                                                case "":
                                                                    var v114 uint8 = 35
                                                                    _ = v114
                                                                case "foo bar":
                                                                    var v115 uint16 = v20
                                                                    _ = v115
                                                                case "hello":
                                                                    if (0 != (v20 & v20)) {
                                                                        v116 := GenS4{F0: -145, F1: "foo bar"}
                                                                        _ = v116
                                                                    }
                                                                default:
                                                                    var v117 int32 = 1
                                                                    _ = v117
                                                                }
                                                                fmt.Println(v21, fi34)
                                                            }
                                                        }
                                                        if ((v20 - 960) >= (v20 % ((670) | 1))) {
                                                            var v118 int32 = -431
                                                            v119 := GenS1{F0: v16, F1: "", F2: -378}
                                                            v2 = v2
                                                            _ = v118
                                                            _ = v119
                                                        } else {
                                                            if (true || (("\n\t" + "hello") > "")) {
                                                                var v120 int32 = -1
                                                                v4 += v4
                                                                if v16 {
                                                                    fmt.Println(fi33, v20)
                                                                    if false {
                                                                        switch (fi34 + int(v2)) {
                                                                        case 793:
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                v1 = (v21 + (v1 - (v21 & v21)))
                                                                            case "":
                                                                                var v121 float64 = (0.0 * v28)
                                                                                _ = v121
                                                                            }
                                                                        case 32767:
                                                                            switch ("" + "foo bar") {
                                                                            case "hello":
                                                                                if (v4 >= 1) {
                                                                                    _ = len(v32)
                                                                                    v5[(fi33 % ((int(v2)) | 1))] = v16
                                                                                }
                                                                            case "a":
                                                                                _ = len(v5)
                                                                            }
                                                                        }
                                                                        fmt.Println(v4, v1, v3)
                                                                        v3 = ((float32(0.0) - (1.5 + ((v3 + -2.25) * v3))) * float32(v28))
                                                                    } else {
                                                                        v3 = v3
                                                                        var v122 uint64 = 2
                                                                        _ = v122
                                                                    }
                                                                } else {
                                                                    v123 := GenS10{F0: "a", F1: (v20 & 32767), F2: v2}
                                                                    v124 := GenS16{F0: (v120 % ((v120) | 1)), F1: v16, F2: -358}
                                                                    _ = v123
                                                                    _ = v124
                                                                }
                                                                switch fi33 {
                                                                case 0:
                                                                    var v125 float32 = 1.5
                                                                    _ = v125
                                                                case 500:
                                                                    v126 := v5[fi34]
                                                                    _ = v126
                                                                case -9223372036854775808:
                                                                    v127 := GenS5{F0: (^(fi34))}
                                                                    _ = v127
                                                                }
                                                                v2 = (v2 + int8(v21))
                                                                if (false && (127 < (v4 ^ 7))) {
                                                                    v128 := GenS7{F0: v20, F1: (v21 << ((1) & 63))}
                                                                    switch "a" {
                                                                    case "":
                                                                        if (uint16(v1) < v20) {
                                                                            v20 = v20
                                                                            v21 += 304
                                                                            delete(v5, int(v20))
                                                                            fmt.Println(v28)
                                                                            v20 -= 255
                                                                        }
                                                                    }
                                                                    if true {
                                                                        switch (fi34 - (^(fi34))) {
                                                                        case 32767:
                                                                            var v129 int = fi33
                                                                            v120 = 601
                                                                            _ = v129
                                                                        }
                                                                        var v130 uint16 = v20
                                                                        if (float32(v28) == v3) {
                                                                            if false {
                                                                                v21 = 42
                                                                            } else {
                                                                                v130 -= 1
                                                                                if (-1 == 1) {
                                                                                    switch fi34 {
                                                                                    case 9223372036854775807:
                                                                                        if (!v16) {
                                                                                            _ = len(v32)
                                                                                            if true {
                                                                                                v131 := GenS17{F0: (v2 | (v2 - v2))}
                                                                                                var v132 int64 = 1
                                                                                                _ = v131
                                                                                                _ = v132
                                                                                            }
                                                                                        }
                                                                                    case -2147483648:
                                                                                        v1 = 2
                                                                                    case 0:
                                                                                        if ((v16 || (((((!(float64(1.5) < float64(3.14159))) || (42 < -565)) || v16) && (!v16)) && false)) && false) {
                                                                                            v133 := v5[(fi34 % ((fi33) | 1))]
                                                                                            switch fi33 {
                                                                                            case -2147483648:
                                                                                                if (2147483647 > (^(v120))) {
                                                                                                    v134 := GenS18{F0: 165}
                                                                                                    fmt.Println(v21, v20, v28)
                                                                                                    v133 = true
                                                                                                    var v135 string = ("hello" + (("\n\t" + "hello") + "a"))
                                                                                                    _ = v134
                                                                                                    _ = v135
                                                                                                } else {
                                                                                                    if false {
                                                                                                        v136 := GenS19{F0: fi33, F1: 4294967295}
                                                                                                        switch (("hello" + "\n\t") + ("a" + (("" + "hello") + "a"))) {
                                                                                                        case "":
                                                                                                            if (!((v21 / ((v21) | 1)) >= (v1 * (v21 % ((v21) | 1))))) {
                                                                                                                if len(v32) >= 2 {
                                                                                                                    v32 = v32[1:len(v32)]
                                                                                                                }
                                                                                                                v120 += v120
                                                                                                            } else {
                                                                                                                switch ("a" + "hello") {
                                                                                                                case "foo bar":
                                                                                                                    v137 := GenS4{F0: -46, F1: "foo bar"}
                                                                                                                    _ = v137
                                                                                                                case "hello":
                                                                                                                    if (false && ((v20 ^ (v20 * v20)) < v20)) {
                                                                                                                        var v138 float64 = ((3.14159 * (0.1 + (v28 - ((float64(((float32(1.5) + v3) * v3)) - (v28 / v28)) - v28)))) - -2.25)
                                                                                                                        v21 = 2
                                                                                                                        _ = v138
                                                                                                                    } else {
                                                                                                                        if (v16 || false) {
                                                                                                                            switch fi33 {
                                                                                                                            case -323:
                                                                                                                                var v139 int = fi34
                                                                                                                                _ = v139
                                                                                                                            case 9223372036854775807:
                                                                                                                                fmt.Println(fi33, v28)
                                                                                                                            case 488:
                                                                                                                                var v140 string = "hello"
                                                                                                                                v3 = v3
                                                                                                                                v20 -= (v20 | 2)
                                                                                                                                _ = v140
                                                                                                                            }
                                                                                                                            switch "foo bar" {
                                                                                                                            case "\n\t":
                                                                                                                                var v141 int64 = -9223372036854775808
                                                                                                                                _ = v141
                                                                                                                            case "":
                                                                                                                                v133 = (!false)
                                                                                                                            default:
                                                                                                                                fmt.Println(v21, v120)
                                                                                                                            }
                                                                                                                            if v133 {
                                                                                                                                if v16 {
                                                                                                                                    switch ("" + "foo bar") {
                                                                                                                                    case "\n\t":
                                                                                                                                        if v16 {
                                                                                                                                            fmt.Println(v28, v20)
                                                                                                                                            v142 := GenS7{F0: 42, F1: v21}
                                                                                                                                            v143 := GenS5{F0: fi33}
                                                                                                                                            var v144 int64 = 1
                                                                                                                                            if len(v32) >= 2 {
                                                                                                                                                v32 = v32[1:len(v32)]
                                                                                                                                            }
                                                                                                                                            var v145 uint16 = uint16(v4)
                                                                                                                                            _ = v142
                                                                                                                                            _ = v143
                                                                                                                                            _ = v144
                                                                                                                                            _ = v145
                                                                                                                                        } else {
                                                                                                                                            v146 := GenS11{F0: (^(v4))}
                                                                                                                                            _ = v146
                                                                                                                                        }
                                                                                                                                    case "foo bar":
                                                                                                                                        var v147 uint8 = v4
                                                                                                                                        _ = v147
                                                                                                                                    }
                                                                                                                                    switch ((("hello" + "foo bar") + ("\n\t" + "\n\t")) + "\n\t") {
                                                                                                                                    case "":
                                                                                                                                        v5[929] = ((v2 | v2) >= int8(v130))
                                                                                                                                    case "a":
                                                                                                                                        var v148 uint64 = 0
                                                                                                                                        _ = v148
                                                                                                                                    case "hello":
                                                                                                                                        _ = len(v5)
                                                                                                                                    default:
                                                                                                                                        v20 -= v130
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    var v149 uint = uint(v21)
                                                                                                                                    if len(v32) > 0 {
                                                                                                                                        v32[0] = int(v4)
                                                                                                                                    }
                                                                                                                                    v150 := GenS20{F0: v4}
                                                                                                                                    _ = v149
                                                                                                                                    _ = v150
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                switch "foo bar" {
                                                                                                                                case "a":
                                                                                                                                    var v151 uint32 = 4294967295
                                                                                                                                    _ = v151
                                                                                                                                default:
                                                                                                                                    if len(v32) > 0 {
                                                                                                                                        v32[0] = 389
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                if true {
                                                                                                                                    switch "hello" {
                                                                                                                                    case "foo bar":
                                                                                                                                        if (!(!((fi34 - (fi33 * (^(fi33)))) <= fi34))) {
                                                                                                                                            if (18446744073709551615 == 42) {
                                                                                                                                                switch 0 {
                                                                                                                                                case 9223372036854775807:
                                                                                                                                                    v32 = append(v32, (fi34 | -541))
                                                                                                                                                case 2147483647:
                                                                                                                                                    switch fi33 {
                                                                                                                                                    case 25:
                                                                                                                                                        var v152 uint = 181
                                                                                                                                                        _ = v152
                                                                                                                                                    }
                                                                                                                                                    v153 := v5[fi34]
                                                                                                                                                    _ = v153
                                                                                                                                                default:
                                                                                                                                                    fmt.Println(v2, v3)
                                                                                                                                                }
                                                                                                                                                if ("a" >= "") {
                                                                                                                                                    v16 = (!false)
                                                                                                                                                    if v133 {
                                                                                                                                                        var v154 int8 = (^(v2))
                                                                                                                                                        _ = v154
                                                                                                                                                    } else {
                                                                                                                                                        var v155 bool = (2 > 1)
                                                                                                                                                        _ = v155
                                                                                                                                                    }
                                                                                                                                                    if len(v32) >= 2 {
                                                                                                                                                        v32 = v32[1:len(v32)]
                                                                                                                                                    }
                                                                                                                                                    var v156 string = (("hello" + "hello") + "a")
                                                                                                                                                    var v157 int64 = 946
                                                                                                                                                    _ = v156
                                                                                                                                                    _ = v157
                                                                                                                                                }
                                                                                                                                                if (!(uint(v21) < 647)) {
                                                                                                                                                    if false {
                                                                                                                                                        v158 := GenS18{F0: 18}
                                                                                                                                                        _ = len(v5)
                                                                                                                                                        _ = v158
                                                                                                                                                    } else {
                                                                                                                                                        v32 = append(v32, fi34)
                                                                                                                                                    }
                                                                                                                                                    if v133 {
                                                                                                                                                        v133 = (int16(v21) > 42)
                                                                                                                                                        v32 = append(v32, 9223372036854775807)
                                                                                                                                                        v159 := GenS21{F0: v130, F1: ("a" + "\n\t"), F2: 18}
                                                                                                                                                        v120 = -2147483648
                                                                                                                                                        _ = v159
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    case "\n\t":
                                                                                                                                        var v160 int8 = v2
                                                                                                                                        _ = v160
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v3 = 3.14159
                                                                                                                        var v161 int64 = 9223372036854775807
                                                                                                                        _ = v161
                                                                                                                    }
                                                                                                                }
                                                                                                                if (!v16) {
                                                                                                                    var v162 int64 = 9223372036854775807
                                                                                                                    var v163 int = (fi33 ^ fi34)
                                                                                                                    var v164 int = (fi33 / (((fi33 | 370)) | 1))
                                                                                                                    _ = v162
                                                                                                                    _ = v163
                                                                                                                    _ = v164
                                                                                                                } else {
                                                                                                                    v32 = append(v32, fi33)
                                                                                                                    switch ((("hello" + "\n\t") + "a") + ("foo bar" + "")) {
                                                                                                                    case "":
                                                                                                                        v165 := GenS22{F0: (v120 >> ((65535) & 63)), F1: (v3 != v3)}
                                                                                                                        _ = v165
                                                                                                                    }
                                                                                                                    v166 := GenS23{F0: fi33}
                                                                                                                    fmt.Println(v3, v2, v130)
                                                                                                                    _ = v166
                                                                                                                }
                                                                                                                v20 += (v130 * v130)
                                                                                                            }
                                                                                                        default:
                                                                                                            fmt.Println(fi34, v28, v4)
                                                                                                            if (uint64(v130) < v1) {
                                                                                                                var v167 uint64 = (v21 ^ v1)
                                                                                                                _ = v167
                                                                                                            } else {
                                                                                                                _ = len(v32)
                                                                                                            }
                                                                                                            v21 = 2
                                                                                                            v168 := GenS24{F0: (v21 / (((v21 ^ v21)) | 1))}
                                                                                                            v169 := GenS13{F0: uint32(v130)}
                                                                                                            var v170 float32 = 0.0
                                                                                                            _ = v168
                                                                                                            _ = v169
                                                                                                            _ = v170
                                                                                                        }
                                                                                                        var v171 uint32 = 29
                                                                                                        _ = v136
                                                                                                        _ = v171
                                                                                                    } else {
                                                                                                        if len(v32) > 0 {
                                                                                                            v32[0] = fi34
                                                                                                        }
                                                                                                        v2 -= (v2 ^ v2)
                                                                                                        v133 = v133
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                if len(v32) > 0 {
                                                                                                    v32[0] = -2147483648
                                                                                                }
                                                                                            }
                                                                                            var v172 int16 = 200
                                                                                            _ = v133
                                                                                            _ = v172
                                                                                        }
                                                                                        if v16 {
                                                                                            v5[32768] = false
                                                                                            var v173 float32 = 1.5
                                                                                            v2 = v2
                                                                                            _ = v173
                                                                                        }
                                                                                        v174 := GenS18{F0: 255}
                                                                                        _ = v174
                                                                                    default:
                                                                                        var v175 string = "\n\t"
                                                                                        _ = v175
                                                                                    }
                                                                                } else {
                                                                                    switch "\n\t" {
                                                                                    case "foo bar":
                                                                                        _ = len(v5)
                                                                                    case "hello":
                                                                                        if v16 {
                                                                                            v176 := GenS18{F0: (^(v4))}
                                                                                            var v177 float32 = (float32(v28) + ((float32(v28) + 1e10) * v3))
                                                                                            _ = v176
                                                                                            _ = v177
                                                                                        }
                                                                                        fmt.Println(v1, fi34)
                                                                                    case "\n\t":
                                                                                        v5[-9223372036854775808] = (2147483648 >= 0)
                                                                                    default:
                                                                                        v178 := GenS25{F0: 42}
                                                                                        _ = v178
                                                                                    }
                                                                                }
                                                                                _ = len(v32)
                                                                            }
                                                                            if v16 {
                                                                                if (!false) {
                                                                                    v32 = append(v32, 42)
                                                                                    _ = len(v5)
                                                                                    switch -9223372036854775808 {
                                                                                    case -425:
                                                                                        v21 += v1
                                                                                    case 32768:
                                                                                        v179 := GenS10{F0: "", F1: (^(v130)), F2: (v2 + 1)}
                                                                                        _ = v179
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v180 float32 = v3
                                                                                v28 = 1.5
                                                                                v16 = (v16 && (!(v130 > v130)))
                                                                                v5[(fi33 - (^(fi33)))] = ((!true) || (v20 != v130))
                                                                                _ = v180
                                                                            }
                                                                        } else {
                                                                            v20 = v20
                                                                            v181 := GenS26{F0: "foo bar"}
                                                                            v182 := GenS27{F0: 4294967295, F1: 2}
                                                                            v183 := GenS19{F0: fi33, F1: v1}
                                                                            _ = v181
                                                                            _ = v182
                                                                            _ = v183
                                                                        }
                                                                        _ = v130
                                                                    } else {
                                                                        if len(v32) >= 2 {
                                                                            v32 = v32[1:len(v32)]
                                                                        }
                                                                        v1 += uint64(v4)
                                                                        if true {
                                                                            v21 = v1
                                                                            v1 = v21
                                                                        } else {
                                                                            v184 := GenS1{F0: (!(279 == 731)), F1: "foo bar", F2: 0}
                                                                            v2 = (v2 * -1)
                                                                            var v185 int64 = 153
                                                                            _ = v184
                                                                            _ = v185
                                                                        }
                                                                        v28 = (v28 / ((v28 * v28) - v28))
                                                                    }
                                                                    var v186 string = "\n\t"
                                                                    var v187 uint32 = 131
                                                                    _ = v128
                                                                    _ = v186
                                                                    _ = v187
                                                                } else {
                                                                    v188 := GenS28{F0: v21, F1: "", F2: -9223372036854775808}
                                                                    var v189 int8 = genFn0(("foo bar" + "\n\t"), 2)
                                                                    _ = v188
                                                                    _ = v189
                                                                }
                                                                _ = len(v5)
                                                                v190 := GenS13{F0: 65535}
                                                                var v191 int8 = genFn0("hello", uint32(v20))
                                                                _ = v120
                                                                _ = v190
                                                                _ = v191
                                                            }
                                                            v192 := GenS17{F0: v2}
                                                            _ = len(v32)
                                                            _ = v192
                                                        }
                                                    }
                                                    var v193 int = (-(fi33))
                                                    v193 += v193
                                                    _ = v193
                                                }
                                                _ = v108
                                            }
                                        case "a":
                                            if (-1 != 1000) {
                                                if (-459 < 9223372036854775807) {
                                                    v3 = v3
                                                } else {
                                                    v194 := GenS18{F0: 128}
                                                    v16 = true
                                                    switch "a" {
                                                    case "foo bar":
                                                        fmt.Println(v2)
                                                    case "":
                                                        v16 = v16
                                                    case "\n\t":
                                                        fmt.Println(v16)
                                                    }
                                                    _ = v194
                                                }
                                            }
                                        default:
                                            if len(v32) >= 2 {
                                                v32 = v32[1:len(v32)]
                                            }
                                        }
                                        _ = len(v32)
                                    } else {
                                        v3 += v3
                                        v1 = 18446744073709551615
                                    }
                                    if (-1 <= -1) {
                                        v32 = append(v32, (-(fi33)))
                                        var v195 int = 536
                                        var v196 uint = 0
                                        _ = v195
                                        _ = v196
                                    } else {
                                        if v16 {
                                            v197 := GenS29{F0: 9223372036854775807, F1: v1}
                                            _ = v197
                                        } else {
                                            _ = len(v5)
                                        }
                                        v21 -= (v1 >> ((2) & 63))
                                    }
                                }
                            }
                        }
                        v198 := GenS30{F0: 1, F1: 42, F2: 42}
                        _ = v198
                        _ = v74
                    }
                    _ = fi34
                }
                v199 := GenS31{F0: ((v1 / ((259) | 1)) > v1), F1: (v1 + 2), F2: uint(v4)}
                switch fi33 {
                case 32767:
                    v28 -= 1e10
                default:
                    v200 := GenS32{F0: 2147483647}
                    _ = v200
                }
                _ = fi33
                _ = v199
            }
        }
        if (v28 >= v28) {
            var v201 int = 0
            _ = v201
        }
    }
    fmt.Println(v21, v21)
    v202 := GenS33{F0: -9223372036854775808, F1: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v202
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
