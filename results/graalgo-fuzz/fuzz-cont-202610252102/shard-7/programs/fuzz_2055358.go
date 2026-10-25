package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    v1 = (v1 * v1)
    v4 := []int{4294967295}
    v5 := GenS1{F0: 42, F1: 127}
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(fi6, v1)
        v1 -= (v1 << ((255) & 63))
        _ = fi6
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 989, F1: 100}
        fmt.Println(v1)
        _ = fi7
        _ = v8
    }
    var v9 bool = (!(true && true))
    var v10 string = "\n\t"
    var v11 int8 = 127
    v9 = (!(1.5 == 1.5))
    var v12 int8 = (v11 % ((0) | 1))
    v12 = (v11 | v12)
    v13 := v2[1]
    v1 -= (^(v1))
    if v9 {
        for fi14 := 0; fi14 < 12; fi14++ {
            if len(v4) > 0 {
                v4[0] = (^(fi14))
            }
            _ = fi14
        }
        v10 = "hello"
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 bool = true
        v16 = v16
        _ = fi15
        _ = v16
    }
    if true {
        var v17 int = -734
        switch (^(v17)) {
        case -9223372036854775808:
            v1 += v1
        case 1:
            var v18 float64 = (3.14159 - -2.25)
            _ = v18
        case -1:
            fmt.Println(v12, v9, v9)
        default:
            v12 = v12
        }
        _ = v17
    }
    v10 = (((v10 + v13) + v10) + ("foo bar" + v10))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
