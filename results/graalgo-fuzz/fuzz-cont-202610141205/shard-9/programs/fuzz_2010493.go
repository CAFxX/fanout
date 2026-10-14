package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 42
    var v2 string = "hello"
    var v3 uint = 65535
    if false {
        for fi4 := 0; fi4 < 8; fi4++ {
            v3 += v3
            _ = fi4
        }
        v1 += v1
        var v5 int32 = 1
        _ = v5
    }
    v6 := []int{251, 1, 4294967295, 127}
    fmt.Println(v1, v3)
    var v7 int8 = v1
    fmt.Println(v7, v7, v2)
    v2 = v2
    v1 -= (v7 - (^(v1)))
    for fi8 := 0; fi8 < 15; fi8++ {
        v7 -= (v7 + (v7 & v1))
        _ = fi8
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
