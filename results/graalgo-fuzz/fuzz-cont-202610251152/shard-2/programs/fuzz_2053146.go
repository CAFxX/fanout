package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint64 = 1
    v2 := []uint32{885, 65535, 65535, 4294967295}
    var v3 int32 = -1
    var v4 int64 = 0
    var v5 uint8 = 2
    var v6 float64 = 0.1
    v3 = (v3 % (((v3 << ((0) & 63))) | 1))
    if len(v2) > 0 {
        v2[0] = 1
    }
    v4 += (v4 >> ((18446744073709551615) & 63))
    _ = len(v2)
    fmt.Println(v4)
    v7 := GenS1{F0: 163}
    var v8 uint8 = (v5 / (((v5 + v5)) | 1))
    var v9 float64 = -2.25
    v10 := GenS2{F0: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
