package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 int; F1 int8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int8; F1 int32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint32; F1 uint16 }
type GenS12 struct { F0 int; F1 int16; F2 int32 }
type GenS13 struct { F0 int32; F1 string }

func main() {
    v1 := []string{"foo bar", "hello", ""}
    var v2 uint64 = 620
    var v3 uint8 = 42
    var v4 float64 = 0.1
    var v5 uint64 = v2
    if true {
        var v6 uint16 = 32767
        switch (((("foo bar" + "a") + ("" + "")) + "a") + "\n\t") {
        case "foo bar":
            switch (("a" + ((("" + "\n\t") + "\n\t") + "hello")) + ("foo bar" + "a")) {
            case "":
                var v7 int8 = -64
                _ = v7
            case "\n\t":
                v3 += v3
            }
        case "hello":
            var v8 int16 = 32767
            _ = v8
        case "\n\t":
            v9 := GenS1{F0: 520}
            _ = v9
        default:
            var v10 bool = true
            _ = v10
        }
        switch (("" + ("" + "a")) + "") {
        case "\n\t":
            var v11 int = 127
            _ = v11
        case "":
            v6 = (v6 - (^(v6)))
        case "foo bar":
            v2 = uint64(v6)
        default:
            switch "" {
            case "a":
                v2 = v2
            case "foo bar":
                fmt.Println(v2)
            case "hello":
                if (v5 <= (^(v5))) {
                    if false {
                        switch ((("" + "\n\t") + (("" + ("" + "")) + "foo bar")) + "\n\t") {
                        case "":
                            v4 = v4
                        case "hello":
                            if false {
                                switch -338 {
                                case 32767:
                                    v12 := GenS1{F0: 101}
                                    _ = v12
                                case 9223372036854775807:
                                    fmt.Println(v5, v2, v2)
                                default:
                                    v6 = v6
                                }
                                if false {
                                    fmt.Println(v4, v5, v2)
                                    for fi13 := 0; fi13 < 1; fi13++ {
                                        fmt.Println(fi13, v5, v5)
                                        if ((v5 ^ (v5 % ((v5) | 1))) > (v2 * uint64(v3))) {
                                            if (true || false) {
                                                fmt.Println(v4)
                                                for fi14 := 0; fi14 < 14; fi14++ {
                                                    v15 := GenS1{F0: 0}
                                                    v2 = 1
                                                    v1 = append(v1, "foo bar")
                                                    var v16 uint64 = v5
                                                    var v17 float32 = 1.5
                                                    v18 := GenS1{F0: uint(v5)}
                                                    if (((("a" + ("hello" + (("hello" + "") + ("foo bar" + "hello")))) + "") + "") < "a") {
                                                        v19 := GenS2{F0: v5, F1: 1, F2: (v2 % (((v5 << ((697) & 63))) | 1))}
                                                        if (!((!(!false)) && true)) {
                                                            if (4294967295 <= 596) {
                                                                v5 = (^(v16))
                                                                var v20 uint = 4294967295
                                                                _ = v20
                                                            } else {
                                                                v17 += (1.5 / (((float32(1.5) - v17) * 0.1) / (0.1 / (1.5 + (((3.14159 - 3.14159) / 0.1) * v17)))))
                                                            }
                                                            var v21 string = ""
                                                            if (65535 >= uint(v6)) {
                                                                var v22 int8 = 64
                                                                v21 = ((v21 + v21) + "\n\t")
                                                                _ = v22
                                                            } else {
                                                                var v23 uint64 = v5
                                                                if (!false) {
                                                                    if true {
                                                                        if ((((!(true && true)) || (false && (v21 <= v21))) || (9223372036854775807 > -868)) || true) {
                                                                            if false {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = "hello"
                                                                                }
                                                                            } else {
                                                                                v21 = "\n\t"
                                                                            }
                                                                            fmt.Println(v2)
                                                                            var v24 string = v21
                                                                            _ = v24
                                                                        }
                                                                        v4 = 3.14159
                                                                    }
                                                                    v3 -= 1
                                                                    var v25 bool = (true && false)
                                                                    _ = v25
                                                                }
                                                                var v26 int = fi14
                                                                if true {
                                                                    var v27 float32 = 0.0
                                                                    if true {
                                                                        switch (v26 | (fi13 ^ v26)) {
                                                                        case 267:
                                                                            v28 := GenS3{F0: (v2 >> ((uint(v3)) & 63)), F1: -32768}
                                                                            _ = v28
                                                                        case 4294967295:
                                                                            v29 := GenS4{F0: (v3 & (v3 * 1)), F1: v3, F2: 2}
                                                                            if true {
                                                                                var v30 int64 = int64(v6)
                                                                                v3 = (v3 / ((v3) | 1))
                                                                                var v31 int8 = 0
                                                                                _ = v30
                                                                                _ = v31
                                                                            }
                                                                            fmt.Println(v27, v16, v23)
                                                                            _ = v29
                                                                        }
                                                                        fmt.Println(v26, v23, v5)
                                                                    } else {
                                                                        if false {
                                                                            if true {
                                                                                fmt.Println(v4, v6)
                                                                                v5 = v23
                                                                                v26 += fi13
                                                                                if ((0 == -32768) && true) {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = v21
                                                                                    }
                                                                                    var v32 int32 = -2147483648
                                                                                    fmt.Println(v4, v17)
                                                                                    _ = v32
                                                                                } else {
                                                                                    fmt.Println(v6, v5, v17)
                                                                                }
                                                                                if false {
                                                                                    switch fi14 {
                                                                                    case -471:
                                                                                        var v33 int = -9223372036854775808
                                                                                        _ = v33
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v23 -= 4294967295
                                                                                var v34 string = v21
                                                                                _ = v34
                                                                            }
                                                                            v1 = append(v1, v21)
                                                                        }
                                                                    }
                                                                    fmt.Println(v2)
                                                                    _ = v27
                                                                } else {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    v3 = (v3 & (v3 | v3))
                                                                }
                                                                v4 = 0.1
                                                                _ = v23
                                                                _ = v26
                                                            }
                                                            var v35 float64 = (0.1 / ((-2.25 / v4) - v4))
                                                            if true {
                                                                if false {
                                                                    switch fi13 {
                                                                    case 937:
                                                                        v17 -= v17
                                                                    case -707:
                                                                        _ = len(v1)
                                                                    default:
                                                                        if (float32((v4 - 3.14159)) < 1e10) {
                                                                            var v36 int16 = 52
                                                                            v4 += v35
                                                                            var v37 int64 = 9223372036854775807
                                                                            _ = v36
                                                                            _ = v37
                                                                        }
                                                                    }
                                                                    fmt.Println(v2, fi14)
                                                                    if (413 < 749) {
                                                                        var v38 int16 = -484
                                                                        var v39 uint8 = (^(v3))
                                                                        fmt.Println(v6)
                                                                        _ = v38
                                                                        _ = v39
                                                                    } else {
                                                                        switch (^(fi14)) {
                                                                        case 32768:
                                                                            v40 := GenS5{F0: (v3 ^ v3), F1: 0, F2: 32767}
                                                                            _ = v40
                                                                        case 127:
                                                                            var v41 int = (fi14 ^ (fi13 | (fi14 ^ -2147483648)))
                                                                            _ = v41
                                                                        default:
                                                                            var v42 float32 = (((1.5 / v17) - float32((v35 + ((((v35 * v35) - (v35 - float64(v17))) / float64(float32(1.5))) + float64(v17))))) / float32(v35))
                                                                            _ = v42
                                                                        }
                                                                        var v43 int8 = -87
                                                                        _ = v43
                                                                    }
                                                                }
                                                            }
                                                            _ = v21
                                                            _ = v35
                                                        }
                                                        if (!true) {
                                                            var v44 int = 46
                                                            switch "" {
                                                            case "a":
                                                                var v45 int = (v44 & (v44 & fi13))
                                                                _ = v45
                                                            case "":
                                                                v46 := GenS3{F0: v16, F1: 32767}
                                                                _ = v46
                                                            default:
                                                                if ((v6 <= 239) || true) {
                                                                    v47 := GenS6{F0: 42, F1: 679}
                                                                    _ = v47
                                                                }
                                                            }
                                                            _ = v44
                                                        } else {
                                                            if (0 < 716) {
                                                                var v48 bool = (false && (!true))
                                                                var v49 uint = 1
                                                                var v50 uint16 = (v6 % ((v6) | 1))
                                                                v17 = (1.5 - float32(0.0))
                                                                v51 := GenS5{F0: uint8(v6), F1: (v49 % (((v49 + (v49 << ((v49) & 63)))) | 1)), F2: 65535}
                                                                _ = v48
                                                                _ = v49
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                            if len(v1) > 0 {
                                                                v1[0] = ""
                                                            }
                                                        }
                                                        _ = v19
                                                    }
                                                    fmt.Println(v17)
                                                    var v52 int16 = 0
                                                    v2 += v2
                                                    _ = fi14
                                                    _ = v15
                                                    _ = v16
                                                    _ = v17
                                                    _ = v18
                                                    _ = v52
                                                }
                                            }
                                            if true {
                                                fmt.Println(fi13, v3, v3)
                                                v3 = v3
                                            } else {
                                                for fi53 := 0; fi53 < 8; fi53++ {
                                                    var v54 int8 = 127
                                                    var v55 int = -604
                                                    _ = fi53
                                                    _ = v54
                                                    _ = v55
                                                }
                                                v5 = v2
                                                var v56 string = ""
                                                _ = v56
                                            }
                                            v4 = v4
                                        } else {
                                            v57 := GenS6{F0: 100, F1: 0}
                                            _ = v57
                                        }
                                        v3 = (v3 / ((2) | 1))
                                        _ = fi13
                                    }
                                    var v58 uint = uint(v6)
                                    _ = v58
                                } else {
                                    var v59 int64 = 697
                                    _ = v59
                                }
                                v1 = append(v1, "foo bar")
                            }
                        case "a":
                            var v60 int16 = 42
                            _ = v60
                        }
                        v61 := GenS3{F0: (v2 * 42), F1: -32768}
                        _ = v61
                    } else {
                        v62 := GenS1{F0: 4294967295}
                        _ = v62
                    }
                    for fi63 := 0; fi63 < 10; fi63++ {
                        var v64 uint16 = uint16(v5)
                        _ = fi63
                        _ = v64
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                }
            default:
                switch "\n\t" {
                case "hello":
                    if (42 <= 394) {
                        fmt.Println(v2, v3)
                        v65 := GenS7{F0: 127, F1: -25}
                        v66 := GenS5{F0: (^(v3)), F1: 2, F2: v6}
                        _ = v65
                        _ = v66
                    }
                case "foo bar":
                    v5 += 1
                default:
                    v5 = v5
                }
            }
        }
        _ = v6
    } else {
        var v67 uint = 0
        if false {
            fmt.Println(v4, v3, v3)
            if false {
                v67 = v67
                v4 -= v4
                fmt.Println(v4, v5)
            }
            for fi68 := 0; fi68 < 9; fi68++ {
                v1 = append(v1, "hello")
                for fi69 := 0; fi69 < 20; fi69++ {
                    var v70 int16 = -32768
                    var v71 bool = false
                    _ = fi69
                    _ = v70
                    _ = v71
                }
                _ = fi68
            }
        } else {
            switch 4294967295 {
            case 1:
                for fi72 := 0; fi72 < 7; fi72++ {
                    for fi73 := 0; fi73 < 7; fi73++ {
                        fmt.Println(v2, v3)
                        _ = fi73
                    }
                    var v74 string = (("a" + "foo bar") + "")
                    _ = fi72
                    _ = v74
                }
            case 54:
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
            }
            v75 := GenS2{F0: (^(v5)), F1: -1, F2: v5}
            if true {
                v2 += (v2 << ((21) & 63))
            } else {
                v76 := GenS8{F0: 2147483648}
                if true {
                    switch ("a" + "a") {
                    case "a":
                        var v77 int = int(v67)
                        _ = v77
                    default:
                        v78 := GenS9{F0: 127, F1: -2147483648}
                        _ = v78
                    }
                    v79 := GenS10{F0: -503}
                    if (float64(0.0) == (((-2.25 - v4) / float64(1e10)) + ((((((v4 * (float64(0.0) - (float64(1e10) * v4))) / 0.1) - (float64((1.5 + (3.14159 * 1e10))) + v4)) + (v4 + (-2.25 / (0.1 / (((v4 + (float64(0.0) - float64(-2.25))) * v4) - (float64((-2.25 + (float32(v4) / 1.0))) * 1.5)))))) - (v4 - v4)) / (1e10 * (((((3.14159 - 1e10) + 3.14159) * -2.25) - (float64(-2.25) - v4)) + v4))))) {
                        if (false && (9223372036854775807 != 832)) {
                            if false {
                                v80 := GenS11{F0: 454, F1: 2}
                                v3 = 2
                                switch "a" {
                                case "foo bar":
                                    v81 := GenS5{F0: 255, F1: v67, F2: 32767}
                                    _ = v81
                                default:
                                    for fi82 := 0; fi82 < 15; fi82++ {
                                        var v83 float64 = 1e10
                                        for fi84 := 0; fi84 < 19; fi84++ {
                                            var v85 uint32 = 42
                                            if false {
                                                v5 -= (v2 ^ (v2 << ((v67) & 63)))
                                                var v86 uint = (v67 ^ v67)
                                                _ = v86
                                            }
                                            _ = fi84
                                            _ = v85
                                        }
                                        fmt.Println(v5)
                                        _ = fi82
                                        _ = v83
                                    }
                                }
                                _ = v80
                            }
                        }
                        v67 += v67
                        var v87 float32 = (-2.25 / (1.5 + ((float32(v4) - (1e10 * 3.14159)) + (-2.25 / 1.0))))
                        _ = v87
                    }
                    _ = v79
                } else {
                    var v88 string = "foo bar"
                    _ = v88
                }
                var v89 int = 2147483647
                _ = v76
                _ = v89
            }
            _ = v75
        }
        _ = v67
    }
    v4 = 0.1
    fmt.Println(v4)
    if (0 == -74) {
        var v90 int32 = -1
        _ = len(v1)
        _ = v90
    } else {
        fmt.Println(v5)
    }
    var v91 uint8 = (v3 << ((65535) & 63))
    v3 = v3
    var v92 int64 = 9223372036854775807
    v93 := GenS12{F0: 4294967295, F1: -1, F2: -2147483648}
    if (-2147483648 < -1) {
        v91 = (v3 + (v3 + v3))
    }
    v4 = (float64((1e10 / (float32(v4) + float32(v4)))) + v4)
    v94 := GenS13{F0: -1, F1: ("foo bar" + "\n\t")}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        if (!true) {
            v3 = v3
            fmt.Println(v3)
            if true {
                var v95 uint = 288
                _ = v95
            } else {
                fmt.Println(v91, v91, v92)
                var v96 int = 1
                _ = v96
            }
        } else {
            for fi97 := 0; fi97 < 3; fi97++ {
                v4 = float64(float32(3.14159))
                var v98 bool = (2147483647 < -1)
                _ = fi97
                _ = v98
            }
            var v99 string = ""
            _ = v99
        }
    } else {
        fmt.Println(v3)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v91
    _ = v92
    _ = v93
    _ = v94
}
