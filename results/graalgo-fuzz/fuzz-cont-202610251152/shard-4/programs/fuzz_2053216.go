package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int }
func genFn0(p0 int, p1 uint32) bool {
    switch int(p1) {
    case 4294967295:
        if false {
            var v5 uint = 4294967295
            _ = v5
        } else {
            v6 := GenS1{F0: -370, F1: (2 == 42), F2: p0}
            _ = v6
        }
    case 127:
        p1 = 2
    }
    var v7 float64 = -2.25
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v7
    return true
}
func genFn1(p0 string, p1 string, p2 uint64) bool {
    v8 := GenS1{F0: -105, F1: (!(32767 > 1000)), F2: 9223372036854775807}
    var v9 float64 = (1e10 * 0.0)
    var v10 int = -947
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 128
    var v3 uint32 = 488
    var v4 uint16 = 531
    fmt.Println(v4)
    var v11 int64 = 1
    var v12 int = 515
    fmt.Println(v2)
    v13 := []int{v12}
    if len(v13) > 0 {
        v13[0] = (-(v12))
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        fmt.Println(fi14, v4)
        var v15 int16 = 1
        _ = fi14
        _ = v15
    }
    v13 = append(v13, -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
