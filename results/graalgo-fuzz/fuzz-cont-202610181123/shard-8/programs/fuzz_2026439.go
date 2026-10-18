package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int32; F1 bool; F2 uint16 }

func main() {
    var v1 bool = false
    var v2 float32 = (float32(0.1) - 0.1)
    v3 := []string{(((("hello" + "\n\t") + "hello") + "a") + "\n\t"), "", ("foo bar" + ""), ("" + "\n\t")}
    var v4 string = "foo bar"
    if v1 {
        var v5 int64 = -9223372036854775808
        v6 := GenS1{F0: 0, F1: -2147483648, F2: -2147483648}
        v1 = (((-1 < -32768) || ((!((((v2 / (v2 * (float32((1e10 + 0.0)) + v2))) + ((float32(1e10) * v2) + 1e10)) > v2) && (!(105 < -128)))) && ((v5 >> ((42) & 63)) != (v5 >> ((42) & 63))))) || v1)
        _ = v5
        _ = v6
    }
    v1 = true
    v1 = (65535 <= 0)
    v7 := GenS2{F0: 255, F1: 42, F2: 360}
    if (535 > 836) {
        var v8 float32 = v2
        fmt.Println(v4, v4)
        _ = v8
    } else {
        v1 = ((!(v4 < (v4 + (v4 + (v4 + (((("\n\t" + "foo bar") + ((v4 + v4) + v4)) + "") + v4)))))) || false)
        fmt.Println(v2)
        v9 := GenS3{F0: 0, F1: v1, F2: 1}
        _ = v9
    }
    var v10 uint32 = 42
    v2 = 3.14159
    v11 := GenS1{F0: 346, F1: -996, F2: -2147483648}
    v2 = float32((0.1 / 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
