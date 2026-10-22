package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint32, p1 int) int32 {
    switch -1 {
    case 21:
        fmt.Println(p1, p0)
    case 128:
        v7 := GenS1{F0: 2, F1: 1000}
        _ = v7
    case 621:
        var v8 uint16 = 621
        _ = v8
    }
    var v9 int16 = 0
    _ = p0
    _ = p1
    _ = v9
    return 1
}

func main() {
    var v1 int16 = 616
    var v2 int32 = 1
    v3 := make(map[string]int)
    var v4 float32 = ((0.0 + 1.5) + -2.25)
    v5 := make(map[int]string)
    var v6 uint16 = 2
    var v10 bool = false
    var v11 float32 = v4
    var v12 float32 = v4
    v3["\n\t"] = -1
    fmt.Println(v12)
    v12 = v4
    v10 = v10
    v10 = (-123 == 297)
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := GenS1{F0: 665, F1: v1}
        fmt.Println(v4)
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
