package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 bool; F1 int8; F2 uint }
type GenS3 struct { F0 string; F1 int16; F2 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 uint64; F1 uint; F2 int }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16; F1 string; F2 uint16 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint; F1 int64; F2 int32 }
type GenS14 struct { F0 int8; F1 bool; F2 int64 }

func main() {
    var v1 float32 = ((0.1 * float32(-2.25)) * 0.1)
    var v2 float64 = 1.5
    v3 := []string{"a", ("hello" + "\n\t"), "hello"}
    var v4 bool = (-1 == -2147483648)
    var v5 uint16 = 255
    v4 = ((1.5 / (-2.25 + v1)) != (((((v1 / 1.5) / ((v1 - (-2.25 + (v1 - float32((3.14159 * v2))))) - ((v1 * v1) / v1))) + ((((((float32((((float64(v1) + v2) + (((v2 * ((v2 - 1e10) - (0.1 + 3.14159))) / float64(v1)) * ((0.1 - (float64(-2.25) + (v2 / v2))) / ((float64(float32(0.0)) / (v2 - float64(1.5))) - 1e10)))) * v2)) * float32((((v2 - v2) * v2) / v2))) - (-2.25 / v1)) + float32(v2)) - (-2.25 / v1)) + v1) / v1)) * 3.14159) - float32(0.1)))
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 int16 = 42
        _ = fi6
        _ = v7
    }
    if true {
        var v8 uint16 = v5
        _ = len(v3)
        var v9 uint16 = 1
        _ = v8
        _ = v9
    }
    if (84 >= 42) {
        if true {
            if (!v4) {
                var v10 uint = 25
                v11 := GenS1{F0: (v10 % ((v10) | 1)), F1: 0}
                _ = v10
                _ = v11
            } else {
                for fi12 := 0; fi12 < 12; fi12++ {
                    if (false && (v4 && (1 >= 1))) {
                        var v13 int64 = int64(v5)
                        v14 := GenS1{F0: 65535, F1: 2}
                        v2 = (0.1 - v2)
                        _ = v13
                        _ = v14
                    }
                    var v15 string = ((("" + "foo bar") + (("\n\t" + (("\n\t" + (("foo bar" + "a") + "foo bar")) + (("a" + ("\n\t" + ("\n\t" + "hello"))) + "foo bar"))) + (("a" + "hello") + ("foo bar" + "a")))) + "")
                    _ = len(v3)
                    _ = fi12
                    _ = v15
                }
                if false {
                    switch 2147483647 {
                    case 539:
                        fmt.Println(v2)
                    default:
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                    }
                    fmt.Println(v4, v1, v1)
                }
                switch ("" + "") {
                case "foo bar":
                    if (v2 > v2) {
                        for fi16 := 0; fi16 < 2; fi16++ {
                            v2 += float64(-2.25)
                            _ = fi16
                        }
                        v17 := GenS1{F0: 367, F1: 2}
                        var v18 int32 = -164
                        _ = v17
                        _ = v18
                    }
                case "\n\t":
                    _ = len(v3)
                case "a":
                    if v4 {
                        for fi19 := 0; fi19 < 15; fi19++ {
                            if (2 <= 255) {
                                var v20 uint32 = uint32(v5)
                                if len(v3) > 0 {
                                    v3[0] = (("a" + "\n\t") + ("" + (("foo bar" + "a") + "")))
                                }
                                v4 = (2 != 0)
                                _ = v20
                            } else {
                                var v21 uint = uint(v5)
                                var v22 string = "hello"
                                _ = v21
                                _ = v22
                            }
                            _ = fi19
                        }
                        switch "foo bar" {
                        case "\n\t":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "foo bar":
                            v23 := GenS2{F0: v4, F1: 127, F2: 252}
                            _ = v23
                        }
                        var v24 float32 = 0.1
                        _ = v24
                    }
                default:
                    switch 2147483647 {
                    case 2147483647:
                        for fi25 := 0; fi25 < 2; fi25++ {
                            var v26 uint16 = (^(v5))
                            if false {
                                var v27 int32 = -1
                                if (v4 && true) {
                                    v28 := GenS1{F0: 65535, F1: 127}
                                    v29 := GenS2{F0: (v4 && false), F1: 100, F2: 18446744073709551615}
                                    v26 -= v26
                                    _ = v28
                                    _ = v29
                                } else {
                                    v30 := GenS1{F0: 1, F1: 71}
                                    _ = len(v3)
                                    _ = v30
                                }
                                _ = v27
                            } else {
                                var v31 int32 = -1
                                for fi32 := 0; fi32 < 11; fi32++ {
                                    var v33 string = "foo bar"
                                    var v34 int = fi25
                                    _ = fi32
                                    _ = v33
                                    _ = v34
                                }
                                var v35 int = fi25
                                _ = v31
                                _ = v35
                            }
                            _ = fi25
                            _ = v26
                        }
                    default:
                        v1 = 0.1
                    }
                }
            }
            v5 -= v5
            if false {
                var v36 int16 = -604
                v5 -= v5
                var v37 int = 32767
                _ = v36
                _ = v37
            } else {
                fmt.Println(v4)
                fmt.Println(v4, v2, v2)
                if v4 {
                    v38 := GenS1{F0: 18446744073709551615, F1: 255}
                    _ = v38
                } else {
                    v4 = (1 >= -128)
                }
            }
        } else {
            fmt.Println(v4, v1)
            if v4 {
                v39 := GenS3{F0: "a", F1: -616, F2: 255}
                if (!(((!(0 <= 0)) || (!(633 == 2))) || v4)) {
                    if (354 >= 42) {
                        v40 := GenS1{F0: 922, F1: 129}
                        v3 = append(v3, "\n\t")
                        v5 = (v5 + (v5 * 416))
                        _ = v40
                    }
                }
                if false {
                    switch 2147483647 {
                    case -925:
                        var v41 bool = false
                        _ = v41
                    case -1:
                        v1 = v1
                    case 128:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    if v4 {
                        v4 = (false || (!v4))
                        for fi42 := 0; fi42 < 4; fi42++ {
                            fmt.Println(v5)
                            v5 += (v5 * (^(v5)))
                            _ = fi42
                        }
                        fmt.Println(v2, v1)
                    }
                } else {
                    fmt.Println(v2, v1)
                    if len(v3) > 0 {
                        v3[0] = ("a" + ("hello" + (("foo bar" + "\n\t") + ("foo bar" + ""))))
                    }
                    for fi43 := 0; fi43 < 19; fi43++ {
                        v44 := GenS2{F0: v4, F1: 100, F2: 40}
                        v5 = (^(v5))
                        if ((-1 != 42) || ((!(-1 > -1)) && (1 < -9223372036854775808))) {
                            var v45 uint32 = 42
                            _ = v45
                        } else {
                            if false {
                                v46 := GenS1{F0: 1, F1: 1}
                                _ = v46
                            }
                            v3 = append(v3, ((("\n\t" + (("a" + "\n\t") + "a")) + ("\n\t" + "")) + ("a" + ("a" + "\n\t"))))
                            fmt.Println(v2)
                        }
                        _ = fi43
                        _ = v44
                    }
                }
                _ = v39
            } else {
                for fi47 := 0; fi47 < 20; fi47++ {
                    for fi48 := 0; fi48 < 9; fi48++ {
                        switch (fi47 ^ (^(fi47))) {
                        case -9223372036854775808:
                            switch (^(fi48)) {
                            case 127:
                                v3 = append(v3, ("hello" + ("" + "hello")))
                                fmt.Println(v4)
                            case 128:
                                var v49 int8 = -1
                                _ = v49
                            default:
                                v5 = (^(v5))
                            }
                            var v50 uint32 = 1
                            if (v1 != 1e10) {
                                v51 := GenS3{F0: (("a" + ("" + ("" + ""))) + "foo bar"), F1: 625, F2: 65535}
                                _ = v51
                            }
                            _ = v50
                        case 2147483647:
                            v3 = append(v3, ("\n\t" + "a"))
                        }
                        fmt.Println(v2, v1, fi47)
                        _ = fi48
                    }
                    v1 = ((0.0 * float32(v2)) * 0.1)
                    _ = fi47
                }
                fmt.Println(v5)
                v52 := GenS3{F0: ("hello" + "foo bar"), F1: -380, F2: (v5 + v5)}
                _ = v52
            }
            v3 = append(v3, "a")
        }
        v1 -= 0.1
        for fi53 := 0; fi53 < 14; fi53++ {
            for fi54 := 0; fi54 < 20; fi54++ {
                if (432 >= 636) {
                    if ((v5 == v5) || ((float64((v1 + v1)) / v2) == 1.5)) {
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        v4 = v4
                        if false {
                            var v55 bool = v4
                            if (float64(v1) == v2) {
                                if (!(true || true)) {
                                    if len(v3) > 0 {
                                        v3[0] = "\n\t"
                                    }
                                } else {
                                    v3 = append(v3, "hello")
                                }
                                var v56 uint16 = (v5 * v5)
                                var v57 string = ""
                                _ = v56
                                _ = v57
                            } else {
                                v1 = v1
                            }
                            _ = v55
                        } else {
                            var v58 uint64 = 223
                            switch fi53 {
                            case 252:
                                var v59 uint16 = (v5 + (v5 % ((v5) | 1)))
                                _ = v59
                            default:
                                fmt.Println(v58, v1)
                            }
                            _ = v58
                        }
                    }
                    var v60 uint16 = v5
                    var v61 uint64 = 4294967295
                    _ = v60
                    _ = v61
                } else {
                    v2 = v2
                    v1 = (((v1 + 0.0) - v1) / (v1 / v1))
                }
                switch fi53 {
                case 42:
                    v2 = v2
                default:
                    switch "foo bar" {
                    case "a":
                        var v62 uint8 = 0
                        _ = v62
                    case "hello":
                        if ((3.14159 == v1) && (!v4)) {
                            v3 = append(v3, (("foo bar" + "") + "hello"))
                            if (false || true) {
                                if (("a" + "a") == "foo bar") {
                                    v5 = v5
                                }
                            } else {
                                v5 += v5
                                var v63 bool = true
                                var v64 int8 = -128
                                _ = v63
                                _ = v64
                            }
                            v3 = append(v3, "")
                            var v65 int64 = -503
                            v1 = v1
                            if ((((("a" + "foo bar") + ("a" + (("\n\t" + "hello") + ("foo bar" + (("hello" + ((("" + ("" + ("foo bar" + "hello"))) + "") + "hello")) + ("" + "")))))) + "foo bar") < "foo bar") && (25 == 42)) {
                                if (385 <= 407) {
                                    if ((v65 | v65) != v65) {
                                        fmt.Println(v4, v1)
                                        var v66 string = ((("" + "") + "hello") + ("foo bar" + ""))
                                        _ = v66
                                    }
                                    v3 = append(v3, "")
                                    if (v4 && v4) {
                                        var v67 int64 = (-(v65))
                                        _ = v67
                                    }
                                }
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = "hello"
                                }
                                if v4 {
                                    v1 = v1
                                    v68 := GenS4{F0: ((329 >= 0) && (true || v4))}
                                    _ = len(v3)
                                    _ = v68
                                } else {
                                    v4 = v4
                                    var v69 uint32 = 2147483648
                                    if v4 {
                                        switch fi54 {
                                        case 9223372036854775807:
                                            v65 = -9223372036854775808
                                        default:
                                            fmt.Println(v69, fi54)
                                        }
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(v65, v69, v4)
                                        default:
                                            var v70 int = (-(fi53))
                                            _ = v70
                                        }
                                        v71 := GenS5{F0: fi53}
                                        _ = v71
                                    } else {
                                        switch -455 {
                                        case -80:
                                            fmt.Println(fi54)
                                        default:
                                            fmt.Println(v1, v4, v69)
                                            switch "foo bar" {
                                            case "":
                                                switch ((("hello" + "\n\t") + ("\n\t" + "hello")) + "hello") {
                                                case "":
                                                    _ = len(v3)
                                                case "a":
                                                    fmt.Println(v69, v1)
                                                }
                                            default:
                                                if (0.1 == ((1.5 / v1) - 0.0)) {
                                                    if (int(v69) == (^(fi53))) {
                                                        var v72 bool = v4
                                                        v73 := GenS3{F0: ("a" + "\n\t"), F1: 1000, F2: v5}
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                } else {
                                                    var v74 int64 = v65
                                                    if v4 {
                                                        var v75 uint8 = 130
                                                        _ = v75
                                                    } else {
                                                        var v76 uint8 = 127
                                                        var v77 bool = (!((v4 && v4) || true))
                                                        if (!(2 >= 0)) {
                                                            v3 = append(v3, "\n\t")
                                                            v78 := GenS4{F0: true}
                                                            if v4 {
                                                                var v79 float32 = v1
                                                                _ = v79
                                                            }
                                                            _ = v78
                                                        } else {
                                                            if false {
                                                                switch 128 {
                                                                case 1:
                                                                    switch ("foo bar" + (("\n\t" + "\n\t") + (("" + ("\n\t" + ("hello" + "\n\t"))) + "hello"))) {
                                                                    case "foo bar":
                                                                        switch int(v69) {
                                                                        case -9223372036854775808:
                                                                            v80 := GenS3{F0: ("foo bar" + ((((((("foo bar" + "hello") + "") + ("a" + "")) + "foo bar") + "") + "a") + ("hello" + "\n\t"))), F1: 32767, F2: (v5 ^ 42)}
                                                                            _ = v80
                                                                        case 32767:
                                                                            if len(v3) > 0 {
                                                                                v3[0] = ("\n\t" + "a")
                                                                            }
                                                                        case -845:
                                                                            var v81 uint32 = (v69 % ((967) | 1))
                                                                            if false {
                                                                                fmt.Println(v5, v74, v76)
                                                                                var v82 float32 = (3.14159 - v1)
                                                                                fmt.Println(v74)
                                                                                _ = v82
                                                                            }
                                                                            if len(v3) > 0 {
                                                                                v3[0] = "hello"
                                                                            }
                                                                            if (v4 && (!false)) {
                                                                                switch ("foo bar" + "") {
                                                                                case "":
                                                                                    var v83 string = "foo bar"
                                                                                    _ = v83
                                                                                case "\n\t":
                                                                                    if true {
                                                                                        v84 := GenS6{F0: 42, F1: 1}
                                                                                        if (v2 >= float64((v1 - v1))) {
                                                                                            v2 = (v2 - v2)
                                                                                            v1 += (-2.25 * 1e10)
                                                                                        } else {
                                                                                            var v85 int32 = -71
                                                                                            var v86 int64 = v65
                                                                                            _ = v85
                                                                                            _ = v86
                                                                                        }
                                                                                        _ = v84
                                                                                    } else {
                                                                                        v76 -= 128
                                                                                        var v87 bool = v4
                                                                                        var v88 int32 = 0
                                                                                        v89 := GenS7{F0: 736, F1: 18446744073709551615, F2: fi53}
                                                                                        if v77 {
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                        } else {
                                                                                            if (float32(v2) >= v1) {
                                                                                                v90 := GenS5{F0: 128}
                                                                                                switch fi53 {
                                                                                                case 4294967295:
                                                                                                    var v91 bool = true
                                                                                                    _ = v91
                                                                                                case 127:
                                                                                                    v2 -= 1.5
                                                                                                    v92 := GenS1{F0: 18446744073709551615, F1: v76}
                                                                                                    v93 := GenS2{F0: (!false), F1: 100, F2: 2}
                                                                                                    var v94 int = 32768
                                                                                                    _ = v92
                                                                                                    _ = v93
                                                                                                    _ = v94
                                                                                                case -9223372036854775808:
                                                                                                    switch "\n\t" {
                                                                                                    case "\n\t":
                                                                                                        var v95 bool = v87
                                                                                                        _ = v95
                                                                                                    case "":
                                                                                                        var v96 uint8 = (v76 % ((v76) | 1))
                                                                                                        _ = v96
                                                                                                    case "a":
                                                                                                        switch ("\n\t" + "") {
                                                                                                        case "a":
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                        case "":
                                                                                                            v2 -= 0.0
                                                                                                        case "foo bar":
                                                                                                            v77 = false
                                                                                                        default:
                                                                                                            fmt.Println(v69, v77)
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    v88 += (^(v88))
                                                                                                }
                                                                                                fmt.Println(v1, v76)
                                                                                                _ = v90
                                                                                            } else {
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = "\n\t"
                                                                                                }
                                                                                                if (-92 <= 127) {
                                                                                                    v97 := GenS7{F0: 42, F1: 1, F2: fi54}
                                                                                                    v76 -= 1
                                                                                                    switch (fi54 % (((fi54 + (fi54 % ((2147483647) | 1)))) | 1)) {
                                                                                                    case 956:
                                                                                                        if len(v3) > 0 {
                                                                                                            v3[0] = ("\n\t" + ("hello" + "\n\t"))
                                                                                                        }
                                                                                                    case 871:
                                                                                                        fmt.Println(fi54, v4)
                                                                                                    default:
                                                                                                        var v98 uint16 = v5
                                                                                                        _ = v98
                                                                                                    }
                                                                                                    _ = v97
                                                                                                }
                                                                                                switch fi53 {
                                                                                                case 9223372036854775807:
                                                                                                    var v99 float32 = v1
                                                                                                    v100 := GenS5{F0: fi54}
                                                                                                    v101 := GenS8{F0: (v76 / ((uint8(v69)) | 1))}
                                                                                                    _ = v100
                                                                                                    _ = v101
                                                                                                    _ = v99
                                                                                                default:
                                                                                                    var v102 string = ("hello" + "foo bar")
                                                                                                    var v103 int = (fi53 % ((2147483647) | 1))
                                                                                                    _ = v102
                                                                                                    _ = v103
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v3 = append(v3, "a")
                                                                                        _ = v87
                                                                                        _ = v88
                                                                                        _ = v89
                                                                                    }
                                                                                default:
                                                                                    switch fi53 {
                                                                                    case -9223372036854775808:
                                                                                        fmt.Println(v81)
                                                                                    case 1:
                                                                                        v77 = (v77 && (1 == -2147483648))
                                                                                        v81 = 42
                                                                                        v74 = v74
                                                                                    }
                                                                                }
                                                                                if v4 {
                                                                                    switch "a" {
                                                                                    case "\n\t":
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                    default:
                                                                                        switch (-(fi54)) {
                                                                                        case 2147483647:
                                                                                            v3 = append(v3, "")
                                                                                        case -1:
                                                                                            if ((v5 % ((v5) | 1)) == 2) {
                                                                                                fmt.Println(v1)
                                                                                            } else {
                                                                                                if (uint32(v76) <= (^(v69))) {
                                                                                                    v104 := GenS9{F0: 1}
                                                                                                    _ = v104
                                                                                                }
                                                                                                if (941 >= 433) {
                                                                                                    switch "hello" {
                                                                                                    case "a":
                                                                                                        v65 -= 1
                                                                                                    }
                                                                                                    switch (fi53 - (fi53 / ((10) | 1))) {
                                                                                                    case 9223372036854775807:
                                                                                                        if true {
                                                                                                            v105 := GenS8{F0: 42}
                                                                                                            _ = v105
                                                                                                        } else {
                                                                                                            v77 = v4
                                                                                                            var v106 int32 = 1
                                                                                                            _ = v106
                                                                                                        }
                                                                                                    }
                                                                                                    v1 = v1
                                                                                                } else {
                                                                                                    fmt.Println(v4)
                                                                                                    if v4 {
                                                                                                        var v107 uint32 = (v81 << ((0) & 63))
                                                                                                        switch ("foo bar" + "foo bar") {
                                                                                                        case "foo bar":
                                                                                                            v74 -= (v74 >> ((0) & 63))
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                        case "":
                                                                                                            fmt.Println(v74, v74)
                                                                                                        default:
                                                                                                            var v108 float64 = float64(1.5)
                                                                                                            _ = v108
                                                                                                        }
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                        v2 += v2
                                                                                                        _ = v107
                                                                                                    } else {
                                                                                                        v69 = v81
                                                                                                    }
                                                                                                    v81 -= 967
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v65 += 825
                                                                                }
                                                                                var v109 float32 = -2.25
                                                                                _ = v109
                                                                            }
                                                                            _ = v81
                                                                        }
                                                                    case "":
                                                                        if true {
                                                                            var v110 uint16 = 32767
                                                                            var v111 float32 = float32(1e10)
                                                                            _ = v110
                                                                            _ = v111
                                                                        }
                                                                    case "\n\t":
                                                                        if false {
                                                                            var v112 int64 = v74
                                                                            if true {
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    v113 := GenS9{F0: 18446744073709551615}
                                                                                    _ = v113
                                                                                default:
                                                                                    v2 = v2
                                                                                }
                                                                            } else {
                                                                                if v4 {
                                                                                    switch "\n\t" {
                                                                                    case "hello":
                                                                                        v114 := GenS10{F0: 42, F1: "hello", F2: v5}
                                                                                        _ = v114
                                                                                    case "a":
                                                                                        if ((("" + "") + "") >= "hello") {
                                                                                            var v115 float32 = v1
                                                                                            var v116 uint = 18446744073709551615
                                                                                            _ = len(v3)
                                                                                            _ = v115
                                                                                            _ = v116
                                                                                        } else {
                                                                                            if (!(587 <= -1)) {
                                                                                                v117 := GenS11{F0: 2147483647}
                                                                                                var v118 string = "hello"
                                                                                                var v119 uint64 = 245
                                                                                                _ = v117
                                                                                                _ = v118
                                                                                                _ = v119
                                                                                            } else {
                                                                                                var v120 int32 = -185
                                                                                                _ = v120
                                                                                            }
                                                                                            var v121 uint32 = v69
                                                                                            _ = v121
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        fmt.Println(v4)
                                                                                        var v122 uint8 = 1
                                                                                        v77 = (2147483647 > 665)
                                                                                        _ = v122
                                                                                    } else {
                                                                                        var v123 bool = (2 > 923)
                                                                                        v124 := GenS1{F0: 18446744073709551615, F1: 1}
                                                                                        var v125 uint16 = (v5 + v5)
                                                                                        _ = v123
                                                                                        _ = v124
                                                                                        _ = v125
                                                                                    }
                                                                                    if ((true || v4) || (-32768 < -978)) {
                                                                                        _ = len(v3)
                                                                                    }
                                                                                    fmt.Println(fi53, v77)
                                                                                    v77 = v4
                                                                                }
                                                                                if v4 {
                                                                                    fmt.Println(v112, v4, v69)
                                                                                }
                                                                            }
                                                                            var v126 uint32 = 0
                                                                            var v127 float32 = v1
                                                                            var v128 int64 = (v74 & v65)
                                                                            v1 = (v1 * v1)
                                                                            v129 := GenS8{F0: v76}
                                                                            _ = v112
                                                                            _ = v126
                                                                            _ = v127
                                                                            _ = v128
                                                                            _ = v129
                                                                        } else {
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                var v130 bool = v77
                                                                                _ = v130
                                                                            case "hello":
                                                                                var v131 int16 = -80
                                                                                _ = v131
                                                                            case "foo bar":
                                                                                v77 = (1 > 72)
                                                                            }
                                                                        }
                                                                    default:
                                                                        _ = len(v3)
                                                                    }
                                                                case 32767:
                                                                    if v4 {
                                                                        var v132 bool = (!v77)
                                                                        fmt.Println(v5)
                                                                        if true {
                                                                            var v133 int32 = 0
                                                                            v134 := GenS5{F0: 125}
                                                                            _ = v133
                                                                            _ = v134
                                                                        } else {
                                                                            switch (fi53 * (fi54 | 42)) {
                                                                            case 42:
                                                                                var v135 int = fi53
                                                                                _ = v135
                                                                            case -356:
                                                                                var v136 uint64 = 0
                                                                                _ = v136
                                                                            default:
                                                                                switch ("foo bar" + "hello") {
                                                                                case "a":
                                                                                    var v137 float64 = 1e10
                                                                                    _ = v137
                                                                                case "hello":
                                                                                    if v132 {
                                                                                        v138 := GenS4{F0: true}
                                                                                        _ = v138
                                                                                    }
                                                                                }
                                                                            }
                                                                            v139 := GenS1{F0: uint(v76), F1: (v76 | (v76 >> ((274) & 63)))}
                                                                            _ = v139
                                                                        }
                                                                        _ = v132
                                                                    }
                                                                case 523:
                                                                    v3 = append(v3, "\n\t")
                                                                default:
                                                                    fmt.Println(v65, fi54, v69)
                                                                }
                                                            } else {
                                                                switch (-(fi54)) {
                                                                case -314:
                                                                    fmt.Println(v69, v74)
                                                                case 2147483647:
                                                                    var v140 int16 = int16(v76)
                                                                    _ = v140
                                                                }
                                                            }
                                                            switch (fi54 | fi54) {
                                                            case 127:
                                                                var v141 uint64 = 0
                                                                _ = v141
                                                            }
                                                            var v142 int32 = 42
                                                            _ = v142
                                                        }
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    _ = v74
                                                }
                                            }
                                        }
                                    }
                                    _ = v69
                                }
                                var v143 string = "a"
                                var v144 int32 = 503
                                fmt.Println(fi54)
                                _ = v143
                                _ = v144
                            }
                            switch 42 {
                            case 128:
                                v1 = v1
                            case 32768:
                                var v145 uint32 = 2147483647
                                _ = v145
                            default:
                                switch "a" {
                                case "":
                                    v3 = append(v3, "a")
                                default:
                                    v65 -= (v65 % (((-(v65))) | 1))
                                }
                                switch ((("hello" + "a") + "hello") + (("hello" + "") + ("hello" + "foo bar"))) {
                                case "\n\t":
                                    if v4 {
                                        var v146 bool = ((!(!(2 > 1))) && v4)
                                        v65 = (-(v65))
                                        if (2 > 65535) {
                                            fmt.Println(v5)
                                            var v147 uint = uint(v5)
                                            v5 = (v5 << ((uint(v5)) & 63))
                                            var v148 uint8 = uint8(v5)
                                            _ = v147
                                            _ = v148
                                        }
                                        v5 -= (v5 & v5)
                                        _ = v146
                                    } else {
                                        switch "" {
                                        case "":
                                            var v149 float64 = 3.14159
                                            _ = v149
                                        default:
                                            switch "\n\t" {
                                            case "a":
                                                switch (-(fi54)) {
                                                case 32767:
                                                    if ((v2 * (0.0 * (v2 + (0.1 / (v2 + (((0.0 * (1e10 - 3.14159)) - (v2 / v2)) / float64(float32(v2)))))))) >= v2) {
                                                        if len(v3) > 0 {
                                                            v3[0] = "hello"
                                                        }
                                                        v4 = ((int32(fi54) != -2147483648) && false)
                                                        if (127 > 54) {
                                                            v65 = -9223372036854775808
                                                            var v150 int16 = -1
                                                            _ = v150
                                                        } else {
                                                            v5 = (v5 * v5)
                                                            var v151 uint = 1
                                                            _ = v151
                                                        }
                                                    } else {
                                                        v152 := GenS9{F0: 182}
                                                        v1 = float32(0.0)
                                                        _ = v152
                                                    }
                                                default:
                                                    v65 = v65
                                                }
                                            case "\n\t":
                                                v153 := GenS10{F0: -952, F1: "hello", F2: 652}
                                                _ = v153
                                            case "hello":
                                                switch fi54 {
                                                case 4294967295:
                                                    if ((("\n\t" + "hello") + "hello") >= "a") {
                                                        v65 -= v65
                                                        switch fi54 {
                                                        case 166:
                                                            v1 += 1.5
                                                        case 1:
                                                            v4 = (false || v4)
                                                        case 32768:
                                                            var v154 int32 = 2147483647
                                                            _ = v154
                                                        }
                                                        v4 = (!v4)
                                                    } else {
                                                        fmt.Println(v4, v5)
                                                        v155 := GenS12{F0: 128}
                                                        _ = v155
                                                    }
                                                }
                                            }
                                        }
                                    }
                                case "":
                                    fmt.Println(fi54, v4, v5)
                                case "a":
                                    var v156 int16 = -779
                                    _ = v156
                                }
                                v157 := GenS13{F0: 42, F1: -463, F2: 2147483647}
                                fmt.Println(fi53, v5, v5)
                                switch fi54 {
                                case 677:
                                    var v158 string = ("\n\t" + "")
                                    _ = v158
                                default:
                                    switch (fi54 - (fi53 % ((1) | 1))) {
                                    case -1:
                                        fmt.Println(v5)
                                    case 512:
                                        v159 := GenS5{F0: (fi54 ^ 32767)}
                                        _ = v159
                                    default:
                                        var v160 uint64 = 18446744073709551615
                                        _ = v160
                                    }
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                fmt.Println(fi54, fi53, v5)
                                if true {
                                    v2 = (3.14159 + float64(float32((((0.0 * v2) + v2) * v2))))
                                    var v161 uint16 = v5
                                    var v162 string = ""
                                    var v163 uint16 = 296
                                    _ = v161
                                    _ = v162
                                    _ = v163
                                }
                                _ = v157
                            }
                            _ = v65
                        } else {
                            var v164 uint32 = 204
                            if ((v2 / v2) == v2) {
                                var v165 float32 = (((v1 - 1.5) * v1) + v1)
                                var v166 uint32 = v164
                                var v167 int64 = 9223372036854775807
                                _ = len(v3)
                                _ = v165
                                _ = v166
                                _ = v167
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = "a"
                                }
                            }
                            v168 := GenS14{F0: -128, F1: (!(uint64(v5) >= 1)), F2: -9223372036854775808}
                            _ = v164
                            _ = v168
                        }
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
                var v169 float64 = (v2 - (v2 * -2.25))
                _ = fi54
                _ = v169
            }
            _ = fi53
        }
    }
    v1 += -2.25
    v4 = (!true)
    v4 = (!v4)
    if v4 {
        var v170 int16 = 1
        _ = v170
    }
    var v171 float64 = 0.1
    var v172 string = "a"
    var v173 float32 = v1
    fmt.Println(v173)
    fmt.Println("v1:", v1)
    fmt.Println("v171:", v171)
    fmt.Println("v172:", v172)
    fmt.Println("v173:", v173)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v171
    _ = v172
    _ = v173
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
