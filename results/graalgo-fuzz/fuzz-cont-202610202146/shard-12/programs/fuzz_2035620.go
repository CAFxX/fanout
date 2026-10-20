package main

import (
    "fmt"
)

func genFn0(p0 bool, p1 uint64) int8 {
    v7 := []int{-705}
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int64 = 0
        p0 = ((p1 ^ p1) > (p1 - 837))
        _ = fi8
        _ = v9
    }
    fmt.Println(p1, p1)
    p0 = (!(182 <= 2))
    if len(v7) > 0 {
        v7[0] = 128
    }
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    var v4 int16 = 42
    v5 := []uint32{435, 516, 193, 2}
    var v6 bool = (!(false && true))
    var v10 int8 = 1
    switch "a" {
    case "hello":
        v4 = v4
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v6 = (!false)
    var v11 float64 = 3.14159
    var v12 bool = v6
    v10 += 127
    fmt.Println(v6, v1)
    if v6 {
        fmt.Println(v6, v10, v11)
    }
    var v13 uint64 = 1
    v6 = v12
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint8 = 0
        var v16 float64 = -2.25
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int64 = 9223372036854775807
    v6 = (18446744073709551615 < 2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
