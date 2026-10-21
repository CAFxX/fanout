package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int32; F1 uint32; F2 int8 }
type GenS6 struct { F0 string; F1 int64; F2 uint32 }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 uint8; F1 string; F2 int8 }
type GenS9 struct { F0 int; F1 int16; F2 int16 }
type GenS10 struct { F0 int; F1 uint16 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint64; F1 uint }
type GenS13 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 int8; F1 int8; F2 uint64 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 int; F1 uint32; F2 int16 }
func genFn0(p0 int16, p1 uint16) uint16 {
    v6 := GenS1{F0: 255}
    fmt.Println(p1, p0)
    p0 = (p0 / ((p0) | 1))
    _ = p0
    _ = p1
    _ = v6
    return 2
}

func main() {
    var v1 uint = 694
    v2 := make(map[uint32]bool)
    v3 := []uint32{}
    var v4 uint16 = 1
    v5 := []uint32{}
    var v7 int64 = -9223372036854775808
    switch "" {
    case "":
        if (-1 < 626) {
            v8 := GenS2{F0: 299, F1: -1, F2: (v4 / (((v4 & v4)) | 1))}
            switch "foo bar" {
            case "foo bar":
                var v9 bool = (606 <= 1)
                _ = v9
            case "":
                if (1e10 >= 0.1) {
                    v10 := GenS1{F0: v1}
                    v11 := GenS3{F0: 1}
                    _ = v10
                    _ = v11
                }
            case "a":
                v1 = v1
            }
            var v12 int16 = 0
            _ = v12
            _ = v8
        }
    default:
        var v13 float32 = 0.0
        _ = v13
    }
    var v14 int8 = 42
    switch ("a" + "hello") {
    case "a":
        if (false || false) {
            v5 = append(v5, 674)
            var v15 uint16 = (v4 & v4)
            _ = v15
        } else {
            switch 127 {
            case -9223372036854775808:
                if ((v4 ^ 637) == v4) {
                    _ = len(v5)
                    if (!false) {
                        var v16 int32 = -337
                        var v17 int = 128
                        fmt.Println(v1, v7, v14)
                        _ = v16
                        _ = v17
                    }
                    var v18 int16 = 32767
                    _ = v18
                } else {
                    v14 = v14
                }
            case -917:
                switch "foo bar" {
                case "hello":
                    v14 = (-(v14))
                case "":
                    fmt.Println(v14)
                case "foo bar":
                    if (v14 == v14) {
                        var v19 uint16 = 1
                        _ = v19
                    }
                }
            default:
                v20 := GenS1{F0: v1}
                _ = v20
            }
        }
    case "\n\t":
        v7 = v7
    default:
        v1 -= (v1 & v1)
    }
    var v21 bool = false
    switch ("foo bar" + "a") {
    case "foo bar":
        var v22 float32 = -2.25
        _ = v22
    default:
        if (1 < 1) {
            v7 = v7
            switch -842 {
            case -2147483648:
                v1 += uint(v4)
            case 1:
                v1 = (^(v1))
            case 127:
                var v23 uint = v1
                _ = v23
            default:
                v1 -= (v1 | (v1 & v1))
            }
            if v21 {
                var v24 uint8 = 128
                _ = v24
            } else {
                v14 -= -1
                for fi25 := 0; fi25 < 10; fi25++ {
                    for fi26 := 0; fi26 < 13; fi26++ {
                        if (-585 != -1) {
                            if v21 {
                                v27 := GenS4{F0: genFn0(42, (v4 << ((v1) & 63))), F1: 1}
                                _ = v27
                            }
                            delete(v2, 65535)
                            var v28 uint64 = 990
                            _ = v28
                        } else {
                            v29 := GenS1{F0: v1}
                            switch (fi26 - (fi25 ^ fi25)) {
                            case 9223372036854775807:
                                fmt.Println(v4)
                            case 32767:
                                if true {
                                    var v30 int32 = 0
                                    _ = v30
                                } else {
                                    switch "foo bar" {
                                    case "a":
                                        switch ("\n\t" + ("hello" + "")) {
                                        case "hello":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        default:
                                            v31 := GenS4{F0: (v4 & (v4 - (v4 / (((v4 / ((v4) | 1))) | 1)))), F1: 1}
                                            _ = v31
                                        }
                                    case "":
                                        v32 := v2[4294967295]
                                        _ = v32
                                    }
                                    var v33 uint16 = v4
                                    _ = v33
                                }
                            case -348:
                                v21 = true
                            }
                            _ = v29
                        }
                        switch "\n\t" {
                        case "":
                            var v34 int16 = 1000
                            _ = v34
                        case "\n\t":
                            if (v7 < (v7 * v7)) {
                                _ = len(v3)
                                var v35 uint = v1
                                _ = v35
                            } else {
                                v36 := GenS5{F0: -1, F1: 2147483647, F2: 42}
                                var v37 int32 = 0
                                var v38 bool = v21
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                        }
                        var v39 uint = (^(v1))
                        _ = fi26
                        _ = v39
                    }
                    _ = fi25
                }
                for fi40 := 0; fi40 < 16; fi40++ {
                    for fi41 := 0; fi41 < 3; fi41++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v42 bool = (2147483648 > 4294967295)
                        _ = fi41
                        _ = v42
                    }
                    for fi43 := 0; fi43 < 9; fi43++ {
                        v14 += -5
                        if v21 {
                            v14 += v14
                            _ = len(v2)
                        } else {
                            switch (^(fi43)) {
                            case 999:
                                var v44 int = (fi43 >> ((v1) & 63))
                                _ = v44
                            case 264:
                                v7 = -9223372036854775808
                            case 0:
                                v45 := GenS6{F0: "", F1: v7, F2: 42}
                                _ = v45
                            }
                            fmt.Println(v21, v1)
                            fmt.Println(v7, v21, v21)
                        }
                        v46 := GenS1{F0: 255}
                        switch (fi43 & (fi43 / (((fi43 | (fi43 << ((0) & 63)))) | 1))) {
                        case -2147483648:
                            var v47 string = "hello"
                            _ = v47
                        }
                        v14 += 95
                        var v48 int8 = 100
                        delete(v2, 485)
                        _ = fi43
                        _ = v46
                        _ = v48
                    }
                    if v21 {
                        if true {
                            fmt.Println(v21)
                            for fi49 := 0; fi49 < 7; fi49++ {
                                v50 := GenS7{F0: (-(v14)), F1: uint64(v4)}
                                _ = fi49
                                _ = v50
                            }
                        } else {
                            fmt.Println(fi40, v14, v7)
                            if v21 {
                                var v51 uint = (^(v1))
                                switch ("foo bar" + ("foo bar" + "")) {
                                case "foo bar":
                                    var v52 uint64 = 18446744073709551615
                                    _ = v52
                                case "a":
                                    var v53 float64 = (((((1e10 * 3.14159) - 3.14159) - 1.5) * 3.14159) - 0.0)
                                    _ = v53
                                case "\n\t":
                                    fmt.Println(v1, v21)
                                }
                                v54 := GenS8{F0: 15, F1: "", F2: v14}
                                _ = v51
                                _ = v54
                            } else {
                                v55 := GenS4{F0: v4, F1: 127}
                                _ = v55
                            }
                            v56 := GenS9{F0: fi40, F1: 0, F2: 1000}
                            _ = v56
                        }
                    } else {
                        for fi57 := 0; fi57 < 16; fi57++ {
                            v58 := GenS10{F0: (fi40 % ((fi40) | 1)), F1: v4}
                            fmt.Println(fi57, v14, fi40)
                            v7 -= (v7 % (((^(v7))) | 1))
                            _ = fi57
                            _ = v58
                        }
                        v59 := GenS10{F0: fi40, F1: v4}
                        _ = v59
                    }
                    _ = fi40
                }
            }
        } else {
            if (v4 >= (v4 % ((genFn0(-32768, (v4 | genFn0(32767, v4)))) | 1))) {
                v4 -= v4
                v4 = (v4 >> (((^(v1))) & 63))
            } else {
                for fi60 := 0; fi60 < 20; fi60++ {
                    var v61 float64 = 3.14159
                    v7 = 0
                    _ = fi60
                    _ = v61
                }
                v1 = 1
                if false {
                    if v21 {
                        for fi62 := 0; fi62 < 20; fi62++ {
                            var v63 int8 = v14
                            if (fi62 == (fi62 - fi62)) {
                                switch fi62 {
                                case 0:
                                    if true {
                                        var v64 uint64 = 795
                                        if (255 >= 42) {
                                            var v65 bool = (!(!((v1 ^ (^(v1))) <= (v1 >> (((^(v1))) & 63)))))
                                            _ = v65
                                        } else {
                                            var v66 int16 = 1
                                            for fi67 := 0; fi67 < 1; fi67++ {
                                                var v68 int = 128
                                                fmt.Println(v63, v68)
                                                v69 := GenS11{F0: "a"}
                                                _ = fi67
                                                _ = v68
                                                _ = v69
                                            }
                                            switch 0 {
                                            case 4294967295:
                                                v21 = v21
                                            case 2147483647:
                                                v1 = (v1 / ((255) | 1))
                                            case 128:
                                                switch "foo bar" {
                                                case "hello":
                                                    if false {
                                                        v66 -= v66
                                                        var v70 int8 = 127
                                                        _ = v70
                                                    }
                                                }
                                            }
                                            _ = v66
                                        }
                                        v71 := GenS9{F0: 4294967295, F1: int16(v7), F2: -1}
                                        _ = v64
                                        _ = v71
                                    }
                                case 418:
                                    v4 += v4
                                case -2147483648:
                                    v63 += v63
                                }
                                v21 = v21
                            }
                            v7 += v7
                            _ = fi62
                            _ = v63
                        }
                    } else {
                        var v72 int = 127
                        v73 := GenS12{F0: 2, F1: 382}
                        _ = v72
                        _ = v73
                    }
                } else {
                    if ((!v21) || (((v1 - 18446744073709551615) <= 1) && true)) {
                        v14 = 108
                        switch "" {
                        case "\n\t":
                            if v21 {
                                fmt.Println(v1, v1)
                                if v21 {
                                    if (1 < int8(v4)) {
                                        if true {
                                            switch "hello" {
                                            case "foo bar":
                                                var v74 float32 = (((1e10 / 1.5) - (0.0 / ((-2.25 + 1.5) / (-2.25 - (0.1 / 1.5))))) - 1e10)
                                                _ = v74
                                            case "":
                                                v14 += 42
                                            }
                                            for fi75 := 0; fi75 < 15; fi75++ {
                                                v21 = (v7 > 42)
                                                _ = fi75
                                            }
                                        }
                                        v76 := GenS5{F0: int32(v4), F1: 592, F2: (v14 - v14)}
                                        _ = v76
                                    }
                                }
                            }
                        case "a":
                            if v21 {
                                v77 := GenS2{F0: 210, F1: 0, F2: (v4 / (((^(v4))) | 1))}
                                _ = v77
                            }
                        default:
                            v7 += v7
                        }
                    }
                }
            }
            _ = len(v5)
        }
    }
    var v78 bool = (!v21)
    for fi79 := 0; fi79 < 18; fi79++ {
        if true {
            fmt.Println(v78, v78, v4)
            var v80 int = 788
            v1 = 65535
            _ = v80
        } else {
            var v81 float32 = (-2.25 / 1e10)
            if true {
                fmt.Println(v7, v7)
            }
            _ = v81
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi79
    }
    var v82 int8 = 42
    for fi83 := 0; fi83 < 10; fi83++ {
        if true {
            v84 := GenS5{F0: 1, F1: 2147483648, F2: int8(v7)}
            v21 = (v21 && v21)
            _ = v84
        } else {
            v85 := GenS2{F0: 42, F1: 32767, F2: v4}
            if false {
                var v86 int8 = (v82 | (v14 - (-(v14))))
                if false {
                    if ((v21 || ((v86 != v86) || (v1 > (^(v1))))) || (18446744073709551615 < 117)) {
                        var v87 uint16 = 32767
                        for fi88 := 0; fi88 < 17; fi88++ {
                            var v89 int32 = -150
                            if (false && (!(!(false && v78)))) {
                                v2[2147483648] = v21
                                var v90 string = ("foo bar" + (("a" + ("foo bar" + (("a" + "foo bar") + "a"))) + "\n\t"))
                                v14 += 1
                                _ = v90
                            }
                            _ = fi88
                            _ = v89
                        }
                        var v91 int32 = 42
                        _ = v87
                        _ = v91
                    }
                    fmt.Println(v82, v78, v86)
                } else {
                    v14 = v14
                    for fi92 := 0; fi92 < 7; fi92++ {
                        v93 := GenS13{F0: v4, F1: 42, F2: 707}
                        _ = fi92
                        _ = v93
                    }
                    v94 := GenS14{F0: 780}
                    _ = v94
                }
                if (!false) {
                    if (((!true) || (0.1 < -2.25)) && v78) {
                        var v95 string = "foo bar"
                        _ = v95
                    } else {
                        v7 -= v7
                        delete(v2, 421)
                        fmt.Println(v1, v78, v82)
                    }
                    v96 := GenS15{F0: v82, F1: int8(v4), F2: 552}
                    v4 -= v4
                    _ = v96
                }
                _ = v86
            } else {
                switch (fi83 / ((fi83) | 1)) {
                case -2147483648:
                    v14 += v82
                case 9223372036854775807:
                    switch "foo bar" {
                    case "":
                        v97 := v2[699]
                        _ = v97
                    }
                }
                v21 = (42 >= 632)
            }
            _ = v85
        }
        var v98 uint64 = 18446744073709551615
        fmt.Println(v7, v1, v7)
        _ = fi83
        _ = v98
    }
    fmt.Println(v7, v82, v7)
    if false {
        v99 := GenS16{F0: -170}
        v78 = v78
        v100 := GenS17{F0: 316, F1: 2147483647, F2: 127}
        _ = v100
        _ = v99
    }
    var v101 int32 = 180
    v101 += v101
    var v102 int8 = (^(v82))
    if false {
        v102 = -1
        var v103 int = -1
        _ = v103
    }
    fmt.Println(v21)
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v78:", v78)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v101
    _ = v102
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v78
    _ = v82
}
