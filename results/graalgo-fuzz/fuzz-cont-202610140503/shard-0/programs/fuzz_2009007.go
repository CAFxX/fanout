package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int; F1 bool; F2 string }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint; F1 int64; F2 int8 }
type GenS6 struct { F0 int32; F1 int16; F2 uint16 }
type GenS7 struct { F0 int16; F1 int64; F2 uint64 }
type GenS8 struct { F0 uint64; F1 string; F2 uint64 }
type GenS9 struct { F0 int16; F1 int; F2 uint }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int8; F1 string; F2 uint32 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint8; F1 int; F2 int8 }
type GenS14 struct { F0 int64; F1 bool }
type GenS15 struct { F0 uint16; F1 uint64; F2 uint64 }
func genFn0(p0 uint8, p1 int8, p2 uint16) uint8 {
    v6 := []int{42, 258, -337, 770}
    fmt.Println(p0)
    if len(v6) > 0 {
        v6[0] = 42
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p0 >> ((4294967295) & 63))
}
func genFn1() int64 {
    switch "" {
    case "\n\t":
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 uint8 = 127
            if false {
                v8 = (^(v8))
            }
            if (!false) {
                if (true && true) {
                    v9 := []int{42}
                    fmt.Println(fi7, fi7)
                    var v10 uint64 = 327
                    _ = v10
                    _ = v9
                }
                fmt.Println(v8, v8)
                for fi11 := 0; fi11 < 16; fi11++ {
                    if true {
                        v8 += (^(v8))
                    } else {
                        v8 = (v8 + genFn0(v8, 100, 65535))
                    }
                    switch (fi11 * 0) {
                    case -9223372036854775808:
                        v8 = (v8 / ((42) | 1))
                    case -38:
                        fmt.Println(fi7, v8)
                    default:
                        if false {
                            switch (("\n\t" + (("a" + ("\n\t" + "foo bar")) + "")) + ("foo bar" + ("\n\t" + "foo bar"))) {
                            case "\n\t":
                                switch (fi7 * fi11) {
                                case 32768:
                                    v8 = 42
                                    fmt.Println(fi7)
                                case 127:
                                    var v12 int32 = 85
                                    _ = v12
                                }
                            }
                        }
                    }
                    if (((3.14159 - -2.25) > 0.0) || (1.5 != (1.5 / 1.5))) {
                        if true {
                            v8 -= (v8 << ((190) & 63))
                            v13 := GenS1{F0: 2, F1: fi11, F2: 2}
                            _ = v13
                        } else {
                            v14 := GenS1{F0: 4294967295, F1: fi7, F2: 2147483647}
                            _ = v14
                        }
                    } else {
                        v15 := GenS1{F0: 196, F1: (fi7 / ((fi11) | 1)), F2: 500}
                        var v16 uint64 = 508
                        var v17 uint32 = 65535
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    _ = fi11
                }
            } else {
                v18 := GenS1{F0: 2, F1: fi7, F2: 42}
                v8 = (^(v8))
                _ = v18
            }
            _ = fi7
            _ = v8
        }
    case "foo bar":
        var v19 uint = 255
        _ = v19
    case "hello":
        var v20 float64 = -2.25
        _ = v20
    }
    v21 := []int{-670, 761}
    _ = v21
    return 469
}

