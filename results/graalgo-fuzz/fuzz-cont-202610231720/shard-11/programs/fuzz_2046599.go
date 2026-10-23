package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []uint32{}
    var v2 int32 = 0
    var v3 int32 = (v2 >> ((65535) & 63))
    v3 += (v2 | v3)
    for fi4 := 0; fi4 < 19; fi4++ {
        _ = len(v1)
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        v5 := GenS1{F0: 1}
        _ = fi4
        _ = v5
    }
    var v6 int16 = 1
    var v7 uint64 = 1
    var v8 bool = (false && true)
    v6 = v6
    v8 = (829 >= 1)
    var v9 uint64 = 18446744073709551615
    for fi10 := 0; fi10 < 9; fi10++ {
        fmt.Println(v7, v2)
        _ = fi10
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
