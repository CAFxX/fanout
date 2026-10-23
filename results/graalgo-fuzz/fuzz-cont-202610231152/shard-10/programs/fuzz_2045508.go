package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[int]string)
    v2 := []int{}
    var v3 uint32 = 127
    v4 := make(map[string]bool)
    var v5 uint32 = 2147483647
    var v6 int8 = -66
    v7 := GenS1{F0: 42}
    delete(v4, "a")
    v5 -= v5
    if (42 < 4294967295) {
        if false {
            fmt.Println(v6, v3)
            for fi8 := 0; fi8 < 14; fi8++ {
                v5 -= v3
                _ = fi8
            }
        } else {
            var v9 uint = 42
            _ = v9
        }
    }
    if false {
        var v10 uint8 = 101
        _ = v10
    } else {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 int32 = 1
            _ = fi11
            _ = v12
        }
    }
    fmt.Println(v6, v3, v5)
    v3 += (v3 - v5)
    v13 := v4["a"]
    v3 = 803
    fmt.Println(v3)
    v6 += v6
    v5 = 2147483648
    for fi14 := 0; fi14 < 10; fi14++ {
        fmt.Println(v5, v13)
        _ = fi14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
