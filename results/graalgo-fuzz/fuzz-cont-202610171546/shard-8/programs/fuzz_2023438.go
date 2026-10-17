package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
func genFn0(p0 int64) uint64 {
    p0 += p0
    p0 = (p0 ^ (p0 - 9223372036854775807))
    p0 = 0
    v7 := GenS1{F0: (((296 <= 1) && (42 > 42)) || true), F1: 2147483647}
    _ = p0
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint = 1
    var v2 uint8 = 50
    var v3 int8 = 100
    var v4 int = int(v2)
    var v5 int = 42
    var v6 int32 = 67
    var v8 uint = (v1 | v1)
    var v9 int32 = (v6 / ((v6) | 1))
    v10 := []int{v4, int(v8)}
    v11 := GenS1{F0: (!((v4 >> ((2) & 63)) != v4)), F1: 251}
    v3 = v3
    v12 := GenS1{F0: false, F1: 1}
    var v13 int = (v5 % ((int(v3)) | 1))
    v6 = v6
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v4 += (v4 + (v5 / (((v4 ^ (v13 ^ v4))) | 1)))
    v13 = (v4 * (v4 / ((v5) | 1)))
    _ = len(v10)
    var v14 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
