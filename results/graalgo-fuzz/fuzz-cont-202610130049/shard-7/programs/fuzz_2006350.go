package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    var v3 int64 = 1
    switch -811 {
    case -2147483648:
        v4 := GenS1{F0: 200}
        _ = v4
    case -764:
        var v5 bool = true
        _ = v5
    }
    var v6 bool = true
    v7 := []int{0}
    for fi8 := 0; fi8 < 6; fi8++ {
        if false {
            var v9 uint64 = 2
            for fi10 := 0; fi10 < 11; fi10++ {
                fmt.Println(v9)
                _ = fi10
            }
            _ = v9
        }
        var v11 float32 = 1e10
        _ = fi8
        _ = v11
    }
    var v12 uint16 = 0
    v13 := GenS1{F0: 128}
    var v14 uint32 = 65535
    v6 = (0 <= 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
