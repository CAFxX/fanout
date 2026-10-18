package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint32 }

func main() {
    var v1 bool = false
    var v2 uint16 = 0
    var v3 uint64 = 42
    v3 -= 38
    v4 := GenS1{F0: 0, F1: 589, F2: 2147483648}
    v5 := GenS1{F0: 1, F1: 754, F2: 0}
    v6 := []int{}
    for fi7 := 0; fi7 < 13; fi7++ {
        v6 = append(v6, 551)
        _ = len(v6)
        var v8 string = "hello"
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 int64 = 0
        v11 := GenS1{F0: 273, F1: 65535, F2: 4294967295}
        _ = fi9
        _ = v10
        _ = v11
    }
    v3 -= v3
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
