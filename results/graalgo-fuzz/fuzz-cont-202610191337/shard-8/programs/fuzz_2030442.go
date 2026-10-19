package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]uint32)
    var v3 uint64 = 861
    var v4 int16 = 32767
    var v5 uint = 881
    var v6 float32 = (1.5 - 0.1)
    fmt.Println(v3)
    switch 1 {
    case -2147483648:
        var v7 bool = (2147483648 <= 194)
        _ = v7
    case -9223372036854775808:
        v6 += (float32(v1) / v6)
    default:
        v6 += float32((float64(v6) * 0.0))
    }
    fmt.Println(v4, v5)
    var v8 float32 = (v6 / 1.5)
    v9 := []int{128, 128, 259}
    var v10 string = (("\n\t" + (("" + "foo bar") + "hello")) + "\n\t")
    v10 = "hello"
    var v11 uint64 = (v3 | (v3 / ((v3) | 1)))
    v1 = float64(3.14159)
    if len(v9) > 0 {
        v9[0] = 456
    }
    switch "" {
    case "hello":
        v9 = append(v9, 27)
    }
    switch v10 {
    case "foo bar":
        if false {
            fmt.Println(v11, v5, v11)
            for fi12 := 0; fi12 < 4; fi12++ {
                if len(v9) > 0 {
                    v9[0] = fi12
                }
                _ = fi12
            }
        }
    }
    if true {
        v13 := GenS1{F0: 2, F1: 127}
        v5 -= v5
        v14 := GenS1{F0: 0, F1: -2147483648}
        _ = v13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
