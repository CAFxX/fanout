package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 bool }
type GenS2 struct { F0 int64; F1 bool; F2 int8 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int; F1 int16; F2 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS8 struct { F0 int; F1 uint; F2 int }
type GenS9 struct { F0 string; F1 uint; F2 int8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 uint32 = 2147483648
    v4 := make(map[uint32]bool)
    v5 := []int{}
    fmt.Println(v3)
    var v6 uint8 = 42
    v7 := v1[42]
    var v8 int16 = -589
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 float32 = (0.1 / float32(((1.5 + 0.1) - 0.0)))
        _ = fi9
        _ = v10
    }
    var v11 uint8 = (v6 + v6)
    v8 = v8
    v7 = (v3 | (v3 % ((v7) | 1)))
    switch 843 {
    case -725:
        switch -631 {
        case -9223372036854775808:
            if (42 > 65535) {
                if true {
                    switch ("a" + "hello") {
                    case "\n\t":
                        if false {
                            fmt.Println(v6, v6)
                            v8 -= v8
                        } else {
                            v5 = append(v5, 187)
                            v12 := GenS1{F0: 32767, F1: v8, F2: false}
                            _ = v12
                        }
                    case "":
                        v3 = (v3 % (((v7 / (((v7 & (v7 % (((v3 + v7)) | 1)))) | 1))) | 1))
                    case "hello":
                        if false {
                            v11 = (v11 >> ((4294967295) & 63))
                        }
                    }
                    for fi13 := 0; fi13 < 11; fi13++ {
                        v7 += 710
                        var v14 float32 = 1.5
                        _ = fi13
                        _ = v14
                    }
                }
                for fi15 := 0; fi15 < 15; fi15++ {
                    if true {
                        var v16 int8 = 105
                        _ = v16
                    } else {
                        if (!(!(!((v8 << ((65535) & 63)) != (v8 << ((65535) & 63)))))) {
                            for fi17 := 0; fi17 < 10; fi17++ {
                                if (18446744073709551615 != 65535) {
                                    if ("hello" == (("a" + "") + "hello")) {
                                        v18 := GenS2{F0: 9223372036854775807, F1: false, F2: 100}
                                        switch "a" {
                                        case "foo bar":
                                            if (false || (1.5 != (1e10 + float32(float64(float32(1.5)))))) {
                                                v8 += (v8 | 42)
                                                if true {
                                                    fmt.Println(v7, v7)
                                                    var v19 uint = 4294967295
                                                    _ = v19
                                                } else {
                                                    fmt.Println(v11, fi17, fi17)
                                                }
                                                v20 := GenS1{F0: 32767, F1: 32767, F2: true}
                                                _ = v20
                                            }
                                        case "":
                                            v21 := GenS3{F0: 801, F1: (519 != -95)}
                                            _ = v21
                                        case "\n\t":
                                            fmt.Println(v11, v6, v7)
                                        }
                                        _ = v18
                                    }
                                }
                                if ((!true) && (2 < 42)) {
                                    v8 = v8
                                    _ = len(v1)
                                    v22 := GenS4{F0: (^(fi17)), F1: 1000, F2: v7}
                                    _ = v22
                                } else {
                                    v23 := GenS1{F0: 1, F1: (v8 - v8), F2: true}
                                    v24 := GenS3{F0: -2147483648, F1: false}
                                    if false {
                                        delete(v4, v7)
                                        fmt.Println(v6, fi17, v11)
                                        v25 := GenS4{F0: -9223372036854775808, F1: v8, F2: (v3 % ((v7) | 1))}
                                        _ = v25
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        fmt.Println(fi17, fi15, v8)
                                    }
                                    v26 := GenS5{F0: "foo bar"}
                                    _ = v23
                                    _ = v24
                                    _ = v26
                                }
                                _ = fi17
                            }
                        } else {
                            for fi27 := 0; fi27 < 1; fi27++ {
                                var v28 int64 = -9223372036854775808
                                if len(v5) > 0 {
                                    v5[0] = (fi15 | 42)
                                }
                                v29 := GenS6{F0: 1}
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                            switch 32768 {
                            case -9223372036854775808:
                                if true {
                                    var v30 int16 = -32768
                                    v7 = v7
                                    fmt.Println(v11, v7, v8)
                                    _ = v30
                                } else {
                                    switch ("" + (((("foo bar" + "foo bar") + "foo bar") + "a") + "foo bar")) {
                                    case "":
                                        v31 := GenS5{F0: "\n\t"}
                                        _ = v31
                                    case "\n\t":
                                        fmt.Println(fi15, v3)
                                    }
                                    for fi32 := 0; fi32 < 13; fi32++ {
                                        delete(v2, fi15)
                                        var v33 int = fi32
                                        _ = fi32
                                        _ = v33
                                    }
                                    v3 -= v7
                                }
                            case 0:
                                v34 := GenS7{F0: (^(v3)), F1: (v11 + v11), F2: -128}
                                _ = v34
                            }
                        }
                    }
                    v11 = (v11 * v11)
                    for fi35 := 0; fi35 < 20; fi35++ {
                        v36 := GenS2{F0: 9223372036854775807, F1: false, F2: 0}
                        _ = fi35
                        _ = v36
                    }
                    _ = fi15
                }
                var v37 float64 = -2.25
                _ = v37
            } else {
                var v38 uint32 = 755
                _ = v38
            }
        case 744:
            if (4294967295 > 4294967295) {
                v5 = append(v5, 18)
                for fi39 := 0; fi39 < 8; fi39++ {
                    var v40 string = ""
                    v41 := GenS8{F0: fi39, F1: 344, F2: fi39}
                    var v42 float64 = (3.14159 - 3.14159)
                    _ = fi39
                    _ = v40
                    _ = v41
                    _ = v42
                }
            }
        case -1:
            if true {
                switch 9223372036854775807 {
                case -123:
                    for fi43 := 0; fi43 < 10; fi43++ {
                        _ = len(v2)
                        fmt.Println(fi43, v11, v8)
                        if ("\n\t" != ("hello" + "a")) {
                            if (v8 >= v8) {
                                var v44 float64 = 1e10
                                var v45 uint64 = 4294967295
                                _ = v44
                                _ = v45
                            }
                            v46 := GenS7{F0: 1, F1: v11, F2: 1}
                            _ = v46
                        } else {
                            _ = len(v4)
                            if false {
                                v47 := GenS5{F0: "foo bar"}
                                _ = v47
                            }
                            switch fi43 {
                            case -9223372036854775808:
                                if (!true) {
                                    var v48 int64 = -498
                                    for fi49 := 0; fi49 < 19; fi49++ {
                                        var v50 uint8 = 2
                                        _ = fi49
                                        _ = v50
                                    }
                                    _ = v48
                                } else {
                                    v11 += (v6 / ((v11) | 1))
                                    var v51 uint = 1
                                    _ = v51
                                }
                            default:
                                switch 1 {
                                case 4294967295:
                                    v7 -= (v7 << ((788) & 63))
                                default:
                                    var v52 int64 = 42
                                    _ = v52
                                }
                            }
                        }
                        _ = fi43
                    }
                case 42:
                    v5 = append(v5, -815)
                case 32767:
                    var v53 uint32 = v3
                    _ = v53
                default:
                    v7 -= (v7 - 2)
                }
                if false {
                    var v54 int = 331
                    _ = v54
                } else {
                    var v55 string = "\n\t"
                    _ = v55
                }
                v56 := GenS6{F0: 1}
                _ = v56
            }
        }
    }
    if ((v8 % ((v8) | 1)) < (v8 / (((v8 >> ((65535) & 63))) | 1))) {
        if true {
            var v57 uint = 2
            _ = v57
        }
        if true {
            var v58 string = ((("hello" + "foo bar") + (("" + "a") + "\n\t")) + ("a" + "hello"))
            if true {
                if ((false && (0 == -128)) && (((361 < 930) && (99 <= 568)) || (42 != 255))) {
                    v59 := v1[475]
                    if true {
                        v6 = (^(v6))
                        if (!(!true)) {
                            var v60 bool = (!true)
                            for fi61 := 0; fi61 < 15; fi61++ {
                                var v62 uint64 = 18446744073709551615
                                var v63 int64 = int64(v7)
                                var v64 float32 = (((3.14159 + (1e10 * (1e10 - (1e10 + (3.14159 + 1e10))))) + 3.14159) / (1e10 * -2.25))
                                _ = fi61
                                _ = v62
                                _ = v63
                                _ = v64
                            }
                            _ = v60
                        }
                        for fi65 := 0; fi65 < 13; fi65++ {
                            v58 = v58
                            v59 = v7
                            _ = fi65
                        }
                    }
                    if (!false) {
                        v8 = (v8 << ((605) & 63))
                        v66 := GenS4{F0: 1, F1: 543, F2: 1}
                        for fi67 := 0; fi67 < 2; fi67++ {
                            if true {
                                var v68 uint = 255
                                var v69 int = fi67
                                _ = v68
                                _ = v69
                            } else {
                                v4[232] = ((0.1 + 1.5) <= 1.5)
                                fmt.Println(v58, v7)
                            }
                            if (0 == 42) {
                                var v70 uint8 = (v6 / ((v6) | 1))
                                _ = v70
                            } else {
                                for fi71 := 0; fi71 < 20; fi71++ {
                                    v5 = append(v5, 127)
                                    _ = fi71
                                }
                                _ = len(v1)
                                if true {
                                    var v72 uint = 4294967295
                                    var v73 int8 = 100
                                    _ = v72
                                    _ = v73
                                }
                            }
                            _ = fi67
                        }
                        _ = v66
                    } else {
                        v8 -= (v8 << ((255) & 63))
                        var v74 uint64 = 368
                        for fi75 := 0; fi75 < 9; fi75++ {
                            var v76 float32 = 3.14159
                            _ = fi75
                            _ = v76
                        }
                        _ = v74
                    }
                    _ = v59
                }
                if (-553 > -1) {
                    v5 = append(v5, 32768)
                }
                v77 := GenS4{F0: 0, F1: v8, F2: (v7 + 65535)}
                _ = v77
            } else {
                if false {
                    if len(v5) > 0 {
                        v5[0] = 9223372036854775807
                    }
                }
            }
            _ = v58
        }
        v7 = v7
    } else {
        if (!false) {
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
            var v78 bool = ((!(-1 != 2147483647)) && (v6 > v6))
            _ = v78
        } else {
            if true {
                var v79 int32 = 591
                v8 = v8
                switch -418 {
                case 0:
                    for fi80 := 0; fi80 < 18; fi80++ {
                        v2[fi80] = false
                        _ = fi80
                    }
                }
                _ = v79
            } else {
                v81 := GenS5{F0: ("hello" + "hello")}
                _ = v81
            }
        }
        switch ("" + ("\n\t" + "a")) {
        case "foo bar":
            if false {
                if (!false) {
                    if (-1 == 1) {
                        var v82 float32 = (1.5 + (-2.25 * -2.25))
                        v11 = (v11 - (^(v6)))
                        v83 := GenS9{F0: "", F1: 666, F2: 124}
                        _ = v82
                        _ = v83
                    }
                }
                switch ("a" + "") {
                case "hello":
                    fmt.Println(v6, v6, v3)
                case "a":
                    var v84 uint8 = v6
                    _ = v84
                case "":
                    fmt.Println(v7, v11)
                }
            }
        default:
            for fi85 := 0; fi85 < 6; fi85++ {
                v7 -= v3
                var v86 uint8 = v6
                _ = fi85
                _ = v86
            }
        }
    }
    v8 = -32768
    if false {
        var v87 string = "foo bar"
        _ = v87
    } else {
        var v88 int64 = 42
        _ = v88
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
