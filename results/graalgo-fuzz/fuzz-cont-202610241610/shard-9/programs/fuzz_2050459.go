package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 0
    var v2 uint = 207
    var v3 uint8 = 2
    if ((true && (false || false)) && (905 >= 319)) {
        fmt.Println(v2)
    }
    fmt.Println(v3)
    v4 := []int{}
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 bool = false
        _ = fi5
        _ = v6
    }
    v1 = -32768
    var v7 uint64 = uint64(v3)
    v2 = (v2 >> ((v2) & 63))
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
