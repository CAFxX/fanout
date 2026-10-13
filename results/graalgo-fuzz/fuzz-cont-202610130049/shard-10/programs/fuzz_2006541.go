package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int16 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 bool; F2 uint64 }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 uint64; F1 int; F2 string }
type GenS7 struct { F0 bool; F1 uint8; F2 string }

func main() {
    v1 := []int{1, 4294967295, 9223372036854775807, 702}
    var v2 int8 = 122
    var v3 float64 = float64(1.5)
    v4 := make(map[int]string)
    var v5 int16 = 46
    delete(v4, -496)
    v2 += (v2 << ((0) & 63))
    v5 = v5
    v5 = v5
    for fi6 := 0; fi6 < 10; fi6++ {
        if false {
            switch "a" {
            case "\n\t":
                var v7 int64 = -9223372036854775808
                _ = v7
            }
        }
        v2 = 1
        _ = fi6
    }
    v8 := GenS1{F0: 0, F1: "\n\t", F2: 557}
    delete(v4, 32767)
    switch -1 {
    case 2147483647:
        if false {
            fmt.Println(v2, v5, v2)
            switch 9223372036854775807 {
            case 2147483647:
                var v9 int8 = (v2 ^ (v2 & v2))
                _ = v9
            case 32767:
                if false {
                    v3 = float64((-2.25 + 1.5))
                    fmt.Println(v5, v3)
                } else {
                    switch -871 {
                    case 127:
                        if false {
                            if ((!true) || (!(false && (!false)))) {
                                for fi10 := 0; fi10 < 20; fi10++ {
                                    if true {
                                        var v11 int32 = -2147483648
                                        _ = v11
                                    } else {
                                        var v12 float32 = 3.14159
                                        _ = v12
                                    }
                                    v5 += v5
                                    for fi13 := 0; fi13 < 9; fi13++ {
                                        switch "\n\t" {
                                        case "a":
                                            v14 := GenS1{F0: 1, F1: "hello", F2: v5}
                                            _ = v14
                                        case "":
                                            fmt.Println(fi10, v3)
                                            if (false && true) {
                                                var v15 float32 = -2.25
                                                _ = v15
                                            } else {
                                                v3 -= (v3 + -2.25)
                                            }
                                        case "\n\t":
                                            v3 -= 0.1
                                        default:
                                            var v16 uint64 = 1
                                            _ = v16
                                        }
                                        var v17 uint8 = 0
                                        _ = fi13
                                        _ = v17
                                    }
                                    _ = fi10
                                }
                                switch ("a" + "hello") {
                                case "foo bar":
                                    v2 += 1
                                case "":
                                    v18 := GenS2{F0: 65535, F1: (v5 >> ((2) & 63)), F2: 18446744073709551615}
                                    _ = v18
                                case "\n\t":
                                    var v19 int32 = 2147483647
                                    _ = v19
                                }
                                if true {
                                    var v20 float64 = v3
                                    _ = v20
                                } else {
                                    v21 := v4[32767]
                                    var v22 uint = 254
                                    _ = v21
                                    _ = v22
                                }
                            } else {
                                if (false || (1 == 9223372036854775807)) {
                                    var v23 int32 = -830
                                    var v24 float32 = 1.5
                                    if false {
                                        v2 = (v2 ^ v2)
                                    } else {
                                        switch -1 {
                                        case -9223372036854775808:
                                            if false {
                                                if ((v3 == v3) && false) {
                                                    for fi25 := 0; fi25 < 9; fi25++ {
                                                        v26 := GenS3{F0: (v5 ^ (v5 * v5))}
                                                        fmt.Println(fi25, v24, v23)
                                                        _ = fi25
                                                        _ = v26
                                                    }
                                                }
                                                for fi27 := 0; fi27 < 15; fi27++ {
                                                    if (101 >= 648) {
                                                        var v28 int = (fi27 + -9223372036854775808)
                                                        var v29 uint32 = 2147483648
                                                        switch "a" {
                                                        case "\n\t":
                                                            for fi30 := 0; fi30 < 5; fi30++ {
                                                                v29 += (v29 | (v29 >> ((0) & 63)))
                                                                _ = fi30
                                                            }
                                                        case "":
                                                            fmt.Println(v28, v5)
                                                        }
                                                        _ = v28
                                                        _ = v29
                                                    }
                                                    for fi31 := 0; fi31 < 11; fi31++ {
                                                        fmt.Println(fi31, v2, fi31)
                                                        var v32 int8 = (v2 ^ 42)
                                                        _ = fi31
                                                        _ = v32
                                                    }
                                                    if (v24 >= (((float32((-2.25 * (float64(((v24 + (v24 + (float32(0.0) - (v24 - (((v24 / v24) / float32(float64(-2.25))) - v24))))) * v24)) + v3))) + v24) + ((v24 + ((v24 + v24) - (((v24 / 0.1) * float32(v3)) * (((v24 / 3.14159) + v24) / 1.5)))) * v24)) / v24)) {
                                                        fmt.Println(fi27)
                                                        v2 += -128
                                                    } else {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        var v33 uint8 = 255
                                                        _ = v33
                                                    }
                                                    _ = fi27
                                                }
                                            }
                                        }
                                    }
                                    _ = v23
                                    _ = v24
                                } else {
                                    var v34 int = 32768
                                    v1 = append(v1, 9223372036854775807)
                                    _ = len(v1)
                                    _ = v34
                                }
                                var v35 uint16 = 17
                                if (("foo bar" + ("hello" + "hello")) > "hello") {
                                    v36 := GenS4{F0: 42, F1: false, F2: 1}
                                    v4[-362] = "\n\t"
                                    _ = v36
                                } else {
                                    v1 = append(v1, 258)
                                    var v37 uint32 = 65535
                                    v2 += (v2 | (v2 & -1))
                                    _ = v37
                                }
                                _ = v35
                            }
                            v38 := GenS3{F0: 1000}
                            if false {
                                if true {
                                    v39 := GenS2{F0: 2147483647, F1: (v5 - v5), F2: 4294967295}
                                    for fi40 := 0; fi40 < 10; fi40++ {
                                        for fi41 := 0; fi41 < 6; fi41++ {
                                            v3 = v3
                                            var v42 uint = 18446744073709551615
                                            _ = fi41
                                            _ = v42
                                        }
                                        var v43 bool = true
                                        if (18446744073709551615 == 255) {
                                            if (false || v43) {
                                                v1 = append(v1, (fi40 | (fi40 / ((fi40) | 1))))
                                                if (0 < 2) {
                                                    for fi44 := 0; fi44 < 3; fi44++ {
                                                        var v45 uint16 = 2
                                                        _ = fi44
                                                        _ = v45
                                                    }
                                                    for fi46 := 0; fi46 < 11; fi46++ {
                                                        var v47 bool = v43
                                                        _ = fi46
                                                        _ = v47
                                                    }
                                                    v48 := GenS2{F0: 1, F1: -730, F2: 1}
                                                    _ = v48
                                                } else {
                                                    v5 -= (v5 + (v5 << ((2) & 63)))
                                                    if true {
                                                        v49 := GenS1{F0: 18446744073709551615, F1: "a", F2: v5}
                                                        if (127 == 0) {
                                                            switch "hello" {
                                                            case "":
                                                                if false {
                                                                    if (true || v43) {
                                                                        var v50 float32 = (1.5 / 1e10)
                                                                        for fi51 := 0; fi51 < 9; fi51++ {
                                                                            var v52 int = (^(fi40))
                                                                            var v53 string = "a"
                                                                            _ = fi51
                                                                            _ = v52
                                                                            _ = v53
                                                                        }
                                                                        _ = v50
                                                                    }
                                                                    var v54 float32 = 1e10
                                                                    _ = v54
                                                                }
                                                            case "\n\t":
                                                                switch "foo bar" {
                                                                case "a":
                                                                    var v55 int8 = 48
                                                                    _ = v55
                                                                case "\n\t":
                                                                    var v56 int8 = (v2 & (v2 >> ((2) & 63)))
                                                                    _ = v56
                                                                default:
                                                                    switch fi40 {
                                                                    case 127:
                                                                        v2 += (v2 % ((int8(fi40)) | 1))
                                                                    default:
                                                                        var v57 uint8 = 255
                                                                        _ = v57
                                                                    }
                                                                }
                                                            }
                                                            var v58 bool = (18446744073709551615 >= 0)
                                                            _ = v58
                                                        }
                                                        _ = v49
                                                    }
                                                    v5 = -104
                                                }
                                                v59 := GenS3{F0: v5}
                                                _ = v59
                                            } else {
                                                fmt.Println(v2, v3, fi40)
                                                var v60 uint64 = 1
                                                _ = v60
                                            }
                                        }
                                        _ = fi40
                                        _ = v43
                                    }
                                    v3 += 0.0
                                    _ = v39
                                }
                                if (false && (!false)) {
                                    var v61 bool = false
                                    fmt.Println(v2)
                                    v62 := GenS2{F0: 1, F1: (v5 % ((v5) | 1)), F2: 1}
                                    _ = v61
                                    _ = v62
                                } else {
                                    if false {
                                        v3 -= float64(1.5)
                                        for fi63 := 0; fi63 < 14; fi63++ {
                                            if false {
                                                if false {
                                                    fmt.Println(fi63)
                                                    _ = len(v4)
                                                } else {
                                                    var v64 int = fi63
                                                    _ = v64
                                                }
                                            }
                                            v65 := GenS5{F0: 9223372036854775807, F1: "hello"}
                                            _ = fi63
                                            _ = v65
                                        }
                                    } else {
                                        for fi66 := 0; fi66 < 17; fi66++ {
                                            var v67 uint32 = 159
                                            _ = fi66
                                            _ = v67
                                        }
                                        var v68 uint = 18446744073709551615
                                        _ = v68
                                    }
                                    for fi69 := 0; fi69 < 5; fi69++ {
                                        v70 := GenS6{F0: 2, F1: fi69, F2: "hello"}
                                        for fi71 := 0; fi71 < 8; fi71++ {
                                            v72 := GenS2{F0: 4294967295, F1: 42, F2: 777}
                                            var v73 int64 = 0
                                            v74 := GenS7{F0: (!(false || false)), F1: 127, F2: ("a" + ("hello" + "\n\t"))}
                                            _ = fi71
                                            _ = v72
                                            _ = v73
                                            _ = v74
                                        }
                                        _ = fi69
                                        _ = v70
                                    }
                                    v75 := v4[1]
                                    _ = v75
                                }
                            } else {
                                fmt.Println(v2, v5, v5)
                                var v76 uint = 0
                                v4[344] = ("a" + "")
                                _ = v76
                            }
                            _ = v38
                        } else {
                            v5 = v5
                            var v77 int16 = v5
                            _ = v77
                        }
                    case 42:
                        v78 := GenS1{F0: 18446744073709551615, F1: "", F2: v5}
                        _ = v78
                    default:
                        if true {
                            v79 := GenS6{F0: 42, F1: 2147483647, F2: "foo bar"}
                            v2 = -93
                            _ = v79
                        }
                    }
                }
            }
        }
    }
    var v80 int8 = (v2 << ((2) & 63))
    switch 1 {
    case 32768:
        fmt.Println(v3)
    case 128:
        v2 -= (v80 + 21)
    case -2147483648:
        var v81 int64 = 1
        _ = v81
    }
    switch -2147483648 {
    case 369:
        v82 := v4[2147483647]
        _ = v82
    case -414:
        if (-346 > 0) {
            fmt.Println(v5, v80, v3)
            v3 = (1e10 / v3)
        }
    case -973:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        var v83 uint8 = 1
        _ = v83
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v80
}
