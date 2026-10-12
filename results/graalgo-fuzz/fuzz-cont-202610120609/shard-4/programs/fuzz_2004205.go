package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    v3 := []int{128, 32768, 110, -9223372036854775808}
    for fi4 := 0; fi4 < 18; fi4++ {
        fmt.Println(fi4)
        _ = len(v3)
        var v5 uint64 = 1
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 225, F1: 0}
    v7 := GenS1{F0: 0, F1: 32767}
    var v8 int16 = -1
    v1[("foo bar" + "\n\t")] = (-109 > -62)
    if true {
        v3 = append(v3, -1)
    }
    var v9 uint8 = 2
    v10 := GenS1{F0: v9, F1: 571}
    if true {
        v11 := GenS1{F0: v9, F1: 255}
        _ = v11
    }
    if false {
        fmt.Println(v8, v8)
        v8 = v8
        fmt.Println(v8, v8)
    } else {
        _ = len(v3)
    }
    v9 += v9
    var v12 uint32 = 224
    v9 = v9
    var v13 float32 = 0.1
    v14 := GenS1{F0: 10, F1: 652}
    v15 := GenS1{F0: v9, F1: 475}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
