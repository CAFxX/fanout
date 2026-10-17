package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint }

func main() {
    var v1 uint64 = 0
    var v2 string = "a"
    v3 := []uint32{4294967295}
    var v4 bool = true
    var v5 uint64 = (v1 << ((255) & 63))
    _ = len(v3)
    _ = len(v3)
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    var v6 uint32 = 1
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            switch "foo bar" {
            case "\n\t":
                v9 := GenS1{F0: fi7, F1: (1 == 2), F2: 18446744073709551615}
                _ = v9
            case "hello":
                var v10 int8 = 106
                _ = v10
            case "":
                _ = len(v3)
            }
            _ = fi8
        }
        _ = fi7
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS1{F0: 1, F1: v4, F2: 255}
        _ = fi11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
