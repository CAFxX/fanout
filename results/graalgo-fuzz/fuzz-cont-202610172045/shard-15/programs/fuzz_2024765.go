package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint16, p1 uint32, p2 uint8) uint32 {
    switch "hello" {
    case "\n\t":
        v6 := GenS1{F0: -508, F1: 0}
        _ = v6
    case "hello":
        v7 := GenS1{F0: -1, F1: 42}
        _ = v7
    case "":
        p1 = 878
    }
    v8 := GenS2{F0: -1, F1: p1}
    fmt.Println(p1, p0, p2)
    var v9 int8 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return p1
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int = -1
    var v3 uint8 = 127
    var v4 bool = (-84 < 1)
    v5 := []string{"a", "\n\t"}
    v2 = (v2 % (((^(v2))) | 1))
    if v4 {
        v2 += v2
        v2 += v2
    }
    var v10 uint = 127
    v3 -= 252
    switch -9223372036854775808 {
    case -2147483648:
        v1 += (v1 % ((v1) | 1))
    default:
        v11 := GenS3{F0: 255}
        _ = v11
    }
    var v12 int = v2
    fmt.Println(v10)
    v13 := GenS1{F0: -710, F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
