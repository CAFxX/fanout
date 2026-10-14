package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 1
    var v3 float32 = (((0.1 * (3.14159 * float32((0.1 + float64(float32((1.5 - 1.5))))))) - 1.5) * (float32(0.0) * 1.5))
    v4 := make(map[string]string)
    var v5 uint16 = 0
    v6 := []string{}
    v7 := GenS1{F0: -327, F1: -2147483648}
    if (42 <= v5) {
        switch "foo bar" {
        case "\n\t":
            v3 += 1.5
        case "":
            var v8 int32 = -1
            _ = v8
        }
    } else {
        var v9 int16 = 0
        _ = v9
    }
    fmt.Println(v3, v3)
    var v10 uint = 317
    var v11 uint32 = 756
    var v12 uint32 = v11
    if len(v6) > 0 {
        v6[0] = ((("hello" + "\n\t") + "") + "\n\t")
    }
    v13 := GenS1{F0: 32767, F1: -29}
    if (("a" + "hello") < "hello") {
        for fi14 := 0; fi14 < 6; fi14++ {
            for fi15 := 0; fi15 < 2; fi15++ {
                if false {
                    var v16 uint16 = v5
                    _ = v16
                }
                var v17 float32 = 0.1
                v5 += (^(v5))
                switch "\n\t" {
                case "":
                    v12 = 2147483648
                case "\n\t":
                    switch 1 {
                    case -1:
                        var v18 uint64 = uint64(v5)
                        _ = v18
                    case 127:
                        v6 = append(v6, "\n\t")
                    default:
                        v19 := GenS2{F0: 781, F1: 2, F2: 32767}
                        v20 := GenS1{F0: 1000, F1: 0}
                        v5 = (v5 | v5)
                        _ = v19
                        _ = v20
                    }
                    if true {
                        if ((1.5 + v3) != (0.1 * (0.0 - (0.1 - v17)))) {
                            v17 += v17
                        } else {
                            var v21 int32 = -306
                            v21 = 903
                            switch (fi15 ^ fi15) {
                            case 127:
                                v21 = v21
                            case 62:
                                _ = len(v6)
                            default:
                                delete(v4, (("" + "\n\t") + (("" + "hello") + "foo bar")))
                            }
                            v11 -= (v12 >> ((196) & 63))
                            v2 = int8(v5)
                            _ = v21
                        }
                        fmt.Println(v5, v12)
                    }
                    v3 -= (3.14159 + v3)
                    var v22 float32 = v17
                    _ = v22
                default:
                    fmt.Println(v5, v11)
                    switch (fi14 >> ((801) & 63)) {
                    case -820:
                        var v23 float32 = (0.1 / v3)
                        _ = v23
                    case 563:
                        v2 += v2
                    }
                    var v24 uint64 = uint64(v10)
                    _ = v24
                }
                var v25 uint32 = v11
                v26 := GenS2{F0: 42, F1: 781, F2: v5}
                var v27 uint32 = (v11 >> (((v10 | v10)) & 63))
                _ = fi15
                _ = v17
                _ = v25
                _ = v26
                _ = v27
            }
            v2 = (v2 - v2)
            _ = fi14
        }
        var v28 uint8 = 255
        v11 = 4294967295
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
