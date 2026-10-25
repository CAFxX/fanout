package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int16 }
func genFn0() uint {
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 float64 = 1.5
        _ = fi5
        _ = v6
    }
    var v7 uint = 631
    v7 -= (^(v7))
    _ = v7
    return 2
}

func main() {
    var v1 string = "hello"
    var v2 uint64 = 4294967295
    var v3 bool = (true || ((-2.25 >= 1.5) || (false || true)))
    v4 := []string{v1, (((v1 + "a") + v1) + v1), v1}
    fmt.Println(v1, v3, v3)
    v8 := GenS1{F0: v3, F1: 466, F2: 0}
    fmt.Println(v1, v3, v2)
    fmt.Println(v2, v1)
    v9 := GenS1{F0: false, F1: 42, F2: 1000}
    var v10 int64 = 0
    v2 += v2
    for fi11 := 0; fi11 < 12; fi11++ {
        v10 += v10
        _ = len(v4)
        var v12 int32 = 2147483647
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        v4 = append(v4, "hello")
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
