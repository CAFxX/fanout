package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int64 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 int16; F1 uint8; F2 int64 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int32 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 string = "hello"
    var v3 string = v2
    var v4 float64 = 0.0
    var v5 int32 = -1
    switch v2 {
    case "":
        var v6 uint32 = 2147483647
        _ = v6
    default:
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: "a", F1: (v5 % ((42) | 1)), F2: v1}
            v9 := GenS1{F0: v3, F1: (v5 / (((v5 >> ((596) & 63))) | 1)), F2: v1}
            fmt.Println(v5)
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    var v10 float64 = ((v4 - 1e10) / v4)
    fmt.Println(v2)
    switch -2147483648 {
    case 2147483647:
        v2 = v2
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v10 = v4
        _ = fi11
    }
    if false {
        v12 := GenS2{F0: 42, F1: ""}
        _ = v12
    }
    if (1e10 >= 1.5) {
        for fi13 := 0; fi13 < 20; fi13++ {
            v14 := []int{fi13}
            _ = fi13
            _ = v14
        }
    }
    if (9223372036854775807 >= -1) {
        v3 = v3
        if true {
            var v15 float64 = ((v10 - 3.14159) * (float64(float32(1.5)) / (v4 + (float64(3.14159) - v4))))
            _ = v15
        } else {
            fmt.Println(v1, v10, v4)
        }
        v16 := []int{}
        _ = v16
    } else {
        v10 = 0.1
    }
    var v17 int = 32767
    for fi18 := 0; fi18 < 8; fi18++ {
        if true {
            v19 := []int{v17, -743, (v17 % ((-304) | 1)), v17}
            if false {
                v5 = -616
                v4 = ((3.14159 / 1e10) - (0.0 - (v4 * 0.0)))
            }
            if true {
                v20 := GenS2{F0: -32768, F1: v2}
                if (false || true) {
                    v5 = v5
                    v5 -= 42
                    var v21 int16 = 1
                    _ = v21
                }
                fmt.Println(v10, v5)
                _ = v20
            }
            _ = v19
        } else {
            v17 = fi18
            v22 := []int{v17}
            for fi23 := 0; fi23 < 13; fi23++ {
                if (((238 == (v5 + (v5 & (v5 ^ v5)))) || (-231 == 32767)) || false) {
                    v3 = v2
                } else {
                    v17 = (v17 >> ((221) & 63))
                    _ = len(v22)
                    if (v5 != 1) {
                        switch "\n\t" {
                        case "a":
                            v1 -= v1
                        default:
                            _ = len(v22)
                            _ = len(v22)
                            fmt.Println(v10, v5)
                            v1 = v1
                            v17 -= (^(fi18))
                        }
                        if false {
                            var v24 float64 = (1.5 * (1.5 * v10))
                            var v25 int = 1
                            if true {
                                v4 -= (((v4 / v4) - 3.14159) / 1.0)
                                if (2 < 453) {
                                    v25 += v17
                                } else {
                                    fmt.Println(v5)
                                    if len(v22) > 0 {
                                        v22[0] = v17
                                    }
                                    var v26 int32 = (v5 >> ((2) & 63))
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                    _ = v26
                                }
                                v4 -= float64(1e10)
                            } else {
                                if (false || false) {
                                    var v27 uint32 = 190
                                    var v28 bool = false
                                    fmt.Println(v3, fi23)
                                    v29 := GenS3{F0: -836, F1: 174}
                                    v30 := GenS4{F0: v25, F1: v5}
                                    switch (v25 + 42) {
                                    case 128:
                                        if v28 {
                                            v31 := GenS1{F0: "\n\t", F1: (v5 >> ((18446744073709551615) & 63)), F2: v1}
                                            var v32 int16 = 42
                                            fmt.Println(v5)
                                            _ = v31
                                            _ = v32
                                        }
                                    }
                                    if true {
                                        if len(v22) > 0 {
                                            v22[0] = fi18
                                        }
                                        if (!v28) {
                                            switch 9223372036854775807 {
                                            case 1:
                                                switch fi23 {
                                                case 9223372036854775807:
                                                    _ = len(v22)
                                                    switch 864 {
                                                    case 420:
                                                        v27 = 978
                                                    default:
                                                        v33 := GenS4{F0: v25, F1: v5}
                                                        _ = v33
                                                    }
                                                case 0:
                                                    v3 = v2
                                                case 1:
                                                    v5 -= v5
                                                default:
                                                    fmt.Println(v25)
                                                }
                                            default:
                                                v34 := GenS5{F0: 473, F1: 55, F2: v1}
                                                _ = v34
                                            }
                                            _ = len(v22)
                                            if (true || v28) {
                                                v35 := GenS5{F0: 32767, F1: 2, F2: v1}
                                                _ = v35
                                            }
                                        } else {
                                            if len(v22) >= 2 {
                                                v22 = v22[1:len(v22)]
                                            }
                                        }
                                        switch "foo bar" {
                                        case "hello":
                                            if (false || v28) {
                                                if true {
                                                    fmt.Println(v4, v25)
                                                } else {
                                                    if len(v22) >= 2 {
                                                        v22 = v22[1:len(v22)]
                                                    }
                                                    var v36 uint32 = v27
                                                    v37 := GenS6{F0: 235, F1: 340}
                                                    _ = v36
                                                    _ = v37
                                                }
                                                v38 := GenS6{F0: 5, F1: -304}
                                                _ = v38
                                            }
                                        case "":
                                            switch (((v2 + (v3 + (v2 + "hello"))) + ((v2 + v2) + ((v2 + (v2 + v3)) + ("hello" + (v2 + (((((v2 + (v2 + v2)) + "a") + ((v2 + "\n\t") + "\n\t")) + v2) + (v2 + v2))))))) + v3) {
                                            case "foo bar":
                                                _ = len(v22)
                                            case "hello":
                                                v10 += (float64((0.0 * (((-2.25 * float32(v4)) - (((0.1 + 1.5) * (0.1 / 1.5)) - (1.5 - (-2.25 * -2.25)))) - ((float32(-2.25) + 0.1) / (0.1 + (0.1 * ((-2.25 / 1e10) * 3.14159))))))) / v24)
                                            case "":
                                                v17 -= v25
                                            default:
                                                switch fi23 {
                                                case 266:
                                                    var v39 uint8 = 1
                                                    _ = v39
                                                case 32767:
                                                    v1 -= v1
                                                    v40 := GenS2{F0: 0, F1: ("foo bar" + ((v2 + (((v2 + "a") + (("a" + "a") + v2)) + v3)) + v3))}
                                                    _ = v40
                                                case -431:
                                                    var v41 int16 = 262
                                                    _ = v41
                                                }
                                            }
                                        default:
                                            switch fi23 {
                                            case 674:
                                                if v28 {
                                                    v42 := GenS7{F0: int32(v17), F1: 0, F2: (^(v27))}
                                                    _ = v42
                                                } else {
                                                    v4 += float64(1.5)
                                                    _ = len(v22)
                                                    v22 = append(v22, 0)
                                                }
                                            default:
                                                v2 = "foo bar"
                                            }
                                        }
                                        switch v3 {
                                        case "a":
                                            switch v2 {
                                            case "hello":
                                                v43 := GenS6{F0: 128, F1: 292}
                                                var v44 uint32 = (v27 - 1)
                                                _ = v43
                                                _ = v44
                                            default:
                                                var v45 string = v2
                                                _ = v45
                                            }
                                        case "\n\t":
                                            if ((^(v5)) > 454) {
                                                if v28 {
                                                    switch v25 {
                                                    case -74:
                                                        v22 = append(v22, fi18)
                                                    case -907:
                                                        v4 -= v24
                                                        if ((765 < 42) || (v28 || false)) {
                                                            fmt.Println(v17)
                                                        } else {
                                                            if len(v22) > 0 {
                                                                v22[0] = v17
                                                            }
                                                            v22 = append(v22, (v25 - 1))
                                                            if v28 {
                                                                fmt.Println(v17, v1)
                                                                fmt.Println(v3, v17)
                                                                v46 := GenS8{F0: 32767}
                                                                _ = v46
                                                            }
                                                        }
                                                    default:
                                                        if (128 <= 2) {
                                                            v5 += (-(v5))
                                                        }
                                                    }
                                                } else {
                                                    if false {
                                                        fmt.Println(v27, v24, v25)
                                                        fmt.Println(fi18, v3)
                                                        fmt.Println(v4, v10, v24)
                                                    } else {
                                                        v2 = "hello"
                                                        var v47 bool = false
                                                        var v48 uint16 = uint16(v27)
                                                        var v49 int16 = -274
                                                        _ = v47
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                    var v50 int16 = 1000
                                                    _ = v50
                                                }
                                                if v28 {
                                                    v51 := GenS9{F0: (^(v5))}
                                                    _ = v51
                                                }
                                            } else {
                                                var v52 uint64 = 18446744073709551615
                                                switch v3 {
                                                case "hello":
                                                    fmt.Println(v52, v10)
                                                case "a":
                                                    v3 = ("hello" + "hello")
                                                case "\n\t":
                                                    var v53 uint16 = 684
                                                    _ = v53
                                                }
                                                var v54 float64 = v4
                                                _ = v52
                                                _ = v54
                                            }
                                        }
                                        v28 = ((!(fi18 < (fi18 >> ((297) & 63)))) && (v28 && true))
                                        var v55 float64 = v24
                                        var v56 string = (v3 + v3)
                                        _ = v55
                                        _ = v56
                                    }
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                                var v57 bool = (!(-9223372036854775808 == (v1 - (v1 * v1))))
                                _ = v57
                            }
                            _ = v24
                            _ = v25
                        }
                    } else {
                        fmt.Println(v10, fi18, v17)
                        v3 = v3
                    }
                    if (!(!((((0.0 + ((float32(v10) + 3.14159) / 1.0)) / (-2.25 / 1e10)) * 0.0) >= (1.5 - 1e10)))) {
                        v5 = v5
                        var v58 uint8 = 0
                        var v59 int64 = 9223372036854775807
                        _ = v58
                        _ = v59
                    } else {
                        v1 = v1
                    }
                }
                var v60 int64 = 1
                _ = fi23
                _ = v60
            }
            _ = v22
        }
        _ = fi18
    }
    v61 := GenS3{F0: 1, F1: 546}
    var v62 int32 = v5
    fmt.Println(v3, v17)
    v63 := GenS2{F0: 42, F1: ((v2 + "a") + v2)}
    v64 := []int{-1, -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    fmt.Println("slicelen v64:", len(v64), v64)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v62
    _ = v63
    _ = v64
}
