package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 bool = (!false)
    v2 := []int{-359, 1, 0}
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    var v5 int16 = 1000
    var v6 int16 = (v5 & v5)
    if (false && v1) {
        v7 := GenS1{F0: 4294967295}
        for fi8 := 0; fi8 < 1; fi8++ {
            v2 = append(v2, -9223372036854775808)
            _ = fi8
        }
        var v9 bool = v1
        _ = v7
        _ = v9
    }
    var v10 uint = 0
    if (3.14159 < 3.14159) {
        v11 := GenS1{F0: 714}
        var v12 float32 = 0.0
        v5 = v6
        _ = v11
        _ = v12
    } else {
        v10 -= v10
        switch ("\n\t" + ("\n\t" + "hello")) {
        case "\n\t":
            fmt.Println(v5)
        case "a":
            v10 += (v10 + 627)
        }
    }
    var v13 int64 = 42
    delete(v3, -167)
    v14 := GenS1{F0: -2147483648}
    if v1 {
        v5 = -1
        v15 := GenS1{F0: 32767}
        _ = v15
    }
    var v16 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
