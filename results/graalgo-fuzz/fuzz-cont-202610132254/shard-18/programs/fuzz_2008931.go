package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int16, p1 int32) bool {
    var v7 uint64 = 360
    v7 -= 1
    var v8 int64 = 252
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return false
}
func genFn1(p0 int8) uint16 {
    fmt.Println(p0)
    if (!true) {
        var v9 int = 32767
        fmt.Println(v9, p0)
        _ = v9
    } else {
        var v10 float64 = (0.0 * 1e10)
        fmt.Println(v10)
        _ = v10
    }
    _ = p0
    return 242
}

func main() {
    var v1 int16 = 0
    var v2 uint16 = 806
    var v3 int32 = 1
    var v4 bool = true
    var v5 string = "foo bar"
    var v6 int16 = (v1 * v1)
    v11 := []int{42, 32768}
    var v12 int8 = -1
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := GenS1{F0: "foo bar"}
        _ = fi13
        _ = v14
    }
    var v15 float64 = 1e10
    var v16 uint = 42
    v5 = v5
    var v17 uint8 = 27
    var v18 string = v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
