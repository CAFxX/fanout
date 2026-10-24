package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8; F1 uint64; F2 uint }
type GenS8 struct { F0 uint; F1 uint64; F2 bool }
type GenS9 struct { F0 uint8; F1 uint32; F2 bool }
type GenS10 struct { F0 bool; F1 int32; F2 string }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint; F1 int16; F2 bool }
type GenS13 struct { F0 uint32; F1 uint }
type GenS14 struct { F0 uint8; F1 string; F2 int32 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint32; F1 bool; F2 int8 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 int32; F1 int32; F2 int32 }
func genFn0(p0 int32, p1 int32, p2 int) bool {
    var v7 int = 9223372036854775807
    v8 := GenS1{F0: -128, F1: 55, F2: 0}
    if (59 <= 127) {
        for fi9 := 0; fi9 < 14; fi9++ {
            fmt.Println(fi9)
            _ = fi9
        }
        var v10 float64 = 3.14159
        var v11 float64 = ((v10 / 0.1) - v10)
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 9; fi13++ {
                v7 = (^(v7))
                if (32767 <= 32767) {
                    fmt.Println(fi12)
                } else {
                    var v14 int32 = (p0 + p0)
                    _ = v14
                }
                v15 := GenS2{F0: 517}
                _ = fi13
                _ = v15
            }
            _ = fi12
        }
        for fi16 := 0; fi16 < 1; fi16++ {
            switch 642 {
            case 32768:
                var v17 float64 = ((0.0 - 3.14159) + -2.25)
                _ = v17
            default:
                v7 = (^(fi16))
            }
            _ = fi16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ((false && false) || (!(((0.1 - 1e10) != float64(0.0)) || false)))
}

func main() {
    var v1 int16 = -373
    v2 := []string{"a", ("\n\t" + ("hello" + ("a" + "foo bar"))), ("\n\t" + "foo bar")}
    var v3 int16 = v1
    var v4 int = 4294967295
    var v5 uint64 = 42
    var v6 uint8 = 1
    var v18 float64 = -2.25
    if (v5 <= (v5 << ((65535) & 63))) {
        _ = len(v2)
        fmt.Println(v5, v18, v6)
    }
    if true {
        switch (^(v4)) {
        case -964:
            if ((("hello" + "") + ("\n\t" + "a")) == "") {
                if ((v4 % (((-(v4))) | 1)) <= 946) {
                    v19 := GenS3{F0: 4294967295}
                    var v20 bool = (18446744073709551615 == 2)
                    v6 += (v6 % ((v6) | 1))
                    _ = v19
                    _ = v20
                }
            } else {
                switch "hello" {
                case "hello":
                    v21 := GenS2{F0: v5}
                    _ = v21
                case "\n\t":
                    if (uint16(v6) > 32767) {
                        for fi22 := 0; fi22 < 6; fi22++ {
                            switch "foo bar" {
                            case "hello":
                                switch 4294967295 {
                                case 2147483647:
                                    var v23 int = (v4 | (v4 | 32768))
                                    _ = v23
                                case 1:
                                    v18 -= float64((-2.25 - 0.1))
                                case -2147483648:
                                    for fi24 := 0; fi24 < 7; fi24++ {
                                        _ = len(v2)
                                        switch "a" {
                                        case "a":
                                            switch (-(fi24)) {
                                            case 32768:
                                                switch int(v6) {
                                                case 127:
                                                    var v25 int8 = -1
                                                    _ = v25
                                                case 4294967295:
                                                    if false {
                                                        if false {
                                                            if true {
                                                                if (false && (!false)) {
                                                                    v18 -= 0.1
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            v3 += v3
                                                                        case "a":
                                                                            var v26 float32 = 1.5
                                                                            _ = v26
                                                                        case "foo bar":
                                                                            v6 += v6
                                                                        }
                                                                    case "foo bar":
                                                                        v27 := GenS3{F0: v5}
                                                                        _ = v27
                                                                    case "hello":
                                                                        v4 = v4
                                                                    default:
                                                                        var v28 int8 = 127
                                                                        _ = v28
                                                                    }
                                                                }
                                                                v5 = 347
                                                            }
                                                            if (877 >= uint(v6)) {
                                                                if len(v2) > 0 {
                                                                    v2[0] = "hello"
                                                                }
                                                            } else {
                                                                var v29 bool = false
                                                                var v30 int32 = 2147483647
                                                                var v31 int8 = 0
                                                                if v29 {
                                                                    v18 += ((v18 * float64(3.14159)) - v18)
                                                                }
                                                                fmt.Println(v30, v5)
                                                                _ = v29
                                                                _ = v30
                                                                _ = v31
                                                            }
                                                            switch (fi22 | -1) {
                                                            case 32767:
                                                                v3 = v3
                                                            case -537:
                                                                fmt.Println(v1, v4)
                                                            case 4294967295:
                                                                if (!(647 == int64(fi22))) {
                                                                    if true {
                                                                        var v32 bool = false
                                                                        v33 := GenS4{F0: "a", F1: 65535}
                                                                        fmt.Println(fi22, v5, v3)
                                                                        _ = v32
                                                                        _ = v33
                                                                    }
                                                                    _ = len(v2)
                                                                }
                                                            }
                                                        }
                                                    }
                                                default:
                                                    v34 := GenS4{F0: "hello", F1: 299}
                                                    _ = v34
                                                }
                                            case -2147483648:
                                                var v35 uint16 = 342
                                                _ = v35
                                            case 0:
                                                var v36 uint8 = v6
                                                _ = v36
                                            default:
                                                switch ("foo bar" + "\n\t") {
                                                case "hello":
                                                    var v37 uint = 255
                                                    _ = v37
                                                case "\n\t":
                                                    v38 := GenS2{F0: 1}
                                                    _ = v38
                                                case "":
                                                    var v39 uint16 = 255
                                                    _ = v39
                                                }
                                            }
                                        case "hello":
                                            var v40 uint32 = 2
                                            _ = v40
                                        case "\n\t":
                                            v41 := GenS5{F0: 1}
                                            _ = v41
                                        }
                                        _ = fi24
                                    }
                                default:
                                    fmt.Println(v1)
                                }
                            case "foo bar":
                                if (true && (v5 >= (v5 >> ((18446744073709551615) & 63)))) {
                                    if false {
                                        for fi42 := 0; fi42 < 4; fi42++ {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if true {
                                                var v43 uint32 = 2147483647
                                                v4 = v4
                                                var v44 string = ""
                                                _ = v43
                                                _ = v44
                                            } else {
                                                v1 = -30
                                                switch v4 {
                                                case -662:
                                                    v1 = v3
                                                case 9223372036854775807:
                                                    var v45 uint16 = 42
                                                    _ = v45
                                                case -136:
                                                    v5 = (v5 / ((v5) | 1))
                                                default:
                                                    v46 := GenS6{F0: 638}
                                                    _ = v46
                                                }
                                                if (!(18446744073709551615 == 0)) {
                                                    var v47 uint32 = 735
                                                    v2 = append(v2, ("hello" + "hello"))
                                                    _ = v47
                                                } else {
                                                    _ = len(v2)
                                                }
                                            }
                                            if (((v5 >> ((18446744073709551615) & 63)) == (v5 << ((548) & 63))) && false) {
                                                v1 -= v1
                                                switch fi42 {
                                                case 127:
                                                    v48 := GenS7{F0: (v6 * v6), F1: v5, F2: 957}
                                                    _ = v48
                                                case 32768:
                                                    v18 += 1.5
                                                default:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                            } else {
                                                if (true && false) {
                                                    v49 := GenS8{F0: 233, F1: v5, F2: true}
                                                    _ = v49
                                                } else {
                                                    if (!false) {
                                                        if false {
                                                            fmt.Println(v6)
                                                        }
                                                    }
                                                    var v50 uint8 = 2
                                                    if len(v2) > 0 {
                                                        v2[0] = ("hello" + "foo bar")
                                                    }
                                                    _ = v50
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            _ = fi42
                                        }
                                        for fi51 := 0; fi51 < 7; fi51++ {
                                            v52 := GenS9{F0: (v6 % ((128) | 1)), F1: 42, F2: true}
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v5 = v5
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            switch (fi22 << ((18446744073709551615) & 63)) {
                                            case 958:
                                                if true {
                                                    if (9223372036854775807 >= fi51) {
                                                        v1 = v3
                                                    }
                                                    switch fi22 {
                                                    case 128:
                                                        switch ("\n\t" + ("foo bar" + "hello")) {
                                                        case "":
                                                            var v53 int8 = 100
                                                            _ = v53
                                                        case "a":
                                                            v54 := GenS10{F0: false, F1: 388, F2: "foo bar"}
                                                            _ = v54
                                                        case "hello":
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        default:
                                                            if true {
                                                                switch (("" + "a") + "hello") {
                                                                case "hello":
                                                                    var v55 int8 = -72
                                                                    _ = v55
                                                                case "a":
                                                                    v3 -= v1
                                                                default:
                                                                    v1 = v1
                                                                }
                                                                if false {
                                                                    if true {
                                                                        if true {
                                                                            var v56 string = ("" + "\n\t")
                                                                            fmt.Println(v18)
                                                                            switch 1 {
                                                                            case 127:
                                                                                v2 = append(v2, "\n\t")
                                                                            case 32768:
                                                                                v3 -= v3
                                                                                v3 = (v1 ^ (v1 - (v3 * v3)))
                                                                                _ = len(v2)
                                                                                fmt.Println(fi22, fi22)
                                                                                fmt.Println(v18, v1)
                                                                                v56 = v56
                                                                            }
                                                                            fmt.Println(v18, v5, v5)
                                                                            v4 -= v4
                                                                            _ = v56
                                                                        } else {
                                                                            var v57 uint64 = (^(v5))
                                                                            _ = v57
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        var v58 uint8 = v6
                                                        _ = v58
                                                    }
                                                } else {
                                                    v59 := GenS1{F0: -86, F1: 1, F2: 0}
                                                    fmt.Println(v5, v4)
                                                    _ = v59
                                                }
                                            case 260:
                                                if (false || (-200 != 1)) {
                                                    var v60 float32 = 1e10
                                                    if true {
                                                        v61 := GenS11{F0: 0}
                                                        v6 -= 16
                                                        var v62 int16 = v3
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = "hello"
                                                    }
                                                    _ = v60
                                                }
                                            default:
                                                v63 := GenS11{F0: -1}
                                                if (2 <= 2147483647) {
                                                    v6 += 2
                                                }
                                                _ = v63
                                            }
                                            _ = fi51
                                            _ = v52
                                        }
                                        var v64 uint32 = 0
                                        _ = v64
                                    }
                                    fmt.Println(v4, fi22)
                                    var v65 float64 = float64(((1e10 - float32((v18 - (float64(3.14159) / (v18 + (v18 + (v18 / v18))))))) + 0.0))
                                    _ = v65
                                } else {
                                    if true {
                                        if true {
                                            for fi66 := 0; fi66 < 18; fi66++ {
                                                v67 := GenS12{F0: 65535, F1: v3, F2: false}
                                                _ = fi66
                                                _ = v67
                                            }
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v68 uint8 = 42
                                                _ = v68
                                            case "hello":
                                                var v69 bool = ((v6 == (v6 + (v6 >> ((42) & 63)))) && false)
                                                _ = v69
                                            case "":
                                                if (100 >= 54) {
                                                    var v70 uint64 = (v5 % ((v5) | 1))
                                                    v5 = (v70 >> ((18446744073709551615) & 63))
                                                    v5 += 456
                                                    _ = v70
                                                }
                                            default:
                                                for fi71 := 0; fi71 < 3; fi71++ {
                                                    switch fi71 {
                                                    case -59:
                                                        fmt.Println(v3, v3, v18)
                                                    case -392:
                                                        v3 = (-(v3))
                                                    }
                                                    v72 := GenS13{F0: 992, F1: 2}
                                                    _ = fi71
                                                    _ = v72
                                                }
                                            }
                                        }
                                        var v73 uint = 726
                                        _ = v73
                                    }
                                    v4 = (v4 >> ((543) & 63))
                                    for fi74 := 0; fi74 < 4; fi74++ {
                                        _ = len(v2)
                                        var v75 uint32 = 65535
                                        _ = fi74
                                        _ = v75
                                    }
                                }
                            case "a":
                                switch (fi22 + (fi22 % ((42) | 1))) {
                                case 0:
                                    var v76 uint8 = v6
                                    _ = v76
                                case 128:
                                    v3 -= v1
                                case -297:
                                    var v77 uint8 = v6
                                    _ = v77
                                }
                            }
                            switch "hello" {
                            case "a":
                                v18 = 0.1
                            case "\n\t":
                                var v78 int = fi22
                                _ = v78
                            case "":
                                if false {
                                    v18 = 1e10
                                } else {
                                    v1 += 32767
                                    for fi79 := 0; fi79 < 17; fi79++ {
                                        if true {
                                            switch v4 {
                                            case 32767:
                                                var v80 uint32 = 65535
                                                _ = v80
                                            case 1:
                                                var v81 bool = (false || (42 >= 255))
                                                _ = v81
                                            default:
                                                switch ("" + "foo bar") {
                                                case "\n\t":
                                                    v3 += v3
                                                case "a":
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v6 += 1
                                                    case "foo bar":
                                                        var v82 bool = false
                                                        _ = v82
                                                    case "hello":
                                                        fmt.Println(v3, v18, v3)
                                                    default:
                                                        v83 := GenS8{F0: uint(v6), F1: v5, F2: false}
                                                        _ = v83
                                                    }
                                                }
                                            }
                                            v3 = -202
                                            fmt.Println(v18, v6, fi22)
                                        } else {
                                            v4 = fi79
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = fi79
                                    }
                                    if ((v5 >> ((135) & 63)) <= v5) {
                                        switch "hello" {
                                        case "foo bar":
                                            if len(v2) > 0 {
                                                v2[0] = (("a" + ("foo bar" + "\n\t")) + "\n\t")
                                            }
                                        case "":
                                            v84 := GenS14{F0: (^(v6)), F1: "hello", F2: 966}
                                            _ = v84
                                        default:
                                            fmt.Println(v6)
                                        }
                                    }
                                }
                            }
                            var v85 int16 = v1
                            _ = fi22
                            _ = v85
                        }
                        switch "hello" {
                        case "foo bar":
                            v18 = 3.14159
                        case "":
                            switch "a" {
                            case "hello":
                                if (!(float32(((-2.25 * v18) + ((v18 * (float64(0.0) * float64(0.1))) - v18))) != (0.1 * -2.25))) {
                                    v86 := GenS11{F0: -912}
                                    v3 -= v1
                                    var v87 int32 = -1
                                    _ = v86
                                    _ = v87
                                }
                            case "foo bar":
                                for fi88 := 0; fi88 < 8; fi88++ {
                                    var v89 int32 = -2147483648
                                    v90 := GenS10{F0: (4294967295 <= 814), F1: (v89 - 2147483647), F2: "hello"}
                                    _ = fi88
                                    _ = v89
                                    _ = v90
                                }
                            case "":
                                var v91 uint64 = (^(v5))
                                _ = v91
                            default:
                                var v92 int16 = v3
                                _ = v92
                            }
                        }
                        v1 += 0
                    }
                }
                v1 -= v3
                var v93 int16 = v1
                _ = v93
            }
        case 9223372036854775807:
            if true {
                var v94 int8 = -88
                v18 += float64(0.0)
                if (false || ((v4 & 959) < (v4 - (v4 % (((v4 / ((v4) | 1))) | 1))))) {
                    v5 -= 2
                } else {
                    v95 := GenS5{F0: 898}
                    _ = v95
                }
                _ = v94
            } else {
                v96 := GenS8{F0: 65535, F1: (^(v5)), F2: true}
                v97 := GenS15{F0: 117}
                if false {
                    for fi98 := 0; fi98 < 3; fi98++ {
                        fmt.Println(v18)
                        _ = fi98
                    }
                    v99 := GenS6{F0: 32767}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v99
                }
                _ = v96
                _ = v97
            }
        }
        if false {
            v6 = v6
            if (!true) {
                for fi100 := 0; fi100 < 13; fi100++ {
                    switch ((("a" + "") + (("foo bar" + ("foo bar" + "\n\t")) + ("" + ("a" + ("\n\t" + "a"))))) + "hello") {
                    case "foo bar":
                        if (1 < uint32(v6)) {
                            v1 -= 42
                            var v101 uint64 = (v5 >> ((1) & 63))
                            for fi102 := 0; fi102 < 4; fi102++ {
                                v103 := GenS16{F0: 65535, F1: true, F2: 126}
                                _ = fi102
                                _ = v103
                            }
                            _ = v101
                        }
                    case "a":
                        var v104 int16 = (v1 * (v3 - 0))
                        _ = v104
                    }
                    if len(v2) > 0 {
                        v2[0] = ("foo bar" + "\n\t")
                    }
                    v4 += v4
                    _ = fi100
                }
            } else {
                for fi105 := 0; fi105 < 14; fi105++ {
                    v3 = (v1 ^ (^(v1)))
                    v1 -= (v1 | (v3 & (v3 + (v3 & 146))))
                    v106 := GenS17{F0: 2147483647}
                    _ = fi105
                    _ = v106
                }
                v4 -= -534
            }
            if true {
                v18 += ((1e10 * (float64(-2.25) * v18)) + (v18 * (3.14159 * ((0.0 / v18) + -2.25))))
                v107 := GenS18{F0: 1, F1: 360, F2: 490}
                fmt.Println(v4, v3)
                _ = v107
            }
        } else {
            v18 = 0.1
            v4 -= (-(v4))
        }
        fmt.Println(v4, v4, v18)
    }
    var v108 int8 = 85
    v3 -= (v1 | v1)
    v5 += 18446744073709551615
    fmt.Println(v108)
    v2 = append(v2, "foo bar")
    v6 = (v6 & 128)
    var v109 float32 = ((1.5 / -2.25) * 3.14159)
    var v110 int = -36
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v108
    _ = v109
    _ = v110
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
