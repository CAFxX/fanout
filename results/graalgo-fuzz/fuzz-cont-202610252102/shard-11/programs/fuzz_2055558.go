package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32; F1 uint8 }
type GenS9 struct { F0 uint; F1 string }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint64; F1 int16 }
type GenS12 struct { F0 string; F1 int16; F2 int64 }
type GenS13 struct { F0 int; F1 uint8 }
type GenS14 struct { F0 int8 }
func genFn0() int8 {
    var v7 int32 = 0
    if true {
        switch "a" {
        case "\n\t":
            v7 = (v7 & v7)
        }
        fmt.Println(v7)
        fmt.Println(v7)
    }
    var v8 bool = false
    var v9 int64 = 42
    _ = v7
    _ = v8
    _ = v9
    return -128
}

func main() {
    var v1 bool = false
    var v2 int8 = -89
    v3 := make(map[int]uint32)
    v4 := []int{-373, 0}
    v5 := make(map[uint32]int)
    v6 := make(map[string]string)
    v2 -= (v2 * v2)
    v10 := GenS1{F0: 42, F1: "\n\t"}
    if v1 {
        switch ("foo bar" + "") {
        case "":
            switch "foo bar" {
            case "a":
                v11 := GenS2{F0: (v2 / ((genFn0()) | 1)), F1: 1}
                _ = v11
            default:
                for fi12 := 0; fi12 < 14; fi12++ {
                    for fi13 := 0; fi13 < 14; fi13++ {
                        v2 -= v2
                        if (!((fi13 - fi13) == (fi13 - 42))) {
                            var v14 uint64 = 0
                            _ = v14
                        }
                        _ = fi13
                    }
                    _ = fi12
                }
            }
        case "hello":
            fmt.Println(v1, v2)
        case "\n\t":
            delete(v3, -581)
        }
        if v1 {
            switch 39 {
            case 9223372036854775807:
                var v15 uint64 = 18446744073709551615
                _ = v15
            case 814:
                fmt.Println(v1, v2)
            case 4294967295:
                if ((1e10 - 3.14159) > ((3.14159 + 1e10) - 1.5)) {
                    for fi16 := 0; fi16 < 17; fi16++ {
                        if v1 {
                            if v1 {
                                var v17 float64 = 1.5
                                if true {
                                    var v18 uint = 4294967295
                                    _ = v18
                                }
                                switch fi16 {
                                case -582:
                                    var v19 uint8 = 42
                                    _ = v19
                                }
                                _ = v17
                            } else {
                                fmt.Println(fi16)
                                switch "hello" {
                                case "\n\t":
                                    v20 := GenS2{F0: v2, F1: 42}
                                    _ = v20
                                default:
                                    for fi21 := 0; fi21 < 10; fi21++ {
                                        v1 = v1
                                        if (641 >= -1) {
                                            v22 := GenS2{F0: v2, F1: 32767}
                                            switch "foo bar" {
                                            case "a":
                                                fmt.Println(fi16, v1, v2)
                                            }
                                            _ = v22
                                        }
                                        v23 := GenS2{F0: (-(v2)), F1: 230}
                                        v24 := GenS3{F0: 1000}
                                        if true {
                                            if (v1 || (!true)) {
                                                var v25 int32 = -1
                                                var v26 float64 = 1.5
                                                var v27 uint16 = 1
                                                _ = v25
                                                _ = v26
                                                _ = v27
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v28 bool = v1
                                            _ = v28
                                        }
                                        _ = fi21
                                        _ = v23
                                        _ = v24
                                    }
                                }
                                v2 = (v2 - 1)
                            }
                            var v29 int64 = 786
                            _ = v29
                        } else {
                            fmt.Println(fi16, v2)
                            if len(v4) > 0 {
                                v4[0] = (^(fi16))
                            }
                        }
                        for fi30 := 0; fi30 < 10; fi30++ {
                            switch "" {
                            case "\n\t":
                                v2 = (v2 - (v2 % (((v2 * genFn0())) | 1)))
                            case "":
                                var v31 float32 = 0.1
                                _ = v31
                            default:
                                var v32 int64 = 9223372036854775807
                                _ = v32
                            }
                            v2 = (v2 ^ (-(v2)))
                            _ = fi30
                        }
                        _ = fi16
                    }
                    v1 = true
                    if ((v1 && false) && v1) {
                        fmt.Println(v1)
                    }
                } else {
                    if true {
                        var v33 uint32 = 567
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 7; fi34++ {
                        if true {
                            var v35 float32 = 3.14159
                            for fi36 := 0; fi36 < 14; fi36++ {
                                var v37 int32 = 1
                                _ = len(v4)
                                if false {
                                    v38 := GenS2{F0: (^(v2)), F1: 32767}
                                    _ = v38
                                } else {
                                    fmt.Println(v1, fi34)
                                    fmt.Println(v1, fi36, v1)
                                }
                                _ = fi36
                                _ = v37
                            }
                            if v1 {
                                switch "a" {
                                case "foo bar":
                                    v39 := GenS1{F0: 2, F1: "a"}
                                    _ = v39
                                default:
                                    var v40 uint16 = 187
                                    _ = v40
                                }
                                if (genFn0() > (v2 + 100)) {
                                    var v41 uint32 = 4294967295
                                    var v42 uint32 = v41
                                    switch "\n\t" {
                                    case "":
                                        var v43 uint32 = 2147483647
                                        _ = v43
                                    case "\n\t":
                                        if v1 {
                                            if (uint(v42) <= 18446744073709551615) {
                                                var v44 int64 = 1
                                                _ = v44
                                            } else {
                                                if v1 {
                                                    for fi45 := 0; fi45 < 10; fi45++ {
                                                        if (!(1 == 916)) {
                                                            var v46 int64 = -9223372036854775808
                                                            var v47 uint16 = 609
                                                            _ = v46
                                                            _ = v47
                                                        } else {
                                                            var v48 string = "\n\t"
                                                            v42 = (^(v42))
                                                            _ = v48
                                                        }
                                                        var v49 uint16 = 290
                                                        _ = fi45
                                                        _ = v49
                                                    }
                                                } else {
                                                    v50 := GenS1{F0: 42, F1: (("\n\t" + "hello") + ("foo bar" + "hello"))}
                                                    _ = v50
                                                }
                                                fmt.Println(fi34, v41, v1)
                                            }
                                            var v51 bool = false
                                            _ = v51
                                        }
                                    default:
                                        if v1 {
                                            fmt.Println(v41, v35, fi34)
                                            switch "foo bar" {
                                            case "foo bar":
                                                v52 := GenS4{F0: v42, F1: 4294967295}
                                                _ = v52
                                            case "\n\t":
                                                _ = len(v5)
                                            case "hello":
                                                var v53 int32 = 1
                                                _ = v53
                                            }
                                        }
                                    }
                                    _ = v41
                                    _ = v42
                                } else {
                                    var v54 int = fi34
                                    _ = v54
                                }
                                if (!(572 == 961)) {
                                    var v55 uint32 = 778
                                    fmt.Println(fi34)
                                    _ = v55
                                } else {
                                    v56 := GenS5{F0: "hello", F1: 1}
                                    v35 = (v35 - (v35 * v35))
                                    _ = v56
                                }
                            }
                            _ = v35
                        } else {
                            if v1 {
                                fmt.Println(v2, fi34)
                            }
                            for fi57 := 0; fi57 < 15; fi57++ {
                                if v1 {
                                    if (387 < 42) {
                                        v58 := GenS6{F0: 127}
                                        fmt.Println(v1, fi57, v2)
                                        fmt.Println(fi57, v2)
                                        _ = v58
                                    }
                                    v2 = v2
                                }
                                var v59 int64 = 42
                                _ = len(v4)
                                _ = fi57
                                _ = v59
                            }
                        }
                        _ = fi34
                    }
                }
            }
            for fi60 := 0; fi60 < 15; fi60++ {
                v61 := GenS7{F0: (2 == 42)}
                for fi62 := 0; fi62 < 10; fi62++ {
                    v6["\n\t"] = ""
                    v1 = false
                    if v1 {
                        var v63 int8 = (^(v2))
                        v2 = 0
                        if (!(v1 || (("foo bar" <= "\n\t") || v1))) {
                            var v64 int8 = genFn0()
                            var v65 int16 = 0
                            _ = v64
                            _ = v65
                        } else {
                            v66 := GenS8{F0: 2147483648, F1: 2}
                            v67 := GenS9{F0: 255, F1: "foo bar"}
                            var v68 int32 = 2147483647
                            _ = v66
                            _ = v67
                            _ = v68
                        }
                        fmt.Println(v1, v63)
                        _ = v63
                    }
                    _ = fi62
                }
                _ = fi60
                _ = v61
            }
        }
        var v69 int8 = v2
        _ = v69
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "hello" {
    case "hello":
        for fi70 := 0; fi70 < 2; fi70++ {
            var v71 int8 = (^(v2))
            v1 = false
            for fi72 := 0; fi72 < 14; fi72++ {
                fmt.Println(v1, fi70)
                _ = fi72
            }
            _ = fi70
            _ = v71
        }
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        if len(v4) > 0 {
            v4[0] = 42
        }
    }
    _ = len(v5)
    fmt.Println(v1)
    delete(v3, -185)
    var v73 int8 = -101
    var v74 uint = 0
    v73 = (v2 / ((1) | 1))
    _ = len(v4)
    v1 = v1
    if (false || v1) {
        fmt.Println(v73, v73, v73)
        switch ("foo bar" + ("hello" + "\n\t")) {
        case "\n\t":
            var v75 int = 42
            _ = v75
        case "hello":
            v2 = genFn0()
        }
        for fi76 := 0; fi76 < 14; fi76++ {
            if v1 {
                var v77 bool = true
                switch -9223372036854775808 {
                case -1:
                    delete(v3, (fi76 & (fi76 - fi76)))
                }
                _ = v77
            } else {
                if (1.5 < 1e10) {
                    fmt.Println(v1, fi76)
                    v74 -= 488
                }
            }
            _ = fi76
        }
    } else {
        v78 := GenS7{F0: (v74 != (v74 + (v74 << ((v74) & 63))))}
        for fi79 := 0; fi79 < 5; fi79++ {
            v80 := GenS10{F0: 4294967295}
            _ = fi79
            _ = v80
        }
        _ = v78
    }
    switch "foo bar" {
    case "hello":
        switch "a" {
        case "hello":
            if true {
                v73 -= v73
                v81 := GenS5{F0: "", F1: 42}
                if false {
                    v82 := GenS4{F0: 693, F1: 2}
                    _ = v82
                }
                _ = v81
            }
        case "\n\t":
            v4 = append(v4, 127)
        case "a":
            var v83 string = ""
            _ = v83
        }
    default:
        switch -9223372036854775808 {
        case 127:
            v84 := GenS11{F0: 18446744073709551615, F1: 32767}
            _ = v84
        default:
            if (1000 != 1) {
                var v85 int16 = 1
                if (325 == 9223372036854775807) {
                    v86 := GenS12{F0: "", F1: v85, F2: -595}
                    for fi87 := 0; fi87 < 13; fi87++ {
                        var v88 int16 = (v85 + (v85 + -1))
                        _ = fi87
                        _ = v88
                    }
                    for fi89 := 0; fi89 < 20; fi89++ {
                        v90 := GenS1{F0: 244, F1: "\n\t"}
                        switch "hello" {
                        case "":
                            _ = len(v3)
                        case "a":
                            if v1 {
                                switch (-(fi89)) {
                                case 2147483647:
                                    for fi91 := 0; fi91 < 3; fi91++ {
                                        switch "hello" {
                                        case "a":
                                            var v92 int8 = genFn0()
                                            _ = v92
                                        default:
                                            var v93 int16 = v85
                                            _ = v93
                                        }
                                        _ = fi91
                                    }
                                case -9223372036854775808:
                                    v94 := GenS7{F0: false}
                                    _ = v94
                                case 9223372036854775807:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            } else {
                                if (454 >= 0) {
                                    v2 = genFn0()
                                    for fi95 := 0; fi95 < 19; fi95++ {
                                        fmt.Println(v73, fi95)
                                        _ = fi95
                                    }
                                }
                                v3[42] = 42
                                v85 -= v85
                            }
                        case "foo bar":
                            v1 = v1
                        }
                        if (!true) {
                            for fi96 := 0; fi96 < 7; fi96++ {
                                v97 := GenS13{F0: (^(fi96)), F1: 165}
                                _ = fi96
                                _ = v97
                            }
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                for fi98 := 0; fi98 < 7; fi98++ {
                                    v99 := GenS14{F0: 60}
                                    if (0.1 <= 3.14159) {
                                        var v100 uint32 = 65535
                                        v74 -= v74
                                        switch "\n\t" {
                                        case "hello":
                                            v85 -= 107
                                        case "a":
                                            var v101 string = ""
                                            _ = v101
                                        default:
                                            v100 += (v100 % ((610) | 1))
                                        }
                                        _ = len(v4)
                                        _ = v100
                                    }
                                    _ = fi98
                                    _ = v99
                                }
                            }
                            if true {
                                var v102 int64 = -1
                                v103 := GenS6{F0: 127}
                                _ = v102
                                _ = v103
                            } else {
                                fmt.Println(v1)
                                v73 += -1
                                var v104 int16 = v85
                                _ = v104
                            }
                        }
                        _ = fi89
                        _ = v90
                    }
                    _ = v86
                }
                var v105 int64 = -369
                _ = v105
                _ = v85
            }
        }
    }
    switch 0 {
    case 2147483647:
        v106 := GenS3{F0: 32767}
        _ = v106
    case 4294967295:
        if true {
            v73 = v73
            if (4294967295 >= 2147483647) {
                v107 := GenS8{F0: 759, F1: 0}
                _ = v107
            }
            v108 := GenS9{F0: 4294967295, F1: ("a" + "\n\t")}
            _ = v108
        } else {
            var v109 uint32 = 2147483648
            _ = v109
        }
    case 802:
        fmt.Println(v73)
    }
    v74 -= (v74 ^ v74)
    fmt.Println(v74)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
}
