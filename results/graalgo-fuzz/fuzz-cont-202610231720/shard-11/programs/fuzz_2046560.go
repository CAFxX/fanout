package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 uint32; F2 int }
type GenS7 struct { F0 int8; F1 bool; F2 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int32; F1 int8; F2 string }
func genFn0(p0 uint16) int32 {
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := []int{(fi5 / (((fi5 & 9223372036854775807)) | 1))}
        p0 += 278
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 863, F1: 4294967295, F2: -1}
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 uint = 18446744073709551615
            v11 := []int{fi9, (fi8 & 1), (fi9 << (((v10 - v10)) & 63))}
            _ = fi9
            _ = v10
            _ = v11
        }
        p0 = p0
        _ = fi8
    }
    v12 := GenS2{F0: false}
    _ = p0
    _ = v12
    _ = v7
    return -2147483648
}
func genFn1(p0 int16) uint {
    v13 := []int{675, 674, 1, 128}
    var v14 int32 = 0
    _ = len(v13)
    var v15 int64 = 0
    var v16 int32 = (v14 * v14)
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    return 294
}

func main() {
    var v1 bool = (-61 >= 86)
    var v2 float64 = (float64(1.5) + 3.14159)
    var v3 float64 = 1e10
    var v4 uint64 = 391
    v17 := []int{42, -962}
    if true {
        if (42 <= 2147483647) {
            fmt.Println(v1)
        } else {
            var v18 uint8 = 127
            var v19 float32 = (-2.25 * -2.25)
            var v20 int64 = 1
            _ = v18
            _ = v19
            _ = v20
        }
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
    }
    v1 = (!v1)
    v21 := GenS3{F0: 2147483647}
    if (-64 == 0) {
        v4 += v4
        switch "foo bar" {
        case "hello":
            for fi22 := 0; fi22 < 11; fi22++ {
                fmt.Println(v2, v4)
                v2 += (0.0 + ((v2 - -2.25) * 0.1))
                _ = fi22
            }
        case "foo bar":
            if len(v17) > 0 {
                v17[0] = -1
            }
        case "":
            switch "hello" {
            case "\n\t":
                var v23 float32 = 0.0
                _ = v23
            default:
                for fi24 := 0; fi24 < 12; fi24++ {
                    for fi25 := 0; fi25 < 8; fi25++ {
                        if (1 >= uint16(v4)) {
                            fmt.Println(v1)
                            switch "hello" {
                            case "hello":
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                            }
                            if (!false) {
                                fmt.Println(v2, v1)
                                var v26 int = (fi25 + fi24)
                                v2 = (v3 + 3.14159)
                                _ = v26
                            } else {
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                            }
                            switch (fi25 << ((1) & 63)) {
                            case 9223372036854775807:
                                var v27 bool = (-907 <= -624)
                                _ = v27
                            default:
                                var v28 bool = v1
                                _ = v28
                            }
                        } else {
                            if true {
                                v29 := GenS1{F0: 4294967295, F1: 1, F2: fi25}
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                v3 = 0.0
                                _ = v29
                            } else {
                                if v1 {
                                    if (1 != 42) {
                                        var v30 uint16 = 1
                                        fmt.Println(v4, v4)
                                        fmt.Println(v30, v1, fi25)
                                        if ((v4 % ((2) | 1)) == (v4 ^ v4)) {
                                            if true {
                                                _ = len(v17)
                                                if len(v17) >= 2 {
                                                    v17 = v17[1:len(v17)]
                                                }
                                                v31 := GenS4{F0: -1, F1: v30}
                                                _ = v31
                                            } else {
                                                var v32 uint8 = 42
                                                _ = v32
                                            }
                                            _ = len(v17)
                                            var v33 int32 = 414
                                            _ = v33
                                        }
                                        _ = v30
                                    } else {
                                        v4 = (v4 % ((18446744073709551615) | 1))
                                    }
                                    var v34 uint8 = 148
                                    _ = v34
                                } else {
                                    switch 191 {
                                    case -880:
                                        var v35 uint8 = 2
                                        _ = v35
                                    case 128:
                                        v36 := GenS3{F0: fi24}
                                        v4 = 18446744073709551615
                                        v3 = v3
                                        _ = v36
                                    case 32767:
                                        v17 = append(v17, 486)
                                    }
                                    var v37 uint = 42
                                    var v38 int64 = -159
                                    _ = v37
                                    _ = v38
                                }
                            }
                            fmt.Println(v1)
                            var v39 uint = 255
                            _ = v39
                        }
                        v40 := GenS5{F0: 9223372036854775807}
                        v41 := GenS2{F0: (65535 == 2)}
                        fmt.Println(v2)
                        _ = fi25
                        _ = v40
                        _ = v41
                    }
                    for fi42 := 0; fi42 < 2; fi42++ {
                        switch "foo bar" {
                        case "\n\t":
                            v4 = v4
                        default:
                            if v1 {
                                v4 = v4
                                if v1 {
                                    v1 = true
                                    var v43 float64 = v2
                                    _ = v43
                                } else {
                                    var v44 int = fi42
                                    v45 := GenS1{F0: 2, F1: 805, F2: (fi24 + 9223372036854775807)}
                                    switch (("a" + ("a" + "foo bar")) + ("\n\t" + (("" + ("hello" + "hello")) + "foo bar"))) {
                                    case "a":
                                        _ = len(v17)
                                    case "foo bar":
                                        v46 := GenS5{F0: 9223372036854775807}
                                        _ = v46
                                    }
                                    v47 := GenS6{F0: ((v4 - (^(v4))) > v4), F1: 42, F2: fi42}
                                    if v1 {
                                        var v48 uint = 568
                                        _ = v48
                                    }
                                    _ = v44
                                    _ = v45
                                    _ = v47
                                }
                                v2 += float64(3.14159)
                                fmt.Println(fi24, v3, v1)
                                if ((v2 - v2) >= float64(0.1)) {
                                    if (65535 >= 2) {
                                        v2 = v2
                                        if (true && false) {
                                            _ = len(v17)
                                            if ((704 >= 65535) || v1) {
                                                v2 = 0.0
                                                v49 := GenS1{F0: 1, F1: 42, F2: 816}
                                                v50 := GenS7{F0: -7, F1: (0 < -1), F2: 18446744073709551615}
                                                _ = v49
                                                _ = v50
                                            } else {
                                                if len(v17) > 0 {
                                                    v17[0] = 1
                                                }
                                            }
                                            var v51 uint = 390
                                            _ = v51
                                        }
                                        var v52 uint32 = 2147483647
                                        _ = v52
                                    } else {
                                        fmt.Println(fi42, v3, fi24)
                                        v3 = v3
                                    }
                                    var v53 uint64 = (v4 - (v4 / (((v4 - 2)) | 1)))
                                    _ = len(v17)
                                    var v54 int = 32768
                                    _ = v53
                                    _ = v54
                                }
                            } else {
                                v3 = float64(3.14159)
                                v1 = true
                                _ = len(v17)
                            }
                            var v55 uint = 42
                            _ = v55
                        }
                        var v56 uint = 0
                        _ = len(v17)
                        _ = fi42
                        _ = v56
                    }
                    _ = fi24
                }
            }
        default:
            for fi57 := 0; fi57 < 20; fi57++ {
                var v58 int16 = 1000
                _ = fi57
                _ = v58
            }
        }
    }
    switch "\n\t" {
    case "foo bar":
        for fi59 := 0; fi59 < 10; fi59++ {
            for fi60 := 0; fi60 < 16; fi60++ {
                v61 := GenS1{F0: 4294967295, F1: 2147483647, F2: (fi59 & (fi60 << ((65535) & 63)))}
                fmt.Println(fi60, v4, v2)
                _ = fi60
                _ = v61
            }
            _ = fi59
        }
    case "":
        var v62 int64 = 9223372036854775807
        _ = v62
    default:
        var v63 int16 = 1
        _ = v63
    }
    var v64 uint32 = 577
    for fi65 := 0; fi65 < 12; fi65++ {
        v1 = v1
        v66 := GenS4{F0: 646, F1: 255}
        _ = fi65
        _ = v66
    }
    var v67 int64 = 0
    v2 = 0.1
    if (v67 > (^(v67))) {
        if (!((v67 + v67) != -9223372036854775808)) {
            fmt.Println(v2, v1, v67)
            v1 = v1
            for fi68 := 0; fi68 < 12; fi68++ {
                if true {
                    switch (fi68 >> ((843) & 63)) {
                    case 0:
                        switch fi68 {
                        case 121:
                            v69 := GenS6{F0: (1 <= 636), F1: v64, F2: fi68}
                            _ = v69
                        }
                    case 746:
                        v70 := GenS8{F0: 351}
                        _ = v70
                    case -523:
                        if len(v17) > 0 {
                            v17[0] = (fi68 & fi68)
                        }
                    }
                    if v1 {
                        v1 = ((v1 || v1) || false)
                    } else {
                        fmt.Println(v2, v2, v67)
                        if (0 != 65535) {
                            var v71 uint16 = 1
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            _ = v71
                        }
                    }
                    v72 := GenS9{F0: 0, F1: -6, F2: "a"}
                    _ = v72
                } else {
                    switch (("hello" + ("hello" + ("a" + "a"))) + ("foo bar" + "a")) {
                    case "hello":
                        if ((v67 - (v67 / ((v67) | 1))) < (v67 ^ v67)) {
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            var v73 int8 = 42
                            var v74 uint8 = 70
                            _ = v73
                            _ = v74
                        }
                    case "\n\t":
                        fmt.Println(v64)
                    case "a":
                        v1 = false
                    }
                }
                _ = fi68
            }
        }
    }
    var v75 int = 71
    fmt.Println(v64, v64, v64)
    switch v75 {
    case 1:
        switch ("" + "") {
        case "foo bar":
            v76 := GenS4{F0: 32767, F1: 255}
            _ = v76
        default:
            v77 := GenS3{F0: v75}
            _ = v77
        }
    case 597:
        switch v75 {
        case 115:
            if true {
                var v78 int8 = -23
                switch ("" + ("hello" + "a")) {
                case "a":
                    switch "foo bar" {
                    case "hello":
                        v17 = append(v17, v75)
                    case "a":
                        var v79 string = ("foo bar" + "a")
                        _ = v79
                    default:
                        var v80 uint64 = v4
                        _ = v80
                    }
                case "foo bar":
                    fmt.Println(v4, v78, v67)
                case "\n\t":
                    if (("foo bar" + "foo bar") < "hello") {
                        var v81 float32 = ((-2.25 / -2.25) - 3.14159)
                        _ = v81
                    }
                default:
                    v75 = v75
                }
                fmt.Println(v78)
                _ = v78
            }
        case 1:
            for fi82 := 0; fi82 < 9; fi82++ {
                var v83 int16 = 1000
                _ = fi82
                _ = v83
            }
        default:
            if (42 <= 10) {
                v4 -= 670
                var v84 int32 = 1
                v64 = v64
                _ = v84
            }
        }
    case 0:
        v1 = (127 == 1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v64:", v64)
    fmt.Println("v67:", v67)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v64
    _ = v67
    _ = v75
}
