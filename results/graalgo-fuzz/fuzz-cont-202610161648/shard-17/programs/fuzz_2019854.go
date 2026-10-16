package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint8; F1 string }

func main() {
    var v1 int = 0
    v2 := make(map[uint32]bool)
    var v3 uint64 = 18446744073709551615
    v4 := make(map[int]string)
    var v5 int16 = 0
    v6 := make(map[string]uint32)
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 float32 = 3.14159
        v8 -= (1e10 / 1.5)
        _ = fi7
        _ = v8
    }
    if (1 != 736) {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 uint16 = 247
            var v11 uint64 = (v3 * 18446744073709551615)
            fmt.Println(v1, v3)
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 int32 = 1
        v4[v1] = ("hello" + ("" + "hello"))
        _ = v12
    } else {
        v5 -= v5
    }
    var v13 int8 = 119
    v14 := GenS1{F0: false}
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 int8 = 43
            if false {
                v1 -= 4294967295
                v17 += -1
            }
            _ = fi16
            _ = v17
        }
        v13 -= (^(v13))
        switch (fi15 % (((v1 & v1)) | 1)) {
        case 42:
            var v18 int16 = v5
            _ = v18
        case 32768:
            for fi19 := 0; fi19 < 12; fi19++ {
                v20 := GenS2{F0: 273, F1: v13}
                var v21 int = (v1 % (((fi19 << ((357) & 63))) | 1))
                v22 := GenS1{F0: ((uint(v3) != 255) || (!false))}
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        case -2147483648:
            v23 := GenS3{F0: 128, F1: ""}
            _ = v23
        }
        _ = fi15
    }
    v5 = v5
    var v24 float64 = -2.25
    var v25 int16 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
