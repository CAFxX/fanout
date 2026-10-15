package main

import (
    "fmt"
)


func main() {
    v1 := []int{43}
    v2 := make(map[int]uint32)
    var v3 int8 = -1
    var v4 int32 = -599
    switch "" {
    case "\n\t":
        v4 = -1
    case "hello":
        var v5 uint64 = 11
        _ = v5
    default:
        var v6 uint32 = 674
        _ = v6
    }
    fmt.Println(v4)
    if len(v1) > 0 {
        v1[0] = 128
    }
    v1 = append(v1, 127)
    v1 = append(v1, 32767)
    fmt.Println(v4, v3)
    v3 = v3
    v3 = (v3 | v3)
    _ = len(v1)
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int64 = -1
        _ = fi7
        _ = v8
    }
    v4 -= (v4 | v4)
    v4 += 0
    if true {
        var v9 uint16 = 255
        _ = v9
    }
    v3 += v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
