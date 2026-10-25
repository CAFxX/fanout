package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 uint64 }
type GenS7 struct { F0 bool; F1 int64 }
func genFn0(p0 uint16, p1 uint16, p2 int) string {
    var v7 bool = (0 == -32768)
    p1 = p0
    if (!v7) {
        switch "\n\t" {
        case "a":
            v8 := GenS1{F0: -1, F1: 246}
            _ = v8
        case "":
            var v9 uint16 = p1
            _ = v9
        case "foo bar":
            var v10 uint32 = 1
            _ = v10
        }
        switch (("hello" + "\n\t") + "a") {
        case "foo bar":
            v11 := GenS2{F0: 2147483647}
            _ = v11
        case "\n\t":
            var v12 int8 = 1
            _ = v12
        case "hello":
            switch -9223372036854775808 {
            case 42:
                var v13 float64 = 1e10
                _ = v13
            case 128:
                var v14 uint = 18446744073709551615
                _ = v14
            case -9223372036854775808:
                var v15 int = p2
                _ = v15
            default:
                switch ((("hello" + "foo bar") + ("" + "a")) + "\n\t") {
                case "hello":
                    if (false && false) {
                        v16 := []int{(^(p2)), p2, (-(p2)), p2}
                        _ = v16
                    }
                case "a":
                    var v17 uint8 = 128
                    _ = v17
                default:
                    var v18 bool = v7
                    _ = v18
                }
            }
        }
    } else {
        v19 := []int{}
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        p1 += 65535
        fmt.Println(p0)
        _ = fi20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "foo bar"
}
func genFn1(p0 uint64) int {
    for fi21 := 0; fi21 < 18; fi21++ {
        v22 := GenS1{F0: 786, F1: 92}
        _ = fi21
        _ = v22
    }
    v23 := GenS3{F0: p0}
    _ = p0
    _ = v23
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    var v3 uint32 = 459
    var v4 float64 = 3.14159
    var v5 int32 = -2147483648
    var v6 uint16 = 1
    v6 += (v6 & v6)
    v5 -= v5
    fmt.Println(v3, v4)
    var v24 uint8 = 128
    var v25 int8 = -128
    if false {
        fmt.Println(v5)
        var v26 int64 = 9223372036854775807
        _ = v26
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        v6 += 65535
        fmt.Println(v6, v3, v3)
        switch "\n\t" {
        case "foo bar":
            for fi28 := 0; fi28 < 1; fi28++ {
                switch 9223372036854775807 {
                case 127:
                    v4 = (v4 + (v4 / v4))
                case 861:
                    switch "hello" {
                    case "":
                        var v29 uint16 = v6
                        _ = v29
                    case "\n\t":
                        var v30 int64 = 42
                        _ = v30
                    case "hello":
                        if (987 == 2) {
                            v31 := []int{-9223372036854775808, (fi27 * genFn1(18446744073709551615))}
                            switch (fi28 * (fi28 | 4294967295)) {
                            case 0:
                                var v32 uint32 = (v3 / ((v3) | 1))
                                _ = v32
                            case 4294967295:
                                v33 := GenS4{F0: -996, F1: 1}
                                _ = v33
                            case 280:
                                var v34 float32 = 3.14159
                                _ = v34
                            default:
                                var v35 uint64 = 584
                                switch genFn1((v35 & 18446744073709551615)) {
                                case 209:
                                    var v36 string = ("\n\t" + "hello")
                                    _ = v36
                                case -2147483648:
                                    v35 = (^(v35))
                                case 32767:
                                    switch genFn1((v35 + v35)) {
                                    case 128:
                                        v4 = v4
                                    default:
                                        switch (fi28 & fi28) {
                                        case -2147483648:
                                            var v37 uint8 = v24
                                            _ = v37
                                        case -813:
                                            var v38 uint8 = v24
                                            _ = v38
                                        case 312:
                                            v5 += v5
                                        default:
                                            v39 := GenS5{F0: 1000}
                                            _ = v39
                                        }
                                    }
                                default:
                                    var v40 float64 = v4
                                    _ = v40
                                }
                                _ = v35
                            }
                            if true {
                                v41 := GenS6{F0: 2, F1: 2}
                                var v42 bool = (0 >= 42)
                                _ = v41
                                _ = v42
                            }
                            _ = v31
                        }
                    default:
                        v6 = (^(v6))
                    }
                default:
                    v43 := GenS5{F0: -684}
                    _ = v43
                }
                v44 := []int{}
                if (genFn1(uint64(v24)) > (fi27 << ((0) & 63))) {
                    v4 += v4
                    var v45 float64 = -2.25
                    var v46 uint64 = 1
                    var v47 uint16 = (^(v6))
                    _ = v45
                    _ = v46
                    _ = v47
                } else {
                    if false {
                        fmt.Println(v24)
                    }
                }
                var v48 string = "hello"
                _ = fi28
                _ = v44
                _ = v48
            }
        case "a":
            for fi49 := 0; fi49 < 14; fi49++ {
                if true {
                    v50 := GenS4{F0: -14, F1: 9223372036854775807}
                    v51 := []int{genFn1(uint64(v3)), genFn1(uint64(v3))}
                    v5 = (v5 % (((v5 ^ 2147483647)) | 1))
                    var v52 int32 = (v5 | (v5 << ((255) & 63)))
                    v53 := GenS4{F0: -74, F1: 1}
                    _ = v50
                    _ = v51
                    _ = v52
                    _ = v53
                }
                _ = fi49
            }
        }
        _ = fi27
    }
    v54 := []int{1, 817, -363, 341}
    v55 := v2[128]
    v6 += v6
    v56 := GenS6{F0: 255, F1: 93}
    v25 -= (v25 << ((255) & 63))
    v55 -= 437
    switch v55 {
    case 209:
        var v57 int32 = 661
        _ = v57
    case 856:
        if (v5 <= v5) {
            switch (v55 / ((v55) | 1)) {
            case 2147483647:
                if ((v5 << ((255) & 63)) < -1) {
                    v1[(v55 >> ((2) & 63))] = (v55 ^ (v55 >> ((42) & 63)))
                    if false {
                        var v58 string = "foo bar"
                        for fi59 := 0; fi59 < 4; fi59++ {
                            if (((-2.25 / (float32(0.1) * (0.1 - 0.0))) / 1.5) > -2.25) {
                                v24 = (v24 ^ (v24 ^ v24))
                                var v60 float64 = (0.1 * (float64(((1e10 - 0.1) - 3.14159)) + ((v4 / (v4 * v4)) - -2.25)))
                                _ = v60
                            }
                            var v61 int8 = v25
                            for fi62 := 0; fi62 < 13; fi62++ {
                                v54 = append(v54, (v55 + 9223372036854775807))
                                switch ("hello" + v58) {
                                case "hello":
                                    v63 := GenS6{F0: 2, F1: 4294967295}
                                    _ = v63
                                default:
                                    var v64 int = fi59
                                    _ = v64
                                }
                                v65 := GenS6{F0: uint(v3), F1: 942}
                                _ = fi62
                                _ = v65
                            }
                            _ = fi59
                            _ = v61
                        }
                        v66 := GenS6{F0: 255, F1: 611}
                        _ = v58
                        _ = v66
                    } else {
                        if (true || true) {
                            for fi67 := 0; fi67 < 11; fi67++ {
                                var v68 string = "a"
                                for fi69 := 0; fi69 < 16; fi69++ {
                                    if ((((((((1e10 * 1e10) - 1e10) + 3.14159) / (1e10 * 3.14159)) + 0.1) / -2.25) + (((3.14159 + 1e10) / 1.5) - 0.0)) < -2.25) {
                                        v55 = (fi69 & fi67)
                                    }
                                    _ = len(v54)
                                    v25 = (v25 << ((18446744073709551615) & 63))
                                    _ = fi69
                                }
                                v6 += (v6 >> ((uint(v24)) & 63))
                                _ = fi67
                                _ = v68
                            }
                        } else {
                            if false {
                                for fi70 := 0; fi70 < 12; fi70++ {
                                    for fi71 := 0; fi71 < 14; fi71++ {
                                        v72 := GenS7{F0: true, F1: 0}
                                        switch (fi70 / ((v55) | 1)) {
                                        case 32767:
                                            if len(v54) >= 2 {
                                                v54 = v54[1:len(v54)]
                                            }
                                        case -9223372036854775808:
                                            var v73 uint64 = 1
                                            _ = v73
                                        case -339:
                                            if false {
                                                if true {
                                                    if len(v54) >= 2 {
                                                        v54 = v54[1:len(v54)]
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v5, fi71)
                                                v24 = (v24 & (v24 % (((v24 ^ v24)) | 1)))
                                                var v74 int8 = (v25 << ((0) & 63))
                                                _ = v74
                                            }
                                        }
                                        switch v55 {
                                        case 0:
                                            var v75 uint32 = (v3 & (v3 << ((65535) & 63)))
                                            var v76 string = "hello"
                                            v3 = 65535
                                            _ = v75
                                            _ = v76
                                        case 1:
                                            v4 = 3.14159
                                        default:
                                            var v77 float64 = v4
                                            _ = v77
                                        }
                                        _ = fi71
                                        _ = v72
                                    }
                                    _ = fi70
                                }
                                fmt.Println(v55, v55)
                                v25 = v25
                            } else {
                                v4 += v4
                                v2[(^(v55))] = -9223372036854775808
                                v4 -= v4
                            }
                            if (433 <= (v3 * 2147483647)) {
                                v78 := GenS7{F0: (!true), F1: -1}
                                _ = v78
                            }
                        }
                    }
                    var v79 float32 = 3.14159
                    _ = v79
                }
            case 9223372036854775807:
                v3 = (v3 >> ((4294967295) & 63))
            }
            v6 = (^(v6))
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v6
}
