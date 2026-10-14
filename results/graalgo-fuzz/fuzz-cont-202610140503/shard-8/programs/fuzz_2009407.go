package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 uint32 = 240
    var v2 uint64 = 235
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 uint = 2
        var v7 uint16 = 919
        if false {
            if (!false) {
                var v8 float64 = (((-2.25 * 0.0) + 3.14159) + ((1.5 + 0.1) + (float64(0.1) / (((0.0 / (1e10 + 0.0)) * (((3.14159 - 0.1) * ((0.1 * (1.5 - 0.1)) * ((-2.25 / 1e10) * 1e10))) * 1.5)) + -2.25))))
                var v9 uint32 = (^(v1))
                _ = v8
                _ = v9
            } else {
                switch "\n\t" {
                case "hello":
                    fmt.Println(v6)
                case "":
                    switch "" {
                    case "foo bar":
                        var v10 int32 = -566
                        _ = v10
                    case "hello":
                        var v11 float32 = 1.5
                        _ = v11
                    case "a":
                        if (false && false) {
                            v12 := GenS1{F0: 255, F1: v7, F2: (v6 / ((255) | 1))}
                            switch (^(fi5)) {
                            case -786:
                                v13 := []int{}
                                _ = v13
                            case 128:
                                v6 = v6
                            }
                            _ = v12
                        }
                    }
                case "\n\t":
                    for fi14 := 0; fi14 < 3; fi14++ {
                        var v15 int64 = 1
                        fmt.Println(v7, v7, v15)
                        _ = fi14
                        _ = v15
                    }
                default:
                    v16 := GenS2{F0: 1, F1: v1, F2: 42}
                    _ = v16
                }
                var v17 float32 = 0.0
                var v18 float32 = 0.0
                _ = v17
                _ = v18
            }
            var v19 int16 = int16(v6)
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 bool = (2 != 1)
                _ = fi20
                _ = v21
            }
            _ = v19
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    if (false && false) {
        v22 := GenS2{F0: 18446744073709551615, F1: v1, F2: 776}
        var v23 int64 = 1
        v24 := []int{521, -369, -461, 1}
        _ = v22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        fmt.Println(v2, fi25, v1)
        for fi26 := 0; fi26 < 18; fi26++ {
            var v27 int8 = 1
            v28 := GenS1{F0: 42, F1: 507, F2: 0}
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        switch fi29 {
        case 4294967295:
            v1 = (^(v1))
        case 127:
            for fi30 := 0; fi30 < 9; fi30++ {
                switch (^(fi29)) {
                case -2147483648:
                    v1 = (v1 | (v1 ^ 42))
                case 2147483647:
                    v31 := []int{}
                    v1 += 0
                    _ = v31
                case 32768:
                    if true {
                        switch fi30 {
                        case 128:
                            v32 := GenS2{F0: v2, F1: v1, F2: 318}
                            _ = v32
                        }
                        v1 += uint32(v2)
                        if (3.14159 <= (-2.25 * ((-2.25 * (0.0 / 1.0)) - 3.14159))) {
                            fmt.Println(fi30, fi30)
                            v33 := []int{(^(fi29)), (fi30 >> ((855) & 63)), fi29}
                            v1 = v1
                            _ = v33
                        } else {
                            v1 -= (v1 % (((v1 * v1)) | 1))
                        }
                        var v34 uint16 = 487
                        _ = v34
                    } else {
                        if false {
                            v35 := GenS3{F0: 42, F1: v2}
                            _ = v35
                        }
                    }
                }
                v36 := v3[fi29]
                _ = fi30
                _ = v36
            }
        case 128:
            v1 = (v1 / ((0) | 1))
        default:
            if false {
                var v37 uint = 177
                _ = v37
            }
        }
        fmt.Println(fi29)
        _ = fi29
    }
    v38 := GenS4{F0: 127}
    v2 -= (v2 * v2)
    var v39 bool = (-2.25 >= 0.1)
    var v40 uint8 = 128
    v2 = v2
    var v41 bool = (v40 != v40)
    var v42 int = 544
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
