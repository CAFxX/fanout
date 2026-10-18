package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -942
    var v3 int = 42
    v4 := make(map[int]bool)
    var v5 uint16 = 2
    var v6 bool = ((!(v3 >= 267)) || (-2.25 != (1e10 / 1e10)))
    var v7 int64 = v2
    v2 = (v7 * v7)
    for fi8 := 0; fi8 < 5; fi8++ {
        switch "\n\t" {
        case "\n\t":
            v9 := []int{v3, -2147483648, -9223372036854775808}
            _ = v9
        case "a":
            v10 := GenS1{F0: ("a" + "a"), F1: 127, F2: 920}
            _ = v10
        }
        _ = fi8
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v7 = -9223372036854775808
        _ = fi11
    }
    v12 := GenS1{F0: (("\n\t" + "") + ("hello" + "foo bar")), F1: 120, F2: 42}
    if v6 {
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 uint32 = 972
            var v15 uint8 = 128
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v4[(^(v3))] = v6
        _ = fi16
    }
    v17 := []int{v3, v3, 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
