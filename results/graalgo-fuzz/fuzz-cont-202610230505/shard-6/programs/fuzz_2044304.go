package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }

func main() {
    var v1 int8 = 42
    v2 := []uint32{225, 4294967295}
    v3 := []int{498, -1}
    var v4 uint16 = 0
    for fi5 := 0; fi5 < 19; fi5++ {
        v1 += (v1 >> ((385) & 63))
        if len(v3) > 0 {
            v3[0] = fi5
        }
        _ = fi5
    }
    v6 := GenS1{F0: 648, F1: (9223372036854775807 > 42)}
    var v7 string = "foo bar"
    var v8 int32 = -2147483648
    fmt.Println(v8, v1, v7)
    for fi9 := 0; fi9 < 15; fi9++ {
        v1 += v1
        v8 = 42
        _ = fi9
    }
    v4 -= (v4 % (((^(v4))) | 1))
    for fi10 := 0; fi10 < 12; fi10++ {
        fmt.Println(v8, fi10, fi10)
        switch v7 {
        case "a":
            var v11 uint8 = 42
            _ = v11
        }
        _ = fi10
    }
    fmt.Println(v4, v4)
    var v12 int8 = v1
    var v13 uint64 = 525
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
