package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 string; F1 int64; F2 uint64 }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint64; F1 string; F2 int8 }
type GenS11 struct { F0 int32; F1 int16; F2 uint }
type GenS12 struct { F0 int64 }

func main() {
    v1 := []int{}
    var v2 int = -1
    v3 := make(map[int]bool)
    var v4 uint32 = 2147483648
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            if (("a" + (("" + "hello") + ("hello" + ("foo bar" + ((("foo bar" + (("hello" + "a") + "foo bar")) + ("\n\t" + "foo bar")) + "hello"))))) < (("a" + "\n\t") + "hello")) {
                v7 := GenS1{F0: (v4 * 0), F1: (v4 ^ v4)}
                var v8 uint = 0
                var v9 int8 = 42
                _ = v7
                _ = v8
                _ = v9
            } else {
                v10 := GenS1{F0: 2147483648, F1: v4}
                var v11 float32 = -2.25
                var v12 uint8 = 255
                var v13 bool = ((!false) || (true && (18446744073709551615 == 796)))
                _ = v10
                _ = v11
                _ = v12
                _ = v13
            }
            if (100 == -128) {
                var v14 uint16 = 65535
                v15 := GenS1{F0: v4, F1: uint32(v14)}
                _ = v14
                _ = v15
            }
            v4 += 2147483648
            if (-2.25 >= (((3.14159 - (1.5 * (1e10 + 0.0))) * ((1e10 / 1.0) * float32(0.1))) * (0.1 * 1.5))) {
                var v16 float32 = 0.1
                var v17 int64 = 9223372036854775807
                if false {
                    if false {
                        v4 = v4
                        v17 += (-(v17))
                        fmt.Println(v4)
                    } else {
                        switch fi6 {
                        case 128:
                            switch fi5 {
                            case 42:
                                var v18 int = (-(fi6))
                                _ = v18
                            default:
                                v3[-851] = (127 >= (fi5 / (((v2 & -884)) | 1)))
                            }
                        case -9223372036854775808:
                            var v19 float64 = -2.25
                            v19 = ((((((((float64((v16 * (v16 + v16))) / 1e10) - v19) - -2.25) + 1.5) / 1.0) * v19) + v19) - v19)
                            _ = v19
                        default:
                            if ("" != "foo bar") {
                                v2 = (fi5 ^ (-(v2)))
                                v4 = (^(v4))
                                v2 -= -611
                            }
                        }
                        v4 = v4
                    }
                    switch v2 {
                    case -849:
                        v16 = (v16 * ((float32(float64(((0.0 / ((v16 / float32(1e10)) * v16)) - 0.1))) * (3.14159 - v16)) / 1e10))
                    case 32767:
                        v4 -= (v4 + (v4 ^ 2))
                    default:
                        if ("a" == "foo bar") {
                            fmt.Println(v17, fi5)
                            _ = len(v1)
                            v20 := GenS1{F0: (v4 * (^(v4))), F1: v4}
                            _ = v20
                        }
                    }
                    v16 = ((0.1 / float32((1.5 - (0.1 / (float64(1e10) - (-2.25 + 0.1)))))) + 3.14159)
                } else {
                    fmt.Println(fi6)
                    var v21 uint64 = 181
                    switch "" {
                    case "hello":
                        var v22 int32 = int32(v17)
                        v23 := GenS2{F0: ""}
                        switch "foo bar" {
                        case "\n\t":
                            v4 -= (v4 / ((2147483647) | 1))
                        default:
                            switch (^(fi6)) {
                            case 1:
                                v24 := GenS1{F0: 0, F1: uint32(v21)}
                                _ = v24
                            case 4294967295:
                                switch "\n\t" {
                                case "hello":
                                    var v25 uint32 = (v4 & 4294967295)
                                    _ = v25
                                }
                            case -1:
                                switch "a" {
                                case "":
                                    if true {
                                        var v26 string = "a"
                                        if (!(-32768 <= -1)) {
                                            v3[(fi5 ^ 4294967295)] = (v17 > v17)
                                            fmt.Println(v4)
                                            if true {
                                                _ = len(v1)
                                                if true {
                                                    v27 := GenS1{F0: v4, F1: v4}
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v28 float64 = 0.1
                                                        _ = v28
                                                    case "\n\t":
                                                        v4 -= (^(v4))
                                                    case "hello":
                                                        if len(v1) > 0 {
                                                            v1[0] = 0
                                                        }
                                                        if ((v16 * ((1e10 + 0.0) * 0.0)) > -2.25) {
                                                            v29 := GenS2{F0: v26}
                                                            v2 += fi6
                                                            _ = v29
                                                        } else {
                                                            var v30 int = -434
                                                            v22 = (v22 ^ 2147483647)
                                                            switch (v26 + (v26 + v26)) {
                                                            case "\n\t":
                                                                v31 := GenS3{F0: "a", F1: (v21 - v21), F2: v21}
                                                                _ = v31
                                                            case "":
                                                                switch v30 {
                                                                case 0:
                                                                    v32 := GenS1{F0: (v4 | (v4 ^ v4)), F1: v4}
                                                                    _ = v32
                                                                }
                                                            case "a":
                                                                fmt.Println(v21, v4)
                                                            }
                                                            _ = v30
                                                        }
                                                    }
                                                    _ = v27
                                                }
                                                v33 := GenS4{F0: 42, F1: 1}
                                                _ = v33
                                            }
                                        } else {
                                            var v34 int64 = v17
                                            fmt.Println(v17, fi6)
                                            v26 = v26
                                            v16 += 1.5
                                            v35 := GenS5{F0: 42, F1: -32768}
                                            _ = v34
                                            _ = v35
                                        }
                                        _ = v26
                                    } else {
                                        v16 = (v16 * v16)
                                        v36 := GenS6{F0: "", F1: 9223372036854775807, F2: (^(v21))}
                                        var v37 uint32 = 4294967295
                                        _ = v36
                                        _ = v37
                                    }
                                case "\n\t":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = v2
                                    }
                                case "foo bar":
                                    v17 += (v17 / (((v17 - 0)) | 1))
                                default:
                                    var v38 int = (^(fi6))
                                    _ = v38
                                }
                            default:
                                if true {
                                    v4 -= (v4 % ((418) | 1))
                                    _ = len(v3)
                                    v3[v2] = true
                                }
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = -136
                        }
                        _ = v22
                        _ = v23
                    case "":
                        v4 -= uint32(v21)
                    default:
                        fmt.Println(v4, v2, v2)
                    }
                    _ = v21
                }
                var v39 uint = 822
                var v40 uint8 = 128
                v2 = 16
                _ = v16
                _ = v17
                _ = v39
                _ = v40
            } else {
                switch 128 {
                case 32767:
                    if false {
                        fmt.Println(fi6)
                        if (4294967295 != 4294967295) {
                            var v41 float32 = (((0.1 / 1.0) * (float32(float64(0.1)) - -2.25)) + (1.5 / 0.1))
                            v42 := GenS7{F0: fi6, F1: (v4 & v4)}
                            _ = v41
                            _ = v42
                        }
                        fmt.Println(v4)
                    }
                case -723:
                    v2 = (fi6 ^ (fi5 >> ((0) & 63)))
                case 127:
                    switch (fi5 >> ((69) & 63)) {
                    case -387:
                        switch ("" + ("a" + "a")) {
                        case "hello":
                            v2 -= (-(fi5))
                        case "":
                            var v43 int32 = -81
                            _ = v43
                        default:
                            var v44 uint32 = (^(v4))
                            _ = v44
                        }
                    case 2147483647:
                        v45 := GenS2{F0: "hello"}
                        _ = v45
                    case 395:
                        var v46 bool = (!(84 != 127))
                        _ = v46
                    default:
                        var v47 float32 = (3.14159 + 0.0)
                        _ = v47
                    }
                default:
                    fmt.Println(v2)
                }
                v2 -= fi6
            }
            _ = fi6
        }
        _ = fi5
    }
    for fi48 := 0; fi48 < 11; fi48++ {
        v4 = (v4 - (v4 >> ((274) & 63)))
        if false {
            var v49 int8 = -41
            _ = v49
        } else {
            v4 += v4
            var v50 int32 = 848
            _ = v50
        }
        _ = fi48
    }
    if true {
        v2 = (-(v2))
        v4 = (v4 * v4)
    } else {
        fmt.Println(v2, v4)
        v2 -= (-(v2))
        if true {
            if false {
                switch (v2 & v2) {
                case -255:
                    v51 := GenS3{F0: ("a" + "hello"), F1: 1, F2: 2}
                    _ = v51
                default:
                    _ = len(v1)
                }
                switch "hello" {
                case "a":
                    v2 += 127
                }
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(v4, v4)
                case "hello":
                    switch "\n\t" {
                    case "foo bar":
                        var v52 uint64 = 42
                        _ = v52
                    }
                case "\n\t":
                    switch ("a" + "a") {
                    case "a":
                        for fi53 := 0; fi53 < 11; fi53++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v54 int8 = 42
                            _ = fi53
                            _ = v54
                        }
                    default:
                        var v55 int8 = 0
                        _ = v55
                    }
                }
            } else {
                fmt.Println(v4, v4)
            }
        }
    }
    var v56 int16 = -364
    if false {
        if false {
            switch "\n\t" {
            case "":
                fmt.Println(v2)
            case "foo bar":
                var v57 int32 = 0
                _ = v57
            }
        }
    }
    fmt.Println(v4, v4, v2)
    var v58 int = v2
    v59 := GenS8{F0: 0, F1: 42}
    for fi60 := 0; fi60 < 1; fi60++ {
        switch 128 {
        case 128:
            for fi61 := 0; fi61 < 19; fi61++ {
                var v62 int = (-(fi60))
                var v63 int16 = v56
                _ = len(v1)
                v64 := GenS9{F0: v63}
                if false {
                    if ((1 > 628) || ((0.0 == 0.0) && false)) {
                        fmt.Println(fi61, fi61, v4)
                    } else {
                        if true {
                            fmt.Println(v58, fi60, v4)
                            v2 -= fi60
                        } else {
                            var v65 bool = (!true)
                            if (2 > 0) {
                                v66 := GenS10{F0: 18446744073709551615, F1: ("\n\t" + "hello"), F2: -73}
                                fmt.Println(v56, v56, fi61)
                                v67 := v3[v62]
                                if (849 == -2147483648) {
                                    if v65 {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    } else {
                                        if v65 {
                                            if (0 < 42) {
                                                if len(v1) > 0 {
                                                    v1[0] = 2147483647
                                                }
                                                switch (-(v2)) {
                                                case -9223372036854775808:
                                                    var v68 bool = true
                                                    _ = v68
                                                case 32768:
                                                    v58 += v2
                                                default:
                                                    if v67 {
                                                        v2 -= 127
                                                        v4 += 65535
                                                        v58 = (-(v2))
                                                    }
                                                }
                                                var v69 float64 = 1e10
                                                _ = v69
                                            } else {
                                                v3[(v58 | (fi61 << ((0) & 63)))] = (0.0 != (((0.1 * 0.0) - 1e10) / (-2.25 * (1e10 * 1.5))))
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v70 bool = true
                                                _ = v70
                                            }
                                        }
                                        var v71 int8 = 127
                                        if v67 {
                                            v2 += (v58 % ((127) | 1))
                                            v67 = ((v2 << ((42) & 63)) != (v62 / (((^(v62))) | 1)))
                                        } else {
                                            var v72 float64 = 0.1
                                            var v73 int = fi61
                                            _ = v72
                                            _ = v73
                                        }
                                        _ = v71
                                    }
                                } else {
                                    _ = len(v1)
                                }
                                _ = v66
                                _ = v67
                            }
                            _ = v65
                        }
                        if false {
                            if len(v1) > 0 {
                                v1[0] = 127
                            }
                            fmt.Println(v62)
                        }
                    }
                    var v74 float64 = 1.5
                    v3[v62] = true
                    _ = v74
                }
                v75 := GenS11{F0: 42, F1: v63, F2: 0}
                var v76 float32 = 0.1
                switch (("a" + "\n\t") + ("foo bar" + ("" + (("a" + "a") + "foo bar")))) {
                case "hello":
                    v77 := v3[(fi60 | 128)]
                    _ = v77
                }
                _ = fi61
                _ = v62
                _ = v63
                _ = v64
                _ = v75
                _ = v76
            }
        case -1:
            if (!(0 <= 42)) {
                v78 := GenS6{F0: "\n\t", F1: 9223372036854775807, F2: 0}
                v56 = v56
                _ = v78
            }
        }
        for fi79 := 0; fi79 < 2; fi79++ {
            switch "hello" {
            case "a":
                var v80 float64 = 0.0
                _ = v80
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "foo bar":
                v81 := GenS12{F0: -25}
                _ = v81
            }
            _ = fi79
        }
        var v82 int64 = 67
        _ = fi60
        _ = v82
    }
    if true {
        var v83 float64 = float64((1e10 - 0.0))
        var v84 int8 = 42
        var v85 uint8 = uint8(v4)
        _ = v83
        _ = v84
        _ = v85
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v56
    _ = v58
    _ = v59
}
