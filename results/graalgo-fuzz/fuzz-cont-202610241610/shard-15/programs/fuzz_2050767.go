package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int16, p1 int8) int16 {
    var v7 int16 = 32767
    for fi8 := 0; fi8 < 14; fi8++ {
        if (-1 != p1) {
            p1 = (^(p1))
        }
        p1 = p1
        _ = fi8
    }
    fmt.Println(p1, p0, p0)
    _ = p0
    _ = p1
    _ = v7
    return v7
}

func main() {
    var v1 bool = false
    var v2 bool = (v1 && (!((v1 && false) || (1 != 0))))
    var v3 uint32 = 600
    var v4 uint32 = (v3 ^ (^(v3)))
    var v5 int = 42
    v6 := make(map[int]int)
    v9 := GenS1{F0: v3}
    for fi10 := 0; fi10 < 12; fi10++ {
        v3 = (v4 % (((v4 >> ((42) & 63))) | 1))
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 int64 = 0
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v3 += v3
    v4 = v3
    for fi13 := 0; fi13 < 12; fi13++ {
        fmt.Println(v2)
        _ = fi13
    }
    fmt.Println(v2, v2, v1)
    v1 = v1
    v14 := GenS1{F0: 1}
    fmt.Println(v2, v5, v5)
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            fmt.Println(v2, v4)
            _ = fi16
        }
        _ = fi15
    }
    v1 = false
    switch (v5 ^ v5) {
    case 42:
        var v17 uint32 = 1
        _ = v17
    default:
        var v18 int64 = 0
        _ = v18
    }
    var v19 float64 = 1e10
    v20 := []int{208}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
