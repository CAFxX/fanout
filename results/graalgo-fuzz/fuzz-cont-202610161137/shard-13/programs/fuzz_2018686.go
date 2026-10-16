package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 string; F1 int8; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS4 struct { F0 string; F1 uint; F2 bool }
type GenS5 struct { F0 int16; F1 string; F2 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8; F1 string; F2 int16 }
type GenS8 struct { F0 uint8; F1 uint }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int8; F1 bool }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    v3 := []string{"hello", "a", "a"}
    v4 := make(map[int]int)
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]uint32)
    switch "foo bar" {
    case "hello":
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := GenS1{F0: 42, F1: (fi7 | 32767)}
            _ = fi7
            _ = v8
        }
    default:
        v9 := GenS1{F0: 105, F1: -685}
        _ = v9
    }
    var v10 uint8 = 2
    for fi11 := 0; fi11 < 6; fi11++ {
        v10 = (^(v10))
        var v12 string = ("foo bar" + "a")
        _ = fi11
        _ = v12
    }
    switch 42 {
    case 127:
        v13 := GenS1{F0: 65535, F1: 9223372036854775807}
        _ = v13
    case 42:
        v10 -= v10
    case 361:
        switch 42 {
        case -960:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 127:
            v10 -= v10
        case -556:
            v10 = v10
        default:
            switch "" {
            case "a":
                _ = len(v3)
            case "":
                for fi14 := 0; fi14 < 3; fi14++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi14
                }
            default:
                v15 := v2[""]
                _ = v15
            }
        }
    }
    var v16 int16 = 42
    fmt.Println(v10, v16)
    if (!false) {
        _ = len(v2)
        switch 127 {
        case 1:
            v16 = (v16 & v16)
        case 0:
            var v17 uint64 = 42
            _ = v17
        default:
            var v18 uint64 = 640
            _ = v18
        }
        fmt.Println(v10, v16)
    }
    if false {
        fmt.Println(v10, v10)
        v19 := GenS1{F0: uint(v10), F1: -9223372036854775808}
        _ = v19
    } else {
        if (true && (true || ("hello" < "foo bar"))) {
            for fi20 := 0; fi20 < 10; fi20++ {
                if false {
                    if false {
                        for fi21 := 0; fi21 < 6; fi21++ {
                            v22 := GenS2{F0: "", F1: 89, F2: 1}
                            _ = fi21
                            _ = v22
                        }
                        var v23 uint16 = 2
                        for fi24 := 0; fi24 < 2; fi24++ {
                            var v25 float32 = 0.0
                            _ = fi24
                            _ = v25
                        }
                        _ = v23
                    }
                    if (0.1 != 1e10) {
                        v16 = v16
                        var v26 uint32 = 307
                        fmt.Println(fi20, v16, v26)
                        _ = v26
                    } else {
                        var v27 uint8 = v10
                        _ = v27
                    }
                    switch "a" {
                    case "\n\t":
                        switch fi20 {
                        case -935:
                            v28 := GenS3{F0: (true && false), F1: 188, F2: 4294967295}
                            _ = v28
                        }
                    case "a":
                        for fi29 := 0; fi29 < 4; fi29++ {
                            switch ("" + "hello") {
                            case "hello":
                                if (18446744073709551615 <= 136) {
                                    if false {
                                        v3 = append(v3, "a")
                                    } else {
                                        var v30 int32 = -1
                                        var v31 bool = true
                                        _ = v30
                                        _ = v31
                                    }
                                    v32 := GenS1{F0: 974, F1: fi20}
                                    _ = v32
                                }
                            case "\n\t":
                                if (!(!(uint32(v10) != 1))) {
                                    fmt.Println(fi29)
                                }
                            default:
                                if true {
                                    switch "hello" {
                                    case "a":
                                        var v33 uint64 = 592
                                        _ = v33
                                    case "foo bar":
                                        v34 := GenS4{F0: "a", F1: 255, F2: (v10 > v10)}
                                        v35 := GenS2{F0: ("hello" + (("a" + (("foo bar" + "a") + "foo bar")) + "\n\t")), F1: 42, F2: 458}
                                        _ = v34
                                        _ = v35
                                    case "":
                                        if (65535 != 564) {
                                            if len(v3) > 0 {
                                                v3[0] = (("hello" + "foo bar") + "foo bar")
                                            }
                                            var v36 uint64 = 42
                                            _ = v36
                                        }
                                    }
                                } else {
                                    if true {
                                        switch "foo bar" {
                                        case "a":
                                            var v37 int32 = 42
                                            _ = v37
                                        }
                                        var v38 float32 = 1e10
                                        fmt.Println(v10)
                                        if false {
                                            if (v16 < (-(v16))) {
                                                fmt.Println(v10, fi29, fi20)
                                                fmt.Println(v16, v16)
                                            } else {
                                                if ((((((3.14159 * (1.5 + 1e10)) + 3.14159) - 0.0) / 1.0) * -2.25) > 1.5) {
                                                    if true {
                                                        var v39 bool = ("hello" < "")
                                                        _ = len(v3)
                                                        _ = v39
                                                    } else {
                                                        var v40 uint8 = (v10 / (((^(v10))) | 1))
                                                        _ = v40
                                                    }
                                                    if false {
                                                        fmt.Println(v16)
                                                    } else {
                                                        var v41 uint = 2
                                                        _ = v41
                                                    }
                                                    fmt.Println(fi29, v10, fi29)
                                                } else {
                                                    _ = len(v2)
                                                    v10 += v10
                                                }
                                                switch ((("hello" + ("foo bar" + (((("\n\t" + "foo bar") + "") + "") + ("\n\t" + ("a" + "foo bar"))))) + "foo bar") + "\n\t") {
                                                case "a":
                                                    switch "hello" {
                                                    case "hello":
                                                        switch "hello" {
                                                        case "foo bar":
                                                            v42 := v5[65535]
                                                            _ = v42
                                                        case "hello":
                                                            v38 += v38
                                                        }
                                                    case "a":
                                                        if true {
                                                            v38 += (v38 - v38)
                                                            v43 := GenS5{F0: v16, F1: "hello", F2: (fi29 + 127)}
                                                            _ = v43
                                                        }
                                                    case "":
                                                        var v44 uint64 = 2
                                                        _ = v44
                                                    }
                                                case "\n\t":
                                                    if true {
                                                        var v45 string = "hello"
                                                        _ = v45
                                                    } else {
                                                        fmt.Println(v16, v16, fi20)
                                                        v10 += 240
                                                    }
                                                case "":
                                                    var v46 uint8 = v10
                                                    _ = v46
                                                default:
                                                    switch "hello" {
                                                    case "\n\t":
                                                        switch fi29 {
                                                        case 32768:
                                                            var v47 int16 = 32767
                                                            _ = v47
                                                        default:
                                                            fmt.Println(fi29, v38, v10)
                                                        }
                                                    case "a":
                                                        var v48 float64 = (0.1 - (1e10 / (1.5 / (1e10 - ((0.1 * 0.0) + 0.0)))))
                                                        _ = v48
                                                    case "":
                                                        fmt.Println(fi29)
                                                    }
                                                }
                                                fmt.Println(v16, fi29)
                                            }
                                        }
                                        var v49 int8 = 12
                                        switch (fi20 >> ((18446744073709551615) & 63)) {
                                        case 32768:
                                            if false {
                                                v50 := GenS2{F0: ("foo bar" + ("" + "hello")), F1: v49, F2: 42}
                                                v49 = v49
                                                v10 = v10
                                                _ = v50
                                            } else {
                                                fmt.Println(v49, v38, v16)
                                            }
                                        case -350:
                                            var v51 float64 = 0.0
                                            _ = v51
                                        }
                                        _ = v38
                                        _ = v49
                                    } else {
                                        v16 = 737
                                    }
                                }
                            }
                            var v52 int64 = 42
                            v16 -= (-(v16))
                            _ = fi29
                            _ = v52
                        }
                    case "hello":
                        v10 = (^(v10))
                    }
                }
                switch "" {
                case "a":
                    v10 += (v10 - v10)
                case "hello":
                    v16 -= v16
                case "foo bar":
                    var v53 int8 = 42
                    _ = v53
                default:
                    fmt.Println(v16, fi20)
                }
                for fi54 := 0; fi54 < 18; fi54++ {
                    switch fi54 {
                    case 631:
                        var v55 uint8 = (^(v10))
                        _ = v55
                    case 127:
                        v16 = (v16 >> ((232) & 63))
                    }
                    switch ((("" + "hello") + "\n\t") + ("hello" + "\n\t")) {
                    case "":
                        var v56 bool = true
                        _ = v56
                    case "\n\t":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    v10 -= (^(v10))
                    _ = fi54
                }
                _ = fi20
            }
            if true {
                for fi57 := 0; fi57 < 1; fi57++ {
                    v58 := GenS6{F0: 42}
                    _ = fi57
                    _ = v58
                }
                var v59 uint = 42
                _ = v59
            } else {
                for fi60 := 0; fi60 < 4; fi60++ {
                    if ((42 != -128) && (-2147483648 > -328)) {
                        var v61 float64 = (3.14159 - 1.5)
                        v10 -= 0
                        v10 += 252
                        _ = v61
                    } else {
                        v10 = v10
                        fmt.Println(fi60, v16, fi60)
                    }
                    _ = fi60
                }
                switch 4294967295 {
                case 0:
                    if (true || ((65535 != 0) && (663 == 4294967295))) {
                        var v62 int8 = int8(v16)
                        for fi63 := 0; fi63 < 6; fi63++ {
                            var v64 uint16 = 796
                            var v65 int32 = -1
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                        for fi66 := 0; fi66 < 9; fi66++ {
                            if false {
                                v10 = v10
                                var v67 uint8 = (v10 + (v10 & v10))
                                _ = v67
                            }
                            _ = fi66
                        }
                        _ = v62
                    }
                case -9223372036854775808:
                    for fi68 := 0; fi68 < 2; fi68++ {
                        for fi69 := 0; fi69 < 3; fi69++ {
                            v5[849] = ("\n\t" + "a")
                            if true {
                                v10 -= v10
                                fmt.Println(fi69, v16, v10)
                            } else {
                                var v70 string = ("a" + "hello")
                                v71 := GenS6{F0: 42}
                                var v72 float32 = (float32(-2.25) / (-2.25 - (1.5 - 0.0)))
                                _ = v70
                                _ = v71
                                _ = v72
                            }
                            _ = fi69
                        }
                        _ = fi68
                    }
                case -181:
                    v73 := GenS7{F0: -128, F1: "a", F2: 42}
                    _ = v73
                default:
                    var v74 int = 127
                    _ = v74
                }
                v16 -= (v16 % (((v16 | 1)) | 1))
            }
            for fi75 := 0; fi75 < 1; fi75++ {
                if (((true || (42 >= 42)) && true) && false) {
                    switch "" {
                    case "hello":
                        for fi76 := 0; fi76 < 3; fi76++ {
                            var v77 uint = 758
                            switch 0 {
                            case 9223372036854775807:
                                fmt.Println(v77)
                                v78 := GenS6{F0: -241}
                                v79 := v4[fi75]
                                _ = v78
                                _ = v79
                            case -9223372036854775808:
                                var v80 uint32 = 2147483647
                                _ = v80
                            case -1:
                                v81 := GenS8{F0: v10, F1: 18446744073709551615}
                                v82 := GenS9{F0: (("foo bar" + ("\n\t" + "a")) + "hello")}
                                _ = v81
                                _ = v82
                            }
                            var v83 int32 = -2147483648
                            _ = fi76
                            _ = v77
                            _ = v83
                        }
                    case "":
                        for fi84 := 0; fi84 < 3; fi84++ {
                            _ = len(v3)
                            _ = fi84
                        }
                    default:
                        switch fi75 {
                        case -9223372036854775808:
                            for fi85 := 0; fi85 < 18; fi85++ {
                                v16 -= v16
                                v3 = append(v3, "\n\t")
                                _ = fi85
                            }
                        default:
                            if (fi75 < fi75) {
                                var v86 int64 = 1
                                _ = v86
                            }
                        }
                    }
                    switch fi75 {
                    case -852:
                        v16 += -1
                    case 127:
                        v87 := GenS9{F0: "a"}
                        _ = v87
                    case 441:
                        if (1 >= (^(v10))) {
                            fmt.Println(v10, v10, fi75)
                            var v88 uint8 = (v10 / (((v10 % (((v10 ^ (v10 + v10))) | 1))) | 1))
                            _ = v88
                        }
                    }
                } else {
                    v10 -= (v10 & v10)
                }
                v16 = -1
                var v89 int32 = 2147483647
                _ = fi75
                _ = v89
            }
        } else {
            v10 = 127
            v90 := GenS9{F0: "foo bar"}
            _ = v90
        }
        v16 -= (v16 - 32767)
    }
    for fi91 := 0; fi91 < 9; fi91++ {
        var v92 uint64 = 502
        switch (fi91 ^ fi91) {
        case -1:
            var v93 int64 = 9223372036854775807
            _ = v93
        case 2147483647:
            v16 = v16
        }
        _ = fi91
        _ = v92
    }
    fmt.Println(v10, v16)
    v16 = 42
    var v94 uint16 = 32767
    switch (("hello" + "\n\t") + ("a" + "foo bar")) {
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    default:
        if true {
            if (true || ((0.1 / (float32(-2.25) - (float32(float64(3.14159)) / 1e10))) > 3.14159)) {
                if false {
                    if true {
                        if (v10 < (v10 - (v10 >> ((783) & 63)))) {
                            for fi95 := 0; fi95 < 5; fi95++ {
                                v10 += v10
                                var v96 int32 = 1
                                _ = fi95
                                _ = v96
                            }
                            for fi97 := 0; fi97 < 10; fi97++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v3 = append(v3, (("foo bar" + (("" + ("" + "")) + ("" + "foo bar"))) + (("" + "hello") + "a")))
                                _ = fi97
                            }
                        }
                        var v98 int16 = v16
                        _ = v98
                    } else {
                        if (v10 >= 128) {
                            v94 += uint16(v10)
                            fmt.Println(v94, v16)
                        }
                    }
                    v10 -= (^(v10))
                    var v99 uint64 = 42
                    _ = v99
                }
            } else {
                var v100 bool = false
                if v100 {
                    if v100 {
                        if false {
                            v101 := GenS6{F0: 2147483647}
                            v102 := GenS6{F0: 1}
                            _ = v101
                            _ = v102
                        }
                        var v103 float64 = 0.0
                        var v104 int = 47
                        _ = v103
                        _ = v104
                    }
                    for fi105 := 0; fi105 < 19; fi105++ {
                        var v106 uint16 = v94
                        for fi107 := 0; fi107 < 17; fi107++ {
                            fmt.Println(v16, v10, v10)
                            var v108 uint32 = 42
                            var v109 string = "foo bar"
                            _ = fi107
                            _ = v108
                            _ = v109
                        }
                        _ = fi105
                        _ = v106
                    }
                }
                switch ("hello" + ((("a" + ("" + ("foo bar" + ""))) + "") + ("\n\t" + ""))) {
                case "foo bar":
                    v110 := GenS8{F0: (v10 - (v10 - v10)), F1: 27}
                    _ = v110
                default:
                    var v111 uint64 = 0
                    _ = v111
                }
                _ = v100
            }
            for fi112 := 0; fi112 < 15; fi112++ {
                if (("foo bar" + ("foo bar" + ("" + "\n\t"))) == ("\n\t" + (("a" + ("hello" + ("a" + ""))) + "foo bar"))) {
                    var v113 int16 = v16
                    if (2 != 1) {
                        if len(v3) > 0 {
                            v3[0] = ("a" + "hello")
                        }
                    } else {
                        if false {
                            for fi114 := 0; fi114 < 14; fi114++ {
                                v115 := GenS10{F0: 1, F1: true}
                                _ = fi114
                                _ = v115
                            }
                            fmt.Println(v16, fi112, fi112)
                        } else {
                            if ((v10 / (((v10 << ((1) & 63))) | 1)) > (^(v10))) {
                                if len(v3) > 0 {
                                    v3[0] = "a"
                                }
                                var v116 float64 = 0.1
                                _ = v116
                            } else {
                                if true {
                                    for fi117 := 0; fi117 < 11; fi117++ {
                                        var v118 uint16 = uint16(v10)
                                        v16 = (v113 % ((v16) | 1))
                                        _ = fi117
                                        _ = v118
                                    }
                                } else {
                                    _ = len(v3)
                                    fmt.Println(v16)
                                }
                                switch "foo bar" {
                                case "":
                                    if ("hello" != "hello") {
                                        var v119 int64 = -9223372036854775808
                                        var v120 uint64 = 0
                                        _ = v119
                                        _ = v120
                                    } else {
                                        if true {
                                            fmt.Println(v113, fi112, v16)
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v3 = append(v3, ("hello" + "\n\t"))
                                        }
                                    }
                                case "a":
                                    var v121 int = (fi112 - (fi112 ^ (fi112 / ((2147483647) | 1))))
                                    _ = v121
                                case "foo bar":
                                    v113 = (v16 % (((v16 >> ((747) & 63))) | 1))
                                default:
                                    for fi122 := 0; fi122 < 20; fi122++ {
                                        var v123 uint32 = 345
                                        _ = fi122
                                        _ = v123
                                    }
                                }
                                var v124 int = 9223372036854775807
                                _ = v124
                            }
                            var v125 int16 = -1
                            _ = v125
                        }
                        var v126 uint16 = v94
                        v10 = v10
                        _ = v126
                    }
                    for fi127 := 0; fi127 < 3; fi127++ {
                        v94 = v94
                        _ = fi127
                    }
                    _ = v113
                }
                v94 = 0
                _ = fi112
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v94:", v94)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v94
}
