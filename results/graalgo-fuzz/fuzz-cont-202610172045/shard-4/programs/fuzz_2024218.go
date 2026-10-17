package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int32 = -1
    var v2 int8 = 0
    var v3 bool = (!true)
    var v4 int8 = (v2 >> ((1) & 63))
    var v5 uint16 = 728
    var v6 bool = false
    v7 := GenS1{F0: 0, F1: v1}
    var v8 bool = false
    v8 = v6
    v8 = ((!(v4 <= 42)) && (!v6))
    fmt.Println(v2, v2)
    if (!(!(42 <= 42))) {
        var v9 int64 = -207
        _ = v9
    } else {
        if true {
            var v10 float64 = 0.0
            fmt.Println(v2, v1, v10)
            var v11 int = -9223372036854775808
            _ = v10
            _ = v11
        } else {
            switch 1 {
            case -990:
                v8 = v6
            case -1:
                if (!false) {
                    switch "" {
                    case "":
                        v12 := []int{-500, 42}
                        _ = v12
                    }
                }
            case 42:
                v13 := GenS1{F0: 128, F1: (v1 & (v1 - (^(v1))))}
                _ = v13
            default:
                v14 := []int{127, -9223372036854775808}
                _ = v14
            }
            v15 := []int{42, -584}
            _ = v15
        }
        var v16 int8 = (v2 << ((42) & 63))
        _ = v16
    }
    v17 := []int{0, 489}
    fmt.Println(v2)
    var v18 uint32 = 4294967295
    v19 := GenS2{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
