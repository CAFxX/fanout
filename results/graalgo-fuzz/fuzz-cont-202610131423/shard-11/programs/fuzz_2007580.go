package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 string }
type GenS2 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 uint16; F2 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32; F1 int8; F2 uint }
type GenS8 struct { F0 uint16; F1 uint; F2 int8 }
type GenS9 struct { F0 bool; F1 int16; F2 bool }
type GenS10 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS11 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint64 }
func genFn0(p0 int) uint16 {
    p0 -= p0
    v6 := GenS1{F0: ((p0 | 4294967295) != (p0 / ((p0) | 1))), F1: 128, F2: "a"}
    p0 = p0
    _ = p0
    _ = v6
    return 65535
}

func main() {
    var v1 uint16 = 789
    v2 := []string{("foo bar" + ((("" + ("a" + "a")) + "\n\t") + ("hello" + (("hello" + ("hello" + "a")) + (("" + ("" + "")) + ("foo bar" + "foo bar")))))), "hello", ""}
    var v3 uint8 = 128
    var v4 float32 = ((0.1 * 1.5) + 1.5)
    var v5 uint64 = 1
    v2 = append(v2, "foo bar")
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 float64 = 1e10
        _ = fi7
        _ = v8
    }
    fmt.Println(v5, v5)
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            v5 += 18446744073709551615
            if ((!true) && (!true)) {
                fmt.Println(v1, v4, v3)
                if (!false) {
                    _ = len(v2)
                    v5 = (^(v5))
                    v1 -= v1
                    if true {
                        v5 = v5
                        if (128 < (^(fi10))) {
                            if (127 != -79) {
                                var v11 int64 = 42
                                v4 = v4
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                                _ = v11
                            }
                            if (!((!false) || (!(175 != -280)))) {
                                var v12 int8 = 100
                                _ = v12
                            } else {
                                v13 := GenS1{F0: false, F1: (-(fi9)), F2: (("hello" + ("\n\t" + "\n\t")) + "foo bar")}
                                v5 += v5
                                v14 := GenS1{F0: true, F1: fi9, F2: "foo bar"}
                                _ = v13
                                _ = v14
                            }
                            var v15 uint8 = (^(v3))
                            _ = v15
                        } else {
                            var v16 bool = true
                            fmt.Println(v16, v4)
                            _ = v16
                        }
                    } else {
                        v2 = append(v2, (("" + ("\n\t" + "a")) + ("foo bar" + "")))
                        var v17 bool = true
                        _ = v17
                    }
                }
            } else {
                v2 = append(v2, ("hello" + "\n\t"))
                switch "\n\t" {
                case "hello":
                    if true {
                        switch "foo bar" {
                        case "":
                            _ = len(v2)
                        case "\n\t":
                            if len(v2) > 0 {
                                v2[0] = ""
                            }
                        }
                    } else {
                        v18 := GenS1{F0: (!(!(!true))), F1: fi9, F2: ("hello" + "hello")}
                        v3 -= (v3 % (((v3 / ((v3) | 1))) | 1))
                        _ = v18
                    }
                default:
                    switch ("" + "") {
                    case "foo bar":
                        v19 := GenS1{F0: true, F1: (fi10 << ((1) & 63)), F2: ("hello" + "a")}
                        switch "a" {
                        case "foo bar":
                            if (v1 < 838) {
                                var v20 bool = false
                                switch (^(fi9)) {
                                case -582:
                                    var v21 uint32 = 0
                                    _ = v21
                                case -9223372036854775808:
                                    v22 := GenS1{F0: (v4 >= v4), F1: fi10, F2: (("" + "foo bar") + "\n\t")}
                                    if true {
                                        v2 = append(v2, "a")
                                    } else {
                                        v23 := GenS2{F0: 255, F1: false, F2: 65535}
                                        switch "" {
                                        case "foo bar":
                                            if v20 {
                                                v3 = (^(v3))
                                                v5 += (v5 % ((v5) | 1))
                                                var v24 string = "hello"
                                                _ = v24
                                            }
                                        default:
                                            fmt.Println(fi9, v3, v5)
                                        }
                                        switch "" {
                                        case "":
                                            fmt.Println(fi9)
                                        case "a":
                                            if ((true || (true || ((((v20 && v20) || ((fi9 ^ fi9) < 127)) || (-1 > 2147483647)) || true))) && (uint16(v3) != v1)) {
                                                if v20 {
                                                    if v20 {
                                                        v20 = v20
                                                        var v25 uint32 = 856
                                                        _ = v25
                                                    } else {
                                                        if (true && true) {
                                                            v5 = v5
                                                        }
                                                        v5 += (v5 | v5)
                                                    }
                                                    _ = len(v2)
                                                } else {
                                                    if (261 < 42) {
                                                        v5 = (v5 << ((0) & 63))
                                                        var v26 string = "hello"
                                                        var v27 uint16 = (v1 + 0)
                                                        _ = v26
                                                        _ = v27
                                                    } else {
                                                        v3 += (^(v3))
                                                        switch "hello" {
                                                        case "hello":
                                                            var v28 int8 = 42
                                                            _ = v28
                                                        case "":
                                                            switch fi10 {
                                                            case 530:
                                                                var v29 bool = (((0.0 + ((-2.25 * (3.14159 - (1e10 + (0.0 / (0.1 / float64(v4)))))) + (1e10 * ((((1.5 + -2.25) * (float64(v4) / ((3.14159 - float64(0.0)) - 1.5))) - 1.5) / float64((0.0 - (v4 * v4))))))) >= (3.14159 / 3.14159)) && true)
                                                                _ = v29
                                                            case 9223372036854775807:
                                                                v1 = (v1 * v1)
                                                            default:
                                                                if ((fi9 / ((42) | 1)) == -2147483648) {
                                                                    v30 := GenS1{F0: v20, F1: 1, F2: "\n\t"}
                                                                    v31 := GenS1{F0: v20, F1: 611, F2: "\n\t"}
                                                                    _ = v30
                                                                    _ = v31
                                                                }
                                                            }
                                                        }
                                                        switch ("\n\t" + "foo bar") {
                                                        case "":
                                                            if true {
                                                                if false {
                                                                    var v32 uint64 = uint64(v1)
                                                                    if v20 {
                                                                        fmt.Println(v1)
                                                                    } else {
                                                                        v33 := GenS1{F0: (!("a" >= "")), F1: (fi9 | (fi10 ^ (fi9 - 2147483647))), F2: ""}
                                                                        _ = v33
                                                                    }
                                                                    var v34 uint16 = uint16(v32)
                                                                    v2 = append(v2, (("hello" + "") + ""))
                                                                    _ = v32
                                                                    _ = v34
                                                                }
                                                                v3 = v3
                                                            }
                                                            var v35 uint = uint(v1)
                                                            _ = v35
                                                        }
                                                        if v20 {
                                                            v1 = (^(v1))
                                                            v4 += v4
                                                            v2 = append(v2, "foo bar")
                                                        }
                                                        var v36 int64 = -9223372036854775808
                                                        _ = v36
                                                    }
                                                    v20 = v20
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = "\n\t"
                                                }
                                                v37 := GenS3{F0: v5, F1: 795}
                                                _ = v37
                                            } else {
                                                _ = len(v2)
                                                var v38 float32 = float32(((-2.25 / (1.5 / (1.5 - (3.14159 + 1e10)))) + (((((float64(float32(1e10)) + 0.0) / 1.5) - -2.25) * float64(float32(1.5))) - 0.0)))
                                                if true {
                                                    var v39 int8 = -128
                                                    switch (fi9 >> ((1) & 63)) {
                                                    case 128:
                                                        var v40 int64 = -1
                                                        _ = v40
                                                    }
                                                    _ = v39
                                                }
                                                _ = v38
                                            }
                                        default:
                                            var v41 uint16 = 32767
                                            _ = v41
                                        }
                                        _ = v23
                                    }
                                    v1 = (^(v1))
                                    switch ("" + "a") {
                                    case "\n\t":
                                        switch "\n\t" {
                                        case "\n\t":
                                            switch "hello" {
                                            case "hello":
                                                var v42 uint16 = v1
                                                _ = v42
                                            case "a":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            case "\n\t":
                                                _ = len(v2)
                                                var v43 int32 = -913
                                                _ = v43
                                            default:
                                                v44 := GenS1{F0: ((((v4 * (v4 + v4)) + ((0.0 - 1.5) * v4)) * (-2.25 / v4)) > 0.1), F1: fi10, F2: ("a" + "foo bar")}
                                                _ = v44
                                            }
                                        case "a":
                                            v1 = (v1 & (v1 << ((77) & 63)))
                                        case "foo bar":
                                            if len(v2) > 0 {
                                                v2[0] = ("\n\t" + "hello")
                                            }
                                            var v45 uint = 625
                                            if (int32(v45) == 0) {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v45 = v45
                                            }
                                            v46 := GenS2{F0: v3, F1: (0 < -40), F2: 0}
                                            switch fi10 {
                                            case -9223372036854775808:
                                                var v47 int8 = -5
                                                _ = v47
                                            }
                                            var v48 int32 = -956
                                            _ = v45
                                            _ = v46
                                            _ = v48
                                        default:
                                            var v49 uint8 = 42
                                            _ = v49
                                        }
                                        v50 := GenS4{F0: v5}
                                        _ = v50
                                    default:
                                        switch "a" {
                                        case "":
                                            v1 = (v1 * (v1 >> ((18446744073709551615) & 63)))
                                        }
                                        switch -1000 {
                                        case 32768:
                                            var v51 int32 = -595
                                            _ = v51
                                        case 9223372036854775807:
                                            switch "foo bar" {
                                            case "":
                                                var v52 uint = 639
                                                _ = v52
                                            case "\n\t":
                                                switch (fi9 >> ((1) & 63)) {
                                                case 9223372036854775807:
                                                    switch (fi10 % ((fi10) | 1)) {
                                                    case 869:
                                                        v1 -= (^(v1))
                                                    default:
                                                        if len(v2) > 0 {
                                                            v2[0] = ""
                                                        }
                                                    }
                                                case 205:
                                                    var v53 string = ""
                                                    v54 := GenS3{F0: 4294967295, F1: -789}
                                                    v55 := GenS4{F0: (v5 + 0)}
                                                    _ = v53
                                                    _ = v54
                                                    _ = v55
                                                default:
                                                    v56 := GenS5{F0: v20, F1: v1, F2: 1}
                                                    _ = v56
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v57 uint64 = (v5 & (v5 * (v5 * (v5 - 42))))
                                                _ = v57
                                            case "foo bar":
                                                v1 = v1
                                            }
                                        case 821:
                                            if (-9223372036854775808 <= 42) {
                                                v58 := GenS6{F0: true}
                                                fmt.Println(v5, fi9, fi10)
                                                _ = v58
                                            } else {
                                                v59 := GenS7{F0: -1, F1: 0, F2: 4294967295}
                                                fmt.Println(v20, fi10, v4)
                                                _ = v59
                                            }
                                        }
                                        v20 = true
                                    }
                                    var v60 int = -2147483648
                                    _ = v22
                                    _ = v60
                                default:
                                    v3 = 2
                                }
                                fmt.Println(fi10, v20)
                                _ = v20
                            } else {
                                _ = len(v2)
                                switch (fi9 >> ((261) & 63)) {
                                case 1:
                                    if ((v1 / ((2) | 1)) >= v1) {
                                        v4 = v4
                                        var v61 uint32 = 532
                                        var v62 uint8 = (v3 * v3)
                                        _ = v61
                                        _ = v62
                                    } else {
                                        if ((true || (42 < 16)) && true) {
                                            v1 += genFn0(2147483647)
                                            if (100 == 1) {
                                                switch (fi10 + 0) {
                                                case 4294967295:
                                                    if (v1 >= v1) {
                                                        v3 += v3
                                                        v63 := GenS7{F0: -2147483648, F1: -1, F2: 18446744073709551615}
                                                        fmt.Println(fi10, fi9)
                                                        _ = v63
                                                    } else {
                                                        switch ("hello" + "\n\t") {
                                                        case "a":
                                                            _ = len(v2)
                                                        case "hello":
                                                            v64 := GenS8{F0: v1, F1: 255, F2: -128}
                                                            _ = v64
                                                        case "foo bar":
                                                            v2 = append(v2, "hello")
                                                        }
                                                        v4 = float32(1.5)
                                                        var v65 bool = ("\n\t" > (("foo bar" + "a") + ("\n\t" + "a")))
                                                        _ = v65
                                                    }
                                                case 462:
                                                    var v66 uint16 = v1
                                                    _ = v66
                                                default:
                                                    v1 -= v1
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                if false {
                                                    var v67 bool = (!(false || ((v3 / ((v3) | 1)) < 42)))
                                                    if false {
                                                        var v68 int32 = 51
                                                        _ = v68
                                                    } else {
                                                        switch "hello" {
                                                        case "\n\t":
                                                            var v69 uint32 = 65535
                                                            _ = v69
                                                        default:
                                                            switch -9223372036854775808 {
                                                            case 0:
                                                                v2 = append(v2, "")
                                                            case 127:
                                                                v5 -= v5
                                                            default:
                                                                if (uint16(v5) < 32767) {
                                                                    v2 = append(v2, "a")
                                                                    var v70 float32 = v4
                                                                    _ = v70
                                                                } else {
                                                                    v5 -= (v5 ^ v5)
                                                                    if (v67 || v67) {
                                                                        var v71 float32 = v4
                                                                        _ = v71
                                                                    }
                                                                    v72 := GenS9{F0: ((2147483648 > 2147483647) && false), F1: 42, F2: v67}
                                                                    if v67 {
                                                                        if v67 {
                                                                            var v73 int = 9223372036854775807
                                                                            var v74 float64 = -2.25
                                                                            fmt.Println(v74, fi9, v67)
                                                                            _ = v73
                                                                            _ = v74
                                                                        }
                                                                        v67 = v67
                                                                    }
                                                                    _ = v72
                                                                }
                                                            }
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = "a"
                                                        }
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    _ = v67
                                                } else {
                                                    if true {
                                                        v1 -= v1
                                                    } else {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v75 uint = 349
                                                            _ = v75
                                                        case "foo bar":
                                                            if true {
                                                                v4 = (0.1 * 1e10)
                                                                v1 -= uint16(v5)
                                                                var v76 uint16 = (v1 % ((v1) | 1))
                                                                v3 = 127
                                                                if true {
                                                                    fmt.Println(v1)
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    v2 = append(v2, ("a" + ("a" + "\n\t")))
                                                                    v76 = (v1 & v1)
                                                                } else {
                                                                    var v77 bool = ((v5 < uint64(v1)) && ((!((("foo bar" + "a") + ("hello" + "foo bar")) > ("" + ("\n\t" + "")))) && (true || true)))
                                                                    v76 = (v76 ^ 255)
                                                                    _ = v77
                                                                }
                                                                _ = v76
                                                            }
                                                        }
                                                        v78 := GenS8{F0: v1, F1: 800, F2: -128}
                                                        _ = v78
                                                    }
                                                    switch (fi10 / (((fi9 & (fi10 & 1))) | 1)) {
                                                    case 9223372036854775807:
                                                        var v79 int32 = -456
                                                        _ = v79
                                                    case -1:
                                                        fmt.Println(fi10, fi9, fi10)
                                                    case 42:
                                                        switch fi9 {
                                                        case 127:
                                                            var v80 uint64 = 2
                                                            _ = v80
                                                        case 32768:
                                                            fmt.Println(v3, v1, v1)
                                                        case -1:
                                                            if false {
                                                                v4 -= v4
                                                                v5 = v5
                                                            }
                                                        }
                                                    }
                                                    var v81 uint16 = 255
                                                    _ = v81
                                                }
                                            }
                                            if ((fi9 ^ int(v3)) < fi10) {
                                                v5 -= (v5 & v5)
                                                v1 = genFn0((fi10 << ((0) & 63)))
                                                v82 := GenS10{F0: 65535, F1: uint8(v5), F2: v5}
                                                var v83 string = "\n\t"
                                                v1 += v1
                                                _ = v82
                                                _ = v83
                                            }
                                        } else {
                                            v84 := GenS11{F0: true, F1: uint16(v5), F2: (^(v1))}
                                            if (!(!(v4 > v4))) {
                                                var v85 int = (fi10 | fi9)
                                                _ = v85
                                            }
                                            switch "\n\t" {
                                            case "a":
                                                if false {
                                                    if ((fi9 != (-(fi10))) || ((1e10 - ((float64(v4) + 1e10) / (0.0 - (0.1 + -2.25)))) > 1.5)) {
                                                        if false {
                                                            v5 = v5
                                                            v5 -= v5
                                                            var v86 int16 = 42
                                                            _ = v86
                                                        } else {
                                                            var v87 float64 = (0.1 + 1.5)
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                if len(v2) > 0 {
                                                                    v2[0] = ("foo bar" + "a")
                                                                }
                                                            case "a":
                                                                switch fi9 {
                                                                case 4294967295:
                                                                    v88 := GenS5{F0: (42 != 4294967295), F1: (v1 - 0), F2: 1}
                                                                    _ = v88
                                                                default:
                                                                    v89 := GenS4{F0: (v5 | (v5 / (((v5 % (((v5 + 365)) | 1))) | 1)))}
                                                                    _ = v89
                                                                }
                                                            default:
                                                                var v90 int32 = -2147483648
                                                                _ = v90
                                                            }
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                switch fi10 {
                                                                case 4294967295:
                                                                    var v91 uint = 1
                                                                    _ = v91
                                                                case 0:
                                                                    v92 := GenS12{F0: "hello"}
                                                                    _ = v92
                                                                case -252:
                                                                    if (("hello" + "a") >= ((("" + "foo bar") + "\n\t") + ("foo bar" + "a"))) {
                                                                        if ((fi9 + fi9) >= fi10) {
                                                                            if (!false) {
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    v1 += v1
                                                                                }
                                                                                v87 -= (v87 + float64((0.1 / (-2.25 + -2.25))))
                                                                                v93 := GenS11{F0: ((^(fi10)) <= 609), F1: v1, F2: 65535}
                                                                                _ = v93
                                                                            } else {
                                                                                _ = len(v2)
                                                                            }
                                                                            switch fi9 {
                                                                            case -143:
                                                                                v4 += v4
                                                                            case 9223372036854775807:
                                                                                v1 = v1
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v94 uint8 = v3
                                                                        v95 := GenS13{F0: v5}
                                                                        _ = v94
                                                                        _ = v95
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                if false {
                                                                    v4 += v4
                                                                    if false {
                                                                        v2 = append(v2, "a")
                                                                        switch -2147483648 {
                                                                        case 128:
                                                                            switch (fi10 - (fi9 << ((187) & 63))) {
                                                                            case 32767:
                                                                                if false {
                                                                                    var v96 float32 = 0.1
                                                                                    var v97 bool = (v1 == genFn0(fi10))
                                                                                    v98 := GenS9{F0: (((0 > -1) || (!v97)) || v97), F1: 1, F2: true}
                                                                                    _ = v96
                                                                                    _ = v97
                                                                                    _ = v98
                                                                                }
                                                                            default:
                                                                                if true {
                                                                                    if false {
                                                                                        switch (-(fi9)) {
                                                                                        case 127:
                                                                                            var v99 int32 = -519
                                                                                            _ = v99
                                                                                        default:
                                                                                            fmt.Println(v1, fi10, v5)
                                                                                        }
                                                                                    }
                                                                                    switch ("hello" + "") {
                                                                                    case "hello":
                                                                                        var v100 uint64 = (v5 << ((117) & 63))
                                                                                        _ = v100
                                                                                    case "":
                                                                                        var v101 int16 = -316
                                                                                        _ = v101
                                                                                    default:
                                                                                        var v102 int16 = 1000
                                                                                        _ = v102
                                                                                    }
                                                                                } else {
                                                                                    var v103 int8 = 127
                                                                                    _ = v103
                                                                                }
                                                                            }
                                                                        case 127:
                                                                            switch fi10 {
                                                                            case 266:
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            case 356:
                                                                                v104 := GenS7{F0: 121, F1: 1, F2: 437}
                                                                                _ = v104
                                                                            case -9223372036854775808:
                                                                                if true {
                                                                                    var v105 int = (fi9 | (^(fi10)))
                                                                                    _ = v105
                                                                                } else {
                                                                                    v106 := GenS13{F0: (^(v5))}
                                                                                    fmt.Println(v4, fi9, v3)
                                                                                    fmt.Println(fi9, v87)
                                                                                    _ = v106
                                                                                }
                                                                            default:
                                                                                var v107 float32 = 1.5
                                                                                _ = v107
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v108 int = -1
                                                                        v87 = v87
                                                                        _ = v108
                                                                    }
                                                                }
                                                            case "a":
                                                                v109 := GenS4{F0: 42}
                                                                _ = v109
                                                            }
                                                            _ = v87
                                                        }
                                                        var v110 uint8 = 2
                                                        v4 = -2.25
                                                        _ = v110
                                                    } else {
                                                        if true {
                                                            fmt.Println(v4)
                                                            if len(v2) > 0 {
                                                                v2[0] = ("hello" + "")
                                                            }
                                                        } else {
                                                            var v111 uint = 33
                                                            _ = v111
                                                        }
                                                        var v112 bool = false
                                                        _ = v112
                                                    }
                                                }
                                                v5 = (v5 & v5)
                                                v1 += (v1 & 32767)
                                            case "hello":
                                                var v113 bool = false
                                                _ = v113
                                            }
                                            _ = v84
                                        }
                                        var v114 uint8 = v3
                                        v115 := GenS1{F0: false, F1: int(v1), F2: "foo bar"}
                                        _ = v114
                                        _ = v115
                                    }
                                case 418:
                                    v1 = (v1 >> ((uint(v1)) & 63))
                                case 0:
                                    var v116 bool = (false || true)
                                    _ = v116
                                default:
                                    v1 += (v1 | 2)
                                }
                            }
                        case "":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case "\n\t":
                            v3 += v3
                        default:
                            v4 += 3.14159
                        }
                        v2 = append(v2, "a")
                        var v117 uint8 = v3
                        _ = v117
                        _ = v19
                    case "a":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    var v118 int8 = 127
                    var v119 bool = true
                    _ = v118
                    _ = v119
                }
                _ = len(v2)
            }
            _ = fi10
        }
        for fi120 := 0; fi120 < 7; fi120++ {
            var v121 uint8 = 1
            _ = fi120
            _ = v121
        }
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        _ = fi9
    }
    var v122 int64 = 1
    _ = len(v2)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v123 int64 = v122
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v124 string = "hello"
    v125 := GenS5{F0: ((v3 - v3) <= (^(v3))), F1: v1, F2: 9223372036854775807}
    v126 := GenS3{F0: (v5 / ((v5) | 1)), F1: (-(v122))}
    switch -442 {
    case 42:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 128:
        v4 -= (v4 + (v4 * 0.0))
    case -853:
        switch 42 {
        case 32768:
            v1 = (v1 * v1)
        case -199:
            var v127 uint32 = 65535
            _ = v127
        }
    }
    var v128 uint = uint(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v122:", v122)
    fmt.Println("v123:", v123)
    fmt.Println("v124:", v124)
    fmt.Println("v128:", v128)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v122
    _ = v123
    _ = v124
    _ = v125
    _ = v126
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
