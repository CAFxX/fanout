package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint64; F1 bool }
type GenS11 struct { F0 string; F1 int16; F2 int8 }
type GenS12 struct { F0 int64; F1 uint8; F2 int }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 string; F1 int32 }
type GenS15 struct { F0 string; F1 int }
type GenS16 struct { F0 uint64 }

func main() {
    var v1 int32 = -1
    v2 := make(map[int]string)
    v3 := []uint32{}
    var v4 uint32 = 531
    v5 := []string{"hello", ((("" + (("a" + "foo bar") + ("hello" + "foo bar"))) + "") + ("\n\t" + "hello"))}
    var v6 string = ("" + "a")
    switch 32768 {
    case 617:
        _ = len(v2)
    }
    switch "\n\t" {
    case "":
        v6 = v6
    case "hello":
        v4 = v4
    case "a":
        for fi7 := 0; fi7 < 18; fi7++ {
            if true {
                v2[fi7] = v6
            }
            var v8 int32 = v1
            for fi9 := 0; fi9 < 13; fi9++ {
                if (false || false) {
                    if false {
                        fmt.Println(fi9, v6, fi9)
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v6, fi9, fi9)
                        }
                        switch fi7 {
                        case 127:
                            if true {
                                if true {
                                    _ = len(v2)
                                    v8 = -37
                                } else {
                                    v10 := v2[fi7]
                                    _ = v10
                                }
                            } else {
                                v8 += (v8 >> ((537) & 63))
                                if true {
                                    if true {
                                        v11 := GenS1{F0: (fi9 ^ int(v8)), F1: 646, F2: 912}
                                        v12 := GenS2{F0: (v4 * v4), F1: 0, F2: (v4 << ((18446744073709551615) & 63))}
                                        _ = v11
                                        _ = v12
                                    } else {
                                        v13 := GenS3{F0: true, F1: 42}
                                        v6 = ("foo bar" + "a")
                                        v8 -= (-(v1))
                                        v14 := GenS4{F0: -957, F1: v4}
                                        switch "a" {
                                        case "":
                                            v8 = (v1 >> ((18446744073709551615) & 63))
                                        }
                                        v4 = (v4 + (v4 ^ v4))
                                        v4 += (v4 | (v4 | (v4 + 2147483647)))
                                        _ = v13
                                        _ = v14
                                    }
                                    v1 += (v8 * (v8 % ((v8) | 1)))
                                    v1 -= -158
                                    v15 := GenS5{F0: 254, F1: (v4 / ((v4) | 1))}
                                    _ = v15
                                } else {
                                    v5 = append(v5, v6)
                                    if false {
                                        v16 := GenS4{F0: 42, F1: v4}
                                        var v17 string = "a"
                                        if false {
                                            v3 = append(v3, v4)
                                            if (v1 < v1) {
                                                var v18 int64 = 1
                                                var v19 float64 = 1e10
                                                _ = v18
                                                _ = v19
                                            } else {
                                                if true {
                                                    var v20 int = (fi9 % (((fi9 / ((128) | 1))) | 1))
                                                    if false {
                                                        fmt.Println(v6)
                                                        v1 = (-(v1))
                                                        if len(v3) > 0 {
                                                            v3[0] = v4
                                                        }
                                                    } else {
                                                        var v21 int = -1
                                                        v1 -= -318
                                                        v22 := GenS6{F0: 0, F1: -744}
                                                        _ = v21
                                                        _ = v22
                                                    }
                                                    v1 = 42
                                                    _ = v20
                                                } else {
                                                    v23 := GenS7{F0: 42}
                                                    var v24 int64 = 0
                                                    var v25 int64 = (v24 * -526)
                                                    _ = v23
                                                    _ = v24
                                                    _ = v25
                                                }
                                                v2[-9223372036854775808] = v17
                                            }
                                            switch ("" + "a") {
                                            case "":
                                                v26 := GenS1{F0: fi9, F1: 0, F2: 32767}
                                                _ = v26
                                            }
                                            v27 := GenS1{F0: fi9, F1: 42, F2: 271}
                                            _ = v27
                                        } else {
                                            v28 := GenS4{F0: -32768, F1: v4}
                                            v29 := v2[(fi7 >> ((547) & 63))]
                                            _ = v28
                                            _ = v29
                                        }
                                        _ = v16
                                        _ = v17
                                    } else {
                                        v8 = v1
                                    }
                                }
                            }
                        case -116:
                            v6 = ((("a" + ((v6 + v6) + "a")) + "\n\t") + v6)
                        default:
                            switch v6 {
                            case "":
                                var v30 float64 = 1.5
                                _ = v30
                            case "foo bar":
                                v6 = ((("a" + v6) + "\n\t") + "a")
                            case "\n\t":
                                var v31 bool = ((^(fi9)) != fi7)
                                _ = v31
                            default:
                                v1 = (v8 >> ((65535) & 63))
                            }
                        }
                        switch v6 {
                        case "":
                            v32 := GenS2{F0: 62, F1: 4294967295, F2: 681}
                            _ = v32
                        }
                    } else {
                        fmt.Println(fi7)
                        v8 = v1
                        var v33 string = (("foo bar" + "\n\t") + (((v6 + v6) + (v6 + "hello")) + "hello"))
                        v8 = v1
                        if (42 <= (v4 & v4)) {
                            var v34 uint = 42
                            fmt.Println(v34, v34)
                            if ((1.5 + -2.25) != (0.1 * 1.5)) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v8 += v1
                            }
                            _ = v34
                        } else {
                            var v35 float32 = 0.1
                            v1 = 0
                            if (9223372036854775807 > -823) {
                                if (111 != 0) {
                                    var v36 uint64 = 346
                                    switch "" {
                                    case "hello":
                                        v37 := GenS8{F0: v4}
                                        _ = v37
                                    case "a":
                                        v38 := GenS9{F0: fi9}
                                        _ = v38
                                    default:
                                        fmt.Println(v6, v33, v6)
                                    }
                                    delete(v2, 127)
                                    v8 -= v8
                                    _ = v36
                                }
                            } else {
                                var v39 uint16 = 481
                                _ = v39
                            }
                            _ = v35
                        }
                        _ = v33
                    }
                }
                var v40 string = v6
                _ = fi9
                _ = v40
            }
            _ = fi7
            _ = v8
        }
    default:
        v1 = (v1 / ((v1) | 1))
    }
    if true {
        var v41 int16 = 1
        _ = v41
    } else {
        if false {
            var v42 float32 = 0.0
            _ = v42
        }
        if true {
            if (!false) {
                var v43 int8 = -15
                _ = v43
            }
            var v44 float64 = (-2.25 - (3.14159 + 0.0))
            v2[-2147483648] = ("foo bar" + (v6 + v6))
            _ = v44
        }
        for fi45 := 0; fi45 < 12; fi45++ {
            if true {
                switch (("foo bar" + v6) + "\n\t") {
                case "":
                    for fi46 := 0; fi46 < 4; fi46++ {
                        v47 := GenS10{F0: 1, F1: (((v6 < v6) || false) || ((-788 == -9223372036854775808) && (-285 == -1)))}
                        _ = fi46
                        _ = v47
                    }
                case "\n\t":
                    var v48 uint32 = v4
                    _ = v48
                }
            } else {
                v1 += (^(v1))
            }
            v49 := GenS3{F0: false, F1: 51}
            delete(v2, (fi45 % ((472) | 1)))
            _ = fi45
            _ = v49
        }
    }
    if ((3.14159 * 0.0) < (0.0 / ((((((0.0 + ((1e10 + 0.0) * -2.25)) / 1.0) / 3.14159) - (3.14159 * 0.1)) / float32(1.5)) / 3.14159))) {
        v1 -= 42
        var v50 int64 = 9223372036854775807
        _ = v50
    }
    v51 := GenS11{F0: "foo bar", F1: 1, F2: -116}
    if (true && false) {
        if (42 == 1) {
            v6 = v6
            switch "foo bar" {
            case "\n\t":
                for fi52 := 0; fi52 < 3; fi52++ {
                    for fi53 := 0; fi53 < 6; fi53++ {
                        if true {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        } else {
                            if (923 > 42) {
                                switch "\n\t" {
                                case "\n\t":
                                    var v54 float64 = 0.1
                                    _ = v54
                                default:
                                    var v55 int = fi53
                                    _ = v55
                                }
                                v4 = 517
                                var v56 int16 = 1000
                                _ = v56
                            }
                            if true {
                                var v57 float64 = 0.1
                                var v58 int32 = v1
                                if false {
                                    var v59 uint16 = 0
                                    _ = v59
                                }
                                v60 := GenS12{F0: -1, F1: 255, F2: (fi53 * fi53)}
                                _ = v57
                                _ = v58
                                _ = v60
                            }
                        }
                        fmt.Println(v1, v1)
                        v1 -= v1
                        _ = fi53
                    }
                    for fi61 := 0; fi61 < 9; fi61++ {
                        fmt.Println(v6)
                        switch (fi61 / (((^(fi52))) | 1)) {
                        case 42:
                            if false {
                                var v62 int16 = 570
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v63 := GenS13{F0: 2}
                                _ = v62
                                _ = v63
                            } else {
                                if false {
                                    if (!false) {
                                        var v64 int = (^(fi52))
                                        _ = v64
                                    } else {
                                        v65 := GenS5{F0: 127, F1: v4}
                                        if (!(float32((0.1 - 3.14159)) >= 1e10)) {
                                            var v66 uint8 = 42
                                            _ = len(v3)
                                            _ = v66
                                        } else {
                                            switch ((v6 + (v6 + ("hello" + (v6 + (v6 + v6))))) + "") {
                                            case "a":
                                                if true {
                                                    v67 := v2[(fi61 % (((fi52 - fi61)) | 1))]
                                                    if (((!true) && true) || (v4 != v4)) {
                                                        var v68 uint32 = 352
                                                        fmt.Println(v67)
                                                        var v69 uint16 = uint16(v68)
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                    v70 := GenS6{F0: 215, F1: fi52}
                                                    _ = v67
                                                    _ = v70
                                                } else {
                                                    v1 -= 42
                                                }
                                            default:
                                                var v71 int32 = 0
                                                _ = v71
                                            }
                                        }
                                        _ = v65
                                    }
                                    _ = len(v2)
                                    fmt.Println(v1, v1, v1)
                                }
                                if false {
                                    if true {
                                        if true {
                                            v4 = v4
                                            fmt.Println(v4, fi52)
                                        }
                                    } else {
                                        var v72 uint8 = 127
                                        var v73 int16 = -158
                                        var v74 string = ("a" + "a")
                                        if (!false) {
                                            if (!((v72 % ((v72) | 1)) == 1)) {
                                                var v75 int64 = -1
                                                var v76 uint8 = 42
                                                _ = v75
                                                _ = v76
                                            }
                                        } else {
                                            if false {
                                                switch v74 {
                                                case "a":
                                                    v2[127] = "\n\t"
                                                default:
                                                    fmt.Println(fi52, v72)
                                                }
                                                v73 = v73
                                            } else {
                                                fmt.Println(v74)
                                            }
                                        }
                                        _ = v72
                                        _ = v73
                                        _ = v74
                                    }
                                    v77 := GenS2{F0: 2147483648, F1: 2, F2: 4294967295}
                                    _ = v77
                                } else {
                                    _ = len(v5)
                                }
                                if true {
                                    v78 := GenS4{F0: 1000, F1: (v4 - 2147483648)}
                                    _ = v78
                                } else {
                                    if (3.14159 >= 1e10) {
                                        v79 := GenS13{F0: 834}
                                        var v80 float32 = (-2.25 + 1.5)
                                        _ = v79
                                        _ = v80
                                    } else {
                                        var v81 int64 = 0
                                        _ = v81
                                    }
                                    if false {
                                        switch ((("foo bar" + v6) + "\n\t") + (((v6 + "foo bar") + (v6 + (("foo bar" + v6) + v6))) + "")) {
                                        case "\n\t":
                                            _ = len(v3)
                                        default:
                                            v6 = v6
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = (v4 ^ (^(v4)))
                                        }
                                    } else {
                                        v4 -= 2
                                        v82 := GenS14{F0: v6, F1: v1}
                                        var v83 uint16 = 1
                                        var v84 float32 = 3.14159
                                        var v85 uint64 = 435
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                        _ = v85
                                    }
                                    switch (fi52 & fi52) {
                                    case 42:
                                        fmt.Println(fi61, v6, fi52)
                                    }
                                }
                                var v86 uint8 = 255
                                _ = v86
                            }
                        case 418:
                            if (2 <= 22) {
                                var v87 uint = 223
                                _ = v87
                            }
                        case -873:
                            var v88 uint8 = 164
                            _ = v88
                        }
                        _ = fi61
                    }
                    _ = fi52
                }
            case "":
                for fi89 := 0; fi89 < 4; fi89++ {
                    var v90 float64 = 0.1
                    v90 = v90
                    v91 := GenS15{F0: ("hello" + v6), F1: fi89}
                    _ = fi89
                    _ = v90
                    _ = v91
                }
            }
            switch 128 {
            case 226:
                v1 = (v1 * v1)
            default:
                fmt.Println(v6, v1, v4)
            }
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v5) > 0 {
                v5[0] = ((((v6 + v6) + ((v6 + v6) + ("a" + v6))) + v6) + ("foo bar" + (("\n\t" + (v6 + v6)) + v6)))
            }
            if true {
                var v92 int64 = 42
                _ = v92
            } else {
                var v93 uint8 = 241
                v94 := GenS16{F0: 1}
                fmt.Println(v4)
                _ = v93
                _ = v94
            }
        }
        v1 -= 1
    } else {
        if (false || false) {
            v6 = (v6 + "")
            var v95 uint = 65535
            v4 = v4
            _ = v95
        } else {
            var v96 uint64 = 1
            var v97 int8 = 17
            _ = v96
            _ = v97
        }
    }
    v6 = ((v6 + v6) + v6)
    var v98 uint64 = 465
    v99 := v2[518]
    fmt.Println(v1, v99, v98)
    v1 = v1
    v100 := v2[0]
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v98
    _ = v99
}
