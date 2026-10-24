package main

import (
    "fmt"
)


func main() {
    v1 := []string{"a", "\n\t", "foo bar", ("foo bar" + "\n\t")}
    v2 := make(map[int]uint32)
    var v3 uint8 = 0
    v3 = (v3 * v3)
    _ = len(v2)
    var v4 int64 = -591
    var v5 bool = (-1 == 127)
    fmt.Println(v3, v4)
    v6 := v2[127]
    var v7 float32 = -2.25
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(fi8, v6, v3)
        _ = fi8
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
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
