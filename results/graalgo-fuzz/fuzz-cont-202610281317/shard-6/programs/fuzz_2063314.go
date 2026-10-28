package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint64 = 0
    var v2 bool = false
    v3 := []uint32{42, 0, 2147483647, 128}
    v4 := make(map[int]int)
    var v5 int8 = 12
    var v6 string = "\n\t"
    var v7 uint16 = 0
    for fi8 := 0; fi8 < 11; fi8++ {
        v6 = "foo bar"
        _ = fi8
    }
    fmt.Println(v6, v6, v6)
    v9 := GenS1{F0: "hello"}
    var v10 int32 = -736
    v5 = v5
    v7 -= 236
    v5 += (v5 | 110)
    v11 := GenS1{F0: (v6 + v6)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
