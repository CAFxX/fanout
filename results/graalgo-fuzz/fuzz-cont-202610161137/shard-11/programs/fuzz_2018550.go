package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 string; F1 int16; F2 int16 }
type GenS4 struct { F0 int; F1 int32; F2 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 uint16; F2 int }
type GenS7 struct { F0 int8; F1 int32; F2 string }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int64; F1 int16 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint32; F1 int8 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 string; F1 int32 }
type GenS18 struct { F0 string; F1 int; F2 string }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 bool; F1 uint8 }
type GenS21 struct { F0 uint }
func genFn0(p0 int16, p1 uint64, p2 uint) string {
    p1 = 42
    v5 := []int{128}
    var v6 float32 = (((0.1 / 0.1) * 3.14159) / 1e10)
    if true {
        if false {
            v7 := GenS1{F0: 439}
            _ = v7
        }
        v8 := GenS2{F0: 1, F1: false}
        v9 := GenS1{F0: 32767}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return "a"
}
func genFn1(p0 int, p1 uint16) int64 {
    var v10 uint64 = 2
    var v11 float64 = ((-2.25 * float64(0.1)) + (0.1 - 1.5))
    if false {
        v12 := []int{p0, p0, p0, p0}
        if false {
            for fi13 := 0; fi13 < 7; fi13++ {
                v14 := GenS2{F0: 589, F1: (-128 != -1)}
                for fi15 := 0; fi15 < 15; fi15++ {
                    p0 -= fi13
                    fmt.Println(fi15, fi13, p0)
                    p0 += 128
                    p0 = -1
                    _ = fi15
                }
                v11 = (v11 / (float64(1e10) * v11))
                _ = fi13
                _ = v14
            }
        } else {
            v11 = ((((v11 * ((((0.1 + ((v11 - v11) / v11)) * v11) - v11) / ((v11 + float64(float32(float64((1e10 / -2.25))))) * 0.0))) * (-2.25 - ((-2.25 - float64((0.0 + 0.1))) - (v11 / (v11 * (v11 + v11)))))) + float64((3.14159 + 3.14159))) + float64(1.5))
        }
        if ((float32(-2.25) <= (-2.25 / 1.5)) || (127 > -128)) {
            if (244 <= 0) {
                if ((9223372036854775807 < 1) && true) {
                    p1 += (p1 | (p1 + 32767))
                    var v16 uint16 = 1
                    for fi17 := 0; fi17 < 7; fi17++ {
                        var v18 uint64 = (v10 | 4294967295)
                        fmt.Println(fi17, v18)
                        v10 += 18446744073709551615
                        _ = fi17
                        _ = v18
                    }
                    _ = v16
                }
                switch (p0 / (((p0 - p0)) | 1)) {
                case 874:
                    for fi19 := 0; fi19 < 18; fi19++ {
                        fmt.Println(p1, p0, fi19)
                        if (4294967295 < -823) {
                            v12 = append(v12, (fi19 % (((^(fi19))) | 1)))
                        }
                        _ = fi19
                    }
                case 1:
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                case 0:
                    var v20 int8 = -128
                    _ = v20
                }
                fmt.Println(p1, v10)
            } else {
                if (32767 > (p1 >> ((18446744073709551615) & 63))) {
                    for fi21 := 0; fi21 < 13; fi21++ {
                        if (!false) {
                            p1 = 0
                            if (!(!(float32(v11) >= (-2.25 / ((0.1 + 1e10) / (-2.25 * 1e10)))))) {
                                for fi22 := 0; fi22 < 5; fi22++ {
                                    switch 128 {
                                    case 32767:
                                        switch "" {
                                        case "\n\t":
                                            var v23 int64 = 42
                                            _ = v23
                                        case "hello":
                                            if ((!false) || (255 > 212)) {
                                                var v24 int64 = 501
                                                _ = v24
                                            } else {
                                                v25 := GenS3{F0: "foo bar", F1: 1000, F2: 1000}
                                                if true {
                                                    var v26 int8 = 100
                                                    v26 += (v26 % ((v26) | 1))
                                                    switch ("hello" + "hello") {
                                                    case "":
                                                        fmt.Println(p1, v11, p0)
                                                        switch 32767 {
                                                        case 0:
                                                            switch (("" + "") + "a") {
                                                            case "\n\t":
                                                                fmt.Println(v26, fi22)
                                                            case "a":
                                                                fmt.Println(fi21, v10)
                                                            }
                                                            switch (p0 / ((-438) | 1)) {
                                                            case -755:
                                                                if false {
                                                                    fmt.Println(p0)
                                                                    if true {
                                                                        p1 = (p1 % (((p1 + p1)) | 1))
                                                                    } else {
                                                                        if true {
                                                                            if false {
                                                                                fmt.Println(v11, fi22, p1)
                                                                            } else {
                                                                                var v27 int64 = -9223372036854775808
                                                                                fmt.Println(v26, fi22, p1)
                                                                                v10 = (v10 + 2)
                                                                                _ = v27
                                                                            }
                                                                        }
                                                                        switch 431 {
                                                                        case 0:
                                                                            fmt.Println(v10)
                                                                        case -643:
                                                                            if (true || true) {
                                                                                v28 := GenS4{F0: int(v26), F1: 2147483647, F2: 2147483648}
                                                                                if len(v12) > 0 {
                                                                                    v12[0] = fi21
                                                                                }
                                                                                if (-1 == 110) {
                                                                                    var v29 bool = false
                                                                                    fmt.Println(fi21, p0)
                                                                                    var v30 uint8 = 1
                                                                                    if (!v29) {
                                                                                        switch (("a" + ("hello" + "hello")) + "\n\t") {
                                                                                        case "foo bar":
                                                                                            var v31 float32 = 0.1
                                                                                            _ = v31
                                                                                        case "hello":
                                                                                            var v32 int16 = -422
                                                                                            _ = v32
                                                                                        case "\n\t":
                                                                                            var v33 float32 = -2.25
                                                                                            _ = v33
                                                                                        default:
                                                                                            fmt.Println(v30)
                                                                                        }
                                                                                        if ((v30 ^ (v30 >> ((2) & 63))) != (^(v30))) {
                                                                                            var v34 int = (^(p0))
                                                                                            _ = v34
                                                                                        } else {
                                                                                            fmt.Println(fi22)
                                                                                            v30 -= (v30 - v30)
                                                                                            v35 := GenS5{F0: (^(p1))}
                                                                                            _ = v35
                                                                                        }
                                                                                    }
                                                                                    v30 = v30
                                                                                    _ = v29
                                                                                    _ = v30
                                                                                } else {
                                                                                    if false {
                                                                                        var v36 int8 = v26
                                                                                        if true {
                                                                                            if len(v12) > 0 {
                                                                                                v12[0] = p0
                                                                                            }
                                                                                        }
                                                                                        _ = len(v12)
                                                                                        _ = v36
                                                                                    }
                                                                                    if (false && false) {
                                                                                        switch "hello" {
                                                                                        case "hello":
                                                                                            var v37 int32 = -1
                                                                                            _ = v37
                                                                                        case "":
                                                                                            var v38 int8 = v26
                                                                                            _ = v38
                                                                                        default:
                                                                                            switch 9223372036854775807 {
                                                                                            case 2147483647:
                                                                                                if len(v12) >= 2 {
                                                                                                    v12 = v12[1:len(v12)]
                                                                                                }
                                                                                            default:
                                                                                                var v39 uint64 = 783
                                                                                                _ = v39
                                                                                            }
                                                                                        }
                                                                                        var v40 float32 = 0.0
                                                                                        if len(v12) > 0 {
                                                                                            v12[0] = p0
                                                                                        }
                                                                                        _ = v40
                                                                                    }
                                                                                }
                                                                                _ = v28
                                                                            } else {
                                                                                fmt.Println(p0)
                                                                                v11 -= -2.25
                                                                            }
                                                                        case 32768:
                                                                            v26 = -128
                                                                        default:
                                                                            v10 += (v10 - (^(v10)))
                                                                        }
                                                                        var v41 int8 = 127
                                                                        switch (p0 % ((0) | 1)) {
                                                                        case -1:
                                                                            if true {
                                                                                v42 := GenS6{F0: v10, F1: (p1 >> ((65535) & 63)), F2: -1}
                                                                                fmt.Println(v11)
                                                                                if false {
                                                                                    v10 = (v10 << ((2) & 63))
                                                                                    p0 = p0
                                                                                    fmt.Println(v26, p0, fi22)
                                                                                    var v43 uint32 = 529
                                                                                    _ = v43
                                                                                } else {
                                                                                    p1 += p1
                                                                                    if (!true) {
                                                                                        v12 = append(v12, int(v10))
                                                                                        if len(v12) > 0 {
                                                                                            v12[0] = (^(fi22))
                                                                                        }
                                                                                        v44 := GenS7{F0: (v41 ^ v26), F1: 849, F2: ("hello" + "")}
                                                                                        _ = v44
                                                                                    }
                                                                                }
                                                                                _ = v42
                                                                            }
                                                                        case 285:
                                                                            var v45 int32 = -2147483648
                                                                            _ = v45
                                                                        case -665:
                                                                            v12 = append(v12, -1)
                                                                        default:
                                                                            if false {
                                                                                v46 := GenS5{F0: p1}
                                                                                p0 -= (p0 / ((1) | 1))
                                                                                if len(v12) >= 2 {
                                                                                    v12 = v12[1:len(v12)]
                                                                                }
                                                                                _ = v46
                                                                            } else {
                                                                                v10 += 269
                                                                                v47 := GenS6{F0: 765, F1: p1, F2: fi21}
                                                                                _ = v47
                                                                            }
                                                                        }
                                                                        v10 = (^(v10))
                                                                        _ = v41
                                                                    }
                                                                }
                                                            default:
                                                                v48 := GenS2{F0: 18446744073709551615, F1: (1 < 1)}
                                                                _ = v48
                                                            }
                                                            var v49 uint64 = 4294967295
                                                            _ = v49
                                                        }
                                                    case "hello":
                                                        v11 -= 1e10
                                                    }
                                                    fmt.Println(v26, fi22)
                                                    _ = v26
                                                } else {
                                                    if len(v12) > 0 {
                                                        v12[0] = (p0 >> ((4294967295) & 63))
                                                    }
                                                    v50 := GenS8{F0: 339, F1: ("\n\t" + ("a" + ("foo bar" + "\n\t")))}
                                                    var v51 int8 = 103
                                                    _ = v50
                                                    _ = v51
                                                }
                                                switch ("foo bar" + ("" + "a")) {
                                                case "foo bar":
                                                    switch "hello" {
                                                    case "":
                                                        v10 = (v10 / (((v10 / ((1) | 1))) | 1))
                                                    default:
                                                        v11 = (v11 * v11)
                                                    }
                                                case "":
                                                    if (17 > -41) {
                                                        p0 = 0
                                                    } else {
                                                        var v52 uint8 = 127
                                                        p0 += (fi22 & fi21)
                                                        _ = v52
                                                    }
                                                default:
                                                    if ((!((!false) || false)) || true) {
                                                        fmt.Println(v11, fi22)
                                                        v11 = float64(1e10)
                                                        if true {
                                                            v10 = 4294967295
                                                            switch (^(p0)) {
                                                            case 1:
                                                                if false {
                                                                    p0 = (^(fi22))
                                                                } else {
                                                                    v53 := GenS6{F0: v10, F1: uint16(v10), F2: (-(fi22))}
                                                                    if ((v11 - v11) < v11) {
                                                                        v54 := GenS8{F0: 868, F1: "foo bar"}
                                                                        var v55 float32 = 0.1
                                                                        _ = v54
                                                                        _ = v55
                                                                    }
                                                                    var v56 int16 = 42
                                                                    _ = v53
                                                                    _ = v56
                                                                }
                                                                var v57 int8 = 42
                                                                var v58 float64 = ((3.14159 - v11) * v11)
                                                                _ = v57
                                                                _ = v58
                                                            case -1:
                                                                p0 -= (fi22 / (((fi21 / ((fi22) | 1))) | 1))
                                                            case 0:
                                                                if false {
                                                                    p1 -= (p1 | (p1 * p1))
                                                                }
                                                            }
                                                            v11 -= (1.5 / -2.25)
                                                            fmt.Println(p1, p1)
                                                        } else {
                                                            var v59 int8 = 118
                                                            v60 := GenS9{F0: ((((("foo bar" + "") + "foo bar") + "hello") + ("foo bar" + "\n\t")) + "foo bar"), F1: 0}
                                                            var v61 bool = (false && (42 > 1))
                                                            _ = v59
                                                            _ = v60
                                                            _ = v61
                                                        }
                                                        var v62 int8 = -27
                                                        var v63 uint16 = 65535
                                                        _ = v62
                                                        _ = v63
                                                    } else {
                                                        if (!true) {
                                                            var v64 int32 = 1
                                                            switch (p0 - (fi21 - 2147483647)) {
                                                            case -291:
                                                                switch (fi22 >> ((4294967295) & 63)) {
                                                                case 4294967295:
                                                                    v65 := GenS9{F0: "foo bar", F1: v64}
                                                                    _ = v65
                                                                case -9223372036854775808:
                                                                    switch (("hello" + ("hello" + "hello")) + ("" + "hello")) {
                                                                    case "a":
                                                                        var v66 int8 = -128
                                                                        _ = v66
                                                                    case "":
                                                                        var v67 uint64 = (v10 << ((42) & 63))
                                                                        _ = v67
                                                                    default:
                                                                        v68 := GenS3{F0: "", F1: 1000, F2: 0}
                                                                        _ = v68
                                                                    }
                                                                case 0:
                                                                    v12 = append(v12, 9223372036854775807)
                                                                default:
                                                                    v69 := GenS10{F0: 42}
                                                                    _ = v69
                                                                }
                                                            case -621:
                                                                p1 = 42
                                                            case -699:
                                                                if (false || (true && (65535 != 458))) {
                                                                    p1 -= 506
                                                                    var v70 int = -550
                                                                    v64 += (v64 ^ 247)
                                                                    p1 += (p1 ^ 572)
                                                                    _ = v70
                                                                } else {
                                                                    _ = len(v12)
                                                                }
                                                            }
                                                            var v71 int16 = -977
                                                            _ = v64
                                                            _ = v71
                                                        } else {
                                                            if false {
                                                                v72 := GenS11{F0: 9223372036854775807, F1: int16(p1)}
                                                                fmt.Println(v10, p1)
                                                                v73 := GenS11{F0: 1, F1: -793}
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                            var v74 int = 119
                                                            v74 += (v74 * fi22)
                                                            var v75 uint64 = v10
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                        if false {
                                                            v76 := GenS5{F0: 40}
                                                            _ = v76
                                                        }
                                                        switch "a" {
                                                        case "foo bar":
                                                            if len(v12) > 0 {
                                                                v12[0] = int(v10)
                                                            }
                                                        case "hello":
                                                            v10 = 4294967295
                                                        default:
                                                            switch p0 {
                                                            case -9223372036854775808:
                                                                p1 += (^(p1))
                                                            case 43:
                                                                if (!(!false)) {
                                                                    v77 := GenS12{F0: 42}
                                                                    var v78 int = int(p1)
                                                                    v12 = append(v12, 1)
                                                                    _ = v77
                                                                    _ = v78
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v25
                                            }
                                        case "foo bar":
                                            if len(v12) > 0 {
                                                v12[0] = 128
                                            }
                                        default:
                                            fmt.Println(p0)
                                        }
                                    }
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    v79 := GenS9{F0: "hello", F1: -1}
                                    _ = fi22
                                    _ = v79
                                }
                            }
                            v10 = v10
                        } else {
                            v80 := GenS13{F0: true}
                            _ = v80
                        }
                        _ = fi21
                    }
                    p1 += p1
                } else {
                    v10 -= (v10 % ((v10) | 1))
                }
                for fi81 := 0; fi81 < 5; fi81++ {
                    var v82 int16 = 32767
                    if true {
                        v83 := GenS7{F0: 1, F1: 42, F2: "\n\t"}
                        if false {
                            if true {
                                v12 = append(v12, (p0 % ((fi81) | 1)))
                                v84 := GenS14{F0: 137}
                                _ = v84
                            } else {
                                var v85 float32 = 0.1
                                var v86 uint32 = 42
                                switch 569 {
                                case 4294967295:
                                    for fi87 := 0; fi87 < 14; fi87++ {
                                        v86 += (v86 & 2147483648)
                                        v86 += (v86 & 42)
                                        if false {
                                            p0 += 669
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            v88 := GenS6{F0: (v10 * (v10 * v10)), F1: (^(p1)), F2: -1}
                                            _ = v88
                                        } else {
                                            if false {
                                                if true {
                                                    var v89 int = fi87
                                                    _ = v89
                                                } else {
                                                    p0 = p0
                                                    v90 := GenS15{F0: (v86 % ((v86) | 1)), F1: -1}
                                                    _ = v90
                                                }
                                            }
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            var v91 float32 = (-2.25 + float32(((-2.25 * 1.5) - 3.14159)))
                                            if false {
                                                var v92 int32 = -2147483648
                                                if false {
                                                    _ = len(v12)
                                                }
                                                _ = v92
                                            }
                                            _ = v91
                                        }
                                        _ = fi87
                                    }
                                }
                                _ = v85
                                _ = v86
                            }
                            if (!false) {
                                v10 = 960
                                if false {
                                    p0 = -93
                                    var v93 string = "foo bar"
                                    for fi94 := 0; fi94 < 15; fi94++ {
                                        v95 := GenS16{F0: (v93 != "a")}
                                        v11 -= (v11 / ((v11 * (v11 - (v11 + (((0.0 * float64(-2.25)) / float64(3.14159)) / 3.14159)))) - v11))
                                        v11 -= (-2.25 + v11)
                                        _ = fi94
                                        _ = v95
                                    }
                                    _ = v93
                                } else {
                                    switch fi81 {
                                    case 42:
                                        switch ("a" + "foo bar") {
                                        case "hello":
                                            v96 := GenS17{F0: "", F1: 315}
                                            _ = v96
                                        case "\n\t":
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                        }
                                    }
                                    if false {
                                        for fi97 := 0; fi97 < 17; fi97++ {
                                            var v98 uint16 = (p1 | (^(p1)))
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            _ = fi97
                                            _ = v98
                                        }
                                        for fi99 := 0; fi99 < 8; fi99++ {
                                            switch "\n\t" {
                                            case "a":
                                                v10 -= v10
                                            }
                                            _ = fi99
                                        }
                                    } else {
                                        var v100 uint64 = v10
                                        v11 -= 0.1
                                        v101 := GenS16{F0: (!true)}
                                        _ = v100
                                        _ = v101
                                    }
                                    if len(v12) > 0 {
                                        v12[0] = p0
                                    }
                                }
                                p0 -= fi81
                            }
                        }
                        _ = len(v12)
                        _ = v83
                    }
                    if len(v12) > 0 {
                        v12[0] = fi81
                    }
                    _ = fi81
                    _ = v82
                }
                var v102 int32 = 42
                _ = v102
            }
            v12 = append(v12, (p0 >> ((65535) & 63)))
        } else {
            v10 = v10
            var v103 uint32 = 2147483648
            _ = v103
        }
        _ = v12
    }
    p1 += p1
    var v104 string = "hello"
    _ = p0
    _ = p1
    _ = v10
    _ = v104
    _ = v11
    return -327
}

func main() {
    v1 := []int{0, 32768, 32767, -603}
    v2 := []int{665, 127, 0, -806}
    var v3 uint8 = 128
    var v4 uint8 = (v3 >> ((18446744073709551615) & 63))
    v4 -= (v3 * 255)
    v3 += 1
    var v105 uint = 2
    _ = len(v2)
    var v106 string = ""
    if len(v1) > 0 {
        v1[0] = 128
    }
    v2 = append(v2, 168)
    if (-128 < 0) {
        var v107 uint32 = uint32(v105)
        var v108 uint = 111
        v4 = v4
        _ = v107
        _ = v108
    }
    v106 = v106
    fmt.Println(v106, v3, v3)
    if (1 > 0) {
        _ = len(v2)
    }
    if len(v1) > 0 {
        v1[0] = 293
    }
    _ = len(v1)
    if false {
        for fi109 := 0; fi109 < 7; fi109++ {
            fmt.Println(v3, v4, v3)
            if (!(!false)) {
                var v110 int = fi109
                switch (fi109 | -867) {
                case 0:
                    v111 := GenS4{F0: v110, F1: 2147483647, F2: 2147483648}
                    _ = v111
                case 42:
                    switch (fi109 | v110) {
                    case -48:
                        for fi112 := 0; fi112 < 12; fi112++ {
                            if true {
                                var v113 int8 = -14
                                v114 := GenS10{F0: 0}
                                v113 += (v113 & (v113 / ((v113) | 1)))
                                _ = v113
                                _ = v114
                            }
                            var v115 uint32 = 638
                            if (-590 == -1) {
                                switch v106 {
                                case "":
                                    var v116 int16 = -370
                                    _ = v116
                                case "hello":
                                    fmt.Println(v3, v105, v106)
                                default:
                                    if false {
                                        v105 = 255
                                        fmt.Println(v110)
                                    } else {
                                        if false {
                                            v115 = (v115 * v115)
                                            fmt.Println(v3)
                                        }
                                        if (((-2.25 / (1e10 * -2.25)) == 0.0) || true) {
                                            v117 := GenS18{F0: ("" + v106), F1: -9223372036854775808, F2: ""}
                                            _ = v117
                                        }
                                    }
                                }
                            }
                            var v118 int32 = 1
                            fmt.Println(v3, fi109)
                            v106 = (v106 + v106)
                            var v119 int8 = 127
                            v120 := GenS19{F0: v119}
                            _ = fi112
                            _ = v115
                            _ = v118
                            _ = v119
                            _ = v120
                        }
                    case 348:
                        v121 := GenS17{F0: "a", F1: 0}
                        _ = v121
                    default:
                        if len(v2) > 0 {
                            v2[0] = v110
                        }
                    }
                case 4294967295:
                    v122 := GenS20{F0: true, F1: (v4 / ((v4) | 1))}
                    _ = v122
                default:
                    var v123 uint32 = 42
                    _ = v123
                }
                _ = v110
            }
            var v124 uint = v105
            _ = fi109
            _ = v124
        }
        var v125 string = ((("foo bar" + v106) + v106) + v106)
        var v126 float64 = float64((-2.25 - (((-2.25 * ((3.14159 / (1.5 / 1.0)) - (0.0 - 3.14159))) * float32(1.5)) / (((1.5 / -2.25) - 3.14159) / 0.1))))
        _ = v125
        _ = v126
    } else {
        if true {
            v1 = append(v1, 127)
            v1 = append(v1, 128)
        }
        var v127 int16 = 0
        var v128 float32 = (0.0 + 0.0)
        _ = v127
        _ = v128
    }
    if true {
        var v129 float64 = float64(((((float32(0.1) * float32(float64(-2.25))) / 1e10) + (1.5 + (float32(0.0) * (0.0 * 3.14159)))) + 0.0))
        v130 := GenS21{F0: v105}
        var v131 string = v106
        _ = v129
        _ = v130
        _ = v131
    } else {
        var v132 uint = (v105 / ((4294967295) | 1))
        for fi133 := 0; fi133 < 4; fi133++ {
            switch 127 {
            case 4294967295:
                var v134 float32 = (1e10 + 1.5)
                _ = v134
            }
            var v135 float64 = 1.5
            _ = fi133
            _ = v135
        }
        v136 := GenS5{F0: 0}
        _ = v132
        _ = v136
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v105
    _ = v106
    _ = v2
    _ = v3
    _ = v4
}
