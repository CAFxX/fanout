package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 int64; F2 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = (-2.25 * float32(-2.25))
    var v3 bool = false
    var v4 int = -693
    _ = len(v1)
    var v5 uint8 = 255
    if (2147483648 != 237) {
        for fi6 := 0; fi6 < 13; fi6++ {
            fmt.Println(v5, v4, v4)
            switch "" {
            case "a":
                if v3 {
                    var v7 string = "foo bar"
                    for fi8 := 0; fi8 < 19; fi8++ {
                        var v9 uint8 = v5
                        v3 = v3
                        switch 127 {
                        case 32767:
                            var v10 uint64 = 18446744073709551615
                            _ = v10
                        case 617:
                            fmt.Println(fi6, v2)
                        case 42:
                            fmt.Println(v3, v3)
                        default:
                            var v11 int32 = 42
                            _ = v11
                        }
                        _ = fi8
                        _ = v9
                    }
                    _ = v7
                } else {
                    v12 := GenS1{F0: (^(v5)), F1: v3, F2: 870}
                    v3 = (!(0 <= 0))
                    _ = v12
                }
            case "hello":
                if (5 > 25) {
                    if v3 {
                        if (100 == -46) {
                            for fi13 := 0; fi13 < 17; fi13++ {
                                var v14 int64 = -1
                                _ = fi13
                                _ = v14
                            }
                            v4 += v4
                            if (128 <= v5) {
                                for fi15 := 0; fi15 < 3; fi15++ {
                                    v4 += -663
                                    fmt.Println(fi6, v4)
                                    v16 := GenS2{F0: 65535}
                                    _ = fi15
                                    _ = v16
                                }
                            }
                        }
                        v17 := GenS3{F0: v3, F1: 0}
                        v18 := []int{(fi6 / (((v4 | v4)) | 1)), fi6}
                        _ = v17
                        _ = v18
                    }
                    var v19 int64 = 0
                    var v20 uint8 = v5
                    _ = v19
                    _ = v20
                } else {
                    v21 := GenS2{F0: 1}
                    _ = v21
                }
            case "":
                if false {
                    v22 := GenS4{F0: (v5 ^ v5)}
                    for fi23 := 0; fi23 < 14; fi23++ {
                        v3 = v3
                        v24 := []int{42, (fi6 & -812)}
                        var v25 int8 = -128
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = v22
                }
            }
            var v26 uint16 = 2
            _ = fi6
            _ = v26
        }
    } else {
        var v27 int32 = -509
        _ = v27
    }
    v28 := []int{v4, 0}
    v2 = v2
    fmt.Println(v3)
    var v29 int16 = 385
    var v30 int = v4
    switch "\n\t" {
    case "a":
        v2 = ((float32(((1e10 + 0.0) - float64(v2))) + v2) - (1e10 - float32((float64((v2 * ((v2 / (float32(1.5) / 0.1)) - 0.1))) + 1.5))))
    default:
        fmt.Println(v3, v29)
    }
    var v31 string = "\n\t"
    var v32 int = v30
    v30 = v32
    v33 := GenS2{F0: 845}
    v34 := GenS5{F0: 902, F1: 328, F2: 65535}
    v5 = (^(v5))
    switch (v4 >> ((18446744073709551615) & 63)) {
    case 2147483647:
        switch v4 {
        case -882:
            var v35 string = v31
            _ = v35
        case 9223372036854775807:
            v29 = (v29 >> ((0) & 63))
        case 42:
            var v36 string = v31
            _ = v36
        default:
            v5 = 0
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
