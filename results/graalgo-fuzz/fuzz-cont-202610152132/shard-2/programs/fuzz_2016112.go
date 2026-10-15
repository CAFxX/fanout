package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]bool)
    var v3 int64 = 42
    v4 := make(map[int]int)
    v5 := make(map[uint32]uint32)
    var v6 int64 = 0
    if false {
        var v7 uint32 = 888
        _ = v7
    }
    v8 := GenS1{F0: 142}
    if true {
        v3 -= v3
        v6 = v3
        var v9 uint = 0
        _ = v9
    }
    v6 = (v6 - v6)
    v10 := v4[-739]
    v6 -= v6
    var v11 int32 = -1
    switch ("" + "a") {
    case "\n\t":
        if (365 != 328) {
            fmt.Println(v3, v6)
            if false {
                fmt.Println(v3)
            }
        }
    }
    v11 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
