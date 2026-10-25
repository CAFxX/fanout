package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }

func main() {
    v1 := []string{"a", ""}
    v2 := []uint32{42, 42}
    var v3 int = 2147483647
    var v4 uint8 = 255
    var v5 int8 = 127
    v6 := GenS1{F0: -311, F1: 814}
    if false {
        if false {
            v5 = 51
        }
    } else {
        v5 -= (v5 % ((v5) | 1))
        v5 += 0
        var v7 bool = true
        _ = v7
    }
    switch "hello" {
    case "":
        v1 = append(v1, ("foo bar" + "\n\t"))
    }
    var v8 int32 = -1
    v4 += (v4 * 2)
    v5 = (v5 & v5)
    var v9 float32 = 1.5
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint = 1
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 float32 = (v9 / -2.25)
        for fi14 := 0; fi14 < 8; fi14++ {
            v15 := GenS1{F0: -2147483648, F1: 18446744073709551615}
            _ = len(v2)
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v9, v8, v9)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
