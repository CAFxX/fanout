package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint = 642
    var v2 uint32 = 1
    v3 := make(map[int]bool)
    var v4 uint32 = (v2 | (v2 / ((2147483647) | 1)))
    v5 := make(map[int]int)
    var v6 uint8 = 0
    var v7 bool = (165 > 2)
    delete(v5, 337)
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int8 = 64
        var v10 bool = v7
        fmt.Println(v2, v6, v10)
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = len(v3)
    v7 = v7
    v6 = 2
    switch 963 {
    case 4294967295:
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(v6, v4)
            _ = fi11
        }
    }
    var v12 uint32 = v2
    fmt.Println(v4, v2)
    v13 := GenS1{F0: 42}
    var v14 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
