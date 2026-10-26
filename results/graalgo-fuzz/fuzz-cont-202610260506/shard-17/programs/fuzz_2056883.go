package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 bool; F1 int; F2 bool }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 int8; F1 int16 }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int32; F1 uint; F2 bool }
type GenS12 struct { F0 uint32; F1 uint32 }
type GenS13 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 int8 = 1
    var v2 float32 = 1e10
    var v3 uint = 65535
    var v4 float32 = 1e10
    var v5 uint64 = 4294967295
    var v6 string = "a"
    v7 := GenS1{F0: (v1 / ((v1) | 1))}
    var v8 float64 = 3.14159
    if true {
        for fi9 := 0; fi9 < 5; fi9++ {
            v8 = (v8 - (-2.25 + (v8 * (((((((-2.25 + v8) + v8) * v8) - float64(v2)) + -2.25) + (v8 / 1e10)) + (-2.25 * v8)))))
            for fi10 := 0; fi10 < 12; fi10++ {
                v6 = (("hello" + v6) + v6)
                v5 = v5
                if (v8 <= 1e10) {
                    v1 = v1
                    if (!("foo bar" != "hello")) {
                        v4 += float32(float64(v4))
                        switch "\n\t" {
                        case "\n\t":
                            v5 -= (v5 * (v5 | v5))
                        }
                    } else {
                        v11 := []int{int(v1)}
                        v3 -= (v3 + (v3 & 255))
                        switch int(v1) {
                        case 4294967295:
                            v2 -= v2
                        case 267:
                            if false {
                                if (2 == 2) {
                                    if true {
                                        var v12 float64 = v8
                                        fmt.Println(v2, v5, v2)
                                        _ = v12
                                    } else {
                                        if (!true) {
                                            var v13 int32 = -2147483648
                                            if (-1 >= int64(v5)) {
                                                v1 += 127
                                                switch "" {
                                                case "\n\t":
                                                    v13 = v13
                                                case "":
                                                    if (0 > 509) {
                                                        v14 := GenS2{F0: 2147483648, F1: 32767}
                                                        _ = v14
                                                    }
                                                }
                                                fmt.Println(v3, v6, v4)
                                            } else {
                                                if len(v11) > 0 {
                                                    v11[0] = fi10
                                                }
                                                if (fi10 == fi10) {
                                                    var v15 bool = ((v3 ^ (v3 ^ 4294967295)) <= v3)
                                                    var v16 string = "foo bar"
                                                    var v17 string = ("a" + (v16 + v6))
                                                    _ = v15
                                                    _ = v16
                                                    _ = v17
                                                } else {
                                                    var v18 int8 = v1
                                                    v19 := GenS1{F0: v1}
                                                    switch "" {
                                                    case "\n\t":
                                                        var v20 bool = ((v13 >> ((4294967295) & 63)) >= (v13 + int32(fi9)))
                                                        _ = v20
                                                    }
                                                    _ = v18
                                                    _ = v19
                                                }
                                                _ = len(v11)
                                            }
                                            v21 := GenS2{F0: 65535, F1: 42}
                                            _ = v13
                                            _ = v21
                                        }
                                    }
                                    if (v1 == v1) {
                                        v6 = (v6 + v6)
                                        _ = len(v11)
                                        v1 += (v1 & -7)
                                    } else {
                                        fmt.Println(fi9)
                                        if (true && (!true)) {
                                            v6 = ((("foo bar" + (v6 + "")) + (v6 + (v6 + ((v6 + v6) + v6)))) + v6)
                                        } else {
                                            v1 = 0
                                        }
                                    }
                                    v8 += 1e10
                                }
                                v6 = ""
                            } else {
                                v22 := GenS2{F0: 2147483648, F1: 36}
                                v3 = v3
                                _ = v22
                            }
                        default:
                            var v23 int32 = 42
                            _ = v23
                        }
                        _ = v11
                    }
                    var v24 int16 = 32767
                    _ = v24
                }
                _ = fi10
            }
            switch fi9 {
            case 9223372036854775807:
                fmt.Println(fi9, v6, v3)
            case -471:
                fmt.Println(v6, fi9)
            default:
                var v25 int8 = 42
                _ = v25
            }
            _ = fi9
        }
    } else {
        if true {
            v26 := GenS2{F0: 416, F1: uint16(v5)}
            var v27 uint32 = 2
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 12; fi28++ {
            v1 += v1
            var v29 uint16 = 2
            _ = fi28
            _ = v29
        }
        switch "foo bar" {
        case "hello":
            v30 := GenS3{F0: (true || (false || true)), F1: 856, F2: (true || false)}
            _ = v30
        case "foo bar":
            fmt.Println(v2, v2, v5)
        }
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 float32 = v2
        _ = fi31
        _ = v32
    }
    var v33 float64 = (v8 - 0.0)
    if false {
        v3 -= v3
        fmt.Println(v5, v4)
    } else {
        fmt.Println(v1, v3)
    }
    v34 := GenS1{F0: v1}
    fmt.Println(v1, v2)
    if (v5 <= (v5 >> ((4294967295) & 63))) {
        if (((65535 > 626) && false) || false) {
            for fi35 := 0; fi35 < 16; fi35++ {
                if true {
                    v36 := GenS1{F0: -128}
                    if (128 >= 40) {
                        v37 := []int{fi35, 768, (fi35 + (fi35 & (-(fi35)))), (fi35 & 32767)}
                        fmt.Println(v6, v5)
                        _ = v37
                    }
                    _ = v36
                }
                _ = fi35
            }
            if false {
                v38 := []int{127, 2147483647, -564, 9223372036854775807}
                v39 := GenS2{F0: 628, F1: 149}
                _ = v38
                _ = v39
            }
        }
        v40 := []int{32768, -193}
        _ = v40
    }
    if (-2147483648 != 66) {
        v1 += v1
    }
    for fi41 := 0; fi41 < 14; fi41++ {
        for fi42 := 0; fi42 < 7; fi42++ {
            v4 -= v2
            if (!(fi41 <= (-(fi42)))) {
                v43 := []int{2147483647, -928}
                fmt.Println(v3)
                v1 = (v1 & v1)
                _ = v43
            }
            v44 := []int{}
            if len(v44) > 0 {
                v44[0] = fi41
            }
            if false {
                var v45 int64 = 1
                var v46 uint16 = 65535
                v47 := GenS1{F0: 127}
                _ = v45
                _ = v46
                _ = v47
            } else {
                v6 = "hello"
                v6 = (v6 + (v6 + (((v6 + "foo bar") + (v6 + v6)) + ("foo bar" + v6))))
            }
            _ = fi42
            _ = v44
        }
        _ = fi41
    }
    var v48 uint16 = 129
    if false {
        fmt.Println(v48, v48)
        fmt.Println(v3, v4)
        switch v6 {
        case "\n\t":
            switch 1 {
            case 4294967295:
                if true {
                    v4 = 1e10
                    var v49 uint32 = uint32(v48)
                    _ = v49
                } else {
                    switch v6 {
                    case "":
                        for fi50 := 0; fi50 < 20; fi50++ {
                            var v51 int64 = int64(v5)
                            for fi52 := 0; fi52 < 4; fi52++ {
                                var v53 int64 = (v51 | 764)
                                if false {
                                    var v54 string = v6
                                    var v55 int64 = v51
                                    var v56 int16 = -1
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                }
                                _ = fi52
                                _ = v53
                            }
                            _ = fi50
                            _ = v51
                        }
                    case "foo bar":
                        v6 = v6
                    case "a":
                        fmt.Println(v3, v4, v33)
                    default:
                        v3 += 829
                    }
                    var v57 bool = (0.0 < 0.1)
                    if (true || false) {
                        v58 := GenS1{F0: v1}
                        if v57 {
                            v33 = (v33 / 1.5)
                        } else {
                            v57 = ((v2 + 1e10) <= v4)
                        }
                        _ = v58
                    } else {
                        v59 := GenS4{F0: (v48 * (v48 & (v48 >> (((v3 & (^(v3)))) & 63)))), F1: -1}
                        if (v57 || false) {
                            v6 = (v6 + v6)
                            fmt.Println(v3, v3, v2)
                        }
                        var v60 int16 = int16(v1)
                        _ = v59
                        _ = v60
                    }
                    _ = v57
                }
            case -267:
                var v61 uint64 = (v5 - 18446744073709551615)
                _ = v61
            case 127:
                var v62 string = ""
                _ = v62
            }
        }
    }
    v63 := GenS5{F0: 781, F1: 65535, F2: v5}
    if (754 != (v3 + v3)) {
        switch "" {
        case "foo bar":
            v6 = v6
        case "a":
            var v64 bool = (618 != 685)
            _ = v64
        case "\n\t":
            if false {
                v65 := GenS6{F0: 0, F1: 0}
                _ = v65
            }
        }
        v66 := GenS7{F0: "\n\t", F1: v6}
        var v67 int64 = 810
        _ = v66
        _ = v67
    } else {
        v4 += v4
        for fi68 := 0; fi68 < 5; fi68++ {
            var v69 int16 = 1
            switch 840 {
            case -1:
                v70 := GenS8{F0: 127, F1: v69}
                _ = v70
            case -488:
                v71 := GenS1{F0: -128}
                _ = v71
            case 42:
                switch "hello" {
                case "":
                    if (true || (((false && ((^(v48)) != 290)) || (false && true)) && true)) {
                        switch (fi68 % ((-2147483648) | 1)) {
                        case -471:
                            var v72 float64 = float64((float32(v8) / v4))
                            _ = v72
                        case -211:
                            for fi73 := 0; fi73 < 2; fi73++ {
                                switch ("" + v6) {
                                case "":
                                    if (true || (v2 != 1e10)) {
                                        v74 := GenS9{F0: 186, F1: v69}
                                        if (109 <= int32(v69)) {
                                            switch fi73 {
                                            case -9223372036854775808:
                                                if true {
                                                    v3 = v3
                                                    var v75 uint = uint(v48)
                                                    _ = v75
                                                } else {
                                                    v76 := []int{fi73, (fi73 * (fi73 % ((int(v3)) | 1)))}
                                                    v5 = v5
                                                    _ = len(v76)
                                                    var v77 int = fi73
                                                    var v78 uint16 = 65535
                                                    if (v1 == (v1 & v1)) {
                                                        if (2147483647 == 2147483648) {
                                                            switch v6 {
                                                            case "\n\t":
                                                                if true {
                                                                    fmt.Println(v8, v48)
                                                                }
                                                            case "a":
                                                                if false {
                                                                    fmt.Println(fi73, v3, v4)
                                                                    if len(v76) >= 2 {
                                                                        v76 = v76[1:len(v76)]
                                                                    }
                                                                }
                                                            }
                                                            v76 = append(v76, -241)
                                                        }
                                                        var v79 string = v6
                                                        if true {
                                                            var v80 float32 = v4
                                                            _ = v80
                                                        } else {
                                                            var v81 int16 = v69
                                                            _ = v81
                                                        }
                                                        _ = v79
                                                    } else {
                                                        var v82 string = v6
                                                        v3 += (v3 - v3)
                                                        _ = v82
                                                    }
                                                    if len(v76) >= 2 {
                                                        v76 = v76[1:len(v76)]
                                                    }
                                                    if (false && true) {
                                                        var v83 int16 = (^(v69))
                                                        _ = v83
                                                    }
                                                    v8 = (1e10 - (v8 + v33))
                                                    var v84 uint32 = 230
                                                    v85 := GenS10{F0: v69}
                                                    _ = v76
                                                    _ = v77
                                                    _ = v78
                                                    _ = v84
                                                    _ = v85
                                                }
                                            }
                                            fmt.Println(v2, v1)
                                            var v86 float64 = v8
                                            _ = v86
                                        }
                                        fmt.Println(v33, v6, v33)
                                        _ = v74
                                    }
                                case "\n\t":
                                    var v87 uint64 = (v5 >> ((v3) & 63))
                                    _ = v87
                                case "a":
                                    v88 := []int{}
                                    _ = v88
                                }
                                v89 := GenS7{F0: v6, F1: v6}
                                _ = fi73
                                _ = v89
                            }
                        case 1:
                            for fi90 := 0; fi90 < 16; fi90++ {
                                if (112 > (v5 | v5)) {
                                    if (v3 != v3) {
                                        var v91 int = (fi68 | (^(fi90)))
                                        var v92 int = v91
                                        var v93 bool = (294 > -351)
                                        _ = v91
                                        _ = v92
                                        _ = v93
                                    } else {
                                        v69 = int16(v1)
                                        var v94 int16 = (v69 - (v69 | (v69 - v69)))
                                        _ = v94
                                    }
                                }
                                v95 := []int{}
                                _ = fi90
                                _ = v95
                            }
                        }
                        var v96 string = "hello"
                        _ = v96
                    } else {
                        v2 += (v4 * (0.1 / 1.0))
                        for fi97 := 0; fi97 < 3; fi97++ {
                            v98 := GenS11{F0: -2147483648, F1: uint(v48), F2: true}
                            _ = fi97
                            _ = v98
                        }
                        for fi99 := 0; fi99 < 16; fi99++ {
                            var v100 uint32 = uint32(v3)
                            _ = fi99
                            _ = v100
                        }
                    }
                case "\n\t":
                    for fi101 := 0; fi101 < 20; fi101++ {
                        switch ("a" + (("hello" + "hello") + v6)) {
                        case "hello":
                            fmt.Println(v3)
                        case "foo bar":
                            var v102 uint64 = v5
                            _ = v102
                        default:
                            v48 = uint16(v3)
                        }
                        _ = fi101
                    }
                default:
                    fmt.Println(v6, v6, v1)
                }
            default:
                v4 -= -2.25
            }
            _ = fi68
            _ = v69
        }
        for fi103 := 0; fi103 < 17; fi103++ {
            v1 += v1
            _ = fi103
        }
    }
    for fi104 := 0; fi104 < 14; fi104++ {
        if false {
            v105 := GenS12{F0: 2, F1: 65535}
            v106 := GenS13{F0: v5, F1: 42}
            _ = v105
            _ = v106
        }
        switch "\n\t" {
        case "\n\t":
            v107 := GenS4{F0: v48, F1: 2147483647}
            _ = v107
        case "hello":
            var v108 float64 = (v8 - v8)
            _ = v108
        case "a":
            switch -9223372036854775808 {
            case 1:
                for fi109 := 0; fi109 < 12; fi109++ {
                    v110 := []int{(-(fi109)), (^(fi109)), 9223372036854775807, fi109}
                    _ = fi109
                    _ = v110
                }
            case 4294967295:
                var v111 uint64 = (v5 / ((v5) | 1))
                _ = v111
            }
        }
        var v112 int8 = (v1 ^ v1)
        _ = fi104
        _ = v112
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v63
    _ = v7
    _ = v8
}
