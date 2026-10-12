package main

import (
    "fmt"
)

func genFn0() int64 {
    v6 := []int{42, -9223372036854775808, 128}
    var v7 int64 = -1
    _ = v6
    _ = v7
    return -914
}

func main() {
    var v1 bool = false
    v2 := []string{(("\n\t" + "") + "foo bar"), "hello", "", ""}
    var v3 uint = 0
    var v4 bool = v1
    var v5 int64 = 1
    fmt.Println(v1)
    fmt.Println(v5)
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint32 = 2
        var v10 bool = true
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 = true
    var v11 int64 = 475
    v5 = genFn0()
    var v12 float32 = (1.5 * 3.14159)
    v5 += (v11 * (v5 * 9223372036854775807))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        fmt.Println(v11, v5, v4)
        _ = fi13
    }
    v12 += v12
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
