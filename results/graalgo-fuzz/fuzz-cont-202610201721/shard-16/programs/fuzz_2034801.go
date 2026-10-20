package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 bool; F2 int8 }
type GenS3 struct { F0 uint; F1 bool; F2 uint32 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 uint8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int32; F1 int32 }
func genFn0(p0 uint32) uint64 {
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{fi6, fi6}
        p0 -= p0
        switch (fi6 - fi6) {
        case -2147483648:
            for fi8 := 0; fi8 < 5; fi8++ {
                switch "hello" {
                case "a":
                    fmt.Println(p0, fi8)
                case "foo bar":
                    var v9 float64 = (3.14159 * (1.5 - (-2.25 * 3.14159)))
                    _ = v9
                case "\n\t":
                    var v10 uint = 18446744073709551615
                    _ = v10
                default:
                    var v11 float32 = 3.14159
                    _ = v11
                }
                if len(v7) > 0 {
                    v7[0] = fi6
                }
                _ = fi8
            }
        case 1:
            if len(v7) > 0 {
                v7[0] = fi6
            }
        case 4294967295:
            var v12 uint32 = p0
            _ = v12
        }
        _ = fi6
        _ = v7
    }
    v13 := GenS1{F0: 4294967295}
    var v14 int64 = -928
    _ = p0
    _ = v13
    _ = v14
    return 0
}

