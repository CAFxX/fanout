package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int8; F1 int32; F2 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint8; F1 int32; F2 int64 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 bool; F1 uint; F2 string }
type GenS14 struct { F0 int8; F1 uint64 }
type GenS15 struct { F0 string; F1 uint16; F2 int64 }
func genFn0(p0 uint, p1 int, p2 string) uint8 {
    if (127 >= -32) {
        var v6 int = p1
        if true {
            v7 := GenS1{F0: -279}
            _ = v7
        }
        _ = v6
    }
    fmt.Println(p1)
    p2 = (((p2 + p2) + ("hello" + (p2 + p2))) + p2)
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 int64, p1 string, p2 int) uint {
    if (0 > (p0 | (p0 - (p0 | (p0 % ((p0) | 1)))))) {
        for fi8 := 0; fi8 < 11; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 int16 = 857
                v11 := GenS1{F0: 438}
                _ = fi9
                _ = v10
                _ = v11
            }
            var v12 uint16 = 42
            _ = fi8
            _ = v12
        }
    }
    switch p1 {
    case "foo bar":
        v13 := GenS2{F0: (p2 << ((18446744073709551615) & 63)), F1: 42}
        _ = v13
    }
    var v14 bool = false
    p1 = (p1 + p1)
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := []int{9223372036854775807, -871}
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 153
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{}
    v3 := make(map[int]uint32)
    v4 := []uint32{916}
    var v5 int64 = 0
    if (false || (false || (42 <= 0))) {
        switch "a" {
        case "\n\t":
            v5 = -53
        case "foo bar":
            _ = len(v3)
        }
        v1["\n\t"] = (!true)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (true || (42 <= -404)) {
        fmt.Println(v5)
        var v17 uint8 = 127
        v18 := GenS2{F0: -1, F1: 42}
        _ = v17
        _ = v18
    } else {
        switch 213 {
        case 2147483647:
            switch 9223372036854775807 {
            case 0:
                var v19 int = -2147483648
                _ = v19
            }
        default:
            switch 128 {
            case 32767:
                switch 32768 {
                case 888:
                    for fi20 := 0; fi20 < 2; fi20++ {
                        switch ("hello" + "\n\t") {
                        case "hello":
                            if false {
                                v5 += 0
                                if true {
                                    v21 := GenS2{F0: fi20, F1: -1}
                                    switch ("\n\t" + "a") {
                                    case "\n\t":
                                        var v22 bool = false
                                        _ = v22
                                    }
                                    _ = v21
                                } else {
                                    v5 = 42
                                    if (2147483648 >= 601) {
                                        for fi23 := 0; fi23 < 8; fi23++ {
                                            if len(v4) > 0 {
                                                v4[0] = 1
                                            }
                                            if true {
                                                switch -289 {
                                                case 751:
                                                    fmt.Println(fi20)
                                                }
                                                var v24 uint = 1
                                                if (!false) {
                                                    var v25 int8 = 113
                                                    _ = v25
                                                } else {
                                                    v5 = 623
                                                }
                                                v5 = v5
                                                _ = v24
                                            } else {
                                                v5 = v5
                                            }
                                            _ = fi23
                                        }
                                        v26 := GenS3{F0: 58, F1: 1, F2: 2}
                                        v5 = -1
                                        _ = v26
                                    } else {
                                        v4 = append(v4, 0)
                                        fmt.Println(fi20, v5)
                                    }
                                    v5 += -9223372036854775808
                                }
                                for fi27 := 0; fi27 < 10; fi27++ {
                                    var v28 int64 = v5
                                    v29 := GenS2{F0: 9223372036854775807, F1: 2147483647}
                                    if len(v4) > 0 {
                                        v4[0] = 4294967295
                                    }
                                    switch "" {
                                    case "\n\t":
                                        var v30 uint64 = 340
                                        _ = v30
                                    case "":
                                        fmt.Println(v5, v28, fi27)
                                    case "foo bar":
                                        v31 := GenS4{F0: 127}
                                        if true {
                                            switch "" {
                                            case "\n\t":
                                                v28 += 1
                                            case "hello":
                                                if true {
                                                    fmt.Println(v5, v28)
                                                    if (42 < -2147483648) {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        v4 = append(v4, 78)
                                                    } else {
                                                        _ = len(v1)
                                                    }
                                                    v28 = v28
                                                }
                                            default:
                                                if (255 <= 43) {
                                                    var v32 uint8 = 255
                                                    var v33 uint = genFn1((v28 % (((v5 * (v5 | (v28 << ((255) & 63))))) | 1)), ("hello" + ("\n\t" + "foo bar")), (fi27 * (fi27 / ((fi27) | 1))))
                                                    switch 839 {
                                                    case 42:
                                                        var v34 int32 = -2147483648
                                                        fmt.Println(v33, fi27, fi20)
                                                        _ = v34
                                                    case 128:
                                                        if true {
                                                            var v35 bool = ("" != ("hello" + "hello"))
                                                            _ = len(v3)
                                                            if v35 {
                                                                var v36 float64 = 0.1
                                                                _ = v36
                                                            } else {
                                                                v28 += v5
                                                                v32 = v32
                                                            }
                                                            _ = v35
                                                        } else {
                                                            var v37 float32 = ((0.0 / 1.0) * float32(0.0))
                                                            _ = v37
                                                        }
                                                    case 32767:
                                                        v32 += v32
                                                        v32 -= 0
                                                        switch "a" {
                                                        case "a":
                                                            if false {
                                                                var v38 uint = v33
                                                                _ = v38
                                                            } else {
                                                                if (true || true) {
                                                                    _ = len(v2)
                                                                    v28 = int64(v32)
                                                                }
                                                            }
                                                        case "foo bar":
                                                            v28 = v28
                                                        case "hello":
                                                            var v39 int32 = 2147483647
                                                            if (((32768 == fi20) || (true && false)) || true) {
                                                                v40 := GenS5{F0: 42, F1: (fi20 ^ -766)}
                                                                var v41 int64 = v5
                                                                var v42 int64 = (v41 | v28)
                                                                _ = v40
                                                                _ = v41
                                                                _ = v42
                                                            }
                                                            v43 := GenS4{F0: 60}
                                                            _ = v39
                                                            _ = v43
                                                        }
                                                    default:
                                                        fmt.Println(fi27, fi20, fi20)
                                                    }
                                                    fmt.Println(fi20, v33)
                                                    var v44 int16 = 242
                                                    _ = v32
                                                    _ = v33
                                                    _ = v44
                                                }
                                            }
                                            switch (fi27 + fi27) {
                                            case 1:
                                                if true {
                                                    if (fi20 > (fi27 + (fi27 + -2147483648))) {
                                                        var v45 float64 = 1e10
                                                        var v46 float32 = ((0.1 / (3.14159 - 1.5)) + 0.0)
                                                        v47 := GenS6{F0: 1}
                                                        _ = v45
                                                        _ = v46
                                                        _ = v47
                                                    } else {
                                                        v48 := GenS7{F0: 128}
                                                        fmt.Println(v28, fi20, v5)
                                                        var v49 int = -222
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                } else {
                                                    if false {
                                                        v50 := GenS8{F0: 100}
                                                        _ = v50
                                                    }
                                                    var v51 uint32 = 42
                                                    _ = v51
                                                }
                                            case 9223372036854775807:
                                                v52 := v3[(-(fi27))]
                                                _ = v52
                                            default:
                                                if false {
                                                    if true {
                                                        var v53 float32 = 1.5
                                                        if len(v2) > 0 {
                                                            v2[0] = ""
                                                        }
                                                        _ = v53
                                                    }
                                                    var v54 int32 = 1
                                                    delete(v1, ("a" + "a"))
                                                    var v55 bool = true
                                                    _ = v54
                                                    _ = v55
                                                } else {
                                                    fmt.Println(fi27, v28)
                                                    var v56 int = (fi20 ^ fi27)
                                                    _ = v56
                                                }
                                            }
                                            fmt.Println(v28, fi27, v5)
                                        }
                                        var v57 uint32 = 2
                                        v58 := GenS6{F0: 65535}
                                        _ = v31
                                        _ = v57
                                        _ = v58
                                    }
                                    var v59 uint8 = 114
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                    _ = v59
                                }
                            } else {
                                v5 += -736
                            }
                        case "foo bar":
                            for fi60 := 0; fi60 < 14; fi60++ {
                                switch fi60 {
                                case 8:
                                    switch "a" {
                                    case "\n\t":
                                        v5 -= 0
                                    }
                                default:
                                    var v61 uint32 = 0
                                    _ = v61
                                }
                                _ = fi60
                            }
                        default:
                            v62 := GenS9{F0: 255}
                            _ = v62
                        }
                        v63 := GenS10{F0: (fi20 >> ((0) & 63))}
                        for fi64 := 0; fi64 < 3; fi64++ {
                            v65 := v3[fi64]
                            v66 := GenS11{F0: 128, F1: 42, F2: v5}
                            fmt.Println(fi64, fi64)
                            _ = len(v1)
                            _ = fi64
                            _ = v65
                            _ = v66
                        }
                        _ = fi20
                        _ = v63
                    }
                }
            case 1:
                if ((4294967295 <= 977) || (!(!true))) {
                    switch 165 {
                    case 525:
                        var v67 uint32 = 2147483647
                        _ = v67
                    case 0:
                        fmt.Println(v5)
                    case 2147483647:
                        switch -2147483648 {
                        case 32768:
                            for fi68 := 0; fi68 < 4; fi68++ {
                                var v69 int16 = -652
                                var v70 uint32 = 4294967295
                                _ = fi68
                                _ = v69
                                _ = v70
                            }
                        case 4294967295:
                            v4 = append(v4, 4294967295)
                        case 2147483647:
                            v1[("hello" + "a")] = (-9223372036854775808 <= v5)
                        }
                    default:
                        for fi71 := 0; fi71 < 11; fi71++ {
                            v72 := GenS4{F0: -1}
                            v5 -= v5
                            var v73 string = ("hello" + "hello")
                            _ = fi71
                            _ = v72
                            _ = v73
                        }
                    }
                    var v74 int32 = 2147483647
                    _ = v74
                } else {
                    v5 += v5
                    delete(v3, 32767)
                }
            }
        }
    }
    v5 -= 650
    if (!true) {
        v5 -= 122
        fmt.Println(v5)
    } else {
        fmt.Println(v5)
        var v75 int8 = 100
        v5 = v5
        _ = v75
    }
    for fi76 := 0; fi76 < 13; fi76++ {
        v77 := GenS12{F0: 665}
        v5 = (v5 | (v5 << ((396) & 63)))
        switch (fi76 >> ((2) & 63)) {
        case 9223372036854775807:
            for fi78 := 0; fi78 < 7; fi78++ {
                delete(v1, "\n\t")
                if false {
                    if ((!true) && (675 != 943)) {
                        v79 := GenS13{F0: (-1 < -104), F1: 255, F2: ((("a" + "foo bar") + ("foo bar" + "foo bar")) + ("foo bar" + ("hello" + ("foo bar" + "hello"))))}
                        var v80 string = ("" + "foo bar")
                        v5 = v5
                        switch v80 {
                        case "":
                            var v81 bool = (false || (3.14159 < 0.0))
                            _ = v81
                        default:
                            switch (v80 + v80) {
                            case "":
                                if false {
                                    v80 = (v80 + (v80 + v80))
                                }
                            case "\n\t":
                                if true {
                                    if false {
                                        fmt.Println(v80, v5)
                                        fmt.Println(v80, v80, v5)
                                    }
                                    fmt.Println(fi76, v5)
                                    v5 -= 42
                                }
                            case "foo bar":
                                v82 := GenS14{F0: 1, F1: 4294967295}
                                _ = v82
                            }
                        }
                        _ = v79
                        _ = v80
                    }
                    var v83 int8 = 0
                    v83 = v83
                    _ = v83
                } else {
                    var v84 int32 = 0
                    v85 := GenS11{F0: 42, F1: (v84 | v84), F2: (v5 - v5)}
                    _ = v84
                    _ = v85
                }
                if (true || false) {
                    v5 = (v5 + int64(fi78))
                    _ = len(v2)
                    v4 = append(v4, 2)
                } else {
                    var v86 uint32 = 140
                    _ = v86
                }
                _ = fi78
            }
        case 328:
            switch "\n\t" {
            case "":
                for fi87 := 0; fi87 < 11; fi87++ {
                    switch fi76 {
                    case -424:
                        v88 := GenS2{F0: -225, F1: -343}
                        _ = v88
                    case 0:
                        fmt.Println(fi76, v5)
                    }
                    _ = fi87
                }
            case "foo bar":
                var v89 uint64 = 314
                _ = v89
            default:
                v5 = 1
            }
        case 122:
            var v90 uint64 = 42
            _ = v90
        default:
            if (!true) {
                if true {
                    v5 -= (v5 << ((255) & 63))
                    delete(v3, (^(fi76)))
                }
                if ((!(1 != -128)) && (!("hello" != "hello"))) {
                    v5 = (v5 + v5)
                    var v91 int8 = -128
                    _ = v91
                }
                switch fi76 {
                case 0:
                    for fi92 := 0; fi92 < 10; fi92++ {
                        if true {
                            v5 = 42
                            v4 = append(v4, 2)
                            delete(v1, (("\n\t" + ("\n\t" + "\n\t")) + "foo bar"))
                        }
                        _ = fi92
                    }
                case 32767:
                    fmt.Println(fi76)
                default:
                    var v93 bool = true
                    _ = v93
                }
            }
        }
        _ = fi76
        _ = v77
    }
    v5 = v5
    v2 = append(v2, "")
    v5 = -9223372036854775808
    v4 = append(v4, 49)
    if (1 != 42) {
        v4 = append(v4, 2)
        fmt.Println(v5)
        v94 := GenS15{F0: "a", F1: 255, F2: (-(v5))}
        _ = v94
    }
    if (false || true) {
        v95 := GenS9{F0: 255}
        v5 = -489
        var v96 int16 = 1000
        _ = v95
        _ = v96
    } else {
        var v97 uint16 = 2
        _ = v97
    }
    var v98 int8 = -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v98
}
