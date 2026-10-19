package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int; F2 int32 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint16 }
func genFn0(p0 uint16, p1 int, p2 string) uint16 {
    fmt.Println(p2, p1)
    v5 := []int{(p1 & p1), (p1 << ((65535) & 63))}
    var v6 int16 = -32768
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 1, F1: -1}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (p0 & (p0 << ((1) & 63)))
}

func main() {
    var v1 uint16 = 670
    v2 := []string{("" + ""), ("foo bar" + ("foo bar" + ("foo bar" + ("a" + "a"))))}
    var v3 float32 = ((((float32((((float64(0.0) - (((-2.25 * (-2.25 + float64(1.5))) / -2.25) - (-2.25 + 3.14159))) - 3.14159) + float64((1e10 * -2.25)))) / 3.14159) / 3.14159) + 0.1) * ((-2.25 - (float32(-2.25) / 0.1)) + (1e10 * ((((((-2.25 + float32(0.0)) / 0.1) + 0.1) * float32(-2.25)) + 3.14159) / ((-2.25 * (3.14159 / (-2.25 - 0.1))) + (1e10 + -2.25))))))
    var v4 bool = (!(42 > -9223372036854775808))
    if (v4 && (!v4)) {
        fmt.Println(v1, v1)
        v4 = (-1 == 42)
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 bool = v4
            if (v4 || (!(v4 || (-2147483648 > 4294967295)))) {
                var v11 int64 = 508
                _ = v11
            }
            _ = fi9
            _ = v10
        }
    }
    if (!v4) {
        v12 := GenS2{F0: (1e10 >= ((1.5 - 0.1) + ((((((v3 / float32(0.1)) - v3) - float32(0.1)) / 1.0) + v3) / v3)))}
        v2 = append(v2, ("a" + "a"))
        _ = v12
    } else {
        if v4 {
            for fi13 := 0; fi13 < 2; fi13++ {
                var v14 uint32 = 2147483647
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = len(v2)
                _ = fi13
                _ = v14
            }
        } else {
            var v15 float64 = float64(v3)
            fmt.Println(v4)
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 float32 = ((float32(((((-2.25 / (0.1 / 3.14159)) - (v15 + v15)) / float64(v3)) / 0.1)) * -2.25) / (v3 + v3))
                _ = fi16
                _ = v17
            }
            _ = v15
        }
        if true {
            if len(v2) > 0 {
                v2[0] = "a"
            }
            for fi18 := 0; fi18 < 11; fi18++ {
                v19 := GenS1{F0: fi18, F1: 42}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 20; fi20++ {
                fmt.Println(v1, v4, v1)
                _ = fi20
            }
        } else {
            fmt.Println(v1, v4)
            for fi21 := 0; fi21 < 6; fi21++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v22 float64 = ((-2.25 * ((1.5 - 1.5) * ((0.1 + (0.0 * 1.5)) - 0.1))) / 0.1)
                var v23 uint16 = genFn0(v1, (fi21 - 42), "")
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 int = 32767
            _ = v24
        }
    }
    var v25 uint32 = 0
    var v26 uint32 = (^(v25))
    if (!(-2147483648 > 42)) {
        var v27 float64 = ((0.1 / (3.14159 + 1e10)) + 0.0)
        _ = v27
    } else {
        _ = len(v2)
    }
    switch -1 {
    case -129:
        v28 := GenS2{F0: (1e10 <= -2.25)}
        _ = v28
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        v1 = (v1 >> ((uint(v25)) & 63))
        for fi30 := 0; fi30 < 13; fi30++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi30
        }
        _ = fi29
    }
    switch -702 {
    case 397:
        for fi31 := 0; fi31 < 9; fi31++ {
            if (uint64(v25) != 0) {
                var v32 int32 = 42
                for fi33 := 0; fi33 < 13; fi33++ {
                    v34 := GenS3{F0: (^(v26)), F1: 4294967295, F2: v32}
                    switch fi31 {
                    case 32768:
                        if v4 {
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                            var v35 uint8 = 2
                            _ = v35
                        } else {
                            switch ("a" + "foo bar") {
                            case "hello":
                                switch "foo bar" {
                                case "":
                                    switch "" {
                                    case "foo bar":
                                        if (((v25 ^ 2147483648) != 42) && ((!false) || ((((0.1 - (((3.14159 + 1.5) * 0.0) * (0.0 * ((0.0 + 1.5) + 1e10)))) - (-2.25 / 3.14159)) * float64(3.14159)) != (((1e10 / (3.14159 + 0.1)) / 1e10) / (0.1 + 1.5))))) {
                                            if len(v2) > 0 {
                                                v2[0] = "\n\t"
                                            }
                                            var v36 uint = 18446744073709551615
                                            _ = v36
                                        } else {
                                            v3 -= v3
                                            if len(v2) > 0 {
                                                v2[0] = "a"
                                            }
                                        }
                                    case "a":
                                        if (0 < -1) {
                                            v3 -= -2.25
                                            var v37 uint32 = (v26 << ((0) & 63))
                                            v38 := GenS4{F0: int64(v1), F1: 100}
                                            v37 = (v25 + 486)
                                            _ = v37
                                            _ = v38
                                        }
                                    }
                                default:
                                    if true {
                                        fmt.Println(fi33, v1)
                                        v2 = append(v2, "hello")
                                        var v39 int64 = -9223372036854775808
                                        _ = v39
                                    }
                                }
                            case "foo bar":
                                var v40 int32 = v32
                                _ = v40
                            default:
                                var v41 float32 = 0.0
                                _ = v41
                            }
                            v4 = (true || true)
                        }
                    }
                    _ = fi33
                    _ = v34
                }
                _ = v32
            }
            if v4 {
                if (42 != int16(fi31)) {
                    var v42 int16 = 1
                    var v43 int64 = 42
                    for fi44 := 0; fi44 < 5; fi44++ {
                        switch "a" {
                        case "\n\t":
                            var v45 int32 = -139
                            _ = v45
                        case "a":
                            var v46 uint64 = 18446744073709551615
                            _ = v46
                        case "hello":
                            fmt.Println(v3, v1, fi44)
                        default:
                            var v47 uint = 42
                            _ = v47
                        }
                        fmt.Println(v26, v4, v25)
                        _ = fi44
                    }
                    _ = v42
                    _ = v43
                }
            } else {
                var v48 string = "hello"
                _ = v48
            }
            _ = fi31
        }
    case 42:
        v4 = (v26 <= 2)
    case 9223372036854775807:
        var v49 float64 = 3.14159
        _ = v49
    default:
        if (2147483647 <= 32768) {
            switch "\n\t" {
            case "\n\t":
                v50 := GenS5{F0: v25, F1: v1, F2: (v1 % (((v1 + genFn0(genFn0(2, 32767, ("a" + "")), 1, "foo bar"))) | 1))}
                _ = v50
            default:
                var v51 uint = 1
                _ = v51
            }
        }
    }
    var v52 int64 = -9223372036854775808
    switch 732 {
    case 1:
        v4 = (v4 && v4)
    case -130:
        var v53 bool = (1.5 != v3)
        _ = v53
    case -2147483648:
        v2 = append(v2, "a")
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v52
}
