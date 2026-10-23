package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 bool; F1 uint16; F2 string }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool }
func genFn0(p0 uint8, p1 uint8) int {
    v5 := []int{-918, 42, -1, -514}
    var v6 int16 = -32768
    fmt.Println(v6, v6)
    v5 = append(v5, -538)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -162
}
func genFn1(p0 uint64) bool {
    var v7 float64 = 1e10
    switch 2147483647 {
    case -9223372036854775808:
        v8 := []int{4294967295, 1, genFn0(171, 255)}
        _ = v8
    default:
        v9 := GenS1{F0: 1, F1: 42, F2: 42}
        _ = v9
    }
    var v10 int32 = 42
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := GenS1{F0: 2, F1: 101, F2: v10}
            p0 = p0
            _ = fi12
            _ = v13
        }
        var v14 uint64 = p0
        _ = fi11
        _ = v14
    }
    switch 9223372036854775807 {
    case 32768:
        var v15 int32 = (v10 & 2147483647)
        _ = v15
    case -2147483648:
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 string = ("foo bar" + "\n\t")
            v18 := GenS1{F0: 255, F1: 1, F2: v10}
            _ = fi16
            _ = v17
            _ = v18
        }
    default:
        var v19 uint = 4294967295
        _ = v19
    }
    _ = p0
    _ = v10
    _ = v7
    return (-994 == 32767)
}

