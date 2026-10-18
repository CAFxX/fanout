package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0() int {
    var v4 uint64 = 0
    v5 := []int{397, 4294967295, 74, 9223372036854775807}
    v6 := GenS1{F0: 0, F1: -1}
    var v7 uint64 = (v4 + v4)
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 635
    var v3 string = "hello"
    v8 := GenS1{F0: -128, F1: -86}
    v3 = v3
    var v9 string = v3
    fmt.Println(v9, v9)
    var v10 uint64 = 366
    v9 = ("" + "\n\t")
    v9 = ("" + v9)
    fmt.Println(v3)
    v11 := GenS1{F0: 127, F1: -128}
    for fi12 := 0; fi12 < 4; fi12++ {
        fmt.Println(fi12, v3)
        v13 := GenS2{F0: -1, F1: 100}
        _ = fi12
        _ = v13
    }
    v14 := []int{829, 879, 128, 0}
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 uint32 = 1
        v16 = v16
        v17 := GenS1{F0: -82, F1: 108}
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
