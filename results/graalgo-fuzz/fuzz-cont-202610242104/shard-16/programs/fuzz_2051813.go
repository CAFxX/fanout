package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int64; F1 uint64; F2 uint32 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{}
    var v3 uint64 = 593
    var v4 int = -9223372036854775808
    v5 := make(map[uint32]string)
    if ((65535 <= 32767) || true) {
        for fi6 := 0; fi6 < 6; fi6++ {
            for fi7 := 0; fi7 < 3; fi7++ {
                v4 = (fi7 ^ 127)
                switch "" {
                case "":
                    v5[348] = "foo bar"
                    if false {
                        switch "a" {
                        case "a":
                            var v8 int32 = 0
                            var v9 float64 = 3.14159
                            _ = v8
                            _ = v9
                        case "hello":
                            fmt.Println(v4, fi6)
                        default:
                            var v10 int16 = -821
                            _ = v10
                        }
                        v3 = 4294967295
                        var v11 uint32 = 4294967295
                        var v12 float32 = 1e10
                        switch (v4 - v4) {
                        case -1:
                            _ = len(v2)
                        default:
                            fmt.Println(fi7)
                        }
                        _ = v11
                        _ = v12
                    }
                case "foo bar":
                    fmt.Println(v3, v3)
                case "\n\t":
                    v2 = append(v2, "hello")
                default:
                    _ = len(v5)
                }
                var v13 int64 = 1
                _ = fi7
                _ = v13
            }
            fmt.Println(v4, v3, fi6)
            var v14 int16 = 32767
            _ = fi6
            _ = v14
        }
        v4 += 9223372036854775807
        v4 -= v4
    }
    if (-88 >= 127) {
        v4 -= v4
    } else {
        var v15 uint16 = 349
        _ = v15
    }
    v16 := GenS1{F0: v3, F1: (v4 << ((2) & 63))}
    if (!(!true)) {
        var v17 bool = (128 == 216)
        v17 = v17
        _ = v17
    }
    var v18 int = (v4 / (((v4 | -9223372036854775808)) | 1))
    v19 := GenS1{F0: 18446744073709551615, F1: (-(v18))}
    if (42 >= 32767) {
        var v20 bool = (v4 >= -311)
        v21 := GenS1{F0: (^(v3)), F1: v4}
        v22 := GenS2{F0: 0, F1: (v3 + v3), F2: 2147483647}
        _ = v20
        _ = v21
        _ = v22
    } else {
        for fi23 := 0; fi23 < 6; fi23++ {
            delete(v5, 1)
            _ = fi23
        }
    }
    var v24 uint32 = 2
    for fi25 := 0; fi25 < 12; fi25++ {
        fmt.Println(v4, v18, v4)
        _ = len(v1)
        v26 := GenS1{F0: v3, F1: fi25}
        _ = fi25
        _ = v26
    }
    if (71 >= 255) {
        for fi27 := 0; fi27 < 4; fi27++ {
            v18 -= fi27
            var v28 int64 = 476
            v4 -= (v18 & (^(v4)))
            _ = fi27
            _ = v28
        }
        var v29 int = (v18 | 9223372036854775807)
        _ = v29
    }
    var v30 int = (v18 | v18)
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
