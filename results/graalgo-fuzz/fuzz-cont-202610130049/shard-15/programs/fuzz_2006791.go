package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 string; F2 uint }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 bool }
func genFn0(p0 int, p1 uint32, p2 uint8) bool {
    fmt.Println(p0, p2)
    var v6 int8 = -128
    switch "foo bar" {
    case "a":
        fmt.Println(v6)
    }
    p0 -= (p0 & p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return false
}

func main() {
    var v1 string = "a"
    v2 := make(map[int]uint32)
    v3 := []string{v1, v1, v1, "foo bar"}
    var v4 float32 = 3.14159
    v5 := make(map[int]string)
    var v7 uint32 = 2
    switch 32767 {
    case -2147483648:
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 int8 = 1
            var v10 int32 = 300
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS1{F0: 515}
            v1 = ("" + v1)
            for fi13 := 0; fi13 < 5; fi13++ {
                v1 = "foo bar"
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
    }
    v14 := GenS1{F0: 2}
    for fi15 := 0; fi15 < 16; fi15++ {
        switch v1 {
        case "a":
            for fi16 := 0; fi16 < 12; fi16++ {
                fmt.Println(v1)
                var v17 uint8 = 127
                v18 := GenS1{F0: 4294967295}
                _ = fi16
                _ = v17
                _ = v18
            }
        case "\n\t":
            v7 += 42
        default:
            v19 := GenS1{F0: 67}
            _ = v19
        }
        if false {
            v20 := GenS2{F0: true, F1: v1, F2: 4294967295}
            _ = v20
        } else {
            var v21 float64 = (1e10 / 0.1)
            for fi22 := 0; fi22 < 18; fi22++ {
                v4 += (v4 / v4)
                _ = fi22
            }
            _ = v21
        }
        v23 := GenS2{F0: (v7 == (^(v7))), F1: "hello", F2: 2}
        _ = fi15
        _ = v23
    }
    if (0 >= 518) {
        v24 := GenS1{F0: 281}
        var v25 uint = 65535
        _ = v24
        _ = v25
    } else {
        fmt.Println(v4, v1)
    }
    v7 -= 65535
    v26 := GenS1{F0: 2}
    switch "hello" {
    case "":
        fmt.Println(v7)
    case "a":
        if ((v7 % ((v7) | 1)) != v7) {
            var v27 string = "foo bar"
            fmt.Println(v7, v7, v4)
            if (false || (-1 >= 1)) {
                switch v27 {
                case "\n\t":
                    v28 := GenS3{F0: 65535, F1: -1}
                    _ = v28
                case "":
                    var v29 uint32 = 2147483647
                    _ = v29
                }
                var v30 uint = 399
                _ = v30
            }
            _ = v27
        } else {
            if len(v3) > 0 {
                v3[0] = ("hello" + ((("hello" + "a") + v1) + v1))
            }
            if (65535 <= 876) {
                v31 := v2[836]
                fmt.Println(v7)
                if false {
                    for fi32 := 0; fi32 < 20; fi32++ {
                        if len(v3) > 0 {
                            v3[0] = (("hello" + (v1 + ("\n\t" + v1))) + v1)
                        }
                        v4 = float32(1.5)
                        for fi33 := 0; fi33 < 3; fi33++ {
                            var v34 uint32 = (^(v7))
                            switch (v1 + v1) {
                            case "a":
                                v35 := GenS4{F0: 923, F1: ((!false) || (false && true))}
                                _ = v35
                            case "hello":
                                var v36 int8 = 21
                                _ = v36
                            case "\n\t":
                                switch (("a" + v1) + (v1 + (v1 + v1))) {
                                case "a":
                                    var v37 float32 = v4
                                    _ = v37
                                case "\n\t":
                                    var v38 bool = ((255 == 299) && true)
                                    _ = v38
                                default:
                                    fmt.Println(fi33, v31, fi33)
                                }
                            }
                            v39 := GenS3{F0: 2, F1: 42}
                            _ = fi33
                            _ = v34
                            _ = v39
                        }
                        _ = fi32
                    }
                    v1 = ((v1 + "foo bar") + v1)
                    v7 -= v7
                } else {
                    for fi40 := 0; fi40 < 10; fi40++ {
                        if (47 > 0) {
                            delete(v2, (-(fi40)))
                            if (false && true) {
                                if true {
                                    v41 := GenS1{F0: 478}
                                    for fi42 := 0; fi42 < 16; fi42++ {
                                        if true {
                                            fmt.Println(fi42)
                                            v7 -= 74
                                        } else {
                                            if (!true) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v43 int = (fi42 | (fi40 & (fi42 * fi40)))
                                                _ = v43
                                            } else {
                                                _ = len(v2)
                                                v44 := GenS1{F0: 0}
                                                _ = v44
                                            }
                                            v31 += (v31 >> ((255) & 63))
                                            fmt.Println(v31)
                                        }
                                        fmt.Println(v7)
                                        _ = fi42
                                    }
                                    var v45 uint16 = 889
                                    _ = v41
                                    _ = v45
                                }
                                for fi46 := 0; fi46 < 13; fi46++ {
                                    var v47 uint = 4294967295
                                    switch ((v1 + v1) + ((v1 + v1) + v1)) {
                                    case "hello":
                                        var v48 uint32 = (v7 ^ (v31 | v7))
                                        _ = v48
                                    }
                                    _ = fi46
                                    _ = v47
                                }
                            } else {
                                v1 = (v1 + v1)
                                switch fi40 {
                                case 781:
                                    var v49 bool = (!false)
                                    _ = v49
                                case 139:
                                    if false {
                                        switch ((v1 + v1) + "") {
                                        case "foo bar":
                                            v50 := GenS1{F0: 4294967295}
                                            _ = v50
                                        }
                                        v4 += (float32(1e10) / (v4 - float32((1e10 - -2.25))))
                                    }
                                default:
                                    var v51 uint64 = 4294967295
                                    _ = v51
                                }
                                switch (v1 + v1) {
                                case "a":
                                    switch fi40 {
                                    case -1:
                                        if (false || ("foo bar" <= v1)) {
                                            if (false && true) {
                                                v1 = ""
                                                var v52 uint64 = 1
                                                fmt.Println(v1, v31)
                                                _ = v52
                                            } else {
                                                v53 := GenS1{F0: 4294967295}
                                                _ = v53
                                            }
                                            switch (fi40 & fi40) {
                                            case 32768:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            default:
                                                for fi54 := 0; fi54 < 15; fi54++ {
                                                    switch (fi54 ^ (fi40 & fi54)) {
                                                    case 32767:
                                                        switch fi40 {
                                                        case 0:
                                                            v55 := GenS4{F0: -1, F1: true}
                                                            _ = v55
                                                        case -870:
                                                            _ = len(v2)
                                                        case 42:
                                                            v5[fi54] = "\n\t"
                                                        }
                                                    }
                                                    v7 = 2147483648
                                                    v56 := GenS1{F0: 4294967295}
                                                    v57 := GenS2{F0: true, F1: v1, F2: 705}
                                                    if true {
                                                        switch fi54 {
                                                        case 589:
                                                            var v58 int8 = 42
                                                            _ = v58
                                                        default:
                                                            v59 := GenS3{F0: 196, F1: 1}
                                                            _ = v59
                                                        }
                                                    }
                                                    v60 := GenS2{F0: false, F1: (v1 + v1), F2: 4294967295}
                                                    if true {
                                                        fmt.Println(fi54, v4, fi54)
                                                        fmt.Println(v31, v31, v7)
                                                        var v61 int8 = 127
                                                        _ = v61
                                                    }
                                                    _ = fi54
                                                    _ = v56
                                                    _ = v57
                                                    _ = v60
                                                }
                                            }
                                            for fi62 := 0; fi62 < 7; fi62++ {
                                                delete(v2, 42)
                                                if (!(!false)) {
                                                    var v63 bool = (true && (true || (2 == 42)))
                                                    v63 = (42 <= 145)
                                                    if (!(!v63)) {
                                                        var v64 int64 = -373
                                                        var v65 uint8 = 1
                                                        v66 := GenS5{F0: v63}
                                                        _ = v64
                                                        _ = v65
                                                        _ = v66
                                                    } else {
                                                        switch (fi40 | 889) {
                                                        case 4294967295:
                                                            fmt.Println(v4, v63, v4)
                                                        case -977:
                                                            fmt.Println(fi40)
                                                        default:
                                                            v31 += (v7 ^ v31)
                                                        }
                                                    }
                                                    v67 := GenS4{F0: -843, F1: (float64((v4 - float32(1.5))) >= (-2.25 / (3.14159 * 0.1)))}
                                                    _ = v63
                                                    _ = v67
                                                }
                                                var v68 int8 = 1
                                                _ = fi62
                                                _ = v68
                                            }
                                        } else {
                                            var v69 int32 = 42
                                            for fi70 := 0; fi70 < 1; fi70++ {
                                                v31 = v7
                                                var v71 float64 = 1e10
                                                if (true && (!(358 >= 0))) {
                                                    v7 += 475
                                                    fmt.Println(fi40, fi40, v71)
                                                    if true {
                                                        var v72 uint16 = 517
                                                        _ = v72
                                                    } else {
                                                        var v73 int = -1
                                                        var v74 uint = 2
                                                        fmt.Println(v4)
                                                        v69 = -2147483648
                                                        if len(v3) > 0 {
                                                            v3[0] = v1
                                                        }
                                                        var v75 int32 = 1
                                                        if (1 == -32768) {
                                                            var v76 float32 = 0.0
                                                            fmt.Println(v76, v76, fi40)
                                                            v69 = (v69 ^ v75)
                                                            _ = v76
                                                        } else {
                                                            v75 = v75
                                                        }
                                                        _ = v73
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                    v1 = v1
                                                }
                                                _ = fi70
                                                _ = v71
                                            }
                                            switch v1 {
                                            case "foo bar":
                                                if len(v3) > 0 {
                                                    v3[0] = v1
                                                }
                                            }
                                            _ = v69
                                        }
                                    case -542:
                                        delete(v5, fi40)
                                    case 9223372036854775807:
                                        fmt.Println(v1)
                                    default:
                                        for fi77 := 0; fi77 < 17; fi77++ {
                                            var v78 int8 = -1
                                            _ = fi77
                                            _ = v78
                                        }
                                    }
                                case "\n\t":
                                    v7 = (^(v7))
                                default:
                                    switch v1 {
                                    case "":
                                        if true {
                                            var v79 uint = 65535
                                            _ = v79
                                        } else {
                                            fmt.Println(v4, v4)
                                        }
                                    default:
                                        _ = len(v2)
                                    }
                                }
                            }
                        } else {
                            fmt.Println(v7)
                            var v80 int8 = 100
                            if (!false) {
                                var v81 bool = true
                                _ = v81
                            } else {
                                v4 += 1e10
                            }
                            _ = v80
                        }
                        if (!true) {
                            var v82 int64 = 915
                            var v83 uint32 = (^(v7))
                            _ = v82
                            _ = v83
                        }
                        _ = fi40
                    }
                    if true {
                        v3 = append(v3, "a")
                        v31 += (v7 >> ((252) & 63))
                    } else {
                        for fi84 := 0; fi84 < 13; fi84++ {
                            switch (v1 + v1) {
                            case "foo bar":
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                            case "hello":
                                if (!false) {
                                    v85 := GenS1{F0: 628}
                                    if false {
                                        v1 = v1
                                    } else {
                                        v31 = (v7 - (v31 & v7))
                                    }
                                    v86 := GenS1{F0: 1}
                                    _ = v85
                                    _ = v86
                                }
                            default:
                                var v87 int64 = -515
                                _ = v87
                            }
                            switch v1 {
                            case "a":
                                for fi88 := 0; fi88 < 3; fi88++ {
                                    _ = len(v2)
                                    v4 -= (((3.14159 * v4) * ((0.1 + (((v4 - (0.0 - v4)) + v4) / ((v4 - 3.14159) * v4))) / 0.1)) - 1e10)
                                    var v89 uint64 = 42
                                    _ = fi88
                                    _ = v89
                                }
                            default:
                                for fi90 := 0; fi90 < 4; fi90++ {
                                    v31 -= v31
                                    if (-7 < 9223372036854775807) {
                                        _ = len(v5)
                                        var v91 int64 = int64(v31)
                                        _ = v91
                                    } else {
                                        switch (^(fi90)) {
                                        case 42:
                                            var v92 int64 = 1
                                            _ = v92
                                        case 355:
                                            switch fi90 {
                                            case -602:
                                                var v93 float32 = float32((((-2.25 * 0.0) * ((1e10 + 3.14159) + 0.1)) - 0.1))
                                                _ = v93
                                            }
                                        default:
                                            switch 2147483647 {
                                            case -9223372036854775808:
                                                var v94 uint64 = 18446744073709551615
                                                _ = v94
                                            }
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = v1
                                        }
                                    }
                                    v1 = ("foo bar" + "a")
                                    _ = fi90
                                }
                            }
                            v7 = v7
                            _ = fi84
                        }
                    }
                    v95 := GenS3{F0: 1, F1: -330}
                    _ = v95
                }
                _ = v31
            } else {
                fmt.Println(v1)
            }
        }
    }
    v7 += v7
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
