package main

import (
    "fmt"
)


func main() {
    var v1 int64 = -755
    var v2 float32 = 0.0
    v3 := []int{-9223372036854775808, -71}
    var v4 uint32 = 4294967295
    v5 := []string{}
    var v6 float64 = 1e10
    v2 -= v2
    var v7 int8 = -4
    fmt.Println(v7, v2)
    for fi8 := 0; fi8 < 16; fi8++ {
        v6 = (0.0 * 1e10)
        fmt.Println(v1, v6, v6)
        v2 = v2
        _ = fi8
    }
    var v9 uint8 = 128
    var v10 string = "hello"
    var v11 int = 4294967295
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 float64 = 0.0
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi12
        _ = v13
    }
    switch v10 {
    case "a":
        _ = len(v3)
    case "hello":
        if true {
            fmt.Println(v1, v4)
            fmt.Println(v4, v7, v11)
            var v14 float32 = v2
            _ = v14
        }
    case "\n\t":
        v11 += v11
    default:
        var v15 string = "foo bar"
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
