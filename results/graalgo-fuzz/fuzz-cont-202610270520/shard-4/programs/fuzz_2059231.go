package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int8; F1 int8; F2 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 int8; F1 uint; F2 int16 }
type GenS9 struct { F0 uint8; F1 uint16; F2 uint }
type GenS10 struct { F0 int; F1 int8 }
type GenS11 struct { F0 int32; F1 bool }
type GenS12 struct { F0 int8 }

func main() {
    var v1 bool = true
    var v2 int32 = 2147483647
    var v3 int8 = 42
    var v4 int64 = -9223372036854775808
    var v5 uint8 = 2
    for fi6 := 0; fi6 < 19; fi6++ {
        v5 = 102
        _ = fi6
    }
    if true {
        v7 := GenS1{F0: 745}
        _ = v7
    } else {
        fmt.Println(v4, v2)
        v5 = (v5 | (v5 & (v5 + v5)))
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        if true {
            var v9 uint16 = 327
            _ = v9
        } else {
            var v10 uint = 0
            switch "" {
            case "":
                if (!v1) {
                    if v1 {
                        if false {
                            v11 := []int{(fi8 | fi8)}
                            _ = v11
                        } else {
                            v12 := GenS1{F0: 2}
                            var v13 int16 = 1
                            switch (fi8 - fi8) {
                            case 205:
                                var v14 float64 = 0.0
                                _ = v14
                            case -627:
                                v2 -= (v2 / ((v2) | 1))
                            case 128:
                                if (42 > (v4 ^ v4)) {
                                    var v15 int16 = v13
                                    v16 := []int{-624, (fi8 / (((fi8 % ((fi8) | 1))) | 1)), fi8}
                                    for fi17 := 0; fi17 < 6; fi17++ {
                                        switch "foo bar" {
                                        case "hello":
                                            if (false || (float64((1.5 / ((0.1 + 0.0) / 1e10))) >= 0.1)) {
                                                var v18 float64 = (float64(0.1) * float64(1e10))
                                                v19 := GenS2{F0: (^(v4)), F1: 1}
                                                _ = v18
                                                _ = v19
                                            }
                                        case "\n\t":
                                            fmt.Println(v5, v2, v2)
                                        case "":
                                            if (!v1) {
                                                switch "a" {
                                                case "hello":
                                                    if ("hello" < "a") {
                                                        switch fi8 {
                                                        case 4294967295:
                                                            var v20 uint16 = 65535
                                                            _ = v20
                                                        case 581:
                                                            fmt.Println(v15, v10, v2)
                                                        }
                                                        if v1 {
                                                            if (65535 > uint16(v5)) {
                                                                if v1 {
                                                                    var v21 int32 = -821
                                                                    _ = v21
                                                                } else {
                                                                    v4 = 0
                                                                    var v22 uint32 = 561
                                                                    if len(v16) > 0 {
                                                                        v16[0] = (^(fi17))
                                                                    }
                                                                    v22 -= 790
                                                                    _ = v22
                                                                }
                                                            } else {
                                                                if (false || false) {
                                                                    v2 += v2
                                                                    var v23 int32 = (v2 - (v2 >> (((^(v10))) & 63)))
                                                                    if ((1 != 0) && v1) {
                                                                        if (!((!("hello" > "foo bar")) && (!v1))) {
                                                                            switch 976 {
                                                                            case -2147483648:
                                                                                v24 := GenS2{F0: 9223372036854775807, F1: 255}
                                                                                _ = v24
                                                                            case 707:
                                                                                v3 = v3
                                                                            case 0:
                                                                                var v25 int32 = (v23 & (v23 % (((v2 & (v23 * (v23 | (v23 + v23))))) | 1)))
                                                                                if (!("hello" != "foo bar")) {
                                                                                    _ = len(v16)
                                                                                } else {
                                                                                    v26 := GenS2{F0: 42, F1: 0}
                                                                                    v27 := GenS3{F0: 646, F1: (v23 >= (v23 & (v25 | (v2 - 42))))}
                                                                                    _ = v26
                                                                                    _ = v27
                                                                                }
                                                                                _ = v25
                                                                            default:
                                                                                if (-9223372036854775808 >= (v4 * v4)) {
                                                                                    if ((-(v15)) > v15) {
                                                                                        if false {
                                                                                            var v28 float64 = float64(-2.25)
                                                                                            var v29 uint = (v10 + 0)
                                                                                            _ = v28
                                                                                            _ = v29
                                                                                        } else {
                                                                                            v13 = 833
                                                                                        }
                                                                                        if (!v1) {
                                                                                            if len(v16) > 0 {
                                                                                                v16[0] = (fi17 % ((fi17) | 1))
                                                                                            }
                                                                                            v13 += v13
                                                                                        } else {
                                                                                            var v30 int16 = v13
                                                                                            v15 = 0
                                                                                            if len(v16) >= 2 {
                                                                                                v16 = v16[1:len(v16)]
                                                                                            }
                                                                                            _ = v30
                                                                                        }
                                                                                    } else {
                                                                                        var v31 int32 = (v2 ^ (v23 - -2147483648))
                                                                                        _ = v31
                                                                                    }
                                                                                    var v32 int64 = v4
                                                                                    v4 = v32
                                                                                    _ = v32
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v1 = (false || (1.5 == 1.5))
                                                                        if v1 {
                                                                            fmt.Println(fi17, v5, v10)
                                                                            if (-769 < (v2 & (v23 / (((^(v2))) | 1)))) {
                                                                                var v33 string = ""
                                                                                _ = v33
                                                                            } else {
                                                                                fmt.Println(v13, v23, v10)
                                                                                v1 = v1
                                                                            }
                                                                        } else {
                                                                            v34 := GenS1{F0: 712}
                                                                            v13 -= v15
                                                                            switch (fi8 - fi8) {
                                                                            case 976:
                                                                                var v35 uint64 = 0
                                                                                _ = v35
                                                                            case 371:
                                                                                if (1 >= 179) {
                                                                                    var v36 uint32 = 190
                                                                                    _ = v36
                                                                                } else {
                                                                                    v10 = 487
                                                                                    fmt.Println(v1, v23, fi8)
                                                                                    v37 := GenS4{F0: v3, F1: v3, F2: int32(v13)}
                                                                                    _ = v37
                                                                                }
                                                                            case 9223372036854775807:
                                                                                fmt.Println(v13)
                                                                            default:
                                                                                v3 = v3
                                                                            }
                                                                            _ = v34
                                                                        }
                                                                    }
                                                                    _ = v23
                                                                }
                                                                if v1 {
                                                                    v4 -= (^(v4))
                                                                    var v38 uint = (v10 / ((255) | 1))
                                                                    switch ("\n\t" + "hello") {
                                                                    case "a":
                                                                        v1 = v1
                                                                    case "hello":
                                                                        v4 = v4
                                                                    }
                                                                    _ = v38
                                                                }
                                                            }
                                                            var v39 uint8 = (v5 % ((v5) | 1))
                                                            v3 += 42
                                                            if v1 {
                                                                switch (fi17 | (fi8 % (((-(fi8))) | 1))) {
                                                                case 32768:
                                                                    var v40 string = "a"
                                                                    _ = v40
                                                                case 9223372036854775807:
                                                                    var v41 uint = v10
                                                                    _ = v41
                                                                }
                                                                v42 := GenS5{F0: v2}
                                                                switch ("\n\t" + "foo bar") {
                                                                case "":
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        v43 := GenS3{F0: v4, F1: false}
                                                                        _ = v43
                                                                    case "foo bar":
                                                                        v44 := GenS6{F0: 65535, F1: 1}
                                                                        _ = v44
                                                                    default:
                                                                        fmt.Println(v1)
                                                                    }
                                                                case "\n\t":
                                                                    var v45 string = "\n\t"
                                                                    _ = v45
                                                                }
                                                                _ = v42
                                                            }
                                                            _ = v39
                                                        } else {
                                                            if false {
                                                                var v46 uint16 = uint16(v5)
                                                                var v47 int32 = 1
                                                                if v1 {
                                                                    v15 += (v15 - int16(fi8))
                                                                    switch fi17 {
                                                                    case 4294967295:
                                                                        v16 = append(v16, 32767)
                                                                    case -531:
                                                                        v47 -= v47
                                                                        fmt.Println(v1)
                                                                        var v48 int8 = (v3 + (v3 >> ((uint(v5)) & 63)))
                                                                        fmt.Println(v15)
                                                                        v49 := GenS7{F0: 0, F1: (((("hello" + "a") + "\n\t") + "hello") + "hello")}
                                                                        v50 := GenS8{F0: int8(fi8), F1: 0, F2: v15}
                                                                        _ = v48
                                                                        _ = v49
                                                                        _ = v50
                                                                    }
                                                                    v51 := GenS9{F0: v5, F1: (v46 >> ((v10) & 63)), F2: 42}
                                                                    _ = v51
                                                                }
                                                                v47 -= (-(v2))
                                                                switch "foo bar" {
                                                                case "a":
                                                                    if v1 {
                                                                        v3 += v3
                                                                        v52 := GenS5{F0: v2}
                                                                        if v1 {
                                                                            v2 = (v47 << (((v10 & (v10 % (((v10 | v10)) | 1)))) & 63))
                                                                            var v53 uint16 = (v46 / ((0) | 1))
                                                                            if (fi8 != fi17) {
                                                                                if len(v16) >= 2 {
                                                                                    v16 = v16[1:len(v16)]
                                                                                }
                                                                                fmt.Println(v4)
                                                                            } else {
                                                                                var v54 uint64 = 0
                                                                                v55 := GenS10{F0: int(v3), F1: 1}
                                                                                _ = v54
                                                                                _ = v55
                                                                            }
                                                                            _ = v53
                                                                        } else {
                                                                            if len(v16) >= 2 {
                                                                                v16 = v16[1:len(v16)]
                                                                            }
                                                                        }
                                                                        _ = v52
                                                                    } else {
                                                                        var v56 int64 = (v4 << ((673) & 63))
                                                                        v1 = (!true)
                                                                        var v57 int = -9223372036854775808
                                                                        _ = v56
                                                                        _ = v57
                                                                    }
                                                                case "\n\t":
                                                                    v16 = append(v16, (^(fi17)))
                                                                default:
                                                                    v10 = (v10 / ((v10) | 1))
                                                                }
                                                                _ = v46
                                                                _ = v47
                                                            }
                                                        }
                                                    }
                                                case "a":
                                                    fmt.Println(fi17, v13)
                                                }
                                                v58 := GenS7{F0: 47, F1: ""}
                                                if len(v16) >= 2 {
                                                    v16 = v16[1:len(v16)]
                                                }
                                                _ = v58
                                            } else {
                                                switch (fi8 << ((0) & 63)) {
                                                case -9223372036854775808:
                                                    var v59 uint16 = 870
                                                    _ = v59
                                                case 394:
                                                    v60 := GenS3{F0: (v4 + v4), F1: false}
                                                    _ = v60
                                                default:
                                                    var v61 int = (fi8 | (fi8 | int(v15)))
                                                    _ = v61
                                                }
                                                fmt.Println(v13, v13)
                                            }
                                        }
                                        _ = fi17
                                    }
                                    _ = v15
                                    _ = v16
                                }
                            }
                            _ = v12
                            _ = v13
                        }
                        fmt.Println(v5, v1)
                    } else {
                        if true {
                            v2 += 0
                        }
                        fmt.Println(fi8, v10)
                        for fi62 := 0; fi62 < 9; fi62++ {
                            v63 := []int{(^(fi62)), fi62}
                            var v64 int = 0
                            v63 = append(v63, (v64 & fi62))
                            _ = fi62
                            _ = v63
                            _ = v64
                        }
                    }
                    v1 = ((v4 & v4) < v4)
                    if ((v5 == v5) || v1) {
                        switch "hello" {
                        case "\n\t":
                            v2 = (^(v2))
                        case "hello":
                            if (3.14159 > (0.0 + (3.14159 / (-2.25 + (-2.25 * (1.5 * 0.1)))))) {
                                v2 += (v2 & 796)
                                for fi65 := 0; fi65 < 11; fi65++ {
                                    var v66 uint8 = v5
                                    if v1 {
                                        v3 = v3
                                        if v1 {
                                            fmt.Println(fi65)
                                            var v67 uint16 = 171
                                            _ = v67
                                        } else {
                                            var v68 float64 = 0.1
                                            var v69 int8 = v3
                                            var v70 int8 = v69
                                            _ = v68
                                            _ = v69
                                            _ = v70
                                        }
                                    }
                                    _ = fi65
                                    _ = v66
                                }
                                if true {
                                    v2 = (v2 / (((v2 * v2)) | 1))
                                    v71 := GenS11{F0: 0, F1: v1}
                                    for fi72 := 0; fi72 < 15; fi72++ {
                                        var v73 int = (^(fi8))
                                        if v1 {
                                            v74 := GenS6{F0: 65535, F1: 0}
                                            if v1 {
                                                v3 = (v3 >> ((2) & 63))
                                                v1 = v1
                                            }
                                            _ = v74
                                        }
                                        _ = fi72
                                        _ = v73
                                    }
                                    _ = v71
                                }
                            } else {
                                v75 := GenS9{F0: (^(v5)), F1: 32767, F2: v10}
                                fmt.Println(v5, fi8)
                                v76 := []int{127, fi8}
                                _ = v75
                                _ = v76
                            }
                        default:
                            v3 = v3
                        }
                        if v1 {
                            v2 = v2
                        } else {
                            var v77 uint = (v10 & (^(v10)))
                            v2 += v2
                            _ = v77
                        }
                        v3 += v3
                    } else {
                        v78 := []int{-761, (^(fi8))}
                        _ = v78
                    }
                }
            default:
                v1 = ("a" != ("" + ("foo bar" + "a")))
            }
            var v79 uint64 = 2
            _ = v10
            _ = v79
        }
        _ = fi8
    }
    var v80 float64 = (0.0 * 0.1)
    for fi81 := 0; fi81 < 5; fi81++ {
        var v82 int32 = (^(v2))
        var v83 int16 = 275
        _ = fi81
        _ = v82
        _ = v83
    }
    if ((1 > 255) || (((float64(-2.25) * v80) - (v80 * v80)) <= v80)) {
        v3 -= -33
    } else {
        v80 += (v80 / v80)
        fmt.Println(v4, v2)
    }
    fmt.Println(v4)
    for fi84 := 0; fi84 < 2; fi84++ {
        var v85 uint8 = (v5 | v5)
        var v86 uint64 = 4294967295
        _ = fi84
        _ = v85
        _ = v86
    }
    for fi87 := 0; fi87 < 11; fi87++ {
        v2 += (v2 * v2)
        _ = fi87
    }
    var v88 bool = false
    fmt.Println(v5)
    v88 = ((v2 >> ((18446744073709551615) & 63)) < (-(v2)))
    var v89 int64 = v4
    v90 := []int{int(v2), -2147483648, int(v89)}
    if v88 {
        if ((!v1) && (!false)) {
            v91 := GenS12{F0: 42}
            switch "" {
            case "hello":
                for fi92 := 0; fi92 < 3; fi92++ {
                    var v93 int32 = v2
                    fmt.Println(v4, v93, v2)
                    _ = fi92
                    _ = v93
                }
            default:
                if len(v90) >= 2 {
                    v90 = v90[1:len(v90)]
                }
            }
            _ = v91
        }
    }
    v90 = append(v90, 9223372036854775807)
    if ((v3 << ((42) & 63)) <= v3) {
        var v94 uint32 = 2
        _ = v94
    }
    if len(v90) >= 2 {
        v90 = v90[1:len(v90)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("slicelen v90:", len(v90), v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v80
    _ = v88
    _ = v89
    _ = v90
}
