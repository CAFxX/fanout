package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int32; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 int16; F1 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int8; F1 uint64; F2 int32 }
type GenS10 struct { F0 int8; F1 uint8 }
type GenS11 struct { F0 int32 }
func genFn0() int8 {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint64 = 42
        v8 := []int{(fi6 + (fi6 * fi6))}
        _ = len(v8)
        _ = fi6
        _ = v7
        _ = v8
    }
    switch -1 {
    case 128:
        var v9 int16 = 213
        _ = v9
    case 32767:
        var v10 string = "\n\t"
        _ = v10
    default:
        v11 := []int{}
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            switch "\n\t" {
            case "hello":
                var v14 int64 = 122
                _ = v14
            case "":
                fmt.Println(fi13)
            case "\n\t":
                var v15 bool = (!(1 < 42))
                switch "foo bar" {
                case "a":
                    if false {
                        var v16 string = ""
                        fmt.Println(fi13, v15, fi13)
                        _ = v16
                    }
                case "hello":
                    v15 = v15
                default:
                    v15 = ((0 == -2147483648) || (4294967295 < 4294967295))
                }
                fmt.Println(fi13, v15, fi13)
                _ = v15
            default:
                var v17 string = "foo bar"
                _ = v17
            }
            fmt.Println(fi13, fi12)
            _ = fi13
        }
        if ((float32(3.14159) != (1e10 / 0.1)) || true) {
            if true {
                switch fi12 {
                case -562:
                    var v18 float64 = 3.14159
                    _ = v18
                default:
                    var v19 string = ("foo bar" + ("foo bar" + "foo bar"))
                    _ = v19
                }
                if true {
                    var v20 uint16 = 0
                    if (!(!false)) {
                        fmt.Println(v20)
                        if (((float64(-2.25) == (((0.1 + 0.1) * 1e10) - 3.14159)) || (100 >= 13)) || (-580 == -762)) {
                            switch (-(fi12)) {
                            case 32:
                                fmt.Println(fi12, v20)
                            case -1:
                                var v21 float64 = ((0.1 / 3.14159) * 0.0)
                                _ = v21
                            default:
                                v22 := GenS1{F0: "\n\t"}
                                _ = v22
                            }
                        } else {
                            v23 := []int{-9223372036854775808, fi12, 4294967295, (fi12 % ((1) | 1))}
                            var v24 int = fi12
                            _ = v23
                            _ = v24
                        }
                    }
                    v20 += v20
                    _ = v20
                } else {
                    v25 := []int{fi12, (fi12 << ((65535) & 63))}
                    v26 := GenS1{F0: ""}
                    _ = v25
                    _ = v26
                }
                if (!(42 >= -1)) {
                    fmt.Println(fi12)
                    if (true && (255 >= 255)) {
                        for fi27 := 0; fi27 < 1; fi27++ {
                            v28 := []int{-9223372036854775808, fi12, fi12, 584}
                            _ = fi27
                            _ = v28
                        }
                        if ("hello" >= "foo bar") {
                            v29 := []int{fi12, 32768, fi12, (fi12 << ((612) & 63))}
                            _ = v29
                        } else {
                            var v30 int32 = 2147483647
                            if true {
                                if (2 < 0) {
                                    if false {
                                        for fi31 := 0; fi31 < 9; fi31++ {
                                            var v32 uint64 = 0
                                            v33 := GenS1{F0: "\n\t"}
                                            v34 := []int{(-(fi31))}
                                            v35 := GenS1{F0: "a"}
                                            v36 := GenS2{F0: -11, F1: v30, F2: 581}
                                            switch fi12 {
                                            case -2147483648:
                                                v32 = v32
                                            case 32767:
                                                if false {
                                                    v34 = append(v34, (fi31 / ((fi31) | 1)))
                                                    v32 = v32
                                                    var v37 int64 = 9223372036854775807
                                                    var v38 int16 = 0
                                                    _ = v37
                                                    _ = v38
                                                }
                                            case 42:
                                                var v39 uint32 = 42
                                                _ = v39
                                            default:
                                                v30 = v30
                                            }
                                            if (1 < -128) {
                                                switch fi31 {
                                                case 127:
                                                    if len(v34) >= 2 {
                                                        v34 = v34[1:len(v34)]
                                                    }
                                                case 1:
                                                    v32 -= v32
                                                }
                                                var v40 float32 = 1.5
                                                var v41 int64 = 705
                                                fmt.Println(v32, fi31, fi31)
                                                _ = v40
                                                _ = v41
                                            } else {
                                                v42 := GenS2{F0: -128, F1: (v30 * v30), F2: v32}
                                                if true {
                                                    v30 += (v30 ^ (v30 | v30))
                                                    if (-1 >= -1) {
                                                        if (!true) {
                                                            var v43 uint64 = v32
                                                            _ = v43
                                                        } else {
                                                            var v44 int8 = -67
                                                            v45 := GenS1{F0: "foo bar"}
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                        v46 := GenS3{F0: 9223372036854775807}
                                                        _ = v46
                                                    }
                                                }
                                                _ = v42
                                            }
                                            v32 = v32
                                            v32 = v32
                                            _ = fi31
                                            _ = v32
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                            _ = v36
                                        }
                                        v30 = -2147483648
                                    }
                                }
                                v47 := GenS1{F0: ((("a" + (("\n\t" + "\n\t") + "foo bar")) + "") + "\n\t")}
                                for fi48 := 0; fi48 < 14; fi48++ {
                                    v30 += (v30 | v30)
                                    var v49 uint8 = 12
                                    _ = fi48
                                    _ = v49
                                }
                                _ = v47
                            }
                            _ = v30
                        }
                        var v50 int8 = 1
                        _ = v50
                    }
                    var v51 int16 = 252
                    _ = v51
                } else {
                    var v52 bool = true
                    var v53 uint32 = 1
                    switch "\n\t" {
                    case "a":
                        var v54 float64 = (0.1 * float64((((1.5 - 0.0) * 1.5) * (1.5 * 0.1))))
                        _ = v54
                    case "foo bar":
                        var v55 int16 = 0
                        _ = v55
                    default:
                        fmt.Println(v53, fi12)
                    }
                    _ = v52
                    _ = v53
                }
            }
        } else {
            v56 := []int{617, fi12, (fi12 % (((fi12 % ((143) | 1))) | 1)), fi12}
            _ = v56
        }
        _ = fi12
    }
    v57 := []int{}
    _ = v57
    return 42
}

