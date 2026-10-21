package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = true
    var v3 int16 = -1
    v4 := make(map[string]bool)
    v5 := []int{340, 42}
    switch -2147483648 {
    case 42:
        v2 = v2
    case 513:
        var v6 uint16 = 715
        _ = v6
    case 4294967295:
        var v7 float32 = 3.14159
        _ = v7
    default:
        v8 := GenS1{F0: v3}
        _ = v8
    }
    v3 = (v3 / ((v3) | 1))
    var v9 int16 = (v3 >> ((13) & 63))
    var v10 string = ("a" + "a")
    if len(v5) > 0 {
        v5[0] = 127
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        switch 32767 {
        case 2147483647:
            switch 42 {
            case 9223372036854775807:
                v10 = v10
            case 762:
                v12 := GenS1{F0: (^(v3))}
                _ = v12
            case -561:
                for fi13 := 0; fi13 < 3; fi13++ {
                    _ = len(v5)
                    fmt.Println(fi11)
                    _ = fi13
                }
            }
        }
        _ = fi11
    }
    delete(v1, 2147483647)
    var v14 float32 = ((3.14159 + (1.5 + 1e10)) - (((-2.25 * (0.0 - 1.5)) / 1.0) - 1.5))
    fmt.Println(v10)
    for fi15 := 0; fi15 < 18; fi15++ {
        v9 += v9
        if len(v5) > 0 {
            v5[0] = (fi15 << ((601) & 63))
        }
        _ = fi15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(v9, v2)
        _ = fi16
    }
    var v17 int64 = -271
    v2 = (9223372036854775807 != v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
