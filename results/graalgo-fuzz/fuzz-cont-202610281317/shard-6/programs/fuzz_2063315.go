package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint32; F1 bool; F2 int }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 uint64; F1 string; F2 uint8 }
type GenS7 struct { F0 uint8; F1 int; F2 int }
type GenS8 struct { F0 uint }

func main() {
    v1 := []uint32{2147483648, 4294967295, 2}
    var v2 int32 = 1
    var v3 int32 = v2
    v4 := make(map[int]bool)
    var v5 uint64 = 18446744073709551615
    v6 := make(map[string]bool)
    var v7 bool = (1 != 714)
    for fi8 := 0; fi8 < 9; fi8++ {
        if false {
            var v9 int16 = 0
            _ = v9
        }
        switch 0 {
        case -2147483648:
            var v10 int64 = 286
            _ = v10
        case -818:
            v3 += (v2 ^ (^(v2)))
        default:
            v3 = -1
        }
        fmt.Println(v5, v2, v5)
        _ = fi8
    }
    v11 := v6[(("\n\t" + ("hello" + "")) + "\n\t")]
    switch 780 {
    case -2147483648:
        delete(v4, 761)
    case -635:
        if v7 {
            if (v11 && true) {
                v5 -= v5
                var v12 float64 = -2.25
                var v13 int16 = int16(v5)
                _ = v12
                _ = v13
            }
        } else {
            var v14 uint16 = 42
            _ = v14
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if ((!(91 != 0)) && ((2147483647 < 2) || v11)) {
        switch -1 {
        case 842:
            var v15 string = ("" + "a")
            _ = v15
        case -2147483648:
            v4[32768] = (1.5 != (-2.25 - 0.0))
        default:
            v3 = 0
        }
        v11 = false
        if (0.0 <= (1e10 * 1.5)) {
            var v16 uint64 = (v5 * v5)
            switch "hello" {
            case "foo bar":
                v17 := GenS1{F0: v3, F1: 42, F2: 255}
                _ = v17
            case "":
                if (127 > 127) {
                    if (!v11) {
                        v1 = append(v1, 346)
                        var v18 float32 = 0.0
                        var v19 int64 = 1
                        _ = v18
                        _ = v19
                    }
                }
            }
            _ = v16
        } else {
            switch 469 {
            case 128:
                v5 += v5
            case -233:
                for fi20 := 0; fi20 < 16; fi20++ {
                    switch "" {
                    case "foo bar":
                        for fi21 := 0; fi21 < 17; fi21++ {
                            var v22 int = (fi20 & fi21)
                            v5 = (^(v5))
                            if v7 {
                                switch fi20 {
                                case -9223372036854775808:
                                    if false {
                                        if (2 > 1) {
                                            var v23 int16 = 0
                                            var v24 int16 = (v23 + 32767)
                                            switch ("a" + "\n\t") {
                                            case "a":
                                                switch 9223372036854775807 {
                                                case -562:
                                                    if v11 {
                                                        var v25 float64 = 0.1
                                                        v26 := GenS1{F0: 1, F1: uint8(v5), F2: 2}
                                                        _ = v25
                                                        _ = v26
                                                    } else {
                                                        v4[-658] = v7
                                                        fmt.Println(v3)
                                                        switch fi21 {
                                                        case 916:
                                                            v27 := v6["\n\t"]
                                                            _ = v27
                                                        case -361:
                                                            var v28 uint = 2
                                                            _ = v28
                                                        default:
                                                            switch fi21 {
                                                            case 32767:
                                                                var v29 bool = v7
                                                                _ = v29
                                                            default:
                                                                v30 := GenS2{F0: fi20, F1: (fi21 >= fi20)}
                                                                _ = v30
                                                            }
                                                        }
                                                        var v31 int32 = v3
                                                        _ = v31
                                                    }
                                                default:
                                                    v5 += (v5 / (((^(v5))) | 1))
                                                    var v32 int32 = v2
                                                    _ = v32
                                                }
                                            case "foo bar":
                                                v2 -= v2
                                            }
                                            var v33 uint16 = 32767
                                            if ((-2147483648 > v3) || ((v24 ^ 42) != v24)) {
                                                v34 := GenS3{F0: v33, F1: v7, F2: 0}
                                                var v35 int64 = 559
                                                _ = v34
                                                _ = v35
                                            }
                                            _ = v23
                                            _ = v24
                                            _ = v33
                                        } else {
                                            v36 := GenS3{F0: 233, F1: v7, F2: 65535}
                                            v5 += 607
                                            _ = v36
                                        }
                                    }
                                case 4294967295:
                                    v37 := GenS4{F0: 1, F1: (false || (-119 <= 0)), F2: fi21}
                                    _ = v37
                                case 1:
                                    if true {
                                        fmt.Println(v22, v11, v7)
                                        var v38 float32 = 1e10
                                        fmt.Println(fi20, v38, v5)
                                        fmt.Println(fi20, v7, v22)
                                        if (128 > 0) {
                                            var v39 int32 = 2147483647
                                            v22 = fi21
                                            if false {
                                                fmt.Println(v11, v22, v2)
                                                switch (fi21 + fi20) {
                                                case -9223372036854775808:
                                                    var v40 bool = (747 < 1)
                                                    _ = v40
                                                default:
                                                    if v7 {
                                                        var v41 uint16 = 0
                                                        v11 = (float32(1e10) == v38)
                                                        _ = v41
                                                    } else {
                                                        var v42 string = ("\n\t" + "foo bar")
                                                        _ = v42
                                                    }
                                                }
                                            }
                                            _ = v39
                                        } else {
                                            if true {
                                                if (-2.25 != 1.5) {
                                                    if len(v1) > 0 {
                                                        v1[0] = 89
                                                    }
                                                } else {
                                                    var v43 int = (fi21 >> ((945) & 63))
                                                    _ = v43
                                                }
                                                var v44 float64 = (3.14159 - (1e10 * 1e10))
                                                var v45 float32 = (0.1 + (-2.25 / 3.14159))
                                                switch (fi20 % ((fi21) | 1)) {
                                                case 401:
                                                    switch "foo bar" {
                                                    case "a":
                                                        fmt.Println(fi20, v22)
                                                    case "foo bar":
                                                        var v46 int = (v22 ^ fi20)
                                                        _ = v46
                                                    }
                                                case -9223372036854775808:
                                                    var v47 int64 = -9223372036854775808
                                                    _ = v47
                                                case 2147483647:
                                                    switch 4294967295 {
                                                    case 648:
                                                        var v48 int8 = 100
                                                        fmt.Println(fi21, fi20, v11)
                                                        _ = v48
                                                    default:
                                                        v11 = (((!v11) || true) && false)
                                                    }
                                                }
                                                v49 := GenS5{F0: 1, F1: 9223372036854775807}
                                                _ = v44
                                                _ = v45
                                                _ = v49
                                            } else {
                                                v50 := v6[("\n\t" + "\n\t")]
                                                v38 = v38
                                                _ = v50
                                            }
                                            switch "" {
                                            case "hello":
                                                fmt.Println(v7, fi20, v22)
                                            case "\n\t":
                                                var v51 int64 = -295
                                                _ = v51
                                            case "":
                                                v52 := GenS2{F0: fi21, F1: false}
                                                _ = v52
                                            default:
                                                var v53 int64 = -565
                                                _ = v53
                                            }
                                        }
                                        _ = v38
                                    } else {
                                        fmt.Println(fi20, v22, fi20)
                                        var v54 float32 = ((-2.25 - 0.0) - 1.5)
                                        v4[fi20] = (-1 != 605)
                                        _ = v54
                                    }
                                }
                                if (v7 && v7) {
                                    if ((^(v22)) != fi21) {
                                        fmt.Println(v3, v7, v2)
                                    }
                                    fmt.Println(v11, fi21)
                                    v3 -= (v3 - (^(v2)))
                                    switch "\n\t" {
                                    case "\n\t":
                                        if false {
                                            v11 = true
                                            if v7 {
                                                var v55 uint16 = 0
                                                _ = v55
                                            } else {
                                                v3 += -2147483648
                                                var v56 int8 = 100
                                                _ = len(v1)
                                                v56 += -1
                                                _ = v56
                                            }
                                            v7 = v7
                                        }
                                    case "foo bar":
                                        _ = len(v6)
                                    default:
                                        var v57 int16 = -767
                                        _ = v57
                                    }
                                    v58 := GenS6{F0: v5, F1: ("\n\t" + "foo bar"), F2: 237}
                                    if (42 >= 2) {
                                        v4[(v22 + (fi20 + (fi20 % ((fi20) | 1))))] = v7
                                    }
                                    v59 := GenS3{F0: 570, F1: false, F2: 128}
                                    v3 = 230
                                    _ = v58
                                    _ = v59
                                }
                            }
                            var v60 int32 = 1
                            _ = fi21
                            _ = v22
                            _ = v60
                        }
                    default:
                        var v61 bool = v11
                        _ = v61
                    }
                    if v11 {
                        var v62 float32 = 3.14159
                        _ = v62
                    }
                    _ = fi20
                }
            }
            var v63 uint32 = 431
            var v64 float64 = float64(1.5)
            _ = v63
            _ = v64
        }
    }
    _ = len(v4)
    fmt.Println(v11, v2)
    delete(v4, 1)
    v65 := GenS7{F0: 255, F1: -2147483648, F2: -1}
    switch 1 {
    case 127:
        v7 = false
    }
    var v66 float64 = -2.25
    for fi67 := 0; fi67 < 17; fi67++ {
        switch fi67 {
        case 128:
            v68 := GenS8{F0: 2}
            _ = v68
        default:
            v69 := v4[-905]
            _ = v69
        }
        var v70 uint32 = 2
        switch ("" + "a") {
        case "hello":
            _ = len(v1)
        case "\n\t":
            fmt.Println(v2)
        default:
            v71 := GenS7{F0: 42, F1: int(v70), F2: fi67}
            _ = v71
        }
        _ = fi67
        _ = v70
    }
    var v72 uint64 = (^(v5))
    delete(v6, "a")
    switch 0 {
    case 165:
        var v73 int = -212
        _ = v73
    case -2147483648:
        fmt.Println(v5, v5)
    }
    switch "\n\t" {
    case "foo bar":
        var v74 bool = ((261 > 128) && (true && (!((0 == 431) && (42 == 4294967295)))))
        _ = v74
    case "hello":
        fmt.Println(v5, v2, v72)
    case "\n\t":
        var v75 float32 = 3.14159
        _ = v75
    default:
        v3 = v3
    }
    if (!(42 <= 100)) {
        if v11 {
            for fi76 := 0; fi76 < 2; fi76++ {
                var v77 uint32 = 2
                var v78 float64 = v66
                fmt.Println(v5, v72, v11)
                _ = fi76
                _ = v77
                _ = v78
            }
        } else {
            delete(v4, int(v5))
        }
    }
    for fi79 := 0; fi79 < 20; fi79++ {
        v2 -= v3
        var v80 uint32 = 4294967295
        _ = fi79
        _ = v80
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v66:", v66)
    fmt.Println("v7:", v7)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v7
    _ = v72
}
