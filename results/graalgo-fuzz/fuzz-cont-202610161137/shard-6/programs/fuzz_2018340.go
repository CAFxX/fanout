package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int32 }
type GenS2 struct { F0 int8; F1 int64 }

func main() {
    var v1 uint = 65535
    var v2 int32 = -2147483648
    var v3 float32 = 3.14159
    var v4 uint = 18446744073709551615
    var v5 int32 = v2
    v6 := make(map[uint32]uint32)
    var v7 uint64 = 0
    var v8 int64 = 0
    v9 := GenS1{F0: v4, F1: (v1 ^ 4294967295), F2: v2}
    for fi10 := 0; fi10 < 20; fi10++ {
        v4 = (^(v4))
        _ = fi10
    }
    v11 := GenS2{F0: 127, F1: v8}
    var v12 uint16 = 698
    var v13 uint64 = (^(v7))
    v13 = (v7 / ((uint64(v12)) | 1))
    var v14 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
