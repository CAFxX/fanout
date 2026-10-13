package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 bool; F2 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 int8; F2 int16 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 uint8; F1 uint32; F2 int }
func genFn0(p0 int8, p1 int32) int8 {
    var v6 uint = 18446744073709551615
    if ((p0 | int8(v6)) >= -45) {
        var v7 bool = true
        if v7 {
            switch "" {
            case "a":
                if v7 {
                    for fi8 := 0; fi8 < 5; fi8++ {
                        if v7 {
                            p1 = (p1 / (((p1 + 6)) | 1))
                            v9 := GenS1{F0: p0}
                            v10 := []int{(fi8 << ((v6) & 63)), (fi8 - fi8)}
                            _ = v10
                            _ = v9
                        } else {
                            if (v7 && false) {
                                var v11 uint = (^(v6))
                                v12 := GenS1{F0: int8(fi8)}
                                if (32767 >= 1) {
                                    for fi13 := 0; fi13 < 9; fi13++ {
                                        v14 := []int{(fi8 << ((255) & 63)), -9223372036854775808, fi13, (fi13 / (((fi13 * (fi8 | (fi8 / ((2147483647) | 1))))) | 1))}
                                        var v15 int64 = 42
                                        var v16 uint16 = 255
                                        _ = fi13
                                        _ = v14
                                        _ = v15
                                        _ = v16
                                    }
                                }
                                _ = v11
                                _ = v12
                            }
                            v17 := []int{1, fi8}
                            _ = v17
                        }
                        v6 = v6
                        _ = fi8
                    }
                } else {
                    switch "foo bar" {
                    case "a":
                        var v18 int = -358
                        _ = v18
                    default:
                        if (false && true) {
                            if false {
                                switch "" {
                                case "\n\t":
                                    fmt.Println(p0)
                                case "a":
                                    for fi19 := 0; fi19 < 13; fi19++ {
                                        var v20 float64 = float64(0.0)
                                        _ = fi19
                                        _ = v20
                                    }
                                }
                            } else {
                                for fi21 := 0; fi21 < 4; fi21++ {
                                    p0 = p0
                                    for fi22 := 0; fi22 < 15; fi22++ {
                                        switch fi22 {
                                        case -2147483648:
                                            var v23 float64 = (1.5 + 0.1)
                                            _ = v23
                                        case 0:
                                            if v7 {
                                                v6 = v6
                                                v6 -= v6
                                            }
                                        case 32767:
                                            var v24 uint = (v6 >> ((106) & 63))
                                            _ = v24
                                        }
                                        _ = fi22
                                    }
                                    p1 = 42
                                    _ = fi21
                                }
                            }
                            v25 := GenS2{F0: 912, F1: v7, F2: 2147483647}
                            var v26 bool = v7
                            _ = v25
                            _ = v26
                        }
                    }
                    v27 := GenS1{F0: (^(p0))}
                    fmt.Println(v7)
                    _ = v27
                }
            default:
                v28 := GenS3{F0: (0 != 42)}
                _ = v28
            }
            v29 := GenS3{F0: (127 != 4294967295)}
            _ = v29
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return (p0 | (p0 - p0))
}
func genFn1(p0 int8) uint {
    var v30 uint8 = 42
    p0 = (-(p0))
    v31 := GenS4{F0: 1}
    v32 := []int{-9223372036854775808, -363, -1}
    _ = p0
    _ = v30
    _ = v31
    _ = v32
    return 4294967295
}

func main() {
    var v1 uint32 = 2
    var v2 float64 = 3.14159
    v3 := []string{("\n\t" + (("" + "foo bar") + "hello")), ("foo bar" + "a"), ("" + "foo bar"), ("\n\t" + "")}
    var v4 float64 = (((v2 + v2) * (v2 / (3.14159 - v2))) / float64(-2.25))
    var v5 uint16 = 121
    switch "a" {
    case "a":
        switch "" {
        case "":
            v33 := GenS1{F0: genFn0(100, -2147483648)}
            _ = v33
        case "foo bar":
            fmt.Println(v4, v2, v2)
        case "a":
            fmt.Println(v4, v4, v4)
        }
    default:
        v2 = v2
    }
    v34 := GenS4{F0: genFn0(-128, -1)}
    v4 += v2
    for fi35 := 0; fi35 < 1; fi35++ {
        var v36 uint8 = 128
        _ = fi35
        _ = v36
    }
    switch "" {
    case "":
        v37 := GenS1{F0: -36}
        _ = v37
    case "hello":
        v3 = append(v3, "a")
    case "a":
        if false {
            v1 = (v1 + (v1 & v1))
            v2 -= (float64(-2.25) - (1.5 - v4))
            for fi38 := 0; fi38 < 15; fi38++ {
                var v39 int = fi38
                v40 := GenS2{F0: fi38, F1: true, F2: -856}
                _ = fi38
                _ = v39
                _ = v40
            }
        }
    }
    var v41 int16 = -1
    v5 += 560
    fmt.Println(v41, v2)
    _ = len(v3)
    _ = len(v3)
    var v42 int = -1
    switch "" {
    case "hello":
        v3 = append(v3, "a")
    case "a":
        if false {
            if false {
                var v43 uint = uint(v1)
                if (true || false) {
                    fmt.Println(v5, v41, v5)
                    fmt.Println(v5)
                }
                for fi44 := 0; fi44 < 8; fi44++ {
                    v4 -= (1e10 - v4)
                    var v45 float32 = (0.0 + (1e10 - -2.25))
                    _ = fi44
                    _ = v45
                }
                _ = v43
            }
            var v46 uint16 = v5
            _ = v46
        }
    default:
        var v47 uint32 = 42
        _ = v47
    }
    if false {
        for fi48 := 0; fi48 < 9; fi48++ {
            v2 = (v2 / 1.0)
            for fi49 := 0; fi49 < 1; fi49++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi49
            }
            _ = fi48
        }
    }
    v5 += 1
    if true {
        for fi50 := 0; fi50 < 2; fi50++ {
            if (!(4294967295 == genFn1(-128))) {
                v41 -= (^(v41))
                v51 := GenS5{F0: 4294967295, F1: 42, F2: (v41 % (((v41 - v41)) | 1))}
                _ = v51
            } else {
                fmt.Println(v5, v1)
            }
            if (-1 >= v41) {
                switch fi50 {
                case -2147483648:
                    if true {
                        if true {
                            switch v42 {
                            case 477:
                                v52 := GenS5{F0: 0, F1: 42, F2: -162}
                                _ = v52
                            }
                            var v53 uint64 = 2
                            _ = v53
                        }
                    } else {
                        switch 168 {
                        case 127:
                            fmt.Println(v2)
                        case 9223372036854775807:
                            _ = len(v3)
                        default:
                            var v54 int64 = 1
                            _ = v54
                        }
                        if false {
                            if len(v3) > 0 {
                                v3[0] = "\n\t"
                            }
                            for fi55 := 0; fi55 < 3; fi55++ {
                                if true {
                                    var v56 float64 = v4
                                    _ = v56
                                }
                                _ = fi55
                            }
                            var v57 float32 = 0.0
                            _ = v57
                        }
                        fmt.Println(fi50, v4, v42)
                    }
                case 9223372036854775807:
                    v58 := GenS6{F0: v41, F1: 255}
                    _ = v58
                case 771:
                    v41 += 1000
                }
            } else {
                if false {
                    _ = len(v3)
                    v41 = 42
                    if ("" > "") {
                        for fi59 := 0; fi59 < 11; fi59++ {
                            v41 = (v41 | (v41 ^ (v41 & v41)))
                            switch (-(v42)) {
                            case -728:
                                _ = len(v3)
                            case -9223372036854775808:
                                v1 += (v1 / ((v1) | 1))
                            default:
                                var v60 uint8 = 2
                                _ = v60
                            }
                            if (v1 < 0) {
                                var v61 int = 110
                                var v62 int64 = 9223372036854775807
                                _ = v61
                                _ = v62
                            } else {
                                switch "" {
                                case "hello":
                                    v1 = 2147483648
                                case "foo bar":
                                    var v63 string = (("hello" + "foo bar") + "foo bar")
                                    _ = v63
                                default:
                                    var v64 uint64 = 42
                                    _ = v64
                                }
                                var v65 int64 = 1
                                _ = v65
                            }
                            _ = fi59
                        }
                        v5 -= uint16(v1)
                        var v66 int8 = genFn0(100, -1)
                        _ = v66
                    } else {
                        var v67 uint16 = v5
                        for fi68 := 0; fi68 < 10; fi68++ {
                            v67 += 0
                            switch 9223372036854775807 {
                            case 32767:
                                if len(v3) > 0 {
                                    v3[0] = ""
                                }
                            case 1:
                                fmt.Println(v1, v42)
                            case 266:
                                var v69 uint8 = 255
                                _ = v69
                            default:
                                v70 := GenS1{F0: -6}
                                _ = v70
                            }
                            if len(v3) > 0 {
                                v3[0] = ("hello" + ("\n\t" + "a"))
                            }
                            _ = len(v3)
                            _ = fi68
                        }
                        _ = v67
                    }
                } else {
                    for fi71 := 0; fi71 < 8; fi71++ {
                        if ((fi71 / (((^(fi50))) | 1)) <= 4294967295) {
                            switch (("a" + "\n\t") + "a") {
                            case "a":
                                v42 -= v42
                            case "\n\t":
                                v4 += v4
                            }
                            v3 = append(v3, "a")
                            v72 := GenS7{F0: 0, F1: 65535, F2: fi71}
                            v4 += float64((0.1 * 0.0))
                            v73 := GenS7{F0: 224, F1: v1, F2: 976}
                            _ = v72
                            _ = v73
                        } else {
                            _ = len(v3)
                        }
                        var v74 uint32 = (v1 / ((v1) | 1))
                        _ = fi71
                        _ = v74
                    }
                    v1 -= v1
                    if true {
                        v4 = v2
                        var v75 uint64 = 42
                        _ = v75
                    }
                }
            }
            v42 -= (fi50 + (fi50 - 0))
            _ = fi50
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
