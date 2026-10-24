package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int16 }

func main() {
    var v1 int64 = 0
    var v2 uint = 2
    var v3 uint = (v2 - (v2 - (^(v2))))
    v4 := make(map[string]uint32)
    var v5 float32 = 0.0
    v6 := make(map[string]bool)
    if true {
        var v7 int8 = 127
        var v8 bool = (v2 != 369)
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: -1, F1: 0, F2: 1000}
    v6[""] = (!(!false))
    fmt.Println(v3, v3, v1)
    var v10 uint8 = 1
    v10 += (v10 ^ 110)
    v11 := v6["hello"]
    var v12 uint8 = v10
    var v13 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
