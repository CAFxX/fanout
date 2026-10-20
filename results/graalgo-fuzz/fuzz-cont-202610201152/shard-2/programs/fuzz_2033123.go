package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint16, p1 uint64, p2 int16) uint {
    p2 -= p2
    for fi4 := 0; fi4 < 10; fi4++ {
        v5 := GenS1{F0: (true || ((^(fi4)) < -9223372036854775808))}
        _ = fi4
        _ = v5
    }
    var v6 uint32 = 493
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 402
}

func main() {
    var v1 int8 = 42
    var v2 uint32 = 42
    var v3 float64 = float64(3.14159)
    var v7 string = ("\n\t" + "a")
    v8 := []int{128, 4294967295, 9223372036854775807}
    if false {
        for fi9 := 0; fi9 < 8; fi9++ {
            v8 = append(v8, fi9)
            var v10 uint32 = (v2 >> ((4294967295) & 63))
            _ = fi9
            _ = v10
        }
        var v11 uint64 = 0
        v11 = 169
        _ = v11
    }
    var v12 float32 = 1e10
    v3 = 0.0
    _ = len(v8)
    var v13 uint = genFn0(260, 42, 1)
    v7 = v7
    if (v3 >= 0.0) {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        var v14 float64 = v3
        v12 = 0.1
        _ = v14
    }
    fmt.Println(v2)
    var v15 uint32 = 136
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
