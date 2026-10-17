package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 int16; F2 int16 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 int64; F1 int8; F2 int16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int8; F1 bool }
func genFn0(p0 int16, p1 uint, p2 uint16) uint64 {
    fmt.Println(p2)
    fmt.Println(p1)
    v5 := GenS1{F0: p1, F1: 42}
    fmt.Println(p2, p0)
    var v6 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int32 = 42
    var v2 float64 = 3.14159
    var v3 uint64 = 0
    v4 := make(map[string]string)
    v3 -= v3
    v7 := []int{}
    v7 = append(v7, 882)
    v8 := GenS1{F0: 65535, F1: genFn0(42, 47, 65535)}
    fmt.Println(v2, v1, v3)
    if ((("foo bar" + "hello") + "\n\t") >= "") {
        _ = len(v4)
        v7 = append(v7, 128)
        var v9 uint64 = v3
        _ = v9
    } else {
        v1 = v1
        switch 32767 {
        case 42:
            switch "\n\t" {
            case "":
                v10 := GenS1{F0: 42, F1: 666}
                _ = v10
            }
        case 687:
            v2 -= 3.14159
        }
    }
    v3 += v3
    v1 = 1
    if ((true || (!true)) || false) {
        switch ("foo bar" + "") {
        case "\n\t":
            v11 := GenS2{F0: 127}
            _ = v11
        case "hello":
            for fi12 := 0; fi12 < 4; fi12++ {
                if true {
                    v2 = v2
                    if len(v7) > 0 {
                        v7[0] = (fi12 & -9223372036854775808)
                    }
                    v1 = (v1 * v1)
                } else {
                    fmt.Println(v1, v3, v2)
                }
                if false {
                    switch fi12 {
                    case -590:
                        fmt.Println(v3, fi12, v2)
                    default:
                        var v13 bool = (v3 != 2)
                        _ = v13
                    }
                    var v14 uint16 = 1
                    _ = v14
                }
                if (209 <= 27) {
                    var v15 int16 = -181
                    if (-449 != fi12) {
                        if true {
                            var v16 int8 = 1
                            _ = v16
                        } else {
                            if false {
                                if (false && false) {
                                    var v17 uint8 = 35
                                    if true {
                                        fmt.Println(v17, fi12, fi12)
                                        switch "" {
                                        case "hello":
                                            for fi18 := 0; fi18 < 9; fi18++ {
                                                var v19 float32 = 0.0
                                                var v20 float32 = v19
                                                var v21 bool = true
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                                if len(v7) > 0 {
                                                    v7[0] = fi18
                                                }
                                                var v22 uint8 = (v17 * v17)
                                                switch fi12 {
                                                case -898:
                                                    if (v21 && (true && v21)) {
                                                        v23 := GenS2{F0: 42}
                                                        _ = v23
                                                    }
                                                case 42:
                                                    fmt.Println(v3, fi12)
                                                }
                                                _ = fi18
                                                _ = v19
                                                _ = v20
                                                _ = v21
                                                _ = v22
                                            }
                                        case "\n\t":
                                            v4[""] = "foo bar"
                                        default:
                                            var v24 uint16 = 42
                                            _ = v24
                                        }
                                    } else {
                                        var v25 int32 = v1
                                        fmt.Println(v1)
                                        var v26 float64 = 0.1
                                        _ = v25
                                        _ = v26
                                    }
                                    var v27 uint64 = v3
                                    _ = v17
                                    _ = v27
                                } else {
                                    fmt.Println(fi12, v3, v2)
                                    for fi28 := 0; fi28 < 8; fi28++ {
                                        switch (fi28 >> ((2) & 63)) {
                                        case 42:
                                            var v29 uint32 = 253
                                            _ = v29
                                        default:
                                            if (fi12 != (^(fi12))) {
                                                v3 += genFn0(1, 42, 42)
                                                if (int32(v15) >= (v1 & (-(v1)))) {
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                } else {
                                                    v3 += v3
                                                    v2 = (v2 + -2.25)
                                                }
                                            } else {
                                                var v30 uint64 = (v3 | (v3 ^ 42))
                                                var v31 uint8 = 127
                                                _ = v30
                                                _ = v31
                                            }
                                        }
                                        var v32 int16 = v15
                                        var v33 int32 = v1
                                        _ = fi28
                                        _ = v32
                                        _ = v33
                                    }
                                    if true {
                                        var v34 int = fi12
                                        for fi35 := 0; fi35 < 15; fi35++ {
                                            v36 := GenS1{F0: 4294967295, F1: genFn0(v15, 255, 2)}
                                            switch fi35 {
                                            case -303:
                                                switch 32768 {
                                                case -541:
                                                    if (!false) {
                                                        var v37 int16 = (v15 ^ (v15 * 0))
                                                        _ = v37
                                                    } else {
                                                        var v38 int32 = (v1 / ((-986) | 1))
                                                        fmt.Println(v34, v1, v3)
                                                        switch 127 {
                                                        case 0:
                                                            switch (fi12 * fi12) {
                                                            case 9223372036854775807:
                                                                var v39 string = "a"
                                                                _ = v39
                                                            default:
                                                                fmt.Println(v38, v1, fi12)
                                                            }
                                                        case -2147483648:
                                                            v40 := GenS1{F0: 235, F1: v3}
                                                            _ = v40
                                                        default:
                                                            v7 = append(v7, fi35)
                                                        }
                                                        _ = v38
                                                    }
                                                    if (!(127 < -128)) {
                                                        v3 -= 0
                                                    } else {
                                                        if false {
                                                            v41 := GenS1{F0: 65535, F1: (v3 - genFn0((-(v15)), 0, 0))}
                                                            _ = len(v7)
                                                            switch 919 {
                                                            case 2147483647:
                                                                var v42 uint8 = 128
                                                                _ = v42
                                                            }
                                                            _ = v41
                                                        }
                                                        switch (fi35 / ((int(v3)) | 1)) {
                                                        case 32767:
                                                            fmt.Println(v2, v1)
                                                            fmt.Println(fi35)
                                                            if ((v15 + (v15 / ((v15) | 1))) > v15) {
                                                                if (-2.25 < float32(v2)) {
                                                                    if len(v7) > 0 {
                                                                        v7[0] = fi35
                                                                    }
                                                                    if true {
                                                                        var v43 uint16 = 255
                                                                        _ = len(v7)
                                                                        var v44 float64 = float64(3.14159)
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                }
                                                                v45 := GenS2{F0: 35}
                                                                if false {
                                                                    _ = len(v4)
                                                                    v1 = v1
                                                                } else {
                                                                    var v46 uint64 = genFn0(v15, 255, 2)
                                                                    _ = v46
                                                                }
                                                                _ = v45
                                                            }
                                                        default:
                                                            var v47 uint64 = (v3 / ((1) | 1))
                                                            _ = v47
                                                        }
                                                    }
                                                case -9223372036854775808:
                                                    if false {
                                                        v15 = v15
                                                        var v48 uint64 = genFn0(v15, 18446744073709551615, 736)
                                                        v3 = 4294967295
                                                        _ = v48
                                                    }
                                                }
                                            case 4294967295:
                                                var v49 uint16 = 500
                                                if ((!true) || false) {
                                                    v34 = (fi35 % ((fi12) | 1))
                                                } else {
                                                    var v50 int = (v34 & 2147483647)
                                                    v51 := v4["foo bar"]
                                                    if ((-2.25 >= -2.25) && (int(v15) <= (v34 + -993))) {
                                                        if false {
                                                            v52 := GenS1{F0: 1, F1: (v3 & (v3 + (^(v3))))}
                                                            v49 = v49
                                                            _ = v52
                                                        }
                                                        var v53 float64 = v2
                                                        if true {
                                                            if true {
                                                                v15 -= v15
                                                            } else {
                                                                v4[v51] = ("a" + v51)
                                                                switch 2147483647 {
                                                                case 42:
                                                                    var v54 uint32 = uint32(v49)
                                                                    v54 -= (v54 + v54)
                                                                    switch (v50 + (v34 * v50)) {
                                                                    case 345:
                                                                        if (true && (0 <= 0)) {
                                                                            if (genFn0((v15 * (v15 / (((^(v15))) | 1))), 255, 0) >= v3) {
                                                                                if (true || false) {
                                                                                    var v55 uint32 = v54
                                                                                    delete(v4, ((((((v51 + v51) + v51) + v51) + v51) + ((("hello" + ("hello" + v51)) + v51) + v51)) + v51))
                                                                                    v7 = append(v7, (v50 * int(v49)))
                                                                                    _ = v55
                                                                                } else {
                                                                                    switch "\n\t" {
                                                                                    case "":
                                                                                        var v56 int64 = 9223372036854775807
                                                                                        _ = v56
                                                                                    case "a":
                                                                                        if true {
                                                                                            v57 := GenS3{F0: 2147483648}
                                                                                            _ = v57
                                                                                        } else {
                                                                                            delete(v4, "\n\t")
                                                                                            var v58 uint8 = 34
                                                                                            if len(v7) >= 2 {
                                                                                                v7 = v7[1:len(v7)]
                                                                                            }
                                                                                            fmt.Println(v3, fi35)
                                                                                            _ = v58
                                                                                        }
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                _ = len(v4)
                                                                            }
                                                                            var v59 bool = (127 >= 1)
                                                                            _ = v59
                                                                        } else {
                                                                            var v60 uint64 = (v3 | genFn0(v15, 4294967295, v49))
                                                                            _ = len(v7)
                                                                            _ = v60
                                                                        }
                                                                    case 851:
                                                                        fmt.Println(fi12, v49, fi12)
                                                                    case -180:
                                                                        fmt.Println(v51, fi12, fi35)
                                                                    default:
                                                                        v7 = append(v7, (v50 % ((-9223372036854775808) | 1)))
                                                                    }
                                                                    _ = v54
                                                                case -2147483648:
                                                                    _ = len(v4)
                                                                default:
                                                                    fmt.Println(v15, v1)
                                                                }
                                                            }
                                                            var v61 bool = (false && true)
                                                            v4[("\n\t" + v51)] = v51
                                                            if false {
                                                                if false {
                                                                    switch v51 {
                                                                    case "\n\t":
                                                                        var v62 uint8 = 127
                                                                        v1 += v1
                                                                        var v63 uint8 = (v62 << ((965) & 63))
                                                                        v64 := GenS4{F0: 0, F1: v15, F2: 0}
                                                                        v3 -= (v3 & genFn0(1000, 2, (v49 * 42)))
                                                                        _ = v62
                                                                        _ = v63
                                                                        _ = v64
                                                                    case "a":
                                                                        _ = len(v7)
                                                                    case "hello":
                                                                        v49 = 784
                                                                    }
                                                                    fmt.Println(v3, v49, v61)
                                                                }
                                                            }
                                                            switch v50 {
                                                            case 32767:
                                                                var v65 string = (("hello" + v51) + v51)
                                                                _ = v65
                                                            case -2147483648:
                                                                fmt.Println(v61)
                                                            }
                                                            _ = v61
                                                        }
                                                        fmt.Println(fi12, fi35)
                                                        switch v34 {
                                                        case 128:
                                                            if false {
                                                                if len(v7) > 0 {
                                                                    v7[0] = -720
                                                                }
                                                                switch -59 {
                                                                case 554:
                                                                    if (2 < 0) {
                                                                        var v66 int16 = int16(v1)
                                                                        v1 -= (v1 - 0)
                                                                        v2 = (v2 / v2)
                                                                        _ = v66
                                                                    } else {
                                                                        var v67 int = v50
                                                                        v3 = (v3 << ((65535) & 63))
                                                                        if true {
                                                                            var v68 float32 = ((float32(((float64(0.0) - 0.0) * v2)) - (float32(((v2 / v53) * (v53 / v2))) / 1.5)) / 1.0)
                                                                            var v69 uint = 2
                                                                            _ = v68
                                                                            _ = v69
                                                                        } else {
                                                                            fmt.Println(v53, v51, v49)
                                                                            var v70 int32 = v1
                                                                            _ = v70
                                                                        }
                                                                        _ = v67
                                                                    }
                                                                case -184:
                                                                    v3 = (v3 - (v3 / (((v3 << ((1) & 63))) | 1)))
                                                                }
                                                                v71 := v4[v51]
                                                                _ = v71
                                                            }
                                                        default:
                                                            v72 := GenS1{F0: 4294967295, F1: (v3 & (v3 << ((179) & 63)))}
                                                            _ = v72
                                                        }
                                                        v73 := v4[v51]
                                                        _ = v53
                                                        _ = v73
                                                    } else {
                                                        _ = len(v7)
                                                        switch (v50 / ((fi35) | 1)) {
                                                        case 2147483647:
                                                            switch "hello" {
                                                            case "\n\t":
                                                                if false {
                                                                    v7 = append(v7, fi35)
                                                                } else {
                                                                    var v74 int8 = 0
                                                                    var v75 uint = 18446744073709551615
                                                                    var v76 int64 = 988
                                                                    _ = v74
                                                                    _ = v75
                                                                    _ = v76
                                                                }
                                                            case "foo bar":
                                                                var v77 uint32 = 923
                                                                _ = v77
                                                            }
                                                        case 876:
                                                            v50 += v34
                                                        }
                                                    }
                                                    _ = v50
                                                    _ = v51
                                                }
                                                _ = v49
                                            default:
                                                var v78 string = ""
                                                _ = v78
                                            }
                                            if false {
                                                v79 := GenS1{F0: 255, F1: 742}
                                                var v80 int64 = 0
                                                var v81 int16 = v15
                                                _ = v79
                                                _ = v80
                                                _ = v81
                                            } else {
                                                if len(v7) > 0 {
                                                    v7[0] = -224
                                                }
                                                v82 := GenS3{F0: 1}
                                                _ = v82
                                            }
                                            _ = fi35
                                            _ = v36
                                        }
                                        _ = v34
                                    } else {
                                        _ = len(v7)
                                    }
                                }
                                fmt.Println(fi12, fi12)
                            } else {
                                delete(v4, "")
                                if (42 >= -49) {
                                    v1 -= -843
                                    var v83 int32 = 2147483647
                                    _ = v83
                                }
                            }
                            var v84 string = ("a" + ("hello" + "foo bar"))
                            if false {
                                if (294 > 4294967295) {
                                    fmt.Println(v1, v1)
                                    if (!false) {
                                        var v85 string = "a"
                                        v15 = v15
                                        _ = v85
                                    }
                                } else {
                                    var v86 int32 = 1
                                    fmt.Println(v3, v3)
                                    _ = v86
                                }
                            }
                            _ = v84
                        }
                    } else {
                        for fi87 := 0; fi87 < 15; fi87++ {
                            fmt.Println(fi12)
                            v2 = (0.0 + v2)
                            _ = fi87
                        }
                        for fi88 := 0; fi88 < 18; fi88++ {
                            v2 += v2
                            _ = fi88
                        }
                    }
                    for fi89 := 0; fi89 < 7; fi89++ {
                        v1 -= 11
                        _ = fi89
                    }
                    _ = v15
                } else {
                    for fi90 := 0; fi90 < 20; fi90++ {
                        var v91 uint16 = 255
                        _ = fi90
                        _ = v91
                    }
                    switch fi12 {
                    case 2147483647:
                        v3 = (v3 ^ (^(v3)))
                    case 4294967295:
                        fmt.Println(v3, v2, v3)
                    case 32767:
                        if ((fi12 * fi12) > (fi12 + -302)) {
                            for fi92 := 0; fi92 < 6; fi92++ {
                                v93 := GenS2{F0: 128}
                                var v94 uint8 = 206
                                fmt.Println(v3, fi12, v2)
                                var v95 int32 = (^(v1))
                                _ = fi92
                                _ = v93
                                _ = v94
                                _ = v95
                            }
                            switch (fi12 << ((966) & 63)) {
                            case 32768:
                                v96 := GenS2{F0: 2}
                                _ = v96
                            }
                            if ((1.5 - (float32(v2) - 0.1)) <= (1e10 * -2.25)) {
                                var v97 uint32 = 42
                                var v98 uint64 = v3
                                _ = v97
                                _ = v98
                            } else {
                                for fi99 := 0; fi99 < 6; fi99++ {
                                    if (2147483648 < 2147483648) {
                                        fmt.Println(v3, v2)
                                    } else {
                                        if false {
                                            v100 := GenS3{F0: 900}
                                            _ = v100
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            fmt.Println(fi12, fi12)
                                        case "a":
                                            v1 += v1
                                        case "":
                                            if false {
                                                switch ("a" + "a") {
                                                case "":
                                                    if true {
                                                        var v101 float32 = 1.5
                                                        fmt.Println(v1, v1, v1)
                                                        _ = v101
                                                    } else {
                                                        if ((fi99 & fi99) < -33) {
                                                            v7 = append(v7, (^(fi99)))
                                                            if len(v7) >= 2 {
                                                                v7 = v7[1:len(v7)]
                                                            }
                                                            var v102 uint = 42
                                                            _ = v102
                                                        }
                                                        v103 := GenS4{F0: 2147483648, F1: int16(v3), F2: -32768}
                                                        _ = v103
                                                    }
                                                case "foo bar":
                                                    switch fi12 {
                                                    case 32767:
                                                        fmt.Println(v2)
                                                        v2 = ((((v2 / (1.5 - (float64(1e10) + v2))) * (((float64(1.5) * v2) - 0.0) / 1e10)) + 0.1) + v2)
                                                    }
                                                default:
                                                    fmt.Println(fi99, v2)
                                                }
                                                v104 := GenS5{F0: fi12, F1: 67}
                                                if ((!((-9223372036854775808 < 204) || false)) && true) {
                                                    v3 = v3
                                                    if (2 <= 893) {
                                                        var v105 float64 = ((v2 * -2.25) + -2.25)
                                                        var v106 int8 = 1
                                                        if ((true || ("foo bar" < "")) && true) {
                                                            var v107 int64 = 0
                                                            var v108 uint32 = 2
                                                            _ = v107
                                                            _ = v108
                                                        } else {
                                                            switch -536 {
                                                            case 128:
                                                                if (!false) {
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        v1 -= int32(v106)
                                                                    case "hello":
                                                                        if len(v7) > 0 {
                                                                            v7[0] = (fi99 >> ((741) & 63))
                                                                        }
                                                                    case "foo bar":
                                                                        var v109 int64 = -9223372036854775808
                                                                        _ = v109
                                                                    default:
                                                                        if ((("" + "") != "foo bar") || (!false)) {
                                                                            if false {
                                                                                _ = len(v7)
                                                                                _ = len(v7)
                                                                                fmt.Println(v106, fi12)
                                                                            }
                                                                        }
                                                                    }
                                                                    fmt.Println(v105, v106)
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        v110 := GenS6{F0: 520, F1: ("a" + "foo bar")}
                                                                        _ = v110
                                                                    default:
                                                                        switch (-(fi99)) {
                                                                        case 1000:
                                                                            var v111 uint = 4294967295
                                                                            _ = v111
                                                                        }
                                                                    }
                                                                } else {
                                                                    v112 := GenS7{F0: 9223372036854775807, F1: 85, F2: -1}
                                                                    v113 := GenS1{F0: 255, F1: 2}
                                                                    var v114 float32 = 0.0
                                                                    var v115 int16 = 0
                                                                    v116 := GenS5{F0: (fi99 >> ((4294967295) & 63)), F1: 73}
                                                                    v117 := GenS6{F0: 0, F1: "hello"}
                                                                    var v118 uint = 2
                                                                    var v119 float64 = (v105 + ((1.5 + 3.14159) + 0.0))
                                                                    fmt.Println(v119, v1)
                                                                    _ = v112
                                                                    _ = v113
                                                                    _ = v114
                                                                    _ = v115
                                                                    _ = v116
                                                                    _ = v117
                                                                    _ = v118
                                                                    _ = v119
                                                                }
                                                            }
                                                        }
                                                        _ = v105
                                                        _ = v106
                                                    }
                                                } else {
                                                    v7 = append(v7, fi12)
                                                    v1 = (^(v1))
                                                }
                                                fmt.Println(v1)
                                                switch -32 {
                                                case 2147483647:
                                                    v120 := v4["a"]
                                                    _ = v120
                                                default:
                                                    if ((v3 % (((^(v3))) | 1)) > v3) {
                                                        fmt.Println(v1)
                                                    }
                                                }
                                                _ = v104
                                            }
                                        default:
                                            v121 := GenS8{F0: 100}
                                            _ = v121
                                        }
                                    }
                                    v122 := GenS9{F0: -108, F1: (false || true)}
                                    _ = fi99
                                    _ = v122
                                }
                                if (!(v3 < genFn0(1, 255, 42))) {
                                    var v123 int = (fi12 | (fi12 - fi12))
                                    _ = v123
                                } else {
                                    switch "hello" {
                                    case "foo bar":
                                        v3 -= v3
                                    }
                                }
                            }
                        } else {
                            v124 := GenS2{F0: 128}
                            if ((0 > 2) || true) {
                                v1 -= v1
                            }
                            _ = v124
                        }
                    }
                }
                _ = fi12
            }
        }
        var v125 int32 = v1
        _ = v125
    } else {
        _ = len(v4)
        if (v2 < v2) {
            if false {
                for fi126 := 0; fi126 < 5; fi126++ {
                    for fi127 := 0; fi127 < 18; fi127++ {
                        fmt.Println(fi127, v2, v2)
                        switch ("foo bar" + ("hello" + "hello")) {
                        case "\n\t":
                            v2 -= 0.0
                        case "":
                            v2 = 1.5
                        default:
                            v2 += (((0.1 - 1.5) * v2) / 1.5)
                        }
                        _ = len(v7)
                        _ = fi127
                    }
                    if (-1 < 677) {
                        fmt.Println(fi126, v1, v1)
                        v7 = append(v7, fi126)
                        var v128 uint = 65535
                        _ = v128
                    }
                    _ = fi126
                }
                for fi129 := 0; fi129 < 11; fi129++ {
                    var v130 int = 9223372036854775807
                    v131 := GenS4{F0: 350, F1: 32767, F2: 152}
                    var v132 float32 = (1e10 - ((((-2.25 / float32((0.0 + v2))) / 1.5) * float32((float64(float32((v2 - 3.14159))) * v2))) * 0.0))
                    _ = fi129
                    _ = v130
                    _ = v131
                    _ = v132
                }
                v3 += genFn0(-32768, 312, 23)
            } else {
                v1 = (v1 << ((65535) & 63))
                v3 += v3
                var v133 int16 = -32768
                _ = v133
            }
            for fi134 := 0; fi134 < 10; fi134++ {
                var v135 uint16 = 2
                _ = len(v4)
                _ = fi134
                _ = v135
            }
            for fi136 := 0; fi136 < 14; fi136++ {
                v137 := GenS3{F0: 72}
                var v138 int = fi136
                _ = fi136
                _ = v137
                _ = v138
            }
        } else {
            var v139 uint = 481
            v139 += (^(v139))
            _ = v139
        }
        fmt.Println(v1)
    }
    var v140 float64 = (0.0 - v2)
    var v141 uint16 = 551
    v142 := GenS1{F0: 490, F1: (v3 - (^(v3)))}
    fmt.Println("v1:", v1)
    fmt.Println("v140:", v140)
    fmt.Println("v141:", v141)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v140
    _ = v141
    _ = v142
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
