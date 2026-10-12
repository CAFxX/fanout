package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 int64; F2 uint32 }
type GenS8 struct { F0 int32; F1 uint64; F2 int16 }
type GenS9 struct { F0 bool; F1 uint }
type GenS10 struct { F0 uint8; F1 uint; F2 int8 }
type GenS11 struct { F0 uint8; F1 uint16; F2 int }
type GenS12 struct { F0 uint64; F1 uint32 }
type GenS13 struct { F0 int; F1 int16; F2 uint16 }

func main() {
    v1 := []string{"\n\t", "", ("hello" + ("foo bar" + "\n\t")), (("hello" + "\n\t") + "a")}
    var v2 uint8 = 2
    var v3 string = ("\n\t" + ("\n\t" + ("\n\t" + "hello")))
    var v4 int = 97
    var v5 int8 = 2
    fmt.Println(v5)
    var v6 float64 = float64(-2.25)
    for fi7 := 0; fi7 < 10; fi7++ {
        if true {
            var v8 bool = false
            switch v3 {
            case "":
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 bool = v8
                    _ = fi9
                    _ = v10
                }
            case "a":
                for fi11 := 0; fi11 < 12; fi11++ {
                    switch (v3 + "\n\t") {
                    case "foo bar":
                        _ = len(v1)
                    case "":
                        if v8 {
                            switch v4 {
                            case 4294967295:
                                switch (fi7 & (fi11 - fi11)) {
                                case -534:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case -387:
                                    v3 = (v3 + "hello")
                                }
                            case 1:
                                v5 += (v5 - v5)
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch (v4 ^ 128) {
                            case -9223372036854775808:
                                v1 = append(v1, ((("\n\t" + (v3 + "\n\t")) + v3) + v3))
                            default:
                                if (42 != 32767) {
                                    var v12 bool = (18446744073709551615 < 32)
                                    switch "hello" {
                                    case "\n\t":
                                        if (!true) {
                                            switch "foo bar" {
                                            case "":
                                                v13 := GenS1{F0: 391, F1: 255}
                                                _ = v13
                                            case "\n\t":
                                                if v8 {
                                                    fmt.Println(v6)
                                                    v14 := GenS1{F0: 588, F1: 836}
                                                    _ = v14
                                                } else {
                                                    switch "foo bar" {
                                                    case "a":
                                                        v12 = v8
                                                    default:
                                                        var v15 float32 = -2.25
                                                        _ = v15
                                                    }
                                                }
                                            }
                                            switch fi7 {
                                            case 32767:
                                                switch (((("a" + (v3 + v3)) + v3) + v3) + "a") {
                                                case "a":
                                                    v12 = ((fi7 / ((fi11) | 1)) == (v4 - (fi11 / ((127) | 1))))
                                                case "\n\t":
                                                    var v16 bool = false
                                                    v8 = (!(v12 && v12))
                                                    _ = v16
                                                }
                                            case -9223372036854775808:
                                                if false {
                                                    switch (("a" + "foo bar") + "\n\t") {
                                                    case "\n\t":
                                                        v17 := GenS2{F0: (((v2 ^ 42) < 1) || (!((845 < 1) || (v2 >= v2)))), F1: 581}
                                                        _ = v17
                                                    case "hello":
                                                        v5 = (v5 * v5)
                                                    default:
                                                        fmt.Println(v3, v6)
                                                    }
                                                    switch fi11 {
                                                    case -9223372036854775808:
                                                        v2 -= 128
                                                    case -2147483648:
                                                        if (18446744073709551615 < 42) {
                                                            if v12 {
                                                                if false {
                                                                    if v8 {
                                                                        v3 = ""
                                                                        var v18 uint32 = 2147483647
                                                                        if ("\n\t" < v3) {
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            var v19 float32 = 0.1
                                                                            v6 = float64(v19)
                                                                            if v12 {
                                                                                fmt.Println(v18, v6, v12)
                                                                                var v20 int8 = (-(v5))
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = v3
                                                                                }
                                                                                var v21 int8 = (v20 | v5)
                                                                                switch ((v3 + ((v3 + ("a" + (v3 + ((v3 + v3) + (v3 + v3))))) + (v3 + v3))) + v3) {
                                                                                case "a":
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        var v22 float64 = (v6 - (((3.14159 - 1e10) - v6) + v6))
                                                                                        _ = v22
                                                                                    case "foo bar":
                                                                                        if (!v12) {
                                                                                            switch ((v3 + (v3 + "")) + "hello") {
                                                                                            case "\n\t":
                                                                                                switch fi11 {
                                                                                                case 42:
                                                                                                    switch v3 {
                                                                                                    case "\n\t":
                                                                                                        var v23 float32 = 0.0
                                                                                                        _ = v23
                                                                                                    case "a":
                                                                                                        if v12 {
                                                                                                            fmt.Println(v21)
                                                                                                            var v24 uint64 = 135
                                                                                                            _ = v24
                                                                                                        } else {
                                                                                                            var v25 int64 = 812
                                                                                                            _ = len(v1)
                                                                                                            _ = v25
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            case "":
                                                                                                var v26 uint32 = v18
                                                                                                _ = v26
                                                                                            case "a":
                                                                                                fmt.Println(v19, v3, v4)
                                                                                            }
                                                                                        }
                                                                                    case "":
                                                                                        v21 -= (v5 ^ 1)
                                                                                    default:
                                                                                        v4 = v4
                                                                                    }
                                                                                case "\n\t":
                                                                                    if false {
                                                                                        var v27 uint16 = 32767
                                                                                        _ = len(v1)
                                                                                        var v28 int64 = -1
                                                                                        if true {
                                                                                            v29 := GenS3{F0: -9, F1: -32768}
                                                                                            _ = v29
                                                                                        }
                                                                                        var v30 bool = (!(fi11 <= (fi11 & 4294967295)))
                                                                                        _ = v27
                                                                                        _ = v28
                                                                                        _ = v30
                                                                                    }
                                                                                }
                                                                                if v8 {
                                                                                    if (!v8) {
                                                                                        if v12 {
                                                                                            fmt.Println(v12, v12)
                                                                                            var v31 int32 = 452
                                                                                            _ = v31
                                                                                        }
                                                                                        var v32 string = "a"
                                                                                        v8 = (!false)
                                                                                        _ = v32
                                                                                    } else {
                                                                                        var v33 bool = (((v6 * v6) * (1.5 / 0.1)) >= v6)
                                                                                        v8 = (1 < 0)
                                                                                        v34 := GenS1{F0: 0, F1: 509}
                                                                                        _ = v33
                                                                                        _ = v34
                                                                                    }
                                                                                    _ = len(v1)
                                                                                }
                                                                                var v35 int32 = -2147483648
                                                                                switch (((("foo bar" + "") + v3) + v3) + v3) {
                                                                                case "":
                                                                                    var v36 float64 = v6
                                                                                    _ = v36
                                                                                default:
                                                                                    var v37 bool = true
                                                                                    _ = v37
                                                                                }
                                                                                v1 = append(v1, ("" + ""))
                                                                                v38 := GenS2{F0: v12, F1: v18}
                                                                                _ = v20
                                                                                _ = v21
                                                                                _ = v35
                                                                                _ = v38
                                                                            }
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                switch (fi7 / ((fi7) | 1)) {
                                                                                case 32767:
                                                                                    if ((v2 | (^(v2))) == 42) {
                                                                                        v19 = ((v19 - v19) * (v19 + 1.5))
                                                                                        v5 -= v5
                                                                                        v39 := GenS2{F0: (255 == 243), F1: v18}
                                                                                        v40 := GenS2{F0: ("hello" <= (v3 + ("hello" + v3))), F1: v18}
                                                                                        _ = v39
                                                                                        _ = v40
                                                                                    } else {
                                                                                        var v41 bool = (1 < 977)
                                                                                        if true {
                                                                                            fmt.Println(v5, v8)
                                                                                            if (false || v8) {
                                                                                                switch fi11 {
                                                                                                case 2147483647:
                                                                                                    _ = len(v1)
                                                                                                case 239:
                                                                                                    var v42 uint = 1
                                                                                                    _ = v42
                                                                                                case -1:
                                                                                                    var v43 uint32 = (^(v18))
                                                                                                    _ = v43
                                                                                                default:
                                                                                                    if true {
                                                                                                        switch v3 {
                                                                                                        case "\n\t":
                                                                                                            switch v3 {
                                                                                                            case "\n\t":
                                                                                                                v8 = v8
                                                                                                            default:
                                                                                                                v19 = 0.0
                                                                                                            }
                                                                                                        default:
                                                                                                            fmt.Println(fi7)
                                                                                                        }
                                                                                                        if (!v12) {
                                                                                                            v44 := GenS4{F0: uint16(v18)}
                                                                                                            _ = v44
                                                                                                        } else {
                                                                                                            v2 -= (v2 / (((v2 / (((v2 | (v2 - v2))) | 1))) | 1))
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v41
                                                                                    }
                                                                                case 127:
                                                                                    if (18446744073709551615 != 295) {
                                                                                        var v45 string = "foo bar"
                                                                                        v46 := GenS1{F0: 1, F1: 42}
                                                                                        var v47 int8 = 42
                                                                                        _ = v45
                                                                                        _ = v46
                                                                                        _ = v47
                                                                                    } else {
                                                                                        v48 := GenS5{F0: fi11}
                                                                                        var v49 int8 = v5
                                                                                        _ = v48
                                                                                        _ = v49
                                                                                    }
                                                                                default:
                                                                                    v50 := GenS1{F0: 0, F1: 592}
                                                                                    _ = v50
                                                                                }
                                                                            case "":
                                                                                v1 = append(v1, (v3 + ((((v3 + v3) + ((((v3 + "hello") + "hello") + v3) + "hello")) + "\n\t") + v3)))
                                                                            default:
                                                                                v4 = (fi7 >> ((255) & 63))
                                                                            }
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            v51 := GenS1{F0: 0, F1: 448}
                                                                            _ = v19
                                                                            _ = v51
                                                                        } else {
                                                                            if (v8 || (!true)) {
                                                                                var v52 int16 = 138
                                                                                v53 := GenS6{F0: -1}
                                                                                _ = v52
                                                                                _ = v53
                                                                            } else {
                                                                                var v54 float32 = float32((v6 + float64(((0.1 + ((1e10 / ((1e10 * 0.1) + 3.14159)) / (3.14159 - (0.1 * 0.0)))) * 3.14159))))
                                                                                _ = v54
                                                                            }
                                                                            fmt.Println(v18, fi11, v18)
                                                                            if false {
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                v2 += v2
                                                                                var v55 int = 0
                                                                                _ = v55
                                                                            }
                                                                        }
                                                                        v18 = v18
                                                                        if false {
                                                                            switch ("a" + v3) {
                                                                            case "a":
                                                                                var v56 uint = 0
                                                                                _ = v56
                                                                            case "\n\t":
                                                                                var v57 int32 = 1
                                                                                _ = v57
                                                                            }
                                                                            fmt.Println(v2, v12)
                                                                            v58 := GenS2{F0: (!v8), F1: v18}
                                                                            _ = v58
                                                                        } else {
                                                                            var v59 int32 = 42
                                                                            fmt.Println(v2, v3)
                                                                            if ((v4 | (fi7 * fi11)) <= -50) {
                                                                                var v60 int64 = -9223372036854775808
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = v3
                                                                                }
                                                                                _ = v60
                                                                            } else {
                                                                                var v61 int8 = int8(v4)
                                                                                v1 = append(v1, (v3 + v3))
                                                                                var v62 float32 = (0.0 + 0.0)
                                                                                _ = v61
                                                                                _ = v62
                                                                            }
                                                                            _ = v59
                                                                        }
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        v1 = append(v1, (((v3 + (v3 + v3)) + ("a" + "foo bar")) + v3))
                                                                        v63 := GenS7{F0: "foo bar", F1: 0, F2: v18}
                                                                        _ = v18
                                                                        _ = v63
                                                                    }
                                                                    v5 = v5
                                                                    var v64 uint16 = 42
                                                                    _ = v64
                                                                }
                                                                if len(v1) > 0 {
                                                                    v1[0] = ("" + (("a" + v3) + v3))
                                                                }
                                                                var v65 int32 = 557
                                                                if len(v1) > 0 {
                                                                    v1[0] = v3
                                                                }
                                                                _ = v65
                                                            }
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = v3
                                                        }
                                                        v2 = 127
                                                    default:
                                                        var v66 uint = 65535
                                                        _ = v66
                                                    }
                                                }
                                            default:
                                                v3 = ((v3 + v3) + v3)
                                            }
                                            var v67 int32 = 463
                                            _ = v67
                                        } else {
                                            if false {
                                                switch (v3 + v3) {
                                                case "hello":
                                                    fmt.Println(v12, fi7)
                                                }
                                                if (18446744073709551615 < 1) {
                                                    var v68 uint16 = 0
                                                    v69 := GenS1{F0: -1, F1: 1}
                                                    v4 = (fi11 | (-(fi11)))
                                                    _ = v68
                                                    _ = v69
                                                } else {
                                                    var v70 int8 = (v5 | (v5 >> ((210) & 63)))
                                                    _ = v70
                                                }
                                            } else {
                                                fmt.Println(v2, fi11, fi7)
                                                var v71 int = fi7
                                                var v72 uint16 = 1
                                                _ = v71
                                                _ = v72
                                            }
                                        }
                                    case "a":
                                        var v73 int16 = 32767
                                        _ = v73
                                    case "foo bar":
                                        v12 = true
                                    default:
                                        if ((894 >= 2147483647) && (!(v8 || v8))) {
                                            if v8 {
                                                var v74 float64 = 1e10
                                                fmt.Println(v5, v12)
                                                v2 += (v2 >> ((2) & 63))
                                                v4 = v4
                                                if false {
                                                    var v75 uint32 = 4294967295
                                                    switch "" {
                                                    case "foo bar":
                                                        v76 := GenS8{F0: -2147483648, F1: 0, F2: 0}
                                                        var v77 bool = v8
                                                        _ = v76
                                                        _ = v77
                                                    default:
                                                        v12 = (1.5 != (0.0 - 0.0))
                                                    }
                                                    if false {
                                                        switch v3 {
                                                        case "hello":
                                                            v75 = (^(v75))
                                                        }
                                                        v78 := GenS9{F0: (18446744073709551615 == 366), F1: uint(v2)}
                                                        _ = v78
                                                    }
                                                    _ = v75
                                                }
                                                _ = v74
                                            } else {
                                                if len(v1) > 0 {
                                                    v1[0] = "hello"
                                                }
                                            }
                                        }
                                    }
                                    _ = v12
                                } else {
                                    if (1 >= 42) {
                                        v1 = append(v1, v3)
                                        v79 := GenS3{F0: -211, F1: 1}
                                        _ = v79
                                    } else {
                                        if false {
                                            if v8 {
                                                v80 := GenS10{F0: (v2 + 1), F1: 849, F2: -1}
                                                var v81 uint64 = 219
                                                fmt.Println(fi11, v81, v5)
                                                _ = v80
                                                _ = v81
                                            }
                                            fmt.Println(v2, v2)
                                            v4 = (v4 & fi7)
                                        }
                                        fmt.Println(v5, v4, v8)
                                        if v8 {
                                            var v82 uint32 = 42
                                            _ = v82
                                        } else {
                                            v83 := GenS11{F0: 1, F1: 65535, F2: fi11}
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            fmt.Println(fi11)
                                            _ = v83
                                        }
                                    }
                                    v84 := GenS9{F0: (1 <= 0), F1: 255}
                                    if false {
                                        if false {
                                            v85 := GenS12{F0: 710, F1: 692}
                                            var v86 bool = ((((v6 + v6) / (v6 + (1e10 / v6))) >= (((0.0 + v6) / 0.1) / v6)) && v8)
                                            _ = v85
                                            _ = v86
                                        } else {
                                            switch ("" + "a") {
                                            case "foo bar":
                                                var v87 int64 = 655
                                                _ = v87
                                            case "":
                                                v88 := GenS13{F0: -1, F1: -223, F2: 2}
                                                _ = v88
                                            case "hello":
                                                var v89 string = v3
                                                _ = v89
                                            default:
                                                var v90 bool = v8
                                                _ = v90
                                            }
                                        }
                                        switch 9223372036854775807 {
                                        case 487:
                                            if (!true) {
                                                var v91 uint8 = 137
                                                var v92 uint = 42
                                                _ = v91
                                                _ = v92
                                            } else {
                                                v2 = v2
                                                v93 := GenS1{F0: 9223372036854775807, F1: 2}
                                                switch (v3 + v3) {
                                                case "\n\t":
                                                    var v94 int64 = -9223372036854775808
                                                    _ = v94
                                                case "":
                                                    if ((v2 << ((2) & 63)) > (v2 ^ (v2 | 100))) {
                                                        v6 = (1e10 - (float64((((0.1 * 0.1) + ((float32((float64((1e10 + float32((v6 * v6)))) / v6)) + ((0.0 - (0.0 * (-2.25 + 0.1))) * ((1e10 - 1e10) + 0.1))) + (0.1 * (1.5 / 1.0)))) * float32((v6 * v6)))) - v6))
                                                        var v95 int = 32767
                                                        var v96 uint32 = 589
                                                        v2 = (v2 + v2)
                                                        var v97 int8 = 0
                                                        if (332 != v96) {
                                                            var v98 string = "\n\t"
                                                            var v99 int32 = -617
                                                            v1 = append(v1, (("foo bar" + v3) + (v3 + (v3 + "foo bar"))))
                                                            _ = v98
                                                            _ = v99
                                                        } else {
                                                            var v100 string = ("a" + (v3 + v3))
                                                            _ = v100
                                                        }
                                                        var v101 string = v3
                                                        _ = v101
                                                        _ = v95
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                }
                                                _ = v93
                                            }
                                        case 32767:
                                            if v8 {
                                                fmt.Println(v5)
                                                switch v3 {
                                                case "\n\t":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case "foo bar":
                                                    v1 = append(v1, ("" + ""))
                                                case "hello":
                                                    if v8 {
                                                        v3 = v3
                                                        v5 = (v5 >> ((1) & 63))
                                                    }
                                                default:
                                                    v2 = 0
                                                }
                                                v102 := GenS11{F0: v2, F1: 32767, F2: fi11}
                                                var v103 uint8 = v2
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = len(v1)
                                                _ = v102
                                                _ = v103
                                            }
                                        }
                                    } else {
                                        switch fi7 {
                                        case 128:
                                            v104 := GenS10{F0: v2, F1: 18446744073709551615, F2: v5}
                                            _ = v104
                                        case 42:
                                            var v105 uint = 733
                                            _ = v105
                                        default:
                                            if v8 {
                                                var v106 int16 = 0
                                                _ = v106
                                            } else {
                                                v2 -= (v2 << ((4294967295) & 63))
                                            }
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = "foo bar"
                                        }
                                    }
                                    if (2 >= 65535) {
                                        switch (^(fi11)) {
                                        case 9223372036854775807:
                                            var v107 float32 = 1.5
                                            _ = v107
                                        default:
                                            v3 = v3
                                        }
                                    }
                                    var v108 float64 = float64(1.5)
                                    switch -651 {
                                    case 42:
                                        if v8 {
                                            if (true || v8) {
                                                v8 = (67 != 793)
                                            }
                                        }
                                    case 503:
                                        v8 = (false || v8)
                                    case 127:
                                        var v109 string = "foo bar"
                                        _ = v109
                                    default:
                                        v3 = v3
                                    }
                                    _ = v108
                                    _ = v84
                                }
                            }
                        }
                    case "hello":
                        _ = len(v1)
                    }
                    _ = fi11
                }
            }
            _ = v8
        }
        _ = fi7
    }
    v5 -= v5
    var v110 int32 = 560
    var v111 uint16 = 851
    var v112 float32 = 3.14159
    v111 = v111
    for fi113 := 0; fi113 < 4; fi113++ {
        v1 = append(v1, v3)
        v2 = (v2 % ((2) | 1))
        fmt.Println(v110, v2, v4)
        _ = fi113
    }
    v1 = append(v1, "hello")
    v6 += v6
    for fi114 := 0; fi114 < 11; fi114++ {
        var v115 int32 = (v110 << ((1) & 63))
        fmt.Println(v2, v2, v4)
        _ = fi114
        _ = v115
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v110:", v110)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v110
    _ = v111
    _ = v112
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
