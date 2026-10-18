package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[string]int)
    var v3 string = "a"
    var v4 int16 = 282
    v5 := []int{1, 401, -2147483648, -661}
    var v6 string = "hello"
    var v7 int = 4294967295
    v8 := GenS1{F0: 0}
    switch v7 {
    case 509:
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 uint32 = 596
            var v11 bool = (1 < 1)
            var v12 int16 = (^(v4))
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    case -641:
        v3 = ""
    case 0:
        v13 := GenS1{F0: -120}
        _ = v13
    default:
        var v14 float64 = 1e10
        _ = v14
    }
    var v15 uint = (^(v1))
    var v16 int16 = v4
    if (-2147483648 != 597) {
        var v17 uint64 = 225
        fmt.Println(v1)
        var v18 int16 = 42
        _ = v17
        _ = v18
    } else {
        delete(v2, v6)
        var v19 bool = (((2 < 990) && false) || (42 != 86))
        v16 = v4
        _ = v19
    }
    var v20 float64 = 0.0
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
