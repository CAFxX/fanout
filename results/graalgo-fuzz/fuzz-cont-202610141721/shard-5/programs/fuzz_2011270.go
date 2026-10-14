package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int16 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int; F2 int16 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int64 = -296
    v2 := make(map[int]bool)
    var v3 uint32 = 92
    var v4 string = "a"
    var v5 int16 = -765
    var v6 uint = 4294967295
    fmt.Println(v5, v3)
    for fi7 := 0; fi7 < 17; fi7++ {
        v3 = 498
        _ = fi7
    }
    if (v5 >= (v5 << ((18446744073709551615) & 63))) {
        fmt.Println(v3)
        delete(v2, 4294967295)
        v8 := GenS1{F0: 100, F1: true, F2: (v5 & (v5 << ((v6) & 63)))}
        _ = v8
    } else {
        var v9 int32 = 0
        if (!false) {
            fmt.Println(v6, v1, v9)
            switch -9223372036854775808 {
            case -1:
                v6 += v6
            default:
                fmt.Println(v4, v6, v4)
            }
        } else {
            v10 := GenS2{F0: 127, F1: 255}
            var v11 bool = true
            var v12 int64 = v1
            _ = v10
            _ = v11
            _ = v12
        }
        _ = v9
    }
    v13 := GenS3{F0: 32767, F1: -9223372036854775808, F2: (v5 << ((368) & 63))}
    v14 := []int{9223372036854775807, -1, -1, 42}
    if true {
        var v15 string = v4
        v2[42] = false
        _ = v15
    }
    v16 := GenS1{F0: 1, F1: true, F2: (v5 - (v5 - v5))}
    var v17 bool = (!true)
    var v18 uint64 = 455
    if (((-2.25 - 1.5) * 3.14159) == float32(float64(0.0))) {
        fmt.Println(v3, v1, v17)
        if v17 {
            if (v17 && (v17 || true)) {
                var v19 string = v4
                var v20 int16 = (-(v5))
                _ = v19
                _ = v20
            }
            switch "hello" {
            case "a":
                v14 = append(v14, 42)
            case "\n\t":
                if true {
                    fmt.Println(v5, v5)
                    if (v17 || (!true)) {
                        fmt.Println(v17)
                        switch "\n\t" {
                        case "\n\t":
                            var v21 int8 = 86
                            _ = v21
                        case "a":
                            v3 += 0
                        case "foo bar":
                            for fi22 := 0; fi22 < 1; fi22++ {
                                v23 := v2[127]
                                v24 := GenS4{F0: 42}
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        default:
                            for fi25 := 0; fi25 < 14; fi25++ {
                                if (v17 || v17) {
                                    if (fi25 == fi25) {
                                        for fi26 := 0; fi26 < 12; fi26++ {
                                            v4 = (v4 + (v4 + (v4 + "")))
                                            _ = fi26
                                        }
                                        v27 := v2[42]
                                        if len(v14) > 0 {
                                            v14[0] = fi25
                                        }
                                        _ = v27
                                    } else {
                                        v28 := GenS3{F0: 1, F1: fi25, F2: int16(v1)}
                                        var v29 uint8 = 44
                                        _ = v28
                                        _ = v29
                                    }
                                    var v30 uint32 = (v3 % ((v3) | 1))
                                    v1 += (v1 % ((v1) | 1))
                                    _ = v30
                                } else {
                                    fmt.Println(v3)
                                    if v17 {
                                        fmt.Println(v4)
                                        for fi31 := 0; fi31 < 19; fi31++ {
                                            v32 := GenS1{F0: 94, F1: v17, F2: v5}
                                            var v33 string = "hello"
                                            _ = fi31
                                            _ = v32
                                            _ = v33
                                        }
                                        for fi34 := 0; fi34 < 5; fi34++ {
                                            fmt.Println(fi34, fi34, v18)
                                            _ = fi34
                                        }
                                    } else {
                                        fmt.Println(v18)
                                        if (0.1 != 1.5) {
                                            v5 -= (v5 & (v5 * (v5 + v5)))
                                        }
                                    }
                                }
                                fmt.Println(v18)
                                fmt.Println(v1, v17, v18)
                                _ = fi25
                            }
                        }
                    }
                    v35 := GenS1{F0: 42, F1: (0.0 > float32((-2.25 / 1.5))), F2: v5}
                    _ = v35
                } else {
                    delete(v2, 1)
                    v17 = ((v1 / ((v1) | 1)) != (v1 | v1))
                }
            }
            switch v4 {
            case "\n\t":
                var v36 int8 = 100
                _ = v36
            }
        }
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        v38 := GenS2{F0: -1, F1: uint16(v18)}
        _ = fi37
        _ = v38
    }
    var v39 uint8 = 128
    var v40 int64 = v1
    v1 -= v1
    v40 -= 9223372036854775807
    v14 = append(v14, -2147483648)
    for fi41 := 0; fi41 < 7; fi41++ {
        v2[fi41] = ((false && (!(0 > int16(v6)))) || (-2.25 <= 0.1))
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
