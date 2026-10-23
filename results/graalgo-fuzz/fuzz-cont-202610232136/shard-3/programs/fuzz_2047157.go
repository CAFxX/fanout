package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]string)
    v2 := []string{}
    var v3 uint = 0
    var v4 float64 = 1e10
    var v5 int16 = 1
    v2 = append(v2, "foo bar")
    var v6 int64 = 42
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 float64 = float64((((0.0 + 1.5) * 3.14159) + 3.14159))
        var v9 uint = (v3 & v3)
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int64 = 847
    if true {
        var v11 float64 = -2.25
        _ = v11
    } else {
        fmt.Println(v10)
    }
    v5 = (v5 ^ (v5 >> ((255) & 63)))
    var v12 uint8 = 127
    fmt.Println(v4, v10, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
