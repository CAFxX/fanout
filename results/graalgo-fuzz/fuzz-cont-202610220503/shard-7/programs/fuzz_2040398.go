package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int; F1 int }

func main() {
    v1 := make(map[int]string)
    var v2 bool = false
    v3 := make(map[string]string)
    var v4 string = "\n\t"
    v5 := GenS1{F0: 1, F1: 1000}
    var v6 int16 = 0
    var v7 uint16 = 2
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 float32 = 3.14159
        var v10 uint64 = 489
        switch (fi8 + fi8) {
        case -174:
            if (v2 || true) {
                var v11 float32 = 1.5
                switch v4 {
                case "\n\t":
                    var v12 int64 = int64(v10)
                    _ = v12
                case "hello":
                    v10 -= (v10 << ((65535) & 63))
                case "":
                    v13 := GenS1{F0: 984, F1: v6}
                    _ = v13
                default:
                    fmt.Println(fi8)
                }
                _ = v11
            }
        case 128:
            for fi14 := 0; fi14 < 6; fi14++ {
                v15 := GenS2{F0: 283, F1: 558}
                if (!(-2147483648 == -51)) {
                    if false {
                        if (uint(v7) <= 18446744073709551615) {
                            fmt.Println(v6)
                            if v2 {
                                var v16 int64 = 0
                                var v17 uint8 = uint8(v10)
                                _ = v16
                                _ = v17
                            } else {
                                switch -541 {
                                case 568:
                                    var v18 uint32 = 214
                                    _ = v18
                                case 4294967295:
                                    var v19 float32 = v9
                                    v1[fi8] = v4
                                    v6 = v6
                                    _ = v19
                                default:
                                    v10 = (v10 / ((4294967295) | 1))
                                }
                                fmt.Println(v4, v6, v10)
                                var v20 int = int(v10)
                                _ = v20
                            }
                            v21 := GenS1{F0: fi14, F1: int16(fi14)}
                            _ = v21
                        }
                        var v22 uint16 = v7
                        _ = v22
                    }
                    v7 = (v7 | v7)
                    var v23 uint = 42
                    switch -1 {
                    case 0:
                        var v24 float32 = (v9 - v9)
                        _ = v24
                    case 42:
                        v23 += v23
                    }
                    fmt.Println(fi14, v9)
                    var v25 float64 = -2.25
                    _ = v23
                    _ = v25
                } else {
                    switch (v4 + v4) {
                    case "":
                        var v26 string = v4
                        _ = v26
                    case "hello":
                        delete(v1, int(v6))
                    }
                }
                _ = fi14
                _ = v15
            }
        default:
            v27 := GenS1{F0: 0, F1: 0}
            _ = v27
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v28 uint16 = (^(v7))
    for fi29 := 0; fi29 < 16; fi29++ {
        v30 := []int{}
        v2 = (1 == 42)
        for fi31 := 0; fi31 < 18; fi31++ {
            v28 -= v7
            v32 := GenS3{F0: 42}
            v6 -= (v6 % (((-(v6))) | 1))
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    v7 = v28
    fmt.Println(v6)
    fmt.Println(v6)
    v2 = true
    var v33 int32 = -1
    v34 := GenS4{F0: 18446744073709551615, F1: -128}
    v35 := GenS5{F0: 4294967295, F1: 112}
    v2 = v2
    switch 0 {
    case 42:
        var v36 uint8 = 127
        _ = v36
    case 32767:
        v33 = v33
    default:
        var v37 uint16 = (v7 >> ((65535) & 63))
        _ = v37
    }
    v6 -= (v6 + (v6 / (((v6 & (-(v6)))) | 1)))
    v38 := []int{-1, 235, 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
