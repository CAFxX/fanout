package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS10 struct { F0 int; F1 int8; F2 uint64 }
type GenS11 struct { F0 uint16 }
func genFn0() int {
    switch ("a" + "hello") {
    case "hello":
        var v6 uint64 = 2
        _ = v6
    case "foo bar":
        var v7 int16 = 569
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := []int{}
            for fi10 := 0; fi10 < 16; fi10++ {
                if true {
                    var v11 uint16 = 203
                    fmt.Println(v11, fi8)
                    _ = v11
                }
                var v12 float64 = (3.14159 + 0.0)
                var v13 int = (^(fi8))
                if ((0 > 255) && (42 >= 915)) {
                    v14 := GenS1{F0: 614, F1: (4294967295 <= 0)}
                    v15 := GenS1{F0: 0, F1: false}
                    _ = v14
                    _ = v15
                }
                _ = fi10
                _ = v12
                _ = v13
            }
            if (!((true || (-68 < 100)) && ((fi8 >> ((2) & 63)) > -605))) {
                if ((1000 != -1) && false) {
                    for fi16 := 0; fi16 < 20; fi16++ {
                        v17 := GenS2{F0: true, F1: 646, F2: 0}
                        _ = fi16
                        _ = v17
                    }
                } else {
                    v9 = append(v9, (-(fi8)))
                    var v18 int32 = -2147483648
                    fmt.Println(v18)
                    _ = v18
                }
                var v19 uint16 = 255
                _ = v19
            } else {
                _ = len(v9)
            }
            _ = fi8
            _ = v9
        }
    default:
        v20 := GenS3{F0: (3.14159 > ((0.0 - ((0.1 - (0.1 + 1e10)) + float64((-2.25 * float32(0.0))))) - 0.1)), F1: 12}
        _ = v20
    }
    var v21 float64 = 0.1
    switch "\n\t" {
    case "a":
        if (true || false) {
            v22 := []int{}
            _ = v22
        }
    case "\n\t":
        for fi23 := 0; fi23 < 9; fi23++ {
            if false {
                v24 := GenS1{F0: 571, F1: true}
                if ((2 >= 789) || false) {
                    v25 := []int{}
                    if (-2147483648 > 42) {
                        var v26 int16 = 322
                        fmt.Println(v21, fi23)
                        _ = v26
                    } else {
                        var v27 int16 = 1
                        for fi28 := 0; fi28 < 16; fi28++ {
                            var v29 float64 = v21
                            var v30 bool = (0.1 >= 3.14159)
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                        _ = v27
                    }
                    fmt.Println(v21, fi23)
                    _ = v25
                } else {
                    v31 := GenS3{F0: true, F1: 0}
                    _ = v31
                }
                _ = v24
            } else {
                fmt.Println(fi23)
                if ((-299 != -32768) || (true || false)) {
                    v32 := []int{-526, 128}
                    if true {
                        v21 += v21
                    }
                    v21 = 0.0
                    _ = v32
                }
                for fi33 := 0; fi33 < 7; fi33++ {
                    var v34 int = fi33
                    var v35 int8 = 0
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            }
            v21 = float64(0.1)
            v36 := GenS4{F0: 0}
            _ = fi23
            _ = v36
        }
    }
    v37 := []int{365}
    var v38 int32 = 668
    _ = v21
    _ = v37
    _ = v38
    return 128
}

func main() {
    v1 := []uint32{65535}
    v2 := make(map[uint32]int)
    var v3 int8 = 0
    var v4 int = -341
    var v5 uint8 = 188
    if true {
        if (!true) {
            _ = len(v1)
            v2[753] = (v4 | (^(v4)))
            v5 -= 128
        }
        v39 := GenS2{F0: true, F1: 556, F2: 65535}
        fmt.Println(v4, v5, v5)
        _ = v39
    } else {
        fmt.Println(v4, v5, v4)
        switch "hello" {
        case "a":
            v4 = 127
        case "":
            var v40 int64 = 817
            _ = v40
        case "hello":
            var v41 uint8 = (v5 ^ (v5 >> ((42) & 63)))
            _ = v41
        }
    }
    v3 -= v3
    if false {
        switch genFn0() {
        case 127:
            if true {
                var v42 float32 = 0.0
                var v43 float64 = 1.5
                _ = v42
                _ = v43
            } else {
                var v44 uint = 877
                v45 := GenS5{F0: (v3 - (v3 % (((v3 >> (((v44 * 565)) & 63))) | 1))), F1: 65535}
                _ = v44
                _ = v45
            }
        case -748:
            delete(v2, 2)
        }
        v46 := GenS6{F0: -9223372036854775808}
        _ = v46
    } else {
        switch (("hello" + ("foo bar" + "")) + "foo bar") {
        case "\n\t":
            var v47 int32 = -717
            _ = v47
        case "":
            v4 += (v4 << ((1) & 63))
        case "hello":
            switch "foo bar" {
            case "\n\t":
                var v48 uint32 = 2147483648
                _ = v48
            case "hello":
                var v49 int = (-(v4))
                _ = v49
            case "foo bar":
                switch "a" {
                case "":
                    v3 = (v3 / ((v3) | 1))
                case "\n\t":
                    fmt.Println(v4, v4, v5)
                default:
                    v50 := GenS7{F0: -874}
                    _ = v50
                }
            default:
                v3 += (v3 & 0)
            }
        default:
            v51 := GenS1{F0: 42, F1: (((-32768 == -1) || (!true)) || true)}
            _ = v51
        }
        for fi52 := 0; fi52 < 16; fi52++ {
            switch fi52 {
            case 32767:
                fmt.Println(v4)
            case 127:
                var v53 bool = false
                _ = v53
            default:
                var v54 uint64 = 42
                _ = v54
            }
            var v55 uint16 = 632
            _ = fi52
            _ = v55
        }
    }
    var v56 int32 = -699
    _ = len(v1)
    v57 := GenS2{F0: true, F1: 4294967295, F2: 65535}
    v3 -= int8(v4)
    var v58 uint8 = 105
    for fi59 := 0; fi59 < 16; fi59++ {
        v60 := GenS8{F0: (v56 ^ (v56 / ((v56) | 1)))}
        if true {
            fmt.Println(v3, v5, v4)
            switch "" {
            case "a":
                fmt.Println(v4, v4)
            case "foo bar":
                if (!(!false)) {
                    var v61 uint64 = 664
                    _ = v61
                }
            case "hello":
                v58 = v5
            default:
                v3 = v3
            }
            switch "foo bar" {
            case "foo bar":
                var v62 uint = 4294967295
                _ = v62
            default:
                fmt.Println(v5)
            }
        }
        fmt.Println(fi59)
        _ = fi59
        _ = v60
    }
    var v63 uint8 = (v5 & v5)
    if (!(!false)) {
        fmt.Println(v5, v63, v58)
        fmt.Println(v3, v5, v3)
    }
    if (!(127 >= (v5 ^ v63))) {
        switch "hello" {
        case "foo bar":
            fmt.Println(v58)
        case "":
            var v64 float64 = -2.25
            _ = v64
        }
        switch ("" + ("hello" + "\n\t")) {
        case "":
            fmt.Println(v58, v3, v4)
        case "a":
            v4 = -268
        case "foo bar":
            for fi65 := 0; fi65 < 8; fi65++ {
                _ = len(v1)
                if false {
                    var v66 float64 = (0.0 + ((((3.14159 * 0.1) * 1.5) - -2.25) - 1e10))
                    _ = len(v2)
                    _ = v66
                } else {
                    for fi67 := 0; fi67 < 3; fi67++ {
                        switch "hello" {
                        case "":
                            var v68 int8 = (v3 | -88)
                            _ = v68
                        case "\n\t":
                            switch v4 {
                            case 480:
                                if (399 < 9223372036854775807) {
                                    v69 := GenS9{F0: 990, F1: 2147483648, F2: 1}
                                    _ = v69
                                }
                                if true {
                                    fmt.Println(v5)
                                }
                                var v70 int64 = -588
                                _ = v70
                            case 9223372036854775807:
                                delete(v2, 173)
                            default:
                                var v71 int64 = -9223372036854775808
                                _ = v71
                            }
                        default:
                            var v72 int64 = 0
                            _ = v72
                        }
                        _ = fi67
                    }
                    fmt.Println(v58, v56, v58)
                    var v73 uint = 328
                    _ = v73
                }
                _ = fi65
            }
        default:
            if (1 >= 4294967295) {
                if false {
                    v1 = append(v1, 65535)
                } else {
                    v3 = v3
                }
            } else {
                switch v4 {
                case 42:
                    v58 -= 38
                case -50:
                    if true {
                        fmt.Println(v56, v56, v4)
                        v4 -= (^(v4))
                        v74 := GenS8{F0: (v56 / ((-243) | 1))}
                        _ = v74
                    }
                }
                switch "" {
                case "\n\t":
                    v58 += (^(v5))
                case "":
                    var v75 int64 = 536
                    _ = v75
                case "foo bar":
                    v4 += (v4 - v4)
                default:
                    var v76 string = ("hello" + "foo bar")
                    _ = v76
                }
            }
        }
    } else {
        for fi77 := 0; fi77 < 7; fi77++ {
            v3 = -128
            _ = fi77
        }
        for fi78 := 0; fi78 < 15; fi78++ {
            for fi79 := 0; fi79 < 7; fi79++ {
                if ((-(v56)) >= -2147483648) {
                    v56 += 1
                }
                if (!((0.1 == 3.14159) && true)) {
                    v80 := GenS2{F0: (!true), F1: 147, F2: 1}
                    v81 := GenS10{F0: fi78, F1: (v3 / ((v3) | 1)), F2: 18446744073709551615}
                    switch ("a" + "") {
                    case "":
                        v58 = v5
                    case "foo bar":
                        var v82 int16 = 0
                        _ = v82
                    default:
                        fmt.Println(fi78)
                    }
                    _ = v80
                    _ = v81
                }
                _ = fi79
            }
            _ = fi78
        }
        for fi83 := 0; fi83 < 7; fi83++ {
            var v84 int16 = 1
            _ = fi83
            _ = v84
        }
    }
    switch "\n\t" {
    case "foo bar":
        v5 = 2
    }
    v85 := GenS11{F0: 0}
    var v86 float64 = 0.0
    var v87 int = (v4 & v4)
    var v88 string = "foo bar"
    v89 := GenS7{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("v63:", v63)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v63
    _ = v85
    _ = v86
    _ = v87
    _ = v88
    _ = v89
}
