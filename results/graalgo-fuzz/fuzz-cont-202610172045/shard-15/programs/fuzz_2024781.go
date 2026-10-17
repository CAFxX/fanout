package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 0.1
    v3 := []uint32{}
    var v4 int8 = 127
    v4 = (v4 & (v4 & v4))
    var v5 uint32 = 36
    var v6 uint64 = uint64(v5)
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint64 = v6
        if false {
            for fi9 := 0; fi9 < 11; fi9++ {
                if (true && ((false || (2 == 1)) || (3.14159 <= 1e10))) {
                    if len(v3) > 0 {
                        v3[0] = (v5 * 42)
                    }
                    fmt.Println(fi9, v2, fi7)
                    v2 -= 1e10
                }
                _ = len(v3)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi9
            }
            fmt.Println(v5)
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v2)
    v3 = append(v3, v5)
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS1{F0: -42, F1: v6}
        var v12 int32 = 0
        v12 += 2147483647
        _ = fi10
        _ = v11
        _ = v12
    }
    if len(v3) > 0 {
        v3[0] = (v5 * v5)
    }
    v13 := GenS1{F0: -1, F1: (v6 + 1)}
    var v14 float32 = float32(0.1)
    _ = len(v3)
    v5 += 672
    var v15 uint8 = 173
    if (v4 > 100) {
        v16 := GenS1{F0: v4, F1: v6}
        var v17 int8 = -128
        if (1 == 42) {
            var v18 int32 = 2147483647
            if (v18 != v18) {
                switch 9223372036854775807 {
                case 32768:
                    switch (("\n\t" + "") + "a") {
                    case "\n\t":
                        v2 = (1e10 / 1.5)
                    case "hello":
                        fmt.Println(v5, v4)
                    default:
                        fmt.Println(v17, v14, v17)
                    }
                }
                var v19 float64 = (0.1 / 0.1)
                _ = v19
            }
            _ = v18
        }
        _ = v16
        _ = v17
    }
    v4 = (v4 / (((v4 - v4)) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
