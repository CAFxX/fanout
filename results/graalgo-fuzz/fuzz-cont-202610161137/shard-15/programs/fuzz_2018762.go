package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int }
type GenS2 struct { F0 uint16; F1 int64; F2 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint; F1 int64 }

func main() {
    var v1 float64 = 0.0
    v2 := []string{("" + ("" + "a"))}
    var v3 uint8 = 127
    var v4 uint8 = (v3 % ((v3) | 1))
    var v5 int8 = 55
    v4 -= v4
    v6 := GenS1{F0: 469, F1: 1, F2: 42}
    switch 127 {
    case 4294967295:
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := GenS1{F0: 611, F1: 804, F2: fi7}
            _ = fi7
            _ = v8
        }
    case -2147483648:
        if false {
            v4 -= v4
            v4 -= (v3 - 42)
            v3 = v3
        } else {
            v3 -= (v3 ^ v4)
        }
    case -194:
        var v9 float32 = 0.0
        _ = v9
    default:
        var v10 bool = true
        _ = v10
    }
    v4 = v4
    v4 = 2
    switch 32768 {
    case -528:
        v11 := GenS1{F0: 264, F1: -1, F2: -1}
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        switch ("foo bar" + "") {
        case "":
            if true {
                v2 = append(v2, "a")
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 uint32 = 1
                    switch "foo bar" {
                    case "hello":
                        var v15 int64 = 42
                        _ = v15
                    case "":
                        v16 := GenS1{F0: 255, F1: -32768, F2: 9223372036854775807}
                        _ = v16
                    default:
                        v5 += (v5 % ((v5) | 1))
                    }
                    _ = fi13
                    _ = v14
                }
                var v17 uint32 = 2
                _ = v17
            } else {
                for fi18 := 0; fi18 < 20; fi18++ {
                    v5 -= (v5 * -105)
                    v4 = v4
                    _ = fi18
                }
                v2 = append(v2, "")
            }
        case "\n\t":
            v19 := GenS1{F0: 2, F1: 1, F2: (fi12 - (fi12 ^ fi12))}
            _ = v19
        default:
            v5 += (v5 ^ (v5 ^ v5))
        }
        if (v1 != v1) {
            v5 = v5
            if (v1 <= v1) {
                v4 = 2
                if true {
                    switch ("\n\t" + "") {
                    case "\n\t":
                        var v20 int8 = -1
                        _ = v20
                    case "":
                        v21 := GenS2{F0: 0, F1: 0, F2: 979}
                        _ = v21
                    case "foo bar":
                        var v22 float64 = (v1 + (float64(1e10) * float64(1.5)))
                        _ = v22
                    default:
                        fmt.Println(v1, v4, fi12)
                    }
                } else {
                    var v23 uint64 = uint64(v3)
                    v24 := GenS2{F0: 157, F1: 0, F2: 42}
                    if ((18446744073709551615 != 18446744073709551615) || false) {
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                    } else {
                        var v25 int64 = 293
                        v26 := GenS3{F0: v3}
                        _ = v25
                        _ = v26
                    }
                    _ = v23
                    _ = v24
                }
                for fi27 := 0; fi27 < 13; fi27++ {
                    var v28 float32 = (3.14159 / ((-2.25 * 0.1) - (1.5 / 1.0)))
                    if false {
                        v29 := GenS1{F0: 835, F1: 0, F2: (fi12 + (fi12 ^ (fi27 ^ fi27)))}
                        switch 128 {
                        case 42:
                            v4 -= (^(v3))
                        case -1000:
                            if (("hello" != ("foo bar" + "foo bar")) || false) {
                                var v30 int32 = 144
                                switch (("a" + "\n\t") + "\n\t") {
                                case "foo bar":
                                    v30 += 2147483647
                                }
                                _ = v30
                            }
                        }
                        _ = v29
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi27
                    _ = v28
                }
            }
        } else {
            for fi31 := 0; fi31 < 3; fi31++ {
                v32 := GenS2{F0: 65535, F1: -161, F2: 32767}
                _ = fi31
                _ = v32
            }
        }
        _ = fi12
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        v1 = (3.14159 - (v1 / (((1e10 + v1) - v1) * ((v1 / v1) / (0.0 + v1)))))
        for fi34 := 0; fi34 < 10; fi34++ {
            fmt.Println(v5, v5)
            _ = fi34
        }
        _ = fi33
    }
    switch "foo bar" {
    case "hello":
        var v35 float64 = 1e10
        _ = v35
    case "\n\t":
        if true {
            fmt.Println(v5, v4)
            v3 += v3
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            if (-1 < -9223372036854775808) {
                fmt.Println(v1, v4)
                if false {
                    if false {
                        switch (((("a" + "") + (("" + ("a" + "\n\t")) + "a")) + "") + "foo bar") {
                        case "a":
                            v5 = 0
                        case "foo bar":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        switch "hello" {
                        case "a":
                            v36 := GenS3{F0: 1}
                            _ = v36
                        case "":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            v5 = v5
                        }
                        var v37 float32 = (0.1 / (float32(v1) - 3.14159))
                        _ = v37
                    } else {
                        v1 = v1
                        switch 2147483647 {
                        case -50:
                            var v38 int16 = 1
                            _ = v38
                        case 4294967295:
                            for fi39 := 0; fi39 < 20; fi39++ {
                                for fi40 := 0; fi40 < 14; fi40++ {
                                    var v41 string = (("\n\t" + "a") + "")
                                    v42 := GenS4{F0: fi40, F1: 65535, F2: 2}
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                }
                                for fi43 := 0; fi43 < 9; fi43++ {
                                    var v44 uint32 = 2
                                    _ = fi43
                                    _ = v44
                                }
                                switch "foo bar" {
                                case "":
                                    v45 := GenS5{F0: 1, F1: 883}
                                    _ = v45
                                case "foo bar":
                                    fmt.Println(fi39)
                                case "a":
                                    v1 = (-2.25 - v1)
                                default:
                                    if len(v2) > 0 {
                                        v2[0] = "a"
                                    }
                                }
                                _ = fi39
                            }
                        default:
                            for fi46 := 0; fi46 < 9; fi46++ {
                                var v47 uint32 = 2
                                v48 := GenS4{F0: fi46, F1: 1, F2: 65535}
                                if (float32(-2.25) <= 0.1) {
                                    v49 := GenS6{F0: 943}
                                    _ = v49
                                }
                                _ = fi46
                                _ = v47
                                _ = v48
                            }
                        }
                    }
                    for fi50 := 0; fi50 < 16; fi50++ {
                        var v51 string = "\n\t"
                        var v52 uint = 1
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    var v53 float32 = 0.1
                    _ = v53
                } else {
                    if (!false) {
                        switch 127 {
                        case 128:
                            fmt.Println(v4)
                        }
                        v1 = 0.0
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        var v54 float32 = 0.1
                        var v55 int8 = 127
                        fmt.Println(v5, v3)
                        _ = v54
                        _ = v55
                    }
                    var v56 int16 = 1
                    var v57 int32 = 1
                    _ = v56
                    _ = v57
                }
                for fi58 := 0; fi58 < 8; fi58++ {
                    for fi59 := 0; fi59 < 2; fi59++ {
                        fmt.Println(fi59)
                        var v60 uint8 = v3
                        switch 32768 {
                        case 667:
                            var v61 uint = 42
                            _ = v61
                        default:
                            var v62 float32 = 1e10
                            if false {
                                var v63 uint = 65535
                                switch "a" {
                                case "a":
                                    v64 := GenS7{F0: 9223372036854775807}
                                    _ = v64
                                case "hello":
                                    var v65 uint64 = 696
                                    _ = v65
                                case "":
                                    if true {
                                        var v66 int8 = v5
                                        if false {
                                            var v67 uint64 = 233
                                            _ = v67
                                        }
                                        fmt.Println(v4, v60, v1)
                                        v5 = v5
                                        var v68 uint8 = (v60 & (^(v4)))
                                        _ = v66
                                        _ = v68
                                    } else {
                                        if false {
                                            v69 := GenS8{F0: v5}
                                            if true {
                                                if true {
                                                    v62 = 1.5
                                                }
                                                var v70 string = "a"
                                                if true {
                                                    if (true && true) {
                                                        if (-899 <= 76) {
                                                            v2 = append(v2, "hello")
                                                            var v71 int16 = 1
                                                            if true {
                                                                if true {
                                                                    var v72 int32 = 42
                                                                    if false {
                                                                        v72 = (v72 & v72)
                                                                        v3 -= uint8(v63)
                                                                    } else {
                                                                        switch (-(fi59)) {
                                                                        case 681:
                                                                            fmt.Println(v72, v5)
                                                                        case 537:
                                                                            fmt.Println(v4, fi58)
                                                                        }
                                                                        v73 := GenS9{F0: 470}
                                                                        v74 := GenS10{F0: (^(v63))}
                                                                        _ = v73
                                                                        _ = v74
                                                                    }
                                                                    v4 = 2
                                                                    switch fi58 {
                                                                    case 127:
                                                                        v71 = v71
                                                                    case 446:
                                                                        if (v63 != 42) {
                                                                            v70 = ""
                                                                            if len(v2) > 0 {
                                                                                v2[0] = v70
                                                                            }
                                                                            v62 -= v62
                                                                        } else {
                                                                            if true {
                                                                                v70 = "hello"
                                                                                var v75 int32 = v72
                                                                                _ = v75
                                                                            }
                                                                            if true {
                                                                                v3 = v4
                                                                                v62 = (v62 / v62)
                                                                                v5 -= (^(v5))
                                                                                if false {
                                                                                    switch fi58 {
                                                                                    case 32768:
                                                                                        var v76 uint64 = 631
                                                                                        _ = v76
                                                                                    default:
                                                                                        v70 = v70
                                                                                    }
                                                                                    if true {
                                                                                        var v77 float32 = v62
                                                                                        v72 += (^(v72))
                                                                                        v62 = 1e10
                                                                                        _ = v77
                                                                                    }
                                                                                    var v78 int8 = -66
                                                                                    var v79 string = (v70 + v70)
                                                                                    _ = v78
                                                                                    _ = v79
                                                                                }
                                                                            }
                                                                        }
                                                                    case 0:
                                                                        var v80 uint8 = (v3 + (v3 - (v4 >> ((v63) & 63))))
                                                                        _ = v80
                                                                    }
                                                                    var v81 int8 = int8(v72)
                                                                    switch v70 {
                                                                    case "":
                                                                        v60 = (v3 % (((v3 / ((v3) | 1))) | 1))
                                                                    }
                                                                    fmt.Println(v63, v70, v4)
                                                                    _ = v72
                                                                    _ = v81
                                                                }
                                                            } else {
                                                                switch fi58 {
                                                                case 60:
                                                                    v63 = (v63 >> ((v63) & 63))
                                                                }
                                                                var v82 int8 = v5
                                                                _ = v82
                                                            }
                                                            v83 := GenS1{F0: 0, F1: -1, F2: fi59}
                                                            var v84 uint64 = 42
                                                            _ = v71
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                        fmt.Println(fi59, fi58)
                                                        _ = len(v2)
                                                    }
                                                    if true {
                                                        var v85 int16 = -240
                                                        _ = v85
                                                    }
                                                } else {
                                                    switch "hello" {
                                                    case "":
                                                        v3 = v3
                                                        var v86 int32 = -2147483648
                                                        _ = v86
                                                    case "hello":
                                                        if (float64(v62) <= 1e10) {
                                                            var v87 float64 = v1
                                                            switch v70 {
                                                            case "":
                                                                var v88 uint32 = 810
                                                                var v89 int32 = -736
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                            _ = v87
                                                        }
                                                    }
                                                }
                                                _ = v70
                                            }
                                            v1 = v1
                                            v5 += v5
                                            var v90 uint = (v63 - v63)
                                            if len(v2) > 0 {
                                                v2[0] = ""
                                            }
                                            var v91 uint16 = 2
                                            _ = v69
                                            _ = v90
                                            _ = v91
                                        }
                                        var v92 float64 = 1.5
                                        fmt.Println(v60, fi58, v60)
                                        _ = v92
                                    }
                                default:
                                    if false {
                                        var v93 float64 = 1e10
                                        if (1 < 65535) {
                                            var v94 int64 = -1
                                            if true {
                                                var v95 float64 = 3.14159
                                                _ = v95
                                            } else {
                                                switch 128 {
                                                case 42:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case -9223372036854775808:
                                                    v4 = v60
                                                case 2147483647:
                                                    _ = len(v2)
                                                }
                                                if false {
                                                    v96 := GenS11{F0: (^(v63)), F1: v94}
                                                    fmt.Println(v3, v63, v4)
                                                    v97 := GenS3{F0: (v4 * v3)}
                                                    _ = v96
                                                    _ = v97
                                                }
                                                var v98 int64 = (v94 * 1)
                                                _ = v98
                                            }
                                            _ = v94
                                        }
                                        _ = v93
                                    } else {
                                        var v99 int32 = 0
                                        if true {
                                            fmt.Println(v3, fi59, v63)
                                            v1 -= float64((v62 / v62))
                                        } else {
                                            v100 := GenS3{F0: v60}
                                            var v101 int32 = (^(v99))
                                            v62 += ((v62 / v62) - (v62 / v62))
                                            _ = v100
                                            _ = v101
                                        }
                                        _ = v99
                                    }
                                }
                                _ = v63
                            } else {
                                v60 = v4
                                v3 -= v60
                            }
                            if false {
                                var v102 uint64 = 2
                                _ = v102
                            } else {
                                var v103 uint32 = 2147483648
                                v60 = v4
                                v103 += 908
                                v60 -= (v4 / (((^(v60))) | 1))
                                _ = v103
                            }
                            var v104 float32 = 1.5
                            _ = v104
                            _ = v62
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if ((!(114 == -2147483648)) && false) {
                            switch ("hello" + "hello") {
                            case "a":
                                v60 = 255
                            default:
                                if false {
                                    v2 = append(v2, "a")
                                    fmt.Println(v4)
                                }
                            }
                        }
                        _ = fi59
                        _ = v60
                    }
                    v105 := GenS1{F0: 2, F1: 621, F2: 128}
                    _ = fi58
                    _ = v105
                }
            }
        }
    default:
        var v106 uint16 = 32767
        _ = v106
    }
    for fi107 := 0; fi107 < 10; fi107++ {
        fmt.Println(v4, v4, v1)
        _ = fi107
    }
    var v108 int = 32767
    v5 -= 42
    if len(v2) > 0 {
        v2[0] = ("hello" + "hello")
    }
    for fi109 := 0; fi109 < 16; fi109++ {
        var v110 uint16 = uint16(v3)
        var v111 int8 = v5
        _ = fi109
        _ = v110
        _ = v111
    }
    if false {
        var v112 uint32 = 2147483647
        switch "\n\t" {
        case "":
            var v113 uint32 = 2
            _ = v113
        case "hello":
            for fi114 := 0; fi114 < 14; fi114++ {
                for fi115 := 0; fi115 < 11; fi115++ {
                    v2 = append(v2, "a")
                    var v116 string = "foo bar"
                    fmt.Println(v5)
                    _ = fi115
                    _ = v116
                }
                switch fi114 {
                case 0:
                    v1 += (3.14159 - v1)
                }
                _ = fi114
            }
        case "foo bar":
            if ((v108 | v108) < (v108 / (((v108 >> ((461) & 63))) | 1))) {
                fmt.Println(v5, v4, v5)
                for fi117 := 0; fi117 < 8; fi117++ {
                    fmt.Println(v5, v4)
                    _ = fi117
                }
            } else {
                var v118 float32 = (1.5 + float32(v1))
                v108 = (v108 - v108)
                switch ("\n\t" + "\n\t") {
                case "foo bar":
                    for fi119 := 0; fi119 < 15; fi119++ {
                        var v120 int64 = 1
                        var v121 uint16 = 674
                        switch ((("a" + (("foo bar" + (("foo bar" + "") + ("\n\t" + "a"))) + ((("foo bar" + "a") + "foo bar") + ((("\n\t" + ("hello" + "")) + ("foo bar" + ("hello" + "hello"))) + "")))) + "") + ("" + "a")) {
                        case "\n\t":
                            fmt.Println(v108, v112, v4)
                        case "foo bar":
                            for fi122 := 0; fi122 < 17; fi122++ {
                                fmt.Println(v5, v112)
                                _ = fi122
                            }
                        case "hello":
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                        default:
                            var v123 uint = 18446744073709551615
                            _ = v123
                        }
                        _ = fi119
                        _ = v120
                        _ = v121
                    }
                }
                _ = v118
            }
        }
        if len(v2) > 0 {
            v2[0] = ("a" + "hello")
        }
        _ = v112
    }
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
