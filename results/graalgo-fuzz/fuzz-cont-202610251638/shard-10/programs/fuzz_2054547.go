package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 327
    v2 := make(map[int]int)
    var v3 uint64 = 1
    var v4 uint = 42
    var v5 string = ("hello" + ("\n\t" + ("foo bar" + "")))
    v4 += (v4 / (((v4 << ((v4) & 63))) | 1))
    v6 := []int{}
    v3 -= v3
    v6 = append(v6, -9223372036854775808)
    var v7 int16 = 0
    var v8 float32 = 0.0
    v2[-2147483648] = 0
    for fi9 := 0; fi9 < 3; fi9++ {
        delete(v2, (fi9 - (fi9 & fi9)))
        var v10 int32 = 1
        _ = fi9
        _ = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
