package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint64; F2 string }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS7 struct { F0 int32; F1 bool; F2 int32 }
type GenS8 struct { F0 bool; F1 uint; F2 int16 }
type GenS9 struct { F0 uint8; F1 int64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int32; F1 bool; F2 uint16 }
type GenS12 struct { F0 int64; F1 int16; F2 int16 }
type GenS13 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS14 struct { F0 string; F1 uint; F2 int32 }
type GenS15 struct { F0 string; F1 int8 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS18 struct { F0 uint8; F1 uint }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS21 struct { F0 int16; F1 uint8 }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 int64 }
type GenS24 struct { F0 uint8 }
type GenS25 struct { F0 int8; F1 uint }
type GenS26 struct { F0 uint32; F1 uint64 }
type GenS27 struct { F0 int8; F1 int64; F2 uint }
type GenS28 struct { F0 uint16; F1 string }
type GenS29 struct { F0 uint; F1 int16; F2 uint }
type GenS30 struct { F0 int16; F1 int16; F2 uint }
type GenS31 struct { F0 int64; F1 string; F2 bool }
type GenS32 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS33 struct { F0 uint8; F1 uint64 }
type GenS34 struct { F0 int16; F1 int8; F2 uint32 }
type GenS35 struct { F0 int }
type GenS36 struct { F0 int64; F1 bool; F2 string }
type GenS37 struct { F0 uint; F1 uint8 }
type GenS38 struct { F0 bool; F1 string; F2 int32 }

