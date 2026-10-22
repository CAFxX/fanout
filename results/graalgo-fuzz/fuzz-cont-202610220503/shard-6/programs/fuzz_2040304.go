package main

import (
    "fmt"
)

func genFn0(p0 uint64, p1 bool) int16 {
    var v6 uint32 = 386
    v7 := []int{-1, -2147483648, 128}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 float64 = ((3.14159 - (0.0 - (((0.1 * 3.14159) + 1e10) / 1.0))) + 1e10)
    v2 := []uint32{0, 0, 234}
    v3 := make(map[string]int)
    v4 := []uint32{}
    var v5 uint = 4294967295
    v5 += (v5 | v5)
    fmt.Println(v5, v5)
    switch "hello" {
    case "a":
        v8 := v3["a"]
        _ = v8
    case "foo bar":
        var v9 int32 = 2147483647
        _ = v9
    default:
        v1 = v1
    }
    v5 -= 1
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(v5, v5)
        for fi11 := 0; fi11 < 4; fi11++ {
            _ = len(v3)
            fmt.Println(v5, fi11, fi10)
            var v12 int32 = -941
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    if false {
        var v13 bool = false
        var v14 string = "\n\t"
        _ = v13
        _ = v14
    }
    var v15 uint16 = 101
    var v16 float32 = 1e10
    var v17 uint32 = 65535
    v18 := v3["\n\t"]
    var v19 uint = (v5 + 65535)
    var v20 uint = (v5 << ((v5) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
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
