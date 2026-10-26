package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 uint16; F2 bool }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 uint; F1 uint32; F2 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int16; F1 string; F2 uint32 }
type GenS10 struct { F0 int32; F1 string; F2 int8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint16; F1 int32; F2 int64 }
type GenS14 struct { F0 uint8; F1 string }
type GenS15 struct { F0 uint64; F1 int16 }
type GenS16 struct { F0 int32; F1 string; F2 uint }
type GenS17 struct { F0 int8; F1 uint }
type GenS18 struct { F0 uint16; F1 uint16 }
type GenS19 struct { F0 uint32; F1 uint32 }
type GenS20 struct { F0 uint32 }
func genFn0(p0 int16, p1 uint64, p2 uint32) uint {
    fmt.Println(p2, p0)
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 uint32 = (^(p2))
        p1 -= p1
        p0 = p0
        _ = fi4
        _ = v5
    }
    v6 := []int{769, 128, 986, 32767}
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            if (uint16(p2) != 65535) {
                if len(v6) > 0 {
                    v6[0] = 4294967295
                }
                if false {
                    var v9 uint64 = 0
                    _ = v9
                }
            } else {
                if (!(!false)) {
                    if false {
                        var v10 int32 = 259
                        fmt.Println(p0, fi7, p1)
                        switch "" {
                        case "foo bar":
                            var v11 uint8 = 42
                            var v12 int64 = int64(v10)
                            var v13 uint8 = v11
                            v14 := GenS1{F0: false, F1: int32(v13)}
                            _ = v11
                            _ = v12
                            _ = v13
                            _ = v14
                        case "":
                            if len(v6) > 0 {
                                v6[0] = 127
                            }
                        case "a":
                            var v15 int8 = 42
                            _ = v15
                        default:
                            var v16 uint8 = 0
                            _ = v16
                        }
                        _ = v10
                    } else {
                        v17 := GenS1{F0: true, F1: 0}
                        if false {
                            if (((!(0.1 == -2.25)) || ((p0 * (p0 << ((2) & 63))) > p0)) || false) {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if len(v6) > 0 {
                                    v6[0] = 32768
                                }
                                if false {
                                    var v18 float32 = 0.0
                                    if ("\n\t" >= (("\n\t" + "a") + ("" + "foo bar"))) {
                                        var v19 int64 = 697
                                        p1 = (^(p1))
                                        v6 = append(v6, fi8)
                                        _ = v19
                                    } else {
                                        switch ("a" + "a") {
                                        case "":
                                            if true {
                                                v6 = append(v6, fi7)
                                                if true {
                                                    var v20 int64 = 42
                                                    _ = v20
                                                } else {
                                                    switch "\n\t" {
                                                    case "":
                                                        var v21 int32 = 166
                                                        _ = v21
                                                    case "hello":
                                                        switch (-(fi8)) {
                                                        case 128:
                                                            _ = len(v6)
                                                        case -908:
                                                            v18 = v18
                                                        case 9223372036854775807:
                                                            p0 = (p0 / ((p0) | 1))
                                                        }
                                                    case "\n\t":
                                                        if (1e10 < 1.5) {
                                                            var v22 int32 = 140
                                                            p1 -= (^(p1))
                                                            _ = v22
                                                        } else {
                                                            if false {
                                                                if len(v6) > 0 {
                                                                    v6[0] = fi7
                                                                }
                                                                v23 := GenS2{F0: (p2 + p2)}
                                                                v6 = append(v6, (-(fi7)))
                                                                _ = v23
                                                            } else {
                                                                if false {
                                                                    fmt.Println(p0, v18, fi8)
                                                                } else {
                                                                    p2 = 42
                                                                    if (true && true) {
                                                                        v18 += float32(1e10)
                                                                        if (42 < 2) {
                                                                            if (42 < -991) {
                                                                                if len(v6) > 0 {
                                                                                    v6[0] = fi7
                                                                                }
                                                                                var v24 int = fi8
                                                                                if (992 == 42) {
                                                                                    var v25 uint = 654
                                                                                    _ = v25
                                                                                }
                                                                                if false {
                                                                                    if (!(!(6 < 100))) {
                                                                                        v18 += 0.0
                                                                                    } else {
                                                                                        fmt.Println(p0, fi7, fi8)
                                                                                    }
                                                                                }
                                                                                var v26 int8 = int8(p2)
                                                                                var v27 bool = true
                                                                                _ = v24
                                                                                _ = v26
                                                                                _ = v27
                                                                            } else {
                                                                                if false {
                                                                                    fmt.Println(v18, p1, p2)
                                                                                }
                                                                                if true {
                                                                                    var v28 uint16 = 682
                                                                                    p1 -= 2
                                                                                    var v29 float64 = (((3.14159 + 0.1) * 0.0) - (float64(1e10) / 1.5))
                                                                                    var v30 int = (fi7 + (^(fi7)))
                                                                                    _ = v28
                                                                                    _ = v29
                                                                                    _ = v30
                                                                                }
                                                                            }
                                                                            switch ("foo bar" + (("foo bar" + (("a" + "a") + "\n\t")) + (("hello" + ("foo bar" + "\n\t")) + "a"))) {
                                                                            case "hello":
                                                                                v18 -= ((v18 * v18) - 0.1)
                                                                            case "foo bar":
                                                                                if len(v6) > 0 {
                                                                                    v6[0] = 604
                                                                                }
                                                                            case "a":
                                                                                if (true && (!true)) {
                                                                                    _ = len(v6)
                                                                                } else {
                                                                                    v31 := GenS1{F0: (-2.25 == 3.14159), F1: 42}
                                                                                    var v32 int64 = 9223372036854775807
                                                                                    _ = v31
                                                                                    _ = v32
                                                                                }
                                                                            default:
                                                                                switch "a" {
                                                                                case "\n\t":
                                                                                    var v33 uint32 = p2
                                                                                    _ = v33
                                                                                case "a":
                                                                                    p0 = p0
                                                                                case "foo bar":
                                                                                    _ = len(v6)
                                                                                default:
                                                                                    v34 := GenS2{F0: p2}
                                                                                    _ = v34
                                                                                }
                                                                            }
                                                                        }
                                                                        if true {
                                                                            p1 = (p1 | 42)
                                                                            fmt.Println(p2)
                                                                            if ((-1 <= 42) || false) {
                                                                                v35 := GenS3{F0: true, F1: int32(p2)}
                                                                                fmt.Println(v18)
                                                                                if (fi7 < (fi7 - fi8)) {
                                                                                    v6 = append(v6, (fi7 / ((fi8) | 1)))
                                                                                    var v36 float32 = v18
                                                                                    fmt.Println(v18, v36)
                                                                                    _ = v36
                                                                                }
                                                                                _ = v35
                                                                            } else {
                                                                                switch "a" {
                                                                                case "a":
                                                                                    var v37 string = ""
                                                                                    _ = v37
                                                                                case "hello":
                                                                                    var v38 int32 = -1
                                                                                    _ = v38
                                                                                case "\n\t":
                                                                                    p0 = p0
                                                                                    if false {
                                                                                        v18 += 1.5
                                                                                        p1 += 0
                                                                                    }
                                                                                    if false {
                                                                                        switch -1 {
                                                                                        case 42:
                                                                                            switch fi7 {
                                                                                            case -1:
                                                                                                v18 = (v18 - v18)
                                                                                            case 32768:
                                                                                                if len(v6) >= 2 {
                                                                                                    v6 = v6[1:len(v6)]
                                                                                                }
                                                                                            case 127:
                                                                                                if len(v6) > 0 {
                                                                                                    v6[0] = (fi8 << ((255) & 63))
                                                                                                }
                                                                                            default:
                                                                                                if len(v6) > 0 {
                                                                                                    v6[0] = -419
                                                                                                }
                                                                                            }
                                                                                        case 128:
                                                                                            var v39 int16 = (-(p0))
                                                                                            p2 = (p2 & 65535)
                                                                                            _ = v39
                                                                                        }
                                                                                        fmt.Println(v18, v18, fi7)
                                                                                    } else {
                                                                                        var v40 uint32 = (p2 >> ((2) & 63))
                                                                                        v41 := GenS4{F0: (p2 << ((2) & 63))}
                                                                                        v6 = append(v6, (^(fi7)))
                                                                                        _ = v40
                                                                                        _ = v41
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                    if true {
                                                        var v42 int8 = -1
                                                        if true {
                                                            v18 = ((3.14159 / (v18 * v18)) * v18)
                                                        } else {
                                                            var v43 bool = true
                                                            var v44 uint = 42
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                        v45 := GenS1{F0: false, F1: 0}
                                                        _ = v42
                                                        _ = v45
                                                    } else {
                                                        v46 := GenS3{F0: true, F1: -2147483648}
                                                        _ = v46
                                                    }
                                                    var v47 string = "a"
                                                    _ = v47
                                                }
                                                var v48 float32 = v18
                                                _ = v48
                                            } else {
                                                v49 := GenS1{F0: ((!true) && true), F1: 0}
                                                _ = v49
                                            }
                                        }
                                    }
                                    fmt.Println(fi8, fi8)
                                    fmt.Println(p0)
                                    _ = v18
                                }
                                p2 -= (p2 / ((p2) | 1))
                                if false {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                } else {
                                    v50 := GenS4{F0: 2}
                                    if (1 >= 9223372036854775807) {
                                        if len(v6) > 0 {
                                            v6[0] = fi8
                                        }
                                        fmt.Println(fi8)
                                    } else {
                                        var v51 int8 = -1
                                        fmt.Println(p2, v51)
                                        p1 -= (p1 * (^(p1)))
                                        _ = v51
                                    }
                                    _ = v50
                                }
                            }
                            if true {
                                if false {
                                    if len(v6) > 0 {
                                        v6[0] = fi8
                                    }
                                    if ((((1.5 - 1.5) / float32(0.1)) + 3.14159) < 0.0) {
                                        _ = len(v6)
                                        var v52 uint64 = p1
                                        if (1 <= 0) {
                                            var v53 int64 = 9223372036854775807
                                            _ = v53
                                        }
                                        _ = v52
                                    }
                                }
                            } else {
                                p0 = (-(p0))
                                var v54 int8 = 0
                                var v55 int8 = v54
                                _ = v54
                                _ = v55
                            }
                        }
                        var v56 uint32 = (p2 ^ (p2 / ((42) | 1)))
                        _ = v17
                        _ = v56
                    }
                    p2 = 2147483647
                    if (645 > 48) {
                        var v57 uint = 259
                        v6 = append(v6, fi8)
                        _ = v57
                    } else {
                        var v58 uint = 789
                        if false {
                            var v59 int16 = 498
                            _ = v59
                        } else {
                            v58 = 592
                        }
                        if (2147483648 == 42) {
                            var v60 int8 = -55
                            if len(v6) > 0 {
                                v6[0] = fi7
                            }
                            _ = v60
                        } else {
                            if (p2 >= (p2 + 2147483648)) {
                                switch (("" + (("\n\t" + "a") + (("a" + "\n\t") + "a"))) + "hello") {
                                case "\n\t":
                                    var v61 uint = (v58 << ((255) & 63))
                                    _ = v61
                                case "a":
                                    p0 += p0
                                case "foo bar":
                                    p2 = 4294967295
                                default:
                                    p1 -= (p1 % (((p1 / ((uint64(p2)) | 1))) | 1))
                                }
                                var v62 uint8 = 255
                                var v63 float32 = 1.5
                                switch (fi7 & fi8) {
                                case 0:
                                    if true {
                                        if (!true) {
                                            fmt.Println(v63)
                                        } else {
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            v6 = append(v6, (^(fi8)))
                                            v64 := GenS2{F0: p2}
                                            if true {
                                                v65 := GenS3{F0: (0 >= -685), F1: 1}
                                                if false {
                                                    if (3.14159 > float32(float64(0.0))) {
                                                        var v66 uint16 = 65535
                                                        v62 -= 42
                                                        _ = v66
                                                    }
                                                } else {
                                                    if true {
                                                        var v67 int8 = -86
                                                        var v68 uint16 = 764
                                                        _ = v67
                                                        _ = v68
                                                    } else {
                                                        v58 = v58
                                                        fmt.Println(v58, p1, fi8)
                                                        switch ("foo bar" + ("hello" + (("foo bar" + "\n\t") + ""))) {
                                                        case "\n\t":
                                                            var v69 uint16 = 213
                                                            _ = v69
                                                        case "":
                                                            switch (^(fi7)) {
                                                            case 0:
                                                                if len(v6) > 0 {
                                                                    v6[0] = fi7
                                                                }
                                                            case -2147483648:
                                                                var v70 uint32 = p2
                                                                _ = v70
                                                            case 2147483647:
                                                                v58 = v58
                                                            default:
                                                                switch "" {
                                                                case "a":
                                                                    v71 := GenS5{F0: "hello", F1: 1, F2: ((v62 | (^(v62))) > v62)}
                                                                    _ = v71
                                                                }
                                                            }
                                                        case "hello":
                                                            if true {
                                                                var v72 int = fi8
                                                                var v73 int8 = -7
                                                                if false {
                                                                    var v74 uint64 = 437
                                                                    v75 := GenS3{F0: (1 != 39), F1: 42}
                                                                    fmt.Println(fi8)
                                                                    _ = v74
                                                                    _ = v75
                                                                } else {
                                                                    p0 = (^(p0))
                                                                    v58 = v58
                                                                }
                                                                _ = v72
                                                                _ = v73
                                                            } else {
                                                                v62 = v62
                                                                p0 += p0
                                                                v76 := GenS3{F0: (2 == (p1 - p1)), F1: -1}
                                                                fmt.Println(fi8, p0)
                                                                _ = v76
                                                            }
                                                        default:
                                                            v63 = ((float32(-2.25) - ((((v63 * v63) * 0.1) / 0.1) / 3.14159)) / (((v63 / v63) + float32((0.1 + 3.14159))) * v63))
                                                        }
                                                    }
                                                    v77 := GenS6{F0: (54 > 2), F1: (v58 & v58)}
                                                    if (!(!((v63 / v63) == (v63 - ((v63 - v63) + v63))))) {
                                                        if (!(((3.14159 - (-2.25 * float64(v63))) / (-2.25 * (-2.25 / ((((1.5 - (1e10 - ((0.1 * 1.5) / ((1e10 + 3.14159) + 1.5)))) + ((1e10 * (-2.25 - 3.14159)) - float64(v63))) - (float64(float32((((1.5 - float64(v63)) - 1.5) + float64(v63)))) - 0.1)) / 0.1)))) == (1e10 + (1e10 + -2.25)))) {
                                                            if (9223372036854775807 >= -767) {
                                                                if true {
                                                                    fmt.Println(p1)
                                                                } else {
                                                                    v78 := GenS4{F0: (p2 - 42)}
                                                                    _ = v78
                                                                }
                                                            }
                                                        } else {
                                                            var v79 int64 = 939
                                                            _ = v79
                                                        }
                                                    }
                                                    v63 = v63
                                                    _ = v77
                                                }
                                                p2 = (p2 + uint32(p1))
                                                _ = v65
                                            }
                                            v80 := GenS4{F0: 338}
                                            fmt.Println(v58)
                                            v81 := GenS2{F0: (p2 ^ (p2 - (p2 - p2)))}
                                            switch "foo bar" {
                                            case "\n\t":
                                                if true {
                                                    p2 = p2
                                                } else {
                                                    var v82 uint32 = (p2 / (((^(p2))) | 1))
                                                    fmt.Println(v58, p1, p2)
                                                    _ = v82
                                                }
                                            case "a":
                                                var v83 uint = v58
                                                _ = v83
                                            case "":
                                                if (-2.25 != (((((float64(v63) / float64((((((((v63 * 0.1) - 1e10) * (v63 / (float32(-2.25) * ((v63 + float32(1.5)) * (v63 * v63))))) / -2.25) / (v63 * -2.25)) / 0.1) + (((3.14159 / ((1.5 + float32(float64(v63))) - v63)) - 1e10) + ((v63 / 1.5) + 3.14159))))) + 0.1) - float64(float32((float64(v63) / (0.0 + ((-2.25 + (1.5 - -2.25)) / -2.25)))))) * 0.1) + 1.5)) {
                                                    p2 = (p2 ^ p2)
                                                    if len(v6) > 0 {
                                                        v6[0] = -9223372036854775808
                                                    }
                                                } else {
                                                    if false {
                                                        var v84 uint = 235
                                                        _ = v84
                                                    } else {
                                                        v85 := GenS4{F0: 2147483648}
                                                        v86 := GenS7{F0: v58, F1: 2, F2: -9223372036854775808}
                                                        _ = v85
                                                        _ = v86
                                                    }
                                                }
                                            }
                                            fmt.Println(fi7, p0)
                                            _ = v64
                                            _ = v80
                                            _ = v81
                                        }
                                    } else {
                                        var v87 uint32 = 65535
                                        if ((p1 % (((^(p1))) | 1)) < uint64(p2)) {
                                            var v88 int32 = 2147483647
                                            if (!(true || ((float64(v63) / ((0.0 * (float64(v63) - ((((0.0 - 0.1) - 3.14159) - float64(v63)) / 1.5))) - 0.0)) != -2.25))) {
                                                var v89 int16 = p0
                                                v90 := GenS2{F0: (v87 / ((v87) | 1))}
                                                if len(v6) > 0 {
                                                    v6[0] = fi8
                                                }
                                                _ = v89
                                                _ = v90
                                            }
                                            fmt.Println(v58, fi7)
                                            _ = v88
                                        } else {
                                            fmt.Println(p2, p1)
                                            switch "foo bar" {
                                            case "hello":
                                                v6 = append(v6, (fi8 & fi8))
                                            }
                                            v91 := GenS8{F0: 181}
                                            _ = v91
                                        }
                                        _ = v87
                                    }
                                case 42:
                                    v63 += 3.14159
                                    var v92 string = "\n\t"
                                    v93 := GenS5{F0: (v92 + (("" + "foo bar") + v92)), F1: 2, F2: false}
                                    _ = v92
                                    _ = v93
                                case 1:
                                    var v94 uint = 4294967295
                                    _ = v94
                                }
                                _ = v62
                                _ = v63
                            } else {
                                if len(v6) > 0 {
                                    v6[0] = fi7
                                }
                                switch "\n\t" {
                                case "a":
                                    _ = len(v6)
                                case "hello":
                                    p0 = -714
                                case "\n\t":
                                    switch (fi7 ^ (fi7 ^ fi7)) {
                                    case 32768:
                                        _ = len(v6)
                                    case 4294967295:
                                        switch (fi7 % (((fi7 >> ((v58) & 63))) | 1)) {
                                        case 9223372036854775807:
                                            var v95 int = 42
                                            _ = v95
                                        default:
                                            switch 127 {
                                            case -2147483648:
                                                var v96 int32 = -164
                                                _ = v96
                                            case 610:
                                                p1 = p1
                                            case 127:
                                                p0 = 42
                                            }
                                        }
                                    default:
                                        fmt.Println(fi8)
                                    }
                                default:
                                    var v97 uint64 = p1
                                    _ = v97
                                }
                            }
                            var v98 int64 = -1
                            v99 := GenS4{F0: (p2 / ((0) | 1))}
                            _ = v98
                            _ = v99
                        }
                        _ = v58
                    }
                    p0 = (p0 * p0)
                    if false {
                        switch "a" {
                        case "":
                            if (!(!(0 >= 0))) {
                                fmt.Println(p0, fi7, p2)
                                fmt.Println(p2, p0, p1)
                                p2 += 4294967295
                            } else {
                                v100 := GenS4{F0: p2}
                                v101 := GenS6{F0: false, F1: 974}
                                _ = v100
                                _ = v101
                            }
                        case "hello":
                            if true {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                fmt.Println(fi8)
                                v102 := GenS3{F0: true, F1: -72}
                                _ = v102
                            } else {
                                fmt.Println(fi7, p1, fi8)
                            }
                        }
                        switch (fi8 % ((int(p1)) | 1)) {
                        case 32768:
                            if ((!true) && false) {
                                p0 = (p0 & 1000)
                                p0 = (p0 * (p0 >> ((209) & 63)))
                                var v103 float32 = 0.0
                                _ = v103
                            }
                        case -1:
                            switch fi7 {
                            case 9223372036854775807:
                                if (127 <= 127) {
                                    if true {
                                        var v104 int8 = 100
                                        p1 -= 18446744073709551615
                                        v105 := GenS9{F0: 1000, F1: ("hello" + ("\n\t" + "\n\t")), F2: p2}
                                        var v106 string = "\n\t"
                                        switch ((v106 + (v106 + "\n\t")) + "foo bar") {
                                        case "a":
                                            var v107 bool = true
                                            _ = v107
                                        case "":
                                            var v108 int16 = int16(p2)
                                            _ = v108
                                        }
                                        var v109 int = (fi8 >> ((42) & 63))
                                        _ = v104
                                        _ = v105
                                        _ = v106
                                        _ = v109
                                    } else {
                                        var v110 string = (("" + (("foo bar" + (("a" + (("hello" + "\n\t") + "\n\t")) + "")) + "foo bar")) + "hello")
                                        fmt.Println(v110, fi7, v110)
                                        var v111 int64 = 74
                                        _ = v110
                                        _ = v111
                                    }
                                    p1 += p1
                                }
                            case 383:
                                _ = len(v6)
                            case -1:
                                if (201 != 4294967295) {
                                    var v112 string = "\n\t"
                                    _ = v112
                                }
                            default:
                                var v113 uint8 = 127
                                _ = v113
                            }
                            switch "a" {
                            case "a":
                                var v114 int32 = 211
                                _ = v114
                            case "foo bar":
                                var v115 bool = ((("\n\t" + "hello") + ("a" + "")) > ("a" + "foo bar"))
                                _ = v115
                            }
                        }
                    } else {
                        var v116 int16 = p0
                        var v117 uint32 = p2
                        _ = v116
                        _ = v117
                    }
                    var v118 uint8 = 128
                    switch (fi8 << ((18446744073709551615) & 63)) {
                    case 42:
                        p2 = (p2 * 1)
                    case -89:
                        var v119 int64 = -1
                        _ = v119
                    case -711:
                        fmt.Println(fi7, p0, p2)
                    default:
                        switch ("\n\t" + (("a" + ("\n\t" + "foo bar")) + "foo bar")) {
                        case "foo bar":
                            if (false || (true && ((0.1 * (-2.25 + (((0.1 * 1e10) - float32((-2.25 - 1e10))) / 3.14159))) < 3.14159))) {
                                if (true || true) {
                                    var v120 uint16 = 255
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    fmt.Println(fi8)
                                    _ = v120
                                }
                                if len(v6) > 0 {
                                    v6[0] = -9223372036854775808
                                }
                            } else {
                                fmt.Println(fi7, v118)
                                fmt.Println(fi7, fi7, fi7)
                            }
                        case "\n\t":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case "hello":
                            if true {
                                p0 += int16(p2)
                                fmt.Println(v118)
                            } else {
                                if true {
                                    var v121 int16 = p0
                                    _ = v121
                                }
                                fmt.Println(v118, p2, fi7)
                            }
                        }
                    }
                    switch (fi8 | (fi7 + 42)) {
                    case 32768:
                        var v122 uint16 = 0
                        _ = v122
                    case 127:
                        p1 -= 159
                    }
                    if true {
                        switch "a" {
                        case "hello":
                            var v123 bool = false
                            _ = v123
                        }
                        p1 -= p1
                        if (42 > 18446744073709551615) {
                            p2 = (p2 >> ((uint(p2)) & 63))
                        } else {
                            var v124 bool = true
                            if (false && v124) {
                                v124 = (1e10 > 0.1)
                            } else {
                                fmt.Println(p0, p2, fi8)
                                var v125 int8 = 127
                                v126 := GenS10{F0: 42, F1: (("hello" + "") + "a"), F2: v125}
                                _ = v125
                                _ = v126
                            }
                            switch "foo bar" {
                            case "\n\t":
                                var v127 uint32 = p2
                                _ = v127
                            }
                            _ = v124
                        }
                        switch ("a" + "foo bar") {
                        case "foo bar":
                            if true {
                                if len(v6) > 0 {
                                    v6[0] = fi8
                                }
                            } else {
                                if true {
                                    switch int(p2) {
                                    case 128:
                                        p2 = p2
                                    case 32768:
                                        fmt.Println(p0, v118)
                                        if (!false) {
                                            fmt.Println(p1, v118, p0)
                                            v118 = (v118 | (v118 | v118))
                                            if (int(v118) == (fi7 >> ((255) & 63))) {
                                                switch (fi7 % ((944) | 1)) {
                                                case 42:
                                                    v128 := GenS9{F0: 0, F1: "foo bar", F2: (^(p2))}
                                                    _ = v128
                                                case 128:
                                                    p1 = p1
                                                    fmt.Println(p2, p0, p2)
                                                }
                                                p2 = (p2 ^ 300)
                                                p1 += 0
                                                var v129 int32 = -2147483648
                                                _ = v129
                                            } else {
                                                p0 = p0
                                                if true {
                                                    var v130 float32 = ((-2.25 / ((1.5 / 1.0) + (float32(0.1) - float32(((0.0 - 1.5) - -2.25))))) / (1.5 + (1e10 + (1.5 / (0.1 / (((0.0 / 1e10) * 3.14159) - 1e10))))))
                                                    _ = v130
                                                }
                                                v131 := GenS11{F0: (-(fi7))}
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                _ = v131
                                            }
                                        }
                                    case 1:
                                        v132 := GenS1{F0: false, F1: -1}
                                        _ = v132
                                    default:
                                        v133 := GenS5{F0: (("\n\t" + ("" + "")) + "foo bar"), F1: uint16(v118), F2: false}
                                        _ = v133
                                    }
                                    v134 := GenS5{F0: ("" + ("hello" + ("\n\t" + "\n\t"))), F1: 661, F2: (int(p0) <= 4294967295)}
                                    fmt.Println(fi8)
                                    v118 = 42
                                    v135 := GenS1{F0: (-128 != 42), F1: 2147483647}
                                    _ = v134
                                    _ = v135
                                }
                                fmt.Println(p2)
                            }
                        case "":
                            v6 = append(v6, fi8)
                            fmt.Println(p0)
                            switch "\n\t" {
                            case "foo bar":
                                fmt.Println(p0, p1)
                            case "hello":
                                fmt.Println(p1, fi7, fi8)
                            }
                        default:
                            fmt.Println(fi7, v118)
                        }
                    }
                    switch fi7 {
                    case 607:
                        v118 += v118
                    case -835:
                        if false {
                            var v136 int64 = 706
                            var v137 int32 = 2147483647
                            var v138 uint8 = (v118 >> ((505) & 63))
                            if false {
                                var v139 string = "\n\t"
                                fmt.Println(p0, p1, fi8)
                                _ = v139
                            }
                            v140 := GenS11{F0: (fi7 | 202)}
                            _ = v136
                            _ = v137
                            _ = v138
                            _ = v140
                        }
                    }
                    _ = v118
                }
                var v141 uint32 = (^(p2))
                _ = v141
            }
            switch fi7 {
            case 4294967295:
                if false {
                    p2 += 2147483647
                    switch fi7 {
                    case -272:
                        switch (fi8 + fi7) {
                        case -357:
                            var v142 uint8 = 42
                            _ = v142
                        case -967:
                            v6 = append(v6, 592)
                        case 699:
                            _ = len(v6)
                        default:
                            var v143 int8 = 0
                            _ = v143
                        }
                    case -597:
                        if false {
                            p0 += p0
                            var v144 int16 = p0
                            _ = v144
                        } else {
                            if (!false) {
                                var v145 int = (-(fi7))
                                _ = v145
                            }
                            p2 -= (p2 / ((p2) | 1))
                        }
                    default:
                        var v146 uint = 152
                        _ = v146
                    }
                    var v147 int16 = p0
                    v147 = 32767
                    switch fi8 {
                    case 42:
                        var v148 string = ((("hello" + "a") + "") + "")
                        _ = v148
                    case -2147483648:
                        p0 -= 1
                    case 352:
                        if false {
                            p1 = p1
                            v149 := GenS12{F0: (p0 * 32767)}
                            p1 += (p1 ^ 42)
                            _ = v149
                        }
                    default:
                        fmt.Println(fi7, v147)
                        switch -363 {
                        case 32768:
                            fmt.Println(p0, p1)
                        case 0:
                            v150 := GenS13{F0: 168, F1: 1, F2: -1}
                            _ = v150
                        case 1:
                            v151 := GenS14{F0: 1, F1: ""}
                            _ = v151
                        default:
                            if true {
                                if (-764 <= 295) {
                                    v147 += (p0 + v147)
                                }
                            }
                        }
                    }
                    _ = v147
                }
            case -2147483648:
                v152 := GenS15{F0: uint64(p2), F1: (p0 / ((p0) | 1))}
                _ = v152
            }
            p1 -= p1
            p1 -= 921
            _ = fi8
        }
        var v153 int64 = 572
        _ = fi7
        _ = v153
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1(p0 uint16, p1 uint, p2 uint) uint64 {
    for fi154 := 0; fi154 < 15; fi154++ {
        var v155 int16 = 32767
        for fi156 := 0; fi156 < 12; fi156++ {
            fmt.Println(p0, fi154, p0)
            _ = fi156
        }
        _ = fi154
        _ = v155
    }
    if true {
        v157 := []int{127, 32768, 2147483647}
        if (2 == 128) {
            var v158 int64 = -1
            _ = v158
        }
        _ = v157
    } else {
        for fi159 := 0; fi159 < 1; fi159++ {
            v160 := GenS13{F0: p0, F1: -335, F2: -88}
            _ = fi159
            _ = v160
        }
        for fi161 := 0; fi161 < 19; fi161++ {
            if (!true) {
                var v162 int32 = -2147483648
                _ = v162
            } else {
                if (!false) {
                    p1 = p1
                }
                v163 := GenS16{F0: 1, F1: "", F2: (^(p1))}
                if ((p2 ^ (p1 % ((42) | 1))) <= (p2 & (p2 ^ p1))) {
                    v164 := GenS4{F0: 249}
                    switch 32768 {
                    case -1:
                        var v165 float32 = 1e10
                        _ = v165
                    case -9223372036854775808:
                        v166 := GenS17{F0: -88, F1: 320}
                        _ = v166
                    case 32767:
                        v167 := []int{}
                        _ = v167
                    }
                    _ = v164
                }
                _ = v163
            }
            for fi168 := 0; fi168 < 14; fi168++ {
                var v169 string = ("foo bar" + ((("" + "foo bar") + "foo bar") + ("a" + "a")))
                _ = fi168
                _ = v169
            }
            _ = fi161
        }
    }
    var v170 int16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v170
    return 1
}

func main() {
    var v1 uint16 = 42
    var v2 int64 = -9223372036854775808
    v3 := make(map[int]int)
    if false {
        v1 -= v1
        v1 += 255
    } else {
        switch "foo bar" {
        case "":
            v171 := GenS18{F0: (^(v1)), F1: 65535}
            _ = v171
        default:
            switch 9223372036854775807 {
            case -2147483648:
                v172 := GenS7{F0: 0, F1: 460, F2: 0}
                _ = v172
            case 2147483647:
                switch "\n\t" {
                case "foo bar":
                    var v173 uint = 42
                    _ = v173
                }
            case 54:
                v2 += v2
            default:
                if (41 < 1) {
                    if ((false && (!false)) || true) {
                        for fi174 := 0; fi174 < 3; fi174++ {
                            var v175 bool = (false || false)
                            v2 += (-(v2))
                            _ = fi174
                            _ = v175
                        }
                    } else {
                        var v176 float32 = -2.25
                        v177 := GenS8{F0: (v1 | (v1 & v1))}
                        for fi178 := 0; fi178 < 2; fi178++ {
                            var v179 bool = (0 <= 1)
                            v179 = v179
                            var v180 uint32 = 790
                            _ = fi178
                            _ = v179
                            _ = v180
                        }
                        _ = v176
                        _ = v177
                    }
                } else {
                    if (!true) {
                        if false {
                            var v181 int64 = v2
                            v181 = v2
                            _ = v181
                        }
                        for fi182 := 0; fi182 < 20; fi182++ {
                            for fi183 := 0; fi183 < 3; fi183++ {
                                if (!(2147483647 == 944)) {
                                    fmt.Println(v2, v1, v1)
                                    var v184 float32 = (0.0 * 1.5)
                                    switch (fi183 >> ((0) & 63)) {
                                    case 32767:
                                        switch (("hello" + "foo bar") + "a") {
                                        case "hello":
                                            var v185 bool = true
                                            _ = v185
                                        case "a":
                                            v184 = ((((v184 - v184) - v184) / (v184 + v184)) / 0.1)
                                        case "foo bar":
                                            var v186 int64 = 9223372036854775807
                                            _ = v186
                                        }
                                    case -9223372036854775808:
                                        v187 := GenS8{F0: v1}
                                        _ = v187
                                    }
                                    v1 = (^(v1))
                                    var v188 int64 = v2
                                    _ = v184
                                    _ = v188
                                } else {
                                    switch fi182 {
                                    case -1:
                                        v189 := GenS1{F0: (true || (genFn1(v1, 255, 18446744073709551615) > 18446744073709551615)), F1: 0}
                                        _ = v189
                                    case 128:
                                        var v190 float64 = 1.5
                                        _ = v190
                                    default:
                                        v191 := []int{}
                                        _ = v191
                                    }
                                }
                                v192 := []int{(fi183 >> ((18446744073709551615) & 63)), fi182, (^(fi183)), fi183}
                                _ = fi183
                                _ = v192
                            }
                            var v193 int64 = -1
                            _ = fi182
                            _ = v193
                        }
                        fmt.Println(v2)
                    }
                    if true {
                        var v194 bool = (true || ((487 <= -2147483648) || (!(32768 > 0))))
                        if (127 < 197) {
                            fmt.Println(v194)
                            v195 := GenS16{F0: 0, F1: ("" + "foo bar"), F2: 882}
                            _ = v195
                        }
                        for fi196 := 0; fi196 < 12; fi196++ {
                            if true {
                                for fi197 := 0; fi197 < 11; fi197++ {
                                    var v198 uint32 = 2
                                    var v199 uint8 = uint8(v1)
                                    if ((^(v198)) != uint32(v1)) {
                                        fmt.Println(fi196)
                                        v2 -= v2
                                        v199 = (v199 << ((255) & 63))
                                    } else {
                                        v200 := v3[(fi196 >> ((1) & 63))]
                                        _ = v200
                                    }
                                    _ = fi197
                                    _ = v198
                                    _ = v199
                                }
                                v2 = v2
                            } else {
                                v201 := GenS9{F0: 1000, F1: "\n\t", F2: 1}
                                fmt.Println(v2, v194)
                                _ = v201
                            }
                            if (true && v194) {
                                fmt.Println(v1)
                                v2 += (-(v2))
                                v202 := []int{(fi196 + (fi196 ^ (fi196 - (fi196 / (((fi196 + fi196)) | 1))))), -560}
                                _ = v202
                            }
                            _ = fi196
                        }
                        _ = v194
                    }
                }
            }
        }
        v203 := GenS19{F0: 2147483648, F1: 2}
        _ = v203
    }
    v204 := []int{-9223372036854775808}
    if (42 > 127) {
        v205 := GenS1{F0: false, F1: -808}
        if true {
            if (true && true) {
                v2 += v2
            } else {
                v206 := GenS20{F0: 484}
                _ = v206
            }
            var v207 float64 = float64(1.5)
            _ = v207
        }
        _ = v205
    } else {
        v2 += (v2 << ((4294967295) & 63))
        v208 := GenS8{F0: 427}
        _ = v208
    }
    var v209 int32 = -2147483648
    var v210 bool = true
    var v211 uint16 = v1
    for fi212 := 0; fi212 < 9; fi212++ {
        v213 := GenS5{F0: "", F1: v211, F2: (!(-63 >= 114))}
        _ = fi212
        _ = v213
    }
    v214 := GenS12{F0: 1000}
    delete(v3, -1)
    switch "foo bar" {
    case "a":
        for fi215 := 0; fi215 < 17; fi215++ {
            var v216 float32 = -2.25
            for fi217 := 0; fi217 < 11; fi217++ {
                fmt.Println(v216, v210, fi217)
                v3[fi215] = 42
                v2 = (v2 << ((255) & 63))
                _ = fi217
            }
            _ = fi215
            _ = v216
        }
    case "hello":
        for fi218 := 0; fi218 < 14; fi218++ {
            v219 := GenS3{F0: (!false), F1: 979}
            v220 := v3[fi218]
            var v221 bool = (42 < genFn1(32767, 1, 531))
            _ = fi218
            _ = v219
            _ = v220
            _ = v221
        }
    case "":
        if (-32768 >= 32767) {
            fmt.Println(v211, v1, v211)
            v211 = (v1 / ((92) | 1))
        } else {
            v209 = 2147483647
        }
    default:
        var v222 string = "hello"
        _ = v222
    }
    for fi223 := 0; fi223 < 2; fi223++ {
        for fi224 := 0; fi224 < 7; fi224++ {
            v210 = ((v209 % ((v209) | 1)) != 490)
            fmt.Println(v209)
            _ = fi224
        }
        _ = fi223
    }
    v209 = 0
    if len(v204) > 0 {
        v204[0] = 195
    }
    v209 = v209
    fmt.Println(v2, v210, v1)
    switch "\n\t" {
    case "":
        var v225 uint32 = 976
        _ = v225
    }
    var v226 int = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v204:", len(v204), v204)
    fmt.Println("v209:", v209)
    fmt.Println("v210:", v210)
    fmt.Println("v211:", v211)
    fmt.Println("v226:", v226)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v204
    _ = v209
    _ = v210
    _ = v211
    _ = v214
    _ = v226
    _ = v3
}
