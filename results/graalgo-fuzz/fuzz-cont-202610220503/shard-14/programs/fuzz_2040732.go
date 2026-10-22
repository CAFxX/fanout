package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int64; F2 uint16 }

func main() {
    v1 := []uint32{65535, 0}
    v2 := make(map[uint32]int)
    var v3 uint64 = 58
    v4 := make(map[string]string)
    switch "a" {
    case "a":
        v3 = v3
    case "hello":
        if true {
            if true {
                for fi5 := 0; fi5 < 8; fi5++ {
                    fmt.Println(fi5, v3)
                    v3 = 42
                    _ = fi5
                }
                fmt.Println(v3)
                v6 := GenS1{F0: 105, F1: 0}
                _ = v6
            }
            if (-474 >= 42) {
                for fi7 := 0; fi7 < 6; fi7++ {
                    var v8 uint32 = 65535
                    _ = fi7
                    _ = v8
                }
                if true {
                    var v9 uint64 = 4294967295
                    _ = v9
                }
            } else {
                v10 := GenS1{F0: 0, F1: 100}
                var v11 float64 = 3.14159
                _ = v10
                _ = v11
            }
        } else {
            var v12 bool = true
            fmt.Println(v3, v3)
            if (v12 && v12) {
                var v13 int = -140
                _ = v13
            } else {
                v14 := GenS1{F0: -1, F1: -1}
                if (v12 && v12) {
                    v12 = ((float64(3.14159) + 0.0) <= ((1.5 / (-2.25 - (-2.25 * (0.1 + float64((0.1 + 0.1)))))) * 1e10))
                    switch 2147483647 {
                    case -9223372036854775808:
                        if false {
                            _ = len(v2)
                            for fi15 := 0; fi15 < 9; fi15++ {
                                var v16 float64 = ((((1.5 / 1.0) + (0.1 / 0.1)) / (float64(-2.25) * (((1.5 * ((0.1 * (1e10 - (0.1 * 3.14159))) - 3.14159)) - 0.1) * ((0.0 - (1.5 - -2.25)) + float64(3.14159))))) + 0.1)
                                v3 -= v3
                                if v12 {
                                    var v17 uint = 1
                                    delete(v2, 244)
                                    _ = v17
                                }
                                _ = fi15
                                _ = v16
                            }
                            v1 = append(v1, 0)
                        }
                    default:
                        v12 = (233 == 1)
                    }
                    switch "" {
                    case "":
                        if len(v1) > 0 {
                            v1[0] = 784
                        }
                    case "\n\t":
                        fmt.Println(v12, v3)
                    }
                }
                _ = v14
            }
            _ = v12
        }
    default:
        var v18 int = 553
        _ = v18
    }
    v19 := GenS1{F0: 0, F1: -1}
    if len(v1) > 0 {
        v1[0] = 42
    }
    switch 41 {
    case 32767:
        v20 := GenS1{F0: -1, F1: 81}
        _ = v20
    }
    fmt.Println(v3)
    v21 := GenS1{F0: 82, F1: -1}
    fmt.Println(v3)
    v3 += v3
    v22 := GenS2{F0: "a", F1: -104}
    v3 = (v3 ^ v3)
    var v23 int8 = -1
    if false {
        v24 := GenS3{F0: 1}
        v2[2147483647] = -2147483648
        v25 := GenS4{F0: -127, F1: -893, F2: 255}
        _ = v24
        _ = v25
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
