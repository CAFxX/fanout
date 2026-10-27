package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint }
type GenS2 struct { F0 uint32; F1 int32; F2 uint }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint8; F1 int; F2 uint64 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 int32; F1 int32; F2 uint8 }
type GenS7 struct { F0 string; F1 uint8; F2 int }
type GenS8 struct { F0 uint16; F1 int }
type GenS9 struct { F0 uint8; F1 int32 }
type GenS10 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS14 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int32, p1 int, p2 uint32) uint16 {
    v7 := []int{p1, p1, 9223372036854775807}
    v8 := GenS1{F0: 2, F1: 809, F2: 0}
    fmt.Println(p1)
    fmt.Println(p1, p1, p1)
    fmt.Println(p0, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint8) int {
    v9 := []int{}
    fmt.Println(p0)
    var v10 bool = false
    v11 := GenS2{F0: 65535, F1: -2147483648, F2: 0}
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 127
}

func main() {
    var v1 uint = 2
    v2 := []uint32{2147483647, 2147483647, 42}
    var v3 uint8 = 73
    v4 := make(map[uint32]uint32)
    var v5 uint8 = v3
    var v6 float32 = ((0.1 + (float32(float64(0.1)) + float32(((1.5 / -2.25) - (0.0 / 0.1))))) * 1.5)
    v12 := GenS3{F0: -2147483648, F1: 42}
    var v13 int16 = -32768
    v6 += 0.1
    v14 := GenS1{F0: 4294967295, F1: 175, F2: v1}
    if false {
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 uint16 = genFn0(-2147483648, (-(fi15)), 1)
            var v17 float32 = (v6 + (1.5 / 1.5))
            var v18 int = fi15
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        v6 = v6
    } else {
        if false {
            fmt.Println(v6, v5, v5)
            fmt.Println(v3, v1, v13)
        } else {
            v5 += (v3 / ((v3) | 1))
            var v19 uint64 = 4294967295
            _ = v19
        }
    }
    switch -709 {
    case 2147483647:
        var v20 float32 = float32((1.5 / 1e10))
        _ = v20
    case 707:
        var v21 uint64 = 116
        _ = v21
    default:
        var v22 int8 = -1
        _ = v22
    }
    fmt.Println(v3)
    if false {
        var v23 float32 = ((((float32(float64(v6)) - v6) + v6) - v6) - 0.1)
        switch ("a" + ("a" + ("foo bar" + "foo bar"))) {
        case "":
            v6 = (v23 + float32(-2.25))
        case "a":
            v3 -= v3
        }
        fmt.Println(v3, v23, v1)
        _ = v23
    }
    if false {
        _ = len(v4)
        v1 = (v1 - (v1 | 255))
    } else {
        if (!false) {
            for fi24 := 0; fi24 < 20; fi24++ {
                v25 := GenS1{F0: 40, F1: 0, F2: 1}
                _ = fi24
                _ = v25
            }
            for fi26 := 0; fi26 < 2; fi26++ {
                v6 = 0.1
                _ = fi26
            }
            if true {
                for fi27 := 0; fi27 < 7; fi27++ {
                    fmt.Println(v3, v13, v6)
                    _ = fi27
                }
                switch ("hello" + "foo bar") {
                case "a":
                    fmt.Println(v6, v1, v5)
                default:
                    v5 = v3
                }
                v28 := GenS4{F0: (v3 * (v5 / ((2) | 1))), F1: 1, F2: 42}
                _ = v28
            } else {
                v1 += v1
                var v29 uint16 = 255
                _ = v29
            }
        } else {
            if (false || ((false && false) && (("foo bar" + "foo bar") >= "\n\t"))) {
                var v30 bool = (true && true)
                _ = v30
            }
            v31 := GenS5{F0: (v1 | v1), F1: "hello"}
            if (v1 <= (v1 & v1)) {
                switch int(v3) {
                case 2147483647:
                    switch -9223372036854775808 {
                    case -180:
                        switch (("a" + "foo bar") + "") {
                        case "a":
                            v32 := GenS6{F0: int32(v1), F1: 0, F2: 2}
                            _ = v32
                        case "hello":
                            for fi33 := 0; fi33 < 18; fi33++ {
                                v13 = 1
                                if false {
                                    if len(v2) > 0 {
                                        v2[0] = 4294967295
                                    }
                                    for fi34 := 0; fi34 < 15; fi34++ {
                                        v35 := GenS7{F0: (("foo bar" + "") + "hello"), F1: v5, F2: fi34}
                                        if false {
                                            v36 := GenS3{F0: 42, F1: -123}
                                            _ = v36
                                        } else {
                                            var v37 float32 = float32((3.14159 - ((0.1 + 0.0) + (((-2.25 * (0.1 * 0.0)) - -2.25) + ((-2.25 - 1.5) * 0.1)))))
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            fmt.Println(v13)
                                            fmt.Println(v13, v13, v6)
                                            if true {
                                                var v38 float64 = 3.14159
                                                _ = v38
                                            } else {
                                                var v39 uint8 = v3
                                                v13 = (v13 * v13)
                                                fmt.Println(v5, v3, v39)
                                                if ("" <= "foo bar") {
                                                    v40 := GenS8{F0: 32767, F1: (fi33 / ((genFn1((^(v5)))) | 1))}
                                                    _ = v40
                                                } else {
                                                    var v41 uint = 1
                                                    v41 -= v41
                                                    v42 := GenS9{F0: 2, F1: -1}
                                                    _ = v41
                                                    _ = v42
                                                }
                                                switch int(v5) {
                                                case -961:
                                                    fmt.Println(v6, v1, fi33)
                                                default:
                                                    var v43 int16 = 187
                                                    _ = v43
                                                }
                                                var v44 uint64 = uint64(v39)
                                                _ = v39
                                                _ = v44
                                            }
                                            _ = v37
                                        }
                                        v13 = (v13 * 658)
                                        _ = fi34
                                        _ = v35
                                    }
                                    delete(v4, 761)
                                }
                                delete(v4, 4294967295)
                                _ = fi33
                            }
                        case "\n\t":
                            for fi45 := 0; fi45 < 12; fi45++ {
                                v46 := v4[65535]
                                for fi47 := 0; fi47 < 16; fi47++ {
                                    v48 := v4[17]
                                    var v49 string = ("\n\t" + "hello")
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                                var v50 int8 = -128
                                _ = fi45
                                _ = v46
                                _ = v50
                            }
                        default:
                            if (1 >= 505) {
                                for fi51 := 0; fi51 < 13; fi51++ {
                                    if true {
                                        if (fi51 <= int(v5)) {
                                            var v52 int16 = (v13 & v13)
                                            for fi53 := 0; fi53 < 7; fi53++ {
                                                v5 = v3
                                                var v54 uint16 = 2
                                                switch (("\n\t" + "\n\t") + ("" + ("hello" + "hello"))) {
                                                case "\n\t":
                                                    if true {
                                                        v2 = append(v2, 463)
                                                        v6 += v6
                                                        var v55 int64 = -111
                                                        _ = v55
                                                    }
                                                case "":
                                                    if len(v2) > 0 {
                                                        v2[0] = 804
                                                    }
                                                default:
                                                    var v56 uint32 = 274
                                                    _ = v56
                                                }
                                                _ = fi53
                                                _ = v54
                                            }
                                            var v57 uint = v1
                                            _ = v52
                                            _ = v57
                                        } else {
                                            v13 = (v13 - (v13 * (v13 & (-(v13)))))
                                            var v58 int32 = -552
                                            _ = v58
                                        }
                                        for fi59 := 0; fi59 < 15; fi59++ {
                                            v3 -= 2
                                            fmt.Println(v3, v3)
                                            _ = fi59
                                        }
                                        delete(v4, 42)
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = 604
                                        }
                                    }
                                    var v60 int8 = -128
                                    _ = fi51
                                    _ = v60
                                }
                            } else {
                                v1 = (v1 * (v1 | v1))
                            }
                        }
                    }
                case 0:
                    switch "hello" {
                    case "hello":
                        fmt.Println(v6, v6, v13)
                    case "foo bar":
                        if (1 > 2147483647) {
                            v61 := GenS10{F0: (-1 <= -1), F1: 2, F2: 1}
                            v1 = uint(v5)
                            var v62 int8 = -65
                            _ = v61
                            _ = v62
                        } else {
                            v6 = v6
                        }
                    default:
                        if ((true && true) || true) {
                            if (("\n\t" + "") > "\n\t") {
                                v4[uint32(v1)] = 1
                                for fi63 := 0; fi63 < 15; fi63++ {
                                    switch "hello" {
                                    case "\n\t":
                                        v13 = v13
                                    case "hello":
                                        fmt.Println(v5)
                                    case "":
                                        var v64 uint8 = v5
                                        _ = v64
                                    default:
                                        fmt.Println(v5, fi63, fi63)
                                    }
                                    _ = fi63
                                }
                                var v65 uint32 = 0
                                _ = v65
                            } else {
                                delete(v4, 898)
                            }
                        }
                    }
                case 4294967295:
                    if false {
                        v6 -= (v6 / float32(1e10))
                    } else {
                        v5 -= (v5 * v3)
                        switch "foo bar" {
                        case "hello":
                            v66 := v4[1]
                            _ = v66
                        }
                        _ = len(v2)
                    }
                default:
                    if true {
                        for fi67 := 0; fi67 < 2; fi67++ {
                            v5 = (v5 + v3)
                            _ = fi67
                        }
                        fmt.Println(v3, v6, v3)
                    } else {
                        switch 488 {
                        case -2147483648:
                            v1 -= 159
                        }
                    }
                }
            } else {
                var v68 uint8 = 127
                _ = v68
            }
            _ = v31
        }
        fmt.Println(v3)
        v69 := GenS3{F0: 42, F1: 55}
        _ = v69
    }
    v2 = append(v2, 42)
    var v70 int = -809
    for fi71 := 0; fi71 < 16; fi71++ {
        switch (^(fi71)) {
        case 9223372036854775807:
            for fi72 := 0; fi72 < 3; fi72++ {
                if (false && (true || true)) {
                    switch (fi72 >> ((v1) & 63)) {
                    case -894:
                        if true {
                            var v73 uint8 = v3
                            _ = v73
                        } else {
                            var v74 uint16 = genFn0(-986, (fi72 + genFn1((v5 + 2))), 936)
                            if true {
                                switch ("" + "foo bar") {
                                case "\n\t":
                                    v2 = append(v2, uint32(v5))
                                case "foo bar":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case "a":
                                    _ = len(v2)
                                default:
                                    var v75 int16 = (v13 % (((^(v13))) | 1))
                                    _ = v75
                                }
                            }
                            fmt.Println(fi72, v3)
                            fmt.Println(v13)
                            var v76 uint64 = 1
                            var v77 int16 = v13
                            _ = v74
                            _ = v76
                            _ = v77
                        }
                    }
                    var v78 uint8 = v3
                    _ = v78
                } else {
                    switch ("" + "") {
                    case "":
                        switch v70 {
                        case -9223372036854775808:
                            fmt.Println(v5, v70)
                        default:
                            var v79 uint16 = 65535
                            _ = v79
                        }
                    case "hello":
                        v80 := GenS7{F0: "", F1: v3, F2: -137}
                        _ = v80
                    case "foo bar":
                        v3 -= (v3 + (^(v3)))
                    default:
                        v6 -= (v6 - v6)
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v81 float32 = (v6 - (3.14159 * v6))
                _ = fi72
                _ = v81
            }
        case -9223372036854775808:
            _ = len(v4)
        default:
            for fi82 := 0; fi82 < 18; fi82++ {
                if false {
                    if false {
                        v1 = (v1 % (((v1 + (v1 >> (((^(v1))) & 63)))) | 1))
                    }
                    _ = len(v4)
                    var v83 float32 = (float32((0.1 * 1.5)) / float32(3.14159))
                    _ = v83
                } else {
                    var v84 int64 = 42
                    _ = v84
                }
                _ = fi82
            }
        }
        for fi85 := 0; fi85 < 10; fi85++ {
            switch (("hello" + "") + "\n\t") {
            case "\n\t":
                switch fi71 {
                case 4294967295:
                    var v86 bool = true
                    var v87 uint32 = 2147483648
                    _ = v86
                    _ = v87
                case 32768:
                    v88 := GenS2{F0: 263, F1: -325, F2: v1}
                    _ = v88
                case 9223372036854775807:
                    switch (fi71 ^ -285) {
                    case 0:
                        v89 := GenS11{F0: ("a" + "\n\t")}
                        _ = v89
                    case 42:
                        var v90 uint8 = 255
                        _ = v90
                    default:
                        var v91 uint64 = 18446744073709551615
                        _ = v91
                    }
                default:
                    if false {
                        switch "a" {
                        case "a":
                            var v92 uint32 = 0
                            _ = v92
                        case "\n\t":
                            var v93 uint8 = (v5 & v5)
                            _ = v93
                        case "":
                            switch "a" {
                            case "foo bar":
                                if (v13 >= v13) {
                                    switch (v70 | v70) {
                                    case -665:
                                        v94 := GenS12{F0: ""}
                                        _ = v94
                                    case -9223372036854775808:
                                        v70 -= v70
                                    }
                                    var v95 uint = 18446744073709551615
                                    _ = v95
                                }
                            }
                        default:
                            var v96 int = genFn1((v3 - v5))
                            _ = v96
                        }
                    }
                }
            case "":
                v97 := GenS13{F0: 0, F1: 2147483648, F2: 2147483647}
                _ = v97
            default:
                switch "" {
                case "\n\t":
                    v70 += fi85
                    delete(v4, 235)
                    v6 = v6
                    if true {
                        v98 := GenS7{F0: "\n\t", F1: 15, F2: 0}
                        fmt.Println(v1)
                        switch fi85 {
                        case 0:
                            var v99 uint64 = 0
                            _ = v99
                        case -347:
                            v100 := GenS12{F0: "a"}
                            _ = v100
                        case 32768:
                            v101 := GenS14{F0: uint32(v5), F1: 439}
                            _ = v101
                        }
                        _ = v98
                    } else {
                        switch 905 {
                        case -9223372036854775808:
                            var v102 bool = true
                            _ = v102
                        case 9223372036854775807:
                            fmt.Println(v1, v6)
                        default:
                            v6 = 1e10
                        }
                    }
                    var v103 float64 = float64(-2.25)
                    _ = v103
                }
            }
            _ = fi85
        }
        if (!false) {
            switch "\n\t" {
            case "a":
                fmt.Println(v1, v5)
            default:
                switch "\n\t" {
                case "":
                    fmt.Println(v3, v3)
                }
            }
            if false {
                if false {
                    v104 := GenS2{F0: 2147483647, F1: 923, F2: 4294967295}
                    v105 := v4[0]
                    _ = v104
                    _ = v105
                } else {
                    v4[uint32(v1)] = 65535
                }
                if false {
                    v2 = append(v2, 4294967295)
                }
                var v106 string = ""
                _ = v106
            }
        } else {
            delete(v4, 4294967295)
            v107 := GenS6{F0: 1, F1: 2147483647, F2: v5}
            _ = v107
        }
        _ = fi71
    }
    v1 = (v1 / (((v1 << (((v1 - v1)) & 63))) | 1))
    if true {
        _ = len(v2)
    } else {
        fmt.Println(v13, v6)
        delete(v4, 2147483647)
    }
    v4[42] = 372
    var v108 float32 = (v6 - float32((-2.25 * 0.1)))
    v3 = (v5 % ((127) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v108
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
}
