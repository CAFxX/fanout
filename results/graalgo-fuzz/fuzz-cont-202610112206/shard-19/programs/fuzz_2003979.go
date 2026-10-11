package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 float32 = ((3.14159 / (1e10 - (0.1 + 0.1))) / 3.14159)
    var v2 int = -7
    v3 := make(map[int]int)
    if true {
        v4 := GenS1{F0: 23}
        var v5 int = (^(v2))
        _ = v4
        _ = v5
    } else {
        v6 := GenS1{F0: 65535}
        _ = v6
    }
    var v7 uint8 = 42
    v8 := GenS2{F0: 2147483648}
    var v9 float32 = float32(1.5)
    var v10 int32 = 1
    v11 := GenS2{F0: 44}
    fmt.Println(v7, v1)
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
