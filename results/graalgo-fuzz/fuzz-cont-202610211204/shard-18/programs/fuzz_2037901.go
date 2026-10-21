package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint32 }
func genFn0(p0 uint64) int32 {
    p0 = 583
    var v4 int32 = 0
    _ = p0
    _ = v4
    return -49
}

func main() {
    var v1 int = 256
    var v2 uint = 18446744073709551615
    v3 := make(map[uint32]bool)
    var v5 int32 = 2147483647
    var v6 int64 = -1
    v3[42] = (544 <= 495)
    var v7 int64 = (^(v6))
    delete(v3, 42)
    v8 := GenS1{F0: (1 == 18446744073709551615), F1: (^(v2)), F2: 314}
    v1 = v1
    switch "hello" {
    case "foo bar":
        var v9 uint64 = 2
        _ = v9
    case "hello":
        v2 += (v2 % (((v2 % (((v2 ^ v2)) | 1))) | 1))
    default:
        switch 127 {
        case 128:
            fmt.Println(v2)
        case 4294967295:
            v10 := GenS2{F0: -1, F1: 255, F2: 102}
            _ = v10
        default:
            fmt.Println(v2, v6, v2)
        }
    }
    var v11 bool = false
    fmt.Println(v6, v5)
    v12 := []int{v1}
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