func main() {
    var v1 bool = false
    var v2 float32 = ((0.1 + -2.25) - (0.1 / float32((1e10 / 1.5))))
    var v3 int64 = 9223372036854775807
    var v4 int64 = v3
    var v5 uint16 = 634
    v6 := []int{341, -629}
    _ = len(v6)
    v6 = append(v6, -320)
    var v7 int64 = v4
    var v8 uint64 = 132
    var v9 uint = 604
    switch (("a" + "a") + "") {
    case "\n\t":
        var v10 uint32 = 2147483647
        _ = v10
    case "a":
        for fi11 := 0; fi11 < 2; fi11++ {
            for fi12 := 0; fi12 < 16; fi12++ {
                switch (^(fi11)) {
                case 2147483647:
                    if len(v6) > 0 {
                        v6[0] = (fi12 << ((uint(v8)) & 63))
                    }
                    v13 := GenS1{F0: 833}
                    _ = v13
                case -2147483648:
                    v1 = (v1 || (!(!(-2147483648 >= 2147483647))))
                }
                if v1 {
                    v14 := GenS1{F0: -2147483648}
                    var v15 bool = (v2 <= v2)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v9 = (^(v9))
                    if v15 {
                        v2 -= float32((0.0 / -2.25))
                    }
                    _ = v14
                    _ = v15
                } else {
                    fmt.Println(v9)
                }
                v16 := GenS1{F0: 220}
                v17 := GenS1{F0: -2147483648}
                var v18 uint8 = 2
                v7 += (v4 + (v3 ^ -1))
                fmt.Println(v2)
                v3 = (v7 >> ((v9) & 63))
                v3 = 0
                v19 := GenS1{F0: 1}
                _ = fi12
                _ = v16
                _ = v17
                _ = v18
                _ = v19
            }
            var v20 uint32 = 896
            _ = fi11
            _ = v20
        }
    default:
        var v21 uint8 = 2
        _ = v21
    }
    if v1 {
        for fi22 := 0; fi22 < 12; fi22++ {
            if ((0 != 187) && (!(!v1))) {
                if (454 == -444) {
                    if len(v6) > 0 {
                        v6[0] = -1
                    }
                    if true {
                        if (fi22 == (^(fi22))) {
                            switch fi22 {
                            case -2147483648:
                                for fi23 := 0; fi23 < 18; fi23++ {
                                    if ((!v1) && (v8 <= 0)) {
                                        switch "" {
                                        case "foo bar":
                                            if v1 {
                                                v24 := GenS1{F0: -2147483648}
                                                var v25 float64 = ((0.0 + (0.1 / (0.0 + (0.1 + (float64(v2) * (0.0 - (0.1 + 0.0))))))) / (0.0 + 0.1))
                                                if len(v6) > 0 {
                                                    v6[0] = fi22
                                                }
                                                _ = v24
                                                _ = v25
                                            } else {
                                                switch (fi22 - -1) {
                                                case 42:
                                                    v6 = append(v6, 801)
                                                    v5 = v5
                                                case 9223372036854775807:
                                                    v2 = v2
                                                default:
                                                    v7 = v4
                                                }
                                                var v26 int32 = 42
                                                v6 = append(v6, fi22)
                                                _ = v26
                                            }
                                        case "\n\t":
                                            _ = len(v6)
                                        case "a":
                                            var v27 int8 = 91
                                            _ = v27
                                        }
                                        v28 := GenS1{F0: 1}
                                        fmt.Println(v8, fi23, fi22)
                                        _ = v28
                                    }
                                    fmt.Println(v7, fi22, v9)
                                    _ = fi23
                                }
                            }
                            fmt.Println(v7, v5, v1)
                            v2 -= float32((3.14159 - ((0.1 / (3.14159 - ((1.5 * 0.0) - (0.1 * 1.5)))) * 1.5)))
                        } else {
                            fmt.Println(v8)
                            for fi29 := 0; fi29 < 7; fi29++ {
                                if ((0.1 != (v2 + (-2.25 / v2))) || (v7 != 1)) {
                                    if v1 {
                                        switch "\n\t" {
                                        case "foo bar":
                                            v7 = -9223372036854775808
                                        default:
                                            v9 = (v9 | (^(v9)))
                                        }
                                    } else {
                                        var v30 int64 = -9223372036854775808
                                        if (((v8 * 847) > v8) && true) {
                                            var v31 int64 = 0
                                            var v32 int64 = v31
                                            if len(v6) > 0 {
                                                v6[0] = (-(fi22))
                                            }
                                            _ = v31
                                            _ = v32
                                        }
                                        v30 -= v7
                                        _ = v30
                                    }
                                    v4 = 1
                                    var v33 uint64 = 2
                                    _ = v33
                                }
                                var v34 int16 = -47
                                _ = fi29
                                _ = v34
                            }
                            fmt.Println(fi22, v8)
                        }
                        if len(v6) > 0 {
                            v6[0] = fi22
                        }
                        var v35 int8 = 0
                        _ = v35
                    } else {
                        for fi36 := 0; fi36 < 10; fi36++ {
                            if len(v6) > 0 {
                                v6[0] = fi22
                            }
                            var v37 uint16 = v5
                            _ = fi36
                            _ = v37
                        }
                        v38 := GenS1{F0: 0}
                        _ = v38
                    }
                    v39 := GenS1{F0: 1}
                    _ = v39
                } else {
                    if v1 {
                        v2 -= v2
                    }
                    if (-596 != -447) {
                        if (1 == v7) {
                            v6 = append(v6, 127)
                            if (v1 || (2147483648 <= uint32(v8))) {
                                if len(v6) > 0 {
                                    v6[0] = (fi22 << ((v9) & 63))
                                }
                            }
                            v3 = v7
                        } else {
                            var v40 uint8 = 127
                            v6 = append(v6, 9223372036854775807)
                            for fi41 := 0; fi41 < 12; fi41++ {
                                v42 := GenS1{F0: -328}
                                fmt.Println(v1, v8, v9)
                                var v43 uint = v9
                                _ = fi41
                                _ = v42
                                _ = v43
                            }
                            _ = v40
                        }
                    } else {
                        if len(v6) > 0 {
                            v6[0] = (fi22 | (^(fi22)))
                        }
                    }
                }
                if len(v6) > 0 {
                    v6[0] = 32768
                }
            } else {
                fmt.Println(v1, v4)
                for fi44 := 0; fi44 < 4; fi44++ {
                    v45 := GenS1{F0: -1}
                    _ = fi44
                    _ = v45
                }
                v46 := GenS1{F0: 0}
                _ = v46
            }
            var v47 string = "a"
            fmt.Println(v4, v8)
            _ = fi22
            _ = v47
        }
        switch 948 {
        case 1:
            if false {
                v48 := GenS1{F0: 1}
                if (true || v1) {
                    if (1 != 0) {
                        switch "foo bar" {
                        case "\n\t":
                            var v49 int = int(v9)
                            _ = v49
                        case "foo bar":
                            fmt.Println(v9, v3)
                        case "hello":
                            var v50 int8 = -128
                            _ = v50
                        }
                        v1 = (v1 && true)
                        switch "" {
                        case "foo bar":
                            for fi51 := 0; fi51 < 19; fi51++ {
                                for fi52 := 0; fi52 < 7; fi52++ {
                                    if v1 {
                                        v53 := GenS1{F0: 198}
                                        fmt.Println(v3, v3, fi51)
                                        v6 = append(v6, (^(fi52)))
                                        _ = v53
                                    } else {
                                        _ = len(v6)
                                    }
                                    switch int(v5) {
                                    case -887:
                                        v7 = v3
                                    default:
                                        switch "hello" {
                                        case "\n\t":
                                            var v54 uint64 = 2
                                            _ = v54
                                        }
                                    }
                                    if (!true) {
                                        var v55 float32 = (1.5 / 1.5)
                                        _ = v55
                                    }
                                    _ = fi52
                                }
                                if (fi51 < fi51) {
                                    var v56 string = ""
                                    _ = v56
                                }
                                var v57 uint8 = 127
                                _ = fi51
                                _ = v57
                            }
                        case "":
                            fmt.Println(v7, v9)
                        case "a":
                            _ = len(v6)
                        default:
                            switch "hello" {
                            case "hello":
                                switch 32768 {
                                case 4294967295:
                                    var v58 uint = v9
                                    _ = v58
                                case 128:
                                    for fi59 := 0; fi59 < 4; fi59++ {
                                        if len(v6) > 0 {
                                            v6[0] = 835
                                        }
                                        for fi60 := 0; fi60 < 1; fi60++ {
                                            _ = len(v6)
                                            switch 849 {
                                            case -1:
                                                if (-1 > -595) {
                                                    if v1 {
                                                        v61 := GenS2{F0: 127, F1: (^(v8)), F2: ("\n\t" + ((("" + "\n\t") + "foo bar") + (("a" + "hello") + "\n\t")))}
                                                        fmt.Println(fi59, v9)
                                                        switch fi59 {
                                                        case -999:
                                                            var v62 uint = 2
                                                            _ = v62
                                                        case 1:
                                                            var v63 string = "foo bar"
                                                            _ = v63
                                                        }
                                                        if (v8 > 273) {
                                                            if ((v4 < v7) || v1) {
                                                                v64 := GenS3{F0: 251, F1: 95}
                                                                v8 = v8
                                                                _ = v64
                                                            } else {
                                                                v6 = append(v6, 2147483647)
                                                                v7 -= (v4 | v4)
                                                                if true {
                                                                    var v65 float32 = v2
                                                                    _ = v65
                                                                } else {
                                                                    v66 := GenS4{F0: v3}
                                                                    switch (fi59 ^ int(v7)) {
                                                                    case 4294967295:
                                                                        fmt.Println(v1, fi60, v2)
                                                                    default:
                                                                        var v67 float32 = v2
                                                                        _ = v67
                                                                    }
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                    var v68 int32 = 42
                                                                    _ = v66
                                                                    _ = v68
                                                                }
                                                            }
                                                        } else {
                                                            fmt.Println(v3, v5)
                                                            var v69 uint16 = 1
                                                            var v70 uint8 = 128
                                                            _ = v69
                                                            _ = v70
                                                        }
                                                        v3 += (v7 % ((42) | 1))
                                                        _ = v61
                                                    }
                                                    v71 := GenS5{F0: true}
                                                    _ = v71
                                                } else {
                                                    v8 += v8
                                                    if v1 {
                                                        var v72 int32 = 1
                                                        fmt.Println(fi60, v72, v8)
                                                        v6 = append(v6, fi59)
                                                        switch ("a" + "") {
                                                        case "hello":
                                                            if false {
                                                                if (("hello" + ("" + "\n\t")) > "hello") {
                                                                    if (v1 || (!(1 >= 38))) {
                                                                        var v73 bool = false
                                                                        _ = v73
                                                                    }
                                                                } else {
                                                                    v74 := GenS1{F0: 913}
                                                                    v6 = append(v6, (fi59 / (((fi59 + (-(fi59)))) | 1)))
                                                                    _ = v74
                                                                }
                                                                switch (fi59 << ((v9) & 63)) {
                                                                case 128:
                                                                    v75 := GenS3{F0: 127, F1: 151}
                                                                    v1 = false
                                                                    v76 := GenS6{F0: 1, F1: 2, F2: v8}
                                                                    var v77 int8 = 98
                                                                    _ = v75
                                                                    _ = v76
                                                                    _ = v77
                                                                case -573:
                                                                    v78 := GenS7{F0: 1, F1: true, F2: (v72 - v72)}
                                                                    _ = v78
                                                                case 0:
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        if (2147483648 <= 0) {
                                                                            v6 = append(v6, 147)
                                                                            fmt.Println(v9, v72, v1)
                                                                            if false {
                                                                                if (!(false && false)) {
                                                                                    var v79 uint64 = v8
                                                                                    if len(v6) >= 2 {
                                                                                        v6 = v6[1:len(v6)]
                                                                                    }
                                                                                    v80 := GenS5{F0: v1}
                                                                                    _ = v79
                                                                                    _ = v80
                                                                                }
                                                                                if v1 {
                                                                                    var v81 int8 = 42
                                                                                    if (false && (!(v5 <= (^(v5))))) {
                                                                                        if ((true && (float32(0.1) > ((v2 - v2) / 1.0))) && (!false)) {
                                                                                            if (((3.14159 + 3.14159) - -2.25) != (v2 / v2)) {
                                                                                                v82 := GenS7{F0: 42, F1: false, F2: v72}
                                                                                                _ = v82
                                                                                            }
                                                                                            fmt.Println(v4, fi59)
                                                                                        }
                                                                                        v83 := GenS8{F0: v1, F1: 972, F2: -719}
                                                                                        _ = v83
                                                                                    } else {
                                                                                        var v84 float32 = v2
                                                                                        var v85 uint = (v9 % ((1) | 1))
                                                                                        v85 = 42
                                                                                        var v86 uint8 = 16
                                                                                        switch (fi59 - (fi59 >> (((^(v85))) & 63))) {
                                                                                        case 394:
                                                                                            switch 303 {
                                                                                            case -2147483648:
                                                                                                v2 += (v2 + 3.14159)
                                                                                            case -9223372036854775808:
                                                                                                var v87 bool = (true && (!("a" != ("a" + "foo bar"))))
                                                                                                v88 := GenS9{F0: (v86 * uint8(v8)), F1: (v7 + v3)}
                                                                                                _ = v87
                                                                                                _ = v88
                                                                                            }
                                                                                        case -925:
                                                                                            fmt.Println(v2, fi60, v86)
                                                                                        case 32767:
                                                                                            v81 -= (v81 ^ (v81 - v81))
                                                                                        default:
                                                                                            v6 = append(v6, (fi60 >> ((v9) & 63)))
                                                                                        }
                                                                                        _ = v84
                                                                                        _ = v85
                                                                                        _ = v86
                                                                                    }
                                                                                    var v89 int = 742
                                                                                    _ = v81
                                                                                    _ = v89
                                                                                }
                                                                                v3 += 9223372036854775807
                                                                            } else {
                                                                                if (255 >= v9) {
                                                                                    if v1 {
                                                                                        switch int(v72) {
                                                                                        case 1:
                                                                                            _ = len(v6)
                                                                                        case 127:
                                                                                            if (255 < 255) {
                                                                                                v6 = append(v6, fi60)
                                                                                                if (2147483648 >= 2) {
                                                                                                    fmt.Println(v1)
                                                                                                } else {
                                                                                                    var v90 uint64 = 18446744073709551615
                                                                                                    v91 := GenS10{F0: -599}
                                                                                                    v4 -= 42
                                                                                                    _ = v90
                                                                                                    _ = v91
                                                                                                }
                                                                                                if len(v6) > 0 {
                                                                                                    v6[0] = fi59
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v92 := GenS8{F0: v1, F1: v9, F2: -718}
                                                                                        if (fi59 < (fi60 / ((fi60) | 1))) {
                                                                                            var v93 int16 = int16(fi59)
                                                                                            if v1 {
                                                                                                var v94 uint16 = (^(v5))
                                                                                                var v95 float32 = ((v2 / v2) / (1.5 * 3.14159))
                                                                                                _ = v94
                                                                                                _ = v95
                                                                                            }
                                                                                            _ = v93
                                                                                        }
                                                                                        _ = v92
                                                                                    } else {
                                                                                        var v96 int64 = 457
                                                                                        _ = v96
                                                                                    }
                                                                                    if false {
                                                                                        if (v4 <= (-(v7))) {
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                _ = len(v6)
                                                                                            }
                                                                                            v97 := GenS9{F0: 255, F1: v7}
                                                                                            var v98 int16 = 1
                                                                                            var v99 bool = false
                                                                                            _ = v97
                                                                                            _ = v98
                                                                                            _ = v99
                                                                                        } else {
                                                                                            fmt.Println(v7)
                                                                                            var v100 int64 = v7
                                                                                            _ = v100
                                                                                        }
                                                                                        var v101 int64 = -9223372036854775808
                                                                                        switch "\n\t" {
                                                                                        case "\n\t":
                                                                                            if (!v1) {
                                                                                                if len(v6) >= 2 {
                                                                                                    v6 = v6[1:len(v6)]
                                                                                                }
                                                                                            } else {
                                                                                                if true {
                                                                                                    v101 = (v3 * (v7 >> (((v9 - (^(v9)))) & 63)))
                                                                                                    v8 = v8
                                                                                                    fmt.Println(v4, fi59, fi59)
                                                                                                }
                                                                                                if ((v5 * 2) != v5) {
                                                                                                    v1 = v1
                                                                                                    if v1 {
                                                                                                        var v102 int16 = 1000
                                                                                                        v72 += v72
                                                                                                        fmt.Println(v101)
                                                                                                        _ = v102
                                                                                                    }
                                                                                                    var v103 int32 = 42
                                                                                                    _ = v103
                                                                                                } else {
                                                                                                    var v104 float32 = (v2 / (((((((v2 * v2) * ((v2 + (v2 * v2)) - 0.0)) + (float32(-2.25) + v2)) - v2) / (((v2 + v2) / (v2 - v2)) + v2)) / v2) * v2))
                                                                                                    switch (fi59 / ((545) | 1)) {
                                                                                                    case -2147483648:
                                                                                                        v3 += (v101 | v4)
                                                                                                    case -721:
                                                                                                        var v105 int16 = 1000
                                                                                                        _ = v105
                                                                                                    case 4294967295:
                                                                                                        if len(v6) >= 2 {
                                                                                                            v6 = v6[1:len(v6)]
                                                                                                        }
                                                                                                    default:
                                                                                                        var v106 int32 = v72
                                                                                                        _ = v106
                                                                                                    }
                                                                                                    var v107 uint8 = 42
                                                                                                    v104 -= 0.0
                                                                                                    if v1 {
                                                                                                        v108 := GenS11{F0: (-(v72)), F1: (!false), F2: v5}
                                                                                                        if len(v6) >= 2 {
                                                                                                            v6 = v6[1:len(v6)]
                                                                                                        }
                                                                                                        v101 -= 42
                                                                                                        switch (fi59 + (-(fi60))) {
                                                                                                        case 4294967295:
                                                                                                            if len(v6) >= 2 {
                                                                                                                v6 = v6[1:len(v6)]
                                                                                                            }
                                                                                                        case -1:
                                                                                                            var v109 uint16 = (^(v5))
                                                                                                            _ = v109
                                                                                                        case -9223372036854775808:
                                                                                                            v101 += (v7 | v101)
                                                                                                        }
                                                                                                        v110 := GenS3{F0: (v107 | v107), F1: 66}
                                                                                                        v4 -= -920
                                                                                                        _ = v108
                                                                                                        _ = v110
                                                                                                    }
                                                                                                    if (-32768 < 1000) {
                                                                                                        fmt.Println(v2, v3, v8)
                                                                                                        v4 = v7
                                                                                                        v101 -= v7
                                                                                                        if v1 {
                                                                                                            v111 := GenS4{F0: 1}
                                                                                                            v8 = v8
                                                                                                            switch 2147483647 {
                                                                                                            case 0:
                                                                                                                v7 += (v3 / ((v3) | 1))
                                                                                                            case -560:
                                                                                                                if (v1 || ((^(v5)) != (v5 | (v5 << (((v9 | (^(v9)))) & 63))))) {
                                                                                                                    v101 += (v101 << (((v9 | v9)) & 63))
                                                                                                                    if v1 {
                                                                                                                        v1 = v1
                                                                                                                        v104 += v2
                                                                                                                        var v112 uint32 = 4294967295
                                                                                                                        v8 = 280
                                                                                                                        _ = v112
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    if len(v6) > 0 {
                                                                                                                        v6[0] = 0
                                                                                                                    }
                                                                                                                    if len(v6) > 0 {
                                                                                                                        v6[0] = (-(fi59))
                                                                                                                    }
                                                                                                                    var v113 uint16 = v5
                                                                                                                    _ = v113
                                                                                                                }
                                                                                                            case 4294967295:
                                                                                                                fmt.Println(v9, v8)
                                                                                                            }
                                                                                                            _ = v111
                                                                                                        } else {
                                                                                                            v114 := GenS3{F0: 0, F1: v107}
                                                                                                            if (false && (216 == 4294967295)) {
                                                                                                                v1 = true
                                                                                                            } else {
                                                                                                                v6 = append(v6, int(v5))
                                                                                                                fmt.Println(v4, v7)
                                                                                                            }
                                                                                                            var v115 float64 = float64((-2.25 / (v104 * (v2 * v104))))
                                                                                                            if len(v6) >= 2 {
                                                                                                                v6 = v6[1:len(v6)]
                                                                                                            }
                                                                                                            fmt.Println(v8)
                                                                                                            v6 = append(v6, fi59)
                                                                                                            _ = v114
                                                                                                            _ = v115
                                                                                                        }
                                                                                                        v116 := GenS4{F0: (v101 / (((v4 * 177)) | 1))}
                                                                                                        _ = v116
                                                                                                    }
                                                                                                    v101 = 516
                                                                                                    _ = v104
                                                                                                    _ = v107
                                                                                                }
                                                                                            }
                                                                                        case "hello":
                                                                                            if v1 {
                                                                                                fmt.Println(v9, v4)
                                                                                                if false {
                                                                                                    _ = len(v6)
                                                                                                } else {
                                                                                                    if len(v6) >= 2 {
                                                                                                        v6 = v6[1:len(v6)]
                                                                                                    }
                                                                                                    fmt.Println(v3, fi59)
                                                                                                }
                                                                                                fmt.Println(v5, v101, v3)
                                                                                            } else {
                                                                                                v1 = v1
                                                                                                if false {
                                                                                                    var v117 float64 = 1e10
                                                                                                    v117 = (0.0 + float64(0.1))
                                                                                                    if ((!(!v1)) && (!(0 > 127))) {
                                                                                                        v6 = append(v6, 720)
                                                                                                        v2 = 1e10
                                                                                                    } else {
                                                                                                        v1 = true
                                                                                                    }
                                                                                                    if len(v6) > 0 {
                                                                                                        v6[0] = fi60
                                                                                                    }
                                                                                                    _ = v117
                                                                                                }
                                                                                                v3 = v101
                                                                                                fmt.Println(v3)
                                                                                                if (v1 && ((42 < uint(v5)) || (42 <= 255))) {
                                                                                                    if false {
                                                                                                        v5 += (v5 * (v5 ^ (^(v5))))
                                                                                                        var v118 int8 = 100
                                                                                                        _ = v118
                                                                                                    } else {
                                                                                                        var v119 uint = v9
                                                                                                        _ = v119
                                                                                                    }
                                                                                                    v8 = v8
                                                                                                    if true {
                                                                                                        var v120 string = "hello"
                                                                                                        if (!(!(932 == 2))) {
                                                                                                            fmt.Println(v1, v5)
                                                                                                            v121 := GenS8{F0: (128 < 191), F1: (v9 - 4294967295), F2: 293}
                                                                                                            _ = v121
                                                                                                        } else {
                                                                                                            var v122 uint = (v9 >> ((2) & 63))
                                                                                                            if true {
                                                                                                                if (2147483647 < 42) {
                                                                                                                    var v123 uint16 = v5
                                                                                                                    var v124 int32 = 2147483647
                                                                                                                    _ = v123
                                                                                                                    _ = v124
                                                                                                                } else {
                                                                                                                    if (!v1) {
                                                                                                                        v125 := GenS12{F0: 42, F1: 32767, F2: -1}
                                                                                                                        _ = v125
                                                                                                                    } else {
                                                                                                                        if (0 == (v122 | (v9 / (((v122 * v122)) | 1)))) {
                                                                                                                            v126 := GenS5{F0: v1}
                                                                                                                            var v127 int8 = 1
                                                                                                                            if true {
                                                                                                                                var v128 int16 = 42
                                                                                                                                if (v7 == (^(v7))) {
                                                                                                                                    v122 -= (^(v122))
                                                                                                                                } else {
                                                                                                                                    v129 := GenS13{F0: (v128 >> ((uint(v5)) & 63)), F1: (v8 ^ (^(v8))), F2: 127}
                                                                                                                                    _ = v129
                                                                                                                                }
                                                                                                                                v130 := GenS12{F0: (v3 & v101), F1: (-(v128)), F2: (-(v128))}
                                                                                                                                v131 := GenS14{F0: ("foo bar" + "foo bar"), F1: v122, F2: -2147483648}
                                                                                                                                _ = v128
                                                                                                                                _ = v130
                                                                                                                                _ = v131
                                                                                                                            }
                                                                                                                            _ = v126
                                                                                                                            _ = v127
                                                                                                                        } else {
                                                                                                                            v132 := GenS15{F0: v120, F1: 108}
                                                                                                                            var v133 uint64 = 167
                                                                                                                            v4 += v101
                                                                                                                            v134 := GenS9{F0: 71, F1: 624}
                                                                                                                            v133 = v8
                                                                                                                            v101 -= (v7 / ((9223372036854775807) | 1))
                                                                                                                            _ = v132
                                                                                                                            _ = v133
                                                                                                                            _ = v134
                                                                                                                        }
                                                                                                                        var v135 string = "\n\t"
                                                                                                                        _ = v135
                                                                                                                    }
                                                                                                                }
                                                                                                                if (!v1) {
                                                                                                                    v136 := GenS6{F0: 42, F1: 1, F2: v8}
                                                                                                                    _ = v136
                                                                                                                }
                                                                                                                _ = len(v6)
                                                                                                            }
                                                                                                            var v137 int16 = -1
                                                                                                            _ = v122
                                                                                                            _ = v137
                                                                                                        }
                                                                                                        _ = v120
                                                                                                    } else {
                                                                                                        v101 = 1
                                                                                                    }
                                                                                                    if len(v6) >= 2 {
                                                                                                        v6 = v6[1:len(v6)]
                                                                                                    }
                                                                                                    var v138 int32 = v72
                                                                                                    var v139 bool = v1
                                                                                                    _ = v138
                                                                                                    _ = v139
                                                                                                } else {
                                                                                                    if len(v6) >= 2 {
                                                                                                        v6 = v6[1:len(v6)]
                                                                                                    }
                                                                                                    var v140 float32 = -2.25
                                                                                                    _ = v140
                                                                                                }
                                                                                            }
                                                                                        case "":
                                                                                            var v141 bool = true
                                                                                            _ = v141
                                                                                        default:
                                                                                            if true {
                                                                                                var v142 uint = (v9 / ((v9) | 1))
                                                                                                if len(v6) >= 2 {
                                                                                                    v6 = v6[1:len(v6)]
                                                                                                }
                                                                                                _ = v142
                                                                                            }
                                                                                        }
                                                                                        _ = v101
                                                                                    }
                                                                                }
                                                                            }
                                                                        } else {
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                            v143 := GenS16{F0: (fi59 << ((v9) & 63))}
                                                                            var v144 int = -2147483648
                                                                            fmt.Println(v2, v7, v144)
                                                                            _ = v143
                                                                            _ = v144
                                                                        }
                                                                    case "a":
                                                                        v6 = append(v6, (fi59 % ((992) | 1)))
                                                                        switch (-(fi60)) {
                                                                        case 128:
                                                                            v3 += (v3 * (v3 ^ (-(v3))))
                                                                        case 0:
                                                                            v4 += (v7 % ((-1) | 1))
                                                                        default:
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                if false {
                                                                                    v9 = v9
                                                                                    if true {
                                                                                        switch (("a" + "a") + ("" + "hello")) {
                                                                                        case "foo bar":
                                                                                            var v145 int8 = -1
                                                                                            _ = v145
                                                                                        case "":
                                                                                            v1 = true
                                                                                        case "a":
                                                                                            _ = len(v6)
                                                                                        }
                                                                                        if false {
                                                                                            v8 = v8
                                                                                            _ = len(v6)
                                                                                        } else {
                                                                                            if (0 > int16(v3)) {
                                                                                                v4 = int64(v9)
                                                                                                v7 = 966
                                                                                            }
                                                                                            v146 := GenS2{F0: 128, F1: v8, F2: "foo bar"}
                                                                                            _ = v146
                                                                                        }
                                                                                    } else {
                                                                                        if true {
                                                                                            var v147 uint8 = 255
                                                                                            if len(v6) >= 2 {
                                                                                                v6 = v6[1:len(v6)]
                                                                                            }
                                                                                            _ = v147
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(fi60, v1)
                                                                                    if ((!v1) || ((^(fi60)) <= fi60)) {
                                                                                        switch fi59 {
                                                                                        case 388:
                                                                                            if len(v6) > 0 {
                                                                                                v6[0] = fi59
                                                                                            }
                                                                                        default:
                                                                                            var v148 int32 = v72
                                                                                            _ = v148
                                                                                        }
                                                                                    } else {
                                                                                        v2 -= 0.1
                                                                                        v149 := GenS14{F0: "\n\t", F1: (v9 ^ v9), F2: (^(v72))}
                                                                                        v150 := GenS16{F0: (fi59 ^ (fi59 ^ (fi60 * -382)))}
                                                                                        _ = v149
                                                                                        _ = v150
                                                                                    }
                                                                                    switch (("hello" + "a") + ("\n\t" + "a")) {
                                                                                    case "":
                                                                                        _ = len(v6)
                                                                                        fmt.Println(v9, v4, v1)
                                                                                    case "hello":
                                                                                        v151 := GenS2{F0: 1, F1: (v8 - 206), F2: ("" + "a")}
                                                                                        _ = v151
                                                                                    default:
                                                                                        v152 := GenS2{F0: uint8(v9), F1: (v8 ^ 18446744073709551615), F2: ("" + "a")}
                                                                                        _ = v152
                                                                                    }
                                                                                }
                                                                            default:
                                                                                switch (fi60 * 32767) {
                                                                                case 9223372036854775807:
                                                                                    v1 = (v8 != (^(v8)))
                                                                                }
                                                                            }
                                                                        }
                                                                        v153 := GenS12{F0: -9223372036854775808, F1: 536, F2: -32768}
                                                                        var v154 float64 = float64(v2)
                                                                        _ = v153
                                                                        _ = v154
                                                                    }
                                                                default:
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                    v6 = append(v6, -9223372036854775808)
                                                                    if v1 {
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                        v155 := GenS9{F0: 0, F1: (v3 >> ((18446744073709551615) & 63))}
                                                                        v7 += (v4 + v7)
                                                                        _ = v155
                                                                    } else {
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                        v156 := GenS9{F0: 127, F1: v4}
                                                                        var v157 float64 = 1.5
                                                                        _ = v156
                                                                        _ = v157
                                                                    }
                                                                    v4 -= (-(v3))
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        fmt.Println(v9, fi59)
                                                                    case "hello":
                                                                        if v1 {
                                                                            v2 += v2
                                                                            v7 -= (v3 / (((v7 / ((v7) | 1))) | 1))
                                                                            if len(v6) > 0 {
                                                                                v6[0] = 467
                                                                            }
                                                                        } else {
                                                                            v158 := GenS14{F0: "", F1: (v9 / (((v9 ^ (v9 ^ (v9 >> (((v9 % ((285) | 1))) & 63))))) | 1)), F2: int32(fi60)}
                                                                            _ = v158
                                                                        }
                                                                    }
                                                                }
                                                                v159 := GenS14{F0: "a", F1: v9, F2: (v72 << ((v9) & 63))}
                                                                _ = v159
                                                            } else {
                                                                var v160 uint16 = v5
                                                                _ = v160
                                                            }
                                                        case "a":
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                        case "":
                                                            switch fi59 {
                                                            case -2147483648:
                                                                v161 := GenS1{F0: (-(v72))}
                                                                _ = v161
                                                            case 42:
                                                                v1 = true
                                                            case 128:
                                                                v72 = v72
                                                            default:
                                                                fmt.Println(v1)
                                                            }
                                                        default:
                                                            var v162 int = (fi60 >> ((v9) & 63))
                                                            _ = v162
                                                        }
                                                        _ = v72
                                                    }
                                                    fmt.Println(v5, v3)
                                                }
                                            default:
                                                switch "a" {
                                                case "a":
                                                    switch fi60 {
                                                    case 306:
                                                        v4 = (-(v3))
                                                    case 2147483647:
                                                        v163 := GenS9{F0: 127, F1: v4}
                                                        _ = v163
                                                    default:
                                                        var v164 int16 = -874
                                                        _ = v164
                                                    }
                                                case "foo bar":
                                                    v1 = ("\n\t" <= "")
                                                }
                                            }
                                            fmt.Println(v7)
                                            _ = fi60
                                        }
                                        _ = fi59
                                    }
                                }
                            default:
                                var v165 uint = (^(v9))
                                _ = v165
                            }
                        }
                    } else {
                        var v166 int32 = 0
                        var v167 int = -9223372036854775808
                        _ = v166
                        _ = v167
                    }
                } else {
                    fmt.Println(v2, v1, v9)
                    if len(v6) > 0 {
                        v6[0] = -1
                    }
                    var v168 uint64 = (^(v8))
                    _ = v168
                }
                fmt.Println(v9, v1, v9)
                _ = v48
            } else {
                for fi169 := 0; fi169 < 7; fi169++ {
                    fmt.Println(v1)
                    for fi170 := 0; fi170 < 3; fi170++ {
                        _ = len(v6)
                        v4 -= (v3 & v4)
                        var v171 uint64 = 2
                        _ = fi170
                        _ = v171
                    }
                    if (((("" + "foo bar") <= ("hello" + "hello")) || (v8 != v8)) && v1) {
                        var v172 int32 = int32(v9)
                        _ = v172
                    } else {
                        v173 := GenS15{F0: ("foo bar" + ("" + (("foo bar" + ("hello" + ("a" + ("a" + "")))) + ((("hello" + "a") + (("\n\t" + "foo bar") + ("" + (("" + (("hello" + ((("a" + "a") + (("" + ("hello" + "hello")) + "a")) + "foo bar")) + "\n\t")) + "")))) + "\n\t")))), F1: 0}
                        fmt.Println(v7, v3)
                        for fi174 := 0; fi174 < 17; fi174++ {
                            if (!true) {
                                v5 = v5
                            } else {
                                if v1 {
                                    v9 -= v9
                                    v3 = 1
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                } else {
                                    if (v9 < (v9 ^ 2)) {
                                        v175 := GenS17{F0: 2, F1: uint32(v8), F2: v7}
                                        if true {
                                            if (0.0 == (v2 / (v2 - float32((1e10 - (1.5 / 0.1)))))) {
                                                var v176 string = "\n\t"
                                                _ = v176
                                            }
                                            v177 := GenS11{F0: -1, F1: (v9 <= (v9 ^ (v9 ^ (^(v9))))), F2: (^(v5))}
                                            _ = v177
                                        }
                                        _ = v175
                                    } else {
                                        _ = len(v6)
                                    }
                                }
                                v3 -= v7
                            }
                            switch fi169 {
                            case -9223372036854775808:
                                if v1 {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    var v178 uint64 = (v8 >> ((v9) & 63))
                                    _ = v178
                                } else {
                                    if true {
                                        v179 := GenS8{F0: false, F1: v9, F2: -32768}
                                        v2 += (3.14159 + ((0.0 / 1.5) / 1.0))
                                        _ = v179
                                    }
                                    var v180 int = fi174
                                    _ = v180
                                }
                                v1 = v1
                                switch -748 {
                                case 5:
                                    if (!((v5 / ((32767) | 1)) < (v5 >> (((v9 | 1)) & 63)))) {
                                        switch "" {
                                        case "\n\t":
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                    } else {
                                        if v1 {
                                            fmt.Println(v2, v4)
                                            v3 += v3
                                        }
                                        v181 := GenS18{F0: 55, F1: 2}
                                        _ = v181
                                    }
                                default:
                                    var v182 int = (fi169 - (fi169 - (fi174 >> (((v9 / (((^(v9))) | 1))) & 63))))
                                    _ = v182
                                }
                            case -2147483648:
                                switch fi169 {
                                case 374:
                                    v6 = append(v6, (fi174 + 9223372036854775807))
                                case 1:
                                    v183 := GenS10{F0: -678}
                                    _ = v183
                                default:
                                    var v184 string = ""
                                    _ = v184
                                }
                                switch fi169 {
                                case -1:
                                    var v185 uint64 = v8
                                    v6 = append(v6, -2147483648)
                                    v6 = append(v6, -634)
                                    _ = v185
                                case 4294967295:
                                    var v186 int64 = v3
                                    _ = v186
                                default:
                                    var v187 int16 = -967
                                    _ = v187
                                }
                                v1 = (!v1)
                            case 308:
                                fmt.Println(v1)
                            default:
                                fmt.Println(v9, fi174, fi169)
                            }
                            _ = fi174
                        }
                        _ = v173
                    }
                    _ = fi169
                }
                switch 32767 {
                case 0:
                    if v1 {
                        v188 := GenS19{F0: 42}
                        _ = v188
                    } else {
                        for fi189 := 0; fi189 < 13; fi189++ {
                            v7 = (v7 % ((v7) | 1))
                            for fi190 := 0; fi190 < 9; fi190++ {
                                v191 := GenS1{F0: 0}
                                if len(v6) > 0 {
                                    v6[0] = -71
                                }
                                _ = fi190
                                _ = v191
                            }
                            _ = fi189
                        }
                    }
                default:
                    switch 997 {
                    case 9223372036854775807:
                        var v192 float32 = 0.1
                        _ = v192
                    case 4294967295:
                        v193 := GenS12{F0: v3, F1: 32767, F2: 0}
                        _ = v193
                    }
                }
                for fi194 := 0; fi194 < 10; fi194++ {
                    if (v1 || v1) {
                        if true {
                            if len(v6) > 0 {
                                v6[0] = 0
                            }
                            for fi195 := 0; fi195 < 16; fi195++ {
                                if (float32((0.1 + -2.25)) == float32(-2.25)) {
                                    if len(v6) > 0 {
                                        v6[0] = fi194
                                    }
                                    var v196 uint32 = 282
                                    var v197 float64 = 0.0
                                    _ = v196
                                    _ = v197
                                } else {
                                    switch "\n\t" {
                                    case "a":
                                        fmt.Println(v5)
                                    case "\n\t":
                                        _ = len(v6)
                                        fmt.Println(v9)
                                        v1 = v1
                                    case "foo bar":
                                        v5 += v5
                                    default:
                                        if (!v1) {
                                            fmt.Println(v7, v3, v7)
                                        } else {
                                            v198 := GenS20{F0: 0, F1: 42, F2: v8}
                                            var v199 int32 = -2147483648
                                            if v1 {
                                                v200 := GenS21{F0: -20, F1: 255}
                                                _ = v200
                                            } else {
                                                fmt.Println(fi195, v7, v7)
                                            }
                                            v7 = 42
                                            var v201 bool = ((false || false) && ((1e10 * (1e10 - float32(-2.25))) <= v2))
                                            _ = v198
                                            _ = v199
                                            _ = v201
                                        }
                                    }
                                }
                                var v202 float32 = 0.0
                                var v203 uint8 = 46
                                _ = fi195
                                _ = v202
                                _ = v203
                            }
                        }
                        var v204 int64 = v3
                        v1 = (fi194 == fi194)
                        _ = v204
                    } else {
                        if v1 {
                            for fi205 := 0; fi205 < 19; fi205++ {
                                v206 := GenS4{F0: v7}
                                _ = fi205
                                _ = v206
                            }
                            var v207 int = (fi194 * fi194)
                            if len(v6) > 0 {
                                v6[0] = fi194
                            }
                            _ = v207
                        }
                        switch "" {
                        case "\n\t":
                            switch "a" {
                            case "\n\t":
                                v208 := GenS22{F0: v1}
                                _ = v208
                            }
                        case "":
                            if len(v6) > 0 {
                                v6[0] = (fi194 % ((fi194) | 1))
                            }
                        case "a":
                            for fi209 := 0; fi209 < 16; fi209++ {
                                fmt.Println(v2)
                                _ = fi209
                            }
                        }
                    }
                    if len(v6) > 0 {
                        v6[0] = fi194
                    }
                    v2 = v2
                    _ = fi194
                }
            }
        case 885:
            v1 = true
        }
    } else {
        for fi210 := 0; fi210 < 9; fi210++ {
            fmt.Println(v8)
            switch "\n\t" {
            case "foo bar":
                var v211 uint16 = 1
                _ = v211
            case "\n\t":
                v6 = append(v6, fi210)
            case "":
                v212 := GenS7{F0: -819, F1: true, F2: 1}
                _ = v212
            }
            _ = fi210
        }
        if (42 >= 28) {
            switch "a" {
            case "\n\t":
                if (((v7 & -1) >= v7) && v1) {
                    fmt.Println(v9, v4, v2)
                } else {
                    v213 := GenS20{F0: 2147483648, F1: (v8 * v8), F2: (^(v8))}
                    v3 = (v4 & v7)
                    var v214 float32 = (-2.25 * v2)
                    _ = v213
                    _ = v214
                }
            case "foo bar":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case "":
                if false {
                    v215 := GenS23{F0: v3}
                    _ = v215
                } else {
                    fmt.Println(v1, v1, v2)
                    fmt.Println(v2, v3, v4)
                    var v216 uint64 = v8
                    _ = v216
                }
            }
            v9 = v9
        } else {
            v217 := GenS8{F0: (!v1), F1: (v9 % ((v9) | 1)), F2: 596}
            var v218 uint16 = v5
            var v219 int16 = 1000
            _ = v217
            _ = v218
            _ = v219
        }
        for fi220 := 0; fi220 < 11; fi220++ {
            for fi221 := 0; fi221 < 20; fi221++ {
                var v222 uint32 = uint32(v5)
                fmt.Println(v7, v5, fi221)
                if ((!(0.0 > float32(((3.14159 * 0.1) - 1e10)))) && false) {
                    var v223 string = ""
                    if (!v1) {
                        switch "" {
                        case "hello":
                            v8 += (v8 ^ (v8 / (((v8 >> (((v9 | v9)) & 63))) | 1)))
                        }
                    }
                    v223 = v223
                    v2 = v2
                    var v224 uint64 = (v8 | 18446744073709551615)
                    var v225 float32 = 3.14159
                    if v1 {
                        if v1 {
                            v7 += -1
                        }
                    } else {
                        v9 -= v9
                        var v226 int32 = 1
                        _ = v226
                    }
                    _ = v223
                    _ = v224
                    _ = v225
                }
                _ = fi221
                _ = v222
            }
            v227 := GenS24{F0: 128}
            for fi228 := 0; fi228 < 1; fi228++ {
                switch fi228 {
                case 388:
                    if (v1 && v1) {
                        switch fi228 {
                        case 2147483647:
                            v5 = 65535
                        case 128:
                            switch fi220 {
                            case -993:
                                if len(v6) > 0 {
                                    v6[0] = (fi220 << ((1) & 63))
                                }
                            case 9223372036854775807:
                                v1 = v1
                            case -785:
                                var v229 int = fi220
                                _ = v229
                            }
                        case 42:
                            switch fi220 {
                            case 127:
                                if false {
                                    var v230 int8 = -56
                                    var v231 int8 = (^(v230))
                                    v6 = append(v6, fi220)
                                    if (391 != 1) {
                                        v231 = (v231 - v230)
                                        switch fi228 {
                                        case 2147483647:
                                            if false {
                                                switch (fi228 - (fi228 + fi228)) {
                                                case 32768:
                                                    var v232 int64 = (v3 % ((v7) | 1))
                                                    _ = v232
                                                case 0:
                                                    switch (fi228 << (((^(v9))) & 63)) {
                                                    case 42:
                                                        v7 = (v4 % ((v3) | 1))
                                                    }
                                                case 4294967295:
                                                    _ = len(v6)
                                                default:
                                                    if (2 > (^(v9))) {
                                                        var v233 uint32 = 2
                                                        if false {
                                                            switch "a" {
                                                            case "a":
                                                                v234 := GenS25{F0: 0, F1: (v9 + v9)}
                                                                var v235 uint64 = 1
                                                                var v236 float32 = ((v2 - (v2 * ((v2 - v2) - float32((0.1 / (1.5 + 3.14159)))))) + 1.5)
                                                                _ = v234
                                                                _ = v235
                                                                _ = v236
                                                            default:
                                                                var v237 uint64 = (v8 ^ (^(v8)))
                                                                _ = len(v6)
                                                                _ = v237
                                                            }
                                                            if true {
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                v238 := GenS26{F0: v233, F1: (v8 >> ((v9) & 63))}
                                                                if (true || v1) {
                                                                    var v239 float32 = 1e10
                                                                    v240 := GenS27{F0: v230, F1: (v3 * v7), F2: (v9 << (((v9 >> ((v9) & 63))) & 63))}
                                                                    v4 = (-(v3))
                                                                    _ = v239
                                                                    _ = v240
                                                                }
                                                                _ = v238
                                                            } else {
                                                                v241 := GenS19{F0: 987}
                                                                v1 = v1
                                                                switch ("" + ((("" + "\n\t") + (("a" + "\n\t") + "\n\t")) + "a")) {
                                                                case "foo bar":
                                                                    fmt.Println(v231, fi220)
                                                                    var v242 bool = (!v1)
                                                                    switch 2147483647 {
                                                                    case -1:
                                                                        var v243 uint = (v9 / (((v9 | (v9 * (^(v9))))) | 1))
                                                                        _ = v243
                                                                    case 32768:
                                                                        var v244 uint8 = 1
                                                                        _ = v244
                                                                    case 599:
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                        var v245 int = fi228
                                                                        _ = v245
                                                                    }
                                                                    _ = v242
                                                                case "a":
                                                                    switch (^(fi220)) {
                                                                    case 778:
                                                                        if v1 {
                                                                            var v246 int32 = int32(v8)
                                                                            v231 = (v230 / ((v230) | 1))
                                                                            v247 := GenS21{F0: 353, F1: 4}
                                                                            _ = v246
                                                                            _ = v247
                                                                        }
                                                                    case 4294967295:
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            switch fi220 {
                                                                            case 32768:
                                                                                var v248 string = "hello"
                                                                                _ = v248
                                                                            }
                                                                            if (0 > 255) {
                                                                                var v249 uint = 42
                                                                                v250 := GenS9{F0: 138, F1: -176}
                                                                                v2 -= 3.14159
                                                                                _ = v249
                                                                                _ = v250
                                                                            }
                                                                        }
                                                                    case 32768:
                                                                        v251 := GenS28{F0: 65535, F1: "a"}
                                                                        _ = v251
                                                                    }
                                                                case "\n\t":
                                                                    v2 += (((v2 + (v2 * float32(float64(-2.25)))) * (0.0 * ((v2 + v2) - float32(float64(1e10))))) + 3.14159)
                                                                    fmt.Println(v230, v230, v231)
                                                                }
                                                                _ = len(v6)
                                                                _ = v241
                                                            }
                                                            if v1 {
                                                                v4 = 1
                                                                v9 = (v9 & 0)
                                                            }
                                                        } else {
                                                            v9 += (v9 % ((v9) | 1))
                                                            v6 = append(v6, fi228)
                                                            v6 = append(v6, 0)
                                                        }
                                                        v252 := GenS29{F0: v9, F1: 558, F2: v9}
                                                        _ = v233
                                                        _ = v252
                                                    } else {
                                                        if (!true) {
                                                            v231 -= (v230 & (v230 ^ (v231 << ((v9) & 63))))
                                                            switch 806 {
                                                            case 786:
                                                                fmt.Println(fi228, v230, v2)
                                                                var v253 uint = (^(v9))
                                                                _ = v253
                                                            case 2147483647:
                                                                fmt.Println(v230)
                                                            default:
                                                                var v254 int = fi220
                                                                _ = v254
                                                            }
                                                        }
                                                        if (42 == 0) {
                                                            var v255 float32 = float32(1e10)
                                                            if (true || true) {
                                                                v256 := GenS4{F0: v3}
                                                                _ = v256
                                                            } else {
                                                                v257 := GenS17{F0: uint8(v8), F1: 4294967295, F2: v4}
                                                                _ = v257
                                                            }
                                                            _ = v255
                                                        }
                                                    }
                                                }
                                                v4 = v4
                                                switch "" {
                                                case "a":
                                                    switch (fi228 ^ (fi220 - (^(fi228)))) {
                                                    case 814:
                                                        var v258 bool = (42 == 635)
                                                        _ = v258
                                                    case 32767:
                                                        v259 := GenS11{F0: 2147483647, F1: false, F2: (v5 & 0)}
                                                        _ = v259
                                                    case -969:
                                                        v5 = v5
                                                    }
                                                default:
                                                    if len(v6) > 0 {
                                                        v6[0] = fi220
                                                    }
                                                }
                                            }
                                        case 1:
                                            var v260 float64 = 0.1
                                            _ = v260
                                        case 127:
                                            var v261 uint16 = v5
                                            _ = v261
                                        default:
                                            var v262 uint = v9
                                            _ = v262
                                        }
                                    }
                                    _ = v230
                                    _ = v231
                                } else {
                                    if ((55 != 2147483647) || v1) {
                                        fmt.Println(v3)
                                        if v1 {
                                            var v263 int16 = 0
                                            if (!v1) {
                                                v264 := GenS4{F0: 640}
                                                if v1 {
                                                    switch ((("foo bar" + "") + "foo bar") + "") {
                                                    case "":
                                                        v9 = (v9 * 18446744073709551615)
                                                    case "foo bar":
                                                        if false {
                                                            v265 := GenS30{F0: (v263 * int16(v8)), F1: (v263 - (v263 | (^(v263)))), F2: v9}
                                                            _ = v265
                                                        } else {
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            var v266 int16 = (v263 ^ v263)
                                                            var v267 uint16 = (^(v5))
                                                            switch ("hello" + (("hello" + "foo bar") + "foo bar")) {
                                                            case "foo bar":
                                                                fmt.Println(fi228, v267, fi228)
                                                            case "hello":
                                                                v266 = (v263 ^ v266)
                                                            case "a":
                                                                if false {
                                                                    v2 += float32((0.1 * (1e10 - -2.25)))
                                                                    v268 := GenS16{F0: 9223372036854775807}
                                                                    switch "" {
                                                                    case "":
                                                                        v2 += 0.1
                                                                    }
                                                                    _ = v268
                                                                } else {
                                                                    if v1 {
                                                                        if len(v6) > 0 {
                                                                            v6[0] = fi220
                                                                        }
                                                                        v269 := GenS31{F0: v3, F1: "", F2: false}
                                                                        if false {
                                                                            v6 = append(v6, 4294967295)
                                                                        }
                                                                        _ = v269
                                                                    }
                                                                    fmt.Println(fi228, v266)
                                                                    var v270 int8 = 100
                                                                    _ = v270
                                                                }
                                                                v271 := GenS32{F0: 0, F1: (v1 || ((v2 > float32(1e10)) || v1)), F2: 259}
                                                                v7 -= -9223372036854775808
                                                                _ = v271
                                                            default:
                                                                switch ("hello" + "\n\t") {
                                                                case "a":
                                                                    v6 = append(v6, fi220)
                                                                }
                                                            }
                                                            var v272 uint8 = 2
                                                            fmt.Println(v5, v272)
                                                            _ = v266
                                                            _ = v267
                                                            _ = v272
                                                        }
                                                    case "hello":
                                                        v9 += v9
                                                    default:
                                                        var v273 uint = 255
                                                        _ = v273
                                                    }
                                                    switch (fi228 + (fi220 & (fi228 % ((int(v8)) | 1)))) {
                                                    case -434:
                                                        v263 -= v263
                                                    case -845:
                                                        v6 = append(v6, (fi228 * 0))
                                                    case 9223372036854775807:
                                                        if false {
                                                            switch fi220 {
                                                            case 32768:
                                                                _ = len(v6)
                                                            default:
                                                                fmt.Println(v9, fi228, fi228)
                                                            }
                                                        } else {
                                                            switch (fi228 << (((v9 & 2)) & 63)) {
                                                            case -2147483648:
                                                                var v274 uint16 = 0
                                                                _ = v274
                                                            }
                                                            _ = len(v6)
                                                        }
                                                    default:
                                                        v275 := GenS33{F0: 127, F1: (v8 | v8)}
                                                        _ = v275
                                                    }
                                                } else {
                                                    v7 -= (v7 % ((v7) | 1))
                                                }
                                                switch "a" {
                                                case "hello":
                                                    v276 := GenS33{F0: 128, F1: 18446744073709551615}
                                                    _ = v276
                                                case "\n\t":
                                                    v5 += (v5 & uint16(v9))
                                                case "":
                                                    v3 = (v4 + (^(v4)))
                                                }
                                                _ = v264
                                            }
                                            _ = v263
                                        } else {
                                            v277 := GenS34{F0: 32767, F1: -115, F2: 937}
                                            var v278 float32 = v2
                                            var v279 bool = v1
                                            fmt.Println(v9, fi228)
                                            switch (("hello" + ("\n\t" + "")) + "a") {
                                            case "":
                                                fmt.Println(v279, fi228)
                                            case "foo bar":
                                                v280 := GenS31{F0: (-(v7)), F1: "hello", F2: (1.5 <= 3.14159)}
                                                if (!v279) {
                                                    fmt.Println(v8, v7)
                                                }
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                _ = v280
                                            default:
                                                v7 = v7
                                            }
                                            _ = v277
                                            _ = v278
                                            _ = v279
                                        }
                                        if len(v6) > 0 {
                                            v6[0] = fi220
                                        }
                                    } else {
                                        var v281 int16 = 1000
                                        _ = v281
                                    }
                                    var v282 uint16 = v5
                                    if (false && ((false || ((!(127 <= 2)) || (!(-32768 >= -411)))) && false)) {
                                        v9 = v9
                                        switch (^(fi220)) {
                                        case 32768:
                                            switch ("a" + ("" + "a")) {
                                            case "a":
                                                v283 := GenS35{F0: fi228}
                                                if v1 {
                                                    v284 := GenS36{F0: 42, F1: true, F2: ("\n\t" + "hello")}
                                                    var v285 bool = (-292 <= -1)
                                                    v9 -= (v9 * (v9 + uint(v8)))
                                                    _ = v284
                                                    _ = v285
                                                }
                                                var v286 int16 = 32767
                                                v287 := GenS5{F0: v1}
                                                _ = v283
                                                _ = v286
                                                _ = v287
                                            }
                                        case 32767:
                                            v6 = append(v6, fi220)
                                        case 42:
                                            var v288 uint8 = 128
                                            _ = v288
                                        }
                                    }
                                    _ = v282
                                }
                            case 32768:
                                if len(v6) > 0 {
                                    v6[0] = fi220
                                }
                                var v289 string = "hello"
                                _ = v289
                            default:
                                var v290 int8 = 0
                                _ = v290
                            }
                        }
                        v5 += v5
                        if v1 {
                            var v291 uint64 = v8
                            v292 := GenS18{F0: 209, F1: 2}
                            _ = v291
                            _ = v292
                        }
                    } else {
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v2)
                        case "foo bar":
                            v5 = (^(v5))
                        default:
                            if true {
                                v5 += uint16(v9)
                                var v293 int = fi228
                                v294 := GenS7{F0: 0, F1: v1, F2: 974}
                                _ = v293
                                _ = v294
                            } else {
                                var v295 int16 = 1000
                                v296 := GenS32{F0: 128, F1: ((^(v5)) < (v5 >> ((v9) & 63))), F2: 1}
                                _ = v295
                                _ = v296
                            }
                            v3 = 9223372036854775807
                            if len(v6) > 0 {
                                v6[0] = -2147483648
                            }
                            switch fi228 {
                            case -849:
                                v297 := GenS37{F0: 430, F1: 2}
                                var v298 uint32 = 213
                                var v299 bool = false
                                if v299 {
                                    fmt.Println(v5, v1)
                                    v300 := GenS6{F0: 42, F1: (v298 ^ v298), F2: 461}
                                    var v301 int64 = (v3 & v7)
                                    _ = v300
                                    _ = v301
                                } else {
                                    switch (fi220 >> ((264) & 63)) {
                                    case 127:
                                        var v302 int = 128
                                        _ = v302
                                    default:
                                        if false {
                                            fmt.Println(v298, v1, v299)
                                            v5 = v5
                                        } else {
                                            if true {
                                                v303 := GenS38{F0: v1, F1: "\n\t", F2: -411}
                                                _ = v303
                                            } else {
                                                fmt.Println(v299, v8)
                                            }
                                        }
                                    }
                                    var v304 uint = uint(v5)
                                    _ = v304
                                }
                                _ = v297
                                _ = v298
                                _ = v299
                            case 42:
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                            var v305 int32 = -577
                            _ = v305
                        }
                        v1 = (3.14159 != ((1e10 * (-2.25 / 1.0)) - (1e10 / 1e10)))
                    }
                case -2147483648:
                    switch "hello" {
                    case "a":
                        fmt.Println(v9)
                    case "hello":
                        if len(v6) > 0 {
                            v6[0] = fi220
                        }
                        if len(v6) > 0 {
                            v6[0] = (fi228 - fi220)
                        }
                        var v306 int16 = 189
                        _ = v306
                    }
                    switch "hello" {
                    case "":
                        v5 += 936
                    }
                    v9 -= (v9 / ((65535) | 1))
                default:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                if (false && (!v1)) {
                    if (693 <= v9) {
                        v2 -= v2
                    }
                    v6 = append(v6, fi220)
                } else {
                    if v1 {
                        v307 := GenS33{F0: 255, F1: v8}
                        fmt.Println(fi220, v8)
                        _ = v307
                    } else {
                        v8 = (v8 & (v8 & v8))
                    }
                    switch "" {
                    case "foo bar":
                        v5 = v5
                    case "a":
                        _ = len(v6)
                    case "\n\t":
                        if (42 == v5) {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                    default:
                        var v308 float32 = ((v2 * 3.14159) * -2.25)
                        _ = v308
                    }
                }
                _ = fi228
            }
            _ = fi220
            _ = v227
        }
    }
    fmt.Println(v8, v7)
    fmt.Println(v2, v4, v1)
    switch int(v5) {
    case 32768:
        var v309 uint32 = 2147483647
        _ = v309
    case 9223372036854775807:
        var v310 int = 127
        _ = v310
    }
    for fi311 := 0; fi311 < 18; fi311++ {
        var v312 int32 = -524
        if (1.5 != float64((-2.25 + 0.0))) {
            var v313 uint64 = v8
            _ = v313
        } else {
            for fi314 := 0; fi314 < 18; fi314++ {
                switch (fi311 & fi314) {
                case -1:
                    var v315 uint32 = 2147483648
                    _ = v315
                default:
                    var v316 uint64 = (^(v8))
                    if true {
                        var v317 int32 = (^(v312))
                        _ = v317
                    } else {
                        if v1 {
                            v318 := GenS33{F0: 67, F1: v316}
                            if len(v6) > 0 {
                                v6[0] = (fi311 * (fi314 ^ fi311))
                            }
                            _ = v318
                        } else {
                            if v1 {
                                v319 := GenS20{F0: 65535, F1: (v8 ^ (v316 ^ v8)), F2: 4294967295}
                                if v1 {
                                    v1 = ((90 <= 42) && ((v312 & -1) > v312))
                                    var v320 int32 = 2147483647
                                    _ = v320
                                }
                                switch (((("" + "\n\t") + ("\n\t" + "foo bar")) + "foo bar") + (("\n\t" + "a") + ("" + "foo bar"))) {
                                case "hello":
                                    v312 = (v312 / ((v312) | 1))
                                case "a":
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    var v321 float32 = v2
                                    _ = v321
                                case "\n\t":
                                    var v322 int = (fi314 << ((v9) & 63))
                                    _ = v322
                                }
                                _ = v319
                            }
                        }
                        v6 = append(v6, fi314)
                        v5 = (v5 >> (((v9 & (v9 % (((v9 | (v9 & v9))) | 1)))) & 63))
                        _ = len(v6)
                    }
                    var v323 uint16 = (v5 + v5)
                    _ = v316
                    _ = v323
                }
                _ = len(v6)
                _ = fi314
            }
        }
        _ = fi311
        _ = v312
    }
    v9 = v9
    v7 = (v7 << ((v9) & 63))
    var v324 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v324:", v324)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v324
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
