package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 string; F2 int32 }
type GenS3 struct { F0 int32; F1 int8; F2 uint64 }
func genFn0() uint32 {
    if (float64(((float32(float64(((float32((1.5 - float64(3.14159))) / (1.5 + (((0.1 / 1.0) - 0.1) / (1e10 / ((0.1 + 0.0) / (-2.25 + 1e10)))))) / 1.5))) * -2.25) / 1.0)) > 3.14159) {
        var v6 int32 = 564
        fmt.Println(v6)
        var v7 uint32 = 277
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 65535}
    v9 := GenS1{F0: 0}
    _ = v8
    _ = v9
    return 2147483648
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint = 42
    var v3 float32 = 0.1
    v4 := []uint32{544}
    var v5 int64 = -9223372036854775808
    for fi10 := 0; fi10 < 9; fi10++ {
        if false {
            var v11 float32 = 0.1
            _ = v11
        } else {
            var v12 float32 = v3
            _ = v12
        }
        if (true && false) {
            if len(v4) > 0 {
                v4[0] = 0
            }
            switch (fi10 | fi10) {
            case -1:
                if false {
                    switch (fi10 ^ fi10) {
                    case -434:
                        var v13 uint64 = 42
                        _ = v13
                    default:
                        v2 = v2
                    }
                } else {
                    v1 = (v5 + v5)
                }
            default:
                _ = len(v4)
            }
        } else {
            var v14 uint16 = 1
            for fi15 := 0; fi15 < 7; fi15++ {
                if (!("hello" == "a")) {
                    v14 -= 255
                    var v16 int64 = v5
                    _ = v16
                }
                v5 = (^(v5))
                v2 += v2
                _ = fi15
            }
            _ = v14
        }
        fmt.Println(v3)
        _ = fi10
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 int8 = 122
        var v19 uint16 = 0
        _ = fi17
        _ = v18
        _ = v19
    }
    if true {
        for fi20 := 0; fi20 < 6; fi20++ {
            v21 := GenS1{F0: 255}
            var v22 float64 = 1e10
            _ = fi20
            _ = v21
            _ = v22
        }
        var v23 uint16 = 42
        var v24 uint8 = 104
        _ = v23
        _ = v24
    } else {
        var v25 float64 = 0.1
        var v26 int = 9223372036854775807
        v26 += v26
        _ = v25
        _ = v26
    }
    switch "" {
    case "":
        v27 := GenS2{F0: 1, F1: ("a" + "a"), F2: 2147483647}
        _ = v27
    case "a":
        v1 = -9223372036854775808
    }
    v2 -= 2
    var v28 int = -9223372036854775808
    fmt.Println(v5, v2, v1)
    fmt.Println(v3, v28)
    v1 = 431
    v29 := GenS3{F0: 956, F1: -128, F2: 18446744073709551615}
    for fi30 := 0; fi30 < 7; fi30++ {
        v1 += v5
        _ = fi30
    }
    var v31 int64 = int64(v2)
    v32 := GenS3{F0: 42, F1: 0, F2: 0}
    var v33 float64 = 3.14159
    v1 -= (v5 + v5)
    v28 = (^(v28))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
