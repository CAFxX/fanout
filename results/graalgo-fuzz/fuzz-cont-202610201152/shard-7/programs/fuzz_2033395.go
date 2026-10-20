package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int64 }
type GenS2 struct { F0 string; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 9223372036854775807
    var v3 int = (v2 ^ v2)
    v4 := []int{(^(v2)), -292, 135}
    switch (("" + "\n\t") + "hello") {
    case "foo bar":
        var v5 int = (v3 ^ v3)
        _ = v5
    case "a":
        var v6 uint16 = 349
        _ = v6
    }
    v2 += v3
    fmt.Println(v2)
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS1{F0: -128, F1: 1, F2: 42}
        _ = fi7
        _ = v8
    }
    var v9 int32 = -496
    if len(v4) > 0 {
        v4[0] = 128
    }
    if (!false) {
        var v10 int64 = 42
        if false {
            var v11 bool = (((v3 % ((v3) | 1)) == (v3 * v3)) && false)
            var v12 float64 = 3.14159
            v1[1] = "a"
            _ = v11
            _ = v12
        } else {
            switch (v2 % (((v3 | -1)) | 1)) {
            case 0:
                switch "a" {
                case "a":
                    var v13 int8 = -79
                    _ = v13
                case "hello":
                    v2 -= v3
                default:
                    switch v3 {
                    case 974:
                        v9 -= v9
                    default:
                        if ((13 > 65535) || (18446744073709551615 == 42)) {
                            for fi14 := 0; fi14 < 16; fi14++ {
                                var v15 float32 = (0.1 / 1e10)
                                for fi16 := 0; fi16 < 17; fi16++ {
                                    _ = len(v4)
                                    _ = fi16
                                }
                                _ = fi14
                                _ = v15
                            }
                            for fi17 := 0; fi17 < 13; fi17++ {
                                var v18 float64 = ((1e10 + 0.0) * ((0.1 / -2.25) * 1e10))
                                var v19 int = (fi17 + (v3 % (((fi17 + (v3 / ((573) | 1)))) | 1)))
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        }
                    }
                }
            case -1:
                if (42 <= 307) {
                    v20 := GenS1{F0: 1, F1: 42, F2: v10}
                    v3 = v3
                    _ = v20
                }
            default:
                if (42 < 1) {
                    var v21 uint16 = 42
                    v10 -= int64(v3)
                    _ = v21
                } else {
                    fmt.Println(v2, v10)
                    v22 := GenS2{F0: "", F1: 9223372036854775807, F2: 0}
                    var v23 uint64 = 638
                    _ = v22
                    _ = v23
                }
            }
            switch v2 {
            case 979:
                var v24 uint16 = 255
                _ = v24
            case 4294967295:
                v3 += 0
            }
        }
        for fi25 := 0; fi25 < 8; fi25++ {
            var v26 bool = ((0.0 + -2.25) == 1e10)
            v9 = (v9 - v9)
            if false {
                v9 = (v9 ^ v9)
                if true {
                    v3 -= fi25
                } else {
                    for fi27 := 0; fi27 < 14; fi27++ {
                        var v28 bool = v26
                        _ = len(v4)
                        _ = fi27
                        _ = v28
                    }
                    fmt.Println(v10, v3)
                    switch "a" {
                    case "hello":
                        switch 266 {
                        case 32768:
                            v3 -= -9223372036854775808
                        case -1:
                            var v29 int64 = v10
                            _ = v29
                        default:
                            v1[4294967295] = "foo bar"
                        }
                    default:
                        v30 := GenS3{F0: 952, F1: 0}
                        _ = v30
                    }
                }
            }
            _ = fi25
            _ = v26
        }
        _ = v10
    } else {
        fmt.Println(v2)
        v3 += v3
    }
    v3 = v2
    var v31 uint = 0
    var v32 string = ("\n\t" + "\n\t")
    v2 += (v3 >> (((v31 % ((255) | 1))) & 63))
    v9 = v9
    v1[42] = v32
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v9
}
