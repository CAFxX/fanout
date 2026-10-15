package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }

func main() {
    var v1 bool = false
    v2 := []uint32{65535, 0, 65535, 2}
    var v3 int16 = 32767
    var v4 int64 = 0
    v4 -= v4
    fmt.Println(v1)
    v5 := GenS1{F0: v1, F1: 255}
    var v6 uint32 = 42
    fmt.Println(v6, v3, v4)
    switch "hello" {
    case "hello":
        v6 -= v6
    default:
        for fi7 := 0; fi7 < 1; fi7++ {
            if v1 {
                _ = len(v2)
            } else {
                v8 := GenS1{F0: ((!(-863 <= -1)) && (v4 < (-(v4)))), F1: 48}
                var v9 int64 = (v4 / (((v4 & v4)) | 1))
                _ = v8
                _ = v9
            }
            var v10 float64 = float64(0.1)
            _ = fi7
            _ = v10
        }
    }
    var v11 float32 = ((-2.25 * 1.5) / (1e10 + 3.14159))
    fmt.Println(v3)
    v3 = (v3 >> ((18446744073709551615) & 63))
    _ = len(v2)
    var v12 uint = 109
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
