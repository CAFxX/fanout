package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS10 struct { F0 bool; F1 int32; F2 uint8 }
type GenS11 struct { F0 uint; F1 int; F2 int }
type GenS12 struct { F0 int; F1 bool }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint; F1 uint64; F2 int }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 int; F1 int32; F2 int32 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 int64; F1 uint; F2 bool }
type GenS21 struct { F0 int16; F1 int8 }
type GenS22 struct { F0 int16; F1 uint; F2 uint8 }
type GenS23 struct { F0 uint16; F1 int32 }
type GenS24 struct { F0 string }

func main() {
    v1 := make(map[int]int)
    var v2 bool = false
    var v3 uint = 255
    v4 := []int{2147483647, -391, 0, 547}
    var v5 int16 = 42
    switch 9223372036854775807 {
    case 1:
        switch ("" + "foo bar") {
        case "foo bar":
            for fi6 := 0; fi6 < 10; fi6++ {
                for fi7 := 0; fi7 < 5; fi7++ {
                    var v8 uint64 = 0
                    v8 -= v8
                    _ = fi7
                    _ = v8
                }
                if true {
                    if (v2 || v2) {
                        v9 := GenS1{F0: 42}
                        _ = v9
                    }
                    for fi10 := 0; fi10 < 11; fi10++ {
                        switch "foo bar" {
                        case "foo bar":
                            if v2 {
                                if false {
                                    v11 := GenS2{F0: 18446744073709551615, F1: 901, F2: 65535}
                                    var v12 bool = (!v2)
                                    fmt.Println(fi6, fi6)
                                    var v13 bool = (((v5 % ((v5) | 1)) <= v5) && v2)
                                    _ = v11
                                    _ = v12
                                    _ = v13
                                }
                                v14 := GenS1{F0: 2147483648}
                                _ = v14
                            }
                        case "hello":
                            v15 := v1[(fi6 - 1)]
                            _ = v15
                        }
                        _ = fi10
                    }
                }
                for fi16 := 0; fi16 < 10; fi16++ {
                    switch "a" {
                    case "a":
                        var v17 float64 = float64(0.0)
                        _ = v17
                    case "hello":
                        fmt.Println(v3, fi16)
                    case "foo bar":
                        v18 := GenS3{F0: 822}
                        _ = v18
                    default:
                        v19 := GenS3{F0: 144}
                        _ = v19
                    }
                    _ = fi16
                }
                _ = fi6
            }
        case "hello":
            v5 = v5
        case "a":
            v20 := GenS1{F0: 2}
            _ = v20
        default:
            v21 := GenS3{F0: 4294967295}
            _ = v21
        }
    }
    if (2147483647 >= -770) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if v2 {
            switch 127 {
            case 4294967295:
                if true {
                    delete(v1, 9223372036854775807)
                    v3 = (v3 | v3)
                } else {
                    fmt.Println(v5, v3, v5)
                    for fi22 := 0; fi22 < 15; fi22++ {
                        if (541 != 2147483648) {
                            v2 = false
                        } else {
                            v23 := GenS3{F0: 212}
                            _ = v23
                        }
                        var v24 int = fi22
                        v2 = (-9223372036854775808 != -9223372036854775808)
                        _ = fi22
                        _ = v24
                    }
                    var v25 float64 = 0.1
                    _ = v25
                }
            case 127:
                if v2 {
                    v5 += v5
                }
            }
        } else {
            if v2 {
                v26 := GenS4{F0: 785}
                switch "a" {
                case "\n\t":
                    v2 = true
                }
                if len(v4) > 0 {
                    v4[0] = -995
                }
                _ = v26
            }
            switch 42 {
            case 4294967295:
                v3 = 516
            }
        }
        switch 0 {
        case 32767:
            for fi27 := 0; fi27 < 18; fi27++ {
                var v28 uint64 = 4294967295
                _ = fi27
                _ = v28
            }
        case 127:
            for fi29 := 0; fi29 < 6; fi29++ {
                var v30 uint64 = 1
                for fi31 := 0; fi31 < 3; fi31++ {
                    v3 -= v3
                    var v32 float64 = ((((1e10 * ((((0.1 + ((float64((1e10 * (((3.14159 - 3.14159) / 1e10) / (0.0 - float32(((-2.25 + -2.25) - 0.1)))))) * float64(0.1)) + 1e10)) + ((3.14159 + 0.1) - 1.5)) / (((1.5 * (3.14159 + (((0.1 * 0.1) / (1e10 * -2.25)) - -2.25))) + 0.1) / 3.14159)) + 1.5)) + 1e10) * -2.25) - 3.14159)
                    switch (fi29 / ((fi29) | 1)) {
                    case -2147483648:
                        v30 -= v30
                    case -125:
                        delete(v1, (fi29 >> (((v3 | (v3 - v3))) & 63)))
                    }
                    switch fi31 {
                    case -2147483648:
                        switch (fi29 - (fi29 - 377)) {
                        case 9223372036854775807:
                            fmt.Println(v32)
                        default:
                            if true {
                                if v2 {
                                    delete(v1, (fi31 >> ((v3) & 63)))
                                } else {
                                    if v2 {
                                        fmt.Println(v5, v3)
                                        v33 := GenS1{F0: 4294967295}
                                        v34 := GenS3{F0: uint64(v3)}
                                        _ = v33
                                        _ = v34
                                    } else {
                                        if v2 {
                                            _ = len(v4)
                                        }
                                        v35 := GenS1{F0: 2147483647}
                                        _ = v35
                                    }
                                    v36 := GenS5{F0: fi31}
                                    _ = v36
                                }
                                v3 += (v3 & v3)
                                if true {
                                    fmt.Println(v30, v2, fi31)
                                    v2 = (true && true)
                                    var v37 uint32 = 731
                                    _ = v37
                                } else {
                                    v38 := GenS4{F0: fi31}
                                    fmt.Println(fi29, v2)
                                    switch ("foo bar" + "hello") {
                                    case "foo bar":
                                        fmt.Println(fi31)
                                        delete(v1, 2147483647)
                                    case "hello":
                                        fmt.Println(v2, v3)
                                    }
                                    switch (fi31 & (fi31 & (^(fi29)))) {
                                    case 32767:
                                        v39 := v1[(-(fi29))]
                                        _ = v39
                                    case 766:
                                        _ = len(v1)
                                    case -293:
                                        v3 -= (v3 >> (((^(v3))) & 63))
                                    default:
                                        var v40 int16 = (v5 % ((v5) | 1))
                                        if (!(true && v2)) {
                                            v41 := GenS6{F0: -1}
                                            _ = v41
                                        } else {
                                            v42 := GenS1{F0: uint32(v3)}
                                            if v2 {
                                                if true {
                                                    if v2 {
                                                        var v43 uint = v3
                                                        v44 := GenS7{F0: 2, F1: (v3 << (((v3 - 344)) & 63))}
                                                        var v45 uint64 = (v30 % ((2) | 1))
                                                        var v46 uint = (^(v3))
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                        _ = v46
                                                    } else {
                                                        var v47 string = ("\n\t" + (("hello" + "hello") + "foo bar"))
                                                        switch v47 {
                                                        case "hello":
                                                            if false {
                                                                v2 = true
                                                                v32 -= (float64(((float32(0.1) - 0.1) * (((-2.25 - ((3.14159 - float32((1.5 + (1e10 / 0.1)))) * 3.14159)) + 1e10) + 0.1))) * (((0.1 / ((v32 + (v32 * ((v32 / v32) + 1e10))) + v32)) - v32) + v32))
                                                                switch (v47 + v47) {
                                                                case "a":
                                                                    if (42 <= 385) {
                                                                        if false {
                                                                            switch fi29 {
                                                                            case 9223372036854775807:
                                                                                if (0 != 42) {
                                                                                    v5 += v40
                                                                                    v48 := GenS4{F0: fi31}
                                                                                    delete(v1, (fi29 >> ((314) & 63)))
                                                                                    _ = v48
                                                                                } else {
                                                                                    fmt.Println(v3)
                                                                                }
                                                                            case 0:
                                                                                var v49 uint8 = 1
                                                                                _ = v49
                                                                            case -22:
                                                                                if (!(v2 && (!v2))) {
                                                                                    v32 = (v32 + (((0.0 + float64(0.0)) + float64(float32((float64(float32((1e10 + v32))) - 0.0)))) / (1.5 * v32)))
                                                                                }
                                                                            default:
                                                                                v32 = 3.14159
                                                                            }
                                                                            var v50 uint = (^(v3))
                                                                            switch fi31 {
                                                                            case -683:
                                                                                var v51 int16 = v5
                                                                                _ = v51
                                                                            default:
                                                                                if (!v2) {
                                                                                    v2 = (!(v40 >= v40))
                                                                                    v52 := GenS8{F0: (v40 % (((v40 << ((v3) & 63))) | 1)), F1: int8(v5)}
                                                                                    var v53 uint = 255
                                                                                    _ = v52
                                                                                    _ = v53
                                                                                }
                                                                            }
                                                                            _ = v50
                                                                        } else {
                                                                            v54 := GenS1{F0: 4294967295}
                                                                            var v55 uint16 = 72
                                                                            if false {
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                                var v56 int = (fi29 | -1)
                                                                                if ((v32 + v32) <= 3.14159) {
                                                                                    v47 = ("a" + (("\n\t" + v47) + ""))
                                                                                    var v57 int64 = 824
                                                                                    switch (fi29 ^ 2147483647) {
                                                                                    case 963:
                                                                                        var v58 bool = false
                                                                                        _ = v58
                                                                                    case -1:
                                                                                        if false {
                                                                                            v4 = append(v4, fi31)
                                                                                            if true {
                                                                                                switch (v47 + v47) {
                                                                                                case "\n\t":
                                                                                                    v1[fi29] = (^(fi31))
                                                                                                default:
                                                                                                    v59 := GenS9{F0: (v30 / ((v30) | 1)), F1: 42, F2: v30}
                                                                                                    _ = v59
                                                                                                }
                                                                                            } else {
                                                                                                _ = len(v4)
                                                                                                v60 := GenS3{F0: v30}
                                                                                                var v61 uint16 = 2
                                                                                                if (v2 || v2) {
                                                                                                    fmt.Println(v2, v5, v55)
                                                                                                }
                                                                                                _ = v60
                                                                                                _ = v61
                                                                                            }
                                                                                            if v2 {
                                                                                                if len(v4) >= 2 {
                                                                                                    v4 = v4[1:len(v4)]
                                                                                                }
                                                                                                v32 += v32
                                                                                                if len(v4) > 0 {
                                                                                                    v4[0] = fi31
                                                                                                }
                                                                                            }
                                                                                            v55 = v55
                                                                                            if (-2.25 >= v32) {
                                                                                                switch v47 {
                                                                                                case "a":
                                                                                                    if ((v3 % (((v3 ^ v3)) | 1)) > (v3 & 18446744073709551615)) {
                                                                                                        switch (fi29 << (((^(v3))) & 63)) {
                                                                                                        case 128:
                                                                                                            fmt.Println(v55)
                                                                                                        }
                                                                                                        if false {
                                                                                                            v57 -= (v57 ^ v57)
                                                                                                        }
                                                                                                        if v2 {
                                                                                                            v62 := GenS10{F0: (((v40 < -32768) || true) && (false && (v2 && ((("foo bar" + "\n\t") == "") && v2)))), F1: -1, F2: 42}
                                                                                                            _ = v62
                                                                                                        } else {
                                                                                                            v63 := GenS7{F0: v55, F1: 65535}
                                                                                                            var v64 uint8 = 52
                                                                                                            if (v55 != v55) {
                                                                                                                if (v55 > v55) {
                                                                                                                    v64 += 0
                                                                                                                    switch -555 {
                                                                                                                    case -1000:
                                                                                                                        v3 += (^(v3))
                                                                                                                    case 1:
                                                                                                                        var v65 int = 4294967295
                                                                                                                        _ = v65
                                                                                                                    default:
                                                                                                                        v66 := GenS3{F0: v30}
                                                                                                                        _ = v66
                                                                                                                    }
                                                                                                                }
                                                                                                                var v67 string = v47
                                                                                                                _ = v67
                                                                                                            } else {
                                                                                                                if ((!(!(!(v2 && true)))) && (255 <= uint8(v55))) {
                                                                                                                    v68 := GenS2{F0: (v30 - 1), F1: int32(v30), F2: (^(v55))}
                                                                                                                    var v69 int = fi31
                                                                                                                    _ = v68
                                                                                                                    _ = v69
                                                                                                                } else {
                                                                                                                    var v70 bool = (fi31 <= (v56 ^ fi31))
                                                                                                                    v47 = ((v47 + "hello") + v47)
                                                                                                                    if v70 {
                                                                                                                        var v71 int8 = 0
                                                                                                                        var v72 int64 = (v57 >> ((v3) & 63))
                                                                                                                        var v73 string = v47
                                                                                                                        _ = v71
                                                                                                                        _ = v72
                                                                                                                        _ = v73
                                                                                                                    } else {
                                                                                                                        v32 = (float64(1e10) - ((0.1 + v32) / v32))
                                                                                                                        var v74 string = (v47 + v47)
                                                                                                                        _ = v74
                                                                                                                    }
                                                                                                                    v2 = v2
                                                                                                                    _ = v70
                                                                                                                }
                                                                                                                v30 = (v30 | (v30 % ((0) | 1)))
                                                                                                                if v2 {
                                                                                                                    v75 := GenS4{F0: (-(fi29))}
                                                                                                                    v76 := GenS9{F0: v30, F1: 4294967295, F2: v30}
                                                                                                                    _ = v75
                                                                                                                    _ = v76
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v63
                                                                                                            _ = v64
                                                                                                        }
                                                                                                    }
                                                                                                case "":
                                                                                                    var v77 float64 = 0.1
                                                                                                    _ = v77
                                                                                                }
                                                                                                var v78 int64 = (v57 + 138)
                                                                                                _ = v78
                                                                                            }
                                                                                            if (0 <= -613) {
                                                                                                v4 = append(v4, 715)
                                                                                                switch (v47 + "foo bar") {
                                                                                                case "hello":
                                                                                                    switch -1 {
                                                                                                    case 32768:
                                                                                                        switch "hello" {
                                                                                                        case "hello":
                                                                                                            v79 := GenS9{F0: (v30 + (v30 >> ((65535) & 63))), F1: 808, F2: uint64(v55)}
                                                                                                            _ = v79
                                                                                                        case "\n\t":
                                                                                                            var v80 uint = (v3 ^ (v3 - (v3 ^ v3)))
                                                                                                            if v2 {
                                                                                                                v5 = v5
                                                                                                            } else {
                                                                                                                var v81 float64 = ((v32 + (v32 - v32)) / v32)
                                                                                                                _ = v81
                                                                                                            }
                                                                                                            _ = v80
                                                                                                        }
                                                                                                    case 9223372036854775807:
                                                                                                        switch "\n\t" {
                                                                                                        case "foo bar":
                                                                                                            v82 := GenS11{F0: v3, F1: (fi31 ^ (^(fi29))), F2: 2147483647}
                                                                                                            _ = v82
                                                                                                        case "":
                                                                                                            switch (v47 + v47) {
                                                                                                            case "foo bar":
                                                                                                                var v83 uint32 = 464
                                                                                                                delete(v1, int(v40))
                                                                                                                _ = v83
                                                                                                            case "":
                                                                                                                fmt.Println(v56)
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            v40 = 42
                                                                                                        }
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    v84 := GenS12{F0: fi29, F1: ((v55 | v55) >= (v55 - v55))}
                                                                                                    _ = v84
                                                                                                case "foo bar":
                                                                                                    var v85 uint8 = 196
                                                                                                    _ = v85
                                                                                                default:
                                                                                                    switch (fi31 % (((-(fi31))) | 1)) {
                                                                                                    case 128:
                                                                                                        fmt.Println(v5)
                                                                                                    case 32767:
                                                                                                        var v86 uint32 = 739
                                                                                                        _ = v86
                                                                                                    case 1:
                                                                                                        if len(v4) >= 2 {
                                                                                                            v4 = v4[1:len(v4)]
                                                                                                        }
                                                                                                    default:
                                                                                                        switch v47 {
                                                                                                        case "":
                                                                                                            if (v2 && v2) {
                                                                                                                if v2 {
                                                                                                                    fmt.Println(v55, v47, v40)
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        v87 := GenS8{F0: v5, F1: 0}
                                                                                                                        fmt.Println(v30, v30)
                                                                                                                        var v88 int32 = -2147483648
                                                                                                                        _ = v87
                                                                                                                        _ = v88
                                                                                                                    }
                                                                                                                    switch "hello" {
                                                                                                                    case "a":
                                                                                                                        var v89 int = (fi31 + fi29)
                                                                                                                        _ = v89
                                                                                                                    case "hello":
                                                                                                                        if (true && ((v30 & (v30 << ((65535) & 63))) != 48)) {
                                                                                                                            fmt.Println(v56)
                                                                                                                            v55 = (v55 << ((676) & 63))
                                                                                                                            v90 := GenS6{F0: 21}
                                                                                                                            fmt.Println(v40)
                                                                                                                            v5 += v40
                                                                                                                            _ = v90
                                                                                                                        }
                                                                                                                    case "foo bar":
                                                                                                                        switch v47 {
                                                                                                                        case "a":
                                                                                                                            var v91 float64 = v32
                                                                                                                            v4 = append(v4, fi31)
                                                                                                                            v47 = v47
                                                                                                                            _ = v91
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v92 := GenS5{F0: 42}
                                                                                                                    var v93 uint64 = v30
                                                                                                                    _ = v92
                                                                                                                    _ = v93
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case 127:
                                                                                        v5 = v40
                                                                                    default:
                                                                                        v40 += 445
                                                                                        if v2 {
                                                                                            fmt.Println(v30, v47)
                                                                                            v57 += (-(v57))
                                                                                        } else {
                                                                                            var v94 string = "hello"
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = 735
                                                                                            }
                                                                                            v47 = "\n\t"
                                                                                            _ = v94
                                                                                        }
                                                                                        v56 += -9223372036854775808
                                                                                    }
                                                                                    _ = v57
                                                                                } else {
                                                                                    v95 := GenS13{F0: 1}
                                                                                    v96 := GenS14{F0: v3, F1: (^(v30)), F2: fi31}
                                                                                    _ = v95
                                                                                    _ = v96
                                                                                }
                                                                                var v97 int32 = 0
                                                                                v40 -= 259
                                                                                _ = v56
                                                                                _ = v97
                                                                            } else {
                                                                                v30 -= v30
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = 32768
                                                                                }
                                                                                switch v47 {
                                                                                case "a":
                                                                                    v47 = ("\n\t" + "foo bar")
                                                                                case "":
                                                                                    v98 := GenS13{F0: 116}
                                                                                    _ = v98
                                                                                }
                                                                            }
                                                                            v5 -= v5
                                                                            v99 := GenS5{F0: (fi31 | (^(fi31)))}
                                                                            v100 := GenS7{F0: (v55 >> ((v3) & 63)), F1: v3}
                                                                            _ = v100
                                                                            _ = v54
                                                                            _ = v55
                                                                            _ = v99
                                                                        }
                                                                        v47 = "\n\t"
                                                                        if false {
                                                                            v40 -= (v5 << ((397) & 63))
                                                                            if (-2.25 != 1.5) {
                                                                                v4 = append(v4, (^(fi31)))
                                                                                var v101 float32 = 1e10
                                                                                v102 := GenS11{F0: v3, F1: (fi29 | (-(fi29))), F2: -9223372036854775808}
                                                                                _ = v101
                                                                                _ = v102
                                                                            } else {
                                                                                v103 := v1[(fi31 << ((306) & 63))]
                                                                                fmt.Println(v5)
                                                                                v5 -= (^(v5))
                                                                                _ = v103
                                                                            }
                                                                        } else {
                                                                            var v104 uint = v3
                                                                            _ = v104
                                                                        }
                                                                    } else {
                                                                        if (2 <= 42) {
                                                                            fmt.Println(v3, v47, fi31)
                                                                        } else {
                                                                            if v2 {
                                                                                delete(v1, (^(fi29)))
                                                                                switch (fi31 | (fi31 | -1)) {
                                                                                case -831:
                                                                                    var v105 uint64 = 2
                                                                                    _ = v105
                                                                                case -2147483648:
                                                                                    switch -1 {
                                                                                    case -552:
                                                                                        v106 := GenS6{F0: 100}
                                                                                        _ = v106
                                                                                    case 32767:
                                                                                        var v107 bool = v2
                                                                                        _ = v107
                                                                                    }
                                                                                }
                                                                            }
                                                                            fmt.Println(v30, v30)
                                                                        }
                                                                        _ = len(v1)
                                                                    }
                                                                case "\n\t":
                                                                    var v108 uint64 = v30
                                                                    fmt.Println(v30)
                                                                    _ = v108
                                                                }
                                                            }
                                                        case "a":
                                                            v32 = 1e10
                                                        case "":
                                                            var v109 bool = (2 > 4294967295)
                                                            _ = v109
                                                        default:
                                                            if (42 <= -100) {
                                                                switch (fi31 << ((v3) & 63)) {
                                                                case 240:
                                                                    var v110 int = (fi29 % (((fi31 % ((int(v3)) | 1))) | 1))
                                                                    _ = v110
                                                                case -107:
                                                                    switch fi29 {
                                                                    case 217:
                                                                        if (v2 || (v32 == float64((3.14159 - ((1e10 + ((0.0 + ((3.14159 + 3.14159) + float32(v32))) / (3.14159 / (0.1 * 1.5)))) / ((1.5 * 0.0) + 1.5)))))) {
                                                                            v111 := GenS1{F0: 4294967295}
                                                                            if ((v32 + 0.0) >= float64(0.0)) {
                                                                                if (!v2) {
                                                                                    v112 := GenS12{F0: fi31, F1: (!(0.0 == v32))}
                                                                                    v5 -= (v40 | (v40 >> ((v3) & 63)))
                                                                                    _ = v112
                                                                                } else {
                                                                                    v30 -= v30
                                                                                }
                                                                            }
                                                                            if ((true && v2) || v2) {
                                                                                v32 = v32
                                                                            } else {
                                                                                v30 = v30
                                                                                var v113 int32 = -652
                                                                                var v114 float64 = v32
                                                                                switch fi31 {
                                                                                case 268:
                                                                                    delete(v1, fi31)
                                                                                default:
                                                                                    var v115 int64 = 0
                                                                                    _ = v115
                                                                                }
                                                                                _ = v113
                                                                                _ = v114
                                                                            }
                                                                            _ = v111
                                                                        }
                                                                        v32 -= v32
                                                                    default:
                                                                        var v116 int = (-(fi29))
                                                                        switch (v116 - int(v5)) {
                                                                        case 127:
                                                                            if (742 > 592) {
                                                                                var v117 int32 = -2147483648
                                                                                v118 := GenS6{F0: -75}
                                                                                fmt.Println(v116, v5, v116)
                                                                                switch v47 {
                                                                                case "foo bar":
                                                                                    if false {
                                                                                        var v119 uint = v3
                                                                                        if true {
                                                                                            v32 = ((v32 - 1.5) * v32)
                                                                                            fmt.Println(fi31, v40)
                                                                                        } else {
                                                                                            fmt.Println(v5, v119)
                                                                                        }
                                                                                        _ = v119
                                                                                    }
                                                                                case "a":
                                                                                    v32 = ((float64(((float32(v32) * 0.1) / 1.0)) / (0.1 / (v32 + (((float64(float32(v32)) / (float64(0.1) / v32)) + v32) * float64(0.1))))) * (1.5 * (0.0 / 3.14159)))
                                                                                }
                                                                                _ = v117
                                                                                _ = v118
                                                                            }
                                                                        default:
                                                                            if ((false || (!true)) && (v5 < v5)) {
                                                                                fmt.Println(fi31, fi31, v3)
                                                                            } else {
                                                                                var v120 uint = (v3 >> ((65535) & 63))
                                                                                v5 = int16(v120)
                                                                                _ = v120
                                                                            }
                                                                        }
                                                                        _ = v116
                                                                    }
                                                                }
                                                            } else {
                                                                v121 := GenS15{F0: 2147483648}
                                                                if ((42 != -9223372036854775808) || false) {
                                                                    v122 := GenS1{F0: 212}
                                                                    _ = v122
                                                                } else {
                                                                    v40 = 602
                                                                    v47 = ((v47 + ("" + (("hello" + v47) + v47))) + (((v47 + (v47 + v47)) + ("foo bar" + "hello")) + "hello"))
                                                                    v123 := GenS6{F0: 42}
                                                                    var v124 uint8 = 42
                                                                    _ = v123
                                                                    _ = v124
                                                                }
                                                                var v125 string = (v47 + "a")
                                                                if (3.14159 == -2.25) {
                                                                    if ((1 > 128) || (v2 && v2)) {
                                                                        switch v125 {
                                                                        case "\n\t":
                                                                            v126 := v1[-464]
                                                                            _ = v126
                                                                        case "hello":
                                                                            v47 = v47
                                                                        default:
                                                                            var v127 uint8 = 255
                                                                            _ = v127
                                                                        }
                                                                        v125 = (v47 + v47)
                                                                        v40 = (v40 + v40)
                                                                    }
                                                                    fmt.Println(fi29)
                                                                }
                                                                fmt.Println(v32, v47, v47)
                                                                if (v2 || ((fi31 >> ((4294967295) & 63)) != fi29)) {
                                                                    var v128 int64 = int64(v40)
                                                                    v129 := GenS1{F0: 124}
                                                                    switch fi31 {
                                                                    case -9223372036854775808:
                                                                        fmt.Println(v5, v2, v2)
                                                                    default:
                                                                        v130 := GenS16{F0: (-(fi29)), F1: 42, F2: 42}
                                                                        _ = v130
                                                                    }
                                                                    _ = v128
                                                                    _ = v129
                                                                } else {
                                                                    if false {
                                                                        var v131 int32 = 0
                                                                        var v132 int16 = v5
                                                                        _ = v131
                                                                        _ = v132
                                                                    } else {
                                                                        v47 = v125
                                                                        v30 -= v30
                                                                        v4 = append(v4, -2147483648)
                                                                    }
                                                                    if (("" + (v125 + v125)) >= "hello") {
                                                                        if (2147483647 > 306) {
                                                                            switch (fi29 ^ -841) {
                                                                            case 0:
                                                                                var v133 int64 = 42
                                                                                _ = v133
                                                                            case 128:
                                                                                v40 = v5
                                                                            default:
                                                                                v134 := GenS17{F0: 4294967295}
                                                                                _ = v134
                                                                            }
                                                                            if len(v4) > 0 {
                                                                                v4[0] = fi29
                                                                            }
                                                                        } else {
                                                                            var v135 int16 = v5
                                                                            v136 := GenS18{F0: -9223372036854775808}
                                                                            _ = v135
                                                                            _ = v136
                                                                        }
                                                                    }
                                                                }
                                                                _ = v121
                                                                _ = v125
                                                            }
                                                        }
                                                        _ = v47
                                                    }
                                                } else {
                                                    var v137 bool = ((v2 || true) && (v2 || v2))
                                                    v138 := GenS14{F0: (^(v3)), F1: v30, F2: -9223372036854775808}
                                                    _ = v137
                                                    _ = v138
                                                }
                                            } else {
                                                if ((v40 & 32767) > v40) {
                                                    v40 -= 1
                                                } else {
                                                    if (v3 != (v3 * (v3 >> ((v3) & 63)))) {
                                                        var v139 int32 = 42
                                                        v40 = (v40 ^ v5)
                                                        _ = v139
                                                    } else {
                                                        v140 := GenS17{F0: 322}
                                                        _ = v140
                                                    }
                                                }
                                            }
                                            v141 := v1[9223372036854775807]
                                            _ = v141
                                            _ = v42
                                        }
                                        v4 = append(v4, fi29)
                                        _ = v40
                                    }
                                    v2 = (42 == -1)
                                    _ = v38
                                }
                            }
                        }
                    default:
                        if v2 {
                            var v142 int = 42
                            v2 = v2
                            v32 -= 1e10
                            _ = v142
                        }
                    }
                    var v143 uint8 = 115
                    _ = fi31
                    _ = v143
                    _ = v32
                }
                _ = fi29
                _ = v30
            }
        case 366:
            _ = len(v1)
        }
    } else {
        switch 479 {
        case 520:
            v144 := GenS4{F0: -2147483648}
            _ = v144
        case -788:
            v145 := GenS16{F0: -477, F1: 93, F2: 0}
            _ = v145
        case 32768:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            for fi146 := 0; fi146 < 19; fi146++ {
                if (false && false) {
                    v3 += (^(v3))
                    if v2 {
                        v147 := GenS3{F0: 1}
                        _ = v147
                    }
                } else {
                    fmt.Println(v2)
                    var v148 int32 = 2147483647
                    _ = v148
                }
                fmt.Println(v3, fi146, v2)
                _ = fi146
            }
        }
        _ = len(v4)
    }
    v3 = 654
    for fi149 := 0; fi149 < 9; fi149++ {
        for fi150 := 0; fi150 < 15; fi150++ {
            if v2 {
                var v151 int = fi149
                _ = v151
            }
            _ = len(v1)
            v152 := GenS19{F0: v5}
            _ = fi150
            _ = v152
        }
        switch fi149 {
        case 128:
            v2 = v2
        default:
            fmt.Println(v5)
        }
        if (42 > 255) {
            v5 = (v5 & (-(v5)))
            v3 -= v3
            if v2 {
                if true {
                    v153 := GenS20{F0: 737, F1: 255, F2: v2}
                    _ = v153
                }
            }
        } else {
            v4 = append(v4, 2147483647)
            v5 = v5
            v5 += (v5 + -1)
        }
        _ = fi149
    }
    v154 := GenS11{F0: (^(v3)), F1: 458, F2: 9223372036854775807}
    for fi155 := 0; fi155 < 4; fi155++ {
        v156 := GenS13{F0: 128}
        _ = fi155
        _ = v156
    }
    var v157 uint8 = 225
    var v158 uint32 = 2
    var v159 int64 = -1
    if (v3 == v3) {
        v5 = (v5 + 1)
        v159 = v159
    } else {
        if v2 {
            for fi160 := 0; fi160 < 11; fi160++ {
                fmt.Println(v5, v2, v3)
                v161 := GenS21{F0: (v5 | -399), F1: -68}
                _ = fi160
                _ = v161
            }
            v162 := v1[4294967295]
            _ = v162
        } else {
            if v2 {
                switch 1 {
                case 42:
                    for fi163 := 0; fi163 < 7; fi163++ {
                        switch ("foo bar" + "hello") {
                        case "\n\t":
                            fmt.Println(fi163)
                        case "a":
                            fmt.Println(fi163, v158, v3)
                        }
                        v3 = (v3 / ((v3) | 1))
                        var v164 string = ""
                        _ = fi163
                        _ = v164
                    }
                default:
                    if (((9223372036854775807 > -9223372036854775808) && ((!v2) && (2 >= 86))) && (!(!v2))) {
                        if ("a" != "foo bar") {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            delete(v1, 32768)
                        } else {
                            switch -1 {
                            case -766:
                                if v2 {
                                    var v165 uint32 = (v158 % ((v158) | 1))
                                    _ = v165
                                }
                            case 2147483647:
                                switch 735 {
                                case -126:
                                    fmt.Println(v159, v2)
                                case 295:
                                    var v166 string = "a"
                                    _ = v166
                                }
                            }
                            if (1 < 0) {
                                switch -9223372036854775808 {
                                case 32767:
                                    var v167 int8 = -106
                                    _ = v167
                                case 315:
                                    switch "" {
                                    case "a":
                                        switch "hello" {
                                        case "hello":
                                            var v168 uint16 = 255
                                            _ = v168
                                        case "\n\t":
                                            for fi169 := 0; fi169 < 4; fi169++ {
                                                v158 += (v158 >> ((v3) & 63))
                                                var v170 bool = ((v5 >> (((v3 - (v3 << (((v3 + v3)) & 63)))) & 63)) < 1)
                                                var v171 bool = (v159 < -806)
                                                _ = fi169
                                                _ = v170
                                                _ = v171
                                            }
                                        default:
                                            _ = len(v4)
                                        }
                                    }
                                }
                            }
                        }
                        v1[9223372036854775807] = 42
                    }
                }
                var v172 uint = v3
                v173 := GenS22{F0: (v5 << ((188) & 63)), F1: (v3 % (((^(v3))) | 1)), F2: (^(v157))}
                _ = v172
                _ = v173
            }
        }
        if ("foo bar" <= "\n\t") {
            for fi174 := 0; fi174 < 12; fi174++ {
                var v175 uint32 = 670
                v176 := GenS23{F0: 2, F1: 2147483647}
                _ = len(v4)
                _ = fi174
                _ = v175
                _ = v176
            }
            if ((^(v5)) < (^(v5))) {
                for fi177 := 0; fi177 < 19; fi177++ {
                    v1[(fi177 / ((fi177) | 1))] = fi177
                    _ = fi177
                }
                for fi178 := 0; fi178 < 3; fi178++ {
                    var v179 uint16 = 65535
                    for fi180 := 0; fi180 < 2; fi180++ {
                        if (0.1 >= float32(1e10)) {
                            v158 -= (v158 | (v158 | v158))
                            if false {
                                if false {
                                    var v181 bool = ((0.1 - 1.5) != 0.0)
                                    _ = v181
                                } else {
                                    var v182 int32 = 233
                                    v179 = v179
                                    _ = v182
                                }
                                if ((v5 | v5) != 42) {
                                    _ = len(v1)
                                    var v183 float64 = (0.1 - 0.0)
                                    _ = v183
                                }
                            } else {
                                switch "" {
                                case "":
                                    fmt.Println(v157)
                                case "hello":
                                    v184 := GenS14{F0: 572, F1: uint64(v179), F2: fi180}
                                    _ = v184
                                }
                                var v185 string = "a"
                                _ = v185
                            }
                            fmt.Println(fi180, v5, v179)
                        } else {
                            v1[int(v3)] = fi178
                        }
                        switch "hello" {
                        case "hello":
                            var v186 float64 = -2.25
                            _ = v186
                        }
                        _ = fi180
                    }
                    v159 -= -1
                    _ = fi178
                    _ = v179
                }
            }
            v157 -= (^(v157))
        } else {
            if (1 != 1) {
                if v2 {
                    for fi187 := 0; fi187 < 16; fi187++ {
                        var v188 uint64 = 0
                        for fi189 := 0; fi189 < 10; fi189++ {
                            v188 += v188
                            _ = fi189
                        }
                        _ = fi187
                        _ = v188
                    }
                }
                var v190 uint = (v3 << ((v3) & 63))
                v191 := GenS11{F0: 0, F1: int(v159), F2: 607}
                _ = v190
                _ = v191
            }
            delete(v1, -643)
            for fi192 := 0; fi192 < 17; fi192++ {
                var v193 bool = (!false)
                for fi194 := 0; fi194 < 4; fi194++ {
                    v195 := v1[(fi194 - (fi192 << (((^(v3))) & 63)))]
                    fmt.Println(v159, v193, v193)
                    _ = fi194
                    _ = v195
                }
                _ = fi192
                _ = v193
            }
        }
        v196 := GenS4{F0: -169}
        _ = v196
    }
    if (-2.25 != 0.0) {
        v197 := GenS2{F0: 898, F1: 0, F2: 1}
        var v198 int8 = 0
        _ = v197
        _ = v198
    }
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    if (v3 <= v3) {
        var v199 bool = true
        _ = v199
    }
    for fi200 := 0; fi200 < 1; fi200++ {
        for fi201 := 0; fi201 < 5; fi201++ {
            v202 := GenS24{F0: "\n\t"}
            _ = len(v4)
            _ = fi201
            _ = v202
        }
        _ = fi200
    }
    var v203 int8 = 100
    var v204 int16 = 1000
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v157:", v157)
    fmt.Println("v158:", v158)
    fmt.Println("v159:", v159)
    fmt.Println("v2:", v2)
    fmt.Println("v203:", v203)
    fmt.Println("v204:", v204)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v154
    _ = v157
    _ = v158
    _ = v159
    _ = v2
    _ = v203
    _ = v204
    _ = v3
    _ = v4
    _ = v5
}
