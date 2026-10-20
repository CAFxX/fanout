package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]string)
    var v2 int32 = 2147483647
    v3 := []int{218}
    var v4 int32 = (v2 ^ v2)
    var v5 uint64 = 1
    switch 32767 {
    case -872:
        var v6 bool = (true || true)
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int64 = -554
        _ = fi7
        _ = v8
    }
    fmt.Println(v4, v4)
    if (-683 != 851) {
        v4 = 1
        fmt.Println(v2)
        v2 = 42
    } else {
        var v9 uint64 = v5
        v3 = append(v3, 127)
        _ = v9
    }
    var v10 string = ""
    fmt.Println(v10)
    for fi11 := 0; fi11 < 13; fi11++ {
        v4 = 330
        _ = fi11
    }
    v2 += -2147483648
    var v12 uint16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
