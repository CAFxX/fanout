package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 string = "foo bar"
    var v2 float64 = 1e10
    v3 := []uint32{574, 65535, 346}
    v4 := make(map[string]string)
    var v5 string = v1
    v5 = v5
    var v6 int8 = 42
    fmt.Println(v5, v6, v6)
    switch 103 {
    case 32768:
        switch v5 {
        case "a":
            v7 := GenS1{F0: 42}
            _ = v7
        case "foo bar":
            fmt.Println(v5, v1)
        }
    default:
        var v8 int8 = v6
        _ = v8
    }
    var v9 uint = 18446744073709551615
    var v10 uint32 = 4294967295
    v4[v1] = (v1 + v5)
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 uint8 = 42
        _ = fi11
        _ = v12
    }
    fmt.Println(v6, v2, v2)
    v10 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
