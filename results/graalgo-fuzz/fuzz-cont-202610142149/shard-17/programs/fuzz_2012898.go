package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    var v4 int64 = 9223372036854775807
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 uint8 = 42
        _ = fi5
        _ = v6
    }
    v4 += v4
    v4 += v4
    var v7 uint32 = 2147483648
    delete(v1, 42)
    v8 := GenS1{F0: 76, F1: 2147483647, F2: -1}
    fmt.Println(v7, v7)
    v4 = v4
    v4 = 691
    fmt.Println(v4)
    v7 = v7
    v1[2147483647] = (42 > 2147483647)
    v7 += v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
