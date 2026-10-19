package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint32 }

func main() {
    v1 := []uint32{1, 65535, 2}
    var v2 uint64 = 4294967295
    var v3 int16 = 886
    var v4 uint16 = 373
    v5 := make(map[string]int)
    var v6 int32 = 409
    if true {
        v7 := GenS1{F0: 1, F1: -1, F2: 4294967295}
        fmt.Println(v4, v4)
        _ = len(v1)
        _ = v7
    }
    fmt.Println(v2)
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int32 = (-(v6))
        _ = fi8
        _ = v9
    }
    var v10 int32 = (v6 + (v6 % ((-2147483648) | 1)))
    if ((v10 > v6) && (!(2 < 65535))) {
        for fi11 := 0; fi11 < 15; fi11++ {
            for fi12 := 0; fi12 < 14; fi12++ {
                v4 += (v4 % ((v4) | 1))
                fmt.Println(fi11)
                _ = fi12
            }
            _ = fi11
        }
    } else {
        v2 = 0
    }
    v6 = (v6 ^ 0)
    var v13 string = "a"
    fmt.Println(v13)
    v10 += (v10 * v6)
    v14 := GenS1{F0: -128, F1: 139, F2: 65535}
    v13 = v13
    if (42 != 166) {
        var v15 int32 = -2147483648
        v10 = 42
        switch v13 {
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = v15
    } else {
        fmt.Println(v6)
        var v16 uint64 = 42
        var v17 uint = 260
        _ = v16
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
