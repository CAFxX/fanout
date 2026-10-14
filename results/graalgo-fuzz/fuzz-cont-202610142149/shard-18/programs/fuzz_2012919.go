package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 string }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint32 }

func main() {
    v1 := []uint32{1, 42}
    var v2 float32 = ((3.14159 / (0.1 - (((1.5 + 0.1) * ((((float32(0.1) / (3.14159 - 0.1)) + (0.1 / 3.14159)) - 3.14159) * 3.14159)) - ((1.5 * 1e10) * (-2.25 * 3.14159))))) - ((float32(0.1) + float32(-2.25)) * 0.1))
    var v3 int = 466
    var v4 float64 = 1e10
    switch (v3 & 572) {
    case 9223372036854775807:
        v5 := GenS1{F0: 773, F1: 71, F2: "a"}
        _ = v5
    case -2147483648:
        fmt.Println(v3, v3)
    }
    v3 = v3
    v4 = float64(float32(-2.25))
    v6 := GenS1{F0: 2147483647, F1: -1, F2: "hello"}
    fmt.Println(v2, v4)
    v7 := GenS1{F0: 332, F1: 127, F2: "a"}
    v3 += v3
    v8 := GenS2{F0: v3, F1: 11}
    v9 := GenS3{F0: 0, F1: 0, F2: 899}
    if true {
        var v10 float32 = (float32(v4) + v2)
        _ = v10
    }
    var v11 string = "a"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
