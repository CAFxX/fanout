package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int) uint16 {
    var v6 float32 = (3.14159 / 1e10)
    var v7 float32 = (3.14159 / 1.0)
    var v8 uint16 = 42
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint32 = 1
    var v2 float32 = (-2.25 * 1.5)
    var v3 uint = 332
    var v4 int = -56
    var v5 int32 = 789
    v9 := GenS1{F0: false}
    v5 = v5
    v10 := []int{}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v2 += (0.0 - float32(((1.5 / 1.0) * 0.1)))
    switch 42 {
    case 32767:
        v5 = (-(v5))
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        switch (^(fi11)) {
        case 127:
            v1 = v1
        case 668:
            var v12 int64 = 0
            _ = v12
        default:
            v2 -= v2
        }
        v13 := GenS1{F0: (58 < uint64(v3))}
        _ = fi11
        _ = v13
    }
    var v14 int64 = -1
    v1 += v1
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 string = "a"
        _ = fi15
        _ = v16
    }
    fmt.Println(v1, v14)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
