package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 string; F2 string }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 uint32; F1 int8 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint16; F1 uint16; F2 uint }
type GenS10 struct { F0 int8; F1 int16 }
type GenS11 struct { F0 int64; F1 int32 }
type GenS12 struct { F0 int; F1 int64 }
type GenS13 struct { F0 int32; F1 uint }
type GenS14 struct { F0 bool; F1 string }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 string }
type GenS17 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS18 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 uint8; F1 int32; F2 int32 }
type GenS21 struct { F0 int16; F1 uint32; F2 int16 }
type GenS22 struct { F0 int8 }
type GenS23 struct { F0 int8; F1 uint8 }
type GenS24 struct { F0 int16; F1 int32; F2 int8 }

func main() {
    var v1 uint16 = 255
    var v2 int64 = 9223372036854775807
    v3 := []string{}
    var v4 int16 = 1000
    var v5 float32 = float32(float64(float32(0.1)))
    var v6 uint = 4294967295
    var v7 int16 = v4
    var v8 uint8 = uint8(v6)
    var v9 int16 = (v7 / ((v7) | 1))
    v10 := GenS1{F0: ("" + "\n\t"), F1: ((("" + "") + "") + "\n\t")}
    if len(v3) > 0 {
        v3[0] = ""
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(fi11, fi11)
        v12 := GenS1{F0: ("hello" + "foo bar"), F1: "hello"}
        var v13 int8 = 0
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 bool = false
    switch ("a" + (("hello" + "a") + "foo bar")) {
    case "a":
        var v15 float64 = ((0.1 / float64(float32(((((3.14159 + (float64(1e10) - 3.14159)) + ((-2.25 * -2.25) * float64(float32(((1.5 / -2.25) + 0.1))))) / 3.14159) / 1.5)))) + (0.1 / 1e10))
        _ = v15
    case "\n\t":
        switch -1 {
        case 128:
            if v14 {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v16 := GenS1{F0: "a", F1: ""}
                v9 += v4
                _ = v16
            }
        }
    default:
        if true {
            if (2147483647 == -2147483648) {
                switch 9223372036854775807 {
                case 361:
                    v2 += v2
                }
            }
        } else {
            v4 = 177
            var v17 uint = 255
            _ = v17
        }
    }
    fmt.Println(v7, v7, v6)
    v6 = 2
    if v14 {
        fmt.Println(v2, v9)
    } else {
        var v18 int = 1
        v19 := GenS1{F0: "", F1: ((("hello" + ("a" + ("foo bar" + "\n\t"))) + ("hello" + "\n\t")) + "")}
        switch "foo bar" {
        case "\n\t":
            for fi20 := 0; fi20 < 20; fi20++ {
                if true {
                    switch ("" + "hello") {
                    case "":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "foo bar":
                        fmt.Println(v5, v1)
                    case "hello":
                        var v21 int16 = (v7 & v4)
                        _ = v21
                    default:
                        v3 = append(v3, ("foo bar" + "foo bar"))
                    }
                } else {
                    var v22 int32 = -1
                    v23 := GenS1{F0: "foo bar", F1: (("\n\t" + "hello") + "a")}
                    _ = v22
                    _ = v23
                }
                v24 := GenS1{F0: ("hello" + ((("\n\t" + "\n\t") + "a") + ((("foo bar" + ("foo bar" + "\n\t")) + "") + ""))), F1: ((((("hello" + ((("" + "") + ("" + ("a" + ("\n\t" + "hello")))) + "foo bar")) + "hello") + "") + ("foo bar" + "foo bar")) + "\n\t")}
                switch (fi20 & 320) {
                case 1:
                    for fi25 := 0; fi25 < 2; fi25++ {
                        var v26 uint = v6
                        _ = fi25
                        _ = v26
                    }
                case 128:
                    v5 = v5
                }
                _ = fi20
                _ = v24
            }
        case "":
            v27 := GenS1{F0: "hello", F1: ("hello" + "")}
            _ = v27
        case "hello":
            for fi28 := 0; fi28 < 18; fi28++ {
                v2 = 9223372036854775807
                if v14 {
                    switch (fi28 & 461) {
                    case 0:
                        var v29 uint64 = 75
                        _ = v29
                    }
                    if ((18446744073709551615 <= 4294967295) && false) {
                        for fi30 := 0; fi30 < 5; fi30++ {
                            v14 = (886 < v1)
                            var v31 uint64 = 101
                            _ = len(v3)
                            _ = fi30
                            _ = v31
                        }
                    } else {
                        for fi32 := 0; fi32 < 6; fi32++ {
                            if v14 {
                                v3 = append(v3, ("\n\t" + "a"))
                                v2 += 1
                            } else {
                                if (3.14159 >= 0.1) {
                                    v3 = append(v3, (("a" + "\n\t") + ""))
                                }
                                if v14 {
                                    fmt.Println(fi28, v8)
                                }
                                v8 = (v8 % (((v8 ^ v8)) | 1))
                            }
                            if (v18 > (fi28 / ((v18) | 1))) {
                                if v14 {
                                    if ((fi28 / ((fi28) | 1)) < fi28) {
                                        switch "\n\t" {
                                        case "hello":
                                            v33 := GenS2{F0: 127, F1: (v9 - v7)}
                                            _ = v33
                                        case "foo bar":
                                            v18 += -1
                                        case "":
                                            if (!(-1 < -1)) {
                                                fmt.Println(v5)
                                            }
                                        default:
                                            if (true || (12 != (v9 * int16(v1)))) {
                                                if (!((true || v14) || v14)) {
                                                    var v34 int = (fi28 * 32768)
                                                    switch fi32 {
                                                    case 722:
                                                        v1 = (v1 * v1)
                                                    case 32767:
                                                        if (v14 && (v8 != 255)) {
                                                            switch -9223372036854775808 {
                                                            case 437:
                                                                if v14 {
                                                                    if (v14 && ((!true) && ((((("a" + "\n\t") + "a") + ("foo bar" + "\n\t")) + ("\n\t" + "")) < ((("" + ("foo bar" + "")) + (("" + "\n\t") + ("hello" + ""))) + "")))) {
                                                                        v35 := GenS3{F0: "hello"}
                                                                        _ = v35
                                                                    }
                                                                    v36 := GenS1{F0: "", F1: "\n\t"}
                                                                    v37 := GenS4{F0: 899, F1: "foo bar", F2: "a"}
                                                                    switch -345 {
                                                                    case 211:
                                                                        fmt.Println(v5, v18, v18)
                                                                    case 2147483647:
                                                                        var v38 int16 = (v9 ^ v4)
                                                                        _ = v38
                                                                    default:
                                                                        v39 := GenS4{F0: uint64(v1), F1: ("hello" + "\n\t"), F2: "hello"}
                                                                        _ = v39
                                                                    }
                                                                    _ = v36
                                                                    _ = v37
                                                                } else {
                                                                    v40 := GenS2{F0: 0, F1: v9}
                                                                    _ = v40
                                                                }
                                                                v41 := GenS5{F0: "foo bar", F1: 512}
                                                                switch "a" {
                                                                case "\n\t":
                                                                    var v42 int16 = v4
                                                                    _ = v42
                                                                case "foo bar":
                                                                    if v14 {
                                                                        v43 := GenS6{F0: uint32(v1), F1: 100}
                                                                        _ = v43
                                                                    } else {
                                                                        fmt.Println(v8, v18, v14)
                                                                    }
                                                                case "hello":
                                                                    var v44 string = ""
                                                                    _ = v44
                                                                default:
                                                                    if len(v3) > 0 {
                                                                        v3[0] = "foo bar"
                                                                    }
                                                                    fmt.Println(v9, v1)
                                                                    var v45 bool = v14
                                                                    _ = v45
                                                                }
                                                                v46 := GenS6{F0: 1, F1: -80}
                                                                var v47 int8 = 100
                                                                _ = v41
                                                                _ = v46
                                                                _ = v47
                                                            case 555:
                                                                if (!(2147483647 > 1)) {
                                                                    if (32767 > uint16(v8)) {
                                                                        if (((371 == 0) && (v14 && (9223372036854775807 == (v2 * v2)))) || v14) {
                                                                            if v14 {
                                                                                var v48 uint = v6
                                                                                v49 := GenS7{F0: v2, F1: 42}
                                                                                _ = v48
                                                                                _ = v49
                                                                            } else {
                                                                                if true {
                                                                                    v3 = append(v3, ("a" + "a"))
                                                                                    v14 = v14
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = len(v3)
                                                                        fmt.Println(fi28, v6)
                                                                    }
                                                                    v50 := GenS1{F0: "\n\t", F1: ("foo bar" + ("\n\t" + ("" + "a")))}
                                                                    _ = v50
                                                                } else {
                                                                    var v51 bool = v14
                                                                    _ = v51
                                                                }
                                                            case 2147483647:
                                                                if (-2.25 == (((3.14159 + 0.0) - ((0.1 / -2.25) * (float64(v5) * -2.25))) / (-2.25 - 0.1))) {
                                                                    if (!(-1 > -1)) {
                                                                        var v52 int8 = 96
                                                                        v8 = v8
                                                                        _ = v52
                                                                    } else {
                                                                        switch (fi28 & v34) {
                                                                        case 202:
                                                                            var v53 int = (fi28 / ((-24) | 1))
                                                                            if true {
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = ("hello" + "a")
                                                                                }
                                                                                if (-96 != -1) {
                                                                                    v5 -= v5
                                                                                    if false {
                                                                                        switch 128 {
                                                                                        case 32768:
                                                                                            switch (fi32 >> ((459) & 63)) {
                                                                                            case 893:
                                                                                                v54 := GenS8{F0: v6}
                                                                                                _ = v54
                                                                                            }
                                                                                        }
                                                                                        var v55 uint32 = uint32(v6)
                                                                                        v56 := GenS9{F0: v1, F1: v1, F2: uint(v8)}
                                                                                        _ = v55
                                                                                        _ = v56
                                                                                    } else {
                                                                                        switch (fi28 >> ((18446744073709551615) & 63)) {
                                                                                        case 140:
                                                                                            var v57 string = ("\n\t" + "")
                                                                                            _ = v57
                                                                                        case 9223372036854775807:
                                                                                            v53 = (fi32 / (((fi32 - (fi28 * fi32))) | 1))
                                                                                        case 2147483647:
                                                                                            var v58 bool = ((!(v14 && v14)) || v14)
                                                                                            _ = v58
                                                                                        default:
                                                                                            var v59 uint16 = (v1 % ((v1) | 1))
                                                                                            _ = v59
                                                                                        }
                                                                                    }
                                                                                    switch fi32 {
                                                                                    case 909:
                                                                                        if ((1.5 / 1.5) < 0.0) {
                                                                                            switch ("hello" + "a") {
                                                                                            case "hello":
                                                                                                fmt.Println(v1, v5)
                                                                                            default:
                                                                                                if true {
                                                                                                    var v60 int = 310
                                                                                                    switch (v34 % ((712) | 1)) {
                                                                                                    case 1:
                                                                                                        v61 := GenS6{F0: 805, F1: -1}
                                                                                                        _ = v61
                                                                                                    case -9223372036854775808:
                                                                                                        fmt.Println(v18, v14, v53)
                                                                                                    case 9223372036854775807:
                                                                                                        if len(v3) > 0 {
                                                                                                            v3[0] = "foo bar"
                                                                                                        }
                                                                                                    default:
                                                                                                        switch (^(v34)) {
                                                                                                        case -1:
                                                                                                            v3 = append(v3, ("foo bar" + ("" + ("foo bar" + ""))))
                                                                                                        case -288:
                                                                                                            v2 = -1
                                                                                                        default:
                                                                                                            var v62 int32 = -659
                                                                                                            _ = v62
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v60
                                                                                                } else {
                                                                                                    v63 := GenS9{F0: 255, F1: v1, F2: v6}
                                                                                                    v7 = (v4 + int16(v1))
                                                                                                    fmt.Println(v4, v18, v8)
                                                                                                    var v64 uint = v6
                                                                                                    if v14 {
                                                                                                        switch 32767 {
                                                                                                        case -9223372036854775808:
                                                                                                            v18 = (-(v34))
                                                                                                            var v65 uint64 = 4294967295
                                                                                                            if (v14 || false) {
                                                                                                                if len(v3) >= 2 {
                                                                                                                    v3 = v3[1:len(v3)]
                                                                                                                }
                                                                                                            }
                                                                                                            var v66 string = "hello"
                                                                                                            v67 := GenS10{F0: 42, F1: v7}
                                                                                                            _ = v65
                                                                                                            _ = v66
                                                                                                            _ = v67
                                                                                                        }
                                                                                                        v34 += v18
                                                                                                    }
                                                                                                    _ = v63
                                                                                                    _ = v64
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v2, v4)
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v68 int16 = (^(v7))
                                                                                _ = v68
                                                                            }
                                                                            _ = v53
                                                                        default:
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    v2 += (-(v2))
                                                                                case "hello":
                                                                                    var v69 uint8 = (v8 & 0)
                                                                                    _ = v69
                                                                                default:
                                                                                    var v70 uint64 = 266
                                                                                    _ = v70
                                                                                }
                                                                            case "\n\t":
                                                                                if ((v6 * (v6 >> ((2) & 63))) <= (v6 % (((^(v6))) | 1))) {
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = "hello"
                                                                                    }
                                                                                } else {
                                                                                    var v71 uint8 = 128
                                                                                    v14 = (v7 < v7)
                                                                                    var v72 bool = false
                                                                                    _ = v71
                                                                                    _ = v72
                                                                                }
                                                                            }
                                                                        }
                                                                        switch fi32 {
                                                                        case 4294967295:
                                                                            if v14 {
                                                                                fmt.Println(fi28)
                                                                            } else {
                                                                                v73 := GenS2{F0: -1, F1: v4}
                                                                                var v74 int32 = -2147483648
                                                                                if (v5 == v5) {
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        var v75 uint32 = uint32(v1)
                                                                                        _ = v75
                                                                                    case "foo bar":
                                                                                        if v14 {
                                                                                            if v14 {
                                                                                                var v76 uint16 = 42
                                                                                                if ((v74 - v74) != (v74 >> ((v6) & 63))) {
                                                                                                    if (!false) {
                                                                                                        v77 := GenS4{F0: 4294967295, F1: "a", F2: ""}
                                                                                                        switch ("a" + ((("\n\t" + "a") + "") + "")) {
                                                                                                        case "hello":
                                                                                                            if v14 {
                                                                                                                v7 = -1
                                                                                                                if true {
                                                                                                                    fmt.Println(v76)
                                                                                                                    fmt.Println(fi32, v8)
                                                                                                                    v78 := GenS11{F0: v2, F1: v74}
                                                                                                                    if ((v1 ^ 120) > v76) {
                                                                                                                        _ = len(v3)
                                                                                                                        if (42 >= 127) {
                                                                                                                            var v79 uint = 661
                                                                                                                            fmt.Println(v7, v1)
                                                                                                                            _ = v79
                                                                                                                        }
                                                                                                                        v6 = (^(v6))
                                                                                                                    }
                                                                                                                    _ = v78
                                                                                                                } else {
                                                                                                                    fmt.Println(fi28)
                                                                                                                    if v14 {
                                                                                                                        v5 += float32(3.14159)
                                                                                                                    }
                                                                                                                    v80 := GenS12{F0: (fi32 ^ (v18 ^ -1)), F1: v2}
                                                                                                                    _ = v80
                                                                                                                }
                                                                                                                if len(v3) >= 2 {
                                                                                                                    v3 = v3[1:len(v3)]
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        var v81 float64 = 1.5
                                                                                                        _ = v77
                                                                                                        _ = v81
                                                                                                    } else {
                                                                                                        v82 := GenS13{F0: v74, F1: v6}
                                                                                                        v6 -= (v6 | 255)
                                                                                                        v1 = v76
                                                                                                        v5 -= v5
                                                                                                        _ = v82
                                                                                                    }
                                                                                                    if len(v3) >= 2 {
                                                                                                        v3 = v3[1:len(v3)]
                                                                                                    }
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = "a"
                                                                                                    }
                                                                                                } else {
                                                                                                    var v83 uint16 = v76
                                                                                                    _ = v83
                                                                                                }
                                                                                                v84 := GenS14{F0: v14, F1: "hello"}
                                                                                                _ = v76
                                                                                                _ = v84
                                                                                            }
                                                                                            v85 := GenS15{F0: v14}
                                                                                            _ = v85
                                                                                        }
                                                                                        if ((0.1 == float64(0.0)) || (!v14)) {
                                                                                            fmt.Println(fi32, fi28, fi28)
                                                                                        }
                                                                                        if true {
                                                                                            if (fi32 != (fi32 / (((v18 >> ((v6) & 63))) | 1))) {
                                                                                                v86 := GenS16{F0: ("" + "\n\t")}
                                                                                                var v87 int64 = (v2 << ((v6) & 63))
                                                                                                v88 := GenS1{F0: "\n\t", F1: ("\n\t" + "foo bar")}
                                                                                                _ = v86
                                                                                                _ = v87
                                                                                                _ = v88
                                                                                            } else {
                                                                                                v89 := GenS4{F0: 638, F1: "foo bar", F2: "hello"}
                                                                                                fmt.Println(v74)
                                                                                                var v90 int16 = 1
                                                                                                _ = v89
                                                                                                _ = v90
                                                                                            }
                                                                                            var v91 int64 = -128
                                                                                            _ = v91
                                                                                        } else {
                                                                                            if v14 {
                                                                                                _ = len(v3)
                                                                                                if true {
                                                                                                    fmt.Println(v2)
                                                                                                    var v92 uint8 = v8
                                                                                                    _ = v92
                                                                                                }
                                                                                            } else {
                                                                                                var v93 uint8 = (v8 >> ((18446744073709551615) & 63))
                                                                                                _ = v93
                                                                                            }
                                                                                            switch (v34 ^ 127) {
                                                                                            case 772:
                                                                                                if v14 {
                                                                                                    if (2 <= v8) {
                                                                                                        fmt.Println(fi28)
                                                                                                    } else {
                                                                                                        if (true && v14) {
                                                                                                            v94 := GenS13{F0: v74, F1: v6}
                                                                                                            _ = v94
                                                                                                        }
                                                                                                    }
                                                                                                    v9 -= -696
                                                                                                    if false {
                                                                                                        fmt.Println(fi28, v6, v18)
                                                                                                        v95 := GenS6{F0: 2147483647, F1: 0}
                                                                                                        _ = v95
                                                                                                    } else {
                                                                                                        switch v18 {
                                                                                                        case 0:
                                                                                                            var v96 string = "hello"
                                                                                                            _ = v96
                                                                                                        case 2147483647:
                                                                                                            switch "a" {
                                                                                                            case "":
                                                                                                                v14 = v14
                                                                                                                v74 += v74
                                                                                                            case "\n\t":
                                                                                                                v97 := GenS10{F0: 85, F1: (-(v7))}
                                                                                                                _ = v97
                                                                                                            }
                                                                                                            var v98 int64 = (v2 | -293)
                                                                                                            v3 = append(v3, "foo bar")
                                                                                                            var v99 uint16 = uint16(v8)
                                                                                                            _ = v98
                                                                                                            _ = v99
                                                                                                        case 32767:
                                                                                                            var v100 string = "a"
                                                                                                            _ = v100
                                                                                                        default:
                                                                                                            v5 = v5
                                                                                                        }
                                                                                                        v18 = (fi28 & v34)
                                                                                                        v5 -= v5
                                                                                                    }
                                                                                                } else {
                                                                                                    var v101 float64 = 0.0
                                                                                                    fmt.Println(v8)
                                                                                                    var v102 int32 = 0
                                                                                                    _ = v101
                                                                                                    _ = v102
                                                                                                }
                                                                                            case -1:
                                                                                                v7 = 1
                                                                                            case -9223372036854775808:
                                                                                                v6 -= (v6 + (v6 + v6))
                                                                                            }
                                                                                            var v103 uint = v6
                                                                                            _ = v103
                                                                                        }
                                                                                    case "\n\t":
                                                                                        fmt.Println(v74, v2, v9)
                                                                                    }
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = "\n\t"
                                                                                    }
                                                                                    switch "hello" {
                                                                                    case "\n\t":
                                                                                        v5 += v5
                                                                                    default:
                                                                                        v1 += (^(v1))
                                                                                    }
                                                                                } else {
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    var v104 uint64 = 18446744073709551615
                                                                                    var v105 string = "a"
                                                                                    _ = v104
                                                                                    _ = v105
                                                                                }
                                                                                _ = v73
                                                                                _ = v74
                                                                            }
                                                                        }
                                                                        var v106 int32 = 42
                                                                        _ = v106
                                                                    }
                                                                    var v107 int8 = 127
                                                                    _ = v107
                                                                } else {
                                                                    v1 += (v1 * v1)
                                                                }
                                                            default:
                                                                fmt.Println(v4)
                                                            }
                                                            var v108 int16 = (v7 >> ((v6) & 63))
                                                            var v109 uint = v6
                                                            _ = v108
                                                            _ = v109
                                                        } else {
                                                            var v110 uint16 = (^(v1))
                                                            var v111 uint64 = 118
                                                            _ = v110
                                                            _ = v111
                                                        }
                                                    case 128:
                                                        v7 = int16(v2)
                                                    default:
                                                        v6 -= v6
                                                    }
                                                    _ = v34
                                                }
                                                v1 -= 680
                                            }
                                        }
                                        if (!false) {
                                            _ = len(v3)
                                            if false {
                                                v3 = append(v3, ("" + "\n\t"))
                                                if (4294967295 <= 628) {
                                                    v112 := GenS1{F0: "", F1: "\n\t"}
                                                    v113 := GenS17{F0: 32767, F1: (v2 | v2), F2: 42}
                                                    if v14 {
                                                        fmt.Println(v6)
                                                        var v114 string = "hello"
                                                        v6 += v6
                                                        v18 -= (v18 % ((-9223372036854775808) | 1))
                                                        var v115 bool = ((-(v2)) <= (v2 ^ v2))
                                                        v2 = (v2 & v2)
                                                        _ = v114
                                                        _ = v115
                                                    } else {
                                                        v116 := GenS14{F0: (!(!v14)), F1: "hello"}
                                                        _ = v116
                                                    }
                                                    v3 = append(v3, ("hello" + ("hello" + (((("a" + "foo bar") + "\n\t") + "\n\t") + (("foo bar" + "foo bar") + "foo bar")))))
                                                    _ = v112
                                                    _ = v113
                                                }
                                                _ = len(v3)
                                            }
                                        } else {
                                            v117 := GenS9{F0: v1, F1: v1, F2: (v6 & v6)}
                                            var v118 float32 = (0.1 / (v5 - v5))
                                            _ = v117
                                            _ = v118
                                        }
                                        switch 128 {
                                        case 152:
                                            v18 = fi28
                                            if v14 {
                                                var v119 uint32 = uint32(v6)
                                                if false {
                                                    if len(v3) > 0 {
                                                        v3[0] = ("" + "foo bar")
                                                    }
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = "a"
                                                    }
                                                    if v14 {
                                                        if (v5 > 3.14159) {
                                                            v120 := GenS18{F0: 2147483647, F1: v6, F2: v8}
                                                            var v121 int = 42
                                                            fmt.Println(v5, v121, v14)
                                                            if true {
                                                                var v122 float64 = (1e10 / 1.0)
                                                                fmt.Println(v8, v6)
                                                                _ = v122
                                                            } else {
                                                                if (3.14159 != v5) {
                                                                    v4 += (v4 - (v4 * (v7 + v7)))
                                                                    var v123 int64 = v2
                                                                    _ = v123
                                                                } else {
                                                                    v124 := GenS3{F0: "foo bar"}
                                                                    _ = v124
                                                                }
                                                                var v125 int16 = v9
                                                                v2 = -1
                                                                _ = v125
                                                            }
                                                            _ = v120
                                                            _ = v121
                                                        }
                                                    } else {
                                                        if false {
                                                            if (uint64(v8) < 693) {
                                                                v4 += v7
                                                                v8 -= v8
                                                                if true {
                                                                    if v14 {
                                                                        if true {
                                                                            v126 := GenS1{F0: ("" + ((("foo bar" + "a") + "hello") + "a")), F1: "foo bar"}
                                                                            _ = v126
                                                                        }
                                                                    } else {
                                                                        switch ("hello" + "") {
                                                                        case "hello":
                                                                            fmt.Println(v6)
                                                                            var v127 string = ""
                                                                            if len(v3) > 0 {
                                                                                v3[0] = v127
                                                                            }
                                                                            _ = v127
                                                                        default:
                                                                            v128 := GenS19{F0: v2}
                                                                            _ = v128
                                                                        }
                                                                        v14 = (((fi32 + (fi32 & 0)) > -861) || (0 != 4294967295))
                                                                        var v129 float64 = float64(float32(-2.25))
                                                                        _ = v129
                                                                    }
                                                                    var v130 int8 = 100
                                                                    if (v14 || v14) {
                                                                        v4 = int16(v6)
                                                                        var v131 int = (-(fi28))
                                                                        if false {
                                                                            fmt.Println(v8, v14)
                                                                        } else {
                                                                            fmt.Println(v18, v9)
                                                                            var v132 int16 = (v7 | v7)
                                                                            var v133 int32 = 137
                                                                            _ = v132
                                                                            _ = v133
                                                                        }
                                                                        fmt.Println(v1, v2)
                                                                        var v134 bool = (v14 || ((^(v8)) == 1))
                                                                        var v135 string = ((("a" + "a") + ("\n\t" + "foo bar")) + "")
                                                                        v8 = (v8 ^ 110)
                                                                        _ = v131
                                                                        _ = v134
                                                                        _ = v135
                                                                    }
                                                                    _ = v130
                                                                } else {
                                                                    fmt.Println(v5)
                                                                    var v136 bool = v14
                                                                    _ = v136
                                                                }
                                                            }
                                                            if (-63 != (fi32 ^ (v18 * (fi28 >> ((v6) & 63))))) {
                                                                var v137 float64 = (1.5 + -2.25)
                                                                switch (fi32 | 32768) {
                                                                case 1:
                                                                    fmt.Println(v14, v8)
                                                                }
                                                                _ = v137
                                                            }
                                                            v119 = v119
                                                            if (v14 && true) {
                                                                if (((!true) && (!true)) && (-128 == -1)) {
                                                                    v138 := GenS10{F0: -119, F1: v7}
                                                                    if true {
                                                                        v139 := GenS13{F0: 496, F1: v6}
                                                                        _ = v139
                                                                    }
                                                                    _ = v138
                                                                }
                                                                v14 = (!false)
                                                                if len(v3) > 0 {
                                                                    v3[0] = ("a" + "foo bar")
                                                                }
                                                            } else {
                                                                v140 := GenS20{F0: 208, F1: 970, F2: 2147483647}
                                                                var v141 float32 = 1.5
                                                                _ = v140
                                                                _ = v141
                                                            }
                                                            v14 = (!(v6 == v6))
                                                        }
                                                        v9 -= v9
                                                        fmt.Println(v6, v18)
                                                    }
                                                    v142 := GenS18{F0: v119, F1: (v6 + (v6 % ((4294967295) | 1))), F2: v8}
                                                    _ = v142
                                                }
                                                _ = v119
                                            }
                                        }
                                        _ = len(v3)
                                        fmt.Println(v7, v9, v7)
                                        v2 += -535
                                        if ((v6 < (v6 ^ (v6 >> (((v6 + uint(v1))) & 63)))) || (uint64(v1) <= 868)) {
                                            v143 := GenS21{F0: (^(v9)), F1: 65535, F2: (v4 | (v4 - v4))}
                                            var v144 uint16 = (^(v1))
                                            fmt.Println(fi28, fi32, fi28)
                                            _ = v143
                                            _ = v144
                                        } else {
                                            switch ("\n\t" + ("foo bar" + ("foo bar" + ""))) {
                                            case "foo bar":
                                                var v145 int32 = -1
                                                _ = v145
                                            default:
                                                var v146 float32 = v5
                                                _ = v146
                                            }
                                        }
                                        v147 := GenS22{F0: 64}
                                        _ = v147
                                    }
                                    v148 := GenS2{F0: 42, F1: (-(v9))}
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                    switch fi28 {
                                    case -9223372036854775808:
                                        v149 := GenS9{F0: v1, F1: (v1 / (((v1 >> ((v6) & 63))) | 1)), F2: uint(v8)}
                                        _ = v149
                                    case -851:
                                        var v150 float64 = (0.0 - (0.0 / ((((1e10 * ((3.14159 + float64(v5)) - 1e10)) / -2.25) / 0.1) - (0.0 + (((float64(3.14159) / 3.14159) / 1.5) / ((3.14159 - (1.5 / 1.0)) * -2.25))))))
                                        _ = v150
                                    case -1:
                                        if true {
                                            v8 += v8
                                            v18 += -1
                                            var v151 string = "\n\t"
                                            var v152 int16 = v4
                                            if v14 {
                                                v3 = append(v3, ((v151 + "hello") + "foo bar"))
                                            }
                                            switch (fi32 / ((103) | 1)) {
                                            case 42:
                                                var v153 float64 = float64(v5)
                                                if len(v3) > 0 {
                                                    v3[0] = "a"
                                                }
                                                _ = v153
                                            case -581:
                                                v154 := GenS23{F0: 42, F1: v8}
                                                _ = v154
                                            case 9223372036854775807:
                                                switch (("foo bar" + "hello") + "\n\t") {
                                                case "a":
                                                    v155 := GenS14{F0: (!(v152 != (v7 * 0))), F1: v151}
                                                    _ = v155
                                                case "hello":
                                                    if true {
                                                        switch (fi28 | -2147483648) {
                                                        case 993:
                                                            v1 -= (v1 - 0)
                                                        case 9223372036854775807:
                                                            switch -509 {
                                                            case 32768:
                                                                var v156 float32 = 0.0
                                                                _ = v156
                                                            }
                                                        case -9223372036854775808:
                                                            v157 := GenS24{F0: (v152 / (((-(v9))) | 1)), F1: 743, F2: 42}
                                                            _ = v157
                                                        }
                                                    } else {
                                                        if false {
                                                            v6 -= v6
                                                            _ = len(v3)
                                                            if len(v3) > 0 {
                                                                v3[0] = "hello"
                                                            }
                                                        }
                                                        v4 -= (v9 & v4)
                                                        var v158 string = (("a" + v151) + ("a" + ((v151 + ("" + (v151 + v151))) + v151)))
                                                        v152 = v4
                                                        _ = v158
                                                    }
                                                case "foo bar":
                                                    if ("hello" >= v151) {
                                                        var v159 int8 = 127
                                                        _ = v159
                                                    } else {
                                                        switch 677 {
                                                        case 32767:
                                                            var v160 int8 = 100
                                                            _ = v160
                                                        case -122:
                                                            v8 = (^(v8))
                                                        }
                                                    }
                                                    if v14 {
                                                        if false {
                                                            if v14 {
                                                                var v161 uint = uint(v1)
                                                                var v162 int16 = 897
                                                                _ = v161
                                                                _ = v162
                                                            } else {
                                                                fmt.Println(v2, v151, v14)
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            }
                                                            v1 = (v1 << ((v6) & 63))
                                                            if (!v14) {
                                                                var v163 uint32 = 961
                                                                _ = v163
                                                            } else {
                                                                if len(v3) > 0 {
                                                                    v3[0] = ((v151 + (v151 + "a")) + ("foo bar" + v151))
                                                                }
                                                                v8 -= v8
                                                                var v164 string = ("" + (((((((v151 + ("hello" + "a")) + (v151 + (v151 + v151))) + v151) + "foo bar") + v151) + (v151 + v151)) + v151))
                                                                _ = v164
                                                            }
                                                        } else {
                                                            fmt.Println(v14)
                                                            var v165 float32 = v5
                                                            v3 = append(v3, v151)
                                                            _ = v165
                                                        }
                                                        var v166 bool = v14
                                                        _ = v166
                                                    } else {
                                                        v1 = 255
                                                        _ = len(v3)
                                                        if true {
                                                            if ((v1 >> (((v6 >> ((v6) & 63))) & 63)) < v1) {
                                                                var v167 string = ""
                                                                _ = v167
                                                            } else {
                                                                if v14 {
                                                                    switch fi28 {
                                                                    case -371:
                                                                        if false {
                                                                            var v168 bool = ((fi32 | (^(v18))) == fi28)
                                                                            if (!true) {
                                                                                if ((fi32 == 2147483647) && ((v5 > float32(0.0)) || ((!(fi32 <= (v18 & (^(v18))))) || (((float32(-2.25) - v5) * -2.25) != v5)))) {
                                                                                    var v169 uint64 = 2
                                                                                    _ = v169
                                                                                }
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = v151
                                                                                }
                                                                                switch v18 {
                                                                                case 2147483647:
                                                                                    v170 := GenS18{F0: 536, F1: v6, F2: v8}
                                                                                    _ = v170
                                                                                case 9223372036854775807:
                                                                                    var v171 int = int(v6)
                                                                                    _ = v171
                                                                                case -832:
                                                                                    fmt.Println(v152, v7)
                                                                                    v172 := GenS3{F0: v151}
                                                                                    v18 = (^(fi28))
                                                                                    _ = v172
                                                                                }
                                                                                var v173 int = 538
                                                                                fmt.Println(v173)
                                                                                _ = v173
                                                                            }
                                                                            v14 = (v151 <= "hello")
                                                                            _ = v168
                                                                        }
                                                                    default:
                                                                        var v174 string = v151
                                                                        _ = v174
                                                                    }
                                                                    v3 = append(v3, (v151 + (((v151 + (v151 + (("foo bar" + v151) + (v151 + v151)))) + v151) + v151)))
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                                v175 := GenS24{F0: (v152 & (v9 % (((v7 | v152)) | 1))), F1: -1, F2: 0}
                                                                _ = v175
                                                            }
                                                            fmt.Println(v14)
                                                            fmt.Println(v1, fi32)
                                                        }
                                                    }
                                                    v3 = append(v3, "hello")
                                                }
                                            default:
                                                if v14 {
                                                    var v176 float64 = 0.0
                                                    _ = v176
                                                }
                                            }
                                            v5 -= 1.5
                                            fmt.Println(v1, v151, v4)
                                            var v177 uint16 = v1
                                            _ = v151
                                            _ = v152
                                            _ = v177
                                        } else {
                                            if true {
                                                var v178 string = "hello"
                                                if len(v3) > 0 {
                                                    v3[0] = v178
                                                }
                                                v4 = -1
                                                _ = v178
                                            }
                                        }
                                    }
                                    _ = v148
                                } else {
                                    var v179 uint32 = 42
                                    _ = v179
                                }
                            } else {
                                fmt.Println(v5)
                            }
                            _ = fi32
                        }
                        if (18446744073709551615 == 433) {
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                            v180 := GenS12{F0: fi28, F1: -318}
                            if (int8(v6) < 104) {
                                var v181 uint16 = v1
                                var v182 float32 = float32(float64(v5))
                                _ = len(v3)
                                _ = v181
                                _ = v182
                            }
                            _ = v180
                        }
                        var v183 uint = v6
                        _ = v183
                    }
                    var v184 float32 = (v5 - (0.0 + 3.14159))
                    _ = v184
                } else {
                    var v185 int = (^(v18))
                    v18 = v18
                    var v186 bool = (-72 >= -46)
                    _ = v185
                    _ = v186
                }
                _ = fi28
            }
        default:
            v187 := GenS17{F0: uint16(v6), F1: -9223372036854775808, F2: 862}
            _ = v187
        }
        _ = v18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
