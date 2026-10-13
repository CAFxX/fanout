package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint16 }

func main() {
    var v1 int8 = 127
    var v2 string = "a"
    v3 := make(map[string]int)
    var v4 float64 = -2.25
    v5 := []string{((v2 + "") + (v2 + v2))}
    var v6 string = "hello"
    var v7 int8 = v1
    switch 839 {
    case -454:
        fmt.Println(v4, v7)
    case 127:
        v1 += v1
    }
    var v8 int64 = 0
    if (!true) {
        if len(v5) > 0 {
            v5[0] = "hello"
        }
    } else {
        switch v2 {
        case "a":
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 bool = (42 <= 255)
                fmt.Println(v1, v1, v1)
                _ = fi9
                _ = v10
            }
        case "hello":
            fmt.Println(v8, v7)
        default:
            var v11 uint8 = 128
            _ = v11
        }
    }
    var v12 int = 128
    v13 := GenS1{F0: 42, F1: 0, F2: 0}
    for fi14 := 0; fi14 < 9; fi14++ {
        v12 += fi14
        _ = fi14
    }
    v3[v2] = v12
    fmt.Println(v6, v12)
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 string = v6
        var v17 uint16 = 0
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint16 = 2
    v19 := GenS1{F0: v7, F1: 2147483648, F2: (^(v18))}
    v4 = (float64(-2.25) + v4)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
