package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint64; F1 int; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int16; F1 bool }
type GenS9 struct { F0 int32; F1 uint }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint32; F1 bool }
type GenS12 struct { F0 int; F1 string; F2 int }
type GenS13 struct { F0 int16; F1 int16 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 uint8; F1 string; F2 string }

func main() {
    var v1 uint64 = 42
    v2 := []string{"\n\t", "\n\t", "foo bar"}
    v3 := []int{127}
    switch (((("foo bar" + "hello") + "hello") + "") + ("" + "a")) {
    case "hello":
        var v4 int8 = -125
        _ = v4
    case "foo bar":
        var v5 int = 460
        _ = v5
    default:
        if false {
            var v6 float32 = 3.14159
            _ = v6
        } else {
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                var v7 uint64 = 81
                _ = v7
            case "foo bar":
                switch 32768 {
                case 9223372036854775807:
                    v1 -= v1
                case 2147483647:
                    v8 := GenS1{F0: (v1 - v1), F1: ""}
                    _ = v8
                default:
                    v9 := GenS2{F0: v1, F1: 67, F2: v1}
                    _ = v9
                }
            case "hello":
                for fi10 := 0; fi10 < 11; fi10++ {
                    if len(v3) > 0 {
                        v3[0] = 128
                    }
                    _ = fi10
                }
            default:
                v1 = v1
            }
            switch 127 {
            case 1:
                if (((1e10 + 0.0) - (3.14159 / -2.25)) != (1.5 + 1e10)) {
                    if true {
                        switch 0 {
                        case 9223372036854775807:
                            var v11 int8 = -55
                            _ = v11
                        default:
                            _ = len(v3)
                        }
                        if true {
                            var v12 uint8 = 93
                            switch (("a" + "hello") + "foo bar") {
                            case "foo bar":
                                for fi13 := 0; fi13 < 20; fi13++ {
                                    var v14 uint32 = 207
                                    _ = fi13
                                    _ = v14
                                }
                            case "hello":
                                var v15 uint32 = 247
                                _ = v15
                            case "a":
                                if len(v3) > 0 {
                                    v3[0] = 740
                                }
                            default:
                                if (v12 != (v12 - 62)) {
                                    switch "foo bar" {
                                    case "":
                                        if true {
                                            if ((-267 <= -861) || (true || (-746 > 42))) {
                                                v1 += uint64(v12)
                                                fmt.Println(v12)
                                            } else {
                                                v16 := GenS3{F0: 0}
                                                v17 := GenS1{F0: (v1 ^ 42), F1: ("foo bar" + "")}
                                                _ = v16
                                                _ = v17
                                            }
                                        } else {
                                            switch (("" + "\n\t") + "\n\t") {
                                            case "foo bar":
                                                var v18 uint32 = 2
                                                _ = v18
                                            default:
                                                if len(v3) > 0 {
                                                    v3[0] = 4294967295
                                                }
                                            }
                                        }
                                    }
                                    v1 += v1
                                } else {
                                    fmt.Println(v1, v1)
                                    var v19 int32 = -1
                                    _ = v19
                                }
                            }
                            _ = v12
                        }
                        var v20 int8 = -1
                        _ = v20
                    } else {
                        fmt.Println(v1)
                        var v21 int64 = 9223372036854775807
                        v1 += (v1 * v1)
                        _ = v21
                    }
                    var v22 bool = true
                    _ = v22
                }
            case -9223372036854775808:
                v23 := GenS1{F0: v1, F1: "hello"}
                _ = v23
            }
        }
    }
    if true {
        if len(v2) > 0 {
            v2[0] = "a"
        }
    } else {
        var v24 int64 = 42
        v25 := GenS2{F0: 18446744073709551615, F1: -2147483648, F2: 18446744073709551615}
        if true {
            var v26 uint64 = 1
            _ = v26
        } else {
            if false {
                if (587 == 255) {
                    switch 128 {
                    case 32768:
                        v1 -= (v1 | 42)
                    case 611:
                        switch 32767 {
                        case -689:
                            v2 = append(v2, "")
                        case -1:
                            var v27 uint32 = 539
                            _ = v27
                        case -419:
                            v28 := GenS1{F0: 2, F1: ""}
                            _ = v28
                        default:
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                        }
                    }
                } else {
                    var v29 uint16 = 32767
                    _ = v29
                }
                for fi30 := 0; fi30 < 5; fi30++ {
                    v1 = v1
                    var v31 uint64 = 0
                    v1 -= v31
                    _ = fi30
                    _ = v31
                }
            } else {
                _ = len(v3)
            }
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
        }
        _ = v24
        _ = v25
    }
    for fi32 := 0; fi32 < 15; fi32++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = len(v3)
        _ = fi32
    }
    var v33 int64 = -1
    switch ("hello" + "hello") {
    case "\n\t":
        v1 = (v1 >> ((0) & 63))
    case "hello":
        if (42 < 2) {
            fmt.Println(v33)
        } else {
            var v34 string = (("hello" + "\n\t") + "\n\t")
            v35 := GenS4{F0: 270}
            _ = v34
            _ = v35
        }
    case "foo bar":
        fmt.Println(v33)
    default:
        if false {
            if false {
                var v36 int8 = 50
                var v37 float32 = 1.5
                switch ("hello" + "\n\t") {
                case "hello":
                    var v38 float64 = (0.0 + 1.5)
                    _ = v38
                case "":
                    var v39 bool = false
                    _ = v39
                case "\n\t":
                    var v40 int32 = 42
                    _ = v40
                default:
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v36)
                    case "a":
                        var v41 float32 = v37
                        _ = v41
                    default:
                        v33 -= v33
                    }
                }
                _ = v36
                _ = v37
            }
        }
    }
    var v42 uint64 = v1
    v43 := GenS1{F0: v42, F1: "hello"}
    fmt.Println(v1, v42, v42)
    for fi44 := 0; fi44 < 14; fi44++ {
        fmt.Println(v42)
        _ = fi44
    }
    for fi45 := 0; fi45 < 8; fi45++ {
        _ = len(v2)
        _ = fi45
    }
    for fi46 := 0; fi46 < 9; fi46++ {
        v47 := GenS2{F0: (v42 | 42), F1: fi46, F2: (v42 / ((v1) | 1))}
        v42 = 1
        for fi48 := 0; fi48 < 5; fi48++ {
            v33 += 1
            switch "foo bar" {
            case "a":
                switch "\n\t" {
                case "a":
                    v49 := GenS3{F0: 673}
                    _ = v49
                case "":
                    if (244 != 1) {
                        v3 = append(v3, (fi48 * fi46))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                case "\n\t":
                    var v50 bool = false
                    _ = v50
                }
            case "\n\t":
                var v51 uint16 = 65535
                _ = v51
            case "":
                if (!(46 >= 42)) {
                    if (!(uint16(v42) != 574)) {
                        v52 := GenS3{F0: 771}
                        if (((v33 & (^(v33))) <= 0) || ((v1 | (v42 & v1)) != v42)) {
                            fmt.Println(fi46, fi46)
                            if (false || true) {
                                if (("hello" + ((((("foo bar" + "") + "hello") + "hello") + "a") + "a")) > "a") {
                                    switch 868 {
                                    case 1:
                                        v53 := GenS1{F0: v42, F1: "a"}
                                        _ = v53
                                    case 157:
                                        fmt.Println(v42, fi48, fi46)
                                    default:
                                        _ = len(v2)
                                    }
                                } else {
                                    v54 := GenS3{F0: 255}
                                    if len(v2) > 0 {
                                        v2[0] = "\n\t"
                                    }
                                    _ = v54
                                }
                                v55 := GenS5{F0: v1}
                                if true {
                                    fmt.Println(fi46)
                                    var v56 int16 = 1
                                    v33 -= (v33 << ((571) & 63))
                                    _ = v56
                                }
                                _ = v55
                            }
                            v42 -= v1
                            fmt.Println(v33, fi46)
                            if (0.1 >= 1.5) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if true {
                                    switch "" {
                                    case "a":
                                        if true {
                                            fmt.Println(v33, fi46, fi48)
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    case "foo bar":
                                        if (-85 != 42) {
                                            v57 := GenS6{F0: (v42 >> ((65535) & 63))}
                                            _ = v57
                                        }
                                    default:
                                        var v58 string = "foo bar"
                                        _ = v58
                                    }
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v59 := GenS7{F0: true}
                                    _ = v59
                                }
                                if true {
                                    if true {
                                        fmt.Println(fi46, fi46)
                                        switch int(v1) {
                                        case -1:
                                            v33 += (^(v33))
                                        case 992:
                                            fmt.Println(v42)
                                            switch "\n\t" {
                                            case "":
                                                v33 += (v33 >> ((363) & 63))
                                            case "\n\t":
                                                if (v33 == v33) {
                                                    v33 -= -9223372036854775808
                                                    var v60 uint64 = v42
                                                    v33 = v33
                                                    _ = v60
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v61 string = ""
                                                    _ = v61
                                                }
                                            default:
                                                _ = len(v3)
                                            }
                                            v1 = 1
                                        case 128:
                                            if (127 < fi46) {
                                                v1 = (v1 / ((v42) | 1))
                                            } else {
                                                v3 = append(v3, 42)
                                                switch "foo bar" {
                                                case "hello":
                                                    if false {
                                                        var v62 bool = (true || false)
                                                        _ = v62
                                                    } else {
                                                        if false {
                                                            var v63 int = fi46
                                                            if true {
                                                                v1 = (v42 ^ v1)
                                                                var v64 uint = 559
                                                                var v65 uint16 = 1
                                                                if false {
                                                                    v66 := GenS8{F0: int16(v1), F1: ((v33 + (^(v33))) < v33)}
                                                                    v64 = v64
                                                                    _ = v66
                                                                }
                                                                _ = v64
                                                                _ = v65
                                                            } else {
                                                                var v67 uint16 = 2
                                                                v3 = append(v3, fi48)
                                                                _ = v67
                                                            }
                                                            var v68 uint32 = 989
                                                            _ = v63
                                                            _ = v68
                                                        }
                                                        var v69 int64 = 1
                                                        v42 -= 650
                                                        _ = v69
                                                    }
                                                case "":
                                                    fmt.Println(v1, fi48, v33)
                                                case "\n\t":
                                                    v42 -= 793
                                                default:
                                                    if (0.0 <= 0.0) {
                                                        switch (fi48 % (((^(fi46))) | 1)) {
                                                        case -9223372036854775808:
                                                            switch 128 {
                                                            case 205:
                                                                v70 := GenS8{F0: -654, F1: true}
                                                                _ = v70
                                                            default:
                                                                v71 := GenS9{F0: -2147483648, F1: uint(v1)}
                                                                _ = v71
                                                            }
                                                        case -219:
                                                            switch fi48 {
                                                            case -9223372036854775808:
                                                                v42 = 349
                                                            case 32767:
                                                                var v72 int = (fi48 << ((644) & 63))
                                                                _ = v72
                                                            case -1:
                                                                _ = len(v3)
                                                            }
                                                        case -869:
                                                            if ((128 < 42) || (true && false)) {
                                                                v3 = append(v3, fi48)
                                                            } else {
                                                                switch fi48 {
                                                                case 42:
                                                                    v73 := GenS10{F0: 2}
                                                                    _ = v73
                                                                default:
                                                                    if false {
                                                                        if (42 == 0) {
                                                                            var v74 uint32 = 189
                                                                            _ = v74
                                                                        } else {
                                                                            var v75 float32 = float32(float64(1.5))
                                                                            _ = v75
                                                                        }
                                                                    }
                                                                }
                                                                fmt.Println(v42)
                                                                switch fi46 {
                                                                case -1:
                                                                    var v76 float32 = 1e10
                                                                    _ = v76
                                                                case 2147483647:
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 0
                                                                    }
                                                                case 32768:
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    case "hello":
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    case "":
                                                                        switch fi48 {
                                                                        case 937:
                                                                            v33 = (v33 / ((v33) | 1))
                                                                            var v77 int32 = 1
                                                                            _ = v77
                                                                        case 32767:
                                                                            if (!false) {
                                                                                v42 -= 0
                                                                                switch int(v33) {
                                                                                case 147:
                                                                                    v33 -= (-(v33))
                                                                                case -404:
                                                                                    var v78 uint16 = 42
                                                                                    _ = v78
                                                                                case -9223372036854775808:
                                                                                    v42 -= 573
                                                                                }
                                                                                v79 := GenS6{F0: (v42 | v42)}
                                                                                if (!(("hello" + ("a" + ("\n\t" + ("" + "")))) != "a")) {
                                                                                    switch 32768 {
                                                                                    case 0:
                                                                                        fmt.Println(fi48, v42, v42)
                                                                                    case 873:
                                                                                        v42 -= (v1 << ((154) & 63))
                                                                                    default:
                                                                                        if true {
                                                                                            if (834 < 65535) {
                                                                                                if true {
                                                                                                    v1 -= v1
                                                                                                }
                                                                                            } else {
                                                                                                switch ("a" + (("\n\t" + (("hello" + "a") + "foo bar")) + "hello")) {
                                                                                                case "hello":
                                                                                                    if (false && false) {
                                                                                                        var v80 bool = (v1 < v1)
                                                                                                        v81 := GenS11{F0: 2147483647, F1: (v80 && v80)}
                                                                                                        switch 9223372036854775807 {
                                                                                                        case -68:
                                                                                                            fmt.Println(v1, v42, fi48)
                                                                                                        default:
                                                                                                            var v82 int8 = 100
                                                                                                            _ = v82
                                                                                                        }
                                                                                                        _ = v80
                                                                                                        _ = v81
                                                                                                    } else {
                                                                                                        _ = len(v2)
                                                                                                        if false {
                                                                                                            fmt.Println(fi46, v1, fi46)
                                                                                                            var v83 bool = ((((-2.25 - 1e10) + 3.14159) - (3.14159 / (1e10 - 1.5))) < 3.14159)
                                                                                                            _ = v83
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                fmt.Println(v42, fi48)
                                                                                                                v1 = 667
                                                                                                            }
                                                                                                            fmt.Println(fi48)
                                                                                                            if true {
                                                                                                                fmt.Println(fi46)
                                                                                                                v84 := GenS6{F0: v42}
                                                                                                                _ = v84
                                                                                                            }
                                                                                                            v1 = (v42 | (v42 % (((v1 << ((710) & 63))) | 1)))
                                                                                                            v85 := GenS11{F0: 304, F1: false}
                                                                                                            _ = v85
                                                                                                        }
                                                                                                        v86 := GenS12{F0: fi48, F1: (("\n\t" + "hello") + ""), F2: (fi48 * fi46)}
                                                                                                        _ = v86
                                                                                                    }
                                                                                                case "a":
                                                                                                    v87 := GenS1{F0: v1, F1: ("hello" + "foo bar")}
                                                                                                    _ = v87
                                                                                                }
                                                                                                if true {
                                                                                                    fmt.Println(fi46, fi46, fi48)
                                                                                                    var v88 int64 = v33
                                                                                                    _ = v88
                                                                                                } else {
                                                                                                    v1 = (v42 | (v1 ^ 42))
                                                                                                }
                                                                                                v1 += v1
                                                                                                if false {
                                                                                                    var v89 uint = 1
                                                                                                    v90 := GenS10{F0: v89}
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = fi48
                                                                                                    }
                                                                                                    var v91 uint = v89
                                                                                                    if true {
                                                                                                        if (true && ((v91 ^ v89) >= (v89 - v89))) {
                                                                                                            switch 128 {
                                                                                                            case 748:
                                                                                                                fmt.Println(v89)
                                                                                                            case -439:
                                                                                                                if false {
                                                                                                                    if (1 < 674) {
                                                                                                                        var v92 int = fi48
                                                                                                                        v93 := GenS11{F0: 42, F1: true}
                                                                                                                        switch -9223372036854775808 {
                                                                                                                        case -2147483648:
                                                                                                                            v94 := GenS3{F0: (v91 / ((uint(v42)) | 1))}
                                                                                                                            _ = v94
                                                                                                                        case 4294967295:
                                                                                                                            if true {
                                                                                                                                switch 771 {
                                                                                                                                case 9223372036854775807:
                                                                                                                                    switch fi46 {
                                                                                                                                    case 42:
                                                                                                                                        var v95 float32 = (1.5 / 0.1)
                                                                                                                                        _ = v95
                                                                                                                                    }
                                                                                                                                case 0:
                                                                                                                                    v96 := GenS13{F0: 0, F1: 1}
                                                                                                                                    _ = v96
                                                                                                                                case -349:
                                                                                                                                    v91 = v89
                                                                                                                                    if false {
                                                                                                                                        v89 += 975
                                                                                                                                    } else {
                                                                                                                                        var v97 uint32 = 0
                                                                                                                                        _ = v97
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    var v98 float32 = (((3.14159 / 1.0) / 1.0) / ((((1e10 / 1.5) * 0.1) + 1.5) / (1.5 * 0.1)))
                                                                                                                                    _ = v98
                                                                                                                                }
                                                                                                                                var v99 uint32 = 888
                                                                                                                                switch v92 {
                                                                                                                                case -2147483648:
                                                                                                                                    if true {
                                                                                                                                        v91 -= v91
                                                                                                                                        var v100 int64 = (^(v33))
                                                                                                                                        if len(v3) > 0 {
                                                                                                                                            v3[0] = fi46
                                                                                                                                        }
                                                                                                                                        _ = v100
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    var v101 int32 = 2147483647
                                                                                                                                    v102 := GenS11{F0: v99, F1: (!true)}
                                                                                                                                    _ = v101
                                                                                                                                    _ = v102
                                                                                                                                }
                                                                                                                                _ = v99
                                                                                                                            }
                                                                                                                        case -1:
                                                                                                                            v89 = (v91 ^ (^(v89)))
                                                                                                                        default:
                                                                                                                            var v103 bool = (!false)
                                                                                                                            _ = v103
                                                                                                                        }
                                                                                                                        _ = v92
                                                                                                                        _ = v93
                                                                                                                    }
                                                                                                                    var v104 float64 = 1e10
                                                                                                                    _ = v104
                                                                                                                }
                                                                                                                v1 += 18446744073709551615
                                                                                                            case 32767:
                                                                                                                v42 += v42
                                                                                                            default:
                                                                                                                _ = len(v3)
                                                                                                            }
                                                                                                            var v105 int8 = 42
                                                                                                            _ = v105
                                                                                                        }
                                                                                                        v106 := GenS2{F0: (v42 << (((^(v91))) & 63)), F1: fi48, F2: 4294967295}
                                                                                                        v89 -= (v89 | 2)
                                                                                                        v42 += (v42 << ((4294967295) & 63))
                                                                                                        v107 := GenS1{F0: (v42 / (((v42 & (v1 | v42))) | 1)), F1: "foo bar"}
                                                                                                        v108 := GenS14{F0: 2147483647}
                                                                                                        _ = v106
                                                                                                        _ = v107
                                                                                                        _ = v108
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                            var v109 float64 = 1.5
                                                                                                            v110 := GenS7{F0: ((-1 <= 546) || true)}
                                                                                                            _ = v109
                                                                                                            _ = v110
                                                                                                        } else {
                                                                                                            v42 -= (v42 / ((v1) | 1))
                                                                                                            v3 = append(v3, (fi48 | -9223372036854775808))
                                                                                                        }
                                                                                                    }
                                                                                                    if true {
                                                                                                        v111 := GenS4{F0: (^(v91))}
                                                                                                        v33 -= (v33 << ((698) & 63))
                                                                                                        switch "a" {
                                                                                                        case "a":
                                                                                                            switch fi48 {
                                                                                                            case 2147483647:
                                                                                                                v89 += v91
                                                                                                            }
                                                                                                            switch "foo bar" {
                                                                                                            case "foo bar":
                                                                                                                var v112 int = (fi48 << ((v91) & 63))
                                                                                                                _ = v112
                                                                                                            }
                                                                                                            v113 := GenS4{F0: 503}
                                                                                                            _ = v113
                                                                                                        default:
                                                                                                            fmt.Println(v33, v91)
                                                                                                        }
                                                                                                        v114 := GenS12{F0: fi46, F1: "foo bar", F2: (fi48 % ((fi46) | 1))}
                                                                                                        _ = v111
                                                                                                        _ = v114
                                                                                                    } else {
                                                                                                        if false {
                                                                                                            fmt.Println(v42, fi48, fi46)
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                            v115 := GenS15{F0: v89}
                                                                                                            v116 := GenS6{F0: (v42 & v1)}
                                                                                                            var v117 uint16 = 32767
                                                                                                            if false {
                                                                                                                var v118 uint64 = (v1 * uint64(v117))
                                                                                                                v1 = (v1 % (((v42 << (((v91 << ((42) & 63))) & 63))) | 1))
                                                                                                                v42 += 87
                                                                                                                v91 = (v89 | v91)
                                                                                                                _ = v118
                                                                                                            }
                                                                                                            _ = v115
                                                                                                            _ = v116
                                                                                                            _ = v117
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v89
                                                                                                    _ = v90
                                                                                                    _ = v91
                                                                                                } else {
                                                                                                    v42 -= v1
                                                                                                    var v119 uint = 255
                                                                                                    var v120 int64 = v33
                                                                                                    _ = v119
                                                                                                    _ = v120
                                                                                                }
                                                                                                var v121 float32 = ((0.0 - 1e10) + -2.25)
                                                                                                var v122 float64 = (float64(v121) * (float64(0.0) + 0.1))
                                                                                                _ = v121
                                                                                                _ = v122
                                                                                            }
                                                                                        }
                                                                                        v42 = (v42 / ((18446744073709551615) | 1))
                                                                                    }
                                                                                    var v123 int32 = -2147483648
                                                                                    _ = v123
                                                                                }
                                                                                var v124 int8 = 0
                                                                                _ = v124
                                                                                _ = v79
                                                                            }
                                                                        case -2147483648:
                                                                            if false {
                                                                                v1 += v1
                                                                            } else {
                                                                                var v125 string = "hello"
                                                                                var v126 int32 = 1
                                                                                _ = v125
                                                                                _ = v126
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    v42 += (v1 ^ (^(v1)))
                                                                }
                                                            }
                                                        }
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                var v127 int = fi48
                                                _ = v127
                                            }
                                        default:
                                            switch ("a" + "hello") {
                                            case "hello":
                                                v1 = 134
                                            case "":
                                                v1 = v42
                                            case "foo bar":
                                                v1 = v1
                                            }
                                        }
                                        if false {
                                            var v128 uint8 = 127
                                            var v129 float64 = 0.1
                                            if true {
                                                var v130 uint = 42
                                                if len(v2) > 0 {
                                                    v2[0] = ("" + "\n\t")
                                                }
                                                _ = v130
                                            }
                                            _ = v128
                                            _ = v129
                                        }
                                    }
                                } else {
                                    switch fi46 {
                                    case 128:
                                        v33 -= 774
                                    default:
                                        var v131 float32 = 1.5
                                        _ = v131
                                    }
                                    var v132 uint8 = 42
                                    var v133 int16 = -167
                                    _ = v132
                                    _ = v133
                                }
                            }
                            var v134 int8 = 100
                            var v135 uint = 1
                            v136 := GenS3{F0: v135}
                            _ = v134
                            _ = v135
                            _ = v136
                        } else {
                            v42 = v42
                            v1 -= v1
                            v42 -= 670
                        }
                        _ = v52
                    } else {
                        v137 := GenS7{F0: (v1 > (^(v1)))}
                        v138 := GenS13{F0: 32767, F1: 1000}
                        _ = v137
                        _ = v138
                    }
                }
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi48
        }
        _ = fi46
        _ = v47
    }
    var v139 int8 = 1
    if true {
        fmt.Println(v33)
        fmt.Println(v33)
        var v140 int = 349
        _ = v140
    }
    v141 := GenS16{F0: 127, F1: (("hello" + ("hello" + "foo bar")) + "a"), F2: ("a" + "a")}
    var v142 uint = 18446744073709551615
    v139 = v139
    var v143 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v139:", v139)
    fmt.Println("v142:", v142)
    fmt.Println("v143:", v143)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v139
    _ = v141
    _ = v142
    _ = v143
    _ = v2
    _ = v3
    _ = v33
    _ = v42
    _ = v43
}
