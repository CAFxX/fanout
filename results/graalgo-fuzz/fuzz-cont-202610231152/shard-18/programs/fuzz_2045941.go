package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 2
    var v2 float32 = float32((3.14159 - ((1e10 * 1.5) - 3.14159)))
    var v3 uint64 = 2
    v4 := make(map[string]string)
    var v5 float64 = (3.14159 - 0.1)
    var v6 int = -9223372036854775808
    fmt.Println(v2, v1, v3)
    var v7 uint16 = 42
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 int = fi8
            _ = fi9
            _ = v10
        }
        switch "hello" {
        case "\n\t":
            v11 := []int{fi8, (v6 / (((fi8 / (((^(v6))) | 1))) | 1)), 32768}
            _ = v11
        case "":
            v12 := v4[""]
            _ = v12
        case "hello":
            var v13 int8 = 1
            _ = v13
        default:
            fmt.Println(v1, v3, v7)
        }
        fmt.Println(v3, v3)
        _ = fi8
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 uint8 = 42
        _ = fi14
        _ = v15
    }
    v16 := []int{-9223372036854775808, (v6 & (v6 << ((255) & 63)))}
    _ = len(v16)
    delete(v4, (("" + "foo bar") + "a"))
    fmt.Println(v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
