package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    var v3 int8 = 57
    v4 := make(map[string]uint32)
    var v5 int = 1
    _ = len(v4)
    var v6 uint8 = 42
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(fi7)
        fmt.Println(v3)
        v3 = 62
        _ = fi7
    }
    fmt.Println(v6, v6, v3)
    v8 := []int{}
    v1[2147483647] = v5
    v3 += (v3 - v3)
    if ((^(v6)) == (v6 - (v6 | (v6 >> ((2) & 63))))) {
        fmt.Println(v5)
        switch ("" + "hello") {
        case "hello":
            v6 = (v6 ^ 255)
        case "a":
            v8 = append(v8, 4294967295)
        case "\n\t":
            v6 += (v6 & 255)
        default:
            var v9 uint64 = 42
            _ = v9
        }
    } else {
        v8 = append(v8, -1)
        v6 += 1
        var v10 int64 = 9223372036854775807
        _ = v10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
