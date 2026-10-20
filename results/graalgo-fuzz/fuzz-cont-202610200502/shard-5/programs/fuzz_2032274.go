package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 int64; F2 uint64 }

func main() {
    var v1 uint32 = 42
    v2 := make(map[uint32]int)
    v3 := []uint32{v1, v1}
    fmt.Println(v1)
    var v4 bool = false
    switch 0 {
    case 2147483647:
        if (float64(0.1) > ((1.5 + (0.0 + 1e10)) * 1e10)) {
            delete(v2, (^(v1)))
        } else {
            var v5 bool = true
            v6 := GenS1{F0: v1}
            _ = v5
            _ = v6
        }
    case 4294967295:
        var v7 bool = false
        _ = v7
    case 542:
        var v8 uint8 = 2
        _ = v8
    }
    if (25 >= 103) {
        v9 := GenS2{F0: 32767}
        v4 = ((true && (1000 <= 344)) || true)
        _ = v9
    } else {
        if (-1 == 100) {
            v10 := GenS3{F0: 18446744073709551615}
            _ = v10
        } else {
            var v11 float64 = 1e10
            v11 = float64(0.0)
            v12 := GenS1{F0: (v1 | v1)}
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 13; fi13++ {
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            var v14 uint32 = (v1 % (((v1 - v1)) | 1))
            if (-468 > -145) {
                for fi15 := 0; fi15 < 19; fi15++ {
                    switch "hello" {
                    case "\n\t":
                        _ = len(v3)
                    case "hello":
                        fmt.Println(v14, v4, v14)
                    default:
                        v4 = v4
                    }
                    v16 := GenS3{F0: 18446744073709551615}
                    _ = fi15
                    _ = v16
                }
                if ("hello" < "\n\t") {
                    if true {
                        var v17 uint32 = 593
                        _ = v17
                    }
                    delete(v2, 65535)
                    fmt.Println(fi13, v14, v4)
                } else {
                    var v18 uint32 = 840
                    _ = v18
                }
                v3 = append(v3, (^(v1)))
            } else {
                var v19 uint8 = 128
                var v20 float32 = 3.14159
                _ = v19
                _ = v20
            }
            _ = fi13
            _ = v14
        }
        var v21 float32 = -2.25
        _ = v21
    }
    var v22 int = 127
    switch v22 {
    case 9:
        v22 = (v22 ^ (v22 / (((v22 % (((v22 & v22)) | 1))) | 1)))
    case 1:
        v4 = false
    }
    switch (("\n\t" + "\n\t") + "a") {
    case "foo bar":
        for fi23 := 0; fi23 < 7; fi23++ {
            v1 = (v1 & v1)
            _ = fi23
        }
    default:
        if ((255 < 319) && (!(42 >= 127))) {
            var v24 bool = v4
            _ = v24
        } else {
            _ = len(v3)
            for fi25 := 0; fi25 < 14; fi25++ {
                if (v4 && false) {
                    delete(v2, v1)
                    if v4 {
                        for fi26 := 0; fi26 < 12; fi26++ {
                            v27 := GenS1{F0: (v1 - v1)}
                            _ = fi26
                            _ = v27
                        }
                        v1 = v1
                    }
                }
                var v28 int32 = 0
                if ((255 < 10) || v4) {
                    for fi29 := 0; fi29 < 17; fi29++ {
                        var v30 float32 = float32(float64(float32((((0.1 + (3.14159 * (float64((-2.25 / ((3.14159 + (3.14159 - 1e10)) * 1e10))) + 3.14159))) / (0.1 + float64(0.0))) / -2.25))))
                        var v31 uint = 4294967295
                        v1 = v1
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                }
                _ = fi25
                _ = v28
            }
        }
    }
    var v32 uint8 = 128
    v33 := GenS4{F0: -9223372036854775808, F1: -1, F2: 18446744073709551615}
    var v34 uint = 42
    v22 = (-(v22))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
