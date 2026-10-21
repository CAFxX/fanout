package main

import (
    "fmt"
)


func main() {
    var v1 uint64 = 2
    var v2 uint16 = 32767
    var v3 uint = 2
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 string = "\n\t"
        v5 = v5
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            fmt.Println(v3)
            var v8 uint32 = 2
            var v9 bool = true
            fmt.Println(v8, v2, v8)
            var v10 int8 = -128
            var v11 float64 = ((0.0 * 1e10) * 0.0)
            _ = fi7
            _ = v10
            _ = v11
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v12 := []int{9223372036854775807, 32768}
    var v13 uint32 = 2147483647
    v12 = append(v12, -441)
    fmt.Println(v13, v1)
    var v14 float64 = 0.1
    var v15 uint16 = v2
    switch 32767 {
    case -1:
        if len(v12) > 0 {
            v12[0] = 9223372036854775807
        }
    default:
        v3 = uint(v1)
    }
    if (true && ((true || (!((v3 % ((0) | 1)) > v3))) && false)) {
        var v16 int = -9223372036854775808
        v3 -= (v3 - (v3 - v3))
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
