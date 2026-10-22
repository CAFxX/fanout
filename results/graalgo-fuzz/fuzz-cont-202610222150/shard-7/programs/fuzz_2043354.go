package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 32768
    v3 := make(map[string]bool)
    v4 := make(map[int]uint32)
    v5 := []uint32{65535, 0, 1}
    v6 := GenS1{F0: 0}
    fmt.Println(v2)
    switch 32767 {
    case 32768:
        v7 := GenS1{F0: 42}
        _ = v7
    case -638:
        v2 += v2
    }
    var v8 int8 = -25
    v9 := GenS1{F0: 1}
    fmt.Println(v2, v8)
    delete(v4, v2)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v10 float32 = -2.25
    var v11 int64 = 1
    switch "" {
    case "foo bar":
        var v12 int64 = v11
        _ = v12
    }
    fmt.Println(v10)
    var v13 int = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
