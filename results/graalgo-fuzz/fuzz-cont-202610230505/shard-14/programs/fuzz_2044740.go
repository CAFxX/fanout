package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint, p1 uint32) bool {
    p1 += p1
    fmt.Println(p0, p1)
    p1 += (p1 % ((p1) | 1))
    v7 := []int{}
    v8 := GenS1{F0: 727, F1: 704}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 int8 = -128
    var v2 int = -9223372036854775808
    var v3 uint16 = 42
    var v4 int8 = -48
    var v5 int = 84
    var v6 int16 = 0
    if false {
        var v9 float32 = -2.25
        _ = v9
    } else {
        v10 := GenS1{F0: 18446744073709551615, F1: v2}
        v11 := GenS1{F0: 4294967295, F1: v5}
        _ = v10
        _ = v11
    }
    var v12 bool = false
    var v13 float64 = float64(-2.25)
    v12 = v12
    v14 := GenS2{F0: -1, F1: v3}
    v5 = v5
    var v15 int64 = 9223372036854775807
    fmt.Println(v13, v13, v1)
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 uint16 = (v3 >> ((147) & 63))
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
