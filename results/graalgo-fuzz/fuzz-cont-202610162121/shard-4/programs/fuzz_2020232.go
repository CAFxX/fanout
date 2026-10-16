package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS4 struct { F0 int32; F1 int32; F2 uint32 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int8; F1 int }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int16; F1 int64; F2 int64 }
type GenS12 struct { F0 int; F1 int8; F2 string }
func genFn0() int {
    v6 := GenS1{F0: 255}
    if (1e10 > (3.14159 / 1e10)) {
        if false {
            for fi7 := 0; fi7 < 9; fi7++ {
                var v8 uint64 = 4294967295
                _ = fi7
                _ = v8
            }
            v9 := []int{}
            _ = v9
        }
        var v10 uint64 = 42
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 int = (^(fi11))
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        if true {
            v13 := GenS2{F0: 396, F1: 37, F2: 255}
            _ = v13
        } else {
            for fi14 := 0; fi14 < 15; fi14++ {
                if (false || false) {
                    var v15 int16 = 62
                    fmt.Println(v15, v15)
                    _ = v15
                } else {
                    v16 := GenS3{F0: 42, F1: 2147483647, F2: -1}
                    for fi17 := 0; fi17 < 9; fi17++ {
                        var v18 float32 = float32(((0.1 / 0.1) * 0.0))
                        switch "a" {
                        case "hello":
                            v19 := GenS4{F0: 590, F1: 1, F2: 4294967295}
                            _ = v19
                        case "a":
                            v18 += v18
                        default:
                            fmt.Println(fi17, fi14, fi14)
                        }
                        var v20 uint64 = 42
                        _ = fi17
                        _ = v18
                        _ = v20
                    }
                    _ = v16
                }
                var v21 uint8 = 127
                v22 := GenS2{F0: 42, F1: 322, F2: (v21 * (^(v21)))}
                _ = fi14
                _ = v21
                _ = v22
            }
            var v23 int8 = 100
            for fi24 := 0; fi24 < 17; fi24++ {
                var v25 float64 = 0.0
                _ = fi24
                _ = v25
            }
            _ = v23
        }
    }
    switch "" {
    case "":
        var v26 int64 = -9223372036854775808
        _ = v26
    case "foo bar":
        if true {
            v27 := []int{9223372036854775807, -1}
            _ = v27
        } else {
            var v28 string = ("a" + "\n\t")
            _ = v28
        }
    case "hello":
        switch (("foo bar" + "foo bar") + "\n\t") {
        case "a":
            var v29 int16 = 153
            _ = v29
        case "foo bar":
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                var v30 bool = true
                _ = v30
            case "":
                var v31 int = 9223372036854775807
                _ = v31
            }
        default:
            var v32 float64 = (1e10 / 1.0)
            _ = v32
        }
    default:
        var v33 uint64 = 748
        _ = v33
    }
    var v34 int64 = 1
    _ = v34
    _ = v6
    return 2147483647
}
func genFn1() int32 {
    var v35 uint = 18446744073709551615
    var v36 string = (("" + (("\n\t" + "foo bar") + "hello")) + "")
    v36 = v36
    v36 = v36
    switch 32767 {
    case -256:
        fmt.Println(v36)
    case -929:
        v37 := GenS2{F0: 32767, F1: 2147483647, F2: 127}
        _ = v37
    case 127:
        for fi38 := 0; fi38 < 11; fi38++ {
            var v39 int32 = -2147483648
            _ = fi38
            _ = v39
        }
    }
    _ = v35
    _ = v36
    return int32(v35)
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{"a", ("" + "\n\t"), "", ("hello" + "\n\t")}
    var v3 int64 = v1
    var v4 uint32 = 65535
    var v5 uint64 = 99
    v5 += (v5 ^ v5)
    fmt.Println(v4)
    v5 = v5
    var v40 int8 = 1
    var v41 uint16 = 42
    if ((((1e10 * 1.5) * 3.14159) * 0.1) == float64(0.0)) {
        switch genFn0() {
        case -2147483648:
            for fi42 := 0; fi42 < 7; fi42++ {
                var v43 int32 = -1
                _ = fi42
                _ = v43
            }
        case -65:
            if (4294967295 != int(v5)) {
                v44 := GenS5{F0: 32768, F1: int64(v4)}
                v3 += v1
                _ = v44
            } else {
                if true {
                    v45 := GenS5{F0: 184, F1: v3}
                    fmt.Println(v4)
                    v2 = append(v2, "hello")
                    _ = v45
                }
                v40 = (v40 - v40)
            }
        case 127:
            for fi46 := 0; fi46 < 10; fi46++ {
                v47 := GenS1{F0: 128}
                _ = len(v2)
                _ = fi46
                _ = v47
            }
        }
        var v48 float32 = -2.25
        _ = v48
    } else {
        if len(v2) > 0 {
            v2[0] = ("a" + "foo bar")
        }
        v4 += (v4 ^ v4)
    }
    v49 := GenS6{F0: (v41 << ((0) & 63))}
    v50 := GenS7{F0: v4, F1: true, F2: (v41 << ((1) & 63))}
    var v51 uint32 = (^(v4))
    for fi52 := 0; fi52 < 7; fi52++ {
        switch ((("\n\t" + "\n\t") + ("foo bar" + "")) + "foo bar") {
        case "\n\t":
            if (v3 >= (v1 * (v3 % ((v1) | 1)))) {
                if ((163 > 2) && false) {
                    switch (fi52 | (fi52 + (fi52 & (fi52 % ((fi52) | 1))))) {
                    case -119:
                        for fi53 := 0; fi53 < 14; fi53++ {
                            switch "foo bar" {
                            case "a":
                                fmt.Println(v3)
                            default:
                                v1 = (v1 & v3)
                            }
                            _ = fi53
                        }
                    case -2147483648:
                        var v54 int16 = 1000
                        _ = v54
                    case 953:
                        for fi55 := 0; fi55 < 2; fi55++ {
                            v3 -= 342
                            var v56 uint64 = v5
                            v3 += v3
                            _ = fi55
                            _ = v56
                        }
                    default:
                        if ((fi52 * int(v5)) == int(v51)) {
                            var v57 float64 = (-2.25 + -2.25)
                            _ = v57
                        }
                    }
                    var v58 uint8 = 128
                    _ = v58
                }
                v40 = v40
                v40 -= v40
            } else {
                if true {
                    var v59 uint8 = 128
                    _ = v59
                }
                v2 = append(v2, "a")
                if (false && false) {
                    switch (fi52 / (((fi52 + (fi52 & int(v41)))) | 1)) {
                    case 240:
                        _ = len(v2)
                    }
                    if false {
                        if false {
                            for fi60 := 0; fi60 < 9; fi60++ {
                                v5 = v5
                                _ = fi60
                            }
                        }
                    }
                    var v61 int8 = v40
                    _ = v61
                } else {
                    for fi62 := 0; fi62 < 20; fi62++ {
                        if (32768 > (^(fi62))) {
                            v1 += (v3 & int64(fi62))
                            v5 -= v5
                            v4 -= v4
                        }
                        var v63 bool = (1 != 65535)
                        v64 := GenS5{F0: fi52, F1: (v1 ^ -9223372036854775808)}
                        var v65 int32 = -1
                        _ = fi62
                        _ = v63
                        _ = v64
                        _ = v65
                    }
                    var v66 uint64 = 447
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v66
                }
            }
        case "":
            if true {
                _ = len(v2)
                _ = len(v2)
            } else {
                for fi67 := 0; fi67 < 15; fi67++ {
                    v68 := GenS2{F0: v41, F1: v51, F2: 127}
                    if false {
                        switch "" {
                        case "a":
                            v69 := GenS8{F0: int32(v4)}
                            _ = v69
                        case "hello":
                            switch ("\n\t" + "hello") {
                            case "\n\t":
                                var v70 uint32 = v4
                                _ = v70
                            }
                            v71 := GenS5{F0: genFn0(), F1: (v1 << ((2) & 63))}
                            if true {
                                v3 += v1
                                v2 = append(v2, "a")
                                v72 := GenS9{F0: v40, F1: fi67}
                                _ = v72
                            } else {
                                if false {
                                    var v73 float32 = -2.25
                                    v41 = (^(v41))
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    _ = v73
                                } else {
                                    var v74 uint8 = 31
                                    var v75 int32 = 0
                                    if false {
                                        switch "foo bar" {
                                        case "hello":
                                            var v76 float32 = -2.25
                                            _ = v76
                                        case "a":
                                            var v77 int8 = (v40 & 90)
                                            _ = v77
                                        case "\n\t":
                                            var v78 int64 = v3
                                            _ = v78
                                        default:
                                            v74 -= (v74 * (v74 & (v74 >> ((1) & 63))))
                                        }
                                        if false {
                                            var v79 int64 = (v1 & (v1 * v1))
                                            _ = v79
                                        }
                                    }
                                    _ = v74
                                    _ = v75
                                }
                                var v80 float32 = 0.1
                                _ = v80
                            }
                            _ = v71
                        }
                    } else {
                        v1 += v1
                        if (true && true) {
                            var v81 float64 = (-2.25 - 0.1)
                            v41 += 65535
                            _ = v81
                        }
                    }
                    _ = fi67
                    _ = v68
                }
                if ((v40 << ((878) & 63)) < v40) {
                    switch 32767 {
                    case 127:
                        var v82 uint16 = v41
                        _ = v82
                    case 42:
                        v2 = append(v2, ("" + "a"))
                    case 32768:
                        fmt.Println(v51, v40, v51)
                    }
                    fmt.Println(v3, v40)
                } else {
                    var v83 uint8 = 91
                    switch ("" + ("\n\t" + "a")) {
                    case "":
                        switch fi52 {
                        case 32768:
                            var v84 float64 = 3.14159
                            _ = v84
                        case 2147483647:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case 66:
                            v40 = v40
                        default:
                            if (false || false) {
                                for fi85 := 0; fi85 < 19; fi85++ {
                                    if (-32768 <= -8) {
                                        var v86 uint64 = (v5 * 0)
                                        _ = v86
                                    }
                                    _ = fi85
                                }
                            }
                        }
                    case "\n\t":
                        var v87 uint16 = (v41 * (v41 * uint16(v5)))
                        _ = v87
                    default:
                        v88 := GenS4{F0: 420, F1: -2147483648, F2: (v4 + v51)}
                        _ = v88
                    }
                    for fi89 := 0; fi89 < 19; fi89++ {
                        v3 -= v1
                        if (false && ((!true) && true)) {
                            v4 = v4
                            if len(v2) > 0 {
                                v2[0] = ("hello" + "a")
                            }
                        }
                        _ = fi89
                    }
                    _ = v83
                }
                _ = len(v2)
            }
        case "foo bar":
            if true {
                var v90 uint = 42
                fmt.Println(v3, v1, v51)
                switch (("\n\t" + "") + ("a" + "a")) {
                case "":
                    if false {
                        var v91 uint = 2
                        var v92 uint = (v90 * 438)
                        _ = v91
                        _ = v92
                    }
                case "a":
                    var v93 int16 = 256
                    _ = v93
                case "hello":
                    switch (("foo bar" + "foo bar") + ("foo bar" + "foo bar")) {
                    case "":
                        _ = len(v2)
                    default:
                        for fi94 := 0; fi94 < 9; fi94++ {
                            fmt.Println(fi52, v5)
                            _ = fi94
                        }
                    }
                default:
                    v95 := GenS10{F0: v41}
                    _ = v95
                }
                _ = v90
            }
        default:
            if true {
                var v96 uint16 = (^(v41))
                fmt.Println(v5, v41, v3)
                _ = v96
            } else {
                if true {
                    v1 -= (v3 * (v1 - v1))
                    fmt.Println(v40, v51, v5)
                    v97 := GenS9{F0: (v40 + v40), F1: fi52}
                    _ = v97
                } else {
                    switch "hello" {
                    case "":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "foo bar":
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                    case "a":
                        for fi98 := 0; fi98 < 5; fi98++ {
                            v41 = (v41 >> ((336) & 63))
                            var v99 uint64 = v5
                            _ = fi98
                            _ = v99
                        }
                    default:
                        switch fi52 {
                        case 42:
                            if ((1000 != -877) && false) {
                                var v100 uint16 = (v41 / ((v41) | 1))
                                if true {
                                    var v101 float32 = (1.5 / float32(0.1))
                                    v5 -= 18446744073709551615
                                    if (v101 >= 0.1) {
                                        v41 = v41
                                        for fi102 := 0; fi102 < 19; fi102++ {
                                            v100 = v100
                                            v101 = (v101 * float32((((1.5 + 0.0) / 1.0) / 0.1)))
                                            switch fi102 {
                                            case 42:
                                                v103 := GenS11{F0: -907, F1: 0, F2: v3}
                                                _ = v103
                                            case 127:
                                                var v104 bool = false
                                                _ = v104
                                            }
                                            fmt.Println(v3, v5, v4)
                                            v105 := GenS11{F0: 997, F1: 578, F2: (v3 * v3)}
                                            _ = fi102
                                            _ = v105
                                        }
                                    } else {
                                        fmt.Println(v4)
                                    }
                                    _ = v101
                                }
                                for fi106 := 0; fi106 < 1; fi106++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v107 float32 = 3.14159
                                    fmt.Println(v100, v5)
                                    _ = fi106
                                    _ = v107
                                }
                                _ = v100
                            }
                        case 32768:
                            if false {
                                if true {
                                    switch "hello" {
                                    case "":
                                        v3 = (v3 - v1)
                                    case "a":
                                        var v108 string = (("" + (("\n\t" + "\n\t") + "foo bar")) + "foo bar")
                                        _ = v108
                                    case "\n\t":
                                        for fi109 := 0; fi109 < 10; fi109++ {
                                            v1 = v1
                                            var v110 int32 = genFn1()
                                            v41 += 1
                                            _ = fi109
                                            _ = v110
                                        }
                                    }
                                } else {
                                    v111 := GenS10{F0: v41}
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v111
                                }
                            } else {
                                fmt.Println(v3, v41)
                                v41 += 32767
                                _ = len(v2)
                            }
                        case 127:
                            for fi112 := 0; fi112 < 2; fi112++ {
                                if (v4 >= v4) {
                                    fmt.Println(v41, fi112)
                                    switch ("" + "hello") {
                                    case "a":
                                        if len(v2) > 0 {
                                            v2[0] = ""
                                        }
                                    case "\n\t":
                                        v113 := GenS8{F0: genFn1()}
                                        _ = v113
                                    case "hello":
                                        var v114 int8 = -81
                                        _ = v114
                                    }
                                } else {
                                    var v115 float64 = (((3.14159 + float64(((3.14159 * (-2.25 - (float32(-2.25) / float32(-2.25)))) + 3.14159))) - 0.0) + 0.1)
                                    _ = v115
                                }
                                if ((42 >= (v5 / ((v5) | 1))) || false) {
                                    v116 := GenS12{F0: (fi112 * -224), F1: v40, F2: "a"}
                                    v117 := GenS11{F0: 1, F1: 9223372036854775807, F2: v1}
                                    _ = v116
                                    _ = v117
                                } else {
                                    v5 = v5
                                    v118 := GenS9{F0: 127, F1: 549}
                                    _ = v118
                                }
                                _ = fi112
                            }
                        }
                    }
                }
            }
        }
        var v119 uint64 = v5
        _ = fi52
        _ = v119
    }
    var v120 int32 = -2147483648
    var v121 uint32 = (v51 * 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v120
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
