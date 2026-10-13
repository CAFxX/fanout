package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint16 = 405
    var v2 int32 = -2147483648
    var v3 bool = (((694 == 690) && false) && false)
    v4 := make(map[int]bool)
    var v5 int8 = -1
    v6 := make(map[string]int)
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: 0}
        if (!false) {
            var v9 uint = 222
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    var v10 int = 722
    fmt.Println(v1, v1)
    var v11 uint = 0
    var v12 int8 = (v5 + 0)
    v10 += v10
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := GenS1{F0: 99}
        _ = fi13
        _ = v14
    }
    v12 += -128
    v15 := []int{-1, (^(v10))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
