package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() uint32 {
    if false {
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 float32 = (0.0 * 1e10)
            v8 := []int{42, (^(fi6)), fi6}
            _ = fi6
            _ = v7
            _ = v8
        }
        var v9 uint32 = 709
        _ = v9
    } else {
        var v10 uint8 = 2
        _ = v10
    }
    var v11 float32 = (0.1 * 3.14159)
    _ = v11
    return 266
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    var v5 int8 = -115
    var v12 bool = true
    for fi13 := 0; fi13 < 20; fi13++ {
        v1 -= 1.5
        _ = fi13
    }
    v5 += 1
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS1{F0: v12}
        _ = fi14
        _ = v15
    }
    var v16 uint32 = genFn0()
    v17 := v4["\n\t"]
    if false {
        v16 = v17
    }
    var v18 int64 = 42
    v19 := []int{128, 2147483647, int(v18)}
    v5 -= v5
    v17 = (v17 + (v16 ^ v17))
    var v20 int8 = v5
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 int = (fi21 + (fi21 / ((fi21) | 1)))
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
