package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint64 }
func genFn0(p0 int, p1 uint8) uint32 {
    v4 := GenS1{F0: 42, F1: 71, F2: 0}
    var v5 string = "hello"
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return uint32(p1)
}

func main() {
    var v1 uint64 = 42
    var v2 string = "foo bar"
    var v3 float64 = (1.5 / 1e10)
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint = 320
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 0
    if true {
        if true {
            var v9 uint16 = 0
            _ = v9
        }
    }
    v10 := []int{-294, -651}
    var v11 bool = false
    v11 = ((-2147483648 >= 0) && (!true))
    fmt.Println(v11)
    if len(v10) > 0 {
        v10[0] = 42
    }
    var v12 uint = 42
    v13 := GenS1{F0: 2147483648, F1: 0, F2: v1}
    fmt.Println(v12)
    if len(v10) > 0 {
        v10[0] = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v8
}
