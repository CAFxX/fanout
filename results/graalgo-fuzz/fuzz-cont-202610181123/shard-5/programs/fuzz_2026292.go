package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 int16; F2 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int8; F1 int }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 uint64; F1 int8 }
type GenS10 struct { F0 int; F1 uint16 }
type GenS11 struct { F0 uint32; F1 int }
type GenS12 struct { F0 uint8; F1 int32 }
type GenS13 struct { F0 int64; F1 uint16; F2 int32 }
type GenS14 struct { F0 uint16 }

func main() {
    v1 := []string{"foo bar", ((("foo bar" + ("\n\t" + (("foo bar" + (("\n\t" + (("\n\t" + "foo bar") + "\n\t")) + "hello")) + "a"))) + ("a" + "a")) + ("hello" + ("foo bar" + "foo bar"))), "a"}
    v2 := []string{"a"}
    var v3 float64 = float64(0.0)
    var v4 int16 = 1
    var v5 uint8 = 11
    v6 := []uint32{2147483648, 42}
    v7 := GenS1{F0: 1, F1: "foo bar", F2: 4294967295}
    var v8 uint64 = 265
    var v9 int64 = 884
    v4 = (v4 ^ v4)
    var v10 uint32 = 2147483647
    switch "foo bar" {
    case "":
        v10 -= (v10 ^ v10)
    case "a":
        v5 = 42
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 uint8 = v5
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        if true {
            v8 = v8
        }
        _ = fi13
    }
    v5 -= v5
    switch 0 {
    case 128:
        var v14 float32 = ((((0.1 / (-2.25 * -2.25)) - (float32(0.0) / 1.5)) - 0.1) * (1.5 - 0.1))
        _ = v14
    default:
        v2 = append(v2, "foo bar")
    }
    v15 := GenS1{F0: -1, F1: ("hello" + ""), F2: v8}
    v3 = (v3 + float64((0.1 / float32(1e10))))
    if false {
        v8 += (v8 * (^(v8)))
        var v16 uint64 = 251
        fmt.Println(v5)
        _ = v16
    }
    if (1 == 127) {
        for fi17 := 0; fi17 < 16; fi17++ {
            v3 -= (float64(3.14159) - float64(1e10))
            v8 = 42
            v8 -= 590
            _ = fi17
        }
        for fi18 := 0; fi18 < 8; fi18++ {
            switch (fi18 | fi18) {
            case -9223372036854775808:
                var v19 string = "a"
                _ = v19
            case 4294967295:
                switch fi18 {
                case 128:
                    var v20 float32 = ((float32(((float64(-2.25) + 0.0) * ((0.0 + v3) / v3))) * (float32(1e10) / 1.5)) + -2.25)
                    _ = v20
                default:
                    var v21 uint32 = v10
                    _ = v21
                }
            default:
                switch fi18 {
                case -9223372036854775808:
                    switch (^(fi18)) {
                    case -340:
                        var v22 string = (("a" + "a") + "\n\t")
                        _ = v22
                    case 42:
                        fmt.Println(v5, v4, v5)
                    }
                case 128:
                    v2 = append(v2, "\n\t")
                }
            }
            if (4294967295 >= 42) {
                v23 := GenS2{F0: v9}
                if true {
                    var v24 string = ("a" + (("a" + ("a" + (("a" + "foo bar") + "a"))) + "foo bar"))
                    if (!(true && ((fi18 << ((65535) & 63)) >= (fi18 ^ (^(fi18)))))) {
                        v25 := GenS3{F0: 657, F1: 65535}
                        _ = v25
                    } else {
                        if true {
                            v3 = v3
                            v24 = ("foo bar" + "\n\t")
                            var v26 uint = 1
                            _ = v26
                        } else {
                            v8 -= 2
                            var v27 uint64 = 4294967295
                            _ = v27
                        }
                        var v28 uint16 = 32767
                        v10 = uint32(v8)
                        _ = v28
                    }
                    if (0 >= 1) {
                        for fi29 := 0; fi29 < 1; fi29++ {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v9 = v9
                            _ = fi29
                        }
                        var v30 bool = (("a" + ("a" + v24)) == (v24 + v24))
                        _ = v30
                    } else {
                        var v31 int64 = (v9 / ((v9) | 1))
                        for fi32 := 0; fi32 < 1; fi32++ {
                            var v33 int = -9223372036854775808
                            if (fi32 <= v33) {
                                v34 := GenS2{F0: (v9 & v9)}
                                _ = v34
                            } else {
                                v8 += (v8 + 1)
                            }
                            switch (((v24 + (v24 + "hello")) + (v24 + ((v24 + ((v24 + v24) + ("a" + "\n\t"))) + v24))) + v24) {
                            case "foo bar":
                                fmt.Println(v24)
                            }
                            _ = fi32
                            _ = v33
                        }
                        switch v24 {
                        case "":
                            for fi35 := 0; fi35 < 13; fi35++ {
                                var v36 uint32 = (v10 ^ v10)
                                if true {
                                    var v37 string = (v24 + ("\n\t" + v24))
                                    if false {
                                        v36 -= (^(v36))
                                        if ((v10 != (^(v36))) || (127 >= -1)) {
                                            if (!true) {
                                                fmt.Println(v4)
                                                var v38 uint32 = v36
                                                if ((v38 < v10) || false) {
                                                    v39 := GenS3{F0: 487, F1: 65535}
                                                    v31 = v31
                                                    _ = v39
                                                } else {
                                                    switch v24 {
                                                    case "\n\t":
                                                        switch (v37 + "\n\t") {
                                                        case "foo bar":
                                                            v40 := GenS4{F0: (v9 | v9)}
                                                            _ = v40
                                                        case "":
                                                            fmt.Println(v36)
                                                        case "\n\t":
                                                            switch (fi35 >> ((449) & 63)) {
                                                            case 32767:
                                                                v6 = append(v6, (v38 % ((v38) | 1)))
                                                            default:
                                                                if len(v1) > 0 {
                                                                    v1[0] = v37
                                                                }
                                                            }
                                                        default:
                                                            if true {
                                                                v24 = v24
                                                                if (false || false) {
                                                                    var v41 uint = 2
                                                                    fmt.Println(v24)
                                                                    if false {
                                                                        if ("hello" != (("a" + "") + v24)) {
                                                                            if ("hello" < (v37 + ((v24 + v24) + ("hello" + "a")))) {
                                                                                switch v37 {
                                                                                case "":
                                                                                    var v42 int32 = 1
                                                                                    v43 := GenS4{F0: (v9 | (v9 * (-(v31))))}
                                                                                    _ = v42
                                                                                    _ = v43
                                                                                case "\n\t":
                                                                                    var v44 uint16 = 1
                                                                                    _ = v44
                                                                                default:
                                                                                    var v45 uint16 = 484
                                                                                    _ = v45
                                                                                }
                                                                                _ = len(v6)
                                                                                v38 = (v10 ^ 2147483647)
                                                                                v10 = 443
                                                                            }
                                                                        }
                                                                        switch v24 {
                                                                        case "a":
                                                                            v46 := GenS2{F0: v9}
                                                                            _ = v46
                                                                        case "hello":
                                                                            if len(v1) > 0 {
                                                                                v1[0] = (("hello" + "\n\t") + v24)
                                                                            }
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v4, v41)
                                                                        v9 += v9
                                                                        var v47 float64 = v3
                                                                        _ = v47
                                                                    }
                                                                    _ = v41
                                                                } else {
                                                                    if true {
                                                                        v38 = 548
                                                                        var v48 uint = 197
                                                                        var v49 uint64 = v8
                                                                        _ = v48
                                                                        _ = v49
                                                                    }
                                                                    if true {
                                                                        v31 = (v9 << ((65535) & 63))
                                                                        var v50 uint32 = (v38 >> ((255) & 63))
                                                                        _ = v50
                                                                    } else {
                                                                        v10 = 4294967295
                                                                        var v51 float32 = (0.0 - float32(v3))
                                                                        v31 -= v9
                                                                        _ = v51
                                                                    }
                                                                    v31 = -9223372036854775808
                                                                }
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    if (v24 < "hello") {
                                                                        if false {
                                                                            v52 := GenS5{F0: (("hello" + v24) + "a"), F1: 1, F2: 42}
                                                                            v31 = (v31 ^ v9)
                                                                            v3 -= 0.0
                                                                            _ = v52
                                                                        }
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        switch (fi18 - (fi18 >> ((42) & 63))) {
                                                                        case -357:
                                                                            if false {
                                                                                switch v24 {
                                                                                case "":
                                                                                    if len(v6) > 0 {
                                                                                        v6[0] = v10
                                                                                    }
                                                                                case "hello":
                                                                                    v36 = v10
                                                                                }
                                                                                v37 = (v37 + v24)
                                                                            } else {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = v24
                                                                                }
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            }
                                                                        case 4294967295:
                                                                            _ = len(v2)
                                                                        case -9223372036854775808:
                                                                            var v53 int16 = 42
                                                                            _ = v53
                                                                        }
                                                                    } else {
                                                                        var v54 int8 = 0
                                                                        v55 := GenS5{F0: v37, F1: 399, F2: v8}
                                                                        _ = v54
                                                                        _ = v55
                                                                    }
                                                                case "":
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                }
                                                            } else {
                                                                v56 := GenS6{F0: 100}
                                                                v8 = uint64(v5)
                                                                switch (fi18 >> ((4294967295) & 63)) {
                                                                case -1:
                                                                    v4 -= (v4 & v4)
                                                                case 9223372036854775807:
                                                                    var v57 uint = uint(v5)
                                                                    _ = v57
                                                                case 590:
                                                                    v38 += (v36 % ((v38) | 1))
                                                                default:
                                                                    v24 = (v24 + ("hello" + (v37 + (((v24 + v37) + v37) + "\n\t"))))
                                                                }
                                                                fmt.Println(v31, v4)
                                                                _ = v56
                                                            }
                                                        }
                                                    case "":
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                    fmt.Println(v10)
                                                }
                                                _ = v38
                                            } else {
                                                var v58 int64 = -149
                                                switch (v37 + "hello") {
                                                case "\n\t":
                                                    var v59 float32 = 1e10
                                                    _ = v59
                                                case "a":
                                                    v6 = append(v6, (v10 ^ (v10 / (((v36 % ((v10) | 1))) | 1))))
                                                default:
                                                    if (0.0 <= (3.14159 - (float32((0.0 - (3.14159 * 3.14159))) * 0.0))) {
                                                        var v60 uint32 = v10
                                                        if true {
                                                            var v61 int64 = v9
                                                            switch v24 {
                                                            case "a":
                                                                if (!(!true)) {
                                                                    switch 9223372036854775807 {
                                                                    case 4294967295:
                                                                        var v62 int16 = (v4 / ((v4) | 1))
                                                                        switch (fi18 << ((4294967295) & 63)) {
                                                                        case 606:
                                                                            v1 = append(v1, "\n\t")
                                                                        case 32768:
                                                                            if (((42 < uint(v60)) || (0.1 == v3)) && ((((v5 * v5) < (v5 & 75)) || false) || (v5 == (^(v5))))) {
                                                                                v4 = -32768
                                                                            }
                                                                        }
                                                                        _ = v62
                                                                    case -2147483648:
                                                                        var v63 float32 = 1e10
                                                                        _ = v63
                                                                    default:
                                                                        var v64 float32 = 0.1
                                                                        _ = v64
                                                                    }
                                                                    var v65 int8 = 42
                                                                    _ = v65
                                                                }
                                                            case "foo bar":
                                                                var v66 int16 = v4
                                                                _ = v66
                                                            }
                                                            v67 := GenS2{F0: 362}
                                                            var v68 bool = (!((v37 + (("foo bar" + v24) + v37)) >= v37))
                                                            if ((v36 ^ v10) >= v60) {
                                                                if v68 {
                                                                    v69 := GenS5{F0: v24, F1: (v4 & (v4 & (v4 | int16(v8)))), F2: v8}
                                                                    v70 := GenS2{F0: 371}
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                                switch fi18 {
                                                                case -2147483648:
                                                                    v71 := GenS4{F0: v31}
                                                                    _ = v71
                                                                case 127:
                                                                    fmt.Println(v5)
                                                                }
                                                                var v72 int32 = -295
                                                                _ = v72
                                                            } else {
                                                                if true {
                                                                    fmt.Println(v68, v58, v31)
                                                                    var v73 int16 = (-(v4))
                                                                    _ = v73
                                                                }
                                                                v10 -= 65535
                                                                if ((v5 & (v5 * v5)) <= (v5 + 234)) {
                                                                    fmt.Println(fi35)
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    var v74 string = (v37 + "")
                                                                    _ = v74
                                                                } else {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    fmt.Println(fi35)
                                                                }
                                                                v75 := GenS2{F0: (v61 % ((v31) | 1))}
                                                                v60 -= (v36 ^ 340)
                                                                _ = v75
                                                            }
                                                            _ = v61
                                                            _ = v67
                                                            _ = v68
                                                        } else {
                                                            v76 := GenS7{F0: 100, F1: 127}
                                                            var v77 uint = 65535
                                                            fmt.Println(v9, v4)
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                        _ = v60
                                                    } else {
                                                        var v78 uint64 = v8
                                                        v79 := GenS2{F0: (^(v9))}
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                }
                                                _ = v58
                                            }
                                            switch (fi18 | fi35) {
                                            case -9223372036854775808:
                                                switch (fi18 << ((255) & 63)) {
                                                case -2147483648:
                                                    var v80 int8 = 0
                                                    _ = v80
                                                }
                                                if true {
                                                    var v81 int = (fi18 * (fi18 * fi35))
                                                    _ = v81
                                                } else {
                                                    if false {
                                                        v4 = v4
                                                        fmt.Println(fi18, v3, v8)
                                                        v36 = (v10 - 1)
                                                    } else {
                                                        var v82 int64 = int64(v4)
                                                        if true {
                                                            switch (v37 + v37) {
                                                            case "hello":
                                                                var v83 int = (fi18 % ((fi35) | 1))
                                                                _ = v83
                                                            default:
                                                                v36 += 2
                                                            }
                                                        } else {
                                                            v84 := GenS8{F0: (v9 * v9), F1: 779}
                                                            _ = v84
                                                        }
                                                        _ = v82
                                                    }
                                                    v9 = 42
                                                }
                                            case 0:
                                                if false {
                                                    v3 = v3
                                                    var v85 uint16 = 42
                                                    _ = v85
                                                }
                                            default:
                                                var v86 uint64 = v8
                                                _ = v86
                                            }
                                            v87 := GenS9{F0: v8, F1: 0}
                                            fmt.Println(fi18)
                                            v9 = v9
                                            _ = v87
                                        } else {
                                            switch fi35 {
                                            case 1:
                                                var v88 float64 = (3.14159 - (v3 - v3))
                                                switch v24 {
                                                case "\n\t":
                                                    switch 42 {
                                                    case 32767:
                                                        if (42 < uint(v8)) {
                                                            var v89 uint32 = 4294967295
                                                            _ = v89
                                                        } else {
                                                            var v90 uint16 = 42
                                                            _ = v90
                                                        }
                                                    case -2147483648:
                                                        var v91 bool = (106 >= 127)
                                                        _ = v91
                                                    }
                                                case "foo bar":
                                                    if len(v2) > 0 {
                                                        v2[0] = v24
                                                    }
                                                }
                                                v92 := GenS3{F0: uint(v5), F1: 255}
                                                _ = v88
                                                _ = v92
                                            default:
                                                switch fi35 {
                                                case -629:
                                                    v2 = append(v2, "\n\t")
                                                case 32767:
                                                    v93 := GenS10{F0: (fi35 / ((int(v31)) | 1)), F1: 42}
                                                    _ = v93
                                                case 2147483647:
                                                    switch v37 {
                                                    case "a":
                                                        v31 -= (v9 >> ((1) & 63))
                                                    case "foo bar":
                                                        var v94 int16 = v4
                                                        _ = v94
                                                    case "":
                                                        var v95 uint16 = 665
                                                        _ = v95
                                                    }
                                                default:
                                                    if (!false) {
                                                        v37 = v24
                                                        v5 -= 0
                                                    }
                                                    switch fi18 {
                                                    case 128:
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v96 := GenS9{F0: 312, F1: 7}
                                                            _ = v96
                                                        default:
                                                            switch "hello" {
                                                            case "a":
                                                                v97 := GenS7{F0: int8(v8), F1: (fi35 % (((^(fi18))) | 1))}
                                                                _ = v97
                                                            default:
                                                                _ = len(v1)
                                                                v36 = 528
                                                            }
                                                        }
                                                    default:
                                                        if ("" >= v24) {
                                                            var v98 string = v37
                                                            v24 = "\n\t"
                                                            _ = v98
                                                        } else {
                                                            var v99 int32 = 976
                                                            switch fi35 {
                                                            case 9223372036854775807:
                                                                v100 := GenS10{F0: (fi18 + fi18), F1: 32767}
                                                                _ = v100
                                                            default:
                                                                v37 = ((("foo bar" + v24) + (v37 + ("" + ("a" + "\n\t")))) + (("foo bar" + (v37 + (v24 + (v24 + "hello")))) + (("foo bar" + (v24 + (v37 + ""))) + "")))
                                                            }
                                                            var v101 string = ("foo bar" + (((((("\n\t" + ((v24 + "hello") + ("foo bar" + v37))) + "hello") + v37) + (v24 + "hello")) + "a") + v24))
                                                            _ = v101
                                                            _ = v99
                                                        }
                                                    }
                                                }
                                            }
                                            if (!true) {
                                                v102 := GenS1{F0: (fi18 % (((fi18 & (fi35 + (fi35 >> ((42) & 63))))) | 1)), F1: ("a" + ("" + "")), F2: v8}
                                                _ = v102
                                            } else {
                                                fmt.Println(v8, v8)
                                            }
                                            v8 -= v8
                                        }
                                        var v103 int8 = -128
                                        v36 -= v10
                                        _ = v103
                                    } else {
                                        if len(v6) > 0 {
                                            v6[0] = v36
                                        }
                                    }
                                    _ = v37
                                }
                                _ = fi35
                                _ = v36
                            }
                        case "a":
                            for fi104 := 0; fi104 < 11; fi104++ {
                                v105 := GenS1{F0: 371, F1: v24, F2: 18446744073709551615}
                                fmt.Println(v5)
                                v106 := GenS11{F0: (^(v10)), F1: fi18}
                                switch (-(fi18)) {
                                case -746:
                                    if ((("foo bar" != ((v24 + ((v24 + (v24 + (v24 + v24))) + (v24 + (v24 + ((("a" + v24) + (("foo bar" + v24) + v24)) + v24))))) + "a")) || (0 != 42)) && (2 < (v5 / ((42) | 1)))) {
                                        if ((v4 * (v4 << ((uint(v10)) & 63))) != -1) {
                                            var v107 int64 = (v9 + int64(v8))
                                            v108 := GenS11{F0: (v10 ^ (v10 | 0)), F1: fi18}
                                            v109 := GenS12{F0: (v5 >> ((255) & 63)), F1: 1}
                                            fmt.Println(v107, v10)
                                            fmt.Println(v3)
                                            _ = v107
                                            _ = v108
                                            _ = v109
                                        }
                                        v5 -= 128
                                    }
                                case 4294967295:
                                    v110 := GenS13{F0: v31, F1: 32767, F2: 0}
                                    _ = v110
                                }
                                _ = fi104
                                _ = v105
                                _ = v106
                            }
                        case "\n\t":
                            for fi111 := 0; fi111 < 8; fi111++ {
                                v10 = v10
                                v31 -= v9
                                fmt.Println(v5, fi111)
                                _ = fi111
                            }
                        }
                        _ = v31
                    }
                    _ = v24
                } else {
                    v8 = (v8 + v8)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    if (255 >= 65535) {
                        var v112 int32 = -2147483648
                        _ = v112
                    }
                }
                _ = v23
            }
            _ = fi18
        }
    }
    var v113 int16 = v4
    var v114 uint = 593
    v115 := GenS14{F0: 2}
    v113 -= (v4 % ((v4) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v113:", v113)
    fmt.Println("v114:", v114)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v113
    _ = v114
    _ = v115
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
