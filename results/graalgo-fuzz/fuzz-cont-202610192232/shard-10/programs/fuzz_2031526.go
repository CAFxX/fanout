package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 int32; F1 string; F2 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 1
    v3 := make(map[string]int)
    var v4 uint32 = 875
    var v5 uint32 = (v4 * v4)
    var v6 float32 = ((1e10 + 0.1) / (((0.1 + -2.25) + 0.1) / -2.25))
    if (9223372036854775807 <= 0) {
        v7 := v3["hello"]
        switch "" {
        case "a":
            v8 := []int{v2, v7, (v7 >> ((18446744073709551615) & 63)), v7}
            _ = v8
        case "hello":
            for fi9 := 0; fi9 < 16; fi9++ {
                switch ("foo bar" + "foo bar") {
                case "":
                    var v10 float64 = float64(1.5)
                    _ = v10
                case "\n\t":
                    fmt.Println(v4, v7)
                case "hello":
                    switch fi9 {
                    case -174:
                        v7 = -494
                    case 959:
                        if ((^(v4)) == 1) {
                            if true {
                                v6 = (v6 - (v6 - 1.5))
                            }
                            v11 := []int{fi9, int(v4)}
                            _ = v11
                        } else {
                            fmt.Println(v6, v6)
                            for fi12 := 0; fi12 < 11; fi12++ {
                                v13 := GenS1{F0: (1 != 42)}
                                switch fi9 {
                                case 42:
                                    v5 -= (^(v4))
                                case 4294967295:
                                    var v14 float64 = (1e10 * 3.14159)
                                    _ = v14
                                default:
                                    v15 := GenS2{F0: 1, F1: 255, F2: 594}
                                    _ = v15
                                }
                                _ = fi12
                                _ = v13
                            }
                        }
                    case -9223372036854775808:
                        var v16 int = fi9
                        _ = v16
                    default:
                        if false {
                            v5 = v5
                            if (42 == 1) {
                                v4 = (v5 << ((255) & 63))
                                var v17 int8 = 75
                                _ = v17
                            } else {
                                v5 += 328
                                if false {
                                    fmt.Println(v2)
                                } else {
                                    fmt.Println(fi9, v6)
                                    v18 := GenS2{F0: 1000, F1: 157, F2: 1}
                                    v19 := GenS2{F0: -63, F1: 18446744073709551615, F2: 801}
                                    _ = v18
                                    _ = v19
                                }
                                var v20 float32 = (v6 - (-2.25 + v6))
                                _ = v20
                            }
                            var v21 float32 = (v6 * v6)
                            _ = v21
                        }
                    }
                }
                switch v2 {
                case 9223372036854775807:
                    _ = len(v1)
                case 887:
                    v22 := v3[""]
                    _ = v22
                case 2147483647:
                    v23 := GenS3{F0: 2, F1: true, F2: (v5 + 2147483648)}
                    _ = v23
                default:
                    var v24 int64 = 0
                    _ = v24
                }
                fmt.Println(v2, fi9)
                _ = fi9
            }
        }
        _ = v7
    }
    switch "a" {
    case "hello":
        switch "a" {
        case "\n\t":
            v4 = v4
        case "hello":
            if false {
                v25 := v3[("a" + "foo bar")]
                _ = v25
            }
        case "a":
            fmt.Println(v6, v4)
        default:
            fmt.Println(v2, v6, v2)
        }
    case "a":
        v4 -= (^(v5))
    default:
        for fi26 := 0; fi26 < 1; fi26++ {
            v4 += (v4 ^ 2147483647)
            _ = fi26
        }
    }
    delete(v1, (v5 >> ((18446744073709551615) & 63)))
    if false {
        v5 += 65535
    }
    v4 = 1
    switch "hello" {
    case "":
        v5 += (v5 & v5)
    case "foo bar":
        switch "foo bar" {
        case "\n\t":
            for fi27 := 0; fi27 < 10; fi27++ {
                if false {
                    if true {
                        var v28 uint32 = (v5 ^ v5)
                        v6 -= v6
                        _ = v28
                    }
                }
                var v29 uint8 = 42
                _ = fi27
                _ = v29
            }
        case "":
            var v30 int16 = 393
            _ = v30
        }
    case "hello":
        v31 := GenS3{F0: 127, F1: (2 == 65535), F2: v4}
        _ = v31
    default:
        if (v2 <= (v2 % (((-(v2))) | 1))) {
            _ = len(v3)
            for fi32 := 0; fi32 < 12; fi32++ {
                var v33 int = v2
                if ((^(v4)) >= 0) {
                    v34 := GenS4{F0: 1, F1: 0, F2: -274}
                    _ = v34
                }
                v2 += -2147483648
                _ = fi32
                _ = v33
            }
            if (2 > 2) {
                switch (-(v2)) {
                case 1:
                    if true {
                        if false {
                            var v35 int32 = 2147483647
                            var v36 uint16 = 2
                            _ = v35
                            _ = v36
                        }
                        if true {
                            if false {
                                v5 = (v4 - (v5 & 167))
                                v1[v4] = "hello"
                            } else {
                                v4 += (v5 + 1)
                                if (false && false) {
                                    v37 := []int{(v2 * (v2 | v2)), v2}
                                    for fi38 := 0; fi38 < 14; fi38++ {
                                        v39 := GenS5{F0: -1}
                                        _ = fi38
                                        _ = v39
                                    }
                                    for fi40 := 0; fi40 < 20; fi40++ {
                                        v2 += fi40
                                        if (!(("foo bar" + ("a" + ("" + (("\n\t" + (("a" + "foo bar") + "")) + (((("" + "hello") + "foo bar") + "\n\t") + ""))))) == ("foo bar" + ("\n\t" + ("hello" + "hello"))))) {
                                            if true {
                                                for fi41 := 0; fi41 < 16; fi41++ {
                                                    var v42 uint16 = 32767
                                                    _ = fi41
                                                    _ = v42
                                                }
                                            }
                                        } else {
                                            v43 := GenS6{F0: 1}
                                            _ = v43
                                        }
                                        _ = fi40
                                    }
                                    _ = v37
                                } else {
                                    for fi44 := 0; fi44 < 12; fi44++ {
                                        v45 := GenS3{F0: 42, F1: (1 <= -100), F2: v5}
                                        switch (fi44 >> ((542) & 63)) {
                                        case -9223372036854775808:
                                            fmt.Println(v5, v5, fi44)
                                        case 887:
                                            v46 := GenS7{F0: 745, F1: (-(v2))}
                                            _ = v46
                                        case -378:
                                            var v47 int16 = -1
                                            _ = v47
                                        }
                                        _ = fi44
                                        _ = v45
                                    }
                                    var v48 string = "a"
                                    _ = v48
                                }
                                var v49 string = "a"
                                _ = v49
                            }
                            fmt.Println(v2, v2, v5)
                        }
                    } else {
                        v50 := []int{32767, v2, v2, v2}
                        _ = v50
                    }
                }
            } else {
                v4 = 2
                var v51 string = "\n\t"
                _ = v51
            }
        }
    }
    v52 := []int{(-(v2)), v2}
    fmt.Println(v4, v2, v6)
    if true {
        switch "" {
        case "\n\t":
            v4 += (v4 >> ((42) & 63))
        case "foo bar":
            for fi53 := 0; fi53 < 18; fi53++ {
                for fi54 := 0; fi54 < 3; fi54++ {
                    if true {
                        var v55 uint = 18446744073709551615
                        _ = v55
                    }
                    v5 -= (^(v5))
                    if (((v4 >> ((1) & 63)) >= 135) && false) {
                        var v56 string = "hello"
                        v57 := v1[v5]
                        v4 = v4
                        v5 -= (v5 ^ (v4 + 539))
                        var v58 float32 = v6
                        _ = v56
                        _ = v57
                        _ = v58
                    } else {
                        var v59 uint16 = 565
                        var v60 bool = ((127 != 2) || (-376 <= -532))
                        if false {
                            v6 = (v6 * ((-2.25 - 0.1) / (v6 - v6)))
                        } else {
                            if len(v52) >= 2 {
                                v52 = v52[1:len(v52)]
                            }
                            var v61 bool = false
                            if false {
                                _ = len(v3)
                                v62 := GenS8{F0: -1, F1: "foo bar", F2: 2}
                                _ = v62
                            }
                            var v63 float32 = (float32(0.1) * 1.5)
                            _ = v61
                            _ = v63
                        }
                        v52 = append(v52, (v2 >> ((2) & 63)))
                        v52 = append(v52, fi54)
                        _ = v59
                        _ = v60
                    }
                    _ = fi54
                }
                if true {
                    var v64 int8 = 127
                    _ = v64
                } else {
                    fmt.Println(v6, v4, v2)
                    v2 = (v2 << ((255) & 63))
                    if len(v52) >= 2 {
                        v52 = v52[1:len(v52)]
                    }
                }
                fmt.Println(fi53)
                _ = fi53
            }
        case "a":
            var v65 int64 = -50
            _ = v65
        }
        switch "hello" {
        case "":
            var v66 float32 = float32(1.5)
            _ = v66
        case "\n\t":
            v4 = 2147483648
        case "foo bar":
            var v67 string = "\n\t"
            _ = v67
        default:
            var v68 uint8 = 0
            _ = v68
        }
    } else {
        for fi69 := 0; fi69 < 1; fi69++ {
            v70 := GenS1{F0: false}
            _ = fi69
            _ = v70
        }
        var v71 int64 = 42
        _ = v71
    }
    v72 := GenS7{F0: 73, F1: v2}
    if ((-881 == 0) && (0 <= 2147483647)) {
        for fi73 := 0; fi73 < 5; fi73++ {
            for fi74 := 0; fi74 < 11; fi74++ {
                v6 += 1.5
                _ = fi74
            }
            fmt.Println(v5, fi73, v6)
            _ = fi73
        }
        v2 -= -1
        v6 -= v6
    } else {
        var v75 int16 = -263
        var v76 uint64 = 843
        switch 9223372036854775807 {
        case 128:
            if true {
                switch (v2 * v2) {
                case 1:
                    v2 = -873
                case 2147483647:
                    var v77 float64 = 3.14159
                    _ = v77
                case 660:
                    v76 += (v76 ^ 798)
                default:
                    v75 += (v75 >> ((561) & 63))
                }
                var v78 uint32 = 65535
                _ = v78
            } else {
                v76 -= (v76 >> ((1) & 63))
                v75 += v75
                var v79 int32 = -346
                _ = v79
            }
        case -622:
            var v80 float64 = 3.14159
            _ = v80
        default:
            v76 += v76
        }
        _ = v75
        _ = v76
    }
    v5 = (v5 | (v5 % (((v5 >> ((657) & 63))) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v72
}
