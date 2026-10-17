package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint8 }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{2147483648, 0, 2}
    var v3 int32 = 42
    v4 := []string{"\n\t", "hello"}
    var v5 int16 = 1000
    v6 := make(map[string]uint32)
    fmt.Println(v5, v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v7 int = 9223372036854775807
    var v8 int8 = 127
    v9 := GenS1{F0: 255}
    v2 = append(v2, 65535)
    var v10 bool = (4294967295 <= 4294967295)
    v10 = v10
    switch int(v3) {
    case 128:
        v11 := GenS2{F0: 1, F1: 128}
        _ = v11
    case -268:
        v8 -= -128
    case -1:
        var v12 int32 = (v3 * (v3 / ((v3) | 1)))
        _ = v12
    default:
        fmt.Println(v10)
    }
    var v13 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
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
