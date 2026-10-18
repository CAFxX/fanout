package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 bool; F1 bool; F2 uint32 }
type GenS6 struct { F0 int; F1 uint64; F2 int8 }
type GenS7 struct { F0 string; F1 uint8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint8; F1 int8; F2 int8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint8; F1 uint8 }
type GenS16 struct { F0 int; F1 uint8; F2 uint16 }
type GenS17 struct { F0 int; F1 uint8; F2 int64 }
type GenS18 struct { F0 uint16; F1 uint8 }
type GenS19 struct { F0 uint64 }
func genFn0(p0 uint16, p1 int16) int16 {
    for fi6 := 0; fi6 < 12; fi6++ {
        fmt.Println(p1, p0, fi6)
        _ = fi6
    }
    fmt.Println(p1)
    for fi7 := 0; fi7 < 17; fi7++ {
        p0 = (^(p0))
        var v8 int8 = -111
        _ = fi7
        _ = v8
    }
    fmt.Println(p1, p1)
    switch 1 {
    case 9223372036854775807:
        for fi9 := 0; fi9 < 11; fi9++ {
            if (!((18446744073709551615 <= uint(p0)) || true)) {
                p0 = (p0 & 377)
                switch fi9 {
                case -2147483648:
                    v10 := []int{}
                    _ = v10
                case 32767:
                    for fi11 := 0; fi11 < 2; fi11++ {
                        p1 = p1
                        fmt.Println(p0, p0, p1)
                        var v12 uint = 42
                        _ = fi11
                        _ = v12
                    }
                }
                var v13 float64 = (3.14159 / 3.14159)
                _ = v13
            } else {
                p0 -= (^(p0))
            }
            fmt.Println(p1, p1, p1)
            v14 := []int{}
            _ = fi9
            _ = v14
        }
    case 127:
        var v15 int8 = -56
        _ = v15
    }
    _ = p0
    _ = p1
    return (-(p1))
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]int)
    v3 := []int{-2147483648, 4294967295, -127, 2147483647}
    var v4 int = -921
    var v5 int16 = -332
    v16 := GenS1{F0: 255, F1: (v4 % ((int(v1)) | 1))}
    v5 = 206
    if (false || false) {
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := GenS1{F0: 903, F1: 9223372036854775807}
            _ = fi17
            _ = v18
        }
        delete(v2, 127)
        for fi19 := 0; fi19 < 6; fi19++ {
            if false {
                var v20 uint16 = 0
                v21 := GenS1{F0: (v20 + (v20 ^ v20)), F1: (v4 + (v4 & (^(fi19))))}
                var v22 uint8 = 1
                _ = v20
                _ = v21
                _ = v22
            } else {
                _ = len(v2)
                v4 -= (v4 % ((v4) | 1))
            }
            _ = len(v3)
            _ = fi19
        }
    } else {
        switch 1 {
        case 1:
            v23 := GenS1{F0: 1, F1: -373}
            _ = v23
        case 128:
            v3 = append(v3, v4)
        case 127:
            for fi24 := 0; fi24 < 8; fi24++ {
                if false {
                    if false {
                        var v25 int = v4
                        if true {
                            v26 := GenS1{F0: 2, F1: v25}
                            v2[(v25 ^ v4)] = (v4 & -1)
                            for fi27 := 0; fi27 < 9; fi27++ {
                                switch "\n\t" {
                                case "hello":
                                    v25 -= -1
                                case "":
                                    v2[943] = fi24
                                }
                                _ = fi27
                            }
                            _ = v26
                        }
                        v5 = v5
                        _ = v25
                    } else {
                        _ = len(v3)
                        for fi28 := 0; fi28 < 19; fi28++ {
                            var v29 int32 = -181
                            _ = fi28
                            _ = v29
                        }
                        v4 = v4
                    }
                    v4 += (fi24 / ((fi24) | 1))
                }
                for fi30 := 0; fi30 < 5; fi30++ {
                    var v31 uint16 = 42
                    fmt.Println(v4, fi24, v1)
                    var v32 float32 = 0.1
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = fi24
            }
        }
    }
    switch ("foo bar" + "foo bar") {
    case "hello":
        var v33 int8 = -1
        _ = v33
    default:
        v1 = v1
    }
    var v34 uint = (^(v1))
    v1 = (v1 << (((v1 + 18446744073709551615)) & 63))
    var v35 uint32 = uint32(v1)
    switch 32768 {
    case 32767:
        v35 = (v35 << ((372) & 63))
    }
    if (v34 < v34) {
        for fi36 := 0; fi36 < 20; fi36++ {
            switch fi36 {
            case -313:
                v37 := GenS1{F0: 2, F1: (fi36 / ((v4) | 1))}
                _ = v37
            case 0:
                v38 := GenS1{F0: 42, F1: fi36}
                _ = v38
            case 9223372036854775807:
                var v39 float32 = (((0.0 + (0.1 + float32((-2.25 + 0.0)))) - 1.5) * 0.0)
                _ = v39
            default:
                if true {
                    fmt.Println(v5)
                } else {
                    fmt.Println(v34, v5)
                }
            }
            _ = fi36
        }
    } else {
        v4 = 675
    }
    switch v4 {
    case -671:
        if true {
            if false {
                v40 := GenS2{F0: 42, F1: (((float64(float32(1.5)) + 0.0) + 0.1) == (1.5 / (float64((1.5 / ((0.1 / 1e10) / ((3.14159 * -2.25) - (-2.25 * 1.5))))) / 1.5)))}
                fmt.Println(v34, v5)
                _ = v40
            } else {
                for fi41 := 0; fi41 < 4; fi41++ {
                    for fi42 := 0; fi42 < 20; fi42++ {
                        _ = len(v2)
                        v5 = v5
                        v4 = (v4 >> ((v1) & 63))
                        var v43 uint32 = v35
                        _ = fi42
                        _ = v43
                    }
                    switch "\n\t" {
                    case "":
                        if false {
                            var v44 uint8 = 0
                            v45 := GenS3{F0: v44, F1: fi41}
                            _ = v44
                            _ = v45
                        } else {
                            if true {
                                if (127 < 1) {
                                    var v46 float32 = ((1.5 * 1.5) + float32((1.5 + 0.1)))
                                    fmt.Println(v4, v46)
                                    delete(v2, fi41)
                                    _ = v46
                                }
                                if (false || (genFn0(192, v5) >= v5)) {
                                    if (!true) {
                                        if (-171 == 1) {
                                            v5 -= (v5 % (((v5 ^ genFn0(65535, genFn0(42, (^(v5)))))) | 1))
                                            v3 = append(v3, v4)
                                        } else {
                                            _ = len(v3)
                                            var v47 uint16 = 875
                                            _ = v47
                                        }
                                        var v48 int16 = (v5 >> (((v34 ^ v34)) & 63))
                                        var v49 uint = uint(v35)
                                        _ = v48
                                        _ = v49
                                    } else {
                                        if (0 <= 76) {
                                            v5 -= v5
                                            var v50 uint = v1
                                            _ = len(v3)
                                            _ = v50
                                        } else {
                                            v1 = 2
                                            switch -2147483648 {
                                            case 128:
                                                if true {
                                                    switch "hello" {
                                                    case "a":
                                                        var v51 int = (fi41 * 1)
                                                        _ = v51
                                                    case "hello":
                                                        v5 -= -647
                                                    case "":
                                                        v35 += (^(v35))
                                                    }
                                                    for fi52 := 0; fi52 < 6; fi52++ {
                                                        switch "foo bar" {
                                                        case "hello":
                                                            _ = len(v3)
                                                        case "foo bar":
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                v1 += 18446744073709551615
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    v34 = (v1 | 581)
                                                                case "":
                                                                    var v53 float64 = (1.5 * 3.14159)
                                                                    _ = v53
                                                                case "a":
                                                                    var v54 uint64 = 2
                                                                    _ = v54
                                                                }
                                                                switch (fi52 * 0) {
                                                                case 361:
                                                                    var v55 uint64 = 42
                                                                    _ = v55
                                                                default:
                                                                    v56 := GenS2{F0: uint64(v35), F1: (18446744073709551615 > 534)}
                                                                    _ = v56
                                                                }
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = (fi52 / ((int(v35)) | 1))
                                                            }
                                                        case "a":
                                                            v57 := GenS4{F0: 1, F1: 0}
                                                            _ = v57
                                                        }
                                                        v34 -= 449
                                                        switch fi52 {
                                                        case 32767:
                                                            v35 -= v35
                                                        default:
                                                            switch 2147483647 {
                                                            case 9223372036854775807:
                                                                if ((true && ((fi52 * (^(fi41))) <= (fi52 / ((fi52) | 1)))) || (!true)) {
                                                                    v58 := GenS5{F0: ((!(-102 == 127)) && true), F1: (((true || (17 <= 100)) || (42 >= (v4 + (fi41 << (((v1 * uint(v35))) & 63))))) || (!(255 == 255))), F2: 2147483648}
                                                                    var v59 float32 = 1e10
                                                                    if true {
                                                                        v4 += fi41
                                                                    }
                                                                    if (2 != v1) {
                                                                        switch "" {
                                                                        case "":
                                                                            var v60 uint8 = 128
                                                                            _ = v60
                                                                        case "foo bar":
                                                                            v61 := GenS4{F0: 42, F1: (v5 / (((v5 >> (((v34 * (^(v34)))) & 63))) | 1))}
                                                                            _ = v61
                                                                        }
                                                                    } else {
                                                                        if true {
                                                                            var v62 int64 = 9223372036854775807
                                                                            v35 += (v35 % ((v35) | 1))
                                                                            _ = v62
                                                                        } else {
                                                                            var v63 int64 = 0
                                                                            var v64 uint = 0
                                                                            v65 := GenS2{F0: 381, F1: (v4 <= fi41)}
                                                                            _ = v63
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                        if (false && false) {
                                                                            var v66 float64 = (1e10 + (3.14159 - float64(1.5)))
                                                                            _ = v66
                                                                        }
                                                                        v67 := GenS1{F0: 255, F1: (-(fi41))}
                                                                        v5 += v5
                                                                        _ = v67
                                                                    }
                                                                    var v68 uint8 = 255
                                                                    if (true && (-1 >= 23)) {
                                                                        var v69 uint8 = v68
                                                                        switch 642 {
                                                                        case 4294967295:
                                                                            _ = len(v3)
                                                                        case 2147483647:
                                                                            var v70 int64 = 1
                                                                            _ = v70
                                                                        case -2147483648:
                                                                            var v71 float64 = (3.14159 * ((1.5 + -2.25) / ((1.5 + 0.1) / float64(1e10))))
                                                                            _ = v71
                                                                        default:
                                                                            var v72 string = "a"
                                                                            _ = v72
                                                                        }
                                                                        _ = v69
                                                                    } else {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = v4
                                                                        }
                                                                        if (!(!(255 < 65535))) {
                                                                            var v73 int16 = v5
                                                                            _ = v73
                                                                        }
                                                                    }
                                                                    if true {
                                                                        v2[fi41] = -627
                                                                        if false {
                                                                            fmt.Println(v68, v34, v4)
                                                                            var v74 int64 = -9223372036854775808
                                                                            _ = v74
                                                                        }
                                                                        if (1e10 == 3.14159) {
                                                                            if true {
                                                                                v5 = (-(v5))
                                                                            }
                                                                            v3 = append(v3, (-(fi52)))
                                                                        } else {
                                                                            v3 = append(v3, fi52)
                                                                        }
                                                                    } else {
                                                                        var v75 uint64 = 2
                                                                        v5 -= v5
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        fmt.Println(fi41)
                                                                        _ = v75
                                                                    }
                                                                    switch (fi41 + 32768) {
                                                                    case -538:
                                                                        v76 := GenS1{F0: 2, F1: -2147483648}
                                                                        _ = v76
                                                                    default:
                                                                        v2[(v4 + (-(fi52)))] = 42
                                                                    }
                                                                    _ = v58
                                                                    _ = v59
                                                                    _ = v68
                                                                }
                                                            default:
                                                                v77 := GenS4{F0: 42, F1: (v5 & v5)}
                                                                v78 := GenS6{F0: (v4 >> ((v34) & 63)), F1: 0, F2: 1}
                                                                v34 = (v34 - 255)
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                            v79 := GenS1{F0: 65535, F1: (fi52 >> ((0) & 63))}
                                                            if (100 == 22) {
                                                                v3 = append(v3, fi41)
                                                            }
                                                            _ = v79
                                                        }
                                                        _ = fi52
                                                    }
                                                    switch ("hello" + ("hello" + "\n\t")) {
                                                    case "foo bar":
                                                        var v80 uint64 = 1
                                                        _ = v80
                                                    case "\n\t":
                                                        v81 := GenS7{F0: ("foo bar" + "hello"), F1: 124}
                                                        _ = v81
                                                    }
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                            default:
                                                v1 += 757
                                            }
                                        }
                                        if (v35 != (v35 / ((2) | 1))) {
                                            var v82 bool = false
                                            _ = v82
                                        } else {
                                            v4 = 4294967295
                                            var v83 int8 = -122
                                            _ = v83
                                        }
                                        v35 -= (v35 & (v35 >> (((v34 + v34)) & 63)))
                                    }
                                } else {
                                    v34 -= (v34 / (((v34 - (v1 << (((v34 & 42)) & 63)))) | 1))
                                }
                                v1 = 18446744073709551615
                            } else {
                                switch (-(fi41)) {
                                case 500:
                                    v3 = append(v3, -6)
                                case 1:
                                    switch 4294967295 {
                                    case 128:
                                        switch ("hello" + ("\n\t" + "hello")) {
                                        case "\n\t":
                                            var v84 uint16 = 2
                                            _ = v84
                                        case "hello":
                                            v1 += (^(v34))
                                        }
                                    case 195:
                                        v5 = v5
                                    case 0:
                                        for fi85 := 0; fi85 < 7; fi85++ {
                                            if (true && (!true)) {
                                                var v86 uint = v1
                                                fmt.Println(v35, v34)
                                                _ = v86
                                            }
                                            v87 := GenS7{F0: "hello", F1: 241}
                                            _ = len(v2)
                                            switch ("" + "hello") {
                                            case "":
                                                switch 1 {
                                                case 51:
                                                    fmt.Println(fi85, v34, v34)
                                                default:
                                                    v88 := GenS8{F0: 143}
                                                    _ = v88
                                                }
                                            case "foo bar":
                                                v89 := GenS9{F0: 42, F1: 100, F2: 42}
                                                _ = v89
                                            }
                                            _ = fi85
                                            _ = v87
                                        }
                                    default:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                case 9223372036854775807:
                                    var v90 int = (v4 - fi41)
                                    _ = v90
                                default:
                                    if (42 != v1) {
                                        if ((!true) || true) {
                                            var v91 int = -9223372036854775808
                                            _ = v91
                                        }
                                        for fi92 := 0; fi92 < 13; fi92++ {
                                            if (-2147483648 < 2147483647) {
                                                if ((0.1 - -2.25) > (float32(float64((float32((((float64((((1.5 * float32(1.5)) * ((-2.25 * (3.14159 + 1e10)) / ((3.14159 / (((((-2.25 * 0.0) - -2.25) / 3.14159) + -2.25) + 1e10)) * 1.5))) / 1e10)) / float64((((0.1 / float32((1e10 - (0.0 + (float64(((0.0 / 3.14159) + float32(((3.14159 - 1.5) + (0.0 - 0.0))))) - 1e10))))) - 1.5) / 0.1))) / 1.0) * 0.0)) + 0.0))) + (-2.25 + 0.1))) {
                                                    var v93 int = fi92
                                                    var v94 int = (v4 / (((-(v4))) | 1))
                                                    _ = v93
                                                    _ = v94
                                                } else {
                                                    if true {
                                                        if true {
                                                            v4 = fi41
                                                            var v95 float64 = (-2.25 + 0.1)
                                                            v96 := GenS10{F0: (fi41 | v4)}
                                                            _ = v95
                                                            _ = v96
                                                        }
                                                        if (!(false && (((4294967295 <= uint64(v1)) && (68 == 127)) && (2147483647 == 42)))) {
                                                            if (!(true || false)) {
                                                                v35 = (v35 | 428)
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                v5 = 0
                                                                v34 = v1
                                                                if ((!(0 >= 255)) || false) {
                                                                    var v97 int16 = (v5 + genFn0(597, (v5 ^ v5)))
                                                                    v98 := GenS6{F0: -1, F1: 18446744073709551615, F2: -1}
                                                                    fmt.Println(v4, v5)
                                                                    _ = v97
                                                                    _ = v98
                                                                } else {
                                                                    if false {
                                                                        fmt.Println(v34)
                                                                        v34 = v1
                                                                        if (!(v35 == (v35 << ((v1) & 63)))) {
                                                                            if true {
                                                                                if false {
                                                                                    fmt.Println(v34, v34)
                                                                                    var v99 uint32 = (v35 ^ (v35 >> ((4294967295) & 63)))
                                                                                    v34 += v1
                                                                                    _ = v99
                                                                                } else {
                                                                                    fmt.Println(v4, fi41, v4)
                                                                                    var v100 uint32 = 42
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        var v101 int32 = 0
                                                                                        _ = v101
                                                                                    }
                                                                                    _ = v100
                                                                                }
                                                                            } else {
                                                                                var v102 uint64 = 0
                                                                                v34 = (^(v34))
                                                                                v35 -= v35
                                                                                _ = v102
                                                                            }
                                                                            v103 := GenS10{F0: fi41}
                                                                            var v104 uint64 = 474
                                                                            if ((v5 * v5) > 1) {
                                                                                v105 := GenS6{F0: 779, F1: v104, F2: -1}
                                                                                if false {
                                                                                    if (32767 < fi92) {
                                                                                        var v106 uint32 = uint32(v1)
                                                                                        v107 := GenS3{F0: 42, F1: fi41}
                                                                                        _ = v106
                                                                                        _ = v107
                                                                                    } else {
                                                                                        var v108 float32 = -2.25
                                                                                        _ = v108
                                                                                    }
                                                                                } else {
                                                                                    v35 = (^(v35))
                                                                                    switch (fi92 / ((127) | 1)) {
                                                                                    case 42:
                                                                                        fmt.Println(fi92, v34)
                                                                                    case 1:
                                                                                        if (true && true) {
                                                                                            fmt.Println(v1, v35)
                                                                                        } else {
                                                                                            var v109 float32 = (((((((1e10 / 1e10) / 1.0) / (1.5 / 1.0)) / (((0.1 + 1e10) + 1.5) / 1e10)) + 1.5) + 3.14159) + (0.1 + 0.0))
                                                                                            _ = v109
                                                                                        }
                                                                                    case 4294967295:
                                                                                        fmt.Println(v34, fi41)
                                                                                    default:
                                                                                        v5 = 323
                                                                                    }
                                                                                    var v110 string = ("" + "")
                                                                                    _ = v110
                                                                                }
                                                                                _ = v105
                                                                            }
                                                                            var v111 bool = (!false)
                                                                            _ = v103
                                                                            _ = v104
                                                                            _ = v111
                                                                        } else {
                                                                            if true {
                                                                                v112 := GenS11{F0: 1, F1: 134, F2: 42}
                                                                                if (32767 == (-(v5))) {
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        v5 = v5
                                                                                    default:
                                                                                        v1 += v1
                                                                                    }
                                                                                    _ = len(v2)
                                                                                } else {
                                                                                    v113 := GenS1{F0: 594, F1: 2147483647}
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = 659
                                                                                    }
                                                                                    _ = v113
                                                                                }
                                                                                var v114 float32 = 1e10
                                                                                _ = v112
                                                                                _ = v114
                                                                            }
                                                                            var v115 float64 = -2.25
                                                                            _ = v115
                                                                        }
                                                                        if (((((3.14159 * (1.5 * ((3.14159 * -2.25) + ((float64(0.1) - float64(1.5)) * -2.25)))) - 1.5) * 0.0) != 0.0) || (!((983 <= -1) || (1 == 935)))) {
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                fmt.Println(v5)
                                                                            case "":
                                                                                v116 := GenS4{F0: 0, F1: v5}
                                                                                var v117 int16 = genFn0(2, (v5 ^ v5))
                                                                                v118 := GenS7{F0: "a", F1: uint8(v35)}
                                                                                _ = v116
                                                                                _ = v117
                                                                                _ = v118
                                                                            }
                                                                            var v119 float32 = 1e10
                                                                            var v120 string = "foo bar"
                                                                            _ = v119
                                                                            _ = v120
                                                                        } else {
                                                                            v1 -= v34
                                                                            if true {
                                                                                fmt.Println(v35, fi92, v5)
                                                                                var v121 uint = 42
                                                                                _ = v121
                                                                            }
                                                                        }
                                                                        v122 := GenS8{F0: 1}
                                                                        _ = v122
                                                                    } else {
                                                                        var v123 int = (fi41 / ((fi92) | 1))
                                                                        _ = len(v2)
                                                                        _ = v123
                                                                    }
                                                                }
                                                                v124 := GenS12{F0: "a"}
                                                                if (0.0 < (((1.5 - 1.5) / (((3.14159 / 0.1) + 3.14159) - 1.5)) / 1.5)) {
                                                                    switch (fi92 + fi92) {
                                                                    case -1:
                                                                        var v125 int16 = (v5 | v5)
                                                                        _ = v125
                                                                    case 127:
                                                                        var v126 float64 = float64((1.5 + (-2.25 - (((0.0 + 0.0) + 1.5) + 0.0))))
                                                                        _ = v126
                                                                    }
                                                                    v127 := GenS13{F0: -9223372036854775808}
                                                                    v5 = genFn0(954, (v5 + 0))
                                                                    _ = v127
                                                                }
                                                                v1 = 65535
                                                                if len(v3) > 0 {
                                                                    v3[0] = 32768
                                                                }
                                                                _ = v124
                                                            }
                                                            v128 := GenS4{F0: 1, F1: (v5 & v5)}
                                                            var v129 float32 = (-2.25 - (((1.5 * ((1e10 * (float32((0.1 + float64(0.0))) * ((float32((3.14159 + 1e10)) * ((3.14159 - float32(0.0)) + float32(0.1))) + 0.1))) * (0.1 / 3.14159))) + 0.0) * 0.0))
                                                            _ = v128
                                                            _ = v129
                                                        }
                                                        fmt.Println(v5, fi41)
                                                        switch 650 {
                                                        case -1:
                                                            v1 -= 73
                                                        default:
                                                            fmt.Println(fi92, fi92)
                                                            v4 += fi41
                                                            var v130 uint64 = 2
                                                            if false {
                                                                var v131 string = "foo bar"
                                                                _ = v131
                                                            } else {
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    var v132 uint32 = v35
                                                                    _ = v132
                                                                case "\n\t":
                                                                    fmt.Println(v35)
                                                                case "a":
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                default:
                                                                    _ = len(v2)
                                                                }
                                                                var v133 uint = (v1 >> (((v1 - 831)) & 63))
                                                                v134 := GenS8{F0: 0}
                                                                _ = v133
                                                                _ = v134
                                                            }
                                                            v135 := GenS14{F0: (v1 * v34)}
                                                            _ = v130
                                                            _ = v135
                                                        }
                                                        if ((v35 <= (v35 | (^(v35)))) || true) {
                                                            v136 := GenS1{F0: 571, F1: 0}
                                                            _ = v136
                                                        }
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                v35 = v35
                                                switch (v4 | (fi41 + (^(fi41)))) {
                                                case 128:
                                                    if ((!true) || true) {
                                                        _ = len(v3)
                                                    } else {
                                                        var v137 uint32 = (v35 * v35)
                                                        v138 := GenS12{F0: "a"}
                                                        v139 := GenS5{F0: false, F1: true, F2: v137}
                                                        _ = v137
                                                        _ = v138
                                                        _ = v139
                                                    }
                                                }
                                                switch ("a" + "a") {
                                                case "":
                                                    if len(v3) > 0 {
                                                        v3[0] = v4
                                                    }
                                                case "hello":
                                                    var v140 uint64 = 0
                                                    _ = v140
                                                }
                                                switch 1 {
                                                case 17:
                                                    var v141 bool = false
                                                    _ = v141
                                                case -900:
                                                    v142 := GenS15{F0: 226, F1: 255}
                                                    _ = v142
                                                case 128:
                                                    var v143 int64 = -9223372036854775808
                                                    _ = v143
                                                default:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                            }
                                            v144 := GenS16{F0: (v4 ^ 32767), F1: 1, F2: 32767}
                                            _ = fi92
                                            _ = v144
                                        }
                                        var v145 int64 = -1
                                        _ = v145
                                    } else {
                                        for fi146 := 0; fi146 < 5; fi146++ {
                                            var v147 uint = v1
                                            var v148 uint = v147
                                            v1 = 18446744073709551615
                                            v1 -= (v1 ^ (v34 % (((v34 % ((18446744073709551615) | 1))) | 1)))
                                            v149 := GenS17{F0: v4, F1: 1, F2: 1}
                                            _ = fi146
                                            _ = v147
                                            _ = v148
                                            _ = v149
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            v4 = 340
                                        case "hello":
                                            switch "\n\t" {
                                            case "a":
                                                v5 = v5
                                            case "hello":
                                                v5 += (v5 - v5)
                                            }
                                        case "a":
                                            if (!true) {
                                                for fi150 := 0; fi150 < 6; fi150++ {
                                                    if (true || false) {
                                                        var v151 uint32 = (^(v35))
                                                        fmt.Println(v35)
                                                        _ = v151
                                                    } else {
                                                        switch fi41 {
                                                        case 42:
                                                            v35 -= (v35 ^ (v35 << ((v1) & 63)))
                                                        }
                                                    }
                                                    var v152 uint32 = 2
                                                    _ = fi150
                                                    _ = v152
                                                }
                                            } else {
                                                delete(v2, int(v35))
                                                v153 := GenS1{F0: 42, F1: (fi41 ^ (fi41 / ((-327) | 1)))}
                                                for fi154 := 0; fi154 < 3; fi154++ {
                                                    if false {
                                                        switch 128 {
                                                        case 1:
                                                            if true {
                                                                var v155 float32 = 1e10
                                                                if false {
                                                                    v35 += 4294967295
                                                                }
                                                                var v156 uint = 4294967295
                                                                _ = v155
                                                                _ = v156
                                                            } else {
                                                                fmt.Println(v4, v34, v1)
                                                                if (100 > 18446744073709551615) {
                                                                    _ = len(v3)
                                                                    if (genFn0(255, (v5 % (((v5 - -506)) | 1))) > genFn0(1, 0)) {
                                                                        var v157 float32 = -2.25
                                                                        _ = v157
                                                                    } else {
                                                                        var v158 float32 = 0.0
                                                                        _ = v158
                                                                    }
                                                                }
                                                                var v159 bool = true
                                                                _ = v159
                                                            }
                                                        case 155:
                                                            if false {
                                                                var v160 int8 = 0
                                                                var v161 uint32 = uint32(v34)
                                                                var v162 uint32 = (v161 | (v35 | v161))
                                                                _ = v160
                                                                _ = v161
                                                                _ = v162
                                                            } else {
                                                                fmt.Println(v1, v34)
                                                                var v163 int16 = (v5 / (((^(v5))) | 1))
                                                                v3 = append(v3, 4294967295)
                                                                _ = v163
                                                            }
                                                        case 29:
                                                            var v164 string = (("hello" + "a") + "")
                                                            _ = v164
                                                        default:
                                                            v1 += (v1 | 137)
                                                        }
                                                    }
                                                    _ = fi154
                                                }
                                                _ = v153
                                            }
                                        default:
                                            v4 += v4
                                        }
                                    }
                                }
                                fmt.Println(fi41)
                                for fi165 := 0; fi165 < 17; fi165++ {
                                    v1 -= 16
                                    _ = fi165
                                }
                            }
                            var v166 uint64 = 889
                            _ = v166
                        }
                    case "foo bar":
                        switch ("a" + "a") {
                        case "foo bar":
                            var v167 int = 428
                            _ = v167
                        default:
                            v168 := GenS8{F0: 170}
                            _ = v168
                        }
                    }
                    v3 = append(v3, (-(v4)))
                    _ = fi41
                }
                v169 := GenS5{F0: (false || (255 <= 1)), F1: (v5 > v5), F2: (v35 + v35)}
                v170 := GenS1{F0: 353, F1: 644}
                _ = v169
                _ = v170
            }
            v4 = int(v34)
            var v171 int16 = -1
            _ = v171
        } else {
            v35 += 708
        }
    default:
        if (!(false || (!true))) {
            var v172 uint32 = 2147483647
            _ = v172
        } else {
            if false {
                if len(v3) > 0 {
                    v3[0] = (v4 + 103)
                }
            } else {
                if false {
                    v35 += v35
                    var v173 uint32 = v35
                    fmt.Println(v173, v34)
                    _ = v173
                } else {
                    for fi174 := 0; fi174 < 3; fi174++ {
                        v175 := GenS4{F0: 1, F1: (v5 | 0)}
                        fmt.Println(v1, v5)
                        _ = fi174
                        _ = v175
                    }
                }
            }
            v176 := GenS9{F0: 128, F1: 1, F2: 100}
            v1 -= 42
            _ = v176
        }
    }
    for fi177 := 0; fi177 < 5; fi177++ {
        v178 := GenS1{F0: 231, F1: (-(v4))}
        _ = fi177
        _ = v178
    }
    v179 := GenS18{F0: 65535, F1: 128}
    switch "foo bar" {
    case "":
        var v180 uint64 = 42
        _ = v180
    case "\n\t":
        v34 -= (v34 / ((v34) | 1))
    default:
        switch ((("\n\t" + "a") + "\n\t") + "") {
        case "":
            for fi181 := 0; fi181 < 9; fi181++ {
                switch "\n\t" {
                case "foo bar":
                    v182 := GenS3{F0: 186, F1: v4}
                    _ = v182
                case "":
                    v35 = v35
                default:
                    var v183 uint8 = 108
                    _ = v183
                }
                var v184 uint32 = v35
                fmt.Println(v34, v34, v35)
                _ = fi181
                _ = v184
            }
        case "hello":
            v1 -= 2
        default:
            v35 += 2147483648
        }
    }
    switch "a" {
    case "":
        var v185 int8 = 6
        _ = v185
    case "foo bar":
        switch "hello" {
        case "":
            fmt.Println(v5, v1, v1)
        case "\n\t":
            var v186 uint = 4294967295
            _ = v186
        case "a":
            if true {
                for fi187 := 0; fi187 < 12; fi187++ {
                    switch (((("hello" + ("a" + "hello")) + (("foo bar" + "\n\t") + "")) + "a") + "a") {
                    case "a":
                        fmt.Println(v4, fi187)
                    }
                    _ = fi187
                }
                fmt.Println(v1, v5, v34)
                for fi188 := 0; fi188 < 19; fi188++ {
                    v189 := GenS5{F0: ((0 != 1) || false), F1: (0 >= 115), F2: v35}
                    v190 := GenS19{F0: 1}
                    _ = fi188
                    _ = v189
                    _ = v190
                }
            }
        default:
            for fi191 := 0; fi191 < 1; fi191++ {
                v2[42] = (fi191 / ((128) | 1))
                _ = fi191
            }
        }
    case "\n\t":
        v2[(-(v4))] = (v4 << ((778) & 63))
    }
    v34 = v1
    fmt.Println(v1, v1, v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v179
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
