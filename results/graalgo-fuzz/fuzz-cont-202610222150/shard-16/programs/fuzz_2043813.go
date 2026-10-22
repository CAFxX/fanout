package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]uint32)
    var v2 string = (("foo bar" + "") + "")
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    var v5 int8 = 77
    var v6 int16 = 1
    var v7 int = -895
    v7 -= v7
    var v8 int16 = v6
    var v9 bool = false
    if (!v9) {
        v5 -= v5
    }
    v10 := []int{(v7 - v7)}
    var v11 uint64 = 18446744073709551615
    var v12 uint64 = (^(v11))
    var v13 int64 = 500
    for fi14 := 0; fi14 < 15; fi14++ {
        if (-150 != (v6 - v8)) {
            var v15 uint16 = 0
            _ = v15
        } else {
            v9 = true
        }
        fmt.Println(v11)
        _ = fi14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
