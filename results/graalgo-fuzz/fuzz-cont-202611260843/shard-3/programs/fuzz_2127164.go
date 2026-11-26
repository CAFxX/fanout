package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 uint16; F1 uint32 }
type GenS8 struct { F0 uint16; F1 int64 }
type GenS9 struct { F0 string }
func genFn0(p0 int) int64 {
    v7 := GenS1{F0: 1}
    v8 := GenS1{F0: 814}
    var v9 string = (("a" + (("\n\t" + "foo bar") + "")) + (("hello" + (("foo bar" + "hello") + "foo bar")) + "\n\t"))
    for fi10 := 0; fi10 < 20; fi10++ {
        p0 = fi10
        for fi11 := 0; fi11 < 17; fi11++ {
            if (-2.25 > 0.0) {
                switch (v9 + "hello") {
                case "foo bar":
                    v12 := []int{fi10}
                    _ = v12
                case "":
                    v9 = "\n\t"
                case "hello":
                    if true {
                        var v13 uint16 = 42
                        _ = v13
                    } else {
                        v14 := GenS2{F0: 355, F1: 369, F2: "hello"}
                        v15 := GenS3{F0: 0}
                        _ = v14
                        _ = v15
                    }
                }
            }
            _ = fi11
        }
        var v16 bool = false
        _ = fi10
        _ = v16
    }
    if false {
        var v17 uint32 = 42
        fmt.Println(v17)
        _ = v17
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint16) int8 {
    for fi18 := 0; fi18 < 18; fi18++ {
        v19 := []int{fi18, fi18}
        p0 -= (p0 << ((42) & 63))
        _ = fi18
        _ = v19
    }
    p0 = p0
    _ = p0
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 float64 = 1e10
    var v5 float64 = 1.5
    var v6 uint = 1
    v5 = (3.14159 / -2.25)
    v6 = 4294967295
    for fi20 := 0; fi20 < 3; fi20++ {
        switch fi20 {
        case 2147483647:
            v21 := []int{1}
            _ = v21
        case 28:
            var v22 uint64 = 1
            _ = v22
        default:
            v4 += (3.14159 * 0.0)
        }
        v23 := GenS3{F0: 1}
        _ = fi20
        _ = v23
    }
    if false {
        if true {
            switch 128 {
            case 1:
                for fi24 := 0; fi24 < 15; fi24++ {
                    v25 := GenS4{F0: 426}
                    var v26 int8 = -77
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            case 746:
                switch "a" {
                case "":
                    v27 := GenS2{F0: 4294967295, F1: 1, F2: ("hello" + "foo bar")}
                    _ = v27
                }
            case -9223372036854775808:
                fmt.Println(v4, v4)
            default:
                v28 := GenS5{F0: 9223372036854775807}
                _ = v28
            }
            v5 = v5
        }
        for fi29 := 0; fi29 < 19; fi29++ {
            var v30 float64 = (v5 * ((((v4 / v5) / (((((1.5 / v4) * v5) - 1.5) / v4) / v5)) * 0.1) - (0.1 + v4)))
            _ = fi29
            _ = v30
        }
        var v31 uint = v6
        _ = v31
    }
    var v32 uint16 = 65535
    switch ("foo bar" + "a") {
    case "":
        switch "a" {
        case "hello":
            v33 := []int{-162, 32767}
            _ = v33
        case "":
            var v34 float64 = float64(-2.25)
            _ = v34
        case "a":
            if true {
                v35 := []int{-706, 196, 518}
                _ = v35
            }
        default:
            var v36 uint32 = 8
            _ = v36
        }
    case "foo bar":
        v5 += v5
    }
    if true {
        v37 := GenS2{F0: 42, F1: 401, F2: "a"}
        var v38 int = 477
        var v39 uint16 = v32
        _ = v37
        _ = v38
        _ = v39
    } else {
        if (546 <= 42) {
            var v40 bool = (false && (false && false))
            var v41 int8 = 110
            if (v40 || ((v6 + (^(v6))) < 612)) {
                if (42 >= -32768) {
                    v42 := v1[((("" + "\n\t") + "\n\t") + ("a" + ""))]
                    _ = v42
                } else {
                    v43 := []int{-1}
                    if v40 {
                        if v40 {
                            if v40 {
                                v41 += v41
                            } else {
                                v41 -= -39
                                if len(v43) >= 2 {
                                    v43 = v43[1:len(v43)]
                                }
                                var v44 uint32 = 605
                                _ = v44
                            }
                            for fi45 := 0; fi45 < 7; fi45++ {
                                switch "" {
                                case "":
                                    var v46 uint = (v6 >> ((457) & 63))
                                    _ = v46
                                case "foo bar":
                                    fmt.Println(v6, v5)
                                case "a":
                                    switch "hello" {
                                    case "foo bar":
                                        var v47 int8 = genFn1(v32)
                                        _ = v47
                                    default:
                                        v4 += v5
                                    }
                                default:
                                    for fi48 := 0; fi48 < 1; fi48++ {
                                        v49 := GenS3{F0: v41}
                                        if (false && (0.1 >= 1.5)) {
                                            v41 = 127
                                            var v50 float64 = float64(0.0)
                                            _ = v50
                                        }
                                        if (!v40) {
                                            if ((v6 != v6) || true) {
                                                if (uint64(v6) >= 0) {
                                                    var v51 uint8 = 32
                                                    _ = v51
                                                }
                                                v6 = v6
                                            } else {
                                                var v52 int32 = 2147483647
                                                if (v6 == (v6 << (((^(v6))) & 63))) {
                                                    var v53 int16 = 1000
                                                    if (!v40) {
                                                        _ = len(v1)
                                                    } else {
                                                        v54 := GenS6{F0: int32(fi45), F1: fi45}
                                                        var v55 int = fi48
                                                        v56 := GenS5{F0: 589}
                                                        _ = len(v43)
                                                        _ = v54
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                    switch fi45 {
                                                    case 32768:
                                                        fmt.Println(v41, v6)
                                                    case -2147483648:
                                                        var v57 int16 = v53
                                                        _ = v57
                                                    case 114:
                                                        switch "" {
                                                        case "\n\t":
                                                            var v58 uint32 = 2147483648
                                                            _ = v58
                                                        case "":
                                                            v41 += v41
                                                        case "hello":
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                var v59 uint = (v6 | v6)
                                                                _ = v59
                                                            case "a":
                                                                switch (fi48 % ((fi48) | 1)) {
                                                                case 695:
                                                                    if false {
                                                                        var v60 float32 = (3.14159 / 0.1)
                                                                        v52 = -2147483648
                                                                        _ = v60
                                                                    } else {
                                                                        switch "a" {
                                                                        case "a":
                                                                            fmt.Println(v53)
                                                                        default:
                                                                            switch ("" + "a") {
                                                                            case "a":
                                                                                switch "\n\t" {
                                                                                case "hello":
                                                                                    v61 := GenS4{F0: 4294967295}
                                                                                    _ = v61
                                                                                default:
                                                                                    var v62 uint64 = 42
                                                                                    _ = v62
                                                                                }
                                                                            default:
                                                                                var v63 bool = (180 != 42)
                                                                                _ = v63
                                                                            }
                                                                        }
                                                                        if (fi48 == 1) {
                                                                            if v40 {
                                                                                fmt.Println(v5)
                                                                            } else {
                                                                                fmt.Println(v53, v6)
                                                                                var v64 string = "a"
                                                                                _ = len(v1)
                                                                                _ = v64
                                                                            }
                                                                            if len(v43) > 0 {
                                                                                v43[0] = (fi48 / (((fi48 % ((fi48) | 1))) | 1))
                                                                            }
                                                                            v65 := GenS5{F0: 1}
                                                                            if false {
                                                                                v66 := GenS6{F0: (v52 - v52), F1: (fi45 ^ fi45)}
                                                                                var v67 int32 = v52
                                                                                _ = v66
                                                                                _ = v67
                                                                            }
                                                                            switch (fi45 >> ((v6) & 63)) {
                                                                            case -1:
                                                                                v6 = (v6 ^ v6)
                                                                            }
                                                                            _ = v65
                                                                        } else {
                                                                            var v68 uint = v6
                                                                            _ = v68
                                                                        }
                                                                    }
                                                                default:
                                                                    if len(v43) > 0 {
                                                                        v43[0] = (fi45 + fi48)
                                                                    }
                                                                }
                                                            default:
                                                                var v69 uint32 = 2147483648
                                                                _ = v69
                                                            }
                                                        }
                                                    default:
                                                        if len(v43) >= 2 {
                                                            v43 = v43[1:len(v43)]
                                                        }
                                                    }
                                                    _ = v53
                                                }
                                                var v70 float32 = 0.0
                                                var v71 int8 = (v41 % (((^(v41))) | 1))
                                                v72 := GenS7{F0: v32, F1: 1}
                                                _ = v52
                                                _ = v70
                                                _ = v71
                                                _ = v72
                                            }
                                            fmt.Println(fi45, v5, v5)
                                        } else {
                                            v6 += v6
                                            fmt.Println(fi48, v5, v6)
                                        }
                                        _ = fi48
                                        _ = v49
                                    }
                                }
                                for fi73 := 0; fi73 < 2; fi73++ {
                                    fmt.Println(v41, v6, v5)
                                    v40 = true
                                    var v74 int32 = -30
                                    v75 := GenS8{F0: (v32 ^ (v32 + v32)), F1: -508}
                                    v76 := GenS2{F0: 42, F1: 0, F2: "\n\t"}
                                    _ = fi73
                                    _ = v74
                                    _ = v75
                                    _ = v76
                                }
                                _ = fi45
                            }
                        }
                        switch "foo bar" {
                        case "a":
                            for fi77 := 0; fi77 < 15; fi77++ {
                                var v78 uint64 = 42
                                v78 += v78
                                _ = fi77
                                _ = v78
                            }
                        case "foo bar":
                            if (!(!v40)) {
                                var v79 bool = true
                                var v80 uint32 = 1
                                _ = v79
                                _ = v80
                            } else {
                                for fi81 := 0; fi81 < 20; fi81++ {
                                    switch "a" {
                                    case "":
                                        var v82 bool = v40
                                        _ = v82
                                    case "\n\t":
                                        if len(v43) >= 2 {
                                            v43 = v43[1:len(v43)]
                                        }
                                    case "foo bar":
                                        if (241 < 18446744073709551615) {
                                            var v83 int8 = v41
                                            if v40 {
                                                v6 = v6
                                            } else {
                                                for fi84 := 0; fi84 < 18; fi84++ {
                                                    var v85 uint8 = 1
                                                    _ = fi84
                                                    _ = v85
                                                }
                                            }
                                            v86 := GenS8{F0: (v32 + v32), F1: 0}
                                            _ = v83
                                            _ = v86
                                        }
                                    default:
                                        var v87 int8 = (v41 ^ (v41 << ((18446744073709551615) & 63)))
                                        _ = v87
                                    }
                                    switch (fi81 * 4294967295) {
                                    case 380:
                                        for fi88 := 0; fi88 < 4; fi88++ {
                                            fmt.Println(v6, v6, v41)
                                            _ = fi88
                                        }
                                    default:
                                        if v40 {
                                            if len(v43) > 0 {
                                                v43[0] = 127
                                            }
                                            switch "" {
                                            case "":
                                                switch "" {
                                                case "foo bar":
                                                    var v89 uint8 = 128
                                                    _ = v89
                                                }
                                            case "\n\t":
                                                var v90 int8 = -45
                                                _ = v90
                                            case "hello":
                                                v6 -= (v6 << ((v6) & 63))
                                            default:
                                                v5 -= -2.25
                                            }
                                        } else {
                                            switch -1 {
                                            case 2147483647:
                                                if ((float32((((v5 + ((v5 / 1.0) * v4)) * v4) * ((0.1 * (((v5 / (-2.25 - (float64(((0.0 - ((0.1 / (-2.25 - (-2.25 / 1.5))) * 1e10)) * float32(3.14159))) + v4))) * (1e10 * 1.5)) * v5)) + v5))) < (1e10 * 3.14159)) && v40) {
                                                    for fi91 := 0; fi91 < 20; fi91++ {
                                                        v3[65535] = 747
                                                        _ = fi91
                                                    }
                                                    var v92 uint64 = 41
                                                    if v40 {
                                                        if len(v43) > 0 {
                                                            v43[0] = (fi81 & fi81)
                                                        }
                                                        if v40 {
                                                            v40 = (((-2147483648 != 2147483647) && ((1 > 128) && (2147483647 == 722))) || v40)
                                                        }
                                                        var v93 int32 = 827
                                                        _ = v93
                                                    }
                                                    _ = v92
                                                }
                                            }
                                            v94 := GenS9{F0: "a"}
                                            fmt.Println(v40)
                                            _ = v94
                                        }
                                    }
                                    var v95 int8 = v41
                                    _ = fi81
                                    _ = v95
                                }
                                var v96 uint64 = 4294967295
                                var v97 uint32 = 4294967295
                                _ = v96
                                _ = v97
                            }
                        }
                    }
                    fmt.Println(v32)
                    _ = v43
                }
                v98 := []int{42, -9223372036854775808}
                v99 := GenS4{F0: 2}
                _ = v98
                _ = v99
            } else {
                delete(v2, -9223372036854775808)
                if v40 {
                    switch 32768 {
                    case 9223372036854775807:
                        v100 := GenS6{F0: -2147483648, F1: 4294967295}
                        _ = v100
                    case 0:
                        v101 := v3[167]
                        _ = v101
                    case 127:
                        if v40 {
                            v4 -= v4
                            if (v40 && (0 == 313)) {
                                var v102 string = "foo bar"
                                _ = v102
                            } else {
                                fmt.Println(v40, v32, v41)
                            }
                        } else {
                            var v103 int = 2147483647
                            for fi104 := 0; fi104 < 7; fi104++ {
                                v105 := []int{fi104, (fi104 / (((v103 / ((v103) | 1))) | 1)), v103}
                                var v106 float64 = v4
                                var v107 uint = 74
                                _ = fi104
                                _ = v105
                                _ = v106
                                _ = v107
                            }
                            fmt.Println(v41, v4, v40)
                            _ = v103
                        }
                    }
                }
                switch 4294967295 {
                case 128:
                    for fi108 := 0; fi108 < 12; fi108++ {
                        fmt.Println(v4)
                        _ = fi108
                    }
                case 377:
                    for fi109 := 0; fi109 < 11; fi109++ {
                        v110 := GenS4{F0: 2}
                        var v111 uint32 = 563
                        _ = fi109
                        _ = v110
                        _ = v111
                    }
                case 95:
                    v40 = v40
                }
            }
            _ = v40
            _ = v41
        } else {
            fmt.Println(v4, v32, v4)
        }
        var v112 float64 = (v4 * float64(3.14159))
        _ = v112
    }
    var v113 int = -161
    v114 := []int{v113, v113}
    v115 := GenS4{F0: 18446744073709551615}
    v116 := GenS7{F0: (^(v32)), F1: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v113:", v113)
    fmt.Println("slicelen v114:", len(v114), v114)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v113
    _ = v114
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
