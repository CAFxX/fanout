package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int }

func main() {
    v1 := []string{("a" + "\n\t")}
    var v2 uint32 = 0
    var v3 uint64 = 4294967295
    var v4 int64 = 398
    var v5 int32 = -1
    v6 := make(map[string]uint32)
    var v7 int64 = v4
    for fi8 := 0; fi8 < 20; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    v9 := GenS1{F0: 42, F1: 1, F2: 32767}
    switch "a" {
    case "a":
        v6[""] = v2
    case "":
        v5 = v5
    case "\n\t":
        fmt.Println(v5, v3)
    }
    v2 = (v2 | v2)
    fmt.Println(v3, v2)
    var v10 uint32 = (v2 ^ (v2 * (^(v2))))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v11 int32 = (v5 << ((918) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
