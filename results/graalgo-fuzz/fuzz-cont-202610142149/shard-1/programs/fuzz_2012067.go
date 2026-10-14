package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int32 }
type GenS2 struct { F0 int16; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 string = "\n\t"
    var v2 float64 = 3.14159
    v3 := []uint32{4294967295, 772, 42, 539}
    v2 -= (v2 - (v2 / (v2 / ((float64((3.14159 * 0.1)) * ((v2 + v2) * (-2.25 + (3.14159 * (v2 - 0.1))))) * v2))))
    _ = len(v3)
    switch v1 {
    case "foo bar":
        var v4 int64 = 42
        _ = v4
    }
    var v5 float32 = ((1.5 + ((float32(-2.25) - 3.14159) + (0.0 + 0.1))) + ((0.0 * (((1.5 / (float32(1e10) + -2.25)) * 3.14159) + 1e10)) + (1e10 / 1.0)))
    v3 = append(v3, 0)
    v6 := GenS1{F0: 255, F1: -9223372036854775808, F2: -1}
    v7 := GenS2{F0: 146, F1: true, F2: 122}
    var v8 int64 = -69
    if (65535 != 1) {
        for fi9 := 0; fi9 < 5; fi9++ {
            v5 = float32(-2.25)
            if (v8 < v8) {
                var v10 uint8 = 255
                fmt.Println(v10, v8)
                v11 := GenS1{F0: 65535, F1: (fi9 * (fi9 >> ((0) & 63))), F2: -110}
                _ = v10
                _ = v11
            }
            _ = fi9
        }
        v12 := GenS3{F0: 437}
        _ = v12
    }
    fmt.Println(v1, v5, v1)
    var v13 uint = 1
    if true {
        v14 := GenS3{F0: 42}
        v15 := GenS2{F0: -1, F1: false, F2: 1}
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 uint64 = 714
            _ = fi16
            _ = v17
        }
        _ = v14
        _ = v15
    }
    var v18 int64 = (v8 * (v8 << ((v13) & 63)))
    v2 += float64(1.5)
    var v19 uint = (v13 % (((^(v13))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
