package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS2 struct { F0 string }

func main() {
    var v1 float64 = (0.0 + (1e10 - 0.0))
    var v2 float64 = float64(1e10)
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    v5 := []uint32{}
    var v6 uint = 1
    v2 += v1
    v1 += (v1 + ((v1 + v2) / v1))
    fmt.Println(v6, v6)
    var v7 uint8 = 255
    if (true || (3.14159 >= 3.14159)) {
        v3["foo bar"] = ((v7 + v7) != (v7 >> (((v6 | v6)) & 63)))
        v8 := GenS1{F0: 4294967295, F1: uint16(v6), F2: 4294967295}
        _ = v8
    }
    if true {
        var v9 int32 = -160
        if true {
            for fi10 := 0; fi10 < 11; fi10++ {
                v5 = append(v5, 512)
                _ = fi10
            }
            v5 = append(v5, 1)
        } else {
            var v11 int = -9223372036854775808
            v5 = append(v5, 4294967295)
            v2 = 1e10
            _ = v11
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v9
    }
    var v12 float32 = float32((1.5 / v2))
    fmt.Println(v6, v1, v1)
    var v13 uint = (v6 ^ (v6 | (v6 >> ((559) & 63))))
    v14 := GenS2{F0: ""}
    fmt.Println(v12, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
