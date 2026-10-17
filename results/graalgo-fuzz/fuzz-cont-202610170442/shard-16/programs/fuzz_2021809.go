package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 string }

func main() {
    var v1 float32 = ((0.1 + 3.14159) / 0.1)
    var v2 uint32 = 65535
    var v3 int8 = 0
    v4 := make(map[int]uint32)
    var v5 int64 = 789
    var v6 uint32 = v2
    v1 += (1e10 + 3.14159)
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            fmt.Println(v3, v2, v3)
            _ = fi8
        }
        v3 = 127
        _ = fi7
    }
    v9 := []int{32767, 2147483647, 590}
    for fi10 := 0; fi10 < 9; fi10++ {
        v5 += v5
        v5 = 42
        _ = fi10
    }
    var v11 int8 = -91
    switch -1 {
    case 0:
        v11 -= (v11 ^ (-(v3)))
    case 4294967295:
        fmt.Println(v1, v2, v3)
    }
    var v12 uint64 = 4294967295
    v13 := GenS1{F0: true, F1: (v2 / ((1) | 1)), F2: "foo bar"}
    _ = len(v9)
    _ = len(v9)
    switch 128 {
    case 0:
        v2 += (v2 << ((42) & 63))
    case 32768:
        var v14 uint32 = (v2 & v2)
        _ = v14
    }
    var v15 uint64 = v12
    var v16 float32 = (v1 * v1)
    switch 544 {
    case 2147483647:
        if false {
            if (!true) {
                if len(v9) > 0 {
                    v9[0] = -581
                }
            }
        }
    case -9223372036854775808:
        var v17 int8 = v11
        _ = v17
    case -8:
        v18 := GenS1{F0: (913 < 0), F1: v2, F2: "hello"}
        _ = v18
    }
    v19 := GenS1{F0: true, F1: (v2 >> ((0) & 63)), F2: (("a" + "a") + "a")}
    var v20 uint = 498
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
