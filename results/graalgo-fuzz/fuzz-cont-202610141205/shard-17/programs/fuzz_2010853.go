package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int32 }

func main() {
    var v1 string = ("a" + ("foo bar" + "\n\t"))
    v2 := []uint32{2}
    v3 := make(map[uint32]int)
    var v4 uint64 = 42
    v5 := make(map[string]bool)
    v6 := make(map[string]uint32)
    fmt.Println(v1, v4)
    v1 = v1
    if (((((1.5 - (3.14159 - 1e10)) / float32(-2.25)) * 1.5) * 3.14159) < (1e10 + 0.1)) {
        switch 0 {
        case 32767:
            for fi7 := 0; fi7 < 6; fi7++ {
                if false {
                    if (-32768 >= 738) {
                        var v8 uint32 = 266
                        for fi9 := 0; fi9 < 2; fi9++ {
                            var v10 float64 = float64(-2.25)
                            var v11 float32 = ((0.0 - 0.1) / ((float32(0.0) / 3.14159) - 0.1))
                            if (290 == fi9) {
                                v4 += v4
                                var v12 uint = 18446744073709551615
                                _ = v12
                            }
                            switch (v1 + v1) {
                            case "\n\t":
                                if ((0 == 1) || (!true)) {
                                    var v13 uint16 = 255
                                    var v14 int32 = 663
                                    _ = v13
                                    _ = v14
                                } else {
                                    v4 = v4
                                }
                            case "foo bar":
                                v4 = 42
                            default:
                                v15 := GenS1{F0: -79}
                                _ = v15
                            }
                            v16 := GenS1{F0: int8(v8)}
                            _ = fi9
                            _ = v10
                            _ = v11
                            _ = v16
                        }
                        for fi17 := 0; fi17 < 2; fi17++ {
                            var v18 uint = 920
                            v19 := GenS1{F0: -128}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        _ = v8
                    }
                    var v20 int64 = 42
                    v2 = append(v2, 1)
                    _ = v20
                }
                var v21 uint = 4294967295
                _ = fi7
                _ = v21
            }
        case 32768:
            v22 := GenS2{F0: v1}
            _ = v22
        }
    }
    var v23 uint8 = 122
    fmt.Println(v4)
    var v24 int32 = 42
    if ((1.5 / 1e10) == ((1e10 / 3.14159) / 1.0)) {
        switch 128 {
        case 4294967295:
            var v25 float64 = (((0.0 * 1e10) + 0.1) + (1.5 + float64(1.5)))
            _ = v25
        case 128:
            var v26 float32 = 0.0
            _ = v26
        case 1:
            v27 := GenS3{F0: 32767, F1: 955}
            _ = v27
        }
        v4 -= v4
    }
    v28 := GenS3{F0: 231, F1: (v24 / ((1) | 1))}
    v23 = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
