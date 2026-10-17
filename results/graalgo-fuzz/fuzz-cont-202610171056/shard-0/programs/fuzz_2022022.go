package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int16) int32 {
    fmt.Println(p0)
    v4 := GenS1{F0: -114}
    _ = p0
    _ = v4
    return int32(p0)
}

func main() {
    var v1 int32 = -1
    v2 := []int{-625, -1}
    var v3 float64 = 3.14159
    v5 := GenS1{F0: int8(v1)}
    var v6 int64 = -9223372036854775808
    v7 := GenS1{F0: -128}
    var v8 int64 = (v6 | 42)
    fmt.Println(v6, v8, v6)
    var v9 bool = false
    switch 4294967295 {
    case -9223372036854775808:
        var v10 bool = (("hello" + "\n\t") != "a")
        _ = v10
    case 28:
        var v11 float64 = v3
        _ = v11
    }
    fmt.Println(v6, v6)
    if false {
        var v12 int16 = -727
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        fmt.Println(v6)
        v2 = append(v2, fi13)
        _ = fi13
    }
    var v14 uint = 1
    v14 += (^(v14))
    v15 := GenS1{F0: int8(v14)}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
