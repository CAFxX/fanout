package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 bool; F1 string; F2 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS8 struct { F0 uint; F1 uint32 }
type GenS9 struct { F0 uint32; F1 uint32 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int8; F1 int64 }
type GenS12 struct { F0 int; F1 int16 }
type GenS13 struct { F0 string; F1 bool; F2 uint8 }
type GenS14 struct { F0 string; F1 int }
type GenS15 struct { F0 bool; F1 int; F2 uint32 }
type GenS16 struct { F0 int64; F1 int64; F2 uint8 }
type GenS17 struct { F0 int32; F1 int16; F2 string }
type GenS18 struct { F0 uint16; F1 int64; F2 int }

func main() {
    var v1 uint = 185
    var v2 float32 = 0.1
    v3 := []string{"a"}
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]string)
    v2 = 1.5
    delete(v4, 2)
    fmt.Println(v1, v2)
    v4[370] = 2147483648
    if (9223372036854775807 > -9223372036854775808) {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := v5["a"]
            _ = fi6
            _ = v7
        }
    } else {
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
    }
    v1 = 42
    fmt.Println(v2, v1)
    var v8 int = 0
    fmt.Println(v1, v8, v1)
    v1 = (^(v1))
    switch v8 {
    case -825:
        if false {
            v1 = v1
            v9 := GenS1{F0: 32767, F1: (v8 - v8)}
            if (0 == 10) {
                if (!false) {
                    switch -9223372036854775808 {
                    case 128:
                        var v10 int8 = 0
                        _ = v10
                    }
                } else {
                    switch -447 {
                    case 127:
                        fmt.Println(v8, v8)
                    case 174:
                        if true {
                            var v11 float32 = 1e10
                            for fi12 := 0; fi12 < 3; fi12++ {
                                v8 = v8
                                _ = fi12
                            }
                            fmt.Println(v11, v2)
                            _ = v11
                        } else {
                            v13 := GenS1{F0: 255, F1: v8}
                            _ = len(v4)
                            v14 := GenS1{F0: 42, F1: (-(v8))}
                            _ = v13
                            _ = v14
                        }
                    case 748:
                        if (-537 == 0) {
                            v15 := GenS2{F0: v8, F1: (-1 == 0)}
                            switch v8 {
                            case 0:
                                v2 += v2
                            case -2147483648:
                                var v16 float32 = 1e10
                                _ = v16
                            }
                            _ = v15
                        }
                    default:
                        for fi17 := 0; fi17 < 6; fi17++ {
                            v8 = (v8 << (((v1 + v1)) & 63))
                            if false {
                                fmt.Println(v1)
                                var v18 float64 = ((0.0 * 0.1) / ((0.1 - 3.14159) - 1.5))
                                var v19 uint16 = 42
                                _ = v18
                                _ = v19
                            } else {
                                v2 = v2
                            }
                            _ = fi17
                        }
                    }
                }
                var v20 uint64 = 4294967295
                var v21 int8 = -1
                _ = v20
                _ = v21
            }
            _ = v9
        }
    case 961:
        switch 128 {
        case 5:
            if true {
                var v22 int32 = -791
                if (!true) {
                    v2 = -2.25
                    v22 += (v22 >> ((v1) & 63))
                    v22 = (v22 ^ 42)
                } else {
                    if (42 == 57) {
                        v1 = 65535
                        fmt.Println(v8)
                        var v23 uint16 = 42
                        _ = v23
                    }
                    var v24 uint64 = 939
                    for fi25 := 0; fi25 < 16; fi25++ {
                        var v26 uint = 1
                        _ = fi25
                        _ = v26
                    }
                    _ = v24
                }
                _ = v22
            } else {
                _ = len(v4)
                fmt.Println(v2, v8)
            }
        case 2147483647:
            switch v8 {
            case -2147483648:
                v1 = 0
            case 81:
                var v27 int8 = -78
                _ = v27
            default:
                for fi28 := 0; fi28 < 7; fi28++ {
                    fmt.Println(v2, fi28, v2)
                    _ = fi28
                }
            }
        case 9223372036854775807:
            var v29 uint = 191
            _ = v29
        default:
            var v30 uint16 = 1
            _ = v30
        }
    }
    _ = len(v4)
    if (65535 >= 65535) {
        for fi31 := 0; fi31 < 17; fi31++ {
            if (!(!((!false) && (-1 < -1)))) {
                var v32 uint64 = 2
                v8 += -827
                _ = v32
            } else {
                if true {
                    var v33 uint8 = 42
                    _ = v33
                } else {
                    if true {
                        var v34 uint = (v1 >> (((v1 & v1)) & 63))
                        v1 += v1
                        switch v8 {
                        case 484:
                            var v35 int32 = -732
                            _ = v35
                        default:
                            fmt.Println(v1)
                        }
                        _ = v34
                    } else {
                        switch "a" {
                        case "\n\t":
                            v1 -= (^(v1))
                        case "foo bar":
                            switch 128 {
                            case 127:
                                if (30 > -9223372036854775808) {
                                    switch v8 {
                                    case 127:
                                        var v36 uint8 = 0
                                        _ = v36
                                    case -2147483648:
                                        switch (("" + "\n\t") + "") {
                                        case "a":
                                            if false {
                                                if (!(720 != 4294967295)) {
                                                    var v37 uint64 = 2
                                                    _ = v37
                                                }
                                            } else {
                                                if (false && false) {
                                                    var v38 int32 = -714
                                                    for fi39 := 0; fi39 < 5; fi39++ {
                                                        v2 = v2
                                                        switch fi31 {
                                                        case -1:
                                                            v1 = (v1 - v1)
                                                        default:
                                                            var v40 uint64 = 212
                                                            _ = v40
                                                        }
                                                        _ = fi39
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v38
                                                }
                                                v41 := GenS2{F0: (v8 | fi31), F1: (804 == v8)}
                                                _ = v41
                                            }
                                        case "":
                                            for fi42 := 0; fi42 < 10; fi42++ {
                                                var v43 bool = false
                                                _ = fi42
                                                _ = v43
                                            }
                                        default:
                                            fmt.Println(v1, v8, v2)
                                        }
                                    case -1:
                                        v1 = 255
                                    }
                                } else {
                                    v2 += (float32(0.1) + float32(float64(1e10)))
                                    var v44 uint = v1
                                    fmt.Println(v44)
                                    _ = v44
                                }
                            case 128:
                                v1 = v1
                            }
                        default:
                            if (127 <= 85) {
                                var v45 int64 = 9223372036854775807
                                var v46 bool = (((v2 * v2) * 1e10) == float32((-2.25 * float64((v2 + v2)))))
                                _ = v45
                                _ = v46
                            }
                        }
                        v1 = v1
                    }
                }
                var v47 int32 = -1
                if true {
                    for fi48 := 0; fi48 < 16; fi48++ {
                        fmt.Println(fi31, v1, v1)
                        _ = fi48
                    }
                    v1 += v1
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                } else {
                    fmt.Println(fi31, v47)
                    var v49 uint = v1
                    _ = v49
                }
                _ = v47
            }
            _ = fi31
        }
        v4[2147483648] = 335
        v2 = v2
    } else {
        for fi50 := 0; fi50 < 13; fi50++ {
            if false {
                var v51 string = "\n\t"
                fmt.Println(v2, v2, v2)
                switch "\n\t" {
                case "hello":
                    switch v8 {
                    case -9223372036854775808:
                        v52 := GenS2{F0: (v8 & (fi50 / ((v8) | 1))), F1: true}
                        _ = v52
                    case 9223372036854775807:
                        v51 = v51
                    case -2147483648:
                        fmt.Println(v2)
                    default:
                        v2 = float32(0.0)
                    }
                default:
                    switch (v8 + (v8 | 9223372036854775807)) {
                    case -780:
                        _ = len(v3)
                    }
                }
                _ = v51
            } else {
                if (18446744073709551615 <= 18446744073709551615) {
                    v2 += (v2 + (1e10 + 1.5))
                    switch "hello" {
                    case "a":
                        if (9223372036854775807 <= -55) {
                            var v53 int32 = 42
                            v54 := GenS2{F0: 128, F1: (32767 <= -1)}
                            v8 = v8
                            _ = v53
                            _ = v54
                        } else {
                            v1 += (v1 ^ 532)
                            var v55 float32 = float32(float64(((1.5 * float32(3.14159)) / v2)))
                            if false {
                                var v56 int8 = -1
                                if (42 > (v1 + (^(v1)))) {
                                    var v57 uint64 = 1
                                    for fi58 := 0; fi58 < 11; fi58++ {
                                        v8 -= (fi50 << (((^(v1))) & 63))
                                        v2 += ((float32(1e10) / v55) + 0.0)
                                        _ = fi58
                                    }
                                    _ = v57
                                } else {
                                    v55 = v2
                                }
                                if true {
                                    for fi59 := 0; fi59 < 6; fi59++ {
                                        delete(v5, ("a" + "\n\t"))
                                        if ((0.0 != ((0.1 - ((0.1 * -2.25) / 3.14159)) / (0.1 / (((0.1 / 3.14159) * 1.5) + 3.14159)))) && ((fi59 > (fi50 << ((65535) & 63))) || false)) {
                                            if (!(!false)) {
                                                if (!false) {
                                                    var v60 uint = (v1 << ((4294967295) & 63))
                                                    _ = v60
                                                } else {
                                                    if true {
                                                        v61 := GenS3{F0: (v56 + (^(v56))), F1: v56}
                                                        v62 := GenS2{F0: v8, F1: true}
                                                        _ = v61
                                                        _ = v62
                                                    } else {
                                                        switch "a" {
                                                        case "":
                                                            switch "foo bar" {
                                                            case "":
                                                                v8 -= (fi59 | (fi59 * int(v56)))
                                                            case "hello":
                                                                var v63 int8 = v56
                                                                _ = v63
                                                            case "foo bar":
                                                                if (v2 <= float32(3.14159)) {
                                                                    var v64 bool = false
                                                                    v8 -= (fi59 / ((fi59) | 1))
                                                                    var v65 string = ""
                                                                    _ = v64
                                                                    _ = v65
                                                                }
                                                            default:
                                                                if false {
                                                                    if false {
                                                                        v8 = -2147483648
                                                                        v66 := GenS2{F0: (fi59 >> (((v1 & (v1 | v1))) & 63)), F1: false}
                                                                        v67 := v5[("" + "a")]
                                                                        if (0 < 241) {
                                                                            v55 = 0.0
                                                                        } else {
                                                                            var v68 uint16 = 1
                                                                            v69 := GenS1{F0: (v68 >> (((^(v1))) & 63)), F1: (fi50 * fi59)}
                                                                            _ = v68
                                                                            _ = v69
                                                                        }
                                                                        _ = v66
                                                                        _ = v67
                                                                    } else {
                                                                        if (9223372036854775807 != 380) {
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            v3 = append(v3, "foo bar")
                                                                        } else {
                                                                            var v70 int8 = v56
                                                                            v2 -= 3.14159
                                                                            v3 = append(v3, "\n\t")
                                                                            _ = v70
                                                                        }
                                                                        if (true && (false || true)) {
                                                                            fmt.Println(fi50)
                                                                        }
                                                                    }
                                                                    v2 -= (1e10 * (1.5 - v2))
                                                                    v55 = v2
                                                                } else {
                                                                    if ((1 >= int32(v56)) && ((v2 != 1e10) && (1e10 >= (0.1 + 0.1)))) {
                                                                        v1 += v1
                                                                        v2 = (v2 / v55)
                                                                    }
                                                                    if false {
                                                                        _ = len(v3)
                                                                        v71 := GenS4{F0: int64(v1), F1: "\n\t"}
                                                                        if false {
                                                                            switch ("a" + "\n\t") {
                                                                            case "\n\t":
                                                                                v72 := GenS5{F0: (-479 > -1), F1: "a", F2: "hello"}
                                                                                _ = v72
                                                                            }
                                                                            v55 = 0.1
                                                                            switch 32767 {
                                                                            case -110:
                                                                                v73 := GenS6{F0: 251}
                                                                                _ = v73
                                                                            case 9223372036854775807:
                                                                                if (1 > 65535) {
                                                                                    var v74 uint = 255
                                                                                    var v75 int32 = 1
                                                                                    _ = v74
                                                                                    _ = v75
                                                                                } else {
                                                                                    if (32767 > 42) {
                                                                                        if false {
                                                                                            v76 := GenS7{F0: 9223372036854775807, F1: 2, F2: 0}
                                                                                            var v77 uint8 = 1
                                                                                            var v78 uint32 = 42
                                                                                            _ = v76
                                                                                            _ = v77
                                                                                            _ = v78
                                                                                        }
                                                                                    }
                                                                                    if true {
                                                                                        if (("a" + "\n\t") > "\n\t") {
                                                                                            var v79 float64 = (3.14159 / 1.5)
                                                                                            _ = v79
                                                                                        } else {
                                                                                            v8 += fi59
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case 0:
                                                                                v80 := GenS2{F0: (v8 << ((178) & 63)), F1: false}
                                                                                _ = v80
                                                                            }
                                                                            var v81 int = 1
                                                                            var v82 int64 = 42
                                                                            _ = v81
                                                                            _ = v82
                                                                        }
                                                                        switch fi59 {
                                                                        case 32767:
                                                                            v1 = (v1 - 643)
                                                                        case 948:
                                                                            if false {
                                                                                v1 = v1
                                                                            } else {
                                                                                var v83 string = ("foo bar" + (("hello" + "hello") + "\n\t"))
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                if (-1 != 42) {
                                                                                    _ = len(v5)
                                                                                }
                                                                                _ = v83
                                                                            }
                                                                        default:
                                                                            v55 = (v55 * 1e10)
                                                                        }
                                                                        if false {
                                                                            switch fi50 {
                                                                            case -278:
                                                                                v55 += 1.5
                                                                            }
                                                                        }
                                                                        if true {
                                                                            if (!(true && false)) {
                                                                                var v84 uint32 = 169
                                                                                _ = v84
                                                                            } else {
                                                                                if (!(2147483647 < 1)) {
                                                                                    v85 := GenS8{F0: 520, F1: 2147483647}
                                                                                    _ = v85
                                                                                } else {
                                                                                    v3 = append(v3, "")
                                                                                    if true {
                                                                                        v86 := GenS1{F0: 255, F1: 42}
                                                                                        switch 4294967295 {
                                                                                        case -905:
                                                                                            var v87 uint8 = 42
                                                                                            _ = v87
                                                                                        case -987:
                                                                                            fmt.Println(v8, v1)
                                                                                        }
                                                                                        fmt.Println(v8, v2, v8)
                                                                                        _ = v86
                                                                                    } else {
                                                                                        if false {
                                                                                            if (1e10 <= 0.1) {
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = "a"
                                                                                                }
                                                                                                switch "a" {
                                                                                                case "a":
                                                                                                    v1 += (v1 + v1)
                                                                                                case "":
                                                                                                    var v88 float64 = -2.25
                                                                                                    _ = v88
                                                                                                default:
                                                                                                    v5[(("" + "foo bar") + "\n\t")] = ""
                                                                                                }
                                                                                                if (!true) {
                                                                                                    v89 := GenS8{F0: 665, F1: 990}
                                                                                                    switch "" {
                                                                                                    case "\n\t":
                                                                                                        fmt.Println(v56)
                                                                                                    case "a":
                                                                                                        if (42 != 1) {
                                                                                                            v90 := GenS9{F0: 0, F1: 42}
                                                                                                            _ = v90
                                                                                                        }
                                                                                                    }
                                                                                                    v8 -= 4294967295
                                                                                                    v91 := GenS6{F0: 42}
                                                                                                    _ = v89
                                                                                                    _ = v91
                                                                                                } else {
                                                                                                    switch (fi50 ^ -9223372036854775808) {
                                                                                                    case 127:
                                                                                                        if (!(float64(v55) >= 1.5)) {
                                                                                                            var v92 int8 = int8(fi50)
                                                                                                            v93 := GenS3{F0: v56, F1: (v92 << (((v1 + 4294967295)) & 63))}
                                                                                                            v3 = append(v3, "")
                                                                                                            v94 := GenS5{F0: (float32(1.5) >= (v55 / v2)), F1: "hello", F2: ("foo bar" + ("" + (("a" + "hello") + "")))}
                                                                                                            _ = v92
                                                                                                            _ = v93
                                                                                                            _ = v94
                                                                                                        } else {
                                                                                                            v2 += ((v2 - v55) - 1.5)
                                                                                                            _ = len(v5)
                                                                                                            var v95 uint32 = 712
                                                                                                            var v96 int64 = 1
                                                                                                            _ = v95
                                                                                                            _ = v96
                                                                                                        }
                                                                                                    case 245:
                                                                                                        if false {
                                                                                                            switch fi50 {
                                                                                                            case 632:
                                                                                                                v56 = (v56 >> ((65535) & 63))
                                                                                                            case -498:
                                                                                                                v55 -= v2
                                                                                                            case -9223372036854775808:
                                                                                                                v97 := GenS9{F0: 65535, F1: 885}
                                                                                                                _ = v97
                                                                                                            default:
                                                                                                                switch "\n\t" {
                                                                                                                case "":
                                                                                                                    fmt.Println(v8, v2, v2)
                                                                                                                case "\n\t":
                                                                                                                    v1 = v1
                                                                                                                case "a":
                                                                                                                    v5[((("foo bar" + (("foo bar" + ("\n\t" + "foo bar")) + ("" + "foo bar"))) + "hello") + "a")] = ""
                                                                                                                default:
                                                                                                                    var v98 int16 = -43
                                                                                                                    _ = v98
                                                                                                                }
                                                                                                            }
                                                                                                            fmt.Println(v2, v56, v56)
                                                                                                        } else {
                                                                                                            fmt.Println(fi50, fi50, v55)
                                                                                                        }
                                                                                                    case 514:
                                                                                                        if false {
                                                                                                            v99 := GenS10{F0: 321}
                                                                                                            _ = v99
                                                                                                        } else {
                                                                                                            v2 = v55
                                                                                                            var v100 uint64 = 648
                                                                                                            var v101 uint8 = 194
                                                                                                            _ = v100
                                                                                                            _ = v101
                                                                                                        }
                                                                                                    default:
                                                                                                        switch 9223372036854775807 {
                                                                                                        case 0:
                                                                                                            var v102 int16 = 477
                                                                                                            _ = v102
                                                                                                        case -1:
                                                                                                            var v103 uint64 = 2
                                                                                                            _ = v103
                                                                                                        }
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                        var v104 int16 = 1
                                                                                                        switch fi59 {
                                                                                                        case 42:
                                                                                                            _ = len(v4)
                                                                                                        default:
                                                                                                            v105 := GenS10{F0: 1}
                                                                                                            _ = v105
                                                                                                        }
                                                                                                        if true {
                                                                                                            var v106 bool = false
                                                                                                            fmt.Println(v2, v104)
                                                                                                            if false {
                                                                                                                var v107 int16 = (v104 ^ (-(v104)))
                                                                                                                delete(v5, "hello")
                                                                                                                v2 -= -2.25
                                                                                                                _ = v107
                                                                                                            }
                                                                                                            _ = v106
                                                                                                        }
                                                                                                        _ = v104
                                                                                                    }
                                                                                                    fmt.Println(fi59, v1)
                                                                                                    v108 := GenS11{F0: (v56 & (v56 << (((v1 >> (((v1 + 18446744073709551615)) & 63))) & 63))), F1: 9223372036854775807}
                                                                                                    _ = v108
                                                                                                }
                                                                                            } else {
                                                                                                var v109 int16 = 0
                                                                                                v110 := GenS12{F0: (fi59 | -43), F1: v109}
                                                                                                var v111 uint32 = 1
                                                                                                _ = v109
                                                                                                _ = v110
                                                                                                _ = v111
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v55, v1)
                                                                                        var v112 uint16 = 0
                                                                                        _ = v112
                                                                                    }
                                                                                    var v113 bool = (1e10 != float32((1e10 + -2.25)))
                                                                                    if v113 {
                                                                                        v114 := GenS13{F0: (("\n\t" + "\n\t") + ""), F1: v113, F2: 186}
                                                                                        fmt.Println(v1)
                                                                                        v115 := GenS1{F0: 563, F1: (fi50 / ((v8) | 1))}
                                                                                        _ = v114
                                                                                        _ = v115
                                                                                    } else {
                                                                                        v1 = v1
                                                                                        v56 -= v56
                                                                                    }
                                                                                    v1 = 65535
                                                                                    _ = v113
                                                                                }
                                                                                v8 = fi59
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            }
                                                                            var v116 uint = 0
                                                                            if len(v3) > 0 {
                                                                                v3[0] = (("foo bar" + (("" + "hello") + "hello")) + "\n\t")
                                                                            }
                                                                            _ = v116
                                                                        } else {
                                                                            var v117 int16 = 1
                                                                            _ = v117
                                                                        }
                                                                        _ = v71
                                                                    } else {
                                                                        if false {
                                                                            var v118 uint16 = 65535
                                                                            if true {
                                                                                v56 -= v56
                                                                                if false {
                                                                                    v118 += 757
                                                                                }
                                                                                var v119 float64 = (0.1 * ((1e10 + -2.25) + -2.25))
                                                                                _ = v119
                                                                            } else {
                                                                                var v120 float32 = v55
                                                                                v121 := GenS9{F0: 65535, F1: uint32(v1)}
                                                                                var v122 uint64 = 1
                                                                                var v123 int8 = 1
                                                                                _ = v120
                                                                                _ = v121
                                                                                _ = v122
                                                                                _ = v123
                                                                            }
                                                                            _ = v118
                                                                        }
                                                                        delete(v5, ("foo bar" + ""))
                                                                        v8 = (v8 / ((-863) | 1))
                                                                        switch -572 {
                                                                        case -2147483648:
                                                                            if false {
                                                                                v124 := GenS14{F0: "", F1: 9223372036854775807}
                                                                                var v125 uint64 = uint64(v1)
                                                                                _ = len(v3)
                                                                                _ = v124
                                                                                _ = v125
                                                                            } else {
                                                                                fmt.Println(v2, v2)
                                                                            }
                                                                        case -1:
                                                                            if ((((int64(v1) != 1) || false) && ((^(v56)) > v56)) && (2 != 255)) {
                                                                                var v126 int8 = int8(fi50)
                                                                                _ = v126
                                                                            } else {
                                                                                fmt.Println(v55, v56, fi50)
                                                                                fmt.Println(fi50)
                                                                                var v127 int = (v8 ^ -874)
                                                                                switch -9223372036854775808 {
                                                                                case 511:
                                                                                    v127 = fi50
                                                                                default:
                                                                                    var v128 string = ((("" + ("foo bar" + "")) + "\n\t") + (("" + "hello") + "\n\t"))
                                                                                    _ = v128
                                                                                }
                                                                                fmt.Println(v1, v56, fi59)
                                                                                _ = v127
                                                                            }
                                                                        case 32768:
                                                                            v129 := GenS8{F0: v1, F1: 2}
                                                                            _ = v129
                                                                        default:
                                                                            fmt.Println(v8, v8, fi59)
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        case "\n\t":
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v8 = fi59
                                                            }
                                                        }
                                                        var v130 uint8 = 127
                                                        _ = v130
                                                    }
                                                    fmt.Println(v56, v1, fi59)
                                                    delete(v5, "")
                                                }
                                                var v131 int = (v8 << ((v1) & 63))
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v131
                                            } else {
                                                var v132 int32 = -847
                                                var v133 uint64 = 4294967295
                                                _ = v132
                                                _ = v133
                                            }
                                            v134 := GenS8{F0: v1, F1: 65535}
                                            fmt.Println(v1)
                                            _ = v134
                                        }
                                        var v135 float32 = v2
                                        switch (fi59 & fi59) {
                                        case 2147483647:
                                            v136 := GenS15{F0: true, F1: (fi59 ^ (fi59 + (fi59 & 32767))), F2: 42}
                                            _ = v136
                                        case 1:
                                            var v137 int32 = 2147483647
                                            _ = v137
                                        case -1:
                                            v55 -= ((1e10 + v55) * 1e10)
                                        }
                                        var v138 uint32 = 42
                                        fmt.Println(v2, v55, v135)
                                        _ = fi59
                                        _ = v135
                                        _ = v138
                                    }
                                    var v139 uint16 = 32767
                                    _ = v139
                                } else {
                                    v140 := GenS6{F0: 0}
                                    v56 -= -53
                                    fmt.Println(v55)
                                    _ = v140
                                }
                                _ = v56
                            }
                            _ = v55
                        }
                    case "foo bar":
                        if true {
                            v141 := GenS16{F0: 169, F1: 627, F2: 255}
                            v2 = v2
                            _ = v141
                        } else {
                            if false {
                                v142 := GenS17{F0: 2147483647, F1: 1, F2: ("" + ("foo bar" + ("" + "a")))}
                                v8 = 4294967295
                                v1 = 18446744073709551615
                                _ = v142
                            }
                        }
                    }
                    switch (v8 << (((^(v1))) & 63)) {
                    case 1:
                        fmt.Println(v8)
                    default:
                        var v143 int16 = 32767
                        _ = v143
                    }
                } else {
                    v8 = (v8 + (^(fi50)))
                    v144 := GenS15{F0: true, F1: fi50, F2: 65535}
                    _ = v144
                }
            }
            v1 -= (v1 | v1)
            _ = fi50
        }
        if false {
            for fi145 := 0; fi145 < 6; fi145++ {
                for fi146 := 0; fi146 < 20; fi146++ {
                    fmt.Println(fi145, fi146)
                    _ = fi146
                }
                for fi147 := 0; fi147 < 15; fi147++ {
                    if (((2 < uint32(v1)) && true) && true) {
                        v2 += v2
                        v148 := GenS1{F0: 0, F1: (fi145 << ((18446744073709551615) & 63))}
                        if false {
                            _ = len(v5)
                        }
                        var v149 uint = v1
                        v150 := GenS18{F0: 65535, F1: -9223372036854775808, F2: 9223372036854775807}
                        _ = v148
                        _ = v149
                        _ = v150
                    }
                    _ = fi147
                }
                var v151 string = "foo bar"
                _ = fi145
                _ = v151
            }
            var v152 string = "hello"
            _ = v152
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (857 < 0) {
        fmt.Println(v1, v8)
        fmt.Println(v1, v2, v1)
        v8 = v8
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
