package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }

func main() {
    var v1 int8 = -76
    v2 := make(map[int]string)
    var v3 int32 = -1
    var v4 string = "foo bar"
    v1 = (-(v1))
    if false {
        var v5 uint64 = 964
        v6 := v2[-1]
        _ = v5
        _ = v6
    }
    fmt.Println(v4, v4)
    var v7 int64 = 9223372036854775807
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: (v3 << ((354) & 63)), F1: (v4 + "hello")}
        var v10 uint8 = 0
        for fi11 := 0; fi11 < 17; fi11++ {
            v3 -= (v3 >> ((42) & 63))
            _ = fi11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v12 uint32 = 0
    fmt.Println(v1)
    v7 = v7
    v12 = v12
    fmt.Println(v1, v7, v12)
    var v13 uint16 = 1
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