func main() {
    var v1 int16 = 1
    var v2 float32 = (0.0 + (-2.25 - 0.1))
    v3 := []uint32{316, 42, 0}
    v4 := make(map[int]int)
    var v5 int64 = 452
    v5 = v5
    if true {
        if (((341 != (v1 & 42)) || (v2 > v2)) || true) {
            v58 := GenS3{F0: v5}
            switch "" {
            case "hello":
                var v59 uint16 = 255
                _ = v59
            }
            _ = v58
        }
    }
    switch -2147483648 {
    case -950:
        v60 := GenS4{F0: 454}
        _ = v60
    case 9223372036854775807:
        if (32768 != 42) {
            if (128 != 0) {
                switch 42 {
                case 1:
                    if (true && (v1 == v1)) {
                        if true {
                            switch 598 {
                            case 42:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                v1 -= v1
                            }
                        } else {
                            fmt.Println(v1)
                        }
                        v1 += (v1 | v1)
                        v61 := GenS1{F0: "hello"}
                        _ = v61
                    } else {
                        v2 += v2
                    }
                case -2147483648:
                    v62 := GenS5{F0: 4294967295, F1: 0}
                    _ = v62
                default:
                    for fi63 := 0; fi63 < 15; fi63++ {
                        v64 := GenS6{F0: "hello", F1: 32767}
                        if false {
                            v1 = v1
                        }
                        v65 := GenS3{F0: (^(v5))}
                        _ = fi63
                        _ = v64
                        _ = v65
                    }
                }
            } else {
                if (true && (false || (false || false))) {
                    v4[9223372036854775807] = 9223372036854775807
                    switch "" {
                    case "a":
                        var v66 string = (("foo bar" + ((("" + "\n\t") + "foo bar") + "a")) + "hello")
                        _ = v66
                    case "foo bar":
                        for fi67 := 0; fi67 < 2; fi67++ {
                            if (2147483647 <= 2147483647) {
                                var v68 string = ((("\n\t" + "\n\t") + ("" + ("hello" + (("a" + ("hello" + ("" + "a"))) + "a")))) + "hello")
                                _ = v68
                            }
                            var v69 uint8 = 2
                            v5 = v5
                            _ = fi67
                            _ = v69
                        }
                    default:
                        v2 -= ((3.14159 * float32((float64(v2) * (((-2.25 / 1.5) / 1.5) + 0.1)))) / (((3.14159 - ((v2 + 1e10) / (float32(0.0) + ((float32((-2.25 - 0.1)) + 1.5) * (v2 / (v2 + v2)))))) * 0.1) - v2))
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                    if (false || false) {
                        switch "\n\t" {
                        case "foo bar":
                            v2 = -2.25
                        }
                        v70 := GenS7{F0: (v1 << ((671) & 63)), F1: 219}
                        v71 := GenS7{F0: (v1 >> ((65535) & 63)), F1: 794}
                        _ = v70
                        _ = v71
                    } else {
                        delete(v4, -573)
                        v2 -= 3.14159
                    }
                }
                fmt.Println(v5)
                fmt.Println(v5, v2)
            }
            if (false && true) {
                if (42 == 127) {
                    v1 += v1
                    var v72 float32 = 0.1
                    switch "foo bar" {
                    case "foo bar":
                        var v73 int8 = 48
                        _ = v73
                    default:
                        for fi74 := 0; fi74 < 13; fi74++ {
                            if len(v3) > 0 {
                                v3[0] = 65535
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi74
                        }
                    }
                    _ = v72
                }
            } else {
                if false {
                    switch (("" + "hello") + ("foo bar" + "a")) {
                    case "foo bar":
                        if (1.5 > (0.0 + 0.0)) {
                            if ((v1 & (v1 | -1)) != (v1 * v1)) {
                                if (!true) {
                                    v4[9223372036854775807] = 1
                                    var v75 string = ("\n\t" + "\n\t")
                                    _ = v75
                                } else {
                                    if false {
                                        var v76 uint = 0
                                        for fi77 := 0; fi77 < 9; fi77++ {
                                            fmt.Println(v5, v76, v1)
                                            var v78 float32 = (v2 * float32(1e10))
                                            _ = fi77
                                            _ = v78
                                        }
                                        for fi79 := 0; fi79 < 12; fi79++ {
                                            fmt.Println(v1)
                                            switch "a" {
                                            case "hello":
                                                for fi80 := 0; fi80 < 1; fi80++ {
                                                    var v81 uint32 = 2147483648
                                                    _ = len(v4)
                                                    var v82 float64 = 0.0
                                                    _ = fi80
                                                    _ = v81
                                                    _ = v82
                                                }
                                            }
                                            v83 := GenS2{F0: -128, F1: 2147483647, F2: 1}
                                            _ = fi79
                                            _ = v83
                                        }
                                        _ = v76
                                    }
                                    var v84 bool = ((4294967295 <= 18446744073709551615) || false)
                                    if ((^(v1)) < (v1 & v1)) {
                                        var v85 uint64 = 2
                                        v84 = (v84 && v84)
                                        switch 1 {
                                        case 32768:
                                            v5 = (v5 >> ((2) & 63))
                                        default:
                                            var v86 uint8 = 1
                                            _ = v86
                                        }
                                        _ = v85
                                    } else {
                                        _ = len(v4)
                                        var v87 float32 = ((-2.25 - v2) * v2)
                                        _ = v87
                                    }
                                    _ = v84
                                }
                                fmt.Println(v1, v2)
                                _ = len(v3)
                            } else {
                                switch (("foo bar" + "") + "\n\t") {
                                case "":
                                    if len(v3) > 0 {
                                        v3[0] = 65535
                                    }
                                case "\n\t":
                                    fmt.Println(v1)
                                case "a":
                                    for fi88 := 0; fi88 < 16; fi88++ {
                                        var v89 int = fi88
                                        _ = fi88
                                        _ = v89
                                    }
                                default:
                                    var v90 float64 = -2.25
                                    _ = v90
                                }
                            }
                            v91 := GenS8{F0: -478}
                            var v92 uint = 65535
                            _ = v91
                            _ = v92
                        } else {
                            if false {
                                _ = len(v4)
                            } else {
                                v93 := GenS9{F0: -1, F1: 600, F2: -1}
                                var v94 uint32 = 668
                                _ = v93
                                _ = v94
                            }
                        }
                    }
                    v95 := GenS2{F0: 42, F1: 1, F2: 4294967295}
                    _ = v95
                }
                var v96 float64 = ((((0.0 * float64(v2)) / 1.0) * (3.14159 / 1.5)) - -2.25)
                v2 = 3.14159
                _ = v96
            }
        } else {
            var v97 uint16 = 1
            _ = v97
        }
    default:
        fmt.Println(v2, v1)
    }
    var v98 string = ""
    for fi99 := 0; fi99 < 2; fi99++ {
        var v100 int64 = v5
        _ = fi99
        _ = v100
    }
    var v101 float32 = 0.1
    v102 := GenS10{F0: -8, F1: 124}
    var v103 float32 = (v2 - (((0.0 / v101) / v101) * 1.5))
    switch int(v5) {
    case 32768:
        var v104 uint = 548
        _ = v104
    }
    var v105 int32 = 0
    if (127 >= -9223372036854775808) {
        _ = len(v4)
    } else {
        switch ((v98 + "\n\t") + (v98 + "a")) {
        case "foo bar":
            for fi106 := 0; fi106 < 20; fi106++ {
                var v107 float32 = v2
                var v108 uint = 65535
                var v109 float64 = 0.1
                _ = fi106
                _ = v107
                _ = v108
                _ = v109
            }
        default:
            v110 := GenS2{F0: -120, F1: v105, F2: 18446744073709551615}
            _ = v110
        }
        for fi111 := 0; fi111 < 14; fi111++ {
            _ = len(v4)
            v105 = (v105 / ((-2147483648) | 1))
            _ = fi111
        }
    }
    v1 = (v1 % ((1) | 1))
    _ = len(v4)
    v3 = append(v3, 65535)
    if (18446744073709551615 > 42) {
        if (false || (true && false)) {
            var v112 uint = 255
            if true {
                v98 = "\n\t"
                fmt.Println(v101, v2)
                fmt.Println(v98, v112)
            }
            _ = v112
        } else {
            v1 -= v1
            v113 := GenS11{F0: (^(v105))}
            var v114 int8 = 118
            _ = v113
            _ = v114
        }
    }
    delete(v4, 628)
    switch (v98 + v98) {
    case "hello":
        v98 = ("hello" + v98)
    case "foo bar":
        var v115 int = 42
        _ = v115
    }
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v103:", v103)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v101
    _ = v102
    _ = v103
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v98
}
