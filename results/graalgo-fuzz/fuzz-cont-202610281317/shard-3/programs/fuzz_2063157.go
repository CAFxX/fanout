package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16; F1 uint16; F2 string }
type GenS9 struct { F0 uint16; F1 int64; F2 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 42
    v3 := make(map[int]bool)
    var v4 uint16 = v2
    v5 := []int{4294967295, -222, 9223372036854775807}
    v5 = append(v5, 128)
    if true {
        fmt.Println(v4, v4)
    }
    v4 = 928
    for fi6 := 0; fi6 < 19; fi6++ {
        v4 += (v2 >> ((42) & 63))
        _ = fi6
    }
    v3[32767] = false
    for fi7 := 0; fi7 < 2; fi7++ {
        if (928 <= 42) {
            v8 := GenS1{F0: (218 <= 65535), F1: 0, F2: 1000}
            for fi9 := 0; fi9 < 2; fi9++ {
                v2 += (v2 + (v4 % (((v4 + (v2 - v4))) | 1)))
                _ = fi9
            }
            _ = v8
        }
        v5 = append(v5, (fi7 + 128))
        var v10 int32 = -171
        _ = fi7
        _ = v10
    }
    switch -1 {
    case 32768:
        if (!(42 >= -1)) {
            v2 -= v4
            var v11 bool = false
            _ = v11
        }
    case 789:
        v12 := GenS2{F0: 195}
        _ = v12
    default:
        v4 -= v4
    }
    switch 42 {
    case -9223372036854775808:
        switch -546 {
        case -636:
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 uint = 145
                v4 -= (v2 % ((1) | 1))
                if ("\n\t" > "\n\t") {
                    v15 := GenS2{F0: 128}
                    v4 -= v2
                    delete(v1, 65535)
                    _ = v15
                } else {
                    switch 0 {
                    case -743:
                        v16 := GenS1{F0: true, F1: -1, F2: 1000}
                        _ = v16
                    case -1:
                        for fi17 := 0; fi17 < 13; fi17++ {
                            var v18 uint64 = 42
                            _ = len(v3)
                            v4 = 168
                            v4 = (v4 >> ((uint(v2)) & 63))
                            var v19 uint8 = 230
                            if true {
                                var v20 uint64 = v18
                                v18 -= uint64(v14)
                                switch (("hello" + "") + ("foo bar" + ("a" + ""))) {
                                case "":
                                    v21 := GenS3{F0: v2, F1: ""}
                                    _ = v21
                                }
                                _ = v20
                            }
                            var v22 uint32 = 2
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v22
                        }
                    case 4294967295:
                        v23 := GenS1{F0: true, F1: 42, F2: 477}
                        _ = v23
                    default:
                        v24 := GenS2{F0: 127}
                        _ = v24
                    }
                }
                _ = fi13
                _ = v14
            }
        case 128:
            var v25 uint = 4294967295
            _ = v25
        }
    }
    if true {
        var v26 float64 = 3.14159
        v27 := GenS1{F0: (100 != 122), F1: 9223372036854775807, F2: -582}
        _ = v26
        _ = v27
    } else {
        for fi28 := 0; fi28 < 4; fi28++ {
            for fi29 := 0; fi29 < 18; fi29++ {
                v2 -= v4
                _ = fi29
            }
            v30 := GenS3{F0: (v2 * v4), F1: ""}
            for fi31 := 0; fi31 < 4; fi31++ {
                var v32 uint8 = 127
                _ = fi31
                _ = v32
            }
            _ = fi28
            _ = v30
        }
    }
    v33 := GenS4{F0: v4, F1: v4, F2: 270}
    for fi34 := 0; fi34 < 12; fi34++ {
        if true {
            if len(v5) > 0 {
                v5[0] = (^(fi34))
            }
            if ((914 != 955) && (true || ((1.5 < 0.0) || false))) {
                if true {
                    for fi35 := 0; fi35 < 2; fi35++ {
                        fmt.Println(v2, v4)
                        var v36 uint32 = 1
                        var v37 uint16 = v4
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    for fi38 := 0; fi38 < 13; fi38++ {
                        fmt.Println(fi38, v2, fi38)
                        fmt.Println(fi38, v4, v4)
                        v1[65535] = 859
                        _ = fi38
                    }
                    var v39 uint64 = 42
                    _ = v39
                } else {
                    if false {
                        var v40 int8 = 1
                        delete(v3, fi34)
                        _ = v40
                    } else {
                        v41 := GenS4{F0: v4, F1: (v4 + 255), F2: (v2 & 218)}
                        v2 += (^(v2))
                        _ = v41
                    }
                    var v42 uint = 255
                    v2 += (^(v4))
                    _ = v42
                }
                switch "" {
                case "a":
                    if false {
                        fmt.Println(fi34, v4, fi34)
                        fmt.Println(fi34)
                        _ = len(v1)
                    } else {
                        v43 := GenS2{F0: 0}
                        delete(v1, 2)
                        v5 = append(v5, (-(fi34)))
                        _ = v43
                    }
                }
                var v44 int16 = 768
                _ = v44
            }
        } else {
            v45 := GenS1{F0: true, F1: 0, F2: 0}
            v46 := GenS5{F0: 42}
            _ = v45
            _ = v46
        }
        if (0.0 < 0.0) {
            for fi47 := 0; fi47 < 13; fi47++ {
                v48 := GenS6{F0: 65535, F1: ((1.5 < 1.5) || (128 > 128)), F2: 0}
                var v49 uint32 = 1
                _ = fi47
                _ = v48
                _ = v49
            }
            var v50 uint32 = 2
            _ = v50
        } else {
            if ((fi34 != (fi34 ^ 4294967295)) && (!true)) {
                if false {
                    v4 = v4
                    for fi51 := 0; fi51 < 20; fi51++ {
                        var v52 uint8 = 193
                        if false {
                            _ = len(v5)
                            var v53 int8 = 0
                            _ = v53
                        } else {
                            switch "hello" {
                            case "\n\t":
                                if true {
                                    switch fi34 {
                                    case 719:
                                        delete(v3, (^(fi51)))
                                    case -345:
                                        v54 := GenS7{F0: 994}
                                        _ = v54
                                    }
                                } else {
                                    if (2147483648 >= 2147483647) {
                                        fmt.Println(v52)
                                    } else {
                                        if (1e10 > float32((float64(-2.25) - (0.1 / 0.1)))) {
                                            var v55 int16 = 42
                                            if false {
                                                v55 = (v55 % (((v55 << ((0) & 63))) | 1))
                                            }
                                            v56 := GenS5{F0: v2}
                                            var v57 float32 = ((1e10 / 0.1) / 1.0)
                                            v57 += v57
                                            _ = v55
                                            _ = v56
                                            _ = v57
                                        }
                                        v2 = v4
                                    }
                                    v58 := GenS1{F0: true, F1: 42, F2: -1}
                                    if (2 == 0) {
                                        if false {
                                            v4 += 2
                                            v59 := GenS4{F0: (^(v2)), F1: 355, F2: (^(v4))}
                                            v5 = append(v5, 32768)
                                            v60 := v3[(fi51 ^ fi51)]
                                            switch "foo bar" {
                                            case "hello":
                                                var v61 uint32 = 0
                                                _ = v61
                                            }
                                            _ = v59
                                            _ = v60
                                        } else {
                                            switch "\n\t" {
                                            case "\n\t":
                                                v52 = 128
                                            case "a":
                                                switch "hello" {
                                                case "foo bar":
                                                    v4 = 32767
                                                    v62 := GenS6{F0: v4, F1: (255 <= 474), F2: 0}
                                                    var v63 uint = 0
                                                    fmt.Println(fi34, v52)
                                                    var v64 int8 = -128
                                                    if (!((v52 ^ v52) == v52)) {
                                                        var v65 int8 = v64
                                                        var v66 uint64 = 42
                                                        if (false || false) {
                                                            var v67 uint16 = (v2 & (v4 + v2))
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            fmt.Println(v63, v67)
                                                            var v68 int64 = 42
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                        _ = v65
                                                        _ = v66
                                                    } else {
                                                        v63 = 459
                                                        v69 := GenS5{F0: (v2 << ((v63) & 63))}
                                                        _ = v69
                                                    }
                                                    v63 += v63
                                                    v64 = v64
                                                    _ = v62
                                                    _ = v63
                                                    _ = v64
                                                case "\n\t":
                                                    v70 := GenS5{F0: 822}
                                                    _ = v70
                                                case "a":
                                                    fmt.Println(v4)
                                                default:
                                                    _ = len(v5)
                                                }
                                            }
                                        }
                                        if true {
                                            if false {
                                                v2 = 2
                                            } else {
                                                fmt.Println(fi34, v2, fi51)
                                                var v71 uint64 = 403
                                                fmt.Println(v71, v2)
                                                _ = v71
                                            }
                                        } else {
                                            fmt.Println(fi34, v2)
                                            v72 := GenS5{F0: (v4 & (v4 * 1))}
                                            if ((!true) || false) {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                if false {
                                                    fmt.Println(fi51, v2, v4)
                                                    if (919 < 18446744073709551615) {
                                                        v2 -= (^(v2))
                                                    } else {
                                                        v73 := GenS2{F0: v52}
                                                        _ = v73
                                                    }
                                                } else {
                                                    var v74 uint8 = 42
                                                    v75 := GenS2{F0: v74}
                                                    fmt.Println(fi34, v74, fi34)
                                                    _ = v74
                                                    _ = v75
                                                }
                                                var v76 int8 = 100
                                                _ = v76
                                            }
                                            fmt.Println(v2, fi34, v2)
                                            _ = v72
                                        }
                                        v3[fi51] = false
                                        v2 += v2
                                        var v77 int16 = -32768
                                        v77 = 749
                                        _ = v77
                                    } else {
                                        var v78 float32 = (3.14159 - 0.1)
                                        if (!true) {
                                            if (!true) {
                                                var v79 uint32 = 571
                                                v80 := GenS1{F0: true, F1: -444, F2: -1}
                                                v81 := GenS3{F0: (v4 / ((v2) | 1)), F1: "hello"}
                                                _ = v79
                                                _ = v80
                                                _ = v81
                                            } else {
                                                if (fi34 == 127) {
                                                    var v82 uint = 18446744073709551615
                                                    fmt.Println(fi34, v4)
                                                    fmt.Println(v2, fi51, v4)
                                                    _ = v82
                                                }
                                                var v83 uint32 = 2147483647
                                                _ = v83
                                            }
                                            switch fi34 {
                                            case -251:
                                                v5 = append(v5, fi51)
                                            case 4294967295:
                                                v84 := GenS8{F0: 2, F1: v4, F2: "\n\t"}
                                                _ = v84
                                            }
                                            switch -2147483648 {
                                            case -1:
                                                v4 -= v2
                                            }
                                        } else {
                                            var v85 int32 = 1
                                            if (42 != v4) {
                                                v4 = (v2 + v2)
                                            } else {
                                                var v86 float64 = ((float64(float32(0.1)) + (float64(3.14159) + -2.25)) / 0.1)
                                                _ = v86
                                            }
                                            var v87 float64 = (1.5 + 0.0)
                                            _ = v85
                                            _ = v87
                                        }
                                        if false {
                                            var v88 string = ("" + "a")
                                            _ = v88
                                        }
                                        _ = v78
                                    }
                                    _ = v58
                                }
                            default:
                                var v89 uint = 1
                                _ = v89
                            }
                        }
                        var v90 int64 = 42
                        _ = fi51
                        _ = v52
                        _ = v90
                    }
                    if (!false) {
                        delete(v3, 128)
                    }
                } else {
                    var v91 float32 = 0.0
                    var v92 float32 = v91
                    v93 := v1[4294967295]
                    _ = v91
                    _ = v92
                    _ = v93
                }
                v4 += 255
            }
            switch "foo bar" {
            case "foo bar":
                var v94 uint32 = 2147483647
                _ = v94
            default:
                _ = len(v3)
            }
            v95 := GenS9{F0: v4, F1: -437, F2: 127}
            _ = v95
        }
        _ = fi34
    }
    var v96 int64 = -9223372036854775808
    fmt.Println(v2)
    if false {
        switch 127 {
        case 0:
            var v97 float64 = (((float64(1e10) / float64((-2.25 * ((3.14159 - 3.14159) - 0.1)))) / 1e10) * (((((0.0 - 1e10) / 1.5) - -2.25) - ((0.0 - 3.14159) + 3.14159)) + -2.25))
            _ = v97
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v96
}
