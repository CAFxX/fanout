package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 bool; F1 int32; F2 string }
type GenS8 struct { F0 int8; F1 int; F2 uint64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int8; F1 int; F2 uint }
type GenS11 struct { F0 string; F1 uint32; F2 uint8 }

func main() {
    var v1 uint16 = 32767
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]uint32)
    v4 := []int{}
    var v5 int64 = v2
    if false {
        if false {
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 uint64 = 549
                v8 := GenS1{F0: 32}
                _ = fi6
                _ = v7
                _ = v8
            }
        } else {
            var v9 bool = (23 < 1)
            fmt.Println(v1, v9, v5)
            _ = v9
        }
        v1 -= (^(v1))
    } else {
        if (v1 != (v1 % (((v1 | 255)) | 1))) {
            for fi10 := 0; fi10 < 3; fi10++ {
                if false {
                    v1 = 42
                }
                _ = fi10
            }
            switch 127 {
            case -2147483648:
                v5 += 0
            case 4294967295:
                var v11 int64 = v2
                _ = v11
            }
            var v12 float64 = float64(0.1)
            _ = v12
        }
    }
    switch 1 {
    case 9223372036854775807:
        fmt.Println(v5, v2)
    case -765:
        var v13 int16 = -1
        _ = v13
    }
    v14 := GenS2{F0: (^(v2)), F1: 244}
    var v15 int8 = 100
    v3["\n\t"] = 128
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 15; fi17++ {
            if false {
                if ((-2.25 <= 1.5) || (-9223372036854775808 >= -1)) {
                    var v18 int16 = -639
                    if (!true) {
                        _ = len(v4)
                        switch (fi17 ^ int(v15)) {
                        case -32:
                            v19 := GenS2{F0: v2, F1: 128}
                            _ = v19
                        case -1:
                            if (false || ("a" <= "hello")) {
                                if true {
                                    v2 = (^(v5))
                                    var v20 uint32 = 2147483647
                                    if (!(18446744073709551615 <= 2)) {
                                        fmt.Println(v2)
                                        v15 = v15
                                        v21 := v3[("a" + ("foo bar" + (("a" + "hello") + "\n\t")))]
                                        _ = v21
                                    } else {
                                        switch (fi17 << ((255) & 63)) {
                                        case 0:
                                            v20 = 2147483647
                                        default:
                                            var v22 uint16 = 366
                                            _ = v22
                                        }
                                        var v23 uint32 = (v20 + v20)
                                        v1 -= 2
                                        _ = v23
                                    }
                                    v4 = append(v4, int(v15))
                                    var v24 int8 = 123
                                    fmt.Println(v5, v24, fi17)
                                    _ = v20
                                    _ = v24
                                }
                            } else {
                                v18 -= v18
                                switch "" {
                                case "hello":
                                    v18 -= v18
                                case "a":
                                    if true {
                                        if true {
                                            v2 += 9223372036854775807
                                            var v25 string = ("" + ("" + "a"))
                                            var v26 int32 = -790
                                            _ = v25
                                            _ = v26
                                        } else {
                                            var v27 uint8 = 216
                                            _ = v27
                                        }
                                        switch "\n\t" {
                                        case "":
                                            if false {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                var v28 int8 = v15
                                                _ = v28
                                            } else {
                                                fmt.Println(v1)
                                                switch "\n\t" {
                                                case "a":
                                                    v18 -= v18
                                                }
                                            }
                                        case "hello":
                                            v5 += v2
                                        case "foo bar":
                                            fmt.Println(fi17, fi17)
                                        }
                                    } else {
                                        if false {
                                            var v29 float32 = float32(3.14159)
                                            v18 = (v18 * (v18 % (((v18 * (v18 ^ (v18 + (^(v18)))))) | 1)))
                                            _ = v29
                                        } else {
                                            if true {
                                                v30 := GenS3{F0: v2, F1: 1}
                                                if true {
                                                    v15 += v15
                                                    if (0 < 490) {
                                                        var v31 uint64 = 42
                                                        _ = v31
                                                    }
                                                }
                                                v3["\n\t"] = 2147483647
                                                _ = v30
                                            }
                                            if (1 < 865) {
                                                var v32 int8 = v15
                                                v33 := v3["hello"]
                                                _ = v32
                                                _ = v33
                                            }
                                        }
                                        if false {
                                            if false {
                                                var v34 float64 = float64(((float32(3.14159) + 0.0) * float32(((-2.25 * float64(1e10)) + 0.1))))
                                                _ = v34
                                            }
                                            v1 += (v1 + v1)
                                        } else {
                                            var v35 float64 = float64(1e10)
                                            _ = v35
                                        }
                                        v2 += 42
                                    }
                                default:
                                    v5 = v2
                                }
                                if ("foo bar" != ("hello" + "")) {
                                    switch ("" + "") {
                                    case "\n\t":
                                        var v36 uint32 = 65535
                                        _ = v36
                                    }
                                }
                            }
                            var v37 int8 = 0
                            var v38 uint8 = 127
                            _ = v37
                            _ = v38
                        }
                        fmt.Println(v1)
                    } else {
                        v18 += (^(v18))
                        switch int(v15) {
                        case 579:
                            var v39 string = ("hello" + ("a" + ("\n\t" + "\n\t")))
                            _ = v39
                        case 9223372036854775807:
                            var v40 int8 = (v15 % (((v15 - 81)) | 1))
                            _ = v40
                        }
                        var v41 int64 = (v5 - (v2 + int64(v15)))
                        var v42 int64 = v5
                        if (!false) {
                            switch 75 {
                            case 127:
                                fmt.Println(v15, fi17)
                            case -743:
                                fmt.Println(v18, fi17)
                            default:
                                v15 = v15
                            }
                            if len(v4) > 0 {
                                v4[0] = (fi17 % ((128) | 1))
                            }
                        } else {
                            fmt.Println(v18, v2)
                        }
                        v15 = (v15 + (v15 << ((42) & 63)))
                        _ = v41
                        _ = v42
                    }
                    var v43 int16 = v18
                    _ = v18
                    _ = v43
                } else {
                    v3["a"] = 65535
                    delete(v3, "foo bar")
                    var v44 int32 = 258
                    fmt.Println(v15, v2)
                    _ = v44
                }
                if ("foo bar" != "") {
                    v1 = 345
                }
                fmt.Println(v1, v1, v2)
                v45 := GenS4{F0: -815, F1: 18446744073709551615}
                _ = v45
            } else {
                v4 = append(v4, (fi17 ^ fi16))
                _ = len(v3)
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = len(v3)
            var v46 bool = true
            if (0.0 < 3.14159) {
                var v47 uint8 = 42
                _ = v47
            } else {
                switch "" {
                case "hello":
                    v5 = (v2 + v2)
                }
                _ = len(v3)
            }
            _ = fi17
            _ = v46
        }
        v48 := GenS3{F0: v5, F1: 213}
        _ = fi16
        _ = v48
    }
    if (794 < 18446744073709551615) {
        switch ("a" + "\n\t") {
        case "\n\t":
            if (!(false && false)) {
                if false {
                    v5 += v5
                    for fi49 := 0; fi49 < 15; fi49++ {
                        for fi50 := 0; fi50 < 20; fi50++ {
                            if len(v4) > 0 {
                                v4[0] = fi49
                            }
                            v15 = v15
                            _ = len(v3)
                            if false {
                                fmt.Println(v5, v15)
                                if false {
                                    var v51 uint32 = 1
                                    _ = v51
                                }
                                fmt.Println(v1, fi49)
                            } else {
                                if (("foo bar" + "") >= "foo bar") {
                                    if false {
                                        var v52 int16 = 1
                                        var v53 float64 = (0.1 + -2.25)
                                        _ = v52
                                        _ = v53
                                    }
                                    v2 = (v5 << ((962) & 63))
                                } else {
                                    fmt.Println(fi49, v5)
                                }
                                fmt.Println(v5, v1, v5)
                                if false {
                                    fmt.Println(v15, fi50)
                                    var v54 float64 = 1e10
                                    var v55 int16 = 0
                                    _ = v54
                                    _ = v55
                                } else {
                                    var v56 int16 = 35
                                    var v57 float64 = 1.5
                                    _ = v56
                                    _ = v57
                                }
                            }
                            _ = fi50
                        }
                        v5 = -9223372036854775808
                        _ = fi49
                    }
                } else {
                    if false {
                        var v58 int8 = (v15 & v15)
                        if false {
                            _ = len(v3)
                        } else {
                            v5 = 42
                        }
                        switch -802 {
                        case 42:
                            switch ("" + "a") {
                            case "hello":
                                for fi59 := 0; fi59 < 19; fi59++ {
                                    for fi60 := 0; fi60 < 1; fi60++ {
                                        v61 := GenS5{F0: 1}
                                        _ = fi60
                                        _ = v61
                                    }
                                    if true {
                                        fmt.Println(v2)
                                    }
                                    v3["foo bar"] = 42
                                    _ = fi59
                                }
                            case "\n\t":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        case 850:
                            fmt.Println(v5, v2, v58)
                        case 951:
                            fmt.Println(v2, v2, v2)
                        }
                        _ = v58
                    } else {
                        _ = len(v3)
                        v1 = 0
                        switch 42 {
                        case 643:
                            var v62 uint = 4294967295
                            _ = v62
                        default:
                            if ((v1 / ((v1) | 1)) != v1) {
                                var v63 int64 = 184
                                if true {
                                    var v64 uint = 255
                                    for fi65 := 0; fi65 < 5; fi65++ {
                                        v4 = append(v4, 32768)
                                        v66 := GenS6{F0: "a", F1: 4294967295}
                                        v63 += int64(fi65)
                                        _ = fi65
                                        _ = v66
                                    }
                                    v4 = append(v4, 127)
                                    _ = v64
                                } else {
                                    var v67 float64 = (1e10 - 3.14159)
                                    _ = v67
                                }
                                _ = v63
                            }
                        }
                    }
                }
                v68 := GenS7{F0: false, F1: 0, F2: "hello"}
                var v69 string = "hello"
                _ = v68
                _ = v69
            }
        case "hello":
            v5 = -534
        case "":
            var v70 float64 = -2.25
            _ = v70
        default:
            fmt.Println(v15)
        }
        var v71 uint8 = 42
        _ = v71
    } else {
        if (907 == -1) {
            if true {
                for fi72 := 0; fi72 < 15; fi72++ {
                    for fi73 := 0; fi73 < 9; fi73++ {
                        v15 -= v15
                        _ = fi73
                    }
                    delete(v3, ("foo bar" + "a"))
                    for fi74 := 0; fi74 < 8; fi74++ {
                        var v75 uint = 42
                        switch ("a" + (("a" + "") + ("\n\t" + "\n\t"))) {
                        case "hello":
                            fmt.Println(fi74, v2, v75)
                        default:
                            switch "foo bar" {
                            case "\n\t":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case "":
                                v76 := GenS3{F0: (^(v2)), F1: 1}
                                _ = v76
                            case "a":
                                _ = len(v4)
                                var v77 int32 = 0
                                var v78 float64 = 0.1
                                v79 := GenS8{F0: v15, F1: fi74, F2: 548}
                                v75 -= (v75 & v75)
                                _ = v77
                                _ = v78
                                _ = v79
                            default:
                                var v80 int = fi74
                                _ = v80
                            }
                        }
                        var v81 uint = 2
                        _ = fi74
                        _ = v75
                        _ = v81
                    }
                    _ = fi72
                }
            } else {
                v15 = v15
                v4 = append(v4, -1)
            }
        }
        for fi82 := 0; fi82 < 10; fi82++ {
            for fi83 := 0; fi83 < 11; fi83++ {
                var v84 uint8 = 42
                v1 = uint16(v84)
                var v85 bool = (0 >= 42)
                _ = len(v3)
                fmt.Println(fi83, v15)
                _ = fi83
                _ = v84
                _ = v85
            }
            var v86 int32 = 0
            _ = fi82
            _ = v86
        }
        v15 -= (v15 << ((255) & 63))
    }
    fmt.Println(v5, v15, v5)
    var v87 float32 = float32(0.0)
    if false {
        var v88 uint = 357
        _ = v88
    } else {
        var v89 float32 = float32((-2.25 - 0.1))
        for fi90 := 0; fi90 < 6; fi90++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v91 uint = 717
            _ = fi90
            _ = v91
        }
        _ = v89
    }
    if false {
        if len(v4) > 0 {
            v4[0] = -2147483648
        }
        for fi92 := 0; fi92 < 8; fi92++ {
            for fi93 := 0; fi93 < 10; fi93++ {
                fmt.Println(fi92, v1)
                _ = fi93
            }
            _ = fi92
        }
        var v94 float32 = (v87 / 3.14159)
        _ = v94
    }
    fmt.Println(v1)
    var v95 int32 = 0
    if true {
        v95 = v95
        if true {
            v96 := v3["hello"]
            for fi97 := 0; fi97 < 5; fi97++ {
                if true {
                    v5 += v5
                }
                v98 := GenS5{F0: v95}
                _ = len(v4)
                _ = fi97
                _ = v98
            }
            v99 := GenS9{F0: "foo bar"}
            _ = v96
            _ = v99
        } else {
            for fi100 := 0; fi100 < 19; fi100++ {
                switch (^(fi100)) {
                case 0:
                    v5 += v5
                case 4294967295:
                    switch "foo bar" {
                    case "":
                        v101 := GenS10{F0: (v15 & 0), F1: -2147483648, F2: 3}
                        _ = v101
                    case "foo bar":
                        v15 = (v15 / ((-128) | 1))
                    default:
                        var v102 bool = true
                        _ = v102
                    }
                case 42:
                    var v103 float32 = v87
                    _ = v103
                }
                v2 += (-(v5))
                v87 = v87
                _ = fi100
            }
            for fi104 := 0; fi104 < 12; fi104++ {
                v3["a"] = 42
                var v105 uint32 = 660
                fmt.Println(v15, v15, v2)
                _ = fi104
                _ = v105
            }
            if (-241 >= (v5 * (v5 >> ((696) & 63)))) {
                v2 -= (v5 - v5)
            } else {
                v106 := GenS10{F0: (v15 >> ((2) & 63)), F1: 127, F2: 178}
                switch 128 {
                case 1:
                    fmt.Println(v2, v15, v95)
                case 32767:
                    switch "" {
                    case "hello":
                        var v107 int64 = v2
                        _ = v107
                    case "foo bar":
                        v108 := v3["a"]
                        _ = v108
                    case "\n\t":
                        var v109 string = ""
                        _ = v109
                    }
                case 42:
                    var v110 float32 = float32(-2.25)
                    _ = v110
                }
                _ = v106
            }
        }
        v1 -= (v1 / ((129) | 1))
    }
    v111 := GenS2{F0: (^(v2)), F1: uint8(v1)}
    if (31 != 67) {
        for fi112 := 0; fi112 < 10; fi112++ {
            switch fi112 {
            case 500:
                v4 = append(v4, (fi112 + fi112))
            case -665:
                for fi113 := 0; fi113 < 15; fi113++ {
                    v5 = v5
                    if true {
                        var v114 int32 = (v95 << ((255) & 63))
                        v5 = 14
                        v115 := GenS6{F0: "", F1: 1}
                        v116 := GenS11{F0: "a", F1: 904, F2: 26}
                        var v117 bool = true
                        _ = v114
                        _ = v115
                        _ = v116
                        _ = v117
                    }
                    v118 := v3["\n\t"]
                    _ = fi113
                    _ = v118
                }
            default:
                v15 -= v15
            }
            _ = fi112
        }
        v2 = v5
        var v119 int32 = v95
        _ = v119
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v87:", v87)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v111
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v87
    _ = v95
}
