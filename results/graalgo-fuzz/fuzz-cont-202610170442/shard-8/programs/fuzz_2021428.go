package main

import (
    "fmt"
)

func genFn0() uint16 {
    var v4 uint32 = 430
    var v5 bool = (!(1.5 > 0.0))
    v5 = false
    _ = v4
    _ = v5
    return 513
}
func genFn1(p0 int, p1 string, p2 int32) int8 {
    v6 := []int{p0, (p0 % ((p0) | 1))}
    switch "foo bar" {
    case "\n\t":
        var v7 uint8 = 2
        _ = v7
    case "a":
        var v8 uint32 = 2147483647
        _ = v8
    }
    _ = len(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 127
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 838
    v3 := []int{-219, 962, -694, 128}
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 uint64 = 562
        _ = fi9
        _ = v10
    }
    var v11 int = 4294967295
    if false {
        var v12 bool = false
        var v13 string = "a"
        v1[v11] = (9223372036854775807 == 42)
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 15; fi14++ {
            v2 = v2
            _ = fi14
        }
        var v15 string = ""
        _ = v15
    }
    _ = len(v1)
    var v16 int8 = 42
    var v17 int64 = 0
    fmt.Println(v16, v17)
    switch 42 {
    case -2147483648:
        v2 = (^(v2))
    }
    fmt.Println(v11, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
