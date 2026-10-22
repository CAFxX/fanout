package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float32 = 3.14159
    var v2 uint8 = 118
    v3 := make(map[string]int)
    for fi4 := 0; fi4 < 15; fi4++ {
        if (0.0 <= ((1e10 / 1.5) * 1e10)) {
            var v5 int64 = 99
            var v6 float64 = 1.5
            _ = v5
            _ = v6
        }
        v7 := []int{(fi4 ^ (fi4 ^ (fi4 + fi4))), fi4, fi4, fi4}
        if true {
            for fi8 := 0; fi8 < 9; fi8++ {
                v2 = (v2 * 2)
                _ = fi8
            }
        }
        _ = fi4
        _ = v7
    }
    v9 := GenS1{F0: 65535, F1: 4294967295}
    v10 := GenS1{F0: 0, F1: 0}
    v2 += v2
    fmt.Println(v1, v2)
    var v11 int = 9223372036854775807
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := v3["foo bar"]
        var v14 int64 = -1
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS1{F0: 915, F1: 255}
        var v17 int16 = -32768
        _ = fi15
        _ = v16
        _ = v17
    }
    v1 = (v1 - (((float32((((0.0 * 1.5) * (3.14159 - (1e10 / float64(0.1)))) / ((1.5 * 0.1) + (float64(0.0) * 0.1)))) / v1) + (3.14159 - float32((0.0 + 3.14159)))) - v1))
    v11 += v11
    v18 := GenS2{F0: 4294967295}
    var v19 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v9
}
