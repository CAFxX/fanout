package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint32; F2 bool }
type GenS3 struct { F0 int; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint64; F1 int32; F2 int8 }
type GenS5 struct { F0 uint64; F1 int8; F2 int }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int64; F1 bool; F2 uint8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint8; F1 int32; F2 int }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 string; F1 uint }
type GenS14 struct { F0 uint16; F1 int; F2 int }
type GenS15 struct { F0 uint; F1 uint16 }
type GenS16 struct { F0 int16; F1 uint; F2 int16 }
type GenS17 struct { F0 bool; F1 uint32; F2 string }

func main() {
    var v1 uint16 = 42
    v2 := []string{"foo bar", "hello"}
    v3 := []string{("" + (("" + "") + "hello")), "hello"}
    var v4 bool = (42 > 266)
    v5 := []uint32{4294967295, 969, 42, 85}
    v1 = 0
    switch 42 {
    case 822:
        var v6 int32 = 491
        _ = v6
    case 9223372036854775807:
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 uint = 348
            var v9 int8 = 1
            if (0.1 < float64(0.0)) {
                if false {
                    v8 = v8
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                }
                var v10 bool = false
                _ = v10
            } else {
                switch (("hello" + ("" + "foo bar")) + "") {
                case "":
                    if (v4 || v4) {
                        if (true || (255 < 255)) {
                            _ = len(v3)
                            v8 += (^(v8))
                        } else {
                            switch "" {
                            case "":
                                v11 := GenS1{F0: 178}
                                _ = v11
                            case "a":
                                for fi12 := 0; fi12 < 6; fi12++ {
                                    switch ((("\n\t" + "") + "hello") + (("" + ("foo bar" + "a")) + "a")) {
                                    case "a":
                                        var v13 uint8 = 127
                                        _ = v13
                                    case "":
                                        v9 += 100
                                    case "hello":
                                        switch 1 {
                                        case -1:
                                            var v14 uint = v8
                                            _ = v14
                                        case 32767:
                                            switch 1 {
                                            case -1:
                                                if false {
                                                    switch ("\n\t" + "hello") {
                                                    case "\n\t":
                                                        if (!v4) {
                                                            fmt.Println(fi12, fi7)
                                                            v15 := GenS2{F0: -2147483648, F1: 1, F2: (true && v4)}
                                                            _ = v15
                                                        } else {
                                                            _ = len(v3)
                                                        }
                                                    case "foo bar":
                                                        v1 = v1
                                                    case "hello":
                                                        fmt.Println(v9)
                                                        if len(v3) > 0 {
                                                            v3[0] = (("" + ("foo bar" + (("hello" + "") + "hello"))) + "a")
                                                        }
                                                        var v16 bool = (2 > 255)
                                                        _ = v16
                                                    default:
                                                        v2 = append(v2, ("a" + ""))
                                                    }
                                                    if ((219 > 42) || (((("hello" == (("a" + "foo bar") + "\n\t")) && false) && (((-2.25 * (float32(3.14159) * -2.25)) <= ((0.1 / 1.0) / 1e10)) || (0 < -144))) || v4)) {
                                                        var v17 float32 = 0.0
                                                        fmt.Println(v4, v8)
                                                        v18 := GenS2{F0: 1, F1: 489, F2: false}
                                                        _ = v17
                                                        _ = v18
                                                    } else {
                                                        var v19 int32 = -2147483648
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v19
                                                    }
                                                    v8 -= (v8 % (((v8 >> ((v8) & 63))) | 1))
                                                } else {
                                                    v8 = v8
                                                    _ = len(v3)
                                                    v4 = v4
                                                }
                                            case 4294967295:
                                                v1 = v1
                                                v20 := GenS3{F0: fi7, F1: -2147483648, F2: 932}
                                                _ = v20
                                            }
                                            if len(v5) > 0 {
                                                v5[0] = 1
                                            }
                                            fmt.Println(v8, v4, v1)
                                        default:
                                            var v21 uint = v8
                                            _ = v21
                                        }
                                    default:
                                        v9 -= (v9 & 91)
                                    }
                                    _ = fi12
                                }
                            case "foo bar":
                                var v22 uint64 = 42
                                _ = v22
                            }
                            v1 = (v1 + 65535)
                        }
                        v23 := GenS4{F0: 18446744073709551615, F1: 1, F2: v9}
                        if true {
                            v4 = (false || ("" >= "a"))
                            v24 := GenS5{F0: 1, F1: v9, F2: fi7}
                            _ = v24
                        }
                        _ = v23
                    }
                }
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    case 303:
        for fi25 := 0; fi25 < 10; fi25++ {
            var v26 bool = (-1 != 127)
            v4 = (false || (v1 != v1))
            _ = fi25
            _ = v26
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        for fi28 := 0; fi28 < 8; fi28++ {
            _ = len(v5)
            fmt.Println(v4)
            v1 = (v1 / (((v1 << ((18446744073709551615) & 63))) | 1))
            if true {
                v2 = append(v2, (("foo bar" + "") + ""))
            } else {
                v29 := GenS2{F0: -1, F1: 1, F2: false}
                _ = v29
            }
            switch ("hello" + "hello") {
            case "a":
                var v30 float64 = 3.14159
                _ = v30
            }
            _ = len(v2)
            _ = fi28
        }
        var v31 uint16 = 255
        _ = fi27
        _ = v31
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        if v4 {
            for fi33 := 0; fi33 < 9; fi33++ {
                var v34 uint64 = 394
                v34 = v34
                if (v4 || false) {
                    v34 = (^(v34))
                    v1 = 255
                    var v35 uint32 = 65535
                    _ = v35
                }
                _ = fi33
                _ = v34
            }
            v1 += (v1 >> ((42) & 63))
            for fi36 := 0; fi36 < 20; fi36++ {
                var v37 uint64 = 456
                v4 = (-1 > 0)
                if len(v5) > 0 {
                    v5[0] = 0
                }
                _ = fi36
                _ = v37
            }
        } else {
            v38 := GenS6{F0: 0, F1: 1}
            for fi39 := 0; fi39 < 12; fi39++ {
                v40 := GenS6{F0: 2147483647, F1: 286}
                if (79 < 567) {
                    switch "" {
                    case "hello":
                        v1 -= 458
                    case "a":
                        var v41 string = ""
                        _ = v41
                    case "foo bar":
                        switch (fi32 & 127) {
                        case -86:
                            var v42 bool = (0 >= 9223372036854775807)
                            _ = v42
                        default:
                            var v43 int32 = 2147483647
                            _ = v43
                        }
                    default:
                        if (651 > -2147483648) {
                            v44 := GenS7{F0: 42, F1: 411}
                            v45 := GenS3{F0: (^(fi32)), F1: -2147483648, F2: 18446744073709551615}
                            _ = v44
                            _ = v45
                        } else {
                            var v46 uint8 = 2
                            _ = v46
                        }
                    }
                    v4 = (!v4)
                    v1 -= (v1 & (v1 * (^(v1))))
                } else {
                    v4 = (v4 && (!(0 != 42)))
                    var v47 int32 = 42
                    if false {
                        fmt.Println(v4)
                    } else {
                        var v48 float32 = float32(((1.5 - float64(-2.25)) / 1e10))
                        if false {
                            v1 -= v1
                            var v49 int64 = 419
                            _ = v49
                        } else {
                            var v50 int16 = -1
                            _ = v50
                        }
                        v51 := GenS8{F0: "foo bar"}
                        switch -287 {
                        case -2147483648:
                            v47 += v47
                        default:
                            if v4 {
                                var v52 int64 = 1
                                var v53 int16 = 42
                                _ = v52
                                _ = v53
                            }
                        }
                        _ = v48
                        _ = v51
                    }
                    _ = len(v3)
                    var v54 int = 2147483647
                    _ = v47
                    _ = v54
                }
                if ((1 >= -45) || false) {
                    if true {
                        v1 += 74
                        var v55 int8 = 42
                        if (("" + (("\n\t" + "hello") + "")) == "foo bar") {
                            v55 -= 0
                            v1 = v1
                        } else {
                            v4 = (!v4)
                        }
                        if false {
                            fmt.Println(v55, fi39, v55)
                            if v4 {
                                v4 = v4
                                var v56 int8 = -1
                                _ = v56
                            }
                        } else {
                            if (false || v4) {
                                if v4 {
                                    var v57 uint32 = 237
                                    var v58 float64 = -2.25
                                    fmt.Println(v1, fi32, v1)
                                    switch fi32 {
                                    case -2147483648:
                                        fmt.Println(fi32, fi39, v4)
                                    case 127:
                                        v59 := GenS8{F0: "a"}
                                        _ = v59
                                    }
                                    _ = v57
                                    _ = v58
                                } else {
                                    var v60 int64 = 791
                                    _ = v60
                                }
                                switch fi39 {
                                case 1:
                                    _ = len(v3)
                                case -844:
                                    var v61 int32 = -2147483648
                                    _ = v61
                                default:
                                    if false {
                                        fmt.Println(v4, fi39, fi32)
                                    }
                                }
                                switch (fi32 >> ((65535) & 63)) {
                                case -9223372036854775808:
                                    v55 -= v55
                                default:
                                    _ = len(v2)
                                }
                            }
                            if (42 > 62) {
                                var v62 uint64 = 444
                                v55 = 0
                                v1 += 233
                                fmt.Println(v4, v55)
                                _ = v62
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = ""
                                }
                            }
                        }
                        fmt.Println(v55, fi32)
                        _ = v55
                    } else {
                        v63 := GenS9{F0: 1, F1: v4, F2: 2}
                        v64 := GenS9{F0: -9223372036854775808, F1: false, F2: 0}
                        v65 := GenS10{F0: -792}
                        _ = v63
                        _ = v64
                        _ = v65
                    }
                    v1 += 699
                    v66 := GenS6{F0: 504, F1: -1}
                    var v67 uint64 = 181
                    var v68 uint = 1
                    switch fi39 {
                    case 2147483647:
                        switch "foo bar" {
                        case "":
                            switch ("" + "\n\t") {
                            case "foo bar":
                                v67 -= (v67 / ((225) | 1))
                            case "\n\t":
                                var v69 uint64 = (v67 | 4294967295)
                                _ = v69
                            }
                        case "\n\t":
                            v67 = 887
                            var v70 int = fi32
                            fmt.Println(v68)
                            _ = v70
                        default:
                            if v4 {
                                if (!true) {
                                    var v71 bool = ((v67 % (((v67 / ((v67) | 1))) | 1)) < 1)
                                    fmt.Println(v67, v1, fi32)
                                    v68 = (v68 - (v68 ^ v68))
                                    _ = v71
                                }
                                if true {
                                    v72 := GenS6{F0: 2147483647, F1: -36}
                                    if len(v3) > 0 {
                                        v3[0] = "a"
                                    }
                                    _ = v72
                                } else {
                                    switch "a" {
                                    case "":
                                        v3 = append(v3, "hello")
                                    case "a":
                                        v5 = append(v5, 407)
                                    default:
                                        var v73 float64 = ((0.0 * float64(3.14159)) + ((0.1 * 3.14159) / 1.0))
                                        _ = v73
                                    }
                                    var v74 int = fi39
                                    v68 = (v68 / ((4294967295) | 1))
                                    _ = v74
                                }
                            }
                        }
                        v67 = v67
                        v2 = append(v2, (("" + "") + "foo bar"))
                        v75 := GenS7{F0: 517, F1: 1}
                        var v76 float32 = (1.5 / 0.1)
                        v77 := GenS7{F0: 304, F1: 42}
                        _ = v75
                        _ = v76
                        _ = v77
                    case 42:
                        switch (^(fi32)) {
                        case 909:
                            if v4 {
                                if len(v5) > 0 {
                                    v5[0] = 2147483648
                                }
                            }
                        case -501:
                            v5 = append(v5, 1)
                        default:
                            v78 := GenS3{F0: -371, F1: -1, F2: 395}
                            _ = v78
                        }
                    default:
                        switch (fi32 >> ((v68) & 63)) {
                        case -9223372036854775808:
                            if (32767 < -32768) {
                                var v79 uint8 = uint8(v68)
                                var v80 uint16 = 65535
                                if (v4 || (v4 && v4)) {
                                    v67 = (v67 | 4294967295)
                                    switch "hello" {
                                    case "\n\t":
                                        v1 = v1
                                    case "a":
                                        var v81 int16 = int16(v67)
                                        _ = v81
                                    default:
                                        if v4 {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            v67 -= v67
                                            v82 := GenS11{F0: v79, F1: -2147483648, F2: (fi39 * (fi39 << ((v68) & 63)))}
                                            _ = v82
                                        } else {
                                            v83 := GenS4{F0: v67, F1: 2147483647, F2: -1}
                                            _ = v83
                                        }
                                        fmt.Println(fi32, v1)
                                        if v4 {
                                            var v84 string = "foo bar"
                                            switch 32768 {
                                            case -400:
                                                v80 = uint16(v79)
                                            case 42:
                                                if ((v79 >> (((v68 ^ (v68 ^ (v68 << ((v68) & 63))))) & 63)) < 255) {
                                                    if (-1 < -798) {
                                                        v80 -= (v80 ^ 820)
                                                    } else {
                                                        v85 := GenS12{F0: 0}
                                                        switch (fi32 ^ fi32) {
                                                        case -9223372036854775808:
                                                            v79 = (v79 << ((305) & 63))
                                                            var v86 float64 = 3.14159
                                                            if ((v1 < v1) && (!((((("foo bar" + "hello") + ("" + (v84 + (v84 + "")))) + v84) + "foo bar") < "foo bar"))) {
                                                                _ = len(v3)
                                                                v79 += v79
                                                                var v87 int64 = 42
                                                                if len(v5) > 0 {
                                                                    v5[0] = uint32(v68)
                                                                }
                                                                _ = v87
                                                            } else {
                                                                if true {
                                                                    var v88 bool = v4
                                                                    var v89 string = (v84 + "a")
                                                                    _ = v88
                                                                    _ = v89
                                                                } else {
                                                                    v3 = append(v3, "a")
                                                                    v3 = append(v3, v84)
                                                                    v5 = append(v5, 0)
                                                                }
                                                                var v90 uint64 = 180
                                                                if false {
                                                                    switch fi39 {
                                                                    case -2147483648:
                                                                        switch fi32 {
                                                                        case -712:
                                                                            v3 = append(v3, (v84 + v84))
                                                                        }
                                                                    case -9223372036854775808:
                                                                        v86 = (-2.25 - v86)
                                                                    default:
                                                                        fmt.Println(v67, v79)
                                                                    }
                                                                    v91 := GenS6{F0: -2147483648, F1: 0}
                                                                    var v92 int8 = -72
                                                                    fmt.Println(v92, v90, v92)
                                                                    if (v86 != float64(0.1)) {
                                                                        var v93 string = "foo bar"
                                                                        v94 := GenS10{F0: -666}
                                                                        v1 = (v1 | 0)
                                                                        if (1000 == 1000) {
                                                                            v90 -= (v90 << ((v68) & 63))
                                                                            if len(v3) > 0 {
                                                                                v3[0] = v84
                                                                            }
                                                                            v95 := GenS13{F0: "foo bar", F1: 65535}
                                                                            _ = v95
                                                                        }
                                                                        _ = v93
                                                                        _ = v94
                                                                    } else {
                                                                        fmt.Println(fi32, v4)
                                                                    }
                                                                    _ = v91
                                                                    _ = v92
                                                                }
                                                                _ = v90
                                                            }
                                                            _ = v86
                                                        default:
                                                            if ((1e10 * (1e10 + -2.25)) == (0.1 / 3.14159)) {
                                                                fmt.Println(v4)
                                                                v84 = (("a" + "a") + v84)
                                                            } else {
                                                                fmt.Println(v67, v4)
                                                                if v4 {
                                                                    var v96 int = fi39
                                                                    if ((127 >= 0) || v4) {
                                                                        v4 = ((!true) && (307 < uint32(v79)))
                                                                        var v97 uint64 = (v67 / ((225) | 1))
                                                                        v1 += (v80 + (v80 % (((v80 ^ (v80 * uint16(v79)))) | 1)))
                                                                        _ = v97
                                                                    }
                                                                    switch 32768 {
                                                                    case 42:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = "a"
                                                                        }
                                                                    case -145:
                                                                        fmt.Println(v79, v80, v67)
                                                                    default:
                                                                        var v98 int64 = 1
                                                                        _ = v98
                                                                    }
                                                                    _ = v96
                                                                } else {
                                                                    fmt.Println(fi39, v84, fi39)
                                                                }
                                                            }
                                                        }
                                                        switch (v84 + v84) {
                                                        case "":
                                                            v99 := GenS7{F0: 0, F1: 2147483647}
                                                            _ = v99
                                                        case "foo bar":
                                                            v80 = 2
                                                        default:
                                                            v84 = "a"
                                                        }
                                                        _ = v85
                                                    }
                                                } else {
                                                    v67 += 0
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    v84 = v84
                                                }
                                                fmt.Println(fi39, v79)
                                            default:
                                                if len(v5) > 0 {
                                                    v5[0] = 2147483647
                                                }
                                            }
                                            v67 += (v67 ^ (v67 >> (((v68 % (((v68 * (v68 << ((v68) & 63)))) | 1))) & 63)))
                                            fmt.Println(fi39)
                                            _ = v84
                                        } else {
                                            if len(v5) > 0 {
                                                v5[0] = 0
                                            }
                                            var v100 uint = 257
                                            fmt.Println(v1, v100, v67)
                                            var v101 int32 = -2147483648
                                            _ = v100
                                            _ = v101
                                        }
                                    }
                                    switch (-(fi32)) {
                                    case 4294967295:
                                        v5 = append(v5, 1)
                                    default:
                                        if (-2.25 < 0.0) {
                                            v2 = append(v2, "\n\t")
                                        }
                                    }
                                }
                                _ = v79
                                _ = v80
                            } else {
                                if ((v4 || true) && (1000 < 1000)) {
                                    v67 = v67
                                    v2 = append(v2, "foo bar")
                                    _ = len(v3)
                                }
                            }
                        case 9223372036854775807:
                            if ((-32768 < -368) || (v1 == (v1 + v1))) {
                                var v102 int64 = -637
                                v103 := GenS8{F0: "hello"}
                                _ = v102
                                _ = v103
                            } else {
                                switch fi39 {
                                case 2147483647:
                                    var v104 uint = (v68 & (v68 + v68))
                                    _ = v104
                                case -2147483648:
                                    var v105 uint32 = 450
                                    _ = v105
                                }
                                v4 = true
                            }
                            v106 := GenS10{F0: 592}
                            _ = v106
                        case -38:
                            v1 -= 255
                        }
                    }
                    _ = v66
                    _ = v67
                    _ = v68
                } else {
                    if true {
                        fmt.Println(v1)
                        var v107 string = "hello"
                        _ = v107
                    }
                    fmt.Println(fi39)
                }
                v4 = false
                _ = fi39
                _ = v40
            }
            _ = v38
        }
        _ = fi32
    }
    v4 = false
    v1 -= (v1 | 65535)
    v108 := GenS14{F0: 2, F1: 127, F2: 4294967295}
    fmt.Println(v1, v4)
    var v109 bool = (((v1 == v1) || (!v4)) || (-9223372036854775808 < 9223372036854775807))
    if v109 {
        switch ("foo bar" + "") {
        case "":
            for fi110 := 0; fi110 < 5; fi110++ {
                for fi111 := 0; fi111 < 15; fi111++ {
                    switch (fi110 >> ((2) & 63)) {
                    case 128:
                        v112 := GenS4{F0: 18446744073709551615, F1: 0, F2: 35}
                        _ = v112
                    case 42:
                        if (65535 > 42) {
                            v113 := GenS14{F0: (v1 * v1), F1: (fi111 + fi111), F2: (fi110 | (fi110 + 704))}
                            v114 := GenS3{F0: fi110, F1: 0, F2: 646}
                            _ = v113
                            _ = v114
                        }
                    case 1:
                        var v115 int = 42
                        _ = v115
                    }
                    var v116 uint64 = 0
                    _ = fi111
                    _ = v116
                }
                var v117 int64 = 9223372036854775807
                _ = fi110
                _ = v117
            }
        }
        if v4 {
            for fi118 := 0; fi118 < 20; fi118++ {
                v1 = v1
                for fi119 := 0; fi119 < 16; fi119++ {
                    fmt.Println(v1, v1)
                    _ = fi119
                }
                if (!((!("hello" >= "")) || false)) {
                    v1 -= (^(v1))
                    v1 = v1
                    fmt.Println(v1)
                } else {
                    if (1e10 < (float64(0.0) - (3.14159 * ((float64(3.14159) - 1.5) / 1e10)))) {
                        if len(v3) > 0 {
                            v3[0] = ("a" + ("" + "a"))
                        }
                        v109 = (v109 && (("hello" + ("\n\t" + ("a" + ("\n\t" + ("foo bar" + "a"))))) <= "hello"))
                    } else {
                        v5 = append(v5, 2)
                    }
                    v120 := GenS15{F0: 126, F1: v1}
                    _ = v120
                }
                _ = fi118
            }
            var v121 uint = 1
            _ = v121
        } else {
            var v122 int32 = -2147483648
            _ = v122
        }
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if len(v5) > 0 {
        v5[0] = 2
    }
    var v123 float64 = (0.1 * 1e10)
    if false {
        var v124 string = "\n\t"
        _ = v124
    } else {
        if v4 {
            switch 310 {
            case 128:
                _ = len(v5)
            case -495:
                var v125 float64 = 3.14159
                _ = v125
            case 410:
                var v126 bool = v4
                _ = v126
            }
            v123 = v123
        } else {
            v1 = (v1 >> ((18446744073709551615) & 63))
        }
        if true {
            v4 = ((!false) && false)
            for fi127 := 0; fi127 < 18; fi127++ {
                _ = len(v5)
                var v128 int = -76
                _ = fi127
                _ = v128
            }
        }
    }
    var v129 int32 = -2147483648
    var v130 float64 = float64(0.0)
    v131 := GenS12{F0: 65535}
    if v109 {
        switch "hello" {
        case "":
            v129 = (v129 * -564)
        case "\n\t":
            for fi132 := 0; fi132 < 17; fi132++ {
                v4 = (v129 >= v129)
                fmt.Println(v130, v130, v123)
                v129 += (v129 + v129)
                _ = fi132
            }
        case "hello":
            var v133 uint32 = 65535
            _ = v133
        }
        for fi134 := 0; fi134 < 19; fi134++ {
            v135 := GenS16{F0: -32768, F1: 225, F2: -32768}
            var v136 int32 = v129
            _ = fi134
            _ = v135
            _ = v136
        }
    } else {
        for fi137 := 0; fi137 < 6; fi137++ {
            for fi138 := 0; fi138 < 2; fi138++ {
                v2 = append(v2, ("a" + "foo bar"))
                v109 = (-541 > -32768)
                _ = fi138
            }
            switch "hello" {
            case "a":
                if v109 {
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                    for fi139 := 0; fi139 < 15; fi139++ {
                        if v4 {
                            var v140 int = fi139
                            v141 := GenS7{F0: 385, F1: v129}
                            _ = v140
                            _ = v141
                        }
                        v142 := GenS4{F0: 2, F1: (v129 | v129), F2: -63}
                        _ = fi139
                        _ = v142
                    }
                } else {
                    for fi143 := 0; fi143 < 19; fi143++ {
                        switch (fi143 ^ 128) {
                        case 158:
                            v130 -= (0.0 / (v123 * 3.14159))
                        case 127:
                            var v144 uint64 = 1
                            _ = v144
                        }
                        switch "" {
                        case "hello":
                            var v145 string = "foo bar"
                            _ = v145
                        case "":
                            switch fi137 {
                            case 9223372036854775807:
                                v146 := GenS4{F0: 4294967295, F1: v129, F2: -40}
                                _ = v146
                            case -9223372036854775808:
                                if ((v129 >= 2147483647) || ((0 != 127) || true)) {
                                    var v147 uint8 = 139
                                    v129 += (v129 / (((v129 / ((2147483647) | 1))) | 1))
                                    fmt.Println(v130, v4, v129)
                                    _ = v147
                                }
                                if (v4 && (!((!(true || v109)) && (((v129 >> ((65535) & 63)) < (v129 / ((v129) | 1))) && v109)))) {
                                    switch "\n\t" {
                                    case "a":
                                        var v148 uint16 = v1
                                        _ = v148
                                    case "foo bar":
                                        var v149 uint8 = 127
                                        _ = v149
                                    }
                                    switch (fi143 ^ (fi137 >> ((255) & 63))) {
                                    case 127:
                                        v130 = (0.1 * float64((0.1 / float32(0.1))))
                                    case -1:
                                        v130 += float64(-2.25)
                                    case 42:
                                        if len(v2) > 0 {
                                            v2[0] = "foo bar"
                                        }
                                    }
                                } else {
                                    v150 := GenS7{F0: 42, F1: int32(v1)}
                                    switch ("foo bar" + "hello") {
                                    case "a":
                                        var v151 int = (fi137 * fi137)
                                        _ = v151
                                    default:
                                        v5 = append(v5, 756)
                                    }
                                    fmt.Println(fi143, fi137)
                                    _ = v150
                                }
                            case 42:
                                if (!v109) {
                                    v130 -= (v123 - (((v123 * 3.14159) - (v123 / v123)) * (1e10 + 0.1)))
                                    var v152 int = fi137
                                    _ = v152
                                }
                                var v153 uint8 = 127
                                var v154 bool = true
                                var v155 int32 = (v129 - (v129 - v129))
                                _ = v153
                                _ = v154
                                _ = v155
                            default:
                                v1 = v1
                            }
                        case "a":
                            var v156 int64 = -1
                            _ = v156
                        }
                        v1 = (^(v1))
                        _ = fi143
                    }
                    var v157 int8 = 100
                    _ = v157
                }
            case "hello":
                _ = len(v2)
            case "\n\t":
                fmt.Println(v130, v1, v1)
            default:
                switch fi137 {
                case -2147483648:
                    v158 := GenS13{F0: "", F1: 18446744073709551615}
                    _ = v158
                default:
                    if (1 == 1) {
                        v159 := GenS17{F0: (v129 == (v129 + (v129 | (v129 >> ((2) & 63))))), F1: 2147483647, F2: "hello"}
                        _ = v159
                    }
                }
            }
            v3 = append(v3, (("\n\t" + "") + ("" + "a")))
            _ = fi137
        }
        v130 -= (float64(1e10) * 0.1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v123:", v123)
    fmt.Println("v129:", v129)
    fmt.Println("v130:", v130)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v108
    _ = v109
    _ = v123
    _ = v129
    _ = v130
    _ = v131
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
