package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint8 = 94
    var v2 bool = false
    var v3 int16 = 0
    v4 := []int{32768, -9223372036854775808}
    v1 -= (v1 >> ((326) & 63))
    v4 = append(v4, -1)
    v5 := GenS1{F0: "hello"}
    if (80 < 42) {
        fmt.Println(v1, v2)
    }
    var v6 uint64 = 964
    v2 = (!false)
    fmt.Println(v1, v1)
    var v7 int64 = -9223372036854775808
    switch -1 {
    case 9223372036854775807:
        v6 = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
