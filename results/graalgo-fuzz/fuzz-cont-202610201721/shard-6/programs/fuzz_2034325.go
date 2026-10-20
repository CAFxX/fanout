package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int) uint {
    var v7 float32 = ((-2.25 + (-2.25 - float32(0.1))) * 3.14159)
    var v8 int8 = 127
    v9 := []int{}
    v10 := GenS1{F0: 50}
    v9 = append(v9, p0)
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 uint = 4294967295
    var v5 uint = (v4 + 42)
    var v6 uint32 = v1
    var v11 int64 = 368
    if false {
        var v12 int = 42
        var v13 float32 = 0.1
        _ = v12
        _ = v13
    }
    var v14 uint64 = 709
    v15 := []int{42, 42}
    v16 := GenS1{F0: 2147483647}
    var v17 int64 = v11
    fmt.Println(v4, v5)
    var v18 float64 = 3.14159
    if (1 > 0) {
        v1 = 0
        var v19 int = -631
        v15 = append(v15, (-(v19)))
        _ = v19
    } else {
        v4 += (v5 << ((684) & 63))
    }
    v18 += (0.1 / (v18 / (v18 * 1e10)))
    v20 := GenS1{F0: v1}
    for fi21 := 0; fi21 < 2; fi21++ {
        v17 += v17
        v4 = 2
        delete(v3, ("a" + "a"))
        _ = fi21
    }
    var v22 uint64 = (v14 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
