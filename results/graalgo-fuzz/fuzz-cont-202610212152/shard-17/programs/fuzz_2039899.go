package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int16; F1 uint32 }
type GenS11 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS12 struct { F0 uint; F1 uint; F2 uint8 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int16; F1 int; F2 int16 }
type GenS16 struct { F0 uint8; F1 uint }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 string; F1 int64; F2 int16 }
type GenS19 struct { F0 uint; F1 int64; F2 string }
type GenS20 struct { F0 int8; F1 bool; F2 int8 }
type GenS21 struct { F0 uint16 }
type GenS22 struct { F0 uint64; F1 int32 }
type GenS23 struct { F0 uint16 }
type GenS24 struct { F0 uint32; F1 int32; F2 int32 }
func genFn0(p0 uint8, p1 uint64, p2 int32) uint32 {
    v5 := GenS1{F0: true, F1: (p0 % ((48) | 1))}
    v6 := GenS1{F0: (0.1 <= (-2.25 - float64(float32(0.1)))), F1: (p0 << ((0) & 63))}
    v7 := []int{0}
    fmt.Println(p0, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint8, p1 int16) uint {
    p1 = p1
    if false {
        var v8 int64 = -1
        v9 := GenS2{F0: (v8 | (v8 - 0))}
        p0 -= (p0 / (((p0 % (((p0 >> ((2) & 63))) | 1))) | 1))
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "foo bar"
    v3 := []string{("" + "foo bar"), "hello", v2, (("foo bar" + v2) + v2)}
    var v4 uint32 = 42
    var v10 uint32 = genFn0(182, 18446744073709551615, -891)
    if len(v3) > 0 {
        v3[0] = ""
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        if (4294967295 != genFn1(149, 42)) {
            var v12 bool = (false && (549 <= 32767))
            v1[v4] = ("a" + v2)
            _ = v12
        }
        _ = fi11
    }
    if (1.5 > 3.14159) {
        v13 := GenS1{F0: true, F1: 255}
        switch v2 {
        case "\n\t":
            v10 -= 4294967295
        }
        _ = v13
    }
    if ((v10 < genFn0(127, 757, 1)) && (2 >= 18446744073709551615)) {
        v2 = ("hello" + "foo bar")
        v4 = v10
    } else {
        fmt.Println(v2, v10)
        v14 := GenS3{F0: 588, F1: v2}
        _ = v14
    }
    if true {
        if false {
            if true {
                if (!false) {
                    var v15 int = -2147483648
                    delete(v1, (v4 + v10))
                    for fi16 := 0; fi16 < 3; fi16++ {
                        var v17 uint32 = genFn0(42, 18446744073709551615, -2147483648)
                        _ = fi16
                        _ = v17
                    }
                    _ = v15
                }
            }
        } else {
            v3 = append(v3, (v2 + v2))
            for fi18 := 0; fi18 < 6; fi18++ {
                for fi19 := 0; fi19 < 17; fi19++ {
                    if false {
                        fmt.Println(fi19, fi18, fi18)
                    } else {
                        fmt.Println(v2, fi18)
                        v2 = v2
                        var v20 int = (fi18 | 127)
                        _ = v20
                    }
                    _ = fi19
                }
                var v21 uint16 = 300
                if (v21 < (^(v21))) {
                    v10 = v4
                    for fi22 := 0; fi22 < 10; fi22++ {
                        var v23 int8 = -1
                        v21 += 65535
                        _ = fi22
                        _ = v23
                    }
                } else {
                    if true {
                        for fi24 := 0; fi24 < 14; fi24++ {
                            var v25 bool = false
                            v26 := GenS1{F0: v25, F1: 0}
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        if (3.14159 >= float32(-2.25)) {
                            for fi27 := 0; fi27 < 7; fi27++ {
                                if true {
                                    var v28 uint32 = (^(v10))
                                    v29 := GenS4{F0: v10}
                                    _ = v28
                                    _ = v29
                                }
                                switch v2 {
                                case "\n\t":
                                    if false {
                                        switch v2 {
                                        case "\n\t":
                                            switch (^(fi27)) {
                                            case 128:
                                                var v30 float64 = float64(3.14159)
                                                _ = v30
                                            }
                                            v4 = v10
                                            v1[v4] = "foo bar"
                                            var v31 int16 = -1
                                            switch "a" {
                                            case "hello":
                                                var v32 uint16 = v21
                                                _ = v32
                                            case "\n\t":
                                                v33 := GenS3{F0: genFn1(255, (-(v31))), F1: v2}
                                                _ = v33
                                            case "a":
                                                switch 4294967295 {
                                                case 32768:
                                                    if false {
                                                        var v34 int = (fi27 >> ((379) & 63))
                                                        _ = v34
                                                    }
                                                case 1:
                                                    v35 := GenS5{F0: 2, F1: 2147483647}
                                                    _ = v35
                                                case 4294967295:
                                                    if false {
                                                        v21 = (^(v21))
                                                        fmt.Println(v31, v21, v2)
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = ("foo bar" + v2)
                                                        }
                                                        if (false || true) {
                                                            fmt.Println(v2, fi27, fi27)
                                                            var v36 int8 = 1
                                                            v37 := v1[(v4 + v10)]
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        var v38 uint64 = 42
                                                        _ = v38
                                                    }
                                                default:
                                                    if false {
                                                        fmt.Println(v2, fi27, v31)
                                                        switch 4294967295 {
                                                        case 32768:
                                                            v21 = v21
                                                        case 9223372036854775807:
                                                            if (((1e10 / float64(((0.1 + ((((float32((-2.25 / 1.5)) + -2.25) * 3.14159) * ((-2.25 - ((-2.25 * 1.5) * (0.0 - (0.1 * 0.0)))) * -2.25)) - 1e10)) - -2.25))) / 1.0) == (0.1 + ((((-2.25 - 1.5) / 1.0) / (1.5 * 3.14159)) * 0.0))) {
                                                                _ = len(v1)
                                                                v2 = ("" + (v2 + (v2 + (v2 + v2))))
                                                                v10 += 65535
                                                            }
                                                        }
                                                        v21 -= (v21 ^ (v21 * v21))
                                                    }
                                                }
                                            default:
                                                if len(v3) > 0 {
                                                    v3[0] = v2
                                                }
                                            }
                                            switch -325 {
                                            case 127:
                                                v39 := GenS2{F0: 1}
                                                switch fi18 {
                                                case 32767:
                                                    v40 := GenS6{F0: -183, F1: -1}
                                                    _ = v40
                                                }
                                                _ = v39
                                            case -425:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            v10 -= (v4 | (v10 % (((v10 + v10)) | 1)))
                                            var v41 int16 = int16(fi18)
                                            _ = v31
                                            _ = v41
                                        default:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        v10 = v4
                                        v1[(^(v4))] = ""
                                    }
                                    if true {
                                        v21 += 42
                                        v2 = (v2 + v2)
                                        v2 = "hello"
                                    } else {
                                        if false {
                                            v42 := v1[(v10 ^ genFn0(127, 2, 42))]
                                            _ = v42
                                        } else {
                                            v21 -= v21
                                            v4 += 2147483648
                                            _ = len(v3)
                                        }
                                        switch 1 {
                                        case 2147483647:
                                            switch v2 {
                                            case "":
                                                v3 = append(v3, ("hello" + v2))
                                            default:
                                                v43 := GenS7{F0: 550}
                                                _ = v43
                                            }
                                            if (-116 != -1) {
                                                var v44 uint64 = 0
                                                v45 := v1[v10]
                                                _ = v44
                                                _ = v45
                                            } else {
                                                var v46 int8 = -1
                                                _ = v46
                                            }
                                            v47 := GenS5{F0: genFn0(128, uint64(v10), 2147483647), F1: -2147483648}
                                            if len(v3) > 0 {
                                                v3[0] = v2
                                            }
                                            _ = v47
                                        }
                                    }
                                case "a":
                                    _ = len(v1)
                                case "hello":
                                    fmt.Println(v21, v2, v10)
                                default:
                                    var v48 int = fi18
                                    _ = v48
                                }
                                var v49 bool = false
                                v49 = true
                                _ = fi27
                                _ = v49
                            }
                            if false {
                                v50 := GenS4{F0: (v10 - (v4 / ((2147483647) | 1)))}
                                for fi51 := 0; fi51 < 16; fi51++ {
                                    v52 := GenS2{F0: 1}
                                    switch ("foo bar" + "\n\t") {
                                    case "\n\t":
                                        var v53 int16 = 853
                                        _ = v53
                                    case "hello":
                                        v2 = ""
                                    default:
                                        switch fi51 {
                                        case -2147483648:
                                            _ = len(v3)
                                        case -1:
                                            v54 := GenS8{F0: "a"}
                                            _ = v54
                                        case 127:
                                            if (((0.0 + 3.14159) + -2.25) > 0.0) {
                                                v1[(v10 ^ v10)] = ((v2 + ((v2 + v2) + v2)) + v2)
                                                if ((true && ((v21 / ((v21) | 1)) > v21)) && (824 != 42)) {
                                                    _ = len(v1)
                                                    switch 9223372036854775807 {
                                                    case 0:
                                                        _ = len(v3)
                                                    default:
                                                        v2 = "hello"
                                                    }
                                                    v1[v4] = ((("" + (v2 + "foo bar")) + v2) + v2)
                                                }
                                            } else {
                                                if true {
                                                    v55 := GenS9{F0: -103}
                                                    v4 += v4
                                                    var v56 int64 = -1
                                                    v57 := GenS10{F0: 1000, F1: genFn0(2, 2, 2147483647)}
                                                    switch (fi51 & -367) {
                                                    case 128:
                                                        var v58 uint16 = 672
                                                        _ = v58
                                                    case -221:
                                                        var v59 uint = 4294967295
                                                        _ = v59
                                                    case -347:
                                                        var v60 uint64 = 0
                                                        _ = v60
                                                    }
                                                    _ = v55
                                                    _ = v56
                                                    _ = v57
                                                }
                                                fmt.Println(v21, v10)
                                            }
                                        default:
                                            v61 := GenS5{F0: (v10 << ((genFn1(127, 42)) & 63)), F1: -1}
                                            _ = v61
                                        }
                                    }
                                    var v62 int = int(v10)
                                    _ = fi51
                                    _ = v52
                                    _ = v62
                                }
                                for fi63 := 0; fi63 < 14; fi63++ {
                                    fmt.Println(fi18)
                                    _ = fi63
                                }
                                _ = v50
                            }
                            v1[(^(v4))] = v2
                        } else {
                            if false {
                                if false {
                                    switch "a" {
                                    case "a":
                                        switch 128 {
                                        case 127:
                                            v64 := GenS4{F0: 2147483647}
                                            _ = v64
                                        case -1:
                                            v65 := GenS7{F0: 2}
                                            _ = v65
                                        case -2147483648:
                                            switch fi18 {
                                            case -9223372036854775808:
                                                var v66 uint32 = v4
                                                _ = v66
                                            case 42:
                                                if true {
                                                    v2 = "a"
                                                } else {
                                                    fmt.Println(v4)
                                                    v3 = append(v3, "\n\t")
                                                    if true {
                                                        v3 = append(v3, v2)
                                                        v3 = append(v3, "\n\t")
                                                    }
                                                }
                                            default:
                                                if (!(!(1 == 147))) {
                                                    if len(v3) > 0 {
                                                        v3[0] = v2
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = v2
                                                    }
                                                    for fi67 := 0; fi67 < 16; fi67++ {
                                                        if true {
                                                            if ((1 != 42) && false) {
                                                                fmt.Println(v4)
                                                            }
                                                        }
                                                        v2 = v2
                                                        v10 -= (v4 * (v4 << ((4294967295) & 63)))
                                                        v2 = v2
                                                        if false {
                                                            _ = len(v3)
                                                            v3 = append(v3, (v2 + v2))
                                                            v2 = v2
                                                        }
                                                        _ = fi67
                                                    }
                                                } else {
                                                    v1[905] = v2
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        for fi68 := 0; fi68 < 7; fi68++ {
                                                            if false {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            } else {
                                                                v4 -= (v4 << ((884) & 63))
                                                            }
                                                            _ = fi68
                                                        }
                                                    default:
                                                        v10 = genFn0(127, 4294967295, -1)
                                                    }
                                                    v4 -= 65535
                                                }
                                            }
                                        }
                                    default:
                                        v2 = "hello"
                                    }
                                    if false {
                                        switch ("" + "\n\t") {
                                        case "":
                                            fmt.Println(fi18)
                                        default:
                                            if ((v21 >> ((65535) & 63)) == (v21 / ((144) | 1))) {
                                                v1[v4] = (v2 + (v2 + v2))
                                            } else {
                                                v69 := v1[v4]
                                                for fi70 := 0; fi70 < 7; fi70++ {
                                                    var v71 int32 = -1
                                                    v2 = ((v2 + "\n\t") + "\n\t")
                                                    _ = fi70
                                                    _ = v71
                                                }
                                                _ = v69
                                            }
                                        }
                                        _ = len(v1)
                                    } else {
                                        for fi72 := 0; fi72 < 3; fi72++ {
                                            v1[v10] = "foo bar"
                                            _ = fi72
                                        }
                                        fmt.Println(fi18)
                                        switch 32767 {
                                        case 838:
                                            v2 = ""
                                        }
                                    }
                                    if (!(0 >= 0)) {
                                        var v73 string = ((("" + v2) + v2) + v2)
                                        fmt.Println(fi18, v21, v4)
                                        var v74 float64 = 1.5
                                        _ = v73
                                        _ = v74
                                    }
                                } else {
                                    v10 += (v4 * (v10 + (v10 & genFn0(237, 135, 0))))
                                    delete(v1, 0)
                                }
                            } else {
                                if false {
                                    var v75 int8 = 127
                                    _ = v75
                                } else {
                                    v76 := GenS11{F0: 0, F1: 121, F2: 0}
                                    _ = v76
                                }
                                v4 -= uint32(v21)
                            }
                        }
                        if true {
                            for fi77 := 0; fi77 < 3; fi77++ {
                                fmt.Println(v10, v2)
                                v21 -= v21
                                if true {
                                    if false {
                                        v78 := GenS1{F0: true, F1: uint8(v21)}
                                        if ((v21 ^ v21) <= 65535) {
                                            var v79 int32 = -190
                                            v4 = 0
                                            var v80 uint64 = 42
                                            v81 := GenS8{F0: (v2 + v2)}
                                            v80 -= v80
                                            switch -731 {
                                            case 32768:
                                                var v82 uint32 = (v10 - v10)
                                                _ = v82
                                            case 128:
                                                var v83 float64 = float64(((0.1 + ((3.14159 * 0.1) + 3.14159)) * (-2.25 * (float32((-2.25 * (-2.25 * ((float64(1e10) - (1.5 * 1.5)) / ((((0.0 / ((-2.25 - ((0.0 - float64(1e10)) - 1e10)) * 1.5)) / 1.5) - 1.5) + 0.0))))) + float32((1.5 / 1e10))))))
                                                _ = v83
                                            }
                                            _ = v79
                                            _ = v80
                                            _ = v81
                                        } else {
                                            v84 := GenS8{F0: v2}
                                            v85 := GenS12{F0: 65535, F1: 335, F2: 1}
                                            _ = v84
                                            _ = v85
                                        }
                                        _ = v78
                                    }
                                    v21 = 255
                                }
                                _ = fi77
                            }
                        } else {
                            v86 := GenS13{F0: 255}
                            fmt.Println(v10, v4)
                            _ = v86
                        }
                    } else {
                        switch v2 {
                        case "hello":
                            if false {
                                if true {
                                    switch (v2 + v2) {
                                    case "a":
                                        switch (^(fi18)) {
                                        case 127:
                                            v87 := GenS14{F0: ""}
                                            _ = v87
                                        case 32768:
                                            var v88 float64 = (1.5 - ((1e10 * 1.5) * (float64((0.0 * 0.1)) + ((-2.25 - 0.1) + ((((float64(0.0) + 1.5) * -2.25) - float64(float32(3.14159))) * (0.1 / ((1e10 / 1.5) * float64(-2.25))))))))
                                            _ = v88
                                        case 474:
                                            for fi89 := 0; fi89 < 14; fi89++ {
                                                fmt.Println(v4, fi89, fi89)
                                                var v90 float64 = 3.14159
                                                v91 := GenS9{F0: 0}
                                                _ = fi89
                                                _ = v90
                                                _ = v91
                                            }
                                        default:
                                            v21 = v21
                                        }
                                    case "foo bar":
                                        var v92 uint64 = 4294967295
                                        _ = v92
                                    case "":
                                        fmt.Println(v21)
                                    default:
                                        fmt.Println(v2)
                                    }
                                }
                                if (4294967295 == 42) {
                                    v4 = 2147483648
                                    v93 := GenS15{F0: 0, F1: (fi18 + fi18), F2: 1000}
                                    _ = v93
                                }
                                switch v2 {
                                case "hello":
                                    v4 = (^(v4))
                                case "foo bar":
                                    if false {
                                        v10 += v10
                                    }
                                case "":
                                    var v94 int64 = -787
                                    _ = v94
                                default:
                                    if ((131 <= 551) && false) {
                                        var v95 uint16 = 73
                                        var v96 uint16 = (v95 << ((974) & 63))
                                        _ = v95
                                        _ = v96
                                    } else {
                                        var v97 int = (fi18 / ((-512) | 1))
                                        _ = len(v3)
                                        for fi98 := 0; fi98 < 13; fi98++ {
                                            v97 = fi18
                                            switch (v97 | (fi18 | (^(v97)))) {
                                            case 127:
                                                var v99 int16 = 1000
                                                _ = v99
                                            case -332:
                                                switch 1 {
                                                case -9223372036854775808:
                                                    v21 -= v21
                                                case -2147483648:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                default:
                                                    if false {
                                                        v4 -= (v4 % ((uint32(v21)) | 1))
                                                    } else {
                                                        v2 = v2
                                                    }
                                                }
                                            case 0:
                                                if false {
                                                    if (-1 > 42) {
                                                        if (345 <= 0) {
                                                            var v100 float32 = 3.14159
                                                            v2 = "\n\t"
                                                            _ = v100
                                                        }
                                                        fmt.Println(v97, v10)
                                                    } else {
                                                        v101 := GenS16{F0: 2, F1: 18446744073709551615}
                                                        v102 := GenS12{F0: genFn1(167, 238), F1: 42, F2: 2}
                                                        _ = v101
                                                        _ = v102
                                                    }
                                                    if (-2.25 <= 0.0) {
                                                        v21 = 42
                                                        v103 := GenS17{F0: 977}
                                                        fmt.Println(v4)
                                                        if false {
                                                            v104 := GenS8{F0: v2}
                                                            delete(v1, genFn0(1, 673, -1))
                                                            fmt.Println(fi18, v21, v97)
                                                            fmt.Println(fi18)
                                                            _ = v104
                                                        } else {
                                                            v21 = 255
                                                            if (0 == 255) {
                                                                switch "\n\t" {
                                                                case "":
                                                                    v97 += fi98
                                                                default:
                                                                    v2 = v2
                                                                }
                                                            } else {
                                                                v105 := GenS18{F0: v2, F1: -1, F2: int16(v97)}
                                                                if (!(!(v21 == v21))) {
                                                                    switch (fi98 + (fi98 * 0)) {
                                                                    case 641:
                                                                        v106 := GenS18{F0: (v2 + v2), F1: 9223372036854775807, F2: 541}
                                                                        var v107 float64 = (3.14159 + ((1e10 / 1.5) / float64(0.1)))
                                                                        _ = v106
                                                                        _ = v107
                                                                    case 42:
                                                                        switch (v97 & (fi98 | (fi18 / (((v97 << ((0) & 63))) | 1)))) {
                                                                        case 2147483647:
                                                                            if false {
                                                                                fmt.Println(v21, v4, fi98)
                                                                                v108 := GenS18{F0: ((v2 + v2) + (v2 + v2)), F1: 1, F2: 1000}
                                                                                switch fi18 {
                                                                                case -45:
                                                                                    v21 -= v21
                                                                                default:
                                                                                    if (0 < 2147483647) {
                                                                                        v21 = (v21 & 65535)
                                                                                        v3 = append(v3, ("\n\t" + v2))
                                                                                    } else {
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = v2
                                                                                        }
                                                                                        if true {
                                                                                            v4 += genFn0(3, 0, 0)
                                                                                            v1[genFn0(0, 1, 746)] = "a"
                                                                                            v1[uint32(v21)] = (v2 + "\n\t")
                                                                                        } else {
                                                                                            if false {
                                                                                                v4 = 732
                                                                                                if true {
                                                                                                    v109 := GenS19{F0: uint(v10), F1: 9223372036854775807, F2: (v2 + v2)}
                                                                                                    _ = v109
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v108
                                                                            }
                                                                        default:
                                                                            v110 := GenS11{F0: 42, F1: (v4 & v4), F2: 980}
                                                                            _ = v110
                                                                        }
                                                                        var v111 float32 = 1.5
                                                                        var v112 int16 = 32767
                                                                        _ = v111
                                                                        _ = v112
                                                                    default:
                                                                        if true {
                                                                            delete(v1, 2)
                                                                            var v113 int8 = 0
                                                                            v114 := GenS9{F0: 127}
                                                                            _ = v113
                                                                            _ = v114
                                                                        } else {
                                                                            var v115 int = (fi98 ^ fi98)
                                                                            var v116 int8 = int8(v115)
                                                                            _ = v115
                                                                            _ = v116
                                                                        }
                                                                        var v117 uint = 42
                                                                        _ = v117
                                                                    }
                                                                }
                                                                _ = v105
                                                            }
                                                            v10 = (v10 / ((v10) | 1))
                                                        }
                                                        fmt.Println(v10, v2)
                                                        _ = v103
                                                    } else {
                                                        if (!false) {
                                                            if true {
                                                                if (1 <= 0) {
                                                                    fmt.Println(fi18, fi18, v10)
                                                                }
                                                                if (!(true || (!true))) {
                                                                    var v118 float64 = float64(0.1)
                                                                    switch v2 {
                                                                    case "\n\t":
                                                                        var v119 uint32 = (v4 % (((v10 & (v4 & v4))) | 1))
                                                                        _ = v119
                                                                    case "hello":
                                                                        switch 9223372036854775807 {
                                                                        case 850:
                                                                            fmt.Println(fi18, v118, v97)
                                                                        case 32768:
                                                                            if len(v3) > 0 {
                                                                                v3[0] = v2
                                                                            }
                                                                        default:
                                                                            v120 := GenS15{F0: 1000, F1: 0, F2: 1000}
                                                                            _ = v120
                                                                        }
                                                                    default:
                                                                        var v121 int16 = 937
                                                                        _ = v121
                                                                    }
                                                                    if false {
                                                                        v4 -= (v10 << ((1) & 63))
                                                                        _ = len(v1)
                                                                        fmt.Println(fi18, v10, v2)
                                                                    } else {
                                                                        v122 := GenS15{F0: -333, F1: fi98, F2: 32767}
                                                                        var v123 uint = 231
                                                                        _ = v122
                                                                        _ = v123
                                                                    }
                                                                    _ = v118
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        fmt.Println(fi18, v4, v21)
                                                                    case "\n\t":
                                                                        fmt.Println(v21, fi18, v21)
                                                                    case "a":
                                                                        fmt.Println(v10, v4)
                                                                    default:
                                                                        var v124 uint = 42
                                                                        _ = v124
                                                                    }
                                                                    fmt.Println(v10, fi18)
                                                                }
                                                            }
                                                            v10 += (v10 << ((408) & 63))
                                                            v125 := GenS20{F0: 116, F1: false, F2: -128}
                                                            _ = v125
                                                        }
                                                    }
                                                } else {
                                                    switch v2 {
                                                    case "\n\t":
                                                        var v126 int16 = 1
                                                        _ = v126
                                                    default:
                                                        if (0.1 > (1e10 / (-2.25 / float32(1e10)))) {
                                                            v97 = (v97 & fi18)
                                                        } else {
                                                            v127 := GenS21{F0: (v21 * (^(v21)))}
                                                            _ = v127
                                                        }
                                                    }
                                                    delete(v1, 1)
                                                    v3 = append(v3, "a")
                                                }
                                            default:
                                                delete(v1, (v4 + (^(v4))))
                                            }
                                            switch 42 {
                                            case 4294967295:
                                                v128 := GenS14{F0: (v2 + v2)}
                                                _ = v128
                                            case -2147483648:
                                                var v129 uint16 = 255
                                                _ = v129
                                            default:
                                                switch ((v2 + ("\n\t" + v2)) + v2) {
                                                case "a":
                                                    var v130 uint8 = 0
                                                    _ = v130
                                                case "\n\t":
                                                    v131 := GenS22{F0: 4294967295, F1: 2147483647}
                                                    _ = v131
                                                }
                                            }
                                            _ = fi98
                                        }
                                        _ = v97
                                    }
                                }
                            }
                        case "foo bar":
                            for fi132 := 0; fi132 < 9; fi132++ {
                                if (1 != 255) {
                                    var v133 float32 = 3.14159
                                    var v134 int64 = -1
                                    if (!true) {
                                        v135 := GenS9{F0: -1}
                                        _ = len(v1)
                                        _ = v135
                                    } else {
                                        var v136 int64 = 1
                                        v137 := GenS23{F0: v21}
                                        _ = v136
                                        _ = v137
                                    }
                                    _ = v133
                                    _ = v134
                                } else {
                                    v138 := v1[(v10 ^ v10)]
                                    v1[2147483647] = ""
                                    _ = v138
                                }
                                v139 := GenS24{F0: (v4 | (v10 ^ 2)), F1: -2147483648, F2: 452}
                                _ = fi132
                                _ = v139
                            }
                        default:
                            var v140 string = v2
                            _ = v140
                        }
                    }
                    v141 := v1[genFn0(180, 1, 8)]
                    fmt.Println(v21, v141)
                    _ = v141
                }
                _ = fi18
                _ = v21
            }
            v142 := GenS8{F0: "hello"}
            _ = v142
        }
        switch ((v2 + v2) + v2) {
        case "hello":
            v4 -= (v10 / ((v10) | 1))
        case "a":
            switch "foo bar" {
            case "hello":
                v2 = "foo bar"
            }
        }
        var v143 int32 = 698
        _ = v143
    } else {
        for fi144 := 0; fi144 < 7; fi144++ {
            delete(v1, genFn0(uint8(v4), 0, -317))
            _ = fi144
        }
        if false {
            var v145 int64 = 1
            v145 += v145
            _ = v145
        }
    }
    if true {
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch "" {
            case "":
                v4 = 348
            }
        }
        v146 := GenS2{F0: 0}
        _ = v146
    }
    switch -26 {
    case 128:
        var v147 int16 = 550
        _ = v147
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
