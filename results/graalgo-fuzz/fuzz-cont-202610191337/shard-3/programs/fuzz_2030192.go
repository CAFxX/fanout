package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint8 = 4
    var v2 int64 = 0
    var v3 uint = 42
    var v4 bool = false
    v5 := make(map[uint32]string)
    if v4 {
        v2 = v2
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: "hello"}
        v4 = true
        _ = fi6
        _ = v7
    }
    v2 += (v2 - (v2 - v2))
    fmt.Println(v4, v3, v1)
    v2 = (v2 >> ((42) & 63))
    var v8 uint64 = 2
    var v9 int = 42
    var v10 int64 = v2
    v4 = true
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
