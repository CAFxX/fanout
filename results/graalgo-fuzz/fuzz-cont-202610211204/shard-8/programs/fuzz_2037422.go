package main

import (
    "fmt"
)

func genFn0() uint8 {
    var v4 uint8 = 127
    v5 := []int{-2147483648, 2147483647}
    var v6 float32 = 1.5
    var v7 uint64 = 1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return v4
}

func main() {
    var v1 float64 = 1.5
    var v2 uint32 = 0
    var v3 int = -262
    var v8 int64 = 1
    v8 -= 0
    for fi9 := 0; fi9 < 13; fi9++ {
        fmt.Println(v3, v1)
        var v10 int = (^(v3))
        _ = fi9
        _ = v10
    }
    var v11 float32 = 1.5
    var v12 int = v3
    var v13 float32 = v11
    var v14 string = ("a" + "a")
    var v15 int8 = -128
    fmt.Println(v2, v1)
    v16 := []int{int(v8), (v3 * v3)}
    var v17 int8 = (v15 / (((v15 << ((1) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v8
}
