package main

import (
    "fmt"
)

func genFn0() int16 {
    var v7 int16 = -652
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 int = fi8
        fmt.Println(v7, v9, v9)
        _ = fi8
        _ = v9
    }
    _ = v7
    return (v7 * v7)
}

func main() {
    var v1 int = 127
    var v2 int8 = 1
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    v5 := []int{(v1 & (v1 + v1)), v1, v1, (v1 & (^(v1)))}
    v6 := make(map[uint32]bool)
    if false {
        var v10 int = v1
        _ = v10
    }
    var v11 int = 213
    var v12 int64 = -1
    if false {
        v2 = (-(v2))
    }
    var v13 int16 = 0
    fmt.Println(v1, v1, v2)
    v14 := v6[42]
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v13 -= (^(v13))
    for fi15 := 0; fi15 < 13; fi15++ {
        if len(v5) > 0 {
            v5[0] = fi15
        }
        var v16 int32 = -2147483648
        _ = fi15
        _ = v16
    }
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
