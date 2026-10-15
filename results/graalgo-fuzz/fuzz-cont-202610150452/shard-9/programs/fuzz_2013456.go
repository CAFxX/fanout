package main

import (
    "fmt"
)


func main() {
    var v1 float64 = (-2.25 * (0.0 + ((-2.25 / (3.14159 + (3.14159 / (3.14159 / 1.0)))) - ((0.1 + float64(0.1)) - -2.25))))
    v2 := []string{}
    v3 := []int{-2147483648, 871}
    if len(v2) > 0 {
        v2[0] = "a"
    }
    v1 = float64(0.0)
    fmt.Println(v1)
    var v4 bool = (!((574 <= 302) && false))
    var v5 uint8 = 1
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 int8 = 100
        var v8 uint32 = 287
        v1 = 0.0
        _ = fi6
        _ = v7
        _ = v8
    }
    if (!(!v4)) {
        var v9 uint64 = 2
        v1 = 0.0
        _ = v9
    }
    fmt.Println(v4, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
