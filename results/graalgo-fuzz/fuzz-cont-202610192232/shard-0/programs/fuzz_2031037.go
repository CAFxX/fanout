package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int32 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{}
    var v3 int64 = 9223372036854775807
    v4 := make(map[uint32]int)
    var v5 bool = (13 < 0)
    var v6 uint8 = 128
    fmt.Println(v3, v6)
    var v7 int16 = 1
    v8 := GenS1{F0: 2, F1: 2, F2: -592}
    var v9 uint64 = 0
    v3 -= v3
    for fi10 := 0; fi10 < 5; fi10++ {
        if false {
            fmt.Println(fi10, v7)
        }
        _ = fi10
    }
    v11 := GenS1{F0: v9, F1: 65535, F2: 1}
    switch -772 {
    case -1:
        v7 = v7
    case 4294967295:
        switch "" {
        case "foo bar":
            var v12 bool = (784 < 333)
            _ = v12
        }
    default:
        v6 -= v6
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
