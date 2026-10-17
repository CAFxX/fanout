package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]uint32)
    v3 := []int{-617, -2147483648, 42}
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]string)
    var v6 uint32 = 2147483648
    if true {
        v7 := GenS1{F0: 42, F1: 345, F2: 2}
        _ = v7
    }
    var v8 uint32 = v6
    var v9 int = 159
    v8 = (v8 | 911)
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            v6 += (v8 | v6)
            v6 -= 4294967295
            _ = fi11
        }
        _ = fi10
    }
    v8 = v8
    fmt.Println(v6)
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int = (-(fi12))
        _ = fi12
        _ = v13
    }
    switch ("hello" + "a") {
    case "foo bar":
        v6 += v8
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
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
