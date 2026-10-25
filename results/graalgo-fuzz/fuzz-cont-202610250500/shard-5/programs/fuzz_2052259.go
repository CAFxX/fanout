package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 uint16; F1 uint64; F2 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 int64; F2 int32 }
type GenS7 struct { F0 string; F1 uint64; F2 int32 }
type GenS8 struct { F0 int16; F1 int8; F2 int16 }
type GenS9 struct { F0 int; F1 uint16; F2 uint64 }
type GenS10 struct { F0 int16; F1 bool; F2 bool }

func main() {
    var v1 uint16 = 439
    var v2 float64 = (0.1 + (1.5 - 1.5))
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    v5 := make(map[string]bool)
    fmt.Println(v2)
    switch 32768 {
    case 4294967295:
        v6 := GenS1{F0: 0}
        _ = v6
    case -296:
        if (42 >= 0) {
            v7 := []int{}
            _ = v7
        }
    }
    v1 = (v1 + 738)
    switch -1 {
    case 0:
        v2 = (3.14159 * v2)
    case -2147483648:
        if (2 <= 324) {
            var v8 uint32 = 335
            _ = v8
        }
    default:
        v1 = v1
    }
    v9 := v3[-9223372036854775808]
    fmt.Println(v2, v9)
    switch -100 {
    case -9223372036854775808:
        if v9 {
            v10 := []int{-1, 128, 0, 32767}
            if (42 >= 709) {
                var v11 uint16 = v1
                var v12 int16 = 1
                _ = v11
                _ = v12
            } else {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                v13 := GenS1{F0: -799}
                v9 = (0 < -1)
                _ = v13
            }
            if (42 >= -456) {
                for fi14 := 0; fi14 < 19; fi14++ {
                    for fi15 := 0; fi15 < 12; fi15++ {
                        v16 := GenS2{F0: 228, F1: "\n\t"}
                        var v17 int8 = -1
                        if v9 {
                            if v9 {
                                v18 := GenS3{F0: 1, F1: ("hello" + "a")}
                                switch (-(fi14)) {
                                case -9223372036854775808:
                                    v17 = 44
                                case 9223372036854775807:
                                    v19 := GenS1{F0: -9223372036854775808}
                                    fmt.Println(v9)
                                    _ = len(v3)
                                    var v20 int16 = 32767
                                    if true {
                                        fmt.Println(fi14)
                                    } else {
                                        v20 = -417
                                        delete(v5, "")
                                        switch fi14 {
                                        case 0:
                                            var v21 uint8 = 127
                                            _ = v21
                                        }
                                    }
                                    _ = v19
                                    _ = v20
                                case -1:
                                    if (v1 != (v1 >> ((2) & 63))) {
                                        var v22 string = (("hello" + "a") + "\n\t")
                                        var v23 int16 = 1
                                        if (v2 <= (-2.25 + ((0.1 * 0.0) + v2))) {
                                            v22 = ((v22 + ((v22 + ((v22 + (v22 + (v22 + v22))) + v22)) + ((("hello" + ((v22 + v22) + v22)) + v22) + "\n\t"))) + "foo bar")
                                            v9 = false
                                            _ = len(v3)
                                        } else {
                                            delete(v4, 694)
                                        }
                                        _ = v22
                                        _ = v23
                                    } else {
                                        v24 := GenS1{F0: -262}
                                        switch "hello" {
                                        case "a":
                                            switch "a" {
                                            case "":
                                                switch "hello" {
                                                case "\n\t":
                                                    var v25 int64 = 685
                                                    _ = v25
                                                case "a":
                                                    var v26 int = fi14
                                                    _ = v26
                                                case "hello":
                                                    switch ("" + "a") {
                                                    case "foo bar":
                                                        var v27 float32 = 1.5
                                                        _ = v27
                                                    case "a":
                                                        var v28 float32 = 0.1
                                                        _ = v28
                                                    default:
                                                        v29 := GenS1{F0: 9223372036854775807}
                                                        switch "foo bar" {
                                                        case "hello":
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                        case "foo bar":
                                                            v30 := GenS4{F0: (v1 | v1), F1: 624, F2: "a"}
                                                            _ = v30
                                                        default:
                                                            _ = len(v10)
                                                        }
                                                        _ = v29
                                                    }
                                                }
                                            }
                                        default:
                                            v5[("foo bar" + "\n\t")] = ((v2 < v2) && (true || v9))
                                        }
                                        _ = v24
                                    }
                                default:
                                    if ((!v9) && v9) {
                                        v31 := GenS4{F0: (v1 | 32767), F1: 69, F2: "a"}
                                        v9 = (-2.25 != 1e10)
                                        if false {
                                            v1 = 255
                                            var v32 uint = 4294967295
                                            fmt.Println(v9, v17, v1)
                                            v9 = v9
                                            _ = v32
                                        }
                                        fmt.Println(fi15, v2, v17)
                                        _ = v31
                                    } else {
                                        fmt.Println(fi15, v2, v17)
                                        if v9 {
                                            v33 := GenS5{F0: (v1 ^ v1)}
                                            _ = v33
                                        }
                                        if (2 != 4294967295) {
                                            fmt.Println(fi14)
                                        } else {
                                            switch (^(fi14)) {
                                            case 9223372036854775807:
                                                var v34 uint16 = (v1 ^ (v1 & v1))
                                                switch "foo bar" {
                                                case "\n\t":
                                                    fmt.Println(v1, v9)
                                                case "foo bar":
                                                    var v35 int64 = -467
                                                    _ = v35
                                                case "a":
                                                    fmt.Println(v34, v34, v2)
                                                }
                                                _ = v34
                                            case 597:
                                                var v36 string = (("a" + ((("" + ("hello" + "\n\t")) + "\n\t") + "a")) + "\n\t")
                                                _ = v36
                                            }
                                            v37 := GenS4{F0: 2, F1: 171, F2: "hello"}
                                            _ = v37
                                        }
                                        if len(v10) > 0 {
                                            v10[0] = 32767
                                        }
                                        if (1.5 <= float64(3.14159)) {
                                            var v38 int = fi15
                                            v10 = append(v10, (fi15 + fi15))
                                            v17 = v17
                                            _ = v38
                                        }
                                    }
                                }
                                if v9 {
                                    var v39 string = "foo bar"
                                    if len(v10) > 0 {
                                        v10[0] = fi14
                                    }
                                    if v9 {
                                        fmt.Println(v1, fi14)
                                    }
                                    if v9 {
                                        fmt.Println(v17, fi15, v17)
                                        v9 = v9
                                        if v9 {
                                            v3[(fi14 ^ 127)] = v9
                                        }
                                    }
                                    switch "a" {
                                    case "hello":
                                        var v40 uint8 = 31
                                        _ = v40
                                    case "":
                                        if (55 >= 2) {
                                            switch "foo bar" {
                                            case "":
                                                v17 = -1
                                            case "hello":
                                                if (!v9) {
                                                    var v41 uint64 = 0
                                                    _ = v41
                                                }
                                            case "a":
                                                v1 = v1
                                            }
                                            v4[2147483647] = ("hello" + v39)
                                            v42 := GenS6{F0: 1, F1: 1, F2: 257}
                                            _ = v42
                                        } else {
                                            v2 += v2
                                            var v43 int = -9223372036854775808
                                            _ = v43
                                        }
                                    }
                                    v2 -= 1e10
                                    if (-32768 < 1) {
                                        var v44 float64 = v2
                                        _ = v44
                                    } else {
                                        var v45 int64 = 42
                                        if ((^(v1)) <= 255) {
                                            v2 = 0.1
                                            if false {
                                                v5[(v39 + "hello")] = (65535 < 138)
                                            }
                                            fmt.Println(v17)
                                        } else {
                                            _ = len(v5)
                                            var v46 uint = 1
                                            v45 = (v45 ^ v45)
                                            if v9 {
                                                v39 = ""
                                                fmt.Println(fi14, v45, v45)
                                                var v47 float32 = (3.14159 / 1.5)
                                                _ = v47
                                            }
                                            _ = v46
                                        }
                                        _ = v45
                                    }
                                    _ = v39
                                }
                                _ = v18
                            } else {
                                v48 := GenS2{F0: 42, F1: ("hello" + "")}
                                _ = v48
                            }
                        } else {
                            if ((41 >= v17) || (-122 < -9223372036854775808)) {
                                var v49 int64 = 570
                                if len(v10) > 0 {
                                    v10[0] = fi15
                                }
                                _ = v49
                            }
                        }
                        fmt.Println(v2)
                        if true {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            fmt.Println(fi15)
                        } else {
                            switch fi15 {
                            case -1:
                                v50 := GenS3{F0: 696, F1: ""}
                                _ = v50
                            case -9223372036854775808:
                                var v51 int16 = -6
                                _ = v51
                            case 42:
                                var v52 int64 = -1
                                v17 = v17
                                _ = v52
                            default:
                                fmt.Println(v17, v1)
                            }
                            var v53 int8 = (v17 - v17)
                            var v54 bool = (true && v9)
                            _ = v53
                            _ = v54
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                }
                v55 := GenS4{F0: v1, F1: 691, F2: ("foo bar" + (("foo bar" + "\n\t") + "foo bar"))}
                _ = v55
            } else {
                delete(v4, 2)
                if (127 > 139) {
                    fmt.Println(v9, v9)
                }
                fmt.Println(v9, v9)
            }
            _ = v10
        }
    }
    var v56 uint = 83
    for fi57 := 0; fi57 < 19; fi57++ {
        v2 -= 0.1
        switch "a" {
        case "hello":
            v9 = v9
        case "":
            if false {
                fmt.Println(fi57)
                switch fi57 {
                case -2147483648:
                    var v58 int8 = 1
                    _ = v58
                case -181:
                    if true {
                        v2 += ((float64((1e10 - float32(v2))) / -2.25) * ((v2 + (0.1 / 0.1)) + -2.25))
                        v3[fi57] = ((v9 && v9) && true)
                    } else {
                        var v59 uint64 = 374
                        v60 := GenS1{F0: 487}
                        v61 := GenS5{F0: (v1 + v1)}
                        _ = v59
                        _ = v60
                        _ = v61
                    }
                case 1:
                    if ((1 == 0) && v9) {
                        v62 := []int{-9223372036854775808, (fi57 % ((fi57) | 1)), 746}
                        v1 = (v1 / (((v1 << ((645) & 63))) | 1))
                        switch (fi57 - (fi57 * -2147483648)) {
                        case -921:
                            for fi63 := 0; fi63 < 19; fi63++ {
                                v64 := GenS7{F0: "\n\t", F1: 969, F2: -1}
                                _ = fi63
                                _ = v64
                            }
                        case 4294967295:
                            var v65 int32 = 2147483647
                            _ = v65
                        case -1:
                            v62 = append(v62, fi57)
                        default:
                            if false {
                                var v66 uint = 255
                                for fi67 := 0; fi67 < 3; fi67++ {
                                    v66 -= v56
                                    if len(v62) >= 2 {
                                        v62 = v62[1:len(v62)]
                                    }
                                    var v68 uint16 = (v1 & (v1 % ((v1) | 1)))
                                    _ = fi67
                                    _ = v68
                                }
                                v9 = v9
                                _ = v66
                            }
                        }
                        _ = v62
                    } else {
                        v1 += 65535
                    }
                }
            }
        }
        _ = fi57
    }
    v69 := GenS5{F0: v1}
    v70 := GenS3{F0: -1, F1: "\n\t"}
    for fi71 := 0; fi71 < 14; fi71++ {
        switch (("" + ("\n\t" + ("a" + ("a" + ("hello" + ("\n\t" + "hello")))))) + "") {
        case "a":
            for fi72 := 0; fi72 < 20; fi72++ {
                var v73 uint = v56
                fmt.Println(v1)
                _ = fi72
                _ = v73
            }
        case "\n\t":
            for fi74 := 0; fi74 < 19; fi74++ {
                v56 = (^(v56))
                var v75 int32 = 2147483647
                v76 := []int{(fi71 * (fi74 ^ fi74)), (fi74 * fi71)}
                _ = fi74
                _ = v75
                _ = v76
            }
        default:
            v2 += v2
        }
        var v77 bool = (297 <= 1)
        _ = fi71
        _ = v77
    }
    var v78 uint32 = 440
    if v9 {
        var v79 float32 = 1e10
        switch "" {
        case "a":
            var v80 int32 = 1
            _ = v80
        case "foo bar":
            v81 := []int{128, 2147483647, -245, int(v1)}
            _ = v81
        case "":
            v82 := GenS8{F0: 949, F1: -40, F2: 42}
            _ = v82
        default:
            if ((-1 >= 100) || v9) {
                for fi83 := 0; fi83 < 2; fi83++ {
                    var v84 int8 = -128
                    _ = fi83
                    _ = v84
                }
                v4[v78] = "a"
                var v85 float32 = 1e10
                _ = v85
            } else {
                var v86 bool = (!false)
                var v87 uint64 = 18446744073709551615
                _ = v86
                _ = v87
            }
        }
        _ = v79
    } else {
        _ = len(v4)
        if (true || true) {
            v88 := GenS9{F0: 42, F1: v1, F2: 1}
            v9 = (("a" + "") != "\n\t")
            _ = v88
        }
    }
    var v89 int = 140
    v90 := GenS10{F0: 0, F1: v9, F2: (!v9)}
    for fi91 := 0; fi91 < 17; fi91++ {
        var v92 float32 = 0.1
        var v93 int32 = 148
        v94 := []int{(v89 + (fi91 >> ((v56) & 63)))}
        _ = fi91
        _ = v92
        _ = v93
        _ = v94
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("v78:", v78)
    fmt.Println("v89:", v89)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v69
    _ = v70
    _ = v78
    _ = v89
    _ = v9
    _ = v90
}
