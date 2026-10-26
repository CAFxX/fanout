package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint; F2 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int64; F1 int8 }
type GenS12 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint16; F1 bool }
type GenS18 struct { F0 uint; F1 bool; F2 uint32 }
type GenS19 struct { F0 int32; F1 uint }
type GenS20 struct { F0 int16; F1 bool }
type GenS21 struct { F0 int16; F1 int32 }
type GenS22 struct { F0 int32 }
func genFn0(p0 uint32, p1 uint8, p2 uint64) uint {
    for fi7 := 0; fi7 < 9; fi7++ {
        p2 -= 42
        p1 -= 1
        _ = fi7
    }
    p0 = (p0 * (p0 - (p0 | p0)))
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 int, p1 uint8) uint32 {
    var v8 string = "a"
    v9 := []int{4294967295}
    if (-128 != 0) {
        switch -50 {
        case 170:
            switch "" {
            case "":
                p1 -= p1
            case "hello":
                for fi10 := 0; fi10 < 11; fi10++ {
                    if ((!(-882 <= -1)) || ((-1 >= (fi10 & 127)) && (!(2 < 42)))) {
                        fmt.Println(p1, v8, p1)
                    }
                    switch (v8 + v8) {
                    case "":
                        var v11 float64 = float64(((((0.1 / (-2.25 * 1.5)) + -2.25) / -2.25) / 1e10))
                        _ = v11
                    default:
                        var v12 int = 2147483647
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 5; fi13++ {
                        fmt.Println(v8, fi13, fi10)
                        _ = fi13
                    }
                    _ = fi10
                }
            }
        case -1:
            var v14 int = -802
            _ = v14
        default:
            p0 = p0
        }
        for fi15 := 0; fi15 < 19; fi15++ {
            if (400 < -1) {
                var v16 int = 740
                _ = v16
            } else {
                fmt.Println(fi15)
            }
            if false {
                p1 -= (p1 << ((255) & 63))
            }
            _ = fi15
        }
        if len(v9) > 0 {
            v9[0] = (p0 * p0)
        }
    }
    if (!true) {
        if false {
            fmt.Println(v8, p1)
            _ = len(v9)
            fmt.Println(p0)
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        if true {
            if false {
                var v17 uint32 = 2
                if (true || false) {
                    var v18 int32 = 2147483647
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = v18
                } else {
                    var v19 float32 = 3.14159
                    switch (v8 + (v8 + v8)) {
                    case "hello":
                        v19 -= (float32(1.5) / v19)
                    default:
                        var v20 float64 = float64(((((1.5 / ((float32((-2.25 * 3.14159)) * (1e10 / ((v19 / (1.5 * (0.1 - (1.5 + v19)))) / v19))) - v19)) + float32(0.0)) - v19) * v19))
                        _ = v20
                    }
                    switch 2147483647 {
                    case 127:
                        v19 += v19
                    case -9223372036854775808:
                        if false {
                            for fi21 := 0; fi21 < 6; fi21++ {
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                for fi22 := 0; fi22 < 9; fi22++ {
                                    if false {
                                        fmt.Println(v8)
                                        if (-23 == -99) {
                                            fmt.Println(fi21, p0)
                                            fmt.Println(v17, v8)
                                            v23 := GenS1{F0: 425}
                                            _ = v23
                                        } else {
                                            _ = len(v9)
                                            if ((true || true) && ((65535 != 42) && (42 > 65535))) {
                                                var v24 uint64 = 605
                                                var v25 int32 = 1
                                                _ = v24
                                                _ = v25
                                            }
                                            switch fi21 {
                                            case -1:
                                                var v26 string = ((((v8 + "hello") + (v8 + "\n\t")) + v8) + v8)
                                                _ = v26
                                            case 9223372036854775807:
                                                var v27 uint32 = v17
                                                _ = v27
                                            default:
                                                if false {
                                                    var v28 float64 = 0.1
                                                    if true {
                                                        var v29 uint32 = v17
                                                        p0 = (p0 & 9223372036854775807)
                                                        fmt.Println(v19, v8)
                                                        if (v28 != v28) {
                                                            v30 := GenS2{F0: 1, F1: (fi22 ^ 0), F2: 889}
                                                            v31 := GenS3{F0: 42, F1: 42, F2: 1000}
                                                            _ = v30
                                                            _ = v31
                                                        } else {
                                                            switch v8 {
                                                            case "foo bar":
                                                                switch "\n\t" {
                                                                case "":
                                                                    var v32 int64 = 703
                                                                    _ = v32
                                                                case "foo bar":
                                                                    p0 -= (p0 | (fi22 * 482))
                                                                case "hello":
                                                                    v33 := GenS2{F0: 1, F1: p0, F2: 65535}
                                                                    _ = v33
                                                                default:
                                                                    v17 = v29
                                                                }
                                                            case "\n\t":
                                                                if len(v9) >= 2 {
                                                                    v9 = v9[1:len(v9)]
                                                                }
                                                            case "hello":
                                                                if len(v9) >= 2 {
                                                                    v9 = v9[1:len(v9)]
                                                                }
                                                            }
                                                            if true {
                                                                switch 32767 {
                                                                case 128:
                                                                    if len(v9) >= 2 {
                                                                        v9 = v9[1:len(v9)]
                                                                    }
                                                                    v28 -= ((v28 * (-2.25 + 0.0)) + (v28 - 1.5))
                                                                    v9 = append(v9, fi21)
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        if true {
                                                                            var v34 uint32 = (^(v29))
                                                                            _ = v34
                                                                        } else {
                                                                            var v35 uint16 = 255
                                                                            v36 := GenS3{F0: v35, F1: 65535, F2: -32768}
                                                                            _ = v35
                                                                            _ = v36
                                                                        }
                                                                    case "a":
                                                                        if false {
                                                                            _ = len(v9)
                                                                            switch ("foo bar" + v8) {
                                                                            case "foo bar":
                                                                                p0 -= (fi22 & fi22)
                                                                            case "":
                                                                                switch fi22 {
                                                                                case 32768:
                                                                                    if len(v9) >= 2 {
                                                                                        v9 = v9[1:len(v9)]
                                                                                    }
                                                                                case 4294967295:
                                                                                    fmt.Println(fi22)
                                                                                    switch -966 {
                                                                                    case -2147483648:
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            if (4294967295 < (^(v29))) {
                                                                                                switch -1 {
                                                                                                case 127:
                                                                                                    if len(v9) >= 2 {
                                                                                                        v9 = v9[1:len(v9)]
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                var v37 string = "a"
                                                                                                var v38 uint8 = (p1 << ((467) & 63))
                                                                                                _ = v37
                                                                                                _ = v38
                                                                                            }
                                                                                        case "hello":
                                                                                            v8 = "\n\t"
                                                                                        }
                                                                                    default:
                                                                                        var v39 int = (fi22 + 128)
                                                                                        _ = v39
                                                                                    }
                                                                                    fmt.Println(v29, p1, p0)
                                                                                    v40 := GenS4{F0: 42}
                                                                                    var v41 int8 = -128
                                                                                    v42 := GenS5{F0: p1}
                                                                                    var v43 float64 = ((((float64(v19) * ((v28 / (-2.25 + v28)) + v28)) + (v28 - v28)) - 3.14159) * v28)
                                                                                    v9 = append(v9, -882)
                                                                                    _ = v40
                                                                                    _ = v41
                                                                                    _ = v42
                                                                                    _ = v43
                                                                                }
                                                                            default:
                                                                                v19 = float32(float64((v19 / float32(v28))))
                                                                            }
                                                                        }
                                                                    case "":
                                                                        var v44 string = ("foo bar" + v8)
                                                                        _ = v44
                                                                    }
                                                                    switch v8 {
                                                                    case "":
                                                                        p1 -= (p1 & p1)
                                                                    case "foo bar":
                                                                        p0 = (p0 - 0)
                                                                    default:
                                                                        var v45 uint8 = p1
                                                                        _ = v45
                                                                    }
                                                                    if len(v9) >= 2 {
                                                                        v9 = v9[1:len(v9)]
                                                                    }
                                                                    fmt.Println(v29)
                                                                    switch (v8 + "") {
                                                                    case "\n\t":
                                                                        if (966 < 0) {
                                                                            var v46 uint8 = (p1 % ((38) | 1))
                                                                            _ = len(v9)
                                                                            var v47 uint8 = 127
                                                                            if len(v9) > 0 {
                                                                                v9[0] = fi21
                                                                            }
                                                                            _ = v46
                                                                            _ = v47
                                                                        }
                                                                    case "a":
                                                                        var v48 int16 = -478
                                                                        _ = v48
                                                                    case "foo bar":
                                                                        v28 -= float64((float32(v28) / (((float32((((1.5 + v28) * v28) / v28)) + (v19 - float32(1.5))) * -2.25) * 0.1)))
                                                                        v29 = 679
                                                                    }
                                                                    v29 += v17
                                                                    v19 += 3.14159
                                                                case 42:
                                                                    if (65535 == 42) {
                                                                        switch v8 {
                                                                        case "foo bar":
                                                                            _ = len(v9)
                                                                        case "":
                                                                            p0 -= p0
                                                                        case "hello":
                                                                            if (65535 == 314) {
                                                                                var v49 uint = genFn0((v17 % (((v29 ^ 42)) | 1)), 2, 0)
                                                                                var v50 int64 = -1
                                                                                v19 = (3.14159 * ((-2.25 * (-2.25 * 0.0)) / 0.1))
                                                                                _ = v49
                                                                                _ = v50
                                                                            }
                                                                        default:
                                                                            var v51 float32 = float32((float64(-2.25) / (0.0 - -2.25)))
                                                                            _ = v51
                                                                        }
                                                                        v19 = 3.14159
                                                                        v28 = 1e10
                                                                        v29 += v17
                                                                    } else {
                                                                        p1 = p1
                                                                        fmt.Println(v19, v17)
                                                                        switch v8 {
                                                                        case "":
                                                                            _ = len(v9)
                                                                        default:
                                                                            var v52 int8 = 1
                                                                            _ = v52
                                                                        }
                                                                    }
                                                                default:
                                                                    switch (fi22 / (((p0 * -9223372036854775808)) | 1)) {
                                                                    case -9223372036854775808:
                                                                        var v53 uint16 = 208
                                                                        var v54 string = ("a" + "foo bar")
                                                                        var v55 int = -9223372036854775808
                                                                        _ = v53
                                                                        _ = v54
                                                                        _ = v55
                                                                    case -918:
                                                                        if ((true && (814 == 32767)) || true) {
                                                                            fmt.Println(v19)
                                                                            fmt.Println(p1)
                                                                        }
                                                                    default:
                                                                        fmt.Println(v28)
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v29
                                                    } else {
                                                        v19 += ((v19 / v19) + v19)
                                                        var v56 uint16 = 432
                                                        _ = v56
                                                    }
                                                    p1 += p1
                                                    _ = v28
                                                }
                                            }
                                        }
                                        v17 = (v17 | v17)
                                    }
                                    if (0 != 0) {
                                        _ = len(v9)
                                        fmt.Println(p0, fi22)
                                    } else {
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                    }
                                    _ = fi22
                                }
                                var v57 uint16 = 32767
                                _ = fi21
                                _ = v57
                            }
                        }
                    case 106:
                        if len(v9) > 0 {
                            v9[0] = (-(p0))
                        }
                    }
                    _ = v19
                }
                if (1 >= -32768) {
                    v58 := GenS6{F0: v8}
                    if (472 >= 42) {
                        _ = len(v9)
                    } else {
                        if false {
                            if ((p0 < p0) && false) {
                                if (2 == 0) {
                                    v9 = append(v9, p0)
                                }
                                p1 = p1
                                if ((p0 > 127) || (-128 == 100)) {
                                    for fi59 := 0; fi59 < 14; fi59++ {
                                        v17 -= v17
                                        _ = fi59
                                    }
                                    var v60 uint8 = 255
                                    _ = v60
                                } else {
                                    v8 = v8
                                    for fi61 := 0; fi61 < 5; fi61++ {
                                        switch v8 {
                                        case "":
                                            var v62 uint32 = 4294967295
                                            _ = v62
                                        case "foo bar":
                                            v17 -= 2147483647
                                        }
                                        v63 := GenS7{F0: 1000}
                                        var v64 uint = 0
                                        _ = fi61
                                        _ = v63
                                        _ = v64
                                    }
                                }
                            }
                        } else {
                            v65 := GenS8{F0: (v17 / ((143) | 1))}
                            if (!false) {
                                v8 = ("a" + v8)
                            } else {
                                if true {
                                    for fi66 := 0; fi66 < 11; fi66++ {
                                        p1 = p1
                                        _ = fi66
                                    }
                                }
                                if true {
                                    for fi67 := 0; fi67 < 9; fi67++ {
                                        p1 = 255
                                        var v68 int64 = -313
                                        fmt.Println(v17)
                                        _ = fi67
                                        _ = v68
                                    }
                                    v17 = v17
                                }
                            }
                            switch (("\n\t" + "hello") + v8) {
                            case "\n\t":
                                var v69 int64 = -9223372036854775808
                                _ = v69
                            case "hello":
                                if ((!true) || true) {
                                    var v70 bool = (v8 != "foo bar")
                                    if (18446744073709551615 >= 4294967295) {
                                        v9 = append(v9, (p0 % ((p0) | 1)))
                                    } else {
                                        var v71 uint = 4294967295
                                        _ = v71
                                    }
                                    for fi72 := 0; fi72 < 1; fi72++ {
                                        switch int(v17) {
                                        case -2147483648:
                                            var v73 string = v8
                                            _ = v73
                                        case -807:
                                            _ = len(v9)
                                        case 42:
                                            switch v8 {
                                            case "a":
                                                _ = len(v9)
                                            case "":
                                                var v74 uint = uint(p1)
                                                _ = v74
                                            case "\n\t":
                                                var v75 uint = 496
                                                _ = v75
                                            }
                                        default:
                                            if v70 {
                                                var v76 float64 = ((0.0 + -2.25) / 1.0)
                                                _ = len(v9)
                                                _ = v76
                                            }
                                        }
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                        _ = fi72
                                    }
                                    _ = v70
                                } else {
                                    switch (p0 ^ (p0 - p0)) {
                                    case -877:
                                        var v77 uint8 = (p1 << ((2) & 63))
                                        _ = v77
                                    case -167:
                                        v9 = append(v9, p0)
                                    case 4294967295:
                                        for fi78 := 0; fi78 < 10; fi78++ {
                                            switch "foo bar" {
                                            case "a":
                                                for fi79 := 0; fi79 < 10; fi79++ {
                                                    v8 = "\n\t"
                                                    fmt.Println(fi79, v17)
                                                    var v80 int8 = 1
                                                    v8 = (v8 + "foo bar")
                                                    fmt.Println(p0)
                                                    _ = fi79
                                                    _ = v80
                                                }
                                            case "\n\t":
                                                v8 = v8
                                            case "":
                                                p0 = p0
                                            default:
                                                if false {
                                                    var v81 bool = ((!false) || (501 > -2147483648))
                                                    fmt.Println(p1)
                                                    fmt.Println(fi78, v81)
                                                    _ = v81
                                                }
                                            }
                                            if true {
                                                for fi82 := 0; fi82 < 6; fi82++ {
                                                    p0 += -870
                                                    _ = fi82
                                                }
                                                v9 = append(v9, 32768)
                                                var v83 int64 = 42
                                                _ = v83
                                            }
                                            _ = fi78
                                        }
                                    default:
                                        var v84 int = (p0 >> ((42) & 63))
                                        _ = v84
                                    }
                                    for fi85 := 0; fi85 < 4; fi85++ {
                                        switch v8 {
                                        case "a":
                                            v86 := GenS9{F0: -2147483648}
                                            _ = v86
                                        }
                                        if len(v9) > 0 {
                                            v9[0] = (fi85 & int(p1))
                                        }
                                        _ = fi85
                                    }
                                    v9 = append(v9, (p0 << ((408) & 63)))
                                }
                            default:
                                p1 -= (p1 ^ p1)
                            }
                            _ = v65
                        }
                        var v87 int32 = 42
                        p0 -= 392
                        _ = v87
                    }
                    switch (p0 / ((p0) | 1)) {
                    case 940:
                        switch v8 {
                        case "hello":
                            var v88 uint = 18446744073709551615
                            _ = v88
                        case "":
                            if ((-(p0)) < (p0 << ((genFn0(v17, p1, 857)) & 63))) {
                                var v89 uint = 743
                                for fi90 := 0; fi90 < 14; fi90++ {
                                    if true {
                                        v91 := GenS10{F0: v17}
                                        _ = v91
                                    } else {
                                        p1 = (p1 >> ((v89) & 63))
                                    }
                                    _ = fi90
                                }
                                if ((0 > 1) || true) {
                                    v17 = v17
                                } else {
                                    var v92 int32 = 679
                                    v93 := GenS1{F0: 0}
                                    _ = v92
                                    _ = v93
                                }
                                _ = v89
                            } else {
                                var v94 uint8 = 128
                                _ = v94
                            }
                        default:
                            _ = len(v9)
                        }
                    default:
                        fmt.Println(p0, v17)
                    }
                    _ = v58
                } else {
                    if ((float64((float32(1e10) - ((0.0 / 1.5) * (1e10 + 1.5)))) + 0.0) <= (0.1 * 3.14159)) {
                        p1 = 0
                    }
                    var v95 int64 = -926
                    _ = v95
                }
                _ = v17
            } else {
                v96 := GenS2{F0: 32767, F1: p0, F2: 0}
                v8 = ((v8 + (("hello" + "a") + ("a" + v8))) + v8)
                _ = v96
            }
            var v97 uint16 = 65535
            _ = v97
        }
    } else {
        v98 := GenS10{F0: 1}
        _ = v98
    }
    v8 = v8
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 uint = 1
    v2 := make(map[uint32]string)
    var v3 string = "a"
    var v4 uint64 = 4294967295
    var v5 uint64 = v4
    var v6 int = -2147483648
    fmt.Println(v5)
    fmt.Println(v6)
    if (false || true) {
        if false {
            v99 := v2[2]
            _ = v99
        }
    } else {
        fmt.Println(v1)
        var v100 int8 = 42
        for fi101 := 0; fi101 < 7; fi101++ {
            v102 := []int{2147483647, (-(fi101)), fi101, (fi101 - 990)}
            if (-2.25 < 0.1) {
                if (!(((0.1 * 0.0) * 1e10) < 3.14159)) {
                    if true {
                        fmt.Println(v4)
                        if (!false) {
                            for fi103 := 0; fi103 < 7; fi103++ {
                                var v104 uint8 = 42
                                if true {
                                    if (32767 <= 260) {
                                        v102 = append(v102, fi103)
                                        if len(v102) > 0 {
                                            v102[0] = fi101
                                        }
                                        fmt.Println(v6, fi103, v100)
                                    } else {
                                        var v105 int32 = 2147483647
                                        _ = v105
                                    }
                                }
                                _ = fi103
                                _ = v104
                            }
                        } else {
                            switch (-(v6)) {
                            case -2147483648:
                                var v106 int8 = 0
                                _ = v106
                            default:
                                v107 := GenS10{F0: genFn1((v6 ^ (v6 & (fi101 - (fi101 / (((fi101 | 127)) | 1))))), 255)}
                                _ = v107
                            }
                            if (456 >= 1) {
                                v1 = v1
                                v108 := GenS7{F0: 42}
                                if len(v102) > 0 {
                                    v102[0] = (fi101 >> (((v1 % (((^(v1))) | 1))) & 63))
                                }
                                _ = v108
                            }
                        }
                        var v109 float32 = 1e10
                        _ = v109
                    }
                }
            }
            v110 := GenS11{F0: int64(v100), F1: -1}
            _ = fi101
            _ = v102
            _ = v110
        }
        _ = v100
    }
    if ((v1 % ((2) | 1)) < v1) {
        if true {
            v5 = (^(v4))
        } else {
            var v111 int64 = -9223372036854775808
            _ = v111
        }
    } else {
        for fi112 := 0; fi112 < 13; fi112++ {
            switch (v6 << ((65535) & 63)) {
            case -2147483648:
                switch -315 {
                case 155:
                    if true {
                        for fi113 := 0; fi113 < 13; fi113++ {
                            v114 := GenS11{F0: 1, F1: 127}
                            if ((fi113 >= fi112) && false) {
                                var v115 float32 = 0.1
                                var v116 float64 = (3.14159 + (0.0 + (1.5 * (float64(((((v115 / v115) / v115) + v115) - v115)) + -2.25))))
                                delete(v2, genFn1(v6, 42))
                                _ = v115
                                _ = v116
                            }
                            _ = fi113
                            _ = v114
                        }
                        v117 := GenS2{F0: 1, F1: fi112, F2: 558}
                        _ = v117
                    } else {
                        v118 := []int{-1}
                        _ = v118
                    }
                default:
                    switch v6 {
                    case 729:
                        fmt.Println(v5)
                    case 850:
                        switch ((v3 + v3) + (v3 + "a")) {
                        case "\n\t":
                            if (42 > (v1 % (((v1 ^ 666)) | 1))) {
                                switch ("foo bar" + "hello") {
                                case "hello":
                                    v1 += (v1 + v1)
                                case "\n\t":
                                    if (false && false) {
                                        v119 := GenS12{F0: 147, F1: -2147483648, F2: 1}
                                        var v120 uint = 4294967295
                                        if ("hello" <= v3) {
                                            for fi121 := 0; fi121 < 6; fi121++ {
                                                var v122 uint16 = 542
                                                _ = fi121
                                                _ = v122
                                            }
                                            var v123 int64 = -690
                                            _ = v123
                                        } else {
                                            var v124 float32 = 1.5
                                            _ = v124
                                        }
                                        _ = v119
                                        _ = v120
                                    } else {
                                        var v125 uint8 = 99
                                        for fi126 := 0; fi126 < 18; fi126++ {
                                            var v127 int32 = 2147483647
                                            v128 := []int{fi126, (-(fi112)), fi126, 2147483647}
                                            _ = fi126
                                            _ = v127
                                            _ = v128
                                        }
                                        _ = v125
                                    }
                                case "foo bar":
                                    switch (v6 | v6) {
                                    case -9223372036854775808:
                                        if ((v4 >= v4) || (0.0 > 1e10)) {
                                            for fi129 := 0; fi129 < 1; fi129++ {
                                                switch -1 {
                                                case 128:
                                                    v130 := GenS4{F0: (v1 | 42)}
                                                    _ = v130
                                                default:
                                                    var v131 uint32 = 4294967295
                                                    _ = v131
                                                }
                                                v6 += fi129
                                                v132 := GenS13{F0: (("\n\t" + "foo bar") + "a")}
                                                _ = fi129
                                                _ = v132
                                            }
                                            v4 -= 619
                                            v3 = "\n\t"
                                        } else {
                                            for fi133 := 0; fi133 < 14; fi133++ {
                                                if (351 < 0) {
                                                    var v134 uint64 = v4
                                                    if ((((-2.25 + 0.0) - 0.0) / (1e10 * -2.25)) == (0.0 / 1.0)) {
                                                        switch (("foo bar" + v3) + v3) {
                                                        case "":
                                                            var v135 int = v6
                                                            _ = v135
                                                        case "foo bar":
                                                            if false {
                                                                v2[2147483648] = (v3 + v3)
                                                                if (0 <= 2147483648) {
                                                                    if true {
                                                                        fmt.Println(v4, fi112, v6)
                                                                    }
                                                                    fmt.Println(v6, v5)
                                                                    switch (fi133 % ((-9223372036854775808) | 1)) {
                                                                    case 400:
                                                                        v136 := GenS4{F0: 65535}
                                                                        _ = v136
                                                                    case 128:
                                                                        if false {
                                                                            v3 = v3
                                                                            var v137 float64 = (0.1 / 1.0)
                                                                            v138 := GenS14{F0: 0}
                                                                            _ = v137
                                                                            _ = v138
                                                                        } else {
                                                                            if true {
                                                                                v139 := GenS10{F0: 953}
                                                                                _ = v139
                                                                            } else {
                                                                                switch "a" {
                                                                                case "":
                                                                                    if (0 > 2) {
                                                                                        var v140 float32 = (1.5 + ((1.5 - 0.1) / float32(float64(float32(-2.25)))))
                                                                                        v6 -= -9223372036854775808
                                                                                        _ = v140
                                                                                    }
                                                                                }
                                                                                switch v3 {
                                                                                case "a":
                                                                                    if false {
                                                                                        if true {
                                                                                            v5 = (v134 / ((65) | 1))
                                                                                            var v141 uint = 42
                                                                                            _ = v141
                                                                                        } else {
                                                                                            var v142 int = fi112
                                                                                            v3 = "hello"
                                                                                            _ = v142
                                                                                        }
                                                                                        v3 = v3
                                                                                        var v143 float64 = (3.14159 * -2.25)
                                                                                        _ = v143
                                                                                    }
                                                                                case "foo bar":
                                                                                    v6 = fi133
                                                                                }
                                                                            }
                                                                            v1 += (v1 >> ((uint(v134)) & 63))
                                                                        }
                                                                    case -2147483648:
                                                                        v6 = fi133
                                                                    }
                                                                    var v144 float64 = 3.14159
                                                                    _ = v144
                                                                } else {
                                                                    if false {
                                                                        fmt.Println(fi133, fi112)
                                                                    } else {
                                                                        v145 := []int{(fi133 >> (((^(v1))) & 63)), (v6 << (((v1 + 257)) & 63)), (fi112 % (((v6 + v6)) | 1))}
                                                                        v4 = v5
                                                                        _ = v145
                                                                    }
                                                                    v5 = (v134 % (((v4 & (v134 * v4))) | 1))
                                                                }
                                                                if false {
                                                                    v146 := []int{v6, (fi112 << ((47) & 63))}
                                                                    v4 = (v4 >> (((v1 + 1)) & 63))
                                                                    _ = v146
                                                                } else {
                                                                    v5 = (v4 | (v5 ^ 0))
                                                                    v3 = v3
                                                                    v5 = v134
                                                                }
                                                            } else {
                                                                var v147 uint = v1
                                                                fmt.Println(fi133)
                                                                v148 := []int{v6}
                                                                _ = v147
                                                                _ = v148
                                                            }
                                                        default:
                                                            if (-2.25 != (0.0 - 1e10)) {
                                                                v6 -= -797
                                                                v149 := []int{(v6 * (-(fi112)))}
                                                                v150 := GenS11{F0: 42, F1: -1}
                                                                _ = v149
                                                                _ = v150
                                                            } else {
                                                                fmt.Println(fi112, v4, v4)
                                                                v151 := v2[genFn1((v6 | -871), 42)]
                                                                switch v3 {
                                                                case "\n\t":
                                                                    switch fi112 {
                                                                    case -260:
                                                                        v152 := []int{128, -659, v6}
                                                                        if (!(-2.25 != -2.25)) {
                                                                            switch (v6 % ((329) | 1)) {
                                                                            case -279:
                                                                                var v153 bool = false
                                                                                _ = v153
                                                                            }
                                                                            var v154 bool = true
                                                                            if v154 {
                                                                                fmt.Println(v154, fi112, fi112)
                                                                                v155 := GenS3{F0: 32767, F1: v1, F2: -32768}
                                                                                var v156 uint64 = (v134 % ((v5) | 1))
                                                                                _ = v155
                                                                                _ = v156
                                                                            }
                                                                            fmt.Println(v5, v134)
                                                                            var v157 uint = v1
                                                                            switch v3 {
                                                                            case "a":
                                                                                v4 += v5
                                                                            case "\n\t":
                                                                                var v158 uint8 = 2
                                                                                _ = v158
                                                                            }
                                                                            var v159 uint = (^(v1))
                                                                            var v160 uint8 = 128
                                                                            var v161 int16 = -32768
                                                                            _ = v154
                                                                            _ = v157
                                                                            _ = v159
                                                                            _ = v160
                                                                            _ = v161
                                                                        }
                                                                        _ = v152
                                                                    case -9223372036854775808:
                                                                        v162 := GenS12{F0: 128, F1: 1, F2: 0}
                                                                        _ = v162
                                                                    }
                                                                case "":
                                                                    delete(v2, 65535)
                                                                default:
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        v163 := []int{(fi112 + (v6 / ((v6) | 1))), (fi112 >> ((835) & 63))}
                                                                        _ = v163
                                                                    }
                                                                }
                                                                _ = v151
                                                            }
                                                        }
                                                        v3 = ("foo bar" + "a")
                                                        v164 := []int{(fi133 % ((-813) | 1)), fi112}
                                                        _ = v164
                                                    }
                                                    v6 -= (fi133 >> ((255) & 63))
                                                    _ = v134
                                                }
                                                if (!false) {
                                                    v165 := GenS5{F0: 2}
                                                    var v166 bool = true
                                                    var v167 uint8 = 128
                                                    v168 := GenS15{F0: (^(v4))}
                                                    _ = v165
                                                    _ = v166
                                                    _ = v167
                                                    _ = v168
                                                }
                                                v6 = (fi133 ^ 0)
                                                _ = fi133
                                            }
                                        }
                                    case 32767:
                                        if true {
                                            fmt.Println(v6, v3)
                                        }
                                    default:
                                        if true {
                                            v169 := GenS7{F0: 873}
                                            _ = v169
                                        }
                                    }
                                }
                                if false {
                                    var v170 int = v6
                                    v171 := []int{(fi112 * v170), (fi112 >> ((genFn0(716, 255, v4)) & 63)), (fi112 >> (((v1 - (v1 | v1))) & 63))}
                                    if len(v171) >= 2 {
                                        v171 = v171[1:len(v171)]
                                    }
                                    _ = v170
                                    _ = v171
                                } else {
                                    v172 := []int{(fi112 % ((fi112) | 1)), (fi112 >> ((v1) & 63)), (fi112 * (^(fi112)))}
                                    _ = v172
                                }
                            }
                        case "a":
                            v173 := GenS10{F0: 2147483647}
                            _ = v173
                        }
                    case 127:
                        var v174 float32 = (0.0 + ((((1.5 * 0.0) * float32(float64(0.1))) / ((float32(-2.25) - (float32(1.5) - 3.14159)) / 1.5)) + 0.0))
                        _ = v174
                    }
                }
            default:
                switch v3 {
                case "a":
                    v175 := GenS9{F0: -271}
                    _ = v175
                case "foo bar":
                    switch ("foo bar" + "foo bar") {
                    case "a":
                        v176 := []int{(fi112 & fi112), (v6 * v6), v6, v6}
                        _ = v176
                    case "\n\t":
                        fmt.Println(v6)
                    default:
                        var v177 bool = (42 != 128)
                        _ = v177
                    }
                }
            }
            _ = fi112
        }
        var v178 int16 = 1
        for fi179 := 0; fi179 < 2; fi179++ {
            fmt.Println(v6, v4, v3)
            _ = fi179
        }
        _ = v178
    }
    fmt.Println(v4, v4)
    v5 = (v5 / ((v5) | 1))
    for fi180 := 0; fi180 < 15; fi180++ {
        v5 = (v4 | v5)
        var v181 float32 = (1e10 * (3.14159 * (0.0 + 0.0)))
        v182 := GenS12{F0: 170, F1: 2147483647, F2: 2}
        _ = fi180
        _ = v181
        _ = v182
    }
    if false {
        v183 := []int{}
        v184 := GenS16{F0: 0}
        _ = v183
        _ = v184
    } else {
        for fi185 := 0; fi185 < 11; fi185++ {
            v186 := GenS16{F0: 837}
            for fi187 := 0; fi187 < 6; fi187++ {
                v188 := GenS17{F0: 751, F1: (((1.5 * (float64((float32(0.1) - (1.5 * 0.0))) / 1.0)) - float64((((0.0 + 0.0) / -2.25) - 0.0))) == 1.5)}
                _ = fi187
                _ = v188
            }
            v189 := GenS10{F0: genFn1(fi185, 127)}
            _ = fi185
            _ = v186
            _ = v189
        }
        switch (v6 - v6) {
        case 32768:
            for fi190 := 0; fi190 < 16; fi190++ {
                fmt.Println(v1)
                v3 = v3
                _ = fi190
            }
        default:
            v191 := GenS10{F0: 545}
            _ = v191
        }
        fmt.Println(v4, v5, v5)
    }
    var v192 uint16 = uint16(v5)
    for fi193 := 0; fi193 < 2; fi193++ {
        v194 := GenS18{F0: (^(v1)), F1: (!(!false)), F2: 95}
        switch ((v3 + v3) + v3) {
        case "a":
            if ((v5 * (v5 | (v5 << (((v1 * 2)) & 63)))) <= 521) {
                if ((v192 + (v192 >> ((v1) & 63))) < v192) {
                    v1 -= genFn0(898, 205, 4294967295)
                    v192 -= 930
                } else {
                    v3 = "\n\t"
                    _ = len(v2)
                }
            } else {
                switch v3 {
                case "a":
                    if ("hello" <= "a") {
                        fmt.Println(v1)
                    } else {
                        var v195 uint64 = v5
                        _ = v195
                    }
                default:
                    for fi196 := 0; fi196 < 16; fi196++ {
                        v197 := GenS19{F0: 42, F1: (^(v1))}
                        v5 = v4
                        v192 = 2
                        _ = fi196
                        _ = v197
                    }
                }
                var v198 uint8 = 0
                _ = v198
            }
        case "foo bar":
            if false {
                var v199 int64 = -9223372036854775808
                _ = v199
            } else {
                v192 = 1
            }
        }
        _ = fi193
        _ = v194
    }
    v200 := []int{(-(v6))}
    switch v6 {
    case -183:
        var v201 string = v3
        _ = v201
    default:
        var v202 int = (v6 / ((-2147483648) | 1))
        _ = v202
    }
    fmt.Println(v3, v192, v5)
    v203 := GenS20{F0: 1, F1: (true && false)}
    v1 -= (v1 + 0)
    if false {
        v4 -= v5
        v4 = 42
        var v204 int32 = 750
        _ = v204
    } else {
        switch v3 {
        case "a":
            v205 := GenS6{F0: v3}
            _ = v205
        case "hello":
            v206 := GenS16{F0: 1}
            _ = v206
        case "\n\t":
            fmt.Println(v6, v5, v6)
        default:
            v207 := GenS21{F0: -1, F1: -2147483648}
            _ = v207
        }
        var v208 int8 = 42
        _ = v208
    }
    v209 := GenS15{F0: (v5 << ((v1) & 63))}
    if true {
        var v210 int16 = -170
        v211 := GenS9{F0: 1}
        _ = v210
        _ = v211
    } else {
        v212 := GenS22{F0: 827}
        _ = v212
    }
    fmt.Println("v1:", v1)
    fmt.Println("v192:", v192)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v200:", len(v200), v200)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v192
    _ = v2
    _ = v200
    _ = v203
    _ = v209
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
