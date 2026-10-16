package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 int16; F1 uint32 }
type GenS9 struct { F0 uint32; F1 int8; F2 string }
type GenS10 struct { F0 int16; F1 string; F2 int }
type GenS11 struct { F0 int; F1 uint8; F2 uint16 }
func genFn0() int32 {
    v5 := []int{1, 963, 127, 128}
    switch "foo bar" {
    case "foo bar":
        v5 = append(v5, 128)
    case "hello":
        switch "a" {
        case "\n\t":
            v6 := GenS1{F0: 0, F1: 758}
            _ = v6
        case "a":
            var v7 int32 = -2147483648
            _ = v7
        case "hello":
            var v8 uint8 = 127
            _ = v8
        }
    default:
        var v9 uint8 = 255
        _ = v9
    }
    _ = v5
    return 319
}
func genFn1(p0 uint32, p1 int8) uint64 {
    v10 := []int{}
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint = 255
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return 42
}

func main() {
    var v1 uint16 = 2
    v2 := []uint32{0}
    v3 := make(map[int]uint32)
    var v4 uint64 = 42
    v13 := GenS2{F0: 127, F1: 0, F2: v4}
    if (100 < 127) {
        _ = len(v3)
        v14 := GenS3{F0: (v1 >> ((181) & 63)), F1: 255, F2: v1}
        _ = v14
    }
    v15 := GenS1{F0: -806, F1: 4294967295}
    for fi16 := 0; fi16 < 1; fi16++ {
        v4 -= v4
        v4 = 1
        if true {
            if false {
                switch (("" + "a") + "\n\t") {
                case "a":
                    for fi17 := 0; fi17 < 3; fi17++ {
                        v2 = append(v2, 657)
                        if true {
                            if len(v2) > 0 {
                                v2[0] = 30
                            }
                            var v18 uint = 18446744073709551615
                            switch fi17 {
                            case 282:
                                v19 := GenS1{F0: -2147483648, F1: 199}
                                _ = v19
                            case -147:
                                switch (fi17 ^ 0) {
                                case -997:
                                    var v20 string = ""
                                    _ = v20
                                case -2147483648:
                                    if (!((!true) || (-2.25 < (1e10 - (1.5 - float64(1e10)))))) {
                                        var v21 uint64 = v4
                                        switch "foo bar" {
                                        case "":
                                            v18 = 18446744073709551615
                                        default:
                                            v18 += v18
                                        }
                                        if (!false) {
                                            v1 = 42
                                        } else {
                                            var v22 string = "hello"
                                            v1 += 1
                                            _ = v22
                                        }
                                        _ = v21
                                    }
                                }
                            }
                            _ = v18
                        } else {
                            switch "" {
                            case "foo bar":
                                delete(v3, -2147483648)
                            case "\n\t":
                                _ = len(v3)
                            }
                            if false {
                                v23 := GenS1{F0: 0, F1: 320}
                                switch fi17 {
                                case -538:
                                    switch "a" {
                                    case "foo bar":
                                        if true {
                                            var v24 int64 = -1
                                            var v25 bool = (65535 <= 2147483648)
                                            v26 := GenS1{F0: genFn0(), F1: 2147483647}
                                            if (((3.14159 - (1e10 * (float64(float32(1.5)) / (0.0 - 0.1)))) * float64(0.0)) < 1e10) {
                                                v27 := GenS3{F0: (^(v1)), F1: 23, F2: 32767}
                                                v28 := GenS1{F0: -868, F1: 4294967295}
                                                v4 = v4
                                                _ = v27
                                                _ = v28
                                            } else {
                                                v29 := GenS2{F0: 116, F1: v1, F2: 42}
                                                v30 := GenS2{F0: 255, F1: (v1 - 2), F2: v4}
                                                _ = v29
                                                _ = v30
                                            }
                                            v31 := GenS3{F0: (v1 - uint16(v4)), F1: 255, F2: 2}
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                            _ = v31
                                        } else {
                                            v3[(fi16 << ((18446744073709551615) & 63))] = 0
                                            v32 := GenS1{F0: 42, F1: 2147483648}
                                            _ = v32
                                        }
                                    case "hello":
                                        v1 = 65535
                                    default:
                                        var v33 uint16 = (v1 ^ v1)
                                        _ = v33
                                    }
                                case -403:
                                    fmt.Println(v1)
                                }
                                _ = v23
                            }
                            if (-629 == -883) {
                                var v34 int64 = 0
                                _ = v34
                            }
                        }
                        var v35 float32 = 0.0
                        var v36 float64 = 0.1
                        v37 := GenS4{F0: ("\n\t" + "\n\t"), F1: 1000}
                        _ = fi17
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                case "foo bar":
                    if true {
                        for fi38 := 0; fi38 < 20; fi38++ {
                            v39 := GenS5{F0: ((v1 & 349) >= 65535), F1: -1}
                            _ = fi38
                            _ = v39
                        }
                    } else {
                        fmt.Println(v4, v1, v1)
                        for fi40 := 0; fi40 < 3; fi40++ {
                            _ = len(v3)
                            _ = fi40
                        }
                    }
                }
                v1 = (v1 * (v1 >> ((4294967295) & 63)))
                var v41 int = 2147483647
                _ = v41
            }
            v42 := GenS2{F0: 48, F1: (^(v1)), F2: v4}
            _ = v42
        } else {
            v43 := GenS6{F0: 18446744073709551615}
            switch "" {
            case "":
                v44 := GenS3{F0: (v1 ^ (v1 + v1)), F1: 2, F2: 65535}
                _ = v44
            }
            if ((false || (v1 == 1)) && (!true)) {
                var v45 uint64 = genFn1(0, 127)
                for fi46 := 0; fi46 < 17; fi46++ {
                    v45 -= v45
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi46
                }
                _ = v45
            }
            _ = v43
        }
        _ = fi16
    }
    fmt.Println(v4, v4)
    v47 := GenS7{F0: 2147483648, F1: 1}
    v1 = v1
    for fi48 := 0; fi48 < 16; fi48++ {
        v49 := GenS4{F0: "", F1: 1}
        var v50 uint8 = uint8(v1)
        _ = fi48
        _ = v49
        _ = v50
    }
    v51 := GenS7{F0: 65535, F1: 515}
    _ = len(v3)
    v4 -= 752
    v52 := GenS5{F0: (3.14159 >= (0.0 - 1.5)), F1: -95}
    var v53 int64 = 9223372036854775807
    for fi54 := 0; fi54 < 6; fi54++ {
        v55 := GenS4{F0: "", F1: -1}
        _ = fi54
        _ = v55
    }
    if (uint64(v1) != (v4 - (^(v4)))) {
        v53 = v53
    } else {
        var v56 uint64 = 0
        fmt.Println(v56, v4, v4)
        _ = v56
    }
    v57 := GenS3{F0: (v1 - (v1 % ((2) | 1))), F1: 127, F2: v1}
    v58 := GenS6{F0: 743}
    if ("\n\t" == "a") {
        if (0 < 255) {
            if (true || true) {
                if false {
                    for fi59 := 0; fi59 < 20; fi59++ {
                        v60 := GenS8{F0: -32768, F1: 335}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v61 uint16 = 610
                        _ = fi59
                        _ = v60
                        _ = v61
                    }
                    v4 = v4
                    v62 := GenS2{F0: 255, F1: v1, F2: (v4 * (v4 ^ (v4 * v4)))}
                    _ = v62
                } else {
                    var v63 uint64 = v4
                    v4 = (v4 - 18446744073709551615)
                    v53 = v53
                    _ = v63
                }
                if true {
                    v64 := GenS6{F0: 18446744073709551615}
                    switch "hello" {
                    case "hello":
                        v53 -= (v53 * v53)
                    case "foo bar":
                        v65 := GenS9{F0: 4294967295, F1: 1, F2: ("hello" + "foo bar")}
                        _ = v65
                    }
                    _ = v64
                } else {
                    v4 = 42
                    if (!(true || false)) {
                        var v66 bool = (!((!(100 < 0)) || false))
                        if v66 {
                            for fi67 := 0; fi67 < 2; fi67++ {
                                var v68 int64 = -49
                                v69 := GenS5{F0: (2147483647 <= 2147483647), F1: 2}
                                _ = len(v3)
                                _ = fi67
                                _ = v68
                                _ = v69
                            }
                            if v66 {
                                v4 += genFn1(4294967295, -21)
                            }
                            if true {
                                for fi70 := 0; fi70 < 9; fi70++ {
                                    switch "a" {
                                    case "hello":
                                        delete(v3, (^(fi70)))
                                    }
                                    _ = fi70
                                }
                                if (v66 || true) {
                                    fmt.Println(v1)
                                }
                                v66 = (v1 <= v1)
                            }
                        } else {
                            if v66 {
                                for fi71 := 0; fi71 < 16; fi71++ {
                                    switch "hello" {
                                    case "hello":
                                        v66 = true
                                    case "\n\t":
                                        var v72 float32 = (1.5 - 1.5)
                                        _ = v72
                                    case "a":
                                        v73 := GenS4{F0: "\n\t", F1: -32768}
                                        _ = v73
                                    }
                                    for fi74 := 0; fi74 < 9; fi74++ {
                                        v1 = (v1 >> ((18446744073709551615) & 63))
                                        v1 = v1
                                        _ = fi74
                                    }
                                    if true {
                                        var v75 bool = ((false && (128 == 0)) || (-2147483648 < fi71))
                                        _ = v75
                                    } else {
                                        if ((v53 % ((v53) | 1)) < v53) {
                                            if (v1 > v1) {
                                                _ = len(v2)
                                            }
                                            switch (fi71 + (fi71 - (fi71 / ((fi71) | 1)))) {
                                            case -888:
                                                v2 = append(v2, 65535)
                                            }
                                        } else {
                                            v66 = v66
                                        }
                                        v66 = (!(v4 != v4))
                                        if len(v2) > 0 {
                                            v2[0] = 725
                                        }
                                    }
                                    _ = fi71
                                }
                                if ("" > "a") {
                                    fmt.Println(v1)
                                }
                            }
                            for fi76 := 0; fi76 < 15; fi76++ {
                                v1 -= (v1 / ((448) | 1))
                                v4 -= v4
                                _ = fi76
                            }
                            if (-1 != 2147483647) {
                                switch "\n\t" {
                                case "a":
                                    v77 := GenS10{F0: 42, F1: "\n\t", F2: 9223372036854775807}
                                    _ = v77
                                case "":
                                    for fi78 := 0; fi78 < 9; fi78++ {
                                        v1 += 32767
                                        for fi79 := 0; fi79 < 17; fi79++ {
                                            if len(v2) > 0 {
                                                v2[0] = 2147483648
                                            }
                                            _ = fi79
                                        }
                                        v1 -= 255
                                        _ = fi78
                                    }
                                case "\n\t":
                                    v2 = append(v2, 4294967295)
                                default:
                                    for fi80 := 0; fi80 < 3; fi80++ {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        for fi81 := 0; fi81 < 10; fi81++ {
                                            fmt.Println(v1, v53, fi80)
                                            _ = fi81
                                        }
                                        for fi82 := 0; fi82 < 17; fi82++ {
                                            var v83 int = (fi80 | (fi82 << ((80) & 63)))
                                            _ = fi82
                                            _ = v83
                                        }
                                        _ = fi80
                                    }
                                }
                                v1 += 32767
                                switch "\n\t" {
                                case "foo bar":
                                    v2 = append(v2, 0)
                                }
                            }
                        }
                        _ = v66
                    }
                }
            }
        }
        for fi84 := 0; fi84 < 19; fi84++ {
            if true {
                var v85 int = 547
                var v86 float32 = float32(-2.25)
                for fi87 := 0; fi87 < 11; fi87++ {
                    if false {
                        v88 := GenS8{F0: 1, F1: 4294967295}
                        if false {
                            switch 763 {
                            case 9223372036854775807:
                                if false {
                                    v89 := GenS11{F0: 9223372036854775807, F1: 1, F2: (^(v1))}
                                    _ = v89
                                } else {
                                    fmt.Println(v1, v85, v86)
                                }
                            default:
                                if true {
                                    fmt.Println(v86, v86)
                                } else {
                                    var v90 int = 4294967295
                                    var v91 uint32 = 615
                                    _ = v90
                                    _ = v91
                                }
                            }
                        }
                        _ = v88
                    }
                    _ = fi87
                }
                _ = v85
                _ = v86
            } else {
                if false {
                    for fi92 := 0; fi92 < 7; fi92++ {
                        var v93 float64 = 1e10
                        _ = fi92
                        _ = v93
                    }
                } else {
                    for fi94 := 0; fi94 < 4; fi94++ {
                        v53 = 42
                        _ = fi94
                    }
                    fmt.Println(fi84, fi84)
                }
                for fi95 := 0; fi95 < 18; fi95++ {
                    var v96 float32 = float32((0.0 + 0.0))
                    v1 += v1
                    _ = fi95
                    _ = v96
                }
                v53 -= (-(v53))
            }
            _ = fi84
        }
    } else {
        var v97 int64 = (v53 + (v53 / ((0) | 1)))
        _ = v97
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v51
    _ = v52
    _ = v53
    _ = v57
    _ = v58
}
