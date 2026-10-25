package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 uint32; F1 uint32; F2 bool }
type GenS9 struct { F0 int8; F1 bool; F2 int }
type GenS10 struct { F0 string }
func genFn0() int16 {
    v4 := GenS1{F0: 263}
    var v5 uint16 = 36
    switch ("hello" + ("hello" + "hello")) {
    case "hello":
        var v6 float64 = 0.0
        _ = v6
    }
    var v7 int32 = 584
    _ = v4
    _ = v5
    _ = v7
    return int16(v7)
}

func main() {
    var v1 float64 = 1e10
    var v2 int8 = 100
    var v3 int8 = -128
    var v8 uint8 = 124
    if (!(329 < 164)) {
        for fi9 := 0; fi9 < 1; fi9++ {
            if false {
                var v10 float64 = v1
                for fi11 := 0; fi11 < 17; fi11++ {
                    fmt.Println(v1, fi11, v3)
                    var v12 uint8 = v8
                    _ = fi11
                    _ = v12
                }
                _ = v10
            } else {
                var v13 uint8 = 234
                v1 -= (-2.25 + 0.0)
                _ = v13
            }
            if false {
                v14 := GenS1{F0: 422}
                if true {
                    v8 -= v8
                    for fi15 := 0; fi15 < 11; fi15++ {
                        var v16 float64 = (v1 + 0.1)
                        var v17 int16 = 1000
                        v2 = (v2 + v3)
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                } else {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        v19 := []int{}
                        _ = fi18
                        _ = v19
                    }
                    v3 = (v2 / ((v2) | 1))
                    var v20 uint = 42
                    _ = v20
                }
                _ = v14
            }
            _ = fi9
        }
        v21 := []int{32768, -2147483648}
        _ = v21
    } else {
        for fi22 := 0; fi22 < 7; fi22++ {
            for fi23 := 0; fi23 < 13; fi23++ {
                var v24 int = (fi23 ^ (fi23 << ((0) & 63)))
                v25 := []int{int(v8), (fi22 * fi22), fi23, fi22}
                if ((2 == 818) && (!true)) {
                    if len(v25) > 0 {
                        v25[0] = fi23
                    }
                    fmt.Println(fi22, v8)
                    var v26 uint64 = 658
                    _ = v26
                }
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi22
        }
        var v27 int64 = -1
        _ = v27
    }
    var v28 int = -729
    v29 := GenS1{F0: 1}
    var v30 bool = true
    var v31 bool = (138 == 778)
    fmt.Println(v8, v8)
    var v32 int = (v28 ^ (v28 * v28))
    v33 := GenS1{F0: 18446744073709551615}
    v34 := []int{833}
    var v35 int = (v28 | 501)
    if false {
        var v36 int64 = -113
        for fi37 := 0; fi37 < 17; fi37++ {
            switch "hello" {
            case "":
                for fi38 := 0; fi38 < 11; fi38++ {
                    var v39 uint64 = 18446744073709551615
                    _ = fi38
                    _ = v39
                }
            default:
                if len(v34) >= 2 {
                    v34 = v34[1:len(v34)]
                }
            }
            if (-2147483648 <= (v32 - (fi37 << ((42) & 63)))) {
                fmt.Println(v36)
                for fi40 := 0; fi40 < 4; fi40++ {
                    var v41 int32 = 2147483647
                    _ = fi40
                    _ = v41
                }
                for fi42 := 0; fi42 < 2; fi42++ {
                    switch (("hello" + (("a" + "\n\t") + "hello")) + "foo bar") {
                    case "":
                        v30 = (48 != 929)
                        var v43 uint64 = 2
                        _ = v43
                    case "\n\t":
                        v2 += (v3 & v2)
                        var v44 int8 = (v2 ^ v3)
                        var v45 float64 = (v1 - 0.0)
                        var v46 float64 = v45
                        var v47 int = (fi42 << ((467) & 63))
                        _ = v44
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    _ = fi42
                }
            } else {
                for fi48 := 0; fi48 < 15; fi48++ {
                    v31 = v30
                    switch "" {
                    case "foo bar":
                        if true {
                            if len(v34) > 0 {
                                v34[0] = v28
                            }
                            if len(v34) > 0 {
                                v34[0] = v32
                            }
                        }
                    case "a":
                        v49 := GenS1{F0: 0}
                        _ = v49
                    case "":
                        v28 = v35
                        fmt.Println(v2)
                        if len(v34) > 0 {
                            v34[0] = (v35 % ((v28) | 1))
                        }
                    }
                    fmt.Println(fi37, v35, v32)
                    _ = fi48
                }
                if v30 {
                    var v50 string = "hello"
                    _ = v50
                } else {
                    v35 = (-(v32))
                }
                for fi51 := 0; fi51 < 5; fi51++ {
                    var v52 uint = 4294967295
                    v31 = (0 == 2)
                    _ = len(v34)
                    _ = fi51
                    _ = v52
                }
            }
            _ = fi37
        }
        _ = v36
    } else {
        v28 = v35
        for fi53 := 0; fi53 < 7; fi53++ {
            v35 = v32
            v54 := GenS1{F0: 2}
            v55 := GenS1{F0: 500}
            _ = fi53
            _ = v54
            _ = v55
        }
    }
    v2 = v3
    fmt.Println(v1)
    for fi56 := 0; fi56 < 15; fi56++ {
        if true {
            var v57 bool = v30
            _ = v57
        } else {
            for fi58 := 0; fi58 < 18; fi58++ {
                switch (^(v28)) {
                case -9223372036854775808:
                    var v59 uint64 = 4294967295
                    _ = v59
                case 718:
                    fmt.Println(v31)
                case 127:
                    if v30 {
                        v8 += (v8 | 122)
                        if (!((v30 && (2 != 1)) || (18446744073709551615 >= 4294967295))) {
                            var v60 string = ("\n\t" + "hello")
                            if v30 {
                                var v61 int16 = 211
                                var v62 uint32 = 1000
                                v60 = v60
                                _ = v61
                                _ = v62
                            }
                            if (v31 && ((^(fi58)) >= 2147483647)) {
                                switch ("" + v60) {
                                case "":
                                    v1 += 0.1
                                case "\n\t":
                                    v3 = v2
                                case "hello":
                                    v31 = (0 != 0)
                                    var v63 int32 = 42
                                    _ = v63
                                default:
                                    v2 = 0
                                }
                                if (!(0 < 42)) {
                                    v31 = (0 != 531)
                                    v64 := GenS2{F0: "", F1: 2147483647}
                                    fmt.Println(v30, fi58)
                                    _ = v64
                                }
                            } else {
                                var v65 string = "a"
                                var v66 bool = (!((("a" <= v60) || ((0 != 331) || v30)) || false))
                                v67 := GenS1{F0: 446}
                                _ = v65
                                _ = v66
                                _ = v67
                            }
                            fmt.Println(v1, v32, fi58)
                            _ = v60
                        } else {
                            switch 32768 {
                            case -9223372036854775808:
                                var v68 string = ""
                                _ = v68
                            default:
                                if (!true) {
                                    v8 = 255
                                } else {
                                    fmt.Println(v3, v2, v3)
                                }
                            }
                            var v69 uint = 0
                            fmt.Println(v3, fi58)
                            _ = v69
                        }
                    } else {
                        v70 := GenS3{F0: false}
                        v28 += fi58
                        _ = len(v34)
                        _ = v70
                    }
                default:
                    v1 -= (v1 + v1)
                }
                _ = fi58
            }
            switch (-(v32)) {
            case 1:
                for fi71 := 0; fi71 < 12; fi71++ {
                    switch "hello" {
                    case "foo bar":
                        v2 = v2
                    }
                    fmt.Println(v2, v2)
                    if len(v34) > 0 {
                        v34[0] = 521
                    }
                    _ = fi71
                }
            case 127:
                var v72 int64 = -9223372036854775808
                _ = v72
            case -141:
                fmt.Println(v28)
            }
            if v30 {
                if true {
                    v73 := GenS2{F0: "\n\t", F1: 0}
                    var v74 uint16 = 42
                    switch (("a" + "\n\t") + ((("\n\t" + ("hello" + "\n\t")) + "foo bar") + "foo bar")) {
                    case "a":
                        for fi75 := 0; fi75 < 14; fi75++ {
                            v76 := GenS4{F0: 1, F1: 2}
                            switch 0 {
                            case 9223372036854775807:
                                switch ("hello" + "a") {
                                case "foo bar":
                                    var v77 uint = 255
                                    v1 = 1e10
                                    var v78 int = (v35 / (((fi56 - fi75)) | 1))
                                    _ = v77
                                    _ = v78
                                case "a":
                                    if len(v34) > 0 {
                                        v34[0] = v28
                                    }
                                    if ((v3 * v3) <= (v2 << ((331) & 63))) {
                                        v8 += v8
                                        var v79 uint16 = v74
                                        v80 := GenS5{F0: 0}
                                        _ = v79
                                        _ = v80
                                    } else {
                                        if v31 {
                                            var v81 int8 = 0
                                            switch v35 {
                                            case 745:
                                                v74 -= v74
                                            }
                                            switch int(v2) {
                                            case 0:
                                                v8 = 42
                                            default:
                                                fmt.Println(fi56, v3, v32)
                                            }
                                            _ = len(v34)
                                            _ = v81
                                        }
                                    }
                                case "":
                                    v82 := GenS6{F0: 1, F1: 316}
                                    if false {
                                        switch "a" {
                                        case "hello":
                                            var v83 uint16 = (v74 % ((405) | 1))
                                            _ = v83
                                        }
                                        v31 = true
                                        v34 = append(v34, fi56)
                                    } else {
                                        v30 = (-680 <= 42)
                                        var v84 float32 = (1.5 / -2.25)
                                        v30 = false
                                        _ = v84
                                    }
                                    v85 := GenS2{F0: "hello", F1: 737}
                                    _ = v82
                                    _ = v85
                                }
                            case -1:
                                var v86 uint64 = 18446744073709551615
                                _ = v86
                            case -9223372036854775808:
                                if len(v34) >= 2 {
                                    v34 = v34[1:len(v34)]
                                }
                            default:
                                if len(v34) > 0 {
                                    v34[0] = 128
                                }
                            }
                            _ = fi75
                            _ = v76
                        }
                    }
                    _ = v73
                    _ = v74
                }
            }
        }
        _ = fi56
    }
    for fi87 := 0; fi87 < 4; fi87++ {
        v88 := GenS5{F0: 18446744073709551615}
        v28 = (v35 | v35)
        _ = fi87
        _ = v88
    }
    v35 = int(v8)
    switch int(v8) {
    case 128:
        v1 = 1.5
    case 814:
        var v89 int = (v32 / (((^(v32))) | 1))
        _ = v89
    default:
        if v30 {
            if ((v8 | (v8 - 45)) < v8) {
                v30 = true
                v90 := GenS7{F0: -2147483648, F1: v3}
                _ = v90
            }
        } else {
            switch v28 {
            case -9223372036854775808:
                if (!false) {
                    for fi91 := 0; fi91 < 17; fi91++ {
                        for fi92 := 0; fi92 < 9; fi92++ {
                            if len(v34) >= 2 {
                                v34 = v34[1:len(v34)]
                            }
                            _ = fi92
                        }
                        _ = fi91
                    }
                    v32 = v32
                } else {
                    var v93 int16 = 42
                    _ = v93
                }
            case 4294967295:
                if true {
                    v34 = append(v34, (v32 & v32))
                    _ = len(v34)
                } else {
                    if len(v34) > 0 {
                        v34[0] = (v35 - 128)
                    }
                }
            case 32767:
                for fi94 := 0; fi94 < 2; fi94++ {
                    v95 := GenS8{F0: 141, F1: 322, F2: ((47 != 0) || (v30 || v31))}
                    _ = fi94
                    _ = v95
                }
            default:
                for fi96 := 0; fi96 < 18; fi96++ {
                    var v97 int8 = (v3 ^ 127)
                    switch "" {
                    case "foo bar":
                        for fi98 := 0; fi98 < 6; fi98++ {
                            v31 = (3.14159 <= float32(float64(0.1)))
                            v99 := GenS8{F0: 0, F1: 1, F2: (2 == 2)}
                            var v100 float32 = (3.14159 - -2.25)
                            _ = fi98
                            _ = v100
                            _ = v99
                        }
                    default:
                        switch "hello" {
                        case "\n\t":
                            var v101 float32 = 3.14159
                            _ = v101
                        case "":
                            v28 = -317
                        default:
                            v97 = (-(v97))
                        }
                    }
                    _ = fi96
                    _ = v97
                }
            }
            for fi102 := 0; fi102 < 10; fi102++ {
                v3 += -30
                if v30 {
                    if (!v30) {
                        for fi103 := 0; fi103 < 17; fi103++ {
                            v104 := GenS9{F0: (v2 | v3), F1: true, F2: -9223372036854775808}
                            var v105 uint = 415
                            _ = fi103
                            _ = v104
                            _ = v105
                        }
                        var v106 uint = 1
                        _ = v106
                    } else {
                        switch ((("hello" + (("a" + "hello") + ("a" + "foo bar"))) + "foo bar") + "foo bar") {
                        case "":
                            var v107 uint64 = 2
                            _ = v107
                        case "\n\t":
                            if v31 {
                                v8 = (^(v8))
                                _ = len(v34)
                                for fi108 := 0; fi108 < 15; fi108++ {
                                    if len(v34) >= 2 {
                                        v34 = v34[1:len(v34)]
                                    }
                                    v1 -= (((v1 / v1) * v1) * float64((0.1 / 1.5)))
                                    if len(v34) > 0 {
                                        v34[0] = 128
                                    }
                                    v109 := GenS4{F0: 42, F1: 2}
                                    _ = fi108
                                    _ = v109
                                }
                            } else {
                                for fi110 := 0; fi110 < 9; fi110++ {
                                    if len(v34) > 0 {
                                        v34[0] = 127
                                    }
                                    var v111 int16 = -1
                                    var v112 uint8 = v8
                                    _ = fi110
                                    _ = v111
                                    _ = v112
                                }
                                v8 -= v8
                                fmt.Println(v31, v35, v31)
                            }
                        }
                    }
                    var v113 uint32 = 0
                    var v114 uint = 255
                    _ = v113
                    _ = v114
                }
                _ = fi102
            }
            v115 := GenS10{F0: "\n\t"}
            _ = v115
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v8
}
