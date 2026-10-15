package main

import (
    "fmt"
)


func main() {
    var v1 int32 = 1
    v2 := make(map[int]int)
    var v3 uint = 42
    var v4 uint16 = 2
    var v5 int64 = 9223372036854775807
    v4 = v4
    var v6 float32 = ((((float32(1.5) - ((-2.25 + -2.25) * -2.25)) - 1.5) * (3.14159 * 0.1)) - 0.0)
    if true {
        v4 = 65535
        v6 = 0.1
    }
    var v7 uint64 = 18446744073709551615
    fmt.Println(v4, v5)
    var v8 int64 = -1
    var v9 bool = (!true)
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 int = fi10
        _ = fi10
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
