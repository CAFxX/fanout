package main

import (
    "fmt"
)


func main() {
    var v1 bool = false
    v2 := make(map[string]uint32)
    var v3 float64 = 3.14159
    v1 = (255 > 1)
    v4 := []int{32767}
    var v5 float64 = (-2.25 / (0.0 * v3))
    delete(v2, "foo bar")
    fmt.Println(v3, v5)
    v1 = (2 >= 4294967295)
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint64 = 0
        _ = fi6
        _ = v7
    }
    if (2147483647 >= 128) {
        delete(v2, "\n\t")
        v5 = v3
    }
    var v8 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
