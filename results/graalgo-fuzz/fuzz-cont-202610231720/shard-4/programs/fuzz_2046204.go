package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 bool }
type GenS2 struct { F0 int16; F1 string; F2 uint32 }
type GenS3 struct { F0 int; F1 int32; F2 int8 }
type GenS4 struct { F0 uint8; F1 string; F2 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16; F1 int8 }
func genFn0(p0 int16, p1 uint) int16 {
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: fi6, F1: fi6, F2: false}
        _ = fi6
        _ = v7
    }
    v8 := GenS2{F0: (p0 / (((p0 * p0)) | 1)), F1: "hello", F2: 2}
    v9 := GenS3{F0: 42, F1: 379, F2: 100}
    if true {
        v10 := GenS1{F0: 0, F1: 128, F2: false}
        p0 = 1
        var v11 int64 = 0
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return (^(p0))
}

func main() {
    v1 := []string{"\n\t", "a"}
    var v2 uint = 18446744073709551615
    v3 := make(map[int]uint32)
    var v4 uint8 = 0
    v5 := []int{-386}
    if (3.14159 == 3.14159) {
        if (0 < -290) {
            v2 = (v2 / ((uint(v4)) | 1))
            v12 := GenS4{F0: v4, F1: "a", F2: 2}
            _ = v12
        }
        switch -526 {
        case 32767:
            var v13 uint32 = 42
            _ = v13
        case 250:
            if false {
                switch "" {
                case "":
                    v3[42] = 42
                case "\n\t":
                    v14 := GenS1{F0: -951, F1: -479, F2: true}
                    _ = v14
                default:
                    if (142 < (v2 ^ v2)) {
                        switch 489 {
                        case 1:
                            var v15 int32 = -2147483648
                            _ = v15
                        case -741:
                            if (float64(-2.25) <= 1e10) {
                                var v16 uint = (v2 * v2)
                                var v17 int32 = 2147483647
                                _ = v16
                                _ = v17
                            } else {
                                v3[2147483647] = 2
                                fmt.Println(v2, v2)
                                v4 -= v4
                            }
                        case -128:
                            v4 -= (v4 % (((^(v4))) | 1))
                        }
                        if (123 <= 0) {
                            fmt.Println(v4)
                            if ((2 <= v2) && false) {
                                for fi18 := 0; fi18 < 6; fi18++ {
                                    for fi19 := 0; fi19 < 6; fi19++ {
                                        var v20 uint8 = 127
                                        fmt.Println(v4, v4, v2)
                                        _ = fi19
                                        _ = v20
                                    }
                                    _ = fi18
                                }
                            }
                            if (v2 != v2) {
                                var v21 int = -1
                                fmt.Println(v2, v21)
                                switch "" {
                                case "foo bar":
                                    for fi22 := 0; fi22 < 1; fi22++ {
                                        switch 32768 {
                                        case 2147483647:
                                            v21 = fi22
                                        default:
                                            v23 := GenS3{F0: fi22, F1: -1, F2: -1}
                                            _ = v23
                                        }
                                        _ = fi22
                                    }
                                case "":
                                    v4 -= (v4 ^ (^(v4)))
                                case "hello":
                                    for fi24 := 0; fi24 < 3; fi24++ {
                                        fmt.Println(v4, v2, v21)
                                        var v25 int64 = -1
                                        _ = fi24
                                        _ = v25
                                    }
                                }
                                _ = v21
                            }
                        }
                    }
                }
                var v26 bool = ((-135 != 1) && false)
                _ = v26
            } else {
                var v27 float32 = float32(3.14159)
                _ = v27
            }
        default:
            for fi28 := 0; fi28 < 9; fi28++ {
                switch "hello" {
                case "":
                    for fi29 := 0; fi29 < 11; fi29++ {
                        if (0.0 == 1e10) {
                            var v30 uint = v2
                            var v31 uint = v2
                            switch (fi28 ^ (^(fi29))) {
                            case -716:
                                switch (fi28 - fi28) {
                                case -9223372036854775808:
                                    v4 = 1
                                    fmt.Println(v31, v30)
                                default:
                                    fmt.Println(v2, fi29)
                                }
                            case 32768:
                                var v32 uint8 = 128
                                _ = v32
                            case 4294967295:
                                fmt.Println(fi29)
                            }
                            _ = v30
                            _ = v31
                        } else {
                            var v33 int8 = 1
                            v34 := GenS2{F0: 42, F1: ("foo bar" + "foo bar"), F2: 4294967295}
                            var v35 string = ("" + "")
                            var v36 uint64 = 4294967295
                            _ = v33
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        _ = fi29
                    }
                case "\n\t":
                    v2 += 524
                }
                v2 += v2
                _ = fi28
            }
        }
    } else {
        var v37 uint16 = 32767
        v5 = append(v5, 584)
        _ = v37
    }
    switch "\n\t" {
    case "":
        switch "a" {
        case "foo bar":
            for fi38 := 0; fi38 < 7; fi38++ {
                for fi39 := 0; fi39 < 8; fi39++ {
                    if true {
                        var v40 uint64 = 0
                        v41 := GenS5{F0: -336}
                        var v42 string = ("foo bar" + "a")
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                    _ = fi39
                }
                v43 := GenS6{F0: -280}
                v2 -= v2
                _ = fi38
                _ = v43
            }
        case "\n\t":
            v4 += 128
        case "a":
            for fi44 := 0; fi44 < 13; fi44++ {
                _ = len(v5)
                _ = fi44
            }
        }
    case "\n\t":
        var v45 int32 = 1
        _ = v45
    default:
        fmt.Println(v4)
    }
    for fi46 := 0; fi46 < 1; fi46++ {
        if len(v1) > 0 {
            v1[0] = "a"
        }
        if (-32768 == -32768) {
            if ((1.5 + -2.25) > -2.25) {
                v2 += 4294967295
            }
            var v47 int32 = 2147483647
            for fi48 := 0; fi48 < 1; fi48++ {
                v4 += (v4 << ((1) & 63))
                if (1 == (fi48 - fi46)) {
                    var v49 float32 = -2.25
                    switch "" {
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "a":
                        fmt.Println(fi48, v2, fi46)
                    case "foo bar":
                        v47 += (^(v47))
                    default:
                        var v50 string = "\n\t"
                        _ = v50
                    }
                    _ = v49
                } else {
                    v47 -= 459
                    switch "" {
                    case "hello":
                        v4 = v4
                    default:
                        v51 := v3[fi48]
                        _ = v51
                    }
                }
                _ = fi48
            }
            _ = v47
        } else {
            var v52 uint32 = 0
            _ = len(v5)
            var v53 bool = true
            _ = v52
            _ = v53
        }
        var v54 string = "foo bar"
        _ = fi46
        _ = v54
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    var v55 int32 = 42
    switch ("hello" + "a") {
    case "a":
        _ = len(v3)
    default:
        if len(v5) > 0 {
            v5[0] = -218
        }
    }
    if ((0.1 - 3.14159) > (float32((0.0 + 3.14159)) + ((0.0 * (float32((0.0 - 3.14159)) * float32((3.14159 * 0.1)))) + 1e10))) {
        var v56 uint = v2
        _ = v56
    }
    _ = len(v3)
    v57 := GenS7{F0: 2, F1: 100}
    fmt.Println(v2, v55, v4)
    for fi58 := 0; fi58 < 1; fi58++ {
        var v59 uint16 = 42
        _ = fi58
        _ = v59
    }
    switch 42 {
    case 165:
        for fi60 := 0; fi60 < 15; fi60++ {
            v55 = v55
            _ = fi60
        }
    default:
        v61 := GenS1{F0: 32768, F1: 127, F2: false}
        _ = v61
    }
    var v62 uint32 = 407
    for fi63 := 0; fi63 < 14; fi63++ {
        v64 := GenS3{F0: 0, F1: (v55 >> ((42) & 63)), F2: 0}
        _ = fi63
        _ = v64
    }
    var v65 uint32 = (v62 / ((v62) | 1))
    if false {
        _ = len(v3)
        var v66 uint16 = 32767
        _ = v66
    } else {
        _ = len(v1)
        if (false || true) {
            for fi67 := 0; fi67 < 9; fi67++ {
                if len(v5) > 0 {
                    v5[0] = (fi67 % ((fi67) | 1))
                }
                _ = fi67
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v68 int64 = 473
            _ = v68
        }
    }
    for fi69 := 0; fi69 < 20; fi69++ {
        var v70 int16 = -1
        v71 := GenS5{F0: -9223372036854775808}
        _ = fi69
        _ = v70
        _ = v71
    }
    var v72 float64 = (0.0 + 3.14159)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v62:", v62)
    fmt.Println("v65:", v65)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v57
    _ = v62
    _ = v65
    _ = v72
}
