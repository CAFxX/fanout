package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int; F2 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 int16; F2 int16 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 string; F1 uint8; F2 uint }
type GenS9 struct { F0 uint; F1 uint8; F2 uint16 }

func main() {
    var v1 uint32 = 1
    var v2 int64 = 0
    var v3 uint64 = 18446744073709551615
    var v4 uint64 = v3
    v5 := make(map[int]bool)
    var v6 uint16 = 856
    v3 -= uint64(v1)
    v6 = (v6 ^ (v6 << ((792) & 63)))
    var v7 uint32 = v1
    v8 := []int{9223372036854775807, 186, 9223372036854775807, 0}
    for fi9 := 0; fi9 < 18; fi9++ {
        switch "\n\t" {
        case "":
            if ((!true) && ((-1 == 32767) && false)) {
                if ((v6 * v6) <= v6) {
                    if true {
                        for fi10 := 0; fi10 < 3; fi10++ {
                            v8 = append(v8, 9223372036854775807)
                            _ = fi10
                        }
                        for fi11 := 0; fi11 < 17; fi11++ {
                            v12 := GenS1{F0: ("foo bar" + "foo bar"), F1: 0}
                            _ = fi11
                            _ = v12
                        }
                        if (!false) {
                            fmt.Println(v2)
                            v13 := GenS1{F0: "", F1: 18446744073709551615}
                            v6 = 862
                            _ = v13
                        }
                    } else {
                        switch "" {
                        case "":
                            v7 -= v7
                        case "foo bar":
                            for fi14 := 0; fi14 < 17; fi14++ {
                                v4 = 4294967295
                                v15 := GenS2{F0: false}
                                _ = fi14
                                _ = v15
                            }
                        case "a":
                            if ((v4 * (v4 & (v3 % ((757) | 1)))) <= (v4 / ((v3) | 1))) {
                                switch (fi9 - fi9) {
                                case 32767:
                                    fmt.Println(v7)
                                }
                                v16 := GenS1{F0: "\n\t", F1: 181}
                                v6 -= 0
                                _ = v16
                            }
                        }
                        v5[42] = true
                    }
                } else {
                    if true {
                        v17 := GenS2{F0: true}
                        for fi18 := 0; fi18 < 13; fi18++ {
                            switch "hello" {
                            case "":
                                switch fi18 {
                                case 42:
                                    var v19 int16 = -517
                                    _ = v19
                                case 128:
                                    fmt.Println(fi18)
                                case -9223372036854775808:
                                    delete(v5, (^(fi9)))
                                default:
                                    v4 -= (v4 / ((824) | 1))
                                }
                            default:
                                _ = len(v5)
                            }
                            fmt.Println(v2)
                            _ = fi18
                        }
                        if true {
                            v7 += (v1 ^ v7)
                        } else {
                            fmt.Println(v3, v4, v7)
                            var v20 uint16 = v6
                            for fi21 := 0; fi21 < 10; fi21++ {
                                if false {
                                    if false {
                                        v2 += (v2 ^ (v2 << ((65535) & 63)))
                                    }
                                }
                                if (366 <= 348) {
                                    v22 := GenS1{F0: "a", F1: 65535}
                                    switch (fi21 - 32768) {
                                    case 128:
                                        v23 := GenS3{F0: 1, F1: fi9, F2: 0}
                                        _ = v23
                                    }
                                    v24 := GenS4{F0: 42}
                                    _ = v22
                                    _ = v24
                                } else {
                                    v20 = (v6 | 65535)
                                    v20 = 722
                                    var v25 uint8 = 1
                                    if len(v8) > 0 {
                                        v8[0] = 4294967295
                                    }
                                    if (533 > v1) {
                                        v26 := GenS4{F0: (v25 * v25)}
                                        v7 = (v7 * v1)
                                        _ = v26
                                    }
                                    _ = v25
                                }
                                _ = fi21
                            }
                            _ = v20
                        }
                        _ = v17
                    } else {
                        fmt.Println(fi9, v7)
                        v27 := v5[(fi9 ^ -9223372036854775808)]
                        _ = v27
                    }
                    fmt.Println(v3, v1)
                    if ((v6 - v6) < v6) {
                        v28 := GenS5{F0: ((fi9 + -2147483648) != int(v6)), F1: 839, F2: int16(v6)}
                        for fi29 := 0; fi29 < 6; fi29++ {
                            v3 += v4
                            v7 = (v7 | 1)
                            v6 += (v6 | uint16(v1))
                            _ = fi29
                        }
                        _ = v28
                    }
                }
                v1 += 0
            }
        case "foo bar":
            v2 -= -143
        case "hello":
            v7 = 2
        default:
            v8 = append(v8, (fi9 << ((42) & 63)))
        }
        _ = fi9
    }
    v7 = (v1 << ((2) & 63))
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 bool = true
        _ = fi30
        _ = v31
    }
    var v32 float64 = (0.0 - 1e10)
    switch 677 {
    case 42:
        fmt.Println(v1)
    case 127:
        v32 += (1.5 - v32)
    }
    var v33 int = 358
    if len(v8) > 0 {
        v8[0] = v33
    }
    v3 = (v4 * v4)
    switch (v33 & -1) {
    case 87:
        switch v33 {
        case 2147483647:
            v6 = v6
        default:
            fmt.Println(v6, v4, v3)
        }
    case -68:
        for fi34 := 0; fi34 < 9; fi34++ {
            if (484 >= 450) {
                if len(v8) > 0 {
                    v8[0] = v33
                }
                for fi35 := 0; fi35 < 13; fi35++ {
                    if (1 == 46) {
                        var v36 int64 = (^(v2))
                        v37 := GenS6{F0: 2, F1: v3}
                        v38 := GenS5{F0: false, F1: 1000, F2: 156}
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    _ = fi35
                }
            } else {
                switch (fi34 + fi34) {
                case -2147483648:
                    for fi39 := 0; fi39 < 19; fi39++ {
                        var v40 float64 = ((v32 * 1.5) / (v32 - ((v32 + 0.0) * v32)))
                        v41 := GenS6{F0: (v6 & 42), F1: (v3 * 18446744073709551615)}
                        v42 := GenS7{F0: (v3 == v4), F1: -971}
                        _ = fi39
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                case -17:
                    var v43 int32 = 0
                    _ = v43
                }
                v32 = (((((-2.25 - v32) * (v32 + v32)) / (v32 - (float64((float32(float64(0.0)) * float32(v32))) - ((v32 * (0.1 + float64(((float32(1e10) - (float32(v32) + 1.5)) * 1e10)))) + v32)))) / (v32 + (((float64(0.1) - (v32 + 1e10)) * v32) + 1e10))) - v32)
            }
            var v44 float32 = 0.1
            _ = fi34
            _ = v44
        }
    }
    switch ("a" + "hello") {
    case "":
        for fi45 := 0; fi45 < 1; fi45++ {
            v7 = v7
            var v46 int8 = 100
            v1 = (^(v7))
            _ = fi45
            _ = v46
        }
    case "\n\t":
        var v47 float32 = float32(1e10)
        _ = v47
    case "a":
        v48 := GenS8{F0: "\n\t", F1: 81, F2: 4294967295}
        _ = v48
    default:
        var v49 float64 = (v32 - 3.14159)
        _ = v49
    }
    if false {
        if len(v8) > 0 {
            v8[0] = v33
        }
        if false {
            switch v33 {
            case 2147483647:
                var v50 uint = 2
                _ = v50
            default:
                _ = len(v5)
            }
            var v51 float32 = 0.1
            _ = v51
        } else {
            if true {
                var v52 uint16 = (v6 * v6)
                v2 += v2
                fmt.Println(v32, v32)
                _ = v52
            }
            switch 127 {
            case -9223372036854775808:
                fmt.Println(v4, v7, v6)
            case 9223372036854775807:
                v32 = (v32 - 0.1)
            case 1:
                var v53 int = -754
                _ = v53
            default:
                for fi54 := 0; fi54 < 7; fi54++ {
                    delete(v5, (-(fi54)))
                    for fi55 := 0; fi55 < 19; fi55++ {
                        var v56 float32 = 3.14159
                        v1 -= v1
                        _ = fi55
                        _ = v56
                    }
                    if (!true) {
                        v3 -= 18446744073709551615
                    } else {
                        fmt.Println(v7)
                    }
                    _ = fi54
                }
            }
        }
        for fi57 := 0; fi57 < 3; fi57++ {
            fmt.Println(v4, v1, v33)
            v58 := GenS1{F0: "a", F1: 0}
            _ = fi57
            _ = v58
        }
    } else {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    for fi59 := 0; fi59 < 1; fi59++ {
        fmt.Println(v7)
        fmt.Println(v2, v33)
        for fi60 := 0; fi60 < 7; fi60++ {
            fmt.Println(fi59, fi59, v6)
            v61 := GenS9{F0: 255, F1: 42, F2: v6}
            _ = fi60
            _ = v61
        }
        _ = fi59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
