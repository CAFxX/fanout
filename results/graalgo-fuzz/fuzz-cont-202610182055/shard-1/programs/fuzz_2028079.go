package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
func genFn0() uint16 {
    if true {
        v7 := []int{2147483647, -9223372036854775808, 128}
        _ = v7
    } else {
        var v8 int16 = 1000
        if false {
            v9 := GenS1{F0: 0, F1: 255, F2: 44}
            v8 -= (v8 / ((v8) | 1))
            v8 = v8
            _ = v9
        }
        _ = v8
    }
    var v10 uint8 = 246
    v11 := GenS1{F0: 603, F1: 194, F2: v10}
    v10 += (v10 - (v10 & v10))
    _ = v10
    _ = v11
    return 42
}
func genFn1(p0 uint64) int64 {
    v12 := GenS1{F0: 1, F1: 65535, F2: 236}
    var v13 int32 = 42
    p0 -= (^(p0))
    _ = p0
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]int)
    var v3 float64 = (1.5 - ((0.0 - (3.14159 * (0.1 - 3.14159))) + 0.1))
    var v4 uint32 = 1
    var v5 uint = 0
    var v6 float64 = (v3 - v3)
    var v14 string = "foo bar"
    switch "" {
    case "a":
        v4 += 65535
    case "":
        if true {
            fmt.Println(v3, v1)
        } else {
            delete(v2, v4)
            fmt.Println(v4, v3, v1)
            if (737 >= 42) {
                v15 := v2[0]
                if true {
                    v4 += v4
                    v16 := []int{v15, (v15 >> ((uint(v4)) & 63)), (v15 * v15), (v15 % ((v15) | 1))}
                    _ = v16
                } else {
                    var v17 int = 874
                    v18 := []int{(v17 % (((-(v15))) | 1)), v17, (-(v17)), 42}
                    switch "" {
                    case "foo bar":
                        switch (v17 << (((v5 / ((v5) | 1))) & 63)) {
                        case 2147483647:
                            var v19 string = (v14 + (v14 + (((v14 + v14) + "foo bar") + v14)))
                            _ = v19
                        case 1:
                            if (0.0 < (0.1 * 0.1)) {
                                v3 = -2.25
                                fmt.Println(v5)
                                v20 := GenS2{F0: v1}
                                _ = v20
                            } else {
                                switch "foo bar" {
                                case "\n\t":
                                    fmt.Println(v3)
                                case "foo bar":
                                    v1 = v1
                                case "a":
                                    switch "foo bar" {
                                    case "hello":
                                        v21 := GenS1{F0: 1, F1: (v5 >> ((uint(v4)) & 63)), F2: uint8(v5)}
                                        _ = v21
                                    case "foo bar":
                                        switch 128 {
                                        case 127:
                                            fmt.Println(v15, v1)
                                        default:
                                            if len(v18) > 0 {
                                                v18[0] = 377
                                            }
                                        }
                                    }
                                default:
                                    if ((v5 + 154) > (v5 - 0)) {
                                        fmt.Println(v6, v3)
                                        var v22 uint8 = uint8(v4)
                                        v14 = (v14 + "foo bar")
                                        _ = v22
                                    }
                                }
                                if ((!false) || (-9223372036854775808 == 42)) {
                                    delete(v2, v4)
                                    v23 := GenS2{F0: 1000}
                                    if (false && false) {
                                        switch (v15 ^ (v15 * v17)) {
                                        case -1:
                                            switch v15 {
                                            case 2147483647:
                                                if true {
                                                    delete(v2, 515)
                                                }
                                            case 128:
                                                v24 := GenS2{F0: (v1 >> ((v5) & 63))}
                                                _ = v24
                                            case 0:
                                                v18 = append(v18, (v17 | -644))
                                            default:
                                                switch v14 {
                                                case "":
                                                    var v25 uint32 = (v4 >> ((255) & 63))
                                                    _ = v25
                                                case "hello":
                                                    v6 = ((v6 + (v3 * v3)) + v3)
                                                }
                                            }
                                        case 355:
                                            for fi26 := 0; fi26 < 3; fi26++ {
                                                if false {
                                                    v18 = append(v18, (v17 & v15))
                                                    if len(v18) > 0 {
                                                        v18[0] = 80
                                                    }
                                                }
                                                _ = fi26
                                            }
                                        }
                                        for fi27 := 0; fi27 < 17; fi27++ {
                                            v28 := GenS2{F0: (^(v1))}
                                            var v29 int16 = 1
                                            _ = fi27
                                            _ = v28
                                            _ = v29
                                        }
                                        v2[(^(v4))] = v15
                                    } else {
                                        var v30 int64 = 42
                                        v31 := GenS1{F0: -2147483648, F1: 2, F2: 2}
                                        _ = v30
                                        _ = v31
                                    }
                                    _ = v23
                                } else {
                                    v18 = append(v18, (v17 ^ v15))
                                }
                            }
                        default:
                            for fi32 := 0; fi32 < 5; fi32++ {
                                var v33 int64 = -9223372036854775808
                                v33 = v33
                                v14 = v14
                                _ = fi32
                                _ = v33
                            }
                        }
                    }
                    _ = v17
                    _ = v18
                }
                _ = v15
            }
        }
    default:
        v5 -= (v5 | 1)
    }
    var v34 uint8 = 128
    v4 = 4294967295
    v3 -= (v6 + float64(0.1))
    v6 = (v3 * float64(1.5))
    v35 := GenS2{F0: -1}
    switch 1 {
    case 32768:
        delete(v2, (v4 % (((v4 | (v4 >> ((v5) & 63)))) | 1)))
    case 1:
        v36 := []int{1, 42}
        _ = v36
    default:
        for fi37 := 0; fi37 < 17; fi37++ {
            v38 := GenS2{F0: 738}
            v4 += 844
            _ = fi37
            _ = v38
        }
    }
    var v39 uint8 = (^(v34))
    v40 := v2[v4]
    v41 := GenS3{F0: (v1 - v1)}
    if (v14 < ("a" + "")) {
        v1 -= (v1 % ((-695) | 1))
        v42 := GenS3{F0: v1}
        _ = v42
    }
    if (!true) {
        v43 := GenS1{F0: 1, F1: v5, F2: (v34 * (v34 * (v34 * (v34 << (((v5 | (^(v5)))) & 63)))))}
        _ = v43
    } else {
        v44 := GenS1{F0: -1, F1: uint(v4), F2: (v39 << (((v5 % ((v5) | 1))) & 63))}
        _ = v44
    }
    v45 := GenS3{F0: v1}
    v14 = v14
    for fi46 := 0; fi46 < 15; fi46++ {
        var v47 float32 = (float32(0.1) - (-2.25 / ((3.14159 + 0.1) - (1e10 + 0.0))))
        _ = fi46
        _ = v47
    }
    if (((-1 == 0) && (!false)) && (v40 < (v40 | v40))) {
        fmt.Println(v1, v1, v40)
    } else {
        fmt.Println(v14, v14, v40)
        v3 += (float64(0.1) + 3.14159)
        var v48 uint = 18446744073709551615
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v45
    _ = v5
    _ = v6
}
