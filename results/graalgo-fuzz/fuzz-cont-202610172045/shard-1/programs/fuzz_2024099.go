package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16; F1 string }
type GenS8 struct { F0 int16; F1 uint64; F2 string }
type GenS9 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 bool; F1 uint32 }
type GenS12 struct { F0 int16; F1 int64; F2 int64 }
type GenS13 struct { F0 int64 }

func main() {
    v1 := []string{"a", "\n\t", "foo bar"}
    var v2 float64 = 3.14159
    var v3 int16 = 1000
    v4 := []string{((("\n\t" + "foo bar") + "\n\t") + (("\n\t" + "a") + "")), (((("hello" + "foo bar") + (("foo bar" + "a") + "hello")) + ("foo bar" + "hello")) + "a")}
    var v5 uint = 2
    v2 = v2
    var v6 float64 = v2
    for fi7 := 0; fi7 < 12; fi7++ {
        if (!true) {
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 float64 = 3.14159
                _ = fi8
                _ = v9
            }
            v6 += ((1e10 - (1e10 - v6)) * v2)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            switch "" {
            case "foo bar":
                if (!(!false)) {
                    var v10 int8 = 0
                    fmt.Println(v10)
                    v3 = 0
                    _ = v10
                } else {
                    for fi11 := 0; fi11 < 18; fi11++ {
                        var v12 uint32 = 2
                        v3 += v3
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                        _ = fi11
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 19; fi13++ {
                        fmt.Println(v3, fi13, v6)
                        var v14 int8 = 100
                        _ = fi13
                        _ = v14
                    }
                }
            case "a":
                v15 := GenS1{F0: (!true), F1: ""}
                _ = v15
            default:
                for fi16 := 0; fi16 < 1; fi16++ {
                    switch "a" {
                    case "":
                        v5 -= 790
                    }
                    fmt.Println(v5)
                    _ = fi16
                }
            }
        }
        var v17 bool = ((1.5 * (v6 - 3.14159)) == ((((v6 * (v6 + (v6 - 1.5))) / (3.14159 / ((v6 + float64(1.5)) * float64((((((3.14159 - (1.5 + -2.25)) / 3.14159) + (float32(v2) + 1e10)) + -2.25) + float32((0.0 / (v6 / 1.0)))))))) / v6) * v2))
        v18 := GenS2{F0: 190, F1: 9223372036854775807}
        _ = fi7
        _ = v17
        _ = v18
    }
    if false {
        switch -135 {
        case 532:
            _ = len(v4)
        case 324:
            var v19 float32 = 1.5
            _ = v19
        case 32768:
            var v20 uint64 = 1
            _ = v20
        default:
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
        }
        switch 127 {
        case 4294967295:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 421:
            v21 := GenS3{F0: 128, F1: -78}
            _ = v21
        case 9223372036854775807:
            var v22 uint64 = 357
            _ = v22
        default:
            for fi23 := 0; fi23 < 12; fi23++ {
                v2 = v2
                var v24 int = fi23
                _ = fi23
                _ = v24
            }
        }
    } else {
        for fi25 := 0; fi25 < 12; fi25++ {
            v2 = 1e10
            _ = fi25
        }
        v6 = (v6 - v6)
    }
    v1 = append(v1, "a")
    if (!true) {
        if true {
            v3 -= (v3 | (v3 & 42))
            v6 = -2.25
            switch "a" {
            case "hello":
                v3 = v3
            default:
                var v26 uint64 = 538
                _ = v26
            }
        }
        var v27 uint32 = 2147483647
        _ = v27
    }
    v28 := GenS4{F0: 18446744073709551615}
    for fi29 := 0; fi29 < 8; fi29++ {
        var v30 bool = false
        _ = len(v4)
        fmt.Println(v3)
        _ = fi29
        _ = v30
    }
    if (!false) {
        v4 = append(v4, "a")
        for fi31 := 0; fi31 < 18; fi31++ {
            switch (((("foo bar" + ("" + "\n\t")) + "hello") + "foo bar") + ("\n\t" + "\n\t")) {
            case "\n\t":
                for fi32 := 0; fi32 < 2; fi32++ {
                    if (!true) {
                        fmt.Println(v2)
                        var v33 bool = ((false && true) || true)
                        _ = v33
                    } else {
                        v34 := GenS5{F0: 4294967295}
                        _ = v34
                    }
                    fmt.Println(v3, fi32)
                    v6 -= v6
                    _ = fi32
                }
            case "a":
                fmt.Println(fi31, v6)
            }
            if true {
                for fi35 := 0; fi35 < 20; fi35++ {
                    if (false && false) {
                        switch "foo bar" {
                        case "foo bar":
                            v36 := GenS4{F0: 162}
                            _ = v36
                        case "a":
                            if len(v4) > 0 {
                                v4[0] = "hello"
                            }
                        case "\n\t":
                            v3 = 1000
                            v37 := GenS6{F0: 42}
                            if false {
                                v2 = float64((((float32(float64(0.1)) / 1.0) + (1e10 + (1e10 / float32(((-2.25 - 0.0) / (v2 - 1.5)))))) * 0.0))
                                v38 := GenS7{F0: (v3 / ((v3) | 1)), F1: "foo bar"}
                                fmt.Println(fi31)
                                _ = v38
                            } else {
                                if (!true) {
                                    if ((v5 >> (((v5 - (^(v5)))) & 63)) > 18446744073709551615) {
                                        fmt.Println(v5, v6)
                                    }
                                    v6 += ((v6 + float64(1e10)) + 3.14159)
                                }
                                var v39 int64 = 0
                                if (3.14159 == -2.25) {
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v40 uint8 = 0
                                        _ = v40
                                    case "hello":
                                        v41 := GenS6{F0: 2}
                                        _ = v41
                                    default:
                                        switch 0 {
                                        case 4294967295:
                                            v2 = 0.0
                                        case 9223372036854775807:
                                            var v42 int = int(v39)
                                            _ = v42
                                        }
                                    }
                                    var v43 float64 = v6
                                    switch -441 {
                                    case 989:
                                        if (243 != 1) {
                                            if (false && (1 >= 48)) {
                                                v44 := GenS4{F0: 311}
                                                if (false && true) {
                                                    var v45 uint16 = 1
                                                    fmt.Println(fi31, v6)
                                                    _ = v45
                                                }
                                                _ = v44
                                            } else {
                                                var v46 uint = (v5 - v5)
                                                var v47 string = ""
                                                var v48 int64 = v39
                                                _ = v46
                                                _ = v47
                                                _ = v48
                                            }
                                        }
                                    case 42:
                                        if len(v1) > 0 {
                                            v1[0] = "\n\t"
                                        }
                                        switch fi35 {
                                        case 32767:
                                            v49 := GenS8{F0: (v3 - int16(v5)), F1: 683, F2: ("hello" + "")}
                                            _ = v49
                                        }
                                    default:
                                        v39 -= (v39 >> ((255) & 63))
                                        switch fi31 {
                                        case 784:
                                            var v50 uint32 = 0
                                            _ = v50
                                        case -1:
                                            var v51 uint16 = 190
                                            _ = v51
                                        case -9223372036854775808:
                                            v39 -= v39
                                        }
                                        var v52 int = (fi31 | int(v39))
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                        if false {
                                            switch ("hello" + (("a" + "") + (("a" + "a") + "a"))) {
                                            case "\n\t":
                                                switch int(v39) {
                                                case -771:
                                                    if true {
                                                        v53 := GenS9{F0: 495, F1: 2, F2: -89}
                                                        _ = v53
                                                    }
                                                    switch v52 {
                                                    case 32767:
                                                        var v54 bool = false
                                                        _ = v54
                                                    }
                                                    var v55 bool = ((fi31 + int(v5)) >= 0)
                                                    _ = v55
                                                case -9223372036854775808:
                                                    if (((false && false) && ((v3 + (-(v3))) == 32767)) || ((v39 * (v39 | v39)) != (^(v39)))) {
                                                        v39 -= 9223372036854775807
                                                        fmt.Println(v2, v5, fi35)
                                                    } else {
                                                        _ = len(v4)
                                                        var v56 float32 = 0.1
                                                        if (v6 > v6) {
                                                            var v57 uint8 = 128
                                                            var v58 float64 = v6
                                                            v59 := GenS2{F0: (v57 / (((v57 + v57)) | 1)), F1: (^(v39))}
                                                            _ = v57
                                                            _ = v58
                                                            _ = v59
                                                        } else {
                                                            var v60 uint8 = 0
                                                            v1 = append(v1, ("\n\t" + "foo bar"))
                                                            switch ("" + "foo bar") {
                                                            case "":
                                                                switch (fi31 / ((v52) | 1)) {
                                                                case -2147483648:
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        var v61 uint8 = (v60 << (((v5 << ((v5) & 63))) & 63))
                                                                        _ = v61
                                                                    }
                                                                default:
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                            case "hello":
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            case "a":
                                                                var v62 uint8 = 1
                                                                _ = v62
                                                            }
                                                            v3 = v3
                                                            _ = v60
                                                        }
                                                        _ = v56
                                                    }
                                                case 4294967295:
                                                    v6 -= v2
                                                    if (0 != 616) {
                                                        v3 += v3
                                                        v4 = append(v4, "hello")
                                                        var v63 bool = true
                                                        _ = v63
                                                    }
                                                    if (96 <= 4294967295) {
                                                        if false {
                                                            v1 = append(v1, "")
                                                        } else {
                                                            var v64 string = "\n\t"
                                                            v65 := GenS2{F0: 201, F1: v39}
                                                            v66 := GenS10{F0: (v5 / ((v5) | 1))}
                                                            _ = v64
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                        fmt.Println(fi31, v6, v39)
                                                        v67 := GenS11{F0: false, F1: 1}
                                                        var v68 int16 = (-(v3))
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                    var v69 int32 = 435
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v69 += 1
                                                    _ = v69
                                                default:
                                                    var v70 uint16 = 1
                                                    _ = v70
                                                }
                                            default:
                                                v3 = (v3 / (((v3 * v3)) | 1))
                                            }
                                            var v71 bool = false
                                            if (v71 && false) {
                                                v3 -= v3
                                            }
                                            _ = v71
                                        }
                                        _ = v52
                                    }
                                    _ = v43
                                } else {
                                    v3 -= (v3 | -1)
                                }
                                _ = v39
                            }
                            var v72 int32 = -163
                            _ = v37
                            _ = v72
                        default:
                            if (!(!true)) {
                                v6 += float64(-2.25)
                                var v73 int32 = 0
                                _ = v73
                            }
                        }
                    } else {
                        if true {
                            if false {
                                if true {
                                    if false {
                                        var v74 int64 = 1
                                        if (128 < 2) {
                                            var v75 int32 = int32(v74)
                                            fmt.Println(v74, v6, fi31)
                                            v76 := GenS12{F0: (v3 - (v3 * (^(v3)))), F1: (v74 >> ((v5) & 63)), F2: (v74 | v74)}
                                            var v77 uint32 = 1
                                            _ = v75
                                            _ = v76
                                            _ = v77
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v74
                                    }
                                    if false {
                                        if len(v1) > 0 {
                                            v1[0] = ("\n\t" + "")
                                        }
                                        if false {
                                            v78 := GenS7{F0: -32768, F1: "hello"}
                                            _ = v78
                                        } else {
                                            fmt.Println(v3, v2, v6)
                                        }
                                        if (42 != -128) {
                                            var v79 int64 = 1
                                            v2 -= v6
                                            if len(v1) > 0 {
                                                v1[0] = "hello"
                                            }
                                            _ = v79
                                        }
                                        var v80 uint = v5
                                        var v81 bool = true
                                        _ = v80
                                        _ = v81
                                    }
                                    var v82 uint64 = 42
                                    v3 += v3
                                    _ = v82
                                } else {
                                    v83 := GenS7{F0: -32768, F1: ""}
                                    v84 := GenS11{F0: (!false), F1: 65535}
                                    _ = v83
                                    _ = v84
                                }
                                var v85 string = "a"
                                _ = v85
                            }
                        }
                    }
                    fmt.Println(v5, v6)
                    var v86 int64 = -9223372036854775808
                    var v87 string = (("a" + "") + (((((("" + (("\n\t" + "foo bar") + "\n\t")) + ("" + "a")) + "\n\t") + "") + ("hello" + "foo bar")) + ("hello" + "")))
                    _ = fi35
                    _ = v86
                    _ = v87
                }
                var v88 uint = v5
                switch "\n\t" {
                case "foo bar":
                    switch (fi31 >> ((291) & 63)) {
                    case -2147483648:
                        var v89 string = "\n\t"
                        _ = v89
                    default:
                        fmt.Println(v5, v88, v6)
                    }
                case "a":
                    v4 = append(v4, (("foo bar" + ("hello" + "hello")) + ("foo bar" + (("foo bar" + "") + "hello"))))
                }
                _ = v88
            }
            _ = fi31
        }
    }
    v90 := GenS13{F0: -1}
    v1 = append(v1, "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v90
}
