package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 int32 = 997
    var v2 float64 = 0.1
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    var v5 int64 = -1
    var v6 uint64 = 18446744073709551615
    v3[128] = 42
    v7 := GenS1{F0: -458}
    fmt.Println(v5, v1)
    v4[953] = "foo bar"
    var v8 int16 = 824
    var v9 int32 = (^(v1))
    v10 := GenS1{F0: -2147483648}
    v11 := GenS1{F0: -2147483648}
    if (352 != 65535) {
        v9 = int32(v5)
        v5 += v5
        var v12 int16 = v8
        _ = v12
    } else {
        var v13 uint64 = v6
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
