package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 bool; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint64; F1 int; F2 string }
type GenS6 struct { F0 uint16; F1 int64; F2 int16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int; F1 int; F2 uint64 }
type GenS10 struct { F0 uint32; F1 int8; F2 bool }
type GenS11 struct { F0 int32; F1 uint; F2 uint16 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint8; F1 string; F2 int }
type GenS14 struct { F0 uint16; F1 bool; F2 int64 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 uint16; F1 bool; F2 int8 }
type GenS18 struct { F0 bool }
func genFn0(p0 uint64, p1 uint64) uint16 {
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 float64 = (1.5 * 0.1)
        _ = fi6
        _ = v7
    }
    p1 = (^(p1))
    if false {
        v8 := GenS1{F0: 42, F1: (-198 > 0), F2: 127}
        _ = v8
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    v1 := []string{"foo bar", "foo bar", ("foo bar" + ("a" + "hello")), ""}
    v2 := []uint32{2147483647, 65535}
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]int)
    var v5 int = 2147483647
    v5 += (v5 << ((1) & 63))
    var v9 uint32 = 4294967295
    for fi10 := 0; fi10 < 16; fi10++ {
        if true {
            for fi11 := 0; fi11 < 1; fi11++ {
                v5 -= -211
                if ((18446744073709551615 >= 2) && (v5 >= fi11)) {
                    v5 += 305
                    var v12 int32 = -1
                    v9 += (v9 & (v9 | v9))
                    _ = v12
                } else {
                    if false {
                        v13 := GenS1{F0: 0, F1: true, F2: fi11}
                        v14 := v3[("foo bar" + "\n\t")]
                        _ = v13
                        _ = v14
                    }
                    v9 -= (v9 * (v9 + v9))
                    if false {
                        var v15 float32 = ((1.5 / 1.0) + 1.5)
                        if (!true) {
                            switch (^(v5)) {
                            case 9223372036854775807:
                                v5 += (-(fi11))
                            default:
                                fmt.Println(v9, v15, v5)
                            }
                            v16 := GenS1{F0: 95, F1: (false || (true && true)), F2: (v5 / ((fi10) | 1))}
                            v17 := GenS1{F0: -128, F1: true, F2: fi11}
                            _ = v16
                            _ = v17
                        } else {
                            fmt.Println(fi10)
                        }
                        if ("a" >= "foo bar") {
                            if (32767 <= fi10) {
                                v4[(v9 * v9)] = (fi10 & (v5 & 128))
                                var v18 int16 = 0
                                if (0 < 42) {
                                    fmt.Println(v18, v9, fi11)
                                }
                                _ = v18
                            } else {
                                switch v5 {
                                case -808:
                                    var v19 uint16 = 255
                                    _ = v19
                                }
                                switch (-(fi11)) {
                                case 579:
                                    var v20 uint8 = 255
                                    _ = v20
                                default:
                                    switch fi10 {
                                    case -1:
                                        v21 := GenS1{F0: 36, F1: (0 != 79), F2: fi10}
                                        _ = v21
                                    case 32768:
                                        var v22 uint64 = 575
                                        _ = v22
                                    default:
                                        fmt.Println(fi10, fi11, fi11)
                                    }
                                }
                            }
                        }
                        _ = v15
                    }
                }
                _ = fi11
            }
            for fi23 := 0; fi23 < 5; fi23++ {
                v24 := GenS1{F0: 9, F1: (!true), F2: (v5 >> ((4294967295) & 63))}
                fmt.Println(fi23, fi23)
                _ = fi23
                _ = v24
            }
            v25 := GenS1{F0: 78, F1: (true && (!(!((true || false) || (50 == 892))))), F2: fi10}
            _ = v25
        } else {
            if true {
                v9 = v9
                fmt.Println(fi10, v5, fi10)
                switch "a" {
                case "hello":
                    var v26 uint64 = 4294967295
                    _ = v26
                default:
                    var v27 int64 = int64(v9)
                    _ = v27
                }
            } else {
                v9 = v9
            }
            _ = len(v4)
            var v28 string = "\n\t"
            _ = v28
        }
        for fi29 := 0; fi29 < 20; fi29++ {
            var v30 string = ""
            _ = fi29
            _ = v30
        }
        _ = fi10
    }
    v5 += (v5 | (v5 % ((2147483647) | 1)))
    var v31 float32 = 0.0
    if true {
        if false {
            v32 := GenS2{F0: (v9 * (v9 >> ((2) & 63))), F1: v9}
            if (!true) {
                if false {
                    for fi33 := 0; fi33 < 20; fi33++ {
                        var v34 uint8 = uint8(v9)
                        delete(v3, "\n\t")
                        if false {
                            v9 = 4294967295
                        } else {
                            if len(v1) > 0 {
                                v1[0] = "foo bar"
                            }
                        }
                        _ = fi33
                        _ = v34
                    }
                    if (v31 >= v31) {
                        if (0 != 1) {
                            for fi35 := 0; fi35 < 7; fi35++ {
                                v36 := GenS1{F0: -74, F1: true, F2: (v5 + (v5 ^ -217))}
                                if false {
                                    for fi37 := 0; fi37 < 8; fi37++ {
                                        fmt.Println(v9)
                                        _ = fi37
                                    }
                                    var v38 uint8 = 255
                                    v39 := GenS3{F0: 2147483647, F1: 4294967295}
                                    _ = v38
                                    _ = v39
                                } else {
                                    v9 = (v9 | v9)
                                    v9 += v9
                                    var v40 float64 = 3.14159
                                    _ = v40
                                }
                                _ = fi35
                                _ = v36
                            }
                            _ = len(v1)
                            v3[""] = (v9 * v9)
                        } else {
                            for fi41 := 0; fi41 < 4; fi41++ {
                                if len(v1) > 0 {
                                    v1[0] = ((("a" + "\n\t") + "hello") + "foo bar")
                                }
                                if false {
                                    _ = len(v4)
                                    switch (fi41 + (v5 - (v5 >> ((646) & 63)))) {
                                    case -1:
                                        switch "\n\t" {
                                        case "hello":
                                            var v42 bool = false
                                            _ = v42
                                        default:
                                            v31 = 3.14159
                                        }
                                    case -354:
                                        if ("a" <= "\n\t") {
                                            var v43 float32 = v31
                                            _ = v43
                                        } else {
                                            if true {
                                                switch (fi41 >> ((18446744073709551615) & 63)) {
                                                case 128:
                                                    v9 = v9
                                                case 469:
                                                    if (!false) {
                                                        v9 = 42
                                                    } else {
                                                        v44 := GenS3{F0: 2, F1: 1}
                                                        for fi45 := 0; fi45 < 11; fi45++ {
                                                            v31 -= -2.25
                                                            _ = fi45
                                                        }
                                                        var v46 float32 = ((((v31 / 1.5) * (v31 - v31)) - v31) * 3.14159)
                                                        _ = v44
                                                        _ = v46
                                                    }
                                                default:
                                                    switch v5 {
                                                    case 909:
                                                        for fi47 := 0; fi47 < 9; fi47++ {
                                                            switch (v5 ^ 42) {
                                                            case 4294967295:
                                                                delete(v3, (("\n\t" + "") + "a"))
                                                            case -1:
                                                                var v48 bool = true
                                                                _ = v48
                                                            case 0:
                                                                v9 = 65535
                                                            default:
                                                                if false {
                                                                    v31 = (3.14159 / v31)
                                                                    if false {
                                                                        switch "" {
                                                                        case "":
                                                                            var v49 int32 = -2147483648
                                                                            _ = v49
                                                                        case "hello":
                                                                            if false {
                                                                                var v50 uint32 = v9
                                                                                v9 = (v9 / ((65535) | 1))
                                                                                _ = v50
                                                                            } else {
                                                                                v9 += v9
                                                                                if ("a" != (("foo bar" + "\n\t") + (("" + ((("foo bar" + "a") + ("hello" + ("hello" + "\n\t"))) + "hello")) + ("" + "")))) {
                                                                                    fmt.Println(v5, fi47, v5)
                                                                                    v5 = v5
                                                                                }
                                                                                if (!false) {
                                                                                    if (v5 <= fi41) {
                                                                                        v5 = fi47
                                                                                        v9 += v9
                                                                                        v31 += -2.25
                                                                                    } else {
                                                                                        var v51 string = "foo bar"
                                                                                        if (-1 != 0) {
                                                                                            v52 := GenS4{F0: false, F1: 2, F2: 42}
                                                                                            var v53 uint = 0
                                                                                            fmt.Println(v31, v53)
                                                                                            if (true || (true || false)) {
                                                                                                v31 = 0.1
                                                                                                fmt.Println(v9, v31)
                                                                                            }
                                                                                            _ = v52
                                                                                            _ = v53
                                                                                        }
                                                                                        v31 = v31
                                                                                        _ = v51
                                                                                    }
                                                                                    if false {
                                                                                        var v54 int64 = 42
                                                                                        fmt.Println(v9, v5, v9)
                                                                                        var v55 uint8 = 127
                                                                                        _ = v54
                                                                                        _ = v55
                                                                                    } else {
                                                                                        fmt.Println(fi41, fi41, v5)
                                                                                        if (true || true) {
                                                                                            switch -1 {
                                                                                            case 766:
                                                                                                var v56 uint8 = 42
                                                                                                _ = v56
                                                                                            case 9223372036854775807:
                                                                                                v5 = fi47
                                                                                            case 67:
                                                                                                v57 := GenS5{F0: 18446744073709551615, F1: (^(v5)), F2: "\n\t"}
                                                                                                _ = v57
                                                                                            }
                                                                                        }
                                                                                        v58 := GenS2{F0: v9, F1: v9}
                                                                                        _ = v58
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(fi41, v31)
                                                                                    v59 := GenS2{F0: v9, F1: (^(v9))}
                                                                                    _ = v59
                                                                                }
                                                                                var v60 float64 = float64(v31)
                                                                                switch (fi47 >> ((1) & 63)) {
                                                                                case -564:
                                                                                    fmt.Println(v5)
                                                                                case -2147483648:
                                                                                    v61 := GenS6{F0: 768, F1: 368, F2: -859}
                                                                                    _ = v61
                                                                                default:
                                                                                    if true {
                                                                                        var v62 int = (fi47 ^ 9223372036854775807)
                                                                                        fmt.Println(v5, fi47, v60)
                                                                                        v63 := GenS7{F0: 845}
                                                                                        _ = v62
                                                                                        _ = v63
                                                                                    } else {
                                                                                        v9 = (^(v9))
                                                                                        fmt.Println(v5)
                                                                                    }
                                                                                    fmt.Println(v60, v5)
                                                                                }
                                                                                delete(v4, (v9 * (v9 - v9)))
                                                                                if ((1e10 == v60) && false) {
                                                                                    switch "\n\t" {
                                                                                    case "":
                                                                                        switch (fi41 >> ((4294967295) & 63)) {
                                                                                        case 0:
                                                                                            var v64 uint = 1
                                                                                            _ = v64
                                                                                        case 2147483647:
                                                                                            if ((2 <= 65535) || true) {
                                                                                                _ = len(v1)
                                                                                                var v65 uint8 = 255
                                                                                                v2 = append(v2, 2147483648)
                                                                                                _ = v65
                                                                                            }
                                                                                        }
                                                                                    case "\n\t":
                                                                                        if false {
                                                                                            v1 = append(v1, ("foo bar" + ""))
                                                                                            delete(v3, ((("a" + ("foo bar" + "foo bar")) + ("hello" + "\n\t")) + "hello"))
                                                                                        } else {
                                                                                            v66 := GenS4{F0: false, F1: 255, F2: -435}
                                                                                            fmt.Println(v60, v31)
                                                                                            _ = v66
                                                                                        }
                                                                                        var v67 int = (^(v5))
                                                                                        var v68 bool = false
                                                                                        _ = v67
                                                                                        _ = v68
                                                                                    case "hello":
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = (v9 - v9)
                                                                                        }
                                                                                    }
                                                                                    if true {
                                                                                        var v69 uint64 = 18446744073709551615
                                                                                        _ = v69
                                                                                    }
                                                                                } else {
                                                                                    v9 += v9
                                                                                    if false {
                                                                                        if false {
                                                                                            fmt.Println(fi47, v60, v31)
                                                                                        } else {
                                                                                            var v70 uint = 4294967295
                                                                                            var v71 int = fi47
                                                                                            v72 := GenS3{F0: (v9 ^ (v9 * v9)), F1: 98}
                                                                                            _ = v70
                                                                                            _ = v71
                                                                                            _ = v72
                                                                                        }
                                                                                        if true {
                                                                                            v73 := GenS7{F0: 1}
                                                                                            v9 -= (v9 ^ (v9 * 0))
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            switch ("foo bar" + "foo bar") {
                                                                                            case "a":
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = 681
                                                                                                }
                                                                                            case "hello":
                                                                                                v31 = 3.14159
                                                                                            }
                                                                                            _ = v73
                                                                                        }
                                                                                        v31 += v31
                                                                                    } else {
                                                                                        if false {
                                                                                            if true {
                                                                                                if (1 >= 18446744073709551615) {
                                                                                                    v74 := GenS8{F0: -975}
                                                                                                    var v75 int32 = 670
                                                                                                    _ = v74
                                                                                                    _ = v75
                                                                                                }
                                                                                            }
                                                                                            v3[("a" + ("" + "\n\t"))] = (v9 & (^(v9)))
                                                                                        }
                                                                                        fmt.Println(v9, v31, fi47)
                                                                                        if (261 > 164) {
                                                                                            v31 = float32(float64(-2.25))
                                                                                            fmt.Println(fi41, v60, fi47)
                                                                                            v4[(v9 >> ((255) & 63))] = fi41
                                                                                        } else {
                                                                                            fmt.Println(v60)
                                                                                            v60 -= -2.25
                                                                                        }
                                                                                    }
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                }
                                                                                v76 := GenS8{F0: -1}
                                                                                if (57 >= 104) {
                                                                                    v31 += (-2.25 / ((float32(v60) - 1e10) - ((v31 + v31) - 3.14159)))
                                                                                }
                                                                                _ = v60
                                                                                _ = v76
                                                                            }
                                                                        case "\n\t":
                                                                            delete(v3, "hello")
                                                                        }
                                                                    } else {
                                                                        var v77 float64 = -2.25
                                                                        fmt.Println(fi41, v9, v5)
                                                                        _ = v77
                                                                    }
                                                                    fmt.Println(v9, fi41, fi47)
                                                                    switch ("" + "hello") {
                                                                    case "hello":
                                                                        var v78 int32 = -148
                                                                        _ = v78
                                                                    case "":
                                                                        var v79 bool = true
                                                                        if v79 {
                                                                            fmt.Println(v5)
                                                                            var v80 uint32 = (v9 & (^(v9)))
                                                                            _ = v80
                                                                        } else {
                                                                            v5 -= -1
                                                                            fmt.Println(v31, fi41, v9)
                                                                            v31 -= v31
                                                                        }
                                                                        fmt.Println(v9)
                                                                        _ = v79
                                                                    case "\n\t":
                                                                        fmt.Println(fi47)
                                                                    }
                                                                } else {
                                                                    switch (^(fi47)) {
                                                                    case 865:
                                                                        fmt.Println(v9, v31, fi41)
                                                                    case 42:
                                                                        v81 := GenS9{F0: fi41, F1: fi47, F2: 0}
                                                                        _ = v81
                                                                    default:
                                                                        var v82 int8 = -128
                                                                        switch fi41 {
                                                                        case 276:
                                                                            var v83 bool = true
                                                                            _ = v83
                                                                        case -1:
                                                                            v31 = float32(0.0)
                                                                        default:
                                                                            v82 += v82
                                                                        }
                                                                        if true {
                                                                            v5 += v5
                                                                        }
                                                                        _ = v82
                                                                    }
                                                                }
                                                            }
                                                            var v84 uint32 = v9
                                                            _ = len(v3)
                                                            switch "a" {
                                                            case "a":
                                                                if false {
                                                                    var v85 float32 = (float32(1.5) / 1.5)
                                                                    _ = v85
                                                                } else {
                                                                    v86 := GenS1{F0: 59, F1: false, F2: fi47}
                                                                    _ = v86
                                                                }
                                                            case "hello":
                                                                v87 := GenS10{F0: v9, F1: -88, F2: (int64(v9) >= 42)}
                                                                _ = v87
                                                            case "":
                                                                switch (v5 / (((v5 >> ((65535) & 63))) | 1)) {
                                                                case 128:
                                                                    v9 = (v84 ^ (v84 % ((v84) | 1)))
                                                                default:
                                                                    fmt.Println(v9)
                                                                }
                                                            default:
                                                                var v88 int32 = 400
                                                                _ = v88
                                                            }
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            delete(v4, v9)
                                                            var v89 uint64 = 4294967295
                                                            _ = fi47
                                                            _ = v84
                                                            _ = v89
                                                        }
                                                    case 127:
                                                        v31 -= 3.14159
                                                    case -1:
                                                        for fi90 := 0; fi90 < 20; fi90++ {
                                                            v5 -= (-(fi90))
                                                            delete(v3, "")
                                                            _ = fi90
                                                        }
                                                    }
                                                }
                                                v31 = v31
                                            } else {
                                                var v91 float64 = 1e10
                                                _ = v91
                                            }
                                        }
                                    default:
                                        if false {
                                            v9 = 42
                                            delete(v4, (v9 << ((uint(v9)) & 63)))
                                            v9 = v9
                                        }
                                    }
                                    _ = len(v4)
                                }
                                _ = fi41
                            }
                            for fi92 := 0; fi92 < 2; fi92++ {
                                v5 = v5
                                _ = fi92
                            }
                        }
                        fmt.Println(v31)
                        switch "" {
                        case "foo bar":
                            var v93 int32 = -2147483648
                            _ = v93
                        case "\n\t":
                            var v94 uint16 = 36
                            _ = v94
                        case "hello":
                            for fi95 := 0; fi95 < 12; fi95++ {
                                v9 += v9
                                _ = fi95
                            }
                        }
                    }
                    v96 := GenS11{F0: 717, F1: 69, F2: genFn0(18446744073709551615, 2)}
                    _ = v96
                } else {
                    v9 = v9
                    var v97 int64 = -1
                    v31 -= (1.5 - v31)
                    _ = v97
                }
                v9 = v9
            }
            if ((int32(v5) == 42) || false) {
                var v98 uint16 = 0
                _ = v98
            } else {
                if (0 > 128) {
                    v31 = ((0.0 + float32(1e10)) / v31)
                }
                var v99 float64 = 0.1
                _ = v99
            }
            _ = v32
        } else {
            _ = len(v3)
            if (v31 >= -2.25) {
                if true {
                    v5 = v5
                    v31 += (float32(1e10) - (((((0.0 * v31) / v31) / (v31 - float32((1e10 + float64(((float32(0.0) * v31) - v31)))))) + float32((0.0 + (((3.14159 * (((1.5 * 0.1) + 0.0) - float64(v31))) * (float64((v31 * v31)) / float64(float32(1e10)))) - (3.14159 - 1.5))))) / v31))
                    fmt.Println(v5, v31, v9)
                }
            }
            var v100 int64 = 794
            _ = v100
        }
        var v101 uint32 = v9
        var v102 string = "foo bar"
        _ = v101
        _ = v102
    } else {
        if false {
            switch ((("" + "\n\t") + "foo bar") + "foo bar") {
            case "a":
                if true {
                    for fi103 := 0; fi103 < 14; fi103++ {
                        if (!true) {
                            switch ("a" + "a") {
                            case "foo bar":
                                v104 := v4[v9]
                                _ = v104
                            case "hello":
                                v9 -= 187
                            default:
                                if (255 > 42) {
                                    switch fi103 {
                                    case 1:
                                        v9 += (v9 / (((v9 * 4294967295)) | 1))
                                    case 0:
                                        for fi105 := 0; fi105 < 7; fi105++ {
                                            var v106 bool = (true && (255 >= 2))
                                            var v107 string = "hello"
                                            switch (fi105 + (fi105 | (^(fi103)))) {
                                            case 673:
                                                switch v107 {
                                                case "":
                                                    v108 := GenS12{F0: 295}
                                                    _ = v108
                                                default:
                                                    if (42 != 2147483647) {
                                                        v9 = 0
                                                    }
                                                }
                                            case 32767:
                                                if (!v106) {
                                                    fmt.Println(v31, v9)
                                                    var v109 int32 = -1
                                                    if (!(true || ((fi103 >> ((65535) & 63)) > fi105))) {
                                                        v110 := GenS10{F0: v9, F1: 124, F2: v106}
                                                        _ = v110
                                                    } else {
                                                        v106 = (-684 >= 395)
                                                        var v111 int = fi103
                                                        var v112 uint64 = 4294967295
                                                        _ = v111
                                                        _ = v112
                                                    }
                                                    delete(v4, v9)
                                                    _ = v109
                                                }
                                            case -1:
                                                v113 := GenS3{F0: v9, F1: 0}
                                                _ = v113
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = fi105
                                            _ = v106
                                            _ = v107
                                        }
                                    case 9223372036854775807:
                                        v4[(v9 * 345)] = 32768
                                    }
                                    v9 = v9
                                    var v114 uint64 = 765
                                    _ = v114
                                } else {
                                    v3[("" + "\n\t")] = (^(v9))
                                }
                            }
                            _ = len(v4)
                            for fi115 := 0; fi115 < 12; fi115++ {
                                if (!true) {
                                    switch "foo bar" {
                                    case "":
                                        fmt.Println(v31)
                                    case "foo bar":
                                        if false {
                                            if (578 <= int64(v9)) {
                                                v9 -= v9
                                            }
                                            switch "a" {
                                            case "\n\t":
                                                var v116 string = "\n\t"
                                                _ = v116
                                            case "hello":
                                                v117 := GenS13{F0: 1, F1: "\n\t", F2: (fi115 ^ (fi115 % ((9223372036854775807) | 1)))}
                                                _ = v117
                                            default:
                                                if len(v2) > 0 {
                                                    v2[0] = 173
                                                }
                                            }
                                            var v118 int32 = 42
                                            _ = v118
                                        }
                                    case "a":
                                        var v119 uint = 1
                                        var v120 uint32 = v9
                                        _ = v119
                                        _ = v120
                                    default:
                                        v121 := v3["foo bar"]
                                        _ = v121
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = (v9 % (((v9 & 2147483648)) | 1))
                                    }
                                } else {
                                    fmt.Println(v9)
                                    if (1 >= 42) {
                                        v122 := GenS11{F0: -2147483648, F1: 42, F2: 32767}
                                        var v123 int16 = -931
                                        if true {
                                            v9 = (v9 / ((0) | 1))
                                            v31 += float32(float64(1.5))
                                            v123 = 1000
                                            if false {
                                                fmt.Println(v31, v5)
                                                v9 = (v9 / (((v9 >> ((601) & 63))) | 1))
                                            } else {
                                                var v124 int16 = (v123 & -32768)
                                                _ = v124
                                            }
                                            if true {
                                                v9 += v9
                                                var v125 float64 = (1.5 + (0.0 + (0.1 * 1e10)))
                                                if ((v9 | 2) == (v9 + v9)) {
                                                    v3[("" + "foo bar")] = (v9 % ((546) | 1))
                                                }
                                                _ = v125
                                            } else {
                                                if (false || ((true && false) && false)) {
                                                    var v126 bool = (!(!((!((-(v123)) > v123)) || (false || false))))
                                                    var v127 uint8 = 128
                                                    _ = v126
                                                    _ = v127
                                                } else {
                                                    v31 = 0.0
                                                }
                                            }
                                            switch (fi103 + (fi103 + -1)) {
                                            case 2147483647:
                                                if (v5 >= (^(v5))) {
                                                    var v128 float32 = 0.0
                                                    fmt.Println(v9)
                                                    _ = v128
                                                }
                                            case 281:
                                                var v129 float64 = ((3.14159 * (1.5 + -2.25)) * (0.1 + float64((-2.25 + 1.5))))
                                                _ = v129
                                            default:
                                                switch "hello" {
                                                case "a":
                                                    if false {
                                                        switch v5 {
                                                        case 42:
                                                            v9 = (v9 - (v9 & 42))
                                                        }
                                                        if false {
                                                            v123 = 32767
                                                        }
                                                        var v130 int16 = (v123 * v123)
                                                        var v131 int16 = v123
                                                        _ = v130
                                                        _ = v131
                                                    } else {
                                                        var v132 int16 = v123
                                                        _ = v132
                                                    }
                                                case "":
                                                    v133 := GenS3{F0: (v9 - (v9 + 4294967295)), F1: 966}
                                                    _ = v133
                                                case "hello":
                                                    var v134 int32 = 0
                                                    _ = v134
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = ("foo bar" + ("a" + ""))
                                                }
                                                v31 = v31
                                            }
                                            var v135 uint16 = 970
                                            _ = v135
                                        } else {
                                            switch "foo bar" {
                                            case "a":
                                                var v136 uint64 = 0
                                                _ = v136
                                            case "hello":
                                                v137 := GenS14{F0: 1, F1: (!false), F2: 403}
                                                _ = v137
                                            default:
                                                v123 = (v123 / ((v123) | 1))
                                            }
                                            v123 = (v123 << ((483) & 63))
                                        }
                                        _ = v122
                                        _ = v123
                                    }
                                    var v138 float32 = 3.14159
                                    _ = v138
                                }
                                if ((((-2.25 + 0.0) + v31) * 0.1) < float32(float64((3.14159 + float32(0.0))))) {
                                    if false {
                                        if (float64(v31) > (0.0 - float64(v31))) {
                                            if (!false) {
                                                var v139 int = fi115
                                                switch 1 {
                                                case 0:
                                                    v140 := GenS7{F0: 42}
                                                    _ = v140
                                                }
                                                _ = v139
                                            }
                                            v31 = v31
                                            v141 := GenS15{F0: 127}
                                            if (v9 == (v9 % (((v9 ^ 803)) | 1))) {
                                                _ = len(v4)
                                                delete(v3, "")
                                            } else {
                                                switch (^(fi115)) {
                                                case 2147483647:
                                                    var v142 int32 = -375
                                                    _ = v142
                                                case 4294967295:
                                                    switch fi103 {
                                                    case 2147483647:
                                                        if false {
                                                            if (true || ((!(true && (42 == 2))) || (false && (!false)))) {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v4[1] = fi103
                                                            }
                                                            var v143 float32 = (((((v31 * 1.5) + v31) + v31) * 3.14159) / v31)
                                                            v31 = float32(0.0)
                                                            _ = v143
                                                        }
                                                    case 42:
                                                        var v144 int16 = 1000
                                                        _ = v144
                                                    }
                                                case 32768:
                                                    fmt.Println(v31, fi103, v9)
                                                }
                                            }
                                            _ = v141
                                        }
                                        v31 += float32(((-2.25 * float64((float32(1e10) + (float32(0.1) - 0.1)))) * 1.5))
                                    }
                                    _ = len(v2)
                                    var v145 int32 = 2147483647
                                    var v146 uint8 = 255
                                    _ = v145
                                    _ = v146
                                } else {
                                    if (!(!true)) {
                                        v9 -= (^(v9))
                                        v4[v9] = 330
                                        if true {
                                            v147 := GenS8{F0: 600}
                                            _ = v147
                                        } else {
                                            fmt.Println(v5, v31)
                                        }
                                        if (32767 <= 415) {
                                            v31 = ((1e10 * (v31 + ((((1e10 - float32(float64(1e10))) / float32((1e10 - 3.14159))) - 0.1) * float32(0.1)))) + 0.0)
                                            v9 += v9
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = (v9 * (v9 - v9))
                                            }
                                            v148 := GenS16{F0: (542 > 18446744073709551615)}
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v148
                                        }
                                        if true {
                                            switch fi103 {
                                            case 32767:
                                                v9 = v9
                                                v9 = v9
                                            case 0:
                                                var v149 int16 = -1
                                                _ = v149
                                            default:
                                                var v150 int = (fi115 << ((143) & 63))
                                                _ = v150
                                            }
                                            switch -9223372036854775808 {
                                            case 811:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            default:
                                                if (true && ((18446744073709551615 > 1) || true)) {
                                                    var v151 int = 4294967295
                                                    switch (fi103 >> ((653) & 63)) {
                                                    case 32767:
                                                        fmt.Println(v9, v151, v9)
                                                    }
                                                    _ = v151
                                                } else {
                                                    v1 = append(v1, ("\n\t" + ""))
                                                }
                                            }
                                        } else {
                                            fmt.Println(fi115)
                                            if (!((372 != 792) || ((((v31 + (v31 - v31)) - (v31 * v31)) <= v31) || false))) {
                                                v152 := GenS5{F0: 1, F1: v5, F2: ("\n\t" + "foo bar")}
                                                _ = v152
                                            }
                                        }
                                    }
                                }
                                _ = fi115
                            }
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if true {
                            v5 = (fi103 >> ((65535) & 63))
                        }
                        _ = fi103
                    }
                    var v153 int = v5
                    var v154 bool = (v9 > v9)
                    _ = v153
                    _ = v154
                }
            case "foo bar":
                var v155 int8 = -1
                _ = v155
            case "\n\t":
                v9 = v9
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = len(v4)
    }
    v5 += v5
    fmt.Println(v5)
    for fi156 := 0; fi156 < 19; fi156++ {
        switch ("foo bar" + "foo bar") {
        case "":
            delete(v4, v9)
        default:
            v157 := GenS17{F0: 65535, F1: true, F2: int8(v9)}
            _ = v157
        }
        v158 := GenS18{F0: (true && false)}
        _ = fi156
        _ = v158
    }
    fmt.Println(v31, v5)
    var v159 float32 = v31
    v5 = (v5 | v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v159:", v159)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v159
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v9
}
