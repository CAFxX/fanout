package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int16; F1 int; F2 int32 }

func main() {
    var v1 uint = 160
    v2 := []uint32{}
    var v3 int = -792
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    fmt.Println(v1)
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 string = "\n\t"
        _ = fi6
        _ = v7
    }
    fmt.Println(v3, v1)
    if (false && true) {
        var v8 uint8 = 2
        if false {
            var v9 uint8 = (v8 - v8)
            v10 := GenS1{F0: 378}
            _ = v10
            _ = v9
        } else {
            v1 += (v1 % (((v1 & 2)) | 1))
            var v11 uint16 = 65535
            v12 := GenS2{F0: 148, F1: (v8 ^ v8), F2: 2147483647}
            _ = v11
            _ = v12
        }
        v13 := GenS2{F0: 180, F1: v8, F2: 2147483647}
        _ = v13
        _ = v8
    } else {
        switch (-(v3)) {
        case 928:
            _ = len(v5)
        }
        v1 -= 2
    }
    _ = len(v4)
    v1 = v1
    var v14 int = (v3 / ((v3) | 1))
    if (!(0.0 < ((0.0 - float64(-2.25)) * 0.0))) {
        v15 := GenS2{F0: 158, F1: 2, F2: 1}
        var v16 int8 = 53
        v14 = v14
        _ = v15
        _ = v16
    } else {
        var v17 uint64 = 42
        var v18 uint32 = 2
        v3 = (v3 ^ (-(v3)))
        _ = v17
        _ = v18
    }
    delete(v5, v14)
    if (!true) {
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
        if (1.5 != 1.5) {
            _ = len(v5)
            var v19 int32 = 2147483647
            _ = v19
        } else {
            fmt.Println(v14, v1)
            v3 = v3
        }
        if (!false) {
            for fi20 := 0; fi20 < 18; fi20++ {
                var v21 uint64 = 4294967295
                if true {
                    for fi22 := 0; fi22 < 6; fi22++ {
                        v21 = (v21 / ((v21) | 1))
                        v23 := GenS1{F0: 985}
                        _ = fi22
                        _ = v23
                    }
                    fmt.Println(v14)
                    var v24 int16 = int16(v21)
                    _ = v24
                } else {
                    fmt.Println(v3, v21)
                    v3 = v3
                    v1 -= v1
                }
                if false {
                    fmt.Println(v14)
                    switch "" {
                    case "a":
                        fmt.Println(v21)
                    }
                    v21 -= v21
                }
                _ = fi20
                _ = v21
            }
            var v25 uint = (v1 % ((v1) | 1))
            v3 = (v14 / (((v14 >> ((831) & 63))) | 1))
            _ = v25
        }
    } else {
        v1 = (^(v1))
        fmt.Println(v3)
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        var v27 uint16 = 225
        fmt.Println(v1)
        if (false && false) {
            for fi28 := 0; fi28 < 4; fi28++ {
                v1 -= (v1 | (v1 / ((2) | 1)))
                var v29 bool = true
                if (0 <= 229) {
                    v14 -= (fi26 >> (((v1 << ((0) & 63))) & 63))
                    if true {
                        var v30 uint64 = 134
                        if (!(v29 || false)) {
                            switch "foo bar" {
                            case "\n\t":
                                switch v14 {
                                case 79:
                                    fmt.Println(fi28, v29)
                                    fmt.Println(fi28)
                                    if v29 {
                                        v29 = (42 < 1)
                                        v14 = (v14 << (((v1 / ((uint(v27)) | 1))) & 63))
                                        v2 = append(v2, 1)
                                    }
                                case 905:
                                    v4[("hello" + "hello")] = (v14 % (((-(v3))) | 1))
                                }
                            case "a":
                                v31 := GenS3{F0: 42, F1: (fi26 * fi26), F2: 0}
                                _ = v31
                            }
                            var v32 uint = (^(v1))
                            switch (fi26 & (fi28 ^ fi28)) {
                            case 32767:
                                var v33 uint8 = 255
                                _ = v33
                            case 32768:
                                var v34 bool = (true || v29)
                                _ = v34
                            default:
                                v32 = uint(v30)
                            }
                            var v35 int8 = 1
                            _ = v32
                            _ = v35
                        }
                        fmt.Println(fi26)
                        v14 = v14
                        _ = v30
                    } else {
                        if len(v2) > 0 {
                            v2[0] = 0
                        }
                        var v36 float32 = ((0.0 * -2.25) + 1e10)
                        _ = v36
                    }
                }
                var v37 bool = v29
                _ = fi28
                _ = v29
                _ = v37
            }
        }
        _ = fi26
        _ = v27
    }
    v38 := GenS2{F0: 76, F1: 2, F2: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
