package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 bool; F1 uint64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int32; F1 int8 }
type GenS10 struct { F0 string; F1 string }
type GenS11 struct { F0 int8; F1 uint64; F2 bool }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 string; F1 uint64; F2 int64 }
type GenS14 struct { F0 uint8; F1 int16 }
type GenS15 struct { F0 int; F1 int8; F2 int }
type GenS16 struct { F0 int; F1 string }
type GenS17 struct { F0 int64; F1 int32 }
type GenS18 struct { F0 uint8; F1 uint }
type GenS19 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS20 struct { F0 int }
func genFn0(p0 uint) bool {
    p0 -= p0
    var v4 uint8 = 2
    p0 -= 4294967295
    _ = p0
    _ = v4
    return false
}
func genFn1(p0 uint, p1 int32, p2 uint8) int16 {
    fmt.Println(p1, p1)
    fmt.Println(p0, p2, p1)
    p2 = 1
    p1 += 42
    switch 2147483647 {
    case -264:
        if false {
            var v5 string = ""
            _ = v5
        } else {
            v6 := []int{2147483647}
            var v7 bool = (p2 != (p2 * p2))
            _ = v6
            _ = v7
        }
    default:
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := []int{fi8}
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 int32 = (p1 << ((972) & 63))
                p1 = (v11 ^ -2147483648)
                if true {
                    var v12 uint32 = 0
                    _ = v12
                } else {
                    var v13 float64 = (0.1 * (1.5 - 0.1))
                    fmt.Println(fi8, p0, fi8)
                    _ = v13
                }
                p2 -= 42
                fmt.Println(fi10)
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 110
}

func main() {
    var v1 int32 = 1
    var v2 uint64 = 2
    var v3 uint = 923
    v14 := []int{-216, 127, 128}
    var v15 float32 = 0.0
    var v16 uint32 = 0
    v2 = 0
    for fi17 := 0; fi17 < 20; fi17++ {
        v15 = 1.5
        var v18 uint8 = 1
        v3 -= v3
        _ = fi17
        _ = v18
    }
    v1 += (-(v1))
    v19 := GenS1{F0: 153, F1: v1}
    for fi20 := 0; fi20 < 19; fi20++ {
        if (2 < (^(v2))) {
            if true {
                if true {
                    switch ("foo bar" + "foo bar") {
                    case "foo bar":
                        if len(v14) > 0 {
                            v14[0] = (fi20 >> ((uint(v2)) & 63))
                        }
                    case "a":
                        v21 := GenS1{F0: (^(fi20)), F1: -2147483648}
                        _ = v21
                    case "\n\t":
                        v22 := GenS2{F0: 83, F1: true}
                        _ = v22
                    default:
                        v23 := GenS3{F0: v3}
                        _ = v23
                    }
                    var v24 bool = (v2 <= uint64(v16))
                    v14 = append(v14, 179)
                    _ = v24
                }
                var v25 uint32 = v16
                _ = v25
            } else {
                if len(v14) > 0 {
                    v14[0] = -404
                }
                if (!(!(false && (0 == 497)))) {
                    switch ("hello" + ("\n\t" + "")) {
                    case "a":
                        for fi26 := 0; fi26 < 8; fi26++ {
                            fmt.Println(v16)
                            v1 += -2147483648
                            if ("hello" <= "\n\t") {
                                v1 -= v1
                            }
                            _ = fi26
                        }
                    case "":
                        if false {
                            v2 -= v2
                        }
                    }
                }
            }
            v2 = (v2 + v2)
        } else {
            fmt.Println(v1)
            v27 := GenS2{F0: 244, F1: true}
            _ = v27
        }
        var v28 uint32 = (v16 * v16)
        var v29 uint8 = 2
        _ = fi20
        _ = v28
        _ = v29
    }
    if true {
        if (!false) {
            for fi30 := 0; fi30 < 9; fi30++ {
                v1 -= -1
                for fi31 := 0; fi31 < 16; fi31++ {
                    v16 -= v16
                    var v32 string = "\n\t"
                    fmt.Println(v16)
                    if ((3.14159 - (1e10 / (0.1 + float64(-2.25)))) < (float64(0.0) + -2.25)) {
                        switch (^(fi31)) {
                        case -9223372036854775808:
                            fmt.Println(v3)
                            var v33 int16 = -32768
                            _ = v33
                        case 32768:
                            switch v32 {
                            case "\n\t":
                                if false {
                                    v1 += v1
                                }
                            }
                            fmt.Println(fi30)
                            if (true || (v2 <= 42)) {
                                fmt.Println(v2, v16)
                            } else {
                                fmt.Println(v3, fi31)
                                if true {
                                    v15 = float32((1.5 + 0.0))
                                } else {
                                    v34 := GenS3{F0: v3}
                                    v15 += float32(0.1)
                                    _ = v34
                                }
                                v35 := GenS4{F0: false, F1: 98}
                                _ = v35
                            }
                            v1 += (v1 % ((v1) | 1))
                            var v36 int = int(v16)
                            _ = v36
                        default:
                            var v37 float32 = v15
                            _ = v37
                        }
                    } else {
                        fmt.Println(v16, fi31, v16)
                        var v38 uint32 = v16
                        _ = v38
                    }
                    _ = fi31
                    _ = v32
                }
                fmt.Println(v3, v15, v2)
                _ = fi30
            }
            var v39 uint16 = 32767
            if false {
                if (false && (v15 != ((v15 - ((v15 * v15) * float32(float64(v15)))) - (v15 * 1.5)))) {
                    if len(v14) > 0 {
                        v14[0] = 127
                    }
                    for fi40 := 0; fi40 < 1; fi40++ {
                        fmt.Println(v2, v15)
                        switch "foo bar" {
                        case "":
                            v39 += (^(v39))
                        case "hello":
                            var v41 uint = v3
                            _ = v41
                        case "\n\t":
                            for fi42 := 0; fi42 < 1; fi42++ {
                                if ((false && ((v39 | (^(v39))) <= 1)) && false) {
                                    var v43 int = (-(fi40))
                                    if (!true) {
                                        if false {
                                            if (true && (0.1 != (v15 / 1e10))) {
                                                v43 = (-(fi42))
                                            }
                                            if false {
                                                if true {
                                                    switch -669 {
                                                    case -561:
                                                        if false {
                                                            fmt.Println(v15, v16, v16)
                                                            var v44 uint8 = uint8(v39)
                                                            _ = v44
                                                        } else {
                                                            if len(v14) > 0 {
                                                                v14[0] = fi40
                                                            }
                                                        }
                                                    case 684:
                                                        if true {
                                                            v45 := GenS5{F0: (v43 | v43)}
                                                            v46 := GenS6{F0: "foo bar", F1: false}
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                    case 128:
                                                        fmt.Println(v39, fi40)
                                                    default:
                                                        fmt.Println(v39, fi42)
                                                    }
                                                    var v47 int32 = v1
                                                    var v48 int8 = 127
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    if false {
                                                        if len(v14) >= 2 {
                                                            v14 = v14[1:len(v14)]
                                                        }
                                                        switch v43 {
                                                        case 42:
                                                            v3 = v3
                                                        case 127:
                                                            v3 -= 747
                                                        case 423:
                                                            var v49 int16 = 1
                                                            _ = v49
                                                        }
                                                    }
                                                    var v50 int8 = int8(v16)
                                                    v1 -= v1
                                                    _ = v50
                                                }
                                                v51 := GenS7{F0: (!(true && false)), F1: v2}
                                                if len(v14) > 0 {
                                                    v14[0] = -1
                                                }
                                                _ = v51
                                            } else {
                                                switch (fi42 | (fi42 / ((v43) | 1))) {
                                                case 441:
                                                    switch "\n\t" {
                                                    case "hello":
                                                        v3 = (v3 ^ (v3 / ((v3) | 1)))
                                                    case "\n\t":
                                                        v52 := GenS3{F0: 18446744073709551615}
                                                        _ = v52
                                                    }
                                                }
                                            }
                                        }
                                        var v53 uint32 = (v16 / ((uint32(v2)) | 1))
                                        v54 := GenS1{F0: v43, F1: v1}
                                        var v55 float32 = (1e10 + 1.5)
                                        _ = v53
                                        _ = v54
                                        _ = v55
                                    }
                                    _ = v43
                                }
                                _ = fi42
                            }
                        default:
                            fmt.Println(v15, v1)
                        }
                        _ = fi40
                    }
                }
                switch int(v1) {
                case -9223372036854775808:
                    var v56 uint64 = uint64(v3)
                    _ = v56
                case 42:
                    switch "a" {
                    case "hello":
                        switch "a" {
                        case "a":
                            for fi57 := 0; fi57 < 13; fi57++ {
                                _ = len(v14)
                                _ = fi57
                            }
                        default:
                            if len(v14) > 0 {
                                v14[0] = 782
                            }
                        }
                    case "":
                        if len(v14) > 0 {
                            v14[0] = 135
                        }
                    case "a":
                        v16 += v16
                    }
                case 127:
                    var v58 int16 = 32767
                    _ = v58
                }
                fmt.Println(v39, v15)
            }
            _ = v39
        }
        v16 += (v16 << ((uint(v16)) & 63))
    }
    v2 += v2
    for fi59 := 0; fi59 < 9; fi59++ {
        if true {
            v2 -= v2
            for fi60 := 0; fi60 < 8; fi60++ {
                if false {
                    var v61 float64 = ((((-2.25 / (-2.25 * ((((float64(v15) / 1.0) / 0.1) * -2.25) / 1e10))) * 0.1) + 0.0) / (0.0 + 0.1))
                    if true {
                        switch fi59 {
                        case 447:
                            if (!true) {
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                v62 := GenS1{F0: (fi60 + 42), F1: (-(v1))}
                                _ = v62
                            }
                        case -2147483648:
                            switch "\n\t" {
                            case "":
                                v61 += 0.0
                            case "\n\t":
                                if (true || ((v61 / 3.14159) >= 1.5)) {
                                    _ = len(v14)
                                    if len(v14) >= 2 {
                                        v14 = v14[1:len(v14)]
                                    }
                                    var v63 int32 = (v1 & 42)
                                    _ = v63
                                } else {
                                    v64 := GenS4{F0: ((0 > 240) || true), F1: 179}
                                    if true {
                                        v16 += (v16 ^ (v16 | v16))
                                        if (!(v16 >= (^(v16)))) {
                                            v65 := GenS7{F0: true, F1: v2}
                                            if (v15 <= v15) {
                                                if (-190 == -1) {
                                                    if true {
                                                        switch fi60 {
                                                        case 9223372036854775807:
                                                            var v66 uint8 = 0
                                                            _ = v66
                                                        case 0:
                                                            fmt.Println(v61, v16, v1)
                                                            if true {
                                                                v14 = append(v14, (fi59 - -9223372036854775808))
                                                                var v67 float32 = v15
                                                                if (!((v3 << ((v3) & 63)) == (v3 * (^(v3))))) {
                                                                    v14 = append(v14, fi59)
                                                                    v15 = v15
                                                                }
                                                                _ = v67
                                                            } else {
                                                                v14 = append(v14, (-(fi60)))
                                                                switch "" {
                                                                case "a":
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        switch (^(fi59)) {
                                                                        case 797:
                                                                            _ = len(v14)
                                                                        case 9223372036854775807:
                                                                            var v68 uint32 = v16
                                                                            _ = v68
                                                                        default:
                                                                            v61 = (float64(-2.25) - v61)
                                                                        }
                                                                    case "foo bar":
                                                                        var v69 bool = (0 > 1)
                                                                        _ = v69
                                                                    default:
                                                                        v70 := GenS4{F0: true, F1: 198}
                                                                        _ = v70
                                                                    }
                                                                    fmt.Println(v2)
                                                                    v16 -= v16
                                                                case "foo bar":
                                                                    _ = len(v14)
                                                                }
                                                                if true {
                                                                    v3 = v3
                                                                    if false {
                                                                        v1 += v1
                                                                        var v71 uint = (v3 >> ((4294967295) & 63))
                                                                        _ = v71
                                                                    } else {
                                                                        var v72 uint32 = 975
                                                                        var v73 int64 = -9223372036854775808
                                                                        v74 := GenS6{F0: (("hello" + ("" + "hello")) + "a"), F1: true}
                                                                        _ = v72
                                                                        _ = v73
                                                                        _ = v74
                                                                    }
                                                                    if false {
                                                                        v15 = float32(0.0)
                                                                    }
                                                                    var v75 uint16 = uint16(v2)
                                                                    _ = v75
                                                                }
                                                            }
                                                            v15 = v15
                                                        default:
                                                            var v76 uint = (^(v3))
                                                            _ = v76
                                                        }
                                                        fmt.Println(v15, v2)
                                                    }
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    if false {
                                                        v1 = (v1 << ((2) & 63))
                                                    } else {
                                                        fmt.Println(v16)
                                                        var v77 uint = 1
                                                        if true {
                                                            if false {
                                                                v78 := GenS8{F0: genFn1(v3, -1, 1)}
                                                                if len(v14) >= 2 {
                                                                    v14 = v14[1:len(v14)]
                                                                }
                                                                var v79 uint32 = v16
                                                                var v80 string = ("" + "hello")
                                                                _ = v78
                                                                _ = v79
                                                                _ = v80
                                                            } else {
                                                                if ((v3 + 18446744073709551615) <= v77) {
                                                                    v61 = 3.14159
                                                                    if len(v14) >= 2 {
                                                                        v14 = v14[1:len(v14)]
                                                                    }
                                                                    if len(v14) >= 2 {
                                                                        v14 = v14[1:len(v14)]
                                                                    }
                                                                }
                                                            }
                                                            switch (fi59 * 127) {
                                                            case 9223372036854775807:
                                                                if true {
                                                                    v14 = append(v14, fi59)
                                                                    var v81 string = ("" + "")
                                                                    _ = v81
                                                                }
                                                            default:
                                                                v14 = append(v14, 9223372036854775807)
                                                            }
                                                            if true {
                                                                if false {
                                                                    v14 = append(v14, int(v16))
                                                                    v82 := GenS8{F0: int16(v2)}
                                                                    v83 := GenS9{F0: v1, F1: 1}
                                                                    var v84 uint32 = 763
                                                                    v16 += v16
                                                                    _ = v82
                                                                    _ = v83
                                                                    _ = v84
                                                                } else {
                                                                    v85 := GenS10{F0: ((("a" + "") + "hello") + "\n\t"), F1: ("a" + "hello")}
                                                                    var v86 int32 = 1
                                                                    if (false && (592 != 42)) {
                                                                        switch 127 {
                                                                        case 2147483647:
                                                                            if true {
                                                                                switch fi60 {
                                                                                case -758:
                                                                                    var v87 bool = (42 >= 42)
                                                                                    _ = v87
                                                                                case 32767:
                                                                                    v16 += (^(v16))
                                                                                    _ = len(v14)
                                                                                case -38:
                                                                                    var v88 uint16 = 140
                                                                                    _ = v88
                                                                                default:
                                                                                    if (-9223372036854775808 <= 9223372036854775807) {
                                                                                        if (("foo bar" != "\n\t") && true) {
                                                                                            var v89 uint8 = 230
                                                                                            _ = v89
                                                                                        } else {
                                                                                            switch -260 {
                                                                                            case -2147483648:
                                                                                                v90 := GenS6{F0: "", F1: (("hello" + "") == "a")}
                                                                                                _ = v90
                                                                                            case 0:
                                                                                                if false {
                                                                                                    v86 -= 1
                                                                                                } else {
                                                                                                    v86 = v86
                                                                                                }
                                                                                            case -9223372036854775808:
                                                                                                switch "a" {
                                                                                                case "":
                                                                                                    fmt.Println(v3, v86, v16)
                                                                                                case "hello":
                                                                                                    switch fi59 {
                                                                                                    case 127:
                                                                                                        if false {
                                                                                                            if false {
                                                                                                                v15 -= v15
                                                                                                                if true {
                                                                                                                    fmt.Println(v86, fi59)
                                                                                                                    switch "" {
                                                                                                                    case "a":
                                                                                                                        var v91 int16 = 1000
                                                                                                                        _ = v91
                                                                                                                    }
                                                                                                                    fmt.Println(v3, v3, v61)
                                                                                                                } else {
                                                                                                                    if (true || false) {
                                                                                                                        v92 := GenS11{F0: -1, F1: 38, F2: false}
                                                                                                                        var v93 uint8 = 161
                                                                                                                        if true {
                                                                                                                            var v94 string = (("" + "") + "")
                                                                                                                            v61 = float64(1.5)
                                                                                                                            if true {
                                                                                                                                if false {
                                                                                                                                    if true {
                                                                                                                                        v95 := GenS9{F0: 269, F1: -1}
                                                                                                                                        fmt.Println(v1)
                                                                                                                                        var v96 float32 = v15
                                                                                                                                        _ = v95
                                                                                                                                        _ = v96
                                                                                                                                    } else {
                                                                                                                                        if (!(true && (v61 < v61))) {
                                                                                                                                            if true {
                                                                                                                                                v16 = 634
                                                                                                                                            } else {
                                                                                                                                                if ((false || false) || false) {
                                                                                                                                                    v77 = 255
                                                                                                                                                    if false {
                                                                                                                                                        var v97 bool = false
                                                                                                                                                        var v98 uint64 = (^(v2))
                                                                                                                                                        fmt.Println(v16)
                                                                                                                                                        _ = v97
                                                                                                                                                        _ = v98
                                                                                                                                                    }
                                                                                                                                                    v14 = append(v14, fi60)
                                                                                                                                                } else {
                                                                                                                                                    v2 += v2
                                                                                                                                                    var v99 float64 = 1e10
                                                                                                                                                    v94 = v94
                                                                                                                                                    _ = v99
                                                                                                                                                }
                                                                                                                                                _ = len(v14)
                                                                                                                                                v100 := GenS8{F0: 1000}
                                                                                                                                                _ = v100
                                                                                                                                            }
                                                                                                                                            var v101 uint64 = 0
                                                                                                                                            v101 = (^(v2))
                                                                                                                                            var v102 int64 = -9223372036854775808
                                                                                                                                            _ = v101
                                                                                                                                            _ = v102
                                                                                                                                        }
                                                                                                                                        v103 := GenS12{F0: (^(v93))}
                                                                                                                                        var v104 uint32 = (v16 << ((65535) & 63))
                                                                                                                                        var v105 int64 = 178
                                                                                                                                        _ = v103
                                                                                                                                        _ = v104
                                                                                                                                        _ = v105
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                v94 = (("hello" + "a") + v94)
                                                                                                                                if ((!false) && (!false)) {
                                                                                                                                    var v106 uint16 = 0
                                                                                                                                    if (v15 == (v15 + (v15 + v15))) {
                                                                                                                                        v2 -= 1
                                                                                                                                        switch ((((v94 + "") + v94) + "\n\t") + v94) {
                                                                                                                                        case "a":
                                                                                                                                            switch v94 {
                                                                                                                                            case "\n\t":
                                                                                                                                                switch "foo bar" {
                                                                                                                                                case "a":
                                                                                                                                                    v16 = v16
                                                                                                                                                case "hello":
                                                                                                                                                    v61 = v61
                                                                                                                                                default:
                                                                                                                                                    v15 = 0.0
                                                                                                                                                }
                                                                                                                                            case "hello":
                                                                                                                                                if false {
                                                                                                                                                    v61 = v61
                                                                                                                                                    v86 -= (v1 | v1)
                                                                                                                                                }
                                                                                                                                            case "foo bar":
                                                                                                                                                if len(v14) > 0 {
                                                                                                                                                    v14[0] = 32768
                                                                                                                                                }
                                                                                                                                            default:
                                                                                                                                                v3 = uint(v106)
                                                                                                                                            }
                                                                                                                                        default:
                                                                                                                                            v106 += 42
                                                                                                                                        }
                                                                                                                                        fmt.Println(v93, v61)
                                                                                                                                    } else {
                                                                                                                                        if true {
                                                                                                                                            v107 := GenS10{F0: v94, F1: v94}
                                                                                                                                            if true {
                                                                                                                                                v108 := GenS5{F0: 0}
                                                                                                                                                v16 = 527
                                                                                                                                                v2 = (v2 - (v2 >> ((v3) & 63)))
                                                                                                                                                _ = v108
                                                                                                                                            }
                                                                                                                                            _ = v107
                                                                                                                                        } else {
                                                                                                                                            fmt.Println(v86, fi60)
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    var v109 uint64 = (v2 << ((192) & 63))
                                                                                                                                    _ = v106
                                                                                                                                    _ = v109
                                                                                                                                }
                                                                                                                                fmt.Println(v77, v15, v1)
                                                                                                                            }
                                                                                                                            _ = v94
                                                                                                                        }
                                                                                                                        _ = v92
                                                                                                                        _ = v93
                                                                                                                    } else {
                                                                                                                        var v110 float32 = float32(0.1)
                                                                                                                        _ = v110
                                                                                                                    }
                                                                                                                    if (v2 < (v2 * 2)) {
                                                                                                                        v15 = v15
                                                                                                                        if (127 < (fi60 << (((v77 & 0)) & 63))) {
                                                                                                                            if false {
                                                                                                                                var v111 string = "hello"
                                                                                                                                _ = v111
                                                                                                                            } else {
                                                                                                                                var v112 int8 = -128
                                                                                                                                v15 -= v15
                                                                                                                                var v113 uint32 = (v16 | (v16 - (v16 | v16)))
                                                                                                                                v14 = append(v14, (fi60 + (fi60 << (((v3 + v3)) & 63))))
                                                                                                                                var v114 int8 = -49
                                                                                                                                _ = v112
                                                                                                                                _ = v113
                                                                                                                                _ = v114
                                                                                                                            }
                                                                                                                            var v115 float32 = 0.0
                                                                                                                            fmt.Println(fi59)
                                                                                                                            _ = v115
                                                                                                                        } else {
                                                                                                                            var v116 string = ("a" + ((("\n\t" + ("" + "a")) + "a") + "a"))
                                                                                                                            fmt.Println(v77)
                                                                                                                            var v117 uint8 = 127
                                                                                                                            _ = v116
                                                                                                                            _ = v117
                                                                                                                        }
                                                                                                                        var v118 bool = ((32767 != 32767) && true)
                                                                                                                        _ = v118
                                                                                                                    } else {
                                                                                                                        switch "hello" {
                                                                                                                        case "hello":
                                                                                                                            switch (fi60 - fi60) {
                                                                                                                            case 9223372036854775807:
                                                                                                                                v2 = 203
                                                                                                                            case 4294967295:
                                                                                                                                v61 = v61
                                                                                                                            case -9223372036854775808:
                                                                                                                                v77 = v3
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            switch "" {
                                                                                                                            case "a":
                                                                                                                                fmt.Println(v77)
                                                                                                                            case "hello":
                                                                                                                                v16 = v16
                                                                                                                            }
                                                                                                                        }
                                                                                                                        if (42 == int8(v2)) {
                                                                                                                            v119 := GenS8{F0: 1}
                                                                                                                            var v120 int32 = (v1 ^ -40)
                                                                                                                            switch 2147483647 {
                                                                                                                            case 1:
                                                                                                                                var v121 uint16 = 32767
                                                                                                                                _ = v121
                                                                                                                            }
                                                                                                                            _ = v119
                                                                                                                            _ = v120
                                                                                                                        }
                                                                                                                        var v122 int32 = v1
                                                                                                                        _ = v122
                                                                                                                    }
                                                                                                                    var v123 int = fi60
                                                                                                                    v1 = v86
                                                                                                                    _ = v123
                                                                                                                }
                                                                                                            } else {
                                                                                                                fmt.Println(v3, v2)
                                                                                                            }
                                                                                                        } else {
                                                                                                            v86 -= v1
                                                                                                            if len(v14) > 0 {
                                                                                                                v14[0] = (fi60 - 127)
                                                                                                            }
                                                                                                            var v124 float64 = v61
                                                                                                            if true {
                                                                                                                if len(v14) > 0 {
                                                                                                                    v14[0] = (^(fi60))
                                                                                                                }
                                                                                                                if false {
                                                                                                                    if ((42 == 607) && (false || (v15 != float32(v61)))) {
                                                                                                                        fmt.Println(v124, v124)
                                                                                                                        v125 := GenS13{F0: "a", F1: (v2 / ((v2) | 1)), F2: 1}
                                                                                                                        _ = v125
                                                                                                                    }
                                                                                                                    var v126 int32 = (v86 >> (((v3 / ((v3) | 1))) & 63))
                                                                                                                    v127 := GenS10{F0: ("" + "a"), F1: "foo bar"}
                                                                                                                    _ = v126
                                                                                                                    _ = v127
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v128 int32 = (v86 >> ((v3) & 63))
                                                                                                                v3 += v3
                                                                                                                _ = v128
                                                                                                            }
                                                                                                            _ = v124
                                                                                                        }
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    _ = len(v14)
                                                                                                }
                                                                                            }
                                                                                            v14 = append(v14, fi59)
                                                                                            var v129 float64 = v61
                                                                                            v3 += 255
                                                                                            _ = v129
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        case 42:
                                                                            var v130 int64 = 1
                                                                            _ = v130
                                                                        case 128:
                                                                            if false {
                                                                                switch fi60 {
                                                                                case -356:
                                                                                    v131 := GenS8{F0: 32767}
                                                                                    _ = v131
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        switch (("" + "\n\t") + "a") {
                                                                        case "hello":
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                _ = len(v14)
                                                                                v3 = (v3 / ((255) | 1))
                                                                                if false {
                                                                                    if false {
                                                                                        if true {
                                                                                            fmt.Println(v3)
                                                                                        } else {
                                                                                            if true {
                                                                                                var v132 bool = ((v61 <= v61) || false)
                                                                                                var v133 int64 = 0
                                                                                                if true {
                                                                                                    v1 = (^(v1))
                                                                                                }
                                                                                                _ = v132
                                                                                                _ = v133
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v1 += (^(v1))
                                                                                }
                                                                            case "a":
                                                                                var v134 uint = (v77 % (((v3 / (((^(v77))) | 1))) | 1))
                                                                                _ = v134
                                                                            default:
                                                                                fmt.Println(v1)
                                                                            }
                                                                        case "foo bar":
                                                                            v86 -= (v86 & 560)
                                                                        default:
                                                                            v135 := GenS14{F0: 127, F1: 42}
                                                                            _ = v135
                                                                        }
                                                                        v2 = v2
                                                                        if false {
                                                                            v2 -= (^(v2))
                                                                        }
                                                                        v14 = append(v14, -157)
                                                                        var v136 uint64 = (v2 % ((4294967295) | 1))
                                                                        _ = v136
                                                                    }
                                                                    _ = v85
                                                                    _ = v86
                                                                }
                                                                v137 := GenS15{F0: (fi60 + (fi59 + (fi60 % ((32767) | 1)))), F1: -128, F2: (^(fi60))}
                                                                switch "a" {
                                                                case "a":
                                                                    var v138 float64 = v61
                                                                    _ = v138
                                                                case "hello":
                                                                    var v139 uint64 = v2
                                                                    _ = v139
                                                                case "":
                                                                    v15 -= v15
                                                                    if (uint32(v2) >= (^(v16))) {
                                                                        v16 = (v16 - 832)
                                                                        var v140 int64 = 425
                                                                        _ = v140
                                                                    } else {
                                                                        var v141 int = fi60
                                                                        if true {
                                                                            v1 = (v1 * -1)
                                                                            v77 = v3
                                                                            switch (fi60 * v141) {
                                                                            case 32767:
                                                                                if false {
                                                                                    v1 += v1
                                                                                    var v142 bool = ((v77 > v77) && ("a" >= ""))
                                                                                    if false {
                                                                                        if (((1.5 + v61) + v61) > v61) {
                                                                                            var v143 float64 = float64(v15)
                                                                                            v144 := GenS12{F0: 181}
                                                                                            v14 = append(v14, fi59)
                                                                                            _ = v143
                                                                                            _ = v144
                                                                                        }
                                                                                        var v145 int16 = -32768
                                                                                        v16 -= 4294967295
                                                                                        _ = v145
                                                                                    }
                                                                                    _ = v142
                                                                                } else {
                                                                                    if len(v14) >= 2 {
                                                                                        v14 = v14[1:len(v14)]
                                                                                    }
                                                                                    if false {
                                                                                        switch 4294967295 {
                                                                                        case -702:
                                                                                            switch "hello" {
                                                                                            case "":
                                                                                                var v146 int32 = (v1 ^ v1)
                                                                                                _ = v146
                                                                                            }
                                                                                        case 1:
                                                                                            _ = len(v14)
                                                                                            var v147 bool = ((0 > -1) && ((!false) && false))
                                                                                            fmt.Println(v147, v141, v16)
                                                                                            _ = v147
                                                                                        }
                                                                                        var v148 uint16 = 32767
                                                                                        v1 -= (v1 & v1)
                                                                                        _ = v148
                                                                                    }
                                                                                }
                                                                            case -2147483648:
                                                                                var v149 int16 = 1000
                                                                                _ = v149
                                                                            case 4294967295:
                                                                                if true {
                                                                                    var v150 int16 = 0
                                                                                    var v151 float32 = 1e10
                                                                                    _ = v150
                                                                                    _ = v151
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v152 := GenS16{F0: (^(fi59)), F1: ""}
                                                                            _ = v152
                                                                        }
                                                                        fmt.Println(v15, v77)
                                                                        _ = v141
                                                                    }
                                                                    var v153 uint8 = 127
                                                                    fmt.Println(fi59)
                                                                    if ((("foo bar" + (("" + "a") + "\n\t")) == ((("hello" + ("hello" + "a")) + "a") + "\n\t")) || (v153 < v153)) {
                                                                        _ = len(v14)
                                                                        v2 -= 0
                                                                        v154 := GenS5{F0: int(v77)}
                                                                        var v155 int8 = 100
                                                                        _ = v154
                                                                        _ = v155
                                                                    } else {
                                                                        var v156 int8 = 103
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            v14 = append(v14, fi60)
                                                                        case "foo bar":
                                                                            var v157 int32 = -221
                                                                            _ = v157
                                                                        case "hello":
                                                                            fmt.Println(fi60)
                                                                        }
                                                                        _ = v156
                                                                    }
                                                                    var v158 uint16 = 170
                                                                    _ = v153
                                                                    _ = v158
                                                                }
                                                                _ = v137
                                                            } else {
                                                                v159 := GenS8{F0: -1}
                                                                if len(v14) >= 2 {
                                                                    v14 = v14[1:len(v14)]
                                                                }
                                                                var v160 bool = (!true)
                                                                _ = v159
                                                                _ = v160
                                                            }
                                                        } else {
                                                            var v161 uint = v77
                                                            _ = len(v14)
                                                            if (!(42 < 106)) {
                                                                v162 := GenS5{F0: -296}
                                                                _ = v162
                                                            } else {
                                                                switch fi60 {
                                                                case 128:
                                                                    v14 = append(v14, fi59)
                                                                default:
                                                                    v16 -= (^(v16))
                                                                }
                                                                var v163 int32 = v1
                                                                _ = v163
                                                            }
                                                            _ = v161
                                                        }
                                                        _ = v77
                                                    }
                                                case "a":
                                                    if len(v14) >= 2 {
                                                        v14 = v14[1:len(v14)]
                                                    }
                                                }
                                            } else {
                                                if len(v14) > 0 {
                                                    v14[0] = (fi59 >> (((v3 + v3)) & 63))
                                                }
                                                fmt.Println(v2)
                                                var v164 float32 = (1e10 + v15)
                                                _ = v164
                                            }
                                            var v165 uint64 = 42
                                            _ = v165
                                            _ = v65
                                        }
                                        v16 += (^(v16))
                                    }
                                    switch "" {
                                    case "\n\t":
                                        v166 := GenS17{F0: 9223372036854775807, F1: v1}
                                        _ = v166
                                    case "hello":
                                        v167 := GenS17{F0: -12, F1: -648}
                                        _ = v167
                                    case "foo bar":
                                        if false {
                                            switch 127 {
                                            case 42:
                                                var v168 int16 = 70
                                                _ = v168
                                            case 0:
                                                var v169 uint32 = v16
                                                _ = v169
                                            case 2147483647:
                                                switch fi60 {
                                                case 127:
                                                    v170 := GenS15{F0: fi59, F1: -1, F2: 2147483647}
                                                    _ = v170
                                                case 58:
                                                    if (!(!false)) {
                                                        v171 := GenS18{F0: 2, F1: (v3 << ((v3) & 63))}
                                                        v16 = v16
                                                        _ = v171
                                                    } else {
                                                        var v172 uint64 = v2
                                                        v2 = 4294967295
                                                        if (v1 <= (v1 - v1)) {
                                                            if len(v14) >= 2 {
                                                                v14 = v14[1:len(v14)]
                                                            }
                                                        } else {
                                                            v173 := GenS5{F0: int(v172)}
                                                            v3 += 18446744073709551615
                                                            var v174 int64 = -80
                                                            _ = v173
                                                            _ = v174
                                                        }
                                                        _ = v172
                                                    }
                                                    fmt.Println(v61)
                                                    if true {
                                                        fmt.Println(v1, fi60, v1)
                                                        v14 = append(v14, fi59)
                                                        switch ("\n\t" + ("\n\t" + ("\n\t" + ("foo bar" + "\n\t")))) {
                                                        case "\n\t":
                                                            var v175 float32 = v15
                                                            _ = v175
                                                        case "":
                                                            v14 = append(v14, fi60)
                                                        case "a":
                                                            if (!true) {
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    if len(v14) > 0 {
                                                                        v14[0] = (fi59 + fi60)
                                                                    }
                                                                case "":
                                                                    v176 := GenS19{F0: 9223372036854775807, F1: 1, F2: (v2 >> (((v3 / (((v3 / ((1) | 1))) | 1))) & 63))}
                                                                    _ = v176
                                                                case "hello":
                                                                    v3 -= v3
                                                                default:
                                                                    v177 := GenS9{F0: (v1 / ((v1) | 1)), F1: 1}
                                                                    _ = v177
                                                                }
                                                            }
                                                        default:
                                                            if len(v14) >= 2 {
                                                                v14 = v14[1:len(v14)]
                                                            }
                                                        }
                                                        var v178 float64 = v61
                                                        fmt.Println(v16, v178, v16)
                                                        v179 := GenS1{F0: fi60, F1: 42}
                                                        _ = v178
                                                        _ = v179
                                                    }
                                                }
                                            default:
                                                fmt.Println(v61, v2, v16)
                                                v2 = v2
                                            }
                                            if false {
                                                if (("" + ("\n\t" + "")) > "\n\t") {
                                                    if ("\n\t" <= ("\n\t" + "")) {
                                                        var v180 bool = (true && false)
                                                        var v181 int8 = 0
                                                        var v182 int8 = 90
                                                        v182 += v182
                                                        _ = v180
                                                        _ = v181
                                                        _ = v182
                                                    } else {
                                                        fmt.Println(v16, v3)
                                                        fmt.Println(v1, v1, v16)
                                                        v16 = 65535
                                                    }
                                                    var v183 int16 = 13
                                                    _ = v183
                                                } else {
                                                    var v184 int32 = v1
                                                    if (!("\n\t" >= ("hello" + "foo bar"))) {
                                                        if len(v14) > 0 {
                                                            v14[0] = fi59
                                                        }
                                                    }
                                                    var v185 uint32 = v16
                                                    if len(v14) >= 2 {
                                                        v14 = v14[1:len(v14)]
                                                    }
                                                    var v186 int16 = genFn1(v3, v1, 42)
                                                    var v187 string = ""
                                                    _ = v184
                                                    _ = v185
                                                    _ = v186
                                                    _ = v187
                                                }
                                            } else {
                                                fmt.Println(v61, v16)
                                                if len(v14) > 0 {
                                                    v14[0] = fi59
                                                }
                                                _ = len(v14)
                                            }
                                        } else {
                                            var v188 bool = ("" <= "hello")
                                            _ = v188
                                        }
                                    }
                                    _ = v64
                                }
                            default:
                                v3 = 1
                            }
                            var v189 uint32 = 317
                            fmt.Println(v61)
                            _ = v189
                        default:
                            v61 = v61
                            v3 = (v3 * (v3 / (((^(v3))) | 1)))
                            v14 = append(v14, fi59)
                        }
                    } else {
                        fmt.Println(v1, v15)
                        var v190 float64 = v61
                        _ = v190
                    }
                    fmt.Println(fi60)
                    _ = v61
                }
                _ = fi60
            }
            for fi191 := 0; fi191 < 9; fi191++ {
                v192 := GenS20{F0: -2147483648}
                var v193 uint32 = v16
                var v194 uint8 = 1
                _ = fi191
                _ = v192
                _ = v193
                _ = v194
            }
        }
        _ = fi59
    }
    v16 = 42
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
