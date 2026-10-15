package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int64; F1 uint32; F2 uint }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 int; F1 uint8; F2 uint32 }
type GenS6 struct { F0 int; F1 uint64; F2 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int16; F1 uint64 }
type GenS9 struct { F0 string; F1 uint64; F2 uint32 }
type GenS10 struct { F0 uint32; F1 uint }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int; F1 uint32; F2 int }
type GenS13 struct { F0 int64; F1 uint64; F2 int8 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int8; F1 uint8; F2 int8 }
type GenS16 struct { F0 uint; F1 uint }
type GenS17 struct { F0 uint32; F1 uint32 }
type GenS18 struct { F0 int8; F1 uint16 }
type GenS19 struct { F0 uint; F1 uint64; F2 int16 }
type GenS20 struct { F0 int64; F1 uint8 }
type GenS21 struct { F0 uint32; F1 uint }
type GenS22 struct { F0 int16 }
type GenS23 struct { F0 uint32; F1 uint; F2 string }
type GenS24 struct { F0 uint }
type GenS25 struct { F0 uint8; F1 uint16 }
type GenS26 struct { F0 int8; F1 int; F2 int8 }
type GenS27 struct { F0 int8; F1 bool }
type GenS28 struct { F0 int64 }
type GenS29 struct { F0 string; F1 bool }
func genFn0(p0 int32, p1 uint8, p2 string) bool {
    var v4 uint = 18446744073709551615
    for fi5 := 0; fi5 < 1; fi5++ {
        p1 = p1
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (true && false)
}
func genFn1(p0 int, p1 string, p2 uint16) bool {
    v6 := GenS1{F0: (((^(p2)) == (p2 / ((65535) | 1))) && (-1 > -128))}
    switch p0 {
    case 42:
        fmt.Println(p2, p2)
    case -1:
        switch (p1 + "foo bar") {
        case "\n\t":
            var v7 int64 = 1
            _ = v7
        case "foo bar":
            v8 := []int{p0, p0}
            _ = v8
        default:
            switch (p0 + (p0 / ((42) | 1))) {
            case -1:
                var v9 string = "hello"
                _ = v9
            }
        }
    case 32767:
        if ((true && true) && false) {
            p2 -= (p2 | (p2 * p2))
            if false {
                v10 := GenS2{F0: 127, F1: p0}
                fmt.Println(p2, p2)
                p2 += (^(p2))
                _ = v10
            } else {
                v11 := GenS1{F0: false}
                _ = v11
            }
        } else {
            fmt.Println(p0, p2)
        }
    }
    if false {
        p2 += (p2 >> ((394) & 63))
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := []int{(fi12 + fi12)}
            var v15 float64 = 0.0
            _ = fi13
            _ = v14
            _ = v15
        }
        v16 := GenS3{F0: 1, F1: 950, F2: 18446744073709551615}
        switch ("hello" + ((("hello" + p1) + p1) + p1)) {
        case "foo bar":
            if ((p2 >= (p2 << ((4294967295) & 63))) || ((2 < 1) || true)) {
                fmt.Println(p2, p0)
            }
        case "hello":
            v17 := GenS3{F0: 1, F1: 65535, F2: 1}
            _ = v17
        }
        _ = fi12
        _ = v16
    }
    if true {
        v18 := GenS3{F0: 9223372036854775807, F1: 42, F2: 1}
        v19 := []int{(p0 - 4294967295), (p0 | p0), (p0 / ((378) | 1))}
        var v20 int8 = -1
        _ = v18
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return false
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 42
    var v3 int16 = 682
    for fi21 := 0; fi21 < 14; fi21++ {
        v22 := []int{}
        fmt.Println(v2, v2, v3)
        if (18446744073709551615 != 0) {
            v3 += (v3 * 42)
            if ((fi21 ^ 0) >= 1) {
                var v23 int16 = v3
                for fi24 := 0; fi24 < 4; fi24++ {
                    fmt.Println(fi21, v23, v23)
                    var v25 int64 = v2
                    var v26 uint32 = 258
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                switch (fi21 >> ((42) & 63)) {
                case 422:
                    for fi27 := 0; fi27 < 7; fi27++ {
                        if false {
                            v3 = (v23 / ((-1) | 1))
                            v28 := GenS3{F0: (v2 / ((1) | 1)), F1: 4294967295, F2: 0}
                            _ = v28
                        }
                        _ = fi27
                    }
                case 32768:
                    if (-104 <= 41) {
                        if (false || (1.5 == (((float32(3.14159) / ((((0.0 / (0.0 + 1.5)) - 1.5) * -2.25) / float32(3.14159))) * -2.25) - (-2.25 * (-2.25 * 1e10))))) {
                            switch (("foo bar" + "hello") + "\n\t") {
                            case "foo bar":
                                v2 = (-(v2))
                            }
                            for fi29 := 0; fi29 < 15; fi29++ {
                                _ = len(v22)
                                v30 := GenS4{F0: v2, F1: (v3 | -7)}
                                _ = fi29
                                _ = v30
                            }
                            var v31 float64 = (-2.25 / (((-2.25 * (((1e10 / -2.25) * ((3.14159 - 3.14159) + float64(0.0))) - 0.1)) / 0.1) / 1.0))
                            _ = v31
                        }
                        v2 += -160
                    }
                case 0:
                    if true {
                        var v32 int8 = 42
                        v23 = (v23 * v23)
                        for fi33 := 0; fi33 < 18; fi33++ {
                            var v34 uint32 = 2147483648
                            if (!(float64(((((((3.14159 / 1e10) + 1e10) / (float32(-2.25) / 3.14159)) / (((((1.5 - (((1.5 * (0.1 / 0.1)) * 0.1) * 0.1)) / 1.0) - (0.1 * float32(((3.14159 - (0.0 - ((1.5 - (float64(-2.25) * (-2.25 - 0.1))) - (0.0 - (-2.25 - ((1e10 + (0.0 - (1.5 + 1e10))) + (0.1 - float64(float32(0.1))))))))) * ((-2.25 * 3.14159) * 0.0))))) * 0.1) - 1e10)) + -2.25) / ((-2.25 / (-2.25 - 3.14159)) / (1e10 - (0.0 / (1.5 - ((1.5 * float32((1.5 * ((((((1e10 + float64(0.1)) * 3.14159) * 1.5) - -2.25) - (float64(0.1) - (float64(0.0) * (0.1 * 1.5)))) * 0.0)))) + 0.1))))))) <= (1.5 - 3.14159))) {
                                if (false || (42 > 42)) {
                                    fmt.Println(v2)
                                } else {
                                    if ((fi21 >> ((34) & 63)) < (fi33 - -639)) {
                                        v2 -= v2
                                        var v35 string = ("hello" + "\n\t")
                                        fmt.Println(v34, fi33)
                                        v2 = v2
                                        v32 = v32
                                        v32 -= v32
                                        switch fi33 {
                                        case 0:
                                            var v36 float32 = -2.25
                                            _ = v36
                                        case 4294967295:
                                            fmt.Println(fi33, v34, v23)
                                        }
                                        v32 -= v32
                                        if true {
                                            v2 = (-(v2))
                                        } else {
                                            var v37 bool = (!((fi33 >> ((0) & 63)) == 568))
                                            _ = v37
                                        }
                                        var v38 uint16 = 2
                                        _ = v35
                                        _ = v38
                                    } else {
                                        v34 += (v34 >> ((213) & 63))
                                        switch "a" {
                                        case "a":
                                            delete(v1, int(v32))
                                        case "\n\t":
                                            switch fi33 {
                                            case -421:
                                                var v39 int = -694
                                                _ = v39
                                            case 127:
                                                var v40 uint64 = 42
                                                _ = v40
                                            case 0:
                                                v41 := v1[(fi33 - 128)]
                                                _ = v41
                                            }
                                        }
                                        v42 := GenS5{F0: fi21, F1: 127, F2: 372}
                                        _ = v42
                                    }
                                    v22 = append(v22, (fi21 % (((fi33 | fi33)) | 1)))
                                    var v43 bool = (128 == -2147483648)
                                    _ = v43
                                }
                                var v44 uint64 = 18446744073709551615
                                v32 -= 42
                                if true {
                                    _ = len(v1)
                                    if ((!(true && false)) && true) {
                                        v3 = (v23 - (^(v3)))
                                    } else {
                                        if len(v22) > 0 {
                                            v22[0] = (fi33 | fi33)
                                        }
                                    }
                                } else {
                                    v3 += v23
                                    if (!true) {
                                        if false {
                                            var v45 bool = ((!((true && false) || true)) && false)
                                            _ = v45
                                        } else {
                                            switch (fi21 % ((-742) | 1)) {
                                            case -9223372036854775808:
                                                fmt.Println(v3, v2, v32)
                                            case 218:
                                                switch fi21 {
                                                case -9223372036854775808:
                                                    var v46 int64 = v2
                                                    v47 := GenS6{F0: 111, F1: v44, F2: 0}
                                                    if true {
                                                        switch "hello" {
                                                        case "a":
                                                            switch (("hello" + "foo bar") + "\n\t") {
                                                            case "\n\t":
                                                                var v48 bool = (fi33 == 32767)
                                                                _ = v48
                                                            case "a":
                                                                if ("foo bar" == "foo bar") {
                                                                    v46 -= (v2 >> ((599) & 63))
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        if false {
                                                                            if true {
                                                                                fmt.Println(v44)
                                                                                fmt.Println(v23, v32, v44)
                                                                                if len(v22) > 0 {
                                                                                    v22[0] = fi21
                                                                                }
                                                                            }
                                                                            v3 = 1
                                                                        } else {
                                                                            v32 = (v32 >> ((412) & 63))
                                                                            v49 := GenS6{F0: fi33, F1: v44, F2: 2}
                                                                            switch "a" {
                                                                            case "a":
                                                                                var v50 int = (fi21 << ((0) & 63))
                                                                                _ = v50
                                                                            case "hello":
                                                                                if ("a" == "") {
                                                                                    v3 -= (v3 >> ((65535) & 63))
                                                                                    switch fi21 {
                                                                                    case 32767:
                                                                                        if ((false && true) || false) {
                                                                                            if true {
                                                                                                if true {
                                                                                                    if (fi33 < fi21) {
                                                                                                        fmt.Println(fi33, v46, v46)
                                                                                                    }
                                                                                                    switch -356 {
                                                                                                    case -67:
                                                                                                        if len(v22) >= 2 {
                                                                                                            v22 = v22[1:len(v22)]
                                                                                                        }
                                                                                                    case 114:
                                                                                                        var v51 int16 = v3
                                                                                                        _ = v51
                                                                                                    case -1:
                                                                                                        var v52 float32 = ((((3.14159 - float32((1e10 / 1.0))) * 3.14159) * float32((0.0 * (0.0 / 0.1)))) + (((0.0 / 1e10) * 1.5) * 0.0))
                                                                                                        _ = v52
                                                                                                    }
                                                                                                    if len(v22) >= 2 {
                                                                                                        v22 = v22[1:len(v22)]
                                                                                                    }
                                                                                                } else {
                                                                                                    var v53 int64 = (v46 | (v2 & (v46 * (v2 << ((643) & 63)))))
                                                                                                    switch -801 {
                                                                                                    case 9223372036854775807:
                                                                                                        fmt.Println(fi21, v34, v3)
                                                                                                    case 4294967295:
                                                                                                        v54 := GenS7{F0: v32}
                                                                                                        _ = v54
                                                                                                    }
                                                                                                    _ = v53
                                                                                                }
                                                                                                fmt.Println(v3, v3, v46)
                                                                                                var v55 float64 = 1.5
                                                                                                _ = v55
                                                                                            }
                                                                                        } else {
                                                                                            v56 := GenS8{F0: 42, F1: v44}
                                                                                            _ = v56
                                                                                        }
                                                                                    }
                                                                                }
                                                                            default:
                                                                                v44 -= v44
                                                                                var v57 int32 = -465
                                                                                _ = v57
                                                                            }
                                                                            _ = v49
                                                                        }
                                                                    case "":
                                                                        if (!false) {
                                                                            v3 = (v3 / ((-1) | 1))
                                                                        } else {
                                                                            switch (fi33 % ((-2147483648) | 1)) {
                                                                            case 2147483647:
                                                                                v32 = v32
                                                                            case 32767:
                                                                                if (!true) {
                                                                                    v58 := GenS4{F0: (v2 % (((v46 % (((v46 % (((v2 % ((v2) | 1))) | 1))) | 1))) | 1)), F1: int16(v34)}
                                                                                    switch fi21 {
                                                                                    case -22:
                                                                                        if false {
                                                                                            v59 := GenS5{F0: 916, F1: 29, F2: (^(v34))}
                                                                                            switch "\n\t" {
                                                                                            case "hello":
                                                                                                v60 := GenS5{F0: fi21, F1: 88, F2: v34}
                                                                                                _ = v60
                                                                                            default:
                                                                                                if false {
                                                                                                    v23 = v3
                                                                                                    if (65535 <= 255) {
                                                                                                        v61 := GenS9{F0: ("a" + "a"), F1: (v44 & (v44 | 4294967295)), F2: 0}
                                                                                                        _ = v61
                                                                                                    }
                                                                                                    var v62 int32 = 129
                                                                                                    _ = v62
                                                                                                } else {
                                                                                                    v63 := GenS5{F0: (fi33 & (fi33 ^ 0)), F1: 255, F2: v34}
                                                                                                    if (true && ((^(v32)) >= v32)) {
                                                                                                        v32 = (v32 - (^(v32)))
                                                                                                        if true {
                                                                                                            v64 := GenS5{F0: fi21, F1: 0, F2: (v34 >> ((65535) & 63))}
                                                                                                            switch (fi33 + (fi33 % ((fi33) | 1))) {
                                                                                                            case 9223372036854775807:
                                                                                                                fmt.Println(v23, v44, v32)
                                                                                                                v2 = v2
                                                                                                                v32 = (^(v32))
                                                                                                            case -861:
                                                                                                                if (4294967295 < fi21) {
                                                                                                                    switch 127 {
                                                                                                                    case 127:
                                                                                                                        var v65 bool = ((!(!(true || (798 != -2147483648)))) && true)
                                                                                                                        _ = v65
                                                                                                                    case 2147483647:
                                                                                                                        v46 = v46
                                                                                                                    case 32768:
                                                                                                                        v22 = append(v22, (-(fi33)))
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    if (((!true) || true) && true) {
                                                                                                                        var v66 float64 = 0.0
                                                                                                                        if len(v22) > 0 {
                                                                                                                            v22[0] = 791
                                                                                                                        }
                                                                                                                        var v67 uint64 = v44
                                                                                                                        fmt.Println(v32, v66, v23)
                                                                                                                        v68 := GenS10{F0: v34, F1: 2}
                                                                                                                        _ = v66
                                                                                                                        _ = v67
                                                                                                                        _ = v68
                                                                                                                    }
                                                                                                                    v2 = v46
                                                                                                                    v69 := GenS9{F0: ("a" + "hello"), F1: (v44 + v44), F2: v34}
                                                                                                                    v34 = (^(v34))
                                                                                                                    _ = v69
                                                                                                                }
                                                                                                            }
                                                                                                            var v70 float32 = (-2.25 - 1e10)
                                                                                                            v71 := GenS11{F0: 65535}
                                                                                                            var v72 uint = 2
                                                                                                            if false {
                                                                                                                fmt.Println(v46, v32)
                                                                                                                delete(v1, (fi33 | 498))
                                                                                                                var v73 float64 = 1e10
                                                                                                                _ = v73
                                                                                                            } else {
                                                                                                                delete(v1, fi33)
                                                                                                                fmt.Println(v3)
                                                                                                                v74 := GenS12{F0: int(v34), F1: (^(v34)), F2: (fi33 % ((-1) | 1))}
                                                                                                                _ = v74
                                                                                                            }
                                                                                                            switch "\n\t" {
                                                                                                            case "hello":
                                                                                                                v3 -= v3
                                                                                                            case "\n\t":
                                                                                                                if true {
                                                                                                                    v44 = 96
                                                                                                                    if false {
                                                                                                                        if true {
                                                                                                                            v75 := GenS8{F0: v3, F1: (v44 >> (((v72 & v72)) & 63))}
                                                                                                                            _ = v75
                                                                                                                        }
                                                                                                                        v70 = v70
                                                                                                                        var v76 uint8 = 1
                                                                                                                        _ = v76
                                                                                                                    } else {
                                                                                                                        var v77 float64 = 0.0
                                                                                                                        switch "\n\t" {
                                                                                                                        case "a":
                                                                                                                            var v78 int = fi33
                                                                                                                            _ = v78
                                                                                                                        }
                                                                                                                        v46 = -1
                                                                                                                        _ = v77
                                                                                                                    }
                                                                                                                    v46 -= v2
                                                                                                                    var v79 float32 = (v70 / v70)
                                                                                                                    _ = v79
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        _ = len(v1)
                                                                                                                        var v80 int = (-(fi33))
                                                                                                                        _ = v80
                                                                                                                    } else {
                                                                                                                        var v81 uint8 = 127
                                                                                                                        switch ("foo bar" + "") {
                                                                                                                        case "a":
                                                                                                                            delete(v1, 0)
                                                                                                                        case "":
                                                                                                                            v82 := GenS10{F0: (^(v34)), F1: (v72 | v72)}
                                                                                                                            _ = v82
                                                                                                                        case "\n\t":
                                                                                                                            if false {
                                                                                                                                var v83 string = "hello"
                                                                                                                                _ = v83
                                                                                                                            } else {
                                                                                                                                if ((false && (65535 >= 65535)) || (!false)) {
                                                                                                                                    v84 := GenS13{F0: (v2 * (v2 & (v2 ^ (v2 * -9223372036854775808)))), F1: 927, F2: (^(v32))}
                                                                                                                                    _ = v84
                                                                                                                                } else {
                                                                                                                                    v81 += (v81 + (v81 + (^(v81))))
                                                                                                                                    v85 := GenS14{F0: 397}
                                                                                                                                    _ = v85
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                        _ = len(v1)
                                                                                                                        var v86 int64 = (v2 + (v46 << (((^(v72))) & 63)))
                                                                                                                        if true {
                                                                                                                            if (!true) {
                                                                                                                                var v87 uint8 = v81
                                                                                                                                fmt.Println(v81, v81)
                                                                                                                                if true {
                                                                                                                                    v23 = v3
                                                                                                                                } else {
                                                                                                                                    switch ("a" + ("hello" + "hello")) {
                                                                                                                                    case "foo bar":
                                                                                                                                        v46 -= (-(v2))
                                                                                                                                    default:
                                                                                                                                        var v88 int16 = v23
                                                                                                                                        _ = v88
                                                                                                                                    }
                                                                                                                                    if false {
                                                                                                                                        var v89 int16 = (v23 ^ (v3 + (v3 ^ int16(fi33))))
                                                                                                                                        _ = v89
                                                                                                                                    }
                                                                                                                                    v90 := v1[fi21]
                                                                                                                                    _ = len(v1)
                                                                                                                                    _ = v90
                                                                                                                                }
                                                                                                                                _ = v87
                                                                                                                            }
                                                                                                                            _ = len(v22)
                                                                                                                        } else {
                                                                                                                            fmt.Println(v86, v3, fi21)
                                                                                                                        }
                                                                                                                        _ = v81
                                                                                                                        _ = v86
                                                                                                                    }
                                                                                                                    var v91 uint8 = 0
                                                                                                                    _ = v91
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v64
                                                                                                            _ = v70
                                                                                                            _ = v71
                                                                                                            _ = v72
                                                                                                        }
                                                                                                        var v92 bool = true
                                                                                                        _ = v92
                                                                                                    } else {
                                                                                                        _ = len(v1)
                                                                                                        v93 := GenS14{F0: 2}
                                                                                                        _ = v93
                                                                                                    }
                                                                                                    _ = v63
                                                                                                }
                                                                                            }
                                                                                            _ = v59
                                                                                        }
                                                                                    case -9223372036854775808:
                                                                                        fmt.Println(v2)
                                                                                    case 32767:
                                                                                        var v94 int16 = (v23 | v3)
                                                                                        _ = v94
                                                                                    default:
                                                                                        var v95 int = fi21
                                                                                        _ = v95
                                                                                    }
                                                                                    _ = v58
                                                                                } else {
                                                                                    if (42 == (v46 >> ((18446744073709551615) & 63))) {
                                                                                        v96 := GenS15{F0: (v32 << ((4294967295) & 63)), F1: 127, F2: v32}
                                                                                        var v97 uint64 = 2
                                                                                        _ = v96
                                                                                        _ = v97
                                                                                    }
                                                                                    var v98 uint64 = (v44 >> ((255) & 63))
                                                                                    _ = v98
                                                                                }
                                                                            default:
                                                                                if true {
                                                                                    if false {
                                                                                        if len(v22) >= 2 {
                                                                                            v22 = v22[1:len(v22)]
                                                                                        }
                                                                                        v99 := v1[(fi33 % (((fi33 / ((fi33) | 1))) | 1))]
                                                                                        _ = v99
                                                                                    }
                                                                                    _ = len(v22)
                                                                                    v46 = v46
                                                                                } else {
                                                                                    var v100 float64 = 1e10
                                                                                    if false {
                                                                                        fmt.Println(v32, v34)
                                                                                        if len(v22) > 0 {
                                                                                            v22[0] = (fi21 & 209)
                                                                                        }
                                                                                        switch 9223372036854775807 {
                                                                                        case 883:
                                                                                            if true {
                                                                                                switch (("" + "foo bar") + ("" + "\n\t")) {
                                                                                                case "foo bar":
                                                                                                    if ((3.14159 + (1e10 / float32(v100))) == 0.0) {
                                                                                                        var v101 int32 = 0
                                                                                                        delete(v1, (fi33 + -995))
                                                                                                        _ = v101
                                                                                                    } else {
                                                                                                        delete(v1, fi21)
                                                                                                        v3 -= v3
                                                                                                        v102 := GenS9{F0: ("a" + "hello"), F1: v44, F2: 65535}
                                                                                                        _ = v102
                                                                                                    }
                                                                                                case "a":
                                                                                                    var v103 uint16 = 42
                                                                                                    _ = v103
                                                                                                }
                                                                                                fmt.Println(v34, v23)
                                                                                            }
                                                                                        default:
                                                                                            v104 := GenS7{F0: v32}
                                                                                            _ = v104
                                                                                        }
                                                                                        switch "hello" {
                                                                                        case "\n\t":
                                                                                            v2 -= v2
                                                                                        case "foo bar":
                                                                                            var v105 int16 = v23
                                                                                            _ = v105
                                                                                        case "hello":
                                                                                            if ((v34 == 1) && false) {
                                                                                                var v106 uint16 = 664
                                                                                                if ((v2 % (((v46 * (v46 ^ 42))) | 1)) == 42) {
                                                                                                    if true {
                                                                                                        var v107 uint16 = 1
                                                                                                        var v108 uint16 = (^(v107))
                                                                                                        v109 := GenS16{F0: 2, F1: 2}
                                                                                                        v110 := v1[(fi21 & fi21)]
                                                                                                        _ = v107
                                                                                                        _ = v108
                                                                                                        _ = v109
                                                                                                        _ = v110
                                                                                                    }
                                                                                                    delete(v1, (fi21 << ((255) & 63)))
                                                                                                    v111 := GenS4{F0: (v2 % (((v2 + v46)) | 1)), F1: 32767}
                                                                                                    _ = v111
                                                                                                }
                                                                                                var v112 int8 = 1
                                                                                                var v113 bool = true
                                                                                                _ = v106
                                                                                                _ = v112
                                                                                                _ = v113
                                                                                            } else {
                                                                                                var v114 int = 1
                                                                                                if len(v22) > 0 {
                                                                                                    v22[0] = fi21
                                                                                                }
                                                                                                var v115 string = "hello"
                                                                                                _ = v114
                                                                                                _ = v115
                                                                                            }
                                                                                        default:
                                                                                            v116 := GenS17{F0: (v34 >> ((646) & 63)), F1: v34}
                                                                                            _ = v116
                                                                                        }
                                                                                    }
                                                                                    _ = v100
                                                                                }
                                                                            }
                                                                            var v117 uint = 42
                                                                            v2 = v46
                                                                            fmt.Println(v46)
                                                                            _ = v117
                                                                        }
                                                                    default:
                                                                        var v118 int = (fi33 | 32768)
                                                                        _ = v118
                                                                    }
                                                                    switch fi21 {
                                                                    case 32768:
                                                                        var v119 uint16 = 2
                                                                        _ = v119
                                                                    case 144:
                                                                        v32 = (v32 / ((v32) | 1))
                                                                    default:
                                                                        if ((v46 & (v2 + (v2 + (v46 | (^(v2)))))) <= 1) {
                                                                            v3 += 32767
                                                                            v120 := GenS2{F0: 2, F1: 1}
                                                                            _ = v120
                                                                        } else {
                                                                            fmt.Println(v2)
                                                                            v3 = (-(v3))
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v121 int8 = (v32 & -1)
                                                                    _ = v121
                                                                }
                                                            case "hello":
                                                                fmt.Println(v32)
                                                            }
                                                        default:
                                                            switch fi33 {
                                                            case -19:
                                                                var v122 uint16 = 768
                                                                _ = v122
                                                            default:
                                                                _ = len(v1)
                                                            }
                                                        }
                                                    }
                                                    _ = v46
                                                    _ = v47
                                                case 128:
                                                    fmt.Println(fi33)
                                                case 9223372036854775807:
                                                    fmt.Println(v34, v2)
                                                    v34 += (v34 + (^(v34)))
                                                default:
                                                    if false {
                                                        var v123 uint = 796
                                                        _ = len(v22)
                                                        if (42 < 127) {
                                                            if false {
                                                                v23 -= v23
                                                                v23 = int16(v44)
                                                                fmt.Println(v3)
                                                                fmt.Println(v123, v2)
                                                                v3 += (v23 + 42)
                                                                v124 := GenS8{F0: (v3 / (((v23 + v23)) | 1)), F1: (^(v44))}
                                                                _ = v124
                                                            }
                                                        }
                                                        _ = v123
                                                    }
                                                }
                                            case 42:
                                                var v125 int8 = v32
                                                if true {
                                                    v1[int(v2)] = int(v34)
                                                    if (v2 >= v2) {
                                                        v2 = (v2 >> ((2) & 63))
                                                        v22 = append(v22, fi33)
                                                    }
                                                    v23 -= 42
                                                }
                                                _ = v125
                                            }
                                            if true {
                                                var v126 uint64 = v44
                                                var v127 string = ""
                                                var v128 int8 = (v32 << ((18446744073709551615) & 63))
                                                _ = v126
                                                _ = v127
                                                _ = v128
                                            }
                                            if false {
                                                v32 = v32
                                                if (42 >= v3) {
                                                    if len(v22) >= 2 {
                                                        v22 = v22[1:len(v22)]
                                                    }
                                                    _ = len(v22)
                                                } else {
                                                    switch (fi21 & fi21) {
                                                    case 2147483647:
                                                        if (!false) {
                                                            var v129 int = (^(fi33))
                                                            _ = v129
                                                        } else {
                                                            var v130 uint = 4294967295
                                                            _ = v130
                                                        }
                                                    case -14:
                                                        if (true && (!((false || ((v3 <= int16(v44)) && (true || (!true)))) || true))) {
                                                            if len(v22) > 0 {
                                                                v22[0] = (fi21 | fi33)
                                                            }
                                                            if true {
                                                                switch fi33 {
                                                                case -650:
                                                                    v131 := GenS2{F0: 244, F1: 1}
                                                                    _ = v131
                                                                }
                                                            }
                                                        } else {
                                                            var v132 uint16 = 679
                                                            var v133 int8 = 1
                                                            _ = v132
                                                            _ = v133
                                                        }
                                                    case 219:
                                                        v2 -= (v2 / ((v2) | 1))
                                                    }
                                                    var v134 bool = true
                                                    _ = v134
                                                }
                                                if ((195 != 22) || (0 <= (v44 & v44))) {
                                                    var v135 int16 = v3
                                                    var v136 uint = 42
                                                    _ = len(v1)
                                                    _ = v135
                                                    _ = v136
                                                } else {
                                                    v2 -= (v2 >> ((523) & 63))
                                                    var v137 uint8 = uint8(v34)
                                                    _ = v137
                                                }
                                            }
                                            v1[fi33] = fi33
                                            if true {
                                                switch "" {
                                                case "a":
                                                    if false {
                                                        var v138 uint16 = 1
                                                        v22 = append(v22, -9223372036854775808)
                                                        _ = v138
                                                    }
                                                case "hello":
                                                    fmt.Println(fi33)
                                                }
                                                if false {
                                                    v34 -= 765
                                                    fmt.Println(v2)
                                                    if true {
                                                        var v139 int32 = -6
                                                        var v140 int32 = v139
                                                        v141 := GenS4{F0: 9223372036854775807, F1: 0}
                                                        _ = v139
                                                        _ = v140
                                                        _ = v141
                                                    }
                                                }
                                            } else {
                                                v142 := GenS15{F0: v32, F1: 42, F2: v32}
                                                fmt.Println(v32, v32)
                                                _ = v142
                                            }
                                        }
                                    } else {
                                        fmt.Println(fi33, v32)
                                        delete(v1, int(v34))
                                        var v143 int64 = -1
                                        _ = v143
                                    }
                                }
                                switch "\n\t" {
                                case "":
                                    v34 -= (^(v34))
                                }
                                _ = v44
                            }
                            _ = fi33
                            _ = v34
                        }
                        _ = v32
                    }
                }
                _ = v23
            } else {
                var v144 uint8 = 1
                var v145 uint64 = 42
                v3 = (v3 * -1)
                _ = v144
                _ = v145
            }
        }
        _ = fi21
        _ = v22
    }
    if false {
        var v146 float32 = 1.5
        for fi147 := 0; fi147 < 7; fi147++ {
            v148 := []int{fi147, fi147, 0, 127}
            switch fi147 {
            case -361:
                v149 := GenS12{F0: 128, F1: 1, F2: (^(fi147))}
                _ = v149
            case -9223372036854775808:
                v150 := GenS18{F0: -1, F1: 2}
                _ = v150
            }
            v146 += v146
            _ = fi147
            _ = v148
        }
        v151 := GenS15{F0: 0, F1: 0, F2: -29}
        _ = v146
        _ = v151
    }
    fmt.Println(v2)
    fmt.Println(v2, v3)
    for fi152 := 0; fi152 < 15; fi152++ {
        for fi153 := 0; fi153 < 4; fi153++ {
            if true {
                v3 += (^(v3))
                switch fi152 {
                case -200:
                    if (false && (0 >= 2)) {
                        fmt.Println(v3, fi152, fi152)
                    } else {
                        delete(v1, fi152)
                        v3 = -1
                    }
                case 1:
                    v3 = (-(v3))
                }
                v154 := GenS19{F0: 4294967295, F1: 86, F2: (v3 % (((v3 / (((v3 ^ v3)) | 1))) | 1))}
                _ = v154
            }
            var v155 int16 = v3
            switch (fi152 / ((2147483647) | 1)) {
            case -2147483648:
                var v156 uint16 = 0
                _ = v156
            case 633:
                switch "\n\t" {
                case "":
                    var v157 uint = 255
                    _ = v157
                case "hello":
                    var v158 uint32 = 65535
                    _ = v158
                case "\n\t":
                    var v159 uint = 255
                    _ = v159
                default:
                    _ = len(v1)
                }
            case -625:
                switch -2147483648 {
                case -2147483648:
                    if true {
                        switch (^(fi152)) {
                        case -898:
                            var v160 int = -331
                            _ = v160
                        case -2147483648:
                            if true {
                                fmt.Println(v3, v155)
                            } else {
                                var v161 string = ("" + "hello")
                                _ = v161
                            }
                        case 944:
                            fmt.Println(v3, fi152)
                        default:
                            v155 = v3
                        }
                    } else {
                        v162 := GenS12{F0: fi153, F1: 2, F2: fi153}
                        _ = v162
                    }
                case 127:
                    v163 := v1[(^(fi153))]
                    _ = v163
                case 9223372036854775807:
                    switch "a" {
                    case "":
                        v1[(fi152 * -250)] = fi153
                    case "hello":
                        if ((fi153 | (fi152 % ((fi153) | 1))) != fi152) {
                            if (!false) {
                                var v164 bool = false
                                fmt.Println(v3)
                                _ = v164
                            } else {
                                var v165 int8 = 89
                                v2 = -283
                                var v166 uint64 = 543
                                v167 := GenS20{F0: (^(v2)), F1: 42}
                                _ = v165
                                _ = v166
                                _ = v167
                            }
                            v2 -= v2
                            switch "\n\t" {
                            case "\n\t":
                                v168 := GenS21{F0: 2147483647, F1: 42}
                                _ = v168
                            case "":
                                v169 := GenS1{F0: true}
                                _ = v169
                            case "a":
                                v170 := GenS9{F0: "hello", F1: 387, F2: 2}
                                _ = v170
                            }
                        }
                    case "foo bar":
                        if (1 < 127) {
                            switch fi152 {
                            case 9223372036854775807:
                                fmt.Println(v155)
                            default:
                                if true {
                                    if (false && (0 >= 976)) {
                                        var v171 int16 = (v155 / ((v3) | 1))
                                        fmt.Println(v3, v3, v2)
                                        if true {
                                            var v172 int = (fi153 << ((0) & 63))
                                            v155 = v155
                                            if true {
                                                v173 := []int{(^(v172))}
                                                v174 := GenS22{F0: (-(v171))}
                                                switch "\n\t" {
                                                case "":
                                                    if true {
                                                        var v175 int64 = v2
                                                        fmt.Println(fi153, v175)
                                                        v2 += (v2 & -682)
                                                        fmt.Println(v155)
                                                        _ = v175
                                                    }
                                                default:
                                                    if (0 < 2147483648) {
                                                        v3 -= v155
                                                        if (0 >= 1) {
                                                            _ = len(v173)
                                                            if len(v173) > 0 {
                                                                v173[0] = (-(v172))
                                                            }
                                                            v2 = 0
                                                        } else {
                                                            fmt.Println(v2, v171, v172)
                                                            if (!(!true)) {
                                                                var v176 uint = 2
                                                                var v177 int = 76
                                                                _ = v176
                                                                _ = v177
                                                            } else {
                                                                v178 := GenS23{F0: 42, F1: 42, F2: "foo bar"}
                                                                var v179 int = 128
                                                                _ = v178
                                                                _ = v179
                                                            }
                                                            var v180 int64 = 42
                                                            _ = v180
                                                        }
                                                        var v181 uint64 = 4294967295
                                                        _ = v181
                                                    }
                                                }
                                                _ = v173
                                                _ = v174
                                            } else {
                                                if (2147483647 <= -468) {
                                                    var v182 uint64 = 2
                                                    var v183 uint64 = (v182 - v182)
                                                    _ = v182
                                                    _ = v183
                                                } else {
                                                    if (1e10 < (0.0 / 3.14159)) {
                                                        var v184 uint = 728
                                                        switch "hello" {
                                                        case "hello":
                                                            v185 := []int{(fi153 + v172), fi152, 128, (fi152 - (fi152 & 0))}
                                                            _ = v185
                                                        case "foo bar":
                                                            v3 -= 32767
                                                        default:
                                                            switch "hello" {
                                                            case "a":
                                                                fmt.Println(v171)
                                                            case "foo bar":
                                                                fmt.Println(fi153)
                                                            case "hello":
                                                                v186 := GenS24{F0: (^(v184))}
                                                                _ = v186
                                                            default:
                                                                if false {
                                                                    if (!false) {
                                                                        switch "hello" {
                                                                        case "a":
                                                                            if (!(0 < (v155 >> ((v184) & 63)))) {
                                                                                _ = len(v1)
                                                                                if (2 != 117) {
                                                                                    var v187 int8 = 109
                                                                                    _ = v187
                                                                                } else {
                                                                                    var v188 uint8 = 2
                                                                                    _ = v188
                                                                                }
                                                                                var v189 int = 968
                                                                                _ = v189
                                                                            } else {
                                                                                if false {
                                                                                    var v190 float32 = -2.25
                                                                                    _ = v190
                                                                                } else {
                                                                                    var v191 uint8 = 2
                                                                                    v2 = v2
                                                                                    v3 += int16(v191)
                                                                                    var v192 int16 = 32767
                                                                                    _ = v191
                                                                                    _ = v192
                                                                                }
                                                                                v171 = v155
                                                                            }
                                                                            v193 := GenS25{F0: 0, F1: 255}
                                                                            _ = v193
                                                                        case "":
                                                                            var v194 uint64 = 42
                                                                            _ = v194
                                                                        }
                                                                        v1[(-(fi152))] = (fi152 << ((v184) & 63))
                                                                    } else {
                                                                        v195 := GenS3{F0: v2, F1: 4294967295, F2: (v184 * (v184 / (((v184 & 342)) | 1)))}
                                                                        _ = v195
                                                                    }
                                                                } else {
                                                                    v196 := GenS7{F0: 0}
                                                                    var v197 float64 = 3.14159
                                                                    v172 = v172
                                                                    _ = v196
                                                                    _ = v197
                                                                }
                                                            }
                                                        }
                                                        _ = v184
                                                    }
                                                }
                                                var v198 int8 = -41
                                                _ = v198
                                            }
                                            _ = v172
                                        } else {
                                            v155 = (v155 & v155)
                                        }
                                        _ = v171
                                    }
                                    v199 := GenS19{F0: 255, F1: 4294967295, F2: v155}
                                    var v200 uint64 = 412
                                    _ = v199
                                    _ = v200
                                } else {
                                    v2 -= (v2 >> ((65535) & 63))
                                }
                            }
                            var v201 int16 = 1
                            v2 = (v2 & v2)
                            var v202 int64 = int64(v155)
                            fmt.Println(v2, fi152)
                            switch fi152 {
                            case -958:
                                if true {
                                    var v203 int = (fi153 % (((fi153 * fi153)) | 1))
                                    var v204 uint = 42
                                    if ((32767 > 65535) && (false || (int16(v204) > v3))) {
                                        var v205 uint16 = 65535
                                        fmt.Println(v204)
                                        switch "\n\t" {
                                        case "foo bar":
                                            v202 = 0
                                        case "a":
                                            var v206 float64 = float64(1.5)
                                            _ = v206
                                        default:
                                            v155 -= (v3 * 32767)
                                        }
                                        _ = v205
                                    }
                                    v155 += v3
                                    var v207 uint32 = 2147483647
                                    _ = v203
                                    _ = v204
                                    _ = v207
                                }
                            case 32768:
                                v2 = 1
                            default:
                                v208 := GenS26{F0: 127, F1: (fi153 % (((fi153 >> ((1) & 63))) | 1)), F2: -124}
                                _ = v208
                            }
                            if false {
                                v209 := GenS6{F0: fi153, F1: 4294967295, F2: 1}
                                switch -484 {
                                case -1:
                                    v2 += v202
                                    v2 = v202
                                case 42:
                                    v1[fi153] = fi152
                                }
                                v2 += (v2 * v202)
                                _ = v209
                            }
                            _ = v201
                            _ = v202
                        } else {
                            if (!((!("foo bar" < ("\n\t" + ("" + "\n\t")))) || (!false))) {
                                fmt.Println(fi152, fi153)
                            } else {
                                var v210 int16 = v3
                                switch ("" + "hello") {
                                case "hello":
                                    fmt.Println(v155, fi152)
                                case "a":
                                    var v211 string = ((("" + "") + ("hello" + "a")) + "foo bar")
                                    _ = v211
                                default:
                                    var v212 uint = 124
                                    _ = v212
                                }
                                _ = v210
                            }
                            fmt.Println(fi153, fi152)
                        }
                    }
                }
            }
            _ = fi153
            _ = v155
        }
        if ((("" + "hello") + "") <= "") {
            v213 := v1[fi152]
            v3 = (v3 / (((v3 % ((v3) | 1))) | 1))
            for fi214 := 0; fi214 < 6; fi214++ {
                v2 = 1
                v215 := []int{}
                if true {
                    v216 := GenS4{F0: v2, F1: 1000}
                    _ = v216
                } else {
                    v217 := GenS2{F0: 255, F1: (v213 | fi214)}
                    _ = v217
                }
                _ = fi214
                _ = v215
            }
            _ = v213
        }
        v218 := []int{fi152, (fi152 & (fi152 + fi152))}
        _ = fi152
        _ = v218
    }
    var v219 int64 = v2
    if (!true) {
        if false {
            fmt.Println(v2, v219)
            v3 = v3
            delete(v1, -2147483648)
        } else {
            var v220 int16 = -1
            switch "" {
            case "a":
                for fi221 := 0; fi221 < 8; fi221++ {
                    if (0.0 > (0.1 * float32(1e10))) {
                        if false {
                            switch (fi221 + 42) {
                            case -139:
                                var v222 float64 = (((-2.25 - (float64((float32(0.0) - (0.1 - (0.1 - 0.1)))) + 3.14159)) - (1.5 + (float64((float32(1e10) + 3.14159)) + -2.25))) + 0.0)
                                _ = v222
                            case -2147483648:
                                if true {
                                    if (42 < 450) {
                                        switch "foo bar" {
                                        case "hello":
                                            v223 := GenS14{F0: 42}
                                            _ = v223
                                        }
                                    } else {
                                        var v224 string = "foo bar"
                                        v224 = v224
                                        if true {
                                            for fi225 := 0; fi225 < 13; fi225++ {
                                                v3 = (v3 >> ((4294967295) & 63))
                                                if true {
                                                    fmt.Println(v2)
                                                }
                                                v226 := GenS27{F0: 1, F1: true}
                                                _ = fi225
                                                _ = v226
                                            }
                                            if (!(2147483648 == 65535)) {
                                                v220 += v3
                                            }
                                        }
                                        _ = v224
                                    }
                                    if false {
                                        fmt.Println(v219, v220)
                                        var v227 int64 = v219
                                        _ = v227
                                    }
                                    v2 = (^(v219))
                                }
                            case 32767:
                                var v228 uint32 = 4294967295
                                _ = v228
                            default:
                                if (0 >= 2) {
                                    fmt.Println(fi221, v219, v2)
                                } else {
                                    v229 := GenS4{F0: -9223372036854775808, F1: v220}
                                    switch ("" + "hello") {
                                    case "a":
                                        var v230 float64 = (3.14159 * 1.5)
                                        _ = v230
                                    case "foo bar":
                                        var v231 uint32 = 4294967295
                                        _ = v231
                                    }
                                    _ = v229
                                }
                            }
                            v232 := v1[fi221]
                            fmt.Println(v2, v220, v232)
                            _ = v232
                        } else {
                            for fi233 := 0; fi233 < 13; fi233++ {
                                var v234 float64 = 0.0
                                v220 -= -322
                                v3 += 0
                                _ = fi233
                                _ = v234
                            }
                            if true {
                                var v235 uint64 = 2
                                var v236 uint32 = 42
                                _ = v235
                                _ = v236
                            } else {
                                var v237 uint32 = 613
                                v3 = (v220 & (v3 | (-(v220))))
                                fmt.Println(v3, v220)
                                _ = v237
                            }
                            v2 += 9223372036854775807
                        }
                        v238 := []int{}
                        v239 := GenS18{F0: 0, F1: 255}
                        _ = v238
                        _ = v239
                    } else {
                        v3 = v3
                        var v240 int64 = 42
                        for fi241 := 0; fi241 < 5; fi241++ {
                            fmt.Println(v3, v219, v220)
                            v2 -= v240
                            v242 := GenS5{F0: (fi241 - fi241), F1: 128, F2: 799}
                            _ = fi241
                            _ = v242
                        }
                        _ = v240
                    }
                    for fi243 := 0; fi243 < 12; fi243++ {
                        v3 -= 506
                        v244 := GenS28{F0: (v2 | v2)}
                        _ = fi243
                        _ = v244
                    }
                    _ = fi221
                }
            case "foo bar":
                var v245 string = "\n\t"
                _ = v245
            default:
                v246 := []int{32768, 748}
                _ = v246
            }
            _ = v220
        }
        v247 := GenS19{F0: 2, F1: 42, F2: 32767}
        _ = v247
    } else {
        v248 := []int{42, 640, 761}
        _ = v248
    }
    for fi249 := 0; fi249 < 1; fi249++ {
        fmt.Println(fi249, fi249, v2)
        fmt.Println(fi249, fi249, fi249)
        _ = fi249
    }
    switch -9223372036854775808 {
    case 229:
        var v250 float32 = 1e10
        _ = v250
    case 396:
        v251 := GenS28{F0: 174}
        _ = v251
    case 42:
        if true {
            for fi252 := 0; fi252 < 3; fi252++ {
                var v253 int8 = 50
                var v254 int = fi252
                _ = fi252
                _ = v253
                _ = v254
            }
        }
    default:
        v2 = v2
    }
    for fi255 := 0; fi255 < 20; fi255++ {
        v256 := []int{(^(fi255)), fi255}
        _ = fi255
        _ = v256
    }
    for fi257 := 0; fi257 < 12; fi257++ {
        for fi258 := 0; fi258 < 14; fi258++ {
            v1[fi258] = fi258
            _ = fi258
        }
        fmt.Println(v3, v219, v219)
        _ = fi257
    }
    if true {
        var v259 int = 2147483647
        for fi260 := 0; fi260 < 7; fi260++ {
            switch "\n\t" {
            case "hello":
                v259 += fi260
            case "":
                var v261 uint = 255
                _ = v261
            }
            _ = fi260
        }
        _ = v259
    }
    v3 = (v3 + (-(v3)))
    v1[-2147483648] = 9223372036854775807
    for fi262 := 0; fi262 < 13; fi262++ {
        switch "\n\t" {
        case "a":
            v263 := GenS29{F0: "a", F1: false}
            _ = v263
        case "hello":
            fmt.Println(v219, v219)
        }
        _ = fi262
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v219:", v219)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v219
    _ = v3
}
