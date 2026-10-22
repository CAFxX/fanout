package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint; F1 int; F2 int64 }
type GenS3 struct { F0 string; F1 uint; F2 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (0 <= 42)
    var v3 string = "a"
    v4 := make(map[uint32]uint32)
    v5 := GenS1{F0: 1, F1: 42}
    v3 = v3
    switch 128 {
    case 32768:
        fmt.Println(v2, v2)
    case 9223372036854775807:
        v2 = (204 < 2)
    case 128:
        var v6 uint64 = 1
        _ = v6
    }
    switch (v3 + "foo bar") {
    case "foo bar":
        var v7 bool = true
        _ = v7
    case "a":
        var v8 uint = 1
        _ = v8
    default:
        v9 := GenS2{F0: 4294967295, F1: 32768, F2: 9223372036854775807}
        _ = v9
    }
    v10 := []int{}
    var v11 uint16 = 32767
    var v12 uint8 = 0
    v13 := GenS3{F0: "foo bar", F1: 540, F2: -1}
    fmt.Println(v2, v11)
    v4[246] = uint32(v12)
    v3 = (v3 + "a")
    fmt.Println(v3, v11, v11)
    fmt.Println(v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
