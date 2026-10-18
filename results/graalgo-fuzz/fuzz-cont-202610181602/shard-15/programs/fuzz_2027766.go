package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int32; F2 int64 }

func main() {
    var v1 bool = (-99 > 100)
    v2 := make(map[int]int)
    var v3 uint16 = 527
    var v4 int64 = -719
    v5 := []string{"", "foo bar"}
    var v6 uint32 = 2147483648
    v7 := GenS1{F0: 42, F1: 42}
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(v6, fi8)
        fmt.Println(v3, v4)
        for fi9 := 0; fi9 < 1; fi9++ {
            fmt.Println(fi8)
            fmt.Println(v4)
            if v1 {
                var v10 uint = 566
                _ = v10
            }
            _ = fi9
        }
        _ = fi8
    }
    v4 -= 42
    var v11 uint8 = 42
    _ = len(v5)
    var v12 int64 = v4
    for fi13 := 0; fi13 < 1; fi13++ {
        switch "a" {
        case "\n\t":
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := GenS2{F0: v6, F1: -522, F2: -1}
                v11 -= v11
                _ = fi14
                _ = v15
            }
        case "foo bar":
            v16 := GenS2{F0: (v6 % ((500) | 1)), F1: 42, F2: -1}
            _ = v16
        default:
            v17 := v2[1]
            _ = v17
        }
        fmt.Println(v3)
        v11 += (^(v11))
        _ = fi13
    }
    v18 := GenS1{F0: 278, F1: 2}
    v11 -= v11
    v6 += v6
    v19 := GenS1{F0: 4294967295, F1: 42}
    delete(v2, 4294967295)
    v20 := GenS2{F0: v6, F1: 2147483647, F2: (v4 - (-(v4)))}
    fmt.Println(v4, v4, v1)
    v4 = v12
    switch 128 {
    case -690:
        fmt.Println(v4)
    case 9223372036854775807:
        v21 := v2[4294967295]
        _ = v21
    case 955:
        v11 += 59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
