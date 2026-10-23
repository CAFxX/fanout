package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint32; F2 int }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint16; F1 uint64; F2 bool }
type GenS5 struct { F0 string; F1 uint16; F2 int8 }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 bool; F1 uint16; F2 string }
type GenS8 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS9 struct { F0 int8; F1 int16; F2 int64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 string; F1 int32 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 uint64; F1 uint8 }
type GenS17 struct { F0 string; F1 int }
type GenS18 struct { F0 uint64 }
type GenS19 struct { F0 int16; F1 int }

func main() {
    var v1 int = -2147483648
    var v2 uint = 65535
    v3 := []string{}
    var v4 string = "foo bar"
    v5 := make(map[string]string)
    var v6 float64 = 0.0
    switch (v1 | (v1 | 0)) {
    case 32768:
        for fi7 := 0; fi7 < 8; fi7++ {
            if (0 <= 1) {
                for fi8 := 0; fi8 < 19; fi8++ {
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    _ = fi8
                }
                v9 := GenS1{F0: 42}
                _ = len(v5)
                _ = v9
            } else {
                for fi10 := 0; fi10 < 20; fi10++ {
                    v11 := GenS2{F0: -398, F1: 160, F2: (fi7 + (fi7 / ((fi7) | 1)))}
                    v1 = 42
                    _ = fi10
                    _ = v11
                }
                var v12 uint = v2
                for fi13 := 0; fi13 < 5; fi13++ {
                    if ((-9223372036854775808 == (v1 ^ fi13)) && true) {
                        switch (v4 + ("" + v4)) {
                        case "a":
                            switch "\n\t" {
                            case "":
                                if (18446744073709551615 == 185) {
                                    v14 := GenS3{F0: 1, F1: 4294967295}
                                    v15 := GenS3{F0: 100, F1: 4294967295}
                                    if true {
                                        var v16 float64 = v6
                                        _ = v16
                                    }
                                    _ = v14
                                    _ = v15
                                }
                            case "a":
                                if false {
                                    v5[""] = ((v4 + v4) + ("" + "a"))
                                    if ("a" != (v4 + (v4 + v4))) {
                                        if (3 > 2) {
                                            if true {
                                                var v17 uint32 = 42
                                                _ = v17
                                            } else {
                                                if false {
                                                    v12 = 4294967295
                                                } else {
                                                    if (-2147483648 <= 42) {
                                                        v18 := GenS4{F0: 340, F1: 4294967295, F2: (!true)}
                                                        if false {
                                                            fmt.Println(v1, fi7, v12)
                                                            fmt.Println(v2)
                                                            v4 = ((v4 + (v4 + "hello")) + v4)
                                                        } else {
                                                            v1 -= (v1 & fi7)
                                                            v19 := GenS1{F0: 0}
                                                            if false {
                                                                v20 := GenS5{F0: "foo bar", F1: 248, F2: -1}
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    switch v4 {
                                                                    case "a":
                                                                        if false {
                                                                            v3 = append(v3, "a")
                                                                        } else {
                                                                            v21 := GenS3{F0: 1, F1: 4294967295}
                                                                            _ = v21
                                                                        }
                                                                    case "":
                                                                        v12 += v12
                                                                        switch "a" {
                                                                        case "hello":
                                                                            fmt.Println(fi13, v2)
                                                                        case "\n\t":
                                                                            var v22 uint8 = 245
                                                                            _ = v22
                                                                        case "a":
                                                                            v23 := v5[(v4 + v4)]
                                                                            _ = v23
                                                                        }
                                                                        v4 = v4
                                                                    case "foo bar":
                                                                        var v24 uint = 647
                                                                        switch fi13 {
                                                                        case 128:
                                                                            v25 := GenS5{F0: "", F1: 2, F2: 42}
                                                                            _ = v25
                                                                        case -2147483648:
                                                                            fmt.Println(v2, fi7, v12)
                                                                        case -617:
                                                                            switch "a" {
                                                                            case "":
                                                                                v26 := GenS2{F0: (^(fi7)), F1: 2147483647, F2: 9223372036854775807}
                                                                                _ = v26
                                                                            case "\n\t":
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v27 float64 = 0.0
                                                                            _ = v27
                                                                        }
                                                                        v28 := GenS3{F0: 42, F1: 223}
                                                                        _ = v24
                                                                        _ = v28
                                                                    }
                                                                case "foo bar":
                                                                    var v29 int32 = 460
                                                                    _ = v29
                                                                default:
                                                                    switch (^(fi7)) {
                                                                    case 32768:
                                                                        v30 := GenS1{F0: -32768}
                                                                        _ = v30
                                                                    case 0:
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            var v31 uint16 = 2
                                                                            _ = v31
                                                                        case "a":
                                                                            if (-2147483648 != -1) {
                                                                                fmt.Println(fi13, v2, fi13)
                                                                                switch v4 {
                                                                                case "":
                                                                                    v12 -= 434
                                                                                case "hello":
                                                                                    var v32 int64 = 0
                                                                                    _ = v32
                                                                                case "\n\t":
                                                                                    v33 := GenS6{F0: 4294967295, F1: 1}
                                                                                    _ = v33
                                                                                default:
                                                                                    switch fi7 {
                                                                                    case 911:
                                                                                        var v34 uint64 = 208
                                                                                        _ = v34
                                                                                    default:
                                                                                        _ = len(v5)
                                                                                    }
                                                                                }
                                                                                var v35 bool = (1 < (^(fi7)))
                                                                                _ = v35
                                                                            }
                                                                            v36 := v5[v4]
                                                                            fmt.Println(v36, v1, fi7)
                                                                            _ = v36
                                                                        }
                                                                    case 4294967295:
                                                                        if (42 != 255) {
                                                                            switch v4 {
                                                                            case "a":
                                                                                if (false || false) {
                                                                                    var v37 float32 = 0.1
                                                                                    if (true && (!false)) {
                                                                                        switch 127 {
                                                                                        case 128:
                                                                                            var v38 int = v1
                                                                                            _ = v38
                                                                                        case 32767:
                                                                                            switch "hello" {
                                                                                            case "foo bar":
                                                                                                if (!true) {
                                                                                                    v4 = v4
                                                                                                    fmt.Println(v4, v6, fi13)
                                                                                                } else {
                                                                                                    if true {
                                                                                                        if (!true) {
                                                                                                            var v39 uint64 = 4294967295
                                                                                                            _ = len(v3)
                                                                                                            _ = v39
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                fmt.Println(v1, v6, v6)
                                                                                                                v40 := GenS7{F0: (((187 <= 255) || ((v6 - (v6 / (0.0 / (float64(1.5) - float64((0.1 / v37)))))) <= 0.1)) || (2147483647 <= -736)), F1: 543, F2: ((v4 + "") + v4)}
                                                                                                                _ = v40
                                                                                                            } else {
                                                                                                                fmt.Println(fi13)
                                                                                                            }
                                                                                                            var v41 int16 = 0
                                                                                                            v1 += (v1 >> ((v2) & 63))
                                                                                                            _ = len(v5)
                                                                                                            var v42 float64 = 3.14159
                                                                                                            v43 := GenS8{F0: 65535, F1: 977, F2: 0}
                                                                                                            v2 -= v12
                                                                                                            _ = v41
                                                                                                            _ = v42
                                                                                                            _ = v43
                                                                                                        }
                                                                                                        var v44 int16 = 42
                                                                                                        v2 += (v2 >> ((500) & 63))
                                                                                                        v45 := GenS9{F0: 91, F1: (v44 % ((-1) | 1)), F2: 608}
                                                                                                        var v46 int64 = -1
                                                                                                        _ = v44
                                                                                                        _ = v45
                                                                                                        _ = v46
                                                                                                    } else {
                                                                                                        fmt.Println(fi13)
                                                                                                        if (true || (!(((true || ("\n\t" == v4)) && true) || false))) {
                                                                                                            v1 -= -1
                                                                                                            fmt.Println(v37, v12, v4)
                                                                                                            _ = len(v3)
                                                                                                        }
                                                                                                    }
                                                                                                    v3 = append(v3, ("hello" + v4))
                                                                                                }
                                                                                            case "":
                                                                                                v47 := GenS10{F0: v2}
                                                                                                _ = v47
                                                                                            default:
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = v4
                                                                                                }
                                                                                            }
                                                                                        case -2147483648:
                                                                                            if ((v1 >= (fi13 >> ((v12) & 63))) && (!true)) {
                                                                                                var v48 float32 = ((v37 * v37) / v37)
                                                                                                var v49 int8 = 106
                                                                                                var v50 int64 = -967
                                                                                                _ = v48
                                                                                                _ = v49
                                                                                                _ = v50
                                                                                            }
                                                                                        default:
                                                                                            delete(v5, v4)
                                                                                        }
                                                                                    } else {
                                                                                        v2 -= v12
                                                                                        switch "a" {
                                                                                        case "a":
                                                                                            fmt.Println(v2, v2)
                                                                                        case "\n\t":
                                                                                            var v51 int32 = -1
                                                                                            _ = v51
                                                                                        }
                                                                                        if true {
                                                                                            switch ("hello" + (v4 + "foo bar")) {
                                                                                            case "\n\t":
                                                                                                switch v4 {
                                                                                                case "\n\t":
                                                                                                    fmt.Println(v12, v6, v2)
                                                                                                case "foo bar":
                                                                                                    v12 = 255
                                                                                                case "a":
                                                                                                    fmt.Println(fi7, v12, v2)
                                                                                                }
                                                                                            case "":
                                                                                                v52 := GenS11{F0: 42}
                                                                                                _ = v52
                                                                                            }
                                                                                            switch "" {
                                                                                            case "hello":
                                                                                                v37 -= v37
                                                                                            case "":
                                                                                                fmt.Println(fi13, v12)
                                                                                            default:
                                                                                                v53 := v5[(("\n\t" + "foo bar") + (("" + v4) + v4))]
                                                                                                switch v1 {
                                                                                                case -201:
                                                                                                    if false {
                                                                                                        fmt.Println(fi13, v1, v1)
                                                                                                    } else {
                                                                                                        v54 := v5["a"]
                                                                                                        _ = v54
                                                                                                    }
                                                                                                case 0:
                                                                                                    if (!true) {
                                                                                                        v55 := GenS9{F0: 55, F1: 0, F2: 9223372036854775807}
                                                                                                        v12 = (v12 ^ v12)
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                        _ = v55
                                                                                                    }
                                                                                                }
                                                                                                if ((990 == -18) || (v37 >= v37)) {
                                                                                                    v56 := v5[(((("hello" + ("foo bar" + v53)) + ("\n\t" + (v4 + v4))) + (v53 + ((((((v53 + (("" + "") + "\n\t")) + v4) + v4) + ("\n\t" + v53)) + v53) + (v53 + v53)))) + v53)]
                                                                                                    switch v4 {
                                                                                                    case "\n\t":
                                                                                                        switch fi13 {
                                                                                                        case 4294967295:
                                                                                                            var v57 float64 = (3.14159 + 1e10)
                                                                                                            _ = v57
                                                                                                        case -2147483648:
                                                                                                            var v58 uint = v2
                                                                                                            _ = v58
                                                                                                        case -43:
                                                                                                            switch (^(fi13)) {
                                                                                                            case -951:
                                                                                                                switch v4 {
                                                                                                                case "":
                                                                                                                    if true {
                                                                                                                        fmt.Println(v6, v4, v6)
                                                                                                                        if false {
                                                                                                                            switch (v53 + "\n\t") {
                                                                                                                            case "\n\t":
                                                                                                                                v59 := GenS12{F0: v2}
                                                                                                                                _ = v59
                                                                                                                            case "":
                                                                                                                                v6 += v6
                                                                                                                            case "hello":
                                                                                                                                if true {
                                                                                                                                    switch (v1 & 313) {
                                                                                                                                    case -696:
                                                                                                                                        v60 := GenS13{F0: 42}
                                                                                                                                        _ = v60
                                                                                                                                    case 848:
                                                                                                                                        fmt.Println(v6)
                                                                                                                                    case 32767:
                                                                                                                                        switch v1 {
                                                                                                                                        case 9223372036854775807:
                                                                                                                                            switch (fi13 << ((v2) & 63)) {
                                                                                                                                            case -594:
                                                                                                                                                if len(v3) > 0 {
                                                                                                                                                    v3[0] = "\n\t"
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        case 32767:
                                                                                                                                            fmt.Println(fi7, v12)
                                                                                                                                        case -131:
                                                                                                                                            v61 := GenS10{F0: (v2 & 1)}
                                                                                                                                            _ = v61
                                                                                                                                        default:
                                                                                                                                            var v62 int64 = -1
                                                                                                                                            _ = v62
                                                                                                                                        }
                                                                                                                                        var v63 uint16 = 32767
                                                                                                                                        v12 = 1
                                                                                                                                        _ = v63
                                                                                                                                    default:
                                                                                                                                        var v64 uint32 = 333
                                                                                                                                        _ = v64
                                                                                                                                    }
                                                                                                                                    var v65 string = "\n\t"
                                                                                                                                    _ = v65
                                                                                                                                }
                                                                                                                            }
                                                                                                                            v3 = append(v3, (v53 + (v53 + v53)))
                                                                                                                            switch (v4 + "a") {
                                                                                                                            case "":
                                                                                                                                if len(v3) >= 2 {
                                                                                                                                    v3 = v3[1:len(v3)]
                                                                                                                                }
                                                                                                                            case "a":
                                                                                                                                var v66 int16 = -4
                                                                                                                                _ = v66
                                                                                                                            }
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        var v67 int64 = 42
                                                                                                                        var v68 uint64 = 30
                                                                                                                        _ = v67
                                                                                                                        _ = v68
                                                                                                                    }
                                                                                                                default:
                                                                                                                    _ = len(v3)
                                                                                                                }
                                                                                                            case -1:
                                                                                                                var v69 uint16 = 0
                                                                                                                _ = v69
                                                                                                            default:
                                                                                                                switch (((v4 + v56) + v56) + (("foo bar" + v53) + v56)) {
                                                                                                                case "foo bar":
                                                                                                                    if false {
                                                                                                                        _ = len(v5)
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        var v70 int16 = -292
                                                                                                        _ = v70
                                                                                                    default:
                                                                                                        fmt.Println(v37, fi7, v37)
                                                                                                    }
                                                                                                    v2 = (v12 - v12)
                                                                                                    _ = v56
                                                                                                }
                                                                                                _ = v53
                                                                                            }
                                                                                            fmt.Println(fi13)
                                                                                        } else {
                                                                                            v2 = (v2 << (((v2 ^ (v12 + 255))) & 63))
                                                                                            v71 := v5["hello"]
                                                                                            v72 := GenS13{F0: 42}
                                                                                            _ = v71
                                                                                            _ = v72
                                                                                        }
                                                                                    }
                                                                                    v73 := GenS10{F0: (^(v12))}
                                                                                    _ = v37
                                                                                    _ = v73
                                                                                }
                                                                            case "hello":
                                                                                switch v4 {
                                                                                case "hello":
                                                                                    fmt.Println(v4, fi13, fi7)
                                                                                }
                                                                            default:
                                                                                if false {
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        v12 += (v2 << (((v2 - 2)) & 63))
                                                                                    case "\n\t":
                                                                                        v74 := GenS7{F0: ("hello" >= ""), F1: 255, F2: (v4 + ("hello" + v4))}
                                                                                        _ = v74
                                                                                    case "hello":
                                                                                        v75 := GenS5{F0: "a", F1: 32767, F2: 67}
                                                                                        _ = v75
                                                                                    default:
                                                                                        delete(v5, "a")
                                                                                    }
                                                                                } else {
                                                                                    var v76 uint16 = 815
                                                                                    _ = v76
                                                                                }
                                                                                if false {
                                                                                    v2 = v2
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        v77 := GenS4{F0: 0, F1: 2, F2: (1 < 0)}
                                                                                        _ = v77
                                                                                    case "a":
                                                                                        if true {
                                                                                            fmt.Println(fi13, fi7)
                                                                                        }
                                                                                    case "hello":
                                                                                        var v78 int32 = -2147483648
                                                                                        _ = v78
                                                                                    default:
                                                                                        if ((v6 + 1e10) >= (1.5 - v6)) {
                                                                                            if (105 <= 127) {
                                                                                                var v79 int32 = 1
                                                                                                var v80 uint32 = 1
                                                                                                var v81 float64 = float64(-2.25)
                                                                                                _ = v79
                                                                                                _ = v80
                                                                                                _ = v81
                                                                                            } else {
                                                                                                v82 := GenS9{F0: 85, F1: 1, F2: -601}
                                                                                                var v83 uint16 = 42
                                                                                                var v84 uint = v2
                                                                                                _ = v82
                                                                                                _ = v83
                                                                                                _ = v84
                                                                                            }
                                                                                            v6 = 0.0
                                                                                        }
                                                                                    }
                                                                                    v3 = append(v3, v4)
                                                                                    switch (fi13 >> ((v12) & 63)) {
                                                                                    case 9223372036854775807:
                                                                                        fmt.Println(v4)
                                                                                    default:
                                                                                        v85 := GenS2{F0: v1, F1: 1, F2: v1}
                                                                                        _ = v85
                                                                                    }
                                                                                    var v86 uint64 = 775
                                                                                    _ = v86
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v1 = (v1 | (^(fi13)))
                                                                        }
                                                                    }
                                                                }
                                                                _ = v20
                                                            } else {
                                                                switch fi13 {
                                                                case 32767:
                                                                    fmt.Println(v6, fi7, v1)
                                                                default:
                                                                    var v87 int = v1
                                                                    var v88 uint16 = 32767
                                                                    _ = len(v3)
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                            }
                                                            _ = v19
                                                        }
                                                        _ = v18
                                                    } else {
                                                        var v89 float32 = 0.1
                                                        v1 = fi13
                                                        _ = v89
                                                    }
                                                    if ((!true) || false) {
                                                        v3 = append(v3, ("\n\t" + v4))
                                                    } else {
                                                        v2 = (v12 & v2)
                                                        var v90 string = (((v4 + "hello") + (v4 + (v4 + ("\n\t" + v4)))) + ((("hello" + (v4 + ("hello" + "hello"))) + ((v4 + v4) + v4)) + (((v4 + (v4 + "")) + "hello") + (v4 + (v4 + v4)))))
                                                        _ = v90
                                                    }
                                                }
                                            }
                                            var v91 float64 = 0.1
                                            if true {
                                                v1 = 32768
                                                v4 = "\n\t"
                                                if false {
                                                    switch "a" {
                                                    case "a":
                                                        v5[((("\n\t" + v4) + "\n\t") + ("hello" + v4))] = ((("hello" + ((("foo bar" + (v4 + "hello")) + (v4 + v4)) + "a")) + v4) + "\n\t")
                                                    case "":
                                                        v2 = 321
                                                    case "foo bar":
                                                        var v92 int32 = 2147483647
                                                        _ = v92
                                                    default:
                                                        v4 = (v4 + "hello")
                                                    }
                                                }
                                                if ((!((2147483647 < 211) && (!true))) && ((!true) || false)) {
                                                    fmt.Println(fi13)
                                                    v12 = v12
                                                    v1 = (-(fi7))
                                                } else {
                                                    v5[(v4 + "foo bar")] = v4
                                                }
                                                v93 := GenS10{F0: v12}
                                                _ = v93
                                            } else {
                                                var v94 uint8 = 128
                                                _ = v94
                                            }
                                            _ = v91
                                        } else {
                                            var v95 uint = 2
                                            _ = v95
                                        }
                                        if ((996 != -32768) || (32767 >= 439)) {
                                            var v96 float32 = -2.25
                                            v6 = (v6 - 0.1)
                                            _ = v96
                                        }
                                        v97 := GenS14{F0: "", F1: 2147483647}
                                        v98 := GenS7{F0: (32767 == 32767), F1: 110, F2: (v4 + "hello")}
                                        var v99 int8 = -72
                                        var v100 int8 = v99
                                        _ = v100
                                        _ = v97
                                        _ = v98
                                        _ = v99
                                    } else {
                                        v4 = ("a" + ((("\n\t" + v4) + v4) + ("hello" + "hello")))
                                        v4 = v4
                                    }
                                }
                            default:
                                if (65535 < 42) {
                                    fmt.Println(v6)
                                    fmt.Println(fi13, v2, v1)
                                }
                            }
                        }
                        _ = len(v3)
                        v6 = (1e10 + (v6 + v6))
                        var v101 float32 = -2.25
                        _ = v101
                    } else {
                        var v102 float64 = v6
                        v103 := GenS15{F0: 42}
                        v2 = v12
                        _ = v102
                        _ = v103
                    }
                    switch v4 {
                    case "hello":
                        v2 = v12
                    }
                    _ = fi13
                }
                _ = v12
            }
            if false {
                if len(v3) > 0 {
                    v3[0] = v4
                }
            } else {
                v1 = fi7
                var v104 int64 = 42
                v105 := GenS3{F0: 127, F1: 65535}
                _ = v104
                _ = v105
            }
            _ = fi7
        }
    }
    switch (v1 & (v1 | v1)) {
    case 4294967295:
        v6 = (v6 * float64(3.14159))
    case 128:
        v106 := GenS15{F0: 229}
        _ = v106
    }
    v4 = v4
    for fi107 := 0; fi107 < 17; fi107++ {
        var v108 float32 = 3.14159
        if ((v1 + v1) != v1) {
            fmt.Println(v108)
            for fi109 := 0; fi109 < 5; fi109++ {
                if len(v3) > 0 {
                    v3[0] = (v4 + v4)
                }
                _ = fi109
            }
        } else {
            for fi110 := 0; fi110 < 7; fi110++ {
                if (253 == 255) {
                    if false {
                        v111 := GenS16{F0: 162, F1: 252}
                        if (true && true) {
                            if (true || (255 > 122)) {
                                v1 -= v1
                                switch 0 {
                                case 2147483647:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                default:
                                    fmt.Println(v108)
                                    _ = len(v5)
                                }
                                if true {
                                    var v112 string = ""
                                    _ = v112
                                } else {
                                    switch -2147483648 {
                                    case -181:
                                        var v113 float64 = -2.25
                                        v3 = append(v3, (v4 + (v4 + v4)))
                                        _ = v113
                                    case 4294967295:
                                        _ = len(v5)
                                    case 42:
                                        switch fi107 {
                                        case 9223372036854775807:
                                            v108 = v108
                                        default:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    default:
                                        fmt.Println(v1)
                                    }
                                    var v114 uint8 = 108
                                    _ = v114
                                }
                                v115 := GenS16{F0: 4294967295, F1: 176}
                                _ = v115
                            } else {
                                if ((v4 > (v4 + v4)) || ((2 != 4294967295) || (false || false))) {
                                    var v116 int = (fi110 & (fi110 | 4294967295))
                                    _ = v116
                                } else {
                                    v1 = fi107
                                }
                                v1 -= (fi110 - v1)
                                switch fi110 {
                                case 642:
                                    v6 = v6
                                case 0:
                                    if (4294967295 >= 4294967295) {
                                        switch v1 {
                                        case -39:
                                            fmt.Println(fi110)
                                        case 32768:
                                            var v117 bool = (0 >= -78)
                                            _ = v117
                                        default:
                                            v118 := GenS9{F0: 100, F1: int16(fi110), F2: -644}
                                            _ = v118
                                        }
                                        v4 = "a"
                                        var v119 uint = 273
                                        _ = v119
                                    }
                                    var v120 uint = 40
                                    var v121 uint = v120
                                    var v122 string = v4
                                    if true {
                                        if len(v3) > 0 {
                                            v3[0] = "a"
                                        }
                                        var v123 uint64 = 4294967295
                                        _ = v123
                                    } else {
                                        v124 := GenS14{F0: "foo bar", F1: 228}
                                        v1 += 128
                                        switch fi107 {
                                        case 2147483647:
                                            var v125 bool = (42 >= 1)
                                            _ = v125
                                        case 4294967295:
                                            switch (fi107 >> (((v120 | v120)) & 63)) {
                                            case 1:
                                                fmt.Println(v120)
                                            case -2147483648:
                                                v126 := GenS3{F0: 1, F1: 2}
                                                _ = v126
                                            case -9223372036854775808:
                                                v120 += v120
                                            default:
                                                v108 = v108
                                            }
                                        }
                                        _ = v124
                                    }
                                    _ = v120
                                    _ = v121
                                    _ = v122
                                case 9223372036854775807:
                                    var v127 int64 = 1
                                    _ = v127
                                default:
                                    var v128 string = ("foo bar" + v4)
                                    _ = v128
                                }
                                v129 := GenS14{F0: v4, F1: -23}
                                var v130 uint64 = 1
                                v131 := GenS17{F0: (v4 + "foo bar"), F1: v1}
                                fmt.Println(v1, v108)
                                _ = v129
                                _ = v130
                                _ = v131
                            }
                        } else {
                            v108 += float32(v6)
                            v6 += (v6 + v6)
                        }
                        delete(v5, "\n\t")
                        _ = v111
                    }
                } else {
                    v108 = float32(v6)
                    v3 = append(v3, (v4 + (("foo bar" + "hello") + ((v4 + "\n\t") + v4))))
                    var v132 bool = ((!false) || false)
                    var v133 int16 = 32767
                    _ = v132
                    _ = v133
                }
                v3 = append(v3, ("\n\t" + (v4 + v4)))
                _ = fi110
            }
            _ = len(v3)
        }
        if (602 < 18446744073709551615) {
            v1 = -27
            if (65535 > 65535) {
                v2 = v2
                var v134 string = "hello"
                v4 = v4
                _ = v134
            } else {
                for fi135 := 0; fi135 < 2; fi135++ {
                    v136 := GenS18{F0: 558}
                    var v137 float64 = v6
                    if ((false || true) && (0 != v2)) {
                        v4 = "\n\t"
                        if (false || (1 != 4294967295)) {
                            var v138 int16 = 848
                            _ = v138
                        }
                        if (int8(v2) == -1) {
                            var v139 float32 = v108
                            _ = v139
                        } else {
                            v108 -= v108
                            _ = len(v5)
                        }
                        _ = len(v3)
                        var v140 float64 = v137
                        _ = v140
                    }
                    _ = fi135
                    _ = v136
                    _ = v137
                }
                v5["hello"] = v4
            }
            v5[v4] = v4
        }
        _ = fi107
        _ = v108
    }
    if false {
        v6 -= ((float64((float32((v6 + v6)) + ((3.14159 - 1e10) / (3.14159 / -2.25)))) * v6) * 0.1)
        for fi141 := 0; fi141 < 5; fi141++ {
            var v142 string = (("" + v4) + (v4 + (("" + ("\n\t" + v4)) + v4)))
            if true {
                var v143 uint64 = 531
                _ = v143
            }
            _ = fi141
            _ = v142
        }
    }
    for fi144 := 0; fi144 < 6; fi144++ {
        var v145 int8 = 51
        var v146 float64 = (v6 - -2.25)
        var v147 bool = false
        _ = fi144
        _ = v145
        _ = v146
        _ = v147
    }
    fmt.Println(v1, v4)
    for fi148 := 0; fi148 < 6; fi148++ {
        var v149 uint32 = 4294967295
        if (182 == 127) {
            for fi150 := 0; fi150 < 14; fi150++ {
                fmt.Println(fi148)
                var v151 float32 = 3.14159
                v4 = ((((("\n\t" + (v4 + v4)) + ("hello" + v4)) + (v4 + "")) + "a") + ("hello" + v4))
                _ = fi150
                _ = v151
            }
        } else {
            var v152 int32 = 0
            _ = v152
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi148
        _ = v149
    }
    v6 = v6
    v2 = 4294967295
    for fi153 := 0; fi153 < 18; fi153++ {
        for fi154 := 0; fi154 < 16; fi154++ {
            if true {
                delete(v5, (v4 + ("a" + v4)))
            }
            _ = fi154
        }
        v4 = v4
        v6 += v6
        _ = fi153
    }
    if true {
        delete(v5, "a")
    } else {
        v6 = v6
        var v155 uint64 = 278
        _ = v155
    }
    if (0 <= 127) {
        v2 += v2
    } else {
        var v156 uint32 = 538
        v6 = float64(3.14159)
        _ = v156
    }
    switch (-(v1)) {
    case 42:
        v2 = (^(v2))
    case 693:
        for fi157 := 0; fi157 < 14; fi157++ {
            var v158 int64 = -9223372036854775808
            _ = fi157
            _ = v158
        }
    }
    v159 := GenS1{F0: -613}
    v6 = float64(((((0.0 / -2.25) + ((-2.25 - 3.14159) + float32(v6))) / 3.14159) * (float32(1e10) - (((float32(3.14159) - (1e10 / (0.0 - ((1.5 - -2.25) / (1.5 * (1.5 - -2.25)))))) + 0.1) + (3.14159 * 0.1)))))
    v2 = 799
    switch v4 {
    case "a":
        for fi160 := 0; fi160 < 4; fi160++ {
            for fi161 := 0; fi161 < 10; fi161++ {
                _ = len(v5)
                var v162 int64 = -59
                _ = fi161
                _ = v162
            }
            _ = fi160
        }
    case "\n\t":
        for fi163 := 0; fi163 < 3; fi163++ {
            if true {
                v2 -= v2
            } else {
                v164 := GenS19{F0: 1, F1: v1}
                _ = v164
            }
            _ = fi163
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v159
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
