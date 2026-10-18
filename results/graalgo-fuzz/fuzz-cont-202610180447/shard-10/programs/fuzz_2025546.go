package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint; F1 uint8; F2 bool }
type GenS4 struct { F0 uint16; F1 uint; F2 int32 }
type GenS5 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 uint; F1 int32 }
type GenS9 struct { F0 bool; F1 int32; F2 uint64 }
type GenS10 struct { F0 int8; F1 uint16 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint }
func genFn0() int8 {
    switch -2147483648 {
    case 128:
        var v7 float32 = 0.1
        _ = v7
    default:
        var v8 uint8 = 0
        _ = v8
    }
    if true {
        var v9 int8 = 127
        switch "foo bar" {
        case "foo bar":
            v9 += v9
        }
        _ = v9
    }
    return 100
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 0.0
    v3 := []string{}
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]uint32)
    var v6 uint = 632
    for fi10 := 0; fi10 < 18; fi10++ {
        _ = len(v3)
        if (true && (false && true)) {
            v11 := GenS1{F0: 2, F1: v6, F2: -1}
            v12 := GenS2{F0: v6, F1: 4294967295}
            _ = v11
            _ = v12
        } else {
            var v13 uint32 = 2147483648
            _ = v13
        }
        v14 := v4[602]
        _ = fi10
        _ = v14
    }
    if (4294967295 > 2) {
        for fi15 := 0; fi15 < 16; fi15++ {
            if false {
                for fi16 := 0; fi16 < 13; fi16++ {
                    var v17 int64 = -524
                    v18 := GenS1{F0: 1, F1: (v6 >> ((v6) & 63)), F2: -2147483648}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 11; fi19++ {
                v6 = v6
                if (42 == -9223372036854775808) {
                    fmt.Println(v2, fi19, fi19)
                } else {
                    var v20 uint64 = 42
                    var v21 uint32 = 68
                    if (-9223372036854775808 >= 210) {
                        fmt.Println(fi15, fi19, v2)
                        switch (fi15 >> ((v6) & 63)) {
                        case -1:
                            v22 := GenS1{F0: 2, F1: v6, F2: 2147483647}
                            _ = v22
                        case 32767:
                            fmt.Println(v20, fi19, fi19)
                        case -12:
                            var v23 float64 = (0.1 / float64(v2))
                            _ = v23
                        default:
                            var v24 uint32 = (v21 & (v21 << (((v6 - (v6 | (v6 | v6)))) & 63)))
                            _ = v24
                        }
                    } else {
                        v21 = v21
                        v25 := v1[0]
                        _ = v25
                    }
                    _ = v20
                    _ = v21
                }
                v6 = (v6 / (((^(v6))) | 1))
                var v26 int32 = -1
                _ = fi19
                _ = v26
            }
            _ = fi15
        }
        fmt.Println(v6)
    }
    var v27 int32 = -2
    var v28 string = ""
    if (true && (32767 == 32767)) {
        v29 := GenS2{F0: 2, F1: (v6 - v6)}
        if (!(32767 != 0)) {
            for fi30 := 0; fi30 < 5; fi30++ {
                fmt.Println(fi30)
                switch v28 {
                case "foo bar":
                    _ = len(v3)
                }
                _ = fi30
            }
            for fi31 := 0; fi31 < 16; fi31++ {
                switch 2147483647 {
                case 369:
                    var v32 int = fi31
                    _ = v32
                }
                v33 := GenS2{F0: (v6 - (v6 / (((v6 / ((v6) | 1))) | 1))), F1: (v6 - v6)}
                v34 := GenS1{F0: 194, F1: v6, F2: 2147483647}
                _ = fi31
                _ = v33
                _ = v34
            }
            for fi35 := 0; fi35 < 8; fi35++ {
                var v36 int = -9223372036854775808
                _ = fi35
                _ = v36
            }
        }
        v27 += 0
        _ = v29
    }
    switch v28 {
    case "\n\t":
        if ((((float32(3.14159) * 0.1) - (v2 * v2)) - v2) > v2) {
            v37 := GenS1{F0: 2, F1: v6, F2: -2147483648}
            delete(v4, 2)
            for fi38 := 0; fi38 < 15; fi38++ {
                _ = len(v3)
                for fi39 := 0; fi39 < 2; fi39++ {
                    fmt.Println(v2, fi38, fi39)
                    switch "hello" {
                    case "":
                        if true {
                            v2 = v2
                            v40 := GenS1{F0: 63, F1: v6, F2: (v27 * 666)}
                            _ = len(v5)
                            _ = v40
                        }
                    case "a":
                        switch (fi38 + (fi38 + fi39)) {
                        case -2147483648:
                            v41 := GenS3{F0: (v6 | (^(v6))), F1: 255, F2: true}
                            _ = v41
                        case 0:
                            v42 := GenS4{F0: 1, F1: (v6 - v6), F2: v27}
                            _ = v42
                        case 32768:
                            v28 = v28
                        }
                    case "foo bar":
                        v43 := GenS5{F0: 127, F1: 42, F2: v27}
                        if true {
                            _ = len(v1)
                            v44 := GenS6{F0: 18446744073709551615}
                            _ = v44
                        } else {
                            v2 = v2
                            if (!false) {
                                v45 := GenS5{F0: 161, F1: 1, F2: (v27 << ((v6) & 63))}
                                fmt.Println(v27)
                                switch fi39 {
                                case -1:
                                    var v46 uint8 = 128
                                    _ = v46
                                case -624:
                                    v47 := GenS7{F0: "foo bar", F1: 42}
                                    v6 = (v6 * (v6 & v6))
                                    _ = v47
                                case 499:
                                    v28 = ((((("" + (v28 + v28)) + ((v28 + (((v28 + v28) + (v28 + v28)) + ("foo bar" + "\n\t"))) + v28)) + v28) + (((v28 + v28) + (v28 + (v28 + (v28 + ("\n\t" + v28))))) + v28)) + ((("\n\t" + (("\n\t" + v28) + (("\n\t" + (v28 + v28)) + (v28 + "\n\t")))) + v28) + (v28 + (v28 + "hello"))))
                                default:
                                    var v48 uint64 = 1
                                    _ = v48
                                }
                                _ = v45
                            } else {
                                var v49 float32 = float32(((3.14159 - 1e10) * 0.0))
                                _ = v49
                            }
                            if true {
                                v50 := GenS7{F0: "", F1: 814}
                                _ = v50
                            }
                        }
                        switch fi39 {
                        case -330:
                            fmt.Println(v2, fi38)
                        case 32767:
                            var v51 uint32 = 65535
                            _ = v51
                        }
                        _ = v43
                    default:
                        v52 := GenS6{F0: uint64(v6)}
                        _ = v52
                    }
                    v3 = append(v3, v28)
                    _ = fi39
                }
                _ = fi38
            }
            _ = v37
        } else {
            var v53 uint = (v6 & (^(v6)))
            delete(v1, 808)
            for fi54 := 0; fi54 < 3; fi54++ {
                fmt.Println(fi54)
                _ = fi54
            }
            _ = v53
        }
    }
    switch ("a" + v28) {
    case "":
        v55 := v5[2147483647]
        _ = v55
    case "\n\t":
        if ((("hello" + "") <= v28) && false) {
            v28 = v28
            if true {
                var v56 uint8 = 2
                if (4294967295 == 2) {
                    var v57 uint = v6
                    _ = v57
                }
                _ = v56
            }
        }
    }
    for fi58 := 0; fi58 < 3; fi58++ {
        fmt.Println(v6, v28, fi58)
        _ = fi58
    }
    if (float64((((v2 * (v2 - v2)) + v2) + (v2 - v2))) == 0.0) {
        v59 := GenS8{F0: (v6 % (((v6 ^ 0)) | 1)), F1: (v27 ^ -663)}
        switch 0 {
        case 628:
            delete(v4, 942)
        default:
            if (980 >= v27) {
                switch -9223372036854775808 {
                case 9223372036854775807:
                    switch v28 {
                    case "\n\t":
                        fmt.Println(v2, v28)
                    case "a":
                        v5[4294967295] = 802
                    case "hello":
                        v60 := GenS4{F0: 42, F1: v6, F2: v27}
                        _ = v60
                    default:
                        v1[2] = (!(-2.25 == 0.1))
                    }
                case 273:
                    switch v28 {
                    case "hello":
                        v2 = float32(1e10)
                    default:
                        v6 -= 841
                    }
                default:
                    if (((false || (9223372036854775807 >= 934)) && (3.14159 != ((3.14159 - v2) - (float32(((0.1 * 0.1) + (0.1 + 0.1))) + (float32(((0.1 - 3.14159) / 0.1)) - 0.0))))) && (2 <= 227)) {
                        fmt.Println(v2, v28, v27)
                    } else {
                        var v61 int8 = 120
                        v61 = 1
                        _ = v61
                    }
                }
            } else {
                if true {
                    var v62 int64 = -860
                    fmt.Println(v62)
                    for fi63 := 0; fi63 < 6; fi63++ {
                        switch ((v28 + v28) + "hello") {
                        case "":
                            var v64 float32 = ((v2 / (1.5 - (float32((((1e10 - (-2.25 * 0.0)) / 0.1) / (-2.25 - 0.1))) * v2))) + 0.1)
                            _ = v64
                        case "hello":
                            v65 := GenS5{F0: 0, F1: 1, F2: v27}
                            _ = v65
                        }
                        _ = fi63
                    }
                    _ = v62
                } else {
                    v66 := GenS8{F0: 436, F1: v27}
                    if true {
                        switch 0 {
                        case 195:
                            switch 707 {
                            case 32768:
                                fmt.Println(v6)
                            case 0:
                                var v67 uint16 = 65535
                                _ = v67
                            case 42:
                                for fi68 := 0; fi68 < 8; fi68++ {
                                    if (-973 < 294) {
                                        var v69 float32 = 1.5
                                        fmt.Println(v27, v27, v2)
                                        _ = v69
                                    }
                                    delete(v5, 116)
                                    _ = fi68
                                }
                            default:
                                v27 -= v27
                            }
                        }
                        for fi70 := 0; fi70 < 18; fi70++ {
                            v6 -= v6
                            switch (fi70 - 4294967295) {
                            case 0:
                                for fi71 := 0; fi71 < 16; fi71++ {
                                    fmt.Println(fi71)
                                    switch v28 {
                                    case "a":
                                        v27 += (^(v27))
                                    default:
                                        switch (v28 + ("a" + "foo bar")) {
                                        case "a":
                                            var v72 uint64 = 957
                                            _ = v72
                                        case "hello":
                                            var v73 uint8 = 2
                                            _ = v73
                                        default:
                                            v74 := GenS9{F0: ((!false) || ((383 >= (fi71 & int(v6))) && false)), F1: v27, F2: 42}
                                            _ = v74
                                        }
                                    }
                                    var v75 float64 = -2.25
                                    if true {
                                        var v76 float32 = v2
                                        _ = v76
                                    } else {
                                        switch v28 {
                                        case "foo bar":
                                            v75 += v75
                                        case "a":
                                            v2 += (float32(v75) + v2)
                                        default:
                                            v77 := GenS8{F0: v6, F1: 667}
                                            _ = v77
                                        }
                                        if ((v6 ^ v6) >= v6) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if (0 < 127) {
                                                v28 = v28
                                            } else {
                                                v27 = v27
                                            }
                                            if true {
                                                var v78 float64 = 0.0
                                                v79 := GenS6{F0: 42}
                                                v27 += v27
                                                _ = v78
                                                _ = v79
                                            } else {
                                                switch (v28 + v28) {
                                                case "a":
                                                    var v80 int64 = 1
                                                    _ = v80
                                                case "foo bar":
                                                    var v81 int16 = 1
                                                    _ = v81
                                                case "\n\t":
                                                    switch fi70 {
                                                    case 1:
                                                        var v82 int8 = 127
                                                        _ = v82
                                                    }
                                                }
                                                v83 := GenS4{F0: 32767, F1: v6, F2: -955}
                                                delete(v5, 42)
                                                v28 = ((v28 + v28) + "a")
                                                _ = v83
                                            }
                                        } else {
                                            fmt.Println(v27, v75, v28)
                                            v27 = (v27 / ((v27) | 1))
                                        }
                                    }
                                    _ = fi71
                                    _ = v75
                                }
                            case -453:
                                for fi84 := 0; fi84 < 18; fi84++ {
                                    v6 -= v6
                                    v85 := GenS2{F0: v6, F1: (v6 % (((v6 + (v6 << ((v6) & 63)))) | 1))}
                                    v86 := GenS10{F0: -128, F1: 65535}
                                    _ = fi84
                                    _ = v85
                                    _ = v86
                                }
                            }
                            _ = fi70
                        }
                    }
                    v27 = (v27 ^ -1)
                    _ = v66
                }
                switch v28 {
                case "hello":
                    var v87 uint64 = 1
                    _ = v87
                case "foo bar":
                    var v88 uint16 = 161
                    _ = v88
                case "a":
                    for fi89 := 0; fi89 < 8; fi89++ {
                        for fi90 := 0; fi90 < 7; fi90++ {
                            var v91 int8 = 127
                            _ = fi90
                            _ = v91
                        }
                        _ = fi89
                    }
                }
                fmt.Println(v2, v27)
            }
        }
        for fi92 := 0; fi92 < 12; fi92++ {
            switch (((("" + v28) + v28) + "a") + "") {
            case "foo bar":
                v2 = 0.0
            default:
                v93 := GenS9{F0: (!true), F1: (-(v27)), F2: 42}
                _ = v93
            }
            if false {
                if (!((42 > 127) || false)) {
                    for fi94 := 0; fi94 < 13; fi94++ {
                        v6 -= (v6 >> ((v6) & 63))
                        if false {
                            _ = len(v5)
                        }
                        v95 := GenS11{F0: true}
                        _ = fi94
                        _ = v95
                    }
                    v3 = append(v3, (v28 + ("\n\t" + "foo bar")))
                    for fi96 := 0; fi96 < 11; fi96++ {
                        var v97 bool = true
                        var v98 uint64 = 1
                        _ = fi96
                        _ = v97
                        _ = v98
                    }
                } else {
                    _ = len(v3)
                }
                if (false && (!true)) {
                    v4[1] = v28
                    if (v28 >= ("" + "hello")) {
                        v28 = ""
                        fmt.Println(v2)
                    } else {
                        if ((!true) || true) {
                            switch 9223372036854775807 {
                            case 32768:
                                v6 = (v6 >> (((v6 / ((v6) | 1))) & 63))
                            case 710:
                                fmt.Println(v28, v27)
                            case -9223372036854775808:
                                v28 = v28
                            default:
                                switch "foo bar" {
                                case "":
                                    if false {
                                        for fi99 := 0; fi99 < 12; fi99++ {
                                            if true {
                                                var v100 uint32 = 617
                                                v101 := GenS10{F0: -1, F1: 53}
                                                var v102 uint64 = 2
                                                _ = v100
                                                _ = v101
                                                _ = v102
                                            }
                                            if false {
                                                fmt.Println(fi92, v6)
                                                v2 -= (((v2 - (1.5 - v2)) + (v2 * (float32(1e10) / v2))) / 1e10)
                                            }
                                            switch v28 {
                                            case "":
                                                v28 = (v28 + "hello")
                                            }
                                            _ = fi99
                                        }
                                        fmt.Println(v6, v28, v6)
                                        switch fi92 {
                                        case 2147483647:
                                            if (!false) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            } else {
                                                switch "a" {
                                                case "a":
                                                    v2 += v2
                                                default:
                                                    v28 = v28
                                                }
                                            }
                                        case 32767:
                                            if (true || false) {
                                                v103 := GenS8{F0: v6, F1: (v27 >> (((v6 - v6)) & 63))}
                                                _ = v103
                                            }
                                        case -2147483648:
                                            var v104 int64 = -1
                                            _ = v104
                                        }
                                    }
                                }
                            }
                        }
                    }
                    v27 = (v27 | 251)
                }
            }
            _ = fi92
        }
        _ = v59
    }
    fmt.Println(v2, v2)
    v105 := GenS12{F0: v6}
    fmt.Println(v6, v27, v27)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v105
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
