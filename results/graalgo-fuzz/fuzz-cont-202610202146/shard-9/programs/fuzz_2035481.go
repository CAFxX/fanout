package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{}
    var v3 bool = (false || (!false))
    v4 := make(map[uint32]bool)
    var v5 bool = (v3 && (!v3))
    var v6 int = -246
    v7 := GenS1{F0: false, F1: 8, F2: 9223372036854775807}
    var v8 int64 = -1
    v4[2] = (v3 || true)
    v2 = append(v2, 65535)
    var v9 int8 = 0
    v9 += v9
    var v10 int8 = -99
    var v11 uint32 = 2147483647
    var v12 bool = v3
    for fi13 := 0; fi13 < 8; fi13++ {
        v6 = 692
        var v14 int = -9223372036854775808
        if (!(!(147 != 0))) {
            v6 += fi13
        } else {
            fmt.Println(v12)
            var v15 uint16 = 584
            v3 = v5
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v3, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
