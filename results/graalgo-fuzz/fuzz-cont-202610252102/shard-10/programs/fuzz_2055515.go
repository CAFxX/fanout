package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    var v4 int = -399
    v5 := make(map[string]bool)
    var v6 int8 = 1
    v7 := GenS1{F0: 65535}
    v8 := GenS1{F0: 526}
    var v9 bool = false
    v4 -= -9223372036854775808
    v9 = true
    v10 := v1[v4]
    var v11 int64 = 70
    fmt.Println(v4, v10)
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 uint32 = v10
        var v14 uint64 = 654
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 bool = true
    var v16 string = ("a" + ("a" + "\n\t"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
