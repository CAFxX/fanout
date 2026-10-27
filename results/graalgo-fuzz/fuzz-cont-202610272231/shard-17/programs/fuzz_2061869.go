package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int8; F2 int64 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 42
    v3 := make(map[uint32]string)
    var v4 uint16 = 192
    var v5 uint = 2
    var v6 string = "a"
    var v7 uint64 = 0
    v4 = (v4 ^ v4)
    v8 := GenS1{F0: uint(v4)}
    v9 := GenS2{F0: 1, F1: 0, F2: -9223372036854775808}
    v10 := GenS3{F0: 573}
    fmt.Println(v5)
    v2 = v2
    var v11 uint8 = 0
    switch 216 {
    case 0:
        v2 = (v2 ^ -224)
    case 32768:
        v2 += v2
    }
    v11 = v11
    var v12 uint = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
