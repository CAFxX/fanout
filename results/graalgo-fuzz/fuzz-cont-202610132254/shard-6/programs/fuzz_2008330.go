package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 bool; F1 int8; F2 string }
type GenS3 struct { F0 int16; F1 uint8; F2 string }
type GenS4 struct { F0 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{2147483647, 810}
    v3 := []uint32{4294967295, 2147483648, 4294967295, 853}
    var v4 uint = 45
    var v5 int32 = 0
    _ = len(v3)
    v4 -= v4
    var v6 bool = (0 < 42)
    var v7 int8 = 92
    v8 := GenS1{F0: 2, F1: -1}
    v7 = v7
    var v9 int8 = (v7 >> (((v4 % (((v4 - v4)) | 1))) & 63))
    v10 := GenS1{F0: 1, F1: 951}
    for fi11 := 0; fi11 < 16; fi11++ {
        if false {
            v12 := GenS2{F0: true, F1: -22, F2: ""}
            var v13 float64 = ((3.14159 + 1e10) / ((((1.5 + 0.0) - float64((float32((float64(1.5) - (-2.25 - 0.0))) + float32(3.14159)))) * ((float64(1.5) * ((1e10 - 0.1) + ((1.5 / 1.0) * float64(-2.25)))) * 1.5)) / -2.25))
            _ = v12
            _ = v13
        }
        _ = fi11
    }
    var v14 int = 1
    var v15 uint = 2
    if v6 {
        var v16 string = ("" + "\n\t")
        _ = v16
    }
    if (!v6) {
        var v17 string = ("" + "a")
        if (v6 || (!false)) {
            var v18 int64 = 42
            _ = v18
        } else {
            v7 = v7
            switch (v14 & (^(v14))) {
            case 818:
                if len(v3) > 0 {
                    v3[0] = 1
                }
            case 2147483647:
                v5 = (v5 * (^(v5)))
            default:
                v7 += v7
            }
            if (v14 >= 32768) {
                var v19 uint16 = 65535
                var v20 uint32 = 2147483648
                _ = v19
                _ = v20
            } else {
                v9 = v7
                if false {
                    switch v14 {
                    case -1:
                        for fi21 := 0; fi21 < 14; fi21++ {
                            fmt.Println(v15)
                            for fi22 := 0; fi22 < 10; fi22++ {
                                if len(v3) > 0 {
                                    v3[0] = 1
                                }
                                if v6 {
                                    if (v6 || (1 == v5)) {
                                        var v23 uint8 = 8
                                        _ = v23
                                    }
                                    var v24 bool = v6
                                    v25 := GenS3{F0: 0, F1: 42, F2: "\n\t"}
                                    fmt.Println(v15)
                                    fmt.Println(v6, v5, v9)
                                    _ = v24
                                    _ = v25
                                }
                                if (18446744073709551615 <= 1) {
                                    v17 = v17
                                    switch fi21 {
                                    case -2147483648:
                                        var v26 int = (fi21 / ((42) | 1))
                                        _ = v26
                                    default:
                                        if len(v3) > 0 {
                                            v3[0] = 42
                                        }
                                    }
                                    var v27 uint16 = 2
                                    _ = v27
                                }
                                _ = fi22
                            }
                            var v28 float32 = (float32(-2.25) - (1e10 * (1e10 * 0.1)))
                            _ = fi21
                            _ = v28
                        }
                    case 793:
                        v29 := GenS4{F0: v17}
                        _ = v29
                    default:
                        v30 := GenS1{F0: uint64(v15), F1: v14}
                        _ = v30
                    }
                    v7 = 1
                }
            }
        }
        _ = v17
    }
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
