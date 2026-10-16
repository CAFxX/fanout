package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int16; F2 int8 }
type GenS3 struct { F0 int32; F1 int8; F2 int64 }
type GenS4 struct { F0 string; F1 int32; F2 uint8 }
type GenS5 struct { F0 int16; F1 int8; F2 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int; F1 uint8; F2 int16 }
type GenS8 struct { F0 int8; F1 int64 }
type GenS9 struct { F0 uint32; F1 int16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int16; F1 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 bool = false
    var v4 bool = (32767 < 32767)
    var v5 uint64 = 0
    v6 := make(map[int]int)
    fmt.Println(v4)
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int64 = int64(fi7)
        v9 := GenS1{F0: -103}
        v3 = v4
        _ = fi7
        _ = v8
        _ = v9
    }
    switch "foo bar" {
    case "":
        for fi10 := 0; fi10 < 11; fi10++ {
            if ((((float64((0.0 / -2.25)) * -2.25) * float64(-2.25)) * 3.14159) > (1e10 / ((1e10 / 3.14159) + 1e10))) {
                var v11 string = "hello"
                var v12 uint64 = (^(v5))
                _ = v11
                _ = v12
            }
            if ((false || false) && true) {
                switch fi10 {
                case -2147483648:
                    var v13 uint16 = 799
                    _ = v13
                case 174:
                    for fi14 := 0; fi14 < 12; fi14++ {
                        var v15 string = "foo bar"
                        v3 = true
                        var v16 uint64 = v5
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                }
            } else {
                for fi17 := 0; fi17 < 8; fi17++ {
                    switch "hello" {
                    case "foo bar":
                        var v18 uint16 = 65535
                        _ = v18
                    case "":
                        switch "\n\t" {
                        case "a":
                            v5 = (v5 / ((18446744073709551615) | 1))
                        case "hello":
                            fmt.Println(v3)
                        case "foo bar":
                            if v3 {
                                v19 := []int{}
                                _ = v19
                            } else {
                                fmt.Println(v3)
                            }
                        }
                    case "a":
                        if (!v4) {
                            v20 := GenS1{F0: -2147483648}
                            var v21 int32 = 1
                            var v22 uint = 18446744073709551615
                            _ = v20
                            _ = v21
                            _ = v22
                        } else {
                            v23 := GenS2{F0: 2, F1: -1, F2: -1}
                            switch 1 {
                            case 1:
                                v24 := GenS2{F0: 32, F1: -950, F2: int8(v5)}
                                _ = v24
                            }
                            switch ("foo bar" + ("a" + (("hello" + "a") + "foo bar"))) {
                            case "a":
                                var v25 int64 = 1
                                _ = v25
                            default:
                                v4 = ((!v3) && (395 > 0))
                            }
                            fmt.Println(v4, v3)
                            _ = v23
                        }
                    }
                    fmt.Println(fi10)
                    if (871 != 1) {
                        fmt.Println(v4, v4)
                        fmt.Println(fi17, fi17)
                    }
                    _ = fi17
                }
                v4 = ((!v3) && v3)
                v5 += (v5 + (v5 ^ v5))
            }
            _ = fi10
        }
    case "foo bar":
        switch 127 {
        case 42:
            v5 -= v5
        case -1:
            v5 = v5
        default:
            for fi26 := 0; fi26 < 1; fi26++ {
                var v27 uint32 = 2147483648
                var v28 int = 2147483647
                _ = fi26
                _ = v27
                _ = v28
            }
        }
    case "a":
        v3 = (!(-623 <= 42))
    default:
        fmt.Println(v3)
    }
    switch ("hello" + "") {
    case "foo bar":
        if ((1 < 57) && (1 != 42)) {
            v29 := GenS1{F0: 2147483647}
            v30 := []int{}
            if (-32768 >= 651) {
                switch 128 {
                case 1:
                    switch 32767 {
                    case 169:
                        if v4 {
                            for fi31 := 0; fi31 < 8; fi31++ {
                                v1[1] = 2
                                v3 = (!v4)
                                _ = fi31
                            }
                            v3 = (-32768 <= 1000)
                        } else {
                            switch "" {
                            case "foo bar":
                                v4 = (!v3)
                            case "":
                                if (0 >= 42) {
                                    var v32 float32 = (1.5 / 1e10)
                                    if v4 {
                                        switch "\n\t" {
                                        case "":
                                            _ = len(v6)
                                        case "foo bar":
                                            delete(v2, ("foo bar" + ("foo bar" + ("" + ("hello" + "\n\t")))))
                                        case "\n\t":
                                            v33 := GenS1{F0: -2147483648}
                                            _ = v33
                                        default:
                                            switch 9223372036854775807 {
                                            case 2147483647:
                                                if false {
                                                    fmt.Println(v3, v32)
                                                } else {
                                                    if false {
                                                        if (343 > (v5 & v5)) {
                                                            for fi34 := 0; fi34 < 15; fi34++ {
                                                                var v35 float64 = 3.14159
                                                                v36 := GenS3{F0: 42, F1: -1, F2: 9223372036854775807}
                                                                if true {
                                                                    v37 := GenS4{F0: "hello", F1: 0, F2: 11}
                                                                    switch fi34 {
                                                                    case 1:
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            var v38 uint = 42
                                                                            _ = v38
                                                                        case "\n\t":
                                                                            var v39 bool = false
                                                                            _ = v39
                                                                        case "":
                                                                            v30 = append(v30, fi34)
                                                                        }
                                                                    case 4294967295:
                                                                        for fi40 := 0; fi40 < 1; fi40++ {
                                                                            v41 := GenS4{F0: "foo bar", F1: 2147483647, F2: 70}
                                                                            v42 := GenS2{F0: 108, F1: -32768, F2: 42}
                                                                            _ = fi40
                                                                            _ = v41
                                                                            _ = v42
                                                                        }
                                                                    case 0:
                                                                        fmt.Println(v3, v32)
                                                                    default:
                                                                        var v43 float32 = (1.5 - v32)
                                                                        _ = v43
                                                                    }
                                                                    _ = v37
                                                                } else {
                                                                    v5 = v5
                                                                    fmt.Println(v5, v4)
                                                                }
                                                                _ = fi34
                                                                _ = v35
                                                                _ = v36
                                                            }
                                                            var v44 bool = false
                                                            _ = v44
                                                        } else {
                                                            fmt.Println(v3)
                                                        }
                                                        if true {
                                                            v45 := GenS3{F0: -1, F1: -1, F2: -9223372036854775808}
                                                            _ = v45
                                                        }
                                                        fmt.Println(v32, v3, v3)
                                                    }
                                                    v32 += 1e10
                                                }
                                            default:
                                                for fi46 := 0; fi46 < 19; fi46++ {
                                                    v4 = (255 > 2)
                                                    _ = fi46
                                                }
                                            }
                                        }
                                        fmt.Println(v32)
                                    }
                                    for fi47 := 0; fi47 < 17; fi47++ {
                                        var v48 float32 = v32
                                        v48 = v48
                                        v48 = v32
                                        _ = fi47
                                        _ = v48
                                    }
                                    _ = v32
                                }
                            case "a":
                                if v4 {
                                    for fi49 := 0; fi49 < 10; fi49++ {
                                        fmt.Println(v4, fi49)
                                        _ = fi49
                                    }
                                    if (!((1.5 - -2.25) == 1.5)) {
                                        _ = len(v2)
                                        v3 = (!(255 != 42))
                                        if v4 {
                                            var v50 float64 = -2.25
                                            v51 := GenS5{F0: 42, F1: -128, F2: false}
                                            _ = v50
                                            _ = v51
                                        } else {
                                            v3 = v3
                                        }
                                    }
                                    for fi52 := 0; fi52 < 3; fi52++ {
                                        v53 := GenS6{F0: 42}
                                        var v54 uint8 = 127
                                        _ = fi52
                                        _ = v53
                                        _ = v54
                                    }
                                } else {
                                    var v55 float64 = (0.0 - -2.25)
                                    var v56 uint8 = 42
                                    _ = v55
                                    _ = v56
                                }
                            default:
                                for fi57 := 0; fi57 < 3; fi57++ {
                                    switch (fi57 % (((fi57 % (((-(fi57))) | 1))) | 1)) {
                                    case 1:
                                        fmt.Println(v3, v4, v5)
                                    case -1:
                                        v5 = v5
                                    default:
                                        fmt.Println(v3, v3)
                                    }
                                    var v58 int8 = -57
                                    v59 := GenS7{F0: 32768, F1: 0, F2: 1000}
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                }
                            }
                            v3 = (v5 < (v5 | v5))
                            switch 32768 {
                            case 2147483647:
                                if v3 {
                                    v60 := GenS8{F0: -24, F1: 0}
                                    _ = v60
                                } else {
                                    v61 := GenS6{F0: 42}
                                    var v62 string = ""
                                    _ = v61
                                    _ = v62
                                }
                            case 32768:
                                v63 := v6[128]
                                _ = v63
                            default:
                                switch 413 {
                                case 128:
                                    fmt.Println(v5, v3)
                                case 1:
                                    _ = len(v30)
                                }
                            }
                        }
                    case 42:
                        switch "\n\t" {
                        case "":
                            var v64 uint = 153
                            _ = v64
                        case "foo bar":
                            if len(v30) >= 2 {
                                v30 = v30[1:len(v30)]
                            }
                        }
                    case 32767:
                        var v65 int8 = 100
                        _ = v65
                    }
                }
                v66 := GenS6{F0: 59}
                for fi67 := 0; fi67 < 1; fi67++ {
                    var v68 string = (("\n\t" + ("\n\t" + "\n\t")) + ("a" + ""))
                    for fi69 := 0; fi69 < 3; fi69++ {
                        _ = len(v2)
                        _ = len(v1)
                        _ = fi69
                    }
                    _ = fi67
                    _ = v68
                }
                _ = v66
            } else {
                if v3 {
                    v4 = v3
                    var v70 float64 = 1e10
                    _ = v70
                } else {
                    fmt.Println(v5, v3, v3)
                    var v71 int16 = 42
                    _ = v71
                }
            }
            _ = v29
            _ = v30
        }
    }
    v5 = v5
    switch ("" + "foo bar") {
    case "hello":
        fmt.Println(v4)
    default:
        if (-9223372036854775808 >= -788) {
            _ = len(v1)
            v5 = 156
            if (!false) {
                for fi72 := 0; fi72 < 7; fi72++ {
                    v5 -= v5
                    if false {
                        v3 = v3
                        switch ("foo bar" + "") {
                        case "hello":
                            v73 := GenS3{F0: -761, F1: 62, F2: -781}
                            _ = v73
                        case "\n\t":
                            v3 = v3
                        default:
                            v3 = (9223372036854775807 < 42)
                        }
                        for fi74 := 0; fi74 < 12; fi74++ {
                            if v4 {
                                v3 = true
                                v75 := []int{fi74, (fi72 >> ((2) & 63)), fi74}
                                _ = v75
                            } else {
                                if ((255 > 127) || (!v4)) {
                                    if (820 < v5) {
                                        v76 := GenS4{F0: ("foo bar" + (("\n\t" + "hello") + "foo bar")), F1: 779, F2: 127}
                                        switch (fi72 & fi72) {
                                        case 42:
                                            v3 = (0.1 > -2.25)
                                        case -2147483648:
                                            switch "" {
                                            case "hello":
                                                var v77 int64 = 0
                                                if v3 {
                                                    if true {
                                                        v77 = v77
                                                        var v78 uint8 = 120
                                                        _ = v78
                                                    }
                                                }
                                                switch ((("\n\t" + "") + ("foo bar" + "a")) + "foo bar") {
                                                case "\n\t":
                                                    var v79 int16 = 129
                                                    _ = v79
                                                default:
                                                    delete(v2, "foo bar")
                                                }
                                                _ = v77
                                            case "a":
                                                var v80 uint64 = (v5 + 18446744073709551615)
                                                _ = v80
                                            }
                                        }
                                        var v81 int32 = -720
                                        _ = v76
                                        _ = v81
                                    } else {
                                        if false {
                                            v2["hello"] = v4
                                            var v82 int32 = 2147483647
                                            _ = v82
                                        }
                                        v83 := GenS2{F0: 24, F1: 1, F2: -61}
                                        if (-1 >= -103) {
                                            var v84 float32 = (-2.25 / 0.1)
                                            _ = v84
                                        } else {
                                            var v85 int8 = 127
                                            _ = v85
                                        }
                                        _ = v83
                                    }
                                    v86 := []int{}
                                    var v87 int16 = 32767
                                    _ = v86
                                    _ = v87
                                }
                                v88 := GenS4{F0: ("" + ("\n\t" + "foo bar")), F1: 715, F2: 255}
                                v89 := GenS9{F0: 141, F1: 0}
                                _ = v88
                                _ = v89
                            }
                            _ = fi74
                        }
                    }
                    _ = fi72
                }
                var v90 uint32 = 65535
                _ = v90
            } else {
                switch ("\n\t" + "hello") {
                case "\n\t":
                    if (0 > 688) {
                        if ((497 != 1) && false) {
                            var v91 int32 = -2147483648
                            var v92 string = (("foo bar" + ("" + "hello")) + "")
                            _ = v91
                            _ = v92
                        } else {
                            var v93 int8 = 42
                            var v94 string = ("\n\t" + "foo bar")
                            _ = v93
                            _ = v94
                        }
                        v3 = v4
                    } else {
                        for fi95 := 0; fi95 < 11; fi95++ {
                            for fi96 := 0; fi96 < 16; fi96++ {
                                var v97 int = fi96
                                switch -9223372036854775808 {
                                case 1:
                                    var v98 uint32 = 331
                                    _ = v98
                                case 9223372036854775807:
                                    v3 = v4
                                default:
                                    v5 = v5
                                }
                                v99 := []int{(fi95 % ((1) | 1)), fi95}
                                v4 = (42 > 42)
                                v100 := GenS1{F0: 2147483647}
                                _ = fi96
                                _ = v100
                                _ = v97
                                _ = v99
                            }
                            v3 = (!false)
                            _ = fi95
                        }
                    }
                case "a":
                    v101 := GenS4{F0: "\n\t", F1: 849, F2: 2}
                    _ = v101
                case "hello":
                    if (1e10 > 1.5) {
                        v3 = false
                    }
                }
            }
        } else {
            fmt.Println(v4)
        }
    }
    _ = len(v6)
    _ = len(v2)
    switch (("\n\t" + "a") + "\n\t") {
    case "a":
        if false {
            for fi102 := 0; fi102 < 10; fi102++ {
                var v103 uint32 = 905
                _ = fi102
                _ = v103
            }
            v4 = (1 == -32768)
        }
    case "":
        v1[628] = 2147483648
    case "\n\t":
        v3 = true
    }
    fmt.Println(v3, v4, v5)
    for fi104 := 0; fi104 < 5; fi104++ {
        for fi105 := 0; fi105 < 8; fi105++ {
            if ((v3 || true) && v3) {
                var v106 string = "\n\t"
                _ = v106
            }
            v107 := []int{(fi105 / ((fi104) | 1))}
            var v108 int = fi104
            var v109 float32 = ((float32(0.0) / float32(0.1)) * float32(3.14159))
            _ = fi105
            _ = v107
            _ = v108
            _ = v109
        }
        v3 = true
        _ = fi104
    }
    for fi110 := 0; fi110 < 20; fi110++ {
        v2[("\n\t" + "a")] = (v4 && false)
        for fi111 := 0; fi111 < 1; fi111++ {
            switch fi110 {
            case 32767:
                var v112 int = (fi110 >> ((2) & 63))
                _ = v112
            case 799:
                v113 := GenS10{F0: (fi110 << ((4294967295) & 63))}
                _ = len(v6)
                var v114 int = 2147483647
                v114 += (fi110 - -9223372036854775808)
                v115 := []int{}
                _ = v113
                _ = v114
                _ = v115
            }
            var v116 int = fi110
            _ = fi111
            _ = v116
        }
        for fi117 := 0; fi117 < 12; fi117++ {
            if v3 {
                var v118 uint32 = 157
                fmt.Println(v118)
                switch -407 {
                case 450:
                    var v119 uint64 = v5
                    _ = v119
                }
                v120 := GenS11{F0: 65535}
                var v121 int = (fi110 ^ (fi117 ^ 32767))
                v1[v118] = (v118 ^ v118)
                v5 = v5
                v122 := GenS12{F0: -442, F1: 255}
                _ = v118
                _ = v120
                _ = v121
                _ = v122
            } else {
                v3 = v4
                _ = len(v2)
            }
            _ = fi117
        }
        _ = fi110
    }
    var v123 uint = 817
    if (42 > 127) {
        switch "hello" {
        case "":
            v4 = true
        }
        switch 0 {
        case 127:
            fmt.Println(v123, v5)
        case -275:
            _ = len(v6)
        case 128:
            v123 = v123
        }
    } else {
        var v124 int32 = -198
        _ = v124
    }
    var v125 int64 = -9223372036854775808
    for fi126 := 0; fi126 < 19; fi126++ {
        var v127 int32 = -2147483648
        var v128 bool = (!((v123 - (v123 ^ (v123 >> ((v123) & 63)))) == v123))
        _ = fi126
        _ = v127
        _ = v128
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v123:", v123)
    fmt.Println("v125:", v125)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v123
    _ = v125
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
