package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int32) uint16 {
    for fi7 := 0; fi7 < 13; fi7++ {
        p0 = p0
        v8 := GenS1{F0: 4294967295}
        _ = fi7
        _ = v8
    }
    var v9 int64 = -9223372036854775808
    v10 := []int{127, -619, 496, 4294967295}
    if true {
        var v11 int16 = 738
        _ = len(v10)
        if len(v10) > 0 {
            v10[0] = -1
        }
        _ = v11
    }
    _ = p0
    _ = v10
    _ = v9
    return 255
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 0
    var v3 uint = 519
    v4 := make(map[string]bool)
    var v5 uint = v3
    v6 := []string{}
    v12 := v4["foo bar"]
    var v13 int8 = -1
    v5 = (v3 % ((v3) | 1))
    v14 := GenS1{F0: 0}
    var v15 float64 = 0.0
    if true {
        fmt.Println(v13)
    }
    var v16 uint32 = 0
    fmt.Println(v16, v2, v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
