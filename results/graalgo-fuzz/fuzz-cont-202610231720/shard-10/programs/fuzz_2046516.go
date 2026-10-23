package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint32 = 438
    var v3 int = 0
    v4 := make(map[uint32]bool)
    var v5 uint64 = 18446744073709551615
    v6 := []string{"a", "a", "\n\t"}
    var v7 float32 = 1e10
    switch "" {
    case "hello":
        switch (v3 >> ((996) & 63)) {
        case 128:
            _ = len(v6)
        }
    default:
        fmt.Println(v3)
    }
    _ = len(v6)
    var v8 uint64 = 0
    v7 += 1.5
    _ = len(v6)
    fmt.Println(v2, v5, v8)
    fmt.Println(v1)
    v6 = append(v6, "foo bar")
    v8 -= 4294967295
    if (0 >= -2147483648) {
        switch "\n\t" {
        case "hello":
            v9 := GenS1{F0: ((("" + "a") + "a") + "foo bar"), F1: -65, F2: 2}
            _ = v9
        }
        switch v3 {
        case -9223372036854775808:
            for fi10 := 0; fi10 < 12; fi10++ {
                var v11 bool = true
                fmt.Println(v2, v2)
                _ = fi10
                _ = v11
            }
        case 127:
            var v12 int16 = -704
            _ = v12
        }
    }
    v7 = (v7 + 0.1)
    v7 -= (1e10 / float32(1e10))
    var v13 uint16 = 42
    fmt.Println(v3, v13)
    var v14 uint64 = (v5 * (^(v5)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
