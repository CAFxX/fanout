package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int = 127
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    var v5 int32 = -970
    var v6 uint64 = 605
    fmt.Println(v5)
    if true {
        fmt.Println(v6, v5, v6)
    }
    var v7 uint32 = 0
    fmt.Println(v6, v7)
    v8 := GenS1{F0: -1}
    v5 = v5
    var v9 uint64 = (^(v6))
    var v10 int = v1
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := GenS1{F0: 0}
        var v13 uint32 = (v7 / (((v7 / ((v7) | 1))) | 1))
        var v14 uint8 = 0
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 int = v10
    v10 -= -678
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
