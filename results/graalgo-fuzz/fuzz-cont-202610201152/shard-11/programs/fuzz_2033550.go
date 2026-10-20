package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 bool; F1 int64; F2 bool }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 int16; F2 string }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int16; F1 bool; F2 int }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 bool }
func genFn0(p0 int32, p1 uint, p2 uint) int {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int32 = (p0 & (p0 * -2147483648))
        var v9 string = "\n\t"
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 uint8 = 198
            switch fi7 {
            case 0:
                v12 := []int{(fi10 >> ((p1) & 63))}
                _ = v12
            case 506:
                v13 := []int{fi7}
                _ = v13
            case -1:
                v14 := []int{fi10, (fi7 / (((fi10 ^ -764)) | 1)), fi7, fi10}
                _ = v14
            default:
                if ((774 == 117) || (true || (492 <= int64(v11)))) {
                    var v15 bool = (4294967295 > 0)
                    _ = v15
                } else {
                    if false {
                        p0 = p0
                    } else {
                        if ((-2.25 - (((0.1 - -2.25) / 1.0) * (float32((1.5 - (0.1 + 3.14159))) + 1.5))) > 1.5) {
                            var v16 int32 = p0
                            v8 -= v8
                            if ((v8 > (v16 << ((p2) & 63))) && true) {
                                if (v9 == ("a" + "\n\t")) {
                                    v11 -= v11
                                    switch -9223372036854775808 {
                                    case 9223372036854775807:
                                        v17 := GenS1{F0: 1, F1: 0}
                                        _ = v17
                                    case 0:
                                        switch fi7 {
                                        case 32768:
                                            p1 = p2
                                        case 42:
                                            p0 += (v16 + p0)
                                        case -1:
                                            v16 = p0
                                        default:
                                            v11 = (v11 | v11)
                                        }
                                    }
                                }
                                var v18 int16 = 32767
                                _ = v18
                            } else {
                                if (((-9223372036854775808 > 508) && false) && (65535 == 1)) {
                                    v19 := GenS1{F0: -1, F1: -32768}
                                    _ = v19
                                }
                                switch ((((v9 + v9) + v9) + "hello") + v9) {
                                case "hello":
                                    var v20 bool = (127 < (fi10 * fi7))
                                    _ = v20
                                case "foo bar":
                                    fmt.Println(v16, p0, p0)
                                case "":
                                    v21 := []int{}
                                    _ = v21
                                }
                            }
                            _ = v16
                        }
                        if (true && (4294967295 <= 307)) {
                            p2 = 65535
                            if (!((p0 + p0) >= (^(v8)))) {
                                var v22 int64 = -518
                                var v23 string = (v9 + v9)
                                var v24 float64 = (0.0 + -2.25)
                                _ = v22
                                _ = v23
                                _ = v24
                            } else {
                                p1 = uint(v11)
                                v9 = "foo bar"
                                switch fi7 {
                                case -791:
                                    switch 428 {
                                    case -1:
                                        p0 = (p0 / ((p0) | 1))
                                    case -513:
                                        v9 = v9
                                    }
                                    if true {
                                        switch (v9 + "foo bar") {
                                        case "a":
                                            var v25 uint32 = 2147483648
                                            _ = v25
                                        default:
                                            v9 = (("hello" + ("\n\t" + v9)) + "\n\t")
                                        }
                                    }
                                    var v26 uint32 = 2
                                    _ = v26
                                default:
                                    fmt.Println(v11, v9, fi7)
                                }
                            }
                        }
                    }
                }
            }
            fmt.Println(v8)
            _ = fi10
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        v28 := GenS2{F0: (p0 - (p0 >> (((p2 / (((p1 % ((94) | 1))) | 1))) & 63))), F1: 701}
        v29 := GenS3{F0: 32767, F1: (p0 * (p0 ^ -1))}
        _ = fi27
        _ = v28
        _ = v29
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}
func genFn1(p0 int32, p1 int32, p2 int) int16 {
    fmt.Println(p0, p1, p0)
    v30 := GenS4{F0: false, F1: -214, F2: (!false)}
    for fi31 := 0; fi31 < 19; fi31++ {
        var v32 float64 = 0.0
        _ = fi31
        _ = v32
    }
    p2 = (p2 & genFn0((p0 ^ p1), 387, 18446744073709551615))
    p2 += -972
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return 339
}

