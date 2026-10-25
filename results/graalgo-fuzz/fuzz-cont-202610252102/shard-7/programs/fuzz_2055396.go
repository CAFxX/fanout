package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }

func main() {
    var v1 float64 = (3.14159 * 1.5)
    var v2 uint8 = 228
    var v3 uint64 = 18446744073709551615
    v4 := []string{"a"}
    switch 1 {
    case 127:
        v1 -= ((float64(-2.25) / v1) - 0.1)
    }
    switch 128 {
    case 32768:
        var v5 int = 128
        _ = v5
    case 108:
        var v6 uint = 467
        _ = v6
    }
    var v7 int16 = 42
    fmt.Println(v3)
    v4 = append(v4, "")
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: v3, F1: 104, F2: 42}
        var v10 string = ("foo bar" + ("hello" + "\n\t"))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint8 = 0
    switch -786 {
    case 32767:
        switch 919 {
        case 310:
            for fi12 := 0; fi12 < 4; fi12++ {
                var v13 int8 = 0
                v11 = v11
                fmt.Println(v13, v1, v13)
                _ = fi12
                _ = v13
            }
        default:
            v14 := GenS1{F0: v3, F1: 839, F2: -962}
            _ = v14
        }
    }
    v1 = v1
    var v15 uint16 = 42
    v15 += v15
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
