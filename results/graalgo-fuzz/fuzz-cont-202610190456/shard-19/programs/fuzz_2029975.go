package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 int64; F1 uint8; F2 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int; F1 uint64; F2 int32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int16; F1 bool }
type GenS11 struct { F0 int16; F1 uint64 }
type GenS12 struct { F0 int8; F1 uint16 }
type GenS13 struct { F0 int; F1 int }
func genFn0(p0 int32) bool {
    v5 := GenS1{F0: -798}
    var v6 uint32 = 4294967295
    var v7 uint8 = 42
    var v8 uint = 345
    if false {
        for fi9 := 0; fi9 < 15; fi9++ {
            switch "foo bar" {
            case "\n\t":
                p0 += p0
            case "foo bar":
                v10 := GenS1{F0: fi9}
                _ = v10
            }
            v11 := GenS1{F0: fi9}
            _ = fi9
            _ = v11
        }
        var v12 uint8 = v7
        v13 := []int{128}
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 15; fi14++ {
            switch (("" + "foo bar") + "a") {
            case "hello":
                switch "a" {
                case "hello":
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v7 += 6
                        v8 = (v8 & v8)
                        _ = fi15
                    }
                case "":
                    if (v7 != (v7 & v7)) {
                        if (!(!(-9223372036854775808 == 9223372036854775807))) {
                            v16 := GenS2{F0: 0, F1: 1}
                            v7 = (v7 & 2)
                            _ = v16
                        } else {
                            p0 -= 0
                        }
                    }
                case "a":
                    v7 += v7
                default:
                    v17 := GenS2{F0: 0, F1: 100}
                    _ = v17
                }
            case "\n\t":
                if true {
                    if true {
                        var v18 uint16 = uint16(v7)
                        switch fi14 {
                        case 9223372036854775807:
                            v7 -= v7
                        default:
                            for fi19 := 0; fi19 < 15; fi19++ {
                                v20 := []int{int(v8), fi14, 128, fi19}
                                var v21 uint64 = 42
                                v22 := GenS2{F0: -128, F1: int8(v6)}
                                _ = fi19
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                        }
                        switch (-(fi14)) {
                        case 42:
                            v23 := []int{(fi14 << ((165) & 63))}
                            _ = v23
                        case 415:
                            v24 := GenS2{F0: 127, F1: 100}
                            _ = v24
                        default:
                            var v25 uint64 = 18446744073709551615
                            _ = v25
                        }
                        _ = v18
                    } else {
                        v26 := GenS3{F0: (v8 + (^(v8)))}
                        _ = v26
                    }
                    v27 := GenS4{F0: -128, F1: 53}
                    _ = v27
                }
            case "a":
                v28 := []int{32767, (-(fi14))}
                _ = v28
            default:
                if false {
                    if false {
                        v29 := []int{9223372036854775807, fi14, fi14, fi14}
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 20; fi30++ {
                        if true {
                            v31 := GenS2{F0: 0, F1: -1}
                            v32 := GenS5{F0: 0, F1: v7, F2: -1}
                            v33 := []int{(^(fi14)), fi14, fi14}
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        _ = fi30
                    }
                    v34 := []int{0, (fi14 + (fi14 ^ fi14))}
                    _ = v34
                }
            }
            for fi35 := 0; fi35 < 13; fi35++ {
                v36 := GenS6{F0: -1}
                v6 = 1
                _ = fi35
                _ = v36
            }
            _ = fi14
        }
        for fi37 := 0; fi37 < 10; fi37++ {
            fmt.Println(v6, v8)
            for fi38 := 0; fi38 < 19; fi38++ {
                fmt.Println(v6)
                if ((1e10 * 1.5) < (1e10 * (1e10 * (1e10 - float64((-2.25 + -2.25)))))) {
                    v39 := GenS2{F0: 1, F1: 124}
                    fmt.Println(p0)
                    _ = v39
                } else {
                    var v40 uint64 = 18446744073709551615
                    v41 := []int{fi38}
                    v42 := GenS5{F0: 9223372036854775807, F1: (v7 & 127), F2: int16(p0)}
                    _ = v40
                    _ = v41
                    _ = v42
                }
                v6 = v6
                if false {
                    if true {
                        var v43 uint64 = 2
                        v44 := []int{fi38, fi38, fi38, fi37}
                        p0 = (p0 % ((p0) | 1))
                        v44 = append(v44, 32767)
                        _ = v43
                        _ = v44
                    }
                    v7 = (^(v7))
                    var v45 int32 = (^(p0))
                    var v46 bool = (0.1 != -2.25)
                    switch (fi37 | fi38) {
                    case 207:
                        var v47 uint = 4294967295
                        _ = v47
                    case 4294967295:
                        v48 := GenS7{F0: (fi38 + fi38)}
                        _ = v48
                    }
                    _ = v45
                    _ = v46
                }
                var v49 float64 = -2.25
                _ = fi38
                _ = v49
            }
            _ = fi37
        }
        v50 := GenS8{F0: -689, F1: 2, F2: (^(p0))}
        _ = v50
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (!false)
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 32767
    var v3 int = -938
    v4 := make(map[string]uint32)
    v3 = (v3 + v3)
    switch (v3 + -447) {
    case 48:
        var v51 uint16 = (v2 << ((776) & 63))
        _ = v51
    case 32768:
        v52 := GenS9{F0: -253}
        _ = v52
    case -9223372036854775808:
        v4[("\n\t" + ("foo bar" + "hello"))] = 2
    }
    v3 -= (v3 + v3)
    v3 += v3
    fmt.Println(v2)
    v2 = v2
    fmt.Println(v3)
    if (v3 != v3) {
        v53 := v4["\n\t"]
        _ = len(v1)
        v2 -= v2
        _ = v53
    } else {
        v1[593] = (!false)
        v3 -= 1
        var v54 int = -870
        _ = v54
    }
    switch v3 {
    case -9223372036854775808:
        for fi55 := 0; fi55 < 19; fi55++ {
            fmt.Println(v3, fi55)
            _ = fi55
        }
    case 0:
        var v56 uint8 = 0
        _ = v56
    case 2147483647:
        fmt.Println(v3, v2)
    default:
        v2 -= (v2 - 32767)
    }
    v57 := GenS9{F0: 32767}
    var v58 float32 = 3.14159
    switch -175 {
    case -86:
        if ((0.0 > 1.5) && (false && (v58 <= float32((1e10 * 3.14159))))) {
            for fi59 := 0; fi59 < 4; fi59++ {
                switch (^(v3)) {
                case -1:
                    _ = len(v4)
                case 42:
                    delete(v4, "")
                case 9223372036854775807:
                    var v60 uint = 952
                    _ = v60
                }
                var v61 int = (fi59 >> ((4294967295) & 63))
                v62 := GenS9{F0: 32767}
                _ = fi59
                _ = v61
                _ = v62
            }
        } else {
            var v63 string = ""
            if (127 < 1) {
                var v64 int32 = 1
                switch (v3 * 1) {
                case 888:
                    switch (v3 << ((42) & 63)) {
                    case 961:
                        var v65 int8 = -128
                        _ = v65
                    case 1:
                        for fi66 := 0; fi66 < 18; fi66++ {
                            var v67 string = v63
                            fmt.Println(v58)
                            fmt.Println(fi66)
                            _ = fi66
                            _ = v67
                        }
                    }
                default:
                    delete(v4, "hello")
                }
                for fi68 := 0; fi68 < 12; fi68++ {
                    if false {
                        if true {
                            for fi69 := 0; fi69 < 13; fi69++ {
                                var v70 float64 = 1e10
                                var v71 string = (v63 + v63)
                                _ = fi69
                                _ = v70
                                _ = v71
                            }
                            v72 := []int{}
                            var v73 float32 = float32(1e10)
                            _ = v72
                            _ = v73
                        }
                        var v74 int64 = 574
                        _ = v74
                    }
                    var v75 int8 = 99
                    switch v63 {
                    case "":
                        for fi76 := 0; fi76 < 13; fi76++ {
                            v58 = (v58 / (v58 / 1e10))
                            _ = fi76
                        }
                    case "hello":
                        v77 := v1[42]
                        _ = v77
                    case "a":
                        for fi78 := 0; fi78 < 16; fi78++ {
                            if true {
                                if true {
                                    var v79 float32 = -2.25
                                    switch (fi68 % (((-(fi68))) | 1)) {
                                    case 472:
                                        v3 = (^(fi78))
                                    default:
                                        if (-1 > -32768) {
                                            var v80 int64 = 1
                                            switch v63 {
                                            case "a":
                                                var v81 float64 = 0.0
                                                v82 := []int{(fi78 % ((fi68) | 1))}
                                                var v83 uint16 = (v2 + 42)
                                                _ = v81
                                                _ = v82
                                                _ = v83
                                            case "hello":
                                                switch "hello" {
                                                case "hello":
                                                    var v84 int8 = 42
                                                    _ = v84
                                                case "\n\t":
                                                    if false {
                                                        var v85 float32 = (v79 + v58)
                                                        _ = v85
                                                    }
                                                    var v86 uint = 1
                                                    _ = v86
                                                }
                                            case "\n\t":
                                                v87 := GenS10{F0: -32768, F1: ("hello" != ("\n\t" + v63))}
                                                _ = v87
                                            default:
                                                v2 = (^(v2))
                                            }
                                            v88 := []int{(fi68 % ((fi78) | 1))}
                                            _ = v80
                                            _ = v88
                                        } else {
                                            var v89 uint = 2
                                            v75 = (v75 - (v75 % ((v75) | 1)))
                                            _ = v89
                                        }
                                    }
                                    _ = v79
                                } else {
                                    if (42 < -369) {
                                        v90 := GenS2{F0: v75, F1: (v75 << ((856) & 63))}
                                        v58 = v58
                                        switch "foo bar" {
                                        case "hello":
                                            var v91 int = (-(fi78))
                                            _ = v91
                                        case "\n\t":
                                            v92 := GenS11{F0: 32767, F1: 2}
                                            _ = v92
                                        case "foo bar":
                                            var v93 uint8 = 42
                                            _ = v93
                                        default:
                                            v94 := GenS12{F0: v75, F1: v2}
                                            _ = v94
                                        }
                                        _ = v90
                                    }
                                    switch "a" {
                                    case "hello":
                                        v75 = v75
                                    case "a":
                                        v95 := GenS1{F0: v3}
                                        v96 := GenS2{F0: v75, F1: v75}
                                        _ = v95
                                        _ = v96
                                    case "":
                                        v97 := GenS1{F0: (v3 % ((fi78) | 1))}
                                        _ = v97
                                    }
                                }
                            }
                            v75 = (v75 ^ v75)
                            _ = fi78
                        }
                    default:
                        if (v75 != v75) {
                            _ = len(v1)
                            fmt.Println(v58, v3, v64)
                            v3 += -818
                        } else {
                            for fi98 := 0; fi98 < 9; fi98++ {
                                v75 = 42
                                v99 := GenS13{F0: (fi98 | v3), F1: fi98}
                                _ = fi98
                                _ = v99
                            }
                        }
                    }
                    _ = fi68
                    _ = v75
                }
                _ = v64
            } else {
                if (328 == 647) {
                    v63 = (v63 + ("a" + (v63 + "hello")))
                    switch (v63 + "") {
                    case "foo bar":
                        v100 := []int{-9223372036854775808}
                        _ = v100
                    case "\n\t":
                        var v101 string = v63
                        _ = v101
                    default:
                        if true {
                            v58 -= (v58 * 1.5)
                        }
                    }
                    for fi102 := 0; fi102 < 20; fi102++ {
                        if (-1 < 1) {
                            v103 := GenS13{F0: (fi102 >> ((1) & 63)), F1: (v3 + (^(v3)))}
                            _ = v103
                        }
                        _ = fi102
                    }
                } else {
                    if true {
                        var v104 int8 = 1
                        fmt.Println(v104)
                        delete(v1, 1)
                        _ = v104
                    } else {
                        v2 += (v2 & 1)
                        if true {
                            switch v3 {
                            case -595:
                                fmt.Println(v3)
                            }
                            var v105 int32 = -419
                            _ = v105
                        } else {
                            v106 := GenS11{F0: 42, F1: 113}
                            v63 = (v63 + (((("a" + ("foo bar" + "foo bar")) + v63) + "") + "\n\t"))
                            _ = v106
                        }
                        v63 = "hello"
                    }
                    var v107 uint32 = 0
                    _ = v107
                }
                switch v3 {
                case 808:
                    v108 := v1[492]
                    _ = v108
                default:
                    v58 = v58
                }
                switch ("hello" + "hello") {
                case "a":
                    v58 += (v58 * v58)
                case "hello":
                    var v109 uint32 = 42
                    _ = v109
                case "\n\t":
                    v3 -= v3
                }
            }
            _ = len(v4)
            _ = v63
        }
    case -265:
        if (1 < 2) {
            v58 -= v58
            fmt.Println(v58, v58, v3)
        }
    }
    v2 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v58
}
