package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 705
    var v3 uint32 = 42
    v4 := make(map[string]bool)
    if true {
        v2 = v2
        v3 = (v3 >> ((341) & 63))
    } else {
        fmt.Println(v3)
        for fi5 := 0; fi5 < 10; fi5++ {
            v6 := GenS1{F0: 2147483647}
            _ = fi5
            _ = v6
        }
    }
    var v7 bool = (!(0 > 140))
    if v7 {
        if v7 {
            v8 := GenS1{F0: -2147483648}
            _ = v8
        } else {
            v9 := []int{128, -9223372036854775808, 508, 2147483647}
            v3 = 4294967295
            _ = v9
        }
        fmt.Println(v2, v7)
    }
    v3 += v3
    var v10 int16 = -1
    var v11 int8 = -1
    for fi12 := 0; fi12 < 1; fi12++ {
        v10 = v10
        _ = fi12
    }
    v7 = v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
