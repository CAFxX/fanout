package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int8 = 0
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    v3[""] = 42
    var v4 string = (("" + ("\n\t" + "")) + "\n\t")
    delete(v2, -387)
    v4 = ((v4 + v4) + "a")
    var v5 float64 = float64(0.1)
    v6 := GenS1{F0: 644}
    delete(v2, -587)
    var v7 int = 128
    if (234 <= 127) {
        var v8 float64 = float64((3.14159 * 0.1))
        var v9 int8 = v1
        var v10 uint32 = 2147483647
        _ = v10
        _ = v8
        _ = v9
    } else {
        fmt.Println(v4, v7, v7)
    }
    fmt.Println(v7, v4, v7)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
