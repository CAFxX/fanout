package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 127
    v2 := make(map[string]uint32)
    var v3 bool = (0.0 >= 1e10)
    var v4 float32 = float32(1.5)
    var v5 float32 = (0.1 * v4)
    var v6 int32 = -926
    switch ("hello" + "") {
    case "a":
        var v7 uint8 = v1
        _ = v7
    case "foo bar":
        v3 = v3
    }
    v2["a"] = 232
    var v8 uint32 = 2
    switch "a" {
    case "foo bar":
        fmt.Println(v4)
    }
    var v9 int32 = v6
    fmt.Println(v5)
    v3 = v3
    var v10 uint32 = 708
    v4 += v5
    fmt.Println(v3, v5, v3)
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := v2["foo bar"]
        _ = fi11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
