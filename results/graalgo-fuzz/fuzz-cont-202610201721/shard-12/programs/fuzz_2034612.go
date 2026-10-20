package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 uint32; F2 uint64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint64; F1 int32 }
type GenS9 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint64; F1 int64; F2 uint64 }

func main() {
    var v1 uint32 = 223
    v2 := []uint32{(v1 * v1), (v1 & 4294967295)}
    var v3 uint32 = v1
    v3 = v3
    if (18446744073709551615 != 18446744073709551615) {
        v1 = 1
        v1 += (v3 + v1)
        if true {
            switch 322 {
            case -419:
                switch "a" {
                case "a":
                    var v4 uint = 290
                    _ = v4
                default:
                    v5 := GenS1{F0: 4294967295, F1: 2, F2: 65535}
                    _ = v5
                }
            case -1:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            if len(v2) > 0 {
                v2[0] = 2
            }
            v1 -= v1
        } else {
            v6 := GenS2{F0: -655, F1: 127}
            switch "\n\t" {
            case "\n\t":
                if true {
                    var v7 int8 = 127
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    _ = v7
                } else {
                    fmt.Println(v1)
                    v1 += 2147483647
                }
            case "foo bar":
                var v8 uint8 = 42
                _ = v8
            case "":
                var v9 uint64 = 170
                _ = v9
            }
            v3 = (^(v1))
            _ = v6
        }
    } else {
        if false {
            v1 = v3
        } else {
            if len(v2) > 0 {
                v2[0] = v1
            }
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := GenS2{F0: -795, F1: (fi10 % ((-309) | 1))}
                _ = fi10
                _ = v11
            }
        }
    }
    switch 102 {
    case 2147483647:
        var v12 string = ("a" + "")
        _ = v12
    case 736:
        v13 := GenS1{F0: 0, F1: 2, F2: 2}
        _ = v13
    }
    if (-1 != 0) {
        v1 += 0
        for fi14 := 0; fi14 < 20; fi14++ {
            fmt.Println(v3)
            var v15 bool = false
            _ = fi14
            _ = v15
        }
    }
    switch "hello" {
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = v3
        }
    case "a":
        for fi16 := 0; fi16 < 5; fi16++ {
            switch -9223372036854775808 {
            case -674:
                for fi17 := 0; fi17 < 15; fi17++ {
                    fmt.Println(v1, fi17, fi16)
                    _ = fi17
                }
            case 4294967295:
                v18 := GenS2{F0: 760, F1: fi16}
                _ = v18
            }
            v19 := GenS3{F0: 255, F1: -2147483648}
            var v20 int = (fi16 % ((127) | 1))
            _ = fi16
            _ = v19
            _ = v20
        }
    case "hello":
        v3 = v1
    default:
        v3 = v1
    }
    switch 956 {
    case -845:
        v1 -= (v3 - (v1 % (((v3 * (^(v3)))) | 1)))
    default:
        switch 128 {
        case 32767:
            var v21 int64 = 42
            _ = v21
        case 127:
            for fi22 := 0; fi22 < 17; fi22++ {
                if false {
                    v2 = append(v2, v3)
                    for fi23 := 0; fi23 < 1; fi23++ {
                        var v24 int64 = -1
                        _ = fi23
                        _ = v24
                    }
                    if (!true) {
                        if ((fi22 - (fi22 * fi22)) > -2147483648) {
                            fmt.Println(v3)
                        } else {
                            v3 = (v3 & (v1 >> ((243) & 63)))
                            fmt.Println(fi22)
                        }
                    } else {
                        fmt.Println(v1, fi22)
                        v3 = (v1 | 590)
                        switch fi22 {
                        case -9223372036854775808:
                            for fi25 := 0; fi25 < 9; fi25++ {
                                switch (fi22 << ((4294967295) & 63)) {
                                case -761:
                                    v1 = v3
                                case 4294967295:
                                    if (false && (3.14159 >= 0.1)) {
                                        if (42 != 1) {
                                            v1 -= v1
                                            var v26 float32 = (1.5 + (1e10 / -2.25))
                                            switch 128 {
                                            case 1:
                                                fmt.Println(v3, fi22, fi25)
                                            case 0:
                                                var v27 int32 = 2147483647
                                                _ = v27
                                            default:
                                                if (false && (!((("\n\t" + "") + "hello") > ("\n\t" + "")))) {
                                                    fmt.Println(v3, fi22)
                                                }
                                            }
                                            _ = v26
                                        } else {
                                            var v28 int8 = 42
                                            if (1.5 >= 0.0) {
                                                fmt.Println(v1, fi25)
                                            } else {
                                                if (true || true) {
                                                    var v29 bool = (4294967295 > (^(v3)))
                                                    switch ((("\n\t" + ("" + "")) + "foo bar") + "foo bar") {
                                                    case "hello":
                                                        switch ((("" + (((("" + "foo bar") + ("\n\t" + "\n\t")) + "a") + "a")) + (("a" + "") + "\n\t")) + "") {
                                                        case "":
                                                            v2 = append(v2, v1)
                                                        }
                                                    case "a":
                                                        switch "a" {
                                                        case "":
                                                            var v30 float32 = -2.25
                                                            _ = v30
                                                        case "foo bar":
                                                            var v31 float64 = 0.0
                                                            _ = v31
                                                        default:
                                                            if ((!true) && true) {
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    v2 = append(v2, v1)
                                                                case "a":
                                                                    if false {
                                                                        v29 = (944 != 619)
                                                                        if v29 {
                                                                            var v32 float64 = 0.0
                                                                            var v33 uint64 = 4294967295
                                                                            _ = v32
                                                                            _ = v33
                                                                        } else {
                                                                            var v34 int = fi22
                                                                            fmt.Println(v3, fi22)
                                                                            var v35 string = "a"
                                                                            _ = v34
                                                                            _ = v35
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    var v36 uint32 = (^(v1))
                                                                    _ = v36
                                                                default:
                                                                    v3 -= 1
                                                                }
                                                            } else {
                                                                v37 := GenS4{F0: 1}
                                                                var v38 int16 = -330
                                                                v3 -= 0
                                                                _ = v37
                                                                _ = v38
                                                            }
                                                        }
                                                    default:
                                                        switch (("foo bar" + "") + ("\n\t" + "hello")) {
                                                        case "":
                                                            if true {
                                                                fmt.Println(v1, fi25, v29)
                                                                var v39 int32 = -1
                                                                if len(v2) > 0 {
                                                                    v2[0] = 793
                                                                }
                                                                var v40 uint32 = (^(v1))
                                                                switch fi25 {
                                                                case -444:
                                                                    var v41 string = "\n\t"
                                                                    _ = v41
                                                                case -913:
                                                                    fmt.Println(v29)
                                                                default:
                                                                    v28 = (v28 << ((1) & 63))
                                                                }
                                                                _ = v39
                                                                _ = v40
                                                            }
                                                        case "foo bar":
                                                            if v29 {
                                                                v29 = (127 < fi25)
                                                            }
                                                        default:
                                                            fmt.Println(fi22)
                                                        }
                                                    }
                                                    v29 = (true && ((((v3 + v1) == v3) && ((9223372036854775807 != 706) && v29)) || (!(v29 && (!v29)))))
                                                    _ = v29
                                                } else {
                                                    v1 -= v1
                                                    v42 := GenS5{F0: ("\n\t" + "\n\t"), F1: v3, F2: 880}
                                                    _ = v42
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = (v3 + (v3 & 2147483648))
                                                }
                                                switch fi25 {
                                                case -9223372036854775808:
                                                    if (!("hello" == "\n\t")) {
                                                        var v43 bool = (312 >= 255)
                                                        var v44 int = 128
                                                        var v45 string = "\n\t"
                                                        v44 = -2147483648
                                                        switch v44 {
                                                        case 613:
                                                            if v43 {
                                                                v1 = 65535
                                                            } else {
                                                                v46 := GenS6{F0: 0}
                                                                _ = v46
                                                            }
                                                        case 0:
                                                            var v47 uint16 = 2
                                                            _ = v47
                                                        default:
                                                            v1 -= (^(v1))
                                                        }
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        fmt.Println(fi22, fi22)
                                                        fmt.Println(v28, v1)
                                                    }
                                                case -125:
                                                    if (!(v28 != int8(fi25))) {
                                                        v2 = append(v2, 2147483648)
                                                        v28 -= v28
                                                        v28 = -88
                                                    } else {
                                                        if false {
                                                            v48 := GenS5{F0: (((("a" + "a") + "a") + "\n\t") + ""), F1: 2147483648, F2: 1}
                                                            v1 = (v3 - v1)
                                                            _ = v48
                                                        } else {
                                                            v49 := GenS7{F0: 65535}
                                                            v28 = (v28 >> ((549) & 63))
                                                            v2 = append(v2, (^(v3)))
                                                            _ = v49
                                                        }
                                                    }
                                                case 873:
                                                    v1 -= v1
                                                default:
                                                    var v50 uint = 657
                                                    _ = v50
                                                }
                                            }
                                            v28 = (v28 | (v28 - v28))
                                            _ = v28
                                        }
                                    }
                                case 32768:
                                    v51 := GenS1{F0: uint64(v1), F1: 4294967295, F2: 470}
                                    _ = v51
                                default:
                                    _ = len(v2)
                                }
                                _ = len(v2)
                                fmt.Println(fi22)
                                _ = fi25
                            }
                        default:
                            v1 += v3
                        }
                    }
                } else {
                    v3 += 690
                    v52 := GenS7{F0: 1}
                    _ = v52
                }
                _ = fi22
            }
        case 370:
            for fi53 := 0; fi53 < 1; fi53++ {
                fmt.Println(fi53, fi53, v3)
                v1 += v3
                v54 := GenS8{F0: 635, F1: -77}
                _ = fi53
                _ = v54
            }
        }
    }
    if true {
        for fi55 := 0; fi55 < 19; fi55++ {
            if (42 > 231) {
                for fi56 := 0; fi56 < 3; fi56++ {
                    var v57 uint8 = 127
                    if (false || (!(true || true))) {
                        v3 += 1
                        v1 = v3
                        switch ("" + ((("hello" + (("" + "\n\t") + "foo bar")) + "a") + "\n\t")) {
                        case "foo bar":
                            v1 = 123
                        case "\n\t":
                            if (42 < 1) {
                                _ = len(v2)
                                if (-628 >= 1) {
                                    fmt.Println(v1, v57, v3)
                                    if true {
                                        v57 += v57
                                        v2 = append(v2, 0)
                                        v1 = (v1 + 356)
                                    } else {
                                        var v58 float64 = 0.1
                                        _ = v58
                                    }
                                }
                            }
                        case "hello":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                    _ = fi56
                    _ = v57
                }
                var v59 string = ("" + ("a" + (("\n\t" + "a") + "a")))
                _ = v59
            }
            var v60 uint32 = (^(v1))
            _ = fi55
            _ = v60
        }
        v3 += (^(v1))
        for fi61 := 0; fi61 < 20; fi61++ {
            v62 := GenS9{F0: 128, F1: -2147483648, F2: 0}
            _ = len(v2)
            _ = fi61
            _ = v62
        }
    } else {
        switch 2147483647 {
        case 0:
            v1 = v3
        case -556:
            if (true && (1 > 65535)) {
                v1 = (v3 - (v1 + (v3 - 4294967295)))
                var v63 int16 = 1
                _ = v63
            }
        default:
            if (0 < 172) {
                var v64 int32 = -681
                v64 += (v64 * 133)
                _ = v64
            } else {
                switch 0 {
                case 1:
                    v65 := GenS5{F0: "a", F1: (v1 - (v1 ^ (^(v1)))), F2: 4294967295}
                    _ = v65
                }
                for fi66 := 0; fi66 < 20; fi66++ {
                    for fi67 := 0; fi67 < 9; fi67++ {
                        if (!false) {
                            switch "hello" {
                            case "foo bar":
                                v1 = v1
                            default:
                                v68 := GenS10{F0: fi67}
                                _ = v68
                            }
                            switch "" {
                            case "":
                                switch "\n\t" {
                                case "foo bar":
                                    fmt.Println(v3, v1)
                                case "":
                                    var v69 int32 = -2147483648
                                    _ = v69
                                default:
                                    var v70 int32 = -1
                                    _ = v70
                                }
                            case "\n\t":
                                switch "\n\t" {
                                case "foo bar":
                                    if (-397 <= 539) {
                                        if true {
                                            switch "a" {
                                            case "\n\t":
                                                var v71 uint32 = v3
                                                _ = v71
                                            }
                                            switch "" {
                                            case "":
                                                v1 -= (v3 >> ((928) & 63))
                                            }
                                        }
                                        if ((0.0 / 1.0) <= ((((-2.25 - (0.0 + (3.14159 / (1e10 * (0.0 - (1e10 / 1e10)))))) - 1e10) + 1e10) + 1.5)) {
                                            fmt.Println(fi67, v3)
                                        } else {
                                            var v72 uint16 = 2
                                            var v73 uint64 = 516
                                            var v74 uint = 56
                                            _ = v72
                                            _ = v73
                                            _ = v74
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                case "hello":
                                    if true {
                                        v3 = v1
                                    }
                                default:
                                    var v75 int16 = 1000
                                    _ = v75
                                }
                            }
                            v1 = (v3 << ((42) & 63))
                        } else {
                            v1 = 4294967295
                            fmt.Println(fi67, v1, fi66)
                            var v76 uint64 = 0
                            _ = v76
                        }
                        switch "foo bar" {
                        case "a":
                            if (!false) {
                                if (!false) {
                                    var v77 int64 = 0
                                    v78 := GenS1{F0: uint64(v1), F1: 759, F2: 1}
                                    var v79 float32 = (1.5 + -2.25)
                                    var v80 uint32 = 2147483648
                                    v81 := GenS4{F0: int32(fi67)}
                                    v82 := GenS7{F0: 2}
                                    fmt.Println(fi67)
                                    _ = v77
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                    _ = v81
                                    _ = v82
                                }
                                v83 := GenS5{F0: "foo bar", F1: (v1 * (v3 >> ((65535) & 63))), F2: 4294967295}
                                _ = v83
                            } else {
                                var v84 uint8 = 128
                                _ = v84
                            }
                        case "\n\t":
                            v1 -= 0
                        case "hello":
                            v85 := GenS7{F0: 4294967295}
                            _ = v85
                        default:
                            v3 = v3
                        }
                        fmt.Println(v1, fi66)
                        _ = fi67
                    }
                    switch (fi66 / ((1) | 1)) {
                    case -337:
                        var v86 int = 127
                        _ = v86
                    case 4294967295:
                        v3 -= (v1 * v1)
                    default:
                        if (!((18446744073709551615 < 0) && (0.1 == 0.0))) {
                            if true {
                                if false {
                                    var v87 int = (fi66 - (fi66 + 2147483647))
                                    _ = v87
                                } else {
                                    var v88 float32 = float32((0.1 + 3.14159))
                                    v89 := GenS11{F0: 0, F1: -1, F2: 1}
                                    var v90 uint32 = 580
                                    _ = v88
                                    _ = v89
                                    _ = v90
                                }
                            } else {
                                var v91 int32 = -623
                                fmt.Println(v1, fi66)
                                var v92 int = (fi66 >> ((643) & 63))
                                _ = v91
                                _ = v92
                            }
                            var v93 int64 = 42
                            _ = v93
                        }
                    }
                    _ = fi66
                }
                v1 -= (v1 ^ 1)
            }
        }
    }
    var v94 uint8 = 42
    v95 := GenS9{F0: 0, F1: 42, F2: 1}
    fmt.Println(v1)
    var v96 uint32 = (v1 / ((428) | 1))
    switch "\n\t" {
    case "\n\t":
        var v97 uint16 = 290
        _ = v97
    case "a":
        var v98 int = 2147483647
        _ = v98
    case "foo bar":
        for fi99 := 0; fi99 < 5; fi99++ {
            fmt.Println(v94, fi99)
            _ = fi99
        }
    }
    var v100 int64 = 42
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        var v101 uint32 = (v1 / ((v3) | 1))
        _ = v101
    } else {
        switch 1 {
        case 2147483647:
            for fi102 := 0; fi102 < 3; fi102++ {
                v103 := GenS2{F0: -450, F1: -206}
                _ = fi102
                _ = v103
            }
        }
        for fi104 := 0; fi104 < 18; fi104++ {
            var v105 string = ("\n\t" + "")
            _ = fi104
            _ = v105
        }
        fmt.Println(v94, v3, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v94:", v94)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v94
    _ = v95
    _ = v96
}
