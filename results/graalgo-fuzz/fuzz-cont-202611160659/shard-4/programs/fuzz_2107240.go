package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint64 = 42
    var v2 int = 32768
    var v3 uint8 = 253
    switch v2 {
    case 703:
        if (1 <= 9223372036854775807) {
            var v4 bool = (98 > -1)
            _ = v4
        } else {
            var v5 uint8 = (^(v3))
            _ = v5
        }
    default:
        var v6 uint8 = (v3 >> ((1) & 63))
        _ = v6
    }
    if (!false) {
        switch "a" {
        case "hello":
            v7 := []int{128, (-(v2)), (v2 / ((v2) | 1)), 1}
            _ = v7
        case "a":
            var v8 uint32 = 525
            _ = v8
        default:
            v9 := []int{v2, (v2 & (-(v2))), v2, 127}
            _ = v9
        }
        v10 := GenS1{F0: v2}
        var v11 uint16 = 423
        _ = v10
        _ = v11
    }
    var v12 uint64 = v1
    v13 := GenS1{F0: v2}
    var v14 int8 = 100
    v15 := []int{340}
    var v16 string = ""
    fmt.Println(v3)
    fmt.Println(v2)
    if false {
        var v17 bool = false
        for fi18 := 0; fi18 < 20; fi18++ {
            v19 := GenS1{F0: v2}
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    var v20 int64 = 9223372036854775807
    var v21 uint64 = 42
    v22 := GenS2{F0: 498}
    var v23 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
