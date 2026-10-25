package main

import (
    "fmt"
)

func genFn0() bool {
    var v6 uint64 = 1
    v6 = (v6 & v6)
    _ = v6
    return (false || (false || false))
}

func main() {
    var v1 float32 = (1.5 + 0.1)
    v2 := make(map[uint32]bool)
    var v3 string = "foo bar"
    var v4 float64 = 3.14159
    v5 := []uint32{65535}
    v1 = v1
    fmt.Println(v1, v4)
    var v7 uint16 = 923
    v7 = (^(v7))
    v4 += v4
    var v8 uint32 = 2147483647
    v7 -= 816
    switch v3 {
    case "a":
        for fi9 := 0; fi9 < 18; fi9++ {
            v1 += 0.0
            _ = fi9
        }
    }
    var v10 int64 = -9223372036854775808
    if false {
        v10 = -1
        fmt.Println(v7, v4, v8)
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
