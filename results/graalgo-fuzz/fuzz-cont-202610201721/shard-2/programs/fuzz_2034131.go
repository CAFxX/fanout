package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int32 }
type GenS2 struct { F0 int32; F1 int; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int; F2 uint32 }
type GenS5 struct { F0 int64; F1 uint; F2 int16 }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 int; F1 uint16; F2 int64 }
type GenS8 struct { F0 int16; F1 uint16 }
type GenS9 struct { F0 uint32; F1 int }
type GenS10 struct { F0 int64; F1 uint; F2 uint8 }
type GenS11 struct { F0 int64; F1 int; F2 int16 }
type GenS12 struct { F0 int32; F1 uint8 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int16; F1 uint8 }
type GenS15 struct { F0 uint16 }
func genFn0(p0 uint, p1 uint16, p2 uint16) uint {
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint64 = 1
        p2 = p2
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    switch 99 {
    case 127:
        var v8 uint16 = (p1 / ((p1) | 1))
        _ = v8
    case 456:
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := []int{(fi9 | -9223372036854775808)}
            v11 := GenS1{F0: (!(true && (!false))), F1: ("" + "hello"), F2: 42}
            _ = fi9
            _ = v10
            _ = v11
        }
    case 2147483647:
        switch 127 {
        case 1:
            switch 32767 {
            case 2147483647:
                v12 := GenS1{F0: ((^(p0)) != p0), F1: "foo bar", F2: 587}
                _ = v12
            case 9223372036854775807:
                if true {
                    for fi13 := 0; fi13 < 4; fi13++ {
                        p2 += 238
                        _ = fi13
                    }
                }
            case 0:
                if false {
                    p0 -= 739
                } else {
                    p1 += (p1 & (p2 << ((p0) & 63)))
                    if true {
                        p1 += (p1 * p2)
                    } else {
                        for fi14 := 0; fi14 < 15; fi14++ {
                            v15 := GenS2{F0: 1, F1: (fi14 << (((p0 << ((uint(p2)) & 63))) & 63)), F2: (fi14 >> (((p0 << (((p0 / ((p0) | 1))) & 63))) & 63))}
                            _ = fi14
                            _ = v15
                        }
                    }
                }
            default:
                var v16 uint16 = 65535
                _ = v16
            }
        }
    }
    v17 := GenS3{F0: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return p0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]int)
    var v4 uint64 = 282
    var v5 int16 = -247
    var v18 int16 = v5
    var v19 uint64 = 1
    v20 := []int{-740, 128, 32767}
    if false {
        if false {
            switch -456 {
            case 32768:
                if false {
                    if true {
                        delete(v3, 0)
                    }
                }
            case -61:
                v21 := GenS4{F0: 0, F1: int(v5), F2: 2147483647}
                _ = v21
            case -1:
                _ = len(v3)
            }
        }
        var v22 int32 = 42
        v5 = 989
        _ = v22
    } else {
        switch (("hello" + "foo bar") + "hello") {
        case "\n\t":
            v4 += v4
        case "hello":
            if (!(3.14159 != (1.5 + 0.0))) {
                if false {
                    v19 = v4
                    v4 -= v4
                    if ((0 == int32(v2)) || (!(42 == 65535))) {
                        v20 = append(v20, 4294967295)
                        v19 = (v19 | v4)
                    } else {
                        v23 := GenS3{F0: 0}
                        var v24 uint32 = 65535
                        switch ("hello" + "a") {
                        case "":
                            v4 -= v19
                        case "foo bar":
                            fmt.Println(v19, v19, v24)
                        }
                        _ = v23
                        _ = v24
                    }
                }
                if false {
                    var v25 uint32 = 991
                    _ = v25
                } else {
                    switch 42 {
                    case 459:
                        v18 -= (v18 >> ((547) & 63))
                    case 511:
                        v26 := v3[2147483647]
                        _ = v26
                    case -258:
                        var v27 int64 = 1
                        _ = v27
                    default:
                        if true {
                            switch -2147483648 {
                            case 4294967295:
                                fmt.Println(v2, v18)
                            case 127:
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                            case 2147483647:
                                v19 = v19
                            }
                            v28 := GenS5{F0: (v2 % (((v2 ^ v2)) | 1)), F1: 4294967295, F2: 1000}
                            if true {
                                for fi29 := 0; fi29 < 7; fi29++ {
                                    v19 = v19
                                    _ = fi29
                                }
                                fmt.Println(v18, v4, v18)
                            }
                            _ = v28
                        } else {
                            v18 -= (v18 | (v5 - 32767))
                            v19 = (v19 % ((v19) | 1))
                            switch "" {
                            case "hello":
                                var v30 int8 = 100
                                _ = v30
                            case "foo bar":
                                v19 += (v4 << ((4294967295) & 63))
                            case "a":
                                switch -9223372036854775808 {
                                case 0:
                                    switch "" {
                                    case "foo bar":
                                        if (4294967295 != -9223372036854775808) {
                                            var v31 int = 1
                                            for fi32 := 0; fi32 < 7; fi32++ {
                                                var v33 float32 = ((1.5 * (((((0.0 / 3.14159) + 1.5) * -2.25) * 0.0) * float32(-2.25))) - -2.25)
                                                switch "" {
                                                case "":
                                                    var v34 int = fi32
                                                    _ = v34
                                                }
                                                v35 := GenS6{F0: -32768, F1: v31}
                                                _ = fi32
                                                _ = v33
                                                _ = v35
                                            }
                                            _ = v31
                                        } else {
                                            if (true || false) {
                                                v2 = v2
                                            }
                                            v36 := GenS4{F0: 1, F1: 773, F2: 638}
                                            var v37 uint64 = (v4 % (((v4 % (((v4 << ((42) & 63))) | 1))) | 1))
                                            _ = v36
                                            _ = v37
                                        }
                                    }
                                case -255:
                                    for fi38 := 0; fi38 < 9; fi38++ {
                                        switch ("a" + "foo bar") {
                                        case "foo bar":
                                            var v39 float64 = ((0.0 - (0.1 + -2.25)) - 3.14159)
                                            _ = v39
                                        default:
                                            v40 := GenS6{F0: (v5 >> ((317) & 63)), F1: (^(fi38))}
                                            _ = v40
                                        }
                                        v2 = v2
                                        delete(v1, 65)
                                        _ = fi38
                                    }
                                case -1:
                                    var v41 int32 = 150
                                    _ = v41
                                default:
                                    var v42 uint32 = 42
                                    _ = v42
                                }
                            }
                        }
                    }
                    var v43 float32 = 1.5
                    _ = v43
                }
            } else {
                for fi44 := 0; fi44 < 1; fi44++ {
                    var v45 bool = true
                    v20 = append(v20, (fi44 + 32767))
                    var v46 uint64 = v4
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                var v47 float32 = (((1.5 * 3.14159) * ((((-2.25 - 1e10) * 0.1) * float32(3.14159)) - 0.1)) - 1e10)
                _ = v47
            }
        case "foo bar":
            v48 := v1[736]
            _ = v48
        default:
            var v49 uint8 = 0
            _ = v49
        }
        switch 2147483647 {
        case 352:
            v18 = -32768
        default:
            switch 32767 {
            case 1:
                v4 -= (v19 ^ v19)
            case 270:
                for fi50 := 0; fi50 < 6; fi50++ {
                    fmt.Println(fi50)
                    var v51 float32 = 0.0
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    _ = fi50
                    _ = v51
                }
            case -9223372036854775808:
                v52 := GenS2{F0: 2147483647, F1: -759, F2: -9223372036854775808}
                _ = v52
            default:
                if (((1000 <= 0) || true) && ((((1.5 * float32((float64(3.14159) * (3.14159 * 1e10)))) + 3.14159) - (((float32(0.1) / 1.0) + (1e10 - 0.0)) + 0.1)) == -2.25)) {
                    if false {
                        if (((42 == 22) && (!true)) && false) {
                            delete(v1, 118)
                            var v53 uint = 361
                            for fi54 := 0; fi54 < 9; fi54++ {
                                switch fi54 {
                                case -1:
                                    switch 9223372036854775807 {
                                    case 0:
                                        if (true || ((fi54 > 32768) || ((false || (true || false)) || true))) {
                                            v55 := v1[4294967295]
                                            _ = v55
                                        } else {
                                            v18 = v5
                                            var v56 int = (fi54 ^ 0)
                                            var v57 int8 = -128
                                            _ = v56
                                            _ = v57
                                        }
                                    case -9223372036854775808:
                                        v4 -= 309
                                    }
                                case -2147483648:
                                    var v58 int = fi54
                                    _ = v58
                                case 0:
                                    for fi59 := 0; fi59 < 13; fi59++ {
                                        switch ("hello" + "a") {
                                        case "foo bar":
                                            v2 -= v2
                                        case "\n\t":
                                            v2 += v2
                                        case "a":
                                            if (true && false) {
                                                var v60 uint16 = 42
                                                var v61 int = -575
                                                v4 = 18446744073709551615
                                                _ = v60
                                                _ = v61
                                            }
                                        default:
                                            v62 := GenS6{F0: v18, F1: fi59}
                                            _ = v62
                                        }
                                        var v63 int8 = -14
                                        var v64 int64 = (v2 - (v2 + int64(fi54)))
                                        _ = fi59
                                        _ = v63
                                        _ = v64
                                    }
                                default:
                                    var v65 int8 = int8(v18)
                                    _ = v65
                                }
                                var v66 int = -421
                                for fi67 := 0; fi67 < 3; fi67++ {
                                    if (1.5 >= (1e10 * (1.5 * 1.5))) {
                                        var v68 bool = false
                                        if (1 == 4294967295) {
                                            if v68 {
                                                fmt.Println(v66)
                                                switch (v66 | (fi67 << ((65535) & 63))) {
                                                case 740:
                                                    v69 := GenS5{F0: v2, F1: (v53 - (v53 % (((v53 << (((v53 - v53)) & 63))) | 1))), F2: v18}
                                                    _ = v69
                                                case 128:
                                                    v4 += 515
                                                case 127:
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v70 float64 = (((((1.5 * float64(1.5)) * (((1.5 + (((0.0 + 1.5) / 3.14159) / 0.1)) + -2.25) / ((0.0 - ((((3.14159 + float64(float32(float64((0.1 - float32((-2.25 * 0.1))))))) + ((1e10 + 0.0) / (-2.25 * 0.1))) + float64(1.5)) + float64(0.1))) / ((-2.25 - (0.1 * -2.25)) / 1e10)))) / 0.1) + float64(1e10)) + ((0.1 + (1e10 + 0.0)) + 1.5))
                                                        _ = v70
                                                    default:
                                                        v71 := GenS7{F0: (fi67 & fi54), F1: 42, F2: (v2 / ((v2) | 1))}
                                                        _ = v71
                                                    }
                                                }
                                                v72 := GenS6{F0: (v5 ^ v18), F1: fi54}
                                                _ = v72
                                            } else {
                                                v19 += (v19 & v4)
                                                v53 = 115
                                            }
                                        } else {
                                            v18 = v18
                                            if len(v20) > 0 {
                                                v20[0] = fi54
                                            }
                                        }
                                        _ = v68
                                    } else {
                                        fmt.Println(fi67, v53, v53)
                                        if ((v53 | 464) >= (v53 * (v53 | 692))) {
                                            var v73 uint16 = 1
                                            _ = v73
                                        } else {
                                            if true {
                                                v2 = v2
                                                v74 := GenS6{F0: v18, F1: 0}
                                                if false {
                                                    if ((fi54 * (fi67 | 9223372036854775807)) >= int(v53)) {
                                                        fmt.Println(v19, fi67, v4)
                                                        if len(v20) >= 2 {
                                                            v20 = v20[1:len(v20)]
                                                        }
                                                        var v75 bool = (!(-2.25 > -2.25))
                                                        _ = v75
                                                    } else {
                                                        if true {
                                                            var v76 uint16 = 65535
                                                            switch (fi54 >> ((928) & 63)) {
                                                            case -810:
                                                                switch ("a" + ("" + "foo bar")) {
                                                                case "a":
                                                                    fmt.Println(v18, v4)
                                                                    if len(v20) > 0 {
                                                                        v20[0] = (v66 | v66)
                                                                    }
                                                                case "\n\t":
                                                                    var v77 int32 = 0
                                                                    _ = v77
                                                                case "hello":
                                                                    if true {
                                                                        v19 += (v4 >> ((v53) & 63))
                                                                    }
                                                                }
                                                            case 4294967295:
                                                                if true {
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        v19 += 750
                                                                    case "":
                                                                        v5 = v5
                                                                    case "\n\t":
                                                                        if false {
                                                                            if false {
                                                                                switch 32767 {
                                                                                case 1:
                                                                                    if false {
                                                                                        v78 := GenS6{F0: v18, F1: v66}
                                                                                        v66 -= v66
                                                                                        switch 9223372036854775807 {
                                                                                        case 367:
                                                                                            switch (fi67 - (v66 / (((fi67 >> ((65535) & 63))) | 1))) {
                                                                                            case -1:
                                                                                                var v79 int16 = (v18 << (((^(v53))) & 63))
                                                                                                _ = v79
                                                                                            case -9223372036854775808:
                                                                                                v80 := GenS2{F0: 2147483647, F1: fi67, F2: fi67}
                                                                                                _ = v80
                                                                                            }
                                                                                        case -9223372036854775808:
                                                                                            fmt.Println(fi54, v76)
                                                                                        default:
                                                                                            if (!(2147483647 <= 958)) {
                                                                                                var v81 int16 = v5
                                                                                                v5 = v18
                                                                                                _ = v81
                                                                                            } else {
                                                                                                var v82 int8 = 0
                                                                                                if true {
                                                                                                    v5 = -240
                                                                                                    var v83 uint64 = (^(v19))
                                                                                                    _ = v83
                                                                                                } else {
                                                                                                    v2 = v2
                                                                                                    v82 = (v82 - (v82 ^ (v82 % ((v82) | 1))))
                                                                                                    v66 += v66
                                                                                                    if true {
                                                                                                        v20 = append(v20, (fi54 + (v66 / ((int(v18)) | 1))))
                                                                                                    } else {
                                                                                                        v84 := GenS5{F0: int64(v18), F1: (v53 + (^(v53))), F2: 1000}
                                                                                                        if (((int16(v76) <= v18) || (0 != (v82 ^ v82))) && true) {
                                                                                                            v76 -= v76
                                                                                                            var v85 uint16 = v76
                                                                                                            v86 := GenS1{F0: false, F1: ("hello" + ""), F2: -1}
                                                                                                            _ = v85
                                                                                                            _ = v86
                                                                                                        } else {
                                                                                                            v18 = (v18 % (((v18 * v5)) | 1))
                                                                                                            v19 = v4
                                                                                                        }
                                                                                                        _ = v84
                                                                                                    }
                                                                                                }
                                                                                                _ = v82
                                                                                            }
                                                                                        }
                                                                                        _ = v78
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v87 uint32 = 2147483647
                                                                                var v88 bool = false
                                                                                var v89 string = "hello"
                                                                                _ = v87
                                                                                _ = v88
                                                                                _ = v89
                                                                            }
                                                                        } else {
                                                                            switch v66 {
                                                                            case 0:
                                                                                if false {
                                                                                    switch (^(v66)) {
                                                                                    case -417:
                                                                                        var v90 bool = (v53 != 146)
                                                                                        var v91 uint16 = 0
                                                                                        var v92 uint16 = (v76 * 769)
                                                                                        var v93 int64 = (v2 >> ((v53) & 63))
                                                                                        var v94 uint64 = v4
                                                                                        _ = v90
                                                                                        _ = v91
                                                                                        _ = v92
                                                                                        _ = v93
                                                                                        _ = v94
                                                                                    case 871:
                                                                                        v95 := GenS8{F0: 42, F1: (^(v76))}
                                                                                        fmt.Println(v4)
                                                                                        v18 = -1
                                                                                        _ = v95
                                                                                    case -957:
                                                                                        if false {
                                                                                            v1[65535] = true
                                                                                        }
                                                                                    }
                                                                                    v96 := GenS8{F0: 1, F1: (v76 % ((v76) | 1))}
                                                                                    fmt.Println(v19, v5, v19)
                                                                                    _ = v96
                                                                                } else {
                                                                                    if (!false) {
                                                                                        v97 := GenS4{F0: (^(v76)), F1: 0, F2: 42}
                                                                                        _ = v97
                                                                                    } else {
                                                                                        var v98 float32 = float32(3.14159)
                                                                                        v19 -= v4
                                                                                        _ = v98
                                                                                    }
                                                                                    v5 += (v18 - (v18 * (v5 | (v18 % ((v5) | 1)))))
                                                                                    if len(v20) > 0 {
                                                                                        v20[0] = fi54
                                                                                    }
                                                                                }
                                                                            case 32767:
                                                                                v99 := GenS5{F0: (v2 / ((-912) | 1)), F1: genFn0((v53 * v53), (v76 ^ 42), v76), F2: v18}
                                                                                _ = v99
                                                                            case 659:
                                                                                v53 += v53
                                                                            default:
                                                                                var v100 uint16 = 32767
                                                                                _ = v100
                                                                            }
                                                                        }
                                                                        fmt.Println(v5)
                                                                        if (0 <= 0) {
                                                                            var v101 int64 = (-(v2))
                                                                            _ = v101
                                                                        }
                                                                    }
                                                                    switch ("hello" + "foo bar") {
                                                                    case "":
                                                                        _ = len(v3)
                                                                    case "\n\t":
                                                                        switch v66 {
                                                                        case -336:
                                                                            v76 = v76
                                                                        case -842:
                                                                            if true {
                                                                                switch fi54 {
                                                                                case 32768:
                                                                                    delete(v1, 1)
                                                                                default:
                                                                                    v102 := GenS7{F0: (v66 << (((v53 << (((v53 % ((v53) | 1))) & 63))) & 63)), F1: v76, F2: 9223372036854775807}
                                                                                    _ = v102
                                                                                }
                                                                            } else {
                                                                                v103 := GenS9{F0: 65535, F1: -38}
                                                                                v53 = v53
                                                                                switch int(v19) {
                                                                                case 4294967295:
                                                                                    switch (v66 - (v66 | fi54)) {
                                                                                    case 32767:
                                                                                        if (!false) {
                                                                                            if false {
                                                                                                var v104 int8 = int8(fi54)
                                                                                                v1[574] = (!((^(v53)) >= 2))
                                                                                                _ = v104
                                                                                            }
                                                                                            v5 = v5
                                                                                        } else {
                                                                                            if true {
                                                                                                var v105 uint8 = uint8(v4)
                                                                                                switch "" {
                                                                                                case "hello":
                                                                                                    if len(v20) > 0 {
                                                                                                        v20[0] = 32768
                                                                                                    }
                                                                                                default:
                                                                                                    switch fi67 {
                                                                                                    case 2147483647:
                                                                                                        fmt.Println(fi67, fi54, v19)
                                                                                                    case 42:
                                                                                                        var v106 int16 = v5
                                                                                                        _ = v106
                                                                                                    case 32768:
                                                                                                        v66 = 128
                                                                                                    default:
                                                                                                        fmt.Println(v2, v2)
                                                                                                        switch (fi67 << ((uint(v105)) & 63)) {
                                                                                                        case -2147483648:
                                                                                                            if true {
                                                                                                                switch ("a" + "a") {
                                                                                                                case "a":
                                                                                                                    _ = len(v3)
                                                                                                                case "hello":
                                                                                                                    v3[42] = (v66 - (fi54 * fi67))
                                                                                                                }
                                                                                                            } else {
                                                                                                                fmt.Println(v66, v2, v5)
                                                                                                                v3[1] = fi67
                                                                                                                switch "\n\t" {
                                                                                                                case "hello":
                                                                                                                    v4 = 0
                                                                                                                default:
                                                                                                                    delete(v1, 0)
                                                                                                                }
                                                                                                            }
                                                                                                        case 9223372036854775807:
                                                                                                            fmt.Println(v2, v4, v105)
                                                                                                        case -1:
                                                                                                            v107 := GenS7{F0: -489, F1: uint16(v19), F2: (v2 << (((v53 & (v53 - v53))) & 63))}
                                                                                                            _ = v107
                                                                                                        default:
                                                                                                            var v108 uint8 = 255
                                                                                                            _ = v108
                                                                                                        }
                                                                                                        v53 -= (v53 * (v53 % (((v53 * genFn0((v53 | (v53 * 42)), (v76 >> (((v53 >> ((634) & 63))) & 63)), (v76 & (^(v76)))))) | 1)))
                                                                                                    }
                                                                                                }
                                                                                                _ = v105
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        if false {
                                                                                            if (127 >= 7) {
                                                                                                v109 := GenS10{F0: 945, F1: 4294967295, F2: 2}
                                                                                                _ = v109
                                                                                            } else {
                                                                                                if true {
                                                                                                    v4 = 124
                                                                                                } else {
                                                                                                    if len(v20) >= 2 {
                                                                                                        v20 = v20[1:len(v20)]
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                case 0:
                                                                                    switch (v66 - fi54) {
                                                                                    case -772:
                                                                                        if len(v20) >= 2 {
                                                                                            v20 = v20[1:len(v20)]
                                                                                        }
                                                                                    case 128:
                                                                                        var v110 uint16 = (v76 % ((42) | 1))
                                                                                        _ = v110
                                                                                    case -1:
                                                                                        fmt.Println(v76, v76, v5)
                                                                                    }
                                                                                    v111 := GenS8{F0: v18, F1: v76}
                                                                                    _ = v111
                                                                                }
                                                                                _ = v103
                                                                            }
                                                                        }
                                                                    default:
                                                                        v112 := GenS4{F0: 32767, F1: fi54, F2: 875}
                                                                        _ = v112
                                                                    }
                                                                }
                                                            case 9223372036854775807:
                                                                if ((^(v2)) > (v2 + 42)) {
                                                                    var v113 int64 = (-(v2))
                                                                    fmt.Println(v113, v113)
                                                                    _ = v113
                                                                }
                                                            }
                                                            v114 := GenS10{F0: v2, F1: 65535, F2: 42}
                                                            v76 = (v76 << (((v53 + (^(v53)))) & 63))
                                                            if len(v20) >= 2 {
                                                                v20 = v20[1:len(v20)]
                                                            }
                                                            _ = v114
                                                            _ = v76
                                                        } else {
                                                            if true {
                                                                fmt.Println(fi54, v4, v4)
                                                                if false {
                                                                    switch v66 {
                                                                    case -788:
                                                                        var v115 uint = v53
                                                                        var v116 int32 = 42
                                                                        if len(v20) >= 2 {
                                                                            v20 = v20[1:len(v20)]
                                                                        }
                                                                        _ = v115
                                                                        _ = v116
                                                                    case -9223372036854775808:
                                                                        v19 = (v4 & (^(v4)))
                                                                        var v117 uint = (^(v53))
                                                                        v2 = (v2 + v2)
                                                                        _ = v117
                                                                    case 42:
                                                                        var v118 string = "hello"
                                                                        _ = v118
                                                                    }
                                                                    _ = len(v20)
                                                                } else {
                                                                    if false {
                                                                        if len(v20) > 0 {
                                                                            v20[0] = fi67
                                                                        }
                                                                        switch "hello" {
                                                                        case "a":
                                                                            v18 = v18
                                                                        case "\n\t":
                                                                            var v119 uint8 = 127
                                                                            _ = v119
                                                                        default:
                                                                            v4 = (v4 << ((2) & 63))
                                                                        }
                                                                    }
                                                                }
                                                                v19 += (v4 | 2)
                                                            } else {
                                                                v20 = append(v20, fi54)
                                                                v5 += (v5 * v5)
                                                            }
                                                        }
                                                        var v120 uint16 = 215
                                                        var v121 uint = v53
                                                        _ = v120
                                                        _ = v121
                                                    }
                                                    fmt.Println(v18)
                                                    v5 = v5
                                                    v122 := GenS7{F0: 363, F1: 2, F2: (v2 >> ((673) & 63))}
                                                    _ = v122
                                                } else {
                                                    switch (("a" + "a") + "hello") {
                                                    case "hello":
                                                        if true {
                                                            var v123 uint16 = 1
                                                            v123 = (v123 % ((32767) | 1))
                                                            var v124 bool = false
                                                            _ = v123
                                                            _ = v124
                                                        }
                                                    default:
                                                        if false {
                                                            if false {
                                                                var v125 uint16 = 65535
                                                                if ("hello" == ("foo bar" + "foo bar")) {
                                                                    delete(v1, 0)
                                                                    var v126 uint = v53
                                                                    v5 = (v18 | v5)
                                                                    _ = v126
                                                                }
                                                                v125 = 0
                                                                _ = v125
                                                            }
                                                            v127 := GenS8{F0: 1, F1: 255}
                                                            v18 -= 0
                                                            _ = v127
                                                        }
                                                    }
                                                    if ((2 <= 2) || (true || ((!false) || ((true || (float64(float32((((1.5 * 1.5) - (-2.25 * 0.0)) * ((float64(0.1) * 3.14159) * 1.5)))) < 1.5)) || (int32(v4) >= 0))))) {
                                                        var v128 uint32 = 4294967295
                                                        var v129 int16 = v5
                                                        var v130 string = "\n\t"
                                                        _ = v128
                                                        _ = v129
                                                        _ = v130
                                                    } else {
                                                        v131 := GenS4{F0: 1, F1: fi54, F2: 2147483648}
                                                        _ = v131
                                                    }
                                                    var v132 uint32 = 762
                                                    if false {
                                                        v3[v132] = v66
                                                        if (-2.25 != (-2.25 + ((-2.25 / 1e10) / (-2.25 - (0.1 + float32(float64(1e10))))))) {
                                                            var v133 uint = v53
                                                            var v134 bool = false
                                                            if len(v20) >= 2 {
                                                                v20 = v20[1:len(v20)]
                                                            }
                                                            _ = v133
                                                            _ = v134
                                                        }
                                                    } else {
                                                        v1[(^(v132))] = (((1 == 102) && true) && false)
                                                    }
                                                    _ = v132
                                                }
                                                v66 = (v66 + 32767)
                                                var v135 int32 = 213
                                                _ = v135
                                                _ = v74
                                            } else {
                                                v136 := GenS11{F0: v2, F1: -543, F2: v18}
                                                _ = len(v20)
                                                if len(v20) >= 2 {
                                                    v20 = v20[1:len(v20)]
                                                }
                                                _ = v136
                                            }
                                            var v137 bool = true
                                            if len(v20) > 0 {
                                                v20[0] = v66
                                            }
                                            _ = v137
                                        }
                                    }
                                    if (!true) {
                                        var v138 int16 = 0
                                        var v139 uint8 = uint8(v4)
                                        v139 += v139
                                        _ = len(v20)
                                        v140 := v3[1]
                                        var v141 string = "foo bar"
                                        switch v141 {
                                        case "a":
                                            v66 = -9223372036854775808
                                        case "hello":
                                            fmt.Println(v18)
                                        default:
                                            _ = len(v20)
                                        }
                                        _ = v138
                                        _ = v139
                                        _ = v140
                                        _ = v141
                                    } else {
                                        var v142 string = ("" + "foo bar")
                                        v53 = (v53 - (^(v53)))
                                        switch v66 {
                                        case 32768:
                                            v19 += (^(v4))
                                        }
                                        _ = v142
                                    }
                                    v53 -= (v53 * (^(v53)))
                                    switch 0 {
                                    case 404:
                                        fmt.Println(v53, v19)
                                    case 9223372036854775807:
                                        var v143 int16 = v18
                                        var v144 int = (^(fi54))
                                        if false {
                                            v145 := GenS2{F0: 42, F1: 9223372036854775807, F2: (v66 % ((-795) | 1))}
                                            v66 = 32768
                                            _ = v145
                                        } else {
                                            v146 := GenS4{F0: 750, F1: fi54, F2: 2}
                                            fmt.Println(v53, v53, v66)
                                            var v147 float64 = 3.14159
                                            var v148 uint8 = 255
                                            v53 = v53
                                            var v149 uint16 = 65535
                                            _ = v146
                                            _ = v147
                                            _ = v148
                                            _ = v149
                                        }
                                        _ = v143
                                        _ = v144
                                    }
                                    _ = fi67
                                }
                                _ = fi54
                                _ = v66
                            }
                            _ = v53
                        } else {
                            v3[42] = 32768
                        }
                    }
                    v19 += 18446744073709551615
                } else {
                    for fi150 := 0; fi150 < 1; fi150++ {
                        for fi151 := 0; fi151 < 10; fi151++ {
                            if (1 <= 275) {
                                fmt.Println(fi151)
                            }
                            if (!(!false)) {
                                var v152 uint64 = (v4 >> ((0) & 63))
                                if false {
                                    v152 += 2
                                    if true {
                                        delete(v3, 175)
                                    }
                                    v153 := GenS3{F0: 42}
                                    _ = v153
                                }
                                _ = v152
                            } else {
                                if (!(v2 <= v2)) {
                                    v4 += (v19 * v4)
                                } else {
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                    if true {
                                        v154 := GenS1{F0: true, F1: ("foo bar" + ""), F2: 0}
                                        var v155 int32 = 1
                                        v3[776] = 2147483647
                                        _ = v154
                                        _ = v155
                                    } else {
                                        v18 -= 1
                                        v19 -= 602
                                    }
                                    if false {
                                        v156 := GenS12{F0: int32(v18), F1: 42}
                                        _ = v156
                                    }
                                }
                                if (fi150 == (fi150 * (^(fi151)))) {
                                    switch (-(fi150)) {
                                    case 709:
                                        var v157 float64 = 1e10
                                        _ = v157
                                    case 9223372036854775807:
                                        if (int64(v18) <= 9223372036854775807) {
                                            v2 -= (-(v2))
                                            fmt.Println(fi151, fi150, v18)
                                        } else {
                                            var v158 int32 = 2147483647
                                            v20 = append(v20, fi150)
                                            _ = v158
                                        }
                                    case -2147483648:
                                        v159 := GenS5{F0: (v2 * (-(v2))), F1: 4294967295, F2: -1}
                                        _ = v159
                                    }
                                    var v160 int16 = (v18 ^ 1000)
                                    _ = v160
                                } else {
                                    if (3.14159 <= ((1e10 / 0.1) + 1.5)) {
                                        var v161 int = 0
                                        _ = v161
                                    }
                                    var v162 int32 = 42
                                    _ = v162
                                }
                                v2 = (v2 ^ (v2 | 0))
                            }
                            v2 = (v2 >> ((865) & 63))
                            _ = fi151
                        }
                        _ = fi150
                    }
                }
            }
        }
        v163 := GenS13{F0: false}
        _ = v163
    }
    v18 -= v5
    v19 += (^(v4))
    var v164 uint16 = 1
    v5 = (-(v18))
    fmt.Println(v19, v2, v19)
    _ = len(v1)
    v4 = v4
    v18 += (v18 % ((42) | 1))
    v165 := GenS14{F0: 196, F1: 127}
    for fi166 := 0; fi166 < 14; fi166++ {
        for fi167 := 0; fi167 < 3; fi167++ {
            v168 := GenS12{F0: 42, F1: 132}
            if ((v2 + -9223372036854775808) != -1) {
                switch "" {
                case "\n\t":
                    var v169 int16 = int16(fi166)
                    _ = v169
                case "a":
                    v170 := GenS2{F0: -1, F1: fi167, F2: int(v19)}
                    _ = v170
                default:
                    switch ((("a" + (("\n\t" + ("foo bar" + "")) + "")) + "hello") + "") {
                    case "":
                        v2 = (-(v2))
                    case "\n\t":
                        delete(v1, 325)
                    case "a":
                        _ = len(v3)
                    default:
                        if (true || true) {
                            var v171 int8 = -100
                            if (true || (2 >= 504)) {
                                var v172 int = fi167
                                _ = v172
                            } else {
                                v164 = 752
                            }
                            _ = v171
                        } else {
                            var v173 uint = 255
                            _ = v173
                        }
                    }
                }
            } else {
                var v174 float32 = float32(1.5)
                switch (((("a" + "\n\t") + "") + "a") + "\n\t") {
                case "":
                    v4 = (v19 * (v19 | 0))
                case "a":
                    if len(v20) > 0 {
                        v20[0] = fi166
                    }
                default:
                    if ((false || true) || false) {
                        v175 := GenS4{F0: (v164 | (v164 ^ 605)), F1: (fi167 % ((265) | 1)), F2: 4294967295}
                        _ = v175
                    } else {
                        var v176 float32 = float32(0.0)
                        v177 := GenS15{F0: (v164 / ((v164) | 1))}
                        var v178 uint8 = 2
                        _ = v176
                        _ = v177
                        _ = v178
                    }
                }
                var v179 bool = (!true)
                _ = v174
                _ = v179
            }
            _ = fi167
            _ = v168
        }
        _ = fi166
    }
    var v180 uint64 = 18446744073709551615
    for fi181 := 0; fi181 < 11; fi181++ {
        var v182 float64 = (-2.25 / 1.5)
        _ = fi181
        _ = v182
    }
    v183 := GenS12{F0: 0, F1: 86}
    for fi184 := 0; fi184 < 15; fi184++ {
        if len(v20) > 0 {
            v20[0] = (^(fi184))
        }
        fmt.Println(v19, v5)
        _ = fi184
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v164:", v164)
    fmt.Println("v18:", v18)
    fmt.Println("v180:", v180)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v164
    _ = v165
    _ = v18
    _ = v180
    _ = v183
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
