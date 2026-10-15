package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int; F2 uint8 }
type GenS3 struct { F0 int64; F1 bool; F2 bool }

func main() {
    var v1 int = 32768
    v2 := make(map[string]int)
    v3 := []int{v1, (v1 | (-(v1))), v1, (^(v1))}
    var v4 float32 = -2.25
    var v5 string = ("a" + ("\n\t" + "hello"))
    v5 = ("" + "")
    var v6 int32 = -376
    v7 := GenS1{F0: -32768}
    v8 := GenS1{F0: 1}
    v5 = v5
    v5 = v5
    if (!(-32768 <= 42)) {
        v3 = append(v3, v1)
        v4 += 0.0
    } else {
        if false {
            v9 := GenS2{F0: -18, F1: v1, F2: 202}
            for fi10 := 0; fi10 < 17; fi10++ {
                if true {
                    var v11 uint = 930
                    switch (("a" + (v5 + v5)) + v5) {
                    case "a":
                        v12 := GenS3{F0: 0, F1: (float64(-2.25) <= 0.1), F2: (1 > 233)}
                        _ = v12
                    }
                    _ = v11
                } else {
                    v6 = (-(v6))
                    var v13 float32 = (((v4 * v4) / v4) / 1e10)
                    var v14 bool = false
                    _ = v13
                    _ = v14
                }
                v6 -= v6
                v5 = v5
                _ = fi10
            }
            _ = v9
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 int8 = -128
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v6 -= (v6 - -132)
            _ = fi15
            _ = v16
        }
        var v17 bool = false
        _ = v17
    }
    v5 = v5
    var v18 int32 = 42
    switch v1 {
    case 42:
        if ((v4 * 1e10) != 0.0) {
            fmt.Println(v1, v1)
            switch 127 {
            case 32768:
                if (!false) {
                    v6 = v18
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            default:
                var v19 float64 = ((0.1 + 1e10) + ((1e10 / (float64(v4) * ((float64(1.5) * (0.1 + 1.5)) * ((float64((1e10 + v4)) - (((1e10 * (1e10 - 1e10)) - float64(((((v4 * v4) / (float32(float64(v4)) - (float32((0.1 * (-2.25 + float64(((float32(((1e10 - 0.1) / 1.0)) * v4) + float32(3.14159)))))) * v4))) * (0.0 / 1.5)) + 0.0))) * 1.5)) / (1.5 / 1.0))))) - float64(v4)))
                _ = v19
            }
        }
    default:
        fmt.Println(v1)
    }
    if true {
        _ = len(v2)
        var v20 string = v5
        if (-2.25 == (float32((3.14159 + float64(float32(-2.25)))) * 0.1)) {
            v18 += v18
            var v21 uint64 = 635
            if true {
                for fi22 := 0; fi22 < 16; fi22++ {
                    delete(v2, "")
                    switch v20 {
                    case "hello":
                        for fi23 := 0; fi23 < 14; fi23++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi23
                        }
                    case "a":
                        switch v5 {
                        case "a":
                            fmt.Println(v18, v21)
                        case "":
                            var v24 int16 = 42
                            _ = v24
                        case "hello":
                            v1 = (v1 * fi22)
                        default:
                            var v25 string = ("" + v5)
                            _ = v25
                        }
                    case "":
                        var v26 uint32 = 0
                        _ = v26
                    default:
                        if len(v3) > 0 {
                            v3[0] = fi22
                        }
                    }
                    var v27 float32 = 1e10
                    _ = fi22
                    _ = v27
                }
                v20 = (v5 + v20)
                v21 -= (^(v21))
            } else {
                v18 = v18
                var v28 int16 = 32767
                _ = v28
            }
            _ = v21
        } else {
            var v29 bool = false
            _ = v29
        }
        _ = v20
    }
    v1 += v1
    v6 = (^(v6))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
