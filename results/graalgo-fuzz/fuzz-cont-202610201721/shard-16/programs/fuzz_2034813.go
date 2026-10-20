package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 uint32; F1 int64 }
type GenS8 struct { F0 int16; F1 bool }
type GenS9 struct { F0 string; F1 uint8 }
type GenS10 struct { F0 bool; F1 uint8; F2 int }
type GenS11 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS12 struct { F0 string; F1 bool }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int32; F1 uint32; F2 int64 }
func genFn0() bool {
    var v7 int16 = -7
    v8 := GenS1{F0: (v7 >> ((4294967295) & 63)), F1: (v7 / ((-1) | 1)), F2: 255}
    v7 -= (v7 << ((876) & 63))
    var v9 uint32 = 33
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 uint = 2
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    var v4 bool = (false || (42 != 0))
    var v5 int32 = -1
    v6 := make(map[string]uint32)
    if (true || (!(((1.5 / 1.0) * ((3.14159 + 0.0) / 3.14159)) > (3.14159 * (0.1 + 0.1))))) {
        v10 := GenS1{F0: 1000, F1: -880, F2: 127}
        _ = v10
    } else {
        var v11 bool = (244 != 4294967295)
        v11 = (v4 || v4)
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 uint8 = 255
        _ = len(v6)
        v14 := []int{}
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint8 = 127
        v1 = v1
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int16 = 0
            var v19 int = fi17
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    v20 := GenS2{F0: 65535, F1: -87}
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := []int{fi21}
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 float64 = 1.5
            var v25 bool = (!v4)
            v26 := GenS2{F0: 42, F1: 1}
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        switch "a" {
        case "\n\t":
            fmt.Println(v4, v5)
        case "hello":
            var v27 float64 = float64(1e10)
            _ = v27
        default:
            switch "a" {
            case "":
                if true {
                    v1 = 298
                } else {
                    var v28 string = ("a" + "foo bar")
                    _ = v28
                }
            default:
                if v4 {
                    v5 += v5
                    switch "a" {
                    case "":
                        v4 = ((v5 + v5) == v5)
                    }
                } else {
                    v29 := GenS3{F0: 607, F1: -1}
                    _ = v29
                }
            }
        }
        _ = fi21
        _ = v22
    }
    v30 := []int{42, 0, 0, 1}
    for fi31 := 0; fi31 < 4; fi31++ {
        var v32 uint8 = 0
        v4 = (!(v4 && v4))
        _ = len(v3)
        _ = fi31
        _ = v32
    }
    v1 -= (v1 - (v1 & 358))
    fmt.Println(v1, v4)
    fmt.Println(v5, v5, v1)
    v33 := GenS1{F0: 32767, F1: 1000, F2: 1}
    switch "hello" {
    case "":
        for fi34 := 0; fi34 < 14; fi34++ {
            v4 = (45 != 1)
            for fi35 := 0; fi35 < 10; fi35++ {
                if true {
                    var v36 float64 = (0.1 - 1.5)
                    if len(v30) > 0 {
                        v30[0] = fi35
                    }
                    if len(v30) >= 2 {
                        v30 = v30[1:len(v30)]
                    }
                    _ = v36
                }
                _ = fi35
            }
            if false {
                var v37 string = "foo bar"
                _ = v37
            } else {
                var v38 int16 = 42
                var v39 string = "hello"
                for fi40 := 0; fi40 < 2; fi40++ {
                    if v4 {
                        v1 = 882
                        switch v39 {
                        case "foo bar":
                            v41 := v2[v39]
                            _ = v41
                        case "hello":
                            if true {
                                if v4 {
                                    v4 = ((v38 >> (((v1 + (v1 << (((v1 >> ((v1) & 63))) & 63)))) & 63)) < v38)
                                    v1 = v1
                                }
                                if (1 < 42) {
                                    var v42 uint32 = uint32(v1)
                                    var v43 string = "a"
                                    switch v39 {
                                    case "hello":
                                        v44 := GenS4{F0: v38, F1: 42, F2: v1}
                                        var v45 uint32 = (v42 ^ v42)
                                        if false {
                                            if v4 {
                                                if (((1.5 * 1e10) + 0.1) == -2.25) {
                                                    v42 = (v45 | (^(v42)))
                                                    v39 = (v43 + "hello")
                                                }
                                                var v46 float32 = 1.5
                                                _ = v46
                                            }
                                            switch int(v1) {
                                            case 128:
                                                var v47 string = "a"
                                                _ = v47
                                            case 469:
                                                v48 := GenS5{F0: -128}
                                                _ = v48
                                            case 42:
                                                switch v39 {
                                                case "a":
                                                    if v4 {
                                                        var v49 uint32 = 236
                                                        if (false || v4) {
                                                            v42 = v42
                                                        } else {
                                                            v50 := GenS5{F0: -121}
                                                            if v4 {
                                                                var v51 uint = 65535
                                                                var v52 int8 = -128
                                                                _ = v51
                                                                _ = v52
                                                            } else {
                                                                v1 = 93
                                                                switch fi40 {
                                                                case 349:
                                                                    v30 = append(v30, -879)
                                                                    v30 = append(v30, fi34)
                                                                case 32767:
                                                                    switch v43 {
                                                                    case "a":
                                                                        switch (fi34 - 1) {
                                                                        case -512:
                                                                            if v4 {
                                                                                _ = len(v30)
                                                                                var v53 string = v43
                                                                                var v54 int16 = v38
                                                                                if v4 {
                                                                                    v1 -= uint(v42)
                                                                                    _ = len(v2)
                                                                                } else {
                                                                                    var v55 int16 = (v54 << ((v1) & 63))
                                                                                    v56 := GenS1{F0: v55, F1: (v38 - v55), F2: 128}
                                                                                    switch "hello" {
                                                                                    case "\n\t":
                                                                                        if (1 >= 1) {
                                                                                            switch v39 {
                                                                                            case "\n\t":
                                                                                                if (v4 || v4) {
                                                                                                    _ = len(v30)
                                                                                                    if len(v30) >= 2 {
                                                                                                        v30 = v30[1:len(v30)]
                                                                                                    }
                                                                                                    fmt.Println(v54, v55, v54)
                                                                                                } else {
                                                                                                    v57 := GenS6{F0: 3, F1: 114}
                                                                                                    switch (fi40 >> ((0) & 63)) {
                                                                                                    case 127:
                                                                                                        fmt.Println(v38, v5, v39)
                                                                                                        v58 := GenS3{F0: (v42 | v45), F1: (-(v5))}
                                                                                                        if (!(v55 <= int16(fi34))) {
                                                                                                            var v59 int64 = int64(v5)
                                                                                                            var v60 bool = v4
                                                                                                            switch "hello" {
                                                                                                            case "":
                                                                                                                if (v4 || (v38 > (v54 << ((v1) & 63)))) {
                                                                                                                    fmt.Println(v43, fi34, v54)
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                if v60 {
                                                                                                                    switch (v53 + v53) {
                                                                                                                    case "\n\t":
                                                                                                                        if (!v4) {
                                                                                                                            var v61 int64 = -9223372036854775808
                                                                                                                            fmt.Println(v1, fi34, v60)
                                                                                                                            _ = v61
                                                                                                                        }
                                                                                                                    case "":
                                                                                                                        switch ((v53 + (v43 + v39)) + ("" + ((v53 + (((v43 + v43) + v39) + "a")) + "\n\t"))) {
                                                                                                                        case "\n\t":
                                                                                                                            if (0.1 > (0.0 - 1e10)) {
                                                                                                                                fmt.Println(v54)
                                                                                                                                var v62 uint64 = 1
                                                                                                                                _ = v62
                                                                                                                            } else {
                                                                                                                                var v63 uint16 = uint16(v49)
                                                                                                                                _ = v63
                                                                                                                            }
                                                                                                                        case "":
                                                                                                                            var v64 bool = (v4 || ((((1.5 * -2.25) <= 3.14159) || v4) || v60))
                                                                                                                            _ = v64
                                                                                                                        case "foo bar":
                                                                                                                            switch "foo bar" {
                                                                                                                            case "\n\t":
                                                                                                                                switch fi40 {
                                                                                                                                case 4294967295:
                                                                                                                                    v39 = (("foo bar" + v39) + "a")
                                                                                                                                default:
                                                                                                                                    fmt.Println(v60)
                                                                                                                                }
                                                                                                                            case "a":
                                                                                                                                v39 = v43
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            v65 := GenS4{F0: v38, F1: (v42 >> ((v1) & 63)), F2: 553}
                                                                                                                            _ = v65
                                                                                                                        }
                                                                                                                    case "hello":
                                                                                                                        if (!v60) {
                                                                                                                            var v66 uint = v1
                                                                                                                            if (255 <= 216) {
                                                                                                                                v66 = (v66 * v66)
                                                                                                                                var v67 string = (v53 + v43)
                                                                                                                                _ = v67
                                                                                                                            } else {
                                                                                                                                var v68 uint64 = 4294967295
                                                                                                                                v45 -= v42
                                                                                                                                _ = v68
                                                                                                                            }
                                                                                                                            fmt.Println(v1, v42)
                                                                                                                            _ = v66
                                                                                                                        } else {
                                                                                                                            switch v39 {
                                                                                                                            case "\n\t":
                                                                                                                                fmt.Println(v38, v38)
                                                                                                                            case "foo bar":
                                                                                                                                v43 = ""
                                                                                                                            case "":
                                                                                                                                v5 = v5
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    if true {
                                                                                                                        if (int8(v59) >= -57) {
                                                                                                                            v30 = append(v30, fi34)
                                                                                                                            v3[42] = false
                                                                                                                        } else {
                                                                                                                            v1 = v1
                                                                                                                        }
                                                                                                                    }
                                                                                                                    switch v53 {
                                                                                                                    case "":
                                                                                                                        var v69 int16 = v38
                                                                                                                        _ = v69
                                                                                                                    case "\n\t":
                                                                                                                        var v70 int32 = (v5 | v5)
                                                                                                                        _ = v70
                                                                                                                    default:
                                                                                                                        v43 = ((v53 + v53) + v39)
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    if v4 {
                                                                                                                        delete(v2, v39)
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                if v60 {
                                                                                                                    switch "a" {
                                                                                                                    case "\n\t":
                                                                                                                        v5 = (v5 & v5)
                                                                                                                    case "foo bar":
                                                                                                                        v55 = (v54 | v55)
                                                                                                                    default:
                                                                                                                        v45 += 548
                                                                                                                    }
                                                                                                                    _ = len(v30)
                                                                                                                    v53 = v53
                                                                                                                    switch -279 {
                                                                                                                    case 42:
                                                                                                                        _ = len(v2)
                                                                                                                    case 0:
                                                                                                                        v71 := GenS2{F0: 1, F1: 127}
                                                                                                                        _ = v71
                                                                                                                    case 734:
                                                                                                                        if v4 {
                                                                                                                            var v72 uint16 = 42
                                                                                                                            _ = v72
                                                                                                                        } else {
                                                                                                                            var v73 float32 = ((1e10 + 0.1) + (1.5 + ((0.0 / 3.14159) / 0.1)))
                                                                                                                            v74 := GenS7{F0: (^(v45)), F1: (v59 * v59)}
                                                                                                                            _ = v73
                                                                                                                            _ = v74
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            v53 = v43
                                                                                                            _ = v59
                                                                                                            _ = v60
                                                                                                        }
                                                                                                        _ = v58
                                                                                                    case 32767:
                                                                                                        v75 := GenS7{F0: v49, F1: -1}
                                                                                                        _ = v75
                                                                                                    }
                                                                                                    _ = v57
                                                                                                }
                                                                                            case "a":
                                                                                                v49 -= 0
                                                                                            }
                                                                                            v1 += 989
                                                                                            v76 := v6[v43]
                                                                                            switch fi40 {
                                                                                            case 32767:
                                                                                                if (1 != (fi40 * -943)) {
                                                                                                    v45 += v76
                                                                                                } else {
                                                                                                    var v77 int16 = (v38 % ((v54) | 1))
                                                                                                    fmt.Println(v39)
                                                                                                    _ = v77
                                                                                                }
                                                                                            case 4294967295:
                                                                                                var v78 float64 = 1e10
                                                                                                _ = v78
                                                                                            default:
                                                                                                if (v4 && v4) {
                                                                                                    switch (fi34 % ((4294967295) | 1)) {
                                                                                                    case -1:
                                                                                                        switch (fi34 - 127) {
                                                                                                        case -212:
                                                                                                            v5 -= v5
                                                                                                        case 127:
                                                                                                            v5 = 1
                                                                                                        default:
                                                                                                            var v79 int = fi40
                                                                                                            _ = v79
                                                                                                        }
                                                                                                    case 127:
                                                                                                        _ = len(v2)
                                                                                                    }
                                                                                                    if (1 != 2) {
                                                                                                        if ((fi34 & 9223372036854775807) < fi40) {
                                                                                                            var v80 int16 = (v54 - (v55 - v38))
                                                                                                            _ = v80
                                                                                                        }
                                                                                                        _ = len(v30)
                                                                                                        var v81 uint64 = 18446744073709551615
                                                                                                        _ = v81
                                                                                                    }
                                                                                                } else {
                                                                                                    v82 := GenS8{F0: (v55 * (v54 / ((v38) | 1))), F1: (0.1 == (-2.25 + ((0.0 / 1.0) - 1.5)))}
                                                                                                    if (!(-9223372036854775808 < 269)) {
                                                                                                        var v83 bool = (v39 == (v53 + ((("\n\t" + v39) + v53) + ("" + v39))))
                                                                                                        if (0.0 == 1.5) {
                                                                                                            if len(v30) >= 2 {
                                                                                                                v30 = v30[1:len(v30)]
                                                                                                            }
                                                                                                            v84 := GenS6{F0: 255, F1: 990}
                                                                                                            _ = v84
                                                                                                        } else {
                                                                                                            if len(v30) >= 2 {
                                                                                                                v30 = v30[1:len(v30)]
                                                                                                            }
                                                                                                            _ = len(v6)
                                                                                                            var v85 float32 = 1.5
                                                                                                            _ = v85
                                                                                                        }
                                                                                                        v39 = v53
                                                                                                        _ = v83
                                                                                                    } else {
                                                                                                        var v86 float64 = float64(float32((3.14159 - -2.25)))
                                                                                                        _ = v86
                                                                                                    }
                                                                                                    v76 = v42
                                                                                                    v54 += -1
                                                                                                    _ = v82
                                                                                                }
                                                                                            }
                                                                                            v5 += (^(v5))
                                                                                            _ = v76
                                                                                        } else {
                                                                                            v53 = v53
                                                                                            v38 += (^(v38))
                                                                                            var v87 int16 = v38
                                                                                            _ = v87
                                                                                        }
                                                                                    case "hello":
                                                                                        if v4 {
                                                                                            v45 = v42
                                                                                            v4 = (v42 != (v49 | (v49 / (((v42 >> ((v1) & 63))) | 1))))
                                                                                            if (v4 && ((fi34 * (fi34 ^ -220)) < 32767)) {
                                                                                                var v88 float32 = (1e10 / 1.5)
                                                                                                var v89 int32 = -2147483648
                                                                                                v30 = append(v30, int(v38))
                                                                                                var v90 int = (fi34 ^ (-(fi34)))
                                                                                                if v4 {
                                                                                                    switch 946 {
                                                                                                    case 2147483647:
                                                                                                        v91 := v6[((v53 + ("foo bar" + v43)) + ("\n\t" + "\n\t"))]
                                                                                                        _ = v91
                                                                                                    case 127:
                                                                                                        v43 = v39
                                                                                                    }
                                                                                                    var v92 uint32 = 0
                                                                                                    _ = v92
                                                                                                }
                                                                                                _ = v88
                                                                                                _ = v89
                                                                                                _ = v90
                                                                                            }
                                                                                        } else {
                                                                                            v54 = 1000
                                                                                        }
                                                                                    default:
                                                                                        _ = len(v2)
                                                                                    }
                                                                                    _ = v55
                                                                                    _ = v56
                                                                                }
                                                                                fmt.Println(v38, fi34, fi34)
                                                                                _ = v53
                                                                                _ = v54
                                                                            } else {
                                                                                var v93 uint32 = v49
                                                                                switch v43 {
                                                                                case "hello":
                                                                                    if ((9223372036854775807 < fi34) || (127 > -102)) {
                                                                                        v94 := GenS8{F0: 0, F1: v4}
                                                                                        if (0.1 >= 3.14159) {
                                                                                            var v95 bool = false
                                                                                            var v96 uint16 = 32767
                                                                                            _ = v95
                                                                                            _ = v96
                                                                                        } else {
                                                                                            var v97 uint8 = 0
                                                                                            if (float32(((1.5 / (-2.25 * (float64((float32((-2.25 - -2.25)) + ((0.0 - 1.5) / 1e10))) + (-2.25 / (((3.14159 + ((1e10 - -2.25) / 3.14159)) + 0.1) * (((-2.25 / 1e10) / -2.25) - 1.5)))))) / 0.1)) <= (-2.25 + 1.5)) {
                                                                                                v98 := GenS8{F0: (v38 | v38), F1: false}
                                                                                                switch v43 {
                                                                                                case "":
                                                                                                    switch ((v39 + (((("hello" + v43) + v43) + v39) + v43)) + (v39 + (("a" + (v43 + v43)) + v43))) {
                                                                                                    case "hello":
                                                                                                        delete(v6, v39)
                                                                                                    default:
                                                                                                        var v99 string = "foo bar"
                                                                                                        _ = v99
                                                                                                    }
                                                                                                    fmt.Println(v5, fi34, v1)
                                                                                                    v43 = "a"
                                                                                                case "a":
                                                                                                    switch v39 {
                                                                                                    case "\n\t":
                                                                                                        switch 127 {
                                                                                                        case 281:
                                                                                                            switch (^(fi40)) {
                                                                                                            case 127:
                                                                                                                v39 = v43
                                                                                                            case -2147483648:
                                                                                                                var v100 string = v43
                                                                                                                _ = v100
                                                                                                            case -9223372036854775808:
                                                                                                                v101 := GenS9{F0: v39, F1: v97}
                                                                                                                _ = v101
                                                                                                            default:
                                                                                                                var v102 int16 = 32767
                                                                                                                _ = v102
                                                                                                            }
                                                                                                        }
                                                                                                    case "":
                                                                                                        v5 = 0
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    if true {
                                                                                                        if v4 {
                                                                                                            v103 := GenS8{F0: (v38 >> ((v1) & 63)), F1: v4}
                                                                                                            v42 += v93
                                                                                                            var v104 bool = v4
                                                                                                            switch v43 {
                                                                                                            case "\n\t":
                                                                                                                v2[v39] = true
                                                                                                            case "hello":
                                                                                                                var v105 float32 = ((3.14159 * (0.0 / (0.1 - (float32(float64((3.14159 / 1.0))) + (3.14159 / 1e10))))) - (3.14159 * ((0.1 / 0.1) - 1e10)))
                                                                                                                _ = v105
                                                                                                            default:
                                                                                                                v106 := GenS9{F0: v39, F1: v97}
                                                                                                                _ = v106
                                                                                                            }
                                                                                                            var v107 int64 = -1
                                                                                                            _ = v103
                                                                                                            _ = v104
                                                                                                            _ = v107
                                                                                                        }
                                                                                                        v38 = 931
                                                                                                        v1 = (v1 ^ (^(v1)))
                                                                                                    }
                                                                                                default:
                                                                                                    v49 -= v93
                                                                                                }
                                                                                                _ = v98
                                                                                            }
                                                                                            v108 := GenS10{F0: v4, F1: 104, F2: 2147483647}
                                                                                            v109 := v6[("hello" + v39)]
                                                                                            _ = v108
                                                                                            _ = v109
                                                                                            _ = v97
                                                                                        }
                                                                                        var v110 uint64 = 0
                                                                                        _ = v110
                                                                                        _ = v94
                                                                                    } else {
                                                                                        var v111 uint16 = 445
                                                                                        _ = len(v30)
                                                                                        v43 = ((v43 + "foo bar") + v39)
                                                                                        _ = v111
                                                                                    }
                                                                                case "\n\t":
                                                                                    v5 -= v5
                                                                                }
                                                                                _ = v93
                                                                            }
                                                                        default:
                                                                            var v112 int8 = 0
                                                                            _ = v112
                                                                        }
                                                                    case "foo bar":
                                                                        v113 := GenS7{F0: uint32(v1), F1: 1}
                                                                        _ = v113
                                                                    case "":
                                                                        if (int8(fi34) < 100) {
                                                                            v4 = v4
                                                                        } else {
                                                                            var v114 bool = ((v38 % ((v38) | 1)) > v38)
                                                                            _ = v114
                                                                        }
                                                                    default:
                                                                        var v115 bool = v4
                                                                        _ = v115
                                                                    }
                                                                default:
                                                                    fmt.Println(v49, v42)
                                                                }
                                                            }
                                                            _ = v50
                                                        }
                                                        var v116 int32 = v5
                                                        _ = v116
                                                        _ = v49
                                                    } else {
                                                        switch (v39 + v39) {
                                                        case "hello":
                                                            if v4 {
                                                                var v117 uint32 = (v45 / ((0) | 1))
                                                                v118 := GenS10{F0: ((v38 & v38) <= (^(v38))), F1: 127, F2: (fi34 - fi40)}
                                                                _ = v117
                                                                _ = v118
                                                            }
                                                        case "a":
                                                            v119 := GenS11{F0: 1, F1: 0, F2: 604}
                                                            _ = v119
                                                        }
                                                        var v120 int = fi34
                                                        _ = v120
                                                    }
                                                case "\n\t":
                                                    v5 += (v5 ^ (v5 ^ v5))
                                                    var v121 int64 = int64(v5)
                                                    _ = v121
                                                case "hello":
                                                    v122 := GenS3{F0: (^(v42)), F1: (v5 - int32(v38))}
                                                    _ = v122
                                                }
                                            default:
                                                v42 -= (v42 - 122)
                                            }
                                        }
                                        _ = v44
                                        _ = v45
                                    }
                                    _ = len(v6)
                                    _ = v42
                                    _ = v43
                                } else {
                                    v38 -= (v38 + 0)
                                    v123 := GenS12{F0: (v39 + v39), F1: v4}
                                    _ = v123
                                }
                            } else {
                                switch (fi40 - (fi40 * 945)) {
                                case 187:
                                    if (42 > 225) {
                                        var v124 bool = v4
                                        var v125 float32 = ((-2.25 / (0.1 + 1e10)) * (-2.25 / (1e10 + (3.14159 + (3.14159 - 1.5)))))
                                        _ = v124
                                        _ = v125
                                    }
                                case 641:
                                    switch (fi40 ^ (fi34 + fi34)) {
                                    case -9223372036854775808:
                                        switch (v39 + v39) {
                                        case "":
                                            fmt.Println(v5, v39)
                                            var v126 float32 = ((float32(0.0) / 1e10) * 1e10)
                                            v38 = v38
                                            _ = v126
                                        case "hello":
                                            var v127 uint = (v1 * 4294967295)
                                            var v128 uint = 927
                                            v129 := GenS3{F0: 2147483647, F1: v5}
                                            _ = v127
                                            _ = v128
                                            _ = v129
                                        }
                                    }
                                case -64:
                                    v130 := GenS2{F0: 255, F1: -1}
                                    var v131 int16 = (v38 + (^(v38)))
                                    var v132 uint = (v1 & (v1 >> ((v1) & 63)))
                                    _ = v130
                                    _ = v131
                                    _ = v132
                                }
                            }
                        default:
                            v133 := GenS5{F0: 0}
                            var v134 uint16 = 65535
                            switch v39 {
                            case "hello":
                                if len(v30) > 0 {
                                    v30[0] = fi34
                                }
                            case "foo bar":
                                var v135 int16 = 795
                                _ = v135
                            case "\n\t":
                                var v136 uint = (^(v1))
                                _ = v136
                            default:
                                v137 := GenS13{F0: 768}
                                _ = v137
                            }
                            _ = v133
                            _ = v134
                        }
                        v138 := GenS8{F0: -478, F1: true}
                        _ = v138
                    } else {
                        if ("\n\t" > v39) {
                            fmt.Println(fi34)
                            var v139 float64 = -2.25
                            v140 := v3[fi34]
                            _ = v139
                            _ = v140
                        }
                    }
                    fmt.Println(v39, fi34)
                    v141 := GenS14{F0: v4}
                    _ = fi40
                    _ = v141
                }
                _ = v38
                _ = v39
            }
            _ = fi34
        }
    case "a":
        v1 -= v1
    default:
        v142 := GenS15{F0: true}
        _ = v142
    }
    if v4 {
        if (32767 != -357) {
            switch "\n\t" {
            case "hello":
                if false {
                    var v143 float64 = 0.1
                    var v144 int32 = v5
                    _ = v143
                    _ = v144
                }
            case "a":
                if v4 {
                    _ = len(v3)
                } else {
                    for fi145 := 0; fi145 < 18; fi145++ {
                        if len(v30) > 0 {
                            v30[0] = fi145
                        }
                        _ = fi145
                    }
                    switch "foo bar" {
                    case "\n\t":
                        for fi146 := 0; fi146 < 6; fi146++ {
                            for fi147 := 0; fi147 < 6; fi147++ {
                                if v4 {
                                    if (fi146 >= fi147) {
                                        fmt.Println(v4, fi146)
                                    }
                                    v148 := GenS10{F0: (fi147 == (^(fi147))), F1: 128, F2: (fi146 + (fi146 & (-(fi147))))}
                                    _ = v148
                                }
                                if false {
                                    v149 := v6[("foo bar" + "foo bar")]
                                    if len(v30) >= 2 {
                                        v30 = v30[1:len(v30)]
                                    }
                                    v1 += 42
                                    _ = v149
                                }
                                _ = fi147
                            }
                            _ = fi146
                        }
                    case "hello":
                        var v150 int = 128
                        _ = v150
                    }
                    fmt.Println(v5)
                }
            default:
                for fi151 := 0; fi151 < 9; fi151++ {
                    v5 -= v5
                    _ = fi151
                }
            }
        } else {
            var v152 int16 = 32767
            v153 := GenS16{F0: 1, F1: 820, F2: 1}
            var v154 string = "hello"
            _ = v152
            _ = v153
            _ = v154
        }
    } else {
        v5 = v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
