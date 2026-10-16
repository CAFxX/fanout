package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := []int{32768, 218, -760, 235}
    var v2 int32 = 2147483647
    v3 := make(map[string]bool)
    var v4 int16 = -32768
    var v5 int32 = v2
    v4 += (v4 ^ (v4 % ((v4) | 1)))
    switch "\n\t" {
    case "foo bar":
        v4 -= (v4 | v4)
    case "":
        var v6 bool = (((1.5 < -2.25) && ("hello" != (("\n\t" + "") + "a"))) || ((!false) || false))
        _ = v6
    case "hello":
        _ = len(v3)
    default:
        var v7 int8 = 42
        _ = v7
    }
    var v8 int32 = 420
    var v9 int = -2147483648
    var v10 int = (^(v9))
    v2 = (v5 << ((1) & 63))
    if false {
        var v11 uint16 = 2
        _ = v11
    }
    switch (v10 + (v10 | (v9 & (^(v9))))) {
    case 128:
        fmt.Println(v5, v10)
    }
    var v12 bool = true
    for fi13 := 0; fi13 < 17; fi13++ {
        v4 = -249
        _ = len(v1)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi13
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := GenS1{F0: 0, F1: 255}
        _ = fi14
        _ = v15
    }
    var v16 uint = 2
    var v17 bool = v12
    v1 = append(v1, v9)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
