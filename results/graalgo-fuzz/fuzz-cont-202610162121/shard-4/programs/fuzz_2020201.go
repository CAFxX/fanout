package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
func genFn0(p0 uint64) uint32 {
    p0 += p0
    var v4 uint16 = 42
    var v5 uint = 0
    _ = p0
    _ = v4
    _ = v5
    return 2147483648
}

func main() {
    var v1 uint = 65535
    var v2 uint64 = 2
    var v3 uint = (^(v1))
    fmt.Println(v1, v3, v1)
    v6 := []int{}
    if true {
        fmt.Println(v3, v2, v2)
        v3 = v1
    }
    var v7 int8 = -128
    v7 = (v7 % ((v7) | 1))
    var v8 float32 = 0.1
    var v9 uint = v1
    v3 = v3
    switch "" {
    case "a":
        if len(v6) > 0 {
            v6[0] = 32768
        }
    }
    var v10 uint64 = v2
    var v11 uint32 = 65535
    v1 -= (^(v3))
    var v12 bool = (false || (1 == 32767))
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := GenS1{F0: fi14, F1: (v11 >> ((v9) & 63))}
            var v16 int = (fi14 ^ fi13)
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 float64 = 1.5
        _ = fi13
        _ = v17
    }
    v7 = -1
    var v18 string = ""
    var v19 bool = v12
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
