package main

import (
    "fmt"
)

func genFn0() uint16 {
    var v7 int = 0
    for fi8 := 0; fi8 < 18; fi8++ {
        v7 -= v7
        _ = fi8
    }
    var v9 int16 = 32767
    var v10 uint32 = 2147483647
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 bool = false
        var v13 uint = 4294967295
        v13 = 18446744073709551615
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = v10
    _ = v7
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{0, 65535, 29}
    v3 := make(map[int]uint32)
    var v4 float64 = ((3.14159 + 1.5) - 0.1)
    var v5 string = "\n\t"
    var v6 int16 = 1000
    v5 = v5
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    if (!(true && false)) {
        var v14 int16 = (v6 * (v6 | -871))
        v14 -= v6
        _ = v14
    }
    v5 = (v5 + (v5 + "\n\t"))
    var v15 uint16 = genFn0()
    v4 -= (-2.25 - ((v4 - float64(float32(-2.25))) - (v4 - v4)))
    switch v5 {
    case "\n\t":
        var v16 int8 = 100
        _ = v16
    default:
        var v17 uint = 18446744073709551615
        _ = v17
    }
    if false {
        var v18 bool = false
        var v19 int16 = (v6 << ((1) & 63))
        v5 = ("hello" + v5)
        _ = v18
        _ = v19
    }
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