func main() {
    v1 := []string{"a", "", ("foo bar" + ("a" + "foo bar"))}
    var v2 float32 = float32(0.1)
    var v3 string = "hello"
    var v4 uint64 = 0
    var v5 uint = 1
    v15 := GenS2{F0: 45, F1: (v2 > 0.0), F2: 0}
    fmt.Println(v2, v4)
    switch 9223372036854775807 {
    case 32767:
        var v16 int8 = -93
        _ = v16
    case -1:
        _ = len(v1)
    case 993:
        var v17 int = -372
        _ = v17
    default:
        v3 = "foo bar"
    }
    if true {
        v1 = append(v1, v3)
        if false {
            var v18 string = "\n\t"
            var v19 bool = false
            for fi20 := 0; fi20 < 17; fi20++ {
                for fi21 := 0; fi21 < 18; fi21++ {
                    fmt.Println(v18, v18, fi20)
                    if (v19 && ((v19 && true) || v19)) {
                        switch fi20 {
                        case 32767:
                            switch (v18 + v3) {
                            case "":
                                var v22 int64 = -9223372036854775808
                                _ = v22
                            }
                        default:
                            v4 += v4
                        }
                    }
                    _ = fi21
                }
                _ = fi20
            }
            _ = v18
            _ = v19
        }
        switch 127 {
        case 697:
            if (!true) {
                for fi23 := 0; fi23 < 9; fi23++ {
                    var v24 float32 = v2
                    _ = fi23
                    _ = v24
                }
                _ = len(v1)
            } else {
                v25 := GenS1{F0: v5}
                v4 = genFn0(987)
                _ = v25
            }
        case -2147483648:
            var v26 uint64 = v4
            _ = v26
        default:
            if (false && (!(!true))) {
                fmt.Println(v5)
                if len(v1) > 0 {
                    v1[0] = "a"
                }
            }
        }
    } else {
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    }
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        for fi28 := 0; fi28 < 6; fi28++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v4 = v4
            _ = fi28
        }
        switch v3 {
        case "\n\t":
            v29 := GenS1{F0: 255}
            _ = v29
        default:
            v3 = (v3 + v3)
        }
        v30 := GenS2{F0: 42, F1: ((v2 / (v2 * v2)) != v2), F2: 0}
        _ = fi27
        _ = v30
    }
    switch -9223372036854775808 {
    case 127:
        v31 := GenS3{F0: (v5 - v5), F1: false, F2: 2}
        _ = v31
    case -9223372036854775808:
        _ = len(v1)
    case 4294967295:
        if (0 != 2147483647) {
            if ((v5 >> (((v5 & (v5 * v5))) & 63)) == 18446744073709551615) {
                switch 396 {
                case -2147483648:
                    v1 = append(v1, v3)
                case 2147483647:
                    for fi32 := 0; fi32 < 2; fi32++ {
                        if false {
                            fmt.Println(v2, fi32)
                            v3 = v3
                            v4 = 113
                        } else {
                            for fi33 := 0; fi33 < 5; fi33++ {
                                _ = len(v1)
                                fmt.Println(v2)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v34 int32 = -91
                                fmt.Println(v5, v34)
                                _ = fi33
                                _ = v34
                            }
                        }
                        v5 += 882
                        v35 := GenS1{F0: v5}
                        _ = fi32
                        _ = v35
                    }
                case 362:
                    v36 := GenS2{F0: -128, F1: (v5 < v5), F2: 0}
                    _ = v36
                }
                v37 := GenS4{F0: 241, F1: -1}
                var v38 uint8 = uint8(v4)
                _ = v37
                _ = v38
            } else {
                switch 127 {
                case 128:
                    v4 += (v4 ^ v4)
                case -2147483648:
                    if true {
                        for fi39 := 0; fi39 < 3; fi39++ {
                            var v40 uint = 4294967295
                            if false {
                                var v41 float32 = (v2 + 0.0)
                                _ = v41
                            }
                            _ = fi39
                            _ = v40
                        }
                    } else {
                        if true {
                            if true {
                                v42 := GenS1{F0: (^(v5))}
                                switch -1 {
                                case 128:
                                    v4 -= 2
                                case 4294967295:
                                    v5 += 271
                                default:
                                    switch ((v3 + "foo bar") + (("hello" + v3) + "foo bar")) {
                                    case "\n\t":
                                        switch -542 {
                                        case 9223372036854775807:
                                            v4 = 18446744073709551615
                                        default:
                                            switch 128 {
                                            case 9223372036854775807:
                                                switch (v3 + v3) {
                                                case "\n\t":
                                                    switch "hello" {
                                                    case "foo bar":
                                                        if true {
                                                            var v43 int8 = 1
                                                            if false {
                                                                for fi44 := 0; fi44 < 2; fi44++ {
                                                                    v45 := GenS3{F0: v5, F1: true, F2: 922}
                                                                    _ = fi44
                                                                    _ = v45
                                                                }
                                                            }
                                                            v2 -= v2
                                                            _ = v43
                                                        }
                                                    default:
                                                        if len(v1) > 0 {
                                                            v1[0] = "a"
                                                        }
                                                    }
                                                }
                                            case 4294967295:
                                                var v46 string = ""
                                                _ = v46
                                            case 1:
                                                _ = len(v1)
                                            default:
                                                _ = len(v1)
                                            }
                                        }
                                    case "":
                                        if len(v1) > 0 {
                                            v1[0] = (v3 + v3)
                                        }
                                    case "hello":
                                        v2 += (v2 / (float32(3.14159) / v2))
                                    }
                                }
                                _ = v42
                            }
                            for fi47 := 0; fi47 < 2; fi47++ {
                                switch fi47 {
                                case 128:
                                    var v48 int16 = 42
                                    _ = v48
                                default:
                                    switch v3 {
                                    case "\n\t":
                                        if true {
                                            v3 = ((v3 + v3) + "a")
                                        }
                                    }
                                }
                                for fi49 := 0; fi49 < 14; fi49++ {
                                    v4 = genFn0(583)
                                    var v50 int8 = 100
                                    if false {
                                        fmt.Println(fi49, v3, v4)
                                    } else {
                                        var v51 uint64 = (^(v4))
                                        _ = v51
                                    }
                                    if false {
                                        switch fi47 {
                                        case -2147483648:
                                            v3 = v3
                                        case 32767:
                                            v52 := GenS2{F0: (v50 / ((v50) | 1)), F1: (fi49 >= (fi49 ^ (fi47 >> ((228) & 63)))), F2: (^(v50))}
                                            _ = v52
                                        }
                                        if false {
                                            if len(v1) > 0 {
                                                v1[0] = (v3 + (v3 + (v3 + v3)))
                                            }
                                            v4 = (v4 / (((^(v4))) | 1))
                                            v4 = genFn0(809)
                                        } else {
                                            if (65535 > (v5 + (v5 | v5))) {
                                                switch -1 {
                                                case 42:
                                                    var v53 uint32 = 1
                                                    _ = v53
                                                case 928:
                                                    switch "a" {
                                                    case "a":
                                                        _ = len(v1)
                                                    case "":
                                                        v54 := GenS4{F0: -1, F1: 1}
                                                        _ = v54
                                                    default:
                                                        switch "a" {
                                                        case "":
                                                            var v55 uint32 = 42
                                                            _ = v55
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        var v56 bool = true
                                                        fmt.Println(v50, v3)
                                                        var v57 uint16 = 691
                                                        v58 := GenS4{F0: 1, F1: 42}
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                }
                                                if true {
                                                    if ((!false) || (1 < -537)) {
                                                        v59 := GenS5{F0: (4294967295 >= 2147483648)}
                                                        _ = v59
                                                    } else {
                                                        v2 -= ((v2 * (v2 - 0.0)) - v2)
                                                        v60 := GenS6{F0: 296}
                                                        var v61 uint8 = 146
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                    var v62 int = int(v50)
                                                    v63 := GenS6{F0: 0}
                                                    _ = v62
                                                    _ = v63
                                                } else {
                                                    fmt.Println(v50, v4)
                                                }
                                                var v64 uint64 = (v4 / ((v4) | 1))
                                                if (488 > int64(fi49)) {
                                                    var v65 uint64 = 276
                                                    _ = v65
                                                }
                                                _ = v64
                                            } else {
                                                v66 := GenS7{F0: 125, F1: 57}
                                                v4 -= genFn0(694)
                                                _ = v66
                                            }
                                            if false {
                                                v50 = v50
                                            }
                                        }
                                        var v67 int16 = int16(v4)
                                        _ = v67
                                    } else {
                                        switch -99 {
                                        case 755:
                                            v3 = v3
                                        case 2147483647:
                                            fmt.Println(fi49, v3, v2)
                                        default:
                                            v68 := GenS2{F0: v50, F1: (!(42 > 0)), F2: v50}
                                            _ = v68
                                        }
                                        var v69 float32 = v2
                                        _ = v69
                                    }
                                    if (385 >= 42) {
                                        fmt.Println(fi49, v50, v3)
                                    }
                                    _ = fi49
                                    _ = v50
                                }
                                var v70 int32 = 1
                                _ = fi47
                                _ = v70
                            }
                            _ = len(v1)
                        } else {
                            v1 = append(v1, v3)
                        }
                    }
                case 669:
                    if true {
                        v1 = append(v1, "\n\t")
                        var v71 uint8 = 0
                        var v72 int32 = -1
                        _ = v71
                        _ = v72
                    } else {
                        fmt.Println(v4)
                        for fi73 := 0; fi73 < 13; fi73++ {
                            switch 152 {
                            case 1:
                                fmt.Println(v3, v3)
                            }
                            switch ((v3 + v3) + v3) {
                            case "a":
                                for fi74 := 0; fi74 < 20; fi74++ {
                                    var v75 int = 32767
                                    fmt.Println(v2, fi74)
                                    _ = fi74
                                    _ = v75
                                }
                            case "foo bar":
                                if (float64(v2) < ((3.14159 / float64(float32(3.14159))) - 1.5)) {
                                    for fi76 := 0; fi76 < 9; fi76++ {
                                        v3 = v3
                                        v77 := GenS1{F0: 999}
                                        v78 := GenS8{F0: false}
                                        var v79 int8 = 127
                                        _ = fi76
                                        _ = v77
                                        _ = v78
                                        _ = v79
                                    }
                                } else {
                                    for fi80 := 0; fi80 < 10; fi80++ {
                                        if (-51 <= -1) {
                                            if true {
                                                var v81 int64 = 9223372036854775807
                                                _ = v81
                                            }
                                        } else {
                                            v2 -= (v2 * 0.1)
                                            v2 = v2
                                            if true {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                switch ((v3 + (v3 + (v3 + v3))) + "a") {
                                                case "a":
                                                    switch v3 {
                                                    case "foo bar":
                                                        v5 = v5
                                                    case "":
                                                        fmt.Println(v2, fi73, fi73)
                                                    default:
                                                        if len(v1) > 0 {
                                                            v1[0] = v3
                                                        }
                                                    }
                                                case "hello":
                                                    var v82 int = (fi80 / ((fi80) | 1))
                                                    _ = v82
                                                }
                                            }
                                        }
                                        var v83 uint16 = 65535
                                        v3 = v3
                                        _ = fi80
                                        _ = v83
                                    }
                                    switch (((v3 + (((((v3 + (("foo bar" + (((v3 + "\n\t") + ("hello" + "hello")) + v3)) + (v3 + ""))) + (v3 + "hello")) + v3) + "hello") + (v3 + ((v3 + v3) + "\n\t")))) + (v3 + (("hello" + ("foo bar" + ((v3 + v3) + (v3 + "a")))) + v3))) + v3) {
                                    case "hello":
                                        fmt.Println(v5, v2, v3)
                                    case "\n\t":
                                        if false {
                                            var v84 uint64 = (v4 & v4)
                                            v85 := GenS9{F0: 965, F1: 0}
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v84
                                            _ = v85
                                        } else {
                                            var v86 int32 = 0
                                            var v87 float64 = 1e10
                                            _ = v86
                                            _ = v87
                                        }
                                    }
                                }
                            case "\n\t":
                                if (!(!true)) {
                                    for fi88 := 0; fi88 < 13; fi88++ {
                                        if false {
                                            if len(v1) > 0 {
                                                v1[0] = v3
                                            }
                                            if (42 != -128) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = len(v1)
                                                v2 = (v2 - v2)
                                            }
                                        }
                                        _ = fi88
                                    }
                                    var v89 float64 = 0.1
                                    _ = v89
                                } else {
                                    if true {
                                        v5 = 638
                                        fmt.Println(v3, v4, v2)
                                    } else {
                                        var v90 string = v3
                                        v3 = v3
                                        if (42 <= -1) {
                                            var v91 int16 = 42
                                            _ = v91
                                        }
                                        _ = v90
                                    }
                                    v3 = v3
                                }
                            default:
                                var v92 uint = v5
                                _ = v92
                            }
                            fmt.Println(v3, fi73)
                            _ = fi73
                        }
                    }
                default:
                    for fi93 := 0; fi93 < 1; fi93++ {
                        for fi94 := 0; fi94 < 4; fi94++ {
                            v5 = (v5 >> ((317) & 63))
                            v95 := GenS3{F0: v5, F1: false, F2: 266}
                            fmt.Println(fi94, fi93, fi93)
                            _ = fi94
                            _ = v95
                        }
                        v5 = v5
                        _ = fi93
                    }
                }
            }
        }
    default:
        for fi96 := 0; fi96 < 13; fi96++ {
            var v97 string = "\n\t"
            var v98 uint64 = v4
            _ = fi96
            _ = v97
            _ = v98
        }
    }
    v3 = v3
    v3 = v3
    v99 := GenS3{F0: (v5 & v5), F1: true, F2: 56}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = ((((v3 + v3) + ("foo bar" + "foo bar")) + (v3 + v3)) + v3)
    if ((v3 > (("a" + v3) + "")) && (-1 <= -1)) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        if ((!((v4 - (v4 / ((uint64(v5)) | 1))) < v4)) || false) {
            v100 := GenS3{F0: v5, F1: false, F2: 2147483647}
            _ = v100
        } else {
            var v101 int = 42
            for fi102 := 0; fi102 < 6; fi102++ {
                fmt.Println(v2, v4)
                var v103 float32 = (((((v2 - v2) * float32(1e10)) * 0.1) / v2) / (v2 * float32(1.5)))
                var v104 uint = (^(v5))
                _ = fi102
                _ = v103
                _ = v104
            }
            _ = v101
        }
        v3 = v3
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v99
}
