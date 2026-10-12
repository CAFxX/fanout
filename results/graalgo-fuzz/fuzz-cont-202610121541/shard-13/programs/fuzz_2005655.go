package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = -232
    var v3 int16 = -683
    v4 := []int{127}
    var v5 uint8 = 127
    v3 += int16(v2)
    v3 = (v3 << ((1) & 63))
    fmt.Println(v5)
    if (2147483647 < 0) {
        fmt.Println(v3, v5, v5)
        v2 -= v2
        switch 0 {
        case 1:
            fmt.Println(v2)
        case 32767:
            _ = len(v1)
        }
    }
    var v6 string = "a"
    v7 := GenS1{F0: -9223372036854775808}
    v6 = v6
    v1[1] = 679
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        v3 += v3
    case "hello":
        var v8 float32 = (0.1 * 0.0)
        _ = v8
    case "a":
        var v9 uint8 = (v5 % ((v5) | 1))
        _ = v9
    default:
        var v10 uint32 = 87
        _ = v10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
