package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]int)
    v3 := []string{}
    fmt.Println(v1)
    _ = len(v3)
    var v4 uint16 = 893
    v1 = 42
    fmt.Println(v4, v4)
    v2[65535] = 127
    var v5 uint64 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
