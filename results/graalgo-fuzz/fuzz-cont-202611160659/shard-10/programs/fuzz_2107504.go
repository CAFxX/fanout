package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int; F1 int16; F2 uint32 }
type GenS3 struct { F0 string; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int16; F1 string; F2 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS8 struct { F0 int8; F1 int; F2 int8 }
type GenS9 struct { F0 uint; F1 uint; F2 uint32 }
type GenS10 struct { F0 int32; F1 bool; F2 bool }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int; F1 uint32 }
type GenS13 struct { F0 uint64; F1 uint32 }
type GenS14 struct { F0 uint8; F1 uint64 }
type GenS15 struct { F0 int; F1 uint64 }
type GenS16 struct { F0 int16; F1 uint16 }
type GenS17 struct { F0 string; F1 int16; F2 uint }

func main() {
    var v1 int16 = -841
    var v2 uint8 = 2
    var v3 uint16 = 914
    v4 := []string{}
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 uint8 = (v2 * 0)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: (v3 << ((18446744073709551615) & 63)), F1: 100}
    for fi8 := 0; fi8 < 17; fi8++ {
        v2 += v2
        for fi9 := 0; fi9 < 1; fi9++ {
            v2 = v2
            _ = fi9
        }
        if (32767 < v3) {
            if true {
                fmt.Println(v1, v1)
                if false {
                    var v10 uint32 = 4294967295
                    fmt.Println(fi8, v2, v3)
                    _ = v10
                }
            } else {
                if (false && true) {
                    for fi11 := 0; fi11 < 6; fi11++ {
                        var v12 float64 = ((-2.25 + ((float64(3.14159) / 1e10) - 3.14159)) - -2.25)
                        _ = fi11
                        _ = v12
                    }
                    v2 -= 1
                    for fi13 := 0; fi13 < 18; fi13++ {
                        v3 += v3
                        var v14 uint32 = 4294967295
                        _ = fi13
                        _ = v14
                    }
                } else {
                    v15 := GenS1{F0: 42, F1: -41}
                    var v16 float32 = (0.1 - (3.14159 - (((float32(float64(-2.25)) + 1e10) - (-2.25 / ((3.14159 * 1.5) * (((3.14159 / (1.5 / 0.1)) + (((1e10 * -2.25) / -2.25) * 3.14159)) + -2.25)))) / float32(0.1))))
                    v1 += v1
                    _ = v15
                    _ = v16
                }
                var v17 float64 = float64(0.1)
                for fi18 := 0; fi18 < 14; fi18++ {
                    v19 := GenS2{F0: fi18, F1: v1, F2: 2147483648}
                    if true {
                        v20 := GenS1{F0: v3, F1: 1}
                        if true {
                            v21 := GenS1{F0: 273, F1: -13}
                            var v22 uint64 = 2
                            _ = v21
                            _ = v22
                        } else {
                            v3 -= (^(v3))
                            v17 = 1.5
                        }
                        v3 += 2
                        _ = v20
                    } else {
                        v23 := GenS3{F0: "a", F1: v2, F2: 32767}
                        var v24 bool = true
                        switch (("" + ("" + "foo bar")) + "\n\t") {
                        case "":
                            fmt.Println(fi8, v2, fi8)
                        default:
                            fmt.Println(fi8)
                        }
                        _ = v23
                        _ = v24
                    }
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
            var v25 string = ("" + "hello")
            _ = v25
        }
        _ = fi8
    }
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    v1 = (-(v1))
    v2 = v2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        if (!true) {
            for fi27 := 0; fi27 < 15; fi27++ {
                switch "a" {
                case "hello":
                    var v28 int8 = -128
                    _ = v28
                case "foo bar":
                    if len(v4) > 0 {
                        v4[0] = "hello"
                    }
                    fmt.Println(fi27, v1, v3)
                default:
                    fmt.Println(v2, fi26)
                }
                if true {
                    v1 += v1
                    if (!true) {
                        if false {
                            v29 := GenS3{F0: "", F1: v2, F2: uint16(v2)}
                            if (("" + "a") <= "a") {
                                if (!(!("foo bar" <= ("hello" + "\n\t")))) {
                                    v30 := GenS3{F0: "foo bar", F1: v2, F2: (v3 + (v3 / ((2) | 1)))}
                                    _ = len(v4)
                                    _ = v30
                                } else {
                                    v3 += (^(v3))
                                    v1 = (v1 + v1)
                                }
                            }
                            _ = v29
                        }
                        switch (fi27 >> ((1) & 63)) {
                        case -1:
                            v4 = append(v4, "")
                        case -866:
                            fmt.Println(v2, v2, fi26)
                            var v31 uint8 = uint8(v3)
                            var v32 uint64 = 0
                            _ = v31
                            _ = v32
                        case 385:
                            if true {
                                fmt.Println(v1)
                                if (1 < -9223372036854775808) {
                                    v33 := GenS4{F0: (v1 ^ (v1 % ((1000) | 1))), F1: ("foo bar" + ("foo bar" + (("a" + "foo bar") + ("a" + "")))), F2: v1}
                                    v4 = append(v4, "a")
                                    v34 := GenS3{F0: "a", F1: 128, F2: (v3 | v3)}
                                    var v35 uint = 42
                                    fmt.Println(v3)
                                    if len(v4) > 0 {
                                        v4[0] = "a"
                                    }
                                    v35 = uint(v2)
                                    v1 += v1
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                                switch 2147483647 {
                                case 32767:
                                    switch fi26 {
                                    case 32768:
                                        var v36 int16 = v1
                                        _ = v36
                                    default:
                                        if true {
                                            v3 -= v3
                                            v37 := GenS5{F0: 2}
                                            _ = v37
                                        }
                                    }
                                }
                            } else {
                                switch fi26 {
                                case 2147483647:
                                    var v38 float64 = 1.5
                                    _ = v38
                                case 694:
                                    if len(v4) > 0 {
                                        v4[0] = ""
                                    }
                                    v4 = append(v4, (("foo bar" + "hello") + "a"))
                                }
                                v2 = (^(v2))
                            }
                        default:
                            if (("a" + (("a" + (("\n\t" + ("\n\t" + ("hello" + "\n\t"))) + "")) + "a")) != ("foo bar" + "")) {
                                v2 += (^(v2))
                                _ = len(v4)
                                v2 = 255
                                if true {
                                    if (false && (v2 >= (v2 | (v2 ^ v2)))) {
                                        fmt.Println(v3, fi26)
                                        v39 := GenS1{F0: v3, F1: 100}
                                        v1 += v1
                                        _ = v39
                                    } else {
                                        var v40 bool = false
                                        _ = v40
                                    }
                                    v3 += (v3 % ((v3) | 1))
                                }
                                if false {
                                    if (!false) {
                                        v3 += 782
                                        if (1e10 > (3.14159 * (1.5 / 1e10))) {
                                            if len(v4) > 0 {
                                                v4[0] = "foo bar"
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = "a"
                                            }
                                        } else {
                                            if false {
                                                _ = len(v4)
                                            } else {
                                                v2 += uint8(v3)
                                            }
                                            var v41 bool = (false || true)
                                            _ = v41
                                        }
                                        if false {
                                            var v42 float32 = 1e10
                                            if (!((!true) || false)) {
                                                switch (("" + "") + "") {
                                                case "foo bar":
                                                    if true {
                                                        if ((v3 * (v3 * 1)) < uint16(v2)) {
                                                            fmt.Println(v42, fi27, v1)
                                                            v42 += (1e10 * v42)
                                                            if false {
                                                                if false {
                                                                    v1 -= (v1 / (((^(v1))) | 1))
                                                                    _ = len(v4)
                                                                } else {
                                                                    v3 = v3
                                                                    var v43 string = "hello"
                                                                    v44 := GenS6{F0: 2}
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                                fmt.Println(fi26)
                                                                var v45 uint = 1
                                                                _ = len(v4)
                                                                var v46 int64 = 248
                                                                _ = v45
                                                                _ = v46
                                                            } else {
                                                                v47 := GenS1{F0: v3, F1: 127}
                                                                if ((v3 % ((v3) | 1)) <= (v3 | v3)) {
                                                                    fmt.Println(v42, v42, v42)
                                                                    fmt.Println(fi27, fi27, v3)
                                                                    if (false || ((v1 != (-(v1))) && ((float32(3.14159) * v42) >= ((v42 * v42) - ((v42 - v42) + (v42 * (float32((((float64((-2.25 / 1.5)) + (1.5 * ((((1.5 - 1.5) * 0.0) / (1e10 - ((-2.25 - (float64(v42) - -2.25)) * float64(0.1)))) + (((1e10 * -2.25) - 1e10) / 0.1)))) / 0.1) - -2.25)) - 0.1))))))) {
                                                                        switch (fi26 + fi26) {
                                                                        case 4294967295:
                                                                            v3 = 2
                                                                            fmt.Println(fi26)
                                                                            v48 := GenS7{F0: 156, F1: 100, F2: (v2 | (v2 % (((v2 - v2)) | 1)))}
                                                                            var v49 int = 526
                                                                            _ = v48
                                                                            _ = v49
                                                                        case -1:
                                                                            v50 := GenS5{F0: 42}
                                                                            _ = v50
                                                                        case 570:
                                                                            var v51 bool = false
                                                                            _ = v51
                                                                        }
                                                                        fmt.Println(fi27)
                                                                        var v52 uint = 255
                                                                        _ = v52
                                                                    } else {
                                                                        switch fi26 {
                                                                        case 812:
                                                                            switch fi26 {
                                                                            case -335:
                                                                                v2 = (v2 / ((v2) | 1))
                                                                            case -2147483648:
                                                                                v4 = append(v4, "\n\t")
                                                                            case -9223372036854775808:
                                                                                var v53 string = "\n\t"
                                                                                _ = v53
                                                                            }
                                                                        case 32768:
                                                                            if true {
                                                                                var v54 uint64 = 4294967295
                                                                                _ = v54
                                                                            } else {
                                                                                v4 = append(v4, "")
                                                                                v55 := GenS5{F0: 18446744073709551615}
                                                                                _ = v55
                                                                            }
                                                                        case 128:
                                                                            var v56 int64 = 42
                                                                            _ = v56
                                                                        default:
                                                                            _ = len(v4)
                                                                        }
                                                                        v2 += (v2 ^ (v2 + 238))
                                                                        switch ("a" + "\n\t") {
                                                                        case "\n\t":
                                                                            v57 := GenS5{F0: 0}
                                                                            _ = v57
                                                                        case "hello":
                                                                            if false {
                                                                                var v58 int32 = -1
                                                                                _ = v58
                                                                            }
                                                                        default:
                                                                            v1 = v1
                                                                        }
                                                                    }
                                                                } else {
                                                                    if (v42 <= v42) {
                                                                        var v59 int64 = 9223372036854775807
                                                                        fmt.Println(v42, v1, v2)
                                                                        fmt.Println(v3, v59, v2)
                                                                        _ = v59
                                                                    }
                                                                    _ = len(v4)
                                                                    v3 = v3
                                                                    v60 := GenS6{F0: 777}
                                                                    _ = v60
                                                                }
                                                                _ = v47
                                                            }
                                                        }
                                                        var v61 string = ("" + ("foo bar" + (("\n\t" + "foo bar") + "foo bar")))
                                                        _ = v61
                                                    } else {
                                                        var v62 uint8 = 1
                                                        if len(v4) > 0 {
                                                            v4[0] = "a"
                                                        }
                                                        _ = v62
                                                    }
                                                case "":
                                                    var v63 uint = 0
                                                    _ = v63
                                                }
                                                if false {
                                                    _ = len(v4)
                                                }
                                            } else {
                                                switch (-(fi26)) {
                                                case 4294967295:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                case -9223372036854775808:
                                                    if (v3 < (v3 >> ((789) & 63))) {
                                                        v64 := GenS8{F0: 0, F1: fi26, F2: -128}
                                                        fmt.Println(v2, v1)
                                                        _ = v64
                                                    } else {
                                                        var v65 uint8 = (v2 | (^(v2)))
                                                        _ = v65
                                                    }
                                                case 32768:
                                                    fmt.Println(fi27, v1)
                                                default:
                                                    switch (fi27 - (fi26 * int(v2))) {
                                                    case 127:
                                                        switch ((("a" + (("a" + ("" + "")) + ("a" + "a"))) + ("\n\t" + ((("" + "hello") + "hello") + "\n\t"))) + ("\n\t" + "")) {
                                                        case "hello":
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        case "a":
                                                            if (-2147483648 <= 2147483647) {
                                                                fmt.Println(fi27, v42, fi27)
                                                            } else {
                                                                v66 := GenS3{F0: ((((("hello" + "a") + "a") + (((("\n\t" + "\n\t") + ("" + "\n\t")) + (((("" + ("a" + "foo bar")) + "hello") + ("foo bar" + ("" + ("hello" + "foo bar")))) + "\n\t")) + ("hello" + "hello"))) + ("a" + "hello")) + ""), F1: v2, F2: v3}
                                                                fmt.Println(fi27, v2)
                                                                v67 := GenS5{F0: 0}
                                                                _ = v66
                                                                _ = v67
                                                            }
                                                        default:
                                                            v68 := GenS9{F0: 2, F1: 4294967295, F2: 583}
                                                            _ = v68
                                                        }
                                                    default:
                                                        v69 := GenS4{F0: v1, F1: "hello", F2: (v1 | v1)}
                                                        _ = v69
                                                    }
                                                }
                                                fmt.Println(v2)
                                                var v70 uint32 = 2
                                                _ = v70
                                            }
                                            switch (((("" + "hello") + "a") + "") + ("a" + ("hello" + "a"))) {
                                            case "":
                                                v1 = (v1 >> ((0) & 63))
                                                if false {
                                                    v1 += v1
                                                    var v71 float32 = v42
                                                    var v72 int32 = 617
                                                    _ = v71
                                                    _ = v72
                                                } else {
                                                    var v73 int = (fi27 / ((fi26) | 1))
                                                    switch "a" {
                                                    case "\n\t":
                                                        v3 += 32767
                                                    case "":
                                                        var v74 int32 = 242
                                                        v1 = (v1 << ((1) & 63))
                                                        _ = v74
                                                    case "a":
                                                        if (127 > 100) {
                                                            var v75 int16 = 0
                                                            var v76 uint8 = 2
                                                            var v77 float32 = (v42 + ((float32(1e10) - float32(float64(v42))) / v42))
                                                            _ = v75
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                    }
                                                    _ = v73
                                                }
                                            case "a":
                                                if len(v4) > 0 {
                                                    v4[0] = "hello"
                                                }
                                            case "foo bar":
                                                switch (fi27 * int(v1)) {
                                                case -346:
                                                    var v78 uint64 = 18446744073709551615
                                                    _ = v78
                                                case 176:
                                                    v79 := GenS3{F0: "foo bar", F1: 127, F2: v3}
                                                    _ = v79
                                                default:
                                                    if ("hello" <= "\n\t") {
                                                        v2 -= v2
                                                        if (2147483648 >= 42) {
                                                            switch ("" + "\n\t") {
                                                            case "hello":
                                                                var v80 int8 = 1
                                                                _ = v80
                                                            case "\n\t":
                                                                var v81 int64 = 1
                                                                if (-115 <= -1) {
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        v1 = (v1 + v1)
                                                                    case "hello":
                                                                        var v82 bool = ((-(v1)) >= (v1 / (((v1 + -877)) | 1)))
                                                                        _ = v82
                                                                    }
                                                                    if true {
                                                                        v83 := GenS10{F0: 17, F1: (fi27 > (fi27 + fi27)), F2: (369 <= 4294967295)}
                                                                        v1 = v1
                                                                        _ = v83
                                                                    } else {
                                                                        v84 := GenS11{F0: 671}
                                                                        _ = v84
                                                                    }
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        var v85 int64 = v81
                                                                        _ = v85
                                                                    case "\n\t":
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            var v86 uint16 = 0
                                                                            _ = v86
                                                                        default:
                                                                            var v87 int64 = (v81 % (((^(v81))) | 1))
                                                                            _ = v87
                                                                        }
                                                                    default:
                                                                        fmt.Println(fi26)
                                                                    }
                                                                }
                                                                if false {
                                                                    v88 := GenS4{F0: v1, F1: ("hello" + (("hello" + "\n\t") + "\n\t")), F2: v1}
                                                                    v89 := GenS12{F0: 32767, F1: 4294967295}
                                                                    _ = v88
                                                                    _ = v89
                                                                }
                                                                _ = v81
                                                            default:
                                                                var v90 int64 = 0
                                                                _ = v90
                                                            }
                                                            var v91 string = ("" + ("hello" + ((("hello" + (((((("foo bar" + "foo bar") + "a") + "a") + "hello") + "foo bar") + "foo bar")) + "foo bar") + ("foo bar" + "\n\t"))))
                                                            _ = v91
                                                        } else {
                                                            v4 = append(v4, "foo bar")
                                                            fmt.Println(fi26, v3)
                                                        }
                                                        fmt.Println(v1)
                                                    }
                                                }
                                            }
                                            v4 = append(v4, "")
                                            _ = v42
                                        } else {
                                            _ = len(v4)
                                            var v92 uint = 0
                                            var v93 uint = 456
                                            _ = v92
                                            _ = v93
                                        }
                                    } else {
                                        var v94 uint64 = 1
                                        _ = v94
                                    }
                                } else {
                                    v95 := GenS13{F0: 183, F1: 2}
                                    _ = v95
                                }
                                _ = len(v4)
                            }
                        }
                        var v96 int32 = 733
                        v2 = (v2 * (v2 ^ (^(v2))))
                        if (false && true) {
                            v3 = v3
                            var v97 bool = false
                            fmt.Println(v1, fi26)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if (((255 <= v2) && true) && (2 == 2147483647)) {
                                v97 = (18446744073709551615 < 692)
                                _ = len(v4)
                            } else {
                                switch fi26 {
                                case 1:
                                    v98 := GenS14{F0: 128, F1: 18446744073709551615}
                                    _ = v98
                                case -440:
                                    fmt.Println(fi26, v2)
                                case 9223372036854775807:
                                    var v99 bool = v97
                                    _ = v99
                                default:
                                    var v100 uint16 = v3
                                    _ = v100
                                }
                            }
                            _ = v97
                        }
                        _ = v96
                    }
                } else {
                    if false {
                        fmt.Println(v3)
                        v4 = append(v4, "")
                    } else {
                        v1 += v1
                        var v101 uint32 = 552
                        if (v1 < v1) {
                            v101 = v101
                            if (false || (240 != v2)) {
                                var v102 int64 = -284
                                _ = v102
                            }
                            var v103 int64 = 9223372036854775807
                            _ = v103
                        } else {
                            v3 -= (v3 << ((49) & 63))
                            v104 := GenS8{F0: -96, F1: fi27, F2: -1}
                            if ((1 >= 18446744073709551615) && true) {
                                if (1 >= 65535) {
                                    v101 += 65535
                                }
                            }
                            fmt.Println(v3, fi27, v3)
                            v105 := GenS10{F0: 0, F1: (!false), F2: false}
                            v106 := GenS15{F0: fi26, F1: 18446744073709551615}
                            _ = v104
                            _ = v105
                            _ = v106
                        }
                        _ = v101
                    }
                }
                _ = fi27
            }
            if true {
                v2 = v2
            } else {
                for fi107 := 0; fi107 < 8; fi107++ {
                    if ((((0.1 + -2.25) * (((1.5 * -2.25) * 1.5) - 0.1)) / 1.0) <= ((((1.5 + ((0.1 * 1.5) / 3.14159)) * (0.0 / 3.14159)) / -2.25) - 3.14159)) {
                        v108 := GenS4{F0: 42, F1: "hello", F2: -140}
                        _ = len(v4)
                        v109 := GenS16{F0: v1, F1: v3}
                        _ = v108
                        _ = v109
                    } else {
                        switch fi26 {
                        case 462:
                            v110 := GenS17{F0: "foo bar", F1: v1, F2: 65535}
                            _ = v110
                        case 2147483647:
                            v2 += (v2 >> ((0) & 63))
                        case 1:
                            if len(v4) > 0 {
                                v4[0] = "a"
                            }
                        default:
                            var v111 int32 = -1
                            _ = v111
                        }
                    }
                    if true {
                        v2 -= uint8(v3)
                    }
                    _ = fi107
                }
            }
        }
        _ = fi26
    }
    var v112 uint64 = 422
    var v113 uint32 = 4294967295
    var v114 bool = false
    var v115 bool = v114
    v1 += (-(v1))
    fmt.Println(v115)
    fmt.Println(v3)
    v115 = v114
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v113:", v113)
    fmt.Println("v114:", v114)
    fmt.Println("v115:", v115)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v112
    _ = v113
    _ = v114
    _ = v115
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