func main() {
    var v1 uint16 = 763
    v2 := make(map[uint32]bool)
    v3 := []uint32{982, 1, 2}
    var v4 int64 = -9223372036854775808
    v5 := make(map[int]int)
    v1 += (v1 / (((^(v1))) | 1))
    switch -2147483648 {
    case 42:
        v22 := GenS2{F0: 1000, F1: 671}
        _ = v22
    case -9223372036854775808:
        var v23 int8 = -1
        _ = v23
    case -926:
        var v24 float32 = (0.0 - float32(0.1))
        _ = v24
    }
    if len(v3) > 0 {
        v3[0] = 14
    }
    var v25 string = "\n\t"
    v26 := GenS2{F0: -32768, F1: 927}
    if (float32((0.1 - 1.5)) > 1.5) {
        if true {
            for fi27 := 0; fi27 < 9; fi27++ {
                v4 -= genFn1()
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    var v28 float64 = 1.5
                    v2[891] = false
                    _ = v28
                }
                _ = fi27
            }
        } else {
            switch (((v25 + ((v25 + v25) + "a")) + ((((("a" + "hello") + v25) + (("hello" + "\n\t") + v25)) + ((("a" + "a") + "a") + v25)) + "hello")) + v25) {
            case "hello":
                var v29 int32 = 579
                _ = v29
            case "":
                switch 128 {
                case 910:
                    if false {
                        if true {
                            switch v25 {
                            case "a":
                                v30 := GenS3{F0: 4294967295, F1: false, F2: v25}
                                _ = v30
                            default:
                                if true {
                                    for fi31 := 0; fi31 < 12; fi31++ {
                                        switch "a" {
                                        case "hello":
                                            var v32 float32 = float32(0.1)
                                            _ = v32
                                        case "":
                                            for fi33 := 0; fi33 < 14; fi33++ {
                                                var v34 string = v25
                                                _ = fi33
                                                _ = v34
                                            }
                                        }
                                        v35 := GenS3{F0: 497, F1: false, F2: "\n\t"}
                                        if false {
                                            fmt.Println(v4)
                                            switch v25 {
                                            case "":
                                                fmt.Println(v25, fi31)
                                            case "a":
                                                for fi36 := 0; fi36 < 8; fi36++ {
                                                    switch (fi36 % ((0) | 1)) {
                                                    case -2147483648:
                                                        if (!(79 < 127)) {
                                                            v1 -= v1
                                                            if false {
                                                                var v37 float64 = -2.25
                                                                if false {
                                                                    switch (v25 + (v25 + v25)) {
                                                                    case "a":
                                                                        if ((fi36 - fi31) != fi31) {
                                                                            switch (fi31 ^ -1) {
                                                                            case 2147483647:
                                                                                v38 := GenS3{F0: 32768, F1: (!true), F2: ""}
                                                                                _ = v38
                                                                            case 772:
                                                                                var v39 uint = 255
                                                                                _ = v39
                                                                            }
                                                                        }
                                                                    case "\n\t":
                                                                        if (!false) {
                                                                            v4 = v4
                                                                            v37 -= 1.5
                                                                        } else {
                                                                            v40 := GenS3{F0: (fi36 - (fi31 / ((-9223372036854775808) | 1))), F1: true, F2: (v25 + v25)}
                                                                            if false {
                                                                                switch (fi36 >> ((18446744073709551615) & 63)) {
                                                                                case 2147483647:
                                                                                    fmt.Println(v25, v4)
                                                                                    var v41 int = fi31
                                                                                    var v42 bool = (false || ((^(v1)) == 65535))
                                                                                    _ = v41
                                                                                    _ = v42
                                                                                case 42:
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = 65535
                                                                                    }
                                                                                }
                                                                                if (33 <= 1) {
                                                                                    fmt.Println(fi31)
                                                                                    if true {
                                                                                        v43 := v5[284]
                                                                                        var v44 uint8 = 230
                                                                                        v4 = v4
                                                                                        _ = v43
                                                                                        _ = v44
                                                                                    } else {
                                                                                        v3 = append(v3, 2147483647)
                                                                                        if false {
                                                                                            if false {
                                                                                                switch v25 {
                                                                                                case "foo bar":
                                                                                                    switch fi31 {
                                                                                                    case -829:
                                                                                                        switch v25 {
                                                                                                        case "":
                                                                                                            if (-528 <= -423) {
                                                                                                                var v45 uint8 = 1
                                                                                                                _ = v45
                                                                                                            }
                                                                                                        case "\n\t":
                                                                                                            if (-81 <= -1) {
                                                                                                                v4 = (v4 | v4)
                                                                                                                var v46 int32 = 2147483647
                                                                                                                v47 := GenS4{F0: 898, F1: 2, F2: 1}
                                                                                                                _ = v46
                                                                                                                _ = v47
                                                                                                            }
                                                                                                        case "a":
                                                                                                            if (true && (65535 <= 469)) {
                                                                                                                var v48 int16 = -15
                                                                                                                var v49 float32 = (float32((0.0 + (((v37 + (v37 + 0.1)) / -2.25) * v37))) + ((1.5 / float32(v37)) + 3.14159))
                                                                                                                var v50 bool = false
                                                                                                                _ = v48
                                                                                                                _ = v49
                                                                                                                _ = v50
                                                                                                            }
                                                                                                        }
                                                                                                        if false {
                                                                                                            v51 := v5[(fi36 | fi31)]
                                                                                                            var v52 uint16 = (v1 | 42)
                                                                                                            delete(v5, fi36)
                                                                                                            _ = v51
                                                                                                            _ = v52
                                                                                                        } else {
                                                                                                            var v53 int64 = (v4 >> ((862) & 63))
                                                                                                            _ = v53
                                                                                                        }
                                                                                                        if (true || false) {
                                                                                                            v37 += v37
                                                                                                            switch (fi31 << ((42) & 63)) {
                                                                                                            case 4294967295:
                                                                                                                switch ("foo bar" + v25) {
                                                                                                                case "a":
                                                                                                                    switch (fi36 / ((9223372036854775807) | 1)) {
                                                                                                                    case -2147483648:
                                                                                                                        v37 -= 0.1
                                                                                                                    case -1:
                                                                                                                        var v54 int32 = 1
                                                                                                                        _ = v54
                                                                                                                    case -9223372036854775808:
                                                                                                                        _ = len(v3)
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                v1 = 65535
                                                                                                            }
                                                                                                            if (false && true) {
                                                                                                                var v55 float32 = -2.25
                                                                                                                _ = v55
                                                                                                            } else {
                                                                                                                v1 = v1
                                                                                                                var v56 uint32 = 2147483647
                                                                                                                _ = v56
                                                                                                            }
                                                                                                            var v57 string = "a"
                                                                                                            var v58 int8 = 127
                                                                                                            v59 := GenS4{F0: 782, F1: 18446744073709551615, F2: 2}
                                                                                                            _ = v57
                                                                                                            _ = v58
                                                                                                            _ = v59
                                                                                                        } else {
                                                                                                            if true {
                                                                                                                switch fi36 {
                                                                                                                case -546:
                                                                                                                    v1 = 2
                                                                                                                case 1:
                                                                                                                    var v60 float64 = v37
                                                                                                                    _ = v60
                                                                                                                default:
                                                                                                                    var v61 uint = 255
                                                                                                                    _ = v61
                                                                                                                }
                                                                                                                v62 := GenS1{F0: 303, F1: int(v1), F2: 42}
                                                                                                                _ = v62
                                                                                                            } else {
                                                                                                                fmt.Println(v37, v4, v4)
                                                                                                                switch v25 {
                                                                                                                case "\n\t":
                                                                                                                    var v63 int16 = 0
                                                                                                                    _ = v63
                                                                                                                }
                                                                                                                v64 := v5[fi36]
                                                                                                                if false {
                                                                                                                    if (v25 > "") {
                                                                                                                        var v65 uint = 255
                                                                                                                        fmt.Println(fi31, v64, v65)
                                                                                                                        if len(v3) > 0 {
                                                                                                                            v3[0] = 42
                                                                                                                        }
                                                                                                                        _ = v65
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v1 = 0
                                                                                                                    fmt.Println(v4, fi36)
                                                                                                                }
                                                                                                                var v66 uint8 = 204
                                                                                                                _ = v64
                                                                                                                _ = v66
                                                                                                            }
                                                                                                            fmt.Println(fi36, fi31, v4)
                                                                                                            v67 := GenS3{F0: fi31, F1: (!false), F2: (((v25 + (v25 + v25)) + v25) + ((("\n\t" + v25) + (((v25 + (v25 + (v25 + "foo bar"))) + ("\n\t" + v25)) + v25)) + "a"))}
                                                                                                            _ = v67
                                                                                                        }
                                                                                                    default:
                                                                                                        switch fi36 {
                                                                                                        case 0:
                                                                                                            var v68 float64 = 3.14159
                                                                                                            _ = v68
                                                                                                        }
                                                                                                    }
                                                                                                case "":
                                                                                                    v37 -= 1e10
                                                                                                case "hello":
                                                                                                    switch ("" + "hello") {
                                                                                                    case "foo bar":
                                                                                                        var v69 float64 = (float64(0.0) / v37)
                                                                                                        _ = v69
                                                                                                    case "hello":
                                                                                                        if (int16(v1) < -1) {
                                                                                                            fmt.Println(v4)
                                                                                                            if false {
                                                                                                                v2[997] = true
                                                                                                                if false {
                                                                                                                    v70 := GenS2{F0: 32767, F1: -1}
                                                                                                                    _ = v70
                                                                                                                }
                                                                                                                fmt.Println(v4, fi36)
                                                                                                            } else {
                                                                                                                v25 = v25
                                                                                                                switch v25 {
                                                                                                                case "":
                                                                                                                    if false {
                                                                                                                        v71 := GenS5{F0: 940, F1: v4, F2: 1}
                                                                                                                        _ = v71
                                                                                                                    } else {
                                                                                                                        if ((!(1 != 2)) && (!false)) {
                                                                                                                            var v72 uint8 = 255
                                                                                                                            fmt.Println(fi31)
                                                                                                                            _ = v72
                                                                                                                        } else {
                                                                                                                            var v73 uint32 = 382
                                                                                                                            v73 = (v73 >> ((uint(v1)) & 63))
                                                                                                                            _ = v73
                                                                                                                        }
                                                                                                                        v74 := v2[4294967295]
                                                                                                                        switch 2147483647 {
                                                                                                                        case -20:
                                                                                                                            var v75 uint64 = 4294967295
                                                                                                                            _ = v75
                                                                                                                        case -254:
                                                                                                                            var v76 uint64 = 4294967295
                                                                                                                            _ = v76
                                                                                                                        default:
                                                                                                                            v4 += v4
                                                                                                                        }
                                                                                                                        _ = v74
                                                                                                                    }
                                                                                                                }
                                                                                                                if (v4 <= v4) {
                                                                                                                    if true {
                                                                                                                        v4 = 42
                                                                                                                        var v77 uint64 = 0
                                                                                                                        _ = v77
                                                                                                                    } else {
                                                                                                                        if true {
                                                                                                                            fmt.Println(v37, fi36, v37)
                                                                                                                            v78 := GenS5{F0: 0, F1: (^(v4)), F2: 127}
                                                                                                                            switch (fi31 & fi36) {
                                                                                                                            case 128:
                                                                                                                                if ((^(v4)) >= -1) {
                                                                                                                                    if ((true && (600 > -32768)) && (1 <= 0)) {
                                                                                                                                        switch fi31 {
                                                                                                                                        case 9223372036854775807:
                                                                                                                                            if len(v3) > 0 {
                                                                                                                                                v3[0] = 42
                                                                                                                                            }
                                                                                                                                        case 4294967295:
                                                                                                                                            var v79 float64 = v37
                                                                                                                                            _ = v79
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        if ((v25 + v25) < v25) {
                                                                                                                                            fmt.Println(v25, v1, v4)
                                                                                                                                        }
                                                                                                                                        if ((true && false) && (false || (0 >= 288))) {
                                                                                                                                            v80 := GenS4{F0: 0, F1: 18446744073709551615, F2: 950}
                                                                                                                                            v1 = 394
                                                                                                                                            _ = len(v5)
                                                                                                                                            _ = v80
                                                                                                                                        }
                                                                                                                                        v1 -= v1
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                            v4 = genFn1()
                                                                                                                            _ = v78
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v81 int8 = int8(v1)
                                                                                                                    v81 = -128
                                                                                                                    _ = v81
                                                                                                                } else {
                                                                                                                    var v82 uint = 18446744073709551615
                                                                                                                    _ = v82
                                                                                                                }
                                                                                                            }
                                                                                                        } else {
                                                                                                            v83 := GenS6{F0: -249, F1: 42, F2: v1}
                                                                                                            var v84 uint16 = (v1 / (((v1 ^ (v1 << ((2) & 63)))) | 1))
                                                                                                            v85 := v5[(fi31 ^ (fi36 << ((0) & 63)))]
                                                                                                            fmt.Println(v37, v84, v85)
                                                                                                            v86 := GenS7{F0: 0, F1: (v4 - (v4 / (((v4 * 9223372036854775807)) | 1))), F2: 503}
                                                                                                            if (32767 < 1000) {
                                                                                                                delete(v2, 2147483648)
                                                                                                            } else {
                                                                                                                v87 := v5[fi31]
                                                                                                                var v88 float32 = 1.5
                                                                                                                _ = v87
                                                                                                                _ = v88
                                                                                                            }
                                                                                                            switch (fi31 - int(v1)) {
                                                                                                            case 32768:
                                                                                                                _ = len(v3)
                                                                                                            case -9223372036854775808:
                                                                                                                delete(v5, v85)
                                                                                                                switch int(v84) {
                                                                                                                case 4294967295:
                                                                                                                    var v89 bool = (-89 >= -128)
                                                                                                                    _ = v89
                                                                                                                case 713:
                                                                                                                    var v90 uint32 = 4294967295
                                                                                                                    _ = v90
                                                                                                                case -194:
                                                                                                                    if false {
                                                                                                                        var v91 uint8 = 39
                                                                                                                        _ = v91
                                                                                                                    } else {
                                                                                                                        switch 320 {
                                                                                                                        case 139:
                                                                                                                            _ = len(v3)
                                                                                                                        }
                                                                                                                        fmt.Println(fi31, v85, v25)
                                                                                                                    }
                                                                                                                }
                                                                                                                v85 = (fi31 | (v85 / (((fi31 ^ (-(fi31)))) | 1)))
                                                                                                            case 902:
                                                                                                                v25 = v25
                                                                                                            }
                                                                                                            if true {
                                                                                                                var v92 uint = 862
                                                                                                                v37 += v37
                                                                                                                _ = v92
                                                                                                            }
                                                                                                            _ = v83
                                                                                                            _ = v84
                                                                                                            _ = v85
                                                                                                            _ = v86
                                                                                                        }
                                                                                                    case "\n\t":
                                                                                                        if false {
                                                                                                            switch 32768 {
                                                                                                            case -7:
                                                                                                                if (false || (65535 > 657)) {
                                                                                                                    v5[(-(fi31))] = -9223372036854775808
                                                                                                                }
                                                                                                            default:
                                                                                                                v4 -= v4
                                                                                                            }
                                                                                                            var v93 bool = true
                                                                                                            _ = v93
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                switch "foo bar" {
                                                                                                case "\n\t":
                                                                                                    v4 += genFn1()
                                                                                                }
                                                                                                v94 := GenS8{F0: 2, F1: "hello", F2: 2}
                                                                                                _ = v94
                                                                                            } else {
                                                                                                if true {
                                                                                                    v95 := GenS2{F0: 0, F1: 0}
                                                                                                    _ = v95
                                                                                                } else {
                                                                                                    if false {
                                                                                                        v96 := GenS9{F0: int16(v4), F1: fi31, F2: 591}
                                                                                                        _ = v96
                                                                                                    }
                                                                                                    v25 = ("" + ("\n\t" + v25))
                                                                                                    var v97 string = "\n\t"
                                                                                                    _ = v97
                                                                                                }
                                                                                                if (true && true) {
                                                                                                    fmt.Println(v4)
                                                                                                    var v98 uint64 = 843
                                                                                                    var v99 int = (^(fi31))
                                                                                                    fmt.Println(fi36)
                                                                                                    if (false && ((true && (0 > genFn1())) && false)) {
                                                                                                        switch v25 {
                                                                                                        case "hello":
                                                                                                            switch (((((v25 + (v25 + (v25 + v25))) + v25) + ((v25 + (v25 + v25)) + v25)) + v25) + (v25 + ("\n\t" + (("" + ("hello" + v25)) + (v25 + v25))))) {
                                                                                                            case "":
                                                                                                                fmt.Println(fi31)
                                                                                                            }
                                                                                                            v100 := GenS7{F0: -32768, F1: v4, F2: v98}
                                                                                                            var v101 uint = 18446744073709551615
                                                                                                            _ = v100
                                                                                                            _ = v101
                                                                                                        case "foo bar":
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                        }
                                                                                                        v102 := GenS8{F0: v98, F1: v25, F2: (v98 % ((v98) | 1))}
                                                                                                        v103 := GenS3{F0: (fi36 + 32768), F1: (42 < 0), F2: ("foo bar" + (v25 + v25))}
                                                                                                        _ = v102
                                                                                                        _ = v103
                                                                                                    } else {
                                                                                                        var v104 int = 920
                                                                                                        _ = v104
                                                                                                    }
                                                                                                    _ = v98
                                                                                                    _ = v99
                                                                                                } else {
                                                                                                    v2[2] = true
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v4 -= v4
                                                                                            var v105 int32 = -294
                                                                                            var v106 uint16 = (v1 / ((v1) | 1))
                                                                                            _ = v105
                                                                                            _ = v106
                                                                                        }
                                                                                    }
                                                                                    v4 -= genFn1()
                                                                                    _ = len(v2)
                                                                                }
                                                                                v107 := GenS3{F0: (fi36 ^ -2147483648), F1: true, F2: (("a" + "\n\t") + ("a" + ("a" + (v25 + ("foo bar" + ("" + v25))))))}
                                                                                v1 -= (v1 - v1)
                                                                                _ = v107
                                                                            } else {
                                                                                if (v4 > v4) {
                                                                                    v1 = 2
                                                                                }
                                                                            }
                                                                            _ = v40
                                                                        }
                                                                    case "":
                                                                        v1 += 42
                                                                    default:
                                                                        fmt.Println(v1, v37)
                                                                    }
                                                                    switch (fi31 / (((fi31 ^ -847)) | 1)) {
                                                                    case 2147483647:
                                                                        v108 := GenS1{F0: 2, F1: fi36, F2: 42}
                                                                        _ = v108
                                                                    case 127:
                                                                        var v109 int32 = -2147483648
                                                                        _ = v109
                                                                    }
                                                                    var v110 uint64 = 817
                                                                    if (v1 != 255) {
                                                                        v111 := GenS10{F0: 1}
                                                                        switch fi36 {
                                                                        case 4:
                                                                            v112 := GenS10{F0: (v4 % ((genFn1()) | 1))}
                                                                            _ = v112
                                                                        default:
                                                                            v113 := GenS11{F0: 0, F1: (v25 + ((v25 + "hello") + "a")), F2: 0}
                                                                            var v114 uint = 609
                                                                            v115 := GenS8{F0: uint64(v1), F1: "\n\t", F2: 992}
                                                                            v116 := GenS12{F0: fi36}
                                                                            _ = v113
                                                                            _ = v114
                                                                            _ = v115
                                                                            _ = v116
                                                                        }
                                                                        _ = v111
                                                                    } else {
                                                                        fmt.Println(v25)
                                                                        if false {
                                                                            var v117 int32 = -2147483648
                                                                            v25 = v25
                                                                            v25 = "foo bar"
                                                                            v117 -= (v117 >> ((1) & 63))
                                                                            _ = v117
                                                                        } else {
                                                                            if false {
                                                                                var v118 int16 = 32767
                                                                                v119 := v5[(fi36 - -133)]
                                                                                _ = v118
                                                                                _ = v119
                                                                            } else {
                                                                                v3 = append(v3, uint32(v110))
                                                                                switch (v25 + "hello") {
                                                                                case "hello":
                                                                                    if ((uint(v110) == 732) && (!false)) {
                                                                                        v120 := GenS1{F0: 2, F1: fi31, F2: 954}
                                                                                        _ = v120
                                                                                    } else {
                                                                                        var v121 uint16 = v1
                                                                                        if (-1 > 792) {
                                                                                            fmt.Println(fi31)
                                                                                            fmt.Println(v121)
                                                                                            var v122 int16 = 805
                                                                                            _ = v122
                                                                                        } else {
                                                                                            fmt.Println(v121, v4)
                                                                                        }
                                                                                        _ = v121
                                                                                    }
                                                                                    delete(v2, 2147483648)
                                                                                    v123 := GenS5{F0: 0, F1: v4, F2: -1}
                                                                                    _ = v123
                                                                                default:
                                                                                    v1 -= v1
                                                                                }
                                                                                switch (^(fi31)) {
                                                                                case 4294967295:
                                                                                    v110 += 873
                                                                                case -9223372036854775808:
                                                                                    if false {
                                                                                        var v124 int32 = 0
                                                                                        var v125 float32 = 1e10
                                                                                        _ = v124
                                                                                        _ = v125
                                                                                    }
                                                                                default:
                                                                                    v4 += (v4 / ((v4) | 1))
                                                                                }
                                                                            }
                                                                            fmt.Println(v25)
                                                                            v126 := GenS13{F0: 127, F1: fi36, F2: 104}
                                                                            _ = v126
                                                                        }
                                                                        v1 += (^(v1))
                                                                    }
                                                                    _ = v110
                                                                } else {
                                                                    v25 = v25
                                                                    switch (fi36 + (fi36 >> ((1) & 63))) {
                                                                    case -1:
                                                                        v2[2147483647] = ((false || ((((float64(1e10) + v37) * v37) * 1e10) < (1e10 * -2.25))) && false)
                                                                    case 127:
                                                                        v127 := GenS7{F0: -629, F1: (v4 + (-(v4))), F2: 42}
                                                                        _ = v127
                                                                    case 243:
                                                                        v128 := GenS8{F0: 0, F1: "hello", F2: 18446744073709551615}
                                                                        _ = v128
                                                                    }
                                                                }
                                                                if (0.1 < 0.1) {
                                                                    var v129 string = "hello"
                                                                    delete(v2, 2147483648)
                                                                    _ = v129
                                                                }
                                                                _ = v37
                                                            }
                                                        }
                                                    }
                                                    v25 = ("hello" + v25)
                                                    _ = fi36
                                                }
                                            default:
                                                _ = len(v2)
                                            }
                                        }
                                        _ = fi31
                                        _ = v35
                                    }
                                }
                            }
                            v25 = ((("" + "foo bar") + "") + ("\n\t" + v25))
                            v2[42] = false
                        } else {
                            var v130 int = 1
                            var v131 int8 = 100
                            _ = v130
                            _ = v131
                        }
                        fmt.Println(v25)
                    } else {
                        var v132 uint16 = (^(v1))
                        switch 0 {
                        case 127:
                            var v133 int32 = 1
                            _ = v133
                        case -9223372036854775808:
                            if (!true) {
                                v134 := GenS13{F0: 183, F1: 762, F2: 35}
                                v135 := GenS14{F0: -9223372036854775808, F1: true}
                                for fi136 := 0; fi136 < 20; fi136++ {
                                    switch (v25 + v25) {
                                    case "a":
                                        var v137 int = fi136
                                        _ = v137
                                    case "foo bar":
                                        var v138 float64 = ((1e10 + 1.5) + 1.5)
                                        _ = v138
                                    default:
                                        var v139 int16 = -32768
                                        _ = v139
                                    }
                                    var v140 int = fi136
                                    switch (v25 + "") {
                                    case "":
                                        v141 := GenS5{F0: 1, F1: v4, F2: -128}
                                        _ = v141
                                    case "a":
                                        var v142 uint32 = 42
                                        _ = v142
                                    }
                                    _ = fi136
                                    _ = v140
                                }
                                _ = v134
                                _ = v135
                            } else {
                                var v143 uint32 = 1
                                var v144 uint16 = (v1 & (v132 | (v1 / ((65535) | 1))))
                                if false {
                                    v1 += (^(v132))
                                    v1 += (^(v144))
                                }
                                _ = v143
                                _ = v144
                            }
                        default:
                            for fi145 := 0; fi145 < 4; fi145++ {
                                var v146 int = fi145
                                var v147 bool = true
                                v146 += int(v4)
                                _ = fi145
                                _ = v146
                                _ = v147
                            }
                        }
                        var v148 uint64 = 2
                        _ = v132
                        _ = v148
                    }
                default:
                    v149 := GenS2{F0: 0, F1: 1}
                    _ = v149
                }
            }
            delete(v2, 2)
        }
        v25 = "a"
        switch 9223372036854775807 {
        case 32767:
            v1 = v1
        case -9223372036854775808:
            v5[2147483647] = 0
        default:
            fmt.Println(v1)
        }
    } else {
        switch -778 {
        case 128:
            v1 = (v1 % (((v1 | (^(v1)))) | 1))
        case 32767:
            v150 := GenS5{F0: 2, F1: (^(v4)), F2: -1}
            _ = v150
        }
        if ((!(false && true)) || (!("" != "a"))) {
            if (4294967295 == 522) {
                fmt.Println(v25)
            }
            v25 = v25
        } else {
            var v151 uint32 = 264
            _ = v151
        }
        for fi152 := 0; fi152 < 8; fi152++ {
            var v153 int8 = -54
            if (!(2147483648 == 401)) {
                var v154 bool = false
                _ = v154
            } else {
                _ = len(v2)
                v1 = 32767
            }
            _ = fi152
            _ = v153
        }
    }
    v25 = v25
    var v155 uint = 42
    var v156 int = 4294967295
    v157 := GenS15{F0: (^(v1)), F1: 42, F2: 42}
    var v158 int16 = 42
    v159 := GenS15{F0: 2, F1: 1, F2: 262}
    switch -386 {
    case 42:
        var v160 bool = ((-(v4)) == -416)
        _ = v160
    case 986:
        var v161 float64 = (0.1 - (1e10 * (0.0 * 0.1)))
        _ = v161
    case -2147483648:
        v155 = (v155 ^ v155)
    default:
        v162 := GenS9{F0: (^(v158)), F1: v156, F2: v155}
        _ = v162
    }
    v4 -= genFn1()
    v156 = v156
    v156 = 127
    switch "foo bar" {
    case "\n\t":
        var v163 int = (v156 >> (((^(v155))) & 63))
        _ = v163
    }
    fmt.Println(v155, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v155:", v155)
    fmt.Println("v156:", v156)
    fmt.Println("v158:", v158)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v155
    _ = v156
    _ = v157
    _ = v158
    _ = v159
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
