package main

import (
    "fmt"
)


func main() {
    var v1 int = 128
    var v2 int8 = -1
    v3 := make(map[string]bool)
    delete(v3, "\n\t")
    v4 := v3["a"]
    if (294 > 613) {
        v4 = true
    }
    v5 := []int{(v1 & v1), 128}
    v4 = v4
    var v6 bool = (1 < 1)
    var v7 int32 = 1
    if len(v5) > 0 {
        v5[0] = (v1 >> ((482) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
