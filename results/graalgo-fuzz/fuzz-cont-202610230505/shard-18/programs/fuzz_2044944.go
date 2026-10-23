package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint16; F1 int; F2 int32 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 uint8; F2 int16 }
type GenS8 struct { F0 int64; F1 int }
type GenS9 struct { F0 uint8 }
func genFn0(p0 uint, p1 uint64, p2 uint32) uint16 {
    v7 := []int{2147483647, 0}
    if false {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    v8 := GenS1{F0: 951, F1: p2, F2: 9223372036854775807}
    switch -868 {
    case -574:
        _ = len(v7)
    case -789:
        p1 -= 4294967295
    case 1:
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := GenS1{F0: 2147483647, F1: p2, F2: 1}
            if ((-1 > -1) && ((p2 >> ((p0) & 63)) >= 42)) {
                v11 := GenS1{F0: -2147483648, F1: (p2 << (((p0 & 480)) & 63)), F2: 9223372036854775807}
                _ = v11
            } else {
                v7 = append(v7, fi9)
                var v12 int64 = 42
                var v13 int64 = v12
                _ = v12
                _ = v13
            }
            var v14 string = ""
            _ = fi9
            _ = v10
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 615
}
func genFn1(p0 int8, p1 int64) uint64 {
    fmt.Println(p0, p0)
    var v15 int32 = 1
    if (0.1 <= (3.14159 / float64((((-2.25 / 1e10) - (float32(3.14159) * float32(0.1))) + 1e10)))) {
        var v16 float32 = float32((((1.5 + 1.5) - (((1.5 * 0.1) / (float64((0.1 + 1.5)) - (float64(1e10) / (1e10 + ((-2.25 - (-2.25 * 3.14159)) - 3.14159))))) * (0.1 + (-2.25 + -2.25)))) + (1.5 * 1.5)))
        if (154 == 65535) {
            v17 := GenS2{F0: 371, F1: -9223372036854775808, F2: v15}
            v18 := []int{32768, 32767}
            _ = v17
            _ = v18
        }
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v15
    return 427
}

func main() {
    var v1 uint = 65535
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    var v5 uint = v1
    v6 := make(map[string]bool)
    for fi19 := 0; fi19 < 13; fi19++ {
        v1 = v5
        v1 += (v1 % ((v5) | 1))
        fmt.Println(fi19)
        _ = fi19
    }
    v20 := []int{-1, 127}
    if true {
        v1 += (^(v5))
        if false {
            if (127 == -19) {
                v1 += v5
                switch -243 {
                case 128:
                    v1 -= v1
                case -1:
                    v1 += v5
                }
                var v21 string = "hello"
                _ = v21
            }
        }
        var v22 uint64 = 298
        _ = v22
    }
    v5 += 42
    for fi23 := 0; fi23 < 7; fi23++ {
        if false {
            for fi24 := 0; fi24 < 20; fi24++ {
                var v25 uint8 = 128
                v26 := GenS1{F0: 2147483647, F1: 571, F2: 42}
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        _ = fi23
    }
    if (!false) {
        fmt.Println(v5)
    }
    v27 := GenS1{F0: 0, F1: 478, F2: 0}
    if (float64(1e10) == 0.1) {
        var v28 int32 = -2147483648
        v29 := GenS2{F0: 2, F1: 32767, F2: v28}
        _ = v28
        _ = v29
    } else {
        _ = len(v20)
    }
    v30 := GenS2{F0: 32767, F1: 1, F2: 523}
    if (false || true) {
        var v31 float32 = 0.0
        v20 = append(v20, 990)
        _ = v31
    }
    fmt.Println(v1)
    if (!true) {
        for fi32 := 0; fi32 < 6; fi32++ {
            v1 += 0
            var v33 string = "\n\t"
            if false {
                for fi34 := 0; fi34 < 8; fi34++ {
                    v5 = v1
                    v3[105] = 42
                    var v35 bool = (382 >= 328)
                    _ = fi34
                    _ = v35
                }
                if true {
                    var v36 string = (v33 + v33)
                    v37 := GenS3{F0: 0, F1: 65535}
                    _ = v36
                    _ = v37
                } else {
                    switch fi32 {
                    case 123:
                        v33 = "\n\t"
                    default:
                        v2[4294967295] = 839
                    }
                    switch fi32 {
                    case 619:
                        for fi38 := 0; fi38 < 11; fi38++ {
                            v5 -= (^(v1))
                            if len(v20) > 0 {
                                v20[0] = fi38
                            }
                            var v39 int64 = 0
                            _ = fi38
                            _ = v39
                        }
                    case 1:
                        if true {
                            v20 = append(v20, (^(fi32)))
                        } else {
                            fmt.Println(v33)
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            if false {
                                if (false && true) {
                                    fmt.Println(v1, v33, v33)
                                    var v40 bool = (0 > 0)
                                    if (65535 >= genFn0(65535, 1, 65535)) {
                                        v41 := GenS4{F0: 0}
                                        var v42 bool = (v40 || v40)
                                        _ = v41
                                        _ = v42
                                    }
                                    _ = v40
                                } else {
                                    v43 := GenS3{F0: 1, F1: uint32(v5)}
                                    if true {
                                        fmt.Println(v1, v5)
                                        v33 = "hello"
                                    } else {
                                        var v44 uint = v5
                                        v33 = ((v33 + (v33 + "hello")) + ((v33 + v33) + v33))
                                        _ = v44
                                    }
                                    if (((false && (false && (((1.5 * 0.0) - float32(((0.0 + 1e10) + -2.25))) == 0.1))) && true) && (!false)) {
                                        v5 -= 786
                                        var v45 bool = (v5 > 42)
                                        _ = v45
                                    }
                                    _ = v43
                                }
                            }
                        }
                    case -1:
                        var v46 int16 = 42
                        _ = v46
                    default:
                        v4[(fi32 + 9223372036854775807)] = 9223372036854775807
                    }
                    _ = len(v6)
                }
            } else {
                for fi47 := 0; fi47 < 1; fi47++ {
                    v33 = (((v33 + v33) + v33) + "")
                    var v48 int16 = 192
                    if true {
                        switch (("" + v33) + v33) {
                        case "hello":
                            fmt.Println(v33, v33)
                            switch ((v33 + ("foo bar" + (v33 + v33))) + v33) {
                            case "hello":
                                v49 := GenS5{F0: v48, F1: -127}
                                _ = v49
                            case "foo bar":
                                var v50 int16 = 42
                                _ = v50
                            case "\n\t":
                                var v51 uint = (v1 - (v1 + v5))
                                _ = v51
                            default:
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                            }
                        case "\n\t":
                            v33 = ("a" + (("" + "foo bar") + ("a" + (v33 + ((v33 + v33) + ("foo bar" + "\n\t"))))))
                        default:
                            switch (fi47 & (fi32 ^ fi32)) {
                            case 32768:
                                v48 = v48
                            case -748:
                                v52 := GenS1{F0: 1, F1: 71, F2: -1}
                                _ = v52
                            }
                        }
                        var v53 uint8 = 128
                        _ = v53
                    }
                    _ = fi47
                    _ = v48
                }
            }
            _ = fi32
            _ = v33
        }
        if true {
            if true {
                v54 := GenS6{F0: 908}
                var v55 float32 = 0.0
                _ = v54
                _ = v55
            }
            switch 2147483647 {
            case 4294967295:
                for fi56 := 0; fi56 < 3; fi56++ {
                    switch fi56 {
                    case 4294967295:
                        switch fi56 {
                        case -344:
                            v1 = v5
                        case -28:
                            var v57 uint = (v5 << (((^(v5))) & 63))
                            _ = v57
                        default:
                            v5 = v1
                        }
                    case 2147483647:
                        switch ("hello" + "a") {
                        case "":
                            if false {
                                var v58 int64 = -862
                                _ = v58
                            } else {
                                v59 := GenS7{F0: (-9223372036854775808 > -9223372036854775808), F1: 42, F2: 1}
                                v5 = v1
                                if (v1 < (v5 % ((v5) | 1))) {
                                    if len(v20) > 0 {
                                        v20[0] = (fi56 * fi56)
                                    }
                                    fmt.Println(v5)
                                }
                                _ = v59
                            }
                        case "a":
                            v1 = v1
                        }
                    case -621:
                        var v60 uint = (v1 / ((v1) | 1))
                        _ = v60
                    default:
                        fmt.Println(v1, v1, v1)
                    }
                    fmt.Println(v1, v1, v5)
                    _ = fi56
                }
            case -2147483648:
                v61 := GenS6{F0: 951}
                _ = v61
            }
        }
    } else {
        v5 -= 366
    }
    v62 := GenS8{F0: -9223372036854775808, F1: 126}
    for fi63 := 0; fi63 < 12; fi63++ {
        if true {
            _ = len(v2)
            if true {
                switch fi63 {
                case 0:
                    if true {
                        for fi64 := 0; fi64 < 9; fi64++ {
                            v65 := GenS4{F0: 480}
                            _ = fi64
                            _ = v65
                        }
                        v5 = 42
                        fmt.Println(v5)
                    }
                case 128:
                    var v66 string = "foo bar"
                    _ = v66
                case -200:
                    if false {
                        v67 := GenS9{F0: 127}
                        var v68 uint64 = 1
                        switch "hello" {
                        case "a":
                            switch ("a" + ("hello" + (("foo bar" + "foo bar") + "hello"))) {
                            case "foo bar":
                                v1 += 4294967295
                            }
                        default:
                            var v69 uint16 = genFn0((v1 / ((v1) | 1)), 42, 4294967295)
                            _ = v69
                        }
                        _ = v67
                        _ = v68
                    }
                default:
                    fmt.Println(fi63, fi63)
                }
            } else {
                _ = len(v6)
            }
            v70 := v2[630]
            _ = v70
        } else {
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            switch fi63 {
            case 1:
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
            }
        }
        _ = fi63
    }
    v5 += 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v62
}
