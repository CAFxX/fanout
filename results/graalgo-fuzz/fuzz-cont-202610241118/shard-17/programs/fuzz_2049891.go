package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint16; F2 string }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8; F1 int8; F2 int32 }
type GenS9 struct { F0 int16; F1 int32 }
type GenS10 struct { F0 string; F1 int32; F2 uint64 }
func genFn0() int32 {
    switch (("a" + "foo bar") + "\n\t") {
    case "\n\t":
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 int = fi6
            _ = fi6
            _ = v7
        }
    case "":
        var v8 uint = 255
        _ = v8
    case "foo bar":
        var v9 int16 = 857
        _ = v9
    }
    var v10 float64 = ((1.5 + (0.0 + ((((1.5 - (0.1 * 0.1)) + ((1e10 + -2.25) * -2.25)) * 0.0) + float64(0.1)))) / 1e10)
    switch 127 {
    case -368:
        var v11 uint = 0
        _ = v11
    case 32768:
        var v12 uint64 = 289
        _ = v12
    }
    v10 = v10
    _ = v10
    return 42
}

func main() {
    var v1 uint = 4294967295
    var v2 int32 = 1
    v3 := make(map[string]int)
    var v4 uint16 = 0
    var v5 uint64 = 42
    var v13 bool = ("foo bar" != "\n\t")
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := GenS1{F0: 78}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 8; fi17++ {
            switch ("" + "hello") {
            case "hello":
                var v18 float32 = ((0.0 - (0.0 - 1e10)) - 0.1)
                _ = v18
            case "\n\t":
                v1 = v1
            case "":
                delete(v3, "hello")
            default:
                if false {
                    if true {
                        v4 = v4
                        if true {
                            v19 := GenS1{F0: -45}
                            _ = v19
                        } else {
                            v5 -= v5
                            switch "hello" {
                            case "foo bar":
                                switch "" {
                                case "":
                                    var v20 int = fi17
                                    _ = v20
                                }
                            case "":
                                var v21 float32 = ((1e10 + (1.5 + (3.14159 + 0.1))) - 1.5)
                                _ = v21
                            }
                            if (-128 >= 1) {
                                if (!true) {
                                    fmt.Println(v13)
                                    v5 = v5
                                    var v22 int16 = 42
                                    _ = v22
                                }
                                v23 := GenS2{F0: 128, F1: v5}
                                v24 := GenS2{F0: fi17, F1: 4294967295}
                                fmt.Println(fi17)
                                _ = v23
                                _ = v24
                            }
                        }
                        v1 = (v1 | (v1 + 1))
                    }
                    v25 := GenS3{F0: 0, F1: (^(v4)), F2: ((("foo bar" + "\n\t") + "a") + "")}
                    _ = v25
                }
                var v26 int = fi16
                fmt.Println(fi16, v26)
                var v27 float32 = 0.0
                v28 := GenS2{F0: (fi17 * fi17), F1: 4294967295}
                _ = v26
                _ = v27
                _ = v28
            }
            v4 = v4
            v29 := GenS3{F0: 0, F1: (v4 - (v4 >> ((uint(v4)) & 63))), F2: ("a" + "hello")}
            _ = fi17
            _ = v29
        }
        v30 := GenS2{F0: 1, F1: 4294967295}
        _ = fi16
        _ = v30
    }
    var v31 int8 = 100
    if (0 >= 594) {
        v4 = v4
        if false {
            v13 = (127 != 32767)
            v32 := []int{9223372036854775807, -1}
            _ = v32
        }
        var v33 int64 = int64(v5)
        _ = v33
    }
    if v13 {
        for fi34 := 0; fi34 < 13; fi34++ {
            for fi35 := 0; fi35 < 19; fi35++ {
                switch ("" + "\n\t") {
                case "hello":
                    if false {
                        switch "\n\t" {
                        case "":
                            v36 := GenS3{F0: 2, F1: (v4 * (v4 & (^(v4)))), F2: "\n\t"}
                            _ = v36
                        default:
                            v37 := GenS4{F0: ("foo bar" + ("hello" + "\n\t")), F1: 1}
                            _ = v37
                        }
                        if (!(-1 > -32768)) {
                            fmt.Println(v1, v1, v1)
                            fmt.Println(v4, v31)
                        } else {
                            switch (fi35 >> ((65535) & 63)) {
                            case -41:
                                v13 = v13
                            case 32767:
                                var v38 uint = (v1 ^ (^(v1)))
                                _ = v38
                            case -2147483648:
                                delete(v3, ((("\n\t" + "hello") + "\n\t") + ""))
                            }
                            if ((v1 << (((v1 << ((1) & 63))) & 63)) >= 843) {
                                if false {
                                    if (v31 < int8(v1)) {
                                        if v13 {
                                            var v39 string = ""
                                            switch (v39 + (v39 + "a")) {
                                            case "foo bar":
                                                if (genFn0() < (v2 >> ((v1) & 63))) {
                                                    v5 = v5
                                                }
                                            case "\n\t":
                                                if true {
                                                    if (!v13) {
                                                        var v40 int8 = -1
                                                        if false {
                                                            switch "\n\t" {
                                                            case "a":
                                                                v41 := GenS3{F0: 1, F1: (v4 ^ (^(v4))), F2: (v39 + "")}
                                                                _ = v41
                                                            default:
                                                                switch v39 {
                                                                case "":
                                                                    v39 = "\n\t"
                                                                case "a":
                                                                    var v42 uint64 = (v5 ^ v5)
                                                                    v43 := GenS5{F0: 1000, F1: (v1 * v1)}
                                                                    v44 := []int{-228, fi35, 673, (fi35 * (fi35 & -1))}
                                                                    _ = v42
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                            }
                                                        } else {
                                                            fmt.Println(v31)
                                                        }
                                                        v13 = ((v1 | 255) != (v1 / ((v1) | 1)))
                                                        _ = v40
                                                    } else {
                                                        if (true || (!v13)) {
                                                            var v45 uint16 = (v4 * (^(v4)))
                                                            if (-362 < 826) {
                                                                if (true && ((^(v31)) >= v31)) {
                                                                    v46 := []int{fi34, (fi34 - (-(fi35))), (fi34 | 42), 703}
                                                                    _ = len(v46)
                                                                    _ = v46
                                                                }
                                                                switch (v39 + v39) {
                                                                case "":
                                                                    fmt.Println(fi35, v39)
                                                                default:
                                                                    if v13 {
                                                                        v2 = v2
                                                                        var v47 int32 = v2
                                                                        _ = v47
                                                                    }
                                                                }
                                                                var v48 uint16 = (^(v4))
                                                                _ = v48
                                                            }
                                                            _ = v45
                                                        } else {
                                                            if v13 {
                                                                v13 = (!v13)
                                                                var v49 int8 = v31
                                                                _ = v49
                                                            }
                                                            v13 = (!v13)
                                                            switch (fi35 ^ 710) {
                                                            case -2147483648:
                                                                fmt.Println(v13, fi35, fi34)
                                                            default:
                                                                v4 = (^(v4))
                                                            }
                                                        }
                                                        var v50 int32 = (-(v2))
                                                        if (!((fi34 ^ fi35) <= fi34)) {
                                                            v1 = v1
                                                            v51 := []int{128, 4294967295}
                                                            fmt.Println(fi34)
                                                            _ = v51
                                                        }
                                                        var v52 int16 = 0
                                                        v53 := []int{446, fi34, 32768, -9223372036854775808}
                                                        _ = v50
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                    fmt.Println(v4, v4)
                                                    v54 := GenS6{F0: v39, F1: v39}
                                                    _ = v54
                                                }
                                            }
                                            var v55 int8 = (v31 << ((v1) & 63))
                                            _ = v39
                                            _ = v55
                                        } else {
                                            var v56 uint16 = 2
                                            _ = v56
                                        }
                                    } else {
                                        var v57 float32 = float32((0.0 - 1.5))
                                        _ = v57
                                    }
                                }
                                var v58 bool = ((^(v1)) >= 2)
                                _ = v58
                            } else {
                                var v59 int16 = 974
                                _ = v59
                            }
                            if true {
                                v3["a"] = -311
                            }
                            switch "\n\t" {
                            case "foo bar":
                                switch (fi34 % (((fi34 + (fi34 | (fi34 / ((fi34) | 1))))) | 1)) {
                                case -2147483648:
                                    var v60 uint32 = 157
                                    _ = v60
                                }
                            case "\n\t":
                                v61 := GenS7{F0: v13}
                                _ = v61
                            }
                            v13 = v13
                        }
                        v5 = 700
                    }
                }
                _ = fi35
            }
            if ((v5 * 619) > v5) {
                for fi62 := 0; fi62 < 20; fi62++ {
                    v2 = -1
                    _ = fi62
                }
                if v13 {
                    v63 := v3["\n\t"]
                    fmt.Println(fi34, v2)
                    if (v2 != (-(v2))) {
                        for fi64 := 0; fi64 < 7; fi64++ {
                            v65 := GenS1{F0: 42}
                            delete(v3, "\n\t")
                            v66 := []int{v63, (fi64 << ((1) & 63))}
                            v67 := GenS8{F0: 184, F1: (-(v31)), F2: (v2 << (((v1 - v1)) & 63))}
                            _ = fi64
                            _ = v65
                            _ = v66
                            _ = v67
                        }
                    } else {
                        var v68 uint8 = 128
                        switch "foo bar" {
                        case "":
                            if (1 <= 1) {
                                if v13 {
                                    v63 += v63
                                    for fi69 := 0; fi69 < 15; fi69++ {
                                        var v70 uint64 = v5
                                        fmt.Println(v70, v2, v63)
                                        var v71 uint32 = 4294967295
                                        _ = fi69
                                        _ = v70
                                        _ = v71
                                    }
                                } else {
                                    v72 := GenS2{F0: fi34, F1: (v5 ^ 42)}
                                    v2 += v2
                                    _ = v72
                                }
                                var v73 string = (("a" + "a") + ("foo bar" + ("hello" + "foo bar")))
                                v68 -= v68
                                _ = v73
                            } else {
                                v63 -= v63
                            }
                        case "hello":
                            var v74 float32 = float32((((1e10 + (0.1 + float64(((0.1 - ((0.0 + 1.5) / 3.14159)) + 0.1)))) / (((float64(float32(1.5)) * (-2.25 + 1.5)) * 1e10) - (1.5 / -2.25))) * 0.0))
                            _ = v74
                        case "a":
                            for fi75 := 0; fi75 < 4; fi75++ {
                                v3["a"] = (v63 / (((fi34 % ((fi34) | 1))) | 1))
                                _ = fi75
                            }
                        default:
                            for fi76 := 0; fi76 < 4; fi76++ {
                                var v77 int8 = (v31 - (v31 % (((v31 << ((540) & 63))) | 1)))
                                _ = fi76
                                _ = v77
                            }
                        }
                        fmt.Println(v63, fi34)
                        _ = v68
                    }
                    _ = v63
                }
                v78 := GenS6{F0: "", F1: "foo bar"}
                _ = v78
            }
            for fi79 := 0; fi79 < 6; fi79++ {
                switch 683 {
                case 414:
                    v2 -= v2
                case 42:
                    v1 += 4294967295
                }
                if v13 {
                    switch "a" {
                    case "hello":
                        fmt.Println(v2)
                    case "":
                        fmt.Println(v2, v31)
                        v13 = (255 > 128)
                    case "\n\t":
                        var v80 string = "foo bar"
                        _ = v80
                    default:
                        if (((2147483647 <= -2147483648) || (65535 == 65535)) && v13) {
                            fmt.Println(v1, v5)
                        }
                    }
                    v5 -= 0
                } else {
                    fmt.Println(v5, v13, v31)
                    v81 := GenS3{F0: 190, F1: 32767, F2: ""}
                    switch -1 {
                    case 801:
                        v3[("a" + ("a" + "foo bar"))] = fi34
                    default:
                        var v82 bool = ((v5 % (((v5 / ((42) | 1))) | 1)) != v5)
                        _ = v82
                    }
                    _ = v81
                }
                _ = fi79
            }
            _ = fi34
        }
    }
    if v13 {
        var v83 bool = (48 <= v31)
        _ = v83
    } else {
        switch "" {
        case "\n\t":
            var v84 int32 = (^(v2))
            _ = v84
        }
        if v13 {
            v1 = (v1 * v1)
            var v85 uint = v1
            _ = v85
        } else {
            v13 = (!false)
        }
        if v13 {
            v31 = v31
        } else {
            v86 := GenS9{F0: -74, F1: genFn0()}
            _ = v86
        }
    }
    if false {
        v87 := GenS4{F0: "a", F1: uint32(v4)}
        v88 := GenS8{F0: 128, F1: v31, F2: (-(v2))}
        v89 := GenS10{F0: "hello", F1: (^(v2)), F2: v5}
        _ = v87
        _ = v88
        _ = v89
    }
    fmt.Println(v1)
    v90 := []int{-790, 243}
    for fi91 := 0; fi91 < 19; fi91++ {
        if (!false) {
            if v13 {
                v13 = (v13 || v13)
                if (((("" + ("\n\t" + "hello")) <= "foo bar") && v13) || false) {
                    fmt.Println(v31, v4)
                } else {
                    fmt.Println(v31, v4, v2)
                    var v92 float32 = 0.1
                    var v93 uint64 = (v5 - (v5 + (v5 / ((18446744073709551615) | 1))))
                    _ = v92
                    _ = v93
                }
                for fi94 := 0; fi94 < 8; fi94++ {
                    v31 = v31
                    fmt.Println(v31)
                    var v95 int32 = (^(v2))
                    _ = fi94
                    _ = v95
                }
            }
            for fi96 := 0; fi96 < 10; fi96++ {
                var v97 int8 = v31
                var v98 int8 = v97
                var v99 uint = 255
                _ = fi96
                _ = v97
                _ = v98
                _ = v99
            }
        }
        var v100 uint64 = (v5 + (v5 % (((v5 | v5)) | 1)))
        _ = fi91
        _ = v100
    }
    var v101 int64 = int64(v2)
    fmt.Println(v101)
    switch 127 {
    case 0:
        var v102 uint32 = uint32(v4)
        _ = v102
    }
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v90:", len(v90), v90)
    _ = v1
    _ = v101
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v90
}
