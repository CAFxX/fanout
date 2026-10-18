package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint }

func main() {
    var v1 uint64 = 536
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    v4 := []uint32{}
    var v5 int = 546
    var v6 int32 = -2147483648
    v5 -= (v5 & v5)
    v4 = append(v4, 124)
    switch v5 {
    case 1:
        v2[("" + "foo bar")] = -432
    case -772:
        v5 -= v5
    default:
        switch (v5 ^ v5) {
        case -1:
            v1 = v1
        case -2147483648:
            var v7 uint = 65535
            _ = v7
        case -998:
            var v8 uint32 = 2147483647
            _ = v8
        default:
            var v9 int = (v5 * 42)
            _ = v9
        }
    }
    fmt.Println(v5, v6, v6)
    var v10 int = (v5 << ((504) & 63))
    v1 += v1
    fmt.Println(v6)
    var v11 string = "hello"
    var v12 bool = (float64(0.0) == float64(-2.25))
    delete(v2, (("hello" + v11) + "a"))
    var v13 float64 = 1e10
    for fi14 := 0; fi14 < 20; fi14++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi14
    }
    if true {
        switch v5 {
        case -9223372036854775808:
            for fi15 := 0; fi15 < 12; fi15++ {
                v10 = (v5 % (((fi15 * (v5 | fi15))) | 1))
                _ = fi15
            }
        case 4294967295:
            fmt.Println(v6)
        default:
            v16 := GenS1{F0: int(v6), F1: (v1 ^ (v1 & (v1 & 0))), F2: 255}
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 int64 = 42
            fmt.Println(v11)
            _ = fi17
            _ = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
