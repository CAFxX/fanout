package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]uint32)
    v2 := []int{9223372036854775807, -84}
    var v3 int = 32768
    var v4 uint = 4294967295
    v5 := make(map[string]uint32)
    v6 := make(map[int]bool)
    if true {
        v4 = (v4 + v4)
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int64 = -1
        var v9 int16 = 46
        var v10 string = "a"
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v3 = v3
    v4 += (^(v4))
    fmt.Println(v4)
    var v11 uint8 = 1
    v12 := v6[(v3 / (((^(v3))) | 1))]
    var v13 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