func main() {
    v1 := []int{-861, -291, -1, 0}
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v4 uint64 = 411
    v4 += v4
    if (-2.25 <= (3.14159 - 3.14159)) {
        var v20 int32 = 42
        if (66 != 1) {
            var v21 float32 = 1.5
            _ = v21
        }
        for fi22 := 0; fi22 < 10; fi22++ {
            fmt.Println(v4, v20, v4)
            var v23 float64 = (((0.0 * float64((3.14159 + -2.25))) - 3.14159) / 1.0)
            switch ("\n\t" + ("a" + "foo bar")) {
            case "a":
                switch genFn0(2, 255) {
                case 9223372036854775807:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 42:
                    for fi24 := 0; fi24 < 10; fi24++ {
                        v4 = v4
                        delete(v2, genFn0(127, uint8(v4)))
                        _ = fi24
                    }
                case -619:
                    var v25 int32 = 454
                    _ = v25
                }
            case "hello":
                v4 = (^(v4))
            default:
                switch "hello" {
                case "hello":
                    v4 = v4
                case "\n\t":
                    if false {
                        v3[414] = true
                    } else {
                        fmt.Println(v23, v23, fi22)
                        v4 = v4
                    }
                default:
                    if true {
                        switch -1 {
                        case -9223372036854775808:
                            for fi26 := 0; fi26 < 6; fi26++ {
                                v27 := GenS1{F0: 959, F1: 255, F2: v20}
                                _ = fi26
                                _ = v27
                            }
                        case 128:
                            v20 -= v20
                        default:
                            switch 196 {
                            case 42:
                                v20 = (^(v20))
                            case 0:
                                fmt.Println(fi22, v23)
                            case -1:
                                if false {
                                    v23 = ((v23 * float64(1e10)) / (((v23 + (1.5 + (v23 / 0.1))) + 1.5) * v23))
                                }
                            }
                        }
                        var v28 float64 = float64((3.14159 + (0.1 / 0.1)))
                        v4 += 2
                        _ = v28
                    } else {
                        v20 = 2147483647
                        v29 := GenS2{F0: 0, F1: 0}
                        v30 := GenS2{F0: 992, F1: 42}
                        _ = v29
                        _ = v30
                    }
                }
            }
            _ = fi22
            _ = v23
        }
        _ = v20
    }
    v4 -= (v4 >> ((4294967295) & 63))
    var v31 int8 = -73
    var v32 int8 = 100
    fmt.Println(v32, v31)
    if true {
        var v33 int = 32767
        if true {
            var v34 int8 = -128
            for fi35 := 0; fi35 < 8; fi35++ {
                v36 := GenS3{F0: true, F1: 255, F2: "\n\t"}
                v31 = int8(v33)
                _ = fi35
                _ = v36
            }
            _ = v34
        } else {
            v37 := GenS3{F0: false, F1: 1, F2: ("a" + ("\n\t" + (("" + ("\n\t" + "\n\t")) + ("a" + ""))))}
            v33 = (v33 - v33)
            _ = len(v3)
            _ = v37
        }
        _ = v33
    }
    if ("foo bar" == "") {
        for fi38 := 0; fi38 < 2; fi38++ {
            for fi39 := 0; fi39 < 7; fi39++ {
                switch fi39 {
                case 42:
                    _ = len(v1)
                    var v40 uint32 = 2147483648
                    switch fi39 {
                    case -363:
                        if true {
                            fmt.Println(v32, fi38, fi39)
                            if true {
                                switch (fi39 / ((-2147483648) | 1)) {
                                case 32767:
                                    v40 = (v40 - (v40 + v40))
                                case 2147483647:
                                    v2[fi39] = fi39
                                }
                                v1 = append(v1, fi39)
                                var v41 bool = ((!(true && (!(!true)))) && true)
                                _ = v41
                            }
                        } else {
                            v32 = (v32 ^ v31)
                            if ((fi39 % (((fi39 % (((fi38 % ((genFn0(1, 127)) | 1))) | 1))) | 1)) == (fi39 / (((fi38 * (fi38 - (fi38 / ((32767) | 1))))) | 1))) {
                                v32 = 127
                                var v42 uint = uint(v4)
                                switch "\n\t" {
                                case "hello":
                                    _ = len(v3)
                                default:
                                    fmt.Println(v42)
                                }
                                _ = v42
                            } else {
                                v31 = v32
                                v32 += int8(v4)
                                v40 -= v40
                                if (32767 <= 372) {
                                    switch "" {
                                    case "":
                                        if false {
                                            if len(v1) > 0 {
                                                v1[0] = fi39
                                            }
                                            if false {
                                                fmt.Println(fi38)
                                            } else {
                                                var v43 uint = 4294967295
                                                if (42 != 9223372036854775807) {
                                                    v43 = (^(v43))
                                                } else {
                                                    v4 -= v4
                                                }
                                                var v44 string = ("a" + "foo bar")
                                                var v45 float32 = ((3.14159 - (float32(float64(3.14159)) + 1e10)) * (1.5 * 1.5))
                                                v32 -= (v32 / ((v32) | 1))
                                                _ = v43
                                                _ = v44
                                                _ = v45
                                            }
                                        } else {
                                            v1 = append(v1, 2147483647)
                                        }
                                    default:
                                        var v46 int16 = 0
                                        _ = v46
                                    }
                                    if false {
                                        v47 := GenS4{F0: 42, F1: 240}
                                        var v48 float64 = 0.0
                                        _ = v47
                                        _ = v48
                                    } else {
                                        fmt.Println(fi38, fi39)
                                        v40 += (v40 & v40)
                                    }
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v49 uint = 2
                                        _ = v49
                                    default:
                                        var v50 uint8 = 2
                                        _ = v50
                                    }
                                    v31 = (v31 * (v32 >> ((18446744073709551615) & 63)))
                                }
                                if (v31 != v31) {
                                    if (21 < 255) {
                                        var v51 uint64 = (v4 + (^(v4)))
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v51
                                    } else {
                                        v52 := GenS1{F0: 83, F1: 127, F2: -812}
                                        v53 := GenS1{F0: 0, F1: 1, F2: 42}
                                        _ = v52
                                        _ = v53
                                    }
                                } else {
                                    switch "hello" {
                                    case "hello":
                                        v32 += int8(fi38)
                                    case "\n\t":
                                        fmt.Println(v4, v31)
                                    case "":
                                        v4 -= (v4 << ((489) & 63))
                                    default:
                                        v4 = (^(v4))
                                    }
                                    _ = len(v3)
                                }
                                _ = len(v1)
                            }
                        }
                    default:
                        v40 += (v40 | (v40 / ((0) | 1)))
                        v54 := GenS4{F0: 1, F1: 0}
                        v40 += (v40 | (v40 | 71))
                        _ = v54
                    }
                    _ = v40
                default:
                    v4 -= v4
                }
                _ = fi39
            }
            v3[1] = true
            _ = fi38
        }
        if (80 < 2) {
            var v55 uint = 42
            if (-855 > int32(v55)) {
                v1 = append(v1, 128)
                v3[1] = (!false)
                for fi56 := 0; fi56 < 16; fi56++ {
                    var v57 int32 = -1
                    v57 += v57
                    _ = fi56
                    _ = v57
                }
            } else {
                v58 := v3[475]
                _ = v58
            }
            if ((1e10 - float64(0.1)) != (0.1 * (float64(-2.25) + 3.14159))) {
                for fi59 := 0; fi59 < 1; fi59++ {
                    if (!(32767 != 32767)) {
                        var v60 uint = 1
                        _ = len(v2)
                        if false {
                            for fi61 := 0; fi61 < 20; fi61++ {
                                v4 += (v4 % ((v4) | 1))
                                v62 := GenS1{F0: 65535, F1: 1, F2: 145}
                                _ = fi61
                                _ = v62
                            }
                            switch 4294967295 {
                            case -53:
                                switch (fi59 ^ fi59) {
                                case 1:
                                    fmt.Println(v55, v32, v32)
                                case 128:
                                    v55 -= v55
                                default:
                                    v31 -= v31
                                }
                            case -290:
                                for fi63 := 0; fi63 < 4; fi63++ {
                                    if false {
                                        v60 = 2
                                    } else {
                                        v55 = v60
                                        if false {
                                            _ = len(v1)
                                            if true {
                                                if false {
                                                    _ = len(v1)
                                                    var v64 float32 = (((3.14159 * -2.25) + (float32(3.14159) + ((3.14159 - (0.0 - 1e10)) + 1.5))) * 1e10)
                                                    var v65 uint16 = 2
                                                    _ = v64
                                                    _ = v65
                                                } else {
                                                    _ = len(v1)
                                                    switch (fi59 & fi63) {
                                                    case 2147483647:
                                                        v4 += 18446744073709551615
                                                    }
                                                }
                                                var v66 int64 = -144
                                                fmt.Println(v55, v4, v55)
                                                _ = v66
                                            }
                                            fmt.Println(fi63, v31, v55)
                                        }
                                    }
                                    fmt.Println(v31, v31)
                                    if (1.5 < 0.1) {
                                        v67 := GenS5{F0: 754, F1: (v4 + 1), F2: 2}
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if (!false) {
                                            var v68 string = ""
                                            var v69 string = "foo bar"
                                            _ = v68
                                            _ = v69
                                        } else {
                                            var v70 int64 = 1
                                            v70 += v70
                                            switch (fi63 << ((v60) & 63)) {
                                            case 837:
                                                fmt.Println(v4, v70)
                                            }
                                            _ = v70
                                        }
                                        _ = v67
                                    }
                                    var v71 float32 = 3.14159
                                    if (127 >= fi59) {
                                        var v72 int64 = 0
                                        _ = v72
                                    }
                                    _ = fi63
                                    _ = v71
                                }
                            case -1:
                                if (true || false) {
                                    v3[625] = ((-(v32)) < -1)
                                } else {
                                    if false {
                                        v73 := GenS2{F0: 4294967295, F1: v31}
                                        _ = v73
                                    }
                                    if (!false) {
                                        v4 -= 42
                                    }
                                    var v74 float32 = 3.14159
                                    _ = v74
                                }
                            default:
                                v55 = v60
                            }
                            switch (fi59 & -783) {
                            case 128:
                                fmt.Println(v60, fi59)
                            default:
                                _ = len(v3)
                            }
                        }
                        _ = v60
                    }
                    v75 := GenS6{F0: 818, F1: 2147483647}
                    if true {
                        v55 -= v55
                        v4 = v4
                        var v76 float32 = 0.0
                        _ = v76
                    } else {
                        fmt.Println(fi59)
                    }
                    _ = fi59
                    _ = v75
                }
                for fi77 := 0; fi77 < 16; fi77++ {
                    v78 := GenS6{F0: 2, F1: 0}
                    _ = fi77
                    _ = v78
                }
                if true {
                    var v79 uint32 = 697
                    _ = v79
                } else {
                    fmt.Println(v55)
                    v80 := GenS2{F0: 4, F1: v32}
                    var v81 uint = 293
                    _ = v80
                    _ = v81
                }
            }
            _ = v55
        }
    } else {
        switch ("hello" + "") {
        case "\n\t":
            var v82 int = 32767
            _ = v82
        case "foo bar":
            v83 := GenS1{F0: 255, F1: 255, F2: -996}
            _ = v83
        }
        if false {
            v4 += (v4 - v4)
        }
        for fi84 := 0; fi84 < 14; fi84++ {
            v3[65535] = false
            _ = fi84
        }
    }
    v85 := GenS7{F0: 471}
    v4 = v4
    v32 = (v31 & (-(v32)))
    v4 -= v4
    if false {
        if (1 != 42) {
            v86 := GenS8{F0: (873 >= 856)}
            _ = v86
        } else {
            var v87 int64 = 1
            _ = v87
        }
        if true {
            fmt.Println(v31, v4, v4)
            var v88 int16 = -1
            _ = v88
        }
    }
    v89 := v2[0]
    if true {
        var v90 int32 = -1
        _ = v90
    }
    if true {
        var v91 string = ("\n\t" + "")
        _ = v91
    } else {
        _ = len(v2)
        if ("\n\t" == "\n\t") {
            var v92 float32 = float32(0.0)
            _ = v92
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v85
    _ = v89
}
