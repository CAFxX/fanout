package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint8 = 128
    v2 := make(map[string]string)
    v3 := make(map[int]uint32)
    var v4 float32 = 1.5
    v4 -= v4
    fmt.Println(v4, v1)
    if (4294967295 > 18446744073709551615) {
        v1 += (v1 * 1)
    }
    v1 = v1
    fmt.Println(v4, v1)
    var v5 uint16 = 32767
    v6 := GenS1{F0: 0}
    v7 := GenS1{F0: 997}
    v4 = v4
    switch "hello" {
    case "hello":
        v4 -= (0.0 * -2.25)
    case "\n\t":
        v8 := []int{-444, 127}
        _ = v8
    case "foo bar":
        v1 -= (v1 & (^(v1)))
    default:
        var v9 bool = (!true)
        _ = v9
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
