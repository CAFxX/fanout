package main

import (
    "fmt"
)


func main() {
    var v1 float32 = (1.5 / 0.1)
    var v2 uint32 = 4294967295
    var v3 uint32 = v2
    var v4 bool = (74 == 195)
    fmt.Println(v4, v2, v2)
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 float64 = 3.14159
        var v7 int16 = 42
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 float64 = 1.5
    var v9 bool = (928 <= 0)
    v4 = v9
    fmt.Println(v8, v9)
    if true {
        v3 = v2
        var v10 int16 = -32768
        v1 += (v1 + 1.5)
        _ = v10
    }
    var v11 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
