package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 string }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 bool; F2 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 uint16; F2 string }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int; F1 string }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int16; F1 uint64 }
type GenS13 struct { F0 int32; F1 uint8; F2 int64 }
func genFn0() uint32 {
    var v5 int = -47
    v6 := []int{(v5 ^ v5), v5, v5, (^(v5))}
    switch v5 {
    case 55:
        var v7 uint32 = 0
        _ = v7
    default:
        for fi8 := 0; fi8 < 4; fi8++ {
            v5 = -9223372036854775808
            if (100 == 1) {
                fmt.Println(fi8, v5)
                for fi9 := 0; fi9 < 6; fi9++ {
                    if len(v6) > 0 {
                        v6[0] = 32767
                    }
                    _ = fi9
                }
                v5 = (v5 & -9223372036854775808)
            } else {
                switch "foo bar" {
                case "hello":
                    var v10 int16 = 587
                    _ = v10
                default:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                fmt.Println(fi8, v5)
                var v11 uint32 = 0
                _ = v11
            }
            _ = fi8
        }
    }
    v5 = 262
    if len(v6) > 0 {
        v6[0] = -2147483648
    }
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    var v3 uint8 = 42
    var v4 int8 = -70
    if (665 > 929) {
        v4 = (v4 % (((v4 >> ((65535) & 63))) | 1))
        if false {
            if (-32768 >= 32767) {
                var v12 uint = 57
                fmt.Println(v3, v12)
                _ = v12
            }
            v3 = (v3 - 128)
        } else {
            switch "foo bar" {
            case "":
                v13 := GenS1{F0: 813, F1: 42, F2: "foo bar"}
                _ = v13
            }
            switch "foo bar" {
            case "foo bar":
                v4 = 100
            }
        }
    } else {
        var v14 float64 = 3.14159
        v15 := GenS2{F0: v3, F1: 4294967295, F2: 425}
        for fi16 := 0; fi16 < 17; fi16++ {
            switch "hello" {
            case "foo bar":
                v3 += 255
            case "\n\t":
                if false {
                    var v17 bool = (2 != 534)
                    v4 = (v4 | (v4 - (v4 << ((2) & 63))))
                    _ = v17
                } else {
                    var v18 uint64 = 18446744073709551615
                    _ = v18
                }
            case "a":
                if false {
                    if (fi16 > (fi16 << ((4294967295) & 63))) {
                        var v19 int8 = (v4 >> ((65535) & 63))
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 8; fi20++ {
                        if true {
                            if false {
                                fmt.Println(v4, fi20)
                            }
                        }
                        var v21 uint = 4294967295
                        _ = fi20
                        _ = v21
                    }
                } else {
                    v3 += (v3 / ((0) | 1))
                    if (!false) {
                        var v22 bool = (false && true)
                        var v23 uint64 = 126
                        _ = v22
                        _ = v23
                    }
                }
            }
            for fi24 := 0; fi24 < 13; fi24++ {
                var v25 bool = false
                v4 = v4
                _ = fi24
                _ = v25
            }
            v26 := GenS3{F0: 4294967295}
            _ = fi16
            _ = v26
        }
        _ = v14
        _ = v15
    }
    fmt.Println(v4, v3)
    v27 := []int{383, 1}
    v28 := GenS1{F0: 42, F1: 4294967295, F2: ""}
    var v29 int = -675
    delete(v2, v29)
    if (0 >= 0) {
        v30 := GenS4{F0: 116, F1: (!true), F2: (v4 >> ((65535) & 63))}
        fmt.Println(v4)
        v29 -= v29
        _ = v30
    }
    switch "a" {
    case "hello":
        switch "hello" {
        case "hello":
            v31 := GenS1{F0: 545, F1: 2, F2: ""}
            _ = v31
        case "\n\t":
            _ = len(v1)
        }
    case "a":
        var v32 int32 = 1
        _ = v32
    default:
        switch "" {
        case "":
            v29 = (v29 & (v29 << ((42) & 63)))
        default:
            v3 = (^(v3))
        }
    }
    v33 := GenS5{F0: 32767}
    switch "" {
    case "a":
        if true {
            var v34 bool = (!((0.0 == (1e10 + 1.5)) || false))
            _ = v34
        } else {
            switch "a" {
            case "hello":
                v35 := GenS5{F0: int16(v29)}
                _ = v35
            case "foo bar":
                v29 -= v29
            case "\n\t":
                for fi36 := 0; fi36 < 16; fi36++ {
                    fmt.Println(v29)
                    var v37 uint32 = 496
                    v37 -= genFn0()
                    _ = fi36
                    _ = v37
                }
            default:
                fmt.Println(v29, v3)
            }
            var v38 int = v29
            _ = len(v27)
            _ = v38
        }
    case "foo bar":
        if (("a" + (("foo bar" + "\n\t") + "\n\t")) > "\n\t") {
            var v39 uint32 = 65535
            v40 := GenS1{F0: 873, F1: (v39 | 42), F2: "a"}
            switch 127 {
            case -36:
                v41 := GenS5{F0: 1}
                _ = v41
            case 1:
                var v42 uint = 2
                _ = v42
            case 127:
                switch ("foo bar" + "") {
                case "\n\t":
                    v4 = v4
                case "foo bar":
                    if (3.14159 == 3.14159) {
                        v4 = (v4 | 42)
                        if (0 == 65535) {
                            var v43 int16 = 0
                            _ = len(v27)
                            _ = v43
                        } else {
                            if (1 >= 42) {
                                v44 := GenS3{F0: v39}
                                switch "hello" {
                                case "a":
                                    v4 += (v4 | (v4 ^ v4))
                                case "foo bar":
                                    for fi45 := 0; fi45 < 12; fi45++ {
                                        v29 = (fi45 & v29)
                                        _ = fi45
                                    }
                                case "":
                                    if (((0.1 + (3.14159 - 1e10)) - 0.1) == (3.14159 + 1.5)) {
                                        switch (v29 * (v29 << ((0) & 63))) {
                                        case -9223372036854775808:
                                            v46 := GenS3{F0: (^(v39))}
                                            _ = v46
                                        case 777:
                                            v47 := GenS4{F0: v4, F1: true, F2: v4}
                                            _ = v47
                                        case 1:
                                            var v48 int16 = -32768
                                            _ = v48
                                        default:
                                            var v49 int64 = 0
                                            _ = v49
                                        }
                                        switch (v29 | 127) {
                                        case 32768:
                                            v50 := v1[v29]
                                            _ = v50
                                        case 2147483647:
                                            for fi51 := 0; fi51 < 20; fi51++ {
                                                var v52 bool = false
                                                v29 += fi51
                                                _ = fi51
                                                _ = v52
                                            }
                                        case 128:
                                            if (1.5 == 0.0) {
                                                switch (v29 ^ v29) {
                                                case 32767:
                                                    var v53 int64 = int64(v4)
                                                    _ = v53
                                                default:
                                                    if false {
                                                        if len(v27) >= 2 {
                                                            v27 = v27[1:len(v27)]
                                                        }
                                                        v54 := v1[(v29 * (v29 ^ (v29 % ((v29) | 1))))]
                                                        fmt.Println(v39)
                                                        _ = v54
                                                    }
                                                }
                                            }
                                        }
                                        for fi55 := 0; fi55 < 11; fi55++ {
                                            v39 = (v39 % ((1) | 1))
                                            _ = fi55
                                        }
                                    } else {
                                        v27 = append(v27, (v29 & 128))
                                        fmt.Println(v4, v39)
                                    }
                                default:
                                    v3 += v3
                                }
                                if len(v27) >= 2 {
                                    v27 = v27[1:len(v27)]
                                }
                                _ = v44
                            } else {
                                var v56 bool = false
                                if len(v27) > 0 {
                                    v27[0] = 1
                                }
                                v57 := GenS6{F0: false, F1: 294, F2: "\n\t"}
                                _ = v56
                                _ = v57
                            }
                        }
                        for fi58 := 0; fi58 < 9; fi58++ {
                            switch ((("hello" + "hello") + ("foo bar" + ("hello" + ("" + "foo bar")))) + "a") {
                            case "foo bar":
                                var v59 float64 = ((3.14159 / 1.0) + float64(3.14159))
                                _ = v59
                            default:
                                for fi60 := 0; fi60 < 1; fi60++ {
                                    v39 -= v39
                                    v4 = (v4 / (((^(v4))) | 1))
                                    _ = fi60
                                }
                            }
                            switch (fi58 - (v29 << ((1) & 63))) {
                            case -689:
                                var v61 uint64 = 2
                                _ = v61
                            default:
                                v4 -= (v4 / ((100) | 1))
                            }
                            v62 := GenS4{F0: v4, F1: (!(!false)), F2: v4}
                            _ = fi58
                            _ = v62
                        }
                    } else {
                        if false {
                            for fi63 := 0; fi63 < 19; fi63++ {
                                v3 = (v3 ^ v3)
                                _ = len(v27)
                                var v64 int32 = -108
                                _ = fi63
                                _ = v64
                            }
                        }
                        v4 += v4
                        v39 += v39
                    }
                default:
                    if (-2147483648 <= 540) {
                        v4 = v4
                        v65 := GenS7{F0: 1}
                        _ = v65
                    } else {
                        fmt.Println(v3, v29, v3)
                        switch "" {
                        case "\n\t":
                            fmt.Println(v4, v3, v29)
                        case "":
                            if (false || true) {
                                if true {
                                    var v66 int = (v29 + v29)
                                    _ = v66
                                }
                                var v67 uint8 = 255
                                v67 = v3
                                _ = v67
                            } else {
                                fmt.Println(v39, v29, v29)
                                var v68 float32 = 1.5
                                _ = v68
                            }
                        }
                    }
                }
            default:
                switch (^(v29)) {
                case 1:
                    if true {
                        if true {
                            v69 := GenS8{F0: 823}
                            v70 := GenS8{F0: -1}
                            for fi71 := 0; fi71 < 15; fi71++ {
                                v72 := GenS8{F0: 770}
                                v73 := GenS9{F0: fi71, F1: "\n\t"}
                                _ = fi71
                                _ = v72
                                _ = v73
                            }
                            _ = v69
                            _ = v70
                        } else {
                            _ = len(v27)
                            v3 = 2
                        }
                        _ = len(v27)
                        v27 = append(v27, -9223372036854775808)
                    }
                case 895:
                    v1[v29] = (v29 - (-(v29)))
                case -1:
                    v3 = v3
                default:
                    v3 -= v3
                }
            }
            _ = v39
            _ = v40
        } else {
            v3 = 160
        }
    case "\n\t":
        fmt.Println(v3, v29)
    }
    v74 := GenS10{F0: 2}
    v3 += v3
    switch (v29 / (((v29 >> ((161) & 63))) | 1)) {
    case 0:
        switch 0 {
        case 2147483647:
            fmt.Println(v29, v29)
        }
    default:
        switch (^(v29)) {
        case -2147483648:
            if true {
                var v75 float32 = 0.1
                v4 -= v4
                _ = v75
            } else {
                switch ("a" + ("a" + ("hello" + ("foo bar" + "hello")))) {
                case "foo bar":
                    switch "a" {
                    case "hello":
                        if ((1e10 * 3.14159) < -2.25) {
                            switch ("\n\t" + "foo bar") {
                            case "foo bar":
                                v2[v29] = ("\n\t" + "")
                            case "a":
                                for fi76 := 0; fi76 < 9; fi76++ {
                                    for fi77 := 0; fi77 < 10; fi77++ {
                                        v29 += 511
                                        switch "\n\t" {
                                        case "":
                                            v78 := GenS6{F0: (true || (false && (1 <= (fi76 >> ((2) & 63))))), F1: 1, F2: "a"}
                                            var v79 uint8 = v3
                                            _ = v78
                                            _ = v79
                                        default:
                                            v29 -= (-(fi77))
                                        }
                                        if true {
                                            v80 := GenS11{F0: fi77}
                                            v4 = (v4 | v4)
                                            _ = v80
                                        }
                                        v81 := GenS11{F0: v29}
                                        _ = fi77
                                        _ = v81
                                    }
                                    var v82 int64 = 1
                                    _ = fi76
                                    _ = v82
                                }
                            case "\n\t":
                                fmt.Println(v3, v4)
                            }
                            v83 := GenS11{F0: (^(v29))}
                            _ = v83
                        }
                    }
                case "\n\t":
                    for fi84 := 0; fi84 < 4; fi84++ {
                        var v85 bool = false
                        v4 = v4
                        fmt.Println(v85, v29, v3)
                        _ = fi84
                        _ = v85
                    }
                default:
                    var v86 int8 = -128
                    _ = v86
                }
            }
        case 9223372036854775807:
            fmt.Println(v3, v29)
        }
    }
    if (-1 < 1) {
        var v87 string = "hello"
        if true {
            delete(v1, 32767)
            v29 += 858
            for fi88 := 0; fi88 < 7; fi88++ {
                var v89 int = (fi88 / (((fi88 | (^(v29)))) | 1))
                _ = fi88
                _ = v89
            }
        }
        _ = v87
    } else {
        v90 := GenS12{F0: 1, F1: 584}
        switch (v29 << ((65535) & 63)) {
        case 0:
            v91 := GenS13{F0: 1, F1: v3, F2: -9223372036854775808}
            _ = v91
        case -727:
            var v92 uint = uint(v3)
            _ = v92
        }
        v29 += (v29 | (v29 ^ (v29 >> ((65535) & 63))))
        _ = v90
    }
    v93 := GenS12{F0: 0, F1: 886}
    v4 -= (-(v4))
    var v94 string = ("" + "\n\t")
    v29 += (v29 + (^(v29)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v74
    _ = v93
    _ = v94
}
