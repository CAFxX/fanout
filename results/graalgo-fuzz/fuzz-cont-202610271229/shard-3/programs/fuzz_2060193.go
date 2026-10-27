package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 bool = (600 != 0)
    var v3 uint8 = 128
    v4 := make(map[int]int)
    v5 := make(map[int]string)
    var v6 int16 = 42
    var v7 uint = 2
    if v2 {
        v8 := GenS1{F0: 4}
        _ = v8
    }
    fmt.Println(v3, v6, v2)
    v9 := []int{-156, 309}
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS1{F0: 500}
        _ = fi10
        _ = v11
    }
    if v2 {
        if v2 {
            switch "\n\t" {
            case "hello":
                var v12 int16 = v6
                _ = v12
            }
            fmt.Println(v7)
        }
        fmt.Println(v6, v3, v6)
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int = fi13
        v6 += v6
        v15 := GenS1{F0: 1}
        _ = fi13
        _ = v14
        _ = v15
    }
    if (v3 < (v3 % ((42) | 1))) {
        if len(v9) > 0 {
            v9[0] = 352
        }
        var v16 uint32 = 65535
        _ = v16
    } else {
        if len(v9) > 0 {
            v9[0] = 1
        }
        v3 -= (^(v3))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
