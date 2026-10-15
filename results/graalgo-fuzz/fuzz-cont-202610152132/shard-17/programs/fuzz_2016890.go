package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int8 }

func main() {
    var v1 int64 = 1
    v2 := make(map[int]uint32)
    var v3 float32 = 1e10
    var v4 uint16 = 2
    if true {
        v4 = 65535
        v4 = 837
    } else {
        var v5 uint16 = (^(v4))
        _ = v5
    }
    var v6 uint64 = 4294967295
    fmt.Println(v4, v6)
    for fi7 := 0; fi7 < 2; fi7++ {
        v6 += (^(v6))
        v8 := GenS1{F0: 0, F1: "hello", F2: 103}
        if true {
            v9 := []int{fi7, (fi7 << ((354) & 63)), fi7, (-(fi7))}
            var v10 int64 = 498
            _ = v10
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    if false {
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := v2[int(v4)]
            fmt.Println(v1)
            _ = fi11
            _ = v12
        }
        fmt.Println(v6, v4, v6)
    } else {
        v13 := []int{32767, -828, 32768, 32768}
        _ = v13
    }
    fmt.Println(v6, v1, v4)
    v1 += (v1 * v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
