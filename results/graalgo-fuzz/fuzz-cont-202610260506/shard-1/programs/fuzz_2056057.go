package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]int)
    v3 := make(map[string]bool)
    var v4 int64 = -1
    v5 := make(map[uint32]int)
    var v6 float32 = v1
    v7 := GenS1{F0: 0, F1: 2}
    v8 := GenS1{F0: 376, F1: 592}
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            v1 = v6
            fmt.Println(v1, fi10, fi9)
            _ = fi10
        }
        _ = fi9
    }
    var v11 uint16 = 993
    fmt.Println(v11, v1, v4)
    v12 := []int{}
    if false {
        v13 := v2["a"]
        if ((!(18446744073709551615 != 0)) || false) {
            var v14 float64 = float64(v1)
            fmt.Println(v14, v6)
            fmt.Println(v4, v14, v11)
            _ = v14
        } else {
            _ = len(v12)
        }
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 int8 = 100
            _ = fi15
            _ = v16
        }
        _ = v13
    } else {
        switch 4294967295 {
        case 281:
            var v17 int16 = -153
            _ = v17
        }
    }
    v18 := GenS1{F0: 42, F1: 0}
    var v19 uint16 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
