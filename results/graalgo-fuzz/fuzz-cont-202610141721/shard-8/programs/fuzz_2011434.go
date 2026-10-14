package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[int]uint32)
    var v3 uint8 = 128
    v4 := make(map[uint32]bool)
    fmt.Println(v3, v1)
    v3 -= (v3 / (((v3 | (v3 ^ 128))) | 1))
    v1 += float32(3.14159)
    switch -797 {
    case -2147483648:
        if (42 <= 2147483648) {
            v1 = v1
            v5 := GenS1{F0: 18}
            switch "" {
            case "foo bar":
                v3 = v3
            }
            _ = v5
        }
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        v3 += 128
        _ = fi6
    }
    switch "hello" {
    case "\n\t":
        v3 = v3
    case "":
        v3 = v3
    }
    v7 := GenS1{F0: 42}
    switch 2147483647 {
    case 0:
        var v8 uint16 = 42
        _ = v8
    default:
        var v9 int32 = 2147483647
        _ = v9
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
