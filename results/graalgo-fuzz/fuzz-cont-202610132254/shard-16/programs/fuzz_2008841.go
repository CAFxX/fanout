package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint8 = 127
    var v2 int64 = -721
    var v3 uint64 = 4294967295
    var v4 int64 = v2
    var v5 float32 = 0.1
    if (int8(v1) < -1) {
        var v6 uint64 = v3
        _ = v6
    }
    v2 = v2
    switch "a" {
    case "":
        v7 := GenS1{F0: (v1 * 247)}
        _ = v7
    case "\n\t":
        v1 = v1
    case "hello":
        fmt.Println(v5, v1)
    default:
        var v8 int = 32767
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        fmt.Println(v4)
        _ = fi9
    }
    v10 := GenS1{F0: v1}
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := []int{}
        fmt.Println(v3, v4, v4)
        _ = fi11
        _ = v12
    }
    v2 = v4
    var v13 int = 570
    var v14 float64 = (-2.25 / 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
