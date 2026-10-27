package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }

func main() {
    var v1 float64 = (1.5 - 0.1)
    var v2 int = 830
    var v3 float32 = float32(float64((((((3.14159 - 3.14159) + (1.5 * 3.14159)) - (3.14159 / float32(v1))) - (0.1 - (1.5 + 0.1))) * 0.0)))
    var v4 uint8 = 1
    var v5 uint16 = 42
    v6 := make(map[string]uint32)
    var v7 int8 = 127
    v3 = v3
    delete(v6, "foo bar")
    fmt.Println(v5, v7)
    v5 = (v5 + v5)
    switch ("hello" + "") {
    case "a":
        var v8 int16 = 0
        _ = v8
    case "foo bar":
        fmt.Println(v4)
    case "hello":
        fmt.Println(v3, v4, v1)
    default:
        v5 -= (v5 >> ((18446744073709551615) & 63))
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint = 417
        _ = fi9
        _ = v10
    }
    if true {
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 int = (fi11 / ((v2) | 1))
            _ = len(v6)
            _ = fi11
            _ = v12
        }
    }
    v3 = v3
    v13 := GenS1{F0: 2, F1: 557}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
