package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS3 struct { F0 string; F1 uint8; F2 int }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 int64; F1 string }
type GenS9 struct { F0 uint64; F1 uint }
type GenS10 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint16; F1 bool; F2 int8 }
type GenS13 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS14 struct { F0 int16; F1 int; F2 int64 }
type GenS15 struct { F0 uint32; F1 int64; F2 int64 }
type GenS16 struct { F0 uint; F1 uint16; F2 bool }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"a", "hello", "hello", "hello"}
    var v3 uint16 = 610
    if (-693 > -9223372036854775808) {
        var v4 string = "foo bar"
        for fi5 := 0; fi5 < 20; fi5++ {
            switch fi5 {
            case -1:
                if true {
                    var v6 float64 = (1e10 - 0.1)
                    v6 = (float64(3.14159) / v6)
                    _ = v6
                } else {
                    v4 = v4
                }
            }
            if false {
                for fi7 := 0; fi7 < 1; fi7++ {
                    fmt.Println(fi5, fi7)
                    var v8 float64 = 1e10
                    _ = fi7
                    _ = v8
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                var v9 uint32 = 2147483648
                v10 := GenS1{F0: true, F1: 42, F2: 42}
                v3 = (v3 >> ((65535) & 63))
                _ = v10
                _ = v9
            }
            _ = fi5
        }
        _ = len(v1)
        _ = v4
    } else {
        switch 127 {
        case -1:
            switch 42 {
            case 9223372036854775807:
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        fmt.Println(v3)
    }
    v1[468] = 25
    if false {
        var v11 uint8 = 255
        _ = v11
    }
    v12 := GenS2{F0: 1, F1: 18446744073709551615, F2: 42}
    var v13 uint8 = 211
    if (!true) {
        v1[1] = 4294967295
        switch -678 {
        case 42:
            switch 4294967295 {
            case -566:
                v14 := GenS1{F0: false, F1: 32767, F2: 157}
                _ = v14
            }
        case -882:
            var v15 uint = 0
            _ = v15
        case -1:
            if (0 != 2147483647) {
                v16 := GenS3{F0: "a", F1: v13, F2: -292}
                var v17 uint = 2
                _ = v16
                _ = v17
            }
        default:
            fmt.Println(v3, v13)
        }
    } else {
        fmt.Println(v13, v3)
    }
    var v18 int16 = 176
    v3 = (v3 * (v3 | (v3 << ((0) & 63))))
    for fi19 := 0; fi19 < 6; fi19++ {
        if false {
            delete(v1, 42)
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 string = (("\n\t" + "") + "\n\t")
                var v22 float64 = -2.25
                _ = len(v2)
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        _ = fi19
    }
    var v23 float64 = 1e10
    for fi24 := 0; fi24 < 19; fi24++ {
        switch "" {
        case "a":
            switch "" {
            case "foo bar":
                if false {
                    v18 = v18
                    v18 += (-(v18))
                    for fi25 := 0; fi25 < 2; fi25++ {
                        switch "hello" {
                        case "\n\t":
                            if true {
                                fmt.Println(v13)
                                switch "\n\t" {
                                case "a":
                                    if (-2147483648 == 2147483647) {
                                        switch 4294967295 {
                                        case 0:
                                            if true {
                                                v3 = (v3 % ((2) | 1))
                                                if false {
                                                    v18 = v18
                                                    v13 = 128
                                                }
                                                fmt.Println(fi24, fi25, v3)
                                            } else {
                                                var v26 string = ("" + ("foo bar" + "a"))
                                                var v27 int32 = 746
                                                fmt.Println(v23, v18, v13)
                                                _ = v26
                                                _ = v27
                                            }
                                            v23 += 1e10
                                        case 898:
                                            fmt.Println(v13, fi24, v13)
                                            v28 := GenS4{F0: -1, F1: 124}
                                            if (789 >= 459) {
                                                switch fi25 {
                                                case 4294967295:
                                                    switch "foo bar" {
                                                    case "":
                                                        fmt.Println(fi24)
                                                    default:
                                                        v3 += 65535
                                                    }
                                                case 64:
                                                    if (500 == 42) {
                                                        if true {
                                                            v29 := GenS3{F0: "", F1: v13, F2: fi25}
                                                            v1[1] = 914
                                                            var v30 string = "a"
                                                            _ = v29
                                                            _ = v30
                                                        } else {
                                                            v31 := GenS3{F0: ("hello" + ((("\n\t" + "\n\t") + (("" + "foo bar") + ("" + "foo bar"))) + "hello")), F1: (v13 % (((v13 - (^(v13)))) | 1)), F2: (^(fi24))}
                                                            switch (fi25 >> ((916) & 63)) {
                                                            case -2147483648:
                                                                if true {
                                                                    v32 := GenS5{F0: v18, F1: (v3 - 175)}
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        v23 = (1e10 + float64((1e10 * ((1e10 * float32(v23)) / (0.0 / (-2.25 - ((1.5 + (((-2.25 + ((0.0 / 0.1) * (float32(float64(1e10)) + 1e10))) * 0.0) - float32((v23 + v23)))) + 0.1)))))))
                                                                    case "a":
                                                                        v33 := GenS6{F0: (v13 + 255), F1: (^(fi25))}
                                                                        _ = v33
                                                                    case "\n\t":
                                                                        fmt.Println(fi24, fi25, v18)
                                                                    default:
                                                                        var v34 uint16 = v3
                                                                        _ = v34
                                                                    }
                                                                    var v35 uint16 = 1
                                                                    _ = v32
                                                                    _ = v35
                                                                } else {
                                                                    if false {
                                                                        if true {
                                                                            var v36 float64 = (1.5 - (((v23 - (float64((3.14159 / 1e10)) / (v23 / v23))) * 1.5) * ((v23 * v23) + v23)))
                                                                            v37 := GenS6{F0: 1, F1: 69}
                                                                            v23 -= ((v36 - 0.1) + v23)
                                                                            _ = v36
                                                                            _ = v37
                                                                        } else {
                                                                            var v38 bool = false
                                                                            fmt.Println(v38, v18, v23)
                                                                            if false {
                                                                                var v39 int = (^(fi25))
                                                                                var v40 float64 = v23
                                                                                fmt.Println(v40)
                                                                                _ = v39
                                                                                _ = v40
                                                                            }
                                                                            _ = v38
                                                                        }
                                                                        switch (("" + "") + "a") {
                                                                        case "\n\t":
                                                                            v13 += 127
                                                                        case "hello":
                                                                            fmt.Println(v3, v18, fi25)
                                                                        default:
                                                                            if false {
                                                                                switch 0 {
                                                                                case 11:
                                                                                    switch (-(fi25)) {
                                                                                    case -90:
                                                                                        v23 = 0.0
                                                                                    case -375:
                                                                                        var v41 uint32 = 1
                                                                                        _ = v41
                                                                                    case -2147483648:
                                                                                        v13 = (^(v13))
                                                                                    }
                                                                                case 9223372036854775807:
                                                                                    var v42 int = (fi24 / (((fi25 >> ((246) & 63))) | 1))
                                                                                    _ = v42
                                                                                }
                                                                                v43 := GenS3{F0: ("\n\t" + "hello"), F1: v13, F2: int(v18)}
                                                                                switch (fi25 + fi25) {
                                                                                case 301:
                                                                                    var v44 int16 = (v18 << ((18446744073709551615) & 63))
                                                                                    _ = v44
                                                                                case 32768:
                                                                                    v45 := GenS5{F0: v18, F1: 337}
                                                                                    _ = v45
                                                                                }
                                                                                _ = v43
                                                                            }
                                                                        }
                                                                        fmt.Println(v13)
                                                                    } else {
                                                                        if false {
                                                                            v46 := GenS7{F0: 4294967295, F1: 856}
                                                                            switch fi25 {
                                                                            case 0:
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            default:
                                                                                var v47 float32 = (1.5 + (float32((v23 * v23)) - 3.14159))
                                                                                if ((!(18446744073709551615 > 18446744073709551615)) || (false || false)) {
                                                                                    fmt.Println(v13)
                                                                                }
                                                                                var v48 bool = false
                                                                                _ = v47
                                                                                _ = v48
                                                                            }
                                                                            _ = v46
                                                                        } else {
                                                                            v23 = v23
                                                                        }
                                                                    }
                                                                    var v49 uint32 = 4294967295
                                                                    _ = v49
                                                                }
                                                            case 4294967295:
                                                                v13 += (v13 - (v13 << ((724) & 63)))
                                                            case 128:
                                                                if ((fi25 >> ((255) & 63)) <= (fi25 | fi25)) {
                                                                    v3 += v3
                                                                    v2 = append(v2, "a")
                                                                    if true {
                                                                        var v50 bool = (("\n\t" + "\n\t") >= "hello")
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        _ = v50
                                                                    }
                                                                    v3 = v3
                                                                    v13 = v13
                                                                }
                                                            default:
                                                                if false {
                                                                    v3 -= (v3 * (v3 - v3))
                                                                } else {
                                                                    v23 = 0.0
                                                                    if (fi24 >= fi25) {
                                                                        var v51 int8 = -99
                                                                        _ = v51
                                                                    } else {
                                                                        delete(v1, 42)
                                                                    }
                                                                }
                                                            }
                                                            var v52 bool = false
                                                            _ = v31
                                                            _ = v52
                                                        }
                                                        v3 += v3
                                                        v53 := GenS8{F0: 1, F1: "a"}
                                                        _ = v53
                                                    }
                                                case 127:
                                                    v13 = 128
                                                }
                                                var v54 uint16 = v3
                                                var v55 uint8 = 127
                                                v55 -= (v55 | (^(v13)))
                                                v56 := GenS9{F0: 951, F1: 0}
                                                v55 -= 128
                                                fmt.Println(fi24, v3)
                                                _ = v54
                                                _ = v55
                                                _ = v56
                                            } else {
                                                var v57 int32 = -1
                                                _ = v57
                                            }
                                            _ = v28
                                        case -2147483648:
                                            if len(v2) > 0 {
                                                v2[0] = ("hello" + "a")
                                            }
                                        }
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v58 uint8 = 1
                                            if true {
                                                if (true && false) {
                                                    v59 := GenS7{F0: 866, F1: 395}
                                                    v3 += (v3 ^ v3)
                                                    _ = v59
                                                }
                                                switch fi24 {
                                                case 42:
                                                    fmt.Println(v58)
                                                case -1:
                                                    v60 := GenS5{F0: (v18 / ((v18) | 1)), F1: (^(v3))}
                                                    _ = v60
                                                case 128:
                                                    v58 = (v13 / (((v58 | 127)) | 1))
                                                default:
                                                    v61 := GenS3{F0: "a", F1: (v58 - (v13 + 231)), F2: (-(fi25))}
                                                    _ = v61
                                                }
                                            }
                                            _ = v58
                                        case "a":
                                            var v62 string = "foo bar"
                                            _ = v62
                                        case "hello":
                                            if (true || (true && true)) {
                                                v13 = 216
                                                if (fi25 >= fi24) {
                                                    v18 = (v18 * v18)
                                                    _ = len(v2)
                                                    var v63 uint16 = v3
                                                    var v64 uint32 = 2
                                                    _ = v63
                                                    _ = v64
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = (("\n\t" + "") + "hello")
                                                }
                                            } else {
                                                switch (-(fi25)) {
                                                case 1:
                                                    v2 = append(v2, "")
                                                case -2147483648:
                                                    switch "hello" {
                                                    case "":
                                                        var v65 uint64 = 42
                                                        if false {
                                                            if ((v65 + (v65 * 42)) > v65) {
                                                                switch ("\n\t" + "a") {
                                                                case "\n\t":
                                                                    if ((v13 / (((^(v13))) | 1)) != v13) {
                                                                        v66 := GenS2{F0: 2, F1: v65, F2: (v13 >> ((438) & 63))}
                                                                        _ = v66
                                                                    } else {
                                                                        v67 := GenS1{F0: (v13 > 1), F1: (v18 - 42), F2: (v65 ^ 396)}
                                                                        if (false || (true && true)) {
                                                                            var v68 uint64 = (^(v65))
                                                                            fmt.Println(v3)
                                                                            fmt.Println(v68, v68, v68)
                                                                            _ = v68
                                                                        }
                                                                        _ = v67
                                                                    }
                                                                }
                                                                var v69 bool = true
                                                                _ = v69
                                                            }
                                                            switch ("hello" + "\n\t") {
                                                            case "":
                                                                v3 = v3
                                                            }
                                                            var v70 uint = 255
                                                            _ = v70
                                                        }
                                                        _ = v65
                                                    }
                                                }
                                                if (!(!(4294967295 < 0))) {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                v71 := GenS5{F0: (v18 >> ((65535) & 63)), F1: v3}
                                                var v72 int16 = (^(v18))
                                                var v73 uint16 = 255
                                                v3 -= 255
                                                _ = v71
                                                _ = v72
                                                _ = v73
                                            }
                                        default:
                                            v74 := GenS10{F0: v13, F1: v13, F2: 65535}
                                            _ = v74
                                        }
                                        v1[uint32(v13)] = 2147483647
                                    } else {
                                        v13 += (^(v13))
                                    }
                                }
                            } else {
                                switch ((("hello" + ("\n\t" + ("foo bar" + ""))) + ("a" + "foo bar")) + "\n\t") {
                                case "\n\t":
                                    fmt.Println(v13, v13)
                                }
                                var v75 uint = 65535
                                v23 = (v23 * v23)
                                _ = v75
                            }
                        case "foo bar":
                            v3 = v3
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        _ = fi25
                    }
                } else {
                    fmt.Println(fi24, v3, fi24)
                }
            case "a":
                fmt.Println(v3, v18, v3)
            case "":
                var v76 uint8 = v13
                _ = v76
            default:
                var v77 int32 = 42
                _ = v77
            }
        case "foo bar":
            _ = len(v2)
        }
        if (138 < 116) {
            switch "" {
            case "":
                if (true && true) {
                    _ = len(v2)
                    for fi78 := 0; fi78 < 17; fi78++ {
                        switch -2147483648 {
                        case -704:
                            if (v3 < 65535) {
                                if ((1 != v18) || true) {
                                    v23 = v23
                                    if false {
                                        if false {
                                            if false {
                                                var v79 uint = 310
                                                _ = v79
                                            }
                                            var v80 float32 = (1.5 / 0.1)
                                            _ = v80
                                        } else {
                                            v81 := GenS11{F0: 0}
                                            _ = v81
                                        }
                                        _ = len(v1)
                                    }
                                } else {
                                    fmt.Println(fi78, v3, v18)
                                }
                                v13 += v13
                                if (v18 != (-(v18))) {
                                    if false {
                                        if false {
                                            switch "a" {
                                            case "foo bar":
                                                var v82 int = (fi78 >> ((1) & 63))
                                                _ = v82
                                            case "hello":
                                                var v83 uint8 = 2
                                                _ = v83
                                            }
                                            var v84 uint32 = 0
                                            _ = v84
                                        }
                                    }
                                    var v85 uint32 = 1
                                    _ = v85
                                } else {
                                    var v86 int8 = -1
                                    v1[179] = 0
                                    _ = v86
                                }
                                if (true || (false || (-2.25 > (0.1 - 1.5)))) {
                                    fmt.Println(v3, v18)
                                    v87 := GenS12{F0: (^(v3)), F1: ((!(v13 < 0)) && (true && (v3 <= 438))), F2: -60}
                                    v88 := GenS13{F0: v3, F1: 298, F2: v3}
                                    _ = v87
                                    _ = v88
                                } else {
                                    v23 = (v23 / float64(0.1))
                                }
                                delete(v1, 2147483648)
                            } else {
                                v3 += (v3 * v3)
                                fmt.Println(v18, v23, v3)
                                var v89 int16 = 1000
                                _ = v89
                            }
                        case 9223372036854775807:
                            switch int(v13) {
                            case 1:
                                v1[2] = uint32(v13)
                            case 9223372036854775807:
                                v90 := GenS14{F0: -883, F1: (fi78 << ((253) & 63)), F2: 9223372036854775807}
                                _ = v90
                            case -9223372036854775808:
                                switch ("" + "") {
                                case "":
                                    switch "foo bar" {
                                    case "a":
                                        v3 = (v3 >> ((714) & 63))
                                    case "foo bar":
                                        v3 -= v3
                                        var v91 uint64 = 2
                                        _ = v91
                                    case "":
                                        var v92 float64 = v23
                                        _ = v92
                                    }
                                    v1[65535] = 0
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = ("a" + ("foo bar" + ""))
                                    }
                                    if false {
                                        fmt.Println(v18, fi78)
                                        var v93 uint8 = v13
                                        _ = v93
                                    }
                                default:
                                    v13 = (v13 - (^(v13)))
                                }
                            default:
                                v1[4294967295] = 322
                            }
                        default:
                            _ = len(v1)
                            fmt.Println(fi78, v3)
                        }
                        switch "a" {
                        case "\n\t":
                            v94 := GenS14{F0: v18, F1: int(v13), F2: -1}
                            _ = v94
                        case "foo bar":
                            v13 = v13
                        default:
                            v18 -= -928
                        }
                        _ = fi78
                    }
                } else {
                    if (!(v23 <= v23)) {
                        v95 := GenS15{F0: 0, F1: 224, F2: -1}
                        v13 -= (v13 >> ((476) & 63))
                        _ = v95
                    }
                }
            case "\n\t":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            var v96 float64 = (float64(0.0) * 3.14159)
            v97 := GenS16{F0: 65535, F1: 255, F2: false}
            _ = v96
            _ = v97
        }
        v98 := GenS15{F0: 65535, F1: 503, F2: 0}
        _ = fi24
        _ = v98
    }
    for fi99 := 0; fi99 < 18; fi99++ {
        switch (fi99 | (fi99 ^ (fi99 + 127))) {
        case 946:
            for fi100 := 0; fi100 < 12; fi100++ {
                var v101 uint8 = v13
                fmt.Println(v101, v13, fi99)
                v101 -= 2
                _ = fi100
                _ = v101
            }
        default:
            _ = len(v2)
        }
        var v102 int16 = 42
        _ = fi99
        _ = v102
    }
    fmt.Println(v18, v13, v18)
    v13 += v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v23
    _ = v3
}
