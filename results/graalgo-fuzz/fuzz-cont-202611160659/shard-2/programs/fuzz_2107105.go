package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 string }
type GenS2 struct { F0 uint32; F1 int; F2 int64 }
type GenS3 struct { F0 uint; F1 string; F2 uint16 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS6 struct { F0 bool; F1 int8 }
type GenS7 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS8 struct { F0 uint8; F1 int16; F2 int }

func main() {
    v1 := []int{409}
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 1
    v4 := make(map[int]string)
    v5 := make(map[string]bool)
    v6 := GenS1{F0: 2, F1: 0, F2: "\n\t"}
    for fi7 := 0; fi7 < 13; fi7++ {
        _ = len(v2)
        switch (fi7 ^ (fi7 * fi7)) {
        case -531:
            fmt.Println(fi7, fi7)
        default:
            switch (fi7 | fi7) {
            case -2147483648:
                var v8 uint16 = 32767
                _ = v8
            case 0:
                v9 := GenS2{F0: 2147483648, F1: (-(fi7)), F2: 0}
                _ = v9
            }
        }
        switch "" {
        case "\n\t":
            switch "foo bar" {
            case "hello":
                fmt.Println(fi7)
            case "foo bar":
                var v10 uint32 = 0
                _ = v10
            default:
                var v11 int64 = 450
                _ = v11
            }
        case "":
            var v12 int = (-(fi7))
            _ = v12
        case "hello":
            var v13 int32 = 384
            _ = v13
        default:
            switch fi7 {
            case 42:
                v14 := GenS1{F0: 0, F1: 42, F2: "a"}
                _ = v14
            case -793:
                switch fi7 {
                case 127:
                    v3 = (v3 ^ 0)
                case 32767:
                    for fi15 := 0; fi15 < 4; fi15++ {
                        if false {
                            switch ("foo bar" + ("a" + "hello")) {
                            case "hello":
                                v5[(("hello" + "") + "a")] = (0 != 9223372036854775807)
                            case "":
                                v3 -= 0
                            case "\n\t":
                                var v16 uint8 = 127
                                _ = v16
                            }
                            var v17 uint8 = v3
                            v17 = (^(v17))
                            _ = v17
                        }
                        _ = fi15
                    }
                default:
                    for fi18 := 0; fi18 < 8; fi18++ {
                        var v19 int8 = 97
                        _ = fi18
                        _ = v19
                    }
                }
            default:
                v3 = (v3 >> ((619) & 63))
            }
        }
        _ = fi7
    }
    v2[388] = 4294967295
    if len(v1) > 0 {
        v1[0] = -1
    }
    v1 = append(v1, 4294967295)
    if true {
        if true {
            fmt.Println(v3)
            if (!((1.5 < (3.14159 - (-2.25 - 0.1))) || true)) {
                v3 = v3
            } else {
                v3 -= 42
                var v20 uint = 423
                v21 := GenS3{F0: 761, F1: "\n\t", F2: 65535}
                _ = v20
                _ = v21
            }
            v3 = 127
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (42 > 4294967295) {
                if (1e10 != (((0.1 * 0.0) - 0.0) / 0.1)) {
                    if (-1 >= -1) {
                        v1 = append(v1, 724)
                    }
                    if len(v1) > 0 {
                        v1[0] = 127
                    }
                } else {
                    v22 := GenS2{F0: 2, F1: 32767, F2: -1}
                    _ = v22
                }
                for fi23 := 0; fi23 < 18; fi23++ {
                    if false {
                        v24 := GenS3{F0: 851, F1: "", F2: 331}
                        var v25 float32 = -2.25
                        _ = v24
                        _ = v25
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            v3 = v3
                        case "\n\t":
                            switch "a" {
                            case "hello":
                                v26 := GenS4{F0: 1000, F1: 18446744073709551615}
                                _ = v26
                            default:
                                for fi27 := 0; fi27 < 14; fi27++ {
                                    v3 -= (v3 & 225)
                                    fmt.Println(v3, fi27)
                                    var v28 int = fi27
                                    _ = fi27
                                    _ = v28
                                }
                            }
                        case "hello":
                            var v29 int64 = 1
                            _ = v29
                        default:
                            if false {
                                v30 := GenS5{F0: 1, F1: 42, F2: 241}
                                _ = v30
                            } else {
                                if (("\n\t" + "foo bar") == "hello") {
                                    v3 -= 252
                                    var v31 int64 = 1
                                    _ = v31
                                }
                                switch (fi23 & fi23) {
                                case 0:
                                    v3 = (v3 % (((^(v3))) | 1))
                                case 42:
                                    if false {
                                        if (42 >= 729) {
                                            v3 = v3
                                            fmt.Println(v3, v3)
                                        } else {
                                            var v32 int64 = -910
                                            _ = v32
                                        }
                                        switch (fi23 >> ((42) & 63)) {
                                        case 1:
                                            v3 = 33
                                        case 42:
                                            fmt.Println(v3, fi23)
                                        case 491:
                                            switch "foo bar" {
                                            case "":
                                                var v33 int64 = 42
                                                _ = v33
                                            case "hello":
                                                var v34 int32 = 2147483647
                                                _ = v34
                                            case "a":
                                                v1 = append(v1, (fi23 >> ((18446744073709551615) & 63)))
                                            }
                                        default:
                                            var v35 float32 = 3.14159
                                            _ = v35
                                        }
                                        _ = len(v4)
                                    } else {
                                        var v36 int64 = -1
                                        for fi37 := 0; fi37 < 11; fi37++ {
                                            var v38 int = 259
                                            var v39 string = "a"
                                            _ = fi37
                                            _ = v38
                                            _ = v39
                                        }
                                        switch "foo bar" {
                                        case "\n\t":
                                            v40 := GenS6{F0: false, F1: 82}
                                            _ = v40
                                        case "":
                                            fmt.Println(v36)
                                        }
                                        _ = v36
                                    }
                                case 128:
                                    v3 += v3
                                default:
                                    fmt.Println(fi23, fi23)
                                }
                                switch fi23 {
                                case -52:
                                    for fi41 := 0; fi41 < 11; fi41++ {
                                        v3 = v3
                                        if ((48 < 2) && (1.5 == 0.1)) {
                                            v3 -= v3
                                            v42 := GenS7{F0: 255, F1: 1, F2: -128}
                                            if ((!true) && (-1 <= -47)) {
                                                v43 := v4[(^(fi41))]
                                                if false {
                                                    var v44 int32 = 42
                                                    switch "a" {
                                                    case "\n\t":
                                                        var v45 string = v43
                                                        _ = v45
                                                    case "a":
                                                        if (2 == 255) {
                                                            if (!true) {
                                                                var v46 uint16 = 0
                                                                _ = v46
                                                            } else {
                                                                var v47 uint8 = (v3 | v3)
                                                                fmt.Println(fi23, fi23, v43)
                                                                fmt.Println(fi23, v44)
                                                                _ = v47
                                                            }
                                                            v44 += v44
                                                            if false {
                                                                v48 := GenS7{F0: 42, F1: uint16(v3), F2: 100}
                                                                _ = v48
                                                            } else {
                                                                var v49 int = fi41
                                                                v50 := GenS3{F0: 65535, F1: ((("\n\t" + v43) + "hello") + v43), F2: 81}
                                                                _ = v49
                                                                _ = v50
                                                            }
                                                            v44 -= -1
                                                            var v51 int32 = (v44 % (((-(v44))) | 1))
                                                            _ = v51
                                                        } else {
                                                            v2[704] = 572
                                                            v44 += (v44 % (((v44 & -2147483648)) | 1))
                                                            switch (v43 + "\n\t") {
                                                            case "hello":
                                                                if (false && false) {
                                                                    if ((float32((float64(float32((0.0 / 3.14159))) + 0.0)) < 1e10) && false) {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = fi23
                                                                        }
                                                                        fmt.Println(v3)
                                                                        var v52 string = "hello"
                                                                        _ = v52
                                                                    }
                                                                    v53 := v5[v43]
                                                                    v43 = v43
                                                                    _ = v53
                                                                } else {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = -2147483648
                                                                    }
                                                                    v43 = (v43 + v43)
                                                                    if len(v1) > 0 {
                                                                        v1[0] = fi23
                                                                    }
                                                                    switch v43 {
                                                                    case "":
                                                                        var v54 float64 = (((-2.25 + ((3.14159 / 0.1) + 1.5)) - 3.14159) - (float64((0.1 + 0.0)) * (float64(1e10) - -2.25)))
                                                                        _ = v54
                                                                    case "hello":
                                                                        v43 = v43
                                                                    default:
                                                                        switch (fi23 | (fi41 ^ fi23)) {
                                                                        case -9223372036854775808:
                                                                            v43 = "a"
                                                                        case 42:
                                                                            fmt.Println(fi41, v44)
                                                                        }
                                                                    }
                                                                    if (false || true) {
                                                                        var v55 int = 9223372036854775807
                                                                        var v56 string = (v43 + ("a" + v43))
                                                                        _ = v55
                                                                        _ = v56
                                                                    } else {
                                                                        var v57 int64 = 1
                                                                        v43 = v43
                                                                        _ = v57
                                                                    }
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    v44 = -448
                                                                }
                                                            }
                                                            switch ("" + v43) {
                                                            case "a":
                                                                var v58 int32 = 626
                                                                _ = v58
                                                            case "hello":
                                                                _ = len(v1)
                                                            }
                                                        }
                                                    }
                                                    _ = v44
                                                } else {
                                                    fmt.Println(v3)
                                                }
                                                var v59 int64 = 0
                                                _ = v43
                                                _ = v59
                                            }
                                            _ = v42
                                        } else {
                                            if false {
                                                v3 = v3
                                                fmt.Println(fi23, v3)
                                                delete(v2, 478)
                                                if false {
                                                    var v60 uint64 = 42
                                                    v61 := GenS4{F0: 42, F1: (v60 & v60)}
                                                    _ = v60
                                                    _ = v61
                                                } else {
                                                    v3 += v3
                                                }
                                            }
                                        }
                                        _ = fi41
                                    }
                                case 42:
                                    var v62 bool = true
                                    _ = v62
                                }
                            }
                        }
                        for fi63 := 0; fi63 < 1; fi63++ {
                            v3 = (v3 + 127)
                            _ = fi63
                        }
                        for fi64 := 0; fi64 < 5; fi64++ {
                            var v65 int8 = 87
                            v66 := GenS2{F0: 910, F1: (fi64 ^ 32767), F2: 408}
                            _ = fi64
                            _ = v65
                            _ = v66
                        }
                    }
                    switch "\n\t" {
                    case "hello":
                        var v67 bool = false
                        _ = v67
                    }
                    if false {
                        var v68 uint32 = 2147483648
                        v68 = (v68 | v68)
                        _ = v68
                    } else {
                        var v69 uint8 = v3
                        _ = v69
                    }
                    _ = fi23
                }
            }
            var v70 uint64 = 1
            _ = v70
        }
        fmt.Println(v3)
        var v71 float32 = -2.25
        _ = v71
    } else {
        for fi72 := 0; fi72 < 6; fi72++ {
            for fi73 := 0; fi73 < 10; fi73++ {
                v74 := GenS7{F0: 1, F1: 255, F2: 1}
                switch -606 {
                case 975:
                    var v75 bool = false
                    _ = v75
                case 707:
                    fmt.Println(v3, v3)
                default:
                    v5[(("a" + "foo bar") + "hello")] = true
                }
                v3 -= (v3 << ((542) & 63))
                _ = fi73
                _ = v74
            }
            v76 := GenS1{F0: uint64(v3), F1: 42, F2: "foo bar"}
            v3 = 42
            _ = fi72
            _ = v76
        }
    }
    v77 := GenS3{F0: 4294967295, F1: "foo bar", F2: 255}
    v78 := GenS8{F0: 0, F1: 1, F2: 32768}
    var v79 int16 = -37
    for fi80 := 0; fi80 < 10; fi80++ {
        v2[42] = 4294967295
        for fi81 := 0; fi81 < 3; fi81++ {
            v4[(fi81 << ((42) & 63))] = ("" + "")
            var v82 uint16 = 255
            var v83 int8 = 0
            _ = fi81
            _ = v82
            _ = v83
        }
        _ = len(v1)
        _ = fi80
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v79 = (v79 % (((v79 - v79)) | 1))
    var v84 int16 = (-(v79))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v79:", v79)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
    _ = v78
    _ = v79
    _ = v84
}
