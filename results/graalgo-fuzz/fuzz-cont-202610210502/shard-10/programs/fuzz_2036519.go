package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 int; F2 uint64 }
type GenS4 struct { F0 int8; F1 int; F2 uint }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 uint8; F1 int32; F2 bool }
type GenS7 struct { F0 int16; F1 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 string; F1 uint64; F2 uint32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int; F1 int }
type GenS12 struct { F0 uint; F1 uint32; F2 uint32 }
func genFn0(p0 uint64) string {
    v6 := []int{}
    if true {
        switch ("a" + ("\n\t" + "")) {
        case "a":
            fmt.Println(p0)
        case "hello":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            switch "a" {
            case "":
                for fi7 := 0; fi7 < 10; fi7++ {
                    var v8 uint16 = 262
                    _ = fi7
                    _ = v8
                }
            case "a":
                fmt.Println(p0)
            case "\n\t":
                for fi9 := 0; fi9 < 11; fi9++ {
                    if (1 >= 18446744073709551615) {
                        if true {
                            for fi10 := 0; fi10 < 4; fi10++ {
                                switch (fi10 << ((485) & 63)) {
                                case 32767:
                                    p0 += 878
                                default:
                                    if true {
                                        p0 = p0
                                        switch "foo bar" {
                                        case "":
                                            var v11 uint32 = 104
                                            _ = v11
                                        default:
                                            var v12 int32 = -1
                                            _ = v12
                                        }
                                    }
                                }
                                v13 := GenS1{F0: 0, F1: -1}
                                if false {
                                    var v14 uint16 = 95
                                    var v15 uint8 = 231
                                    _ = v14
                                    _ = v15
                                } else {
                                    v16 := GenS1{F0: 65535, F1: -555}
                                    v17 := GenS2{F0: 442}
                                    var v18 float32 = 1.5
                                    if (!true) {
                                        var v19 int64 = 42
                                        v19 += v19
                                        if (fi9 <= fi10) {
                                            v6 = append(v6, -2147483648)
                                        }
                                        _ = v19
                                    }
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                }
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                p0 += 0
                                _ = fi10
                                _ = v13
                            }
                            fmt.Println(fi9, p0)
                        } else {
                            p0 += p0
                        }
                        switch "a" {
                        case "":
                            var v20 int64 = 0
                            _ = v20
                        default:
                            var v21 string = "a"
                            _ = v21
                        }
                    } else {
                        v22 := GenS3{F0: 32767, F1: fi9, F2: p0}
                        for fi23 := 0; fi23 < 16; fi23++ {
                            var v24 int = (fi23 * 128)
                            _ = fi23
                            _ = v24
                        }
                        _ = v22
                    }
                    if len(v6) > 0 {
                        v6[0] = fi9
                    }
                    v25 := GenS3{F0: -242, F1: (^(fi9)), F2: 692}
                    _ = fi9
                    _ = v25
                }
            }
        }
    }
    if len(v6) > 0 {
        v6[0] = 127
    }
    _ = p0
    _ = v6
    return "a"
}
func genFn1(p0 bool) int64 {
    if p0 {
        v26 := []int{4294967295}
        if (!(-73 > 127)) {
            for fi27 := 0; fi27 < 3; fi27++ {
                var v28 uint = 42
                for fi29 := 0; fi29 < 4; fi29++ {
                    var v30 bool = (-9223372036854775808 < 0)
                    if len(v26) > 0 {
                        v26[0] = (fi29 & fi29)
                    }
                    _ = fi29
                    _ = v30
                }
                var v31 float64 = float64(1e10)
                _ = fi27
                _ = v28
                _ = v31
            }
            if false {
                var v32 int = -875
                if p0 {
                    var v33 uint32 = 183
                    v34 := GenS2{F0: -9223372036854775808}
                    _ = v33
                    _ = v34
                } else {
                    fmt.Println(v32)
                    v32 -= (v32 - (v32 + v32))
                    var v35 float64 = 0.0
                    _ = v35
                }
                _ = v32
            } else {
                if p0 {
                    v36 := GenS4{F0: 127, F1: 127, F2: 0}
                    var v37 uint = 0
                    _ = v36
                    _ = v37
                }
            }
            p0 = (32768 != -383)
        } else {
            for fi38 := 0; fi38 < 15; fi38++ {
                fmt.Println(fi38, fi38)
                p0 = (42 <= 1)
                _ = fi38
            }
            v26 = append(v26, -769)
            p0 = (127 > 28)
        }
        for fi39 := 0; fi39 < 8; fi39++ {
            var v40 float32 = (1.5 / 1.5)
            var v41 int = 2147483647
            p0 = (!(255 == 364))
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = v26
    }
    v42 := GenS1{F0: 4294967295, F1: 32767}
    _ = p0
    _ = v42
    return 471
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    v4 := []string{("foo bar" + ("a" + ("" + ("hello" + "foo bar")))), ("hello" + "foo bar")}
    var v5 int = 359
    fmt.Println(v1, v1)
    var v43 int8 = 88
    v43 -= v43
    var v44 uint16 = 2
    v1 += 146
    var v45 float32 = 3.14159
    delete(v3, v5)
    if true {
        v1 = (^(v1))
        if (!(("foo bar" >= ("\n\t" + "a")) || false)) {
            for fi46 := 0; fi46 < 13; fi46++ {
                if ("a" != ("foo bar" + "\n\t")) {
                    _ = len(v4)
                    for fi47 := 0; fi47 < 1; fi47++ {
                        var v48 bool = false
                        switch (fi46 + v5) {
                        case 969:
                            v49 := GenS2{F0: 0}
                            _ = v49
                        case 4294967295:
                            var v50 int64 = genFn1(v48)
                            _ = v50
                        }
                        if len(v4) > 0 {
                            v4[0] = "a"
                        }
                        _ = fi47
                        _ = v48
                    }
                    fmt.Println(v43)
                }
                _ = fi46
            }
            if ((true && (32767 != v44)) && ((148 > 0) || (981 > 1))) {
                v5 -= v5
            }
            var v51 uint8 = v1
            _ = v51
        }
        v1 = v1
    } else {
        v45 += (float32(float64(0.1)) / v45)
    }
    if ((v43 & v43) >= v43) {
        v1 = (v1 + v1)
        var v52 int32 = 719
        _ = v52
    } else {
        var v53 int32 = -1
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v53
    }
    if (false && true) {
        v43 += v43
        switch -1 {
        case -9223372036854775808:
            var v54 int = (v5 % ((4294967295) | 1))
            _ = v54
        case 462:
            switch "" {
            case "foo bar":
                var v55 uint64 = 1
                _ = v55
            case "\n\t":
                if true {
                    for fi56 := 0; fi56 < 15; fi56++ {
                        var v57 uint = 1
                        _ = fi56
                        _ = v57
                    }
                    switch (^(v5)) {
                    case -2147483648:
                        v58 := GenS5{F0: 257, F1: (v43 <= (v43 & 100))}
                        _ = v58
                    case 32768:
                        if true {
                            var v59 int64 = -552
                            v44 = v44
                            _ = v59
                        } else {
                            var v60 uint16 = 42
                            fmt.Println(v1, v5, v5)
                            var v61 int64 = -286
                            _ = v60
                            _ = v61
                        }
                    default:
                        switch "a" {
                        case "a":
                            v62 := GenS6{F0: v1, F1: -1, F2: (("" + "") < ("" + (("hello" + "a") + ("\n\t" + "foo bar"))))}
                            _ = v62
                        }
                    }
                    for fi63 := 0; fi63 < 10; fi63++ {
                        if false {
                            v64 := GenS6{F0: (^(v1)), F1: 2147483647, F2: (!(!((true && true) && ((float64(((v45 + (-2.25 + (1e10 - -2.25))) - (1.5 / (v45 * v45)))) * ((((1.5 * 0.0) * (0.0 * 1.5)) + 0.1) + ((((0.0 - float64(v45)) / (0.1 / 1.0)) - float64((v45 * (float32(0.1) / v45)))) / 1.5))) != 3.14159))))}
                            _ = v64
                        }
                        v65 := GenS7{F0: -993, F1: 2}
                        for fi66 := 0; fi66 < 11; fi66++ {
                            v45 = (-2.25 + v45)
                            switch v5 {
                            case -793:
                                v44 = v44
                            case 42:
                                var v67 string = (("hello" + "\n\t") + "")
                                _ = v67
                            default:
                                fmt.Println(v44)
                            }
                            _ = fi66
                        }
                        _ = fi63
                        _ = v65
                    }
                }
            }
        default:
            for fi68 := 0; fi68 < 12; fi68++ {
                if false {
                    for fi69 := 0; fi69 < 4; fi69++ {
                        v3[(fi68 >> ((1) & 63))] = fi68
                        var v70 uint32 = 2
                        _ = len(v3)
                        var v71 bool = true
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                    var v72 uint8 = 255
                    switch (v5 * (fi68 | (v5 / ((v5) | 1)))) {
                    case -2147483648:
                        fmt.Println(v45, fi68)
                    case 818:
                        var v73 uint16 = (^(v44))
                        _ = v73
                    default:
                        v74 := GenS5{F0: 841, F1: (0 < 65535)}
                        _ = v74
                    }
                    _ = v72
                }
                for fi75 := 0; fi75 < 8; fi75++ {
                    fmt.Println(v45)
                    v1 = v1
                    switch ("a" + "foo bar") {
                    case "hello":
                        if false {
                            fmt.Println(fi75, fi68, fi75)
                        }
                    case "a":
                        var v76 bool = true
                        _ = v76
                    case "\n\t":
                        var v77 uint64 = 4294967295
                        _ = v77
                    }
                    _ = fi75
                }
                switch "foo bar" {
                case "foo bar":
                    var v78 float32 = v45
                    _ = v78
                case "\n\t":
                    v79 := GenS8{F0: v43}
                    _ = v79
                default:
                    var v80 uint = 355
                    _ = v80
                }
                _ = fi68
            }
        }
        v81 := v2["foo bar"]
        _ = v81
    } else {
        v82 := GenS9{F0: "hello", F1: 455, F2: 42}
        for fi83 := 0; fi83 < 10; fi83++ {
            _ = len(v2)
            _ = fi83
        }
        var v84 bool = (true || (172 >= 0))
        _ = v82
        _ = v84
    }
    if (true || (2 > 2)) {
        v85 := GenS10{F0: (v1 | (v1 ^ 127))}
        if (((0.1 / -2.25) + ((((0.1 / (0.1 / 1e10)) + -2.25) * ((float64(v45) / -2.25) * (0.0 - (0.0 / (1e10 / 0.1))))) + 0.1)) <= float64(v45)) {
            v86 := v3[9223372036854775807]
            var v87 int8 = v43
            _ = v86
            _ = v87
        } else {
            v44 = 42
            if ((!true) && (-1 < -1)) {
                v43 += v43
            }
        }
        _ = v85
    } else {
        _ = len(v4)
        var v88 uint8 = v1
        var v89 float64 = 1.5
        _ = v88
        _ = v89
    }
    switch (v5 | -2147483648) {
    case 0:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case -767:
        v90 := GenS11{F0: (v5 - (v5 * (v5 | v5))), F1: v5}
        _ = v90
    case 424:
        var v91 uint32 = 1
        _ = v91
    default:
        if (1000 <= -107) {
            _ = len(v2)
        } else {
            var v92 uint16 = v44
            _ = v92
        }
    }
    if ((v5 / (((v5 * v5)) | 1)) > -2147483648) {
        v93 := GenS12{F0: 842, F1: 943, F2: 213}
        _ = v93
    }
    switch 1 {
    case 32768:
        if (!(0 == (v1 * uint8(v44)))) {
            if true {
                fmt.Println(v5, v5, v45)
            } else {
                v94 := GenS4{F0: 42, F1: (v5 % (((v5 & (-(v5)))) | 1)), F2: 374}
                _ = v94
            }
        } else {
            var v95 bool = true
            _ = v95
        }
    }
    if (true && false) {
        v1 = v1
        for fi96 := 0; fi96 < 6; fi96++ {
            if (2147483647 >= 1) {
                var v97 float32 = (v45 + v45)
                v43 += (v43 >> ((uint(v44)) & 63))
                v45 -= v45
                _ = v97
            } else {
                v1 = v1
                v98 := GenS11{F0: fi96, F1: 128}
                _ = v98
            }
            _ = fi96
        }
    } else {
        v99 := GenS12{F0: 4294967295, F1: uint32(v1), F2: 2147483647}
        var v100 float64 = -2.25
        _ = v100
        _ = v99
    }
    v101 := v2[("foo bar" + (("\n\t" + "") + ""))]
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
