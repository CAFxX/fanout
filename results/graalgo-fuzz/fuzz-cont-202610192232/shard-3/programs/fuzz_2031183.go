package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
func genFn0(p0 int16, p1 uint64, p2 string) int16 {
    v7 := []int{42, 4294967295, -14, 468}
    v8 := GenS1{F0: 146, F1: 42}
    var v9 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (p0 * (^(p0)))
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]uint32)
    var v3 uint64 = v1
    var v4 uint8 = 0
    var v5 int16 = 0
    var v6 bool = true
    var v10 int = 128
    if v6 {
        v5 = (v5 >> ((42) & 63))
        v10 += -9223372036854775808
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 uint8 = (v4 - v4)
        _ = fi11
        _ = v12
    }
    fmt.Println(v3, v10, v5)
    fmt.Println(v10, v6)
    v4 += 2
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS1{F0: v4, F1: 1}
        v15 := GenS1{F0: v4, F1: 42}
        v4 = v4
        _ = fi13
        _ = v14
        _ = v15
    }
    switch ("foo bar" + "foo bar") {
    case "":
        fmt.Println(v5)
    default:
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
