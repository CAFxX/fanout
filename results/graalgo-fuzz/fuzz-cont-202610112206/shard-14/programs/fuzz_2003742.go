package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 uint64; F2 int32 }
type GenS5 struct { F0 int16; F1 string; F2 string }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int64; F1 int32; F2 uint }
type GenS8 struct { F0 uint32; F1 uint16 }
type GenS9 struct { F0 int64; F1 int16 }
type GenS10 struct { F0 string }

func main() {
    var v1 uint = 255
    var v2 int64 = 239
    var v3 float64 = (3.14159 / 1e10)
    fmt.Println(v3)
    for fi4 := 0; fi4 < 18; fi4++ {
        switch fi4 {
        case 32768:
            for fi5 := 0; fi5 < 14; fi5++ {
                v1 += 255
                v1 = v1
                _ = fi5
            }
        default:
            for fi6 := 0; fi6 < 13; fi6++ {
                if (!(430 > 0)) {
                    v3 = v3
                    v7 := []int{}
                    var v8 uint = v1
                    _ = v7
                    _ = v8
                }
                var v9 int = (fi4 % (((fi4 << ((v1) & 63))) | 1))
                fmt.Println(v3, v3)
                _ = fi6
                _ = v9
            }
        }
        _ = fi4
    }
    v3 += float64(3.14159)
    var v10 uint32 = 360
    v11 := []int{838, 2147483647, 773}
    v1 -= v1
    switch "\n\t" {
    case "hello":
        _ = len(v11)
    case "\n\t":
        var v12 int32 = 2147483647
        _ = v12
    }
    if (-32768 <= 42) {
        v1 = v1
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            if false {
                switch "foo bar" {
                case "hello":
                    if len(v11) > 0 {
                        v11[0] = 32768
                    }
                case "":
                    switch "hello" {
                    case "foo bar":
                        v1 -= v1
                    case "":
                        switch 0 {
                        case -429:
                            if false {
                                if ((v10 + 504) <= (v10 * v10)) {
                                    v13 := GenS1{F0: ("a" + "hello"), F1: uint8(v10), F2: v10}
                                    v14 := GenS2{F0: "hello"}
                                    var v15 int64 = (v2 & int64(v10))
                                    _ = v13
                                    _ = v14
                                    _ = v15
                                } else {
                                    v16 := GenS2{F0: "a"}
                                    var v17 int64 = (-(v2))
                                    v18 := GenS3{F0: "foo bar"}
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                }
                            }
                        case 32767:
                            var v19 int8 = 127
                            _ = v19
                        default:
                            var v20 float64 = (v3 / 3.14159)
                            _ = v20
                        }
                    case "\n\t":
                        if (1.5 > v3) {
                            switch 127 {
                            case -9223372036854775808:
                                var v21 int8 = 108
                                _ = v21
                            }
                            v3 += 0.1
                        }
                    default:
                        if (!true) {
                            for fi22 := 0; fi22 < 10; fi22++ {
                                v2 = int64(fi22)
                                fmt.Println(fi22, v3)
                                for fi23 := 0; fi23 < 11; fi23++ {
                                    v24 := GenS4{F0: 32767, F1: 2, F2: -2147483648}
                                    fmt.Println(fi23, v10, v1)
                                    _ = len(v11)
                                    if true {
                                        fmt.Println(v3)
                                    }
                                    _ = fi23
                                    _ = v24
                                }
                                _ = fi22
                            }
                        } else {
                            v10 += 2
                        }
                    }
                default:
                    if false {
                        for fi25 := 0; fi25 < 7; fi25++ {
                            switch -697 {
                            case 0:
                                if (231 > 42) {
                                    for fi26 := 0; fi26 < 14; fi26++ {
                                        if true {
                                            v2 = (^(v2))
                                            switch fi25 {
                                            case 128:
                                                var v27 uint64 = 18446744073709551615
                                                _ = v27
                                            case -2147483648:
                                                v11 = append(v11, (fi25 | (fi26 % (((fi26 | (^(fi26)))) | 1))))
                                            default:
                                                v28 := GenS2{F0: ((((("foo bar" + "\n\t") + "hello") + "a") + "\n\t") + "\n\t")}
                                                _ = v28
                                            }
                                            var v29 string = ""
                                            _ = v29
                                        } else {
                                            fmt.Println(fi25)
                                            fmt.Println(v10, v2, v10)
                                            var v30 float64 = 1.5
                                            if len(v11) > 0 {
                                                v11[0] = fi25
                                            }
                                            v31 := GenS5{F0: 42, F1: "\n\t", F2: (("" + "hello") + ("\n\t" + (("a" + "a") + (("\n\t" + ("\n\t" + ("" + ""))) + ("hello" + "a")))))}
                                            _ = v30
                                            _ = v31
                                        }
                                        if ("foo bar" > "foo bar") {
                                            var v32 int = 9223372036854775807
                                            v33 := GenS3{F0: "foo bar"}
                                            switch -91 {
                                            case 314:
                                                if false {
                                                    v32 = 1
                                                    v11 = append(v11, v32)
                                                } else {
                                                    v1 -= v1
                                                    var v34 uint64 = 1
                                                    _ = v34
                                                }
                                                if ((42 == 0) || (!false)) {
                                                    var v35 float64 = -2.25
                                                    var v36 int16 = 42
                                                    v10 += uint32(v1)
                                                    fmt.Println(fi26)
                                                    _ = v35
                                                    _ = v36
                                                } else {
                                                    v32 = v32
                                                }
                                                v37 := GenS6{F0: (v10 + 65535)}
                                                if len(v11) > 0 {
                                                    v11[0] = fi25
                                                }
                                                _ = v37
                                            }
                                            v38 := GenS6{F0: 105}
                                            v1 += (^(v1))
                                            v39 := GenS1{F0: "", F1: 128, F2: v10}
                                            _ = len(v11)
                                            v10 = v10
                                            if (!false) {
                                                v32 += -2147483648
                                                v40 := GenS1{F0: "a", F1: 1, F2: v10}
                                                fmt.Println(v1, v32, fi26)
                                                v3 = 0.0
                                                _ = v40
                                            } else {
                                                v1 += 42
                                            }
                                            v10 -= (^(v10))
                                            var v41 int16 = -32768
                                            _ = v32
                                            _ = v33
                                            _ = v38
                                            _ = v39
                                            _ = v41
                                        }
                                        _ = fi26
                                    }
                                }
                            case 32767:
                                v1 -= (^(v1))
                            case 445:
                                v42 := GenS7{F0: (v2 << ((v1) & 63)), F1: 1, F2: v1}
                                _ = v42
                            default:
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                            }
                            v3 = v3
                            var v43 uint = (^(v1))
                            _ = fi25
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 3; fi44++ {
                            v45 := GenS6{F0: v10}
                            var v46 uint32 = (v10 & v10)
                            switch (fi44 % ((0) | 1)) {
                            case -965:
                                var v47 float32 = float32(v3)
                                _ = v47
                            default:
                                if false {
                                    for fi48 := 0; fi48 < 9; fi48++ {
                                        var v49 int64 = v2
                                        _ = fi48
                                        _ = v49
                                    }
                                    var v50 uint16 = 255
                                    _ = v50
                                } else {
                                    switch fi44 {
                                    case 9223372036854775807:
                                        switch "hello" {
                                        case "a":
                                            if true {
                                                for fi51 := 0; fi51 < 5; fi51++ {
                                                    var v52 int64 = (v2 & (-(v2)))
                                                    if (0 != 0) {
                                                        v53 := GenS1{F0: "hello", F1: 0, F2: 2147483647}
                                                        v54 := GenS3{F0: "\n\t"}
                                                        _ = v53
                                                        _ = v54
                                                    } else {
                                                        v55 := GenS6{F0: 2147483647}
                                                        v11 = append(v11, fi51)
                                                        fmt.Println(v2)
                                                        _ = v55
                                                    }
                                                    _ = fi51
                                                    _ = v52
                                                }
                                            }
                                        default:
                                            _ = len(v11)
                                        }
                                    case -2147483648:
                                        fmt.Println(v46, v46, v3)
                                    case -43:
                                        fmt.Println(v3, v2, v3)
                                    }
                                    fmt.Println(v3, v1, v3)
                                    if false {
                                        _ = len(v11)
                                    } else {
                                        v11 = append(v11, int(v1))
                                        var v56 int8 = -5
                                        var v57 int16 = 42
                                        _ = v56
                                        _ = v57
                                    }
                                }
                            }
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                        var v58 uint64 = 208
                        _ = v58
                    }
                }
                for fi59 := 0; fi59 < 18; fi59++ {
                    if false {
                        v2 -= (v2 % ((v2) | 1))
                        if len(v11) > 0 {
                            v11[0] = 42
                        }
                    } else {
                        if (!(3.14159 != (((1.5 / (((((0.0 - 1e10) + ((0.0 / (3.14159 - -2.25)) - 0.1)) + float32(3.14159)) / (0.1 * (0.1 - (0.0 + (3.14159 + (1e10 - 1.5)))))) - 1e10)) * ((-2.25 * (((float32(v3) / 0.1) + 0.1) / 0.1)) * 3.14159)) - -2.25))) {
                            v1 = 65535
                            var v60 int64 = v2
                            var v61 uint32 = (v10 | 2147483648)
                            _ = v60
                            _ = v61
                        }
                    }
                    fmt.Println(v2, v10, v1)
                    switch 32767 {
                    case 326:
                        if false {
                            for fi62 := 0; fi62 < 19; fi62++ {
                                v2 = v2
                                _ = fi62
                            }
                        }
                    default:
                        var v63 int16 = int16(v2)
                        _ = v63
                    }
                    _ = fi59
                }
                v2 += v2
            } else {
                v1 = v1
                switch 42 {
                case -1:
                    if (v1 != (v1 & (v1 << ((v1) & 63)))) {
                        v64 := GenS3{F0: "hello"}
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        var v65 float32 = ((1.5 + 3.14159) - 0.0)
                        _ = v64
                        _ = v65
                    } else {
                        if true {
                            v10 += (v10 / (((v10 & v10)) | 1))
                            v3 += v3
                            if (true || (42 < 100)) {
                                v66 := GenS8{F0: v10, F1: 65535}
                                if false {
                                    v3 += v3
                                    if ((!false) || (!false)) {
                                        if len(v11) > 0 {
                                            v11[0] = 0
                                        }
                                    } else {
                                        if (!false) {
                                            for fi67 := 0; fi67 < 1; fi67++ {
                                                v2 -= 1
                                                v1 = v1
                                                _ = fi67
                                            }
                                            v3 = (v3 + (v3 / v3))
                                            var v68 string = "a"
                                            _ = v68
                                        } else {
                                            v69 := GenS9{F0: (^(v2)), F1: -1}
                                            if true {
                                                v70 := GenS10{F0: ("\n\t" + "")}
                                                if (true && (42 != 255)) {
                                                    _ = len(v11)
                                                    v1 -= v1
                                                }
                                                switch -9223372036854775808 {
                                                case 32768:
                                                    if len(v11) >= 2 {
                                                        v11 = v11[1:len(v11)]
                                                    }
                                                case 4294967295:
                                                    for fi71 := 0; fi71 < 12; fi71++ {
                                                        var v72 float64 = v3
                                                        var v73 int64 = 652
                                                        var v74 float32 = 1.5
                                                        _ = fi71
                                                        _ = v72
                                                        _ = v73
                                                        _ = v74
                                                    }
                                                case 128:
                                                    _ = len(v11)
                                                }
                                                _ = v70
                                            }
                                            for fi75 := 0; fi75 < 2; fi75++ {
                                                v3 = 1.5
                                                var v76 float64 = v3
                                                _ = fi75
                                                _ = v76
                                            }
                                            _ = v69
                                        }
                                        v77 := GenS10{F0: "hello"}
                                        var v78 int32 = -2147483648
                                        _ = v77
                                        _ = v78
                                    }
                                    v11 = append(v11, 1)
                                }
                                fmt.Println(v1, v10, v1)
                                _ = v66
                            }
                        }
                    }
                case 369:
                    v1 += 574
                case 1:
                    var v79 int32 = 1
                    _ = v79
                default:
                    var v80 float64 = v3
                    _ = v80
                }
                var v81 bool = false
                _ = v81
            }
        default:
            fmt.Println(v2, v2)
        }
    } else {
        var v82 int64 = (v2 ^ -742)
        fmt.Println(v3, v2, v1)
        _ = v82
    }
    var v83 uint8 = 22
    fmt.Println(v2, v10)
    v83 = (v83 * 41)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v83
}
