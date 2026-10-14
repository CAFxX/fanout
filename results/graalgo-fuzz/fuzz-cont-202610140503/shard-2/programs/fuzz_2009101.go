package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 uint }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 int64 = 0
    var v4 int32 = 2147483647
    v5 := make(map[int]bool)
    v6 := []int{}
    var v7 bool = (-2.25 > -2.25)
    v8 := GenS1{F0: (("hello" + ("\n\t" + "hello")) + "a")}
    var v9 uint64 = 42
    v10 := GenS2{F0: 76, F1: 0}
    switch 32767 {
    case 32767:
        v1[326] = 4294967295
    case 32768:
        v9 += (v9 & 0)
    default:
        if v7 {
            var v11 int8 = 127
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        switch 731 {
        case -184:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case 128:
            fmt.Println(v3)
        case 4294967295:
            for fi13 := 0; fi13 < 4; fi13++ {
                if false {
                    if (v7 || false) {
                        var v14 int = 590
                        if (false && v7) {
                            v15 := GenS3{F0: 23, F1: (v3 / ((v3) | 1))}
                            v6 = append(v6, fi13)
                            switch 32767 {
                            case -1:
                                v16 := GenS2{F0: 0, F1: 731}
                                _ = v16
                            case 0:
                                v3 = 0
                            default:
                                v9 += 4294967295
                            }
                            _ = v15
                        }
                        v3 -= 735
                        _ = v14
                    } else {
                        var v17 uint16 = 65535
                        switch "\n\t" {
                        case "\n\t":
                            if ((-2.25 > 0.0) && v7) {
                                v18 := GenS4{F0: (v3 >> ((0) & 63)), F1: -9223372036854775808}
                                fmt.Println(fi12, fi12, fi12)
                                v19 := GenS4{F0: v3, F1: v3}
                                _ = v18
                                _ = v19
                            }
                        case "hello":
                            _ = len(v1)
                        case "":
                            var v20 float64 = (float64(0.0) - (-2.25 / 0.1))
                            _ = v20
                        default:
                            var v21 uint8 = 42
                            _ = v21
                        }
                        var v22 float32 = 3.14159
                        delete(v1, 2147483647)
                        _ = v17
                        _ = v22
                    }
                    fmt.Println(fi12, v9)
                    var v23 int = fi13
                    v24 := GenS3{F0: uint8(v9), F1: (v3 << ((uint(v9)) & 63))}
                    _ = v23
                    _ = v24
                }
                _ = fi13
            }
        }
        _ = fi12
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        v26 := GenS5{F0: (v3 * (v3 + (v3 + (v3 ^ 1)))), F1: (fi25 + (-(fi25)))}
        switch (fi25 & fi25) {
        case 32768:
            v27 := GenS6{F0: v4}
            _ = v27
        case -311:
            v4 = int32(v3)
        case 2147483647:
            var v28 uint32 = 307
            _ = v28
        default:
            for fi29 := 0; fi29 < 10; fi29++ {
                var v30 uint16 = 2
                _ = len(v5)
                v6 = append(v6, int(v3))
                _ = fi29
                _ = v30
            }
        }
        _ = fi25
        _ = v26
    }
    var v31 uint64 = (v9 % (((v9 / ((733) | 1))) | 1))
    switch 42 {
    case 260:
        v9 = 0
    case 346:
        if (v4 != (v4 / ((v4) | 1))) {
            v31 -= (v31 & 59)
            if v7 {
                switch -2147483648 {
                case 0:
                    _ = len(v2)
                default:
                    fmt.Println(v9)
                }
                v4 += v4
                if v7 {
                    v6 = append(v6, 41)
                }
            } else {
                var v32 int8 = -4
                var v33 int = -392
                v4 += v4
                _ = v32
                _ = v33
            }
        }
    case 482:
        var v34 uint32 = 1
        _ = v34
    }
    if (1e10 > 0.0) {
        var v35 uint = 858
        v4 += (v4 & (^(v4)))
        _ = v35
    }
    if (!v7) {
        v4 -= (^(v4))
        _ = len(v6)
        var v36 int64 = v3
        _ = v36
    } else {
        v31 += v9
        for fi37 := 0; fi37 < 6; fi37++ {
            v38 := GenS7{F0: (fi37 ^ (fi37 & (fi37 + -802))), F1: uint(v9)}
            v9 += 4294967295
            v9 = v31
            _ = fi37
            _ = v38
        }
        v6 = append(v6, 4294967295)
    }
    if v7 {
        v3 = 42
        v31 = (v9 >> ((4294967295) & 63))
        var v39 float32 = 1.5
        _ = v39
    }
    v2["hello"] = "hello"
    for fi40 := 0; fi40 < 1; fi40++ {
        v1[(fi40 % (((fi40 - (fi40 | int(v31)))) | 1))] = 65535
        _ = fi40
    }
    fmt.Println(v9)
    var v41 int32 = (v4 | (v4 / (((v4 >> ((65535) & 63))) | 1)))
    var v42 float32 = 3.14159
    v9 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
