package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint64 }

func main() {
    var v1 int16 = 61
    var v2 int64 = 0
    v3 := make(map[string]int)
    var v4 int16 = v1
    var v5 uint32 = 628
    v6 := v3[""]
    switch "a" {
    case "":
        for fi7 := 0; fi7 < 20; fi7++ {
            v6 = 128
            _ = fi7
        }
    case "\n\t":
        v6 = (v6 >> ((65535) & 63))
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        if true {
            v5 -= v5
            var v9 uint32 = (^(v5))
            v5 += 65535
            _ = v9
        } else {
            fmt.Println(v1, v4, v2)
            fmt.Println(v4)
        }
        _ = fi8
    }
    var v10 uint = 957
    v5 -= (v5 + (v5 >> (((v10 + (^(v10)))) & 63)))
    switch "hello" {
    case "":
        switch 0 {
        case 0:
            fmt.Println(v6, v10)
        default:
            v2 += v2
        }
    case "a":
        switch (v6 | v6) {
        case 42:
            if true {
                switch v6 {
                case -9223372036854775808:
                    var v11 int = v6
                    _ = v11
                default:
                    var v12 float64 = 1.5
                    _ = v12
                }
                var v13 int32 = int32(v1)
                if true {
                    var v14 uint32 = v5
                    var v15 int32 = v13
                    v4 += v4
                    _ = v14
                    _ = v15
                } else {
                    switch -9223372036854775808 {
                    case 577:
                        v2 += -1
                    }
                    switch v6 {
                    case -963:
                        if (uint64(v5) == 1) {
                            if (!(0 > 1)) {
                                for fi16 := 0; fi16 < 9; fi16++ {
                                    if false {
                                        for fi17 := 0; fi17 < 15; fi17++ {
                                            var v18 float64 = -2.25
                                            if true {
                                                switch (v6 | (v6 | fi17)) {
                                                case -527:
                                                    v19 := GenS1{F0: (v13 & v13), F1: (v6 * fi17)}
                                                    _ = v19
                                                case 1:
                                                    v20 := GenS2{F0: 1, F1: 0}
                                                    _ = v20
                                                case -2147483648:
                                                    v1 += int16(v13)
                                                default:
                                                    v13 -= v13
                                                }
                                                v21 := []int{0, (fi16 | (fi17 >> ((v10) & 63))), (v6 % (((fi17 % ((288) | 1))) | 1))}
                                                fmt.Println(v5)
                                                _ = v21
                                            } else {
                                                var v22 uint64 = 4294967295
                                                if false {
                                                    fmt.Println(fi16, v5, v6)
                                                } else {
                                                    if true {
                                                        v22 = 1
                                                        if (!(2 == 42)) {
                                                            fmt.Println(v22)
                                                        }
                                                    } else {
                                                        fmt.Println(v6)
                                                        if (true && ((v2 >> ((v10) & 63)) >= (v2 * 230))) {
                                                            var v23 int8 = 0
                                                            v1 -= (v4 + (-(v1)))
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                fmt.Println(v18, v6, v23)
                                                            default:
                                                                v5 = (v5 * v5)
                                                            }
                                                            _ = v23
                                                        } else {
                                                            var v24 uint8 = 128
                                                            _ = v24
                                                        }
                                                        v18 += ((((float64(-2.25) - v18) - (v18 + (3.14159 / (float64(0.1) + v18)))) / 1.0) * 1.5)
                                                    }
                                                    v25 := v3[("\n\t" + "hello")]
                                                    _ = v25
                                                }
                                                if true {
                                                    if false {
                                                        if (42 < 900) {
                                                            v26 := GenS1{F0: v13, F1: (fi17 % (((-(v6))) | 1))}
                                                            _ = v26
                                                        } else {
                                                            if false {
                                                                fmt.Println(v4)
                                                            }
                                                        }
                                                    }
                                                    if (1 < 42) {
                                                        v4 = v4
                                                        v2 = (v2 * (v2 + int64(v5)))
                                                        v27 := GenS2{F0: 65535, F1: (v22 - (v22 / (((v22 % (((^(v22))) | 1))) | 1)))}
                                                        _ = v27
                                                    } else {
                                                        v28 := []int{32767, -2147483648}
                                                        v29 := GenS2{F0: 255, F1: (v22 - (v22 & v22))}
                                                        if (true || (!false)) {
                                                            var v30 float32 = float32(v18)
                                                            _ = v30
                                                        } else {
                                                            switch -698 {
                                                            case -24:
                                                                v31 := GenS2{F0: 205, F1: v22}
                                                                _ = v31
                                                            }
                                                        }
                                                        fmt.Println(v4)
                                                        _ = v28
                                                        _ = v29
                                                    }
                                                } else {
                                                    v32 := []int{}
                                                    var v33 int32 = 42
                                                    _ = v32
                                                    _ = v33
                                                }
                                                _ = v22
                                            }
                                            fmt.Println(v5, v10, fi17)
                                            _ = fi17
                                            _ = v18
                                        }
                                        switch (fi16 << ((v10) & 63)) {
                                        case 9223372036854775807:
                                            var v34 uint = (v10 ^ v10)
                                            _ = v34
                                        case 42:
                                            for fi35 := 0; fi35 < 18; fi35++ {
                                                fmt.Println(v6, v10, fi16)
                                                fmt.Println(v1, v5, v10)
                                                _ = fi35
                                            }
                                        default:
                                            var v36 int64 = v2
                                            _ = v36
                                        }
                                        var v37 bool = (("hello" + ("hello" + ("" + ("a" + "")))) == "")
                                        _ = v37
                                    } else {
                                        for fi38 := 0; fi38 < 8; fi38++ {
                                            fmt.Println(v6)
                                            var v39 bool = (int8(v13) == -1)
                                            if ((true || ((v10 << (((v10 % (((v10 ^ (v10 & (v10 >> (((v10 | v10)) & 63))))) | 1))) & 63)) < (^(v10)))) && v39) {
                                                v5 = (v5 - (v5 ^ (^(v5))))
                                                var v40 uint16 = 130
                                                var v41 float32 = 3.14159
                                                _ = v40
                                                _ = v41
                                            } else {
                                                v39 = v39
                                                var v42 string = "hello"
                                                _ = v42
                                            }
                                            var v43 int64 = v2
                                            if v39 {
                                                if true {
                                                    fmt.Println(v1)
                                                }
                                            }
                                            fmt.Println(v1, v2)
                                            v44 := GenS1{F0: -2147483648, F1: -2147483648}
                                            v45 := v3[(("" + "foo bar") + "")]
                                            _ = fi38
                                            _ = v39
                                            _ = v43
                                            _ = v44
                                            _ = v45
                                        }
                                    }
                                    var v46 int32 = (v13 * v13)
                                    _ = fi16
                                    _ = v46
                                }
                                v13 = (^(v13))
                                v47 := []int{(v6 | (v6 | v6)), (^(v6)), v6}
                                _ = v47
                            }
                            var v48 int16 = (v4 >> ((255) & 63))
                            _ = v48
                        }
                    case 127:
                        v49 := GenS3{F0: 0, F1: 315, F2: 4294967295}
                        _ = v49
                    case -203:
                        fmt.Println(v2, v4)
                    }
                    for fi50 := 0; fi50 < 12; fi50++ {
                        v6 -= fi50
                        _ = fi50
                    }
                }
                _ = v13
            } else {
                if (((("\n\t" + ("\n\t" + ("\n\t" + "a"))) + ("hello" + ("a" + ("" + "hello")))) + ("\n\t" + ("\n\t" + "\n\t"))) != ("hello" + "hello")) {
                    v2 = v2
                    switch v6 {
                    case -186:
                        if (v1 == -1) {
                            v3["a"] = v6
                            v1 = v4
                        } else {
                            v4 = 0
                            if false {
                                fmt.Println(v4, v5)
                            }
                            switch (v6 + v6) {
                            case 128:
                                v5 = 4294967295
                            case -2147483648:
                                v51 := GenS2{F0: 0, F1: 42}
                                _ = v51
                            case 9223372036854775807:
                                fmt.Println(v5, v4)
                            }
                        }
                    case 32768:
                        var v52 bool = false
                        _ = v52
                    default:
                        for fi53 := 0; fi53 < 4; fi53++ {
                            v5 += (v5 * v5)
                            switch ((("\n\t" + "\n\t") + "\n\t") + ((("a" + "\n\t") + "a") + "hello")) {
                            case "foo bar":
                                switch 253 {
                                case 2147483647:
                                    v54 := []int{(v6 % (((v6 >> ((18446744073709551615) & 63))) | 1)), fi53}
                                    _ = v54
                                }
                            case "\n\t":
                                v55 := GenS4{F0: v1, F1: (("" + ("a" + "a")) + "a")}
                                _ = v55
                            case "hello":
                                if (!true) {
                                    v10 = v10
                                    for fi56 := 0; fi56 < 7; fi56++ {
                                        v57 := GenS3{F0: (v1 * int16(v5)), F1: 42, F2: 42}
                                        v58 := []int{(v6 / (((fi53 ^ 9223372036854775807)) | 1)), (v6 << ((2) & 63))}
                                        if false {
                                            v3["\n\t"] = 9223372036854775807
                                            var v59 int64 = (v2 << ((18446744073709551615) & 63))
                                            switch fi56 {
                                            case 32768:
                                                v6 -= (v6 | (v6 * (fi53 << ((v10) & 63))))
                                                var v60 string = ""
                                                v59 = -111
                                                _ = v60
                                            case 42:
                                                if true {
                                                    switch ("a" + ("" + "a")) {
                                                    case "hello":
                                                        fmt.Println(v6, v2, v5)
                                                    case "\n\t":
                                                        v6 += (v6 | v6)
                                                    }
                                                    var v61 float64 = -2.25
                                                    v62 := GenS2{F0: 2, F1: 416}
                                                    fmt.Println(v6, fi53, v4)
                                                    _ = v61
                                                    _ = v62
                                                }
                                            default:
                                                _ = len(v58)
                                            }
                                            _ = v59
                                        } else {
                                            v4 -= -32768
                                            if (true || true) {
                                                if len(v58) > 0 {
                                                    v58[0] = -2147483648
                                                }
                                                switch "" {
                                                case "a":
                                                    switch ("hello" + "\n\t") {
                                                    case "hello":
                                                        v63 := GenS3{F0: (^(v1)), F1: 0, F2: 372}
                                                        _ = v63
                                                    case "foo bar":
                                                        v5 = (^(v5))
                                                    default:
                                                        switch "hello" {
                                                        case "a":
                                                            if (!false) {
                                                                var v64 uint8 = 42
                                                                if (-923 < v4) {
                                                                    v6 -= (^(fi56))
                                                                    v6 += (fi53 / ((128) | 1))
                                                                    if false {
                                                                        if true {
                                                                            v2 = (v2 & v2)
                                                                            if (!(-2.25 > 1e10)) {
                                                                                v65 := GenS2{F0: 32767, F1: 1}
                                                                                v1 = (v1 * (v4 << ((v10) & 63)))
                                                                                _ = v65
                                                                            } else {
                                                                                var v66 float32 = float32((-2.25 / (((0.0 / 1.0) + 1.5) + (float64(3.14159) + (0.1 - 1e10)))))
                                                                                var v67 float64 = (((1e10 / 1e10) * ((((0.1 - 1.5) / 1.5) + 3.14159) + (((float64(v66) + 0.1) / (0.1 * 3.14159)) * -2.25))) / float64((v66 / (v66 - v66))))
                                                                                _ = v66
                                                                                _ = v67
                                                                            }
                                                                            if true {
                                                                                delete(v3, "a")
                                                                            } else {
                                                                                v68 := GenS2{F0: 255, F1: 42}
                                                                                if (v10 >= (^(v10))) {
                                                                                    v3[(("" + "hello") + ("a" + (("foo bar" + (("foo bar" + "") + ((("hello" + "hello") + "foo bar") + "hello"))) + "hello")))] = v6
                                                                                    if len(v58) >= 2 {
                                                                                        v58 = v58[1:len(v58)]
                                                                                    }
                                                                                    var v69 int = (v6 << ((v10) & 63))
                                                                                    _ = v69
                                                                                } else {
                                                                                    var v70 int8 = -128
                                                                                    _ = v70
                                                                                }
                                                                                _ = v68
                                                                            }
                                                                        } else {
                                                                            v71 := v3["hello"]
                                                                            _ = v71
                                                                        }
                                                                        fmt.Println(v64, fi53, v64)
                                                                    }
                                                                    v2 += v2
                                                                    v2 = 1
                                                                    if (621 == 1) {
                                                                        delete(v3, "foo bar")
                                                                        v10 += (^(v10))
                                                                    }
                                                                    if ((^(v5)) <= v5) {
                                                                        if (-1 >= -385) {
                                                                            fmt.Println(v1)
                                                                            var v72 float64 = 1e10
                                                                            v5 += v5
                                                                            v73 := v3[("\n\t" + ("\n\t" + ((("foo bar" + "a") + "") + "a")))]
                                                                            _ = len(v58)
                                                                            v74 := GenS5{F0: v2}
                                                                            _ = v72
                                                                            _ = v73
                                                                            _ = v74
                                                                        } else {
                                                                            v64 -= 42
                                                                        }
                                                                    } else {
                                                                        var v75 string = ""
                                                                        _ = v75
                                                                    }
                                                                    fmt.Println(v64, fi56, v10)
                                                                } else {
                                                                    v5 = (v5 | uint32(v64))
                                                                    v4 += 383
                                                                    var v76 bool = true
                                                                    _ = v76
                                                                }
                                                                _ = v64
                                                            }
                                                        case "foo bar":
                                                            if true {
                                                                v5 -= (^(v5))
                                                                switch "\n\t" {
                                                                case "a":
                                                                    var v77 float64 = 1e10
                                                                    _ = v77
                                                                case "foo bar":
                                                                    var v78 int32 = -384
                                                                    _ = v78
                                                                case "\n\t":
                                                                    if true {
                                                                        v10 = 1
                                                                        if (!(!false)) {
                                                                            v79 := GenS4{F0: (v4 | v1), F1: ("\n\t" + ("\n\t" + ("foo bar" + ("hello" + "foo bar"))))}
                                                                            v80 := GenS6{F0: -70, F1: v2}
                                                                            if false {
                                                                                v2 += (-(v2))
                                                                            }
                                                                            switch "hello" {
                                                                            case "a":
                                                                                var v81 string = "hello"
                                                                                _ = v81
                                                                            }
                                                                            v82 := GenS7{F0: false, F1: v5, F2: v5}
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                v6 += 2147483647
                                                                            default:
                                                                                v83 := GenS8{F0: (v5 + v5)}
                                                                                _ = v83
                                                                            }
                                                                            _ = v79
                                                                            _ = v80
                                                                            _ = v82
                                                                        } else {
                                                                            v84 := GenS1{F0: -762, F1: fi53}
                                                                            fmt.Println(v1, v2, v5)
                                                                            var v85 uint = (^(v10))
                                                                            _ = v84
                                                                            _ = v85
                                                                        }
                                                                    }
                                                                default:
                                                                    fmt.Println(v1, v2, fi56)
                                                                }
                                                            } else {
                                                                v58 = append(v58, fi53)
                                                            }
                                                        case "":
                                                            var v86 int32 = 368
                                                            if len(v58) >= 2 {
                                                                v58 = v58[1:len(v58)]
                                                            }
                                                            _ = v86
                                                        }
                                                        switch (-(fi56)) {
                                                        case 127:
                                                            v1 -= -32768
                                                        default:
                                                            var v87 int = 1
                                                            _ = v87
                                                        }
                                                        v1 += v4
                                                    }
                                                default:
                                                    if (false || ("a" == "")) {
                                                        fmt.Println(fi56, fi56, fi56)
                                                        v2 = (v2 - v2)
                                                        v10 = v10
                                                    }
                                                }
                                            } else {
                                                v10 = (v10 >> (((v10 >> ((v10) & 63))) & 63))
                                            }
                                            var v88 uint = (v10 & 2)
                                            _ = v88
                                        }
                                        _ = fi56
                                        _ = v57
                                        _ = v58
                                    }
                                } else {
                                    if (-1 > 373) {
                                        for fi89 := 0; fi89 < 11; fi89++ {
                                            var v90 int32 = 797
                                            v3["\n\t"] = fi53
                                            v4 = 614
                                            _ = fi89
                                            _ = v90
                                        }
                                        var v91 string = (((("" + "a") + ("\n\t" + "")) + "") + "")
                                        v92 := []int{}
                                        _ = v91
                                        _ = v92
                                    }
                                }
                            default:
                                _ = len(v3)
                            }
                            var v93 int64 = v2
                            _ = fi53
                            _ = v93
                        }
                    }
                }
                switch ("hello" + "") {
                case "hello":
                    var v94 string = "foo bar"
                    _ = v94
                case "foo bar":
                    var v95 float32 = 1e10
                    _ = v95
                }
                fmt.Println(v1, v2, v6)
            }
        }
    default:
        var v96 uint32 = (v5 / ((v5) | 1))
        _ = v96
    }
    v97 := GenS9{F0: int32(v5)}
    fmt.Println(v1, v10, v4)
    v98 := GenS10{F0: 755}
    switch ((("a" + ("a" + "hello")) + "foo bar") + "\n\t") {
    case "\n\t":
        if (!(true && false)) {
            var v99 uint = 255
            _ = v99
        }
    case "hello":
        if (!(v5 <= 2147483648)) {
            v10 = 75
            var v100 int32 = 1
            _ = v100
        } else {
            for fi101 := 0; fi101 < 19; fi101++ {
                for fi102 := 0; fi102 < 2; fi102++ {
                    var v103 int64 = (v2 - 9223372036854775807)
                    _ = fi102
                    _ = v103
                }
                var v104 int = fi101
                for fi105 := 0; fi105 < 7; fi105++ {
                    v10 += (v10 >> (((^(v10))) & 63))
                    var v106 string = "\n\t"
                    v107 := []int{}
                    _ = fi105
                    _ = v106
                    _ = v107
                }
                _ = fi101
                _ = v104
            }
            _ = len(v3)
        }
    default:
        v108 := []int{v6, v6}
        _ = v108
    }
    v6 += v6
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v97
    _ = v98
}
