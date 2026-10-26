package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int8; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 uint32; F2 int }
type GenS7 struct { F0 uint; F1 uint; F2 int8 }
type GenS8 struct { F0 int64; F1 int64; F2 uint32 }
type GenS9 struct { F0 uint8; F1 uint8 }
func genFn0() int64 {
    if (32767 != -1) {
        for fi6 := 0; fi6 < 12; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                v8 := GenS1{F0: ((0.0 / float64(1e10)) > 1.5)}
                v9 := []int{fi7}
                if (-212 != 2147483647) {
                    var v10 bool = ((42 < 42) || (!(false || true)))
                    v11 := GenS1{F0: true}
                    _ = v10
                    _ = v11
                } else {
                    switch (fi6 | (fi6 << ((4294967295) & 63))) {
                    case -114:
                        switch (("" + "hello") + "foo bar") {
                        case "hello":
                            var v12 uint8 = 163
                            v12 -= 42
                            _ = v12
                        case "a":
                            var v13 bool = false
                            _ = v13
                        case "foo bar":
                            var v14 float64 = 1.5
                            _ = v14
                        }
                    case 32768:
                        var v15 uint16 = 652
                        _ = v15
                    case 0:
                        if ((0.1 >= float64(3.14159)) || (0 != 1)) {
                            var v16 uint8 = 1
                            _ = v16
                        } else {
                            v9 = append(v9, -2147483648)
                            fmt.Println(fi7)
                            var v17 int = 9223372036854775807
                            _ = v17
                        }
                    }
                    fmt.Println(fi6, fi6)
                }
                var v18 string = ("a" + ("" + ("hello" + "hello")))
                v19 := GenS1{F0: false}
                _ = fi7
                _ = v18
                _ = v19
                _ = v8
                _ = v9
            }
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 int64 = 9223372036854775807
                _ = fi20
                _ = v21
            }
            switch fi6 {
            case 960:
                for fi22 := 0; fi22 < 1; fi22++ {
                    var v23 string = ("hello" + "a")
                    if false {
                        switch (fi6 | fi6) {
                        case 128:
                            fmt.Println(fi6, fi6, fi6)
                        default:
                            v23 = v23
                        }
                        fmt.Println(v23)
                        v24 := []int{}
                        _ = v24
                    }
                    v23 = ("hello" + ("a" + v23))
                    switch fi22 {
                    case 956:
                        v25 := GenS1{F0: (fi22 >= fi6)}
                        _ = v25
                    }
                    v26 := GenS1{F0: (!false)}
                    _ = fi22
                    _ = v23
                    _ = v26
                }
            case 1:
                v27 := GenS2{F0: 0}
                _ = v27
            default:
                v28 := []int{}
                _ = v28
            }
            _ = fi6
        }
        if true {
            if ("hello" == "\n\t") {
                var v29 uint16 = 65535
                v30 := []int{}
                v30 = append(v30, 127)
                _ = v29
                _ = v30
            }
            for fi31 := 0; fi31 < 9; fi31++ {
                fmt.Println(fi31)
                v32 := GenS3{F0: 42, F1: -128, F2: "\n\t"}
                if false {
                    var v33 uint16 = 421
                    v34 := []int{816}
                    if (!true) {
                        if true {
                            switch "foo bar" {
                            case "foo bar":
                                switch (fi31 << ((42) & 63)) {
                                case 32767:
                                    _ = len(v34)
                                case -1:
                                    fmt.Println(v33, fi31)
                                }
                            case "hello":
                                if false {
                                    for fi35 := 0; fi35 < 16; fi35++ {
                                        fmt.Println(fi35, v33, fi31)
                                        _ = fi35
                                    }
                                }
                            case "\n\t":
                                if (4294967295 != 0) {
                                    fmt.Println(v33, fi31)
                                    v33 = (^(v33))
                                    v36 := GenS4{F0: 402}
                                    _ = v36
                                }
                            default:
                                if true {
                                    if ("foo bar" != ("a" + "")) {
                                        _ = len(v34)
                                    }
                                    var v37 int16 = -32768
                                    var v38 uint32 = 328
                                    _ = v37
                                    _ = v38
                                } else {
                                    switch fi31 {
                                    case 78:
                                        for fi39 := 0; fi39 < 18; fi39++ {
                                            v33 = v33
                                            _ = fi39
                                        }
                                    case 128:
                                        for fi40 := 0; fi40 < 1; fi40++ {
                                            _ = len(v34)
                                            if (!(float32((0.0 * 0.1)) <= 0.1)) {
                                                var v41 int32 = 2147483647
                                                v34 = append(v34, 215)
                                                v33 = 1
                                                _ = v41
                                            } else {
                                                var v42 uint64 = 42
                                                fmt.Println(fi40, fi40, fi40)
                                                switch (fi31 % (((fi40 * (fi40 * -9223372036854775808))) | 1)) {
                                                case 0:
                                                    v42 += (v42 / (((v42 * 42)) | 1))
                                                case -1:
                                                    v43 := GenS1{F0: (uint32(v33) != 1)}
                                                    _ = v43
                                                default:
                                                    v44 := GenS1{F0: true}
                                                    _ = v44
                                                }
                                                _ = v42
                                            }
                                            _ = len(v34)
                                            _ = fi40
                                        }
                                    default:
                                        for fi45 := 0; fi45 < 12; fi45++ {
                                            v46 := GenS5{F0: (("" + ("hello" + "hello")) < "foo bar")}
                                            v47 := GenS3{F0: -765, F1: int8(v33), F2: "foo bar"}
                                            _ = fi45
                                            _ = v46
                                            _ = v47
                                        }
                                    }
                                    for fi48 := 0; fi48 < 13; fi48++ {
                                        v33 = (v33 + v33)
                                        _ = fi48
                                    }
                                }
                            }
                            for fi49 := 0; fi49 < 20; fi49++ {
                                v50 := GenS1{F0: (false && false)}
                                _ = fi49
                                _ = v50
                            }
                        } else {
                            v33 += (v33 | 255)
                            var v51 bool = (128 == 0)
                            _ = v51
                        }
                    }
                    _ = v33
                    _ = v34
                } else {
                    v52 := []int{(^(fi31)), (fi31 * (-(fi31))), fi31}
                    _ = v52
                }
                _ = fi31
                _ = v32
            }
            var v53 uint16 = 0
            _ = v53
        }
    } else {
        var v54 uint8 = 205
        _ = v54
    }
    var v55 uint8 = 164
    v56 := GenS2{F0: 1}
    var v57 int8 = 83
    _ = v55
    _ = v56
    _ = v57
    return 0
}
func genFn1(p0 int16, p1 uint, p2 uint8) uint64 {
    for fi58 := 0; fi58 < 10; fi58++ {
        var v59 uint8 = p2
        _ = fi58
        _ = v59
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 uint64 = 647
    v2 := make(map[int]uint32)
    var v3 float64 = 3.14159
    var v4 uint16 = 986
    var v5 string = "a"
    fmt.Println(v1, v5)
    var v60 float64 = v3
    var v61 int16 = 1000
    fmt.Println(v1)
    var v62 string = v5
    v63 := []int{128, 0, 1, 2147483647}
    v4 = (v4 / ((2) | 1))
    for fi64 := 0; fi64 < 7; fi64++ {
        var v65 uint16 = (v4 % ((v4) | 1))
        _ = fi64
        _ = v65
    }
    for fi66 := 0; fi66 < 8; fi66++ {
        v60 = 0.0
        _ = fi66
    }
    var v67 int8 = 1
    if true {
        switch -2147483648 {
        case 306:
            var v68 int16 = 1000
            _ = v68
        }
        var v69 int16 = v61
        _ = v69
    }
    var v70 int32 = -1
    for fi71 := 0; fi71 < 9; fi71++ {
        for fi72 := 0; fi72 < 4; fi72++ {
            var v73 float32 = float32(v3)
            if true {
                v67 -= (v67 << ((876) & 63))
                var v74 uint8 = 42
                _ = v74
            }
            _ = fi72
            _ = v73
        }
        for fi75 := 0; fi75 < 12; fi75++ {
            if (!(127 <= 128)) {
                if false {
                    v76 := GenS6{F0: 18446744073709551615, F1: 861, F2: int(v1)}
                    _ = v76
                }
            }
            v63 = append(v63, (fi75 * fi75))
            _ = fi75
        }
        switch v5 {
        case "":
            switch (fi71 - -403) {
            case 9223372036854775807:
                v67 = v67
            case 127:
                fmt.Println(v5)
            default:
                if false {
                    for fi77 := 0; fi77 < 6; fi77++ {
                        if false {
                            var v78 float32 = (1.5 + float32(v3))
                            v4 = uint16(v1)
                            _ = v78
                        } else {
                            v79 := GenS7{F0: 2, F1: 4294967295, F2: v67}
                            _ = v79
                        }
                        v80 := GenS2{F0: 127}
                        _ = fi77
                        _ = v80
                    }
                    if (191 == 127) {
                        for fi81 := 0; fi81 < 10; fi81++ {
                            switch fi71 {
                            case -537:
                                switch "hello" {
                                case "\n\t":
                                    if true {
                                        var v82 int64 = -1
                                        var v83 float64 = 0.0
                                        switch v5 {
                                        case "a":
                                            var v84 bool = (fi71 <= int(v70))
                                            _ = v84
                                        }
                                        _ = v82
                                        _ = v83
                                    }
                                    var v85 float64 = v3
                                    _ = v85
                                case "":
                                    delete(v2, (-(fi71)))
                                }
                            }
                            var v86 uint8 = uint8(v1)
                            fmt.Println(v67, v1, v70)
                            _ = fi81
                            _ = v86
                        }
                    }
                    if false {
                        v87 := GenS8{F0: -9223372036854775808, F1: 239, F2: 65535}
                        _ = v87
                    } else {
                        v88 := v2[int(v70)]
                        if (3.14159 <= float64((float32((v60 + -2.25)) + -2.25))) {
                            fmt.Println(v62, v61)
                            v2[fi71] = v88
                            for fi89 := 0; fi89 < 3; fi89++ {
                                v63 = append(v63, fi89)
                                fmt.Println(v62)
                                var v90 uint = 988
                                _ = fi89
                                _ = v90
                            }
                        }
                        _ = v88
                    }
                } else {
                    switch 702 {
                    case -9223372036854775808:
                        var v91 int16 = (v61 + 1000)
                        _ = v91
                    }
                }
            }
        case "a":
            for fi92 := 0; fi92 < 1; fi92++ {
                v93 := GenS9{F0: uint8(v4), F1: 255}
                v94 := GenS8{F0: 9223372036854775807, F1: -1, F2: 4294967295}
                _ = fi92
                _ = v93
                _ = v94
            }
        case "foo bar":
            v67 = 127
        }
        _ = fi71
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("slicelen v63:", len(v63), v63)
    fmt.Println("v67:", v67)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v61
    _ = v62
    _ = v63
    _ = v67
    _ = v70
}
