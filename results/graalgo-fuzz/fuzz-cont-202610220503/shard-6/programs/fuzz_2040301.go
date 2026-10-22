package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int = -330
    v2 := make(map[int]uint32)
    var v3 string = "foo bar"
    var v4 uint32 = 42
    var v5 float64 = (1e10 - float64(((1e10 / 1.5) / 1.0)))
    v6 := GenS1{F0: 1}
    var v7 float64 = 3.14159
    switch v3 {
    case "a":
        fmt.Println(v1, v4)
    case "":
        v8 := v2[32768]
        _ = v8
    default:
        v1 -= v1
    }
    var v9 int64 = 9223372036854775807
    switch ("foo bar" + "\n\t") {
    case "":
        var v10 bool = false
        _ = v10
    default:
        for fi11 := 0; fi11 < 11; fi11++ {
            fmt.Println(v3, v5)
            var v12 uint16 = 65535
            _ = fi11
            _ = v12
        }
    }
    var v13 bool = true
    var v14 bool = v13
    var v15 string = (v3 + v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
