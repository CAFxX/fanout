package main

import (
    "fmt"
)


func main() {
    var v1 float32 = ((1e10 - -2.25) / (-2.25 / 1.5))
    var v2 float32 = (float32(((1e10 * 0.0) + (-2.25 * 0.1))) - 1.5)
    v3 := []int{32767, 127, 128}
    v3 = append(v3, 42)
    switch "hello" {
    case "\n\t":
        v1 += v2
    case "a":
        for fi4 := 0; fi4 < 19; fi4++ {
            v1 += v2
            _ = fi4
        }
    case "":
        var v5 uint64 = 4294967295
        _ = v5
    }
    var v6 int8 = 86
    fmt.Println(v6)
    var v7 float64 = (1.5 - (((1.5 + 3.14159) / 3.14159) / 1.0))
    fmt.Println(v7, v7)
    v3 = append(v3, 128)
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 int = fi8
        v2 = v2
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 int8 = v6
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    var v12 bool = true
    var v13 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
