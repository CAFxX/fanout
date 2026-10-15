package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int8 }

func main() {
    var v1 float64 = 3.14159
    var v2 float32 = (-2.25 + (1e10 + (-2.25 - 0.0)))
    var v3 string = "hello"
    var v4 uint = 127
    var v5 int32 = -651
    v6 := []int{4294967295, 9223372036854775807, -1, 558}
    var v7 uint64 = 4294967295
    v8 := GenS1{F0: 54, F1: 2, F2: -94}
    fmt.Println(v5, v5)
    var v9 bool = (-9223372036854775808 <= 42)
    var v10 float64 = v1
    for fi11 := 0; fi11 < 18; fi11++ {
        fmt.Println(v5, v3)
        v5 = (v5 ^ (v5 / ((-320) | 1)))
        _ = fi11
    }
    if v9 {
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 uint32 = 1
            fmt.Println(v2)
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: 32767, F1: 127, F2: 127}
        fmt.Println(v2, v2)
        _ = v14
    } else {
        fmt.Println(v9, v7, v5)
        switch 32768 {
        case -832:
            var v15 int16 = -32768
            _ = v15
        case 4294967295:
            v1 = 3.14159
        default:
            switch "a" {
            case "foo bar":
                v5 = 42
            }
        }
        var v16 uint64 = (v7 ^ v7)
        _ = v16
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
