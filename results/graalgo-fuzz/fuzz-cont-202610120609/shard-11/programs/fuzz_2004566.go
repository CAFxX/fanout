package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS3 struct { F0 int32; F1 uint16; F2 int64 }
type GenS4 struct { F0 int16; F1 int64; F2 int32 }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int8 }

func main() {
    var v1 float64 = (((((((((3.14159 / 0.1) - 0.0) + 1.5) + 1e10) * 0.0) / 1e10) - 0.1) + 0.1) + (1.5 / 1.5))
    v2 := []string{("hello" + "a")}
    var v3 int = -9223372036854775808
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]string)
    v6 := make(map[uint32]bool)
    var v7 uint8 = 141
    var v8 uint64 = 42
    v8 -= v8
    _ = len(v2)
    var v9 float32 = 0.1
    v10 := GenS1{F0: -614, F1: 1}
    var v11 string = "hello"
    if true {
        var v12 int8 = 122
        switch v11 {
        case "a":
            if (v11 < v11) {
                var v13 int8 = (v12 & v12)
                for fi14 := 0; fi14 < 8; fi14++ {
                    for fi15 := 0; fi15 < 2; fi15++ {
                        var v16 string = v11
                        _ = fi15
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 7; fi17++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi17
                    }
                    v12 = 100
                    _ = fi14
                }
                _ = v13
            } else {
                for fi18 := 0; fi18 < 2; fi18++ {
                    switch "hello" {
                    case "":
                        if (((v7 % (((v7 << ((2) & 63))) | 1)) < (^(v7))) || true) {
                            if (!(!(v7 == uint8(v8)))) {
                                v11 = "foo bar"
                                v19 := v5[v11]
                                _ = v19
                            } else {
                                if false {
                                    var v20 uint16 = 2
                                    for fi21 := 0; fi21 < 7; fi21++ {
                                        v1 = v1
                                        if false {
                                            var v22 uint32 = 509
                                            var v23 string = (v11 + v11)
                                            v24 := GenS1{F0: 9223372036854775807, F1: 42}
                                            _ = v22
                                            _ = v23
                                            _ = v24
                                        }
                                        _ = fi21
                                    }
                                    if (49 >= 145) {
                                        v20 = v20
                                        var v25 uint64 = 42
                                        _ = v25
                                    }
                                    _ = v20
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                var v26 uint32 = 2147483647
                                _ = v26
                            }
                        }
                    case "a":
                        for fi27 := 0; fi27 < 10; fi27++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v28 uint8 = v7
                            _ = fi27
                            _ = v28
                        }
                    case "hello":
                        v7 += 23
                    default:
                        if false {
                            for fi29 := 0; fi29 < 8; fi29++ {
                                fmt.Println(fi18, v7)
                                _ = fi29
                            }
                            if len(v2) > 0 {
                                v2[0] = ((((v11 + v11) + v11) + "\n\t") + "hello")
                            }
                        } else {
                            if false {
                                v30 := GenS2{F0: (^(v8)), F1: (v12 + v12), F2: 975}
                                var v31 string = v11
                                _ = v30
                                _ = v31
                            }
                            if true {
                                v9 = 1.5
                                v32 := GenS2{F0: 4294967295, F1: (v12 >> ((42) & 63)), F2: 4294967295}
                                for fi33 := 0; fi33 < 16; fi33++ {
                                    var v34 uint32 = 1
                                    var v35 uint64 = uint64(v7)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v36 uint = 50
                                    _ = fi33
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                                _ = v32
                            } else {
                                v1 += ((v1 * v1) + v1)
                                if len(v2) > 0 {
                                    v2[0] = v11
                                }
                                switch fi18 {
                                case 128:
                                    v37 := GenS1{F0: 42, F1: -9223372036854775808}
                                    _ = v37
                                case -870:
                                    fmt.Println(v9, v9, v12)
                                case 857:
                                    switch v11 {
                                    case "foo bar":
                                        if false {
                                            v11 = ""
                                            v9 -= (v9 * v9)
                                        } else {
                                            v8 -= 726
                                            var v38 int8 = (v12 * 0)
                                            _ = v38
                                        }
                                    }
                                default:
                                    switch "\n\t" {
                                    case "a":
                                        for fi39 := 0; fi39 < 8; fi39++ {
                                            fmt.Println(v12, v11, v1)
                                            switch v11 {
                                            case "hello":
                                                switch v3 {
                                                case 32768:
                                                    var v40 string = (v11 + v11)
                                                    _ = v40
                                                default:
                                                    v41 := GenS2{F0: v8, F1: v12, F2: 1}
                                                    _ = v41
                                                }
                                            case "":
                                                if (207 >= 42) {
                                                    fmt.Println(v3, v11)
                                                    switch v3 {
                                                    case -154:
                                                        v6[634] = ((^(v12)) > (v12 | v12))
                                                    case 298:
                                                        switch "" {
                                                        case "foo bar":
                                                            var v42 int = (-(fi18))
                                                            _ = v42
                                                        }
                                                    }
                                                    v43 := v5[v11]
                                                    v9 = (v9 + (v9 + v9))
                                                    _ = v43
                                                }
                                            }
                                            _ = fi39
                                        }
                                    case "hello":
                                        _ = len(v6)
                                    case "":
                                        v9 -= v9
                                    }
                                }
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                    v12 = v12
                    v8 += (v8 + v8)
                    _ = fi18
                }
                fmt.Println(v9, v11)
                var v44 int = (v3 - 4294967295)
                _ = v44
            }
        }
        _ = v12
    } else {
        fmt.Println(v8, v8, v8)
        v2 = append(v2, "\n\t")
        v9 = v9
    }
    v7 = (v7 + (v7 ^ 60))
    switch (v3 - v3) {
    case 4294967295:
        var v45 string = "\n\t"
        _ = v45
    case -170:
        fmt.Println(v8, v1)
    }
    switch 436 {
    case -2147483648:
        v7 += (v7 / (((v7 * (^(v7)))) | 1))
    default:
        var v46 int32 = 0
        _ = v46
    }
    switch "" {
    case "hello":
        var v47 float32 = (((v9 + 1e10) + v9) - 1.5)
        _ = v47
    case "foo bar":
        for fi48 := 0; fi48 < 13; fi48++ {
            fmt.Println(v8, v8)
            _ = fi48
        }
    }
    for fi49 := 0; fi49 < 4; fi49++ {
        for fi50 := 0; fi50 < 9; fi50++ {
            fmt.Println(v11, fi50, fi50)
            v3 = (fi49 << ((57) & 63))
            if (371 == 379) {
                if true {
                    v4[923] = 2
                    v1 -= v1
                    if false {
                        switch "" {
                        case "\n\t":
                            fmt.Println(v1, v1, v7)
                        case "hello":
                            v1 -= v1
                        default:
                            if false {
                                var v51 int64 = 973
                                switch v11 {
                                case "a":
                                    var v52 int8 = 1
                                    _ = v52
                                default:
                                    switch (((((v11 + v11) + "\n\t") + v11) + "") + v11) {
                                    case "":
                                        v53 := GenS2{F0: v8, F1: 1, F2: 2147483648}
                                        fmt.Println(fi49, fi49)
                                        _ = v53
                                    case "a":
                                        switch v11 {
                                        case "":
                                            v54 := GenS2{F0: v8, F1: 100, F2: 4294967295}
                                            _ = v54
                                        case "hello":
                                            if ((!false) && ((!(float32(v1) < v9)) || false)) {
                                                v11 = "foo bar"
                                            }
                                        default:
                                            fmt.Println(v3)
                                        }
                                    case "\n\t":
                                        var v55 uint32 = 65535
                                        _ = v55
                                    default:
                                        v56 := GenS2{F0: 18446744073709551615, F1: -27, F2: 65535}
                                        _ = v56
                                    }
                                }
                                switch v3 {
                                case 127:
                                    fmt.Println(v9)
                                case -2147483648:
                                    if (65535 == 560) {
                                        v57 := GenS2{F0: 0, F1: -30, F2: 760}
                                        var v58 int8 = 42
                                        _ = v57
                                        _ = v58
                                    } else {
                                        v59 := GenS2{F0: (^(v8)), F1: -1, F2: 826}
                                        var v60 int8 = 26
                                        _ = v59
                                        _ = v60
                                    }
                                case 42:
                                    if (!true) {
                                        if ((-2.25 >= 0.1) || (0.1 >= (1.5 - 0.1))) {
                                            v61 := v5[v11]
                                            switch (v3 % (((v3 >> ((4294967295) & 63))) | 1)) {
                                            case -1:
                                                var v62 uint64 = 0
                                                _ = v62
                                            case 137:
                                                if ((((-1 > 32767) && (false && ((v61 + ("foo bar" + "a")) < ((v61 + (v11 + (v11 + "\n\t"))) + v11)))) && (!true)) || (42 != 123)) {
                                                    switch v61 {
                                                    case "foo bar":
                                                        var v63 int32 = -16
                                                        _ = v63
                                                    case "":
                                                        var v64 int16 = int16(v8)
                                                        _ = v64
                                                    case "\n\t":
                                                        if ((v8 | v8) <= (v8 >> ((0) & 63))) {
                                                            fmt.Println(v51)
                                                            if (32767 > 32767) {
                                                                var v65 int = -9223372036854775808
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                var v66 uint8 = v7
                                                                _ = v65
                                                                _ = v66
                                                            } else {
                                                                delete(v4, 624)
                                                            }
                                                            v67 := v5[v61]
                                                            _ = v67
                                                        } else {
                                                            var v68 bool = (-1 >= -128)
                                                            var v69 uint = 773
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                    default:
                                                        if true {
                                                            var v70 int8 = 1
                                                            _ = v70
                                                        } else {
                                                            switch ((v61 + ((v61 + v61) + "a")) + ("\n\t" + v61)) {
                                                            case "":
                                                                var v71 uint32 = 2
                                                                _ = v71
                                                            default:
                                                                v8 -= (v8 - v8)
                                                            }
                                                        }
                                                    }
                                                    v72 := v5[v11]
                                                    var v73 bool = true
                                                    _ = v72
                                                    _ = v73
                                                }
                                            default:
                                                v74 := GenS1{F0: 969, F1: (v51 >> ((433) & 63))}
                                                _ = v74
                                            }
                                            if true {
                                                fmt.Println(v9)
                                                var v75 string = ""
                                                _ = v75
                                            } else {
                                                v76 := GenS3{F0: -295, F1: 0, F2: -1}
                                                var v77 int8 = 42
                                                _ = v76
                                                _ = v77
                                            }
                                            _ = v61
                                        } else {
                                            var v78 uint16 = 2
                                            switch v3 {
                                            case 723:
                                                var v79 float32 = ((-2.25 - 1.5) * -2.25)
                                                _ = v79
                                            default:
                                                v6[4294967295] = (v8 != v8)
                                                var v80 uint64 = (v8 + v8)
                                                _ = v80
                                            }
                                            if (0 < 32767) {
                                                v81 := v4[2147483648]
                                                v3 -= v3
                                                _ = v81
                                            } else {
                                                switch v11 {
                                                case "foo bar":
                                                    v78 -= 641
                                                case "a":
                                                    if (968 != 65535) {
                                                        var v82 int32 = -2147483648
                                                        if (fi50 > fi50) {
                                                            v2 = append(v2, v11)
                                                        }
                                                        _ = v82
                                                    } else {
                                                        var v83 uint16 = v78
                                                        _ = v83
                                                    }
                                                case "":
                                                    v84 := GenS4{F0: 865, F1: -1, F2: -2147483648}
                                                    _ = v84
                                                }
                                                fmt.Println(v3, v3)
                                            }
                                            fmt.Println(v51, v51, fi49)
                                            _ = v78
                                        }
                                        v51 += 1
                                    } else {
                                        fmt.Println(fi49, v7, v3)
                                        fmt.Println(v3, v7)
                                    }
                                }
                                _ = v51
                            } else {
                                v85 := GenS5{F0: 2147483648, F1: 18446744073709551615}
                                _ = v85
                            }
                        }
                    } else {
                        var v86 float64 = v1
                        v87 := GenS6{F0: -128}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v86
                        _ = v87
                    }
                }
            } else {
                if true {
                    if (75 <= 127) {
                        v3 -= -1
                        fmt.Println(v1, fi50, v3)
                        v7 = (^(v7))
                        v8 = v8
                        delete(v5, "a")
                    }
                    var v88 int8 = 127
                    switch ("\n\t" + "foo bar") {
                    case "a":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "":
                        v8 += (v8 * (^(v8)))
                    case "foo bar":
                        v11 = ""
                    }
                    _ = v88
                } else {
                    v1 = float64(3.14159)
                    fmt.Println(fi49, v8)
                    v9 = (v9 * (v9 * 3.14159))
                }
                var v89 uint32 = 42
                _ = v89
            }
            _ = fi50
        }
        _ = fi49
    }
    var v90 uint = 836
    if true {
        if (!(v7 != v7)) {
            fmt.Println(v11, v11)
            var v91 int64 = 1
            _ = v91
        }
        for fi92 := 0; fi92 < 16; fi92++ {
            v93 := GenS3{F0: -1, F1: uint16(v90), F2: -961}
            _ = fi92
            _ = v93
        }
    }
    switch v3 {
    case 298:
        var v94 uint16 = uint16(v8)
        _ = v94
    }
    for fi95 := 0; fi95 < 14; fi95++ {
        if (!true) {
            var v96 int8 = -1
            _ = v96
        } else {
            fmt.Println(v7)
            v97 := v6[uint32(v90)]
            _ = v97
        }
        _ = fi95
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    _ = v90
}
