package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]int)
    v2 := []int{-905, -1, 9223372036854775807, 0}
    var v3 int = 42
    var v4 uint = 339
    for fi5 := 0; fi5 < 7; fi5++ {
        v4 -= (^(v4))
        _ = fi5
    }
    var v6 uint = 18446744073709551615
    var v7 int64 = -1
    v7 = 0
    if false {
        var v8 int32 = 42
        v4 += 444
        _ = v8
    }
    fmt.Println(v4, v4, v4)
    var v9 float32 = 3.14159
    for fi10 := 0; fi10 < 2; fi10++ {
        v3 = (fi10 * 127)
        var v11 int = v3
        _ = fi10
        _ = v11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v9
}
