package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 float64 = 0.1
    v2 := []int{534}
    var v3 int8 = 42
    var v4 uint8 = 1
    fmt.Println(v4)
    v4 -= (v4 % ((v4) | 1))
    var v5 float32 = 3.14159
    var v6 uint32 = 2147483648
    switch "hello" {
    case "hello":
        var v7 uint16 = 2
        _ = v7
    case "foo bar":
        switch "" {
        case "foo bar":
            fmt.Println(v5, v4, v3)
        default:
            var v8 uint64 = 1
            _ = v8
        }
    case "a":
        for fi9 := 0; fi9 < 20; fi9++ {
            _ = len(v2)
            v1 += v1
            _ = fi9
        }
    default:
        var v10 int16 = -32768
        _ = v10
    }
    fmt.Println(v1, v1)
    v11 := GenS1{F0: (v6 / ((65535) | 1))}
    v6 -= 1
    var v12 uint16 = 0
    v6 += 677
    var v13 float32 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
