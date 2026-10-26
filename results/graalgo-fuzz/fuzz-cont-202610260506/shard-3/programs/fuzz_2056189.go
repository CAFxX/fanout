package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    var v3 bool = (true && (-1 > 89))
    var v4 float32 = (-2.25 * (0.0 + (0.1 * -2.25)))
    v5 := make(map[string]string)
    var v6 int64 = 42
    var v7 int8 = -1
    if (-32768 <= 744) {
        v5["\n\t"] = (("\n\t" + "a") + ("foo bar" + "\n\t"))
        fmt.Println(v6, v4, v6)
    }
    fmt.Println(v7, v4, v6)
    v4 = float32(3.14159)
    var v8 int32 = 30
    v9 := GenS1{F0: 2}
    var v10 uint64 = 237
    var v11 bool = (!v3)
    if (1 >= 707) {
        v10 += (v10 & v10)
        var v12 int64 = (v6 * v6)
        _ = v12
    } else {
        var v13 uint16 = 65535
        _ = v13
    }
    fmt.Println(v3, v7, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
