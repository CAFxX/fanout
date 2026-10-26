package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 1.5
    v3 := make(map[uint32]string)
    var v4 int8 = -1
    for fi5 := 0; fi5 < 9; fi5++ {
        fmt.Println(fi5, v2)
        v4 = (v4 | (v4 + (v4 / (((v4 << ((366) & 63))) | 1))))
        _ = fi5
    }
    v4 += v4
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := v3[655]
        fmt.Println(v2, v2, fi6)
        _ = fi6
        _ = v7
    }
    fmt.Println(v4, v4)
    v8 := GenS1{F0: 255}
    v2 = v2
    var v9 int32 = -1
    v4 += (v4 << ((0) & 63))
    var v10 string = ""
    v10 = ((v10 + (((v10 + "hello") + v10) + ("a" + ""))) + v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
