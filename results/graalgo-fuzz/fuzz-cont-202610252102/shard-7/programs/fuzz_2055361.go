package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int32; F1 int64; F2 bool }
type GenS3 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int; F2 uint32 }
type GenS5 struct { F0 uint32; F1 string; F2 string }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 int64; F1 int; F2 int16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int8; F1 uint16; F2 uint }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int32; F1 int }
type GenS14 struct { F0 uint16; F1 bool; F2 int16 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int8; F1 int32; F2 int64 }
type GenS17 struct { F0 bool; F1 int8 }
type GenS18 struct { F0 uint; F1 uint8; F2 bool }
type GenS19 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS20 struct { F0 uint16 }
type GenS21 struct { F0 bool; F1 int8; F2 int }
type GenS22 struct { F0 int32; F1 uint64 }
type GenS23 struct { F0 uint32 }
type GenS24 struct { F0 uint; F1 uint8; F2 int }
type GenS25 struct { F0 int32; F1 int32; F2 bool }
type GenS26 struct { F0 string; F1 int64 }
func genFn0(p0 uint32, p1 int8, p2 uint) uint64 {
    v7 := []int{}
    if len(v7) > 0 {
        v7[0] = -1
    }
    switch -9223372036854775808 {
    case 127:
        if (false || ((((p2 << ((p2) & 63)) != 2) && (!(false && (18446744073709551615 == 375)))) || (1 == 565))) {
            var v8 float64 = 0.0
            if (855 != 255) {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                if (p0 <= 2147483647) {
                    switch "" {
                    case "":
                        _ = len(v7)
                    case "foo bar":
                        v9 := GenS1{F0: -32768, F1: 1, F2: p0}
                        _ = v9
                    case "\n\t":
                        v10 := GenS1{F0: 408, F1: 42, F2: p0}
                        _ = v10
                    }
                    var v11 float32 = (0.1 + float32(((((((v8 + v8) + (v8 * 1e10)) / v8) + 1.5) * v8) - v8)))
                    _ = v11
                }
            } else {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            _ = v8
        }
    }
    v12 := GenS1{F0: 1, F1: 2, F2: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    return 2
}

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    v5 := make(map[string]string)
    v6 := make(map[uint32]string)
    for fi13 := 0; fi13 < 2; fi13++ {
        switch ("a" + (("hello" + "a") + "hello")) {
        case "hello":
            for fi14 := 0; fi14 < 8; fi14++ {
                v1 += fi13
                _ = fi14
            }
        case "\n\t":
            v1 += (v1 << ((4294967295) & 63))
        case "":
            if true {
                _ = len(v6)
                switch 32767 {
                case 1:
                    if false {
                        v15 := GenS1{F0: -295, F1: 638, F2: 0}
                        v1 += (fi13 << ((42) & 63))
                        v16 := []int{}
                        _ = v15
                        _ = v16
                    }
                case -1:
                    if (32767 == -1) {
                        if false {
                            delete(v3, "")
                            if (156 > 200) {
                                v1 = 0
                                v1 = 348
                                switch fi13 {
                                case 338:
                                    var v17 string = "hello"
                                    _ = v17
                                case 828:
                                    var v18 uint64 = 18446744073709551615
                                    _ = v18
                                case -702:
                                    switch (("hello" + ("a" + "a")) + "") {
                                    case "a":
                                        var v19 int16 = -212
                                        _ = v19
                                    case "":
                                        _ = len(v2)
                                    case "hello":
                                        if true {
                                            if ((!false) && false) {
                                                v20 := []int{fi13, (fi13 ^ (fi13 | 0)), (^(fi13)), (fi13 >> ((4294967295) & 63))}
                                                for fi21 := 0; fi21 < 3; fi21++ {
                                                    if false {
                                                        v22 := GenS1{F0: -1, F1: 0, F2: 2147483647}
                                                        switch (("a" + ("foo bar" + ("" + ("foo bar" + ((((("foo bar" + "a") + "") + (("foo bar" + ("hello" + ("foo bar" + "foo bar"))) + "\n\t")) + "\n\t") + ((("a" + ("hello" + "\n\t")) + "") + "\n\t")))))) + "") {
                                                        case "\n\t":
                                                            var v23 int16 = 42
                                                            _ = v23
                                                        case "":
                                                            fmt.Println(v1, fi21)
                                                        case "foo bar":
                                                            var v24 int8 = 123
                                                            _ = v24
                                                        default:
                                                            var v25 float64 = -2.25
                                                            _ = v25
                                                        }
                                                        _ = v22
                                                    } else {
                                                        v1 = (fi21 / ((v1) | 1))
                                                    }
                                                    _ = fi21
                                                }
                                                _ = v20
                                            }
                                        } else {
                                            v26 := v4["a"]
                                            v6[2147483648] = "a"
                                            _ = v26
                                        }
                                    }
                                }
                            }
                        } else {
                            var v27 int = 341
                            switch (v27 % ((1) | 1)) {
                            case 2147483647:
                                v1 += v27
                            case 9223372036854775807:
                                switch (v27 / ((v27) | 1)) {
                                case -2147483648:
                                    if true {
                                        var v28 float64 = -2.25
                                        if (!(true && false)) {
                                            if false {
                                                var v29 int = (v1 + 9223372036854775807)
                                                _ = v29
                                            } else {
                                                for fi30 := 0; fi30 < 16; fi30++ {
                                                    var v31 int64 = -9223372036854775808
                                                    var v32 uint8 = 42
                                                    _ = fi30
                                                    _ = v31
                                                    _ = v32
                                                }
                                                fmt.Println(v28, v28, fi13)
                                                if true {
                                                    var v33 int16 = 42
                                                    if (548 < 42) {
                                                        var v34 int = (v1 & -693)
                                                        switch (fi13 % (((v34 * v27)) | 1)) {
                                                        case 194:
                                                            for fi35 := 0; fi35 < 14; fi35++ {
                                                                if (true || (1 >= 18446744073709551615)) {
                                                                    v34 = (-(fi35))
                                                                    v33 -= v33
                                                                    v4[("a" + "")] = v34
                                                                    var v36 bool = (!((v27 ^ (v34 ^ fi13)) <= fi35))
                                                                    _ = v36
                                                                } else {
                                                                    var v37 int8 = -1
                                                                    _ = v37
                                                                }
                                                                v38 := []int{}
                                                                v39 := GenS2{F0: 908, F1: -1, F2: ((v33 & v33) <= v33)}
                                                                if false {
                                                                    var v40 int = 664
                                                                    fmt.Println(v27, v33)
                                                                    _ = v40
                                                                }
                                                                if (1.5 > 3.14159) {
                                                                    v6[794] = ("" + "\n\t")
                                                                }
                                                                if ((!false) && (1.5 <= float64(0.0))) {
                                                                    var v41 uint64 = 2
                                                                    v34 += (^(fi13))
                                                                    if (2 < 1) {
                                                                        fmt.Println(v28, v33, fi35)
                                                                    } else {
                                                                        v41 += 42
                                                                        v42 := GenS1{F0: (v33 + 1), F1: 65535, F2: 65535}
                                                                        _ = v42
                                                                    }
                                                                    if false {
                                                                        v43 := GenS3{F0: 782, F1: false, F2: 0}
                                                                        v27 += (^(fi13))
                                                                        var v44 uint8 = 79
                                                                        _ = v43
                                                                        _ = v44
                                                                    } else {
                                                                        delete(v2, (fi35 ^ 42))
                                                                        v28 += v28
                                                                        v27 -= (v34 >> ((68) & 63))
                                                                    }
                                                                    _ = v41
                                                                }
                                                                v45 := GenS3{F0: 2, F1: (v27 == (fi13 & (v27 << ((332) & 63)))), F2: 65535}
                                                                var v46 uint64 = 42
                                                                _ = fi35
                                                                _ = v38
                                                                _ = v39
                                                                _ = v45
                                                                _ = v46
                                                            }
                                                        case 495:
                                                            if false {
                                                                v33 = (v33 << ((65535) & 63))
                                                                switch "\n\t" {
                                                                case "":
                                                                    v47 := GenS3{F0: 797, F1: (!true), F2: 0}
                                                                    _ = v47
                                                                case "a":
                                                                    var v48 uint64 = genFn0(2147483647, 1, 823)
                                                                    _ = v48
                                                                }
                                                                for fi49 := 0; fi49 < 10; fi49++ {
                                                                    v34 += v34
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        if false {
                                                                            if true {
                                                                                v50 := GenS1{F0: v33, F1: 42, F2: 1}
                                                                                if true {
                                                                                    v28 += (v28 + (v28 / 1e10))
                                                                                    delete(v5, "")
                                                                                    var v51 float64 = 1.5
                                                                                    v33 = -1
                                                                                    if (426 >= 2) {
                                                                                        var v52 int64 = 42
                                                                                        _ = v52
                                                                                    }
                                                                                    _ = v51
                                                                                }
                                                                                _ = v50
                                                                            }
                                                                            switch 42 {
                                                                            case 2147483647:
                                                                                v28 += (v28 + v28)
                                                                            case 398:
                                                                                switch "foo bar" {
                                                                                case "":
                                                                                    v34 -= v1
                                                                                }
                                                                            case 9223372036854775807:
                                                                                switch v34 {
                                                                                case 32768:
                                                                                    if (!(2147483647 > -2147483648)) {
                                                                                        switch 42 {
                                                                                        case -9223372036854775808:
                                                                                            switch "foo bar" {
                                                                                            case "foo bar":
                                                                                                v1 = 959
                                                                                            }
                                                                                        default:
                                                                                            if true {
                                                                                                v53 := GenS4{F0: 65535, F1: fi13, F2: 873}
                                                                                                fmt.Println(fi13)
                                                                                                if true {
                                                                                                    v28 = ((float64(float32(float64(0.0))) / -2.25) - v28)
                                                                                                } else {
                                                                                                    if true {
                                                                                                        v54 := GenS3{F0: 510, F1: (-1 >= -9223372036854775808), F2: 65535}
                                                                                                        _ = v54
                                                                                                    } else {
                                                                                                        v33 = 767
                                                                                                    }
                                                                                                }
                                                                                                v28 = v28
                                                                                                v28 -= (v28 + v28)
                                                                                                _ = v53
                                                                                            } else {
                                                                                                var v55 int16 = (v33 * (v33 - (^(v33))))
                                                                                                var v56 int64 = 826
                                                                                                v57 := GenS4{F0: 590, F1: (^(v34)), F2: 2}
                                                                                                switch fi49 {
                                                                                                case 1:
                                                                                                    var v58 uint32 = 42
                                                                                                    _ = v58
                                                                                                case 2147483647:
                                                                                                    v33 -= v55
                                                                                                default:
                                                                                                    if true {
                                                                                                        v33 += 32767
                                                                                                        if (!(-128 <= -128)) {
                                                                                                            v59 := GenS5{F0: 624, F1: "", F2: ("hello" + "a")}
                                                                                                            fmt.Println(fi13, v33, v55)
                                                                                                            _ = v59
                                                                                                        } else {
                                                                                                            if ((1 < 443) && true) {
                                                                                                                if true {
                                                                                                                    var v60 uint16 = 942
                                                                                                                    v61 := v4["foo bar"]
                                                                                                                    var v62 uint32 = 630
                                                                                                                    _ = v60
                                                                                                                    _ = v61
                                                                                                                    _ = v62
                                                                                                                }
                                                                                                                var v63 string = "\n\t"
                                                                                                                _ = v63
                                                                                                            }
                                                                                                        }
                                                                                                        v64 := GenS5{F0: 2147483647, F1: "hello", F2: ("hello" + (("hello" + "") + ""))}
                                                                                                        if true {
                                                                                                            fmt.Println(v27)
                                                                                                            v56 += (v56 & v56)
                                                                                                            v55 = v55
                                                                                                            var v65 int = 597
                                                                                                            var v66 uint = 255
                                                                                                            v34 += (fi49 % (((fi49 ^ 4294967295)) | 1))
                                                                                                            v67 := []int{-976, (v65 % ((v34) | 1))}
                                                                                                            _ = v65
                                                                                                            _ = v66
                                                                                                            _ = v67
                                                                                                        } else {
                                                                                                            if (("\n\t" + "\n\t") > "") {
                                                                                                                switch 4294967295 {
                                                                                                                case 39:
                                                                                                                    fmt.Println(fi13, v56)
                                                                                                                default:
                                                                                                                    v68 := []int{fi49, v1}
                                                                                                                    _ = v68
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v69 float32 = float32(float64(0.1))
                                                                                                                var v70 uint64 = 4294967295
                                                                                                                if (((v34 % ((0) | 1)) <= 313) && ((("hello" == ("\n\t" + "")) || (float64(v69) <= 1e10)) && true)) {
                                                                                                                    var v71 int64 = 1
                                                                                                                    var v72 int64 = (v56 / ((int64(v55)) | 1))
                                                                                                                    var v73 float32 = 3.14159
                                                                                                                    switch "" {
                                                                                                                    case "\n\t":
                                                                                                                        if true {
                                                                                                                            var v74 int8 = 127
                                                                                                                            v75 := []int{(v27 / (((v1 / ((v1) | 1))) | 1))}
                                                                                                                            var v76 bool = (2 != 0)
                                                                                                                            _ = v74
                                                                                                                            _ = v75
                                                                                                                            _ = v76
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v71
                                                                                                                    _ = v72
                                                                                                                    _ = v73
                                                                                                                }
                                                                                                                _ = v69
                                                                                                                _ = v70
                                                                                                            }
                                                                                                            v1 -= fi49
                                                                                                        }
                                                                                                        _ = v64
                                                                                                    }
                                                                                                }
                                                                                                v77 := []int{}
                                                                                                _ = v55
                                                                                                _ = v56
                                                                                                _ = v57
                                                                                                _ = v77
                                                                                            }
                                                                                        }
                                                                                        v27 += (fi49 | (v34 * 0))
                                                                                        v33 -= (v33 / ((v33) | 1))
                                                                                    } else {
                                                                                        v33 = (v33 | 0)
                                                                                        if true {
                                                                                            v78 := GenS5{F0: 729, F1: ("hello" + ((("foo bar" + "") + ("hello" + (("" + "a") + (("foo bar" + "\n\t") + "hello")))) + "foo bar")), F2: ""}
                                                                                            v33 -= (v33 * (^(v33)))
                                                                                            _ = v78
                                                                                        } else {
                                                                                            fmt.Println(fi49, fi13)
                                                                                        }
                                                                                    }
                                                                                }
                                                                            default:
                                                                                fmt.Println(v33, fi49)
                                                                            }
                                                                            v28 = (0.1 / float64(0.1))
                                                                        } else {
                                                                            switch "" {
                                                                            case "hello":
                                                                                v79 := GenS6{F0: 100, F1: 118}
                                                                                _ = v79
                                                                            }
                                                                        }
                                                                    case "hello":
                                                                        var v80 float64 = v28
                                                                        _ = v80
                                                                    case "a":
                                                                        v81 := GenS1{F0: (^(v33)), F1: 42, F2: 2147483647}
                                                                        _ = v81
                                                                    }
                                                                    _ = fi49
                                                                }
                                                            } else {
                                                                var v82 int32 = -2147483648
                                                                _ = v82
                                                            }
                                                        }
                                                        _ = v34
                                                    } else {
                                                        var v83 float32 = 3.14159
                                                        v1 += 0
                                                        var v84 uint64 = 42
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                    v85 := []int{(v27 & v1)}
                                                    _ = v33
                                                    _ = v85
                                                }
                                            }
                                            v86 := GenS7{F0: 9223372036854775807, F1: 128, F2: -1}
                                            _ = v86
                                        }
                                        _ = v28
                                    } else {
                                        var v87 float32 = 0.1
                                        if true {
                                            if false {
                                                v88 := []int{(^(v1))}
                                                fmt.Println(v1, v1, v27)
                                                if (9223372036854775807 >= 1000) {
                                                    v89 := GenS4{F0: 734, F1: v1, F2: 891}
                                                    if false {
                                                        v90 := GenS8{F0: 2}
                                                        if false {
                                                            v1 = 9223372036854775807
                                                            var v91 int64 = 9223372036854775807
                                                            _ = v91
                                                        } else {
                                                            for fi92 := 0; fi92 < 18; fi92++ {
                                                                fmt.Println(fi92, v87, fi13)
                                                                var v93 float64 = 0.1
                                                                fmt.Println(fi13, v1, v93)
                                                                _ = fi92
                                                                _ = v93
                                                            }
                                                        }
                                                        v87 = float32(1e10)
                                                        _ = v90
                                                    }
                                                    _ = v89
                                                }
                                                _ = v88
                                            }
                                        } else {
                                            if true {
                                                v94 := []int{}
                                                if (2147483647 >= 2147483648) {
                                                    v95 := GenS8{F0: 65535}
                                                    v96 := GenS2{F0: -2147483648, F1: 0, F2: true}
                                                    v1 = 128
                                                    _ = v95
                                                    _ = v96
                                                } else {
                                                    for fi97 := 0; fi97 < 14; fi97++ {
                                                        v98 := GenS9{F0: 4294967295}
                                                        v87 -= v87
                                                        _ = fi97
                                                        _ = v98
                                                    }
                                                }
                                                _ = v94
                                            }
                                            if true {
                                                for fi99 := 0; fi99 < 1; fi99++ {
                                                    fmt.Println(v1, v1)
                                                    v1 -= fi13
                                                    _ = fi99
                                                }
                                            } else {
                                                switch "" {
                                                case "foo bar":
                                                    fmt.Println(v87, v1)
                                                case "\n\t":
                                                    var v100 int64 = 0
                                                    _ = v100
                                                case "":
                                                    if true {
                                                        if true {
                                                            v87 -= 3.14159
                                                            switch 9223372036854775807 {
                                                            case 127:
                                                                switch ("a" + "hello") {
                                                                case "a":
                                                                    v27 = 4294967295
                                                                default:
                                                                    for fi101 := 0; fi101 < 8; fi101++ {
                                                                        var v102 string = "\n\t"
                                                                        _ = fi101
                                                                        _ = v102
                                                                    }
                                                                }
                                                            case -11:
                                                                switch fi13 {
                                                                case 42:
                                                                    v1 += fi13
                                                                case -799:
                                                                    fmt.Println(v1, v27, v1)
                                                                default:
                                                                    var v103 uint = 2
                                                                    _ = v103
                                                                }
                                                            default:
                                                                v1 = fi13
                                                            }
                                                        } else {
                                                            var v104 uint32 = 65535
                                                            var v105 int = (fi13 << ((711) & 63))
                                                            switch 396 {
                                                            case 128:
                                                                v27 += (v27 & (fi13 >> ((979) & 63)))
                                                            case 32768:
                                                                fmt.Println(fi13)
                                                            }
                                                            _ = v104
                                                            _ = v105
                                                        }
                                                        switch (v27 >> ((0) & 63)) {
                                                        case 0:
                                                            switch 32768 {
                                                            case -2147483648:
                                                                var v106 string = "a"
                                                                _ = v106
                                                            default:
                                                                v107 := GenS9{F0: 2147483648}
                                                                _ = v107
                                                            }
                                                        case 128:
                                                            v108 := GenS10{F0: 42, F1: 32767, F2: 255}
                                                            _ = v108
                                                        case 32768:
                                                            v109 := GenS1{F0: 440, F1: 255, F2: 2147483647}
                                                            _ = v109
                                                        }
                                                        var v110 uint64 = 42
                                                        _ = v110
                                                    } else {
                                                        if (-572 >= -533) {
                                                            v111 := GenS7{F0: -979, F1: v1, F2: -196}
                                                            _ = v111
                                                        } else {
                                                            var v112 int64 = -486
                                                            _ = v112
                                                        }
                                                        v113 := GenS11{F0: 42}
                                                        fmt.Println(v1)
                                                        _ = v113
                                                    }
                                                }
                                                switch ("" + "foo bar") {
                                                case "\n\t":
                                                    v27 = v27
                                                default:
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        if ((!false) || false) {
                                                            v1 = fi13
                                                        }
                                                    case "\n\t":
                                                        fmt.Println(v87, v27, v87)
                                                    case "":
                                                        if (v1 <= v1) {
                                                            v87 = (v87 + (v87 - ((v87 / v87) + (v87 + 0.1))))
                                                        } else {
                                                            v1 += v1
                                                            if (!true) {
                                                                v114 := GenS10{F0: 0, F1: 1, F2: 634}
                                                                _ = v114
                                                            }
                                                            for fi115 := 0; fi115 < 13; fi115++ {
                                                                v116 := []int{(fi115 - (-(fi13))), -2147483648, (v1 - (fi13 ^ fi115))}
                                                                v87 -= 0.1
                                                                var v117 uint8 = 2
                                                                _ = fi115
                                                                _ = v116
                                                                _ = v117
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        for fi118 := 0; fi118 < 6; fi118++ {
                                            var v119 float32 = (v87 * v87)
                                            if (0 >= -18) {
                                                if (!(!true)) {
                                                    var v120 float32 = (v119 + float32((3.14159 * 1.5)))
                                                    fmt.Println(v87, fi13, fi13)
                                                    delete(v3, ("foo bar" + "\n\t"))
                                                    _ = v120
                                                } else {
                                                    fmt.Println(v87, v1, v119)
                                                    v1 = fi118
                                                }
                                                fmt.Println(fi13, v87, v27)
                                                v121 := GenS12{F0: 42}
                                                _ = v121
                                            }
                                            _ = fi118
                                            _ = v119
                                        }
                                        _ = v87
                                    }
                                case -1:
                                    fmt.Println(v1, fi13)
                                }
                            case -1:
                                for fi122 := 0; fi122 < 14; fi122++ {
                                    v27 = (^(v1))
                                    _ = fi122
                                }
                            }
                            fmt.Println(v27, fi13)
                            _ = v27
                        }
                    } else {
                        for fi123 := 0; fi123 < 14; fi123++ {
                            switch "a" {
                            case "foo bar":
                                v124 := GenS13{F0: 297, F1: (v1 % (((fi123 / ((v1) | 1))) | 1))}
                                _ = v124
                            }
                            v1 += (v1 ^ (fi123 + (fi123 << ((0) & 63))))
                            v1 += fi13
                            _ = fi123
                        }
                        switch "a" {
                        case "\n\t":
                            var v125 uint8 = 2
                            _ = v125
                        case "":
                            for fi126 := 0; fi126 < 2; fi126++ {
                                if (!false) {
                                    fmt.Println(v1)
                                } else {
                                    fmt.Println(fi126, fi13)
                                    if true {
                                        if false {
                                            switch fi13 {
                                            case 703:
                                                v127 := []int{762, 2147483647, fi13, fi13}
                                                _ = v127
                                            case 127:
                                                var v128 float32 = (1e10 * 0.0)
                                                _ = v128
                                            case 0:
                                                if true {
                                                    switch ("foo bar" + "foo bar") {
                                                    case "foo bar":
                                                        fmt.Println(v1, fi126)
                                                    case "hello":
                                                        v1 += v1
                                                    }
                                                    if true {
                                                        var v129 uint = 255
                                                        var v130 string = ("\n\t" + "foo bar")
                                                        _ = v129
                                                        _ = v130
                                                    } else {
                                                        if (!true) {
                                                            v131 := GenS14{F0: 868, F1: (0 >= 0), F2: 1000}
                                                            v132 := GenS5{F0: 194, F1: "hello", F2: (("a" + "") + "\n\t")}
                                                            _ = v131
                                                            _ = v132
                                                        }
                                                    }
                                                } else {
                                                    v133 := GenS11{F0: 183}
                                                    v134 := []int{-9223372036854775808, (fi13 ^ 128), -21, (fi13 & fi13)}
                                                    _ = v133
                                                    _ = v134
                                                }
                                            default:
                                                var v135 float32 = (((1e10 + 0.0) / 1e10) - (1e10 / 1.5))
                                                _ = v135
                                            }
                                        }
                                        if true {
                                            fmt.Println(fi126, v1, fi13)
                                            switch ("a" + ("\n\t" + ("hello" + ""))) {
                                            case "foo bar":
                                                if false {
                                                    v1 = (v1 ^ (-(fi13)))
                                                }
                                            default:
                                                var v136 int64 = 763
                                                _ = v136
                                            }
                                            if false {
                                                _ = len(v6)
                                                fmt.Println(v1, v1)
                                            } else {
                                                if (!true) {
                                                    v137 := GenS7{F0: 9223372036854775807, F1: v1, F2: 1000}
                                                    var v138 uint = 0
                                                    _ = v137
                                                    _ = v138
                                                }
                                                v139 := []int{}
                                                v1 -= v1
                                                if true {
                                                    if true {
                                                        fmt.Println(v1, fi13, v1)
                                                        var v140 int64 = 1
                                                        v1 = fi126
                                                        var v141 int64 = (v140 % ((v140) | 1))
                                                        _ = v140
                                                        _ = v141
                                                    } else {
                                                        switch v1 {
                                                        case 9223372036854775807:
                                                            v1 -= fi13
                                                        case 2147483647:
                                                            if len(v139) >= 2 {
                                                                v139 = v139[1:len(v139)]
                                                            }
                                                        default:
                                                            v142 := GenS15{F0: 566}
                                                            _ = v142
                                                        }
                                                    }
                                                    v143 := GenS16{F0: 1, F1: 0, F2: -1}
                                                    _ = v143
                                                } else {
                                                    v1 += 42
                                                    v144 := GenS17{F0: (14 > -1), F1: -100}
                                                    _ = v144
                                                }
                                                var v145 int16 = 1
                                                _ = v139
                                                _ = v145
                                            }
                                        }
                                    } else {
                                        v146 := []int{fi13, -9223372036854775808}
                                        var v147 string = ""
                                        _ = v146
                                        _ = v147
                                    }
                                    v148 := GenS6{F0: 100, F1: 0}
                                    var v149 uint8 = 255
                                    _ = v148
                                    _ = v149
                                }
                                var v150 uint = 18446744073709551615
                                _ = fi126
                                _ = v150
                            }
                        case "a":
                            for fi151 := 0; fi151 < 9; fi151++ {
                                switch (fi151 / (((fi13 & (v1 - -2147483648))) | 1)) {
                                case 32767:
                                    switch "a" {
                                    case "foo bar":
                                        var v152 uint8 = 255
                                        _ = v152
                                    }
                                case -9223372036854775808:
                                    v153 := v6[42]
                                    _ = v153
                                case 0:
                                    if (32767 >= -32768) {
                                        switch (fi151 & 227) {
                                        case -610:
                                            if ("foo bar" >= "a") {
                                                v1 -= -109
                                                v154 := GenS18{F0: 0, F1: 128, F2: false}
                                                var v155 uint8 = 255
                                                _ = v154
                                                _ = v155
                                            }
                                        case 32767:
                                            v1 -= 736
                                        case 226:
                                            if (42 <= 65) {
                                                v156 := []int{fi13, fi151, -2147483648}
                                                _ = v156
                                            }
                                            var v157 bool = false
                                            v158 := GenS19{F0: 916, F1: 333, F2: 4294967295}
                                            v1 = -9223372036854775808
                                            _ = v157
                                            _ = v158
                                        default:
                                            if (true || true) {
                                                v1 -= v1
                                                var v159 uint32 = 1
                                                v160 := GenS11{F0: -1}
                                                _ = v159
                                                _ = v160
                                            } else {
                                                switch "a" {
                                                case "a":
                                                    delete(v3, "foo bar")
                                                case "foo bar":
                                                    fmt.Println(fi151, v1, fi151)
                                                case "\n\t":
                                                    switch "hello" {
                                                    case "a":
                                                        var v161 bool = (0 != 2147483647)
                                                        _ = v161
                                                    default:
                                                        if (1 == 0) {
                                                            var v162 string = "hello"
                                                            _ = v162
                                                        }
                                                    }
                                                    var v163 uint = 2
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v164 := GenS5{F0: 303, F1: "foo bar", F2: "\n\t"}
                                                        _ = v164
                                                    case "foo bar":
                                                        _ = len(v2)
                                                    default:
                                                        fmt.Println(fi151, fi151, v1)
                                                    }
                                                    _ = v163
                                                default:
                                                    if false {
                                                        var v165 int8 = -16
                                                        if ((float32(1e10) != 0.0) && false) {
                                                            var v166 int16 = 0
                                                            if (0.0 <= 3.14159) {
                                                                v1 += (fi13 / ((v1) | 1))
                                                                v165 = v165
                                                            }
                                                            v1 = (v1 >> ((255) & 63))
                                                            var v167 float32 = ((3.14159 * (-2.25 - 1.5)) + 0.0)
                                                            switch -9223372036854775808 {
                                                            case 755:
                                                                v168 := []int{0}
                                                                _ = v168
                                                            case 0:
                                                                v169 := GenS20{F0: 778}
                                                                _ = v169
                                                            default:
                                                                v166 += v166
                                                            }
                                                            fmt.Println(fi13, v1, v165)
                                                            v167 += float32(1e10)
                                                            _ = v166
                                                            _ = v167
                                                        }
                                                        v170 := GenS15{F0: 0}
                                                        _ = v165
                                                        _ = v170
                                                    }
                                                }
                                            }
                                        }
                                        if (0 < 1) {
                                            var v171 uint64 = 76
                                            v172 := GenS21{F0: (!false), F1: 127, F2: (fi151 + fi13)}
                                            v1 = fi13
                                            _ = v171
                                            _ = v172
                                        } else {
                                            v1 -= (fi151 & (^(fi151)))
                                            v1 -= (fi13 % (((fi13 / (((v1 & fi13)) | 1))) | 1))
                                            v1 = fi13
                                        }
                                        if (0 >= 847) {
                                            v173 := []int{(v1 + (v1 + (v1 + 128))), (v1 % ((2147483647) | 1)), 32768, fi151}
                                            _ = v173
                                        } else {
                                            if (true && true) {
                                                v174 := GenS6{F0: 42, F1: 11}
                                                var v175 uint64 = 77
                                                _ = v174
                                                _ = v175
                                            } else {
                                                v176 := GenS9{F0: 845}
                                                v177 := GenS10{F0: -119, F1: 0, F2: 42}
                                                _ = v176
                                                _ = v177
                                            }
                                            if true {
                                                v1 = fi13
                                                v1 = (fi151 / (((fi13 >> ((4294967295) & 63))) | 1))
                                                v178 := GenS22{F0: 2147483647, F1: 403}
                                                var v179 uint8 = 17
                                                v180 := GenS23{F0: 832}
                                                _ = v178
                                                _ = v179
                                                _ = v180
                                            }
                                        }
                                        v181 := []int{v1, 127}
                                        _ = v181
                                    }
                                }
                                var v182 bool = false
                                if (127 >= -96) {
                                    fmt.Println(fi13, fi151)
                                    var v183 float64 = (0.0 - 0.1)
                                    switch "hello" {
                                    case "\n\t":
                                        var v184 uint32 = 310
                                        _ = v184
                                    case "foo bar":
                                        v183 -= (v183 + v183)
                                    default:
                                        v185 := v2[v1]
                                        _ = v185
                                    }
                                    _ = v183
                                } else {
                                    var v186 uint = 0
                                    v1 += (v1 - fi151)
                                    _ = v186
                                }
                                _ = fi151
                                _ = v182
                            }
                        default:
                            v1 = v1
                        }
                    }
                default:
                    var v187 int16 = 1000
                    _ = v187
                }
                fmt.Println(fi13, v1)
            } else {
                if (!true) {
                    _ = len(v4)
                    v188 := GenS19{F0: 360, F1: 945, F2: genFn0(42, 56, 0)}
                    v189 := GenS18{F0: 0, F1: 0, F2: (!(0 != -736))}
                    _ = v188
                    _ = v189
                } else {
                    switch ("hello" + "foo bar") {
                    case "foo bar":
                        fmt.Println(fi13, v1)
                    case "":
                        _ = len(v2)
                    }
                }
                fmt.Println(v1, fi13)
                fmt.Println(v1)
            }
        }
        if true {
            v190 := GenS18{F0: 255, F1: 0, F2: false}
            for fi191 := 0; fi191 < 15; fi191++ {
                if true {
                    var v192 uint = 0
                    var v193 float32 = 1e10
                    _ = v192
                    _ = v193
                } else {
                    v194 := GenS24{F0: 939, F1: 185, F2: fi191}
                    switch 0 {
                    case -38:
                        switch "\n\t" {
                        case "a":
                            v195 := []int{(fi13 % ((v1) | 1)), -2147483648, (v1 - v1)}
                            _ = v195
                        case "hello":
                            v1 -= (v1 + (v1 << ((397) & 63)))
                        }
                    default:
                        v196 := []int{fi13, -714, v1}
                        _ = v196
                    }
                    _ = v194
                }
                _ = fi191
            }
            for fi197 := 0; fi197 < 18; fi197++ {
                v198 := GenS25{F0: 1, F1: -308, F2: false}
                var v199 uint64 = 233
                _ = fi197
                _ = v198
                _ = v199
            }
            _ = v190
        } else {
            var v200 uint8 = 1
            _ = v200
        }
        v201 := []int{fi13}
        _ = fi13
        _ = v201
    }
    var v202 int64 = 1
    switch "\n\t" {
    case "\n\t":
        var v203 float32 = 1e10
        _ = v203
    case "":
        var v204 uint64 = genFn0(10, -128, 42)
        _ = v204
    }
    v205 := GenS26{F0: ("hello" + "foo bar"), F1: (v202 - -1)}
    var v206 int64 = v202
    switch "\n\t" {
    case "hello":
        var v207 string = ""
        _ = v207
    case "foo bar":
        v206 += v206
    case "a":
        var v208 float64 = (1.5 * 1.5)
        _ = v208
    default:
        v209 := GenS8{F0: 630}
        _ = v209
    }
    v210 := GenS6{F0: -126, F1: 231}
    delete(v5, (("a" + "a") + "hello"))
    if true {
        var v211 int16 = 1
        switch v1 {
        case -1:
            var v212 float64 = 3.14159
            _ = v212
        case 127:
            v213 := []int{-528}
            _ = v213
        default:
            var v214 int8 = -1
            _ = v214
        }
        var v215 float64 = 1.5
        _ = v211
        _ = v215
    }
    fmt.Println(v206, v206, v202)
    var v216 float64 = 1.5
    v206 = v202
    switch "hello" {
    case "a":
        if (127 != 102) {
            v202 -= (v206 * v202)
            var v217 float64 = 0.1
            for fi218 := 0; fi218 < 5; fi218++ {
                var v219 string = "a"
                var v220 float64 = v216
                _ = fi218
                _ = v219
                _ = v220
            }
            _ = v217
        }
    default:
        if (!(255 > 42)) {
            v221 := []int{}
            var v222 bool = false
            v206 = 712
            _ = v221
            _ = v222
        }
    }
    v223 := []int{}
    v202 -= v202
    fmt.Println(v206, v206)
    v224 := v4["foo bar"]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v202:", v202)
    fmt.Println("v206:", v206)
    fmt.Println("v216:", v216)
    fmt.Println("slicelen v223:", len(v223), v223)
    fmt.Println("v224:", v224)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v202
    _ = v205
    _ = v206
    _ = v210
    _ = v216
    _ = v223
    _ = v224
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
