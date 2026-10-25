package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    var v3 uint64 = 42
    switch 763 {
    case 9223372036854775807:
        v4 := []int{2147483647}
        _ = v4
    default:
        for fi5 := 0; fi5 < 17; fi5++ {
            v6 := []int{32768, (fi5 >> ((206) & 63)), (fi5 ^ (fi5 >> ((752) & 63)))}
            _ = fi5
            _ = v6
        }
    }
    fmt.Println(v3)
    v3 -= 1
    v3 = (v3 << ((18446744073709551615) & 63))
    switch 847 {
    case 2147483647:
        v7 := []int{9223372036854775807}
        _ = v7
    case 128:
        fmt.Println(v3)
    case -802:
        var v8 float64 = float64(3.14159)
        _ = v8
    default:
        fmt.Println(v3)
    }
    var v9 uint64 = 129
    v10 := v2[42]
    for fi11 := 0; fi11 < 8; fi11++ {
        v3 = (^(v9))
        _ = fi11
    }
    fmt.Println(v10, v10)
    v3 = v3
    v9 = v9
    switch 689 {
    case 128:
        v12 := GenS1{F0: 42}
        _ = v12
    case -822:
        fmt.Println(v3, v9)
    }
    v13 := GenS1{F0: 32767}
    v14 := GenS1{F0: 32768}
    var v15 uint64 = v3
    var v16 int64 = 46
    var v17 int64 = (v16 + 0)
    for fi18 := 0; fi18 < 15; fi18++ {
        fmt.Println(v10)
        _ = fi18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
