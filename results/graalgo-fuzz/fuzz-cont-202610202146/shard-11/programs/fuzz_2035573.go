package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int }

func main() {
    var v1 float64 = 1.5
    var v2 string = "a"
    var v3 uint32 = 262
    var v4 int = 4294967295
    v3 += v3
    var v5 string = v2
    v1 -= (v1 + 3.14159)
    v4 = 42
    var v6 string = v2
    var v7 uint = 1
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 int8 = 1
        var v10 bool = true
        switch v5 {
        case "":
            v10 = ((true || v10) || (true || v10))
        case "hello":
            v7 -= uint(v3)
        case "foo bar":
            for fi11 := 0; fi11 < 9; fi11++ {
                v12 := GenS1{F0: false, F1: (v9 - v9), F2: (^(fi11))}
                var v13 uint = v7
                v9 -= v9
                _ = fi11
                _ = v12
                _ = v13
            }
        default:
            v10 = v10
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    if (true || (42 <= 32767)) {
        var v14 float32 = 3.14159
        var v15 int16 = 1000
        fmt.Println(v15)
        _ = v14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
