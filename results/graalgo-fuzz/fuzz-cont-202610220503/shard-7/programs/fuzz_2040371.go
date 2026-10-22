package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 string }
func genFn0() int32 {
    var v4 float32 = (1e10 * (1e10 + 1e10))
    var v5 float64 = 1e10
    _ = v4
    _ = v5
    return -74
}

func main() {
    var v1 uint64 = 562
    var v2 string = "hello"
    var v3 bool = (!(1000 == -32768))
    v3 = v3
    if true {
        var v6 string = ("hello" + "a")
        var v7 uint32 = 65535
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        v2 = v2
        fmt.Println(v3)
        _ = fi8
    }
    fmt.Println(v1)
    v9 := []int{128, 128}
    v10 := GenS1{F0: 100, F1: v1, F2: v2}
    var v11 uint32 = 2
    if len(v9) > 0 {
        v9[0] = 128
    }
    v2 = v2
    v12 := GenS1{F0: 0, F1: 18446744073709551615, F2: ("foo bar" + v2)}
    var v13 int32 = -273
    v3 = true
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
