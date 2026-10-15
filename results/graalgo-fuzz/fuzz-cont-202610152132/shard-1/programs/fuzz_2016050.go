package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    v3 := make(map[int]string)
    var v4 uint64 = 2
    v4 -= 636
    v4 -= (v4 + 4294967295)
    v4 -= v4
    v4 += v4
    for fi5 := 0; fi5 < 8; fi5++ {
        if true {
            fmt.Println(v4)
            v4 -= v4
        }
        _ = fi5
    }
    fmt.Println(v4)
    _ = len(v2)
    v4 = 249
    var v6 int64 = 962
    var v7 uint16 = 776
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
