package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint8 }

func main() {
    v1 := []uint32{}
    var v2 int32 = 42
    var v3 int = 32767
    var v4 uint = 1
    var v5 float32 = 1.5
    var v6 float64 = (float64(v5) - (-2.25 - -2.25))
    v7 := GenS1{F0: v4, F1: int8(v3)}
    if false {
        v4 = (v4 * v4)
        for fi8 := 0; fi8 < 8; fi8++ {
            if (false && (v6 > (((((0.1 / (float64(0.1) + -2.25)) / 1.0) - (1.5 * (((0.0 / 3.14159) * (v6 / v6)) + (0.0 - v6)))) / v6) * v6))) {
                v9 := GenS1{F0: v4, F1: 100}
                for fi10 := 0; fi10 < 10; fi10++ {
                    switch (("\n\t" + (("hello" + "hello") + "foo bar")) + "a") {
                    case "\n\t":
                        var v11 int32 = v2
                        _ = v11
                    case "":
                        if false {
                            v2 -= (v2 % (((v2 ^ v2)) | 1))
                            var v12 float32 = v5
                            var v13 float32 = float32(v6)
                            _ = v12
                            _ = v13
                        }
                    }
                    fmt.Println(fi10)
                    _ = fi10
                }
                var v14 uint8 = 2
                _ = v14
                _ = v9
            }
            _ = fi8
        }
        var v15 uint64 = 303
        _ = v15
    }
    var v16 float64 = 1.5
    for fi17 := 0; fi17 < 18; fi17++ {
        fmt.Println(v3, fi17)
        if ("foo bar" < ((("\n\t" + "") + "\n\t") + "foo bar")) {
            if (!false) {
                switch (fi17 + (fi17 - (v3 << (((v4 | (v4 % (((v4 << (((^(v4))) & 63))) | 1)))) & 63)))) {
                case 128:
                    var v18 int16 = 1000
                    _ = v18
                case -2147483648:
                    if true {
                        var v19 uint8 = 2
                        var v20 uint32 = 2
                        var v21 int32 = (v2 << (((v4 * v4)) & 63))
                        _ = v19
                        _ = v20
                        _ = v21
                    } else {
                        v22 := GenS1{F0: 637, F1: -78}
                        _ = v22
                    }
                case 32767:
                    var v23 uint8 = 32
                    _ = v23
                default:
                    var v24 int16 = 0
                    _ = v24
                }
                if false {
                    var v25 float64 = (v6 - v16)
                    var v26 uint32 = 1
                    _ = v25
                    _ = v26
                } else {
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                    fmt.Println(v16, v3, v2)
                }
            } else {
                v27 := GenS2{F0: (v4 ^ (v4 - (v4 * (v4 | v4))))}
                _ = v27
            }
            var v28 int8 = -124
            _ = v28
        } else {
            v1 = append(v1, 41)
        }
        _ = fi17
    }
    var v29 int32 = v2
    var v30 string = "foo bar"
    var v31 uint32 = 42
    for fi32 := 0; fi32 < 12; fi32++ {
        var v33 string = v30
        var v34 int64 = 1
        fmt.Println(v34, v33)
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 float64 = v16
    v6 += ((((v16 + v16) / (float64(float32(v35)) * float64((v5 - (1e10 / (((v5 - float32(((v35 / 0.1) * float64((0.0 + 1.5))))) + v5) + v5)))))) * v16) + v6)
    switch (v30 + (v30 + ("foo bar" + v30))) {
    case "foo bar":
        v30 = (v30 + (v30 + "hello"))
    }
    v36 := GenS3{F0: v30, F1: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
