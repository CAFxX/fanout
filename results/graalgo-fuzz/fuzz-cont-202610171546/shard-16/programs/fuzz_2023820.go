package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int8 = 127
    var v2 string = ""
    var v3 int64 = 959
    var v4 int8 = v1
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: v4}
    var v8 int16 = 1
    v3 += 42
    v1 += v1
    var v9 float32 = (((1e10 + -2.25) / (-2.25 * ((0.0 - (float32(3.14159) / (1e10 - ((3.14159 - 0.1) - 0.1)))) - (-2.25 + 1.5)))) * 1e10)
    v10 := v6[v2]
    var v11 uint16 = 65535
    var v12 int16 = -465
    v11 = (^(v11))
    v10 += v10
    var v13 uint32 = 65535
    v12 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
    _ = v7
    _ = v8
    _ = v9
}
