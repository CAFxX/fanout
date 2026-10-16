package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 string }

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    var v4 uint64 = 42
    var v5 uint16 = 2
    var v6 uint8 = 31
    v6 = v6
    if (false || (2147483647 != 42)) {
        delete(v2, 42)
    } else {
        var v7 int = 154
        _ = v7
    }
    var v8 int8 = int8(v4)
    v9 := GenS1{F0: uint32(v6), F1: v5}
    fmt.Println(v5, v6)
    v10 := GenS2{F0: "\n\t"}
    fmt.Println(v5)
    v11 := GenS2{F0: ("" + ("foo bar" + ""))}
    v12 := GenS1{F0: 65535, F1: 2}
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := GenS2{F0: "\n\t"}
        _ = fi13
        _ = v14
    }
    var v15 uint = 794
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
