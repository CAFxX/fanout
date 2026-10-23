package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 bool = false
    var v2 int = -515
    var v3 int32 = -1
    var v4 float64 = 3.14159
    var v5 uint = 710
    v5 = v5
    v4 = v4
    fmt.Println(v4, v1)
    v6 := GenS1{F0: 1, F1: 4294967295}
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 string = "hello"
        v4 += (1e10 - (v4 - ((v4 / v4) + (v4 / v4))))
        fmt.Println(v4)
        _ = fi7
        _ = v8
    }
    var v9 uint = 4294967295
    if (32767 <= v2) {
        v10 := []int{(v2 % (((-(v2))) | 1)), v2, v2, (-(v2))}
        _ = v10
    }
    fmt.Println(v4)
    v11 := GenS1{F0: 666, F1: 4294967295}
    v2 -= 673
    if (!(2 > 42)) {
        v12 := []int{42, -9223372036854775808}
        _ = len(v12)
        _ = v12
    }
    var v13 uint8 = 127
    v14 := []int{(-(v2)), (-(v2)), 0}
    fmt.Println(v13, v13, v13)
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