func main() {
    v1 := []string{}
    var v2 int16 = -287
    var v3 uint32 = 65535
    var v4 float64 = 0.0
    var v5 bool = true
    var v6 float64 = v4
    for fi33 := 0; fi33 < 20; fi33++ {
        var v34 bool = v5
        v5 = (v6 != v6)
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        switch 9223372036854775807 {
        case 32768:
            fmt.Println(fi35, v4, v3)
        case 127:
            v6 -= (v4 + (v6 + (v6 - v4)))
        }
        _ = fi35
    }
    v36 := GenS4{F0: v5, F1: int64(v2), F2: true}
    switch genFn0(-687, 105, 42) {
    case 1:
        if true {
            if (true || (!v5)) {
                v37 := GenS5{F0: "foo bar", F1: -2147483648}
                for fi38 := 0; fi38 < 2; fi38++ {
                    v39 := GenS2{F0: 418, F1: 4294967295}
                    fmt.Println(v6, fi38, v3)
                    switch 0 {
                    case 127:
                        if v5 {
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                        }
                    case 4294967295:
                        v3 -= v3
                    }
                    _ = fi38
                    _ = v39
                }
                if (0.1 != 0.1) {
                    v40 := GenS3{F0: -32768, F1: 55}
                    if v5 {
                        v2 = (-(v2))
                    }
                    _ = v40
                } else {
                    if v5 {
                        v5 = v5
                    }
                }
                _ = v37
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    case 514:
        for fi41 := 0; fi41 < 16; fi41++ {
            if v5 {
                switch 32767 {
                case -1:
                    v6 = ((0.0 + v6) * 0.1)
                default:
                    if (-2147483648 == 923) {
                        v1 = append(v1, ("hello" + ("" + ("foo bar" + "a"))))
                    } else {
                        if v5 {
                            if (42 <= 0) {
                                if (v5 || false) {
                                    _ = len(v1)
                                    var v42 float64 = v6
                                    _ = v42
                                } else {
                                    fmt.Println(v4)
                                }
                                switch (("" + ("" + "hello")) + "\n\t") {
                                case "a":
                                    var v43 uint64 = 0
                                    _ = v43
                                case "foo bar":
                                    switch ("" + "foo bar") {
                                    case "\n\t":
                                        var v44 uint = 1
                                        _ = v44
                                    case "":
                                        v3 = (v3 ^ (v3 ^ 1))
                                    case "foo bar":
                                        _ = len(v1)
                                    default:
                                        if v5 {
                                            v4 = 0.0
                                            v6 = 1.5
                                            var v45 uint8 = 128
                                            _ = v45
                                        }
                                    }
                                case "hello":
                                    if v5 {
                                        if true {
                                            switch (fi41 >> ((2) & 63)) {
                                            case -536:
                                                v2 = v2
                                            case 9223372036854775807:
                                                if (v2 == (v2 & 32767)) {
                                                    if (1 <= 42) {
                                                        v4 -= -2.25
                                                        for fi46 := 0; fi46 < 16; fi46++ {
                                                            var v47 float32 = 0.1
                                                            v3 = v3
                                                            _ = fi46
                                                            _ = v47
                                                        }
                                                        _ = len(v1)
                                                    }
                                                    var v48 string = "\n\t"
                                                    v6 += v6
                                                    _ = v48
                                                } else {
                                                    v49 := GenS6{F0: "\n\t"}
                                                    _ = v49
                                                }
                                            }
                                        }
                                        for fi50 := 0; fi50 < 10; fi50++ {
                                            v51 := GenS6{F0: "hello"}
                                            _ = fi50
                                            _ = v51
                                        }
                                        for fi52 := 0; fi52 < 7; fi52++ {
                                            var v53 uint32 = v3
                                            v54 := GenS4{F0: true, F1: 9223372036854775807, F2: true}
                                            _ = fi52
                                            _ = v53
                                            _ = v54
                                        }
                                    }
                                }
                            } else {
                                fmt.Println(fi41, v5, v4)
                            }
                        }
                    }
                }
                var v55 string = ("hello" + ("hello" + "a"))
                v56 := GenS6{F0: v55}
                _ = v55
                _ = v56
            } else {
                v5 = v5
            }
            _ = fi41
        }
    case 127:
        if ((("a" + "hello") + "") < "hello") {
            switch "" {
            case "foo bar":
                for fi57 := 0; fi57 < 17; fi57++ {
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    _ = fi57
                }
            case "":
                var v58 int = -739
                _ = v58
            default:
                if v5 {
                    v3 = (v3 * v3)
                    var v59 int16 = (v2 << ((411) & 63))
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v59
                } else {
                    var v60 int16 = 1
                    var v61 int = genFn0(810, 18446744073709551615, 18446744073709551615)
                    var v62 int64 = 9223372036854775807
                    _ = v60
                    _ = v61
                    _ = v62
                }
            }
        }
    default:
        v63 := GenS7{F0: 0, F1: v2, F2: "foo bar"}
        _ = v63
    }
    v64 := GenS8{F0: 18446744073709551615}
    switch "a" {
    case "":
        var v65 uint8 = 128
        _ = v65
    case "\n\t":
        _ = len(v1)
    case "hello":
        fmt.Println(v2)
    }
    v4 += (-2.25 - 0.1)
    var v66 float32 = -2.25
    v67 := GenS4{F0: (18446744073709551615 <= 0), F1: 0, F2: v5}
    v68 := GenS8{F0: 761}
    v5 = v5
    if (!((genFn0(678, uint(v3), 18446744073709551615) != -222) || true)) {
        v6 += 0.1
        switch genFn0(42, 472, 556) {
        case 0:
            var v69 uint32 = 39
            _ = v69
        case 4294967295:
            v66 = v66
        default:
            if v5 {
                if ((0.1 + float64(0.1)) != v6) {
                    v1 = append(v1, "hello")
                    for fi70 := 0; fi70 < 19; fi70++ {
                        for fi71 := 0; fi71 < 9; fi71++ {
                            var v72 uint32 = 2147483648
                            switch (("\n\t" + "a") + "") {
                            case "a":
                                v66 -= 1.5
                            case "\n\t":
                                v1 = append(v1, "")
                            case "foo bar":
                                if true {
                                    switch "" {
                                    case "hello":
                                        if true {
                                            v73 := GenS1{F0: 1, F1: 32767}
                                            var v74 int = 654
                                            v66 += (float32(v6) / v66)
                                            _ = v73
                                            _ = v74
                                        } else {
                                            fmt.Println(v3, v2)
                                        }
                                    case "":
                                        if v5 {
                                            if ((42 == 0) && ((false || v5) && (v5 || (!false)))) {
                                                v75 := GenS7{F0: 0, F1: (v2 - v2), F2: "a"}
                                                v1 = append(v1, "")
                                                v76 := GenS9{F0: v2, F1: (!(("hello" + "\n\t") <= (("" + "foo bar") + "foo bar"))), F2: genFn0(42, 1, 18446744073709551615)}
                                                _ = v75
                                                _ = v76
                                            } else {
                                                var v77 float32 = v66
                                                v4 = (3.14159 + (((v6 * (v4 / (v6 - ((v4 - v4) - 1.5)))) * (-2.25 + v6)) * v6))
                                                _ = v77
                                            }
                                            v72 -= v3
                                        } else {
                                            var v78 int = fi71
                                            v79 := GenS10{F0: 384}
                                            switch (("hello" + "\n\t") + "hello") {
                                            case "":
                                                v66 += 3.14159
                                            default:
                                                if len(v1) > 0 {
                                                    v1[0] = "\n\t"
                                                }
                                            }
                                            _ = v78
                                            _ = v79
                                        }
                                    default:
                                        if v5 {
                                            v72 = (v3 % (((v3 * (^(v3)))) | 1))
                                            if len(v1) > 0 {
                                                v1[0] = ""
                                            }
                                        }
                                    }
                                    v3 -= v3
                                    fmt.Println(v66, fi70)
                                }
                            default:
                                v4 = 0.1
                                var v80 float32 = (1.5 * v66)
                                fmt.Println(fi71, v4)
                                _ = v80
                            }
                            if (!v5) {
                                fmt.Println(v5, v72, v3)
                            }
                            _ = fi71
                            _ = v72
                        }
                        switch (fi70 << ((2) & 63)) {
                        case 42:
                            fmt.Println(v2)
                        }
                        _ = fi70
                    }
                    for fi81 := 0; fi81 < 7; fi81++ {
                        v3 += v3
                        var v82 int16 = 781
                        fmt.Println(v3, v82)
                        _ = fi81
                        _ = v82
                    }
                } else {
                    var v83 bool = v5
                    for fi84 := 0; fi84 < 13; fi84++ {
                        for fi85 := 0; fi85 < 6; fi85++ {
                            var v86 int64 = 0
                            v87 := GenS3{F0: v2, F1: 0}
                            _ = fi85
                            _ = v86
                            _ = v87
                        }
                        switch (("foo bar" + ("" + ("hello" + (("foo bar" + "hello") + "foo bar")))) + "") {
                        case "\n\t":
                            for fi88 := 0; fi88 < 12; fi88++ {
                                var v89 uint = 2
                                _ = len(v1)
                                fmt.Println(fi84, v83)
                                _ = fi88
                                _ = v89
                            }
                        }
                        _ = fi84
                    }
                    _ = len(v1)
                    _ = v83
                }
            }
        }
        v90 := GenS10{F0: 255}
        _ = v90
    }
    v6 -= ((v6 * float64(v66)) / 1e10)
    v91 := GenS5{F0: "hello", F1: genFn0(1, 18446744073709551615, 65535)}
    v92 := GenS11{F0: v2}
    for fi93 := 0; fi93 < 13; fi93++ {
        v94 := GenS12{F0: v5}
        _ = fi93
        _ = v94
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v66
    _ = v67
    _ = v68
    _ = v91
    _ = v92
}
