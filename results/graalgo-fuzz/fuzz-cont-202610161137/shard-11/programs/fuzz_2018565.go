package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16; F1 string }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 bool; F1 int8; F2 int32 }
type GenS12 struct { F0 uint8; F1 string }
type GenS13 struct { F0 bool; F1 int8 }
type GenS14 struct { F0 int16; F1 bool }
type GenS15 struct { F0 uint32; F1 string }

func main() {
    var v1 int32 = 1
    var v2 uint64 = 0
    v3 := []string{(("" + (("hello" + ("foo bar" + "hello")) + ((("hello" + "hello") + "foo bar") + (((("a" + "foo bar") + (("foo bar" + "") + "a")) + ((("a" + "\n\t") + "foo bar") + "\n\t")) + "")))) + "foo bar"), "a", "\n\t", (("foo bar" + "") + "a")}
    v4 := []int{}
    var v5 int8 = -98
    var v6 uint8 = 0
    fmt.Println(v2, v5)
    for fi7 := 0; fi7 < 14; fi7++ {
        if true {
            v3 = append(v3, "")
            v5 = v5
        }
        switch fi7 {
        case -763:
            if (1 <= -9223372036854775808) {
                v8 := GenS1{F0: 32767, F1: (v2 % ((18446744073709551615) | 1))}
                if ((true || (false || true)) || false) {
                    switch "foo bar" {
                    case "foo bar":
                        var v9 uint32 = 4294967295
                        _ = v9
                    case "":
                        v3 = append(v3, "\n\t")
                    default:
                        for fi10 := 0; fi10 < 2; fi10++ {
                            v4 = append(v4, (fi10 / (((-(fi7))) | 1)))
                            switch "foo bar" {
                            case "foo bar":
                                var v11 float32 = 1.5
                                _ = v11
                            case "a":
                                var v12 int8 = 42
                                _ = v12
                            default:
                                v5 -= 100
                            }
                            if (9223372036854775807 >= 201) {
                                var v13 uint8 = (^(v6))
                                if false {
                                    if false {
                                        if ((!false) && ((0.1 - 0.0) != (3.14159 * -2.25))) {
                                            v14 := GenS2{F0: 9223372036854775807}
                                            v15 := GenS3{F0: 4294967295, F1: (v6 ^ v6)}
                                            _ = v14
                                            _ = v15
                                        }
                                    }
                                    if (true && false) {
                                        var v16 string = "a"
                                        _ = v16
                                    } else {
                                        v2 = (v2 % (((v2 & v2)) | 1))
                                    }
                                    v1 = 2147483647
                                } else {
                                    var v17 string = "a"
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if (true || false) {
                                        v18 := GenS3{F0: 42, F1: v6}
                                        _ = v18
                                    } else {
                                        v17 = ((v17 + "a") + (v17 + v17))
                                    }
                                    _ = v17
                                }
                                _ = v13
                            } else {
                                v1 -= -1
                                if false {
                                    v1 = v1
                                }
                            }
                            var v19 bool = (true && true)
                            _ = fi10
                            _ = v19
                        }
                    }
                } else {
                    fmt.Println(v1, fi7, v5)
                    v1 += v1
                }
                fmt.Println(fi7)
                _ = v8
            }
        default:
            v20 := GenS3{F0: 65535, F1: (v6 + (v6 >> ((2) & 63)))}
            _ = v20
        }
        var v21 int16 = 42
        _ = fi7
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        for fi23 := 0; fi23 < 1; fi23++ {
            v6 -= v6
            var v24 uint32 = 42
            var v25 bool = (949 != 2)
            v26 := GenS2{F0: -978}
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        if false {
            var v27 uint32 = 42
            fmt.Println(v27)
            _ = v27
        }
        var v28 uint = 65535
        _ = fi22
        _ = v28
    }
    v5 -= (v5 >> ((42) & 63))
    if (234 <= 32767) {
        if false {
            v29 := GenS1{F0: 2, F1: 0}
            if ((32767 < 32767) || false) {
                v30 := GenS4{F0: 840, F1: 0}
                _ = v30
            } else {
                switch -9223372036854775808 {
                case 9223372036854775807:
                    v2 -= v2
                case 128:
                    v31 := GenS1{F0: 1, F1: v2}
                    _ = v31
                }
            }
            _ = v29
        }
    } else {
        fmt.Println(v6, v6)
        v3 = append(v3, ("foo bar" + (("foo bar" + "foo bar") + ("\n\t" + "foo bar"))))
        var v32 float64 = (-2.25 / (3.14159 + ((1.5 * (0.0 + (0.1 * 1e10))) + (-2.25 * 3.14159))))
        _ = v32
    }
    if ((int(v6) != -2147483648) || (!false)) {
        for fi33 := 0; fi33 < 14; fi33++ {
            v6 += (v6 - 2)
            _ = fi33
        }
        v34 := GenS3{F0: 1, F1: 42}
        _ = v34
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        var v36 float64 = ((3.14159 - 1e10) / 0.1)
        switch fi35 {
        case 4294967295:
            for fi37 := 0; fi37 < 17; fi37++ {
                if len(v3) > 0 {
                    v3[0] = ("" + (("hello" + (("hello" + "hello") + ("" + ("\n\t" + "hello")))) + ""))
                }
                _ = fi37
            }
        }
        _ = fi35
        _ = v36
    }
    if false {
        if len(v4) > 0 {
            v4[0] = 1
        }
        v1 -= 642
        fmt.Println(v1, v5, v6)
    } else {
        switch 4294967295 {
        case 2147483647:
            v6 -= v6
        case 31:
            if true {
                if len(v4) > 0 {
                    v4[0] = 32767
                }
            } else {
                var v38 int8 = (v5 % ((v5) | 1))
                _ = v38
            }
        }
    }
    if false {
        if len(v3) > 0 {
            v3[0] = ""
        }
        for fi39 := 0; fi39 < 14; fi39++ {
            if true {
                _ = len(v4)
            }
            v6 = (v6 * (^(v6)))
            v3 = append(v3, "hello")
            _ = fi39
        }
        switch 2147483647 {
        case -2147483648:
            switch ("a" + ((("hello" + "hello") + "foo bar") + ("a" + ("" + "hello")))) {
            case "a":
                for fi40 := 0; fi40 < 5; fi40++ {
                    if (127 >= (v6 + v6)) {
                        v41 := GenS5{F0: 18446744073709551615}
                        fmt.Println(v2, v5, v5)
                        switch ("a" + (("\n\t" + ("foo bar" + "a")) + "a")) {
                        case "foo bar":
                            var v42 int = fi40
                            _ = v42
                        case "hello":
                            fmt.Println(v6, v5, v6)
                        case "\n\t":
                            switch fi40 {
                            case 1:
                                v1 = (v1 << ((1) & 63))
                            case 128:
                                _ = len(v3)
                            default:
                                var v43 int8 = (v5 | (v5 >> ((2) & 63)))
                                _ = v43
                            }
                        }
                        _ = v41
                    } else {
                        var v44 uint64 = (v2 * v2)
                        var v45 int16 = 159
                        for fi46 := 0; fi46 < 15; fi46++ {
                            switch "foo bar" {
                            case "hello":
                                var v47 float64 = (0.1 / 1.5)
                                _ = v47
                            case "\n\t":
                                if (18446744073709551615 > 1) {
                                    if (770 >= 42) {
                                        v1 = -2147483648
                                        switch ("foo bar" + ("\n\t" + "hello")) {
                                        case "":
                                            switch "a" {
                                            case "foo bar":
                                                if true {
                                                    v48 := GenS2{F0: 9223372036854775807}
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    if (-2.25 < -2.25) {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        var v49 uint32 = 2147483648
                                                        _ = v49
                                                    } else {
                                                        _ = len(v3)
                                                        if true {
                                                            v5 += v5
                                                            v3 = append(v3, (("" + "") + ("hello" + ("" + ("" + "foo bar")))))
                                                            v50 := GenS2{F0: -788}
                                                            _ = v50
                                                        } else {
                                                            if len(v4) > 0 {
                                                                v4[0] = fi46
                                                            }
                                                        }
                                                    }
                                                    _ = v48
                                                }
                                            default:
                                                switch int(v44) {
                                                case 4294967295:
                                                    v2 -= v44
                                                case 881:
                                                    v44 = (v2 << ((2) & 63))
                                                    if (!(fi46 > fi46)) {
                                                        v51 := GenS6{F0: 65535, F1: -735}
                                                        _ = v51
                                                    }
                                                    if false {
                                                        v4 = append(v4, (fi40 << ((4294967295) & 63)))
                                                        if (1e10 > 0.0) {
                                                            v52 := GenS1{F0: 437, F1: 405}
                                                            var v53 int32 = v1
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                        fmt.Println(v5)
                                                    } else {
                                                        fmt.Println(v44)
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                }
                                            }
                                        case "\n\t":
                                            v5 = (v5 / ((v5) | 1))
                                        case "foo bar":
                                            _ = len(v4)
                                        default:
                                            var v54 uint64 = 42
                                            _ = v54
                                        }
                                        v55 := GenS6{F0: 2, F1: 42}
                                        _ = v55
                                    }
                                    fmt.Println(fi46, v45)
                                    if len(v4) > 0 {
                                        v4[0] = 4294967295
                                    }
                                    var v56 int64 = -9223372036854775808
                                    switch "hello" {
                                    case "hello":
                                        if false {
                                            fmt.Println(v1, v45, v45)
                                            switch 1 {
                                            case 138:
                                                if len(v3) > 0 {
                                                    v3[0] = "\n\t"
                                                }
                                            case 2147483647:
                                                fmt.Println(fi46)
                                            case 9223372036854775807:
                                                var v57 int16 = v45
                                                switch "" {
                                                case "foo bar":
                                                    if len(v3) > 0 {
                                                        v3[0] = (("" + "foo bar") + "")
                                                    }
                                                    _ = len(v4)
                                                    switch "\n\t" {
                                                    case "":
                                                        v44 -= uint64(v6)
                                                    case "hello":
                                                        if true {
                                                            var v58 int64 = v56
                                                            _ = v58
                                                        } else {
                                                            v3 = append(v3, "\n\t")
                                                            if true {
                                                                var v59 float64 = ((((1e10 + (-2.25 - (((((3.14159 - 0.1) - 0.0) - ((0.0 - (-2.25 / 1.0)) / -2.25)) - 1.5) * -2.25))) / 0.1) - 1.5) + (1e10 / -2.25))
                                                                _ = v59
                                                            } else {
                                                                switch (-(fi46)) {
                                                                case 32768:
                                                                    v4 = append(v4, (fi46 ^ (fi40 | fi40)))
                                                                case 127:
                                                                    v60 := GenS7{F0: int64(v5)}
                                                                    if true {
                                                                        v61 := GenS8{F0: ((!((fi40 ^ (fi40 % ((-118) | 1))) > fi46)) || false)}
                                                                        fmt.Println(v2, v45, fi40)
                                                                        v62 := GenS3{F0: 817, F1: v6}
                                                                        _ = v61
                                                                        _ = v62
                                                                    } else {
                                                                        if len(v4) > 0 {
                                                                            v4[0] = fi46
                                                                        }
                                                                        if (false && true) {
                                                                            v5 += (-(v5))
                                                                            if false {
                                                                                var v63 string = "foo bar"
                                                                                var v64 uint8 = v6
                                                                                _ = v63
                                                                                _ = v64
                                                                            } else {
                                                                                var v65 string = ((("\n\t" + "") + (("hello" + ((("hello" + ("\n\t" + ("foo bar" + ("a" + (("hello" + ("hello" + ("hello" + ("a" + ("hello" + "\n\t"))))) + (((("a" + "\n\t") + ("\n\t" + (("\n\t" + "") + "\n\t"))) + "hello") + "\n\t")))))) + "a") + "hello")) + ("a" + (("a" + "hello") + "foo bar")))) + "")
                                                                                _ = v65
                                                                            }
                                                                        } else {
                                                                            _ = len(v4)
                                                                            if true {
                                                                                v44 -= (v44 >> ((255) & 63))
                                                                                if (!true) {
                                                                                    switch (fi46 & 32767) {
                                                                                    case 190:
                                                                                        switch "a" {
                                                                                        case "":
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = ("" + "")
                                                                                            }
                                                                                        case "a":
                                                                                            fmt.Println(v44)
                                                                                        case "hello":
                                                                                            v66 := GenS9{F0: int16(v5), F1: "foo bar"}
                                                                                            _ = v66
                                                                                        default:
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = "foo bar"
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        var v67 int32 = v1
                                                                                        _ = v67
                                                                                    }
                                                                                    switch "hello" {
                                                                                    case "":
                                                                                        if true {
                                                                                            v68 := GenS10{F0: (!(uint8(v44) >= 2))}
                                                                                            v69 := GenS5{F0: 1}
                                                                                            if true {
                                                                                                v5 = v5
                                                                                                v57 -= (v57 << ((255) & 63))
                                                                                                _ = len(v4)
                                                                                            }
                                                                                            v45 -= int16(fi40)
                                                                                            _ = v68
                                                                                            _ = v69
                                                                                        } else {
                                                                                            var v70 uint32 = 1
                                                                                            v57 = v57
                                                                                            _ = v70
                                                                                        }
                                                                                    case "a":
                                                                                        var v71 int = 32768
                                                                                        _ = v71
                                                                                    case "hello":
                                                                                        switch (fi46 | fi46) {
                                                                                        case -332:
                                                                                            var v72 float32 = 3.14159
                                                                                            _ = v72
                                                                                        }
                                                                                        v56 += v56
                                                                                        switch (fi40 - (fi46 * fi40)) {
                                                                                        case -9223372036854775808:
                                                                                            v73 := GenS11{F0: false, F1: 65, F2: (v1 >> ((1) & 63))}
                                                                                            _ = v73
                                                                                        default:
                                                                                            if (-84 != (v56 ^ v56)) {
                                                                                                var v74 uint16 = 42
                                                                                                var v75 int8 = (v5 % ((v5) | 1))
                                                                                                _ = v74
                                                                                                _ = v75
                                                                                            }
                                                                                        }
                                                                                        if (!((v57 & (v57 & 32767)) < (v45 % ((-177) | 1)))) {
                                                                                            switch "a" {
                                                                                            case "\n\t":
                                                                                                if true {
                                                                                                    switch fi46 {
                                                                                                    case -9223372036854775808:
                                                                                                        var v76 int8 = 89
                                                                                                        _ = v76
                                                                                                    default:
                                                                                                        var v77 float64 = (3.14159 / 1.5)
                                                                                                        _ = v77
                                                                                                    }
                                                                                                    v5 = 127
                                                                                                    if len(v4) >= 2 {
                                                                                                        v4 = v4[1:len(v4)]
                                                                                                    }
                                                                                                    switch (fi40 * int(v56)) {
                                                                                                    case 9223372036854775807:
                                                                                                        fmt.Println(v5, v5)
                                                                                                    default:
                                                                                                        fmt.Println(fi46, v45, v45)
                                                                                                        fmt.Println(v5, fi40)
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            if false {
                                                                                                var v78 uint = 1
                                                                                                var v79 int32 = (^(v1))
                                                                                                var v80 int64 = (v56 << ((265) & 63))
                                                                                                v4 = append(v4, fi40)
                                                                                                _ = v78
                                                                                                _ = v79
                                                                                                _ = v80
                                                                                            }
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = fi40
                                                                                            }
                                                                                            v81 := GenS8{F0: (false && (!false))}
                                                                                            _ = v81
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v82 := GenS10{F0: (!(!((v57 ^ 32767) < v45)))}
                                                                                _ = v82
                                                                            }
                                                                        }
                                                                    }
                                                                    switch "" {
                                                                    case "":
                                                                        v3 = append(v3, "foo bar")
                                                                    case "foo bar":
                                                                        if (false && true) {
                                                                            v5 -= (^(v5))
                                                                        }
                                                                    }
                                                                    _ = v60
                                                                }
                                                            }
                                                            v44 = (v2 + (v44 | (^(v2))))
                                                            v45 -= v57
                                                        }
                                                    default:
                                                        if (false && true) {
                                                            var v83 int64 = v56
                                                            v56 = (v83 * (v56 / ((int64(v57)) | 1)))
                                                            v83 = v56
                                                            _ = v83
                                                        } else {
                                                            var v84 int64 = int64(v44)
                                                            _ = v84
                                                        }
                                                    }
                                                    v85 := GenS5{F0: 2}
                                                    _ = v85
                                                case "a":
                                                    if false {
                                                        fmt.Println(v56, fi46, v57)
                                                        var v86 bool = true
                                                        v87 := GenS12{F0: uint8(v2), F1: "a"}
                                                        _ = v86
                                                        _ = v87
                                                    }
                                                case "\n\t":
                                                    v45 = (-(v45))
                                                }
                                                v88 := GenS13{F0: ("" >= "a"), F1: (v5 * (v5 + v5))}
                                                _ = v57
                                                _ = v88
                                            }
                                            switch fi40 {
                                            case -724:
                                                fmt.Println(fi46)
                                            case 32768:
                                                v5 = v5
                                            case -9223372036854775808:
                                                fmt.Println(v5, v6, v45)
                                            }
                                        } else {
                                            switch (("" + "\n\t") + "") {
                                            case "foo bar":
                                                v2 = (^(v2))
                                            case "":
                                                var v89 int64 = 548
                                                _ = v89
                                            default:
                                                var v90 uint32 = uint32(v2)
                                                _ = v90
                                            }
                                            switch 42 {
                                            case 32767:
                                                fmt.Println(v45, fi40, v5)
                                            }
                                        }
                                        v45 = (v45 << ((589) & 63))
                                        var v91 uint32 = uint32(v6)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v91
                                    }
                                    _ = v56
                                } else {
                                    var v92 uint32 = 626
                                    v93 := GenS14{F0: v45, F1: false}
                                    _ = v92
                                    _ = v93
                                }
                                var v94 int32 = int32(v45)
                                v95 := GenS9{F0: (v45 / ((v45) | 1)), F1: "\n\t"}
                                _ = v94
                                _ = v95
                            case "foo bar":
                                switch (fi40 - fi40) {
                                case 93:
                                    var v96 int = (fi40 / ((fi40) | 1))
                                    _ = v96
                                case 4294967295:
                                    v97 := GenS3{F0: 669, F1: (v6 - 1)}
                                    if (-379 <= -1) {
                                        v4 = append(v4, fi40)
                                        var v98 float64 = 0.0
                                        if true {
                                            v4 = append(v4, (-(fi46)))
                                        }
                                        _ = v98
                                    }
                                    _ = v97
                                case 0:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            default:
                                v6 = v6
                            }
                            var v99 uint8 = (v6 + v6)
                            _ = fi46
                            _ = v99
                        }
                        _ = v44
                        _ = v45
                    }
                    _ = fi40
                }
            case "foo bar":
                v5 -= v5
            case "":
                var v100 uint32 = 0
                _ = v100
            }
        case -1:
            fmt.Println(v2, v5, v1)
        case -702:
            var v101 int8 = (-(v5))
            _ = v101
        }
    } else {
        if true {
            for fi102 := 0; fi102 < 11; fi102++ {
                v103 := GenS15{F0: 1, F1: ("\n\t" + ("" + "a"))}
                _ = fi102
                _ = v103
            }
        }
    }
    fmt.Println(v1)
    v2 += (v2 >> ((1) & 63))
    fmt.Println(v2, v6, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
