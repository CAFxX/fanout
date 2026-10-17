package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
func genFn0() int8 {
    v7 := GenS1{F0: 4294967295, F1: 0}
    var v8 float32 = 3.14159
    v9 := GenS1{F0: 1, F1: 4294967295}
    v8 += (1e10 / float32(-2.25))
    var v10 uint16 = 403
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 float64 = (0.0 * float64(1.5))
    var v2 uint = 0
    var v3 uint32 = 504
    var v4 int = 127
    var v5 uint32 = uint32(v2)
    var v6 uint16 = 65535
    var v11 float32 = ((0.1 / (float32(v1) / float32(float64((1.5 - 0.0))))) + float32(float64(float32(v1))))
    v4 = v4
    fmt.Println(v5)
    v1 = -2.25
    var v12 float32 = ((0.0 * v11) + v11)
    switch "hello" {
    case "":
        v6 = 42
    case "hello":
        v11 -= v12
    }
    var v13 int = v4
    v1 = -2.25
    v6 += 417
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
