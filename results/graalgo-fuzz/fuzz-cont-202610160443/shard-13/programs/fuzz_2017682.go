package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 string; F1 uint8; F2 uint64 }
type GenS9 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int8; F1 string }
type GenS12 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS13 struct { F0 uint8 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]string)
    var v3 uint8 = 192
    var v4 int64 = 0
    v5 := make(map[string]uint32)
    var v6 int64 = -506
    v3 = (v3 * 233)
    v7 := GenS1{F0: 4294967295, F1: 127, F2: v1}
    var v8 uint64 = 4294967295
    switch -9223372036854775808 {
    case 32768:
        v5["a"] = 42
    case 127:
        v3 = v3
    case 90:
        v9 := GenS1{F0: 2, F1: uint8(v1), F2: 2}
        _ = v9
    }
    v10 := []int{957, 9223372036854775807, -2147483648}
    switch 9223372036854775807 {
    case 9223372036854775807:
        for fi11 := 0; fi11 < 17; fi11++ {
            delete(v5, ("foo bar" + ("" + (("\n\t" + "foo bar") + "foo bar"))))
            switch int(v6) {
            case -2147483648:
                v10 = append(v10, -2147483648)
            default:
                v1 -= 42
            }
            v3 += v3
            _ = fi11
        }
    case 42:
        var v12 int64 = v4
        _ = v12
    default:
        if ((^(v6)) > v4) {
            var v13 string = ("\n\t" + ((("hello" + "foo bar") + "hello") + ""))
            for fi14 := 0; fi14 < 6; fi14++ {
                for fi15 := 0; fi15 < 13; fi15++ {
                    var v16 int8 = -1
                    switch fi14 {
                    case -9223372036854775808:
                        v6 = (v4 % ((844) | 1))
                    case 9223372036854775807:
                        switch fi14 {
                        case 128:
                            if true {
                                v1 -= 18446744073709551615
                                v3 += 42
                                v17 := GenS1{F0: v8, F1: (v3 >> ((v1) & 63)), F2: 2}
                                fmt.Println(v8, v13)
                                v18 := GenS1{F0: (v8 % (((^(v8))) | 1)), F1: 128, F2: v1}
                                var v19 bool = ((fi15 | (^(fi15))) <= (fi14 << ((v1) & 63)))
                                v20 := GenS2{F0: 0}
                                v21 := GenS3{F0: 994}
                                _ = v17
                                _ = v18
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                        case -604:
                            if false {
                                if true {
                                    v3 -= 255
                                }
                                var v22 uint64 = 2
                                _ = v22
                            } else {
                                var v23 bool = (v8 != 4294967295)
                                var v24 uint = v1
                                _ = v23
                                _ = v24
                            }
                        case -9223372036854775808:
                            v3 = 255
                        }
                    default:
                        _ = len(v2)
                        fmt.Println(v4, v3)
                        if ((!(!false)) && false) {
                            switch fi14 {
                            case 128:
                                var v25 float32 = 0.0
                                _ = v25
                            case -489:
                                v2[(fi14 / ((-1) | 1))] = "foo bar"
                            }
                            var v26 int8 = v16
                            v27 := GenS3{F0: 65535}
                            _ = v26
                            _ = v27
                        }
                        var v28 int = (fi14 * (fi14 & fi15))
                        _ = v28
                    }
                    _ = fi15
                    _ = v16
                }
                for fi29 := 0; fi29 < 1; fi29++ {
                    v4 = 497
                    _ = fi29
                }
                var v30 uint = (^(v1))
                _ = fi14
                _ = v30
            }
            for fi31 := 0; fi31 < 17; fi31++ {
                for fi32 := 0; fi32 < 18; fi32++ {
                    v8 = v8
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = fi32
                }
                fmt.Println(v6)
                _ = fi31
            }
            _ = v13
        }
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        switch "a" {
        case "foo bar":
            var v34 int16 = -654
            _ = v34
        case "a":
            for fi35 := 0; fi35 < 20; fi35++ {
                var v36 float64 = -2.25
                _ = fi35
                _ = v36
            }
        }
        for fi37 := 0; fi37 < 17; fi37++ {
            var v38 uint64 = v8
            if false {
                var v39 string = ("" + "hello")
                v40 := GenS1{F0: (^(v38)), F1: (v3 / ((v3) | 1)), F2: (v1 << (((v1 ^ 433)) & 63))}
                var v41 bool = (3.14159 > (3.14159 + -2.25))
                _ = len(v10)
                _ = v39
                _ = v40
                _ = v41
            }
            _ = fi37
            _ = v38
        }
        if true {
            var v42 bool = (v1 <= (^(v1)))
            _ = v42
        }
        _ = fi33
    }
    if false {
        for fi43 := 0; fi43 < 3; fi43++ {
            switch ("a" + "") {
            case "foo bar":
                fmt.Println(fi43, v6)
            case "":
                fmt.Println(v8, v3, v1)
            }
            _ = fi43
        }
    } else {
        switch ("foo bar" + "a") {
        case "\n\t":
            v3 += (^(v3))
        default:
            var v44 int8 = 100
            _ = v44
        }
        var v45 int = -9223372036854775808
        _ = v45
    }
    var v46 bool = (1.5 < -2.25)
    v47 := GenS1{F0: (v8 << ((0) & 63)), F1: (v3 << ((v1) & 63)), F2: 99}
    if v46 {
        var v48 int8 = 127
        if true {
            for fi49 := 0; fi49 < 15; fi49++ {
                v4 -= v4
                _ = fi49
            }
        }
        fmt.Println(v48, v6, v3)
        _ = v48
    } else {
        v50 := v5[(("\n\t" + "a") + ("a" + "\n\t"))]
        v2[32767] = "hello"
        var v51 int16 = -1
        _ = v50
        _ = v51
    }
    v52 := GenS4{F0: 1, F1: (v6 * (v6 / ((v6) | 1)))}
    switch "hello" {
    case "a":
        if v46 {
            v46 = true
        } else {
            fmt.Println(v3)
            var v53 uint64 = 641
            switch 32767 {
            case -731:
                if ((v6 == v4) && v46) {
                    v54 := GenS5{F0: int64(v53), F1: 0}
                    if v46 {
                        v55 := GenS4{F0: 835, F1: int64(v53)}
                        v3 = v3
                        var v56 uint16 = 203
                        _ = v55
                        _ = v56
                    } else {
                        if v46 {
                            var v57 uint64 = v8
                            _ = v57
                        }
                    }
                    for fi58 := 0; fi58 < 19; fi58++ {
                        v59 := v2[4294967295]
                        _ = fi58
                        _ = v59
                    }
                    _ = v54
                } else {
                    switch "" {
                    case "":
                        v60 := v2[0]
                        _ = v60
                    case "a":
                        v8 = (v53 | v53)
                    case "\n\t":
                        _ = len(v10)
                    default:
                        for fi61 := 0; fi61 < 4; fi61++ {
                            if v46 {
                                if ((!true) || true) {
                                    v62 := GenS6{F0: 32767, F1: 32767}
                                    if true {
                                        fmt.Println(v6, v4)
                                        v63 := GenS2{F0: -1}
                                        _ = v63
                                    } else {
                                        v6 -= -9223372036854775808
                                        v10 = append(v10, 9223372036854775807)
                                        if false {
                                            var v64 bool = true
                                            _ = v64
                                        }
                                    }
                                    for fi65 := 0; fi65 < 18; fi65++ {
                                        var v66 uint = v1
                                        var v67 int16 = 195
                                        switch "\n\t" {
                                        case "hello":
                                            v46 = true
                                        case "":
                                            fmt.Println(v66)
                                        default:
                                            var v68 int64 = 0
                                            _ = v68
                                        }
                                        _ = fi65
                                        _ = v66
                                        _ = v67
                                    }
                                    _ = v62
                                } else {
                                    switch "a" {
                                    case "a":
                                        var v69 uint8 = (v3 << ((v1) & 63))
                                        _ = v69
                                    case "\n\t":
                                        var v70 float32 = 1.5
                                        _ = v70
                                    default:
                                        switch ("" + "a") {
                                        case "foo bar":
                                            var v71 int = fi61
                                            _ = v71
                                        }
                                    }
                                }
                            }
                            _ = fi61
                        }
                    }
                    switch "a" {
                    case "hello":
                        for fi72 := 0; fi72 < 4; fi72++ {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            if false {
                                if (v3 < 128) {
                                    switch fi72 {
                                    case 128:
                                        v1 -= 670
                                    }
                                    if (v6 > v4) {
                                        v73 := GenS7{F0: (fi72 + 619), F1: 42}
                                        _ = v73
                                    } else {
                                        switch "a" {
                                        case "a":
                                            var v74 int = 0
                                            _ = v74
                                        }
                                        fmt.Println(v3, fi72)
                                        v75 := GenS8{F0: "foo bar", F1: (^(v3)), F2: (v53 & (v53 ^ v8))}
                                        _ = v75
                                    }
                                    var v76 float64 = 1.5
                                    _ = v76
                                } else {
                                    for fi77 := 0; fi77 < 11; fi77++ {
                                        if true {
                                            switch (-(fi77)) {
                                            case 127:
                                                if v46 {
                                                    var v78 int32 = 42
                                                    if v46 {
                                                        switch "hello" {
                                                        case "hello":
                                                            switch 1 {
                                                            case 42:
                                                                var v79 uint8 = v3
                                                                _ = v79
                                                            case -2147483648:
                                                                if v46 {
                                                                    v4 += 1
                                                                    var v80 int32 = v78
                                                                    v10 = append(v10, fi72)
                                                                    var v81 int = fi72
                                                                    fmt.Println(v80, v3)
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        if true {
                                                                            switch ("foo bar" + "\n\t") {
                                                                            case "\n\t":
                                                                                var v82 int16 = 32767
                                                                                _ = v82
                                                                            }
                                                                            v78 = v80
                                                                            var v83 int8 = -128
                                                                            fmt.Println(v83)
                                                                            var v84 int16 = int16(fi77)
                                                                            fmt.Println(v4, v83, v6)
                                                                            _ = v83
                                                                            _ = v84
                                                                        }
                                                                        fmt.Println(v53, v53, fi72)
                                                                        if (!true) {
                                                                            if (v1 >= (v1 << ((v1) & 63))) {
                                                                                if true {
                                                                                    delete(v5, ("\n\t" + ("hello" + "hello")))
                                                                                    var v85 uint = 1
                                                                                    var v86 int16 = -32768
                                                                                    _ = v85
                                                                                    _ = v86
                                                                                } else {
                                                                                    v81 -= fi77
                                                                                    v78 += (-(v78))
                                                                                    switch "hello" {
                                                                                    case "foo bar":
                                                                                        delete(v5, ("foo bar" + "a"))
                                                                                    case "hello":
                                                                                        if (-221 == -32768) {
                                                                                            v3 -= 1
                                                                                            _ = len(v2)
                                                                                        }
                                                                                    case "":
                                                                                        delete(v2, (^(fi77)))
                                                                                    default:
                                                                                        _ = len(v10)
                                                                                    }
                                                                                    v87 := GenS8{F0: ("" + ""), F1: v3, F2: v8}
                                                                                    _ = v87
                                                                                }
                                                                                v88 := GenS9{F0: 753, F1: v80, F2: 1}
                                                                                _ = v88
                                                                            } else {
                                                                                v81 = int(v3)
                                                                                var v89 float32 = float32(1e10)
                                                                                _ = v89
                                                                            }
                                                                            switch "" {
                                                                            case "a":
                                                                                v1 += v1
                                                                            case "\n\t":
                                                                                if v46 {
                                                                                    var v90 int8 = -128
                                                                                    v91 := GenS7{F0: (fi77 / ((32768) | 1)), F1: 1}
                                                                                    switch (v81 / ((fi72) | 1)) {
                                                                                    case 42:
                                                                                        if len(v10) >= 2 {
                                                                                            v10 = v10[1:len(v10)]
                                                                                        }
                                                                                    case 1:
                                                                                        _ = len(v2)
                                                                                    case 805:
                                                                                        var v92 int64 = (v6 | v6)
                                                                                        _ = v92
                                                                                    }
                                                                                    _ = v90
                                                                                    _ = v91
                                                                                } else {
                                                                                    var v93 uint16 = 32767
                                                                                    _ = v93
                                                                                }
                                                                            }
                                                                            var v94 uint64 = 4294967295
                                                                            _ = v94
                                                                        } else {
                                                                            v95 := GenS4{F0: -1, F1: v4}
                                                                            var v96 int32 = (v78 >> ((v1) & 63))
                                                                            _ = v95
                                                                            _ = v96
                                                                        }
                                                                        fmt.Println(fi77, v6, v6)
                                                                        var v97 int = v81
                                                                        _ = v97
                                                                    case "":
                                                                        fmt.Println(fi72)
                                                                        switch v81 {
                                                                        case 32768:
                                                                            _ = len(v10)
                                                                        case 1:
                                                                            var v98 int = -896
                                                                            _ = v98
                                                                        case 929:
                                                                            v4 = (v4 << ((990) & 63))
                                                                        }
                                                                        if (v46 || ((302 != 2) || false)) {
                                                                            v3 -= (v3 << (((^(v1))) & 63))
                                                                            v1 += 2
                                                                        }
                                                                        if v46 {
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                _ = len(v2)
                                                                            default:
                                                                                v78 = v80
                                                                            }
                                                                            v99 := GenS4{F0: 0, F1: -9223372036854775808}
                                                                            var v100 float32 = float32(1.5)
                                                                            _ = v100
                                                                            _ = v99
                                                                        } else {
                                                                            _ = len(v5)
                                                                            v53 = v8
                                                                            fmt.Println(v46, v3, fi77)
                                                                            if (v46 && v46) {
                                                                                v53 -= v8
                                                                                v101 := GenS10{F0: 866}
                                                                                v5["hello"] = 525
                                                                                _ = v101
                                                                            } else {
                                                                                v102 := v5[""]
                                                                                if len(v10) >= 2 {
                                                                                    v10 = v10[1:len(v10)]
                                                                                }
                                                                                v103 := GenS3{F0: 1}
                                                                                _ = v102
                                                                                _ = v103
                                                                            }
                                                                        }
                                                                        if ((v78 | (-(v78))) == 79) {
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                v104 := GenS7{F0: 914, F1: uint32(v8)}
                                                                                _ = v104
                                                                            case "hello":
                                                                                if (!((v1 == 65535) && (true || ((v53 << (((v1 | 938)) & 63)) <= (v8 + (v8 * 1)))))) {
                                                                                    v80 += -2147483648
                                                                                    if len(v10) >= 2 {
                                                                                        v10 = v10[1:len(v10)]
                                                                                    }
                                                                                    fmt.Println(v8, v6)
                                                                                    v105 := GenS11{F0: 0, F1: "\n\t"}
                                                                                    _ = v105
                                                                                }
                                                                            case "foo bar":
                                                                                if v46 {
                                                                                    v1 = uint(v3)
                                                                                } else {
                                                                                    if (2 <= 0) {
                                                                                        if len(v10) > 0 {
                                                                                            v10[0] = fi77
                                                                                        }
                                                                                        var v106 int8 = -1
                                                                                        v107 := GenS4{F0: -1, F1: 1}
                                                                                        _ = v106
                                                                                        _ = v107
                                                                                    }
                                                                                    v2[fi77] = ("\n\t" + ("hello" + ("hello" + "\n\t")))
                                                                                    v10 = append(v10, 1)
                                                                                }
                                                                            }
                                                                            if v46 {
                                                                                fmt.Println(v8, v6, v6)
                                                                            }
                                                                            var v108 int16 = -833
                                                                            _ = v108
                                                                        } else {
                                                                            var v109 uint = (v1 >> ((18446744073709551615) & 63))
                                                                            fmt.Println(v78)
                                                                            _ = v109
                                                                        }
                                                                    default:
                                                                        if len(v10) > 0 {
                                                                            v10[0] = (fi77 + int(v6))
                                                                        }
                                                                    }
                                                                    _ = v80
                                                                    _ = v81
                                                                } else {
                                                                    switch (fi72 % (((-(fi72))) | 1)) {
                                                                    case 32767:
                                                                        var v110 string = "hello"
                                                                        _ = v110
                                                                    case 9223372036854775807:
                                                                        var v111 uint8 = v3
                                                                        _ = v111
                                                                    case -9223372036854775808:
                                                                        v4 = v4
                                                                    }
                                                                }
                                                            case 144:
                                                                v2[4294967295] = (("a" + "a") + "")
                                                            default:
                                                                var v112 uint = 1
                                                                _ = v112
                                                            }
                                                        }
                                                        delete(v5, "a")
                                                        var v113 uint8 = (v3 + (v3 >> (((v1 >> ((v1) & 63))) & 63)))
                                                        var v114 uint64 = v8
                                                        v10 = append(v10, (^(fi77)))
                                                        var v115 int64 = -9223372036854775808
                                                        _ = v113
                                                        _ = v114
                                                        _ = v115
                                                    }
                                                    var v116 uint16 = 540
                                                    switch "a" {
                                                    case "a":
                                                        if len(v10) > 0 {
                                                            v10[0] = (fi72 << ((v1) & 63))
                                                        }
                                                    case "\n\t":
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            var v117 float64 = (1.5 * 0.1)
                                                            _ = v117
                                                        default:
                                                            switch (fi77 | -389) {
                                                            case 2147483647:
                                                                if len(v10) >= 2 {
                                                                    v10 = v10[1:len(v10)]
                                                                }
                                                            default:
                                                                if (true && (((0.0 >= 1e10) || false) || v46)) {
                                                                    v118 := GenS12{F0: 477, F1: (v116 << ((255) & 63)), F2: 757}
                                                                    v119 := GenS2{F0: -32768}
                                                                    _ = len(v10)
                                                                    _ = v118
                                                                    _ = v119
                                                                }
                                                            }
                                                        }
                                                    case "foo bar":
                                                        var v120 uint64 = v53
                                                        _ = v120
                                                    default:
                                                        v53 = (v8 - (v53 & v8))
                                                    }
                                                    _ = v116
                                                    _ = v78
                                                }
                                                v53 -= 0
                                            }
                                        } else {
                                            var v121 uint = (v1 - v1)
                                            _ = v121
                                        }
                                        _ = fi77
                                    }
                                    var v122 string = "hello"
                                    _ = v122
                                }
                            } else {
                                v3 = 198
                            }
                            v8 = uint64(v1)
                            _ = fi72
                        }
                    case "":
                        v123 := GenS13{F0: uint8(v1)}
                        _ = v123
                    case "foo bar":
                        v10 = append(v10, int(v8))
                    }
                    for fi124 := 0; fi124 < 15; fi124++ {
                        fmt.Println(v4)
                        _ = fi124
                    }
                }
            default:
                v4 += -169
            }
            _ = v53
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v52
    _ = v6
    _ = v7
    _ = v8
}
