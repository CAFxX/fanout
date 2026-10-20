package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 float64 = (0.0 / ((1.5 - (3.14159 - 0.1)) + (float64(1e10) / 1.5)))
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]bool)
    var v4 string = "\n\t"
    var v5 uint8 = 1
    if true {
        var v6 uint16 = 65535
        _ = v6
    }
    v5 = (v5 & 88)
    var v7 bool = ((v5 & (v5 - v5)) <= (v5 >> ((791) & 63)))
    var v8 string = v4
    v1 -= v1
    var v9 uint = 251
    v10 := GenS1{F0: -128}
    v11 := GenS1{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
